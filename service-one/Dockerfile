#
# BUILD CONTAINER TO CREATE BINARY
#
FROM golang:latest AS builder
LABEL Service One Builder

# Setup files
RUN mkdir /app
WORKDIR /app
ADD . .

# Build the binary in container
RUN CGO_ENABLED=0 GOOS=linux go build main.go



#
# BUILD CONTAINER TO RUN BINARY
#
FROM alpine:latest AS production
LABEL Service One

COPY --from=builder /app .

# Expose port
EXPOSE 4000

# Run service as binary version
CMD [ "./main" ]
