---
source: "https://github.com/zulman/peer-review"
hn_url: "https://news.ycombinator.com/item?id=48400663"
title: "CLI loop: A panel of AI agents that peer-reviews anything"
article_title: "GitHub - zulman/peer-review: A panel of AI agents that peer-reviews your document and hands back an improved version. · GitHub"
author: "gamescodedogs"
captured_at: "2026-06-04T16:11:07Z"
capture_tool: "hn-digest"
hn_id: 48400663
score: 1
comments: 0
posted_at: "2026-06-04T16:08:25Z"
tags:
  - hacker-news
  - translated
---

# CLI loop: A panel of AI agents that peer-reviews anything

- HN: [48400663](https://news.ycombinator.com/item?id=48400663)
- Source: [github.com](https://github.com/zulman/peer-review)
- Score: 1
- Comments: 0
- Posted: 2026-06-04T16:08:25Z

## Translation

タイトル: CLI ループ: あらゆるものをピアレビューする AI エージェントのパネル
記事のタイトル: GitHub - zulman/peer-review: ドキュメントをピアレビューし、改善されたバージョンを返す AI エージェントのパネル。 · GitHub
説明: ドキュメントをピアレビューし、改善されたバージョンを返す AI エージェントのパネル。 - ズルマン/査読

記事本文:
GitHub - zulman/peer-review: ドキュメントをピアレビューし、改善されたバージョンを返す AI エージェントのパネル。 · GitHub
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
ズルマン
/
査読
パブリックテンプレート
通知
通知設定を変更するにはサインインする必要があります
追加

アルナビゲーションオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット lib lib 役割 役割 実行 実行 スクリプト スクリプト .gitignore .gitignore README.md README.md ピアレビュー.sh ピアレビュー.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェントのパネルが文書をピアレビューし、改善されたバージョンを返します。
あなたはそれに草案、計画、質問、またはタスクを与えます。専門的なチーム
査読者がさまざまな角度から議論し、編集者が問題を解決します。
意見の相違がある場合、クリーンでよく考えられた文書が戻ってきます。さらに、完全な文書が返されます。
パネルがそこに到達するまでの記録。
設計文書、研究ノート、製品計画、ポリシーなど、あらゆるものに適用できます。
メモ、ストーリー、またはまだ計画が必要な生のタスク。
あなたのテキスト
│
▼
[ editor: Frame ] 提出物が何であるか、何をチェックすべきかを理解する
│
▼
┌─────── 1 回のレビューラウンド (解決するまで繰り返します) ─────────┐
│ 5 人のレビュー担当者が並行して実行されます: │
│ 楽観主義者、悲観主義者、懐疑論者、ググラー、常識人 │
│ │ │
│ ▼ │
│ [ 質問者 ] まだ未解決の質問はありますか？ → 続行 / 完了 │
│ │ (オープンの場合) 適切なレビュー担当者への的を絞ったフォローアップ │
│ ▼ │
│ [編集者: 解決する ] 議論を具体的な決定に変える │
━━━━━━━━━━━━━━━━━━━━━━━━━━━┘
│ (DONE、またはラウンド/安全キャップに到達したとき)
▼
[ エディター: ファイナライズ ] 決定事項に基づいて改善された文書を作成します
│
▼
最終.md + ダイアログ.md + 履歴.md
ストーリーテラーが端末内で進行状況をライブでナレーションします。

これは実行されます。
査読者
何をするのか
楽観主義者
前向きなシナリオと真の利点を見つけます。
悲観主義者
ネガティブなシナリオ、リスク、障害モードを見つけます。
懐疑的
証拠のない自明ではない主張を信用しません。
グーグラー
インターネットで同様の事例や過去の経験を検索します。
常識的な
その努力に価値があるかどうか疑問。忙しい仕事を減らします。ゴールまでの最短距離を保ちます。
サポートエージェント: エディター (フレーム → 解決 → ファイナライズ)、
質問者 (まだ回答されていないものがあるかどうかを検証)、および
ストーリーテラー（生ナレーション）。
bash に加えて、標準の Unix ツール ( awk 、 sed 、 grep )。
PATH 上のカーソル エージェント CLI と
CURSOR_API_KEY 。 (または、標準入力でプロンプトを読み取る CLI で --agent を指定します。
そしてマークダウンを出力します。)
--dry-run を使用して、モック出力を使用し、API 呼び出しを行わずにパイプラインを試します。
別のエージェントの使用 (Codex、opencode など)
デフォルトでは、システムはロールごとに Cursor のエージェント コマンド ラインを実行します。
CLI が標準入力のプロンプトを読み取り、マークダウンを標準出力に出力するだけの場合、
コードに触れずに切り替えることができます。
./peer-review.sh --agent " my-cli --some-flag " " ... "
# または:export PEER_REVIEW_AGENT="my-cli --some-flag"
CLI に別のフラグが必要な場合 (例: codex 、 opencode 、Claude Code など)、
統合ポイントは、 lib/common.sh の run_agent 関数です。欲しくない
手動で編集するには?以下のプロンプトをコピーして、コーディング エージェントに渡します —
<...> の部分に記入します。
私はこのフォルダーで「ピアレビュー」ツールを使用します。デフォルトでは、カーソルの「エージェント」が実行されます。
すべてのレビュー担当者ロールの CLI。代わりに <MY_CLI> を使用したいと考えています。
<MY_CLI> をシステムに接続してください。
- 唯一の統合ポイントは、`lib/common.sh` の `run_agent` 関数です。
`PEER_REVIEW_AGENT`が`agent`の場合にコマンドを構築するブロック。ブランチを追加する
<MY_CLI> の場合 (またはデフォルトにします)

）。
- 各ロールには、1 つのプロンプト文字列として完全な指示が与えられます。 <MY_CLI> は必須です
そのプロンプトを受け取り、完全に非対話型/ヘッドレスで実行し、読み取りを許可します。
「$REPO_ROOT」のワークスペースに保存し、結果のマークダウンのみを標準出力に出力します。
(チャット UI、進行状況スピナー、ストリーミング コントロール キャラクターはありません)。
- "$PEER_REVIEW_MODEL" が設定されている場合は、それをモデルとして渡します。
- レビュー ループ、役割、3 つの出力など、他のすべてを変更しないでください。
ファイル (final.md、dialog.md、history.md)。
- 完了したら確認します: `PEER_REVIEW_PROGRESS=0 ./peer-review.sh --dry-run "test"` まだ
これら 3 つのファイルを正確に生成し、<MY_CLI> を使用して実際に 1 回実行します。
私の CLI の詳細:
- 1 つのヘッドレス プロンプトを実行するコマンド: <例: `codex exec` / `opencode run`>
- プロンプトの受け取り方法 (標準入力 / `-p` フラグ / ファイル パス): <fill in>
- 最終的な回答のみを出力するフラグ: <存在する場合は記入>
- モデルを設定するフラグ (サポートされている場合): <存在する場合は入力>
使用法
./peer-review.sh [オプション] " 文書、計画、質問、またはタスク "
オプション
意味
--ワークスペースDIR
ディレクトリ エージェントは証拠を読み取ることができます (デフォルト: 現在のディレクトリ)。
--最大ラウンド数 N
レビューラウンドの上限 (デフォルト: 無制限)。
--モデルM
すべてのロールのエージェント モデル (例: Sonnet-4 )。
--エージェントCMD
各ロールの実行に使用される CLI (デフォルト: エージェント )。
--ドライラン
エージェントを生成する代わりにモック出力を使用します。
-h 、 --help
ヘルプを表示します。
環境オーバーライド: PEER_REVIEW_MAX_ROUNDS 、 PEER_REVIEW_MODEL 、
PEER_REVIEW_AGENT 、 PEER_REVIEW_AGENT_TIMEOUT 、 PEER_REVIEW_RUNS_DIR 、
PEER_REVIEW_PROGRESS (ライブ ステータス行を沈黙させるには 0 に設定)、
CURSOR_API_KEY 。
実行ごとに、runs/ の下に 3 つのファイルを含むセッション フォルダーが作成されます。
実行中に生成されるその他のものはすべて自動的にクリーンアップされます。
# 現在のディレクトリを証拠として使用して設計ドキュメントをレビューする
./peer-review.sh " の設計ドキュメント

オフライン同期」
# 特定のリポジトリに対するレビュー、最大 3 ラウンド
./peer-review.sh --workspace ~ /my-repo --max-rounds 3 " この分析メモを確認する "
# エージェントを呼び出さずにパイプラインを試してみる
PEER_REVIEW_PROGRESS=0 ./peer-review.sh --dry-run " アーキテクチャ レビューのインタビュー スクリプト "
プロジェクトのレイアウト
peer-review.sh エントリ ポイント: 引数を解析し、ループを実行し、結果を出力します。
ライブラリ/
common.sh 設定、セッション セットアップ、プロンプト レンダリング、エージェント ランナー
Orchestrator.sh のレビュー ループ: フレーム → ラウンド → ファイナライズ
Finalize.sh ファイナライズ ヘルパー + フォールバック ドキュメント ビルダー
Revid_writer.sh は、final.md を書き込み、dialog.md を構築し、セッションをクリーンアップします。
storyteller.sh ライブナレーション + History.md
progress.sh ターミナルステータス行
text.sh 小さなテキスト/役割ラベル ヘルパー
ロール/エージェントロールごとに 1 つの Markdown プロンプト
実行/セッション出力 (git 無視)
について
AI エージェントのパネルが文書をピアレビューし、改善されたバージョンを返します。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
貢献者
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A panel of AI agents that peer-reviews your document and hands back an improved version. - zulman/peer-review

GitHub - zulman/peer-review: A panel of AI agents that peer-reviews your document and hands back an improved version. · GitHub
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
zulman
/
peer-review
Public template
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits lib lib roles roles runs runs scripts scripts .gitignore .gitignore README.md README.md peer-review.sh peer-review.sh View all files Repository files navigation
A panel of AI agents that peer-reviews your document and hands back an improved version.
You give it a draft, a plan, a question, or a task. A team of specialized
reviewers argues over it from different angles, an editor resolves the
disagreements, and you get back a clean, well-thought-out document — plus a full
record of how the panel got there.
It works on anything : a design doc, a research note, a product plan, a policy
memo, a story, or a raw task that still needs a plan.
your text
│
▼
[ editor: frame ] understand what the submission is and what to check
│
▼
┌──────────────── one review round (repeats until settled) ───────────────┐
│ 5 reviewers run in parallel: │
│ optimist · pessimist · skeptic · googler · common-sense │
│ │ │
│ ▼ │
│ [ questioner ] are there still open questions? → CONTINUE / DONE │
│ │ (if open) targeted follow-up to the right reviewer │
│ ▼ │
│ [ editor: resolve ] turn the debate into concrete decisions │
└──────────────────────────────────────────────────────────────────────────┘
│ (when DONE, or round/safety cap reached)
▼
[ editor: finalize ] write the improved document from the decisions
│
▼
final.md + dialog.md + history.md
A storyteller narrates progress live in your terminal while this runs.
Reviewer
What it does
optimist
Finds positive scenarios and the real upside.
pessimist
Finds negative scenarios, risks, and failure modes.
skeptic
Distrusts any non-obvious claim made without proof.
googler
Searches the internet for similar cases and prior experience.
common-sense
Questions whether the effort is worth it; cuts busywork; keeps the shortest path to the goal.
Supporting agents: the editor (frame → resolve → finalize), the
questioner (validates whether anything is still unanswered), and the
storyteller (live narration).
bash , plus standard Unix tools ( awk , sed , grep ).
The Cursor agent CLI on your PATH and a
CURSOR_API_KEY . (Or point --agent at any CLI that reads a prompt on stdin
and prints Markdown.)
Use --dry-run to try the pipeline with mock output and no API calls.
Using a different agent (Codex, opencode, …)
By default the system runs Cursor's agent command line for every role.
If your CLI just reads a prompt on stdin and prints Markdown to stdout, you
can switch without touching the code:
./peer-review.sh --agent " my-cli --some-flag " " ... "
# or: export PEER_REVIEW_AGENT="my-cli --some-flag"
If your CLI needs different flags (e.g. codex , opencode , Claude Code, etc.),
the integration point is the run_agent function in lib/common.sh . Don't want
to edit it by hand? Copy the prompt below and give it to your coding agent —
fill in the <...> parts:
I use the `peer-review` tool in this folder. By default it runs Cursor's `agent`
CLI for every reviewer role. I want it to use <MY_CLI> instead.
Please wire <MY_CLI> into the system:
- The only integration point is the `run_agent` function in `lib/common.sh` — the
block that builds the command when `PEER_REVIEW_AGENT` is `agent`. Add a branch
for <MY_CLI> (or make it the default).
- Each role is given its full instructions as one prompt string. <MY_CLI> must
receive that prompt, run fully non-interactively / headless, be allowed to read
the workspace at "$REPO_ROOT", and print ONLY the resulting Markdown to stdout
(no chat UI, no progress spinners, no streaming control characters).
- If "$PEER_REVIEW_MODEL" is set, pass it as the model.
- Keep everything else unchanged: the review loop, the roles, and the three output
files (final.md, dialog.md, history.md).
- Verify when done: `PEER_REVIEW_PROGRESS=0 ./peer-review.sh --dry-run "test"` still
yields exactly those three files, then do one real run with <MY_CLI>.
My CLI details:
- Command to run one headless prompt: <e.g. `codex exec` / `opencode run`>
- How it takes the prompt (stdin / a `-p` flag / a file path): <fill in>
- Flag(s) to print only the final answer: <fill in, if any>
- Flag to set the model, if supported: <fill in, if any>
Usage
./peer-review.sh [options] " document, plan, question, or task "
Option
Meaning
--workspace DIR
Directory agents may read for evidence (default: current dir).
--max-rounds N
Cap on review rounds (default: unlimited).
--model M
Agent model for every role (e.g. sonnet-4 ).
--agent CMD
CLI used to run each role (default: agent ).
--dry-run
Use mock outputs instead of spawning agents.
-h , --help
Show help.
Environment overrides: PEER_REVIEW_MAX_ROUNDS , PEER_REVIEW_MODEL ,
PEER_REVIEW_AGENT , PEER_REVIEW_AGENT_TIMEOUT , PEER_REVIEW_RUNS_DIR ,
PEER_REVIEW_PROGRESS (set to 0 to silence the live status line),
CURSOR_API_KEY .
Each run creates a session folder under runs/ with exactly three files :
Everything else produced during the run is cleaned up automatically.
# Review a design doc using the current directory as evidence
./peer-review.sh " Design doc for offline sync "
# Review against a specific repo, capped at 3 rounds
./peer-review.sh --workspace ~ /my-repo --max-rounds 3 " Review this analytics memo "
# Try the pipeline without calling any agent
PEER_REVIEW_PROGRESS=0 ./peer-review.sh --dry-run " Interview script for architecture review "
Project layout
peer-review.sh entry point: parses args, runs the loop, prints results
lib/
common.sh config, session setup, prompt rendering, agent runner
orchestrator.sh the review loop: frame → rounds → finalize
finalize.sh finalize helpers + fallback document builder
revised_writer.sh writes final.md, builds dialog.md, cleans the session
storyteller.sh live narration + history.md
progress.sh the terminal status line
text.sh small text/role-label helpers
roles/ one Markdown prompt per agent role
runs/ session outputs (git-ignored)
About
A panel of AI agents that peer-reviews your document and hands back an improved version.
There was an error while loading. Please reload this page .
0
forks
Report repository
Contributors
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
