CREATE TABLE inventory (
 id text default gen_random_uuid(),
 name text,
 amount int default 0
)