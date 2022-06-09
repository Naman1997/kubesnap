FROM golang:1.18-alpine as builder
WORKDIR /app
COPY cmd/kube-snap/ ./
RUN go mod download
# RUN go get -d -v ./...
# RUN go vet -v
# RUN go test -v
RUN CGO_ENABLED=0 go build -o app .

FROM alpine:latest
RUN apk add --no-cache git
COPY --from=builder /app/app ./
CMD ["./app"]