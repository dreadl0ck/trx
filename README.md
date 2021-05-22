# TRX

A simple example [Maltego](https://maltego.com) transform server using the [maltego](github.com/dreadl0ck/maltego) package.

## Docker containers

- **docker/alpine**: golang basend transform server on alpine linux
- **docker/ubuntu**: python and gunicorn transform server on ubuntu linux

You pull the prebuilt containers from my docker hub, eg:

    docker pull dreadl0ck/trx-alpine

or 
    
    docker pull dreadl0ck/trx-ubtuntu

## Note

- Automatic TLS certs via letsencrypt if desired. 
- No built-in basic auth at the moment, you need to run it either internally reachable or behind a reverse proxy with authentication.