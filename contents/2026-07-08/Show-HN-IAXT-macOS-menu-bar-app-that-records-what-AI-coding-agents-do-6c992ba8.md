---
source: "https://iaxt.com/"
hn_url: "https://news.ycombinator.com/item?id=48830666"
title: "Show HN: IAXT – macOS menu-bar app that records what AI coding agents do"
article_title: "IAXT: AI you can trust. | Audit AI coding agents on macOS"
author: "laurencoral"
captured_at: "2026-07-08T12:16:03Z"
capture_tool: "hn-digest"
hn_id: 48830666
score: 1
comments: 1
posted_at: "2026-07-08T11:44:27Z"
tags:
  - hacker-news
  - translated
---

# Show HN: IAXT – macOS menu-bar app that records what AI coding agents do

- HN: [48830666](https://news.ycombinator.com/item?id=48830666)
- Source: [iaxt.com](https://iaxt.com/)
- Score: 1
- Comments: 1
- Posted: 2026-07-08T11:44:27Z

## Translation

タイトル: Show HN: IAXT – AI コーディング エージェントの動作を記録する macOS メニューバー アプリ
記事のタイトル: IAXT: 信頼できる AI。 | macOS 上の AI コーディング エージェントを監査する
説明: IAXT は、AI コーディング エージェント用の macOS フライト レコーダーです。 Claude Code、Cursor、Aider、その他のエージェントが実行するすべてのコマンドと、彼らが触れるすべてのファイルが表示され、セッション終了後に確認できます。無料ダウンロード。

記事本文:
IAXT: 信頼できる AI。 | macOS 上の AI コーディング エージェントを監査する
IAXT
仕組み
ユースケース
スクリーンショット
よくある質問
IAXTを入手
インテリジェント エージェント実行トラッカー (IAXT)
AI コーディングを確認する
エージェントは実際にそうしました。
ブロックはありません。ワークフローの変更はありません。個人向けのクラウドアップロードはありません。ただ、
ローカル監査証跡は、コーディング セッションの終了後に確認できます。
仕組み
論文
AI コーディング エージェントは、マシン上でコマンドを実行し、ファイルを変更します。
幅広い権限と高速性。
2 つのことが、これに追いつくのを難しくしています。
疲労: 長いセッションで許可ダイアログが表示される
私たちのほとんどは、コントロールしているかのような錯覚に陥り、クリックスルーしてしまいます。
プロンプト インジェクション: 攻撃者は次の命令を非表示にします。
通常のテキスト内に AI が侵入し、攻撃はますます頻繁になっています
そしてより洗練されたものに。
IAXTはブロックしません。 IAXTは制限しません。 IAXT ウォッチ、属性、および
覚えているので、それが疲れたクリックであっても、隠れた指示であっても、
今日、ほとんどのチームが答えられない質問に答えることができます。
AIは実際に何をしたのですか？
単独の開発者が自分のマシンを監査するか、創設者が
チーム全体の監査証跡。 IAXT は両方に適合し、それらは重要です
さまざまなことについて。
安心
自分のマシン。
アプリをインストールし、コーディングに戻ります。 IAXT がメニューに加わります
バー、すべての AI エージェントのアクティビティをローカルに記録し、
完了すると、30 秒間のレビュー画面が表示されます。
事故を捕捉し（AI がリポジトリの半分を削除）、フラグを立てます
まれに、意図的な exfil (プロンプトによって挿入される README)
SSH キーでcurl -X POSTをトリガーしました)、そして
それ以外の時間は顔にかからないようにします。
01
無料でダウンロードできます。
アカウントもサインアップもアップセルもありません。 DMG をドラッグ アンド ドロップし、1 回起動すれば完了です。
02
マシンからは何も残りません。
テレメトリーはありません。分析はありません。 [アップデートの確認] を明示的にクリックしない限り、ネットワーク呼び出しはゼロです。ダ

ta は ~/Library/Logs/IAXT/ にあり、 sqlite3 で開くことができ、 rm -rf で削除できます。
03
目立った減速はありません。
受動的観察のみ — IAXT は何もブロック、傍受、変更することはありません。エージェントはその存在を知りませんし、ビルド時間も知りません。
macOS 13 Ventura以降。 Appleシリコンとインテル。 ～2MB。
公証された DMG、Apple によって検証済み
監査証跡
デューデリジェンスが求めます。
エンジニアは IAXT をインストールします - 個人と同じアプリ
余分な摩擦はありません。 1 日 1 回、レビューの概要
(生のログではない) 中央のエンドポイントにプッシュします
あなたがコントロールします。セキュリティ リーダー、CTO、または創設者は、
エンジニアごとのロールアップ。
顧客が安全性を審査するとき、投資家は義務を負う
勤勉さ、または買収者の技術監査が「どうやってやっていますか？」と尋ねます。
AI エージェントのリスクを管理しますか?」と尋ねて、ダッシュボードを開きます。それが
答えてください。今日、これを示すことができるチームはほとんどありません。
部屋の中で目立つようになります。
私たちはチーム全体で Claude Code、Cursor、Aider を使用しています。
すべてのセッションはローカルに記録されます。レビュー層のイベント —
永続化メカニズム、認証情報へのアクセス、漏洩
パターン — 自動的にフラグが付けられ、次のようにロールアップされます。
私たちのセキュリティレビュー。先週のレポートはこちらです。
— 顧客、投資家、買収者に伝えること
01
エンジニアにとって摩擦のないもの。
個々の層と同じメニューバー アプリ。一度インストールすれば、忘れてしまいます。ワークフローの変更、IDE プラグイン、プロンプトはありません。
02
生のログではなく、概要。
集計されたレビュー データのみが各エンジニアのマシンから残ります。完全なフォレンジックトレイルはローカルに残ります。
03
セルフホストまたはマネージド。
エンドポイントを独自のインフラ上で実行するか、弊社にホストさせてください。プライベート ベータ中のデザイン パートナー価格。
チーム エンドポイントは招待専用ですが、次のように強化されています。
デザインパートナーを設立。あなたのチームについて教えてください。
連絡させていただきます。
3つのこと
IAXTは黙ってやります。
コマンドの実行、ファイルの作成または変更

d、パッケージのインストール、git 操作、cron エントリ、エージェントの起動。すべてのアクションは、それを作成したツールに起因するものであり、確認済み、可能性が高い、または可能です。
注目に値するアクションを表すゴールド ストライプ — プロジェクト外のファイル システムへの書き込み、予期しないネットワーク呼び出し、スケジュールされたタスクの変更。単に一見の価値があるものにはバイオレット。
AI の使用パターンの毎日の概要。セッションカード、​​統計、注目アイテム。すべてのエージェントを一度に監視するため、1 つのツールだけでなく、クロード コード、カーソルなど、すべてのツールにわたって実際に AI を使用してコーディングする方法を明確に把握できます。チームレビュー用の CSV エクスポート。すべてローカルで、テレメトリーやクラウドはありません。
誰が手を伸ばすのか
こんな記録。
AI のコーディング内容に対する公平な観察者
エージェントが実際にあなたのマシン上で実行したことです。事故を捕まえて、
あなたの行動を信頼せずに、本来のものではないまれなアクションを見つけます。
長いセッションの思い出。
エージェントが全体でどのように使用されているかを可視化
ワークフローの摩擦がゼロのチーム。誰もがインストールします
同じ静かなアプリ。ブラインドの代わりにレビュー面が得られます
スポット。
AI エージェントのリスクが実際に存在するという証拠
監視: セキュリティレビューで指摘できる監査証跡、または
デューデリジェンス。レビュー層のイベントに自動的にフラグが付けられます。
macOS アプリの実際のスクリーンショット。ビューを選択して内容を確認してください
を示し、それがなぜ重要なのかを説明します。
何かが必要になるまで、メニューバーで静かにしてください。
IAXT はメニュー バー内に存在し、メニューバーからは外されます。
観察しながらの様子。そのアイコンは単純な円です。
すべてが正常でモニタリングが正常な場合は白
アクティブ、価値のあるものにフラグを立てた場合は紫色
あなたのレビュー。メニューから監視を一時停止できます。
メイン ウィンドウを開くか、Finder でログを表示します。
記録されたものはすべて Mac に保存されます。
単一の SQLite データベース、完全な
任意の SQLite ブラウザで検査可能

いつでも削除可能です。
システム全体には何も散在していません。
あなたのマシンから何も離れることはありません。
注目に値するものから毎日を始めましょう。
「概要」には、IAXT が考えるあらゆる内容が表示されます。
検討する価値があります。これらは
必ずしも問題があるわけではありません: それは重要です
AI コーディング エージェントのように、一見する価値のある変更
システムの変更、ログインの追加
項目、または機密ファイルに触れた場合。 1 つ
クリックすると、アクション ログ内のそれらの行に直接移動します。 3
タイルはその日、AI セッション、アクション、ファイルを要約します
変更されたため、その間に何が起こったかを常に知ることができます
働いていました。
実際に AI をどのように活用しているかをご覧ください。
概要は、あなた自身の AI を静かに映し出す鏡でもあります。
使用方法。その日のレビュー項目に加えて、あなたの状況をグラフ化します。
一日の活動リズム、
最も混雑する時間帯とファイルが変更された場所
ほとんどの場合、時間の経過とともに週ごとにロールアップされます
パターンと合計。本当に便利な方法です
クロード コードやカーソルなどのツールにどれだけ依存しているかを理解します。
そしていつ。
レビューが必要なものを確認し、それを任意の AI にコピーします。
アクション ログ (レビュー) にフィルタリングされ、最も高い
層。ここで AI エージェントは 2 つの永続性を残しました
メカニズム: シェル起動ファイルを変更しました
そして、起動エージェントを作成しました。両方とも実行されます
またまた自分たちで。任意の行を右クリックして選択します
行をコンテキストとしてコピーして、表示されている内容を正確に取得します。
紫色のボックス: クリーンで自己記述的な要約
誰が、どこで、何をしたのか、なぜフラグが立てられたのか。
セカンドオピニオンとして任意の LLM に貼り付けます。つまり、IAXT
AI の査読者、監査証跡となる
選択したツールを使用して再確認できます。
フラグが立っているのは中間層です: では珍しいです
コンテキストですが、多くの場合は正当なものであり、ちょっと見てみる価値があります
アラームよりも。ここで、Claude Code はファイルをダウンロードしました。
ネットワークにはcurlがあり、起動エージェントは
削除されました。見てください

Who 列: すべてのアクションは次のとおりです
原因となった正確なエージェント セッションに起因する
それ、自信の丸薬で（確認済み、
おそらく、可能性があります) したがって、IAXT がどれほど確実であるかがわかります。ノイズはもう
なくなってしまえば、見る価値のあるものは 1% 程度しか残っていないのです。
IAXT は、あなたの行動に関する忠実な事後記録を提供します。
AI コーディング エージェントがあなたのマシン上で実行しました。まったく異なる 2 つの
状況によってはそれが価値のあるものになります。
善意：承認疲れ。あなたは承認します
仕事中は機密性の高いアクションが必要ですが、何時間も何日もコーディングが必要です
疲労と脱感作をもたらします。クリックするのは人間です
日常から外れることを許可する、または自分の何かに手を振る
完全には理解できませんでした。 IAXT では、落ち着いて元に戻ることができます。
後で、実際に何に同意したかを確認してください。
悪意があります: 即時注入。攻撃者が隠れる
AI が読み取るコンテンツ内の指示。以下は実際のメールです。
本文は無害なオンボーディングリマインダーのように見えますが、
ソースは、処理する AI アシスタントを対象としたコマンドを隠します。
というメッセージ。
Windows をターゲットにしており、ここでは実行されませんでした。しかし、そのテクニックは、
リアルでより洗練されています。 IAXT は受動的です: 実行します
ブロックではなく、記録します。 Mac 上のエージェントの場合
このような隠された指示に従って行動したことがあれば、わかるでしょう
セッション後のアクション ログに記録される内容とまったく同じです。
個人層: ゼロ。テレメトリーはありません。いいえ
分析。アカウントがありません。アプリが実行できる唯一のネットワーク呼び出し
make は「アップデートの確認」メニュー項目です - 1 つ
GitHub へのリクエストは、クリックした場合にのみ行われます。
企業レベル: 1 日 1 回、レビューの概要
(エージェントごとのカウント、フラグ付き/レビューアクションの合計 - 生データなし
コマンド) は、制御するエンドポイントに送信されます。私たちには何もありません。
クロードコード、カーソル、エイダー、コーデックス、ウィンドサーフィン、キロコード、
OpenCode、副操縦士、Cody。新しいエージェントはリクエストに応じて追加されます —
GitHub で問題を開く
ツールのプロセス名を追加します。
今のところはそうではない。 IAXTは

