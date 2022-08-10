package cmd

import (
	"challenge2019/pkg/model"
)

var input []model.InputCSV

func SetInputCollectionElement(model model.InputCSV) {
	input = append(input, model)
}

func GetInputCollectionElements() []model.InputCSV {
	return input
}
