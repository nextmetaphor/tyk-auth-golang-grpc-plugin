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

## Licence ##
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
