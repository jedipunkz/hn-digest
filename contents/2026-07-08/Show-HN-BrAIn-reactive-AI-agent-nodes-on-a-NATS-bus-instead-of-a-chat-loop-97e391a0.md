---
source: "https://github.com/tibzejoker/brAIn"
hn_url: "https://news.ycombinator.com/item?id=48832740"
title: "Show HN: BrAIn, reactive AI agent nodes on a NATS bus instead of a chat loop"
article_title: "GitHub - tibzejoker/brAIn: A node based agent / orchestration framework · GitHub"
author: "tibzejoker"
captured_at: "2026-07-08T15:10:33Z"
capture_tool: "hn-digest"
hn_id: 48832740
score: 1
comments: 0
posted_at: "2026-07-08T14:50:21Z"
tags:
  - hacker-news
  - translated
---

# Show HN: BrAIn, reactive AI agent nodes on a NATS bus instead of a chat loop

- HN: [48832740](https://news.ycombinator.com/item?id=48832740)
- Source: [github.com](https://github.com/tibzejoker/brAIn)
- Score: 1
- Comments: 0
- Posted: 2026-07-08T14:50:21Z

## Translation

タイトル: HN を表示: BrAIn、チャット ループの代わりに NATS バス上のリアクティブ AI エージェント ノード
記事タイトル: GitHub - tibzejoker/brAIn: ノードベースのエージェント/オーケストレーション フレームワーク · GitHub
説明: ノードベースのエージェント/オーケストレーション フレームワーク。 GitHub でアカウントを作成して、tibzejoker/brAIn の開発に貢献してください。

記事本文:
GitHub - tibzejoker/brAIn: ノードベースのエージェント/オーケストレーション フレームワーク · GitHub
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
ティブゼジョーカー
/
脳
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動

コード 「その他のアクション」メニューを開く フォルダーとファイル
282 コミット 282 コミット .github .github .idea .idea パッケージ パッケージ スクリプト スクリプト テスト テスト .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .npmrc .npmrc AGENTS.md AGENTS.md ARCHITECTURE.md ARCHITECTURE.md CHANGELOG.md CHANGELOG.md Dockerfile Dockerfile ライセンス ライセンス README.md README.md TODO.md TODO.md docker-compose.yml docker-compose.yml eslint.config.mjs eslint.config.mjs package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml run run run.cmd run.cmd sonar-project.properties sonar-project.properties tsconfig.base.json tsconfig.base.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
バス反応型アンビエント インテリジェント ノード
私は Flutter モバイルと AI のエンジニアですが、正直なところ、一度も経験したことがありません。
世の中のエージェント フレームワークに満足しています。それらのほとんどはループまたは cron に要約されます。
タイマーでモデルを突いて、すべてを 1 つのツールに集中させます。
チャット。そこで実際に使いたいものを作ってみました。フル
透明性: このコードを書くために AI に大きく依存しました (これは
この実験も)、結果として得られたものは共有する価値があると思います。
内部的には、実際には長寿命デーモンを備えた NATS パブ/サブ バスです。
ノード、精神的にはチャット フレームワークよりも ROS に近いです。 LLM はただ
ノードが到達できるものの 1 つは、重心ではありません。
それは正しいと思うことがいくつかあります:
待機中のトークンは消費されません。ノードはリアクティブではなく、
予定されています。彼らは駐車したままで、メッセージが届いた場合にのみ LLM を呼び出します。
実は気になる到着。 cron やアイドルポーリングがないため、料金を支払う必要があります
カチカチするのではなく考える。
すべてのノードは、ローカルまたはリモートの独自のインターフェイスを持つことができます。ノード
単なるハンドラーではありません。 UI を出荷でき、その UI はそのまま残ります

使用可能であり、
ノードが別のマシンで実行されている場合でもアクセス可能です。
チャット ボックスではなく、エージェントを監視します。 1回の会話の代わりに
すべてを実行すると、小さな専用ノードのライブ グラフが得られます。
1 つの仕事と独自のビューを持っています。
荷重を分散することができます。同じノードがインプロセスまたはプロセス上で実行されます。
リモートマシンが共有バスに参加しました。作業を分散する
コードに触れずにハードウェアにアクセスできます。
それは洗練された製品ではなく、何かを置き換えようとするものでもありません。
これは、まだ解決されていないと思われる問題に対する私の正直な試みです。もしあれば
これは共感できるので、見てください。
Brain-explainer-web.mp4
brAIn の違いを説明する 2 分間のツアー - リアクティブな多対多ノード、
プリエンプションによる重要性、ライブ ダッシュボード、ノードごとの任意のモデル、
分散 + MCP、そしてそれを使って構築できるもの。
▶︎ YouTubeでもご覧いただけます。
voice.chat.with.intent.2026-07-08.at.16.02.16.mp4
サウンドオン。ビデオは仮想カメラとマイクとして再生され、
認識スタック全体がそれをライブ入力として扱います。声
二人の発言を文字に起こして日記にし、視線は誰が話しているのかを追跡する
どこを探し、意図が 2 つをペアにするか: それぞれに話しかける人々
他のものは文脈のままですが、話者がカメラに向かって話しかけた瞬間
脳が目覚め、応答し、ココロの tts ノードが発話します。
大声で返事をする。ウェイクワードはなく、立ち聞きされたおしゃべりに費用がかかることはありません
LLMコール。
ネットワーク全体のライブグラフ
ダッシュボードはネットワークです。つまり、すべてのノード、ノード間を流れるメッセージです。
それらと、ドラッグして編集する配線。単一のチャットではなく、それぞれのみ
エージェントは独自の仕事をしています。
2 台のマシンが 1 つのバスと 1 つのキャンバスを共有します。ピアショーで生成されたノード
ローカルのものの隣にあるため、ハードウェア全体に負荷を分散できます。
配線では、ノードが実際にどこに存在するかは関係ありません。
ノードは単なる機能ではなく、独自のインターフェイスを提供できます。ここで、
sm

すべての「ブレインペット」ノードは専用の UI を備え、再利用可能で同じに到達します
ローカルで実行されるかリモートピアで実行されるかによって異なります。
どこからでもノードにアクセス
Telegram ブリッジを介して再生される三目並べノード。同じノード
ロジックは外部チャネル経由で公開されるため、インターフェイスは再利用可能です
バスの向こう側でも、マシンから離れていても。
電話.孤独.mp4
brAIn-mobile は電話をノードに変えます。そのセンサー (光、動き、
バッテリー）他の出版社と同じようにバスにストリーミングし、脳を動かします。
デバイスに物理的に起こったことに反応します。ここは強いです
どこに置かれたのかについての意見。応答はループバックします。
携帯電話自体の画面とスピーカーなので、認識と反応の両方がライブです
端で。
brAIn と他のツールとの関係
brAIn は、あなたが知っているほとんどのエージェント ツールと競合するものではありません。
隣接する問題を解決し、適切なものは形状によって異なります
希望するエージェントの。
LangGraph / Vercel AI SDK / Mastra : チャット型エージェント用
(ユーザータイプ → LLM が考える → ツール → LLM が応答する)、これらは次のとおりです。
brAIn のチャット サポートよりも優れており、より詳細です。いつでも彼らに手を伸ばしてください
エージェントは基本的に会話型インターフェイスです。
AutoGen / CrewAI : 役割を伴うマルチエージェントの会話。使用する
複数の LLM ペルソナが議論したりコラボレーションしたりする場合にこれらを使用します。
単一の対話内で。
ROS 2 : アーキテクチャ的に最も近いもの。パブ/サブバス、デーモン
ノード、多言語、DDS を介したクロスマシン。 brAIn は多くのことを共有しています
ROS を使用したメンタル モデルの構築。違いはドメイン固有です。
brAIn は、LLM 制約 (トークン バジェット、中止可能) を中心に構築されています。
推論、ツール呼び出しループ、MCP) は、DDS ではなく NATS 上で実行されます。
(リアルタイム保証が必要ない場合は導入が簡単です)、および
間ではなく進行中のレベルでプリエンプトします。
コールバック。
インジェスト / Trigger.dev / 時間 : du

素晴らしいワークフロー。もしあなたの
エージェントは、再試行とバックオフを備えた有限形状の DAG です。
プロダクショングレードのオプション。 brAIn のノードはオープンエンドのデーモンです。
両者は異なるメンタルモデルに基づいて動作します。
n8n / Flowise / Langflow / Dify / Node-RED : ビジュアル
ノードベースのオーサリング。 brAIn ノードはコードで記述されます (a
TypeScript ハンドラーと config.json 、または任意の HTTP/WS サービス
トランスポート経由: 「ウェブ」)。ダッシュボードは観察第一です。
Claude Cowork / OpenAI スケジュールされたタスク / cron 駆動のエージェント :
時間によってトリガーされるプロンプト。 brAIn はイベントによってトリガーされます。あなたのエージェントの場合
ケイデンスは「毎週月曜日の朝に X を要約する」、スケジューラは
右のプリミティブ。
brAIn のアーキテクチャ上の特徴を凝縮して比較すると、
コンテキストとしては次のようになります。
多対多のイベント内で動作する自律エージェント用のランタイム
世界。
ノードは存続期間の長いデーモンです。それぞれが複数の入力をサブスクライブします
ストリーム (チャット メッセージ、センサー イベント、Webhook、内部バス)
トラフィック、公開できるものなら何でも）、関連するものがあったときに反応します
が表示され、同時に多くの出力にパブリッシュできます。ないよ
単一のトリガー チャネル: エージェントが監視し、いつトリガーするかを決定します。
行動する。
ノードは、より高い重要度の場合、飛行中にプリエンプトされることもあります。
遅い操作 (LLM 呼び出し、ツール) 中にメッセージが到着した
呼び出し、CLI エージェント）、ランナーは進行中の内容を中止し、
表示された新しいコンテキストでハンドラーを再実行します。
ctx.preemptionContext 。同じノード構成をプロセス内で実行できます。
WebSocket、または共有 NATS に参加しているリモート ブレイン エージェント上
バス。
これを使用して構築できるいくつかの形状:
カメラとマイクのみを監視するアンビエントルームエージェント
誰かが話しているときにそれを見ているときに話します（フラッグシップ）
デモ: 音声 + 視線 + 意図 + 脳)。
スレッド内に存在する Slack チャネルのリスナーが、
メッセージ全体のコンテキスト、および

要約したり返信したりする
会話が一時停止します。
Grafana アラートをサブスクライブしたモニタリング エージェント +
oncall-rotation イベント + 最近のデプロイにより、
相関関係がしきい値を超えた場合の仮説。
温度、動き、カレンダー、
環境をいつ変更するかを決定する時刻。
デーモン モデル: ノードは反復にわたって存続します。枠組み
アイドル時に自動駐車し、次の時間にバスがウェイクアップします。
購読したメッセージ。状態は実行後も維持されます。
多対多 I/O : 各ノードは N 個のトピック パターンをサブスクライブします
(ワイルドカードを使用) 多くのユーザーに公開します。
プリエンプションによる重要性: すべてのメッセージには
批判性。重要度の高いメッセージがハンドラーの途中で到着する
実行中の反復を中止し、再実行をトリガーします。
ctx.preemptionContext.{中断中のメッセージ、前のメッセージ}
利用可能です。
分散ランタイム: 同じノード構成がプロセス内またはプロセス上で実行されます。
リモートのブレインエージェント。ライフサイクル (停止/開始) とリードバック
(ログ/メールボックス/DLQ) はマシン間で透過的に動作します。
NATS。
MCP ネイティブ : mcp-config (マネージャー) + mcp-server (1 つあたり 1 つ)
アップストリーム）MCP サーバーのツールをバスにブリッジします。
エージェントがファイルシステム、git、Slack、Linear、Notion、に到達します。
Sentry … 他のノードと同じです。
package/core/src/bus : NatsBusService は IBusService を実装します
NATS ブローカーの上にあります。フレームワークは組み込みの
nats-server (バンドルされている Go バイナリ、ポストインストールによってダウンロードされる)
フック) 空きローカルホスト ポート上。リモートのブレインエージェントプロセス
同じブローカーに接続し、バスを共有します。セット
BRAIN_NATS_URL は組み込みブローカーをスキップし、外部ブローカーに参加します
複数のホスト間で実行する場合は一般的に、代わりに 1 つが使用されます。
BusService (メモリ内) は引き続きエクスポートされますが、テストとしてのみエクスポートされます
治具;実稼働コードのパスは常に NATS を経由します。
ノ

