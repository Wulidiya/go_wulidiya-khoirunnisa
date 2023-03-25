# 12_Section_Concurrent Programming
-	Saat melakukan pencarian sebuah website akan terdapat 3 hasil yaitu website, video, dan image. Adapun cara kerja pencarian dengan ketiga hasil tersbeut.
1.	Sequential program yaitu secara berurutan atau satu pes satu. Contohnya video 2 detik, website 3 detik, dan image 4 detik sehingga totalnya adalah 9 detik
2.	Parallel program yaitu mengerjakan 3 pekerjaan secara bersamaan dengan syarat adanya multi-core. Contohnya video, website, dan image akan diproses secara bersama-sama selama 4 detik
3.	Concurrent program yaitu menjalankan tugas secara bersamaan dan independent atau berdiri sendiri dan tidak membutuhkan multi-core. Contohnya website, video, image dapat dipecah ke beberapa detik.
-	Goroutines yaitu sebuah proses concurrent yang dijalankan dalam golang. goroutines adalah sebuah function yang akan dijalankan secara concurrently dibandingkan function lain. Contohnya membuat function hello yang akan menampilkan “hello world” dan menjalankan function hello secara concurrent dengan menulis go hello. Maka saat di run akan menghasilkan kosong karena berjalan dengan independent. Agar bisa output akan dihandel dengan channels dan select. Namun untuk menyelesaikan goroutines menyelesaikan tugasnya dapat menambahkan time.Sleep(1 * time.Second) Sehingga tercetaklah hello world. Terdapat beberapa eksekusi goroutine yaitu:
1.	Main waktu tunggu 200 detik
2.	Getchars waktu runggu 10 detik (paling cepat)
3.	Getdigits waktu tunggu bisa lebih lama dan konverensi berjalan beringingan atau selang-seling
-	Channel dan selesct digunakan saat melakukan sebuah komunikasi goroutine 1 dengan go routin lainnya karena ada tugas yang independent dan saling tunggu. 
-	Channels digunakan untuk komunikasi saat concurretnt dan memiliki hubungan dengan select. Pendeklarasian dan penggunaan channel menggunakan panah ke kiri (<-) untuk menginput sebuah data. Channel dapat menyimpan 1 type data yang dilambangkan “chan” pada golang. contoh:

-	Select untuk mengontrol jalannya concurrent dari channel.channel yang digunakan adalah unbuffer “ ”.
-	Race condition terjadi Ketika 2 threads atau goroutine mengakses memori yang sama, dimana salah satunya adalah weiting. Race condition terjadi karena tidak sinkronnya kesamann Ketika mengakses variabel yang global atau alamat memori yang dicari. Contoh:

Race condition merupakan salah satu kelemahan saat menggunakan concurrent programing karena dapat mengakses file yang sama. Untuk menghendel race condition
-	Pengamanan fixing data race dapat menggunakan:
1.	Waitgroups
2.	Channel 
3.	Mutex 
