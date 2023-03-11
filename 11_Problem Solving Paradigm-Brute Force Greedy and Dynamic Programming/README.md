# 11_Section_Problem Solving Paradigm-Brute Force Greedy and Dynamic Programming
-	Problem solving paradigm adalah penyelesaian sebuah masalah yang dapat diselesaikan dengan: 
1.	Complete search (brute force) yaitu penyelesaian masalah dengan pengecekan semua data, dari awal sampai akhir. Digunakan saat melakukan pencarian yang searah. Contohnya mencari nilai minimal dan maksimal.
2.	Divide and conquer akan memecah permsalahan ke bagian permasalahan kecil. Contohnya mengimplementasikan binary search yang memiliki data sebagai berikut binary search datanya harus disorting atau diurutkan terlebih dahulu. Jika dibandingkan, sequential dan binary yang paling optimal adalah binary. Divide and conquer secara kompleksitas akan berkaitan dengan o log n.
3.	Greedy adalah sebuah algoritma yang menyelesaikan masalah dengan mencari nilai local optimal dan global optimal. Greedy akan lebih focus ke local optimal. Greedy lebih cepat, bisa optimal dan bisa tidak. Contohnya pada coin change yang memiliki nilai 10, 25, 5, 1 dengan mencari nilai kombinasi terkecil dari 42. Dari 42 akan dilakukan perulangan untuk mencari local. 42-25(nilai terbessar)= 17. Kemudian 17-10=7. Lalu 7-5=2 dan 2-1=1. Ketika 1-1=0 maka perulangan sudah selesai.
4.	Dynamic programming adalah algoritma teknik untuk menyelesaikan problem secara optimal dengan mengutamakan fatka dari setiap problem-problem. Karakteristik:
    a. Overlapping subproblem yaitu saat mencari solusi subproblem akan dipanggil beberapa kali. Contohnya pada fibonanci yang memanggil angka secara berkali-kali.         b. Optimal substructure property yaitu mengoptimalkan masalah yang ditemukan.
    Cara yang dapat digunakan dalam dynamic programming:
    a. Top down with memorization yaitu memecahkan masalah yang paling besar ke masalah yang paling kecil.
    b. Bottom up with tabulation yaitu memecahkan masalah yang kecil dan digabungkan masih kecil-kecil tersebut untuk menyelesaikan masalah yang besar.
