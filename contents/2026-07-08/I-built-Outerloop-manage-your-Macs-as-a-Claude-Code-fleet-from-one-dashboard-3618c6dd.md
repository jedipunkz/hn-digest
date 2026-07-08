---
source: "https://github.com/phyolim/outerloop"
hn_url: "https://news.ycombinator.com/item?id=48837653"
title: "I built Outerloop – manage your Macs as a Claude Code fleet from one dashboard"
article_title: "GitHub - phyolim/outerloop · GitHub"
author: "phylim"
captured_at: "2026-07-08T22:01:35Z"
capture_tool: "hn-digest"
hn_id: 48837653
score: 1
comments: 0
posted_at: "2026-07-08T21:25:54Z"
tags:
  - hacker-news
  - translated
---

# I built Outerloop – manage your Macs as a Claude Code fleet from one dashboard

- HN: [48837653](https://news.ycombinator.com/item?id=48837653)
- Source: [github.com](https://github.com/phyolim/outerloop)
- Score: 1
- Comments: 0
- Posted: 2026-07-08T21:25:54Z

## Translation

タイトル: アウターループを構築しました – 1 つのダッシュボードから Mac をクロード コード フリートとして管理します
記事タイトル: GitHub - phyolim/outerloop · GitHub
説明: GitHub でアカウントを作成して、phyolim/outerloop の開発に貢献します。

記事本文:
GitHub - フィオリム/アウターループ · GitHub
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
フィオリム
/
アウターループ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル

s
63 コミット 63 コミット デプロイ デプロイ アウターループ アウターループ プロンプト/エージェント プロンプト/エージェント スクリプト スクリプト テスト テスト ui ui .gitignore .gitignore DEPLOY.md DEPLOY.md ライセンス ライセンス README.md README.md crontab.example crontab.example innerloop.plist.example innerloop.plist.example schema.sql schema.sql すべてのファイルを表示 リポジトリ ファイルのナビゲーション
アウターループは、あなたがすでに所有している Mac をクロード コード ワーカーの小さな集団に変えます。
1 つの Web ダッシュボードから管理できます。常時稼働している 1 台の Mac がハブとなり、チケットを保持します
キュー、スケジューラー、ダッシュボード。他の Mac はワーカーとして参加します。シングル
Mac は完全なフリートです。ハブも独自の作業を行います。
それが解決する問題: 作業を実行できるマシンは、通常、現在のマシンではありません。
あなたの目の前に。デスク Mac には Claude Code がログインしており、リポジトリは git/gh で複製されています。
設定されています。携帯電話や別のラップトップからの修正を考えるとき、選択肢は次のとおりです。
リモート接続 — 画面共有またはインタラクティブ セッションへの SSH 接続、低速で扱いにくい
電話するか、後でコピー＆ペーストできるように書き留めておきます。アウターループを使用すると、チケットを提出します
どこにいても、自宅の有能なマシンがそれを拾います。 PR に戻ります。
作業単位はチケットであり、チケットは記録でもあります。エージェントはチケットにコードを書き込みます。
ブランチがチケットをコンテキストとして使用し、2 番目のエージェントがそれをレビューします (作成者とレビュー担当者は
別々の;レビュー担当者はマージできません)、PR はチケットを参照します。マージには次のことが必要です
ダッシュボードでの手動承認、および CI が失敗すると、承認後であってもマージがブロックされます。
マシンは、Bonjour によって、または送信 SSH トンネルを通じて LAN 上で互いを見つけます。
外出時に格安の VPS を利用できます。 VPN、OAuth、クラウド サービスはありません。
すべては http://<hub>.local:8765 の Web ダッシュボードで行われます。
受信箱 – あなたの自宅。何があなたを待っていても (承認

、質問、失敗）
一番上に座ります。フリートが現在行っていることはその下にあります。
ボード — すべてのチケット、Jira スタイル。ここでチケットを作成し ( c を押します)、次の条件でフィルターします
ステータスを確認し、チケットのスレッドをドリルダウンします。
フリート — マシン: オンライン/オフライン、それぞれの実行内容、一時停止/再開、
ペアリング、予算。
アクティビティ — 完全なログ。すべてのアクションは理由とともに記録されます。
コーディング チケットは、次のパスを独自にたどります。
あなたがファイルする → トリアージして優先順位を付ける → エージェントがブランチにコードを書く
→ 2 人目のエージェントがレビューします (3 ラウンド以内) → 安定したら PR が開始されます
→ マージを承認する → マージ、完了
途中、エージェントが一時停止して質問する場合があります。チケットに答えてください。
スレッドが開始され、あなたの回答で作業が再開されます。合流ゲートには PR リンクが表示されます。
diff stat、および CI ステータス。承認とマージ、変更のリクエスト (メモは次のようになります)
エージェントに直接次のパスを求めてください）、停止する場合は拒否してください。メモを追加することもできます
いつでも任意のチケットにアクセスして、次の実行を制御できます。
実行したままにしておくのが安全な理由
すべてのマージを承認します。エージェントは提案するだけです。決定キューが唯一のパスです
取り返しのつかないことまで。 CI が失敗すると、「承認」をクリックした後でもマージがブロックされます。
著者は自分の作品をレビューすることはなく、レビュー担当者にはマージする方法がありません。
何も逃げません。チケットごとおよびフリート全体のトークン予算、実行ごとのタイムアウト、
レビューラウンドには厳しい制限があり、すべての新しい作業を停止するワンクリックキルスイッチがあります。
驚かないでください。通知 URL (outerloop config Notice_url https://ntfy.sh/<topic> ) を設定すると、決定、失敗、およびエラーが発生するたびに電話機がブザー音を鳴らします。
ジャンク駐車券。
LAN に公開されたハブは常に自身をロックします: ワーカー認証がオン、ダッシュボードのパスワードがオン
(自己生成され、設定しない場合は 1 回表示されます。アウターループのステータスによって呼び出されます)。
メニューバー アプリをインストールします。これが最もスムーズな方法です。

は署名され、公証されており、それは
セットアップ全体 (ダッシュボードのパスワード、サービスの開始、ペアリング) を行います。
醸造タップ フィオリム/タップ
brew trust phyolim/tap # 1回限り、カスクに必要
brew install アウターループ # CLI + サービス
brew install --cask innerloop-app # メニューバーアプリ
メニュー バーからアウターループを開き、この Mac が何であるかを選択します。
ハブ + ワーカー — 最初または唯一の Mac での通常の選択: キューを保持し、
ダッシュボードを使用し、コーディング作業自体を行います。
ワーカー — 1 つおきの Mac。ポップオーバーにはネットワーク上のハブがリストされます。[参加] をクリックし、
6 文字のコードが表示されます。そのコードをハブのフリート ページに入力します。ペアリングしました。
ハブ — 珍しい: 調整のみを行い、コーディングを行わないハブ。
それでおしまい。ダッシュボードはメニュー バーまたは http://<hub>.local:8765 から開きます。
1 つのハブで任意の数のワーカーを使用可能。
コーディング作業を行う Mac には、claude (Claude Code CLI、ログイン済み) も必要です。
(認証済み)、および git (コミット ID セット) — アプリは最初の起動時にすべてをチェックします。
Python ≥ 3.9 (組み込みの /usr/bin/python3 ) で十分です。 pip インストールはゼロです。
樽を飛ばしてください。 brew install innerloop の後、Mac にその役割を自分で伝えます。
# 最初/Mac のみ — 次のコードも記述するハブ:
アウターループのセットアップ - 両方の && 醸造サービスがアウターループを開始します # `setup-hub` = 座標のみ
# 1 つおきの Mac — ワーカー:
アウターループローカルロールワーカー
アウターループ ローカル ハブ URL http:// < ハブ > .local:8765
brew サービスがアウターループを開始します
アウターループステータスはダッシュボードのパスワードを出力します。アウターループの医師がチェックする
クロード / gh / git ツールチェーン。ハブから CLI ワーカーをペアリングします: アウターループ トークン <ワーカー名> <シークレット> 、次にワーカー アウターループ ローカル ワーカー <ワーカー名> と
アウターループのローカル トークン <secret> を実行し、再起動します。
アプリをインストールし、 Worker を選択して、その [参加] ボタン (上) を使用します。
ペアリング全体。各労働者に許可されていること

作業内容はその機能によって設定されます。
ハブのフリート ページでライブ編集されます。
ハブは SSH 経由で安価な VPS にダイヤルアウトできるため、ダッシュボードにアクセスできます。
どこでも (HTTPS + パスワード) — 自宅ではポートが開いていません。そのウォークスルーに加えて、
管理対象フリートのゼロタッチ .pkg インストーラーは、
デプロイ/README.md 。
モデル: 安価なモデルは些細な作業を行い、有能なモデルは深い作業を行います (トリアージ →
俳句、評論→ソネット、作家→オーパス）。 --models "author=opus" でオーバーライドするか、
OUTERLOOP_MODEL_<役割> 。
ペルソナ: チームへの人員配置など、エージェントにアイデンティティと専門分野を与えます。をドロップします
ハブのエージェント/データ ディレクトリ内のマークダウン ファイル (「フィンテックに精通した著者による、
銀行業-* プロジェクト」、「食品配達の UX レビュー担当者-*」）およびチケット
プロジェクトごとにマッチングスペシャリストにルートします。プロンプト/エージェント/README.md を参照してください。
状態は ~/Library/Application Support/outerloop に存在します — 1 つのストアが共有されます
CLI、デーモン、アプリ。マシンごとに 1 つのハブを実行し、brew-services を混在させないでください。
同じボックス上の .pkg ハブを備えたハブ。
アップグレード: メニューバー アプリ自体が更新されます。 CLI/サービスの場合、
brew アップグレード アウターループ && 醸造サービス アウターループを再起動します。
brew サービスがアウターループを停止します
ブリューアンインストールアウターループ
brew uninstall --zap --cask innerloop-app
rm -rf ~ /Library/Application \ Support/outerloop # state — 削除したい場合のみ
何もインストールせずに試してみる
FAKE モード — クローンからのデフォルト — はエージェントと GitHub をシミュレートするため、全体
ループは、Mac 独自の Python 以外は何もインストールされていない状態でエンドツーエンドで実行されます。
git clone https://github.com/phyolim/outerloop && cd アウターループ
python3 -m アウターループ初期化
python3 -m innerloop は http://127.0.0.1:8765 で # ダッシュボードを提供します
python3 -m innerloop pick # ループを進める (2 番目のシェル; 繰り返し実行)
チケットをファイルし、数回チェックを入れて、マージに到達するのを確認します

ゲート、承認、もう一度チェック —
合併した。それは製品全体のミニチュアです。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
16
v0.3.15
最新の
2026 年 7 月 8 日
+ 15 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to phyolim/outerloop development by creating an account on GitHub.

GitHub - phyolim/outerloop · GitHub
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
phyolim
/
outerloop
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
63 Commits 63 Commits deploy deploy outerloop outerloop prompts/ agents prompts/ agents scripts scripts tests tests ui ui .gitignore .gitignore DEPLOY.md DEPLOY.md LICENSE LICENSE README.md README.md crontab.example crontab.example outerloop.plist.example outerloop.plist.example schema.sql schema.sql View all files Repository files navigation
Outerloop turns the Macs you already own into a small fleet of Claude Code workers,
managed from one web dashboard. One always-on Mac is the hub : it holds the ticket
queue, the scheduler, and the dashboard. Any other Mac joins as a worker . A single
Mac is a complete fleet — the hub does its own work too.
The problem it solves: the machine that can do the work is usually not the machine in
front of you. The desk Mac has Claude Code logged in, the repos cloned, git/gh
configured. When you think of a fix from your phone or another laptop, the options were
remoting in — screen sharing or SSH into an interactive session, slow and awkward from a
phone — or writing it down to copy-paste later. With outerloop you file a ticket from
wherever you are and a capable machine at home picks it up. You come back to a PR.
The unit of work is a ticket, and the ticket is also the record: an agent writes code on
a branch using the ticket as context, a second agent reviews it (author and reviewer are
separate; the reviewer cannot merge), and the PR references the ticket. Merges require
manual approval in the dashboard, and failing CI blocks a merge even after approval.
Machines find each other on your LAN by Bonjour, or through an outbound SSH tunnel to a
cheap VPS when you're away. No VPN, no OAuth, no cloud service.
Everything happens in the web dashboard at http://<hub>.local:8765 :
Inbox — your home. Whatever is waiting on you (approvals, questions, failures)
sits at the top; what the fleet is doing right now is below it.
Board — every ticket, Jira-style. Create tickets here (press c ), filter by
status, drill into any ticket's thread.
Fleet — your machines: online/offline, what each is running, pause/resume,
pairing, budget.
Activity — the full log; every action is recorded with a why.
A coding ticket walks this path on its own:
you file it → triaged & prioritized → agent writes code on a branch
→ a second agent reviews it (≤3 rounds) → PR opens once it's stable
→ you approve the merge → merged, done
Along the way the agent may pause to ask you a question — answer in the ticket
thread and work resumes with your answer. At the merge gate you see the PR link,
diff stat, and CI status; Approve & merge , Request changes (your note goes
straight to the agent for another pass), or Reject to stop. You can also add a note
to any ticket at any time to steer its next run.
Why it's safe to leave running
You approve every merge. Agents only propose; the decision queue is the sole path
to anything irreversible. Failing CI blocks a merge even after you click approve.
The author never reviews its own work , and the reviewer has no way to merge.
Nothing runs away. Per-ticket and fleet-wide token budgets, per-run timeouts, a
hard cap on review rounds, and a one-click kill switch that stops all new work.
Get pinged, not surprised. Set a notify URL ( outerloop config notify_url https://ntfy.sh/<topic> ) and your phone buzzes for every decision, failure, and
junk-parked ticket.
A LAN-exposed hub always locks itself: worker auth on, and a dashboard password
(self-generated and shown once if you don't set one — outerloop status recalls it).
Install the menu-bar app — it's the smoothest way in. It's signed + notarized, and it
does the whole setup for you (dashboard password, service start, pairing):
brew tap phyolim/tap
brew trust phyolim/tap # one-time, required for the cask
brew install outerloop # CLI + service
brew install --cask outerloop-app # the menu-bar app
Open Outerloop from the menu bar and pick what this Mac is:
Hub + Worker — the usual choice on your first or only Mac: it holds the queue and
dashboard and does coding work itself.
Worker — every other Mac. Its popover lists hubs on your network — click Join ,
it shows a 6-character code, type that code on the hub's Fleet page. Paired.
Hub — rare: a hub that only coordinates and never codes.
That's it. The dashboard opens from the menu bar, or at http://<hub>.local:8765 .
One hub, any number of workers.
Any Mac doing coding work also needs claude (Claude Code CLI, logged in), gh
(authed), and git (commit identity set) — the app checks all of it on first launch.
Python ≥ 3.9 (the built-in /usr/bin/python3 ) is enough; zero pip installs.
Skip the cask; after brew install outerloop , tell the Mac its role yourself:
# first / only Mac — hub that also codes:
outerloop setup-both && brew services start outerloop # `setup-hub` = coordinate-only
# every other Mac — a worker:
outerloop local role worker
outerloop local hub_url http:// < hub > .local:8765
brew services start outerloop
outerloop status prints the dashboard password; outerloop doctor checks the
claude / gh / git toolchain. Pair a CLI worker from the hub: outerloop token <worker-name> <secret> , then on the worker outerloop local worker <worker-name> and
outerloop local token <secret> , and restart it.
Install the app on it, pick Worker , and use its Join button (above) — that's the
whole pairing. What each worker is allowed to work on is set by its capabilities ,
edited live on the hub's Fleet page.
The hub can dial out to a cheap VPS over SSH so you can reach the dashboard from
anywhere (HTTPS + password) — no port opened at home. That walkthrough, plus the
zero-touch .pkg installers for a managed fleet, lives in
deploy/README.md .
Models: cheap models do the trivial work, capable ones the deep work (triage →
Haiku, review → Sonnet, author → Opus). Override with --models "author=opus" or
OUTERLOOP_MODEL_<ROLE> .
Personas: give agents identities and specialties — like staffing a team. Drop a
markdown file in the hub's agents/ data dir ("a fintech-savvy author for
banking-* projects", "a food-delivery UX reviewer for food-* ") and tickets
route to the matching specialist by project. See prompts/agents/README.md .
State lives in ~/Library/Application Support/outerloop — one store shared by the
CLI, the daemon, and the app. Run one hub per machine , and don't mix a brew-services
hub with a .pkg hub on the same box.
Upgrade: the menu-bar app updates itself; for the CLI/service,
brew upgrade outerloop && brew services restart outerloop .
brew services stop outerloop
brew uninstall outerloop
brew uninstall --zap --cask outerloop-app
rm -rf ~ /Library/Application \ Support/outerloop # state — only if you want it gone
Try it without installing anything
FAKE mode — the default from a clone — simulates the agents and GitHub, so the whole
loop runs end to end with nothing installed but a Mac's own Python:
git clone https://github.com/phyolim/outerloop && cd outerloop
python3 -m outerloop init
python3 -m outerloop serve # dashboard at http://127.0.0.1:8765
python3 -m outerloop tick # advance the loop (second shell; run repeatedly)
File a ticket, tick a few times, watch it reach the merge gate, approve it, tick again —
merged. That's the whole product in miniature.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
16
v0.3.15
Latest
Jul 8, 2026
+ 15 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
