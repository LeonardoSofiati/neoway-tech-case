package service_test

import (
	"bytes"
	"neoway_test/internal/domain/customer/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteParseTxtFileService_ValidFile(t *testing.T) {
	fileContent := `CPF                PRIVATE     INCOMPLETO  DATA DA ÚLTIMA COMPRA TICKET MÉDIO          TICKET DA ÚLTIMA COMPRA LOJA MAIS FREQUÊNTE LOJA DA ÚLTIMA COMPRA
026.987.379-13     0           0           2011-01-20            159,31                159,31                  79.379.491/0001-83  79.379.491/0001-83
041.091.641-25     0           1           NULL                  NULL                  NULL                    NULL                NULL`

	reader := bytes.NewReader([]byte(fileContent))
	service := service.NewParseTxtFileService()

	customers, err := service.ExecuteParseTxtFileService(reader)

	assert.Nil(t, err)
	assert.Len(t, customers, 2)
	assert.Equal(t, "026.987.379-13", customers[0].Cpf)
	assert.Equal(t, "0", customers[0].Private)
	assert.Equal(t, "0", customers[0].Incompleto)
	assert.NotNil(t, customers[0].DataUltimaCompra)
	assert.Equal(t, 159.31, customers[0].TicketMedio)
	assert.Equal(t, 159.31, customers[0].TicketUltimaCompra)
	assert.Equal(t, "79.379.491/0001-83", customers[0].LojaMaisFrequente)
	assert.Equal(t, "79.379.491/0001-83", customers[0].LojaUltimaCompra)

	assert.Equal(t, "041.091.641-25", customers[1].Cpf)
	assert.Equal(t, "0", customers[1].Private)
	assert.Equal(t, "1", customers[1].Incompleto)
	assert.Nil(t, customers[1].DataUltimaCompra)
	assert.Equal(t, 0.0, customers[1].TicketMedio)
	assert.Equal(t, 0.0, customers[1].TicketUltimaCompra)
	assert.Equal(t, "NULL", customers[1].LojaMaisFrequente)
	assert.Equal(t, "NULL", customers[1].LojaUltimaCompra)
}

func TestExecuteParseTxtFileService_InvalidLineLength(t *testing.T) {
	fileContent := `CPF               PRIVATE        INCOMPLETO     DATA_ULTIMA_COMPRA
922.488.109-20   0              0              2011-01-27`

	reader := bytes.NewReader([]byte(fileContent))
	service := service.NewParseTxtFileService()

	customers, err := service.ExecuteParseTxtFileService(reader)

	assert.Nil(t, customers)
	assert.EqualError(t, err, "invalid file format: line too short")
}
