DROP TABLE IF EXISTS product;

CREATE TABLE product (
  id INT AUTO_INCREMENT NOT NULL,
  title VARCHAR(128) NOT NULL,
  description VARCHAR(255) NOT NULL,
  price DECIMAL(5, 2) NOT NULL,
  discount_percentage DECIMAL(5, 2),
  rating DECIMAL(5, 2),
  stock INT,
  brand VARCHAR(255),
  category VARCHAR(255),
  thumbnail VARCHAR(255),
  PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS image;

CREATE TABLE image (
  id INT AUTO_INCREMENT NOT NULL,
  product_id INT NOT NULL,
  img VARCHAR(128) NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (product_id) REFERENCES product(id)
);

-- Insert dummy data into the product table
INSERT INTO
  product (
    title,
    description,
    price,
    discount_percentage,
    rating,
    stock,
    brand,
    category,
    thumbnail
  )
VALUES
  (
    'Product 1',
    'Description for product 1',
    10.99,
    5.00,
    4.5,
    10,
    'Brand 1',
    'Category 1',
    'thumbnail1.jpg'
  ),
  (
    'Product 2',
    'Description for product 2',
    15.50,
    10.00,
    3.0,
    5,
    'Brand 2',
    'Category 2',
    'thumbnail2.jpg'
  ),
  (
    'Product 3',
    'Description for product 3',
    20.00,
    0.00,
    4.0,
    8,
    'Brand 3',
    'Category 3',
    'thumbnail3.jpg'
  );

-- Insert dummy data into the image table
INSERT INTO
  image (product_id, img)
VALUES
  (1, 'product1_image1.jpg'),
  (1, 'product1_image2.jpg'),
  (1, 'product1_image3.jpg'),
  (2, 'product2_image1.jpg'),
  (2, 'product2_image2.jpg'),
  (3, 'product3_image1.jpg'),
  (3, 'product3_image2.jpg'),
  (3, 'product3_image3.jpg');