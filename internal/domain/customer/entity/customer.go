package entity

import (
	shared "neoway_test/internal/domain/shared/entity"
	internalerrors "neoway_test/internal/internal-errors"
	"strings"
	"time"

	"github.com/klassmann/cpfcnpj"
	"github.com/mozillazg/go-unidecode"
)

type Customer struct {
	shared.BaseEntity
	Cpf                         string     `json:"cpf" gorm:"size:20;not null"`
	CpfValido                   bool       `json:"cpf_valido" gorm:"not null"`
	Private                     string     `json:"private"`
	Incompleto                  string     `json:"incompleto"`
	DataUltimaCompra            *time.Time `json:"data_ultima_compra"`
	TicketMedio                 float64    `json:"ticket_medio" gorm:"type:numeric(10,2)"`
	TicketUltimaCompra          float64    `json:"ticket_ultima_compra" gorm:"type:numeric(10,2)"`
	LojaMaisFrequente           string     `json:"loja_mais_frequente" gorm:"size:20"`
	CnpjLojaMaisFrequenteValido bool       `json:"cnpj_loja_mais_frequente_valido" gorm:"not null"`
	LojaUltimaCompra            string     `json:"loja_ultima_compra" gorm:"size:20"`
	CnpjLojaUltimaCompraValido  bool       `json:"cnpj_loja_ultima_compra_valido" gorm:"not null"`
}

func NewCustomer(
	cpf string,
	private string,
	incompleto string,
	dataUltimaCompra *time.Time,
	ticketMedio float64,
	ticketUltimaCompra float64,
	lojaMaisFrequente string,
	lojaUltimaCompra string,
) (*Customer, error) {

	cpf = sanitizeInput(cpf)
	lojaMaisFrequente = sanitizeInput(lojaMaisFrequente)
	lojaUltimaCompra = sanitizeInput(lojaUltimaCompra)

	customer := &Customer{
		BaseEntity:                  shared.NewBaseEntity(),
		Cpf:                         cpf,
		CpfValido:                   validateCpf(cpf),
		Private:                     private,
		Incompleto:                  incompleto,
		DataUltimaCompra:            dataUltimaCompra,
		TicketMedio:                 ticketMedio,
		TicketUltimaCompra:          ticketUltimaCompra,
		LojaMaisFrequente:           lojaMaisFrequente,
		CnpjLojaMaisFrequenteValido: validateCnpj(lojaMaisFrequente),
		LojaUltimaCompra:            lojaUltimaCompra,
		CnpjLojaUltimaCompraValido:  validateCnpj(lojaUltimaCompra),
	}

	err := internalerrors.ValidateStruct(customer)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func sanitizeInput(input string) string {
	result := strings.ToUpper(unidecode.Unidecode(input))
	if result == "" {
		return "NULL"
	}
	return result
}

func validateCpf(value string) bool {
	cpf := cpfcnpj.NewCPF(value)

	return cpf.IsValid()
}

func validateCnpj(value string) bool {
	cnpj := cpfcnpj.NewCNPJ(value)

	return cnpj.IsValid()
}
