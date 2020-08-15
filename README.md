# Auth system

### Run
```
docker-compose up backend
```

Backend  http://localhost:8000

### Run in debug mode

```
docker-compose -f docker-compose.yml -f docker-compose-debug.yml up backend
```

For debugging payments
```
ngrok http 8000
```

### Run test
```bash
docker-compose run --rm backend make test
```

For running in JetBrains need set environment variable
```
DB_TEST_URL=postgres://admin:admin@localhost/auth_test?sslmode=disable
```


### Linting

```bash
docker-compose run --rm backend make hooks
```

### Migrations

Create new migration
```
docker-compose run --rm backend migrate create -ext sql -dir migrations -seq create_project_table
```

Migrate
```
docker-compose run --rm migrate -path=/migrations/ -database "postgres://admin:admin@postgres/auth?sslmode=disable" up
```

### Deploy
```
nixops create ./deploy/amazon-ec2.nix ./deploy/machine-config.nix -d onlinePayments
nixops deploy -d onlinePayments
```
Info 
```
nixops info -d onlinePayments
```

Connect to machine
```
nixops ssh -d onlinePayments online-payments
```
