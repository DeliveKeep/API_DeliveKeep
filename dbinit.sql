CREATE TABLE IF NOT EXISTS usuarios (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    cpf VARCHAR(14),
    endereco VARCHAR(255) NOT NULL,
    telefone CHAR(20) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    senha VARCHAR(128) NOT NULL,
    perfil VARCHAR(1) NOT NULL
);

CREATE TABLE IF NOT EXISTS pedidos (
    id SERIAL PRIMARY KEY,
    id_operador INT NOT NULL,
    cpf_cliente INT NOT NULL,
    nome_remetente VARCHAR(100) NOT NULL,
    endereco_remetente VARCHAR(255) NOT NULL,
    nome_destinatario VARCHAR(100) NOT NULL,
    endereco_destinatario VARCHAR(255) NOT NULL,
    codigo_rastreamento VARCHAR(25) NOT NULL,
    status_pedido VARCHAR(255) NOT NULL DEFAULT 'Pedido criado',
    altura REAL NOT NULL,
    comprimento REAL NOT NULL,
    peso REAL NOT NULL,
    largura REAL NOT NULL,
    descricao VARCHAR(255) NOT NULL,
    FOREIGN KEY (cpf_cliente) REFERENCES usuarios(cpf),
    FOREIGN KEY (id_operador) REFERENCES usuarios(id)
)

CREATE TABLE IF NOT EXISTS notificacoes (
    id_notificacao SERIAL PRIMARY KEY,
    id_pedido INT NOT NULL,
    conteudo VARCHAR(255) NOT NULL,
    FOREIGN KEY (id_pedido) REFERENCES pedidos(id)
);