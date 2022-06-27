package main

import (
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {

	var persons = []Person{
		{Nama: "Andi", Alamat: "Jakarta", Pekerjaan: "Pegawai Negeri", Alasan: "Training"},
		{Nama: "Budi", Alamat: "Tangerang", Pekerjaan: "Wiraswasta", Alasan: "Tantangan baru"},
		{Nama: "Carl", Alamat: "Bogor", Pekerjaan: "Software Enginner", Alasan: "Pekerjaan"},
		{Nama: "Delon", Alamat: "Bekasi", Pekerjaan: "IT Support", Alasan: "Untuk peningkatan karir"},
	}
	args := os.Args

	if len(args) <= 1 {
		fmt.Println("Invalid Parameter")
	} else {
		arg, error := strconv.Atoi(args[1])
		if error != nil {
			fmt.Println(error)
		} else if arg <= len(persons)-1 {
			fmt.Println(persons[arg])
		} else {
			fmt.Println("Parameter Out of Range")
		}
	}

}
