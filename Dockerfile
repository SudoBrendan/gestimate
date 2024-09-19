################################################################################
# Use an official Golang image as the base image for building
################################################################################
FROM golang:1.21-alpine as build

# Install dependencies
RUN apk add --no-cache make

# Build
WORKDIR /app
COPY . .
RUN make build

################################################################################
# Use a minimal base image for the final stage
################################################################################
FROM alpine:latest

WORKDIR /root/
COPY --from=build /app/gestimate .

# Command to run the app
ENTRYPOINT ["./gestimate"]
