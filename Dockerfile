FROM alpine:latest
ADD tyk-auth-golang-grpc-plugin /tmp/tyk-auth-golang-grpc-plugin
ENTRYPOINT /tmp/tyk-auth-golang-grpc-plugin