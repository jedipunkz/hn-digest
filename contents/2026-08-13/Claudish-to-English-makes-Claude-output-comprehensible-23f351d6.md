---
source: "https://github.com/gvzdv/claudish-to-english"
hn_url: "https://news.ycombinator.com/item?id=49287484"
title: "Claudish-to-English makes Claude output comprehensible"
article_title: "GitHub - gvzdv/claudish-to-english · GitHub"
author: "rickcarlino"
captured_at: "2026-08-13T15:50:06Z"
capture_tool: "hn-digest"
hn_id: 49287484
score: 1
comments: 0
posted_at: "2026-08-13T15:28:42Z"
tags:
  - hacker-news
  - translated
---

# Claudish-to-English makes Claude output comprehensible

- HN: [49287484](https://news.ycombinator.com/item?id=49287484)
- Source: [github.com](https://github.com/gvzdv/claudish-to-english)
- Score: 1
- Comments: 0
- Posted: 2026-08-13T15:28:42Z

## Translation

タイトル: クロード語から英語への変換により、クロードの出力がわかりやすくなります
記事のタイトル: GitHub - gvzdv/claudish-to-english · GitHub
説明: GitHub でアカウントを作成して、gvzdv/claudish-to-english の開発に貢献します。

記事本文:
GitHub - gvzdv/claudish-to-english · GitHub
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
gvzdv
/
クローディッシュから英語へ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
11 コミット 11 コミット .claude-plugin .claude-plugin フック フック .gitignore .gitignore ライセンス ライセンス README.md README.md rewrite-md.sh rewrite-md.sh rewrite.sh rewrite.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
プレイを表示する Claude Code プラグイン

各アシスタントの n-English リライト
メッセージ。 ollam を介してローカル LLM によって生成されます。表示専用です : Claude's
自分の推論と保存されたトランスクリプトには、オリジナルのテキストが保持されます。
画面上の変化を読み取ります。
オプションの 2 番目のフックは、Markdown ファイルを平易な英語に書き換えます。
書き込みまたは編集されます (オプトイン、デフォルトではオフ)。
ステータス: 動作中のプロトタイプ。すべてのフックは失敗して開きます - 何か問題が発生した場合
(ollam ダウン、タイムアウト、依存関係の欠落)、Claude のオリジナルが表示されるだけです。
テキスト。プラグインは回答を飲み込んだり破損したりすることはありません。
要件（最初にこれをお読みください）
このプラグインはローカル モデルにシェルアウトします。これらが整うまでは何も機能しません。
オラマのサーブ後にモデルを一度ウォームします (最初の呼び出しはゆっくりとしたコールドロードです)。
オラマ ラン gemma4:26b-mlx 「こんにちは」
ローカル モデルの準備ができていない場合、プラグインはテキストに対して何も行いません。
Claude の出力は通常どおり、変化せずに表示されます。これは仕様によるものであり、バグではありません。スキップします
(フェールオープン) ollam がダウンしている場合、リクエストがタイムアウトしている場合、またはモデルが正常に動作していない場合
引っ張られた。セッション中に初めてこの問題が発生すると、その理由が表示されます。
フックは画面上に 1 行の通知を追加し、Markdown フックは
システムメッセージ 。したがって、サイレントスキップは決して謎ではありません（セッションごとに1回、設定されます）
CLAUDISH_NOTICE=0 で無音になります)。
実際にお持ちのモデルをお選びください。デフォルトは gemma4:26b-mlx です。
Apple-silicon (MLX) ビルド — Mac では正しい選択ですが、macOS のみです。オン
Windows では実行されないため、通常のタグに切り替える必要があります (「
Windows セットアップ)。 (上記のように) 引くか、より小さい/より速い速度で引っ張ります。
モデルを選択し、CLAUDISH_MODEL をそのモデルのモデルに設定してプラグインをポイントします。
環境内の正確な ollam タグ (を参照)
プラグインの設定)。 CLAUDISH_MODEL が
プルしていないモデルの場合、すべての書き換えはスキップされます - 1 回限りの通知
上。
フックはバーです

