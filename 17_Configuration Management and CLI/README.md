# 17_Section_Configuration Management and CLI
# CLI (Command line interface)
CLI dapat lebih cepat memanagement tugas besar dan dapat melakukan penyimpanan sebuah perintah secara berurutan yang dapat menjalankan tugas secara otomatis. Untuk melakukan perintah disebut shell dan tempat untuk menjalankan perintah disebut terminal.
# Contoh shell legend 
me@linuxbox: ~ $
- Me: username
- Linuxbox: hostname
- ~: current path (home)
- $: normal user bisa memanimupalisi manajemen folder tugas yang tempatnya di dalam home/username
- #: root user dapat mengakses semua folder tugas
# Perintah yang sering digunakan dalam unix shell
-	Directory 
1.	pwd untuk mengetahui direktori yang sedang dibuka
2.	ls untuk menampilkan file yang ada di work directory, ls -a untuk menampilkan file yang terhiden, ls -l untuk menampilkan file mendetail
3.	mkdir untuk membuat suatu folder
4.	cd untuk menampilkan nama direktori atau mengubah lokasi direktori
5.	rm untuk mendelet file/folder, rm-r untuk menghapus folder dan isinya 
6.	cp untuk mengcopy file 1 ke alamat yang akan dituju
7.	mv untuk memindah 1file ke directory lain
8.	ln untuk membuat sebuah link yang bersifat softlink dan hardlink. Penggunaan soflink dengan perintah ln -s
-	Files
1.	Create: touch untuk membuat sebuah file ex: touch[spasi]nama file
2.	View: head, cat, tail, less untuk menampilkan data yang ada di dalam file. Ex: cat[spasi]nama file
3.	Editor: vim, nano. ex: vim[spasi]file yang ingin diedit; nano[spasi]nama file
4.	Permission: chown, chmod
5.	Different: diff untuk membandingkan file 1 dengan lain
-	Network 
1.	ping untuk melakukan tes koneksi internet ex: ping 8.8.8.8
2.	ssh untuk koneksi ke remote computer
3.	netstat untuk mengetahui didalam jaringan terjadi apa saja
4.	nmap
5.	ifconfig untuk menampilkan 
6.	wget dan curl untuk unduh file dari url ex: wget[spasi]url
7.	telnet untuk koneksi remote computer namun lebih baik menggunakan s$h
8.	netcat untuk monitoring jaringan
-	utility
1.	man sebuah informasi perintah apa yang dibingungkan ex: man ls maka akan muncul deskripsi ls
2.	env
3.	echo untuk menampilkan ex: encho “Hello CLI”
4.	date menampilkan waktu
5.	which untuk mengetahui lokasi file
6.	watch
7.	sudo untuk meminta permission dari user ke root
8.	history
9.	grep untuk mencari file yang mengandung suatu kata
10.	locate untuk pencarian file
