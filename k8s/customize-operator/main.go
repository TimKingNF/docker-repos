package main

import (
  "flag"
  "time"

  "github.com/golang/glog"
  "k8s.io/client-go/kubernetes"
  "k8s.io/client-go/tools/clientcmd"

  clientset "datahub.txzing.com/mysql-gr-operator/pkg/client/clientset/versioned"
  informers "datahub.txzing.com/mysql-gr-operator/pkg/client/informers/externalversions"
  "datahub.txzing.com/mysql-gr-operator/pkg/signals"
)

var (
  masterURL  string
  kubeconfig string
)

func init() {
  flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
  flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
}

func main() {
  flag.Parse()

  // 获取信号量
  stopCh := signals.SetupSignalHandler()

  // 获取config, 如果为空, 默认使用集群内的配置
  cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
  if err != nil {
    glog.Fatalf("Error building kubeconfig: %s", err.Error())
  }

  // 创建 kubernetes client
  kubeClient, err := kubernetes.NewForConfig(cfg)
  if err != nil {
    glog.Fatalf("Error building Kubernetes clientset: %s", err.Error())
  }

  // 创建 mysql-gr-operator client
  mysqlGROperatorClient, err := clientset.NewForConfig(cfg)
  if err != nil {
    glog.Fatalf("Error building mysql-gr-operator clientset: %s", err.Error())
  }

  // 创建工厂类,并生成 informer 对象
  mysqlGROperatorInformerFactory := informers.NewSharedInformerFactory(mysqlGROperatorClient, time.Second*30)
  mysqlGROperatorInformer := mysqlGROperatorInformerFactory.Datahub().V1().MysqlGROperators()

  // 传递给控制器
  controller := NewController(kubeClient, mysqlGROperatorClient, mysqlGROperatorInformer)

  go mysqlGROperatorInformerFactory.Start(stopCh) // 启动informer

  // 启动自定义控制器
  if err = controller.Run(2, stopCh); err != nil {
    glog.Fatalf("Error running controller: %s", err.Error())
  }
}
