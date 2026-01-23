FROM golang:1.24.3
WORKDIR /app
RUN go install github.com/air-verse/air@latest
COPY ../go.mod ../go.sum ./
RUN go mod download
COPY . .
RUN go build -o ./goapp cmd/main.go
CMD ["air", "-c", ".air.toml"]