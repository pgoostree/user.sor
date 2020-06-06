-- this runs inside docker container
\i /docker-entrypoint-initdb.d/schema/schema.sql
\i /docker-entrypoint-initdb.d/tables/user.sql
\i /docker-entrypoint-initdb.d/tables/group.sql
\i /docker-entrypoint-initdb.d/tables/user_group.sql
\i /docker-entrypoint-initdb.d/procs/insert_user.sql
\i /docker-entrypoint-initdb.d/procs/get_user.sql
\i /docker-entrypoint-initdb.d/procs/update_user.sql
\i /docker-entrypoint-initdb.d/procs/delete_user.sql
\i /docker-entrypoint-initdb.d/procs/insert_group.sql
\i /docker-entrypoint-initdb.d/procs/delete_group.sql
\i /docker-entrypoint-initdb.d/procs/get_user_group.sql
\i /docker-entrypoint-initdb.d/procs/insert_user_group.sql
\i /docker-entrypoint-initdb.d/procs/update_user_group.sql