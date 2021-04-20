package model

type Instansi struct {
	Id           int    `form:"-" gorm:"primaryKey"`
	NamaInstansi string `form:"nama_instansi" json:"nama_instansi"`
	KdInstansi   string `form:"kd_instansi" json:"kd_instansi"`
	Uuid         string `form:"-"`
}
