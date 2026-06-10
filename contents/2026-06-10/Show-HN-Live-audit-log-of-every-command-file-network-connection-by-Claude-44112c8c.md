---
source: "https://github.com/yeet-src/claudefeed"
hn_url: "https://news.ycombinator.com/item?id=48468650"
title: "Show HN: Live audit log of every command, file, network connection by Claude"
article_title: "GitHub - yeet-src/claudefeed: Live audit log of every command, file, and network connection a Claude Code (or any matched) session makes, from the kernel. · GitHub"
author: "r3tr0"
captured_at: "2026-06-10T01:01:05Z"
capture_tool: "hn-digest"
hn_id: 48468650
score: 4
comments: 0
posted_at: "2026-06-09T22:26:05Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Live audit log of every command, file, network connection by Claude

- HN: [48468650](https://news.ycombinator.com/item?id=48468650)
- Source: [github.com](https://github.com/yeet-src/claudefeed)
- Score: 4
- Comments: 0
- Posted: 2026-06-09T22:26:05Z

## Translation

タイトル: HN を表示: クロードによるすべてのコマンド、ファイル、ネットワーク接続のライブ監査ログ
記事のタイトル: GitHub - yeet-src/claudefeed: クロード コード (または一致するもの) セッションが行うすべてのコマンド、ファイル、ネットワーク接続のカーネルからのライブ監査ログ。 · GitHub
説明: クロード コード (または一致するもの) セッションが行うすべてのコマンド、ファイル、およびネットワーク接続のカーネルからのライブ監査ログ。 -yeet-src/claudefeed

記事本文:
GitHub - yeet-src/claudefeed: クロード コード (または一致するもの) セッションが行うすべてのコマンド、ファイル、およびネットワーク接続のカーネルからのライブ監査ログ。 · GitHub
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
まだ-src
/
クロードフィード
公共
通知
nを変更するにはサインインする必要があります

通知設定
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
6 コミット 6 コミット アセット .gitignore .gitignore Makefile Makefile README.md README.md claudefeed.bpf.c claudefeed.bpf.c config.js config.js feed.js feed.js main.js main.js model.js model.js すべてのファイルを表示 リポジトリ ファイル ナビゲーション
クロード コードの tail -f。セッションが実行するすべてのコマンド、セッションが開くすべてのファイル、到達またはバインドするすべての TCP ポートがデコードされ、端末にライブでストリーミングされます。そのセッションのプロセス サブツリーをスコープとします。他の PID やノイズはありません。
claudefeed は、クロード コード セッションをライブのデコードされた監査ログに変換します。つまり、一致したセッションとそのサブツリー全体のすべての exec (完全な argv 付き)、すべての openat 、およびすべての送信またはバインドされた TCP ポートです。
兄弟の claudetree が誰がライブ プロセス ツリーとして実行されているかを示しますが、 claudefeed はツリーが約束した仲間であり、実行したことの尾部 -f です。
これをただ踏むことはできません。セッションはプロセスの移動ツリーです。Claude はシェルを生成し、シェルは git を生成し、git は ssh を生成します。すべての openat のシステム全体のフィードはファイアホースとなり、PID ごとのトレースではすべての子が見逃されます。 kernel 内の claudefeed フィルター: セッション tgid の追跡されたセットがプローブをゲートし、そのセットが exec 全体に自己伝播し、ノイズがユーザー空間に侵入することはありません。
カール -fsSL https://yeet.cx |しー
yeet github:yeet-src/claudefeed を実行します
手動インストールガイド | Linuxのみ
クロード コード セッションがボックス上のどこででも実行されれば、それだけです。claudefeed はライブ クロード プロセスを見つけ、その追跡セットをシードし、ストリーミングを開始します。 -- 以降はすべて claudefeed に渡されます。
yeet run github:yeet-src/claudefeed -- --before=open --secs=30 # コマンド + ネットワークのみ、30 秒
yeet github:yeet-src/claudefeed を実行します -- --onl

y=exec,conn # 「実行された内容」と「ダイヤルされた内容」のみ
yeet run github:yeet-src/claudefeed -- --match=node # 代わりにノード セッションを監査します
追跡セットに関する 60 秒の入門書
セッションの監査で難しいのは、イベントをキャプチャすることではありません。プロセス ツリーが成長したり縮小したりするときに、セッションのイベントのみをキャプチャすることです。 claudefeed は、セッション tgid の追跡されたハッシュ マップをカーネル内に保持します。ファイルとネットワーク プローブは、その中に含まれていない PID に対しては何も操作しません。セットは 3 つの方向から最新の状態を維持します。
1. JS がシードします。起動時の 1 つの yeet.graph クエリは、ライブ セッション プロセスとそのすべての子孫を検索し、1 つの updateBatch システムコールで追跡される tgid を書き込みます。そのため、アタッチ時にすでに実行されているプロセスのアクティビティもキャプチャされます。
2. exec はそれを自己伝播します。追跡対象のプロセス実行者が子である場合、その子はメンバーシップを継承します。新しいセッション (親が追跡されていないセッション) は、実行時に JS 側がプログラムの .data セクションにパッチするニードルと、実行されたプログラムのベース名を照合することによって捕捉されます。同じニードルがシードとカーネル側のキャッチの両方を駆動するため、何が重要であるかについては 2 つが常に一致します。
3. exit はそれをプルーニングします。スレッドグループのリーダーが離脱すると、その tgid が削除されます。 tgid によってキー設定された別の開始マップは、 exec → exit をブリッジするため、終了行はプロセスの生存期間を報告できます (アタッチ時にすでに生存しているプロセスにはスタンプがなく、不明な生存期間が報告されます)。
一致はプログラムのベース名に対してのみテストされ、コマンド ライン全体に対してテストされることはありません。パスまたは引数で claude に言及するだけのプロセスはセッションとして偽装されません。
claudefeed は、真の自律性を持って Claude Code (または任意のエージェント) を実行していて、それが何を行ったかではなく、実際に何をしたかの真実の記録を必要としている人のためのものです。
自律エージェントの実行が終了しました。何を命令するか

実際に実行されるのですが、どのような順序で実行されますか?
セッションがホストに接続したか、関係のないファイルを開いた。ワイヤーで捕まえてください。
長時間無人の実行にはトリップワイヤーが必要です。何か注目すべきことを行った瞬間にページングされます (「アラート」を参照)。
フォレンジックやレビューのために、クリーンで grep 可能でアーカイブ可能なセッションのログが必要です。
claudefeed — 「claude」セッションの監査 · 288 ライブからシードされた 9 プロセス · ストリーミング exec/exit/open/conn/listen …
16:11:22.023 3003897 スリープを終了します 0 を終了します
16:11:22.023 3003924 exec uname · /usr/bin/uname -a
16:11:22.024 3003924 exit uname exit 0 1ms 後に
16:11:22.024 3003925 exec カール · カール -s -m 3 http://example.com
16:11:22.033 3003925 コンカール → 104.20.23.154:80
16:11:22.053 3003925 終了カール 29 ミリ秒後に終了 0
各行は time、pid、class、detail です。 pid とそのプロセスのプログラム名はプロセスごとに安定した色を共有するため、アクティビティのバーストは 1 つのアクターとして読み取られます。同じ PID を再実行すると、その色が維持されます。コマンド ラインと開いたパス内では、フラグはミュートされ、パスはシアン色になるため、 bash -c '…' パイプラインは一目で判読できます。クラスは色キー付きのバッジです。
色は yeet のスタイル グローバル (標準の 16 色パレット、トゥルーカラーなし) から取得され、stdout が TTY ではない場合、プレーン テキストへの操作は行われません。そのため、パイプされた実行は grep またはアーカイブできるクリーンなプレーン テキスト ログになります。
コアは claudefeed.bpf.c と main.js にあります。
1 つの BPF オブジェクトは 3 つのトレースポイントと 2 つの kprobe を接続し、SEC() 名によって start() に自動的に接続されます。
3 つのマップはカーネルをユーザー空間に接続します。
追跡 — セッション tgid のハッシュ。 JS からシードされ、その後はカーネルによって維持されます。すべてのファイル/ネットワーク プローブをゲートします。
start — tgid によってキー化された HASH。exec → exit をブリッジするため、exit 行は有効期間を報告できます。
events — btf_struct (event) によってバインドされた RINGBUF、1 つのデコード

イベントごとのレコード。
ファイル
責任
main.js
配線: バインド/スタート、ニードルのパッチ、追跡されたシード、リングのストリーミング
config.js
yeet.args → 定数 ( match 、 secs 、のみ / 除く、イベントの種類)
モデル.js
sysgraph からセッション サブツリーをシードします。レコードフィールドをデコードする
フィード.js
1 つのレコードを色付きの監査行にフォーマットする
RingBuf.subscribe は、各レコードをデコードされたオブジェクトとしてストリーミングして返します。char[] フィールドは JS 文字列として、__u8[N] フィールド (生のアドレス) はバイト オブジェクトとして到着するため、argv はカーネル側で 1 つの char[] にスペース結合されますが、アドレスはバイナリ クリーンのままです。
syscall ラッパーではなく、トレースポイントと kprobe を使用する理由
メンバーシップ トリックは、カーネルが実行する瞬間に、実行内 (子に伝播するため) と出口内 (プルーニングするため) を起動する必要があります。コマンド ラインで PID を指定することなく、ツリー全体に対して一度に実行されます。トレースポイントと kprobe はまさにそれを提供し、カーネル内で追跡されるゲートは、Firehose がユーザー空間に到達しないことを意味します。
フィードはローカルですが、各行を出力する同じリングバッファ コールバックによってページングすることもできます。 yeet は Slack に投稿します。 yeet.cx/settings にログインし、ワークスペースに一度接続すると、yeet.alert が到達可能な任意のチャネルに送信できるようになります。
main.js の購読コールバックはすでにデコードされたすべてのレコードを保持しているため、アラートはその中の単なる分岐にすぎません。クロード セッションが関係のないホストにダイヤルアウトした瞬間を知るには、次のようにします。
// feed.js が使用するのと同じデコーダを取り込みます
import { cstr , fmtAddr } from "./model.js" ;
const WATCH = "yeet.cx" ; // IP、ポート、またはコマンドニードルも機能します
const sub = リングを待ちます。サブスクライブ ( async ( rec ) => {
const e = rec .イベント??記録;
統計。イベント++ ;
if(!SHOW.has(KIND_KEY[e.kind]))return;
コンソール。 log(line(e,Date.now()));
if ( e .kind === EVT_CONNECT ) {
const ピア = fmtAd

dr(e.family、e.addr、e.port);
if (peer . include (WATCH) ) {
まだ待ってください。アラート ( {
メソッド:「スラック」、
チャンネル : "#alerts" 、
text : `claude (pid ${ e . pid } ) が ${peer } ` に接続しました、
ブロック: [
{ type : "header" , text : { type : "plain_text" , text : "クロードが家に電話しました" } } ,
{
タイプ: "セクション" 、
フィールド: [
{ type : "mrkdwn" 、 text : `*Process:*\n ${ cstr ( e . comm ) } ( ${ e . pid } )` } 、
{ タイプ : "mrkdwn" 、テキスト : `*Peer:*\n ${ ピア } ` } 、
]、
} 、
]、
} ) ;
}
}
} ) ;
同じ形状がどのクラスにも適合します。cstr(e.cmdline) が rm -rf またはcurl … と一致する EVT_EXEC のアラートです。 sh 、または ~/.ssh/id_* に触れる EVT_OPEN 。コールバックにはデコードされたレコードがすでに手元にあります。述語と yeet.alert を追加するだけです。
BTF を備えた Linux カーネル — トレースポイント コンテキスト構造体、 task_struct 、およびネットワーク kprobes が読み取る sock / ソケット構造体に必要です。現在の Arch、Fedora、Ubuntu、および Debian のデフォルト。
yeet デーモン。特権付き BPF ロードを処理します。カール -fsSL https://yeet.cx | sh がインストールします。
claudefeed は強制ではなく観察可能性です。それは何が起こったのかを教えてくれます。何も起こらなくなるわけではありません。
claudefeed は監査ログです。セッションが何をしたかを報告します。何もブロックしたりサンドボックス化したりしません。 (アクションをブロックするには、お問い合わせください)
何が読み取られたか、何が書き込まれたかではなく、ファイルが開かれたことが記録されます。オープン イベントは openat からのみ発生します。古いオープン システムコールおよびメモリ マップされた I/O パスやその他の I/O パスは表示されていません。 (カスタムyeetスクリプトについてはお問い合わせください)
クロード自身のログを読んでみませんか?
これらは、クロードがそれが何をしたと考えているかを示しています。 claudefeed はカーネルを読み取るので、実際に execve 、 openat 、および TCP スタックを通過した内容を記録します。これには、生成されたシェル、 git 、およびサブプロセスが独自に実行したすべてのものが含まれます。
セッションが遅くなりますか?

わぁ？
意味のあるオーバーヘッドはありません。プローブは受動的なオブザーバーであり、カーネル内で追跡されるゲートは、リング バッファーへの書き込みが発生する前に、非セッション イベントをドロップします。
アタッチ後に開始したセッションはキャッチされますか?
はい。新しいセッションは、ベース名がニードルと一致するプログラムを実行した瞬間に捕捉されます。これがメンバーシップのトリックのカーネル側の半分です。
クロードにしか効かないのか？
いいえ、 --match は単なるプログラム名の針です。 --match=node 、 --match=python 、 --match=bash — 認識可能な名前で実行されるものはすべて同じように機能します。
データストリームをエクスポートできますか?
現在は組み込まれていません。フィードは stdout (またはパイプされた場合はプレーン テキスト) にレンダリングされるため、必要な --only= / --reason= フィルターを使用して claudefeed を実行し、出力をログ ファイルまたは選択した配送業者に保存するのが簡単な方法です。 HTTP 経由の JSON、Kafka トピック、S3 シンク、または SIEM パイプラインなどの構造化エクスポートの場合、main.js の Ring.subscribe コールバックを追加します。上記の Slack アラートの例と同じ形状は、どのシンクでも機能します。管理されたエクスポート パイプラインを設定するには、お問い合わせください。
make # bpftool 経由で vmlinux.h をダンプし、claudefeed.bpf.o をビルドします
Clang (BPF ターゲット) と bpftool に加えて、BTF を備えたカーネルが必要です。生成された include/vmlinux.h および *.bpf.o はビルド アーティファクトです。
GPL-2.0。 B

[切り捨てられた]

## Original Extract

Live audit log of every command, file, and network connection a Claude Code (or any matched) session makes, from the kernel. - yeet-src/claudefeed

GitHub - yeet-src/claudefeed: Live audit log of every command, file, and network connection a Claude Code (or any matched) session makes, from the kernel. · GitHub
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
yeet-src
/
claudefeed
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
6 Commits 6 Commits assets assets .gitignore .gitignore Makefile Makefile README.md README.md claudefeed.bpf.c claudefeed.bpf.c config.js config.js feed.js feed.js main.js main.js model.js model.js View all files Repository files navigation
tail -f for Claude Code. Every command a session runs, every file it opens, every TCP port it reaches or binds — decoded and streamed live to your terminal. Scoped to that session's process subtree. No other PIDs, no noise.
claudefeed turns a Claude Code session into a live, decoded audit log: every exec (with full argv), every openat , and every outbound or bound TCP port — for the matched session and its entire subtree.
Where its sibling claudetree paints who is running as a live process tree, claudefeed is the companion that tree promised: the tail -f of what they did .
You can't just strace this. A session is a moving tree of processes — Claude spawns a shell, the shell spawns git , git spawns ssh . A system-wide feed of every openat would be a firehose, and a per-PID trace misses every child. claudefeed filters in the kernel : a tracked set of session tgids gates the probes, the set self-propagates across exec , and the noise never crosses into userspace.
curl -fsSL https://yeet.cx | sh
yeet run github:yeet-src/claudefeed
Manual install guide | Linux only
With a Claude Code session running anywhere on the box, that's it — claudefeed finds the live claude processes, seeds its tracked set, and starts streaming. Everything after -- is passed to claudefeed:
yeet run github:yeet-src/claudefeed -- --except=open --secs=30 # commands + network only, 30s
yeet run github:yeet-src/claudefeed -- --only=exec,conn # just "what ran" and "what it dialed"
yeet run github:yeet-src/claudefeed -- --match=node # audit node sessions instead
A 60-second primer on the tracked set
The hard part of auditing a session isn't capturing events — it's capturing only the session's events, as its process tree grows and shrinks. claudefeed keeps a tracked hash map of session tgids in the kernel; the file and network probes are no-ops for any pid not in it. The set stays current from three directions:
1. JS seeds it. One yeet.graph query at startup finds the live session processes and all their descendants and writes the tgids into tracked in a single updateBatch syscall — so activity from processes already running when you attach is captured too.
2. exec self-propagates it. When a tracked process exec's a child, the child inherits membership. A fresh session — one whose parent isn't tracked — is caught by matching the exec'd program's basename against a needle the JS side patches into the program's .data section at runtime. The same needle drives both the seed and the kernel-side catch, so the two always agree on what counts.
3. exit prunes it. The thread-group leader leaving drops its tgid. A separate start map keyed by tgid bridges exec → exit , so the exit line can report how long the process lived (processes already alive at attach have no stamp and report an unknown lifetime).
Because match is tested against the program basename only — never the whole command line — a process that merely mentions claude in a path or argument doesn't masquerade as a session.
claudefeed is for anyone running Claude Code (or any agent) with real autonomy and wanting a ground-truth record of what it actually did — not what it said it did.
An autonomous agent run finished. What commands did it actually execute, and in what order?
A session reached out to a host or opened a file it had no business touching. Catch it on the wire.
A long unattended run needs a tripwire — get paged the moment it does something notable (see Alerting ).
You want a clean, grep-able, archivable log of a session for forensics or review.
claudefeed — auditing "claude" sessions · seeded 9 processes from 288 live · streaming exec/exit/open/conn/listen …
16:11:22.023 3003897 exit sleep exit 0
16:11:22.023 3003924 exec uname · /usr/bin/uname -a
16:11:22.024 3003924 exit uname exit 0 after 1ms
16:11:22.024 3003925 exec curl · curl -s -m 3 http://example.com
16:11:22.033 3003925 conn curl → 104.20.23.154:80
16:11:22.053 3003925 exit curl exit 0 after 29ms
Each line is time · pid · class · detail . The pid and that process's program name share a stable per-process color, so a burst of activity reads as one actor; re-exec of the same pid keeps its color. Within a command line and an opened path, flags are muted and paths are cyan , so a bash -c '…' pipeline is legible at a glance. The class is a color-keyed badge:
Colors come from yeet's style global (standard 16-color palette, no truecolor), which no-ops to plain text when stdout isn't a TTY — so a piped run is a clean plain-text log you can grep or archive.
The core is in claudefeed.bpf.c and main.js .
One BPF object attaches three tracepoints and two kprobes, auto-attached on start() by their SEC() names:
Three maps connect kernel to userspace:
tracked — HASH of session tgids. Seeded from JS, maintained by the kernel thereafter; gates every file/network probe.
start — HASH keyed by tgid, bridging exec → exit so the exit line can report a lifetime.
events — RINGBUF bound by its btf_struct ( event ), one decoded record per event.
file
responsibility
main.js
wiring: bind/start, patch the needle, seed tracked , stream the ring
config.js
yeet.args → constants ( match , secs , only / except , event kinds)
model.js
seed the session subtree from the sysgraph; decode record fields
feed.js
format one record into a colored audit line
RingBuf.subscribe streams each record back as a decoded object: a char[] field arrives as a JS string and an __u8[N] field (the raw address) as a byte object, so argv is space-joined kernel-side into one char[] while addresses stay binary-clean.
Why tracepoints and kprobes, not a syscall wrapper
The membership trick needs to fire inside exec (to propagate to children) and inside exit (to prune), at the moment the kernel does them — for the whole tree at once, with no PID named on the command line. Tracepoints and kprobes give exactly that, and the in-kernel tracked gate means the firehose never reaches userspace.
The feed is local — but the same ring-buffer callback that prints each line can also page you . yeet posts to Slack: log in at yeet.cx/settings , connect your workspace once, and yeet.alert can then send to any channel it can reach.
The subscribe callback in main.js already holds every decoded record, so an alert is just a branch inside it. To know the moment a Claude session dials out to a host it has no business touching:
// pull in the same decoders feed.js uses
import { cstr , fmtAddr } from "./model.js" ;
const WATCH = "yeet.cx" ; // an IP, a port, or a command needle work too
const sub = await ring . subscribe ( async ( rec ) => {
const e = rec . event ?? rec ;
stats . events ++ ;
if ( ! SHOW . has ( KIND_KEY [ e . kind ] ) ) return ;
console . log ( line ( e , Date . now ( ) ) ) ;
if ( e . kind === EVT_CONNECT ) {
const peer = fmtAddr ( e . family , e . addr , e . port ) ;
if ( peer . includes ( WATCH ) ) {
await yeet . alert ( {
method : "slack" ,
channel : "#alerts" ,
text : `claude (pid ${ e . pid } ) connected to ${ peer } ` ,
blocks : [
{ type : "header" , text : { type : "plain_text" , text : "Claude phoned home" } } ,
{
type : "section" ,
fields : [
{ type : "mrkdwn" , text : `*Process:*\n ${ cstr ( e . comm ) } ( ${ e . pid } )` } ,
{ type : "mrkdwn" , text : `*Peer:*\n ${ peer } ` } ,
] ,
} ,
] ,
} ) ;
}
}
} ) ;
The same shape fits any class: alert on an EVT_EXEC whose cstr(e.cmdline) matches rm -rf or curl … | sh , or an EVT_OPEN that touches ~/.ssh/id_* . The callback already has the decoded record in hand — you only add the predicate and the yeet.alert .
A Linux kernel with BTF — needed for the tracepoint context structs, task_struct , and the sock / socket structs the network kprobes read. Default on current Arch, Fedora, Ubuntu, and Debian.
The yeet daemon, which handles the privileged BPF load. curl -fsSL https://yeet.cx | sh installs it.
claudefeed is observability, not enforcement. It tells you what happened; it does not stop anything from happening.
claudefeed is an audit log. It reports what a session did; it does not block or sandbox anything. (to block actions, contact us )
It records that a file was opened, not what was read or written. open events come from openat only; the older open syscall and memory-mapped or other I/O paths are not shown. ( contact us for custom yeet scripts)
Why not just read Claude's own logs?
Those tell you what Claude thinks it did. claudefeed reads the kernel, so it records what actually crossed execve , openat , and the TCP stack — including everything the spawned shells, git , and subprocesses did on their own.
Does it slow the session down?
No meaningful overhead. The probes are passive observers and the in-kernel tracked gate drops non-session events before they ever cost a ring-buffer write.
Will it catch a session I start after attaching?
Yes. A fresh session is caught the moment it exec's a program whose basename matches the needle — that's the kernel-side half of the membership trick.
Does it only work for Claude?
No. --match is just a program-name needle. --match=node , --match=python , --match=bash — anything that exec's under a recognizable name works the same way.
Can I export the data stream?
Not built in today. The feed renders to stdout (or plain text when piped), so the quick path is to run claudefeed with whatever --only= / --except= filter you want and tee the output into a log file or your shipper of choice. For a structured export, say JSON over HTTP, a Kafka topic, an S3 sink, or a SIEM pipeline, the ring.subscribe callback in main.js is where you'd add it; the same shape as the Slack alerting example above works for any sink. To set up a managed export pipeline, contact us .
make # dumps vmlinux.h via bpftool, builds claudefeed.bpf.o
Needs clang (BPF target) and bpftool , plus a kernel with BTF. The generated include/vmlinux.h and *.bpf.o are build artifacts.
GPL-2.0. The B

[truncated]
