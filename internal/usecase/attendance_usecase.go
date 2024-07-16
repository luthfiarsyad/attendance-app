package usecase

import (
	"attendance-app/internal/repository"
	"attendance-app/model"
)

type AttendanceUsecase interface {
	TambahKaryawan(id int, nama string)
	UpdateKehadiran(id int, hadir bool)
	HapusKaryawan(id int)
	TampilkanKehadiran() []*model.Karyawan
}

type AttendanceUsecaseImpl struct {
	repo repository.AttendanceRepository
}

func NewAttendanceUsecase(repo repository.AttendanceRepository) AttendanceUsecase {
	return &AttendanceUsecaseImpl{repo: repo}
}

func (usecase *AttendanceUsecaseImpl) TambahKaryawan(id int, nama string) {
	karyawanBaru := &model.Karyawan{
		ID:        id,
		Nama:      nama,
		Kehadiran: false,
	}
	usecase.repo.TambahKaryawan(karyawanBaru)
}

func (usecase *AttendanceUsecaseImpl) UpdateKehadiran(id int, hadir bool) {
	usecase.repo.UpdateKehadiran(id, hadir)
}

func (usecase *AttendanceUsecaseImpl) HapusKaryawan(id int) {
	usecase.repo.HapusKaryawan(id)
}

func (usecase *AttendanceUsecaseImpl) TampilkanKehadiran() []*model.Karyawan {
	return usecase.repo.TampilkanKehadiran()
}