独自の制限 (スケーリング): 各インスタンスは全体をサブスクライブします。
階層 ( <prefix>.> ) と matchTopic を使用してローカルにフィルターして保持します
brAIn のワイルドカード セマンティクス。したがって、すべてのインスタンスはすべてのバスを認識します
トラフィックを処理し、メモリ内でフィルタリングします。小規模/LAN 規模では問題ありませんが、
インスタンス/メッセージ数が増加するにつれて最初のボトルネックが発生します。対象となる修正は
実際のサブジェクトごとの NATS サブスクリプションの数が狭くなります。
ワイルドカード トピックの一致 (alerts.*)。
構成可能な max_size と
最新/最低優先度の保持ポリシー。落として
容量が公開されるため、ダッシュボードでメールボックスごとに表示できるようになります
背圧。
因果関係のトレース: すべてのメッセージには、trace_id +parent_id が含まれます。
NATS エンコード/デコードを存続し、クエリ可能
/network/traces/:id を取得します。各トレースは新しいものとして再生できます
POST /network/traces/:id/replay を介した排出: 新しい ID、
書き換えられた親チェーン、元のtrace_idは次のように保持されます
デバッグ用のmetadata.replayed_from。
履歴スライディング ウィンドウ (デフォルトでは 10,000 メッセージ)。
Packages/core/src/runner 、ノードのタグから選択されます（ノードのタグではありません）
name): llm タグを持つノードは LLMRunner を取得します。
それ以外の場合は ServiceRunner を取得します (Web トランスポートは両方をオーバーライドします)。の
以下の括弧内のノード名はノードの単なる例です。

