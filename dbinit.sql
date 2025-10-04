CREATE TABLE IF NOT EXISTS usuarios (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    cpf VARCHAR(14) NOT NULL,
    endereco VARCHAR(255) NOT NULL,
    telefone CHAR(20) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    senha VARCHAR(128) NOT NULL
);

CREATE TABLE IF NOT EXISTS encomendas (
    id SERIAL PRIMARY KEY,
    descricao VARCHAR(255) NOT NULL,
    codigo_rastreamento VARCHAR(100) UNIQUE,
    remetente_nome VARCHAR(150) NOT NULL,
    remetente_endereco TEXT NOT NULL,
    destinatario_nome VARCHAR(150) NOT NULL,
    destinatario_endereco TEXT NOT NULL,
    altura_cm DECIMAL(10, 2) NOT NULL,
    largura_cm DECIMAL(10, 2) NOT NULL,
    comprimento_cm DECIMAL(10, 2) NOT NULL,
    peso_kg DECIMAL(10, 2) NOT NULL,
    criado_em TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    
    -- Chave estrangeira para ligar a encomenda ao usuário
    usuario_id INTEGER NOT NULL,
    FOREIGN KEY (usuario_id) REFERENCES usuarios(id) ON DELETE CASCADE
);
