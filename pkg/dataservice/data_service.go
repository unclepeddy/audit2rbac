package main

import (
	"encoding/json"
	"fmt"

	"k8s.io/apiserver/pkg/apis/audit"
)

func createUser(name string, groups []string) audit.UserInfo{
	return audit.UserInfo{
		Username: name,
		Groups: groups,
	}
}

func createObjectRef(resource, namespace, apiVersion string) *audit.ObjectReference {
	return &audit.ObjectReference{
		Resource: resource,
		Namespace: "true",
		APIVersion: apiVersion,
	}
}

func createEvent(user audit.UserInfo, ref *audit.ObjectReference, verb, requestUri string) audit.Event{
	return audit.Event {
		User: user,
		Verb: verb,
		RequestURI: requestUri,
		ObjectRef: ref,
	}
}

func main() {

	user := createUser("alice", []string{})
	verb := "list"
	requestUri  := "/api/v1/namespaces/default/pods" 
	objectRef := createObjectRef("pods", "namespace", "v1")

	events := []audit.Event{}

	event := createEvent(user, objectRef, verb, requestUri)
	events = append(events, event)

	fmt.Println(events)
}