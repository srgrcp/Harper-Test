package resolver

// Query resolver
type Query struct{}

// NewRoot method
func NewRoot() *Query {
	return &Query{}
}

// Test endpoint
func (*Query) Test() string {
	return "GraphQL Server is working."
}
