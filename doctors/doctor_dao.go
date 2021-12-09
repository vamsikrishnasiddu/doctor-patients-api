package doctors

import (
	"errors"
	"fmt"

	"github.com/vamsikrishnasiddu/doctors_patiens-api/datasources/mysql/doctors_db"
	"github.com/vamsikrishnasiddu/doctors_patiens-api/rest_errors"
)

const (
	queryInsertDoctor = "INSERT INTO doctors(doctor_name, date_created) VALUES(?, ?);"
	queryGetDoctor    = "SELECT id, doctor_name, date_created  FROM doctors WHERE id=?;"
	queryUpdateDoctor = "UPDATE doctors SET doctor_name=? WHERE id=?;"
	queryDeleteDoctor = "DELETE FROM doctors WHERE id=?;"
)

func (doctor *Doctor) Get() rest_errors.RestErr {
	stmt, err := doctors_db.Client.Prepare(queryGetDoctor)
	if err != nil {
		fmt.Println("err", err)
		return rest_errors.NewInternalServerError("error when tying to get doctor", errors.New("database error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(doctor.Id)

	if getErr := result.Scan(&doctor.Id, &doctor.DoctorName, &doctor.DateCreated); getErr != nil {
		fmt.Println("getErr", getErr)
		return rest_errors.NewInternalServerError("error when tying to get user", errors.New("database error"))
	}
	return nil
}

func (doctor *Doctor) Save() rest_errors.RestErr {
	stmt, err := doctors_db.Client.Prepare(queryInsertDoctor)
	if err != nil {
		fmt.Println("err:", err)
		return rest_errors.NewInternalServerError("error when tying to save doctor", errors.New("database error"))
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(doctor.DoctorName, doctor.DateCreated)
	if saveErr != nil {
		fmt.Println("saveErr", saveErr)
		return rest_errors.NewInternalServerError("error when tying to save doctor", errors.New("database error"))
	}

	doctor_id, err := insertResult.LastInsertId()
	if err != nil {
		return rest_errors.NewInternalServerError("error when tying to save doctor", errors.New("database error"))
	}
	doctor.Id = int(doctor_id)

	return nil
}

func (doctor *Doctor) Update() rest_errors.RestErr {
	stmt, err := doctors_db.Client.Prepare(queryUpdateDoctor)
	if err != nil {
		fmt.Println("err", err)
		return rest_errors.NewInternalServerError("error when tying to update user", errors.New("database error"))
	}
	defer stmt.Close()

	_, err = stmt.Exec(doctor.DoctorName, doctor.Id)
	if err != nil {
		fmt.Println("updateErr", err)
		return rest_errors.NewInternalServerError("error when tying to update user", errors.New("database error"))
	}
	return nil
}

func (doctor *Doctor) Delete() rest_errors.RestErr {
	stmt, err := doctors_db.Client.Prepare(queryDeleteDoctor)
	if err != nil {
		fmt.Println("deleteErr1", err)
		return rest_errors.NewInternalServerError("error when tying to update user", errors.New("database error"))
	}
	defer stmt.Close()

	if _, err = stmt.Exec(doctor.Id); err != nil {
		fmt.Println("deleteErr2", err)
		return rest_errors.NewInternalServerError("error when tying to save user", errors.New("database error"))
	}
	return nil
}
