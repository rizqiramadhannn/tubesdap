//	Muhammad Rizqi Ramadhan (1301193325)

package main

import "fmt"

type(
	typepelanggan struct {
		nama, telp, motor string
	}
	typedata struct {
		pelanggan typepelanggan
		arrtransaksi [100] typetransaksi
	}
	typetransaksi struct {
		tipe, namasp, servis string
		bayar int
		tanggal, bulan, tahun int
	}
	typesparepart struct {
		nama, pabrikan string
		tahun, harga int
	}
)

var(
	arrdata [10000] typedata
	arrsparepart [10000] typesparepart
	countsp, countpel, countbayar int
)

func printpelanggan() {
	i:= 0
	nomor := 0
	for i < countpel {
		if arrdata[i].pelanggan.nama != "" {
			fmt.Printf("%d.	Nama : %s\n", nomor+1, arrdata[i].pelanggan.nama)
			fmt.Printf("	No. Telpon: %s\n",arrdata[i].pelanggan.telp)
			fmt.Printf("	Motor: %s\n",arrdata[i].pelanggan.motor)
			nomor++
		}
		i++
	}
}

func printpel() {
	i := 0
	nomor := 0
	for i < countsp {
		if arrsparepart[i].nama != "" {
			fmt.Printf("%d. %s\n", nomor+1, arrsparepart[i].nama)
			nomor++
		}
		i++
	}
}

func printsparepart() {
	i := 0
	nomor := 0
	for i < countsp {
		if arrsparepart[i].nama != "" {
			fmt.Printf("%d. Nama Sparepart : %s\n", nomor+1, arrsparepart[i].nama)
			fmt.Printf("	Pabrikan Sparepart : %s\n", arrsparepart[i].pabrikan)
			fmt.Printf("	Tahun Keluar : %d\n", arrsparepart[i].tahun)
			fmt.Printf("	Harga : %d\n", arrsparepart[i].harga)
			nomor++
		}
		i++
	}
}

func printtransaksi(index int) {
	countbayar = 100

	for i:= 0 ; i < countbayar ; i++ {
		if arrdata[index].arrtransaksi[i].tipe != "" {
			fmt.Printf("%d.	Tipe transaksi : %s\n", i+1, arrdata[index].arrtransaksi[i].tipe)
			if arrdata[index].arrtransaksi[i].tipe != "servis" {
				fmt.Printf("	Sparepart yang dibeli : %s\n", arrdata[index].arrtransaksi[i].namasp)
			} else {
				fmt.Printf("	Jenis servis : %s\n", arrdata[index].arrtransaksi[i].servis)
			}
			fmt.Printf("	Tanggal transaksi : %d %d %d\n", arrdata[index].arrtransaksi[i].tanggal, arrdata[index].arrtransaksi[i].bulan, arrdata[index].arrtransaksi[i].tahun)
			fmt.Printf("	Jumlah yang harus dibayar : %d\n", arrdata[index].arrtransaksi[i].bayar)
		}
	}
}

func menu () {
	/*	I.S Program dimulai dengan fungsi ini sebagai menu utama
		F.S Program menampilkan menu utama dalam program */
	fmt.Println("Selamat Datang di app Service Motor")
	fmt.Println("")
	fmt.Println("Pilih nomor dari menu yang ingin dibuka")
	fmt.Println("1. Tambah data")
	fmt.Println("2. Edit data")
	fmt.Println("3. Hapus data")
	fmt.Println("4. Cari pelanggan Service")
	fmt.Println("5. Cari pelanggan pembeli spare-part")
	fmt.Println("6. Daftar pabrikan")
	fmt.Println("7. Pemasukan tempat service perbulan")
}

func tambah(pilih int) {
	/*	I.S Tersedia submenu dari menu tambah data yang menanyakan user akan memasukkan data apa
		F.S User dapat memilih akan masuk ke submenu mana dengan memilih berdasarkan nomor submenu */
	var nama, pabrikan, telp, motor, servis, tipe, namasp string
	var tahun, harga, bayar, tanggal, bulan int

	pilih = 1

	fmt.Println("Pilih jenis data")
	fmt.Println("1. Data Sparepart")
	fmt.Println("2. Data Pelanggan")
	fmt.Println("3. Data Transaksi")
	fmt.Print("Pilihan : ")
	fmt.Scanln(&pilih)
	fmt.Println("")

	for pilih != 1 && pilih != 2 && pilih != 3 {
		fmt.Println("Maaf menu tidak tersedia")
		fmt.Println("")
		fmt.Println("Pilih jenis data")
		fmt.Println("1. Data Sparepart")
		fmt.Println("2. Data Pelanggan")
		fmt.Println("3. Data Transaksi")
		fmt.Print("Pilihan : ")
		fmt.Scanln(&pilih)
		fmt.Println("")
	}
	if pilih == 1{
		tambahsukucadang(nama, pabrikan, tahun, harga)
	} else if pilih == 2{
		tambahpelanggan(nama, telp, motor)
	} else if pilih == 3{
		tambahtransaksi(tipe, namasp, bayar, tanggal, bulan, tahun, servis)
	}
}

