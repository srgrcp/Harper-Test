package resolver

import "github.com/srgrcp/Harper-Test/orm"

// GetService endpoint resolver
func (*Query) GetService(args GetServiceArgs) *ServiceResolver {
	service := orm.GetServiceOrder(args.UUID)
	return &ServiceResolver{
		service: *service,
	}
}
