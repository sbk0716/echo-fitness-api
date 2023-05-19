# 1. Project Overview
* This is a project for an app called `fitness-api`.

## (1)App Features
* This app is able to use below function.

### User Story
* Create a user with a username and password.
* Create a measurement with user_id, weight and body_fat.



## (2)Project Structure
```sh
.
├── Dockerfile
├── Dockerfile.production
├── Makefile
├── README.md
├── cmd
│   ├── handlers
│   │   ├── handelUsers.go
│   │   ├── handleMeasurements.go
│   │   ├── middleware.go
│   │   └── rootHandler.go
│   ├── models
│   │   └── user.go
│   ├── repositories
│   │   ├── measurementsDb.go
│   │   └── userDb.go
│   └── storage
│       └── db.go
├── docker-compose.production.yml
├── docker-compose.yml
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── main.go
└── postgres
    └── init
        └── init.sql
```

# 2. Usage
## 2.1. Start a postgres server and create a database.

```sh
% ls | grep docker-compose.yml
docker-compose.yml
% docker-compose up -d
% docker ps
CONTAINER ID   IMAGE                     COMMAND                  CREATED              STATUS              PORTS                    NAMES
6d56f7dd3feb   fitness-api-fitness_api   "go run main.go"         About a minute ago   Up About a minute   0.0.0.0:8000->8000/tcp   fitness_api
e8372c8c7c2c   postgres:14               "docker-entrypoint.s…"   About a minute ago   Up About a minute   0.0.0.0:5432->5432/tcp   fitness_pg
% 
% docker exec -it fitness_pg /bin/bash
root@1a4b66aaadc4:/# printenv | grep PASS
POSTGRES_PASSWORD=*******
root@1a4b66aaadc4:/# psql --username super-user --dbname fitness_app
psql (14.7 (Debian 14.7-1.pgdg110+1))
Type "help" for help.

fitness_app=# \l
                                      List of databases
   Name    |   Owner    | Encoding |  Collate   |   Ctype    |       Access privileges       
-----------+------------+----------+------------+------------+-------------------------------
 fitness_app | super-user | UTF8     | en_US.utf8 | en_US.utf8 | 
 postgres  | super-user | UTF8     | en_US.utf8 | en_US.utf8 | 
 template0 | super-user | UTF8     | en_US.utf8 | en_US.utf8 | =c/"super-user"              +
           |            |          |            |            | "super-user"=CTc/"super-user"
 template1 | super-user | UTF8     | en_US.utf8 | en_US.utf8 | =c/"super-user"              +
           |            |          |            |            | "super-user"=CTc/"super-user"
(4 rows)

fitness_app=# show search_path;
   search_path   
-----------------
 "$user", public
(1 row)

fitness_app=# SET search_path = "$user", public, private;
SET
fitness_app=# show search_path;
       search_path        
--------------------------
 "$user", public, private
(1 row)

fitness_app=# \dt;
              List of relations
 Schema  |     Name     | Type  |   Owner    
---------+--------------+-------+------------
 private | measurements | table | super-user
 private | users        | table | super-user
(2 rows)

fitness_app=# \q
root@1a4b66aaadc4:/# exit
exit
% 
```

## 2.2. Run the app(local)

```sh
% make up
docker-compose up
[+] Running 2/2
 ⠿ Container fitness_pg   Recreated                                         0.1s
 ⠿ Container fitness_api  Recreated                                         0.1s
Attaching to fitness_api, fitness_pg
fitness_pg   | 
fitness_pg   | PostgreSQL Database directory appears to contain a database; Skipping initialization
fitness_pg   | 
fitness_pg   | 2023-05-19 21:00:04.340 UTC [1] LOG:  starting PostgreSQL 14.8 (Debian 14.8-1.pgdg110+1) on aarch64-unknown-linux-gnu, compiled by gcc (Debian 10.2.1-6) 10.2.1 20210110, 64-bit
fitness_pg   | 2023-05-19 21:00:04.340 UTC [1] LOG:  listening on IPv4 address "0.0.0.0", port 5432
fitness_pg   | 2023-05-19 21:00:04.340 UTC [1] LOG:  listening on IPv6 address "::", port 5432
fitness_pg   | 2023-05-19 21:00:04.346 UTC [1] LOG:  listening on Unix socket "/var/run/postgresql/.s.PGSQL.5432"
fitness_pg   | 2023-05-19 21:00:04.351 UTC [28] LOG:  database system was shut down at 2023-05-19 21:00:01 UTC
fitness_pg   | 2023-05-19 21:00:04.356 UTC [1] LOG:  database system is ready to accept connections
fitness_api  | filePath: ".env.local"
fitness_api  | Successfully connected to database
fitness_api  | 
fitness_api  |    ____    __
fitness_api  |   / __/___/ /  ___
fitness_api  |  / _// __/ _ \/ _ \
fitness_api  | /___/\__/_//_/\___/ v4.10.0
fitness_api  | High performance, minimalist Go web framework
fitness_api  | https://echo.labstack.com
fitness_api  | ____________________________________O/_______
fitness_api  |                                     O\
fitness_api  | ⇨ http server started on [::]:8080
...
...
```

