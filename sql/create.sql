CREATE TABLE IF NOT EXISTS categories (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS items (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	price DECIMAL NOT NULL,
	category_id INTEGER NOT NULL, 
	CONSTRAINT fk_categories
		FOREIGN KEY (category_id) REFERENCES categories(id)
);
CREATE TABLE IF NOT EXISTS orders (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	total DECIMAL NOT NULL,
	status TEXT NOT NULL,
	time_ordered INTEGER NOT NULL,
	time_complete INTEGER
);
CREATE TABLE IF NOT EXISTS order_item (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
  itemid INTEGER NOT NULL,
  orderid INTEGER NOT NULL,
  FOREIGN KEY (itemid) REFERENCES items(id),
  FOREIGN KEY (orderid) REFERENCES orders(id)
);
