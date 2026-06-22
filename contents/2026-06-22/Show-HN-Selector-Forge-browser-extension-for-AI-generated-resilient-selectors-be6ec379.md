---
source: "https://github.com/Intuned/selector-forge"
hn_url: "https://news.ycombinator.com/item?id=48630515"
title: "Show HN: Selector Forge – browser extension for AI-generated resilient selectors"
article_title: "GitHub - Intuned/selector-forge: Browser extension to create reliable selectors (CSS and Xpath) using AI · GitHub"
author: "ahmadilaiwi"
captured_at: "2026-06-22T14:53:10Z"
capture_tool: "hn-digest"
hn_id: 48630515
score: 11
comments: 0
posted_at: "2026-06-22T14:21:24Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Selector Forge – browser extension for AI-generated resilient selectors

- HN: [48630515](https://news.ycombinator.com/item?id=48630515)
- Source: [github.com](https://github.com/Intuned/selector-forge)
- Score: 11
- Comments: 0
- Posted: 2026-06-22T14:21:24Z

## Translation

タイトル: Show HN: Selector Forge – AI 生成の回復力のあるセレクターのブラウザー拡張機能
記事タイトル: GitHub - Intuned/selector-forge: AI を使用して信頼性の高いセレクター (CSS および Xpath) を作成するためのブラウザー拡張機能 · GitHub
説明: AI を使用して信頼性の高いセレクター (CSS および Xpath) を作成するためのブラウザー拡張機能 - Intuned/selector-forge
HN テキスト: こんにちは、HN、Intuned ( https://intunedhq.com ) チームの Ahmad です。本日、AI を使用して信頼性の高い CSS/XPath セレクターを生成するブラウザー拡張機能である Selector Forge ( https://selectorforge.ai/ ) をリリースし、オープンソース化します。これを使用して、単一の要素または要素の配列のセレクターを作成できます。作成されるセレクターは「セマンティック」であり、Chrome DevTool の「Copy Selector」（および他の同様の拡張機能）が提供するものよりもページ変更に対する復元力が高くなります。これらは、`#top > div.w-100.ph0-l.ph3.ph4-m > h1 > span` のような脆弱なものを渡す傾向があり、最小限のページ変更で壊れる可能性があります。 Selector Forge は、壊れにくいセレクターを目指しています。以下に、Selector Forge が作成したセレクターをいくつか示します。 `//div[@aria-label="Showing Weekly downloads"]//p[@aria-live="polite"]` (項目セレクター) と `//*[local-name()='svg' および @aria-label="Download統計"]/following-sibling::div` (リスト セレクター) です。拡張機能の使用方法のビデオデモは次のとおりです: https://www.youtube.com/watch?v=8IjjeDQkKmo Chrome の Selector Forge: https://chromewebstore.google.com/detail/lbendfnlmhdakbeblaj... Firefox の Selector Forge:
https://addons.mozilla.org/en-US/firefox/addon/selector-forg... Selector Forge コード:
https://github.com/Intuned/selector-forge バックストーリー: 過去 2 年間、私たちはブラウザ自動化を構築および保守するためのコーディング エージェントである Intuned Agent を構築してきました。私たちはすぐに、最も壊れやすい部分を理解しました。

通常、ブラウザー コードで最も重要なのはセレクターであり、優れたセレクターを作成することは、自動化自体の品質と信頼性の向上に大いに役立ちます。そこで、セレクターの作成を独自のエージェントに抽象化し、それをツールとしてラップし、codegen エージェントから呼び出せるようにしました。デフォルトでは、LLM は適切なセレクターを生成するのにあまり機能しません。そのため、これは非常に便利であり、エージェントが生成するコードを改善しました。私たちは最近、この部分 (セレクター エージェント/作成物) はそれ自体 (私たちのプラットフォームの外で) 役立つと考えたので、ブラウザー拡張機能としてパッケージ化しました。 That’s this post! Selector Forge はオープン ソースであり、ブラウザ ストア (Chrome および Firefox) のバージョンは、月あたり最大 200 セレクターに対して無料です。無制限の使用は有料プランの一部です。ほとんどの開発者はもうこの種のコードを手動で記述していないことを認識しているため、次のステップでは、コーディング エージェントが CLI または MCP 経由で直接呼び出すことができる方法でこの機能を公開します。私たちのロードマップは次のとおりです: https://github.com/Intuned/selector-forge#roadmap 皆様のご意見、ご質問、フィードバックをお待ちしております。

記事本文:
GitHub - Intuned/selector-forge: AI を使用して信頼性の高いセレクター (CSS および Xpath) を作成するためのブラウザー拡張機能 · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
調律された
/
セレクターフォージ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション最適化

オン
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
28 コミット 28 コミット .github/ workflows .github/ workflows .ladle .ladle アセット アセット e2e e2e エントリポイント エントリポイント lib lib public/ icon public/ icon スクリプト スクリプト ストーリー ストーリー テスト テスト .gitignore .gitignore ARCHITECTURE.md ARCHITECTURE.md ライセンス ライセンス README.md README.md package.json package.json playwright.config.ts playwright.config.ts tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts wxt.config.ts wxt.config.ts 糸.ロック 糸.ロック すべてのファイルを表示 リポジトリ ファイルのナビゲーション
任意のページ上の要素を選択すると、信頼性の高いセレクターが返されます。AI によって生成および判断され、実際に表示される前にライブ DOM に対して再検証されます。
Selector Forge は、見ているページから直接堅牢な CSS または XPath セレクターを構築するのに役立つスタンドアロンのブラウザー拡張機能 (Chrome および Firefox、MV3) です。あなたは望むものを指さします。拡張機能と Intuned のセレクター バックエンドが残りの作業を実行します。つまり、候補を提案し、実際のページに対してテストし、正しく解決されないものはすべて破棄します。
これは、エンドツーエンドのテストを作成したり、スクレイパーを構築したり、脆弱なセレクターによって後でコストがかかるページを自動化したりする場合に役立ちます。
任意のページを開いて拡張機能をクリックします。
選択モードを選択し、ライブ ページ上で要素を直接選択します。
この拡張機能は、選択したもの (選択したターゲット、DOM コンテキスト、シード候補) のコンパクトなスナップショットをキャプチャし、バックエンドに送信します。
バックエンドは候補セレクターを提案し、ランク付けします。この拡張機能は、すべての候補をライブ DOM に対してテストし、結果をフィードバックします。
このループは、バックエンドが勝者に決まるまで繰り返されます。
ポップアップには再検証されたセレクターのみが表示され、それぞれにコピー ボタンが付いています。
ブラウザは常に真実の情報源です

r セレクターが実際に一致するもの。 AI が提案してランク付けします。正しさについて最終的な判断が下されることはありません。
拡張機能は、ループの継続性のソースであるセレクター作成セッション状態を保持します。
ブラウザは真実の情報源であり、すべての結果に対して再検証が必須です。
AI がセレクターを提案し、ランク付けします。正しさを証明するものではありません。
リストの場合、検証では目的のセット全体がチェックされるため、過剰一致および過小一致のセレクターは拒否されます。
モジュール マップ、メッセージング レイヤー、背景/コンテンツ/ポップアップ コンテキスト、および認証と CLI の継ぎ目については、ARCHITECTURE.md を参照してください。
モード
そうですよね
あなたは得ます
シングル
要素を 1 つ選択してください
ボタン、入力、リンク、ラベル、一回限りのターゲットなど、その要素そのもののセレクター候補を検証しました。
リスト
繰り返しセットから 2 つの例を選択します
フルセットの検証済みコンテナセレクター。保存する前にプレビューされます。
開発クイックスタート
yarn install # `wxt prepare` も実行します
糸 dev # watch + Chrome で .output/chrome-mv3 をロード (解凍)
糸 dev:firefox # Firefox も同様
最初のyarn dev の後、解凍された拡張機能を .output/chrome-mv3 の chrome://extensions からロードします (開発者モードを有効にします)。
コマンド
何をするのか
糸開発 / 糸開発:Firefox
ビルドを監視、解凍された拡張機能としてロード可能
糸コンパイル
tsc --noEmit タイプチェック
糸テスト
Vitest — ユニット + リアル Chromium ブラウザ プロジェクト
糸ビルド / 糸ビルド:Firefox
実稼働拡張バンドル
糸のビルド:e2e
<all_urls> ホスト権限を持つ E2E バリアント - これを出荷しないでください
糸e2e
build:e2e、次にパッケージ化された拡張機能に対する Playwright
ヤーン ジップ / ヤーン ジップ:Firefox
ストア対応のzip
糸のアイコン
アイコンアセットを再生成する
糸ひしゃく
http://localhost:61010 でポップアップ コンポーネントを単独でプレビューします。
コンポーネントのプレビュー (レードル)
ヤーンレードルは、設計のためにポップアップの React コンポーネントを分離して提供します

拡張機能のリロードも実際のバックエンドもありません。ストーリーは stories/ ( *.stories.tsx ) に存在します。お玉の設定は .ladle/ にあります。ポップアップは WXT の挿入されたブラウザー グローバルを予期しているため、 .ladle/wxt-globals.ts はそのための no-op スタブをインストールします。 yarn ladle:build は dist/ladle の下に静的バンドルを生成します。
ユニット — セレクター ロジック、状態変換、ストレージ、および決定論的フォールバックの高速 Vitest テスト (ノード/happy-dom)。
ブラウザ — 実際の DOM に対してセレクター生成を実行し、各候補が期待される要素セットに正確に解決されることを証明する Vitest ブラウザモード テスト。これが正しさの神託です。両方の層は糸テストの下で実行されます。
E2E — 実際のページ、ポインター フロー、ポップアップ、コンテンツ スクリプト、およびバックグラウンド ワーカーを備えたパッケージ化された MV3 拡張機能に対する Playwright。糸 e2e で実行します。
エントリポイント/
background.ts バックグラウンド サービス ワーカー - セッション状態、エージェント ループ、ネットワーク I/O
content.ts コンテンツ スクリプト — ピッカー オーバーレイ、DOM アクセス、セレクター テスト
Popup/ React Popup — モードコントロール、結果、コピーアクション
ライブラリ/
エージェント/エージェント ループ コントローラー (バックエンド ターンテイキング)
コンテンツ/ピッカー オーバーレイ、要素レジストリ、DOM インスペクション
バックグラウンド/ハンドラー、コンテキスト メニュー、セッション接続、CLI ブリッジ
メッセージング/型付き、方向分割されたランタイム メッセージ プロトコル
状態/セッション状態、履歴、スキーマ、設定
認証/認証クライアント + トークン処理
graphql/ワークスペース + 使用状況のクエリ
config.ts API ベース + ランタイム構成
テスト/vitest (ユニット + ブラウザ)
e2e/ ビルドされた拡張機能に対する playwright
ARCHITECTURE.md モジュール マップ、信頼境界、エージェント ループ、シーム
ポップアップ用に React を使用して WXT 上に構築されています。
CLI 制御 - Intuned CLI から拡張機能を駆動します。Intuned IDE サポート、エンドツーエンドのテストと自動化を実行するローカル エージェント、MCP を介した公開。 (基本的な配線 - タブの権限と

CDP 主導のセッション開始 — はすでに導入されています。)
スマート ピッカー — 1 つのフローで多くの要素を選択し、拡張機能でそれらを単一の項目とリストのようなセットにグループ化できる複数モードに加え、ページに有用なフィールド、名前、セレクターを自動的に提案する AI フィールド検出を備えています。
ドリルダウン モード — 選択後の精度の向上: XPath/DOM ツリーを実際に意図した要素 (子スパン → ボタン → 行 → ラベル → 親コンテナ) まで移動し、リスト選択を親レベルまたは子レベルに移動し、必要な例を追加するか、間違った例を除外します。
独自のバックエンドを導入する — 現在、拡張機能は認証とセレクターの生成のために Intuned と通信します。私たちは、信頼性の高いセレクターを生成および判断するオープンソース エージェントを含む、その継ぎ目に組み込まれ、Intuned を完全に置き換える小型の自己ホスト可能なリファレンス バックエンドを出荷する予定です。これにより、独自のインフラストラクチャでループ全体を実行できるようになります。
さらに言えば、セレクター/オートメーション履歴、Playwright またはプレーン JavaScript へのエクスポート、自動ページネーション検出、クロス iframe/shadow-DOM サポート。
問題やプルリクエストは大歓迎です。 PR を開く前に、yarn コンパイルとyarn テストを実行してください。
AI を使用して信頼性の高いセレクター (CSS および Xpath) を作成するためのブラウザー拡張機能
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
0
フォーク
レポートリポジトリ
リリース
2
セレクターフォージ v0.0.3
最新の
2026 年 6 月 22 日
+ 1 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Browser extension to create reliable selectors (CSS and Xpath) using AI - Intuned/selector-forge

Hi HN, I'm Ahmad from the Intuned ( https://intunedhq.com ) team. Today, we're releasing and open-sourcing Selector Forge ( https://selectorforge.ai/ ), a browser extension that generates reliable CSS/XPath selectors using AI. You can use it to create a selector for a single element or for an array of elements. The selectors it creates are meant to be "semantic" and more resilient to page changes than what Chrome DevTool’s “Copy Selector” (and other similar extensions) give you. Those tend to hand you something brittle like `#top > div.w-100.ph0-l.ph3.ph4-m > h1 > span`, which can break with a minimal page change. Selector Forge aims for selectors that don't break as easily. Here are some selectors that Selector Forge created: `//div[@aria-label="Showing weekly downloads"]//p[@aria-live="polite"]` (item selector) and `//*[local-name()='svg' and @aria-label="Download statistics"]/following-sibling::div` (list selector). Here is a video demo of using the extension: https://www.youtube.com/watch?v=8IjjeDQkKmo Selector Forge on Chrome: https://chromewebstore.google.com/detail/lbendfnlmhdakbeblaj... Selector Forge on Firefox:
https://addons.mozilla.org/en-US/firefox/addon/selector-forg... Selector Forge code:
https://github.com/Intuned/selector-forge Backstory: For the past couple of years we've been building Intuned Agent, a coding agent for building and maintaining browser automations. We quickly figured out that the most fragile part of any browser code is usually the selectors and that creating good selectors can go a long way towards improving the quality and reliability of the automation itself. So we abstracted selector creation into its own agent, wrapped it as a tool, and let our codegen agent call it. LLMs by default don't do a great job generating good selectors, so this turned out to be really useful and improved the code our agent generates. We recently thought that this piece (the selector agent/creation) is useful on its own (outside our platform) so we packaged it as a browser extension. That’s this post! Selector Forge is open source, and the version in the browser stores (Chrome and Firefox) is free for up to 200 selectors/month. Unlimited usage is part of our paid plans. We realize most developers aren't writing this kind of code by hand anymore, so the next step is exposing this functionality in a way coding agents can call directly, over a CLI or MCP. Here's our roadmap: https://github.com/Intuned/selector-forge#roadmap Excited to hear your thoughts, questions, and feedback!

GitHub - Intuned/selector-forge: Browser extension to create reliable selectors (CSS and Xpath) using AI · GitHub
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
Intuned
/
selector-forge
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
28 Commits 28 Commits .github/ workflows .github/ workflows .ladle .ladle assets assets e2e e2e entrypoints entrypoints lib lib public/ icon public/ icon scripts scripts stories stories tests tests .gitignore .gitignore ARCHITECTURE.md ARCHITECTURE.md LICENSE LICENSE README.md README.md package.json package.json playwright.config.ts playwright.config.ts tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts wxt.config.ts wxt.config.ts yarn.lock yarn.lock View all files Repository files navigation
Pick an element on any page, get back a reliable selector — generated and judged by AI, then re-verified against the live DOM before you ever see it.
Selector Forge is a standalone browser extension (Chrome & Firefox, MV3) that helps you build robust CSS or XPath selectors directly from the pages you're looking at. You point at what you want; the extension and Intuned 's selector backend do the rest — proposing candidates, testing them against the real page, and discarding anything that doesn't resolve correctly.
It's useful for writing end-to-end tests, building scrapers, and automating any page where a brittle selector would cost you later.
Open any page and click the extension.
Choose a selection mode and pick element(s) directly on the live page.
The extension captures a compact snapshot of your picks (selected targets, DOM context, seed candidates) and sends it to the backend.
The backend proposes and ranks candidate selectors; the extension tests every candidate against the live DOM and feeds the results back.
This loop repeats until the backend settles on a winner.
The popup shows only re-verified selectors, each with a copy button.
The browser is always the source of truth for what a selector actually matches. The AI proposes and ranks; it never gets the final word on correctness.
The extension holds the selector-creation session state — the source of continuity for the loop.
The browser is the source of truth — re-verification is mandatory for every result.
The AI proposes and ranks selectors; it does not prove correctness.
For lists, verification checks the full intended set, so over-matching and under-matching selectors are rejected.
See ARCHITECTURE.md for the module map, the messaging layer, the background/content/popup contexts, and the auth + CLI seams.
Mode
You do
You get
Single
Pick one element
Verified selector candidates for that exact element — buttons, inputs, links, labels, one-off targets.
List
Pick two examples from a repeating set
A verified container selector for the full set, previewed before you save it.
Dev quickstart
yarn install # also runs `wxt prepare`
yarn dev # watch + load .output/chrome-mv3 in Chrome (unpacked)
yarn dev:firefox # same for Firefox
After the first yarn dev , load the unpacked extension from .output/chrome-mv3 at chrome://extensions (enable Developer mode).
Command
What it does
yarn dev / yarn dev:firefox
Watch build, loadable as an unpacked extension
yarn compile
tsc --noEmit typecheck
yarn test
Vitest — unit + real-Chromium browser projects
yarn build / yarn build:firefox
Production extension bundle
yarn build:e2e
E2E variant with <all_urls> host permission — never ship this
yarn e2e
build:e2e then Playwright against the packaged extension
yarn zip / yarn zip:firefox
Store-ready zip
yarn icons
Regenerate icon assets
yarn ladle
Preview popup components in isolation at http://localhost:61010
Component previews (Ladle)
yarn ladle serves the popup's React components in isolation for design and review — no extension reload, no real backend. Stories live in stories/ ( *.stories.tsx ); Ladle config is in .ladle/ . The popup expects WXT's injected browser global, so .ladle/wxt-globals.ts installs a no-op stub for it. yarn ladle:build produces a static bundle under dist/ladle .
Unit — fast Vitest tests (node/happy-dom) for selector logic, state transforms, storage, and deterministic fallbacks.
Browser — Vitest browser-mode tests that run selector generation against a real DOM and prove each candidate resolves to exactly the expected element set. This is the correctness oracle. Both layers run under yarn test .
E2E — Playwright against the packaged MV3 extension with a real page, pointer flow, popup, content script, and background worker. Run with yarn e2e .
entrypoints/
background.ts background service worker — session state, agent loop, network I/O
content.ts content script — picker overlay, DOM access, selector testing
popup/ React popup — mode controls, results, copy actions
lib/
agent/ agent loop controller (backend turn-taking)
content/ picker overlay, element registry, DOM inspection
background/ handlers, context menu, session wiring, CLI bridge
messaging/ typed, direction-partitioned runtime-message protocol
state/ session state, history, schema, preferences
auth/ auth client + token handling
graphql/ workspace + usage queries
config.ts API base + runtime config
tests/ vitest (unit + browser)
e2e/ playwright against the built extension
ARCHITECTURE.md module map, trust boundary, agent loop, seams
Built on WXT with React for the popup.
CLI control — drive the extension from the Intuned CLI: Intuned IDE support, local agents running end-to-end tests and automations, and exposure through MCP. (Foundational wiring — the tabs permission and CDP-driven session start — is already in place.)
Smart picker — a multiple mode that lets you select many elements in one flow and have the extension group them into single items and list-like sets, plus AI field detection that suggests useful fields, names, and selectors for a page automatically.
Drill-down modes — precision refinement after a pick: walk the XPath/DOM tree to the element you actually meant (child span → button → row → label → parent container), move a list selection to a parent or child level, and add required examples or exclude wrong ones.
Bring your own backend — today the extension talks to Intuned for authentication and selector generation. We plan to ship a small, self-hostable reference backend that drops into that seam and replaces Intuned entirely — including an open-source agent that generates and judges reliable selectors — so you can run the whole loop on your own infrastructure.
Further out: selector/automation history, export to Playwright or plain JavaScript, automatic pagination detection, and cross-iframe / shadow-DOM support.
Issues and pull requests are welcome. Please run yarn compile and yarn test before opening a PR.
Browser extension to create reliable selectors (CSS and Xpath) using AI
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
2
Selector Forge v0.0.3
Latest
Jun 22, 2026
+ 1 release
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
