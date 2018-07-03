package models

import (
	"errors"
	"fmt"

	"github.com/dchest/uniuri"
)

// errorsList - list of errors
type errorsList []error

// Error - compress the list to a single string
func (el errorsList) Error() string {
	var errStr string
	for _, v := range el {
		errStr += fmt.Sprintf("%s\n", v.Error())
	}
	return errStr
}

// Validate - the Query submission
func (qu *VehicleQuery) Validate() error {
	var err errorsList
	if qu.Id == "" && qu.Query == "" {
		if qu.Id == "" {
			err = append(err, errors.New("Missing id information"))
		}
		if qu.Query == "" {
			err = append(err, errors.New("Missing query information"))
		}
	}

	// Check the errors list
	if len(err) == 0 {
		return nil
	}
	return err
}

// Validate - the Vehicle submission
func (veh *Vehicle) Validate(put bool) error {
	var err errorsList
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

	// Check the errors list
	if len(err) == 0 {
		return nil
	}
	return err
}