sh スクリプト。 Windows では、Claude Code は Git 経由で実行します。
Bash (Windows 用 Git)。
インストール後にターミナルを再起動して、 jq 、 ollama 、および Git Bash を有効にします。
PATH ( jq --version と ollama --version を確認してください)。
デフォルトのモデルは macOS のみです。Windows ユーザーはこれをオーバーライドする必要があります。の
プラグインのデフォルトの gemma4:26b-mlx は、Apple-silicon (MLX) ビルドです。
Windows 上で実行されるため、未設定のままにしておくと、すべての書き換えが警告なしにスキップされることになります。
Windows では、常に CLAUDISH_MODEL を通常の (MLX 以外の) タグに設定します。テーブル
上記では例として gemma4:26b を使用しています。あなたのマシンに適合する場合は別のものを選択してください
もっと良い。
Ollama の起動後にモデルを一度ウォームします (最初の呼び出しは低速なコールド ロードです)。
オラマ ラン ジェマ4:26b 「こんにちは」
次に、settings.json の env ブロックで CLAUDISH_MODEL を設定します (「
プラグインの設定 — その方法は次の場合と同じです。
Windows)、または PowerShell からの 1 回限りのセッションの場合:
$ env:CLAUDISH_MODEL = " gemma4:26b " ;クロード
Windows でのセッション中のキル スイッチに相当するもの
(セッション中の切り替え):
New-Item - ItemType ファイル $HOME \.claude\claudish - off # 一時停止リライト
Remove-Item $HOME \.claude\claudish - オフ # 再開
(Git Bash では、そのセクションの touch / rm コマンドはそのまま機能します。)
CLAUDISH_MD_DIR をスラッシュで記述します
( C:/dev/docs/plain ) bash 側のパスチェックが一致するようにする
CLAUDISH_DEBUG=1 ログは Git Bash の一時ディレクトリに配置されます。
( $TMPDIR/claudish-to-english/ 、通常は
C:\Users\<you>\AppData\Local\Temp\claudish-to-english\ )。
このリポジトリから直接 (独自のマーケットプレイスも提供します):
/プラグイン マーケットプレイス add gvzdv/claudish-to-english
/plugin install claudish-to-english@gvzdv-plugins
Anthropic チームによるレビュー後、プラグインはコミュニティ マーケットプレイスからインストールできるようになります。
/プラグイン マーケットプレイス追加 anthropics/claude-plugins-community
/プラグイン

claudish-to-english@claude-community をインストールします
インストールの概要に「/reload-plugins を実行してアクティブ化する」と表示されている場合は、 、そのコマンドを実行します。
インストールする前に試してください (1 セッションでロードされ、インストールはされません)。
クロード --plugin-dir /path/to/claudish-to-english
編集後に /reload-plugins を実行します。ロードされない場合は、/plugin を確認してください
「エラー」タブ。
すべての動作は CLAUDISH_* 環境変数によって制御されます (完全なリストは次のとおりです)。
構成は以下の通り）。からインストールする場合
マーケットプレイス、settings.json の Claude Code の env ブロックに設定します。
プラグイン自体のフック/hooks.json は編集しないでください。これは読み取り専用で存在します。
プラグイン キャッシュ ( ~/.claude/plugins/cache/… ) は更新のたびに上書きされます。
個人的なすべてのプロジェクトのセットアップの場合は、 ~/.claude/settings.json を使用します。
{
"環境" : {
"CLAUDISH_MODEL" : " gemma4:26b-mlx " ,
"CLAUDISH_MODE" : " 追加 "
}
}
フックはクロード コードが生成するサブプロセスであるため、これらを継承します。いくつか
知っておくべきこと:
env を編集した後、Claude Code を再起動します。値は起動時に取得されます。
したがって、実行中のセッションでは古いセッションが保持されます。
env はスコープ間でマージされません。最も優先順位の高い設定ファイル
env を定義するものはブロック全体を提供します。これは下位のブロックとは結合されません。
スコープ。優先順位: 管理対象 → ローカル → プロジェクト → ユーザー。すべてを保管してください
勝ったファイルの CLAUDISH_* 変数。
スコープ: ~/.claude/settings.json (すべてのプロジェクト) ·
.claude/settings.json (リポジトリと共有、チェックイン) ·
.claude/settings.local.json (あなただけ、このリポジトリだけ)。
ファイルを編集せずに簡単に 1 回限り — フックは起動シェルを継承します。
CLAUDISH_MODEL=llama3.2:3b クロード
フックが起動していることを確認するには、CLAUDISH_DEBUG=1 を設定して監視します。
"$TMPDIR"/claudish-to-english/debug.log 。
クロード コードは、ストリーミング チャンクごとに MessageDisplay イベントを 1 回発生させます。
メッセージごとに 1 回。それぞれの火は別個のプロセスで実行されます

