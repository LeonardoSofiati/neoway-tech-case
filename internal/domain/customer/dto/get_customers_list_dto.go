package dto

import "time"

type InputGetCustomersListDto struct {
	Page int
}

type OutputGetCustomersListDto struct {
	ID                          string     `json:"id"`
	Cpf                         string     `json:"cpf"`
	CpfValido                   bool       `json:"cpf_valido"`
	Private                     string     `json:"private"`
	Incompleto                  string     `json:"incompleto"`
	DataUltimaCompra            *time.Time `json:"data_ultima_compra"`
	TicketMedio                 float64    `json:"ticket_medio"`
	TicketUltimaCompra          float64    `json:"ticket_ultima_compra"`
	LojaMaisFrequente           string     `json:"loja_mais_frequente"`
	CnpjLojaMaisFrequenteValido bool       `json:"cnpj_loja_mais_frequente_valido"`
	LojaUltimaCompra            string     `json:"loja_ultima_compra"`
	CnpjLojaUltimaCompraValido  bool       `json:"cnpj_loja_ultima_compra_valido"`
	CreatedAt                   time.Time  `json:"created_at"`
}
