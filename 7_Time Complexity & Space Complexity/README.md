# 7_Section_Time Complexity & Space Complexity
Time Complexity adalah bagaimana kita menghitung sebuah sistem akan berjalan dengan waktu. Waktu dihitung dengan menentukan seberapa banyak operasi dominan yang dilakukan.
Operasi dominan (Big-O notation) dapat dilihat dari contoh berikut:
Func dominant (n int) int {
Var result int = 0
For I := 0; I < n; i++ {
Result += 1
}
Result = result + 10
Return result
}
Dari contoh diatas yang disebut operasi dominan adalah Result += 1 karena dilakukan sebanyak n kali, artinya jika kita menginput nilai 50 maka Result += 1 dilakukan sebanyak 50 kali. Operasi dominan (big O notation yang merupakan jenis linier complexity) artinya jika menginput n maka operasi yang paling dominan akan dilakukan sebanyak n.
Time complexities terdiri dari:
1. Constanta time-O(1), artinya Ketika menginput berapapun maka notasi dalam algoritma dinotasikan sebanyak 1 kali (tidak berpengaruh terhadap inputan).
2. Linier time–O (n) merupakan complexities paling normal dan paling mudah.
3. Linier time–O (n+m) yang memiliki 2 operasi dominan yaitu berdasarkan m dan n.
4. Logarithmic time–O (log n) merupakan complexities yang paling efektif jika ingin melakukan problem solving terkait percepatan. O log n memiliki cara kerja 2^x=n. misalnya n=8 maka x bernilai 3.
5. Quadratic time–O (n^2), biasanya sering digunakan ketika problem solving untuk mencari kombinasi dan permutasi. Namun, tidak efektif. Misal diinputkan 10 maka akan butuh melakukan sebuah operasi sebanyak 100 kali karena 10x10.
6. Exponential time O(2^n) and factorial time O(n!)
Ketika menyelesaikan problem solving untuk time complexity diusahakan menggunakan O(1) atau  O (log n). sedangkan untuk space complexity diusahakan konstan. Jadi diusahakan variabel itu minimal sehingga kecepatan operasi yang dijalankan cukup cepat.
