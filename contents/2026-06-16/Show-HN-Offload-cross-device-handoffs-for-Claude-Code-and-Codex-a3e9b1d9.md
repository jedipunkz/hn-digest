---
source: "https://github.com/ToxicPine/offloads"
hn_url: "https://news.ycombinator.com/item?id=48555644"
title: "Show HN: Offload, cross-device handoffs for Claude Code and Codex"
article_title: "GitHub - ToxicPine/offloads: /offload for Claude Code and Codex -- Like Cursor's Cloud Agents, but for Codex and Claude Code. · GitHub"
author: "cardellifan"
captured_at: "2026-06-16T16:38:43Z"
capture_tool: "hn-digest"
hn_id: 48555644
score: 3
comments: 0
posted_at: "2026-06-16T14:13:00Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Offload, cross-device handoffs for Claude Code and Codex

- HN: [48555644](https://news.ycombinator.com/item?id=48555644)
- Source: [github.com](https://github.com/ToxicPine/offloads)
- Score: 3
- Comments: 0
- Posted: 2026-06-16T14:13:00Z

## Translation

タイトル: HN を表示: クロード コードとコーデックスのオフロード、クロスデバイス ハンドオフ
記事のタイトル: GitHub - ToxicPine/offloads: Claude Code と Codex の /offload -- Cursor の Cloud Agents に似ていますが、Codex と Claude Code 用です。 · GitHub
説明: Claude Code および Codex 用の /offload -- Cursor の Cloud Agent と似ていますが、Codex および Claude Code 用です。 - ToxicPine/オフロード
HN テキスト: プロンプトの先頭に `/offload` を追加するだけで、別のデバイスの /goal で実行されます (Mac Mini のように、VPS をセットアップすることもできます)。環境キーの安全な移動、サインインしていることの確認、ツールや依存関係の正しいバージョンの固定、開発サーバーを表示するための安全なトンネリングなどの詳細を処理します。Claude Code または Codex のリモート コントロール機能を介して直接介入できます。開発作業のために自分の VPS を維持するのはあまりにも面倒だと感じましたし、Claude Code と Codex のクラウド エージェントも面倒だと感じました。また、他のオプションでは、OpenCode + OpenRouter、同じプロジェクト内で異なるハーネスの使用などはサポートされていません。

記事本文:
GitHub - ToxicPine/offloads: クロード コードとコーデックスの /offload -- Cursor のクラウド エージェントと似ていますが、コーデックスとクロード コード用です。 · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
有毒パイン
/
オフロード
公共
通知
通知設定を変更するにはサインインする必要があります
さらに

l ナビゲーションオプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
25 コミット 25 コミット .agents/ skill/ caveman .agents/ skill/ caveman メディア メディア パッケージ パッケージ .gitattributes .gitattributes .gitignore .gitignore LICENSE.md LICENSE.md README.md README.md flake.lock flake.lock flake.nix flake.nix install-offload install-offload skill-lock.json skill-lock.json skill.nix skill.nix すべてのファイルを表示 リポジトリ ファイルのナビゲーション
offload-demo-25s-v2-youtube-1080p30.mp4
について
/offload を使用すると、コーディング エージェントは別のマシンで長時間実行されるタスクを継続できます。
/offload lead.csv を調べ、請求書に関して最も一致する 500 件を電子メールで送信します
午前 9 時から午後 5 時までの清掃サービス。誰が返信し、何を尋ねたかを追跡します。
そして水曜日に短いレビューの概要と推奨される次のステップを発表します。
エージェントは現在のプロジェクトの状態を別のマシンに送信し、タスクを実行します。
そこにあり、変更を新しいブランチに保存します。使えるマシンはどれでも使えます
access または /offload を使用すると、ジョブをクラウド インスタンスに簡単にディスパッチできます。
Fly 、Cursor Cloud Agents に似ています。
ラップトップがスリープ状態になったり切断されたりしても、実行は継続できます。送信できます
ステータス更新、リモート localhost:<port> dev サーバーを公開
認証されたパブリック URL を使用し、公式のクロード コードまたはコーデックス リモートを使用します。
コントロールを使用してチェックインまたは操作します。
/offload はエージェント スキルです。永続的なローカル ソフトウェアは必要ありません。
カール -fsSL https://raw.githubusercontent.com/ToxicPine/offloads/master/install-offload |バッシュ
bash -s -- の後にオプションを渡します。たとえば、プロンプトをスキップするには --yes です。
--silent は静かにインストールします。
このソフトウェアは、完全なエージェント スキルとともに配布されます。
ドキュメント。スキルをインストールしてから、LLM に質問することをお勧めします。
あるいは、「」を読んでもよいでしょう。

テキストを自分で削除して開始します
これで。
/offload には、Nix パッケージを実行するためのゼロインストール、ルートレスの方法が付属しています。
このスキルは、プロジェクトをマネージャーとして実行できるようにする作業を管理するのに役立ちます。
ニックスフレーク。これにより、リポジトリでリモートのプロジェクト環境を定義できるようになります。
マシンが正しく動作するために必要なものです。
次の部分はリモート ターゲットです。作業用に設計されたコンテナで、
インストールされた適切なツール、ランタイム Nix ビルド、永続的な状態、および
独自の localhost:<port> 上で実行されている開発サーバーを、認証されたものを通じて公開します。
パブリック URL。スキルには、そのコンテナを Fly にデプロイするための手順が含まれています。
残りは、GitHub 認証情報の確認とシードなどの統合の磨きです。
Codex などのツールの git ID、リポジトリの状態、およびコーディング エージェントのログイン状態
クロード・コード。目標はシンプルです。同じプロジェクトを別の場所で実行し、戻ってくることです。
結果は通常のブランチになります。
オプション:Hermes エージェントと Telegram の統合
このマシンにはエルメスが含まれています
デフォルト。オプションで、進行状況の ping のために Telegram 統合を有効にすることができます。
電話による監視や、リモート エージェントに質問するなどの迅速な介入が可能です。
開発サーバーのリンク。 /offload はなくても機能します。
オフロードは主要なエージェント スキルです。リモート コンピューターの準備ができているかどうかを判断します。
必要に応じてセットアップを支援し、オフロードされた実行を開始します。
offloader は、現在の git プロジェクトからリモート コンピューターにコマンドを送信します。
そして、作業を受け取る run ブランチを作成します。
offloader-configurator はリモート コンピュータをチェックし、アカウントをシードします。
GitHub やアシスタントのログインなど、必要な状態を示します。
offloader-container は、
デフォルトの設定。
offloader-transports は、そのコンピュータにアクセスするための小さな一連の方法です。
Fly.io、プレーン SSH、および Tailscale SSH。
ネステールT

リモート コンピューター上の localhost:<port> を開くことができる URL に変換します。
そして、そのポートに対して保護された共有リンクを生成できます。
vusperize は長時間の作業をラップするため、実行中にライブ ステータスの更新を送信できます。
boondoggler は Codex に目標を与え、それをリモート ブランチから動作させます。
コミットして結果をプッシュバックします。
Claude Code と Codex の /offload -- Cursor の Cloud Agent と似ていますが、Codex と Claude Code 用です。
読み込み中にエラーが発生しました。このページをリロードしてください。
1
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

/offload for Claude Code and Codex -- Like Cursor's Cloud Agents, but for Codex and Claude Code. - ToxicPine/offloads

You just prepend `/offload` to a prompt, and it will run on /goal on another device (like a Mac Mini, it can also set-up a VPS for you). It handles details like moving env keys safely, making sure you're signed-into gh, pinning the right versions of any tools or dependencies, tunnelling securely to display your dev servers, etc. You can intervene directly via the Remote Control feature in Claude Code or Codex. I found it way too tedious to maintain my own VPS for dev work, and I found the cloud agent stuff for Claude Code and Codex annoying also. Also, other options doesn't support OpenCode + OpenRouter, using different harnesses within the same project, etc.

GitHub - ToxicPine/offloads: /offload for Claude Code and Codex -- Like Cursor's Cloud Agents, but for Codex and Claude Code. · GitHub
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
ToxicPine
/
offloads
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
25 Commits 25 Commits .agents/ skills/ caveman .agents/ skills/ caveman media media packages packages .gitattributes .gitattributes .gitignore .gitignore LICENSE.md LICENSE.md README.md README.md flake.lock flake.lock flake.nix flake.nix install-offload install-offload skills-lock.json skills-lock.json skills.nix skills.nix View all files Repository files navigation
offload-demo-25s-v2-youtube-1080p30.mp4
About
/offload lets coding agents continue long-running tasks on another machine.
/offload look through leads.csv, email the 500 best matches about our invoice
cleanup service between 9am and 5pm, track who replies and what they ask,
and stop Wednesday with a short review summary and recommended next steps
The agent sends your current project state to another machine, runs the task
there, and saves any changes on a new branch. You can use any machine you can
access, or /offload can easily dispatch the job to a cloud instance on
Fly , similar to Cursor Cloud Agents.
The run can continue even if your laptop sleeps or disconnects. It can send
status updates, expose remote localhost:<port> dev servers through
authenticated public URLs, and use the official Claude Code or Codex remote
controls to check in or steer it.
/offload is an agent skill; no persistent local software is needed.
curl -fsSL https://raw.githubusercontent.com/ToxicPine/offloads/master/install-offload | bash
Pass options after bash -s -- , for example --yes to skip the prompt or
--silent for quiet installs.
This software is distributed with an agent skill, which serves as complete
documentation. I suggest that you install the skill, then ask an LLM about
its usage, etc. Alternatively, you may read the skill text yourself, starting
with this .
/offload ships with a zero-install, rootless way to run the Nix package
manager, and the skill helps manage the work of making your project runnable as
a Nix flake. That lets the repo define the project environment the remote
machine needs in order to work with it correctly.
The next part is the remote target: a container designed for work, with the
right tooling installed, runtime Nix builds, persistent state, and a way to
expose dev servers running on its own localhost:<port> through authenticated
public URLs. The skill includes instructions for deploying that container on Fly.
The rest is integration polish, like checking and seeding GitHub credentials,
git identity, repo state, and coding-agent login state for tools like Codex and
Claude Code. The goal is simple: run the same project somewhere else and return
the result as a normal branch.
Optional: Hermes Agent Integration With Telegram
The machine includes Hermes by
default. You can optionally enable its Telegram integration for progress pings,
phone supervision, and quick interventions like asking the remote agent for a
dev-server link. /offload still works without it.
offload is the main agent skill. It decides whether a remote computer is ready,
helps set one up when needed, and starts the offloaded run.
offloader sends a command from your current git project to the remote computer
and creates the run branch that receives the work.
offloader-configurator checks and seeds the remote computer with the account
state it needs, such as GitHub and assistant login.
offloader-container is the ready-to-run remote computer image used by the
default setup.
offloader-transports is the small set of ways to reach that computer, including
Fly.io, plain SSH, and Tailscale SSH.
nestail turns localhost:<port> on the remote computer into an openable URL,
and can generate protected share links for that port.
vusperize wraps long work so it can send live status updates while it runs.
boondoggler gives Codex a goal, lets it work from the remote branch, then
commits and pushes the result back.
/offload for Claude Code and Codex -- Like Cursor's Cloud Agents, but for Codex and Claude Code.
There was an error while loading. Please reload this page .
1
fork
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