## 2.3. Run the app(production)

```sh
% make up/prod
docker-compose -f docker-compose.production.yml up
[+] Running 2/2
 ⠿ Container fitness_pg   Recreated                                         0.1s
 ⠿ Container fitness_api  Recreated                                         0.1s
Attaching to fitness_api, fitness_pg
fitness_pg   | 
fitness_pg   | PostgreSQL Database directory appears to contain a database; Skipping initialization
fitness_pg   | 
fitness_pg   | 2023-05-19 21:01:06.053 UTC [1] LOG:  starting PostgreSQL 14.8 (Debian 14.8-1.pgdg110+1) on aarch64-unknown-linux-gnu, compiled by gcc (Debian 10.2.1-6) 10.2.1 20210110, 64-bit
fitness_pg   | 2023-05-19 21:01:06.053 UTC [1] LOG:  listening on IPv4 address "0.0.0.0", port 5432
fitness_pg   | 2023-05-19 21:01:06.053 UTC [1] LOG:  listening on IPv6 address "::", port 5432
fitness_pg   | 2023-05-19 21:01:06.056 UTC [1] LOG:  listening on Unix socket "/var/run/postgresql/.s.PGSQL.5432"
fitness_pg   | 2023-05-19 21:01:06.062 UTC [27] LOG:  database system was shut down at 2023-05-19 21:00:45 UTC
fitness_pg   | 2023-05-19 21:01:06.066 UTC [1] LOG:  database system is ready to accept connections
fitness_api  | filePath: ".env.production"
fitness_api  | Successfully connected to database
fitness_api  | 
fitness_api  |    ____    __
fitness_api  |   / __/___/ /  ___
fitness_api  |  / _// __/ _ \/ _ \
fitness_api  | /___/\__/_//_/\___/ v4.10.0
fitness_api  | High performance, minimalist Go web framework
fitness_api  | https://echo.labstack.com
fitness_api  | ____________________________________O/_______
fitness_api  |                                     O\
fitness_api  | ⇨ http server started on [::]:8080
...
...
```


# 3. Check the operation of each API endpoint
## 3.1. `POST /users`
* Create a user with a username and password.

```sh
% curl -s -H "Content-Type: application/json" \
    -X POST \
    -d '{"name": "test user","email": "test_user@mail.com","password": "1234567fjoafhouf"}' \
    http://localhost:8080/users | jq -r '.'
{
  "id": 7,
  "name": "test user",
  "email": "test_user@mail.com",
  "password": "1234567fjoafhouf",
  "created_at": "2023-05-20T06:53:42.079452Z",
  "updated_at": "2023-05-20T06:53:42.079452Z"
}
%
```

## 3.2. `POST /measurements`
* Create a measurement with user_id, weight and body_fat.

```sh
% curl -s -H "Content-Type: application/json" \
    -X POST \
    -d '{"user_id": 1,"weight": 80,"height": 180,"body_fat": 20}' \
    http://localhost:8080/measurements | jq -r '.'
{
  "id": 1,
  "user_id": 1,
  "weight": 80,
  "height": 180,
  "body_fat": 20,
  "created_at": "0001-01-01T00:00:00Z"
}
% 
```


## 3.3. `PUT /users`
* Update a user with a username and password.

```sh
% curl -s -H "Content-Type: application/json" \
    -X PUT \ 
    -d '{"name": "updated test01","email": "test01@mail.com","password": "password"}' \
    http://localhost:8080/user/1 | jq -r '.'
{
  "id": 1,
  "name": "updated test01",
  "email": "test01@mail.com",
  "password": "password",
  "created_at": "0001-01-01T00:00:00Z",
  "updated_at": "0001-01-01T00:00:00Z"
}
```