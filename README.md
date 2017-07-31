# Kubernetes Envoy Service Discovery Service

The `kubernetes-envoy-sds` service implements the [Envoy Service Discovery REST API](https://lyft.github.io/envoy/docs/configuration/cluster_manager/sds_api.html)


## Deploy

```
kubectl apply -f deployments/kubernetes-envoy-sds.yaml 
```

```
kubectl apply -f services/kubernetes-envoy-sds.yaml
```
