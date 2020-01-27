package task

import (
	"context"
	"fmt"
	"log"

	apiextv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	apijson "k8s.io/apimachinery/pkg/util/json"
	"k8s.io/client-go/discovery"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"

	"github.com/kudobuilder/kudo/pkg/apis/kudo/v1beta1"
	"github.com/kudobuilder/kudo/pkg/engine/health"
)

// ApplyTask will apply a set of given resources to the cluster. See Run method for more details.
type ApplyTask struct {
	Name      string
	Resources []string
}

// Run method for the ApplyTask. Given the task context, it renders the templates using context parameters
// creates runtime objects and enhances them, and applies them using the controller client. Finally,
// resources are checked for health.
func (at ApplyTask) Run(ctx Context) (bool, error) {
	// 1. - Render task templates -
	rendered, err := render(at.Resources, ctx)
	if err != nil {
		return false, fatalExecutionError(err, taskRenderingError, ctx.Meta)
	}

	// 2. - Enhance them with metadata -
	enhanced, err := enhance(rendered, ctx.Meta, ctx.Enhancer)
	if err != nil {
		return false, fatalExecutionError(err, taskEnhancementError, ctx.Meta)
	}

	// 3. - Apply them using the client -
	applied, err := apply(enhanced, ctx.Client)
	if err != nil {
		return false, err
	}

	// 4. - Check health for all resources -
	err = isHealthy(applied)
	if err != nil {
		// so far we do not distinguish between unhealthy resources and other errors that might occur during a health check
		// an error during a health check is not treated task execution error
		log.Printf("TaskExecution: %v", err)
		return false, nil
	}
	return true, nil
}

// apply method takes a slice of k8s object and applies them using passed client. If an object
// doesn't exist it will be created. An already existing object will be patched.
func apply(rr []runtime.Object, c client.Client) ([]runtime.Object, error) {
	applied := make([]runtime.Object, 0)

	for _, r := range rr {
		existing := r.DeepCopyObject()

		key, err := ObjectKeyFromObject(r)
		if err != nil {
			return nil, err
		}

		err = c.Get(context.TODO(), key, existing)

		switch {
		case apierrors.IsNotFound(err): // create resource if it doesn't exist
			err = c.Create(context.TODO(), r)
			// c.Create always overrides the input, in this case, the object that had previously set GVK loses it (at least for integration tests)
			// and this was causing problems in health module
			// with error failed to convert *unstructured.Unstructured to *v1.Deployment: Object 'Kind' is missing in 'unstructured object has no kind'
			// so re-setting the GVK here to be sure
			// https://github.com/kubernetes/kubernetes/issues/80609
			r.GetObjectKind().SetGroupVersionKind(existing.GetObjectKind().GroupVersionKind())
			if err != nil {
				return nil, err
			}
			applied = append(applied, r)
		case err != nil: // raise any error other than StatusReasonNotFound
			return nil, err
		default: // update existing resource
			err := patch(r, c)
			if err != nil {
				return nil, err
			}
			applied = append(applied, r)
		}
	}

	return applied, nil
}

// ObjectKeyFromObject method wraps client.ObjectKeyFromObject method by additionally checking if passed object is
// a cluster-scoped resource (e.g. CustomResourceDefinition, ClusterRole etc.) and removing the namespace from the
// key since cluster-scoped resources are not namespaced.
func ObjectKeyFromObject(r runtime.Object) (client.ObjectKey, error) {
	key, err := client.ObjectKeyFromObject(r)
	if err != nil {
		return client.ObjectKey{}, fmt.Errorf("failed to get an object key from object %v: %v", r.GetObjectKind(), err)
	}

	// if the resource is cluster-scoped we need to clear then namespace from the key
	namespaced, err := isNamespaced(r)
	if err != nil {
		return client.ObjectKey{}, fmt.Errorf("failed to determine if the resource %v is cluster-scoped: %v", r.GetObjectKind(), err)
	}

	if !namespaced { // not namespaced means cluster-scoped!
		key.Namespace = ""
	}
	return key, nil
}

// isNamespaced method return true if given runtime.Object is a namespaced (not cluster-scoped) resource
func isNamespaced(r runtime.Object) (bool, error) {
	namespaced, err := namespacedGVKResources()
	if err != nil {
		return false, err
	}

	gvk := r.GetObjectKind().GroupVersionKind()
	return namespaced[gvk], nil
}

// patch calls update method on kubernetes client to make sure the current resource reflects what is on server
//
// an obvious optimization here would be to not patch when objects are the same, however that is not easy
// kubernetes native objects might be a problem because we cannot just compare the spec as the spec might have extra fields
// and those extra fields are set by some kubernetes component
// because of that for now we just try to apply the patch every time
// it mutates the object passed in to be consistent with the kubernetes client behavior
func patch(newObj runtime.Object, c client.Client) error {
	newObjJSON, _ := apijson.Marshal(newObj)
	key, _ := client.ObjectKeyFromObject(newObj)
	_, isUnstructured := newObj.(runtime.Unstructured)
	_, isCRD := newObj.(*apiextv1beta1.CustomResourceDefinition)

	if isUnstructured || isCRD || isKudoType(newObj) {
		// strategic merge patch is not supported for these types, falling back to merge patch
		err := c.Patch(context.TODO(), newObj, client.ConstantPatch(types.MergePatchType, newObjJSON))
		if err != nil {
			return fmt.Errorf("failed to apply merge patch to object %s/%s: %w", key.Namespace, key.Name, err)
		}
	} else {
		err := c.Patch(context.TODO(), newObj, client.ConstantPatch(types.StrategicMergePatchType, newObjJSON))
		if err != nil {
			return fmt.Errorf("failed to apply StrategicMergePatch to object %s/%s: %w", key.Namespace, key.Name, err)
		}
	}

	return nil
}

func isKudoType(object runtime.Object) bool {
	_, isOperator := object.(*v1beta1.OperatorVersion)
	_, isOperatorVersion := object.(*v1beta1.Operator)
	_, isInstance := object.(*v1beta1.Instance)
	return isOperator || isOperatorVersion || isInstance
}

func isHealthy(ro []runtime.Object) error {
	for _, r := range ro {
		err := health.IsHealthy(r)
		if err != nil {
			key, _ := client.ObjectKeyFromObject(r)
			return fmt.Errorf("object %s/%s is NOT healthy: %w", key.Namespace, key.Name, err)
		}
	}
	return nil
}

// namespacedGVKResources method uses the discovery client to fetch all API resources (with Groups and Versions)
// and returns a mapping of GVK -> isNamespaced. Useful when deciding whether a given object is a cluster-scoped resource.
func namespacedGVKResources() (map[schema.GroupVersionKind]bool, error) {
	namespaced := map[schema.GroupVersionKind]bool{}

	restCfg, err := config.GetConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch cluster REST config: %v", err)
	}
	// Use a config object to create a discovery client
	discClient, err := discovery.NewDiscoveryClientForConfig(restCfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create a discovery client: %v", err)
	}
	// Fetch namespaced API resources
	_, apiResources, err := discClient.ServerGroupsAndResources()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch server groups and resources: %v", err)
	}

	for _, rr := range apiResources {
		gv, err := schema.ParseGroupVersion(rr.GroupVersion)
		if err != nil {
			continue
		}
		for _, r := range rr.APIResources {
			gvk := gv.WithKind(r.Kind)
			namespaced[gvk] = r.Namespaced
		}
	}

	return namespaced, nil
}
