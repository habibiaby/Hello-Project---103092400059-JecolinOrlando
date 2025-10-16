package main

import "fmt"

const MAX = 100

type Mahasiswa struct {
	ID, Nama, Jurusan, Status string
	NilaiTes                  int
}
type Jurusan struct{ Nama string }

var mhs [MAX]Mahasiswa
var jur [MAX]Jurusan
var jmlMhs, jmlJur int

func newMhs(id, nama, jurusan string, nilai int) Mahasiswa {
	status := "Ditolak"
	if nilai >= 70 {
		status = "Diterima"
	}
	return Mahasiswa{id, nama, jurusan, status, nilai}
}
func cariJur(n string) int {
	for i := 0; i < jmlJur; i++ {
		if jur[i].Nama == n {
			return i
		}
	}
	return -1
}
func cariMhs(id string) int {
	for i := 0; i < jmlMhs; i++ {
		if mhs[i].ID == id {
			return i
		}
	}
	return -1
}

func tambahJur() {
	var n string
	fmt.Print("Nama Jurusan: ")
	fmt.Scan(&n)
	if cariJur(n) != -1 {
		fmt.Println("Sudah ada.")
		return
	}
	if jmlJur < MAX {
		jur[jmlJur] = Jurusan{n}
		jmlJur++
		fmt.Println("Berhasil ditambah.")
	} else {
		fmt.Println("Penuh.")
	}
}
func hapusJur() {
	var n string
	fmt.Print("Nama Jurusan: ")
	fmt.Scan(&n)
	idx := cariJur(n)
	if idx == -1 {
		fmt.Println("Tidak ditemukan.")
		return
	}
	for i := idx; i < jmlJur-1; i++ {
		jur[i] = jur[i+1]
	}
	jmlJur--
	fmt.Println("Berhasil dihapus.")
}

func tambahMhs() {
	var id, nama, j string
	var n int
	fmt.Print("ID: ")
	fmt.Scan(&id)
	fmt.Print("Nama: ")
	fmt.Scan(&nama)
	fmt.Print("Jurusan: ")
	fmt.Scan(&j)
	if cariJur(j) == -1 {
		fmt.Println("Jurusan tidak terdaftar.")
		return
	}
	fmt.Print("Nilai Tes: ")
	fmt.Scan(&n)
	if jmlMhs < MAX {
		mhs[jmlMhs] = newMhs(id, nama, j, n)
		jmlMhs++
		fmt.Println("Mahasiswa ditambah.")
	} else {
		fmt.Println("Kapasitas penuh.")
	}
}

func editMhs() {
	var id, nama, j string
	var n int
	fmt.Print("ID Mahasiswa: ")
	fmt.Scan(&id)
	idx := cariMhs(id)
	if idx == -1 {
		fmt.Println("Tidak ditemukan.")
		return
	}
	fmt.Print("Nama Baru: ")
	fmt.Scan(&nama)
	fmt.Print("Jurusan Baru: ")
	fmt.Scan(&j)
	if cariJur(j) == -1 {
		fmt.Println("Jurusan tidak terdaftar.")
		return
	}
	fmt.Print("Nilai Tes Baru: ")
	fmt.Scan(&n)
	mhs[idx] = newMhs(id, nama, j, n)
	fmt.Println("Data diubah.")
}

func hapusMhs() {
	var id string
	fmt.Print("ID Mahasiswa: ")
	fmt.Scan(&id)
	idx := cariMhs(id)
	if idx == -1 {
		fmt.Println("Tidak ditemukan.")
		return
	}
	for i := idx; i < jmlMhs-1; i++ {
		mhs[i] = mhs[i+1]
	}
	jmlMhs--
	fmt.Println("Data dihapus.")
}

