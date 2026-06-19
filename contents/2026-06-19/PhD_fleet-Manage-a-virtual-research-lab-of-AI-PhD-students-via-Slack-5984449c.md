---
source: "https://github.com/canatara/phd_fleet"
hn_url: "https://news.ycombinator.com/item?id=48600475"
title: "PhD_fleet – Manage a virtual research lab of AI PhD students via Slack"
article_title: "GitHub - canatara/phd_fleet: Manage your own AI research lab. Act as the advisor to a fleet of virtual PhD students. Assign tasks, review reported results, and iterate on their independent research projects directly through a Slack interface. · GitHub"
author: "canatara"
captured_at: "2026-06-19T18:44:06Z"
capture_tool: "hn-digest"
hn_id: 48600475
score: 2
comments: 2
posted_at: "2026-06-19T16:45:59Z"
tags:
  - hacker-news
  - translated
---

# PhD_fleet – Manage a virtual research lab of AI PhD students via Slack

- HN: [48600475](https://news.ycombinator.com/item?id=48600475)
- Source: [github.com](https://github.com/canatara/phd_fleet)
- Score: 2
- Comments: 2
- Posted: 2026-06-19T16:45:59Z

## Translation

タイトル: PhD_fleet – Slack 経由で AI 博士課程の学生の仮想研究室を管理する
記事タイトル: GitHub - canatara/phd_fleet: 自分の AI 研究ラボを管理する。仮想の博士課程学生群のアドバイザーとして機能します。 Slack インターフェースを介して直接、タスクを割り当て、報告された結果を確認し、独立した研究プロジェクトを繰り返し実行します。 · GitHub
説明: 自分の AI 研究ラボを管理します。仮想の博士課程学生群のアドバイザーとして機能します。 Slack インターフェースを介して直接、タスクを割り当て、報告された結果を確認し、独立した研究プロジェクトを繰り返し実行します。 - カナタラ/phd_fleet

記事本文:
GitHub - canatara/phd_fleet: 独自の AI 研究ラボを管理します。仮想の博士課程学生群のアドバイザーとして機能します。 Slack インターフェースを介して直接、タスクを割り当て、報告された結果を確認し、独立した研究プロジェクトを繰り返し実行します。 · GitHub
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
{{

メッセージ }}
カナタラ
/
phd_fleet
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット ライブラリ library mentor_template mentor_template src src Student_template Student_template .env.example .env.example .gitignore .gitignore LAB_CONTEXT.md LAB_CONTEXT.md ライセンス ライセンス README.md README.md bot.py bot.py pyproject.toml pyproject.toml要件.txt要件.txtslack-manifest.yamlslack-manifest.yamlすべてのファイルを表示リポジトリ ファイルのナビゲーション
phd_fleet — Slack 上で「博士課程の学生」エージェントのフリートを実行する
1 人の研究者 (アドバイザー) が生成し、
クロード・コードの艦隊と会話する
Slack 経由でエージェントに連絡します。各エージェントは、独自のクロード コード セッションです。
ワークスペースディレクトリ。 Slack メッセージがターンを促進します。ターンの間にエージェントの
ファイルシステムはその長期メモリです。
別のコーチエージェントが、アドバイザーがどのようにアドバイスし、与えるかを監視します。
メンタリング技術に関する証拠に基づいたフィードバック。絡み合った 2 つの目標: 実現する
研究を行い、メンターとして成長します。
ホストには、Python と Claude Code CLI を実行するものであれば何でもかまいません (ラップトップ、
ラボ ワークステーション、クラウド VM、HPC ログイン ノード。ボットは Slack に接続します
ソケット モードなので、ホストはインバウンド HTTP を必要とせず、NAT の背後またはネットワーク内で動作します。
プライベートサブネット。
エージェントごとに 1 つの Slack チャネル。 #student-<name> は、生成する生徒ごとに指定します。
コーチの #mentor-coach (最初の起動時に自動的に作成されます)。
3 つのスラッシュ コマンド。
/new-student <名前> <briefing> ワークスペースをスキャフォールディングし、チャネルを作成し、
そして最初のターンが始まります。
/coach-review <名前> [日] は、コーチにあなたのメンタリングをレビューするよう依頼します。
特定の生徒。
/claude-status は、簡単なローカル レポートを出力します。

adout — ターン、最後のコンテキスト サイズ、
エージェントごとの累積トークン、モデル、コスト、GitHub リンク。クロードからの電話はありません。
単なるレジストリビューです。
エージェントごとのジャーナル。各エージェントのワークスペースの JOURNAL.md は追加専用です。
ターンごとに 1 つのセクションがあり、Did / Found / Next ブロックで終わります。
紙の共有ライブラリ。プロジェクトのルートにある library/ は、次ごとに 1 つのプールです。
エージェントは読み取りと貢献を行います。論文の最初の読者は次のように書きます
正規の要約。後で読む人は別のメモ ファイルを追加します。ライブラリ/README.md
は再生成されたインデックスです。エージェントは読み取りのみを行います。
エージェントごとの GitHub ブランチへのターンごとのコミット (オプション)。設定する場合
オリジン 、ボットがエージェントのワークスペースを新しいステージに切り替えるたびに
エージェント/<名前> でコミットし、リースを使用して強制プッシュします。スラック決勝
メッセージはレビュー用のブランチへのリンクです。原点がない場合、このステップはスキップされます
静かに。
デフォルトでは静かです。ターン中、ボットは最大でも 1 つの短い「開始」を投稿します。
メッセージとエージェントの最終応答 - ツールごとの装飾的なストリームはありません
メッセージ。失敗とタイムアウトが発生すると、:warning: 行が表示されます。
PATH 上のクロード コード CLI。ボットはそれに攻撃を仕掛けるので、claude --version
ボットを実行するのと同じシェルで動作する必要があります。
Claude.ai サブスクリプションまたは Anthropic Console API キー (「
クロード認証）。
カスタム アプリをインストールできる Slack ワークスペース (無料または有料)。
オプション: ターンごとのレビュー ブランチが必要な場合は、GitHub リポジトリ。ボット
何もなくても問題なく動作します。
git clone <このリポジトリの URL >
cd phd_fleet
python3 -m venv .venv
.venv/bin/pip install -rrequirements.txt
cp .env.example .env # 次に、3 つの必要な値を入力します
.venv/bin/python bot.py
次に、Slack で /new-student alice "your project Briefing" を実行して、
最初のエージェント。このセクションの残りの部分では、各手順について説明します。
上記のクイック スタートを参照してください。仮想

env は推奨されますが、必須ではありません - 任意
requirements.txt に依存関係がある環境では機能します。
リポジトリには、slack-manifest.yaml が同梱されます。
すでに 3 つのスラッシュ コマンド、ボット スコープ、イベントがリストされています。
サブスクリプション。使用するには:
https://api.slack.com/apps → 新しいアプリの作成 → マニフェストから に移動します。
ワークスペースを選択し、slack-manifest.yaml の内容を貼り付けて確認します。
[基本情報] → [アプリレベルのトークン] → [トークンとスコープの生成] で、
Connections:write スコープを使用してトークンを作成します。それをコピーします ( xapp-… ) — それは
あなたの SLACK_APP_TOKEN 。
[OAuth と権限] で、アプリをワークスペースにインストールします。をコピーします
ボット ユーザー OAuth トークン ( xoxb-… ) — これが SLACK_BOT_TOKEN です。
cp .env.example .env
# .env を編集し、SLACK_BOT_TOKEN、SLACK_APP_TOKEN、ADVISOR_SLACK_USER_ID を入力します
ADVISOR_SLACK_USER_ID を確認するには、Slack でプロフィール写真をクリック →
フルプロフィールを表示 → […] メニュー → [メンバー ID をコピー] 。のように見えます
U0123ABC456 。ボットはユーザー ID が一致しない人からのメッセージを拒否します
この値は、唯一のアクセス制御面です。
2 つのパス (優先順):
定期購読（推奨）。クロードが OAuth 認証されている場合、
Claude.ai アカウント — つまり、~/.claude/.credentials.json は実行中に存在します。
claude /login Once — ボットはそれを使用するだけです。 ANTHROPIC_API_KEY は設定しないままにしておきます。
APIキー。 Anthropic コンソールを通じてトークンごとに支払うには、次のように設定します。
ANTHROPIC_API_KEY=sk-ant-… .env 内。 Agent SDK はそれを取得して無視します。
サブスクリプションのパス。
.venv/bin/python bot.py
以下が表示されるはずです:
… 情報準備完了 — ソケット モードでリッスン中 (advisor=U0123ABC456)
… 情報 新しいセッション (s_…) が確立されました
… お知らせ ⚡️Bolt アプリが稼働中です!
ヘルスチェック: ボットが存在する任意のチャネルで、@<bot> ping が返されます。
ポン。

長時間実行する操作の場合は、プロセス スーパーバイザーの下でボットを実行します。tmux は
クラッシュ時に自動再起動したい場合は、最も単純なパス systemd --user を使用します。ボット
一時的な切断時に Slack に再接続しますが、完全なプロセスを続行できません
監督者なしで終了します。
/new-student alice 「無秩序な領域に関する AlphaFold の信頼度を調査してください。最近の文献を読んでから、小さな実験を提案してください。」
これが行うこと:
^[a-z0-9][a-z0-9_-]{0,40}$ に対して名前を検証します。
Student_template/ から students/alice/ を足場にし、名前と名前を入力します。
プロジェクトの説明会。
Slack で #student-alice を作成して招待します。
Agents.json (ランタイム レジストリ) にアリスを登録します。
最初のターンが開始されます。エージェントは CLAUDE.md を読み取り、指示を受けて、
が報告します。
その後、#student-alice で送信するすべてのメッセージが次のプロンプトになります。
エージェントのセッションは、ターン間およびボットの再起動後に再開されます。
コーチには、最初のボット起動時に作成される独自のチャネル #mentor-coach があります。 2
使い方:
無料チャット。 #mentor-coach に書き込んだものはすべてプロンプトになります —
「間違っていると思う方法を提案する学生にどう対処すればよいでしょうか?」コーチ
コーチングのような声で答え、明確な質問をし、関連する質問をします。
必要に応じてフレームワークを使用します。
構造化されたレビュー。 /coach-review alice 7 は過去 7 日間を取得します
#student-alice に、Alice の JOURNAL.md の最近の抜粋を加えて、
アリスへのあなたの指導をレビューするコーチ — 何がうまくいったか、何がうまくいく可能性があるか
より鮮明で、それぞれが特定の瞬間に結びついています。結果は #mentor-coach に投稿されます。
また、コーチは長期的な mentor/coach/notes/advisees/alice.md も更新します。
/クロードステータス
すべてのエージェントをリストする一時的な応答: 種類、ターン数、最後のコンテキスト サイズ、
累積入出力トークン、モデル、総コスト、および GitHub ブランチ

リンクする場合
利用可能です。純粋なローカル読み取り — クロードを呼び出しません。
わざとマニュアル化。何を持っているかを確認するには、students/ を参照してください。アーカイブするには
Student: mv students/<name> students/_archived/ からそのエントリを削除します
エージェント.json 。年に 2 回行う操作にスラッシュ コマンドを使用する価値はありません
彼らの体重。
各学生は、students/<name>/ に住んでいます。
CLAUDE.md — ペルソナとプロジェクトの概要。作成時に入力されます。
Student_template/ 。
JOURNAL.md — 追加専用の調査ログ、セッションごとに 1 セクション。
Notes/ — プライベートのスクラッチパッド、設計メモ、中間分析。
ディレクトリの残りの部分 - 実際の作業成果物 (コード、データ、結果)。
各生徒には、何かを読む前に図書館/（共有プール）を確認するように指示されています
新しい、そこに新しい論文の要約を書き、決して git を実行しない - ボットが処理します
出版すること。
ラボ全体の習慣と慣習は LAB_CONTEXT.md にあります。
すべてのエージェントのシステム プロンプトに自動的に追加されます。一度編集すれば、すべてのエージェントが選択します
次のターンで新しいルールが適用されます。
コーチは、学生と同じワークスペース構造を持つ、mentor/coach/ に住んでいます。
加えて、長期的な観察用にnotes/advisees/<name>.mdを追加します。そのCLAUDE.md
コーチングのペルソナと名前付きフレームワーク (GROW、SBI、
ヴィゴツキーの ZPD。フィードフォワード）。同じランナー、同じエージェントごとのロックを使用します。
そして学生と同じ足場。コーチは受け身のみです。コーチはいつ話すのかを話します。
召喚 — 学生のターンごとに自動レビューは行われません。
library/ は、すべてのエージェントが読み取るプロジェクト ルートにある単一のディレクトリです。
そして以下に書きます：
論文 X の最初の読者が library/<citekey>.md (マークダウンの概要) を書き込みます
YAML フロントマター付き) と library/<citekey>.pdf (自由にダウンロードできる場合)。
後で自分の意見を追加したい読者は、別のファイルを作成します —
library/<citekey>__notes_<彼らの

-name>.md 。彼らはピアの概要を編集することはありません。
library/README.md がインデックスです。エージェントは読むだけです。ボット
library/*.md を実行し、frontmatter を解析することで、各ターン後にそれを再生成します。
そしてテーブルを書き換えます。
Citekey の衝突 (同じ年の同じ著者による 2 つの論文) が解決されました
文字の接尾辞が付いています — Jumper2021a 、 Jumper2021b 。完全な規約が存在します
LAB_CONTEXT.md 。このファイル所有権ごとの形状により、失敗が回避されます。
2 つのエージェントが同時に同じインデックスに追加し、相互に破壊するモード。
Origin を構成している場合、ボットは各エージェントのワークスペースを
各ターン後のエージェントごとのブランチ (agent/<name>)、
リース。一時的な git インデックスに段階的に移行するため、ボットのコミットが妨げられることはありません
あなたの作業ツリー。 Slack の最終メッセージは GitHub のブランチにリンクしています。の
段階的な理論的根拠はコメントに記載されています。
src/agents.py ( commit_and_push )。
これを望まない場合は、origin を追加しないでください。ボットは公開をスキップします
サイレントステップを実行すると、Slack メッセージにレビューリンクが含まれなくなります。
ほとんどの動作は .env で設定されます。エージェントごとの .claude/settings.json ファイルには次の内容が含まれます。
許可拒否リスト (「セキュリティ モデル」を参照)。 .env ノブ:
これは多層防御であり、サンドボックスではありません。各エージェントの .claude/settings.json
許可を持っています d

[切り捨てられた]

## Original Extract

Manage your own AI research lab. Act as the advisor to a fleet of virtual PhD students. Assign tasks, review reported results, and iterate on their independent research projects directly through a Slack interface. - canatara/phd_fleet

GitHub - canatara/phd_fleet: Manage your own AI research lab. Act as the advisor to a fleet of virtual PhD students. Assign tasks, review reported results, and iterate on their independent research projects directly through a Slack interface. · GitHub
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
canatara
/
phd_fleet
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit library library mentor_template mentor_template src src student_template student_template .env.example .env.example .gitignore .gitignore LAB_CONTEXT.md LAB_CONTEXT.md LICENSE LICENSE README.md README.md bot.py bot.py pyproject.toml pyproject.toml requirements.txt requirements.txt slack-manifest.yaml slack-manifest.yaml View all files Repository files navigation
phd_fleet — run a fleet of "PhD student" agents over Slack
A small Python toolkit that lets one researcher — the advisor — spawn and
converse with a fleet of Claude Code
agents through Slack. Each agent is its own Claude Code session in its own
workspace directory; Slack messages drive turns; between turns, the agent's
filesystem is its long-term memory.
A separate coach agent watches how the advisor advises and gives
evidence-based feedback on mentoring craft. The two intertwined goals: get real
research done and grow as a mentor.
The host can be anything that runs Python and the Claude Code CLI — a laptop, a
lab workstation, a cloud VM, an HPC login node. The bot connects to Slack over
Socket Mode, so the host needs no inbound HTTP and works behind a NAT or in a
private subnet.
One Slack channel per agent. #student-<name> for each student you spawn;
#mentor-coach for the coach (created automatically at first startup).
Three slash commands.
/new-student <name> <briefing> scaffolds a workspace, creates the channel,
and kicks off the first turn.
/coach-review <name> [days] asks the coach to review your mentoring of a
specific student.
/claude-status prints a quick local readout — turns, last context size,
cumulative tokens, model, cost, and GitHub link per agent. No Claude calls;
just a registry view.
Per-agent journal. JOURNAL.md in each agent's workspace is append-only,
one section per turn, ending with a Did / Found / Next block.
Shared paper library. library/ at the project root is a single pool every
agent reads from and contributes to. First reader of a paper writes the
canonical summary; later readers add a separate notes file. library/README.md
is a regenerated index — agents only read it.
Per-turn commits to per-agent GitHub branches (optional). If you configure
an origin , after each turn the bot stages that agent's workspace into a fresh
commit on agent/<name> and force-pushes it with a lease. The Slack final
message links to the branch for review. With no origin , this step is skipped
silently.
Quiet by default. During a turn the bot posts at most one short "started"
message and the agent's final reply — no per-tool stream of decorative
messages. Failures and timeouts post a :warning: line.
The Claude Code CLI on PATH . The bot shells out to it, so claude --version
must work in the same shell where you run the bot.
A Claude.ai subscription or an Anthropic Console API key (see
Claude authentication ).
A Slack workspace (free or paid) where you can install a custom app.
Optional: a GitHub repository if you want per-turn review branches. The bot
runs fine without one.
git clone < this repo URL >
cd phd_fleet
python3 -m venv .venv
.venv/bin/pip install -r requirements.txt
cp .env.example .env # then fill in the three required values
.venv/bin/python bot.py
Then in Slack, run /new-student alice "your project briefing" to spawn your
first agent. The rest of this section explains each step.
See the Quick start above. A virtualenv is recommended but not required — any
environment with the dependencies in requirements.txt works.
The repo ships slack-manifest.yaml — a manifest that
already lists the three slash commands, the bot scopes, and the event
subscriptions. To use it:
Go to https://api.slack.com/apps → Create New App → From a manifest .
Pick your workspace, paste the contents of slack-manifest.yaml , and confirm.
Under Basic Information → App-Level Tokens → Generate Token and Scopes ,
create a token with the connections:write scope. Copy it ( xapp-… ) — that's
your SLACK_APP_TOKEN .
Under OAuth & Permissions , install the app to your workspace. Copy the
Bot User OAuth Token ( xoxb-… ) — that's your SLACK_BOT_TOKEN .
cp .env.example .env
# edit .env and fill in SLACK_BOT_TOKEN, SLACK_APP_TOKEN, ADVISOR_SLACK_USER_ID
To find your ADVISOR_SLACK_USER_ID : in Slack, click your profile picture →
View full profile → the … menu → Copy member ID . It looks like
U0123ABC456 . The bot rejects messages from anyone whose user ID does not match
this value — the only access-control surface.
Two paths, in order of preference:
Subscription (recommended). If claude is OAuth-authenticated to your
Claude.ai account — i.e., ~/.claude/.credentials.json exists from running
claude /login once — the bot just uses it. Leave ANTHROPIC_API_KEY unset.
API key. To pay per token through the Anthropic Console, set
ANTHROPIC_API_KEY=sk-ant-… in .env . The Agent SDK picks it up and ignores
the subscription path.
.venv/bin/python bot.py
You should see:
… INFO ready — listening on Socket Mode (advisor=U0123ABC456)
… INFO A new session (s_…) has been established
… INFO ⚡️ Bolt app is running!
Health check: in any channel where the bot is present, @<bot> ping returns
pong .
For long-running operation, run the bot under a process supervisor — tmux is
the simplest path, systemd --user if you want auto-restart on crash. The bot
reconnects to Slack on transient disconnects but cannot survive a full process
exit without a supervisor.
/new-student alice "Investigate AlphaFold confidence on disordered regions. Read the recent literature, then propose a small experiment."
What this does:
Validates the name against ^[a-z0-9][a-z0-9_-]{0,40}$ .
Scaffolds students/alice/ from student_template/ , filling in the name and
project briefing.
Creates #student-alice in Slack and invites you.
Registers Alice in agents.json (the runtime registry).
Kicks off the first turn — the agent reads its CLAUDE.md , gets oriented, and
reports back.
After that, every message you send in #student-alice becomes the next prompt.
The agent's session resumes across turns and across bot restarts.
The coach has its own channel, #mentor-coach , created at first bot startup. Two
ways to use it:
Free chat. Anything you write in #mentor-coach becomes a prompt —
"How should I handle a student proposing a method I think is wrong?" The coach
responds in a coaching voice, asks clarifying questions, and names a relevant
framework when appropriate.
Structured review. /coach-review alice 7 pulls the last 7 days of
#student-alice plus a recent excerpt of Alice's JOURNAL.md , and asks the
coach to review your mentoring of Alice — what was done well, what could be
sharper, each tied to a specific moment. The result is posted in #mentor-coach ,
and the coach also updates a longitudinal mentor/coach/notes/advisees/alice.md .
/claude-status
Ephemeral reply listing every agent: kind, turns taken, last context size,
cumulative input/output tokens, model, total cost, and the GitHub branch link if
available. Pure local read — does not call Claude.
Manual on purpose. To see what you have, look in students/ . To archive a
student: mv students/<name> students/_archived/ and remove its entry from
agents.json . Slash commands for operations you'll do twice a year aren't worth
their weight.
Each student lives in students/<name>/ :
CLAUDE.md — persona and project briefing, filled in at create time from
student_template/ .
JOURNAL.md — append-only research log, one section per session.
notes/ — private scratchpads, design notes, intermediate analyses.
the rest of the directory — actual work artifacts (code, data, results).
Each student is told to check library/ (the shared pool) before reading anything
new, to write new paper summaries there, and never to run git — the bot handles
publishing.
The lab-wide habits and conventions live in LAB_CONTEXT.md ,
auto-appended to every agent's system prompt. Edit it once and every agent picks
up the new rules on its next turn.
The coach lives in mentor/coach/ , with the same workspace structure as a student
plus notes/advisees/<name>.md for longitudinal observations. Its CLAUDE.md
carries a coaching persona and a vocabulary of named frameworks (GROW; SBI;
Vygotsky's ZPD; feedforward). It uses the same runner, the same per-agent lock,
and the same scaffolding as a student. The coach is reactive only: it speaks when
summoned — there is no auto-review after every student turn.
library/ is a single directory at the project root that every agent reads from
and writes to:
First reader of paper X writes library/<citekey>.md (a markdown summary
with YAML frontmatter) and library/<citekey>.pdf if it's freely downloadable.
Later readers who want to add their take write a separate file —
library/<citekey>__notes_<their-name>.md . They never edit a peer's summary.
library/README.md is the index. Agents only read it. The bot
regenerates it after each turn by walking library/*.md , parsing frontmatter,
and rewriting the table.
Citekey collisions (two papers by the same author in the same year) are resolved
with letter suffixes — jumper2021a , jumper2021b . The full conventions live in
LAB_CONTEXT.md . This per-file-ownership shape avoids the failure
mode where two agents append to the same index at once and clobber each other.
If you've configured origin , the bot publishes each agent's workspace to a
per-agent branch ( agent/<name> ) after every turn, using a force-push with a
lease. It stages into a temporary git index so the bot's commits never disturb
your working tree. The Slack final message links to the branch on GitHub. The
step-by-step rationale is documented in the comments of
src/agents.py ( commit_and_push ).
If you don't want this, simply don't add an origin . The bot skips the publish
step silently and the Slack messages won't include review links.
Most behavior is set in .env ; the per-agent .claude/settings.json files carry
the permission deny-list (see Security model ). The .env knobs:
This is defense-in-depth, not a sandbox. Each agent's .claude/settings.json
carries a permission d

[truncated]
