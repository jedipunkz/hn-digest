---
source: "https://cloudywithachanceoflatency.net/blog_02-how-llm-accelerated-cloudy-dev.html"
hn_url: "https://news.ycombinator.com/item?id=49229929"
title: "How LLMs accelerated Cloudy development? not by typing code for me"
article_title: "How LLMs accelerated Cloudy development"
author: "rixed"
captured_at: "2026-08-09T10:22:40Z"
capture_tool: "hn-digest"
hn_id: 49229929
score: 1
comments: 1
posted_at: "2026-08-09T09:49:15Z"
tags:
  - hacker-news
  - translated
---

# How LLMs accelerated Cloudy development? not by typing code for me

- HN: [49229929](https://news.ycombinator.com/item?id=49229929)
- Source: [cloudywithachanceoflatency.net](https://cloudywithachanceoflatency.net/blog_02-how-llm-accelerated-cloudy-dev.html)
- Score: 1
- Comments: 1
- Posted: 2026-08-09T09:49:15Z

## Translation

タイトル: LLM はどのように Cloudy 開発を加速したか?コードを入力することではありません
記事のタイトル: LLM がどのように Cloudy 開発を加速したか
説明: ネタバレ: コードを入力することではありません。

記事本文:
遅延の可能性がある曇り ユースケース AI / MCP の機能 お問い合わせ ログイン ロードマップの開始 ドキュメント --> ブログ デモ ブログ LLM がどのようにして Cloudy 開発を加速したか
スポイラー: コードを入力することではありません。
LLM はコードを書くときに役立ちます (あらゆる状況で div を中央に配置する方法がついにわかりました)。
しかし、大規模で重要なプロジェクトを編集する場合は、作業の遅れを避けるために多くのガイダンスが必要です。これらにより、開発者はコードを深く考えるのではなく、コードを入力する手間が軽減されます。
したがって、タイピストが非常に遅いか、ひどく冗長な言語を使用する人でない限り、その方向からはあまり利益は得られません。
Cloudy は市場で最高のキーボードで入力されており、エージェントから Web フロントエンドに至るまでのほぼすべてのコードは、冗長ではない言語である OCaml で書かれています。
それでも、LLM により開発が大幅に高速化されています。どうして？
LLM によって廃止される: UI、ドキュメント、オンボーディング、プログラマー
Cloudy のような非常に技術的で低レベルの製品を開発する場合、ほとんどの時間はシステム プログラミング、ネットワーク分析、インフラストラクチャに費やされると思われるかもしれません。
しかし、それは間違いです。開発時間のほとんどは、UI の実装、テスト、デバッグに費やされます。
はい、ありふれた電子商取引 Web サイトと同様です。
Cloudy の git リポジトリに関する大まかな統計は次のとおりです。
コード行の 60% は Web UI 用です (全体で同じ言語なので、比較は成り立ちます)
コミットによってタッチされるファイルの 50% は Web UI 用です
したがって、作業の半分は UI であると推定するのが妥当だと思われます。
その理由の一部は、Web テクノロジによってもたらされる偶発的な複雑さによるものですが (それについてはやめておきます)、多くの UI の問題が非常に複雑であるためでもあります。
その 1 つは、抽象構文ツリー Cloudy u など、多くの特殊なデータ型を使用して、任意の深くネストされたデータ構造を編集することです。

ses はテスト プログラムを表します。
UI 作業では、小さな DSL (ドメイン固有言語) を編集する必要性が頻繁に発生しますが、これに対する適切なソリューションを見たことがありません。私はこれがユーザーインターフェースの聖杯だと考えるようになりました。
たとえ優れたリアクティブなレイアウトとレスポンシブなバリデータを備えていたとしても、ほとんどのユーザーはそのようなフォームを威圧的に感じるでしょう。
UI 自体に加えて、次のように記述する必要があります。
そして、ソフトウェアを学ぶために積極的に投資してくれるユーザーも必要です。
最も貴重な人々、つまりアプリを 15 分間テストドライブするだけの人は、その説明には当てはまりません。
根本的な問題は、テクノロジーにおける概念の多くは何らかの形式的な言語の形で自然に生まれるのに、人間はそれが得意ではないということです。
彼らは忍耐力のない説明と訓練を必要としています。
しかし、形式言語が得意な人は誰でしょうか? LLM。
LLM にその仕様といくつかの例を含む正式な構文を与えると、0.5 秒後にはパワー ユーザーになります。
上のスクリーンショットをこの Claude-Code セッションと比較してください。
┌─────────────┬───────────────┬─────┐
│ 探査機 │ 実際の場所 │ ヨーロッパ？ │
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┤
│ Agent-hetzner-nbg1 │ デラウェア州ニュルンベルク │ はい │
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┤
│ エージェント-hetzner-hel1 │

ドイツ (hel1 にもかかわらず) │ はい │
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┤
│ Agent-ovh-UK1 │ フランス、OVH SAS FR (英国にもかかわらず) │ はい │
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┤
│ … │ … │ … │
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┤
│ Agent-aws-us-west-1 │ 米国 │ いいえ │
━━━━━━━━━━━━━━━━━━━━━━━━━━━┘
曇りを 3 回呼び出す (Ctrl+O で展開)
tools/program-upsert (プログラムを作成するため)
tools/program-config (プログラムを開始するため)
ツール/プログラム (確認用)
● プログラムが作成され、有効化され、開始されます (開始: 1786255999)。
Cloudy-Net-monitor — 300 秒ごと。
構造 — 3 つのタスクのシーケンスをラップする 300 秒ごとの繰り返し:
1. すべてのプローブからの HTTP (from: "") — 1 つの接続再利用配列内の 8 リクエストのバッチ、HTTP/1.0、1.1、2、3 × TLS 1.2、1.3。
2. 4 つのヨーロッパのプローブ、TCP/443、maxHops 30、3 パケットから 104.21.25.188 にルーティングします。
3. 同じソースセットとパラメータで 172.67.134.122 にルーティングします。
Cloudflare が 2 つのアドレスを配布し、それらが異なる PoP に到達する可能性があるため、両方の A レコードを追跡しました。 ICMP ではなく TCP/443 なので、測定されたパスは

HTTPS トラフィックは実際にかかります。
httping を最初にシーケンスするということは、TLS/HTTP チェックが各サイクルの 2 つのtraceroute より遅れないことを意味します。
また、LLM は自然言語にも精通しているため、ユーザーのリクエストを平易な英語からアプリが使用する正式な言語に簡単に変換できます。
そして、変換するだけでなく、独自の専門知識を追加してタスクを改良し、落とし穴を回避します。
探査機の名前を信頼するのではなく、探査機の実際の位置を確認する。
Cloudflareのエニーキャストには注意してください。
テストの暗黙の目標に従って、traceroute に適切なプロトコルを選択します。
パフォーマンスを向上させるために、最初に httping の順序を変更します。
作成者にとっては、MCP サーバー内に留まる簡潔なリファレンス以上のものを記述する必要はなく、LLM がすでに持っているすべての知識を UI にエンコードする必要もありません。
ユーザーにとっては、アプリケーションの詳細や複雑な UI の操作方法を学ぶ必要はありません。
LLM はユーザー インターフェイスとなり、そこから本当のメリットが得られます。
生き残るもの: ウィジェットは千の言葉に匹敵するでしょうか?
番組編集者？これまで見てきたように、LLM ははるかに優れています。
時系列？もちろん、Cloudy にはさまざまな機能が備わっていますが、LLM には利用できる優れた組み込み機能もあります。あるいは、より強力な BI ツールを使用して、そのデータを他のソースと簡単に関連付けることもできます。それでも、データのライブ ストリームを渡すための規約がまだないため、対話は貧弱になります。
トレースルートグラフ? LLM は、クラスター化されたホストとそれらの間のリンクを含む間に合わせの世界地図をユーザーに表示できますが、それには大量のトークンがかかり、Leaflet と Python に頼る必要があり、結果は一貫性も対話性もありません。
したがって、今のところ Cloudy の UI はどこにも行きませんが、いくつかの特殊なウィジェットは

本当にユニークなものは、LLM がユーザーにそれらを指示できるように、単独で利用可能になっています。
他の、より些細なウィジェット (正直に言うと、ほとんどの UI) は使用されなくなる可能性があります。
UI は重要なことを前面に押し出します。誰かがいないと私たちは暗闇にいる
Cloudy の UI は、次のようなビジュアル コードを使用して、異常なものを可視化しようと努めます。
チャートやグラフのデータは、悪いデータに注意を引くために色分けされています。
タイムラインには、不正なデータが存在する場合にマーカーが表示されます。
プローブの失敗、BGP のフラッピングなどの不正なデータは、派手な色で目立つように表示されます。
不正なデータが含まれるパネルを開くボタンは、ユーザーの注意を引くために点滅します。
実行可能ですが、大変な作業です。
一方、LLM は、異常または不良に見えるものを自発的に報告し、それにコンテキストを追加します。
逆に、些細な理由でプローブが失敗したとしても、UI はそのような微妙な違いを認識できませんが、それを心配することはありません。
前の例に続いて、失敗したプローブを冷静さを失うことなく適切に無視していることに注目してください。
これは、LLM のファジー ロジックが自動動作を無効にする領域です。
UI がなければ、ユーザーは何ができるかをどうやって見つけられるでしょうか?
人間が運転する必要があるため、発見可能性は単に要件となります。
ユーザーは何ができるかを知りたいのではなく、目標を述べたいのです。
Cloudy は、エージェント SRE のツールボックスに含まれる、調査能力を拡張するためのツールの 1 つにすぎません。
ユーザーが知っておくべきことは次のとおりです。Cloudy は、SRE に、多くのネットワーク テストを継続的に実行し、高レベルの傾向から個々のネットワーク パケットに至るまで、それらのテストの履歴結果を参照するためのコスト効率の高い方法を提供します。
MCP の説明は、LLM にすべての詳細を伝えます。
LLM は、ツールを理解するだけでなく、

自分の出力だけでなく、どのような状況でその出力に到達するかを知ることも必要です。
クールで強力に見えても役に立たないツールもあれば、目立たないように見えても現実世界の多くのシナリオで役立つツールもあります。
Cloudy の有用性の実際の尺度は、UI の簡単なツアーだけでは簡単に評価できません。
これは、エージェント SRE がそのツールボックスに到達する頻度によって最もよく測定されます。
LLM とは異なり、UI を使用すると、アプリへの迅速かつ無料かつプライベートなアクセスが可能になります。
これは否定の余地のないことであり、まさにそれが Cloudy が依然として UI を備えている理由です。それは、平易な英語で説明するよりも右ボタンをクリックする方が速い場合、匿名化に関係なく第三者をループに参加させることができない場合、またはエージェント自体がダウンしている場合に備えています。
しかし、正直に言うと、この反対意見は過去 20 年間の IT インフラストラクチャの進化にほとんど影響を与えていません。
では、AI を使って生産性を高めるにはどうすればよいでしょうか
上記すべてに基づいて、いくつかの推奨事項を次に示します。
MCP はメインのアクセス ポイントであり、唯一のドキュメントです。
ナビゲーションは要約すると、人間のユーザーはハーネスによって、特定のデータを選択したカスタム ウィジェットに送信され、その後終了します。
ユーザー、セッション、またはアクセス制御を管理しないでください。リクエストがサービスに到達するまでに、ハーネスは上流でユーザーをすでに認証しています。
形式的な言語を敬遠しないでください。
彼らは人間に対しては悪い印象を持っていますが、LLM に対してはそうではありません。
UI が邪魔にならなければ、サービスをプログラム可能にしてサービスの範囲を拡大するのは非常に低コストです。
LLM が実験できる予行演習モード、または遊び場を提供します。
LLM により、ユーザー受け入れテストが安価になります。
早めに LLM を MCP に接続し、実験の時間を節約します。
LLM に何が必要かを尋ねます。ほとんどの場合、LLM はサービスで何ができるかについての良いアイデアを持っています。

動作に注意してください。役に立たないツールの呼び出しは、何かが曖昧であるか欠落していることを示す明らかな兆候です。
(たとえば、クロードは、一見理由もなくプログラムのステータスを尋ねていました。デフォルト値があるときに回答から省略されていたいくつかの重要なフィールドによって混乱していることが判明しました。一度特定すれば簡単に修正できます。)
予測: ハーネスは未来のオペレーティング システムです
自然言語はアプリケーションへの主要なインターフェイスになります。
アプリケーションは単なるデータプロバイダーになるでしょう。
あるアプリからのデータは別のアプリからのデータとシームレスに結合されます。これは、すべてのデータが最終的に LLM コンテキスト内に置かれ、ユーザーが実際に必要とするデータを構築するために混合されるためです。
データが特定のアプリケーションに属していたウォールド ガーデンの一時的な時代は時代遅れになります。
答えが言葉でうまく表現できない場合、LLM は、既知の形状とセマンティクス (オントロジー) を持つデータのストリームを、組み込みまたはサードパーティによって提供される汎用ビジュアライザーに渡します。
そのため、ユーザーはアプリケーションをアプリケーションとして見ることはなくなり、一時的に呼び出されるウィジェットだけが見えるようになります。アプリケーションはユーザーを認識しなくなり、プロファイル、アカウント、セッションはなくなります。これらすべてが LLM コンテキストです。
将来のオペレーティング システムは、LLM、サービス、データ、ウィジェットのいずれかを提供する MCP サーバー、およびローカルの主に一時的なファイルを接続するハーネスになる可能性があります。
変わらないのは、これらすべてを接続するネットワークと、それらを担当する専任の SRE です。
ただし、現在 SRE にはさまざまなツールを備えたエージェントが出向しています。
© 2026 - CloudyWithAChanceOfLatency |すぐに設計されました。

## Original Extract

Spoiler: not by typing code for me.

Cloudy with a chance of latency Use Cases AI / MCP Features Contact Login Start Roadmap Documentation --> Blog Demo Blog How LLMs accelerated Cloudy development
Spoiler: not by typing code for me.
LLMs do help when writing code (we finally know how to center divs in all circumstances).
But when editing large, non-trivial projects, they need a lot of guidance to avoid slop; they relieve developers of typing the code, not of thinking it through.
So unless one is a very slow typist or using terribly verbose languages, not much gain will come from that direction.
Cloudy is typed on the best keyboards on the market and almost all its code from the agents to the web front-end is written in OCaml , a language that’s everything but verbose.
Still, LLMs are making the development much faster. How come?
Obsoleted by LLM: UI, documentation, onboarding, programmers
You would think that in developing a very technical, low level product such as Cloudy , most of the time is spent doing system programming, network analysis and infrastructure.
But you would be wrong: most of the development time is spent in implementing, testing and debugging the UI.
Yes, like any run-of-the-mill e-commerce website.
Here are some rough stats on Cloudy ’s git repositories:
60% of lines of code are for the web UI (same language throughout, so the comparison holds)
50% of files touched by any commit are for the web UI
So it seems reasonable to estimate that half the work has been the UI.
Part of that is due to the accidental complexity brought by web technologies (let’s not get started on that), but also because many UI problems are genuinely hairy.
One of them is editing arbitrary, deeply nested data structures with many specialized data types, such as the abstract syntax tree Cloudy uses to represent its testing programs.
The need to edit a small DSL (Domain-Specific Language) arises frequently in UI work, and I’ve never seen a good solution for it; I came to consider this the Holy Grail of user interfaces.
Even with a good reactive layout and responsive validators, most users will find such a form intimidating:
In addition to the UI itself, one needs to write:
And then one needs users who are willing to make the investment of learning your software.
The most precious of them, those who are just giving your app a 15-minute test drive, don’t fit that description.
The fundamental issue is that a lot of concepts in tech come naturally in the form of some formal language, and humans are not good at that.
They need explanations and training that they have no patience for.
But guess who is good with formal languages? LLMs.
Give an LLM a formal syntax with its specification, a couple of examples, and half a second later it is a power user.
Compare the above screenshot with this Claude-Code session:
┌─────────────────────────┬─────────────────────────────────┬─────────┐
│ Probe │ Actual location │ Europe? │
├─────────────────────────┼─────────────────────────────────┼─────────┤
│ agent-hetzner-nbg1 │ Nuremberg, DE │ yes │
├─────────────────────────┼─────────────────────────────────┼─────────┤
│ agent-hetzner-hel1 │ Germany (despite hel1) │ yes │
├─────────────────────────┼─────────────────────────────────┼─────────┤
│ agent-ovh-UK1 │ France, OVH SAS FR (despite UK) │ yes │
├─────────────────────────┼─────────────────────────────────┼─────────┤
│ … │ … │ … │
├─────────────────────────┼─────────────────────────────────┼─────────┤
│ agent-aws-us-west-1 │ US │ no │
└─────────────────────────┴─────────────────────────────────┴─────────┘
Called cloudy 3 times (ctrl+o to expand)
tools/program-upsert (to create the program)
tools/program-config (to start the program)
tools/programs (to check)
● The program is created, enabled, and started (started: 1786255999).
cloudy-net-monitor — every 300s.
Structure — RepeatEvery 300s wrapping a Sequence of three tasks:
1. Httping from all probes (from: "") — the 8-request batch, HTTP/1.0, 1.1, 2, 3 × TLS 1.2, 1.3, in one connection-reusing array.
2. Route to 104.21.25.188 from the four European probes, TCP/443, maxHops 30, 3 packets.
3. Route to 172.67.134.122, same source set and parameters.
I traced both A records because Cloudflare hands out two addresses and they can land on different PoPs; TCP/443 rather than ICMP so the measured path is the one HTTPS traffic actually takes.
Sequencing the httping first means the TLS/HTTP checks aren't delayed behind two traceroutes each cycle.
And since an LLM is also proficient in natural languages, it can easily convert users’ requests from plain English into whatever formal language the app uses.
And not only convert them: it also adds its own expertise to refine the task and avoid the pitfalls:
checking the probes’ actual locations rather than trusting their names;
being careful with Cloudflare’s anycast;
picking the right protocol for the traceroute, in accordance with the unstated goal of the test;
sequencing the httping first for better performance.
For the author: no need to write anything more than a concise reference that stays within the MCP server, and no need to encode in the UI all that knowledge the LLM already has.
For the user: no need to learn an application’s details, nor how to navigate its complex UI.
The LLM has become the user interface, and that’s where the real gains are coming from.
What survives: is a widget worth a thousand words?
Program editor? As we have just seen, the LLM is so much better!
Timeseries? Cloudy has a large variety of those of course, but the LLM also has good built-ins it can reach out to; or it could use a more powerful BI tool, which would make it easier to correlate that data with other sources. Still, since there is not yet a convention to pass a live stream of data, the interaction would be poor.
Traceroute graphs? The LLM can present the user with a makeshift world map with clustered hosts and links between them, but that costs a lot of tokens, it has to resort to Leaflet and Python, and the result is neither consistent nor interactive.
So for now Cloudy ’s UI is going nowhere, but the few specialized widgets that are truly unique are being made available in isolation, so that the LLM can direct the user at them.
The other, more trivial widgets (most of the UI, honestly) might fall into disuse.
A UI pushes important things in front; without one we are in the dark
Cloudy ’s UI tries hard to make anything unusual visible, using visual codes such as:
Data in charts and graphs are color coded to attract the attention to bad data;
Timelines feature markers when bad data is present;
Bad data such as failing probes, BGP flapping, etc will be displayed prominently in flashy colors;
Buttons opening a panel that has bad data will blink to attract user’s attention.
It’s doable, but it’s hard work.
On the other hand, the LLM will spontaneously report anything that looks unusual or bad, and will add context to it.
Conversely, it will not freak out over a probe that fails for a trivial reason, whereas the UI is incapable of such nuance.
Following up on the previous example, notice how it rightly ignores a failing probe without losing its calm:
That’s an area where the fuzzy logic of LLMs defeats automatic behavior.
Without a UI, how can users discover what’s doable?
Discoverability is only a requirement because the human had to drive.
Users don’t want to know what’s doable, they want to state a goal.
Cloudy is just another tool in the agentic SRE’s toolbox, there to expand its investigative power.
Here is all the user needs to know: Cloudy gives the SRE a cost efficient way to continuously perform many network tests, and to consult the historical results of those tests, from high level trends down to individual network packets.
The MCP descriptions will tell the LLM all the details.
The LLM already has the background knowledge required to not only make sense of the tools and their output, but also to know in which situations to reach for them.
A tool might look cool and powerful yet be useless, whereas another might look obscure but come in handy in many real world scenarios.
The real measure of Cloudy ’s usefulness is not easily assessed from a quick tour of the UI;
it is best measured by how often the agentic SRE reaches for that toolbox.
A UI gives a quick, free and private access to the app, unlike an LLM
This is undeniable, and it is precisely why Cloudy still has a UI: for those cases when clicking the right button is faster than describing it in plain English, for those situations where no third party can be admitted in the loop regardless of anonymisation, or for when the agent itself is down.
But let’s be honest: this objection has had very little influence over the evolution of IT infrastructure for the last two decades.
So, how to be more productive with AI
Based on all the above, here are a few recommendations:
The MCP is the main access point and the only documentation.
The navigation boils down to: human user is sent by the harness to a custom widget with a specific selection of data, then leaves.
Do not manage users, sessions or access control: by the time a request reaches your service, the harness has already authenticated the user upstream.
Don’t shy away from formal languages:
They have a bad rep with humans, but not with LLMs.
Once the UI is out of the way, it is very cheap to expand the reach of a service by making it programmable.
Provide a dry-run mode, or a playground for the LLM to experiment in.
LLMs make user acceptance testing cheap:
Connect an LLM to your MCP early, and save time on experimentation.
Ask the LLM what it needs: more often than not, it has a good idea of what it could do with your service.
Watch what it does: a useless tool call is a clear sign that something is ambiguous or missing.
(for instance, Claude used to ask for program statuses for seemingly no reason. It turned out to be confused by some important fields that were omitted from the answer when they had their default value; an easy fix, once identified.)
Prediction: harnesses are the operating system of the future
Natural language will become the main interface to applications.
Applications will become mere data providers.
Data from one app will be seamlessly combined with data from another, because all of it ends up in the LLM context, which mixes it to build the data that users actually want.
The temporary era of walled gardens, when data belonged to specific applications, will go out of fashion.
When the answer is not best expressed in words, LLMs will hand off streams of data with well-known shapes and semantics (ontologies) to generic visualizers, either built-in or provided by third parties.
So users won’t see applications as applications any more, only widgets summoned for a moment; and applications won’t see users: no more profiles, accounts or sessions - all that is the LLM context.
The operating system of the future could be just that: a harness that connects LLMs, MCP servers providing either services, data or widgets, and local, mostly transient files.
What will remain unchanged are the networks connecting all this and the SREs dedicated to take care of them,
except that now the SREs are seconded by agents equipped with a variety of tools.
© 2026 - CloudyWithAChanceOfLatency | Promptly engineered.
