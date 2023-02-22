package model

type BidangSubBidang struct {
	KdBidangSubBidang string `json:"kd_bidang_sub_bindang"`
	BidangSubBidang   string `json:"bidang_sub_bidang"`
	UUID              string `json:"uuid"`
	ID                uint64 `json:"-"`
}

type BidangSubBidangResponse struct {
	Data []BidangSubBidang `json:"data"`
}
