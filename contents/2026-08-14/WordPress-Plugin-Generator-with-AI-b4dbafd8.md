---
source: "https://steem.dev"
hn_url: "https://news.ycombinator.com/item?id=49296681"
title: "WordPress Plugin Generator with AI"
article_title: "Steem: AI WordPress Plugin Generator | Build Plugins From a Prompt"
author: "startuphubai"
captured_at: "2026-08-14T10:55:44Z"
capture_tool: "hn-digest"
hn_id: 49296681
score: 1
comments: 1
posted_at: "2026-08-14T10:03:26Z"
tags:
  - hacker-news
  - translated
---

# WordPress Plugin Generator with AI

- HN: [49296681](https://news.ycombinator.com/item?id=49296681)
- Source: [steem.dev](https://steem.dev)
- Score: 1
- Comments: 1
- Posted: 2026-08-14T10:03:26Z

## Translation

タイトル: AI を使用した WordPress プラグイン ジェネレーター
記事のタイトル: Steem: AI WordPress プラグイン ジェネレーター |プロンプトからプラグインを構築する
説明: AI を使用して、テキスト プロンプトから本番環境に対応した WordPress プラグインを生成します。クリーンな PHP、ライブ プレビュー、ローカル セットアップなしで、カスタム WooCommerce 拡張機能、管理ダッシュボード、サイト アーキテクチャを数分で構築できます。

記事本文:
Steem: AI WordPress プラグイン ジェネレーター |プロンプトからプラグインを構築する Steem の機能 仕組み ギャラリー ガイド 料金 FAQ 構築を開始する スニペットではありません。インストールするプラグイン。 ZIP で配布される WordPress プラグイン ジェネレーター。
必要なものを説明するか、すでに使用しているプラ​​グインを示します。 Steem はクリーンで安全な PHP を作成し、どこにでもインストールする前に実行をプレビューします。
75 の無料クレジット。1 つの完全なプラグインを構築および改良するのに十分です。カードはなく、ダウンロードする前にブラウザの実際の WordPress で実行できます。
生成された WordPress コードで懸念されるのは、出力がエスケープされておらず、ノンスが欠落していることです。実際に戻ってきたものの形がこちらです。
public function save_settings() { if ( ! current_user_can( 'manage_options' ) ) { wp_die( esc_html__( '許可されません。', 'my-plugin' ) ); check_admin_referer( 'my_plugin_save', 'my_plugin_nonce' ); update_option( 'my_plugin_email', sanitize_email( wp_unslash( $_POST['email'] ?? '' ) ) );グローバル $wpdb; $rows = $wpdb->get_results( $wpdb->prepare( "SELECT id, name FROM {$wpdb->prefix}my_plugin WHERE status = %s", 'active' ) ); echo esc_html( $rows[0]->name );デフォルトで安全
機能チェック、nonce、サニタイズされた入力、およびエスケープされた出力は最初から書き込まれ、レビューで不足していることが判明した後に追加されるものではありません。
普通のPHP。ライセンスキーやランタイム依存関係はなく、Steem へのコールバックもありません。修正し、出荷し、販売します。
プラグイン ヘッダー、設定 API、適切なファイル構造、翻訳可能な文字列。コンパイルするファイルではなく、プラグインのように動作します。
WordPress は Web で大きなシェアを占めています。それに関連するツールの開発が追いついていません。
WordPress 専用に構築された PHP を対象とした一般的なコード ジェネレーターではありません。 Steem はフック、フィルター、設定 API、nonce、プラグインのディレクトリ レイアウトを認識しているため、記述されたものは次のように動作します。

たまたまコンパイルされるファイルではなく、プラグインです。本番環境に対応したコード 実際のファイル構造を備えたクリーンで読みやすい PHP: プレースホルダー、中途半端なスタブ、ロジックがあるべき場所に「これを自分で実装してください」というコメントはありません。デフォルトのセキュリティ 機能チェック、nonce 検証、サニタイズされた入力、エスケープされた出力は最初から組み込まれており、レビューで不足していることが判明した後に追加されるものではありません。インストールする前にプレビューする すべてのプラグインは、ブラウザ上の実際の WordPress インスタンスで起動します。サイトに何かが触れる前に、管理画面とフロントエンドをクリックしてください。 SEO と AI の可視性 スキーマ マークアップとクリーンなセマンティック出力が標準装備されているため、プラグインがレンダリングするページは、検索クローラーやアンサー エンジンでも同様に判読できます。数分で発送します。説明し、構築を観察し、ZIP をダウンロードします。ローカル環境、構築ステップ、手作業で仕上げる足場は必要ありません。仕組み
プラグインが何をすべきかを平易な英語で書きます。仕様も定型文もありません。
Steem はファイルを書き込み、その進行状況をそれぞれ表示します。
実際の WordPress で実行されているものをクリックして、ZIP を取得します。
Steem は WordPress 用のプラグインを構築するため、プラグインをインストールする場所が必要になります。何もないところから始めるのであれば、これが近道です。
ウェブの 40% 以上が WordPress で実行されており、これがデフォルトになっているのには理由があります。
サイトはあなたのものです。あなたのドメイン、あなたのデータベース、あなたのコンテンツです。他人のプラットフォームのサブドメインではない
ページは HTML として届くため、検索エンジンと回答エンジンは最初に JavaScript を実行せずにページを読み取ります。
MCP サーバーは WordPress 用に存在するため、Claude または ChatGPT をサイトに直接公開できます
または SiteGround SiteGround の特典は、このリンク経由で提供されます プレミアム ホスティングが最大 80% オフ 新しいプロジェクト用の無料ドメイン 別のホスト、WordPress.com、EuroDNS、または Namecheap からの無料移行

AI サイト ビルダーを使用して WordPress ホスティングを管理します。任意のプラグインがインストールされます。 SiteGround : Google Cloud 上の自己ホスト型 WordPress、管理された更新機能。任意のプラグインがインストールされます。 WordPress.com : WordPress 自体を開発している会社が提供する、フルマネージドのオプションです。公式 MCP サーバーを出荷するので、Claude または ChatGPT をサイトに直接公開できます。プラグインには月額 4 ドルからの有料プランが必要です。無料利用枠ではインストールできません。 EuroDNS : ヨーロッパの 1 つのプロバイダーから管理された WordPress とドメイン。ワンクリックでインストールできるため、どのプラグインも機能します。 Namecheap : ドメインと並行して、EasyWP を通じて 1 つのプロバイダーから WordPress を管理します。プラグインはすべてのプランにインストールされます。 EasyWP は独自のキャッシュを行うため、キャッシュ プラグインはブロックされます。
75 クレジット、無料。決定する前に、実際のものを構築するのに十分です。
© 2026 Steem. WordPress プラグイン、生成されました。

## Original Extract

Generate production-ready WordPress plugins from text prompts using AI. Build custom WooCommerce extensions, admin dashboards, and site architectures in minutes, with clean PHP, a live preview, and no local setup.

Steem: AI WordPress Plugin Generator | Build Plugins From a Prompt Steem Features How it works Gallery Guides Pricing FAQ Start building Not a snippet. A plugin you install. The WordPress plugin generator that ships a ZIP.
Describe what you need, or point at a plugin you already use. Steem writes clean, secure PHP and previews it running before you install it anywhere.
75 free credits, enough to build and refine one complete plugin. No card, and you can run it in a real WordPress in your browser before downloading.
The fear with generated WordPress code is unescaped output and missing nonces. Here is the shape of what actually comes back.
public function save_settings() { if ( ! current_user_can( 'manage_options' ) ) { wp_die( esc_html__( 'Not allowed.', 'my-plugin' ) ); } check_admin_referer( 'my_plugin_save', 'my_plugin_nonce' ); update_option( 'my_plugin_email', sanitize_email( wp_unslash( $_POST['email'] ?? '' ) ) ); global $wpdb; $rows = $wpdb->get_results( $wpdb->prepare( "SELECT id, name FROM {$wpdb->prefix}my_plugin WHERE status = %s", 'active' ) ); echo esc_html( $rows[0]->name ); } Secure by default
Capability checks, nonces, sanitised input and escaped output are written from the start, not added after a review finds them missing.
Ordinary PHP. No licence key, no runtime dependency, and nothing calling back to Steem. Modify it, ship it, sell it.
Plugin header, Settings API, proper file structure and translation-ready strings. It behaves like a plugin, not a file that compiles.
WordPress runs a large share of the web. The tooling around it has not kept pace.
Purpose-built for WordPress Not a general code generator pointed at PHP. Steem knows hooks, filters, the settings API, nonces, and the plugin directory layout, so what it writes behaves like a plugin rather than a file that happens to compile. Production-ready code Clean, readable PHP with a real file structure: no placeholders, no half-finished stubs, no "implement this yourself" comments where the logic should be. Security by default Capability checks, nonce verification, sanitised input and escaped output are written in from the start, not bolted on after a review finds them missing. Preview before you install Every plugin boots in a real WordPress instance in your browser. Click through the admin screens and the front end before anything touches your site. SEO and AI visibility Schema markup and clean semantic output come as standard, so the pages your plugin renders are legible to search crawlers and answer engines alike. Ship in minutes Describe it, watch it build, download the ZIP. No local environment, no build step, no scaffolding you have to finish by hand. How it works
Write what the plugin should do, in plain English. No spec, no boilerplate.
Steem writes the files and shows you each one as it goes.
Click through it running in real WordPress, then take the ZIP.
Steem builds plugins for WordPress, so you will need somewhere to install one. If you are starting from nothing, this is the short way there.
Over 40% of the web runs on WordPress, and it is the default for a reason
The site is yours: your domain, your database, your content. Not a subdomain on someone else’s platform
Pages arrive as HTML, so search engines and answer engines read them without running any JavaScript first
MCP servers exist for WordPress, so Claude or ChatGPT can publish to your site directly
or SiteGround SiteGround offers, via this link Up to 80% off premium hosting Free domain for a new project Free migration from another host or WordPress.com or EuroDNS or Namecheap Managed WordPress hosting with an AI site builder. Any plugin installs. SiteGround : Self-hosted WordPress on Google Cloud, with managed updates. Any plugin installs. WordPress.com : The fully managed option, from the company behind WordPress itself. Ships an official MCP server, so Claude or ChatGPT can publish straight to the site. Plugins need a paid plan, from $4 a month; the free tier cannot install them. EuroDNS : Managed WordPress and the domain from one European provider. One-click install, so any plugin works. Namecheap : Managed WordPress through EasyWP, alongside the domain, from one provider. Plugins install on every plan. Caching plugins are blocked, because EasyWP does its own caching.
75 credits, free. Enough to build something real before deciding.
© 2026 Steem. WordPress plugins, generated.
