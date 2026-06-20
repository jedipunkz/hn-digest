---
source: "https://github.com/nikitadoudikov/claude-pulse"
hn_url: "https://news.ycombinator.com/item?id=48612844"
title: "Pulse – a local dashboard for Claude Code, approve tool calls from your phone"
article_title: "GitHub - nikitadoudikov/claude-pulse: Local, zero-dependency dashboard for Claude Code: live token usage and context, lost-session recovery, full-text search, and approve tool calls from your phone. · GitHub"
author: "nikitadvd"
captured_at: "2026-06-20T21:32:55Z"
capture_tool: "hn-digest"
hn_id: 48612844
score: 2
comments: 0
posted_at: "2026-06-20T20:46:50Z"
tags:
  - hacker-news
  - translated
---

# Pulse – a local dashboard for Claude Code, approve tool calls from your phone

- HN: [48612844](https://news.ycombinator.com/item?id=48612844)
- Source: [github.com](https://github.com/nikitadoudikov/claude-pulse)
- Score: 2
- Comments: 0
- Posted: 2026-06-20T20:46:50Z

## Translation

タイトル: Pulse – Claude Code のローカル ダッシュボード、携帯電話からのツール呼び出しを承認
記事のタイトル: GitHub - nikitadoudikov/claude-pulse: Claude Code のローカル、依存性ゼロのダッシュボード: ライブ トークンの使用法とコンテキスト、失われたセッションの回復、全文検索、電話からのツール呼び出しの承認。 · GitHub
説明: クロード コード用の依存関係のないローカル ダッシュボード: ライブ トークンの使用法とコンテキスト、失われたセッションの回復、全文検索、電話からのツール呼び出しの承認。 - ニキタドゥディコフ/クロード・パルス

記事本文:
GitHub - nikitadoudikov/claude-pulse: クロード コードのローカル、依存性ゼロのダッシュボード: ライブ トークンの使用状況とコンテキスト、失われたセッションの回復、全文検索、電話からのツール呼び出しの承認。 · GitHub
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
ニキタドゥディコフ
/
クロード・パルス
プ

ブリック
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
10 コミット 10 コミット bin bin docs docs フック フック public public src src .gitignore .gitignore ライセンス ライセンス README.md README.md config.example.json config.example.json package.json package.json すべてのファイルを表示 リポジトリ ファイル ナビゲーション
Claude Code のローカル ダッシュボード。Claude が何をしているか、何を費やしているかを表示し、携帯電話からツール呼び出しを承認できるようにします。依存関係がゼロなので、マシンから何も離れません。
Claude Code はすでにすべてのセッションをディスクに書き込んでいます。 Pulse はこれらのファイル (読み取り専用) を読み取り、ライブ ダッシュボードに変換します。時間、日、週ごとのトークンの使用状況、アクティブなセッションのコンテキストフィル、仕事中のクロードのアンビエント ビュー、これまでに実行したすべての内容の全文検索、およびクロードが必要なときにデスクトップまたは携帯電話に [許可] / [すべて許可] / [拒否] ボタンを備えた通知が表示されます。アカウント、テレメトリ、ネットワーク通話はありません。
携帯電話から承認してください。 「許可」/「すべて許可」/「拒否」ボタンが機能する状態で押す。 Wi-Fi セットアップ、IP、オープンポートは必要ありません。どこからでも、携帯電話でも動作します。
セッションを失うことはありません。 1 つのコマンドで最後のセッションを読み取り可能なトランスクリプトとして復元し、Pulse がアクティブなセッションを自動スナップショットするため、ラップトップのクラッシュやフリーズによってコンテキストが失われることはありません。
支出を参照してください。設定した予算に対して、時間、日、週、モデル、プロジェクトごとのライブ トークンと API 相当コストが表示され、予算を超えた場合には電話でアラートが送信されます。
雰囲気のあるオフィス。小さなマスコットが働いたり、休憩したり、待機している様子を、大まかな到着予定時刻とともに全画面表示します。セカンドモニターに静かにハマります。
すべてを検索してください。ディスク上のすべてのセッションを全文検索し、ワンクリックでトランスクリプトを表示します。
ローカル

そしてプライベート。 ~/.claude は読み取り専用で読み取り、 127.0.0.1 で機能し、依存関係はなく、テレメトリもありません。
ノード 18 以降が必要です。インストールせずに実行します。
npx パルス・フォー・クロード・コード
またはクローンを作成します。
git clone https://github.com/nikitadoudikov/claude-pulse.git
cd クロードパルス
ノード bin/cli.js
どちらの方法でも http://127.0.0.1:4317 が開きます。デスクトップと電話を取得するには
通知を受け取り、ツール呼び出しを承認するには、フックを接続します (1 つのコマンドで安全に実行できます)
再実行):
claude-pulse install-hooks # フックを ~/.claude/settings.json に追加します
その後、Claude Code を再起動すれば準備完了です。その他のオプション:
claude-pulse --port 4317 # ポートを変更します
claude-pulse --no-open # ブラウザを開きません
実行し続けてください
フォアグラウンドで実行すると、そのターミナルを閉じると Pulse が停止します。それを保つために
独立して動作している場合は、バックグラウンドで実行します。
claude-pulse start # 切り離されて実行され、ターミナルを閉じても存続します
クロードパルスステータス # 実行中ですか?
クロード・パルス ストップ # やめて
claude-pulse restart # 停止して再び開始する
ターミナルがクラッシュした場合、claude-pulse start を実行すると 1 つのコマンドで元の状態に戻ります。
バックグラウンド インスタンスはそもそもクラッシュの影響を受けません。
macOS では、Pulse をシステムに渡して、ログイン時に開始して再生成することができます。
それが死んだ場合にはそれ自体:
claude-pulse install-service # ログイン時に開始、自動再起動
claude-pulse uninstall-service # 削除します
失われたセッションを回復する
端末がクラッシュしたり、ラップトップがフリーズしたり、セッション制限に達したりしていませんか?何も失われない：クロード
コードは、すべてのセッションを発生時にディスクに書き込みます。 1 つのコマンドで最後のコマンドが実行されます
戻って要約を印刷し、判読可能なトランスクリプトを保存します。
クロードパルスは # 最新のセッションを回復します
クロードパルスリカバリー2 #その前のもの
claude-pulse reverse < id > # 特定のセッション
ライト マークダウン ファイルを ~/.claude-pulse/exports/ に保存します (15 MB のログ)
約 180 KB のファイルになります) と印刷します

