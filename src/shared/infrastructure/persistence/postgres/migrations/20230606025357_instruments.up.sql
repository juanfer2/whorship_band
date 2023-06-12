BEGIN;
  CREATE TABLE instruments (
    -- id SERIAL
    id uuid DEFAULT uuid_generate_v4 ()
    --
    ,name VARCHAR ( 255 ) UNIQUE
    --
    ,created_at timestamptz  NOT NULL  DEFAULT current_timestamp
    ,updated_at timestamptz  NOT NULL  DEFAULT current_timestamp
    -- soft delete
    ,deleted_at   timestamptz
    --
    ,PRIMARY KEY (id)
  );

  CREATE INDEX "idx_instruments_deleted_at" ON "instruments" ("deleted_at");
  CREATE INDEX "idx_instruments_name" ON "instruments" ("name");
COMMIT;
