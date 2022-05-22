package resolver

import "github.com/durudex/durudex-gateway/internal/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{ service *service.Service }

func NewResolver(service *service.Service) *Resolver {
	return &Resolver{service: service}
}