func sortNama() {
	var temp Mahasiswa
	for i := 1; i < jmlMhs; i++ {
		temp = mhs[i]
		j := i - 1
		for j >= 0 && mhs[j].Nama > temp.Nama {
			mhs[j+1] = mhs[j]
			j--
		}
		mhs[j+1] = temp
	}
}
func sortNilai() {
	var temp Mahasiswa
	for i := 0; i < jmlMhs-1; i++ {
		min := i
		for j := i + 1; j < jmlMhs; j++ {
			if mhs[j].NilaiTes < mhs[min].NilaiTes {
				min = j
			}
		}
		temp = mhs[i]
		mhs[i] = mhs[min]
		mhs[min] = temp
	}
}

func tampil(filter string) {
	fmt.Println()
	for i := 0; i < jmlMhs; i++ {
		if filter == "Diterima" && mhs[i].Status != "Diterima" {
			continue
		}
		if filter == "Ditolak" && mhs[i].Status != "Ditolak" {
			continue
		}
		fmt.Println(mhs[i])
	}
}

func tampilByJur() {
	var j string
	fmt.Print("Nama Jurusan: ")
	fmt.Scan(&j)
	if cariJur(j) == -1 {
		fmt.Println("Jurusan tidak terdaftar.")
		return
	}
	fmt.Println("\n--- Mahasiswa", j, "---")
	found := false
	for i := 0; i < jmlMhs; i++ {
		if mhs[i].Jurusan == j {
			fmt.Println(mhs[i])
			found = true
		}
	}
	if !found {
		fmt.Println("Belum ada mahasiswa.")
	}
}

func cekStatus() {
	var id string
	fmt.Print("Masukkan ID Anda: ")
	fmt.Scan(&id)
	for i := 0; i < jmlMhs; i++ {
		if mhs[i].ID == id {
			fmt.Printf("\nID: %s\nNama: %s\nJurusan: %s\nNilai: %d\nStatus: %s\n",
				mhs[i].ID, mhs[i].Nama, mhs[i].Jurusan, mhs[i].NilaiTes, mhs[i].Status)
			return
		}
	}
	fmt.Println("ID tidak ditemukan.")
}

func menuAdmin() {
	for {
		fmt.Println("\n=== Menu Admin ===")
		fmt.Println("1. Tambah Jurusan\n2. Tambah Mahasiswa\n3. Edit Mahasiswa\n4. Hapus Mahasiswa\n5. Hapus Jurusan")
		fmt.Println("6. Urut Nama\n7. Urut Nilai\n8. Berdasar Jurusan\n9. Diterima\n10. Ditolak\n0. Kembali")
		fmt.Print("Pilih: ")
		var p int
		fmt.Scan(&p)
		if p == 1 {
			tambahJur()
		} else if p == 2 {
			tambahMhs()
		} else if p == 3 {
			editMhs()
		} else if p == 4 {
			hapusMhs()
		} else if p == 5 {
			hapusJur()
		} else if p == 6 {
			sortNama()
			fmt.Println("\n--- Mahasiswa Urut Nama ---")
			tampil("")
		} else if p == 7 {
			sortNilai()
			fmt.Println("\n--- Mahasiswa Urut Nilai ---")
			tampil("")
		} else if p == 8 {
			tampilByJur()
		} else if p == 9 {
			fmt.Println("\n--- Mahasiswa Diterima ---")
			tampil("Diterima")
		} else if p == 10 {
			fmt.Println("\n--- Mahasiswa Ditolak ---")
			tampil("Ditolak")
		} else if p == 0 {
			return
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func main() {
	for {
		fmt.Println("\n=== Sistem Pendaftaran ===")
		fmt.Println("1. Admin\n2. Calon Mahasiswa\n0. Keluar")
		fmt.Print("Pilih: ")
		var p int
		fmt.Scan(&p)
		if p == 1 {
			menuAdmin()
		} else if p == 2 {
			if jmlMhs == 0 {
				fmt.Println("Belum ada data.")
			} else {
				cekStatus()
			}
		} else if p == 0 {
			fmt.Println("Terima kasih!")
			return
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
