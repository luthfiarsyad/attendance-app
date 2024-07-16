package handler

import (
	"attendance-app/internal/usecase"
	"fmt"
)

type AttendanceHandler struct {
	usecase usecase.AttendanceUsecase
}

func NewAttendanceHandler(usecase usecase.AttendanceUsecase) *AttendanceHandler {
	return &AttendanceHandler{usecase: usecase}
}

func (handler *AttendanceHandler) TambahKaryawan(id int, nama string) {
	handler.usecase.TambahKaryawan(id, nama)
	fmt.Println("Karyawan berhasil ditambahkan")
}

func (handler *AttendanceHandler) UpdateKehadiran(id int, hadir bool) {
	handler.usecase.UpdateKehadiran(id, hadir)
	fmt.Println("Status kehadiran berhasil diperbarui")
}

func (handler *AttendanceHandler) HapusKaryawan(id int) {
	handler.usecase.HapusKaryawan(id)
	fmt.Println("Karyawan berhasil dihapus")
}

func (handler *AttendanceHandler) TampilkanKehadiran() {
	daftarKaryawan := handler.usecase.TampilkanKehadiran()
	for _, karyawan := range daftarKaryawan {
		fmt.Printf("ID: %d, Nama: %s, Kehadiran: %t\n", karyawan.ID, karyawan.Nama, karyawan.Kehadiran)
	}
}
