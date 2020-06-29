package main

import (
	"fmt"

	"github.com/princeparmar/stripe/constants"
	"github.com/stripe/stripe-go/v71"
	"github.com/stripe/stripe-go/v71/paymentintent"
)

func init() {
	stripe.Key = constants.StripeKey
}

func main() {
	stripe.Key = "sk_test_WEmL0vr6zksZXBhfOEjutNqR"

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(1099),
		Currency: stripe.String(string(stripe.CurrencyINR)),
	}
	// Verify your integration in this guide by including this parameter
	params.AddMetadata("integration_check", "accept_a_payment")

	pi, _ := paymentintent.New(params)
	fmt.Printf("%#v\n", pi)
}
