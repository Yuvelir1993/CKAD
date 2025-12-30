Preparing for Application Deployment (20%) section.

## Blue/Green
Blue-Green deployment model is having 2 versions of your application in a separate environments. Blue - is a production one, Green - testing/staging. When Green is fully tested, traffic is switched to Green, making it a new Blue, then old Blue becomes a new Green.

### Deployment + Service
blue/green deployment scenario using only Deployment and Service k8s resources.

1. Apply the whole file: `kubectl apply -n ckad -f blue-green.yaml`
2. Check to which Pods both Services are pointing:
   1. `kubectl -n ckad get endpoints webapp-svc`
   2. `kubectl -n ckad get endpoints webapp-preview`
3. Switch label of `webapp-svc` selector to point to green, apply: `kubectl apply -n ckad -f blue-green.yaml`
4. Check to which Pods the `webapp-svc` is now pointing: `kubectl -n ckad get endpoints webapp-svc`
5. Switch label of `webapp-svc` selector back to `blue`, apply: `kubectl apply -n ckad -f blue-green.yaml`
6. Check to which Pods the `webapp-svc` is now pointing: `kubectl -n ckad get endpoints webapp-svc`
   
### Ingress
Deploy the `blue-green.yaml` in it's initial state and create `Ingress` resource to point to one of the Services.
To enable ingress controller on minikube, type `minikube addons enable ingress`.

Then, create Ingress resource:
`k create ingress bluegreen --rule="webapp.ckad.local/*=webapp-svc:80" --rule="preview.ckad.local/*=webapp-preview:80" -n ckad --class=nginx`

## Canary
Canary deployment model is when part of a traffic is moved to a staging version of app or microservice. If there is no failures and desired effect is achieved, 100% of traffic can be switched to a new version.

### Deployment + Service
Deploy resources with `kubectl apply -f canary-simple.yaml`.
Check to which Pods Service is pointing `kubectl -n ckad get endpoints svc-app`.
Increase traffic to canary by changing `replicas` number to higher in `canary-deployment` and down in `stable-deployment`.

### Ingress
Practicing [ingress nginx with weighting traffic](https://kubernetes.github.io/ingress-nginx/examples/canary/).

Apply `canary-ingress.yaml`
Create Ingress resources. First, create Ingress for a stable traffic:
```bash
kubectl -n ckad-canary-ingress create ingress canary-stable \
  --class=nginx \
  --rule="webapp.ckad.local/=svc-stable:80"
```

After, create Ingress for canary traffic and annotate with needed weight:
```bash
kubectl -n ckad-canary-ingress create ingress canary-canary \
  --class=nginx \
  --rule="webapp.ckad.local/=svc-canary:80"

kubectl -n ckad-canary-ingress annotate ingress canary-canary \
  nginx.ingress.kubernetes.io/canary="true" \
  nginx.ingress.kubernetes.io/canary-weight="10"
```

Check ingress correctness:
`kubectl -n ckad-canary-ingress describe ingress canary-stable`
`kubectl -n ckad-canary-ingress describe ingress canary-canary`