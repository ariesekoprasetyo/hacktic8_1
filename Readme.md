# Program Golang Sederhana untuk Pencarian Data

Ini adalah program Golang sederhana yang mencari data berdasarkan parameter input nama seseorang. Program mencari data dari slice tipe `DataSiswa` yang telah ditentukan sebelumnya di dalam package `models`.

## Deskripsi Kode

### Package Utama (`main`)

Package `main` berisi fungsi `main()` program, yang menerima satu parameter input (nama seseorang) dari argumen baris perintah. Program mencari data berdasarkan nama masukan menggunakan fungsi `SearchData()` dari package `models`. Jika data ditemukan, program menampilkan data tersebut di console.

### Package Models (`models`)

Package `models` berisi struct `DataSiswa` yang merepresentasikan data seseorang. Package ini juga berisi slice tipe `DataSiswa` yang digunakan sebagai database data siswa. Fungsi `SearchData()` adalah fungsi yang digunakan untuk mencari data siswa berdasarkan nama. Fungsi ini mengembalikan data siswa yang sesuai atau struct kosong jika tidak ditemukan.
