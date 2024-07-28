package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"

	"github.com/gabelchinmay/kubernetes-go-client/handler"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
)

const (
	namespace  = "default"
	secretName = "sample-secret"
)

// Sample code to demonstrate how to use the handler package

func main() {
	handler, err := handler.NewHandler(context.TODO())
	if err != nil {
		log.Fatalf("failed to create handler: %v", err)
	}

	corev1.AddToScheme(handler.Scheme)
	retrievedSecret := &corev1.Secret{}

	if err := handler.Get(types.NamespacedName{Name: secretName, Namespace: namespace}, retrievedSecret); err != nil {
		log.Fatalf("failed to get secret: %v", err)
	}

	usernameBase64 := retrievedSecret.Data["username"]
	passwordBase64 := retrievedSecret.Data["password"]

	decodedUsername, err := base64.StdEncoding.DecodeString(string(usernameBase64))
	if err != nil {
		log.Fatalf("failed to decode username: %v", err)
	}

	decodedPassword, err := base64.StdEncoding.DecodeString(string(passwordBase64))
	if err != nil {
		log.Fatalf("failed to decode password: %v", err)
	}

	fmt.Printf("Username: %s\n", string(decodedUsername))
	fmt.Printf("Password: %s\n", string(decodedPassword))
}
