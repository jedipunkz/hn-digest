---
source: "https://github.com/bhpark1013/claudenews"
hn_url: "https://news.ycombinator.com/item?id=48674128"
title: "Show HN: Turn Claude Code's \"thinking \" wait into a dev news feed"
article_title: "GitHub - bhpark1013/claudenews: Fill Claude Code's agent wait time with dev news — Hacker News, GitHub Trending & per-language sources in your status line, auto-translated with short summaries. · GitHub"
author: "goodboy1013"
captured_at: "2026-06-25T14:57:35Z"
capture_tool: "hn-digest"
hn_id: 48674128
score: 2
comments: 2
posted_at: "2026-06-25T14:38:53Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Turn Claude Code's "thinking " wait into a dev news feed

- HN: [48674128](https://news.ycombinator.com/item?id=48674128)
- Source: [github.com](https://github.com/bhpark1013/claudenews)
- Score: 2
- Comments: 2
- Posted: 2026-06-25T14:38:53Z

## Translation

タイトル: Show HN: Claude Code の「思考」待機を開発ニュース フィードに変える
記事のタイトル: GitHub - bhpark1013/claudenews: Claude Code のエージェントの待ち時間を開発ニュースで埋める — ステータス ラインにハッカー ニュース、GitHub トレンド、および言語ごとのソースが表示され、短い概要が自動翻訳されます。 · GitHub
説明: Claude Code のエージェントの待ち時間を開発ニュース (ハッカー ニュース、GitHub トレンド、および言語ごとのソース) でステータス行に表示し、短い概要を自動翻訳します。 - bhpark1013/claudenews

記事本文:
GitHub - bhpark1013/claudenews: Claude Code のエージェントの待ち時間を開発ニュースで埋めます — ハッカー ニュース、GitHub トレンド、およびステータス行に言語ごとのソースが表示され、短い概要が自動翻訳されます。 · GitHub
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
bhpark1013
/
クロードニュース
公共
ノーティ

fications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
99 Commits 99 Commits .claude-plugin .claude-plugin .github/ workflows .github/ workflows docs docs plugin plugin web web .gitignore .gitignore README.md README.md install.sh install.sh install.sh uninstall navigation
claudenews — DevFeed for Claude Code
Switching to another task while Claude works costs you context.
Just staring at the terminal is boring. claudenews fills those short
waits with dev news — so the idle time isn't wasted and you never lose
あなたの場所。
Claudeが他のコマンドを処理している間に別のことをしましょう。
費用がかかり、ターミナルだけ見ていると思っていましたか？その短い時間を開発
ニュースで記入し、流れはそのままにしてください。
A Claude Code plugin that surfaces Hacker News , GitHub Trending , and
per-language dev & news sources in the status line during agent wait time —
auto-translated into your OS language, with short summaries that fade in as
Claude fetches them in the background.
The title is a clickable link, ▲ score / 💬 comments show for HN
items, and the trailing hint rotates through the commands (dropped
automatically if the terminal is too narrow). Korean / Japanese /
Chinese users see the same item auto-translated — no configuration:
/plugin marketplace add bhpark1013/claudenews
/plugin install claudenews@claudenews
/reload-plugins
/claudenews:setup
/reload-plugins activates the hooks and slash commands without a full
restart。 Then restart Claude Code once so the status line takes effect —
the next time the agent thinks, a news item rotates in.
No API key. T

翻訳と要約には既存のクロード コード セッションを使用します
（俳句、バックグラウンド価格）。 OpenAI キーや個別の請求はありません。
OS 言語を自動検出します。 macOS AppleLocale 、次に $LANG 、
次に Python ロケール。韓国語、日本語、中国語、スペイン語などはすべて言語なしで動作します。
構成。英語ユーザーにはタイトルが翻訳されずに表示されます。
既存のステータス行と共存します。 OMC HUD、git ステータス、カスタム
スクリプト — ~/.claudenews/config.json のparentStatusLine 経由でスクリプトを指定します
そして、claudenews は置き換えるのではなく、その下に追加します。
端末の幅に自動的にフィットします。 cmux → tmux → iTerm2 → stty so を検出
タイトルと概要は、静的な 120 ではなく、実際の列数に切り捨てられます。
GitHub リポジトリの概要は正確です。 GitHub リポジトリの説明が次の場合
空の場合、claudenews は GitHub 経由で README の最初の段落に戻ります
「GitHub でアカウントを作成して…に貢献」を表示する代わりの API。
独立したウィンドウはありません。すべてはクロード コード内でインラインで行われます
セッション — ソース リストは、ID によって読み取って切り替えるプレーンなメニューです。
ソース リストはメニューです: /claudenews:list を実行してすべてのソースを表示します
(ID + フラグ + オン/オフ)、ID によって切り替えます。クロードに尋ねることもできます
わかりやすい言葉 - 「ニュースソースを表示」、「이 뉴스 자세히 보여줘」、
「ニュース言語を韓国語に変更する」、「claudenews のフィードバックを送信する」、そして
適切なコマンドが自動的に実行されます。
キーボード ナビゲーション (オプション、macOS)
デフォルトでは、エージェントが考えている間、ステータス行が自動的にニュースを循環します。
自分で項目をステップ実行したい場合は、キー ナビゲーションをオンにします。
/claudenews:ナビオン
Ctrl+Shift+→ 次へ · Ctrl+Shift+← 前へ
キーを再マップします: ~/.claudenews/config.json に navKeys を設定します (任意の
ctrl / shift / cmd / alt コンボ + 前/次キー)、再実行します
/claudenews:nav on — 以下のパワーユーザー設定を参照してください。
ニュースになる

es global : すべてのクロード コード セッションで同じ項目が表示されるため、
ステップを踏むとそれらがすべて一緒に動きます。
iTerm2、Apple Terminal、WezTerm、kitty で動作します。他のアプリでは、
ショートカットは変更されないままになります (したがって、ブラウザ/エディタでテキストが選択されたままになります)。
鍵は小さなハンマースプーンのタップで捕らえられ、
そのため、ナビゲーションには Hammerspoon がインストールされている必要があります。
brew install --cask ハンマースプーン
唯一の手動手順は、Hammerspoon に対する 1 回限りの macOS アクセシビリティ許可です。
— macOS は初めてプロンプトを表示します。 「システム設定を開く」をクリックして切り替えます
それをオンにします。 nav on は Hammerspoon を再起動し、ステータス行の更新を次のように設定します。
1秒なので踏み込みが瞬時に感じられます。いつでも無効化（および自動回転の復元）できます。
/claudenews:nav をオフにします。
/claudenews:更新
/reload-プラグイン
/claudenews:update は最新バージョンをプルし、キャッシュを更新し、
ステータスラインランチャーをワンステップで更新します。 /reload-plugins その後
新しいフックとコマンド (クロード コードのランタイム アクション) をアクティブ化します。
プラグイン自体は実行できません）。 /claudenews:setup を再実行する必要はありません —
ステータス行は、インストールされている最新バージョンを自動追跡します。
ソースは、 /claudenews:list で選択したカタログです。
ハッカー ニュース — トップ 50 (デフォルトで、すべてのユーザー)
GitHub Trending — 過去 7 日間に作成された人気のリポジトリ (デフォルトですべてのユーザー)
言語ごとのネイティブ開発コミュニティ - OS 言語に一致するもの
自動的にオンになります。すべて誰でも切り替え可能です。
geeknews GeekNews — 韓国語 ( ko )
qiita Qiita 人気 — 日本語 ( ja ) · zenn 、hatena-it (オプトイン)
v2ex V2EX — 中国語 ( zh ) · infoq-cn (オプトイン)
jdh ハッカージャーナル — フランス語 ( fr )
heise-dev heise 開発者 — ドイツ語 ( de )
tabnews TabNews — ポルトガル語 / ブラジル ( pt )
グローバル英語 (オプトイン): ロブスター ロブスター、devto DEV コミュニティ
一般的な世界/国内ニュース (開発者向けではなくオプトイン)

発生、デフォルトではオフ):
cnn CNN、bbc BBC ニュース、nyt NYT ワールド、聯合ニュース、ハニ 한겨레
(ネイバーニュースには公開一般 RSS がありません。聯合ニュースは KR 通信フィードです)
プラグインを更新せずにサーバー側でさらに追加できる
プラグインは、最初の実行時に OS 言語から適切なデフォルトを選択します。
/claudenews:list は完全なメニューを出力します。 /claudenews:リスト cnn bbc hn
ID によって 1 つ以上を切り替えます。パブリック ソースは、
バックエンド;選択内容は、何をマージするかを指示するだけです。
パワーユーザー構成 ( ~/.claudenews/config.json )
手動で設定できるいくつかのオプションのノブ (残りはプラグインが管理します):
clientFeeds — バックエンドが到達できない、フェッチされた RSS/Atom フィードを追加します
代わりに自分のマシンから。 Reddit が事件の見出し: 403s
データセンターの IP ですが、www.reddit.com/r/<sub>/.rss は通常の IP アドレスから動作します。
機械。アイテムはローテーション/ナビゲーション/翻訳/要約を通じて正確に流れます
組み込みソースと同様に、ソースごとのコードはなく、URL のみが存在します。
追加する最も簡単な方法: /claudenews:list add r/<sub> (または質問してください)
クロード — "r/rust を追加" ); /claudenews:list rmfeed r/<sub> で削除されます。の
コマンドによってこれが書き込まれますが、手動で編集することもできます。
"クライアントフィード" : [
{ "名前" : " 👽 r/ClaudeAI " 、 "url" : " https://www.reddit.com/r/ClaudeAI/.rss " },
{ "name" : " 🐘 #rust " 、 "url" : " https://mastodon.social/tags/rust.rss " 、 "lang" : " en " }
】
navKeys — キーボードとナビゲーションのコンボを再マップします (デフォルトは ctrl+shift+←/→ )。セット
修飾子と prev/next キーを使用してから、/claudenews:nav を再実行して次を適用します。
"navKeys" : { "modifiers" : [ "ctrl" 、 "shift" ]、 "prev" : " left " 、 "next" : " right " }
修飾子は、ctrl/shift/cmd/alt (一緒に保持) のいずれかです。前へ / 次へ
は任意のキー名 ( left 、 right 、 [ 、 ] 、 …) です。シェルのコンボを維持する
(ctrl+shift ファミリー) は使用しません

— 実行できるものを選択します (例: cmd / alt
+矢印) は、ターミナル内でそのショートカットが失われることを意味します。
summaryColor — 概要行の SGR カラー (デフォルトは 38;5;245 、
読みやすい灰色）。例えば「37」プレーンホワイト、「38;5;250」ライター、「38;5;240」
暗くします。
parentStatusLine — ステータス行 claudenews がその下に連鎖します。セット
以前のものからセットアップ時に自動的に。
プロンプトもコードもキーストロークも収集されません
ニュースはプロンプト送信時に取得され、レートはセッションごとに 30 秒に制限されます
翻訳/要約はローカルのクロード コード セッションを使用します。コンテンツは決して使用されません。
記事のメタディスクリプションと
リポジトリ概要用の GitHub API
すべてのキャッシュは ~/.claudenews/ の下に存在します。
匿名インストール数: /claudenews:setup はバックエンドに 1 回 ping を実行します。
サーバーは 1 つのカウンターのみを増分します。IP、ユーザー エージェント、または
識別子が読み取られ、保存され、または記録されます。オフラインでは黙ってスキップするだけです。
フィードバック ( /claudenews:フィードバック ) は明示的でオプトインです。
入力したメッセージとプラグインのバージョンが送信および保存されます (IP なし)。
ユーザーエージェント、マシン情報、または識別子。
/claudenews:分解
/plugin 削除 クロードニュース
分解すると、インストール前のステータス行が復元されます (
バックアップはセットアップ時に作成されます)、 ~/.claudenews/ を削除します。
CDウェブ
npmインストール
npm 実行開発
バックエンドは、Vercel 上のステートレス Next.js 16 アプリです。 /api/ニュースインターリーブ
選択したソース (HN、GitHub Trending、および RSS フィード)、5 分間キャッシュされました
関数インスタンスごとのメモリ内。データベースもユーザーごとの追跡もありません —
唯一永続的な状態は、プライバシーに配慮した匿名カウンター (インストール
カウント、フィードバック）、IP/識別子は保存されていません。
Claude Code のエージェントの待ち時間を開発ニュース (ハッカー ニュース、GitHub トレンド、およびステータス ラインに表示される言語ごとのソース) で埋め、短い概要を自動翻訳します。
そこには

読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Fill Claude Code's agent wait time with dev news — Hacker News, GitHub Trending & per-language sources in your status line, auto-translated with short summaries. - bhpark1013/claudenews

GitHub - bhpark1013/claudenews: Fill Claude Code's agent wait time with dev news — Hacker News, GitHub Trending & per-language sources in your status line, auto-translated with short summaries. · GitHub
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
bhpark1013
/
claudenews
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
99 Commits 99 Commits .claude-plugin .claude-plugin .github/ workflows .github/ workflows docs docs plugin plugin web web .gitignore .gitignore README.md README.md install.sh install.sh uninstall.sh uninstall.sh View all files Repository files navigation
claudenews — DevFeed for Claude Code
Switching to another task while Claude works costs you context.
Just staring at the terminal is boring. claudenews fills those short
waits with dev news — so the idle time isn't wasted and you never lose
your place.
Claude가 다른 명령을 처리하는 동안 다른 일을 하자니 컨텍스트 스위칭
비용이 들고, 터미널만 보고 있자니 심심하셨죠? 그 짧은 시간을 개발
뉴스로 채우고, 흐름은 그대로 유지하세요.
A Claude Code plugin that surfaces Hacker News , GitHub Trending , and
per-language dev & news sources in the status line during agent wait time —
auto-translated into your OS language, with short summaries that fade in as
Claude fetches them in the background.
The title is a clickable link, ▲ score / 💬 comments show for HN
items, and the trailing hint rotates through the commands (dropped
automatically if the terminal is too narrow). Korean / Japanese /
Chinese users see the same item auto-translated — no configuration:
/plugin marketplace add bhpark1013/claudenews
/plugin install claudenews@claudenews
/reload-plugins
/claudenews:setup
/reload-plugins activates the hooks and slash commands without a full
restart. Then restart Claude Code once so the status line takes effect —
the next time the agent thinks, a news item rotates in.
No API key. Translation and summary use your existing Claude Code session
(Haiku, background-priced). No OpenAI key, no separate billing.
Auto-detects your OS language. macOS AppleLocale , then $LANG ,
then Python locale. Korean, Japanese, Chinese, Spanish, etc. all work without
configuration. English users get titles untranslated.
Coexists with your existing status line. OMC HUD, git status, custom
scripts — point at them via parentStatusLine in ~/.claudenews/config.json
and claudenews appends underneath rather than replacing.
Auto-fits your terminal width. Detects cmux → tmux → iTerm2 → stty so
titles & summaries truncate to the actual column count, not a static 120.
GitHub repo summaries are accurate. When a GitHub repo's description is
empty, claudenews falls back to the README's first paragraph via the GitHub
API instead of showing "Contribute to … by creating an account on GitHub".
No separate windows. Everything happens inline in your Claude Code
session — the source list is a plain menu you read and toggle by id.
The source list is the menu: run /claudenews:list to see every source
(id + flag + on/off) and toggle by id. You can also just ask Claude in
plain language — "show my news sources", "이 뉴스 자세히 보여줘",
"change the news language to Korean", "send claudenews feedback" — and
the right command runs automatically.
Keyboard navigation (optional, macOS)
By default the status line auto-rotates through news while the agent thinks.
If you'd rather step through items yourself, turn on key navigation:
/claudenews:nav on
ctrl+shift+→ next · ctrl+shift+← previous
Remap the keys: set navKeys in ~/.claudenews/config.json (any
ctrl / shift / cmd / alt combo + prev/next keys), then rerun
/claudenews:nav on — see Power-user config below.
News becomes global : every Claude Code session shows the same item, so
stepping moves them all together.
Works in iTerm2, Apple Terminal, WezTerm, kitty . In any other app the
shortcut is left untouched (so it still selects text in your browser/editor).
The key is captured by a tiny Hammerspoon tap,
so nav on needs Hammerspoon installed:
brew install --cask hammerspoon
The only manual step is a one-time macOS Accessibility grant for Hammerspoon
— macOS pops the prompt the first time; click Open System Settings and toggle
it on. nav on restarts Hammerspoon for you and sets the status-line refresh to
1s so stepping feels instant. Disable it (and restore auto-rotation) anytime
with /claudenews:nav off .
/claudenews:update
/reload-plugins
/claudenews:update pulls the latest version, refreshes the cache, and
updates the status-line launcher in one step. /reload-plugins then
activates the new hooks and commands (a Claude Code runtime action a
plugin can't do itself). You never need to re-run /claudenews:setup —
the status line auto-tracks the newest installed version.
Sources are a catalog you choose from with /claudenews:list :
Hacker News — top 50 (on by default, all users)
GitHub Trending — popular repos created in the last 7 days (on by default, all users)
Per-language native dev communities — the one matching your OS language
turns on automatically; all are toggleable for anyone:
geeknews GeekNews — Korean ( ko )
qiita Qiita 人気 — Japanese ( ja ) · zenn , hatena-it (opt-in)
v2ex V2EX — Chinese ( zh ) · infoq-cn (opt-in)
jdh Journal du hacker — French ( fr )
heise-dev heise Developer — German ( de )
tabnews TabNews — Portuguese / Brazil ( pt )
Global English (opt-in): lobsters Lobsters, devto DEV Community
General world / national news (opt-in, not dev-focused, off by default):
cnn CNN, bbc BBC News, nyt NYT World, yonhap 연합뉴스, hani 한겨레
(Naver News has no public general RSS; Yonhap is the KR wire-service feed)
More can be added server-side without a plugin update
The plugin picks sensible defaults on first run from your OS language.
/claudenews:list prints the full menu; /claudenews:list cnn bbc hn
toggles one or more by id. Public sources are fetched and cached by the
backend; your selection just tells it what to merge.
Power-user config ( ~/.claudenews/config.json )
A few optional knobs you can set by hand (the plugin manages the rest):
clientFeeds — add any RSS/Atom feed the backend can't reach, fetched
from your own machine instead. Reddit is the headline case: it 403s
datacenter IPs, but www.reddit.com/r/<sub>/.rss works from a normal
machine. Items flow through rotation / nav / translation / summary exactly
like built-in sources — there's no per-source code, just URLs.
Easiest way to add one: /claudenews:list add r/<sub> (or just ask
Claude — "add r/rust" ); /claudenews:list rmfeed r/<sub> removes it. The
command writes this for you, and you can also edit it by hand:
"clientFeeds" : [
{ "name" : " 👽 r/ClaudeAI " , "url" : " https://www.reddit.com/r/ClaudeAI/.rss " },
{ "name" : " 🐘 #rust " , "url" : " https://mastodon.social/tags/rust.rss " , "lang" : " en " }
]
navKeys — remap the keyboard-nav combo (default ctrl+shift+←/→ ). Set
the modifiers and the prev/next keys, then rerun /claudenews:nav on to apply:
"navKeys" : { "modifiers" : [ " ctrl " , " shift " ], "prev" : " left " , "next" : " right " }
modifiers is any of ctrl / shift / cmd / alt (held together); prev / next
are any key names ( left , right , [ , ] , …). Keep to a combo the shell
doesn't use (the ctrl+shift family) — picking one it does (e.g. cmd / alt
+arrow) means losing that shortcut inside the terminal.
summaryColor — SGR color for the summary lines (default 38;5;245 , a
readable gray). e.g. "37" plain white, "38;5;250" lighter, "38;5;240"
dimmer.
parentStatusLine — the status line claudenews chains underneath; set
automatically at setup from whatever you had before.
No prompts, no code, no keystrokes collected
News is fetched on prompt submit, rate-limited to 30s per session
Translation/summary uses your local Claude Code session — content never
leaves your machine except to fetch the article meta description and the
GitHub API for repo summaries
All caches live under ~/.claudenews/
Anonymous install count: /claudenews:setup pings the backend once .
The server only increments a single counter — no IP, user-agent, or
identifier is read, stored, or logged. Offline just skips it silently.
Feedback ( /claudenews:feedback ) is explicit and opt-in: only the
message you type and the plugin version are sent and stored — no IP,
user-agent, machine info, or identifier.
/claudenews:teardown
/plugin remove claudenews
teardown restores whatever status line you had before install (from the
backup created at setup time) and removes ~/.claudenews/ .
cd web
npm install
npm run dev
Backend is a stateless Next.js 16 app on Vercel. /api/news interleaves
your selected sources (HN, GitHub Trending, and RSS feeds), cached 5 min
in-memory per function instance. No database and no per-user tracking —
the only persisted state is privacy-safe anonymous counters (install
count, feedback) with no IP/identifier ever stored.
Fill Claude Code's agent wait time with dev news — Hacker News, GitHub Trending & per-language sources in your status line, auto-translated with short summaries.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
