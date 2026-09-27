FROM node:24-alpine

WORKDIR /app_frontend

COPY package*.json ./
RUN npm install

# Copiar código fonte
COPY . .

# Expor a porta
EXPOSE 3000

# Comando para executar a aplicação
CMD ["npm", "start:prod"]