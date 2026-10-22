#!/bin/sh
set -eu
# No shell trace; credentials come from an ignored generated local file.
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname postgres --set=ip="$IDENTITY_PASSWORD" --set=gp="$GAME_PASSWORD" <<'SQL'
CREATE ROLE identity_owner LOGIN PASSWORD :'ip';
CREATE ROLE game_owner LOGIN PASSWORD :'gp';
CREATE DATABASE identity OWNER identity_owner;
CREATE DATABASE game OWNER game_owner;
REVOKE ALL ON DATABASE identity FROM PUBLIC;
REVOKE ALL ON DATABASE game FROM PUBLIC;
GRANT CONNECT ON DATABASE identity TO identity_owner;
GRANT CONNECT ON DATABASE game TO game_owner;
SQL
