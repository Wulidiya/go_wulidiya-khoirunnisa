# 19_Section_Intro RESTful API
- Implementasi: mengintegrasikan aplikasi frontend dan backend ataupun backend dan backend
- Cara kerja: terjadi proses request dari frontend ke backend dan API akan memberikan respon sesuai dengan permintaannya.
- REST (Representational State Transfer) memiliki format JSON (JavaScript Object Notation), XML, dan SOAP. Format JSON: Code, massage, status, data
# HTTP request method:
1.	GET untuk meminta data dari resource yang telah ditentukan pada UR
2.	POST untuk melakukan update dan jika data tidak ditemukan maka akan ditambahkan data baru
3.	PUT untuk mengirimkan data ke server untuk memperbarui atau membuat resource
4.	DELETE untuk menghapus resource yang telah ditentukan pada URI.
# HTTP Code response:
1.	200 (OK), server telah berhasil memproses permintaan dan memberikan respons sesuai yang diminta.
2.	201 (Created), server telah berhasil membuat resource yang baru, seperti membuat data baru di database.
3.	400 (Bad request), jika terjadi kesalahan ketika client memasukkan data ke dalam request endpoint Rest API
4.	404 (not found) resource yang diminta oleh client tidak ditemukan atau tidak tersedia di server.
5.	401 (unauthorized), client tidak memiliki akses yang sah ke resource yang diminta.
6.	405 (method not allowed), client telah mengirimkan HTTP method yang tidak diizinkan oleh resource yang diminta.
7.	500 (internal server error), terdapat kesalah didalam server
# API dapat diakses dengan:
1.	Katalon
2.	Apache JmMeter
3.	SoapUI
4.	Postman adalah http client yang digunakan untuk melakukan testing website, membuat development dan dokumen dengan mudah. Didalam testing rest api di postman, agar kita tidak perlu menginputkan value berkali-berkali sebagai contoh host url / token maka fitur apa yang perlu kita gunakan environtment variable
