package app

import (
	"github.com/vamsikrishnasiddu/doctors_patiens-api/controllers/doctors"
	"github.com/vamsikrishnasiddu/doctors_patiens-api/controllers/ping"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.POST("/doctors", doctors.CreateDoctor)
	router.GET("/doctors/:doctor_id", doctors.GetDoctor)
	router.PATCH("/doctors/:doctor_id", doctors.UpdateDoctor)
	router.PUT("/doctors/:doctor_id", doctors.UpdateDoctor)
	router.DELETE("/doctors/:doctor_id", doctors.DeleteDoctor)

}
