package order

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
	"github.com/xuri/excelize/v2"
)

type OrderUseCase struct {
	rmq *shared.RabbitMqProducer
}

func NewOrderUseCase(rmq *shared.RabbitMqProducer) *OrderUseCase {
	return &OrderUseCase{
		rmq: rmq,
	}
}

func (u *OrderUseCase) ReadFile(r *http.Request) error {
	file, header, err := r.FormFile("file")

	if err != nil {
		return shared.WriteError(500, fmt.Sprintf("failed to read file %s", err.Error()))
	}

	defer file.Close()

	isFileExtensionCorrect := shared.IsAllowedExtension(header.Filename)

	if !isFileExtensionCorrect {
		return shared.WriteError(400, "invalid file extension")
	}

	excel, err := excelize.OpenReader(file)

	if err != nil {
		return shared.WriteError(500, fmt.Sprintf("error when reading excel: %s", err.Error()))
	}

	defer excel.Close()

	sheets := excel.GetSheetList()

	rows, err := excel.GetRows(sheets[0])

	if err != nil {
		return shared.WriteError(500, fmt.Sprintf("error when getting excel sheets: %s", err.Error()))
	}

	var orders []OrderImport

	for i, row := range rows {
		if i == 0 {
			continue
		}

		if len(row) >= 3 {
			quantity, err := strconv.Atoi(row[2])
			if err != nil {
				continue
			}

			orders = append(orders, OrderImport{
				Email:       row[0],
				ProductName: row[1],
				Quantity:    quantity,
			})
		}
	}

	log.Printf("all users : %v", orders)

	err = u.rmq.Publish(shared.QueueOrderDirectImport, orders)

	if err != nil {
		return shared.WriteError(500, err.Error())
	}

	return nil

}
