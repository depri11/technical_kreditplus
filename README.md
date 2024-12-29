# Tes Kredit Plus Loan Tenor System
This project is a loan tenure system designed with a microservice architecture to ensure scalability. It uses gRPC for communication between services and can be managed using Docker Compose.

# Overview
The Loan Tenure System includes several microservices, each responsible for different aspects of the application, such as customer management, loan limit management, and transaction processing. By leveraging microservices and gRPC, the system ensures efficient communication, scalability, and maintainability.

# Architecture Diagram

# Flowchart Diagram
![flowchart](flowchart_pinjaman.png)

# Structure Database
#### Customer
![DB Customer](customer_service.png)

#### Transaction
![DB Customer](transaction_service.png)

# Run
you just have to run the command below
`docker-compose up -d`
and wait for all services to run successfully
