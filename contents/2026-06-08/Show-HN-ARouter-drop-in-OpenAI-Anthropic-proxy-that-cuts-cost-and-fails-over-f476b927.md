---
source: "https://github.com/sricola/arouter"
hn_url: "https://news.ycombinator.com/item?id=48440512"
title: "Show HN: ARouter – drop-in OpenAI/Anthropic proxy that cuts cost and fails over"
article_title: "GitHub - sricola/arouter · GitHub"
author: "sricola"
captured_at: "2026-06-08T04:36:51Z"
capture_tool: "hn-digest"
hn_id: 48440512
score: 1
comments: 1
posted_at: "2026-06-08T01:49:33Z"
tags:
  - hacker-news
  - translated
---

# Show HN: ARouter – drop-in OpenAI/Anthropic proxy that cuts cost and fails over

- HN: [48440512](https://news.ycombinator.com/item?id=48440512)
- Source: [github.com](https://github.com/sricola/arouter)
- Score: 1
- Comments: 1
- Posted: 2026-06-08T01:49:33Z

## Translation

タイトル: Show HN: ARouter – コスト削減とフェイルオーバーを実現するドロップイン OpenAI/Anthropic プロキシ
記事タイトル: GitHub - sricola/arouter · GitHub
説明: GitHub でアカウントを作成して、sricola/arouter の開発に貢献します。

記事本文:
GitHub - sricola/arouter · GitHub
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
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
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
スリコラ
/
ルーター
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
25

0 コミット 250 コミット .github/ workflows .github/ workflows conformance conformance crates cratesdeploy/ helm/ arouterdeploy/ helm/ arouter docs docs e2e e2e 例 例 .dockerignore .dockerignore .gitignore .gitignore CHANGELOG.md CHANGELOG.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml Dockerfile Dockerfile README.md README.md SECURITY.md SECURITY.md arouter.example.toml arouter.example.toml ポリシー.example.yaml ポリシー.example.yaml Rust-toolchain.toml Rust-toolchain.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
OpenAI および Anthropic 用のドロップイン LLM ルーティング プロキシ — Rust で書かれています。
既存の OpenAI または Anthropic SDK を ARouter に指定すると、すべてのものがルーティングされます。
ポリシーによるリクエスト — コストの削減、プロバイダー間のフェールオーバー、または適切なサイズの調整のため
最も安価な機能モデル — アプリケーション コードを変更する必要はありません。
openaiインポートからOpenAI
# 同じ SDK、同じ呼び出し - Base_url のみが変更されます。
# ARouter は、ポリシーに従って、実際に各リクエストを処理するモデル/プロバイダーを決定します。
client = OpenAI (base_url = "http://localhost:8080/v1" 、api_key = "unused" )
クライアント。チャット 。完成品。 create (モデル = "gpt-4o-mini" , メッセージ = [{ "ロール" : "ユーザー" , "コンテンツ" : "hi" }])
ARルーターを選ぶ理由
🔌ドロップイン。 OpenAI または Anthropic SDK を保持します — ベース URL のみ
変化します。どちらの SDK も他のプロバイダーに到達することもできます。 ARouter の翻訳
ワイヤー。 (適合スイートは、CI を実行するたびに両方の実際の SDK を実行します。)
💸支出を減らします。些細な大部分を適切なサイズにして、
コード変更ではなく、1 行のポリシー ルールです。
(コストデモ → ダウンシフトのリクエストにより最大 94% 安くなります。)
🛟起きてください。モデルを他のプロバイダーのバックアップにヘッジするため、
停止はサイレントにフェイルオーバーします。
(フェイルオーバーのデモ → 停止してもクライアントに表示されるエラーは 0 件。)
✅ より良い出力を出荷します。 V

JSON スキーマに対して応答を検証し、
コントラクトが失敗した場合に、より強力なモデルに自己修復します。
🔭 すべてを見る。すべての応答には x-arouter-* 決定ヘッダーが含まれます。
Prometheus メトリクスと OpenTelemetry トレースが組み込まれています。
2 つのプロバイダー (OpenAI + Anthropic) · 2 つのワイヤ フォーマット · 1 つの静的
バイナリ。現在のリリース: v3.0.0 (変更履歴)。
git clone https://github.com/sricola/arouter
CDルーター
カーゴビルド --release
cp arouter.example.toml arouter.toml
エクスポート AROUTER_OPENAI_API_KEY=sk-...
import AROUTER_ANTHROPIC_API_KEY=sk-ant-... # オプション — クロード* モデルを有効にします
./target/release/arouter --config arouter.toml
ARouter は現在、 http://localhost:8080 をリッスンしています。任意の OpenAI をポイントするか、
Anthropic SDK — モデル名によって、どのアップストリームがサービスを提供するかが決まります。
リクエスト (claude-* → Anthropic、その他すべて → OpenAI):
openaiインポートからOpenAI
client = OpenAI (base_url = "http://localhost:8080/v1" 、api_key = "anything")
# 通常の OpenAI 呼び出し — 変更なし:
クライアント。チャット 。完成品。 create (モデル = "gpt-4o-mini" , メッセージ = [{ "ロール" : "ユーザー" , "コンテンツ" : "hi" }])
# 同じ SDK ですが、Claude モデル — ARouter はリクエストとレスポンスを変換します。
クライアント。チャット 。完成品。 create (model = "claude-sonnet-4-5" ,messages = [{ "role" : "user" , "content" : "hi" }])
ARouter は、arouter.toml 内のキー (または一致するキー) を使用してアップストリームを認証します。
AROUTER_* 環境変数);アプリが送信する api_key は無視されます。少なくとも 1 つ
プロバイダーを構成する必要があります。両方は共存できます。
⚠️ ARouter はデフォルトで 127.0.0.1 にバインドし、認証しません
受信クライアント — アクセスできる人は誰でもプロバイダー キーを使用できます。以前
localhost を超えて公開する場合は、「セキュリティとデプロイメント」を参照してください。
ARouter は、単一のルーティング上で 2 つのエンドポイント (ワイヤ形式ごとに 1 つ) を公開します。
パイプライン。 E

どちらの SDK もいずれかのプロバイダーを駆動できます。応答は常に返されます。
どのアップストリームが応答したかに関係なく、リクエストが到着した形式に戻ります。
# 逆の方向: Anthropic SDK が OpenAI モデルを呼び出す — も機能します。
anthropic インポートより Anthropic
client = Anthropic (base_url = "http://localhost:8080" 、api_key = "unused" )
クライアント。メッセージ。 create (model = "gpt-4o-mini" , max_tokens = 256 ,messages = [{ "role" : "user" , "content" : "hi" }])
ストリーミング、ツール呼び出し、およびクロスワイヤ変換はすべて、
適合スイート。仕組み:
2 線式フォーマット。
router.policy_file をルールの YAML ファイルに指定します。ルールが評価される
あらゆるリクエストに対してトップダウンで対応します。最初の一致によりモデルを書き換え、検証することができます。
応答、フォールバック チェーンの設定、または別のプロバイダーへのヘッジを行います。一致しない→
リクエストはそのまま通過します。
最高の ROI ルールは、プロバイダー内で適切なサイジングを行うことです。つまり、些細なメッセージを送信します。
大部分を安価な層に配置し、必要なもののためにフラッグシップを予約します。の
安価な製品と主力製品の価格差は通常 10 ～ 20 倍であり、同じプロバイダーのスワップは
モデル文字列が異なるだけです (同じワイヤ形式、トークナイザー、および認証)。
リスクはほとんどありません。
# Policy.yaml — 完全なリファレンスについては、policy.example.yaml を参照してください。
ルール：
# OpenAI 内の些細な多数派 -> 安価な層。
- 名前 : trivial_to_mini
いつ:
tokens_input_lt : 250
no_tools : true
no_image : true
次に:
モデル: gpt-4o-mini
# ヘビーコンテキスト -> 旗艦を予約します。
- 名前：heavy_context_to_flagship
いつ:
tokens_input_gt : 6000
次に:
モデル：gpt-4.1
📊 実際に動作を確認してください。例/コスト削減-デモ/
この正確なポリシーを通じて、公式の openai および anthropic SDK を推進します
ライブ API と比較して — ダウンシフトされたリクエストで最大 94% 安くなります
(OpenAI; Anthropic の狭い価格はしごで最大 55%)、

同等かそれ以上の水準で
アプリケーションの変更を一切行わずにスループットを向上させます。
コストを超えて、ルールは検証と自己修復を行うことができます (次の場合にモデル ラダーを再試行します)。
JSON スキーマ コントラクトが失敗する）とヘッジ（JSON スキーマ コントラクトの後で 2 番目のプロバイダーと競合する）
遅延があるため、停止またはスローテールが発生してもサイレントにフェイルオーバーします)。完全な参照:
ルール・
条件・
バリデータ ·
フォールバックチェーン ·
ヘッジ。
完全なドキュメントは docs/ (mdBook) にあります。トピック別:
ここから始めましょう — はじめに · OpenAI クイックスタート · Anthropic クイックスタート
ルーティングとポリシー — ルール、条件、決定、バリデーター、フォールバック チェーン、ヘッジ
アーキテクチャ — 自己修復 · 忠実なプロキシの不変条件 · 事前フラッシュバッファ
ワイヤの互換性 — 2 つのワイヤ形式 · 応答形状デルタ
可観測性 — 応答ヘッダー · Prometheus メトリクス · OpenTelemetry
運用 — セキュリティと展開 · 構成 · Docker · Helm · ベンチマーク
例 — コストデモ、フェイルオーバーデモ、適合スイート
貨物テスト --ワークスペース
カーゴクリッピー --workspace --all-targets -- -D 警告
カーゴFMT --all
荷台
ほとんどのテストは、ルーターをインプロセスで実行します。ブラックボックス層が実際のレイヤーを起動します
ローカルのモックアップストリームに対するソケット経由のルーターバイナリ (API キーなし)、
構成のロード、ソケットバインディング、および回線上の応答について説明します。「」を参照してください。
e2e/ 。適合/スイートは実際の openai を実行します
そして、CI 実行ごとに ARouter に対する人為的な Python SDK を使用します。
sricola.github.io/arouter/
リソース
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
4
v3.0.0
最新の
2026 年 6 月 7 日
+ 3 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私のものを共有しないでください

