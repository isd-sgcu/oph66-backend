-- Create "feature_flags" table
CREATE TABLE "feature_flags" (
    "key" VARCHAR(50) NOT NULL,
    "enabled" BOOLEAN NOT NULL,
    "cache_duration" INT NOT NULL,
    "extra_info" JSONB NOT NULL,
    PRIMARY KEY ("key")
);
-- Create "faculties" table
CREATE TABLE "faculties" (
    "code" VARCHAR(10) NOT NULL,
    "name_en" VARCHAR(80) NULL,
    "name_th" VARCHAR(80) NULL,
    PRIMARY KEY ("code")
);
-- Create "departments" table
CREATE TABLE "departments" (
    "code" VARCHAR(10) NOT NULL,
    "name_th" VARCHAR(80) NULL,
    "name_en" VARCHAR(80) NULL,
    "faculty_code" VARCHAR(10) NOT NULL,
    PRIMARY KEY ("code", "faculty_code"),
    CONSTRAINT "fk_departments_faculty" FOREIGN KEY ("faculty_code") REFERENCES "faculties" ("code") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "sections" table
CREATE TABLE "sections" (
    "code" VARCHAR(10) NOT NULL,
    "faculty_code" VARCHAR(10) NOT NULL,
    "department_code" VARCHAR(10) NOT NULL,
    "name_th" VARCHAR(80) NULL,
    "name_en" VARCHAR(80) NULL,
    PRIMARY KEY ("code", "department_code", "faculty_code"),
    CONSTRAINT "fk_sections_department" FOREIGN KEY ("faculty_code", "department_code") REFERENCES "departments" ("faculty_code", "code") ON UPDATE NO ACTION ON DELETE NO ACTION,
    CONSTRAINT "fk_sections_faculty" FOREIGN KEY ("faculty_code") REFERENCES "faculties" ("code") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "users" table
CREATE TABLE "users" (
    "id" SERIAL NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL,
    "updated_at" TIMESTAMPTZ NOT NULL,
    "first_name" VARCHAR(80) NULL,
    "last_name" VARCHAR(80) NULL,
    "email" VARCHAR(80) NOT NULL,
    "birth_date" VARCHAR(80) NULL,
    "join_cu_reason" VARCHAR(300) NULL,
    "status" VARCHAR(80) NULL,
    "country" VARCHAR(80) NULL,
    "province" VARCHAR(80) NULL,
    "educational_level" VARCHAR(80) NULL,
    "desired_round" VARCHAR(40) NULL,
    PRIMARY KEY ("id")
);

ALTER SEQUENCE "users_id_seq" RESTART 10000;

-- Create index "idx_users_email" to table: "users"
CREATE UNIQUE INDEX "idx_users_email" ON "users" ("email");
-- Create "rounds" enum
CREATE TYPE "rounds" AS ENUM (
	'1',
	'2',
	'3',
	'4',
	'5'
);

CREATE TYPE news_source AS ENUM (
    'facebook',
    'instagram',
    'faculty',
    'chula-student',
    'friend',
    'parent',
    'school',
    'other'
);

-- Create "news_sources" table
CREATE TABLE "news_sources_users" (
    "user_id" INT NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL,
    "updated_at" TIMESTAMPTZ NOT NULL,
    "news_source" news_source NOT NULL,
    PRIMARY KEY ("user_id", "news_sources"),
    CONSTRAINT "fk_news_sources_users" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);

-- Create "visiting_faculties" table
CREATE TABLE "visiting_faculties" (
    "user_id" bigint NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL,
    "updated_at" TIMESTAMPTZ NOT NULL,
    "order" INT NOT NULL,
    "faculty_code" VARCHAR(10) NOT NULL,
    "department_code" VARCHAR(10) NOT NULL,
    "section_code" VARCHAR(10) NOT NULL,
    PRIMARY KEY ("user_id", "order"),
    CONSTRAINT "fk_visiting_faculties_section" FOREIGN KEY ("section_code", "department_code", "faculty_code") REFERENCES "sections" ("code", "department_code", "faculty_code") ON UPDATE NO ACTION ON DELETE NO ACTION,
    CONSTRAINT "fk_visiting_faculties_department" FOREIGN KEY ("department_code", "faculty_code") REFERENCES "departments" ("code", "faculty_code") ON UPDATE NO ACTION ON DELETE NO ACTION,
    CONSTRAINT "fk_visiting_faculties_faculty" FOREIGN KEY ("faculty_code") REFERENCES "faculties" ("code") ON UPDATE NO ACTION ON DELETE NO ACTION,
    CONSTRAINT "fk_users_visiting_faculties" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_visiting_faculties_user_id" to table: "visiting_faculties"
CREATE INDEX "idx_visiting_faculties_user_id" ON "visiting_faculties" ("user_id");

-- Create "events" table
CREATE TABLE "events" (
    "id" VARCHAR(30) NOT NULL,
    "name_en" VARCHAR(80) NULL,
    "name_th" VARCHAR(80) NULL,
    "faculty_code" VARCHAR(10) NULL,
    "department_code" VARCHAR(10) NULL,
    "require_registration" BOOLEAN NULL,
    "max_capacity" INT NULL,
    "location_en" VARCHAR(300) NULL,
    "location_th" VARCHAR(300) NULL,
    "description_en" TEXT NULL,
    "description_th" TEXT NULL,
    PRIMARY KEY ("id"), 
    CONSTRAINT "fk_events_faculty" FOREIGN KEY ("faculty_code") REFERENCES "faculties" ("code") ON UPDATE NO ACTION ON DELETE NO ACTION,
    CONSTRAINT "fk_events_department" FOREIGN KEY ("faculty_code", "department_code") REFERENCES "departments" ("faculty_code", "code") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "interested_faculties" table
CREATE TABLE "interested_faculties" (
    "user_id" bigint NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL,
    "updated_at" TIMESTAMPTZ NOT NULL,
    "order" INT NOT NULL,
    "faculty_code" VARCHAR(10) NOT NULL,
    "department_code" VARCHAR(10) NOT NULL,
    "section_code" VARCHAR(10) NOT NULL,
    PRIMARY KEY ("user_id", "order"),
    CONSTRAINT "fk_interested_faculties_section" FOREIGN KEY ("section_code", "department_code", "faculty_code") REFERENCES "sections" ("code", "department_code", "faculty_code") ON UPDATE NO ACTION ON DELETE NO ACTION,
    CONSTRAINT "fk_interested_faculties_department" FOREIGN KEY ("department_code", "faculty_code") REFERENCES "departments" ("code", "faculty_code") ON UPDATE NO ACTION ON DELETE NO ACTION,
    CONSTRAINT "fk_interested_faculties_faculty" FOREIGN KEY ("faculty_code") REFERENCES "faculties" ("code") ON UPDATE NO ACTION ON DELETE NO ACTION,
    CONSTRAINT "fk_users_interested_faculties" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_interested_faculties_user_id" to table: "interested_faculties"
CREATE INDEX "idx_interested_faculties_user_id" ON "interested_faculties" ("user_id");
-- Create "schedules" table
CREATE TYPE "schedule_period" AS ENUM (
    '20-morning',
    '20-afternoon',
    '21-morning',
    '21-afternoon'
);
CREATE TABLE "schedules" (
    "id" SERIAL NOT NULL,
    "event_id" VARCHAR(30) NOT NULL,
    "starts_at" timestamptz NULL,
    "ends_at" timestamptz NULL,
    "period" schedule_period NULL,
    "current_attendee" INT NOT NULL DEFAULT 0,
    PRIMARY KEY ("id"),
    CONSTRAINT "fk_events_schedules" FOREIGN KEY ("event_id") REFERENCES "events" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
CREATE INDEX "idx_event_id" ON "schedules" ("event_id");

-- Create event regestration table
CREATE TABLE event_registrations (
    "user_id" INT NOT NULL,
    "schedule_id" INT NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL,
    "updated_at" TIMESTAMPTZ NOT NULL,
    PRIMARY KEY ("user_id", "schedule_id")
);

CREATE INDEX "idx_event_registrations_user_id" ON "event_registrations" ("user_id"); 

INSERT INTO feature_flags(key, enabled, cache_duration, extra_info) VALUES ('livestream', FALSE, 10, '{"url": "https://www.youtube.com/watch?v=0tOXxuLcaog"}');

CREATE OR REPLACE FUNCTION update_attandee_count()
 RETURNS void
 LANGUAGE plpgsql
AS $function$
declare
	sid INTEGER;
	cnt INTEGER;
BEGIN
for sid, cnt in select id, count(user_id) from schedules s left join event_registrations er on s.id = er.schedule_id group by s.id loop
	update schedules set current_attendee = cnt where id = sid;
end loop;
end;
$function$;
