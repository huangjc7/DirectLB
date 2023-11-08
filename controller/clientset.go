package controller

import (
	"flag"
	"github.com/huangjc7/directLB/pkg/generated/clientset/versioned"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/klog/v2"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

func InitClient() (*versioned.Clientset, *kubernetes.Clientset) {
	var kubeconfig *string
	var config *rest.Config
	var err error
	var dtlbClientSet *versioned.Clientset
	var ClientSet *kubernetes.Clientset

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(可选) kubeconfig 文件的绝对路径")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "kubeconfig 文件的绝对路径")
	}
	flag.Parse()

	if config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig); err != nil {
		klog.Fatal("build config from tags error:", err)
	}
	dtlbClientSet, err = versioned.NewForConfig(config)
	if err != nil {
		klog.Fatal("New for dtlb clientset config err:", err)
	}
	//
	ClientSet, err = kubernetes.NewForConfig(config)
	if err != nil {
		klog.Fatal("New for clientset config err:", err)
	}
	return dtlbClientSet, ClientSet
}

func SetupSignalHandler() (stopCh <-chan struct{}) {
	stop := make(chan struct{})
	c := make(chan os.Signal, 2)
	signal.Notify(c, []os.Signal{syscall.SIGINT, syscall.SIGTERM}...)
	go func() {
		<-c
		close(stop)
		<-c
		os.Exit(1)
	}()
	return stop
}
