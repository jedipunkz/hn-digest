---
source: "https://capacitor.kurrent.io/"
hn_url: "https://news.ycombinator.com/item?id=48734560"
title: "Show HN: Capacitor – shared mem for Claude Code, Cursor and other coding agents"
article_title: "Capacitor — shared memory for coding agents"
author: "lougarou"
captured_at: "2026-06-30T16:44:38Z"
capture_tool: "hn-digest"
hn_id: 48734560
score: 1
comments: 1
posted_at: "2026-06-30T15:56:30Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Capacitor – shared mem for Claude Code, Cursor and other coding agents

- HN: [48734560](https://news.ycombinator.com/item?id=48734560)
- Source: [capacitor.kurrent.io](https://capacitor.kurrent.io/)
- Score: 1
- Comments: 1
- Posted: 2026-06-30T15:56:30Z

## Translation

タイトル: Show HN: Capacitor – クロード コード、カーソル、その他のコーディング エージェントの共有メモリ
記事のタイトル: コンデンサ — コーディング エージェントの共有メモリ
説明: コンデンサーがチームを記録します。
HN テキスト: こんにちは、HN!私は Lokhesh です。小さなスタートアップ企業、Kurrent の製品および AI チームの一員です。私たちは Capacitor を構築しました。そして、それがどのように始まったかを正直に説明します。ある夜、当社のエンジニアリング リードである Alexey は、タスクを完了するために一晩中クロード コードを実行したままにしました。朝になると彼のマシンは OS を自動更新し、セッションは消えていました。文脈も、何が行われたかの記録も、何もありません。彼はクロード コードのドキュメントを詳しく調べ、セッションの履歴書とトランスクリプト ファイルについて学び、通常は目に見えない膨大な量のトランスクリプトが存在することに気づきました。エージェントがタスクをどのように推論したか。ファイルが変更されるたびに。すべてのツールが呼び出されます。失敗したすべてのコマンドと行き止まり。全部。金！そのとき彼は、「マシンが故障してすべてを失ったらどうなるだろう?」と考え、さらに「それをチームと共有できたらどうなるだろう?」と考えたのが Capacitor の誕生です。 Capacitor は、エージェント セッション (プロンプト、応答、ツール呼び出し、コード変更など) をイベント ストア (内部の KurrentDB) に記録します。各セッションは一意の ID を取得します。あらゆるエージェントやマシン上で再開したり、すべてを再説明することなくチームメイトに渡したりすることができます。これを使用する目的のいくつかを次に示します。 - 再起動時に存続する
- 移動する必要がある場合は、デスクトップからラップトップにセッションを引き継ぎます
- 以前にやったことを思い出してください (+ その理由と方法)
- セカンドオピニオンを求める（人間またはエージェントから）
- 同僚と知識を共有する
- エージェントが作成したコードの背後にある理由を説明する
- 病気になったときや休暇をとったときは引き継ぎます。この投稿をしているのは、HN コミュニティからコンデンサーに賭ける方法についてフィードバックをもらいたいからです。

えーっと。何がうまくいきますか?改善できる点は何でしょうか?何を追加すればよいでしょうか?ここから入手: https://capacitor.kurrent.io/api/auth/start?tier=free
ドキュメントをチェックしてください: https://capacitor.kurrent.io/docs/getting-started/quickstart... 注: セットアップを簡単にするために、Capacitor をクラウド サービスとして設計しましたが、最適なパッケージ化方法に関する提案を歓迎します。

記事本文:
">
コンデンサのドキュメント サインインについて 開始する プレビュー · 招待のみ チームのコーディング エージェント用の共有メモリ。
Capacitor は、エージェントが何を試みたのか、チームメイトが拒否したもの、最終的に何が通過したのか、そしてそれがなぜ重要だったのかなど、作業の背後にあるセッションを記録します。
KurrentDB の背後にあるチームによって構築されました。
チームのコーディング エージェントのためのセッション ストリーム。検索可能。共有可能。ベンダー中立。得点しました。
セッションを共有して、一緒に作業しましょう。
スラックカード → オープンスペック
ホスト型エージェント → コドライブ
要約 → 引き継ぎ Slack でセッション リンクを送信したり、ホストされているエージェントを共有したり、中断された作業を要約して別の開発者が作業を進められるようにします。
最初からやり直すことなく、コーディング エージェントを切り替えます。
コーデックス · 47 ターン · スタック
↓
クロード コード · 要約を読む コーデックスで壁にぶつかりましたか? Claude Code は要約から新しいセッションを開始します。失敗した仮説は除外されたままになります。
差分だけでなく、理由を付けて PR をレビューします。
- バッチサイズ: 500
+バッチサイズ: 5 レビュー エージェントはすでにコードを読み取っています。チームメイトがそれを書く前に試したことを読み取ることができるようになりました。
エージェントに同じリポジトリ ルールを 2 回教えるのはやめましょう。
事実判断→クラスター→ガイドライン
次の SessionStart でロードされます。 ルーブリックに対してセッションをスコアリングします。心に残ったレッスンは次走へのレポ指導になります。
任意のブラウザからコーディング エージェントを起動します。
ダッシュボード→起動
ラップトップ → 分離されたワークツリー
2 人のオペレーター → 1 つのセッション ターミナルを開かずに、分離されたワークツリーでクロード コードまたはコーデックスをスピンアップします。 2 人のチームメイトが同じエージェントを運転することができ、トランスクリプトには誰が何を入力したかが記録されます。
尋ねてみてください。以前にもこれに取り組んだことがありますか?
検索 → 1件ヒット、6ヶ月前
概要 → キューと直接呼び出し
トランスクリプト → 決定したターン エージェントは、表示できるすべてのセッションを検索し、チームメイトが前四半期に下した決定を見つけて、正確なターンをチャットに引き出します。
動作します

チームがすでに実行しているエージェント。
どのエージェントが実行したかに関係なく、すべてのセッションが自動的にキャプチャされます。ダッシュボードから起動されるホスト型エージェントは、現在、Claude Code と Codex で利用可能です。
チームが実行するすべてのエージェント セッションの完全で永続的な記録に基づいて構築されています。
数分で無料で開始できます。または、Enterprise についてご相談ください。Capacitor は、チームを Capacitor サーバーに導入します。
むしろ会話を始めませんか？チームに相談してください。私たちはすでにコーディング エージェントを使用しているチームと一緒に構築しています。
KurrentDB の背後にあるチームによって構築されました。本番環境でのイベント ストリームが私たちの仕事です。コーディングエージェントは単に新しい種類を生成するだけです。
capacitor.kurent.io の Cookie
製品分析には PostHog を使用し、Reddit 広告ピクセルを使用して
当社の広告および Reo.dev からの訪問とコンバージョンを測定し、
どの企業が訪問しているかを把握します。 3 つすべては、実行後にのみロードされます
受け入れると、Reddit ピクセルと Reo が追跡 Cookie を設定します。私たち
データを売らないでください。いつでも考えを変えることができます
Cookie ポリシーのページ 。

## Original Extract

Capacitor records your team

Hi HN! I’m Lokhesh, I am part of the Product and AI team at a small startup, Kurrent. We built Capacitor, and here’s honestly how it started: One night our Engineering Lead, Alexey, left Claude Code running overnight to finish a task. In the morning his machine had auto-updated the OS and his session was gone. No context, no record of what had been done, nothing. He dug into the Claude Code docs, learned about session resume and transcript files, and realized there’s an enormous amount in those transcripts that’s normally invisible. How the agent reasoned through the task. Every file change. Every tool call. Every failed command and dead end. All of it. Gold! That’s when he thought, ‘what if my machine dies and I lose all of that?’ And then, ‘what if I could share it with my team?’ That’s how Capacitor came to be. Capacitor records agent sessions (prompts, responses, tool calls, code changes, etc.) into an event store (KurrentDB under the hood). Each session gets a unique ID. You can resume it in any agent, on any machine, or hand it to a teammate without re-explaining everything. Here are some of the things we use it for: - Survive on restart
- Hand a session over from a desktop to a laptop when you need to travel
- Remember what I’ve done before (+ why and how)
- Ask a second opinion (from a human or agent)
- Share knowledge with peers
- Explain the reasoning behind the code written by an agent
- Hand over when I get sick or take a day off I’m doing this post because I’d love to get feedback from the HN community on how we can make Capacitor better. What works well? What could be improved? What should we add? Get it here: https://capacitor.kurrent.io/api/auth/start?tier=free
Check out the docs: https://capacitor.kurrent.io/docs/getting-started/quickstart... Note: We designed Capacitor as a cloud service to make setup easy, but are open to suggestions on how best to package it.

">
C apacitor Docs About Sign in Get started Preview · invite-only Shared memory for your team’s coding agents.
Capacitor records the session behind the work: what agents tried, what teammates rejected, what finally passed, and why it mattered.
Built by the team behind KurrentDB .
A session stream for your team’s coding agents. Searchable. Shareable. Vendor-neutral. Scored.
Share the session, then work together.
Slack card → open spec
hosted agent → co-drive
recap → hand off Send a session link in Slack, share a hosted agent, or recap interrupted work so another dev can move it forward.
Switch coding agents without starting over.
codex · 47 turns · stuck
↓
claude code · reads recap Hit a wall in Codex? Claude Code starts a new session with the recap. The failed hypotheses stay ruled out.
Review PRs with the reasoning, not just the diff.
- batchSize: 500
+ batchSize: 5 Your review agent already reads the code. Now it can read what your teammate tried before they wrote it.
Stop teaching agents the same repo rules twice.
judge fact → cluster → guideline
loaded at next SessionStart Score sessions against your rubric. The lessons that stick become repo guidance for the next run.
Launch a coding agent from any browser.
dashboard → launch
your laptop → isolated worktree
two operators → one session Spin up Claude Code or Codex in an isolated worktree without opening a terminal. Two teammates can drive the same agent and the transcript records who typed what.
Ask: have we worked on this before?
search → 1 hit, 6 months ago
summary → queue vs direct call
transcript → the turn where we decided Your agent searches every session you can see, finds the decision a teammate made last quarter, and pulls the exact turn into the chat.
Works with the agents your team already runs.
Every session is captured automatically, whichever agent ran it. Hosted agents — launched from the dashboard — are available for Claude Code and Codex today.
Built on a complete, durable record of every agent session your team runs.
Start free in minutes, or talk to us about Enterprise — Capacitor gets your team onto a Capacitor server.
Rather start a conversation? Talk to the team — we’re building with teams that already use coding agents.
Built by the team behind KurrentDB — event streams in production are what we do. Coding agents just produce a new kind.
Cookies on capacitor.kurrent.io
We use PostHog for product analytics, the Reddit Ads pixel to
measure visits and conversions from our advertising, and Reo.dev to
understand which companies visit. All three load only after you
accept, and the Reddit pixel and Reo set tracking cookies. We
don’t sell your data. You can change your mind any time on the
cookie policy page .
