---
source: "https://github.com/5dive-com/5dive"
hn_url: "https://news.ycombinator.com/item?id=48490081"
title: "Show HN: 5dive – Run a Team of Claude Code Agents from Telegram"
article_title: "GitHub - 5dive-com/5dive: Spawn and manage AI agents that talk to each other. systemd-managed multi-agent runtime on your own VM. · GitHub"
author: "lodar"
captured_at: "2026-06-11T16:30:43Z"
capture_tool: "hn-digest"
hn_id: 48490081
score: 3
comments: 0
posted_at: "2026-06-11T13:28:09Z"
tags:
  - hacker-news
  - translated
---

# Show HN: 5dive – Run a Team of Claude Code Agents from Telegram

- HN: [48490081](https://news.ycombinator.com/item?id=48490081)
- Source: [github.com](https://github.com/5dive-com/5dive)
- Score: 3
- Comments: 0
- Posted: 2026-06-11T13:28:09Z

## Translation

タイトル: 表示 HN: 5dive – Telegram のクロード コード エージェントのチームを実行する
記事のタイトル: GitHub - 5dive-com/5dive: 相互に通信する AI エージェントを生成および管理します。独自の VM 上の systemd 管理のマルチエージェント ランタイム。 · GitHub
説明: 相互に通信する AI エージェントを生成および管理します。独自の VM 上の systemd 管理のマルチエージェント ランタイム。 - 5dive-com/5dive

記事本文:
GitHub - 5dive-com/5dive: 相互に通信する AI エージェントを生成および管理します。独自の VM 上の systemd 管理のマルチエージェント ランタイム。 · GitHub
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
5ダイブコム
/
5ダイブ
公共
通知
通知設定を変更するにはサインインする必要があります
追加

アルナビゲーションオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
274 コミット 274 コミット .github .github docker docker フック フック スキル/ 通知ユーザー スキル/ 通知ユーザー src src systemd systemd チーム テンプレート チーム テンプレート .gitignore .gitignore 5dive 5dive 5dive-agent-start 5dive-agent-start 5dive-refresh-plugins.sh 5dive-refresh-plugins.sh CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md build.sh build.sh install.sh install.sh project-CLAUDE.md Projects-CLAUDE.md telegram-agent-CLAUDE.md telegram-agent-CLAUDE.md すべてのファイルを表示リポジトリ ファイルのナビゲーション
相互に通信する AI エージェントを生成および管理します。
マサチューセッツ工科大学。 5dive.com 上のすべてのエージェントを実行する同じバイナリ。オープンコア分割はありません。マネージド VM を使用した操作をスキップするか、独自の VM を実行します。
#1.インストール
カール -fsSL https://install.5dive.com | sudo bash
# 2. 最初のエージェントを作成します。ウィザードは Telegram にも接続します。
# ボット トークンを貼り付け (BotFather が提供します)、ボットを送信します /start,
# そしてそれは自動的にペアリングされます。コードはありません。
sudo 5dive 初期化
全体をカットせず、新しい箱に入れて、応答するクロードエージェントにインストールします
Telegram の場合 (アイドル時間圧縮):
代わりにスクリプトを作成しますか (CI、プロビジョニング)?非対話型パスには 1 つのパスが必要です
追加のステップ — ボットは最初の DM にペアリング コードを返信します。
sudo 5dive エージェント create my-agent --type=claude --channels=telegram --telegram-token= < トークン >
sudo 5dive エージェント ペア my-agent --code= < ペアリング コード >
仕組み
各エージェントは、公式のエージェント AI CLI セッション ( claude 、 codex 、 antigravity 、 grok など) を systemd サービスとして実行する独自の Linux ユーザーです。複数のエージェントが同じ CLI バイナリとサブスクリプションを共有できます。エージェントは、同じ 5dive CLI、つまりバスを呼び出すことによって相互に連絡します。 Telegra のようなチャンネル

エージェントごとに添付します。
1 つのホスト
┌─────────────────┐
│コーダーライター午後 │
│ (クロード) (コーデックス) (クロード) │
│ │ │ │ │
│ ━━━ 5dive CLI ───┘ │
│ 送信・質問・ログ │
━━━━━━━━━━━━━━━━━━━━━━┘
↕ 電報 / Discord
(エージェントごとに添付)
ブローカーもプロトコルもオーケストレーターもありません。共有ファイルシステム、共有 CLI。
あなたがいなくても機能するチーム。 1 つのホスト上の複数のエージェントが相互に調整します。
彼らにバックログを渡します。定期的なタスクを含む共有タスク キューに加えて、キューにある作業がある場合にのみエージェントをウェイクアップするハートビート。決定は、タップして回答するボタンとして携帯電話に表示されます。
セッションではなくサービスとして実行されます。ターミナルを閉じてもエージェントは生き続けます。 Telegram からいつでもメッセージを送信してください。
すべての主要なエージェント AI CLI。 claude 、 codex 、 antigravity 、 grok 、 hermes 、 openclaw 、 opencode 、すべて 1 つのチームの下にあります。
あなただけのサブスクリプション。自分の Pro/Max 上の公式クロード CLI。仲介者なし、OAuth プロキシなし、人為的ポリシーは安全です。
デフォルトでは安全です。各エージェントは、3 つの分離層のいずれかに属する独自の Linux ユーザーです。エージェントをサンドボックス化すると、ホーム ディレクトリを読み取ることも、ボックスを sudo することもできません。
CLI は OSS サーフェスです。ここにあるすべての動詞、すべてのエージェント、すべてのホストは、すべて /usr/local/bin/5dive から駆動されます。
ssh ではなくクリックしたい場合は、5dive.com がマネージド バージョンです。内部では同じ CLI を使用しますが、VM、強化、更新、ダッシュボードが自動的に実行されます。
種類
モデルファミリー
認証
チャンネル
クロード
Anthropic Claude、または Anthropic と互換性のあるエンドポイント
OAuth / APIキー / --provider
電報、ディスコード
コーデックス
OpenAI コーデックス
OAuth / APIキー
電報
反重力
Google反重力
Google OA

うーん
電報
グロク
xAI グロク
OAuth (xAI) / APIキー
電報
エルメス
サードパーティのマルチプロバイダー ハーネス
OAuth (OpenAI) / APIキー
電報、ディスコード
オープンクロー
サードパーティのマルチプロバイダー ハーネス
OAuth (OpenAI) / APIキー
電報、ディスコード
オープンコード
オープンコード
APIキー
電報
hermes と openclaw は、多くのプロバイダー (OpenRouter、Anthropic、Google、Moonshot、DeepSeek、Z.ai など) にルーティングできるコミュニティによって構築されたハーネスです。 2026 年 4 月 4 日の時点で、Anthropic はサードパーティのハーネスを介した消費者の Claude Pro/Max OAuth のルーティングを許可しなくなりました。この作業には、独自の API キーを持つ公式のクロード タイプを使用します。背景: クロードのために OpenClaw を捨てました → 。
クロード タイプは、独自のキーを使用して、サードパーティの Anthropic 互換エンドポイントに対して公式のクロード コード ハーネスを実行することもできます。
sudo 5dive エージェント作成格安コーダー --type=claude --provider=deepseek --api-key= < キー >
プロバイダー数: deepseek (DeepSeek)、moonshot (Kimi)、zai (GLM)
コマンドの概要
5diveエージェントリスト/作成/開始/停止/再起動/rm
5dive エージェントは <名前> <テキスト> を送信します
5dive エージェントは <名前> <テキスト> を尋ねます [--timeout=120]
5dive エージェントのログ <名前> [--follow]
5dive エージェント設定 <名前> セットモデル=<id> / 努力=<低|中|高|x高|最大>
5dive エージェント <名前> トゥイ
5dive タスクの追加 / ls / 割り当て / 開始 / 完了 / 必要 / 受信トレイ / 回答
5dive ハートビート オン / オフ / ls / 作業をキューに入れているウェイク エージェントの数をチェック
5dive 組織セット / ツリー番号 誰が誰に報告するか
5diveアカウントの追加/ログイン/リスト/表示/使用状況/名前変更/削除
5dive 認証セット / ログイン / ステータス # 下位レベル;アカウントは人の道です
5diveスキルの追加/一覧/削除
5dive ドクター [--修復] [--json]
5dive watch #htop風ライブビュー
5dive up / down / ps / 5dive.yaml 経由で # 宣言型エージェントをエクスポート
5dive Team import <slug> # 1 回の呼び出しでチーム全体のテンプレートをプロビジョニングします
5dive 自己アップデート

# CLI + プラグインを更新してからエージェントを再起動します
フラグの完全なリファレンス: 5dive --help (または 5dive <verb> --help )。 --json を介した任意のコマンドの機械可読出力。
アカウント (共有認証プロファイル)
sudo 5diveアカウント追加作業
sudo 5dive アカウントのログイン作業 --type=claude
sudo 5dive エージェント作成エージェント-a --type=claude --auth-profile=work
sudo 5dive エージェント作成エージェント-b --type=claude --auth-profile=work
アカウントの名前を変更またはローテーションすると、バインドされているすべてのエージェントが自動的に再バインドされます。 5dive アカウントの使用状況は、各アカウントのレート制限のヘッドルームを示します。
ボックス上のエージェントはタスク キューを共有します (SQLite、サーバーなし)。作業をファイルして割り当て、何かする必要がある場合にのみハートビートで担当者を目覚めさせます。定期的なテンプレートは cron スケジュールに従って具体化されます。
5dive タスク追加「夜間の CI 障害のトリアージ」 --assignee=ops --quirting= " 0 7 * * * "
sudo 5dive オペレーションでのハートビート --every=30m
エージェントが人間だけが判断できる何かに到達すると、そのタスクをあなたに任せます。
5dive タスクには DIVE-42 が必要です --type=approval --ask= " 出荷価格設定 v2? " --options= " ship|hold " --recommend=ship
それは、タップして応答するボタンとして Telegram に届きます。いずれかをタップすると、所有エージェントのブロックが解除され、再開します。 5dive タスクの受信箱には人間が待っているすべてのものがリストされ、5dive 組織にはレポート チャートが保存されるので、誰が誰のために働いているかを確認できます。
エージェントごとのボットはオプションです。 Telegram グループ (トピックが有効) で 1 つの共有ボットを指定すると、すべてのエージェントが独自のフォーラム トピックを取得します。
sudo 5dive エージェント チームボット共有 --group= < chat_id > --agents=coder,writer,pm --token= < ボットトークン >
新しいエージェントは独自のトピックで自動的にアタッチします ( --no-team-bot を使用してエージェントごとにオプトアウトします)。 Team-bot Discover はグループ ID を検索し、team-bot インターコムはエージェント間の会話を専用のトピックにミラーリングするので、チームの調整を監視できます。
階層
アクセス
管理者 (デフォルト)
フルホスト

標準
共有読み取り、制限付き書き込み
サンドボックス化された
自分のホームのみ、sudo なし、systemd リソース制限
sudo 5dive エージェント create my-agent --type=claude --isolation=sandboxed
仲介業者はいません
5dive はサーバー上で実行されます。認証トークンはモデルプロバイダーに直接送信され、私たちに送信されることはありません。テレメトリ、エラー報告、使用状況データは箱から出ません。各エージェントは、独自のログインを持つ 1 人の Linux ユーザーです。
長い形式: あなたの認証トークンは私たちに影響を与えません → 。
5dive は、シェル アクセスを使用してエージェントを実行します。標準的な衛生状態が適用されます:
OS にパッチを適用します (無人アップグレード)
ベースライン: devsec.os_hardening、Lynis、fail2ban 。または、チェックリストをスキップします。 5dive.com が対応します。
ドッカー。ホストをインストールせずにタイヤを動かしましょう:
docker build -f docker/Dockerfile -t 5dive 。
docker run -d --name 5dive-demo --privileged 5dive
docker exec -it 5dive-demo bash
オフライン/エアギャップ。 install.sh は $REPO (デフォルトの GitHub raw) から読み取ります。 REPO=file:///path/to/local/tree でオーバーライドし、apt deps をプレインストールします。フェッチされたファイルは、 install.sh の先頭にリストされます。
更新中。 5dive は自動更新しません。コードがいつ変更されるかはユーザーが制御できます。オンデマンドで更新するには:
sudo 5dive 自己更新
これにより、CLI、フック、スキル、およびプラグイン (同じ install.sh --upgrade パス) が更新され、実行中の各エージェントが再起動され、新しいバージョンが実際にロードされます。ライブ エージェントは、再起動するまで古いプラグインをメモリ内に保持します。エージェント AI CLI ( claude 、 codex など) は、独自の自動アップデータを介して自己更新します。再起動すると最新のものがロードされます。スケジュール通りにやりたいですか？ cron にドロップします。
0 4 * * * /usr/local/bin/5dive 自己更新 >/dev/null 2>&1
文脈の腐敗。長時間のセッションは劣化します。上記の毎日の自己更新によってエージェントも再起動され、それぞれに新しいセッションが提供されます。クロード ランタイム エージェントは、再起動後もプロジェクト メモリを ~/.claude/projects/<dir>/memory/ の下に保持します。セッションのリセット、わかっています

エッジは残ります。
すでに Claude Code / Codex / Antigravity / Grok / opencode を使用している場合は、このプロンプトを貼り付けてください。エージェントは 5dive をインストールし、スキルを学習し、チャットを通じてエージェントの管理を続けます。
この Linux ホストに 5dive をインストールして、5dive エージェントの管理に使用できるようにします。
1. インストーラーを実行します (冪等、再実行しても安全です)。
カール -fsSL https://install.5dive.com | sudo bash
2. 確認します: `5dive --version` は「5dive 0.1.x」を出力します。
3. 5dive-cli スキルをインストールします。 <runtime> を次のいずれかに置き換えます。
クロード コード、コーデックス、反重力、grok、hermes エージェント、openclaw、opencode:
npx -y スキル追加 https://github.com/5dive-com/skills --skill 5dive-cli --agent <ランタイム> --yes
4. スキルをロードするために再起動するよう指示し、最初に作成するエージェントを尋ねます。
SSH 経由でリモート VM にインストールしますか?同じプロンプトで、インストール行の先頭に ssh -t <user@host> を付けます。リモートではなく、SSH を発行しているラップトップにスキルをインストールします。 TTY が必要なものには ssh -t を使用します (例: 5dive エージェント認証ログイン )。
すべてのコマンドは --json を受け入れます。出力は成功した場合は {ok:true,data:...}、失敗した場合は {ok:false,error:{code,class,message}} です。終了コードが error.code と一致するため、シェル パイプラインは解析せずに分岐します。進行状況ラインは標準エラー出力に残ります。 stdout は常に有効な JSON です。
{ "ok" : true , "data" : [ { "name" : " main " 、 "type" : " クロード " 、 "active" : " active " } ] }
{ "ok" : false 、 "error" : { "code" : 4 、 "class" : " not_found " 、 "message" : " 'foo' という名前のエージェントはありません "

[切り捨てられた]

## Original Extract

Spawn and manage AI agents that talk to each other. systemd-managed multi-agent runtime on your own VM. - 5dive-com/5dive

GitHub - 5dive-com/5dive: Spawn and manage AI agents that talk to each other. systemd-managed multi-agent runtime on your own VM. · GitHub
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
5dive-com
/
5dive
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
274 Commits 274 Commits .github .github docker docker hooks hooks skills/ notify-user skills/ notify-user src src systemd systemd team-templates team-templates .gitignore .gitignore 5dive 5dive 5dive-agent-start 5dive-agent-start 5dive-refresh-plugins.sh 5dive-refresh-plugins.sh CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md build.sh build.sh install.sh install.sh projects-CLAUDE.md projects-CLAUDE.md telegram-agent-CLAUDE.md telegram-agent-CLAUDE.md View all files Repository files navigation
spawn and manage AI agents that talk to each other.
MIT. The same binary that runs every agent on 5dive.com . No open-core split. Skip the ops with the managed VM, or run your own.
# 1. install
curl -fsSL https://install.5dive.com | sudo bash
# 2. create your first agent — the wizard wires Telegram too:
# paste a bot token (BotFather gives you one), send the bot /start,
# and it pairs itself. No codes.
sudo 5dive init
The whole thing, uncut, on a fresh box — install to a Claude agent answering
on Telegram (idle time compressed):
Scripting it instead (CI, provisioning)? The non-interactive path needs one
extra step — the bot replies to your first DM with a pairing code:
sudo 5dive agent create my-agent --type=claude --channels=telegram --telegram-token= < token >
sudo 5dive agent pair my-agent --code= < pairing-code >
How it works
Each agent is its own Linux user running an official agentic AI CLI session ( claude , codex , antigravity , grok , …) as a systemd service. Multiple agents can share the same CLI binary and subscription. Agents reach each other by invoking the same 5dive CLI — that is the bus. Channels like Telegram attach per agent.
one host
┌──────────────────────────────────┐
│ coder writer pm │
│ (claude) (codex) (claude) │
│ │ │ │ │
│ └──── 5dive CLI ─────┘ │
│ send · ask · logs │
└──────────────────────────────────┘
↕ Telegram / Discord
(attach per agent)
No broker, no protocol, no orchestrator. Shared filesystem, shared CLI.
A team that works without you. Multiple agents on one host, coordinating with each other.
Hand them a backlog. A shared task queue with recurring tasks, plus a heartbeat that wakes an agent only when it has queued work. Decisions land on your phone as tap-to-answer buttons.
Runs as a service, not a session. Your agents stay alive when you close the terminal. Message them from Telegram any time.
Every major agentic AI CLI. claude , codex , antigravity , grok , hermes , openclaw , opencode , all under one team.
A subscription that's yours. Official claude CLI on your own Pro/Max. No middleman, no OAuth proxy, Anthropic-policy safe.
Safe by default. Each agent is its own Linux user under one of three isolation tiers. Sandbox an agent and it can't read your home dir or sudo your box.
The CLI is the OSS surface. Every verb here, every agent, every host, all driven from /usr/local/bin/5dive .
If you'd rather click than ssh , 5dive.com is the managed version: same CLI under the hood, but the VM, hardening, updates, and dashboard are run for you.
Type
Model family
Auth
Channels
claude
Anthropic Claude, or any Anthropic-compatible endpoint
OAuth / API key / --provider
Telegram, Discord
codex
OpenAI Codex
OAuth / API key
Telegram
antigravity
Google Antigravity
Google OAuth
Telegram
grok
xAI Grok
OAuth (xAI) / API key
Telegram
hermes
third-party multi-provider harness
OAuth (OpenAI) / API key
Telegram, Discord
openclaw
third-party multi-provider harness
OAuth (OpenAI) / API key
Telegram, Discord
opencode
OpenCode
API key
Telegram
hermes and openclaw are community-built harnesses that can route to many providers (OpenRouter, Anthropic, Google, Moonshot, DeepSeek, Z.ai, etc.). As of April 4, 2026, Anthropic no longer permits routing consumer Claude Pro/Max OAuth through third-party harnesses. For that work, use the official claude type with your own API key. Background: We Ditched OpenClaw for Claude → .
The claude type can also run the official Claude Code harness against a third-party Anthropic-compatible endpoint, bring your own key:
sudo 5dive agent create cheap-coder --type=claude --provider=deepseek --api-key= < key >
# providers: deepseek (DeepSeek), moonshot (Kimi), zai (GLM)
Commands at a glance
5dive agent list / create / start / stop / restart / rm
5dive agent send <name> <text>
5dive agent ask <name> <text> [--timeout=120]
5dive agent logs <name> [--follow]
5dive agent config <name> set model=<id> / effort=<low|medium|high|xhigh|max>
5dive agent <name> tui
5dive task add / ls / assign / start / done / need / inbox / answer
5dive heartbeat on / off / ls / tick # wake agents that have queued work
5dive org set / tree # who reports to whom
5dive account add / login / list / show / usage / rename / remove
5dive auth set / login / status # lower-level; account is the human path
5dive skill add / list / remove
5dive doctor [--repair] [--json]
5dive watch # htop-style live view
5dive up / down / ps / export # declarative agents via 5dive.yaml
5dive team import <slug> # provision a whole team template in one call
5dive self-update # update CLI + plugins, then restart agents
Full flag reference: 5dive --help (or 5dive <verb> --help ). Machine-readable output on any command via --json .
Accounts (shared auth profiles)
sudo 5dive account add work
sudo 5dive account login work --type=claude
sudo 5dive agent create agent-a --type=claude --auth-profile=work
sudo 5dive agent create agent-b --type=claude --auth-profile=work
Rename or rotate the account, every bound agent rebinds automatically. 5dive account usage shows each account's rate-limit headroom.
Agents on a box share a task queue (sqlite, no server). File work, assign it, and let the heartbeat wake the assignee only when there's something to do. Recurring templates materialize on a cron schedule:
5dive task add " triage overnight CI failures " --assignee=ops --recurring= " 0 7 * * * "
sudo 5dive heartbeat on ops --every=30m
When an agent hits something only a human can decide, it parks the task on you:
5dive task need DIVE-42 --type=approval --ask= " Ship pricing v2? " --options= " ship|hold " --recommend=ship
That arrives on your Telegram as tap-to-answer buttons. Tap one, and the owning agent is unblocked and resumes. 5dive task inbox lists everything waiting on a human, and 5dive org keeps a reporting chart so you can see who works for whom.
Per-agent bots are optional. Point one shared bot at a Telegram group (topics enabled) and every agent gets its own forum topic:
sudo 5dive agent team-bot shared --group= < chat_id > --agents=coder,writer,pm --token= < bot-token >
New agents auto-attach with their own topic (opt out per agent with --no-team-bot ). team-bot discover finds the group id for you, and team-bot intercom mirrors inter-agent chatter into a dedicated topic so you can watch the team coordinate.
Tier
Access
admin (default)
full host
standard
shared read, limited write
sandboxed
own home only, no sudo, systemd resource limits
sudo 5dive agent create my-agent --type=claude --isolation=sandboxed
No middlemen
5dive runs on your server. Auth tokens go to model providers directly, never to us. No telemetry, no error reporting, no usage data leaves the box. Each agent is one Linux user with its own login.
Long form: your auth tokens don't touch us → .
5dive runs agents with shell access. Standard hygiene applies:
patch the OS ( unattended-upgrades )
Baselines: devsec.os_hardening · Lynis · fail2ban . Or skip the checklist; 5dive.com handles it.
Docker . Kick the tires without a host install:
docker build -f docker/Dockerfile -t 5dive .
docker run -d --name 5dive-demo --privileged 5dive
docker exec -it 5dive-demo bash
Offline / air-gapped. install.sh reads from $REPO (default GitHub raw). Override with REPO=file:///path/to/local/tree and pre-install apt deps. The fetched files are listed at the top of install.sh .
Updating. 5dive doesn't auto-update — you stay in control of when code changes land. To update on demand:
sudo 5dive self-update
This refreshes the CLI, hooks, skills, and plugins (the same install.sh --upgrade path), then restarts each running agent so the new versions actually load — a live agent keeps its old plugin in memory until it restarts. The agent AI CLIs ( claude , codex , …) self-update via their own autoupdaters; the restart loads the latest. Want it on a schedule? Drop it in cron:
0 4 * * * /usr/local/bin/5dive self-update >/dev/null 2>&1
Context rot. Long sessions degrade — the daily self-update above also restarts agents, giving each a fresh session. Claude-runtime agents keep project memory under ~/.claude/projects/<dir>/memory/ across restarts. Session resets, knowledge stays.
If you already use Claude Code / Codex / Antigravity / Grok / opencode, paste this prompt. Your agent installs 5dive, learns the skill, then keeps managing agents through chat:
Install 5dive on this Linux host so I can use you to manage 5dive agents.
1. Run the installer (idempotent, safe to rerun):
curl -fsSL https://install.5dive.com | sudo bash
2. Confirm: `5dive --version` prints "5dive 0.1.x".
3. Install the 5dive-cli skill. Replace <runtime> with one of
claude-code, codex, antigravity, grok, hermes-agent, openclaw, opencode:
npx -y skills add https://github.com/5dive-com/skills --skill 5dive-cli --agent <runtime> --yes
4. Tell me to restart so the skill loads, then ask which agent to create first.
Installing onto a remote VM over SSH? Same prompt, prefix the install line with ssh -t <user@host> . Install the skill on the laptop where you're issuing ssh from, not the remote. Use ssh -t for anything needing a TTY (e.g. 5dive agent auth login ).
Every command accepts --json . Output is {ok:true,data:...} on success or {ok:false,error:{code,class,message}} on failure. Exit code matches error.code so shell pipelines branch without parsing. Progress lines stay on stderr; stdout is always valid JSON.
{ "ok" : true , "data" : [ { "name" : " main " , "type" : " claude " , "active" : " active " } ] }
{ "ok" : false , "error" : { "code" : 4 , "class" : " not_found " , "message" : " no agent named 'foo' "

[truncated]
