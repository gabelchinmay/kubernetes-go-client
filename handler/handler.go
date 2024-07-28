package handler

import (
	"context"
	"fmt"
	"os"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

type Handler struct {
	Ctx    context.Context
	Client client.Client
	Scheme *runtime.Scheme
}

func NewHandler(ctx context.Context) (*Handler, error) {
	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig == "" {
		return nil, fmt.Errorf("KUBECONFIG environment variable not set")
	}

	cfg, err := config.GetConfig()
	if err != nil {
		return nil, fmt.Errorf("error getting kubeconfig: %v", err)
	}

	s := runtime.NewScheme()

	c, err := client.New(cfg, client.Options{Scheme: s})
	if err != nil {
		return nil, fmt.Errorf("error creating Kubernetes client: %v", err)
	}

	return &Handler{Ctx: ctx, Client: c, Scheme: s}, nil
}

func (h *Handler) Create(obj client.Object) error {
	return h.Client.Create(h.Ctx, obj)
}

func (h *Handler) Get(key types.NamespacedName, obj client.Object) error {
	return h.Client.Get(h.Ctx, key, obj)
}