func tambahpelanggan(nama, telp, motor string) {
	/*	I.S Program akan meminta user memasukkan data dari pelanggan
		F.S Data pelanggan akan tersimpan di array sesuai sudah berapa pelanggan yang sudah diinput */

	var back string

	back = "Y"
	for back == "Y" || back == "y" {
		fmt.Print("Masukkan nama pelanggan : ")
		fmt.Scanln(&nama)
		arrdata[countpel].pelanggan.nama = nama
		fmt.Print("Masukkan nomor telepon : ")
		fmt.Scanln(&telp)
		arrdata[countpel].pelanggan.telp = telp
		fmt.Print("Masukkan tipe motor : ")
		fmt.Scanln(&motor)
		arrdata[countpel].pelanggan.motor = motor
		fmt.Print("Masukkan Y untuk memasukkan data lagi : ")
		fmt.Scanln(&back)
		fmt.Println("")
		countpel++
	}
}

func tambahsukucadang(nama, pabrikan string, tahun, harga int) {
	/*	I.S Program akan meminta user memasukkan data dari sparepart
		F.S Data sparepart akan tersimpan di array sesuai sudah berapa sparepart yang sudah diinput */
	var back string

	back = "Y"
	for back == "Y" || back == "y" {
		fmt.Print("Masukkan nama sparepart : ")
		fmt.Scan(&nama)
		arrsparepart[countsp].nama = nama
		fmt.Print("Masukkan pabrikan sparepart : ")
		fmt.Scan(&pabrikan)
		arrsparepart[countsp].pabrikan = pabrikan
		fmt.Print("Masukkan tahun keluar : ")
		fmt.Scan(&tahun)
		arrsparepart[countsp].tahun = tahun
		fmt.Print("Masukkan harga : ")
		fmt.Scan(&harga)
		arrsparepart[countsp].harga = harga
		fmt.Print("Masukkan Y untuk memasukkan data lagi : ")
		fmt.Scan(&back)
		fmt.Println("")
		countsp++
	}
}

