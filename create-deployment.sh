kubectl create deployment \
	-n go-sample-consumer kafka-go-consumer \
	--image=gcr.io/solutionsarchitect-01/golang-kafka-client:0.1 \
	--replicas=1
