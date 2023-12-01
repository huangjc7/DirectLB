# 不使用controller-runtime框架
# 基于informer原生实现kube-operator能力
* 核心逻辑总结：通过分别控制不同的informer(deployment/directlb)来进行ListAndWatch 持续获取事件
* 通过将directlb字段信息传递至deployment字段内通过clientset进行创建及更新操作
