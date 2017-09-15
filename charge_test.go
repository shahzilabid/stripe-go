package stripe

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	"github.com/stripe/stripe-go/form"
)

func TestChargeParams_AppendTo(t *testing.T) {
	{
		params := &ChargeParams{Amount: 123}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"123"}, body.Get("amount"))
	}

	{
		params := &ChargeParams{Dest: "acct_123"}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"acct_123"}, body.Get("destination[account]"))
	}

	{
		params := &ChargeParams{Fraud: "suspicious"}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"suspicious"}, body.Get("fraud_details[user_report]"))
	}

	{
		params := &ChargeParams{Source: &SourceParams{Card: &CardParams{Number: "4242424242424242"}}}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"4242424242424242"}, body.Get("source[number]"))
		assert.Equal(t, []string{"card"}, body.Get("source[object]"))
	}
}