func tambahtransaksi(tipe, namasp string, bayar int, tanggal, bulan, tahun int, servis string) {
	var back string
	var i, no int

	back = "Y"
	for back == "Y" || back == "y" {
		if countpel == 0 {
			fmt.Println("Tidak ada data pelanggan")
		} else {
			fmt.Println("Data Pelanggan")
			printpelanggan()
			fmt.Print("Pilih Pelanggan: ")
			fmt.Scanln(&no)
			for no > i && no < 0 {
				fmt.Println("Maaf menu tidak tersedia")
				fmt.Println("")
				fmt.Println("Data Pelanggan")
				printpelanggan()
				fmt.Print("Pilih pelanggan : ")
				fmt.Scanln(&no)
			}

			countbayar = 0
			for arrdata[no-1].arrtransaksi[countbayar].tipe != "" {
				countbayar++
			}
			fmt.Print("Masukkan nama tipe pelanggan (servis/sparepart) : ")
			fmt.Scanln(&tipe)
			arrdata[no-1].arrtransaksi[countbayar].tipe = tipe
			for tipe != "servis" && tipe != "Servis" && tipe != "sparepart" && tipe != "Sparepart" {
				fmt.Println("Maaf menu tidak tersedia")
				fmt.Println("")
				fmt.Print("Masukkan nama tipe pelanggan (servis/sparepart) : ")
				fmt.Scanln(&tipe)
				arrdata[no-1].arrtransaksi[countbayar].tipe = tipe
			}
			if tipe == "servis" || tipe == "Servis" {
				fmt.Print("Masukkan tipe servis (ringan/sedang/berat) : ")
				fmt.Scanln(&servis)
				arrdata[no-1].arrtransaksi[countbayar].servis = servis
				for servis != "ringan" && servis != "sedang" && servis != "berat"{
					fmt.Println("Maaf menu tidak tersedia")
					fmt.Println("")
					fmt.Print("Masukkan tipe servis (ringan/sedang/berat) : ")
					fmt.Scanln(&servis)
					arrdata[no-1].arrtransaksi[countbayar].servis = servis
				}
				if servis == "ringan" {
					arrdata[no-1].arrtransaksi[countbayar].bayar = 25000
				} else if servis == "sedang" {
					arrdata[no-1].arrtransaksi[countbayar].bayar = 50000
				} else {
					arrdata[no-1].arrtransaksi[countbayar].bayar = 100000
				}
				fmt.Printf("Jumlah yang harus dibayar : %d\n", arrdata[no-1].arrtransaksi[countbayar].bayar)
			} else if tipe == "sparepart" || tipe == "Sparepart" {
				fmt.Print("Masukkan nama sparepart : ")
				fmt.Scanln(&namasp)
				arrdata[no-1].arrtransaksi[countbayar].namasp = namasp
				for i = 0 ; i < countsp && namasp != arrsparepart[i].nama ; i++ {
				}
				arrdata[no-1].arrtransaksi[countbayar].bayar = arrsparepart[i].harga
			}
			fmt.Print("Masukkan tanggal : ")
			fmt.Scanln(&tanggal)
			arrdata[no-1].arrtransaksi[countbayar].tanggal = tanggal
			fmt.Print("Masukkan bulan : ")
			fmt.Scanln(&bulan)
			arrdata[no-1].arrtransaksi[countbayar].bulan = bulan
			fmt.Print("Masukkan tahun : ")
			fmt.Scanln(&tahun)
			arrdata[no-1].arrtransaksi[countbayar].tahun = tahun
			fmt.Print("Masukkan Y untuk memasukkan data lagi : ")
			fmt.Scanln(&back)
			fmt.Println("")
			countbayar++
		}
	}
}

func edit(pilih int) {
	/*	I.S Tersedia submenu dari menu edit data yang menanyakan user akan meng-edit data apa
		F.S User dapat memilih akan masuk ke submenu mana dengan memilih nomor submenu */
	var nama, pabrikan, telp, motor, tipe string
	var tahun, tanggal, bulan, harga int

	pilih = 2

	fmt.Println("Pilih Data Yang Akan di Edit")
	fmt.Println("1. Edit Data Sparepart")
	fmt.Println("2. Edit Data Pelanggan")
	fmt.Println("3. Edit Data Transaksi")
	fmt.Print("Pilihan : ")
	fmt.Scanln(&pilih)
	fmt.Println("")

	for pilih != 1 && pilih != 2 && pilih != 3 {
		fmt.Println("Maaf Menu Tidak Tersedia")
		fmt.Println("")
		fmt.Println("Pilih Data Yang Akan di Edit")
		fmt.Println("1. Edit Data Sparepart")
		fmt.Println("2. Edit Data Pelanggan")
		fmt.Print("Pilihan : ")
		fmt.Scanln(&pilih)
		fmt.Println("")
	}
	if pilih == 1 {
		editsparepart(nama, pabrikan, tahun, harga)
	} else if pilih == 2 {
		editpelanggan(nama, telp, motor)
	} else if pilih == 3 {
		edittransaksi(tipe, tanggal, bulan, tahun)
	}
}

