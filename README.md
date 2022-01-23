# go-url-shortener
It's a simple url shortner. It provides an api and a cli interface.

# Techs
- [golang](https://go.dev/) - programming language 
  - [echo](https://echo.labstack.com/) - web framework
  - [go-redis](https://github.com/go-redis/redis) - redis client
  - [viper](https://github.com/spf13/viper) - config file
- [docker](https://docs.docker.com/engine/) - build, deploy and manager containers
  - [docker-compose](https://docs.docker.com/compose/) -  define and run multiple containers
- [redis](https://redis.io/) - key/value database
- [mongdb](https://www.mongodb.com) - NoSQL database
- [insomia](https://github.com/Kong/insomnia) - api client

# Deploy

You must have installed:
- Docker
- docker-compose

## env file
Use `.env-example` in case you need to change some defualt config


## Run
```
cd deployments
docker-compose up
```



