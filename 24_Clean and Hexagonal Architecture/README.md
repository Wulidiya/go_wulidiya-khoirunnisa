# 24_Section_Clean architecture
 
# Clean Architecture 
Konsep arsitektur perangkat lunak memiliki tujuan untuk menghasilkan sistem yang lebih bersih, modular, dan terpisah dari detail teknis. Konsep ini dibuat oleh Robert C. Martin yang merupakan seorang ahli pemrograman komputer terkenal. Clean Architecture memungkinkan pengembang perangkat lunak untuk mengembangkan sistem yang fleksibel dan mudah untuk diuji. Hal dikarenakan setiap lapisan dalam Clean Architecture memiliki batasan-batasan yang jelas dan terdefinisi dengan baik, sehingga perubahan dalam satu lapisan tidak akan berdampak pada lapisan-lapisan lainnya. Selain itu, Clean Architecture memungkinkan pengembang untuk mengganti implementasi dari lapisan-lapisan tertentu tanpa mengganggu sistem secara keseluruhan.

# Lapisan dalam Clean Architecture
Dalam Clean Architecture, sistem perangkat lunak dibagi menjadi beberapa lapisan yang terpisah secara fungsional dan teknis. Lapisan-lapisan ini memiliki batasan-batasan yang jelas dan terdefinisi dengan baik. Ada lima lapisan dalam Clean Architecture, yaitu:
1. Entities berisi representasi dari objek-objek bisnis dalam aplikasi. Entities merepresentasikan data dan perilaku yang ada dalam bisnis aplikasi dan dapat berupa objek-objek yang kompleks dan memiliki hubungan antar satu sama lain.
2. Use Cases berisi definisi-definisi dari proses-proses bisnis dalam aplikasi. Use Cases merepresentasikan fitur-fitur dan fungsi-fungsi aplikasi yang ingin dicapai oleh pengguna dan dapat berupa urutan-urutan interaksi antar objek bisnis.
3. Interfaces berisi definisi-definisi dari interface-interface yang digunakan untuk berkomunikasi antara aplikasi dan luar aplikasi. Interfaces dapat berupa REST API, message queues, atau user interface.
4. Infrastructure berisi implementasi dari detail-detail teknis dalam aplikasi, seperti database, third-party libraries, dan frameworks. Lapisan ini bertanggung jawab untuk melakukan interaksi dengan dunia luar.
5. Frameworks and Drivers berisi implementasi dari framework dan driver yang digunakan oleh aplikasi. Lapisan ini bertanggung jawab untuk menghubungkan aplikasi dengan lingkungan pelaksanaannya.
