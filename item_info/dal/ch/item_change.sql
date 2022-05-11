CREATE TABLE default.item_change
(
    item_id Int64,
    time Int64,
    item_before String,
    item_after String
)
ENGINE = MergeTree()
PRIMARY KEY (item_id, time)
