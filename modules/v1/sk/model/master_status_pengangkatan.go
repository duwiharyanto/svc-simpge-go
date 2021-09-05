package model

type StatusPengangkatan struct {
	KdStatusPengangkatan string `json:"kd_status_pengangkatan"`
	StatusPengangkatan   string `json:"status_pengangkatan"`
	UUID                 string `json:"uuid"`
	ID                   uint64 `json:"-"`
}

type StatusPengangkatanResponse struct {
	Data []StatusPengangkatan `json:"data"`
}