func edittransaksi(tipe string, tanggal, bulan, tahun int) {
	var back, servis, namasp string
	var i, no, pilihedit int

	back = "Y"
	for back == "Y" || back == "y" {
		if countpel == 0 {
			fmt.Println("Tidak ada data pelanggan")
		} else {
			fmt.Println("Data Pelanggan")
			printpelanggan()
			fmt.Print("Pilih Pelanggan: ")
			fmt.Scanln(&no)
			for no > i && no < 0 {
				fmt.Println("Maaf menu tidak tersedia")
				fmt.Println("")
				fmt.Println("Data Pelanggan")
				printpelanggan()
				fmt.Print("Pilih pelanggan : ")
				fmt.Scanln(&no)
			}

			countbayar = 0
			for arrdata[no-1].arrtransaksi[countbayar].tipe != "" {
				countbayar++
			}
			printtransaksi(no-1)
			fmt.Print("Pilih transaksi yang akan diedit : ")
			fmt.Scanln(&pilihedit)
			fmt.Print("Masukkan nama tipe pelanggan (servis/sparepart) : ")
			fmt.Scanln(&tipe)
			arrdata[no-1].arrtransaksi[pilihedit-1].tipe = tipe
			for tipe != "servis" && tipe != "Servis" && tipe != "sparepart" && tipe != "Sparepart" {
				fmt.Println("Maaf menu tidak tersedia")
				fmt.Println("")
				fmt.Print("Masukkan nama tipe pelanggan (servis/sparepart) : ")
				fmt.Scanln(&tipe)
				arrdata[no-1].arrtransaksi[pilihedit-1].tipe = tipe
			}
			if tipe == "servis" || tipe == "Servis" {
				fmt.Print("Masukkan tipe servis (ringan/sedang/berat) : ")
				fmt.Scanln(&servis)
				arrdata[no-1].arrtransaksi[pilihedit-1].servis = servis
				for servis != "ringan" && servis != "sedang" && servis != "berat"{
					fmt.Println("Maaf menu tidak tersedia")
					fmt.Println("")
					fmt.Print("Masukkan tipe servis (ringan/sedang/berat) : ")
					fmt.Scanln(&servis)
					arrdata[no-1].arrtransaksi[pilihedit-1].servis = servis
				}
				if servis == "ringan" {
					arrdata[no-1].arrtransaksi[pilihedit-1].bayar = 25000
				} else if servis == "sedang" {
					arrdata[no-1].arrtransaksi[pilihedit-1].bayar = 50000
				} else {
					arrdata[no-1].arrtransaksi[pilihedit-1].bayar = 100000
				}
				fmt.Printf("Jumlah yang harus dibayar : %d\n", arrdata[no-1].arrtransaksi[pilihedit-1].bayar)
			} else if tipe == "sparepart" || tipe == "Sparepart" {
				fmt.Print("Masukkan nama sparepart : ")
				fmt.Scanln(&namasp)
				arrdata[no-1].arrtransaksi[pilihedit-1].namasp = namasp
				for i = 0 ; i < countsp && namasp != arrsparepart[i].nama ; i++ {
				}
				arrdata[no-1].arrtransaksi[pilihedit-1].bayar = arrsparepart[i].harga
			}
			fmt.Print("Masukkan tanggal : ")
			fmt.Scanln(&tanggal)
			arrdata[no-1].arrtransaksi[pilihedit-1].tanggal = tanggal
			fmt.Print("Masukkan bulan : ")
			fmt.Scanln(&bulan)
			arrdata[no-1].arrtransaksi[pilihedit-1].bulan = bulan
			fmt.Print("Masukkan tahun : ")
			fmt.Scanln(&tahun)
			arrdata[no-1].arrtransaksi[pilihedit-1].tahun = tahun
			fmt.Print("Masukkan Y untuk edit data lagi : ")
			fmt.Scanln(&back)
			fmt.Println("")
		}
	}
}

func editpelanggan(nama, telp, motor string) {
	var i, no, pilihnomor int
	var back string

	back = "Y"
	for back == "Y" || back == "y" {
		if countpel == 0 {
			fmt.Println("Tidak ada data pelanggan")
		} else {
			fmt.Println("Data Pelanggan")
			printpelanggan()
			fmt.Print("Pilih data: ")
			fmt.Scanln(&no)
			for no > i && no < 0 {
				fmt.Println("Maaf menu tidak tersedia")
				fmt.Println("")
				fmt.Println("Data Pelanggan")
				printpelanggan()
				fmt.Print("Pilih data: ")
				fmt.Scanln(&no)
			} 

			fmt.Printf("1. Nama : %s\n", arrdata[no-1].pelanggan.nama)
			fmt.Printf("2. No. Telpon: %s\n",arrdata[no-1].pelanggan.telp)
			fmt.Printf("3. Motor: %s\n",arrdata[no-1].pelanggan.motor)
			fmt.Print("Pilih data yang akan diedit : ")
			fmt.Scanln(&pilihnomor)
			for pilihnomor > 3 && pilihnomor < 1 {
				fmt.Println("Maaf menu tidak tersedia")
				fmt.Println("")
				fmt.Printf("1. Nama : %s\n", arrdata[no-1].pelanggan.nama)
				fmt.Printf("2. No. Telpon: %s\n",arrdata[no-1].pelanggan.telp)
				fmt.Printf("3. Motor: %s\n",arrdata[no-1].pelanggan.motor)
				fmt.Print("Pilih data yang akan diedit : ")
				fmt.Scanln(&pilihnomor)
			}
			if pilihnomor == 1 {
				fmt.Print("Masukkan nama pelanggan : ")
				fmt.Scanln(&nama)
				arrdata[no-1].pelanggan.nama = nama	
			} else if pilihnomor == 2 {
				fmt.Print("Masukkan nomor telepon : ")
				fmt.Scanln(&telp)
				arrdata[no-1].pelanggan.telp = telp
			} else {
				fmt.Print("Masukkan tipe motor : ")
				fmt.Scanln(&motor)
				arrdata[no-1].pelanggan.motor = motor
			}
		}
		fmt.Print("Masukkan Y untuk mengedit data lagi : ")
		fmt.Scanln(&back)
		fmt.Println("")
	}
}

