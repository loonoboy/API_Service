DO
$do$
BEGIN
	IF NOT EXISTS (SELECT FROM   pg_catalog.pg_roles WHERE rolname = 'api_user') THEN
CREATE ROLE api_user WITH LOGIN PASSWORD 'qwerty';
END IF;
END
$do$;

DO
$do$
BEGIN
	IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'api_db') THEN
    	CREATE DATABASE api_db
				WITH OWNER = api_user
					ENCODING = 'utf8'
					TABLESPACE = pg_default
					LC_COLLATE = 'en_US.utf8'
					LC_CTYPE = 'en_US.utf8'
					CONNECTION LIMIT = -1;
END IF;
	GRANT CONNECT, TEMPORARY ON DATABASE api_db TO public;
	GRANT ALL ON DATABASE api_db TO api_user;
	ALTER ROLE api_user WITH CREATEDB CREATEROLE LOGIN;
END
$do$;
