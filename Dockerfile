FROM golang:latest
WORKDIR /app
COPY . .
RUN  go build -o groupie ./cmd/
CMD ["./groupie"]