package repository

import "attendance-app/model"

type AttendanceRepository interface {
	TambahKaryawan(karyawan *model.Karyawan)
	UpdateKehadiran(id int, hadir bool)
	HapusKaryawan(id int)
	TampilkanKehadiran() []*model.Karyawan
}

type AttendanceRepoImpl struct {
	karyawanList []*model.Karyawan
}

func NewAttendanceRepository() AttendanceRepository {
	return &AttendanceRepoImpl{}
}

func (repo *AttendanceRepoImpl) TambahKaryawan(karyawan *model.Karyawan) {
	repo.karyawanList = append(repo.karyawanList, karyawan)
}

func (repo *AttendanceRepoImpl) UpdateKehadiran(id int, hadir bool) {
	for _, karyawan := range repo.karyawanList {
		if karyawan.ID == id {
			karyawan.Kehadiran = hadir
			return
		}
	}
}

func (repo *AttendanceRepoImpl) HapusKaryawan(id int) {
	for i, karyawan := range repo.karyawanList {
		if karyawan.ID == id {
			repo.karyawanList = append(repo.karyawanList[:i], repo.karyawanList[i+1:]...)
			return
		}
	}
}

func (repo *AttendanceRepoImpl) TampilkanKehadiran() []*model.Karyawan {
	return repo.karyawanList
}
