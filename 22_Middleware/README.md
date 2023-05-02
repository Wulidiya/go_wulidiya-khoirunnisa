# 22_Section_Middleware

# Middlewere 
- Middlewere adalah sebuah entity yang dipasangkan pada sebuah server
- Contoh third party middlewere
1.	Negroni
2.	Echo
3.	Interpose
4.	Alice
5.	Dan lainnya

- Middleware yang disediakan oleh Echo secara default adalah:
1.	Logger: Middleware untuk melakukan logging pada setiap request yang masuk ke server.
2.	CORS: Middleware untuk mengatur Cross-Origin Resource Sharing (CORS) agar bisa mengontrol akses dari domain yang berbeda.
3.	Recover: Middleware untuk melakukan recovery dari panic yang terjadi di dalam handler, sehingga server tetap berjalan meskipun terdapat error.
4.	BodyDump: Middleware untuk melakukan logging atau dump data body dari request yang masuk.
5.	BasicAuth: Middleware untuk melakukan autentikasi basic HTTP.
6.	Gzip: Middleware untuk melakukan kompresi Gzip pada response.
7.	Static: Middleware untuk menyajikan file statis seperti file HTML, CSS, JavaScript, dan gambar.
8.	Method Override: Middleware untuk melakukan override terhadap method HTTP yang digunakan dalam request.
9.	CSRF: Middleware untuk melindungi aplikasi dari serangan Cross-Site Request Forgery (CSRF). 

- Middleware echo dibagi menjadi 2 jenis yaitu:
1.	echo#Pre() ketika akan set up middlewere maka akan tereksekusi sebeleum root masuk ke controller. Ex: HTTSRerdirect, HTTPSWWWRedirect, Rewrite
2.	echo#Use() middlewere akan dieksekusi setelah mendapatkan akses dari konteks echo yang biasanya terletak di controlles. Ex: BodyLimit, Logger, Gzip, JWTAuth

# Authentication Middleware
- Authentication Middleware merupakan proses pengidentifikasian user. Menentukan apakah request/client berhak mengakses API atau tidak. Tujuannya adalah untuk keamanan data dan dapat melihat request/client telah terdaftar atau belum, apakah sebagai admin atau user biasa.
- 2 jenis middleware:
1.	Basic authentication, proses autentifikasi dengan mengirim data username dan password melalui data headers. Aturan basic authentication dengan mengirim key
2.	JSON web token

# Logger middleware 
- Logger middleware digunakan untuk mencatat informasi yang terjadi pada server terutama dalam request. Saat client melakukan request ke server API misalnya dengan alamat http://localhost:8000/users untuk mendapatkan data users. Maka akan mendapatkan alamat host, url, dan data lainnya. Log akan ditangkap kemudian diolah misalnya membuat file log dan kemudian log akan dilakukan analisis. Jika terdapat keanehan log maka akan dapat diketahui.

# Authentication middleware
- Authentication middleware menggunakan Basic + Token
1.	Basic authentication, proses autentifikasi dengan mengirim data username dan password melalui data headers. Aturan basic authentication dengan mengirim key dan value. 
2.	JSON web token

# JWT middleware 
- Struktur JWT terdiri dari 3 token yang dipisahkan dengan titik. 
1. Token pertama adalah sebuah header atau sebuah hasil encode64 dari JSON yang berisi algoritma yang akan digunakan untuk memvalidasi token yang ke 3. 
2. Token kedua merupakan encode64 dari JSON, dimana saat user mengirimkan token maka dapat mengambil data. 
3. Token ketiga adalah gabungan base64 token pertama (header) dan base64 kedua (payload) yang akan ditambahkan sebuah kunci, kemudian dilakukan sebuah algoritma untuk mengecek validasi token ketiga. JWT menggunakan Bearer + Token.
