package main

import (
  "fmt"
  "time"

  "github.com/golang/glog"
  corev1 "k8s.io/api/core/v1"
  "k8s.io/apimachinery/pkg/api/errors"
  "k8s.io/apimachinery/pkg/util/runtime"
  utilruntime "k8s.io/apimachinery/pkg/util/runtime"
  "k8s.io/apimachinery/pkg/util/wait"
  "k8s.io/client-go/kubernetes"
  "k8s.io/client-go/kubernetes/scheme"
  typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
  "k8s.io/client-go/tools/cache"
  "k8s.io/client-go/tools/record"
  "k8s.io/client-go/util/workqueue"

  crdV1 "datahub.txzing.com/mysql-gr-operator/pkg/apis/txz_datahub/v1"
  clientSet "datahub.txzing.com/mysql-gr-operator/pkg/client/clientset/versioned"
  crdScheme "datahub.txzing.com/mysql-gr-operator/pkg/client/clientset/versioned/scheme"
  informers "datahub.txzing.com/mysql-gr-operator/pkg/client/informers/externalversions/txz_datahub/v1"
  listers "datahub.txzing.com/mysql-gr-operator/pkg/client/listers/txz_datahub/v1"
)

const (
  controllerAgentName = "mysql-gr-operator-controller"
  SuccessSynced = "Synced"
  MessageResourceSynced = "MysqlGROperator Synced successfully"
)

// 自定义控制器的结构
type Controller struct {
  // 用于存放 kubeclient 的集合
  kubeClientSet kubernetes.Interface
  // 用于存放 mysql-gr-crd 的集合
  crdClientSet clientSet.Interface

  // 事件记录器
  recorder record.EventRecorder

  // 使用一个队列, 用于协调 informer 与 控制循环之间的速率
  workqueue workqueue.RateLimitingInterface

  crdSynced cache.InformerSynced
  crdLister listers.MysqlGROperatorLister
}

// 返回自定义控制器
func NewController(
  kubeClientSet kubernetes.Interface,
  crdClientSet clientSet.Interface,
  informer informers.MysqlGROperatorInformer) *Controller {

  // 创建事件广播, 这里主要是为了能够正常接受事件通知
  utilruntime.Must(crdScheme.AddToScheme(scheme.Scheme))
  glog.V(4).Info("Creating event broadcaster")
  eventBroadcaster := record.NewBroadcaster()
  eventBroadcaster.StartLogging(glog.Infof)
  eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{Interface: kubeClientSet.CoreV1().Events("")})
  recorder := eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{Component: controllerAgentName})

  controller := &Controller{
    kubeClientSet: kubeClientSet,
    crdClientSet: crdClientSet,
    recorder: recorder,
    workqueue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "MysqlGROperators"),
    crdLister: informer.Lister(),
    crdSynced: informer.Informer().HasSynced,
  }

  glog.Info("Setting up event handlers")

  // informer 添加事件监听, 以回调不同的处理
  informer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
    AddFunc: controller.enqueueMysqlGROperator,
    UpdateFunc: func(old, new interface{}) {
      oldOperator := old.(*crdV1.MysqlGROperator)
      newOperator := new.(*crdV1.MysqlGROperator)
      if oldOperator.ResourceVersion == newOperator.ResourceVersion {
        // 说明对象并没有变化
        return
      }
      controller.enqueueMysqlGROperator(new)
    },
    DeleteFunc: controller.enqueueMysqlGROperatorForDelete,
  })

  return controller
}

// 当一个MysqlGROperator 新增事件发生时, 包装成对应的Key, 放入队列之中
func (c *Controller) enqueueMysqlGROperator(obj interface{}) {
  var key string
  var err error
  if key, err = cache.MetaNamespaceKeyFunc(obj); err != nil {
    runtime.HandleError(err)
    return
  }
  c.workqueue.AddRateLimited(key)
}

// 当一个MysqlGROperator 删除事件发生时, 包装成对应的Key, 放入队列之中
func (c *Controller) enqueueMysqlGROperatorForDelete(obj interface{}) {
  var key string
  var err error
  key, err = cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
  if err != nil {
    runtime.HandleError(err)
    return
  }
  c.workqueue.AddRateLimited(key)
}

func (c *Controller) Run(threadiness int, stopCh <-chan struct {}) error {
  defer runtime.HandleCrash()
  defer c.workqueue.ShutDown()

  glog.Info("Starting control loop")

  glog.Info("Waiting for informer caches to sync")
  if ok := cache.WaitForCacheSync(stopCh, c.crdSynced); !ok {
    return fmt.Errorf("failed to wait for caches to sync")
  }

  glog.Info("Starting workers")
  // Launch two workers to process resources
  for i := 0; i < threadiness; i++ {
    go wait.Until(c.runWorker, time.Second, stopCh)
  }

  glog.Info("Started workers")
  <-stopCh
  glog.Info("Shutting down workers")

  return nil
}

func (c *Controller) runWorker() {
  for c.processNextWorkItem() {
  }
}


func (c *Controller) processNextWorkItem() bool {
  obj, shutdown := c.workqueue.Get()  // 从事件队列中取出事件进行处理
  if shutdown {
    return false
  }

  err := func(obj interface{}) error {
    defer c.workqueue.Done(obj)
    var key string
    var ok bool

    if key, ok = obj.(string); !ok {
      c.workqueue.Forget(obj)
      runtime.HandleError(fmt.Errorf("expected string in workqueue but got %#v", obj))
      return nil
    }

    if err := c.syncHandler(key); err != nil {
      return fmt.Errorf("error syncing '%s': %s", key, err.Error())
    }

    c.workqueue.Forget(obj)
    glog.Infof("Successfully synced '%s'", key)
    return nil
  }(obj)

  if err != nil {
    runtime.HandleError(err)
    return true
  }
  return true
}

func (c *Controller) syncHandler(key string) error {
  namespace, name, err := cache.SplitMetaNamespaceKey(key)
  if err != nil {
    runtime.HandleError(fmt.Errorf("invalid resource key: %s", key))
    return nil
  }

  crd, err := c.crdLister.MysqlGROperators(namespace).Get(name)
  if err != nil {
    // 如果这里找不到, 说明对象已经在local cache中被删除, 即kubectl执行了一个delete操作删除掉了crd对象
    if errors.IsNotFound(err) {

      glog.Warningf("MysqlGROperator: %s/%s does not exist in local cache, will delete it from Neutron ...", namespace, name)

      glog.Infof("[Neutron] Deleting MysqlGROperator: %s/%s ...", namespace, name)

      // 这里写删除crd之后的逻辑
      return nil
    }

    runtime.HandleError(fmt.Errorf("failed to list crd by: %s/%s", namespace, name))
    return err
  }


  glog.Infof("[Neutron] Try to process MysqlGROperator: %#v ...", crd)

  // 编排逻辑


  c.recorder.Event(crd, corev1.EventTypeNormal, SuccessSynced, MessageResourceSynced)
  return nil
}
