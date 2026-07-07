---
source: "https://github.com/mi60dev/visionaire-engine"
hn_url: "https://news.ycombinator.com/item?id=48822611"
title: "Doing front end changes via LLMs is pain in the a*"
article_title: "GitHub - mi60dev/visionaire-engine: Eyes for AI coding agents: deterministic MCP server that tells the LLM which CSS rule wins, in which file, on which line — and why. Cascade verdicts, blast radius, interaction timelines, pixel-perfect audits. · GitHub"
author: "mishonyAI"
captured_at: "2026-07-07T19:42:14Z"
capture_tool: "hn-digest"
hn_id: 48822611
score: 1
comments: 1
posted_at: "2026-07-07T19:34:54Z"
tags:
  - hacker-news
  - translated
---

# Doing front end changes via LLMs is pain in the a*

- HN: [48822611](https://news.ycombinator.com/item?id=48822611)
- Source: [github.com](https://github.com/mi60dev/visionaire-engine)
- Score: 1
- Comments: 1
- Posted: 2026-07-07T19:34:54Z

## Translation

タイトル: LLM を介してフロントエンドの変更を行うのは苦痛です*
記事のタイトル: GitHub - mi60dev/visionaire-engine: AI コーディング エージェントの目: どの CSS ルールが、どのファイルで、どの行で、そしてその理由で優先されるかを LLM に伝える決定論的な MCP サーバー。カスケード判定、爆発範囲、インタラクションタイムライン、ピクセルパーフェクトな監査。 · GitHub
説明: AI コーディング エージェントの目: どのファイルのどの行でどの CSS ルールが勝つのか、そしてその理由を LLM に伝える決定論的な MCP サーバー。カスケード判定、爆発範囲、インタラクションタイムライン、ピクセルパーフェクトな監査。 - mi60dev/visionaire-engine

記事本文:
GitHub - mi60dev/visionaire-engine: AI コーディング エージェントの目: どのファイルのどの行でどの CSS ルールが勝つのか、そしてその理由を LLM に伝える決定論的な MCP サーバー。カスケード判定、爆発範囲、インタラクションタイムライン、ピクセルパーフェクトな監査。 · GitHub
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

t
{{ メッセージ }}
mi60dev
/
ビジョネアエンジン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
40 コミット 40 コミット .github .github bench bench docs docs harness harness scripts scripts src src test test .dockerignore .dockerignore .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile ライセンス ライセンス通知 通知 README.md README.md Glama.json Glama.json hero-2.jpeg hero-2.jpeg hero.png hero.png package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
どのルール、どのファイル、どの行が勝つのか、そしてなぜそれが勝つのか。
問題: 何かがおかしいようです。それをスクリーンショットして説明すると、LLM が間違っていると推測したので、再度説明します。
解決策: Visionaire はライブ ページを読み取り、LLM に正確なルール、ファイル、行を渡します。正しい修正、まずは試してください。
2px マージンのバグを説明するために段落を書く必要はありませんが、今はその必要がありません。説明は少なく、修正が多くなります: 開発者、バイブコーダー、ループ内で LLM を使用してサイトのデザイン変更を出荷する人向けに構築されています。
ステータス: v0.7 — 28 のツール、435 のテスト (実際の Chrome で 252 のユニット + 183 のエンドツーエンド)、24 件のシードバグ ベンチマーク ( npm run bench )、wordpress.org に対してライブで検証済み。
v0.7 — 検証レイヤー:assert_visual (17 タイプのアサーション文法 — 測定されたピクセル、問題のある UID、および再実行可能な名前付きスイートによる合否判定)、visual_diff (モックアップまたは記録されたベースラインとのピクセルの差、要素 UID にマッピングされた発散領域)、impact_preview (共有セレクターを編集する前のブラスト半径 + サンドボックス化されたドライラン)、診断 (「なぜこれが壊れているのか」のランク付け)手段を持った犯人

証拠を取得）、応答性のスイープ（1 回の呼び出し→ビューポートごとの評決マトリックス）、capture_proof（評決デルタを含む前後の証拠バンドル）;クロード コードとカーソルの編集後検証ハーネス ( npxvisionaire-engine init-harness )。 style_diff { Capture_pixels } ベースライン; check_alignment は非推奨になり、assert_visual が優先されます
v0.6 — ピクセルパーフェクトパック: check_alignment (グループアライメント / ギャップリズム / グリッド / ピクセルスナップ監査) および pick_color (実際のペイントされたピクセルのサンプリング + WCAG コントラスト判定)
v0.5: inject_css — ライブ修正ループ (ページ上で修正を試し、何が変更されたかを確認し、収束し、ソースを一度書き込みます)。古いスタイルシートのハードリロードについては、{ bypassCache } に移動します。 Explain_styles に関する blast-radius + スコープ付き修正レポート (すべてのボタンではなく、ボタンを変更します)
v0.4 (フィールド レポート アイテム): UI をある状態に駆動し、それを検査するために対話します。サブピクセルグリフ/テキストインクのセンタリング用のmeasure_element。避難ハッチを評価する。 annotated_screenshot の要素スコープのクロップ/ズーム。 find_elements の match:"any" /visibleOnly:false ;ゼロ設定のコールドスタート Chrome ディスカバリー
v0.3: 時間次元 — イベントリスナー属性、アニメーション診断、ソース属性のインタラクションタイムライン
信頼できないページに対する強化: プロンプトインジェクションサニタイズ、フェイルファストウォッチドッグ、ダイアログの自動終了
今日、LLM に視覚的なバグの修正を依頼すると、2 つの不完全な画像のうちの 1 つが得られます。1 つはコードへのリンクがないピクセル、もう 1 つは画面上で実際に何が勝っているのか正確にレンダリングされていないコードです。既存のブラウザ MCP は、特にデザイン作業においてこれをさらに悪化させます。これらの MCP は、意図的にすべてのスタイルを削除したアクセシビリティ スナップショットを提供します。欠けているのは説明と帰属です:
このプロパティのカスケードでどの CSS ルールが勝ちましたか?なぜ他の CSS ルールが負けたのでしょうか?
どのファイルのどの行に w が実行されますか

inner live in — またはどの Elementor ウィジェット コントロールか、どのカスタマイザー エントリでしょうか?
この要素が表示されない、位置がずれている、またはサイズが間違っているのはなぜですか?
Visionaire は、内部 AI を使用せずにこれらの質問に答えます。すべてが Chrome DevTools プロトコルとクローズド ルールセットから決定論的に計算されます。あいまいな部分 (「ヒーローの下のボタンがオフに見える」を実際の要素に一致させる) は、呼び出し元の LLM に残ります。LLM は、uid キー付きのスナップショット、検索ツール、および注釈付きのスクリーンショットを取得して、それを安価に実行します。
なぜ色 = rgb(255, 255, 255):
勝者 [class*=wp-block] .wp-block-button__link { color: var(--wp--custom--button--color--text) } spec(0,2,0)
→ テーマ/wporg-parent-2021/build/style.css:499 [行 |テーマ: wporg-parent-2021 — テーマ/wporg-parent-2021/build/style.css を編集]
失われた (特異性) :root :where(.wp-element-button, .wp-block-button__link) { color: #fff } spec(0,1,0)
→ global-styles-inline-css:2 [db-entity |グローバル スタイル — サイト エディター → スタイル (theme.json / wp_global_styles)]
失われた (原点) a:-webkit-any-link { color: -webkit-link } spec(0,1,1)
→ ユーザーエージェントスタイルシート
勝者、決定的な損失理由のある敗者、およびそれぞれに対する正直な編集ポインター。生成されたファイルへの無駄なパスの代わりに「サイト エディター → スタイル」などの WordPress を意識した回答が含まれます。
Node ≥ 20 および Chrome/Chromium がインストールされている必要があります。
最速のパス — クロード コードを使用して npm から直接登録します (クローンやビルドは不要):
クロード・MCP ビジョネアを追加 -- npx -y ビジョネア・エンジン
または、クローンから実行します (開発または固定されたローカル ビルドの場合)。
git clone https://github.com/mi60dev/visionaire-engine && cdvisionaire-engine
npm install && npm run build
クロード・MCP ビジョネアを追加 -- ノード " $PWD /dist/index.js "
GitHub Copilot、Cursor、Claude Desktop、Google Antigravity、または別のクライアントを使用していますか?見る

docs/clients.md には、それぞれの設定のコピー＆ペーストと、Linux/WSL/Docker のブラウザーインストールのヘルプが含まれています。
プロジェクトのルート ディレクトリから実行します。Visionaire は、エージェントが実行中のサイトとそのソースの両方をディスク上に持っている場合に最大限の効果を発揮し、その 2 つを相互参照できます。検索する前に確認してください。セレクターを推測するのではなく、page_snapshot を撮って (またはソースを読んで)、実際の要素名を取得してください。セレクターが何も一致しない場合、エラーはページ上で最も近い実際の ID/クラスを提案します。
connect { url: "https://your-site.com" } — Chrome を起動します (または、ログインしている実際のブラウザに接続するには {browserUrl: "http://127.0.0.1:9222" })
page_snapshot {} — 表示されているものの uid キー付きの調査。作成されたセレクターではなく、 uid によって要素をターゲットにします
Explain_styles { uid: "e17", property: "margin-bottom" } — file:line によるカスケード判定
npm run デモ # バンドルされたフィクスチャ
npm 実行デモ -- https://wordpress.org --selector " a.wp-block-button__link "
28 のツール
セッションとグラウンディング — 推測せずに接続し、適切な要素を見つけます。
説明 - 「理由」と領収書。
ピクセルレベルのチェック — v0.6 の新機能。
インタラクションと時間 — スナップショットだけでなく状態も。
検証と証明 — v0.7 の新機能。
検証ループ (CSS ガスライティングを停止する)
エージェントは、レンダリングされたピクセルを見ることなく、スタイルシートを編集し、自身の差分を読み取り、「カードの高さが同じになった」と宣言します。 Visionaire は、エージェントにレンダリングされた真実に対する決定的な目を与えます。ループ:
共有クラスの爆発半径をプレビュー → Impact_preview 。
クレームをアサート→assert_visual (または名前付き suite_id を再実行)。 PASS/FAIL + 実際に測定されたピクセル + 問題のある要素の UID が得られます。
FAIL を診断します。→ 診断すると、ランク付けされた犯人が証拠とともに返されます。
スイープ応答 → 応答性の高いスイープはビューポートごとに返します

評決マトリックス。
証明する → スクリーンショット前後の Capture_proof バンドル + 判定デルタ。
実際の出力 - 修正の前後で同じスイート:
{
"評決" : " 不合格 " 、
"summary" : " 1 つのアサーション: 0 合格、1 失敗 — スイート「カード」として登録されました ({ \" スイート ID \" : \" カード \" } だけで再実行します) " ,
「結果」: [
{ "type" : "equal_height " 、 "verdict" : " FAIL " 、 "id" : "cards-equal " 、
"測定" : { "値" : [ 412 , 388 ]、 "単位" : " px " 、 "デルタ" : 24 、 "tolerance_px" : 1 }、
"offending_uids" : [ " e1 " , " e2 " ] }
]、
"truncated" : false 、"suite_id" : "cards"
}
…CSS を修正し、 { "suite_id": "cards" } だけを指定して再実行します。
{
"評決" : " 合格 " 、
"summary" : " 1 アサーション: 1 合格、0 失敗 " ,
「結果」:[
{ "type" : "equal_height" 、"verdict" : "PASS" 、"id" : "cards-equal" 、
"測定" : { "値" : [ 412 , 412 ]、 "単位" : " px " 、 "デルタ" : 0 、 "tolerance_px" : 1 } }
]、
"truncated" : false 、"suite_id" : "cards"
}
プロジェクト ルートから npxvisionaire-engine init-harness を実行して、含まれているクロード コード フック (またはカーソル ルール) を接続します。これにより、記録上の検証パスがなければエージェントが「修正された」と主張してターンを物理的に終了できなくなります。マーカー、フック、および停止ゲートの仕組み: docs/harness.md 。
docs/clients.md — Claude、Copilot、Cursor、Antigravity、およびその他の MCP クライアントにインストールします
docs/tools.md — 実際の例を含むツールごとのリファレンス
docs/harness.md — 編集後の検証ハーネス (フック、マーカー、init-harness )
docs/architecture.md — 決定論的パイプラインの仕組み
docs/wordpress.md — WordPress オリジン解決ガイド
docs/development.md — 構築、テスト、拡張
内部 LLM なし - 決定的、キャッシュ可能、テスト可能、ホストに依存しない。
ファジーグラウンディングは呼び出し元の LLM に属します。私たちはそれを安くします。
現職の仲間を補完する

wser MCP (同じ uid イディオム)、競合しないでください。
すべての属性の正直さラダー: 行 > ファイル > db-entity > コンポーネント > 生成された > 不明 。
トークン予算に基づいた出力 — ドシエは 300 ～ 800 トークンであり、決してダンプではありません。
Visionaire は任意の、信頼できないページを指しているため、ページのコンテンツを敵対的なものとして扱います。
即時注射ディフェンス。ページ派生文字列 (要素テキスト、クラス名、ID、属性値) は、ツール出力に入る単一のチョーク ポイントでサニタイズされます。つまり、1 行に折りたたまれ、制御文字と双方向オーバーライド文字が削除され、長さに上限が設定されます。ページは、「システム メッセージ」としてフォーマットされた命令形式のテキストを呼び出し元の LLM にこっそり持ち込むことはできません。このようなコンテンツは、不活性で引用符で切り詰められたフラグメントとしてのみ表示されます。
フェイルファストでハングしません。すべてのツール呼び出しはウォッチドッグ内にラップされます (デフォルトは 60 秒、オーバーライドするには VISIONAIRE_TOOL_TIMEOUT_MS、pick_element / Record_interaction は宣言された待機とスラックを取得します)。ウェッジブラウザは、クライアントをブロックするのではなく、再度接続するように指示する対処可能なエラーを返します。
デッドロックダイアログはありません。ページのalert() /confirm()/prompt()呼び出しは自動的に終了します。そうしないと、すべてのevaluate-family CDP呼び出しが無期限にブロックされてしまいます。
Visionaire は、ページで作成されたコードを命令として実行することはありません。読み取りと属性のみを実行します。呼び出し元の LLM は、

[切り捨てられた]

## Original Extract

Eyes for AI coding agents: deterministic MCP server that tells the LLM which CSS rule wins, in which file, on which line — and why. Cascade verdicts, blast radius, interaction timelines, pixel-perfect audits. - mi60dev/visionaire-engine

GitHub - mi60dev/visionaire-engine: Eyes for AI coding agents: deterministic MCP server that tells the LLM which CSS rule wins, in which file, on which line — and why. Cascade verdicts, blast radius, interaction timelines, pixel-perfect audits. · GitHub
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
mi60dev
/
visionaire-engine
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
40 Commits 40 Commits .github .github bench bench docs docs harness harness scripts scripts src src test test .dockerignore .dockerignore .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE NOTICE NOTICE README.md README.md glama.json glama.json hero-2.jpeg hero-2.jpeg hero.png hero.png package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts View all files Repository files navigation
Which rule, which file, which line — and why it wins.
The problem: Something looks off. You screenshot it, explain it, the LLM guesses wrong, you re-explain.
The solution: Visionaire reads the live page and hands the LLM the exact rule, file, and line. Right fix, first try.
You shouldn't have to write a paragraph to explain a 2px margin bug — and now you don't. Less explaining, more fixing: built for developers, vibe coders, and anyone shipping site design changes with an LLM in the loop.
Status: v0.7 — 28 tools, 435 tests (252 unit + 183 end-to-end on real Chrome), a 24-case seeded-bug benchmark ( npm run bench ), verified live against wordpress.org.
v0.7 — the verification layer: assert_visual (a 17-type assertion grammar — PASS/FAIL verdicts with measured pixels, offending uids, and re-runnable named suites), visual_diff (pixel diff vs a mockup or recorded baseline, divergent regions mapped to element uids), impact_preview (blast radius + sandboxed dry-run before editing a shared selector), diagnose (ranked "why is this broken" culprits with measured evidence), responsive_sweep (one call → per-viewport verdict matrix), capture_proof (before/after evidence bundles with a verdict delta); the verify-after-edit harness for Claude Code and Cursor ( npx visionaire-engine init-harness ); style_diff { capture_pixels } baselines; check_alignment deprecated in favor of assert_visual
v0.6 — the pixel-perfect pack: check_alignment (group alignment / gap-rhythm / grid / pixel-snap audit) and pick_color (actual painted-pixel sampling + WCAG contrast verdicts)
v0.5: inject_css — the live fix loop (trial a fix on the page, see what changed, converge, write source once); navigate { bypassCache } for stale-stylesheet hard reloads; blast-radius + scoped-fix reporting on explain_styles (change the button, not all buttons)
v0.4 (field-report items): interact to drive the UI into a state and inspect it; measure_element for sub-pixel glyph/text-ink centering; an evaluate escape hatch; element-scoped crops/zoom on annotated_screenshot ; match:"any" / visibleOnly:false on find_elements ; zero-config cold-start Chrome discovery
v0.3: the time dimension — event-listener attribution, animation diagnosis, source-attributed interaction timelines
Hardened for untrusted pages: prompt-injection sanitization, fail-fast watchdog, dialog auto-dismiss
Ask an LLM to fix a visual bug today and it gets one of two incomplete pictures: pixels, with no link back to code, or code, with no rendering truth about what's actually winning on screen. Existing browser MCPs make this worse for design work specifically — they ship accessibility snapshots that deliberately strip out all styling. What's missing is explanation and attribution :
Which CSS rule wins the cascade for this property, and why did the others lose?
Which file, which line does the winner live in — or which Elementor widget control, or which Customizer entry?
Why is this element invisible, misaligned, or the wrong size?
Visionaire answers those questions with zero AI inside — everything is computed deterministically from the Chrome DevTools Protocol plus closed rulesets. The fuzzy part (matching "the button under the hero looks off" to an actual element) stays with the calling LLM, which gets uid-keyed snapshots, search tools, and annotated screenshots to do that cheaply.
why color = rgb(255, 255, 255):
WINNER [class*=wp-block] .wp-block-button__link { color: var(--wp--custom--button--color--text) } spec(0,2,0)
→ themes/wporg-parent-2021/build/style.css:499 [line | theme: wporg-parent-2021 — edit themes/wporg-parent-2021/build/style.css]
lost (specificity) :root :where(.wp-element-button, .wp-block-button__link) { color: #fff } spec(0,1,0)
→ global-styles-inline-css:2 [db-entity | Global Styles — Site Editor → Styles (theme.json / wp_global_styles)]
lost (origin) a:-webkit-any-link { color: -webkit-link } spec(0,1,1)
→ user-agent stylesheet
Winner, losers with the decisive loss reason, and an honest edit pointer for each — including WordPress-aware answers like "Site Editor → Styles" instead of a useless path to a generated file.
Requires Node ≥ 20 and Chrome/Chromium installed.
Fastest path — register straight from npm with Claude Code (no clone, no build):
claude mcp add visionaire -- npx -y visionaire-engine
Or run from a clone (for development or a pinned local build):
git clone https://github.com/mi60dev/visionaire-engine && cd visionaire-engine
npm install && npm run build
claude mcp add visionaire -- node " $PWD /dist/index.js "
Using GitHub Copilot, Cursor, Claude Desktop, Google Antigravity , or another client? See docs/clients.md for a copy-paste config for each, plus browser-install help for Linux/WSL/Docker.
Run it from your project's root directory — Visionaire is at its best when the agent has both the running site and its source on disk, so it can cross-reference the two. Ground before you search: take a page_snapshot (or read the source) to get real element names instead of guessing selectors. If a selector matches nothing, the error suggests the closest real ids/classes on the page.
connect { url: "https://your-site.com" } — launches Chrome (or { browserUrl: "http://127.0.0.1:9222" } to attach to your real, logged-in browser)
page_snapshot {} — a uid-keyed census of what's visible; target elements by their uid , not invented selectors
explain_styles { uid: "e17", property: "margin-bottom" } — cascade verdict with file:line
npm run demo # bundled fixture
npm run demo -- https://wordpress.org --selector " a.wp-block-button__link "
The 28 tools
Session & grounding — get connected and find the right element without guessing.
Explanation — the "why," with a receipt.
Pixel-level checks — new in v0.6.
Interaction & time — states, not just snapshots.
Verification & proof — new in v0.7.
The verify loop (stop the CSS gaslighting)
An agent edits a stylesheet, reads its own diff, and declares "now the cards are equal height" — without ever seeing a rendered pixel. Visionaire gives your agent deterministic eyes on rendered truth. The loop:
Preview shared-class blast radius → impact_preview .
Assert your claim → assert_visual (or re-run a named suite_id ). You get PASS/FAIL + the actual measured pixels + the offending element uids.
Diagnose any FAIL → diagnose returns the ranked culprit with evidence.
Sweep responsive → responsive_sweep returns a per-viewport verdict matrix.
Prove it → capture_proof bundles before/after screenshots + verdict delta.
Real output — the same suite before and after a fix:
{
"verdict" : " FAIL " ,
"summary" : " 1 assertion: 0 PASS, 1 FAIL — registered as suite 'cards' (re-run with just { \" suite_id \" : \" cards \" }) " ,
"results" : [
{ "type" : " equal_height " , "verdict" : " FAIL " , "id" : " cards-equal " ,
"measured" : { "values" : [ 412 , 388 ], "unit" : " px " , "delta" : 24 , "tolerance_px" : 1 },
"offending_uids" : [ " e1 " , " e2 " ] }
],
"truncated" : false , "suite_id" : " cards "
}
…fix the CSS, re-run with just { "suite_id": "cards" } :
{
"verdict" : " PASS " ,
"summary" : " 1 assertion: 1 PASS, 0 FAIL " ,
"results" : [
{ "type" : " equal_height " , "verdict" : " PASS " , "id" : " cards-equal " ,
"measured" : { "values" : [ 412 , 412 ], "unit" : " px " , "delta" : 0 , "tolerance_px" : 1 } }
],
"truncated" : false , "suite_id" : " cards "
}
Run npx visionaire-engine init-harness from your project root to wire the included Claude Code hooks (or Cursor rule) so the agent physically cannot end a turn claiming "it's fixed" without a verification pass on record. How the markers, hooks, and Stop gate work: docs/harness.md .
docs/clients.md — install in Claude, Copilot, Cursor, Antigravity, and other MCP clients
docs/tools.md — tool-by-tool reference with real examples
docs/harness.md — the verify-after-edit harness (hooks, markers, init-harness )
docs/architecture.md — how the deterministic pipeline works
docs/wordpress.md — WordPress origin resolution guide
docs/development.md — building, testing, extending
No internal LLM — deterministic, cacheable, testable, host-agnostic.
Fuzzy grounding belongs to the calling LLM ; we make it cheap.
Complement the incumbent browser MCPs (same uid idiom), don't compete.
Honesty ladder on every attribution: line > file > db-entity > component > generated > unknown .
Token-budgeted output — a dossier is 300–800 tokens, never a dump.
Visionaire is pointed at arbitrary, untrusted pages, so it treats page content as hostile:
Prompt-injection defense. Page-derived strings (element text, class names, ids, attribute values) are sanitized at the single choke point where they enter tool output — collapsed to one line, stripped of control and bidirectional-override characters, and length-capped. A page cannot smuggle instruction-shaped text formatted as a "system message" toward the calling LLM; such content can only appear as an inert, quoted, truncated fragment.
Fail-fast, never hang. Every tool call is wrapped in a watchdog (default 60s, VISIONAIRE_TOOL_TIMEOUT_MS to override; pick_element / record_interaction get their declared wait plus slack). A wedged browser returns an actionable error telling you to connect again, instead of blocking the client.
No dead-locking dialogs. Page alert() / confirm() / prompt() calls are auto-dismissed — otherwise they would block every evaluate-family CDP call indefinitely.
Visionaire never executes page-authored code as instructions; it only reads and attributes. The calling LLM should

[truncated]
