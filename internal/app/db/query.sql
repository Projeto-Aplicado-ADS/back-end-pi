-- name: ListUsers :many
SELECT * FROM users;

-- name: ListAdmins :many
SELECT * FROM users
WHERE is_admin = 1;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = ?;

-- name: CreateUser :execresult
INSERT INTO users (
  id, full_name, email, phone, password, birthday, created_at
) VALUES (
  ?, ?, ?, ?, ?, ?, ?
);


-- name: GetUserByEmail :one
SELECT email, password, is_admin FROM users 
WHERE email = ?;

-- name: ListUserByEmail :one
SELECT * FROM users
WHERE email = ?;


-- name: CreateReserva :execresult
INSERT INTO reservas (
  id, tipo_reserva, data_reserva, data_checkin, data_checkout, status_reserva, valor_reserva, id_quarto, id_hospede, created_at
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
);

-- name: ListReservas :many
select s.id ,s.data_reserva, s.data_checkin, s.data_checkout, s.valor_reserva, s.tipo_reserva , s.status_reserva ,h.nome , h.cpf , q.numero_quarto , q.numero_andar , q.tipo_quarto , q.status_quarto, s.created_at  from reservas s 
inner join hospedes h on s.id_hospede = h.id 
inner join quartos q on s.id_quarto = q.id;

-- name: AlterarReserva :execresult
UPDATE reservas
SET status_reserva = ?
WHERE id = ?;

-- name: RemoverReserva :execresult
DELETE FROM reservas
WHERE id = ?;

-- name: CreateHospede :execresult
INSERT INTO hospedes (
  id, nome, email, telefone, cpf, data_nascimento, sexo, created_at
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?
);

-- name: ListHospedes :many
SELECT * FROM hospedes where deleted_at = 0;

-- name: ListHospedeById :one
SELECT * FROM hospedes
WHERE id = ?;

-- name: AlterarHospede :execresult
UPDATE hospedes
SET nome = ?, email = ?, telefone = ?, cpf = ?, data_nascimento = ?, sexo = ?
WHERE id = ?;

-- name: RemoverHospede :exec
UPDATE hospedes SET deleted_at = ? WHERE id = ?;

/* -- name: RemoverQuarto :exec
UPDATE quartos SET deleted_at = ? WHERE id = ?;  */

-- name: CreateQuarto :execresult
INSERT INTO quartos (
  id, numero_quarto, numero_andar, descricao, tipo_quarto, status_quarto, created_at
) VALUES (
  ?, ?, ?, ?, ?, ?, ?
);

-- name: ListQuartos :many
SELECT * FROM quartos;  

-- name: ListQuartoById :one
SELECT * FROM quartos
WHERE id = ?; 

-- name: AlterarQuarto :execresult
UPDATE quartos
SET numero_quarto = ?, numero_andar = ?, descricao = ?, tipo_quarto = ?, status_quarto = ?
WHERE id = ?; 

-- name: AlterarStatusQuarto :exec
UPDATE quartos SET status_quarto = ? WHERE id = ?;
