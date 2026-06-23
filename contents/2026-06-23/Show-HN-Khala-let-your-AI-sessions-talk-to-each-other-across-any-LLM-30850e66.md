---
source: "https://khala.to/"
hn_url: "https://news.ycombinator.com/item?id=48644317"
title: "Show HN: Khala – let your AI sessions talk to each other, across any LLM"
article_title: "Khala. The messenger between AI tools."
author: "lanakim9410"
captured_at: "2026-06-23T13:52:04Z"
capture_tool: "hn-digest"
hn_id: 48644317
score: 1
comments: 0
posted_at: "2026-06-23T13:00:08Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Khala – let your AI sessions talk to each other, across any LLM

- HN: [48644317](https://news.ycombinator.com/item?id=48644317)
- Source: [khala.to](https://khala.to/)
- Score: 1
- Comments: 0
- Posted: 2026-06-23T13:00:08Z

## Translation

タイトル: Show HN: Khala – AI セッションを LLM 間で相互に通信させます
記事のタイトル: カーラ。 AIツール間のメッセンジャー。
説明: Khala は、AI ワークフロー全体に会話とコンテキストを自動的に伝達します。 MCP がサポートされている場所ならどこでも利用可能です。
HN テキスト: 自分の AI セッション間でコンテキストをベビーシッターし、さまざまな目的でセッション間で概要をコピーアンドペーストしたり MD ファイルを更新したりすることにうんざりしたので、これを構築しました。すべての AI セッションは白紙の状態から始まります。あなたがこれまで築き上げてきたものは何でも、
最後のセッションでは、アーキテクチャの決定、デバッグ履歴、あらゆる選択の背後にある理由など、すべてが失われています。じゃあ一から説明し直すんですね。 Khala はソースでハンドオフを修正します。概要ファイルを作成する代わりに、
送信セッションでは、完全なコンテキストがパッケージ化され、指定された受信箱に直接配信されます。受信セッションは、送信者が中断したところから正確に再開されます。あらゆる MCP 互換ツール (Claude、Cursor、Codex、ChatGPT など) で動作します。デモ: https://youtube.com/shorts/i8xgbrY_Y0g 今のところは招待制ですが、今月無料で試したい場合はコメントを残してください。コードをお送りします。

記事本文:
カラ。 AIツール間のメッセンジャー。私たちのチームについて JA AI ツール間のメッセンジャーになるのはやめましょう
AI セッションは相互に直接メッセージを送ります。
クロードはコンテキストを Codex に渡します。コピーアンドペーストは必要ありません。
MCP がサポートされている場所ならどこでも利用可能です。
khala khala C Codex · auth-build ❯ codex / auth-build ❯ auth-build で何かキューに入れられていますか?未読メッセージはありません。 auth-build の受信トレイを監視して、計画からの引き継ぎを確認します。 ❯ コマンド Codex の場合は「/」と入力します。 ● 先週の火曜日に Claude での機能を計画しました。
あなたは、Khala を通じて Codex に計画を送信しました。
Codex はそれを自動的に受信します。
そしてすぐに開発を開始します。
AI ワークフローは引き継ぎのたびに中断される
あなたはツール間のメッセンジャーです
企画はクロード・コード、出荷はコーデックス。どちらも、相手が何をしたのか知りません。
ハンドオフのたびにワークフローがリセットされる
新しいチャット。概要を貼り付けます。制約を再度説明します。今日すでにこれを 2 回行っています。
ワークフローはタブを閉じると存続できません
ツールを切り替えると糸が切れます。次のセッションは毎回ゼロから始まります。
すべてのコンテキストを 1 つの LLM セッションから別の LLM セッションに直接接続して送信します
クロードで計画し、カーソルで構築するなど、仕事の各部分に専門家によるセッションを設定し、Khala を行ったり来たりして手動で作業させます。
Khala 受信箱を作成し、セッションに接続します
LLM セッションごとに受信箱を登録します。 MCP コネクタを一度貼り付けます。これで、Claude、Cursor、または Codex が Khala を介して送受信できるようになります。
一方のセッションにコンテキストをもう一方のセッションに送信するように指示します。
コンテキストを名前付きセッションに引き渡すように LLM に指示します。完全なプラン、スレッド、またはアーティファクトは、そのセッションの受信箱に自動的に配信されます。
他のセッションが指示に従って受信して実行するのを監視します
受信セッションは受信トレイを読み取り、送信者が中断したところから正確に再開します。再ブリーフィングはありません。

コピー＆ペーストは禁止です。
クロード コードの計画 ソロ ワークフロー用
計画のために Claude Code を実行し、出荷のために Codex を並行して実行します。 Khala はそれらの間で決定を伝達するため、コンテキストを再度説明する必要はありません。計画セッションとコーディング セッションを並行して設定すると、一方で行った決定がそのままもう一方に反映されます。コピー＆ペーストしたり、タブを切り替えたり、追加のツールを使用したりする必要はありません。
あるセッションで診断し、別のセッションで修正します。コンテキストはすでに存在します。
クロード コード カーソル 調査から報告まで 収集→分析→書き込み
ChatGPT はソースを収集し、Claude は同じ概要を使用して合成します。
ChatGPT Claude データ分析 クランチ→可視化→レポート
Claude が数値を処理し、Codex が出力をフォーマットします。
クロード・コーデックス 契約レビュー 監査 → フラグ → 再草案
あるセッションでは監査が行われ、別のセッションでは改訂版の草稿が作成されます。同じドキュメントであり、コピー＆ペーストは不要です。
クロード クロード コンテンツ企画 企画→ドラフト→公開
クロードがカレンダーを計画し、ChatGPT がコピーを作成します。それは、Khala を通じて 1 つの概要です。
Claude ChatGPT デバッグ パイプライン エラー → トレース → パッチ
Claude Code がエラーを追跡し、Codex がパッチを作成して適用します。
Claude Code Codex khala あなたはチームメイトです チーム向け
チームメイト間で仕事を引き継ぐ
クロード コードまたはコーデックス セッションから、Khala を通じてチームメイトにコンテキストを直接渡します。引き継ぎは数秒で行われ、すべての詳細が損なわれず、その間に人的ミスはありません。
フロントエンドは仕様をバックエンドのセッションに送信します。会議はありません。Khala を通じて直接調整します。
フロントエンド開発 バックエンド開発 エンジニア → レビュー担当者 PR 引き継ぎ 構築 → 引き継ぎ → レビュー
完全な PR コンテキストがレビュー担当者のセッションに送信されます。 Slack スレッドも再説明もありません。
エンジニア コードレビュー担当者 デザイナー → 開発仕様書の納品 仕様 → 納品 → 構築
設計仕様は、完全に形成された開発セッションに届きます。ツール間での転写はゼロです。
デザイナーのフロントエンド開発令状

er → エディタパイプライン 書き込み → 編集 → 出荷
草案と完全な概要を編集者のセッションに渡します。編集者はライターが中断したところから正確に再開します。
ライター 編集者 PM → エンジニア チケット引き継ぎ 計画 → チケット → 実装
PM のセッションは仕様を作成してエンジニアのセッションに送信します。チケット システムは必要ありません。
プロダクトマネージャー エンジニア QA → 開発バグレポート 発見 → レポート → 修正
QA のセッションは、完全な再現手順を開発者のセッションに送信します。開発者はそれをすぐに修正するためのすべてを持っています。
QA エンジニア 開発者 対象者
クロード コードの計画 ソロ ワークフロー用
計画のために Claude Code を実行し、出荷のために Codex を並行して実行します。 Khala はそれらの間で決定を伝達するため、コンテキストを再度説明する必要はありません。計画セッションとコーディング セッションを並行して設定すると、一方で行った決定がそのままもう一方に反映されます。コピー＆ペーストしたり、タブを切り替えたり、追加のツールを使用したりする必要はありません。
チームメイト間で仕事を引き継ぐ
クロード コードまたはコーデックス セッションから、Khala を通じてチームメイトにコンテキストを直接渡します。引き継ぎは数秒で行われ、すべての詳細が損なわれず、その間に人的ミスはありません。
モバイルから直接通信
セッションの更新を送信し、エージェントにチェックインし、電話から作業を引き継ぎます。デスクトップは必要ありません。
あらゆるモバイル LLM クライアントで動作します
セッションエイリアスはどこでも機能します
すべてのハンドオフはエンドツーエンドで暗号化されます。 Khala ですらあなたのセッションを読むことはできません。
メッセージはデバイスから送信される前に暗号化されます
セッションコンテンツがKhalaサーバーに保存されることはありません
ピアツーピア配信、仲介者なし
実際のワークフローから得た実際のストーリー
「同じ API に対して 2 つのクロード セッションを開いていました。1 つは私の半分、もう 1 つはもう一方の半分でした。私はセッション間でコピー＆ペーストを繰り返しました。最終的には、大声で尋ねました。二人で自分たちで解決できないのですか? Khala がそれを文字通り可能にしてくれました。」
「チームメイトと私は

インターフェースを同期する必要がありました。会議の代わりに、私は彼に Khala をインストールするように言いました。私たちの 2 つのクロード セッションにより、契約が直接調整されました。私たちは同じ部屋にいる必要はなく、同じタイムゾーンにいる必要さえありませんでした。 」
「 私のチームの誰かが体調を崩して参加できなくなりました。Khala と一緒に、私は彼らのクロード セッションに直接接続し、自分で仕事を受け取りました。プロジェクトは一日も滞りませんでした。 」
「当社のマーケティング戦略は常に反復されています。以前は、更新のたびに上司に説明しなければなりませんでした。現在、彼女は必要なときに、Khala を通じて最終的な計画にアクセスするだけです。 」
「私は新しいセッションごとに自分のブランド、目標、背景を再紹介していました。今では、これらすべてが入った Khala の受信箱が手元にあります。開いたセッションは最初にそれを読み取ります。一気に実際の作業に移ります。 」
「 私たちの設計から開発への引き継ぎには、Figma リンク、Notion ドキュメント、および Slack スレッドが関係していました。現在、私の設計セッションは、Khala を通じてエンジニアのセッションに完全な仕様を送信します。彼らはすぐに構築を開始します。 」
無料アカウントを作成します。ベータ期間中はクレジット カードは必要ありません。
app.khala.to にサインインすると、プロンプト ボックスが表示されます。これをコピーして、MCP がサポートするツールに貼り付けます。 Khala は MCP コネクタとして一発で接続します。
任意のツールでセッションを開いて作業を開始します。 Khala はコンテキストを次のツールに自動的にルーティングします。
LLM セッションとチームメイトを数分で接続します。ベータ版の間は無料で、クレジット カードは必要ありません。
スタークラフトのプロトスが心の間で思考を運ぶために共有するサイオニックリンク、カーラにちなんで名付けられました。
なぜなら私たちはカーラに縛られているからです。私たちのあらゆる思考と感情の神聖な結合。
Khala を使用してエージェントとエージェントを接続します。エージェント向けのメッセンジャー。

## Original Extract

Khala carries conversations and context across your AI workflows automatically. Available anywhere MCP is supported.

I got tired of babysitting context between my own AI sessions and having to copy-paste summaries or update MD files across sessions for different purposes, so I built this. Every AI session starts with a blank slate. Whatever you built up in the
last session, like architecture decisions, debugging history, the reasoning behind every choice, is all gone. So youu re-explain from scratch. Khala fixes the handoff at the source. Instead of writing a summary file,
the sending session packages its full context and delivers it directly to a named inbox. The receiving session picks up exactly where the sender left off. Works across any MCP-compatible tool (Claude, Cursor, Codex, ChatGPT, etc.). Demo: https://youtube.com/shorts/i8xgbrY_Y0g Invite-only for now, drop a comment if you want to try it free this month and I'll send you a code.

Khala. The messenger between AI tools. Meet our team EN Stop being the messenger between your AI tools
Your AI sessions message each other directly.
Claude passes context to Codex, no copy-paste needed.
Available anywhere MCP is supported.
khala khala C Codex · auth-build ❯ codex / auth-build ❯ Anything queued for me on auth-build ? No unread messages. Watching auth-build inbox for a handoff from planning. ❯ Type / for commands Codex ● You planned your feature in Claude last Tuesday.
You sent the plan to Codex through Khala .
Codex receives it automatically.
And starts developing right away.
AI workflows break at every handoff
You're the messenger between your tools
Claude Code for planning, Codex for shipping. Neither one knows what the other did.
Every handoff resets the workflow
New chat. Paste the brief. Re-explain the constraints. You've done this twice today already.
Your workflow can't survive a tab close
Switch tools and the thread snaps. The next session starts from zero, every time.
Connect and send all your contexts directly from one LLM session to another
Wire up a specialist session for each part of the job — planning in Claude, building in Cursor — then let them hand work back and forth through Khala.
Create a Khala inbox and connect it to your sessions
Register an inbox for each LLM session. Paste the MCP connector once — Claude, Cursor, or Codex can now send and receive through Khala.
Instruct one session to send context to the other
Tell your LLM to hand off context to a named session. The full plan, thread, or artifact is delivered to that session's inbox automatically.
Watch the other session receive and execute as directed
The receiving session reads its inbox and picks up exactly where the sender left off — no re-briefing, no copy-paste.
planning claude code For solo workflows
Run Claude Code for planning and Codex for shipping in parallel. Khala carries decisions between them so you never re-explain context. Set up a planning session and a coding session side by side, and the decision you made in one flows straight to the other — no copy-paste, no switching tabs, no extra tooling.
Diagnose in one session, fix in another — context already there.
Claude Code Cursor Research to report Gather → analyze → write
ChatGPT gathers sources, Claude synthesizes with the same brief.
ChatGPT Claude Data analysis Crunch → visualize → report
Claude handles the numbers, Codex formats the output.
Claude Codex Contract review Audit → flag → redraft
One session audits, another drafts revisions — same doc, no copy-paste.
Claude Claude Content planning Plan → draft → publish
Claude plans the calendar, ChatGPT writes the copy — one brief through Khala.
Claude ChatGPT Debugging pipeline Error → trace → patch
Claude Code traces the error, Codex writes and applies the patch.
Claude Code Codex khala you teammate For teams
Hand off work between teammates
Pass context from your Claude Code or Codex session straight to a teammate's, through Khala. The handoff happens in seconds — every detail intact, no human error in between.
Frontend sends the spec to backend's session. No meeting — they align directly through Khala.
Frontend Dev Backend Dev Engineer → Reviewer PR handoff Build → hand off → review
Full PR context ships to the reviewer's session. No Slack thread, no re-explaining.
Engineer Code Reviewer Designer → Dev spec delivery Spec → deliver → build
Design spec arrives in the dev's session fully formed. Zero transcription between tools.
Designer Frontend Dev Writer → Editor pipeline Write → edit → ship
Draft and full brief pass to the editor's session. Editor picks up exactly where writer left off.
Writer Editor PM → Eng ticket handoff Plan → ticket → implement
PM's session writes the spec and sends it to the engineer's session — no ticket system needed.
Product Manager Engineer QA → Dev bug report Find → report → fix
QA's session sends the full repro steps to the dev's session. Dev has everything to fix it immediately.
QA Engineer Developer Who it's for
planning claude code For solo workflows
Run Claude Code for planning and Codex for shipping in parallel. Khala carries decisions between them so you never re-explain context. Set up a planning session and a coding session side by side, and the decision you made in one flows straight to the other — no copy-paste, no switching tabs, no extra tooling.
Hand off work between teammates
Pass context from your Claude Code or Codex session straight to a teammate's, through Khala. The handoff happens in seconds — every detail intact, no human error in between.
Communicate straight from mobile
Send session updates, check in with agents, and hand off work from your phone, no desktop required.
Works in any mobile LLM client
Session aliases work everywhere
Every handoff is end-to-end encrypted. Not even Khala can read your sessions.
Messages encrypted before they leave your device
Session content never stored on Khala servers
Peer-to-peer delivery, no one in the middle
Real stories from real workflows
“ I had two Claude sessions open for the same API — one for my half, one for the other half. I kept copy-pasting between them. Eventually I just asked out loud: can't you two just work it out yourselves? Khala made that literally possible. ”
“ My teammate and I needed to sync our interfaces. Instead of a meeting, I just told him to install Khala. Our two Claude sessions aligned the contract directly. We didn't have to be in the same room — or even the same timezone. ”
“ Someone on my team got sick and couldn't come in. With Khala, I connected to their Claude session directly and picked up the work myself. The project didn't slip a single day. ”
“ Our marketing strategy is always being iterated on. Before, I had to walk my boss through every update. Now she just accesses the finalized plan through Khala whenever she needs it. ”
“ I used to re-introduce my brand, goals, and context every new session. Now I have a Khala inbox with all of that. Any session I open reads it first. I skip straight to the actual work. ”
“ Our design-to-dev handoffs used to involve a Figma link, a Notion doc, and a Slack thread. Now my design session sends the full spec to the engineer's session through Khala. They start building immediately. ”
Create your free account. No credit card needed during beta.
Sign in to app.khala.to and you'll see a prompt box. Copy it and paste into any MCP-supported tool. Khala connects as an MCP connector in one shot.
Open a session in any tool and start working. Khala routes context to the next tool automatically.
Connect your LLM sessions and teammates in minutes. Free while we're in beta, no credit card needed.
Named after the Khala, the psionic link the Protoss of StarCraft share to carry thoughts between minds.
For we are bound by the Khala. The sacred union of our every thought and emotion.
Connect agent to agent with Khala. A messenger for your agents.
