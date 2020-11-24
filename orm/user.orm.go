package orm

import "github.com/srgrcp/Harper-Test/model"

// CreateCustomer handler
func CreateCustomer(customer *model.Customer) error {
	result := db.Create(customer)
	return result.Error
}

// GetCustomerByNit getter
func GetCustomerByNit(nit string) (*model.Customer, error) {
	customer := &model.Customer{}
	result := db.First(customer, "nit = ?", nit)
	if result.Error != nil {
		return nil, result.Error
	}
	return customer, nil
}

// GetTechnicianByNit getter
func GetTechnicianByNit(nit string) (*model.Technician, error) {
	technician := &model.Technician{}
	result := db.First(technician, "nit = ?", nit)
	if result.Error != nil {
		return nil, result.Error
	}
	return technician, nil
}

// GenerateTestData populates technician table
func GenerateTestData() (*[]*model.Technician, error) {
	result := db.First(&model.Technician{}, "nit = ?", "16854215")
	if result.Error == nil {
		return nil, result.Error
	}
	technicians := &[]*model.Technician{
		{User: model.User{Name: "Eduardo Saavedra", Nit: "16854215"}},
		{User: model.User{Name: "José Pedro Freyre", Nit: "55134542"}},
		{User: model.User{Name: "Giovanni Rovelli", Nit: "133854844"}},
		{User: model.User{Name: "Simon Crowther", Nit: "94677152"}},
		{User: model.User{Name: "Antonio Moreno", Nit: "41135842"}},
		{User: model.User{Name: "Martín Sommer", Nit: "64822458"}},
		{User: model.User{Name: "Pedro Afonso", Nit: "71563521"}},
	}
	for _, technician := range *technicians {
		db.Create(technician)
	}
	return technicians, nil
}

// GetRandomTechnician get a random technician
func GetRandomTechnician() *model.Technician {
	technician := &model.Technician{}
	db.Order("RANDOM()").Take(technician)
	return technician
}
