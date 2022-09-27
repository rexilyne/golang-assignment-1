package main

import (
	"fmt"
	"os"
	"strconv"
)

type Teman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

// var (
// 	MapTeman = map[int]Teman{}
// )

var MapTeman map[int]Teman = map[int]Teman{}

func main() {
	in := os.Args[1]
	getTemanByKey(in)
}

func init() {
	MapTeman[0] = Teman{
		Nama:      "Brian",
		Alamat:    "Sleman",
		Pekerjaan: "Programmer",
		Alasan:    "Mencari teman baru",
	}

	MapTeman[1] = Teman{
		Nama:      "Iddo",
		Alamat:    "Bekasi",
		Pekerjaan: "IT Support",
		Alasan:    "Gabut",
	}

	MapTeman[2] = Teman{
		Nama:      "Agus",
		Alamat:    "Bali",
		Pekerjaan: "Machine Learning Developer",
		Alasan:    "Menambah ilmu",
	}

	MapTeman[3] = Teman{
		Nama:      "Rico",
		Alamat:    "Yogyakarta",
		Pekerjaan: "Mahasiswa",
		Alasan:    "Ingin nilai A",
	}
}

func getTemanByKey(in string) {
	num, err := strconv.Atoi(in)

	if err != nil {
		fmt.Println("ID harus integer")
		return
	}

	if num < 0 {
		fmt.Printf("ID harus positif")
		return
	}

	if _, ok := MapTeman[num]; !ok {
		fmt.Printf("Student dengan ID %d tidak ada pada database", num)
		return
	}

	fmt.Printf("Nama : %s\n", MapTeman[num].Nama)
	fmt.Printf("Alamat : %s\n", MapTeman[num].Alamat)
	fmt.Printf("Pekerjaan : %s\n", MapTeman[num].Pekerjaan)
	fmt.Printf("Alasan : %s\n", MapTeman[num].Alasan)
}
