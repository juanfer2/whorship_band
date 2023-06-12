BEGIN;
  CREATE TABLE performance_group_mates (
    -- id SERIAL
    id uuid DEFAULT uuid_generate_v4 ()
    --
    ,performance_id uuid
    ,group_mate_id uuid
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
    ,CONSTRAINT fk_performance
      FOREIGN KEY(performance_id)
	    REFERENCES performances(id)
  );
COMMIT;
