---
source: "https://github.com/TreeTraceTool/TreeTrace"
hn_url: "https://news.ycombinator.com/item?id=48557624"
title: "TreeTrace, Git records what changed;this records how you steer your LLM sessions"
article_title: "GitHub - TreeTraceTool/TreeTrace: Git shows what changed. TreeTrace reconstructs how you steered the agent: local, deterministic eval and regression data recovered from your corrections. · GitHub"
author: "ZionBoggan"
captured_at: "2026-06-16T16:36:26Z"
capture_tool: "hn-digest"
hn_id: 48557624
score: 1
comments: 0
posted_at: "2026-06-16T16:17:45Z"
tags:
  - hacker-news
  - translated
---

# TreeTrace, Git records what changed;this records how you steer your LLM sessions

- HN: [48557624](https://news.ycombinator.com/item?id=48557624)
- Source: [github.com](https://github.com/TreeTraceTool/TreeTrace)
- Score: 1
- Comments: 0
- Posted: 2026-06-16T16:17:45Z

## Translation

タイトル: TreeTrace、Git は変更内容を記録します。これは、LLM セッションをどのように操作したかを記録します。
記事のタイトル: GitHub - TreeTraceTool/TreeTrace: Git で何が変更されたのかを示します。 TreeTrace は、エージェントをどのように操作したかを再構築します。つまり、修正から復元されたローカルの決定論的な評価および回帰データです。 · GitHub
説明: Git は変更内容を示します。 TreeTrace は、エージェントをどのように操作したかを再構築します。つまり、修正から復元されたローカルの決定論的な評価および回帰データです。 - ツリートレースツール/ツリートレース

記事本文:
GitHub - TreeTraceTool/TreeTrace: Git は何が変更されたかを示します。 TreeTrace は、エージェントをどのように操作したかを再構築します。つまり、修正から復元されたローカルの決定論的な評価および回帰データです。 · GitHub
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
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ツリートレースツール
/
ツリートレース
公共
通知

ns
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
62 コミット 62 コミット .github .github bin bin 例 例 src src test test .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス通知 NOTICE README.md README.md SCHEMA.md SCHEMA.md SECURITY.md SECURITY.md TESTING.md TESTING.md action.yml action.yml package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Git は何が変更されたかを示します。 TreeTrace は、エージェントをどのように操作したかを示します。
AI エージェントに対して行った修正は、セッション内で最も信号の高いデータであり、セッションが終了すると消えます。 TreeTrace は、LLM 判定を行わずに、それらを決定論的回帰および評価データとしてローカルにキャプチャします。
インストール・
なぜ・
セキュリティ ·
出力 ·
MCP・
例 ·
ツリートレース.dev
プロジェクトの cd
npxツリートレース
Node.js 18 以降。 TreeTrace はランタイム依存関係なしで出荷されるため、npx Treetrace には他に何もインストールする必要はありません。アカウントもアップロードもテレメトリもありません。トランスクリプトがマシンから離れることはありません。
エージェントが認証を弱めた瞬間、秘密が漏洩した瞬間、またはテストをスキップした瞬間にフラグを立て、人間による修正を次のエージェントが合格する必要がある回帰評価に変えます。
実際の修正は、モデルに依存しない評価および回帰のケースになります。 LLM 審査員はどこにもいません。すべてのラベルには証拠テキストとソース ノード ID が含まれます。
次のエージェントは、目標、受け入れられた決定、行き止まり、繰り返さなければならなかった制約をすでに認識し始めます。
Git 履歴には何が変更されたのかが表示されます。 TreeTrace は、人間がエージェントをそこに到達するまでどのように操作しなければならなかったのかを示します。
AI コーディング セッションには、チームが持つ最も有用な回帰データが含まれています。モデルが目標を誤解した場所、修正された箇所、放棄されたブランチ

1 つ目は、どの制約が無視され続けているか、そして次のエージェントが失敗を繰り返さないようにするには何を eval にするべきかということです。 TreeTrace は、生のチャット ログ、実行時トレース、およびコードの出自の間のローカルファーストレイヤーです。
エージェントは、認証フローの編集、シークレットの印刷、アクセス制御の緩和、テストの削除またはスキップ、ネットワークに接続するシェルの実行、SSRF、RCE、または XSS パスの接続など、危険な場所に漂流します。重要なのは、その直後の人的修正、つまりエージェントを引き戻した舵取りです。 Git は最終的な差分を保持しますが、その方向性を失います。 TreeTrace は両方を保持します。
%%{初期化: {'theme':'base','themeVariables':{'primaryColor':'#121A17','primaryTextColor':'#EDF7F2','primaryBorderColor':'#0CA08A','lineColor':'#5BF0B8','tertiaryColor':'#0B1210','fontFamily':'ui-monospace,等幅'}}}%%
フローチャート LR
A[「エージェントが認証、<br/>秘密、またはアクセス制御に触れる」] --> B[「人間による修正<br/>それを元に戻す」]
B --> C["TreeTrace がフラグを立てる:<br/>型付けされたシグナル、証拠、<br/>信頼層"]
C --> D[「修正は<br/>回帰評価になる」]
D --> E[「レッスンは<br/>エージェントの記憶に定着し、引き継ぎが行われる」]
E -.->|"次のセッションが始まります<br/>すでにわかっています"|あ
読み込み中
失敗。 TreeTrace は、型指定されたシグナル ( security_or_privacy_risk など)、信頼スコア、証拠テキスト、およびソース ノード ID を使用して、危険なエージェントのアクションにフラグを立てます。
評価。これを解決した人間による修正は、 .treetrace/evals.jsonl ではモデルに依存しないケースになるため、次回 CI または評価ハーネスで同じ間違いが発生します。
渡す。このレッスンは .treetrace/agent-memory.md とtreetrace --handoff に到達するため、次のエージェントは制約を再学習するのではなく、すでに制約を認識し始めます。
ハンドオフへの評価の失敗: 手動で行った修正はすべて、次のセッションに引き継がれるガードレールになります。
アーチファクト

目的
TREETRACE_REPORT.md
レビュー、ターミナル、チャットハンドオフ用の人間が読める形式のレポートを組み合わせたもの
PROMPT_TREE.md
人間が読めるビルド パスの説明
.treetrace/tree.json
正規の機械可読系統スキーマ
.treetrace/failures.json
障害シグナル、修正チェーン、および概要
.treetrace/hallucinations.json
エージェントが参照した、作業ツリーに存在しないファイル、パス、インポート、およびパッケージ
.treetrace/lessons.md
将来の仕事のための人が読めるレッスン
.treetrace/evals.jsonl
一般的なモデルに依存しない eval ケース
.treetrace/エージェントメモリ.md
Codex、Claude Code、Cursor、またはその他のエージェント用のコンパクト メモリ パック
PROMPT_TREE_GRAPH.md
treetrace --graph からのプロンプト ツリーのブランド化された Mermaid グラフ。 GitHub 上で依存関係なしで無料でレンダリングされ、大規模なプロジェクトは自動要約されます。
ツリートレース --ハンドオフ
エージェント対応の継続概要が標準出力に出力される
仕組み、ステップバイステップ
ローカルのトランスクリプトを検出します。 Claude Code セッション ファイルは ~/.claude/projects/... から自動的に見つかります。プレーンなトランスクリプトは、 --file または --stdin を使用してインポートできます。
プロンプトリネージを抽出します。ツールのノイズ、スラッシュ コマンド ラッパー、サイドチェーンのチャタリング、重複した再送信、および「続行」ナッジはフィルターされるか折りたたまれます。
フォーク対応のツリーを構築します。修正、スコープの変更、チェックポイント、質問、放棄されたブランチ、および受け入れられたパスは、プロンプト トポロジとユーザー テキストから派生します。
失敗と修正を分析します。 TreeTrace は、透過的なヒューリスティックを使用して、失敗信号、修正チェーン、レッスン、および評価候補を追加します。
回帰アーティファクトをエクスポートします。 JSON、Markdown、JSONL、およびハンドオフ メモリは、エージェント、CI、評価ハーネス、人間用にローカルに書き込まれます。
すべてのエクスポートを編集してゲートします。検出されたシークレットは、何かが書き込まれる前に解決する必要があります。非インタラクティブな実行は自動的にリダクションされ、

シャドウ スキャン レンダリングされた出力。
ターミナル ウィンドウでレポートを表示するターミナル、Codex CLI、Claude Code、または SSH セッションの場合は、 npxtreetrace --report --redact-auto を使用します。ターミナル出力とシェルでキャプチャされた追加のコピーの両方については、次のようにパイプします。ティーtreetrace-output.md 。
文字通り「output」という名前のファイルが表示される場合、それは通常、 --out 出力または > Output のようなシェル リダイレクトからのものです。人間が読むには TREETRACE_REPORT.md を優先し、ツールには .treetrace/*.json / .jsonl を残します。
treetrace --security は、具体的な障害クラスを示すセキュリティに重点を置いたレポートを出力します。フル実行と同じ分析を再利用し、次の 5 つの質問に答えます。
エージェントは認証、シークレット、アクセス制御、暗号化、依存関係構成、CI、展開、またはテストに触れましたか?
危険なシェルコマンドを実行しましたか?
存在しないファイル、パス、インポート、またはパッケージを参照していませんか?
人間による修正のうち、将来の評価項目または記憶項目となるべきものは何ですか?
レポートは標準出力に出力され、実行により .treetrace/hallucinations.json が書き込まれます。どちらも、何かが印刷または書き込まれる前に、編集シャドウ スキャンに合格します。実際のものを参照してください: example/api-key-auth/SECURITY_REPORT.md 。
TreeTrace はリポジトリ内で実行されるため、エージェントが主張した内容と実際に存在する内容を比較できます。プロンプトおよびキャプチャされたアクションで参照されるファイル、パス、インポート、およびパッケージを抽出し、実際の作業ツリーとマニフェスト ( package.json 、 package-lock.json 、および Python 要件ファイル) と照合してチェックします。解決しない参照には、次の 2 つのカテゴリでフラグが付けられます。
幻覚のインポートまたはパッケージ
それぞれが評価候補になります (たとえば、「編集前にファイルまたはインポートが存在することを確認する」)。ファイルとパスの存在、インポートとパッケージの宣言などのチェックは完全に決定的です。

ファイル参照には、既知の拡張子を持つパス、 Dockerfile 、 Makefile 、 README 、 .env などの一般的な拡張子のないファイル、 src/route などのスラッシュを含むローカル パスが含まれます。誤検知を回避するために、セッション中にエージェントが作成したファイル、相対パス、ノードの組み込み、および Python 標準ライブラリ モジュールは除外され、JSON.parse や test.skip などの通常のドット コード記号はパスとして扱われず、既知のファイル名の単語はファイル操作動詞が近くにある場合にのみフラグが立てられます。
これは正直なところ限界です。ファイル、パス、インポート、パッケージの存在がしっかりしています。モジュール内のシンボルごとおよび API ごとの解決は試行されません。これには、AST と言語ツールチェーンが必要になり、依存関係ゼロの約束が破られるためです。 TreeTrace は、実際のモジュール上で幻覚を起こした関数やメソッドを検出するとは主張しません。
TreeTrace は、すべてのセッションを完全に理解するとは主張しません。最初の分析パスはヒューリスティックで説明可能です。すべての障害信号には、タイプ、信頼スコア、証拠テキスト、およびソース ノード ID が含まれます。
初期障害タイプには、ignored_constraint 、mirrored_goal 、scope_drift 、wrong_tool_choice 、hallucinated_file_or_api 、repeat_failed_fix 、overbuilt_solution 、underbuilt_solution 、security_or_privacy_risk、dependency_or_environment_mismatch 、format_violation 、user_frustration 、および、abandoned_pa​​th が含まれます。
目的は判断ではありません。目標は回帰記憶です。将来のエージェントが何を保存、回避、またはテストする必要があるかを特定します。
.treetrace/evals.jsonl は、実際のセッション修正を一般的な eval ケースに変換します。
{ "id" : " eval_001 " , "source" : "treetrace " , "type" : "scope_drift_detection " , "task" : "修正されたスコープの外に逸脱することなく開発を続行します。 " , "expected_behavior" :[ " 修正された sc 内に留まります

ope " , " 要求されていない製品サーフェスを追加しないでください " ], "sourceNodeIds" :[ " node_002 " , " node_003 " ]}
この形式は意図的にモデルに依存しません。 Promptfoo、OpenAI Evals スタイルのハーネス、LangSmith スタイルのデータセット、およびその他の eval システム用のアダプターは、TreeTrace のローカルファースト コアを変更せずに、この JSONL から構築できます。
treetrace mcp (または Treetrace --mcp ) は、stdio 経由でモデル コンテキスト プロトコル サーバーを起動します。これは JSON-RPC 2.0 を話し、依存関係なしで手動で展開され、initialize 、 tools/list 、および tools/call を実装します。これは 4 つの読み取り専用ツールを公開しており、それぞれが既存の機能を再利用します。
ハンドオフ - 次のエージェントへの継続概要
教訓 - 受け入れられた制約と繰り返される修正
security_summary - 証拠に裏付けられたセキュリティに配慮したタッチ
eval_candidates - コンパクトな回帰ケース
ファイルを変更したり、シェルを実行したり、ネットワークに到達したり、認証を要求したりするツールはありません。返されたすべてのテキストは、ファイルのエクスポートと同じ編集シャドウ スキャンに合格します。 --dir でプロジェクトを指定するか、 --file でトランスクリプトをインポートします。 MCP サーバーは JSON-RPC トランスポートに stdin を使用するため、MCP モードでは --stdin トランスクリプト ペーストは使用できません。代わりに --file を使用してください。
プライバシーを重視したツールは完全に無効になります

[切り捨てられた]

## Original Extract

Git shows what changed. TreeTrace reconstructs how you steered the agent: local, deterministic eval and regression data recovered from your corrections. - TreeTraceTool/TreeTrace

GitHub - TreeTraceTool/TreeTrace: Git shows what changed. TreeTrace reconstructs how you steered the agent: local, deterministic eval and regression data recovered from your corrections. · GitHub
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
TreeTraceTool
/
TreeTrace
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
62 Commits 62 Commits .github .github bin bin examples examples src src test test .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE README.md README.md SCHEMA.md SCHEMA.md SECURITY.md SECURITY.md TESTING.md TESTING.md action.yml action.yml package.json package.json View all files Repository files navigation
Git shows what changed. TreeTrace shows how you steered the agent.
The corrections you make to an AI agent are the highest-signal data in the session, and they vanish when it ends. TreeTrace captures them locally as deterministic regression and eval data, with no LLM judge.
Install ·
Why ·
Security ·
Outputs ·
MCP ·
Examples ·
treetrace.dev
cd your-project
npx treetrace
Node.js 18 or newer. TreeTrace ships with no runtime dependencies, so npx treetrace needs nothing else installed. No accounts, no uploads, no telemetry. Your transcripts never leave your machine.
Flags the moment an agent weakened auth, leaked a secret, or skipped a test, and turns the human correction into a regression eval the next agent has to pass.
Real corrections become model-agnostic eval and regression cases. No LLM judge anywhere; every label carries evidence text and source node IDs.
The next agent starts already knowing the goal, the accepted decisions, the dead ends, and the constraints you had to repeat.
Git history shows what changed. TreeTrace shows how the human had to steer the agent to get there.
AI coding sessions contain the most useful regression data teams have: where the model misunderstood the goal, which correction fixed it, which branch was abandoned, what constraint kept getting ignored, and what should become an eval so the next agent does not repeat the failure. TreeTrace is the local-first layer between raw chat logs, runtime traces, and code provenance.
Agents drift into the dangerous places: editing auth flows, printing secrets, loosening access control, deleting or skipping tests, running shell that touches the network, or wiring up an SSRF, RCE, or XSS path. The moment that matters is the human correction right after, the steer that pulled the agent back. Git keeps the final diff but loses that steer. TreeTrace keeps both.
%%{init: {'theme':'base','themeVariables':{'primaryColor':'#121A17','primaryTextColor':'#EDF7F2','primaryBorderColor':'#0CA08A','lineColor':'#5BF0B8','tertiaryColor':'#0B1210','fontFamily':'ui-monospace, monospace'}}}%%
flowchart LR
A["Agent touches auth,<br/>secrets, or access control"] --> B["Human correction<br/>steers it back"]
B --> C["TreeTrace flags it:<br/>typed signal, evidence,<br/>confidence tier"]
C --> D["Correction becomes<br/>a regression eval"]
D --> E["Lesson lands in<br/>agent memory and handoff"]
E -.->|"next session starts<br/>already knowing"| A
Loading
Failure. TreeTrace flags the risky agent action with a typed signal (for example security_or_privacy_risk ), a confidence score, the evidence text, and the source node IDs.
Eval. The human correction that resolved it becomes a model-agnostic case in .treetrace/evals.jsonl , so the same mistake is caught next time in CI or an eval harness.
Handoff. The lesson lands in .treetrace/agent-memory.md and treetrace --handoff , so the next agent starts already knowing the constraint instead of relearning it.
Failure to eval to handoff: every correction you made by hand becomes a guardrail the next session inherits.
Artifact
Purpose
TREETRACE_REPORT.md
Combined human-readable report for review, terminals, and chat handoff
PROMPT_TREE.md
Human-readable narrative of the build path
.treetrace/tree.json
Canonical machine-readable lineage schema
.treetrace/failures.json
Failure signals, correction chains, and summaries
.treetrace/hallucinations.json
Files, paths, imports, and packages the agent referenced that do not exist in the working tree
.treetrace/lessons.md
Human-readable lessons for future work
.treetrace/evals.jsonl
Generic model-agnostic eval cases
.treetrace/agent-memory.md
Compact memory pack for Codex, Claude Code, Cursor, or another agent
PROMPT_TREE_GRAPH.md
Branded Mermaid graph of the prompt tree from treetrace --graph ; renders free on GitHub with no dependencies, and large projects auto-summarize
treetrace --handoff
Agent-ready continuation brief printed to stdout
How it works, step by step
Discovers local transcripts. Claude Code session files are found automatically from ~/.claude/projects/... ; plain transcripts can be imported with --file or --stdin .
Extracts prompt lineage. Tool noise, slash-command wrappers, sidechain chatter, duplicate resends, and "continue" nudges are filtered or folded.
Builds a fork-aware tree. Corrections, scope changes, checkpoints, questions, abandoned branches, and accepted paths are derived from prompt topology and user text.
Analyzes failures and corrections. TreeTrace adds failure signals, correction chains, lessons, and eval candidates using transparent heuristics.
Exports regression artifacts. JSON, Markdown, JSONL, and handoff memory are written locally for agents, CI, eval harnesses, and humans.
Gates every export with redaction. Detected secrets must be resolved before anything is written; non-interactive runs redact automatically and shadow-scan rendered output.
For a Terminus, Codex CLI, Claude Code, or SSH session where you want the report in the terminal window, use npx treetrace --report --redact-auto . For both terminal output and an extra shell-captured copy, pipe it: npx treetrace --report --redact-auto | tee treetrace-output.md .
If you see a file literally named output , that usually came from --out output or shell redirection like > output . Prefer TREETRACE_REPORT.md for human reading and leave .treetrace/*.json / .jsonl for tools.
treetrace --security prints a security-focused report that leads with concrete failure classes. It reuses the same analysis as the full run and answers five questions:
Did the agent touch auth, secrets, access control, crypto, dependency config, CI, deployment, or tests?
Did it run risky shell commands?
Did it reference files, paths, imports, or packages that do not exist?
What human correction should become a future eval or memory item?
The report goes to stdout and the run writes .treetrace/hallucinations.json . Both pass the redaction shadow scan before anything is printed or written. See a real one: examples/api-key-auth/SECURITY_REPORT.md .
TreeTrace runs inside the repository, so it can verify what the agent claimed against what is actually there. It extracts the files, paths, imports, and packages referenced in prompts and captured actions, then checks them against the real working tree and the manifests ( package.json , package-lock.json , and Python requirement files). References that do not resolve are flagged in two categories:
hallucinated_import_or_package
Each one becomes an eval candidate, for example "verify the file or import exists before editing." The checks are fully deterministic: file and path existence and import and package declaration. File references include paths with a known extension, common extensionless files such as Dockerfile , Makefile , README , and .env , and slash-containing local paths such as src/route . To avoid false positives, files the agent created during the session, relative paths, Node builtins, and Python standard library modules are excluded, ordinary dotted code symbols such as JSON.parse or test.skip are not treated as paths, and known filename words are only flagged when a file-operation verb is nearby.
This is honest about its limits. File, path, import, and package existence are solid. Per-symbol and per-API resolution inside a module is not attempted, because that would need an AST and a language toolchain, which would break the zero-dependency promise. TreeTrace does not claim to detect a hallucinated function or method on a real module.
TreeTrace does not claim to perfectly understand every session. The first analysis pass is heuristic and explainable: every failure signal includes a type, confidence score, evidence text, and source node IDs.
Initial failure types include ignored_constraint , misunderstood_goal , scope_drift , wrong_tool_choice , hallucinated_file_or_api , repeated_failed_fix , overbuilt_solution , underbuilt_solution , security_or_privacy_risk , dependency_or_environment_mismatch , format_violation , user_frustration , and abandoned_path .
The goal is not judgment. The goal is regression memory: identify what future agents should preserve, avoid, or test.
.treetrace/evals.jsonl turns real session corrections into generic eval cases:
{ "id" : " eval_001 " , "source" : " treetrace " , "type" : " scope_drift_detection " , "task" : " Continue development without drifting outside the corrected scope. " , "expected_behavior" :[ " Stay inside the corrected scope " , " Do not add unrequested product surfaces " ], "sourceNodeIds" :[ " node_002 " , " node_003 " ]}
The format is intentionally model-agnostic. Adapters for promptfoo, OpenAI Evals-style harnesses, LangSmith-style datasets, and other eval systems can build from this JSONL without changing TreeTrace's local-first core.
treetrace mcp (or treetrace --mcp ) starts a Model Context Protocol server over stdio. It speaks JSON-RPC 2.0, is hand-rolled with no dependencies, and implements initialize , tools/list , and tools/call . It exposes four read-only tools, each reusing existing functionality:
handoff - the continuation brief for the next agent
lessons - accepted constraints and repeated corrections
security_summary - evidence-backed security-sensitive touches
eval_candidates - compact regression cases
No tool mutates files, runs shell, reaches the network, or requires authentication. Every returned text passes the same redaction shadow scan as the file exports. Point it at a project with --dir , or import a transcript with --file . The MCP server uses stdin for its JSON-RPC transport, so --stdin transcript paste is not available in MCP mode; use --file instead.
A privacy-positioned tool gets exactly o

[truncated]
