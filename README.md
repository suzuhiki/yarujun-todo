# yarujun-todo
タスク管理アプリ「やる順Todo」

## 構想
タグや期日でタスクをソートするアプリがある。日頃利用しているタスク管理アプリではタグを使って優先度を表現しているが、それを優先度リストを使って表現できないかと考えた。
「やる順Todo」では10個までのタスクを自由に並べ替えて「やる順」に表示できる。

## 実装の概要
- JWTによるログイン認証
- GinによるAPI実装
- SQL文の学習のためORM不採用
- SwaggerによるAPIドキュメント

未実装の内容
- Flutterの画面においてAPI通信時のリロードを減らしたい
- 特にAPI周りの処理の共通化
- DBの暗号化や環境変数の管理

## 環境
フロントエンドはFlutter、バックエンドはGo、データベースはPostgreSQLを利用する。

DBに初期に作成されるアカウントは `id: admin, pass: admin`
JWTの秘密鍵はコードにベタ書きされている

### フロントエンド
`./frontend/`ディレクトリ以下に配置

環境：MacにFlutterをインストールし、Android StudioかAndroid実機でデバッグ

実行コマンド： `flutter run`

### バックエンド
`./backend/`ディレクトリ以下に配置

環境：MacにGoをインストールして利用

実行コマンド：`air`
Swaggerの更新： `swag init`
DB起動： `docker compose up`
※DBの初期化は `docker compose down -v`

テストAPI：http://localhost:8080/api/v1/test
swagger：http://localhost:8080/swagger/index.html

