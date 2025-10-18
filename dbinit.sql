CREATE TABLE IF NOT EXISTS clientes (
    id_cliente SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    cpf VARCHAR(14) UNIQUE NOT NULL,
    endereco VARCHAR(255) NOT NULL,
    telefone VARCHAR(20) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    senha VARCHAR(128) NOT NULL
);

CREATE TABLE IF NOT EXISTS galpoes (
    id_galpao SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    endereco VARCHAR(255) NOT NULL
);


CREATE TABLE IF NOT EXISTS operadores (
    id_operador SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    telefone VARCHAR(20) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    senha VARCHAR(128) NOT NULL,
    galpao INT NOT NULL,
    FOREIGN KEY (galpao) REFERENCES galpoes(id_galpao) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS administradores (
    id_administrador SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    telefone VARCHAR(20) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    senha VARCHAR(128) NOT NULL,
    galpao INT NOT NULL,
    FOREIGN KEY (galpao) REFERENCES galpoes(id_galpao) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS encomendas (
    id_encomenda SERIAL PRIMARY KEY,
    nome_remetente VARCHAR(100) NOT NULL,
    endereco_remetente VARCHAR(255) NOT NULL,
    codigo_rastreamento VARCHAR(25) NOT NULL,
    status_pedido VARCHAR(255) NOT NULL DEFAULT 'Pedido criado',
    altura REAL NOT NULL,
    comprimento REAL NOT NULL,
    peso REAL NOT NULL,
    largura REAL NOT NULL,
    descricao VARCHAR(255) NOT NULL,
    cpf_cliente VARCHAR(14) NOT NULL,
    galpao INT NOT NULL,
    FOREIGN KEY (cpf_cliente) REFERENCES clientes(cpf) ON DELETE CASCADE,
    FOREIGN KEY (galpao) REFERENCES galpoes(id_galpao) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS notificacoes (
    id_notificacao SERIAL PRIMARY KEY,
    id_encomenda INT NOT NULL,
    conteudo VARCHAR(255) NOT NULL,
    FOREIGN KEY (id_encomenda) REFERENCES encomendas(id_encomenda)
);