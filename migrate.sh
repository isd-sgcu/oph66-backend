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
				
				DROP TABLE IF EXISTS events;
				DROP TABLE IF EXISTS faculties;

				CREATE TABLE faculties (
					code SMALLINT PRIMARY KEY,
					name_en VARCHAR(128) NOT NULL,
					name_th VARCHAR(128) NOT NULL
				);
				INSERT INTO faculties (code, name_en, name_th) VALUES (21, 'Faculty of Engineering', 'คณะวิศวกรรมศาสตร์');
				INSERT INTO faculties (code, name_en, name_th) VALUES (23, 'Faculty of Science', 'คณะวิทยาศาสตร์');
				INSERT INTO faculties (code, name_en, name_th) VALUES (34, 'Faculty of Law', 'คณะนิติศาสตร์');

        CREATE TABLE events (
					id VARCHAR(128) PRIMARY KEY,
					name VARCHAR NOT NULL,
					faculty_code SMALLINT NOT NULL REFERENCES faculties(code),
					department VARCHAR(128) NOT NULL,
					require_registration BOOLEAN NOT NULL,
					max_capacity INTEGER,
					start_time TIMESTAMP WITH TIME ZONE NOT NULL,
					location_en VARCHAR(128) NOT NULL,
					location_th VARCHAR(128) NOT NULL,
					description VARCHAR(128)
				);
				INSERT INTO events (id, name, faculty_code, department, require_registration, max_capacity, start_time, location_en, location_th, description) VALUES (
					'first-event',
					'First Event',
					21,
					'ภาควิชาคอมพิวเตอร์',
					TRUE,
					'250',
					'2023-12-19 10:00:00+00',
					'Engineering Building 3',
					'ตึก 3',
					'The first event'
				);
				INSERT INTO events (id, name, faculty_code, department, require_registration, max_capacity, start_time, location_en, location_th, description) VALUES (
					'second-event',
					'Second Event',
					23,
					'ภาควิชาเคมี',
					TRUE,
					'250',
					'2023-12-19 10:00:00+00',
					'Mahamakut Building',
					'ตึกมหามกุฎ',
					'The second event'
				);
				INSERT INTO events (id, name, faculty_code, department, require_registration, max_capacity, start_time, location_en, location_th, description) VALUES (
					'third-event',
					'Third Event',
					34,
					'ส่วนกลาง',
					FALSE,
					NULL,
					'2023-12-19 10:00:00+00',
					'Deebdavaravati Building',
					'ตึกเทพทวารวดี',
					'The third event'
				);
        
			\" | psql postgres://postgres:123456@host.docker.internal:5432/postgres
		"