個人情報

## Original Extract

Contribute to sricola/arouter development by creating an account on GitHub.

GitHub - sricola/arouter · GitHub
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
sricola
/
arouter
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
250 Commits 250 Commits .github/ workflows .github/ workflows conformance conformance crates crates deploy/ helm/ arouter deploy/ helm/ arouter docs docs e2e e2e examples examples .dockerignore .dockerignore .gitignore .gitignore CHANGELOG.md CHANGELOG.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml Dockerfile Dockerfile README.md README.md SECURITY.md SECURITY.md arouter.example.toml arouter.example.toml policies.example.yaml policies.example.yaml rust-toolchain.toml rust-toolchain.toml View all files Repository files navigation
A drop-in LLM routing proxy for OpenAI and Anthropic — written in Rust.
Point your existing OpenAI or Anthropic SDK at ARouter and it routes every
request by policy — to cut cost, fail over between providers, or right-size to
the cheapest capable model — with no change to your application code .
from openai import OpenAI
# Same SDK, same call — only the base_url changes.
# ARouter decides which model/provider actually serves each request, per your policy.
client = OpenAI ( base_url = "http://localhost:8080/v1" , api_key = "unused" )
client . chat . completions . create ( model = "gpt-4o-mini" , messages = [{ "role" : "user" , "content" : "hi" }])
Why ARouter?
🔌 Drop-in. Keep your OpenAI or Anthropic SDK — only the base URL
changes. Either SDK can even reach the other provider; ARouter translates on
the wire. (A conformance suite runs both real SDKs against it on every CI run.)
💸 Spend less. Right-size the trivial majority to a cheaper model with a
one-line policy rule, not a code change.
( cost demo → ~94% cheaper on the requests it downshifts.)
🛟 Stay up. Hedge a model onto a backup on the other provider so an
outage fails over silently.
( failover demo → 0 client-visible errors through an outage.)
✅ Ship better output. Validate responses against a JSON schema and
self-heal to a stronger model when the contract fails.
🔭 See everything. Every response carries x-arouter-* decision headers;
Prometheus metrics and OpenTelemetry traces are built in.
Two providers (OpenAI + Anthropic) · two wire formats · one static
binary . Current release: v3.0.0 ( CHANGELOG ).
git clone https://github.com/sricola/arouter
cd arouter
cargo build --release
cp arouter.example.toml arouter.toml
export AROUTER_OPENAI_API_KEY=sk-...
export AROUTER_ANTHROPIC_API_KEY=sk-ant-... # optional — enables claude-* models
./target/release/arouter --config arouter.toml
ARouter is now listening on http://localhost:8080 . Point any OpenAI or
Anthropic SDK at it — the model name decides which upstream serves the
request ( claude-* → Anthropic, everything else → OpenAI):
from openai import OpenAI
client = OpenAI ( base_url = "http://localhost:8080/v1" , api_key = "anything" )
# Your normal OpenAI call — unchanged:
client . chat . completions . create ( model = "gpt-4o-mini" , messages = [{ "role" : "user" , "content" : "hi" }])
# Same SDK, but a Claude model — ARouter translates the request and response:
client . chat . completions . create ( model = "claude-sonnet-4-5" , messages = [{ "role" : "user" , "content" : "hi" }])
ARouter authenticates upstream with the key in arouter.toml (or the matching
AROUTER_* env var); the api_key your app sends is ignored. At least one
provider must be configured — both can coexist.
⚠️ ARouter binds 127.0.0.1 by default and does not authenticate
inbound clients — anyone who can reach it can use your provider keys. Before
exposing it beyond localhost, read Security & deployment .
ARouter exposes two endpoints — one per wire format — over a single routing
pipeline. Either SDK can drive either provider: the response always comes
back in the shape the request arrived in, no matter which upstream answered.
# The other direction: Anthropic SDK calling an OpenAI model — also works.
from anthropic import Anthropic
client = Anthropic ( base_url = "http://localhost:8080" , api_key = "unused" )
client . messages . create ( model = "gpt-4o-mini" , max_tokens = 256 , messages = [{ "role" : "user" , "content" : "hi" }])
Streaming, tool calls, and the cross-wire translation are all exercised by the
conformance suite . How it works:
Two wire formats .
Point router.policy_file at a YAML file of rules. Rules are evaluated
top-down on every request; the first match can rewrite the model, validate the
response, set a fallback chain, or hedge to another provider. No match → the
request passes through untouched.
The highest-ROI rule is right-sizing within a provider — send the trivial
majority to the cheap tier and reserve the flagship for what needs it. The
cheap-vs-flagship price gap is routinely 10–20×, and a same-provider swap is
just a different model string (same wire format, tokenizer, and auth), so it
carries almost no risk.
# policies.yaml — see policies.example.yaml for the full reference.
rules :
# The trivial majority -> the cheap tier, within OpenAI.
- name : trivial_to_mini
when :
tokens_input_lt : 250
no_tools : true
no_image : true
then :
model : gpt-4o-mini
# Heavy context -> reserve the flagship.
- name : heavy_context_to_flagship
when :
tokens_input_gt : 6000
then :
model : gpt-4.1
📊 See it in action. examples/cost-savings-demo/
drives the official openai and anthropic SDKs through this exact policy
against the live APIs — ~94% cheaper on the requests it downshifts
(OpenAI; ~55% on Anthropic's narrower price ladder), at comparable-or-better
throughput, with zero application changes.
Beyond cost, rules can validate-and-self-heal (retry up a model ladder when
a JSON-schema contract fails) and hedge (race a second provider after a
delay, so an outage or a slow tail fails over silently). Full reference:
Rules ·
Conditions ·
Validators ·
Fallback chains ·
Hedging .
Full docs live in docs/ (an mdBook). By topic:
Start here — Introduction · OpenAI quickstart · Anthropic quickstart
Routing & policy — Rules · Conditions · Decisions · Validators · Fallback chains · Hedging
Architecture — Self-healing · Faithful-proxy invariants · Pre-flush buffer
Wire compatibility — Two wire formats · Response-shape deltas
Observability — Response headers · Prometheus metrics · OpenTelemetry
Operations — Security & deployment · Configuration · Docker · Helm · Benchmarks
Examples — Cost demo · Failover demo · Conformance suite
cargo test --workspace
cargo clippy --workspace --all-targets -- -D warnings
cargo fmt --all
cargo bench
Most tests drive the router in-process. A black-box layer launches the real
arouter binary over a socket against a local mock upstream (no API keys),
covering config loading, socket binding, and the on-the-wire response — see
e2e/ . The conformance/ suite runs the real openai
and anthropic Python SDKs against ARouter on every CI run.
sricola.github.io/arouter/
Resources
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
4
v3.0.0
Latest
Jun 7, 2026
+ 3 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
