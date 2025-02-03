# Potato Proxy

An extremely simple HTTP reverse proxy in Go. Useful for proxying containers which you do not want to expose to the external network.

## How to use

The following two environment variables are required:

| Variable  | Description           | Example      |
| --------- | --------------------- | ------------ |
| PORT      | The port to listen on | 3000         |
| PROXY_URL | The URL to proxy      | service:4000 |

Pull the Docker image:

```bash
docker pull ghcr.io/ke126/potato-proxy:latest
```

See [compose.yaml](./compose.yaml) for an example of how to use Potato Proxy.
