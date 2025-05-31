package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

const MAXDATA int = 100

type polusi struct {
	lokasi, sumberPolusi, waktu, tingkatBahayaPolusi string
	indeksPolusi                                     float64
}
type tabPolusi [MAXDATA]polusi

func login() {
	/*
	I.S. Program sedang dijalankan, pengguna belum login
	F.S. Pengguna berhasil login (atau gagal dan diminta login ulang)
	*/
	var username, password string
	var login bool

	login = false
	for !login {
	fmt.Println("=========================================")
	fmt.Println("|               LOGIN                  |")
	fmt.Println("=========================================")
	fmt.Print("Masukkan username: ")
	fmt.Scan(&username)
	fmt.Print("Masukkan Password: ")
	fmt.Scan(&password)

		if username == "admin" && password == "admin" {
		login = true
		} else {
		fmt.Println("\nLogin gagal! Username atau password salah.")
		}
	}
	clearScreen()
}
func menu() { 
	/*
	I.S. Pengguna sudah login dan masuk ke program utama
   	F.S. Program menampilkan pilihan menu dan membaca input pilihan pengguna
	*/
	fmt.Println("\n\n================================================================")
	fmt.Println("                            MENU							 	 ")
	fmt.Println("================================================================")
	fmt.Println("1.  Tambah Data")
	fmt.Println("2.  Lihat Data")
	fmt.Println("3.  Edit Data")
	fmt.Println("4.  Hapus Data")
	fmt.Println("5.  Cari Data")
	fmt.Println("6.  Keluar")
	fmt.Println("----------------------------------------------------------------")
	fmt.Print("Pilih menu [1 - 6]: ")
}
func addData(T *tabPolusi, N *int) {
	/*
	   I.S. Terdefinisi array T dan n sembarang nilai.
	   F.S. Array T terisi sebanyak N
	*/
	fmt.Print("Masukkan lokasi: ")
	fmt.Scan(&T[*N].lokasi)
	fmt.Print("Masukkan Waktu (dd-mm-yyyy): ")
	fmt.Scan(&T[*N].waktu)
	fmt.Print("Masukkan indeks polusi: ")
	fmt.Scan(&T[*N].indeksPolusi)
	fmt.Print("Masukkan sumber polusi: ")
	fmt.Scan(&T[*N].sumberPolusi)

	if T[*N].indeksPolusi <= 50 {
		T[*N].tingkatBahayaPolusi = "Kualitas udara baik."
	} else if T[*N].indeksPolusi <= 100 {
		T[*N].tingkatBahayaPolusi = "Kualits udara sedang"
	} else if T[*N].indeksPolusi <= 150 {
		T[*N].tingkatBahayaPolusi = "Kualitas udara tidak sehat untuk kelompok orang yang sensitif"
	} else if T[*N].indeksPolusi <= 200 {
		T[*N].tingkatBahayaPolusi = "Kualitas udara tidak sehat"
	} else if T[*N].indeksPolusi <= 300 {
		T[*N].tingkatBahayaPolusi = "Kualitas udara sangat tidak sehat"
	} else {
		T[*N].tingkatBahayaPolusi = "Kualitas udara berbahaya"
	}
}
func viewData(A tabPolusi, N int) {
	/*
	I.S. Data polusi sejumlah N tersedia dalam array
	F.S. Data ditampilkan ke layar sesuai urutan yang dipilih (berdasarkan waktu
	     atau indeks polusi, ascending/descending)
	*/	
	var memilih, metode int
	fmt.Println("1. Ascending Sorting")
	fmt.Println("2. Descending Sorting")
	fmt.Println("3. Exit")
	fmt.Print("Pilih menu [1 - 3]: ")
	fmt.Scan(&memilih)

		if memilih == 1 {
			fmt.Println("1. Indeks Polusi")
			fmt.Println("2. Waktu")
			fmt.Println("3. Back")
			fmt.Print("Pilih jenis [1 - 3]: ")
			fmt.Scan(&metode)
			
			if metode == 3 {
				viewData(A, N)
				return
			} else if metode == 1 {
				insertionSortIndeksPolusiAsc(&A, N)
				tampilData(A, N)
			} else if metode == 2 {
				selectionSortWaktuAsc(&A, N)
				tampilData(A, N)
			}
		} else if memilih == 2 {
			fmt.Println("1. Indeks Polusi")
			fmt.Println("2. Waktu")		
			fmt.Println("3. Back")		
			fmt.Print("Pilih jenis [1 - 3]: ")
			fmt.Scan(&metode)

			if metode == 3 {
				viewData(A, N)
				return
			} else if metode == 1 {
				insertionSortIndeksPolusiDsc(&A, N)
				tampilData(A, N)
			} else if metode == 2 {
				selectionSortWaktuDsc(&A, N)
				tampilData(A, N)
			}
		} else if memilih == 3 {
			clearScreen()
			return
		} else {
			clearScreen()
			tampilData(A, N)
			viewData(A, N)
		}
}
func editData(T *tabPolusi, N int) {
	/*
	I.S. Data polusi sebanyak N di dalam array
	F.S. Salah satu data di dalam array telah diperbarui
	*/
	var edit int
	fmt.Println("Ketik 0 jika ingin kembali ke menu")
	fmt.Print("Edit data (1 - 100): ")
	fmt.Scan(&edit) // indeks yang mau diedit
	if edit == 0 {
		clearScreen()
		return
	} else {
			edit--
			addData(T, &edit)
			fmt.Println("Data Berhasil diedit.")
			clearScreen()
	}
}
func deleteData(T *tabPolusi, N *int, lok int) {
	/*
	I.S. Data polusi sebanyak N ada di array, dan indeks data yang ingin dihapus adalah lok
	F.S. Data di indeks lok dihapus, data digeser, dan N berkurang 1
	*/
	var i, idx int
	fmt.Println("Ketik 0 jika ingin kembali ke menu. ")
	fmt.Print("Hapus data [1 - 100]: ")
	fmt.Scan(&lok)
	if lok != 0 {
		idx = -1
		for i = 0; i < *N; i++ {
			if i == lok-1 {
				idx = i
			}
		}
		for i = idx; i < *N-1; i++ {
			T[i] = T[i+1]
		}
		*N--
		clearScreen()
	} else {
		clearScreen()
		return
	}
}
func searchData(T tabPolusi, N int, cari string) {	
	/*
	I.S. Data sebanyak N tersedia dalam array, nilai string cari telah dimasukkan
	F.S. Program menampilkan hasil pencarian berdasarkan cari (misal nama atau waktu) 
	     dan menampilkan nilai maks atau min berdasarkan indeks polusi
	*/

	var found bool
	var i, idx, milih int
	var lokasi, waktu string
	milih = 0
	fmt.Println("1. Lokasi")
	fmt.Println("2. Waktu")
	fmt.Println("3. Nilai tertinggi berdasarkan indeks polusi")
	fmt.Println("4. Nilai terendah berdasarkan indeks polusi ")
	fmt.Println("5. Exit ")
	fmt.Print("Cari data [1 - 5]: ")
	fmt.Scan(&milih)
	if milih == 1 {
		fmt.Println("Ketik n jika ingin kembali ke menu.")
		fmt.Print("Masukkan lokasi: ")
		fmt.Scan(&lokasi)
		if lokasi == "n" {
			searchData(T, N, cari)
			return
			} else {
				cari = lokasi
				insertionSortLokasiAsc(&T, N)
				idx = binSearch(&T, N, cari)
				if idx == -1 {
					fmt.Println("|===========================================|")
					fmt.Println("|Data dengan lokasi tersebut tidak ditemukan|")
					fmt.Println("|===========================================|")
					} else {
					fmt.Println("==========================================================================================================================================")
					fmt.Printf("| %-3s | %-10s | %-12s | %-13s | %-20s | %-61s |\n", "No", "Lokasi", "Waktu", "Indeks Polusi", "Sumber Polusi", "Tingkat Bahaya Polusi")
					fmt.Println("==========================================================================================================================================")
					fmt.Printf("| %-3d | %-10s | %-12s | %-13.1f | %-20s | %-61s |\n", idx+1, T[idx].lokasi, T[idx].waktu, T[idx].indeksPolusi, T[idx].sumberPolusi, T[idx].tingkatBahayaPolusi)
					fmt.Println("==========================================================================================================================================")
				}
			}
		} else if milih == 2 {
			fmt.Println("Ketik n jika ingin kembali ke menu.")
			fmt.Print("Masukkan waktu (dd-mm-yyyy): ")
			fmt.Scan(&waktu)
			if waktu == "n" {
				searchData(T, N, cari)
				return
			} else {
				found = false
				for i = 0; i < N; i++ {
					if T[i].waktu == waktu {
						idx = i
						fmt.Println("==========================================================================================================================================")
						fmt.Printf("| %-3s | %-10s | %-12s | %-13s | %-20s | %-61s |\n", "No", "Lokasi", "Waktu", "Indeks Polusi", "Sumber Polusi", "Tingkat Bahaya Polusi")
						fmt.Println("==========================================================================================================================================")
						fmt.Printf("| %-3d | %-10s | %-12s | %-13.1f | %-20s | %-61s |\n", i+1, T[idx].lokasi, T[idx].waktu, T[idx].indeksPolusi, T[idx].sumberPolusi, T[idx].tingkatBahayaPolusi)
						fmt.Println("==========================================================================================================================================")
						found = true
					}
				}
			if !found{
				fmt.Println("|==========================================|")
				fmt.Println("|Data dengan waktu tersebut tidak ditemukan|")
				fmt.Println("|==========================================|")
			}
		}
		} else if milih == 3 {
			idx = findMax(T, N)
			fmt.Println("==========================================================================================================================================")
			fmt.Printf("| %-3s | %-10s | %-12s | %-13s | %-20s | %-61s |\n", "No", "Lokasi", "Waktu", "Indeks Polusi", "Sumber Polusi", "Tingkat Bahaya Polusi")
			fmt.Println("==========================================================================================================================================")
			fmt.Printf("| %-3d | %-10s | %-12s | %-13.1f | %-20s | %-61s |\n", idx+1, T[idx].lokasi, T[idx].waktu, T[idx].indeksPolusi, T[idx].sumberPolusi, T[idx].tingkatBahayaPolusi)
			fmt.Println("==========================================================================================================================================")
		} else if milih == 4 {
			idx = findMin(T, N)
			fmt.Println("==========================================================================================================================================")
			fmt.Printf("| %-3s | %-10s | %-12s | %-13s | %-20s | %-61s |\n", "No", "Lokasi", "Waktu", "Indeks Polusi", "Sumber Polusi", "Tingkat Bahaya Polusi")
			fmt.Println("==========================================================================================================================================")
			fmt.Printf("| %-3d | %-10s | %-12s | %-13.1f | %-20s | %-61s |\n", idx+1, T[idx].lokasi, T[idx].waktu, T[idx].indeksPolusi, T[idx].sumberPolusi, T[idx].tingkatBahayaPolusi)
			fmt.Println("==========================================================================================================================================")
		} else if milih == 5 {
			clearScreen()
			return
		} else {
			clearScreen()
			tampilData(T, N)
			searchData(T, N, cari)
		}

}
func tampilData(T tabPolusi, N int) {
	/*
	I.S. Diterima array T dengan banyak N
	F.S. Mencetak array T sebanyak N
	*/
	var j int
	fmt.Println("==========================================================================================================================================")
	fmt.Printf("| %-3s | %-10s | %-12s | %-13s | %-20s | %-61s |\n", "No", "Lokasi", "Waktu", "Indeks Polusi", "Sumber Polusi", "Tingkat Bahaya Polusi")
	fmt.Println("==========================================================================================================================================")
	for j = 0; j < N; j++ {
		fmt.Printf("| %-3d | %-10s | %-12s | %-13.1f | %-20s | %-61s |\n", j+1, T[j].lokasi, T[j].waktu, T[j].indeksPolusi, T[j].sumberPolusi, T[j].tingkatBahayaPolusi)
	}
	fmt.Println("==========================================================================================================================================")
}
func insertionSortIndeksPolusiAsc(A *tabPolusi, N int) {
	/*
	I.S. Data belum terurut berdasarkan indeks polusi (ascending)
	F.S. Data terurut naik berdasarkan indeks polusi
	*/
	var i, pass int
	var temp polusi
	
	pass = 1
	for pass < N {
		i = pass
		temp = A[pass]
		for i > 0 && temp.indeksPolusi < A[i-1].indeksPolusi {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
		pass++
	}
}
func selectionSortWaktuAsc(A *tabPolusi, N int) {
	/*
	I.S. Data belum terurut berdasarkan waktu (ascending)
	F.S. Data terurut naik berdasarkan waktu
	*/
	var pass, i, idx int
	var temp polusi
	var tgl1, tgl2 int
	var dd1, mm1, yy1, dd2, mm2, yy2 int

	pass = 1
	for pass < N {
		idx = pass - 1
		i = pass
		for i < N {
			dd1 = int(A[i].waktu[0]-'0')*10 + int(A[i].waktu[1]-'0')
			mm1 = int(A[i].waktu[3]-'0')*10 + int(A[i].waktu[4]-'0')
			yy1 = int(A[i].waktu[6]-'0')*1000 + int(A[i].waktu[7]-'0')*100 + int(A[i].waktu[8]-'0')*10 + int(A[i].waktu[9]-'0')
			tgl1 = yy1 *10000 + mm1*100 + dd1

			dd2 = int(A[idx].waktu[0]-'0')*10 + int(A[idx].waktu[1]-'0')
			mm2 = int(A[idx].waktu[3]-'0')*10 + int(A[idx].waktu[4]-'0')
			yy2 = int(A[idx].waktu[6]-'0')*1000 + int(A[idx].waktu[7]-'0')*100 + int(A[idx].waktu[8]-'0')*10 + int(A[idx].waktu[9]-'0')
			tgl2 = yy2*10000 + mm2*100 + dd2

			if tgl1 < tgl2 {
				idx = i
			}
			i++
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass++
	}
}
func selectionSortWaktuDsc(A *tabPolusi, N int) {
	/*
	I.S. Data belum terurut berdasarkan waktu (descending)
	F.S. Data terurut turun berdasarkan waktu
	*/
	var pass, i, idx int
	var temp polusi
	var tgl1, tgl2 int
	var dd1, mm1, yy1, dd2, mm2, yy2 int

	pass = 1
	for pass < N {
		idx = pass - 1
		i = pass
		for i < N {
			dd1 = int(A[i].waktu[0]-'0')*10 + int(A[i].waktu[1]-'0')
			mm1 = int(A[i].waktu[3]-'0')*10 + int(A[i].waktu[4]-'0')
			yy1 = int(A[i].waktu[6]-'0')*1000 + int(A[i].waktu[7]-'0')*100 + int(A[i].waktu[8]-'0')*10 + int(A[i].waktu[9]-'0')
			tgl1 = yy1 *10000 + mm1*100 + dd1

			dd2 = int(A[idx].waktu[0]-'0')*10 + int(A[idx].waktu[1]-'0')
			mm2 = int(A[idx].waktu[3]-'0')*10 + int(A[idx].waktu[4]-'0')
			yy2 = int(A[idx].waktu[6]-'0')*1000 + int(A[idx].waktu[7]-'0')*100 + int(A[idx].waktu[8]-'0')*10 + int(A[idx].waktu[9]-'0')
			tgl2 = yy2*10000 + mm2*100 + dd2

			if tgl1 > tgl2 {
				idx = i
			}
			i++
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass++
	}

}
func insertionSortIndeksPolusiDsc(A *tabPolusi, N int) {
	/*
	I.S. Data belum terurut berdasarkan indeks polusi (descending)
	F.S. Data terurut turun berdasarkan indeks polusi
	*/
	var i, pass int
	var temp polusi
	
	pass = 1
	for pass < N {
		i = pass
		temp = A[pass]
		for i > 0 && temp.indeksPolusi > A[i-1].indeksPolusi {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
		pass++
	}
}
func findMax(T tabPolusi, n int) int{
	/*
	{Diterima array T yang berisi n bilangan bulat untuk mengembalikan nilai minimum dari array T}
	*/
	var i, max int

	max = 0
	for i = 1; i < n; i++ {
		if T[i].indeksPolusi > T[max].indeksPolusi {
			max = i
		}
	}
	return max
}
func findMin(T tabPolusi, n int) int {
	/*
	{Diterima array T yang berisi n bilangan bulat untuk mengembalikan nilai maksimum dari array T}
	*/
	var i, min int

	min = 0
	for i = 1; i < n; i++ {
		if T[i].indeksPolusi < T[min].indeksPolusi {
			min = i
		}
	}
	return min
}
func insertionSortLokasiAsc(T *tabPolusi, N int) {
	/*
	I.S. Data belum terurut berdasarkan lokasi (ascending)
	F.S. Data terurut naik berdasarkan lokasi
	*/
	var i, pass int
	var temp polusi

	pass = 1
	for pass < N {
		i = pass
		temp = T[pass]
		for i > 0 && temp.lokasi < T[i-1].lokasi {
			T[i] = T[i-1]
			i--
		}
		T[i] = temp
		pass++
	}
}
func binSearch(T *tabPolusi, n int, x string) int {
 	/*
	{mengembalikan indeks pencarian apabila x ada di dalam array T yang berisi 
		n bilangan atau -1 apabila tidak ditemukan, T terurut membesar atau ASCENDING}

 	*/
	var left, right, mid int
	var idx int

	idx = -1
	left = 0
	right = n - 1
	for left <= right && idx == -1 {
		mid = (left + right) / 2 
		if x <  T[mid].lokasi {
			right = mid - 1
		} else if x > T[mid].lokasi {
			left = mid + 1
		} else if x == T[mid].lokasi {
			idx = mid
		}
	}
	return idx
}
func clearScreen() {
	/*
	I.S. Layar berisi teks atau tampilan dari proses sebelumnya
	F.S. Layar menjadi kosong (bersih dari tampilan sebelumnya)
	*/
	time.Sleep(1 * time.Second)

	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls") // untuk Windows
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear") // untuk Linux dan Mac
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
func main() {
	/*
	I.S. Program belum menerima input menu dari pengguna.
	F.S. Menu dipilih dan prosedur yang sesuai dijalankan berdasarkan input pengguna.
	*/
	var pilihan int
	var A tabPolusi
	var n int
	var x1 int
	var x2 string
	login()
	fmt.Println("================================================================")
	fmt.Println("       APLIKASI MANAJEMEN & PEMANTAUAN POLUSI UDARA LOKAL")
	fmt.Println("================================================================")
	fmt.Println("                      Login berhasil.")
	fmt.Println("Selamat datang di sistem pantau dan kelola data kualitas udara.")
	fmt.Println("		Waktu akses :", time.Now().Format("02 January 2006 15:04:05"))
	for pilihan != 6 {
		menu()
		fmt.Scan(&pilihan)
		clearScreen()

		if pilihan == 1 {
			fmt.Println("+---------------------------+")
			fmt.Println("|      ADD DATA MENU        |")
			fmt.Println("+---------------------------+")
			addData(&A, &n)
			fmt.Println("DATA BERHASIL DITAMBAHKAN.")

		} else if pilihan == 2 {
			fmt.Println("+---------------------------+")
			fmt.Println("|      ViEW DATA MENU       |")
			fmt.Println("+---------------------------+")
			tampilData(A, n)
			viewData(A, n)

		} else if pilihan == 3 {
			fmt.Println("+---------------------------+")
			fmt.Println("|      EDIT DATA MENU       |")
			fmt.Println("+---------------------------+")
			tampilData(A, n)
			editData(&A, n)

		} else if pilihan == 4 {
			fmt.Println("+---------------------------+")
			fmt.Println("|      DELETE DATA MENU     |")
			fmt.Println("+---------------------------+")
			tampilData(A, n)
			deleteData(&A, &n, x1)

		} else if pilihan == 5 {
			fmt.Println("+---------------------------+")
			fmt.Println("|     SEARCH DATA MENU      |")
			fmt.Println("+---------------------------+")
			insertionSortLokasiAsc(&A, n)
			tampilData(A, n)
			searchData(A, n, x2)
		} else if pilihan == 6 {
			clearScreen()
			return
		}
	}
}
