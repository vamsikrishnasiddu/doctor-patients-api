package patients

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vamsikrishnasiddu/doctors_patiens-api/patients"
	"github.com/vamsikrishnasiddu/doctors_patiens-api/rest_errors"
	"github.com/vamsikrishnasiddu/doctors_patiens-api/services"
	"github.com/vamsikrishnasiddu/doctors_patiens-api/utils"
)

func getPatientId(patientId string) (int64, error) {
	patient_id, err := strconv.ParseInt(patientId, 10, 64)
	if err != nil {
		return 0, rest_errors.NewBadRequestError("doctor id should be a number")
	}

	return patient_id, nil
}

func CreatePatient(c *gin.Context) {
	var patient patients.Patient

	//patientsArray := utils.PatientsArray
	doctorId := c.Param("doctor_id")

	doctor_id, err := getPatientId(doctorId)
	if err != nil {
		return
	}

	if err := c.ShouldBindJSON(&patient); err != nil {
		fmt.Println("err,err")
		return
	}

	doctor, err := services.DoctorsService.GetDoctor(doctor_id)

	if err != nil {
		fmt.Println("err", err)
	}
	doctor.Patients = append(doctor.Patients, patient)

	c.JSON(http.StatusCreated, patient)

}

func GetPatient(c *gin.Context) {
	var patient patients.Patient

	patientsArray := utils.PatientsArray

	patient_id, err := getPatientId(c.Param("patient_id"))

	if err != nil {
		fmt.Println("err", err)
		return
	}

	for i := range patientsArray {
		if patientsArray[i].Id == int(patient_id) {
			patient = patientsArray[i]
		}
	}

	c.JSON(http.StatusOK, patient)

}
