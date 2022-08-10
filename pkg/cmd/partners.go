package cmd

import (
	"challenge2019/pkg/model"
)

var partners []model.PartnerCSV

func SetPartnersCollectionElement(model model.PartnerCSV) {
	partners = append(partners, model)
}

func GetPartnersCollectionElements() []model.PartnerCSV {
	return partners
}
