CREATE DATABASE user_db;
CREATE USER user_db_user WITH PASSWORD 'user_db_password';
GRANT ALL PRIVILEGES ON DATABASE user_db TO user_db_user;

\c user_db
GRANT ALL ON SCHEMA public TO user_db_user;


CREATE DATABASE group_db;
CREATE USER group_db_user WITH PASSWORD 'group_db_password';
GRANT ALL PRIVILEGES ON DATABASE group_db TO group_db_user;

\c group_db
GRANT ALL ON SCHEMA public TO group_db_user;
