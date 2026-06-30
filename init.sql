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


CREATE DATABASE lesson_db;
CREATE USER lesson_db_user WITH PASSWORD 'lesson_db_password';
GRANT ALL PRIVILEGES ON DATABASE lesson_db TO lesson_db_user;

\c lesson_db
GRANT ALL ON SCHEMA public TO lesson_db_user;


CREATE DATABASE subject_db;
CREATE USER subject_db_user WITH PASSWORD 'subject_db_password';
GRANT ALL PRIVILEGES ON DATABASE subject_db TO subject_db_user;

\c subject_db
GRANT ALL ON SCHEMA public TO subject_db_user;


CREATE DATABASE schedule_db;
CREATE USER schedule_db_user WITH PASSWORD 'schedule_db_password';
GRANT ALL PRIVILEGES ON DATABASE schedule_db TO schedule_db_user;

\c schedule_db
GRANT ALL ON SCHEMA public TO schedule_db_user;