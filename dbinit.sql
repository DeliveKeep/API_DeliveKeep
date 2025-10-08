CREATE TABLE IF NOT EXISTS usuarios (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    cpf VARCHAR(14) NOT NULL,
    endereco VARCHAR(255) NOT NULL,
    telefone CHAR(20) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    senha VARCHAR(128) NOT NULL
);

CREATE TABLE IF NOT EXISTS pedidos (
    id SERIAL PRIMARY KEY,
    nome_remetente VARCHAR(100) NOT NULL,
    endereco_remetente VARCHAR(255) NOT NULL,
    nome_destinatario VARCHAR(100) NOT NULL,
    endereco_destinatario VARCHAR(255) NOT NULL,
    codigo_rastreamento VARCHAR(25) NOT NULL,
    altura REAL NOT NULL,
    comprimento REAL NOT NULL,
    peso REAL NOT NULL,
    largura REAL NOT NULL,
    descricao VARCHAR(255) NOT NULL
)