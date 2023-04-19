package main

import (
	"Assigment1/models"
	"fmt"
	"os"
)

func main() {
	// Dummy Data
	dataSiswa := []models.Siswa{
		{ID: 1, Nama: "Aldi", Alamat: "Jakarta", Pekerjaan: "Karyawan", Alasan: "Karena ingin meningkatkan skill programming"},
		{ID: 2, Nama: "Tashya", Alamat: "Jakarta", Pekerjaan: "Mahasiswa", Alasan: "Karena ingin membuat aplikasi yang lebih efisien"},
		{ID: 3, Nama: "Ihza", Alamat: "Jakarta", Pekerjaan: "Programmer", Alasan: "Karena ingin mengembangkan aplikasi berperforma tinggi\""},
		{ID: 4, Nama: "Syarif", Alamat: "Jakarta", Pekerjaan: "Karyawan", Alasan: "Karena ingin mempelajari bahasa pemrograman yang baru"},
		{ID: 5, Nama: "Valdo", Alamat: "Jakarta", Pekerjaan: "Karyawan", Alasan: "Karena ingin meningkatkan kesempatan karir di bidang IT"},
		{ID: 6, Nama: "Husin", Alamat: "Jakarta", Pekerjaan: "Backend Developer", Alasan: "Karena tertarik dengan fitur-fitur bahasa Go"},
		{ID: 7, Nama: "Aries", Alamat: "Bekasi", Pekerjaan: "Karyawan Swasta", Alasan: "Karena Go merupakan bahasa pemrograman yang mudah dipelajari"},
	}

	getByName := models.SearchData(dataSiswa)
	if len(os.Args) != 2 {
		fmt.Println("Masukkan Hanya 1 Parameter")
		return
	}
	result := getByName(os.Args[1])
	if result.ID == 0 {
		fmt.Println("Data Tidak Ada")
		return
	}
	fmt.Printf("ID : %d\nNama : %s\nAlamat : %s\nPekerjaan : %s\nAlasan : %s", result.ID, result.Nama, result.Alamat, result.Pekerjaan, result.Alasan)
}
