Exercise on understanding how probes affecting Pod's lifecycle.
Useful links:
- [Liveness, Readiness, and Startup Probes](https://kubernetes.io/docs/concepts/configuration/liveness-readiness-startup-probes/)
- [Configure Liveness, Readiness and Startup Probes](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/)

The example shows a liveness-driven restart loop.

1. Apply file `kubectl apply -f pod.yaml`
2. Check the state of the Pod `kubectl -n ckad-probes describe pod probes`
3. Get the restart count of Pod `kubectl -n ckad-probes get pod probes`
4. Constantly track events `kubectl get events -n ckad-probes --watch`