# Estef.link

  Una app web construida con Go + Astro que permite al usuario acortar URLs mediante una autenticacion con Google.
  
 
  
  [![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
  [![Astro](https://img.shields.io/badge/Astro-BC52EE?style=for-the-badge&logo=astro&logoColor=white)](https://astro.build/)
  [![MongoDB](https://img.shields.io/badge/MongoDB-47A248?style=for-the-badge&logo=mongodb&logoColor=white)](https://www.mongodb.com/)
  
  ## Configuración del entorno
  
  Antes de iniciar el proyecto, es necesario configurar las siguientes variables en un archivo .env o en las variables de sistema:
  
  ```bash
  # Conexión a Base de Datos
  MONGODB_URI=<tu_uri_de_mongodb>
  
  # Configuración del Servidor
  PORT=":4040"
  
  # Google Auth (OAuth2)
  GOOGLE_CLIENT_ID=<tu_client_id>
  GOOGLE_CLIENT_SECRET=<tu_client_secret>
  ```
  
  ## Instalacion y Ejecución
  
  *Requisitos Previos*
  - Go >= 1.21
  - Node.js
  
  Para iniciar la aplicación web en modo desarrollo, ejecuta el siguiente comando en la terminal:
  
  ```bash
  go run main.go
  ```
  
  Para compilar y ejecutar la aplicación web en modo producción, sigue estos comandos:   
  
  ```bash
  # Compilar
  go build -o bin/estef.link main.go
  
  # Ejecutar
  ./bin/estef.link
  ```

## Arquitectura del Proyecto
- `/db` Conexion a MongoDB
- `/handlers` Manejo de Rutas y Controladores
- `/models` Definición de Estructuras de Datos