message_id 、
Index 、最終フラグ、およびこのチャンクのデルタ (テキストのフラグメント。
メッセージ全体）。したがって、フックはすべてのデルタを一時ファイルにバッファリングします (キーは次のとおりです)。
message_id )、最後のチャンクでのみモデルを呼び出します。
メッセージは既知です:
チャンク 0 (最終:false) ─┐
チャンク 1 (final:false) ─┤ 各デルタを $TMPDIR/claudish-to-english/<session>/<message>/<index>.part に追加します
チャンク 2 (final:false) ─┘ → 何も発行しない (追加) または "" (置換)
チャンク 3 (final:true) ──► メッセージ全体を再構築 → ollam を 1 回呼び出し → 書き換えを表示
→ バッファを削除する
最後のチャンクでは、元のユーザーの質問も読み取ります。
書き換えを維持するために、トランスクリプトを作成し、それをコンテキストとしてのみモデルに渡します。
話題に上ります。モデルには、決して質問に答えたり繰り返したりしないように指示されます。それだけ
アシスタントのメッセージを書き換えます。
クローディッシュモード
画面上
注意事項
追加 (デフォルト)
元のストリームは通常通り、その後 💬 簡単に言うと、ブロックが追加されます。
最も安全です。ストリーミング損失はありません。 LLM が失敗した場合は、追加のブロックを取得できないだけです。
交換する
簡易バージョンのみ (ストリーミング中に元のチャンクが抑制されます)。
実験的。 LLM レイテンシの後にすべてが一度に表示されます。失敗すると、完全なオリジナルが再表示されます。
Markdown ファイルの書き換え (オプションの 2 番目のフック)
PostToolUse フック ( rewrite-md.sh ) は Markdown ファイルをプレーンに書き換えます
執筆または編集されるときは英語です。表示フックとは異なり、これは変わります
ディスク上のバイト数。
ディレクトリごとにオプトインします。 CLAUDISH_MD_DIR が設定されていない限り何もしません。
解決されたパスがそのディレクトリ内にある *.md ファイルのみに触れます。毎
編集した他の README 、 CLAUDE.md 、またはドキュメントはそのまま残されます。
どちらのモードでも: YAML フロントマターは分割され、逐語的に再アタッチされ、フェンスされます。
コードはモデル命令に任せられ、短いファイルはスキップされ、書き込みは

原子的な。ここでのフェールオープンとは、ファイルがエージェントが書き込んだとおりにそのまま残されることを意味します。
大きなファイルは遅いです。 gemma4:26b-mlx (デフォルト) は約 60 で書き換えます
トークン/秒なので、長いプランや仕様には 30 ～ 120 秒かかる場合があります。このフックにより、最大で次のことが可能になります
CLAUDISH_MD_TIMEOUT (150 秒) 180 秒以内 PostToolUse フック バジェット。書き直しの場合
それでもタイムアウトになる場合は、上記の 1 回限りの通知が表示されます — 制限を上げるか、設定してください
CLAUDISH_MODEL をより小さいモデルに変更します。
兄弟モード (安全なデフォルト) で、同じ方法で 1 つのディレクトリに対して有効にします。
他のすべての設定と同様に、 settings.json の env ブロック:
{
"環境" : {
"CLAUDISH_MD_DIR" : " /ABS/PATH/docs/plain " ,
"CLAUDISH_MD_MODE" : "兄弟"
}
}
上書きモードでは、マーカー コメントは YAML の後に書き込まれます。
したがって、フロントマターはパーサーが期待する行 1 に留まります。
ヴァール
デフォルト
意味
CLAUDISH_ENABLED
1
マスタースイッチ。 0 = すべてを通過させます。セッション開始時に一度読み取ります。
CLAUDISH_OFF_FILE
~/.claude/claudish-off
ランタイムキルスイッチ。このファイルが存在する間、書き換えは一時停止します。つまり、すべてのメッセージが再チェックされるため、env vars とは異なり、セッション中に動作します。 「セッション中の切り替え」を参照してください。
クローディッシュモード
追加する
追加または置換 (フックの表示)。
クローディッシュ_モデル
gemma4:26b-mlx
ollam モデル名 (MLX = Apple-silicon のみ。Windows ユーザーは上書きする必要があります)。
クローディッシュ・オラマ
http://ローカルホスト:11434
オラマのベース URL。
CLAUDISH_MIN_CHARS
200
散文 (コードを取り除いた) がこれより短いメッセージ/ファイルをスキップします。
CLAUDISH_STUB
0
1 = モデルの代わりに決定的なスタブ (表示機構のテスト用)。
CLAUDISH_TIMEOUT
45
表示フックの LLM クライアントのタイムアウト (秒)。フックのタイムアウト (60 秒) 未満に保ちます。
CLAUDISH_MD_TIMEOUT
150
Markdown ファイルフックの LLM クライアントのタイムアウト (秒)。意図的に高くする — 大規模なモデルで長いドキュメントを書き直すと時間がかかります。 PostToolUse の下に置いてください

