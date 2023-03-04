# 9_Section_String-Advance Function-Pointer-Method-Struct and Interface
- String adalah salah satu variabel penting dalam membuat program. Untuk memulai string jangan lupa untuk import dan memasukkan fmt dan strings. Pada golang string dibedakan menjadi:
1. len string
Untuk menghitung panjang string bisa menggunakan keyword len string. Contohnya memiliki sebuah variabel string sentence yang berisi hello dan variabel lensentence yang berisi panjang sentence. Setelah itu di Println(lenSentence) untuk mengetahui panjang string dari jumlah huruf dari hello. Nah setelah itu akan menunjukkan nilai 5 yang merupakan huruf dari hello.
2. Compire string
Misalnya ingin mengcompire string 1 berupa abc dan string 2 berupa abd. Nah untuk membandingkan kedua string tersebut bisa di di Println(str1 == str2). Setelah di run maka akan muncul hasilnya yaitu berupa false karena string 1 tidak sama dengan string 2.
3. Contain 
Untuk membandingkan apakah string ke-2 merupakan bagian dari string1 atau tidak. Contohnya membuat konstantan yang berisi string “something” dan substring “some”. Contain dapat dibuat dengan res := strings.contains(str, substr) dan di println(res). Ketika di run akan membuktikan apakah string 2 adalah bagian dari string 1 maka hasilnya some merupakan bagian dari something.
4. Mengambil indeks tertentu dari sebuah string (substring). Contohnya memiliki string value yang berisi cat, dog dan ingin mengambil nilai dog. Caranya yaitu membuat substring yang berisi value[4:len(value)] 4 berasal dari indeks dog yang berasal di no.4. kemudian println(substring). Ketika di run akan menghasilkan dog.
5. Replace yang digunakan untuk mengganti string. Contohnya memiliki string S yang berisi “katak” dan akan mereplace huruf a menjadi o. caranya dengan t := string.Replace(s, “a”, ”o”, -1). A merupakan huruf yang akan diganti, o merupakan huruf pengganti, dan -1 artinya akan mengganti semua huruf a yang ada di kata tersebut. Setelah itu fmt.Println(t). Ketika di run akan menghasilkan “kotok”.
6. Insert 
Contohnya string p berisi “green” dengan index berisi 2. Kemudian menginsert index kata kedua green dengan q yang berisi p[:indeks] + “HAY” + p[indeks:]. Lalu di println(p,q). Setelah di run maka akan menghasilkan “green grHAYeen”
- Advance function
1. Variadic function, ketika ingin memanggil sebuah function tapi parameter yang inputkan memiliki jumlah yang berbeda. Kita membuat tempo reslice untuk dilakukan sebagai parameter. Pada slice integer diganti pada variadic function dengan mengganti “…”
2. Anonymous function atau literal function yang sering digunakan di sebuah projek, dimana function tidak mempunyai nama. Biasanya untuk mengelompokkan sebuah kumpulan code di dalam function tanpa nama. Misalnya functionnya terlalu banyak maka dapat dikelompokkan agar lebih mudah.
3. Closure adalah sebuah tindak lanjut anonymous function dimana ketika memakai anonymous function itu variabel di luar, maka akan melakukan references. Misalnya ada variabel dipanggil berkali-kali maka akan melanjutkan tidak membuat baru.
4. Defer function adalah penjabaran dari anonymous function dimana kelompok yang depannya diberi defer berarti kelompok function tersebut akan dirunning di akhir.
- Pointer adalah sebuah variabel yang dapat menyimpan memory address dari variabel yang lain. Pointer biasanya disajikan dengan penambahan operator “ * ” dan “ & ”. Contohnya yaitu *p dimana artinya adalah variabel p adalah sebuah pointer dan untuk mengakses nilai pada alamat memori. Lalu simbol & merupakan penunjukkan untuk alamat dari suatu pointer.
- Struct berisi collection of named fields/function/method.
- Methods adalah sebuah function yang ditempel pada sebuah type. Method dapat membantu menulis object oriented style, membantu menghindari naming conflic, dan agar codingan mudah dimengerti maupun dibaca. Method bisa diakses melalui lewat variabel objek (struct).
