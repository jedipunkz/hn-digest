---
source: "https://github.com/andrewyng/aisuite"
hn_url: "https://news.ycombinator.com/item?id=48562975"
title: "Andrew Ng ships Open CoWorker desktop AI agent"
article_title: "GitHub - andrewyng/aisuite: Simple, unified interface to multiple Generative AI providers · GitHub"
author: "Gaishan"
captured_at: "2026-06-17T01:03:10Z"
capture_tool: "hn-digest"
hn_id: 48562975
score: 5
comments: 1
posted_at: "2026-06-16T22:17:06Z"
tags:
  - hacker-news
  - translated
---

# Andrew Ng ships Open CoWorker desktop AI agent

- HN: [48562975](https://news.ycombinator.com/item?id=48562975)
- Source: [github.com](https://github.com/andrewyng/aisuite)
- Score: 5
- Comments: 1
- Posted: 2026-06-16T22:17:06Z

## Translation

タイトル: Andrew Ng が Open CoWorker デスクトップ AI エージェントを出荷
記事のタイトル: GitHub - andrewyng/aisuite: 複数の生成 AI プロバイダーへのシンプルで統合されたインターフェイス · GitHub
説明: 複数のジェネレーティブ AI プロバイダーへのシンプルで統合されたインターフェイス - GitHub - andrewyng/aisuite: 複数のジェネレーティブ AI プロバイダーへのシンプルで統合されたインターフェイス

記事本文:
GitHub - andrewyng/aisuite: 複数の生成 AI プロバイダーへのシンプルで統合されたインターフェイス · GitHub
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
アンドリューイング
/
アイスイート
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
本支店

es タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
373 コミット 373 コミット .github/ workflows .github/ workflows aisuite-js aisuite-js aisuite aisuite cli/ py/ aisuite-code-cli cli/ py/ aisuite-code-cli docs docs サンプル サンプル ガイド ガイド プラットフォーム プラットフォーム スクリプト スクリプト テスト テスト viewer-ui viewer-ui .env.sample .env.sample .git-blame-ignore-revs .git-blame-ignore-revs .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README-alternative.md README-alternative.md README.md README.md 詩.ロック 詩.ロック pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
aisuite 上に構築された、デスクトップ上に常駐する AI エージェント。
OpenCoworker は、チャットだけでなく、詳細な調査やタスクの実行もできるデスクトップ AI エージェントです。
コンピューター上のあなた。 (許可を得て) ファイルを読み取り、コンテキストを取得したり、メッセージ (Slack、電子メールなど) を読み取り/送信したりできます。
PDF レポート、ドキュメント、スプレッドシートなどの実際の成果物を作成します。スケジュールされた自動化もサポートしています。
毎日のニュースの概要を提供するなど。
独自の API キー (OpenAI、Anthropic、Google) を使用するか、Ollama を使用して完全にローカルで実行する必要があります。データはマシン上に残ります。
⬇ macOS 用ダウンロード
macOS 13+ (Apple シリコン)
⬇ Windows 用ダウンロード
Windows 10/11 (x64) ·
クイックスタート: — モデルのインストール、接続、最初のタスク、自動化。
そのソースは、このリポジトリの platform/ にあります。これは、aisuite 上で独自のエージェント ハーネスを構築するための作業参照です。
aisuite は、LLM を使用して構築するための軽量の Python ライブラリで、プロバイダー間で統合されたチャット完了 API と、その上にツールとツールキットを備えたエージェント API の 2 つのレイヤーで構成されています。このリポジトリは、aisuite を使用して構築されたデスクトップ AI コワーカーである OpenCoworker の本拠地でもあります。
┌─

───────────────────────┐
│ OpenCoworker │ 日常業務を実行するためのエージェントハーネス
━━━━━━━━━━━━━━━━━━━━┤
│ エージェント API、ツールキット、MCP │ 複数の LLM にわたるエージェントを構築
━━━━━━━━━━━━━━━━━━━━┤
│ チャット完了 API │ 複数の LLM プロバイダーにわたる 1 つの API
├───┬─────┬───┬───┬───┤
│ OpenAI │ Anthropic │ Google │ Ollama │ その他 │
━───┴───────┴───┴───┴─┘
Chat Completions API — OpenAI、Anthropic、Google、Mistral、Hugging Face、AWS、Cohere、Ollama、OpenRouter などのための統合された OpenAI スタイルのインターフェイス。 1 つの文字列を変更してプロバイダーを切り替えます。
エージェント API · ツールキット · MCP — モデルに実際の Python 関数をツールとして提供し、マルチターン ループを実行し、既製のツールキット (ファイル、git、シェル) または任意の MCP サーバーを接続し、ツール ポリシーですべてを管理します。
OpenCoworker — aisuite を使用して構築されたデスクトップ AI コワーカーで、日常業務用のアプリとして出荷されます。
基本パッケージをインストールするか、使用する予定のプロバイダーの SDK を含めます。
pip install aisuite # 基本パッケージ、プロバイダー SDK なし
pip install ' aisuite[anthropic] ' # 特定のプロバイダーの SDK を使用
pip install ' aisuite[all] ' # すべてのプロバイダ SDK を使用
呼び出すプロバイダーの API キーも必要です

— チャット完了クイックスタートでは、主要なセットアップと最初の通話について説明します。
OpenCoworker アプリ (デスクトップ)
インストーラーをダウンロードし、独自の API キーを持参します (または Ollama でローカル モデルを実行します)。
⬇ macOS (Apple Silicon) · ⬇ Windows 10/11 (x64) · OpenCoworker クイックスタート
チャット完了 — 複数のプロバイダーにわたる 1 つの API
チャット API は、モデルの対話のための高レベルの抽象化を提供します。プロバイダーに依存しない方法ですべてのコア パラメーター ( pressure 、 max_tokens 、 tools など) をサポートし、リクエストとレスポンスの構造を標準化するため、SDK の違いではなくロジックに集中できます。
モデル名には <provider>:<model-name> という形式が使用されます。 aisuite は、適切なパラメータを使用して通話を適切なプロバイダーにルーティングします。
aisuiteをaiとしてインポート
クライアント = ai 。クライアント ()
モデル = [ "openai:gpt-4o" , "anthropic:claude-3-5-sonnet-20240620" ]
メッセージ = [
{ "role" : "system" , "content" : "海賊英語で応答します。" }、
{ "role" : "user" , "content" : "冗談を言ってください。" }、
】
モデル内のモデルの場合:
応答 = クライアント 。チャット 。完成品。作成(
モデル = モデル、
メッセージ = メッセージ 、
温度 = 0.75
)
print (応答の選択肢「0」のメッセージの内容)
→ クイックスタート: docs/chat-completions-quickstart.md — インストール、キーのセットアップ、ローカル モデル、その他の例。
エージェント — モデルに実際のツールを与える
aisuite はツール呼び出しをワンライナーに変換します。単純な Python 関数を渡すと、スキーマが生成され、呼び出しが実行され、結果がモデルにフィードバックされます。
def will_it_rain (場所: str 、時刻: str ):
"""今日の特定の時間にその場所で雨が降るかどうかを確認します。
引数:
location (文字列): 都市の名前
time_of_day (文字列): HH:MM 形式の時刻。
「」
「はい」を返します
クライアント = ai 。クライアント ()
応答 = クライアント 。チャット 。完成品。作成(
モデル = "openai:gpt-4o" ,
メッセージ = [{

"ロール" : "ユーザー" 、
"content" : "私はサンフランシスコに住んでいます。天気を調べてもらえますか"
「それで、午後 2 時に屋外ピクニックを計画してくれる？」
}]、
ツール = [ will_it_rain ]、
max_turns = 2 # 往復ツール呼び出しの最大数
)
print (応答の選択肢「0」のメッセージの内容)
max_turns を設定すると、aisuite はメッセージを送信し、モデル リクエストを呼び出すツールを実行し、結果をモデルに返し、会話が完了するまで繰り返します。会話を継続したい場合は、response.choices[0].intermediate_messages に完全なツール対話履歴が保存されます。
完全な手動制御を好みますか? max_turns を省略し、OpenAI 形式の JSON ツール仕様を渡します。aisuite はモデルのツール呼び出しリクエストを返し、ループを自分で実行します。両方のスタイルについては、examples/tool_calling_abstraction.ipynb を参照してください。
長時間実行される構造化された作業には、ファーストクラスのエージェント API があります。エージェントを 1 回宣言し、 Runner で実行し、ファイル、git、およびシェル用の事前に構築されたサンドボックス ツール ファミリであるツールキットをアタッチします。
aisuiteをaiとしてインポート
aisuiteインポートエージェント、ランナーより
エージェント = エージェント (
name = "リポヘルパー" ,
モデル = "anthropic:claude-sonnet-4-6" ,
指示 = "あなたは慎重なリポジトリアシスタントです。ツールを使用してコードから答えてください。" 、
tools = [ * ai .ツールキット 。ファイル ( root = "." )、 * ai .ツールキット 。 git ( root = "." )]、
)
結果 = ランナー 。 run (agent , "最後のコミットで何が変更されましたか? 3 つの箇条書きで要約してください。" )
print (結果.final_output)
エージェント API は、実稼働ハーネスに必要な要素も提供します。
ツール ポリシー — RequireApprovalPolicy 、許可/拒否リスト、または実行するツール呼び出しを決定する独自の呼び出し可能オブジェクト。
状態ストア — 実行 (メモリ内、ファイル、または Postgres) を保持および再開し、プロセス間で会話を継続します。
アーティファクトとトレース — エージェントが生成したものをキャプチャし、

途中で必要なすべてのステップ。
aisuite はモデル コンテキスト プロトコルをネイティブにサポートしているため、任意の MCP サーバーのツールをボイラープレートなしでモデルに渡すことができます ( pip install 'aisuite[mcp]' )。
クライアント = ai 。クライアント ()
応答 = クライアント 。チャット 。完成品。作成(
モデル = "openai:gpt-4o" ,
messages = [{ "role" : "user" , "content" : "現在のディレクトリ内のファイルを一覧表示する" }],
ツール = [{
"タイプ" : "mcp" ,
"名前" : "ファイルシステム" ,
"コマンド" : "npx" ,
"args" : [ "-y" , "@modelcontextprotocol/server-filesystem" , "/path/to/directory" ]
}]、
最大回転数 = 3
)
print (応答の選択肢「0」のメッセージの内容)
再利用可能な接続、セキュリティ フィルター、およびツールのプレフィックスには、明示的な MCPClient を使用します。
→ クイックスタート: docs/agents-quickstart.md — 手動ツールの処理、完全なエージェント API、ポリシー、状態ストア、および MCP の詳細。
aisuite の拡張: プロバイダーの追加
新しいプロバイダーは、軽量アダプターを実装することで追加できます。システムは検出に次の命名規則を使用します。
# プロバイダー/openai_provider.py
クラス OpenaiProvider (BaseProvider):
...
この規則により、一貫性が確保され、新しい統合の自動ロードが可能になります。
貢献は大歓迎です。投稿ガイドを確認し、Discord に参加してディスカッションしてください。
MIT ライセンスに基づいてリリースされており、商用および非商用を問わず無料で使用できます。
複数の Generative AI プロバイダーへのシンプルで統合されたインターフェイス
読み込み中にエラーが発生しました。このページをリロードしてください。
1.5k
フォーク
レポートリポジトリ
リリース
4
オープンコワーカー 0.1.1
最新の
2026 年 6 月 11 日
+ 3 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私のペルソナを共有しないでください

情報

## Original Extract

Simple, unified interface to multiple Generative AI providers - GitHub - andrewyng/aisuite: Simple, unified interface to multiple Generative AI providers

GitHub - andrewyng/aisuite: Simple, unified interface to multiple Generative AI providers · GitHub
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
andrewyng
/
aisuite
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
373 Commits 373 Commits .github/ workflows .github/ workflows aisuite-js aisuite-js aisuite aisuite cli/ py/ aisuite-code-cli cli/ py/ aisuite-code-cli docs docs examples examples guides guides platform platform scripts scripts tests tests viewer-ui viewer-ui .env.sample .env.sample .git-blame-ignore-revs .git-blame-ignore-revs .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README-alternative.md README-alternative.md README.md README.md poetry.lock poetry.lock pyproject.toml pyproject.toml View all files Repository files navigation
An AI agent that lives on your desktop, built on aisuite.
OpenCoworker is a desktop AI agent that can not only chat, but also do deep research and carry out tasks for
you on your computer. It can read files (with permission) to gain context, read/send messages (slack, email, etc.),
and create real deliverables like PDF reports, documents, spreadsheets. It also supports scheduled automations,
such as providing you a daily news summary.
Requires bringing your own API key (OpenAI, Anthropic, Google) or run fully local with Ollama. Your data stays on your machine.
⬇ Download for macOS
macOS 13+ (Apple Silicon)
⬇ Download for Windows
Windows 10/11 (x64) ·
Quickstart: — install, connect a model, first tasks, automations.
Its source lives in this repository under platform/ — a working reference for building your own agent harness on aisuite.
aisuite is a lightweight Python library for building with LLMs, in two layers: a unified Chat Completions API across providers, and an Agents API with tools and toolkits on top. This repo is also home to OpenCoworker , a desktop AI coworker built using aisuite:
┌───────────────────────────────────────────────┐
│ OpenCoworker │ agent harness for doing everyday tasks
├───────────────────────────────────────────────┤
│ Agents API · Toolkits · MCP │ build agents across multiple LLMs
├───────────────────────────────────────────────┤
│ Chat Completions API │ one API across multiple LLM providers
├────────┬───────────┬────────┬────────┬────────┤
│ OpenAI │ Anthropic │ Google │ Ollama │ Others │
└────────┴───────────┴────────┴────────┴────────┘
Chat Completions API — a unified, OpenAI-style interface for OpenAI, Anthropic, Google, Mistral, Hugging Face, AWS, Cohere, Ollama, OpenRouter , and more. Swap providers by changing one string.
Agents API · Toolkits · MCP — give models real Python functions as tools, run multi-turn loops, attach ready-made toolkits (files, git, shell) or any MCP server, and govern it all with tool policies.
OpenCoworker — a desktop AI coworker built using aisuite, shipped as an app for everyday tasks.
Install the base package, or include the SDKs of the providers you plan to use:
pip install aisuite # base package, no provider SDKs
pip install ' aisuite[anthropic] ' # with a specific provider's SDK
pip install ' aisuite[all] ' # with all provider SDKs
You'll also need API keys for the providers you call — the Chat Completions quickstart covers key setup and your first calls.
The OpenCoworker app (desktop)
Download the installer and bring your own API key (or run local models with Ollama):
⬇ macOS (Apple Silicon) · ⬇ Windows 10/11 (x64) · OpenCoworker quickstart
Chat Completions — one API across providers
The chat API provides a high-level abstraction for model interactions. It supports all core parameters ( temperature , max_tokens , tools , etc.) in a provider-agnostic way, and standardizes request and response structures so you can focus on logic rather than SDK differences.
Model names use the format <provider>:<model-name> ; aisuite routes the call to the right provider with the right parameters:
import aisuite as ai
client = ai . Client ()
models = [ "openai:gpt-4o" , "anthropic:claude-3-5-sonnet-20240620" ]
messages = [
{ "role" : "system" , "content" : "Respond in Pirate English." },
{ "role" : "user" , "content" : "Tell me a joke." },
]
for model in models :
response = client . chat . completions . create (
model = model ,
messages = messages ,
temperature = 0.75
)
print ( response . choices [ 0 ]. message . content )
→ Quickstart: docs/chat-completions-quickstart.md — install, key setup, local models, and more examples.
Agents — give models real tools
aisuite turns tool calling into a one-liner: pass plain Python functions and it generates the schemas, executes the calls, and feeds results back to the model.
def will_it_rain ( location : str , time_of_day : str ):
"""Check if it will rain in a location at a given time today.
Args:
location (str): Name of the city
time_of_day (str): Time of the day in HH:MM format.
"""
return "YES"
client = ai . Client ()
response = client . chat . completions . create (
model = "openai:gpt-4o" ,
messages = [{
"role" : "user" ,
"content" : "I live in San Francisco. Can you check for weather "
"and plan an outdoor picnic for me at 2pm?"
}],
tools = [ will_it_rain ],
max_turns = 2 # Maximum number of back-and-forth tool calls
)
print ( response . choices [ 0 ]. message . content )
With max_turns set, aisuite sends your message, executes any tool calls the model requests, returns the results to the model, and repeats until the conversation completes. response.choices[0].intermediate_messages carries the full tool interaction history if you want to continue the conversation.
Prefer full manual control? Omit max_turns and pass OpenAI-format JSON tool specs — aisuite returns the model's tool-call requests and you run the loop yourself. See examples/tool_calling_abstraction.ipynb for both styles.
For longer-running, structured work there is a first-class Agents API: declare an agent once, run it with a Runner , and attach toolkits — prebuilt, sandboxed tool families for files, git, and shell:
import aisuite as ai
from aisuite import Agent , Runner
agent = Agent (
name = "repo-helper" ,
model = "anthropic:claude-sonnet-4-6" ,
instructions = "You are a careful repo assistant. Use your tools to answer from the code." ,
tools = [ * ai . toolkits . files ( root = "." ), * ai . toolkits . git ( root = "." )],
)
result = Runner . run ( agent , "What changed in the last commit? Summarize in 3 bullets." )
print ( result . final_output )
The Agents API also gives you the pieces a production harness needs:
Tool policies — RequireApprovalPolicy , allow/deny lists, or your own callable deciding which tool calls run.
State stores — persist and resume runs (in-memory, file, or Postgres) and continue conversations across processes.
Artifacts & tracing — capture what an agent produced and every step it took along the way.
aisuite natively supports the Model Context Protocol , so any MCP server's tools can be handed to a model without boilerplate ( pip install 'aisuite[mcp]' ):
client = ai . Client ()
response = client . chat . completions . create (
model = "openai:gpt-4o" ,
messages = [{ "role" : "user" , "content" : "List the files in the current directory" }],
tools = [{
"type" : "mcp" ,
"name" : "filesystem" ,
"command" : "npx" ,
"args" : [ "-y" , "@modelcontextprotocol/server-filesystem" , "/path/to/directory" ]
}],
max_turns = 3
)
print ( response . choices [ 0 ]. message . content )
For reusable connections, security filters, and tool prefixing, use the explicit MCPClient .
→ Quickstart: docs/agents-quickstart.md — manual tool handling, the full Agents API, policies, state stores, and MCP in depth.
Extending aisuite: Adding a Provider
New providers can be added by implementing a lightweight adapter. The system uses a naming convention for discovery:
# providers/openai_provider.py
class OpenaiProvider ( BaseProvider ):
...
This convention ensures consistency and enables automatic loading of new integrations.
Contributions are welcome. Please review the Contributing Guide and join our Discord for discussions.
Released under the MIT License — free for commercial and non-commercial use.
Simple, unified interface to multiple Generative AI providers
There was an error while loading. Please reload this page .
1.5k
forks
Report repository
Releases
4
OpenCoworker 0.1.1
Latest
Jun 11, 2026
+ 3 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
