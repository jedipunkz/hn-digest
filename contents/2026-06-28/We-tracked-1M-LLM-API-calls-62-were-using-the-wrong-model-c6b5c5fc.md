---
source: "https://tokonomics.ca/blog/we-tracked-1m-llm-api-calls-most-were-wasting-money"
hn_url: "https://news.ycombinator.com/item?id=48708666"
title: "We tracked 1M LLM API calls – 62% were using the wrong model"
article_title: "We Tracked 1M LLM API Calls — 60% Were on the Wrong Model â”€ Tokonomics"
author: "aitoukhrib"
captured_at: "2026-06-28T16:25:35Z"
capture_tool: "hn-digest"
hn_id: 48708666
score: 3
comments: 0
posted_at: "2026-06-28T16:13:23Z"
tags:
  - hacker-news
  - translated
---

# We tracked 1M LLM API calls – 62% were using the wrong model

- HN: [48708666](https://news.ycombinator.com/item?id=48708666)
- Source: [tokonomics.ca](https://tokonomics.ca/blog/we-tracked-1m-llm-api-calls-most-were-wasting-money)
- Score: 3
- Comments: 0
- Posted: 2026-06-28T16:13:23Z

## Translation

タイトル: 100 万件の LLM API 呼び出しを追跡しました – 62% が間違ったモデルを使用していました
記事のタイトル: 100 万件の LLM API 呼び出しを追跡しました — 60% が間違ったモデルを使用していました – Tokonomics
説明: ほとんどの開発者は AI に対して 10 ～ 18 倍の過剰な支払いを行っています。 100 万回の通話で数値を実行しました。これが私たちが見つけたものです。

記事本文:
100 万件の LLM API 呼び出しを追跡しました - 60% は間違ったモデル – Tokonomics を使用していました
トコノミクス
特長
ツール
トークンカウンター
コスト計算ツール
プロンプトオプティマイザー
APIリクエストビルダー
モデルマトリックス
ROI 計算ツール
特長
仕組み
LLMの価格設定
比較する
ツール
トークンカウンター
コスト計算ツール
プロンプトオプティマイザー
APIリクエストビルダー
モデルマトリックス
ROI 計算ツール
ブログ
ログイン
無料で始める
ブログ
llm-コストの最適化
ai-API-価格設定
モデルルーティング
2026 年 6 月 8 日
12 分で読めます
100 万件の LLM API 呼び出しを追跡しました - 60% が間違ったモデルを使用していました
2026 年には、開発者の 82% が OpenAI GPT モデルをデフォルトで使用しています ( Stack Overflow Developer Survey、2025 年) が、運用 API 呼び出しの 60 ～ 70% はフロンティア モデルを必要としない単純なタスクです。
分類呼び出しと抽出呼び出しを GPT-4o から DeepSeek V3 に切り替えると、入力トークンが 18 倍節約されます (100 万あたり 2.50 ドル → 0.14 ドル)。
モデル ルーティングとプロンプト キャッシュを組み合わせると、LLM の総支出が 80 ～ 95% 削減されます。
2025 年の平均月間 AI 支出は 1 社あたり 85,500 ドルに達し、前年比 36% 増加しました (CloudZero、2025)。
AI 機能を現在出荷している場合に気になる点があります。
私たちは、Tokonomics を通じて送信された最初の 100 万件の API 呼び出しを、47 のテナント、9 つのプロバイダー、数十のモデルにわたって調べました。このパターンはほぼどこでも同じでした。チームはすべてにおいてデフォルトで GPT-4o を使用していました。カスタマーサポートのチャットボット? GPT-4o。 JSON抽出? GPT-4o。 ５つに分類すると？ GPT-4o。これは AI 開発の SELECT * です。それは機能し、誰も疑問を持たず、静かにお金を生み出しています。
無駄は理論的なものではありません。これは毎月請求ダッシュボードに表示されますが、ほとんどのチームはその存在に気づきません。
開発者の 82% がデフォルトで GPT-4o を使用するのはなぜですか?
2025 年、Stack Overflow の年次開発者調査では、開発者の 82% が OpenAI GPT モデルを次の目的で使用していることがわかりました。

彼らの AI 作品 (Stack Overflow、2025)。これにより、GPT-4o が事実上の標準になり、人々がコードに貼り付けて決して変更されないモデルになります。
それは理にかなっています。 OpenAI には最高のドキュメントがあります。すべてのチュートリアルでは GPT-4o を使用します。深夜にプロトタイプを作成しているときは、6 つのプロバイダーにわたってベンチマークを実行しているわけではありません。機能するものを掴むのです。
しかし、ここに問題があります。プロトタイピングの習慣が制作コストになるのです。 2 月に機能をテストするために選択したモデルは、6 月になっても運用環境で稼働しており、1 日あたり 50,000 件の呼び出しを処理しています。0.14 ドル/M のモデルが 2.50 ドル/M のモデルと同じ結果が得られるかどうかは誰も疑問に思っていません。
私たちの発見: Tokonomics を構築したとき、私たち自身の内部チャットボットは、誰かがチェックする前に 3 か月間 GPT-4o 上で実行されました。 FAQ 部分を GPT-4o-mini に切り替えると、評価セットで測定可能な品質の違いはなく、そのコンポーネントのコストが 94% 削減されました。
これは私たちに限ったことではありません。 2026 年、Divyam.AI は、まさにこのパターンを説明するために「LLMflation」という用語を作りました。チームは、より安価な代替モデルが追いついた後も、長い間レガシー モデルの選択に固執します ( Divyam.AI 、2026)。慣性は本物です。
モデルの選択に実際にかかる費用はいくらですか?
パーセンテージで話すのはやめて、実際のお金を見てみましょう。ここでは、呼び出しごとに平均 500 個の入力トークンと 200 個の出力トークン (分類、抽出、または短い形式の生成のための一般的な運用ワークロード) を想定した場合に、モデル全体で 100 万件のリクエストにかかるコストを示します。
100 万リクエストの月額コスト (500 入力トークン / 200 出力トークン)
GPT-4o
クロード・ソネット 4
クロード俳句 3.5
GPT-4o-ミニ
ディープシーク V3
GPT-4.1ナノ
3,250ドル
4,500ドル
1,200ドル
$195
$126
130ドル
出典: OpenAI、Anthropic、DeepSeek 公式価格ページ、2026 年 6 月
出典: OpenAI、Anthropic、DeepSeek 公式価格ページ、2026 年 6 月
これは、GPT-4o と GPT-4.1 Nano のコスト差が 25 倍です。同じ百万件のリクエストに対して。
アル

それらの通話にはすべて GPT-4o が必要だと思いますか?なぜなら、私たちのデータでは、60〜70％の人がそうしていないからです。
フロンティア モデルを必要としない API 呼び出しはどれですか?
2026 年、複数の運用オペレーターは、典型的な SaaS アプリの API 呼び出しの 60 ～ 70% が、分類、短い要約、構造化された抽出、ルーティングの決定など、予算モデルにとって十分単純であると報告しました (Prem AI、2026)。これは、トコノミクスを通じて私たちが見てきたことと一致します。
実際にどのように分解されるかは次のとおりです。
予算モデルに送信します ($0.10 ～ $0.80/M 入力):
意図の分類 (「これは返金リクエストですか、それとも一般的な質問ですか?」)
テキストからの JSON/構造化データの抽出
短い要約（200ワード未満）
コンテンツの管理/安全性チェック
フロンティア モデルを維持します (入力 100 万ドルあたり 2.50 ドルから 3.00 ドル):
複雑なコードの生成とデバッグ
品質が重要な長編コンテンツの作成
ベンチマークを実施し、品質ギャップが重要であることを確認したタスク
不快な真実?ほとんどのチームはベンチマークを実施しません。彼らは、GPT-4o-mini または DeepSeek が許容可能な結果を​​与えるかどうかをテストしたことがないため、GPT-4o が必要であると想定しています。 LLM オペレーターからのフィールドレポートによると、本番アプリのトークン予算の 40 ～ 60% は純粋な無駄であり、予算モデルの作業を行うフロンティア モデルに費やされています ( エディトリアルGE 、2026)。
企業は実際にいくら支出しているのでしょうか?
2025 年に、CloudZero は従業員数 250 ～ 10,000 人の企業に勤める米国のソフトウェア エンジニア 500 人を対象に調査を実施しました。 AI の月額平均支出は 63,000 ドルから 85,500 ドルに急増し、前年比 36% 増加しました。そして現在、組織の 45% が AI に月額 10 万ドル以上を支出する予定であり、これは 2024 年に同様の回答をした 20% の 2 倍以上です (CloudZero、AI コストの現状、2025 年)。
問題は、AI ROI を自信を持って評価できる組織は 51% だけであるということです。 AI に月に 6 桁を費やしている企業の半数は判断できない

それが価値があるかどうかはあなた次第です。
エンタープライズ AI 支出は、トークンあたりのコストが劇的に低下したにもかかわらず、2024 年から 2025 年の間に 115 億ドルから 370 億ドルへと 320% 増加しました (BuildMVP Fast、2026)。このギャップ、つまり請求額が上昇する一方で価格は下落していることは、最適化よりも使用量の拡大が速いことを示す明らかな兆候です。
私たちの調査結果: LLM に最も多くの費用を費やしているチームは、最も洗練された AI 機能を備えているチームではありません。彼らは、AI 機能を早期にリリースし、モデルの選択を一度も再検討せず、自動操縦で使用量を拡大させた人たちです。私たちが Tokonomics を構築するきっかけとなった 47,000 ドルの請求書は、まさにこのパターンから来ました。
なぜ価格はこれを解決するのに十分な速さで下落しないのでしょうか?
彼らは異常な速度で減少しています。 2025 年に Epoch AI が LLM 推論価格を追跡したところ、年間 50 倍の中央値で下落していることがわかりました。 2024 年 1 月以降、各機能レベルで入手可能な最も安価なモデルの減少率は年間 200 倍に加速しました (Epoch AI、2025)。
LLM 推論の価格は下落しているが、請求額は上昇し続けている
価格/トークン
総支出額
2022年
2024年
2025年
2026年
出典: Epoch AI + CloudZero、2025-2026
出典: Epoch AI + CloudZero、2025-2026
しかし、使用量の増加は価格の下落を上回っています。モデル API への支出は 2024 年末から 2025 年半ばまでに 35 億ドルから 84 億ドルへと倍増し、72% の企業が LLM への支出をさらに増やす予定です (CloudZero、2026)。
いいえ、価格が下がるのをただ待っていて、請求額が自動的に修正されることを期待することはできません。 GPT-4o の価格が現在の DeepSeek の価格と同じになる頃には、さらに新しいフロンティア モデルで 10 倍以上の呼び出しを実行することになります。
修正: タスクごとのルート、パターンごとのキャッシュ、予算ごとの上限
複合する 3 つの動きがあります。ほとんどの記事ではそのうちの 1 つを取り上げています。 3つすべてを重ねるとこうなります。
手順 1: 呼び出しを適切なモデルにルーティングする
すべての API ca にタグを付ける

タスクの種類ごとに異なります。各呼び出しが何を行うかを確認したら、次のようにルーティングできます。
分類と抽出 → GPT-4o-mini または DeepSeek V3
会話サポート → Claude Haiku 3.5
複雑な推論 → GPT-4o または Claude Sonnet 4
通話の 60% が GPT-4o から低予算モデルに移行した場合、それらの通話だけで 10 ～ 15 倍のコスト削減になります。月額 3,250 ドルの場合、約 1,950 ドルの節約になります。
段階的なセットアップについては、機能ごとの LLM コスト追跡に関するガイドを参照してください。
2026 年には、Anthropic のプロンプト キャッシュにより、キャッシュされた入力トークンが 90% 節約されます。 OpenAI の自動キャッシュは、コード変更なしで 50% を節約します。システム プロンプトを繰り返し送信する場合にのみ機能します ( Anthropic 、 OpenAI 、 2026)。
アプリが呼び出しごとに同じシステム プロンプトを送信する場合 (ほとんどのチャットボットはそうします)、プロバイダーがすでに処理したトークンの全額を支払っていることになります。それは燃えているお金です。
OpenAI と Anthropic のプロンプト キャッシュ ガイドでは、実装の詳細が説明されています。
行動 3: 厳しい支出上限を設定する
ルーティングとキャッシュを使用しても、バグは発生します。終了条件のない再試行ループ。同じデータを 2 回処理するバッチ ジョブ。一晩の暴走電話が 1 か月の最適化を元に戻す可能性があります。
毎月の予算の上限を設定し、上限に達すると API 呼び出しをブロックします。午前9時に読むようなアラートではなく、誰も見ていない午前3時に出血を止めるハードブロックです。
支出上限が実際にどのように機能するかは次のとおりです。
3 つの動きをすべて積み重ねても、直線的に加算されるわけではありません。それは以下を複合させます：
モデル ルーティングのみ: 50 ～ 70% の節約
プロンプト キャッシュを追加: さらに 30 ～ 50% が上乗せされます (残りの費用に適用されます)
予算の上限を追加: 貯蓄をすべて消し去る 100% 超過を防ぎます。
すべての GPT-4o トラフィックに月額 3,250 ドルを費やすチームは、同じ出力品質で現実的に月額 300 ～ 650 ドルに到達する可能性があります。それは丸め誤差ではありません。それがあなたの次の雇用者です

毎月のツールの予算。
アーキテクチャを書き直す必要はありません。ここから始めてください:
現在のモデルの使用状況を監査します。機能とモデルごとに呼び出しにタグを付けます。 GPT-4o-mini が処理できる作業に GPT-4o を使用している通話を調べます。まだ可視性がない場合は、Tokonomics が 5 分以内にこれを表示します。
通話量の多い上位 3 つの通話タイプを選択してください。単純な評価を使用して、より安価なモデルでテストを実行します (出力は依然として品質基準を満たしていますか?)。ほとんどのチームは、3 つのうち 2 つがうまく機能していると感じています。モデルを切り替える前に、プロンプト コスト オプティマイザーを使用して、肥大化したプロンプトをトリミングします。
予算アラートを 80% に設定します。何かを最適化する前に、コストがいつ急増するかを必ず把握してください。設定方法は次のとおりです。
今週は通話タイプを 1 つ切り替えます。一度にすべてを最適化しようとしないでください。最も量が多く、複雑性が最も低い通話を予算モデルに移行します。 1週間測定します。それから次のものを移動します。
目標は、どこでも最も安価なモデルを使用することではありません。それは、各タスクに適切なモデルを使用し、それを必要としない作業にフロンティア価格を支払うのをやめるということです。
DeepSeek V3 は GPT-4o と比べてどれくらい安いですか?
DeepSeek V3 のコストは 100 万入力トークンあたり 0.14 ドルですが、GPT-4o の 2.50 ドルは、入力では 18 倍、出力では 36 倍安くなります (0.28 ドル対 10.00 ドル)。 500 入力 + 200 出力トークンで 100 万リクエストの場合、GPT-4o の料金は月額 3,250 ドルですが、DeepSeek の料金は月額 126 ドルです。品質のギャップは複雑な推論では重要ですが、分類や抽出のタスクでは無視できます。
LLM API 呼び出しの何パーセントが安価なモデルを使用できますか?
2026 年には、複数の運用オペレーターが、典型的な SaaS アプリケーションの API 呼び出しの 60 ～ 70% が予算モデルにとって十分単純であると報告しています (Prem AI、2026)。これらには、分類、構造化データの抽出、短い要約、センチメント分析が含まれます。残りの 30 ～ 40% は依然としてフロンティアの恩恵を受けています

GPT-4o や Claude Sonnet 4 などのモデル。
プロンプト キャッシュにより LLM コストはどのくらい節約できますか?
Anthropic のプロンプト キャッシュにより、キャッシュされた入力トークンが 90% 節約されます。 OpenAI の自動キャッシュは、コード変更なしで 50% を節約します ( Anthropic 、 OpenAI 、 2026)。モデル ルーティングとキャッシュを組み合わせると、LLM の総支出を 80 ～ 95% 削減できます。
2025 年の企業の月平均 LLM API 支出はいくらですか?
AI の月間平均支出は 1 社あたり 85,500 ドルに達し、前年比 36% 増加しました (CloudZero、2025 年)。また、組織の 45% が AI に月額 10 万ドル以上を支出する予定であり、2024 年にはその割合の 2 倍以上になります。
各 API 呼び出しが使用する LLM モデルを追跡するにはどうすればよいですか?
アプリと LLM プロバイダーの間にある Tokonomics のような計測プロキシを使用します。カスタム ヘッダーを使用して、各呼び出しにメタデータをタグ付けします。プロキシはモデル、トークン数、呼び出しごとのコストを記録し、コードを変更せずに機能ごとの内訳を提供します。 5 分でセットアップする方法は次のとおりです。
すべてのソースは 2026 年 6 月に取得されました。価格データは公式プロバイダーの価格ページと照合して検証されました。
Stack Overflow、2025 年の開発者アンケート、2026 年 6 月 8 日取得
CloudZero、2025 年の AI コストの状況、2026-06-08 取得
CloudZero、LLM API の価格比較 2026 、2026-06-08 取得
Epoch AI、LLM 推論価格トレンド、2026-06-08 取得

[切り捨てられた]

## Original Extract

Most developers overpay for AI by 10-18x. We ran the numbers on 1M calls. Here's what we found.

We Tracked 1M LLM API Calls — 60% Were on the Wrong Model â”€ Tokonomics
Tokonomics
Features
Tools
Token Counter
Cost Calculator
Prompt Optimizer
API Request Builder
Model Matrix
ROI Calculator
Features
How it works
LLM Pricing
Compare
Tools
Token Counter
Cost Calculator
Prompt Optimizer
API Request Builder
Model Matrix
ROI Calculator
Blog
Log in
Start free
â†Â Blog
llm-cost-optimization
ai-api-pricing
model-routing
June 8, 2026
12 min read
We Tracked 1M LLM API Calls — 60% Were on the Wrong Model
In 2026, 82% of developers default to OpenAI GPT models ( Stack Overflow Developer Survey , 2025), but 60-70% of production API calls are simple tasks that don't need a frontier model.
Switching classification and extraction calls from GPT-4o to DeepSeek V3 saves 18x on input tokens ($2.50 → $0.14 per million).
Combining model routing with prompt caching cuts total LLM spend by 80-95%.
Average monthly AI spend hit $85,500 per company in 2025 — a 36% jump from the year before ( CloudZero , 2025).
Here's something that'll bother you if you're shipping AI features right now.
We looked at the first million API calls that came through Tokonomics — across 47 tenants, 9 providers, dozens of models. The pattern was the same almost everywhere: teams default to GPT-4o for everything. Customer support chatbots? GPT-4o. JSON extraction? GPT-4o. Classification into 5 categories? GPT-4o. It's the SELECT * of AI development — it works, nobody questions it, and it's silently bleeding money.
The waste isn't theoretical. It shows up in the billing dashboard every month, and most teams have no idea it's there.
Why Do 82% of Developers Default to GPT-4o?
In 2025, Stack Overflow's annual Developer Survey found that 82% of developers use OpenAI GPT models for their AI work ( Stack Overflow , 2025). That makes GPT-4o the de facto standard — the model people paste into their code and never change.
It makes sense. OpenAI has the best docs. Every tutorial uses GPT-4o. When you're prototyping at midnight, you're not running benchmarks across 6 providers. You grab what works.
But here's the problem: prototyping habits become production costs. That model you picked to test a feature in February is still running in production in June, processing 50,000 calls a day, and nobody's asked whether a $0.14/M model would give the same result as a $2.50/M model.
Our finding: When we built Tokonomics, our own internal chatbot ran on GPT-4o for three months before anyone checked. Switching the FAQ portion to GPT-4o-mini cut that component's cost by 94% with no measurable quality difference on our eval set.
This isn't unique to us. In 2026, Divyam.AI coined the term "LLMflation" to describe this exact pattern — teams sticking with legacy model choices long after cheaper alternatives catch up ( Divyam.AI , 2026). The inertia is real.
What Does Model Selection Actually Cost You?
Let's stop speaking in percentages and look at real money. Here's what 1 million requests cost across models, assuming an average of 500 input tokens and 200 output tokens per call — a typical production workload for classification, extraction, or short-form generation.
Monthly Cost for 1M Requests (500 in / 200 out tokens)
GPT-4o
Claude Sonnet 4
Claude Haiku 3.5
GPT-4o-mini
DeepSeek V3
GPT-4.1 Nano
$3,250
$4,500
$1,200
$195
$126
$130
Source: OpenAI, Anthropic, DeepSeek official pricing pages, June 2026
Source: OpenAI, Anthropic, DeepSeek official pricing pages, June 2026
That's a 25x cost difference between GPT-4o and GPT-4.1 Nano. For the same million requests.
Are you sure every one of those calls needs GPT-4o? Because in our data, 60-70% of them don't.
Which API Calls Don't Need a Frontier Model?
In 2026, multiple production operators reported that 60-70% of API calls in typical SaaS apps are simple enough for budget models — classification, short summarization, structured extraction, and routing decisions ( Prem AI , 2026). That matches what we've seen through Tokonomics.
Here's how it breaks down in practice:
Send to a budget model ($0.10-$0.80/M input):
Intent classification ("Is this a refund request or a general question?")
JSON/structured data extraction from text
Short summaries (under 200 words)
Content moderation / safety checks
Keep on a frontier model ($2.50-$3.00/M input):
Complex code generation and debugging
Long-form content creation where quality is critical
Tasks where you've benchmarked and confirmed the quality gap matters
The uncomfortable truth? Most teams never benchmark. They assume GPT-4o is required because they never tested whether GPT-4o-mini or DeepSeek gives an acceptable result. Field reports from LLM operators suggest that 40-60% of token budgets in production apps are pure waste — spent on frontier models doing budget-model work ( EditorialGE , 2026).
How Much Are Companies Actually Spending?
In 2025, CloudZero surveyed 500 US software engineers at companies with 250 to 10,000 employees. Average monthly AI spend jumped from $63,000 to $85,500 — a 36% increase year over year. And 45% of organizations now plan to spend over $100,000 per month on AI, more than double the 20% who said the same in 2024 ( CloudZero, State of AI Costs , 2025).
Here's the kicker: only 51% of those organizations can confidently evaluate their AI ROI. Half the companies spending six figures a month on AI can't tell you whether it's worth it.
Enterprise AI spending rose 320% from $11.5 billion to $37 billion between 2024 and 2025, even as per-token costs dropped dramatically ( BuildMVP Fast , 2026). That gap — prices falling while bills rise — is the clearest sign that usage is scaling faster than optimization.
Our finding: The teams spending the most on LLMs aren't the ones with the most sophisticated AI features. They're the ones who shipped AI features early, never revisited model selection, and let usage scale on autopilot. The $47,000 invoice that led us to build Tokonomics came from exactly this pattern.
Why Aren't Prices Falling Fast Enough to Fix This?
They are falling — at an extraordinary rate. In 2025, Epoch AI tracked LLM inference prices and found they dropped at a median rate of 50x per year. After January 2024, the decline accelerated to 200x per year for the cheapest available models at each capability level ( Epoch AI , 2025).
LLM Inference Prices Are Falling — But Bills Keep Rising
Price/token
Total spend
2022
2024
2025
2026
Source: Epoch AI + CloudZero, 2025-2026
Source: Epoch AI + CloudZero, 2025-2026
But usage growth is outpacing the price declines. Model API spending doubled from $3.5 billion to $8.4 billion between late 2024 and mid-2025, and 72% of companies plan to increase their LLM spending further ( CloudZero , 2026).
So no — you can't just wait for prices to drop and hope your bill fixes itself. By the time GPT-4o costs what DeepSeek costs today, you'll be running 10x more calls on an even newer frontier model.
The Fix: Route by Task, Cache by Pattern, Cap by Budget
There are three moves that compound together. Most articles cover one of them. Here's what happens when you stack all three.
Move 1: Route calls to the right model
Tag every API call by task type. Once you can see what each call does, you can route:
Classification and extraction → GPT-4o-mini or DeepSeek V3
Conversational support → Claude Haiku 3.5
Complex reasoning → GPT-4o or Claude Sonnet 4
If 60% of your calls shift from GPT-4o to a budget model, that's a 10-15x cost reduction on those calls alone. On a $3,250/month bill, that's roughly $1,950 saved.
For a step-by-step setup, see our guide to per-feature LLM cost tracking .
In 2026, Anthropic's prompt caching saves 90% on cached input tokens. OpenAI's automatic caching saves 50% with zero code changes — it just works if you're sending repeated system prompts ( Anthropic , OpenAI , 2026).
If your app sends the same system prompt on every call (most chatbots do), you're paying full price for tokens the provider has already processed. That's money on fire.
Our prompt caching guide for OpenAI and Anthropic covers the implementation details.
Move 3: Set hard spending caps
Even with routing and caching, bugs happen. A retry loop with no exit condition. A batch job that processes the same data twice. One night of runaway calls can undo a month of optimization.
Set a monthly budget cap that blocks API calls once you hit the limit. Not an alert you'll read at 9 AM — a hard block that stops the bleeding at 3 AM when nobody's watching.
Here's how hard spending caps work in practice .
Stacking all three moves doesn't add up linearly. It compounds:
Model routing alone: 50-70% savings
Add prompt caching: another 30-50% on top (applies to remaining spend)
Add budget caps: prevents the 100% overruns that wipe out your savings
A team spending $3,250/month on all-GPT-4o traffic can realistically land at $300-$650/month with the same output quality. That's not a rounding error. That's your next hire's monthly tooling budget.
You don't need to rewrite your architecture. Start here:
Audit your current model usage. Tag calls by feature and model. Find out which calls are using GPT-4o for work that GPT-4o-mini can handle. If you don't have visibility yet, Tokonomics shows this in under 5 minutes .
Pick your top 3 highest-volume call types. Run them through a cheaper model with a simple eval (does the output still pass your quality bar?). Most teams find 2 out of 3 work fine. Use our prompt cost optimizer to trim bloated prompts before switching models.
Set a budget alert at 80%. Before you optimize anything, make sure you'll know when costs spike. Here's how to set one up .
Switch one call type this week. Don't try to optimize everything at once. Move your highest-volume, lowest-complexity call to a budget model. Measure for a week. Then move the next one.
The goal isn't to use the cheapest model everywhere. It's to use the right model for each task — and stop paying frontier prices for work that doesn't need it.
How much cheaper is DeepSeek V3 compared to GPT-4o?
DeepSeek V3 costs $0.14 per million input tokens vs GPT-4o's $2.50 — that's 18x cheaper on input and 36x cheaper on output ($0.28 vs $10.00). For 1 million requests at 500 input + 200 output tokens, GPT-4o costs $3,250/month vs DeepSeek's $126/month. The quality gap matters for complex reasoning but is negligible for classification and extraction tasks.
What percentage of LLM API calls can use a cheaper model?
In 2026, multiple production operators report that 60-70% of API calls in typical SaaS applications are simple enough for budget models ( Prem AI , 2026). These include classification, structured data extraction, short summaries, and sentiment analysis. The remaining 30-40% still benefit from frontier models like GPT-4o or Claude Sonnet 4.
How much can prompt caching save on LLM costs?
Anthropic's prompt caching saves 90% on cached input tokens. OpenAI's automatic caching saves 50% with zero code changes ( Anthropic , OpenAI , 2026). Combined with model routing, caching can reduce total LLM spend by 80-95%.
What is the average monthly LLM API spend for companies in 2025?
Average monthly AI spend reached $85,500 per company — a 36% increase from the year before ( CloudZero , 2025). And 45% of organizations plan to spend over $100,000/month on AI, more than double the rate in 2024.
How do I track which LLM model each API call uses?
Use a metering proxy like Tokonomics that sits between your app and the LLM provider. Tag each call with metadata using a custom header. The proxy records model, token count, and cost per call — giving you a per-feature breakdown without code changes. Here's how to set it up in 5 minutes .
All sources retrieved June 2026. Pricing data verified against official provider pricing pages.
Stack Overflow, 2025 Developer Survey , retrieved 2026-06-08
CloudZero, State of AI Costs 2025 , retrieved 2026-06-08
CloudZero, LLM API Pricing Comparison 2026 , retrieved 2026-06-08
Epoch AI, LLM Inference Price Trends , retrieved 2026-06-08

[truncated]
