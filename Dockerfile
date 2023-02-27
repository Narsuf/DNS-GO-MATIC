FROM golang:1.19-alpine AS build_base
WORKDIR /tmp/dns-go-matic
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
# RUN CGO_ENABLED=0 go test -cover ./...
RUN go build -o ./out/dns-go-matic .

FROM alpine:3.9
RUN apk add --no-cache ca-certificates
COPY --from=build_base /tmp/dns-go-matic/out/dns-go-matic .
COPY --from=build_base /tmp/dns-go-matic/app.env .
CMD ["/dns-go-matic"]
