package services

import (
	"fmt"

	"github.com/vamsikrishnasiddu/doctors_patiens-api/doctors"
	"github.com/vamsikrishnasiddu/doctors_patiens-api/rest_errors"
	"github.com/vamsikrishnasiddu/doctors_patiens-api/utils/date_utils"
)

var (
	DoctorsService doctorsServiceInterface = &doctorService{}
)

type doctorService struct{}

type doctorsServiceInterface interface {
	CreateDoctor(doctors.Doctor) (*doctors.Doctor, rest_errors.RestErr)
	GetDoctor(int64) (*doctors.Doctor, rest_errors.RestErr)
	UpdateDoctor(bool, doctors.Doctor) (*doctors.Doctor, rest_errors.RestErr)
	DeleteDoctor(int64) rest_errors.RestErr
}

func (s *doctorService) CreateDoctor(doctor doctors.Doctor) (*doctors.Doctor, rest_errors.RestErr) {
	if err := doctor.Validate(); err != nil {
		return nil, err
	}
	doctor.DateCreated = date_utils.GetNowDBFormat()
	fmt.Println("dateCreated", doctor.DateCreated)
	if err := doctor.Save(); err != nil {
		return nil, err
	}
	return &doctor, nil
}

func (s *doctorService) GetDoctor(doctorId int64) (*doctors.Doctor, rest_errors.RestErr) {
	dao := &doctors.Doctor{Id: int(doctorId)}
	if err := dao.Get(); err != nil {
		return nil, err
	}
	return dao, nil
}

func (s *doctorService) UpdateDoctor(isPartial bool, doctor doctors.Doctor) (*doctors.Doctor, rest_errors.RestErr) {

	current := &doctors.Doctor{Id: doctor.Id}

	if err := current.Get(); err != nil {
		return nil, err
	}

	if isPartial {
		if doctor.DoctorName != "" {
			current.DoctorName = doctor.DoctorName
		}
	} else {
		current.DoctorName = doctor.DoctorName
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil

}

func (s *doctorService) DeleteDoctor(doctorId int64) rest_errors.RestErr {
	dao := &doctors.Doctor{Id: int(doctorId)}

	return dao.Delete()

}
