package main

import (
	"fmt"
	"os"
	"strconv"
)

type siswa struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

var db = []siswa{
	{Nama: "Ahmad Faris", Alamat: "Indonesia", Pekerjaan: "Mahasiswa", Alasan: "Ingin Belajar Hal Baru"},
	{Nama: "Raka Rizqllah Pratama Saputra", Alamat: "Indonesia", Pekerjaan: "APIs Engineer", Alasan: "Ingin lebih mendalami golang"},
	{Nama: "I Gede Diva Dwijayana", Alamat: "Indonesia", Pekerjaan: "System Analyst", Alasan: "Ingin mempelajari golang"},
	{Nama: "Muhammad Daffa Haryadi Putra", Alamat: "Indonesia", Pekerjaan: "Back-End Programmer", Alasan: "Ingin Lebih tau lagi tentang golang"},
	{Nama: "M Yoga Irvandi", Alamat: "Indonesia", Pekerjaan: "Database Engineer", Alasan: "Ingin belajar bahasa golang"},
	{Nama: "Wafianda Azhar", Alamat: "Indonesia", Pekerjaan: "Web Developer", Alasan: "Ingin mencari tahi apa itu golang"},
	{Nama: "Muhammad Haiqal Malik", Alamat: "Indonesia", Pekerjaan: "Mobile App Developer", Alasan: "Ingin menjadi expert di bahasa golang"},
	{Nama: "Luqman Fauzi", Alamat: "Indonesia", Pekerjaan: "Fullstack Programmer", Alasan: "Ingin menjadi programmer bahasa golang"},
	{Nama: "Rubenhard Manat Hasiholan Lumbanraja", Alamat: "Indonesia", Pekerjaan: "Data Analyst", Alasan: "Ingin memahami bahasa golang"},
	{Nama: "Denada Ramschie", Alamat: "Indonesia", Pekerjaan: "Network Administrator", Alasan: "Ingin belajar bahasa yang belum pernah dipelajari"},
}

// fungsi untuk mengambil data sesuai inputan
func getData(id string) {
	// konversi string ke int
	var num0, err = strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("Error: Tolong masukkan angka 1-10")
		return
	}

	//Percabangan
	num := num0 - 1
	if num < 0 || num > 9 {
		fmt.Println("Tolong masukkan angka 1-10")
		return
	}

	// Output
	fmt.Println("Nama: ", db[num].Nama)
	fmt.Println("Alagmat: ", db[num].Alamat)
	fmt.Println("Pekerjaan: ", db[num].Pekerjaan)
	fmt.Println("Alasan memilih kelas Golang: ", db[num].Alasan)
	return
}

func main() {
	id := os.Args[1]
	getData(id)
}
// Silahkan input data melalui terminal
// Contoh: go run main.go 1