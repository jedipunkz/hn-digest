---
source: "https://github.com/MasihMoafi/Elpis"
hn_url: "https://news.ycombinator.com/item?id=49017977"
title: "Show HN: Elpis – a Rust TUI for LLM agents with context pruning"
article_title: "GitHub - MasihMoafi/Elpis: You put an agent into an Elpis, and it becomes Elpis; Be Elpis my friend. · GitHub"
author: "masihmoafi"
captured_at: "2026-07-23T08:12:27Z"
capture_tool: "hn-digest"
hn_id: 49017977
score: 1
comments: 0
posted_at: "2026-07-23T07:20:16Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Elpis – a Rust TUI for LLM agents with context pruning

- HN: [49017977](https://news.ycombinator.com/item?id=49017977)
- Source: [github.com](https://github.com/MasihMoafi/Elpis)
- Score: 1
- Comments: 0
- Posted: 2026-07-23T07:20:16Z

## Translation

タイトル: HN を表示: Elpis – コンテキスト プルーニングを備えた LLM エージェント用の Rust TUI
記事のタイトル: GitHub - MasihMoafi/Elpis: Elpis にエージェントを入れると、Elpis になります。エルピスと友達になってください。 · GitHub
説明: Elpis にエージェントを入れると、それが Elpis になります。エルピスと友達になってください。 - MasihMoafi/エルピス

記事本文:
GitHub - MasihMoafi/Elpis: Elpis にエージェントを入れると、それは Elpis になります。エルピスと友達になってください。 · GitHub
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
マシモアフィ
/
エルピス
公共
通知
通知設定を変更するにはサインインする必要があります

s
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
497 コミット 497 コミット .github/ workflows .github/ workflows codex-rs codex-rs docs docs example/ gemini-cli example/ gemini-cli scripts scripts src src testing testing .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md GUIDE.md GUIDE.md ライセンス ライセンス通知 NOTICE TASKS.md TASKS.md pyproject.toml pyproject.toml readme.md readme.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ターミナルエージェントはすべてを忘れてしまいます。エルピスはそうではありません。
今日、すべてのコーディング エージェントが同じ病気を患っています。長く働くほど、症状は悪化します。
目標は記録に埋もれてしまいます。決定は圧縮時に失われます。すべての新しいセッションは次のように始まります
タスクを再度説明すると、すべてのモデル リクエストがこれまでよりも太くなった履歴を再送信します。
あなたが払っているのです。
Elpis は、コンテキストをコンテキストではなく管理資産として扱うターミナル エージェント シェルです。
増え続けるログ:
ターン後のコンテキストの枝刈り。ターンが答えを導き出した後、生の探索が始まる —
コマンド ダンプ、ファイル読み取り、失敗したプローブ — 次のリクエストでコンパクトによって置き換えられます
正確な証拠が記載された領収書。完全なトランスクリプトはディスク上に残り、
要求;デフォルトでは再送信が停止されるだけです。他のオープンエージェントはターンごとにこれを行いません。
コンテキスト台帳。目に見えるスクロール可能なパネルには、入力された内容が正確に表示されます。
次のモデル リクエスト — すべてのルール ファイル、ゴール、チェックポイント、および追加されたファイルをそれぞれ独自の行として、
それぞれ切り替え可能です。モデルが何を認識するかを決めるのはあなたです。 /独自のファイルを追加します。
ポータブルな継続性。アクティブな目標 ( GOAL.md ) とリーン チェックポイント ( ES.md )
圧縮、セッションの終了、さらにはランタイムの切り替えでも存続します。スレッドを正確に再開し、
または、目標、前回の結果、結果がすでにわかっている新しいものを開始します。

次のアクション —
履歴を再生せずに。
獲得した記憶。事実は、有用な思い出を繰り返して初めて永続的な記憶となります。
異なるコンテキスト。それ以外はすべて検索可能な証拠として残ります。消された思い出は、
リセット前にアーカイブされるため、サイレントに失われることはありません。
ランタイムに依存しない。どのモデルがループを実行しても、シェルはその下で存続します。
ChatGPT/Codex サブスクリプション ログイン、OpenRouter、ネイティブ Anthropic アダプターおよび Gemini アダプター、
ベッドロック、オラマ、LMスタジオ。 Elpis にモデルを入れると、それが Elpis になります。あなたの目標、
コンテキスト、記憶、ルール。
実行基盤 (ターミナル UI、パッチ、権限、サンドボックス、セッション) は、
OpenAI の Apache-2.0 Codex CLI のフォークを差し引いたもので、アップストリームによって強化され、ここで所有されています。
現在の状態: 最初のリリースに近づいています。 v0.1.0 はライブ後にのみタグ付けされます
承諾は TASKS.md に記録されます。緑色のバッジはチェックに合格したことを意味します —
それは決して未完成の仕事が終わったという意味ではありません。
ストリーミング コマンド、パッチ、許可モードを備えたネイティブ Ratatui ターミナル インターフェイス
サンドボックス、マウス選択、セッション、圧縮。
ChatGPT サブスクリプション認証; OPENROUTER_API_KEY を介した OpenRouter ;ネイティブ
Anthropic メッセージと Google Gemini アダプター ( --provider anthropic 、
--provider google-gemini ;ライブベンダーの受け入れは保留中です）。
1 つの内部読み取り専用 RAG サービス: /rag <query> 、 /rag <path> -- <query> 、および
自律的な検索。
ポータブル GOAL.md + ES.md の連続性。正確な再開または無駄のない継続。
ファイルごとの行と切り替えを備えたコンテキスト台帳。 /許可されたすべての使用状況レポート
サイズ、寿命、および理由を含むソース。
リコール追跡、プロモーション、来歴、および
フェイルクローズされたアーカイブ。
決定的な初回パスのコンテキスト クリーニング: 古い長いツールの出力が制限される
耐久性のある rollout:// 証拠ポインタを含む抜粋。
いっぱい

l プルーニング エンジン — で説明されているエージェント作成のターン結果レコード
docs/CONTEXT_AND_SESSIONS.md — 主な機能です
アクティブな開発では、ターンごとの「コンテキスト保存」メトリクスが表示されます。
Elpis は、選択されたランタイムがモデル ループを実行している間、周囲の制御環境を安定に保ちます。正確な証拠は永続的に残ります。小規模で合理的なワーキング セットのみが次のリクエストを入力します。
フローチャート LR
user([ユーザーリクエスト]) --> tui["Elpis TUI"]
tui --> control["Elpis コントロール層"]
rag["読み取り専用 RAG"] --> コントロール
ポータブル["GOAL.md + ES.md"] --> コントロール
メモリ["境界のあるローカルメモリ"] --> 制御
コントロール --> ランタイム{"ランタイムの選択"}
ランタイム --> openai["OpenAI / Codex"]
ランタイム --> ルーター["OpenRouter · Anthropic · Gemini"]
ランタイム --> その他["Bedrock / Ollama / LM Studio"]
openai --> 実行["ツール · 編集 · コマンド"]
ルーター --> 実行
その他 --> 実行
実行 --> 証拠[("ワークスペース + 正確な証拠")]
読み込み中
コンテキスト管理
Elpis では、ルール、現在のリクエスト、移植可能な状態、および関連するメモリを小さなワーキング セットに入れることができます。 Context Ledger と /usage は、完全なアーティファクトがディスク上に残っているにもかかわらず、各ソースが存在する理由を明らかにします。
作業コンテキストはトランスクリプトではありません。完全な会話、ターミナルイベント、アーティファクト
証拠として利用可能なままですが、後のタスクで必要になった場合にのみ取得されます。
証拠ポインタが最初、RAG が 2 番目です。どちらも完全な履歴をデフォルトのプロンプトにしません
付属品。目的は、控えめなコンテキスト ウィンドウで十分で読みやすいものにすることです。
成長し続けるものを再送信するには料金を支払います。
フローチャート LR
rules["ルール + 現在のリクエスト"] --> working["許可された小規模なワーキング セット"]
目標["GOAL.md ≤ 6,000 chars"] --> 動作中
チェックポイント["ES.md ≤ 8,000 chars"] --> 動作中
メモリ["関連メモリ"] --> 動作中
ステータス["コン

テキスト元帳 + /使用法"] -.-> 作業中
作業中 --> ランタイム["選択されたランタイム リクエスト"]
実行時 --> 結果["ツールと関数の結果"]
結果 --> ディスク[("完全なトランスクリプト + ディスク上のアーティファクト")]
results -->large{"古いツールの出力 > 1,200 文字?"}
大 -->|はい|マーカー["コンパクトレシート + 証拠ポインタ"]
大 -->|いいえ| keep["リクエストコンテキストに保持"]
読み込み中
セッションの継続性
Elpis は、有用なネイティブ スレッドを正確に再開するか、 GOAL.md および ES.md から無駄のないスレッドを開始します。圧縮前の同期は、ハンドオフが中断される危険を冒す代わりに、失敗して閉じられます。
フローチャート LR
boundary(["終了・圧縮・ランタイム変更"]) --> sync["GOAL.md + ES.mdを同期"]
sync -->safe{"ファイルは安全に同期されましたか?"}
安全 -->|いいえ| stop["圧縮を停止 + エラーを表示"]
安全 -->|はい| thread{"ネイティブ スレッドはまだ役に立ちますか?"}
スレッド -->|はい|正確["正確な履歴書"]
スレッド -->|いいえ|リーン[「リーン継続」]
正確 --> 次["保存された次のアクションから続行"]
リーン --> 次へ
次 --> 証拠[「必要な場合にのみ正確な証拠を読む」]
証拠 --> verify["検証結果 + チェックポイントを更新"]
読み込み中
メモリ管理
新しい証拠は、繰り返し有用な想起によって永続記憶の対象となるまで、検索可能のままです。耐久性のあるアーティファクトは制限されており、削除または色あせた行はリセット前にアーカイブされます。
フローチャートTB
サブグラフプロモーション["プロモーション"]
LR方向
work(["完成した作品"]) --> raw["検索可能な証拠 + 来歴"]
raw --> remember["トラックリコール + 個別のコンテキスト"]
リコール --> ゲート{"3 は 2 つのコンテキストにわたってリコールしますか?"}
ゲート -->|いいえ| short[「検索可能な証拠を残す」]
ゲート -->|はい|耐久性["MEMORY.md の対象"]
耐久性 --> 制限付き["MEMORY.md ≤ 30k · summary ≤ 10k"]
終わり
サブグラフ 退職["退職"]
LR方向
変更されました{"削除されましたか、それともフェードされましたか?"}
変更されました -->|いいえ|リセット["メモリのベースラインをリセット"]
変更されました -->|はい

| archive["削除された行を archive.md に追加する"]
アーカイブ --> 書き込まれました{"アーカイブの書き込みに成功しましたか?"}
書かれています -->|はい|リセット
書かれています -->|いいえ| block["証拠の損失を防ぐためにブロックをリセット"]
終わり
短い --> 変更されました
境界あり --> 変更されました
読み込み中
読み取り専用 RAG
起動パスは、取得スタックをロードせずに 1 つの最小限の読み取り専用ツールを公開します。埋め込みとインデックス作成は、明示的なセマンティック クエリの後にのみ遅延して読み込まれます。
フローチャート LR
launch(["Elpis の起動"]) --> host["最小限の stdio MCP ホスト"]
ホスト --> ツール["1 つのツール: query_knowledge_base"]
ツール --> クエリ{"明示的な RAG クエリ?"}
クエリ -->|いいえ| idle["インデックス作成またはモデルのロードなし"]
クエリ -->|はい| search["遅延ロード rag.fetch"]
検索 --> スコープ{"ワークスペースまたは指定されたパス?"}
スコープ --> ワークスペース["ワークスペースの起動"]
スコープ --> provided["正規化された指定されたパス"]
ワークスペース --> チャンク["ソースされたセマンティック チャンク"]
提供された --> チャンク
チャンク --> 正確[「コード変更前に正確な読み取りを使用する」]
読み込み中
安全な執行と証拠
結果として生じるアクションは、実行前に表示されるアクセス許可とサンドボックス ポリシーを通過します。 Elpis は結果と証拠を記録し、検証された成功と失敗または未解決のギャップを区別します。
フローチャート LR
action(["提案された結果的なアクション"]) --> ポリシー["権限 + サンドボックス ポリシーを表示"]
ポリシー --> 承認{"承認が必要ですか?"}
承認 -->|拒否|拒否されました[「実行しないでください」]
承認 -->|許可| execute[「ランタイムはツール・編集・コマンドを実行する」]
実行 --> ストリーム["コマンド・パッチ・ツールイベントのストリーム"]
ストリーム --> レコード["レコードのステータス · 変更されたパス · 証拠ポインタ"]
記録 --> verify{"検証に合格しましたか?"}
確認 -->|はい| success["検証された結果"]
確認 -->|いいえ|ギャップ["障害・ブロッカー・未解決のギャップ"]
読み込み中
インストール
現在の Linux x86_64。 macOS および Windows のビルドは、CI リリース ランナーを通じて計画されています。
T

タグ付きリリースはチェックサム付きバイナリを公開します。チェックアウトから:
scripts/install-elpis.sh
インストーラーはチェックサムを検証し、elpis を ~/.local/bin にアトミックにインストールします。
内部 RAG サービス ( /rag ) は別個の Python サイドカーであり、デフォルトではオフになっています。有効にするには
それ：
scripts/setup-rag.sh
これにより、venv が作成され、config.toml に mcp_servers.elpis-rag エントリが書き込まれます。
このマシン上でリポジトリが実際に存在する場所から計算された絶対パス - 決して
これらのパスを手動で編集し、リポジトリを移動した後、または新しいデバイスでこのスクリプトを再実行します。
OpenAI サブスクリプション ログインがデフォルトです。他のルート:
# OpenRouter (別キー)
import OPENROUTER_API_KEY= " あなたのキー "
elpis --provider openrouter --model " プロバイダー/モデル "
# ネイティブ ベンダー アダプター
import ANTHROPIC_API_KEY= " あなたのキー "
elpis --provider anthropic
import GEMINI_API_KEY= " あなたのキー "
elpis --provider google-gemini
elpis --provider claude|gemini|gemini-flash は OpenRouter 互換性ショートカットです。
上記のネイティブ アダプターとは異なります。
Linux の検証とバイナリ ビルドの実行
.github/workflows/embedded-elpis-linux.yml 。
通常の変更では、重点を置いた初回リリース チェックが実行され、Elpis バイナリが構築されます。網羅的
継承された TUI/アプリサーバー回帰は毎晩手動で実行され、

[切り捨てられた]

## Original Extract

You put an agent into an Elpis, and it becomes Elpis; Be Elpis my friend. - MasihMoafi/Elpis

GitHub - MasihMoafi/Elpis: You put an agent into an Elpis, and it becomes Elpis; Be Elpis my friend. · GitHub
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
MasihMoafi
/
Elpis
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
497 Commits 497 Commits .github/ workflows .github/ workflows codex-rs codex-rs docs docs examples/ gemini-cli examples/ gemini-cli scripts scripts src src tests tests .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md GUIDE.md GUIDE.md LICENSE LICENSE NOTICE NOTICE TASKS.md TASKS.md pyproject.toml pyproject.toml readme.md readme.md View all files Repository files navigation
Your terminal agent forgets everything. Elpis doesn't.
Every coding agent today has the same disease: the longer you work, the worse it gets.
Goals drown in transcripts. Decisions vanish at compaction. Every new session starts with
you explaining your task again , and every model request resends an ever-fatter history
you're paying for.
Elpis is a terminal agent shell that treats context as a managed asset, not an
ever-growing log:
Post-turn context pruning. After a turn delivers its answer, raw exploration —
command dumps, file reads, failed probes — is replaced in the next request by compact
receipts with exact evidence pointers. The full transcript stays on disk, retrievable on
demand; it just stops being resent by default. No other open agent does this per-turn.
The Context Ledger. A visible, scrollable panel showing exactly what enters your
next model request — every rule file, goal, checkpoint, and added file as its own row,
each one toggleable. You decide what the model sees. /add a file of your own.
Portable continuity. The active goal ( GOAL.md ) and a lean checkpoint ( ES.md )
survive compaction, session death, and even switching runtimes. Resume a thread exactly,
or start a fresh one that already knows your goal, last result, and next action —
without replaying history.
Earned memory. Facts become durable memory only after repeated useful recall across
distinct contexts. Everything else stays as searchable evidence. Deleted memories are
archived before reset — never silently lost.
Runtime-agnostic. The shell survives underneath whichever model performs the loop:
ChatGPT/Codex subscription login, OpenRouter, native Anthropic and Gemini adapters,
Bedrock, Ollama, LM Studio. Put a model into Elpis and it becomes Elpis: your goals,
context, memory, and rules.
The execution foundation (terminal UI, patches, permissions, sandboxing, sessions) is a
subtracted fork of OpenAI's Apache-2.0 Codex CLI, hardened by upstream, owned here.
Current state: approaching first release. v0.1.0 is tagged only after live
acceptance recorded in TASKS.md . A green badge means those checks passed —
it never means unfinished work is finished.
Native Ratatui terminal interface with streaming commands, patches, permission modes,
sandboxing, mouse selection, sessions, and compaction.
ChatGPT subscription authentication; OpenRouter through OPENROUTER_API_KEY ; native
Anthropic Messages and Google Gemini adapters ( --provider anthropic ,
--provider google-gemini ; live vendor acceptance pending).
One internal, read-only RAG service: /rag <query> , /rag <path> -- <query> , and
autonomous retrieval.
Portable GOAL.md + ES.md continuity; exact resume or lean continuation.
The Context Ledger with per-file rows and toggles; /usage reporting every admitted
source with size, lifetime, and reason.
Bounded local memory with recall tracking, promotion, provenance, and a
fail-closed archive.
Deterministic first-pass context cleaning: older long tool outputs become bounded
excerpts with durable rollout:// evidence pointers.
The full pruning engine — the agent-authored turn outcome record described in
docs/CONTEXT_AND_SESSIONS.md — is the flagship feature
in active development, alongside a visible per-turn "context saved" metric.
Elpis keeps the surrounding control environment stable while the selected runtime performs the model loop. Exact evidence remains durable; only a small, reasoned working set enters the next request.
flowchart LR
user([User request]) --> tui["Elpis TUI"]
tui --> control["Elpis control layer"]
rag["Read-only RAG"] --> control
portable["GOAL.md + ES.md"] --> control
memory["Bounded local memory"] --> control
control --> runtime{"Select runtime"}
runtime --> openai["OpenAI / Codex"]
runtime --> router["OpenRouter · Anthropic · Gemini"]
runtime --> other["Bedrock / Ollama / LM Studio"]
openai --> execution["Tools · edits · commands"]
router --> execution
other --> execution
execution --> evidence[("Workspace + exact evidence")]
Loading
Context management
Elpis admits rules, the current request, portable state, and relevant memory into a small working set. The Context Ledger and /usage expose why each source is present while full artifacts stay on disk.
The working context is not the transcript. Full conversations, terminal events, and artifacts
remain available as evidence, but are retrieved only when a later task needs them — by exact
evidence pointer first, RAG second. Neither makes the full history a default prompt
attachment. The aim is to make a modest context window sufficient and legible, rather than
pay to resend an ever-growing one.
flowchart LR
rules["Rules + current request"] --> working["Small admitted working set"]
goal["GOAL.md ≤ 6,000 chars"] --> working
checkpoint["ES.md ≤ 8,000 chars"] --> working
memory["Relevant memory"] --> working
status["Context Ledger + /usage"] -.-> working
working --> runtime["Selected runtime request"]
runtime --> results["Tool and function results"]
results --> disk[("Full transcript + artifacts on disk")]
results --> large{"Old tool output > 1,200 chars?"}
large -->|Yes| marker["Compact receipt + evidence pointer"]
large -->|No| keep["Keep in request context"]
Loading
Session continuity
Elpis resumes the useful native thread exactly or starts a lean thread from GOAL.md and ES.md . Pre-compaction synchronization fails closed instead of risking a broken handoff.
flowchart LR
boundary(["Exit · compaction · runtime change"]) --> sync["Synchronize GOAL.md + ES.md"]
sync --> safe{"Files safely synchronized?"}
safe -->|No| stop["Stop compaction + show error"]
safe -->|Yes| thread{"Native thread still useful?"}
thread -->|Yes| exact["Exact resume"]
thread -->|No| lean["Lean continuation"]
exact --> next["Continue from preserved next action"]
lean --> next
next --> evidence["Read exact evidence only when needed"]
evidence --> verify["Verify result + refresh checkpoint"]
Loading
Memory management
New evidence stays searchable until repeated useful recall makes it eligible for durable memory. Durable artifacts are bounded, and deleted or faded lines are archived before reset.
flowchart TB
subgraph promotion["Promotion"]
direction LR
work(["Completed work"]) --> raw["Searchable evidence + provenance"]
raw --> recall["Track recalls + distinct contexts"]
recall --> gate{"3 recalls across 2 contexts?"}
gate -->|No| short["Remain searchable evidence"]
gate -->|Yes| durable["Eligible for MEMORY.md"]
durable --> bounded["MEMORY.md ≤ 30k · summary ≤ 10k"]
end
subgraph retirement["Retirement"]
direction LR
changed{"Deleted or faded?"}
changed -->|No| reset["Reset memory baseline"]
changed -->|Yes| archive["Append removed lines to archive.md"]
archive --> written{"Archive write succeeded?"}
written -->|Yes| reset
written -->|No| block["Block reset to prevent evidence loss"]
end
short --> changed
bounded --> changed
Loading
Read-only RAG
The startup path exposes one minimal read-only tool without loading the retrieval stack. Embeddings and indexing load lazily only after an explicit semantic query.
flowchart LR
launch(["Elpis launches"]) --> host["Minimal stdio MCP host"]
host --> tool["One tool: query_knowledge_base"]
tool --> query{"Explicit RAG query?"}
query -->|No| idle["No indexing or model load"]
query -->|Yes| search["Lazy-load rag.fetch"]
search --> scope{"Workspace or supplied path?"}
scope --> workspace["Launch workspace"]
scope --> supplied["Normalized supplied path"]
workspace --> chunks["Sourced semantic chunks"]
supplied --> chunks
chunks --> exact["Use exact reads before code changes"]
Loading
Safe execution and evidence
Consequential actions pass through visible permission and sandbox policy before execution. Elpis records outcomes and evidence, then distinguishes verified success from failure or unresolved gaps.
flowchart LR
action(["Consequential action proposed"]) --> policy["Show permission + sandbox policy"]
policy --> approval{"Approval required?"}
approval -->|Denied| denied["Do not execute"]
approval -->|Allowed| execute["Runtime executes tool · edit · command"]
execute --> stream["Stream command · patch · tool events"]
stream --> record["Record status · changed paths · evidence pointers"]
record --> verify{"Verification passed?"}
verify -->|Yes| success["Verified outcome"]
verify -->|No| gap["Failure · blocker · unresolved gap"]
Loading
Install
Linux x86_64 today; macOS and Windows builds are planned through CI release runners.
Tagged releases publish a checksummed binary. From a checkout:
scripts/install-elpis.sh
The installer verifies the checksum and installs elpis into ~/.local/bin atomically.
The internal RAG service ( /rag ) is a separate Python sidecar, off by default. To enable
it:
scripts/setup-rag.sh
This creates the venv and writes the mcp_servers.elpis-rag entry in config.toml with
absolute paths computed from wherever the repo actually lives on this machine — never
hand-edit those paths, and re-run this script after moving the repo or on a fresh device.
OpenAI subscription login is the default. Other routes:
# OpenRouter (separate key)
export OPENROUTER_API_KEY= " your-key "
elpis --provider openrouter --model " provider/model "
# Native vendor adapters
export ANTHROPIC_API_KEY= " your-key "
elpis --provider anthropic
export GEMINI_API_KEY= " your-key "
elpis --provider google-gemini
elpis --provider claude|gemini|gemini-flash are OpenRouter compatibility shortcuts,
distinct from the native adapters above.
Linux verification and binary builds run through
.github/workflows/embedded-elpis-linux.yml .
Ordinary changes run focused first-release checks and build the Elpis binary. Exhaustive
inherited TUI/app-server regression runs nightly, manually, and

[truncated]
