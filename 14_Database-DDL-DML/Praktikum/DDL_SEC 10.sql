CREATE TABLE users (
	id int(11) not null AUTO_INCREMENT PRIMARY KEY,
    gender char(1) NOT NULL,
    dob DATE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
CREATE TABLE payment_methods (
	id int(11) not null AUTO_INCREMENT PRIMARY KEY,
    name varchar(255),
    status smallint,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
CREATE TABLE operators (
	id int(11) not null AUTO_INCREMENT PRIMARY KEY,
	name varchar(255),
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE product_types (
	id int(11) not null AUTO_INCREMENT PRIMARY KEY,
	name varchar(255),
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE transactions (
	id int(11) not null AUTO_INCREMENT PRIMARY KEY,
	user_id int(11),
	payment_method_id int(11),
	status varchar(10),
	total_qty int(11),
	total_price NUMERIC(25,2),
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP DEFAULT NOW(),
	FOREIGN KEY(user_id) REFERENCES users(id),
	FOREIGN KEY(payment_method_id) REFERENCES payment_methods(id)
);

CREATE TABLE products (
	id int(11) not null AUTO_INCREMENT PRIMARY KEY,
	product_type_id int(11),
	operator_id int(11),
	code varchar(50),
	status smallint,
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP DEFAULT NOW(),
	FOREIGN KEY(product_type_id) REFERENCES product_types(id),
	FOREIGN KEY(operator_id) REFERENCES operators(id)
);

CREATE TABLE product_descriptions (
	id int(11) not null AUTO_INCREMENT PRIMARY KEY,
	product_id int(11),
	description TEXT,
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP DEFAULT NOW(),
	FOREIGN KEY (product_id) REFERENCES products(id)
);


CREATE TABLE transaction_details (
	id int(11) not null AUTO_INCREMENT PRIMARY KEY,
	transaction_id int(11),
	product_id int(11),
	status varchar(10),
	qty int(11),
	price NUMERIC(25,2),
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP DEFAULT NOW(),
	FOREIGN KEY (transaction_id) REFERENCES transactions(id),
	FOREIGN KEY (product_id) REFERENCES products(id)
);
CREATE TABLE kurir (
	id int(11) not null AUTO_INCREMENT PRIMARY KEY,
    name varchar(255),
    nomor_telephon char(13),
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP DEFAULT NOW()
);
alter TABLE kurir add onngkos_dasar numeric(25,2) not null;
alter TABLE kurir rename to shipping;
drop TABLE shipping;

CREATE TABLE payment_method_description (
	id int(11) not null AUTO_INCREMENT PRIMARY KEY,
    deskripsi varchar(255),
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP DEFAULT NOW()
);
CREATE TABLE alamat (
	id int(11) not null AUTO_INCREMENT PRIMARY KEY,
    user_id int(11),
    nama_jalan varchar(11),
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id)
);
CREATE TABLE user_payment_method_detail (
	id int(11) not null AUTO_INCREMENT PRIMARY KEY,
    user_id int(11),
    payment_method int(11),
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (payment_method) REFERENCES payment_methods(id)
);

alter TABLE users add FOREIGN KEY (users_payment_method_detail) REFERENCES user_payment_method_detail (id);
alter TABLE users add users_payment_method_detail int(11) not null;
use alta_online_shop
