package model

type UnitPengangkat struct {
	KdUnitPengangkat string `json:"kd_unit_pengangkat"`
	UnitPengangkat   string `json:"unit_pengangkat"`
	UUID             string `json:"uuid"`
	ID               uint64 `json:"-"`
}

type UnitPengangkatResponse struct {
	Data []UnitPengangkat `json:"data"`
}
