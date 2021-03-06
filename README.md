# twitter-client-app

## Overview

Twitter クライアントアプリ。
機能は以下。

1. ツイートする
2. 他ユーザーをフォローする ※検索したツイート結果に対してフォロー可能
3. タイムラインを見る
4. ツイートを検索する(追加実装)
5. タイムライン上のツイートに「いいね」する(追加実装)

## Demo

![Animation3](https://user-images.githubusercontent.com/24238933/171014845-7219bf49-fa35-4c27-98c1-34d65fd5bea8.gif)

## Usage

### 0. 事前準備

#### ClientID/ClientSecret 取得

Twitter の [Developer POrtal](https://developer.twitter.com/) にて OAuth2.0 アプリケーションを登録し、クライアントID及びクライアントシークレットを取得する。また、Callback URL/Redirect URL を設定する。

### .envファイル or 環境変数を設定

次に、取得した文字列を `backend/.env` に設定する。`REDIRECT_UTL` は固定。
```
CLIENT_ID=xxx
CLIENT_SECRET=xxx
REDIRECT_URL=http://localhost:8081/callback
```

※`CLIENT_ID`及び`CLIENT_SECRET`は置き換えてください。
※上記3つの変数を export するでもOKです。

### 1. サーバ起動

#### Option 1. docker-compose

```
docker-compose up -d
```

上記コマンドでフロントエンド(8080)、バックエンド(8081)両方起動します。

#### Option 2. 手動起動

バックエンドサーバ起動

```
cd backend
make
```
⇒`localhost:8081`で起動

フロントエンドサーバ起動

```
cd frontend
npm install
npm run serve
```
⇒`localhost:8080`で起動

### 2. ブラウザでアプリにアクセス
ブラウザをゲストモードで開いて、`localhost:8080`とアクセスしてください。
（現時点では、Twitterログイン済みのブラウザでは検証していません。）

ブラウザで開いた後は、[Demo](https://github.com/Kaminashi-Inc/ENG-1004_a2-ito/tree/develop#demo) を参考にしながら触ってみてください。まずは右上のボタンからTwitterログインをしてください。

## 注意点

- クライアント環境（ブラウザ）の Cookie やキャッシュに関する制御を実装できていないため、現時点ではChromeのゲストモードを利用して、新規ログインするケースのみ動作確認をすることができます。
- 現時点では `localhost` 上での動作確認のみ行っています。
  - 特定のドメインにする場合の修正箇所は[こちら](https://github.com/Kaminashi-Inc/ENG-1004_a2-ito/tree/develop#%E6%9C%AC%E6%A0%BC%E7%A8%BC%E5%83%8D%E3%81%AB%E5%90%91%E3%81%91%E3%81%A6%E3%82%84%E3%82%8B%E3%81%B9%E3%81%8D%E3%81%93%E3%81%A8)に記載しています。
- トークン情報はクライアントに保持せず、サーバサイドのみ保持しています。また現時点では複数ユーザに対応していないため、別ユーザでログインする場合はバックエンドを再起動してから再度ログインしてください。
  - 複数ユーザ対応は、ユーザID毎に Key-Value 形式で保存すれば良いので難しい対応ではないですが、作業時間の都合上割愛しています。
- Twitter API のスコープ設定は backend/config/config.go 上に埋め込んでいますので、機能追加時に必要に応じて修正・外だししてください。
- 現時点ではローカル環境の構築容易性のみの観点のみで docker を利用しており、 Dockerfile は利用していません。

## Design

### 基本的な考え方

- 全体
  - `golang.org/x/oauth2` 等、極力ライブラリを利用して、実装すべきコード量を最小化する。
    - ライブラリは、公式もしくはスターの数で信頼性を鑑みて判断する。信頼性の判断が難しいものは使わないようにする。（最低でも1k程度）
- Back-end
  - 任意のユーザ利用想定し、また Application としては Public Client であることを想定し、Authorization Code Flow で認証・認可を行った上でAPIを実行できるようにする。
  - OAuth用のクレデンシャル(Client ID/Client Secret)は、envで管理する。リポジトリにはpushしない。
    - .envファイルを利用するか、環境変数で設定。
  - 将来的な拡張等を見据えて、レイヤードアーキテクチャを意識していくつかのディレクトリに分けて実装。(モデルやDB層がないためrepositoryはない)
- Front-end
  - SPA でシンプルに実装。デバッグや実装のしやすさから、特定の単一componentファイルでフロントエンドを実装する。
  - 容易にきれいなデザインにできるため、Vuetify を利用する。

### Components

動作確認済みのバージョンを併せて記載。

#### Frontend
- npm 7.21.0
- node 16.8.0
- Vue 2.6.11
- Vuetify 2.6.0

#### Backend
- go 1.18
- echo v4.7.2

#### Utils
- docker-compose 1.29.2
- docker 20.10.8
- Make
- gofmt
- staticcheck v0.3.2

### エンドポイント
- `GET /login`
  - OAuth2.0 Authorization Code Flow - 認可エンドポイントへリダイレクトする
- `GET /callback`
  - OAuth2.0 Authorization Code Flow - アクセストークンを取得する
- `POST /tweet`
  - ツイートする
- `POST /follow`
  - 特定のユーザをフォローする
- `GET /timelines`
  - 特定のユーザをフォローする
- `GET /search`
  - ツイートを任意の文字列で検索する
- `POST /likes`
  - 特定のツイートを いいね する

## 本格稼働に向けてやるべきこと
- ドキュメンテーション
  - 全体構成図作成
  - Backend: openapi 定義作成
- 機能系
  - Backend: 全般的にパラメータチェックを実装（必須パラメータが入力されてない場合は弾く）
  - Backend: post 系の処理において、クライアントに200を返却して正常にツイートされた、フォローできた、旨を通知する。
  - Frontend: 処理中は in-progress のUIを付ける（処理待ちなのかどうかが不明）
- 非機能系
  - Backend: ログイン済みユーザの対応(再ログインを促す) ⇒おそらくキャッシュの問題
  - Backend: セッション管理
    - redis 等のキャッシュストアを用いて、トークン等の一時的な情報をユーザ毎に保持する。
    - 現状の仕様では他のクライアントが用意になりすましができるので、クライアントに一時的なセッションIDを払い出し、そのセッションIDをもったクライアントでなければトークン利用（API実行）ができないようにする。
  - Backend: テストコード作成
  - Backend: ロギング
  - Backend: Docker化(using builcpacks)
  - Front/Backend コンテナイメージ化
    - クラウドへの可搬性や起動時間の高速化等の要件に応じて実施。
  - Backend: エラーハンドリング
    - エラーを返却しユーザ側にも何が問題かを通知する。
  - Front/Backend: 特定ドメイン化
    - CORS 許可を外す
      - 現状はどちらも`localhost:8080` と `localhost:8081` でレスポンスを返却するため、API側で当該エラーがでないように `middleware.CORS` を設定している。本来は当該設定なしに構成すべきであり、同ドメイン・ポートで動くようにしてから当該設定をはずすこと。
    - Redirect URL 変更
      - ドメイン設定に応じて変更する。
  - Backend: 流量制限
  - Frontend: 静的コンテンツのエクスポート、CDN/Webサーバからの配信
    - 静的コンテンツとしてビルド・エクスポートし、CDN/Webサーバから配信するようにする。

## 参考：実施済みTodoリスト

- [x] 基本的な考え方整理
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
- [x] Frontend: 雛形
- [x] Frontend: ログインしてユーザIDを表示
- [x] Frontend: タイムライン参照機能
- [x] Frontend: ツイート機能
- [x] Frontend: ツイートする前に確認ポップアップを出す
- [x] Frontend: ログイン後にユーザ名を表示する
- [x] Frontend: 「いいね」ボタンをアイコンにする
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
- [x] Backend: ゴミ掃除
- [x] Backend: コメント追加
- [x] デモ動画撮影
- [x] Frontend: npm install 等の事前手順確認
- [x] docker-compose 化
- [x] エンドポイントの説明追加
- [x] Backend: いいねをつける `/likes`

