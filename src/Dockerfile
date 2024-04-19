FROM golang:1.22-alpine as builder
RUN apk add --no-cache gcc musl-dev
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -a -tags timetzdata -installsuffix cgo -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
RUN chmod +x ./main

CMD ["./main"]
