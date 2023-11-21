package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfItGetsAnErrorIfIDIsBlank(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "id is required")
}

func TestIfThePriceIsBlank(t *testing.T) {
	order := Order{ID: "1213"}
	assert.Error(t, order.Validate(), "price must be greater than zero")
}

func TestIfTheTaxIsBlank(t *testing.T) {
	order := Order{ID: "1213", Price: 10.0}
	assert.Error(t, order.Validate(), "invalid tax")
}

func TestFinalPrice(t *testing.T) {
	order := Order{ID: "1213", Price: 10.0, Tax: 5.0}
	assert.NoError(t, order.Validate())
	assert.Equal(t, "1213", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 5.0, order.Tax)
	order.CalculateFinalPrice()
	assert.Equal(t, 15.0, order.FinalPrice)
}
