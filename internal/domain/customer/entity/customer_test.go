package entity

import (
	internalerrors "neoway_test/internal/internal-errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewCustomer(t *testing.T) {
	dataUltimaCompra := time.Date(2011, 1, 27, 0, 0, 0, 0, time.UTC)
	customer, err := NewCustomer(
		"922.488.109-20",
		"0",
		"0",
		&dataUltimaCompra,
		130.54,
		130.54,
		"79.379.491/0001-83",
		"79.379.491/0001-83",
	)

	assert.Nil(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, "922.488.109-20", customer.Cpf)
	assert.Equal(t, "0", customer.Private)
	assert.Equal(t, "0", customer.Incompleto)
	assert.True(t, customer.CpfValido)
	assert.Equal(t, "79.379.491/0001-83", customer.LojaMaisFrequente)
	assert.Equal(t, "79.379.491/0001-83", customer.LojaUltimaCompra)
	assert.Equal(t, dataUltimaCompra, *customer.DataUltimaCompra)
	assert.Equal(t, 130.54, customer.TicketMedio)
	assert.Equal(t, 130.54, customer.TicketUltimaCompra)
	assert.True(t, customer.CnpjLojaMaisFrequenteValido)
	assert.True(t, customer.CnpjLojaUltimaCompraValido)
}

func TestValidateCpf(t *testing.T) {
	assert.True(t, validateCpf("922.488.109-20"))
	assert.False(t, validateCpf("123.456.789-00"))
}

func TestValidateCnpj(t *testing.T) {
	assert.True(t, validateCnpj("79.379.491/0001-83"))
	assert.False(t, validateCnpj("12.312.312/3123-12"))
}

func TestCustomerWithNullDataUltimaCompra(t *testing.T) {
	customer, err := NewCustomer(
		"041.091.641-25",
		"0",
		"0",
		nil,
		50.00,
		50.00,
		"79.379.491/0001-83",
		"79.379.491/0001-83",
	)
	assert.Nil(t, err)
	assert.NotNil(t, customer)
	assert.Nil(t, customer.DataUltimaCompra)
}

func TestInvalidCustomerCnpjValidation(t *testing.T) {
	customer, err := NewCustomer(
		"922.488.109-20",
		"0",
		"0",
		nil,
		100.00,
		100.00,
		"InvãlidLójaCNPJ",
		"79.379.491/0001-83",
	)
	assert.Nil(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, customer.LojaMaisFrequente, "INVALIDLOJACNPJ")
	assert.Equal(t, customer.LojaUltimaCompra, "79.379.491/0001-83")
}

func TestEmptyCustomerFields(t *testing.T) {
	customer, err := NewCustomer(
		"", "", "", nil, 0.0, 0.0, "", "",
	)
	assert.Nil(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, "NULL", customer.Cpf)
	assert.False(t, customer.CpfValido)
	assert.Equal(t, "NULL", customer.LojaMaisFrequente)
	assert.Equal(t, "NULL", customer.LojaUltimaCompra)
	assert.Nil(t, customer.DataUltimaCompra)
	assert.Equal(t, 0.0, customer.TicketMedio)
	assert.Equal(t, 0.0, customer.TicketUltimaCompra)
	assert.False(t, customer.CnpjLojaMaisFrequenteValido)
	assert.False(t, customer.CnpjLojaUltimaCompraValido)
}

func TestCustomer_Validate_ValidCPF(t *testing.T) {
	dataUltimaCompra := time.Now()
	customer := &Customer{
		Cpf:                         "922.488.109-20",
		CpfValido:                   true,
		Private:                     "1",
		Incompleto:                  "0",
		DataUltimaCompra:            &dataUltimaCompra,
		TicketMedio:                 130.54,
		TicketUltimaCompra:          130.54,
		LojaMaisFrequente:           "79.379.491/0001-83",
		CnpjLojaMaisFrequenteValido: true,
		LojaUltimaCompra:            "79.379.491/0001-83",
		CnpjLojaUltimaCompraValido:  true,
	}

	err := internalerrors.ValidateStruct(customer)
	assert.Nil(t, err)
}

func TestSanitizeInput(t *testing.T) {
	assert.Equal(t, "NULL", sanitizeInput(""))
	assert.Equal(t, "922.488.109-20", sanitizeInput("922.488.109-20"))
	assert.Equal(t, "LOJA MAIS FREQUENTE", sanitizeInput("Loja Mais Freqüente"))
	assert.Equal(t, "INCOMPLETO", sanitizeInput("Incompleto"))
	assert.Equal(t, "0", sanitizeInput("0"))
	assert.Equal(t, "NULL", sanitizeInput("NULL"))
	assert.Equal(t, "PRIVATE", sanitizeInput("Private"))
}
