package model

type Bidang struct {
	KdBidang string `json:"kd_bidang"`
	Bidang   string `json:"bidang"`
	UUID     string `json:"uuid"`
	ID       uint64 `json:"-"`
}

type BidangResponse struct {
	Data []Bidang `json:"data"`
}
