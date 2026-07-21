---
source: "https://github.com/mohit17mor/SquadAI"
hn_url: "https://news.ycombinator.com/item?id=48993046"
title: "Show HN: SquadAI is the Kubernetes-like control plane for Codex agents"
article_title: "GitHub - mohit17mor/SquadAI · GitHub"
author: "mohit17mor"
captured_at: "2026-07-21T14:55:55Z"
capture_tool: "hn-digest"
hn_id: 48993046
score: 2
comments: 1
posted_at: "2026-07-21T14:48:04Z"
tags:
  - hacker-news
  - translated
---

# Show HN: SquadAI is the Kubernetes-like control plane for Codex agents

- HN: [48993046](https://news.ycombinator.com/item?id=48993046)
- Source: [github.com](https://github.com/mohit17mor/SquadAI)
- Score: 2
- Comments: 1
- Posted: 2026-07-21T14:48:04Z

## Translation

タイトル: Show HN: SquadAI は Codex エージェント用の Kubernetes のようなコントロール プレーンです
記事タイトル: GitHub - mohit17mor/SquadAI · GitHub
説明: GitHub でアカウントを作成して、mohit17mor/SquadAI の開発に貢献します。

記事本文:
GitHub - mohit17mor/SquadAI · GitHub
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
コードの品質 マージ時に品質を強制する
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
モヒット17モル
/
SquadAI
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード

「その他のアクション」メニューを開く フォルダーとファイル
105 コミット 105 コミット .github/ workflows .github/ workflows codex-agent-manager codex-agent-manager codex-control codex-control docs/assets docs/assets .gitattributes .gitattributes .gitignore .gitignore ライセンス ライセンス README.md README.md すべてのファイルを表示リポジトリ ファイルのナビゲーション
SquadAI は、作業がすでに存在するマシン上で、あらゆるイベントを適切な Codex タスクに変換する Codex エージェント用の Kubernetes のようなコントロール プレーンです。
SquadAI は、次のような場合でも、作業を 1 か所で管理し、Codex エージェントに送信できるようにします。
これらのエージェントは別のマシンで実行されます。あなたのプロジェクト、ツール、スキル、
資格情報は、すでにセットアップされているマシン上に残ります。 SquadAIより
ダッシュボードでは、エージェントと直接話すことも、何かがあったときにエージェントに作業を送信することもできます
Webhook、モニター、Telegram メッセージなどの別のツールで発生します。
最速のセットアップでは、コントロール プレーンとエージェントが 1 台のマシン上で実行されます。必要です
Git、Node.js 22.13 以降、およびサインインできる ChatGPT アカウント
コーデックス。
powershell - ExecutionPolicy ByPass - c " irm https://chatgpt.com/codex/install.ps1 | iex "
コーデックスログイン
git clone https://github.com/mohit17mor/SquadAI.git
Set-Location SquadAI\codex - コントロール
npmci
npm ビルドを実行する
Set-Location ..\codex - エージェント - マネージャー
npmci
npm ビルドを実行する
npm start -- -- 埋め込みモード -- ホスト 127.0。 0.1 -- ポート 4317
macOS
ターミナルを開いて次を実行します。
brew install --cask codex
コーデックスログイン
git クローン https://github.com/mohit17mor/SquadAI.git
cd SquadAI/コーデックスコントロール
npmci
npm ビルドを実行する
cd ../codex-agent-manager
npmci
npm ビルドを実行する
npm start -- --modeembedded --host 127.0.0.1 --port 4317
Linux
ターミナルを開いて次を実行します。
カール -fsSL https://chatgpt.com/codex/install.sh |しー
コーデックスログイン
git クローン https://github.com/mohit17mor/SquadAI.git
cd SquadAI/コーデックス-co

コントロール
npmci
npm ビルドを実行する
cd ../codex-agent-manager
npmci
npm ビルドを実行する
npm start -- --modeembedded --host 127.0.0.1 --port 4317
http://127.0.0.1:4317 を開き、エージェントを作成し、選択します
作業ディレクトリを開き、タスクを送信します。 SquadAI はエージェントとそのエージェントを保存します
会話はローカルで行われるため、再起動後も引き続き使用できます。
[エージェントの作成] (またはトポロジ ビューでエージェントの追加) をクリックします。
プロジェクトが存在するマシンを選択し、その作業ディレクトリを選択します。
指示を追加し、エージェントを作成します。
エージェントを選択し、会話でメッセージを送信します。例:
このリポジトリを確認して、最も危険な 3 つの領域を教えてください。
イベント駆動型の作業をテストするには、my-coder をエージェント ID (短い、
エージェント用に作成された小文字の名前。たとえば、リポジトリコーダーは次のようになります。
repository-coder ) を実行し、コントロール プレーン マシンの端末からこれを送信します。
カール http://127.0.0.1:4317/api/sensor-events \
-H ' コンテンツタイプ: application/json ' \
-d ' {
"ソース": "クイックスタート",
"タイプ": "タスク.リクエスト",
"body": "現在のプロジェクトを検査し、最も重要な次のステップを報告します。",
"targetAgentId": "私のコーダー",
"実行ポリシー": "再利用"
} '
タスクは作業キューに表示され、そのエージェントのマシン上で実行されます。
Telegram グループにも接続したいですか?から始める
npm start -- --modeembedded --telegram-token YOUR_CONTROL_BOT_TOKEN 、その後
Telegram Group Control に従ってください。
SquadAI は codex app-server 上に構築されています。これは、
コントロール プレーン全体にわたる Codex セッションの作成、再開、観察、管理
機械。
プロジェクト全体は、GPT-5.6 Sol および GPT-5.6 Terra を使用して Codex で作成されました。
他のモデルは使用されませんでした。 Sol は主要なアーキテクチャの議論に使用されました
そしてコアの分散コントロールプレーン部分。 Terra はより少ない実装に貢献しました
複雑な機能とシャープなアイデア

以前と同様に、それらは要件になりました。
Codex は UI の形成にも役立ちました。視覚的な指示が生成され、そのうちの 1 つが選択されました。
を参照として使用し、インターフェイスはアノテーションベースのポイントで洗練されました。
フィードバック。プロジェクト全体を通じて、GPT-5.6 モデルは特に次の用途に役立ちました。
大まかな要件を理解し、それを具体的な作業に落とし込む
まだ特定されていないエッジケースを積極的に特定します。
コーディング エージェントは、人が 1 つの会話を開始した時点ですでに効果を発揮しており、
それに 1 つのタスクを与えます。複数のエージェントが必要な場合に粗さが発生し、繰り返し発生します。
作業、並列タスク、または長時間実行されるワークフロー:
会話を見つけたり監視したりすることが困難になります。
作業は、あるツールまたはエージェントから別のツールまたはエージェントに手動でコピーする必要があります。
あなたが離れている間は、何も新しい仕事を聞いていません。
1 つのリポジトリを対象とした複数のタスクには、分離されたワークスペースが必要です。
承認、失敗、モデルの変更、未完了の作業には、目に見える 1 つのホームが必要です。
エージェントは、別のラップトップ、ワークステーション、または VM 上で実行する必要がある場合があります。
SquadAI は、不足している運用レイヤーを提供します。コーデックスは今後も責任を負います
推論、コーディング、ツール、MCP サーバー、スキル、プラグイン、サンドボックス。 SquadAI
エージェントを組織し、それを回避する責任があります。
1 つのコマンド センター: エージェント、ライブ ステータス、会話、解説、
ツールのアクティビティ、承認、イベント、キューに入れられた作業を 1 つのブラウザー UI で実行できます。
永続エージェント: すべてのタスクを開始するのではなく、Codex スレッドを再開します。
空虚な会話から。
イベント駆動型の作業: 課題トラッカー、Webhook、モニター、
スケジューラ、または HTTP エンドポイントを呼び出すことができるシステム。
再利用可能な実行または分離された実行: ストリームに対して 1 つの長時間実行エージェントを再利用します。
または、タスクごとに個別のエージェント インスタンスを自動的に作成します。
人間による制御: [承認を求める]、[私の代わりに承認]、または [完全] を選択します。

アクセス、
会話からの承認リクエストに答えます。
リポジトリの分離: Git リポジトリは管理されたワークツリーを使用するため、並列化されます。
エージェント タスクは同じチェックアウトを編集しません。
リモート ランナー: エージェントの実行中、コントロール プレーンを 1 台のマシン上に維持します。
リポジトリ、認証情報、スキル、プラグイン、MCP サーバーが存在する場所。
簡単なマシン登録: Windows、macOS、または Linux ランナーを
プライベート Tailscale 接続を介した 1 つの期限切れコマンドを含む UI。
ライブ ランナー インベントリ: すべてのコントロール プレーンとリモート マシンを確認します。
接続ステータス、割り当てられたエージェント、アクティブな作業、および最後のハートビート。
Telegram チーム チャット: 選択したエージェントに独自の Telegram ボット、タグを付与します。
彼らを 1 つのグループにまとめ、作業、承認、最終的な要約を 1 つのグループで受け取ります。
同じ会話。
共有スキル ライブラリ: 1 人のランナーからユーザー スキルをインポートしてインストールします
一時的なエージェントを作成したり、モデル トークンを消費したりせずに、別のエージェントで実行できます。
アップグレードの認識: 互換性のない固定モデル設定を検出し、リクエストします。
説明もなく失敗を繰り返すのではなく、移行を決定する必要があります。
人 / Webhook / モニター / Telegram
|
v
+----------------------+
| SquadAI コントロール プレーン |
| UI、API、SQLite、仕事 |
|キュー、承認 |
+----------+----------+
|
コマンドとイベント
|
+-----+-----+
| |
vv
ローカルランナー リモートランナー
| |
vv
各ランナー マシン上の codex アプリサーバー
|
v
リポジトリ、Git ワークツリー、ツール、スキル、プラグイン、MCP サーバー
コントロール プレーンは調整状態を保存し、UI を表示します。ランナー
独自のマシン上で Codex セッションを実行し、コントロールに外部接続します
飛行機。これは、ソース コードとマシンローカル ツールをコピーする必要がないことを意味します。
コントロールプレーンホストに送信します。
エージェントは再利用可能な構成です: 名前、手順、作業ディレクトリ、
モデル設定、権限、選択済み

スキルもランナーも。そのコーデックススレッドは
遅延して作成され、後の会話のために永続化されます。
タスク指向の作業の場合、SquadAI はベースから分離されたインスタンスを作成できます。
エージェントが自動的に起動します。各インスタンスは独自の会話を受け取り、Git の場合は
リポジトリ、独自のワークツリーとブランチ。デフォルトでは 3 つのアクティブが許可されます
ベース エージェントごとにインスタンスと 5 つの未解決のインスタンス。ベースエージェントを選択してください
トポロジを選択し、詳細オプションを開いてその両方の制限を設定します
エージェント。未解決の制限は常にアクティブな制限以上である必要があります。
制限;制限が引き下げられても、既存の作業が停止されることはありません。
イベントは外部信号です。エージェントを直接ターゲットにするか、待機する可能性があります。
割り当て/ルーティング。一度受け入れられると、SquadAI が使用できる耐久性のある作業アイテムになります。
ターゲットが利用可能なときにディスパッチされます。
コントロール プレーンは可視性と調整を所有します。ランナーが実行を所有します。
これらは、1 台のコンピューター上の 1 つのプロセスで実行することも、別個のマシン上で実行することもできます。
別のマシンを追加する (推奨)
これは、別の Windows、macOS、または Linux でエージェントを実行する最も簡単な方法です
機械。コントロール プレーンはメイン マシン上に残ります。ソースコード、コーデックス、
資格情報、MCP サーバー、およびローカル ツールはランナー マシン上に残ります。
Tailscale、Node.js、Codex、および SquadAI を両方のマシンにインストールし、署名します
同じ Tailscale ネットワークに接続します。
メイン マシンでコントロール プレーンを起動します。
npm start -- --mode control --host 127.0.0.1 --port 4317
SquadAI で [トポロジ] を開き、 [ランナーの追加] を選択します。
[登録コマンドの生成] を選択します。 SquadAI は、たとえ
コマンドは PATH 上にないため、プライベート アドレスを作成し、それを提供します
コピーするコマンド。
新しいマシンでそのコマンドを実行します。ランナーを登録し、保存します。
~/.squadai/runner.json 内のランナー固有の認証情報を取得し、接続します。
初めてのT

ailscale Serve が使用されるため、ブラウザーの承認が必要な場合があります。使用する
SquadAI によって表示されたリンクを一度承認してから、登録コマンドを生成します
またまた。登録コマンドは 10 分後に期限切れになり、使用できるのは 1 回だけです。
ランナー マシン上で、後で次のように再接続します。
分隊ランナースタート
最後に記録された状態を確認するには:
分隊ランナーのステータス
v1 にはネイティブ インストーラーやバックグラウンド サービスは必要ありません。ランナー
これは、作業を実行する必要があるマシン上で開始するプロセスにすぎません。
SquadAI は 3 つの単純なプリセットを公開しています。
権限はエージェント設定またはチャットコンポーザから変更できます。モデルと
推論の変更は会話を破棄せずに後続のターンに適用されます
スレッド。
エージェントが Git リポジトリを指定すると、SquadAI は管理対象ワークツリーを準備します
ユーザーの Codex データ ディレクトリの下にあります。インスタンス化されたタスクは個別に受け取ります
元のリポジトリ ブランチに基づくブランチとワークツリー。これにより、
1 つの作業ツリーを共有せずに、同じリポジトリを変更する並列タスク。
エージェントの会話から VS Code で開くを使用して、そのエージェントの
チェックアウトして直接差分します。ワークツリーは、次の場合に自動的に削除されません。
コミットされていない変更が含まれています。
あらゆるモニター、スケジューラー、WebH

[切り捨てられた]

## Original Extract

Contribute to mohit17mor/SquadAI development by creating an account on GitHub.

GitHub - mohit17mor/SquadAI · GitHub
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
Code Quality Enforce quality at merge
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
mohit17mor
/
SquadAI
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
105 Commits 105 Commits .github/ workflows .github/ workflows codex-agent-manager codex-agent-manager codex-control codex-control docs/ assets docs/ assets .gitattributes .gitattributes .gitignore .gitignore LICENSE LICENSE README.md README.md View all files Repository files navigation
SquadAI is the Kubernetes-like control plane for Codex agents turning every event into the right Codex task, on the machine where the work already lives!
SquadAI gives you one place to manage and send work to Codex agents, even when
those agents run on different machines. Your projects, tools, skills, and
credentials stay on the machine where they are already set up. From the SquadAI
dashboard, you can talk to an agent directly or send it work when something
happens in another tool, such as a webhook, monitor, or Telegram message.
The fastest setup runs the control plane and agents on one machine. You need
Git, Node.js 22.13 or newer, and a ChatGPT account that can sign in to
Codex.
powershell - ExecutionPolicy ByPass - c " irm https://chatgpt.com/codex/install.ps1 | iex "
codex login
git clone https: // github.com / mohit17mor / SquadAI.git
Set-Location SquadAI\codex - control
npm ci
npm run build
Set-Location ..\codex - agent - manager
npm ci
npm run build
npm start -- -- mode embedded -- host 127.0 . 0.1 -- port 4317
macOS
Open Terminal and run:
brew install --cask codex
codex login
git clone https://github.com/mohit17mor/SquadAI.git
cd SquadAI/codex-control
npm ci
npm run build
cd ../codex-agent-manager
npm ci
npm run build
npm start -- --mode embedded --host 127.0.0.1 --port 4317
Linux
Open Terminal and run:
curl -fsSL https://chatgpt.com/codex/install.sh | sh
codex login
git clone https://github.com/mohit17mor/SquadAI.git
cd SquadAI/codex-control
npm ci
npm run build
cd ../codex-agent-manager
npm ci
npm run build
npm start -- --mode embedded --host 127.0.0.1 --port 4317
Open http://127.0.0.1:4317 , create an agent, choose
its working directory, and send it a task. SquadAI stores your agents and their
conversations locally, so they remain available after a restart.
Click Create Agent (or Add agent in the topology view).
Choose the machine where the project lives, select its working directory,
add instructions, and create the agent.
Select the agent and send it a message in its conversation, for example:
Review this repository and tell me the three riskiest areas.
To test event-driven work, replace my-coder with your agent ID (the short,
lowercase name created for the agent; for example, Repository Coder becomes
repository-coder ) and send this from a terminal on the control-plane machine:
curl http://127.0.0.1:4317/api/sensor-events \
-H ' content-type: application/json ' \
-d ' {
"source": "quick-start",
"type": "task.requested",
"body": "Inspect the current project and report the most important next step.",
"targetAgentId": "my-coder",
"executionPolicy": "reuse"
} '
The task will appear in the work queue and run on that agent's machine.
Want your Telegram group connected too? Start with
npm start -- --mode embedded --telegram-token YOUR_CONTROL_BOT_TOKEN , then
follow Telegram Group Control .
SquadAI is built on codex app-server , which is the backbone that lets the
control plane create, resume, observe, and manage Codex sessions across
machines.
The entire project was written with Codex using GPT-5.6 Sol and GPT-5.6 Terra;
no other model was used. Sol was used for the major architecture discussions
and the core distributed control-plane pieces. Terra helped implement less
complex features and sharpen ideas before they became requirements.
Codex also helped shape the UI: it generated visual directions, one was chosen
as the reference, and the interface was refined with annotation-based pointed
feedback. Across the project, the GPT-5.6 models were especially useful for
understanding high-level requirements, turning them into concrete work, and
proactively identifying edge cases that had not been specified yet.
Coding agents are already effective when a person opens one conversation and
gives it one task. The roughness appears when you need several agents, recurring
work, parallel tasks, or long-running workflows:
conversations become difficult to find and supervise;
work must be copied manually from one tool or agent to another;
nothing is listening for new work while you are away;
multiple tasks aimed at one repository need isolated workspaces;
approvals, failures, model changes, and unfinished work need one visible home;
agents may need to run on different laptops, workstations, or VMs.
SquadAI provides the missing operational layer. Codex remains responsible for
reasoning, coding, tools, MCP servers, skills, plugins, and sandboxing. SquadAI
is responsible for organizing agents and work around it.
One command center: See agents, live status, conversations, commentary,
tool activity, approvals, events, and queued work in one browser UI.
Persistent agents: Resume Codex threads instead of starting every task
from an empty conversation.
Event-driven work: Send work from issue trackers, webhooks, monitors,
schedulers, or any system that can call an HTTP endpoint.
Reusable or isolated execution: Reuse one long-running agent for a stream
of events, or automatically create a separate agent instance for every task.
Human control: Choose Ask for approval, Approve for me, or Full access,
and answer approval requests from the conversation.
Repository isolation: Git repositories use managed worktrees so parallel
agent tasks do not edit the same checkout.
Remote runners: Keep the control plane on one machine while agents run
where the repositories, credentials, skills, plugins, and MCP servers exist.
Easy machine enrollment: Add a Windows, macOS, or Linux runner from the
UI with one expiring command over a private Tailscale connection.
Live runner inventory: See every control-plane and remote machine,
connection status, assigned agents, active work, and last heartbeat.
Telegram team chat: Give selected agents their own Telegram bots, tag
them in one group, and receive their work, approvals, and final summaries in
the same conversation.
Shared skill library: Import a user skill from one runner and install it
on another without creating a temporary agent or spending model tokens.
Upgrade awareness: Detect incompatible pinned model settings and request
a migration decision instead of repeatedly failing without explanation.
People / webhooks / monitors / Telegram
|
v
+-----------------------+
| SquadAI control plane |
| UI, API, SQLite, work |
| queue, approvals |
+-----------+-----------+
|
commands and events
|
+------+------+
| |
v v
Local runner Remote runner(s)
| |
v v
codex app-server on each runner machine
|
v
Repositories, Git worktrees, tools, skills, plugins, and MCP servers
The control plane stores coordination state and presents the UI. A runner
executes Codex sessions on its own machine and connects outward to the control
plane. This means source code and machine-local tools do not have to be copied
to the control-plane host.
An agent is a reusable configuration: name, instructions, working directory,
model settings, permissions, selected skills, and runner. Its Codex thread is
created lazily and persisted for later conversations.
For task-oriented work, SquadAI can create an isolated instance from a base
agent automatically. Each instance receives its own conversation and, for Git
repositories, its own worktree and branch. The defaults allow three active
instances and five unresolved instances per base agent. Select a base agent in
the topology and open Advanced options to configure both limits for that
agent. The unresolved limit must always be equal to or greater than the active
limit; existing work is never stopped when a limit is lowered.
An event is an external signal. It may directly target an agent or wait for
assignment/routing. Once accepted, it becomes a durable work item that SquadAI
dispatches when the target is available.
The control plane owns visibility and coordination. The runner owns execution.
They can run in one process on one computer or on separate machines.
Add Another Machine (Recommended)
This is the simplest way to run agents on another Windows, macOS, or Linux
machine. The control plane remains on your main machine; source code, Codex,
credentials, MCP servers, and local tools remain on the runner machine.
Install Tailscale, Node.js, Codex, and SquadAI on both machines, then sign
in to the same Tailscale network.
Start the control plane on your main machine:
npm start -- --mode control --host 127.0.0.1 --port 4317
In SquadAI, open Topology and choose Add runner .
Select Generate enrollment command . SquadAI finds Tailscale even if its
command is not on PATH , creates a private address, and gives you one
command to copy.
Run that command on the new machine. It enrolls the runner, saves its
runner-specific credential in ~/.squadai/runner.json , and connects it.
The first time Tailscale Serve is used, it may require a browser approval. Use
the link shown by SquadAI, approve it once, then generate the enrollment command
again. Enrollment commands expire after ten minutes and can only be used once.
On the runner machine, later reconnect it with:
squadai runner start
To check its last recorded state:
squadai runner status
There is no native installer or background service required for v1. The runner
is simply a process you start on the machine where work should run.
SquadAI exposes three simple presets:
Permissions can be changed from the agent settings or chat composer. Model and
reasoning changes apply to subsequent turns without discarding the conversation
thread.
When an agent points at a Git repository, SquadAI prepares a managed worktree
under the user's Codex data directory. Instantiated tasks receive separate
branches and worktrees based on the original repository branch. This allows
parallel tasks to modify the same repository without sharing one working tree.
Use Open in VS Code from the agent conversation to inspect that agent's
checkout and diff directly. Worktrees are not deleted automatically when they
contain uncommitted changes.
Any monitor, scheduler, webh

[truncated]
