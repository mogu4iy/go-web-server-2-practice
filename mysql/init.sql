CREATE DATABASE IF NOT EXISTS web_server_db;
CREATE USER 'web-server'@'%' IDENTIFIED WITH mysql_native_password BY 'bAssKLjmYnQf3!bJGy6BJ@';
GRANT ALL PRIVILEGES ON web_server_db.* TO 'web-server'@'%';
FLUSH PRIVILEGES;