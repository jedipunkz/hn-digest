---
source: "https://github.com/Asyncinnovator/hallint"
hn_url: "https://news.ycombinator.com/item?id=48859105"
title: "Show HN: Hallint – lint AI-generated code for security issues"
article_title: "GitHub - Asyncinnovator/hallint: Detect security and quality issues in AI-generated code · GitHub"
author: "coder_xyz"
captured_at: "2026-07-10T13:04:17Z"
capture_tool: "hn-digest"
hn_id: 48859105
score: 1
comments: 1
posted_at: "2026-07-10T12:43:06Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Hallint – lint AI-generated code for security issues

- HN: [48859105](https://news.ycombinator.com/item?id=48859105)
- Source: [github.com](https://github.com/Asyncinnovator/hallint)
- Score: 1
- Comments: 1
- Posted: 2026-07-10T12:43:06Z

## Translation

タイトル: Show HN: Hallint – セキュリティ問題に対して lint AI が生成したコード
記事のタイトル: GitHub - Asyncinnovator/hallint: AI 生成コードのセキュリティと品質の問題を検出 · GitHub
説明: AI 生成コードのセキュリティと品質の問題を検出 - Asyncinnovator/hallint

記事本文:
GitHub - Asyncinnovator/hallint: AI 生成コードのセキュリティと品質の問題を検出 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
アシンシンノベーター
/
ハリント
公共
通知
chan にサインインする必要があります

GE通知設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
33 コミット 33 コミット .github/ workflows .github/ workflows パッケージ パッケージ .eslintrc.json .eslintrc.json .gitignore .gitignore ライセンス ライセンス README.md README.md Hallint.config.ts Hallint.config.ts package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
██╗ ██╗ █████╗ ██╗ ██╗ ██╗███╗ ██╗████████╗
██║ ██║██╔══██╗██║ ██║ ██║████╗ ██║╚══██╔══╝
███████║███████║██║ ██║ ██║██╔██╗ ██║ ██║
██╔══██║██╔══██║██║ ██║ ██║██║╚██╗██║ ██║
██║ ██║██║ ██║███████╗███████╗██║██║ ╚████║ ██║
╚═╝ ╚═╝╚═╝ ╚═╝╚═════╝╚══════╝╚═╝╚═╝ ╚═══╝ ╚═╝
AI が生成したコードの静的分析。
人工知能が残したセキュリティ ホールを、本番環境に到達する前に発見します。
AI コーディング アシスタント (Copilot、Cursor など) は高速ですが、ハードコードされたシークレット、SQL インジェクション、認証の欠落、安全でない評価など、同じクラスのバグが常に発生します。ハリントは彼らを捕まえます。
npm インストール @asyncinnovator/hallint
または、インストールせずに実行します。
npx @asyncinnovator/hallint-cli ./src
要件: Node.js >= 18
npx @asyncinnovator/hallint-cli ./src
グロブパターンでスキャン
npx @asyncinnovator/hallint-cli

" ./src/**/*.ts "
重要かつ重大な問題のみを表示する
npx @asyncinnovator/hallint-cli ./src --min-severity high
すべてのルールを実行する
npx @asyncinnovator/hallint-cli ./src --rules all
カラー出力を無効にする (ログ/CI 用)
npx @asyncinnovator/hallint-cli ./src --no-color
オプション
オプション
説明
デフォルト
--ルール
ルールセット: 推奨またはすべて
推奨
--min-severity
最小重大度: クリティカル、高、中、低情報
情報
--色なし
カラー出力を無効にする
オフ
--助けて
ヘルプを表示する
--バージョン
バージョンを表示
終了コード
コード
意味
0
問題は見つかりませんでした
1
1 つ以上の重大なまたは重要な所見
2
予期しないエラー
図書館の利用状況
import { scan } から '@asyncinnovator/hallint'
const result = スキャンを待ちます ( {
ファイル: [ './src/**/*.ts' ] 、
})
結果。所見。 forEach ( f => {
コンソール。 log (` ${ f . 重大度 } [ ${ f . ルール ID } ] ${ f . ファイルパス } : ${ f . line } ` )
コンソール。ログ (` ${ f . メッセージ } `)
コンソール。ログ ( ` 修正: ${ f . 修正 } ` )
} )
オプションあり
const result = スキャンを待ちます ( {
ファイル: [ './src/**/*.ts' ] 、
ルール : '推奨' 、
minSeverity : '高' 、
無視: [ '**/node_modules/**' 、 '**/dist/**' ] 、
})
文字列を直接スキャンします (ファイル システムなし)
'@asyncinnovator/hallint' から { scanSource } をインポートします
const ソース = `const apiKey = "sk-abc123def456ghi789jk"`
const 結果 = scanSource (source , 'example.ts' )
所見。 forEach ( f => console . log ( f . rulesId , f . message ) )
ScanResult の形状
{
結果: 検索 [ ] // すべての問題が見つかりました
scannedFiles: string [ ] // スキャンされたファイルのリスト
durationMs:number // かかった時間
summary: { // 重大度ごとのカウント
クリティカル: 数値
高: 数値
中: 数字
低い: 数値
情報: 番号
}
}
形を見つける
{
rulesId: string // 例: 「ハードコードされたシークレット」
重大度: 文字列 // "重大" | 「高い」 | 「中」 | 「低い」 | 「情報」
message: string // 人間が判読できる説明

に
fix: string // 提案された修正
filePath: string // ファイルへの絶対パス
line:number // 行番号
スニペット: string // 問題のあるコード行
}
ルール
ルール
重大度
何が引っかかるのか
ハードコードされたシークレット
クリティカル
ソースコード内のAPIキー、トークン、パスワード
SQLインジェクション
クリティカル
SQL クエリに挿入されるユーザー入力
安全でない評価
クリティカル
eval() または動的入力を使用した new Function()
認証チェックの欠落
高い
認証ミドルウェアを持たないルート ハンドラー
xss-innerHTML
高い
innerHTML に割り当てられたサニタイズされていない文字列
許容的なコア
高い
ルートハンドラーの cors({origin: '*' })
非同期キャッチなし
中程度
try/catch または .catch() のない非同期関数
https ではなく http
中程度
fetch または axios 呼び出しでハードコードされた http:// URL
構成
プロジェクトのルートに Hallint.config.ts を作成します。
'@asyncinnovator/hallint' からタイプ { ScanConfig } をインポートします
デフォルトのエクスポート {
ルール : '推奨' 、
minSeverity : '中' 、
無視します: [
'**/node_modules/**' ,
'**/距離/**' ,
'**/*.test.ts' ,
]、
} は省略 < ScanConfig , 'files' > を満たします
CIの統合
Hallint を PR ゲートとして追加します。重大な検出結果または重要な検出結果があれば、コード 1 で終了します。
名前：ハリント
に:
プッシュ：
ブランチ: [メイン]
プルリクエスト:
ブランチ: [メイン]
ジョブ:
ハリント :
実行: ubuntu-最新
手順:
- 使用:actions/checkout@v4
- 使用:actions/setup-node@v4
付き:
ノードバージョン: 20
- 実行: npx @asyncinnovator/hallint-cli ./src --min-severity high
プリコミットフック
npm install --save-dev ハスキー
npxハスキー初期化
echo " npx @asyncinnovator/hallint-cli ./src --min-severity high " > .husky/pre-commit
貢献する
貢献は大歓迎です。各ルールは、bad.ts/good.ts フィクスチャを含む 1 つのファイル (約 30 行) です。新しいルールは、最初のコントリビューションとして適切です。
ブランチを作成します: git checkout -b feat/rule-your-rule-name
ルールを package/core/src/rules/index.ts に追加します
追加

package/core/tests/fixtures/your-rule-name/ 内のフィクスチャ
「良好な最初の号」とラベル付けされた号は事前にスコープが設定されており、すぐに入手できます。
MIT — 個人および商用プロジェクトで無料で使用できます。
AI が生成したコード内のセキュリティと品質の問題を検出
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
0
フォーク
レポートリポジトリ
リリース
3
v0.1.3
最新の
2026 年 7 月 9 日
+ 2 リリース
貢献者
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Detect security and quality issues in AI-generated code - Asyncinnovator/hallint

GitHub - Asyncinnovator/hallint: Detect security and quality issues in AI-generated code · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
Asyncinnovator
/
hallint
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
33 Commits 33 Commits .github/ workflows .github/ workflows packages packages .eslintrc.json .eslintrc.json .gitignore .gitignore LICENSE LICENSE README.md README.md hallint.config.ts hallint.config.ts package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts View all files Repository files navigation
██╗ ██╗ █████╗ ██╗ ██╗ ██╗███╗ ██╗████████╗
██║ ██║██╔══██╗██║ ██║ ██║████╗ ██║╚══██╔══╝
███████║███████║██║ ██║ ██║██╔██╗ ██║ ██║
██╔══██║██╔══██║██║ ██║ ██║██║╚██╗██║ ██║
██║ ██║██║ ██║███████╗███████╗██║██║ ╚████║ ██║
╚═╝ ╚═╝╚═╝ ╚═╝╚══════╝╚══════╝╚═╝╚═╝ ╚═══╝ ╚═╝
Static analysis for AI-generated code.
Catch the security holes Artificial Intelligence leave behind — before they reach production.
AI coding assistants (Copilot, Cursor, etc.) are fast, but they consistently introduce the same classes of bugs: hardcoded secrets, SQL injection, missing auth, unsafe eval. hallint catches them.
npm install @asyncinnovator/hallint
Or run without installing:
npx @asyncinnovator/hallint-cli ./src
Requirements: Node.js >= 18
npx @asyncinnovator/hallint-cli ./src
Scan with glob pattern
npx @asyncinnovator/hallint-cli " ./src/**/*.ts "
Only show high and critical issues
npx @asyncinnovator/hallint-cli ./src --min-severity high
Run all rules
npx @asyncinnovator/hallint-cli ./src --rules all
Disable color output (for logs/CI)
npx @asyncinnovator/hallint-cli ./src --no-color
Options
Option
Description
Default
--rules
Rule set: recommended or all
recommended
--min-severity
Minimum severity: critical high medium low info
info
--no-color
Disable colored output
off
--help
Show help
--version
Show version
Exit codes
Code
Meaning
0
No issues found
1
One or more critical or high findings
2
Unexpected error
Library Usage
import { scan } from '@asyncinnovator/hallint'
const result = await scan ( {
files : [ './src/**/*.ts' ] ,
} )
result . findings . forEach ( f => {
console . log ( ` ${ f . severity } [ ${ f . ruleId } ] ${ f . filePath } : ${ f . line } ` )
console . log ( ` ${ f . message } ` )
console . log ( ` fix: ${ f . fix } ` )
} )
With options
const result = await scan ( {
files : [ './src/**/*.ts' ] ,
rules : 'recommended' ,
minSeverity : 'high' ,
ignore : [ '**/node_modules/**' , '**/dist/**' ] ,
} )
Scan a string directly (no file system)
import { scanSource } from '@asyncinnovator/hallint'
const source = `const apiKey = "sk-abc123def456ghi789jk"`
const findings = scanSource ( source , 'example.ts' )
findings . forEach ( f => console . log ( f . ruleId , f . message ) )
ScanResult shape
{
findings: Finding [ ] // all issues found
scannedFiles: string [ ] // list of files scanned
durationMs: number // time taken
summary: { // count per severity
critical: number
high: number
medium: number
low: number
info: number
}
}
Finding shape
{
ruleId: string // e.g. "hardcoded-secret"
severity: string // "critical" | "high" | "medium" | "low" | "info"
message: string // human-readable description
fix: string // suggested fix
filePath: string // absolute path to the file
line: number // line number
snippet: string // the offending line of code
}
Rules
Rule
Severity
What it catches
hardcoded-secret
critical
API keys, tokens, passwords in source code
sql-injection
critical
User input interpolated into SQL queries
unsafe-eval
critical
eval() or new Function() with dynamic input
missing-auth-check
high
Route handlers with no auth middleware
xss-innerHTML
high
Unsanitized strings assigned to innerHTML
permissive-cors
high
cors({ origin: '*' }) in route handlers
async-no-catch
medium
async functions with no try/catch or .catch()
http-not-https
medium
Hardcoded http:// URLs in fetch or axios calls
Configuration
Create a hallint.config.ts at your project root:
import type { ScanConfig } from '@asyncinnovator/hallint'
export default {
rules : 'recommended' ,
minSeverity : 'medium' ,
ignore : [
'**/node_modules/**' ,
'**/dist/**' ,
'**/*.test.ts' ,
] ,
} satisfies Omit < ScanConfig , 'files' >
CI Integration
Add hallint as a PR gate — it exits with code 1 on any critical or high finding:
name : hallint
on :
push :
branches : [main]
pull_request :
branches : [main]
jobs :
hallint :
runs-on : ubuntu-latest
steps :
- uses : actions/checkout@v4
- uses : actions/setup-node@v4
with :
node-version : 20
- run : npx @asyncinnovator/hallint-cli ./src --min-severity high
Pre-commit hook
npm install --save-dev husky
npx husky init
echo " npx @asyncinnovator/hallint-cli ./src --min-severity high " > .husky/pre-commit
Contributing
Contributions are welcome. Each rule is a single file (~30 lines) with a bad.ts / good.ts fixture — a new rule is a good first contribution.
Create a branch: git checkout -b feat/rule-your-rule-name
Add your rule in packages/core/src/rules/index.ts
Add fixtures in packages/core/tests/fixtures/your-rule-name/
Issues labeled good first issue are pre-scoped and ready to pick up.
MIT — free to use in personal and commercial projects.
Detect security and quality issues in AI-generated code
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
3
v0.1.3
Latest
Jul 9, 2026
+ 2 releases
Contributors
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
