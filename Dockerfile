FROM golang:alpine
WORKDIR /build
COPY . .
RUN go build -o storage main.go
EXPOSE 8000
CMD ["./storage"]