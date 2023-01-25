## Build WebApp ##

# Use latest golang base image
FROM golang:latest as APIBuilder

ARG ENV_FILE=.env.example

# Set the current working directory & copy source to it inside the container
WORKDIR /app

# Download all dependencies defined in the go.mod & go.sum files
## Done separately to leverage caching layers
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY ${ENV_FILE} .env
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/package/api cmd/api/main.go

# Pre-Requisit for pain-free deployment to EC2.
# Generate self-signed certificates (only for local dev environment)
#RUN go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits 1024 --host 127.0.0.1,::1,localhost --ca --start-date "Jan 1 00:00:00 1970" --duration=100000h

## Run WebApp ##

FROM alpine:latest
RUN apk --no-cache add ca-certificates

ARG VERSION="dev"

LABEL Name="xyz-api" \
      Version="$VERSION" \
      Vendor="decantera.dev" \
      Descrition="API" \
      Maintainer="decantera.dev Maintenance Team <anesovic@decantera.dev>"

WORKDIR /run/app

# Copy pre-build binary file from the 1st stage
COPY --from=APIBuilder /app/build/package/api .
COPY --from=APIBuilder /app/views ./views

# Expose 9450 port to the outside world
EXPOSE 9540

# Command to run the binary
CMD ["./main"]