func editsparepart(nama, pabrikan string, tahun, harga int) {
	var i, no, pilihnomor int
	var back string

	back = "Y"
	for back == "Y" || back == "y" {
		if countsp == 0 {
			fmt.Println("Tidak ada data sparepart")
		} else {
			fmt.Println("Data Sparepart")
			printsparepart()
			fmt.Print("Pilih data: ")
			fmt.Scanln(&no)
			for no > i && no < 0 {
				fmt.Println("Maaf menu tidak tersedia")
				fmt.Println("")
				fmt.Println("Data Sparepart")
				printsparepart()
				fmt.Print("Pilih data: ")
				fmt.Scanln(&no)
			} 

			fmt.Printf("1. Nama Sparepart : %s\n",arrsparepart[no-1].nama)
			fmt.Printf("2. Pabrikan Sparepart : %s\n", arrsparepart[no-1].pabrikan)
			fmt.Printf("3. Tahun Keluar : %d\n", arrsparepart[no-1].tahun)
			fmt.Printf("4. Harga : %d\n", arrsparepart[no-1].harga)
			fmt.Print("Pilih data yang akan diedit : ")
			fmt.Scanln(&pilihnomor)
			for pilihnomor > 3 && pilihnomor < 1 {
				fmt.Println("Maaf menu tidak tersedia")
				fmt.Println("")
				fmt.Printf("1. Nama Sparepart : %s\n",arrsparepart[no-1].nama)
				fmt.Printf("2. Pabrikan Sparepart : %s\n", arrsparepart[no-1].pabrikan)
				fmt.Printf("3. Tahun Keluar : %d\n", arrsparepart[no-1].tahun)
				fmt.Printf("4. Harga : %d\n", arrsparepart[no-1].harga)
				fmt.Print("Pilih data yang akan diedit : ")
				fmt.Scanln(&pilihnomor)
			}
			if pilihnomor == 1 {
				fmt.Print("Masukkan nama sparepart : ")
				fmt.Scanln(&nama)
				arrsparepart[no-1].nama = nama	
			} else if pilihnomor == 2 {
				fmt.Print("Masukkan pabrikan sparepart : ")
				fmt.Scanln(&pabrikan)
				arrsparepart[no-1].pabrikan = pabrikan
			} else if pilihnomor == 3 {
				fmt.Print("Masukkan tahun keluar : ")
				fmt.Scanln(&tahun)
				arrsparepart[no-1].tahun = tahun
			} else {
				fmt.Print("Masukkan Harga Sparepart : ")
				fmt.Scanln(&harga)
				arrsparepart[no].harga = harga
			}
		}
		fmt.Print("Masukkan Y untuk mengedit data lagi : ")
		fmt.Scanln(&back)
		fmt.Println("")
	}
}

func hapus(pilih int) {
	/*	I.S Tersedia submenu dari menu hapus data yang menanyakan user akan meng-hapus data apa
		F.S User dapat memilih akan masuk ke submenu mana dengan memilih nomor submenu */
	var nama, pabrikan, telp, motor, tipe string
	var tahun, harga, bulan, tanggal int

	fmt.Println("Pilih Data Yang Akan di Hapus")
	fmt.Println("1. Hapus Data Sparepart")
	fmt.Println("2. Hapus Data Pelanggan")
	fmt.Println("3. Hapus Data Transaksi")
	fmt.Print("Pilihan : ")
	fmt.Scanln(&pilih)
	fmt.Println("")
	if pilih == 1 {
		hapussparepart(nama, pabrikan, tahun, harga)
	} else if pilih == 2 {
		hapuspelanggan(nama, telp, motor, tipe)
	} else {
		hapustransaksi(tipe, tanggal, bulan, tahun)
	}
}