現在 macOS のみ (macOS 13 Ventura)
以降)、macOS ネイティブのイベント ストリームに基づいて構築されています。窓
バージョンは別の作業になります。に基づいて検討するかもしれません
要求。
いいえ、そんなことはあり得ません。 App Store アプリは Apple 内で実行する必要があります
サンドボックス。各アプリを残りのアプリから隔離します。
システム。 IAXT の仕事全体は、他のプロセスが何をしているかを監視することです
マシン全体でコマンドラインを読み、従う
独自のフォルダーを超えたファイル アクティビティ。サンドボックスブロック
まさにそれ。したがって、IAXT は最も本格的な Mac を出荷する方法です
開発者ツールは、署名され公証された DMG として、
直接ダウンロードし、最初に Apple の Gatekeeper によって検証される
打ち上げ。
あなたが感じるような方法ではありません。 IAXT は macOS のネイティブをサブスクライブします
イベント ストリーム (FSEvents、kqueue、定期的 sysctl) - OS は
すでにこの作業を行っており、イベントの約 95% が削除されます。
何かがディスクに触れる前に構築されます。遅れはありません
エディターでは、ビルドが遅くなったり、エージェントが途切れたりすることはありません。
正直な注意点はバッテリーです。 IAXT は時計があるから
バックグラウンドで継続的に動作するため、バッテリでラップトップが動作する可能性があります
もう少し早くダウンする必要があります。常に稼働しているツールにはこのコストがかかります。
私たちは、サイズを小さく保つために懸命に努力してきました。適応ポーリングにより、
何も起こらないときはすぐに後退し、完全に一時停止します
Mac がスリープしている間に。電源に関しては問題ありません。バッテリー上で
それは数パーセントであり、崖ではありません。
IAXT.appをゴミ箱にドラッグし、
rm -rf ~/Library/Logs/IAXT を実行して SQLite を削除します
データベースと監査ログ。それだけです — 隠しファイルもありません
LaunchAgent をアンロードします。
Apple の Endpoint Security フレームワークは、より多くの情報をキャッチします (ファイルの読み取り、
すべての実行)、ただし手動レビューが必要
Apple からの権利とシステム拡張機能のインストール フロー。
IAXT v1 は、配布上の摩擦によりユーザー空間で実行されます。
最後のイベントをすべて把握することよりも重要です。 ESレベル
検出はロードマップにあります。
どちらも利用可能になります。

