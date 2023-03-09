package model

type SubBidang struct {
	KdBidang    string `json:"kd_bidang"`
	Bidang      string `json:"bidang"`
	KdSubBidang string `json:"kd_sub_bidang"`
	SubBidang   string `json:"sub_bidang"`
	UUID        string `json:"uuid"`
	ID          uint64 `json:"-"`
}

type SubBidangResponse struct {
	Data []SubBidang `json:"data"`
}
