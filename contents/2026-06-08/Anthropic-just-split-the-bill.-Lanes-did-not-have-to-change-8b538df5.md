---
source: "https://lanes.sh/blog/claude-billing-split"
hn_url: "https://news.ycombinator.com/item?id=48446272"
title: "Anthropic just split the bill. Lanes did not have to change."
article_title: "Anthropic just split the bill. Lanes did not have to change. | Blog | Lanes"
author: "MayCXC"
captured_at: "2026-06-08T16:29:50Z"
capture_tool: "hn-digest"
hn_id: 48446272
score: 1
comments: 1
posted_at: "2026-06-08T15:00:51Z"
tags:
  - hacker-news
  - translated
---

# Anthropic just split the bill. Lanes did not have to change.

- HN: [48446272](https://news.ycombinator.com/item?id=48446272)
- Source: [lanes.sh](https://lanes.sh/blog/claude-billing-split)
- Score: 1
- Comments: 1
- Posted: 2026-06-08T15:00:51Z

## Translation

タイトル: Anthropic は割り勘です。車線を変更する必要はありませんでした。
記事のタイトル: Anthropic はちょうど割り勘です。車線を変更する必要はありませんでした。 |ブログ |レーン
説明: 6 月 15 日、クロードのサブスクリプションはプログラマティック エージェントの使用をカバーしなくなります。 Lanes は実際のターミナルで公式の claude CLI を実行するため、Max プランは従来どおり機能し続けます。

記事本文:
Anthropic はちょうど割り勘です。車線を変更する必要はありませんでした。 |ブログ | Lanes Lanes Docs Get Started Blog / Industry Anthropic はちょうど割り勘です。車線を変更する必要はありませんでした。
Anthropic はエージェントの使用料金を再設定しました。 Zed、Conductor、T3 Code、superset.sh、または Agent SDK 上に構築されたその他のものを通じて Claude を実行する場合、Max 5x プランでカバーできる範囲は先週よりも大幅に減ります。レーンズは変更されません。レーンズは、価格が変更されたばかりの部品に基づいて構築されたものではないためです。ここでは、何が変更され、誰が影響を受けるのか、そしてなぜレーンズのセッションがいつもと同じサブスクリプションプールから引き続けられるのかを説明します。
5 月 14 日、Anthropic はプラン加入者に、6 月 15 日発効の新しい月次エージェント SDK クレジットについて電子メールを送信しました。完全なポリシーは、新しいエージェント SDK クレジットに関する Anthropic のサポート記事に記載されています。
2026 年 6 月 15 日以降、Claude Agent SDK と claude -p の使用量は、Claude プランの使用制限にカウントされなくなります。サブスクリプションの使用制限は変更されず、Claude Code、Claude Cowork、および Claude のインタラクティブな使用のために予約されたままになります。
Agent SDK の使用量をカバーするために、対象となるすべてのプランに個別の月次クレジットが付与されるようになりました。
チームスタンダードシート: $20、チームプレミアムシート: $100
エンタープライズ (使用量ベース): 20 ドル、エンタープライズ (シートベースのプレミアム): 200 ドル
クレジットはユーザーごとで、毎月更新され、繰り越されず、最初に消費されます。それがなくなると、追加のエージェント SDK の使用量は標準 API レートでの追加使用量にルーティングされますが、これは明示的に追加使用量を有効にしている場合に限ります。そうしないと、SDK リクエストは次のサイクルまで停止します。
Anthropic はクレジットの意図についても明確にしています。
Agent SDK の月間クレジットは、個々の実験と自動化に合わせてサイズ設定されています。共有プロダクション自動化を実行しているチームは、予測可能な従量課金制を実現する API キーを備えた Claude Developer Platform を使用する必要があります。
あなたがrなら

SDK を駆動するツールを通じて複数のエージェントを並行して実行する、その段落はあなたについて話しています。クレジットのサイズがそれに見合うものではありませんでした。
境界線は、ワークフローが SDK 境界のどちら側に位置するかです。 Anthropic は、対象となる表面として、独自のプロジェクトの Agent SDK、非対話モードの claude -p、Claude Code GitHub Actions の統合、および「Agent SDK を介して Claude サブスクリプションで認証するサードパーティ アプリ」を挙げています。既存のサブスクリプションに残るものは次のとおりです。
ターミナルまたは IDE でのインタラクティブなクロード コード
Web、デスクトップ、またはモバイル アプリでのクロードの会話
2 番目のリストは、ワークフローが請求書のどちら側に着地するかを決定するリストです。
変化の規模を知るために、変更に関する Zed 自身の投稿では、以前はサブスクリプションが API 料金と比較して約 15 ～ 30 倍のエージェント使用料を補助していたことが記載されています。新しいクレジットはフル API レートで請求されます。
ワークフローがプログラムでクロードを駆動する場合、あなたは新しいプールに入ります。
Zed は Agent Client Protocol を通じて Claude を実行しており、ACP を介した使用はクレジットから引き出されることを自身の投稿で確認しました。 Theo Browne のオープンソース コントロール プレーンである T3 コードは、Claude に同じ種類の SDK サーフェスを介して実行させます。Theo は、有効なサブスクリプション価値が約 40 分の 1 に減少することについて公に声を上げています。 Conductor と superset.sh は、こ​​の議論で Zed と並んで最も一般的に名前が挙げられるパラレル エージェント オーケストレーターであり、どちらも開発者コミュニティによって影響を受けると報告されています。 Claude Code GitHub Actions、OpenClaw、および @anthropic-ai/claude-agent-sdk 上に構築されたその他のものはすべて同じボートにあります。
Zed はすでに回避策を自分の言葉で公開しています。「既存のサブスクリプション制限で Claude Pro または Max プランで Claude を使い続けたい場合は、次のように実行できます。

Anthropic の公式クロード CLI は、ACP 経由ではなく、Zed 内のターミナルで実行されます。公式のクロード CLI がターミナルで実行されると、新しいクレジットではなく、サブスクリプションの制限が使用されます。」
この回避策は Lanes では回避策ではありません。これはデフォルトのパスです。それが唯一の道です。
レーンズが新しいプールに参加しない理由
Lanes は Agent SDK をラップしません。 Lanes は claude -p を呼び出しません。 Lanes は ACP を実装していません。問題に関するセッションを開始すると、Lanes は実際の疑似端末 (ネイティブ PTY ライブラリを使用) を開き、実際のドットファイルで実際のログイン シェルを起動し、自分で入力した場合とまったく同じように公式の claude コマンドを実行します。
同じバイナリ。同じフラグです。同じ認証です。同じサブスクリプション プール。
Anthropic の課金の観点から見ると、Lanes セッションは、端末に座って claude --resume <session-id> を実行している開発者と区別できません。それがそういうものだからです。 Anthropic 自身のサポート記事 (「ターミナルまたは IDE のインタラクティブ クロード コード」) の正確なフレーズは、すべての Lanes セッションが実行されるサーフェスを説明しています。
現在のツールを使用するために必要なものはすべて、動かないアーキテクチャにマッピングされています。
並行したオーケストレーションと監視。各エージェントは、計画、実装、レビュー、完了をドラッグしてボード上に独自の問題カードを取得します。各カードには、ビジー、入力待ち、停止、エラーなどのリアルタイム ステータスを示すライブ ターミナルがあります。エージェントがあなたを必要とするとき、ベルが鳴ります。同じ並列エージェント サーフェス Conductor と superset.sh が販売されていますが、その下のランタイムがネイティブ CLI である点が異なります。
セッションごとにワークツリーを分離。 Lanes は、課題ごとに生成されたブランチ名を持つ git ワークツリーを自動作成し、コミットされていない状態とマージされていない状態をリアルタイムで追跡し、課題が完了したときにクリーンアップします。衝突も、踏みつけも、手でジャグリングすることもありません

アークツリーの小道。
オートメーション。リニアと GitHub の統合により、チケットがボードに直接プルされ、PR リンクが書き戻されます。クイック コマンドは、プリセット プロンプトを任意のセッションに挿入します。
すべてはインタラクティブなプールの上にあります。 Max プランは先週と同じことを継続します。エージェント群は稼働し続けます。プログラマティック側にメーター制はなく、100 ドルの上限もなく、スプリント途中で予期せぬバーンダウンもありません。
ソフトウェアは価格決定段階にあります。 Anthropic は分裂を明確にした最初の企業であり、彼らが最後ではありません。ファーストパーティ CLI を介してルーティングするアーキテクチャは、これらの変化を乗り越えて動作し続けます。 SDK をラップするアーキテクチャは、ポリシーが変更されるたびに再価格設定税を支払い続けます。
レーンズは意図的に最初の方法で建設されました。
高価になったばかりのツールを使用して複数のエージェントを実行している場合は、ここから始めるか、Discord に参加してください。 Max プランをご持参ください。
2026 年 6 月 4 日 v0.42 の新機能: データベースの内部を見てみる
v0.42 では、データベース エクスプローラーがコンテキスト サイドパネルに追加されます。作業フォルダー内の SQLite データベースの自動検出、テーブルとビューの参照、組み込みエディターからの読み取り専用クエリの実行をすべて、レーンを離れることなく実行できます。さらに、重複問題アクションと一連の安定性修正も含まれます。
2026 年 5 月 21 日 v0.41 の新機能: 多数のセッション、1 つの問題
同じ問題に対して複数のエージェント セッションを並行して実行します。実行する計画を連鎖させて、レビューしたり、アプローチを並べて比較したり、単に作業を整理したりすることができます。クイック アクションはセッション セレクターから起動できるようになり、セッション設定では CLI、モデル、作業量を事前に選択できるようになりました。さらに、入力待ち、トークン取り消し、プロジェクト切り替え全体の安定性も修正されました。

## Original Extract

On June 15, Claude subscriptions stop covering programmatic agent usage. Lanes runs the official claude CLI in a real terminal, so your Max plan keeps working as it did.

Anthropic just split the bill. Lanes did not have to change. | Blog | Lanes Lanes Docs Get Started Blog / Industry Anthropic just split the bill. Lanes did not have to change.
Anthropic just repriced agent usage. If you run Claude through Zed, Conductor, T3 Code, superset.sh, or anything else built on the Agent SDK, your Max 5x plan now covers a lot less than it did last week. Lanes does not change, because Lanes was never built on the part that just got repriced. Here is what changed, who it hits, and why your Lanes sessions keep drawing from the same subscription pool they always did.
On May 14, Anthropic emailed plan subscribers about a new monthly Agent SDK credit, effective June 15. The full policy is in Anthropic's support article on the new Agent SDK credit :
Starting June 15, 2026 , Claude Agent SDK and claude -p usage no longer counts toward your Claude plan's usage limits. Your subscription usage limits stay the same and stay reserved for interactive use of Claude Code, Claude Cowork, and Claude.
To cover Agent SDK usage, every eligible plan now gets a separate monthly credit:
Team Standard seats: $20, Team Premium seats: $100
Enterprise (usage-based): $20, Enterprise (seat-based Premium): $200
The credit is per-user, refreshes monthly, does not roll over, and drains first. When it runs out, additional Agent SDK usage routes to extra usage at standard API rates, but only if you have explicitly enabled extra usage. Otherwise your SDK requests stop until next cycle.
Anthropic is also clear about the credit's intent:
The Agent SDK monthly credit is sized for individual experimentation and automation. Teams running shared production automation should use the Claude Developer Platform with an API key for predictable pay-as-you-go billing.
If you are running multiple agents in parallel through a tool that drives the SDK, that paragraph is talking about you. The credit was not sized for it.
The dividing line is which side of the SDK boundary your workflow sits on. Anthropic lists the covered surface as: the Agent SDK in your own projects, claude -p in non-interactive mode, the Claude Code GitHub Actions integration, and "third-party apps that authenticate with your Claude subscription through the Agent SDK." What stays on your existing subscription is:
Interactive Claude Code in the terminal or IDE
Claude conversations on the web, desktop, or mobile apps
That second list is the one that decides which side of the bill your workflow lands on.
For a sense of the scale of the shift, Zed's own post on the change notes that subscriptions previously subsidised agent usage at roughly 15 to 30x compared to API pricing. The new credit is billed at full API rates.
If your workflow drives Claude programmatically, you are in the new pool.
Zed runs Claude through the Agent Client Protocol, and has confirmed in their own post that any usage going through ACP draws from the credit. T3 Code, Theo Browne's open-source control plane, drives Claude through the same kind of SDK surface, and Theo has been publicly vocal about a roughly 40x reduction in effective subscription value. Conductor and superset.sh are the parallel-agent orchestrators most commonly named alongside Zed in this discussion, both flagged by the developer community as affected. Claude Code GitHub Actions, OpenClaw, and anything else built on @anthropic-ai/claude-agent-sdk are in the same boat.
Zed has already published the workaround in their own words: "If you want to keep using Claude with your Claude Pro or Max plan at your existing subscription limits, you can run Anthropic's official claude CLI in a terminal inside Zed instead of through ACP. When the official claude CLI runs in the terminal, it uses your subscription's limits, not the new credit."
That workaround is not a workaround in Lanes. It is the default path. It is the only path.
Why Lanes is not in the new pool
Lanes does not wrap the Agent SDK. Lanes does not call claude -p . Lanes does not implement ACP. When you start a session on an issue, Lanes opens a real pseudo-terminal (using a native PTY library), boots your real login shell with your real dotfiles, and runs the official claude command exactly as you would if you typed it yourself.
Same binary. Same flags. Same authentication. Same subscription pool.
From Anthropic's billing perspective, a Lanes session is indistinguishable from a developer sitting at a terminal and running claude --resume <session-id> . Because that is what it is. The exact phrase from Anthropic's own support article ("Interactive Claude Code in the terminal or IDE") describes the surface every Lanes session runs on.
Everything you came to your current tool for, mapped onto an architecture that did not move.
Parallel orchestration and oversight. Each agent gets its own issue card on a board you drag through Planning, Implementation, Review, and Done. Each card has a live terminal with real-time status: busy, awaiting input, stopped, error. The bell rings when an agent needs you. The same parallel-agent surface Conductor and superset.sh sell, except the runtime underneath is the native CLI.
Isolated worktrees per session. Lanes auto-creates a git worktree with a generated branch name for each issue, tracks uncommitted and unmerged state in real time, and cleans up when the issue lands in Done. No collisions, no stomping, no hand-juggling of worktree paths.
Automation. Linear and GitHub integrations pull tickets straight onto the board and write PR links back. Quick commands inject preset prompts into any session.
All of it sits on top of the interactive pool. Your Max plan keeps doing what it did last week. Your fleet of agents keeps running. No metering on the programmatic side, no $100 ceiling, no surprise burndown halfway through a sprint.
Software is in a pricing-discovery phase. Anthropic is the first to make the split explicit, and they will not be the last. Architectures that route through the first-party CLI keep working through these shifts. Architectures that wrap the SDK keep paying the re-pricing tax every time the policy moves.
Lanes was built the first way on purpose.
If you have been running multiple agents through a tool that just got more expensive, start here or join our Discord . Bring your Max plan with you.
Jun 4, 2026 What's New in v0.42: Look Inside Your Databases
v0.42 adds a Databases Explorer to the context sidepanel: auto-discover SQLite databases in your working folder, browse tables and views, and run read-only queries from a built-in editor, all without leaving Lanes. Plus a Duplicate Issue action and a run of stability fixes.
May 21, 2026 What's New in v0.41: Many Sessions, One Issue
Run multiple agent sessions on the same issue in parallel. Chain plan to implement to review, compare approaches side by side, or just keep work organised. Quick Actions now launch from the session selector, and Session Settings lets you pick CLI, model, and effort upfront. Plus stability fixes across awaiting-input, token revocation, and project switching.
