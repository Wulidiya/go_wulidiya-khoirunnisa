# 23_Section_Unit Testing

- Softweare testing adalah proses untuk menganalisa sebuah sistem atau software untuk mendeteksi perbedaan antara kondisi fitur saat ini dengan fitur yang diinginkan oleh user. 
- Tujuan melakukan testing: 
1.	Untuk mengetahui sistem yang dibuat saat ini sudah berjalan dengan baik.
2.	Dapat mencegah terjadinya regression
3.	 Dapat meningkatkan kepercayaan dalam refactoring
4.	 Dapat meningkatkan desain kode
5.	 Dapat melakukan code dokumentation
6.	 Dapat melakukan scaling dengan tim
- Level dari testing
1.	UI test: melakukan test secara end to end melalui user interface sistem yang telah dibuat
2.	Integration
3.	Unit: melakukan testing terhadap unit terkecil dari aplikasi yang telah dibuat
- Cara melakukan test:
1.	Membuat file test baru pada golang dengan membuat package file dan menulis functionnya
2.	Melakukan run testing dengan menuliskan go test ./lib/calculate -cover pada terminal

 
