FROM golang:latest
RUN mkdir /app
ADD . /app
WORKDIR /app
COPY . ./
RUN go build  ./command/main.go
RUN go mod tidy
CMD ["./main"]