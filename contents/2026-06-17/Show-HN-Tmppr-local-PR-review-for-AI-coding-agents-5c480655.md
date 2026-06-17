---
source: "https://tmppr.com/"
hn_url: "https://news.ycombinator.com/item?id=48574002"
title: "Show HN: Tmppr – local PR review for AI coding agents"
article_title: "tmppr — Autonomous PR coordination for AI agents and humans"
author: "douglaswlance"
captured_at: "2026-06-17T18:57:17Z"
capture_tool: "hn-digest"
hn_id: 48574002
score: 2
comments: 0
posted_at: "2026-06-17T17:50:30Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Tmppr – local PR review for AI coding agents

- HN: [48574002](https://news.ycombinator.com/item?id=48574002)
- Source: [tmppr.com](https://tmppr.com/)
- Score: 2
- Comments: 0
- Posted: 2026-06-17T17:50:30Z

## Translation

タイトル: Show HN: Tmppr – AI コーディング エージェントのローカル PR レビュー
記事のタイトル: tmppr — AI エージェントと人間のための自律的な PR 調整
説明: エージェントが作業を開いたり、CI を実行したり、相互にレビューしたり、フォローアップ エージェントをトリガーしたり、強制されたゲートを介してマージしたりできるローカル GitHub スタイルのプル リクエスト。

記事本文:
› tmppr / ローカル PR レビュー 機能 プラグイン 仕組み ワークフロー ホスティング FAQ tmppr のダウンロード コーディング エージェントがプル リクエストを通じて連携できるようにします。
すべてのエージェントに共有ローカル フォージを提供します。エージェントは、PR を開いたり、CI を実行したり、相互にレビューしたり、問題をファイルしたり、フォローアップ エージェントをトリガーしたり、設定したルールに従ってマージしたりできます。
仕組みを確認してください すでに使用しているスタックとうまく連携します
GitHub アクション クロード コード カーソル コーデックス (OpenAI) プラグイン SDK · 新規
コピーするのではなく、拡張するように構築されています。
@tmppr/plugin-sdk を使用すると、フォークせずに tmppr を拡張できます。単一のプラグインでエージェント プロバイダー、自動化トリガー、マージ ルール、フック エグゼキューター、イベント サブスクライバー、HTTP ルート、構成を追加できます。これらはすべて起動時に読み込まれ、コアへのパッチは必要ありません。
コードを漏らすことなく、自律エージェントの調整層を実現します。
tmppr は、PR、問題、CI、レビュー スレッド、マージ ルール、イベント トリガーの自動化など、人間のチームが期待するのと同じ手術室をエージェントに提供します。チャット ログに消えることなく、独立して作業できます。
GitHub スタイルの差分ビューア 構文強調表示を備えた並べて統合された差分。インライン コードのコメント、スレッド化された返信、ファイルの表示状態、およびリビジョンごとの相互差分により、前回以降に変更された部分のみを再レビューできます。実際の CI、ローカルで実行 既存の GitHub Actions ワークフロー YAML をマシン上で実行します。ライブ ログ ストリーミング、ANSI カラー、構造化されたジョブとステップ - マージの準備に直接フィードします。分単位の共有やより大きなボックスが必要な場合は、tmppr クラウド ランナーにバーストします。強制マージ ゲート tmppr は、レビューの承認、CI ステータス、競合チェックから GO / NO-GO をレポートします。切り離されたワークツリーでの実際の git merge --no-ff — 驚くべきことではありません。エージェント用の最上級の CLI すべてのアクション (開く、同期、コメント、返信、レビュー、CI の再実行、マージ) には、スクリプト可能な CLI があります。

tmppr をクロード コード、カーソル、またはコーデックス ループにドロップすると、エージェントが自分の作業をレビューできるようになります。交換可能なエージェント プロバイダー tmppr を Claude Code 、 Codex 、 opencode 、 aider 、またはカスタム コマンドに指定します。マシンまたはリポジトリごとのデフォルトを設定します。あなたのエージェント、あなたのモデル。フルライフサイクルエージェントは、準備済みの問題の実装、新しい PR のレビュー、各リビジョンの QA、失敗した CI の修正、要求された変更のフォローアップ、およびマージ対応 PR の出荷を行うイベント駆動型エージェントをトリガーします。ループ全体が自動的に実行されます。 1 つのコマンドでオンボード tmppr init を実行すると、任意の git リポジトリを tmppr リモートに変換し、エージェント トリガー バンドルをインストールし、デフォルトのプロバイダを設定し、MCP サーバーを登録します。これを 1 つのコマンドで実行します。 MCP サーバーの組み込み すべての tmppr アクションは MCP ツールとして Claude Code および Codex に公開され、init で自動登録されます。エージェントは PR、CI、レビュー、マージをネイティブ ツールとして呼び出します。接着コードは必要ありません。イベントの自動化 リポジトリ イベント ( pr.created 、 pr.updated 、 ci.updated 、 review.updated 、 issue.updated 、および hook.completed ) に対して任意のターミナル コマンドを実行します。スタック型 PR ワークフロー 依存するブランチをスタックにチェーンし、全体を操作します。積み重ねられた PR を手作業で解くことなく、リベース、同期、および出荷を順番に行います。コードがボックスから離れることはありません 127.0.0 にバインドされます
[切り捨てられた]
発行から統合まで - エージェントが主導し、あなたが承認します。
GitHub と同じプルリクエスト ループが、完全にマシン上で実行されます。実際のアプリでスクロールして確認してください。
開いた作品がボード上に置かれます。エージェントは準備ができている問題をピックアップします。何が飛んでいるのかが正確にわかります。
すべてのブランチがプル リクエスト、つまり完全な会話、タイムライン、レビュー スレッドになります。
インライン コメントとスレッド形式の返信を含む差分を並べて表示でき、すべてお使いのマシン上で実行できます。
GitHub アクションは、ライブ グリーン チェックを使用してローカルで実行されます。強制ゲートを通って合流します。
リポジトリは独自に実行されるため、

ソフトウェアのライフサイクル。
1 つの .tmppr/workflow.yaml により、リポジトリがイベント駆動型のアセンブリ ラインに変わります。各ライフサイクル イベント (発行の準備が整う、PR が開始される、CI が赤になる) ごとに、役割を起動してそのジョブを実行し、次のイベントを発行します。チェーン全体はマージされた PR まで実行され、ループには人が関与せず、マージ ルールによってのみ制御されます。
1 automation.tick のディスパッチャー 準備が整ったブロックされていない問題を、制限されたバッチでパイプラインにプロモートします。
2 issue.updated の実装者 分離されたワークツリーでドラフト PR を開き、小さなコミットでコードを書き込みます。
3 pr.created のレビュー担当者 差分を読み、承認または変更リクエストを送信します。
4 pr.updated の QA ブランチを構築し、新しいリビジョンごとに完全なテスト スイートを実行します。
5 Fixer on ci.updated 赤の CI を診断し、修正をブランチにプッシュします。
6 レビューのフォローアップ。更新 要求された変更と別のパスの再プッシュに対処しました。
7 ci.updated の荷送人 チェックが緑色になり、承認がマージ ゲートを満たしたらマージします。
# .tmppr/workflow.yaml
ID: tmppr-swdlc
フック:
- ID: 実装準備完了問題
イベント: 問題.更新
アクション:
- タイプ: start_agent
役割: 実装者
プロンプト: プロンプト/implement-ready-issue.md
- ID: ship-ready-pr
イベント: ci.更新
アクション:
- タイプ: start_agent
役割: 出荷者 分離されたワークツリー 実行ごとに独自の git ワークツリーが取得されるため、エージェントは相互にステップを踏むことなく並行して動作します。交換可能なエージェント 各ロールは、選択したプロバイダー (Claude Code、Codex、opencode、aider、または独自のコマンド) を実行します。強制的なマージ ゲート 必要な承認とステータス チェックに合格するまではマージは行われません。エージェントはマージを獲得し、バイパスしません。比較した
GitHub PRs Forgejo Gitea ソースはローカルに留まります はい — 127.0.0.1 いいえ セルフホスト セルフホスト リポジトリ機能のパリティ (コード、PR、問題、wiki、パッケージ、リリース、アクション) はい はい はい Gitea

/Forgejo 互換 API ドロップイン いいえ ネイティブ ネイティブ Gitea および Forgejo からのインポート / ミラー はい インポートのみ はい はい GitHub スタイルの差分 UI はい はい はい スレッド化されたインライン コメント + レビュー はい はい はい CI はハードウェア上でローカルに実行されます 組み込みのクラウド時間分 セルフホスト ランナー セルフホスト ランナー Git LFS はい はい はい ブランチとタグの保護 はい はい はい はい CODEOWNERS はい はい はい Webhook および git フック はいWebhook のみ ○ ○ プロジェクト / ボード ○ ○ ○ ○ パッケージ レジストリ (マルチエコシステム) ○ ○ ○ ○ スタック PR ワークフロー ○ 手動 × × 強制マージ ゲート (GO / NO-GO) ○ 必須チェック 必須チェック 必須チェック ActivityPub / ForgeFed フェデレーション ○ × 実験的 × AI エージェント用に構築 ネイティブ × × × エージェント用のスクリプト可能な CLI ファーストクラス gh CLI tea CLI tea CLI 交換可能なエージェント モデル Claude Code、 Codex、opencode、aider、カスタム No No No フルライフサイクル エージェント トリガー 組み込み DIY アクション DIY アクション DIY アクション MCP サーバー内蔵 Yes No No No プラグイン SDK 経由で拡張可能 フォーク不要 アプリ + アクション フォーク / webhook フォーク / webhook AI エージェント価格 ローカルシートなし シートベース 無料 / OSS 無料 / OSS 有料サーフェス クラウド ランナー + ホストされたリポジトリ クラウド アカウント シート なし — OSS Gitea Cloud メーカーより
私はエージェントに毎日コードを書いてもらっています。出荷されたものをレビューすることが、今では実際の仕事になっています。そして、それをチャット ウィンドウ内で行うと、差分スレッド、インライン コメント、CI ゲートがないため、見逃してはいけないものを見逃してしまいます。
レビュー UI を取得するためだけにすべての実験的ブランチを GitHub にプッシュすると時間がかかり、コードが漏洩し、AI エージェントに対するシートごとの税金がすぐに加算されます。私は、エージェントが操作できる CLI を備えた、ローカルでの GitHub レビュー サーフェスと、何かがマージされる前の実際の G​​O/NO-GO ゲートが必要でした。
tmpprがそれです。ローカルレビューがデフォルトである必要があります。有料部分はtmpprです

クラウド — ラップトップでは不十分な場合、またはチームで共有時間が必要な場合のクラウド CI ランナー、およびプッシュ、フェッチ、レビューに共有オリジンが必要なチームとマシン用のホストされたリモート リポジトリ。
現地宿泊は無料。必要なときにクラウドを利用できます。
tmppr local — レビュー、CI、自分のマシン上のエージェント — は無料です。 tmppr Cloud は有料サーフェスです。共有コンピューティング用のホストされた CI ランナーと、共有オリジン用のホストされたリモート リポジトリです。
127.0.0.1 で自己ホストされます。無制限のローカル リポジトリ、レビュー、CI、エージェント。コードがマシンから離れることはありません。
ネイティブ デスクトップ アプリ — macOS、Linux、Windows
署名付きビルドとローカルアップデート
無制限のローカル リポジトリ、ブランチ、PR
ローカル CI + 信頼できる LAN ランナー (まだ無料)
交換可能なエージェントプロバイダー + ライフサイクルトリガー
内蔵MCPサーバー+プラグインSDK
ローカルだけでは不十分な場合に使用できる 2 つの有料サーフェス (マネージド CI ランナーと共有リモート リポジトリ)。同じレビュー画面、同じ CLI。
クラウド CI ランナー - マネージド プールに CI をバーストし、tmppr クラウド実行からログをストリーミングします。
ホストされたリモート リポジトリ — チームと CI がプッシュ、フェッチ、レビューするための共有オリジン
クラウド ログイン — tmppr クラウド ログイン、SSO 対応
月額 20 ドル。クラウド ランナー + ホストされたリポジトリ。現地宿泊は無料。
ローカル リポジトリ、ローカル PR レビュー、ローカル CI、エージェント、独自のマシンが含まれており、それらは無料です。 tmppr Cloud は、ホストされた CI ランナーとホストされたリモート リポジトリの有料サーフェスです。
コア エンジンは MIT ライセンスを取得し、オープンに開発されています。パッケージ化されたデスクトップ アプリには、インストーラー、署名済みビルド、自動更新チャネル、およびポリッシュが追加されます。必要に応じてソースからビルドしてください。
いいえ。tmppr はデフォルトで 127.0.0.1 にバインドされます。ローカルで使用するためのアカウント、必要なクラウド、またはテレメトリはありません。ホストされたリモート リポジトリ、オプションの更新チェック、または明示的に CI を実行することを選択した場合にのみ、コードがマシンから離れます。

ガー。
シェル コマンドを実行できるエージェントはどれも tmppr を駆動できます。エージェントに tmppr CLI を指示して、PR を開き、レビュー コメントを残し、変更を要求し、CI を再実行し、準備状況を確認します。著者名が記​​録されるため、人間によるレビューとエージェントによるレビューが一目で区別できます。
はい。 @tmppr/plugin-sdk を使用すると、単一のパッケージでエージェント プロバイダー、自動化トリガー、マージ ルール、フック エグゼキューター、イベント サブスクライバー、HTTP ルート、構成を追加できます。 tmppr プラグイン install <pkg> を使用してインストールします。コアへのパッチやフォークとの同期を維持するものはありません。プラグインはローカルで無料で実行できます。
tmppr は、必要な数の自分のマシンで使用できます。信頼できる LAN ランナーを使用すると、ラップトップで確認しながら、より強力なデスクトップで CI を実行できます。プッシュ、フェッチ、レビューのために共有リモート場所が必要な場合は、tmppr でホストされるリモート リポジトリが有料パスになります。
tmppr Cloud run を使用してディスパッチするマネージド CI ランナー プール。ログはリアルタイムでマシンにストリーミングされます。同じワークフロー YAML、同じレビュー画面 — ラップトップではないコンピューティングだけを使用します。ローカルのハードウェアでは不十分な場合、またはチームが共有ランを必要とする場合に使用します。
tmppr Cloud には、クラウド CI ランナー (共有分数またはより大きなボックス用の管理されたコンピューティング プール) とホストされたリモート リポジトリ (チーム用の共有オリジン) という 2 つの有料サーフェスがあります。ローカル リポジトリ、レビュー、CI、エージェントはローカルにあり、無料です。
チャット ウィンドウで AI コードをレビューするのはやめてください。
本物のレビューツールを手に入れましょう。必要に応じてローカルで、またはクラウドで。

## Original Extract

Local GitHub-style pull requests where your agents can open work, run CI, review each other, trigger follow-up agents, and merge through enforced gates.

› tmppr / local PR review Features Plugins How it works Workflow Hosting FAQ Download tmppr Let your coding agents work together through pull requests.
Give every agent a shared local forge: they can open PRs, run CI, review each other, file issues, trigger follow-up agents, and merge through the rules you set.
See how it works Plays nicely with the stack you already use
GitHub Actions Claude Code Cursor Codex (OpenAI) Plugin SDK · new
Built to be extended, not copied.
@tmppr/plugin-sdk lets you extend tmppr without forking it. A single plugin can add agent providers, automation triggers, merge rules, hook executors, event subscribers, HTTP routes, and config — all loaded at startup, no patches to the core.
A coordination layer for autonomous agents, without leaking your code.
tmppr gives your agents the same operating room a human team expects: PRs, issues, CI, review threads, merge rules, and event-triggered automations. They can work independently without disappearing into chat logs.
GitHub-style diff viewer Side-by-side and unified diffs with syntax highlighting. Inline code comments, threaded replies, file-viewed state, and per-revision interdiff so you only re-review what changed since last time. Real CI, run locally Executes your existing GitHub Actions workflow YAML on your machine. Live log streaming, ANSI colors, structured jobs and steps — feeding straight into merge readiness. Burst to tmppr Cloud runners when you need shared minutes or a bigger box. Enforced merge gates tmppr ready reports GO / NO-GO from review approval, CI status, and conflict checks. Real git merge --no-ff in a detached worktree — no surprises. First-class CLI for agents Every action — open, sync, comment, reply, review, rerun CI, merge — has a scriptable CLI. Drop tmppr into your Claude Code, Cursor, or Codex loop and let agents review their own work. Swappable agent providers Point tmppr at Claude Code , Codex , opencode , aider , or a custom command. Set a default for the machine or per-repo. Your agents, your models. Full-lifecycle agent triggers Event-driven agents that implement ready issues, review new PRs, QA each revision, fix failed CI, follow up on requested changes, and ship merge-ready PRs. The whole loop runs itself. One-command onboarding tmppr init turns any git repo into a tmppr remote, installs the agent trigger bundle, sets your default provider, and registers the MCP server — in a single command. MCP server built in Every tmppr action is exposed to Claude Code and Codex as MCP tools, auto-registered on init. Your agents call PRs, CI, reviews, and merges as native tools — no glue code. Event automations Run any terminal command on repo events — pr.created , pr.updated , ci.updated , review.updated , issue.updated , and hook.completed . Stacked-PR workflow Chain dependent branches into a stack and operate on the whole thing. Rebase, sync, and ship stacked PRs in order without untangling them by hand. Your code never leaves your box Binds to 127.0.0
[truncated]
From issue to merge — your agents drive, you approve.
The same pull-request loop you know from GitHub, running entirely on your machine. Scroll to walk through it in the real app.
Open work sits on a board. Agents pick up ready issues; you see exactly what's in flight.
Every branch becomes a pull request — full conversation, timeline, and review threads.
Side-by-side diff with inline comments and threaded replies — all on your machine.
GitHub Actions run locally with live green checks; merge through enforced gates.
Your repo runs its own software lifecycle.
One .tmppr/workflow.yaml turns your repo into an event-driven assembly line. Each lifecycle event — an issue going ready, a PR opening, CI turning red — wakes a role that does its job and emits the next event. The whole chain runs to a merged PR with no human in the loop , governed only by your merge rules.
1 Dispatcher on automation.tick Promotes ready, unblocked issues into the pipeline in bounded batches.
2 Implementer on issue.updated Opens a draft PR in an isolated worktree and writes the code in small commits.
3 Reviewer on pr.created Reads the diff and submits an approval or a change request.
4 QA on pr.updated Builds the branch and runs the full test suite on every new revision.
5 Fixer on ci.updated Diagnoses red CI and pushes a fix back onto the branch.
6 Follow-up on review.updated Addresses requested changes and re-pushes for another pass.
7 Shipper on ci.updated Merges once checks are green and approvals satisfy the merge gate.
# .tmppr/workflow.yaml
id: tmppr-swdlc
hooks:
- id: implement-ready-issue
event: issue.updated
actions:
- type: start_agent
role: implementer
prompt: prompts/implement-ready-issue.md
- id: ship-ready-pr
event: ci.updated
actions:
- type: start_agent
role: shipper Isolated worktrees Every run gets its own git worktree, so agents work in parallel without stepping on each other. Swappable agents Each role runs the provider you choose — Claude Code, Codex, opencode, aider, or your own command. Enforced merge gate Nothing merges until required approvals and status checks pass — agents earn the merge, they don't bypass it. Compared
GitHub PRs Forgejo Gitea Source stays local Yes — 127.0.0.1 No Self-hosted Self-hosted Repo feature parity (code, PRs, issues, wiki, packages, releases, Actions) Yes Yes Yes Yes Gitea/Forgejo-compatible API Drop-in No Native Native Import / mirror from Gitea & Forgejo Yes Import only Yes Yes GitHub-style diff UI Yes Yes Yes Yes Threaded inline comments + reviews Yes Yes Yes Yes CI runs locally on your hardware Built-in Cloud minutes Self-hosted runner Self-hosted runner Git LFS Yes Yes Yes Yes Branch & tag protection Yes Yes Yes Yes CODEOWNERS Yes Yes Yes Yes Webhooks & git hooks Yes Webhooks only Yes Yes Projects / boards Yes Yes Yes Yes Package registry (multi-ecosystem) Yes Yes Yes Yes Stacked-PR workflow Yes Manual No No Enforced merge gates (GO / NO-GO) Yes Required checks Required checks Required checks ActivityPub / ForgeFed federation Yes No Experimental No Built for AI agents Native No No No Scriptable CLI for agents First-class gh CLI tea CLI tea CLI Swappable agent models Claude Code, Codex, opencode, aider, custom No No No Full-lifecycle agent triggers Built in DIY Actions DIY Actions DIY Actions MCP server built in Yes No No No Extensible via plugin SDK No fork needed Apps + Actions Fork / webhooks Fork / webhooks AI agent pricing No local seats Seat-based Free / OSS Free / OSS Paid surface Cloud runners + hosted repos Cloud account seats None — OSS Gitea Cloud From the maker
I have agents writing code for me every day. Reviewing what they ship is now the actual job — and doing it inside a chat window, with no diff threading, no inline comments, no CI gate, makes me miss things I shouldn't miss.
Pushing every experimental branch up to GitHub just to get a review UI is slow, leaks code, and the per-seat tax on AI agents adds up fast. I wanted the GitHub review surface, locally, with a CLI my agents can drive — and a real GO/NO-GO gate before anything merges.
tmppr is that. Local review should be the default. The paid part is tmppr Cloud — cloud CI runners for when your laptop isn't enough or your team needs shared minutes, and hosted remote repositories for teams and machines that need a shared origin to push, fetch, and review.
Local stays free. Cloud when you need it.
tmppr local — review, CI, agents on your own machine — is free. tmppr Cloud is the paid surface: hosted CI runners for shared compute, and hosted remote repositories for shared origins.
Self-hosted on 127.0.0.1 . Unlimited local repos, reviews, CI, and agents. Your code never leaves your machine.
Native desktop app — macOS, Linux, Windows
Signed builds and local updates
Unlimited local repos, branches, and PRs
Local CI + trusted LAN runners (still free)
Swappable agent providers + lifecycle triggers
Built-in MCP server + plugin SDK
Two paid surfaces for when local isn't enough: managed CI runners and shared remote repositories. Same review surface, same CLI.
Cloud CI runners — burst CI to a managed pool, stream logs from tmppr cloud run
Hosted remote repositories — shared origin for teams and CI to push, fetch, and review
Cloud login — tmppr cloud login , SSO-ready
$20/month. Cloud runners + hosted repos. Local stays free.
Local repositories, local PR review, local CI, agents, and your own machines are included — and free. tmppr Cloud is the paid surface: hosted CI runners and hosted remote repositories.
The core engine is MIT-licensed and developed in the open. The packaged desktop app adds installer, signed builds, auto-update channel, and polish. Build from source if you'd rather.
No. tmppr binds to 127.0.0.1 by default. There's no account, required cloud, or telemetry for local use. Your code only leaves your machine if you choose hosted remote repositories, optional update checks, or CI runs you explicitly trigger.
Any agent that can run a shell command can drive tmppr. Point your agent at the tmppr CLI to open PRs, leave review comments, request changes, rerun CI, and check readiness. Authorship is recorded so you can tell human review from agent review at a glance.
Yes. @tmppr/plugin-sdk lets a single package add agent providers, automation triggers, merge rules, hook executors, event subscribers, HTTP routes, and config. Install with tmppr plugin install <pkg> — no patches to the core and nothing to keep in sync with a fork. Plugins run locally and free.
Use tmppr on as many of your own machines as you want. Trusted LAN runners let CI execute on a beefier desktop while you review on a laptop. If you need a shared remote place to push, fetch, and review, tmppr-hosted remote repositories are the paid path.
A managed CI runner pool you dispatch to with tmppr cloud run . Logs stream back to your machine in real time. Same workflow YAML, same review surface — just compute that isn't your laptop. Use it when local hardware isn't enough or when a team wants shared runs.
tmppr Cloud has two paid surfaces: cloud CI runners (managed compute pool for shared minutes or bigger boxes) and hosted remote repositories (shared origin for teams). Local repos, reviews, CI, and agents stay local and free.
Stop reviewing AI code in a chat window.
Get a real review tool. Locally — or in the cloud when you need it.
