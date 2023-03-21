package stripe_client

import (
	"api/config"
	"net/http"

	"github.com/stripe/stripe-go/v74/client"
	"github.com/stripe/stripe-go/v74"
)

var StripeClient *client.API

func CreateStripeClient() {
  sc := client.New(config.Get("STRIPE_SECRET_KEY"), stripe.NewBackends(http.DefaultClient))
  StripeClient = sc
}
