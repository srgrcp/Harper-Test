package resolver

import "github.com/srgrcp/Harper-Test/model"

// ServiceResolver struct
type ServiceResolver struct {
	service model.Service
}

// UUID resolver
func (r *ServiceResolver) UUID() string {
	return r.service.UUID
}

// Customer resolver
func (r *ServiceResolver) Customer() *UserResolver {
	return &UserResolver{customer: r.service.Customer}
}

// Technician resolver
func (r *ServiceResolver) Technician() *UserResolver {
	return &UserResolver{technician: r.service.Technician}
}

// Comment resolver
func (r *ServiceResolver) Comment() *string {
	comment := &r.service.Comment
	if *comment == "" {
		return nil
	}
	return comment
}

// State resolver
func (r *ServiceResolver) State() string {
	return r.service.State
}

// Rate resolver
func (r *ServiceResolver) Rate() int32 {
	return r.service.Rate
}

// Link resolver
func (r *ServiceResolver) Link() string {
	return "http://localhost:8080/api/order/" + r.service.UUID
}

// NewServiceOrderArgs mutation parameters
type NewServiceOrderArgs struct {
	CustomerNit string
	Comment     *string
}

// FinishServiceArgs mutation parameters
type FinishServiceArgs struct {
	UUID          string
	TechnicianNit string
}

// RateServiceArgs mutation parameters
type RateServiceArgs struct {
	UUID        string
	CustomerNit string
	Rate        int32
}

// GetServiceArgs query parameters
type GetServiceArgs struct {
	UUID string
}
