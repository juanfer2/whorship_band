ALTER TABLE performance_songs DROP FOREIGN KEY fk_song;
ALTER TABLE performance_songs DROP FOREIGN KEY fk_performance;
--
DROP INDEX IF EXISTS idx_performance_songs_deleted_at;
DROP INDEX IF EXISTS idx_performance_songs_name;
DROP TABLE IF EXISTS performance_songs;
