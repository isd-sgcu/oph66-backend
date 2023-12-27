CREATE TABLE event_registrations (
    "user_id" INT NOT NULL,
    "schedule_id" INT NOT NULL,
    PRIMARY KEY ("user_id", "schedule_id")
);

ALTER TABLE "schedules" ADD COLUMN "current_attendee" INT DEFAULT 0;
