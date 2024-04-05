 CREATE TABLE orders (
   order_id SERIAL PRIMARY KEY,
   customer_name VARCHAR(255),
   ordered_at TIMESTAMP
 );

 CREATE TABLE items (
   item_id SERIAL PRIMARY KEY,
   item_code VARCHAR(255),
   description TEXT,
   quantity INTEGER,
   order_id INTEGER REFERENCES orders(order_id)
 );
