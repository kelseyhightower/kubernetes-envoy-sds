# Deploy the Consumer Service

The consumer service is designed to test the [Envoy DaemonSet](deploy-envoy-daemonset.md) and [Kubernetes Envoy Service Discovery Service](deploy-kubernetes-envoy-sds.md) by making HTTP request through a local Envoy proxy.

## Deploy the Consumer Service

```
kubectl apply -f deployments/consumer.yaml
```

Review the consumer service logs

```
kubectl logs -f \
  $(kubectl get pods \
      -l app=consumer \
      -o jsonpath='{.items[0].metadata.name}')
```

The consumer is able to communicate with the local Envoy proxy by referencing the host IP address at runtime as shown below:

```
spec:
  containers:
	- name: consumer
	  image: gcr.io/hightowerlabs/consumer:0.0.1
	  imagePullPolicy: Always
	  env:
		- name: HOST_IP
		  valueFrom:
			fieldRef:
			  fieldPath: status.hostIP
	  args:
		- "-proxy=http://$(HOST_IP):80"
```
