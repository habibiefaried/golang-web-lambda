FROM golang:1.22-alpine as builder
RUN apk add --no-cache gcc musl-dev
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -a -tags timetzdata -installsuffix cgo -o main .
RUN chmod +x main

FROM gcr.io/distroless/static
COPY --from=builder /app/main /main
ENTRYPOINT [ "/main" ]
