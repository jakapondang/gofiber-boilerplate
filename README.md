# Boileplate 0.0.1

## Overview


## Requirements

- **Docker Compose / Desktop**
- **Go Version**: 1.22

## Getting Started

1. **Clone the Repository**

   ```sh
   git clone <repository-url>
   cd <repository-directory>
2. **Build and Run the Application**

   ```sh
   docker-compose up --build

3. **Check Postman Collection**


4. **Run on your local**

   Update .env
   ```sh
   #url: postgres://admin:admin@db:5432/goboilerplate
   url: postgres://admin:admin@localhost:5432/goboilerplate

4. **Problem**
   1. If error ..failed to initialize database, got error 
```
   failed to connect to `user=root database=goboilerplate`:hostname resolving error:lookup db: no such host lookup db: no such host 
   ```
   Try changing the config file database.url from db to localhost

5. **Notes**

## File Structure

Within the download you'll find the following directories and files:

```
root_dir/
├── configs/
├── internal/
├───└──v1
│      ├── app/
│      │   ├── dto/
│      │   ├── handlers/
│      │   ├── usecases/
│      ├── domain/
│      │   ├── models/
│      │   ├── repositories/
│      │   ├── services/
│      └── Interface/
│         ├── grpc/
│         └── http/
│             ├── middlewares/
│             └── routes/
├── pkg/
│   ├── infra/
│   │   ├── config/
│   │   └── database/
│   │       └── sql/
│   ├── msg/
│   └── utils/
│        ├── auth/
│        ├── jwt/
│        └── msg/
├── scripts/
├── main.go
└── go.mod

```
## Thank you
for reviewing this project. Please feel free to reach out with any questions or feedback regarding this project.