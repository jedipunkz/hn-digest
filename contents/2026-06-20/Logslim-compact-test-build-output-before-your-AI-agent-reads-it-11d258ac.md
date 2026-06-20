---
source: "https://github.com/P156HAM/logslim"
hn_url: "https://news.ycombinator.com/item?id=48607333"
title: "Logslim – compact test/build output before your AI agent reads it"
article_title: "GitHub - P156HAM/logslim: Structured CI failure summaries on PRs + compact test/build logs for AI agents. GitHub Action, CLI, MCP. · GitHub"
author: "P156HAM"
captured_at: "2026-06-20T10:05:24Z"
capture_tool: "hn-digest"
hn_id: 48607333
score: 1
comments: 0
posted_at: "2026-06-20T08:13:20Z"
tags:
  - hacker-news
  - translated
---

# Logslim – compact test/build output before your AI agent reads it

- HN: [48607333](https://news.ycombinator.com/item?id=48607333)
- Source: [github.com](https://github.com/P156HAM/logslim)
- Score: 1
- Comments: 0
- Posted: 2026-06-20T08:13:20Z

## Translation

タイトル: Logslim – AI エージェントが読み取る前のコンパクトなテスト/ビルド出力
記事のタイトル: GitHub - P156HAM/logslim: PR に関する構造化された CI 障害の概要 + AI エージェントのコンパクトなテスト/ビルド ログ。 GitHub アクション、CLI、MCP。 · GitHub
説明: PR に関する構造化された CI 障害の概要 + AI エージェントのコンパクトなテスト/ビルド ログ。 GitHub アクション、CLI、MCP。 - P156HAM/ログスリム

記事本文:
GitHub - P156HAM/logslim: PR に関する構造化された CI 障害の概要 + AI エージェントのコンパクトなテスト/ビルド ログ。 GitHub アクション、CLI、MCP。 · GitHub
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
P156ハム
/
ログスリム
公共
通知
通知設定を変更するにはサインインする必要があります
追加

最後のナビゲーションオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
4 コミット 4 コミット .github .github アクション アクション エラー エラー スクリプト スクリプト src src test test .gitignore .gitignore .npmignore .npmignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
CIは失敗しましたか？ 400 行のアクション ログではなく、5 行の PR 概要を取得します。
エージェントはテスト出力を読み取りますか?トークンの 80 ～ 95% をカットします。
テストまたはビルドが失敗すると、GitHub Actions ログをスクロールします。クロードコードまたはカーソルの実行時
npm test 、エージェントはすべてを読み取ります - 進行状況バー、120 個の同一の警告、40 フレーム
のnode_modules。 logslim は次の両方を修正します。
CI / 人間 — GitHub Action は PR に構造化されたエラーを投稿します (ファイル、行、修正ヒント)
エージェント / トークン — CLI + MCP は、LLM が読み取る前にノイズの多い出力を圧縮します (障害を最大 80 ～ 95% 節約)
npx logslim -- npm テスト
アカウントがありません。 APIキーがありません。 MIT オープンソース。
GitHub アクション — PR 失敗の概要
CI が失敗した場合は、レビュー担当者に詳細を調査させるのではなく、読みやすい概要をプル リクエストに投稿します。
アクションログを通じて。
権限:
内容：読む
プルリクエスト : 書き込み
ジョブ:
テスト:
実行: ubuntu-最新
手順:
- 使用:actions/checkout@v4
- 使用:actions/setup-node@v4
付き:
ノードバージョン: 20
- 実行: npm ci
- 名前 : テストの実行
ID : テスト
実行：npm test 2>&1 |ティーテスト出力.log
エラー時続行 : true
- 名前 : 失敗後の概要
if :steps.test.outcome == '失敗' && github.event_name == 'pull_request'
使用: P156HAM/logslim/action@v0.3.0
付き:
ログファイル: test-output.log
終了コード: 1
github トークン: ${{ Secrets.GITHUB_TOKEN }}
- 名前 : 失敗したジョブ
if : ステップ.テスト.アウトカム ==

「失敗」
実行: 1 番出口
PR に掲載される内容:
file:line およびメッセージに関する構造化エラー
既知のコード ( TS2339 、 ERESOLVE など) のヒントを修正
トークン/ログ削減統計 (エージェントも出力を読み取る場合に役立ちます)
アクションは、圧縮、エラー抽出、コード カードなど、CLI と同じエンジンを使用します。
人間が読める PR コメントとエージェントがすぐに使える JSON を 1 つのツールから取得できます。
git クローン https://github.com/P156HAM/logslim.git
cd logslim && npm install && npm run build && npm run Demon
またはクローンを作成しない場合:
npx logslim -- node -e " console.log('ok'); for(let i=0;i<30;i++)console.log('warn '+i); throw Error('fail') "
前 → 後
BEFORE — エージェントが今日読み取る内容 (ここでは約 18 行。実際の実行は 500 ～ 3000 行です):
src/utils.test.ts にパスします
console.warn 非推奨の prop id=1000
console.warn 非推奨の prop id=1001
console.warn 非推奨の prop id=1002
... (同じ警告×120)
失敗 src/checkout/cart.test.ts
予想: 89.10
受信: 99.00
カート.テスト.ts:48:27 で
node_modules/jest-circus/build/utils.js:298:28 で
node_modules/jest-circus/build/utils.js:231:10 で
node_modules/jest-circus/build/run.js:252:3 にあります。
... (さらに 15 個の node_modules フレーム)
テストスイート: 1 つ不合格、1 つ合格、合計 2 つ
AFTER — logslim がエージェントに与えるもの:
src/utils.test.ts にパスします
console.warn 非推奨の prop id=1000
console.warn 非推奨の prop id=1001
console.warn 非推奨の prop id=1002
(+5 同様の行が logslim によって省略されました)
失敗 src/checkout/cart.test.ts
予想: 89.10
受信: 99.00
カート.テスト.ts:48:27 で
node_modules/jest-circus/build/utils.js:298:28 で
… 3 つのベンダー/内部フレームが Logslim によって折りたたまれました
テストスイート: 1 つ不合格、1 つ合格、合計 2 つ
同じ失敗です。同じ修正です。 ~ 92% 少ないトークン。
npm テスト ──► logslim ──► エージェント / CI / あなた
│
§─ 1. ANSI カラーとスピナー ガベージを削除する
§─ 2. 重複排除の繰り返し

d 行 (スパム警告)
§─ 3. node_modules スタック フレームを折りたたむ
§─ 4. 類似した行をグループ化する (タイムスタンプ/ID はマスクされる)
§─ 5. 構造化エラーの抽出（ファイル、行、メッセージ）
§─ 6. 既知のコード (TS2339、ERESOLVE…) の修正カードを添付します。
━─ 7. オプションのトークンバジェット (中央をトリミングし、エラーを保持)
失敗モード (デフォルト): テストに合格 → 軽いクリーンアップのみ。テストが失敗→パイプラインがフルになる。
圧縮コストを支払うのは、実際に何かが壊れた場合のみです。
削除されたすべてのセクションはその場でマークされます ( (+ logslim によって省略された同様の 47 行) )
そのため、エージェントはデータが削除されたことを認識し、必要に応じて生のコマンドを再実行できます。
npm install -g logslim
# またはゼロインストール:
npx logslim -- npm テスト
ノード 18 以降が必要です。
logslim -- npm テスト
logslim -- python -m pytest -x
logslim -- npx tsc --noEmit
終了コードは保存されます。標準出力の出力は圧縮されます。標準エラー出力の統計。
npm テスト 2>&1 |ログスリム
npm テスト ; logslim --終了コード $? 2>&1 < full.log # 出力を保存した場合
JSON — エージェントとCI用
logslim --json -- npm テスト
{
"終了コード" : 1 、
"失敗" : true 、
"compacted" : " src/checkout/cart.test.ts が失敗しました \n ... " ,
「エラー」: [
{
"ファイル" : " cart.test.ts " ,
「行」：48、
"message" : " 期待値: 89.10、受信値: 99.00 " ,
"種類" : " アサーション "
}
]、
「コード」: [
{
"id" : " TS2339 " 、
"lang" : " typescript " ,
"意味" : " プロパティがタイプに存在しません " 、
"fix_steps" : [
「タイプミスをチェックしてください」、
" インターフェースを拡張する " 、
「オプションのチェーンを使用する」
】
}
]、
"統計" : {
"トークンイン" : 3296 、
"トークンアウト" : 252 、
「保存」: 0.92 、
"適用済み" : " フル "
}
}
エージェントは、数千行の散文ではなく、圧縮された + エラー + コードを読み取ります。
logslim --json --attach git,ci -- npm テスト
先頭に追加: ブランチ: feat/x |コミット: a3f2c1d | pr: #42 (GITHUB_* 環境変数から)。
旗
何をするのか
--モードの失敗
障害時のみコンパクトハード (d

デフォルト）
--モードフル
常にハードにコンパクトに
--モードライト
ANSI のみを削除し、積極的な重複排除は行わない
--json
構造化された出力 (上記を参照)
--git,ci を添付します
ブランチ/コミット/CI メタデータを先頭に追加する
-- 予算 2000
ハードトークンの上限。エラー + 先頭/末尾が生き残る
--終了コード N
終了コードがわかっている場合のパイプモードの場合
--コードなし
エラーコード修正カードをスキップする
--統計なし
標準エラー出力節約フッターを非表示にする
MCPサーバー(クロードコード/カーソル)
エージェントは圧縮をツールとして呼び出すことができます。手動でパイプする必要はありません。
プロジェクト .mcp.json または Claude デスクトップ構成:
{
"mcpサーバー": {
"ログスリム" : {
"コマンド" : " npx " ,
"args" : [ " -y " 、 " logslim-mcp " ]
}
}
}
ツール: Compact_output — 出力 (生のログ テキスト) とオプションの exit_code を渡します。
圧縮されたテキスト、抽出されたエラー、修正カード、および統計を返します。
npm run build && npm run mcp
エージェントにそれを使用するように伝えてください
CLAUDE.md 、 AGENTS.md 、または .cursor/rules に追加します。
詳細な出力を生成するテスト、ビルド、またはリンターを実行する場合:
- 推奨: ` logslim --mode Failure --json -- <command> `
- デバッグの前に、`compacted`、`errors`、および `codes` フィールドを読み取ります。
- 出力が省略された場合は、完全なログが必要な場合にのみ生のコマンドを再実行してください。
エラーコード修正カード
ログに既知のコードが含まれている場合、logslim はショート フィックス カード (~30 トークン) を添付します。
エージェントにドキュメントを推測させたり検索させたりするのではなく、
スクレイピングされたドキュメントではなく、手作業で厳選されたポケットリファレンス。 PR はコードを追加することを歓迎します。
ログの種類
ライン
トークン
保存されました
Jest (スパムを警告 + 1 回の失敗)
149→25
~3,300 → ~250
92%
Webpack ビルド (アセット ノイズ + 2 TS エラー)
548→55
～8,900 → ～1,000
88%
Pytest (25 個の同一の失敗)
356 → 153
~4,300 → ~1,500
64%
トークン数は推定されます (トークンあたり最大 4 文字)。請求ではなく、相対的な節約に適しています。
「logslim」から { コンパクト , プロセス } をインポートします。
const { text , stats } = Compact ( rawLog , { mode : "failure" , exitCode : 1 }

) ;
const result = プロセス ( rawLog , {
モード: "失敗" 、
終了コード : 1 、
アタッチ: [ "git" , "ci" ] ,
} ) ;
// result.text、result.errors、result.codes、result.stats
ログスリムを使用する場合
使ってください
スキップしてください
CI が失敗したため、400 行のログではなく PR の概要が必要な場合
テストは合格しましたが、出力はすでに短いです
AI エージェントがローカルまたは CI でテスト/ビルドを実行する
監査のために完全なログをすでにディスクに保存しています
長い繰り返し障害の出力 (jest、pytest、webpack)
プラットフォームはすでに十分に切り詰められています
ツールの出力がコンテキストの制限に達する MCP ワークフロー
コンプライアンスアーカイブには完全なログが必要です
必要に応じて完全なログを保存してください。
npm テスト 2>&1 |ティーフル.log |ログスリム
開発
npmインストール
npmテスト
npm ビルドを実行する
npm デモを実行する
貢献する
logslim は、新しいエラー コードや新しいログ形式、あるいはその両方を学習するたびに、より鮮明になります。
最初の貢献は簡単です:
エラー修正カード (TypeScript / Node / npm) を追加します。これは、約 5 分間の純粋な JSON PR です。
あまり圧縮されていないログを共有します。logslim マングル ツールからの実際の出力を貼り付けます。
新しいランナーのサポートを追加 — Playwright、pytest、vitest、cargo、gradle…
ここから始めましょう: 良い最初の問題
· 貢献.md
MIT — アカウントは必要なく、自由に使用できます。
PR に関する構造化された CI 障害の概要 + AI エージェントのコンパクトなテスト/ビルド ログ。 GitHub アクション、CLI、MCP。
www.npmjs.com/package/logslim
トピックス
読み込み中にエラーが発生しました。このページをリロードしてください。
1
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Structured CI failure summaries on PRs + compact test/build logs for AI agents. GitHub Action, CLI, MCP. - P156HAM/logslim

GitHub - P156HAM/logslim: Structured CI failure summaries on PRs + compact test/build logs for AI agents. GitHub Action, CLI, MCP. · GitHub
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
P156HAM
/
logslim
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits .github .github action action errors errors scripts scripts src src test test .gitignore .gitignore .npmignore .npmignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
CI failed? Get a 5-line PR summary — not a 400-line Actions log.
Agent reading test output? Cut 80–95% of the tokens.
When tests or builds fail, you scroll GitHub Actions logs. When Claude Code or Cursor runs
npm test , the agent reads everything — progress bars, 120 identical warnings, 40 frames
of node_modules . logslim fixes both:
CI / humans — GitHub Action posts structured failures on your PR (file, line, fix hints)
Agents / tokens — CLI + MCP compacts noisy output before an LLM reads it (~80–95% savings on failures)
npx logslim -- npm test
No account. No API key. MIT open source.
GitHub Action — PR failure summary
When CI fails, post a readable summary on the pull request instead of making reviewers dig
through Actions logs.
permissions :
contents : read
pull-requests : write
jobs :
test :
runs-on : ubuntu-latest
steps :
- uses : actions/checkout@v4
- uses : actions/setup-node@v4
with :
node-version : 20
- run : npm ci
- name : Run tests
id : test
run : npm test 2>&1 | tee test-output.log
continue-on-error : true
- name : Post failure summary
if : steps.test.outcome == 'failure' && github.event_name == 'pull_request'
uses : P156HAM/logslim/action@v0.3.0
with :
log-file : test-output.log
exit-code : 1
github-token : ${{ secrets.GITHUB_TOKEN }}
- name : Fail job
if : steps.test.outcome == 'failure'
run : exit 1
What gets posted on the PR:
Structured failures with file:line and messages
Fix hints for known codes ( TS2339 , ERESOLVE , …)
Token/log reduction stats (useful when agents also read the output)
The Action uses the same engine as the CLI — compaction, error extraction, and code cards.
You get human-readable PR comments and agent-ready JSON from one tool.
git clone https://github.com/P156HAM/logslim.git
cd logslim && npm install && npm run build && npm run demo
Or without cloning:
npx logslim -- node -e " console.log('ok'); for(let i=0;i<30;i++)console.log('warn '+i); throw Error('fail') "
Before → After
BEFORE — what the agent reads today (~18 lines here; real runs are 500–3000):
PASS src/utils.test.ts
console.warn deprecated prop id=1000
console.warn deprecated prop id=1001
console.warn deprecated prop id=1002
... (same warning ×120)
FAIL src/checkout/cart.test.ts
Expected: 89.10
Received: 99.00
at cart.test.ts:48:27
at node_modules/jest-circus/build/utils.js:298:28
at node_modules/jest-circus/build/utils.js:231:10
at node_modules/jest-circus/build/run.js:252:3
... (15 more node_modules frames)
Test Suites: 1 failed, 1 passed, 2 total
AFTER — what logslim gives the agent:
PASS src/utils.test.ts
console.warn deprecated prop id=1000
console.warn deprecated prop id=1001
console.warn deprecated prop id=1002
(+5 similar lines omitted by logslim)
FAIL src/checkout/cart.test.ts
Expected: 89.10
Received: 99.00
at cart.test.ts:48:27
at node_modules/jest-circus/build/utils.js:298:28
… 3 vendor/internal frames collapsed by logslim
Test Suites: 1 failed, 1 passed, 2 total
Same failure. Same fix. ~ 92% fewer tokens .
npm test ──► logslim ──► agent / CI / you
│
├─ 1. Strip ANSI colors & spinner garbage
├─ 2. Dedupe repeated lines (warn spam)
├─ 3. Collapse node_modules stack frames
├─ 4. Group similar lines (timestamps/ids masked)
├─ 5. Extract structured errors (file, line, message)
├─ 6. Attach fix cards for known codes (TS2339, ERESOLVE…)
└─ 7. Optional token budget (trim middle, keep errors)
Failure mode (default): tests pass → light cleanup only. Tests fail → full pipeline.
You only pay the compaction cost when something actually broke.
Every removed section is marked in place ( (+47 similar lines omitted by logslim) )
so the agent knows data was elided and can re-run the raw command if needed.
npm install -g logslim
# or zero-install:
npx logslim -- npm test
Requires Node 18+.
logslim -- npm test
logslim -- python -m pytest -x
logslim -- npx tsc --noEmit
Exit code is preserved. Output on stdout is compacted. Stats on stderr.
npm test 2>&1 | logslim
npm test ; logslim --exit-code $? 2>&1 < full.log # if you saved output
JSON — for agents and CI
logslim --json -- npm test
{
"exitCode" : 1 ,
"failed" : true ,
"compacted" : " FAIL src/checkout/cart.test.ts \n ... " ,
"errors" : [
{
"file" : " cart.test.ts " ,
"line" : 48 ,
"message" : " Expected: 89.10, Received: 99.00 " ,
"kind" : " assertion "
}
],
"codes" : [
{
"id" : " TS2339 " ,
"lang" : " typescript " ,
"meaning" : " Property does not exist on type " ,
"fix_steps" : [
" Check for typos " ,
" Extend the interface " ,
" Use optional chaining "
]
}
],
"stats" : {
"tokensIn" : 3296 ,
"tokensOut" : 252 ,
"saved" : 0.92 ,
"applied" : " full "
}
}
The agent reads compacted + errors + codes — not thousands of lines of prose.
logslim --json --attach git,ci -- npm test
Prepends: branch: feat/x | commit: a3f2c1d | pr: #42 (from GITHUB_* env vars).
Flag
What it does
--mode failure
Compact hard only on failure (default)
--mode full
Always compact hard
--mode light
Strip ANSI only, never aggressive dedupe
--json
Structured output (see above)
--attach git,ci
Prepend branch/commit/CI metadata
--budget 2000
Hard token cap; errors + head/tail survive
--exit-code N
For pipe mode when you know the exit code
--no-codes
Skip error code fix cards
--no-stats
Hide stderr savings footer
MCP server (Claude Code / Cursor)
Lets the agent call compaction as a tool — no manual piping.
Project .mcp.json or Claude Desktop config:
{
"mcpServers" : {
"logslim" : {
"command" : " npx " ,
"args" : [ " -y " , " logslim-mcp " ]
}
}
}
Tool: compact_output — pass output (raw log text) and optional exit_code .
Returns compacted text, extracted errors, fix cards, and stats.
npm run build && npm run mcp
Tell your agent to use it
Add to CLAUDE.md , AGENTS.md , or .cursor/rules :
When running tests, builds, or linters that produce verbose output:
- Prefer: ` logslim --mode failure --json -- <command> `
- Read the ` compacted ` , ` errors ` , and ` codes ` fields before debugging.
- If output was elided, re-run the raw command only if you need full logs.
Error code fix cards
When logs contain known codes, logslim attaches a short fix card (~30 tokens)
instead of making the agent guess or search docs.
Hand-curated pocket references — not scraped docs. PRs welcome to add codes.
Log type
Lines
Tokens
Saved
Jest (warn spam + 1 failure)
149 → 25
~3,300 → ~250
92%
Webpack build (asset noise + 2 TS errors)
548 → 55
~8,900 → ~1,000
88%
Pytest (25 identical failures)
356 → 153
~4,300 → ~1,500
64%
Token counts are estimated (~4 chars/token). Good for relative savings, not billing.
import { compact , process } from "logslim" ;
const { text , stats } = compact ( rawLog , { mode : "failure" , exitCode : 1 } ) ;
const result = process ( rawLog , {
mode : "failure" ,
exitCode : 1 ,
attach : [ "git" , "ci" ] ,
} ) ;
// result.text, result.errors, result.codes, result.stats
When to use logslim
Use it
Skip it
CI failed and you want a PR summary, not a 400-line log
Tests passed and output is already short
AI agents running tests/builds locally or in CI
You already tee full logs to disk for audit
Long repetitive failure output (jest, pytest, webpack)
Platform already truncates well enough for you
MCP workflows where tool output hits context limits
You need full logs for compliance archive
Keep full logs if you need them:
npm test 2>&1 | tee full.log | logslim
Development
npm install
npm test
npm run build
npm run demo
Contributing
logslim gets sharper every time it learns a new error code or a new log format — and both
are easy first contributions:
Add an error fix card (TypeScript / Node / npm) — a ~5-minute, pure-JSON PR.
Share a log that compacts badly — paste real output from a tool logslim mangles.
Add support for a new runner — Playwright, pytest, vitest, cargo, gradle…
Start here: good first issues
· CONTRIBUTING.md
MIT — use freely, no account required.
Structured CI failure summaries on PRs + compact test/build logs for AI agents. GitHub Action, CLI, MCP.
www.npmjs.com/package/logslim
Topics
There was an error while loading. Please reload this page .
1
fork
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
