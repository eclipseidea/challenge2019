package helper

import (
	c "challenge2019/pkg/cmd"
	"encoding/csv"
	"log"
	"os"
)

func WriteFile(fileName string) {
	file, err := os.Create(fileName)
	defer file.Close()

	w := csv.NewWriter(file)
	if err != nil {
		log.Println("Cannot create CSV file:", err)
	}

	for _, val := range c.GetOutPutCollectionElements() {
		if err = w.Write([]string{val.ID, val.IndicationDeliveryISPossible, val.Partner, val.Cost}); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	w.Flush()

	if err = w.Error(); err != nil {
		log.Fatal(err)
	}
}
