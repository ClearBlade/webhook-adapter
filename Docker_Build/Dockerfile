#docker build -f ../docker_build/Dockerfile -t amqpAdapter --target runtime  ../..

FROM alpine:latest
LABEL description="Runtime container - webhook-adapter"

#Copy the GooglePubSubAdapter binary into the image
COPY ./webhook-adapter /usr/local/bin/webhook-adapter

RUN apk --no-cache --update add ca-certificates

ENTRYPOINT ["/usr/local/bin/webhook-adapter"]