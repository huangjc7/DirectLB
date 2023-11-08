package main

import (
	"github.com/huangjc7/directLB/controller"
	"github.com/huangjc7/directLB/pkg/generated/informers/externalversions"
	"k8s.io/client-go/informers"
	"time"
)

func main() {
	DtlbClientset, ClientSet := controller.InitClient()
	stopCh := controller.SetupSignalHandler()
	// 原生deployment informer
	deploySharedInformerFactory := informers.NewSharedInformerFactory(ClientSet, time.Second*5)
	// dtlb informer
	dtlbSharedInformerFactory := externalversions.NewSharedInformerFactory(DtlbClientset, time.Second*5)
	dtlbController := controller.NewDtlbController(ClientSet,
		dtlbSharedInformerFactory.Dtlb().V1beta1().DirectLBs(),
		deploySharedInformerFactory.Apps().V1().Deployments())

	go dtlbSharedInformerFactory.Start(stopCh)
	go deploySharedInformerFactory.Start(stopCh)

	dtlbController.Run(5, stopCh)

}
