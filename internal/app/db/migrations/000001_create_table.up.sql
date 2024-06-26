CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(36) NOT NULL PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    email VARCHAR(200) NOT NULL UNIQUE,
    phone VARCHAR(14) NOT NULL,
    password VARCHAR(255) NOT NULL, 
    is_admin BOOLEAN NOT NULL DEFAULT 0,
    created_at BIGINT NOT NULL,
    update_at BIGINT NOT NULL DEFAULT 0
);



CREATE TABLE IF NOT EXISTS quartos (
    id VARCHAR(36) NOT NULL PRIMARY KEY,
    numero_quarto INT NOT NULL,
    numero_andar INT NOT NULL,
    tipo_quarto ENUM('simples', 'casal', 'familia', 'luxo') NOT NULL,
    status_quarto ENUM('livre', 'ocupado', 'reservado') NOT NULL DEFAULT 'livre',
    descricao VARCHAR(255) NOT NULL,
    created_at BIGINT NOT NULL,
    update_at BIGINT NOT NULL DEFAULT 0
); 


CREATE TABLE IF NOT EXISTS hospedes (
    id VARCHAR(36) NOT NULL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    email VARCHAR(200) NOT NULL UNIQUE,
    telefone VARCHAR(17) NOT NULL,
    cpf VARCHAR(11) NOT NULL UNIQUE,
    data_nascimento VARCHAR(14) NOT NULL,
    sexo ENUM('Masculino', 'Feminino') NOT NULL,
    created_at BIGINT NOT NULL,
    update_at BIGINT NOT NULL DEFAULT 0,
    deleted_at BIGINT NOT NULL DEFAULT 0
); 



CREATE TABLE IF NOT EXISTS reservas (
    id VARCHAR(36) NOT NULL PRIMARY KEY,
    tipo_reserva ENUM('diaria', 'semanal', 'mensal') NOT NULL,
    data_reserva VARCHAR(14) NOT NULL,
    data_checkin VARCHAR(24) NOT NULL,
    data_checkout VARCHAR(24) NOT NULL,
    status_reserva ENUM('pendente', 'confirmada', 'cancelada') NOT NULL DEFAULT 'pendente',
    valor_reserva DECIMAL(10,2) NOT NULL,
    created_at BIGINT NOT NULL,
    updated_at BIGINT DEFAULT NULL DEFAULT 0,
    deleted_at BIGINT NOT NULL DEFAULT 0,
    id_quarto VARCHAR(36) NOT NULL,
    id_hospede VARCHAR(36) NOT NULL,
    FOREIGN KEY (id_quarto) REFERENCES quartos(id) ON UPDATE CASCADE,
    FOREIGN KEY (id_hospede) REFERENCES hospedes(id) ON UPDATE CASCADE
); 
