---
source: "https://github.com/varmax2511/trail"
hn_url: "https://news.ycombinator.com/item?id=49150222"
title: "Trail – signed OpenTelemetry spans for AI agents"
article_title: "GitHub - varmax2511/trail · GitHub"
author: "varmax2511"
captured_at: "2026-08-03T01:52:31Z"
capture_tool: "hn-digest"
hn_id: 49150222
score: 1
comments: 0
posted_at: "2026-08-03T01:30:58Z"
tags:
  - hacker-news
  - translated
---

# Trail – signed OpenTelemetry spans for AI agents

- HN: [49150222](https://news.ycombinator.com/item?id=49150222)
- Source: [github.com](https://github.com/varmax2511/trail)
- Score: 1
- Comments: 0
- Posted: 2026-08-03T01:30:58Z

## Translation

タイトル: トレイル – AI エージェントの署名済み OpenTelemetry スパン
記事タイトル: GitHub - varmax2511/trail · GitHub
説明: GitHub でアカウントを作成して、varmax2511/trail の開発に貢献します。

記事本文:
GitHub - varmax2511/trail · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
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
varmax2511
/
トレイル
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード 開く mo

[アクション] メニュー フォルダーとファイル
59 コミット 59 コミット .github/ workflows .github/ workflows docs docs サンプル サンプル テスト トレイル トレイル trail_hook trail_hook .gitignore .gitignore CLAUDE.md CLAUDE.md IMPLEMENTATION_PLAN.md IMPLEMENTATION_PLAN.md ライセンス ライセンス通知 通知 README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml ruff.toml ruff.toml trail_hld.md trail_hld.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェント用の署名付き OpenTelemetry GenAI スパン。キャプチャ、正規化、検証 — 独自のバックエンドを導入します。
Trail は、AI エージェントが実際に行うこと (すべての LLM 呼び出し、ツールの呼び出し、MCP 呼び出し、スキルの実行) を、OpenTelemetry が小さな Trail 拡張名前空間にまたがるときにキャプチャーする Python SDK です。各セッションに Ed25519 で署名し、OTLP 経由で任意の OTel バックエンド (Grafana、Honeycomb、Datadog、Chronosphere など) にエクスポートします。 Trail は、ダッシュボードの保存、クエリ、またはダッシュボードを行いません。ストレージとクエリは既存のバックエンドの仕事です。
実稼働環境でエージェントが不正な動作をした場合、驚くほど答えるのが難しい 3 つの質問があります。
どのツールがどのような入力で、どのような順序で実行されましたか?既存のトレーサは LLM コール形式であり、エージェント形式ではありません。
MCP サーバーの応答は命令を挿入しようとしていたのでしょうか?主流のトレーサーはこれを報告しません。
このスキルは昨日と同じコードですか？デフォルトでは、スキルの置換は痕跡を残しません。
Trail は、エージェント対応ツール分類法 (internal / mcp / skill /builtin)、ツール応答に対する MCP インジェクション フラグ設定、およびサイレント置換を検出するスキル ハッシュという、欠けている 3 つの要素を追加します。これらはすべて標準の OpenTelemetry スパンであるため、OTel バックエンドは翻訳レイヤーなしでそれらを取り込みます。
トレイルは、OpenTelemetry スパン ツリーとして実行されるエージェントをモデル化します。すべての LLM コール、ツール、MCP コール、およびスキルを使用して、セッションごとに 1 つの invoke_agent ルート スパンを使用します。

l はその下にネストされ、trail.* 属性の名前空間がその上に重ねられます。この構造と 3 つの専用の属性によって、上記の各質問がクエリに変換されます。
どのツールがどのような入力で、どのような順序で実行されましたか?
すべてのツールの呼び出しは、gen_ai.tool.name と trail.tool_type (internal / mcp / skill /builtin) でタグ付けされたexecute_tool スパンになります。これは、単純な LLM トレーサーでは決して描画されないエージェント形式の区別です。順序とネストは、OpenTelemetry SDK の contextvars 伝播から得られ、async / await および同時 asyncio タスク全体で正しい状態が保たれるため、各スパンは正しい親に接続されます。入力と出力は、trail.input_hash / trail.output_hash (SHA-256、ホット パスから計算される) と機密性フラグとして記録されます。これは、ペイロード自体を保存せずにペイロードの改ざんが明らかな ID です。
gen_ai.operation.name = "実行ツール"
gen_ai.tool.name = "get_customer_record"
trail.tool_type = "mcp"
trail.input_hash = "sha256:..."
MCP サーバーの応答は命令を挿入しようとしていたのでしょうか?
Trail が MCP call_tool をラップすると、YAML インジェクション ルールセット (命令オーバーライド、システム プロンプト インジェクション、ロール オーバーライド、資格情報 Exfil フレージング ( TRAIL_MCP_RULES によるオーバーライド)) を通じて応答が実行され、スパンに trail.mcp.injection_flag がスタンプされます。 「以前の指示を無視して…」という応答は、来歴用の trail.mcp.server_id の隣に、 trail.mcp.injection_flag = true の通常のスパンとして表示されます。
このスキルは昨日と同じコードですか？
Wrap_skill() は、trail.skill.hash を記録します。スキルのソース ( trail.skill.hash_method = "source" 、C 拡張機能とラムダの qualname-fallback を含む) に対する SHA-256 です。同じスキル→同じハッシュ。サイレント スワップ → 今日のスパンと昨日のスパンの異なるハッシュ。属性の差分

2 つのセッションにまたがって議論すると、置換が表示されます。
次に、あなたはすでにどこを見ているかを尋ねます。トレイルはキャプチャのみです。質問はバックエンドで回答されます。開発モードでは、これはセッション JSONL ( ~/.trail/sessions/{trace_id}.jsonl ) であり、トレイルの verify-export は、それが事後に変更されていないことを証明します (変更されていた場合はスパンを特定します)。 prod では、これは Grafana、Honeycomb、または Datadog の通常の属性フィルター (trail.tool_type = "mcp" AND trail.mcp.injection_flag = true) です。
クイックスタート — OpenAI エージェント (60 秒)
pip インストール ' trail-otel[openai] '
輸入オープンアイ
インポートトレイル
トレイル。 auto_instrument () # openai を検出し、計測します
トレイル付き。セッション (agent_id = "コンテンツパイプライン" ):
クライアント = openai 。オープンAI ()
クライアント。チャット 。完成品。作成 (モデル = "gpt-4o" 、メッセージ = [...])
それだけです。デフォルトでは、Trail はスパンを ~/.trail/sessions/{trace_id}.jsonl に書き込み、短いサマリーを stderr に書き込みます。インフラストラクチャゼロ。
代わりに OTel バックエンドに送信するには:
エクスポート TRAIL_EXPORT=otlp
エクスポート OTEL_EXPORTER_OTLP_ENDPOINT=https://your-collector:4317
Async OpenAI ( AsyncOpenAI ) は、同じ trail.auto_instrument() 呼び出しによって自動的に計測されます。
Trail には、Trail-hook コンソール スクリプトが同梱されています。それを ~/.claude/settings.json に接続します。単一のバイナリが 3 つのイベントすべてを処理します。これは、Claude Code の標準入力ペイロードからイベント名を読み取り、内部的にディスパッチします。
{
「フック」: {
"PreToolUse" : [
{ "マッチャー" : " * " 、 "フック" : [{ "タイプ" : " コマンド " 、 "コマンド" : " トレイルフック " }] }
]、
"PostToolUse" : [
{ "マッチャー" : " * " 、 "フック" : [{ "タイプ" : " コマンド " 、 "コマンド" : " トレイルフック " }] }
]、
"セッション終了" : [
{ "フック" : [{ "タイプ" : " コマンド " , "コマンド" : " トレイルフック " }] }
】
}
}
すでにフックをお持ちですか?クロード コードのフック。<イベント名> は配列です — トレイルは w と並んで構成されます

憎しみはすでにそこにあります。配列を置き換えるのではなく、イベントごとに新しい {matcher,hooks} ブロックを追加します。 Trail は既存のフックを使用して順次実行され、フックをブロックすることはありません (内部エラーが発生しても 0 で終了します)。
エンドツーエンドで見てください:examples/demo_claude_code/run_demo.sh は、トレイルフックに対して実際のクロード コード セッションをリプレイします (インフラストラクチャなし、API キーなし)。ツール分類をキャプチャし、MCP で取得した GitHub の問題に基づくプロンプト インジェクションにフラグを立ててから、verify-export でセッションを証明し、改ざんを捕捉します。
MCP 呼び出しやスキルを含むすべての Claude Code ツール呼び出しがキャプチャされるようになりました。 SessionEnd フックは、セッションが Merkle ルート + Ed25519 署名を取得した瞬間です。 (署名は、セッション終了時に一度起動する SessionEnd に関連付けられています。Stop ではなく、各ターンの終わりに起動され、後のターンのスパンは未署名のままになります。)
Google のエージェント開発キットは OpenTelemetry ネイティブであるため、Trail は ADK 独自の開発キットを利用します
再インスツルメントではなく、execute_tool スパン — 1 回の auto_instrument() 呼び出し
ADK が生成しないツール分類と MCP インジェクション フラグを追加し、ラップします
Trail.session() の実行により署名されます。
インポートトレイル
グーグルから。アドク。ランナーのインポート ランナー
トレイル。 auto_instrument () # google.adk を検出し、そのツール範囲を強化します
トレイル付き。セッション (agent_id = "support-triage" 、プロバイダー = "gcp.vertex" ):
ランナー。 run ( user_id = "u1" 、 session_id = "s1" 、 new_message = msg )
すべての ADK ツール呼び出しには、trail.tool_type ( McpTool → mcp 、
ADK が提供する検索/メモリ ツール → 組み込み 、FunctionTool s → 内部 )
MCP 応答は注入のためにスキャンされます ( trail.mcp.injection_flag )。
まずは開発モードで試してみてください。インフラストラクチャは必要ありません。開発モードがデフォルトなので、
上の 2 行はすでにすべての ADK スパンを次のように書き込んでいます。
~/.trail/sessions/{トレース

_id}.jsonl をローカルに (ネットワークなし)。エージェントを実行してから、
キャプチャされたものを検査します。
cat ~ /.trail/sessions/ <trace_id > .jsonl | jq 。スパン数 + trail.tool_type
署名はオプトインです。キーが存在しない場合、セッションは単純に署名されません。
最小限のセットアップ: スパン + trail.tool_type + MCP インジェクション フラグ、いいえ
改ざんの証拠があり、署名のオーバーヘッドはありません。必要なときにオンにします。
証跡生成キー # 1 回;署名を有効にする
トレイル検証エクスポート ~ /.trail/sessions/ <trace_id > .jsonl
# → VALID (N スパン、署名有効、キー fpr ...)
開発モードの注意: ADK 独自の Cloud Trace / OTel エクスポータも有効にしないでください
開発モードの実行中。 Trail はトレーサー プロバイダーを構成します。 ADK が設定した場合
まず、Trail のローカル JSONL は添付されません。 2 本のトレイル ラインを追加するだけです。
ADK 自身のトレースはオフのままにしておきます。
ローカルで問題がなければ、同じコードをバックエンドに送信します — ADK エクスポート
OTLP なので、TRAIL_EXPORT=otlp と OTEL_EXPORTER_OTLP_ENDPOINT を設定します (例:
Chronosphere) とスパンが代わりにそこに流れます。参照
docs/backends/chronosphere.md 、および
例/adk_manual_instrumentation.py
フレームワークに依存しない手動パスの場合 (アダプターは必要ありません)。
ADK にはファーストクラスの「スキル」がないため、スキルハッシュは trail.wrap_skill のままになります。
ADKで構成されています。並列/マージされたツール呼び出しは、文書化された v1 のギャップです。
標準の OpenTelemetry GenAI 属性:
gen_ai.operation.name = "チャット" | "実行ツール" | 「invoke_agent」
gen_ai.provider.name = "openai" | 「人類的」
gen_ai.request.model = "gpt-4o"
gen_ai.tool.name = "get_customer_record"
gen_ai.usage.input_tokens = 1240
さらに、Trail 拡張機能 - 斬新な部分:
trail.tool_type = "内部" | "mcp" | 「スキル」 | 「内蔵」
trail.mcp.server_id = "acme-crm-mcp"
トレイル.mcp.injection_flag = false
トレイル.スキル.ハッシュ = "sha256:..."
trail.input_hash = "sha256:..."
trail.output_hash = "sha256:..."
で

Grafana または Honeycomb、これらは通常の GenAI スパンとしてレンダリングされます。 trail.* 属性は、他の属性と同様にクエリ可能です ( trail.tool_type = "mcp" および trail.mcp.injection_flag = true )。
指標?コレクターのスパンメトリクス コネクタを使用する
Trail はスパンのみを出力します。Prometheus のスクレイピング エンドポイントや OTel メトリクスはありません。レート / エラー / 期間カウンターまたは「1 分あたりの MCP インジェクション」パネルを取得するには、OpenTelemetry Collector のスパンメトリクス コネクタをパイプラインにドロップし、 trail.tool_type 、 trail.mcp.injection_flag などでラベルを付けます。 スパン バックエンド (Tempo のメトリクス ジェネレーター、Datadog APM メトリクス、Honeycomb 派生列) は、同等のバックエンド側導出を提供します。ソースに 2 つの信号タイプがあると信号が複製されますが、コレクタはそれらをきれいに合成します。
各セッションは、セッション終了時にそのスパン コンテンツのマークル ルートを介して Ed25519 で 1 回署名されます。公開キーを持っている人は誰でも後でそれを検証できます。Trail インフラストラクチャは必要ありません。
トレイル検証-エクスポート session.jsonl
# → 有効 (132 スパン、署名 2026-06-06T10:02:14Z、キー fpr sha256:abcd...)
変更されたスパン、削除されたスパン、および追加されたスパンはすべて、マークル ルートの不一致によって検出されます。
証跡生成キー
# → ~/.trail/keys/trail.key (プライベート、chmod 600)
# → ~/.trail/keys/trail.pub (パブリック)
開発モードと製品モード
モード
スト

[切り捨てられた]

## Original Extract

Contribute to varmax2511/trail development by creating an account on GitHub.

GitHub - varmax2511/trail · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
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
varmax2511
/
trail
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
59 Commits 59 Commits .github/ workflows .github/ workflows docs docs examples examples tests tests trail trail trail_hook trail_hook .gitignore .gitignore CLAUDE.md CLAUDE.md IMPLEMENTATION_PLAN.md IMPLEMENTATION_PLAN.md LICENSE LICENSE NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml ruff.toml ruff.toml trail_hld.md trail_hld.md View all files Repository files navigation
Signed OpenTelemetry GenAI spans for AI agents. Capture, normalize, verify — bring your own backend.
Trail is a Python SDK that captures what AI agents actually do — every LLM call, tool invocation, MCP call, and skill execution — as OpenTelemetry spans with a small Trail extension namespace. It signs each session with Ed25519 and exports via OTLP to any OTel backend (Grafana, Honeycomb, Datadog, Chronosphere, ...). Trail does not store, query, or dashboard. Storage and query are your existing backend's job.
When an agent misbehaves in production, three questions are surprisingly hard to answer:
Which tool ran, in what order, with what inputs? Existing tracers are LLM-call-shaped, not agent-shaped.
Was that MCP server response trying to inject instructions? No mainstream tracer flags this.
Is this skill the same code it was yesterday? Skill substitution leaves no trace by default.
Trail adds the three things that are missing: an agent-aware tool taxonomy ( internal / mcp / skill / builtin ), MCP injection flagging on tool responses, and a skill hash that detects silent substitution — all as standard OpenTelemetry spans, so any OTel backend ingests them with no translation layer.
Trail models an agent run as an OpenTelemetry span tree — one invoke_agent root span per session, with every LLM call, tool, MCP call, and skill nested underneath — and layers a trail.* attribute namespace on top. That structure, plus three purpose-built attributes, is what turns each question above into a query.
Which tool ran, in what order, with what inputs?
Every tool invocation becomes an execute_tool span tagged with gen_ai.tool.name and trail.tool_type ( internal / mcp / skill / builtin ) — the agent-shaped distinction a plain LLM tracer never draws. Order and nesting come from the OpenTelemetry SDK's contextvars propagation, which stays correct across async / await and concurrent asyncio tasks, so each span attaches to the right parent. Inputs and outputs are recorded as trail.input_hash / trail.output_hash (SHA-256, computed off the hot path) plus a sensitivity flag — tamper-evident identity of the payloads without storing the payloads themselves.
gen_ai.operation.name = "execute_tool"
gen_ai.tool.name = "get_customer_record"
trail.tool_type = "mcp"
trail.input_hash = "sha256:..."
Was that MCP server response trying to inject instructions?
When Trail wraps an MCP call_tool , it runs the response through a YAML injection ruleset — instruction-override, system-prompt injection, role override, credential-exfil phrasing (override via TRAIL_MCP_RULES ) — and stamps the span with trail.mcp.injection_flag . A response that says "ignore your previous instructions and…" lands as an ordinary span with trail.mcp.injection_flag = true , next to trail.mcp.server_id for provenance.
Is this skill the same code it was yesterday?
wrap_skill() records trail.skill.hash — a SHA-256 over the skill's source ( trail.skill.hash_method = "source" , with a qualname-fallback for C-extensions and lambdas). Same skill → same hash; a silent swap → a different hash on today's span versus yesterday's. Diff the attribute across two sessions and substitution is visible.
Then you ask where you already look. Trail only captures — the questions get answered in your backend. In dev mode that's the session JSONL ( ~/.trail/sessions/{trace_id}.jsonl ), and trail verify-export proves none of it was altered after the fact (and pinpoints the span if it was). In prod, it's an ordinary attribute filter — trail.tool_type = "mcp" AND trail.mcp.injection_flag = true — in Grafana, Honeycomb, or Datadog.
Quickstart — OpenAI agent (60 seconds)
pip install ' trail-otel[openai] '
import openai
import trail
trail . auto_instrument () # detects openai, instruments it
with trail . session ( agent_id = "content-pipeline" ):
client = openai . OpenAI ()
client . chat . completions . create ( model = "gpt-4o" , messages = [...])
That's it. By default Trail writes spans to ~/.trail/sessions/{trace_id}.jsonl and a short summary to stderr. Zero infrastructure.
To ship to your OTel backend instead:
export TRAIL_EXPORT=otlp
export OTEL_EXPORTER_OTLP_ENDPOINT=https://your-collector:4317
Async OpenAI ( AsyncOpenAI ) is instrumented automatically by the same trail.auto_instrument() call.
Trail ships a trail-hook console script. Wire it into ~/.claude/settings.json . A single binary handles all three events — it reads the event name from Claude Code's stdin payload and dispatches internally.
{
"hooks" : {
"PreToolUse" : [
{ "matcher" : " * " , "hooks" : [{ "type" : " command " , "command" : " trail-hook " }] }
],
"PostToolUse" : [
{ "matcher" : " * " , "hooks" : [{ "type" : " command " , "command" : " trail-hook " }] }
],
"SessionEnd" : [
{ "hooks" : [{ "type" : " command " , "command" : " trail-hook " }] }
]
}
}
Already have hooks? Claude Code's hooks.<EventName> is an array — Trail composes alongside whatever is already there. Append a new {matcher, hooks} block per event rather than replacing the array. Trail runs sequentially with your existing hooks and never blocks them (it exits 0 even on internal errors).
See it end to end: examples/demo_claude_code/run_demo.sh replays a real Claude Code session against trail-hook (no infrastructure, no API key) — captures the tool taxonomy, flags a prompt-injection riding in on an MCP-fetched GitHub issue, then verify-export proves the session and catches a tamper.
Every Claude Code tool call — including MCP calls and skills — is now captured. The SessionEnd hook is the moment the session gets its Merkle root + Ed25519 signature. (Signing is wired to SessionEnd , which fires once when the session terminates — not Stop , which fires at the end of every turn and would leave later turns' spans unsigned.)
Google's Agent Development Kit is OpenTelemetry-native, so Trail rides ADK's own
execute_tool spans rather than re-instrumenting — one auto_instrument() call
adds the tool taxonomy and MCP injection flag ADK doesn't produce, and wrapping
the run in trail.session() signs it.
import trail
from google . adk . runners import Runner
trail . auto_instrument () # detects google.adk, enriches its tool spans
with trail . session ( agent_id = "support-triage" , provider = "gcp.vertex" ):
runner . run ( user_id = "u1" , session_id = "s1" , new_message = msg )
Every ADK tool call now carries trail.tool_type ( McpTool → mcp ,
ADK-provided search/memory tools → builtin , your FunctionTool s → internal )
and MCP responses are scanned for injection ( trail.mcp.injection_flag ).
Try it in dev mode first — zero infrastructure. Dev mode is the default, so
the two lines above already write every ADK span to
~/.trail/sessions/{trace_id}.jsonl locally (no network). Run your agent, then
inspect what was captured:
cat ~ /.trail/sessions/ < trace_id > .jsonl | jq . # spans + trail.tool_type
Signing is opt-in. With no keys present, sessions are simply unsigned —
the minimal setup: spans + trail.tool_type + MCP injection flag, no
tamper-evidence, no signing overhead. Turn it on when you want it:
trail generate-keys # once; enables signing
trail verify-export ~ /.trail/sessions/ < trace_id > .jsonl
# → VALID (N spans, signature valid, key fpr ...)
Dev-mode note: don't also enable ADK's own Cloud Trace / OTel exporter
while running dev mode. Trail configures the tracer provider; if ADK sets one
first, Trail's local JSONL won't attach. Just add the two Trail lines and
leave ADK's own tracing off.
When it looks right locally, ship the same code to your backend — ADK exports
OTLP, so set TRAIL_EXPORT=otlp and OTEL_EXPORTER_OTLP_ENDPOINT (e.g.
Chronosphere) and the spans flow there instead. See
docs/backends/chronosphere.md , and
examples/adk_manual_instrumentation.py
for the framework-agnostic manual path (no adapter required).
ADK has no first-class "skill", so skill-hashing stays with trail.wrap_skill ,
which composes with ADK. Parallel/merged tool calls are a documented v1 gap.
Standard OpenTelemetry GenAI attributes:
gen_ai.operation.name = "chat" | "execute_tool" | "invoke_agent"
gen_ai.provider.name = "openai" | "anthropic"
gen_ai.request.model = "gpt-4o"
gen_ai.tool.name = "get_customer_record"
gen_ai.usage.input_tokens = 1240
Plus the Trail extension — the novel part:
trail.tool_type = "internal" | "mcp" | "skill" | "builtin"
trail.mcp.server_id = "acme-crm-mcp"
trail.mcp.injection_flag = false
trail.skill.hash = "sha256:..."
trail.input_hash = "sha256:..."
trail.output_hash = "sha256:..."
In Grafana or Honeycomb, these render as ordinary GenAI spans. The trail.* attributes are queryable like any other attribute ( trail.tool_type = "mcp" AND trail.mcp.injection_flag = true ).
Metrics? Use the Collector's spanmetrics connector
Trail emits spans only — no Prometheus scrape endpoint and no OTel metrics. To get rate / error / duration counters or a "MCP injections per minute" panel, drop the OpenTelemetry Collector's spanmetrics connector into your pipeline and label by trail.tool_type , trail.mcp.injection_flag , etc. Span backends (Tempo's metrics-generator, Datadog APM metrics, Honeycomb derived columns) offer equivalent backend-side derivations. Two signal types at the source would duplicate the signal — the Collector composes them cleanly.
Each session is signed once at session end with Ed25519 over a Merkle root of its span content. Anyone with the public key can verify it later — no Trail infrastructure required:
trail verify-export session.jsonl
# → VALID (132 spans, signed 2026-06-06T10:02:14Z, key fpr sha256:abcd...)
Modified spans, removed spans, and added spans are all detected by the Merkle root mismatch.
trail generate-keys
# → ~/.trail/keys/trail.key (private, chmod 600)
# → ~/.trail/keys/trail.pub (public)
Dev mode vs prod mode
Mode
Sto

[truncated]
