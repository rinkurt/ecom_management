namespace go item

include "./base.thrift"

struct GetItemRequest {
    1: list<i64> item_id_list,
}

struct Item {
    1: i64 item_id,
    2: i64 user_id,
}

struct GetItemResponse {
    1: map<i64, Item> item_map,
    255: base.BaseResp BaseResp,
}

struct UpdateItemRequest {
    1: required i64 item_id,
    2: optional i64 user_id,
}

struct UpdateItemResponse {
    255: base.BaseResp BaseResp,
}

service ItemService {
    GetItemResponse GetItem(1: GetItemRequest req)
    UpdateItemResponse UpdateItem(1: UpdateItemRequest req)
}
