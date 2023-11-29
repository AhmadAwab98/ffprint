# Go training 
Repository for the training of Go language.

## Description

The project contains:
- list all files and folders in a given path
- cache it for 5 minutes using redis
- docker-compose.yml file for
    - starting redis server
    - running the project

## Getting Started

### Execution

Continued commands from main readme.md

- Run the go mod tidy command to add any missing dependencies to the go.mod file
```
go mod tidy
```

- Run docker-compose up to start redis-server and run the project
```
docker-compose up
```

- On the Postman set the method to Get/Post, write localhost:8080/hello in URL field and in the body write
```
{
    "path": "/home/ahmadawab/go-training"
}
```
- Click send, you will see the result in the body tag of the output.



## Authors

Ahmad Awab
