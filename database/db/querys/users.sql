-- name: CreateUser :one
INSERT INTO users (
     id,

    first_name ,
    last_name ,
    username,
    email ,
    password_hash  ,

    image_id


)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
)
RETURNING *;