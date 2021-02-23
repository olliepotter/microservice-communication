#
# BUILD CONTAINER TO CREATE BINARY
#
FROM golang:latest AS builder
LABEL Service Two Builder

# Setup files
RUN mkdir /app
WORKDIR /app
ADD . .

# Get build args
ARG SERVICE_ONE_URL
ENV SERVICE_ONE_URL=${SERVICE_ONE_URL}

# Build the binary in container
RUN CGO_ENABLED=0 GOOS=linux go build main.go

#
# BUILD CONTAINER TO RUN BINARY
#
FROM alpine:latest AS production
LABEL Service Two

# Get build args
ARG SERVICE_ONE_URL
ENV SERVICE_ONE_URL=${SERVICE_ONE_URL}
COPY --from=builder /app .

# Expose port
EXPOSE 4002

# Run service as binary version
CMD [ "./main" ]
