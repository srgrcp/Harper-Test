package resolver

import (
	"github.com/srgrcp/Harper-Test/model"
	"github.com/srgrcp/Harper-Test/orm"
)

// PopulateTechnicians endpoint
func (*Query) PopulateTechnicians() *[]*UserResolver {
	technicians, _ := orm.GenerateTestData()
	if technicians == nil {
		return nil
	}
	resolvers := []*UserResolver{}
	for _, technician := range *technicians {
		resolvers = append(resolvers, &UserResolver{
			technician: *technician,
		})
	}
	return &resolvers
}

// CreateCustomer endpoint
func (*Query) CreateCustomer(newCustomerArgs NewCustomerArgs) (*UserResolver, error) {
	customer := &model.Customer{
		User: model.User{
			Name: newCustomerArgs.Name,
			Nit:  newCustomerArgs.Nit,
		},
	}
	err := orm.CreateCustomer(customer)
	if err != nil {
		return nil, err
	}
	return &UserResolver{
		customer: *customer,
	}, nil
}
