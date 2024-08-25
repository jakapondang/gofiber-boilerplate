# Boileplate 0.0.1

## Overview


## Ingredients

- **Deployment**
   - Docker & Docker compose
- **Go Version**: 1.22
   - Framework : Fiber V3 https://docs.gofiber.io/next/whats_new/
   - ORM : https://gorm.io/
   - Config : https://github.com/spf13/viper
   - Logs :
      - https://github.com/sirupsen/logrus
      - https://gopkg.in/natefinch/lumberjack.v2
   - Mail : https://gopkg.in/gomail.v2
   - Validator : github.com/go-playground/validator/v10

## Getting Started

1. **Clone the Repository**
   - Go to your working dir
      ```sh
      git clone <repository-url>
      cd <repository-directory>
       ```
2. **Build Application**
   1. Run this command on your CLI
      ```sh 
      docker-compose up --build -d
       ```
   2. Make sure your container running on docker environment .
      ```sh 
      docker ps
      ```
3. **Run local**
   1. Stop your API container.
   2. Update your configs/config.yaml
      ```sh
        #url: postgres://admin:admin@db:5432/goboilerplate
        url: postgres://admin:admin@localhost:5432/goboilerplate
        ```
   2. go run main.go
   3. Problem
      - If error ..failed to initialize database got error below ,Check / try changing the config file database.url from db to localhost
         ```sh
          failed to connect to `user=root database=goboilerplate`:hostname resolving error:lookup db: no such host lookup db: no such host 
         ```

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