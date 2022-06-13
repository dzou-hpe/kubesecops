#!/usr/bin/env bash

set -e

if [[ -z $1 ]]; then
  echo "Warning: Use default kubeconfig"
else
  export KUBECONFIG=$1
fi

kubectl delete -f examples/zap/crd.yaml

kubectl apply -f examples/zap/crd.yaml

# create a custom resource of type Zap
kubectl apply -f examples/zap/example.yaml

sleep 5

# check deployments created through the custom resource
kubectl get deployments