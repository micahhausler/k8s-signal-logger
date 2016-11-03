# k8s-signal-logger

k8s-signal-logger is a simple program for observing how Kubernetes responds to
[liveness and readiness probes](http://kubernetes.io/docs/api-reference/v1/definitions/#_v1_container)

There are 3 endpoints to the container:

- `/healthz` - Used for the health check
- `/ready` - Used for the readiness check
- `/` - A homepage handler - always responds with `200`

If a `SIGTERM` is sent by Kubernetes, the signal is logged to STDOUT and ignored. This way, you can see
if/how much traffic is still being sent to the pod once it is supposed to be in a 'shutdown' mode.

## Try it out

The following commands run:

- Create the deployment and service
- Get the service's nodePort and a node's IP
- Run Apache bench against it
- Update the deployment. Once the deployment is completed, you can see the results from apache bench

```
kubectl create -f deployment.yaml -f service.yaml
export PORT=$(kubectl get -f service.yaml -o json | jq -r .spec.ports[0].nodePort)
export NODE=$(kubectl get no -o json | jq -r .items[0].status.addresses[0].address)
ab -n 1000000 -c 20 http://$NODE:$PORT/

# In a separate terminal
# - change a label on deployment.yaml
kubectl replace -f deployment.yaml
```

## License
MIT License. See [License](/LICENSE) for full text
