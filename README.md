# Go-platform

![CleanShot 2023-08-14 at 12 21 35@2x](https://github.com/eternaleight/go-platform/assets/96198088/56572fe9-715e-4f9b-9a3e-128bfcc0ffb9)
### User
![User 1degree](https://github.com/eternaleight/go-platform/assets/96198088/2d400696-741b-47b4-b4cd-9587d6d53f81)
### Post
![Post 1degree](https://github.com/eternaleight/go-platform/assets/96198088/aa0ce912-e8cb-4e3b-ad28-1f73cc71981a)
### Profile
![Profile 1degree](https://github.com/eternaleight/go-platform/assets/96198088/f02fd20f-f270-4a44-9e28-6b3e2182c83e)
### Relationships
![relationships real large](https://github.com/eternaleight/go-platform/assets/96198088/947c9dd0-3b12-4d80-a0a7-3efa642c3b9d)

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
