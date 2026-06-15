---
source: "https://polis-protocol.vercel.app/#"
hn_url: "https://news.ycombinator.com/item?id=48535529"
title: "Polis – the #1 tool for AI agent coordination"
article_title: "POLIS — control plane for coding agents"
author: "lucius_gc"
captured_at: "2026-06-15T04:37:58Z"
capture_tool: "hn-digest"
hn_id: 48535529
score: 2
comments: 0
posted_at: "2026-06-15T01:43:49Z"
tags:
  - hacker-news
  - translated
---

# Polis – the #1 tool for AI agent coordination

- HN: [48535529](https://news.ycombinator.com/item?id=48535529)
- Source: [polis-protocol.vercel.app](https://polis-protocol.vercel.app/#)
- Score: 2
- Comments: 0
- Posted: 2026-06-15T01:43:49Z

## Translation

タイトル: Polis – AI エージェント調整のための #1 ツール
記事のタイトル: POLIS — コーディング エージェント用のコントロール プレーン
説明: Claude Code、Codex、Gemini、および Cursor を 1 つのチームとして、リポジトリ内、Git 内、すべてのベンダーにわたって実行します。すべてのタスクにはオーナーがいて、すべての引き継ぎには証拠があり、チームは目に見えて自らの間違いを繰り返すのをやめます。

記事本文:
ポリス // コントロールプレーン // REV 2.0 // MIT
01_問題 02_証明 03_インストールの手順 →
01
02
03
04
05
[ ローカルファースト · git ネイティブ · MCP ]
コーディング エージェントを 1 つのチームのように調整します。
Claude Code、Codex、Gemini、および Cursor を 1 つのリポジトリに対して実行します。すべてのタスクにはオーナーがいて、すべての引き継ぎには証拠があり、チームは Git 内で、すべてのベンダーにわたって、自らの間違いを繰り返すのを目に見えて止めています。
70% 0%
メモリーレス 65% ポリス 8%
単一エージェント / 管理されていない群集ポリス · 蓄積 + 再注入
[ 01 ] 問題 // 障害モード
01.A — 衝突
2 つのエージェント、1 つのファイルを同時に実行。
Claude と Codex は両方とも src/auth/login.py を開きます。一方が他方を静かに上書きします。作業が失われるか、午後はマージを解く作業に移ります。プレーンなリポジトリでは、調整は運任せであり、並列ワーカーを追加するたびに衝突が増加します。これは最悪の規模の失敗です。
すべてのセッションはゼロから始まります。同じ問題を毎回学び直しました。
ベンダーを切り替えると、コンテキストが失われます。ツール間でメモリを共有することはありません。
契約が成立しました
申し立て済み + ファイルは予約済み
証拠により解決
レッスン + ガードレール
次のタスクがプリロードされています
一致する将来のタスクに再注入されます
−88%
エラーの繰り返し · 65% → 8%
Polis は、各失敗を教訓またはガードレールとして記録し、それを将来の対応するタスクに再注入します。反復エラー率はメモリレス設定よりも 88% 低くなり、各障害クラスが最大 1 回再発します。単独のエージェントではそのような曲線を描くことはできません。それは決して蓄積されません。
コードに触れる前に、エージェントはファイルを予約します。重複する予約は完全に拒否されます。モデルの判断や人種の判断はありません。単一のエージェントや管理されていない群れでは構造的に提供できない保証です。
ワークスペースを作成します。エージェントを国民として登録します。
コントラクトには、所有者、受け入れ基準、および透明なルーティング理由が与えられます。
作業ファイルのレッスンが完了しました

n;次に一致するタスクが事前ロードされて開始されます。
+ 2 人以上のエージェントが 1 つの実際のリポジトリにアクセスするため、衝突が発生するリスクがあります
+ ベンダーを切り替えると、そのたびにコンテキストが失われます
+ あなたはチームが間違いを繰り返すのをやめてほしいと考えています — おそらく
+ 誰が何をしたのか、なぜ行ったのか監査可能な記録が必要です
− 1 人の適切なプロンプトを備えたエージェントがすでにその仕事を行っています
− 能力カードは正確であり、タスクが逸脱することはありません
− ホストされたランタイムが必要です — ポリス座標、実行されません

## Original Extract

Run Claude Code, Codex, Gemini and Cursor as one team — in your repo, in git, across every vendor. Every task has an owner, every handoff has evidence, and the team measurably stops repeating its own mistakes.

POLIS // CONTROL PLANE // REV 2.0 // MIT
01_Problem 02_Proof 03_Sequence Install →
01
02
03
04
05
[ Local-first · git-native · MCP ]
Coordinate your coding agents like one team.
Run Claude Code, Codex, Gemini and Cursor against one repo. Every task has an owner, every handoff has evidence, and the team measurably stops repeating its own mistakes — in git, across every vendor.
70% 0%
MEMORYLESS 65% POLIS 8%
SINGLE AGENT / UNMANAGED SWARM POLIS · ACCUMULATES + RE-INJECTS
[ 01 ] The problem // FAILURE MODES
01.A — COLLISION
Two agents, one file, at once.
Claude and Codex both open src/auth/login.py . One silently overwrites the other; work is lost, or the afternoon goes to untangling a merge. A plain repo leaves coordination to luck — and collisions grow with every parallel worker you add. This is the failure that scales worst.
Every session starts at zero. The same gotcha, re-learned every time.
Switch vendor, lose the context. No shared memory across tools.
CONTRACT OPENED
CLAIMED + FILES RESERVED
SETTLED WITH EVIDENCE
LESSON + GUARDRAIL
NEXT TASK PRE-LOADED
RE-INJECTED INTO MATCHING FUTURE TASKS
−88%
Repeat errors · 65% → 8%
Polis records each failure as a lesson or guardrail and re-injects it into matching future tasks. The repeat-error rate falls 88% below a memoryless setup — each failure class recurs at most once. A lone agent can't draw that curve; it never accumulates.
Before touching code, an agent reserves the files. An overlapping reservation is rejected outright — no model judgement, no race. The guarantee a single agent or an unmanaged swarm structurally cannot give.
Create the workspace; register your agents as citizens.
Contracts get an owner, acceptance criteria, and a transparent routing reason.
Settled work files a lesson; the next matching task starts pre-loaded with it.
+ 2+ agents touch one real repo and collisions are a risk
+ You switch between vendors and lose context each time
+ You want the team to stop repeating mistakes — provably
+ You need an auditable record of who did what, and why
− One well-prompted agent already does the job
− Your capability cards are accurate and tasks never drift
− You want a hosted runtime — Polis coordinates, it doesn't execute
