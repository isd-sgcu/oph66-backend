#!/bin/sh

docker run \
		-it \
		--rm \
		--add-host host.docker.internal:host-gateway \
		postgres:15.5-bookworm \
		sh -c "
			echo \"
				DROP TABLE IF EXISTS feature_flags;
				CREATE TABLE feature_flags(key VARCHAR(50) PRIMARY KEY, value BOOLEAN NOT NULL, cache_duration INT NOT NULL);
				INSERT INTO feature_flags(key, value, cache_duration) VALUES ('livestream', TRUE, 10);
				
				DROP TABLE IF EXISTS faculties;
				CREATE TABLE faculties (
					id VARCHAR(8) PRIMARY KEY,
					name_en VARCHAR(128) NOT NULL,
					name_th VARCHAR(128) NOT NULL
				);
				INSERT INTO faculties (id, name_en, name_th) VALUES ('ENG', 'Faculty of Engineering', 'คณะวิศวกรรมศาสตร์');
				INSERT INTO faculties (id, name_en, name_th) VALUES ('SCI', 'Faculty of Science', 'คณะวิทยาศาสตร์');
				INSERT INTO faculties (id, name_en, name_th) VALUES ('LAW', 'Faculty of Law', 'คณะนิติศาสตร์');

				DROP TABLE IF EXISTS events;
        CREATE TABLE events (
					id VARCHAR(128) PRIMARY KEY,
					name VARCHAR NOT NULL,
					faculty VARCHAR(8) REFERENCES faculties(id),
					department VARCHAR(128) NOT NULL,
					require_registration BOOLEAN NOT NULL,
					max_capacity INTEGER,
					start_time TIMESTAMP WITH TIME ZONE NOT NULL,
					location VARCHAR(128) NOT NULL,
					description VARCHAR(128)
				);
				INSERT INTO events (id, name, faculty, department, require_registration, max_capacity, start_time, location, description) VALUES (
					'first-event',
					'First Event',
					'ENG',
					'ภาควิชาคอมพิวเตอร์',
					TRUE,
					'250',
					'2023-12-19 10:00:00+00',
					'ตึก 3',
					'The first event'
				);
				INSERT INTO events (id, name, faculty, department, require_registration, max_capacity, start_time, location, description) VALUES (
					'second-event',
					'Second Event',
					'SCI',
					'ภาควิชาเคมี',
					TRUE,
					'250',
					'2023-12-19 10:00:00+00',
					'ตึกมหามกุฎ',
					'The second event'
				);
				INSERT INTO events (id, name, faculty, department, require_registration, max_capacity, start_time, location, description) VALUES (
					'third-event',
					'Third Event',
					'LAW',
					'ส่วนกลาง',
					FALSE,
					NULL,
					'2023-12-19 10:00:00+00',
					'ตึกเทพทวารวดี',
					'The third event'
				);
        
			\" | psql postgres://postgres:123456@host.docker.internal:5432/postgres
		"