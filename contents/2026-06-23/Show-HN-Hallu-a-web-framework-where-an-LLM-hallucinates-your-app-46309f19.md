---
source: "https://github.com/alehlopeh/hallu"
hn_url: "https://news.ycombinator.com/item?id=48648823"
title: "Show HN: Hallu – a web framework where an LLM hallucinates your app"
article_title: "GitHub - alehlopeh/hallu: This web app does not exist · GitHub"
author: "alehlopeh"
captured_at: "2026-06-23T18:39:32Z"
capture_tool: "hn-digest"
hn_id: 48648823
score: 1
comments: 0
posted_at: "2026-06-23T18:03:57Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Hallu – a web framework where an LLM hallucinates your app

- HN: [48648823](https://news.ycombinator.com/item?id=48648823)
- Source: [github.com](https://github.com/alehlopeh/hallu)
- Score: 1
- Comments: 0
- Posted: 2026-06-23T18:03:57Z

## Translation

タイトル: Show HN: Hallu – LLM がアプリを幻覚させる Web フレームワーク
記事タイトル: GitHub - alehlopeh/hallu: この Web アプリは存在しません · GitHub
説明: この Web アプリは存在しません。 GitHub でアカウントを作成して、alehlopeh/hallu の開発に貢献してください。

記事本文:
GitHub - alehlopeh/hallu: この Web アプリは存在しません · GitHub
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
アレロペ
/
ハル
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード もっと開く acti

オンメニュー フォルダーとファイル
7 コミット 7 コミット デモ デモの例 例 スクリプト スクリプト src src test test .gitignore .gitignore ライセンス ライセンス README.md README.md bun.lock bun.lock package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイル ナビゲーション
Hallu: この Web アプリは存在しません
Hallu は、LLM がアプリ全体を幻覚させる Web フレームワークです。
Hallu を使用すると、思いもよらなかった機能がすでに提供されています。これは、リクエストを受け取り、実際のデータベースに対して Web アプリのように動作するエージェント ハーネスです。これは、Claude Code の「コンパイラー」に対する「インタープリター」です。アプリケーション ロジックは、ビルド時ではなく実行時に生成されます。非公式の散文仕様に対してリクエストごとの (証明のない) プログラム合成を考えてみましょう。
すべてのリクエストは、SQL ツールと HTML を返す命令を使用してモデルにルーティングされます。モデルは、コントローラー、ORM、およびテンプレート エンジンです。
例の 1 つを Haiku にポイントし、あちこちをクリックして、動作するアプリが一度に 1 つの機能を自動的に組み立てるのを確認します。
クロードだけにアプリを書かせてはいけません。クロードをあなたのアプリにしてください。
ブン+ホノで運営。デフォルトで SQLite を使用し、Postgres をサポートします。ご自身のモデルをご持参ください。
GET /* はモデルにページ本文を要求します。クライアント ランタイムはフォームの送信をインターセプトし、 /__hallu/action に POST します。
モデルは、ループ内の単一の SQL ツールを通じてデータベースの読み取りと書き込みを行います。
アクションは <hallu-update target="id">...</hallu-update> ブロックをストリームバックします。クライアントは、ID によって各リージョンを交換します。まるでSPAのような雰囲気です。
streamResponses を使用すると、アクションはストリーム ツールを通じてコン​​テナーにストリーミングされます。フレームワークはラッパー マークアップを追加し、トークンが到着するとそれを埋め、最後に Hallu:finalize イベントを起動します。 2 つのモード: デフォルトでリテラル テキスト (マークアップ エスケープ)、または html: true でトークンが到着するとストリームをライブ HTML としてレンダリングします。

pageChat を使用すると、フローティング パネルを使用して、ユーザーは指示によって現在のページを修正できます。フレームワークはページと命令を /__hallu/revise にポストし、モデルはそれを所定の位置に書き換えます。編集内容はパス グロブによって保存され、その後一致するページをレンダリングするたびに再適用されるため、固定されます。
レンダリングされたページはパスごとにキャッシュされるため、ウォーム ロードではモデルがスキップされます。 DB 書き込みでは、影響を受けるページ (デフォルトでは粗いページ、またはグロブでスコープ指定されたページ) が無効になります。
固定 (テーブルがスキーマ全体) または autoSchema (モデルがまだ見ていないパスに対してその場でテーブルを作成します)
スキーマは参照するにつれて増加します。初めてパスにアクセスすると、モデルがテーブルを設計して作成し、ページを HTML にレンダリングします。 /new をパスに追加すると、モデルがフォームを生成します。フォームを送信すると、モデルがレコードを挿入し、ページをレンダリングします。 /delete をパスに追加すると、モデルは「よろしいですか?」というメッセージを表示します。フォームを作成し、要求に応じてレコードを削除します。
スキーマを自分で固定します。構成内でテーブルを宣言すると、フレームワークが起動時にテーブルを作成します。モデルはその内部でのみ読み取りと書き込みを行い、DDL を実行することはありません。そのため、データの形式は固定されており、既知です。
「現実世界」の百科事典。すべての記事はモデルのナレッジ グラフからオンデマンドで生成され、同じモデルが SQL と HTML を生成します。
Postgres、ハイブリッド autoSchema、宣言されたテーブルに関するプログラマの Q&A サイト。
streamResponses 上の ChatGPT スタイルのチャット アプリ。モデルは HTML トークンをトークンごとにスレッドにストリーミングします。
スキーマレスの CRM。オブジェクトにアクセスすると、モデルがそのテーブルを設計および作成し、シードします。
製品、価格、レビューなどのカタログ全体がオンデマンドで作成され、戻ってきたときにページが安定するように保存されるストア。
bunx Hallujs は myapp を生成します # デフォルト: Anthropic + SQLite
cd myapp && bun install
エコー「ANTHROPIC_AP」

I_KEY=sk-ant-... " > .env
パン開発者
generate は、バックエンドとプロバイダーのフラグ --postgres および --anthropic (デフォルト) / --openai / --ollama を受け取ります。
Hallu アプリでのバイブは、電話ゲームによるバイブコーディングのようなものです。あなたのバイブの健全性をチェックするには、パン品質の Hallu.config.ts を使用してください。
Hallu.config.tsのデフォルトエクスポートdefineConfig({...}) 。必須: 名前、説明、モデル。
テナントは { アカウント、コンテキスト? }
SqlEvent は { クエリ、ok、変更されました、エラー? }
SQLite パスのデフォルトは ./hallu.db ( dir 、 ./data の下のアカウントごとのファイル)
Postgres スキーマのデフォルトは public です。
遅い: 1 ページはモデル呼び出しであり、場合によっては複数の呼び出しです。 Haiku ではページの読み込みに約 2 秒かかります。ただし、これは LLM を呼び出すツールを利用した Web サイトです。何を期待していましたか？
トークンがかかります。すべてのコールド リクエストがモデルにヒットします。キャッシュがあるので非常識ではありませんし、Haiku は比較的安価で完璧に仕事をしてくれます。
これは非決定的です。同じ URL でも、コールド ロードとキャッシュの無効化のたびにわずかに異なるレンダリングが行われます。
セキュリティ モデルは完璧ではありません。フレームワークは信頼できない入力を LLM に渡し、任意の SQL を書くように要求します。それは危険です。リポジトリには、一般的な SQL インジェクション手法をテストするスクリプトがあります。しかし、そうです、これは LLM 幻覚 SQL です。これを重要なことには使用しないでください。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
貢献者
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

This web app does not exist. Contribute to alehlopeh/hallu development by creating an account on GitHub.

GitHub - alehlopeh/hallu: This web app does not exist · GitHub
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
alehlopeh
/
hallu
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
7 Commits 7 Commits demo demo examples examples scripts scripts src src test test .gitignore .gitignore LICENSE LICENSE README.md README.md bun.lock bun.lock package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
Hallu: this web app does not exist
Hallu is a web framework where an LLM hallucinates your entire app.
With Hallu, you've already shipped features you never even thought of. It's an agent harness that takes a request and acts like a web app against a real database. It's the "interpreter" to Claude Code's "compiler": Application logic is generated at runtime instead of at build time. Think per-request (proofless) program synthesis against an informal prose spec.
Every request routes to a model with a SQL tool and instructions to return html. The model is the controller, the ORM, and the template engine.
Point one of the examples to Haiku and click around to watch a working app assemble itself one feature at a time.
Don't just have Claude write your app. Have Claude be your app.
Runs on Bun + Hono. Uses SQLite by default, supports Postgres. Bring your own model.
GET /* asks the model for the page body. A client runtime intercepts form submits and POSTs them to /__hallu/action .
The model reads and writes the database through a single sql tool in a loop.
Actions stream back <hallu-update target="id">...</hallu-update> blocks. The client swaps each region in by id. It feels like an SPA.
With streamResponses , an action streams into a container through a stream tool. The framework appends your wrapper markup, fills it as tokens arrive, and fires a hallu:finalize event at the end. Two modes: literal text by default (markup escaped), or html: true to render the stream as live HTML as the tokens land.
With pageChat , a floating panel lets the user revise the current page by instruction. The framework posts the page and instruction to /__hallu/revise and the model rewrites it in place. The edit is saved by path glob and re-applied on every later render of a matching page, so it sticks.
Rendered pages are cached per path so warm loads skip the model. A DB write invalidates affected pages (coarse by default, or scoped with a glob).
Fixed ( tables is the whole schema) or autoSchema (the model creates tables on the fly for paths it hasn't seen)
The schema grows as you browse. Visit a path for the first time and the model will design a table, create it, and render the page to html. Add /new to the path, and the model will generate a form. Submit the form and the model will insert the record and render the page. Add /delete to the path, and the model will present you with an "Are you sure?" form and then delete the record, as requested.
Pin the schema yourself. Declare your tables in the config and the framework creates them on boot. The model reads and writes only within them and never runs DDL, so the shape of your data is fixed and known.
A "real world" encyclopedia. Every article is generated on-demand from the model's knowledge graph, and the same model generates the SQL and html.
A programmer Q&A site on Postgres, hybrid autoSchema and declared tables.
A ChatGPT-style chat app on streamResponses . The model streams html token-by-token into the thread.
A schema-less CRM. Visit an object and the model designs and creates its table and seeds it.
A store where the whole catalog is invented on demand: products, prices, reviews, then saved so the page is stable on return.
bunx hallujs generate myapp # defaults: Anthropic + SQLite
cd myapp && bun install
echo " ANTHROPIC_API_KEY=sk-ant-... " > .env
bun dev
generate takes flags for the backend and provider: --postgres , and --anthropic (default) / --openai / --ollama .
Vibing on a Hallu app is kind of like vibe coding through a game of telephone. Use bun quality hallu.config.ts to sanity check your vibed vibes.
hallu.config.ts default-exports defineConfig({...}) . Required: name , description , model .
Tenant is { account, context? }
SqlEvent is { query, ok, mutated, error? }
SQLite path defaults to ./hallu.db (per-account files under dir , ./data )
Postgres schema defaults to public .
It's slow: A page is a model call, sometimes several. Pages take ~2s to load on Haiku. It's a website powered by a tool calling LLM though. What'd you expect?
It costs tokens: Every cold request hits the model. There's a cache, so it's not insane, and Haiku is relatively cheap and does the job perfectly.
It's non-deterministic: The same URL renders slightly differently on every cold load and cache invalidation.
The security model isn't perfect: The framework passes untrusted input to an LLM and asks it to write arbitrary SQL. That's dangerous. There's a script in the repo that tests common SQL-injection techniques. But yeah, it's an LLM hallucinating SQL. Don't use this for anything important.
There was an error while loading. Please reload this page .
0
forks
Report repository
Contributors
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
