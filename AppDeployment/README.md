Preparing for Application Deployment (20%) section.

blue/green deployment scenario using only Deployment and Service k8s resources.

1. Apply the whole file: `kubectl apply -n ckad -f blue-green.yaml`
2. Check to which Pods both Services are pointing:
   1. `kubectl -n ckad get endpoints webapp-svc`
   2. `kubectl -n ckad get endpoints webapp-preview`
3. Switch label of `webapp-svc` selector to point to green, apply: `kubectl apply -n ckad -f blue-green.yaml`
4. Check to which Pods the `webapp-svc` is now pointing: `kubectl -n ckad get endpoints webapp-svc`
5. Switch label of `webapp-svc` selector back to `blue`, apply: `kubectl apply -n ckad -f blue-green.yaml`
6. Check to which Pods the `webapp-svc` is now pointing: `kubectl -n ckad get endpoints webapp-svc`