namespace go item

include "./base.thrift"

struct GetItemRequest {
    1: list<i64> item_id_list,
}

struct Item {
    1: i64 item_id,
    2: i64 user_id,
    3: string title,
    4: string video_url,
    5: string label,

    6: i64 status,
    7: i64 rate,
    8: bool is_ecom,
    9: i64 content_level,
    10: Counter counter, // TODO

    11: i64 create_time,
    12: i64 update_time,
}

struct Counter {
    1: i64 play_count,
    2: i64 like_count,
    3: i64 share_count,
    4: i64 comment_count,
}

struct GetItemResponse {
    1: map<i64, Item> item_map,
    255: base.BaseResp BaseResp,
}

struct UpdateItemRequest {
    1: required i64 item_id,
    2: optional string title,
    3: optional string label,
    4: optional i64 status,
    5: optional i64 rate,
    6: optional bool is_ecom,
    7: optional i64 content_level,
}

struct UpdateItemResponse {
    255: base.BaseResp BaseResp,
}

struct CreateItemRequest {
    1: required i64 item_id,
    2: required i64 user_id,
    3: required string title,
    4: required string video_url,
    5: required string label,
}

struct CreateItemResponse {
    255: base.BaseResp BaseResp,
}

struct GetItemListRequest {
    1: optional list<i64> item_id,
    2: optional string title,
    3: optional i64 create_time_from,
    4: optional i64 create_time_to,
    5: optional list<string> label,
    6: optional list<i64> user_id,

    7: optional list<i64> status,
    8: optional list<i64> rate,
    9: optional list<i64> content_level,

    100: required i32 order, // 0:DESC, 1:ASC
    101: required string order_by,
    102: required i64 size,
    103: optional i64 offset,
}

struct GetItemListResponse {
    1: list<Item> item_list,
    2: i64 total,
    255: base.BaseResp BaseResp,
}

struct GetItemChangeHistoryRequest {
    1: optional i64 item_id,
    2: optional i64 time_from,
    3: optional i64 time_to,
    100: optional i32 order, // 0:DESC, 1:ASC
    101: optional string order_by,
    102: optional i64 size,
    103: optional i64 offset,
}

struct ItemChange {
    1: i64 item_id,
    2: i64 time,
    3: string item_before, // json
    4: string item_after, // json
}

struct GetItemChangeHistoryResponse {
    1: list<ItemChange> history,
    2: i64 total,
    255: base.BaseResp BaseResp,
}

service ItemService {
    GetItemResponse GetItem(1: GetItemRequest req)
    UpdateItemResponse UpdateItem(1: UpdateItemRequest req)
    CreateItemResponse CreateItem(1: CreateItemRequest req)
    GetItemListResponse GetItemList(1: GetItemListRequest req)
    GetItemChangeHistoryResponse GetItemChangeHistory(1: GetItemChangeHistoryRequest req)
}
