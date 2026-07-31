---
source: "https://getbrainstorm.online/"
hn_url: "https://news.ycombinator.com/item?id=49128167"
title: "Show HN: Brainstorm – a local-first, AI-native OS for knowledge work"
article_title: "Brainstorm — a local-first, AI-native operating system for knowledge management, with end-to-end encrypted sync"
author: "th3-br41n"
captured_at: "2026-07-31T20:54:27Z"
capture_tool: "hn-digest"
hn_id: 49128167
score: 1
comments: 0
posted_at: "2026-07-31T20:26:56Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Brainstorm – a local-first, AI-native OS for knowledge work

- HN: [49128167](https://news.ycombinator.com/item?id=49128167)
- Source: [getbrainstorm.online](https://getbrainstorm.online/)
- Score: 1
- Comments: 0
- Posted: 2026-07-31T20:26:56Z

## Translation

タイトル: Show HN: Brainstorm – ナレッジワーク用のローカルファースト、AI ネイティブ OS
記事のタイトル: Brainstorm — エンドツーエンドの暗号化された同期を備えた、ナレッジ管理のためのローカル ファーストの AI ネイティブ オペレーティング システム
説明: Brainstorm は、ナレッジ管理のためのローカル ファーストの AI ネイティブ オペレーティング システムです。アプリ、データ、AI はすべてマシン上で実行され、オプションでエンドツーエンドの暗号化同期が行われ、すべてのアプリとすべてのエージェントは許可されたもののみにアクセスします。
HN テキスト: こんにちは。私は AI ネイティブ企業として、ローカル ファーストのオープンソース ナレッジ マネジメント オペレーティング システムを 3 か月間構築してきました。Hacker News コミュニティからの本物のフィードバックをぜひ聞きたいと思っています。私は以前、同様のプロジェクトに数年間携わったことがありますが、共通のナレッジ管理アプリですべてが接着されており、1 つを修正すると他のユーザー エクスペリエンスが損なわれることが多いという問題を解決しようとしました。私が構築しようとしていたのは、物事を切り離し、アプリを第一級市民にし、たとえばジャーナリングやホワイトボードの使用など、単一のユーザー ワークフローを担当させることです。私が取り組んでいたアプリにはデータ主権を維持するために厳重なセキュリティ対策が講じられていたため、セキュリティをシステムの中核に組み込むことに決めました。そのため、すべてのアプリはサンドボックス化され、独自の権限と機能を備えた小さなソフトウェアになります。さらに、システムは現在、アプリ、テーマ、エージェントといった資産のマーケットプレイスとなっています。もう 1 つの大きな特徴は、エージェントと人間が協働するオペレーティング システムでありながら、設計されたシステム内でエージェントが行うすべてのことに対して制御層と可観測性層があるため、データ管理が本物のままであることです。言及する価値があるのは、システムはインストールするアプリに応じて進化し、各アプリは他のアプリの MCP ツールとして機能することです。

エージェントアプリがインストールおよび設定されている場合は、Notes でエージェント機能を使用できますが、そうでない場合は、これらの機能がなくてもアプリを通常どおり使用できます。フィードバックや機能リクエストをお待ちしております。 Web サイト: https://getbrainstorm.online Github: https://github.com/brainstorm-os セキュリティ層を検証してフィードバックをお寄せいただくこともできます。大変感謝しています。

記事本文:
訪問数をカウントし、どのページが読まれたかを確認したいと考えています。つまり、デバイス上にランダムな ID を 1 つ保持し、2 ページ目を読んでも 2 人目としてカウントされず、ページ名を EU 内のサーバー上の Amplitude に送信することになります。選択するまで何も保持されず、いつでも好きなときに変更できます。私たちが保管するもの
Brainstorm は無料のオープンソースです — 役に立ったら、GitHub でスターを付けてください Brainstorm 概要 アプリ 仕組み 価格 ロードマップ EN DE FR ドキュメント ダウンロード ◇ ローカルファースト · AI ネイティブ
あなたと AI のためのデスクトップ OS。
Brainstorm は、アプリ、データ、AI を自分のマシン上で実行します。必要なものをインストールし、すべてのファイルをディスク上に保存し、AI に手渡した部分のみを支援させます。
ドキュメントを読む ツアーを見る プレイ中に問題がありますか? YouTube で開く ↗ ローカルファースト · アカウントは不要 · 実際の動作をご覧ください ↓
メモアプリではありません。
これは、あなたが知っているすべてのものを備えたオペレーティング システムです
画面は実際のデスクトップ、壁紙、アプリのアイコン、ウィンドウです。各アイコンはサンドボックス化された独自のアプリであり、すべてがユーザーと一緒に存在する 1 セットのデータから動作します。つまり、間に他人のクラウドを介さずにナレッジ管理を行うことができます。
シェルはアプリのみをホストします。メモ、データベース、ファイル、グラフ、カレンダー、コード エディター - 必要なものを追加し、残りを削除し、それぞれを独自に更新します。
すべてのドキュメントは、あなたが所有するフォルダー内のプレーン ファイルです。同期はオプションであり、エンドツーエンドで暗号化されます。リレーはスクランブルされたバイトのみを認識します。サインアップするアカウントがありません。
シェルはアプリと AI の間に位置します。独自の API キーまたはローカルで実行されるモデルを使用して、各アプリとエージェントが何にアクセスできるかを正確に決定します。
オブジェクトの 1 つのセット。
見方ごとに異なるアプリ
各アプリは同じボールトの読み取りと書き込みを行うため、ここに書いたメモはそこに行があり、gr 内のノードになります。

aph — 決してコピーやエクスポートではありません。
シェル デスクトップが開きます。アプリは、選択した壁紙の上のランチャー内に表示されます。シェルのすべての仕事はそれらをホストすることです。シェルはメモやタスクが何であるかを知りません。それらはアプリ内に存在します。シェル デスクトップに開きます データベース 1 つの場所、あらゆる種類のメモ すべてにリンクするメモ タスク プロジェクト、仕事のやり方 カレンダー 同じデータからの週 連絡先 人々、そしてそのつながり グラフ 仕事の形を見る ホワイトボード 無限のキャンバスで考える ブックマーク Web を保存し、タグを付け、検索する ジャーナル 毎日の習慣、マシン上に常駐するエージェント AI が組み込まれている すべてのアプリ、完全に説明 — 機能とスクリーンショット →
いくつかの決定事項がどこにでも保管されている
あなたとエージェントは 1 か所から作業できます
ほとんどのツールは、AI をその上にボルトで固定します。つまり、他人のクラウドにあるボタンが、自分が所有していない、何をしたのか記録のないデータを操作します。ブレーンストーミングは別の場所から始まります。
アプリが到達できる範囲を決定するのと同じ台帳が、エージェントが到達できる範囲を決定します。すべての許可は具体的で、ログに記録され、簡単に取り消すことができます。システムがリクエストを検証できない場合は、単に「いいえ」と表示されます。
データとキーはマシン上に残るため、見知らぬ人のサーバーに送信することのない AI への手作業が可能になります。これは、クラウド ツールでは提供できない部分です。暗号化されたデータ、ハードウェア上のモデル、そして誰が何をしたかの証跡です。
今日はここにあります: AI ブローカー、来歴、予算。独自のアイデンティティと履歴を持つエージェントは、すでに出荷されているかのように着飾ることなく、以下のロードマップに載っています。
01 バンドルされている機能ではなく、お客様が選択するアプリ。
Notes、データベース、ファイル、グラフをインストールします。カタログまたは URL からサードパーティ アプリをインストールします。使わないものはアンインストールしてください。殻は小さいままです。
すべてのドキュメントはマシン上のファイルです。同期はオプションです

Yjs CRDT を使用します。いつでも標準形式にエクスポートできます。アカウントは必要ありません。
03 必要なときに AI を使用し、必要のないときは AI を使用しません。
シェルはすべての AI 呼び出しを仲介します。独自のプロバイダー キーを使用するか、シェルに同梱されているローカル モデルを実行します。アプリごとの予算によって暴走コストが制限されます。
04 共有ワークスペースを汚さずにカスタマイズします。
データベース ビュー、ダッシュボード レイアウト、ショートカット バインディング、テーマ - デフォルトではすべて個人用です。明示的な「チームとの共有」は、必要に応じて組織スコープに昇格します。
サンドボックス化され、機能ゲート化され、安定した契約が行われます。コード アプリを使用して Brainstorm 内からアプリを構築し、ローカル パッケージまたは URL からインストールします。
すべての決定、トレードオフ、未解決の質問は、出荷前に書き留められ、議論されます。製品はそれらのメモの結果であり、その上のマーケティング層ではありません。
私たちはオープンに構築します。計画と未解決の質問はすべてリポジトリにあります。現在のビルドの内容、現在進行中の内容、および今後の展開を以下に示します。
ビルド内シェルと 20 個のアプリ
機能台帳、ボールト、およびすべてが 1 つのオブジェクト スペースを共有する 20 個のファーストパーティ アプリを備えたサンドボックス ホスト。
ビルドではローカル データ、暗号化された同期
すべてのドキュメントはディスク上の CRDT です。マルチデバイス同期はオプションで、エンドツーエンドで暗号化されます。キーがデバイスから離れることはありません。
すべての AI 呼び出しに対する 1 つのパス (キーまたはローカル モデル) には、すべてのオブジェクトのレコードとすべてのアプリの予算が含まれます。
実際に利用できる無料の暗号化されたマルチデバイス コア。macOS、Windows、Linux で利用可能になり、ワンクリックで既存のメモを共有できます。
独自のアイデンティティ、独自の権限、および実行履歴を持つ AI、つまりエージェント上で実行されるチーム用のオペレーティング システムです。
ロードマップ上 モバイルコンパニオン
ポケットの中に保管庫 — 読み取りとキャプチャのアプリ

r iOS と Android は同じエンドツーエンドの暗号化チャネル経由で同期するため、外出先で作成したメモは同じオブジェクト空間に保存されます。
ロードマップ上 アプリ マーケットプレイス
今すぐオープン SDK に基づいて構築し、サードパーティ アプリを検出、インストール、公開します。それぞれのアプリは同じサンドボックス内で、同じ機能許可を制御して実行されます。
製品は無料です。インフラストラクチャはオプションです。
すべてのアプリ、独自の同期、独自の AI キーなど、マシン上で実行されるものはすべて永久に無料です。有料プランでは、当社が運営するホスト型インフラストラクチャが追加され、商用リリースとともに提供されます。
製品全体をあなたのマシン上で。
すべてのバンドルアプリ、無制限のローカルストレージ
独自のネットワーク経由で同期 – サーバーは不要
セルフホスト型リレーとピアツーピア共有
ローカル AI と独自のプロバイダー キーを追加
すべてをインポートおよびエクスポートします - ロックインはありません
$49.99 / 年 — 2 か月間無料
あなたのデバイスはシームレスに通話します。
ホスト型エンドツーエンド暗号化リレー、最大 5 台のデバイス
暗号化キーのバックアップとリカバリ
パワーユーザー向けの余裕と優先順位。
最大 15 台のデバイス、優先リレー ルーティング
月額 5 ドルのバンドル AI クレジット
365 日の同期履歴と優先サポート
一緒に考えるグループのための共有スペース。
Pro のすべてをチーム全体でプール
ロールベースのアクセスを備えた共有組織スペース
39 ドル/シート/月のベースライン — SSO (SAML / OIDC)、SCIM プロビジョニング、カスタム DPA、構成可能な監査保持、およびオンプレミスのリレー オプション。ボリュームディスカウントは料金表に公表されており、暗黙の交渉ではありません。
無料はお試しではありません。ローカル永続性、バンドルされているすべてのアプリ、ピアツーピア同期、セルフホスティング、独自の AI キー、完全なインポート/エクスポートは永久に無料です。
ホスト型プランは商用リリースとともに開始されます。表示価格は発売目標です。
開いた瞬間にホストされた同期が必要ですか?
無料、ローカルファースト、そしてあなたのもの。 D

macOS、Windows、または Linux 用のパブリック ベータ版を独自にロードします。
ドキュメントを読む アカウントは必要ありません · ボールトはあなたが管理するフォルダーです · AGPL-3.0 に基づく無料のオープンソース
ナレッジワークのためのローカルファーストの AI ネイティブ オペレーティング システム。データはマシン上に存在します。アカウントは必要ありません。

## Original Extract

Brainstorm is a local-first, AI-native operating system for knowledge management. Your apps, your data, and your AI all run on your machine, with optional end-to-end encrypted sync — and every app and every agent only touches what you allow.

Hello, I've been building a local-first open-source knowledge management operating system for 3 months as an AI-native company and would love to hear genuine feedback from the Hacker News community. I've previously worked on a similar project for several years and tried to solve a problem that everything is glued together in common knowledge management apps and fixing one thing usually led to breaking other user experiences. What I was trying to build is disconnecting things and making apps first class citizens, responsible for one single piece of user workflow, like journalling or using whiteboards for example. As the app I was working on had heavy security measures to preserve data sovereignty, I decided that I want the security to be weaved into the system at the core, so all apps are small pieces of software, sandboxed and with their own permissions and capabilities. Furthermore, the system now is a marketplace of assets - apps, themes and agents next. Another major feature is that while being an operating system where agents and humans work together, there is a control and observability layer on everything agents do in the system designed, so the data control remains genuine. Worth mentioning that the system evolves with the apps you install, each app serving as an MCP tool for other apps, for example you can use Agent features in Notes if you have agent app installed and configured, but if you do not you can still use the app normally without these features. Hope to receive some feedback or feature requests. Website: https://getbrainstorm.online Github: https://github.com/brainstorm-os You can verify the security layer and give feedback as well, much appreciated!

We'd like to count visits and see which pages get read. That means keeping one random id on your device, so reading a second page isn't counted as a second person, and sending page names to Amplitude on servers in the EU. Nothing is kept until you choose, and you can change your mind whenever you like. What we store
Brainstorm is free and open source — if it's useful to you, star it on GitHub Brainstorm Overview Apps How it works Pricing Roadmap EN DE FR Docs Download ◇ Local-first · AI-native
A desktop OS for you and your AI.
Brainstorm runs your apps, your data, and your AI on your own machine. Install what you need, keep every file on your disk, and let AI help — only with the parts you hand it.
Read the docs Watch the tour Trouble playing? Open on YouTube ↗ Local-first · No account required · See it in action ↓
It's not a notes app.
It's an operating system for everything you know
The screen is a real desktop — a wallpaper, app icons, windows. Each icon is its own app, sandboxed, all working from one set of data that lives with you: knowledge management without someone else's cloud in the middle.
The shell only hosts apps. Notes, Database, Files, Graph, a calendar, a code editor — add the ones you want, drop the rest, and update each on its own.
Every document is a plain file in a folder you own. Sync is optional and end-to-end encrypted — the relay only ever sees scrambled bytes. There's no account to sign up for.
The shell sits between your apps and any AI. Use your own API key or a model that runs locally, and decide exactly what each app and agent is allowed to touch.
One set of objects.
A different app for each way of seeing it
Each app reads and writes the same vault, so a note you write here is a row there and a node in the graph — never a copy, never an export.
The shell It opens to a desktop Your apps sit in a launcher over a wallpaper you pick. The shell's whole job is to host them — it has no idea what a note or a task is. Those live inside the apps. The shell It opens to a desktop Database One place, every kind of thing Notes Notes that link to everything Tasks Projects, the way you work Calendar Your week, from the same data Contacts People, and how they connect Graph See the shape of your work Whiteboard Think on an infinite canvas Bookmarks Save the web, tag it, find it Journal A daily habit, built in Agent AI that stays on your machine Every app, fully described — capabilities and screenshots →
A handful of decisions, kept everywhere
You and your agents, working from one place
Most tools bolt AI on top: a button in someone else's cloud, working on data you don't own, with no record of what it did. Brainstorm starts somewhere else.
The same ledger that decides what an app can reach decides what an agent can reach. Every grant is specific, logged, and easy to take back — and if the system can't verify a request, it simply says no.
Your data and your keys stay on your machine, so you can hand work to AI you'd never send to a stranger's server. That's the part no cloud tool can offer: your data encrypted, models on your hardware, and a trail of who did what.
Here today: the AI broker, provenance, and budgets. Agents with their own identity and history are on the roadmap below — not dressed up as if they've already shipped.
01 Apps you choose, not features bundled in.
Install Notes, Database, Files, Graph from us. Install third-party apps from the catalog or a URL. Uninstall what you don't use. The shell stays small.
Every document is a file on your machine. Sync is optional and uses Yjs CRDTs. Export to standard formats anytime. No account required, ever.
03 AI when you want it, never when you don't.
The shell brokers every AI call. Bring your own provider key, or run the local model that ships with the shell. Per-app budgets cap runaway cost.
04 Customise without polluting the shared workspace.
Database views, dashboard layouts, shortcut bindings, theme — all personal by default. Explicit "share with team" elevates to org scope when you want it.
Sandboxed, capability-gated, with a stable contract. Build an app from inside Brainstorm with the Code app, and install it from a local package or a URL.
Every decision, trade-off and open question gets written down and argued through before it ships. The product is the result of those notes, not a marketing layer on top.
We build in the open — the plan and the open questions are all in the repo. Here's what's in the current build, what we're on now, and where it goes.
In the build The shell and twenty apps
A sandboxed host with a capability ledger, the vault, and twenty first-party apps that all share one object space.
In the build Local data, encrypted sync
Every document is a CRDT on your disk. Multi-device sync is optional and end-to-end encrypted — your keys never leave your devices.
One path for every AI call — your key or a local model — with a record on every object and a budget on every app.
A free, encrypted, multi-device core you can actually live in — available now for macOS, Windows, and Linux, with a one-click way to bring your existing notes across.
AI that has its own identity, its own permissions, and a history of what it did — an operating system for a team that runs on agents.
On the roadmap A mobile companion
Your vault in your pocket — a read-and-capture app for iOS and Android that syncs over the same end-to-end encrypted channel, so notes you take on the go land in the same object space.
On the roadmap An app marketplace
Build on the open SDK today, then discover, install, and publish third-party apps — each running in the same sandbox, under the same capability grants you control.
The product is free. The infrastructure is optional.
Everything that runs on your machine is free, forever — every app, your own sync, your own AI keys. Paid plans add hosted infrastructure we run for you, and they arrive with the commercial release.
The whole product, on your machine.
All bundled apps, unlimited local storage
Sync over your own network — no server needed
Self-hosted relay and peer-to-peer sharing
Local AI, plus bring-your-own provider keys
Import and export everything — no lock-in
$49.99 / year — two months free
Your devices, talking seamlessly.
Hosted end-to-end encrypted relay, up to 5 devices
Encrypted key backup and recovery
Headroom and priority for power users.
Up to 15 devices, priority relay routing
$5 / month of bundled AI credits
365-day sync history and priority support
Shared spaces for groups that think together.
Everything in Pro, pooled across the team
Shared org spaces with role-based access
$39 / seat / month baseline — SSO (SAML / OIDC), SCIM provisioning, custom DPA, configurable audit retention, and an on-prem relay option. Volume discounts are published on the rate card, not negotiated in the dark.
Free is not a trial. Local persistence, every bundled app, peer-to-peer sync, self-hosting, your own AI keys, and full import/export stay free, forever.
Hosted plans open with the commercial release; listed prices are launch targets.
Want hosted sync the moment it opens?
Free, local-first, and yours. Download the public beta for macOS, Windows, or Linux.
Read the docs No account required · Your vault is a folder you control · Free and open source under AGPL-3.0
A local-first, AI-native operating system for your knowledge work. Your data lives on your machine — no account required.
