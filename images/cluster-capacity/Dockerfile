FROM registry.svc.ci.openshift.org/openshift/release:golang-1.13 AS builder
WORKDIR /go/src/sigs.k8s.io/cluster-capacity
COPY . .
RUN GO111MODULE=auto go build -o hypercc ./cmd/hypercc

FROM registry.svc.ci.openshift.org/openshift/origin-v4.0:base
COPY --from=builder /go/src/sigs.k8s.io/cluster-capacity/hypercc /usr/bin/
RUN ln -sf /usr/bin/hypercc /usr/bin/cluster-capacity
RUN ln -sf /usr/bin/hypercc /usr/bin/genpod
CMD ["/usr/bin/cluster-capacity", "--help"]
