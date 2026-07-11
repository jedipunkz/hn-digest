---
source: "https://github.com/rokrak1/agentkindergarten"
hn_url: "https://news.ycombinator.com/item?id=48870547"
title: "AgentKindergarten – daycare for your AI coding agents"
article_title: "GitHub - rokrak1/agentkindergarten: Mirror your AI agent sessions (Claude Code, Codex) and their dev servers to any browser or phone via a tiny self-hosted relay. · GitHub"
author: "rokrak"
captured_at: "2026-07-11T10:55:21Z"
capture_tool: "hn-digest"
hn_id: 48870547
score: 2
comments: 0
posted_at: "2026-07-11T10:08:05Z"
tags:
  - hacker-news
  - translated
---

# AgentKindergarten – daycare for your AI coding agents

- HN: [48870547](https://news.ycombinator.com/item?id=48870547)
- Source: [github.com](https://github.com/rokrak1/agentkindergarten)
- Score: 2
- Comments: 0
- Posted: 2026-07-11T10:08:05Z

## Translation

タイトル: AgentKindergarten – AI コーディング エージェントのための保育園
記事のタイトル: GitHub - rokrak1/agentkindergarten: AI エージェント セッション (Claude Code、Codex) とその開発サーバーを、小さな自己ホスト型リレーを介してブラウザまたは電話にミラーリングします。 · GitHub
説明: AI エージェント セッション (Claude Code、Codex) とその開発サーバーを、小さな自己ホスト型リレーを介して任意のブラウザーまたは携帯電話にミラーリングします。 - rokrak1/エージェント幼稚園

記事本文:
GitHub - rokrak1/agentkindergarten: AI エージェント セッション (Claude Code、Codex) とその開発サーバーを、小さな自己ホスト型リレーを介してブラウザまたは電話にミラーリングします。 · GitHub
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
ロクラク1
/
エージェント幼稚園
公共
通知
署名が必要です

で通知設定を変更します
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
3 コミット 3 コミット .github/ workflows .github/ workflows ランディング ランディング パッケージ パッケージ .dockerignore .dockerignore .gitattributes .gitattributes .gitignore .gitignore DEPLOY.md DEPLOY.md Dockerfile Dockerfile ライセンス ライセンス README.md README.md SPEC.md SPEC.md biome.json biome.json package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI コーディング エージェントのデイケア - ログオフせずにログオフできます。
あなたが72時間連続でバイブコーディングするたびに小言を言う配偶者がいますか？ 「家族のピクニック」は、あなたと完全に良好なマージ競合の間の人質交渉のように感じますか?学校で1時間しか待っていない子供を迎えに行くために、キーボードから体を剥がす必要がありますか?
おめでとうございます。これらは現代のバイブコーダーにとって日常的で耐え難い困難ですが、私はその治療法を構築しました。
AgentKindergarten は、開発環境全体をポケットに直接ストリーミングします。あなたの体がピクニックに参加している間、エージェントは自宅の PC で仕事を続けます。あなたはうなずきます。あなたは「うーん」と言います。クロードがポテトサラダ テーブルから認証レイヤーをリファクタリングする様子を観察します。親指を 1 つ携帯電話に当て、もう 1 つの手で頼んでもいない紙皿を受け取ります。
子供はブランコを押したいですか?代わりにコミットをプッシュします。スイングから。それがマルチタスクです。
義理の親があなたの「キャリアプラン」について尋ねますか？電話の向きを変えると、あなたは今この瞬間、芝生の椅子に座って 7 人の自律型エンジニアを監督しています。彼らは尋ねることをやめるでしょう。
誰かが「ハンバーガーの準備ができました」と叫びましたか？開発サーバーは localhost:3000 にあります - ライブスルーで開きます

h ケチャップが出てくる前のトンネル。
エージェントは決して休みを取ることはありません。今はその必要もありません。草のある場所で行うだけです。
何のために家族を無視しているのか
🧒 学生 — すべてのエージェント セッションは、入力できる実際の端末としてライブ ストリーミングされます (はい、ベンチから)。
🏫 クラスルーム — マシン上のセッションごとに 1 つのダッシュボードがあるため、すべてを 1 か所で無視できます。
📺 Show & Tell — エージェントが起動する開発サーバー (localhost:3000 と友人) は自動検出され、トンネルを通じてライブで開くことができます。車から。ピックアップラインにて。エンジン作動中。
📋 点呼 — どのデーモンとセッションが存在するのかを一目で確認できます。つまり、「存在している」ときに実際に誰が作業しているのかがわかります。
🪞 どこでも同じセッション — セッションは PC 上に存在します。端末、ブラウザ、電話はすべて単なるビューです。自分の机から始めて、義理の両親のトイレから続けて、自分の机に戻ります。フォークも受け渡しも目撃者もいません。
🍎 先生のミス・クランカー — クラス全体の通知表をオプトインします: 各生徒の課題、その進捗状況、およびそれを証明する領収書。彼女はエージェントに小言を言うので、あなたはその必要がなくなります。ついには小言を外注することになりました。独自のクロード CLI、エージェント CLI、または OpenRouter キーによってグレード付けされます。彼女の成績やトークンの請求書などをライブで視聴してください。
🖐️ 挙手 – 失速したり、エラーを出したり、許可のプロンプトに座っていたりした生徒が手を挙げて携帯電話をウェブプッシュする – 感謝祭での猶予を免除するために使用できる真の緊急事態です。
ブランコを押し忘れたり、冷えたハンバーガーを食べたり、お子様の顔に現れたゆっくりとした気づきに対して、私は責任を負いません。
電話/ブラウザ ──HTTPS/WSS──▶ リレー (VPS) ◀──WSS (アウトバウンド)── デーモン (PC)
│ │
━─ Web UI + 認証をホストします ─ エージェント CLI を PTY にラップします。
パー

パススルー、ストレージなし、ポート、プロキシ プレビューを監視
デーモンは開発マシン上で実行され、ダイヤルアウトします。ポート転送や受信ファイアウォール ホールはありません。
リレーは小型の自己ホスト型サーバーであり、ユーザーを認証し、トラフィックをルーティングします。これは純粋なパススルーです。データベースやセッションのログはなく、何も保持されません。
エージェントの資格情報がマシンから離れることはありません。デーモンは、すでにインストールされてログインしている CLI (独自のクロード コード ログイン、独自の API キー) を実行するだけです。 AgentKindergarten はエージェント認証トークンに触れることはありません。
デーモンは、 agk という 1 つの CLI を公開します。エージェントが存在するマシン上で実行します。
npm install - g Agentkindergarten # `agk` コマンドをインストールします
agk init -- リレー wss: // your.domain -- トークン < PAIRING_TOKEN >
agk start # リレーに接続します。実行したままにしておきます
次に、ブラウザ (または携帯電話の PWA) でリレーを開き、次のアカウントでログインします。
アクセス トークンを使用すると、このマシンが 📋 点呼の下に表示されます。常時接続用
(再起動しても存続します)、 DEPLOY.md のタスク スケジューラのレシピを参照してください。
1 台のマシンで 1 つのデーモンが実行されます。 agk start は 2 番目の起動を拒否します
(オーバーライドするには --force を渡します)。何かがおかしい場合は確認してください
~/.agentkindergarten/daemon.out.log 。
毎日使うワンライナー
agkクロード
claude の代わりにこれを入力します。教室でセッションが生まれ、
同時に端末に接続されるため、端末、ブラウザ、
電話は 1 つのセッションの 3 つのビューです。分岐することはありません。すべてのエージェントフラグが合格します
まっすぐに：
Ctrl+] を押して端末を切断します。セッションは継続して実行されます。
教室; agkattach でいつでも再アタッチします。
コマンド
何をするのか
agk スタート
デーモンを開始します。リレーに接続し、生徒を待ちます。実行したままにしておきます。
agk クロード [引数…]
ワンステップで登録と添付が可能です。クロードを 1 番目に開始します

教室でもこの端末でも。 ( agk コーデックスも。)
agk run "<cmd>" --name <n>
アタッチせずにコマンドを生徒として登録します (教室へのストリームのみ)。 --attach を追加して接続することもできます。
agk Attach <ID または名前>
この端末を既存の生徒に tmux スタイルで接続します。 Ctrl+] を押すと、強制終了せずに切り離されます。
AGK LS
現在教室にいる生徒をリストします。
agk rename <id-or-name> <new>
学生のラベルを変更し、<新しい> 作品を agk で添付します。
agk stop <id>
生徒を家に送り返します（プロセスを強制終了します）。スクロールバックは教室内に残ります。
agk close <id-or-name>
生徒タイルを完全に削除します (まだ実行中の場合は、最初に生徒タイルを強制終了します)。
agk先生[オプション]
Hire Miss Clanker: --claude は独自の claude -p (キーなし) で採点します。 --cli "<cmd>" は任意のエージェント CLI を使用します。 --key <k> は OpenRouter を使用します。 --オフは彼女を解雇します。
agkウォッチ[オプション]
注意深い教師: ストール/エラー/許可プロンプトに対する挙手アラート ( --stall <min> 、教師のワンライナーの場合は --escalate on)。
agkスキル[インストール]
バンドルされているクロード コード スキルをリストします。 install はそれらを ~/.claude/skills にコピーします。
agk init [オプション]
~/.agentkindergarten/config.json にデーモン構成を (再) 作成します。
セッションは OS プロセスとして PC 上に存在します。端末やブラウザを決して閉じないでください
それらに影響を与えます。エージェント自身の会話履歴はディスク上に残ります。
再開可能 ( agk claude -c / --resume )。
クロード コードからの駆動 (オプションのスキル)
Claude Code を使用する場合、4 つのスキルがパッケージに同梱されています。これらを一度インストールすると、
チャット内から教室を制御できるため、「会話の途中で、次のことをしたい」という場合に最適です。
携帯電話で続行するには":
agk skill install # それらを ~/.claude/skills にコピーします
(代わりにリポジトリのクローンから: package/daemon/skills/ のフォルダーを ~/.claude/skills/ にコピーします。)
レポートカードと挙手 (ミス・クランカー)
agk 先生 -- クロード # ミス C を雇う

