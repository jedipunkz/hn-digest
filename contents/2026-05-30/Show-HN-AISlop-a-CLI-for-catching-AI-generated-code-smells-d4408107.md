---
source: "https://github.com/scanaislop/aislop"
hn_url: "https://news.ycombinator.com/item?id=48322956"
title: "Show HN: AISlop, a CLI for catching AI generated code smells"
article_title: "GitHub - scanaislop/aislop: Catch the slop AI coding agents leave in your code. 40+ rules, 7 languages, deterministic, no LLM. MIT. · GitHub"
author: "Heavykenny"
captured_at: "2026-05-30T11:47:49Z"
capture_tool: "hn-digest"
hn_id: 48322956
score: 72
comments: 60
posted_at: "2026-05-29T13:37:38Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AISlop, a CLI for catching AI generated code smells

- HN: [48322956](https://news.ycombinator.com/item?id=48322956)
- Source: [github.com](https://github.com/scanaislop/aislop)
- Score: 72
- Comments: 60
- Posted: 2026-05-29T13:37:38Z

## Translation

タイトル: Show HN: AISlop、AI が生成したコードの匂いをキャッチするための CLI
記事のタイトル: GitHub - scanaislop/aislop: AI コーディング エージェントがコードに残すスロップをキャッチします。 40 以上のルール、7 つの言語、決定論的、LLM なし。マサチューセッツ工科大学· GitHub
説明: AI コーディング エージェントがコードに残した残骸をキャッチします。 40 以上のルール、7 つの言語、決定論的、LLM なし。マサチューセッツ工科大学- スキャナイスロップ/アイスロップ
HN テキスト: こんにちは、私は Kenny です。アイスロップを構築しています。クロードコード、コーデックス、オープンコードを何度か使用し、いくつかのスロップに気づいた後、これに取り組み始めました。これらは構文ではなく、ほとんどのテストに合格します。空の catch ブロック、役に立たないコメント、重複したヘルパー、デッド コードなどのパターンです。そこで、これらのパターンをスキャンしてチェックするツールを構築し、それをフックに接続して、ツール呼び出しのたびにエージェントがスロップをチェックするようにしました。 npx aislop scan を使って試すことができます。すべてローカルであり、コードは転送されません。ありがとう。

記事本文:
GitHub - scanaislop/aislop: AI コーディング エージェントがコードに残すスロップをキャッチします。 40 以上のルール、7 つの言語、決定論的、LLM なし。マサチューセッツ工科大学· GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Spark インテリジェントなアプリを構築してデプロイする
GitHub モデルのプロンプトの管理と比較
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
スキャナイスロップ
/
アイスロップ
公共
通知
サインインする必要があります

o 通知設定を変更する
追加のナビゲーション オプション
コード
この GitHub アクションをプロジェクトで使用する このアクションを既存のワークフローに追加するか、新しいワークフローを作成します マーケットプレイスで表示する メイン ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
120 コミット 120 コミット .aislop .aislop .github .github アセット アセット ドキュメント ドキュメントエディター/ vscode エディター/ vscode サンプル サンプル スキーマ スキーマ スクリプト スクリプト src src テスト テスト .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore .nvmrc .nvmrc .pre-commit-hooks.yaml .pre-commit-hooks.yaml AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md action.yml action.yml biome.json biome.json knip.json knip.json package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.json tsconfig.json tsdown.config.ts tsdown.config.ts vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI コーディング エージェントがコード内に残した不備を把握します。
Claude Code、Cursor、Codex、および OpenCode が残したパターン: 自明のコード上の説明的なコメント、キャストのような飲み込まれた例外、幻覚的なインポート、重複したヘルパー、デッド コード、ToDo スタブ、過大な関数。テストは合格です。糸くずが通ります。とにかくコードが腐る。
アイスロップが彼らを捕まえる。 7 つの言語 (TS/JS、Python、Go、Rust、Ruby、PHP、Java) にわたる 40 以上のルール。すべての変更に 0 ～ 100 のスコアを付けます。サブセカンド。決定的 — ランタイム パスに LLM はなく、同じコードが入力され、同じスコアが出力されます。 MIT ライセンスの無料 CLI。
npxアイスロップスキャン
インストールは必要ありません。あらゆるプロジェクトに対応します。数秒でスコアを取得します。
npx aislop fix # 自動修正の問題
npx aislop fix -f # アグレッシブな修正 (deps、未使用の fi

レ）
npx aislop ci # CI モード (JSON + ゲート)
npx aislop フック install --claude # 編集ごとのフック
公開バッジ : README にスコアを表示します
[![ aislop ] ( https://badges.scanaislop.com/score/<owner>/<repo>.svg )] ( https://scanaislop.com )
npx aislop バッジを実行して自動生成します。 scanaislop.com で無料でご利用いただけます。
# インストールせずに実行する
npxアイスロップスキャン
#npm
npm install --save-dev aislop
# 糸
糸追加 --dev aislop
#pnpm
pnpm add -D aislop
# グローバル
npm install -g aislop
GitHub パッケージでは @scanaislop/aislop としても利用できます。
npx aislop scan # 現在のディレクトリ
npx aislop scan ./src # 特定のディレクトリ
npx aislop scan --changes # HEAD から変更されたファイル
npx aislop scan --staged # ステージングされたファイルのみ
npx aislop scan --json # JSON 出力
npx aislop scan --sarif # SARIF 2.1.0 出力 (GitHub コード スキャン)
除外ファイル: node_modules 、 .git 、 dist 、 build 、カバレッジはデフォルトで除外されます。 .aislop/config.yml にさらに追加します。
除外する:
- " **/*.test.ts "
- ソース/生成
または CLI 経由: npx aislop scan --exclude "**/*.test.ts,dist"
ルールごとの重大度 : ID によってルールの重大度を上書きするか、オフにします。
# .aislop/config.yml
ルール：
ai-slop/narrative-comment : 警告 # エラー |警告 |オフ
ai-slop/trivial-comment : " off " # このルールを完全に削除します
セキュリティ/ハードコードされたシークレット: エラー
off は一致する診断を削除します。エラー/警告は、スコアリングとレポートの前に重大度を書き換えます。マップが存在しない場合は、デフォルトの動作が維持されます。
Extend config : プロジェクト構成は親を拡張できます。
# .aislop/config.yml
拡張: ../../.aislop/base.yml
シ：
failedBelow : 80 # 特定のキーをオーバーライドします
エディターの検証: .aislop/config.yml のオートコンプリートと検証のために、エディターで schema/aislop.config.schema.json の JSON スキーマを指定します。 pnpm gen:schema を使用してソース構成スキーマから再生成します。
mec とは何かを自動修正する

hanical (フォーマッタ、未使用のインポート、デッドコード)。コンテキストが必要な問題については、完全な診断情報をエージェントに渡します。
npx aislop fix # 安全な自動修正
npx aislop fix -f # アグレッシブ: deps、未使用ファイル
エージェントに引き継ぎます
自動修正で解決できない場合は、残りの問題を完全なコンテキストとともにコーディング エージェントに渡します。
npx aislop fix --claude # クロードコード
npx aislop fix --cursor # カーソル (クリップボードにコピー)
npx aislop fix --gemini # Gemini CLI
npx aislop fix --codex # Codex CLI
# 関連項目: --windsurf、--amp、--aider、--goose、--pi、--crush、--opencode、--warp、--kimi、--antigravity、--deep-agents、--vscode
npx aislop fix --prompt # プロンプトの印刷 (エージェントに依存しない)
フックを取り付ける
エージェントが編集されるたびに実行されます。フィードバックはすぐに戻ってきます。
npx aislop フック install --claude # クロード コード
npx aislop フック install --cursor # カーソル
npx aislop フック インストール --gemini # Gemini CLI
npx aislop フック インストール --pi # pi
npx aislop フック インストール # サポートされているすべてのエージェント
npx aislop フック インストール クロード カーソル # 特定のエージェント
ランタイム アダプター (スキャン + フィードバック): claude 、cursor 、gemini 、pi 。
ルールのみ (エージェントがルールを読み取ります): codex 、 Windsurf 、 cline 、 kmocode 、 antigravity 、 copilot 。
クオリティゲート モード : スコアがベースラインを下回った場合にブロックします。
npx aislop フック インストール --claude --quality-gate
npx aislop フックベースライン # ベースラインを再キャプチャ
npx aislop フック ステータス # インストールされているリスト
npx aislop フックアンインストール --claude # 削除
ドキュメント: /docs/hooks
Aislop を Claude Desktop、Cursor、Codex の MCP ツールとして公開します。
// ~/.cursor/mcp.json または Claude デスクトップ構成
{
"mcpサーバー": {
"アイスロップ" : {
"コマンド" : " npx " ,
"args" : [ " -y " 、 " aislop-mcp " ]
}
}
}
ツール: aislop_scan 、 aislop_fix 、 aislop_why 、 aislop_baseline
npx aislop ci # JSON 出力、スコア < しきい値の場合は 1 を終了します
その他のコマンド
npxアイスロップini

t # .aislop/config.ymlを作成します
npx aislop init --strict # エンタープライズ グレード ゲート: すべてのエンジン、タイプチェック、failBelow 85
npx aislop ルール # ルールのリスト
npx Aislop バッジ # バッジの URL を印刷する
npx aislop トレンド # スコア履歴を経時的に表示
npx aislop # インタラクティブメニュー
スコア履歴: 通常の (プロジェクト全体、対話型) スキャンでは、コンパクトなレコードが .aislop/history.jsonl に追加されます (タイムスタンプ、スコア、エラー/警告数、ファイル数、CLI バージョン)。 aislop トレンドはそれを読み取り、テーブルと最近のスコアの ASCII スパークラインを出力します。履歴はローカルの副作用のみです。履歴は、CI 内、または AISLOP_NO_HISTORY=1 が設定されている場合に --json / --sarif 出力に対して書き込まれることはないため、マシンの出力はクリーンなままになります。
npx aislop scan --staged
または、バンドルされたフックを介してプリコミット フレームワークに接続します。
# .pre-commit-config.yaml
リポジトリ:
- リポジトリ : https://github.com/scanaislop/aislop
リビジョン：v0.9.4
フック:
- ID : アイスロー
GitHub アクション
npx aislop init を実行してワークフロー プロンプトを受け入れるか、手動で追加します。
- 使用:actions/checkout@v4
- 使用:actions/setup-node@v4
付き:
ノードバージョン: 20
- 実行: npx aislop@latest ci 。
複合アクション:
- 使用:actions/checkout@v4
- 使用: scanaislop/aislop@v0.8
GitHub コード スキャン (SARIF) : SARIF 2.1.0 レポートを生成し、アップロードして、結果が [セキュリティ] タブに表示されるようにします。
- 実行: npx aislop@latest scan 。 --sarif > aislop.sarif
- 使用: github/codeql-action/upload-sarif@v3
付き:
sarif_file : aislop.sarif
クオリティゲート
.aislop/config.yml で最小スコアを設定します。
シ：
以下の失敗: 70
aislop ci は、スコア < しきい値の場合に 1 を終了します。ドキュメント: CI/CD
scanaislop はチーム向けのホスト型プラットフォームです。
スコアしきい値を備えた PR ゲート
標準階層 (組織 → チーム → プロジェクト)
ダッシュボードとエージェントの属性
同じエンジン、同じスコア。 CLI は MIT ライセンスを取得しています。さらに詳しく→
AI コーディング ツールは、

nd はテストに合格しますが、エンジニアが作成しないパターンが付属しています。 aislop は 1 つのスコア、1 つのゲートを提供し、可能なものは自動修正します。
1 つのスコア: 0 ～ 100、CI で適用されます。重み付けされているため、雑なパターンはスタイル ノイズよりも強くヒットします。
最初に自動修正 : フォーマッタ、未使用のインポート、デッドコードを機械的にクリアします。残りの部分は完全なコンテキストとともにエージェントに渡します。
決定的: Regex + AST + 標準ツール。 LLM や API 呼び出しはありません。同じコードが入力され、同じスコアが出力されます。
Zero-config start : npx aislop scan はどのリポジトリでも機能します。 .aislop/config.yml を追加して調整します。
6 つの決定論的エンジンが並行して実行されます。
完全なルールのリファレンスを参照してください。
aislop ルールは、ローカル フィクスチャだけでなく、パブリック スキャンとベンチマークから派生した障害モードによって形成されます。研究プログラムでは、反復可能なオープンソース スキャンの実行方法を定義しています。つまり、コホートの固定、生の JSON の保存、結果の分類、回帰テストによるノイズの多いルールの修正、および制限の公開です。
インストール · コマンド · ルール · 設定 · スコアリング · CI/CD · テレメトリ · 研究プログラム
質問、ルールリクエスト、誤検知のトリアージに関するディスカッション · バグの問題
COTRIBUTING.md を参照してください。 AI アシスタント: AGENTS.md 。
構築元: Biome 、 oxlint 、 knip 、 ruff 、 golangci-lint 、 expo-doctor
.github/workflows/contributors.yml によって自動更新されます。コミットメールをリンクするか、 .github/contributors-overrides.json に追加します。
AI コーディング エージェントがコード内に残した不備を把握します。 40 以上のルール、7 つの言語、決定論的、LLM なし。マサチューセッツ工科大学
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
9
フォーク
レポートリポジトリ
リリース
28
v0.10.0
最新の
2026 年 5 月 30 日
+ 27 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitH

株式会社ユーブ
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Catch the slop AI coding agents leave in your code. 40+ rules, 7 languages, deterministic, no LLM. MIT. - scanaislop/aislop

Hi, I’m Kenny, I’ve been building aislop. I starting working on this after using Claude Code, codex and opencode several times and noticing some slops. They aren’t syntax and passes most tests, they are patterns like empty catch blocks, useless comments, duplicated helpers, dead code and many more. So I built a tool to scan and check for these patterns and wired it into hooks so after each tool call, the agent checks for the slops. You can try it out with npx aislop scan. It’s all local and no code is transferred. Thank you.

GitHub - scanaislop/aislop: Catch the slop AI coding agents leave in your code. 40+ rules, 7 languages, deterministic, no LLM. MIT. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Spark Build and deploy intelligent apps
GitHub Models Manage and compare prompts
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
scanaislop
/
aislop
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Use this GitHub action with your project Add this Action to an existing workflow or create a new one View on Marketplace main Branches Tags Go to file Code Open more actions menu Folders and files
120 Commits 120 Commits .aislop .aislop .github .github assets assets docs docs editors/ vscode editors/ vscode examples examples schema schema scripts scripts src src tests tests .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore .nvmrc .nvmrc .pre-commit-hooks.yaml .pre-commit-hooks.yaml AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md action.yml action.yml biome.json biome.json knip.json knip.json package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.json tsconfig.json tsdown.config.ts tsdown.config.ts vitest.config.ts vitest.config.ts View all files Repository files navigation
Catch the slop AI coding agents leave in your code.
The patterns Claude Code, Cursor, Codex, and OpenCode leave behind: narrative comments above self-explanatory code, swallowed exceptions, as any casts, hallucinated imports, duplicated helpers, dead code, todo stubs, oversized functions. Tests pass. Lint passes. The code rots anyway.
aislop catches them. 40+ rules across 7 languages (TS/JS, Python, Go, Rust, Ruby, PHP, Java). Scores every change 0–100. Sub-second. Deterministic — no LLM in the runtime path, same code in, same score out. MIT-licensed, free CLI.
npx aislop scan
No install needed. Works on any project. Get your score in seconds.
npx aislop fix # auto-fix issues
npx aislop fix -f # aggressive fixes (deps, unused files)
npx aislop ci # CI mode (JSON + gate)
npx aislop hook install --claude # per-edit hook
Public badge : Show your score on your README
[ ![ aislop ] ( https://badges.scanaislop.com/score/<owner>/<repo>.svg )] ( https://scanaislop.com )
Run npx aislop badge to auto-generate. Free at scanaislop.com .
# Run without installing
npx aislop scan
# npm
npm install --save-dev aislop
# yarn
yarn add --dev aislop
# pnpm
pnpm add -D aislop
# Global
npm install -g aislop
Also available as @scanaislop/aislop on GitHub Packages.
npx aislop scan # current directory
npx aislop scan ./src # specific directory
npx aislop scan --changes # changed files from HEAD
npx aislop scan --staged # staged files only
npx aislop scan --json # JSON output
npx aislop scan --sarif # SARIF 2.1.0 output (GitHub code scanning)
Exclude files : node_modules , .git , dist , build , coverage excluded by default. Add more in .aislop/config.yml :
exclude :
- " **/*.test.ts "
- src/generated
Or via CLI: npx aislop scan --exclude "**/*.test.ts,dist"
Per-rule severity : Override the severity of any rule by id, or turn it off:
# .aislop/config.yml
rules :
ai-slop/narrative-comment : warning # error | warning | off
ai-slop/trivial-comment : " off " # drop this rule entirely
security/hardcoded-secret : error
off drops matching diagnostics; error / warning rewrites severity before scoring and reporting. Absent map keeps default behavior.
Extend config : Project config can extend a parent:
# .aislop/config.yml
extends : ../../.aislop/base.yml
ci :
failBelow : 80 # override specific keys
Editor validation : Point your editor at the JSON Schema in schema/aislop.config.schema.json for autocomplete and validation of .aislop/config.yml . Regenerate it from the source config schema with pnpm gen:schema .
Auto-fix what's mechanical (formatters, unused imports, dead code). For issues that need context, hand off to your agent with full diagnostic info.
npx aislop fix # safe auto-fixes
npx aislop fix -f # aggressive: deps, unused files
Hand off to agent
When auto-fix can't solve it, pass the remaining issues to your coding agent with full context:
npx aislop fix --claude # Claude Code
npx aislop fix --cursor # Cursor (copies to clipboard)
npx aislop fix --gemini # Gemini CLI
npx aislop fix --codex # Codex CLI
# Also: --windsurf, --amp, --aider, --goose, --pi, --crush, --opencode, --warp, --kimi, --antigravity, --deep-agents, --vscode
npx aislop fix --prompt # print prompt (agent-agnostic)
Install hook
Runs after every agent edit. Feedback flows back immediately.
npx aislop hook install --claude # Claude Code
npx aislop hook install --cursor # Cursor
npx aislop hook install --gemini # Gemini CLI
npx aislop hook install --pi # pi
npx aislop hook install # all supported agents
npx aislop hook install claude cursor # specific agents
Runtime adapters (scan + feedback): claude , cursor , gemini , pi .
Rules-only (agent reads rules): codex , windsurf , cline , kilocode , antigravity , copilot .
Quality-gate mode : Blocks if score regresses below baseline.
npx aislop hook install --claude --quality-gate
npx aislop hook baseline # re-capture baseline
npx aislop hook status # list installed
npx aislop hook uninstall --claude # remove
Docs: /docs/hooks
Expose aislop as MCP tools for Claude Desktop, Cursor, Codex:
// ~/.cursor/mcp.json or Claude Desktop config
{
"mcpServers" : {
"aislop" : {
"command" : " npx " ,
"args" : [ " -y " , " aislop-mcp " ]
}
}
}
Tools : aislop_scan , aislop_fix , aislop_why , aislop_baseline
npx aislop ci # JSON output, exits 1 if score < threshold
Other commands
npx aislop init # create .aislop/config.yml
npx aislop init --strict # enterprise-grade gate: all engines, typecheck, failBelow 85
npx aislop rules # list rules
npx aislop badge # print badge URL
npx aislop trend # show score history over time
npx aislop # interactive menu
Score history : a normal (full-project, interactive) scan appends a compact record to .aislop/history.jsonl (timestamp, score, error/warning counts, file count, CLI version). aislop trend reads it and prints a table plus an ASCII sparkline of recent scores. History is a local side effect only: it is never written for --json / --sarif output, in CI, or when AISLOP_NO_HISTORY=1 is set, so machine output stays clean.
npx aislop scan --staged
Or wire it into the pre-commit framework via the bundled hook:
# .pre-commit-config.yaml
repos :
- repo : https://github.com/scanaislop/aislop
rev : v0.9.4
hooks :
- id : aislop
GitHub Actions
Run npx aislop init and accept the workflow prompt, or add manually:
- uses : actions/checkout@v4
- uses : actions/setup-node@v4
with :
node-version : 20
- run : npx aislop@latest ci .
Composite action :
- uses : actions/checkout@v4
- uses : scanaislop/aislop@v0.8
GitHub code scanning (SARIF) : emit a SARIF 2.1.0 report and upload it so findings appear in the Security tab:
- run : npx aislop@latest scan . --sarif > aislop.sarif
- uses : github/codeql-action/upload-sarif@v3
with :
sarif_file : aislop.sarif
Quality gate
Set minimum score in .aislop/config.yml :
ci :
failBelow : 70
aislop ci exits 1 when score < threshold. Docs: CI/CD
scanaislop is the hosted platform for teams:
PR gates with score thresholds
Standards hierarchy (org → team → project)
Dashboards and agent attribution
Same engines, same scores. CLI is MIT-licensed. Learn more →
AI coding tools generate code that compiles and passes tests but ships with patterns no engineer would write. aislop gives you one score, one gate, and auto-fixes what it can.
One score : 0-100, enforced in CI. Weighted so sloppy patterns hit harder than style noise.
Auto-fix first : Clears formatters, unused imports, dead code mechanically. Hands off the rest to your agent with full context.
Deterministic : Regex + AST + standard tooling. No LLMs, no API calls. Same code in, same score out.
Zero-config start : npx aislop scan works on any repo. Add .aislop/config.yml to tune.
Six deterministic engines run in parallel:
See the full rules reference .
aislop rules are shaped by public scans and benchmark-derived failure modes, not only local fixtures. The research program defines how to run repeatable open-source scans: pin the cohort, store raw JSON, classify findings, fix noisy rules with regression tests, and publish the limits.
Installation · Commands · Rules · Config · Scoring · CI/CD · Telemetry · Research program
Discussions for questions, rule requests, and false-positive triage · Issues for bugs
See CONTRIBUTING.md . AI assistants: AGENTS.md .
Built on: Biome , oxlint , knip , ruff , golangci-lint , expo-doctor
Auto-updated by .github/workflows/contributors.yml . Link commit email or add to .github/contributors-overrides.json .
Catch the slop AI coding agents leave in your code. 40+ rules, 7 languages, deterministic, no LLM. MIT.
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
9
forks
Report repository
Releases
28
v0.10.0
Latest
May 30, 2026
+ 27 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
