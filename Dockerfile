FROM scratch
ADD kubernetes-sds /kubernetes-sds
ENTRYPOINT ["/kubernetes-sds"]
