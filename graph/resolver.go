//go:generate go run github.com/99designs/gqlgen
package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	MarketSocket map[int][]chan string
}

func NewResolver() *Resolver {
	return &Resolver{MarketSocket: map[int][]chan string{}}
}
