FROM alpine:latest

WORKDIR /app

# Copy the pre-built binary into the container
COPY bin/linux_amd64/api ./api

EXPOSE 4000

ENTRYPOINT ["./api"]