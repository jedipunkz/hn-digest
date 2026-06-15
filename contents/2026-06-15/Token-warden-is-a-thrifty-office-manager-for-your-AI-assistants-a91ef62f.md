---
source: "https://github.com/vukkt/token-warden"
hn_url: "https://news.ycombinator.com/item?id=48537725"
title: "Token-warden is a thrifty office manager for your AI assistants"
article_title: "GitHub - vukkt/token-warden: Claude Code plugin that makes coding agents measurably cheaper over time: collect token costs, distill candidate rules, benchmark them on a frozen golden suite, and keep only rules that earn their context rent. · GitHub"
author: "vukkt"
captured_at: "2026-06-15T08:10:00Z"
capture_tool: "hn-digest"
hn_id: 48537725
score: 1
comments: 0
posted_at: "2026-06-15T07:22:18Z"
tags:
  - hacker-news
  - translated
---

# Token-warden is a thrifty office manager for your AI assistants

- HN: [48537725](https://news.ycombinator.com/item?id=48537725)
- Source: [github.com](https://github.com/vukkt/token-warden)
- Score: 1
- Comments: 0
- Posted: 2026-06-15T07:22:18Z

## Translation

タイトル: Token-warden は AI アシスタントのための倹約的なオフィス マネージャーです
記事のタイトル: GitHub - vukkt/token-warden: コーディング エージェントを時間の経過とともに目に見えて安くする Claude Code プラグイン: トークン コストを収集し、候補ルールを抽出し、凍結されたゴールデン スイートでベンチマークを行い、コンテキスト レントを得るルールのみを保持します。 · GitHub
説明: Claude Code プラグインは、コーディング エージェントを時間の経過とともに目に見えて安くします。トークン コストを収集し、候補ルールを抽出し、凍結されたゴールデン スイートでベンチマークを行い、コンテキスト レントを得るルールのみを保持します。 - vukkt/トークンウォーデン

記事本文:
GitHub - vukkt/token-warden: コーディング エージェントを時間の経過とともに目に見えて安くする Claude Code プラグイン: トークン コストを収集し、候補ルールを抽出し、凍結されたゴールデン スイートでベンチマークを行い、コンテキスト レントを得るルールのみを保持します。 · GitHub
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
{{私

セージ }}
ヴクト
/
トークンウォーデン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
35 コミット 35 コミット .claude-plugin .claude-plugin .github .github エージェント エージェント ベンチマーク ベンチマーク コマンド コマンド フック フック src src test test .gitignore .gitignore CHANGELOG.md CHANGELOG.md DECISIONS.md DECISIONS.md ライセンス ライセンス README.md README.md TOKEN-WARDEN-SPEC.md TOKEN-WARDEN-SPEC.md biome.json biome.json package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Claude Code プラグインは、コーディング エージェントを時間の経過とともに目に見えて安くします。
ほとんどの「エージェントの記憶」には、誰も検証しないアドバイスが蓄積されています。 token-warden はエージェントを扱います
エンジニアリング上の問題としてのメモリ: エージェントのコンテキスト内にスペースを必要とするすべてのルールは、
固定ベンチマークで、コストよりも多くのトークンを節約できることを証明するか、そうでない場合は
追い出された。結果は、測定されたルールのみを含むエージェントごとのメモリ ファイルになります。
プラスのリターン。
雰囲気ではなく測定 — すべてのルールには実際のベンチマーク実行からのトークン デルタが含まれます
自己資金 — ルールでは、維持するために独自のコンテキスト レントの 2 倍以上を節約する必要があります
自己監査 - アクティブなルールはラウンドロビンで再ベンチマークされ、ルールが無効になると削除されます。
稼ぐのをやめる
セッション オーバーヘッドなし - コレクションはブロックや失敗を起こさない Stop フックで実行されます。
あなたの仕事
オプティマイザーは 4 段階のフィードフォワード ループです。教訓は完成したものから抽出されます
セッションが実行され、将来のセッションに適用されます。過去の作業がやり直されることはありません。
エージェント セッション (任意のプロジェクト、任意のリポジトリ)
│
│ フックの停止 · トランスクリプトを解析します。
│ トークン、ツール呼び出し、ファイルの再読み取り、完了
▼
┌────

───────────────┐
│ 1・収集 │
│ SQLite ではセッションごとに 1 行 │
━━━━━━━━━━━━━━┘
│
│ 実行が
│ エージェントのローリング p75 トークン コスト
▼
┌─────────────────────┐
│ 2・蒸留 │
│ 廃棄物をめぐる一句の呼びかけ │
│ 統計 → 0-2 候補ルール │
━━━━━━━━━━━━━━┘
│
│ 候補者は SQLite で待機します —
│ 測定するまで決して注射しない
▼
┌─────────────────────┐
│ 3・ベンチ │
│ 凍った備品上の黄金のスイート、│
│ 候補者ありとなしの立候補 │
━━━━━━━━━━━━━━┘
│
│ 測定されたデルタとコンテキストレントの比較
▼
┌─────────────────────┐
│ 4・選択 │
│ 貯蓄額が家賃の 2 倍以上の場合は維持、そうでない場合は維持 │
│ 最も古いルールを削除・再監査する │
━━━━━━━━━━━━━━┘
│
▼
~/.claude/エージェントメモリ/<エージェント>/MEMORY.md
生き残ったルールから大規模にコンパイルされ、注入される
エージェントのシステム プロンプトに次のセッションを入力します
1・集めます。 Stop フックと SubagentStop フックは、ターンごとに起動します (メイン セッションと
サブエージェントはそれぞれ動作します）およびセッション記録を解析します
1つの台帳にr

ow: 入力/出力/キャッシュ トークン (API メッセージ ID によって重複排除 —
トランスクリプトはストリーミングされたブロックごとに使用を繰り返します）、ツール呼び出し数、複数回読み取られたファイル、
セッションが完了したかどうか。フックは2秒のバジェットの下でハードキャップされており、
すべての失敗をラップし、関係なく 0 で終了します。セッションが中断されることはありません。
2・蒸留します。収集は安価ですが、分析はそうではないため、分析は配給制です。実行のみです。
エージェントのローリング 75 パーセンタイル コスト (過去の最低 5 回の実行) を超える部分が抽出されます。あ
単一の独立した Haiku-tier 呼び出しは、無駄な統計と 8 KB のアクション トレースを受け取ります。
そして、厳密な JSON を返す必要があります。つまり、最大 2 つの 1 文の一般化可能なルールです。無効な出力
ドロップされ、再試行されません。既存のルールのほぼ重複 — 削除されたものを含む
それらはトライグラムの類似性によって拒否されるため、改ざんされたルールを再提案することはできません。
3・ベンチ。候補者は、フリーズされたフィクスチャに対してゴールデン タスク スイートで評価されます
リポジトリ (「ベンチマーク システム」を参照)。各構成が実行される
候補を一時的にコンパイルした使い捨てコピー内のスイートをヘッドレスで作成し、
完全に分離されたエージェントのメモリ。
4・選択します。ルールの判定はスペックの不等式です: デルタ = 保存された平均トークン数
完了したゴールデン ランとレンタルごとに = ルール自体のトークン サイズ、ルールがアクティブになります
デルタ × セッション/週 ≥ 2 × レンタル × セッション/週 の場合。以前に成功したタスクの失敗
トークンに関係なく即座に削除されます。セレクターを実行するたびに、
最も最近監査されていないアクティブ ルール - メモリはその場所を獲得し続ける必要があります。生存者は、
MEMORY.md にコンパイルされ、クロード コードによってエージェントのシステム プロンプトに挿入されます。
クロード コード v2.1+ ( クロード --version )
macOS または Linux (WSL 経由の Windows — ベンチマークには POSIX シェルが必要です)
git clone https://github.com/vukkt/token-warden.git

CD トークン監視者
npm install # フックはプラグイン独自の tsx + better-sqlite3 経由で実行されます
2・プラグインをロードする
クロード --plugin-dir /path/to/token-warden
または、永続的にインストールします。このリポジトリは独自のマーケットプレイスでもあります。
/plugin マーケットプレイス add vukkt/token-warden
/plugin install token-warden@vukkt-plugins
Marketplace のインストールは、node_modules なしで ~/.claude/plugins/cache にコピーされます。
Stop フックは、最初の実行時に独自の依存関係をブートストラップします (1 回限りの npm install 、
沈黙）;遅くとも 2 回目のセッションから収集が開始されます。
通常どおり 1 ～ 2 ターン作業してから、次のようにします。
/監視員のステータス
main の実行数が表示されるはずです。すべてのプロジェクトのすべてのセッションが現在進行中です
~/.token-warden/warden.db に測定されます。
4 · ベースラインを凍結します (1 回、エージェントあたり約 20 分)
npm run bench -- --agent all # または一度に 1 つのエージェント
これにより、各エージェントの 3 つのゴールデン タスクが 2 回実行され、永続的な run1_tokens がフリーズされます。
将来のすべての改善要求の分母。これは、ルールが存在する前に 1 回だけ実行してください。
実際の作業には 4 つのサブエージェント (frontend 、 backend 、 sql 、 testing ) を使用します。
高価なセッションは自動的に候補に絞り込まれます。 /warden-status が表示される場合
保留中の候補者を測定します。
npx tsx src/select.ts --agent SQL
アクティブなルールはエージェントのメモリに記録されます。次回のセッションは安く始まります。
コマンド
何をするのか
/監視員のステータス
読み取り専用レポート: エージェントごとの実行/ルール数、スイートの合計と凍結されたベースライン (絶対 + %)、長期にわたる学習曲線、測定されたデルタと来歴を含むアクティブなルール、理由付きの最近のエビクション、プロジェクトごとの実際の作業トークン、エージェント間の質問量
/warden-bench <エージェント|すべて> [--実行数 N] [--タスク ID]
ゴールデン スイートを実行し、 run1 および best と比較し、メタコストのベンチマークをレポートします (週の実作業トークンの 10% を超えると警告します)。
/ウォーデンセレ

ct <エージェント> [--runs N] [--top-up N]
保留中の候補を測定し、それらを削除またはアクティブ化し、最も古いアクティブなルールを再監査し、エージェントのメモリを再コンパイルします。
/warden-modelbench <エージェント> --model <id> [--baseline <id>] [--runs N]
エージェントのゴールデン スイートを 2 つのモデル (候補とエージェントの現在のモデル、ルールは一定) で実行し、そのワークロードでより少ないトークンを使用するレポートを作成します。
/warden-promptbench <エージェント> --variant <file.md> [--runs N]
エージェントのゴールデン スイートを 2 つのプロンプト (さまざまなエージェント定義と出荷されたもの、ルールとモデルは一定) で実行し、使用するトークンが少ないレポートを作成します。
/warden-evolve <エージェント> [--N を実行]
エージェントのプロンプト (モデル呼び出し) のトークンをより安価に書き換えることを提案し、それをベンチマークし、それが勝つことが証明された場合にのみ推奨します。決して自動適用されません。
候補ルールが待機しているときに、軽量の SessionStart フックが 1 行の
新しいセッションへのナッジ — 選択自体は常にユーザーの決定によるものです。
実際のベンチマークトークンを消費します。
セッション終了時にエージェントの費用が異常に高かった場合 (エージェントの最近の中央値の 2 倍以上、
前のセッションが 5 つ以上ある場合)、Stop フックにより、コスト異常のヘッドアップが 1 行で表面化します。
systemMessage 経由で通知します。モデルにはフィードしません (動作ループはありません)。
TOKEN_WARDEN_NO_ALERTS=1 を指定してオプトアウトします。
ヘッドレスまたは名前が衝突する場合は、名前空間形式を使用します
( /token-warden:warden-status )。 CLI に相当するもの:
npx tsx src/status.ts # ステータスレポート
npm run bench -- --agent sql [--rule N] # ベンチマーク ランナー
npx tsx src/select.ts --agent sql # セレクター (測定 + エビクト + コンパイル)
npx tsx src/modelbench.ts --agent sql --model haiku # モデルをエージェントのデフォルトと比較します
npx tsx src/promptbench.ts --agent sql --variant v.md # プロンプト バリアントを出荷されたものと比較します
npx tsx src/evol

ve.ts --agent sql # より安価なプロンプト バリアントを提案および測定します
ベンチマークシステム
測定の精度は制御変数によって決まります。 token-warden がそれらを制御します
積極的に：
フィクスチャ ( benchmarks/fixture/ ) は小さいながらも現実的なフルスタック TypeScript です
プロジェクト — SQLite 上の Express ルート → サービス → リポジトリ、React 管理 UI、
部分的な vitest スイート — フェーズ 2 で凍結され、決して変更されないため、ベースラインは維持されます
月をまたいで比較可能。文書化された意図的な欠陥 ( BUGS.md ) が同梱されています。
エージェントは決して見ることはありません: ベンチマーク ランナーはすべてのコピーからそれを除外します)。
ターゲット。
ゴールデン タスク ( benchmarks/<agent>/golden-NN.md ) — エージェントごとに 3 つ、それぞれがフロントマター
一文プロンプトとシェル success_check (greps および/または完全な
ビテスト実行)。実行は、チェックに合格した場合にのみ完了としてカウントされます: 安価な失敗した実行
高価な成功したものよりも悪く、不完全な実行はすべての実行から除外されます。
貯蓄の計算。
フィクスチャを一時ディレクトリにコピーします (node_modules はシンボリックリンクされています。BUGS.md は除外されています)。
エージェント定義をコピーにインストールし、メモリ スコープを次のように書き換えます。
project なので、テスト対象のコンパイル済み MEMORY.md は temp ディレクトリ内で解決されます。
実際のエージェントのメモリは、ベンチマークによって読み書きされることはありません。
テスト対象のルール セット (アクティブなルール ± 1 つの候補) をそのメモリにコンパイルします。
スコープ付き権限を使用して claude -p --agent <name> をヘッドレスで実行します: acceptEdits
さらに、テスト コマンドのみの Bash ホワイトリストを追加します。決して bypassPermissions を使用しないでください。
success_check を実行します。トランスクリプトを解析します。 1 つの実行行を記録します。
(エージェント、タスク) ごとに最初に完了した実行では、baselines.run1_tokens が永久にフリーズします。
後で完了した実行は、best_tokens を下方にラチェットするだけです。
多様性と誠実さ。各構成は 2 回実行され、実行のペアは次のように異なります。
25%以上

出力にフラグが立てられます。 LLM 分散は、次の場合の主な誤差原因です。
効果サイズが小さい — 以下の記録されたデモは、ルールを排除することを示しています。の
セレクターは分散を認識します: st を計算します。

[切り捨てられた]

## Original Extract

Claude Code plugin that makes coding agents measurably cheaper over time: collect token costs, distill candidate rules, benchmark them on a frozen golden suite, and keep only rules that earn their context rent. - vukkt/token-warden

GitHub - vukkt/token-warden: Claude Code plugin that makes coding agents measurably cheaper over time: collect token costs, distill candidate rules, benchmark them on a frozen golden suite, and keep only rules that earn their context rent. · GitHub
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
vukkt
/
token-warden
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
35 Commits 35 Commits .claude-plugin .claude-plugin .github .github agents agents benchmarks benchmarks commands commands hooks hooks src src test test .gitignore .gitignore CHANGELOG.md CHANGELOG.md DECISIONS.md DECISIONS.md LICENSE LICENSE README.md README.md TOKEN-WARDEN-SPEC.md TOKEN-WARDEN-SPEC.md biome.json biome.json package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts View all files Repository files navigation
A Claude Code plugin that makes coding agents measurably cheaper over time.
Most "agent memory" accumulates advice nobody ever verifies. token-warden treats agent
memory as an engineering problem: every rule that wants space in an agent's context must
prove, on a fixed benchmark, that it saves more tokens than it costs — or it gets
evicted. The result is a per-agent memory file containing only rules with measured,
positive return.
Measured, not vibes — every rule carries a token delta from real benchmark runs
Self-funding — rules must save ≥ 2× their own context rent to stay
Self-auditing — active rules are re-benchmarked round-robin and evicted when they
stop earning
Zero session overhead — collection runs in a Stop hook that never blocks or fails
your work
The optimizer is a four-stage, feed-forward loop. Lessons are extracted from finished
sessions and applied to future ones — past work is never re-done.
agent session (any project, any repo)
│
│ Stop hook · parses the transcript:
│ tokens, tool calls, file re-reads, completion
▼
┌─────────────────────────────────────┐
│ 1 · COLLECT │
│ one row per session in SQLite │
└─────────────────────────────────────┘
│
│ fires only when a run exceeds the
│ agent's rolling p75 token cost
▼
┌─────────────────────────────────────┐
│ 2 · DISTILL │
│ one haiku call over the waste │
│ stats → 0–2 candidate rules │
└─────────────────────────────────────┘
│
│ candidates wait in SQLite —
│ never injected until measured
▼
┌─────────────────────────────────────┐
│ 3 · BENCH │
│ golden suite on a frozen fixture, │
│ run with vs. without the candidate │
└─────────────────────────────────────┘
│
│ measured delta vs. context rent
▼
┌─────────────────────────────────────┐
│ 4 · SELECT │
│ keep if savings ≥ 2× rent, else │
│ evict · re-audit the oldest rule │
└─────────────────────────────────────┘
│
▼
~/.claude/agent-memory/<agent>/MEMORY.md
compiled wholesale from surviving rules and injected
into the agent's system prompt next session
1 · Collect. Stop and SubagentStop hooks fire after every turn (main session and
subagent work respectively) and parse the session transcript
into one ledger row: input/output/cache tokens (deduplicated by API message id — the
transcript repeats usage per streamed block), tool-call count, files read more than once,
and whether the session completed. The hook is hard-capped under the 2-second budget,
wraps every failure, and exits 0 regardless — it can never break your session.
2 · Distill. Collection is cheap, analysis is not, so analysis is rationed: only runs
above the agent's rolling 75th-percentile cost (minimum 5 prior runs) are distilled. A
single detached haiku-tier call receives the waste statistics plus an 8 KB action trace
and must return strict JSON: at most two one-sentence, generalizable rules. Invalid output
is dropped, never retried. Near-duplicates of any existing rule — including evicted
ones — are rejected by trigram similarity, so a falsified rule cannot be re-proposed.
3 · Bench. Candidates are measured on a golden task suite against a frozen fixture
repository (see The benchmark system ). Each configuration runs
the suite headlessly in a throwaway copy with the candidate compiled into a temporary,
fully isolated agent memory.
4 · Select. A rule's verdict is the spec inequality: with delta = mean tokens saved
per completed golden run and rent = the rule's own size in tokens, the rule goes active
iff delta × sessions/week ≥ 2 × rent × sessions/week . Failing a previously-passing task
is instant eviction regardless of tokens. Every selector run also re-benchmarks the
least-recently-audited active rule — memory must keep earning its place. Survivors are
compiled into MEMORY.md , which Claude Code injects into the agent's system prompt.
Claude Code v2.1+ ( claude --version )
macOS or Linux (Windows via WSL — benchmarks need a POSIX shell)
git clone https://github.com/vukkt/token-warden.git
cd token-warden
npm install # the hooks run via the plugin's own tsx + better-sqlite3
2 · Load the plugin
claude --plugin-dir /path/to/token-warden
Or install persistently — this repository is also its own marketplace:
/plugin marketplace add vukkt/token-warden
/plugin install token-warden@vukkt-plugins
Marketplace installs are copied to ~/.claude/plugins/cache without node_modules .
The Stop hook bootstraps its own dependencies on first run (one-time npm install ,
silent); collection begins from the second session at the latest.
Work normally for a turn or two, then:
/warden-status
You should see a runs count for main . Every session in every project is now being
measured into ~/.token-warden/warden.db .
4 · Freeze the baselines (one-time, ~20 min per agent)
npm run bench -- --agent all # or one agent at a time
This runs each agent's three golden tasks twice and freezes run1_tokens — the permanent
denominator of every future improvement claim. Do this once, before any rules exist.
Use the four subagents ( frontend , backend , sql , testing ) for real work.
Expensive sessions distill into candidates automatically. When /warden-status shows
candidates pending, measure them:
npx tsx src/select.ts --agent sql
Active rules land in the agent's memory; the next session starts cheaper.
Command
What it does
/warden-status
Read-only report: per-agent run/rule counts, suite total vs. frozen baseline (absolute + %), learning curve over time, active rules with measured deltas and provenance, recent evictions with reasons, real-work tokens by project, cross-agent question volume
/warden-bench <agent|all> [--runs N] [--task id]
Runs the golden suite, compares against run1 and best , and reports benchmarking meta-cost (warns above 10% of the week's real-work tokens)
/warden-select <agent> [--runs N] [--top-up N]
Measures pending candidates, evicts or activates them, re-audits the oldest active rule, and recompiles the agent's memory
/warden-modelbench <agent> --model <id> [--baseline <id>] [--runs N]
Runs the agent's golden suite under two models (candidate vs. the agent's current model, rules held constant) and reports which uses fewer tokens for that workload
/warden-promptbench <agent> --variant <file.md> [--runs N]
Runs the agent's golden suite under two prompts (a variant agent definition vs. the shipped one, rules and model held constant) and reports which uses fewer tokens
/warden-evolve <agent> [--runs N]
Proposes a token-cheaper rewrite of the agent's prompt (model call), benchmarks it, and recommends it only if it provably wins — never auto-applied
When candidate rules are waiting, a lightweight SessionStart hook injects a one-line
nudge into new sessions — selection itself always stays a user decision, because it
spends real benchmark tokens.
When a session ends unusually expensive for its agent (≥ 2× the agent's recent median,
given ≥ 5 prior sessions), the Stop hook surfaces a one-line cost-anomaly heads-up to
you via systemMessage — it informs, it does not feed the model (no behavioral loop).
Opt out with TOKEN_WARDEN_NO_ALERTS=1 .
Headless or when names collide, use the namespaced forms
( /token-warden:warden-status ). CLI equivalents:
npx tsx src/status.ts # status report
npm run bench -- --agent sql [--rule N] # benchmark runner
npx tsx src/select.ts --agent sql # selector (measure + evict + compile)
npx tsx src/modelbench.ts --agent sql --model haiku # compare a model against the agent's default
npx tsx src/promptbench.ts --agent sql --variant v.md # compare a prompt variant against the shipped one
npx tsx src/evolve.ts --agent sql # propose + measure a cheaper prompt variant
The benchmark system
Measurement is only as good as its control variables. token-warden controls them
aggressively:
The fixture ( benchmarks/fixture/ ) is a small but realistic full-stack TypeScript
project — Express routes → services → repositories over SQLite, a React admin UI, a
partial vitest suite — frozen at Phase 2 and never modified , so baselines stay
comparable across months. It ships with documented, deliberate flaws ( BUGS.md , which
agents never see: the benchmark runner excludes it from every copy) that the golden tasks
target.
Golden tasks ( benchmarks/<agent>/golden-NN.md ) — three per agent, each a frontmatter
file with a one-sentence prompt and a shell success_check (greps and/or a full
vitest run ). A run only counts as completed if its check passes: a cheap failed run
is worse than an expensive successful one, and incomplete runs are excluded from all
savings math.
Copy the fixture to a temp dir ( node_modules symlinked; BUGS.md excluded).
Install the agent definition into the copy with its memory scope rewritten to
project , so the compiled MEMORY.md under test resolves inside the temp dir —
real agent memory is never read or written by benchmarks.
Compile the rule set under test (active rules ± one candidate) into that memory.
Run claude -p --agent <name> headlessly with scoped permissions : acceptEdits
plus a Bash allowlist of test commands only — never bypassPermissions .
Run the success_check ; parse the transcript; record one runs row.
First-ever completed run per (agent, task) freezes baselines.run1_tokens forever;
later completed runs only ratchet best_tokens downward.
Variance and honesty. Each configuration runs twice and pairs of runs differing by
more than 25% are flagged in the output. LLM variance is the dominant error source at
small effect sizes — the recorded demonstration below shows it evicting a rule. The
selector is variance-aware: it computes the st

[truncated]
