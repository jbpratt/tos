/*
sandwiches, salads, teas
*/
CREATE TABLE IF NOT EXISTS item_kinds (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  deleted INTEGER DEFAULT 0,
  name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS items (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  deleted INTEGER DEFAULT 0,
  available INTEGER DEFAULT 1,
  kind_id INTEGER NOT NULL,
  name TEXT NOT NULL,
  price INTEGER NOT NULL,
  FOREIGN KEY(kind_id) REFERENCES item_kinds(id)
);

/* 
meats, cheeses, sauces, breads, dressings, prep instruction
*/
CREATE TABLE IF NOT EXISTS option_kinds (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL
);

/* 
meat: smoked pulled pork, brisket
cheeses: american, gouda, provalone, swiss
prep instructions: crispy, over easy, bleu
sizes: LG, RG
*/
CREATE TABLE IF NOT EXISTS options (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  deleted INTEGER DEFAULT 0,
  available INTEGER DEFAULT 1,
  kind_id INTEGER NOT NULL,
  name TEXT NOT NULL,
  price INTEGER NOT NULL,
  light_price INTEGER,
  heavy_price INTEGER,
  FOREIGN KEY(kind_id) REFERENCES option_kinds(id)
);

/* item:option associations with optional price overrides */
CREATE TABLE IF NOT EXISTS item_options (
  item_id INTEGER NOT NULL,
  option_id INTEGER NOT NULL,
  is_default INTEGER DEFAULT 0,
  price INTEGER,
  light_price INTEGER,
  heavy_price INTEGER,
  PRIMARY KEY(item_id, option_id),
  FOREIGN KEY(item_id) REFERENCES items(id),
  FOREIGN KEY(option_id) REFERENCES options(id)
);

/*
modifiable side items like potatos, side salads, toast, eggs...
*/
CREATE TABLE IF NOT EXISTS item_sides (
  item_id INTEGER NOT NULL,
  side_item_id INTEGER NOT NULL,
  is_default INTEGER DEFAULT 0,
  price INTEGER,
  PRIMARY KEY(item_id, side_item_id),
  FOREIGN KEY(item_id) REFERENCES items(id),
  FOREIGN KEY(side_item_id) REFERENCES items(id)
);

CREATE TABLE IF NOT EXISTS orders (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  created_at INTEGER NOT NULL,
  completed_at INTEGER
);

/*
etc payment state
*/

CREATE TABLE IF NOT EXISTS order_items (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  order_id INTEGER NOT NULL,
  item_id INTEGER NOT NULL,
  main_order_item_id INTEGER, /* defined for sides */
  price INTEGER,
  FOREIGN KEY(order_id) REFERENCES orders(id),
  FOREIGN KEY(item_id) REFERENCES items(id),
  FOREIGN KEY(main_order_item_id) REFERENCES order_items(id)
);

CREATE TABLE IF NOT EXISTS order_item_options (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  order_id INTEGER NOT NULL,
  option_id INTEGER NOT NULL,
  order_item_id INTEGER NOT NULL,
  price INTEGER,
  FOREIGN KEY(order_id) REFERENCES orders(id),
  FOREIGN KEY(option_id) REFERENCES options(id),
  FOREIGN KEY(order_item_id) REFERENCES order_items(id)
);
