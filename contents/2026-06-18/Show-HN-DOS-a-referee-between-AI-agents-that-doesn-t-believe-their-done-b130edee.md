---
source: "https://github.com/anthony-chaudhary/dos-kernel"
hn_url: "https://news.ycombinator.com/item?id=48588950"
title: "Show HN: DOS – a referee between AI agents that doesn't believe their \"done\""
article_title: "GitHub - anthony-chaudhary/dos-kernel: Catch your AI agents when they lie about what they shipped — verifies claims against git instead of believing the agent. · GitHub"
author: "anthonysarkis"
captured_at: "2026-06-18T18:54:34Z"
capture_tool: "hn-digest"
hn_id: 48588950
score: 2
comments: 0
posted_at: "2026-06-18T17:52:32Z"
tags:
  - hacker-news
  - translated
---

# Show HN: DOS – a referee between AI agents that doesn't believe their "done"

- HN: [48588950](https://news.ycombinator.com/item?id=48588950)
- Source: [github.com](https://github.com/anthony-chaudhary/dos-kernel)
- Score: 2
- Comments: 0
- Posted: 2026-06-18T17:52:32Z

## Translation

タイトル: Show HN: DOS – 自分たちの「完了」を信じない AI エージェント間の審判
記事のタイトル: GitHub - anthony-chaudhary/dos-kernel: AI エージェントが出荷したものについて嘘をついたときを捕まえる — エージェントを信じるのではなく、git に対する主張を検証します。 · GitHub
説明: AI エージェントが出荷したものについて嘘をついているのを捕まえます。エージェントを信じるのではなく、git に対して主張を検証します。 - アンソニー・チョーダリー/dos-kernel

記事本文:
GitHub - anthony-chaudhary/dos-kernel: AI エージェントが出荷したものについて嘘をついたときに捕まえます。エージェントを信じるのではなく、git に対して主張を検証します。 · GitHub
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
アンソニー・チョーダリー
/
dos-カーネル
公共
通知
あなたは署名している必要があります

通知設定を変更するには
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
709 コミット 709 コミット .claude-plugin .claude-plugin .claude .claude .github .github ベンチマーク ベンチマーク claude-plugin claude-plugin docker docker docs docs 例 例 gitlab-ci gitlab-ci go go リスト リスト パッケージング/ aur パッケージング/ aur 紙 紙スクリプト スクリプト スパイク スパイク src src テスト テスト verify-action verify-action .dockerignore .dockerignore .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore .pre-commit-hooks.yaml .pre-commit-hooks.yaml AGENTS.md AGENTS.md CITATION.cff CITATION.cff CLAUDE.md CLAUDE.md COTRIBUTING.md COTRIBUTING.md Dockerfile Dockerfile GEMINI.md GEMINI.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md context7.json context7.json dos.toml dos.toml gemini-extension.json gemini-extension.json Glama.json Glama.json install.ps1 install.ps1 install.py install.py install.sh install.sh llms-full.txt llms-full.txt llms-install.md llms-install.md llms.txt llms.txt pyproject.toml pyproject.toml server.json server.json smithery.yaml smithery.yaml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
DOS — ディスパッチ オペレーティング システム
AI エージェントが出荷内容について嘘をついた場合は捕まえてください。
📊 実際のリポジトリでの実行を確認します: スコアボード
15 の人気のある AI 構築リポジトリ (roborev、オープン インタプリタ、crewAI、autogen など) をスコア化
— エージェントがどれだけの内容を書いたか、どの内容を書いたか、そして各コミットの主張が裏付けられているかどうか
独自の差分によって。スコアを付けます: dos commit-audit --cookie --workspace 。ベース..ヘッド。
1 つの録音にすべてのピッチが含まれています。エージェントは 2 つの機能が出荷されていると主張しています。 git は 1 を返します。
dos はコミットからの回答を検証します。嘘は exit 1 であり、その上のゲートです。
終了コードは f を拒否します

「完了」も。すべての行が実際の CLI のそのままの出力です。
scripts/build_caught_lie_cast.py は、出力が変更されるたびにそれを再記録します。
1 つのリポジトリでエージェント フリートを実行します。左のループはただ進んでいるように感じます。あなたが操縦できるのは正しいものです。
唯一の違いは、DOS がエージェントの言葉ではなく、現実世界 (ここでは git) から読み取る評決です。
AI エージェントが完了したことを伝えます。 DOS は代わりに現実世界をチェックします。
その言葉をそのまま受け入れると、現実世界に最も近い部分は git 履歴です。
エージェントは、ログイン エンドポイントを出荷したと述べています。そうでしたか？ 1 つのコマンドを実行します。
dos verify は、作業が残した成果物からではなく、成果物から応答します。
エージェントが入力した内容: コミットによりクレームが返される → SHIPPED 、 exit 0 ;何もない
着陸 → NOT_SHIPPED、出口 1 。エージェントの話は決して入りません。 (Git
これは、DOS が読み取る最初の証人にすぎません。ファイル ツリー、クロック、CI ステータス、
テスト環境自体の状態は他の状態、つまりエージェントが作成したものではありません。)
dos verify AUTH AUTH1 # → 出荷された AUTH AUTH1 e62f74d (終了 0)
dos verify AUTH AUTH2 # → NOT_SHIPPED AUTH AUTH2 (終了 1)
それが最小バージョンです。スケールアップも可能: 十数人のエージェントを 1 人に向ける
リポジトリ — CI 内、フリート内、同じファイル上で競合 — そして DOS も通知します
どれがお互いを踏んでいるのか、どれがぐるぐる回っているのか、そして
どちらの「完了」という主張が真実なのか。すべての答えは成果物 (git、
ファイルツリー、時計）、決してナレーションではありません。プレーンな Git リポジトリで動作します
設定はゼロで、教えれば教えるほど賢くなる、これが唯一のことです
install は 1 つの小さな Python パッケージです。
⚡ 追加するだけです。コマンドは 2 つで、決定は必要ありません。あなたの場所のリポジトリから
エージェントの動作:
pip インストール dos-kernel
dos init --hooks auto # すでに使用しているエージェント ランタイムを検索し、チェックに接続します
から

その後: 実際に作業が行われない限り、エージェントは「完了」とは言えません。
着陸すると、2 人のエージェントがお互いのファイルを黙って上書きすることはできず、実行されます。
そのストールには静かに回転する代わりにフラグが立てられます。あなたについては何もありません
ワークフローが変わるため、以下の語彙を学ぶ必要はありません。
カバーされる。作成した 1 つの構成ファイルを出力します。 dosフックを削除する
そこにエントリを入力すると元に戻ります。 (ランタイムが検出されませんか? その旨が表示され、
名前を選ぶことができます - 推測はできません。)
v0.28.0 · 5,600 以上のテスト · CI: Linux 上の Python 3.11 ～ 3.13 + Windows 3.13
スモークラン - 唯一のランタイム依存関係は PyYAML - MIT です。
🧭 次に進むべき方向: 理由と証拠 (わかりやすい言葉のストーリー、bash の 20 行の答え、証明された内容)、
それをスタックに配線します (MCP、フック、インストール)。
syscall + CLI リファレンス 、またはこれを AI エージェントとして読みますか? 、AGENTS.md — 3 行でビルド/テスト/チェック。全体マップはすぐ下のルーターです。
🔤 このページの残りの部分は 5 つの単語に依存します。計画とは名前付きの目標です
(認証);フェーズは、その出荷可能な 1 つのステップ ( AUTH1 ) です。車線というのは
エージェントがアクセスできるファイル ツリーのスライス。オラクルは DOS の一部です
それは証拠とルールを読みます。スタンプは出荷段階のマークです
コミットサブジェクト (AUTH1: …) に残ります。これは、オラクルが grep するものです。
それが語彙全体です。
コーディング エージェントは実際に機能し、その結果がどうなったかを教えてくれます。通常、その話は真実です。
時には「すべての作業が完了しました！」という陽気な声も聞こえます。出荷した作業員から
何もない。 1 つのエージェントを使用すると、その出力を読み直すことで自分でそれを把握できます。
すでに支払っている税金。一度に20台走れば税金は支払えなくなる：誰も
すべてを読み取り、各労働者が自分の宿題を採点し、未チェックの問題を解決します。
コードベースがある程度機能し、誰も安全に変更できなくなるまで、静かに積み重ねてください。

t.
DOS はストーリーを決して読むのではなく、何が起こったのかを読む審判です (
コミット、ファイル、時計）そしてナレーションが動かせない評決をあなたに渡します。それ
コストは午後約 1 つで、ランタイム依存関係が 1 つあり、そのレーンに留まります。
コードが良いかどうかではなく、何が起こったのかを知らせます - 品質は維持されます
あなたのテストとレビュー。 （完全な平文バージョン。）
ここにあるすべての数値は、エージェントが偽造できない事実 (テスト) に対してスコア付けされます。
環境の DB 状態、git 履歴)。 DOS ゲートが「出荷しました」という 15 の嘘を発見
2 つのモデルにわたる 258 のタスクで誤警報はゼロ。同じ主審が止めた
1 つの共有レコードで 8 件中 6 件のサイレント コリジョン。運命のランをやめる
適切な瞬間により、フリートの最大 11% が 1,634 件中 0 件の勝者で誤って計算されなくなりました
殺された;報酬セットの入場ラベルにより、承認精度が 60% 向上しました →
自己採点コレクターが保管している毒を浄化することで100%。方法論、その２つ
マネーモーメントの数値と、予想される対賭けの正直さの勾配は次のとおりです。
何が証明され、何がまだ賭けなのか。
残りのドキュメントはどこにあるのか
このページには、フック、デモ、および修正された障害が保存されています。より深いものすべて
焦点を当てたページに住んでいます — たどり着いた質問を見つけてジャンプしてください:
端末はありますか？これにより、コマンド 1 つですべてが使い捨てリポジトリ内で実行されます。
それをスキャフォールディングし、実際のコミットを作成し、検証し、それ自体の後にクリーンアップします。
pip install dos-kernel # PyYAML は唯一のランタイム デプロイです
dos クイックスタート # → SHIPPED AUTH AUTH1 … 次に NOT_SHIPPED AUTH AUTH2
1 つは SHIPPED 、もう 1 つは NOT_SHIPPED : 1 つ目は git can back のクレームで、2 つ目は
それは何の根拠もない主張です。そのコントラストが商品です。デモは終了します
すでにエージェントを実行している場所にルーターを使用して接続 — クロード コード / カーソル タブ
( dos init --hooks )、MCP ホスト、CI ステップ、またはフリート - したがって、次の

移動は
ドキュメントの掘り出し物ではなく、一行です。 ( --keep ./demo を追加してリポジトリを保持し、それを確認します。
インストールもしたくないですか？ uvx --from dos-kernel dos クイックスタートは
同じデモは一時的に行われます。何も残らないのです。) 同じものを手作業で 5 回に分けて作成します。
行、 docs/QUICKSTART.md です。
同様に自信のある 2 つの主張、それぞれ 1 つの評決 — 1 つの git can back については SHIPPED、何も着陸しなかった場合には NOT_SHIPPED。すべての文字列は、examples/demo/verify_demo.sh のそのままの出力です。クリックスルー バージョンの場合はローカルでステップ実行します (これは HTML ファイルです。リポジトリを複製してブラウザで開きます。GitHub には実行中のページではなくソースが表示されます)。
最小の実際の勝利: CI ステップまたはディスパッチ ループで、次の行を置き換えます。
dos verify PLAN PHASE でエージェントの「完了」を信頼し、終了時に分岐します
コード (0 出荷 / 1 未出荷)。解析も計画も構成も必要ありません。
CI 統合クックブックで説明
端から端まで。あなたのものと同じような形状のリポジトリでそれを実行するには、次から始めます
10 分でリポジトリをオンボードします。
コミットがより速く蓄積した場合、レビューキューで同じ監視者をポイントします。
誰でも読むことができます。残留レビュー
commit-audit のコミットごとの判定を 3 つのバンド - CLEARED (
diff はその主張を目撃しているので、「それが何をしたのか」を再質問することにほとんど注意を費やしません
言った」）、RESIDUAL（Git は反論できなかった主張 - 人間の 100%）、
ノークレーム休み。このリポジトリ自体の最後の 200 コミットでは、171 件中 170 件がクリアされました
チェック可能なクレーム: それはスキップする再レビューであり、
モデルの信頼スコア。 (CLEARED は、変更の形状がその変更の形状と一致したことを意味します
主張 — コードが正しいということではありません。正確性レビューは引き続き適用されます
すべてのコミット。バンドはより多くの目を求めることしかできませんが、それを減らすことは決してありません。)
次のレベル — 判定を独自のスタックに配線します。 に配線します。
誰にも紹介されずに大量のエージェントを一度に実行する

ええと、それは次のとおりです。
各労働者はそれぞれの成功を報告し、あなたはその報告を信じます。
他に続けることはありますか？チェックされていない問題は静かに山積みですが、ここでは嘘です。
2 人のエージェントが同じファイルを上書きし、スコープが少しクリープし、ワーカーが 1 人
コードベースがある程度機能し、誰も安全にできなくなるまで、ぐるぐるぐるぐる回る
それを変えてください。
問題は、エージェントを起動した後、エージェントに独自の評価を行わせることです。
宿題。 DOS は、欠落している信号 (グラウンド トゥルースからの判定) を提供します。
ループが閉じます。以下は両方の体制下の同じ艦隊です。
フローチャート LR
サブグラフ OPEN["NO REFEREE — あなたはナレーションを信じています"]
TB方向
A1["エージェント:「完了!」"] --> B1[["信じている"]]
A2["エージェント: 「完了!」"] --> B1
A3["エージェント: 「完了!」"] --> B1
B1 --> C1[「静かに腐敗が積み重なる<br/>(嘘・衝突・スピン)」]
C1 --> D1["まあまあ動作します" — 変更できません"]
終わり
サブグラフ CLOSED["DOS ADJUDICATES — あなたは評決に基づいて舵を取ります"]
TB方向
A4["エージェント: '完了!'"] --> V{{"dos verify<br/>git を読み取ります"}}
V -->|git の祖先内| S["出荷済み (出口 0)"]
V -->|どこにも見つからない| N["発送されていません (出口 1)"]
S --> L[「着陸する」]
N --> R["再派遣 / フラグ — 捕獲"]
R -.verdict がループを制御します。 -> A4
終わり
読み込み中
以下に、フリートが実際に生成する失敗を示します。それぞれの真実の隣にあります。
それは労働者の話に静かに矛盾している

[切り捨てられた]

## Original Extract

Catch your AI agents when they lie about what they shipped — verifies claims against git instead of believing the agent. - anthony-chaudhary/dos-kernel

GitHub - anthony-chaudhary/dos-kernel: Catch your AI agents when they lie about what they shipped — verifies claims against git instead of believing the agent. · GitHub
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
anthony-chaudhary
/
dos-kernel
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
709 Commits 709 Commits .claude-plugin .claude-plugin .claude .claude .github .github benchmark benchmark claude-plugin claude-plugin docker docker docs docs examples examples gitlab-ci gitlab-ci go go listings listings packaging/ aur packaging/ aur paper paper scripts scripts spikes spikes src src tests tests verify-action verify-action .dockerignore .dockerignore .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore .pre-commit-hooks.yaml .pre-commit-hooks.yaml AGENTS.md AGENTS.md CITATION.cff CITATION.cff CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile GEMINI.md GEMINI.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md context7.json context7.json dos.toml dos.toml gemini-extension.json gemini-extension.json glama.json glama.json install.ps1 install.ps1 install.py install.py install.sh install.sh llms-full.txt llms-full.txt llms-install.md llms-install.md llms.txt llms.txt pyproject.toml pyproject.toml server.json server.json smithery.yaml smithery.yaml View all files Repository files navigation
DOS — the Dispatch Operating System
Catch your AI agents when they lie about what they shipped.
📊 See it run on real repos: the scoreboard
scores 15 popular AI-built repos (roborev, open-interpreter, crewAI, autogen, …)
— how much agents wrote, which ones, and whether each commit's claim is backed
by its own diff. Score yours: dos commit-audit --sweep --workspace . BASE..HEAD .
The whole pitch in one recording: the agent claims two features shipped; git backs one.
dos verify answers from the commits, the lie exits 1 , and a gate on that
exit code refuses the false "done". Every line is the real CLI's verbatim output —
scripts/build_caught_lie_cast.py re-records it whenever the output changes.
Run a fleet of agents on one repo. The left loop just feels like progress; the right one you can steer.
The only difference is a verdict DOS reads from the real world — here, git — never the agent's word.
An AI agent will tell you it finished. DOS checks the real world instead of
taking its word — and the nearest piece of the real world is your git history.
An agent says it shipped the login endpoint; did it? Run one command,
dos verify , and it answers from the artifacts the work left behind, not from
what the agent typed: a commit backs the claim → SHIPPED , exit 0 ; nothing
landed → NOT_SHIPPED , exit 1 . The agent's story never enters into it. (Git
is just the first witness DOS reads; the file tree, the clock, a CI status, a
test environment's own state are others — anything the agent didn't author.)
dos verify AUTH AUTH1 # → SHIPPED AUTH AUTH1 e62f74d (exit 0)
dos verify AUTH AUTH2 # → NOT_SHIPPED AUTH AUTH2 (exit 1)
That's the smallest version. It scales up, too: point a dozen agents at one
repo — in CI, in a fleet, racing on the same files — and DOS also tells you
which ones are stepping on each other, which one is spinning in circles, and
which claim of "done" is real. Every answer comes from the artifacts (git, the
file tree, the clock), never the narration. It works on a plain git repo with
zero config and gets smarter the more you tell it, and the only thing you ever
install is one small Python package.
⚡ Just add it — two commands, zero decisions. From the repo where your
agent works:
pip install dos-kernel
dos init --hooks auto # finds the agent runtime(s) you already use, wires in the checks
From then on: your agent can't tell you "done" unless the work actually
landed, two agents can't silently overwrite each other's files, and a run
that stalls gets flagged instead of quietly spinning. Nothing about your
workflow changes, and you don't need to learn any of the vocabulary below to
be covered. It prints the one config file it wrote; deleting the dos hook
entries there undoes it. (No runtime detected? It says so and lists the
names to pick from — it never guesses.)
v0.28.0 · 5,600+ tests · CI: Python 3.11–3.13 on Linux + a Windows 3.13
smoke run · the only runtime dependency is PyYAML · MIT .
🧭 Where to go next: the why & evidence (plain-words story, the 20-lines-of-bash answer, what's proven),
wire it into your stack (MCP · hooks · install), the
syscall + CLI reference , or, reading this as an AI agent? , AGENTS.md — build/test/check in three lines. The full map is the router just below.
🔤 Five words the rest of this page leans on. A plan is a named goal
( AUTH ); a phase is one shippable step of it ( AUTH1 ); a lane is the
slice of the file tree one agent may touch; the oracle is the part of DOS
that reads the evidence and rules; a stamp is the mark a shipped phase
leaves in a commit subject ( AUTH1: … ) — the thing the oracle greps for.
That's the whole vocabulary.
A coding agent does work, then tells you how it went. Usually the story is true;
sometimes it's the cheerful "all work completed!" from a worker that shipped
nothing. With one agent you catch that yourself by re-reading its output — a real
tax you already pay. Run twenty at once and that tax stops being payable: nobody
reads everything, each worker grades its own homework, and the unchecked problems
pile up quietly until the codebase sorta works and nobody can safely change it.
DOS is the referee that never reads the story — it reads what happened (the
commit, the file, the clock) and hands you a verdict no narration can move. It
costs about an afternoon, has one runtime dependency, and stays in its lane: it
tells you what happened , never whether the code is good — quality stays with
your tests and reviews. ( The full plain-words version .)
Every number here is scored against a fact the agent can't fake (a test
environment's DB state, git history). A DOS gate caught 15 "I shipped it" lies
in 258 tasks across two models with zero false alarms ; the same referee stopped
6 of 8 silent collisions on one shared record; quitting doomed runs at the
right moment saved ~11% of fleet compute with 0 of 1,634 winners wrongly
killed ; and the reward-set admission label lifted acceptance precision 60% →
100% by purging poison a self-graded collector keeps. The methodology, the two
money-moment figures, and the projected-vs-bet honesty gradient are in
what's proven and what's still a bet .
Where the rest of the docs are
This page keeps the hook, the demo, and the failure it fixes. Everything deeper
lives on a focused page — find the question you arrived with and jump:
Got a terminal? This runs the whole thing in a throwaway repo — one command
scaffolds it, makes a real commit, verifies it, and cleans up after itself:
pip install dos-kernel # PyYAML is the only runtime dep
dos quickstart # → SHIPPED AUTH AUTH1 … then NOT_SHIPPED AUTH AUTH2
One SHIPPED , one NOT_SHIPPED : the first is a claim git can back, the second
is a claim nothing landed for. That contrast is the product. The demo closes
with a router to wherever you already run agents — a Claude Code / Cursor tab
( dos init --hooks ), an MCP host, a CI step, or a fleet — so your next move is
one line, not a docs dig. (Add --keep ./demo to keep the repo and poke at it.
Don't even want the install? uvx --from dos-kernel dos quickstart runs the
same demo ephemerally — nothing left behind.) The same thing by hand, in five
lines, is docs/QUICKSTART.md .
Two equally confident claims, one verdict each — SHIPPED for the one git can back, NOT_SHIPPED for the one nothing landed for. Every string is verbatim output of examples/demo/verify_demo.sh . Step through it locally for the click-through version (it's an HTML file — clone the repo and open it in a browser; GitHub shows its source, not the running page).
The smallest real win: in a CI step or dispatch loop, replace the line that
trusts an agent's "done" with dos verify PLAN PHASE and branch on its exit
code ( 0 shipped / 1 not). No parsing, no plan, no config — the
CI integration cookbook walks it
end-to-end. To run it on a repo shaped like yours, start with
Onboard a repo in 10 minutes .
Point the same witness at a review queue when commits pile up faster than
anyone can read them. Residual review
folds commit-audit 's per-commit verdict into three bands — CLEARED (the
diff witnessed the claim, so spend ~0 attention re-asking "did it do what it
said"), RESIDUAL (a claim git couldn't back — the human's 100%), and the
no-claim rest. On this repo's own last 200 commits it cleared 170 of 171
checkable claims: that's the re-review you skip, proven by git rather than a
model's confidence score. (CLEARED means the change's shape matched its
claim — not that the code is correct; correctness review still applies to
every commit. The band can only ever ask for more eyes, never fewer.)
Next level up — wire the verdict into your own stack: Wire it in .
Run a pile of agents at once with nobody refereeing, and here's how it goes:
each worker reports its own success, and you believe the reports, because what
else is there to go on? The unchecked problems pile up quietly — a lie here,
two agents clobbering the same file there, a little scope creep, one worker
spinning in circles — until the codebase sorta works and nobody can safely
change it.
The trouble is you launched the agents and then let them grade their own
homework. DOS gives you the missing signal — a verdict from ground truth — so
the loop closes. Here is the same fleet under both regimes:
flowchart LR
subgraph OPEN["NO REFEREE — you believe the narration"]
direction TB
A1["agent: 'done!'"] --> B1[["believed"]]
A2["agent: 'done!'"] --> B1
A3["agent: 'done!'"] --> B1
B1 --> C1["silent corruption piles up<br/>(lies · collisions · spin)"]
C1 --> D1["'sorta works' — can't be changed"]
end
subgraph CLOSED["DOS ADJUDICATES — you steer on a verdict"]
direction TB
A4["agent: 'done!'"] --> V{{"dos verify<br/>reads git"}}
V -->|in git ancestry| S["SHIPPED (exit 0)"]
V -->|found nowhere| N["NOT_SHIPPED (exit 1)"]
S --> L["land it"]
N --> R["re-dispatch / flag — caught"]
R -.verdict steers the loop.-> A4
end
Loading
Here are the failures a fleet actually produces, each next to the ground truth
that quietly contradicts the worker's story

[truncated]
