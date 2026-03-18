CREATE TABLE products (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(255) NOT NULL,
    quantity INT NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Dados iniciais para validação do CRUD
INSERT INTO products (id, name, type, quantity) VALUES 
('1', 'Guitarra Fender', 'Instrumento', 5),
('2', 'Pedal Boss', 'Acessório', 10);