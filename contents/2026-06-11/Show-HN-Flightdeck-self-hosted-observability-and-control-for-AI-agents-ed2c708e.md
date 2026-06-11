---
source: "https://github.com/flightdeckhq/flightdeck"
hn_url: "https://news.ycombinator.com/item?id=48490231"
title: "Show HN: Flightdeck – self-hosted observability and control for AI agents"
article_title: "GitHub - flightdeckhq/flightdeck: Observability and control plane for AI agents. · GitHub"
author: "pykul"
captured_at: "2026-06-11T16:30:07Z"
capture_tool: "hn-digest"
hn_id: 48490231
score: 2
comments: 0
posted_at: "2026-06-11T13:41:57Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Flightdeck – self-hosted observability and control for AI agents

- HN: [48490231](https://news.ycombinator.com/item?id=48490231)
- Source: [github.com](https://github.com/flightdeckhq/flightdeck)
- Score: 2
- Comments: 0
- Posted: 2026-06-11T13:41:57Z

## Translation

タイトル: Show HN: Flightdeck – AI エージェントの自己ホスト型可観測性と制御
記事のタイトル: GitHub - Flightdeckhq/flightdeck: AI エージェントの可観測性とコントロール プレーン。 · GitHub
説明: AI エージェントの可観測性とコントロール プレーン。 GitHub でアカウントを作成して、flightdeckhq/flightdeck の開発に貢献してください。

記事本文:
GitHub - Flightdeckhq/flightdeck: AI エージェントの可観測性とコントロール プレーン。 · GitHub
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
飛行甲板
/
飛行甲板
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
主要支店

タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
80 コミット 80 コミット .claude-plugin .claude-plugin .claude .claude .github/ workflows .github/ workflows api API アセット アセット ダッシュボード ダッシュボード docker docker helm helm インジェスト インジェスト プレイグラウンド プレイグラウンド プラグイン プラグイン スクリプト スクリプト センサー センサー テスト テスト ワーカー ワーカー .gitignore .gitignore ARCHITECTURE.md ARCHITECTURE.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md DECISIONS.md DECISIONS.md LICENSE LICENSE METHODOLOGY.md METHODOLOGY.md Makefile Makefile README.md README.md RELEASING.md RELEASING.md SECURITY.md SECURITY.md 価格.yaml 価格.yaml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Flightdeck は、プロダクションおよびコーディング エージェントのための自己ホスト型の可観測性およびコントロール プレーンです。
LLM コール、MCP イベント、およびツール呼び出しが発生するたびに、エージェントはダッシュボードにストリームを作成し、エージェントごとのタイムラインとして、またフリート全体のライブ フィードとして表示されます。
実稼働エージェントでトークンの予算、MCP の許可/ブロック ルール、ライブ ディレクティブを設定します。
コーディング エージェントは、このリポジトリの Claude Code プラグインを介して接続されます。
実稼働エージェントは、flightdeck-sensor Python パッケージをエントリポイント ( init() + patch() ) に追加します。その他のコードは変更しません。
前提条件: Docker Engine 28 以降と Compose v2。センサー パスには Python 3.10 以降。クロード プラグイン パスのコード。
git clone https://github.com/flightdeckhq/flightdeck
CD フライトデッキ
開発を行う
http://localhost:4000 のダッシュボード。開発スタックは、テスト トークン tok_dev を自動的にシードします。
Claude Code を起動し、REPL 内のこのリポジトリのマーケットプレイスからプラグインをインストールします。
/プラグインマーケットプレイス追加flightdeckhq/flightdeck
/plugin install Flightdeck@flightdeck-plugins
/reload-プラグイン
ローカルスタックについては以上です — p

lugin のデフォルトは http://localhost:4000 で、開発トークン tok_dev が設定されているため、Claude Code セッションは数秒以内にフリート ビューに表示されます。ツール入力と LLM 呼び出しの内容はデフォルトでキャプチャされます (オプトインするまで Capture_prompts=False を維持する Python センサーとは異なります)。そのため、追加の設定を行わなくても [プロンプト] タブに値が設定されます。
プラグインを別のスタック (本番環境、リモート開発サーバーなど) に向けるには、claude を起動する前にシェルで環境変数をエクスポートします。プラグインは SessionStart ごとに環境変数を読み取ります。
エクスポート FLIGHTDECK_SERVER= " https://flightdeck.example.com "
エクスポート FLIGHTDECK_TOKEN= " ftd_... "
クロード
マーケットプレイスの代わりにローカル チェックアウトを使用するには: claude --plugin-dir /path/to/flightdeck/plugin 。
センサーをインストールし、エージェントにセンサーを向けます。
pip インストール フライトデッキセンサー
フライトデッキセンサーをインポートする
フライトデッキ_センサー 。初期化(
サーバー = "http://localhost:4000/ingest" ,
トークン = "tok_dev" ,
)
フライトデッキ_センサー 。パッチ()
# 既存のエージェント コード。何も変わりません。
輸入人間
クライアント = 人間。人間性 ()
クライアント。メッセージ。 create (model = "claude-sonnet-4-6" , ...)
エージェントは数秒以内にフリート ビューに表示されます。
PyPI ではなくソースからセンサーを実行するには、リポジトリのルートから pip install -e sensor/ を実行します。
サポートされているすべてのフレームワークの実際のサンプルは、 playground/ にあります。各スクリプトの実行あたりのコストは数セントで、実際の LLM API に対してセンサーを実行します。
playground-anthropic を作成する # Anthropic direct
make playground-openai # OpenAI direct
make playground-langchain # LangChain + ChatAnthropic / ChatOpenAI
make playground-langgraph # LangGraph エージェント ループ
make playground-llamaindex # LlamaIndex
playground-crewai # CrewAI マルチエージェントを作成する
make playground-mcp # MCP ツール呼び出し
make playground-policies # トークンポリシーの適用
遊び場をすべて作る # ev

すべて (~$0.50/実行)
API キーが設定されていない場合、各スクリプトは自己スキップするため、make playground-all はどのボックスでも問題なく実行され、資格情報を持っているものだけを実行します。各セッションのフレーバー フィールドには、それを生成したプレイグラウンド スクリプトの名前が付けられるため、ダッシュボードで見つけることができます。完全なマトリックスについては、playground/README.md を参照してください。
プロバイダー
チャット
埋め込み
ストリーミング
エラー
人間的
messages.create 、messages.stream 、beta.messages.* (同期 + 非同期)
litellm経由でVoyageまでルート
同期 + 非同期
14 エントリの llm_error 分類法
OpenAI
chat.completions.create 、response.create (同期 + 非同期)
embeddings.create (同期 + 非同期)
同期 + 非同期
同じ
リテルム
litellm.completion 、 litellm.acompletion (チャット パスのみ)
litellm.embedding 、litellm.aembedding
同期のみ
同じ
ストリーミング イベントは、 payload.streaming = {ttft_ms, chunk_count, inter_chunk_ms,final_outcome, abort_reason} を公開します。ストリームの途中で中止すると、部分チャンクおよび部分トークン データを含む llm_error{error_type="stream_error"} が生成されます。
init() + patch() の後、Anthropic または OpenAI クライアントを内部的に構築するフレームワークは、ユーザー側のラッピングなしでインターセプトされます。
イベントごとのフレームワーク フィールドには、裸の名前 ( langchain 、 crewai など) が含まれます。上位レベルのフレームワークが SDK トランスポートに優先します。LangChain パイプラインは、litellm を介してルーティングされ、 Framework=langchain を報告します。
Claude Code エージェントは、このリポジトリのマーケットプレイスを通じて配布される別のプラグインを介して接続します。
/プラグインマーケットプレイス追加flightdeckhq/flightdeck
/plugin install Flightdeck@flightdeck-plugins
プラグインのデフォルトは、local-dev パスとして http://localhost:4000 + tok_dev です。クロードを起動して別のスタックをポイントする前に、FLIGHTDECK_SERVER + FLIGHTDECK_TOKEN をエクスポートしてください。完全なフローについては、「クイックスタート > コーディング エージェント」を参照してください。ツール入力と LLM 呼び出しの内容はデフォルトでキャプチャされるため、プロンプトは

追加の設定を行わなくても、タブに値が設定されます。
セッションには、flavor=claude-code 、agent_type=coding 、および client_type=claude_code が含まれます。このプラグインはフックベースであり、呼び出し中にディレクティブに作用することはできません。これらのセッションでは、「エージェントの停止」ボタンは非表示になります。書き込み/編集によって書き込まれた生のファイル本体は転送されません。ツール入力はサニタイズされたホワイトリストを通過します。
マルチエージェント フレームワークは、オーケストレーターの親セッションと、parent_session_id でリンクされ、agent_role でラベル付けされたサブエージェント実行ごとの個別の子セッションというツリーとしてレンダリングされます。
マルチエージェント フレームワーク外での直接 SDK 呼び出しはルート セッションを生成します。アイデンティティは変わらない。 Capture_prompts=True の場合、各子セッションは親の入力を incoming_message として運び、子の応答を outcoming_message として返します。これは、実行ドロワーの [サブエージェント] タブに表示されます。
Flightdeck は、MCP トラフィックをチャットやエンベディングと並んでファーストクラスのイベント サーフェスとして観察します。 6 つのイベント タイプ ( mcp_tool_list 、 mcp_tool_call 、 mcp_resource_list 、 mcp_resource_read 、 mcp_prompt_list 、 mcp_prompt_get ) は操作ごとに発行されます。センサーは mcp.client.session.ClientSession に直接パッチを適用するため、公式 SDK を通じて MCP を仲介するすべてのフレームワークが監視されます。 langchain-mcp-adapters 経由の LangChain、同じ経由の LangGraph、llama-index-tools-mcp 経由の LlamaIndex、 mcpadapt 経由の CrewAI、および生の mcp SDK です。
Claude Code プラグインの MCP 対象範囲はツール呼び出しに限定されています。リソースの読み取りとプロンプトのフェッチはプラグイン フック層の下にあります。
ライブフリートビュー ( / )。共有タイムライン上のすべてのエージェント、エージェントごとに 1 行、サブエージェントは親の下にインデントされます。 LLM コール、埋め込み、ツールの使用、ポリシー イベント、および構造化エラーは、到着時に各エージェントの行にプロットされます。実行境界グリフは、各実行の開始時と終了時をマークします。任意のイベントをクリックしてインラインで検査します。
年

nts ダッシュボード ( /agents )。フリート内のすべてのエージェントを並べ替え可能な行として、トークン、レイテンシ、エラー率、セッション数、過去 7 日間のコスト傾向を表示します。フィルター チップは、状態、エージェント タイプ、クライアント タイプ、またはフレームワークによってテーブルを絞り込みます。エージェントをクリックすると、そのエージェントの実行、イベント、サブエージェント関係、および MCP サーバーが含まれるフォーカスされたドロワーが表示されます。独自の時間範囲ピッカーとサブエージェント切り替えを備えた単一エージェントのスイムレーン モーダルのステータス バッジをクリックします。
イベント検索 ( /events )。エージェント、イベント タイプ、エラー タイプ、MCP サーバー、フレームワーク、モデル、終了理由、ポリシー イベント タイプなどによるファセット フィルタリングを備えたフリート全体の個々のイベント。詳細を表示するには、任意のイベントをクリックしてください。実行バッジをクリックして、それを発行した実行を検査します。
検査を実行します。実行を開くと、そのイベントが時系列で表示され、トークン使用状況バー、実行時コンテキスト (git commit、k8s 名前空間、インストールされているフレームワーク、ホスト名、OS)、および生成されたサブエージェントが表示されます。 Capture_prompts=True を使用すると、システム プロンプト、メッセージ、ツール定義、モデル応答、埋め込み入力など、すべての LLM 呼び出しの完全なペイロードが利用可能になります。プロバイダーの形状は保持されます (Anthropic セッションでは、 system 、messages 、 tools 、および response が個別のフィールドとして表示されます。OpenAI セッションでは、システム ロールが含まれたメッセージが表示されます)。
トークンポリシーの施行。フレーバーごとにトークンの予算を一元的に定義します。各エージェントはセッション開始時にポリシーを取得し、コードを変更せずにローカルで適用します。設定可能な警告しきい値（デフォルトは予算の 80%）になると、policy_warn イベントが発生し、通話が続行されます。劣化しきい値 (デフォルトは 90%) になると、policy_degrade はより安価なモデルに切り替わります。ブロックしきい値 (デフォルトは 100%) で、policy_block は BudgetExceededError を発生させます。すべての施行に関する決定は、実行タイムライン上の構造化されたイベントです。
エージェント制御。個人を止めてください

ダッシュボードからの単一エージェントまたはフレーバーのすべてのエージェント。 Kill シグナルは、エージェントの次の LLM 呼び出しで配信されます。カスタム ディレクティブ ( @flightdeck_sensor.directive で装飾された Python 関数) を登録し、ダッシュボードから呼び出します。結果は実行タイムラインに directive_result イベントとして表示されます。
@flightdeck_sensor 。ディレクティブ (
名前 = "clear_cache" ,
description = "プロンプトキャッシュをクリア" ,
パラメータ = [
フライトデッキ_センサー 。パラメータ (
名前 = "キャッシュタイプ" 、
type = "文字列" 、
オプション = [ "すべて" , "プロンプト" ],
デフォルト = "すべて" 、
)
]、
)
def clear_cache ( context 、cache_type = "all" ):
return { "クリア" : my_cache .クリア (キャッシュタイプ)}
MCP サーバー ポリシー。エージェントが通信できる MCP サーバーのフレーバーごとのホワイトリストまたはブロックリスト。センサーは MCP 呼び出しごとに強制します。設定が間違っているサーバーまたは認識されないサーバーは、ポリシーに従って警告またはブロックされます。 「MCP 保護ポリシー」を参照してください。
分析 ( /analytics )。トークン消費量、実行数、レイテンシ、モデル分布、ポリシー イベント量、および共有時間範囲の推定コスト。フレーバー、モデル、フレームワーク、ホスト、エージェント タイプ、チーム、プロバイダー、エージェント ロール、親セッション ID のいずれかによってグループ化されます。
推定コスト。実行ごとのコストは、キャッシュを考慮して公開定価から計算されます。

[切り捨てられた]

## Original Extract

Observability and control plane for AI agents. Contribute to flightdeckhq/flightdeck development by creating an account on GitHub.

GitHub - flightdeckhq/flightdeck: Observability and control plane for AI agents. · GitHub
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
flightdeckhq
/
flightdeck
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
80 Commits 80 Commits .claude-plugin .claude-plugin .claude .claude .github/ workflows .github/ workflows api api assets assets dashboard dashboard docker docker helm helm ingestion ingestion playground playground plugin plugin scripts scripts sensor sensor tests tests workers workers .gitignore .gitignore ARCHITECTURE.md ARCHITECTURE.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md DECISIONS.md DECISIONS.md LICENSE LICENSE METHODOLOGY.md METHODOLOGY.md Makefile Makefile README.md README.md RELEASING.md RELEASING.md SECURITY.md SECURITY.md pricing.yaml pricing.yaml View all files Repository files navigation
Flightdeck is a self-hosted observability and control plane for production and coding agents.
Every LLM call, MCP event, and tool call your agents make streams to the dashboard as it happens, surfaced as a per-agent timeline and as a live fleet-wide feed.
Set token budgets, MCP allow/block rules, and live directives on your production agents.
Coding agents attach via the Claude Code plugin in this repo.
Production agents add the flightdeck-sensor Python package to their entrypoint - init() + patch() , no other code changes.
Prerequisites: Docker Engine 28+ with Compose v2. Python 3.10+ for the sensor path; Claude Code for the plugin path.
git clone https://github.com/flightdeckhq/flightdeck
cd flightdeck
make dev
Dashboard at http://localhost:4000 . The dev stack seeds a test token tok_dev automatically.
Launch Claude Code, then install the plugin from this repo's marketplace inside the REPL:
/plugin marketplace add flightdeckhq/flightdeck
/plugin install flightdeck@flightdeck-plugins
/reload-plugins
That's it for a local stack — the plugin defaults to http://localhost:4000 with the dev token tok_dev , so the Claude Code session shows up in the fleet view within seconds. Tool inputs and LLM call content are captured by default — unlike the Python sensor, which keeps capture_prompts=False until you opt in — so the Prompts tab is populated without extra setup.
To point the plugin at a different stack (production, a remote dev server, etc.) export the env vars in the shell before launching claude — the plugin reads them at every SessionStart :
export FLIGHTDECK_SERVER= " https://flightdeck.example.com "
export FLIGHTDECK_TOKEN= " ftd_... "
claude
To use a local checkout instead of the marketplace: claude --plugin-dir /path/to/flightdeck/plugin .
Install the sensor and point your agent at it:
pip install flightdeck-sensor
import flightdeck_sensor
flightdeck_sensor . init (
server = "http://localhost:4000/ingest" ,
token = "tok_dev" ,
)
flightdeck_sensor . patch ()
# Your existing agent code. Nothing changes.
import anthropic
client = anthropic . Anthropic ()
client . messages . create ( model = "claude-sonnet-4-6" , ...)
The agent shows up in the fleet view within seconds.
To run the sensor from source instead of PyPI: pip install -e sensor/ from the repo root.
Working examples for every supported framework live in playground/ . Each script costs cents per run and exercises the sensor against real LLM APIs.
make playground-anthropic # Anthropic direct
make playground-openai # OpenAI direct
make playground-langchain # LangChain + ChatAnthropic / ChatOpenAI
make playground-langgraph # LangGraph agent loops
make playground-llamaindex # LlamaIndex
make playground-crewai # CrewAI multi-agent
make playground-mcp # MCP tool calls
make playground-policies # token policy enforcement
make playground-all # everything (~$0.50/run)
Each script self-skips when its API keys aren't set, so make playground-all runs cleanly on any box and only exercises what you have credentials for. The flavor field on each session names the playground script that produced it, so you can find them on the dashboard. See playground/README.md for the full matrix.
Provider
Chat
Embeddings
Streaming
Errors
Anthropic
messages.create , messages.stream , beta.messages.* (sync + async)
route via litellm to Voyage
sync + async
14-entry llm_error taxonomy
OpenAI
chat.completions.create , responses.create (sync + async)
embeddings.create (sync + async)
sync + async
same
litellm
litellm.completion , litellm.acompletion (chat path only)
litellm.embedding , litellm.aembedding
sync only
same
Streaming events expose payload.streaming = {ttft_ms, chunk_count, inter_chunk_ms, final_outcome, abort_reason} . Mid-stream aborts emit llm_error{error_type="stream_error"} with partial-chunk and partial-token data.
After init() + patch() , frameworks that build Anthropic or OpenAI clients internally are intercepted with no user-side wrapping.
The per-event framework field carries the bare name ( langchain , crewai , etc.). Higher-level framework wins over SDK transport: a LangChain pipeline routing through litellm reports framework=langchain .
Claude Code agents attach via a separate plugin distributed through this repo's marketplace:
/plugin marketplace add flightdeckhq/flightdeck
/plugin install flightdeck@flightdeck-plugins
The plugin defaults to http://localhost:4000 + tok_dev for the local-dev path; export FLIGHTDECK_SERVER + FLIGHTDECK_TOKEN before launching claude to point at a different stack. See Quickstart > Coding agents for the full flow. Tool inputs and LLM call content are captured by default, so the Prompts tab is populated without extra setup.
Sessions carry flavor=claude-code , agent_type=coding , and client_type=claude_code . The plugin is hook-based and cannot act on directives mid-call; the Stop Agent button is hidden for these sessions. Raw file bodies written by Write / Edit are never forwarded; tool inputs go through a sanitised whitelist.
Multi-agent frameworks render as a tree: a parent session for the orchestrator and a separate child session per sub-agent execution, linked by parent_session_id and labeled with agent_role .
Direct SDK calls outside a multi-agent framework emit root sessions; identity is unchanged. When capture_prompts=True , each child session carries the parent's input as incoming_message and the child's response back as outgoing_message , visible in the run drawer's Sub-agents tab.
Flightdeck observes MCP traffic as a first-class event surface alongside chat and embeddings. Six event types ( mcp_tool_list , mcp_tool_call , mcp_resource_list , mcp_resource_read , mcp_prompt_list , mcp_prompt_get ) emit per operation. The sensor patches mcp.client.session.ClientSession directly, so every framework that mediates MCP through the official SDK is observed: LangChain via langchain-mcp-adapters , LangGraph via the same, LlamaIndex via llama-index-tools-mcp , CrewAI via mcpadapt , plus the raw mcp SDK.
The Claude Code plugin's MCP coverage is limited to tool calls; resource reads and prompt fetches are below the plugin hook layer.
Live fleet view ( / ). Every agent on a shared timeline, one row per agent, sub-agents indented under their parent. LLM calls, embeddings, tool uses, policy events, and structured errors plot on each agent's row as they arrive. Run boundary glyphs mark when each run started and ended. Click any event to inspect it inline.
Agents dashboard ( /agents ). Every agent in the fleet as a sortable row with token, latency, error rate, session count, and cost trends over the last 7 days. Filter chips narrow the table by state, agent type, client type, or framework. Click an agent for a focused drawer with that agent's runs, events, sub-agent relationships, and MCP servers. Click the status badge for a single-agent swimlane modal with its own time-range picker and sub-agent toggle.
Events search ( /events ). Every individual event across the fleet with facet filtering by agent, event type, error type, MCP server, framework, model, close reason, policy event type, and more. Click any event for full detail. Click the run badge to inspect the run that emitted it.
Run inspection. Open any run to see its events in chronological order, the token usage bar, the runtime context (git commit, k8s namespace, frameworks installed, hostname, OS), and the sub-agents it spawned. With capture_prompts=True , every LLM call's full payload is available: system prompt, messages, tool definitions, model response, embedding inputs. Provider shape is preserved (Anthropic sessions display system , messages , tools , and response as separate fields; OpenAI sessions display messages with system role included).
Token policy enforcement. Define token budgets centrally per flavor. Each agent pulls its policy on session start and enforces it locally with no code changes. At a configurable warn threshold (default 80% of budget) a policy_warn event fires and the call proceeds; at the degrade threshold (default 90%) policy_degrade swaps to a cheaper model; at the block threshold (default 100%) policy_block raises BudgetExceededError . Every enforcement decision is a structured event on the run timeline.
Agent control. Stop an individual agent or every agent of a flavor from the dashboard; the kill signal delivers on the agent's next LLM call. Register custom directives (Python functions decorated with @flightdeck_sensor.directive ) and invoke them from the dashboard. Results land as directive_result events on the run timeline.
@ flightdeck_sensor . directive (
name = "clear_cache" ,
description = "Clear the prompt cache" ,
parameters = [
flightdeck_sensor . Parameter (
name = "cache_type" ,
type = "string" ,
options = [ "all" , "prompt" ],
default = "all" ,
)
],
)
def clear_cache ( context , cache_type = "all" ):
return { "cleared" : my_cache . clear ( cache_type )}
MCP server policy. Per-flavor allowlist or blocklist of MCP servers your agents can talk to. The sensor enforces at every MCP call; misconfigured or unrecognized servers either warn or block per the policy. See MCP Protection Policy .
Analytics ( /analytics ). Token consumption, run counts, latency, model distribution, policy event volume, and estimated cost on a shared time range, grouped by any of: flavor, model, framework, host, agent_type, team, provider, agent_role, parent_session_id.
Estimated cost. Per-run cost computed from public list prices, accounting for cach

[truncated]
