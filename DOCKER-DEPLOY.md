# FitStore - Deploy com Docker

## 🐳 Deploy Rápido com Docker Compose

### Pré-requisitos na EC2
```bash
# Instalar Docker
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh

# Instalar Docker Compose
sudo curl -L "https://github.com/docker/compose/releases/download/v2.21.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

# Adicionar usuário ao grupo docker
sudo usermod -aG docker $USER
```

### Como usar

1. **Clone o projeto na EC2:**
```bash
git clone <seu-repositorio>
cd FitStore
```

2. **Configure as variáveis de ambiente:**
```bash
# Edite o arquivo .env com suas configurações
nano .env
```

3. **Suba a aplicação:**
```bash
# Opção 1: Usando Make
make up

# Opção 2: Docker Compose direto
docker-compose up -d
```

4. **Verificar se está funcionando:**
```bash
# Ver logs
make logs
# ou
docker-compose logs -f

# Verificar status
docker-compose ps
```

### Comandos Úteis

```bash
# Parar a aplicação
make down

# Reiniciar serviços
make restart

# Rebuild das imagens
make build

# Limpar tudo (cuidado!)
make clean
```

### Estrutura dos Serviços

- **Backend**: Roda na porta 80
- **PostgreSQL**: Roda na porta 5432
- **Volumes**: Os dados do banco são persistidos

### Troubleshooting

```bash
# Ver logs específicos
docker-compose logs backend
docker-compose logs postgres

# Entrar no container
docker-compose exec backend sh
docker-compose exec postgres psql -U fitstore_user -d fitstore

# Verificar rede
docker network ls
```