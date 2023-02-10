package postgres

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres/general"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres/model"
)

type OrderRepository struct {
	db  *sqlx.DB
	dao general.DaoI[model.Order]
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{
		db:  db,
		dao: NewOrderDao(db),
	}
}

func (r *OrderRepository) Save(ctx context.Context, order *model.Order) (int, error) {
	params := map[string]interface{}{
		"from":         order.From,
		"to":           order.To,
		"date":         time.Now(),
		"status":       model.InProgress,
		"driver_id":    order.DriverId,
		"taxi_type_id": order.TaxiTypeId,
		"user_id":      order.UserId,
	}

	return r.dao.Save(params)
}

func (r *OrderRepository) Update(ctx context.Context, order *model.Order) error {
	params := map[string]interface{}{}

	if order.From != "" {
		params["from"] = order.From
	}

	if order.To != "" {
		params["to"] = order.To
	}

	if order.Status != 0 {
		params["status"] = order.Status
	}

	return r.dao.Update(params, order.Id)
}

func (r *OrderRepository) FindAll(ctx context.Context) ([]model.Order, error) {
	return r.dao.FindByFields([]model.Order{}, nil)
}

func (r *OrderRepository) FindByFields(ctx context.Context, params map[string]interface{}) ([]model.Order, error) {
	return r.dao.FindByFields([]model.Order{}, params)
}
