package models

import (
	"strings"
)

type DataSiswa struct {
	ID        int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

var Name = []DataSiswa{
	{ID: 1, Nama: "Aldi", Alamat: "Jakarta", Pekerjaan: "Karyawan", Alasan: "Karena ingin meningkatkan skill programming"},
	{ID: 2, Nama: "Tashya", Alamat: "Jakarta", Pekerjaan: "Mahasiswa", Alasan: "Karena ingin membuat aplikasi yang lebih efisien"},
	{ID: 3, Nama: "Ihza", Alamat: "Jakarta", Pekerjaan: "Programmer", Alasan: "Karena ingin mengembangkan aplikasi berperforma tinggi\""},
	{ID: 4, Nama: "Syarif", Alamat: "Jakarta", Pekerjaan: "Karyawan", Alasan: "Karena ingin mempelajari bahasa pemrograman yang baru"},
	{ID: 5, Nama: "Valdo", Alamat: "Jakarta", Pekerjaan: "Karyawan", Alasan: "Karena ingin meningkatkan kesempatan karir di bidang IT"},
	{ID: 6, Nama: "Husin", Alamat: "Jakarta", Pekerjaan: "Wiraswasta", Alasan: "Karena tertarik dengan fitur-fitur bahasa Go"},
	{ID: 7, Nama: "Aries", Alamat: "Jakarta", Pekerjaan: "Wiraswasta", Alasan: "Karena Go merupakan bahasa pemrograman yang mudah dipelajari"},
}

func SearchData(dataSiswa []DataSiswa) func(string) DataSiswa {
	return func(nama string) DataSiswa {
		for _, v := range dataSiswa {
			if strings.ToLower(v.Nama) == strings.ToLower(nama) {
				return v
			}
		}
		return DataSiswa{}
	}
}
