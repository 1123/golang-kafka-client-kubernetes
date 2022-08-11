while true; do
  date >> resources.log
  kubectl top pod kafka-go-consumer-5d8d956df6-gbtk6 -n go-sample-consumer >> resources.log
  sleep 300
done
