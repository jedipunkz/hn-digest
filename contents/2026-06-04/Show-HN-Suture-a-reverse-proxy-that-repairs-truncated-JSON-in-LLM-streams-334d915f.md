---
source: "https://github.com/tensorhq/suture-stream-repair"
hn_url: "https://news.ycombinator.com/item?id=48393431"
title: "Show HN: Suture – a reverse proxy that repairs truncated JSON in LLM streams"
article_title: "GitHub - tensorhq/suture-stream-repair: Ultra-low-latency reverse proxy that repairs truncated & malformed JSON in LLM streaming responses (OpenAI, Anthropic, Vertex AI, Bedrock) — fixes JSONDecodeError / serde_json EOF on truncated tool calls. · GitHub"
author: "nabucodonosor"
captured_at: "2026-06-04T04:35:07Z"
capture_tool: "hn-digest"
hn_id: 48393431
score: 1
comments: 0
posted_at: "2026-06-04T03:36:32Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Suture – a reverse proxy that repairs truncated JSON in LLM streams

- HN: [48393431](https://news.ycombinator.com/item?id=48393431)
- Source: [github.com](https://github.com/tensorhq/suture-stream-repair)
- Score: 1
- Comments: 0
- Posted: 2026-06-04T03:36:32Z

## Translation

タイトル: HN を表示: Suture – LLM ストリーム内の切り捨てられた JSON を修復するリバース プロキシ
記事のタイトル: GitHub - tensorhq/suture-stream-repair: LLM ストリーミング レスポンス (OpenAI、Anthropic、Vertex AI、Bedrock) 内の切り捨てられた不正な JSON を修復する超低遅延リバース プロキシ — 切り捨てられたツール呼び出しでの JSONDecodeError / serde_json EOF を修正します。 · GitHub
説明: LLM ストリーミング応答 (OpenAI、Anthropic、Vertex AI、Bedrock) 内の切り捨てられた不正な JSON を修復する超低遅延リバース プロキシ — 切り捨てられたツール呼び出しでの JSONDecodeError / serde_json EOF を修正します。 - tensorhq/縫合糸ストリーム修復

記事本文:
GitHub - tensorhq/suture-stream-repair: LLM ストリーミング レスポンス (OpenAI、Anthropic、Vertex AI、Bedrock) 内の切り捨てられた不正な JSON を修復する超低遅延リバース プロキシ — 切り捨てられたツール呼び出しでの JSONDecodeError / serde_json EOF を修正します。 · GitHub
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

t
{{ メッセージ }}
テンソルハイク
/
縫合糸修復
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
51 コミット 51 コミット .github .github crates cratesdeploydeploy docs/ blog docs/ blog site site .dockerignore .dockerignore .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md Cargo.toml Cargo.toml Dockerfile Dockerfile LICENSE-APACHE LICENSE-APACHE LICENSE-MIT LICENSE-MIT README.md README.md SECURITY.md SECURITY.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
LLM ストリーミング応答内の切り詰められた不正な JSON をその場で修復する、超低遅延のリバース プロキシ。
📝 ストーリー: LLM ツールがサイレントに Break を呼び出す理由 — および ~10µs の修正
上流の LLM ストリームが max_tokens 、コンテキスト ウィンドウ制限、または
ドロップされたソケット — 出力していた JSON (ツール呼び出しの引数、または構造化出力)
content ) は終了しないままになり、アプリケーションは JSONDecodeError / をスローします。
serde_json 「解析中の EOF」エラー。 Suture はアプリとプロバイダーの間に位置し、
ストリームを監視し、欠落している文字を正確に出力して再構築します。
JSON は有効です — ストリームをバッファリングしたり、意味のある遅延を追加したりすることはありません。
ツール呼び出しストリームが max_tokens で切り捨てられると、クライアントは無効な JSON を再構築したままになります。
// クライアントがデルタ イベントから再構築するもの:
{ "city" : " Par // ← 未終了 → JSONDecodeError / serde_json: 文字列の解析中に EOF
Suture はネットワーク上でそれを閉じるため、クライアントは代わりに有効な JSON を取得します。SDK の変更や変更は必要ありません。
再試行します。再生成されたトークンはありません:
{ "city" : " Par " } // ← 有効;文字列とオブジェクトは安全に閉じられています
これで治る症状
LLM アプリにスローがある場合は、正しい場所にいます。

ストリーミング応答の次のいずれか:
json.decoder.JSONDecodeError: 終了していない文字列が次の位置から始まります: 行 1 列 …
json.decoder.JSONDecodeError: 予期される値: 行 1 列 … (char …)
serde_json::Error: 文字列解析中の EOF / オブジェクト解析中の EOF
切り捨てられたツール呼び出し引数での pydantic_core.ValidationError
モデルが max_tokens に達すると解析されないツール/関数呼び出しの引数
ストリーミングされたデルタ全体で切り詰められた構造化出力/JSON モード コンテンツ
…OpenAI、Anthropic、Google Vertex AI (Gemini / Claude)、または AWS Bedrock 上。
OpenAI ( /v1/chat/completions )、Anthropic ( /v1/messages )、を修復します。
GCP Vertex AI (Gemini + Claude-on-Vertex)、および AWS Bedrock (ConverseStream)
ストリーミング応答。
SSE 対応 — 再アセンブルされたツール呼び出し引数/構造化コンテンツを修復します
生のワイヤーバイトだけでなく、デルタイベント全体で蓄積されます。
ストリーミング + 圧縮 — gzip/brotli/deflate を透過的にデコードし、修復し、
クライアントの Accept-Encoding に従って再エンコードします。体全体を緩衝することはありません。追加されました
オーバーヘッドはチャンクあたり約 10 μs です。
認証情報は保持されません。プロバイダー API キー/ベアラー トークンはそのまま転送されます。
バイトレベルの修復エンジンはスタンドアロン ライブラリとして使用できます。
カーゴは suture-repair を追加します (その後 suture::… を使用します)、または経由でエンジンのみを追加します
カーゴは suture-repair-core を追加します ( suture_core::repair_str を使用します)。
カーゴインストール suture-repair # `suture` バイナリをインストールします。または: docker build -t suture 。
縫合 # 127.0.0.1:8787 でリッスンします
SDK のベース URL を Suture に指定します (API キーは引き続き通過します)。
openaiインポートからOpenAI
client = OpenAI (base_url = "http://localhost:8787/v1" 、api_key = os .environ [ "OPENAI_API_KEY" ])
ルート: POST /v1/chat/completions → OpenAI、POST /v1/messages → Anthropic、
POST /v1/projects/* → 頂点、POST /model/*

→ Bedrock (それぞれ有効になっている場合)、GET /health 。
3 つのレイヤーがそれぞれ独立してテストされています。
suture-core — バイトレベルの JSON 修復ステート マシン。有効なプレフィックスを指定すると、
JSON 値を閉じるために必要な文字を計算します (または、入力が次であることを報告します)。
矛盾しているため、そのまま通過する必要があります)。ネストの深さを超える割り当てはありません。
suture-sse — インクリメンタル SSE パーサー + 再構築するプロバイダーごとのエクストラクター
デルタ イベント全体にわたる JSON を含むフィールドを処理し、コア エンジンを駆動して、
ストリームの最後 (ターミネータの前) での終了イベント。
suture — axum/reqwest リバース プロキシ。リクエストをそのまま転送します。
応答: text/event-stream は SSE 層経由で修復されます。単一の
application/json 本体はコア エンジンで閉じられます。それ以外のものはすべて流れます
変わらない。
環境変数
デフォルト
目的
縫合_聞く
127.0.0.1:8787
リッスンアドレス
SUTURE_OPENAI_BASE
https://api.openai.com
OpenAI アップストリーム
SUTURE_ANTHROPIC_BASE
https://api.anthropic.com
人類の上流
SUTURE_VERTEX_ENABLED
0
Vertex ルート (パスから派生したホスト) を有効にします。
SUTURE_VERTEX_BASE
—
オプションの頂点上流オーバーライド
SUTURE_BEDROCK_ENABLED
0
Bedrock ルートを有効にします (検証済みの Host ヘッダーからのホスト)
SUTURE_BEDROCK_BASE
—
オプションの Bedrock 上流オーバーライド
導入
Dockerfile と Cloud Run、ECS/Fargate については、deploy/ を参照してください。
Kubernetes サイドカー マニフェストと運用上の注意事項 (ストリームをバッファリングしないでください。TLS は
エッジ、ヘルスチェック)。サイドカー パターン (同じ場所にあるローカルホスト) は、
低遅延設計。
OpenAI、Anthropic、GCP Vertex AI、AWS Bedrock ( ConverseStream ) がサポートされています。
透過的な圧縮処理。 Bedrock は認証情報不要の SigV4 パススルーを使用します。
クライアントは実際の Bedrock ホストに署名し、Suture はそのまま転送します。

したがって、縫合糸は決して
再利用可能な AWS シークレット (秘密キーはクライアントから離れることはなく、リクエストごとの署名のみが残されます)
トランジット）。
MIT または Apache-2.0 のいずれかでのデュアルライセンス
オプション。明示的に別段の記載がない限り、意図的に提出された投稿は、
あなたが作品に含める場合は、追加の条件なしで、上記の二重ライセンスが適用されるものとします。
LLM ストリーミング応答 (OpenAI、Anthropic、Vertex AI、Bedrock) 内の切り捨てられた不正な JSON を修復する超低レイテンシーのリバース プロキシ — 切り捨てられたツール呼び出しでの JSONDecodeError / serde_json EOF を修正します。
tensorhq.github.io/suture-stream-repair/
トピックス
Apache-2.0、MITライセンスが見つかりました
ライセンスが見つかりました
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
1
フォーク
レポートリポジトリ
リリース
2
v0.1.1 — Vertex v1beta1 ルーティングの修正
最新の
2026 年 6 月 4 日
+ 1 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Ultra-low-latency reverse proxy that repairs truncated & malformed JSON in LLM streaming responses (OpenAI, Anthropic, Vertex AI, Bedrock) — fixes JSONDecodeError / serde_json EOF on truncated tool calls. - tensorhq/suture-stream-repair

GitHub - tensorhq/suture-stream-repair: Ultra-low-latency reverse proxy that repairs truncated & malformed JSON in LLM streaming responses (OpenAI, Anthropic, Vertex AI, Bedrock) — fixes JSONDecodeError / serde_json EOF on truncated tool calls. · GitHub
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
tensorhq
/
suture-stream-repair
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
51 Commits 51 Commits .github .github crates crates deploy deploy docs/ blog docs/ blog site site .dockerignore .dockerignore .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md Cargo.toml Cargo.toml Dockerfile Dockerfile LICENSE-APACHE LICENSE-APACHE LICENSE-MIT LICENSE-MIT README.md README.md SECURITY.md SECURITY.md View all files Repository files navigation
Ultra-low-latency reverse proxy that repairs truncated and malformed JSON in LLM streaming responses, on the fly.
📝 The story: Why your LLM tool calls silently break — and a ~10µs fix
When an upstream LLM stream is cut off — by max_tokens , a context-window limit, or a
dropped socket — the JSON it was emitting (a tool call's arguments , or structured-output
content ) is left unterminated, and your application throws JSONDecodeError /
serde_json "EOF while parsing" errors. Suture sits between your app and the provider,
watches the stream, and emits exactly the missing characters to make the reassembled
JSON valid — without buffering the stream or adding meaningful latency.
A tool-call stream truncated at max_tokens leaves your client reassembling invalid JSON:
// what the client reassembles from the delta events:
{ "city" : " Par // ← unterminated → JSONDecodeError / serde_json: EOF while parsing a string
Suture closes it on the wire, so the client gets valid JSON instead — no SDK changes, no
retry, no regenerated tokens:
{ "city" : " Par " } // ← valid; the string and object are safely closed
Symptoms this fixes
You're in the right place if your LLM app has thrown any of these on a streaming response:
json.decoder.JSONDecodeError: Unterminated string starting at: line 1 column …
json.decoder.JSONDecodeError: Expecting value: line 1 column … (char …)
serde_json::Error: EOF while parsing a string / EOF while parsing an object
pydantic_core.ValidationError on a truncated tool-call arguments
Tool / function-call arguments that won't parse when the model hits max_tokens
Truncated structured-output / JSON-mode content across streamed deltas
…on OpenAI, Anthropic, Google Vertex AI (Gemini / Claude), or AWS Bedrock.
Repairs OpenAI ( /v1/chat/completions ), Anthropic ( /v1/messages ),
GCP Vertex AI (Gemini + Claude-on-Vertex), and AWS Bedrock ( ConverseStream )
streaming responses.
SSE-aware — repairs the reassembled tool-call arguments / structured content
accumulated across delta events, not just raw wire bytes.
Streaming + compressed — transparently decodes gzip/brotli/deflate, repairs, and
re-encodes per the client's Accept-Encoding ; never buffers the whole body. Added
overhead is ~10 µs per chunk.
Holds no credentials — your provider API key / bearer token is forwarded verbatim.
The byte-level repair engine is usable as a standalone library:
cargo add suture-repair (then use suture::… ), or just the engine via
cargo add suture-repair-core ( use suture_core::repair_str ).
cargo install suture-repair # installs the `suture` binary; or: docker build -t suture .
suture # listens on 127.0.0.1:8787
Point your SDK's base URL at Suture (your API key still flows through):
from openai import OpenAI
client = OpenAI ( base_url = "http://localhost:8787/v1" , api_key = os . environ [ "OPENAI_API_KEY" ])
Routes: POST /v1/chat/completions → OpenAI, POST /v1/messages → Anthropic,
POST /v1/projects/* → Vertex, POST /model/* → Bedrock (each when enabled), GET /health .
Three layers, each independently tested:
suture-core — a byte-level JSON repair state machine. Given any prefix of a valid
JSON value, it computes the characters needed to close it (or reports that the input is
inconsistent and should pass through untouched). No allocation beyond nesting depth.
suture-sse — an incremental SSE parser + per-provider extractors that reassemble
the JSON-bearing field across delta events, drive the core engine, and synthesize a
closing event at stream end (before the terminator).
suture — an axum/reqwest reverse proxy. Forwards your request verbatim, then on
the response: text/event-stream is repaired via the SSE layer; a single
application/json body is closed with the core engine; anything else streams through
unchanged.
Env var
Default
Purpose
SUTURE_LISTEN
127.0.0.1:8787
listen address
SUTURE_OPENAI_BASE
https://api.openai.com
OpenAI upstream
SUTURE_ANTHROPIC_BASE
https://api.anthropic.com
Anthropic upstream
SUTURE_VERTEX_ENABLED
0
enable the Vertex route (host derived from the path)
SUTURE_VERTEX_BASE
—
optional Vertex upstream override
SUTURE_BEDROCK_ENABLED
0
enable the Bedrock route (host from the validated Host header)
SUTURE_BEDROCK_BASE
—
optional Bedrock upstream override
Deployment
See deploy/ for a Dockerfile and Cloud Run, ECS/Fargate, and
Kubernetes-sidecar manifests, plus operational notes (don't buffer the stream, TLS at the
edge, health checks). The sidecar pattern (co-located, localhost) best matches the
low-latency design.
OpenAI, Anthropic, GCP Vertex AI, and AWS Bedrock ( ConverseStream ) are supported, with
transparent compression handling. Bedrock uses credential-free SigV4 passthrough — the
client signs for the real Bedrock host and Suture forwards verbatim, so Suture never sees a
reusable AWS secret (the secret key never leaves the client; only a per-request signature
transits).
Dual-licensed under either of MIT or Apache-2.0 , at your
option. Unless you explicitly state otherwise, any contribution intentionally submitted for
inclusion in the work by you shall be dual-licensed as above, without any additional terms.
Ultra-low-latency reverse proxy that repairs truncated & malformed JSON in LLM streaming responses (OpenAI, Anthropic, Vertex AI, Bedrock) — fixes JSONDecodeError / serde_json EOF on truncated tool calls.
tensorhq.github.io/suture-stream-repair/
Topics
Apache-2.0, MIT licenses found
Licenses found
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
1
fork
Report repository
Releases
2
v0.1.1 — Vertex v1beta1 routing fix
Latest
Jun 4, 2026
+ 1 release
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
