---
source: "https://aliothpress.com/cms-for-ai-agents-webmcp-built-in"
hn_url: "https://news.ycombinator.com/item?id=49384446"
title: "AI CMS with WebMCP tools for agents in admin panel"
article_title: "CMS for AI Agents: WebMCP Built Into AliothPress - AliothPress"
image: "https://aliothpress.com/cover-og.webp"
author: "nikariguel"
captured_at: "2026-08-21T06:28:36Z"
capture_tool: "hn-digest"
hn_id: 49384446
score: 1
comments: 0
posted_at: "2026-08-21T06:20:14Z"
tags:
  - hacker-news
  - translated
---

# AI CMS with WebMCP tools for agents in admin panel

- HN: [49384446](https://news.ycombinator.com/item?id=49384446)
- Source: [aliothpress.com](https://aliothpress.com/cms-for-ai-agents-webmcp-built-in)
- Score: 1
- Comments: 0
- Posted: 2026-08-21T06:20:14Z

## Translation

タイトル: 管理パネルのエージェント向け WebMCP ツールを備えた AI CMS
記事タイトル: AI エージェント用 CMS: AliothPress に組み込まれた WebMCP - AliothPress
説明: AliothPress は、ネイティブ WebMCP サポートを備えたセルフホスト型 CMS です。 AI エージェントはサイトを検索し、フォームに記入して送信し、コンテンツの管理を支援します。

記事本文:
アリオスプレス
ホーム
特長
CMS の機能を詳しく見る
プラグイン
インストールガイド
ドキュメント
コンテンツ管理
多言語システム
ページビルダー
AIアシスタント
メディアライブラリ
ファイルマネージャー
フォームビルダー
ニュースレター
SEO、AEO、パフォーマンス
テーマ
メニュー
ユーザー管理
セキュリティ
バックアップと復元
ファビコンジェネレーター
SSL/HTTPS
ライセンスとアップデート
プライバシーとGDPR
アクセシビリティ
技術スタック
CMSをダウンロード
デモ
JP
DE
ES
アリオスプレス
ホーム
特長
CMS の機能を詳しく見る
プラグイン
インストールガイド
ドキュメント
コンテンツ管理
多言語システム
ページビルダー
AIアシスタント
メディアライブラリ
ファイルマネージャー
フォームビルダー
ニュースレター
SEO、AEO、パフォーマンス
テーマ
メニュー
ユーザー管理
セキュリティ
バックアップと復元
ファビコンジェネレーター
SSL/HTTPS
ライセンスとアップデート
プライバシーとGDPR
アクセシビリティ
技術スタック
CMSをダウンロード
デモ
AliothPress: AI エージェント用 CMS (WebMCP 内蔵)
AliothPress は、WebMCP を介した AI エージェントのネイティブ サポートを備えたセルフホスト型 CMS です。その上に構築されたサイトでは、AI エージェントがコンテンツを検索し、フォームを記述して送信し、管理パネル内で投稿の作成、ページの構築、画像のアップロード、コンテンツの翻訳を行うことができます。この機能はデフォルトではオフになっており、管理パネルの AI アシスタント ページにある 2 つのスイッチで制御されます。 AliothPress が AI エージェントへのアクセスを自動的に有効にすることはありません。サイト所有者は、パブリック ツール、管理ツール、または両方を公開するか、どちらも公開しないかを明示的に決定します。
WebMCP は、robots.txt が検索エンジンにルールを提供するのと同じように、Web サイトが AI エージェントにツールを提供できるようにする新しい Web 標準です。サイトは、エージェントに何が許可されているかを宣言します。エージェントはそれを自動的に検出します。
Discovery : すべての AliothPress サイトは、/.well-known/webmcp で機械可読マニフェストで応答します。
ページ内登録: document.modelContext / navigator.modelContext 経由でツールを登録します。
オプションのポリ

fill : 互換性ポリフィルは、ネイティブ サポートのないブラウザに API を追加します。AI アシスタント ページの 1 つのチェックボックスでオンにします。オリジントライアルトークンは同じページに貼り付けられます
WebMCP サポートは CMS コアに同梱されており、インストールされた状態ですぐに使用でき、プラグインやサードパーティ サービスに依存しません。
AI エージェントがサイト上でできること
パブリック サーフェスを有効にすると、サイトの任意のページにアクセスしたエージェントは次のことが可能になります。
コンテンツを検索: search_site は公開された投稿とページをクエリします
サイトの概要を読む: 自動生成された /llms.txt は、言語ごとにグループ化された AI のサイトを要約しています。
フォームを理解する: description_form はフィールド、タイプ、要件を返します。
訪問者用のフォームを送信します: 予約、連絡先、サインアップ。訪問者が質問し、エージェントがフォームに記入して送信します
例: 訪問者がブラウザで AI エージェントに「土曜日のワークショップを予約してください」と伝えます。エージェントはページ上でフォームを見つけ、入力して送信します。ハニーポットのスパム保護とレート制限は引き続き適用されます。
AI エージェントが管理内でできること
管理画面を有効にすると、認証された管理セッション内で作業するエージェントは最大 24 個のツールを利用できるようになります。
コンテンツ : 投稿とページを生成し、同じ呼び出しでスラッグ、SEO、ソーシャル カード、FAQ、および Schema.org メタデータを設定し、1 つの翻訳グループにリンクされた言語バージョンを作成し、テキストを翻訳し、SEO メタデータを最適化し、既存の投稿とページをフィールドごとに編集します。エージェント名が変更するフィールドのみが変更され、その他はすべて保持されます。
メディア : 画像を一度に 1 つずつ、または 1 回の承認の下でセット全体としてメディア ライブラリにアップロードし (人間によるアップロードと同じパイプライン: 最適化、WebP/AVIF バリアント、EXIF ストリッピング、必須の代替テキスト)、ワン ステップで注目の画像、オープン グラフ、または Twitter 画像として添付します。
ページ構築 : Page Builder ページを作成します。

gly は、バッチで、または変換として、すべてのブロック タイプの機械可読参照を読み取り、ページ ブロックを設定または更新します
一括作業: 1 つの承認ダイアログで数十の投稿またはページを一度に作成できます。ページ ビルダーのレイアウトと画像セットが含まれており、移行やドキュメントの実行に最適です。
サイト管理: メニューの作成、スラッグのチェック、コンテンツのリストと検索、管理セクションの移動
エージェントはサインインしたユーザーとして機能し、同じ権限、同じ検証、同じサニタイズ、同じ監査ログなど、人間が使用する正確なルートを再利用します。別個のエージェント バックドアはコードベースに存在しません。
エージェントが決してできないことの 1 つは、公開です。エージェントはドラフトを作成および編集します。公開リクエストは、エージェントまたはそのツールの主張とは無関係に、サーバー側のドラフトに降格されます。このルールはツール レイヤーとサーバー上で 2 回適用されるため、公開は管理 UI で行われる人間の決定となります。
業務に応じて拡張できる承認
書き込みのたびに、サインインしているユーザーに確認が求められます。 1 つの投稿に対して 1 つのダイアログが表示されます。一括作業の場合、ダイアログは倍増するのではなく拡大縮小されます。
バッチ承認 : エージェントが一度に多くの投稿やページを作成する場合、または一連の画像をアップロードする場合、単一のダイアログにチェックボックス付きのすべての項目が一覧表示されます。すべてを選択するか、厳選したサブセットを承認します。スキップされたアイテムはエージェントに報告されるため、各アイテムに何が起こったかを正確に把握できます。
自動操縦 : どの承認ダイアログでも、そのブラウザ タブでのみ、15 分または 60 分のタイムボックス自動承認を有効にできます。表示されるバッジはカウントダウンされ、ワンクリックで取り消されます。人間だけがスイッチを入れることができます。エージェントツールとしては存在しません。自動操縦下であっても、公開は依然として不可能です
エクスポートせずに別の CMS から移行する
なぜなら、エージェントは 1 つのブラウザ タブで古いサイトを読み、AliothPress に下書きを書き込むことができるからです。

別の方法では、ファイルのエクスポート、プラグインのインポート、または古いサイトのマークアップに沿ったドラッグを行わずにコンテンツを移行できます。エージェントは各ページをきれいに書き換え、画像 (代替テキスト付きで最適化された画像) をアップロードし、注目の画像やソーシャル画像を添付し、ページ ビルダーでデザインされたレイアウトをネイティブに再構築し、すべてをレビュー用の下書きとしてファイルします。古いサイトのインライン スタイル、破損したショートコード、蓄積された間違いは、コードが 1 行もコピーされないため、そのまま残ります。段階的なチュートリアルはドキュメント「Migration to AliothPress」に記載されています。
WebMCP と AI アシスタント: 2 つの異なる機能
AliothPress には、管理パネルに AI アシスタントも組み込まれています。これは別の機能であり、2 つは簡単に区別できます。
AI アシスタント : CMS は独自の API キーを使用して AI プロバイダー (Anthropic Claude、DeepSeek、または Google Gemini) とサーバー間で通信し、コンテンツの生成、翻訳、SEO の最適化を行います。
WebMCP : 訪問者または編集者のブラウザで実行される外部 AI エージェント。サイトが公開するツールを使用します。
これらは連携して動作します。生成、翻訳、SEO 用の WebMCP 管理ツールは、同じ AI アシスタント エンジンと構成したプロバイダーを通じて実行されます。
つまり、AI アシスタントはサイトのために働く AI です。 WebMCP はサイトと連動する AI です。
デフォルトでオフ: 新規インストールではエージェント ツールが公開されません
2 つの独立したスイッチ: パブリック サイトと管理者は個別に有効になります
下書きのみ: エージェントは下書きを保存および編集できますが、公開することはできません。ツールだけでなくサーバー側でも適用されます。
無効の場合は正直です: マニフェストは有効: false を報告します。管理エージェントのエンドポイントが存在しないかのように 404 を返す
訪問者のデータは非公開のままです: フォームの送信はエージェントのエンドポイントから完全に除外されます
すべてのアクションが記録される: エージェント主導の変更アプリ

他の管理アクションと同様に監査ログに記録されます
管理パネルで AI アシスタント ページを開きます
公開ツール (検索 + フォーム) 、管理パネル、またはその両方をオンにします
完了: /.well-known/webmcp のマニフェストがツールのアナウンスを開始します
ウェブサイトを閲覧する代わりに、AI エージェントやチャット インターフェイスを通じて Web サイトにアクセスする人が増えています。エージェントが読んだり、検索したり、行動したりできるサイトは、推奨され、引用され、使用されます。エージェントが解析できないサイトはスキップされます。
AliothPress は、その変化の両方の側面に向けてサイトを準備します。
AI 回答エンジンの場合: /llms.txt 、FAQPage スキーマ、Speakable マークアップ、回答優先のコンテンツ構造
AI エージェント向け: 検索、フォーム、コンテンツ管理用の WebMCP ツール
AliothPress は、お客様が完全に制御できるセルフホスト型 CMS 上で、追加費用なしで、コアでの WebMCP サポートを現在提供しています。 AI ファースト Web 向けに構築されたサイトのプラットフォームを選択している場合、これがそのために作られたものです。 AliothPress を無料でダウンロードし、AI エージェントのガイドまたは CMS ドキュメントをお読みください。
© 2026 アリオスプレス。無断転載を禁じます。

## Original Extract

AliothPress is a self-hosted CMS with native WebMCP support. AI agents can search your site, fill and submit forms, and help you manage content.

AliothPress
Home
Features
Explore CMS features
Plugins
Installation Guide
Documentation
Content Management
Multilingual System
Page Builder
AI Assistant
Media Library
File Manager
Form Builder
Newsletter
SEO, AEO & Performance
Themes
Menus
User Management
Security
Backup & Restore
Favicon Generator
SSL / HTTPS
Licensing & Updates
Privacy & GDPR
Accessibility
Tech Stack
Download CMS
Demo
EN
DE
ES
AliothPress
Home
Features
Explore CMS features
Plugins
Installation Guide
Documentation
Content Management
Multilingual System
Page Builder
AI Assistant
Media Library
File Manager
Form Builder
Newsletter
SEO, AEO & Performance
Themes
Menus
User Management
Security
Backup & Restore
Favicon Generator
SSL / HTTPS
Licensing & Updates
Privacy & GDPR
Accessibility
Tech Stack
Download CMS
Demo
AliothPress: a CMS for AI Agents (WebMCP Built In)
AliothPress is a self-hosted CMS with native support for AI agents through WebMCP. Any site built on it can let AI agents search its content, describe and submit its forms, and, inside the admin panel, create posts, build pages, upload images, and translate content. The feature is off by default and controlled by two switches on the AI Assistant page in the admin panel. AliothPress never enables AI agent access automatically. Site owners explicitly decide whether to expose public tools, admin tools, both, or neither.
WebMCP is an emerging web standard that lets websites offer tools to AI agents, the same way robots.txt offers rules to search engines. A site declares what an agent is allowed to do. The agent discovers it automatically.
Discovery : every AliothPress site answers at /.well-known/webmcp with a machine-readable manifest
In-page registration : tools register via document.modelContext / navigator.modelContext
Optional polyfill : a compatibility polyfill adds the API on browsers without native support, switched on with one checkbox on the AI Assistant page. The origin trial token is pasted on the same page
WebMCP support ships in the CMS core, ready to use as installed, and relies on zero plugins or third-party services.
What AI agents can do on your site
With the public surface enabled, an agent visiting any page of your site can:
Search your content : search_site queries published posts and pages
Read the site overview : auto-generated /llms.txt summarizes the site for AI, grouped by language
Understand your forms : describe_form returns fields, types, and requirements
Submit forms for the visitor : booking, contact, signup. The visitor asks, the agent fills and submits the form
Example: a visitor tells the AI agent in their browser "book me into the Saturday workshop". The agent finds the form on your page, fills it, and submits it. Your honeypot spam protection and rate limits still apply.
What AI agents can do in your admin
With the admin surface enabled, an agent working inside your authenticated admin session gets up to 24 tools:
Content : generate posts and pages, set their slug, SEO, social card, FAQ, and Schema.org metadata in the same call, create language versions linked into one translation group, translate text, optimize SEO metadata, and edit existing posts and pages field by field: only the fields the agent names change, everything else is preserved
Media : upload images into the media library, one at a time or as a whole set under a single approval (the same pipeline as human uploads: optimization, WebP/AVIF variants, EXIF stripping, mandatory alt text), and attach them as the featured, Open Graph, or Twitter image in one step
Page building : create Page Builder pages singly, in batches, or as translations, read a machine-readable reference of every block type, and set or update page blocks
Bulk work : create dozens of posts or pages in one go with a single approval dialog, Page Builder layouts and image sets included, ideal for migrations and documentation runs
Site management : build menus, check slugs, list and search content, navigate admin sections
The agent acts as the signed-in user and reuses the exact routes humans use: same permissions, same validation, same sanitization, same audit log. A separate agent backdoor simply does not exist in the codebase.
One thing an agent can never do is publish. Agents create and edit drafts. A request to publish is demoted to a draft server-side, independently of what the agent or its tools claim. The rule is enforced twice, in the tool layer and on the server, so publishing stays a human decision, made in the admin UI.
Approvals that scale with the job
Every write asks the signed-in user for confirmation. For one post, that is one dialog. For bulk work, the dialogs scale instead of multiplying:
Batch approval : when an agent creates many posts or pages at once, or uploads a set of images, a single dialog lists every item with checkboxes: select all, or approve a hand-picked subset. Skipped items are reported back to the agent, so it knows exactly what happened to each one
Autopilot : any approval dialog can enable time-boxed auto-approval: 15 or 60 minutes, in that browser tab only. A visible badge counts down and revokes with one click. Only a human can switch it on. It does not exist as an agent tool. And even under autopilot, publishing remains impossible
Migrate from another CMS without an export
Because an agent can read your old site in one browser tab and write drafts into AliothPress in another, you can migrate content without export files, import plugins, or dragging along the old site's markup. The agent rewrites each page cleanly, uploads the images (which come out optimized, with alt text), attaches featured and social images, rebuilds designed layouts natively in the Page Builder, and files everything as drafts for your review. The old site's inline styles, broken shortcodes, and accumulated mistakes stay behind, because not a single line of its code is copied. The step-by-step walkthrough is in the documentation: Migration to AliothPress .
WebMCP and the AI Assistant: two different features
AliothPress also ships a built-in AI Assistant in the admin panel. It is a separate feature, and the two are easy to tell apart:
AI Assistant : your CMS talks to an AI provider (Anthropic Claude, DeepSeek, or Google Gemini) server-to-server, using your own API key, to generate content, translate, and optimize SEO for you
WebMCP : external AI agents, running in a visitor's or editor's browser, use tools your site exposes to them
They work together : the WebMCP admin tools for generation, translation, and SEO run through the same AI Assistant engine and the provider you configured
In short: the AI Assistant is AI working for your site. WebMCP is AI working with your site.
Off by default : a fresh install exposes zero agent tools
Two independent switches : public site and admin are enabled separately
Drafts only : agents can save and edit drafts but can never publish, enforced server-side, not just in the tools
Honest when disabled : the manifest reports enabled: false . Admin agent endpoints return 404 as if they do not exist
Visitor data stays private : form submissions are excluded from agent endpoints entirely
Every action logged : agent-driven changes appear in the audit log like any admin action
Open the AI Assistant page in the admin panel
Turn on Public tools (search + forms) , Admin panel , or both
Done: the manifest at /.well-known/webmcp starts announcing your tools
People increasingly reach websites through AI agents and chat interfaces instead of browsing. A site an agent can read, search, and act on gets recommended, cited, and used. A site an agent cannot parse gets skipped.
AliothPress prepares your site for both sides of that shift:
For AI answer engines : /llms.txt , FAQPage schema, Speakable markup, answer-first content structure
For AI agents : WebMCP tools for search, forms, and content management
AliothPress ships WebMCP support today, in the core, at no extra cost, on a self-hosted CMS you fully control. If you are choosing a platform for a site built for the AI-first web, this is what it was made for. Download AliothPress free , read the Guide for AI Agents or CMS documentation .
© 2026 AliothPress. All rights reserved.
