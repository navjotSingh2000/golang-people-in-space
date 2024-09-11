FROM golang:1.23

# Set destination for COPY
WORKDIR /app

# get the air live reload
RUN go install github.com/air-verse/air@latest

# Download Go modules
COPY go.mod ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY *.go ./
COPY templates/ ./templates/

# copy .air.toml file for air live reload
COPY .air.toml ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8080

# Run
# CMD ["/docker-gs-ping"]
CMD ["air", "-c", ".air.toml"]