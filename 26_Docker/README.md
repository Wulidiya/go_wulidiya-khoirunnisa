# 26_Section_Docker

# Infrastructure and deployment
1.	Docker
2.	System and software deployment
3.	CI/CD
4.	Kubernetes 

# Perbedaan container dan virtual mechine
- Alokasi Resource
--> Pada virtual mechine, alokasi resource dilakukan pada awal instalasi sehingga ketika ada 2 virtual mechine yang sudah ditentukan resource-nya dan salah satunya kehabisan resource, maka virtual mechine yang kehabisan resource tidak dapat mengambil resource dari virtual mechine lain. Sedangkan pada Container, alokasi resource dapat dilakukan oleh host server sehingga host dapat melakukan pengambilan resource pada hardware sesuai yang dibutuhkan container itu sendiri.
- Hypervisor
--> Hypervisor adalah sebuah software induk yang sering digunakan untuk menjalankan sistem di virtual mechine. Jika tidak ada hypervisor, maka virtual mechine tidak dapat berjalan. Namun, container dapat menjalankan program secara langsung di OS itu sendiri.
- Kernel
--> Kernel merupakan sebuah program komputer yang memungkinkan Virtual Machine dapat memiliki akses ke kernel untuk dapat menjalankan aplikasi yang ada di dalamnya. Pada container, akses ke kernel tidak tersedia.

# Docker infrastructure
- 3 bagian docker:
1.	Client yaitu berupa coment pengoperasian docker
2.	Docker host yang didalamnya terdapat docker daemon yang digunakan untuk mengelola komponen yang ada di dalam docker host. Ex: images dan containers
3.	Registry digunakan untuk menyimpan images dari docker agar lebih mudah untuk mengelolanya

- Yang dapat dilakukan dengan docker
1.	Membutuhkan file yang bernama dockerfile dan dibuild. Dockerfile berisi ururtan perintah.
2.	Build docker images yang berisi aplikasi, bins/libs, dan base image
3.	Setelah build berhasil maka disimpan ke sebuah image repository untuk mempermudah deliver yang umumnya image repository menggunakan docker hub
4.	Kemudian untuk menjalankan aplikasi dapat melakukan download and run ke dalam docker host.

- Syntax docker secara umum:
1.	FROM: mendapatkan image dari docker registry
2.	RUN: menjalankan bash command ketika building container
3.	ENV: mengatur variabe; didalam container
4.	ADD: copy file dengan beberapa proses lainnya
5.	COPY: melakukan copy file
6.	WORKDIR: mengatur tugas file directory
7.	ENTRYPOINT: menjalankan command ketika building container
8.	CMD: menjalankan command yang dapat ditumpuk dengan command lain

 
