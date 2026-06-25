---
source: "https://hezo.ai"
hn_url: "https://news.ycombinator.com/item?id=48669046"
title: "Show HN: Hezo – Self-hosted teams of AI agents that never see your real secrets"
article_title: "Hezo — A whole AI workforce. And you're the boss."
author: "hiddentao"
captured_at: "2026-06-25T05:19:40Z"
capture_tool: "hn-digest"
hn_id: 48669046
score: 1
comments: 0
posted_at: "2026-06-25T04:50:17Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Hezo – Self-hosted teams of AI agents that never see your real secrets

- HN: [48669046](https://news.ycombinator.com/item?id=48669046)
- Source: [hezo.ai](https://hezo.ai)
- Score: 1
- Comments: 0
- Posted: 2026-06-25T04:50:17Z

## Translation

タイトル: Show HN: Hezo – あなたの本当の秘密を決して見ることのない AI エージェントの自己ホスト型チーム
記事のタイトル: Hezo — AI 労働力全体。そしてあなたは上司です。
説明: AI エージェントのチーム全体を会社のように運営します。目標を設定すると、彼らが仕事を出荷します。組織図、プロジェクト、予算、長期記憶、承認はすべて自分が所有するハードウェア上にあります。

記事本文:
Hezo — AI 労働力全体。そしてあなたは上司です。 hezo Docs GitHub はじめる AI 従業員全体。
そしてあなたは上司です。
Hezo は AI エージェントを雇用し、実行し、その作品を送信します。キーを渡す必要はありません。
クラウドへの依存関係なし シークレットはエージェント コンテキストに含まれない 署名済みコミット エージェントごとの予算上限 ​​01 仕組み 3 つの作業が作業チームに移行します。
CEOに仕事の内容を説明してください。プロジェクトの範囲を定め、チームをそれぞれ独自のコンテナー内でプロビジョニングします。
プロジェクト計画を策定し、エージェントを雇用またはカスタマイズし、プロンプトを調整して、エージェントに独自のモデルを与えます。
エージェントはハートビートに応じて自律的に動作します。ライブを視聴し、機密性の高いアクションを承認し、支出を制限し、いつでも方向を変更できます。
CEO とチャットしてください。残りはコーチがやってくれます。
1 つの会話 · いつでもワンクリックですぐ · 中断したところから再開 CEO あなたの連絡窓口 CEO はすべてのプロジェクト、チケット、名簿を確認します。状況を尋ねたり、役割を雇うように指示したりできます。返答はライブでストリーミングされ、重大な結果はすべて承認として返されます。
チケットが完了すると、コーチはそれをレビューし、永続的に学習されたルールをエージェントに書き込みます。同じ間違いは二度起こりません。手動でのプロンプト調整は必要ありません。
エージェントがあなたの秘密を握ることはありません。
キーとトークンは、メモリ内にのみ存在し、ディスク上には決して存在しないマスター キーの背後にあります。 Hezo はあなたなしではロックを解除できません。
すべてのエージェントはプロジェクトごとのコンテナ内で実行されます。ホストへのアクセスはなく、すべてのトラフィックはプロキシを経由します。悪いランの爆発半径は 1 ボックスです。
マシン、キー、支出、データの所有者はあなたです。 Git コミットは、プロジェクト キーを使用してホスト側で署名されます。
独自のプロバイダーを導入してください。自由に混ぜてください。
エージェントのチームが発送する必要があるすべてのもの。
出力プロキシでのシークレット置換 - プレースホルダが使用され、許可されたホストに対してのみ実際のキーが交換されます。
暗号化する

あなただけが保持する 1 つのマスター キーの背後で保存中 (AES-256-GCM) で保存されます。
プロジェクトごとの Docker 分離。すべてのエージェント トラフィックがプロキシを強制的に通過します。
git commit を確認し、プロジェクト キーを使用してホスト側で署名しました。
すべてのアクションとシークレットの使用に関する追加専用の監査証跡。
CEO、コーチ、キャプテン、従業員などの役割を調整する組織図。
タスクごとのルールとエージェントが管理する進捗状況の概要を備えたタスク ボード。
ハートビートの実行 : エージェントは予算に応じてスケジュールに従って起床し、作業を開始します。
複数のプロジェクト。それぞれが独自のコンテナ内の独立したチームです。
独自のプロバイダーを導入してください。エージェントごとに 1 つまで、モデルを自由に組み合わせることができます。
エージェントごと、プロジェクトごとに、日次、週次、月次のハード予算の上限があります。
エージェントは、ウィンドウが使い果たされると一時停止し、ウィンドウがロールオーバーすると再開します。
長期記憶 – CEO は、会話のたびにあなたの現在の好みを覚えています。
耐久性のあるプロジェクト文書 — PRD、仕様、調査結果が完全なバージョン履歴とともに保存されます。
作業はセッション間で蒸発するのではなく、実行をまたいできれいに保持されます。
参考資料の持ち込み — チームが作業するためのモックアップ、画像、PDF をアップロードします。
エージェントはテキストだけでなく、インタラクティブな HTML および SVG 成果物を作成します。
作成した作品は、どのデバイスでもアプリ内でプレビューできます。
モバイルファーストの Web アプリ — あらゆるデバイスから監視、チャット、承認できます。
MCP のインとアウト — 内蔵サーバーにより、どのクライアントでもチームを推進できるほか、エージェントにすでに使用しているツールを提供する外部 MCP サーバーも含まれます。
1 つの自己完結型バイナリ: Web アプリ、API、リアルタイム、データベース、ボールト。
タブではありません。他人のクラウドではありません。
稼働中
1つのコマンドで。
コピーcurl -fsSL https://hezo.ai/install.sh | sh 次に、「最初のプロジェクト」→「localhost:3100 を開く」を参照してください。セットアップ フローでは、マスター キーとモデルの接続について説明します。 hezo AI の仕事全体

セ。そしてあなたは上司です。

## Original Extract

Run a whole team of AI agents like a company — you set the goals, they ship the work. Org charts, projects, budgets, long-term memory, and approvals, all on hardware you own.

Hezo — A whole AI workforce. And you're the boss. hezo Docs GitHub Get started A whole AI workforce.
And you're the boss.
Hezo hires AI agents, runs them, and ships their work — without ever handing them your keys.
No cloud dependency Secrets never in agent context Signed commits Per-agent budget caps 01 How it works Three moves to a working team.
Describe the work to the CEO. It scopes the project and provisions a team — each in its own container.
Lay out the project plan, then hire or customize agents, tune their prompts, and give any agent its own model .
Agents work autonomously on a heartbeat . You watch live, approve sensitive actions, cap the spend, and change direction any time .
Chat with the CEO. The Coach does the rest.
one conversation · always one click away · picks up where you left off CEO Your point of contact The CEO sees every project, ticket, and roster. Ask how things are going or tell it to hire a role — replies stream back live , and anything consequential returns as an approval .
When a ticket completes, the Coach reviews it and writes durable learned rules back onto the agent. The same mistake doesn't happen twice — no prompt-tuning by hand.
Agents never hold your secrets.
Keys and tokens sit behind a master key that lives in memory only, never on disk. Hezo can't unlock itself without you.
Every agent runs in a per-project container — no host access, all traffic through the proxy. A bad run's blast radius is one box.
You own the machine, the keys, the spend, and the data. Git commits are signed host-side with your project key.
Bring your own providers. Mix freely.
Everything a team of agents needs to ship.
Secret substitution at the egress proxy — placeholders in, real keys swapped in only for allowed hosts.
Encrypted at rest (AES-256-GCM) behind one master key only you hold.
Per-project Docker isolation , with all agent traffic forced through the proxy.
Verified git commits , signed host-side with your project key.
An append-only audit trail of every action and secret use.
An org chart of roles — CEO, Coach, Captain, and workers — that coordinate.
A task board with per-task rules and an agent-maintained progress summary.
Heartbeat execution : agents wake on a schedule to pick up work, gated by budget.
Multiple projects , each an independent team in its own container.
Bring your own providers ; mix models freely, down to one per agent.
Hard budget caps — daily, weekly, monthly — per agent and per project.
Agents pause when a window is exhausted and resume when it rolls over.
Long-term memory — the CEO remembers your standing preferences across every conversation.
Durable project documents — PRDs, specs, and research, kept with full version history.
Work carries cleanly across runs instead of evaporating between sessions.
Bring references in — upload mockups, images, and PDFs for the team to work from.
Agents produce interactive HTML & SVG deliverables , not just text.
Preview their work in-app on any device, as it's built.
A mobile-first web app — oversee, chat, and approve from any device.
MCP in and out — a built-in server so any client can drive your teams, plus external MCP servers that give agents the tools you already use.
One self-contained binary : web app, API, realtime, database, and vault.
Not tabs. Not someone else's cloud.
Up and running
in one command.
Copy curl -fsSL https://hezo.ai/install.sh | sh Then see Your first project → Open localhost:3100 — the setup flow walks you through your master key and connecting a model. hezo A whole AI workforce. And you're the boss.
