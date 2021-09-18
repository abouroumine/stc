CREATE
DATABASE stc_auth;
CREATE
DATABASE stc_cc;
CREATE
USER ayoubbouroumine WITH PASSWORD 'password';
GRANT ALL PRIVILEGES ON DATABASE
"stc_auth" to ayoubbouroumine;
GRANT ALL PRIVILEGES  ON DATABASE
"stc_cc" to ayoubbouroumine;