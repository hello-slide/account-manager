# Account Manager

アカウント操作をするAppです。

- アカウント作成
- アカウント削除
- 認証

を行います。

## 環境変数

```env
GOOGLE_CLIENT_ID=******
GOOGLE_CLIENT_SECRET=*****
SEED=*****
TOKEN_MANAGER="token-manager"
USER_DATA_STATE="user-data-state"
USER_EMAIL_STATE="user-email-state"
REFRESH_TOKEN_STATE="login-token-state"
API_URL="https://api.hello-slide.jp"
```

## Paths

- `/account/login`
  - ログイン
- `/account/login/redirect`
  - Google OAuthのリダイレクト先
- `/account/update`
  - リフレッシュトークンを使用したセッショントークンの更新
- `/account/logout`
  - ログアウト。ユーザ情報は保持します。
- `/account/delete`
  - アカウント削除。ユーザ情報も削除されます。

## LICENSE

[MIT](./LICENSE)
