---
source: "https://www.theclimatebrink.com/p/the-real-energy-use-of-agentic-ai"
hn_url: "https://news.ycombinator.com/item?id=49203792"
title: "The energy use of agentic AI"
article_title: "The real energy use of agentic AI - by Zeke Hausfather"
author: "alphabetatango"
captured_at: "2026-08-06T23:50:15Z"
capture_tool: "hn-digest"
hn_id: 49203792
score: 1
comments: 0
posted_at: "2026-08-06T22:59:18Z"
tags:
  - hacker-news
  - translated
---

# The energy use of agentic AI

- HN: [49203792](https://news.ycombinator.com/item?id=49203792)
- Source: [www.theclimatebrink.com](https://www.theclimatebrink.com/p/the-real-energy-use-of-agentic-ai)
- Score: 1
- Comments: 0
- Posted: 2026-08-06T22:59:18Z

## Translation

タイトル: エージェント AI のエネルギー利用
記事のタイトル: エージェント AI の実際のエネルギー利用 - Zeke Hausfather 著
説明: エージェントは単純な AI プロンプトよりも約 600 倍多くのエネルギーを使用します

記事本文:
エージェント AI の実際のエネルギー利用 - Zeke Hausfather 著
気候の瀬戸際
エージェント AI の実際のエネルギー使用量
エージェントは単純な AI プロンプトよりも約 600 倍多くのエネルギーを使用します
Zeke Hausfather 2026 年 8 月 5 日 99 28 25 シェア AI のエネルギー利用は現在、大きな物議を醸しているトピックです。信頼できる推計では、2030 年までに米国の電力使用量の約 12% が AI データセンターによるものになるとされています。しかし同時に、消費者には、自身の AI 使用の影響について、安心させるような小さな数字が与えられており、その数字は、総使用量の驚異的な規模と一見矛盾しているように見えます。
2025年にGoogleは、Geminiのテキストプロンプトの中央値の使用量はわずか0.24ワット時（Wh）で、「9秒間のテレビ視聴」よりも少ないエネルギーであると計算した記事を発表した。同じ頃、Sam Altman 氏は、平均的な ChatGPT クエリは約 0.34 Wh を使用し、Epoch AI も同様の数字を出したと述べました。アンディ・マスリーやハンナ・リッチーなどの作家は、このような割合では個人がチャットボットを使用しても影響はほとんどなく、1回のプロンプトは平均的なアメリカ人の1日あたりの排出量のおよそ15万分の1にすぎないことを示している。
これらの数字は基本的に正しいです。彼らはまた、今日の AI の実際の使用方法からもますます離れています。
ソフトウェア エンジニアや科学者が実際に AI を使用する方法として急速に普及しているのは、チャット ボックスに質問を入力することではありません。むしろ、Claude Code や Codex などのツールを通じて AI エージェントを使用し、計画、コードの作成、実行、結果の読み取り、反復処理を独自に実行します。これらのエージェントは、人間のプロンプトごとに数十のモデル呼び出しを行い、同じ質問に対する複数の回答の試行と評価を含む複雑な推論チェーンに従事します。
私はシリコンバレーの会社 (Stripe) で働いており、最新の AI ツールを明らかに使用しています。

ほとんどの人が。しかし、私は、過去 8 週間にわたる私自身の AI の使用を深く掘り下げ、私が担当していた実際のエネルギー使用量を計算することが有益であると考えました。
過去 8 週間で、私は 1,138 個のプロンプトをクロード コードに入力しました。これらのプロンプトにより 14,000 を超えるモデル呼び出しがトリガーされ、32 億のトークンが処理されました。私の最も正確な推定では、これにより約 170 kWh のデータセンター電力が使用されたと考えられます (方法と仮定の不確実性の範囲はおよそ 70 ～ 330 kWh)。これは、プロンプトあたり約 150 Wh (60 ～ 290 Wh) に相当し、チャット プロンプトのエネルギー中央値の約 600 倍 (250 ～ 1,200) に相当します。 「乗車」が運転の測定値であるのと同様、「プロンプト」は究極的には AI 使用の単位ではありません。どれだけ遠くまで行くかが重要です。
この投稿のきっかけの一部は、企業の AI 排出量会計の標準化されたフレームワークを提案する Watershed の新しいホワイトペーパー (Bistline et al. 2026) の出版です。これは、公開されたクエリごとの数値が桁違いに異なる理由 (主にシステム境界) について、私がこれまでに見た中で最も慎重に扱ったものであり、議論全体を再構成する必要がある図が含まれています。AI タスクあたりの電力量は、テキスト分類の 1,000 分の 1 ワット時から、5 ～ 50 回のフロンティア モデル呼び出しを行うエージェント ワークフローの 50 ～ 500 Wh まで、5 桁以上に及びます。彼らの言葉を借りれば、1 つの「インタラクション」による排出量は、実際に消費されるコンピューティング量を「一桁以上」過小評価している可能性があります。
他の研究者も同様の結果を発見しました。バイら。 (2026) は、実際のソフトウェア タスクでコーディング エージェントを測定し、通常のチャットボット インタラクションのおよそ 1,000 倍のトークンを消費することを発見しました。そして、この種のエージェントのタスクは、AI 使用量の増加を最も急速に促進する要因となります。アントロピックのエコノ

mic Index は、API 使用量の 97% がエージェントに関連する「自動化が優勢な」パターンを示していることを発見しました。
これらの値を客観的に見るために、以下の図は、公開されているプロンプトおよびタスクごとの推定値 (青) と、クロード コードの使用から測定した値 (オレンジ色) およびエネルギー使用量の一般的なベンチマーク (電子レンジ、冷蔵庫、または家全体の稼働) を比較しています。
AI タスクごとの電力消費量。公開されている推定値 (青) と私自身のクロード コード セッション ログから計算された値 (オレンジ) を含みます。 Bistline (2026) アクティビティ層エネルギー係数を使用して変換された測定トークン数。オレンジ色の範囲は、キャッシュ読み取りエネルギーの想定 1% ～ 25% です。私のクロード コード セッションの使用量の中央値は約 0.6 kWh (0.25 ～ 1.2 kWh) で、これは Watershed の一般的なエージェント使用量推定の上限にあたり、携帯電話の充電に使用されるエネルギーの 50 倍です。クロード コードを使用する私の平均的な 1 日 (3.0 kWh、範囲 1.2 ～ 5.9 kWh) は、冷蔵庫 2 台を稼働させるよりも多くの電力を消費します。
Claude Code は、モデル呼び出しごとに API がレポートする正確なトークン数を含む、すべてのセッションの完全なローカル トランスクリプトを保持します。 1 これにより、公開されているベンチマークから単純に推定するのではなく、自分が担当していた AI の使用量を正確に知ることができます。想定を必要とする、使用されるトークンをエネルギーに変換する唯一のステップです。
最初に気づいたのは、「プロンプト」と現実の間には大きなギャップがあるということです。1,138 個のプロンプトを入力した結果、14,000 個をわずかに超える個別のモデル呼び出し (プロンプトごとに 12 個) が発生し、各プロンプトは平均 290 万トークンを消費しました。比較のために言うと、推論や Web 検索を行わない一般的な Web ベースの AI チャット交換では、約 1,000 トークンしか使用されません。
過去 8 週間で、私のクロード コードは 32 億トークンを使用しました。これらの圧倒的な部分は、エージェントが自身の作品を再読したことによるものです。

懐かしい思い出。エージェントがステップを実行するたびに (コマンドの実行、ファイルの読み取り、ツールの呼び出しなど)、モデルは蓄積されたコンテキスト全体を再処理します。以下の図は、トークンの使用方法の内訳と総電力使用量に占めるトークンの割合を示しています。
2026 年 5 月 31 日から 7 月 25 日までの私のクロード コード使用量のトークンと推定電力構成。「キャッシュ読み取り」は、モデル呼び出しごとにキーと値のキャッシュから事前に処​​理されたコンテキストを再読み取りしたものです。 「キャッシュ書き込み」は、処理および保存される新しいコンテキストです。 「出力」とは、モデルによって生成されたテキストとコードです。電力シェアは、新規入力エネルギー率の 10% でキャッシュ読み取りを行う Bistline (2026) 係数を使用します。実際に表示されるテキスト (モデルの出力など) は、処理されたトークン全体の約 0.4% にすぎません。トークンの約 96% はキャッシュ読み取りであり、エージェントは 14,000 ステップごとに自身のコンテキストを再読み取りします。キャッシュされたトークンは、新しいトークンを処理するよりも再読み取りの方がはるかに安価であるため、これはエネルギーの見積もりにとって非常に重要です。 AI 企業は、フレッシュ コンテンツと比較してキャッシュ読み取りの価格の約 10% を請求します。私はその比率をエネルギーの中心的な想定として使用し、1% と 25% を境界とします。 2
フロンティアモデルの本当のトークン当たりのエネルギーを実際に知っている人は研究室の外に誰もいないので（Anthropic はプロンプト当たりやトークン当たりの数値を公表していないが、ウォーターシェッドの論文は丁寧だがこの分野最大のデータギャップであると断固として警告している）、私は 3 つの独立した公開された方法論を通してトークン数の測定を実行した：ウォーターシェッドのアクティビティ層要因、トークン当たりの要因、Epoch AI の作業から導き出された Simon Couch の推定、およびクロードカーボンツールの価格推定係数。
公開されている 3 つの方法論に基づいて使用した 32 億トークンの推定電力消費量は次のとおりです。

3 つのキャッシュ読み取り仮定に基づく流域アクティビティ層係数、Couch (2026) のトークンごとの係数、およびモデルごとのクロード炭素係数。これらのどの方法論でも、8 週間でおよそ 70 ～ 330 キロワット時の間の答えが得られます。この推定値は、どちらの方向でも最大 2 倍不確実です。しかし、より広範な結論はそうではありません。公開されているチャット プロンプトごとのレートで私の 1,138 件のプロンプトを数えると、約 0.3 kWh になることが示唆されますが、実際はその 150 ～ 1,200 倍です。
私の毎日のエネルギー使用パターンは下の図に示されています。日ごとの変動は大きく、私の最も負荷が高かった日 (中央推定値 11 kWh) には、大規模な地理空間分析を通じて複数の並列エージェントが稼働し、米国の平均的な家庭の 1 日あたりの総電力量の 3 分の 1 以上を使用しました。これは、エージェント使用のカテゴリ内であっても、タスクの複雑さと同時に使用されるサブエージェントの数が、結果として生じるエネルギー使用に大きく影響するという事実を反映しています。
クロード コードの使用による毎日の推定電力消費量 (バー: 入力エネルギーの 10% でのキャッシュ読み取り、ひげ: 1% ～ 25%)。基準線は、冷蔵庫と米国の平均的な家庭の典型的な毎日の電力使用量を示しています。私の数値は、他の公表されている薬剤使用の推定値よりも少し高いですが、その理由を判断するには少し掘り下げてみる価値があります。 Couch 氏は、Claude Code セッションの中央値では、24 のモデル呼び出しと 592,000 のトークンが含まれ、約 41 Wh を使用すると推定しました。 Andy Masley の 2026 年 6 月の計算では、100,000 トークンの Claude Opus エージェント セッションは約 459 Wh になります。私のセッションの中央値は約 600 Wh で、100 回以上のコールと、大規模データ分析プロジェクトの多数のサブエージェントを含む約 1,000 万のトークンが関係します。ハンナ・リッチーの仮説上のヘビー ユーザー (1 日あたり 24 のエージェント クエリ) が判明

実際の使用量の中心推定値は 3.0 kWh/日 (1.2 ～ 5.9 kWh) でしたが、2.4 kWh/日でした。
これらの見積もりはいずれも必ずしも間違っているわけではなく、実際の使用状況の幅広い想定を反映しているだけです。ソフトウェア エンジニア、研究者、データ アナリスト (たとえば、私のような人々) は、おそらく使用量分布のかなり下の方に位置しています。同時に、より複雑なエージェント ツールがますます標準になるため、使用量は時間の経過とともに増加する可能性があります。
今年はどんな年になるだろうか
この 8 週間がかなり一般的であると仮定すると、私がエージェント クロード コードを 1 年間使用すると、データセンターの電力は約 1.1 MWh (0.4 ～ 2.2 MWh) 消費されると推定できます。これは、米国の平均的な家庭が使用する電力の約 10 分の 1 です。米国の平均グリッド強度を適用すると、年間約 370 kgCO2e (150 ～ 730 kgCO2e) になります。 3
一般的なアクティビティの年間排出量を、私の年間換算したクロード コード使用量と比較しました。車: EPA の典型的な乗用車 (22.2 mpg、11,500 マイル/年)。 EV: カリフォルニアの電力網で 0.30 kWh/マイルで 11,500 マイル/年。フライト: ICAO 方式のエコノミー往復、CO2 のみ。家庭用電力: 米国の平均送電網に接続されている EIA の米国の平均世帯。乾燥機: 米国の平均送電網で最大 770 kWh/年 (DOE) の典型的な電気衣類乾燥機。現在、私の個人的および職業上の AI 使用による年間排出量は、電気衣類乾燥機を稼働させるよりも若干多く、カリフォルニアで電気自動車を 11,500 マイル運転するか、エコノミーでサンフランシスコとニューヨークを往復する場合の約半分です。 4 これは典型的なアメリカのガソリン車の年間排出量の約 8% に相当し、平均的なアメリカ人の年間温室効果ガス排出量約 18 トンの約 2% に相当します。
これは大規模な排出源であると同時に、私の総二酸化炭素排出量の比較的小さな部分でもあります。私は通常往復飛行機に乗ります

年に2回、年老いた両親を訪ねるためにサンフランシスコから東海岸まで行きますが（仕事の旅行はもちろんのこと）、その選択で眠れなくなることはほとんどありません。また、基本的に航空よりもはるかに脱炭素化が容易な最終用途でもあります (詳細は下記を参照)。しかし、地球の気温が急上昇し、排出削減目標がますます軌道から外れている現在、これは正味新たな排出源でもある。
この投稿の大部分を費やして、エージェント AI の使用はチャットボットの数値が示すよりも何百倍もエネルギーを消費するという主張をしてきましたが、答えは罪悪感や禁欲ではないと考えていることを明確にしておきます。しかし、ここには、今後の AI のエネルギー使用と排出の軌道を形作るために使用できる本当の手段があります。
個人的な面では、エージェント ツールを軽薄に使用しないように努めることができます。 5 人の並行エージェントに難しい研究問題を突きつけることと、バーの賭けを解決するために同じことを行うこと (または、私の場合、娘と一緒にアホロートルをテーマにしたゲームを作ること) の間には、実際の違いがあります。どのようなモデルを使用するかも重要です。単純なタスクを小規模なモデルに送信すると、フロンティア モデルをデフォルトで使用する場合に比べて、トークンあたりのエネルギー使用量がおそらく 5 ～ 7 倍少なくなります 5。検索や機械的な作業では、これを行うことが増えています。とはいえ、過大評価はしたくない

[切り捨てられた]

## Original Extract

Agents use about 600x more energy than simple AI prompts

The real energy use of agentic AI - by Zeke Hausfather
The Climate Brink
Subscribe Sign in The real energy use of agentic AI
Agents use about 600x more energy than simple AI prompts
Zeke Hausfather Aug 05, 2026 99 28 25 Share AI energy use is a huge and controversial topic at the moment. Credible estimates have AI data centers accounting for around 12% US electricity use by 2030. But at the same time consumers have been given reassuringly small numbers about the impact of their own AI use, numbers that seem on their face somewhat inconsistent with the staggering size of their aggregate usage.
In 2025 Google published an article calculating that median Gemini text prompt used only 0.24 watt-hours (Wh), less energy than “watching nine seconds of television”. Around the same time, Sam Altman said that an average ChatGPT query uses about 0.34 Wh, and Epoch AI came out with similar numbers . Writers like Andy Masley and Hannah Ritchie have shown that at these rates an individual using chatbots has a pretty negligible impact, with one prompt only amounting to roughly 1/150,000th of an average American’s daily emissions.
Those numbers are basically right. They are also increasingly divorced from how AI is actually being used today.
The fastest-growing way that software engineers and scientists actually use AI is not typing questions into a chat box. Rather, we use AI agents through tools like Claude Code and Codex that plan, write code, run it, read the results, and iterate on their own. These agents make dozens of model calls per human prompt, and engage in complex reasoning chains that involve attempting and evaluating multiple answers to the same question.
I work for a company in Silicon Valley ( Stripe ) and admittedly use the latest AI tools more than most people. But I thought it would be instructive to take a deep dive into my own AI use over the past 8 weeks and calculate the actual energy use I was responsible for.
Over the past 8 weeks I typed 1,138 prompts into Claude Code. Those prompts triggered more than 14,000 model calls that processed 3.2 billion tokens. My best estimate is that this used around 170 kWh of data center electricity (with an uncertainty range of roughly 70 to 330 kWh across methods and assumptions). That works out to around 150 Wh per prompt (60 to 290 Wh), which is roughly 600 times (250 to 1,200) the energy of a median chat prompt. A “prompt” is ultimately not a unit of AI use any more than “trips” is a measurement of driving; it’s how far you go that matters.
Part of the impetus for this post is the publication of a new white paper from Watershed ( Bistline et al. 2026 ) proposing a standardized framework for corporate AI emissions accounting. It is the most careful treatment I have seen of why published per-query numbers differ by orders of magnitude (system boundaries, mostly), and it contains a figure that should reframe the whole discussion: electricity per AI task spans more than five orders of magnitude, from thousandths of a watt-hour for text classification to 50-500 Wh for an agentic workflow making 5-50 frontier model calls. As they put it, emissions attributed to one “interaction” may understate the compute actually consumed “by an order of magnitude or more.”
Other researchers have found similar results. Bai et al. (2026) measured coding agents on real software tasks and found they consume roughly 1,000 times the tokens of an ordinary chatbot interaction. And these sort of agents tasks represent the most rapid driver of increased AI usage; Anthropic’s Economic Index found that 97% of their API usage now show “automation-dominant” patterns associated with agents.
To put these values in perspective, the figure below compares published per-prompt and task estimates (blue) with what I measured from my Claude Code use (orange) as well as common benchmarks for energy use (running a microwave, a fridge, or a whole home):
Electricity consumption per AI task, including published estimates (blue) and values computed from my own Claude Code session logs (orange). Measured token counts converted using Bistline (2026) activity-tier energy factors; orange ranges span cache-read energy assumptions of 1% to 25%. My median Claude Code session uses around 0.6 kWh (0.25 to 1.2 kWh), which is at the top end of Watershed’s generic agentic usage estimate, and fifty times the energy used to charge a cellphone. My average day of Claude Code (3.0 kWh, range 1.2 to 5.9 kWh) uses more electricity than running two refrigerators.
Claude Code keeps complete local transcripts of every session, including the exact token counts the API reports for every model call. 1 This lets me precisely know how much AI usage I was responsible for rather than simply extrapolating it from published benchmarks; its only the step to convert tokens used to energy that requires assumptions.
The first thing I found is that the gap between “prompts” and reality is massive: my 1,138 typed prompts resulted in just over 14,000 distinct model calls (12 per prompt), and each prompt consumed on average 2.9 million tokens. For comparison, typical web-based AI chat exchanges with no reasoning or web searches only use around a thousand tokens.
Over the past the 8 weeks, my Claude Code used 3.2 billion tokens. These overwhelmingly came from the agent re-reading its own working memory. Every time an agent takes a step (e.g. runs a command, reads a file, or calls a tool), the model re-processes its entire accumulated context. The figure below shows the breakdown of how tokens were used and their share of total electricity use.
Token and estimated electricity composition of my Claude Code usage, May 31 to July 25, 2026. “Cache reads” are previously processed context re-read from the key-value cache on each model call; “cache writes” are new context being processed and stored; “output” is text and code generated by the model. Electricity shares use Bistline (2026) factors with cache reads at 10% of the fresh-input energy rate. The text I actually see (e.g. the model’s output) is only around 0.4% of total tokens processed. Some 96% of the tokens are cache reads where the agent re-reads its own context at each of those 14,000 steps. This matters enormously for the energy estimate, because a cached token is much cheaper to re-read than a fresh one is to process. AI companies charge about 10% of the price for cache reads compared with fresh content, and I use that ratio as my central energy assumption, with 1% and 25% as bounds. 2
Since nobody outside of the labs actually knows the true per-token energy of a frontier model (Anthropic has published no per-prompt or per-token figures, something the Watershed paper politely but firmly flags as the field’s biggest data gap), I ran my measured token counts through three independent published methodologies: Watershed’s activity-tier factors , the per-token factors Simon Couch ’s estimates derived from Epoch AI’s work, and the claude-carbon tool’s pricing-inferred coefficients.
Estimated electricity consumption for the 3.2 billion tokens I used under three published methodologies: Watershed activity-tier factors under three cache read assumptions, Couch (2026) per-token factors, and claude-carbon per-model coefficients. Every one of these methodologies gives an answer between roughly 70 and 330 kilowatt-hours over 8 weeks. The estimate is genuinely uncertain, by a factor of ~2 in either direction. But the broader conclusion is not: counting my 1,138 prompts at published per-chat-prompt rates would have suggested about 0.3 kWh, while the reality is 150 to 1,200 times that.
My daily pattern of energy use is shown in the figure below. The day to day variability is huge: my heaviest day (11 kWh central estimate) involved multiple parallel agents churning through a large geospatial analysis, and used more than a third of the total daily electricity of an average US home. This reflects that fact that even within the category of agentic usage, the complexity of the task and the number of simultaneous sub-agents used will greatly influence the resulting energy use.
Estimated daily electricity consumption of my Claude Code usage(bars: cache reads at 10% of input energy; whiskers: 1% to 25%). Reference lines show typical daily electricity use of a refrigerator and of an average US household. My numbers are a bit higher than some of the other published estimates of agentic use, and it is worth digging in a bit to determine why. Couch estimated that a median Claude Code session uses around 41 Wh, involving 24 model calls and 592k tokens. Andy Masley’s June 2026 calculator puts a 100k-token Claude Opus agent session at ~459 Wh. My median session is ~600 Wh, involving a hundred-plus calls and around ten million tokens including numerous subagents for large data analyses projects. Hannah Ritchie’s hypothetical heavy user (24 agentic queries a day) came out at 2.4 kWh/day, while I measured a central estimate of 3.0 kWh/day (1.2 to 5.9 kWh) for my actual usage.
None of these estimates are necessarily wrong, they just reflect a wide range of actual usage assumptions. Software engineers, researchers, and data analysts (e.g. folks like me) probably lie pretty far down the tail of the usage distribution. At the same time, usage will likely grow over time as more complex agentic tools increasingly become the norm.
What a year of this looks like
If we assume that these 8 weeks are fairly typical, we can estimate that a full year of my agentic Claude Code use would consume roughly 1.1 MWh of data center electricity (0.4 to 2.2 MWh), which is about a tenth of what an average US household uses. Applying the US-average grid intensity, that is roughly 370 kgCO2e per year (150 to 730 kgCO2e). 3
Annual emissions of common activities compared with my annualized Claude Code usage. Car: EPA typical passenger vehicle (22.2 mpg, 11,500 mi/yr). EV: 11,500 mi/yr at 0.30 kWh/mi on the California grid. Flight: ICAO-method economy round trip, CO2 only. Home electricity: EIA average US household on the US-average grid. Dryer: typical electric clothes dryer at ~770 kWh/yr (DOE) on the US-average grid. My personal and professional AI usage now emits a bit more per year than running an electric clothes dryer, and about half as much as driving an electric car 11,500 miles in California or taking one San Francisco to New York round-trip flight in economy. 4 It is about 8% of the annual emissions of a typical American gasoline car, and roughly 2% of the average American’s ~18-ton annual greenhouse gas footprint.
This is simultaneously a large emissions source and a relatively modest part of my total carbon footprint. I typically take a round trip flight from San Francisco to the East Coast twice a year to visit my aging parents (not to mention work travel), and I generally don’t lose sleep over that choice. It is also fundamentally a much easier-to-decarbonize end-use than aviation (more on that below). But this also represents a net new source of emissions, at a time when global temperatures are skyrocketing and our emissions reduction goals are increasingly off track.
Having spent most of this post arguing that agentic AI use is hundreds of times more energy intensive than the chatbot numbers suggest, let me be clear that I don’t think the answer is guilt or abstinence. But there are real levers here that we can use to shape the trajectory of AI energy use and emissions going forward.
On the personal side we can try and not be frivolous with agentic tools. There is a real difference between pointing five parallel agents at a hard research problem and doing the same to settle a bar bet (or, in my case, making axolotl-themed games with my daughter). What models you use matters too: sending simple tasks to smaller models uses perhaps 5 to 7 times less energy per token than defaulting to a frontier model, 5 and it is what I increasingly do for searches and mechanical work. That said, I don’t want to oversel

[truncated]
