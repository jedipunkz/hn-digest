---
source: "https://github.com/fresswolf/Slopera"
hn_url: "https://news.ycombinator.com/item?id=48845268"
title: "Show HN: Slopera, a browser that hallucinates every page with an LLM"
article_title: "GitHub - fresswolf/Slopera: The browser for the slop era · GitHub"
author: "fresswolf"
captured_at: "2026-07-09T13:43:58Z"
capture_tool: "hn-digest"
hn_id: 48845268
score: 3
comments: 0
posted_at: "2026-07-09T13:10:14Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Slopera, a browser that hallucinates every page with an LLM

- HN: [48845268](https://news.ycombinator.com/item?id=48845268)
- Source: [github.com](https://github.com/fresswolf/Slopera)
- Score: 3
- Comments: 0
- Posted: 2026-07-09T13:10:14Z

## Translation

タイトル: HN を表示: Slopera、LLM ですべてのページを幻覚させるブラウザ
記事タイトル: GitHub - fresswolf/Slopera: スロップ時代のブラウザ · GitHub
説明: スロップ時代のブラウザ。 GitHub でアカウントを作成して、fresswolf/Slopera の開発に貢献してください。
HN テキスト: 私は、実際の Web に接続する代わりに、LLM が想定するとおりに各ページをレンダリングするデスクトップ ブラウザーである Slopera を構築しました。 404 が発生することはありません。基本的に、あなたまたは LLM が (偽の Google 結果などで) 思いついた URL はすべて存在します。リンクをクリックすると、オンデマンドで次のページが作成されます。画像やインタラクティブな JS コンテンツも生成されます。 Electron + TypeScript です。実行するには Anthropic キーまたは OpenRouter キーが必要です。ローカル モデルのサポートが計画されており、これにより完全にオフラインのブラウザーになります。名前はオペラ/スロップ/時代をもじったものです。

記事本文:
GitHub - fresswolf/Slopera: スロップ時代のブラウザ · GitHub
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
フレスウルフ
/
スローペラ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード もっと開く

アクション メニュー フォルダーとファイル
30 コミット 30 コミット .github/ workflows .github/ workflows build build scripts scripts src src testing testing .gitignore .gitignore 1998.png 1998.png CLAUDE.md CLAUDE.md ライセンス ライセンス README.md README.md SPEC.md SPEC.md childlike.png childlike.png 電子ビルダー.yml 電子ビルダー.yml 電子.vite.config.1781460665864.mjs 電子.vite.config.1781460665864.mjs 電子.vite.config.ts 電子.vite.config.ts eslint.config.js eslint.config.js Fishrain.gif Fishrain.gif game.gif game.gif logo.png logo.png package-lock.json package-lock.json package.json package.json playwright.config.ts playwright.config.ts tsconfig.json tsconfig.json tsconfig.node.json tsconfig.node.json tsconfig.web.json tsconfig.web.json vitest.config.ts vitest.config.ts すべてのファイルを表示リポジトリ ファイルのナビゲーション
Slopera は、実際の Web には決して触れないデスクトップ ブラウザです。任意の URL (本物または創作) を入力すると、LLM によってそのページがオンザフライで幻覚表示されます。
上記の Google と Reddit は本物ではありません。すべてのページ、そのテキストと画像はすべてその場で作成されます。
インタラクティブな Web アプリでも動作します
上の HTML ゲームは AI によって考案されました。
レンズは、ウェブ全体が夢見るレジスターです。
ツールバーのドロップダウン:
ストレート : 現実の現場を真剣に再現し、もっともらしいことを詰め込んだもの
これまで存在しなかったコンテンツ。
追加のスロップ: 最大限のアルゴリズムに汚染されたインターネット: 不可能な製品、
リストの数が減り、CTA ボタンが多すぎます。
1998 : すべてのサイトが Geocities 時代の姿に: テーブル レイアウト、ヒット
カウンタ、「Netscape Navigator 4 で最もよく表示される」。
Childlike : 6 歳児がクレヨンで描いたようなページ。
[設定] で独自のレンズを作成することもできます。名前と短いものを付けます。
フレーバー プロンプトが表示され、ビルトインの横にドロップダウンに表示されます。それぞれ
レンズキープ

夢見たページの独自のキャッシュがあるため、レンズを切り替えると、
別の現実で同じ URL を再訪問します。
macOS、Windows、または Linux のインストーラーを次の場所から入手します。
リリースページ 、またはから実行
ソース:
npmインストール
npm 実行開発
いずれの場合も、設定 (歯車アイコン) を開き、以下を貼り付けます。
Anthropic または OpenRouter からの API キー。必須、ページを夢見る
fal.ai キー。オプションで、イメージを生成します。 OpenRouter キーを使用すると、代わりに OpenRouter 経由で提供される代替 (ただし低速な) イメージ モデルから選択できます。どちらもないと、画像はキャプション付きのプレースホルダーに劣化します。
キーは OS キーチェーン (safeStorage) を介して暗号化されて保存され、
それぞれの API に送信されます。すでに夢見ていたページの閲覧には費用はかかりません。
ご自身の責任でご使用ください。夢見たばかりのページと画像はすべて有料 API です
通話はキーに請求され、生成されたコンテンツはモデレートされていない LLM 出力です。
OpenRouter : https://openrouter.ai/ にアクセスして「サインアップ」をクリックし、クレジットを追加します。
https://openrouter.ai/settings/credits で新しい API キーを作成します。
https://openrouter.ai/workspaces/default/keys 。
fal.ai : https://fal.ai/ に移動し、「始める」をクリックしてアカウントを作成し、次のアドレスにチャージします。
https://fal.ai/dashboard/usage-billing/credits にアクセスして API キーを作成します。
https://fal.ai/dashboard/keys 。
Claude (Anthropic) : https://platform.claude.com/ でアカウントを作成し、資金を追加してください。
https://platform.claude.com/dashboard でキーを作成します。
https://platform.claude.com/settings/workspaces/default/keys 。
すべてのページは新しい LLM 生成であるため、モデルの設定に応じて異なります。
Slopera はトークンをすぐに焼き尽くしてしまう可能性があります。 [設定] でトレードオフを選択します。
速くて安い : Claude Haiku は現在、最速のテキスト モデルです。
妥当なトークンコスト、および FLUX Schnell (fal.ai 経由でのみ利用可能)
画像に関しては超高速かつ安価です。良いデフォルトとしては、

カジュアルなブラウジング。
高忠実度 : リッチでインタラクティブなページ (動作する Web アプリ、ゲーム) 用
クロード・オーパスのようなより大きなモデルとハイエンドのモデルが必要になる場合があります。
GPT Image 2 などの画像モデル。ページあたりの速度が著しく遅くなり、料金も高くなります。 GLM 5.2 は、比較的低コストで高い忠実度を提供することが示されていますが、速度が遅いです。
npm run typecheck # strict TS、メイン + レンダラー
npm 実行 lint # eslint
npm テスト # vitest: オムニボックスの解析、フェンスストリッピング、プロンプト、抽出
npm run build #electron-vite 本番ビルド
npm run test:e2e # playwright: ビルドされたアプリを起動します (最初に `npm run build` を実行します)
npm run package:mac # .dmg into release/
npm run package:win # NSIS インストーラー (x64) を release/ に追加
npm run package:linux # AppImage (x64) into release/
npm run icons # logo.png からアプリアイコンを再生成
SLOPERA_FAKE_GEN=1 npm run dev は、缶詰に対してブラウザ全体を実行します。
オフライン ジェネレーター — UI 作業に便利で、e2e テストで使用されます。 CI
(GitHub Actions、 .github/workflows/release.yml ) lint、typecheck、および
プッシュごとに単体テストを行い、macOS、Windows、Linux 用のインストーラーを構築します。
v* タグをプッシュすると、それらがドラフト GitHub リリースに収集されます。
src/main/ タブ、プロトコル ハンドラー、生成、ストア
src/preload/ 型付き IPC ブリッジ
src/renderer/ブラウザクローム (React)
src/shared/ 純粋なロジック: URL 処理、フェンスストリッピング、レンズ、タイプ
テスト/vitest 単体テスト + Playwright スモーク テスト
SPEC.md の全機能仕様とアーキテクチャの決定
ライセンス
マサチューセッツ工科大学Slopera がレンダリングするものは何も本物ではありません。実際のウェブサイトとの類似点、
生きているか死んでいるか、それがポイントだ。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
v1.1.0
最新の
2026 年 7 月 9 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
©

2026年 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

The browser for the slop era. Contribute to fresswolf/Slopera development by creating an account on GitHub.

I built Slopera, a desktop browser that, instead of connecting to the real web, renders each page as what an LLM thinks it might look like. You'll never get a 404, basically any URL you or the LLM (for example, in fake google results) comes up with exists. Clicking a link creates the next page on demand. Images and interactive JS content are generated too. It's Electron + TypeScript. You need an Anthropic or OpenRouter key to run it. Support for local models is planned, which would make this a fully offline browser. The name is a pun on Opera / slop / era.

GitHub - fresswolf/Slopera: The browser for the slop era · GitHub
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
fresswolf
/
Slopera
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
30 Commits 30 Commits .github/ workflows .github/ workflows build build scripts scripts src src tests tests .gitignore .gitignore 1998.png 1998.png CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md SPEC.md SPEC.md childlike.png childlike.png electron-builder.yml electron-builder.yml electron.vite.config.1781460665864.mjs electron.vite.config.1781460665864.mjs electron.vite.config.ts electron.vite.config.ts eslint.config.js eslint.config.js fishrain.gif fishrain.gif game.gif game.gif logo.png logo.png package-lock.json package-lock.json package.json package.json playwright.config.ts playwright.config.ts tsconfig.json tsconfig.json tsconfig.node.json tsconfig.node.json tsconfig.web.json tsconfig.web.json vitest.config.ts vitest.config.ts View all files Repository files navigation
Slopera is a desktop browser that never touches the real web. Type any URL (real or invented) and the page is hallucinated on the fly by an LLM.
The Google and Reddit above are not real. Every page, all its text and images, is made up on the spot.
Also works with interactive web apps
The HTML game above was dreamed up by AI.
A lens is the register the whole web gets dreamed in. Pick one from the
toolbar dropdown:
Straight : earnest renditions of real sites, filled with plausible
content that never existed.
Extra slop : maximum algorithm-poisoned internet: impossible products,
listicles that lose count, too many call-to-action buttons.
1998 : every site as its Geocities-era self: table layouts, hit
counters, "best viewed in Netscape Navigator 4".
Childlike : pages as drawn by a 6-year-old with crayons.
You can also create your own lens in Settings : give it a name and a short
flavor prompt, and it appears in the dropdown alongside the built-ins. Each
lens keeps its own cache of dreamed pages, so switching lenses lets you
revisit the same URL in a different reality.
Grab an installer for macOS, Windows or Linux from the
Releases page , or run from
source:
npm install
npm run dev
Either way, open Settings (gear icon) and paste:
an API key from Anthropic or OpenRouter . Required, dreams the pages
a fal.ai key . Optional, generates images. With an OpenRouter key you can instead pick from alternative (but slower) image models served via OpenRouter. Without either, images degrade into captioned placeholders.
Keys are stored encrypted via the OS keychain ( safeStorage ) and only ever
sent to their respective APIs. Browsing already-dreamed pages costs nothing.
Use at your own risk. Every freshly dreamed page and image is a paid API
call billed to your keys, and generated content is unmoderated LLM output.
OpenRouter : go to https://openrouter.ai/ and click "Sign Up", add credits at
https://openrouter.ai/settings/credits , then create a new API key at
https://openrouter.ai/workspaces/default/keys .
fal.ai : go to https://fal.ai/ and click "Get Started" to create an account, top up at
https://fal.ai/dashboard/usage-billing/credits , then create an API key at
https://fal.ai/dashboard/keys .
Claude (Anthropic) : create an account at https://platform.claude.com/ , add funds through
https://platform.claude.com/dashboard , then create a key at
https://platform.claude.com/settings/workspaces/default/keys .
Every page is a fresh LLM generation, so depending on your model settings
Slopera can burn through tokens quickly. Pick your trade-off in Settings:
Fast & cheap : Claude Haiku is currently the fastest text model at a
reasonable token cost, and FLUX schnell (only available via fal.ai) is
super fast and cheap for images. A good default for casual browsing.
High fidelity : for rich, interactive pages (working web apps, games)
you may need bigger models like Claude Opus , paired with a higher-end
image model such as GPT Image 2 . Noticeably slower and pricier per page. GLM 5.2 has been shown to give high fidelity at a comparably low cost, but it's slow.
npm run typecheck # strict TS, main + renderer
npm run lint # eslint
npm test # vitest: omnibox parsing, fence-stripping, prompts, extraction
npm run build # electron-vite production build
npm run test:e2e # playwright: boots the built app (run `npm run build` first)
npm run package:mac # .dmg into release/
npm run package:win # NSIS installer (x64) into release/
npm run package:linux # AppImage (x64) into release/
npm run icons # regenerate app icons from logo.png
SLOPERA_FAKE_GEN=1 npm run dev runs the whole browser against a canned
offline generator — useful for UI work and used by the e2e test. CI
(GitHub Actions, .github/workflows/release.yml ) runs lint, typecheck and
unit tests on every push, and builds installers for macOS, Windows and Linux;
pushing a v* tag collects them into a draft GitHub Release.
src/main/ tabs, protocol handlers, generation, stores
src/preload/ typed IPC bridge
src/renderer/ browser chrome (React)
src/shared/ pure logic: URL handling, fence-stripping, lenses, types
tests/ vitest unit tests + playwright smoke test
SPEC.md full feature spec & architecture decisions
License
MIT. Nothing Slopera renders is real; any resemblance to actual websites,
living or dead, is the point.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
v1.1.0
Latest
Jul 9, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
