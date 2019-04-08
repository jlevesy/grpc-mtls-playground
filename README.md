# mTLS & gRPC playground

Sandbox demonstrating how to use gPRC with mTLS and generate certificates using go.

- Direct connection client to server
- Through [Traefik](https://traefik.io), which terminate TLS and forwards to the application over h2c.

## How to use ?

### Client & Server directly
```
make generate
make run_server

# In another terminal
make run_client
```

### Throught Traefik

```
make generate
docker-compose up -d

# In another terminal
make run_client
```
