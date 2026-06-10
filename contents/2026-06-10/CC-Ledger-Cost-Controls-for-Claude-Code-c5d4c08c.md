---
source: "https://ccledger.dev"
hn_url: "https://news.ycombinator.com/item?id=48479838"
title: "CC-Ledger: Cost Controls for Claude Code"
article_title: "cc-ledger — Every prompt has a price."
author: "tejpalv"
captured_at: "2026-06-10T19:00:39Z"
capture_tool: "hn-digest"
hn_id: 48479838
score: 2
comments: 1
posted_at: "2026-06-10T17:38:53Z"
tags:
  - hacker-news
  - translated
---

# CC-Ledger: Cost Controls for Claude Code

- HN: [48479838](https://news.ycombinator.com/item?id=48479838)
- Source: [ccledger.dev](https://ccledger.dev)
- Score: 2
- Comments: 1
- Posted: 2026-06-10T17:38:53Z

## Translation

タイトル: CC-Ledger: クロードコードのコスト管理
記事のタイトル: cc-ledger — すべてのプロンプトには価格があります。
説明: cc-ledger は、リポジトリ内のすべてのクロード コード、カーソル、およびコパイロット セッションをキャプチャし、それをユーザーごと、リポジトリごとのメトリクス (トークン コスト、ライン アトリビューション、スループット、ストリーク) に変換します。そのため、エンジニアリングのリーダーは、エージェントが実際に何をしているのかを知っています。

記事本文:
cc-ledger — すべてのプロンプトには価格が付いています。 cc / 元帳 サインイン AI コーディング テレメトリ すべてのプロンプトには価格が付いています。
cc-ledger は両方を示します。
cc-ledger は、リポジトリ内のすべてのクロード コード、カーソル、およびコパイロット セッションをキャプチャし、エンジニアリング リーダーが尋ねる 3 つの質問に答えます。PR の実際のコストはいくらか、どのセッションが暴走するのか、その費用はどこに使われるのか (計画かコーディングか) です。
GitHub でソースを見る events.livestreaming 12:04:23 →prompt.start cc-ledger/api 12:04:24 ·tool.use bash 12:04:27 ·tool.use edit src/auth.rs 12:04:31 $cost +$0.0142 12:04:32 +lines.ai +47 12:04:48 ·tool.use bash 12:05:02 ⇒ commit 7e2c1a3 12:05:11 ⇒ pr.opened #221 12:05:12 $cost.session $1.84 12:05:13 ◇ session.end 14m32s 12:04:23 →prompt.start cc-ledger/api 12:04:24 · tools.use bash 12:04:27 ·tool.use edit src/auth.rs 12:04:31 $cost +$0.0142 12:04:32 +lines.ai +47 12:04:48 ·tool.use bash 12:05:02 ⇒ commit 7e2c1a3 12:05:11 ⇒ pr.opened #221 12:05:12 $cost.session $1.84 12:05:13 ◇ session.end 14m32s メトリクス cc-ledger が追跡するもの
トークンの支出はマージされた PR ごとにロールアップされます。キャッシュ読み取りは 10%、バッチ層が適用されます。リポジトリ、チーム、または作成者ごとにドリルインします。
セッションコストの分布 - p50、p95、p99 - したがって、請求書に表示される前にテールが表示されます。クリックして違反者にアクセスします。
P95 $3.10 P99 $18.40 カテゴリ · 分割 ドルの行き先
エージェントは計画、コーディング、デバッグ、レビューなどを自動分類するため、リーダーは支出が探索なのか実行なのかを確認できます。
ディレクター、副社長、チームリーダーは、AI への支出は正当なのか、そしてどのチームが活用されるのかという 1 つの質問に答える必要があります。
cc-ledger はダッシュボードであり、生のログではありません。 JSONL トランスクリプトによる洞窟探索はありません。ページを開いて、PR あたりのコスト、ランナウェイ テール、計画の終了点を確認します。

dコーディングが始まります。任意のユーザー、任意のリポジトリを毎週ドリルダウンします。
1 つのコマンドで、Claude Code ライフサイクル フックをリポジトリに接続します。デーモンや SaaS ping はなく、すべてがローカルで実行されます。
$カール -fsSL https://ccledger.dev/install | bash コピー 02 キャプチャされたセッション
すべてのプロンプト、すべてのツール呼び出し、すべての差分はローカル SQLite 台帳に記録され、(オプションで) コードと一緒にプライベート git ブランチにコミットされます。
GitHub アプリを承認し、cc-ledger をリポジトリに指定すると、ダッシュボードがリアルタイムで表示されます。コスト、アトリビューション、スループット — ユーザーごと、リポジトリごと、週ごと。

## Original Extract

cc-ledger captures every Claude Code, Cursor, and Copilot session in your repos and turns it into per-user, per-repo metrics — token cost, line attribution, throughput, streaks. So engineering leaders know what their agents are actually doing.

cc-ledger — Every prompt has a price. cc / ledger Sign in ai coding telemetry Every prompt has a price.
cc-ledger shows you both.
cc-ledger captures every Claude Code, Cursor, and Copilot session in your repos and answers the three questions every engineering leader is asking: what does a PR actually cost, which sessions are runaway burns, and where does the spend go — planning or coding?
View source on GitHub events.live streaming 12:04:23 → prompt.start cc-ledger/api 12:04:24 · tool.use bash 12:04:27 · tool.use edit src/auth.rs 12:04:31 $ cost +$0.0142 12:04:32 + lines.ai +47 12:04:48 · tool.use bash 12:05:02 ⇒ commit 7e2c1a3 12:05:11 ⇒ pr.opened #221 12:05:12 $ cost.session $1.84 12:05:13 ◇ session.end 14m32s 12:04:23 → prompt.start cc-ledger/api 12:04:24 · tool.use bash 12:04:27 · tool.use edit src/auth.rs 12:04:31 $ cost +$0.0142 12:04:32 + lines.ai +47 12:04:48 · tool.use bash 12:05:02 ⇒ commit 7e2c1a3 12:05:11 ⇒ pr.opened #221 12:05:12 $ cost.session $1.84 12:05:13 ◇ session.end 14m32s metrics What cc-ledger tracks
Token spend rolled up per merged PR — cache reads at 10%, batch tier applied. Drill in by repo, team, or author.
Session-cost distribution — p50, p95, p99 — so the tail is visible before it shows up on the invoice. Click through to the offender.
P95 $3.10 P99 $18.40 category · split Where the dollars go
Agent turns auto-classified — planning, coding, debugging, review — so leaders see whether spend is exploration or execution.
Directors, VPs, and team leads need to answer one question: is the AI spend justified, and which teams get leverage?
cc-ledger is the dashboard, not the raw logs. No spelunking through JSONL transcripts. Open the page, see cost per PR, the runaway tail, and where planning ends and coding begins. Drill into any user, any repo, any week.
One command wires Claude Code lifecycle hooks into your repo. No daemons, no SaaS pings — everything runs locally.
$ curl -fsSL https://ccledger.dev/install | bash copy 02 Sessions captured
Every prompt, every tool call, every diff is recorded into a local SQLite ledger and (optionally) committed to a private git branch alongside your code.
Authorize the GitHub App, point cc-ledger at your repos, and the dashboard fills in real time. Cost, attribution, throughput — per user, per repo, per week.
