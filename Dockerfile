# Specifies a parent image
FROM golang:1.21.1-bullseye
 
# Creates an app directory to hold your app’s source code
WORKDIR /app
 
# Copies everything from your root directory into /app
COPY . .
 
# Installs Go dependencies
RUN go mod download
 
# Builds your app with optional configuration
RUN go build -o ./godocker
 
# Specifies the executable command that runs when the container starts
CMD [ "/godocker" ]

# # ** Check out this below for a better setup for prod **

# # This is a standard Dockerfile for building a Go app.
# # It is a multi-stage build: the first stage compiles the Go source into a binary, and
# #   the second stage copies only the binary into an alpine base.

# # -- Stage 1 -- #
# # Compile the app.
# FROM golang:1.12-alpine as builder
# WORKDIR /app
# # The build context is set to the directory where the repo is cloned.
# # This will copy all files in the repo to /app inside the container.
# # If your app requires the build context to be set to a subdirectory inside the repo, you
# #   can use the source_dir app spec option, see: <https://www.digitalocean.com/docs/app-platform/reference/app-specification-reference/>
# COPY . .
# RUN go build -mod=vendor -o bin/hello

# # -- Stage 2 -- #
# # Create the final environment with the compiled binary.
# FROM alpine
# # Install any required dependencies.
# RUN apk --no-cache add ca-certificates
# WORKDIR /root/
# # Copy the binary from the builder stage and set it as the default command.
# COPY --from=builder /app/bin/hello /usr/local/bin/
# CMD ["hello"]