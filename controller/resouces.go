package controller

import (
	"github.com/huangjc7/directLB/pkg/apis/dtlb.io/v1beta1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func MutateDeloyment(lb *v1beta1.DirectLB) *appsv1.Deployment {
	// 设置Deployment的详细配置
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: lb.Name,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: lb.Spec.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "example",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "example",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "web",
							Image: lb.Spec.Image,
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									Protocol:      corev1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}
	return deployment
}
