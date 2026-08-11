---
source: "https://github.com/karanb192/hindcast"
hn_url: "https://news.ycombinator.com/item?id=49258409"
title: "Show HN: Hindcast (search, replay, and resume any Claude Code session on Mac)"
article_title: "GitHub - karanb192/hindcast: Browse, search & replay every Claude Code session on your Mac. Local-first session viewer, timeline scrubber & cost tracker (ccusage alternative). · GitHub"
author: "karanb192"
captured_at: "2026-08-11T14:13:13Z"
capture_tool: "hn-digest"
hn_id: 49258409
score: 4
comments: 1
posted_at: "2026-08-11T13:54:15Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Hindcast (search, replay, and resume any Claude Code session on Mac)

- HN: [49258409](https://news.ycombinator.com/item?id=49258409)
- Source: [github.com](https://github.com/karanb192/hindcast)
- Score: 4
- Comments: 1
- Posted: 2026-08-11T13:54:15Z

## Translation

タイトル: HN を表示: Hindcast (Mac 上のクロード コード セッションを検索、再生、再開)
記事のタイトル: GitHub - karanb192/hindcast: Mac 上のすべてのクロード コード セッションを参照、検索、再生します。ローカルファーストのセッション ビューアー、タイムライン スクラバー、コスト トラッカー (Ccusage の代替)。 · GitHub
説明: Mac 上ですべてのクロード コード セッションを参照、検索、再生します。ローカルファーストのセッション ビューアー、タイムライン スクラバー、コスト トラッカー (Ccusage の代替)。 - karanb192/ハインドキャスト

記事本文:
GitHub - karanb192/hindcast: Mac 上のすべてのクロード コード セッションを参照、検索、再生します。ローカルファーストのセッション ビューアー、タイムライン スクラバー、コスト トラッカー (Ccusage の代替)。 · GitHub
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
カランブ192
/
後者
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
14 コミット 14 コミット Assets lib lib scripts scripts src src .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING。

md ライセンス ライセンス README.md README.mdindex.htmlindex.htmlmain.jsmain.jspackage-lock.jsonpackage-lock.jsonpackage.jsonpackage.jsonpreload.jspreload.jsstyles.cssstyles.css すべてのファイルを表示 リポジトリ ファイルのナビゲーション
これまでに実行したすべての Claude Code セッション。インデックス付き、検索可能、スクラブ可能。
3 週間前のセッションを見つけて、エージェントが行ったことを確認し、
ターミナルで 1 回貼り付けて再開します。ローカルファースト: を読み取ります。
トランスクリプト クロード コードはすでに ~/.claude/projects/ に書き込んでいます。
Mac からは何も残されません。
私自身のアーカイブ、2026 年 8 月時点: 65 プロジェクトにわたる 114 セッション、半分
ギガバイトの JSONL、1,700 万の出力トークン、4 月にまで遡ります。ハインドキャスト
M3 Pro では約 4 秒ですべてのインデックスがコールドインデックス化され、インデックスが保持されます。
新しいセッションが始まると新鮮です。
Hindcast-explainer.mp4
インストール
醸造タップ karanb192/タップ
醸造トラストkaranb192/タップ
brew install --cask ハインドキャスト
brew install ripgrep # 検索エンジン;それなしでは、⌘K はタイトルのみに一致します
xattr -rd com.apple.quarantine /Applications/Hindcast.app
今のところAppleシリコンのみ。アプリはそうではないため、xattr ステップが必要です。
まだ公証されていません (開発者 ID 登録中)。それはあっという間に消えてしまう
今後の署名入りリリース。ダウンロードをご希望ですか?からダメージを取得します
をリリースします。
今日実行する価値のあることの 1 つは、Claude Code がファイルの古さによってトランスクリプトを整理することです。
~/.claude/settings.json の cleanupPeriodDays、デフォルトは 30 日。ハインドキャスト
インデックス付けできるのは、まだ存在するもののみです。私のものは60に設定されているため、アーカイブは
上記はまだ 4 月に達しています。これらのセッションは再開され続けていたため、
ファイルは若いままでした。剪定者があなたの履歴にアクセスする前に、あなたの値を上げてください。
⌘K あらゆるものを検索 : すべてのトランスクリプト、ツールの全文を検索
出力が含まれています。結果を開くと、テープが一致した場所に直接ジャンプします
イベント。
テープ：イベントごとの抽選

n は目盛りとして表示されます (真鍮 = あなた、アイボリー/インク = クロード、
紫 = 思考、青緑 = ツール、レンガ = エラー)。クリックまたはドラッグしてスクラブします。
再生ヘッドはスクロールに従います。
アーカイブから再開: すべてのセッション ヘッダーには再開チップがあり、
cd -- <project dir> && claude --resume <session id> をコピーし、準備完了
ターミナルに貼り付けます。このコマンドはその場で編集できるテンプレートであるため、
シェルエイリアスは機能します（それがあなたのものであれば、cd {cwd} && ccr {sessionId}）。
The Ledger : 告発スタイル
使用量とコストのビュー: モデルごとの日ごと、週ごと、月ごとの表
内訳、公開されている API レートでの推定ドルコスト、および過去 7 日間
図。
インスタントフィルタリング: セッションリストをタイトルごとにゼロで入力してフィルタリングします。
キーストロークのラグ、日付プリセット、カスタム範囲、およびモデルの複数選択。
フィルターはリストとアーカイブ統計の両方を推進します。
エクスポート : 任意のセッションをマークダウンまたは自己完結型 HTML (アクティブ ブランチ、
画像はインライン化されます)、 ~/Downloads/Hindcast Exports に保存されます。
サブエージェント リール: セッションはサブエージェントのトランスクリプトをリストします。
subagents/workflows/wf_*/ の下にワークフローでネストされたもの。それぞれがそのように開きます
バックリンク付きの独自のテープ。
ダーク、ライト、オートのテーマが持続するため、ウィンドウが右側にペイントされます
最初のフレームからの色。
git クローン https://github.com/karanb192/hindcast.git
CDの後キャスト
npmインストール
npmスタート
Electron のバイナリがインストール中にダウンロードに失敗した場合 (npm スクリプト保護)、
ノード node_modules/electron/install.js を 1 回実行してから、 npm start を実行します。
適切なアプリとしてインストールするには、npm run Pack によって Hindcast.app がビルドされます。
dist/mac-arm64/ ;それを /Applications にドラッグします。
インデクサー ( lib/scanner.js ): すべてのトップレベルの *.jsonl をストリーミングします。
トランスクリプト、セッション タイトルを抽出します (custom-title > ai-title > first
プロンプト）、タイムスタンプ、モデル、トークンの使用状況、およびツールの数。結果は
ファイル mtime + サイズ + for によってキャッシュされます

アプリのユーザーデータディレクトリ内のマットバージョン、
また、fs.watch の再インデックスにより、作業中にリストが最新の状態に保たれます。
検索 ( lib/search.js ): にシェルアウトします。
すべてのトランスクリプトにわたる ripgrep
(引用符または
バックスラッシュ)、一致する行のみを再読み取りしてスニペットを作成します。
トランスクリプトはリストではなくツリーです。分岐したブランチを巻き戻します。視聴者
トランスクリプトの最後のプロンプト行からのleafUuidポインタに従います
(最後のメッセージに戻ります)、ライブ会話のみを表示します。
隠蔽された巻き戻されたイベントの数を報告します。値下げをレンダリングし、折りたたみ可能
思考、インラインスクリーンショット結果を含むツールカード、貼り付けられた画像、
圧縮マーカー、およびモデルフォールバックノート。
使用法 dedup : 1 つの API 応答が複数の JSONL 行にまたがり、それぞれが
同じ使用オブジェクトを繰り返し、セッションを再開またはフォークした場合の履歴のコピー
逐語的に。したがって、使用状況はメッセージ ID によってグローバルに重複排除されます。
ファイル、サブエージェントのトランスクリプトが含まれており、ccuseage がとっているのと同じアプローチです。
キャッシュ読み取りの料金は 0.1 倍、5 分間のキャッシュ書き込みは 1.25 倍、1 時間の書き込みです
2x ( lib/pricing.js )。コストは API に相当する価値の推定値であり、そうではありません。
請求書。
トランスクリプト形式は正式にはクロード コードの内部にあるため、解析は
意図的に寛容: 不明な線種とブロック タイプはスキップされます。
決して致命的ではありません。
「機械作業のための閲覧室」暖かい暗室/暖かい紙パレット、
ディスプレイタイプはニューヨークセリフ、機械音声はSFモノ、真鍮1個
アクセント。ネオンはありません。
Hindcast は ~/.claude/projects/ を読み取り、正確に 2 つの場所に書き込みます。
独自の: ユーザーデータ ディレクトリ (インデックス キャッシュ、テーマ、設定) および
~/Downloads/Hindcast [エクスポート] をクリックするとエクスポートされます。決して書き込むことはありません
~/.claude/ は、独自のネットワーク要求を行わず、テレメトリもありません。
設定するための API キーはありません。

決してクロードに電話をかけないでください：それは読みます
ディスク上のファイル、これがすべてのトリックです。 2026年を行ったり来たり
サードパーティエージェントツールのサブスクリプション認証
(最新ラウンド)
決してそれに適用されませんでした。
開いているセッションでは、クリックして離れて戻るまで、新しいメッセージは表示されません。
意図的に: バックグラウンド更新によってスクロール位置が移動してはなりません。の
セッションリストはどちらの場合でもライブのままです。
巻き戻された (放棄された) ブランチでの検索一致によりセッションが開始されますが、セッションは開始されません
隠しイベントにジャンプします。
アクティブなセッションのライブストリーミングまたはモニタリング: 端末はすでに
ストリームを示します。 Hindcast は読み取り専用のアーカイブのままです。
オーケストレーションとチーム/クラウド同期: ローカルファーストおよび読み取り専用、
永久に。
署名済み + 公証済みビルド: 署名されていない DMG は本日出荷されます。開発者ID
登録は保留中であり、それとともに xattr ステップは終了します。
ripgrep をバンドルするので検索
セットアップゼロで動作します。
オプトインのボールトにより、
アーカイブは、Claude Code の 30 日間の自動クリーンアップ後も残ります。
オンデバイスのセマンティック検索
同じ⌘Kの後ろにあります。
より多くのアイデアが生きています
問題点 ;
CONTRIBUTING.md には基本ルールが記載されています。ハインドキャストが掘ったら
あなたがすでに諦めていたセッションを再開するとき、スターは他の人がそれを見つけるのに役立ちます。
Mac 上ですべての Claude Code セッションを参照、検索、再生します。ローカルファーストのセッション ビューアー、タイムライン スクラバー、コスト トラッカー (Ccusage の代替)。
hindcast.karanbansal.in トピックス
Readme MIT ライセンス
1 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Browse, search & replay every Claude Code session on your Mac. Local-first session viewer, timeline scrubber & cost tracker (ccusage alternative). - karanb192/hindcast

GitHub - karanb192/hindcast: Browse, search & replay every Claude Code session on your Mac. Local-first session viewer, timeline scrubber & cost tracker (ccusage alternative). · GitHub
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
karanb192
/
hindcast
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
14 Commits 14 Commits assets assets lib lib scripts scripts src src .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md index.html index.html main.js main.js package-lock.json package-lock.json package.json package.json preload.js preload.js styles.css styles.css View all files Repository files navigation
Every Claude Code session you've ever run. Indexed, searchable, scrubbable.
Find that session from three weeks ago, scrub through what the agent did,
and resume it in your terminal with one paste. Local-first: it reads the
transcripts Claude Code already writes to ~/.claude/projects/ ,
and nothing leaves your Mac.
My own archive, as of August 2026: 114 sessions across 65 projects, half a
gigabyte of JSONL, 17M output tokens, reaching back to April. Hindcast
cold-indexes all of it in about 4 seconds on an M3 Pro, then keeps the index
fresh as new sessions land.
hindcast-explainer.mp4
Install
brew tap karanb192/tap
brew trust karanb192/tap
brew install --cask hindcast
brew install ripgrep # search engine; without it ⌘K matches titles only
xattr -rd com.apple.quarantine /Applications/Hindcast.app
Apple Silicon only for now. The xattr step is needed because the app isn't
notarized yet (Developer ID enrollment in progress); it disappears in an
upcoming signed release. Prefer a download? Grab the DMG from
Releases .
One thing worth doing today: Claude Code prunes transcripts by file age,
cleanupPeriodDays in ~/.claude/settings.json , default 30 days. Hindcast
can only index what still exists. Mine is set to 60, which is why the archive
above still reaches April: those sessions kept getting resumed, so their
files stayed young. Raise yours before the pruner gets to your history.
⌘K search across everything : full text over every transcript, tool
output included. Opening a result jumps the tape straight to the matched
event.
The tape : every event drawn as a tick (brass = you, ivory/ink = Claude,
violet = thinking, teal = tools, brick = errors). Click or drag to scrub;
the playhead follows your scroll.
Resume from the archive : every session header has a resume chip that
copies cd -- <project dir> && claude --resume <session id> , ready to
paste into a terminal. The command is a template you can edit in place, so
shell aliases work ( cd {cwd} && ccr {sessionId} if that's your thing).
The Ledger : a ccusage -style
usage and cost view: per-day / per-week / per-month tables with per-model
breakdowns, estimated dollar cost at published API rates, and a last-7-days
figure.
Instant filtering : type-to-filter the session list by title with zero
keystroke lag, plus date presets, a custom range, and a model multi-select.
Filters drive both the list and the Archive stats.
Export : any session to markdown or self-contained HTML (active branch,
images inlined), into ~/Downloads/Hindcast Exports .
Subagent reels : sessions list their subagent transcripts, including
workflow-nested ones under subagents/workflows/wf_*/ ; each opens as its
own tape with a back link.
Dark, light, and auto themes , persisted so the window paints the right
color from the first frame.
git clone https://github.com/karanb192/hindcast.git
cd hindcast
npm install
npm start
If Electron's binary fails to download during install (npm script protection),
run node node_modules/electron/install.js once, then npm start .
To install it as a proper app, npm run pack builds Hindcast.app into
dist/mac-arm64/ ; drag it to /Applications.
Indexer ( lib/scanner.js ): streams every top-level *.jsonl
transcript, extracts the session title (custom-title > ai-title > first
prompt), timestamps, models, token usage, and tool counts. Results are
cached by file mtime + size + format version in the app's user-data dir,
and an fs.watch re-index keeps the list current while you work.
Search ( lib/search.js ): shells out to
ripgrep across all transcripts
(also matching the JSON-escaped spelling of queries containing quotes or
backslashes), then re-reads only the matching lines to build snippets.
Transcripts are trees , not lists: rewinds fork branches. The viewer
follows the leafUuid pointer from the transcript's last-prompt line
(falling back to the last message), shows only the live conversation, and
reports how many rewound events it hid. Renders markdown, collapsible
thinking, tool cards with inline screenshot results, pasted images,
compaction markers, and model-fallback notes.
Usage dedup : one API response spans several JSONL lines that each
repeat the same usage object, and resumed or forked sessions copy history
verbatim. So usage is deduplicated globally by message id across every
file, subagent transcripts included, the same approach ccusage takes.
Cache reads priced at 0.1x, 5-minute cache writes at 1.25x, 1-hour writes
at 2x ( lib/pricing.js ). Costs are estimates of API-equivalent value, not
an invoice.
The transcript format is officially internal to Claude Code, so parsing is
deliberately tolerant: unknown line types and block types are skipped,
never fatal.
"A reading room for machine work." Warm darkroom / warm paper palettes,
New York serif for display type, SF Mono for the machine's voice, one brass
accent. No neon.
Hindcast reads ~/.claude/projects/ and writes to exactly two places, both
its own: its user-data dir (index cache, theme, settings) and
~/Downloads/Hindcast Exports when you click export. It never writes into
~/.claude/ , makes no network requests of its own, and has no telemetry.
There is no API key to configure because it never calls Claude: it reads
files on disk, and that's the whole trick. The 2026 back-and-forth over
subscription auth for third-party agent tools
( latest round )
never applied to it.
An open session doesn't show new messages until you click away and back.
Deliberate: a background refresh must never move your scroll position. The
session list stays live either way.
A search match on a rewound (abandoned) branch opens the session but can't
jump to the hidden event.
Live streaming or monitoring of active sessions : your terminal already
shows the stream. Hindcast stays a read-only archive.
Orchestration and team/cloud sync : local-first and read-only,
permanently.
Signed + notarized builds : an unsigned DMG ships today; Developer ID
enrollment is pending, and the xattr step dies with it.
Bundle ripgrep so search
works with zero setup.
An opt-in vault so your
archive survives Claude Code's 30-day auto-cleanup.
On-device semantic search
behind the same ⌘K.
More ideas live in
Issues ;
CONTRIBUTING.md covers the ground rules. If Hindcast dug
up a session you'd already given up on, a star helps other people find it.
Browse, search & replay every Claude Code session on your Mac. Local-first session viewer, timeline scrubber & cost tracker (ccusage alternative).
hindcast.karanbansal.in Topics
Readme MIT license Contributing
1 fork Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
