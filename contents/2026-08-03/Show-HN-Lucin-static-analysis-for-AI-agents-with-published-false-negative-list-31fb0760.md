---
source: "https://lucin.pages.dev/"
hn_url: "https://news.ycombinator.com/item?id=49156727"
title: "Show HN: Lucin -static analysis for AI agents with published false-negative list"
article_title: "Lucin — Your agent can be talked into anything"
author: "Madhav23"
captured_at: "2026-08-03T15:38:22Z"
capture_tool: "hn-digest"
hn_id: 49156727
score: 1
comments: 0
posted_at: "2026-08-03T15:03:08Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Lucin -static analysis for AI agents with published false-negative list

- HN: [49156727](https://news.ycombinator.com/item?id=49156727)
- Source: [lucin.pages.dev](https://lucin.pages.dev/)
- Score: 1
- Comments: 0
- Posted: 2026-08-03T15:03:08Z

## Translation

タイトル: Show HN: Lucin - 公開された偽陰性リストによる AI エージェントの静的分析
記事のタイトル: Lucin — エージェントは何でも話しかけることができます
説明: 他のスキャナーはツール名を読み取り、推測します。 Lucin は、AI エージェントが呼び出すことができるすべてのツール内のコードを読み取り、汚染されたプロンプトから建物の外に出るデータまでの正確なパスを追跡します。無料、オープンソース、API キーなし。

記事本文:
コンテンツにスキップ
ルシン
ドキュメント
ルール
ベンチマーク
限界
比較する
ブログ
GitHub
エージェントをスキャンする
ドキュメント
ルール
ベンチマーク
限界
比較する
ブログ
GitHub
エージェントをスキャンする
AI エージェントのセキュリティ
エージェントは何でも話し合うことができます。
Lucin がそれにかかる費用を教えてくれます。
他のスキャナはツール名を読み取って推測します。 Lucin は、エージェントが呼び出すことができるすべてのツール内のコードを読み取り、有害なプロンプトから建物を出るデータまでの正確なパスを追跡し、それを閉じるカットを渡します。
read_email
クエリ_顧客
__llm__
post_webhook
スキャンツール本体…
|
AG-トライフェクタ · クリティカル
1 つのモデルで 3 つの使い方が可能
ほとんどのスキャナーは推測します。ルーシンが証明します。
エージェントのモデルの 1 つである情報フロー グラフでは、3 つの方法が使用されていました。
ツール内の実際のコードを読み取ります。 file:line を使用して、信頼できない入力から危険なアクションまでのすべてのパスをマップします。コマンド 1 つで、サインアップは不要です。ほとんどのチームにとって、これは製品全体であり、永久に無料です。
パス SCAN にフラグが立てられ、ライブで強制されます。プロンプトではなくフロー上の決定的なゲートなので、モデルが完全に侵害されても流出することはありません。
自分のツールに対して攻撃が生成されるため、発見はそれが本物であるという証拠と、修正されたという証拠を伴います。
一度読んだだけで行動できる発見。
重大度、証明証のパス、それを閉じる正確なカット、およびコード行。 3 つの表面でも同じ結果が得られました。
$ ルシンスキャン ./support-agent/
ターゲット: ./support-agent/ (エージェント 2 つ、ツール 14 つ、MCP サーバー 1 つ)
スキャンは 842 ミリ秒で完了しました
── セキュリティスコア ─────────────
████████████░░░░░░░░ 62/100 — 注意が必要
── リスクの概要 ─────────────
クリティカル███ 1

高██████ 2
ミディアム ██████ 3
── CRITICAL・AG-TRIFECTA ──────────
信頼できない入力が外部シンクに到達する
エージェント: support_agent ツール: post_webhook
証明:
コントロール: read_email → __llm__ → post_webhook
データ: query_customers → __llm__ → post_webhook
最小カットの修正: すべての exfil パスを壊すために 1 つのツールを制限します
post_webhook (ホストを許可リストに登録するか、承認が必要です)
OWASP: LLM06 過剰な主体性
場所: エージェント/support.py:88
lucin-report.html ルシンスキャン 。 --html形式
62 /100
注意が必要です
1 クリティカル
2 高
3 中
4 低
2エージェント・14ツール
1 MCP サーバー · 842ms
クリティカル
信頼できない入力が外部シンクに到達する
AG-トリフェクタ
メールを読む→
__llm__ →
post_webhook
Min-cut 修正: post_webhook を許可リストに登録されたホストに制限するか、起動する前に承認を要求します。
エージェント/support.py:88
高
ツールは境界なしでシェル入力を実行します
AG-011
高
シークレットを含む環境読み取りが取得から到達可能
AG-002
+ スクロールせずに見える範囲にさらに 7 つのインタラクティブなフロー グラフ
.github/workflows/security.yml
出口 1 — ブロックされています
- 使用: Madhav2310/lucinlabs @v1
と:
フェイルオン: クリティカル
フォーマット: サリフ
─────────────────
エラー: AG-TRIFECTA · Agents/support.py:88
信頼できない入力が外部シンクに到達する
read_email → __llm__ → post_webhook
クリティカル 1 件、高 2 件 — 549 合格 — SARIF がコード スキャンにアップロードされました
プロセスは終了コード 1 で完了しました。
ベンチマーク
領収書を発行しております。
上記のすべての数値は、コミットされたコマンドから再生成されます。ご自身で実行してください。当社のマーケティングを信頼するよりも、数字を再現していただきたいと考えています。足りないもの→
17,600 のアクション。 2日半です。流れを見ている人は誰もいない。
ハグファ

CE、2026 年 7 月。洗練されたエクスプロイトではありません。1 つの悪いエッジがあり、何も監視していなかったために 4 日間のログにわたって約 17,600 回走査されました。出荷するエージェントはすべて同じスケルトンを持っています。
誰よりも早く GUARD を手に入れましょう。
少数のデザインパートナーを募集しています。あなたが何を実行しているのか教えてください。 1日以内に返信させていただきます。

## Original Extract

Other scanners read tool names and guess. Lucin reads the code inside every tool your AI agent can call and traces the exact path from a poisoned prompt to your data leaving the building. Free, open source, no API key.

Skip to content
LUCIN
Docs
Rules
Benchmarks
Limits
Compare
Blog
GitHub
Scan your agent
Docs
Rules
Benchmarks
Limits
Compare
Blog
GitHub
Scan your agent
SECURITY FOR AI AGENTS
Your agent can be talked into anything.
Lucin shows you what that costs.
Other scanners read tool names and guess. Lucin reads the code inside every tool your agent can call and traces the exact path from a poisoned prompt to your data leaving the building — then hands you the cut that closes it.
read_email
query_customers
__llm__
post_webhook
SCANNING TOOL BODIES…
|
AG-TRIFECTA · CRITICAL
ONE MODEL, THREE WAYS TO USE IT
Most scanners guess. Lucin proves.
One model of the agent — the information-flow graph — used three ways.
Reads the real code inside your tools. Maps every path from untrusted input to a dangerous action, with file:line . One command, no signup. For most teams this is the whole product, and it's free forever.
The path SCAN flagged, enforced live. A deterministic gate on the flow — not the prompt — so the model can be fully compromised and still not exfiltrate.
Attacks generated against your own tools, so a finding ships with proof it's real — and proof it's fixed.
A finding you can act on in one read.
Severity, the proof-witness path, the exact cut that closes it, and the line of code. Same finding, three surfaces.
$ lucin scan ./support-agent/
Target: ./support-agent/ (2 agents, 14 tools, 1 MCP server)
Scan completed in 842ms
── SECURITY SCORE ─────────────────────────────
████████████░░░░░░░░ 62/100 — Needs attention
── RISK SUMMARY ───────────────────────────────
CRITICAL ███ 1
HIGH ██████ 2
MEDIUM ██████ 3
── CRITICAL · AG-TRIFECTA ─────────────────────
Untrusted input reaches an external sink
Agent: support_agent Tool: post_webhook
Proof:
control: read_email → __llm__ → post_webhook
data: query_customers → __llm__ → post_webhook
Min-cut fix: restrict 1 tool to break every exfil path
post_webhook (allow-list hosts, or require approval)
OWASP: LLM06 Excessive Agency
Location: agents/support.py:88
lucin-report.html lucin scan . --format html
62 /100
Needs attention
1 Critical
2 High
3 Medium
4 Low
2 agents · 14 tools
1 MCP server · 842ms
CRITICAL
Untrusted input reaches an external sink
AG-TRIFECTA
read_email →
__llm__ →
post_webhook
Min-cut fix: restrict post_webhook to an allow-listed host, or require approval before it fires.
agents/support.py:88
HIGH
Tool executes shell input without a boundary
AG-011
HIGH
Secret-bearing env read reachable from retrieval
AG-002
+ 7 more interactive flow graph below the fold
.github/workflows/security.yml
exit 1 — blocked
- uses: Madhav2310/lucinlabs @v1
with:
fail-on: critical
format: sarif
──────────────────────────────────────────────
Error : AG-TRIFECTA · agents/support.py:88
Untrusted input reaches an external sink
read_email → __llm__ → post_webhook
1 critical, 2 high — 549 passing — SARIF uploaded to code scanning
Process completed with exit code 1.
BENCHMARKS
We publish the receipts.
Every number above regenerates from a committed command. Run them yourself — we'd rather you reproduce the numbers than trust our marketing. What it misses →
17,600 actions. Two and a half days. Nobody watching the flow.
Hugging Face, July 2026. Not a sophisticated exploit — one bad edge, traversed roughly 17,600 times over four days of logs because nothing was watching. Every agent you ship has the same skeleton.
Get GUARD before everyone else.
We're taking a small number of design partners. Tell us what you're running; we reply within a day.
