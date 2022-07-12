# Env vars

SHIPA_GEN_MONGO_URI - Mongo DB uri, default: "mongodb://localhost:27017"
SHIPA_GEN_BACKEND_PORT - server port, default:  "8080"


# Development

## Run mongo db in container

    docker-compose up -d

## Stop mongo db

    docker-compose down

## Run server

    go run server.go


# Troubleshooting

In case you getting issue during building application like this

    # golang.org/x/sys/unix
    /Users/vmanilo/go/pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/syscall_darwin.1_13.go:25:3: //go:linkname must refer to declared function or variable
    /Users/vmanilo/go/pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/zsyscall_darwin_arm64.1_13.go:27:3: //go:linkname must refer to declared function or variable
    /Users/vmanilo/go/pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/zsyscall_darwin_arm64.1_13.go:40:3: //go:linkname must refer to declared function or variable
    /Users/vmanilo/go/pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/zsyscall_darwin_arm64.go:28:3: //go:linkname must refer to declared function or variable
    /Users/vmanilo/go/pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/zsyscall_darwin_arm64.go:43:3: //go:linkname must refer to declared function or variable
    /Users/vmanilo/go/pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/zsyscall_darwin_arm64.go:59:3: //go:linkname must refer to declared function or variable
    /Users/vmanilo/go/pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/zsyscall_darwin_arm64.go:75:3: //go:linkname must refer to declared function or variable
    /Users/vmanilo/go/pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/zsyscall_darwin_arm64.go:90:3: //go:linkname must refer to declared function or variable
    /Users/vmanilo/go/pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/zsyscall_darwin_arm64.go:105:3: //go:linkname must refer to declared function or variable
    /Users/vmanilo/go/pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/zsyscall_darwin_arm64.go:121:3: //go:linkname must refer to declared function or variable
    /Users/vmanilo/go/pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/zsyscall_darwin_arm64.go:121:3: too many errors

Suggestion is to run cmd: `go get -u golang.org/x/sys`