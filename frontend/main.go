package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/stripe/stripe-go"
)

stripe.Key = "sk_test_ow8Kmmfn4zvqB9MavwfOpLHT00AHft5LPw"

params := &stripe.PaymentIntentParams{
  Amount: stripe.Int64(1099),
  Currency: stripe.String(string(stripe.CurrencyUSD)),
}
// Verify your integration in this guide by including this parameter
params.AddMetadata("integration_check", "accept_a_payment")

pi, _ := paymentintent.New(params)
