#!/bin/sh

docker run \
		-it \
		--rm \
		--add-host host.docker.internal:host-gateway \
		postgres:15.5-bookworm \
		sh -c "
			echo \"
				CREATE TABLE feature_flags(key VARCHAR(50) PRIMARY KEY, value BOOLEAN NOT NULL, cache_duration INT NOT NULL);
				INSERT INTO feature_flags(key, value, cache_duration) VALUES ('livestream', TRUE, 10);
				
        CREATE TABLE events (
					id VARCHAR(128) PRIMARY KEY,
					name VARCHAR NOT NULL,
					description VARCHAR(128),
					start_time TIMESTAMP WITH TIME ZONE NOT NULL,
					location VARCHAR(128) NOT NULL,
					max_capacity INTEGER NOT NULL,
					department VARCHAR(128) NOT NULL
				);
			\" | psql postgres://postgres:123456@host.docker.internal:5432/postgres
		"