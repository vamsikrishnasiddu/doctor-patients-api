package doctors

import (
	"strings"

	"github.com/vamsikrishnasiddu/doctors_patiens-api/patients"
	"github.com/vamsikrishnasiddu/doctors_patiens-api/rest_errors"
)

type Doctor struct {
	Id          int                `json:"id"`
	DoctorName  string             `json:"doctor_name"`
	DateCreated string             `json:"date_created"`
	Patients    []patients.Patient `json:"patients"`
}

type Doctors []Doctor

func (doctor *Doctor) Validate() rest_errors.RestErr {
	doctor.DoctorName = strings.TrimSpace(doctor.DoctorName)

	if doctor.DoctorName == "" {
		return rest_errors.NewBadRequestError("invalid name")
	}

	return nil
}
