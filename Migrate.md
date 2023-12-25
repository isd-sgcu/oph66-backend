## Migrating with Atlas

### Overview

Atlas provides a streamlined way to manage and migrate database schemas effortlessly. Follow the steps below to set up and migrate your database using Atlas.

### Step-by-Step Guide

#### 1. Installation

Firstly, install Atlas using the provided script:

```bash
curl -sSf https://atlasgo.sh | sh
```

This will ensure you have the necessary tools to inspect and apply schema changes seamlessly.

#### 2. Inspecting the Database and Generating Schema 

Whenever you make changes to your database schema, you must update your schema definition. To do this, inspect your current database setup and generate a `schema.sql` file:

```bash
atlas schema inspect \
  --url "postgres://postgres:123456@127.0.0.1:5432/postgres?search_path=public&sslmode=disable" \
  --format "{{ sql . }}" > schema.sql
```

This command fetches the current schema structure and outputs it to a `schema.sql` file, ensuring you have an up-to-date representation of your database schema.

#### 3. Applying Schema Changes

Once you've made the necessary updates to the `schema.sql` file, you can apply these changes to your database:

```bash
atlas schema apply \
  --url "postgres://postgres:123456@127.0.0.1:5432/postgres?&sslmode=disable" \
  --to "file://schema.sql" \
  --dev-url "docker://postgres/15"
```

Here's what each parameter does:

- `--url`: Specifies the connection URL to your target database where changes will be applied.
- `--to`: Indicates the path to the `schema.sql` file containing the schema changes.
- `--dev-url`: Provides a development URL for rolling back changes if necessary, ensuring a safe migration process.

#### 4. Confirm and Apply

After executing the migration command, review the changes to ensure everything aligns with your expectations. If satisfied, proceed with the migration to finalize the schema changes in your database.

---

This improved documentation offers a structured approach, providing clarity on each step and its purpose.