package models

import (
	"strings"
)

type Siswa struct {
	ID        int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func SearchData(dataSiswa []Siswa) func(string) Siswa {
	return func(nama string) Siswa {
		for _, v := range dataSiswa {
			if strings.EqualFold(v.Nama, nama) {
				return v
			}
		}
		return Siswa{}
	}
}