の完全なトランスクリプトを読むためのリンクです。
ブラウザまたは携帯電話で。ダッシュボードで任意のセッションを開いて使用することもできます。
トランスクリプトを開き、.md をダウンロードします。
Pulse の実行中、最近アクティブなすべてのセッションの自動スナップショットも作成されます。
~/.claude-pulse/exports/snapshots/ (セッションごとに 1 つのファイル、次の場合にのみ書き換えられます)
変わります）。したがって、一度も実行しなくても、最新の状態が常にディスク上にあります。
回復します。オフにするには、~/.claude-pulse.json で snapshotMinutes を 0 に設定します。
すべてを一度にバックアップするには、claude-pulse export-all がすべてのセッションを書き込みます
単一の小さな gzip マークダウン ファイルにまとめるか、すべての履歴をダウンロードします。
セッション画面。
どこで何かをしたのか分からなくなりましたか？ 「セッション」画面には検索ボックスがあります。
ディスク上のすべてのセッションをスキャンして単語または語句を探し、そのセッションに直接ジャンプします。
転写。携帯電話からも機能します。
最も単純な電話制御は ntfy 通知自体です。これは動作を伝えます。
[許可] / [すべて許可] / [拒否] ボタン (上記を参照)、ネットワーク設定はまったくありません。
より詳細なビューを表示するには、同じ Wi-Fi で http://<your-machine>:4317/phone を開きます。
(bindLan: true が必要) クロードが現在何をしているのかを確認し、一時停止を加えます /
再開ボタン。一時停止すると、クロードは次のツールの実行を停止します。
再開します。どちらも PreToolUse フックを配線する必要があります。
┌───────────┐ .jsonl を書き込みます ┌───────────┐ SSE ┌──────────┐
│ クロード・コード │ ───────▶ │ パルス（読み取り専用） │ ───────▶ │ ダッシュボード │
│ (端末) │ │ 127.0.0.1:4317 │ │ + 電話 │
━─────┬───────┘ └─────────────┘ └────────

