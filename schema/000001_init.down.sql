ALTER TABLE
    "usr_role"
    DROP CONSTRAINT IF EXISTS "usr_role_usr_id_foreign";
ALTER TABLE
    "order"
    DROP CONSTRAINT IF EXISTS "order_driver_id_foreign";
ALTER TABLE
    "order"
    DROP CONSTRAINT IF EXISTS "order_taxi_type_id_foreign";
ALTER TABLE
    "order"
    DROP CONSTRAINT IF EXISTS "order_usr_id_foreign";
ALTER TABLE
    "usr_token"
    DROP CONSTRAINT IF EXISTS "usr_token_user_id_foreign";
ALTER TABLE
    "usr_role"
    DROP CONSTRAINT IF EXISTS "usr_role_role_id_foreign";
ALTER TABLE
    "driver"
    DROP CONSTRAINT IF EXISTS "driver_usr_id_foreign";
ALTER TABLE
    "driver"
    DROP CONSTRAINT IF EXISTS "driver_taxi_type_id_foreign";
ALTER TABLE
    "feedback"
    DROP CONSTRAINT IF EXISTS "feedback_driver_id_foreign";
ALTER TABLE
    "feedback"
    DROP CONSTRAINT IF EXISTS "feedback_customer_id_foreign";
ALTER TABLE
    "feedback"
    DROP CONSTRAINT IF EXISTS "feedback_order_id_foreign";

DROP TABLE IF EXISTS "usr";
DROP TABLE IF EXISTS "usr_token";
DROP TABLE IF EXISTS "order";
DROP TABLE IF EXISTS "role";
DROP TABLE IF EXISTS "usr_role";
DROP TABLE IF EXISTS "driver";
DROP TABLE IF EXISTS "taxi_type";
DROP TABLE IF EXISTS "feedback";
