# oph66-backend
# Stack
- Gorm
- Postgres
- Redis
- Gin

# Contribution

## Development
`docker-compose.yaml` provide every component to start the backend.
```sh
docker compose up -d
```
Migration is need to still be made manually because migration method is 
currently to be determined. Feel free to give any suggestion :)
```sh
docker run \
    -it \
    --rm \
    --add-host host.docker.internal:host-gateway \
    postgres:15.5-bookworm \
    sh -c "
        echo \"
            CREATE TABLE feature_flags(key VARCHAR(50) PRIMARY KEY, value BOOLEAN NOT NULL, cache_duration INT NOT NULL);
            INSERT INTO feature_flags(key, value, cache_duration) VALUES ('livestream', TRUE, 10);
        \" | psql postgres://postgres:123456@host.docker.internal:5432/postgres
    "
```

## Using wire
This repository use [wire](https://github.com/google/wire) as dependency 
injection tools. To add provider/injector, take a look at `/di/wire.go`. Don't 
get confused with `/di/wire_gen.go`. It is generated file. To generate code
run.
```sh
go run github.com/google/wire/cmd/wire@latest ./...
```

## Conventional Commit Format
Before making a commit, please make sure that you have run formatter with
```sh
gofmt -w **/*.go
```
In short, the commit message should look like this:

```bash
git commit -m "feat: <what-you-did>"

# or

git commit -m "fix: <what-you-fixed>"

# or

git commit -m "refactor: <what-you-refactored>"
```

The commit message should start with one of the following types:

- feat: A new feature
- fix: A bug fix
- refactor: A code change that neither fixes a bug nor adds a feature
- style: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
- docs: Documentation only changes
- chore: Changes to the build process or auxiliary tools and libraries

For more information, please read the [conventional commit format](https://www.conventionalcommits.org/en/v1.0.0/) documentation.