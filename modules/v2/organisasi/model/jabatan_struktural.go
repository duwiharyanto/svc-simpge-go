package model

type JabatanStruktural struct {
	Id                   uint64 `form:"-" json:"-"`
	IdStrukturOrganisasi uint64 `form:"-" json:"-"`
	IdJenisJabatan       uint64 `form:"-" json:"-"`
	NamaJenisJabatan     string `form:"-" json:"jenis_jabatan"`
	IdJenisUnit          uint64 `form:"-" json:"-"`
	NamaJenisUnit        string `form:"-" json:"jenis_unit"`
	IdUnit               uint64 `form:"-" json:"-"`
	NamaUnit             string `form:"-" json:"unit"`
	KdUnit               string `form:"-" json:"kd_unit"`
	FlagAktif            int    `form:"-" json:"-"`
	Uuid                 string `form:"-" json:"uuid"`
}

type PejabatStruktural struct {
	Id            uint64 `form:"-" json:"-"`
	Nama          string `form:"-" json:"nama"`
	GelarDepan    string `form:"-" json:"gelar_depan"`
	GelarBelakang string `form:"-" json:"gelar_belakang"`
	FlagAktif     int    `form:"-" json:"-"`
	Uuid          string `form:"-" json:"uuid"`

	IdStrukturOrganisasi uint64            `form:"-" json:"-"`
	JabatanStruktural    JabatanStruktural `form:"-" json:"-" gorm:"foreignKey:IdStrukturOrganisasi;references:IdStrukturOrganisasi"`
}
