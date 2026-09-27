		# Comandos para facilitar o gerenciamento do projeto

.PHONY: help build up down restart logs clean

# Mostrar ajuda
help:
	@echo "Comandos disponíveis:"
	@echo "  build     - Construir as imagens Docker"
	@echo "  up        - Subir os serviços"
	@echo "  down      - Parar os serviços"
	@echo "  restart   - Reiniciar os serviços"
	@echo "  logs      - Ver logs dos serviços"
	@echo "  clean     - Limpar containers e volumes"

# Construir as imagens
build:
	docker-compose build

# Subir os serviços
up:
	docker-compose up -d

# Parar os serviços
down:
	docker-compose down

# Reiniciar os serviços
restart:
	docker-compose restart

# Ver logs
logs:
	docker-compose logs -f

# Limpar tudo (cuidado: remove volumes!)
clean:
	docker-compose down -v
	docker system prune -f