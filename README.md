### Vendoring
```
gvt fetch --branch master github.com/TykTechnologies/tyk-protobuf 
gvt fetch --revision v0.11.2 github.com/sirupsen/logrus
```

### Building and Running Docker Image
```
# required for alpine image
export GOOS=linux
export GOARCH=amd64

# build the binary
go build -i

# create a docker image from the built binary
docker build . -t tyk-auth-golang-grpc-plugin:1.0

# run the docker image
docker run -d tyk-auth-golang-grpc-plugin:1.0 -d
```