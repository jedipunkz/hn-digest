---
source: "https://github.com/faisalishfaq2005/loopflow"
hn_url: "https://news.ycombinator.com/item?id=48612954"
title: "Show HN: LoopFlow – loop engineering for Claude Code"
article_title: "GitHub - faisalishfaq2005/loopflow · GitHub"
author: "faisalishfaq"
captured_at: "2026-06-20T21:32:37Z"
capture_tool: "hn-digest"
hn_id: 48612954
score: 2
comments: 0
posted_at: "2026-06-20T21:01:33Z"
tags:
  - hacker-news
  - translated
---

# Show HN: LoopFlow – loop engineering for Claude Code

- HN: [48612954](https://news.ycombinator.com/item?id=48612954)
- Source: [github.com](https://github.com/faisalishfaq2005/loopflow)
- Score: 2
- Comments: 0
- Posted: 2026-06-20T21:01:33Z

## Translation

タイトル: Show HN: LoopFlow – クロード コードのループ エンジニアリング
記事タイトル: GitHub - faisalishfaq2005/loopflow · GitHub
説明: GitHub でアカウントを作成して、faisalishfaq2005/loopflow の開発に貢献します。

記事本文:
GitHub - faisalishfaq2005/loopflow · GitHub
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
フェイサリッシュfaq2005
/
ループフロー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く

フォルダーとファイル
21 コミット 21 コミット ループ ループ src src テンプレート テンプレート テスト テスト .gitattributes .gitattributes .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
コーディング エージェントにプロンプトを表示するのをやめてください。それを促すループを設計します。
LoopF​​low は、Claude Code を自動的に実行するシステムに変えます。つまり、目標、エージェントのパイプライン、検証ゲートを 1 つの YAML ファイルで宣言します。LoopF​​low は、ゲートが通過するか、予算がなくなるか、試行制限に達するまで繰り返します。 1 人のエージェントが書き込み、別のエージェントがチェックを行い、メモリ ファイルによって毎回の実行が前回よりもスマートになります。
$loopflow テストと修正の実行
反復 1/3
▸ 修正してください…
完了 · $0.31 · 再開: クロード --resume 072f1abb…
▸ レビュー（ゲート） …
ゲート失敗 · $0.12
│ 日付パーサーの修正は ISO 文字列のみを処理します。失敗したテスト
│ エポックミリスも供給します。根本原因が解決されていない。
反復 2/3
▸ 修正してください…
完了 · $0.28
▸ レビュー（ゲート） …
完了 · 0.11 ドル
✓ 成功 · 2 回の反復 · $0.82
📺 デモビデオ
LoopFlow デモ — リリース - チェック ループ、2 回の反復
2 年間のワークフローは、プロンプトを作成し、出力を読み取り、次のプロンプトを作成するというものでした。あなたはずっとツールを握っていました。
それは変わりつつあります。 Boris Cherny (Claude Code の作成者) は次のように述べています。「私はもうクロードにプロンプ​​トを表示しません。私はクロードにプロンプ​​トを表示するループを実行しています。」
ループは再帰的な目標です。「完了」がどのようなものかを定義すると、エージェントはそこに到達するまで反復します。ただし、これを raw で実行すると、次の 3 つの鋭いエッジがあります。
エージェントは自分で宿題を採点します。修正を作成したモデルは、修正が機能すると喜んで宣言します。
無人のループはお金を消費します。ループ走行自体も

ミスを犯し、トークンを消費するループが無人で行われます。
エージェントは実行の間にすべてを忘れます。実行するたびに、前回の実行で既に学習した内容が再導出されます。
LoopFlow は、まさにこれら 3 つの問題を中心に構築された小さくて鋭いツールです。
API キーもデーモンもクラウドもありません。 claude がターミナルで動作する場合、loopflow は動作します。
npm install -g @loopflow/cli # または: npx @loopflow/cli
プロジェクトの cd
loopflow init # scaffolds .loopflow/ と 3 つのスターター ループ
oopflow run test-and-fix --dry-run # 各エージェントに何を伝えるかを正確に確認する
Loopflow run test-and-fix # 実際に実行してみる
要件: ノード 18 以降、クロード コードがインストールされ、認証されている。
# .loopflow/loops/test-and-fix.yaml
名前 : テストと修正
説明 : テスト スイートを実行し、障害を修正し、修正を確認します。
予算：
max_usd : 2.00 # 実行全体のハード上限 (すべての反復を含む)
max_iterations : 3 # ゲートが拒否できる試行回数
worktree : false # 分離された git ワークツリーで実行するには true を設定します
デフォルト:
許可モード : acceptEdits
手順:
- ID : 修正
role : > # persona — クロードのシステム プロンプトに追加
あなたは慎重な管理者です。最小限の変更を加えれば問題は解決します
問題を解決するためにテストを弱めることは決してありません。
プロンプト: |
このプロジェクトのテスト スイートを実行します。あらゆる問題の根本原因を診断して修正します
失敗。確認のために再実行します。何を変更したのか、その理由をまとめます。
- ID : レビュー
Gate : true # ← このステップが PASS になるまでループは成功しません
役割: >
あなたは懐疑的な上級エンジニアで、自分が行っていない変更をレビューしています
書く。証拠がなければ何も信用しません。
プロンプト: |
以前のエージェントは、失敗したテストを修正したと主張しています。差分を調べて、
スイートを自分で再実行し、テストが弱体化または削除されていないことを確認してください。
実行すると何が起こるか
┌─────────

─────────────┐
│ 反復回数 (≤ 最大) │
│ │
記憶 ──▶ │ ステップ：修正 ──▶ ステップ：復習（ゲート） ──┐ │
▲ │ ▲ │ │
│ │ └─ レビュアーフィードバック ◀─ 不合格 ─┤ │
│ ━━━━━━━━━━━━━━━━━━┼───────┘
│ │ パス
━─────── 走行記録 ◀─────────┘
各ステップは 1 回のヘッドレス クロード コード実行 ( claude -p ) です。ステップでは、ループのメモリ、反復内の以前のステップの出力、および再試行時のゲートのフィードバックが表示されます。
ゲートは VERDICT: PASS または VERDICT: FAIL で終了する必要があります。判定がない場合は FAIL としてカウントされません。未検証のパスはパスではありません。
FAIL の場合、ループは最初からやり直し、レビュー担当者のフィードバックがすべてのプロンプトに挿入されます。
実行ごとに、レコード (結果、コスト、および最終概要) が .loopflow/memory/<loop>.md に追加され、次の実行で読み取られます。
実際の実行では次のようになります。修正ステップで見逃したデバッグ アーティファクトをキャッチするリリース チェック ループです。
loopflow init は、盗まれるように設計された 3 つのループを提供します。
テストと修正 — 修正者 + 懐疑的なレビュー者のゲート。正規の書き込み/検証ペア。
債務監査 — 発見ループ。 .loopflow/reports/debt-audit.md を維持し、メモリを使用して修正されたもの、新しいもの、無視され続けるものを追跡します。
docs-sync — コードから逸脱したドキュメントを検索し、分離されたワークツリーで修正し、ゲートがソースに対するすべての主張を検証します。
自分だけのループはありますか？クックブックに投稿してください。コミュニティ ループは、loops/ に存在します。
LoopF​​low は意図的にデーモンを同梱しません。すでに持っているスケジューラーを使用する

e:
# cron (Linux/macOS) — 毎週月曜日の午前 9 時に負債を監査します
0 9 * * 1 cd /path/to/project && ループフロー実行債務監査
# Windows タスク スケジューラ
schtasks /create /tn "債務監査" /sc 毎週 /d MON /st 09:00 ^
/tr " cmd /c cd /d C:\path\to\project && ループフロー実行債務監査 "
CI も機能します。loopflow run docs-sync を毎週実行し、保持されているワークツリー ブランチから PR を開く GitHub アクションは、最大 20 行です。
ループは仕事を変えますが、あなたをその仕事から消し去るわけではありません。 LoopFlow の設計では、次の 3 つのことが真実であることを前提としています。
検証はまだあなたにあります。ゲートは明らかな失敗をキャッチしますが、 --verbose と claude --resume <session-id> が存在するため、ループが実際に何をしたかを読み取ることができます。読んでみてください。
理解の負債は現実のものです。あなたが書いていないコードをループが速く出荷するほど、存在するものとあなたが理解しているものとの間のギャップはより速く広がります。メモリ ファイルと保持されたワークツリーは、機械だけでなく人間によって読み取られるように設計されています。
楽な姿勢こそが危険な姿勢なのです。ループが自動的に実行されると、意見を持つことをやめたくなります。判断しながらループを設計し、出力を判断し続けます。
ループを構築します。ただし、単に退職を押す人ではなく、エンジニアであり続けるつもりの人のようにそれを構築してください。
コマンド
何をするのか
ループフロー初期化 [--force]
スターターループを含む Scaffold .loopflow/
ループフローリスト
ステップ、ゲート、バジェットを含むループをリストする
ループフロー検証 [名前]
ループ定義を検証します (すべてデフォルト)
ループフロー実行 <名前>
ループを実行する
--ドライラン
作成されたすべてのプロンプトを出力します。何も呼び出さない
-i、--反復 <n>
Budget.max_iterations をオーバーライドする
-b, --budget <米ドル>
Budget.max_usd をオーバーライドする
-v、--verbose
フルステップ出力を印刷する
終了コード: 0 成功 · 1 ループ失敗 (ゲート枯渇、バジェット、エラー) · 2 構成エラー。 Cron および CI に優しい。
CLI の動作はすべてエクスポートされます。
インポート { l

oadLoop , runLoop } "@loopflow/cli" ;
const ループ = loadLoop ( process . cwd ( ) , "テストと修正" ) ;
const result = await runLoop (loop, {root:process.cwd()});
コンソール。 log (結果 . 結果 , 結果 . コストUsd ) ;
ロードマップ
ループフロー デーモン — ループ.yaml 内の cron 式を使用した組み込みスケジューラ
並列ステップ (ワークツリー全体にファンアウト)
--json-schema による構造化ゲートの判定
ループ実行履歴とループフローログ
他のヘッドレス エージェント用のアダプター (Codex CLI など)
最も価値のある貢献は、実際の問題を解決したループです。 CONTRIBUTING.md を参照してください。コードの貢献: エンジンは、型付けされ、テストされた最大 600 行の TypeScript です。 npm テストは 1 秒以内に実行されます。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to faisalishfaq2005/loopflow development by creating an account on GitHub.

GitHub - faisalishfaq2005/loopflow · GitHub
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
faisalishfaq2005
/
loopflow
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
21 Commits 21 Commits loops loops src src templates templates tests tests .gitattributes .gitattributes .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts View all files Repository files navigation
Stop prompting your coding agent. Design the loop that prompts it.
LoopFlow turns Claude Code into a system that runs itself: you declare a goal, a pipeline of agents, and a verification gate in one YAML file — LoopFlow iterates until the gate passes, the budget runs out, or the attempt limit is hit. One agent writes, a different agent checks, and a memory file makes every run smarter than the last.
$ loopflow run test-and-fix
Iteration 1/3
▸ fix …
done · $0.31 · resume: claude --resume 072f1abb…
▸ review (gate) …
gate FAIL · $0.12
│ The date parser fix only handles ISO strings; the failing test
│ also feeds epoch millis. Root cause not addressed.
Iteration 2/3
▸ fix …
done · $0.28
▸ review (gate) …
done · $0.11
✓ success · 2 iteration(s) · $0.82
📺 Demo video
LoopFlow demo — release-check loop, 2 iterations
For two years the workflow was: write a prompt, read the output, write the next prompt. You held the tool the whole time.
That's changing. As Boris Cherny (creator of Claude Code) put it: "I don't prompt Claude anymore. I have loops running that prompt Claude."
A loop is a recursive goal : you define what "done" looks like, and the agent iterates until it gets there. But doing this raw has three sharp edges:
Agents grade their own homework. The model that wrote the fix will happily declare it works.
Unattended loops burn money. A loop running itself is also a loop making mistakes — and spending tokens — unattended.
The agent forgets everything between runs. Every run re-derives what the last run already learned.
LoopFlow is a small, sharp tool built around exactly those three problems:
No API keys, no daemon, no cloud. If claude works in your terminal, loopflow works.
npm install -g @loopflow/cli # or: npx @loopflow/cli
cd your-project
loopflow init # scaffolds .loopflow/ with three starter loops
loopflow run test-and-fix --dry-run # see exactly what each agent will be told
loopflow run test-and-fix # run it for real
Requirements: Node 18+, Claude Code installed and authenticated.
# .loopflow/loops/test-and-fix.yaml
name : test-and-fix
description : Run the test suite, fix failures, verify the fix.
budget :
max_usd : 2.00 # hard ceiling for the whole run, all iterations included
max_iterations : 3 # how many attempts the gate may reject
worktree : false # set true to run in an isolated git worktree
defaults :
permission_mode : acceptEdits
steps :
- id : fix
role : > # persona — appended to Claude's system prompt
You are a careful maintainer. You make the smallest change that fixes
the problem, and you never weaken a test to make it pass.
prompt : |
Run this project's test suite. Diagnose and fix the root cause of any
failure. Re-run to confirm. Summarize what you changed and why.
- id : review
gate : true # ← the loop cannot succeed until this step says PASS
role : >
You are a skeptical senior engineer reviewing a change you did not
write. You trust nothing without evidence.
prompt : |
A previous agent claims to have fixed failing tests. Inspect the diff,
re-run the suite yourself, and check no test was weakened or deleted.
What happens when you run it
┌──────────────────────────────────────────────┐
│ iteration (≤ max) │
│ │
memory ──▶ │ step: fix ──▶ step: review (gate) ──┐ │
▲ │ ▲ │ │
│ │ └── reviewer feedback ◀── FAIL ─┤ │
│ └──────────────────────────────────────┼───────┘
│ │ PASS
└──────────────── run record ◀──────────────────┘
Each step is one headless Claude Code run ( claude -p ). Steps see the loop's memory , the outputs of earlier steps in the iteration, and — on retries — the gate's feedback .
A gate must end with VERDICT: PASS or VERDICT: FAIL . No verdict counts as FAIL: an unverified pass is not a pass.
On FAIL, the loop starts over with the reviewer's feedback injected into every prompt.
Every run appends a record to .loopflow/memory/<loop>.md — outcome, cost, and the final summary — which the next run reads.
Here's what that looks like in a real run — a release-check loop catching a debug artifact the fix step missed:
loopflow init gives you three loops designed to be stolen from:
test-and-fix — fixer + skeptical reviewer gate. The canonical write/verify pair.
debt-audit — a discovery loop. Maintains .loopflow/reports/debt-audit.md and uses memory to track what got fixed, what's new, and what keeps being ignored.
docs-sync — finds documentation that drifted from the code, fixes it in an isolated worktree, and a gate verifies every claim against the source.
Got a loop of your own? Contribute it to the cookbook — community loops live in loops/ .
LoopFlow deliberately ships no daemon. Use the scheduler you already have:
# cron (Linux/macOS) — audit debt every Monday at 9am
0 9 * * 1 cd /path/to/project && loopflow run debt-audit
# Windows Task Scheduler
schtasks /create /tn " debt-audit " /sc weekly /d MON /st 09:00 ^
/tr " cmd /c cd /d C:\path\to\project && loopflow run debt-audit "
CI works too — a GitHub Action that runs loopflow run docs-sync weekly and opens a PR from the kept worktree branch is ~20 lines.
A loop changes the work — it doesn't delete you from it. LoopFlow's design assumes three things stay true:
Verification is still on you. Gates catch the obvious failures, but --verbose and claude --resume <session-id> exist so you can read what the loop actually did. Read it.
Comprehension debt is real. The faster a loop ships code you didn't write, the faster the gap grows between what exists and what you understand. Memory files and kept worktrees are designed to be read by humans , not just machines.
The comfortable posture is the dangerous one. When the loop runs itself, it's tempting to stop having an opinion. Design the loop with judgment — then keep judging the output.
Build the loop. But build it like someone who intends to stay the engineer, not just the person who presses go.
Command
What it does
loopflow init [--force]
Scaffold .loopflow/ with starter loops
loopflow list
List loops with steps, gates, and budgets
loopflow validate [name]
Validate loop definitions (all by default)
loopflow run <name>
Run a loop
--dry-run
Print every composed prompt; invoke nothing
-i, --iterations <n>
Override budget.max_iterations
-b, --budget <usd>
Override budget.max_usd
-v, --verbose
Print full step outputs
Exit codes: 0 success · 1 loop failed (gate exhausted, budget, error) · 2 configuration error. Cron- and CI-friendly.
Everything the CLI does is exported:
import { loadLoop , runLoop } from "@loopflow/cli" ;
const loop = loadLoop ( process . cwd ( ) , "test-and-fix" ) ;
const result = await runLoop ( loop , { root : process . cwd ( ) } ) ;
console . log ( result . outcome , result . costUsd ) ;
Roadmap
loopflow daemon — built-in scheduler with cron expressions in loop.yaml
Parallel steps (fan-out across worktrees)
Structured gate verdicts via --json-schema
Loop run history & loopflow logs
Adapters for other headless agents (Codex CLI, …)
The most valuable contribution is a loop that solved a real problem for you — see CONTRIBUTING.md . Code contributions: the engine is ~600 lines of typed, tested TypeScript; npm test runs in under a second.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
