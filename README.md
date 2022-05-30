# twitter-client-app

※まだ課題としては仕掛り中ですが、確認のため一旦 push しております。

## Overview

Twitter クライアントアプリ。
機能は以下。

1. ツイートする
2. 他ユーザーをフォローする
3. タイムラインを見る
4. 「いいね」を付ける(追加実装)
5. ツイートを検索する(追加実装)

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
- [x] Backend: タイムラインを参照する `/timelines`
- [x] Backend: ツイートする `/tweet`
- [x] Backend: 他のユーザをフォローする `/follow`
- [ ] Backend: いいねをつける `/likes`
- [x] Frontend: 雛形
- [x] Frontend: ログインしてユーザIDを表示
- [x] Frontend: タイムライン参照機能
- [x] Frontend: ツイート機能
- [x] Frontend: ツイートする前に確認ポップアップを出す
- [x] Frontend: ログイン後にユーザ名を表示する
- [x] Frontend: 「いいね」ボタンをアイコンにする
- [ ] Frontend: 「いいね」機能実装
- [x] Backend: ツイートを検索 `/search`
- [x] Backend: 日本語でツイート
- [x] Backend: 日本語で検索
- [x] Frontend: ツイート検索結果に対してフォローできること
- [x] Frontend: ツイート入力値 validation
- [x] Frontend: ツイート検索入力値 validation
- [x] Frontend: ゴミ掃除
- [x] Backend: ディレクトリ構成変更 (backend ディレクトリ作成)
- [x] Backend: Formatter - gofmt
- [x] Backend: Lint - staticcheck
- [ ] Backend: ゴミ掃除
- [ ] Backend: コメント追加
- [ ] デモ動画撮影
- [ ] Backend: Docker化
- [ ] Backend: 複数ユーザ同時接続対応(トークンストア実装)
- [ ] Backend: ログイン済みユーザの対応(再ログインを促す) ⇒おそらくキャッシュの問題
- [ ] Backend: openapi 定義作成
- [ ] Backend: Authorization Code Flow 確認：他ユーザをフォローする
- [ ] Backend: Authorization Code Flow 確認：チャレンジコードをブラウザ・セッションごとに変える
- [ ] Backend: エラーハンドリング
- [ ] Backend: ロギング
- [ ] デプロイの方法
- [ ] 特定のドメイン環境で稼働させる場合の対応手順
- [ ] Backend: セッション管理
- [ ] Backend: 流量制限
- [ ] Backend: テストコード作成

## Components

### Frontend
- Vue
- Docker

### Backend
- go
- echo
- Docker

### Utils
- golangci-lint v1.46.2 (Docker)
- a

## Design

### 基本的な考え方

- 全体
  - 極力ライブラリを利用して、実装すべきコード量を最小化する。
  - ライブラリは、公式もしくはスターの数で信頼性を鑑みて判断する。信頼性の判断が難しいものは使わないようにする。
  - OAuth用のクレデンシャルは、envファイルで管理する。リポジトリにはpushしない。
- Back-end
  - 任意のユーザ利用を想定し、Authorization Code Flow で認証・認可を行う。
  - 短期間&バックエンドの機能性はほぼないため、レイヤードアーキテクチャにせずシンプルに実装する。
  - 複数ユーザのセッション管理に対応
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
