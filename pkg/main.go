package main

import (
	c "challenge2019/pkg/cmd"
	h "challenge2019/pkg/helper"
	"challenge2019/pkg/model"
)

func main() {
	h.ReadFile("partners.csv", 1)
	h.ReadFile("input.csv", 2)

	var delivery model.Delivery

	for _, val := range c.GetInputCollectionElements() {
		delivery = model.Delivery{
			ID:                           val.DeliveryID,
			IndicationDeliveryISPossible: c.CheckDeliveryISPossible(val.Theatre, val.GB),
			Partner:                      c.FindPartner(),
			Cost:                         c.GetOptimalCost(),
		}

		c.SetOutPutCollectionElement(delivery)
	}

	h.WriteFile("output.csv")
}
