package payment

import (
	"github.com/pressly/chi"

	"github.com/ecommerce/auth"
	"github.com/ecommerce/log"
	"github.com/ecommerce/route"
)

// API struct
type API struct{}

// New api
func New() *API { return new(API) }

func (api *API) Register(r chi.Router) {
	log.Debug("Registering order endpoints...")

	w := route.NewWrapper(r, route.Options{
		Timeout: route.Timeout{
			Timeout:  2,
			Response: map[string]string{"Timeout": "max 2 seconds"},
		},
	})

	handleGet(w)
	handlePost(w)

}

func handleGet(w *route.RouterWrapper) {

}

func handlePost(w *route.RouterWrapper) {
	w.Post("/payment/do-payment", auth.AuthSession(doPayment))
	w.Post("/payment/check-payment", auth.AuthSession(checkPayment))
}
