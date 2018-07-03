package models

import (
	"errors"
	"fmt"

	"github.com/dchest/uniuri"
)

// validationErrors - custom error type that can take many errors
type validationErrors []error

// Error - implementation of error interface
func (ve validationErrors) Error() string {
	var errStr string
	for _, v := range ve {
		errStr += fmt.Sprintf("%s\n", v.Error())
	}
	return errStr
}

// Validate - implementation of Validatable
func (veh *VehicleQuery) Validate() error {
	var err validationErrors
	if veh.Id == "" && veh.Query == "" {
		if veh.Id == "" {
			err = append(err, errors.New("Missing id information"))
		}
		if veh.Query == "" {
			err = append(err, errors.New("Missing query information"))
		}
	}
	return err
}

// Validate - implementation of Validatable
func (veh *Vehicle) Validate(put bool) error {
	var err validationErrors
	if veh.Id == "" && !put {
		veh.Id = uniuri.NewLen(16)
	} else if veh.Id == "" && put {
		err = append(err, errors.New("Id must be present"))
	}
	if veh.Manufacturer == "" {
		err = append(err, errors.New("Manufacturer must be present"))
	}
	if veh.Model == "" {
		err = append(err, errors.New("Model must be present"))
	}
	if veh.Price <= 0 {
		err = append(err, errors.New("Price must be real"))
	}
	if veh.Mileage <= 0 {
		err = append(err, errors.New("Mileage must be real"))
	}
	if len(err) == 0 {
		return nil
	}
	return err
}
