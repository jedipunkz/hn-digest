---
source: "https://github.com/only-cli/oc"
hn_url: "https://news.ycombinator.com/item?id=49367419"
title: "Show HN: Turning websites into micro CLIs for Claude Code to save on tokens"
article_title: "GitHub - only-cli/oc: Turn any website into a compact CLI tailored for AI agents. Browse the web in hundreds of tokens, not tens of thousands. · GitHub"
image: "https://opengraph.githubassets.com/246b240eec01972894dc87573a9c1e553ce1fcf6c926d477a146cc8a8092476c/only-cli/oc"
author: "hyes"
captured_at: "2026-08-19T22:14:57Z"
capture_tool: "hn-digest"
hn_id: 49367419
score: 3
comments: 0
posted_at: "2026-08-19T21:26:15Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Turning websites into micro CLIs for Claude Code to save on tokens

- HN: [49367419](https://news.ycombinator.com/item?id=49367419)
- Source: [github.com](https://github.com/only-cli/oc)
- Score: 3
- Comments: 0
- Posted: 2026-08-19T21:26:15Z

## Translation

タイトル: Show HN: Web サイトをクロード コードのマイクロ CLI に変換してトークンを節約する
記事のタイトル: GitHub - Only-cli/oc: あらゆる Web サイトを AI エージェント向けに調整されたコンパクトな CLI に変えます。数万ではなく、数百のトークンで Web を閲覧します。 · GitHub
説明: あらゆる Web サイトを AI エージェント向けに調整されたコンパクトな CLI に変えます。数万ではなく、数百のトークンで Web を閲覧します。 - のみ-cli/oc

記事本文:
GitHub - Only-cli/oc: あらゆる Web サイトを AI エージェント向けに調整されたコンパクトな CLI に変えます。数万ではなく、数百のトークンで Web を閲覧します。 · GitHub
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
のみ-cli
/
oc
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
52 コミット 52 コミット フォルダーとファイル
.github .github clis clis docs docs skill/only-cli skill/only-cli src src テスト テスト .gitignore .gitignore CONTRIBUTING.md CON

TRIBUTING.md README.md README.md llms.txt llms.txt package-lock.json package-lock.json package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Web サイトを AI エージェントのコマンド ライン インターフェイスに変えます。 oc open <url> はページを取得し、生の HTML やスクリーンショットの代わりにコンパクトな番号付きビューを返すため、Claude Code、Codex、Antigravity などのエージェントはトークンを焼くことなく閲覧できます。また、実際のブラウザと同じようにページと通信することで、一部のサイトで単純なフェッチャーを停止するブロックも回避します。
$ oc open news.ycombinator.com
#ハッカーニュース
[1] Show HN: 小さな CSV ツールキットを構築しました
[2] 312 コメント
...
アクション: <n> を実行します | <n> を読む |次へ |生の
$ oc は 1 を実行します
一般的なページには、数万のマークアップ トークンが含まれます。上のビューは数百に収まります。サイトごとのアダプター、ブラウザ拡張機能、デーモンは必要ありません。
このリポジトリを読んでいる LLM の場合、llms.txt は短縮バージョンです。
npm install -g @only-cli/oc
ノード 20 以上が必要です。リクエストは impers を介して Chrome になりすます。 impers が利用できない場合はネイティブフェッチに戻ります。
エージェントの指示ファイル (CLAUDE.md、AGENTS.md、または同等のもの) に 1 行を追加します。
Web ページからコンテンツが必要な場合は、生の HTML を取得する代わりに、npx @only-cli/oc open <url> を実行します。 npx @only-cli/oc --help を 1 回実行して、コマンドを学習します。
Claude Code ユーザーは、代わりにスキルをインストールできます。skills/only-cli/ を .claude/skills/ にコピーするか、npx skill addonly-cli/oc を実行します。これは、skills.sh を介して Cursor、Codex、Copilot などでも機能します。
セットアップをまったく行わなくても機能します。npx @only-cli/oc はグローバル インストールなしで実行され、レンダリングごとに --help とアクション: 行を通じて独自のコマンドを教えます。
oc open <url> 番号付きのアクションを含むページを取得してレンダリングします
oc do <n> 番号付きリンク [n] をたどるか、テキストの場合は [n] を読み取ります
oc find <クエリ>

すでに開いているページに文字列が表示される場所
oc は、[n] の領域の <n> 個の全文を読み取ります (最大 2000 トークン)
oc next すでに開いているページの次の予算相当
oc raw [url] ページ全体のマークダウンを抽出
oc fill <n> <text> を番号付き入力に入力します (v0.2)
oc submit [n] フォームを送信する (v0.2)
フラグ: --budget <トークン> (デフォルトは 500)、 --json 、 --html (クリーンな HTML として生)、 --session <name> 、 --verbose / -v (標準エラー出力のメトリクス、またはエクスポート OC_VERBOSE=1 )。
oc open は ~/.only-cli ( OC_HOME でオーバーライド) の下のセッションごとに JSON ファイルにレンダリングしたページを記憶するため、エージェントが URL を処理することなく oc do 3 は [3] に従います。予算を超えるページには、省略した内容が記載されています。 oc find 、 oc read <n> 、および oc next はページを再フェッチせずに残りを読み取ります。予算はハードキャップではなく目標です。少し長くしか実行されないページは、カットされるのではなく丸ごと印刷されます。これは、追加のツール呼び出し 1 回のコストが、節約できるトークンよりもはるかに高くなるためです。
ニュース サイト、ブログ、ドキュメント、フォーラム、検索エンジンなど、サイトごとのセットアップを必要としないほとんど静的なサイトで動作します。それに加えて、clis/ には以下の調整されたショートカットが同梱されています。
これらのいくつか (X、Stack Overflow、YouTube) は、サーバーによってレンダリングされた HTML、フィード、またはログインなしでページにすでに配信されているインライン データを見つけることによって、外部からはログインゲートまたは JS のみに見えるページを読み取ります。まだサポートされていません: JavaScript でのみレンダリングするページ、ログインの背後にあるサイト、フィードを公開しないハード ボット チャレンジのあるサイト。
そのリストにウェブサイトを載せたいですか?プル リクエストを開くか、サイトに名前を付ける問題を開きます。 CONTRIBUTING.md を参照してください。
完全な方法論、タスクごとの番号、およびその他のエージェント/モデルは、only-cli/benchmarks に存在します。ニュースのトップページ、Reddit のディスカッション、検索結果ページなどのライブ サイトに対して測定された短いバージョン:
早いです。 v0

.1 では、静的ページ、予算を意識したレンダリング、オフライン テストがカバーされています。セッション、 oc do <n> 、 oc find <query> 、 oc read <n> 、および oc next が含まれ、残りのアクション ( fill 、 submit 、 back ) は v0.2 に導入され、スクリプトの多いページ用の遅延ヘッドレス フォールバックは v0.3 に導入されます。
正直なところ、既知の制限: JavaScript レンダリングはまだなく、ログインの背後にあるサイトはまだなく、ハード ボット チャレンジの背後にあるページは引き続きツールを拒否する可能性があります。
Only-cli 、作成者およびメンテナ
あらゆる Web サイトを AI エージェント向けにカスタマイズされたコンパクトな CLI に変えます。数万ではなく、数百のトークンで Web を閲覧します。
1 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Turn any website into a compact CLI tailored for AI agents. Browse the web in hundreds of tokens, not tens of thousands. - only-cli/oc

GitHub - only-cli/oc: Turn any website into a compact CLI tailored for AI agents. Browse the web in hundreds of tokens, not tens of thousands. · GitHub
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
only-cli
/
oc
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
52 Commits 52 Commits Folders and files
.github .github clis clis docs docs skills/ only-cli skills/ only-cli src src tests tests .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md README.md README.md llms.txt llms.txt package-lock.json package-lock.json package.json package.json View all files Repository files navigation
Turns websites into a command line interface for AI agents. oc open <url> fetches a page and hands back a compact, numbered view instead of raw HTML or a screenshot, so agents like Claude Code, Codex, and Antigravity can browse without burning tokens. It also gets past blocks that stop naive fetchers on some sites, by talking to the page the way a real browser would.
$ oc open news.ycombinator.com
# Hacker News
[1] Show HN: I built a tiny CSV toolkit
[2] 312 comments
...
actions: do <n> | read <n> | next | raw
$ oc do 1
A typical page is tens of thousands of tokens of markup; the view above fits in a few hundred. No per-site adapters required, no browser extension, no daemon.
If you are an LLM reading this repository, llms.txt is the short version.
npm install -g @only-cli/oc
Requires Node 20+. Requests impersonate Chrome via impers ; falls back to native fetch if impers is unavailable.
Add one line to your agent's instructions file (CLAUDE.md, AGENTS.md, or equivalent):
When you need content from a web page, run npx @only-cli/oc open <url> instead of fetching raw HTML. Run npx @only-cli/oc --help once to learn the commands.
Claude Code users can install the skill instead: copy skills/only-cli/ into .claude/skills/ , or run npx skills add only-cli/oc , which also works in Cursor, Codex, Copilot, and others via skills.sh .
No setup at all also works: npx @only-cli/oc runs without a global install, and teaches its own commands through --help and the actions: line on every render.
oc open <url> fetch and render a page with numbered actions
oc do <n> follow the numbered link [n], or read [n] if it is text
oc find <query> where a string appears on the page already open
oc read <n> full text of the region at [n], up to 2000 tokens
oc next the next budget worth of the page already open
oc raw [url] distilled markdown of the whole page
oc fill <n> <text> type into a numbered input (v0.2)
oc submit [n] submit a form (v0.2)
Flags: --budget <tokens> (default 500), --json , --html (raw as cleaned HTML), --session <name> , --verbose / -v (metrics on stderr, or export OC_VERBOSE=1 ).
oc open remembers the page it rendered in a JSON file per session under ~/.only-cli (override with OC_HOME ), so oc do 3 follows [3] without the agent ever handling a URL. Pages longer than the budget say what they left out; oc find , oc read <n> , and oc next read the rest without refetching the page. The budget is a target rather than a hard cap: a page that would only run a little long is printed whole rather than cut, since one extra tool call costs far more than the tokens it would have saved.
Works on any mostly-static site with no per-site setup: news sites, blogs, documentation, forums, search engines. On top of that, clis/ ships tuned shortcuts for:
A few of these (X, Stack Overflow, YouTube) read pages that look login-gated or JS-only from the outside, by finding the server-rendered HTML, feed, or inline data the page already ships without a login. Not supported yet: pages that only render with JavaScript, sites behind logins, and sites with hard bot challenges that expose no feed.
Want a website on that list? Open a pull request, or an issue naming the site — see CONTRIBUTING.md .
Full methodology, per-task numbers, and other agents/models live in only-cli/benchmarks . The short version, measured against live sites across a news front page, a Reddit discussion, a search results page, and more:
Early. v0.1 covers static pages, budget-aware rendering, and offline tests. Sessions, oc do <n> , oc find <query> , oc read <n> , and oc next are in, the rest of the actions ( fill , submit , back ) land in v0.2, and a lazy headless fallback for script-heavy pages in v0.3.
Known limits, honestly: no JavaScript rendering yet, no sites behind logins yet, and pages behind hard bot challenges may still refuse the tool.
only-cli , creator and maintainer
Turn any website into a compact CLI tailored for AI agents. Browse the web in hundreds of tokens, not tens of thousands.
1 fork Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
