package serverless

import (
	"context"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

func DeleteIstioNativeSidecar(ctx context.Context, m manager.Manager) error {
	m.GetLogger().Info("Deleting Istio native sidecar annotations from Functions")
	annotation := "sidecar.istio.io/nativeSidecar"

	deployments, err := listAnnotatedDeployments(ctx, m.GetAPIReader(), annotation)
	if err != nil {
		return fmt.Errorf("failed to list annotated deployments: %w", err)
	}

	var collectedErrors []string

	// delete the annotation from each deployment
	for i := range deployments {
		deployment := &deployments[i]
		base := deployment.DeepCopy()

		// Remove annotation from Deployment pod template
		delete(deployment.Spec.Template.ObjectMeta.Annotations, annotation)
		if err := m.GetClient().Patch(ctx, deployment, client.MergeFrom(base)); err != nil {
			collectedErrors = append(collectedErrors, fmt.Sprintf("failed to delete annotation from deployment %s/%s: %s", deployment.Namespace, deployment.Name, err))
		}
	}

	if len(collectedErrors) > 0 {
		return fmt.Errorf("errors occurred while deleting Istio native sidecar annotations: %v", collectedErrors)
	}

	m.GetLogger().Info("Cleanup finished", "deploymentsProcessed", len(deployments))
	return nil
}

func listAnnotatedDeployments(ctx context.Context, m client.Reader, annotation string) ([]appsv1.Deployment, error) {
	labelSelector := client.MatchingLabels{
		"serverless.kyma-project.io/managed-by": "function-controller",
	}

	var allDeployments appsv1.DeploymentList
	if err := m.List(ctx, &allDeployments, labelSelector); err != nil {
		return nil, err
	}

	var filtered []appsv1.Deployment
	for _, dep := range allDeployments.Items {
		if dep.Annotations != nil {
			if _, exists := dep.Annotations[annotation]; exists {
				filtered = append(filtered, dep)
				continue
			}
		}
		if dep.Spec.Template.Annotations != nil {
			if _, exists := dep.Spec.Template.Annotations[annotation]; exists {
				filtered = append(filtered, dep)
			}
		}
	}

	return filtered, nil
}
