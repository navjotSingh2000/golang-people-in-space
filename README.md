# golang-people-in-space

## Docker Compose Instructions

```bash
Run the app:
docker compose up -d --build

(-d flag to run in detached mode)
(--build flag forces Docker Compose to build image from the Dockerfile before starting the container)
```

```bash
Attach to the terminal
docker attach golang-people-in-space-web-1
```

```bash
Execute commands in terminal
docker exec -it golang-people-in-space-web-1 sh
```
