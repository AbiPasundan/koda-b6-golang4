package main

import (
	"fmt"
	. "golang4/lib"
	"os"
)

func main() {
	defer fmt.Println("0. Exit")

	var number int
	fmt.Println("\nmasukan dulu angka")
	fmt.Scanln(&number)

	fmt.Println("Pilih Konversi")
	ShowMenu()

	for true {
		var input int
		fmt.Scanln(&input)

		switch input {
		case 1:
			// ( °C × 9/5) + 32 = 32 °F
			// 900°C * (9/5) + 32 = 1652°F
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
			os.Exit(0)
		default:
			// os.Exit(0)
			fmt.Println("Pilihan tidak ada pilih yang benar")
		}
	}

}
