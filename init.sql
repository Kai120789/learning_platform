CREATE DATABASE user_db;
CREATE USER user_db_user WITH PASSWORD 'user_db_password';
GRANT ALL PRIVILEGES ON DATABASE user_db TO user_db_user;

\c user_db
GRANT ALL ON SCHEMA public TO user_db_user;