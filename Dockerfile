FROM golang:1.23
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /bin/service ./cmd/service
EXPOSE 8080
CMD ["/bin/service"]