フックのタイムアウト (180 秒)。
CLAUDISH_DEBUG
0
1 = デバッグ ログを $TMPDIR/claudish-to-english/ に書き込みます。
クローディッシュ_通知
1
1 = ollam に到達できない、呼び出しがタイムアウトした、またはモデルがプルされないために書き換えがスキップされたときに、セッションごとに 1 回だけ通知を表示します (display フックは画面上に追加します。Markdown フックは systemMessage を使用します)。 0 = 完全にサイレントを維持します (純粋なフェールオープン)。
CLAUDISH_MD_DIR
(未設定)
マークダウンフックのオプトイン。このディレクトリ下の *.md のみが書き換えられます。未設定 = Markdown フックは何も行いません。
クローディッシュ_MD_モード
兄弟
兄弟 ( NAME.plain.md ) または上書き (その場で)。
CLAUDISH_MD_SUFFIX
平野
兄弟中置: NAME.<suffix>.md 。
hooks/hooks.json では、表示フック ( MessageDisplay ) には 60 秒のタイムアウトがあり、
Markdown フック ( PostToolUse ) のタイムアウトは 180 秒です。ファイル フックの方がタイムアウトが高くなります。
大規模なモデルで長いドキュメントを書き直すには数分かかる場合があるためです。
CLAUDISH_TIMEOUT と CLAUDISH_MD_TIMEOUT は、LLM 呼び出し自体を制限します。
天井の下にあるので、代わりにきれいに開かなくなります
[切り捨てられた]
クイックキルスイッチ: CLAUDISH_ENABLED=0 を設定するか、プラグインを無効にします (両方とも適用されます)
次のセッション開始時からのみ)、または ~/.claude/claudish-off をタッチしてセッションを一時停止します
すでに実行中のセッション — セッション中の切り替えを参照
以下。
CLAUDISH_ENABLED と他の環境変数が読み取られます

[切り捨てられた]

## Original Extract

Contribute to gvzdv/claudish-to-english development by creating an account on GitHub.

