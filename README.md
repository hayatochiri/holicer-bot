# Holicer bot

## Commands

- グループ追加
    * お店ID(MUST)
- グループ削除
    * グループID(MUST)
- ユーザー追加
    * ユーザー名(MUST)
    * グループID(MUST)
    * 日時指定(MAY:or現在時刻)
- ユーザー離脱(支払)
    * ユーザーID(MUST)
    * 金額指定(MUST)
    * 日時指定(MAY:or現在時刻)
- ユーザーリスト取得
    * お店ID(MAY)
    * グループID(MAY)
- お店追加
    * 店名(MUST)
- お店リスト取得
- お店編集
    * お店ID(MUST)
    * 店名(MUST)
- お店削除
    * お店ID(MUST)
- 飲み物追加
    * お店ID(MUST)
    * 飲み物名(MUST)
    * 代金(MUST)
- 飲み物リスト取得
    * お店ID(MAY)
- 飲み物編集
    * 飲み物ID(MUST)
    * 飲み物名(MAY)
    * 代金(MAY)
- 飲み物削除
    * 飲み物ID(MUST)
- ドリンク代追加
    * ユーザーID(MUST)
    * 金額指定 or 飲み物ID(MUST)
    * 割り勘(MAY:or 1) 1つの飲み物を割り勘する人数(単純にその数で割るだけ)
- ユーザー飲み物オーダーリストを取得
    * ユーザーID(MAY)
    * グループID(MAY)
    * お店ID(MAY)
- ユーザー飲み物オーダーリストを注文済みに変更
    * オーダーID(MUST)
- ユーザー飲み物オーダーリストから削除
    * オーダーID(MUST)
- ユーザー現在の代金照会
    * ユーザーID(MUST)
- 清算
    * グループID(MUST)
    * 金額指定(MUST)
- 清算記録の参照
    * グループID(MUST)
- ユーザーBAN
    * ユーザーID(MUST)
- ユーザーUNBAN
    * ユーザーID(MUST)
- 離脱ユーザーリスト取得
    * お店ID(MAY)
    * グループID(MAY)
- BANユーザーリスト取得
- 支払い情報クリア
    * グループID(MUST)
- データベースクリア
- ログ取得
    * お店ID(MAY)
    * グループID(MAY)
    * ユーザーID(MAY)

## Database

### master
* db_version

### tavern
* id
* name_jp
* name_en
* is_removed (boolean)

### gropu
* id
* uuid
* name_jp
* name_en
* started-time
* tavern-id
* total-price
* tax-rate ( float or null:tax included )
* cleard-time

### user
* id
* twitter-id
* e-mail
* name
* avator
* status { leave, join, ban }
* group-id (when status:join)

### drink
* id
* tavern-id
* name_jp
* name_en
* price
* tax-rate ( float or null:tax included )
* is_removed (boolean)

### user-log
* id
* user-id
* group-id
* status { join, leave, ban, unban }
* timestamp

### leave-log
* id
* user-log-id
* pay (integer or null:credit)

### order-log
* id
* user-id
* group-id
* drink-id
* split (integer 1..)
* status { request, orderd, removed }
* timestamp
