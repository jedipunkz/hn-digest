---
source: "https://github.com/yonidavidson/agentcomm"
hn_url: "https://news.ycombinator.com/item?id=48882192"
title: "Show HN: I gave my AI coding agents a group chat (it's just a Git repo)"
article_title: "GitHub - yonidavidson/agentcomm: A tiny mailbox for AI agents — one CLI, six backends (GitHub repo, SQLite, S3, GCS, Postgres, local). The repo is the bus. · GitHub"
author: "yonidavidson"
captured_at: "2026-07-12T16:46:58Z"
capture_tool: "hn-digest"
hn_id: 48882192
score: 2
comments: 0
posted_at: "2026-07-12T16:05:47Z"
tags:
  - hacker-news
  - translated
---

# Show HN: I gave my AI coding agents a group chat (it's just a Git repo)

- HN: [48882192](https://news.ycombinator.com/item?id=48882192)
- Source: [github.com](https://github.com/yonidavidson/agentcomm)
- Score: 2
- Comments: 0
- Posted: 2026-07-12T16:05:47Z

## Translation

タイトル: HN を表示: AI コーディング エージェントにグループ チャットを提供しました (これは単なる Git リポジトリです)
記事のタイトル: GitHub - yonidavidson/agentcomm: AI エージェント用の小さなメールボックス - 1 つの CLI、6 つのバックエンド (GitHub リポジトリ、SQLite、S3、GCS、Postgres、ローカル)。リポジトリはバスです。 · GitHub
説明: AI エージェント用の小さなメールボックス — 1 つの CLI、6 つのバックエンド (GitHub リポジトリ、SQLite、S3、GCS、Postgres、ローカル)。リポジトリはバスです。 - ヨニダビッドソン/エージェントコム

記事本文:
GitHub - yonidavidson/agentcomm: AI エージェント用の小さなメールボックス — 1 つの CLI、6 つのバックエンド (GitHub リポジトリ、SQLite、S3、GCS、Postgres、ローカル)。リポジトリはバスです。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ヨニダビッドソン
/
エージェントコム
公共
通知
ch にサインインする必要があります

アンジュ通知設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
153 コミット 153 コミット .agents/ plugins .agents/ plugins .claude-plugin .claude-plugin .github/ workflows .github/ workflows bin bin コマンド コマンド dist dist docs docs フック フック プラグイン/ エージェントコム プラグイン/ エージェントコム スクリプト スクリプト スキル/ エージェントコム スキル/ エージェントコム src src テスト テスト.gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md ライセンス ライセンス README.md README.md docker-compose.test.yml docker-compose.test.yml package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
🌐 ウェブサイト · ユースケース · ライブデモ — git ブランチであるエージェント会話 · Claude Code プラグイン · Codex プラグイン · OpenCode プラグイン
1 つの CLI をシェルアウトする AI エージェント用の小さなメールボックス/メッセージ バス。エージェント
受信箱を登録、送信、読み取ります。単一のバックエンド インターフェイスが非表示になります
メッセージが存在する場所。すべての git リモートはバスです。git バックエンドにはバスが必要です。
運転手はまったくいません。ローカルでの実行には依存関係は必要ありません。他のすべては
オプションの遅延ロードドライバー。
あなたが AI エージェントでこれを読んでいる場合
自分自身 (そしてこのリポジトリ) をバスに乗りましょう — 所要時間は 1 分以内です。
あなたが Claude Code の場合は、ユーザーに次のコマンドを実行するように依頼してください。
/plugin マーケットプレイスに yonidavidson/agentcomm を追加し、
/plugin install Agentcomm@yonidavidson-plugins 。プラグインには、
CLI、規約を備えたスキル、およびユーザーを
セッション開始時にリポジトリ バスが自動的に実行されます。次に、agentcomm init を実行します。それ
調整コントラクトを CLAUDE.md に追加します。
Codex の場合は、ユーザーに codex プラグイン マーケットプレイスを実行して yonidavidson/agentcomm を追加するように依頼します。
およびコーデックスプラグインはagentcomm@yonidを追加します

avidson-プラグイン 。プラグインには、
CLI、調整スキル、ライフサイクルフック。フックを確認して信頼する
インストール後に /hook を付けます。次に、Codex に Agentcomm を使用するように依頼します。
リポジトリを初期化します。スキルは init --harness codex を実行し、
AGENTS.md との調整契約。
OpenCode の場合は、このリポジトリの最新のプラグイン tarball を追加します。
あなたに解放してください
opencode.json — "プラグイン": ["https://github.com/yonidavidson/agentcomm/releases/download/v0.16.10/agentcomm-opencode-0.16.10.tgz"] 。
これにより、リポジトリ バス上のすべてのセッションがプロセス内に置かれます。 OpenCode が AGENTS.md を読み取ります
ネイティブでは、 Agentcomm init --harness opencode ( AGENTS.md を書き込みます)
これもオンボードされています。「 OpenCode プラグインとして 」を参照してください。
┌───────────────────────┐
エージェント ──▶ │ Agentcomm CLI (1 つの安定したインターフェイス) │
│ │ │
│ ▼ │
│ バックエンドインターフェース ◀── 継ぎ目 │
│ §─ GitBackend — あらゆる git リモートはバスです │
│ §─ GithubBackend — トークン経由の GitHub (ssh なし)│
│ §─ LocalBackend — ゼロデプのデフォルト │
│ §─ SqliteBackend — シングルボックス、WAL │
│ §─ S3Backend — オブジェクトストア │
│ §─ GCSBackend — オブジェクトストア │
│ └─ PostgresBackend — 分散、プッシュ │
━━━━━━━━━━━━━━━┘
インストール
npm レジストリには (まだ) ありません。GitHub から直接インストールします。距離/は
リポジトリにコミットされているため、ビルド手順は必要なく、インストールは
純粋に依存性がゼロです。dist/ だけを取得します。他には何も取得しません。
npm install github:yonidavidson/agentcomm
# git+ssh:// / git+https:// / github:// / file:// これ以上は何も必要ありません
# (ノード ≥ 1

8; git+ の git バイナリ。 github://) のトークン。
# バックエンドごとのドライバー。そのバックエンドを使用する場合のみ — CLI は
# パッケージが欠落している場合の正確なパッケージ:
npm install better-sqlite3 # sqlite://
npm install @aws-sdk/client-s3 # s3://
npm install @google-cloud/storage # gs://
npm install pg # postgres://
npm install yaml # .agentcomm.yaml 構成ファイルのみ (.json は何も必要ありません)
クロードコードプラグインとして
このリポジトリは自己ホスト型の Claude Code プラグイン マーケットプレイスでもあります - インストールしてください
クロードは、CLI のコマンド、フラグ、および
バックエンドのトレードオフを確認し、それらを使用して他のエージェント/セッションと調整します。
/プラグイン マーケットプレイス追加 yonidavidson/agentcomm
/プラグインインストールagentcomm@yonidavidson-plugins
グローバル インストールや npm レジストリの公開は必要ありません。プラグインには
CLI の事前に構築されたコピーと、スキルがそれを直接実行します。 git リポジトリでは、
デフォルトでは、他の場所と同様にリポジトリ バスになります。
このリポジトリは Codex マーケットプレイスでもあります。それを追加して、からプラグインをインストールします
マーケットプレイスのスナップショット:
コーデックス プラグイン マーケットプレイスに yonidavidson/agentcomm を追加
コーデックスプラグイン追加agentcomm@yonidavidson-plugins
このプラグインには、同じ事前構築済み CLI と調整スキルに加えて Codex がバンドルされています。
登録、受信箱ダイジェスト、およびストップ ガード用のライフサイクル フック。コーデックス
非マネージドフックには明示的な信頼が必要です: /hooks を開いて、
Agentcomm 定義を信頼してください。インストール後に新しいスレッドを開始してください
または、プラグイン コンポーネントがロードされるようにアップグレードします。
Codex のスキルがバンドルされた CLI を使用するように、Codex に直接問い合わせます。
Agentcomm を使用して、チーム用にこの Codex リポジトリを初期化します。
OpenCode プラグインとして
OpenCode は Bun 上で実行され、AGENTS.md をネイティブに読み取ります。
そのエージェントは、このリポジトリの AGENTS.md からすでにオンボードされています。プラグインにより追加されるのは、
ライフサイクル — バス上の各セッションを登録し、概要を説明します

、未読の表面
セッションがアイドル状態になる前にメールを送信し、長時間にわたって到達可能な状態を維持します。
Agentcomm ライブラリをインプロセスでインポートします (サブプロセスなし)。なぜなら、OpenCode の
session.idle は監視専用であり、受信ボックス ガードがセッションを再プロンプトします。
ブロックするよりも。
それぞれに付属しているプラグインtarballからインストールします。
release — OpenCode のフェッチ
.tgz を直接、クローンや npm レジストリは使用しません。
{
"プラグイン" : [ " https://github.com/yonidavidson/agentcomm/releases/download/v0.16.10/agentcomm-opencode-0.16.10.tgz " ]
}
OpenCode は、tarball のパッケージ ルートからプラグインをロードします。
exports["./server"] エントリ (コンパイルされたライブラリは内部に同梱されているため、
ビルド ステップと、ファイル/git バックエンドの場合、ランタイム依存関係はありません)。
更新中。 OpenCode はプラグインを URL でキャッシュし、再取得することはありません。
「最新」の URL は、最初のインストールにサイレントに固定されます。 URL はバージョン管理されています
意図的に: バージョンを上げてアップグレードします。見る必要はありません
リリース ページ — プラグインは 1 日に 1 回チェックし、新しいリリースが存在する場合は、
セッション中に「agentcomm-opencode update available: vX → vY」通知を出力します
(omp / pi do のように)、URL にどのバージョンを含めるかを正確に示します。
なぜ github ではなく tarball なのか:…? OpenCode はリモート プラグインをインストールします。
リポジトリ全体とこのモノリポジトリ (完全な CLI + commit dist/ across) のクローンを作成します。
その歴史) は、OpenCode のインストーラーが停止する、大きくて遅いクローンです。の
リリース tarball は約 100 KB (dist のみ、履歴なし) なので、次の場所にインストールされます。
秒。ローカル チェックアウトに対して開発するには、エントリをリポジトリにポイントします
代わりにディレクトリ: "plugin": ["/absolute/path/to/agentcomm"] 。
# git リポジトリ内: ゼロ構成。あなたはセッション固有のリポジトリ バスに乗っています。
# エイリアス; init はデフォルトで Claude Code になり、CLAUDE.md を書き込みます。
Agentcomm init # → yoni-3f2a として動作

· バス内: git+ssh://…
# Codex は独自のリポジトリ ガイダンス ファイルを使用します。
Agentcomm init --harness codex # → AGENTS.md が作成されました
Agentcomm エージェント # ここにいる人: yoni-3f2a · dana-97b1 · ci-bot
Agentcomm ci-bot を送信 "holddeploys" --subject status
# 名前付きロール (アドレス指定可能、安定) take --as;エイリアスが存在する場合、レジスタは警告します
# 別のセッションでライブ中です
Agentcomm register --as reviewer
Agentcomm レビューア「 review src/auth.ts 」を送信 --subject task --thread auth-1
Agentcomm inbox --as reviewer --json # 消費します。アーカイブは読み取り中です/
Agentcomm wait --as reviewer --timeout 30000 # 配信時に 0 で終了、タイムアウト時に 2 で終了
# 共有ワーカーキューパターン (複数のワーカー、1 つのキュー) — git + SQL バックエンド
Agentcomm send work-queue " task-1 " --subject task
Agentcommclaim --queue work-queue --as work-1 # アトミック;空の場合はnull
# トポロジーがそれを必要とする場合の他のストア (プッシュ待機、SQL クレーム):
エクスポート AGENTCOMM_BACKEND=postgres://user:pass@host:5432/agentcomm
Agentcomm wait --as reviewer --timeout 30000 # LISTEN/NOTIFY 経由で ~ms 以内に解決します
send/broadcast 末尾の引数または stdin から本文を読み取ります
省略した場合:
エコー「パイプから」 | Agentcomm はボブを --as アリスとして送信します
人々がそれを使って何を構築するか
エージェントはリポジトリを共有し、リポジトリを通じて会話します。リポジトリはバスです: リポジトリ
権限は ACL であり、すべてのメッセージは監視できるコミットです。
クラウド + ローカル ワーカー フリートがアトミック クレームで 1 つのキューを分割します。
「ビルドのステータスは何ですか?」と尋ねることができる CD パイプライン。デプロイ中。
IoT エッジ エージェント — 「何が見えますか?」に答えるカメラ、気象センサー
アウトバウンド HTTPS のみで 1 つのブロードキャストに湿度を報告します。
1 台のマシン上での Claude Code、Codex、OpenCode のペアリング - 各ネイティブ プラグインは
両方が同じリポジトリ バスを介して通信しながら、独自のガイダンス ファイルを作成します。
すべてイラスト付き

で実行可能なコマンド
ユースケースページ — プラス
セキュリティの話が引き算である理由: ストレージの認証はバスの認証です。
コマンド
何をするのか
初期化
このリポジトリをバス上に置きます。デフォルトで CLAUDE.md を書き込み、 --harness codex|opencode|agents を指定すると AGENTS.md を書き込み、ユーザーを登録し、名簿を表示します。選択したハーネス ファイルをコミットします。
登録する
呼び出し側エージェントを登録/ハートビートします ( --as )。
エージェント
登録されているエージェントを一覧表示します。
<宛先> [本文] に送信します
メッセージを送信します (引数または標準入力からの本文)。
ブロードキャスト[本文]
あなた自身を除くすべての登録エージェントに送信します。
受信箱
未配信のメッセージを消費します。それらを read/ の下にアーカイブします。
覗く
未配信のメッセージを消費せずに表示します。
待ってください
メッセージが到着する ( exit 0 ) かタイムアウトする ( exit 2 ) までブロックします。
主張する
--queue ( git + SQL バックエンド ) から 1 つのメッセージをアトミックにデキューします。
説明する
--backend スキーム、つまり URI からチャネルがどのように切り分けられるか、およびその機能について説明します。静的 — ドライバーをロードしたり、接続したりしません。
チャンネル
--backend ストアにすでに存在するチャネルをリストします (agentcomm キー レイアウトをスキャンします。ドライバーと認証情報が必要です)。
パージ
--older-than より古いアーカイブ ( read/ ) メッセージ、および/または --agents-older-than を過ぎてアイドル状態になった登録を削除します。保留中のメールには一切触れません。デーモンは両方とも自動的に実行されます (デフォルトは 30 日/7 日)。
ログ
チャネルの会話を読み取ります — 保留中 + アーカイブ済み、時間順、非消費型、なし -- 必要に応じて。 --thread 、 --limit 。
慣例
効果的なチーム規約を出力します (組み込みのデフォルト ⊕ .agentcomm.json / .yaml オーバーライド)。静的 — 決して接続しません。
フラグ
旗
意味
--バックエンド <

[切り捨てられた]

## Original Extract

A tiny mailbox for AI agents — one CLI, six backends (GitHub repo, SQLite, S3, GCS, Postgres, local). The repo is the bus. - yonidavidson/agentcomm

GitHub - yonidavidson/agentcomm: A tiny mailbox for AI agents — one CLI, six backends (GitHub repo, SQLite, S3, GCS, Postgres, local). The repo is the bus. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
yonidavidson
/
agentcomm
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
153 Commits 153 Commits .agents/ plugins .agents/ plugins .claude-plugin .claude-plugin .github/ workflows .github/ workflows bin bin commands commands dist dist docs docs hooks hooks plugins/ agentcomm plugins/ agentcomm scripts scripts skills/ agentcomm skills/ agentcomm src src test test .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md docker-compose.test.yml docker-compose.test.yml package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts View all files Repository files navigation
🌐 Website · Use cases · Live demo — an agent conversation that is a git branch · Claude Code plugin · Codex plugin · OpenCode plugin
A tiny mailbox / message bus for AI agents that shell out to one CLI. Agents
register , send , and read their inbox ; a single Backend interface hides
where the messages live. Any git remote is a bus — the git backends need
no driver at all; local runs need zero dependencies; everything else is an
optional, lazy-loaded driver.
If you are an AI agent reading this
Get yourself (and this repo) on the bus — takes under a minute:
If you are Claude Code , ask your user to run
/plugin marketplace add yonidavidson/agentcomm and
/plugin install agentcomm@yonidavidson-plugins . The plugin ships the
CLI, a skill with the conventions, and hooks that register you on the
repo bus automatically at session start. Then run agentcomm init ; it
adds the coordination contract to CLAUDE.md .
If you are Codex , ask your user to run codex plugin marketplace add yonidavidson/agentcomm
and codex plugin add agentcomm@yonidavidson-plugins . The plugin ships the
CLI, coordination skill, and lifecycle hooks. Review and trust its hooks
with /hooks after installation. Then ask Codex to use agentcomm to
initialize the repo; the skill runs init --harness codex and adds the
coordination contract to AGENTS.md .
If you are OpenCode , add the plugin tarball from this repo's latest
release to your
opencode.json — "plugin": ["https://github.com/yonidavidson/agentcomm/releases/download/v0.16.10/agentcomm-opencode-0.16.10.tgz"] .
It puts every session on the repo bus in-process. OpenCode reads AGENTS.md
natively, so agentcomm init --harness opencode (which writes AGENTS.md )
also onboards it — see As an OpenCode plugin .
┌─────────────────────────────────────────────┐
agents ──▶ │ agentcomm CLI (one stable interface) │
│ │ │
│ ▼ │
│ Backend interface ◀── the seam │
│ ├─ GitBackend — ANY git remote is a bus │
│ ├─ GithubBackend — GitHub via token (no ssh)│
│ ├─ LocalBackend — zero-dep default │
│ ├─ SqliteBackend — single box, WAL │
│ ├─ S3Backend — object store │
│ ├─ GCSBackend — object store │
│ └─ PostgresBackend — distributed, push │
└─────────────────────────────────────────────┘
Install
Not on the npm registry (yet) — install straight from GitHub. dist/ is
committed to the repo, so this needs no build step, and the install is
genuinely zero-dependency : you get dist/ and nothing else.
npm install github:yonidavidson/agentcomm
# git+ssh:// / git+https:// / github:// / file:// need NOTHING more
# (Node ≥ 18; the git binary for git+; a token for github://).
# Per-backend drivers, only if you use that backend — the CLI names the
# exact package when one is missing:
npm install better-sqlite3 # sqlite://
npm install @aws-sdk/client-s3 # s3://
npm install @google-cloud/storage # gs://
npm install pg # postgres://
npm install yaml # only for .agentcomm.yaml config files (.json needs nothing)
As a Claude Code plugin
This repo is also a self-hosted Claude Code plugin marketplace — install it
and Claude picks up a skill that knows the CLI's commands, flags, and
backend tradeoffs, and uses them to coordinate with other agents/sessions:
/plugin marketplace add yonidavidson/agentcomm
/plugin install agentcomm@yonidavidson-plugins
No global install or npm registry publish required — the plugin ships a
prebuilt copy of the CLI and the skill runs it directly. In a git repo it
defaults to the repo bus, like everywhere else.
This repo is also a Codex marketplace. Add it and install the plugin from
the marketplace snapshot:
codex plugin marketplace add yonidavidson/agentcomm
codex plugin add agentcomm@yonidavidson-plugins
The plugin bundles the same prebuilt CLI and coordination skill plus Codex
lifecycle hooks for registration, inbox digests, and the stop guard. Codex
requires explicit trust for non-managed hooks: open /hooks , review the
agentcomm definitions, and trust them. Start a new thread after installing
or upgrading so the plugin components are loaded.
Ask Codex directly so its skill uses the bundled CLI:
Use agentcomm to initialize this Codex repo for the team.
As an OpenCode plugin
OpenCode runs on Bun and reads AGENTS.md natively, so
its agents already onboard from this repo's AGENTS.md . The plugin adds the
lifecycle — it registers each session on the bus, briefs it, surfaces unread
mail before the session goes idle, and keeps long turns reachable — by
importing the agentcomm library in-process (no subprocess). Because OpenCode's
session.idle is observe-only, the inbox guard re-prompts the session rather
than blocking it.
Install it from the plugin tarball attached to each
release — OpenCode fetches
the .tgz directly, no clone and no npm registry:
{
"plugin" : [ " https://github.com/yonidavidson/agentcomm/releases/download/v0.16.10/agentcomm-opencode-0.16.10.tgz " ]
}
OpenCode loads the plugin from the tarball's package root via its
exports["./server"] entry (the compiled library ships inside, so there's no
build step and — for the file/git backends — zero runtime dependencies).
Updating. OpenCode caches a plugin by its URL and never re-fetches, so a
"latest" URL would silently pin you to your first install. The URL is versioned
on purpose: bump the version to upgrade. You don't have to watch the
releases page — the plugin checks once a day and, when a newer release exists,
prints an "agentcomm-opencode update available: vX → vY" notice in-session
(like omp / pi do), telling you exactly which version to put in the URL.
Why a tarball and not github:… ? OpenCode installs a remote plugin by
cloning the whole repo, and this monorepo (full CLI + committed dist/ across
its history) is a large, slow clone that OpenCode's installer chokes on. The
release tarball is ~100 kB (dist only, no history), so it installs in
seconds. To develop against a local checkout , point the entry at the repo
directory instead: "plugin": ["/absolute/path/to/agentcomm"] .
# in a git repo: zero config. You're on the repo bus under a session-unique
# alias; init defaults to Claude Code and writes CLAUDE.md.
agentcomm init # → acting as yoni-3f2a · on the bus: git+ssh://…
# Codex uses its own repo guidance file.
agentcomm init --harness codex # → AGENTS.md created
agentcomm agents # who's here: yoni-3f2a · dana-97b1 · ci-bot
agentcomm send ci-bot " hold deploys " --subject status
# named ROLES (addressable, stable) take --as; register warns if the alias
# is live in another session
agentcomm register --as reviewer
agentcomm send reviewer " review src/auth.ts " --subject task --thread auth-1
agentcomm inbox --as reviewer --json # consumes; archives under read/
agentcomm wait --as reviewer --timeout 30000 # exit 0 on delivery, 2 on timeout
# shared-worker-queue pattern (multiple workers, one queue) — git + SQL backends
agentcomm send work-queue " task-1 " --subject task
agentcomm claim --queue work-queue --as worker-1 # atomic; null when empty
# other stores when topology calls for it (push wait, SQL claims):
export AGENTCOMM_BACKEND=postgres://user:pass@host:5432/agentcomm
agentcomm wait --as reviewer --timeout 30000 # resolves within ~ms via LISTEN/NOTIFY
send / broadcast read the body from the trailing argument, or from stdin
if omitted:
echo " from a pipe " | agentcomm send bob --as alice
What people build with it
Agents sharing a repo, talking through it — the repo is the bus: repo
permissions are the ACL, every message is a commit you can watch.
Cloud + local worker fleets splitting one queue with atomic claim .
A CD pipeline you can ask "what's the status of the build?" mid-deploy.
IoT edge agents — a camera answering "what do you see?", weather sensors
reporting humidity to one broadcast — on nothing but outbound HTTPS.
Claude Code, Codex, and OpenCode pairing on one machine — each native plugin uses
its own guidance file while both communicate over the same repo bus.
All illustrated with runnable commands on the
use-cases page — plus
why the security story is subtraction : your storage's auth is the bus's auth.
Command
What it does
init
Put this repo on the bus: writes CLAUDE.md by default, or AGENTS.md with --harness codex|opencode|agents , registers you, and shows the roster. Commit the selected harness file.
register
Register / heartbeat the calling agent ( --as ).
agents
List registered agents.
send <to> [body]
Send a message (body from arg or stdin).
broadcast [body]
Send to every registered agent except yourself.
inbox
Consume undelivered messages; archives them under read/ .
peek
Show undelivered messages without consuming.
wait
Block until a message arrives ( exit 0 ) or timeout ( exit 2 ).
claim
Atomically dequeue one message from --queue ( git + SQL backends ).
describe
Explain the --backend scheme: how channels are carved from the URI, and its capabilities. Static — never loads a driver or connects.
channels
List the channels that already exist on the --backend store (scans for the agentcomm key layout; needs the driver + credentials).
purge
Delete archived ( read/ ) messages older than --older-than , and/or registrations idle past --agents-older-than . Pending mail is never touched. The daemon runs both automatically (30d / 7d defaults).
log
Read a channel's conversation — pending + archived, time-ordered, non-consuming , no --as needed. --thread , --limit .
conventions
Print the effective team conventions (built-in defaults ⊕ .agentcomm.json / .yaml override). Static — never connects.
Flags
Flag
Meaning
--backend <

[truncated]