GitHub - gvzdv/claudish-to-english · GitHub
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
gvzdv
/
claudish-to-english
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
11 Commits 11 Commits .claude-plugin .claude-plugin hooks hooks .gitignore .gitignore LICENSE LICENSE README.md README.md rewrite-md.sh rewrite-md.sh rewrite.sh rewrite.sh View all files Repository files navigation
A Claude Code plugin that shows a plain-English rewrite of each assistant
message, produced by a local LLM via ollama . It is display-only : Claude's
own reasoning and the saved transcript keep the original text — only what you
read on screen changes.
An optional second hook rewrites Markdown files into plain English when they
are written or edited (opt-in, off by default).
Status: working prototype. Every hook fails open — if anything goes wrong
(ollama down, timeout, missing dependency), you simply see Claude's original
text. The plugin can never swallow or corrupt an answer.
Requirements (read this first)
This plugin shells out to a local model. Nothing works until these are in place:
Warm the model once after ollama serve (the first call is a slow cold load):
ollama run gemma4:26b-mlx " hi "
If the local model isn't ready, the plugin does nothing to your text —
Claude's output shows normally, unchanged. That is by design, not a bug. It skips
(fails open) when ollama is down, the request times out, or the model isn't
pulled. The first time that happens in a session it tells you why: the display
hook appends a one-line notice on screen, and the Markdown hook shows a
systemMessage . So a silent skip is never a mystery (once per session; set
CLAUDISH_NOTICE=0 to silence it).
Pick a model you actually have. The default is gemma4:26b-mlx , an
Apple-silicon (MLX) build — the right choice on a Mac, but macOS-only . On
Windows it doesn't run, so you must switch to a regular tag (see
Windows setup ). Pull it (as above), or pull a smaller/faster
model and point the plugin at it by setting CLAUDISH_MODEL to that model's
exact ollama tag in your env (see
Configuring the plugin ). If CLAUDISH_MODEL names a
model you have not pulled, every rewrite is skipped — with the one-time notice
above.
The hooks are bash scripts; on Windows, Claude Code runs them through Git
Bash (Git for Windows).
Restart your terminal after installing so jq , ollama , and Git Bash are on
PATH (check jq --version and ollama --version ).
The default model is macOS-only — Windows users must override it. The
plugin's default, gemma4:26b-mlx , is an Apple-silicon (MLX) build that doesn't
run on Windows, so leaving it unset means every rewrite is silently skipped.
Always set CLAUDISH_MODEL to a regular (non-MLX) tag on Windows. The table
above uses gemma4:26b as an example; choose another if it fits your machine
better.
Warm the model once after launching Ollama (the first call is a slow cold load):
ollama run gemma4:26b " hi "
Then set CLAUDISH_MODEL in the env block of your settings.json (see
Configuring the plugin — that method is identical on
Windows), or for a one-off session from PowerShell:
$ env: CLAUDISH_MODEL = " gemma4:26b " ; claude
Windows equivalents of the mid-session kill switch
( Toggling mid-session ):
New-Item - ItemType File $HOME \.claude\claudish - off # pause rewrites
Remove-Item $HOME \.claude\claudish - off # resume
(In Git Bash the touch / rm commands from that section work as-is.)
Write CLAUDISH_MD_DIR with forward slashes
( C:/dev/docs/plain ) so the bash-side path checks match
The CLAUDISH_DEBUG=1 log lands under Git Bash's temp directory
( $TMPDIR/claudish-to-english/ , typically
C:\Users\<you>\AppData\Local\Temp\claudish-to-english\ ).
Directly from this repository (also serves its own marketplace):
/plugin marketplace add gvzdv/claudish-to-english
/plugin install claudish-to-english@gvzdv-plugins
After review by the Anthropic team, the plugin will be available to install from the community marketplace:
/plugin marketplace add anthropics/claude-plugins-community
/plugin install claudish-to-english@claude-community
If the install summary says Run /reload-plugins to activate. , run that command.
Try before installing (loads it for one session, no install):
claude --plugin-dir /path/to/claudish-to-english
Run /reload-plugins after edits; if it doesn't load, check the /plugin
Errors tab.
All behavior is controlled by CLAUDISH_* environment variables (full list in
Configuration below). When you install from a
marketplace, set them in Claude Code's env block in settings.json — do
not edit the plugin's own hooks/hooks.json , which lives in the read-only
plugin cache ( ~/.claude/plugins/cache/… ) and is overwritten on every update.
For a personal, all-projects setup, use ~/.claude/settings.json :
{
"env" : {
"CLAUDISH_MODEL" : " gemma4:26b-mlx " ,
"CLAUDISH_MODE" : " append "
}
}
The hooks are subprocesses Claude Code spawns, so they inherit these. A few
things to know:
Restart Claude Code after editing env . The value is captured at launch,
so a running session keeps the old one.
env does not merge across scopes. The highest-precedence settings file
that defines env supplies the entire block — it isn't combined with lower
scopes. Precedence: managed → local → project → user. Keep all your
CLAUDISH_* vars in whichever file wins.
Scopes: ~/.claude/settings.json (all your projects) ·
.claude/settings.json (shared with a repo, checked in) ·
.claude/settings.local.json (just you, just this repo).
Quick one-off without editing a file — hooks inherit the launching shell:
CLAUDISH_MODEL=llama3.2:3b claude
To confirm the hook is firing, set CLAUDISH_DEBUG=1 and watch
"$TMPDIR"/claudish-to-english/debug.log .
Claude Code fires the MessageDisplay event once per streamed chunk , not
once per message. Each fire is a separate process carrying message_id ,
index , a final flag, and this chunk's delta (a text fragment, not the
whole message). So the hook buffers every delta to a temp file (keyed by
message_id ) and only calls the model on the final chunk, once the whole
message is known:
chunk 0 (final:false) ─┐
chunk 1 (final:false) ─┤ append each delta to $TMPDIR/claudish-to-english/<session>/<message>/<index>.part
chunk 2 (final:false) ─┘ → emit nothing (append) or "" (replace)
chunk 3 (final:true) ──► reconstruct full message → call ollama once → show the rewrite
→ delete the buffer
On that final chunk it also reads the original user question from the
transcript and passes it to the model as context only — to keep the rewrite
on-topic. The model is told never to answer or repeat the question; it only
rewrites the assistant's message.
CLAUDISH_MODE
On screen
Notes
append (default)
Original streams normally, then a 💬 In plain English: block is appended.
Safest. No streaming loss; if the LLM fails you just don't get the extra block.
replace
Only the simplified version (original chunks suppressed while streaming).
Experimental. Appears all at once after LLM latency; on failure it re-shows the full original.
Markdown file rewrite (optional second hook)
A PostToolUse hook ( rewrite-md.sh ) rewrites Markdown files into plain
English when they are written or edited. Unlike the display hook, this changes
bytes on disk.
Opt-in by directory. It does nothing unless CLAUDISH_MD_DIR is set, and it
only touches *.md files whose resolved path is inside that directory. Every
other README , CLAUDE.md , or doc you edit is left alone.
In both modes: YAML frontmatter is split off and re-attached verbatim , fenced
code is left to the model instruction, short files are skipped, and the write is
atomic. Fail-open here means the file is left exactly as the agent wrote it .
Large files are slow. gemma4:26b-mlx (the default) rewrites at roughly 60
tokens/s, so a long plan or spec can take 30–120s. This hook allows up to
CLAUDISH_MD_TIMEOUT (150s) inside a 180s PostToolUse hook budget; if a rewrite
still times out you get the one-time notice above — raise those limits, or set
CLAUDISH_MODEL to a smaller model.
Enable it for one directory, in sibling mode (the safe default), the same way
as every other setting — the env block of your settings.json :
{
"env" : {
"CLAUDISH_MD_DIR" : " /ABS/PATH/docs/plain " ,
"CLAUDISH_MD_MODE" : " sibling "
}
}
In overwrite mode the marker comment is written after any YAML
frontmatter, so the frontmatter stays on line 1 where parsers expect it.
Var
Default
Meaning
CLAUDISH_ENABLED
1
Master switch. 0 = pass everything through. Read once at session start.
CLAUDISH_OFF_FILE
~/.claude/claudish-off
Runtime kill switch. While this file exists, rewrites pause — re-checked every message, so unlike env vars it works mid-session. See Toggling mid-session .
CLAUDISH_MODE
append
append or replace (display hook).
CLAUDISH_MODEL
gemma4:26b-mlx
ollama model name (MLX = Apple-silicon only; Windows users must override).
CLAUDISH_OLLAMA
http://localhost:11434
ollama base URL.
CLAUDISH_MIN_CHARS
200
Skip messages/files whose prose (code stripped) is shorter than this.
CLAUDISH_STUB
0
1 = deterministic stub instead of the model (for testing display mechanics).
CLAUDISH_TIMEOUT
45
LLM client timeout for the display hook (seconds). Keep it below that hook's timeout (60s).
CLAUDISH_MD_TIMEOUT
150
LLM client timeout for the Markdown file hook (seconds). Higher on purpose — a large model rewriting a long doc is slow. Keep it below the PostToolUse hook timeout (180s).
CLAUDISH_DEBUG
0
1 = write a debug log to $TMPDIR/claudish-to-english/ .
CLAUDISH_NOTICE
1
1 = show a one-time, once-per-session notice when a rewrite is skipped because ollama is unreachable, the call timed out, or the model isn't pulled (display hook appends it on screen; Markdown hook uses a systemMessage ). 0 = stay fully silent (pure fail-open).
CLAUDISH_MD_DIR
(unset)
Markdown hook opt-in. Only *.md under this directory is rewritten. Unset = the Markdown hook does nothing.
CLAUDISH_MD_MODE
sibling
sibling ( NAME.plain.md ) or overwrite (in place).
CLAUDISH_MD_SUFFIX
plain
Sibling infix: NAME.<suffix>.md .
In hooks/hooks.json the display hook ( MessageDisplay ) has a 60s timeout and
the Markdown hook ( PostToolUse ) has a 180s timeout — the file hook is higher
because a large model rewriting a long document can take a couple of minutes.
CLAUDISH_TIMEOUT and CLAUDISH_MD_TIMEOUT keep the LLM call itself bounded
below those ceilings, so it fails open cleanly instead
[truncated]
Quick kill switch: set CLAUDISH_ENABLED=0 or disable the plugin (both apply
only from the next session start), or touch ~/.claude/claudish-off to pause a
session that's already running — see Toggling mid-session
below.
CLAUDISH_ENABLED and the other env vars are read on

[truncated]
