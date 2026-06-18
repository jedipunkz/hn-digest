---
source: "https://github.com/emiliolugo/clawmark"
hn_url: "https://news.ycombinator.com/item?id=48585493"
title: "clawmark: open-source CLAUDE.md A/B Testing CLI tool"
article_title: "GitHub - emiliolugo/clawmark: clawmark run evaluates 2 CLAUDE.md files on standardized benchmarks · GitHub"
author: "emiliolugo"
captured_at: "2026-06-18T16:15:55Z"
capture_tool: "hn-digest"
hn_id: 48585493
score: 2
comments: 1
posted_at: "2026-06-18T14:00:00Z"
tags:
  - hacker-news
  - translated
---

# clawmark: open-source CLAUDE.md A/B Testing CLI tool

- HN: [48585493](https://news.ycombinator.com/item?id=48585493)
- Source: [github.com](https://github.com/emiliolugo/clawmark)
- Score: 2
- Comments: 1
- Posted: 2026-06-18T14:00:00Z

## Translation

タイトル: clawmark: オープンソース CLAUDE.md A/B テスト CLI ツール
記事のタイトル: GitHub - emiliolugo/clawmark: clawmark 実行により、標準化されたベンチマークで 2 つの CLAUDE.md ファイルが評価される · GitHub
説明: clawmark の実行は、標準化されたベンチマークで 2 つの CLAUDE.md ファイルを評価します - emiliolugo/clawmark

記事本文:
GitHub - emiliolugo/clawmark: clawmark run evaluates 2 CLAUDE.md files on standardized benchmarks · GitHub
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
emiliolugo
/
爪痕
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メートル

ain Branches Tags Go to file Code Open more actions menu Folders and files
4 コミット 4 コミット .github/ workflows .github/ workflows data data src src .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md すべてのファイルを表示 リポジトリ ファイル ナビゲーション
clawmark is a local Rust CLI for answering one focused question:
Which of these two CLAUDE.md files performs better on a small SWE-bench Lite smoke set?
v1 compares exactly two local variant files against five bundled SWE-bench Lite tasks. It runs Claude locally, evaluates the generated patches with the official SWE-bench harness in Docker, and writes a simple A/B report.
clawmark doctor checks local prerequisites.
clawmark run evaluates variant A and variant B on the same five tasks.
clawmark report reads an existing output directory and prints the A/B summary again.
There is no config file, web UI, remote execution, retries, resume, progress UI, repeated trials, or full 300-task SWE-bench run in v1.
Install these yourself before running clawmark :
cargo run -- doctor
doctor prints a status table and exits non-zero if a required check fails. A missing SWE-bench Docker image is only a warning;最初の評価で引く可能性があります。
Create two variant files somewhere inside the current working directory:
mkdir -p のバリアント
$EDITOR variants/a.md
$EDITOR バリアント/b.md
A/B 煙ベンチマークを実行します。
貨物の実行 -- 実行 \
--a variants/a.md \
--b バリアント/b.md \
--model sonnet \
--タイムアウト秒 300 \
--アウトアウト
This performs:
2 variants x 5 tasks x 1 trial = 10 Claude invocations
run は新しい出力ディレクトリを作成します。 It fails if --out already exists, so use a new directory for each run.
Print the report from existing output:
カーゴラン -- レポート --out out
ビルド後

ng a release binary, the same commands can be run as:
cargo build --release
./ターゲット/リリース/爪痕ドクター
./target/release/clawmark run --a variants/a.md --b variants/b.md --model sonnet --timeout-secs 300 --out out
./target/release/clawmark レポート --out out
Runtime And Budget Warnings
v1 is intentionally minimal and does not enforce a turn limit, token budget, retry policy, or per-task cost cap. --timeout-secs is only a wall-clock timeout around each Claude invocation. A broad CLAUDE.md can spend the full timeout exploring the repo, installing dependencies, or running tests without producing a patch.
For first e2e runs, use strict benchmark-oriented variants:
あなたは自動ベンチマーク内で実行しています。 Make the smallest code change that addresses the issue.
ルール:
- 完全なテスト スイートを実行しないでください。
- 問題に必要なファイルのみを検査します。
- If you run tests, run at most one targeted test command.
- Do not spend time on formatting, docs, or unrelated cleanup.
- もっともらしい最小限のパッチが作成されたら、停止します。
推奨される開始設定:
有界スモークテストには --timeout-secs 600 を使用します。
Use --timeout-secs 1800 only when you want to give Claude enough time to solve harder tasks.
試行ごとに新しい --out ディレクトリを使用してください。
Run cargo run -- doctor first so failures happen before any Claude calls.
立ち去る前に最初のタスクを見てください。 if it reaches the timeout with an empty patch, tighten your variant instructions before running all 10 invocations.
Budget expectation varies heavily by model behavior. v1 スモーク実行は 10 回の Claude 呼び出しを実行するため、オープンエンドのバリアントは、パッチ中心の短いバリアントよりも大幅に多くの時間と使用量クォータを消費する可能性があります。
各タスクとバリアントについて、clawmark :
Clones the SWE-bench target repository into a temporary workspace.
小切手

タスクのベースコミットをksアウトします。
選択したバリアント ファイルを CLAUDE.md としてリポジトリ ルートに書き込みます。
Invokes Claude with the task problem statement.
git diff HEAD をモデルパッチとしてキャプチャします。
クロードは次のようにローカルで呼び出されます。
claude -p --output-format json --dangerously-skip-permissions --model < model > --add-dir < workspace > -- < problem_statement >
すべての予測がバリアントに対して書き込まれた後、clawmark は SWE ベンチ ハーネスを A に対して 1 回、B に対して 1 回呼び出します。レポートでは、ハーネスのsolved_ids 配列を真実のソースとして扱います。
clawmark doctor
Docker、Claude CLI、Claude 認証、Python、swebench 、SWE ベンチ ハーネス CLI、git、Docker Hub レジストリの到達可能性、および SWE ベンチ Docker イメージが既に存在するかどうかを確認します。
clawmark run --a < パス > --b < パス > --model < モデル > --timeout-secs < 秒 > --out < ディレクトリ >
固定の 5 タスク A/B ベンチマークを実行します。 --timeout-secs は 1 ～ 86400 の範囲にする必要があります。 it applies to each Claude invocation and is also passed to the SWE-bench harness.
爪跡レポート --out < ディレクトリ >
既存のハーネス出力を読み取り、解決されたカウント、A の勝利、B の勝利、両方が解決された同点、および両方が失敗した同点を出力してから、 report.json を書き込みます。
--a と --b は存在し、シンボリックリンク解決後の通常のファイルである必要があります。
どちらのバリアント パスも、プロセスの現在の作業ディレクトリ内に存在する必要があります。
A と B は異なる正規ファイルに解決される必要があります。
--model は空ではない文字列である必要があり、1 つの引数として claude --model に渡されます。
run --out requires an existing parent directory and a destination that does not already exist.
report --out requires an existing v1 output directory with harness results.
バリアント ファイル名は CLAUDE.md である必要はありません。それらの内容は、各一時タスク ワークスペース内に CLAUDE.md として挿入されます。
アウト/
run_records.jsonl
予測/
a.jsonl

b.jsonl
ハーネス/
a.json
b.json
レポート.json
run_records.jsonl stores one record per variant/task attempt. predictions/a.jsonl and predictions/b.jsonl are the SWE-bench harness inputs. harness/a.json および harness/b.json は、生の SWE ベンチ概要ファイルの安定したコピーです。 report.json には、最終的な A/B 集計レポートが保存されます。
clawmark が所有するすべてのレコードには、 schema_version: 1 が含まれます。
タスクごとのエラーのほとんどが記録され、そのタスクの空のパッチを使用して実行が続行されます。
モデルが利用できない、またはレート制限エラー
クロード認証に失敗すると、実行全体が中止されます。ハーネスの障害は、report.json が書き込まれる前に中止されますが、すでに書き込まれた予測は検査のために出力ディレクトリに残ります。
clawmark is a local developer tool for user-controlled inputs.
クロードはコンテナ内ではなくホスト上で実行されます。 The command uses --dangerously-skip-permissions , and variant instructions are not OS-sandboxed. v1 では、信頼できない CLAUDE.md バリアント、信頼できないベンチマーク データ、または信頼できないプロンプトを実行しないでください。
SWE ベンチ テストの実行は、公式ハーネスを介して Docker 内で実行されます。モデルで生成されたパッチはハーネスによって評価されます。 clawmark はホスト上でパッチを直接実行しません。
clawmark は、サブプロセス argv 配列を使用することでシェル インジェクションを軽減し、パスを正規化して現在の作業ディレクトリと照合することでバリアント パス トラバーサルを軽減し、アトミック ファイル書き込みによる部分的な書き込み破損を軽減します。 It does not prevent malicious model behavior from accessing host files, environment variables, network resources, or other local credentials available to the process.
実行するインスタンスがありません。 — The harness filters out empty predictions. This almost always means Claude produced no patch (often from too small a --timeout-secs ).十分なタイムアウト (例: --timeout-secs 600 ) と新しい --o を使用します。

ut directory, then inspect run_records.jsonl to confirm patches are non-empty.
lookup registry-1.docker.io: no such host / Docker image pull errors — Docker cannot resolve Docker Hub, so the harness cannot pull SWE-bench images. clawmark Doctor は、「Docker レジストリに到達可能」チェックでこれにフラグを立てます。ネットワーク/VPN および Docker デスクトップの DNS 設定を修正してから、再試行してください。これは環境の問題であり、爪跡のバグではありません。
Hugging Face 404 lines during dataset load — These are normal probes the datasets library makes for optional files (e.g. SWE-bench_Lite.py , dataset_infos.json ).これらはエラーではなく、実行には影響しません。
clawmark はテレメトリや使用状況データを送信しません。クロードおよび SWE ベンチは、通常の動作の一部として独自のネットワーク アクティビティを実行する場合があります。
プルリクエストは大歓迎です。大きな変更や新機能については、問題をオープンしてください
まず変更について説明します。 CONTRIBUTING.md を参照してください。
MIT ライセンスに基づいてライセンスされています。
clawmark の実行は、標準化されたベンチマークで 2 つの CLAUDE.md ファイルを評価します
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

clawmark run evaluates 2 CLAUDE.md files on standardized benchmarks - emiliolugo/clawmark

GitHub - emiliolugo/clawmark: clawmark run evaluates 2 CLAUDE.md files on standardized benchmarks · GitHub
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
emiliolugo
/
clawmark
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits .github/ workflows .github/ workflows data data src src .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md View all files Repository files navigation
clawmark is a local Rust CLI for answering one focused question:
Which of these two CLAUDE.md files performs better on a small SWE-bench Lite smoke set?
v1 compares exactly two local variant files against five bundled SWE-bench Lite tasks. It runs Claude locally, evaluates the generated patches with the official SWE-bench harness in Docker, and writes a simple A/B report.
clawmark doctor checks local prerequisites.
clawmark run evaluates variant A and variant B on the same five tasks.
clawmark report reads an existing output directory and prints the A/B summary again.
There is no config file, web UI, remote execution, retries, resume, progress UI, repeated trials, or full 300-task SWE-bench run in v1.
Install these yourself before running clawmark :
cargo run -- doctor
doctor prints a status table and exits non-zero if a required check fails. A missing SWE-bench Docker image is only a warning; the first evaluation may pull it.
Create two variant files somewhere inside the current working directory:
mkdir -p variants
$EDITOR variants/a.md
$EDITOR variants/b.md
Run the A/B smoke benchmark:
cargo run -- run \
--a variants/a.md \
--b variants/b.md \
--model sonnet \
--timeout-secs 300 \
--out out
This performs:
2 variants x 5 tasks x 1 trial = 10 Claude invocations
run creates a fresh output directory. It fails if --out already exists, so use a new directory for each run.
Print the report from existing output:
cargo run -- report --out out
After building a release binary, the same commands can be run as:
cargo build --release
./target/release/clawmark doctor
./target/release/clawmark run --a variants/a.md --b variants/b.md --model sonnet --timeout-secs 300 --out out
./target/release/clawmark report --out out
Runtime And Budget Warnings
v1 is intentionally minimal and does not enforce a turn limit, token budget, retry policy, or per-task cost cap. --timeout-secs is only a wall-clock timeout around each Claude invocation. A broad CLAUDE.md can spend the full timeout exploring the repo, installing dependencies, or running tests without producing a patch.
For first e2e runs, use strict benchmark-oriented variants:
You are running inside an automated benchmark. Make the smallest code change that addresses the issue.
Rules:
- Do not run the full test suite.
- Only inspect files needed for the issue.
- If you run tests, run at most one targeted test command.
- Do not spend time on formatting, docs, or unrelated cleanup.
- When a plausible minimal patch is made, stop.
Recommended starting settings:
Use --timeout-secs 600 for a bounded smoke test.
Use --timeout-secs 1800 only when you want to give Claude enough time to solve harder tasks.
Use a fresh --out directory for every attempt.
Run cargo run -- doctor first so failures happen before any Claude calls.
Watch the first task before walking away; if it reaches the timeout with an empty patch, tighten your variant instructions before running all 10 invocations.
Budget expectation varies heavily by model behavior. The v1 smoke run performs 10 Claude invocations, so open-ended variants can consume materially more time and usage quota than short, patch-focused variants.
For each task and variant, clawmark :
Clones the SWE-bench target repository into a temporary workspace.
Checks out the task's base commit.
Writes the selected variant file as CLAUDE.md at the repo root.
Invokes Claude with the task problem statement.
Captures git diff HEAD as the model patch.
Claude is invoked locally with:
claude -p --output-format json --dangerously-skip-permissions --model < model > --add-dir < workspace > -- < problem_statement >
After all predictions are written for a variant, clawmark invokes the SWE-bench harness once for A and once for B. The report treats the harness resolved_ids arrays as the source of truth.
clawmark doctor
Checks Docker, Claude CLI, Claude authentication, Python, swebench , the SWE-bench harness CLI, git, Docker Hub registry reachability, and whether the SWE-bench Docker image is already present.
clawmark run --a < path > --b < path > --model < model > --timeout-secs < seconds > --out < dir >
Runs the fixed five-task A/B benchmark. --timeout-secs must be between 1 and 86400 ; it applies to each Claude invocation and is also passed to the SWE-bench harness.
clawmark report --out < dir >
Reads existing harness output, prints resolved counts, A wins, B wins, both-resolved ties, and both-failed ties, then writes report.json .
--a and --b must exist and be regular files after symlink resolution.
Both variant paths must be inside the process current working directory.
A and B must resolve to different canonical files.
--model must be a non-empty string and is passed as one argument to claude --model .
run --out requires an existing parent directory and a destination that does not already exist.
report --out requires an existing v1 output directory with harness results.
Variant filenames do not need to be CLAUDE.md ; their contents are injected as CLAUDE.md inside each temporary task workspace.
out/
run_records.jsonl
predictions/
a.jsonl
b.jsonl
harness/
a.json
b.json
report.json
run_records.jsonl stores one record per variant/task attempt. predictions/a.jsonl and predictions/b.jsonl are the SWE-bench harness inputs. harness/a.json and harness/b.json are stable copies of the raw SWE-bench summary files. report.json stores the final A/B aggregate report.
All clawmark-owned records include schema_version: 1 .
Most per-task failures are recorded and the run continues with an empty patch for that task:
model unavailable or rate limit errors
Claude authentication failures abort the whole run. Harness failures abort before report.json is written, but already-written predictions remain in the output directory for inspection.
clawmark is a local developer tool for user-controlled inputs.
Claude runs on the host, not in a container. The command uses --dangerously-skip-permissions , and variant instructions are not OS-sandboxed. Do not run untrusted CLAUDE.md variants, untrusted benchmark data, or untrusted prompts through v1.
SWE-bench test execution runs inside Docker through the official harness. The model-generated patch is evaluated by the harness; clawmark does not execute the patch directly on the host.
clawmark mitigates shell injection by using subprocess argv arrays, variant path traversal by canonicalizing and checking paths against the current working directory, and partial write corruption with atomic file writes. It does not prevent malicious model behavior from accessing host files, environment variables, network resources, or other local credentials available to the process.
No instances to run. — The harness filters out empty predictions. This almost always means Claude produced no patch (often from too small a --timeout-secs ). Use a generous timeout (e.g. --timeout-secs 600 ) and a fresh --out directory, then inspect run_records.jsonl to confirm patches are non-empty.
lookup registry-1.docker.io: no such host / Docker image pull errors — Docker cannot resolve Docker Hub, so the harness cannot pull SWE-bench images. clawmark doctor flags this with the "Docker registry reachable" check. Fix your network/VPN and Docker Desktop DNS settings, then retry. This is an environment issue, not a clawmark bug.
Hugging Face 404 lines during dataset load — These are normal probes the datasets library makes for optional files (e.g. SWE-bench_Lite.py , dataset_infos.json ). They are not errors and do not affect the run.
clawmark does not send telemetry or usage data. Claude and SWE-bench may perform their own network activity as part of their normal operation.
Pull requests are welcome. For major changes or new features, open an issue
first to discuss the change — see CONTRIBUTING.md .
Licensed under the MIT License .
clawmark run evaluates 2 CLAUDE.md files on standardized benchmarks
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
