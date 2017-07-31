# Deploy The Envoy DaemonSet

The [Envoy](https://lyft.github.io/envoy/docs/intro/what_is_envoy.html) will be deployed to each cluster node using a [DaemonSet](https://kubernetes.io/docs/concepts/workloads/controllers/daemonset)

## Deploy Envoy to each node

```
kubectl create configmap envoy \
  --namespace kube-system \
  --from-file envoy.json 
```

Envoy will be deployed to each host using a DaemonSet. Envoy will be configured to listen on the host network of each node. Envoy will also be configured to use the Kubernetes cluster DNS service which will make it easy to locate the `kubernetes-envoy-sds` service. 

```
kubectl apply -f daemonsets/envoy.yaml
```

> To ensure Envoy has access to the cluster DNS service, while running in the host network namespace, the `dnsPolicy` of Envoy DaemonSet is set to ClusterFirstWithHostNet.
