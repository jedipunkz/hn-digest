---
source: "https://github.com/gridproof/gridproof"
hn_url: "https://news.ycombinator.com/item?id=49298638"
title: "Show HN: GridProof – spacing/grid QA for AI-generated UIs, in the agent loop"
article_title: "GitHub - gridproof/gridproof: Spacing & grid QA for AI-generated UIs, in the agent loop. MCP server that audits computed geometry and returns fix hints. · GitHub"
author: "seomarlboro"
captured_at: "2026-08-14T14:09:15Z"
capture_tool: "hn-digest"
hn_id: 49298638
score: 1
comments: 0
posted_at: "2026-08-14T13:48:32Z"
tags:
  - hacker-news
  - translated
---

# Show HN: GridProof – spacing/grid QA for AI-generated UIs, in the agent loop

- HN: [49298638](https://news.ycombinator.com/item?id=49298638)
- Source: [github.com](https://github.com/gridproof/gridproof)
- Score: 1
- Comments: 0
- Posted: 2026-08-14T13:48:32Z

## Translation

タイトル: HN の表示: GridProof – AI 生成の UI の間隔/グリッド QA (エージェント ループ内)
記事のタイトル: GitHub - グリッドプルーフ/グリッドプルーフ: エージェント ループでの AI 生成 UI のスペーシングとグリッド QA。計算されたジオメトリを監査し、修正ヒントを返す MCP サーバー。 · GitHub
説明: エージェント ループにおける、AI 生成の UI の間隔とグリッドの QA。計算されたジオメトリを監査し、修正ヒントを返す MCP サーバー。 - グリッドプルーフ/グリッドプルーフ

記事本文:
GitHub - グリッドプルーフ/グリッドプルーフ: エージェント ループでの、AI 生成の UI のスペースとグリッドの QA。計算されたジオメトリを監査し、修正ヒントを返す MCP サーバー。 · GitHub
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
グリッドプルーフ
/
グリッドプルーフ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
30 コミット 30 コミット キャリブレーション キャリブレーション デモ デモ docs docs fixtures fixtures scripts scripts src src test test .gitignore .gitignore LICENSE

ライセンス README.md README.md SKILL.md SKILL.md Gridproof-spec.md Gridproof-spec.md package-lock.json package-lock.json package.json package.json server.json server.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
UI の GridProof — エージェント ループ内の自動間隔とグリッド QA。
GridProof は、
Playwright でフロントエンドを実行し、すべての計算されたジオメトリを測定します。
要素をスペース/トークン ルール セットと照合してチェックし、
構造化された修正レポート — コーディング エージェントがループ自体を閉じることができるようにするため、生成、
監査、修正、再監査。
AI コーディング エージェントは UI の生成は得意ですが、UI をグリッド上に維持するのは苦手です。
py-3 の代わりに py-[13px] 、3 つの異なるギャップを持つ兄弟カード、
アイコンは 24 ピクセルの隣の 17 ピクセルにあります。何も壊れていないので、出荷されます -
エージェント ループでは何もチェックしないためです。 GridProof はそのチェックです。
エージェントは UI → gp_audit(url) → 修正ヒントを含む JSON 違反を生成します
→ エージェントがソースを編集 → gp_audit(url) → クリーンレポート = 完了
サーバーがソース ファイルにアクセスすることはありません。レンダリングされたページを測定し、
ポイント。エージェント (コードベースを開いている) が編集を行います。
# 1 回限り: Playwright が使用する Chromium ビルドをインストールします (~150MB)
npx playwright インストール クロム
クロードコードに登録する
クロード mcp グリッドプルーフを追加 -- npx -y グリッドプルーフ
ローカルチェックアウトから:
npm install && npm run build
クロード mcp グリッドプルーフを追加 -- ノード /absolute/path/to/gridproof/dist/index.js
3つの使い方
1. gp_audit — エージェント ループ内。エージェントはこの MCP ツールを呼び出します
実行中の開発サーバーに対して直接実行し、構造化された JSON を取得します。
(違反 + 修正のヒント) に基づいて行動します。
2. gp_report — HTML レポートも作成する MCP ツール。と同じ入力
gp_audit に加えて、自己

含まれる共有可能な HTML ファイルをディスクに保存します。
3. npx Gridproof --report <url> — ワンショット CLI。 MCP クライアントは必要ありません。
簡単な手動チェックやスクリプト作成に役立ちます。
npx Gridproof --report http://localhost:5173
# ./gridproof-report.html を書き込み、そのパスを出力します
npx グリッドプルーフ --レポート http://localhost:5173 --out ./qa/report.html --viewport 375x812
ルール
4つのルール。デフォルトではすべてのレポートが警告します - 何もブロックせず、何もブロックしません
終了コードのセマンティクス。提案するが、禁止しないでください。唯一の例外はタップです
これはスタイルに関する意見ではなく、アクセシビリティ フロアであるため、エラーが発生します。
追い風のページと非追い風のページ
GridProof は Tailwind プロジェクト用に構築されています。そこに 4 つのルールがすべて適用されます。
spacing-scale 、 arbitrary-value 、および gap-consistency の理由から、
Tailwind の間隔スケールとユーティリティ クラス。
Tailwind として検出されないページでは、自動的に次のページにフォールバックします。
アクセシビリティのみのチェック: canonical-size は引き続き実行されます (ターゲットをタップ、アイコン
サイズ)、Tailwind 固有の 3 つのルールはスキップされ、レポートにはそのように記載されています。
黙って過小報告するのではなく、明示的に報告する。これを強制的に実行できます
設定で assignTailwind: false。
プロジェクト ルートにあるオプションの Gridproof.config.json (すべてのフィールドはオプションです。
デフォルトを示します):
{
"ベースユニット" : 4 、
"allowedValues" : [ 1 , 2 ],
"canonicalSizes" : [ 12 , 14 , 16 , 20 , 24 , 32 , 40 , 48 ],
"minTapTarget" : 44 、
"タップターゲットブレークポイント" : 768 、
"アイコン許容値" : 2 、
"assumeTailwind" : " 自動 " 、
「ルール」: {
"spacing-scale" : " warn " ,
"任意の値" : " 警告 " ,
"gap-consistency" : " warn " ,
"canonical-size" : " エラー "
}、
「抑制」 : [
{ "selector" : " .hero-art * " , "rules" : [ " spacing-scale " ] },
{ "値" : " 13px " 、 "理由" : "光学補正、ロゴのロックアップ " }
]
}
インライン抑制: data-gp-ignore (すべてのルール) または
data-gp-ignore="間隔スケール

任意の要素の「gap-consistency」はその要素をスキップします
それらのルールのサブツリー。抑制された所見はカウントされますが、リストされることはありません。
意図的に行わないこと
コンピュータービジョン/スクリーンショット分析はありません。計算されたジオメトリを読み取ります。
ピクセルではありません。スクリーンショットは HTML レポートに添付されますが、分析はされません。
CI ランナーはありません。これはエージェントのインループ ツールであり、マージ ゲートではありません - いいえ
終了コードがあれば、ビルドに失敗することはありません。
ソース編集はありません。サーバーは測定して提案します。エージェント（
コードベースがあります) が編集を行います。
認証も SaaS も請求もありません。これはローカル MCP サーバーと CLI です。
まだ (v2 候補、未実装): 列グリッド クラスタリング、
ブレークポイント間のアラインメント ドリフト、Figma トークンのインポート。
Playwright はターゲット ページをヘッドレスでレンダリングし、単一のページ内スクリプトがターゲット ページを実行します。
DOM を使用して、計算されたジオメトリ (マージン、パディング、ギャップ、四角形) を収集します。
ルール エンジンは各値を設定と照合してチェックし、違反を発行します。
セレクター、実際の値/期待値、および修正ヒント。おおよそに対して調整されています
60 の現実世界のサイトで誤検知を低く抑える - サブピクセルの四捨五入
許容値、許容値リスト、および重大度のデフォルトはすべてそこから導き出されます。
推測ではなく校正です。
npmインストール
npm run build # tsc → dist/
npm test # vitest (ユニット + Playwright 統合)
npm run dev # TypeScript (tsx) からサーバーを実行します
ライセンス
エージェント ループでの、AI 生成の UI の間隔とグリッドの QA。計算されたジオメトリを監査し、修正ヒントを返す MCP サーバー。
www.npmjs.com/package/gridproof トピック
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Spacing & grid QA for AI-generated UIs, in the agent loop. MCP server that audits computed geometry and returns fix hints. - gridproof/gridproof

GitHub - gridproof/gridproof: Spacing & grid QA for AI-generated UIs, in the agent loop. MCP server that audits computed geometry and returns fix hints. · GitHub
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
gridproof
/
gridproof
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
30 Commits 30 Commits calibration calibration demo demo docs docs fixtures fixtures scripts scripts src src test test .gitignore .gitignore LICENSE LICENSE README.md README.md SKILL.md SKILL.md gridproof-spec.md gridproof-spec.md package-lock.json package-lock.json package.json package.json server.json server.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts View all files Repository files navigation
GridProof your UI — automated spacing & grid QA in the agent loop.
GridProof is an MCP server that renders your
running frontend with Playwright, measures the computed geometry of every
element, checks it against a spacing/token rule set, and hands back a
structured fix report — so a coding agent can close the loop itself: generate,
audit, fix, re-audit.
AI coding agents are good at generating UI and bad at keeping it on a grid:
py-[13px] instead of py-3 , sibling cards with three different gaps,
icons at 17px next to 24px. None of it breaks anything, so it ships —
because nothing in the agent loop checks for it. GridProof is that check.
agent generates UI → gp_audit(url) → JSON violations with fix hints
→ agent edits source → gp_audit(url) → clean report = done
The server never touches your source files. It measures a rendered page and
points; the agent (which has your codebase open) makes the edit.
# One-time: install the Chromium build Playwright uses (~150MB)
npx playwright install chromium
Register in Claude Code
claude mcp add gridproof -- npx -y gridproof
From a local checkout:
npm install && npm run build
claude mcp add gridproof -- node /absolute/path/to/gridproof/dist/index.js
Three ways to use it
1. gp_audit — inside the agent loop. The agent calls this MCP tool
directly against your running dev server and gets back structured JSON
(violations + fix hints) to act on.
2. gp_report — MCP tool that also writes an HTML report. Same inputs as
gp_audit , plus it writes a self-contained, shareable HTML file to disk.
3. npx gridproof --report <url> — one-shot CLI. No MCP client needed;
useful for a quick manual check or scripting.
npx gridproof --report http://localhost:5173
# writes ./gridproof-report.html, prints its path
npx gridproof --report http://localhost:5173 --out ./qa/report.html --viewport 375x812
The rules
Four rules. All report warn by default — nothing blocks, nothing has
exit-code semantics. Suggest, don't forbid ; the one exception is tap
targets, which error because it's an accessibility floor, not a style opinion.
Tailwind, and non-Tailwind pages
GridProof is built for Tailwind projects — that's where all four rules apply,
since spacing-scale , arbitrary-value , and gap-consistency reason about
Tailwind's spacing scale and utility classes.
On a page it doesn't detect as Tailwind, it auto-falls-back to
accessibility-only checks: canonical-size still runs (tap targets, icon
sizes), the three Tailwind-specific rules are skipped, and the report says so
explicitly rather than silently under-reporting. You can force this with
assumeTailwind: false in config.
Optional gridproof.config.json at your project root (all fields optional;
defaults shown):
{
"baseUnit" : 4 ,
"allowedValues" : [ 1 , 2 ],
"canonicalSizes" : [ 12 , 14 , 16 , 20 , 24 , 32 , 40 , 48 ],
"minTapTarget" : 44 ,
"tapTargetBreakpoint" : 768 ,
"iconTolerance" : 2 ,
"assumeTailwind" : " auto " ,
"rules" : {
"spacing-scale" : " warn " ,
"arbitrary-value" : " warn " ,
"gap-consistency" : " warn " ,
"canonical-size" : " error "
},
"suppress" : [
{ "selector" : " .hero-art * " , "rules" : [ " spacing-scale " ] },
{ "value" : " 13px " , "reason" : " optical correction, logo lockup " }
]
}
Inline suppression: data-gp-ignore (all rules) or
data-gp-ignore="spacing-scale gap-consistency" on any element skips its
subtree for those rules. Suppressed findings are counted, never listed.
What it deliberately does NOT do
No computer vision / screenshot analysis. It reads computed geometry,
not pixels. A screenshot is attached to the HTML report, not analyzed.
No CI runner. It's an in-loop tool for an agent, not a merge gate — no
exit codes, nothing fails a build.
No source editing. The server measures and suggests; the agent (which
has your codebase) makes the edits.
No auth, no SaaS, no billing. It's a local MCP server and a CLI.
Not yet (v2 candidates, not implemented): column-grid clustering,
cross-breakpoint alignment drift, Figma token import.
Playwright renders the target page headless, a single in-page script walks the
DOM and collects computed geometry (margins, padding, gap, rects), and the
rule engine checks each value against your config and emits violations with
selectors, actual/expected values, and fix hints. It's tuned against roughly
60 real-world sites to keep false positives low — a subpixel rounding
tolerance, an allowed-values list, and severity defaults all come out of that
calibration, not guesswork.
npm install
npm run build # tsc → dist/
npm test # vitest (unit + Playwright integration)
npm run dev # run the server from TypeScript (tsx)
License
Spacing & grid QA for AI-generated UIs, in the agent loop. MCP server that audits computed geometry and returns fix hints.
www.npmjs.com/package/gridproof Topics
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
