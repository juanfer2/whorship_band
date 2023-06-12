BEGIN;
  CREATE TABLE group_mates (
    -- id SERIAL
    id uuid DEFAULT uuid_generate_v4 ()
    --
    ,name VARCHAR ( 255 ) UNIQUE
    ,url_image  text
    --
    ,created_at timestamptz  NOT NULL  DEFAULT current_timestamp
    ,updated_at timestamptz  NOT NULL  DEFAULT current_timestamp
    -- soft delete
    ,deleted_at   timestamptz
    --
    ,PRIMARY KEY (id)
  );

  CREATE INDEX "idx_group_mates_deleted_at" ON "group_mates" ("deleted_at");
  CREATE INDEX "idx_group_mates_name" ON "group_mates" ("name");
COMMIT;
