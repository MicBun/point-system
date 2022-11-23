# point-system
a transactional database system using golang, mysql, and docker

PreInstall
1. Docker https://docs.docker.com/desktop/install/windows-install/

How to use
1. Clone this Repository using `git clone`
2. Run Docker Desktop
3. use CMD Command on this project root
    1. `docker-compose build`
    2. `docker-compose up`
4. access http://localhost:8080/swagger/index.html#/
5. The app is ready to use.
---
# Database
Database included and can be accessed at http://localhost:5003

---
# Structure Folder
```
root
-db // Contain Db Configuration file
-controller // Contains Controllers
-docker // Contain app.dockerfile for dockerise
-docs // auto-generated file by Swagger
-middleware // middlewares for auth
-model // Contains Models for database structure
-route // Contains Routes
-util // Conatins utility file for multiple uses
.env // an env file
.docker-compose.yml // docker related file
go.mod // go mod
main.go // the main.go
README.md // this Readme file
```
---
# Credential
1. phpmyadmin
```
username: root
password: root
```