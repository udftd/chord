package eviction

import (
	"context"

	"github.com/udftd/chord/pkg/pods"

	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/errors"
	"k8s.io/client-go/kubernetes"
)

// GracefullyDrainNode drains the node by evicting all the pods gracefully
func GracefullyDrainNode(ctx context.Context, client kubernetes.Interface, node corev1.Node) error {
	// list all pods on the node
	nodePods, err := pods.ListNodePods(ctx, client, node.Name)
	if err != nil {
		if k8serrors.IsNotFound(err) {
			return nil
		}
		return err
	}
	// get pods that are managed by a controller
	for _, pod := range nodePods.Items {
		if pod.DeletionTimestamp != nil {
			continue
		}

	}
	return nil
}
