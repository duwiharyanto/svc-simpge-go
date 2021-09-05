package model

type Unit2 struct {
	ID     uint64 `json:"-"`
	KdUnit string `json:"kd_unit"`
	Unit   string `json:"unit"`
	UUID   string `json:"uuid"`
}

type UnitKerja struct {
	ID            uint64 `json:"-"`
	KdUnitKerja   string `json:"kd_unit_kerja"`
	NamaUnitKerja string `json:"nama_unit_kerja"`
	UUID          string `json:"uuid"`
}

type UnitKerjaResponse struct {
	Data []UnitKerja `json:"data"`
}
