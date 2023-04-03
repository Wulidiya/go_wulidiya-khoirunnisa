# 21_Section_ORM and Code Structure (MVC)
# ORM (Object Relational Mapping) 
- ORM akan membantu dalam melakukan konversi data dari tabel menjadi objek. 
- ORM akan otomatis mendapatkan data dari database dan terconvert ke dalam bentuk object atau array object.
- ORM dapat otomatis melakukan chache query (case di beberapa ORM).
- ORM dapat mempercepat proses query database jika dilakukan join lebih dari 10 tabel. 
- Kelebihan ORM:
1.	Tidak memerlukan query yang banyak 
2.	Secara otomatis mengambil data ke dalam menggunakan objek
3.	Memiliki cara sederhana jika ingin melakukan screening data sebelum menyimpan di database
4.	Memiliki fitur cache query
- Kekurangan ORM:
1.	Memuat data hubungan yang tidak diperlukan
2.	Fungsi SQL tertentu yang terkait dengan satu vendor mungkin tidak didukung atau tidak ada fungsi khusus (MySQL : FORCE INDEX)

# GORM 
- GORM merupakan library ORM yang sangat luar biasa untuk Golang dengan tujuan untuk membantu atau mempermudah pengembang perangkat lunak. 
- GORM juga bisa dipasangkan pada echo. Cara mengoneksikan database menggunakan gorm memerlukan username dan password.
- Arti code gorm DB.AutoMigrate(&User{}) adalah melakukan migrasi secara otomatis untuk struct User, yang akan membuat struktur tabel User secara otomatis jika tabel User belum ada di dalam database. Migrasi berarti menyesuaikan struktur tabel dengan struktur model, sehingga jika terdapat perubahan pada struktur model, maka struktur tabel akan diperbarui sesuai dengan perubahan tersebut.
# Database migration 
- Database migration dapat memperbarui versi basis data untuk menyesuaikan perubahan versi aplikasi. Perubahan dapat berupa upgrade ke versi terbaru atau rollback ke versi sebelumnya.
- Database migration digunakan karena dapat mengupdate database lebih simpel, dapat melakukan rollback database dengan simpel, dapat melacak perubahan pada struktur database, history struktur database dapat ditulis disebuah kode, dan selalu kompatibel dengan perubahan versi aplikasi
# MVC (Model, View, dan Controller) 
- MVC adalah cara populer untuk mengatur kode
- Pada structure format MVC, folder model berisi daftar struct untuk pemodelan database, dto, maupun untuk struct bisnis.
- Pada structure format MVC, folder yang digunakan untuk mengisi bisnis-bisnis logik dan validasi bisnis adalah model.