Lanker — 独自の claude CLI を使用して採点します。キーは必要ありません。
agk watch -- stop 10 # 10 分間沈黙した生徒が手を上げる 🖐️
レポートカード (📋 Web UI のレポート) はオプトインです。各生徒は以下に基づいて採点されます。
その任務は何だったのか、その進捗状況、そして領収書（トランスクリプトの抜粋）
それをバックアップしてください。採点が行われる様子をライブでご覧ください — Miss Clanker 自身のトークン請求書も含まれています。
バックエンド: --claude (既存のサブスクリプション。マシンからは何も残りません)、
--cli "codex exec" (任意のエージェント CLI)、または --key <openrouter-key> (トランスクリプトは
OpenRouter へ — あなたからの呼び出しです)。
挙手はデフォルトで有効であり、ストール、エラー、許可など完全にローカルです。
プロンプトは教室内の生徒にフラグを立てます。電話のアラートを有効にして、
ウェブプッシュ。 agk watch --escalate on は、教師による理由の 1 行の概要を追加します。
手は上がっています。
PC 上の何も公開されていないものはありません。デーモンがリレーにダイヤルアウトします —
ポート転送や受信ファイアウォール ホールはありません。デーモンを強制終了すると、PC は
またしてもインターネットの見知らぬ人。
アクセス トークン - ブラウザのログイン。 HTTPS のみ、レート制限があり、
キーがトークン自体から派生する署名付き Cookie - したがって、
リレー上のアクセス トークンにより、全員が即座にログアウトされます。それがあなたのキルスイッチです。
ペアリング トークン — リレーに対してデーモンを認証します。に住んでいます
~/.agentkindergarten/config.json 。
アクセス トークンを保持している人は誰でも、デーモンが実行されるマシンを駆動できます。
それは文字通り製品です。長いランダムトークンを使用し、常にリレーを実行します
HTTPS の背後にあり、トークンを SSH キーのように扱います。
デーモンはリレーソケットを信頼しているため、侵害されたリレーが実際のリレーになります。
最悪の場合。 2 つのデーモン側コントロールがこれを防御します。
したがって、リレー/VPS が完全に所有されている場合でも保持されます。
さらに 2 つの保護が常に適用されます

