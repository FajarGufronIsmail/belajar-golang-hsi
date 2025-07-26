package mahasiswa

func hitungRataRata(nilai ...int) float64 {
	jumlahNilai := 0
	for _, nilaiMhs := range nilai {
		jumlahNilai += nilaiMhs
	}

	return float64(jumlahNilai) / float64(len(nilai))
}

func BuatMahasiswa(nama string, umur int, nilai ...int) Mahasiswa {
	mhs := Mahasiswa{
		Nama:     nama,
		umur:     umur,
		Nilai:    nilai,
		nilaiAvg: hitungRataRata(nilai...),
	}
	return mhs

}

func SumAge(umur ...int) int {
	total := 0
	for _, u := range umur {
		total += u
	}

	return total
}