[切り捨てられた]

## Original Extract

A node based agent / orchestration framework. Contribute to tibzejoker/brAIn development by creating an account on GitHub.

GitHub - tibzejoker/brAIn: A node based agent / orchestration framework · GitHub
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
tibzejoker
/
brAIn
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
282 Commits 282 Commits .github .github .idea .idea packages packages scripts scripts tests tests .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .npmrc .npmrc AGENTS.md AGENTS.md ARCHITECTURE.md ARCHITECTURE.md CHANGELOG.md CHANGELOG.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md TODO.md TODO.md docker-compose.yml docker-compose.yml eslint.config.mjs eslint.config.mjs package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml run run run.cmd run.cmd sonar-project.properties sonar-project.properties tsconfig.base.json tsconfig.base.json vitest.config.ts vitest.config.ts View all files Repository files navigation
Bus-Reactive Ambient Intelligent Nodes
I'm a Flutter mobile and AI engineer, and honestly I've never been
satisfied with the agent frameworks out there. Most of them boil down to a loop or a cron
poking a model on a timer, with everything funnelled through a single
chat. So I tried to build the thing I actually wanted to use. Full
transparency: I leaned on AI heavily to write this code (that's part of
the experiment too), and I think what came out is worth sharing.
Under the hood it's really a NATS pub/sub bus with long-lived daemon
nodes, closer in spirit to ROS than to a chat framework. The LLM is just
one thing a node can reach for, not the center of gravity.
A few things I think it gets right:
It doesn't burn tokens waiting. Nodes are reactive , not
scheduled. They stay parked and only call the LLM when a message they
actually care about arrives. No cron, no idle polling, so you pay for
thinking instead of ticking.
Every node can have its own interface, local or remote. A node
isn't just a handler; it can ship a UI, and that UI stays reusable and
reachable even when the node runs on another machine.
You watch your agents, not a chat box. Instead of one conversation
doing everything, you get a live graph of small, dedicated nodes, each
with one job and its own view.
You can spread the load. The same node runs in-process or on a
remote machine joined to the shared bus. Distribute work across
hardware without touching the code.
It's not a polished product and it's not trying to replace anything.
It's my honest attempt at a problem I don't think is solved yet. If any
of this resonates, take a look.
brain-explainer-web.mp4
A 2-minute tour of what makes brAIn different — reactive many-to-many nodes,
criticality with preemption, the live dashboard, any-model-per-node,
distributed + MCP, and what you can build with it.
▶︎ Also on YouTube .
voice.chat.with.intent.2026-07-08.at.16.02.16.mp4
Sound on. A video is replayed as a virtual camera and microphone, and
the whole perception stack treats it as live input. voice
transcribes and diarizes the two speakers, gaze tracks who is
looking where, and intent pairs the two: people talking to each
other stays context, but the moment a speaker addresses the camera
the brain wakes up, answers, and the Kokoro tts node speaks the
reply out loud. No wake word, and overheard chatter never costs an
LLM call.
A live graph of your whole network
The dashboard is the network: every node, the messages flowing between
them, and the wiring you edit by dragging. No single chat, just each
agent doing its own job.
Two machines share one bus and one canvas. Nodes spawned on a peer show
up next to the local ones, so you can spread the load across hardware.
The wiring doesn't care where a node actually lives.
A node is more than a function: it can serve its own interface. Here a
small "brainpet" node with a dedicated UI, reusable and reached the same
way whether it runs locally or on a remote peer.
Reach your nodes from anywhere
A tic-tac-toe node played through the Telegram bridge. The same node
logic is exposed over an external channel, so interfaces stay reusable
across the bus, even off your machine.
phone.loneliness.mp4
brAIn-mobile turns a phone into a node: its sensors (light, motion,
battery) stream onto the bus like any other publisher, and the brain
reacts to what physically happens to the device. Here it has strong
opinions about where it just got put down. The reply loops back to the
phone's own screen and speaker, so perception and reaction both live
at the edge.
How brAIn relates to other tools
brAIn isn't a competitor to most of the agentic tools you may know;
they solve adjacent problems and the right one depends on the shape
of the agent you want.
LangGraph / Vercel AI SDK / Mastra : for chat-shaped agents
(user types → LLM thinks → tools → LLM responds), these are
excellent and deeper than brAIn's chat support. Reach for them when
the agent is fundamentally a conversational interface.
AutoGen / CrewAI : multi-agent conversations with roles. Use
these when you want several LLM personas debating or collaborating
within a single dialogue.
ROS 2 : the closest architectural cousin. Pub/sub bus, daemon
nodes, multi-language, cross-machine over DDS. brAIn shares a lot
of its mental model with ROS. The differences are domain-specific:
brAIn is built around LLM constraints (token budgets, abortable
inference, tool-call loops, MCP), runs over NATS rather than DDS
(easier to deploy when you don't need real-time guarantees), and
preempts at the work-in-progress level rather than between
callbacks.
Inngest / Trigger.dev / Temporal : durable workflows. If your
agent is a finite-shape DAG with retries and backoff, those are
production-grade options. brAIn's nodes are open-ended daemons;
the two run on different mental models.
n8n / Flowise / Langflow / Dify / Node-RED : visual
node-based authoring. brAIn nodes are written in code (a
TypeScript handler plus a config.json , or any HTTP/WS service
via transport: "web" ). The dashboard is observation-first.
Claude Cowork / OpenAI scheduled tasks / cron-driven agents :
time-triggered prompts. brAIn is event-triggered; if your agent's
cadence is "every Monday morning summarise X", a scheduler is the
right primitive.
A condensed comparison of the architectural traits brAIn happens to
have, for context:
A runtime for autonomous agents that live in a many-to-many event
world .
Nodes are long-lived daemons. Each one subscribes to several input
streams (chat messages, sensor events, webhooks, internal bus
traffic, anything you can publish), reacts when something relevant
shows up, and can publish to as many outputs in parallel. There's no
single triggering channel: the agent watches and decides when to
act.
A node may also be preempted mid-flight: when a higher-criticality
message lands during a slow operation (an LLM call, a tool
invocation, a CLI agent), the runner aborts what's in progress and
re-runs the handler with the new context surfaced in
ctx.preemptionContext . Same node config can run in-process, behind
a WebSocket, or on a remote brain-agent joined to a shared NATS
bus.
A few shapes you can build with this:
An ambient room agent that watches camera + mic and only
speaks when someone is looking at it while talking (the flagship
demo: voice + gaze + intent + brain).
A Slack-channel listener that lives in a thread, picks up
context across messages, and summarises or replies when the
conversation pauses.
A monitoring agent subscribed to Grafana alerts +
oncall-rotation events + recent deploys, that surfaces a
hypothesis when a correlation crosses a threshold.
An IoT controller that fuses temperature, motion, calendar,
and time-of-day to decide when to change the environment.
Daemon model : nodes live across iterations; the framework
auto-parks them when idle and the bus wakes them on the next
subscribed message. State persists across runs.
Many-to-many I/O : each node subscribes to N topic patterns
(with wildcards) and publishes to as many.
Criticality with preemption : every message carries a
criticality. A higher-criticality message arriving mid-handler
aborts the running iteration and triggers a re-run with
ctx.preemptionContext.{interrupting_message, previous_messages}
available.
Distributed runtime : same node config runs in-process or on
a remote brain-agent . Lifecycle (stop / start) and read-back
(logs / mailbox / DLQ) work transparently across machines over
NATS.
MCP-native : mcp-config (manager) + mcp-server (one per
upstream) bridge any MCP server's tools onto the bus, so the
agent reaches into filesystem, git, Slack, Linear, Notion,
Sentry … as it would call any other node.
packages/core/src/bus : NatsBusService implements IBusService
on top of a NATS broker. The framework boots an embedded
nats-server (the bundled Go binary, downloaded by the postinstall
hook) on a free localhost port; remote brain-agent processes
connect to that same broker and share the bus. Set
BRAIN_NATS_URL to skip the embedded broker and join an external
one instead, typical when running across multiple hosts.
BusService (in-memory) is still exported but only as a test
fixture; the production code path always goes through NATS.
Known limitation (scaling): each instance subscribes to the whole
hierarchy ( <prefix>.> ) and filters locally with matchTopic to keep
brAIn's wildcard semantics. Every instance therefore sees all bus
traffic and filters it in memory, fine at small/LAN scale, but the
first bottleneck as instance/message counts grow. The targeted fix is
narrower NATS subscriptions per real subject.
Wildcard topic matching ( alerts.* ).
Per-subscription mailbox with configurable max_size and a
latest / lowest_priority retention policy. dropped and
capacity are exposed so the dashboard can show per-mailbox
backpressure.
Causal traces : every message carries trace_id + parent_id .
Survives NATS encode/decode, queryable via
GET /network/traces/:id . Each trace can be replayed as fresh
emissions through POST /network/traces/:id/replay : fresh ids,
rewritten parent chain, original trace_id carried as
metadata.replayed_from for debugging.
History sliding window (10k messages by default).
packages/core/src/runner , picked from a node's tags (not its
name): a node carrying the llm tag gets the LLMRunner , anything
else gets the ServiceRunner (a web transport overrides both). The
node names in parentheses below are just examples of nodes th

[truncated]
