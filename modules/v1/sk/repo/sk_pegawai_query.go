package repo

import (
	"fmt"
)

// sk list

func getAllSKDosenQuery(uuidDosen string) string {
	return fmt.Sprintf(`
	SELECT 
		COALESCE(d.nama_unit_kerja, ''), 
		COALESCE(e.nama_sk,''),
		COALESCE(e.kd_jenis_sk,''),
		COALESCE(a.nomor_sk, ''), 
		COALESCE(a.tmt, ''), c.uuid
	FROM 
		sk_pegawai a
	JOIN 
		pegawai b ON a.id_pegawai = b.id
	JOIN 
		sk_pengangkatan_dosen c ON a.id = c.id_sk_pegawai
	LEFT JOIN 
		unit_kerja d ON c.id_unit_kerja = d.id
	LEFT JOIN 
		jenis_sk e ON c.id_jenis_sk = e.id 
	WHERE 
		b.uuid = %q`,
		uuidDosen,
	)
}

func getAllSKTendikQuery(uuidTendik string) string {
	return fmt.Sprintf(`
	SELECT 
		COALESCE(d.nama_unit_kerja, ''),
		COALESCE(e.nama_sk,''),
		COALESCE(e.kd_jenis_sk,''),
		COALESCE(a.nomor_sk, ''), 
		COALESCE(a.tmt, ''), c.uuid
	FROM 
		sk_pegawai a
	JOIN 
		pegawai b ON a.id_pegawai = b.id
	JOIN 
		sk_pengangkatan_tendik c ON a.id = c.id_sk_pegawai
	LEFT JOIN 
		unit_kerja d ON c.id_unit_pegawai = d.id
	LEFT JOIN 
		jenis_sk e ON c.id_jenis_sk = e.id 
	WHERE 
		b.uuid = %q`,
		uuidTendik,
	)
}

func getAllSkQuery(uuid string) string {
	return fmt.Sprintf(`
	SELECT 
		COALESCE(d.nama_unit_kerja, '') AS unit_kerja, 
		COALESCE(e.nama_sk,'') AS jenis_sk ,
		COALESCE(e.kd_jenis_sk,'') AS kd_jenis_sk,
		COALESCE(e.uuid,'') AS uuid_jenis_sk,
		COALESCE(a.nomor_sk, '') AS nomor_sk, 
		COALESCE(a.tmt, '') AS tmt,
		c.tgl_update AS tgl_update,
		c.uuid AS uuid
	FROM 
		sk_pegawai a
	JOIN 
		pegawai b ON a.id_pegawai = b.id
	JOIN 
		sk_pengangkatan_dosen c ON a.id = c.id_sk_pegawai
	LEFT JOIN 
		unit_kerja d ON c.id_unit_kerja = d.id
	LEFT JOIN 
		jenis_sk e ON c.id_jenis_sk = e.id 
	WHERE 
		b.uuid =  %q AND a.flag_aktif=1
		
	UNION

	SELECT 
		COALESCE(d.nama_unit_kerja, '') AS unit_kerja,
		COALESCE(e.nama_sk,'') AS jenis_sk ,
		COALESCE(e.kd_jenis_sk,'') AS kd_jenis_sk,
		COALESCE(e.uuid,'') AS uuid_jenis_sk,
		COALESCE(a.nomor_sk, '') AS nomor_sk, 
		COALESCE(a.tmt, '') AS tmt,
		c.tgl_update AS tgl_update, 
		c.uuid AS uuid
	FROM 
		sk_pegawai a
	JOIN 
		pegawai b ON a.id_pegawai = b.id
	JOIN 
		sk_pengangkatan_tendik c ON a.id = c.id_sk_pegawai
	LEFT JOIN 
		unit_kerja d ON c.id_unit_pegawai = d.id
	LEFT JOIN 
		jenis_sk e ON c.id_jenis_sk = e.id 
	WHERE 
		b.uuid = %q AND a.flag_aktif=1

	UNION

	SELECT  
		COALESCE (uk.nama_unit_kerja,'') AS unit_kerja,
		COALESCE (js.nama_sk,'') AS jenis_sk,
		COALESCE (js.kd_jenis_sk,'') AS kd_jenis_sk,
		COALESCE (js.uuid,'') AS uuid_jenis_sk,
		sk.nomor_sk AS nomor_sk, 
		COALESCE (sk.tgl_ditetapkan,'') AS tmt, 
		sk.tgl_update AS tgl_update, 
		sk.uuid AS uuid
	FROM 
		sk_penetapan as sk 
	LEFT JOIN
		pegawai as p ON sk.id_pegawai = p.id
	LEFT JOIN
		unit_kerja as uk ON sk.id_unit_kerja = uk.id  
	LEFT JOIN 
		jenis_sk as js ON sk.id_jenis_sk = js.id
	WHERE 
	p.uuid = %q AND sk.flag_aktif = 1
	
	ORDER BY tgl_update DESC`,
		uuid, uuid, uuid,
	)
}
