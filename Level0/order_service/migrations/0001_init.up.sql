CREATE TABLE "orders" (
                          "order_uid" varchar UNIQUE PRIMARY KEY NOT NULL,
                          "track_number" varchar NOT NULL,
                          "entry" varchar NOT NULL,
                          "locale" varchar NOT NULL,
                          "internal_signature" varchar NOT NULL,
                          "customer_id" varchar NOT NULL,
                          "delivery_service" varchar NOT NULL,
                          "shardkey" varchar NOT NULL,
                          "sm_id" int NOT NULL,
                          "date_created" timestamp with time zone NOT NULL,
                          "oof_shard" varchar NOT NULL
);

CREATE TABLE "deliveries" (
                              "order_uid" varchar UNIQUE PRIMARY KEY NOT NULL,
                              "name" varchar NOT NULL,
                              "phone" varchar NOT NULL,
                              "zip" varchar NOT NULL,
                              "city" varchar NOT NULL,
                              "address" varchar NOT NULL,
                              "region" varchar NOT NULL,
                              "email" varchar NOT NULL
);

CREATE TABLE "payments" (
                            "order_uid" varchar UNIQUE PRIMARY KEY NOT NULL,
                            "transaction" varchar NOT NULL,
                            "request_id" varchar NOT NULL,
                            "currency" varchar NOT NULL,
                            "provider" varchar NOT NULL,
                            "amount" int NOT NULL,
                            "payment_dt" bigint NOT NULL,
                            "bank" varchar NOT NULL,
                            "delivery_cost" int NOT NULL,
                            "goods_total" int NOT NULL,
                            "custom_fee" int NOT NULL
);

CREATE TABLE "items" (
                         "id" serial PRIMARY KEY,
                         "order_uid" varchar NOT NULL,
                         "chrt_id" bigint NOT NULL,
                         "track_number" varchar NOT NULL,
                         "price" int NOT NULL,
                         "rid" varchar NOT NULL,
                         "name" varchar NOT NULL,
                         "sale" int NOT NULL,
                         "size" varchar NOT NULL,
                         "total_price" int NOT NULL,
                         "nm_id" bigint NOT NULL,
                         "brand" varchar NOT NULL,
                         "status" int NOT NULL
);

CREATE INDEX ON "items" ("order_uid");

ALTER TABLE "deliveries" ADD FOREIGN KEY ("order_uid") REFERENCES "orders" ("order_uid") ON DELETE CASCADE;

ALTER TABLE "payments" ADD FOREIGN KEY ("order_uid") REFERENCES "orders" ("order_uid") ON DELETE CASCADE;

ALTER TABLE "items" ADD FOREIGN KEY ("order_uid") REFERENCES "orders" ("order_uid") ON DELETE CASCADE;

CREATE INDEX ON "items" ("order_uid");

CREATE INDEX orders_date_created_idx ON "orders" USING btree (date_created);