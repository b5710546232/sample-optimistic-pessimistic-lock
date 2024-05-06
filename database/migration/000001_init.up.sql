CREATE TABLE inventory (
 id text default gen_random_uuid() PRIMARY KEY,
 name text,
 amount int default 0
)