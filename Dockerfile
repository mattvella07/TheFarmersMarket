FROM golang
WORKDIR /
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o farmersMarket main.go
CMD ["./farmersMarket"]
