package user

import (
	"fmt"
	"log"
	"net/http"

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

func (u *UserUseCase) ReadFile(r *http.Request) error {
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
		return shared.WriteError(500, "error when reading excel")
	}

	defer excel.Close()

	sheets := excel.GetSheetList()

	rows, err := excel.GetRows(sheets[0])

	if err != nil {
		return shared.WriteError(500, "error when getting sheets")
	}

	var users []UserImport

	for i, row := range rows {
		if i == 0 {
			continue
		}

		if len(row) >= 3 {
			users = append(users, UserImport{
				Name:        row[0],
				Email:       row[1],
				PhoneNumber: row[2],
			})
		}
	}

	log.Printf("all users : %v", users)

	err = u.rmq.Publish(shared.RoutingKeyUserDirectImport, users)

	if err != nil {
		return shared.WriteError(500, err.Error())
	}

	return nil
}
