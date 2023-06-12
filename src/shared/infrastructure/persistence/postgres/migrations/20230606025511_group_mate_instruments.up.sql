BEGIN;
  CREATE TABLE group_mate_instruments (
    -- id SERIAL
    id uuid DEFAULT uuid_generate_v4 ()
    --
    ,group_mate_id uuid
    ,instrument_id uuid
    --
    ,created_at timestamptz  NOT NULL  DEFAULT current_timestamp
    ,updated_at timestamptz  NOT NULL  DEFAULT current_timestamp
    -- soft delete
    ,deleted_at   timestamptz
    --
    ,PRIMARY KEY (id)
    --
    ,CONSTRAINT fk_group_mate
      FOREIGN KEY(group_mate_id) 
	    REFERENCES group_mates(id)
    --
    ,CONSTRAINT fk_instrument
      FOREIGN KEY(instrument_id)
	    REFERENCES instruments(id)
  );
COMMIT;
