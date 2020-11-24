package orm

import (
	uuid "github.com/satori/go.uuid"
	"github.com/srgrcp/Harper-Test/model"
)

// CreateServiceOrder handler
func CreateServiceOrder(customer *model.Customer, technician *model.Technician, comment string) (*model.Service, error) {
	service := &model.Service{
		UUID:       uuid.NewV4().String(),
		Customer:   *customer,
		Technician: *technician,
		State:      "CREATED",
		Comment:    comment,
		Rate:       0,
	}
	result := db.Create(service)
	if result.Error != nil {
		return nil, result.Error
	}
	return service, nil
}

// GetServiceOrder by uuid
func GetServiceOrder(uuid string) *model.Service {
	service := &model.Service{}
	result := db.Preload("Customer").Preload("Technician").First(service, "uuid = ?", uuid)
	if result.Error != nil {
		return nil
	}
	return service
}

func updateService(uuid string, column string, value interface{}) error {
	result := db.Model(&model.Service{}).Where("uuid = ?", uuid).Update(column, value)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// FinishService handler
func FinishService(uuid string) error {
	err := updateService(uuid, "state", "FINISHED")
	if err != nil {
		return err
	}
	return nil
}

// RateService handler
func RateService(uuid string, rate int32) error {
	err := updateService(uuid, "rate", rate)
	if err != nil {
		return err
	}
	return nil
}
