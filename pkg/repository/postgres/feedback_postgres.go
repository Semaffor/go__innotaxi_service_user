package postgres

import (
	"context"

	"github.com/guregu/null"
	"github.com/jmoiron/sqlx"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres/general"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository/postgres/model"
)

type FeedbackRepository struct {
	db  *sqlx.DB
	dao general.DaoI[model.Feedback]
}

func NewFeedbackRepository(db *sqlx.DB) *FeedbackRepository {
	return &FeedbackRepository{
		db:  db,
		dao: NewFeedbackDao(db),
	}
}

func (r *FeedbackRepository) Save(ctx context.Context, feedback *model.Feedback) (int, error) {
	params := map[string]interface{}{
		"driver_id":   feedback.DriverId,
		"customer_id": feedback.CustomerId,
		"order_id":    feedback.OrderId,
	}

	return r.dao.Save(params)
}

func (r *FeedbackRepository) Update(ctx context.Context, feedback *model.Feedback) error {
	params := map[string]interface{}{}

	if feedback.MarkFromUser != null.FloatFrom(0) {
		params["mark_from_user"] = feedback.MarkFromUser
	}
	if feedback.MarkFromDriver != null.FloatFrom(0) {
		params["mark_from_driver"] = feedback.MarkFromUser
	}
	if feedback.FeedbackFromUser != null.StringFrom("") {
		params["feedback_from_user"] = feedback.FeedbackFromUser
	}
	if feedback.FeedbackFromDriver != null.StringFrom("") {
		params["feedback_from_driver"] = feedback.FeedbackFromDriver
	}

	return r.dao.Update(params, feedback.Id)
}

func (r *FeedbackRepository) FindAll(ctx context.Context) ([]model.Feedback, error) {
	return r.dao.FindByFields([]model.Feedback{}, nil)
}

func (r *FeedbackRepository) FindByFields(
	ctx context.Context,
	params map[string]interface{},
) ([]model.Feedback, error) {
	return r.dao.FindByFields([]model.Feedback{}, params)
}
