# Kubernetes Envoy Service Discovery Service

The `kubernetes-envoy-sds` service implements the [Envoy Service Discovery REST API](https://lyft.github.io/envoy/docs/configuration/cluster_manager/sds_api.html) on top of the [Kubernetes Services API](https://kubernetes.io/docs/concepts/services-networking/service).  

Each Kubernetes service can be referenced in an Envoy config by its FQDN. The following FQDN maps to the `nginx` service running in the `default` namespace:

```
nginx.default.svc.cluster.local
```

See the [envoy.json](envoy.json) configuration file for a complete example.

## Usage

```
kubernetes-envoy-sds -h
```

```
Usage of kubernetes-envoy-sds:
  -cluster-domain string
    	The cluster domain (default "svc.cluster.local")
  -http string
    	The HTTP listen address. (default "127.0.0.1:8080")
```

## Usage Tutorial

This tutorial will walk you through deploying the `kubernetes-envoy-sds` service and an Envoy service mesh across each node in a Kubernetes cluster. Once the Envoy infrastructure is in place you'll have a chance to test it using the `consumer` example service.

Kubernetes 1.6+ is required.

* [Deploy The Kubernetes Envoy Service Discovery Service](docs/deploy-kubernetes-envoy-sds.md)
* [Deploy The Envoy DaemonSet](docs/deploy-envoy-daemonset.md)
* [Deploy The Consumer Service](docs/deploy-consumer-service.md)
* [Cleanup](docs/cleanup.md)
