
FROM golang:latest
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build -o MediumUZ ./command/main.go
ENV PORT 8080
CMD ["./MediumUZ"]
