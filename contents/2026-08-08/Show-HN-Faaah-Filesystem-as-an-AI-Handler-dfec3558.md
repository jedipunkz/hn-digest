---
source: "https://github.com/sebastiancarlos/faaah"
hn_url: "https://news.ycombinator.com/item?id=49226789"
title: "Show HN: Faaah – Filesystem as an AI Handler"
article_title: "GitHub - sebastiancarlos/faaah: FAAAH (Filesystem As An AI Handler) - Reuse your AI Agent subscription via text files. · GitHub"
author: "httbs"
captured_at: "2026-08-08T23:18:43Z"
capture_tool: "hn-digest"
hn_id: 49226789
score: 1
comments: 0
posted_at: "2026-08-08T23:13:46Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Faaah – Filesystem as an AI Handler

- HN: [49226789](https://news.ycombinator.com/item?id=49226789)
- Source: [github.com](https://github.com/sebastiancarlos/faaah)
- Score: 1
- Comments: 0
- Posted: 2026-08-08T23:13:46Z

## Translation

タイトル: HN を表示: Faaah – AI ハンドラーとしてのファイルシステム
記事のタイトル: GitHub - sebastiancarlos/faaah: FAAAH (AI ハンドラーとしてのファイルシステム) - テキスト ファイル経由で AI エージェント サブスクリプションを再利用します。 · GitHub
説明: FAAAH (AI ハンドラーとしてのファイルシステム) - テキスト ファイル経由で AI エージェント サブスクリプションを再利用します。 - セバスチャンカルロス/ファーア

記事本文:
GitHub - sebastiancarlos/faaah: FAAAH (AI ハンドラーとしてのファイルシステム) - テキスト ファイル経由で AI エージェント サブスクリプションを再利用します。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
セバスチャンカルロス
/
ふぁあ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
12 コミット 12 コミット docs docs 例 例 src/ faaah src/ faaah .gitignore .gitignore ライセンス ライセンス README.md README.md pyproject.toml pyproject.toml test.py テスト。

py uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
🗣️ FAAAH (AI ハンドラーとしてのファイルシステム)
これまでに必要となる最もシンプルな OpenAI 互換 LLM プロキシ™️
FAAAH を使用すると、AI Agent サブスクリプションを汎用サブスクリプションとして再利用できます。
OpenAI互換のローカルサーバー。
FAAAH は依存関係がなく、プレーンテキスト ファイル プロトコルとして実装されています
(UNIX 哲学認定済み) :
プロンプトをクラウド LLM API に送信する代わりに、FAAAH に送信します。
これは、OpenAI 互換のリクエストを読み取り、次のようにフォルダーにダンプします。
.txt ファイル。
次に、既存の AI コーディング エージェント (Claude Code、opencode など) に次のように指示します。
ファイルを読み取り、他の .txt ファイルに応答を書き込みます。
FAAAH は応答を OpenAI 互換の JSON にパッケージ化して返します。
それらをアプリに追加します。
なぜなら、あなたはすでに AI コーディング アシスタントの料金を支払っているからです。 API キーの支払いをやめる
週末のサイドプロジェクトにぴったりです！ふぁあ！
このツールをローカルで非営利的に使用することは禁止されていると理解しています。
AI エージェント プロバイダーの既存の ToS を破ります。
しかし、もし反対する弁護士がいたら、私にメッセージを送ってください。それから私は紹介します
あなたを私の友人、バーブラ・ストライサンドさんに伝えます。
FAAAH を使用して大規模なデータセットを処理する場合は注意してください。一部のプロバイダー
（どれかはご存知でしょうが）彼らの作品をカフカ風に解釈するかもしれません。
あいまいな ToS を使用し、違反を検出するために Orwellian テレメトリを導入します (
まだ私に起こったことです、YOLO!)
依存関係なし: Python の http.server を使用します。それでおしまい。
308 行のコード: この中にある他のツールの肥大化を見たことがありますか?
スペース？やあ。
Unix の哲学: すべてはファイルです。一つのことをしっかりやる。そのままキスして、
やあ、ヤグニ。
ユニバーサル互換性: ツールが事実上の OpenAI API をサポートしている場合
フォーマット (GraphRAG、LangChain、LlamaIndex、LiteLLM、openai SDK)、
FAAAHをサポートします。
エージェントに依存しない: エージェントのエントリポイントがプロンプトであるため、

そうじゃない
特定のエージェント プロバイダー/バージョンに関連付けられます。将来も安心。
人間参加型フォールバック: AI エージェントがスタックした場合、または使用量に達した場合
制限がある場合は、文字通り現在の応答ファイルを開くことができます (たとえば、
response-0004.txt )、回答を自分で入力します (または、リクエストをコピーして貼り付けます)
お気に入りの Web チャットボット) を選択し、[保存] をクリックします。 FAAAHは成功するでしょう。
# このリポジトリ内から
UVツールをインストールします。 # `faaah` コマンドを PATH にインストールします
または、Git リポジトリから直接:
uv ツールをインストール git+https://github.com/sebastiancarlos/faaah
1. サーバーを起動します
faaah # 127.0.0.1:8000 でリッスン、キュー ~/.cache/faaah/queue
faaah --port 8080 # ポートを上書きします
faaah --queue /tmp/q # キューディレクトリを上書きします
2. エージェントにキューを指示する
エージェントのプロンプトは起動時に出力されます。もう一度取得するには:
faaah --エージェントメッセージ
これをコーディング エージェントに貼り付けると、FAAAH コーディネーターが起動します。
ループ:
faaah --watch を呼び出して次のリクエストを取得します (リクエストが存在するまでブロックされます)。
リクエストをワーカー サブエージェントに委任します (蓄積を防ぐため)
コンテキスト)。
繰り返す。ワーカーが応答ファイルを残さない場合、faaah --watch は単に返します。
同じパスを再度実行するため、コーディネーターは再試行します。
注: FAAAH はサブエージェントを使用して、複数のコンテキストの枯渇を防ぎます。
リクエスト。したがって、AI エージェントは、サブエージェントの作成をサポートする必要があります。
プロンプトによるリクエスト。
たとえば次のように、curl を使用できます。
カール http://127.0.0.1:8000/v1/chat/completions \
-H " Content-Type: application/json " \
-H " 権限: ベアラー何でも " \
-d ' {
"モデル": "ふぁあ",
"messages": [{"role": "user", "content": "俳句を書く。"}]
} '
または OpenAI-API 形式のクライアント:
openaiインポートからOpenAI
client = OpenAI (base_url = "http://127.0.0.1:8000/v1" 、api_key = "任意")
応答 = クライアント 。チャット 。完成品。作成(
model = "faaah" , # このフィールドはとにかく無視されます
メートル

essages = [{ "role" : "user" , "content" : "俳句を書きます。" }]、
timeout = None 、# 個のエージェントが遅くなる可能性があります
）。選択肢 [ 0 ]。メッセージ 。内容
印刷（応答）
依存関係のないサンプルは、examples/chat.py にあります。
FAAAH を介して完全に駆動される GraphRAG の高度な使用法については、を参照してください。
グラフラグファーア 。
使用法: faaah [-h] [--host HOST] [--port PORT] [--queue QUEUE] [--timeout タイムアウト] [--agent-message] [--watch]
AI ハンドラーとしてのファイルシステム: テキスト ファイルを処理する AI エージェントによってサポートされる OpenAI 互換プロキシ。
オプション:
-h、--help このヘルプ メッセージを表示して終了します
--host バインドするホスト アドレス (デフォルト: 127.0.0.1)。
--port PORT リッスンするポート (デフォルト: 8000)。
--queue QUEUE プロンプト/レスポンス ファイルが存在するディレクトリ (デフォルト: ~/.cache/faaah/queue)。
--timeout TIMEOUT N 秒後に各呼び出しを中止します。 0 (デフォルト) は永久に待機します。
--agent-message エージェントのプロンプトのみを出力して終了します (起動時にも出力されます)。
--watch 保留中のプロンプトが存在するまでブロックし、そのパスを出力します。
プロトコル (「ファイルシステム API」)
このプロトコルは、キュー ディレクトリ (~/.cache/faaah/queue) 上のファイルに依存します。
デフォルト）。
各リクエストは、prompt-<id>.txt ファイルを生成します。最初のリクエストの ID は、
00001から単調増加します。
FAAAH は、エージェント (または実際のもの) が対応するメッセージを生成することを期待します。
応答-<id>.txt 。
サブエージェントのワーカーは、最初のパスを次のように記述するように求められます。
response-<id>.txt.draft 、名前を変更する前に修正する可能性があります。
最終応答-<id>.txt は、final と見なされます。
再試行は自動的に行われます。応答ファイルのないプロンプトは、単純に再提供されます。
faaah -- コーディネーターが現れるまで監視してください。
FAAAH (AI ハンドラーとしてのファイルシステム) - テキスト ファイル経由で AI エージェント サブスクリプションを再利用します。
Readme MIT ライセンス アクティビティ スター
フォーク数 0 レポート リポジトリの寄稿者
© 2026 GitHub, Inc.
フッターナビガ

ション
私の個人情報を共有しないでください

## Original Extract

FAAAH (Filesystem As An AI Handler) - Reuse your AI Agent subscription via text files. - sebastiancarlos/faaah

GitHub - sebastiancarlos/faaah: FAAAH (Filesystem As An AI Handler) - Reuse your AI Agent subscription via text files. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
sebastiancarlos
/
faaah
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
12 Commits 12 Commits docs docs examples examples src/ faaah src/ faaah .gitignore .gitignore LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml test.py test.py uv.lock uv.lock View all files Repository files navigation
🗣️ FAAAH (Filesystem As An AI Handler)
The simplest OpenAI-compatible LLM proxy you will ever need.™️
FAAAH allows you to reuse your AI Agent subscription as a generic
OpenAI-compatible local server.
FAAAH is dependency-free, implemented as a plain-text file protocol
(UNIX-philosophy certified) :
Instead of sending your prompts to cloud LLM APIs, send them to FAAAH,
which reads OpenAI-compatible requests and dumps them into a folder as
.txt files.
Then, tell your existing AI coding agent (Claude Code, opencode, etc) to
read the files and write responses to other .txt files.
FAAAH then packages the responses into OpenAI-compatible JSON, and returns
them to your app.
Because you already pay for an AI coding assistant. Stop paying for API keys
just for your weekend side projects! faaah!
It is my understanding that local, non-commercial use of this tool doesn't
break the existing ToS of any AI agent provider.
But if any lawyer disagrees, kindly send me a message. I would then introduce
you to a friend of mine: Miss Barbra Streisand.
Be cautious about using FAAAH to process massive datasets. Some providers
(you know which ones) might do some Kafkaesque interpretations of their
ambiguous ToS, and deploy Orwellian telemetry to detect infractions (hasn't
happened to me yet, YOLO!)
Zero Dependencies: Uses Python's http.server . That's it.
308 lines of code: Have you seen the bloat of other tools in this
space? Yuck.
Unix Philosophy: Everything is a file. Do one thing well. Keep it KISS,
ya YAGNI.
Universal Compatibility: If a tool supports the de-facto OpenAI API
format (GraphRAG, LangChain, LlamaIndex, LiteLLM, the openai SDK), it
supports FAAAH.
Agent Agnostic: Due to the agent entrypoint being a prompt , it's not
tied to any specific agent provider/version. Future-proof.
Human-in-the-loop Fallback: If the AI agent gets stuck or hits usage
limits, you can literally open the current response file (say
response-0004.txt ), type the answer yourself (or copy-paste the request to
your favorite web chatbot), and hit save. FAAAH will succeed.
# from inside this repo
uv tool install . # installs the `faaah` command on PATH
Or straight from the git repository:
uv tool install git+https://github.com/sebastiancarlos/faaah
1. Start the server
faaah # listens on 127.0.0.1:8000, queue ~/.cache/faaah/queue
faaah --port 8080 # override port
faaah --queue /tmp/q # override queue directory
2. Point your agent at the queue
The agent prompt is printed on startup. To grab it again:
faaah --agent-message
Paste it into your coding agent, which then starts a FAAAH coordinator
loop :
Call faaah --watch to obtain the next request (blocks until one exists).
Delegate the request to a worker subagent (to prevent accumulating
context).
Repeat. If a worker leaves no response file, faaah --watch simply returns
the same path again, so the coordinator retries it.
Note: FAAAH uses subagents to prevent exhaustion of context on multiple
requests. Thereby, your AI Agents must support creation of subagents on
request by prompt .
You can use curl , for example:
curl http://127.0.0.1:8000/v1/chat/completions \
-H " Content-Type: application/json " \
-H " Authorization: Bearer anything " \
-d ' {
"model": "faaah",
"messages": [{"role": "user", "content": "Write a haiku."}]
} '
Or any OpenAI-API shaped client:
from openai import OpenAI
client = OpenAI ( base_url = "http://127.0.0.1:8000/v1" , api_key = "anything" )
response = client . chat . completions . create (
model = "faaah" , # this field is ignored anyway
messages = [{ "role" : "user" , "content" : "Write a haiku." }],
timeout = None , # agents can be slow
). choices [ 0 ]. message . content
print ( response )
A dependency-free example lives in examples/chat.py .
For an advanced usage, GraphRAG fully driven through FAAAH , see
graphrag-faaah .
usage: faaah [-h] [--host HOST] [--port PORT] [--queue QUEUE] [--timeout TIMEOUT] [--agent-message] [--watch]
Filesystem As An AI Handler: an OpenAI-compatible proxy backed by an AI agent working over text files.
options:
-h, --help show this help message and exit
--host HOST Address to bind (default: 127.0.0.1).
--port PORT Port to listen on (default: 8000).
--queue QUEUE Directory where prompt/response files live (default: ~/.cache/faaah/queue).
--timeout TIMEOUT Abort each call after N seconds. 0 (default) waits forever.
--agent-message Print ONLY the agent prompt and exit (it's also printed on launch).
--watch Block until a pending prompt exists, print its path.
Protocol (The "Filesystem API")
The protocol relies on files on the queue directory ( ~/.cache/faaah/queue by
default).
Each request produces a prompt-<id>.txt file, where the first one's ID will be
00001 and increase monotonically.
FAAAH then expects the agent (or anything really) to generate a corresponding
response-<id>.txt .
The subagent workers are prompted to write a first pass as
response-<id>.txt.draft , which they may revise, before renaming it to the
final response-<id>.txt they consider final .
Retry is automatic: a prompt with no response file is simply re-offered by
faaah --watch to the coordinator until one appears.
FAAAH (Filesystem As An AI Handler) - Reuse your AI Agent subscription via text files.
Readme MIT license Activity Stars
0 forks Report repository Contributors
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
