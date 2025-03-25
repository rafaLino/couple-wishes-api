```bash
docker network create --driver bridge postgres-network
```

```bash
docker run --name postgresdb --network=postgres-network -p 5432:5432 -v ~/tmp/database:/var/lib/postgresql/data -e POSTGRES_PASSWORD=1234 -d postgres
```

```bash
docker run --name pgadmin-container --network=postgres-network -p 5050:80 -e PGADMIN_DEFAULT_EMAIL=user@domain.com -e PGADMIN_DEFAULT_PASSWORD=1234 -d dpage/pgadmin4
```
