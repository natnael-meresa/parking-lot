package repository

import (
	"context"
	"database/sql"
	"parking-lot/internal/model"
	"parking-lot/pkg/logger"
	"time"
)

type RentalRepository interface {
	CreateRental(ctx context.Context, rentalDto model.Rental) (int, error)
	ListRentals(ctx context.Context) ([]model.Rental, error)
	GetRentalByCarID(ctx context.Context, carID int) (*model.Rental, error)
	UpdateRentalKilometersDriven(ctx context.Context, carID int, kilometers_driven int) error
	UpdateRentalEndDate(ctx context.Context, carID int, end_date time.Time) error
}

type rentalRepository struct {
	db *sql.DB
	log *logger.Logger
}

func NewRental(db *sql.DB, log *logger.Logger) RentalRepository {
	return &rentalRepository{
		db: db,
		log: log,
	}
}

func (r *rentalRepository) CreateRental(ctx context.Context, rentalDto model.Rental) (int, error) {
	var id int

	err := r.db.QueryRowContext(ctx, "INSERT INTO rentals(car_id, start_date, kilometers_driven) VALUES($1, $2, $4) RETURNING id", rentalDto.CarID, rentalDto.StartDate, rentalDto.KilometersDriven).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *rentalRepository) ListRentals(ctx context.Context) ([]model.Rental, error) {
	var rentals []model.Rental
	var rental model.Rental

	rows, err := r.db.QueryContext(ctx, "SELECT * FROM rentals")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&rental.ID, &rental.CarID, &rental.StartDate, &rental.EndDate, &rental.KilometersDriven)
		if err != nil {
			return nil, err
		}
		rentals = append(rentals, rental)
	}

	return rentals, nil
}

func (r *rentalRepository) GetRentalByCarID(ctx context.Context, carID int) (*model.Rental, error) {
	var rental model.Rental

	err := r.db.QueryRowContext(ctx, "SELECT * FROM rentals WHERE car_id = $1", carID).Scan(&rental.ID, &rental.CarID, &rental.StartDate, &rental.EndDate, &rental.KilometersDriven)
	if err != nil {
		return nil, err
	}

	return &rental, nil
}

func (r *rentalRepository) UpdateRentalKilometersDriven(ctx context.Context, carID int, kilometers_driven int) error {
	_, err := r.db.ExecContext(ctx, "UPDATE rentals SET kilometers_driven = $1 WHERE car_id = $2", kilometers_driven, carID)
	if err != nil {
		return err
	}

	return nil
}

func (r *rentalRepository) UpdateRentalEndDate(ctx context.Context, carID int, end_date time.Time) error {
	_, err := r.db.ExecContext(ctx, "UPDATE rentals SET end_date = $1 WHERE car_id = $2", end_date, carID)
	if err != nil {
		return err
	}

	return nil
}
