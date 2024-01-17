CREATE TYPE feedback_first_part_answer AS ENUM (
    '1',
    '2',
    '3',
    '4',
    '5'
);

CREATE TYPE feedback_second_part_answer AS ENUM (
    '1',
    '2',
    '3',
    '4',
    '5',
    'did-not-attend'
);

CREATE TABLE feedbacks (
    "user_id" INT NOT NULL,
    "q1" feedback_first_part_answer NOT NULL,
    "q2" feedback_first_part_answer NOT NULL,
    "q3" feedback_first_part_answer NOT NULL,
    "q4" feedback_first_part_answer NOT NULL,
    "q5" feedback_first_part_answer NOT NULL,
    "q6" feedback_first_part_answer NOT NULL,
    "q7" feedback_second_part_answer NOT NULL,
    "q8" feedback_second_part_answer NOT NULL,
    "q9" feedback_second_part_answer NOT NULL,
    "q10" feedback_second_part_answer NOT NULL,
    "q11" feedback_second_part_answer NOT NULL,
    "q12" feedback_second_part_answer NOT NULL,
    "q13" feedback_second_part_answer NOT NULL,
    "q14" feedback_second_part_answer NOT NULL,
    "q15" feedback_second_part_answer NOT NULL,
    "q16" feedback_second_part_answer NOT NULL,
    "q17" feedback_second_part_answer NOT NULL,
    "q18" feedback_second_part_answer NOT NULL,
    "q19" feedback_second_part_answer NOT NULL,
    "comment" VARCHAR(600),
    PRIMARY KEY ("user_id"),
    CONSTRAINT "fk_feedbacks_users" FOREIGN KEY ("user_id") REFERENCES "users" ("id")
);
