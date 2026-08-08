---
source: "https://holycode.cloud"
hn_url: "https://news.ycombinator.com/item?id=49225061"
title: "Show HN: HolyCode Cloud – persistent Linux workstations for AI"
article_title: "HolyCode Cloud — a cloud AI workstation with 60+ tools already installed"
author: "CoderLuii"
captured_at: "2026-08-08T20:17:46Z"
capture_tool: "hn-digest"
hn_id: 49225061
score: 1
comments: 0
posted_at: "2026-08-08T19:29:51Z"
tags:
  - hacker-news
  - translated
---

# Show HN: HolyCode Cloud – persistent Linux workstations for AI

- HN: [49225061](https://news.ycombinator.com/item?id=49225061)
- Source: [holycode.cloud](https://holycode.cloud)
- Score: 1
- Comments: 0
- Posted: 2026-08-08T19:29:51Z

## Translation

タイトル: Show HN: HolyCode Cloud – AI 用の永続的な Linux ワークステーション
記事のタイトル: HolyCode Cloud — 60 以上のツールがすでにインストールされているクラウド AI ワークステーション
説明: Claude Code、Codex、Gemini、OpenCode に加え、Postgres、Playwright、Jupyter、pandas、pandoc、ffmpeg がすでにインストールされているクラウド Linux ワークステーション。自分のマシン、つまり作業内容を保存するディスク上でルートします。月額定額、メーターなし。

記事本文:
ホーリーコードクラウド
得られるもの
何がインストールされているか
価格設定
ビョク
ドキュメント
サインイン
始めましょう
常時オン · アイドル時にサスペンド · セッションが実行中の状態で約 1 秒で復帰します
60 以上のツールがすでにインストールされているクラウド AI ワークステーション
Postgres、Playwright、Jupyter、pandas、pandoc、ffmpeg と同じ Linux ボックス上の Claude Code、Codex、Gemini、OpenCode。自分のマシン、つまり作業内容を保存し、どのブラウザからでもアクセスできるディスク上でルートします。回収されるセッションではなく、コンピューターは明日もあなたのものになります。
マシン上で記録された実際のセッション
サンドボックスやシン Web IDE ではなく、独自の VM、永続的なホームおよびワークスペースをルートとします。
最初の起動時から PATH 上にある Claude Code、Codex、Gemini、OpenCode、Cursor、Pi。サインインして行きましょう。
ブラウザのタブです。デスクでランニングを始めて、電車の中でスマートフォンからチェックしましょう。
アイドル状態の場合はサスペンドし、放置した場所にあるすべてのファイルとプロセスで約 1 秒以内に復帰します。正直な利点の 1 つは、数日間アイドル状態が続くと、代わりにプラットフォームがコールドスタートする可能性があることです。1 分近く経つと、ファイルは変更されません。
並行エージェント、無制限のプロジェクト
ラップトップがスリープしている間に、マシンが必要なだけセッションを一晩実行します。
BYOK — 私たちは AI プロバイダーではありません
Claude/OpenAI/Gemini キーは箱に保管されています。私たちはコンピューティングをホストします。 AI に対して二重請求されることはありません。
エージェントがインストールを要求しなくなる
空のボックス上のエージェントは、最初の 20 分間を apt と npm の実行に費やします。これは、すでに存在するツールチェーンを使用して起動します。以下のすべての名前は、取得したマシン上のコマンドとして解決されます。
クロード · コーデックス · ジェミニ · オープンコード
ノード 24 · Python 3.11 · gcc · make · build-essential
npm · pnpm · 糸 · uv · uvx · pip
typescript · tsx · vite · esbuild · eslint · prettier · nodemon · pm2 · dotenv-cli
psql 17 · redis-cli · sqlite3 · プリズマ · drizz

ルキット
jupyterlab · ipython · pandas · numpy · Polars · matplotlib · seaborn
トーチ (CPU) · 文変換 · tiktoken · sqlite-vec · chromadb
playwright + chromium · lighthouse · fastapi · uvicorn · リクエスト · httpx · beautifulsoup4 · http-server · json-server
pandoc · tesseract OCR · ffmpeg · imagemagick · pdftotext · pypdf · python-docx · openpyxl · yt-dlp ·pillow
SDK · ファイルシステム · メモリ · 逐次的思考
git、gh、tmux、vim、nano、ripgrep、fd、bat、jq、tree、htop、rsync、curl、wget、strace、lsof、dig、ssh、zip
wrangler · と root なので、他のものは何でもすぐに使えます
// を実行しているライブ ボックス上のすべてのコマンドを解決することで、2026 年 8 月 5 日に検証されました。
配布イメージ。Dockerfile から読み取られるものではありません。そのうち 60 個は Python ライブラリの前にあります。
Web 上でクロード コードを使用してみませんか?
多くの場合、そうすべきです。 Pro、Max、Team では無料で、VM のコンピューティング料金は別途かかりません。セッションはブラウザを閉じても存続し、Claude アプリから視聴できます。 GitHub リポジトリで週にいくつかのエージェント タスクを実行しますか?それが正しいツールであり、このページではそうでないふりをするつもりはありません。
私たちはあなたに代理店を売るつもりはありません。私たちはそれが組み込まれたマシンをあなたに販売します。同じプラン、同じキー、同じクロード コード — どちらにしても、レート制限はあなたと Anthropic の間であり、ここでそれらを変更するものは何もありません。変化するのは、エージェントが立っている場所です。彼らの VM は再利用されるセッションであり、これは明日もあなたのものになるコンピューターであり、他の 3 人のエージェントがすでにその上にいます。
クラウド セッションは一定期間非アクティブな状態が続くと停止し、セッションの VM は再利用されます。会話履歴が復元された状態で新しい VM を再度開きます。
毎回同じマシン。インストールしたものはインストールされたままになります。ルートとあなたのディスク。
クロード コード、コーデックス、ジェミニ、オープンコード

、カーソルと円周率。 6 つがプリインストールされており、すべて PATH 上にあります。
リポジトリのクローン作成とプル リクエストの作成には GitHub が必要です。他のホストをバンドルとして送信できますが、そのセッションは結果をリモートにプッシュできません。
あなたのボックス、あなたのリモコン。 GitLab、Bitbucket、所有するサーバー上のベア リポジトリ。
ネットワーク アクセスはデフォルトで制限されています。
// 上記のすべての引用は Anthropic 独自の表現であり、2026 年 8 月 4 日に取得したものです。
code.claude.com/docs/en/claude-code-on-the-web →
そして、私たちがやってはいけないことも書き留めておきます。強制的に再起動すると、実行中のセッションが終了します。
バックアップはありません: 何が生き残るか →
毎月定額。メーターはありません。驚くべき請求はありません。
馬力をお選びください。すべてのプランは製品全体です。アップグレードしてマシンを増やすことはできますが、機能のロックを解除することはありません。
エージェントの料金はすでに支払っています。Claude Pro は月額 20 ドル、Max は 100 ドルからです。これが動作するマシンです。料金は 19 ドルから、1 日あたり約 63 セントです。同じプラン、同じキー、タブを閉じても消えない場所。
すべてのプランには、HolyCode 製品全体 (6 つすべてのエージェント CLI、自分のマシン上の root、無制限のプロジェクト、並列セッション、選択したリージョン、TLS を使用した独自のワークステーション ドメイン) が含まれています。階層の違いはリソースとサポートのみです。
// 毎月定額 — マシンのコストは、1 時間稼働しても、1 か月稼働しても同じです。
// オープンソースの HolyCode コンテナの作者によって構築され、Docker Hub で 8,727 回プルされました
プランごとのサポート応答時間 → · キャンセルの仕組み →
私たちはモデルではなくマシンをホストします。
HolyCode Cloud は AI ではなくワークステーションです。既存の Claude、OpenAI、または Gemini キーを持ち込むと、箱に入ったままになります。弊社ではデータを保存したり、プロキシしたりすることはなく、AI に対して二重料金が請求されることもありません。
これは正直なところです。エージェントの料金はすでに支払っているのです。私たちはそれに、実際の常時稼働マシンを実行させます。

ラップトップがスリープしている間、永続的に、あなただけのもので、どのブラウザからでもアクセスできます。
あらゆるエージェントを実行できる、いつでも利用できるクラウド コーディング ワークステーション。鍵はご自身でご持参ください。 AI を搭載した開発者向けに構築されています。

## Original Extract

A cloud Linux workstation with Claude Code, Codex, Gemini and OpenCode plus Postgres, Playwright, Jupyter, pandas, pandoc and ffmpeg already installed. Root on your own machine, a disk that keeps your work. Flat monthly, no meter.

HolyCode Cloud
What you get
What's installed
Pricing
BYOK
Docs
Sign in
Get started
Always-on · suspends when idle · wakes in ~1 second with your sessions still running
A cloud AI workstation with 60+ tools already installed
Claude Code, Codex, Gemini and OpenCode on the same Linux box as Postgres, Playwright, Jupyter, pandas, pandoc and ffmpeg. Root on your own machine, a disk that keeps your work, reachable from any browser. Not a session that gets reclaimed — a computer that is still yours tomorrow.
A real session, recorded on the machine
Root on your own VM, your own persistent home & workspace — not a sandbox or a thin web IDE.
Claude Code, Codex, Gemini, OpenCode, Cursor and Pi on your PATH from first boot. Sign in and go.
It's a browser tab. Start a run at your desk, check it from your phone on the train.
Suspends when idle, wakes in ~1 second with every file and process where you left it. One honest edge: after days idle the platform may cold-start instead — closer to a minute, files untouched.
Parallel agents, unlimited projects
Run as many sessions as your machine takes, overnight, while your laptop sleeps.
BYOK — we're not an AI provider
Your Claude/OpenAI/Gemini key stays on your box. We host compute; you're never double-charged for AI.
Your agent stops asking to install things
An agent on an empty box spends its first twenty minutes running apt and npm. This one boots with the toolchain already there. Every name below resolves as a command on the machine you get.
claude · codex · gemini · opencode
node 24 · python 3.11 · gcc · make · build-essential
npm · pnpm · yarn · uv · uvx · pip
typescript · tsx · vite · esbuild · eslint · prettier · nodemon · pm2 · dotenv-cli
psql 17 · redis-cli · sqlite3 · prisma · drizzle-kit
jupyterlab · ipython · pandas · numpy · polars · matplotlib · seaborn
torch (CPU) · sentence-transformers · tiktoken · sqlite-vec · chromadb
playwright + chromium · lighthouse · fastapi · uvicorn · requests · httpx · beautifulsoup4 · http-server · json-server
pandoc · tesseract OCR · ffmpeg · imagemagick · pdftotext · pypdf · python-docx · openpyxl · yt-dlp · pillow
sdk · filesystem · memory · sequential-thinking
git · gh · tmux · vim · nano · ripgrep · fd · bat · jq · tree · htop · rsync · curl · wget · strace · lsof · dig · ssh · zip
wrangler · and root, so anything else is one apt away
// verified 2026-08-05 by resolving every command on a live box running the
shipping image, not read off a Dockerfile. 60 of them, before the Python libraries.
Why not just use Claude Code on the web?
Often you should. It's free with Pro, Max and Team, there's no separate compute charge for the VM, sessions survive a closed browser, and you can watch them from the Claude app. A few agent tasks a week on a GitHub repo? That's the right tool, and this page won't pretend otherwise.
We're not selling you an agent. We're selling you the machine it lives on. Same plan, same key, same Claude Code — your rate limits are between you and Anthropic either way, and nothing here changes them. What changes is what the agent is standing on: their VM is a session that gets reclaimed, this is a computer that's still yours tomorrow, with three other agents already on it.
Cloud sessions stop after a period of inactivity and the session's VM is reclaimed. You reopen to a fresh VM with your conversation history restored.
The same machine every time. What you install stays installed. Root, and a disk that's yours.
Claude Code, Codex, Gemini, OpenCode, Cursor and Pi. Six preinstalled, all on your PATH.
Repository cloning and pull request creation require GitHub. Other hosts can be sent up as a bundle, but that session can't push results back to the remote.
Your box, your remotes. GitLab, Bitbucket, a bare repo on a server you own.
Network access is limited by default.
// every quote above is Anthropic's own wording, retrieved 2026-08-04 from
code.claude.com/docs/en/claude-code-on-the-web →
and what we don't do is written down too — a hard restart ends running sessions, and
there are no backups: what survives what →
Flat monthly. No meter. No surprise bills.
Pick your horsepower. Every plan is the whole product — you upgrade for more machine, never to unlock a feature.
You already pay for the agent: Claude Pro is $20 a month, Max starts at $100 . This is the machine it runs on, from $19 — about 63¢ a day. Same plan, same key, somewhere that doesn't disappear when you close the tab.
Every plan includes the entire HolyCode product — all six agent CLIs, root on your own machine, unlimited projects, parallel sessions, your choice of region, and your own workstation domain with TLS. Tiers differ only by resources and support.
// flat monthly — the machine costs the same whether it runs an hour or all month
// built by the author of the open-source HolyCode container, pulled 8,727 times on Docker Hub
support response times by plan → · how cancelling works →
We host the machine, not the model.
HolyCode Cloud is the workstation — not the AI. You bring your existing Claude, OpenAI, or Gemini key, and it stays on your box. We never store it, never proxy it, and you're never double-charged for AI.
That's the honest wedge: you already pay for the agent. We give it a real, always-on machine to run on — persistent, yours alone, and reachable from any browser, while your laptop sleeps.
An always-available cloud coding workstation that runs any agent. Bring your own key. Built for developers who ship with AI.