func hapuspelanggan(nama, telp, motor, tipe string) {
	var i, no int
	var back string

	back = "Y"
	for back == "Y" || back == "y" {
		if countpel == 0 {
			fmt.Println("Tidak ada data pelanggan")
		} else {
			fmt.Println("Data Pelanggan")
			printpelanggan()
			fmt.Print("Masukkan nama pelanggan yang ingin dihapus : ")
			fmt.Scanln(&nama)
			no = carihapus(nama)
			for no > i && no < 0 {
				fmt.Println("Maaf menu tidak tersedia")
				fmt.Println("")
				fmt.Println("Data Pelanggan")
				printpelanggan()
				fmt.Print("Masukkan nama pelanggan yang ingin dihapus : ")
				fmt.Scanln(&nama)
				no = carihapus(nama)
			}
			arrdata[no].pelanggan.nama = ""
			arrdata[no].pelanggan.telp = ""
			arrdata[no].pelanggan.motor = ""
		}
		fmt.Print("Masukkan Y untuk menghapus data lagi : ")
		fmt.Scanln(&back)
		fmt.Println("")
	}
}

func hapussparepart(nama, pabrikan string, tahun, harga int) {
	var no int
	var back string

	back = "Y"
	for back == "Y" || back == "y" {
		if countsp == 0 {
			fmt.Println("Tidak ada data sparepart")
		} else {
			fmt.Println("Data Sparepart")
			printsparepart()
			fmt.Print("Pilih data: ")
			fmt.Scanln(&no)
			for no < 0 {
				fmt.Println("Maaf menu tidak tersedia")
				fmt.Println("")
				fmt.Println("Data Sparepart")
				printsparepart()
				fmt.Print("Pilih data: ")
				fmt.Scanln(&no)
			}

			arrsparepart[no-1].nama = ""
			arrsparepart[no-1].pabrikan = ""
			arrsparepart[no-1].tahun = -1
			arrsparepart[no-1].harga = -1

			fmt.Print("Masukkan Y untuk menghapus data lagi : ")
			fmt.Scanln(&back)
			fmt.Println("")
		}
	}
}

func hapustransaksi(tipe string, tanggal, bulan, tahun int) {
	var back string
	var i, no, pilihhapus int

	back = "Y"
	for back == "Y" || back == "y" {
		if countpel == 0 {
			fmt.Println("Tidak ada data pelanggan")
		} else {
			fmt.Println("Data Pelanggan")
			printpelanggan()
			fmt.Print("Pilih Pelanggan: ")
			fmt.Scanln(&no)
			for no > i && no < 0 {
				fmt.Println("Maaf menu tidak tersedia")
				fmt.Println("")
				fmt.Println("Data Pelanggan")
				printpelanggan()
				fmt.Print("Pilih pelanggan : ")
				fmt.Scanln(&no)
			}

			countbayar = 0
			for arrdata[no-1].arrtransaksi[countbayar].tipe != "" {
				countbayar++
			}
			printtransaksi(no-1)
			fmt.Print("Pilih transaksi yang akan dihapus : ")
			fmt.Scanln(&pilihhapus)

			arrdata[no-1].arrtransaksi[pilihhapus-1].tipe = ""
			arrdata[no-1].arrtransaksi[pilihhapus-1].namasp = ""
			arrdata[no-1].arrtransaksi[pilihhapus-1].servis = ""
			arrdata[no-1].arrtransaksi[pilihhapus-1].tanggal = -1
			arrdata[no-1].arrtransaksi[pilihhapus-1].bulan = -1
			arrdata[no-1].arrtransaksi[pilihhapus-1].tahun = -1
			arrdata[no-1].arrtransaksi[pilihhapus-1].bayar = -1
			fmt.Print("Masukkan Y untuk menghapus data lagi : ")
			fmt.Scanln(&back)
			fmt.Println("")
		}
	}
}

