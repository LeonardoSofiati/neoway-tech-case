package dto

import "time"

type InputCreateCustomerDto struct {
	Cpf                string
	Private            string
	Incompleto         string
	DataUltimaCompra   string
	TicketMedio        float64
	TicketUltimaCompra float64
	LojaMaisFrequente  string
	LojaUltimaCompra   string
}

type OutputCreateCustomerDto struct {
	ID                          string
	Cpf                         string
	CpfValido                   bool
	Private                     string
	Incompleto                  string
	DataUltimaCompra            *time.Time
	TicketMedio                 float64
	TicketUltimaCompra          float64
	LojaMaisFrequente           string
	CnpjLojaMaisFrequenteValido bool
	LojaUltimaCompra            string
	CnpjLojaUltimaCompraValido  bool
	CreatedAt                   time.Time
}
