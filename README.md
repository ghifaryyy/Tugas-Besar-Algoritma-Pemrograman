# Aplikasi Manajemen dan Pemantauan Polusi Udara Lokal

Aplikasi ini dirancang sebagai alat bantu yang memungkinkan pengguna untuk mencatat, memantau, dan mengelola data tingkat polusi udara di berbagai wilayah. Fokus utama aplikasi adalah pendataan dan pemantauan kualitas udara berdasarkan beberapa parameter penting, yakni:
1.  Lokasi pemantauan
2.  Waktu pencatatan
3.  Indeks polusi udara (Air Quality Index / AQI)
4.  Sumber penyebab polusi
5.  Tingkat bahaya dari polusi tersebut

Aplikasi ini ditujukan bagi individu maupun komunitas yang memiliki kepedulian terhadap lingkungan, khususnya terhadap kualitas udara. Dengan menggunakan aplikasi ini, pengguna dapat secara aktif menambahkan data pemantauan kualitas udara, serta melihat dan mengelola data yang telah mereka catat. Adapun fitur-fitur yang dikembangkan pada aplikasi ini, yaitu:

## Fitur Login

Pengguna diharuskan untuk melakukan proses login terlebih dahulu sebelum dapat mengakses menu utama. Ini bertujuan untuk memastikan penggunaan aplikasi yang lebih aman dan terstruktur.

## Fitur Menu Utama

Setelah login berhasil, aplikasi menampilkan menu utama yang berisi daftar fitur yang dapat diakses. Pengguna dapat memilih fitur sesuai kebutuhan, seperti menambahkan data, melihat data, mengedit, atau menghapus data.

## Fitur Tambah Data

Pengguna dapat menambahkan data pemantauan kualitas udara dengan mengisi beberapa informasi, yaitu:
1.  Lokasi pemantauan
2.  Waktu pengukuran (format: `dd-mm-yyyy`)
3.  Indeks polusi (AQI)
4.  Sumber polusi (misalnya: kendaraan bermotor, industri, pembakaran sampah)
5.  Berdasarkan nilai AQI yang diinput, aplikasi secara otomatis mengkategorikan tingkat bahaya polusi, seperti:
    * "Kualitas udara baik"
    * "Tidak sehat untuk kelompok sensitif"
    * "Tidak sehat"
    * "Sangat tidak sehat"
    * "Berbahaya"

## Fitur Tampilkan Data.

Data yang telah ditambahkan dapat ditampilkan kembali dalam bentuk tabel. Selain itu, pengguna juga dapat mengurutkan data:

Berdasarkan **waktu** atau **lokasi** dalam urutan menaik (ascending) atau menurun (descending). Untuk proses pengurutan, aplikasi menggunakan algoritma **selection sort** dan juga **insertion sort**.

## Fitur Ubah Data (Edit)

Pengguna diberikan opsi untuk mengubah data yang telah ditambahkan sebelumnya. Perubahan dapat dilakukan pada semua atribut data (lokasi, waktu, indeks, sumber).

## Fitur Hapus Data

Pengguna dapat menghapus entri data tertentu berdasarkan pencarian nilai tertentu. Proses pencarian data yang akan dihapus menggunakan algoritma **Sequential Search**, yang mencocokkan data satu per satu hingga ditemukan.

## Fitur Cari Data

Melalui fitur ini, pengguna dapat:

* **Melakukan pencarian data spesifik** berdasarkan kriteria berikut:
    * **Lokasi:** Masukkan nama lokasi yang dicari.
    * **Waktu:** Masukkan tanggal (dengan format `dd-mm-yyyy`) atau rentang waktu tertentu.
    * **Sumber Polusi:** Masukkan kata kunci sumber polusi yang dicari (jika sumber yang dicari menggunakan garis bawah, sertakan dalam pencarian, misal: `Kendaraan_Bermotor`).
    * Masukkan kata kunci pencarian pada kolom yang sesuai dan tekan "Cari" atau Enter. Hasil pencarian yang relevan akan ditampilkan.
* **Melihat Indeks Polusi Tertinggi atau Terendah:**
    * Pilih opsi untuk "Lihat Indeks Polusi Tertinggi" atau "Lihat Indeks Polusi Terendah" yang tersedia di dalam fitur "Cari Data".
    * Sistem akan secara otomatis menganalisis seluruh data dan menampilkan informasi mengenai:
        * **Indeks Polusi Tertinggi:** Menampilkan data dengan nilai indeks polusi paling tinggi beserta detail lokasi dan waktunya.
        * **Indeks Polusi Terendah:** Menampilkan data dengan nilai indeks polusi paling rendah beserta detail lokasi dan waktunya.

# Deskripsi Penggunaan Aplikasi Pemantau Polusi

Aplikasi ini dirancang untuk membantu pengguna dalam mengelola dan menganalisis data polusi melalui serangkaian menu yang terstruktur. Berikut adalah panduan langkah demi langkah penggunaan aplikasi:

## 1. Login Pengguna

-   Untuk memulai, pengguna harus masuk ke dalam sistem.
-   Masukkan **Username:** `admin`
-   Masukkan **Password:** `admin`
-   Klik tombol "Login" atau tekan Enter.

