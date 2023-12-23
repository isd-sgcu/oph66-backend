# oph66-backend
# Stack
- Gorm
- Postgres
- Redis
- Gin

# Contribution

## Development
### Getting started
`docker-compose.yaml` provide every component to start the backend.
```sh
docker compose up -d
```
Migration is need to still be made manually because migration method is 
currently to be determined. Feel free to give any suggestion :)
```sh
docker exec \
    -it \
    oph66-db \
    psql -U postgres -c $'
        CREATE TABLE feature_flags(key VARCHAR(50) PRIMARY KEY, enabled BOOLEAN NOT NULL, cache_duration INT NOT NULL, extra_info JSONB NOT NULL);
        INSERT INTO feature_flags(key, enabled, cache_duration, extra_info) VALUES (\'livestream\', TRUE, 10, \'{\"url\": \"https://www.youtube.com/watch?v=0tOXxuLcaog\"}\');
    '    
```

### Using wire
This repository use [wire](https://github.com/google/wire) as dependency 
injection tools. To add provider/injector, take a look at `/di/wire.go`. Don't 
get confused with `/di/wire_gen.go`. It is generated file. To generate code
run.
```sh
go run github.com/google/wire/cmd/wire@latest ./...
```

## Workflow
1. Pull the latest code from `beta` branch.
```sh
git pull origin beta
```
2. Create new branch. Branch name from linear is preferable.
```sh
git checkout -b <branch-name>
```
3. Writing your wonderful error-prone code.
4. Since we do not have test setup, please thoroughly test your code.
5. Push your code to your branch. Also checkout [Conventional Commit Format](#conventional-commit-format).
```sh
git commit -am <commit-message>
git push origin <branch-name>
```
6. Create/Update pull request and tag any maintainer.
7. Wait for maintainer to give you review.
8. Update your code as maintainer suggest by go back to step 3.
9. Once the pull request is approved, then do rebase, squash and merge into `beta` branch.

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