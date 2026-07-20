---
source: "https://github.com/llmrelay/relay"
hn_url: "https://news.ycombinator.com/item?id=48986115"
title: "Show HN: Relay – a self-hosted LLM gateway with eval-gated routing"
article_title: "GitHub - llmrelay/relay: Self-hosted LLM gateway + model router. One static binary, OpenAI & Anthropic dialects in, any provider out, zero telemetry. · GitHub"
author: "olopadef"
captured_at: "2026-07-20T23:47:43Z"
capture_tool: "hn-digest"
hn_id: 48986115
score: 1
comments: 0
posted_at: "2026-07-20T23:18:00Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Relay – a self-hosted LLM gateway with eval-gated routing

- HN: [48986115](https://news.ycombinator.com/item?id=48986115)
- Source: [github.com](https://github.com/llmrelay/relay)
- Score: 1
- Comments: 0
- Posted: 2026-07-20T23:18:00Z

## Translation

タイトル: HN の表示: リレー – 評価ゲート型ルーティングを備えたセルフホスト型 LLM ゲートウェイ
記事のタイトル: GitHub - llmrelay/relay: セルフホステッド LLM ゲートウェイ + モデル ルーター。 1 つの静的バイナリ、OpenAI および Anthropic 方言が入力され、任意のプロバイダーが出力され、テレメトリはゼロです。 · GitHub
説明: セルフホステッド LLM ゲートウェイ + モデル ルーター。 1 つの静的バイナリ、OpenAI および Anthropic 方言が入力され、任意のプロバイダーが出力され、テレメトリはゼロです。 - llmrelay/リレー

記事本文:
GitHub - llmrelay/relay: セルフホステッド LLM ゲートウェイ + モデル ルーター。 1 つの静的バイナリ、OpenAI および Anthropic 方言が入力され、任意のプロバイダーが出力され、テレメトリはゼロです。 · GitHub
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
コードの品質 マージ時に品質を強制する
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
読み込み中にエラーが発生しました。 PL

このページを簡単にリロードしてください。
llmリレー
/
リレー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
56 コミット 56 コミット .claude/ スキル/ 検証 .claude/ スキル/ 検証 .github .github アセット アセット cmd/ リレー cmd/ リレー docs docs 例 例 future/ package-installers future/ package-installers 内部 内部ツール ツール .gitattributes .gitattributes .gitignore .gitignore .golangci.yml .golangci.yml .goreleaser.yaml .goreleaser.yaml CHANGELOG.md CHANGELOG.md COTRIBUTING.md COTRIBUTING.md DESIGN.md DESIGN.md Dockerfile Dockerfile LAUNCH.md LAUNCH.md ライセンス ライセンス通知 NOTICE README.md README.md SECURITY.md SECURITY.md go.mod go.mod go.sum go.sum すべてのファイルを表示 リポジトリ ファイルのナビゲーション
セルフホステッド LLM ゲートウェイ + モデル ルーター。両方を話す 1 つの静的 Go バイナリ
OpenAI および Anthropic API 方言のインバウンド、任意のプロバイダー間のルート
アウトバウンド (独自のキーの持ち込み)、静的エイリアスから外部へのルーティング ポリシーを使用
学習済みスマート層 — 同梱の評価ハーネスによってゲートされます。ゼロ
テレメトリ — オプトイン ping さえありません。 Apache-2.0。
任意の OpenAI SDK ─┐ ┌─ OpenAI / Anthropic / Gemini
クロード コード ────┤→ リレー (1 バイナリ、:4000) → §─ Groq / DeepSeek / Mistral / xAI / …
プレーンカール ───┘ ルート、フェイルオーバー、ログ、キャッシュ └─ ローカルの Ollama
60 秒のクイックスタート
# 1. インストール (またはリリース バイナリをダウンロードします。テスト済みの Dockerfile が含まれているため、自分でビルドすることができます)
醸造インストール llmrelay/tap/relay
# 2. すでに持っているキー、設定なし
import GEMINI_API_KEY=... # および/または OPENAI_API_KEY、ANTHROPIC_API_KEY、GROQ_API_KEY、...
リレーサーブ
#3. 何かを向ける
カール http://localhost:4000/v1/chat/co

mpletions -H " Content-Type: application/json " \
-d ' {"モデル":"gemini/gemini-3.1-flash-lite","メッセージ":[{"役割":"ユーザー","コンテンツ":"hello"}]} '
import ANTHROPIC_BASE_URL=http://localhost:4000 # クロード コードがリレー経由で実行されるようになりました
リレー初期化は完全なrelay.yamlのスキャフォールディングを行います — すべてのリレーの入手先を含む
プロバイダーのキー。クロード コード、カーソル、SDK の例を参照してください。
スニペット、および機密ローカル/バルク安価のエイリアス レシピ。
なぜリレーするのか（そして正直な代替案）
リレー
LiteLLM
オープンルーター
ルートLLM
それは何ですか
セルフホステッドゲートウェイ + ルーター
Python プロキシ/SDK ゲートウェイ
ホスト型アグリゲーター API
リサーチルーティングフレームワーク
デプロイ
1 つの静的バイナリ/ディストロレス イメージ
Python サービス + デプス
何もない（彼らの雲）
Pythonライブラリ
インバウンド方言
OpenAI と Anthropic、完全な相互翻訳
OpenAI (+ パススルー)
OpenAI対応
該当なし
学習したルーティング
段階的、ローカル、ログでトレーニング、評価ゲート
手動ルーティング戦略
彼らのルーティング
従来技術¹
評価用ハーネスが同梱
はい (リレー評価、コミットされたセット、ライブ判定)
いいえ
いいえ
オフライン研究ハーネス
プロンプトは第三者を経由します
なし (自己ホスト型、リモート組み込みルーティングには明示的なオプトインが必要)
自己ホスト型
はい、それが製品です
該当なし
テレメトリー
何も、決して
ドキュメントごとに何もありません
ホスト型サービス
該当なし
生態系の広がり
4 ネイティブ + 12 プリセット、アダプター ガイド
最も広範なプロバイダー マトリクス、チーム/予算/仮想キー
非常に広い
該当なし
ライセンス/ランタイム
Apache-2.0、Go
混合 (OSS + エンタープライズ)、Python
独自のサービス
Apache-2.0、Python
¹ リレーの Tier-2 KNN は RouteLLM と同じファミリーです
(Ong et al.)、独自のトラフィック上でローカルに実行します。グラフルーター
( GraphRouter 、 ICLR 2025) はロードマップです。
LiteLLM と OpenRouter は、異なるトレードオフを持つ優れた製品です。
あなたにとって重要な行。
インバウンド
ステータス
OpenAI チャット補完 ( /v1/chat/completions )

フル、含むストリーミング、ツール、ビジョン
OpenAI レスポンス API ( /v1/responses )
まだ — 追跡済み v1.1 ファストフォロー
人間的メッセージ ( /v1/messages 、 count_tokens )
フル、含むストリーミング、ツール
埋め込み ( /v1/embeddings )
OpenAI の方言。埋め込み API を持たないプロバイダーは正直に 404 に応答します
/v1/models 、 /metrics (Prometheus)、 /dashboard 、 /v1/フィードバック
はい
パフォーマンス
アップストリームのループバック モックに対して測定されたゲートウェイのオーバーヘッド ( go test -run TestOverheadBudget ./internal/server/ — 予算はハード CI ゲートです。
DESIGN.md §11 の方法論。 Windows 11 / Go 1.25、2026-07-18):
プロバイダーの遅延がエンドツーエンド時間の大半を占めます。リレーの仕事は目に見えないようにすることです。
4 つのポリシー (フォールバック、最安、最速、
重み付け)、さらにすべてのチェーンの下にある信頼性: ジッターのある再試行
バックオフ、レート制限クールダウンを備えた API キー プール、サーキット ブレーカー、
ファーストトークン前のストリーミングフェイルオーバー。
スマート ルーティング — デフォルトではオフ、それが機能です
リレーは難易度ベースのスマートルーターを出荷します（トラフィックが簡単→安価なチェーン、
ハード → フロンティア チェーン) 評価ハーネスとハーネス自体
信頼する前に、トラフィックのベースラインを上回る必要があるという判断です。
したがって、スマート ルーティングはデフォルトでオフになっています。
リレー評価 # ドライラン: コミットされたセット、候補
リレー eval --live-judge --dry-run # 実際の完了数、ジャッジスコア。プリントが最初に消費される
ホールドアウト セット v2 でハーネスが測定したもの、ライブ判定 - 本物
各ルーティングモデルからの完成度、claude-opus-4-8 によってスコア付けされた品質
(2026-07-20; 49 プロンプト、各ポリシーの有効な N = 49、完了の上限は
700 個の出力トークン。 Gemini の無料枠のため、ミッドバンドは claude-sonnet-5 です。
gemini-3.5-flash の上限は 1 日あたり 20 リクエストです。裁判官の返答147件中6件が埋葬された
解説内の評決番号

から決定論的に再解析されました。
スコア 0 ではなく生の返信を保存します — 内部の修正ログ
それぞれを判定する JSON レコード):
数値は四捨五入して表示しています。ゲートの決定では、丸められていない値が使用されます。
コミットされた判定アセット (正確に -0.0200 のデルタが合格します。ルールは次のとおりです)
包括的;層 1 の四捨五入されていないデルタは -0.020408…、厳密には -1/49 です。
その下）。 CI テストでは、このテーブルが完全な精度でアセットと一致することがアサートされます。
この 49 回のプロンプト、700 個の出力トークン、単一審査員の評価では、静的で安価な
ベースラインは、設定されたフロンティア ベースラインよりも高い判定スコアを受け取りました。
その結果は、テストされたモデル、プロンプト、トークンキャップ、およびジャッジに固有です。
これは一般的なモデルの品質を主張するものではありません。それが確立するのはバーです
リレーは次のことを守ります: ルーティングの賢さは強力な攻撃に対して負けないようにする必要があります
トラフィックに基づいて測定された、愚かなベースライン。 Tier 1 はホールドアウトでそのバーに失敗しました
データ。レベル 2 は許容範囲内で合格しましたが、v0.1.0 ではオプトインのままです。賢くない
この層はデフォルトでオンになっています。
(履歴: v1 セットでは、Tier 1 は -0.017 で合成に「合格」しました — サンプル内
その重みは v1 で調整されているため、お世辞です。 v2 合成ラベルでは、
-0.071で失敗しました。 v1 はキャリブレーション セットです。 v2ライブジャッジはスタンディング
評決。完全なテーブル:assets/eval/ 。)
これを有効にするのは 1 つの構成ブロックです。明示的な層があり、何もないためです。
サイレントデフォルトでトラフィックをルーティングします。
プロバイダー:
ジェミニ :
api_key : ["${GEMINI_API_KEY}"]
人間性 :
api_key : ["${ANTHROPIC_API_KEY}"]
オラマ :
タイプ : オラマ
ルーティング:
デフォルト: スマート
賢い：
簡単：gemini/gemini-3.1-flash-lite
ハード：人間/クロード・ソネット-5
embeddings : ollama/nomic-embed-text # knn 層 (文書化されたパス) を選択します
# 層: 語彙 # 実験的: 保持されたゲートに失敗しました
Tier 2 (KNN) が取得

リレートレイン経由でトラフィックを改善 — 暗黙的に
シグナル、オプションのリプレイ + ジャッジ (常に支出を見積もり、最初に質問します)、および
POST /v1/フィードバック スコアにより、参照セットが増加します。それから自分で測定してください
リレー eval --refs ~/.relay/smart_refs.json とのクロスオーバー。すべてのスマート
決定はその証拠をログに記録します ( knn: 5 neighbors [seed-math-03 d=0.75 sim=0.95; …] ) — 「モデルがそう感じた」というルーティング理由は受け入れられません。リモコン
embedder には、allow_remote_embeddings: true が必要です。決してサイレントにルーティングしない
プロンプトをどこにでも送信します。
独自のプロンプトでモデルを比較する
リレー比較 --models gemini/gemini-3.1-flash-lite、anthropic/claude-haiku-4-5、groq/llama-3.3-70b \
--html Compare.html " バックエンド エンジニアに CRDT を 5 つの文で説明します。"
1 つのテーブル (および共有可能な HTML レポート): 出力、コスト、レイテンシー、TTFT 側
トラフィックを処理するのと同じアダプターを介して並行して使用できます。
docs/models-landscape.md は日付が変わっています
「どのモデルを何に使うか」の相棒。
すべてのリクエストは、プライバシー層 (デフォルトではオフ) でローカル SQLite にログされます。
メタデータのみ。埋め込みにはクエリ ベクトルが格納されますが、テキストは格納されません。いっぱいです
明示的な選択）。 /dashboard には、日別の支出、レイテンシのパーセンタイル、および
人間が判読できる理由を伴う最近のルーティング決定。 /metrics は
プロメテウス。価格設定レジストリにないモデルは価格が設定されていないものとして表示されます
— サイレント 0 ドルは決してありません。
デフォルトではループバック。インバウンド API キーのない非ループバック バインドは拒否されます
を開始します。テレメトリーはありません。 SECURITY.md を参照してください。
COTRIBUTING.md — アダプター作成ガイドがそこにあります。
OpenAI 互換プロバイダーは 1 エントリのプリセットです。デザインの変更が得られます
DESIGN.md §0 で最初に議論されました。そうすることでドキュメントは真実であり続けるのです。
セルフホステッド LLM ゲートウェイ + モデル ルーター。 1 つの静的バイナリ、OpenAI および Anthropic 方言が入力され、任意のプロバイダーが出力され、テレメトリはゼロです。
Apache-2.0ライセンス
貢献する
そこで

読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
0
フォーク
レポートリポジトリ
リリース
2
v0.1.0
最新の
2026 年 7 月 20 日
+ 1 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Self-hosted LLM gateway + model router. One static binary, OpenAI & Anthropic dialects in, any provider out, zero telemetry. - llmrelay/relay

GitHub - llmrelay/relay: Self-hosted LLM gateway + model router. One static binary, OpenAI & Anthropic dialects in, any provider out, zero telemetry. · GitHub
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
Code Quality Enforce quality at merge
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
llmrelay
/
relay
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
56 Commits 56 Commits .claude/ skills/ verify .claude/ skills/ verify .github .github assets assets cmd/ relay cmd/ relay docs docs examples examples future/ package-installers future/ package-installers internal internal tools tools .gitattributes .gitattributes .gitignore .gitignore .golangci.yml .golangci.yml .goreleaser.yaml .goreleaser.yaml CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md DESIGN.md DESIGN.md Dockerfile Dockerfile LAUNCH.md LAUNCH.md LICENSE LICENSE NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md go.mod go.mod go.sum go.sum View all files Repository files navigation
Self-hosted LLM gateway + model router. One static Go binary that speaks both
the OpenAI and Anthropic API dialects inbound , routes across any provider
outbound (bring your own keys), with routing policies from static aliases to a
learned smart tier — gated by an eval harness that ships in the box. Zero
telemetry — not even opt-in pings. Apache-2.0.
any OpenAI SDK ─┐ ┌─ OpenAI / Anthropic / Gemini
Claude Code ────┤→ relay (one binary, :4000) → ├─ Groq / DeepSeek / Mistral / xAI / …
plain curl ─────┘ route · failover · log · cache └─ your local Ollama
60-second quickstart
# 1. install (or download a release binary; a tested Dockerfile is included to build yourself)
brew install llmrelay/tap/relay
# 2. keys you already have, zero config
export GEMINI_API_KEY=... # and/or OPENAI_API_KEY, ANTHROPIC_API_KEY, GROQ_API_KEY, ...
relay serve
# 3. point anything at it
curl http://localhost:4000/v1/chat/completions -H " Content-Type: application/json " \
-d ' {"model":"gemini/gemini-3.1-flash-lite","messages":[{"role":"user","content":"hello"}]} '
export ANTHROPIC_BASE_URL=http://localhost:4000 # Claude Code now runs through relay
relay init scaffolds a full relay.yaml — including where to get every
provider's key. See examples/ for Claude Code, Cursor, SDK
snippets, and the sensitive-local / bulk-cheap alias recipe.
Why relay (and honest alternatives)
relay
LiteLLM
OpenRouter
RouteLLM
What it is
self-hosted gateway + router
Python proxy/SDK gateway
hosted aggregator API
research routing framework
Deploy
one static binary / distroless image
Python service + deps
nothing (their cloud)
Python library
Inbound dialects
OpenAI and Anthropic, full cross-translation
OpenAI (+ passthroughs)
OpenAI-compatible
n/a
Learned routing
tiered, local, trained on your logs, eval-gated
manual routing strategies
their routing
the prior art¹
Eval harness in the box
yes ( relay eval , committed sets, live-judge)
no
no
offline research harness
Your prompts transit a third party
never (self-hosted; remote-embedder routing requires explicit opt-in)
self-hosted
yes — that's the product
n/a
Telemetry
none, ever
none per their docs
hosted service
n/a
Ecosystem breadth
4 native + 12 presets, adapter guide
broadest provider matrix, teams/budgets/virtual keys
very broad
n/a
License / runtime
Apache-2.0, Go
mixed (OSS + enterprise), Python
proprietary service
Apache-2.0, Python
¹ relay's tier-2 KNN is the same family as RouteLLM
(Ong et al.), run locally over your own traffic; graph routers
( GraphRouter , ICLR 2025) are roadmap.
LiteLLM and OpenRouter are good products with different trade-offs — pick the
row that matters to you.
Inbound
Status
OpenAI Chat Completions ( /v1/chat/completions )
full, incl. streaming, tools, vision
OpenAI Responses API ( /v1/responses )
not yet — tracked v1.1 fast-follow
Anthropic Messages ( /v1/messages , count_tokens )
full, incl. streaming, tools
Embeddings ( /v1/embeddings )
OpenAI dialect; providers without an embeddings API answer an honest 404
/v1/models , /metrics (Prometheus), /dashboard , /v1/feedback
yes
Performance
Gateway overhead measured against a loopback mock upstream ( go test -run TestOverheadBudget ./internal/server/ — the budget is a hard CI gate,
methodology in DESIGN.md §11 ; Windows 11 / Go 1.25, 2026-07-18):
Provider latency dominates end-to-end time; relay's job is to stay invisible.
Static routes and aliases with four policies (fallback / cheapest / fastest /
weighted), plus reliability underneath every chain: retries with jittered
backoff, API-key pools with rate-limit cooldowns, circuit breakers, and
pre-first-token streaming failover.
Smart routing — off by default, and that's the feature
relay ships a difficulty-based smart router (easy traffic → your cheap chain,
hard → your frontier chain) with an eval harness — and the harness's own
verdict is that you should beat our baseline on your traffic before trusting
it , so smart routing is off by default:
relay eval # dry-run: the committed sets, your candidates
relay eval --live-judge --dry-run # real completions, judge-scored; prints spend first
What our harness measured on the held-out set v2, live-judged — real
completions from each routed model, quality scored by claude-opus-4-8
(2026-07-20; 49 prompts, valid N = 49 for every policy; completions capped at
700 output tokens; mid band is claude-sonnet-5 because Gemini's free tier
caps gemini-3.5-flash at 20 requests/day; 6 of 147 judge replies buried the
verdict number in commentary and were re-parsed deterministically from the
preserved raw replies rather than scored 0 — the corrections log inside the
verdict JSON records each one):
Values are shown rounded. Gate decisions use the unrounded values from the
committed verdict asset (a delta of exactly −0.0200 would pass — the rule is
inclusive; tier 1's unrounded delta is −0.020408…, exactly −1/49, strictly
below it). A CI test asserts this table matches the asset at full precision.
On this 49-prompt, 700-output-token, single-judge evaluation, the static-cheap
baseline received a higher judged score than the configured frontier baseline.
That result is specific to the tested models, prompts, token cap, and judge —
it is not a general model-quality claim. What it does establish is the bar
relay holds itself to: routing cleverness must earn its keep against strong
dumb baselines, measured on your traffic . Tier 1 failed that bar on held-out
data; tier 2 passed within tolerance but stays opt-in for v0.1.0; no smart
tier is on by default.
(History: on the v1 set tier 1 had "passed" synthetic at −0.017 — in-sample
flattery, since its weights were calibrated on v1. On v2 synthetic labels it
failed at −0.071. v1 is the calibration set; v2 live-judged is the standing
verdict. Full tables: assets/eval/ .)
Enabling it is one config block — with an explicit tier, because nothing
routes your traffic by silent default:
providers :
gemini :
api_key : ["${GEMINI_API_KEY}"]
anthropic :
api_key : ["${ANTHROPIC_API_KEY}"]
ollama :
type : ollama
routing :
default : smart
smart :
easy : gemini/gemini-3.1-flash-lite
hard : anthropic/claude-sonnet-5
embeddings : ollama/nomic-embed-text # selects the knn tier (documented path)
# tier: lexical # experimental: failed the held-out gate
Tier 2 (KNN) gets better on YOUR traffic via relay train — implicit
signals, optional replay+judge (always estimates spend and asks first), and
POST /v1/feedback scores grow its reference set; then measure your own
crossover with relay eval --refs ~/.relay/smart_refs.json . Every smart
decision logs its evidence ( knn: 5 neighbors [seed-math-03 d=0.75 sim=0.95; …] ) — "the model felt like it" is not an accepted routing reason. A remote
embedder requires allow_remote_embeddings: true ; routing never silently
ships prompts anywhere.
Compare models on your own prompts
relay compare --models gemini/gemini-3.1-flash-lite,anthropic/claude-haiku-4-5,groq/llama-3.3-70b \
--html compare.html " Explain CRDTs to a backend engineer in five sentences. "
One table (and a shareable HTML report): output, cost, latency, and TTFT side
by side, through the same adapters that serve your traffic.
docs/models-landscape.md is the dated
"which model for what" companion.
Every request logs to local SQLite with privacy tiers ( off by default —
metadata only; embeddings stores query vectors, never text; full is an
explicit choice). /dashboard shows spend by day, latency percentiles, and
recent routing decisions with human-readable reasons. /metrics is
Prometheus. Models missing from the pricing registry surface as unpriced
— never a silent $0.
Loopback by default; a non-loopback bind without inbound API keys refuses
to start . No telemetry. See SECURITY.md .
CONTRIBUTING.md — the adapter-authoring guide is there;
an OpenAI-compatible provider is a one-entry preset. Design changes get
argued in DESIGN.md §0 first; it's how the doc stays true.
Self-hosted LLM gateway + model router. One static binary, OpenAI & Anthropic dialects in, any provider out, zero telemetry.
Apache-2.0 license
Contributing
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
2
v0.1.0
Latest
Jul 20, 2026
+ 1 release
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
