INSERT INTO operators (name) values ('TOKO BAHARI');
INSERT INTO operators (name) values ('TOKO MAJU JAYA');
INSERT INTO operators (name) values ('TOKO ANDALAS');
INSERT INTO operators (name) values ('TOKO INDOMERAH');
INSERT INTO operators (name) values ('TOKO ALFA BETA');

SELECT * from operators;

INSERT INTO product_types (name) values ('SEPATU');
INSERT INTO product_types (name) values ('TAS');
INSERT INTO product_types (name) values ('BAJU');

SELECT * from product_types;

INSERT INTO products (product_type_id, operator_id, code, status)
VALUES ( 1, 3, 'Bata55', 1),( 2, 5, 'Eksport99', 1),( 3, 2, 'Elzatta29', 1);

SELECT * from products;

INSERT INTO products (product_type_id, operator_id, code, name, status)
VALUES ( 1, 3, 'Bata55', 'Flat shoes', 1),( 2, 1, 'Eksport99', 'Tas sekolah',1),( 3, 4, 'Elzatta29', 'Gamis anak',1);

SELECT * from products;

delete from products
where code='Bata55';

use alta_online_shop