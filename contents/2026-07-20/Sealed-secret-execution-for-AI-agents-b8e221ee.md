---
source: "https://ironrun.dev/"
hn_url: "https://news.ycombinator.com/item?id=48986254"
title: "Sealed secret execution for AI agents"
article_title: "ironrun - An encrypted workspace for AI agents"
author: "skeehn"
captured_at: "2026-07-20T23:47:31Z"
capture_tool: "hn-digest"
hn_id: 48986254
score: 1
comments: 0
posted_at: "2026-07-20T23:33:32Z"
tags:
  - hacker-news
  - translated
---

# Sealed secret execution for AI agents

- HN: [48986254](https://news.ycombinator.com/item?id=48986254)
- Source: [ironrun.dev](https://ironrun.dev/)
- Score: 1
- Comments: 0
- Posted: 2026-07-20T23:33:32Z

## Translation

タイトル: AI エージェントの封印された秘密の実行
記事のタイトル: Ironrun - AI エージェントのための暗号化されたワークスペース
説明: シークレットと AI エージェント用の暗号化されたワークスペース。キーを 1 か所に保管し、エージェントにタイムボックス アクセスを許可し、エージェントの下に挿入され出力から編集されたシークレットを使用してコマンドを実行します。オープンソース、MIT。

記事本文:
アイアンラン
v0.4.0
☾
☀
GitHub
インストール
人間と AI エージェントのための暗号化されたワークスペース
エージェントを実行します。
あなたの秘密ではありません。
すべての AI エージェントはシェル内で実行され、実際の API キー、データベース URL、トークンに直接接続されます。読んだ内容も忘れません。 Ironrun はそれらを 1 つの暗号化されたワークスペースに封印し、値の代わりにタイムボックス化されたリースをエージェントに渡し、出力に含まれるすべてのものを編集します。エージェントはジョブを終了します。秘密は決して見えません。
エージェントプロンプトのコピー
01
[ 01 / 07 ] 露出
エージェントは読むことができます
あなたが持っているすべての秘密。
キー、データベース URL、トークン: 環境内にある場合、エージェントはそれを読み取ることができます。そしてそれは忘れません。
Claude コードは、~/.claude/projects へのすべてのツール呼び出しを引数ごとに記録します。通じた秘密
現在、そのファイルには 1 つのコマンドが平文で保存されています。
1 つの暗号化されたワークスペース。
値がモデルに到達することはありません。
YAML ファイルをプロジェクトにドロップします。 Ironrun は、ボールト、サブプロセス、および出力を処理します。
エージェントは 1 つのツールを呼び出します。終了コード、タイミング、編集された出力が返されます。内部で実行された秘密の値は決してありません。
シークレットは、マシン上の Ironrun のローカル暗号化保管庫 (AES-256-GCM) に保存されます。または、1Password、Vault、Doppler、または Infisical と連携します。実行時にプルされます。ディスクには書き込まれません。
バージョン：「2」
環境セット : アクティブ
コマンド:
- ID : デプロイ
シークレット: [ STRIPE_SECRET 、 AWS_SECRET ]
2
実行する
クリーンな環境でシェルなしで実行されます。注入するものは何もありません。指定したコマンドのみが実行されます。
コマンド:
- ID : デプロイ
argv : [ ./deploy.sh 、 --prod ]
ttl : 30秒
no_network : true
3
編集
出力のすべてのバイトは、エージェントがそれを認識する前に、ローリング パターン マッチを通じて実行されます。シークレットが標準出力に表示されますか?プロセスを離れる前に消えます。
# エージェントが受け取るもの:
終了コード : 0
その間

ation_ms : 191
stdout : "デプロイされました。key=[編集済み]"
# シークレット値: [編集済み]
03
[ 03 / 07 ] セキュリティモデル
1枚では足りないので8枚重ね。
Ironrun は、エージェントが侵害されていると想定します。すべてのレイヤーが独自に流出をブロックします。
1 つが失敗しても、他のものは保持されます。
エージェントがアクセスを要求し、あなたがタイムボックス リースを承認してから、初めて
封印されたプロセスは秘密に触れることはありません。秘密はそのプロセス内にのみ存在します。
決してシリーズ化されたことはありません。ログに記録されたことはありません。エージェントのコンテキスト内では決してありません。これは、
コマンドのラッパー。それはあなたの秘密とモデルの間にある壁です。
はい、金庫です。
エージェント境界用に構築されています。
Ironrun は、シークレットをローカルの AES-256-GCM ボールトに保存するか、すでに実行している 1Password、Vault、Doppler、または Infisical に連携します。
追加内容: エージェントはタイムボックス化されたアクセスを取得し、その下にシークレットが挿入され、表示されるものすべてから値が編集されます。
対象となるものと対象外となるものは次のとおりです。
アイアンランは多層防御であり、保証ではありません。装甲車ではなくシートベルトです。これは、すでに信頼しているコントロール (スコープ指定されたトークン、読み取り専用データベース ユーザー、サンドボックス エージェント) で構成され、それらが欠落しているレイヤーを追加します。シークレット値がエージェントの環境やコンテキストに入ることは決してないため、ログに記録されたり、キャッシュされたり、そこから抽出されたりすることはありません。
Ironrun はエージェントとシェルの間に位置します。シェルコマンドを実行する場合は、ironrun で動作します。
ネイティブ MCP サポートにより、設定不要でエージェントに 12 のツール (シールドの実行、アクセスの要求、コマンドの提案) が提供されます。初期化をアイロン実行するだけで完了です。
インストールして、 Ironrun init を実行し、
そしてあなたのエージェントはネットワークに接続されています。ポリシー、MCP 構成、および CLAUDE.md が自動的に作成されます。
単一のバイナリ。 macOS と Linux、arm64 と amd64。他に実行するものは何もありません。
# 実行前にチェックサムを検証します
$カール -fsSL https://ironrun.dev/install.sh |バッシュ
2
初期化する
アイアンラン私

nit はプロジェクトを検出し、暗号化された開発環境を作成し、ポリシー、MCP 構成、およびエージェントの指示を書き込みます。コマンドとシークレットをワークスペースに追加します。
$ cd あなたのプロジェクト
$ アイアンラン初期化
- Ironrun.ymlを作成しました
- 暗号化環境開発の作成
- .mcp.jsonを作成しました
- CLAUDE.mdを作成しました
3
密封して実行
エージェントは一度アクセスを要求します。 TUI で承認します。次に、シェルの代わりに run_sealed を呼び出します。秘密が入力され、編集された出力が戻ってきます。
# エージェントは MCP ツールを呼び出します。
run_sealed( "テスト" )
終了コード : 0
stdout : "ok - key=[編集済み]"
エージェントが探さないことを期待するのはやめましょう。
1 つの YAML ファイル。エージェント以下に封印された秘密。 MIT ライセンスを永久に取得します。
エージェントプロンプトのコピー
フッター スクロール ブリッジ (スクロールの中心から外側にヘアラインが描画されます) -->
アイアンラン
シークレットがエージェントのコンテキスト内に残されることなく、エージェントを通じてコマンドを実行します。

## Original Extract

An encrypted workspace for your secrets and your AI agents. Keep keys in one place, grant agents time-boxed access, and run commands with secrets injected below the agent and redacted from output. Open source, MIT.

ironrun
v0.4.0
☾
☀
GitHub
Install
An encrypted workspace for humans and AI agents
Run agents.
Not your secrets.
Every AI agent runs inside your shell, wired straight to your real API keys, database URLs, and tokens. It doesn't forget what it reads, either. ironrun seals them in one encrypted workspace, hands the agent a time-boxed lease instead of the value, and redacts anything that slips into output. The agent finishes the job. It never sees the secret.
Copy agent prompt
01
[ 01 / 07 ] The Exposure
Your agent can read
every secret you've got.
Keys, database URLs, tokens: if it's in your environment, the agent can read it. And it doesn't forget.
Claude Code logs every tool call to ~/.claude/projects, argument by argument. A secret that passed through
one command is sitting in that file, in plaintext, right now.
One encrypted workspace.
The value never reaches the model.
Drop a YAML file in your project. ironrun handles the vault, the subprocess, and the output.
Your agent calls one tool. It gets back exit code, timing, and redacted output. Never the secret values that ran inside.
Secrets live in ironrun's local encrypted vault: AES-256-GCM, on your machine. Or federate to 1Password, Vault, Doppler, or Infisical. Pulled at exec time. Never written to disk.
version : "2"
environment_set : active
commands :
- id : deploy
secrets : [ STRIPE_SECRET , AWS_SECRET ]
2
Execute
Runs with a clean environment and no shell. Nothing to inject through. Only the command you named runs.
commands :
- id : deploy
argv : [ ./deploy.sh , --prod ]
ttl : 30s
no_network : true
3
Redact
Every byte of output runs through a rolling pattern match before the agent sees it. A secret shows up in stdout? Gone before it leaves the process.
# What the agent receives:
exit_code : 0
duration_ms : 191
stdout : "Deployed. key=[REDACTED]"
# secret values: [REDACTED]
03
[ 03 / 07 ] Security Model
Eight layers because one isn't enough.
ironrun assumes your agent is compromised. Every layer blocks exfiltration on its own.
Fail one, the others hold.
The agent asks for access, you approve a time-boxed lease, and only then does a
sealed process ever touch a secret. Secrets exist only inside that process.
Never serialized. Never logged. Never in the agent's context. This isn't a
wrapper around your commands. It's a wall between your secrets and the model.
Yes, it's a vault.
Built for the agent boundary.
ironrun stores secrets in a local AES-256-GCM vault, or federates to the 1Password, Vault, Doppler, or Infisical you already run.
What it adds: agents get time-boxed access, secrets inject below them, and values are redacted from everything they see.
Here is what that covers and what it does not.
ironrun is defense in depth, not a guarantee . A seatbelt, not an armored car. It composes with the controls you already trust (scoped tokens, read-only database users, a sandboxed agent) and adds the layer they miss: the secret value never enters the agent's environment or context, so it can't be logged, cached, or exfiltrated through it.
ironrun sits between your agent and the shell. If it runs shell commands, it works with ironrun.
Native MCP support gives your agent twelve tools with zero config: run sealed, request access, propose commands. Just ironrun init and done.
Install it, run ironrun init ,
and your agent is wired up. The policy, MCP config, and CLAUDE.md get written for you.
A single binary. macOS and Linux, arm64 and amd64. Nothing else to run.
# verifies the checksum before running
$ curl -fsSL https://ironrun.dev/install.sh | bash
2
Initialize
ironrun init detects your project, creates an encrypted dev environment, and writes the policy, MCP config, and agent instructions. Add your commands and secrets to the workspace.
$ cd your-project
$ ironrun init
- Created ironrun.yml
- Created encrypted environment dev
- Created .mcp.json
- Created CLAUDE.md
3
Run sealed
Your agent asks for access once; you approve in the TUI. Then it calls run_sealed instead of the shell. Secrets go in, redacted output comes back.
# Agent calls the MCP tool:
run_sealed( "test" )
exit_code : 0
stdout : "ok - key=[REDACTED]"
Stop hoping the agent won't look.
One YAML file. Secrets sealed below the agent. MIT licensed, forever.
Copy agent prompt
FOOTER SCROLL BRIDGE (hairline draws outward from center on scroll) -->
ironrun
Run commands through your agent without your secrets ending up in its context.
