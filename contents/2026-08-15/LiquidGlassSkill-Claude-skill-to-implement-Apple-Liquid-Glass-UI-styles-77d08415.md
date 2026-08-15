---
source: "https://github.com/stormaref/LiquidGlassSkill"
hn_url: "https://news.ycombinator.com/item?id=49309425"
title: "LiquidGlassSkill: Claude skill to implement Apple Liquid Glass UI styles"
article_title: "GitHub - stormaref/LiquidGlassSkill: Claude skill to implement Apple Liquid Glass UI style · GitHub"
author: "avestura"
captured_at: "2026-08-15T11:10:50Z"
capture_tool: "hn-digest"
hn_id: 49309425
score: 2
comments: 0
posted_at: "2026-08-15T10:31:58Z"
tags:
  - hacker-news
  - translated
---

# LiquidGlassSkill: Claude skill to implement Apple Liquid Glass UI styles

- HN: [49309425](https://news.ycombinator.com/item?id=49309425)
- Source: [github.com](https://github.com/stormaref/LiquidGlassSkill)
- Score: 2
- Comments: 0
- Posted: 2026-08-15T10:31:58Z

## Translation

タイトル: LiquidGlassSkill: Apple 液体ガラス UI スタイルを実装するためのクロード スキル
記事タイトル: GitHub - stormaref/LiquidGlassSkill: Apple リキッド グラス UI スタイルを実装するためのクロード スキル · GitHub
説明: Apple Liquid Glass UI スタイルを実装するクロード スキル - stormaref/LiquidGlassSkill

記事本文:
GitHub - stormaref/LiquidGlassSkill: Apple リキッド グラス UI スタイルを実装するためのクロード スキル · GitHub
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
ストーマレフ
/
液体ガラススキル
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
8 コミット 8 コミット .claude-plugin .claude-plugin docs docs plugins/liquid-glass plugins/liquid-glass LICENSE LICENSE README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
液体ガラス — Cl

オードコードスキル
Apple スタイルのつや消しの屈折サーフェス: 色合いと縁をペイントする CSS レイヤー、および本物のガラスのようにパネルの端を通して背景を曲げる SVG ディスプレイスメント フィルター。
ほとんどの「glassmorphism」レシピは、backdrop-filter:blur() と、ぼやけた長方形として読み取られる白いオーバーレイです。このスキルには、残りの半分 (屈折) と、ガラスがガラスのように見えるか灰色の箱のように見えるかを決定するいくつかのルールが含まれます。ここでのすべてのルールは、正しく見える前に間違って見えた UI に由来しています。
どちらのショットも、assets/demo.html をそのままレンダリングしたものです。CSS レイヤーのみで、屈折はありません。これは、Chromium 以外のブラウザーが取得するフォールバックでもあります。パネルの色は、周囲の背景が透けて見える色です。ティント自体は無色です。
このリポジトリは、Claude Code プラグイン マーケットプレイスです。クロードコードでは:
/プラグイン マーケットプレイス追加 stormaref/LiquidGlassSkill
/プラグインのインストールLiquid-glass@stormaref-skills
そして、自分の言葉でグラスを求めてください。スキルはモデルによって呼び出されます。「このカードをすりガラスにする」、「サイドバーにガラスモーフィズムを追加する」、「ガラスパネルが平らに見える」はすべてそれに到達します。手動で実行するには、プラグイン - /liquid-glass:liquid-glass の下に名前空間が設定されます。
プラグイン機構のない単純なスキル ディレクトリを好みますか?スキルフォルダーをコピーします。
git clone -- Depth 1 https://github.com/stormaref/LiquidGlassSkill.git /tmp/lgs
cp -r /tmp/lgs/plugins/liquid-glass/skills/liquid-glass ~ /.claude/skills/
何が入っているのか
スキルは plugins/liquid-glass/skills/liquid-glass/ にあります。以下のパスはそれに対する相対パスです。
assets/liquid-glass.css はプレーンなスタイルシートです。リンクして、ページシェルにliquid-glass-backdrop、パネルにliquid-glassを配置します。 Angular プロジェクトはディレクティブをそのままインポートできます。それ以外のすべて:references/refraction.md は完全な移植仕様です。または

クロードセッションを向けて尋ねてください。
始める前に知っておくべき 2 つのこと:
背景は必須です。ガラスはその背後にあるものを示します。平らな塗りつぶしの上には何も表示されません。最初にアンビエント メッシュを構築し、ページにガラスがない状態で不均一に見えることを確認します。
屈折は Chromium のみであり、機能検出できません (Safari は backdrop-filter: url(#…) を解析し、何も描画しません)。他の場所では、スタイルシートのぼかしフォールバックが外観を伝えます。ディレクティブはエンジンにゲートをかけ、不活性なままになります。
屈折フィールドは、Armagan Amcalar (MIT) による Liquid-glass-js のフラグメント シェーダーのポートです。このポートは、ライブラリの html2canvas ページ スナップショットではなく、ライブ背景によって供給される SVG ディスプレイスメント フィルターをターゲットとしているため、スクロール、テーマの変更、コンテンツの更新を無料で追跡します。アップストリームの通知は、assets/liquid-glass.directive.ts の先頭にあります。コピーまたはポート上に保持してください。
これが主な表面処理である実稼働の Angular アプリから抽出されたものです。カード、トップバー、ダイアログ、メニュー、タスク ボードはすべてその上にあります。
Apple リキッド グラス UI スタイルを実装するためのクロード スキル
Readme MIT ライセンス アクティビティ スター
フォーク数 0 レポート リポジトリの寄稿者
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Claude skill to implement Apple Liquid Glass UI style - stormaref/LiquidGlassSkill

GitHub - stormaref/LiquidGlassSkill: Claude skill to implement Apple Liquid Glass UI style · GitHub
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
stormaref
/
LiquidGlassSkill
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
8 Commits 8 Commits .claude-plugin .claude-plugin docs docs plugins/ liquid-glass plugins/ liquid-glass LICENSE LICENSE README.md README.md View all files Repository files navigation
Liquid Glass — a Claude Code skill
Apple-style frosted, refracting surfaces: a CSS layer that paints the tint and rim, and an SVG displacement filter that bends the backdrop through the panel's edge the way real glass does.
Most "glassmorphism" recipes are backdrop-filter: blur() plus a white overlay, which reads as a blurred rectangle. This skill carries the other half — the refraction — plus the handful of rules that decide whether glass looks like glass or like a gray box. Every rule here came from a UI that looked wrong before it looked right.
Both shots are assets/demo.html rendered as-is — the CSS layer alone, no refraction, which is also the fallback every non-Chromium browser gets. The colour in the panels is the ambient backdrop showing through; the tint itself is colorless.
This repo is a Claude Code plugin marketplace. In Claude Code:
/plugin marketplace add stormaref/LiquidGlassSkill
/plugin install liquid-glass@stormaref-skills
Then ask for glass in your own words. The skill is model-invoked: "make this card frosted glass", "add glassmorphism to the sidebar", "the glass panel looks flat" all reach it. To run it by hand it is namespaced under the plugin — /liquid-glass:liquid-glass .
Prefer a plain skill directory with no plugin machinery? Copy the skill folder out:
git clone --depth 1 https://github.com/stormaref/LiquidGlassSkill.git /tmp/lgs
cp -r /tmp/lgs/plugins/liquid-glass/skills/liquid-glass ~ /.claude/skills/
What's in it
The skill lives at plugins/liquid-glass/skills/liquid-glass/ ; paths below are relative to it.
assets/liquid-glass.css is a plain stylesheet: link it, put liquid-glass-backdrop on the page shell, liquid-glass on panels. Angular projects can import the directive as-is. Everything else: references/refraction.md is a complete porting spec — or point a Claude session at it and ask.
Two things to know before you start:
The backdrop is required. Glass shows what is behind it; over a flat fill there is nothing to show. Build the ambient mesh first and check the page looks uneven with no glass on it.
Refraction is Chromium-only , and cannot be feature-detected (Safari parses backdrop-filter: url(#…) and paints nothing). Everywhere else the stylesheet's blur fallback carries the look. The directive gates on engine and stays inert.
The refraction field is a port of the fragment shader in liquid-glass-js by Armagan Amcalar (MIT). This port targets an SVG displacement filter fed by the live backdrop rather than the library's html2canvas page snapshot, so it tracks scrolling, theme changes and content updates for free. The upstream notice sits at the top of assets/liquid-glass.directive.ts — keep it there on any copy or port.
Extracted from a production Angular app where this is the primary surface treatment: cards, topbars, dialogs, menus and the task board all sit on it.
Claude skill to implement Apple Liquid Glass UI style
Readme MIT license Activity Stars
0 forks Report repository Contributors
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
