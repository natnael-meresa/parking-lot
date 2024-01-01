package usecase

import (
	"context"
	"parking-lot/internal/controller/dto"
	"parking-lot/internal/mocks"
	"parking-lot/internal/model"
	"parking-lot/pkg/logger"
	"testing"

	"github.com/stretchr/testify/mock"
)


func TestCreateCar(t *testing.T) {
	t.Run("success", func(t *testing.T) {

		mockCar := model.Car{
			Registration: "1234",
			Model:        "Corolla",
			Mileage:      1000,
		}

		mockCarRepo := new(mocks.MockCarRepository)
		mockRentRepo := new(mocks.MockRentalRepository)
		logger, _ := logger.New()
		carUs := NewCar(mockCarRepo, mockRentRepo, logger)

		mockCarRepo.On("CreateCar", mock.Anything, mockCar).Return(1, nil)

		id, err := carUs.CreateCar(context.Background(), dto.Car{
			Registration: "1234",
			Model:        "Corolla",
			Mileage:      1000,
		})

		if err != nil {
			t.Errorf("error should be nil, got %v", err)
		}

		if id != 1 {
			t.Errorf("id should be 1, got %v", id)
		}
	})

	t.Run("duplicate registration", func(t *testing.T) {
		
		mockCar := model.Car{
			Registration: "1234",
			Model:        "Corolla",
			Mileage:      1000,
		}

		mockCarRepo := new(mocks.MockCarRepository)
		mockRentRepo := new(mocks.MockRentalRepository)
		logger, _ := logger.New()

		carUs := NewCar(mockCarRepo, mockRentRepo, logger)

		mockCarRepo.On("GetCarByRegistration", mock.Anything, mockCar.Registration).Return(&mockCar, nil)

		_, err := carUs.CreateCar(context.Background(), dto.Car{
			Registration: "1234",
			Model:        "Corolla",
			Mileage:      1000,
		})

		if err == nil {
			t.Errorf("error should not be nil, got %v", err)
		}
	})

	t.Run("error", func(t *testing.T) {
		
		mockCar := model.Car{
			Registration: "1234",
			Model:        "Corolla",
			Mileage:      1000,
		}

		mockCarRepo := new(mocks.MockCarRepository)
		mockRentRepo := new(mocks.MockRentalRepository)
		logger, _ := logger.New()
		carUs := NewCar(mockCarRepo, mockRentRepo, logger)

		mockCarRepo.On("CreateCar", mock.Anything, mockCar).Return(0, mock.Anything)

		_, err := carUs.CreateCar(context.Background(), dto.Car{
			Registration: "1234",
			Model:        "Corolla",
			Mileage:      1000,
		})

		if err == nil {
			t.Errorf("error should not be nil, got %v", err)
		}
	})
}

func TestListCars(t *testing.T) {
	t.Run("success", func(t *testing.T) {

		mockCars := []model.Car{
			{
				ID:           1,
				Registration: "1234",
				Model:        "Corolla",
				Mileage:      1000,
			},
			{
				ID:           2,
				Registration: "5678",
				Model:        "Corolla",
				Mileage:      2000,
			},
		}

		mockCarRepo := new(mocks.MockCarRepository)
		mockRentRepo := new(mocks.MockRentalRepository)
		logger, _ := logger.New()
		carUs := NewCar(mockCarRepo, mockRentRepo, logger)

		mockCarRepo.On("ListCars", mock.Anything).Return(mockCars, nil)

		cars, err := carUs.ListCars(context.Background())

		if err != nil {
			t.Errorf("error should be nil, got %v", err)
		}

		if len(cars) != 2 {
			t.Errorf("cars length should be 2, got %v", len(cars))
		}
	})

	t.Run("error", func(t *testing.T) {

		mockCarRepo := new(mocks.MockCarRepository)
		mockRentRepo := new(mocks.MockRentalRepository)
		logger, _ := logger.New()

		carUs := NewCar(mockCarRepo, mockRentRepo, logger)

		mockCarRepo.On("ListCars", mock.Anything).Return([]model.Car{}, mock.Anything)

		_, err := carUs.ListCars(context.Background())

		if err == nil {
			t.Errorf("error should not be nil, got %v", err)
		}
	})
}

func TestRentCar(t *testing.T) {
	t.Run("success", func(t *testing.T) {

		mockCar := model.Car{
			ID:           1,
			Registration: "1234",
			Model:        "Corolla",
			Mileage:      1000,
			Available:    "available",
		}

		mockCarRepo := new(mocks.MockCarRepository)
		mockRentRepo := new(mocks.MockRentalRepository)
		logger, _ := logger.New()

		carUs := NewCar(mockCarRepo, mockRentRepo, logger)

		mockCarRepo.On("GetCarByRegistration", mock.Anything, mockCar.Registration).Return(&mockCar, nil)
		mockRentRepo.On("CreateRental", mock.Anything, mock.Anything).Return(1, nil)

		err := carUs.RentCar(context.Background(), mockCar.Registration)

		if err != nil {
			t.Errorf("error should be nil, got %v", err)
		}
	})

	t.Run("not found", func(t *testing.T) {

		mockCar := model.Car{
			ID:           1,
			Registration: "1234",
			Model:        "Corolla",
			Mileage:      1000,
			Available:    "available",
		}

		mockCarRepo := new(mocks.MockCarRepository)
		mockRentRepo := new(mocks.MockRentalRepository)
		logger, _ := logger.New()
		carUs := NewCar(mockCarRepo, mockRentRepo, logger)

		mockCarRepo.On("GetCarByRegistration", mock.Anything, mockCar.Registration).Return(&mockCar, mock.Anything)

		err := carUs.RentCar(context.Background(), mockCar.Registration)

		if err == nil {
			t.Errorf("error should not be nil, got %v", err)
		}
	})

	t.Run("error", func(t *testing.T) {

		mockCar := model.Car{
			ID:           1,
			Registration: "1234",
			Model:        "Corolla",
			Mileage:      1000,
			Available:    "available",
		}

		mockCarRepo := new(mocks.MockCarRepository)
		mockRentRepo := new(mocks.MockRentalRepository)
		logger, _ := logger.New()
		carUs := NewCar(mockCarRepo, mockRentRepo, logger)

		mockCarRepo.On("GetCarByRegistration", mock.Anything, mockCar.Registration).Return(&mockCar, nil)
		mockRentRepo.On("CreateRental", mock.Anything, mock.Anything).Return(0, mock.Anything)

		err := carUs.RentCar(context.Background(), mockCar.Registration)

		if err == nil {
			t.Errorf("error should not be nil, got %v", err)
		}

	})
}
