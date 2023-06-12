ALTER TABLE performance_group_mates DROP FOREIGN KEY fk_group_mate;
ALTER TABLE performance_group_mates DROP FOREIGN KEY fk_performance;
--
DROP INDEX IF EXISTS idx_performance_group_mates_deleted_at;
DROP INDEX IF EXISTS idx_performance_group_mates_name;
DROP TABLE IF EXISTS performance_group_mates;
