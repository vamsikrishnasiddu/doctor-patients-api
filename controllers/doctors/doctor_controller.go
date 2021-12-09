package doctors

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vamsikrishnasiddu/doctors_patiens-api/doctors"
	"github.com/vamsikrishnasiddu/doctors_patiens-api/rest_errors"
	"github.com/vamsikrishnasiddu/doctors_patiens-api/services"
)

func getDoctorId(userIdParam string) (int64, rest_errors.RestErr) {
	doctorId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		return 0, rest_errors.NewBadRequestError("doctor id should be a number")
	}

	return doctorId, nil
}

func CreateDoctor(c *gin.Context) {

	var doctor doctors.Doctor

	if err := c.ShouldBindJSON(&doctor); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	fmt.Println(doctor)
	result, saveErr := services.DoctorsService.CreateDoctor(doctor)

	if saveErr != nil {
		c.JSON(saveErr.Status(), saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)

}

func GetDoctor(c *gin.Context) {
	doctorId, idErr := getDoctorId(c.Param("doctor_id"))

	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}
	doctor, getErr := services.DoctorsService.GetDoctor(doctorId)

	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	c.JSON(http.StatusOK, doctor)
}

func UpdateDoctor(c *gin.Context) {
	doctorId, idErr := getDoctorId(c.Param("doctor_id"))

	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	var doctor doctors.Doctor
	if err := c.ShouldBindJSON(&doctor); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return

	}

	doctor.Id = int(doctorId)

	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.DoctorsService.UpdateDoctor(isPartial, doctor)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, result)

}

func DeleteDoctor(c *gin.Context) {
	doctorId, idErr := getDoctorId(c.Param("doctor_id"))

	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	if err := services.DoctorsService.DeleteDoctor(doctorId); err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}
