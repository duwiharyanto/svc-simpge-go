package model

import "sync"

type Pengaturan struct {
	nilai map[string]string
	mu    sync.Mutex
}

func InitPengaturan() *Pengaturan {
	nm := make(map[string]string)
	return &Pengaturan{nilai: nm}
}

func (p *Pengaturan) Get(atribut string) string {
	return p.nilai[atribut]
}

func (p *Pengaturan) Set(atribut, nilai string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.nilai[atribut] = nilai
}

type PengaturanResponse struct {
	Message string                 `json:"message,omitempty"`
	Data    map[string]interface{} `json:"data"`
}
