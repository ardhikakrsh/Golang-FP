# build stage
FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# RUN
FROM alpine
WORKDIR /app
COPY --from=builder /app/main .
COPY config.yaml .
EXPOSE 8000
CMD [ "/app/main" ]