FROM golang:1.20 AS builder

RUN apt-get update && apt-get install -y git

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN chmod +x generate-sha.sh && ./generate-sha.sh
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o demo-app


FROM scratch
WORKDIR /app
COPY --from=builder /app/demo-app /app/demo-app
COPY --from=builder /app/templates /app/templates
COPY --from=builder /app/static /app/static
COPY --from=builder /app/sha.txt /app/sha.txt

EXPOSE 8080
CMD ["/app/demo-app"]
