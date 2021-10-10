package mpayment

type DoPaymentResponse struct {
	UserID     	int64       `json:"user_id"`
	OrderID    	int64 		`json:"order_id"`
	PaymentID	int64     	`json:"payment_id"`
	ItemList   	[]cart.Item `json:"item_list"`
}


func DoPayment(ctx context.Context, userID, totalPrice int64) (resp DoPaymentResponse, err error) {
	err = ValidateRequest(userID, totalPrice)
	if err != nil {
		return resp, errors.AddTrace(err)
	}

	db, err := database.Get(database.Core)
	if err != nil {
		return resp, errors.AddTrace(err)
	}

	tx, err := db.Master().Beginx()
	if err != nil {
		return resp, errors.AddTrace(err)
	}
	defer tx.Rollback()

	orderID, err := order.CreateNewOrder(userID, totalPrice)
	if err != nil {
		return resp, errors.AddTrace(err)
	}

	itemList, err := cart.GetItemListByUserID(userID)
	if err != nil {
		return resp, errors.AddTrace(err)
	}

	_, err = order.CreateNewOrderItem(orderID, itemList)
	if err != nil {
		return resp, errors.AddTrace(err)
	}

	paymentID, err := payment.CreatePaymentRequest(userID, orderID, totalPrice)
	if err != nil {
		return resp, errors.AddTrace(err)
	}

	_, err = order.UpdateOrderPaymentID(orderID, paymentID)
	if err != nil {
		return resp, errors.AddTrace(err)
	}

	tx.Commit()

	err = publisher.Publish(common.NSQClearCart, ClearCartParam{
		UserID: userID,
	})
	if err != nil {
		log.ErrorTrace(err, fmt.Sprintf("[DoPayment][PublishClearCart][UserID:%d] Failed to publish clear cart", userID))
		return resp, errors.AddTrace(err)
	}

	resp = DoPaymentResponse{
		UserID: userID,
		OrderID: orderID,
		PaymentID: paymentID,
		ItemList: itemList,
	}

	return resp, nil
}