-- Create "inscriptions" table
CREATE TABLE "inscriptions" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "inscription_id" bigint NOT NULL, "uid" character varying NOT NULL, "address" character varying NOT NULL, "output_value" bigint NOT NULL, "content_length" bigint NOT NULL, "content_type" character varying NOT NULL, "timestamp" timestamptz NOT NULL, "genesis_height" bigint NOT NULL, "genesis_fee" bigint NOT NULL, "genesis_tx" character varying NOT NULL, "location" character varying NOT NULL, "output" character varying NOT NULL, "offset" bigint NOT NULL, PRIMARY KEY ("id"));
-- Create index "inscriptions_inscription_id_key" to table: "inscriptions"
CREATE UNIQUE INDEX "inscriptions_inscription_id_key" ON "inscriptions" ("inscription_id");
-- Create index "inscriptions_uid_key" to table: "inscriptions"
CREATE UNIQUE INDEX "inscriptions_uid_key" ON "inscriptions" ("uid");