┘
│
│ フック: 通知 · 停止 · PreToolUse
▼
┌───────────────────┐
│ ~/.claude-pulse/ │ 保留中の承認・決定・イベント
━━━━━━━━━━━━┘
Claude Code は、すべてのセッションを ~/.claude/projects/ の下に JSONL として記録します。各アシスタント
メッセージには、タイムスタンプ付きの使用ブロック (入力、出力、およびキャッシュ トークン) が含まれます。
Pulse はこれらのファイルを読み取り (読み取り専用)、変更時間ごとに各ファイルをキャッシュします。
変更されていないセッションは再解析されることはなく、数値が集計されます。ブラウザ
Server-Sent Events を通じてライブ更新を取得します。 3 つの小さなフックがクロード コードに伝えます
あなたが必要なとき、ターンが終了したとき、ツールを実行したいときにパルスを送ります。
クロードがあなたを必要とするときの通知
Claude Code は、注意が必要なときにフックを実行できます。通知をポイントする
バンドルされたスクリプトでイベントが発生すると、Pulse はバナーを表示し、デスクトップを起動します
タブがバックグラウンドにある場合でも、通知が表示されます。
claude-pulse install-hooks # フックを ~/.claude/settings.json に接続します (再実行しても安全です)
claude-pulse uninstall-hooks # それらを削除します
設定を一度バックアップし、既存のフックの隣にマージします。
重複を追加することはありません。その後、Claude Code を再起動してください。代わりに手動で行うには、
これを ~/.claude/settings.json に追加します (クローンへの絶対パスを使用します)。
{
「フック」: {
「通知」: [
{
"マッチャー" : " " ,
「フック」: [
{ "タイプ" : " コマンド " 、 "コマンド" : " ノード /absolute/path/to/claude-pulse/hooks/notify-hook.js " }
]
}
]、
「停止」: [
{
"マッチャー" : " " ,
「フック」: [
{ "type" : " コマンド " 、 "コマンド" : " ノード /absolute/path/to/claude-pulse/hooks/stop-hook.js " }
]
}
]、
"PreToolUse" : [
{
"マッチャー" : " " ,
「フック」: [
{ "タイプ" : " コマンド " ,

"command" : " ノード /absolute/path/to/claude-pulse/hooks/permission-hook.js " }
]
}
]
}
}
claude-pulse を実行し続ければ準備完了です。
ダッシュボード (および携帯電話) からツールを承認します。
PreToolUse フックを接続した状態で、Claude が必要なものを実行したいとき
許可すると、Pulse に [Allow]、[Allow all]、および [Allow all] を含む承認カードが表示されます。
否定します。残りの実行を要求するすべての停止を許可します。
これはクロードを決して吊るさないように作られています。読み取り専用ツールはそのまま通過します。
Pulse が実行されていません。承認タイムアウト (60 秒) 以内に連絡がありません。
デフォルト、approvalTimeoutMs を設定）、またはエラーが発生した場合は、通常の状態に戻ります。
ターミナルプロンプト。無視しても何も壊れません。電話プッシュでは、Allow が送信されます。
「すべて許可」ボタンと「拒否」ボタン。
電話から承認するには、ntfyTopic (下記) と ntfy
アプリ。プッシュ通知には [許可] 、 [すべて許可] 、 [拒否] ボタンがあり、
タップすると、ntfy を介してプライベート返信トピックに回答が返信されます。
パルスは聞き続けます。同じ Wi-Fi、IP、開いているポートはありません。どこからでも機能します。
携帯電話でも。 Pulse は、実際に応答を待っている間のみ応答に作用します。
そのため、古い通知では何もできません。
クロードはツールを実行したいと考えています
│
▼
PreToolUseフック──▶パルス──プッシュ──▶電話通知
│ │
│「許可」をタップ
│ │
│ ntfy 経由で応答が返されます
│ │
▼ ▼
フックはまだ待機中 ◀── 決定 ◀── パルス (返信トピックに登録済み)
│
▼
フックは「allow」を返します ──▶ クロードがツールを実行します
電話プッシュ (オプション)
クロードがあなたを必要とするとき、または仕事が終わったときに携帯電話にプッシュしてもらうには、押しにくいものを選択してください。
トピック名を推測し、無料の ntfy アプリをインストールしてサブスクライブします。
そのトピックを ~/.claude-pulse.json に設定します。
{ "ntfyTopic" : " クロードパルス-9f3a7c " }
上記のフックを配線すると、

