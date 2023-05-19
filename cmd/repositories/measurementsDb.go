package repositories

import (
	"fitness-api/cmd/models"
	"fitness-api/cmd/storage"
	"fmt"
	"time"
)

func CreateMeasurement(measurement models.Measurements) (models.Measurements, error) {
	db := storage.GetDB()
	sqlStatement := `
		INSERT INTO private.measurements (user_id, weight, height, body_fat, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, user_id, weight, height, body_fat, created_at`
	err := db.QueryRow(sqlStatement, measurement.UserId, measurement.Weight, measurement.Height, measurement.BodyFat, time.Now()).Scan(&measurement.Id, &measurement.UserId, &measurement.Weight, &measurement.Height, &measurement.BodyFat, &measurement.CreatedAt)
	if err != nil {
		return measurement, err
	}
	fmt.Printf("Created measurement: %#v\n", measurement)
	return measurement, nil
}

func UpdateMeasurement(measurement models.Measurements, id int) (models.Measurements, error) {
	db := storage.GetDB()
	sqlStatement := `
		UPDATE private.measurements
		SET weight = $2, height = $3, body_fat = $4, created_at = $5
		WHERE id = $1
		RETURNING id, user_id, weight, height, body_fat, created_at`
	err := db.QueryRow(sqlStatement, id, measurement.Weight, measurement.Height, measurement.BodyFat, time.Now()).Scan(&measurement.Id, &measurement.UserId, &measurement.Weight, &measurement.Height, &measurement.BodyFat, &measurement.CreatedAt)
	if err != nil {
		return models.Measurements{}, err
	}
	fmt.Printf("Updated measurement: %#v\n", measurement)
	measurement.Id = id
	fmt.Printf("Updated measurement: %#v\n", measurement)
	return measurement, nil
}
