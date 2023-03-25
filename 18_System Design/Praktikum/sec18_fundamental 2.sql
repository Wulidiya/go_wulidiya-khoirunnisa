-- Terdapat sebuah query pada SQL yaitu `SELECT * FROM users;`

-- Dengan tujuan yang sama, tuliskan dalam bentuk perintah:

-- 1. Redis
SELECT * FROM users;

-- 2. Neo4j
MATCH * (x:user) RETURN x;

-- 3. Cassandra
SELECT * FROM users_table;