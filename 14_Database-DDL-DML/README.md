# 14_Section_Database-DDL-DML
-	Database adalah sekumpulan data yang terorganisir. Database memiliki hubungan yang sangat erat dengan API.
-	Database relationship terdiri dari:
1. One to one diartikan sebagai tabel 1 user hanya menyimpan satu tweet atay foto profil. Pada foto profil berisi column id (data type integer), url (data type varchar), dan accountid (data type integer). Agar lebih aman, accountid diseting UN, dimana datanya tidak boleh ada yang sama.
2. One to many artinya tabel 1 user menyimpan banyak tweet.
3. Many to many dapat dicontohkan dengan 1 mahasiswa mengambil banyak mata kuliah dan setiap matakuliah dapat diambil oleh banyak orang.
-	Pengimplementasian relationship one to one, one to many, dan many to many dalam bentuk diagram dapat menggunakan erd (entity raltionship diagram). Dengan erd dapat mendesain sebuah diagaram database. Contohnya tabel pertama yaitu department yang memiliki atribut id (bersifat pk artinya unik), name. tabel kedua yaitu tabel employee yang berisi id (bersifat pk artinya unik), name, phone, department_id.
-	Salah satu contoh software yang menggunakan relational database model sebagai dasarnya adalah MySQL. Jenis perintah SQL:
1.	DDL (data definition languange)
2.	DML (data manipulation languange) digunakan untuk memamnipulasi data dalam tabel suatu database. Statemen operation:
    a. INSERT, ex: input data ke tabel user dengan INSERT INTO USERS (username, fullname, status gender, email, password, location)
    b. SELECT, ex: menampilka semua data pada tabel user dengan SELECT * FROM USERS
    c. UPDATE, ex: mengubah data fullname ke tabel user dengan id  dengan UPDATE users SET fullname=’Bima Sakti’ WHERE id = 1
    d. DELETE, ex: hapus data pada tabel user dengan id 1 dengan DELETE FROM users WHERE id = 1
3.	DCL (data control languange)
