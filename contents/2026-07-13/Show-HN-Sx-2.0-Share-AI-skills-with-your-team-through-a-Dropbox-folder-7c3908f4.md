---
source: "https://sleuth-io.github.io/sx/2026/07/10/your-dropbox-is-now-a-skill-server.html"
hn_url: "https://news.ycombinator.com/item?id=48900319"
title: "Show HN: Sx 2.0 – Share AI skills with your team through a Dropbox folder"
article_title: "Your Dropbox is now a skill server | sx blog"
author: "detkin"
captured_at: "2026-07-13T23:43:20Z"
capture_tool: "hn-digest"
hn_id: 48900319
score: 1
comments: 0
posted_at: "2026-07-13T23:26:52Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Sx 2.0 – Share AI skills with your team through a Dropbox folder

- HN: [48900319](https://news.ycombinator.com/item?id=48900319)
- Source: [sleuth-io.github.io](https://sleuth-io.github.io/sx/2026/07/10/your-dropbox-is-now-a-skill-server.html)
- Score: 1
- Comments: 0
- Posted: 2026-07-13T23:26:52Z

## Translation

タイトル: Show HN: Sx 2.0 – Dropbox フォルダを通じて AI スキルをチームと共有する
記事のタイトル: Dropbox がスキル サーバーになりました | SXブログ
説明: sx 2.0 がリリースされました: Mac、Windows、Linux 用のネイティブ アプリで、誰でも AI スキルをチームと共有できます。git もターミナルも必要ありません。

記事本文:
SXブログ
SXブログ
Dropbox がスキル サーバーになりました
sx 2.0 がリリースされました。Mac、Windows、Linux 用のネイティブ アプリで、誰でも AI スキルをチームと共有できます。git もターミナルも必要ありません。
数か月前、私は AI ツール スタックに npm 型の穴があると書きました。優秀な AI ユーザーはスキル、MCP 構成、コマンドを構築して成果を倍増させますが、その知識を配布するクリーンな方法がないため、その知識はマシンに閉じ込められたままになります。私たちはその穴を埋めるために、AI アセット用のオープンソース パッケージ マネージャーである sx を構築しました。うまくいきました。開発者はこれを使用して、Git Vault 内のスキルをバージョン管理し、ロック ファイルと決定的インストールを使用して Claude Code、Cursor、Copilot、Codex、Gemini、Cline、および Kiro 全体にインストールします。
そして私は、アトラシアンが失敗するのを二度見てきたので、いつか来ると予想していたミスを犯してしまいました。私は開発者用の共有レイヤーを構築し、他の全員が最終的にはコマンド ラインで会うことになると想定していました。そうはなりません。今年私たちが実施した約60回の発掘面接では、スキルを最大限に活用している人材はマーケティング、法務、営業、運用の分野で増えています。彼らは素晴らしいスキルを書いています。彼らは git もターミナルも持たず、それらを取得することにも興味がありません。マーケティング チームに sx init --type git を要求することは、スキルを共有しないように要求していることになります。
sx 2.0 が修正です。これは本物のデスクトップ アプリであり、依存する配布モデルはすべてのチームがすでに持っている共有フォルダーです。
共有 Dropbox フォルダがバックエンド全体です
ここでは、技術者以外のチームのワークフローを示します。アプリを開き、Dropbox、Google Drive、OneDrive、または iCloud 内のフォルダーにライブラリを指定して、スキルをドラッグします。Markdown が入力され、スキルが出力されます。チームメイトが同じフォルダーにあるアプリを指定すると、あなたが公開したすべてのものがチームメイトに表示されます。サーブが無い

r でアカウントはありません。会社がすでに料金を支払っているファイル同期製品がレプリケーションを実行します。
これは、2.0 のもう 1 つの大きな変更、ボールト形式 v2 のおかげで機能します。すべてのアセットの最新バージョンは、プレーンで読み取り可能なマークダウンとして、ディスクのassets/<name>/に直接保存されるようになりました。バージョン履歴は、その隣の .sx/versions/ にあります。 Vault を grep できます。 Obsidian で開くことができます。 .claude/skills を直接指定すると、解凍するものが何もないため、そのまま機能します。
明白な比較は、まさに Obsidian のセットアップ、つまり同期フォルダー内のマークダウン ボールトであり、現在多くのチームがその方法でスキルを実行しています。違いは、ファイルの同期後に何が起こるかです。 sx は AI クライアントについてネイティブに認識しています。アプリで [同期] をクリックすると、バックグラウンドで sx インストールが実行されます。何をインストールする必要があるかを解決し、各アセットをすべてのクライアントの形式 (クロード コード スキル、カーソル ルール、コパイロット命令など) に変換し、マシン上の適切な場所に書き込みます。チームメイトが共有フォルダーにスキルをドロップし、ボタンを 1 つクリックすると、そのスキルが AI クライアントに反映されます。その翻訳レイヤーは、マークダウンでいっぱいのフォルダーでは実行できない部分です。
ここで開発者が失うものは何もありません。 CLI は依然として同じ Go バイナリであり、git および skill.new ボールト タイプは引き続き機能し、アプリと CLI は同じボールトを読み取ります。 2.0 ではコレクションが追加され、関連するスキルをグループ化し、読み取り時に解決されるユニットとしてインストールされるため、来月共有コレクションに追加されたスキルは、誰も何も再実行することなく自動的にチーム全体に届きます。
拡張機能。あなたのチームの問題は私のロードマップではないため
このリリースラインの後半は、アプリの拡張システムです。私はプラグイン エコシステムに何年も費やしてきました (共同創設者の Don がプラグイン エコシステムを構築したとき、私は Atlassian にいました)

後継者）チームツールの興味深い問題は常にそのチームに特有のものであり、各チームの答えをコアに詰め込むと製品が肥大化して遅くなるのだと信じています。
これでアプリはプラグイン可能になりました。拡張機能はマニフェストと 1 つの ES モジュールを含むフォルダーであり、必要になるまでビルド手順は必要ありません。拡張機能を使用すると、ダッシュボード ウィジェット、公開時チェック、エディタ コマンド、まったく新しいビューを追加できます。そのうちの 15 個が同梱されているマーケットプレイスがあり、人々が最初にアクセスするのはチームの健康に関するものです。
Collection Doctor はコレクションを 0 ～ 100 でスコア付けし、問題点に名前を付けます。内容が薄い、廃止すべき古いアセット、ほぼ重複したスキル、コンテキストを食いつぶす大きすぎるスキルです。各検出結果は、名前を付けたアセットから 1 回クリックするだけです。
導入ウィジェットには、チームの誰が実際にどのスキルを使用しているかが表示されるため、誰かが 1 週間かけて取り組んだそのスキルが継続して稼いでいるかどうかを確認できます。
Review Rota では、すべてのアセットにその使用頻度に応じたレビュー期限が設定され、チーム全体で公平にレビューがローテーションされます。
コレクションのエクスポートは、Vault の外で共有できるように、コレクションを Claude Code、Codex、または Gemini プラグイン、またはプレーン zip にコンパイルします。
デザインに関する 2 つの決定については、コメントで主張していただければ擁護します。まず、拡張機能はパーミッションゲート型です。拡張機能が明示的に宣言したホスト以外には、ファイルシステムもノードもネットワークもありません。拡張機能を有効にすると、その拡張機能が操作できる正確な内容の単純な言語リストが表示され、更新によってそのリストが変更されるたびに再プロンプトが表示されます。第二に、拡張機能は単なる SX アセットです。これらは、スキルと同じパイプラインを通じて公開、バージョン管理、チームへのスコープ設定、固定、監査を行います。つまり、プラグインがアウトオブバンドで自動更新されるエコシステムのように、拡張機能の更新がレビューをすり抜けて行うことはできません。組織管理者は許可リストに登録するか、サード p を無効にすることができます

ボールト全体にわたる芸術的な拡張機能。マーケットプレイス自体は単なる SX ボールトであるため、アプリを自分のリポジトリに指定すると、プライベート リポジトリが得られます。
sx の背後にある前提は最初のリリースから変わっていません。AI の利益は個人に閉じ込められており、欠けている層は配布です。変わったのは、それを必要としている人についての私のイメージです。 2.0 のラインは、次のスキル作成者の波が開発者ではなくなるという賭けであり、彼らと出会うということは、ネイティブ アプリ、彼らがすでに共有しているフォルダー、そしてその下にパッケージ マネージャーを隠す 1 つの同期ボタンを意味するということです。
すべては Apache-2.0 であり、GitHub: github.com/sleuth-io/sx 上にあります。このアプリは 3 つのプラットフォームすべてのリリース アセットに含まれており、brew install sx で引き続き CLI を取得できます。チームで共有フォルダーの設定を試してみたら、どこで問題が発生するのか聞きたいです。このリリースの一部は、実際の Dropbox の遅延に影響されても生き残れるかどうかはほとんどわかりません。
sx の構築に関するメモ — AI アセット (スキル、ルール、MCP 構成、コマンド) のパッケージ マネージャー。

## Original Extract

sx 2.0 is out: a native app for Mac, Windows, and Linux that lets anyone share AI skills with their team, no git and no terminal required.

sx blog
sx blog
Your Dropbox is now a skill server
sx 2.0 is out: a native app for Mac, Windows, and Linux that lets anyone share AI skills with their team, no git and no terminal required.
A few months ago I wrote that there’s an npm-shaped hole in the AI tooling stack . Your best AI users build up skills, MCP configs, and commands that multiply their output, and that knowledge stays trapped on their machines because there’s no clean way to distribute it. We built sx, an open source package manager for AI assets, to fill that hole. It worked. Developers use it to version skills in git vaults and install them across Claude Code, Cursor, Copilot, Codex, Gemini, Cline, and Kiro with a lock file and deterministic installs.
Then I made a mistake I should have seen coming, because I’ve watched Atlassian make it twice. I built the sharing layer for developers and assumed everyone else would eventually meet us at the command line. They won’t. In the sixty or so discovery interviews we’ve run this year, the people getting the most out of skills are increasingly in marketing, legal, sales, and ops. They write great skills. They have no git, no terminal, and no interest in acquiring either. Asking a marketing team to sx init --type git is asking them to not share skills.
sx 2.0 is the fix. It’s a real desktop app, and the distribution model it leans on is one every team already has: a shared folder.
A shared Dropbox folder is the whole backend
Here’s the workflow for a non-technical team. You open the app, point your library at a folder in Dropbox, Google Drive, OneDrive, or iCloud, and drag your skills in. Markdown goes in, skills come out. Your teammates point the app at the same folder and everything you publish shows up for them. There’s no server and no accounts. The file sync product your company already pays for does the replication.
This works because of the other big change in 2.0: vault format v2. The latest version of every asset now lives directly on disk at assets/<name>/ as plain, readable markdown. Version history lives in .sx/versions/ next to it. You can grep the vault. You can open it in Obsidian. You can point .claude/skills straight at it and it just works, because there’s nothing to unpack.
The obvious comparison is exactly that Obsidian setup, a markdown vault in a synced folder, and plenty of teams do run their skills that way today. The difference is what happens after the files sync. sx knows about the AI clients natively. When you hit Sync in the app, it runs an sx install in the background: it resolves what should be installed for you, translates each asset into every client’s format (Claude Code skills, Cursor rules, Copilot instructions, and so on), and writes it to the right place on your machine. Your teammate drops a skill in the shared folder, you click one button, and it’s live in your AI client. That translation layer is the part a folder full of markdown can’t do for you.
Developers lose nothing here. The CLI is still the same Go binary, the git and skills.new vault types still work, and the app and CLI read the same vaults. 2.0 adds collections, which group related skills and install as a unit resolved at read time, so a skill added to a shared collection next month reaches the whole team automatically without anyone re-running anything.
Extensions, because your team’s problems aren’t my roadmap
The second half of this release line is an extension system for the app. I’ve spent enough years around plugin ecosystems (I was at Atlassian when my co-founder Don built theirs) to believe that the interesting problems in a team tool are always specific to the team, and that cramming every team’s answer into core is how products get bloated and slow.
So the app is now pluggable. An extension is a folder with a manifest and one ES module, no build step until you want one. Extensions can add dashboard widgets, publish-time checks, editor commands, and whole new views. There’s a marketplace that ships with fifteen of them, and the ones people reach for first are the team-health ones:
Collection Doctor scores a collection 0 to 100 and names the problems: thin descriptions, stale assets that should be retired, near-duplicate skills, oversized skills eating context. Each finding is one click from the asset it names.
The adoption widgets show who on the team is actually using which skills, so you can see whether that skill someone spent a week on is earning its keep.
Review Rota gives every asset a review due date that adapts to how heavily it’s used, and rotates reviews fairly across the team.
Collection Export compiles a collection into a Claude Code, Codex, or Gemini plugin, or a plain zip, for sharing outside your vault.
Two design decisions I’ll defend if you push on them in the comments. First, extensions are permission-gated: no filesystem, no Node, no network beyond hosts an extension explicitly declares, and enabling one shows you a plain-language list of exactly what it can touch, re-prompted whenever an update changes that list. Second, extensions are just sx assets. They publish, version, scope to teams, pin, and audit through the same pipeline as a skill, which means an extension update can never sneak past review the way it can in ecosystems where plugins auto-update out-of-band. Org admins can allowlist or disable third-party extensions vault-wide. The marketplace itself is just another sx vault, so pointing the app at your own repo gives you a private one.
The premise behind sx hasn’t changed since the first release: AI gains are trapped in individuals, and the missing layer is distribution. What changed is my picture of who needs it. The 2.0 line is a bet that the next wave of skill authors won’t be developers, and that meeting them means a native app, a folder they already share, and one sync button that hides a package manager underneath.
Everything is Apache-2.0 and on GitHub: github.com/sleuth-io/sx . The app is in the release assets for all three platforms, and brew install sx still gets you the CLI. If you try the shared-folder setup with your team, I want to hear where it breaks. That’s the part of this release I’m least sure survives contact with real Dropbox latency.
Notes from building sx — a package manager for AI assets (skills, rules, MCP configs, commands).
