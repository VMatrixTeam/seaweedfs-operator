package controllers

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	seaweedv1 "github.com/seaweedfs/seaweedfs-operator/api/v1"
)

func (r *SeaweedReconciler) createVolumeServerService(m *seaweedv1.Seaweed) *corev1.Service {
	labels := labelsForVolumeServer(m.Name)

	dep := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      m.Name + "-volume",
			Namespace: m.Namespace,
			Labels:    labels,
			Annotations: map[string]string{
				"service.alpha.kubernetes.io/tolerate-unready-endpoints": "true",
			},
		},
		Spec: corev1.ServiceSpec{
			ClusterIP:                "None",
			PublishNotReadyAddresses: true,
			Ports: []corev1.ServicePort{
				{
					Name:       "swfs-volume",
					Protocol:   corev1.Protocol("TCP"),
					Port:       8444,
					TargetPort: intstr.FromInt(8444),
				},
				{
					Name:       "swfs-volume-grpc",
					Protocol:   corev1.Protocol("TCP"),
					Port:       18444,
					TargetPort: intstr.FromInt(18444),
				},
			},
			Selector: labels,
		},
	}
	return dep
}
