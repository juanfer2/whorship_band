BEGIN;
  CREATE TABLE products (
    id SERIAL
    --
    ,type int
    ,token  text
    ,source_id int
    ,bin text
    ,name VARCHAR ( 255 )
    ,mask VARCHAR ( 255 )
    --
    ,created_at timestamptz  NOT NULL  DEFAULT current_timestamp
    ,updated_at timestamptz  NOT NULL  DEFAULT current_timestamp
    -- soft delete
    ,deleted_at   timestamptz
    --
    ,PRIMARY KEY (id)
  );

  CREATE INDEX "idx_products_deleted_at" ON "products" ("deleted_at");
  CREATE INDEX "idx_products_token" ON "products" ("token");
COMMIT;
