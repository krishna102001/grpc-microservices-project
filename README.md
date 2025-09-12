## gRPC Microservices Project

This is blog microservices project which was created to learn `Kafka` use.

### Services

- Api Gateway - Its the single entry point of all the routes and its convert REST API to gRPC call.
- Auth Service - Its handle the authentication part like login, signup and forget password.
- Blog Service - Its handle the CRUD operation of Blog service.
- Notification Service - Its handle the notification parts like sending mail or sms.

### Tech Stack

- Golang
- PostgreSQL
- Kafka
- gRPC
- docker

### Future Features

- [x] Authentication
- [x] CRUD Operation
- [ ] Retry Mechanism of notification services
- [ ] Upload of Images

### Steps To Run The Project

1. Clone the project
   ```bash
    git clone https://github.com/krishna102001/grpc-microservices-project.git
   ```
2. Create `.env` file in each services
3. Run the docker command
   ```bash
    docker compose up
   ```
