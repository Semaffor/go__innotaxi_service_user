CREATE TABLE "usr"
(
    "id"            SERIAL PRIMARY KEY,
    "username"      VARCHAR(100) NOT NULL,
    "phone_number"  VARCHAR(15)  NOT NULL,
    "email"         VARCHAR(32),
    "password_hash" VARCHAR(45)  NOT NULL,
    "role"          BIGINT       NOT NULL,
    "total_mark"    decimal(3, 2),
    "is_deleted"    boolean      NOT NULL DEFAULT FALSE
);

CREATE TABLE "usr_token"
(
    "id"      SERIAL PRIMARY KEY,
    "user_id" BIGINT      NOT NULL,
    "token"   VARCHAR(45) NOT NULL
);

CREATE TABLE "order"
(
    "id"           SERIAL PRIMARY KEY,
    "from"         VARCHAR(255) NOT NULL,
    "to"           VARCHAR(255) NOT NULL,
    "date"         timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "status"       VARCHAR(16)  NOT NULL,
    "driver_id"    BIGINT       NOT NULL,
    "taxi_type_id" BIGINT       NOT NULL,
    "user_id"      BIGINT       NOT NULL
);

CREATE TABLE "role"
(
    "id"          SERIAL PRIMARY KEY,
    "name"        VARCHAR(16)  NOT NULL,
    "description" VARCHAR(255) NOT NULL
);

CREATE TABLE "usr_role"
(
    "id"      SERIAL PRIMARY KEY,
    "usr_id"  BIGINT NOT NULL,
    "role_id" BIGINT NOT NULL
);

CREATE TABLE "driver"
(
    "id"           SERIAL PRIMARY KEY,
    "usr_id"       BIGINT NOT NULL,
    "status"       BIGINT NOT NULL,
    "taxi_type_id" BIGINT NOT NULL,
    "total_mark"   decimal(3, 2)
);

CREATE TABLE "taxi_type"
(
    "id"          SERIAL PRIMARY KEY,
    "type"        VARCHAR(16)  NOT NULL,
    "description" VARCHAR(255) NOT NULL
);

CREATE TABLE "feedback"
(
    "id"                   SERIAL PRIMARY KEY,
    "driver_id"            BIGINT NOT NULL,
    "customer_id"          BIGINT NOT NULL,
    "mark_from_user"       decimal(3, 2),
    "mark_from_driver"     decimal(3, 2),
    "feedback_from_user"   VARCHAR(255),
    "feedback_from_driver" VARCHAR(255),
    "order_id"             BIGINT NOT NULL
);

ALTER TABLE
    "usr_role"
    ADD CONSTRAINT "usr_role_usr_id_foreign" FOREIGN KEY ("usr_id") REFERENCES "usr" ("id");
ALTER TABLE
    "order"
    ADD CONSTRAINT "order_driver_id_foreign" FOREIGN KEY ("driver_id") REFERENCES "driver" ("id");
ALTER TABLE
    "order"
    ADD CONSTRAINT "order_taxi_type_id_foreign" FOREIGN KEY ("taxi_type_id") REFERENCES "taxi_type" ("id");
ALTER TABLE
    "order"
    ADD CONSTRAINT "order_usr_id_foreign" FOREIGN KEY ("user_id") REFERENCES "usr" ("id");
ALTER TABLE
    "usr_token"
    ADD CONSTRAINT "usr_token_user_id_foreign" FOREIGN KEY ("user_id") REFERENCES "usr" ("id");
ALTER TABLE
    "usr_role"
    ADD CONSTRAINT "usr_role_role_id_foreign" FOREIGN KEY ("role_id") REFERENCES "role" ("id");
ALTER TABLE
    "driver"
    ADD CONSTRAINT "driver_usr_id_foreign" FOREIGN KEY ("usr_id") REFERENCES "usr" ("id");
ALTER TABLE
    "driver"
    ADD CONSTRAINT "driver_taxi_type_id_foreign" FOREIGN KEY ("taxi_type_id") REFERENCES "taxi_type" ("id");
ALTER TABLE
    "feedback"
    ADD CONSTRAINT "feedback_driver_id_foreign" FOREIGN KEY ("driver_id") REFERENCES "driver" ("id");
ALTER TABLE
    "feedback"
    ADD CONSTRAINT "feedback_customer_id_foreign" FOREIGN KEY ("customer_id") REFERENCES "usr" ("id");
ALTER TABLE
    "feedback"
    ADD CONSTRAINT "feedback_order_id_foreign" FOREIGN KEY ("order_id") REFERENCES "order" ("id");