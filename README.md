# twitter-client-app

## Overview

Twitter クライアントアプリ。
機能は以下。

1. ツイートする
2. 他ユーザーをフォローする
3. タイムラインを見る
4. aaa(追加実装)

Todo: ここに全体構成図を挟む

## Todo

- [ ] 全体構成図作成
- [ ] 基本的な考え方整理
- [x] Twitter App 用のクライアントID、シークレットを取得
- [x] Backend: Twitterログイン(Authorization Code Flow)：認可エンドポイントへのリダイレクト～トークンの取得
- [x] Backend: シークレットを環境変数に設定
- [x] Backend: コード構成見直し(レイヤード構成)
- [x] Backend: Authorization Code Flow 確認：token を使った Twitter API 実行確認 https://api.twitter.com/2/users/me
- [ ] openapi 定義作成
- [ ] Backend: Authorization Code Flow 確認：stateの検証(stateユニーク化)
- [ ] Backend: Authorization Code Flow 確認：ツイート取得
- [ ] Backend: Authorization Code Flow 確認：ツイートする
- [ ] Backend: Authorization Code Flow 確認：他ユーザをフォローする
- [ ] Backend: Authorization Code Flow 確認：タイムラインを参照する
- [ ] Backend: config 外だし
- [ ] Backend: state のユニーク化
- [ ] Backend: エラーハンドリング
- [ ] Backend: ロギング
- [ ] Frontend: 雛形・サンプルアプリ作成
- [ ] Frontend: ツイート機能実装
- [ ] Frontend: フォロー機能実装
- [ ] Frontend: タイムライン参照機能
- [ ] Makefile:
- [ ] デプロイの方法

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
  - 短期間&バックエンドの機能性はほぼないため、レイヤードアーキテクチャにせずシンプルに実装する。
- Front-end
  - aaa
- Enviroment
  - aaa

### エンドポイント
- `/login`
- `/callback`

## usage

```
make
```

### how to deploy

Todo: CloudRun にデプロイ
