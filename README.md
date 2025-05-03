```bash
docker network create --driver bridge postgres-network
```

```bash
docker run --name postgresdb --network=postgres-network -p 5432:5432 -v ~/tmp/database:/var/lib/postgresql/data -e POSTGRES_PASSWORD=1234 -d postgres:alpine
```

```bash
docker run --name pgadmin-container --network=postgres-network -p 5050:80 -e PGADMIN_DEFAULT_EMAIL=user@domain.com -e PGADMIN_DEFAULT_PASSWORD=1234 -d dpage/pgadmin4
```

```bash
sudo snap install sqlc
```


```bash
wget http://github.com/golang-migrate/migrate/releases/latest/download/migrate.linux-amd64.deb
sudo dpkg -i migrate.linux-amd64.deb
```

```bash
docker run --net=host -it -e NGROK_AUTHTOKEN=<TOKEN> ngrok/ngrok:latest http 9000
```