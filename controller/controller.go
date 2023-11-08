package controller

import (
	"context"
	"fmt"
	"github.com/huangjc7/directLB/pkg/apis/dtlb.io/v1beta1"
	informer "github.com/huangjc7/directLB/pkg/generated/informers/externalversions/dtlb.io/v1beta1"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	appsv1 "k8s.io/client-go/informers/apps/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"
	"time"
)

type DtlbController struct {
	clientset          *kubernetes.Clientset
	workqueue          workqueue.RateLimitingInterface
	dtlbInformer       informer.DirectLBInformer
	deploymentInformer appsv1.DeploymentInformer
}

type controller interface {
	Run(threadiness int, stopCh <-chan struct{})
}

func NewDtlbController(clientset *kubernetes.Clientset,
	dtlbInformer informer.DirectLBInformer,
	deploymentInformer appsv1.DeploymentInformer) controller {
	d := &DtlbController{
		clientset:          clientset,
		workqueue:          workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter()),
		dtlbInformer:       dtlbInformer,
		deploymentInformer: deploymentInformer,
	}

	klog.Info("settings controller")
	//注册dtlb informer事件
	dtlbInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    d.onAdd,
		UpdateFunc: d.onUpdate,
		DeleteFunc: d.onDelete,
	})
	//注册deployment informer事件
	deploymentInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    d.onAdd,
		UpdateFunc: d.onUpdate,
		DeleteFunc: d.onDelete,
	})

	return d
}

func (d *DtlbController) onAdd(obj interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		runtime.HandleError(err)
		return
	}
	d.workqueue.AddRateLimited(key)
}
func (d *DtlbController) onUpdate(oldObj, newObj interface{}) {
	if oldDirectLB, ok := oldObj.(*v1beta1.DirectLB); ok {
		if newDirectLB, ok := newObj.(*v1beta1.DirectLB); ok {
			if oldDirectLB.ResourceVersion == newDirectLB.ResourceVersion {
				klog.Info("DirectLB resource version is the same. No update needed.")
				return
			}
			d.onAdd(newDirectLB)
		}
	}

	if oldDeploy, ok := oldObj.(*v1.Deployment); ok {
		if newDeploy, ok := newObj.(*v1.Deployment); ok {
			if oldDeploy.ResourceVersion == newDeploy.ResourceVersion {
				klog.Info("DirectLB resource version is the same. No update needed.")
				return
			}
			d.onAdd(newDeploy)
		}
	}
}
func (d *DtlbController) onDelete(obj interface{}) {
	klog.Info("Delete event occurs")
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
	if err != nil {
		runtime.HandleError(err)
		return
	}
	d.workqueue.AddRateLimited(key)
}

// 等待缓存同步以后正式启动循环控制
func (d *DtlbController) Run(threadiness int, stopCh <-chan struct{}) {
	defer runtime.HandleCrash()

	//控制器停止后关闭队列
	defer d.workqueue.ShutDown()
	//实际启动informer factory
	klog.Info("Starting DirectLB Controller")
	//等待缓存同步以后才处理队列中的数据
	if !cache.WaitForCacheSync(stopCh, d.dtlbInformer.Informer().HasSynced, d.deploymentInformer.Informer().HasSynced) {
		runtime.HandleError(fmt.Errorf("timed out waiting caches to sync"))
		return
	}

	klog.Info("Informer caches to sync completed")

	for i := 0; i < threadiness; i++ {
		go wait.Until(d.runWork, time.Second, stopCh)
	}
	<-stopCh
	klog.Info("Stopping DirectLB controller")
}

func (d *DtlbController) runWork() {
	for d.processNextItem() {
	}
}

func (d *DtlbController) processNextItem() bool {
	obj, shutdown := d.workqueue.Get()
	if shutdown {
		return false
	}

	err := func(obj interface{}) error {
		//告诉队列我们已经完成此key的处理
		//这将为其他work解锁该key
		//这将确保并行处理，因为永远不会并行处理相同key的两个pod
		defer d.workqueue.Done(obj)

		//将obj断言为string
		var ok bool
		var key string
		if key, ok = obj.(string); !ok {
			d.workqueue.Forget(obj)
			return fmt.Errorf("expected string in workqueue but get %#v", obj)
		}

		//具体的处理业务逻辑
		if err := d.syncHandler(key); err != nil {
			return fmt.Errorf("sync error: %v", err)
		}
		d.workqueue.Forget(obj)
		klog.Infof("successfully synced %s", obj)
		return nil
	}(obj)
	if err != nil {
		runtime.HandleError(err)
	}
	return true
}

// key -> directlb ->indexer
func (d *DtlbController) syncHandler(key string) error {
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}

	//获取directlb 从index中获取的
	directlb, err := d.dtlbInformer.Lister().DirectLBs(namespace).Get(name)
	if err != nil {
		if errors.IsNotFound(err) {
			klog.Warningf("DirectLB deleting:%s/%s...", namespace, name)
			return nil
		}
		return err
	}
	//klog.Infof("DirectLB try to process %#v ...", directlb)
	//如果没有则创建
	deployment, err := d.deploymentInformer.Lister().Deployments(namespace).Get(name)
	if err != nil {
		if errors.IsNotFound(err) {
			d.createDeploymentResouce(namespace, directlb)
			return nil
		}
		return err
	}

	//如果dtlb的yaml文件发生改变则修改deployment
	if *deployment.Spec.Replicas == *directlb.Spec.Replicas {
		return nil
	} else {
		klog.Info("资源发生更改")
		d.updateDeploymentResouce(namespace, directlb)
	}
	return nil
}

func (d *DtlbController) createDeploymentResouce(namespace string, lb *v1beta1.DirectLB) {

	deploymentClient := d.clientset.AppsV1().Deployments(namespace)
	result, err := deploymentClient.Create(context.TODO(), MutateDeloyment(lb), metav1.CreateOptions{})
	if err != nil {
		klog.Errorf("Failed to create Deployment: %v", err)
	}
	klog.Infof("Created deployment %q.\n", result.GetObjectMeta().GetName())
}

func (d *DtlbController) updateDeploymentResouce(namespace string, lb *v1beta1.DirectLB) {

	deploymentClient := d.clientset.AppsV1().Deployments(namespace)
	result, err := deploymentClient.Update(context.TODO(), MutateDeloyment(lb), metav1.UpdateOptions{})
	if err != nil {
		klog.Errorf("Failed to create Deployment: %v", err)
	}
	klog.Infof("Update deployment %q.\n", result.GetObjectMeta().GetName())
}
