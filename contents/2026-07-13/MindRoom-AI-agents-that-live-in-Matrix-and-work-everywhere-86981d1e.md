---
source: "https://www.nijho.lt/post/mindroom/"
hn_url: "https://news.ycombinator.com/item?id=48899360"
title: "MindRoom: AI agents that live in Matrix and work everywhere"
article_title: "MindRoom: AI agents that live in Matrix and work everywhere 🧠 | Bas Nijholt"
author: "AdamGibbins"
captured_at: "2026-07-13T22:45:50Z"
capture_tool: "hn-digest"
hn_id: 48899360
score: 2
comments: 0
posted_at: "2026-07-13T21:48:14Z"
tags:
  - hacker-news
  - translated
---

# MindRoom: AI agents that live in Matrix and work everywhere

- HN: [48899360](https://news.ycombinator.com/item?id=48899360)
- Source: [www.nijho.lt](https://www.nijho.lt/post/mindroom/)
- Score: 2
- Comments: 0
- Posted: 2026-07-13T21:48:14Z

## Translation

タイトル: MindRoom: マトリックスに住んでどこでも働く AI エージェント
記事のタイトル: MindRoom: マトリックスに住んでどこでも働く AI エージェント 🧠 |バス・ナイホルト
説明: MindRoom は、Matrix 内に存在する AI エージェントを作成するために私が構築したオープンソース システムです。 Matrix は Slack、Telegram、Discord などとの橋渡しをするため、永続メモリ、マルチエージェントのコラボレーション、100 を超える組み込みツールの統合により、エージェントはどこにいてもあなたをフォローします。

記事本文:
MindRoom: マトリックスに住んでどこでも働く AI エージェント 🧠 |バス・ナイホルト
検索
MindRoom: マトリックスに住んでどこでも働く AI エージェント 🧠
MindRoom: マトリックスに住んでどこでも働く AI エージェント 🧠
Matrix、Python、永続メモリ、100 以上の組み込みツール統合を備えたクロスプラットフォーム AI エージェント
1. 問題: AI がアプリに閉じ込められる
4. マトリックスに基づく構築: 良い点と難しい点
6. 100以上の組み込みツール統合
7. チーム: 協力するエージェント
8. ホットリロード: ダウンタイムなしで設定とプラグインを変更します
9. 音声、スケジュール、その他の機能
10. 強迫観念、SaaS の夢、そして燃え尽き症候群
11. OpenClaw とそれが重要な理由
1. 問題: AI がアプリに閉じ込められている #
プログラミングは私の人生最大の情熱です。
私は 10 年以上にわたってオープンソースに積極的に関わっており、Python、JavaScript、Rust にわたる 40 以上のインストール可能なパッケージを保守しています。
これらの中には、ユーザーが 100,000 人を超えるものもあります。他の人はまさに 1 つ、つまり私を持っています。
このことを言及したのは、今後のことにとって重要だからです。AI コーディング ツールが登場したとき、私はゼロから始めたわけではありません。
エージェント コーディングの投稿で書いたように、AI は私の仕事のやり方を完全に変えました。AFK のときは、携帯電話に音声入力し、それを独自の文字起こしソフトウェアにルーティングすることで、全機能またはプロジェクト全体を定期的に構築しています。
これらすべてを通じて、私は AI エージェントが未来であると深く確信しました。
プログラマにとって単なる便利なツールではなく、誰もがコンピュータと対話する方法が根本的に変わります。
しかし、それが本当であれば、今日のエージェントの働き方は破綻していることになります。
私たちは人工知能の間に人工の障壁を築きました。
ChatGPT は 1 つのタブであなたを認識し、Claude は別のタブであなたを認識し、Slack ボットは 3 番目のタブであなたを認識します。
メール AI はカレンダー AI にその会議の招待状を伝えることはできません

ステーション。
コーディング アシスタントは、ドキュメント AI に保存されているプロジェクトの仕様についてまったく知りません。
プラットフォームを切り替えると、AI がゼロから始まります。
ほとんどのエージェント フレームワークでは、プログラミングする必要があります。
それは私のような開発者にとってはうまくいきますが、エージェントが本当に未来であるならば、エージェントはすべての人のために働く必要があります。
プログラマーではない人はコードを書きたいのではなく、ただチャットしたいだけです。
そしてプライバシーもあります。
Gemini とメールを共有しても問題ありません。Google はすでに私の Gmail を所有しています。
しかし、自分の財務データ、健康記録、または個人的なメモをクラウド プロバイダーに送信したいでしょうか?
あまり。
私は機密性の高いタスクにはローカル モデルを使用し、汎用の詳細な研究には最適で安価な中国製モデルを選択したいと考えています。
そこで私は、チャット プロトコルに基づいて MindRoom を構築しました。単にチャットが必要な場合は、チャットを提供します。
エージェントはあなたと同じ部屋に住んでおり、スレッドで共同作業し、あらゆるプラットフォームにわたってあなたをフォローします。
最初のコミットは 2025 年 7 月 29 日に着地しました (OpenClaw の最初のコミットの 4 か月前 😅)。つまり、私がこの記事を公開するまでに、MindRoom は 11 か月近くも私と一緒にいたことになります。最初は熱心なサイド プロジェクトとして、そして 3 月以降は、より継続的に重点を置くようになりました。
最初は夢中になりすぎて、最終的には燃え尽きて完全に離れる必要がありました。
最近、OpenClaw が同様の問題を解決して大きな注目を集めているのを見て、私が構築してきたものが単なるニッチな強迫観念ではなかったことを思い出しました。
そこで私はここで、MindRoom のほこりを払い、MindRoom が何であるか、どのように機能するか、そしてなぜこのアプローチが依然として重要であると考えるのかについて書いています。
MindRoom は、Matrix プロトコル内に存在する AI エージェントを作成するオープンソース システムです。
Matrix について詳しくない方のために説明しますと、Matrix はフェデレーション型のエンドツーエンドの暗号化通信標準です。
フランス政府が550万人の国民に使用しているのと同じプロトコル

15 万以上の組織向けのドイツのヘルスケアと、何百万人もの人々が毎日使用する Element アプリによって、従業員に提供されています。
重要な洞察: Matrix は、Slack、Telegram、Discord、WhatsApp、IRC、電子メール、さらには SMS など、多くの主要なプラットフォームへの橋渡しをしています。
したがって、AI エージェントが Matrix 内に存在する場合、どのプラットフォームでもアクセスできます。
1 つのエージェント、すべてのプラットフォーム、継続的なメモリ。
実際的な注意点が 1 つあります。マトリックス ブリッジには完成度が異なります。 Telegram ブリッジのように、非常にうまく機能するものもありますが、扱いにくいものもあります。走行距離は、必要なプラットフォームによって異なる場合があります。
config.yaml での一般的なセットアップは次のようになります。
エージェント:
コード:
表示名 : コードエージェント
役割 : コードの生成、ファイルの管理、シェルコマンドの実行
モデル：opus-4.8
ツール: [ファイル、シェル、github]
指示:
- ファイルを変更する前に、必ずファイルを読み取ります。
部屋 : [ ロビー、開発者 ]
研究 :
表示名 : リサーチエージェント
役割 : Web の検索、論文の要約、情報の検索
モデル：gpt-5.5
ツール: [tavily、arxiv、wikipedia]
部屋：[ロビー、リサーチ]
チーム：
スーパーチーム :
表示名 : スーパーチーム
エージェント : [ コード、リサーチ ]
モード: コラボレーション
エージェントを定義し、ツールと部屋を提供すると、エージェントはアバター、入力インジケーター、オンライン ステータスとともに実際のユーザーとして Matrix に表示されます。
中心となるのは、私が MultiAgentOrchestrator と呼ぶものです。これは、構成されたすべてのエンティティ (ルーター、エージェント、チーム) を起動し、 matrix ユーザー アカウントを matrix-nio 経由で各エンティティにプロビジョニングし、同期ループを維持する bot.py 内のクラスです。
エージェント自体は、AI モデル プロバイダー全体に統一されたインターフェイスを提供する Agno フレームワークを利用しています。
誰かがマトリックス ルームでメッセージを送信すると:
明示的なメンションは指定されたエージェントに直接送信されます (@mindroom_code help me debug)
スレッドの継続により、同じエージェントが応答し続けます (奇妙なコンテキストの切り替えはありません)

会話の途中でチェス)
新しい会話がルーターに到達します。AI がメッセージを分析し、能力に基づいて最適なエージェントを選択します。
すべての会話はスレッドで行われるため、部屋が整理されます。
エージェントは応答をリアルタイムでストリーミングし、新しいメッセージをスパム送信するのではなく、考えたとおりに 1 つのメッセージを編集します。
ツール呼び出しがライブで行われていることがわかります。
🔧 ツール呼び出し: search_web(query="マトリックス プロトコル ブリッジ")
【結果待ち…】
✅ 検索ウェブ結果:
【結果はこちら】
以下の記録はメインのパーソナル ルームで開始され、「MindRoom の方法と何ができるかを説明してください」というプロンプトが送信されます。 、その後、MindRoom が新しいスレッドを開き、返信を完了までストリーミングします。流れをスキミングできるようにするために、2.5 倍にスピードアップされます。
4. マトリックスに基づく構築: 良い点と難しい点 #
Matrix 上に構築することの最も優れた点の 1 つは、無料で得られるものです。
エンドツーエンドの暗号化、多くのプラットフォームへのブリッジを備えた高度に相互運用可能なチャット プロトコル、クライアントの選択、組織間のフェデレーションなど、すべてがこのプロトコルに付属しています。
自分で何かを構築する必要はありません。
しかし、Matrix の仕様は非常に厳しいため、課題も生じます。
このプロトコルはストリーミングをサポートしていません。
AI エージェントが 30 秒間考えてから壁のテキストを吐き出すと、チャット エクスペリエンスがひどいものになるため、新しいトークンが到着したときに同じメッセージを迅速に編集して、ストリーミング インをハッキングしました。
エージェントがまだ考えている間、⋯ マーカーが表示されます。これは小さなタッチですが、エクスペリエンスに応答性があり、生き生きとしているように感じられます。
メッセージの内容にもサイズ制限があります。
人間のチャットには問題ありませんが、AI の応答は、特にツールの呼び出しとその結果が含まれる場合には長くなる可能性があります。
私は、Matrix の添付機能を使用してこの問題を回避しました。応答が制限を超えると、コンテンツは継続されます。

メッセージがストリーミングされ続けると更新される添付ファイル内。
これには、添付ファイルがダウンロード可能なファイルとしてではなくインラインで表示されるように、Cinny チャット クライアントをフォークする必要があり、全体がシームレスになりました。
MindRoom は現在、2 つのメモリ実装をサポートしています。
1 つ目は、 Mem0 の AsyncMemory 上に構築されたより伝統的なセマンティック メモリ システムで、構成可能な埋め込みプロバイダー ( OpenAI 、 Ollama 、または HuggingFace ) と ChromaDB 経由のベクトル ストレージを備えています。
これにより、エージェントは設定、プロジェクトのコンテキスト、意思決定、繰り返し発生する事実などを検索可能な記憶を得ることができます。
2 つ目は、OpenClaw スタイルのエージェント ファイルのシンプルさにインスピレーションを得たファイルベースのメモリ システムです。
ID をデータベースに埋め込む代わりに、永続的なコンテキストをエージェントと一緒に移動するプレーンな Markdown ファイル内に保存できます。
移植性が重要です。同じエージェント ID を、最初から開始しなくても、OpenClaw、Hermes、MindRoom に移動できます。
MindRoom には、会話から永続的な事実を定期的に抽出してメモリに書き戻すメモリ フラッシュ ステップもあります。
したがって、チャットの記録はチャットのままですが、安定した設定、決定、アイデンティティの詳細は将来のターンで再利用可能なコンテキストになります。
6. 100 を超える組み込みツールの統合 #
エージェントは 100 を超える組み込み統合を使用できます。
ツールは遅延読み込みされ、資格情報が管理されるため、エージェントは必要なものだけを読み込みます。
7. チーム: 協力するエージェント #
単一のエージェントは便利ですが、チームが必要な場合もあります。
MindRoom は 2 つのコラボレーション モードをサポートしています。
調整モード: リード エージェントが他のエージェントを調整します。
あなたが質問すると、リーダーがサブタスクを委任し、結果を収集して、統一された回答をまとめます。
コラボレーション モード: すべてのエージェントが同じタスクを並行して処理し、それぞれが独立した分析を提供します。
その後、システムは彼らの応答をマージします。

コンセンサスサマリ。
実際には、1 人のエージェントが学術論文を検索し、もう 1 人が業界ニュースをチェックし、3 人目が主張を検証する調査チームがあり、これらすべてが 1 つのメッセージによって引き起こされる場合があります。
以下のライブ ルーム ビューは、通常のエージェントが使用するのと同じマトリックス スレッド サーフェスでの実際のチームの応答を示しています。スレッド、ルーティングされたエージェント、およびルーム履歴はすべて、個別のダッシュボードではなく、第一級のチャット オブジェクトです。
8. ホットリロード: ダウンタイムなしで構成とプラグインを変更します #
config.yaml は実行時に監視されます。
エージェントの追加、モデルの変更、手順の更新などを編集すると、MindRoom は古い構成と新しい構成を比較し、影響を受けるエージェントのみを正常に再起動し、ルームに再参加させます。
ダウンタイムがなく、ランタイムを完全に再起動する必要もありません。
同じ開発ループがプラグインにも適用されます。
MindRoom には豊富な Python フック システムがあり、システムの実行中にプラグインをライブ開発できます。
プラグイン ファイルを変更して保存すると、MindRoom は Python ランタイムを再起動せずにプラグインを自動的に再読み込みします。
これは些細なことのように聞こえますが、エージェントの動作やプラグイン フックを反復処理するときに、コードと構成を微調整して数秒で結果を確認できると、開発ループが大幅に改善されます。
9. 音声、スケジュール、その他の機能 #
いくつかの機能は、純粋に自分自身のために必要だったために構築しました (私のローカル AI の旅を読んだ人なら誰でも証明できるように、私のプロジェクトで繰り返されるテーマです)。
ターミナルからエージェントと対話する方法が必要だったので、Matrix CLI クライアントである Matty を真夜中にベッドから構築しました。
音声メッセージ : Matrix の音声メッセージは Whisper 経由で自動転写され、通常のテキスト入力として扱われます。
エージェントと話すことができます。
スケジュールされたタスク: cron ジョブによる自然言語スケジューリング (!schedule "毎朝午前 9 時にメールをチェックしてください" )。
エージェントは次のことができます

バックグラウンドでタスクを実行し、必要に応じてエスカレーションします。
DM サポート : エージェントは 1 対 1 の会話で、メンションを必要とせずに自然に応答します。
組織間のフェデレーション : Matrix はフェデレーションされているため、2 つの企業の AI エージェントが共有ルームで共同作業できます。これは、独自のプラットフォームでは難しいことです。
10. 強迫観念、SaaS の夢、そして燃え尽き症候群 #
MindRoom で実際に何が起こったのかについては正直に話す必要があります。なぜなら、「人生に邪魔が入った」というのはこの話の無害化されたバージョンだからです。
実際に何が起こったのかというと、私は完全に、完全に夢中になってしまったのです。
私は本業に取り組んでいる間も寝ている間も常に MindRoom に取り組んでいました。
何百時間も費やしました。
私は仕事を辞めて、AI関連のスタートアップを立ち上げることを真剣に考えました。
Armin Ronacher の投稿「Agent Psychosis: Are We Going Insane?」を読んだことがあれば、 、私が何を言っているか正確にわかるでしょう。
AI エージェントを使用した構築によるドーパミンの打撃は信じられないほど現実的です。
アーミンは次のように書いています。「生産的だと感じ、すべてが素晴らしいと感じます。そして、何のチェックもせずに、同じことに興味を持っている人々だけと付き合っていると、これはすべて完全に理にかなっているという信念にますます深く陥ります。」
それが私でした。
私は構築して構築し、次から次へと機能を出荷していましたが、それは信じられないほどの気分でした。
コードベースは成長しました

[切り捨てられた]

## Original Extract

MindRoom is an open-source system I built that creates AI agents living in Matrix. Because Matrix bridges to Slack, Telegram, Discord, and more, your agents follow you everywhere—with persistent memory, multi-agent collaboration, and 100+ built-in tool integrations.

MindRoom: AI agents that live in Matrix and work everywhere 🧠 | Bas Nijholt
Search
MindRoom: AI agents that live in Matrix and work everywhere 🧠
MindRoom: AI agents that live in Matrix and work everywhere 🧠
Cross-platform AI agents with Matrix, Python, persistent memory, and 100+ built-in tool integrations
1. The problem: your AI is trapped in apps
4. Building on Matrix: the good and the tricky
6. 100+ built-in tool integrations
7. Teams: agents that collaborate
8. Hot-reload: change config and plugins without downtime
9. Voice, scheduling, and other features
10. The obsession, the SaaS dream, and the burnout
11. OpenClaw and why it matters
1. The problem: your AI is trapped in apps #
Programming is my biggest passion in life.
I’ve been actively involved in open source for over ten years now, maintaining 40+ installable packages across Python, JavaScript, and Rust.
Some of these have over 100,000 users; others have exactly one—me.
I mention this because it matters for what comes next: when AI coding tools arrived, I didn’t start from zero.
As I wrote about in my agentic coding post , AI has completely transformed how I work—when I’m AFK, I regularly build full features or entire projects by dictating to my phone and routing it through my own transcription software .
Through all of this, I became deeply convinced that AI agents are the future.
Not just a useful tool for programmers—a fundamental shift in how everyone will interact with computers.
But if that’s true, then the way agents work today is broken.
We’ve built artificial barriers between artificial intelligences.
ChatGPT knows you in one tab, Claude knows you in another, your Slack bot knows you in a third.
Your email AI can’t tell your calendar AI about that meeting invitation.
Your coding assistant has no idea about the project specs sitting in your document AI.
Switch platforms, and your AI starts from scratch.
Most agent frameworks require you to program them.
That works for developers like me, but if agents are truly the future, they need to work for everyone.
Non-programmers don’t want to write code—they just want a chat.
And then there’s privacy.
I’m fine sharing my email with Gemini—Google already owns my Gmail.
But do I want to send my financial data, my health records, or my personal notes to a cloud provider?
Not really.
I’d rather use a local model for sensitive tasks and pick the best, cheapest Chinese model for general-purpose deep research.
So I built MindRoom on a chat protocol—if people just want a chat, give them a chat.
Agents live in the same rooms as you, collaborate in threads, and follow you across every platform.
The first commit landed on July 29, 2025 (four months before OpenClaw’s first commit 😅), so by the time I’m publishing this MindRoom has been with me for nearly eleven months: first as an obsessive side project, then, since March, with much more sustained focus.
At first I got so obsessed with it that I eventually burned out and had to step away completely.
Recently, seeing OpenClaw gain massive traction solving a similar problem reminded me that what I’d been building wasn’t just a niche obsession.
So here I am, dusting off MindRoom and writing about what it is, how it works, and why I think the approach still matters.
MindRoom is an open-source system that creates AI agents living inside the Matrix protocol .
If you’re not familiar with Matrix—it’s a federated, end-to-end encrypted communication standard.
The same protocol used by the French government for 5.5 million civil servants, by German healthcare for 150K+ organizations, and by the Element app that millions of people use daily.
The key insight: Matrix has bridges to many major platforms—Slack, Telegram, Discord, WhatsApp, IRC, email, even SMS.
So if your AI agent lives in Matrix, it can reach you on any platform.
One agent, every platform, continuous memory.
One practical caveat: Matrix bridges vary in maturity. Some, like the Telegram bridge , work very well, while others can be finicky. Your mileage may vary depending on which platforms you need.
Here’s what a typical setup looks like in config.yaml :
agents :
code :
display_name : CodeAgent
role : Generate code, manage files, execute shell commands
model : opus-4.8
tools : [ file, shell, github]
instructions :
- Always read files before modifying them.
rooms : [ lobby, dev]
research :
display_name : ResearchAgent
role : Search the web, summarize papers, find information
model : gpt-5.5
tools : [ tavily, arxiv, wikipedia]
rooms : [ lobby, research]
teams :
super_team :
display_name : Super Team
agents : [ code, research]
mode : collaborate
Define your agents, give them tools and rooms, and they show up in Matrix as real users—with avatars, typing indicators, and online status.
At the core sits what I call the MultiAgentOrchestrator —a class in bot.py that boots every configured entity (router, agents, teams), provisions Matrix user accounts for each one via matrix-nio , and keeps sync loops alive.
The agents themselves are powered by the Agno framework, which provides a unified interface across AI model providers.
When someone sends a message in a Matrix room:
Explicit mentions go directly to the named agent ( @mindroom_code help me debug this )
Thread continuations keep the same agent responding (no weird context switches mid-conversation)
New conversations hit the router —an AI that analyzes the message and picks the best agent based on capabilities
All conversations happen in threads, which keeps rooms organized.
Agents stream their responses in real-time, editing a single message as they think rather than spamming new ones.
You see tool calls happening live:
🔧 Tool Call: search_web(query="matrix protocol bridges")
[waiting for result...]
✅ search_web result:
[results here]
The recording below starts in the main Personal room, sends the prompt Explain how MindRoom and what you can do. , and then follows MindRoom as it opens the new thread and streams the reply to completion. It is sped up 2.5x to keep the flow skimmable.
4. Building on Matrix: the good and the tricky #
One of the best things about building on Matrix is what you get for free.
End-to-end encryption, a deeply interoperable chat protocol with bridges to many platforms, a choice of clients, federation between organizations—all of that comes with the protocol.
You don’t have to build any of it yourself.
But because Matrix has such a tight specification, it also brings challenges.
The protocol doesn’t support streaming.
AI agents that think for 30 seconds before dumping a wall of text make for a terrible chat experience, so I hacked streaming in by rapidly editing the same message as new tokens arrive.
An ⋯ marker shows while the agent is still thinking—a small touch, but it makes the experience feel responsive and alive.
There’s also a size limit on message content.
That’s fine for human chat, but AI responses can get long—especially when tool calls and their results are included.
I worked around this by using Matrix’s attachment feature: when a response exceeds the limit, the content continues in an attachment that gets updated as the message keeps streaming in.
This required forking the Cinny chat client so that attachments display inline rather than as downloadable files, making the whole thing seamless.
MindRoom currently supports two memory implementations.
The first is the more traditional semantic memory system built on Mem0 ’s AsyncMemory , with configurable embedding providers ( OpenAI , Ollama , or HuggingFace) and vector storage via ChromaDB .
This gives agents searchable memory for preferences, project context, decisions, and recurring facts.
The second is a file-based memory system inspired by the simplicity of OpenClaw-style agent files.
Instead of burying identity in a database, durable context can live in plain Markdown files that travel with the agent.
That portability matters: I can move the same agent identity from OpenClaw to Hermes to MindRoom without starting from scratch.
MindRoom also has a memory flush step that periodically extracts durable facts from conversations and writes them back into memory.
So the chat transcript remains chat, while stable preferences, decisions, and identity details become reusable context for future turns.
6. 100+ built-in tool integrations #
Agents can use over 100 built-in integrations:
Tools are lazy-loaded and credential-managed, so an agent only loads what it needs.
7. Teams: agents that collaborate #
Single agents are useful, but sometimes you need a team.
MindRoom supports two collaboration modes:
Coordinate mode : A lead agent orchestrates others.
You ask a question, the lead delegates subtasks, collects results, and synthesizes a unified response.
Collaborate mode : All agents work on the same task in parallel, each providing their independent analysis.
The system then merges their responses with a consensus summary.
In practice, you might have a research team where one agent searches academic papers, another checks industry news, and a third validates claims—all triggered by a single message.
The live room view below shows a real team response in the same Matrix thread surface normal agents use: threads, routed agents, and room history are all first-class chat objects rather than a separate dashboard.
8. Hot-reload: change config and plugins without downtime #
config.yaml is watched at runtime.
When you edit it—add an agent, change a model, update instructions—MindRoom diffs the old and new config, gracefully restarts only the affected agents, and has them rejoin their rooms.
No downtime, no full runtime restart required.
The same development loop applies to plugins.
MindRoom has a rich Python hook system, and plugins can be live-developed while the system is running.
Change a plugin file, save it, and MindRoom automatically reloads the plugin without restarting the Python runtime.
This sounds minor, but when you’re iterating on agent behavior or plugin hooks, being able to tweak code and config and see results in seconds significantly improves the development loop.
9. Voice, scheduling, and other features #
Some features I built purely because I wanted them for myself (a recurring theme in my projects, as anyone who’s read my local AI journey can attest).
I even built Matty , a Matrix CLI client, from bed at midnight because I needed a way to interact with my agents from the terminal:
Voice messages : Audio messages in Matrix are auto-transcribed via Whisper and treated as regular text input.
You can talk to your agents.
Scheduled tasks : Natural language scheduling ( !schedule "check my email every morning at 9 AM" ) backed by cron jobs.
Agents can run tasks in the background and escalate to you when needed.
DM support : Agents respond naturally in 1:1 conversations without needing mentions.
Cross-organization federation : Because Matrix is federated, two companies’ AI agents can collaborate in a shared room—something that’s hard to do on proprietary platforms.
10. The obsession, the SaaS dream, and the burnout #
I need to be honest about what actually happened with MindRoom, because “life got in the way” is a sanitized version of the story.
What really happened is that I got completely, utterly obsessed.
Every single second I wasn’t working at my day job or sleeping, I was working on MindRoom.
Many hundreds of hours went into it.
I seriously considered quitting my job and starting an AI startup around it.
If you’ve read Armin Ronacher’s post Agent Psychosis: Are We Going Insane? , you’ll know exactly what I’m talking about.
The dopamine hit from building with AI agents is incredibly real.
As Armin writes: “You feel productive, you feel like everything is amazing, and if you hang out just with people that are into that stuff too, without any checks, you go deeper and deeper into the belief that this all makes perfect sense.”
That was me.
I was building and building, shipping feature after feature, and it felt incredible.
The codebase grew t

[truncated]