func caripelanggan(pilih int) {
	var back string
	var i, tgl1, tgl2, bln1, bln2, thn1, thn2, j int
	var valid, awal, akhir int

	back = "Y"
	for back == "Y" || back == "y" {
		fmt.Print("Masukkan tanggal mulai (dd mm yyyy) :")
		fmt.Scan(&tgl1, &bln1, &thn1)
		awal = bln1*30 + thn1*365 + tgl1
		fmt.Print("Masukkan tanggal akhir (dd mm yyyy) :")
		fmt.Scan(&tgl2, &bln2, &thn2)
		akhir = bln2*30 + thn2*365 + tgl2
		i = 0
		for i < countpel{
			j = 0
			for j < countbayar {
				valid = arrdata[i].arrtransaksi[j].tanggal + arrdata[i].arrtransaksi[j].bulan*30 + arrdata[i].arrtransaksi[j].tahun*365 
				if arrdata[i].arrtransaksi[j].tipe == "servis" && valid >= awal && valid <= akhir{
					fmt.Printf("Nama : %s\n", arrdata[i].pelanggan.nama)
					fmt.Printf("No. Telpon: %s\n",arrdata[i].pelanggan.telp)
					fmt.Printf("Motor: %s\n",arrdata[i].pelanggan.motor)
					fmt.Println("")
				}	
				j++
			}
			i++
		}
		fmt.Print("Masukkan Y untuk mencari pelanggan lagi : ")
		fmt.Scan(&back)
		fmt.Println("")
	}
}

func carihapus(nama string) int {
	var i int

	for i < countpel && nama != arrdata[i].pelanggan.nama{
		i++
	}
	if nama == arrdata[i].pelanggan.nama {
		return i
	} else {
		return -1
	}
}

func caripembeli(pilih int) {
	var i, nomor int
	var back string 

	back = "Y"
	for back == "y" || back == "Y" {
		if countsp == 0 {
			fmt.Println("Maaf Tidak Ada Data Sparepart")
		} else {
            fmt.Println("Data Nama Sparepart")
			printpel()
            fmt.Print("PIlih Sparepart: ")
            fmt.Scanln(&nomor)
            for nomor > i && nomor < 0 {
                fmt.Println("Maaf Data Sparepart Tidak Ada")
                fmt.Println("")
                fmt.Println("Data Nama Sparepart")
                printpel()
                fmt.Print("Pilih Sparepart: ")
                fmt.Scanln(&nomor)
			}
			for i < countpel{
				hitung:= 0
				countbayar = 100
				for j:= 0 ; j < countbayar ; j++ {
					if arrsparepart[nomor-1].nama == arrdata[i].arrtransaksi[j].namasp {
						hitung++
						fmt.Printf("%d.	%s\n", hitung, arrdata[i].pelanggan.nama)
					}
				}
				i++
			}
        }
        fmt.Print("Masukkan Y untuk mencari pelanggan lagi : ")
		fmt.Scan(&back)
		fmt.Println("")
	}
}

func daftarpabrikan(pilih int) {
	fmt.Println("Pilih Urutan : ")
	fmt.Println("1. Berdasarkan Nama Sparepart")
	fmt.Println("2. Berdasarkan Nama Pabrikan")
	fmt.Println("3. Berdasarkan Tahun Keluar")
	fmt.Print("Pilihan : ")
	fmt.Scanln(&pilih)
	fmt.Println("")
	for pilih != 1 && pilih != 2 && pilih != 3 {
		fmt.Println("Maaf Menu Tidak Tersedia")
		fmt.Println("")
		fmt.Println("Pilih Urutan : ")
		fmt.Println("1. Berdasarkan Nama Sparepart")
		fmt.Println("2. Berdasarkan Nama Pabrikan")
		fmt.Println("3. Berdasarkan Tahun Keluar")
		fmt.Println("4. Berdasarkan Harga")
		fmt.Print("Pilihan : ")
		fmt.Scanln(&pilih)
	}
	if pilih == 1 {
		dafnama(pilih)
	} else if pilih == 2 {
		dafpabrikan(pilih)
	} else if pilih == 3 {
		daftahun(pilih)
	}
}

func dafnama(pilih int) {
	var i int

	for i = 0; i < countsp; i++ {
		min := i 
		for j := i+1; j < countsp; j++ {
			if arrsparepart[j].nama < arrsparepart[min].nama {
				min = j
			}
		}
		t := arrsparepart[i].nama
		arrsparepart[i].nama = arrsparepart[min].nama
		arrsparepart[min].nama = t 
	}
	for i = 0 ; i < countsp ; i++ {
        fmt.Printf("%d.	%s\n" , i+1, arrsparepart[i].nama)
        fmt.Printf("	%s\n" , arrsparepart[i].pabrikan)
        fmt.Printf("	%d\n" , arrsparepart[i].tahun)
        fmt.Printf("	%d\n" , arrsparepart[i].harga)
    }
}

