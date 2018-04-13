#
# This is the cluster capacity tool.
#
# The standard name for this image is openshift/origin-cluster-capacity
#
FROM openshift/origin-base

RUN INSTALL_PKGS="openshift-cluster-capacity" && \
    yum --enablerepo=origin-local-release install -y ${INSTALL_PKGS} && \
    rpm -V ${INSTALL_PKGS} && \
    yum clean all

RUN ln -sf /usr/bin/hypercc /usr/bin/cluster-capacity
RUN ln -sf /usr/bin/hypercc /usr/bin/genpod
CMD ["/usr/bin/cluster-capacity --help"]
