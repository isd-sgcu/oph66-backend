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
-- Create "rounds" table
CREATE TABLE "rounds" (
    "round_no" INT NOT NULL,
    PRIMARY KEY ("round_no")
);
-- Create "users" table
CREATE TABLE "users" (
    "id" SERIAL NOT NULL,
    "gender" VARCHAR(80) NULL,
    "first_name" VARCHAR(80) NULL,
    "last_name" VARCHAR(80) NULL,
    "email" VARCHAR(80) NOT NULL,
    "school" VARCHAR(80) NULL,
    "birth_date" VARCHAR(80) NULL,
    "address" VARCHAR(300) NULL,
    "from_abroad" VARCHAR(80) NULL,
    "allergy" VARCHAR(150) NULL,
    "medical_condition" VARCHAR(150) NULL,
    "join_cu_reason" VARCHAR(300) NULL,
    "news_source" VARCHAR(100) NULL,
    "status" VARCHAR(80) NULL,
    "grade" VARCHAR(50) NULL,
    PRIMARY KEY ("id")
);
-- Create index "idx_users_email" to table: "users"
CREATE INDEX "idx_users_email" ON "users" ("email");
-- Create "desired_rounds" table
CREATE TABLE "desired_rounds" (
    "user_id" INT NOT NULL,
    "order" INT NOT NULL,
    "round_code" INT NOT NULL,
    PRIMARY KEY ("user_id", "order"),
    CONSTRAINT "fk_desired_rounds_round" FOREIGN KEY ("round_code") REFERENCES "rounds" ("round_no") ON UPDATE NO ACTION ON DELETE NO ACTION,
    CONSTRAINT "fk_users_desired_rounds" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_desired_rounds_user_id" to table: "desired_rounds"
CREATE INDEX "idx_desired_rounds_user_id" ON "desired_rounds" ("user_id");
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
    "order" INT NOT NULL,
    "faculty_code" VARCHAR(10) NOT NULL,
    "department_code" VARCHAR(10) NOT NULL,
    PRIMARY KEY ("user_id", "order"),
    CONSTRAINT "fk_interested_faculties_department" FOREIGN KEY ("faculty_code", "department_code") REFERENCES "departments" ("faculty_code", "code") ON UPDATE NO ACTION ON DELETE NO ACTION,
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
    PRIMARY KEY ("id"),
    CONSTRAINT "fk_events_schedules" FOREIGN KEY ("event_id") REFERENCES "events" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
CREATE INDEX "idx_event_id" ON "schedules" ("event_id");

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
