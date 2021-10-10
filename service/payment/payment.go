package payment

import (
	"fmt"
	"net/http"
	"time"

	mpayment "github.com/ecommerce/model/payment"

	"github.com/ecommerce/errors"
	"github.com/ecommerce/log"
	"github.com/ecommerce/monitoring"
	"github.com/ecommerce/route"
	"github.com/ecommerce/util"
)

func doPayment(r *http.Request) (ho route.HandleObject, errs *errors.Errs) {
	ctx := r.Context()

	ddName := "DoPayment"
	ddTag := []string{}
	ddStatus := "failed"
	ddStartTime := time.Now()
	defer func() {
		monitoring.RecordDD(ddName, ddStatus, ddTag, ddStartTime, "APIhandler")
	}()

	resp := new(route.WS)

	userID, err := util.GetCtxUserID(ctx)
	if err != nil {
		log.ErrorTrace(err, fmt.Sprintf("[doPayment][GetCtxUserID][userID:%d]", r.FormValue("user_id")))
		return resp, errors.AddTrace(err)
	}

	totalPrice := util.Atoi64(r.FormValue("total_price"), 0)

	result, err := mpayment.DoPayment(userID, totalPrice)
	if err != nil {
		log.ErrorTrace(err, fmt.Sprintf("[doPayment][DoPayment][userID:%d]", userID))
		return resp, errors.AddTrace(err)
	}

	ddStatus = "success"

	resp.Data = result

	return resp, nil
}
