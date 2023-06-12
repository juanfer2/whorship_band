BEGIN;
  CREATE TABLE performances (
    -- id SERIAL
    id uuid DEFAULT uuid_generate_v4 ()
    --
    ,date timestamptz  NOT NULL
    ,notes text
    --
    ,created_at timestamptz  NOT NULL  DEFAULT current_timestamp
    ,updated_at timestamptz  NOT NULL  DEFAULT current_timestamp
    -- soft delete
    ,deleted_at   timestamptz
    --
    ,PRIMARY KEY (id)
  );

  CREATE INDEX "idx_performances_deleted_at" ON "performances" ("deleted_at");
  CREATE INDEX "idx_performances_date" ON "performances" ("date");
COMMIT;
