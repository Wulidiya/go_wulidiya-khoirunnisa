# 16_Section_Join-Union-Agregasi-Subquery-Function (DBMS)
-	Join adalah sebuah query untuk menggabungkan data tabel 1 dengan data tabel lain. Join yang sering digunakan:
1.	Inner dapat digambarkan sebagai irisan antara 2 tabel yang mana akan mendapatkan nilai yang ada di tabel A dan tabel B.
2.	Left akan menampilkan seluruh data tabel yang ada di kiri yang dikenai kondisi ON dan hanya berisi baris dari tabel disebelah kanan yang memenuhi kondisi join. 
3.	Right akan menampilkan seluruh data tabel yang ada di kanan yang dikenai kondisi ON dan hanya berisi baris dari tabel disebelah kiri yang memenuhi kondisi join.
-	Union akan menggabungkan data 1 dengan data tabel yang kedua. Syarat union adalah kolom harus sama (tabel A dan tabel B yang akan ditampilkan harus sama jumlah tabelnya). Selain itu, type data juga harus sama.
-	Fungsi agregasi merupakan fungsi dimana nilai beberapa baris yang dikelompokkan Bersama untuk nilai ringkasan tinggal. Jenis fungsi agregasi dalam SQL:
1.	Min yaitu fungsi agregasi untuk mencari nilai minimal yang ada di sebuah data. ex: menampilkan id terkecil dari tabel users yaitu SELECT MIN(id) AS id FROM users
2.	Max yaitu fungsi agregasi untuk mencari nilai maksimal yang ada di sebuah data. ex: menampilkan nilai id terbesar dari tabel users dengan SELECT MAX(id) FROM USERS
3.	Sum yaitu fungsi agregasi untuk menentukan jumlah nilai di dalam sebuah kelompok. ex: menampilkan jumlah total favourite_count dari tabel tweets dengan user_id1 yaitu SELECT SUM(favourite_count) FROM tweets WHERE user_id1
4.	Avg digunakan untuk mencari nilai rata-rata. ex: menampilkan nilai rata-rata favourite_count dari tabel tweets dengan user_id1 yaitu SELECT AVG(favourite_count) FROM tweets WHERE user_id=1
5.	Count digunakan untuk mencari jumlah data dari sebuah baris yang ditentukan. ex: menampilkan jumlah data dari tabel tweets dengan user_id1 yaitu SELECT COUNT(1) FROM tweets WHERE user_id=1
6.	Having digunakan untuk menyeleksi data berdasarkan kriteria berupa fungsi agregat. ex: menampilkan data dari tweets dengan jumlah favourite_count per use >2 yaitu SELECT user_id FROM tweets GROUP BY user_id HAVING SUM(favourite_count)> 2
-	Sub-query yaitu query di dalam query SQL lain. Sebuah subquery digunakan untuk mengembalikan data yang akan digunakan dalam query utama sebagai syarat untuk lebih membatasi data yang akan diambil. Sub query dapat digunakan dengan SELECT, INSERT, UPDATE, dan DELETE statements dengan operator seperti “ =, <, >, >=, <=, IN, BETWEEN, dan lainnya”. ex: saat menampilkan data tabel users yang user-id nya ada pada tabel tweets yaitu SELECT * FROM USERS WHERE id IN (SELECT user_id FROM tweets GROUP BY user_id) ;
-	Function adalah sebuah kumpulan statement yang akan mengembalikan sebuah nilai balik pada pemanggilnya.
