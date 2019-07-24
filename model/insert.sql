

--Insert of records in to users table

INSERT INTO users
(
    created_at,
    updated_at,
    deleted_at,
    name,
    email
)

SELECT DATETIME('now')              [created_at],
       DATETIME('now')              [updated_at],
       NULL                         [deleted_at],
       'Ian Peter Freely'           [name],
       'IP.Freely@email.com'    [   email]

UNION

SELECT DATETIME('now')              [created_at],
       DATETIME('now')              [updated_at],
       NULL                         [deleted_at],
       'Ivor Biggun'                [name],
       'ivor.biggun@email.com'      [email]

UNION

SELECT DATETIME('now')              [created_at],
       DATETIME('now')              [updated_at],
       NULL                         [deleted_at],
       'Hugh J Axey'                [name],
       'hughj.axey@email.com'       [email]






--DELETE FROM users WHERE id = 2
--SELECT * FROM users

