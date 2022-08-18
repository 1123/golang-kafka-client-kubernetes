POD=$(kubectl get pods --no-headers -o custom-columns=":metadata.name" -n go-sample-consumer)
kubectl top pod $POD -n go-sample-consumer
