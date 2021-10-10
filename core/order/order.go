package order

import (
	"context"
	"log"

	"github.com/ecommerce/sqlt"
	"github.com/ecommerce/database"
)

type Order struct {
	ID int64 `json:"id"  db:"id"`
}

func (order *Order) CreateNewOrder(ctx context.Context, tx *sqlx.Tx) (orderID int64, err error) {
	orderID, err = tx.NamedExec(QueryCreateOrder, order)
	if err != nil {
		return orderID, errors.AddTrace(err)
	}

	DeleteOrderCache(ctx, order.UserID)

	return orderID, nil
}

func GetOrderByID(ctx context.Context, orderID int64) (orderData Order, err error) {

	orderData, err = GetOrderByID(ctx, orderID)
	if err == nil {
		return orderData, nil
	}

	err = stmt.GetOrderByID.Get(&orderData, orderID)
	if err != nil {
		return orderData, errors.AddTrace(err)
	}

	SetOrderByID(ctx, orderData)

	return orderData, nil
}