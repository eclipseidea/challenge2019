package model

type PartnerCSV struct {
	TheatreID string
	SizeSlab  string
	MinCost   string
	CostPerGB string
	PartnerID string
}

type PartnerData struct {
	ID        string
	MinCost   int
	CostPerGB int
	TotalCost int
}
