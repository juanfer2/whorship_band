BEGIN;
  CREATE TABLE performance_songs (
    -- id SERIAL
    id uuid DEFAULT uuid_generate_v4 ()
    --
    ,performance_id uuid
    ,song_id uuid
    ,tone VARCHAR(255) 
    --
    ,created_at timestamptz  NOT NULL  DEFAULT current_timestamp
    ,updated_at timestamptz  NOT NULL  DEFAULT current_timestamp
    -- soft delete
    ,deleted_at   timestamptz
    --
    ,PRIMARY KEY (id)
    --
    ,CONSTRAINT fk_song
      FOREIGN KEY(song_id) 
	    REFERENCES songs(id)
    --
    ,CONSTRAINT fk_performance
      FOREIGN KEY(performance_id)
	    REFERENCES performances(id)
  );
COMMIT;
