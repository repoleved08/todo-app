# Products, Todo, and Auth REST API

This is a combined REST API for managing products and todo tasks, with JWT-based authentication and authorization. The API uses PostgreSQL for database management, Echo v4 as the web framework, and godotenv for environment variable management. It also includes image uploads for products, and static files are served using Nginx.

## Features

- **JWT Authentication**: User registration and login with secure JWT tokens.
- **Todo Management**: Create, read, update, and delete todo items.
- **Product Management**: Full CRUD operations for products, including image uploads.
- **PostgreSQL Integration**: Uses PostgreSQL as the database.
- **Nginx for Static Files**: Nginx serves static files for products' images.
  
## Tech Stack

- **Golang**: Backend API with Echo framework.
- **PostgreSQL**: Database for storing todos, products, and users.
- **JWT**: For secure authentication and authorization.
- **Echo v4**: Web framework used for routing.
- **godotenv**: For environment variable management.
- **Nginx**: For serving static files.

## API Endpoints

### Authentication

- `POST /api/auth/register`: Register a new user.
- `POST /api/auth/login`: Login and obtain a JWT token.

### Todos

- `POST /api/todo/create`: Create a new todo (JWT required).
- `GET /api/todos`: Get all todos.
- `GET /api/todos/:id`: Get a todo by ID.
- `PUT /api/todos/update/:id`: Update a todo by ID (JWT required).
- `DELETE /api/todos/delete/:id`: Delete a todo by ID (JWT required).

### Products

- `POST /api/products/create`: Create a new product (JWT required).
- `GET /api/products`: Get all products.
- `GET /api/products/:id`: Get a product by ID.
- `PUT /api/products/update/:id`: Update a product by ID (JWT required).
- `DELETE /api/products/delete/:id`: Delete a product by ID (JWT required).

### Static Files

- Product images are served via Nginx at `http://localhost/static/products`.

## Setup Instructions

### Prerequisites

- Golang 1.18+
- PostgreSQL
- Nginx

### Environment Variables

Set up your `.env` file with the following variables:

```bash
DB_HOST=localhost 
DB_PORT=5432 
DB_USER=your_db_user 
DB_PASSWORD=your_db_password 
DB_NAME=your_db_name 
JWT_SECRET=your_jwt_secret
```


### Install Dependencies

1. Clone the repository.
2. Run `go mod tidy` to install the required dependencies.

### Run the Application

1. Set up your PostgreSQL database and update the `.env` file accordingly.
2. Run the migrations and start the server:
    ```bash
    go run cmd/main.go
    ```
3. Nginx will handle static files. Configure it with the following block in your Nginx config:
    ```nginx
    server {
        listen 80;
        server_name localhost;

        location /static/ {
            root /path/to/your/project;
        }

        location / {
            proxy_pass http://localhost:8080;
            proxy_set_header Host $host;
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        }
    }
    ```

### Swagger Documentation

Swagger documentation is available at `http://localhost:8080/swagger/index.html`.

## Contact developer
techxtrasol.design@gmail.com
