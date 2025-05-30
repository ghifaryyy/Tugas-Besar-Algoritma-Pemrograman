# Tugas-Besar-Algoritma-Pemrograman (Kelompok burcang)
Aplikasi Manajemen dan Pemantauan Polusi Udara Lokal 

Aplikasi ini dirancang sebagai alat bantu yang memungkinkan pengguna untuk mencatat, memantau, dan mengelola data tingkat polusi udara di berbagai wilayah. Fokus utama aplikasi adalah pendataan dan pemantauan kualitas udara berdasarkan beberapa parameter penting, yakni: 
1. Lokasi pemantauan 
2. Waktu pencatatan 
3. Indeks polusi udara (Air Quality Index / AQI) 
4. Sumber penyebab polusi 
5. Tingkat bahaya dari polusi tersebut 

Aplikasi ini ditujukan bagi individu maupun komunitas yang memiliki kepedulian terhadap lingkungan, khususnya terhadap kualitas udara. Dengan menggunakan aplikasi ini, pengguna dapat secara aktif menambahkan data pemantauan kualitas udara, serta melihat dan mengelola data yang telah mereka catat. Adapun fitur-fitur yang dikembangkan pada aplikasi ini, yaitu: 

Fitur Login 
 Pengguna diharuskan untuk melakukan proses login terlebih dahulu sebelum dapat mengakses menu utama. Ini bertujuan untuk memastikan penggunaan aplikasi yang lebih aman dan terstruktur. 

Fitur Menu Utama 
 Setelah login berhasil, aplikasi menampilkan menu utama yang berisi daftar fitur yang dapat diakses. Pengguna dapat memilih fitur sesuai kebutuhan, seperti menambahkan data, melihat data, mengedit, atau menghapus data. 

Fitur Tambah Data 
 Pengguna dapat menambahkan data pemantauan kualitas udara dengan mengisi beberapa informasi, yaitu: 
1. Lokasi pemantauan 
2. Waktu pengukuran (format: dd-mm-yyyy) 
3. Indeks polusi (AQI) 
4. Sumber polusi (misalnya: kendaraan bermotor, industri, pembakaran sampah) 
5. Berdasarkan nilai AQI yang diinput, aplikasi secara otomatis mengkategorikan tingkat bahaya polusi, seperti: 
"Kualitas udara baik" 
"Tidak sehat untuk kelompok sensitif" 
"Tidak sehat" 
"Sangat tidak sehat" 
"Berbahaya"

Fitur Tampilkan Data.
Data yang telah ditambahkan dapat ditampilkan kembali dalam bentuk tabel. Selain itu, pengguna juga dapat mengurutkan data: 
Berdasarkan waktu atau lokasi dalam urutan menaik (ascending) atau menurun (descending). Untuk proses pengurutan, aplikasi menggunakan algoritma selection sort dan juga insertion sort. 

Fitur Ubah Data (Edit) 
 Pengguna diberikan opsi untuk mengubah data yang telah ditambahkan sebelumnya. Perubahan dapat dilakukan pada semua atribut data (lokasi, waktu, indeks, sumber). 

Fitur Hapus Data 
 Pengguna dapat menghapus entri data tertentu berdasarkan pencarian nilai tertentu. Proses pencarian data yang akan dihapus menggunakan algoritma Sequential Search, yang mencocokkan data satu per satu hingga ditemukan. 
