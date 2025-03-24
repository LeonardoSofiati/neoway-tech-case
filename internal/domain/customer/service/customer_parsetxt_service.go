package service

import (
	"bufio"
	"errors"
	"io"
	"neoway_test/internal/domain/customer/dto"

	// usecase "neoway_test/internal/usecase/customer/create"

	"strconv"
	"strings"
	"time"
)

type ParseTxtFileService struct{}
type ParseService struct{}

// NewParseTxtFileService cria uma nova instância do serviço.
func NewParseTxtFileService() *ParseTxtFileService {
	return &ParseTxtFileService{}
}

func NewParseService() *ParseService {
	return &ParseService{}
}

func (s *ParseTxtFileService) ExecuteParseTxtFileService(file io.Reader) ([]dto.OutputCreateCustomerDto, error) {
	var customers []dto.OutputCreateCustomerDto
	reader := bufio.NewScanner(file)

	lineIndex := 0
	for reader.Scan() {
		line := reader.Text()
		// Skip header
		if lineIndex == 0 {
			lineIndex++
			continue
		}

		if len(line) < 135 {
			return nil, errors.New("invalid file format: line too short")
		}

		customer := dto.OutputCreateCustomerDto{
			Cpf:                parseNull(strings.TrimSpace(line[0:19])),
			Private:            strings.TrimSpace(line[19:31]),
			Incompleto:         strings.TrimSpace(line[31:43]),
			DataUltimaCompra:   parseDate(strings.TrimSpace(line[43:65])),
			TicketMedio:        parseFloat(strings.TrimSpace(line[65:87])),
			TicketUltimaCompra: parseFloat(strings.TrimSpace(line[87:111])),
			LojaMaisFrequente:  parseNull(strings.TrimSpace(line[111:131])),
			LojaUltimaCompra:   parseNull(strings.TrimSpace(line[131:])),
		}

		customers = append(customers, customer)
		lineIndex++
	}

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}

func (s *ParseService) ExecuteParseService(input dto.InputCreateCustomerDto) (dto.OutputCreateCustomerDto, error) {

	customer := dto.OutputCreateCustomerDto{
		Cpf:                parseNull(input.Cpf),
		Private:            strings.TrimSpace(input.Private),
		Incompleto:         strings.TrimSpace(input.Incompleto),
		DataUltimaCompra:   parseDate(strings.TrimSpace(input.DataUltimaCompra)),
		TicketMedio:        input.TicketMedio,
		TicketUltimaCompra: input.TicketUltimaCompra,
		LojaMaisFrequente:  parseNull(strings.TrimSpace(input.LojaMaisFrequente)),
		LojaUltimaCompra:   parseNull(strings.TrimSpace(input.LojaUltimaCompra)),
	}

	return customer, nil
}

// parseDate converts a string date to *time.Time (or nil if "NULL")
func parseDate(value string) *time.Time {
	if strings.ToUpper(value) == "NULL" || value == "" {
		return nil
	}
	t, err := time.Parse("2006-01-02", value)
	if err != nil {
		return nil
	}
	return &t
}

// parseFloat converts a string to float64 (or 0 if "NULL")
func parseFloat(value string) float64 {
	if strings.ToUpper(value) == "NULL" || value == "" {
		return 0
	}
	f, err := strconv.ParseFloat(strings.ReplaceAll(value, ",", "."), 64)
	if err != nil {
		return 0
	}
	return f
}

// parseNull converts "NULL" to upper
func parseNull(value string) string {
	return strings.ToUpper(value)
}
