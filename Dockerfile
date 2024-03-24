FROM golang:1.21.6


# Set destination for COPY
WORKDIR /app

# Download Go modules (optional - either check in go module folder or download now)
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY . .

# Build
RUN go build cmd/main_server/main.go

ENV DB_HOST=172.17.0.2
ENV DB_PORT=5432
ENV DB_USER=postgres
ENV DB_PASS=123
ENV DB_NAME=test

EXPOSE 8888

# Run
CMD [ "./main" ]
