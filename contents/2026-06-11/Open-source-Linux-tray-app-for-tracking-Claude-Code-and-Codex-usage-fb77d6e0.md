---
source: "https://github.com/openusage-community/openusage"
hn_url: "https://news.ycombinator.com/item?id=48487393"
title: "Open-source Linux tray app for tracking Claude Code and Codex usage"
article_title: "GitHub - openusage-community/openusage: Community-driven continuation of OpenUsage. Building the best AI usage tracker for Linux, macOS and Windows. · GitHub"
author: "symonbaikov"
captured_at: "2026-06-11T07:48:49Z"
capture_tool: "hn-digest"
hn_id: 48487393
score: 1
comments: 1
posted_at: "2026-06-11T07:28:40Z"
tags:
  - hacker-news
  - translated
---

# Open-source Linux tray app for tracking Claude Code and Codex usage

- HN: [48487393](https://news.ycombinator.com/item?id=48487393)
- Source: [github.com](https://github.com/openusage-community/openusage)
- Score: 1
- Comments: 1
- Posted: 2026-06-11T07:28:40Z

## Translation

タイトル: クロード コードとコーデックスの使用状況を追跡するためのオープンソース Linux トレイ アプリ
記事のタイトル: GitHub - openusage-community/openusage: コミュニティ主導の OpenUsage の継続。 Linux、macOS、Windows に最適な AI 使用状況トラッカーを構築します。 · GitHub
説明: コミュニティ主導の OpenUsage の継続。 Linux、macOS、Windows に最適な AI 使用状況トラッカーを構築します。 - openusage-コミュニティ/openusage

記事本文:
GitHub - openusage-community/openusage: コミュニティ主導の OpenUsage の継続。 Linux、macOS、Windows に最適な AI 使用状況トラッカーを構築します。 · GitHub
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
オープン使用法コミュニティ
/
使用法を開く
公共
robinebers/openusage からフォークされました
ああ、ああ！
の

読み込み中にエラーが発生しました。このページをリロードしてください。
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
458 コミット 458 コミット .agents/ スキル .agents/ スキル .claude/ スキル .claude/ スキル .codex/ 環境 .codex/ 環境 .cursor .cursor .github .github .local-include .local-include docs docs plugins plugins public public scripts scripts src-tauri src-tauri src src --help --help .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンスREADME.md README.md SECURITY.md SECURITY.md TRADEMARK.md TRADEMARK.md bun.lock bun.lock bunfig.toml bunfig.toml コンポーネント.json コンポーネント.json 導体.json 導体.json copy-bundled.cjs copy-bundled.cjs Index.html Index.html package.json package.json Screenshot.pngスクリーンショット.png tsconfig.json tsconfig.json tsconfig.node.json tsconfig.node.json vite.config.ts vite.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
すべての AI コーディングのサブスクリプションを 1 か所で追跡します。
OpenUsage コミュニティは、元の OpenUsage プロジェクトの独立した、コミュニティによって維持される継続です。
このフォークの目標は、macOS のサポートを維持し、将来の Windows サポートの余地を残しつつ、Linux サポートに重点を置いてクロスプラットフォームの Tauri ベースの方向性を継続することです。
メニュー バーまたはシステム トレイから使用状況を一目で確認できます。ダッシュボードを掘り下げる必要はありません。
OpenUsage コミュニティは現在、次のことに重点を置いています。
安定した AppImage、.deb、および .rpm リリース
さまざまな Linux デスクトップ環境におけるシステム トレイの動作
軽量の Tauri ベースのアーキテクチャを維持する
コ

コミュニティ主導のメンテナンスとプロバイダーの貢献
このプロジェクトは、元の OpenUsage プロジェクトから独立しています。これは、元のプロジェクトが Swift/macOS ファーストの方向に向かって進んでいる一方で、このフォークはクロスプラットフォームのアプローチを継続しているために存在します。
サポートされている場合、アプリは自動更新されます。一度インストールすれば完了です。
OpenUsage Community はシステム トレイで実行されます。
AppIndicator 拡張機能のない GNOME など、StatusNotifierItem/AppIndicator サポートのないデスクトップでは、トレイ アイコンを左クリックしてもパネルが開かないことがあります。代わりに、トレイ メニューの [統計を表示] エントリまたはグローバル ショートカットを使用してください。
Wayland では、パネルの配置はベストエフォートであり、コンポジターに依存します。デスクトップ環境で許可されている場合、パネルはトレイ アイコンの下に表示されます。
システム キーリングに保存されている資格情報を読み取るには、通常、 libsecret または libsecret-tools パッケージによって提供される Secret-tool が必要です。
最新リリースからディストリビューションのアセットを取得します。
新しいリリースが利用可能な場合は、以下の例のバージョンを置き換えてください。
dnf は依存関係を自動的に取り込みます。
sudo dnf install https://github.com/openusage-community/openusage/releases/download/v0.6.24/OpenUsage-0.6.24-1.x86_64.rpm
Debian / Ubuntu
カール -LO https://github.com/openusage-community/openusage/releases/download/v0.6.24/OpenUsage_0.6.24_amd64.deb
sudo apt install ./OpenUsage_0.6.24_amd64.deb
任意の Linux ディストリビューション
ポータブル ビルドには AppImage を使用します。
カール -L -o OpenUsage.AppImage https://github.com/openusage-community/openusage/releases/download/v0.6.24/OpenUsage_0.6.24_amd64.AppImage
chmod +x OpenUsage.AppImage
./OpenUsage.AppImage
.rpm または .deb 経由でインストールした後、アプリのメニューから OpenUsage Community を起動します。アプリはシステム トレイで起動します。
デスクトップ環境にトレイがサポートされていない場合は、AppIndicator をインストールしてください

または StatusNotifier 拡張機能を使用するか、グローバル ショートカットでパネルを開きます。
実行時の依存関係は、.rpm および .deb パッケージによって自動的に処理されます。
システムキーリングから読み取るプロバイダー用のシークレットツール
OpenUsage Community はメニュー バーまたはシステム トレイに常駐し、AI コーディング サブスクリプションをどれだけ使用したかを表示します。
プログレスバー、バッジ、クリアラベル。暗算は必要ありません。
ひと目。すべての AI ツールを 1 つのパネルにまとめました。
常に最新の状態。選択したスケジュールで自動的に更新されます。
グローバルショートカット。カスタマイズ可能なキーボード ショートカットを使用して、どこからでもパネルを切り替えます。
軽量。素早く開き、邪魔になりません。
プラグインベース。アプリ全体を変更せずに、新しいプロバイダーを追加できます。
ローカル HTTP API 。他のアプリは 127.0.0.1:6736 から使用状況データを読み取ることができます。
プロキシのサポート。プロバイダーの HTTP リクエストを SOCKS5 または HTTP プロキシ経由でルーティングします。
アンプ/無料枠、ボーナス、クレジット
クロード / セッション、週次、追加使用量、ccusage によるローカル トークンの使用量
コーデックス / セッション、毎週、レビュー、クレジット
コパイロット / プレミアム、チャット、入力完了
カーソル/クレジット、合計使用量、自動使用量、API 使用量、オンデマンド、CLI 認証
ファクトリー / ドロイド / 標準、プレミアム トークン
Gemini / プロ、フラッシュ、ワークスペース/無料/有料レベル
Grok / 使用したクレジット、プラン、従量課金制の上限
JetBrains AI Assistant / 割り当て、残り
キロ / クレジット、ボーナスクレジット、超過額
OpenCode Go / 5 時間、毎週、毎月の支出制限
ウィンドサーフィン / プロンプト クレジット、フレックス クレジット
Z.ai / セッション、毎週、Web 検索
コミュニティへの貢献は大歓迎です。
リストにないプロバイダーをご希望ですか?問題を開きます。
OpenUsage Community はコミュニティ プロジェクトとして維持されます。
実践的なクロスプラットフォーム開発
特に以下のような貢献を歓迎します。
明確な再現手順を含むバグレポート
UI修正の前後のスクリーンショット
プラグインは現在バンドルされていますが、

プラグイン API は開発中です。長期的な目標は、プロバイダーをより柔軟にして、ユーザーと貢献者が独自の統合を構築して読み込めるようにすることです。
プロバイダーを追加します。各プロバイダーはプラグインとして実装されます。 「プラグイン API」を参照してください。
バグを修正します。プルリクエストは大歓迎です。関連する場合、再現手順とスクリーンショットを提供します。
Linux サポートを改善します。 GNOME、KDE ​​Plasma、Xfce、Wayland、および X11 にわたってテストします。
機能をリクエストします。問題を提起し、ユースケースを説明します。
変更に焦点を当て続けます。フィーチャークリープを回避します。 AI が生成したコードを適切にレビュー、テスト、文書化することなく送信しないでください。
OpenUsage コミュニティは、元の OpenUsage プロジェクトに基づいています。
このフォークは独立しており、コミュニティによって維持されています。 Linux を主な焦点として、Tauri ベースのクロスプラットフォームの方向性を継続します。
元のプロジェクト、その名前、および以前の作品は、元の作成者および貢献者に帰属します。
Robin Ebers によるオリジナルの OpenUsage プロジェクトに基づいています。
@sreipete による CodexBar からインスピレーションを受けています。
OpenUsage のオリジナルのコントリビューター全員と、クロスプラットフォーム サポートの改善に協力してくれた皆さんに感謝します。
警告: メイン ブランチには未リリースの変更が含まれている可能性があります。通常使用する場合は、タグ付きリリースをお勧めします。タグ付きバージョンは公開前にテストされることが期待されます。
ビルドする前に、Tauri が必要とするシステム ライブラリをインストールします。
sudo apt-get install -y libwebkit2gtk-4.1-dev libgtk-3-dev \
libayatana-appindicator3-dev librsvg2-dev patchelf \
libsecret-1-dev ビルド必須
フェドーラ
sudo dnf install -y webkit2gtk4.1-devel gtk3-devel \
libappindicator-gtk3-devel librsvg2-devel \
libsecret-devel patchelf
スタック
OpenUsage コミュニティは次のもので構築されています。
コミュニティ主導の OpenUsage の継続。 Linux、macOS、Windows に最適な AI 使用状況トラッカーを構築します。
読み込み中にエラーが発生しました。リロードしてください

のページです。
アクティビティ
カスタムプロパティ
スター
1
フォーク
レポートリポジトリ
リリース
5
v0.6.28
最新の
2026 年 6 月 10 日
+ 4 リリース
このプロジェクトのスポンサーになる
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Community-driven continuation of OpenUsage. Building the best AI usage tracker for Linux, macOS and Windows. - openusage-community/openusage

GitHub - openusage-community/openusage: Community-driven continuation of OpenUsage. Building the best AI usage tracker for Linux, macOS and Windows. · GitHub
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
openusage-community
/
openusage
Public
forked from robinebers/openusage
Uh oh!
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
458 Commits 458 Commits .agents/ skills .agents/ skills .claude/ skills .claude/ skills .codex/ environments .codex/ environments .cursor .cursor .github .github .local-include .local-include docs docs plugins plugins public public scripts scripts src-tauri src-tauri src src --help --help .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md TRADEMARK.md TRADEMARK.md bun.lock bun.lock bunfig.toml bunfig.toml components.json components.json conductor.json conductor.json copy-bundled.cjs copy-bundled.cjs index.html index.html package.json package.json screenshot.png screenshot.png tsconfig.json tsconfig.json tsconfig.node.json tsconfig.node.json vite.config.ts vite.config.ts View all files Repository files navigation
Track all your AI coding subscriptions in one place.
OpenUsage Community is an independent, community-maintained continuation of the original OpenUsage project.
The goal of this fork is to continue the cross-platform Tauri-based direction with a strong focus on Linux support, while keeping macOS support and leaving room for future Windows support.
See your usage at a glance from your menu bar or system tray. No digging through dashboards.
OpenUsage Community is currently focused on:
Stable AppImage, .deb , and .rpm releases
System tray behavior across different Linux desktop environments
Preserving the lightweight Tauri-based architecture
Community-driven maintenance and provider contributions
This project is independent from the original OpenUsage project. It exists because the original project is moving toward a Swift/macOS-first direction, while this fork continues the cross-platform approach.
The app auto-updates where supported. Install once and you're set.
OpenUsage Community runs in the system tray.
On desktops without StatusNotifierItem/AppIndicator support, such as GNOME without the AppIndicator extension, a left-click on the tray icon may not open the panel. Use the tray menu's Show Stats entry or the global shortcut instead.
On Wayland, panel positioning is best-effort and depends on the compositor. The panel appears under the tray icon where the desktop environment allows it.
Reading credentials stored in the system keyring requires secret-tool , usually provided by the libsecret or libsecret-tools package.
Grab the asset for your distro from the latest release .
Replace the version in the examples below if a newer release is available.
dnf pulls in the dependencies automatically:
sudo dnf install https://github.com/openusage-community/openusage/releases/download/v0.6.24/OpenUsage-0.6.24-1.x86_64.rpm
Debian / Ubuntu
curl -LO https://github.com/openusage-community/openusage/releases/download/v0.6.24/OpenUsage_0.6.24_amd64.deb
sudo apt install ./OpenUsage_0.6.24_amd64.deb
Any Linux distro
Use the AppImage for a portable build:
curl -L -o OpenUsage.AppImage https://github.com/openusage-community/openusage/releases/download/v0.6.24/OpenUsage_0.6.24_amd64.AppImage
chmod +x OpenUsage.AppImage
./OpenUsage.AppImage
After installing via .rpm or .deb , launch OpenUsage Community from your app menu. The app starts in the system tray.
If your desktop environment has no tray support, install an AppIndicator or StatusNotifier extension, or open the panel with the global shortcut.
Runtime dependencies are handled automatically by .rpm and .deb packages:
secret-tool for providers that read from the system keyring
OpenUsage Community lives in your menu bar or system tray and shows how much of your AI coding subscriptions you have used.
Progress bars, badges, clear labels. No mental math required.
One glance. All your AI tools in one panel.
Always up to date. Refreshes automatically on a schedule you pick.
Global shortcut. Toggle the panel from anywhere with a customizable keyboard shortcut.
Lightweight. Opens quickly and stays out of your way.
Plugin-based. New providers can be added without changing the whole app.
Local HTTP API . Other apps can read your usage data from 127.0.0.1:6736 .
Proxy support . Route provider HTTP requests through a SOCKS5 or HTTP proxy.
Amp / free tier, bonus, credits
Claude / session, weekly, extra usage, local token usage with ccusage
Codex / session, weekly, reviews, credits
Copilot / premium, chat, completions
Cursor / credits, total usage, auto usage, API usage, on-demand, CLI auth
Factory / Droid / standard, premium tokens
Gemini / pro, flash, workspace/free/paid tier
Grok / credits used, plan, pay-as-you-go cap
JetBrains AI Assistant / quota, remaining
Kiro / credits, bonus credits, overages
OpenCode Go / 5h, weekly, monthly spend limits
Windsurf / prompt credits, flex credits
Z.ai / session, weekly, web searches
Community contributions are welcome.
Want a provider that is not listed? Open an issue.
OpenUsage Community is maintained as a community project.
practical cross-platform development
Contributions are welcome, especially:
bug reports with clear reproduction steps
before/after screenshots for UI fixes
Plugins are currently bundled while the plugin API is being developed. The long-term goal is to make providers more flexible so users and contributors can build and load their own integrations.
Add a provider. Each provider is implemented as a plugin. See the Plugin API .
Fix a bug. Pull requests are welcome. Provide reproduction steps and screenshots where relevant.
Improve Linux support. Test across GNOME, KDE Plasma, Xfce, Wayland, and X11.
Request a feature. Open an issue and explain the use case.
Keep changes focused. Avoid feature creep. Do not submit AI-generated code without reviewing, testing, and documenting it properly.
OpenUsage Community is based on the original OpenUsage project.
This fork is independent and community-maintained. It continues the Tauri-based cross-platform direction, with Linux as the primary focus.
The original project, its name, and its prior work are credited to the original author and contributors.
Based on the original OpenUsage project by Robin Ebers .
Inspired by CodexBar by @steipete .
Thanks to all original OpenUsage contributors and everyone helping improve cross-platform support.
Warning : The main branch may contain unreleased changes. For regular use, prefer tagged releases. Tagged versions are expected to be tested before publishing.
Install the system libraries Tauri needs before building.
sudo apt-get install -y libwebkit2gtk-4.1-dev libgtk-3-dev \
libayatana-appindicator3-dev librsvg2-dev patchelf \
libsecret-1-dev build-essential
Fedora
sudo dnf install -y webkit2gtk4.1-devel gtk3-devel \
libappindicator-gtk3-devel librsvg2-devel \
libsecret-devel patchelf
Stack
OpenUsage Community is built with:
Community-driven continuation of OpenUsage. Building the best AI usage tracker for Linux, macOS and Windows.
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
1
fork
Report repository
Releases
5
v0.6.28
Latest
Jun 10, 2026
+ 4 releases
Sponsor this project
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