func dafpabrikan(pilih int) {
	var i, j int
	var temp string

	i = 0
	for i < countsp {
		temp = arrsparepart[i].pabrikan
		j = i
		for j > 0 && arrsparepart[j-1].pabrikan > temp {
			arrsparepart[j].pabrikan = arrsparepart[j-1].pabrikan
			j--
		}
		arrsparepart[j].pabrikan = temp
		i++
	}
	for i = 0 ; i < countsp ; i++ {
        fmt.Printf("%d.	%s\n" , i+1, arrsparepart[i].nama)
        fmt.Printf("	%s\n" , arrsparepart[i].pabrikan)
        fmt.Printf("	%d\n" , arrsparepart[i].tahun)
        fmt.Printf("	%d\n" , arrsparepart[i].harga)
    }
}

func daftahun(pilih int) {
	var i int
	for i = 0; i < countsp; i++ {
		min := i 
		for j := i+1; j < countsp; j++ {
			if arrsparepart[j].tahun < arrsparepart[min].tahun {
				min = j
			}
		}
		g := arrsparepart[i].tahun 
		arrsparepart[i].tahun = arrsparepart[min].tahun 
		arrsparepart[min].tahun = g
	}
	for i = 0; i < countsp; i++ {
		fmt.Printf("%d.	%s\n" , i+1, arrsparepart[i].nama)
        fmt.Printf("	%s\n" , arrsparepart[i].pabrikan)
        fmt.Printf("	%d\n" , arrsparepart[i].tahun)
        fmt.Printf("	%d\n" , arrsparepart[i].harga)
	}
}

func dafharga(pilih int){
	var i int

	for i = 0; i < countsp; i++ {
		min := i 
		for j := i+1; j < countsp; j++ {
			if arrsparepart[j].harga < arrsparepart[min].harga {
				min = j
			}
		}
		t := arrsparepart[i].harga
		arrsparepart[i].harga = arrsparepart[min].harga
		arrsparepart[min].harga = t 
	}
	for i = 0 ; i < countsp ; i++ {
        fmt.Printf("%d.	%s\n" , i+1, arrsparepart[i].nama)
        fmt.Printf("	%s\n" , arrsparepart[i].pabrikan)
        fmt.Printf("	%d\n" , arrsparepart[i].tahun)
        fmt.Printf("	%d\n" , arrsparepart[i].harga)
    }
}

func pemasukan(pilih int) {
	var bulan, tahun, valid, i, j int

	fmt.Print("Masukkan Bulan (mm): ")
	fmt.Scanln(&bulan)
	fmt.Print("Masukkan Taahun (yyyy): ")
	fmt.Scanln(&tahun)
	valid = (bulan*30) + (tahun*365)
	total := 0
	i = 0
	for i < countpel {
		j = 0
		if countbayar == 0 {
			fmt.Println("Maaf Tidak Ada Transaksi")
		} else {
			for j < countbayar {
				tglbayar := arrdata[i].arrtransaksi[j].bulan*30 + arrdata[i].arrtransaksi[j].tahun*365
				if valid == tglbayar {
					total = total + arrdata[i].arrtransaksi[j].bayar
				}
				j++
			}
			i++
		}
	}
	fmt.Println("Total pemasukan pada bulan", bulan, "tahun", tahun, "sebesar", total)
}
func main () {
	var(
		pilih int
		back string
	)
	countsp = 0
	countpel = 0
	countbayar = 0
	back = "Y"
	for back == "Y" || back == "y" {
		menu()
		fmt.Print("Pilihan : ")
		fmt.Scanln(&pilih)
		fmt.Println("")
		for pilih != 1 && pilih != 2 && pilih != 3 && pilih != 4 && pilih != 5 && pilih != 6 && pilih != 7{
			fmt.Println("Maaf menu tidak tersedia")
			fmt.Println("")
			menu()
			fmt.Print("Pilihan : ")
			fmt.Scanln(&pilih)
			}
		if pilih == 1 {
			tambah(pilih)
		} else if pilih == 2 {
			edit(pilih)
		} else if pilih == 3 {
			hapus(pilih)
		} else if pilih == 4 {
			caripelanggan(pilih)
		} else if pilih == 5 {
			caripembeli(pilih)
		} else if pilih == 6 {
			daftarpabrikan(pilih)
		} else if pilih == 7{
			pemasukan(pilih)
		}
		fmt.Println("")
		fmt.Print("Masukkan Y untuk kembali ke menu utama : ")
		fmt.Scanln(&back)
		fmt.Println("")
	}
	fmt.Println("Terima kasih telah menggunakan aplikasi service motor!")
}