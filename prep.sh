#!/bin/sh

# This will just symlink github.com/kubernetes-incubator/cluster-capacity -> github.com/openshift/cluster-capacity

mkdir $GOPATH/src/github.com/kubernetes-incubator
ln -s $GOPATH/src/github.com/kubernetes-incubator/cluster-capacity $GOPATH/src/github.com/openshift/cluster-capacity
