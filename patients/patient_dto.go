package patients

type Patient struct {
	Id            int    `json:"id"`
	PatientName   string `json:"patient_name"`
	BloodPressure int    `json:"blood_pressure"`
	Sugar         int    `json:"sugar"`
	HeartBeat     int    `json:"heart_beat"`
	OxygenLevel   int    `json:"oxygen_level"`
}
