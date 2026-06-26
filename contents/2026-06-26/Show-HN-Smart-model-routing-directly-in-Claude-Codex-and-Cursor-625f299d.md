---
source: "https://github.com/workweave/router"
hn_url: "https://news.ycombinator.com/item?id=48688700"
title: "Show HN: Smart model routing directly in Claude, Codex and Cursor"
article_title: "GitHub - workweave/router: Model router for agentic systems. Routes every prompt to the right model in <50ms. Cut costs 40-70% with just an endpoint change. · GitHub"
author: "adchurch"
captured_at: "2026-06-26T17:33:42Z"
capture_tool: "hn-digest"
hn_id: 48688700
score: 25
comments: 4
posted_at: "2026-06-26T16:40:11Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Smart model routing directly in Claude, Codex and Cursor

- HN: [48688700](https://news.ycombinator.com/item?id=48688700)
- Source: [github.com](https://github.com/workweave/router)
- Score: 25
- Comments: 4
- Posted: 2026-06-26T16:40:11Z

## Translation

タイトル: HN を表示: クロード、コーデックス、カーソルで直接スマート モデル ルーティングを行う
記事のタイトル: GitHub - workweave/router: エージェント システムのモデル ルーター。すべてのプロンプトを 50 ミリ秒未満で適切なモデルにルーティングします。エンドポイントを変更するだけでコストを 40 ～ 70% 削減します。 · GitHub
説明: エージェント システム用のモデル ルーター。すべてのプロンプトを 50 ミリ秒未満で適切なモデルにルーティングします。エンドポイントを変更するだけでコストを 40 ～ 70% 削減します。 - ワークウィーブ/ルーター
HN テキスト: 私たちは、コーディング エージェント (Claude Code、Codex、Cursor など) にプラグインし、それらに対応する最適なモデルにリクエストをインテリジェントに送信するモデル ルーターを構築しました。ローカルで実行する簡単なデモは次のとおりです: https://www.youtube.com/watch?v=isKhAyivtfM 。 Weave では、すべてのコードを AI で作成していますが、そのコストはますます高くなっています。この問題は、Opus 4.7 がリリースされたときに問題に直面し、トークナイザーの変更のおかげでコストが跳ね上がりました。すべてに Opus が必要なわけではないことはわかっていましたが、本当に必要な場合に備えたインテリジェンスを失いたくありませんでした。そこで、これを処理するモデル ルーターを構築することにしました。 Weave Router は、コーディング エージェント専用の Anthropic/OpenAI エンドポイントとして機能します。すべての推論リクエストを調べて、どのモデルに送信するかをインテリジェントに決定し (これについてはすぐに説明します)、途中で必要なすべての変換を処理します。そのため、可能な場合はより高速/安価なモデル (例: DeepSeek v4、GLM 5.2、Kimi K2.6) を使用し、必要に応じてフロンティア モデル (Opus 4.8 および GPT 5.5 (および Fable が復活するたびに)) を使用できます。どのモデルにルーティングするかをどうやって知るのでしょうか?私たちは、数万 (これまでのところ!) のエージェント トレースに基づいて RL モデルをトレーニングしました。指定されたタスクを正常に完了した LLM を選択すると、ルーティング モデルに報酬が与えられます。以下に例を示します。ルーターに複雑な変更を計画するよう依頼すると、ルーターはそのリクエストを (おそらく) Opus にルーティングします。

4.8.コンテキストを収集するためにコードベースを探索するサブエージェントは、より適切なモデル (DeepSeek V4 Flash など) にルーティングされます。その後、計画を実装する準備ができたら、(おそらく) より迅速なモデル (GLM 5.2 など) に渡されて実行されます。ここ 1 か月間ほど社内でこれを使用しています。品質や速度に目立った違いはなく、トークンを他の場合に支払う金額と比較して 40% 節約できました。ルーターは Elastic License 2.0 でソースから入手できるため、セルフホストできます。または、必要に応じて、当社のホスト型バージョン (weaverouter.com) を使用することもできます。ご質問があればお答えします。

記事本文:
GitHub - workweave/router: エージェント システムのモデル ルーター。すべてのプロンプトを 50 ミリ秒未満で適切なモデルにルーティングします。エンドポイントを変更するだけでコストを 40 ～ 70% 削減します。 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
作業着

ve
/
ルーター
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
455 コミット 455 コミット .claude/ スキル .claude/ スキル .github/ ワークフロー .github/ ワークフロー cmd cmd db db docs docs フロントエンド フロントエンド インストール インストール 内部 内部スクリプト スクリプト .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md docker-compose.yml docker-compose.yml go.mod go.mod go.sum go.sum すべてのファイルを表示 リポジトリ ファイルのナビゲーション
エンドポイントは 1 つ。どのモデルも。常に正しいものです。
最適なモデルを選択する Anthropic、OpenAI、Gemini 用のドロップイン プロキシ
すべてのリクエストに対して、バイブベースのプロンプトではなく、小さなオンボックスエンベッダーを使用します。
🥇 RouterArena リーダーボードで #1 — Acc-Cost Arena 76.09 。
Weave によって構築: ナンバー 1 のエンジニアリング インテリジェンス プラットフォーム、
Robinhood、PostHog、Reducto、その他何百もの人々に愛されています。
localhost:8080 にあるクロード コード、コーデックス、カーソル、または独自のアプリを指定します。ルーター:
🎯 リクエストごとのルート。以下から派生したクラスター スコアラー
アベンジャーズ プロ 2 は正しい選択をします
有効なプロバイダーからのモデルを毎回提供します。
🔌 みんなの API を話します。人間的メッセージ、OpenAI チャット補完、
双子座出身。ストリーミング、ツール、ビジョン、作品。
🧠 OSSにも詳しい。 DeepSeek、キミ、GLM、クウェン、ラマ、ミストラル経由
OpenRouter (または任意の OpenAI 互換エンドポイント)。
🔒 デフォルトでは BYOK。プロバイダー キーはボックスに保存され、暗号化されて保存されます。
📊 観察可能。 OTLP トレースはすぐに使用できます。 Weave ダッシュボード ( http://localhost:8080/ui/dashboard ) でダッシュボードを確認するか、Honeycomb、Datadog、
グラファナ、w

嫌いです。
最も速い方法: ホストされている場所で Claude Code、Codex、または opencode を指定します。
1 つのコマンドでルーターを織ります。クローンも Docker も Postgres もありません。
npx @workweave/ルーター
それだけです。インストーラーはどのツール (Claude Code、Codex、または opencode) を尋ねます。
スコープ (ユーザーとプロジェクト) を順を追って説明し、ルーター キーを取得して配線します。
適切な設定ファイル。他のフレーバー:
npx @workweave/router --claude # ピッカーをスキップします、クロード コード
npx @workweave/router --codex # ピッカーをスキップ、OpenAI Codex CLI
npx @workweave/router --opencode # ピッカーをスキップし、opencode
npx @workweave/router --scope プロジェクト # リポジトリごと、settings.json (または .codex/ / opencode.json) をコミット
npx @workweave/router --local # 自己ホスト型 localhost:8080
npx @workweave/router --base-url https://router.acme.internal
npx @workweave/router@0.1.0 # バージョンを固定する
Node ≥ 18 が必要です (Claude Code および opencode パスにも jq が必要です)。フル
フラグ参照: install/npm/README.md 。
ルーター (およびダッシュボード) を独自のボックスで実行したい場合:
# 1. プロバイダー キーをドロップします。OpenRouter が推奨されるベースラインです。
echo " OPENROUTER_API_KEY=sk-or-v1-... " >> .env.local
# 2. :8080 で Postgres + ルーターを起動し、rk_ キーをシードします。
フルセットアップを行う
ルーターは http://localhost:8080 で稼働しており、ダッシュボードは次の場所にあります。
http://localhost:8080/ui/ (パスワード: admin )、および rk_... キー
ログに出力されます。
# それを人類的と呼んでください
カール -sS http://localhost:8080/v1/messages \
-H " 権限: ベアラー rk_... " \
-d ' {"モデル":"クロード・ソネット-4-5","max_tokens":256,
"メッセージ":[{"ロール":"ユーザー","コンテンツ":"こんにちは"}]} '
# ...または OpenAI のようなもの
カール -sS http://localhost:8080/v1/chat/completions \
-H " 権限: ベアラー rk_... " \
-d ' {"モデル":"gpt-4o-mini",
"メッセージ":[{"ロール":"ユーザー","コンテンツ":"こんにちは"}]} '
# プロキシを使用せずにルーティング決定を確認する
カール -sS http://localhost:8

080/v1/route -H " 認証: ベアラー rk_... " -d ' ... '
ツールに接続します
クロード・コード。 make install-cc を実行して、ローカルでクロード コードを接続します。
セルフホステッド ルーター (終了時にも自動的に呼び出されます)
フルセットアップを行います）。ホスト型ルーターの場合は、npx @workweave/router を使用します。
上。
コーデックス (OpenAI CLI)。 npx @workweave/router --codex パッチ
~/.codex/config.toml (または --scope project を使用した <repo>/.codex/config.toml)
マネージド [model_providers.weave] ブロックを使用し、 model_provider = "weave" を設定します。
Codex の既存の OPENAI_API_KEY は、api.openai.com に流れます。
プランベースのパススルー。ルーターキーは X-Weave-Router-Key HTTP に組み込まれます。
ヘッダー。再インストールおよび --uninstall --codex の管理対象のみを書き換え/削除します
ブロックし、Codex 構成の残りの部分はそのままにしておきます。
オープンコード。 npx @workweave/router --opencode は Provider.weave をマージします
~/.config/opencode/opencode.json (または <repo>/opencode.json) へのエントリ
--scope プロジェクトを使用)。 Opencode にバンドルされている @ai-sdk/anthropic を使用します
ルーターの /v1 エンドポイントを指すプロバイダー — ルーターは、
Anthropic Messages API はネイティブなので、オープンコードは変更せずに動作します。ルーター
キーと ID ヘッダーはプロバイダー設定と一緒に使用されます。再インストールする
マネージド ブロックのみを書き換え、 --uninstall --opencode によってそれが削除されます。
カーソル (初期のベータ版、パフォーマンスは最高ではない可能性があります)。設定→
モデル → OpenAI ベース URL をオーバーライド → http://localhost:8080/v1 を貼り付けます
rk_... を API キーとして使用します。
オン/オフを切り替えます。インストール後、npx @workweave/router off --claude
(または --codex / --opencode ) クライアントをプロバイダーに直接ルーティングします
ルーター設定を破棄せずに再度実行します。 on で元に戻り、ステータス
それがどちらの方向を指しているかを報告します。クロード コードも /router-off を取得します。
/router-on および /router-status スラッシュ コマンド。カーソルも同じように切り替わります
セ

