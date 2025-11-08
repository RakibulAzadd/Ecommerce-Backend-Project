
<xaiArtifact artifact_id="cd0be36d-d1db-445f-92d5-ec0c9eaba476" artifact_version_id="d6575379-b4b8-4a46-8d3d-8fadcf8ca20f" title="README.md" contentType="text/markdown">

# Ecommerce Project with Golang

This is a **backend application** for an e-commerce platform, built using **Golang**. The project provides **RESTful APIs** for managing users, products, orders, and authentication. It is designed to be **modular, secure, and easy to integrate** with any frontend application. This README includes an **ER diagram**, **Swagger documentation link**, and detailed **Postman instructions** to make the project structure and workflow clear.

---

## Features

### User Authentication
- User **signup** and **login**
- **JWT-based authentication**
- Password hashing using secure algorithms (bcrypt)
- Token-based session management

### Product Management
- **CRUD operations**: Create, Read, Update, Delete
- Product categorization and filtering
- Product search by name or category

### Order Management
- Place orders and track status
- View **order history**
- Secure checkout process

### Database
- **PostgreSQL** as the primary database
- Supports relational data and transactions
- Database migrations using **Golang-migrate**

### Error Handling
- Structured JSON error responses
- Input validation using **validator**
- Graceful handling of server errors

---

## Tech Stack

- **Language:** Golang
- **Database:** PostgreSQL
- **Authentication:** JWT (JSON Web Token)
- **API Testing:** Postman
- **Environment Variables:** Managed via `.env` file using **godotenv**
- **API Documentation:** Swagger (OpenAPI 3.0)
- **Database Migration:** Golang-migrate
- **Logging:** Zerolog

---

## Folder Structure

```
ecommerce-golang/
│
├── main.go                # Entry point of the application
├── go.mod                 # Go module file
├── go.sum                 # Go dependencies checksum
├── config/                # Configuration files
│   └── config.go          # Database and environment setup
├── controllers/           # HTTP request handlers
│   ├── auth.go
│   ├── product.go
│   └── order.go
├── models/                # Database models
│   ├── user.go
│   ├── product.go
│   └── order.go
├── routes/                # API routes
│   └── routes.go
├── middleware/            # Authentication & error handling
│   ├── auth.go
│   └── error.go
├── utils/                 # Helper functions
│   ├── jwt.go
│   └── validator.go

```

---

## Entity-Relationship (ER) Diagram

Below is the ER diagram representing the database schema:

```
[Users]                        [Products]
+----+                         +----+
| id | <--- PK                 | id | <--- PK
+----+                         +----+
| username |                   | name |
| email    |                   | description |
| password |                   | price       |
| created_at |                 | category    |
| updated_at |                 | stock       |
+----+                         | created_at  |
      |                        | updated_at  |
      |                        +----+
      |
      |                        [Orders]
      |                        +----+
      +---< one-to-many >-----| id | <--- PK
                               | user_id | <--- FK (Users)
                               | total_price |
                               | status      |
                               | created_at  |
                               | updated_at  |
                               +----+
                                    |
                                    |
                                    | [Order_Items]
                                    | +----+
                                    | | id | <--- PK
                                    | | order_id | <--- FK (Orders)
                                    | | product_id | <--- FK (Products)
                                    | | quantity   |
                                    | | unit_price |
                                    | +----+
```

- **Users**: Stores user information (id, username, email, password).
- **Products**: Stores product details (id, name, description, price, category, stock).
- **Orders**: Stores order details (id, user_id, total_price, status).
- **Order_Items**: Junction table for orders and products (id, order_id, product_id, quantity, unit_price).

---

## Installation & Setup

### 1. Clone the Repository

```bash
git clone https://github.com/RakibulAzadd/Ecommerce-Project-with-Golang.git
cd ecommerce-golang
```

### 2. Set Up PostgreSQL

Run the following SQL commands to create the database and user:

```sql
CREATE DATABASE ecommerce_db;
CREATE USER ecommerce_user WITH PASSWORD 'yourpassword';
GRANT ALL PRIVILEGES ON DATABASE ecommerce_db TO ecommerce_user;
```

### 3. Configure Environment Variables

Create a `.env` file in the project root:

```env
VERSION=1.0.0
SERVICE_NAME=ECOMMERCE
HTTP_PORT=4000

DB_HOST     = 127.0.0.1
DB_PORT     = 5432
DB_USER     = postgres
DB_PASSWORD = 12345
DB_NAME   = myapp
```

### 4. Install Dependencies

```bash
go mod tidy
```


### 6. Run the Server

```bash
go run main.go
```

The server will be available at `http://localhost:4000`.

---

## API Endpoints

### Authentication

| Method | Endpoint        | Description                  |
| ------ | --------------- | ---------------------------- |
| POST   | `/api/v1/signup`| Register a new user          |
| POST   | `/api/v1/login` | Login user and get JWT token |

### Products

| Method | Endpoint             | Description          |
| ------ | -------------------- | -------------------- |
| GET    | `/api/v1/products`   | List all products    |
| GET    | `/api/v1/products/:id` | Get product by ID    |
| POST   | `/api/v1/products`   | Create a new product (Admin) |
| PUT    | `/api/v1/products/:id` | Update product (Admin) |
| DELETE | `/api/v1/products/:id` | Delete product (Admin) |

### Orders

| Method | Endpoint           | Description          |
| ------ | ------------------ | -------------------- |
| POST   | `/api/v1/orders`   | Create a new order   |
| GET    | `/api/v1/orders`   | List all user orders |
| GET    | `/api/v1/orders/:id` | Get order details    |

---

## Swagger Documentation

The API is documented using **Swagger (OpenAPI 3.0)**. Access the Swagger UI at:

```
http://localhost:4000/docs
```

The Swagger YAML file is located at `docs/swagger.yaml`. You can also import it into tools like Swagger Editor.

---

## Postman Instructions

To test the APIs, use the provided Postman collection:

1. Import the `postman_collection.json` file (located in the project root) into Postman.
2. Set up a Postman environment with the following variables:
   - `BASE_URL`: `http://localhost:4000`
   - `JWT_TOKEN`: (Set this after logging in via the `/api/v1/login` endpoint)
3. Use the collection to test endpoints:
   - **Authentication**: Test signup and login to retrieve a JWT token.
   - **Products**: Test CRUD operations (use JWT for admin routes).
   - **Orders**: Test order creation and retrieval (requires JWT).

Sample Postman request for login:

```json
POST http://localhost:4000/api/v1/login
Content-Type: application/json

{
  "email": "rakibulazad@example.com",
  "password": "12345678"
}
```

---

## Contributing

Contributions are welcome! Follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/feature-name`).
3. Make your changes.
4. Commit your changes (`git commit -m 'Add some feature'`).
5. Push to the branch (`git push origin feature/feature-name`).
6. Open a Pull Request.

---

## License

This project is licensed under the **MIT License**.

---

## Contact

**Rakibul Azad**  
Email: `rakibulazad4049@example.com`  
Contant: 01611580264

---

## Future Improvements

- Add **payment gateway integration** (e.g., Stripe).
- Implement **role-based access control** (admin/user).
- Add **unit and integration tests** using Go’s testing package.
- Optimize database queries for better performance.
- Add **caching** using Redis for frequently accessed data.
- Implement **rate limiting** for API endpoints.

</xaiArtifact>