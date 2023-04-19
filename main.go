package main

import (
	"Assigment1/models"
	"fmt"
	"os"
)

func main() {
	var namaSearch = models.SearchData(models.Name)
	if len(os.Args) != 2 {
		fmt.Println("Masukkan Hanya 1 Parameter")
		return
	}
	result := namaSearch(os.Args[1])
	if result.ID == 0 {
		fmt.Println("Data Tidak Ada")
		return
	}
	fmt.Printf("ID : %d\nNama : %s\nAlamat : %s\nPekerjaan : %s\nAlasan : %s", result.ID, result.Nama, result.Alamat, result.Pekerjaan, result.Alasan)
}
