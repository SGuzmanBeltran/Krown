# Useful commands for Krown

## Docker
- Run docker container for redpanda with the command

    `docker run -d --name=redpanda -p 9092:9092 --memory=2g --memory-swap=2g vectorized/redpanda start --overprovisioned --smp 1 --memory 2G --reserve-memory 0M --node-id 0`

- Execute bash terminal in docker so you have access to it

    `docker exec -it redpanda /bin/bash`

- Review resources used by the redpanda container

    `docker stats redpanda`

- Delete all containers

    `docker rm -vf $(docker ps -aq)`

- Delete all images

    `docker rmi -f $(docker images -aq)`

- Delete volumes (need to be unused)

    `docker volume prune`

- Delete all

    `docker system prune -a --volumes`

- Delete all unused things

    `docker system prune`

## Redpanda

- Use redpanda's CLI command to create a new topic

    `rpk topic create test_topic`

- You can list topics like

    `rpk topic list`

## Golang

- Generate SQLC code (Uses configuration file `sqlc.yaml`)

    `sqlc generate`

- Generate migrations (`db/migration`) uses `Makefile`

    `make db-up`