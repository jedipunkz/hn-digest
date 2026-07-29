---
source: "https://github.com/vedaant00/opendot"
hn_url: "https://news.ycombinator.com/item?id=49100984"
title: "Opendot: A terminal AI agent that snapshots every action so you can undo it"
article_title: "GitHub - vedaant00/opendot: A terminal AI agent you can fully undo - every file and shell action is snapshotted and reversible. Model-agnostic (any LLM), and connects to 1000+ app tools and MCP servers. · GitHub"
author: "vedaant00"
captured_at: "2026-07-29T18:58:19Z"
capture_tool: "hn-digest"
hn_id: 49100984
score: 1
comments: 0
posted_at: "2026-07-29T18:11:00Z"
tags:
  - hacker-news
  - translated
---

# Opendot: A terminal AI agent that snapshots every action so you can undo it

- HN: [49100984](https://news.ycombinator.com/item?id=49100984)
- Source: [github.com](https://github.com/vedaant00/opendot)
- Score: 1
- Comments: 0
- Posted: 2026-07-29T18:11:00Z

## Translation

Title: Opendot: A terminal AI agent that snapshots every action so you can undo it
Article title: GitHub - vedaant00/opendot: A terminal AI agent you can fully undo - every file and shell action is snapshotted and reversible. Model-agnostic (any LLM), and connects to 1000+ app tools and MCP servers. · GitHub
Description: A terminal AI agent you can fully undo - every file and shell action is snapshotted and reversible. Model-agnostic (any LLM), and connects to 1000+ app tools and MCP servers. - vedaant00/オープンドット

記事本文:
GitHub - vedaant00/opendot: 完全に元に戻すことができるターミナル AI エージェント - すべてのファイルとシェルのアクションはスナップショットが作成され、元に戻すことができます。モデルに依存せず (LLM を問わず)、1000 以上のアプリ ツールと MCP サーバーに接続します。 · GitHub
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
{{ メッサ

げ }}
ヴェダント00
/
オープンドット
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
46 コミット 46 コミット .github .github アセット アセット src/ opendot src/ opendot テスト テスト .gitignore .gitignore COTRIBUTING.md COTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml ビューすべてのファイル リポジトリ ファイルのナビゲーション
完全に元に戻すことができる対話型ターミナル AI エージェント。
opendot は実際のファイルとシェルで直接動作しますが、他のターミナルとは異なります
エージェントは、実行されるすべてのアクションが最初にスナップショットされるため、正確に確認できます。
それが何をしたか、そしてそれをきれいに元に戻します。ファイルとシェルコマンドだけでなく、
リポジトリ内の編集。効果がワークスペースから外れるコマンド (ネットワーク、sudo、
git Push 、作業ディレクトリ外の削除) にフラグが立てられ、前に確認されます。
彼らは、取り返しのつかないことについて正直にメモを残して走ります。
それが opendot のポイントです。何もしないので放っておいてもよいエージェントです
それは驚くべきことであり、（ほとんど）取り返しのつかないことは何もありません。
opendot はモデルに依存しません。LiteLLM (OpenAI、
Anthropic、Google、DeepSeek など)、Ollama 経由で完全にローカルで実行されます。オラマはただ
ゼロセットアップローカルオプション。好みのバックエンドを使用してください。
# インストールせずにすぐに試してください
uvxオープンドット
# 推奨 (分離されたグローバル CLI)
uv ツール install opendot # または: pipx install opendot
# も動作します
pip インストール opendot
使用する
opendot # インタラクティブなチャットを開く
opendot -p " このプロジェクトを要約する " # ワンショット、スクリプト / CI 用
opendot --model claude-opus-4-5 # 特定のモデルで起動します (以下を参照)
opendot log # Audit: エージェントはここで何をしましたか?
opendot undo # 最後の操作を元に戻します
opendot 元に戻す 000004 # 残り

アクション #4 の前にワークスペースを取得します
チャット内で、スラッシュ コマンド: /model (検索可能なモデル ピッカー)、
/provider (プロバイダーに接続 + API キーを貼り付け)、 /log 、 /undo 、 /clear 、
/compact 、 /help 。
クラウド、ローカル、または顔を抱き締めるなど、どのモデルでも機能します。 API キーが必要です。
使用するプロバイダー (opendot は BYO キーです。モデルをホストしません)。選択してください
モデルを作成し、 /model と /provider を使用してチャット内にキーを貼り付けるか、または
環境にキーを設定し、 --model を渡します。
プロバイダー名はキーの取得場所にリンクしています。
推論モデルは思考をライブでストリーミングします。
ローカルの OpenAI 互換サーバー (llama.cpp / llama-server 、vLLM、LM Studio):
--api-base と openai/ プレフィックス付きモデルを使用して、opendot をサーバーにポイントします。
# 例: llama.cpp: llama-server -m model.gguf --port 8080
opendot --model openai/local --api-base http://localhost:8080/v1
どのモデルが動作するか。デフォルトは gpt-5.1 です。そのキー ( OPENAI_API_KEY ) の場合
は設定されていませんが、別のプロバイダーのキーが設定されている場合、opendot は自動的にそれに切り替えます
起動時のプロバイダー — 例: DEEPSEEK_API_KEY のみが設定されている、裸のオープンドット
deepseek/deepseek-chat を使用します。プロバイダーキーが見つからない場合、opendot が起動します
それは問題ありませんが、最初のメッセージには、キーを設定するか /provider を実行するためのヒントが表示されます（むしろ、
生のプロバイダーエラーよりも）。 ollam/* モデルにはキーは必要ありません。ローカルの Ollama だけです。
opendot は MCP クライアントです: 任意の MCP に接続します
サーバーとそのツールは、組み込みのものと同様にエージェントで使用できるようになります。
/mcp (サーバーとサーバーのドロップダウン) を使用してチャット内からそれらを管理します。
ステータス (「➕ サーバーの追加」)、またはコマンドラインから:
# stdio サーバー — 起動コマンドを `--` の後に置きます
opendot mcp add < name > --env KEY=VALUE -- < command > [args...]
# リモートサーバー (http/sse)
opendot mcp add < 名前 > --url < https url >
# 認証が必要なリモートサーバー

— HTTPヘッダーを渡す
opendot mcp 追加 supabase \
--url " https://mcp.supabase.com/mcp?project_ref=<id>&read_only=true " \
--header " Authorization=Bearer <your-supabase-access-token> "
opendot mcp list # 構成されたサーバーを表示
opendot mcp delete < name > # 1 つ削除します
サーバーは ~/.opendot/mcp.json に保存され、自動的に接続されます。
次の打ち上げ。接続されているサーバーがサイドバーに表示されます。認証済みリモートの場合
サーバーでは、opendot はヘッダー/トークン メソッドをサポートしています (例: Supabase のアクセス)
トークン) — インタラクティブなブラウザー OAuth フローはまだ実装されていません。
opendot は外部ツールが何を行うのかを知ることができないため、すべての MCP ツール呼び出しは
不可逆的なものとして扱われます — 実行前に確認され、
台帳。組み込みのファイル/シェル アクションはスナップショットを保持し、通常どおり元に戻すことができます。
MCP を超えて、opendot は Composio の 1000 以上のアプリに接続できます
独自の Composio API を使用したツール (Gmail、Slack、GitHub、Notion、Linear など)
鍵。チャットで /composio を使用するだけです。
最初の /composio は、Composio API キー (次の場所に保存されています) を要求します。
~/.opendot/composio.json 、所有者のみ読み取り可能)。
その後、/composio は利用可能なアプリをリストします。必要に応じて 1 つ選択してください
OAuth、opendot はブラウザを開いて認証し、完了するのを待ちます。
直接/API キー コネクタはすぐにアクティブになります。
有効なアプリはサイドバーに表示されます。彼らのツールは次回の起動時にロードされます。
Composio ツールは外部サービスにアクセスできるため、MCP と同様に、すべての呼び出しが処理されます。
不可逆的として: 最初に確認され、台帳で ✗ とマークされます。
OPENDOT.md をプロジェクトにドロップします。その散文はエージェントに次のように与えられます。
コンテキスト。 opendot ブロックを使用してスナップショットを取得するものを制御することもできます。
``` オープンドット
# 通常はスキップされる場合でも、これらのスナップショットを作成します。
スナップショット: 距離
# これらのスナップショットは決して作成しないでください:
スキップ: データ、*.log
「」
デフォルトでは、opendot は

kips .git 、node_modules 、virtualenvs、およびビルドキャッシュ
スナップショットを作成するとき — ルールはどちらの方向でもルールをオーバーライドします。
すべてのファイル書き込みまたはシェルコマンドの前に、opendot は作業中のスナップショットを作成します。
ディレクトリを ~/.opendot のコンテンツ アドレス ストアにコピーします (それぞれ固有のファイル
一度保存されるため、スナップショットは安価です)。
すべてのアクションは追加専用台帳に記録され、次の方法で検査できます。
オープンドットログ。
opendot undo は、ワークスペースを選択した時点に正確に復元します。
保守的な分類子により、どのシェル コマンドがワークスペースであるかが決定されます。
包含 (自動実行、取り消し可能) とエスケープ (最初に確認、マーク付き)
不可逆的）。不明な場合は質問します。
正直な境界: opendot では、マシンから離れたエフェクト (送信されたエフェクト) を元に戻すことはできません。
電子メール、削除されたリモート データベース、git Push )。走る前に教えてくれる
そうでないふりをするのではなく、それらを。
問題や PR は歓迎です — セットアップと詳細については CONTRIBUTING.md を参照してください。
唯一の厳しいルール (可逆性を壊さないこと)。セキュリティレポートが通過する
セキュリティ.md 。
git clone https://github.com/vedaant00/opendot
cd オープンドット
uv pip install -e " .[dev] " # または: pip install -e ".[dev]"
pytest
ステータス
初期（アルファ）。インタラクティブ エージェント、ローカル ツール、および完全な可逆性
エンジンは動作し、テストされています。ストリーミング、スラッシュコマンド、OPENDOT.md ルール
より充実した TUI とより多くのツールが登場します。
完全に元に戻すことができるターミナル AI エージェント - すべてのファイルとシェルのアクションはスナップショットを取得され、元に戻すことができます。モデルに依存せず (LLM を問わず)、1000 以上のアプリ ツールと MCP サーバーに接続します。
pypi.org/project/opendot トピック
Readme MIT ライセンス
セキュリティポリシー アクティビティスター
3 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A terminal AI agent you can fully undo - every file and shell action is snapshotted and reversible. Model-agnostic (any LLM), and connects to 1000+ app tools and MCP servers. - vedaant00/opendot

GitHub - vedaant00/opendot: A terminal AI agent you can fully undo - every file and shell action is snapshotted and reversible. Model-agnostic (any LLM), and connects to 1000+ app tools and MCP servers. · GitHub
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
vedaant00
/
opendot
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
46 Commits 46 Commits .github .github assets assets src/ opendot src/ opendot tests tests .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml View all files Repository files navigation
An interactive terminal AI agent you can fully undo.
opendot works directly on your real files and shell — but unlike other terminal
agents, every action it takes is snapshotted first , so you can see exactly
what it did and cleanly walk it back. Files and shell commands, not just
in-repo edits. Commands whose effects escape your workspace (network, sudo,
git push , deleting outside the working dir) are flagged and confirmed before
they run, with an honest note about what can't be undone.
That's the point of opendot: an agent you can let loose because nothing it does
is a surprise, and (almost) nothing is irreversible.
opendot is model-agnostic — it works with any model through LiteLLM (OpenAI,
Anthropic, Google, DeepSeek, …) and runs fully local via Ollama. Ollama is just
the zero-setup local option; use whatever backend you prefer.
# try it instantly, no install
uvx opendot
# recommended (isolated global CLI)
uv tool install opendot # or: pipx install opendot
# also works
pip install opendot
Use
opendot # open an interactive chat
opendot -p " summarize this project " # one-shot, for scripts / CI
opendot --model claude-opus-4-5 # launch with a specific model (see below)
opendot log # audit: what has the agent done here?
opendot undo # revert the last action
opendot undo 000004 # restore the workspace to before action #4
Inside the chat, slash-commands: /model (searchable model picker),
/provider (connect a provider + paste an API key), /log , /undo , /clear ,
/compact , /help .
Any model works — cloud, local, or Hugging Face. You need an API key for the
provider you want to use (opendot is BYO-key; it doesn't host models). Pick a
model and paste a key right inside the chat with /model and /provider , or
set the key in your environment and pass --model :
Provider names link to where you get a key.
Reasoning models stream their thinking live.
Local OpenAI-compatible servers (llama.cpp / llama-server , vLLM, LM Studio):
point opendot at the server with --api-base and an openai/ -prefixed model.
# e.g. llama.cpp: llama-server -m model.gguf --port 8080
opendot --model openai/local --api-base http://localhost:8080/v1
Which model runs. The default is gpt-5.1 . If its key ( OPENAI_API_KEY )
isn't set but another provider's key is, opendot automatically switches to that
provider on launch — e.g. with only DEEPSEEK_API_KEY set, a bare opendot
uses deepseek/deepseek-chat . If no provider key is found, opendot starts
fine but the first message shows a hint to set a key or run /provider (rather
than a raw provider error). ollama/* models need no key — just a local Ollama.
opendot is an MCP client: connect any MCP
server and its tools become available to the agent alongside the built-in ones.
Manage them from inside the chat with /mcp (a dropdown of your servers and
their status, with "➕ Add a server"), or from the command line:
# a stdio server — put its launch command after `--`
opendot mcp add < name > --env KEY=VALUE -- < command > [args...]
# a remote server (http/sse)
opendot mcp add < name > --url < https url >
# a remote server that needs auth — pass an HTTP header
opendot mcp add supabase \
--url " https://mcp.supabase.com/mcp?project_ref=<id>&read_only=true " \
--header " Authorization=Bearer <your-supabase-access-token> "
opendot mcp list # show configured servers
opendot mcp remove < name > # remove one
Servers are stored in ~/.opendot/mcp.json and connect automatically on the
next launch; connected servers appear in the sidebar. For authenticated remote
servers, opendot supports the header/token method (e.g. Supabase's access
token) — the interactive browser-OAuth flow is not implemented yet.
Because opendot can't know what an external tool does, every MCP tool call is
treated as irreversible — it's confirmed before running and marked ✗ in the
ledger. Your built-in file/shell actions stay snapshotted and undoable as usual.
Beyond MCP, opendot can connect to Composio 's 1000+ app
tools (Gmail, Slack, GitHub, Notion, Linear, …) using your own Composio API
key. Just use /composio in the chat:
The first /composio asks for your Composio API key (stored in
~/.opendot/composio.json , owner-readable only).
After that, /composio lists the available apps. Pick one — if it needs
OAuth, opendot opens your browser to authorize and waits for you to finish;
direct/API-key connectors activate immediately.
Enabled apps appear in the sidebar; their tools load on the next launch.
Composio tools reach external services, so — like MCP — every call is treated
as irreversible : confirmed first, marked ✗ in the ledger.
Drop an OPENDOT.md in your project. Its prose is given to the agent as
context. You can also control what gets snapshotted with an opendot block:
``` opendot
# snapshot these even though they'd normally be skipped:
snapshot: dist
# never snapshot these:
skip: data, *.log
```
By default opendot skips .git , node_modules , virtualenvs, and build caches
when snapshotting — your rules override those in either direction.
Before every file write or shell command, opendot snapshots the working
directory into a content-addressed store in ~/.opendot (each unique file
stored once, so snapshots are cheap).
Every action is recorded in an append-only ledger you can inspect with
opendot log .
opendot undo restores the workspace to a chosen point, exactly.
A conservative classifier decides which shell commands are workspace-
contained (auto-run, undoable) vs. escaping (confirmed first, marked
irreversible). When unsure, it asks.
Honest boundary: opendot cannot undo effects that leave your machine (a sent
email, a dropped remote database, a git push ). It tells you before running
those, rather than pretending otherwise.
Issues and PRs welcome — see CONTRIBUTING.md for setup and
the one hard rule (don't break reversibility). Security reports go through
SECURITY.md .
git clone https://github.com/vedaant00/opendot
cd opendot
uv pip install -e " .[dev] " # or: pip install -e ".[dev]"
pytest
Status
Early (alpha). The interactive agent, local tools, and the full reversibility
engine work and are tested. Streaming, slash-commands, and OPENDOT.md rules
are in. A richer TUI and more tools are coming.
A terminal AI agent you can fully undo - every file and shell action is snapshotted and reversible. Model-agnostic (any LLM), and connects to 1000+ app tools and MCP servers.
pypi.org/project/opendot Topics
Readme MIT license Contributing
Security policy Activity Stars
3 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
