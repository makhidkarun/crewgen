# crewgen
Web based generic ship crew generator



## Building the binary

```
go build -o crewgen main.go
```

## Using crewgen with Docker

Create a directory for the Docker build, and create the subdirectory "app".
### Dockerfile

```
FROM voidlinux/voidlinux:latest

RUN mkdir /app
WORKDIR /app
COPY app /app

EXPOSE 8080

ENTRYPOINT [ "/app/crewgen" ]

```

Copy the crewgen binary, the "data" and "web" directories to the working "app"
directory. You should end up with a file structure like this:

```
├── Dockerfile
└── app
    ├── crewgen
    ├── data
    │   └── names.db
    └── web
        ├── crew.tmpl
        ├── form.tmpl
        ├── htmlClose.tmpl
        ├── htmlOpen.tmpl
        ├── layout.tmpl
        └── person.tmpl

```

Build the binary.

```
docker build -t crewgen .
```

Run detached, on a previously defined network ("frontend", 172.18.0.XXX).

```
docker run --detach --name crewgen -p 8080:8080 --network frontend --ip 172.18.0.24  crewgen:latest 
```

