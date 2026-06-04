---
source: "https://github.com/desplega-ai/claude-bridge"
hn_url: "https://news.ycombinator.com/item?id=48403472"
title: "Claude -p drop-in replacement after Jun 15th"
article_title: "GitHub - desplega-ai/claude-bridge: A simple `claude -p` bridge, drop-in replacement · GitHub"
author: "tarasyarema"
captured_at: "2026-06-04T21:38:02Z"
capture_tool: "hn-digest"
hn_id: 48403472
score: 2
comments: 1
posted_at: "2026-06-04T19:28:09Z"
tags:
  - hacker-news
  - translated
---

# Claude -p drop-in replacement after Jun 15th

- HN: [48403472](https://news.ycombinator.com/item?id=48403472)
- Source: [github.com](https://github.com/desplega-ai/claude-bridge)
- Score: 2
- Comments: 1
- Posted: 2026-06-04T19:28:09Z

## Translation

タイトル: 6 月 15 日以降のクロード -p ドロップイン交換
記事タイトル: GitHub - desplega-ai/claude-bridge: シンプルな `claude -p` ブリッジ、ドロップイン置換 · GitHub
説明: シンプルな `claude -p` ブリッジ、ドロップイン置換。 GitHub でアカウントを作成して、desplega-ai/claude-bridge の開発に貢献してください。

記事本文:
GitHub - desplega-ai/claude-bridge: シンプルな `claude -p` ブリッジ、ドロップイン置換 · GitHub
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
デスプレガアイ
/
クロード・ブリッジ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスターB

牧場 タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
17 コミット 17 コミット .github/ workflows .github/ workflows docs docs scripts scripts src src .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md COTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md bun.lock bun.lock package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
claude-bridge は、一般的な claude -p のブリッジ所有の代替品です。
自動化。
raw claude -p に委任する代わりに、通常の対話型 Claude を開始します。
分離された tmux ペイン内のコードは、
tmux を介してプロンプトを送信し、クロード自身のディスク上のトランスクリプトを追跡し、応答をフォーマットします。
そしてターンエンドで出ます。
これにより、プロンプトディスパッチ、トランスクリプトキャプチャ、 --output-format 、JSON スキーマが維持されます。
検証、およびブリッジ内のプロセス終了動作。
# 生のクロードコード印刷モード
クロード -p " こんにちは " --output-format json
# ブリッジ所有の代替品
bunx @desplega.ai/claude-bridge -p "say hi" --output-format json
トランスクリプトのソースは、Claude が書き込む JSONL ファイルと同じです。
~/.claude/projects/<slug>/<session-uuid>.jsonl 。外部印刷モード、パイプ接続
消費者はブリッジ エンベロープを取得し、TTY ユーザーはコンパクトで読みやすいビューを取得します。の
トランスクリプトの末尾が続きます
シャノンのテクニック: スナップショットを作成する
起動前に既存の *.jsonl を設定し、新しいファイルをポーリングし、
100 ミリ秒ごとにポーリングして再解析します。
オーケストレーターは、ブロックされる可能性があるプロンプトも事前にクリアします。
クロードの UI:
クロードのグローバル設定は編集されているため、
プロジェクト[<workdir>].hasTrustDialogAccepted および
hasCompletedProjectOnboarding が設定されています。これは ~/.claude.json によるものです
デフォルト、または CLAUDE_CONFIG_DIR が指定されている場合は $CLAUDE_CONFIG_DIR/.claude.json
セット。以前のファイルも一緒にバックアップされます

それを次のようにします
.claude.json.claude-bridge-backup 。
workdir ごとの .claude/settings.local.json セット
デフォルトモード：「bypassPermissions」および
SkipDangerousModePermissionPrompt: true 。
クロードは --dangerously-skip-permissions で起動されます。
テーマ/セキュリティの起動プロンプトは監視することで自動的に受け入れられます
マーカー テキストと Enter の送信用の tmux キャプチャ ペイン。と
--desplega-local-auth 、カスタム API キーの確認プロンプトも
自動的に受け入れられます。ログイン方法の選択は意図的に自動受け入れられません。
+----------------------+
|クロードブリッジ |
| - tmux ペースト |
| - トランスクリプト尾 |
+----------+----------+
|
| tmux ペーストバッファ + Enter
v
+------+----------------------------+
| tmux セッション クロードブリッジ-<id> |
|ペイン 0: クロード --dangerously-... |
+-------------------------------------------------+
要件
PATH 上のクロード CLI、バージョン >= 2.1.80
クロード 生成されたクロード プロセスに対して認証されたコード。
# ドロップイン印刷モードの使用法。
bunx @desplega.ai/claude-bridge -p " こんにちは "
bunx @desplega.ai/claude-bridge -p "say hi" --output-format json
bunx @desplega.ai/claude-bridge -p "say hi" --output-format stream-json
# ブリッジ固有のコンシューマー向けに Desplega/ブリッジ エンベロープをオプトインします。
bunx @desplega.ai/claude-bridge -p "say hi" --output-format stream-json --desplega-format
Bun を使用してグローバルにインストールします。
bun install -g @desplega.ai/claude-bridge
クロードブリッジ -p "こんにちは"
クロードブリッジ --ヘルプ
npm を使用してグローバルにインストールします。
npm install -g @desplega.ai/claude-bridge
クロードブリッジ -p "こんにちは"
インストールされているコマンドは claude-bridge です。ランタイムには Bun が引き続き必要です
公開された bin は #!/usr/bin/env bun を使用するためです。
クロードブリッジ -p "こんにちは"
クロードブリッジ -p " こんにちは " --model ソネット
クロードブリッジ -p " こんにちは " --output-format json
クロードブリッジ -p " こんにちは " --output-format stream-json
printf ' こんにちは\n '

|クロードブリッジ --print
印刷モードは、シェルの自動化を目的としており、それ以外の場合は
クロード -p :
クロード -p " こんにちは " --output-format json
クロードブリッジ -p " こんにちは " --output-format json
クロード -p " こんにちは " --output-format stream-json
クロードブリッジ -p " こんにちは " --output-format stream-json
これは、一般的な claude -p 自動化のドロップイン置換として意図されています。
印刷モードでは、ラッパーは tmux で対話型の Claude セッションを開始し、待機します。
ペインの準備ができるようにするには、tmux を介してプロンプトを送信し、
要求された形式を指定すると、tmux セッションが強制終了されます。
デフォルトでは、印刷モードの標準出力は要求されたクロード互換用に予約されています。
出力。ブリッジ エンベロープとブリッジ デバッグ イベントは標準出力に書き込まれません。
--desplega-format を明示的に渡す場合を除き、json または stream-json モード。
claude-bridge は Anthropic API 自体を呼び出しません。ローカルを起動します
クロード CLI を使用し、クロード プロセスが実行できるあらゆる認証に依存します。
使用します。
ローカルの対話型マシンの場合は、まず claude が動作することを確認します。
クロードの認証ステータス
クロード -P 「こんにちは」
次に、ブリッジを実行します。
クロードブリッジ -p "こんにちは"
ヘッドレス CI の場合は、次の長期有効な Claude Code OAuth トークンを使用します。
クロードセットアップトークン
印刷されたとおりに正確に設定します。
エクスポート CLAUDE_CODE_OAUTH_TOKEN=sk-ant-oat01-...
クロードブリッジ -p " こんにちは " --output-format json
デフォルトでは、生成された Claude プロセスは HOME 、 CLAUDE_CONFIG_DIR 、
および CLAUDE_CODE_OAUTH_TOKEN ;次のような Anthropic プロバイダーの環境変数
ANTHROPIC_API_KEY と ANTHROPIC_AUTH_TOKEN がクリアされるため、ブリッジは
誤って別の認証パスをテストしないでください。
意図的に生成されたクロードが必要な場合は、--desplega-local-auth を使用してください。
ローカル認証関連の環境変数を受信するプロセス。クロードがカスタム API を見せた場合
キーの確認プロンプトが表示される場合、このモードでは API キーのパスを選択します。
アンソロピック_A

PI_KEY=... クロードブリッジ --desplega-local-auth -p " こんにちは "
クロードがブラウザ ログインまたはログイン方法セレクタを表示した場合、ブリッジは
それを自動選択します。 claude auth status を実行するか、 claude setup-token を実行するか、またはattach
バナーに表示されている tmux ペインに移動し、プロンプトを手動で完了します。
-p / --print には、プロンプト引数またはパイプされた標準入力が必要です。 --出力形式
印刷モードが必要で、 text 、 json 、または stream-json を受け入れます。デフォルトは
テキスト。 --json-schema も印刷専用です。
互換モードがデフォルトです。スクリプト内で claude -p を置き換える場合は、
--desplega-format を渡さないでください。
最終結果はトランスクリプトから得られます。クロードがシステムを書くとき
turn_duration 行では、ラッパーはその中で確認した最新のアシスタント テキストを使用します。
回ります。
text : 最終的な回答テキストと末尾の改行のみを出力します。ラッパー
エラーは標準エラー出力に送られ、ゼロ以外で終了します。
json : クロード互換の最終的な JSON 結果オブジェクトを 1 つ出力します。
result の回答と、 session_id などの利用可能なトランスクリプト メタデータ
duration_ms と使用量。
stream-json : 生の Claude トランスクリプト JSONL 行を書き込み時にストリーミングします。
ブリッジはそれらをカスタム エンベロープで包みません。
ブリッジが所有する古い JSON エンベロープが必要な場合は、--desplega-format を使用します。
json モードまたは stream-json モード。このフラグはブリッジ固有のコンシューマ向けであり、
ドロップイン claude -p 置換スクリプト:
クロードブリッジ -p " 挨拶 " --output-format stream-json --desplega-format
--desplega-format を使用すると、json にはブリッジ デバッグ メタデータが含まれます。
--desplega-verbose が設定されており、stream-json は改行区切りのブリッジを出力します。
実行の進行に応じてイベントが発生し、その後最終結果イベントが発生します。これは習慣です
クロードのネイティブの stream-json スキーマではなく、claude-bridge イベント ストリームです。
一般的な --desplega-format --output-format stream-json イベント タイプは次のとおりです。
{ "タイプ" : "プッシュ" 、 "id" :

" ab12cd34 " , "content" : " こんにちは " }
{ "タイプ" : " トランスクリプトフォルダー " 、 "パス" : " /Users/.../.claude/projects/... " }
{ "type" : "transcript_open " 、 "path" : " /Users/.../<uuid>.jsonl " 、 "session_id" : " ... " }
{ "タイプ" : " トランスクリプト " 、 "行" :{ "タイプ" : " アシスタント " 、 "メッセージ" :{ ... }}}
{ "type" : " result " 、 "subtype" : " success " 、 "is_error" : false 、 "result" : " こんにちは。 " 、 "session_id" : " ... " }
これらのカスタム トランスクリプト イベントは、 --desplega-format でのみ存在します。で
デフォルトの互換モードでは、同じクロード データがトップレベルとして書き込まれます。
JSONL行。
--json-schema <スキーマ|ファイル> はブリッジが所有しています。 raw には転送されません
クロード -p ;ラッパーは通常の tmux/transcript パスを維持し、スキーマを挿入します
--append-system-prompt を使用したガイダンスは、ファイルから最後の JSON 値を抽出します。
最終的なアシスタント テキストを作成し、Zod を使用してローカルで検証します。
既存のユーザー指定の --append-system-prompt 値は保持されます。とき
スキーマが存在する場合、ラッパーはそれらのプロンプトをそのスキーマ命令とマージします。
それらを交換する代わりに。
スキーマ印刷モードでは、グローバルな Claude Code Stop フックもインストールされます。
~/.claude/settings.json 。フックはクロードブリッジスキーマの外では不活性です
走る。スキーマの実行中に、クロードの前に最終アシスタント テキストをチェックします。
停止し、検証されない場合は停止をブロックします。それはクロードに限界を与える
ラッパーが終了するまでに有効な JSON で応答するための追加ターン数。
そのフックを次のように明示的に制御します。
クロードブリッジ --desplega-install
クロードブリッジ --desplega-uninstall
インストールは追加専用で冪等です。無関係なフックは保持され、古くなります。
古い claude-bridge フック コマンドは現在のコマンドに置き換えられます。
スキーマ引数は、インライン JSON または JSON ファイルへのパスです。
claude-bridge -p " リポジトリ名を返す " \
--json-schema ' {" タイプ

":"オブジェクト","必須":["名前"],"プロパティ":{"名前":{"タイプ":"文字列"}}} ' \
--出力形式 json
claude-bridge -p " リポジトリ名を返す " \
--json-schema ./schema.json \
--出力形式のテキスト
抽出は意図的に単純かつ決定的です。
それ以外の場合は、最後にフェンスされた JSON ブロックを使用します。
それ以外の場合は、応答で最終的なバランスのとれた JSON オブジェクトまたは配列を使用します。
検証には Zod の z.fromJSONSchema() コンバーターが使用されます。その API はまだマークされています
Zod による実験的ですが、ブリッジを Zod の JSON スキーマに合わせて維持します。
ここで手書きのバリデーターを維持する代わりにサポートします。ゾッドができないなら
スキーマを変換すると、ラッパーはそれを印刷モード エラーとして扱います。
--output-format text を指定すると、スキーマ モードが成功すると、抽出された JSON が出力されます。
コンパクトな JSON としての値。 --output-format json を使用すると、最終結果には次のものが含まれます。
result 内の元の返信テキストと一緒に、 Structured_output を追加します。と
--desplega-format 、ブリッジ JSON 結果も含まれます
構造化_出力_ソース 。
クロードが応答した後にスキーマの抽出または検証が失敗した場合、json と
stream-json エラーの結果には、未変更の Claude を含む raw_response が含まれます
返事をする。テキストモードでは、同じ生の応答が次の標準エラー出力に出力されます。
生のクロードの返答: 。
コンパクトな文字列化スキーム

[切り捨てられた]

## Original Extract

A simple `claude -p` bridge, drop-in replacement. Contribute to desplega-ai/claude-bridge development by creating an account on GitHub.

GitHub - desplega-ai/claude-bridge: A simple `claude -p` bridge, drop-in replacement · GitHub
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
desplega-ai
/
claude-bridge
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
17 Commits 17 Commits .github/ workflows .github/ workflows docs docs scripts scripts src src .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md bun.lock bun.lock package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
claude-bridge is a bridge-owned replacement for common claude -p
automation.
Instead of delegating to raw claude -p , it starts normal interactive Claude
Code inside a detached tmux pane, sends your
prompt through tmux, tails Claude's own on-disk transcript, formats the reply,
and exits at turn end.
That keeps prompt dispatch, transcript capture, --output-format , JSON schema
validation, and process exit behavior inside the bridge.
# Raw Claude Code print mode
claude -p " say hi " --output-format json
# Bridge-owned replacement
bunx @desplega.ai/claude-bridge -p " say hi " --output-format json
The transcript source is the same JSONL file Claude writes under
~/.claude/projects/<slug>/<session-uuid>.jsonl . Outside print mode, piped
consumers get bridge envelopes and TTY users get a compact readable view. The
transcript tailing follows the
Shannon technique: snapshot the
pre-existing *.jsonl set before launch, poll for a fresh file, and
poll-and-reparse it every 100 ms.
The orchestrator also pre-clears the prompts that would otherwise block
Claude's UI:
Claude's global config is edited so
projects[<workdir>].hasTrustDialogAccepted and
hasCompletedProjectOnboarding are set. This is ~/.claude.json by
default, or $CLAUDE_CONFIG_DIR/.claude.json when CLAUDE_CONFIG_DIR is
set. The previous file is backed up alongside it as
.claude.json.claude-bridge-backup .
A per-workdir .claude/settings.local.json sets
defaultMode: "bypassPermissions" and
skipDangerousModePermissionPrompt: true .
claude is launched with --dangerously-skip-permissions .
Theme/security startup prompts are auto-accepted by watching
tmux capture-pane for marker text and sending Enter . With
--desplega-local-auth , the custom API key confirmation prompt is also
auto-accepted. Login-method selection is deliberately not auto-accepted.
+--------------------+
| claude-bridge |
| - tmux paste |
| - transcript tail |
+----------+---------+
|
| tmux paste-buffer + Enter
v
+------+-----------------------------+
| tmux session claude-bridge-<id> |
| pane 0: claude --dangerously-... |
+------------------------------------+
Requirements
claude CLI on PATH, version >= 2.1.80
Claude Code authenticated for the spawned claude process.
# Drop-in print-mode usage.
bunx @desplega.ai/claude-bridge -p " say hi "
bunx @desplega.ai/claude-bridge -p " say hi " --output-format json
bunx @desplega.ai/claude-bridge -p " say hi " --output-format stream-json
# Opt in to Desplega/bridge envelopes for bridge-specific consumers.
bunx @desplega.ai/claude-bridge -p " say hi " --output-format stream-json --desplega-format
Install globally with Bun:
bun install -g @desplega.ai/claude-bridge
claude-bridge -p " say hi "
claude-bridge --help
Install globally with npm:
npm install -g @desplega.ai/claude-bridge
claude-bridge -p " say hi "
The installed command is claude-bridge . Bun is still required at runtime
because the published bin uses #!/usr/bin/env bun .
claude-bridge -p " say hi "
claude-bridge -p " say hi " --model sonnet
claude-bridge -p " say hi " --output-format json
claude-bridge -p " say hi " --output-format stream-json
printf ' say hi\n ' | claude-bridge --print
Print mode is intended for shell automation that would otherwise call
claude -p :
claude -p " say hi " --output-format json
claude-bridge -p " say hi " --output-format json
claude -p " say hi " --output-format stream-json
claude-bridge -p " say hi " --output-format stream-json
This is intended as a drop-in replacement for common claude -p automation.
In print mode the wrapper starts an interactive Claude session in tmux, waits
for the pane to become ready, sends the prompt through tmux, prints the
requested format, then kills the tmux session.
By default, print-mode stdout is reserved for the requested Claude-compatible
output. Bridge envelopes and bridge debug events are not written to stdout in
json or stream-json mode unless you explicitly pass --desplega-format .
claude-bridge does not call the Anthropic API itself. It launches the local
claude CLI and relies on whatever authentication that claude process can
use.
For local interactive machines, first make sure claude works:
claude auth status
claude -p " say hi "
Then run the bridge:
claude-bridge -p " say hi "
For headless CI, use the long-lived Claude Code OAuth token from:
claude setup-token
Set it exactly as printed:
export CLAUDE_CODE_OAUTH_TOKEN=sk-ant-oat01-...
claude-bridge -p " say hi " --output-format json
By default, the spawned Claude process receives HOME , CLAUDE_CONFIG_DIR ,
and CLAUDE_CODE_OAUTH_TOKEN ; Anthropic provider env vars such as
ANTHROPIC_API_KEY and ANTHROPIC_AUTH_TOKEN are cleared so the bridge does
not accidentally test a different auth path.
Use --desplega-local-auth when you intentionally want the spawned Claude
process to receive local auth-related env vars. If Claude shows the custom API
key confirmation prompt, this mode selects the API-key path:
ANTHROPIC_API_KEY=... claude-bridge --desplega-local-auth -p " say hi "
If Claude shows a browser login or login-method selector, the bridge will not
auto-select it. Run claude auth status , run claude setup-token , or attach
to the tmux pane shown in the banner and complete the prompt manually.
-p / --print requires a prompt argument or piped stdin. --output-format
requires print mode and accepts text , json , or stream-json ; the default is
text . --json-schema is also print-only.
Compatibility mode is the default. If you are replacing claude -p in scripts,
do not pass --desplega-format .
The final result comes from the transcript. When Claude writes a system
turn_duration row, the wrapper uses the latest assistant text it saw in that
turn.
text : prints only the final answer text plus a trailing newline. Wrapper
errors go to stderr and exit non-zero.
json : prints one final Claude-compatible JSON result object with the
answer in result , plus available transcript metadata such as session_id ,
duration_ms , and usage .
stream-json : streams raw Claude transcript JSONL rows as they are written.
The bridge does not wrap them in custom envelopes.
Use --desplega-format when you want the older bridge-owned JSON envelopes in
json or stream-json modes. This flag is for bridge-specific consumers, not
drop-in claude -p replacement scripts:
claude-bridge -p " say hi " --output-format stream-json --desplega-format
With --desplega-format , json includes bridge debug metadata when
--desplega-verbose is set, and stream-json prints newline-delimited bridge
events as the run progresses, then a final result event. This is a custom
claude-bridge event stream, not Claude's native stream-json schema.
Typical --desplega-format --output-format stream-json event types are:
{ "type" : " push " , "id" : " ab12cd34 " , "content" : " say hi " }
{ "type" : " transcript_folder " , "path" : " /Users/.../.claude/projects/... " }
{ "type" : " transcript_open " , "path" : " /Users/.../<uuid>.jsonl " , "session_id" : " ... " }
{ "type" : " transcript " , "row" :{ "type" : " assistant " , "message" :{ ... }}}
{ "type" : " result " , "subtype" : " success " , "is_error" : false , "result" : " Hi. " , "session_id" : " ... " }
These custom transcript events only exist with --desplega-format . In the
default compatibility mode, the same Claude data is written as the top-level
JSONL row.
--json-schema <schema|file> is bridge-owned. It is not forwarded to raw
claude -p ; the wrapper keeps the normal tmux/transcript path, injects schema
guidance with --append-system-prompt , extracts the last JSON value from the
final assistant text, and validates it locally with Zod.
Existing user-provided --append-system-prompt values are preserved. When a
schema is present, the wrapper merges those prompts with its schema instruction
instead of replacing them.
Schema print mode also installs a global Claude Code Stop hook in
~/.claude/settings.json . The hook is inert outside claude-bridge schema
runs; during a schema run it checks the final assistant text before Claude
stops and blocks the stop if it does not validate. That gives Claude a bounded
number of extra turns to answer with valid JSON before the wrapper exits.
Control that hook explicitly with:
claude-bridge --desplega-install
claude-bridge --desplega-uninstall
Install is append-only and idempotent: unrelated hooks are preserved, and stale
old claude-bridge hook commands are replaced with the current command.
The schema argument may be inline JSON or a path to a JSON file:
claude-bridge -p " Return the repo name " \
--json-schema ' {"type":"object","required":["name"],"properties":{"name":{"type":"string"}}} ' \
--output-format json
claude-bridge -p " Return the repo name " \
--json-schema ./schema.json \
--output-format text
Extraction is intentionally simple and deterministic:
Otherwise use the last fenced json block.
Otherwise use the final balanced JSON object or array in the reply.
Validation uses Zod's z.fromJSONSchema() converter. That API is still marked
experimental by Zod, but it keeps the bridge aligned with Zod's JSON Schema
support instead of maintaining a handwritten validator here. If Zod cannot
convert the schema, the wrapper treats that as a print-mode error.
With --output-format text , successful schema mode prints the extracted JSON
value as compact JSON. With --output-format json , the final result includes
structured_output alongside the original reply text in result . With
--desplega-format , bridge JSON results also include
structured_output_source .
If schema extraction or validation fails after Claude replies, json and
stream-json error results include raw_response with the unmodified Claude
reply. In text mode the same raw reply is printed to stderr under
Raw Claude reply: .
The compact stringified schem

[truncated]
