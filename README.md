If error

failed to initialize database, got error failed to connect to `user=root database=goboilerplate`:
hostname resolving error: lookup db: no such host
lookup db: no such host

Try changing the config file database.url from db to localhost

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
   #url: postgres://root:Iinvite3@db:5432/goboilerplate
   url: postgres://root:Iinvite3@localhost:5432/goboilerplate

4. **Issue**

## Thank you
for reviewing this project. Please feel free to reach out with any questions or feedback regarding this project.