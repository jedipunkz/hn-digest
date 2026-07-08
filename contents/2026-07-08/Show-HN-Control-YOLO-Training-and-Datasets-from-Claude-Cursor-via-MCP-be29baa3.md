---
source: "https://github.com/amanharshx/ultralytics-mcp"
hn_url: "https://news.ycombinator.com/item?id=48830826"
title: "Show HN: Control YOLO Training and Datasets from Claude/Cursor via MCP"
article_title: "GitHub - amanharshx/ultralytics-mcp: MCP for Ultralytics Platform workflows, datasets, training, prediction, and model operations. · GitHub"
author: "amanharshx"
captured_at: "2026-07-08T12:15:14Z"
capture_tool: "hn-digest"
hn_id: 48830826
score: 3
comments: 0
posted_at: "2026-07-08T12:04:13Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Control YOLO Training and Datasets from Claude/Cursor via MCP

- HN: [48830826](https://news.ycombinator.com/item?id=48830826)
- Source: [github.com](https://github.com/amanharshx/ultralytics-mcp)
- Score: 3
- Comments: 0
- Posted: 2026-07-08T12:04:13Z

## Translation

タイトル: HN を表示: MCP を介してクロード/カーソルから YOLO トレーニングとデータセットを制御
記事のタイトル: GitHub - amanharshx/ultralytics-mcp: Ultralytics Platform ワークフロー、データセット、トレーニング、予測、モデル操作用の MCP。 · GitHub
説明: Ultralytics Platform ワークフロー、データセット、トレーニング、予測、モデル操作用の MCP。 - amanharshx/ultralytics-mcp

記事本文:
GitHub - amanharshx/ultralytics-mcp: Ultralytics Platform ワークフロー、データセット、トレーニング、予測、モデル操作用の MCP。 · GitHub
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
アマンハルシュクス
/
ウルトラリティクス-mcp
公共
通知
通知設定を変更するにはサインインする必要があります

ngs
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
69 コミット 69 コミット .github/ workflows .github/ workflows アセット アセット フィクスチャ/ パリティ フィクスチャ/ パリティ スクリプト スクリプト src src テスト テスト .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンスREADME.md README.md SECURITY.md SECURITY.md TOOLS.md TOOLS.md biome.json biome.json package-lock.json package-lock.json package.json package.json server.json server.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Ultralytics プラットフォーム用の MCP サーバー
ワークフロー: プロジェクト、データセット、モデル、トレーニング、予測、エクスポート、
データセットのアップロード。
独立したコミュニティプロジェクト。 Ultralytics と提携または承認されていません。
Ultralytics-mcp-demo-1080p.mp4
目次
ローカルビデオファイルからデータセットをアップロードするための PATH 上の ffmpeg および ffprobe
Claude Code、Codex、または標準入出力サーバーを起動できる別の MCP クライアント
Ultralytics プラットフォームにサインインします。
キーの作成、使用法、取り消しの詳細については、公式の Ultralytics API キー ドキュメントを参照してください。
変数
必須
説明
ULTRALYTICS_API_KEY
✅
Ultralytics API キー。予期される形式: ul_ の後に 40 個の 16 進文字が続く
ULTRALYTICS_API_BASE
❌
詳細: API ベース URL をオーバーライドします。デフォルト: https://platform.ultralytics.com/api
ULTRALYTICS_API_KEY をベアラー トークンとして扱います。 MCP クライアント経由で渡します。
環境設定のみ。実際のキーをプロンプト、スクリプト、またはプロンプトに貼り付けないでください。
コミットされた設定ファイル。プロジェクト スコープの .mcp.json ファイルはこのリポジトリでは無視されます
偶発的なキーコミットを減らすため。キーが公開された場合は、Ultralytics でキーを取り消します
プラットフォームを選択し、代替品を作成します。
JSON stdio サーバー定義を受け入れる多くの MCP クライアントで動作します。
{
"mcpサーバー":

{
"ウルトラリティクス" : {
"コマンド" : " npx " ,
"args" : [ " -y " 、 " Ultralytics-mcp@latest " ]、
"環境" : {
"ULTRALYTICS_API_KEY" : " ul_your_api_key_here "
}
}
}
}
反重力
Antigravity 設定を介して追加するか、構成ファイルを更新して追加します。
{
"mcpサーバー": {
"ウルトラリティクス" : {
"コマンド" : " npx " ,
"args" : [ " -y " 、 " Ultralytics-mcp@latest " ]、
"環境" : {
"ULTRALYTICS_API_KEY" : " ul_your_api_key_here "
}
}
}
}
クロード・コード
Claude Code CLI を使用してサーバーを追加します。
クロード mcp add Ultralytics --env ULTRALYTICS_API_KEY=ul_your_api_key_here -- npx -y Ultralytics-mcp@latest
または、プロジェクト スコープのサーバーを repo-root .mcp.json に追加します。
{
"mcpサーバー": {
"ウルトラリティクス" : {
"コマンド" : " npx " ,
"args" : [ " -y " 、 " Ultralytics-mcp@latest " ]、
"環境" : {
"ULTRALYTICS_API_KEY" : " ul_your_api_key_here "
}
}
}
}
クロードデスクトップ
MCP インストール ガイドに従い、上記の標準構成を使用します。
codex mcp add Ultralytics --env ULTRALYTICS_API_KEY=ul_your_api_key_here -- npx -y Ultralytics-mcp@latest
または、 ~/.codex/config.toml に直接追加します。
[ mcp_servers .ウルトラリティクス]
コマンド = " npx "
args = [ " -y " , " Ultralytics-mcp@latest " ]
[ mcp_servers .ウルトラリティクス。環境]
ULTRALYTICS_API_KEY = " ul_your_api_key_here "
カーソル
ボタンをクリックしてインストールします。
重要
インストール ボタンはプレースホルダー キーを書き込みます。インストール後、Cursor MCP 構成を開き、ul_your_api_key_here を Ultralytics API キーに置き換えて、Cursor を再起動します。
カーソル設定 -> MCP -> 新しい MCP サーバーの追加 (または ~/.cursor/mcp.json を直接編集) に移動し、上記の標準構成を使用します: コマンドは npx に設定され、引数は ["-y", "ultralytics-mcp@latest"] に設定され、 env には ULTRALYTICS_API_KEY が設定されます。
MCP インストール ガイドに従い、上記の標準構成を使用します。
重要
インストール ボタンはプレースホルダー キーを書き込みます。インストールしたら、開いてください

VS Code MCP 構成を開き、ul_your_api_key_here を Ultralytics API キーに置き換えて、VS Code を再起動します。
MCP インストール ガイドに従い、上記の標準構成を使用します。 VS Code CLI を使用してサーバーをインストールすることもできます。
# VSコードの場合
コード --add-mcp ' {"name":"ultralytics","command":"npx","args":["-y","ultralytics-mcp@latest"],"env":{"ULTRALYTICS_API_KEY":"ul_your_api_key_here"}} '
インストール後、Ultralytics MCP サーバーは VS Code の GitHub Copilot エージェントで使用できるようになります。
これらの例は、最新の公開された npm リリースを追跡します。アップグレード後に MCP クライアントまたはセッションを再起動し、新しいサーバー プロセスが最新のパッケージを選択するようにします。
クロードMCPリスト
構成された MCP サーバーに Ultralytics が表示されるはずです。
コーデックス mcp リスト
構成された MCP サーバーに Ultralytics が表示されるはずです。
プロジェクト、データセット、モデル、エクスポート、GPU の可用性を参照する
プロジェクトとデータセットの参照を id、slug、username/slug、または ul:// で解決します。
Ultralytics Explore で公開プロジェクトとデータセットを検索
データセット取り込みジョブを開始し、アーカイブ ファイル、フォルダー、またはビデオをアップロードします
トレーニングの進行状況を監視し、最新のメトリクスまたは最近のメトリクス履歴を検査します。
画像 URL または Base64 入力からモデル予測を実行する
モデルの重みをローカル パスにダウンロードする
コストを明示的に確認してエクスポートおよびトレーニング ジョブを作成する
training_start.train_args を通じて高度な YOLO トレーニング設定を渡します
既存のプロジェクト モデルまたは公式 YOLO ベース チェックポイントからトレーニングを開始
パラメータの完全なリファレンス、安全上の注意、ローカル パスの動作、扱いにくいツールの例については、TOOLS.md を参照してください。
ULTRALYTICS_API_KEY はベアラー トークンです。 MCP クライアント環境を介して渡し、実際のキーを決してコミットしない
export_create にはconfirm_cost が必要です: true
training_start にはconfirm_cost が必要です: true
あいまいなプロジェクトまたはデータセット参照は推測ではなく失敗します
署名付きのアップロードとダウンロード

oad URL が認証を転送しない
ローカル アップロード ツールは、MCP クライアント ホストからファイルを読み取ります。 Ultralytics と共有することが予想されるパスの呼び出しのみを承認する
model_download は、要求されたローカル パスに書き込みます。承認する前に、output_path を確認して上書きしてください
ラベル付きコピーを再アップロードすると、既存の画像にラベルが添付されるのではなく、データセット画像レコードが複製される場合があります。
画像のみのデータセットのアップロードは、classify として推論される場合があります。タスクの保存が重要な場合はラベルを含める
ULTRALYTICS_API_KEY は ul_ で始まり、ちょうど 40 16 進数を含む必要があります
プレフィックスの後の文字。
サーバーがクロードコードにロードされない
npx と Node.js がインストールされていることを確認します
サーバーを追加するときに ULTRALYTICS_API_KEY が --env で渡されたことを確認します
必要に応じて、claude mcp get Ultralytics を使用してサーバー構成を検査します。
npx と Node.js がインストールされていることを確認します
~/.codex/config.toml または codex mcp add コマンドの ULTRALYTICS_API_KEY 値を確認してください
ULTRALYTICS_API_KEY=ul_your_api_key_here npx -y Ultralytics-mcp@latest
コマンドが構成エラーですぐに終了する場合は、まず環境を修正してください。
認証、レート制限、またはエンドポイントの動作については、
公式 Ultralytics プラットフォーム REST API ドキュメント 。
サポートを求める場合は、ツール名、リクエストの概要、対応ステータス、
編集されたレスポンスボディと最小限の複製。実際の API は含めないでください
キー、署名付き URL、プライベート データセットのコンテンツ、またはプライベート モデル アーティファクト。
npmインストール
npm実行チェック
npmテスト
npm ビルドを実行する
npm run generated:ツール
について
Ultralytics Platform のワークフロー、データセット、トレーニング、予測、モデル操作用の MCP。
www.npmjs.com/package/ultralytics-mcp
トピックス
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
10
v0.1.9
最新の
2026 年 7 月 8 日
+ 9 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください

年 。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

MCP for Ultralytics Platform workflows, datasets, training, prediction, and model operations. - amanharshx/ultralytics-mcp

GitHub - amanharshx/ultralytics-mcp: MCP for Ultralytics Platform workflows, datasets, training, prediction, and model operations. · GitHub
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
amanharshx
/
ultralytics-mcp
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
69 Commits 69 Commits .github/ workflows .github/ workflows assets assets fixtures/ parity fixtures/ parity scripts scripts src src tests tests .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md TOOLS.md TOOLS.md biome.json biome.json package-lock.json package-lock.json package.json package.json server.json server.json tsconfig.json tsconfig.json View all files Repository files navigation
MCP server for Ultralytics Platform
workflows: projects, datasets, models, training, prediction, exports, and
dataset uploads.
Independent community project. Not affiliated with or endorsed by Ultralytics.
ultralytics-mcp-demo-1080p.mp4
Table of Contents
ffmpeg and ffprobe on PATH to upload a dataset from a local video file
Claude Code, Codex, or another MCP client that can launch stdio servers
Sign in at Ultralytics Platform .
See the official Ultralytics API key docs for key creation, usage, and revocation details.
Variable
Required
Description
ULTRALYTICS_API_KEY
✅
Ultralytics API key. Expected format: ul_ followed by 40 hex characters
ULTRALYTICS_API_BASE
❌
Advanced: override API base URL. Default: https://platform.ultralytics.com/api
Treat ULTRALYTICS_API_KEY as a bearer token. Pass it through your MCP client's
environment configuration only. Never paste real keys into prompts, scripts, or
committed config files. Project-scoped .mcp.json files are ignored by this repo
to reduce accidental key commits; if a key is exposed, revoke it in Ultralytics
Platform and create a replacement.
Works in many MCP clients that accept JSON stdio server definitions.
{
"mcpServers" : {
"ultralytics" : {
"command" : " npx " ,
"args" : [ " -y " , " ultralytics-mcp@latest " ],
"env" : {
"ULTRALYTICS_API_KEY" : " ul_your_api_key_here "
}
}
}
}
Antigravity
Add via the Antigravity settings or by updating your configuration file:
{
"mcpServers" : {
"ultralytics" : {
"command" : " npx " ,
"args" : [ " -y " , " ultralytics-mcp@latest " ],
"env" : {
"ULTRALYTICS_API_KEY" : " ul_your_api_key_here "
}
}
}
}
Claude Code
Add server with Claude Code CLI:
claude mcp add ultralytics --env ULTRALYTICS_API_KEY=ul_your_api_key_here -- npx -y ultralytics-mcp@latest
Or add a project-scoped server in repo-root .mcp.json :
{
"mcpServers" : {
"ultralytics" : {
"command" : " npx " ,
"args" : [ " -y " , " ultralytics-mcp@latest " ],
"env" : {
"ULTRALYTICS_API_KEY" : " ul_your_api_key_here "
}
}
}
}
Claude Desktop
Follow the MCP install guide , use the standard config above.
codex mcp add ultralytics --env ULTRALYTICS_API_KEY=ul_your_api_key_here -- npx -y ultralytics-mcp@latest
Or add it directly to ~/.codex/config.toml :
[ mcp_servers . ultralytics ]
command = " npx "
args = [ " -y " , " ultralytics-mcp@latest " ]
[ mcp_servers . ultralytics . env ]
ULTRALYTICS_API_KEY = " ul_your_api_key_here "
Cursor
Click the button to install:
Important
The install button writes a placeholder key. After installing, open your Cursor MCP config and replace ul_your_api_key_here with your Ultralytics API key, then restart Cursor.
Go to Cursor Settings -> MCP -> Add new MCP Server (or edit ~/.cursor/mcp.json directly) and use the standard config above: command set to npx , args set to ["-y", "ultralytics-mcp@latest"] , and ULTRALYTICS_API_KEY in env .
Follow the MCP install guide , use the standard config above.
Important
The install button writes a placeholder key. After installing, open your VS Code MCP config and replace ul_your_api_key_here with your Ultralytics API key, then restart VS Code.
Follow the MCP install guide , use the standard config above. You can also install the server using the VS Code CLI:
# For VS Code
code --add-mcp ' {"name":"ultralytics","command":"npx","args":["-y","ultralytics-mcp@latest"],"env":{"ULTRALYTICS_API_KEY":"ul_your_api_key_here"}} '
After installation, the Ultralytics MCP server will be available for use with your GitHub Copilot agent in VS Code.
These examples track latest published npm release. Restart MCP client or session after upgrading so new server process picks up latest package.
claude mcp list
You should see ultralytics in configured MCP servers.
codex mcp list
You should see ultralytics in configured MCP servers.
Browse projects, datasets, models, exports, GPU availability
Resolve project and dataset refs by id, slug, username/slug , or ul://
Search public projects and datasets on Ultralytics Explore
Start dataset ingest jobs and upload archive files, folders, or videos
Monitor training progress and inspect latest metrics or recent metric history
Run model prediction from image URL or base64 input
Download model weights to local path
Create exports and training jobs with explicit cost confirmation
Pass advanced YOLO training settings through training_start.train_args
Start training from existing project models or official YOLO base checkpoints
See TOOLS.md for full parameter reference, safety notes, local-path behavior, and examples for tricky tools.
ULTRALYTICS_API_KEY is a bearer token; pass it via MCP client env and never commit real keys
export_create requires confirm_cost: true
training_start requires confirm_cost: true
Ambiguous project or dataset refs fail instead of guessing
Signed upload and download URLs do not forward Authorization
Local upload tools read files from the MCP client host; approve calls only for paths you expect to share with Ultralytics
model_download writes to the requested local path; review output_path and overwrite before approving
Re-uploading labeled copies may duplicate dataset image records instead of attaching labels to existing images
Images-only dataset uploads may be inferred as classify ; include labels when task preservation matters
ULTRALYTICS_API_KEY must start with ul_ and contain exactly 40 hex
characters after prefix.
Server not loading in Claude Code
verify npx and Node.js are installed
verify ULTRALYTICS_API_KEY was passed with --env when adding server
if needed, inspect server config with claude mcp get ultralytics
verify npx and Node.js are installed
verify ULTRALYTICS_API_KEY value in ~/.codex/config.toml or codex mcp add command
ULTRALYTICS_API_KEY=ul_your_api_key_here npx -y ultralytics-mcp@latest
If command exits immediately with config error, fix environment first.
For authentication, rate-limit, or endpoint behavior, compare against the
official Ultralytics Platform REST API docs .
When asking for help, include the tool name, request summary, response status,
redacted response body, and a minimal reproduction. Do not include real API
keys, signed URLs, private dataset contents, or private model artifacts.
npm install
npm run check
npm test
npm run build
npm run generate:tools
About
MCP for Ultralytics Platform workflows, datasets, training, prediction, and model operations.
www.npmjs.com/package/ultralytics-mcp
Topics
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
10
v0.1.9
Latest
Jul 8, 2026
+ 9 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
