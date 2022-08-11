FROM alpine:latest
WORKDIR deploy
RUN apk add --no-cache libc6-compat
RUN apk add --no-cache gcompat
COPY GoConsumerImage /deploy
CMD ["/deploy/GoConsumerImage"]
