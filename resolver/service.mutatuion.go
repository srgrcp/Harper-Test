package resolver

import "github.com/srgrcp/Harper-Test/orm"

// CreateServiceOrder endpoint
func (*Query) CreateServiceOrder(newServiceOrderArgs NewServiceOrderArgs) (*ServiceResolver, error) {
	technician := orm.GetRandomTechnician()
	customer, err := orm.GetCustomerByNit(newServiceOrderArgs.CustomerNit)
	if customer == nil {
		return nil, err
	}
	service, err := orm.CreateServiceOrder(customer, technician, *newServiceOrderArgs.Comment)
	if service == nil {
		return nil, err
	}
	return &ServiceResolver{
		service: *service,
	}, nil
}

// FinishService endpoint
func (*Query) FinishService(args FinishServiceArgs) (*ServiceResolver, error) {
	_, err := orm.GetTechnicianByNit(args.TechnicianNit)
	if err != nil {
		return nil, err
	}
	err = orm.FinishService(args.UUID)
	if err != nil {
		return nil, err
	}
	service := orm.GetServiceOrder(args.UUID)
	return &ServiceResolver{
		service: *service,
	}, nil
}

// RateService endpoint
func (*Query) RateService(args RateServiceArgs) (*ServiceResolver, error) {
	_, err := orm.GetCustomerByNit(args.CustomerNit)
	if err != nil {
		return nil, err
	}
	err = orm.RateService(args.UUID, args.Rate)
	if err != nil {
		return nil, err
	}
	service := orm.GetServiceOrder(args.UUID)
	return &ServiceResolver{
		service: *service,
	}, nil
}
