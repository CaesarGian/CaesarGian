package main

import "fmt"

const NMAX int = 100

type tenant struct {
	NamaTenant, IDTenant string
	tabTransaksi         [NMAX]transaksi
	jmlTransaksi         int
}

type transaksi struct {
	IDTransaksi, namaPembeli, namaBarang string
	totalBeli, hargaBarang, totalHarga   int
}

type kantin [NMAX]tenant

func main() {
	var data kantin
	mainMenu(&data)
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func mainMenu(data *kantin) {
	var Menu int
	fmt.Println()
	fmt.Println("*** ------------------------------------------------------------------------------- ***")
	fmt.Println("                                 Aplikasi Kantin Tel-U                                 ")
	fmt.Println("                         Created by Caesar Gian & Yudis Akbar                          ")
	fmt.Println("                                   Class: IF-47-10                                     ")
	fmt.Println("*** ------------------------------------------------------------------------------- ***")
	fmt.Println("Pilih menu:")
	fmt.Println("1. Tambah data")
	fmt.Println("2. Hapus data")
	fmt.Println("3. Ubah data")
	fmt.Println("4. Admin")
	fmt.Println("0. Exit")
	fmt.Println("*** ------------------------------------------------------------------------------- ***")
	fmt.Print("Pilih (1/2/3/4/0): ")
	fmt.Scan(&Menu)

	if Menu == 0 {
		fmt.Println("Terima kasih telah menggunakan aplikasi Kantin Tel-U.")
		return
	} else if Menu == 1 {
		tambahData(data)
	} else if Menu == 2 {
		hapusData(data)
	} else if Menu == 3 {
		ubahData(data)
	} else if Menu == 4 {
		menuAdmin(data)
	} else {
		fmt.Println("Menu tidak valid. Silakan pilih menu yang tersedia.")
	}
	mainMenu(data)
}

func tambahData(data *kantin) {
	var Menu int
	fmt.Println("*** ------------------------------------------------------------------------------- ***")
	fmt.Println("1. Data Tenant")
	fmt.Println("2. Data Transaksi")
	fmt.Println("3. Kembali ke menu utama")
	fmt.Print("Pilih (1/2/3): ")
	fmt.Scan(&Menu)
	fmt.Println("*** ------------------------------------------------------------------------------- ***")

	if Menu == 1 {
		DataTenant(data)
	} else if Menu == 2 {
		DataTransaksi(data)
	} else if Menu == 3 {
		return
	} else {
		fmt.Println("Menu tidak valid. Silakan pilih menu yang tersedia.")
	}
	tambahData(data)
}

func DataTenant(data *kantin) {
	var namaTenant, IDTenant string
	fmt.Print("Nama tenant: ")
	fmt.Scan(&namaTenant)
	fmt.Print("ID Tenant: ")
	fmt.Scan(&IDTenant)

	addTenant(data, namaTenant, IDTenant, 0)
	fmt.Println("Data tenant berhasil ditambahkan.")
}

func addTenant(data *kantin, namaTenant, IDTenant string, index int) {
	if index < NMAX && data[index].NamaTenant == "" {
		data[index].NamaTenant = namaTenant
		data[index].IDTenant = IDTenant
		data[index].jmlTransaksi = 0
	} else if index < NMAX {
		addTenant(data, namaTenant, IDTenant, index+1)
	}
}

func DataTransaksi(data *kantin) {
	var IDTenant, IDTransaksi, namaPembeli, namaBarang string
	var totalBeli, hargaBarang int

	fmt.Print("ID Tenant: ")
	fmt.Scan(&IDTenant)
	indexTenant := cariTenant(*data, IDTenant, 0)

	if indexTenant == -1 {
		fmt.Println("Tenant tidak ditemukan.")
		return
	}

	fmt.Print("ID Transaksi: ")
	fmt.Scan(&IDTransaksi)
	if IDTransaksi == "0" {
		return
	}
	fmt.Print("Nama pembeli: ")
	fmt.Scan(&namaPembeli)
	fmt.Print("Nama barang: ")
	fmt.Scan(&namaBarang)
	fmt.Print("Total beli: ")
	fmt.Scan(&totalBeli)
	fmt.Print("Harga barang: ")
	fmt.Scan(&hargaBarang)

	totalHarga := totalBeli * hargaBarang

	addTransaksi(&data[indexTenant], IDTransaksi, namaPembeli, namaBarang, totalBeli, hargaBarang, totalHarga, data[indexTenant].jmlTransaksi)

	totalPembayaran := hargaBarang * totalBeli
	komisiAdmin := totalPembayaran * 25 / 100

	fmt.Printf("Total pembayaran: %d\n", totalPembayaran)
	fmt.Printf("Komisi untuk admin kantin: %d\n", komisiAdmin)
}

func addTransaksi(t *tenant, IDTransaksi, namaPembeli, namaBarang string, totalBeli, hargaBarang, totalHarga, indexTransaksi int) {
	if indexTransaksi < NMAX {
		t.tabTransaksi[indexTransaksi] = transaksi{
			IDTransaksi: IDTransaksi,
			namaPembeli: namaPembeli,
			namaBarang:  namaBarang,
			totalBeli:   totalBeli,
			hargaBarang: hargaBarang,
			totalHarga:  totalHarga,
		}
		t.jmlTransaksi++
	}
}

func cariTenant(data kantin, IDTenant string, index int) int {
	if index >= NMAX {
		return -1
	}
	if data[index].IDTenant == IDTenant {
		return index
	}
	return cariTenant(data, IDTenant, index+1)
}

func hapusData(data *kantin) {
	var Menu int
	fmt.Println("*** ------------------------------------------------------------------------------- ***")
	fmt.Println("1. Hapus Tenant")
	fmt.Println("2. Hapus Transaksi")
	fmt.Println("3. Kembali ke menu utama")
	fmt.Print("Pilih (1/2/3): ")
	fmt.Scan(&Menu)
	fmt.Println("*** ------------------------------------------------------------------------------- ***")

	if Menu == 1 {
		hapusTenant(data)
	} else if Menu == 2 {
		hapusTransaksi(data)
	} else if Menu == 3 {
		return
	} else {
		fmt.Println("Menu tidak valid. Silakan pilih menu yang tersedia.")
	}
	hapusData(data)
}

func hapusTenant(data *kantin) {
	var IDTenant string
	fmt.Print("Masukkan ID Tenant yang ingin dihapus: ")
	fmt.Scan(&IDTenant)

	index := cariTenant(*data, IDTenant, 0)
	if index == -1 {
		fmt.Println("Tenant tidak ditemukan.")
		return
	}

	data[index] = tenant{}
	fmt.Println("Data tenant telah dihapus.")
}

func hapusTransaksi(data *kantin) {
	var IDTenant, IDTransaksi string
	fmt.Print("Masukkan ID Tenant: ")
	fmt.Scan(&IDTenant)
	fmt.Print("Masukkan ID Transaksi yang ingin dihapus: ")
	fmt.Scan(&IDTransaksi)

	indexTenant := cariTenant(*data, IDTenant, 0)
	if indexTenant == -1 {
		fmt.Println("Tenant tidak ditemukan.")
		return
	}

	indexTransaksi := cariTransaksi(data[indexTenant], IDTransaksi, 0)
	if indexTransaksi == -1 {
		fmt.Println("Transaksi tidak ditemukan.")
		return
	}

	data[indexTenant].tabTransaksi[indexTransaksi] = transaksi{}
	fmt.Println("Data transaksi telah dihapus.")
}

func cariTransaksi(tenant tenant, IDTransaksi string, index int) int {
	if index >= tenant.jmlTransaksi {
		return -1
	}
	if tenant.tabTransaksi[index].IDTransaksi == IDTransaksi {
		return index
	}
	return cariTransaksi(tenant, IDTransaksi, index+1)
}

func ubahData(data *kantin) {
	var Menu int
	fmt.Println("*** ------------------------------------------------------------------------------- ***")
	fmt.Println("1. Ubah Data Tenant")
	fmt.Println("2. Ubah Data Transaksi")
	fmt.Println("3. Kembali ke menu utama")
	fmt.Print("Pilih (1/2/3): ")
	fmt.Scan(&Menu)
	fmt.Println("*** ------------------------------------------------------------------------------- ***")

	if Menu == 1 {
		ubahDataTenant(data)
	} else if Menu == 2 {
		ubahDataTransaksi(data)
	} else if Menu == 3 {
		return
	} else {
		fmt.Println("Menu tidak valid. Silakan pilih menu yang tersedia.")
	}
	ubahData(data)
}

func ubahDataTenant(data *kantin) {
	var IDTenant, namaTenantBaru string
	fmt.Print("Masukkan ID Tenant yang ingin diubah: ")
	fmt.Scan(&IDTenant)

	index := cariTenant(*data, IDTenant, 0)
	if index == -1 {
		fmt.Println("Tenant tidak ditemukan.")
		return
	}

	fmt.Print("Masukkan nama tenant baru: ")
	fmt.Scan(&namaTenantBaru)
	data[index].NamaTenant = namaTenantBaru
	fmt.Println("Data tenant berhasil diubah.")
}

func ubahDataTransaksi(data *kantin) {
	var IDTenant, IDTransaksi, namaPembeliBaru, namaBarangBaru string
	var totalBeliBaru, hargaBarangBaru int

	fmt.Print("Masukkan ID Tenant: ")
	fmt.Scan(&IDTenant)

	indexTenant := cariTenant(*data, IDTenant, 0)
	if indexTenant == -1 {
		fmt.Println("Tenant tidak ditemukan.")
		return
	}

	fmt.Print("Masukkan ID Transaksi yang ingin diubah: ")
	fmt.Scan(&IDTransaksi)

	indexTransaksi := cariTransaksi(data[indexTenant], IDTransaksi, 0)
	if indexTransaksi == -1 {
		fmt.Println("Transaksi tidak ditemukan.")
		return
	}

	fmt.Print("Masukkan nama pembeli baru: ")
	fmt.Scan(&namaPembeliBaru)
	fmt.Print("Masukkan nama barang baru: ")
	fmt.Scan(&namaBarangBaru)
	fmt.Print("Masukkan total beli baru: ")
	fmt.Scan(&totalBeliBaru)
	fmt.Print("Masukkan harga barang baru: ")
	fmt.Scan(&hargaBarangBaru)

	totalHargaBaru := totalBeliBaru * hargaBarangBaru

	data[indexTenant].tabTransaksi[indexTransaksi] = transaksi{
		IDTransaksi: IDTransaksi,
		namaPembeli: namaPembeliBaru,
		namaBarang:  namaBarangBaru,
		totalBeli:   totalBeliBaru,
		hargaBarang: hargaBarangBaru,
		totalHarga:  totalHargaBaru,
	}
	fmt.Println("Data transaksi berhasil diubah.")
}

func menuAdmin(data *kantin) {
	fmt.Println("*** ------------------------------------------------------------------------------- ***")
	fmt.Println("                                     Menu Admin                                        ")
	fmt.Println("1. Lihat data")
	fmt.Println("2. Lihat komisi admin")
	fmt.Println("3. Kembali ke menu utama")
	fmt.Println("*** ------------------------------------------------------------------------------- ***")
	var choice int
	fmt.Print("Pilih (1/2/3): ")
	fmt.Scan(&choice)

	if choice == 1 {
		printData(data, 0)
	} else if choice == 2 {
		printKomisiAdmin(data, 0, 0)
	} else if choice == 3 {
		return
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
	menuAdmin(data)
}

func printData(data *kantin, index int) {
	if index >= NMAX {
		return
	}

	t := data[index]
	if t.NamaTenant != "" {
		fmt.Printf("Tenant %d: %s (ID: %s)\n", index+1, t.NamaTenant, t.IDTenant)
		fmt.Printf("%10s %10s %15s %15s %10s %10s %10s\n", "Transaksi", "ID", "Pembeli", "Barang", "TotalBeli", "Harga", "TotalHarga")
		printTransaksi(t, 0)
	}
	printData(data, index+1)
}

func printTransaksi(t tenant, index int) {
	if index >= t.jmlTransaksi {
		return
	}

	trans := t.tabTransaksi[index]
	fmt.Printf("%10d %10s %15s %15s %10d %10d %10d\n",
		index+1, trans.IDTransaksi, trans.namaPembeli, trans.namaBarang, trans.totalBeli, trans.hargaBarang, trans.totalHarga)
	printTransaksi(t, index+1)
}

func printKomisiAdmin(data *kantin, index int, totalKomisi int) {
	if index >= NMAX {
		fmt.Printf("Total komisi untuk admin: %d\n", totalKomisi)
		return
	}

	t := data[index]
	if t.NamaTenant != "" {
		totalKomisi = calculateKomisi(t, 0, totalKomisi)
	}
	printKomisiAdmin(data, index+1, totalKomisi)
}

func calculateKomisi(t tenant, transIndex int, totalKomisi int) int {
	if transIndex >= t.jmlTransaksi {
		return totalKomisi
	}

	totalPembayaran := t.tabTransaksi[transIndex].hargaBarang * t.tabTransaksi[transIndex].totalBeli
	totalKomisi += totalPembayaran * 25 / 100
	return calculateKomisi(t, transIndex+1, totalKomisi)
}