プライベートベータ期間中、私たちはそれをホストします
反復の速度。自己ホスト型は一般提供前に出荷されます。必要な場合は
初日から自己ホスト型、電子メール
[メールで保護されています]
—それに応じて優先順位を付けます。
現在、個々のアプリはクローズドソースです。エンタープライズ向けおよび
ビジネスのお客様は、ご要望に応じて喜んでソースを共有いたします
セキュリティ監査とデューデリジェンスのため — 電子メール
[email protected] 。
監査はしません
私たちには見えないもの。
透明性は双方向を切り取ります。これらは構造的なものであると主張する人は誰でも
そうでなければ、あなたに何かを売りつけていることになります。監査ストーリーは次の場合にのみ機能します。
私たちが過度に約束することはないと信じているので、ここに完全なリストを示します。
IAXT は意図的に中間に位置します。何もないよりもはるかに優れています。
フル EDR よりも軽量です。カーネル拡張機能もシステム拡張機能もありません。
マシンの速度を低下させるものは何もありません。個人にとっても、社会にとっても権利
重さや負担のない実際の可視性を求めるスタートアップやチーム
エンドポイント セキュリティ ソフトウェアの侵入性。
ベンダー インフラストラクチャ上のリモート サンドボックス。ローカルプロセスなし
あなたのマシンに触れます。 IAXT は、リモート対応の場合にバナーを表示します。
アプリが実行されているため、死角が存在することがわかります。
ブラウザーホスト型のコードジェネレーター。ローカルでは何も実行されません。
設計上の範囲外 — ベンダーのチャット記録
サイトが唯一の監査対象です。
IAXTr

