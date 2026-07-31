---
source: "https://github.com/stalebrainlabs/stalebrain"
hn_url: "https://news.ycombinator.com/item?id=49122715"
title: "Show HN: Provenance and decay for AI agent memory"
article_title: "GitHub - stalebrainlabs/stalebrain: Provenance and decay for AI agent memory. Audits CLAUDE.md, AGENTS.md, .cursorrules, GEMINI.md and 18+ agent memory files against the live repo: typed claims, cited commits, half-life decay, token meter, approve-only fixes. Works with any agent. · GitHub"
author: "souravroy78"
captured_at: "2026-07-31T13:41:02Z"
capture_tool: "hn-digest"
hn_id: 49122715
score: 1
comments: 0
posted_at: "2026-07-31T13:13:53Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Provenance and decay for AI agent memory

- HN: [49122715](https://news.ycombinator.com/item?id=49122715)
- Source: [github.com](https://github.com/stalebrainlabs/stalebrain)
- Score: 1
- Comments: 0
- Posted: 2026-07-31T13:13:53Z

## Translation

タイトル: HN を表示: AI エージェントの記憶の来歴と衰退
記事のタイトル: GitHub - stalbrainlabs/stalbrain: AI エージェント メモリの来歴と衰退。 CLAUDE.md、AGENTS.md、.cursorrules、GEMINI.md、および 18 個以上のエージェント メモリ ファイルをライブ リポジトリに対して監査します: 入力されたクレーム、引用されたコミット、半減期の減衰、トークン メーター、承認のみの修正。どのエージェントでも動作します。 · GitHub
説明: AI エージェントの記憶の起源と衰退。 CLAUDE.md、AGENTS.md、.cursorrules、GEMINI.md、および 18 個以上のエージェント メモリ ファイルをライブ リポジトリに対して監査します: 入力されたクレーム、引用されたコミット、半減期の減衰、トークン メーター、承認のみの修正。どのエージェントでも動作します。 - ステイルブレインラボ/ステイルブレイン

記事本文:
GitHub - stalbrainlabs/stalbrain: AI エージェントの記憶の来歴と衰退。 CLAUDE.md、AGENTS.md、.cursorrules、GEMINI.md、および 18 個以上のエージェント メモリ ファイルをライブ リポジトリに対して監査します: 入力されたクレーム、引用されたコミット、半減期の減衰、トークン メーター、承認のみの修正。どのエージェントでも動作します。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
アカウントを切り替えました o

別のタブまたはウィンドウ。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
スタイルブレインラボ
/
鈍い脳
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
11 コミット 11 コミット .github/ workflows .github/ workflows アセット アセット リファレンス リファレンス src/ stalbrain src/ stalebrrain .gitignore .gitignore LICENSE LICENSE PORTABLE.md PORTABLE.md README.md README.md SKILL.md SKILL.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェントのメモリ ファイルは嘘をついています。 stalbrain はあらゆる嘘を見つけ、コミットで証明し、修正を渡します。
インストール・
デモ・
機能 ·
比較する方法 ·
トラブルシューティング ·
ライセンス
AI コーディング アシスタントに /stale-brain と入力すると、リポジトリ内のすべてのエージェント メモリ ファイル (CLAUDE.md、AGENTS.md、.cursorrules、GEMINI.md、Copilot 命令、13 エージェントにわたる 18 以上の場所) が、現在存在するコードに対して監査されます。チェックアウトした請求には日付のスタンプが押されます。評決が得られない主張、それを証明するコミット、すぐに適用できる修正。
リンターではありません。信頼モデル: 1 回の監査の後、エージェントの記憶にあるすべての主張には、年齢、検証日、引用された証拠が記録されます。
前提条件: git 履歴のあるリポジトリ。それでおしまい。
uv または pip を使用する (推奨、クロード コードまたは任意の CLI 内で動作します)
uv ツール install stalbrain # または: pip install stalbrain / pipx install stalbrain
stalbrain install # ユーザーレベル: すべてのリポジトリ、Claude Code CLI、デスクトップ アプリ
stalbrain インストール --project 。 # またはこのリポジトリだけ
永続的なものをインストールせずにワンショットで実行:
uvxステイルブレインのインストール
git だけで
# プロジェクトレベル
git clone https://github.com/staleebrainlabs/stalebrrain .claude/skills/stale-

脳
# またはユーザーレベル (すべてのリポジトリで利用可能)
git clone https://github.com/staleebrainlabs/stalebrrain ~/.claude/skills/stale-brain
次に、 /stale-brain と入力するか、「エージェントのメモリを監査してください」と言います。
stalbrain インストールは、Claude Code デスクトップ アプリをカバーします (同じユーザーレベルのスキル フォルダーを読み取ります)。他のすべてのデスクトップ チャット (ChatGPT、Gemini、Claude.ai) については、stalebrrainportable を実行して単一ファイル プロトコルを印刷するか、PORTABLE.md をコピーしてチャットに貼り付けます。
その他のアシスタント (Cursor、Copilot、Gemini CLI、ChatGPT、ローカル モデル)
PORTABLE.md をチャットに貼り付けるか、リポジトリにドロップして「PORTABLE.md で stale-brain プロトコルを実行する」と言います。これは完全に自己完結型で、正常に機能を低下させます。ツールを持たないアシスタントはコマンド出力 (1 つのブロックにまとめて) を貼り付けるように要求し、ファイル アクセス権を持たないアシスタントは適用できるように差分を出力します。
常時オンのトリップワイヤー (オプション、任意のエージェント)
常時ロードされるメモリ ファイル (CLAUDE.md、AGENTS.md、または rules) に 1 行を追加します。
このファイル内の指示が観察された現実と矛盾する場合は、黙って従うのではなく、その旨を 1 行で述べ (⚡ 古い脳)、古い脳の監査を提案してください。
実際の動作を確認してください
検証は機械的に行われます: glob、grep、読み取り専用 git。埋め込み、サーバー、データベース、ネットワークはありません。マシンからは何も出ず、何も実行されません。スクリプトは、実行することではなく、その定義 (スクリプト ブロック、Makefile、CI) によって検証されます。
エージェントは「ルールを無視」し続けているのか？半分の場合、ルールは守れません。彼らは、移動されたパス、名前が変更されたスクリプト、3 月に交換されたパッケージ マネージャーに名前を付けています。エージェントに矛盾した記憶を与えても、90% 正解することはできません。セッションごとに、トークンコストを支払うことで、自信を持って間違います。
能力
どうやって
クレーム抽出
すべての文は入力されたクレームになります: PATH

、スクリプト、シンボル、DEP、事実、所有者、規約。スタイルに関する意見は無視され、決して判断されません。
検証
ライブ リポジトリに対するタイプごとのレシピ。パスが見つからない場合は、git Archeology が名前変更と新しい宛先を見つけます。間違った所有者は CODEOWNERS および git shortlog に対してチェックされます。
自信の低下
タイプごとの半減期 (パス 30 日、規則 120 日)。半減期を過ぎると、事実は再検証されるか、仮説に格下げされます。
来歴スタンプ
<!-- stale-brain: 2026-07-31 に検証済み --> 、レンダリングされたマークダウンでは表示されず、永久に grepable になります。スタンプにより、再監査が段階的に行われます。
矛盾の証拠
決して「これは間違っている」と思わないでください。常に「a1b2c3d 以来間違っています。差分は次のとおりです」。
ファイル間の競合
CLAUDE.md は、yarn と言い、.cursorrules は、どちらが正しいか誰も分からない場合でも、npm: フラグが立てられると言います。
トークンメーター
セッションあたりのメモリのコストと、モデルを積極的に誤解させるメモリの割合。
中間タスクのトリップワイヤー
エージェントはタスクの途中で現実と矛盾する指示に気づくと、黙って従うのではなく一行でその旨を言います。
承認のみの修正
すべての評決は説明されています。すべての編集はあなたの「はい」を待ちます。非対話型の実行では何も適用されません。
扱うファイル
エージェント
ファイル
クロード・コード
CLAUDE.md (ルートおよびネストされた)、CLAUDE.local.md、.claude/**/*.md
コーデックス/クロスツール
AGENTS.md (ルートおよびネストされた)
カーソル
.cursorrules、.cursor/rules/**/*.mdc
GitHub コパイロット
.github/copilot-instructions.md、.github/instructions/*.instructions.md
ジェミニ CLI
ジェミニ.md
ウィンドサーフィン
.windsurfrules、.windsurf/rules/**
クライン
.clinerules (ファイルまたはディレクトリ)
Rooコード
.roo/ルール/**
補助者
CONVENTIONS.md (.aider.conf.yml から参照される場合)
ゼッド
.rules
アマゾンQ
.amazonq/rules/**
ジェットブレインズ ジュニー
.junie/ガイドライン.md
オープンハンズ
.openhands/microagents/*.md
ヒューマン ドキュメント (README、CONTRIBUTING、docs/) は対象外です。ロックファイル、おいおい

ifest と CI 構成はメモリではなく証拠です。これらはクレームがチェックされる対象です。
評決は4つ。すべての主張は、次の 1 つに集約されます。
健全性 = (確認済み + 0.5 × 古い) / (確認済み + 古い + 矛盾) × 100。有効なスタンプは確認としてカウントされます。知らないことは間違っていることと同じではないため、検証不可能な主張は除外されます。
出力システムは機能であり、副産物ではありません。レポートは、パルス グラフ プリミティブでレンダリングされた CALM ルール (ブロックごとに 5 つの結果に分割、最初に評決、ナレーションなし) に従います。パルス グラフ プリミティブは、監査の進行状況を一度に 1 行ずつストリーミングする追加専用のティッカーと、健全性、トークン コスト、トレンドのターミナル グラフです。すべての証拠行には、(直接) または (推定) 出所タグが付いています。なぜ C7 はあらゆる発見をその完全な証拠連鎖として再印刷するのか。 Brief はレポートを 10 行に折りたたみます。ファイルにはファイルごとのヘルス ストリップが表示されます。減衰線は次の半減期の期限を示します。適用された修正は 1 行の GAIN 概要 (健全性、誤解を招くトークン、回復されたトークン、前→後) で終わります。すべては完全な ASCII フォールバックを備えたプレーンな Unicode であり、依存関係はなくゼロから構築されています。詳細な仕様はreferences/output-format.mdにあります。各監査は .stale-brain/audit-YYYY-MM-DD.md に記録され、次回の実行時に傾向グラフにフィードされます。
機能チェックは、各プロジェクトの公開されたドキュメントとソースに対して 2026 年 7 月 31 日に検証されました。 ✓ 発送済み、~ 部分的、✗ 欠品。
1 行が最も重要であるとすれば、それはスタンプです。来歴がなければ、他のすべてのツールはすべてをゼロから再監査するため、「事実」がどれくらい古いのかを知ることができません。これらを使用すると、stalbrain は grep できる信頼モデルになります。
SKILL.md プロトコル (クロード コード エントリ ポイント)
PORTABLE.md 自己完結型の任意のエージェント バージョン
参考文献/
すべてのエージェントがその頭脳を保管するmemory-sources.md
claim-types.md 分類法、半減期、検証レシピ
ああ

utput-format.md CALM ルール、パルス グラフ、トークン メーター、レポート + 監査形式
src/stalebrain/ インストーラー CLI (staleebrain install/portable/path)
アセット/ヒーローとデモアート
pyproject.toml uv / pipx / pip パッケージ化
トラブルシューティング
スキルは発動しない。 /stale-brain を直接入力するか、トリガーの説明に「エージェントのメモリを監査する」、「CLAUDE.md を確認する」、「エージェントがルールを無視する」という語句を使用します。スキルの発動は説明と一致しており、どこでも不正確です。
判決は間違っているように見える。すべての評決では、ロックファイル、スクリプトブロック、コミットハッシュなどの証拠が引用されています。まず引用を確認してください。証拠が正しくても評決が依然として間違っている場合、それは報告する価値のあるバグです。
それは私が保持したいものにフラグを立てました。特定の差分の承認がなければ何も適用されません。修正を拒否します。主張はそのままです。スタイル設定にはフラグがまったく立てられません (OPINION クレームはスキップされます)。
トークンの数字がずれています。これらはファイル バイト ÷ 4 であり、推定値としてラベル付けされています (英語テキストの場合は ±20%)。シグナルは、絶対数ではなく、誤解を招くものの合計に対する比率です。
出力シンボルが文字化けして見える。 「アスキー」と言ってください。すべてのシンボルには、同じ行構造を持つプレーンテキストのフォールバックがあります。
私のリポジトリは浅いクローンです。ドリフト検出には履歴が必要です。 stalbrain は浅いクローンに気づき、空のドリフト パスを信頼する代わりに再検証しますが、完全な履歴の方がより適切な引用を提供します。
マサチューセッツ工科大学。 「ライセンス」を参照してください。内容はすべてゼロから書き下ろしたオリジナル作品です。
AI エージェントの記憶の出自と衰退。 CLAUDE.md、AGENTS.md、.cursorrules、GEMINI.md、および 18 個以上のエージェント メモリ ファイルをライブ リポジトリに対して監査します: 入力されたクレーム、引用されたコミット、半減期の減衰、トークン メーター、承認のみの修正。どのエージェントでも動作します。
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション

イゲーション
私の個人情報を共有しないでください

## Original Extract

Provenance and decay for AI agent memory. Audits CLAUDE.md, AGENTS.md, .cursorrules, GEMINI.md and 18+ agent memory files against the live repo: typed claims, cited commits, half-life decay, token meter, approve-only fixes. Works with any agent. - stalebrainlabs/stalebrain

GitHub - stalebrainlabs/stalebrain: Provenance and decay for AI agent memory. Audits CLAUDE.md, AGENTS.md, .cursorrules, GEMINI.md and 18+ agent memory files against the live repo: typed claims, cited commits, half-life decay, token meter, approve-only fixes. Works with any agent. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
stalebrainlabs
/
stalebrain
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
11 Commits 11 Commits .github/ workflows .github/ workflows assets assets references references src/ stalebrain src/ stalebrain .gitignore .gitignore LICENSE LICENSE PORTABLE.md PORTABLE.md README.md README.md SKILL.md SKILL.md pyproject.toml pyproject.toml View all files Repository files navigation
Your AI agent's memory files are lying to it. stalebrain finds every lie, proves it with commits, and hands you the fix.
Install ·
Demo ·
What it does ·
How it compares ·
Troubleshooting ·
License
Type /stale-brain in your AI coding assistant and it audits every agent memory file in your repo (CLAUDE.md, AGENTS.md, .cursorrules, GEMINI.md, Copilot instructions, 18+ locations across 13 agents) against the code as it exists today. Claims that check out get a dated stamp. Claims that don't get a verdict, the commits that prove it, and a ready-to-apply fix.
Not a linter. A trust model: after one audit, every claim in agent memory carries an age, a verification date, and cited evidence.
Prerequisite: a repo with git history. That's it.
With uv or pip (recommended, works inside Claude Code or any CLI)
uv tool install stalebrain # or: pip install stalebrain / pipx install stalebrain
stalebrain install # user-level: every repo, Claude Code CLI and desktop app
stalebrain install --project . # or just this repo
One-shot without installing anything permanent:
uvx stalebrain install
With git alone
# project-level
git clone https://github.com/stalebrainlabs/stalebrain .claude/skills/stale-brain
# or user-level (available in every repo)
git clone https://github.com/stalebrainlabs/stalebrain ~/.claude/skills/stale-brain
Then type /stale-brain , or say "audit my agent memory".
stalebrain install covers the Claude Code desktop app (it reads the same user-level skills folder). For every other desktop chat (ChatGPT, Gemini, Claude.ai), run stalebrain portable to print the single-file protocol, or just copy PORTABLE.md and paste it into the chat.
Any other assistant (Cursor, Copilot, Gemini CLI, ChatGPT, local models)
Paste PORTABLE.md into the chat, or drop it in your repo and say "run the stale-brain protocol in PORTABLE.md". It is fully self-contained and degrades gracefully: an assistant with no tools asks you to paste command output (batched into one block), and one with no file access prints the diffs for you to apply.
Always-on tripwire (optional, any agent)
Add one line to your always-loaded memory file (CLAUDE.md, AGENTS.md, or rules):
If an instruction in this file contradicts observed reality, say so in one line (⚡ stale-brain) instead of silently complying, and suggest a stale-brain audit.
See it in action
Verification is mechanical: glob, grep, and read-only git. No embeddings, no server, no database, no network. Nothing leaves your machine, and nothing gets executed: scripts are verified by their definition (scripts block, Makefile, CI), never by running them.
Agents keep "ignoring the rules"? Half the time the rules are unfollowable. They name paths that moved, scripts that were renamed, package managers that were swapped out in March. An agent fed contradictory memory doesn't get 90% right; it gets confidently wrong, every session, at a token cost you pay every session.
Capability
How
Claim extraction
Every sentence becomes a typed claim: PATH, SCRIPT, SYMBOL, DEP, FACT, OWNER, CONVENTION. Style opinions are skipped, never judged.
Verification
Per-type recipes against the live repo. A missing path gets git archaeology to find the rename and the new destination. A wrong owner gets checked against CODEOWNERS and git shortlog.
Confidence decay
Per-type half-lives (paths 30d, conventions 120d). Past its half-life a fact is re-verified or downgraded to a hypothesis.
Provenance stamps
<!-- stale-brain: verified 2026-07-31 --> , invisible in rendered markdown, greppable forever. Stamps make re-audits incremental.
Contradiction evidence
Never "this looks wrong". Always "wrong since a1b2c3d, here's the diff".
Cross-file conflicts
CLAUDE.md says yarn, .cursorrules says npm: flagged even when nobody knows which is right.
Token meter
What your memory costs per session, and what share of it is actively misleading the model.
Mid-task tripwire
When the agent notices an instruction contradicting reality mid-task, it says so in one line instead of silently complying.
Approve-only fixes
Every verdict is explained; every edit waits for your yes. Non-interactive runs apply nothing.
What files it handles
Agent
Files
Claude Code
CLAUDE.md (root and nested), CLAUDE.local.md, .claude/**/*.md
Codex / cross-tool
AGENTS.md (root and nested)
Cursor
.cursorrules, .cursor/rules/**/*.mdc
GitHub Copilot
.github/copilot-instructions.md, .github/instructions/*.instructions.md
Gemini CLI
GEMINI.md
Windsurf
.windsurfrules, .windsurf/rules/**
Cline
.clinerules (file or directory)
Roo Code
.roo/rules/**
Aider
CONVENTIONS.md (when referenced from .aider.conf.yml)
Zed
.rules
Amazon Q
.amazonq/rules/**
JetBrains Junie
.junie/guidelines.md
OpenHands
.openhands/microagents/*.md
Human docs (README, CONTRIBUTING, docs/) are out of scope. Lockfiles, manifests, and CI configs are evidence, not memory: they're what claims get checked against.
Four verdicts. Every claim lands in exactly one:
Health = (confirmed + 0.5 × stale) / (confirmed + stale + contradicted) × 100. A valid stamp counts as a confirmation; unverifiable claims are excluded because not knowing is not the same as being wrong.
The output system is a feature, not a byproduct. Reports follow CALM rules (chunked to 5 findings per block, verdict first, zero narration) rendered with Pulse graph primitives: an append-only ticker that streams audit progress one line at a time, and terminal graphs for health, token cost, and trend. Every evidence line carries a (direct) or (inferred) provenance tag; why C7 reprints any finding as its full evidence chain; brief collapses the report to 10 lines; files shows per-file health strips; a decay line names the next half-life expiry; and applied fixes end in a one-line GAIN summary (health, misleading tokens, and recovered tokens, before → after). All of it is plain Unicode with a complete ASCII fallback, built from scratch, zero dependencies. Full spec in references/output-format.md . Each audit is recorded in .stale-brain/audit-YYYY-MM-DD.md , which feeds the trend graph on the next run.
Feature checks verified 2026-07-31 against each project's published docs and source. ✓ shipped, ~ partial, ✗ absent.
If one row matters most, it's the stamps: without provenance, every other tool re-audits everything from zero and can't tell you how old a "fact" is. With them, stalebrain is a trust model you can grep.
SKILL.md the protocol (Claude Code entry point)
PORTABLE.md self-contained any-agent version
references/
memory-sources.md where every agent keeps its brain
claim-types.md taxonomy, half-lives, verification recipes
output-format.md CALM rules, Pulse graphs, token meter, report + audit formats
src/stalebrain/ the installer CLI (stalebrain install / portable / path)
assets/ hero and demo art
pyproject.toml uv / pipx / pip packaging
Troubleshooting
The skill doesn't trigger. Type /stale-brain directly, or use the phrases in the trigger description: "audit my agent memory", "verify my CLAUDE.md", "my agent ignores the rules". Skill activation is description-matched and inexact everywhere.
A verdict looks wrong. Every verdict cites its evidence: a lockfile, a scripts block, a commit hash. Check the citation first; if the evidence is right and the verdict is still wrong, that's a bug worth reporting.
It flagged something I want to keep. Nothing is applied without your approval of the specific diff. Decline the fix; the claim stays as it was. Style preferences are never flagged at all (OPINION claims are skipped).
The token numbers look off. They're file bytes ÷ 4, labeled as estimates (±20% for English text). The signal is the ratio of misleading to total, not the absolute count.
Output symbols look garbled. Say "ascii". Every symbol has a plain-text fallback with the same line structure.
My repo is a shallow clone. Drift detection needs history. stalebrain notices shallow clones and re-verifies instead of trusting an empty drift pass, but full history gives better citations.
MIT. See LICENSE . All content is original work, written from scratch.
Provenance and decay for AI agent memory. Audits CLAUDE.md, AGENTS.md, .cursorrules, GEMINI.md and 18+ agent memory files against the live repo: typed claims, cited commits, half-life decay, token meter, approve-only fixes. Works with any agent.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