## 2. Tampilan Menu Utama

-   Setelah berhasil login, pengguna akan diarahkan ke tampilan menu utama.
-   Dari sini, pengguna dapat memilih berbagai menu fungsional:
    * Tambah Data
    * Lihat Data
    * Ubah Data
    * Hapus Data
    * Mencari Data

## 3. Menu: Tambah Data

-   Pilih opsi **"Tambah Data"** pada menu utama.
-   Isi formulir tambah data dengan informasi berikut:
    * **Lokasi:** Masukkan nama atau deskripsi lokasi tempat data polusi diambil.
    * **Waktu:** Masukkan tanggal pengambilan data dengan format `dd-mm-yyyy` (contoh: `30-05-2025`).
    * **Indeks Polusi:** Masukkan nilai numerik indeks polusi yang terukur.
    * **Sumber Polusi:** Masukkan deskripsi sumber utama polusi. **Jika sumber polusi terdiri lebih dari satu kata, gunakan garis bawah (`_`) sebagai pengganti spasi** (contoh: `Kendaraan_Bermotor`, `Industri_Tekstil`, `Pembakaran_Sampah_Terbuka`).
-   Klik tombol "Simpan" atau "Tambah" untuk menyimpan data baru ke dalam sistem.

## 4. Menu: Lihat Data

-   Pilih opsi **"Lihat Data"** pada menu utama.
-   Seluruh data polusi yang tersimpan akan ditampilkan, kemungkinan dengan nomor urut untuk setiap entri data.
-   Pengguna dapat mengurutkan data yang ditampilkan berdasarkan:
    * **Indeks Polusi:** Urutkan secara menaik (Ascending, nilai terendah ke tertinggi) atau menurun (Descending, nilai tertinggi ke terendah).
    * **Waktu:** Urutkan secara menaik (Ascending, data terlama ke terbaru) atau menurun (Descending, data terbaru ke terlama).
-   Pilih kolom yang diinginkan dan opsi pengurutan (Ascending/Descending) untuk menampilkan data sesuai preferensi.

## 5. Menu: Ubah Data

-   Pilih opsi **"Ubah Data"** pada menu utama.
-   Pengguna akan diminta untuk memilih data spesifik yang ingin diubah (misalnya dengan memasukkan nomor urut data atau melalui mekanisme pencarian singkat terlebih dahulu, lalu memilihnya).
-   Setelah data dipilih, formulir akan muncul dengan data yang sudah terisi sebelumnya.
-   Lakukan perubahan yang diperlukan pada kolom lokasi, waktu, indeks polusi, atau sumber polusi. **Ingat, jika mengubah Sumber Polusi menjadi lebih dari satu kata, gunakan garis bawah (`_`)** (contoh: `Asap_Pabrik`).
-   Klik tombol "Simpan Perubahan" atau "Update" untuk menyimpan modifikasi.

## 6. Menu: Hapus Data

-   Pilih opsi **"Hapus Data"** pada menu utama.
-   Aplikasi akan meminta pengguna untuk memasukkan **nomor urutan data** yang ingin dihapus (misalnya, jika ada 100 data, masukkan angka antara 1 sampai 100).
-   Setelah memasukkan nomor urutan, akan muncul konfirmasi untuk memastikan pengguna benar-benar ingin menghapus data tersebut.
-   Klik "Ya" atau "Hapus" untuk mengkonfirmasi dan menghapus data secara permanen dari sistem.

## 7. Menu: Cari Data

-   Pilih opsi **"Cari Data"** pada menu utama.
-   Di dalam menu ini, pengguna dapat:
    * **Melakukan pencarian data spesifik** berdasarkan kriteria berikut:
        * **Lokasi:** Masukkan nama lokasi yang dicari.
        * **Waktu:** Masukkan tanggal (dengan format `dd-mm-yyyy`) atau rentang waktu tertentu.
        * **Sumber Polusi:** Masukkan kata kunci sumber polusi yang dicari (jika sumber yang dicari menggunakan garis bawah, sertakan dalam pencarian, misal: `Kendaraan_Bermotor`).
        * Masukkan kata kunci pencarian pada kolom yang sesuai dan tekan "Cari" atau Enter. Hasil pencarian yang relevan akan ditampilkan.
    * **Melihat Indeks Polusi Tertinggi atau Terendah:**
        * Pilih opsi untuk "Lihat Indeks Polusi Tertinggi" atau "Lihat Indeks Polusi Terendah" yang tersedia di dalam menu "Cari Data".
        * Sistem akan secara otomatis menganalisis seluruh data dan menampilkan informasi mengenai:
            * **Indeks Polusi Tertinggi:** Menampilkan data dengan nilai indeks polusi paling tinggi beserta detail lokasi dan waktunya.
            * **Indeks Polusi Terendah:** Menampilkan data dengan nilai indeks polusi paling rendah beserta detail lokasi dan waktunya.

Dengan mengikuti panduan ini, pengguna dapat memanfaatkan seluruh fitur yang tersedia dalam aplikasi melalui menu-menu yang telah ditentukan untuk mengelola informasi polusi secara efektif.
