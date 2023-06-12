BEGIN;
  CREATE TABLE songs (
    -- id SERIAL
    id uuid DEFAULT uuid_generate_v4 ()
    --
    ,name text
    ,url text
    ,tone VARCHAR ( 255 )
    --
    ,created_at timestamptz  NOT NULL  DEFAULT current_timestamp
    ,updated_at timestamptz  NOT NULL  DEFAULT current_timestamp
    -- soft delete
    ,deleted_at   timestamptz
    --
    ,PRIMARY KEY (id)
  );

  CREATE INDEX "idx_songs_deleted_at" ON "songs" ("deleted_at");
  CREATE INDEX "idx_songs_name" ON "songs" ("name");
COMMIT;
