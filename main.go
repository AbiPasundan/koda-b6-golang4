package main

import (
	"fmt"
	. "golang4/lib"
	"os"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	var number int
	fmt.Println("\nmasukan dulu angka")
	_, err := fmt.Scanln(&number)
	if err != nil {
		panic(err)
	}

	fmt.Println("Pilih Konversi")
	ShowMenu()

	for true {
		var input int
		fmt.Scanln(&input)

		switch input {
		case 1:
			// ( °C × 9/5) + 32 = 32 °F
			result := number*(9/5) + 32
			fmt.Print("\nHasilnya ", result)
			fmt.Printf("\n\nJika ingin konversi lagi silahkan pilih jika tidak masukan angka 0 \n")
			main()
		case 2:
			// °C * (4/5) = 8°Ré
			result := number * 4 / 5
			fmt.Print("Hasilnya ", result)
			fmt.Printf("\n\nJika ingin konversi lagi silahkan pilih jika tidak masukan angka 0 \n")
			main()
		case 3:
			// °C + 273.15 = 293.15 K
			result := number + 273
			fmt.Print("Hasilnya ", result)
			fmt.Printf("\n\nJika ingin konversi lagi silahkan pilih jika tidak masukan angka 0 \n")
			main()
		case 0:
			fmt.Println("Bay bay samapai jumpah")
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak ada pilih yang benar")
		}
	}

}
