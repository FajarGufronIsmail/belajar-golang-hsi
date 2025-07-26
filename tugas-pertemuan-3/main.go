package main

import (
	"fmt"
	"tugas-pertemuan-3/mahasiswa"
)

func main() {
	//Membuat objek menggunakan fungsi BuatMahasiswa
	fajar := mahasiswa.BuatMahasiswa("Fajar", 26, 80, 90, 95)
	ismail := mahasiswa.BuatMahasiswa("Ismail", 24, 85, 95, 88)

	mhs := []mahasiswa.Mahasiswa{fajar, ismail} //Membuat slice data Mahasiswa agar ditampilkan berupa list

	for _, mahas := range mhs { //for each untuk menampilkan semua list data mahasiswa
		fmt.Println(mahas.Info())
		fmt.Printf("Rata-rata nilai: %.2f \n--- \n", mahas.RataRata())
	}
	//hardcode untuk versi package dan nilai maksimum
	fmt.Printf("Versi Package: %s \n", mahasiswa.Version)
	fmt.Printf("Nilai Maksimum: %d \n", mahasiswa.GetMaxNilai())
	sum := []int{fajar.GetUmur(), ismail.GetUmur()} //Membuat slice untuk menghitung total umur mahasiswa
	fmt.Printf(
		"Total Umur Mahasiswa: %d \n", mahasiswa.SumAge(sum...),
	)

}
