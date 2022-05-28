# twitter-client-app

## Overview

Twitter クライアントアプリ。
機能は以下。

1. ツイートする
2. 他ユーザーをフォローする
3. タイムラインを見る
4. 「いいね」を付ける(追加実装)

Todo: ここに全体構成図を挟む

### 特記事項
- 特定のドメインではなく、現時点では localhost で動く前提にしています。
  - 特定のドメインにする場合は、.env(もしくは環境変数に設定）する。

## Todo

- [ ] 全体構成図作成
- [ ] 基本的な考え方整理
- [x] Twitter App 用のクライアントID、シークレットを取得
- [x] Backend: Twitterログイン(Authorization Code Flow)：認可エンドポイントへのリダイレクト～トークンの取得
- [x] Backend: シークレットを環境変数に設定
- [x] Backend: コード構成見直し(レイヤード構成)
- [x] Backend: Authorization Code Flow 確認：token を使った Twitter API 実行確認 https://api.twitter.com/2/users/me
- [x] Backend: .envだけでなく、環境変数に対応する
- [x] Backend: stateの検証(stateユニーク化)
- [ ] Backend: タイムラインを参照する `/timelines`
- [ ] Backend: openapi 定義作成
- [ ] Backend: Authorization Code Flow 確認：ツイート取得
- [ ] Backend: Authorization Code Flow 確認：ツイートする
- [ ] Backend: Authorization Code Flow 確認：他ユーザをフォローする
- [ ] Backend: Authorization Code Flow 確認：チャレンジコードをブラウザ・セッションごとに変える
- [ ] Backend: エラーハンドリング
- [ ] Backend: ロギング
- [ ] Frontend: 雛形・サンプルアプリ作成
- [ ] Frontend: ツイート機能実装
- [ ] Frontend: フォロー機能実装
- [ ] Frontend: タイムライン参照機能
- [ ] Makefile:
- [ ] デプロイの方法
- [ ] 特定のドメイン環境で稼働させる場合の対応手順

## a

### Frontend
- Vue
- Docker

### Backend
- go
- echo
- Docker

## Design

### 基本的な考え方

- 全体
  - 極力ライブラリを利用して、実装すべきコード量を最小化する。
  - ライブラリは、公式もしくはスターの数で信頼性を鑑みて判断する。信頼性の判断が難しいものは使わないようにする。
  - OAuth用のクレデンシャルは、envファイルで管理する。リポジトリにはpushしない。
- Back-end
  - 任意のユーザ利用を想定し、Authorization Code Flow で認証・認可を行う。
  - 短期間&バックエンドの機能性はほぼないため、レイヤードアーキテクチャにせずシンプルに実装する。
- Front-end
  - aaa
- Enviroment
  - 任意の環境で容易に開発・ビルド・テストできるよう実行環境として Docker を利用する。

### エンドポイント
- `/login`
- `/callback`

## usage

```
make
```

### how to deploy

Todo: CloudRun にデプロイ
