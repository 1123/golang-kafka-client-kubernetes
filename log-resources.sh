while true; do
  POD=$(kubectl get pods --no-headers -o custom-columns=":metadata.name" -n go-sample-consumer)
  date >> resources.log
  kubectl top pod $POD -n go-sample-consumer >> resources.log
  sleep 60
done
