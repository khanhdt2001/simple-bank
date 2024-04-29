-- name: CreateSession :one
insert into sessions (
    id,
    username,
    refresh_token,
    user_agent,
    client_ip,
    is_blocked,
    expires_at
) values (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
) returning *;

-- name: GetSessionById :one
Select * from sessions where id = $1 Limit 1;