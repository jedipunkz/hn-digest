---
source: "https://agnost.ai"
hn_url: "https://news.ycombinator.com/item?id=48908950"
title: "Launch HN: Agnost AI (YC S26) – Extract user feedback from agent conversations"
article_title: "Agnost AI: Catch Agent Failures Your Evals Miss"
author: "laalshaitaan"
captured_at: "2026-07-14T17:05:04Z"
capture_tool: "hn-digest"
hn_id: 48908950
score: 12
comments: 0
posted_at: "2026-07-14T16:06:18Z"
tags:
  - hacker-news
  - translated
---

# Launch HN: Agnost AI (YC S26) – Extract user feedback from agent conversations

- HN: [48908950](https://news.ycombinator.com/item?id=48908950)
- Source: [agnost.ai](https://agnost.ai)
- Score: 12
- Comments: 0
- Posted: 2026-07-14T16:06:18Z

## Translation

タイトル: HN: Agnost AI (YC S26) を開始 – エージェントの会話からユーザー フィードバックを抽出
記事のタイトル: Agnost AI: エージェントの失敗を見つけて評価を逃す
説明: Agnost AI は、本番環境での会話を継続的に分析し、評価で見逃したエージェントの障害を検出し、最も影響の大きいパターンをレビュー済みの修正に変換します。 Google と 25 以上の AI チームから信頼されています。
HN テキスト: HN さん、私たちは Shubham と Parth です。幼馴染みで Agnost AI ( https://agnost.ai ) を構築しており、チャットおよび音声エージェントを構築するチーム向けの製品分析を行っています。私たちは本番環境での会話を読んで、ユーザーが激怒してプロンプトする（エージェントを罵る）、同じリクエストを繰り返し言い換える、エージェントを修正する、不足している機能を尋ねる、技術的に成功した回答後に退席するなど、行動上の失敗を見つけます。サインアップ不要のインタラクティブなデモがここにあります: https://app.agnost.ai?demo=true デモビデオはこちらです: https://www.tella.tv/video/agnost-ai-launch-hn-demo-9haa 中心的な問題は、チャットおよび音声製品には Web アプリと同じメトリクスがないことです。製品のインターフェースが言語である場合、クリックやファネルはあまり役に立たなくなります。また、ユーザーが明確なフィードバックを与えることはほとんどなく、フィードバックを与える場合も通常は表面化されたものになります。私自身は、Claude や Codex に /フィードバック をほとんど入力しません。ほとんどのユーザーはただ悪態をついたり、再度質問したり、エージェントを訂正したり、あるいは去ったりするだけです。そのため、製品エンジニアは遅延、エラー、トレースから技術的な可視性を得ることができますが、ユーザーが望むものを得ることができたかどうかを推測する必要があります。私たちは昨年エージェントを中心に構築した後にここにたどり着きましたが、数人の創設者から、構築中の AI アシスタント向けの会話用の PostHog のようなものを求められました。私たちは可観測性空間や evals 空間に入ろうとしているわけではありません。可観測性により、技術的に何が起こったかがわかります。 Evals は、すでに知っているケースを検証します。私たちはもっと

発見側では、ユーザーが何を望んでいるのか、どこで不満を感じているのか、何を繰り返し要求しているのか、どのような新しい eval が存在する必要があるのか​​などです。チームは、SDK または OTel を介して、オプションでアカウント、プラン、ソース、組織などのメタデータを含むエージェント会話メッセージを私たちに送信します。私たちは会話を製品固有のインテントに分類します。機能リクエストとバグはデフォルトのカテゴリです。他のほとんどのクラスターは顧客のデータから動的に作成され、時間の経過とともに進化します。わかりやすい英語で独自のクラスターを作成できます。クラスターが広すぎる場合は、クラスターを分割します。新しいパターンが出現した場合は、それを提案します。ある AI ビデオ編集会社は Agnost AI を使用して、チャット内に隠された機能リクエストを見つけました。最も大きな問題は、約 70 人のユーザーが自動字幕を望んでいたことでしたが、ユーザーは 1 回のセッションで「このフレームにこのテキストを追加」を 12 回要求し、「字幕を付けてもらえますか」、「音声のトランスクリプトを提供してください」と言語間のバリエーションを要求しました。チームは後にこの機能を構築しました。すべてを LLM に送信せずに、何百万ものメッセージに対してこれを行うのは、当初は困難な部分でした。 ClickHouse では、「会話全体で時間ごとに最後の 50 件のイベントを取得する」と「この会話内のすべてのイベントを取得する」には異なる並べ替え順序が必要なので、キー、パーティション、マテリアライズド ビュー、およびプロジェクションの並べ替えを何度も繰り返す必要がありました。新しいクラスターを見つける場合、LLM 経由ですべてを送信すると、時間がかかりすぎ、コストがかかりすぎます。 HDBSCAN スタイルの埋め込みクラスタリングも、ペアごとの比較のため、大規模になると苦痛が生じます。まず、コサイン ドリフトに基づいて会話をセグメントに分割し、BIRCH を実行して候補空間を圧縮し、次に小さいセットに対して HDBSCAN のようなクラスタリングを使用します。既存のクラスターを照合するために、埋め込み、より小さな分類子/BERT スタイルのモデル、および LLM をあいまいな場合のフォールバックとしてのみ使用します。私たちは

は複数の企業と提携しており、1 日あたり最大 100 万件のチャットと音声メッセージを取り込んでいます。価格は公開されています。Starter は無料、Pro は月額 499 ドル、Enterprise はより多くのボリューム、セキュリティ、保持のニーズに対応します。当社は各顧客のデータをその顧客に対してのみ使用します。当社は SOC 2 Type 1 に準拠しており、Type 2 は進行中であり、SDK は PyPI と npm 上にあります。 HN コミュニティや、チャットまたは音声エージェントを構築している人々からのフィードバックをお待ちしています。現在、これらの信号をどのように検出しているか、どのようなフィードバック方法が機能しているか、これを試すのに妨げになるものは何ですか?喜んで質問に答え、批判を受け付けます。

記事本文:
エージェントの失敗をキャッチする
あなたの評価はミスです。
Agnost AI は本番での会話を継続的に分析し、ユーザーがどこで行き詰まったり、イライラしたり、変換に失敗したりしているかを特定し、最も影響の大きいパターンをレビュー済みの修正に変えてエージェントに提供します。
エージェントの失敗をキャッチする
あなたの評価はミスです。
Agnost AI は本番での会話を継続的に分析し、ユーザーがどこで行き詰まったり、イライラしたり、変換に失敗したりしているかを特定し、最も影響の大きいパターンをレビュー済みの修正に変えてエージェントに提供します。
演出信号・つついてみる
演出信号・つついてみる
あなたの eval は合格します。生産は依然として失敗しています。
Agnost は、実際の会話で見逃していた失敗を見つけて、チームがレビューできる修正に変換します。
「Agnost AI と協力して、包括的な可観測性機能を MCP Toolbox for Databases に統合しました。」
「 Agnost AI は、分析とエラー率の追跡に非常に役立ちます。」
「Agnost AI が実際にコンバージョンにつながった会話の背後にあるパターンを明らかにしてから、当社の Voice BDR は会議の予約において大幅に改善されました。」
「 Agnost AI のおかげで、ユーザーが私たちにはない機能を求めていることがわかりました。私にはまったくわかりませんでした。そのフィードバックはすべて、私たちがすでに行っていた会話の中に留まっていただけでした。」
「 Agnost AI は、エージェントをより良くする方法です。私たちは何を改善し、出荷し、より速く行動するかを検討します。とてもシンプルです。」
「Agnost AI はエージェントの会話に埋もれているバグを発見し、PR を開いて一晩で修正してくれました。」
意図、シグナル、改善
サイレント本番環境での失敗から、修正版の出荷まで。
あらゆる LLM、あらゆるフレームワークで動作します。セットアップは 2 分。 OpenTelemetry ネイティブ。
eval が見逃した失敗を確認します。
接続した瞬間に、Agnost AI は実際の本番環境での会話を読み取り、製品にとって重要な障害カテゴリ (ワークフローの中断、再試行の繰り返し、異常など) を生成します。

摩擦や解約リスクなどが増加します。
無料で始めましょう。準備ができたらスケールします。
会話の取り込みを開始するのが早ければ早いほど、エージェントは賢くなるため、今すぐ接続してください。
意図とセンチメントの信号抽出
障害の検出と解決
迅速に改善する必要がある配送エージェントのチームに最適です。
営業電話をかけずにエージェントのボリュームを拡張する成長チームに最適です。
創設者との無制限の時間
エージェントを大規模に実行している組織に最適です。
Agnost AI は、eval が見逃すエージェントの失敗をキャッチします。実際の運用中のチャットや音声会話を読み取り、ユーザーが行き詰まったりイライラしたりしている箇所を明らかにし、レビュー済みの PR を開いてエージェントを修正します。レビューしてマージします。
Agnost AI は、eval が見逃すエージェントの失敗をキャッチします。実際の運用中のチャットや音声会話を読み取り、ユーザーが行き詰まったりイライラしたりしている箇所を明らかにし、レビュー済みの PR を開いてエージェントを修正します。レビューしてマージします。
Agnost AI は、eval が見逃すエージェントの失敗をキャッチします。実際の運用中のチャットや音声会話を読み取り、ユーザーが行き詰まったりイライラしたりしている箇所を明らかにし、レビュー済みの PR を開いてエージェントを修正します。レビューしてマージします。
Agnost AI はエージェント可観測性ツールとどう違うのですか?
Agnost AI はどの LLM およびフレームワークと連携しますか?
Agnost AI は実際に何を改善しますか?
私の会話データは安全ですか?

## Original Extract

Agnost AI continuously analyzes production conversations, catches agent failures your evals miss, and turns the highest-impact patterns into reviewed fixes. Trusted by Google and 25+ AI teams.

Hey HN, we’re Shubham & Parth, childhood friends building Agnost AI ( https://agnost.ai ), product analytics for teams building chat and voice agents. We read production conversations and find behavioral failures like users rageprompting (cursing at the agent), repeatedly rephrasing the same request, correcting the agent, asking for missing features, or leaving after an answer that was technically successful. We have an interactive demo with no signup here: https://app.agnost.ai?demo=true Here's a demo video: https://www.tella.tv/video/agnost-ai-launch-hn-demo-9haa The core problem is that chat and voice products do not have the same metrics as web apps. When the product interface is language, clicks and funnels become much less useful. Users also rarely give explicit feedback, and when they do it's usually sugarcoated. I barely type /feedback in Claude or Codex myself. Most users just curse, ask again, correct the agent, or leave. So product engineers get technical visibility from latency, errors, and traces, but still have to guess whether users got what they wanted. We got here after building around agents for the last year and got a couple of founders asking for something like a PostHog for conversations for the AI assistants they were building. We are not trying to be in the observability or evals space. Observability tells you what happened technically. Evals validate cases you already know. We're more on the discovery side like what users wanted, where they got frustrated, what they asked for repeatedly, and what new evals should exist. Teams send us agent conversation messages through SDKs or OTel, optionally with metadata like account, plan, source, organization, etc. We cluster conversations into product-specific intents. Feature requests and bugs are default categories; most other clusters are created dynamically from the customer’s data and evolve over time. You can create your own cluster in plain English. If a cluster gets too broad, we split it. If a new pattern appears, we suggest it. One AI video editor company used Agnost AI to find feature requests hidden inside chat. The biggest one was that around 70 users wanted auto-subtitles, but users said it as “add this text in this frame” 12x in a single session, “can you caption it”, “give me transcript of audio” and variations across languages. The team later built the feature. Doing this over millions of messages without sending everything to an LLM was the hard part initially. In ClickHouse, “fetch the last 50 events by time across conversations” and “fetch all events in this conversation” want different sort orders, so we had to iterate a lot on sorting keys, partitions, materialized views, and projections. For finding new clusters, sending everything through an LLM was too slow and expensive. HDBSCAN-style embedding clustering also gets painful at scale because of pairwise comparisons. We first split conversations into segments based on cosine drift, run BIRCH to compress the candidate space, and then use HDBSCAN-like clustering on the smaller set. For matching existing clusters, we use embeddings, smaller classifiers/BERT-style models, and LLMs only as fallback for ambiguous cases. We’re live with multiple companies and ingesting ~1M chat and voice messages per day. Pricing is public: Starter is free, Pro is $499/month, and Enterprise is for higher volume, security, retention needs. We use each customer’s data only for that customer. We are SOC 2 Type 1 compliant, Type 2 is in progress, and our SDKs are on PyPI and npm. We’d love feedback from the HN community and people building chat or voice agents: how do you detect these signals today, what feedback methods have worked, and what would block you from trying this? Happy to answer questions and take criticism.

Catch agent failures
your evals miss.
Agnost AI continuously analyzes production conversations, finds where users get stuck, frustrated, or fail to convert, and turns the highest-impact patterns into reviewed fixes for your agent.
Catch agent failures
your evals miss.
Agnost AI continuously analyzes production conversations, finds where users get stuck, frustrated, or fail to convert, and turns the highest-impact patterns into reviewed fixes for your agent.
Production signal · poke around
Production signal · poke around
Your evals pass. Production still fails.
Agnost finds the missed failures in real conversations and turns them into fixes your team can review.
“ In collaboration with Agnost AI, we have integrated comprehensive observability features into MCP Toolbox for Databases. ”
“ Agnost AI is very useful in tracking analytics and error rates. ”
“ Our Voice BDRs got meaningfully better at booking meetings once Agnost AI surfaced the patterns behind the conversations that actually converted. ”
“ Agnost AI showed me users were asking for features we don't even have. I had no idea. All that feedback was just sitting inside conversations we already had. ”
“ Agnost AI is how we make our agents better. We see what to improve, ship it, and move faster. Simple as that. ”
“ Agnost AI spotted bugs buried in our agent's conversations and opened PRs to fix them overnight. ”
INTENTS · SIGNALS · IMPROVEMENT
From silent production failures to the fix that ships.
Works with any LLM, any framework. 2-min setup. OpenTelemetry native.
See the failures your evals missed.
The moment you connect, Agnost AI reads real production conversations and generates the failure categories that matter to your product: broken workflows, repeated retries, setup friction, churn risk, and more.
Start free. Scale when you're ready.
The earlier you start ingesting conversations, the smarter your agent gets, so connect today.
Intent & sentiment signal extraction
Failure detection & resolution
Best for teams shipping agents that need to get better fast.
Best for growing teams scaling agent volume without a sales call.
Unlimited time with the founders
Best for organizations running agents at scale.
Agnost AI catches agent failures your evals miss. It reads real production chat and voice conversations, surfaces where users get stuck or frustrated, and opens reviewed PRs to fix your agent. You review and merge.
Agnost AI catches agent failures your evals miss. It reads real production chat and voice conversations, surfaces where users get stuck or frustrated, and opens reviewed PRs to fix your agent. You review and merge.
Agnost AI catches agent failures your evals miss. It reads real production chat and voice conversations, surfaces where users get stuck or frustrated, and opens reviewed PRs to fix your agent. You review and merge.
How is Agnost AI different from agent observability tools?
Which LLMs and frameworks does Agnost AI work with?
What does Agnost AI actually improve?
Is my conversation data secure?
