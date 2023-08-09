FROM golang:latest
WORKDIR /app
COPY . .
RUN  go build -o groupie
CMD ["./groupie"]