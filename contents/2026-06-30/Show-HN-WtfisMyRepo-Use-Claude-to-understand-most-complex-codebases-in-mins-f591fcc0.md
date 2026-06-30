---
source: "https://github.com/nandnijaiswal/wtfismyrepo"
hn_url: "https://news.ycombinator.com/item?id=48730022"
title: "Show HN: WtfisMyRepo – Use Claude to understand most complex codebases in mins"
article_title: "GitHub - nandnijaiswal/wtfismyrepo: Stop being lost in a codebase. CLI + Claude Code skill that builds an import graph, runs PageRank to find the files holding a repo together, scores fragility from git churn, and reads GitHub PRs/issues to show you exactly where to start. · GitHub"
author: "udit_50"
captured_at: "2026-06-30T09:03:40Z"
capture_tool: "hn-digest"
hn_id: 48730022
score: 2
comments: 1
posted_at: "2026-06-30T08:43:01Z"
tags:
  - hacker-news
  - translated
---

# Show HN: WtfisMyRepo – Use Claude to understand most complex codebases in mins

- HN: [48730022](https://news.ycombinator.com/item?id=48730022)
- Source: [github.com](https://github.com/nandnijaiswal/wtfismyrepo)
- Score: 2
- Comments: 1
- Posted: 2026-06-30T08:43:01Z

## Translation

タイトル: Show HN: WtfisMyRepo – クロードを使用して最も複雑なコードベースを数分で理解する
記事のタイトル: GitHub - nandnijaiswal/wtfismyrepo: コードベースで迷うのはやめましょう。 CLI + Claude Code スキルは、インポート グラフを構築し、PageRank を実行してリポジトリをまとめて保持しているファイルを検索し、git churn から脆弱性をスコア化し、GitHub PRs/issue を読み取ってどこから始めるべきかを正確に示します。 · GitHub
説明: コードベースで迷子になるのをやめます。 CLI + Claude Code スキルは、インポート グラフを構築し、PageRank を実行してリポジトリをまとめて保持しているファイルを検索し、git churn から脆弱性をスコア化し、GitHub PRs/issue を読み取ってどこから始めるべきかを正確に示します。 - nandnijaiswal/wtfismyrepo
HN テキスト: ストーリー タイム - 私のガールフレンドが私のプロジェクトに取り組みたがっていたので、最初にコードベースにアクセスできるようにしました。彼女はコードベース全体を非常に詳細に理解するためにこのクロード コード スキルを作成しました。超ショック、超素晴らしい。ぜひチェックしてみてください:)

記事本文:
GitHub - nandnijaiswal/wtfismyrepo: Stop being lost in a codebase. CLI + Claude Code スキルは、インポート グラフを構築し、PageRank を実行してリポジトリをまとめて保持しているファイルを検索し、git churn から脆弱性をスコア化し、GitHub PRs/issue を読み取ってどこから始めるべきかを正確に示します。 · GitHub
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
別のタブまたはウィンドウでアカウントを切り替えました。リロードして更新する

あなたのセッション。
アラートを閉じる
{{ メッセージ }}
ナンドニジャイシュワル
/
wtfismyrepo
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
3 コミット 3 コミット .github/ workflows .github/ workflows bin bin docs docs ドキュメント ドキュメント src src test test .gitignore .gitignore LICENSE LICENSE README.md README.md SKILL.md SKILL.md package.json package.json すべてのファイルを表示 リポジトリ ファイル ナビゲーション
wtfismyrepo — エージェントと人間のオンボーディング
あなたは 20 万行のコードベースに放り込まれたのです。何が起こっているのかわかりません。これはそれを修正します。
ほとんどの「コードの説明」ツールはコードを読み取って停止しますが、コードは存在するものを伝えるだけです。その理由と地雷は構造と歴史の中に存在します。 wtfismyrepo は、オンボーディングを読み取りの問題ではなくグラフの問題として扱います。インポート グラフを構築し、PageRank を実行してシステムを保持しているファイルを見つけ、git churn から脆弱性をスコア付けし、GitHub 履歴 (オープン PR = ホット ゾーン、問題 = 痛み) を読み取り、どこから始めるべきかを正確に示します。
これは、CLI/ライブラリ (決定論的エンジン) とクロード コード スキル (ガイド付きの部分ごとの説明) の両方として出荷されます。
リポジトリを初めて使用する場合、紛失した場合、レガシー コードを継承している場合、または「どこから始めればよいの?」と考えている場合は、いつでもこのツールに手を伸ばしてください。
並列: 1 日目、負け vs. wtfismyrepo
🟦 手動の方法
🟧 wtfismyrepo
リポジトリのクローンを作成します。ファイルツリーを開きます。 47 個の最上位フォルダー。
3 年前の README を読んでください。 main を grep します。 12 個のファイルを開いてスレッドを失い、11 個のファイルを閉じます。どのモジュールに負荷がかかっていて、どのモジュールが機能していないのかを推測してみてください。
どこに触れても安全なのかわかりません。今、あなたの下で何が変わっているのか分かりません。最初に何に取り組めばよいのか分かりません。

そこであなたは非難し、Slack で質問し、自分が 1 週間「調子を上げている」ことに誰も気づかないことを静かに望みます。
欠けているもの: あらゆる信号。リポジトリを一度も見たことがない場合、どのファイルも同様に重要であるように見えます。
npx github:nandnijaiswal/wtfismyrepo 。
数秒で次の結果が得られます。
スパイン - 構造的に中心的なファイル (インポート グラフ上の PageRank)。まずこれらをお読みください。
取り扱いには注意してください — 中心的で、頻繁に使用されるファイル (ユーザーを攻撃する可能性が最も高い)。
ホット ゾーン — 開いている PR の下にあるファイルは衝突する可能性があります。
最初の優れた問題 — GitHub から直接取得しました。
「最初の動き」 - 開くべき具体的なファイルと安全な最初のタスク。
47 フォルダーではなくシグナルです。昼食前にどこから始めればよいかわかります。
未編集の出力。このリポジトリで実行します。
━━ THE SPINE — ほとんどの中心ファイル (PageRank) ━━━━━━━━━━━━━━━
まずこれらを読んでください。すべては彼ら次第だ
██████████ src/install.mjs
█████████░ src/graph.mjs
█████████░ src/pagerank.mjs
████████░░ src/detect.mjs
█████░░░░░ src/git.mjs
━━ 歴史 — GitHub で何が起こっているのか ━━━━━━━━━━━━━━━━
🔥 オープン PR は 0 件 (ホット ゾーン - 人を驚かせることは避けてください)
🩹 3 つの未解決の問題 (どこが痛いのか)
#3 ディレクトリごとの責任の概要
#2 GitHub REST API 経由の gh-free 履歴
#1 Rustインポート解析を追加
🌱 良い最初の問題 — ここから始めてください:
#1 Rustインポート解析を追加
━━ あなたの最初の一歩 ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
1. bin/wtfismyrepo.mjs を開き、上から下に読んでください。
2. 次に、src/install.mjs をトレースします。これはスパインです。
3. 問題 #1:「Rust インポート解析の追加」を取り上げます。
取り扱い注意/壊れやすさパス

s (centrality × git-churn ) を点灯するにはコミット履歴が必要です。このような 1 コミット リポジトリではコミット履歴は空であり、成熟したコードベースでは本当に危険なファイルをランク付けします。
1 つのコマンド — スキルを Claude Code エージェントにインストールします (npm アカウントは必要ありません)。
npx github:nandnijaiswal/wtfismyrepo インストール
Claude Code を再起動し、次のように尋ねます。「なんと、このリポジトリは何ですか」 · 「ここに来たのは初めてですが、どこから始めればよいですか?」
カーソル/コーデックス/別のエージェントを使用していますか? @~/.claude/skills/wtfismyrepo/SKILL.md を CLAUDE.md / AGENTS.md / .cursorrules に追加します。完全なマトリックスは document/install.md にあります。
npx github:nandnijaiswal/wtfismyrepo 。 # 現在いるリポジトリを分析する
wtfismyrepo 。 --json # 機械可読 (スキルが消費するもの)
wtfismyrepo ../service --no-history --top 5 # オフライン、上位 5 つだけ
図書館:
「wtfismyrepo」からインポート {analyze , renderText } ;
const report =analyze(".", {history:true,top:10});
コンソール。 log (renderText (レポート) ) ;
// report.importance · report.fragility · report.hotZones · report.entryPoints · report.history
完全なリファレンス: document/api.md 。
ファイル ─► インポート パーサー ─► 指示されたインポート グラフ ─► PageRank ─┐
§─►重要度ランキング（背骨）
git ログ ──────────► ファイルごとのチャーン ──────────┤
└─► 脆弱性 = 中心性 × チャーン
gh CLI ──► オープン PR (ホット ゾーン) + 問題 (ペイン) ──────────► どこから始めるべきか
スキャン — git 追跡ファイル ( .gitignore を尊重) をリストし、解析可能なソース (JS/TS/JSX/Python/Go) を読み取ります。
グラフ — 実際の import / require / from ステートメントを解析し、内部ファイルに解決し、誰が誰に依存するグラフを構築します。
PageRank — ダンピング + ダングリング ノード処理によるパワー反復。上位 = s

松。
チャーン — ファイルごとの git ログの変更頻度。
脆弱性 — 正規化(中心性) × 正規化(チャーン) 。中心部と撹拌 = 取り扱いには注意してください。
歴史 — gh は、オープン PR (→ ホット ゾーン)、マージされた PR (→ 規約)、問題 (→ ペインと良好な優先問題) をプルします。
レポート — ターミナルまたは --json 、具体的な「最初の動き」で終わります。
エンジンはテストされ (PageRank の正確性、インポート解析、スコアリング、インストール)、ノード 18 / 20 / 22 上の CI で実行されます。詳細については、documentation/how-it-works.md を参照してください。
ページ
何が入っているのか
インストール
すべてのインストール パス - スキル、CLI、ライブラリ、エージェントごと
仕組み
パイプライン、アルゴリズム、スキルが重要な理由
CLIとAPI
CLI フラグ、ライブラリ タイプ、分析オブジェクト
また: SKILL.md — 実行可能なクロード コード スキル。
bin/wtfismyrepo.mjs CLI エントリ (分析、インストール、アンインストール)
ソース/
Index.mjs 公共ライブラリの表面
分析.mjs オーケストレーター
scan.mjs ファイルの検出 + ソースの読み取り
graph.mjsインポート解析+解決
pagerank.mjs PageRank アルゴリズム
git.mjs チャーン + 追跡されたファイル
fragility.mjs 重要度 + 脆弱性スコアリング
entrypoints.mjs エントリポイント検出
History.mjs gh PR/シグナルの発行
report.mjs ターミナル + JSON レンダリング
install.mjs スキルインストーラー
SKILL.md クロードコードスキル
ドキュメント/ドキュメント · テスト/単体テスト · .github/ CI
ロードマップと貢献
これは若いプロジェクトです。初期の問題や PR は特に歓迎されます。
Rust / Java / Ruby インポート解析
GitHub REST API 経由の gh -free 履歴
ディレクトリごとの「責任」の概要
npm に公開 ( npx wtfismyrepo 、github: プレフィックスなし)
オンボーディングレポートのHTML/JSONエクスポート
スパインのランキングが間違っていると思われるリポジトリ、または解析すべき言語を見つけましたか?問題をオープンしてください。これらは現時点で最も役立つ投稿です。
決定論的エンジンは src/ にあります。実行可能なSK

病気はSKILL.mdです。さあ、迷子になるのはやめてください。
コードベースで迷うのはもうやめましょう。 CLI + Claude Code スキルは、インポート グラフを構築し、PageRank を実行してリポジトリをまとめて保持しているファイルを検索し、git churn から脆弱性をスコア化し、GitHub PRs/issue を読み取ってどこから始めるべきかを正確に示します。
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

Stop being lost in a codebase. CLI + Claude Code skill that builds an import graph, runs PageRank to find the files holding a repo together, scores fragility from git churn, and reads GitHub PRs/issues to show you exactly where to start. - nandnijaiswal/wtfismyrepo

Story time - My girlfriend wanted to work on a project of mine so I gave her codebase access to first go thru.. she made this claude code skill to understand whole codebase in like super detailed. Supershocked, superamazing.. do check it out :)

