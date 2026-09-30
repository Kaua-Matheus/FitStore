FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copiar go.mod e go.sum primeiro (para cache do Docker)
COPY go.mod go.sum ./

# Download das dependências
RUN go mod download

# Copiar código fonte
COPY . .

# Build da aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd

# Stage final com imagem mínima
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Copiar o binário do stage anterior
COPY --from=builder /app/app .

# Criar diretório para uploads
RUN mkdir -p uploads/banners uploads/Perfil uploads/Produtos

# Expor a porta
EXPOSE 80

# Comando para executar a aplicação
CMD ["./app"]