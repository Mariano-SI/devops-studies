# Alterado de 1.21 para 1.22 ou superior (ou use 'latest' para pegar a mais atual)
FROM golang:1.22-alpine AS builder

# O restante do arquivo permanece igual
WORKDIR /app
COPY go.mod ./
# COPY go.sum ./  <-- Se você não tiver esse arquivo, mantenha comentado
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o healthchecker main.go

# Estágio de Runtime
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/healthchecker .
CMD ["./healthchecker"]