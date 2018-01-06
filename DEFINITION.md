# 各種仕様

## Database

SQLite3 database version 1

#### master

DB自体を管理するためのテーブル。DB設計に変更が必要になった際にアップデート処理が必要となるため、DB設計バージョンを管理する必要がある。

|  Key name  | Variable type | Primary key | Auto increment | not null | Unique | Reference |                 Description                 |
|------------|---------------|-------------|----------------|----------|--------|-----------|---------------------------------------------|
| id         | integer       | o           | o              |          |        |           | DB管理用データのレコードID。値は `1` 固定。 |
| db_version | integer       |             |                |          |        |           | DB設計のバージョンを記録。                  |

#### taverns

居酒屋を管理するためのテーブル。

|  Key name  | Variable type | Primary key | Auto increment | not null | Unique | Reference |                     Description                      |
|------------|---------------|-------------|----------------|----------|--------|-----------|------------------------------------------------------|
| id         | integer       | o           | o              |          |        |           | 居酒屋レコードのID。                                 |
| name_jp    | text          |             |                |          |        |           | 居酒屋の店舗名(日本語)。                             |
| name_en    | text          |             |                |          |        |           | 居酒屋の店舗名(英語)。                               |
| is_removed | integer       |             |                | o        |        |           | 初期値は0。1ならその居酒屋を削除したものとして扱う。 |

#### groups

グループを管理するためのテーブル。メンバーが同一であっても飲み会毎に新しくグループを作成する。(グループ毎に会計を管理するため)

|   Key name   | Variable type | Primary key | Auto increment | not null | Unique | Reference  |                    Description                     |
|--------------|---------------|-------------|----------------|----------|--------|------------|----------------------------------------------------|
| id           | integer       | o           | o              |          |        |            | グループレコードのID。                             |
| uuid         | text          |             |                | o        | o      |            | グループのアクセス用文字列(UUID)。URLに使用する。  |
| name_jp      | text          |             |                |          |        |            | グループ名(日本語)。                               |
| name_en      | text          |             |                |          |        |            | グループ名(英語)。                                 |
| started_time | text          |             |                |          |        |            | 飲み会開始時刻(ISO8601)。                          |
| tavern_id    | integer       |             |                | o        |        | taverns.id | 飲み会を行う居酒屋のID。                           |
| total_price  | integer       |             |                |          |        |            | 飲み会の合計金額。                                 |
| tax_rate     | integer       |             |                |          |        |            | `total_price` が税別なら税率。税込みなら `null` 。 |
| cleard_time  | text          |             |                |          |        |            | 飲み会の清算を行った時刻(ISO8601)。                |

#### users

ユーザーを管理するためのテーブル。

|  Key name  | Variable type | Primary key | Auto increment | not null | Unique | Reference |                                                   Description                                                   |
|------------|---------------|-------------|----------------|----------|--------|-----------|-----------------------------------------------------------------------------------------------------------------|
| id         | integer       | o           | o              |          |        |           | ユーザーレコードのID。                                                                                          |
| twitter_id | text          |             |                |          | o      |           | ツイッターアカウントのID。                                                                                      |
| email      | text          |             |                |          | o      |           | メールアドレス。                                                                                                |
| name       | text          |             |                | o        |        |           | ユーザー名。                                                                                                    |
| avator     | blob          |             |                |          |        |           | ユーザーのアイコン。                                                                                            |
| status     | text          |             |                | o        |        |           | 「 `leave` (離脱), `join` (グループに参加中), `ban` (参加禁止)」のいずれかの状態を記録。デフォルトは `leave` 。 |
| group_id   | integer       |             |                |          |        | groups.id | `status` が `join` の場合のみ、参加しているグループのID。                                                       |

#### menus

メニューを管理するためのテーブル。一度入力したメニューは同じ名前から再度入力できるようにデータベースに入れる。

|  Key name  | Variable type | Primary key | Auto increment | not null | Unique | Reference |                            Description                             |
|------------|---------------|-------------|----------------|----------|--------|-----------|--------------------------------------------------------------------|
| id         | integer       | o           | o              |          |        |           | メニューのレコードのID。                                             |
| tavern_id  | integer       |             |                | o        |        | tavern.id | メニューを紐付ける居酒屋レコードのID。                             |
| name_jp    | text          |             |                |          |        |           | メニュー名(日本語)。                                               |
| name_en    | text          |             |                |          |        |           | メニュー名(英語)。                                                 |
| price      | integer       |             |                | o        |        |           | メニューの値段。                                                   |
| tax_rate   | integer       |             |                |          |        |           | `price` が税別なら税率。税込みなら `null` 。                       |
| is_removed | integer       |             |                | o        |        |           | 初期値は0。1ならそのメニューを削除したものとして扱う(税率変更等)。 |

#### users_log

ユーザーの参加・離脱等ステータス変更をログ管理するテーブル。最終的な料金の清算にも使用する。

|  Key name | Variable type | Primary key | Auto increment | not null | Unique | Reference |                                            Description                                             |
|-----------|---------------|-------------|----------------|----------|--------|-----------|----------------------------------------------------------------------------------------------------|
| id        | integer       | o           | o              |          |        |           | ユーザーログのレコードID。                                                                         |
| user_id   | integer       |             |                | o        |        | users.id  | 状態変化を記録するユーザーのID。                                                                   |
| group_id  | integer       |             |                | o        |        | groups.id | 状態変化を記録した際にユーザーが所属する(していた)グループID。                                     |
| status    | text          |             |                | o        |        |           | 「 `join` (参加), `leave` (離脱), `ban` (参加拒否), `unban` (参加拒否解除)」のいずれかの状態変化。 |
| timestamp | text          |             |                | o        |        |           | ログを記録した時刻(ISO8601)。                                                                      |

#### leave_log

ユーザーの離脱および離脱時の支払い金額を管理するテーブル。

|   Key name  | Variable type | Primary key | Auto increment | not null | Unique |  Reference   |                  Description                   |
|-------------|---------------|-------------|----------------|----------|--------|--------------|------------------------------------------------|
| id          | integer       | o           | o              |          |        |              | ユーザー離脱ログのレコードID。                 |
| user_log_id | integer       |             |                | o        | o      | users_log.id | ユーザーログのレコードID。                     |
| pay         | integer       |             |                |          |        |              | 離脱時に先払いした金額。 `null` ならば後払い。 |

#### orders_log

注文ログを管理するテーブル。

|  Key name | Variable type | Primary key | Auto increment | not null | Unique | Reference |                                       Description                                        |
|-----------|---------------|-------------|----------------|----------|--------|-----------|------------------------------------------------------------------------------------------|
| id        | integer       | o           | o              |          |        |           | 注文ログのレコードID。                                                                   |
| user_id   | integer       |             |                | o        |        | users.id  | 注文したユーザーのID。                                                                   |
| group_id  | integer       |             |                | o        |        | groups.id | 注文したユーザーが所属するグループID。                                                   |
| menu_id   | integer       |             |                | o        |        | menus.id  | 注文したメニューのID。                                                                   |
| split     | integer       |             |                | o        |        |           | 人数割り勘の分母数(デフォルト `1` 、複数人で割り勘する場合に使用)。                      |
| status    | text          |             |                | o        |        |           | 「 `request` (注文待ち), `orderd` (注文済み), `removed` (キャンセル)」のいずれかの状態。 |
| timestamp | text          |             |                | o        |        |           | ログを記録した時刻(ISO8601)。                                                            |
