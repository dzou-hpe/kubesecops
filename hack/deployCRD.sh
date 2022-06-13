#!/usr/bin/env bash

set -e

if [[ -z $1 ]]; then
  echo "Warning: Use default kubeconfig"
else
  export KUBECONFIG=$1
fi

kubectl delete -f examples/zap/crd.yaml || true
#kubectl delete -f https://gist.githubusercontent.com/sdenel/1bd2c8b5975393ababbcff9b57784e82/raw/f1b885349ba17cb2a81ca3899acc86c6ad150eb1/nginx-hello-world-deployment.yaml || true

kubectl apply -f https://gist.githubusercontent.com/sdenel/1bd2c8b5975393ababbcff9b57784e82/raw/f1b885349ba17cb2a81ca3899acc86c6ad150eb1/nginx-hello-world-deployment.yaml
kubectl apply -f examples/zap/crd.yaml

# create a custom resource of type Zap
kubectl apply -f examples/zap/example.yaml

sleep 5

# check deployments created through the custom resource
kubectl get deployments