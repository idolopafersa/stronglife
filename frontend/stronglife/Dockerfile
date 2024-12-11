# Usa la imagen oficial de Node.js 22.10
FROM node:22.10-slim

# Establece el directorio de trabajo
WORKDIR /app

# Copia todo el código del proyecto
COPY . .

# Expone el puerto para Expo
EXPOSE 19000 19001 19002 8081

# Comando por defecto para iniciar Expo
CMD ["npm", "start"]
