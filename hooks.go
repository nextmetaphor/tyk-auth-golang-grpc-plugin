package main

import (
	"github.com/TykTechnologies/tyk-protobuf/bindings/go"
)

func MyPostHook(object *coprocess.Object) (*coprocess.Object, error) {
	secondaryAuthHeader := object.Request.Headers["Secondary-Authorization"]
	object.Request.SetHeaders = map[string]string{"Authorization": secondaryAuthHeader}
	object.Request.DeleteHeaders = []string{"Secondary-Authorization"}

	return object, nil
}
