FROM scratch
ADD kubernetes-envoy-sds /kubernetes-envoy-sds
ENTRYPOINT ["/kubernetes-envoy-sds"]
