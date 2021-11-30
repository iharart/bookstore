# Simple Golang REST API (CRUD) using GORM and Gorilla Mux with unit testing

*by Ihar Artsiomenka*
### Project description
&nbsp;&nbsp;&nbsp;There is a simple REST API (CRUD) project for book store on golang with using docker compose and two docker containers - 
for application and for mysql database
### <center>Directories structure</center>

    app/configs/ - database connection data 
    app/database/ - database layer
    app/handler/ - database/api handlers
    app/router/ - middleware layer
    app/router/ - unit tests
    scrips - dockerfile and script to create database 
    and populate it inside docker container
## Installing

Setup golang: 

https://go.dev/doc/install

Setup Docker: 

https://docs.docker.com/get-docker/

Clone the source down to your machine

git clone git://github.com/iharart/bookstore.git

## Running
1. Go to project root directory bookstore and firstly run from the console

`make app_build`

it will set up necessary dependencies and will try to run the project

2. If project wasn't started run

 `make app_start`

3. To stop application run

`make app_stop`

4. To run test 

`make test_run`
