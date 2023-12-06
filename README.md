# task-5-pbi-btpns-andrifirmansyah
Environment
Struktur dokumen / environment dari GoLang yang akan dibentuk kurang lebih sebagai berikut :

app :Menampung pembuatan struct dalam kasus ini menggunakan struct user untuk keperluan data dan authentication
controllers : Berisi antara logic database yaitu models dan query
database: Berisi konfigurasi database serta digunakan untuk menjalankan koneksi database dan migration
helpers : Berisi fungsi-fungsi yang dapat digunakan di setiap tempat dalam hal ini jwt, bcrypt, headerValue
middlewares :Berisi fungsi yang digunakan untuk proses otentikasi jwt yang digunakan untuk proteksi api
models : Berisi models yang digunakan untuk relasi database 
router : Berisi konfigurasi routing / endpoint yang akan digunakan untuk mengakses api
go mod : Yang digunakan untuk manajemen package / dependency berupa library
