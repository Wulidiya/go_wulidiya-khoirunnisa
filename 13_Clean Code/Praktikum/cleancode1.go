package main

type user struct {
	id       int
	username int
	password int
}

// Kekurangan pada deklarasi tipe data int pada user.
// Alasan: Tipe data username dan password seharusnya merupakan string bukan int karena username dan password biasanya berbentuk teks, bukan bilangan.

type userservice struct {
	t []user
}

// Kekurangan pada deklarasi tipe data []user pada userservice
// Alasan: Nama variabel t pada userservice tidak jelas dan tidak menggambarkan tujuan dari variabel tersebut. Nama variabel sebaiknya lebih deskriptif. contohnya mengubah nama variabel t menjadi users atau userList untuk meningkatkan kejelasan dan membantu pemahaman kode

func (u userservice) getallusers() []user {
	return u.t
}

// Kekurangan pada metode getallusers dan getuserbyid pada userservice
// Alasan: Metode getallusers dan getuserbyid tidak membutuhkan objek userservice secara keseluruhan. Seharusnya metode-metode ini dideklarasikan sebagai metode yang menerima pointer *userservice agar dapat mengakses data users yang ada pada userservice
// Jadi,mengubah metode getallusers dan getuserbyid untuk menerima pointer *userservice

func (u userservice) getuserbyid(id int) user {
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}

	return user{}
}
