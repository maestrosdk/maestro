package plan

import (
	"encoding/json"
	"fmt"
	maestrov1alpha1 "github.com/maestrosdk/maestro/pkg/apis/maestro/v1alpha1"
	"github.com/spf13/cobra"
	"github.com/xlab/treeprint"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func NewPlanStatusCmd() *cobra.Command {
	statusCmd := &cobra.Command{
		//Args:  cobra.ExactArgs(1),
		Use:   "status",
		Short: "Shows the status of a particular plan of an instance.",
		Long: `
	# View plan status
	maestroctl plan status --instance=<instanceName>`,
		Run: planStatusCmd,
	}

	statusCmd.Flags().StringVar(&instance, "instance", "", "The instance name available from 'kubectl get instances'")
	statusCmd.Flags().StringVar(&kubeConfig, "kubeconfig", "", "The file path to kubernetes configuration file; defaults to $HOME/.kube/config")
	statusCmd.Flags().StringVar(&namespace, "namespace", "default", "The namespace where the instance is running.")

	return statusCmd
}

func planStatusCmd(cmd *cobra.Command, args []string) {

	instanceFlag, err := cmd.Flags().GetString("instance")
	if err != nil || instanceFlag == "" {
		log.Printf("Flag Error: %v", err)
	}

	mustKubeConfig()

	_, err = cmd.Flags().GetString("kubeconfig")
	if err != nil || instanceFlag == "" {
		log.Printf("Flag Error: %v", err)
	}

	err = planStatus()
	if err != nil {
		log.Printf("Error: %v", err)
	}
}

func planStatus() error {

	tree := treeprint.New()

	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		return err
	}

	//  Create a Dynamic Client to interface with CRDs.
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return err
	}

	instancesGVR := schema.GroupVersionResource{
		Group:    "maestro.k8s.io",
		Version:  "v1alpha1",
		Resource: "instances",
	}

	instObj, err := dynamicClient.Resource(instancesGVR).Namespace(namespace).Get(instance, metav1.GetOptions{})
	if err != nil {
		return err
	}
	mInstObj, err := instObj.MarshalJSON()

	instance := maestrov1alpha1.Instance{}

	err = json.Unmarshal(mInstObj, &instance)
	if err != nil {
		return err
	}

	frameworkVersionNameOfInstance := instance.Spec.FrameworkVersion.Name

	frameworkGVR := schema.GroupVersionResource{
		Group:    "maestro.k8s.io",
		Version:  "v1alpha1",
		Resource: "frameworkversions",
	}

	//  List all of the Virtual Services.
	frameworkObj, err := dynamicClient.Resource(frameworkGVR).Namespace(namespace).Get(frameworkVersionNameOfInstance, metav1.GetOptions{})
	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}
	mFrameworkObj, err := frameworkObj.MarshalJSON()

	framework := maestrov1alpha1.FrameworkVersion{}

	err = json.Unmarshal(mFrameworkObj, &framework)
	if err != nil {
		return err
	}

	planExecutionsGVR := schema.GroupVersionResource{
		Group:    "maestro.k8s.io",
		Version:  "v1alpha1",
		Resource: "planexecutions",
	}

	activePlanObj, err := dynamicClient.Resource(planExecutionsGVR).Namespace(namespace).Get(instance.Status.ActivePlan.Name, metav1.GetOptions{})
	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}
	mPlanObj, err := activePlanObj.MarshalJSON()

	activePlanType := maestrov1alpha1.PlanExecution{}

	err = json.Unmarshal(mPlanObj, &activePlanType)
	if err != nil {
		return err
	}

	rootDisplay := fmt.Sprintf("%s (Framework-Version: \"%s\" Active-Plan: \"%s\")", instance.Name, instance.Spec.FrameworkVersion.Name, instance.Status.ActivePlan.Name)
	rootBranchName := tree.AddBranch(rootDisplay)

	for name, plan := range framework.Spec.Plans {
		if name == activePlanType.Spec.PlanName {
			planDisplay := fmt.Sprintf("Plan %s (%s strategy) [%s]", name, plan.Strategy, activePlanType.Status.State)
			planBranchName := rootBranchName.AddBranch(planDisplay)
			for _, phase := range activePlanType.Status.Phases {
				phaseDisplay := fmt.Sprintf("Phase %s (%s strategy) [%s]", phase.Name, phase.Strategy, phase.State)
				phaseBranchName := planBranchName.AddBranch(phaseDisplay)
				for _, steps := range phase.Steps {
					stepsDisplay := fmt.Sprintf("Step %s (%s)", steps.Name, steps.State)
					phaseBranchName.AddBranch(stepsDisplay)
				}
			}
		} else {
			planDisplay := fmt.Sprintf("Plan %s (%s strategy) [NOT ACTIVE]", name, plan.Strategy)
			planBranchName := rootBranchName.AddBranch(planDisplay)
			for _, phase := range plan.Phases {
				phaseDisplay := fmt.Sprintf("Phase %s (%s strategy) [NOT ACTIVE]", phase.Name, phase.Strategy)
				phaseBranchName := planBranchName.AddBranch(phaseDisplay)
				for _, steps := range plan.Phases {
					stepsDisplay := fmt.Sprintf("Step %s (%s strategy) [NOT ACTIVE]", steps.Name, steps.Strategy)
					stepBranchName := phaseBranchName.AddBranch(stepsDisplay)
					for _, step := range steps.Steps {
						stepDisplay := fmt.Sprintf("%s [NOT ACTIVE]", step.Name)
						stepBranchName.AddBranch(stepDisplay)
					}
				}
			}
		}
	}

	fmt.Printf("Plan(s) for \"%s\" in namespace \"%s\":\n", instance.Name, namespace)
	fmt.Println(tree.String())

	return nil
}
