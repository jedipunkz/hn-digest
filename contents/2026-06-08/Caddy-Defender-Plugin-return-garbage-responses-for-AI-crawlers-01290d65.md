---
source: "https://github.com/jasonlovesdoggo/caddy-defender"
hn_url: "https://news.ycombinator.com/item?id=48444179"
title: "Caddy Defender Plugin: return garbage responses for AI crawlers"
article_title: "GitHub - JasonLovesDoggo/caddy-defender: Caddy module to block or manipulate requests originating from AIs or cloud services trying to train on your websites · GitHub"
author: "hamburgererror"
captured_at: "2026-06-08T13:35:41Z"
capture_tool: "hn-digest"
hn_id: 48444179
score: 3
comments: 1
posted_at: "2026-06-08T11:52:07Z"
tags:
  - hacker-news
  - translated
---

# Caddy Defender Plugin: return garbage responses for AI crawlers

- HN: [48444179](https://news.ycombinator.com/item?id=48444179)
- Source: [github.com](https://github.com/jasonlovesdoggo/caddy-defender)
- Score: 3
- Comments: 1
- Posted: 2026-06-08T11:52:07Z

## Translation

タイトル: Caddy Defender プラグイン: AI クローラーのガベージ レスポンスを返す
記事のタイトル: GitHub - JasonLovesDoggo/caddy-defender: Web サイトでトレーニングしようとする AI またはクラウド サービスからのリクエストをブロックまたは操作するための Caddy モジュール · GitHub
説明: Web サイトでトレーニングしようとする AI またはクラウド サービスからのリクエストをブロックまたは操作する Caddy モジュール - JasonLovesDoggo/caddy-defender

記事本文:
GitHub - JasonLovesDoggo/caddy-defender: Web サイトでトレーニングしようとする AI またはクラウド サービスからのリクエストをブロックまたは操作するための Caddy モジュール · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ジェイソン大好き犬
/
キャディディフェンダー
公共
ああ、ああ！
途中でエラーが発生しました

オーディング。このページをリロードしてください。
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
JasonLovesDoggo/キャディディフェンダー
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
316 コミット 316 コミット .github .github .idea .idea ベンチマーク/ランダム ベンチマーク/ランダム キャッシュ キャッシュ開発 開発ドキュメント ドキュメントの例 例 マッチャー マッチャー 範囲 範囲 レスポンダー レスポンダー .dockerignore .dockerignore .gitignore .gitignore .golangci.yaml .golangci.yaml CONTRIBUTING.md COTRIBUTING.md Dockerfile Dockerfile FEATURED FEATURED LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md config.go config.go config_test.go config_test.go go.mod go.mod go.sum go.sum middleware.go middleware.go middleware_test.go middleware_test.go mkdocs.yml mkdocs.yml plugin.go plugin.go すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Caddy Defender プラグインは、クライアントの IP アドレスに基づいてリクエストをブロックまたは操作できるようにする Caddy のミドルウェアです。これは、不要なトラフィックを防止したり、ガベージ応答を返すことによって AI トレーニング データを汚染したりする場合に特に役立ちます。
IP 範囲フィルタリング : 特定の IP 範囲からのリクエストをブロックまたは操作します。
埋め込み IP 範囲 : 一般的な AI サービス (OpenAI、DeepSeek、GitHub Copilot など) の事前定義された IP 範囲。
カスタム IP 範囲 : Caddyfile 設定を介して独自の IP 範囲を追加します。
複数のレスポンダー バックエンド:
ブロック : 403 Forbidden 応答を返します。
カスタム : カスタム メッセージを返します。
Garbage : AI トレーニングを汚染するためにゴミデータを返します。
Redirect : カスタム URL を含む 308 Permanent Redirect 応答を返します。
Ratelimit : Ratelimit リクエスト。 caddy-ratelimit 経由で設定可能。
Tarpit : 低速ですが設定可能な速度でデータをストリーミングし、ボットを停止させ、AI トレーニングを汚染します。
簡単なこと

Caddy Defender プラグインを使用する最初の方法は、事前に構築された Docker イメージを使用することです。
docker pull ghcr.io/jasonlovesdoggo/caddy-defender:latest
コンテナを実行します。
次のコマンドを使用して、 Caddyfile でコンテナーを実行します。
docker run -d \
-- キャディの名前 \
-v /path/to/Caddyfile:/etc/caddy/Caddyfile \
-p 80:80 -p 443:443 \
ghcr.io/jasonlovesdoggo/caddy-defender:最新
/path/to/Caddyfile を Caddyfile へのパスに置き換えます。
他のインストール方法については、オンライン ドキュメントを参照してください。
defender ディレクティブは、Caddy Defender プラグインを設定するために使用されます。次の構文があります。
ディフェンダー <レスポンダー> {
メッセージ <カスタムメッセージ>
範囲 <ip_ranges...>
URL <URL>
}
<responder> : 使用するレスポンダー バックエンド。サポートされている値は次のとおりです。
block : 403 Forbidden 応答を返します。
custom : カスタム メッセージを返します ( message が必要です)。
ガベージ : AI トレーニングを汚染するガベージ データを返します。
redirect : 308 Permanent Redirect 応答を返します ( url が必要です)。
ratelimit : レート制限のリクエストをマークします (Caddy-Ratelimit もインストールする必要があります)。
tarpit : 低速ですが設定可能な速度でデータをストリーミングし、ボットを停止させて AI トレーニングを汚染します。
<ip_ranges...> : クライアントの IP と照合する CIDR 範囲または事前定義された範囲キーのオプションのリスト。デフォルトは aws azurepubliccloud deepseek gcloud githubcopilot openai です。
<custom message> : カスタム レスポンダーの使用時に返されるカスタム メッセージ。
<url> : リダイレクト レスポンダーがリダイレクトする URI。
設定の詳細については、Web サイトの設定ページを参照してください。
ドキュメント Web サイトには、プラグインの構成、コード例などの情報が含まれています。
簡単に開始するには、「Getting Started」ガイドに従って、Caddy Defender Plugin を使用してサーバーを保護します。
のために

例については、docs/examples.md を確認してください。
このプラグインには、一般的な AI サービス用の事前定義された IP 範囲が含まれています。これらの範囲はバイナリに埋め込まれているため、追加の構成を行わずに使用できます。
デフォルトでは無効になっています (ビルド時に手動で含める必要があります)
サービス
キー
IP範囲
Tor 出口ノード
トール
トルゴー
ASN (自律システム番号)
として
asn.go
それ以上の方も大歓迎です！プリコンパイルされたリストについては、埋め込まれた結果を参照してください。
寄付を歓迎します!開始するには、 CONTRIBUTING.md を参照してください。
このプロジェクトは MIT ライセンスに基づいてライセンスされています。詳細については、LICENSE ファイルを参照してください。
このプロジェクトのインスピレーション。
bart - Karl Gaissmaier の効率的なルーティング テーブル実装 (バランス ART 適応) により、高性能 IP マッチングが可能になります。
Web サイト上でトレーニングを試みる AI またはクラウド サービスからのリクエストをブロックまたは操作する Caddy モジュール
読み込み中にエラーが発生しました。このページをリロードしてください。
20
フォーク
レポートリポジトリ
リリース
9
v0.10.1
最新の
2026 年 5 月 14 日
+ 8 リリース
このプロジェクトのスポンサーになる
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Caddy module to block or manipulate requests originating from AIs or cloud services trying to train on your websites - JasonLovesDoggo/caddy-defender

GitHub - JasonLovesDoggo/caddy-defender: Caddy module to block or manipulate requests originating from AIs or cloud services trying to train on your websites · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
JasonLovesDoggo
/
caddy-defender
Public
Uh oh!
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
JasonLovesDoggo/caddy-defender
main Branches Tags Go to file Code Open more actions menu Folders and files
316 Commits 316 Commits .github .github .idea .idea benchmarks/ random benchmarks/ random cache cache development development docs docs examples examples matchers matchers ranges ranges responders responders .dockerignore .dockerignore .gitignore .gitignore .golangci.yaml .golangci.yaml CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile FEATURED FEATURED LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md config.go config.go config_test.go config_test.go go.mod go.mod go.sum go.sum middleware.go middleware.go middleware_test.go middleware_test.go mkdocs.yml mkdocs.yml plugin.go plugin.go View all files Repository files navigation
The Caddy Defender plugin is a middleware for Caddy that allows you to block or manipulate requests based on the client's IP address. It is particularly useful for preventing unwanted traffic or polluting AI training data by returning garbage responses.
IP Range Filtering : Block or manipulate requests from specific IP ranges.
Embedded IP Ranges : Predefined IP ranges for popular AI services (e.g., OpenAI, DeepSeek, GitHub Copilot).
Custom IP Ranges : Add your own IP ranges via Caddyfile configuration.
Multiple Responder Backends :
Block : Return a 403 Forbidden response.
Custom : Return a custom message.
Garbage : Return garbage data to pollute AI training.
Redirect : Return a 308 Permanent Redirect response with a custom URL.
Ratelimit : Ratelimit requests, configurable via caddy-ratelimit .
Tarpit : Stream data at a slow, but configurable rate to stall bots and pollute AI training.
The easiest way to use the Caddy Defender plugin is by using the pre-built Docker image.
docker pull ghcr.io/jasonlovesdoggo/caddy-defender:latest
Run the Container :
Use the following command to run the container with your Caddyfile :
docker run -d \
--name caddy \
-v /path/to/Caddyfile:/etc/caddy/Caddyfile \
-p 80:80 -p 443:443 \
ghcr.io/jasonlovesdoggo/caddy-defender:latest
Replace /path/to/Caddyfile with the path to your Caddyfile .
Please see the online documentation for other methods of installation.
The defender directive is used to configure the Caddy Defender plugin. It has the following syntax:
defender <responder> {
message <custom message>
ranges <ip_ranges...>
url <url>
}
<responder> : The responder backend to use. Supported values are:
block : Returns a 403 Forbidden response.
custom : Returns a custom message (requires message ).
garbage : Returns garbage data to pollute AI training.
redirect : Returns a 308 Permanent Redirect response (requires url ).
ratelimit : Marks requests for rate limiting (requires Caddy-Ratelimit to be installed as well ).
tarpit : Stream data at a slow, but configurable rate to stall bots and pollute AI training.
<ip_ranges...> : An optional list of CIDR ranges or predefined range keys to match against the client's IP. Defaults to aws azurepubliccloud deepseek gcloud githubcopilot openai .
<custom message> : A custom message to return when using the custom responder.
<url> : The URI that the redirect responder would redirect to.
For more information about the configuration, refer to the configuration page on the website.
The documentation website has info that includes the configurations of the plugin, code examples, and more.
For a quick start, follow the Getting Started guide to protect your server using the Caddy Defender Plugin .
For examples, check out docs/examples.md
The plugin includes predefined IP ranges for popular AI services. These ranges are embedded in the binary and can be used without additional configuration.
Disabled by default (require manual inclusion at build time)
Service
Key
IP Ranges
Tor Exit Nodes
tor
tor.go
ASN (Autonomous System Numbers)
asn
asn.go
More are welcome! for a precompiled list, see the embedded results
We welcome contributions! To get started, see CONTRIBUTING.md .
This project is licensed under the MIT License . See the LICENSE file for details.
The inspiration for this project .
bart - Karl Gaissmaier 's efficient routing table implementation (Balanced ART adaptation) enabling our high-performance IP matching
Caddy module to block or manipulate requests originating from AIs or cloud services trying to train on your websites
There was an error while loading. Please reload this page .
20
forks
Report repository
Releases
9
v0.10.1
Latest
May 14, 2026
+ 8 releases
Sponsor this project
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
