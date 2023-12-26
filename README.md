# oph66-backend
# Stack
- Gorm
- Postgres
- Redis
- Gin
- Atlas

# Contribution

## Development
### Getting started
`docker-compose.yaml` provide every component to start the backend.
```sh
docker compose up -d
```

### Migrating with Atlas
Atlas provides a streamlined way to manage and migrate database schemas effortlessly. Follow the steps below to set up and migrate your database using Atlas.
1. Installation

Firstly, install Atlas using the provided script:

```bash
curl -sSf https://atlasgo.sh | sh
```

This will ensure you have the necessary tools to inspect and apply schema changes seamlessly.

2. Inspecting the Database and Generating Schema 

Whenever you make changes to your database schema, you must update your schema definition. To do this, inspect your current database setup and generate a `init.sql` file:

```bash
atlas schema inspect \
  --url "postgres://postgres:123456@127.0.0.1:5432/postgres?search_path=public&sslmode=disable" \
  --format "{{ sql . }}" > ./migrations/init.sql
```

This command fetches the current schema structure and outputs it to a `init.sql` file, ensuring you have an up-to-date representation of your database schema.

3. Applying Schema Changes

Once you've made the necessary updates to the `init.sql` file, you can apply these changes to your database:

```bash
atlas schema apply \
  --url "postgres://postgres:123456@127.0.0.1:5432/postgres?&sslmode=disable" \
  --to "file://./migrations/init.sql" \
  --dev-url "docker://postgres/15"
```

Here's what each parameter does:

- `--url`: Specifies the connection URL to your target database where changes will be applied.
- `--to`: Indicates the path to the `schema.sql` file containing the schema changes.
- `--dev-url`: Provides a development URL for rolling back changes if necessary, ensuring a safe migration process.

4. Confirm and Apply

After executing the migration command, review the changes to ensure everything aligns with your expectations. If satisfied, proceed with the migration to finalize the schema changes in your database.

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