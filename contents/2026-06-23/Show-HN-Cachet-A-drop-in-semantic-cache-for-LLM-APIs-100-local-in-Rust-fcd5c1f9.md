---
source: "https://github.com/abhix2112/Cachet"
hn_url: "https://news.ycombinator.com/item?id=48643854"
title: "Show HN: Cachet – A drop-in semantic cache for LLM APIs, 100% local, in Rust"
article_title: "GitHub - abhix2112/Cachet: A transparent, 100%-local semantic cache for LLM APIs — drop-in proxy, one line to integrate, written in Rust · GitHub"
author: "Abhi_2112"
captured_at: "2026-06-23T12:53:39Z"
capture_tool: "hn-digest"
hn_id: 48643854
score: 2
comments: 0
posted_at: "2026-06-23T12:20:09Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Cachet – A drop-in semantic cache for LLM APIs, 100% local, in Rust

- HN: [48643854](https://news.ycombinator.com/item?id=48643854)
- Source: [github.com](https://github.com/abhix2112/Cachet)
- Score: 2
- Comments: 0
- Posted: 2026-06-23T12:20:09Z

## Translation

タイトル: Show HN: Cachet – Rust の 100% ローカルの LLM API 用のドロップイン セマンティック キャッシュ
記事のタイトル: GitHub - abhix2112/Cachet: LLM API 用の透過的で 100% ローカルのセマンティック キャッシュ — ドロップイン プロキシ、統合するのは 1 行、Rust · GitHub
説明: LLM API 用の透過的で 100% ローカルのセマンティック キャッシュ — ドロップイン プロキシ、1 行で統合、Rust で記述 - abhix2112/Cachet

記事本文:
GitHub - abhix2112/Cachet: LLM API 用の透過的で 100% ローカルのセマンティック キャッシュ — ドロップイン プロキシ、1 行で統合、Rust で記述 · GitHub
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
abhix2112
/
カシェット
公共
通知
通知設定を変更するにはサインインする必要があります

s
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
3 コミット 3 コミット アセット アセット スクリプト スクリプト src src .dockerignore .dockerignore .gitignore .gitignore COTRIBUTING.md COTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml Dockerfile Dockerfile LICENSE-APACHE LICENSE-APACHE LICENSE-MIT LICENSE-MIT README.md README.md SECURITY.md SECURITY.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
LLM API 用の透過的で 100% ローカルのセマンティック キャッシュ。何かの前にドロップしてください
アプリを 1 行変更するだけで、通話の繰り返しや言い直しのコストを削減し、視聴できるようになります。
節約はライブで加算されます。
/__cachet/ ダッシュボード: 推定節約額が刻々と増加し、繰り返し/言い換えられた呼び出しがキャッシュから提供されるとストリーミングされる緑色の「ヒット」行が表示されます。
アプリは、同じまたはほぼ同じプロンプトで LLM API を何度もヒットします (再試行、
質問の共有、再実行、わずかに言い換えたクエリ - すべての質問に対して全額を支払います
1つ。通常の答えは、コードに配線するキャッシュ ライブラリか、
重量級のゲートウェイ/マネージド サービスを立ち上げてルーティングします。
カシェットはそのどちらでもない。これは、OpenAI の前に置く単一の Rust バイナリです。
ベース URL を 1 行変更することで、Anthropic 互換の API を作成できます。リクエストは転送されます
透明に;意味的に同等のリクエストが以前に存在した場合、
キャッシュされた応答はローカルに提供され、アップストリーム呼び出しはスキップされます。セマンティック
マッチングは完全にマシン上で実行されます (埋め込み API やベクトル データベースはありません)。
内蔵ダッシュボードには、ヒット率と推定節約額がリアルタイムで表示されます。
# 1. ビルド (単一バイナリ、libc 以降のシステム DEP なし — 純粋な Rust TLS)
カーゴビルド --release
# 2. 設定なしで実行: 127.0.0.1:8080 でリッスンし、https://api.openai.com に転送します。
./ターゲット/r

リリース/カシェット
次に、アプリに 1 行でそれを指定します。
openaiインポートからOpenAI
クライアント = OpenAI (
base_url = "http://localhost:8080/v1" , # ← 唯一の変更
api_key = "sk-..." ,
)
# 他はすべてまったく同じです
curl — ホストを交換するだけです。
# 前: https://api.openai.com/v1/chat/completions
カール http://localhost:8080/v1/chat/completions \
-H " 認可: ベアラー $OPENAI_API_KEY " \
-H " Content-Type: application/json " \
-d ' {"モデル":"gpt-4o","messages":[{"role":"user","content":"フランスの首都はどこですか?"}]} '
それを 2 回送信します。2 番目の応答は X-Cachet で返されます: hit とnever
OpenAIに触れます。次に http://localhost:8080/__cachet/ を開いて動作を確認します。
API キー、ヘッダー、本文、ステータス、ストリーミングはすべて変更されずに通過します。で
キャッシュミスの場合、リクエストは、キャッシュが存在しない場合とまったく同じように、構成されたプロバイダーに送信されます。
カシェット;ヒットすると、ローカル メモリから提供されます。
Rust ツールチェーンは必要ありません。イメージは、最大 47 MB​​ のディストリビューションなしビルドです。 TLS は純粋な Rust です
(CA ルートをコンパイルした Rustls) なので、バイナリにはシステム OpenSSL は必要ありません。
docker build -t cachet 。
# ゼロ構成: https://api.openai.com、ダッシュボード /__cachet/ に転送します。
docker run -p 8080:8080 カシェット
# オーバーライドあり — 例: Anthropic をポイントし、一致のしきい値を緩めます
docker run -p 8080:8080 \
-e CACHET_UPSTREAM=https://api.anthropic.com \
-e CACHET_THRESHOLD=0.85 \
名状
(イメージはコンテナ内で 0.0.0.0 をバインドするため、-p が機能します。1 つおきに
CACHET_* var は -e でオーバーライドできます。)
アプリのキャッシュ (1 つのローカル バイナリ、:8080) アップストリーム
┌─────────┐ 一行 ┌───────────────┐ ┌─────────┐
│base_url = │ b

ase_url │ /v1/* リクエスト │ │ api.openai │
│ localhost │ ────swap────▶ │ │ │ │ .com / any │
│ :8080 │ │ ▼ ① 正確なハッシュ ② ローカル埋め込み │ │ OpenAI-/ │
└───────┘ │ キャッシュ検索 ──────────┐ │ │ Anthropic- │
│ │ ヒットミス│ フォワード │ ────▶ │ 互換性 │
│ ▼ ▼ + ティー │ │ HTTP API │
│ ローカルでクライアントにストリーム配信 │ └───────────┘
│ ($ 保存 ++) & キャプチャ;キャッシュ │
│ きれいな仕上がりについて │
━━━━━━━━━━━━━━━━━━━┘
ダッシュボード: GET /__cachet/
デュアルレイヤーキャッシュ。キャッシュ可能なすべてのリクエストはモデルと正規化によってキー化されます
プロンプト。まず正確なレイヤー (プロンプトの O(1) ハッシュ) — 無料で即時です。で
ミスの場合、セマンティック レイヤーがプロンプトを埋め込み、最も近い保存されたエントリを検索します。
コサイン類似度。設定可能なしきい値をクリアした場合にのみそれを返します。
ローカル語彙埋め込みツール。デフォルトのエンベッダーはニューラル モデルではありません。それは、
単語レベルおよび文字トリグラム特徴 (ハッシュ
トリック)、ストップワードを削除して内容ワードを優先させます。高速、無料、プライベート、
そして依存性がありません — そしてそれには実際の上限があります: 共有する言い換えをキャッチします
語彙または単語の形 (「フランスの首都はどこですか?」 ≈ 「フランスの首都は何ですか?」
首都?」) が、重複ゼロの言い換えではありません (「どの都市に都市があり、
エリゼ宮？」）。 Embedder トレイトは、ニューラル/API の単一のスワップイン ポイントです。
よりシャープなマッチングが必要な場合は、embedder を使用します。
ストリーミング T シャツ + リプレイ。ストリーミング (SSE) 応答もキャッシュされます。見逃した場合は、
応答はチャンクごとにクライアントにストリーミングされ、

d 並行してキャプチャ。それは
ストリームが正常に終了した場合、つまりドロップまたはエラーが発生した場合にのみキャッシュに書き込まれます。
ストリームは決してキャッシュされません。ヒットすると、保存された SSE が適切にフレーム化されたイベントとして再生されます。
ストリーム。
ライブ貯蓄ダッシュボード。 /__cachet/ にある自己完結型ページ (外部/CDN なし)
リクエスト）SSE 経由でメトリクスをストリーミングします。実行中の推定 $saved カウンター、ヒット
レート、保存されたトークン、色分けされたライブ リクエスト フィード。
すべての設定は環境変数です。すべてオプションです。
Cache がユースケースに適合するかどうかは、正直に言ってください。
デフォルトのエンベッダーはニューラルではなく字句です。高速、無料、そして完全にローカルです。
しかし、共有された単語や文字の形に一致します。言い換えるのに最適です。
語彙の重複があり、言い換えが存在しないため盲目です。真のセマンティックマッチングが必要な場合は、
Embedder トレイトを介してニューラル エンベッダーに交換します (ロードマップを参照)。
「節約された金額」の数値は概算であり、請求額ではありません。トークンは次のように近似されます。
chars ÷ 4 で計算され、おおよその公開定価の編集可能な表に基づいて価格が設定されます。
提供された回答のテキストのみ (JSON/SSE フレーム化ではない) をカウントするため、エラーになります。
保守的ですが、これは推定値です。プランに合わせて価格を編集します。
キャッシュは、決定論的なワークロードに最も役立ちます。共有Q&A、ドキュメント/サポートボット、
eval と再試行には大きなメリットがあります。高度にパーソナライズされた、ユーザーごと、または
めったに繰り返されない高温のプロンプトはほとんど効果がありません。
キャッシュはメモリ内にあります。 TTL と最大エントリの上限によって制限されており、空です
再起動時 (ディスク上の永続性はまだありません - ロードマップを参照)。
ミスした場合は、設定したプロバイダーにプロンプトが表示されます。
カシェなし。 「100% ローカル」とは、キャッシュと埋め込みを指します。Cachet は何も追加しません。
サードパーティの埋め込み API またはベクター データベース、およびキャッシュ ヒットは完全に提供されます
ローカルの記憶から。
まだマルチテンには達していない

アリの玄関口。認証、レート制限、またはキーごとの割り当てはありません。
API 管理プラットフォームではなく、ドロップイン キャッシュです。
Cachet は API キーをアップストリームに転送し、応答をキャッシュするため、いくつかのプロパティが重要です。
Authorization ヘッダーがログに記録されたり、保存されたり、ダッシュボードに表示されたりすることはありません。の
キャッシュはモデル + プロンプト + フォーマットによってキー設定され、キャッシュのすべての呼び出し元で共有されます。
インスタンス (信頼境界ごとに 1 つ実行)。ダッシュボードには認証がありません (公開しないでください)
ポートは公開されています）。完全な信頼モデル — 呼び出し元間のキャッシュ共有を含む
資格情報の処理とネットワークへの露出 - SECURITY.md にあります。お願いします
独自のアプリの前以外の場所に Cachet をデプロイする前に読んでください。
貢献は歓迎です — 以下から始めるのが良いでしょう:
既存の Embedder トレイトの背後にあるオプションのニューラル/API エンベッダーにより、より鮮明になります。
セマンティック マッチング (重複ゼロの言い換え)。
再起動後に存続する永続的/ディスク上のキャッシュ。
ANN インデックスは、大規模なキャッシュの線形類似性スキャンを置き換えます。
より多くのプロバイダー/応答形状とより豊富なトークン アカウンティング。
MIT または Apache-2.0 のいずれかに基づいてライセンスが付与されている
オプション。
貢献は大歓迎です。最初に重要な変更について議論するために問題を開いてください。によって
寄稿すると、あなたの作品が同じ条件の下でデュアルライセンスされることに同意したことになります。
LLM API 用の透過的で 100% ローカルのセマンティック キャッシュ — ドロップイン プロキシ、1 行で統合、Rust で記述
Apache-2.0、MITライセンスが見つかりました
ライセンスが見つかりました
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A transparent, 100%-local semantic cache for LLM APIs — drop-in proxy, one line to integrate, written in Rust - abhix2112/Cachet

GitHub - abhix2112/Cachet: A transparent, 100%-local semantic cache for LLM APIs — drop-in proxy, one line to integrate, written in Rust · GitHub
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
abhix2112
/
Cachet
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits assets assets scripts scripts src src .dockerignore .dockerignore .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml Dockerfile Dockerfile LICENSE-APACHE LICENSE-APACHE LICENSE-MIT LICENSE-MIT README.md README.md SECURITY.md SECURITY.md View all files Repository files navigation
A transparent, 100%-local semantic cache for LLM APIs. Drop it in front of any
app with a one-line change, cut the cost of repeated and rephrased calls, and watch
the savings add up live.
The /__cachet/ dashboard: estimated $ saved ticking up and green “hit” rows streaming in as repeated/rephrased calls are served from cache.
Apps hit LLM APIs with the same and nearly-the-same prompts over and over — retries,
shared questions, re-runs, slightly reworded queries — and pay full price for every
one. The usual answers are either a caching library you wire into your code, or a
heavyweight gateway/managed service you stand up and route through.
Cachet is neither. It's a single Rust binary you put in front of any OpenAI- or
Anthropic-compatible API by changing one line — the base URL. Requests are forwarded
transparently; when a semantically equivalent request has been seen before, the
cached answer is served locally and the upstream call is skipped. The semantic
matching runs entirely on your machine (no embedding API, no vector database), and a
built-in dashboard shows the hit rate and estimated dollars saved in real time.
# 1. Build (single binary, no system deps beyond libc — pure-Rust TLS)
cargo build --release
# 2. Run with zero config: listens on 127.0.0.1:8080, forwards to https://api.openai.com
./target/release/cachet
Now point your app at it with one line :
from openai import OpenAI
client = OpenAI (
base_url = "http://localhost:8080/v1" , # ← the only change
api_key = "sk-..." ,
)
# everything else is exactly the same
curl — just swap the host:
# before: https://api.openai.com/v1/chat/completions
curl http://localhost:8080/v1/chat/completions \
-H " Authorization: Bearer $OPENAI_API_KEY " \
-H " Content-Type: application/json " \
-d ' {"model":"gpt-4o","messages":[{"role":"user","content":"What is the capital of France?"}]} '
Send that twice — the second response comes back with X-Cachet: hit and never
touches OpenAI. Then open http://localhost:8080/__cachet/ to watch it work.
Your API key, headers, body, status, and streaming all pass through unchanged. On a
cache miss the request goes to your configured provider exactly as it would without
Cachet; on a hit it's served from local memory.
No Rust toolchain needed. The image is a ~47 MB distroless build; TLS is pure-Rust
(rustls with CA roots compiled in), so the binary needs no system OpenSSL.
docker build -t cachet .
# zero config: forwards to https://api.openai.com, dashboard at /__cachet/
docker run -p 8080:8080 cachet
# with overrides — e.g. point at Anthropic and loosen the match threshold
docker run -p 8080:8080 \
-e CACHET_UPSTREAM=https://api.anthropic.com \
-e CACHET_THRESHOLD=0.85 \
cachet
(The image binds 0.0.0.0 inside the container so -p works; every other
CACHET_* var is overridable with -e .)
your app Cachet (one local binary, :8080) upstream
┌────────────┐ one-line ┌──────────────────────────────────────────┐ ┌───────────────┐
│ base_url = │ base_url │ /v1/* request │ │ api.openai │
│ localhost │ ───swap────▶ │ │ │ │ .com / any │
│ :8080 │ │ ▼ ① exact hash ② local embedding │ │ OpenAI-/ │
└────────────┘ │ cache lookup ──────────────┐ │ │ Anthropic- │
│ │ hit miss│ forward │ ────▶ │ compatible │
│ ▼ ▼ + tee │ │ HTTP API │
│ served locally stream to client │ └───────────────┘
│ ($ saved ++) & capture; cache │
│ on clean finish │
└──────────────────────────────────────────┘
dashboard: GET /__cachet/
Dual-layer cache. Every cacheable request is keyed by model + normalized
prompt. First an exact layer (an O(1) hash of the prompt) — free, instant. On a
miss, a semantic layer embeds the prompt and finds the nearest stored entry by
cosine similarity, returning it only if it clears a configurable threshold.
Local lexical embedder. The default embedder is not a neural model. It turns a
prompt into a vector from word-level and character-trigram features (the hashing
trick), with stopword removal so content words dominate. It's fast, free, private,
and dependency-free — and it has a real ceiling: it catches rephrasings that share
vocabulary or word-shape ("What is the capital of France?" ≈ "What's France's
capital city?") but not zero-overlap paraphrases ("Which city houses the
Élysée Palace?"). The Embedder trait is a single swap-in point for a neural/API
embedder when you want sharper matching.
Streaming tee + replay. Streaming (SSE) responses are cached too. On a miss the
response is streamed to your client chunk-by-chunk and captured in parallel; it's
written to the cache only if the stream finishes cleanly — a dropped or errored
stream is never cached. On a hit, the stored SSE is replayed as a well-framed event
stream.
Live savings dashboard. A self-contained page at /__cachet/ (no external/CDN
requests) streams metrics over SSE: a running estimated $ saved counter, hit
rate, tokens saved, and a color-coded live request feed.
All configuration is environment variables; all are optional.
Honesty up front, because it determines whether Cachet fits your use case:
The default embedder is lexical, not neural. It's fast, free, and fully local,
but it matches on shared words and character shape — great for rephrasings with
lexical overlap, blind to paraphrases with none. If you need true semantic matching,
swap in a neural embedder via the Embedder trait (see Roadmap).
The "$ saved" number is an estimate, not a bill. Tokens are approximated as
chars ÷ 4 and priced against an editable table of approximate public list prices.
We count only the served answer's text (not JSON/SSE framing), so it errs
conservative — but it is an estimate. Edit the prices to match your plan.
Caching helps deterministic-ish workloads most. Shared Q&A, docs/support bots,
evals, and retries benefit a lot. Highly personalized, per-user, or
high-temperature prompts that are rarely repeated benefit little.
The cache is in-memory. It's bounded by TTL and a max-entry cap, and it's empty
on restart (no on-disk persistence yet — see Roadmap).
On a miss, your prompt goes to your configured provider — exactly as it would
without Cachet. "100% local" refers to the caching and embedding: Cachet adds no
third-party embedding API or vector database, and cache hits are served entirely
from local memory.
Not yet a multi-tenant gateway. No auth, rate limiting, or per-key quotas — it's
a drop-in cache, not an API management platform.
Cachet forwards your API key upstream and caches responses, so a few properties matter:
the Authorization header is never logged, stored, or shown in the dashboard; the
cache is keyed by model + prompt + format and is shared across all callers of an
instance (run one per trust boundary); and the dashboard has no auth (don't expose
the port publicly). The full trust model — including cross-caller cache sharing,
credential handling, and network exposure — is in SECURITY.md . Please
read it before deploying Cachet anywhere other than in front of your own app.
Contributions welcome — these are good places to start:
Optional neural/API embedder behind the existing Embedder trait, for sharper
semantic matching (zero-overlap paraphrases).
Persistent / on-disk cache that survives restarts.
ANN index to replace the linear similarity scan for large caches.
More providers / response shapes and richer token accounting.
Licensed under either of MIT or Apache-2.0 at your
option.
Contributions are welcome — open an issue to discuss substantial changes first. By
contributing you agree your work is dual-licensed under the same terms.
A transparent, 100%-local semantic cache for LLM APIs — drop-in proxy, one line to integrate, written in Rust
Apache-2.0, MIT licenses found
Licenses found
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
