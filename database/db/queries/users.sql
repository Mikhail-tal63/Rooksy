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
-- name: GetUserByEmail :one
SELECT 
    id ,
    email   ,
password_hash 

     
FROM users 
WHERE  email = $1
LIMIT 1;

-- name: GetUserByUsername :one
SELECT 
    id ,

    first_name ,
    last_name ,
    username  ,
    email   ,
password_hash ,
bullet_elo  ,
blitz_elo ,
rapid_elo  ,
image_id,
is_active   ,
created_at 
     
FROM users 
WHERE  username = $1
LIMIT 1;