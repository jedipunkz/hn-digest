---
source: "https://github.com/hit9/nanocode"
hn_url: "https://news.ycombinator.com/item?id=48421523"
title: "Show HN: Nanocode-CLI – A lightweight terminal-based AI coding assistant"
article_title: "GitHub - hit9/nanocode: A lightweight terminal-based AI coding assistant · GitHub"
author: "hit9"
captured_at: "2026-06-06T07:11:21Z"
capture_tool: "hn-digest"
hn_id: 48421523
score: 2
comments: 0
posted_at: "2026-06-06T04:54:05Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Nanocode-CLI – A lightweight terminal-based AI coding assistant

- HN: [48421523](https://news.ycombinator.com/item?id=48421523)
- Source: [github.com](https://github.com/hit9/nanocode)
- Score: 2
- Comments: 0
- Posted: 2026-06-06T04:54:05Z

## Translation

タイトル: Show HN: Nanocode-CLI – 軽量のターミナルベースの AI コーディング アシスタント
記事タイトル: GitHub - hit9/nanocode: 軽量のターミナルベースの AI コーディング アシスタント · GitHub
説明: 軽量のターミナルベースの AI コーディング アシスタント。 GitHub でアカウントを作成して、hit9/nanocode の開発に貢献してください。

記事本文:
GitHub - hit9/nanocode: 軽量のターミナルベースの AI コーディング アシスタント · GitHub
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
ヒット9
/
ナノコード
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動

コード 「その他のアクション」メニューを開く フォルダーとファイル
689 コミット 689 コミット .github/ workflows .github/ workflows スナップショット スナップショット テスト テスト .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE MANIFEST.in MANIFEST.in Makefile Makefile README.md README.md README.zh-CN.md README.zh-CN.md nanocode.py nanocode.py pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Python で書かれた小さなターミナル コーディング エージェント。
nanocode は 1.0 より前のソフトウェアです。コマンド、構成、およびツールの動作は、安定版リリース前に変更される可能性があります。
ライブ ターン コントロール: 現在のツール フローを失うことなく、エージェントがまだ動作している間にフォローアップ入力を追加します。
ファイル状態脳 : 読み取りと編集により、現在重要なファイルの現在の行番号付きビューが構築されます。
古い編集の保護: line:hash アンカーは、ターゲット コードがずれている場合に編集を拒否します。
プロジェクト対応ナビゲーション : シンボル インデックスを使用して、アウトライン、参照、変更されたファイルをすばやく移動します。
回復可能なコンテキスト: ツールの出力はプロンプト内で制限されたままですが、生の tr.N 結果は呼び出し可能なままです。
キャッシュを意識したコンテキスト: プロンプトキャッシュの再利用を改善するために、安定したセクションは早めに残り、ノイズの多い動作状態は遅く残ります。
集中的な作業メモリ: メモは、目標、計画、既知の事実をノイズの多い実行ログから分離します。
ターミナルファーストのワークフロー: モデルの選択、履歴の検索、確認、ライブ コマンド出力、追加された入力、ステータスはすべて 1 つの CLI に留まります。
UVツールをインストールnanocode-cli
アップグレード:
UV ツールのアップグレード nanocode-cli
ローカル開発の場合:
UV 同期 --extra dev
UV ナノコードを実行する
使用法
--config <パス> : TOML 構成ファイルを使用します。
--init-config : デフォルトの構成ファイルを作成します。
--yolo : 変異ツールの確認をスキップします。
-v 、 --version : バージョンを表示します。
実行中のターン中、+> プロンプトは次の入力を受け入れます。

xtモデルのリクエスト。
/help : コマンドとツールを表示します。
/status : 実行時のステータスを表示します。
/api [auto|chat|anthropic] : プロバイダー API 形式を表示または設定します。
/debug [on|off] : モデル I/O デバッグ トレースを切り替えます。
/compact : コンテキストをコンパクトにしました。
/index [force] : コード シンボル インデックスを同期または再構築します。
/provider [名前] : プロバイダーを表示または設定します。
/model [MODEL] : モデルを表示または設定します。
/reason : 推論の努力を選択します。
/set KEY VALUE : プロバイダー/ランタイム値を設定します。
/yolo : ツールの確認を切り替えます。
対話型セレクターは、 j / k 、矢印、/ 検索、Enter、および Esc をサポートします。入力は、履歴、補完、Ctrl-R 履歴検索をサポートしています。
ファイル: Read 、 LineCount 、 List 、 Find 、 Search 。
編集: ファイルのコンテンツを作成またはパッチします。
Read 、 Search 、および InspectCode は、必要に応じて行アンカーを返します。編集では、現在の行:ハッシュ アンカーを使用して、古い編集を拒否します。
nanocode --init-config
デフォルトの構成の場所は ~/.nanocode/config.toml です。
[provider.<name>] : url 、 key 、 model 、 api 、 prompt_cache_key 、 available_models 、reasoning 、 chat_reasoning 、 pressure 、 timeout
[ランタイム] shell_timeout 、 max_agent_steps 、 max_context_tokens 、 yolo
api = "auto" は、プロバイダー/モデル プロファイルを使用して、チャット完了と人為的なメッセージのどちらかを選択します。プロンプト_キャッシュ_キー = "自動" は、プロバイダー、モデル、ワークスペース、およびツールのスキーマ名から安定したキーを派生します。
次のプロバイダーは nanocode でテストされています。
aliyun : チャット補完による Alibaba Cloud (Tongyi Qianwen) API
llama.cpp : llama.cpp サーバー経由のローカル推論
各モデル リクエストは、明示的なメッセージから手動で構築されます。安定したコンテキストが最初にあり、会話はメッセージとして残り、作業メモリが続き、最新のファイル状態が最後に追加されます。
モデルリクエスト
+------------------------------------------------+
|システム |
|簡潔なエージェント契約とツールのルール |
+----

------------------------------------------+
|ユーザー |
|環境 |
+------------------------------------------------+
|ユーザー/アシスタント |
|会話、要約、ツール |
+------------------------------------------------+
|ユーザー |
|記憶: 目標、計画、既知の日付 |
+------------------------------------------------+
|ユーザー |
|ファイル状態: 最新の読み取り/編集ファイル ビュー |
+------------------------------------------------+
基本的なルール:
ターン途中のアシスタント テキストと追加のユーザー入力は会話として保持されます。
コンテキストが大きくなりすぎると、以前の会話は明示的な要約に圧縮されます。
ファイル状態は、正常に実行された読み取りおよび編集ツールによって更新され、現在リストされているファイル範囲が最近のファイルから順に表示されます。
新しいファイル行は古い行を上書きします。無効化を編集すると、古い範囲がクリアされます。
ファイル行は、表示される前に現在のファイル統計または行ハッシュと照合してチェックされます。
成功した読み取りおよび編集ツールのメッセージは、ファイル本体を繰り返すのではなく、ファイル状態を示します。
他のツールの出力は会話メッセージ内に制限されており、 tr.N によって呼び出すことができます。
nanocode は、起動された環境でファイルを編集したり、シェル コマンドを実行したりできます。サンドボックス保護は提供されません。必要に応じて、独自のサンドボックス、コンテナ、VM、またはその他の隔離された環境内で実行します。
軽量のターミナルベースの AI コーディング アシスタント
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A lightweight terminal-based AI coding assistant. Contribute to hit9/nanocode development by creating an account on GitHub.

GitHub - hit9/nanocode: A lightweight terminal-based AI coding assistant · GitHub
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
hit9
/
nanocode
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
689 Commits 689 Commits .github/ workflows .github/ workflows snapshots snapshots tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE MANIFEST.in MANIFEST.in Makefile Makefile README.md README.md README.zh-CN.md README.zh-CN.md nanocode.py nanocode.py pyproject.toml pyproject.toml View all files Repository files navigation
A small terminal coding agent written in Python.
nanocode is pre-1.0 software. Commands, configuration, and tool behavior may change before a stable release.
Live turn control : Add follow-up input while the agent is still working, without losing the current tool flow.
File-state brain : Reads and edits build a current, line-numbered view of the files that matter now.
Stale-edit protection : line:hash anchors reject edits when the target code has drifted.
Project-aware navigation : Use the symbol index to jump through outlines, references, and changed files quickly.
Recoverable context : Tool output stays bounded in the prompt, while raw tr.N results remain recallable.
Cache-aware context : Stable sections stay early and noisy working state stays late to improve prompt-cache reuse.
Focused working memory : Note separates goal, plan, and known facts from noisy execution logs.
Terminal-first workflow : Model selection, history search, confirmations, live command output, appended input, and status all stay in one CLI.
uv tool install nanocode-cli
Upgrade:
uv tool upgrade nanocode-cli
For local development:
uv sync --extra dev
uv run nanocode
Usage
--config <path> : use a TOML config file.
--init-config : create a default config file.
--yolo : skip confirmations for mutating tools.
-v , --version : show the version.
During a running turn, the +> prompt accepts follow-up input for the next model request.
/help : show commands and tools.
/status : show runtime status.
/api [auto|chat|anthropic] : show or set provider API format.
/debug [on|off] : toggle model I/O debug traces.
/compact : compact context now.
/index [force] : sync or rebuild the code symbol index.
/provider [NAME] : show or set provider.
/model [MODEL] : show or set model.
/reason : choose reasoning effort.
/set KEY VALUE : set provider/runtime values.
/yolo : toggle tool confirmations.
Interactive selectors support j / k , arrows, / search, Enter, and Esc. Input supports history, completion, and Ctrl-R history search.
File: Read , LineCount , List , Find , Search .
Edit: Edit creates or patches file content.
Read , Search , and InspectCode return line anchors where useful. Edit uses current line:hash anchors to reject stale edits.
nanocode --init-config
Default config location is ~/.nanocode/config.toml .
[provider.<name>] : url , key , model , api , prompt_cache_key , available_models , reasoning , chat_reasoning , temperature , timeout
[runtime] shell_timeout , max_agent_steps , max_context_tokens , yolo
api = "auto" chooses between Chat Completions and Anthropic Messages using provider/model profiles. prompt_cache_key = "auto" derives a stable key from provider, model, workspace, and tool schema names.
The following providers have been tested with nanocode:
aliyun : Alibaba Cloud (Tongyi Qianwen) API via Chat Completions
llama.cpp : Local inference via llama.cpp server
Each model request is built manually from explicit messages. Stable context comes first, conversation stays as messages, working memory follows, and the latest file state is appended at the end.
model request
+--------------------------------------------------+
| system |
| concise agent contract and tool rules |
+--------------------------------------------------+
| user |
| Environment |
+--------------------------------------------------+
| user/assistant |
| conversation, compacted summaries, tools |
+--------------------------------------------------+
| user |
| Memory: goal, plan, known, date |
+--------------------------------------------------+
| user |
| FILE STATE: latest Read/Edit file view |
+--------------------------------------------------+
Core rules:
Mid-turn assistant text and appended user input are kept as conversation.
Earlier conversation is compacted into an explicit summary when the context grows too large.
FILE STATE is updated by successful Read and Edit tools and shows current listed file ranges, with recent files first.
Newer file lines overwrite older lines; edit invalidations clear stale ranges.
File lines are checked against current file stat or line hash before being shown.
Successful Read and Edit tool messages point to FILE STATE instead of repeating file bodies.
Other tool outputs are bounded in conversation messages and can be recalled by tr.N .
nanocode can edit files and run shell commands in the environment where it is started. It does not provide sandbox protection. Run it inside your own sandbox, container, VM, or other isolated environment when needed.
A lightweight terminal-based AI coding assistant
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
