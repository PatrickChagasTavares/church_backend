FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN GOFLAGS=-buildvcs=false CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o church_backend
CMD ["/app/church_backend"]