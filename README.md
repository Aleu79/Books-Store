# 📚 Practica-go

Backend escrito en **Go (Golang)** siguiendo buenas prácticas de arquitectura limpia (Clean Architecture) y principios **SOLID**.  
El objetivo del proyecto es implementar un sistema modular y escalable para gestionar usuarios y libros, con separación clara de capas y responsabilidades. Inicialmente, lo empecé como un CRUD para prepararme para una entrevista técnica, pero decidí completarlo y hacerlo un proyecto más completo y funcional.

---

## 🧠 Estructura del proyecto

```
practica-go/
├── 📁 internal/
│   ├── 📁 model/
│   │   ├── 📄 book.go       # struct Book
│   │   └── 📄 user.go       # struct User (Email, Password, Role, etc.)
│   │
│   ├── 📁 store/            # capa de persistencia (repositorios SQL)
│   │   ├── 📄 store.go      # inicializa DB y agrupa repositorios
│   │   ├── 📁 books/
│   │   │   └── 📄 book_store.go   # CRUD libros
│   │   └── 📁 users/
│   │       └── 📄 user_store.go   # CRUD usuarios, búsqueda por email, etc.
│   │
│   ├── 📁 service/          # capa de negocio
│   │   ├── 📁 books/
│   │   │   └── 📄 service.go      # lógica de libros
│   │   └── 📁 users/
│   │       └── 📄 service.go      # login, registro, validaciones, cookies
│   │
│   ├── 📁 transport/        # capa HTTP (handlers)
│   │   ├── 📄 utils.go           # helpers: writeJSON, parseJSON, manejo de errores
│   │   ├── 📁 books/
│   │   │   ├── 📄 handler.go     # CRUD de libros
│   │   │   ├── 📄 search.go      # búsqueda de libros
│   │   │   └── 📄 exists.go      # comprobar existencia
│   │   └── 📁 users/
│   │       ├── 📄 handler.go     # registro, login, logout
│   │       └── 📄 auth.go        # validación de sesiones
│   │
│   └── 📁 middleware/       # middlewares HTTP
│       ├── 📄 auth.go       # valida cookie o token JWT
│       └── 📄 logging.go    # logs de requests/responses
│
├── 📁 security/             # funciones de seguridad reutilizables
│   ├── 📄 cookies.go        # manejo seguro de cookies de sesión
│   └── 📄 hash.go           # hashing de contraseñas con bcrypt
│
├── 📁 tests/                # pruebas unitarias y de integración
│   ├── 📁 books/
│   │   └── 📄 book_service_test.go
│   └── 📁 users/
│       └── 📄 user_service_test.go
│
├── 📄 .gitignore
├── 📄 go.mod
├── 📄 main.go               # entrypoint: inicializa DB, servicios y servidor HTTP
├── 📄 README.md
├── 📄 LICENSE 
```

### 🧩 Explicación rápida

- **📁 model/** → Define las estructuras (`Book`, `User`) que representan las entidades del dominio.
- **📁 store/** → Capa de acceso a datos (repositorios SQL), separada por dominio (`books`, `users`).
- **📁 service/** → Contiene la lógica de negocio (reglas, validaciones, login, etc.).
- **📁 transport/** → Capa de presentación HTTP, donde se definen los endpoints REST.
- **📁 middleware/** → Funcionalidades transversales como autenticación o logging.
- **📁 security/** → Hash de contraseñas, cookies seguras y funciones de seguridad reutilizables.
- **📁 tests/** → Pruebas separadas por dominio (unitarias y de integración).
- **📄 main.go** → Inicializa la base de datos, servicios y servidor HTTP.

---

## 🚀 Requisitos previos

Antes de comenzar, asegúrate de tener instaladas las siguientes herramientas:

- 🟢 **Go 1.22+** – Lenguaje principal del proyecto
- 🟢 **Git** – Para control de versiones
- 🟢 **SQLite / PostgreSQL / MySQL** – Según la base de datos que uses
- 🟢 **curl** o **Postman** – Para probar los endpoints
- ⚙️ **make** (opcional) – Para automatizar comandos comunes


---
## ⚙️ Instalación y ejecución

### 🐧 Debian / Ubuntu / Linux

1. **Clonar el repositorio**
   ```bash
   git clone https://github.com/tuusuario/practica-go.git
   cd practica-go
2. **Instalar Go (si no lo tenés)**
    ```bash
    sudo apt update
    sudo apt install golang-go
3. **Configurar variables de entorno (opcional)**
   ```bash
    export PORT=8080
    export DB_URL="postgres://user:password@localhost:5432/practica?sslmode=disable"
4. **Descargar dependencias**
    ```bash
    go mod tidy
5. **Ejecutar el servidor**
    ```bash
    go run main.go
6. **Probar el backend**
    ```bash
    curl http://localhost:8080/api/
    ```
    o bien usar Postman para probar los endpoints /api/books y /api/users.

---

## 🍎MacOS

1. **Instalar Go (si no lo tenés)**
    ```bash
    brew install go
2. **Clonar el proyecto**
    ```bash   
    git clone https://github.com/tuusuario/practica-go.git
    cd practica-go
3. **Configurar variables de entorno (opcional)**
    ```bash
    export PORT=8080
    export DB_URL="sqlite3://./data.db"
4. **Instalar dependencias**
    ```bash    
    go mod tidy
5. **Ejecutar el proyecto**
    ```bash
    go run main.go
6. **Probar que el servidor está corriendo**
    ```bash
    curl http://localhost:8080/api/health

## ✅ Ejecutar Test
    go test ./tests/..

**Esto ejecutará todas las pruebas unitarias e integradas definidas en la carpeta tests/.**




## 📦 Tecnologías usadas y buenas practicas

- 🟢 Go – Lenguaje principal.
- 🌐 net/http – Servidor web nativo.
- 🗄️ database/sql – Abstracción de base de datos.
- 🏗️ Clean Architecture – Separación de capas y responsabilidades.
- Pronto seguir agrande mas!!




