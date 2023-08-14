# Go-platform

![CleanShot 2023-08-15 at 04 16 56@2x](https://github.com/eternaleight/go-platform/assets/96198088/ab862886-94d2-49f6-8b6d-6e15055064fd)

### User
![User 1degree](https://github.com/eternaleight/go-platform/assets/96198088/2d400696-741b-47b4-b4cd-9587d6d53f81)
### Post
![Post 1degree](https://github.com/eternaleight/go-platform/assets/96198088/aa0ce912-e8cb-4e3b-ad28-1f73cc71981a)
### Profile
![Profile 1degree](https://github.com/eternaleight/go-platform/assets/96198088/f02fd20f-f270-4a44-9e28-6b3e2182c83e)
### Product
![Product 1degree](https://github.com/eternaleight/go-platform/assets/96198088/c3fd9ebe-b4d2-4661-9331-138589e54d80)
### Purchase
![Purchase 2degrees](https://github.com/eternaleight/go-platform/assets/96198088/53c8c042-97f9-47d8-af46-33ff108e18a6)
### Relationships
![relationships real compact](https://github.com/eternaleight/go-platform/assets/96198088/cd6e48f3-764e-4a8c-b9d7-144c13ab6962)


## Project structure
```
.
├── README.md                      // プロジェクトの基本情報
├── api                            // API関連の主要なコードやツールを格納
│   ├── handlers                   // Webリクエストを処理するための関数
│   │   ├── auth.go                // 認証処理（ログイン、登録）
│   │   ├── posts.go               // 投稿の作成や取得に関する処理
│   │   └── user.go                // ユーザー情報の取得や更新のための処理
│   ├── middlewares                // リクエストやレスポンスの前後で実行される関数
│   │   └── isAuthenticated.go     // 認証状態の確認と処理
│   └── responses                  // APIからの応答を生成するヘルパー関数
│       ├── error.go               // エラー応答の生成
│       └── success.go             // 成功応答の生成
├── go.mod                         // プロジェクトの依存関係やモジュール情報
├── go.sum                         // 依存関係の確認用のチェックサムデータ
├── main.go                        // アプリケーションの開始点。サーバの設定や初期化を含む
├── models                         // データベースのテーブルと一致するGoの構造体
│   └── models.go                  // User, Postなどのデータ構造の定義
└── store                          // データベースとのやり取りを行う関数
    ├── auth_store.go              // ユーザーの認証や登録のためのデータベース処理
    ├── post_store.go              // 投稿の作成、取得、更新のデータベース処理
    └── user_store.go              // ユーザー情報の取得や更新のデータベース処理
```
