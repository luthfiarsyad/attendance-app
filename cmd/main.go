package main

import (
	"attendance-app/internal/handler"
	"attendance-app/internal/repository"
	"attendance-app/internal/usecase"
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

func clearScreen() {
	var clear map[string]func() //create a map for storing clear funcs
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear") // Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") // Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["darwin"] = func() { // macOS
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else {
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func main() {
	repo := repository.NewAttendanceRepository()
	usecase := usecase.NewAttendanceUsecase(repo)
	handler := handler.NewAttendanceHandler(usecase)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		clearScreen()

		fmt.Println("Menu:")
		fmt.Println("1. Tambah Karyawan")
		fmt.Println("2. Update Kehadiran")
		fmt.Println("3. Hapus Karyawan")
		fmt.Println("4. Tampilkan Kehadiran")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih menu: ")

		if !scanner.Scan() {
			break
		}

		pilihan := scanner.Text()

		clearScreen()

		switch pilihan {
		case "1":
			fmt.Print("Masukkan ID Karyawan: ")
			if !scanner.Scan() {
				break
			}
			id, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("ID harus berupa angka.")
				continue
			}

			fmt.Print("Masukkan Nama Karyawan: ")
			if !scanner.Scan() {
				break
			}
			nama := scanner.Text()

			handler.TambahKaryawan(id, nama)

		case "2":
			fmt.Print("Masukkan ID Karyawan: ")
			if !scanner.Scan() {
				break
			}
			id, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("ID harus berupa angka.")
				continue
			}

			fmt.Print("Masukkan Kehadiran (true/false): ")
			if !scanner.Scan() {
				break
			}
			hadir, err := strconv.ParseBool(scanner.Text())
			if err != nil {
				fmt.Println("Kehadiran harus berupa true atau false.")
				continue
			}

			handler.UpdateKehadiran(id, hadir)

		case "3":
			fmt.Print("Masukkan ID Karyawan: ")
			if !scanner.Scan() {
				break
			}
			id, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("ID harus berupa angka.")
				continue
			}

			handler.HapusKaryawan(id)

		case "4":
			handler.TampilkanKehadiran()
			fmt.Println("\nTekan Enter untuk kembali ke menu.")
			scanner.Scan()

		case "5":
			fmt.Println("Keluar dari program.")
			return

		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
