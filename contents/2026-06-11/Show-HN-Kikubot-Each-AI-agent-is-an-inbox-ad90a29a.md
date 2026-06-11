---
source: "https://github.com/mxaiorg/kikubot"
hn_url: "https://news.ycombinator.com/item?id=48492034"
title: "Show HN: Kikubot – Each AI agent is an inbox"
article_title: "GitHub - mxaiorg/kikubot: Email-based AI agent network · GitHub"
author: "asp68"
captured_at: "2026-06-11T16:26:59Z"
capture_tool: "hn-digest"
hn_id: 48492034
score: 2
comments: 0
posted_at: "2026-06-11T15:50:54Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Kikubot – Each AI agent is an inbox

- HN: [48492034](https://news.ycombinator.com/item?id=48492034)
- Source: [github.com](https://github.com/mxaiorg/kikubot)
- Score: 2
- Comments: 0
- Posted: 2026-06-11T15:50:54Z

## Translation

タイトル: 表示 HN: キクボット – 各 AI エージェントは受信箱です
記事タイトル: GitHub - mxaiorg/kikubot: メールベースの AI エージェント ネットワーク · GitHub
説明: 電子メールベースの AI エージェント ネットワーク。 GitHub でアカウントを作成して、mxaiorg/kikubot の開発に貢献してください。
HN テキスト: こんにちは。私は、メッセージ バス (メッセージ キュー、ベクター ストア、オーケストレーターなし) として電子メールを使用することを主な理念とする OSS エージェント フレームワークを投稿しました。基本的に、すべてのエージェントは電子メール アドレスです。このプロジェクトの背後にある主な理由は、企業による AI エージェントの導入と展開を促進することです。 # どのように動作するか (どのように見えるか)
ユーザーは特定の電子メール アドレス (例: kiku@agent.acme.com) に電子メールを送信します。キクは、「添付記事を取得してソーシャル メディアへの投稿を準備する」などの要求されたタスクを実行し、それを電子メール経由で 1 つ以上の内部エージェントにルーティングする、指定されたコーディネーター エージェントとして作成されました。キクは、「彼女」がクラスター内のエージェントの「名簿」を管理しているため、どのエージェントにルーティングするかを知っています。この名簿には、各エージェントの能力が記載されています。エージェントは、送信されたタスクの一部を処理し、結果をコーディネーターに返し、コーディネーターはそれをユーザーに返します。重要な設計面の 1 つは、プロジェクトが電子メール スレッドを状態メモリとして使用していることです。エージェントの LLM が呼び出されると、スレッドが履歴コンテキストとして渡されます。キクボット フレームワークの主な利点: - ユーザーは電子メールを介して AI エージェントと連携するため、組織への展開が容易になります。 - エンドユーザーによるインストールは必要なく、電子メール アドレスだけでタスクを送信することもできます。
- エージェントごとのユーザー アクセス制御
- エージェントごとのコンテナ化
- 拡張性の高い機能 - エージェント チームは他のチームを「ネスト」できます。
- エージェントごとの LLM 選択によるコスト抑制
- 標準の電子メール ツールを介したエージェント内の「会話」の高い可視性

配備:
私たちは、導入とカスタマイズを容易にすることに取り組んでいます。弊社では、セットアップをより簡単にするために大いに役立つと信じている Configurator ツールを用意しています。ローカル電子メール サーバー、Web メール クライアント、Kikubot エージェントを含む完全なデモをマシン上で起動する ./demo.sh スクリプトを追加しました。 ./demo.sh を実行するだけです。 OpenClawとの比較:
世の中には 10 億ものエージェント フレームワークが存在します。以下は、Kikubot がどこに適しているかを理解するために OpenClaw と並べて比較した表です。キクボット |オープンクロー
----------+--------------------------+--------------------------
メッセージバス |電子メール (IMAP/SMTP) |チャットアプリ（テレグラム、
| |スラック、ディスコード、
| | WhatsAppなど）
----------+--------------------------+--------------------------
建築 |マルチエージェント設計 |単一のゲートウェイ
| - コーディネーター + |チャネル/脳/身体
|専門サブエージェント |層分離
----------+--------------------------+--------------------------
目標規模 |組織 / |個人の生産性
|企業 (数百人 | / 個人アシスタント)
|エージェントの数、 |
|部門) |
----------+--------------------------+--------------------------
エージェント トポロジ |分散化 - それぞれ |ローカルファーストゲートウェイ
|独立したエージェント |あなたのプロセス
| | 内の任意の場所にデプロイされます。マシンまたは VPS
|組織 |
----------+--------------------------+--------------------------
エンドユーザーによるインストール |なし - ユーザーのみ |インストールが必要です
|アドレスをメールで送信する |そして、
| |ゲートウェイランタイム
----------+--------------------------+--------------------------
第一言語 |行く | Node.js
それにしても、長文です。プロジェクトをチェックしてください。面白いと思ったらぜひご参加ください！さまざまな種類の組織に、Kikubot の非常に興味深いアプリケーションがすでにいくつかあります。

記事本文:
GitHub - mxaiorg/kikubot: 電子メールベースの AI エージェント ネットワーク · GitHub
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
mxaiorg
/
キクボット
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード さらにアクトを開く

イオンメニュー フォルダーとファイル
81 コミット 81 コミット アセット アセット cmd/ kikubot cmd/ kikubot configs configs docs docs external 内部スクリプト scripts services services .gitignore .gitignore CLAUDE.md CLAUDE.md CONFIGURATION.md CONFIGURATION.md COTRIBUTING.md COTRIBUTING.md Dockerfile Dockerfile LICENSE ライセンスREADME-Observability.md README-Observability.md README.md README.md Demon.sh Demon.sh docker-compose-demo.yml docker-compose-demo.yml docker-compose-example.yml docker-compose-example.yml go.mod go.mod go.sum go.sum すべてのファイルを表示 リポジトリ ファイルのナビゲーション
電子メール主導のマルチエージェント フレームワーク。各エージェントは受信箱です。
Kikubot は、電子メール アカウントを自律エージェントに変えます。実行中の各コンテナは 1 つの IMAP メールボックスをポーリングし、構成可能なツール セットを使用して LLM エージェント ループを通じてすべての新しい電子メールを実行し、SMTP 経由で応答します。エージェントは相互に電子メールを送信することで共同作業を行うため、一般的な展開は、複数のエージェント (コーディネーターと少数のスペシャリスト) が 1 つのメール サーバーを共有するように見えます。
なぜ電子メールを送信するのでしょうか?これはユニバーサルな非同期メッセージ バスです。人間はすでにこれを使用しており、すべてのシステムがメッセージ バスに送信でき、スレッドには独自の履歴が保持され (参照: / In-Reply-To: )、アカウントによりエージェントごとの ID、ACL、耐久性が無料で提供されます。
AI ユーザー インターフェイスとしての電子メール。最もよく使用され理解されているテクノロジである電子メールを介して AI を組織に導入します。トレーニングやソフトウェアのインストールは必要ありません。必要なのは電子メール アドレスだけです。
高い拡張性。エージェントのクラスター。各エージェントを独自のクラスターにすることができ、理論的には大規模なスケーラビリティが実現します。
可観測性。エージェントは標準の電子メールを介して相互に通信します。エージェント アカウントにアクセスして、内部会話履歴を確認します。
コストの抑制。各エージェントは異なる LLM に設定できます。エージェントの役割とツールに最適な LLM を選択する

など。
より高いパフォーマンス。エージェントは専門化することができます。機能はエージェント ネットワーク全体に分散されます。各エージェントは自分の専門分野に重点を置いています。
セキュリティの強化。ユーザーのマシン上で AI エージェントが実行されるリスクはありません。 Kokubot Agent はコンテナ内で実行され、分離とセキュリティを提供します。エージェントは、スコープ指定された API キーを介してツールと統合にアクセスします。エージェントへのアクセスは、ACL (ホワイトリストまたはブラックリストのドメイン、または電子メール アドレス) を介して簡単に制御できます。
回復力。キクボット ネットワークは、世界で最も復元力の高いテクノロジーの 1 つである電子メール上で実行されます。
スレッドごとのメモリ。各電子メールのスレッドは長時間にわたる会話です。エージェントの履歴は、スレッドのルート Message-Id をキーとした JSON として保持されます。
プラグイン可能なツール。組み込みツールは、メッセージング、ステータス レポート、スヌーズ、メールボックス検索をカバーします。オプションのツールには、Salesforce、WordPress、Buffer、Box、Helpjuice、Tavily Web 検索、Apache Tika file-to-text、および任意のローカル/HTTP MCP サーバーが含まれます。
プラグ可能な LLM。 Anthropic API (デフォルト、プロンプト キャッシュあり) または OpenRouter (バックアップ モデル フォールバックあり)。
知識ベース。システム プロンプトに追加されるエージェントごとの共有マークダウン ファイル - 再構築せずにライブで編集可能、ホットリロード可能。
マルチエージェントの調整。エージェントは、message_tool コア ツールを介して相互に通信します。コーディネーター エージェントは、保留中の作業を委任、ファンアウト、およびスヌーズできます。
定期的なタスク。エージェントは、特定の時間または間隔でタスクを実行するようにスケジュールを設定できます。
自動返信/バウンスの安全性。 DSN と不在時の応答は、無限委任ループを防ぐために LLM をバイパスします。
1 人のエージェントから数千のエージェントまで
このリポジトリを使用して、同じマシン上に 1 つ以上のエージェント コンテナを生成できます。各コンテナは単一のエージェントを実行します。このリポジトリを複数のマシンにデプロイし、組織全体にエージェントを生成することもできます。唯一の要件は、

コーディネーター エージェントは電子メールを介して相互に連絡できること。
コーディネーター エージェントはチームに編成でき、各チームに複数のエージェントを含めることができます。コーディネーター エージェントのチーム メンバー自体がコーディネーターになることができます。組織が部門に構造化され、各部門が複数の部門を代表し、さらに複数のチームが代表されるのと同じように、エージェントのネットワークを構築することもできます。各コーディネーターは、直接連携するエージェントのサブセットのみを知る必要があります。理論的には、Kikubot の導入は数十万のエージェントまで拡張できます。
あなたの町の天気についてアルファに尋ねてください。クエリを alpha@labtest.mxhero.com に電子メールで送信してください。
mxHERO Labs は、Kikubot インスタンスを 1 台のマシンにデプロイしました。このインスタンスは、Alpha と Weatherman の 2 つのエージェントで構成されています。コーディネーターはアルファです。どこかの都市の天気を尋ねる電子メールをアルファに送信すると、アルファは気象予報士に天気予報を尋ねます。応答を受信すると、アルファから返信が送信されます。
この小さな Kikubot デモは、複数のエージェントが電子メールを介してどのように連携できるかを示しています。
メールアカウントやAPIを使用せずに、独自のKikubotを1分以内に実行したい
キーが必要ですか?デモでは、使い捨てメール サーバー、Web メール UI、および 1 つの機能を起動します。
エージェント — すべて Docker 内にあり、すべて自分のマシン上にあり、インターネットには何も公開されていません。
git clone https://github.com/mxaiorg/kikubot && cd kikubot
./demo.sh
次に:
http://localhost:8000 で Web メール UI を開きます。
human@demo.local としてログインします (任意のパスワード - デモ メール サーバーの認証は無効になっています)
新しいメールを作成して kiku@demo.local に送信します。
30 秒ほど待って、受信箱を更新してください — キクが返信します。 🎉
すぐに使える (API キーなし) キクは、短い「私は生きています — のキーを追加してください」と応答します。
本当の答え」に気づくので、往復の時間をゼロコストで得ることができます。ロックを解除するには
r

eal エージェントの応答を確認するには、キーを configs/demo/secrets.env にドロップして実行します。
./demo.sh をもう一度 — ANTHROPIC_API_KEY または OPENROUTER_API_KEY を貼り付けます
スクリプトは、一致するプロバイダーを自動的に選択します (構成の編集は必要ありません)。
./demo.sh down ですべてを停止します。
内部 ( docker-compose-demo.yml ): GreenMail
SMTP/IMAP を提供し、Roundcube は作成ウィンドウであり、
エージェントは本番環境と同じイメージです - のデモ設定をポイントしているだけです
configs/demo/ 。
このプロジェクトは、mxHERO Labs の研究に基づいています。詳細については、ブログ投稿を参照してください。
ランタイム、構成、メモリ、ツール、統合に関する開発者向けの視覚的なマップについては、 docs/architecture.md を参照してください。
┌───────┐ ┌─────────┐
│ ユーザー │─┐ │ コーディネーター │ ◀─┐
└───────┘ │ │ (菊受信箱) │ │
▼ ━───┬─────┘ │
┌─────┐ │ │ メール
│ IMAP / │ ▼ デリゲート │ スレッド
│ SMTP │ ┌─────┐ ┌─────┐ ┌─────┐
━━━┬─────┘ │ ベータ │ │ ガンマ │ │ デルタ │
│ │ (CRM) │ │(ソーシャル) │ │ (ウェブ) │
━───┴─────┘ └─────┘ └─────┘
│ │ │
Salesforce バッファー WordPress
mxMCP タヴィリー ヘルプジュース
ボックス、ティカ
各エージェント コンテナは、共有 configs/agents.yaml (名簿 + 共通のデフォルト + エージェントごとのオーバーライド) および共有 configs/secrets.env (API キー + エージェントごとのメールボックス パスワード) によってパラメータ化された同一の Go バイナリを実行します。コンテナは AGENT_EMAIL から ID を取得します。

docker-compose によって挿入された環境変数。
LLM API キー — ANTHROPIC_API_KEY および/または OPENROUTER_API_KEY 。
エージェントごとに 1 つのメールボックスを持つ IMAP + SMTP サーバー。任意のプロバイダーを使用できます。このリポジトリには、必要に応じて、services/dms/ に自己ホスト可能な docker-mailserver サイドカーが含まれています。
(オプション) 有効にした統合 (Salesforce、WordPress、Buffer、Helpjuice、Box、Xero、Tavily、mxMCP) のツール認証情報。
ダッシュボード構成ツールはスクリプト ディレクトリにあります。これは、展開を構成できる Web アプリです。エージェントを定義し、必要に応じて含まれる電子メール サーバーを構成します。
go run ./scripts/configurator # 127.0.0.1:50042 で動作します
詳細については、scripts/configurator/README.md を参照してください。少し古いですが、Configurator ビデオ チュートリアルもあります。
LLM ガイド付き展開の構成ガイド。使用するには、Claude Code などの LLM コーディング エージェントを開き、次のプロンプトを表示するだけです。
CONFIGURATION.md ファイルを読み、その指示に従ってください。
> キクボットを設定します。
または、独自の言語で設定したい場合は、次のようにします。
CONFIGURATION.md ファイルを読み、その指示に従ってください。
キクボットを設定します。日本語で私とコミュニケーションしてください。
git clone https://github.com/mxaiorg/kikubot
CDキクボット
# 1. 名簿と共通のデフォルトを設定します。
cp configs/agents-example.yaml configs/agents.yaml
cp configs/secrets-example.env configs/secrets.env
# configs/agents.yaml を編集: 共通: デフォルト (メールサーバー、
# プロンプト、予算) およびエージェント: エントリ (アイデンティティ、役割、ツール、オプション)
# エージェントごとのオーバーライド)。次に、configs/secrets.env を編集します。
# ANTHROPIC_API_KEY (または OPENROUTER_API_KEY) と 1 つ
# エージェントごとの <UPPERCASED_LOCAL_PART>_EMAIL_PASSWORD、およびツールの資格情報。
# 2. (オプション) ナレッジ ファイルを configs/knowledge/<agent>/*.md にドロップします。
# 3. docker-compose を編集する

.yml を名簿に一致させます。
cp docker-compose-example.yml docker-compose.yml
# - エージェントごとに 1 つのサービス。各サービスは `environment:` に AGENT_EMAIL を設定します。
# そして env_file を configs/secrets.env にポイントします。
# - ボリューム マウント: ./data/<stem>:/app/data (stem = 小文字のローカル部分)。
#4. 検証します。
go run scripts/kikudoctor/ * .go
#5. 起動。
docker 構成 -d --build
ホワイトリストに登録されたアドレスからエージェントに電子メールを送信し、受信箱に返信が届くのを確認します。
ログに記録されたエージェント間の会話を確認するには:
docker compose ログ -f
Kikubot は電子メール ベースのシステムであるため、任意の電子メール クライアントを使用してエージェントの内部および外部の会話を追跡できます。可観測性を参照
デプロイメント構成 — configs/agents.yaml
configs/agents.yaml は、非機密デプロイメント構成の信頼できる唯一の情報源です。これには 2 つのセクションがあります。
共通：
email_server : mail.agents.example.com:993
smtp_server : mail.agents.example.com:587
email_insecure_tls : false
max_history_chars : 200000
max_tokens : 6000
エージェントタイムアウト: 300
最大回転数 : 20
システムプロンプト: |
あなたは役に立つエージェントです...
{{同僚}}
コーディネーターシステムプロンプト: |
あなたは有能なコーディネーター エージェントです...
{{同僚}}
エージェント:
- 名前 : キク
メールアドレス：kiku@agents.example.com
役割 : コーディネーター
description : ユーザーと通信します。他のエージェントを調整します。
ツール: [レポート、スヌーズ、tavily_mcp]
# common: フィールドはここでオーバーライドできます。
llm_provider : オープンルーター
llm_model : anthropic/claude-sonnet-4.6
最大回転数 : 40
ホワイトリスト: [example.com、agents.example.com]
- 名前 : ベータ
電子メール : beta@agents.example.com
役割 : CRM、電子メール アーキビスト
説明 : Salesforce とカンプへのアクセスを管理します

[切り捨てられた]

## Original Extract

Email-based AI agent network. Contribute to mxaiorg/kikubot development by creating an account on GitHub.

Hi All, I’ve posted an OSS Agent framework with its main philosophy being the use of email as the message bus (no message queue, vector store, orchestrator). Essentially every agent is an email address. The main reason behind this project is to facilitate company adoption & deployment of AI agents. # How it behaves (what it looks like)
A user sends an email to a specific email address, e.g., kiku@agent.acme.com. Kiku was created as a designated coordinator agent that takes the requested task, like “take the attached article and prepare social media posts”, and routes it to one or more internal agents via email. Kiku knows to which agents to route to because “she” maintains a “roster” of the agents in the cluster. This roster describes the capabilities of each agent. Agents act on the part of the task sent to them and return the results to the coordinator who then returns it to the user. One of the key design aspects is the project uses email threading as state memory. When the agent’s LLM is called, the the thread is passed in as history context. Top Benefits of the Kikubot framework: - Users engage AI agents via email which makes it easy for organizational deployment - no end-user installs, just an email address to send tasks too
- Per agent user access controls
- Per agent containerization
- Highly scalable capabilities - agent teams can “nest” other teams.
- Cost containment via per agent LLM selection
- High visibility into intra-agent “conversations” via standard email tools Ease of Deployment:
We are working on making it easy to deploy and customize. We have a Configurator tool that we believe goes a long way in making setup much easier. We just added a ./demo.sh script that launches a complete demo on your machine included a local email server, webmail client and Kikubot agent. Just run ./demo.sh. Comparison with OpenClaw:
There’s a billion agent frameworks out there. Here is a side by side table comparison with OpenClaw to help understand where Kikubot fits: Dimension | Kikubot | OpenClaw
-----------------+------------------------+-----------------------
Message bus | Email (IMAP/SMTP) | Chat apps (Telegram,
| | Slack, Discord,
| | WhatsApp, etc.)
-----------------+------------------------+-----------------------
Architecture | Multi-agent by design | Single gateway with a
| - coordinator + | channel/brain/body
| specialized sub-agents | layer separation
-----------------+------------------------+-----------------------
Target scale | Organizational / | Personal productivity
| enterprise (hundreds | / individual assistant
| of agents, |
| departments) |
-----------------+------------------------+-----------------------
Agent topology | Decentralized - each | Local-first gateway
| agent independently | process on your
| deployed anywhere in | machine or VPS
| the org |
-----------------+------------------------+-----------------------
End-user install | None - users just | Requires installing
| email an address | and running the
| | gateway runtime
-----------------+------------------------+-----------------------
Primary language | Go | Node.js
Anyway, a long winded post. Check out the project. If you find it interesting please join in! There are already some very interesting applications of Kikubot across different types of organizations!

GitHub - mxaiorg/kikubot: Email-based AI agent network · GitHub
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
mxaiorg
/
kikubot
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
81 Commits 81 Commits assets assets cmd/ kikubot cmd/ kikubot configs configs docs docs internal internal scripts scripts services services .gitignore .gitignore CLAUDE.md CLAUDE.md CONFIGURATION.md CONFIGURATION.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE README-Observability.md README-Observability.md README.md README.md demo.sh demo.sh docker-compose-demo.yml docker-compose-demo.yml docker-compose-example.yml docker-compose-example.yml go.mod go.mod go.sum go.sum View all files Repository files navigation
An email-driven, multi-agent framework. Each agent is an inbox.
Kikubot turns an email account into an autonomous agent. Each running container polls one IMAP mailbox, runs every new email through an LLM agentic loop with a configurable tool set, and replies via SMTP. Agents collaborate by emailing each other, so a typical deployment looks like several agents — a coordinator and a few specialists — sharing one mail server.
Why email? It's the universal asynchronous message bus: humans already use it, every system can send to it, threads carry their own history ( References: / In-Reply-To: ), and accounts give you free per-agent identity, ACLs, and durability.
Email as the AI User Interface. Deploy AI to an organization via the most used and understood technology - email. No training required, no software to install - just an email address.
High Scalability. Clusters of agents, each agent can be its own cluster - results in theoretically massive scalability.
Observability. Agents communicate with each other via standard email. Access agent accounts to see their internal conversation history.
Cost Containment. Each agent can be configured to a different LLM. Choose the best LLM for the agent's role and toolset.
Higher Performance. Agents can specialize. Capabilities are distributed across the agent network. Each agent focuses on its area of expertise.
Greater Security. No risk of AI agents running on user machines. Kikubot Agents run in containers, providing isolation and security. Agents access tools and integrations via scoped API keys. Access to agents is easily controlled via ACLs (white or blacklist domains, or email addresses).
Resilience. Kikubot networks run over email - one of the most resilient technologies in the world.
Per-thread memory. Each email thread is a long-running conversation; the agent's history is persisted as JSON keyed by the thread's root Message-Id.
Pluggable tools. Built-in tools cover messaging, status reporting, snoozing, and mailbox search. Optional tools include Salesforce, WordPress, Buffer, Box, Helpjuice, Tavily web search, Apache Tika file-to-text, and arbitrary local/HTTP MCP servers.
Pluggable LLMs. Anthropic API (default, with prompt caching) or OpenRouter (with backup-model fallback).
Knowledge base. Per-agent and shared markdown files appended to the system prompt — editable live and hot-reloaded without a rebuild.
Multi-agent coordination. Agents talk to each other via the message_tool core tool; coordinator agents can delegate, fan out, and snooze pending work.
Recurring tasks. Agents can schedule tasks to run at specific times or intervals.
Auto-reply / bounce safety. DSNs and out-of-office replies bypass the LLM to prevent infinite delegation loops.
One Agent to Thousands of Agents
You can spawn one or more agent containers with this repository on the same machine. Each container runs a single agent. You can also deploy this repository across multiple machines and spawn agents across your organization. The only requirement is that coordinator agents can reach each other via email.
Coordinator agents can be organized into teams, and each team can have multiple agents. Coordinator agents team members can in themselves be coordinators. Much like how organizations are structured into divisions, with each division representing multiple departments which in turn represent multiple teams – so can you structure your network of agents. Each coordinator only needs to know the subset of agents it works with directly. Theoretically, a Kikubot deployment can scale to hundreds of thousands of agents.
Ask Alpha about the weather in your town. Email your query to alpha@labtest.mxhero.com .
mxHERO Labs has deployed a Kikubot instance to a single machine. The instance is configured with 2 agents: Alpha and Weatherman. Alpha is the coordinator. If you send an email to Alpha asking about the weather in some city, Alpha will ask the weatherman for the weather forecast. Upon receiving the response, Alpha will send a reply to you.
This small Kikubot demo illustrates how multiple agents can work together over email
Want to run your own Kikubot in under a minute, with no mail account and no API
key required ? The demo spins up a throwaway mail server, a webmail UI, and one
agent — all in Docker, all on your machine, nothing exposed to the internet.
git clone https://github.com/mxaiorg/kikubot && cd kikubot
./demo.sh
Then:
Open the webmail UI at http://localhost:8000
Log in as human@demo.local (any password — the demo mail server has auth disabled)
Compose a new email to kiku@demo.local and send it
Wait ~30s, refresh the inbox — Kiku replies. 🎉
Out of the box (no API key) Kiku replies with a short "I'm alive — add a key for
real answers" notice, so you get the round-trip moment at zero cost. To unlock
real agent responses, drop a key into configs/demo/secrets.env and run
./demo.sh again — paste either an ANTHROPIC_API_KEY or an OPENROUTER_API_KEY
and the script auto-selects the matching provider for you (no config editing).
Stop everything with ./demo.sh down .
Under the hood ( docker-compose-demo.yml ): GreenMail
provides SMTP/IMAP, Roundcube is the compose-window, and
the agent is the same image as production — just pointed at the demo config in
configs/demo/ .
This project is based on the research of mxHERO Labs. See our blog post for more details.
For a developer-oriented visual map of the runtime, configuration, memory, tools, and integrations, see docs/architecture.md .
┌────────────┐ ┌──────────────────┐
│ Users │──┐ │ Coordinator │ ◀──┐
└────────────┘ │ │ (Kiku inbox) │ │
▼ └────────┬─────────┘ │
┌──────────┐ │ │ email
│ IMAP / │ ▼ delegate │ threads
│ SMTP │ ┌─────────┐ ┌─────────┐ ┌─────────┐
└────┬─────┘ │ Beta │ │ Gamma │ │ Delta │
│ │ (CRM) │ │(social) │ │ (web) │
└────────┴─────────┘ └─────────┘ └─────────┘
│ │ │
Salesforce Buffer WordPress
mxMCP Tavily Helpjuice
Box, Tika
Each agent container runs an identical Go binary, parameterised by a shared configs/agents.yaml (roster + common defaults + per-agent overrides) and a shared configs/secrets.env (API keys + per-agent mailbox passwords). The container picks its identity from the AGENT_EMAIL env var injected by docker-compose.
An LLM API key — ANTHROPIC_API_KEY and/or OPENROUTER_API_KEY .
An IMAP + SMTP server with one mailbox per agent. You can use any provider; this repo includes a self-hostable docker-mailserver sidecar at services/dms/ if you want one.
(Optional) tool credentials for any integrations you enable (Salesforce, WordPress, Buffer, Helpjuice, Box, Xero, Tavily, mxMCP).
A dashboard configuration tool can be found in the scripts directory. It's a web app that lets you configure your deployment: define your agents and optionally configure the included email server.
go run ./scripts/configurator # serves on 127.0.0.1:50042
See scripts/configurator/README.md for more details. There is also a, slightly outdated, Configurator Video Tutorial
A configuration guide for LLM guided deployment. To use simply open an LLM coding agent like, Claude Code, and prompt:
Read the CONFIGURATION.md file and follow its instructions to help me
> configure kikubot.
Or if you want to configure in your own language:
Read the CONFIGURATION.md file and follow its instructions to help me
configure kikubot. Communicate with me in Japanese.
git clone https://github.com/mxaiorg/kikubot
cd kikubot
# 1. Configure the roster + common defaults.
cp configs/agents-example.yaml configs/agents.yaml
cp configs/secrets-example.env configs/secrets.env
# Edit configs/agents.yaml: keep/edit the common: defaults (mail server,
# prompts, budgets) and the agents: entries (identity, role, tools, optional
# per-agent overrides). Then edit configs/secrets.env: fill in
# ANTHROPIC_API_KEY (or OPENROUTER_API_KEY) and one
# <UPPERCASED_LOCAL_PART>_EMAIL_PASSWORD per agent, plus any tool credentials.
# 2. (Optional) Drop knowledge files into configs/knowledge/<agent>/*.md
# 3. Edit docker-compose.yml to match your roster.
cp docker-compose-example.yml docker-compose.yml
# - One service per agent. Each service sets AGENT_EMAIL in `environment:`
# and points env_file at configs/secrets.env.
# - Volume mount: ./data/<stem>:/app/data (stem = lowercased local-part).
# 4. Validate.
go run scripts/kikudoctor/ * .go
# 5. Launch.
docker compose up -d --build
Send the agent an email from a whitelisted address and watch the reply land in your inbox.
To watch the conversation between agents recorded in the logs:
docker compose logs -f
Because Kikubot is an email based system, you can use any email client to follow the internal and external conversation of your agents. See Observability
Deployment config — configs/agents.yaml
configs/agents.yaml is the single source of truth for non-secret deployment config. It has two sections:
common :
email_server : mail.agents.example.com:993
smtp_server : mail.agents.example.com:587
email_insecure_tls : false
max_history_chars : 200000
max_tokens : 6000
agent_timeout : 300
max_turns : 20
system_prompt : |
You are a helpful agent...
{{coworkers}}
coordinator_system_prompt : |
You are a helpful Coordenator Agent...
{{coworkers}}
agents :
- name : Kiku
email : kiku@agents.example.com
role : Coordinator
description : Communicates with users. Coordinates other agents.
tools : [report, snooze, tavily_mcp]
# Any common: field may be overridden here.
llm_provider : openrouter
llm_model : anthropic/claude-sonnet-4.6
max_turns : 40
whitelist : [example.com, agents.example.com]
- name : Beta
email : beta@agents.example.com
role : CRM, Email Archivist
description : Manages Salesforce and access to the comp

[truncated]
