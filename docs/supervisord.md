# Supervisord

Below is an example of a supervisor configuration for running allwrite on
startup.

```bash
[program:allwrite]
command=/usr/local/bin/allwrite-docs s
autostart=true
autorestart=true
startretries=10
user=root
directory=/etc/allwrite
environment=
        USER="root",
        HOME="/root",
        ACTIVE_DIR="...",
        CLIENT_SECRET="/etc/allwrite/client_secret.json",
        STORAGE="postgres",
        PG_USER="allwrite",
        PG_PASS="allwrite",
        PG_DB="allwrite",
        PG_HOST="localhost",
        PORT=":443",
        DOMAIN="my-domain.com",
        CERTBOT_EMAIL="my@email.com",
        FREQUENCY="300000"
redirect_stderr=true
stdout_logfile=/var/log/supervisor/allwrite.log
stdout_logfile_maxbytes=50MB
stdout_logfile_backups=10
```
