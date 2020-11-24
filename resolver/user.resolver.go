package resolver

import "github.com/srgrcp/Harper-Test/model"

// UserResolver struct
type UserResolver struct {
	//user model.User
	customer   model.Customer
	technician model.Technician
}

func isCustomer(resolver *UserResolver) bool {
	return resolver.customer.Nit != ""
}

// Nit resolver
func (r *UserResolver) Nit() string {
	if isCustomer(r) {
		return r.customer.Nit
	}
	return r.technician.Nit
}

// Name resolver
func (r *UserResolver) Name() string {
	if isCustomer(r) {
		return r.customer.Name
	}
	return r.technician.Name
}

// Services resolver
func (r *UserResolver) Services() *[]*ServiceResolver {
	var services []model.Service
	if isCustomer(r) {
		services = r.customer.Service
	} else {
		services = r.technician.Service
	}
	resolvers := []*ServiceResolver{}
	for _, service := range services {
		resolvers = append(resolvers, &ServiceResolver{service: service})
	}
	return &resolvers
}

// NewCustomerArgs Create customer mutation parameters
type NewCustomerArgs struct {
	Name string
	Nit  string
}
