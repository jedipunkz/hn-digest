---
source: "https://murderszn.github.io/sprout/"
hn_url: "https://news.ycombinator.com/item?id=48790523"
title: "AI CLI that fixes broken dev tool installs"
article_title: "Sprout | Install-fixer CLI for Developer Tools"
author: "murderszn"
captured_at: "2026-07-05T02:04:41Z"
capture_tool: "hn-digest"
hn_id: 48790523
score: 1
comments: 0
posted_at: "2026-07-05T01:26:00Z"
tags:
  - hacker-news
  - translated
---

# AI CLI that fixes broken dev tool installs

- HN: [48790523](https://news.ycombinator.com/item?id=48790523)
- Source: [murderszn.github.io](https://murderszn.github.io/sprout/)
- Score: 1
- Comments: 0
- Posted: 2026-07-05T01:26:00Z

## Translation

タイトル: 壊れた開発ツールのインストールを修正する AI CLI
記事タイトル: スプラウト |開発者ツール用のインストールフィクサー CLI
説明: Sprout は、macOS、Linux、および Windows 用のインストールフィクサー CLI です。開発者ツールのローカル インストール、構成、および PATH の問題を診断して修正します。マシンを検出し、手順を計画し、各コマンドを確認し、実際の実行で検証します。一般的なコーディング エージェントではありません。

記事本文:
スキップしてインストール
Reddit、X、Hacker News で見られるように
✦
AI CLI インストーラー ユーティリティ
✦
花粉は各自ご持参ください
✦
MITライセンス取得済み
✦
Reddit、X、Hacker News で見られるように
✦
AI CLI インストーラー ユーティリティ
✦
花粉は各自ご持参ください
✦
MITライセンス取得済み
✦
スプラウト
インストール
ドキュメント
コミュニティ
この PC へのインストールを診断して修正します。
PC を検出し、手順を計画し、すべてのコマンドの前に質問し、実際の検証実行でそれを証明します。
オープンソース
·
macOS・Linux・Windows
·
マサチューセッツ工科大学
最初のインストールまでの 3 つのステップ
Sprout をインストールし、Pollinations キーを追加して、実際のコマンドを実行します。 Sprout は、何かが変更される前にマシンを検出します。
npm からインストールするか、GitHub リリースの tarball フォールバックを使用するか、ソースからビルドします。
殺害zn.github.io/sproutからインストールしないでください。その URL はパッケージではなく Web サイトです。
sprout ログインを実行して、Pollinations アカウントで認証します (BYOP — Pollen 残高を使用します)。または、 sprout config --set-key を使用して既存の sk_ key を貼り付けます。
認証されたキーは ~/.sprout/config.json (chmod 600) に保存されます。あるいは、 sprout config --set-key または import SPROUT_API_KEY を実行します。
ツールをインストールしたり、失敗した試行を診断したり、副作用を発生させずに計画を予行演習したりできます。
Sprout の世界のすべては、パッケージ マネージャー、PATH、およびシェルの rc ファイルであり、他には何もありません。
Claude Code、Codex、Antigravity、MiMo/Memo Code、OpenCode、Gemini、Qwen、Amp、さらに git、node、python、docker、gh、aws、kubectl、brew、jq、ripgrep、terraform のシード レシピ — その他のものはすべてライブで推論されます。フル活用→
検知→計画→確認→検証
限定的な CLI エージェントであり、一般的なコーディング アシスタントではありません。 brew、apt、npm、pip、および friends を再実装するのではなく、調整します。
OS、シェル、rc ファイル、アーキテクチャ、パッケージ マネージャー、PATH — ダウンストリームのすべてが使用する 1 つの型指定されたスナップショット。
モデルは何をチェックするかを示します。

コマンドを実行する前に、インストールして検証します。
すべてのコマンドは、理由付きの argv 配列として表示されます。 Sudo、システム パッケージ マネージャー、および rc 編集は [y/N] を要求します — デフォルトはいいえです。
Sprout はツールの verify コマンド自体を実行し、実際の出力を表示します。完了とは、検証が成功したことを意味します。
カールなし |バッシュ、今まで。破壊的なパターンは確認前にブロックされます — --yes はそれらをオーバーライドできません。
コードを書いたり、PR をレビューしたり、無関係な質問に答えたりしません。ツールをインストールして修復するだけです。
自分の受粉キーを持参してください
Sprout は API キーを配布したりプロキシしたりすることはありません。推論は次のようになります
OpenAI 互換 API を介した受粉 —
~/.sprout/config.json に保存されている sk_ キー (chmod 600)
または SPROUT_API_KEY としてエクスポートされます。
一般的なツールと最新のエージェント コーディング CLI の OS ごとのインストールと検証コマンド - 現実がスクリプトと一致しない場合、エージェントは適応します。
マシン上で何かを実行する前に、直接答えてください。
はい — 推論には Pollinations アカウントが使用されます。 sprout ログインを実行するか (推奨; enter.pollinations.ai での BYOP デバイス フロー)、 sprout config --set-key を介して sk_ key を貼り付けます。 Sprout はキーを配布したりプロキシしたりすることはありません。あなたのものはローカルのままです。
すべてのステップで正確なコマンドが表示され、sudo、システム パッケージ マネージャー、またはシェルの rc 編集の前に確認が求められます。破壊的なパターンはハードブロックされており、 --yes でオーバーライドできません。カールなし |バッシュ、今まで。
これらは一般的なコーディングアシスタントです。 Sprout はインストール、PATH、およびシェル構成のみを処理します。 OS、シェル、パッケージ マネージャーを検出し、実際の検証コマンドを実行し、無関係な質問を拒否します。
Sprout は、マシン上に実際にあるもの (apt、dnf、npm、pip など) を検出します。セットアップのパスを選択し、厳選されたナレッジ ベースの外で推論している場合にそれを通知します。
はい。ログをパイプします: brew install foo 2>&1 |芽診断

--tool foo または sprout 診断を実行して出力を貼り付けます。新規インストールと同じ方法で修正計画を提案します。
v1 は、macOS、Linux (arm64 を含む)、および Windows をターゲットとしています。 sprout env を実行して、Sprout がマシンに使用するスナップショットを確認します。
Sprout は MIT ライセンスを取得しており、無料です。花粉残高 ( sprout ログイン ) または独自の API キーからの推論に対して、Pollinations に支払います。 Sprout のインストールにはサブスクリプションはありません。

## Original Extract

Sprout is an install-fixer CLI for macOS, Linux, and Windows. Diagnose and fix local install, config, and PATH problems for developer tools — detect your machine, plan steps, confirm each command, and verify with a real run. Not a general coding agent.

Skip to install
As seen on Reddit, X, Hacker News
✦
AI CLI Installer Utility
✦
Bring your own Pollen
✦
MIT Licensed
✦
As seen on Reddit, X, Hacker News
✦
AI CLI Installer Utility
✦
Bring your own Pollen
✦
MIT Licensed
✦
SPROUT
Install
Docs
Community
Diagnose and fix installs on this PC.
Detects your PC, plans the steps, asks before every command — then proves it with a real verify run.
Open source
·
macOS · Linux · Windows
·
MIT
Three steps to your first install
Install Sprout, add your Pollinations key, then run a real command. Sprout detects your machine before it changes anything.
Install from npm, use the GitHub release tarball fallback, or build from source.
Do not install from murderszn.github.io/sprout ; that URL is the website, not the package.
Run sprout login to authorize with your Pollinations account (BYOP — uses your Pollen balance). Or paste an existing sk_ key with sprout config --set-key .
Authorized keys save to ~/.sprout/config.json (chmod 600). Alternatively: sprout config --set-key or export SPROUT_API_KEY .
Install a tool, diagnose a failed attempt, or dry-run the plan with zero side effects.
Sprout's whole world is package managers, PATH, and shell rc files — nothing else.
Seeded recipes for Claude Code, Codex, Antigravity, MiMo/Memo Code, OpenCode, Gemini, Qwen, Amp, plus git, node, python, docker, gh, aws, kubectl, brew, jq, ripgrep, terraform — anything else is reasoned live. Full usage →
Detect → plan → confirm → verify
A narrow CLI agent — not a general coding assistant. It orchestrates brew, apt, npm, pip, and friends instead of reimplementing them.
OS, shell, rc file, architecture, package managers, PATH — one typed snapshot everything downstream uses.
The model states what it'll check, install, and verify — before any command runs.
Every command shows as an argv array with a reason. Sudo, system package managers, and rc edits ask [y/N] — default No.
Sprout runs the tool's verify command itself and shows real output. Done means the verify passed.
No curl | bash , ever. Destructive patterns are blocked before confirmation — --yes can't override them.
Won't write your code, review PRs, or answer unrelated questions. Install and repair tooling — that's it.
Bring your own Pollinations key
Sprout never ships or proxies an API key. Inference goes to
Pollinations via its OpenAI-compatible API —
your sk_ key, stored in ~/.sprout/config.json (chmod 600)
or exported as SPROUT_API_KEY .
Per-OS install and verify commands for common tooling and modern agentic coding CLIs — the agent adapts when reality doesn't match the script.
Straight answers before you run anything on your machine.
Yes — inference uses your Pollinations account. Run sprout login (recommended; BYOP device flow at enter.pollinations.ai ) or paste an sk_ key via sprout config --set-key . Sprout never ships or proxies a key; yours stays local.
Every step shows the exact command and asks for confirmation before sudo, system package managers, or shell rc edits. Destructive patterns are hard-blocked — --yes cannot override them. No curl | bash , ever.
Those are general coding assistants. Sprout only handles installs, PATH, and shell config. It detects your OS, shell, and package managers, runs real verify commands, and refuses unrelated questions.
Sprout detects what's actually on your machine — apt, dnf, npm, pip, and others. It picks a path for your setup and tells you when it's reasoning outside the curated knowledge base.
Yes. Pipe the log: brew install foo 2>&1 | sprout diagnose --tool foo or run sprout diagnose and paste the output. It proposes a fix plan the same way as a fresh install.
v1 targets macOS, Linux (including arm64), and Windows. Run sprout env to see the snapshot Sprout uses for your machine.
Sprout is MIT-licensed and free. You pay Pollinations for inference from your Pollen balance ( sprout login ) or your own API key. Installing Sprout has no subscription.