on : プレビュー トンネルはポートのみに到達できます。
自分の生徒が開いています (データベースや他のローカルホスト サービスにはアクセスできません)
モデル サーバー)、ブラウザは明示的に接続したセッションにのみ入力できます。
パブリックまたは共有リレーの推奨される姿勢: agk init --lock-commands --view-only 。
以下のすべては 1 台のマシン上で実行されます。配線は本番環境と同じですが、ローカルホスト上のリレーのみを使用します。
pnpmインストール
pnpm - r ビルド
#1.リレー（端子1）
$ env: AGK_INSECURE = ' 1 ' ; $ env: AGK_ACCESS_TOKEN = ' トークンの選択 ' ; $ env: AGK_PAIRING_TOKEN = ' 別のものを選択 '
ノードパッケージ/リレー/dist/server.js
#2. デーモン(端末2)
ノード パッケージ / デーモン / dist / cli.js init -- リレー ws: // 127.0 . 0.1 : 8787 -- トークンピック - 別の
ノードパッケージ / デーモン / dist / cli.js 開始
# 3. http://127.0.0.1:8787/app/ を開く — pick-a-token でログインします。
# 「+ 新入生」を押し、`claude` (またはその他) を実行して、ライブで視聴します。
学生内で開発サーバー (例: pnpm dev ) を起動すると、数秒以内にそのポートが 📺 Show & Tell に表示されます。[ライブを開く] をクリックすると、トンネルを通してアプリが表示されます。
DEPLOY.md を参照してください — Docker Compose はリレー + Caddy (自動 Let's Encrypt HTTPS) を起動します。 PC 上のデーモンは agk init --relay wss://your.domain で接続します。 Windows タスク スケジューラのレシピが含まれているため、外出中でもデーモンが起動し続けます。相対

[切り捨てられた]

## Original Extract

Mirror your AI agent sessions (Claude Code, Codex) and their dev servers to any browser or phone via a tiny self-hosted relay. - rokrak1/agentkindergarten

GitHub - rokrak1/agentkindergarten: Mirror your AI agent sessions (Claude Code, Codex) and their dev servers to any browser or phone via a tiny self-hosted relay. · GitHub
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
rokrak1
/
agentkindergarten
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits .github/ workflows .github/ workflows landing landing packages packages .dockerignore .dockerignore .gitattributes .gitattributes .gitignore .gitignore DEPLOY.md DEPLOY.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md SPEC.md SPEC.md biome.json biome.json package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json View all files Repository files navigation
Daycare for your AI coding agents — so you can log off without logging off.
Do you have a spouse who nags every time you vibe-code for 72 hours straight? Does a "family picnic" feel like a hostage negotiation between you and a perfectly good merge conflict? Do you need to peel yourself off the keyboard to go pick up your kid who's only been waiting at school for an hour?
Congratulations: these are the ordinary, unbearable hardships of the modern vibe coder, and I built the cure.
AgentKindergarten streams your entire dev environment straight to your pocket. Your agents keep grinding on your home PC while your body attends the picnic. You nod. You say "mhm." You watch Claude refactor the auth layer from the potato-salad table — one thumb on your phone, one hand accepting a paper plate you did not ask for.
Kid wants a push on the swing? Push a commit instead. From the swing. That's multitasking.
In-law asking about your "career plans"? Turn the phone around — you are, at this very moment, supervising seven autonomous engineers from a lawn chair. They will stop asking.
Someone shouts "burgers are ready"? So is your dev server on localhost:3000 — open it live through the tunnel before the ketchup comes out.
Your agents never take a day off. Now you don't have to either — you just do it somewhere with grass.
What you're neglecting your family for
🧒 Students — every agent session streamed live as a real terminal you can type into (yes, from the bench).
🏫 Classroom — one dashboard for every session on your machine, so you can ignore them all in one place.
📺 Show & Tell — dev servers your agents start ( localhost:3000 & friends) are auto-detected and openable live through the tunnel. From the car. In the pickup line. Engine running.
📋 Roll call — see at a glance which daemons and sessions are present — i.e. who's actually working while you're "present."
🪞 Same session, everywhere — a session lives on your PC; terminal, browser, and phone are all just views of it. Start it at your desk, keep going from the in-laws' bathroom, come back to your desk. No fork, no handoff, no witnesses.
🍎 Miss Clanker, the Teacher — opt-in report cards for the whole class: each student's assignment, how it's going, and the receipts to prove it. She nags the agents so you don't have to — outsource the nagging at last. Graded by your own claude CLI, any agent CLI, or an OpenRouter key; watch her grade live, token bill and all.
🖐️ Raised hands — a student that stalls, errors out, or sits at a permission prompt raises its hand and web-pushes your phone — a genuine emergency you can use to excuse yourself from grace at Thanksgiving.
I accept no liability for missed swing-pushes, cold burgers, or the slow dawning realization on your child's face.
Phone/Browser ──HTTPS/WSS──▶ Relay (your VPS) ◀──WSS (outbound)── Daemon (your PC)
│ │
└─ hosts web UI + auth └─ wraps agent CLIs in PTYs,
pure pass-through, no storage watches ports, proxies previews
The daemon runs on your dev machine and dials out — no port-forwarding, no inbound firewall holes.
The relay is a tiny self-hosted server: it authenticates you and routes traffic. It is a pure pass-through — no database, no logs of your sessions, nothing persisted.
Your agent credentials never leave your machine. The daemon simply runs the CLIs you already have installed and logged in (your own Claude Code login, your own API keys). AgentKindergarten never touches agent auth tokens.
The daemon exposes one CLI: agk . You run it on the machine your agents live on.
npm install - g agentkindergarten # installs the `agk` command
agk init -- relay wss: // your.domain -- token < PAIRING_TOKEN >
agk start # connects to the relay; leave it running
Then open your relay in a browser (or the PWA on your phone), log in with your
access token , and you'll see this machine under 📋 Roll call. For always-on
(survives reboots), see the Task Scheduler recipe in DEPLOY.md .
One machine runs exactly one daemon. agk start refuses to start a second one
(pass --force to override). If something looks off, check
~/.agentkindergarten/daemon.out.log .
The one-liner you'll use every day
agk claude
Type this instead of claude . The session is born in the classroom and
attached to your terminal at the same time — so your terminal, browser, and
phone are three views of the one session. It never forks. All agent flags pass
straight through:
Press Ctrl+] to detach your terminal — the session keeps running in the
classroom; reattach any time with agk attach .
Command
What it does
agk start
Start the daemon: connect to the relay and wait for students. Leave it running.
agk claude [args…]
Enroll + attach in one step. Start Claude in the classroom and in this terminal. ( agk codex too.)
agk run "<cmd>" --name <n>
Enroll a command as a student without attaching (streams to the classroom only). Add --attach to also attach.
agk attach <id-or-name>
Attach this terminal to an existing student, tmux-style. Ctrl+] detaches without killing it.
agk ls
List the students currently in the classroom.
agk rename <id-or-name> <new>
Re-label a student — then agk attach <new> works.
agk stop <id>
Send a student home (kill the process); its scrollback stays in the classroom.
agk close <id-or-name>
Remove a student tile entirely (kills it first if still running).
agk teacher [options]
Hire Miss Clanker: --claude grades with your own claude -p (no key), --cli "<cmd>" uses any agent CLI, --key <k> uses OpenRouter. --off fires her.
agk watch [options]
The watchful teacher: raised-hand alerts for stalls/errors/permission prompts ( --stall <min> , --escalate on for a Teacher one-liner).
agk skills [install]
List the bundled Claude Code skills; install copies them into ~/.claude/skills .
agk init [options]
(Re)create the daemon config at ~/.agentkindergarten/config.json .
Sessions live on your PC as OS processes; closing a terminal or browser never
affects them. Your agent's own conversation history persists on disk and is
resumable ( agk claude -c / --resume ).
Driving it from Claude Code (optional skills)
If you use Claude Code, four skills ship with the package — install them once and
control the classroom from inside a chat, great for "I'm mid-conversation and want
to continue on my phone":
agk skills install # copies them into ~/.claude/skills
(From a repo clone instead: copy the folders in packages/daemon/skills/ into ~/.claude/skills/ .)
Report cards & raised hands (Miss Clanker)
agk teacher -- claude # hire Miss Clanker — grades with your own claude CLI, no key needed
agk watch -- stall 10 # a student that's silent for 10 min raises its hand 🖐️
Report cards (📋 Reports in the web UI) are opt-in: each student gets graded on
what its assignment was, how it's going, and the receipts (transcript excerpts) to
back it up. Watch the grading happen live — Miss Clanker's own token bill included.
Backends: --claude (your existing subscription, nothing leaves your machine),
--cli "codex exec" (any agent CLI), or --key <openrouter-key> (transcripts go
to OpenRouter — your call).
Raised hands are on by default and fully local: stalls, errors, and permission
prompts flag the student in the classroom; enable alerts on the phone to get a
web-push. agk watch --escalate on adds a one-line Teacher summary of why the
hand is up.
Nothing on your PC listens publicly. The daemon dials out to your relay —
no port-forwarding, no inbound firewall holes. Kill the daemon and your PC is a
stranger to the internet again.
Access token — your browser login. HTTPS-only, rate-limited, and it mints a
signed cookie whose key is derived from the token itself — so rotating the
access token on the relay instantly logs everyone out. That's your kill switch.
Pairing token — authenticates the daemon to the relay. Lives in
~/.agentkindergarten/config.json .
Anyone holding your access token can drive the machine the daemon runs on —
that is literally the product. Use a long random token, always run the relay
behind HTTPS, and treat the token like an SSH key.
Because the daemon trusts the relay socket, a compromised relay is your real
worst case. Two daemon-side controls defend against that — enforced on the
daemon, so they hold even if the relay/VPS is fully owned:
Two more protections are always on : the preview tunnel can only reach ports
your own students opened (no reaching other localhost services like a database or
model server), and a browser can only type into a session it explicitly attached.
Recommended posture for a public or shared relay: agk init --lock-commands --view-only .
Everything below runs on one machine — the same wiring as production, just with the relay on localhost.
pnpm install
pnpm - r build
# 1. relay (terminal 1)
$ env: AGK_INSECURE = ' 1 ' ; $ env: AGK_ACCESS_TOKEN = ' pick-a-token ' ; $ env: AGK_PAIRING_TOKEN = ' pick-another '
node packages / relay / dist / server.js
# 2. daemon (terminal 2)
node packages / daemon / dist / cli.js init -- relay ws: // 127.0 . 0.1 : 8787 -- token pick - another
node packages / daemon / dist / cli.js start
# 3. open http://127.0.0.1:8787/app/ — log in with pick-a-token,
# hit "+ New student", run `claude` (or anything), and watch it live.
Start a dev server inside a student (e.g. pnpm dev ) and its port shows up under 📺 Show & Tell within seconds — click "open live" to view the app through the tunnel.
See DEPLOY.md — Docker Compose brings up the relay + Caddy (automatic Let's Encrypt HTTPS); the daemon on your PC connects with agk init --relay wss://your.domain . Includes a Windows Task Scheduler recipe so the daemon stays up while you're away. The rel

[truncated]
