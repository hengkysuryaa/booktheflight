package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/hengkysuryaa/booktheflight/backend/models"
	"gorm.io/gorm"
)

type IRepository interface {
	GetFlight(ctx context.Context, id uuid.UUID) (f models.FlightSegment, err error)
	GetPassenger(ctx context.Context, id uuid.UUID) (p models.Passenger, err error)
	GetAircraft(ctx context.Context, code string) (a models.Aircraft, err error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) IRepository {
	return &repository{
		db: db.Debug(),
	}
}

func (r *repository) GetFlight(ctx context.Context, id uuid.UUID) (f models.FlightSegment, err error) {
	err = r.db.Preload("Bookings").Where("uuid = ?", id).First(&f).Error
	return
}

func (r *repository) GetPassenger(ctx context.Context, id uuid.UUID) (p models.Passenger, err error) {
	err = r.db.
		Preload("Document").
		Preload("Address").
		Preload("Preferences").
		Preload("FrequentFlyers").
		Where("uuid = ?", id).First(&p).Error
	return
}

func (r *repository) GetAircraft(ctx context.Context, code string) (a models.Aircraft, err error) {
	err = r.db.
		Preload("Cabins").
		Preload("Cabins.RowDisabledCauses").
		Preload("Cabins.SeatRows.Seats.Prices").
		Preload("Cabins.SeatRows.Seats.Taxes").
		Preload("Cabins.SeatRows.Seats.Totals").
		Where("code = ?", code).First(&a).Error

	return
}
