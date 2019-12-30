package v1

import (
  metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 为下面的类型生成对应的Client代码
// +genclient
// 下面类型的无需Status
// +genclient:noStatus
// 在生成 DeepCopy 的时候，实现 Kubernetes 提供的 runtime.Object 接口。
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type MysqlGROperator struct {
  // TypeMeta is the metadata for the resource, like kind and apiversion
  metav1.TypeMeta `json:",inline"`
  // ObjectMeta contains the metadata for the particular object, including
  // things like...
  //  - name
  //  - namespace
  //  - self link
  //  - labels
  //  - ... etc ...
  metav1.ObjectMeta `json:"metadata,omitempty"`

  // Spec is the custom resource spec
  Spec MysqlGROperatorSpec `json:"spec"`
}

type MysqlGROperatorSpec struct {
  Size int `json:"size"`  // 集群大小
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type MysqlGROperatorList struct {
  metav1.TypeMeta `json:",inline"`
  metav1.ListMeta `json:"metadata"`

  Items []MysqlGROperator `json:"items"`
}
