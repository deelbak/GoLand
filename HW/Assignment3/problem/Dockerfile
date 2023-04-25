# # syntax=docker/dockerfile:1

# FROM golang:1.19

# # Set destination for COPY
# WORKDIR /app

# # Download Go modules
# COPY go.mod go.sum ./
# RUN go mod download

# # Copy the source code. Note the slash at the end, as explained in
# # https://docs.docker.com/engine/reference/builder/#copy
# COPY *.go ./

# # Build
# RUN CGO_ENABLED=0 GOOS=linux go build -o /problem

# # To bind to a TCP port, runtime parameters must be supplied to the docker command.
# # But we can (optionally) document in the Dockerfile what ports
# # the application is going to listen on by default.
# # https://docs.docker.com/engine/reference/builder/#expose
# 

# # Run
# CMD [ "/problem" ]
FROM golang:1.19

RUN go version
ENV GOPATH=/

WORKDIR /app

COPY ./ ./

RUN go mod download
RUN go build -o final-go-kbtu ./cmd/web
EXPOSE 8080
CMD ["./final-go-kbtu"]
