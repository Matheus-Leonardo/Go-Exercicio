-- Cross-check: Definição exata conforme a struct entities.Product em Go
CREATE TABLE IF NOT EXISTS product (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(255) NOT NULL,
    quantity INT NOT NULL
);

-- Dados iniciais para teste
INSERT INTO product (id, name, type, quantity) VALUES ('1', 'Guitarra Fender', 'Instrumento', 5);
INSERT INTO product (id, name, type, quantity) VALUES ('2', 'Teclado Yamaha', 'Instrumento', 2);