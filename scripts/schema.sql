USE orders;
CREATE TABLE orders (id varchar(255) NOT NULL, price double NOT NULL, tax double NOT NULL, final_price double NOT NULL, PRIMARY KEY (id));