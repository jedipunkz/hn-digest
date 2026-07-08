---
source: "https://docs.eventsourcingdb.io/blog/2026/07/09/claude-start-my-database/"
hn_url: "https://news.ycombinator.com/item?id=48836805"
title: "Claude, Start My Database"
article_title: "Claude, Start My Database - EventSourcingDB"
author: "goloroden"
captured_at: "2026-07-08T20:57:28Z"
capture_tool: "hn-digest"
hn_id: 48836805
score: 3
comments: 1
posted_at: "2026-07-08T20:08:15Z"
tags:
  - hacker-news
  - translated
---

# Claude, Start My Database

- HN: [48836805](https://news.ycombinator.com/item?id=48836805)
- Source: [docs.eventsourcingdb.io](https://docs.eventsourcingdb.io/blog/2026/07/09/claude-start-my-database/)
- Score: 3
- Comments: 1
- Posted: 2026-07-08T20:08:15Z

## Translation

タイトル: クロード、データベースを始めてください
Article title: Claude, Start My Database - EventSourcingDB

記事本文:
クロード、データベースを始めてください - EventSourcingDB
コンテンツにスキップ
イベントソーシングDB
クロード、データベースを開始してください
検索を初期化しています
イベントソーシングDB
イベントソーシングDBについて
イベントソーシングDBについて
イベントソーシングの概要
EventSourcingDB と他のツールの比較
はじめに
はじめに
EventSourcingDB のインストール
基本
基本
イベントの種類
導入と運用
導入と運用
Docker を使用したデプロイ
事前に構築されたバイナリのデプロイ
構成とシークレットの管理
診断とトラブルシューティング
クライアントSDK
クライアントSDK
概要
拡張機能
拡張機能
概要
クロードコードプラグイン
クロードコードプラグイン
はじめに
MCPサーバー
MCPサーバー
はじめに
ベストプラクティス
ベストプラクティス
モデリングとドメイン設計
モデリングとドメイン設計
イベントのモデリング
動的整合性境界 (DCB)
実装と開発
実装と開発
バージョン管理イベント
イベント駆動型アプリケーションの構築
データ管理とパフォーマンス
データ管理とパフォーマンス
スナップショットとパフォーマンス
読み取りモデルの一貫性と遅延
運用、コンプライアンス、インフラストラクチャ
運用、コンプライアンス、インフラストラクチャ
セキュリティに関する考慮事項
イベントの署名と検証
EventSourcingDB インスタンス間の移行
実践における最終的な一貫性
カテゴリー
カテゴリー
お知らせ
実行中のインスタンスへの 1 つの文
最初から最後までのセッション
クロードは CloudEvents を心から知っています
インデックスに戻る
ゴロー・ローデン
CTO 兼創設者
メタデータ
2026 年 7 月 9 日
イースターに EventSourcingDB 用の Claude Code Plugin をリリースしたとき、SDK も MCP 構成も、Docker コンテナさえも含めず、必要のないものをすべて誇らしげにリストしました。クロードは EventSourcingDB を流暢に話せるようになっていました。しかし、すべての会話には、どこかで誰かがいるという静かな仮定が焼き込まれていました。

クロードが会話できるデータベースをすでに開始していました。
プラグインのバージョン 1.1.0 では、その前提が削除されています。 Claude は、ローカルの EventSourcingDB を独自に開始できるようになり、一貫したイベント ソース、統一されたイベント タイプのプレフィックス、CloudEvents 形式の深い理解など、長時間のセッションを著しくスムーズにするいくつかの習慣を身に付けました。この投稿では、空のターミナルで始まり、クエリ可能なイベント ストアで終わるセッションに沿って、新機能について説明します。
実行中のインスタンスへの 1 つの文 ¶
1.1.0 の目玉機能は、新しい /esdb:start-server スキルです。 「使い捨ての EventSourcingDB をスピンアップしてください」と入力します。この 1 文は、覚えておくか、少なくとも調べる必要がある Docker コマンドを置き換えます。
docker run -d --rm --name eventsourcingdb-dev -p 3000 :3000 \
thenativeweb/eventsourcingdb:最新の実行 \
--api-token = シークレット \
--データディレクトリ一時\
--http 対応 \
--https 有効 = false \
--with-ui
このセットアップについては、「ローカル開発セットアップ: 5 分でわかる EventSourcingDB」で説明しました。これは複雑ではありませんが、まさにプロジェクトの合間に忘れがちな呪文のようなものです。今、クロードがあなたのためにそれを覚えています。起動されるインスタンスは開発モードで実行されます。データは一時ディレクトリに存在し、サーバーは HTTPS ではなくプレーン HTTP を通信し、組み込みの管理 UI が有効になり、API トークンは単に Secret になります。数秒後、Claude は API ベース URL、UI のアドレス、トークン、およびすべてが一時的なものであることを通知するレポートを返します。
スキルはコマンドを実行するだけではありません。何かを開始する前に、Docker が実際に実行されていることを確認し、既存の開発インスタンスを探して、ポートが空いていることを確認します。インスタンスがすでに起動している場合、クロードはその旨を伝え、詐欺を手渡します。

2 番目の接続を開始する代わりに、接続の詳細を確認します。別のプロセスがポート 3000 を占有している場合、Claude は代わりにどのポートを使用するかを尋ねます。特定のイメージ バージョンや別のトークンが必要な場合は、そう言うだけです。デフォルトはデフォルトであり、定説ではありません。
これらのデフォルトのうち 2 つは、インスタンスの適性を形作るものであるため、一時停止する価値があります。一時データ ディレクトリは、データベースが毎回空から開始されることを意味します。これは機能であり、制限ではありません。デモは白紙の状態から始まり、テスト実行が互いに汚染することはなく、先週の火曜日に残った状態が実験を歪めているのではないかと考える必要はありません。また、API と同じポートで提供される管理 UI により、セッションを第 2 の視点で見ることができます。ターミナルでクロードと会話している間、ブラウザに表示される件名やイベントを観察したり、EventQL クエリを対話的に試したりすることができます。 UI をまだ見ていない場合は、「管理 UI の使用」で説明します。
このスキルが意図的に厳密に定めている点が 1 つあります。それは、このセットアップは開発専用であるということです。一時ストレージとプレーンな HTTP は実験には最適ですが、運用環境にはまったく適していません。実際にデプロイする準備ができたら、「EventSourcingDB の実行」で適切な方法を説明します。
セッションの開始から終了まで ¶
部分がどのように組み合わされるかを確認するために、セッション全体をプレイしてみましょう。何も実行していない空のディレクトリで Claude Code を開きます。
データベースを開始しています。 「実験用にローカル EventSourcingDB を開始します。」と入力します。クロードはプリフライト チェックを実行し、コンテナを起動し、インスタンスに ping を実行して正常であることを確認し、他のすべてのスキルが自動的に接続設定を取得できるように接続設定をエクスポートすることを提案します。
ESDB_URL をエクスポート = "http://localhost:3000"
エクスポート ESDB_API_TOKEN = "シークレット"
イベントの書き込み。あなたは続けます

: 「アーサー・C・クラーク著『2001年宇宙の旅』という本を入手し、アリスという名前の読者を登録してください。」 Claude は、イベント ソースとタイプ プレフィックスを一度尋ねます (これについては後ほど詳しく説明します)。その後、イベント ( /books/42 の書籍取得イベントと /readers/1 の読者登録イベント) を書き込みます。
歴史を読む。 「これまで/books/42 に何が起こったのですか?」クロードは、サブジェクトのイベント ストリームを読み取り、タイムスタンプ、ID、ペイロードを含む 1 つの取得イベントを提示します。
質問する。 「図書館には何冊の本がありますか？」クロードは質問を EventQL クエリに変換し、実行して回答します。自然言語から離れたことは一度もありませんし、JSON を 1 行も手書きで書いたこともありません。
データを保護します。これは新しいインスタンスであるため、結果を伴わないスキーマを試すのに最適な時期でもあります。「タイトルと著者を両方とも文字列として必要とする書籍取得イベントのスキーマを登録する」。クロードは JSON スキーマを登録し、それ以降、データベースは一致しない書籍取得イベントを拒否します。インスタンス全体が使い捨てであるため、スキーマの設計を自由に繰り返すことができます。間違った場合でも、不変のスキーマを使い続けて後悔するのではなく、再起動して再試行できます。
生で見てる。完全な効果を得るには、2 番目の端末を開き、そこでクロードに「/books の下にあるものをすべて観察してください」と尋ねます。次に、最初のターミナルから別のイベントを書き込み、それが 2 番目のターミナルに到着するのをリアルタイムで監視します。監視接続はイベントが発生するとストリーミングし、タイムアウト後に自動的に閉じるため、誤ってダングリング接続が残されることはありません。小さなことですが、平易な英語で説明したイベントが別のウィンドウにライブで表示されるのを見ると、図では表現できない方法でアーキテクチャが具体的に理解できるようになります。
掃除中。最後に「やめてください」

クロードは、コンテナと保存されているすべてのデータが失われることを通知し、インスタンスは一時ストレージで起動されたものであるため、停止します。マシンはセッション前とまったく同じようにクリーンになっています。
自分で操作したい場合は、インストールにかかる時間は 1 分もかかりません。 /plugin Marketplace add thenativeweb/claude-plugins でマーケットプレイスを追加し、 /plugin install esdb@thenativeweb でプラグインをインストールすれば完了です。 「はじめに」ガイドでは詳細が説明されており、「利用可能なスキル」ページにはプラグインでできるすべてのことがリストされています。
1.1.0 の 2 番目の変更は、自動起動データベースほど派手ではありませんが、少なくとも実際には同じくらい価値があります。これは、構造化された作業に LLM を使用したことのある人なら誰でも認識するであろう、微妙な問題に対処しています。つまり、言語モデルを独自のデバイスに任せると、言語モデルがもっともらしい値を発明してしまうという問題です。
EventSourcingDB 内のすべてのイベントには、それを生成したシステムを識別するソースと、 io.eventsourcingdb.library.book-acquired などの逆ドメイン表記のタイプが含まれます。どちらも自由形式であるため、クロードは合理的だが一貫性のない選択をしていました。 1 回の書き込みでは https://library.eventsourcingdb.io をソースとして使用し、次の書き込みでは tag:eventsourcingdb.io,2026:library を使用し、3 回目の書き込みではまったく別のものを使用する可能性があります。個々のイベントは完全に有効でした。まとめると、それらは混乱していました。ソースとイベント タイプはまさにフィルターやクエリの基準となるフィールドであるため、その混乱には代償が伴いました。
バージョン 1.1.0 では、 ESDB_SOURCE と ESDB_EVENT_TYPE_PREFIX という 2 つのセッション デフォルトが導入されています。設定されている場合、イベントを構築するすべてのスキルがそれらを使用します。設定されていない場合は、初めて値が必要になったときにクロードが 1 回質問し、セッションの残りの時間は即興で答えるのではなく、その答えに固執します。

g.選択をセッション間で永続的にするには、接続設定と一緒に変数をエクスポートします。
import ESDB_SOURCE = "https://library.eventsourcingdb.io"
import ESDB_EVENT_TYPE_PREFIX = "io.eventsourcingdb.library"
結果は最高の意味で退屈だ。セッションに書き込まれるすべてのイベントは、あたかも規律ある開発者が記述したかのように、1 つのソースと 1 つの型プレフィックスを共有します。一貫性は魅力的なものではありませんが、6 か月後にイベント ストアをクエリできるようにするものです。
これには、プラグインを超えた広範な教訓があります。意思決定を会話から設定に移すと、LLM の使用が劇的に向上します。モデルが選択する必要のない値はすべて、一貫性を欠いて選択できない値です。セッションのデフォルトは、ドリフトが最も問題となる 2 つのフィールドにその原則を適用します。環境変数は存在する中で最も退屈な構成メカニズムであるため、儀式なしで実行されます。 AI アシスタントが長時間のセッションで生成したデータを確認し、同じ識別子のスペルが 3 つ見つかったことがある場合は、これによってどの問題が解決されるのかが正確にわかるでしょう。
クロードは CloudEvents を心から知っています ¶
3 番目の改善は、クロードが生み出すものの品質を見るまでは見えません。このプラグインには、すべてのイベント関連スキルが利用する共有 CloudEvents 参照が同梱されるようになりました。これは、書き込み時にどのフィールドを指定する必要があるか (source 、 subject 、 type 、 data )、およびサーバーが管理するフィールド ( specversion 、 id 、 time など) をクロードに教えます。また、サブジェクトがスラッシュで区切られたパスやデータが JSON オブジェクトであるなど、EventSourcingDB が適用する検証ルールについても説明します。
実際には、これは拒否される書き込みが少なくなることを意味します。クロードは、被験者ができないことを試行錯誤によって発見する必要がなくなりました。

スペースが含まれていないこと、またはイベント タイプに逆ドメイン構造が必要であること。ルールはプラグインとともに伝達されるため、最初の試行で正しい形状が得られます。
また、読書時の回答が向上することも意味します。 EventSourcingDB によって返されるイベントには、CloudEvents の標準フィールドを超える拡張属性 (ハッシュ、先行ハッシュ、および署名) が含まれています。ハッシュはストアのすべてのイベントを暗号的に連鎖させます。これが改ざん検出と監査の基礎となります。署名は、サーバーが署名キーで構成されるときに設定されます。クロードは現在、これらのフィールドの意味を理解しており、それらをごまかさずに文脈に沿って説明できるようになりました。イベントがどのように構成されているかの全体像が必要な場合は、CloudEvents のドキュメントが最近拡張され、まさにこれらの詳細が記載されています。
CloudEvents リファレンスと並行して、ストリーミング NDJSON 応答の処理方法、ハートビートのフィルタリング方法、接続が既に成功を報告した後にストリームの途中で到着するエラーを特定する方法など、すべてのスキルが従う規則を統合しました。以前は、各スキルにはこれらの指示の独自のコピーが含まれていました。今、彼らは同じ場所に住んでいます。ユーザーにとって、目に見える効果は均一性です。つまり、すべてのスキルが設定を処理し、

[切り捨てられた]

## Original Extract

Claude, Start My Database - EventSourcingDB
Skip to content
EventSourcingDB
Claude, Start My Database
Initializing search
EventSourcingDB
About EventSourcingDB
About EventSourcingDB
Introduction to Event Sourcing
Comparing EventSourcingDB with Other Tools
Getting Started
Getting Started
Installing EventSourcingDB
Fundamentals
Fundamentals
Event Types
Deployment and Operations
Deployment and Operations
Deploying with Docker
Deploying the Pre-Built Binary
Managing Configuration and Secrets
Diagnostics and Troubleshooting
Client SDKs
Client SDKs
Overview
Extensions
Extensions
Overview
Claude Code Plugin
Claude Code Plugin
Introduction
MCP Server
MCP Server
Introduction
Best Practices
Best Practices
Modeling and Domain Design
Modeling and Domain Design
Modeling Events
Dynamic Consistency Boundaries (DCBs)
Implementation and Development
Implementation and Development
Versioning Events
Building Event-Driven Applications
Data Management and Performance
Data Management and Performance
Snapshots and Performance
Read-Model Consistency and Lag
Operations, Compliance and Infrastructure
Operations, Compliance and Infrastructure
Security Considerations
Event Signatures and Validation
Migrating Between EventSourcingDB Instances
Eventual Consistency in Practice
Categories
Categories
Announcements
One Sentence to a Running Instance
A Session from Start to Finish
Claude Knows CloudEvents by Heart
Back to index
Golo Roden
CTO and founder
Metadata
July 9, 2026
When we released the Claude Code Plugin for EventSourcingDB at Easter, we proudly listed everything it did not need: no SDK, no MCP configuration, not even a Docker container. Claude had learned to speak EventSourcingDB fluently. But there was a quiet assumption baked into every conversation – that somewhere, someone had already started a database for Claude to talk to.
Version 1.1.0 of the plugin removes that assumption. Claude can now start a local EventSourcingDB on its own , and it has picked up a few habits that make longer sessions noticeably smoother: consistent event sources, uniform event type prefixes, and a deeper understanding of the CloudEvents format. In this post, we walk through what is new – along a session that begins with an empty terminal and ends with a queryable event store.
One Sentence to a Running Instance ¶
The headline feature of 1.1.0 is the new /esdb:start-server Skill. You type: "Spin up a throwaway EventSourcingDB for me." That single sentence replaces a Docker command you would otherwise have to remember, or at least look up:
docker run -d --rm --name eventsourcingdb-dev -p 3000 :3000 \
thenativeweb/eventsourcingdb:latest run \
--api-token = secret \
--data-directory-temporary \
--http-enabled \
--https-enabled = false \
--with-ui
We wrote about this setup in Local Development Setup: EventSourcingDB in 5 Minutes – it is not complicated, but it is exactly the kind of incantation you forget between projects. Now Claude remembers it for you. The instance it starts runs in development mode : data lives in a temporary directory, the server speaks plain HTTP instead of HTTPS, the built-in management UI is enabled, and the API token is simply secret . A few seconds later, Claude reports back with the API base URL, the address of the UI, the token, and a reminder that everything is ephemeral.
The Skill does more than fire off a command. Before starting anything, it checks that Docker is actually running , looks for an existing development instance, and verifies that the port is free. If an instance is already up, Claude tells you so and hands you the connection details instead of starting a second one. If another process occupies port 3000, Claude asks you which port to use instead. And if you want a specific image version or a different token, you just say so – the defaults are defaults, not dogma.
It is worth pausing on two of those defaults, because they shape what the instance is good for. The temporary data directory means the database starts empty every single time – and that is a feature, not a limitation. Demos begin from a clean slate, test runs cannot contaminate each other, and you never have to wonder what leftover state from last Tuesday is skewing your experiment. And the management UI , served on the same port as the API, gives you a second pair of eyes on the session: while you talk to Claude in the terminal, you can watch subjects and events appear in the browser, or try an EventQL query interactively. If you have not seen the UI yet, Using the Management UI gives you the tour.
One thing the Skill is deliberately strict about: this setup is for development only. Temporary storage and plain HTTP are wonderful for experiments and exactly wrong for production. When you are ready to deploy for real, Running EventSourcingDB covers the proper way.
A Session from Start to Finish ¶
To see how the pieces fit together, let us play through a complete session. You open Claude Code in an empty directory, with nothing running.
Starting the database. You type: "Start a local EventSourcingDB for me to experiment with." Claude runs the preflight checks, starts the container, pings the instance to confirm it is healthy, and suggests exporting the connection settings so all other Skills pick them up automatically:
export ESDB_URL = "http://localhost:3000"
export ESDB_API_TOKEN = "secret"
Writing events. You continue: "Acquire a book called '2001 – A Space Odyssey' by Arthur C. Clarke, and register a reader named Alice." Claude asks you once for your event source and type prefix – more on that in a moment – then writes the events: a book-acquired event for /books/42 and a reader-registered event for /readers/1 .
Reading the history. "What happened to /books/42 so far?" Claude reads the event stream for the subject and presents it: one acquisition event, with its timestamp, ID, and payload.
Asking questions. "How many books does the library have?" Claude translates the question into an EventQL query, runs it, and answers. You never left natural language, and you never wrote a single line of JSON by hand.
Guarding the data. Since this is a fresh instance, it is also the perfect moment to try schemas without consequences: "Register a schema for book-acquired events that requires a title and an author, both as strings." Claude registers the JSON Schema, and from then on the database rejects any book-acquired event that does not match. Because the whole instance is disposable, you can iterate on the schema design freely – if you get it wrong, you restart and try again, instead of living with an immutable schema you regret.
Watching it live. For the full effect, open a second terminal and ask Claude there: "Observe everything under /books." Now write another event from the first terminal and watch it arrive in the second one in real time. The observe connection streams events as they happen and closes itself after a timeout, so you cannot accidentally leave a dangling connection behind. It is a small thing, but seeing an event you just described in plain English show up live in another window makes the architecture tangible in a way no diagram can.
Cleaning up. Finally: "Stop the database again." Claude reminds you that the container and all stored data will be gone – the instance was started with temporary storage, after all – and stops it. Your machine is exactly as clean as it was before the session.
If you want to follow along yourself, installation takes less than a minute: add the marketplace with /plugin marketplace add thenativeweb/claude-plugins , install the plugin with /plugin install esdb@thenativeweb , and you are set. The Getting Started guide walks you through the details, and the Available Skills page lists everything the plugin can do.
The second change in 1.1.0 is less flashy than a self-starting database, but at least as valuable in practice. It addresses a subtle problem that anyone who has used an LLM for structured work will recognize: left to its own devices, a language model invents plausible values.
Every event in EventSourcingDB carries a source that identifies the system which produced it, and a type in reverse domain notation, such as io.eventsourcingdb.library.book-acquired . Both are free-form enough that Claude used to make reasonable but inconsistent choices. One write might use https://library.eventsourcingdb.io as the source, the next one tag:eventsourcingdb.io,2026:library , and a third one something else entirely. Each individual event was perfectly valid. Taken together, they were a mess – and since sources and event types are precisely the fields you filter and query by, that mess had a cost.
Version 1.1.0 introduces two session defaults: ESDB_SOURCE and ESDB_EVENT_TYPE_PREFIX . If they are set, every Skill that constructs events uses them. If they are not set, Claude asks you once – the first time a value is needed – and then sticks to your answer for the rest of the session instead of improvising. To make the choice permanent across sessions, you export the variables alongside the connection settings:
export ESDB_SOURCE = "https://library.eventsourcingdb.io"
export ESDB_EVENT_TYPE_PREFIX = "io.eventsourcingdb.library"
The result is boring in the best possible way. Every event written in a session shares one source and one type prefix, exactly as if a disciplined developer had written them. Consistency is not glamorous, but it is what makes an event store queryable six months later.
There is a broader lesson in this that goes beyond the plugin. Working with an LLM gets dramatically better when you move decisions out of the conversation and into configuration. Every value the model does not have to choose is a value it cannot choose inconsistently. The session defaults apply that principle to the two fields where drift hurts the most – and they do it without ceremony, because an environment variable is the most boring configuration mechanism there is. If you have ever reviewed data an AI assistant produced over a long session and found three spellings of the same identifier, you know exactly which problem this solves.
Claude Knows CloudEvents by Heart ¶
The third improvement is invisible until you look at the quality of what Claude produces. The plugin now ships a shared CloudEvents reference that all event-related Skills draw on. It teaches Claude which fields are yours to provide when writing – source , subject , type , and data – and which ones the server manages, such as specversion , id , and time . It also covers the validation rules EventSourcingDB enforces, like subjects being slash-separated paths and data being a JSON object.
In practice, this means fewer rejected writes . Claude no longer has to discover by trial and error that a subject may not contain spaces or that an event type needs its reverse domain structure. It gets the shape right on the first attempt, because the rules travel with the plugin.
It also means better answers when reading. Events returned by EventSourcingDB carry extension attributes beyond the CloudEvents standard fields: a hash , a predecessorhash , and a signature . The hashes chain all events of the store together cryptographically, which is the foundation for tamper detection and auditing; the signature is set when the server is configured with a signing key. Claude now knows what these fields mean and can explain them in context instead of glossing over them. If you want the full picture of how events are structured, the CloudEvents documentation has recently been extended with exactly these details.
Alongside the CloudEvents reference, we consolidated the conventions that every Skill follows – how to handle streaming NDJSON responses, how to filter out heartbeats, how to spot an error that arrives mid-stream after the connection already reported success. Previously, each Skill carried its own copy of these instructions; now they live in one shared place. For you as a user, the visible effect is uniformity : every Skill treats configuration,

[truncated]
