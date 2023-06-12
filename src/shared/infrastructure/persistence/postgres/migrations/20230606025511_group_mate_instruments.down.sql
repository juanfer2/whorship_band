ALTER TABLE group_mate_instruments DROP FOREIGN KEY fk_group_mate;
ALTER TABLE group_mate_instruments DROP FOREIGN KEY fk_instrument;
--
DROP INDEX IF EXISTS idx_group_mate_instruments_deleted_at;
DROP INDEX IF EXISTS idx_group_mate_instruments_name;
DROP TABLE IF EXISTS group_mate_instruments;
