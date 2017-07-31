# Deploy The Kubernetes Envoy Service Discovery Service

The `kubernetes-envoy-sds` service runs in the `kube-system` namespace serving the Envoy service discovery REST API to Envoy servers.

## Deploy the kubernetes-envoy-sds service

```
kubectl apply -f deployments/kubernetes-envoy-sds.yaml
```

```
kubectl apply -f services/kubernetes-envoy-sds.yaml
```

At this point the `kubernetes-envoy-sds` service is available in the `kube-system` namespace.
