package model

import (
	jenisPegawai "svc-insani-go/modules/v1/master-jenis-pegawai/model"
	statusPegawai "svc-insani-go/modules/v1/master-status-pegawai/model"
)

type KelompokPegawai struct {
	KdKelompokPegawai string `json:"kd_kelompok_pegawai"`
	KelompokPegawai   string `json:"kelompok_pegawai"`
	FlagAktif         int    `json:"-"`
	UUID              string `json:"uuid"`
	ID                uint64 `json:"-" gorm:"primaryKey"`

	KdJenisPegawai string                    `json:"kd_jenis_pegawai"`
	JenisPegawai   jenisPegawai.JenisPegawai `json:"-" gorm:"foreignKey:KdJenisPegawai;references:KDJenisPegawai"`

	KdStatusPegawai string                      `json:"kd_status_pegawai"`
	StatusPegawai   statusPegawai.StatusPegawai `json:"-" gorm:"foreignKey:KdStatusPegawai;references:KDStatusPegawai"`
}

type KelompokPegawaiResponse struct {
	Data []KelompokPegawai `json:"data"`
}
