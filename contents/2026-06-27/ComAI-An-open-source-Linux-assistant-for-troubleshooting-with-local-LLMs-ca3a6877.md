---
source: "https://github.com/hossbit/comai-linux-assistant"
hn_url: "https://news.ycombinator.com/item?id=48693507"
title: "ComAI – An open-source Linux assistant for troubleshooting with local LLMs"
article_title: "GitHub - hossbit/comai-linux-assistant: Local AI Linux terminal assistant written in Bash. Explain commands, analyze files and logs, and use local LLMs or OpenAI-compatible APIs. · GitHub"
author: "mirhacker"
captured_at: "2026-06-27T00:31:06Z"
capture_tool: "hn-digest"
hn_id: 48693507
score: 2
comments: 0
posted_at: "2026-06-26T23:46:50Z"
tags:
  - hacker-news
  - translated
---

# ComAI – An open-source Linux assistant for troubleshooting with local LLMs

- HN: [48693507](https://news.ycombinator.com/item?id=48693507)
- Source: [github.com](https://github.com/hossbit/comai-linux-assistant)
- Score: 2
- Comments: 0
- Posted: 2026-06-26T23:46:50Z

## Translation

タイトル: ComAI – ローカル LLM を使用したトラブルシューティングのためのオープンソース Linux アシスタント
記事のタイトル: GitHub - hossbit/comai-linux-assistant: Bash で書かれたローカル AI Linux ターミナル アシスタント。コマンドの説明、ファイルとログの分析、ローカル LLM または OpenAI 互換 API の使用。 · GitHub
説明: Bash で書かれたローカル AI Linux ターミナル アシスタント。コマンドの説明、ファイルとログの分析、ローカル LLM または OpenAI 互換 API の使用。 - hossbit/comai-linux-assistant

記事本文:
GitHub - hossbit/comai-linux-assistant: Bash で書かれたローカル AI Linux ターミナル アシスタント。コマンドの説明、ファイルとログの分析、ローカル LLM または OpenAI 互換 API の使用。 · GitHub
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
ホスビット
/
comai-linux-アシスタント
公共
通知

通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
22 コミット 22 コミット bin bin config config lib/ comai lib/ comai scripts scripts LICENSE LICENSE README.md README.md すべてのファイルを表示 リポジトリ ファイル ナビゲーション
ComAI - ローカル AI Linux ターミナル アシスタント
ComAI は、Bash で書かれたオープンソースの AI 搭載 Linux ターミナル アシスタントです。
Linux について質問し、シェル コマンドを説明し、設定ファイルを分析し、
ログにエラーがないか検査し、ローカル LLM、Ollama、または OpenAI と直接対話します。
端末から。
llama.cpp および OpenAI 互換 API を介したローカル AI モデル
ローカル チャットおよびファイル分析用の Ollama モデル
ファイルと構成の分析
ログ分析とエラー検出
Bash フレンドリーなインストールと構成
comai 説明 chmod 755
comai 1GB を超えるファイルを見つけるにはどうすればよいですか?
comai エラーはありますか? -f アプリケーション.ログ
comai ollam このスクリプトについて説明します -f script.sh
comai gpt この nginx 設定を説明します -f nginx.conf
ローカル モードがデフォルトです。 Ollama モードは、comai の後の最初の単語で実行されます。
オラマです。 ChatGPT モードは、最初の単語が gpt または chatgpt の場合に実行されます。
ローカル モードの場合は、最初に hossbit/localai をインストールします。 ComAI は次の時点でそれを期待しています。
～/あい
Ollama または ChatGPT モードのみが必要な場合は、localai をスキップして次を使用できます。
comai ollam ... または comai gpt ... 。
git clone https://github.com/hossbit/comai-linux-assistant.git
cd comai-linux-assistant
chmod +x スクリプト/install.sh
./scripts/install.sh
インストールされたファイルは ~/localcomai に移動します。
インストーラーは、ファイルを変更する前に各セクションについて説明します。どこに行くかを尋ねます
ComAI をインストールします。デフォルトで ~/localcomai を使用します。ファイルが既に存在する場合は表示されます。
新しいデフォルト設定キーを追加しながら、既存の設定値を保持します。
コマイ何

/etc は Linux にありますか?
comai このコマンドの動作 -command " find . -type f -size +100M "
ChatGPT を使用します。
コマイGPTこんにちは
comai gpt 説明 chmod 755
オラマを使用します。
コマイ・オラマ こんにちは
コマイ・オラマがchmod 755を説明する
ファイルを読み取ります:
このファイルについて説明します -f script.sh
comai ollam このファイルを要約します -f README.md
comai gpt このファイルを要約します -f llama-swap.log
簡単なローカル ファイル/ログ チェックを依頼します。
comaiの最新ファイル
comai最大のファイルはここにあります
comai エラーはありますか? -f ラマ-スワップ.ログ
1 つのリクエストに対してモデルを選択します。
comai --model=Qwen2.5-7B-Instruct-Q4_K_M こんにちは
comai ollam --model=qwen2.5-coder:7b こんにちは。
comai gpt --model=gpt-5.5 こんにちは
ChatGPT セットアップ
エクスポート OPENAI_API_KEY= " your_api_key "
または、インストールされた構成にキーを入力します。
openai_api_key : your_api_key
インストールされた構成:
~ /localcomai/config/comai.yaml
実際の API キーを git にコミットしないでください。
config/comai.yaml # ソースのデフォルト
~ /localcomai/config/comai.yaml # インストールされた構成
例:
プロバイダー : ローカル
ai_dir : ~/ai
api_base_url : http://127.0.0.1
api_base_port : 11435
モデル: Qwen2.5-Coder-7B-Instruct-Q4_K_M
gpt_model : gpt-5.5
ollama_api_base : http://127.0.0.1:11434
ollama_model : qwen2.5-coder:7b
openai_api_base : https://api.openai.com
openai_api_key :
最大トークン数 : 420
タイムアウト: 120
ファイル最大バイト数 : 24000
dir_context_max : 120
error_regex : エラー|エラー|失敗|失敗|例外|致命的|パニック|タイムアウト|警告|警告|トレースバック
error_intent_regex : error|errors|failed|failure|warning|warnings|problem|problems|issue|issues|wrong|bad|broken|fail|crash|crashed|panic|timeout|traceback|healthy|health|(^|[[:space:]]])ok([[:space:]]|$)|oky|check (this )?log|scan (this )?ログ
主なキーの意味:
プロバイダー : デフォルトのプロバイダー。 local 、 ollam 、または openai を使用します。
ai_dir : localai がインストールされている場所。デフォルトは ~/ai です。
api_base_url : ローカルの OpenAI 互換 API URL

ポートなしで。
api_base_port : ローカルの OpenAI 互換 API ポート。
モデル: comai hi のデフォルトのローカル モデル。
gpt_model : comai gpt hi のデフォルトの OpenAI モデル。
ollama_api_base : Ollama API URL。デフォルトは http://127.0.0.1:11434 です。
ollama_model : comai ollama こんにちは のデフォルトの Ollama モデル。
openai_api_base : OpenAI API URL。別の互換性のあるサーバーが必要であることがわかっている場合を除き、これを https://api.openai.com のままにしておきます。
openai_api_key : ChatGPT モードの OpenAI キーを保存するオプションの場所。 OPENAI_API_KEY の方が安全であり、これをオーバーライドします。
max_tokens : 回答の最大長。
timeout : リクエストのタイムアウトを秒単位で指定します。
file_max_bytes : 各 -f ファイルから読み取られる最大バイト数。
dir_context_max : コンテキストとして送信される現在のディレクトリ エントリの最大数。
error_regex : ローカル ログ/エラー チェックで使用される単語。
error_intent_regex : 質問がログ/エラー チェックを求めているかどうかを決定するために使用される単語。
COMAI_MODEL=Qwen2.5-7B-Instruct-Q4_K_M こんばんは。
COMAI_PROVIDER=オラマ・コマイ、こんにちは
OPENAI_API_KEY=your_api_key comai gpt こんにちは
COMAI_MAX_TOKENS=120 こんまい、こんにちは
ローカルAI
ComAI は、 ~/ai にローカルの OpenAI 互換サーバーを想定しています。
systemctl --user start comai-localai.service
または手動で:
~ /ai/start.sh
ローカル モデルを確認します。
カール -s http://127.0.0.1:11435/v1/models | jq -r ' .data[].id '
オラマ
コマイ・オラマ こんにちは
comai ollam このファイルを要約します -f README.md
Ollama モデルを確認します。
カール -s http://127.0.0.1:11434/api/tags | jq -r ' .models[].name '
トラブルシューティング
comai gpt ... 429 と表示されます: OpenAI がレート制限またはクォータのリクエストを拒否しました。請求、クレジット、プロジェクト、またはレート制限を確認します。
comai gpt ... キーをエクスポートせずに動作します。おそらく ~/localcomai/config/comai.yaml から openai_api_key を読み取っています。
comai ... ローカル AI にアクセスできません: comai-localai.service を開始するか、 ~/ai/start.sh を実行してください。
comai ollam ... Ollama に到達できません: start Olma

lama を実行し、 config/comai.yaml の ollama_api_base を確認してください。
bashカールjq検索ソートヘッドsed awk grep wc tr readlink systemctl
オプション:
~ /localcomai/scripts/uninstall.sh
これにより ComAI ファイルが削除され、~/ai だけが残ります。
Bash で書かれたローカル AI Linux ターミナル アシスタント。コマンドの説明、ファイルとログの分析、ローカル LLM または OpenAI 互換 API の使用。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
5
ComAI 2.1.2
最新の
2026 年 6 月 26 日
+ 4 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Local AI Linux terminal assistant written in Bash. Explain commands, analyze files and logs, and use local LLMs or OpenAI-compatible APIs. - hossbit/comai-linux-assistant

GitHub - hossbit/comai-linux-assistant: Local AI Linux terminal assistant written in Bash. Explain commands, analyze files and logs, and use local LLMs or OpenAI-compatible APIs. · GitHub
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
hossbit
/
comai-linux-assistant
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
22 Commits 22 Commits bin bin config config lib/ comai lib/ comai scripts scripts LICENSE LICENSE README.md README.md View all files Repository files navigation
ComAI - Local AI Linux Terminal Assistant
ComAI is an open-source AI-powered Linux terminal assistant written in Bash.
Ask Linux questions, explain shell commands, analyze configuration files,
inspect logs for errors, and interact with local LLMs, Ollama, or OpenAI directly
from your terminal.
Local AI models through llama.cpp and OpenAI-compatible APIs
Ollama models for local chat and file analysis
File and configuration analysis
Log analysis and error detection
Bash-friendly installation and configuration
comai explain chmod 755
comai how do I find files larger than 1GB ?
comai do you see any error ? -f application.log
comai ollama explain this script -f script.sh
comai gpt explain this nginx configuration -f nginx.conf
Local mode is the default. Ollama mode runs when the first word after comai
is ollama . ChatGPT mode runs when the first word is gpt or chatgpt .
For local mode, install hossbit/localai first. ComAI expects it at:
~ /ai
If you only want Ollama or ChatGPT mode, you can skip localai and use
comai ollama ... or comai gpt ... .
git clone https://github.com/hossbit/comai-linux-assistant.git
cd comai-linux-assistant
chmod +x scripts/install.sh
./scripts/install.sh
Installed files go to ~/localcomai .
The installer explains each section before it changes files. It asks where to
install ComAI, uses ~/localcomai by default, shows when files already exist,
and preserves existing config values while adding new default config keys.
comai what is /etc in linux ?
comai how this command work -command " find . -type f -size +100M "
Use ChatGPT:
comai gpt hi
comai gpt explain chmod 755
Use Ollama:
comai ollama hi
comai ollama explain chmod 755
Read a file:
comai explain this file -f script.sh
comai ollama summarize this file -f README.md
comai gpt summarize this file -f llama-swap.log
Ask simple local file/log checks:
comai newest file
comai biggest file here
comai do you see any error ? -f llama-swap.log
Choose a model for one request:
comai --model=Qwen2.5-7B-Instruct-Q4_K_M hi
comai ollama --model=qwen2.5-coder:7b hi
comai gpt --model=gpt-5.5 hi
ChatGPT Setup
export OPENAI_API_KEY= " your_api_key "
Or put the key in the installed config:
openai_api_key : your_api_key
Installed config:
~ /localcomai/config/comai.yaml
Do not commit a real API key to git.
config/comai.yaml # source default
~ /localcomai/config/comai.yaml # installed config
Example:
provider : local
ai_dir : ~/ai
api_base_url : http://127.0.0.1
api_base_port : 11435
model : Qwen2.5-Coder-7B-Instruct-Q4_K_M
gpt_model : gpt-5.5
ollama_api_base : http://127.0.0.1:11434
ollama_model : qwen2.5-coder:7b
openai_api_base : https://api.openai.com
openai_api_key :
max_tokens : 420
timeout : 120
file_max_bytes : 24000
dir_context_max : 120
error_regex : error|errors|failed|failure|exception|fatal|panic|timeout|warn|warning|traceback
error_intent_regex : error|errors|failed|failure|warning|warnings|problem|problems|issue|issues|wrong|bad|broken|fail|crash|crashed|panic|timeout|traceback|healthy|health|(^|[[:space:]])ok([[:space:]]|$)|okay|check (this )?log|scan (this )?log
What the main keys mean:
provider : default provider. Use local , ollama , or openai .
ai_dir : where localai is installed. Default is ~/ai .
api_base_url : local OpenAI-compatible API URL without the port.
api_base_port : local OpenAI-compatible API port.
model : default local model for comai hi .
gpt_model : default OpenAI model for comai gpt hi .
ollama_api_base : Ollama API URL. Default is http://127.0.0.1:11434 .
ollama_model : default Ollama model for comai ollama hi .
openai_api_base : OpenAI API URL. Keep this as https://api.openai.com unless you know you need another compatible server.
openai_api_key : optional place to store your OpenAI key for ChatGPT mode. OPENAI_API_KEY is safer and overrides this.
max_tokens : maximum answer length.
timeout : request timeout in seconds.
file_max_bytes : maximum bytes read from each -f file.
dir_context_max : maximum current-directory entries sent as context.
error_regex : words used by local log/error checks.
error_intent_regex : words used to decide whether a question is asking for a log/error check.
COMAI_MODEL=Qwen2.5-7B-Instruct-Q4_K_M comai hi
COMAI_PROVIDER=ollama comai hi
OPENAI_API_KEY=your_api_key comai gpt hi
COMAI_MAX_TOKENS=120 comai hi
Local AI
ComAI expects a local OpenAI-compatible server in ~/ai .
systemctl --user start comai-localai.service
Or manually:
~ /ai/start.sh
Check local models:
curl -s http://127.0.0.1:11435/v1/models | jq -r ' .data[].id '
Ollama
comai ollama hi
comai ollama summarize this file -f README.md
Check Ollama models:
curl -s http://127.0.0.1:11434/api/tags | jq -r ' .models[].name '
Troubleshooting
comai gpt ... says 429 : OpenAI rejected the request for rate limit or quota. Check billing, credits, project, or rate limits.
comai gpt ... works without exporting a key: it is probably reading openai_api_key from ~/localcomai/config/comai.yaml .
comai ... cannot reach local AI: start comai-localai.service or run ~/ai/start.sh .
comai ollama ... cannot reach Ollama: start Ollama and check ollama_api_base in config/comai.yaml .
bash curl jq find sort head sed awk grep wc tr readlink systemctl
Optional:
~ /localcomai/scripts/uninstall.sh
This removes ComAI files and leaves ~/ai alone.
Local AI Linux terminal assistant written in Bash. Explain commands, analyze files and logs, and use local LLMs or OpenAI-compatible APIs.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
5
ComAI 2.1.2
Latest
Jun 26, 2026
+ 4 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
