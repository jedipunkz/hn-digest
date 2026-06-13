---
source: "https://github.com/YasserCR/galdor"
hn_url: "https://news.ycombinator.com/item?id=48520360"
title: "Show HN: Galdor – a Go LLM agent framework with built-in tracing and replay"
article_title: "GitHub - YasserCR/galdor: A Go-native framework for LLM agents, with OpenTelemetry observability built in. · GitHub"
author: "yassros16"
captured_at: "2026-06-13T21:32:57Z"
capture_tool: "hn-digest"
hn_id: 48520360
score: 4
comments: 0
posted_at: "2026-06-13T19:04:09Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Galdor – a Go LLM agent framework with built-in tracing and replay

- HN: [48520360](https://news.ycombinator.com/item?id=48520360)
- Source: [github.com](https://github.com/YasserCR/galdor)
- Score: 4
- Comments: 0
- Posted: 2026-06-13T19:04:09Z

## Translation

タイトル: Show HN: Galdor – トレースとリプレイが組み込まれた Go LLM エージェント フレームワーク
記事のタイトル: GitHub - YasserCR/galdor: OpenTelemetry の可観測性が組み込まれた LLM エージェント用の Go ネイティブ フレームワーク。 · GitHub
説明: OpenTelemetry オブザーバビリティが組み込まれた、LLM エージェント用の Go ネイティブ フレームワーク。 - YasserCR/galdor

記事本文:
GitHub - YasserCR/galdor: OpenTelemetry 可観測性が組み込まれた LLM エージェント用の Go ネイティブ フレームワーク。 · GitHub
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
ヤセルCR
/
ガルドール
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
C

頌歌
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
114 コミット 114 コミット .github .github cmd/ galdor cmd/ galdor docs docs 例 例 内部 内部メモリ メモリ pkg pkg プロバイダー プロバイダー プロバイダーセット プロバイダーセット スクリプト スクリプト .gitattributes .gitattributes .gitignore .gitignore .golangci.yml .golangci.yml ARCHITECTURE.md ARCHITECTURE.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md COTRIBUTING.md DCO.txt DCO.txt GOVERNANCE.md GOVERNANCE.md ライセンス ライセンス通知 通知 README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md go.mod go.mod go.sum go.sum go.work go.work go.work.sum go.work.sum すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ガルドール (名詞、古英語、9 世紀頃): 呪文、呪文、現実を曲げる唱える言葉。
AI エージェントを構築、オーケストレーション、監視するための Go ネイティブ フレームワーク。
ネイティブの OpenTelemetry。埋め込み型ダッシュボード。バイナリが 1 つ。外部 SaaS はありません。アパッチ2.0。
以下の表は、2026 年 5 月に各プロジェクトのリポジトリ、リリース、公式ドキュメントに対して最後に検証されました。ソースは表の下にリンクされています。何かが漂っているときは PR を歓迎します。
galdor の特徴的なポジション: OTel ネイティブ + シングルバイナリの自己ホスト型ダッシュボード + ファーストパーティ MCP サーバー + ファーストパーティ A2A サーバー、すべて Go で構成されています。他の 4 つのプロジェクトはどれも、今日それらすべてを出荷しません。
スタックで Python が快適に実行され、LangSmith にお金を払っても満足する場合は、LangChain が最も成熟したオプションです。現在、広範囲の Go プロバイダーをカバーする必要がある場合 (ガルドールの 4 つよりも多くのアダプター)、OTel や A2A を犠牲にしても、Eino がさらに進化しています。 Go と MCP のサーバー側公開と A2A 相互運用を 1 か所で行う必要がある場合、現時点では両方のファーストパーティを提供するフレームワークは galdor だけです。
出典 (2026 年 5 月に検証): langchain

-ai/langchain 、 LangSmith セルフホスト ドキュメント 、 tmc/langchaingo 、 cloudwego/eino + eino-ext 、 firebase/genkit/go/plugins/mcp 、 firebase/genkit/go/plugins 、 a2aproject/a2a-go 。
v1.0.0がリリースされました。初期のインテグレータを探しています。
10 フェーズのロードマップは機能的に完成しています: プロバイダーの抽象化 (Anthropic、OpenAI/MiniMax/Groq/Together/DeepSeek/vLLM/Ollama (BaseURL またはプロバイダーセット、Google Gemini、AWS Bedrock 経由)) · リフレクション派生の JSON スキーマを備えたタイプセーフ ツール · チェックポイント、割り込み/再開、ブランチマップの条件付きエッジを備えた有向グラフ ランタイム · ReAct および Plan-and-Execute エージェントヘルパー · 埋め込み SQLite トレース ストア、自動 WAL チェックポイント エクスポータ、自動スタンプされた実行 ID、および孤立スパン警告バナーによるネイティブ OTel 可観測性 · ライブ SSE、実行ごとの DAG、タイムトラベルを備えた埋め込み Web ダッシュボード · 短期メモリ ウィンドウ + 長期メモリ バックエンド (in-mem、SQLite/BM25、pgvector、qdrant) · プロバイダー支援の HTTP/TEI エンベダー· カウンシル マルチエージェント パターン (スーパーバイザー、スウォーム) · 標準入出力、SSE、ストリーミング可能な HTTP 経由の MCP クライアント + サーバー · A2A プロトコル (Google) · LLM-as-judge を備えたインライン評価フレームワーク · スキーマ バインドされた構造化出力 (Go 構造体が入力、デコードされた値が出力) · プロンプト フィンガープリントによる確定的再生 · プロバイダーごとの再試行/バックオフ、実行/ノード タイムアウト、パニック回復、構造化ロギング、Goroutine リーク ゲート、機能を意識した検証、OpenAI 互換の思考モデル用の思考ブロック ストリップ ミドルウェア。
次は、現実世界の統合フィードバックです。 Go で配送エージェントを操作していて、上部のテーブルが共鳴する場合は、スタックで galdor を試してイシューを開いてください。フレームワークは表面を覆っています。残りのエッジは実際の展開でのみ表示されます。 pragma-galdor Retro はそのようなレポートの 1 つであり、v0.1.0 の大部分を形成しました。もっと

歓迎いたします。
v1.0.0 の時点では、 pkg/ のパブリック API は SemVer で安定しています。破壊的な変更は将来の v2 にのみ適用されます。完全な位相追跡と次の内容については、ROADMAP.md を参照してください。
github.com/YasserCR/galdor@v1.0.0 を取得してください
# 必要なプロバイダーに加えて:
github.com/YasserCR/galdor/providers/anthropic@v1.0.0 を取得してください
github.com/YasserCR/galdor/providers/openai@v1.0.0 を取得してください
# または、実行時に環境変数を介してプロバイダーを選択します。
github.com/YasserCR/galdor/providerset@v1.0.0 を取得してください
コア モジュールは必要なものだけをプルします。プロバイダー、メモリ バックエンド、プロトコル アダプターは独自の Go モジュール内に存在するため、依存関係ツリーは緊密なままです。
github.com/YasserCR/galdor/cmd/galdor@v1.0.0 をインストールしてください
galdor ui --db ./traces.db # http://127.0.0.1:7777 を開きます
インストール後に galdor が見つからない場合は、 go install bin ディレクトリが PATH 上にありません。追加するか (export PATH="$(go env GOPATH)/bin:$PATH" )、またはフル パスから galdor Doctor を実行してセットアップを診断します。
20 行で構成される完全な ReAct エージェント:
パッケージメイン
インポート(
「コンテキスト」
「fmt」
「ログ」
「オス」
「github.com/YasserCR/galdor/pkg/agent」
人間的な「github.com/YasserCR/galdor/providers/anthropic」
)
関数メイン () {
p、err:= 人間性。新規 (anthropic.Config { APIKey : os .Getenv ( "ANTHROPIC_API_KEY" )})
エラーの場合 != nil {
ログ。致命的 (エラー)
}
答え、エラー:= エージェント。実行 ( context . バックグラウンド ()、エージェント. Config {
プロバイダー:p、
モデル: "クロード-俳句-4-5" 、
}, 「エクアドルの首都はどこですか?」 )
エラーの場合 != nil {
ログ。致命的 (エラー)
}
fmt 。 Println (答え)
}
anthropic を openai (BaseURL 経由で MiniMax / Groq / Together / Mistral で動作)、Google (Gemini)、または Bedrock に交換しても、他は何も変わりません。
タイプセーフ ツール (ジェネリック + リフレクション派生 JSON スキーマ)
インポート(
「コンテキスト」
「github.com/YasserCR/galdor/pkg/tool」
)
タイプ天気In構造体{
都市

tring `json:"city" jsonschema:"required、検索する都市"`
}
type 天気出力構造体 {
一時 float64 `json:"temp_c"`
空の文字列 `json:"sky"`
}
天気:=ツール。 MustNewTool ( "weather" , "都市の天気を調べる" ,
func (ctx context.Context , in WeatherIn ) (weatherOut , error ) {
return WeatherOut { 気温 : 18.5 、空 : "晴れ" }、nil
})
reg 、 _ := ツール 。 NewRegistry (天気)
答え、_ := エージェント。実行 ( ctx 、エージェント。構成 {
プロバイダー: p 、ツール: reg 、モデル: "claude-haiku-4-5" 、
}, 「キトの天気はどうですか?」 )
In と Out は実際の Go タイプです。LLM に公開される JSON スキーマは、 In のリフレクション メタデータから派生します。マジック ストリングもインターフェースもありません。{}
ネイティブ OpenTelemetry — 組み込まれており、ボルトオンではありません
インポート(
sdktrace "go.opentelemetry.io/otel/sdk/trace"
「github.com/YasserCR/galdor/pkg/observability」
)
エクスポーター、_ := 可観測性。 NewSQLiteExporter ( "./traces.db" )
tp:= sdktrace 。 NewTracerProvider ( sdktrace .WithBatcher ( エクスポーター ))
トレーサー := tp 。トレーサー (「私のエージェント」)
// プロバイダーをラップします。すべての LLM 呼び出しでスパンが生成されます。
p = 可観測性。 InstrumentProvider ( p 、トレーサー 、
可観測性。 WithCaptureContent ( true ))
すべての LLM 呼び出し、ツール呼び出し、およびグラフ ノードは、GenAI セマンティック規則に従って OTel スパンになります。それらを galdor ui で検査するか、既存の Datadog / Honeycomb / Grafana スタックにパイプします。同じデータ、コンシューマの選択です。
マルチエージェント: スーパーバイザーとスウォームの組み込み
インポート「github.com/YasserCR/galdor/pkg/council」
監督者、_ := 評議会。 NewSupervisor (評議会.SupervisorConfig {
プロバイダー : p 、モデル : "claude-haiku-4-5" 、
労働者 : []評議会。労働者 {
{ 名前 : "請求" 、説明 : "請求書、返金を処理します" 、
実行: billingWorker },
{ 名前 : "技術的" 、説明 : "バグを診断します。

年齢" 、
実行：technicalWorker }、
}、
})
最終、_ := スーパーバイザー。呼び出し ( ctx 、 Council.SupervisorState { 入力 : userMessage })
各ターンをスペシャリストに委任する、スクリプト化された LLM ルーティング スーパーバイザ。完全な例は、examples/integration-support-bot を参照してください。
InterruptBefore による人間参加型
g := グラフ 。新しい [ TransferState ]()。
AddNode ( "validate" , validate )。
AddNode ( "実行" 、実行)。
AddEdge ( グラフ . START 、 "検証" )。
AddEdge (「検証」、「実行」)。
InterruptBefore ( "execute" ) // ← 人間の承認のために一時停止します
r 、 _ := g 。コンパイル ()
ckpt := グラフ 。 NewMemoryCheckpointer [ TransferState ]()
// フェーズ 1: ゲートまで走ります。 ErrInterrupted を返します。
_ 、エラー:= r 。 InvokeWith ( ctx 、 init 、graph.RunOptions [ TransferState ]{
実行ID : runID 、チェックポインタ : ckpt 、
})
// フェーズ 2: 人間が状態をレビューして編集します。
ck 、 _ 、 _ := ckpt 。ロード ( ctx 、 runID )
決定 := プロンプトヒューマン ( ck . State ) // UI / Slack ボット / など。
// フェーズ 3: 決定を注入して再開します。
最終、_ := r 。再開 ( ctx 、グラフ。RunOptions [ TransferState ]{
RunID : runID 、チェックポインタ : ckpt 、OverrideState : & 決定 、
})
監査可能な安全な建設の承認フロー。 「examples/integration-approval-gate」を参照してください。
リプレイ: 有料 API → フィクスチャ → 決定論的テスト
// 1 回限り: プロンプト/完了キャプチャをオンにして実際の実行を記録します。
// 次に録音をエクスポートします。
//
// ガルドールの占術リプレイ <run-id> -o fixture.json
// 以降永遠に: CI で実行を無料でリプレイします。
記録 、 _ := 再生 。 LoadFromFile ( "fixture.json" )
モック:= リプレイ。 NewProvider (rec.Calls、replay.ModeStrict)
r , _ := エージェント。 NewReAct (エージェント.Config {プロバイダー:モック、モデル: "..."、ツール:reg })
最終、_ := r 。呼び出し ( ctx 、 state )
// プロンプトがずれている場合、ErrPromptMismatch は正確にどの c かを示します。

全て。
ネットワークにアクセスせず、トークンを焼き付けないプロンプトとエージェントの回帰テスト。補完的な予算執行パターンについては、examples/integration-cost-tracked を参照してください。
MCP サーバー: 20 行でツールを Claude Desktop に公開します
インポート(
「github.com/YasserCR/galdor/pkg/mcp」
「github.com/YasserCR/galdor/pkg/tool/builtins」
)
関数メイン () {
今、_ := 組み込み関数。新しい時間ツール ()
math 、 _ := 組み込み関数。新しい数学ツール ()
reg 、 _ := ツール 。 NewRegistry ( now 、 math 、 yourCustomTool )
srv := mcp 。 NewServer ( reg , mcp.ServerInfo { 名前 : "my-tools" , バージョン : "0.1" })
トランスポート := mcp 。 NewStdioTransport ( os . Stdin 、 os . Stdout )
_ = サーバー。提供 (コンテキスト . バックグラウンド ()、トランスポート )
}
バイナリをビルドし、Claude Desktop の claude_desktop_config.json を指定して、Claude Desktop を再起動します。ツールがピッカーに表示されます。詳細な手順は、examples/integration-mcp-server にあります。
多くのクライアントが共有する存続期間の長いデーモンの場合は、トランスポートを交換します。現在の IDE 互換性のために SSE を、2024 年 11 月 5 日以降の仕様のために Streamable HTTP を使用します。
// 2024-11-05 より前の仕様 (SSE トランスポートの Cursor/Claude Desktop は依然としてデフォルトです)
トランスポート := mcp 。 NewSSETransport ( ":4000" )
// 2024-11-05 仕様 (単一エンドポイント、Mcp-Session-Id ヘッダーによるセッション ID)
輸送:=

[切り捨てられた]

## Original Extract

A Go-native framework for LLM agents, with OpenTelemetry observability built in. - YasserCR/galdor

GitHub - YasserCR/galdor: A Go-native framework for LLM agents, with OpenTelemetry observability built in. · GitHub
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
YasserCR
/
galdor
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
114 Commits 114 Commits .github .github cmd/ galdor cmd/ galdor docs docs examples examples internal internal memory memory pkg pkg providers providers providerset providerset scripts scripts .gitattributes .gitattributes .gitignore .gitignore .golangci.yml .golangci.yml ARCHITECTURE.md ARCHITECTURE.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md DCO.txt DCO.txt GOVERNANCE.md GOVERNANCE.md LICENSE LICENSE NOTICE NOTICE README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md go.mod go.mod go.sum go.sum go.work go.work go.work.sum go.work.sum View all files Repository files navigation
galdor (n., Old English, c. 9th century): incantation, spell, a chanted word that bends reality.
A Go-native framework for building, orchestrating and observing AI agents.
Native OpenTelemetry. Embedded dashboard. One binary. No external SaaS. Apache 2.0.
The table below was last verified against each project's repo, releases and official docs in May 2026. Sources are linked under the table; PRs welcome when something drifts.
galdor's distinctive position: OTel-native + a single-binary self-hosted dashboard + first-party MCP server + first-party A2A server , all in Go. None of the other four projects ship all of those today.
If your stack runs Python comfortably and you're happy paying for LangSmith, LangChain is the most mature option. If you need broad Go provider coverage today (more adapters than galdor's four), Eino is further along — at the cost of no OTel and no A2A. If you need Go and MCP server-side exposure and A2A interop in one place, galdor is currently the only framework that ships both first-party.
Sources (verified May 2026): langchain-ai/langchain , LangSmith self-host docs , tmc/langchaingo , cloudwego/eino + eino-ext , firebase/genkit/go/plugins/mcp , firebase/genkit/go/plugins , a2aproject/a2a-go .
v1.0.0 released. Looking for early integrators.
The 10-phase roadmap is functionally complete: provider abstraction (Anthropic, OpenAI/MiniMax/Groq/Together/DeepSeek/vLLM/Ollama via BaseURL or providerset , Google Gemini, AWS Bedrock) · type-safe tools with reflection-derived JSON schemas · directed graph runtime with checkpoints, interrupt/resume and branch-map conditional edges · ReAct and Plan-and-Execute agent helpers · native OTel observability with embedded SQLite trace store, auto-WAL-checkpointing exporter, auto-stamped run ids, and an orphan-span warning banner · embedded web dashboard with live SSE, per-run DAG, time-travel · short-term memory windows + long-term memory backends (in-mem, SQLite/BM25, pgvector, qdrant) · provider-backed and HTTP/TEI embedders · Council multi-agent patterns (Supervisor, Swarm) · MCP client + server over stdio, SSE, and Streamable HTTP · A2A protocol (Google) · inline eval framework with LLM-as-judge · schema-bound structured output (a Go struct in, a decoded value out) · deterministic replay with prompt fingerprinting · per-provider retry/backoff, run/node timeouts, panic recovery, structured logging, goroutine leak gates, capability-aware validation · thinking-block strip middleware for OpenAI-compat thinking models.
What's next: real-world integration feedback. If you're shipping agents in Go and the table at the top resonates, try galdor on your stack and open an issue — the framework has covered the surface; the remaining edges only show up in actual deployments. The pragma-galdor retro is one such report, and it shaped most of v0.1.0; more would be welcome.
As of v1.0.0 , the public API under pkg/ is stable under SemVer : breaking changes only land in a future v2. See ROADMAP.md for full phase tracking and what's next.
go get github.com/YasserCR/galdor@v1.0.0
# plus the provider(s) you need:
go get github.com/YasserCR/galdor/providers/anthropic@v1.0.0
go get github.com/YasserCR/galdor/providers/openai@v1.0.0
# or pick a provider at runtime via env var:
go get github.com/YasserCR/galdor/providerset@v1.0.0
The core module pulls only what it needs — providers, memory backends and protocol adapters live in their own Go modules so your dependency tree stays tight.
go install github.com/YasserCR/galdor/cmd/galdor@v1.0.0
galdor ui --db ./traces.db # open http://127.0.0.1:7777
If galdor isn't found after installing, the go install bin directory isn't on your PATH — add it ( export PATH="$(go env GOPATH)/bin:$PATH" ), or run galdor doctor from the full path to diagnose your setup.
A complete ReAct agent in 20 lines:
package main
import (
"context"
"fmt"
"log"
"os"
"github.com/YasserCR/galdor/pkg/agent"
anthropic "github.com/YasserCR/galdor/providers/anthropic"
)
func main () {
p , err := anthropic . New (anthropic. Config { APIKey : os . Getenv ( "ANTHROPIC_API_KEY" )})
if err != nil {
log . Fatal ( err )
}
answer , err := agent . Run ( context . Background (), agent. Config {
Provider : p ,
Model : "claude-haiku-4-5" ,
}, "What is the capital of Ecuador?" )
if err != nil {
log . Fatal ( err )
}
fmt . Println ( answer )
}
Swap anthropic for openai (works with MiniMax / Groq / Together / Mistral via BaseURL ), google (Gemini), or bedrock and nothing else changes.
Type-safe tools (generics + reflection-derived JSON Schema)
import (
"context"
"github.com/YasserCR/galdor/pkg/tool"
)
type weatherIn struct {
City string `json:"city" jsonschema:"required, city to look up"`
}
type weatherOut struct {
Temp float64 `json:"temp_c"`
Sky string `json:"sky"`
}
weather := tool . MustNewTool ( "weather" , "Look up the weather for a city" ,
func ( ctx context. Context , in weatherIn ) ( weatherOut , error ) {
return weatherOut { Temp : 18.5 , Sky : "clear" }, nil
})
reg , _ := tool . NewRegistry ( weather )
answer , _ := agent . Run ( ctx , agent. Config {
Provider : p , Tools : reg , Model : "claude-haiku-4-5" ,
}, "How's the weather in Quito?" )
In and Out are real Go types — the JSON schema published to the LLM is derived from In 's reflection metadata. No magic strings, no interface{} .
Native OpenTelemetry — built in, not bolted on
import (
sdktrace "go.opentelemetry.io/otel/sdk/trace"
"github.com/YasserCR/galdor/pkg/observability"
)
exporter , _ := observability . NewSQLiteExporter ( "./traces.db" )
tp := sdktrace . NewTracerProvider ( sdktrace . WithBatcher ( exporter ))
tracer := tp . Tracer ( "my-agent" )
// Wrap your provider — every LLM call now produces a span.
p = observability . InstrumentProvider ( p , tracer ,
observability . WithCaptureContent ( true ))
Every LLM call, tool invocation, and graph node becomes an OTel span following the GenAI semantic conventions. Inspect them with galdor ui or pipe them to your existing Datadog / Honeycomb / Grafana stack — same data, your choice of consumer.
Multi-agent: Supervisor and Swarm built in
import "github.com/YasserCR/galdor/pkg/council"
supervisor , _ := council . NewSupervisor (council. SupervisorConfig {
Provider : p , Model : "claude-haiku-4-5" ,
Workers : []council. Worker {
{ Name : "billing" , Description : "handles invoices, refunds" ,
Run : billingWorker },
{ Name : "technical" , Description : "diagnoses bugs, outages" ,
Run : technicalWorker },
},
})
final , _ := supervisor . Invoke ( ctx , council. SupervisorState { Input : userMessage })
A scripted-LLM routing supervisor that delegates each turn to specialists. See the full example: examples/integration-support-bot .
Human-in-the-loop with InterruptBefore
g := graph . New [ TransferState ]().
AddNode ( "validate" , validate ).
AddNode ( "execute" , execute ).
AddEdge ( graph . START , "validate" ).
AddEdge ( "validate" , "execute" ).
InterruptBefore ( "execute" ) // ← pause for human approval
r , _ := g . Compile ()
ckpt := graph . NewMemoryCheckpointer [ TransferState ]()
// Phase 1: run until the gate. Returns ErrInterrupted.
_ , err := r . InvokeWith ( ctx , init , graph. RunOptions [ TransferState ]{
RunID : runID , Checkpointer : ckpt ,
})
// Phase 2: human reviews and edits state.
ck , _ , _ := ckpt . Load ( ctx , runID )
decision := promptHuman ( ck . State ) // your UI / Slack bot / etc.
// Phase 3: resume with the decision injected.
final , _ := r . Resume ( ctx , graph. RunOptions [ TransferState ]{
RunID : runID , Checkpointer : ckpt , OverrideState : & decision ,
})
Auditable, safe-by-construction approval flows. See examples/integration-approval-gate .
Replay: paid-API → fixture → deterministic test
// One-time: record a real run with prompt/completion capture on,
// then export the recording.
//
// galdor scry replay <run-id> -o fixture.json
// Forever after: replay the run for free in CI.
rec , _ := replay . LoadFromFile ( "fixture.json" )
mock := replay . NewProvider ( rec . Calls , replay . ModeStrict )
r , _ := agent . NewReAct (agent. Config { Provider : mock , Model : "..." , Tools : reg })
final , _ := r . Invoke ( ctx , state )
// If your prompts drifted, ErrPromptMismatch tells you exactly which call.
Regression tests for prompts and agents that don't hit the network and don't burn tokens. See examples/integration-cost-tracked for the complementary budget-enforcement pattern.
MCP server: expose your tools to Claude Desktop in 20 lines
import (
"github.com/YasserCR/galdor/pkg/mcp"
"github.com/YasserCR/galdor/pkg/tool/builtins"
)
func main () {
now , _ := builtins . NewTimeTool ()
math , _ := builtins . NewMathTool ()
reg , _ := tool . NewRegistry ( now , math , yourCustomTool )
srv := mcp . NewServer ( reg , mcp. ServerInfo { Name : "my-tools" , Version : "0.1" })
transport := mcp . NewStdioTransport ( os . Stdin , os . Stdout )
_ = srv . Serve ( context . Background (), transport )
}
Build the binary, point Claude Desktop's claude_desktop_config.json at it, restart Claude Desktop. Your tools appear in the picker. Full instructions in examples/integration-mcp-server .
For long-lived daemons that many clients share, swap the transport — SSE for IDE-compatibility today, Streamable HTTP for the post-2024-11-05 spec:
// pre-2024-11-05 spec (the SSE transport Cursor/Claude Desktop still default to)
transport := mcp . NewSSETransport ( ":4000" )
// 2024-11-05 spec (single endpoint, session id via Mcp-Session-Id header)
transport :=

[truncated]
