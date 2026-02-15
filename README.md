Getting Started

# Clone the repository
```
git clone https://github.com/your-username/mini-s3-service.git
cd mini-s3-service
```

# Create environment file

Create a .env file based on the provided example:
```
cp .env.example .env
```
You can adjust the values if needed.

# Run the project
```
docker-compose up --build
```

The application will be available at: http://localhost:8080

Healthcheck endpoint: http://localhost:8080/health

MinIO console: http://localhost:9001