クロードが待っているときに化フックがプッシュされます
ターンが終了すると停止フックがプッシュされます (30 秒までデバウンスされるため、
行ったり来たりしてもスパムメールにはなりません）。トピックを知っている人なら誰でも読むことができますので、
ランダムな名前。
予算を設定すると (下記)、ローリング ウィンドウが 80% を超えたときにも Pulse がプッシュします。
その場合は予算の 100% なので、調べるのではなく自分のポケットからわかります。
config.example.json を ~/.claude-pulse.json にコピーして編集します。すべてのフィールドはオプションです。
{
"プラン" : "最大20" 、
"contextLimit" : 200000 、
"アイドル分" : 10 、
"approvalTimeoutMs" : 60000 、
"予算" : { "5 時間" : 140 、 "日" : 360 、 "週" : 1100 }
}
制限について
Anthropic は正確なサブスクリプション制限を公開しておらず、使用量に基づいています。
固定のトークン数ではなく。 Pulse は実際のプランの上限を読み取ることができないため、
上記の予算は、API に相当する大まかな見積もりであり、予算に合わせて調整します。
観察する。 pro、max5、max20 のプリセットは出発点であり、公式のものではありません
数字。トークンのコストは、純粋に使用量としてのパブリック API の定価から見積もられます。
代理人;サブスクリプション ユーザーはトークンごとに料金を支払いません。
Pulse はローカルファーストであり、オプトインです。そのままでは 127.0.0.1 のみにバインドされます。
発信呼び出しは行わず、依存関係はなく (サプライ チェーンなし)、読み取ります。
~/.claude 読み取り専用。マシンからは何も出ず、分析も行われません。 2
オプション機能を使用するとこれが変更され、オンにするまで両方ともオフになります。
電話プッシュ ( ntfyTopic ) はパブリック経由でルーティングされます
ntfy.shリレー。承認プロンプト (短いコマンドを使用)
概要）そしてあなたのタップはあなたが指定したトピックを通過するので、そのことを学んだ人は誰でも
トピックはそれらのプロンプトを読んで答えることができます。長いランダムなトピックを使用し、
セルフホスト ntfy を使用するか、より強力な保証が必要な場合は ntfy アクセス トークンを使用します。

[切り捨てられた]

## Original Extract

Local, zero-dependency dashboard for Claude Code: live token usage and context, lost-session recovery, full-text search, and approve tool calls from your phone. - nikitadoudikov/claude-pulse

GitHub - nikitadoudikov/claude-pulse: Local, zero-dependency dashboard for Claude Code: live token usage and context, lost-session recovery, full-text search, and approve tool calls from your phone. · GitHub
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
nikitadoudikov
/
claude-pulse
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
10 Commits 10 Commits bin bin docs docs hooks hooks public public src src .gitignore .gitignore LICENSE LICENSE README.md README.md config.example.json config.example.json package.json package.json View all files Repository files navigation
A local dashboard for Claude Code that shows what Claude is doing, what it is spending, and lets you approve its tool calls from your phone. Zero dependencies, nothing leaves your machine.
Claude Code already writes every session to disk. Pulse reads those files (read only) and turns them into a live dashboard: token spend by hour, day and week, the context fill of your active session, an ambient view of Claude at work, full-text search across everything you have ever run, and a notification with Allow / Allow all / Deny buttons when Claude needs you, on your desktop or your phone. No account, no telemetry, no network calls.
Approve from your phone. A push with working Allow / Allow all / Deny buttons. No Wi-Fi setup, no IP, no open port: it works from anywhere, even on cellular.
Never lose a session. One command recovers your last session as a readable transcript, and Pulse auto-snapshots active ones, so a crash or a frozen laptop never costs you context.
See the spend. Live tokens and API-equivalent cost by hour, day, week, model and project, against budgets you set, with a phone alert when you cross one.
Ambient office. A full-screen view of a little mascot working, resting, or waiting on you, with a rough ETA. Quietly addictive on a second monitor.
Search everything. Full-text search across every session on disk, one click to the transcript.
Local and private. Reads ~/.claude read only, serves on 127.0.0.1 , zero dependencies, no telemetry.
Requires Node 18+. Run it with no install:
npx pulse-for-claude-code
Or clone it:
git clone https://github.com/nikitadoudikov/claude-pulse.git
cd claude-pulse
node bin/cli.js
Either way it opens http://127.0.0.1:4317 . To get desktop and phone
notifications and to approve tool calls, wire the hooks (one command, safe to
re-run):
claude-pulse install-hooks # adds the hooks to ~/.claude/settings.json
Then restart Claude Code, and you are set. Other options:
claude-pulse --port 4317 # change the port
claude-pulse --no-open # do not open the browser
Keep it running
Run in the foreground and Pulse dies when you close that terminal. To keep it
alive independently, run it in the background:
claude-pulse start # run detached, survives closing the terminal
claude-pulse status # is it running?
claude-pulse stop # stop it
claude-pulse restart # stop and start again
If your terminal crashes, claude-pulse start brings it back in one command,
and a background instance is not affected by the crash in the first place.
On macOS you can hand Pulse to the system so it starts at login and respawns
itself if it ever dies:
claude-pulse install-service # start at login, auto-restart
claude-pulse uninstall-service # remove it
Recover a lost session
Terminal crashed, laptop froze, hit a session limit? Nothing is lost: Claude
Code writes every session to disk as it happens. One command brings the last one
back, prints a recap and saves a readable transcript:
claude-pulse recover # the most recent session
claude-pulse recover 2 # the one before that
claude-pulse recover < id > # a specific session
It saves a light markdown file under ~/.claude-pulse/exports/ (a 15 MB log
becomes a ~180 KB file) and prints a link to read the full transcript in the
browser or on your phone. You can also open any session in the dashboard and use
open transcript / download .md .
While Pulse runs it also auto-snapshots every recently active session to
~/.claude-pulse/exports/snapshots/ (one file per session, rewritten only when
it changes). So the latest state is always on disk even if you never run
recover . Set snapshotMinutes to 0 in ~/.claude-pulse.json to turn it off.
To back up everything at once, claude-pulse export-all writes every session
into a single small gzipped markdown file, or use download all history on the
Sessions screen.
Lost where you did something? The Sessions screen has a search box that
scans every session on disk for a word or phrase and jumps you straight to the
transcript. It works from your phone too.
The simplest phone control is the ntfy notification itself: it carries working
Allow / Allow all / Deny buttons (see above), no network setup at all.
For a richer view, open http://<your-machine>:4317/phone on the same Wi-Fi
(needs bindLan: true ) to see what Claude is doing right now plus a Pause /
Resume button. Pausing stops Claude from running further tools until you
resume. Both need the PreToolUse hook wired.
┌──────────────┐ writes .jsonl ┌──────────────────────┐ SSE ┌──────────────┐
│ Claude Code │ ─────────────────▶ │ Pulse (read only) │ ───────▶ │ dashboard │
│ (terminal) │ │ 127.0.0.1:4317 │ │ + phone │
└──────┬───────┘ └──────────────────────┘ └──────────────┘
│
│ hooks: Notification · Stop · PreToolUse
▼
┌─────────────────────────────┐
│ ~/.claude-pulse/ │ pending approvals · decisions · events
└─────────────────────────────┘
Claude Code logs every session as JSONL under ~/.claude/projects/ . Each assistant
message carries a usage block (input, output and cache tokens) with a timestamp.
Pulse reads those files (read only), caches each file by modification time so
unchanged sessions are never re-parsed, and aggregates the numbers. The browser
gets live updates over Server-Sent Events. Three small hooks let Claude Code tell
Pulse when it needs you, when a turn ends, and when it wants to run a tool.
Notifications when Claude needs you
Claude Code can run a hook when it needs your attention. Point its Notification
event at the bundled script and Pulse will show a banner and fire a desktop
notification, even if the tab is in the background.
claude-pulse install-hooks # wires the hooks into ~/.claude/settings.json (safe to re-run)
claude-pulse uninstall-hooks # removes them
It backs up your settings once, merges next to any hooks you already have, and
never adds a duplicate. Restart Claude Code afterwards. To do it by hand instead,
add this to ~/.claude/settings.json (use the absolute path to your clone):
{
"hooks" : {
"Notification" : [
{
"matcher" : " " ,
"hooks" : [
{ "type" : " command " , "command" : " node /absolute/path/to/claude-pulse/hooks/notify-hook.js " }
]
}
],
"Stop" : [
{
"matcher" : " " ,
"hooks" : [
{ "type" : " command " , "command" : " node /absolute/path/to/claude-pulse/hooks/stop-hook.js " }
]
}
],
"PreToolUse" : [
{
"matcher" : " " ,
"hooks" : [
{ "type" : " command " , "command" : " node /absolute/path/to/claude-pulse/hooks/permission-hook.js " }
]
}
]
}
}
Keep claude-pulse running and you are set.
Approve tools from the dashboard (and your phone)
With the PreToolUse hook wired, when Claude wants to run something that needs
permission, an approval card appears in Pulse with Allow , Allow all and
Deny . Allow all stops asking for the rest of the run.
This is built to never hang Claude. Read only tools pass straight through, and if
Pulse is not running, has not heard from you within the approval timeout (60s by
default, set approvalTimeoutMs ), or hits any error, it falls back to the normal
terminal prompt. Nothing breaks if you ignore it. The phone push carries Allow ,
Allow all and Deny buttons.
To approve from your phone, you only need an ntfyTopic (below) and the ntfy
app. The push notification carries Allow , Allow all and Deny buttons, and
tapping one sends the answer back through ntfy to a private reply topic that
Pulse listens on. No same Wi-Fi, no IP, no open port: it works from anywhere,
even on cellular. Pulse only acts on a reply while it is actually waiting for
that request, so a stale notification can do nothing.
Claude wants to run a tool
│
▼
PreToolUse hook ──▶ Pulse ──push──▶ phone notification
│ │
│ tap "Allow"
│ │
│ answer returns over ntfy
│ │
▼ ▼
hook is still waiting ◀── decision ◀── Pulse (subscribed to the reply topic)
│
▼
hook returns "allow" ──▶ Claude runs the tool
Phone push (optional)
To get a push on your phone when Claude needs you or finishes, pick a hard to
guess topic name, install the free ntfy app and subscribe to
that topic, then set it in ~/.claude-pulse.json :
{ "ntfyTopic" : " claude-pulse-9f3a7c " }
With the hooks above wired, the Notification hook pushes when Claude is waiting
for you, and the Stop hook pushes when a turn finishes (debounced to 30s so a
back and forth does not spam you). Anyone who knows the topic can read it, so use
a random name.
If you set budgets (below), Pulse also pushes when a rolling window crosses 80%
then 100% of its budget, so you find out from your pocket, not by checking.
Copy config.example.json to ~/.claude-pulse.json and edit. Every field is optional.
{
"plan" : " max20 " ,
"contextLimit" : 200000 ,
"idleMinutes" : 10 ,
"approvalTimeoutMs" : 60000 ,
"budgets" : { "fiveHour" : 140 , "day" : 360 , "week" : 1100 }
}
About limits
Anthropic does not publish exact subscription limits, and they are usage based
rather than a fixed token count. Pulse cannot read your real plan ceiling, so the
budgets above are rough API-equivalent estimates you adjust to match what you
observe. The pro , max5 and max20 presets are starting points, not official
numbers. Token cost is estimated from public API list prices purely as a usage
proxy; subscription users do not pay per token.
Pulse is local-first and opt-in. Out of the box it binds to 127.0.0.1 only,
makes no outbound calls, has zero dependencies (no supply chain), and reads
~/.claude read only. Nothing leaves your machine and there is no analytics. Two
optional features change that, and both are off until you turn them on:
Phone push ( ntfyTopic ) routes through the public
ntfy.sh relay. Approval prompts (with a short command
summary) and your taps pass through a topic you name, so anyone who learns the
topic can read those prompts and answer them. Use a long random topic, and
self-host ntfy or use ntfy access tokens if you want stronger guarantees.

[truncated]
