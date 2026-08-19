---
source: "https://nexusrun.dev/"
hn_url: "https://news.ycombinator.com/item?id=49362263"
title: "AI Agents that are Portable"
article_title: "NexusRun — portable AI units"
image: "https://nexusrun.dev/og.png"
author: "lseidman1104"
captured_at: "2026-08-19T15:22:32Z"
capture_tool: "hn-digest"
hn_id: 49362263
score: 3
comments: 0
posted_at: "2026-08-19T14:38:36Z"
tags:
  - hacker-news
  - translated
---

# AI Agents that are Portable

- HN: [49362263](https://news.ycombinator.com/item?id=49362263)
- Source: [nexusrun.dev](https://nexusrun.dev/)
- Score: 3
- Comments: 0
- Posted: 2026-08-19T14:38:36Z

## Translation

タイトル: ポータブルな AI エージェント
記事のタイトル: NexusRun — ポータブル AI ユニット
説明: AI エージェントを、Docker も Python 環境もセットアップも必要とせず、どこでも実行できる 1 つの小さなポータブル ファイルに変換します。

記事本文:
NexusRun — ポータブル AI ユニット
ネクサスラン
セクション
見てください
なぜ
スタート
GitHub ↗
ポータブル AI エージェント
AI エージェント全体
1 つの小さなファイルに
ハードウェア上で動作します。
そのモデル、プロンプト、ツール、および資格情報が 1 つのポータブルなパッケージになります。
同僚に渡したり、キオスクに押し込んだり、店舗まで運んだりできます。
エアギャップされたマシンを USB スティックに接続すると、そのまま実行されます。
ドッカーはありません。 Python環境はありません。トークンごとの請求やプロンプトはありません
ネットワークから離れます。
チャットボットではありません。作業を行うエージェント。
ここでは、行って答えを見つけることで質問に答えています。
ラップトップ。ネットワークも API キーもありません。モックアップではなく実際の出力。
nexus run Notes-agent -p "起動ウィンドウはいつ開きますか?"
llama.cpp/server経由でCPU上でnotes-agent:1.0.0を実行
MCP サーバーの起動に関する注意事項
ノート v1.0.0 — 2 つのツール:notes__list_notes、notes__read_note
→notes__list_notes ({})
✓ Notes__list_notes — launch.txt sites.txt
→notes__read_note ({"名前":"launch.txt"})
✓ Notes__read_note — 開始期間は 2027 年 3 月 14 日に開きます。
打ち上げ期間は 2027 年 3 月 14 日に開きます。
— 54 トークン · 7.0 トーク/秒 · 3 ターンにわたる 2 つのツール呼び出し
# nexus.yaml — エージェント全体
APIバージョン: nexusrun.dev/v1
名前: メモエージェント
バージョン: 1.0.0
モデル:
＃すでにオラマで引いていますか？再利用する
- ソース: ollam:llama3.1:8b
エントリポイント:
タイプ: チャット
システムプロンプト: |
メモツールを使用して回答します。
mcp_servers:
注：
# 取得、固定、サンドボックス化
ソース: github:acme/notes-mcp#v1.0.0
コマンド: ["ノード", "server.js"]
サンドボックス:
allowed_paths: [ /home/me/notes ]
秘密:
# ここで宣言されていますが、ここには保存されません
- 名前: API_KEY
必須: true
Dockerfile ではなく、読み取り可能な 1 つのファイル
エージェントがどのように行動するかを決定するすべてがあなたの目の前にあります。それ
モデルを含むのではなくモデルを指すため、ファイルはそのままになります。

うーん
5 ギガバイトではなく 1 キロバイトです。
ツール サーバーは依存関係であり、他のサーバーと同様に固定されています。誰もそうする必要はありません
最初に何かをインストールすると、それぞれがカーネルによって
リストしたパス — 未宣言は拒否されたことを意味します。
秘密はアーティファクトに含まれない
ファイルには、どのキーが必要かが記載されています。値は暗号化されて保存されます。
そのため、エージェントは安全にコミットでき、安全に共有できます。
それはあなたが制御できないマシンのために作られています。
他人のラップトップ、混合ハードウェアのラック、店舗のキオスク、ボックス
インターネットを見たことがない人。エージェントが実際に行う必要があるのはそこです
そして、通常の想定が静かに失敗する場所です。
データがすでに存在する場所で実行されます
レコードを保持するマシン上、VLAN の背後、または
ネットワークがまったくありません。モデルを密閉し、1 つのファイルを横に運びます。何もない
がアップロードされ、頻繁に実行しても追加料金はかかりません。
検証済み: 空のマシンにインポートされ、ネットワークをまったく使用せずに生成されました。
間違ったものを黙って使用することはありません
インストールされているランタイムが駆動するように構築されていない優れた GPU が通常です。
この場合、ほとんどのツールは何も言わずに CPU にフォールバックします。こいつはこう言う
したがって、そのマシン上のエージェントをスコア付けします。答えは次のように変化するためです。
プロンプトだけでなく、ハードウェアとモデルも含めてください。
検証済み: スコア 0/3 のモデルは拒否され、より良いモデルがマシンごとに自動的に選択されます。
他人のエージェントを安全に実行できる
ファイルを読み取り、その内容を正確に確認して、それを理解した上で実行します。
カーネルはそれを保持します。ツール、ファイル、ネットワーク — 何もありません
可能であるとは宣言していません。
サンドボックス内から検証: 野良ファイルの書き込みや送信接続はありません。
NexusRun は、llama.cpp や Ollama を置き換えるものではなく、それらの上で実行されます。何
さらに、パッケージ化、共有、そして明確な答えも追加されています。
Wr

1 つの小さな YAML ファイルです。プロンプトを編集して独自のものにします。
フォルダーから直接アクセスできるため、プロンプトを変更して再度実行できます。
数百バイト、すでに使用しているレジストリへ。にネットワークがありません
もう一方の端は？ nexus build --seal パック
モデル化されているため、USB スティックで移動できます。
ランタイムは永久に無料です。 NexusRun Cloud のホストとサイン
実際のフリートに出荷するチームのユニット —
早期アクセスが公開されています。

## Original Extract

Turn an AI agent into one small portable file that runs anywhere — no Docker, no Python environment, no setup.

NexusRun — portable AI units
nexusrun
Sections
See it
Why
Start
GitHub ↗
Portable AI agents
An entire AI agent
in one small file
that runs on your hardware.
Its model, prompt, tools and credentials become a single portable package.
Hand it to a colleague, push it to a fleet of kiosks, or carry it to an
air-gapped machine on a USB stick — and it just runs.
No Docker. No Python environment. No per-token bill, and no prompt
leaving your network.
Not a chatbot. An agent that does the work.
Here it is answering a question by going and finding the answer — on a
laptop, with no network and no API key. Real output, not a mockup.
nexus run notes-agent -p "When does the launch window open?"
running notes-agent:1.0.0 on CPU via llama.cpp/server
starting MCP server notes
notes v1.0.0 — 2 tool(s): notes__list_notes, notes__read_note
→ notes__list_notes ({})
✓ notes__list_notes — launch.txt sites.txt
→ notes__read_note ({"name":"launch.txt"})
✓ notes__read_note — The launch window opens on 14 March 2027.
The launch window opens on 14 March 2027.
— 54 tokens · 7.0 tok/s · 2 tool call(s) over 3 turns
# nexus.yaml — the whole agent
apiVersion: nexusrun.dev/v1
name: notes-agent
version: 1.0.0
models:
# already pulled with Ollama? reuse it
- source: ollama:llama3.1:8b
entrypoint:
type: chat
system_prompt: |
Answer using the note tools.
mcp_servers:
notes:
# fetched, pinned and sandboxed for you
source: github:acme/notes-mcp#v1.0.0
command: ["node", "server.js"]
sandbox:
allowed_paths: [ /home/me/notes ]
secrets:
# declared here, never stored here
- name: API_KEY
required: true
One readable file, not a Dockerfile
Everything that decides how the agent behaves is in front of you. It
points at the model rather than containing it, so the file stays about
a kilobyte instead of five gigabytes.
Tool servers are dependencies, pinned like any other. Nobody has to
install anything first, and each one is confined by the kernel to the
paths you listed — undeclared means denied.
Secrets stay out of the artifact
The file says which key it needs. The value lives encrypted on the
machine that runs it, so the agent is safe to commit and safe to share.
It's built for machines you don't control.
Someone else's laptop, a rack of mixed hardware, a kiosk in a shop, a box
that has never seen the internet. That is where agents actually have to
work, and where the usual assumptions quietly fail.
It runs where your data already is
On the machine that holds the records, behind the VLAN, or on a box with
no network at all. Seal the model in and carry one file across. Nothing
is uploaded, and running it more often costs nothing extra.
Verified: imported to an empty machine and generated with no network at all.
It won't quietly use the wrong one
A good GPU the installed runtime was never built to drive is the normal
case, and most tools fall back to the CPU without a word. This one says
so — then scores the agent on that machine, because answers change with
the hardware and the model, not just the prompt.
Verified: a model that scored 0/3 there was rejected and a better one chosen, automatically, per machine.
Safe to run someone else's agent
Read the file, see exactly what it will reach for, then run it knowing
the kernel holds it to that. Its tools, its files, its network — nothing
it did not declare is possible.
Verified from inside the sandbox: no stray file writes, no outbound connections.
NexusRun doesn't replace llama.cpp or Ollama — it runs on top of them. What
it adds is the packaging, the sharing, and the straight answers.
Writes one small YAML file. Edit the prompt to make it yours.
Straight from the folder, so you can change the prompt and go again.
A few hundred bytes, to any registry you already use. No network at the
other end? nexus build --seal packs the
model in so it travels on a USB stick.
The runtime is free forever. NexusRun Cloud hosts and signs
your units for teams shipping to real fleets —
early access is open .
