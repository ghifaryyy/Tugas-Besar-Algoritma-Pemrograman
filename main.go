package main
import (
	"os"
	"fmt"
	"time"
	"os/exec"
	"runtime"
)

const MAXDATA int = 100

type polusi struct {
	lokasi, sumberPolusi, waktu, tingkatBahayaPolusi string
	indeksPolusi                                     float64
}
type tabPolusi [MAXDATA]polusi

func login() {
	var username, password string

	fmt.Println("=========================================")
	fmt.Println("|               LOGIN                  |")
	fmt.Println("=========================================")

	for {
		fmt.Print("Masukkan username: ")
		fmt.Scan(&username)
		fmt.Print("Masukkan Password: ")
		fmt.Scan(&password)

		if username == "admin" && password == "admin" {
			break
		}
		fmt.Println("\nLogin gagal! Username atau password salah.")
	}
	fmt.Println("Login berhasil!")
	clearScreen()
}
func menu() {
	/*
	   I.S.
	   F.S.
	*/
	fmt.Println("============================")
	fmt.Println("           MENU             ")
	fmt.Println("============================")
	fmt.Println("  1. Add Data               ")
	fmt.Println("  2. View Data              ")
	fmt.Println("  3. Edit Data              ")
	fmt.Println("  4. Delete Data            ")
	fmt.Println("  5. Search Data            ")
	fmt.Println("  6. Exit                   ")
	fmt.Println("----------------------------")
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
		T[*N].tingkatBahayaPolusi = "Kualitas udara tidak sehat untuk kelompok orang yang sensitif"
	} else if T[*N].indeksPolusi <= 200 {
		T[*N].tingkatBahayaPolusi = "Kualitas udara tidak sehat"
	} else if T[*N].indeksPolusi <= 300 {
		T[*N].tingkatBahayaPolusi = "Kualitas udara sangat tidak sehat"
	} else {
		T[*N].tingkatBahayaPolusi = "Kualitas udara berbahaya"
	}
	*N++
}
func viewData(A tabPolusi, N int) {
	var memilih, metode int
	fmt.Println("1. Ascending Sorting")
	fmt.Println("2. Descending Sorting")
	fmt.Println("3. Exit")
	fmt.Print("Pilih menu [1 - 3]: ")
	fmt.Scan(&memilih)

	for memilih != 3 {
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
				selectionSortIndeksPolusiAsc(&A, N)
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
				selectionSortIndeksPolusiDsc(&A, N)
				tampilData(A, N)
			} else if metode == 2 {
				selectionSortWaktuDsc(&A, N)
				tampilData(A, N)
			}
		}
	}
}
func editData(T *tabPolusi, N int) {
	/*
	   I.S.
	   F.S.
	*/
	var edit int
	fmt.Println("Ketik 0 jika ingin kembali ke menu")
	fmt.Print("Edit data (1 - 100): ")
	fmt.Scan(&edit) // indeks yang mau diedit
	if edit != 0 {
		edit--
		addData(T, &edit)
		fmt.Println("Data Berhasil diedit.")
	}
}
func deleteData(T *tabPolusi, N *int, lok int) {
	/*
	   I.S.
	   F.S.
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
	}
}
func searchData(T tabPolusi, N int, cari string) {
	/*
	   I.S.
	   F.S.
	*/
	var i, idx, milih int
	var lokasi, waktu string
	fmt.Println("1. Location")
	fmt.Println("2. Time")
	fmt.Println("3. Exit")
	fmt.Print("Cari data [1 - 3]: ")
	fmt.Scan(&milih)
	for milih != 3 {
		if milih == 1 {
			fmt.Println("Ketik n jika ingin kembali ke menu.")
			fmt.Print("Masukkan lokasi: ")
			fmt.Scan(&lokasi)
			if lokasi == "n" {
				searchData(T, N, cari)
				return
			} else {
				for i = 0; i < N; i++ {
					if T[i].lokasi == lokasi {
						idx = i
						fmt.Printf(" %-3s  %-10s  %-12s  %-13s  %-20s  %-61s \n", "No", "Lokasi", "Waktu", "Indeks Polusi", "Sumber Polusi", "Tingkat Bahaya Polusi")
						fmt.Printf(" %-3d  %-10s  %-12s  %-13.1f  %-20s  %-61s \n", i+1, T[idx].lokasi, T[idx].waktu, T[idx].indeksPolusi, T[idx].sumberPolusi, T[idx].tingkatBahayaPolusi)
					}
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
				for i = 0; i < N; i++ {
					if T[i].waktu == waktu {
						idx = i
						fmt.Printf(" %-3s  %-10s  %-12s  %-13s  %-20s  %-61s \n", "No", "Lokasi", "Waktu", "Indeks Polusi", "Sumber Polusi", "Tingkat Bahaya Polusi")
						fmt.Printf(" %-3d  %-10s  %-12s  %-13.1f  %-20s  %-61s \n", i+1, T[idx].lokasi, T[idx].waktu, T[idx].indeksPolusi, T[idx].sumberPolusi, T[idx].tingkatBahayaPolusi)
					}
				}
			}
		}
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
func selectionSortIndeksPolusiAsc(A *tabPolusi, N int) {
	var i, idx, pass int
	var temp polusi
	
	pass = 1
	for pass < N {
		idx = pass - 1
		i = pass
		for i < N {
			if A[i].indeksPolusi < A[idx].indeksPolusi {
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
func selectionSortWaktuAsc(A *tabPolusi, N int) {
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
func selectionSortIndeksPolusiDsc(A *tabPolusi, N int) {
	var i, idx, pass int
	var temp polusi
	
	pass = 1
	for pass < N {
		idx = pass - 1
		i = pass
		for i < N {
			if A[i].indeksPolusi > A[idx].indeksPolusi {
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
func clearScreen() {
	// Mengecek sistem operasi
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
		I.S.: Program belum menerima input menu dari pengguna.
		F.S.: Menu dipilih dan prosedur yang sesuai dijalankan berdasarkan input pengguna.
	*/
	var pilihan int
	var A tabPolusi
	var n int
	var x1 int
	var x2 string
	login()
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
			tampilData(A, n)
			searchData(A, n, x2)

		}
	}
}