[切り捨てられた]

## Original Extract

IAXT is a macOS flight recorder for AI coding agents. See every command Claude Code, Cursor, Aider and other agents run, and every file they touch, reviewable after the session ends. Free download.

IAXT: AI you can trust. | Audit AI coding agents on macOS
IAXT
How it works
Use cases
Screenshots
FAQ
Get IAXT
Intelligent Agent eXecution Tracker (IAXT)
See what your AI coding
agent actually did.
No blocking. No workflow changes. No cloud upload for individuals. Just a
local audit trail you can review after the coding session ends.
How it works
The thesis
AI coding agents run commands and change files on your machine, with
broad permissions and at high speed.
Two things make that hard to keep up with.
Fatigue: over a long session the permission dialogs
blur into an illusion of control , and most of us click through.
Prompt injection: attackers now hide instructions for
the AI inside ordinary text, and the attacks keep getting more frequent
and more sophisticated.
IAXT doesn't block. IAXT doesn't restrict. IAXT watches, attributes, and
remembers, so whether it was a tired click or a hidden instruction, you
can answer the question most teams can't answer today:
what did the AI actually do?
Solo developer auditing your own machine, or a founder rolling an
audit trail up to the whole team. IAXT fits both — and they care
about different things.
Peace of mind for
your own machine.
Install the app, go back to coding. IAXT sits in your menu
bar, logs every AI agent's activity locally, and surfaces a
30-second review screen when you're done.
Catches the accidents (the AI deleted half the repo), flags
the rarer intentional exfil (a prompt-injected README
triggered a curl -X POST on your SSH key), and
stays out of your face the rest of the time.
01
Free to download.
No account, no signup, no upsell. Drag-and-drop DMG, launch once, done.
02
Nothing leaves your machine.
No telemetry. No analytics. Zero network calls unless you explicitly click Check for Updates. Data sits in ~/Library/Logs/IAXT/ , openable with sqlite3 , deletable with rm -rf .
03
No noticeable slowdown.
Passive observation only — IAXT never blocks, intercepts, or modifies anything. Your agents don't know it exists, and neither will your build times.
macOS 13 Ventura or later. Apple Silicon & Intel. ~2 MB.
Notarized DMG, verified by Apple
The audit trail
due diligence asks for.
Your engineers install IAXT — same app as the individual
tier, no extra friction. Once a day, a review summary
( not raw logs ) pushes to a central endpoint
you control. Your security lead, CTO, or founder sees a
per-engineer roll-up.
When a customer's security review, an investor's due
diligence, or an acquirer's tech audit asks "how do you
manage AI-agent risk?", you open the dashboard. That's the
answer. Few teams can show this today, which is why having
it now sets you apart in the room.
We use Claude Code, Cursor, and Aider across the team.
Every session is logged locally. Review-tier events —
persistence mechanisms, credential access, exfiltration
patterns — are flagged automatically and rolled up to
our security review. Here's last week's report.
— What you tell a customer, investor, or acquirer
01
Frictionless for engineers.
Same menu-bar app as the individual tier. Install once, forget. No workflow changes, no IDE plugin, no prompts.
02
Summaries, not raw logs.
Only aggregated review data leaves each engineer's machine. Full forensic trail stays local.
03
Self-host or managed.
Run the endpoint on your own infra, or let us host it. Design-partner pricing during private beta.
The team endpoint is invite-only while we harden it with
founding design partners. Tell us about your team and
we'll be in touch.
Three things
IAXT does quietly.
Commands run, files created or modified, packages installed, git operations, cron entries, launch agents. Every action attributed to the tool that made it — confirmed, likely, or possible.
A gold stripe for actions that deserve attention — filesystem writes outside the project, unexpected network calls, changes to scheduled tasks. Violet for things simply worth a look.
A daily Overview of your AI usage patterns. Session cards, stats, attention items. Because it watches every agent at once, it also becomes a clear picture of how you actually code with AI across all your tools, Claude Code, Cursor, and the rest, not just one. CSV export for team review. Everything local, no telemetry, no cloud.
Who reaches for
a record like this.
An impartial observer of what your AI coding
agents actually did on your machine. Catch the accidents, and
spot the rare action that does not belong, without trusting your
memory of a long session.
Visibility into how agents are used across the
team , with zero workflow friction. Everyone installs
the same quiet app; you get a review surface instead of a blind
spot.
Evidence that AI-agent risk is actually
watched: an audit trail you can point to in a security review or
due diligence, with review-tier events flagged automatically.
Real screenshots of the macOS app. Pick a view to see what it
shows and why it matters.
Quiet in your menu bar, until something needs you.
IAXT lives in your menu bar and stays out of the
way while it observes. Its icon is a simple circle:
white when everything is normal and monitoring is
active, violet when it has flagged something worth
your review. From the menu you can pause monitoring ,
open the main window, or show your logs in Finder .
Everything it records stays on your Mac in a
single SQLite database , fully
inspectable with any SQLite browser and deletable any time.
Nothing is scattered across your system, and
nothing ever leaves your machine .
Start each day with what deserves a look.
The Overview opens with anything IAXT thinks is
worth reviewing . These are
not necessarily problems : they are the important
changes worth a glance, like an AI coding agent
modifying your system , adding a login
item , or touching a sensitive file . One
click takes you straight to those rows in the Action Log. Three
tiles summarize the day, AI sessions, actions, and files
changed , so you always know how much happened while you
were working.
See how you actually work with AI.
The Overview is also a quiet mirror of your own AI
usage . Beyond the day's review items, it charts your
activity rhythm across the day, your
busiest hours , and where files changed
most , then rolls it up over time into weekly
patterns and totals. It is a genuinely useful way to
understand how much you lean on tools like Claude Code and Cursor,
and when.
See what needs review, then copy it into any AI.
The Action Log, filtered to Review , the highest
tier. Here an AI agent left two persistence
mechanisms : it modified a shell startup file
and created a launch agent , both of which would run
again on their own. Right-click any row and choose
Copy row as context to get exactly what you see in
the violet box: a clean, self-describing summary of
who did what, where, and why it was flagged .
Paste it into any LLM for a second opinion , so IAXT
becomes a peer reviewer for your AI , an audit trail
you can double-check with the tool of your choice.
Flagged is the middle tier: unusual in
context but often legitimate , worth a quick look rather
than an alarm. Here Claude Code downloaded files over the
network with curl , and a launch agent was
removed. Look at the Who column: every action is
attributed to the exact agent session that caused
it, with a confidence pill ( confirmed ,
likely, possible) so you know how sure IAXT is. The noise is already
gone, only the ~1% worth seeing remains.
IAXT gives you a faithful, after-the-fact record of what your
AI coding agents did on your machine. Two very different
situations make that worth having.
Good faith: approval fatigue. You approve
sensitive actions as you work, but hours and days of coding
bring fatigue and desensitization. It is human to click
allow out of routine, or to wave through something you
did not fully understand. IAXT lets you go back, calmly and
later, and see what you actually agreed to.
Bad faith: prompt injection. Attackers hide
instructions inside content an AI reads. Below is a real email:
the body looks like a harmless onboarding reminder, but its
source hides a command aimed at any AI assistant that processes
the message.
It targeted Windows and never ran here. But the technique is
real and getting more sophisticated. IAXT is passive: it does
not block, it records . If an agent on your Mac
ever acted on a hidden instruction like this, you would see
exactly what it did in the Action Log, after the session.
Individual tier: zero. No telemetry. No
analytics. No account. The only network call the app can
make is the Check for Updates menu item — one
request to GitHub, only when you click it.
Company tier: once a day, a review summary
(counts per agent, flagged/review action totals — no raw
commands) goes to the endpoint you control. Nothing to us.
Claude Code, Cursor, Aider, Codex, Windsurf, Kilo Code,
OpenCode, Copilot, Cody. New agents are added on request —
open an issue on GitHub
with your tool's process name and we'll add it.
Not for the moment. IAXT is macOS-only today (macOS 13 Ventura
or later), built on macOS-native event streams. A Windows
version would be a separate effort; we may consider it based on
demand.
No, and it can't be. App Store apps must run inside Apple's
sandbox, which walls each app off from the rest of the
system. IAXT's whole job is to watch what other processes do
across your machine, read their command lines, and follow
file activity beyond its own folder. The sandbox blocks
exactly that. So IAXT ships the way most serious Mac
developer tools do, as a signed and notarized DMG you
download directly, verified by Apple's Gatekeeper on first
launch.
Not in a way you'll feel. IAXT subscribes to macOS's native
event streams (FSEvents, kqueue, periodic sysctl) — the OS is
already doing this work — and drops roughly 95% of events by
construction before anything touches disk. No lag in your
editor, no slower builds, no stutter in your agents.
The honest caveat is battery. Because IAXT watches
continuously in the background, a laptop on battery may run
down a little sooner — any always-on tool has this cost.
We've worked hard to keep it small: adaptive polling that
backs right off when nothing is happening, and a full pause
while your Mac sleeps. On power it's a non-issue; on battery
it's a few percent, not a cliff.
Drag IAXT.app to the Trash, then
rm -rf ~/Library/Logs/IAXT to remove the SQLite
database and audit logs. That's it — no hidden files, no
LaunchAgent to unload.
Apple's Endpoint Security framework catches more (file reads,
every exec ) but requires a manual-review
entitlement from Apple and a system-extension install flow.
IAXT v1 runs in user-space because distribution friction
matters more than catching every last event. ES-level
detection is on the roadmap.
Both will be available. During private beta we host it for
speed of iteration. Self-hosted ships before GA. If you need
self-hosted from day one, email
[email protected]
— we'll prioritize accordingly.
The individual app is closed-source today. For enterprise and
business customers, we're happy to share the source on request
for security auditing and due diligence — email
[email protected] .
We won't audit
what we can't see.
Transparency cuts both ways. These are structural — anyone claiming
otherwise is selling you something. The audit story only works if
you believe we won't overpromise, so here's the full list.
IAXT sits deliberately in the middle: far more than nothing, far
lighter than a full EDR. No kernel extension, no system extension,
nothing to slow the machine down. Right for individuals, and for the
startups and teams that want real visibility without the weight and
intrusiveness of endpoint security software.
Remote sandboxes on vendor infrastructure. No local process
touches your machine. IAXT shows a banner when a remote-capable
app is running so you know the blind spot exists.
Browser-hosted code generators. Nothing executes locally.
Out of scope by design — the chat transcript on the vendor's
site is the only audit surface.
IAXT r

[truncated]