GitHub - nandnijaiswal/wtfismyrepo: Stop being lost in a codebase. CLI + Claude Code skill that builds an import graph, runs PageRank to find the files holding a repo together, scores fragility from git churn, and reads GitHub PRs/issues to show you exactly where to start. · GitHub
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
nandnijaiswal
/
wtfismyrepo
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits .github/ workflows .github/ workflows bin bin docs docs documentation documentation src src test test .gitignore .gitignore LICENSE LICENSE README.md README.md SKILL.md SKILL.md package.json package.json View all files Repository files navigation
wtfismyrepo — onboarding for agents & humans
You just got dropped into a 200k-line codebase. You have no idea what's going on. This fixes that.
Most "explain my code" tools read the code and stop — but code only tells you what exists. The why , and the landmines , live in the structure and the history . wtfismyrepo treats onboarding as a graph problem, not a reading problem — it builds an import graph, runs PageRank to find the files that hold the system together, scores fragility from git churn, and reads the GitHub history (open PRs = hot zones, issues = pain) to tell you exactly where to start .
It ships as both a CLI/library (the deterministic engine) and a Claude Code skill (the guided, part-by-part explanation).
Reach for it whenever you're new to a repo , lost , inheriting legacy code , or thinking "where do I even start with this?"
Side-by-side: day one, lost vs. wtfismyrepo
🟦 The manual way
🟧 wtfismyrepo
Clone the repo. Open the file tree. 47 top-level folders.
Read a README that's three years stale. Grep around for main . Open twelve files, lose the thread, close eleven. Try to guess which module is load-bearing and which is dead.
No idea what's safe to touch. No idea what's changing under you right now. No idea what to work on first.
So you git blame , ask in Slack, and quietly hope nobody notices you've been "ramping up" for a week.
What's missing: any signal . Every file looks equally important when you've never seen the repo before.
npx github:nandnijaiswal/wtfismyrepo .
In seconds you get:
The spine — the structurally central files (PageRank over the import graph). Read these first.
Handle-with-care — files that are central and high-churn (most likely to bite you).
Hot zones — files under open PRs you'd collide with.
Good first issues — pulled straight from GitHub.
"Your first move" — a concrete file to open and a safe first task.
Signal, not 47 folders. You know where to start before lunch.
Unedited output, run on this very repo:
━━ THE SPINE — most central files (PageRank) ━━━━━━━━━━━━━━━
read these first; everything depends on them
██████████ src/install.mjs
█████████░ src/graph.mjs
█████████░ src/pagerank.mjs
████████░░ src/detect.mjs
█████░░░░░ src/git.mjs
━━ THE HISTORY — what's happening on GitHub ━━━━━━━━━━━━━━━━
🔥 0 open PRs (hot zones — avoid surprising people)
🩹 3 open issues (where it hurts)
#3 Per-directory responsibility summaries
#2 gh-free history via GitHub REST API
#1 Add Rust import parsing
🌱 good first issues — START HERE:
#1 Add Rust import parsing
━━ YOUR FIRST MOVE ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
1. Open bin/wtfismyrepo.mjs and read top-to-bottom.
2. Then trace into src/install.mjs — it's the spine.
3. Pick up issue #1: "Add Rust import parsing".
The Handle-with-care / fragility pass ( centrality × git-churn ) needs commit history to light up — it's empty on a one-commit repo like this one, and ranks the genuinely risky files on any mature codebase.
One command — installs the skill into your Claude Code agent (no npm account needed):
npx github:nandnijaiswal/wtfismyrepo install
Restart Claude Code, then just ask: "wtf is this repo" · "I'm new here, where do I start?"
Using Cursor / Codex / another agent? Add @~/.claude/skills/wtfismyrepo/SKILL.md to your CLAUDE.md / AGENTS.md / .cursorrules . Full matrix in documentation/install.md .
npx github:nandnijaiswal/wtfismyrepo . # analyze the repo you're in
wtfismyrepo . --json # machine-readable (what the skill consumes)
wtfismyrepo ../service --no-history --top 5 # offline, just the top 5
Library:
import { analyze , renderText } from "wtfismyrepo" ;
const report = analyze ( "." , { history : true , top : 10 } ) ;
console . log ( renderText ( report ) ) ;
// report.importance · report.fragility · report.hotZones · report.entryPoints · report.history
Full reference: documentation/api.md .
files ─► import parser ─► directed import graph ─► PageRank ──┐
├─► importance ranking (the spine)
git log ───────────────► per-file churn ───────────────────┤
└─► fragility = centrality × churn
gh CLI ──► open PRs (hot zones) + issues (pain) ──────────────► where to start
Scan — list git-tracked files (respects .gitignore ), read parseable sources (JS/TS/JSX/Python/Go).
Graph — parse real import / require / from statements, resolve to internal files, build a who-depends-on-whom graph.
PageRank — power iteration with damping + dangling-node handling. High rank = the spine.
Churn — git log change-frequency per file.
Fragility — normalize(centrality) × normalize(churn) . Central and churned = handle with care.
History — gh pulls open PRs (→ hot zones), merged PRs (→ conventions), issues (→ pain & good-first-issues).
Report — terminal or --json , ending in a concrete "Your first move."
The engine is tested (PageRank correctness, import parsing, scoring, install) and runs in CI on Node 18 / 20 / 22. Deep dive: documentation/how-it-works.md .
Page
What's in it
Install
Every install path — skill, CLI, library, per-agent
How it works
The pipeline, the algorithm, why the skill matters
CLI & API
CLI flags, library types, the Analysis object
Also: SKILL.md — the runnable Claude Code skill.
bin/wtfismyrepo.mjs CLI entry (analyze · install · uninstall)
src/
index.mjs public library surface
analyze.mjs orchestrator
scan.mjs file discovery + source reading
graph.mjs import parsing + resolution
pagerank.mjs the PageRank algorithm
git.mjs churn + tracked files
fragility.mjs importance + fragility scoring
entrypoints.mjs entry-point detection
history.mjs gh PR/issue signals
report.mjs terminal + JSON rendering
install.mjs skill installer
SKILL.md the Claude Code skill
documentation/ docs · test/ unit tests · .github/ CI
Roadmap & contributing
This is a young project — early issues and PRs are especially welcome:
Rust / Java / Ruby import parsing
gh -free history via the GitHub REST API
Per-directory "responsibility" summaries
Published to npm ( npx wtfismyrepo , no github: prefix)
HTML/JSON export of the onboarding report
Found a repo where the spine ranking looks wrong, or a language it should parse? Open an issue — those are the most useful contributions right now.
The deterministic engine lives in src/ ; the runnable skill is SKILL.md . Go forth and stop being lost.
Stop being lost in a codebase. CLI + Claude Code skill that builds an import graph, runs PageRank to find the files holding a repo together, scores fragility from git churn, and reads GitHub PRs/issues to show you exactly where to start.
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
