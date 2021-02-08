package model

type UnitKerja struct {
	ID            string `json:"-"`
	KdUnitKerja   string `json:"kd_unit_kerja"`
	NamaUnitKerja string `json:"nama_unit_kerja"`
	UUID          string `json:"uuid"`
}

type UnitKerjaResponse struct {
	Data []UnitKerja `json:"data"`
}
