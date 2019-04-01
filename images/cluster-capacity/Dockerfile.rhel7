FROM registry.svc.ci.openshift.org/ocp/builder:golang-1.10 AS builder
WORKDIR /go/src/github.com/kubernetes-incubator/cluster-capacity
COPY . .
RUN make build; \
    mkdir -p /tmp/build; \
    cp /go/src/github.com/kubernetes-incubator/cluster-capacity/_output/local/bin/linux/$(go env GOARCH)/hypercc /tmp/build/hypercc

FROM registry.svc.ci.openshift.org/ocp/4.0:base
COPY --from=builder /tmp/build/hypercc /usr/bin/
RUN ln -sf /usr/bin/hypercc /usr/bin/cluster-capacity
RUN ln -sf /usr/bin/hypercc /usr/bin/genpod
CMD ["/usr/bin/cluster-capacity", "--help"]
