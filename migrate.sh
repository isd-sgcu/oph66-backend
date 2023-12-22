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
				
				DROP TABLE IF EXISTS schedules;
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
					name_en VARCHAR(128) NOT NULL,
					name_th VARCHAR(128) NOT NULL,
					faculty_code SMALLINT NOT NULL REFERENCES faculties(code),
					department_en VARCHAR(128) NOT NULL,
					department_th VARCHAR(128) NOT NULL,
					require_registration BOOLEAN NOT NULL,
					max_capacity INTEGER,
					location_en VARCHAR(128) NOT NULL,
					location_th VARCHAR(128) NOT NULL,
					description_en VARCHAR(2048),
					description_th VARCHAR(2048)
				);
				INSERT INTO events (id, name_en, name_th, faculty_code, department_en, department_th, require_registration, max_capacity, location_en, location_th, description_en, description_th) VALUES (
					'first-event',
					'First Event',
					'อีเวนท์แรก',
					21,
					'Computer Engineering',
					'ภาควิชาคอมพิวเตอร์',
					TRUE,
					'250',
					'Engineering Building 3',
					'ตึก 3',
					'The first event',
					'รายละเอียดอีเวนท์แรก'
				);
				INSERT INTO events (id, name_en, name_th, faculty_code, department_en, department_th, require_registration, max_capacity, location_en, location_th, description_en, description_th) VALUES (
					'second-event',
					'Second Event',
					'อีเวนท์สอง',
					23,
					'Chemistry',
					'ภาควิชาเคมี',
					TRUE,
					'150',
					'Mahamakut Building',
					'ตึกมหามกุฎ',
					'The second event',
					'รายละเอียดอีเวนท์ที่สอง'
				);
				INSERT INTO events (id, name_en, name_th, faculty_code, department_en, department_th, require_registration, max_capacity, location_en, location_th, description_en, description_th) VALUES (
					'third-event',
					'Third Event',
					'อีเวนท์สาม',
					34,
					'Faculty',
					'ส่วนกลาง',
					FALSE,
					NULL,
					'Debdavaravati Building',
					'ตึกเทพทวารวดี',
					'The third event',
					'รายละเอียดอีเวนท์ที่สาม'
				);

				CREATE TABLE schedules (
					event_id VARCHAR(128) REFERENCES events(id),
					starts_at TIMESTAMP WITH TIME ZONE NOT NULL,
					ends_at TIMESTAMP WITH TIME ZONE NOT NULL,
					PRIMARY KEY (event_id, starts_at, ends_at)
				);
				INSERT INTO schedules (event_id, starts_at, ends_at) VALUES (
					'first-event',
					'2024-01-20 03:00:00+00',
					'2024-01-20 10:00:00+00'
				);
				INSERT INTO schedules (event_id, starts_at, ends_at) VALUES (
					'first-event',
					'2024-01-21 03:00:00+00',
					'2024-01-21 10:00:00+00'
				);
				INSERT INTO schedules (event_id, starts_at, ends_at) VALUES (
					'second-event',
					'2024-01-20 02:00:00+00',
					'2024-01-20 09:00:00+00'
				);
				INSERT INTO schedules (event_id, starts_at, ends_at) VALUES (
					'second-event',
					'2024-01-21 02:00:00+00',
					'2024-01-21 09:00:00+00'
				);
				INSERT INTO schedules (event_id, starts_at, ends_at) VALUES (
					'third-event',
					'2024-01-20 02:00:00+00',
					'2024-01-20 09:00:00+00'
				);
				INSERT INTO schedules (event_id, starts_at, ends_at) VALUES (
					'third-event',
					'2024-01-21 02:00:00+00',
					'2024-01-21 09:00:00+00'
				);
        
			\" | psql postgres://postgres:123456@host.docker.internal:5432/postgres
		"