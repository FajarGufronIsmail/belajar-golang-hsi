package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	validErrName = errors.New("Isi nama Anda!")
	validErrAge  = errors.New("Umur tidak valid (minimal 18 tahun)")
)

func getData(label string) string {
	fmt.Print(label)
	r := bufio.NewReader(os.Stdin)
	s, _ := r.ReadString('\n')
	return strings.TrimSpace(s)
}

func validateData(name string, age int) error {
	if name == "" {
		return validErrName
	} else if age < 18 {
		return validErrAge
	} else {
		return nil
	}
}

func main() {
	name := getData("Nama: ")
	age, convErr := strconv.Atoi(getData("Umur: "))

	err := validateData(name, age)
	if err != nil {
		if errors.Is(err, validErrName) {
			fmt.Println(validErrName)
		} else if errors.Is(err, validErrAge) {
			fmt.Println(validErrAge)
		}
	}

	if convErr != nil {
		logrus.WithError(convErr).Error("Input umur tidak valid")
	}

	if err != nil {
		logrus.WithError(err).Error("Data invalid")
	} else {
		fmt.Printf("Selamat datang, %s!", name)
	}

}
