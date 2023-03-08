package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//"testing"

func Test_If_I_Get_An_Error_If_ID_Is_Blank(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "invalid id")
}

func Test_If_I_Get_An_Error_If_Price_Is_Blank(t *testing.T) {
	order := Order{Id: "123"}
	assert.Error(t, order.Validate(), "invalid price")
}
func Test_If_I_Get_An_Error_If_Tax_Is_Blank(t *testing.T) {
	order := Order{Id: "123", Price: 10.0}
	assert.Error(t, order.Validate(), "invalid Tax")
}
func Test_WithAllValidParams(t *testing.T) {
	order := Order{Id: "123", Price: 10.0, Tax: 1.0}
	assert.NoError(t, order.Validate())
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 1.0, order.Tax)
	order.CalculateFinalPrice()
	assert.Equal(t, 11.0, order.FinalPrice)
}