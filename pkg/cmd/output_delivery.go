package cmd

import (
	"challenge2019/pkg/model"
	"regexp"
	"strconv"
)

var (
	outPut           []model.Delivery
	filteredPartners []model.PartnerCSV
	optimalPartner   model.PartnerData
)

var (
	trafficGB                  int
	sizeSlabPartnerOpportunity string
	deliveryIsPossible         bool
	optimalPrice               int
)

func SetOutPutCollectionElement(model model.Delivery) {
	outPut = append(outPut, model)
}

func GetOutPutCollectionElements() []model.Delivery {
	return outPut
}

func CheckDeliveryISPossible(theatreID, GB string) string {
	filteredPartners = nil
	_trafficGB, _ := strconv.ParseInt(GB, 0, 0)
	trafficGB = int(_trafficGB)

	for _, val := range GetPartnersCollectionElements() {
		sizeSlabPartnerOpportunity = val.SizeSlab
		if val.TheatreID == theatreID && sizeSlabOpportunity() {
			filteredPartners = append(filteredPartners, val)
		}
	}

	if filteredPartners != nil {
		deliveryIsPossible = true
	} else {
		deliveryIsPossible = false
	}

	return strconv.FormatBool(deliveryIsPossible)
}

func sizeSlabOpportunity() bool {
	a, b := sizeSlabStringConvertor()

	if trafficGB >= a && trafficGB <= b {
		return true
	}

	return false
}

func sizeSlabStringConvertor() (minSize, maxSize int) {
	re := regexp.MustCompile("[0-9]+")
	nums := re.FindAllString(sizeSlabPartnerOpportunity, -1)
	minTrafficOpportunity, _ := strconv.ParseInt(nums[0], 0, 0)
	maxTrafficOpportunity, _ := strconv.ParseInt(nums[1], 0, 0)

	return int(minTrafficOpportunity), int(maxTrafficOpportunity)
}

func FindPartner() string {
	var partnersCollection []model.PartnerData
	var partnerID = " "

	for _, val := range filteredPartners {
		_minCost, _ := strconv.ParseInt(val.MinCost, 0, 0)
		_pricePerGB, _ := strconv.ParseInt(val.CostPerGB, 0, 0)

		partner := model.PartnerData{
			ID:        val.PartnerID,
			MinCost:   int(_minCost),
			CostPerGB: int(_pricePerGB),
			TotalCost: int(_pricePerGB) * trafficGB,
		}

		partnersCollection = append(partnersCollection, partner)
	}

	if partnersCollection != nil && len(partnersCollection) == 1 {
		optimalPartner = partnersCollection[0]

		if optimalPartner.TotalCost <= optimalPartner.MinCost {
			optimalPartner.TotalCost = optimalPartner.MinCost

		} else if optimalPartner.TotalCost > optimalPartner.MinCost {
			optimalPartner.MinCost = optimalPartner.TotalCost
		}

	} else if len(partnersCollection) > 1 {
		optimalPartner = findMinTotalCostValue(partnersCollection)
	}

	if deliveryIsPossible {
		partnerID = optimalPartner.ID
		optimalPrice = optimalPartner.MinCost
	}

	return partnerID
}

func GetOptimalCost() string {
	if deliveryIsPossible {
		return strconv.Itoa(optimalPrice)
	}
	return " "
}

func findMinTotalCostValue(data []model.PartnerData) model.PartnerData {
	var partnersData []model.PartnerData
	var partner model.PartnerData

	for i, _ := range data {
		partner = data[i]
		if partner.TotalCost <= partner.MinCost {
			partner.TotalCost = partner.MinCost
		} else if partner.TotalCost > partner.MinCost {
			partner.MinCost = partner.TotalCost
		}

		partnersData = append(partnersData, partner)
	}

	for _, val := range partnersData {
		if partner.MinCost > val.MinCost {
			partner = val
		}
	}

	return partner
}
