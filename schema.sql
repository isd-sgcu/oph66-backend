-- Create "feature_flags" table
CREATE TABLE "feature_flags" (
    "key" character varying(50) NOT NULL,
    "value" boolean NOT NULL,
    "cache_duration" integer NOT NULL,
    PRIMARY KEY ("key")
);
-- Create "users" table
CREATE TABLE "users" (
    "id" bigserial NOT NULL,
    "gender" text NULL,
    "first_name" text NULL,
    "last_name" text NULL,
    "email" text NULL,
    "school" text NULL,
    "birth_date" text NULL,
    "address" text NULL,
    "from_abroad" text NULL,
    "allergy" text NULL,
    "medical_condition" text NULL,
    "join_cu_reason" text NULL,
    "news_source" text NULL,
    "status" text NULL,
    "grade" text NULL,
    PRIMARY KEY ("id")
);
-- Create "desired_round" table
CREATE TABLE "desired_round" (
    "id" serial NOT NULL,
    "order" integer NOT NULL,
    "round" character varying(128) NOT NULL,
    "user_id" bigint NULL,
    PRIMARY KEY ("id"),
    CONSTRAINT "desired_round_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "faculties" table
CREATE TABLE "faculties" (
    "code" smallint NOT NULL,
    "name_en" character varying(128) NOT NULL,
    "name_th" character varying(128) NOT NULL,
    PRIMARY KEY ("code")
);
-- Create "events" table
CREATE TABLE "events" (
    "id" character varying(128) NOT NULL,
    "name_en" character varying(128) NOT NULL,
    "name_th" character varying(128) NOT NULL,
    "faculty_code" smallint NOT NULL,
    "department_en" character varying(128) NOT NULL,
    "department_th" character varying(128) NOT NULL,
    "require_registration" boolean NOT NULL,
    "max_capacity" integer NULL,
    "location_en" character varying(128) NOT NULL,
    "location_th" character varying(128) NOT NULL,
    "description_en" character varying(2048) NULL,
    "description_th" character varying(2048) NULL,
    PRIMARY KEY ("id"),
    CONSTRAINT "events_faculty_code_fkey" FOREIGN KEY ("faculty_code") REFERENCES "faculties" ("code") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "interested_faculty" table
CREATE TABLE "interested_faculty" (
    "id" bigserial NOT NULL,
    "order" integer NOT NULL,
    "faculty" text NULL,
    "department" text NULL,
    "section" text NULL,
    "user_id" bigint NULL,
    PRIMARY KEY ("id"),
    CONSTRAINT "fk_users_interested_faculty" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_interested_faculty_user_id" to table: "interested_faculty"
CREATE INDEX "idx_interested_faculty_user_id" ON "interested_faculty" ("user_id");
-- Create "schedules" table
CREATE TABLE "schedules" (
    "event_id" character varying(128) NOT NULL,
    "starts_at" timestamptz NOT NULL,
    "ends_at" timestamptz NOT NULL,
    PRIMARY KEY ("event_id", "starts_at", "ends_at"),
    CONSTRAINT "schedules_event_id_fkey" FOREIGN KEY ("event_id") REFERENCES "events" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);