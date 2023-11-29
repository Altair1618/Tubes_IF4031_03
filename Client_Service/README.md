## Development

Make sure to create .env file based on .env.example before running the server.

Install dependencies:
```bash
bunx install
```

Make sure to create network used for all containers:
```bash
docker network create tessera_network
```
To start the development server run:  
```bash
docker compose up -d
```

Make sure to push the schema to db:  
```bash
bunx drizzle-kit push:pg
```

To see db with drizzle studio:
```bash
bunx drizzle-kit studio
```

Open http://localhost:3000/ with your browser to see the result.