設定 → モデルは上記をオーバーライドします。 install/README.md を参照してください。
sk-or-... / sk-ant-... / sk-... = 上流プロバイダーキー。 .env.local に住んでいます。
rk_... = ルーターキー。クライアントはこれをベアラー トークンとして送信します。
エンドポイント
フォーマット
POST /v1/messages
人間的なメッセージ、ルーティング済み
POST /v1/chat/completions
OpenAI チャット完了、ルーティング済み
POST /v1beta/models/:action
GeminigenerateContent 、ルーティング済み
POST /v1/ルート
決定を返します。アップストリーム呼び出しはありません
GET /v1/models · POST /v1/messages/count_tokens
人為的パススルー
GET /health · GET /validate
活性+キーチェック
より詳細なドキュメント
📐 設定リファレンス : すべての環境変数、
BYOK 暗号化、OTel ノブ、クラスター ルーティング。
🛠️ 貢献 : 階層化ルール、ホットリロード開発、
移行、テスト、エンジニアリング ループ全体。
🏗️ アーキテクチャ : パッケージ レイアウト、インポート コントラクト、
エンドポイント/プロバイダー/戦略を追加するためのレシピ。
トークン対応レート制限 (インストールごとの Redis スライディング ウィンドウ)
テナント階層のサブインストール
投機的なディスパッチ + テール レイテンシーのヘッジ
Lu, Y.、Liu, R.、Yuan, J.、Cui, X.、Zhang, S.、Liu, H.、Xing, J.
RouterArena: LLM を包括的に比較するためのオープン プラットフォーム
ルーター。 arXiv:2510.00202、2025。 https://arxiv.org/abs/2510.00202 ↩
Zhang、Y.ら。 GPT-5 を超えて: LLM をより安く、より良くする
パフォーマンス効率に最適化されたルーティング (アベンジャーズ プロ)。
arXiv:2508.12631、2025。 https://arxiv.org/abs/2508.12631 ↩
エージェント システムのモデル ルーター。すべてのプロンプトを 50 ミリ秒未満で適切なモデルにルーティングします。エンドポイントを変更するだけでコストを 40 ～ 70% 削減します。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
7
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。

© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Model router for agentic systems. Routes every prompt to the right model in <50ms. Cut costs 40-70% with just an endpoint change. - workweave/router

We built a model router that plugs into coding agents (e.g. Claude Code, Codex, Cursor, etc.) and intelligently sends requests to the best model to serve them. Here's a quick demo of running it locally: https://www.youtube.com/watch?v=isKhAyivtfM . At Weave, we write ~all our code with AI, and it's been getting more expensive. This came to a head when Opus 4.7 was released and, thanks to its tokenizer changes, our costs shot up. We knew we didn't need Opus for everything but we didn't want to lose out on the intelligence for the cases where you really need it. So we decided to build a model router to handle this for us. The Weave Router acts as an Anthropic/OpenAI endpoint specifically for coding agents. It looks at every inference request and intelligently (more on that in a sec) decides what model to send it to, handling all the translations required along the way. So it can use faster/cheaper models (e.g. DeepSeek v4, GLM 5.2, Kimi K2.6) when possible, and frontier models (Opus 4.8 & GPT 5.5 (& Fable whenever it's back)) when necessary. How do we know what model to route to? We trained an RL model on tens of thousands (so far!) of agent traces. We reward the routing model when it selects an LLM that successfully completes the given task. Here's an example: if you ask the router to plan a complex change, it will (probably) route that request to Opus 4.8. Subagents exploring the codebase to gather context will be routed to more suitable models (e.g. DeepSeek V4 Flash). Then when you have the plan ready to implement, it will be (most likely) be handed to a quicker model (e.g. GLM 5.2) to carry it out. We've been using this internally for the last month or so. We've saved 40% on tokens vs. what we otherwise would have paid, with no noticeable differences in quality or velocity. The router is source-available under Elastic License 2.0, so you can self-host it. Or if you prefer, you can also use our hosted version: weaverouter.com. I'll be here to answer any questions you may have!

GitHub - workweave/router: Model router for agentic systems. Routes every prompt to the right model in <50ms. Cut costs 40-70% with just an endpoint change. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
workweave
/
router
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
455 Commits 455 Commits .claude/ skills .claude/ skills .github/ workflows .github/ workflows cmd cmd db db docs docs frontend frontend install install internal internal scripts scripts .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md docker-compose.yml docker-compose.yml go.mod go.mod go.sum go.sum View all files Repository files navigation
One endpoint. Every model. Always the right one.
A drop-in proxy for Anthropic, OpenAI, and Gemini that picks the best model
for every request: using a tiny on-box embedder, not a vibes-based prompt.
🥇 #1 on the RouterArena leaderboard 1 — Acc-Cost Arena 76.09 .
Built by Weave : The #1 engineering intelligence platform,
loved by Robinhood, PostHog, Reducto, and hundreds of others.
Point Claude Code, Codex, Cursor, or your own app at localhost:8080 . The router:
🎯 Routes per request. A cluster scorer derived from
Avengers-Pro 2 picks the right
model from your enabled providers, every turn.
🔌 Speaks everyone's API. Anthropic Messages, OpenAI Chat Completions,
Gemini native. Streaming, tools, vision, the works.
🧠 Knows OSS too. DeepSeek, Kimi, GLM, Qwen, Llama, Mistral via
OpenRouter (or any OpenAI-compatible endpoint).
🔒 BYOK by default. Provider keys stay on your box, encrypted at rest.
📊 Observable. OTLP traces out of the box. See your dashboard in the Weave dashboard ( http://localhost:8080/ui/dashboard ) or drop in Honeycomb, Datadog,
Grafana, whatever.
The fastest way: point Claude Code, Codex, or opencode at the hosted
Weave Router with one command. No clone, no Docker, no Postgres.
npx @workweave/router
That's it. The installer asks which tool (Claude Code, Codex, or opencode),
walks you through scope (user vs. project), grabs a router key, and wires
the right config file. Other flavors:
npx @workweave/router --claude # skip the picker, Claude Code
npx @workweave/router --codex # skip the picker, OpenAI Codex CLI
npx @workweave/router --opencode # skip the picker, opencode
npx @workweave/router --scope project # per-repo, commits settings.json (or .codex/ / opencode.json)
npx @workweave/router --local # self-hosted localhost:8080
npx @workweave/router --base-url https://router.acme.internal
npx @workweave/router@0.1.0 # pin a version
Requires Node ≥ 18 (Claude Code and opencode paths also need jq ). Full
flag reference: install/npm/README.md .
If you want the router (and dashboard) running on your own box:
# 1. Drop a provider key in. OpenRouter is the recommended baseline.
echo " OPENROUTER_API_KEY=sk-or-v1-... " >> .env.local
# 2. Boot Postgres + router on :8080 and seed an rk_ key.
make full-setup
The router is up at http://localhost:8080 , the dashboard at
http://localhost:8080/ui/ (password: admin ), and your rk_... key
prints in the logs.
# Call it like Anthropic
curl -sS http://localhost:8080/v1/messages \
-H " Authorization: Bearer rk_... " \
-d ' {"model":"claude-sonnet-4-5","max_tokens":256,
"messages":[{"role":"user","content":"hi"}]} '
# ...or like OpenAI
curl -sS http://localhost:8080/v1/chat/completions \
-H " Authorization: Bearer rk_... " \
-d ' {"model":"gpt-4o-mini",
"messages":[{"role":"user","content":"hi"}]} '
# Peek at the routing decision without proxying
curl -sS http://localhost:8080/v1/route -H " Authorization: Bearer rk_... " -d ' ... '
Wire it into your tools
Claude Code. Run make install-cc to wire Claude Code at the local
self-hosted router (it's also invoked automatically at the end of
make full-setup ). For the hosted router, use npx @workweave/router
above.
Codex (OpenAI CLI). npx @workweave/router --codex patches
~/.codex/config.toml (or <repo>/.codex/config.toml with --scope project )
with a managed [model_providers.weave] block and sets model_provider = "weave" .
Codex's existing OPENAI_API_KEY flows through to api.openai.com for the
plan-based passthrough; the router key rides in an X-Weave-Router-Key HTTP
header. Re-install and --uninstall --codex rewrite/remove only the managed
block, leaving the rest of your Codex config untouched.
opencode. npx @workweave/router --opencode merges a provider.weave
entry into ~/.config/opencode/opencode.json (or <repo>/opencode.json
with --scope project ). It uses opencode's bundled @ai-sdk/anthropic
provider pointed at the router's /v1 endpoint — the router speaks the
Anthropic Messages API natively, so opencode works unmodified. The router
key and identity headers ride alongside the provider config; re-install
rewrites only the managed block and --uninstall --opencode strips it.
Cursor (early beta, performance may not be the best). Settings →
Models → Override OpenAI Base URL → http://localhost:8080/v1 , paste
rk_... as the API key.
Switching on/off. After installing, npx @workweave/router off --claude
(or --codex / --opencode ) routes that client straight to its provider
again without discarding the router config; on flips it back, and status
reports which way it's pointing. Claude Code also gets /router-off ,
/router-on , and /router-status slash commands. Cursor toggles via the same
Settings → Models override above. See install/README.md .
sk-or-... / sk-ant-... / sk-... = your upstream provider key. Lives in .env.local .
rk_... = your router key. Clients send this as a Bearer token.
Endpoint
Format
POST /v1/messages
Anthropic Messages, routed
POST /v1/chat/completions
OpenAI Chat Completions, routed
POST /v1beta/models/:action
Gemini generateContent , routed
POST /v1/route
Returns the decision, no upstream call
GET /v1/models · POST /v1/messages/count_tokens
Anthropic passthrough
GET /health · GET /validate
liveness + key check
Deeper docs
📐 Configuration reference : every env var,
BYOK encryption, OTel knobs, cluster routing.
🛠️ Contributing : layering rules, hot-reload dev,
migrations, tests, the whole engineering loop.
🏗️ Architecture : package layout, import contracts,
recipes for adding endpoints / providers / strategies.
Token-aware rate limiting (Redis sliding window per installation)
Sub-installations for tenant hierarchies
Speculative dispatch + hedging for tail latency
Lu, Y., Liu, R., Yuan, J., Cui, X., Zhang, S., Liu, H., & Xing, J.
RouterArena: An Open Platform for Comprehensive Comparison of LLM
Routers. arXiv:2510.00202, 2025. https://arxiv.org/abs/2510.00202 ↩
Zhang, Y. et al. Beyond GPT-5: Making LLMs Cheaper and Better via
Performance–Efficiency Optimized Routing (Avengers-Pro).
arXiv:2508.12631, 2025. https://arxiv.org/abs/2508.12631 ↩
Model router for agentic systems. Routes every prompt to the right model in <50ms. Cut costs 40-70% with just an endpoint change.
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
7
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
