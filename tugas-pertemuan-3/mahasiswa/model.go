package mahasiswa

import (
	"fmt"
)

var Version string = "v1.0.0"
var maxNilai int = 100

type Mahasiswa struct {
	Nama     string
	Nilai    []int
	umur     int
	nilaiAvg float64
}

type Derkripsi interface {
	Info() string
	RataRata() float64
	GetUmur() int
}

func (mahasiswa *Mahasiswa) Info() string {
	return fmt.Sprintf("Nama: %s, Umur: %d", mahasiswa.Nama, mahasiswa.umur)
}

func (mahasiswa *Mahasiswa) RataRata() float64 {
	return mahasiswa.nilaiAvg
}

func (mahasiswa *Mahasiswa) GetUmur() int {
	return mahasiswa.umur
}

func GetMaxNilai() int {
	return maxNilai
}
