#!/bin/sh

docker exec \
		-it \
		oph66-db \
		psql -U postgres -c $'
				INSERT INTO faculties (code, name_en, name_th) VALUES
				(21, \'Faculty of Engineering\', \'คณะวิศวกรรมศาสตร์\'),
				(23, \'Faculty of Science\', \'คณะวิทยาศาสตร์\'),
				(34, \'Faculty of Law\', \'คณะนิติศาสตร์\');

				INSERT INTO events (id, name_en, name_th, faculty_code, department_en, department_th, require_registration, max_capacity, location_en, location_th, description_en, description_th) VALUES
				(
					\'first-event\',
					\'First Event\',
					\'อีเวนท์แรก\',
					21,
					\'Computer Engineering\',
					\'ภาควิชาคอมพิวเตอร์\',
					TRUE,
					\'250\',
					\'Engineering Building 3\',
					\'ตึก 3\',
					\'The first event\',
					\'รายละเอียดอีเวนท์แรก\'
				),
				(
					\'second-event\',
					\'Second Event\',
					\'อีเวนท์สอง\',
					23,
					\'Chemistry\',
					\'ภาควิชาเคมี\',
					TRUE,
					\'150\',
					\'Mahamakut Building\',
					\'ตึกมหามกุฎ\',
					\'The second event\',
					\'รายละเอียดอีเวนท์ที่สอง\'
				),
				(
					\'third-event\',
					\'Third Event\',
					\'อีเวนท์สาม\',
					34,
					\'Faculty\',
					\'ส่วนกลาง\',
					FALSE,
					NULL,
					\'Debdavaravati Building\',
					\'ตึกเทพทวารวดี\',
					\'The third event\',
					\'รายละเอียดอีเวนท์ที่สาม\'
				);

				INSERT INTO schedules (event_id, starts_at, ends_at, period) VALUES (
					\'first-event\',
					\'2024-01-20 03:00:00+00\',
					\'2024-01-20 05:00:00+00\',
					\'20-morning\'
				),
				(
					\'first-event\',
					\'2024-01-21 06:00:00+00\',
					\'2024-01-21 08:00:00+00\',
					\'21-afternoon\'
				),
				(
					\'second-event\',
					\'2024-01-20 06:00:00+00\',
					\'2024-01-20 08:00:00+00\',
					\'20-afternoon\'
				),
				(
					\'second-event\',
					\'2024-01-21 03:00:00+00\',
					\'2024-01-21 05:00:00+00\',
					\'21-morning\'
				),
				(
					\'third-event\',
					\'2024-01-20 02:00:00+00\',
					\'2024-01-20 05:00:00+00\',
					\'20-morning\'
				),
				(
					\'third-event\',
					\'2024-01-20 06:00:00+00\',
					\'2024-01-20 09:00:00+00\',
					\'20-afternoon\'
				),
				(
					\'third-event\',
					\'2024-01-21 02:00:00+00\',
					\'2024-01-21 05:00:00+00\',
					\'21-morning\'
				),
				(
					\'third-event\',
					\'2024-01-21 06:00:00+00\',
					\'2024-01-21 09:00:00+00\',
					\'21-afternoon\'
				);

			INSERT INTO feature_flags(key, enabled, cache_duration, extra_info) VALUES (\'livestream\', TRUE, 10, \'{\"url\": \"https://www.youtube.com/watch?v=0tOXxuLcaog\"}\');
		'
