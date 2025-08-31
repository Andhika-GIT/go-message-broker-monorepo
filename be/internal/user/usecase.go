package user

import (
	"log"
	"mime/multipart"

	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
	"github.com/xuri/excelize/v2"
)

type UserUseCase struct {
	rmq *shared.RabbitMqProducer
}

func NewUserUseCase(rmq *shared.RabbitMqProducer) *UserUseCase {
	return &UserUseCase{
		rmq: rmq,
	}
}

func (u *UserUseCase) ReadFile(file multipart.File) error {
	excel, err := excelize.OpenReader(file)

	log.Printf("excel: %+v", excel)

	if err != nil {
		return shared.WriteError(500, "error when reading excel")
	}

	defer excel.Close()

	sheets := excel.GetSheetList()

	rows, err := excel.GetRows(sheets[0])

	if err != nil {
		return shared.WriteError(500, "error when getting sheets")
	}

	for i, row := range rows {
		log.Printf("Row %d: %v\n", i+1, row)
	}

	return nil
}
