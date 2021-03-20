package repository

import (
	"ApiRest/app/model/billing"
	"ApiRest/internal/storage"
)

// billingRepository handles communication with the user store
type billingRepository struct {
	db *storage.DbStore
}

//BillingRepositoryInterface define the user repository interface methods
type BillingRepositoryInterface interface {
	CreateBillingService(identity billing.Identify, key string, userID int) error
}

// NewBillingRepository implements the billing repository interface.
func NewBillingRepository(db *storage.DbStore) BillingRepositoryInterface {
	return &billingRepository{
		db,
	}
}

// CreateBillingService Create implements the method to persist a Payment user
func (r *billingRepository) CreateBillingService(identify billing.Identify, PaymentUserKey string, userID int) error {
	createUserQuery := `INSERT INTO billing (identify, key, user_id) 
		VALUES ($1, $2, $3)
		RETURNING id`

	stmt, err := r.db.Prepare(createUserQuery)
	defer stmt.Close()

	if err != nil {
		return nil
	}

	var paymentID int
	err = stmt.QueryRow(identify, PaymentUserKey, userID).Scan(&paymentID)
	return err
}
