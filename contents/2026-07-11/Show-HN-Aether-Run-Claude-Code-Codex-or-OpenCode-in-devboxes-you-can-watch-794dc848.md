---
source: "https://www.runaether.dev/"
hn_url: "https://news.ycombinator.com/item?id=48872847"
title: "Show HN: Aether – Run Claude Code, Codex, or OpenCode in devboxes you can watch"
article_title: "Aether: Not a black box. A devbox."
author: "pranav100000"
captured_at: "2026-07-11T15:45:10Z"
capture_tool: "hn-digest"
hn_id: 48872847
score: 1
comments: 0
posted_at: "2026-07-11T15:25:00Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Aether – Run Claude Code, Codex, or OpenCode in devboxes you can watch

- HN: [48872847](https://news.ycombinator.com/item?id=48872847)
- Source: [www.runaether.dev](https://www.runaether.dev/)
- Score: 1
- Comments: 0
- Posted: 2026-07-11T15:25:00Z

## Translation

タイトル: Show HN: Aether – 視聴できる開発ボックスでクロード コード、コーデックス、または OpenCode を実行する
記事のタイトル: エーテル: ブラック ボックスではありません。開発ボックス。
説明: クラウド開発ボックス内の自律コーディング エージェント。独自のサブスクリプションで Claude Code、Codex、または OpenCode を実行して、監視、操作、検証できます。
HN text: Claude Code や Codex のようなコーディング エージェントが登場して以来、私はかなり夢中になっています。推論で 20 倍の割引を受けられるのであれば、そうしないわけにはいきません。私も彼らにたくさんイライラしてきました。携帯電話のホットスポットに接続しているときにラップトップの蓋を開けたまま歩き回るのが嫌いでした。ルート ディレクトリを失わないようにコマンドを退屈に受け入れなければならず、異なるワークツリーで開発サーバーを同時に起動している複数のエージェントを実行すると RAM が不足してしまいます。上記の問題はどのクラウド エージェントにも当てはまらないため、多くの人が既存のクラウド エージェントがこの問題の解決策であると言うでしょうが、私が試したクラウド エージェント製品はすべてローカル エージェントよりも深刻な問題を抱えていました。タスクに応答するときに VM がセットアップされるのを待っていると、ファイアアンドフォーゲット ワークフローだけを目的としているように感じられ、エージェントとやり取りしたり、プラン モードを使用したりするのが恋しくなりました。その上、VM 全体がブラック ボックスであり、ブランチをチェックアウトして変更を自分でテストする以外にエージェントの動作を確認する方法がなかったので、すぐに面倒になってしまいました。これが私が両方の長所を生かした Aether を構築するきっかけとなりました。依存関係がインストールされ、開発サーバーが実行されている状態で、任意のサイズのリポジトリでエージェントを 5 秒以内に起動できます。応答すると、VM は 1 秒以内に準備が整うため、エージェントとのやり取りはローカル セッションのように感じられます。プラン モードとスキルは、マシン上で動作するのと同じように機能します。 VM もブラック ボックスではありません。すべてのワークスペースは、

ターミナル、ファイル エディタ、ポート プレビュー、Docker を備えた完全な開発ボックス。ローカルで行うのと同じように、いつでも介入して操作したり、ワークスペースの公開ポートを通じてエージェントの作業を手動でテストしたりできます。クラウドで実行すると、ローカル エージェントではまったく実行できない多くのことも可能になります。 Aether は、カスタム レビュー プロンプトを使用して PR をレビューし、PR にコメントを残すことができます。エージェントは PR コメントに応答するように設定することもでき、両方を有効にすると、ライター エージェントとレビューアー エージェントの間にループが作成され、エージェントは収束するまで動作し続けます (おそらく私の個人的なお気に入りの機能です)。ローカル エージェントで同じ結果を得るには、ターミナル タブ間でコピー アンド ペーストするという非常に退屈な数時間の作業が必要になります。このループ内を行き来するエージェントの PR の例を次に示します: https://github.com/Aether-Runtime/acme-checkout/pull/4 (これは公開デモ リポジトリであり、バグは意図的にシードされており、Readme で開示されています) Aether のその他の機能には以下が含まれます: - 視覚的な変更を行う PR には、変更のビデオ デモとスクリーンショットが付属します - Linear 問題または Slack チャネルからエージェントを起動し、コメント/スレッドでコミュニケーションします - Sentry との統合により、アラート発生時に 1 クリックでエージェントを起動できます。エージェントは Sentry MCP へのフル アクセスを取得するため、スタック トレースとコンテキスト自体を取得します。 - スケジュールされたジョブを実行するための自動化を設定します。サンドボックス プロバイダーをラップするのではなく、Firecracker レイヤーを自分たちで実行します。これにより、価格面で非常に競争力を発揮できます。推論は独自のサブスクリプションで実行されるため、巨額の補助金が得られ、Aether はコンピューティングに対してのみ料金を請求します。無料枠があり、有料プランは月額 30 ドルから始まります。肯定的でも否定的でも、フィードバックをお待ちしております。ご質問があれば喜んでお答えいたします。

記事本文:
エーテル: ブラックボックスではありません。開発ボックス。コンテンツにスキップ ループ 料金ドキュメント サインイン Claude Code、Codex、OpenCode 用の無料クラウド エージェントの構築を開始する
地元の監視。クラウドの超大国。クラウド開発ボックスで実行される自律型コーディング エージェントを監視、操作、検証できます。あなた自身のサブスクリプションで。 GitHub、Slack、Linear、端末、または API から。
動作を確認してください。独自のサブスクリプションを導入してください。あなたのエージェント、あなたの鍵。
デモ: 失敗したテストが再現され、エージェントによって修正され、テストが合格すると、レビューされたプル リクエストが開きます。
ライブ録画 · 未編集のタスクは 5 秒以内に開始 · アイドル時は 0 クレジット · 秒ごとの計測
自動化された PR レビュー レビュー エージェントは、到着した瞬間にすべての差分を読み取り、GitHub に実際のコメントを残します。修正をプッシュすると、承認されるまで再度レビューされます。コメントに応答すると、応答が返されます。
スケジュールされた自動化 プロンプトにスケジュールを設定します: 毎時、毎晩、または毎週月曜の 9 時。寝ている間にバックログをトリアージし、不安定なテストを毎朝追跡し、朝起きると ToDo リストの代わりに PR が表示されます。
Slack の統合 バグが報告されたスレッドで Aether について言及すると、その場で開発ボックスが起動します。進行状況の投稿を同じスレッドに戻します。答えてハンドルを切るか、「停止」と言ってプラグを抜きます。
線形統合 チームメイトのように Aether に問題を割り当てます。チケットは機能している間は「進行中」に移動し、PR リンクが自動的に戻り、エージェントのセッション メモが問題に記録されます。
Sentry の自動修正 新しい運用エラーは、スタック トレース、パンくずリスト、および問題のある行が含まれるアラートを読み終える前に修正タスクになります。修正は領収書とともに PR として届きます。
デプロイされたブレッドクラムを修正します · POST /api/webhooks/payments 500 ブレッドクラム · retryCheckout co_9f21 · cart=null 41 イベントで開始された自動修正タスク PR #1849 ✓ 実行に必要な場合はメールで返信してください

電話をかけると、質問が受信箱に届きます。電子メールに応答すると、エージェントが停止した場所から正確に応答します。ラップトップを開かずに夜間作業のブロックを解除します。
CLI brew install を実行し、次に aether run によってクラウド タスクがプロンプトにストリーミングされます。タスク、ファイル、git、シークレット、および請求はすべてコマンドであり、devbox 全体は 1 つのイーサ ssh でアクセスできます。
REST と WebSocket API このページのすべてはスクリプト可能です。プロジェクト、タスク、オートメーション用のプラットフォーム キーを備えた REST API に加え、ライブ ワークスペース、ターミナルから Git への WebSocket フィードも含まれます。
リポジトリのスキル SKILL.md をコミットすると、リポジトリにアクセスするすべてのエージェントは、テスト規則、レビュー バー、デプロイ ルールに従います。個人のスキルはプロジェクト全体にわたって役立ちます。
ライブ ポート プレビュー エージェントが開くすべてのポートは、独自の URL を取得します。実行中のアプリをクリックしてローカルホストなどでテストするか、ホットリロードを含めてテストするか、リンクをチームメイトに渡します。デフォルトでは非公開です。
PR でのビデオ証拠 エージェントは実際のディスプレイ上に実際のブラウザを備えています。構築したものをクリックして実行し、セッションを記録し、MP4 と前後のスクリーンショットを PR の説明に埋め込みます。
Claude Code、Codex、または OpenCode の任意のモデルを独自のサブスクリプションで提供し、40 以上のモデルをサポートします。タスクごとにモデルと推論労力を選択します。さらに必要な場合は、セッション中にホットスワップを行ってください。
すでに作業している場所でタスクを開始します。
CLI を brew install すると、aether run がクラウド タスクをプロンプトにストリーミングします。 Slack で Aether について言及し、それを Linear 問題に割り当てるか、プラットフォーム キーで API を押します。すべてのサーフェスが同じループに到達します。
見守って引き継げるループ。
ループの実行中は、ターミナル、ファイル エディター、ポート プレビュー、Docker などの開発ボックス全体が開いています。すべてのコマンドをライブで視聴します。走行中にジャンプしたり、同じ端末に入力したり、ハンドルを握ったり、ハンドルを完全に握ったりできます。盲目的な真実を決して求めない自律性

セント
計画モード: エージェントが提案し、あなたが承認してから実行します。
実行中にフォローアップをキューに追加すると、エージェントは再起動せずにフォローアップを組み込みます
いつでもエージェントを停止すると、開発ボックスは生きたままになり、使用できるようになります
デモ: エージェントが devbox ターミナルでコマンドをストリーミングしている間に、人間がクリックしてコマンドを入力します。エージェントが一時停止すると、チップが「あなたが運転しています」と表示され、その後コントロールハンドが戻ります。
初稿を決して出荷しない AI。
代理店より発送されます。エージェントによるレビュー。あなたによって承認されました。
Aether レビューをオンにすると、エージェントが開くすべての PR があなたに届く前に難題を実行します。レビュー エージェントが GitHub 内の差分を分解し、作成者が修正してプッシュし、コードが保持されるまでレビューが再度実行されます。リポジトリにあるものはすでにレビューに耐えており、それが機能することを証明するビデオが付いています。
ライブ録画、未編集: レビュー担当者が呼び出し元を通じて二重請求のバグを追跡し、作成者がそれを修正し、ループは 12 分で収束しました。人間は誰も触れなかった。
デモ: Aether がプル リクエストを開き、レビュー エージェントが二重請求のバグを見つけて変更をリクエストし、作成者エージェントが修正をプッシュし、レビューが承認し、チェックとデモ ビデオに合格してプル リクエストが解決されます。
ライブ録音・未編集を見て発送します。
視覚的な変更を行うすべての PR には、実際に実行されている変更の Playwright ビデオ、スクリーンショット、グリーン チェック、および完全なレビュー スレッドなどの領収書が付属します。
playwright · checkout-retry.mp4 0:42 ✓ 8 チェックが合格 3 スクリーンショット +18 -3 エーテル レビューが承認 冪等性ガードが重複配信テストに対して検証されました。承認します。 ✓
サブスクリプションが倍増しました。
Claude Code、Codex、または OpenCode の料金はすでに支払っています。ラップトップでは、1 つのエージェント、1 つのタスク、注意力、バッテリーを購入できます。同じサブスクリプションを Aether に向けると、フリートになります: 並列分離デ

vboxes, each ready in seconds, running while you sleep, wired into GitHub, Slack, Linear, and Sentry.モデルはあなたのものです。私たちはその周りのすべてです。
1 エージェント · マシンの修正: 不安定な認証テストの特技: CSV エクスポート 3:12 AM 修正: セントリー #4231 の修正: マージ競合の特技: ログのページ付けの修正: CI の失敗 インテリジェンスに二重に支払わないでください。
Black-box agents rent you a model at a markup and hide it inside an opaque meter. Aether charges only for compute: the intelligence runs on the subscription you already own.
モデル推論のバンドル、マークアップ
モデル推論 $0 (お客様自身のサブスクリプション)
Comparison uses list pricing for the leading black-box agent as of July 4, 2026.
超過なし: 上限に達するとアップグレードします
すべてのサイズ、最大 4 perf vCPU / 16 GB
クレジット パックは 1 クレジットあたり 0.10 ドル
一人のエンジニアの毎日のフリート。
すべてのサイズ、最大 4 perf vCPU / 16 GB
クレジット パックは 1 クレジットあたり 0.10 ドル
決して怠惰に過ごすことのない艦隊のために。
すべてのサイズ、最大 4 perf vCPU / 16 GB
クレジット パックは 1 クレジットあたり 0.10 ドル
中開発ボックスでは 1 クレジット = 6 分。 Small = 0.5×, Large = 2×, XL = 4×. Idle and suspended devboxes bill nothing.
Credit packs: 100 for $10, 250 for $25, 500 for $50, 1,000 for $100.パックの有効期限は購入後 12 か月です。
請求ポータルからいつでもキャンセルできます。
Encrypted at rest with AES-256-GCM under a key derived per user, and injected into your isolated VM only for the duration of a session. They're never written to logs, and you can revoke them at any time.
Each task runs in its own Firecracker microVM on a machine that is used once and destroyed afterward;セッション間では何も共有またはリサイクルされません。 The devbox is open the whole time, so you can watch every command your agent runs.
No. And because you bring your own keys, what your model provider does with your code is governed by yo

あなたのアカウントとその設定。私たちはそのループの中にはいません。
クレジットは、アクティブなランタイムのみを 1 秒あたりに計測します。 Medium devbox では 1 クレジットは 6 分です。 Small は 0.5 倍、Large は 2 倍、XL は 4 倍で実行されます。アイドル状態および一時停止された Devbox には料金はかかりません。また、ウォーム スロットはクレジットを消費しないプランの権利です。
どのエージェントとモデルを使用できますか?
Claude Code、Codex、OpenCode はそれぞれ独自のサブスクリプションまたはキーで実行されます。 OpenCode を通じて、Anthropic、OpenAI、Google、DeepSeek などの 40 以上のモデルにアクセスでき、タスクごとにモデルと推論の労力を設定できます。
はい。プラットフォーム キーを含む文書化された REST API はプロジェクト、タスク、自動化、使用法をカバーし、WebSocket API はライブ ワークスペース (ターミナル、ファイル、エージェント イベント、ポート、git) をストリーミングします。 CLI は同じ表面上に構築されます。
はい: 請求ポータルを介したセルフサービスで、請求サイクルの終了時に有効になります。無料プランは永続的なものであるため、そのままステップダウンすることもできます。
他に何か気になることがありますか?メールでお問い合わせください。
リポジトリを接続します。エージェントを連れてきてください。行く。
エージェントを虚空に発射するのはやめてください。
実際に表示される開発ボックスで最初のエージェントを実行します。無料。
ご自身のサブスクリプションをご持参ください。あなたのエージェント、あなたの鍵。
開発ボックス内の自律型コーディング エージェントを監視、操作、検証できます。

## Original Extract

Autonomous coding agents in a cloud devbox you can watch, steer, and verify, running Claude Code, Codex, or OpenCode on your own subscription.

Since coding agents like Claude Code and Codex came out, I've been pretty obsessed with them. It's hard not to when you're getting a 20x discount on inference. I've also been frustrated by them a lot. I hated walking around with my laptop lid open while connected to a hotspot on my phone, having to tediously accept commands so I don't lose my root directory, and running out of RAM when running multiple agents that were spinning up dev servers in different worktrees at once. A lot of people would say existing cloud agents are the solution to this, as none of the above problems apply to any cloud agent, but all the cloud agent products I tried had problems worse than local agents. Waiting for the VM to get set up when responding to a task made them feel like they were only for fire and forget workflows, and I missed going back and forth with the agent and using plan mode. On top of that, the whole VM was a black box and I had no way of verifying the agent's work besides checking out its branch and testing its changes myself, which got annoying super quickly. This is what led me to build Aether, which gives you the best of both worlds. You can start an agent on any repo of any size within 5 seconds, with dependencies installed and dev servers running. On responses the VM is ready in less than a second, so going back and forth with the agent feels like a local session. Plan mode and skills work like they do on your machine. The VM also isn't a black box. Every workspace is a full devbox with a terminal, file editor, port previews, and Docker. You can step in to steer at any point, or manually test the agent's work through the workspace's exposed ports, just like you would locally. Running in the cloud also enables a lot of stuff you can't do with a local agent at all. Aether can review PRs with a custom review prompt and leave comments on the PR. Agents can also be set to respond to PR comments, and when both are enabled, it creates a loop between the writer agent and reviewer agent, and the agents keep working until they converge (probably my personal favorite feature). Getting the same result with local agents would be an extremely tedious multi hour chore of copying and pasting between terminal tabs. Here's an example PR of the agents going back and forth in this loop: https://github.com/Aether-Runtime/acme-checkout/pull/4 (it's our public demo repo, the bugs are seeded on purpose, disclosed in the readme) Other features of Aether include: - PRs making visual changes come with video demos and screenshots of the changes - Start agents from Linear issues or Slack channels and communicate with them in the comments/thread - Set up the Sentry integration with 1 click to start an agent when an alert fires. The agent gets full access to the Sentry MCP, so it pulls the stack trace and context itself - Set up automations to run scheduled jobs We run the Firecracker layer ourselves instead of wrapping a sandbox provider, which enables us to be very competitive on pricing. The inference runs on your own subscription, so you still get the huge subsidization and Aether only charges for compute. There's a free tier and paid plans start at $30/mo. Really appreciate any feedback, positive or negative, and happy to answer any questions!

Aether: Not a black box. A devbox. Skip to content The loop Pricing Docs Sign in Start building free Cloud agents for Claude Code, Codex & OpenCode
Local oversight. Cloud superpowers. Autonomous coding agents that run in a cloud devbox you can watch, steer, and verify. On your own subscription. From GitHub, Slack, Linear, your terminal, or the API.
See it work Bring your own subscription. Your agent, your keys.
Demo: a failing test is reproduced, fixed by the agent, tests pass, and a reviewed pull request opens.
Recorded live · unedited tasks start in ~5s · 0 credits while idle · per-second metering
Automated PR review A review agent reads every diff the moment it lands and leaves real comments in GitHub. Push a fix and it reviews again, until it approves. Answer its comments and it answers back.
Scheduled automations Give a prompt a schedule: hourly, nightly, or every Monday at 9. Triage the backlog while you sleep, chase flaky tests every morning, and wake up to PRs instead of a to-do list.
Slack integration Mention Aether in the thread where the bug was reported and a devbox spins up on the spot. Progress posts back to the same thread; reply to steer it, or say stop to pull the plug.
Linear integration Assign an issue to Aether like a teammate. The ticket moves to In Progress while it works, the PR links back automatically, and the agent's session notes land in the issue.
Sentry auto-fix A new production error becomes a fix task before you finish reading the alert: stack trace, breadcrumbs, and offending line included. The fix arrives as a PR with the receipts.
fix deployed breadcrumb · POST /api/webhooks/payments 500 breadcrumb · retryCheckout co_9f21 · cart=null auto-fix task started at 41 events PR #1849 ✓ Reply by email When a run needs your call, the question lands in your inbox. Answer the email and the agent picks up exactly where it stopped. Unblock overnight work without opening a laptop.
CLI brew install , then aether run streams a cloud task into your prompt. Tasks, files, git, secrets, and billing are all commands, and the whole devbox is one aether ws ssh away.
REST & WebSocket API Everything on this page is scriptable: a REST API with platform keys for projects, tasks, and automations, plus a WebSocket feed of the live workspace, terminal to git.
Skills from your repo Commit a SKILL.md and every agent that touches the repo follows it: your test conventions, your review bar, your deploy rules. Personal skills follow you across projects.
Live port previews Every port the agent opens gets its own URL. Click into the running app and test it like localhost, hot reload included, or hand the link to a teammate. Private by default.
Video proof on PRs The agent has a real browser on a real display. It clicks through what it built, records the session, and embeds the MP4 and before/after screenshots in the PR description.
Bring any model Claude Code, Codex, or OpenCode, on your own subscription, with 40+ models behind them. Pick the model and reasoning effort per task; hot-swap mid-session when you need more.
Start tasks where you already work.
brew install the CLI and aether run streams a cloud task into your prompt. Mention Aether in Slack, assign it a Linear issue, or hit the API with a platform key: every surface reaches the same loop.
A loop you can watch and take over.
The whole devbox is open while the loop runs: terminal, file editor, port previews, Docker. Watch every command live. Jump in mid-run, type in the same terminal, steer, or take the wheel entirely. Autonomy that never asks for blind trust.
Plan mode: the agent proposes, you approve, then it executes
Queue follow-ups mid-run and the agent folds them in without restarting
Stop the agent anytime and the devbox stays alive, yours to use
Demo: while an agent streams commands in the devbox terminal, a human clicks in and types a command; the agent pauses, a chip flips to "you're driving," then control hands back.
AI that never ships its first draft.
Shipped by an agent. Reviewed by an agent. Approved by you.
Turn on Aether review and every PR the agent opens runs a gauntlet before it reaches you: a review agent tears the diff apart in GitHub, the author fixes and pushes, and the review runs again until the code holds up. What lands in your repo has already survived review, with a video to prove it works.
Recorded live, unedited: the reviewer traced a double-charge bug through the callers, the author fixed it, and the loop converged in 12 minutes. No human touched it.
Demo: Aether opens a pull request, its review agent finds a double-charge bug and requests changes, the author agent pushes a fix, the review approves, and the pull request resolves with passing checks and a demo video.
Recorded live · unedited Seeing is shipping.
Every PR making visual changes arrives with the receipts: a Playwright video of the change actually running, screenshots, green checks, and the full review thread.
playwright · checkout-retry.mp4 0:42 ✓ 8 checks passing 3 screenshots +18 −3 aether review approved Idempotency guard verified against the duplicate-delivery test. Approving. ✓
Your subscription, multiplied.
You already pay for Claude Code, Codex, or OpenCode. On your laptop that buys one agent, one task, your attention, your battery. Point the same subscription at Aether and it becomes a fleet: parallel isolated devboxes, each ready in seconds, running while you sleep, wired into GitHub, Slack, Linear, and Sentry. The model is yours. We're everything around it.
1 agent · your machine fix: flaky auth test feat: export CSV 3:12 AM fix: sentry #4231 fix: merge conflict feat: paginate logs fix: failing CI Don't pay for intelligence twice.
Black-box agents rent you a model at a markup and hide it inside an opaque meter. Aether charges only for compute: the intelligence runs on the subscription you already own.
model inference bundled, marked up
model inference $0 to us (your own subscription)
Comparison uses list pricing for the leading black-box agent as of July 4, 2026.
No overage: upgrade when you hit the cap
All sizes, up to 4 perf vCPUs / 16 GB
Credit packs at $0.10 per credit
The everyday fleet for one engineer.
All sizes, up to 4 perf vCPUs / 16 GB
Credit packs at $0.10 per credit
For a fleet that never sits idle.
All sizes, up to 4 perf vCPUs / 16 GB
Credit packs at $0.10 per credit
1 credit = 6 minutes on a Medium devbox. Small = 0.5×, Large = 2×, XL = 4×. Idle and suspended devboxes bill nothing.
Credit packs: 100 for $10, 250 for $25, 500 for $50, 1,000 for $100. Packs expire 12 months after purchase.
Cancel anytime via the billing portal.
Encrypted at rest with AES-256-GCM under a key derived per user, and injected into your isolated VM only for the duration of a session. They're never written to logs, and you can revoke them at any time.
Each task runs in its own Firecracker microVM on a machine that is used once and destroyed afterward; nothing is shared or recycled between sessions. The devbox is open the whole time, so you can watch every command your agent runs.
No. And because you bring your own keys, what your model provider does with your code is governed by your account and its settings; we're not in that loop.
Credits meter active runtime only, per second. One credit is six minutes on a Medium devbox; Small runs at 0.5×, Large at 2×, and XL at 4×. Idle and suspended devboxes bill nothing, and warm slots are a plan entitlement that never consumes credits.
Which agents and models can I use?
Claude Code, Codex, and OpenCode, each running on your own subscription or key. Through OpenCode you can reach 40+ models across Anthropic, OpenAI, Google, DeepSeek, and more, and you can set the model and reasoning effort per task.
Yes. A documented REST API with platform keys covers projects, tasks, automations, and usage, and a WebSocket API streams the live workspace: terminal, files, agent events, ports, and git. The CLI is built on the same surface.
Yes: self-serve via the billing portal, effective at the end of your billing cycle. The Free plan is permanent, so you can also just step down to it.
Something else on your mind? Email us.
Connect a repo. Bring your agent. Go.
Stop firing agents into the void.
Run your first agent in a devbox you can actually see. Free.
Bring your own subscription. Your agent, your keys.
Autonomous coding agents in a devbox you can watch, steer, and verify.
