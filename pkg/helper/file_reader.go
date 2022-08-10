package helper

import (
	"challenge2019/pkg/cmd"
	"challenge2019/pkg/model"
	"encoding/csv"
	"log"
	"os"
	"strings"
)

func ReadFile(fileName string, numID int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Println("Cannot open CSV file:", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	switch numID {

	case 1:
		_, err = reader.Read()
		if err != nil {
			return
		}

		rows, err := reader.ReadAll()
		if err != nil {
			log.Println("Cannot read CSV file:", err)
		}

		for _, row := range rows {
			partner := model.PartnerCSV{
				TheatreID: strings.TrimSpace(row[0]),
				SizeSlab:  strings.TrimSpace(row[1]),
				MinCost:   strings.TrimSpace(row[2]),
				CostPerGB: strings.TrimSpace(row[3]),
				PartnerID: strings.TrimSpace(row[4]),
			}

			cmd.SetPartnersCollectionElement(partner)
		}

	case 2:
		rows, err := reader.ReadAll()
		if err != nil {
			log.Println("Cannot read CSV file:", err)
		}

		for _, row := range rows {
			input := model.InputCSV{
				DeliveryID: strings.TrimSpace(row[0]),
				GB:         strings.TrimSpace(row[1]),
				Theatre:    strings.TrimSpace(row[2]),
			}

			cmd.SetInputCollectionElement(input)
		}
	}

}
