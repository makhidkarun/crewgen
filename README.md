# crewgen

Generic 2d6 OGL NPC generators

## Basic usage

Copy npcGen.zip file into a new directory, and unzip it. Read the README.txt file.




## Building the npcGen binary

```
cd cmd/npcGen
go build -o npcGen main.go
```

## Cross-compiling 

```
make build
```

## Contributing

The datafiles can always use new material. Check out 
cmd/npcGen/data/careers.txt and cmd/npcGen/data/jobs.txt for places to start.
You can also customize those for your own campaign.

If you want to do some coding, Go is pretty simple, and there's a TODO list 
here for things that need to be worked on. 

## Using crewgen with Docker

Create a directory for the Docker build, and create the sub-directory "app".

### Dockerfile

```
FROM voidlinux/voidlinux:latest

RUN mkdir /app
WORKDIR /app
COPY app /app

EXPOSE 8080

CMD [ "/app/crewgen" ]

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

