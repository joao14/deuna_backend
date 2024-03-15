Indications:

1. To init proyect execute

docker-compose -f docker-compose.yml up

2. Create file .env and .env.local and put the follow environments en booth.

# Database credentials
DB_HOST="localhost"
DB_USER="postgres"
DB_PASSWORD="postgres"
DB_NAME="pasarela"
DB_PORT="5432"

# Authentication credentials
TOKEN_TTL="2000"
JWT_PRIVATE_KEY="deuna2024"