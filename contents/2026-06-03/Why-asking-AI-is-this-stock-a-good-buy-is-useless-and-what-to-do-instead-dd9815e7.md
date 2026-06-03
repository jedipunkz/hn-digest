---
source: "https://pg1.github.io/paper-profit/experiments/rating-stocks-llm/"
hn_url: "https://news.ycombinator.com/item?id=48370305"
title: "Why asking AI \"is this stock a good buy?\" is useless – and what to do instead"
article_title: "How to decides whether a stock is worth buying using AI | paper-profit"
author: "pg1"
captured_at: "2026-06-03T00:45:14Z"
capture_tool: "hn-digest"
hn_id: 48370305
score: 3
comments: 0
posted_at: "2026-06-02T13:54:20Z"
tags:
  - hacker-news
  - translated
---

# Why asking AI "is this stock a good buy?" is useless – and what to do instead

- HN: [48370305](https://news.ycombinator.com/item?id=48370305)
- Source: [pg1.github.io](https://pg1.github.io/paper-profit/experiments/rating-stocks-llm/)
- Score: 3
- Comments: 0
- Posted: 2026-06-02T13:54:20Z

## Translation

タイトル: なぜ AI に「この株は買って良いのか?」と尋ねるのか?役に立たない – 代わりに何をすべきか
記事タイトル: AI を使用して株を買う価値があるかどうかを判断する方法 |紙の利益
説明: PaperProfit は、実際のお金を危険にさらすことなく、実際の市場データを使用して、投資、取引、リスク管理、市場、ポートフォリオ管理、および幅広い投資戦略について実践しながら学ぶことができる、包括的なペーパー トレーディング プラットフォームです。

記事本文:
AIを使用して株を買う価値があるかどうかを判断する方法
最初に PaperProfit の構築を開始したとき、解決しなければならなかった問題の 1 つは、株式を評価する方法を理解することでした。
私の最初の試みは恥ずかしいほど世間知らずでした。 AI に「Apple は今、良い株ですか?」と尋ねてみました。
返ってきた答えは自信に満ち、流暢で、ほとんど役に立たなかった。このモデルには、Apple の現在の価格がいくらなのか、市場が上がっているのか下がっているのか、あるいは Apple が自社の歴史に比べて安いのか高いのか、まったく分かりませんでした。それは、6か月間昏睡状態にあった人に株式のヒントを教えてもらうようなものでした。その言葉は正しく聞こえました。実体はそこにはなかった。
そこで私は別のアプローチを試みました。ウォール街のアナリストによる評価、つまり大手銀行による「買い/ホールド/売り」の推奨事項のデータ フィードを見つけて、それらを平均化しました。確かに専門家は自分たちが何をしているのか知っていたでしょうか？
それも駄目。アナリストの格付けは時代遅れで、矛盾しており（銀行は対象企業から利益を得ている）、絶望的に楽観的であることが多いことが判明した。すべてのアナリストがすべてを「買い」と評価すると、シグナルは消えます。
行き詰まってしまいました。そして私は疑問に思い始めました。実際に株式をどのように評価する人がいるのでしょう?
ウォール街の考え方を学ぶ短期集中講座
私はこれを理解するのに何週間も費やしました。本を読んだり、金融関連の YouTube を見たり、AI アシスタントと長時間会話したり、投資に関する学術論文を調べたりしました。
これが私が学んだことです。プロの株式アナリスト（ヘッジファンドや銀行の人々で、文字通りの仕事はどの銘柄を所有すべきかを判断することです）は、単一のシグナルに依存しません。複数のレンズを組み合わせて、それぞれのレンズが企業に関する異なる質問に答えます。
目標は「市場を予測する」ことではないと私は気づきました。教育の枠組みを構築することであるべきです。

株式アナリストが実際にどう考えているか。
これをソフトウェア プロジェクトで実際に実装できるものに単純化した後、3 つの核となる柱にたどり着きました。
第 1 の柱: 数字 (ファンダメンタルズ分析)
これが財務上の健康診断です。会社の成績表を見るようなものだと考えてください。
会社は儲かってますか？成長していますか？不況に陥りかねないほどの多額の負債を抱えているのだろうか？企業の実際の収益と比較して、株価は安いのでしょうか、それとも割高なのでしょうか?
私たちが注目する数字のいくつかは次のとおりです。
収益と成長 — 同社の今年の売上は昨年よりも多いですか?
利益率 — 彼らが稼いだ1ドルのうち、実際に手元に残るのはいくらでしょうか?
借金 — 彼らは厳しい状況を乗り切ることができるでしょうか、それともあと 1 四半期で問題が起きるのでしょうか?
評価 — PER は、簡単に言うと、何年分の収益に対して支払っているのかを表します。 PER 20 は、会社の収益 1 ドルごとに 20 ドルを支払うことを意味します。高いことは必ずしも悪いことではありません（市場が成長を期待していることを意味する可能性があります）が、プレミアムを支払っていることを意味します。
第 2 の柱: チャート (テクニカル分析)
これは価格の動きに関するものであり、企業の価値ではなく、市場がその価値をどのように考えているか、そして勢いがどこを指しているのかについてです。
それはちょっとサーフィンに似ています。ファンダメンタル分析は、その波がキャッチする価値があるかどうかを示します。テクニカル分析は、波が上昇しているのか、それともブレイクしているのかを示します。
移動平均 — 過去 50 日間または 200 日間、株価は上昇傾向にありますか、それとも下落傾向にありますか?短期平均が長期平均を上回った場合、それは多くの場合強気のサインです。
RSI (相対強度指数) — 0 ～ 100 のスコア。 70を超えると、株が買われすぎている（熱意が高すぎる）ことを意味することがよくあります。 30未満は売られすぎ（過剰な恐怖）を意味することがよくあります。 60 前後は、多くの場合、健全で上昇傾向にある銘柄です。
MACD

— 購入圧力が増加しているか減少しているかを示すモメンタム指標。
これらはどれも未来を予測するものではありません。しかし、それらは群衆が現在何をしているかを説明しており、市場では群衆が重要です。
柱 3: ストーリー (定性分析)
これは自動化するのが最も難しい柱であり、最も興味深いものです。
数字は何が起こったのかを伝えるだけです。彼らはその理由や次に何が起こるのかを教えてくれません。
定性層は、次のような質問に答えようとします。
この会社を経営している人たちは、自分たちが何をしているのかを実際に知っているのでしょうか?
経営陣は前回の決算発表で自信に満ちたガイダンスを与えたのか、それとも曖昧で回避的なものであったのか?
この会社は成長市場にいますか、それとも縮小市場にいますか?
インサイダーは自社株を買っているのか、それともひそかに売っているのか？
ここで AI 推論が真に役立ちます。言語モデルは、決算会見の記録を読み、数字に見落としているもの、つまり経営陣の口調、ガイダンスを引き上げたのか引き下げたのか、悪い四半期についての説明が本当の計画のように聞こえたのかカバーストーリーのように聞こえたのかを拾い上げることができます。また、最新の収益レポートと以前の収益レポートを簡単に比較できます。
この時点で、「わかった、でもアプリは実際にこれらすべてをどのように行うのですか?」と疑問に思うかもしれません。
正直な答えは、本当に必要な部分には AI を使用し、それ以外の部分にはプレーン コードを使用しているということです。
面倒な作業のほとんどは決定的です。株価の取得、比率の計算、移動平均の計算、危険信号の通知など、これらすべてが通常の Python コードとして実行されます。速く、安く、完璧に再現可能です。 AIは必要ありません。
LLM を導入するのは定性層、つまり非構造化テキストの読み取り、解釈、要約を必要とする層です。具体的には、アプリは少数の対象を絞った API 呼び出しを実行して、2 種類の

書類:
決算コールの記録 — AI は最新のコールを読み取り、経営陣の信頼、収益見通し、マージンの軌道、競争力のあるポジショニング、株式に対する最も強い強気派と弱気派の議論を抽出します。
SEC への提出書類 — AI は今年の 10-K を前年のものと比較し、収益の質、収益の質、バランスシートの健全性の変化、および報告する価値のある新たなリスクや危険信号を探します。
コードからの定量的データ、AI からの定性的シグナルなど、すべての準備が整うと、スコアリング システムがそれらを組み合わせて 1 つの推奨事項を作成します。
3 つの柱すべてにわたるデータを取得したら、それを 1 つの推奨事項に変える方法が必要です。その仕組みは次のとおりです。
各銘柄を次の 5 つの次元でスコア付けします。
各次元には −2 から +2 までのスコアが与えられます。ゼロは中立です。正は信号が良好であることを意味します。マイナスは危険信号であることを意味します。
次に、これらのスコアを重み付けと組み合わせます。長期的な投資決定において、すべての要素が同等に重要であるわけではないためです。
総合スコア =
30% × 品質
+ 25% × 成長
+ 20% × 評価額
+ 15% × 勢い
+ 10% × リスク
強力なファンダメンタルズを備えた企業は長期的に優れたパフォーマンスを発揮する傾向があることが調査で一貫して示されているため、品質と成長が最も重視されます。たとえ素晴らしいビジネスであっても、間違ったタイミングで購入すると悪い投資になる可能性があるため、勢いが含まれています。リスクはペナルティとして機能します。多額の負債や高いボラティリティは、本来は良いスコアを引き下げる可能性があります。
これが実際にどのように見えるかは次のとおりです。 Meridian Corp という架空のソフトウェア会社を評価していると想像してください。
加重合計: +0.95 → シグナル: 🔵 弱い買い
システムによれば、この企業は基本的に健全で割安に見えるが、市場はまだそれに報いておらず、成長は鈍化している。それ

』は、焦らずに見る価値があります。
AI の仕事は分析を置き換えることではありません。人間にとっては面倒でも、言語モデルにとっては簡単な分析の部分を実行することです。
60 ページの 10-K ファイリングを読み、最も重要な 10 の事実を抽出します。わかりやすい英語で決算報告を要約します。経営陣の言葉が自信満々から守りに転じたときに警告する。この四半期の結果を過去 4 四半期と比較し、傾向に注目します。
最終的な判断を下すのはやはり人間です。 AI は、推測ではなく、全体像からスタートしていることを確認します。
スコアがどれほど優れているように見えても、システムは次のような危険信号を自動的に表示します。
複数四半期連続で収益または利益率が減少
金利上昇環境における多額の負債
不正会計（利益の修正修正、監査人の変更）
理由を明らかにしない大量のインサイダー販売
収益の 50% 以上が 1 人の顧客から
株価はあらゆる指標で魅力的に見えても、いずれかが当てはまれば危険である可能性があります。
ここで私が学んだことは、どんな金融本も前もって教えてくれるわけではないということです。優秀なアナリストであっても、多くの間違いを犯しているということです。
市場は部分的には数学であり、部分的には人間の心理学です。企業は完璧なファンダメンタルズを持っていても、センチメントの変化により 40% 下落する可能性があります。ミームと Reddit スレッドのせいで、ひどいビジネスが 3 倍になる可能性があります。
このような体系的なアプローチがもたらすものは水晶玉ではありません。それはプロセス、つまり感情や誇大広告ではなく証拠に基づいて、一貫した合理的な意思決定を行う方法を提供します。多くの決定では、推測よりもうまくいく傾向があります。
それがPaperProfitの目標です。未来を予測するためではなく、専門家が行うように、明確に、一貫して、雑音に流されずに株式について考えるのに役立ちます。
*W

この背後にあるコードを詳しく調べてみませんか?技術ドキュメントを参照してください *

## Original Extract

PaperProfit is a comprehensive paper-trading platform that lets you learn by doing—teaching you about investing, trading, risk management, markets, portfolio management, and a wide range of investment strategies using real market data, all without risking real money.

How to decides whether a stock is worth buying using AI
When I first started building PaperProfit , one of the problems I had to solve was figuring out how to evaluate a stock.
My first attempt was embarrassingly naive. I just asked an AI: “Is Apple a good stock right now?”
The answer came back confident, fluent, and almost completely useless. The model had no idea what Apple’s current price was, whether the market was up or down, or whether Apple was cheap or expensive relative to its own history. It was like asking someone who’d been in a coma for six months to give you stock tips. The words sounded right. The substance wasn’t there.
So I tried another approach: I found a data feed of Wall Street analyst ratings — the “buy/hold/sell” recommendations from big banks — and averaged them together. Surely the professionals knew what they were doing?
Also useless. Analyst ratings, it turns out, are often outdated, conflicted (banks make money from the companies they cover), and hopelessly optimistic. When every analyst rates everything a “buy,” the signal disappears.
I was stuck. And I started to wonder: how does anyone actually evaluate a stock?
A crash course in how Wall Street thinks
I spent weeks trying to figure this out. I read books, watched finance YouTube, had long conversations with AI assistants, and dug through academic papers on investing.
Here’s what I learned: professional equity analysts — the people at hedge funds and banks whose literal job is figuring out which stocks to own — don’t rely on any single signal. They combine multiple lenses, each answering a different question about a company.
The goal, I realized, shouldn’t be to “predict the market.” It should be to build an educational framework that mirrors how equity analysts actually think.
After simplifying this into something a software project could actually implement, I landed on three core pillars:
Pillar 1: The Numbers (Fundamental Analysis)
This is the financial health check. Think of it like looking at a company’s report card.
Is the company making money? Is it growing? Does it have a mountain of debt that could sink it in a recession? Is the stock cheap or overpriced relative to what the business actually earns?
A few of the numbers we look at:
Revenue and growth — Is the company selling more this year than last year?
Profit margins — Of every dollar they earn, how much do they actually keep?
Debt — Could they survive a rough patch, or are they one bad quarter away from trouble?
Valuation — The P/E ratio, in plain English, is how many years of earnings you’re paying for. A P/E of 20 means you’re paying $20 for every $1 the company earns. High isn’t always bad (it might mean the market expects growth), but it does mean you’re paying a premium.
Pillar 2: The Chart (Technical Analysis)
This one is about price movement — not what the company is worth, but what the market thinks it’s worth, and where momentum is pointing.
It’s a bit like surfing. The fundamental analysis tells you whether a wave is worth catching. The technical analysis tells you whether the wave is building or breaking.
Moving averages — Is the stock trending up or down over the past 50 or 200 days? When the short-term average crosses above the long-term average, that’s often a bullish sign.
RSI (Relative Strength Index) — A 0–100 score. Above 70 often means a stock is overbought (too much enthusiasm). Below 30 often means it’s oversold (too much fear). Around 60 is often a healthy, upward-trending stock.
MACD — A momentum indicator that shows whether buying pressure is increasing or decreasing.
None of these predict the future. But they describe what the crowd is currently doing — and crowds matter in markets.
Pillar 3: The Story (Qualitative Analysis)
This is the hardest pillar to automate, and the most interesting.
Numbers only tell you what happened. They don’t tell you why , or what’s coming next.
The qualitative layer tries to answer questions like:
Do the people running this company actually know what they’re doing?
Did management give confident guidance on the last earnings call, or were they vague and evasive?
Is this company in a growing market, or a shrinking one?
Are insiders buying their own stock — or quietly selling?
This is where AI reasoning becomes genuinely useful. A language model can read through an earnings call transcript and pick up on things the numbers miss — the tone of management, whether they raised or lowered guidance, whether their explanation for a bad quarter sounded like a real plan or a cover story. Also it can easily compare latest earning reports with the previous one.
At this point you might be wondering: okay, but how does the app actually do all of this?
The honest answer is that we use AI for the parts that genuinely need it — and plain code for everything else.
Most of the heavy lifting is deterministic. Fetching stock prices, calculating ratios, computing moving averages, flagging red flags — all of that runs as regular Python code. It’s fast, cheap, and perfectly repeatable. No AI required.
Where we bring in an LLM is for the qualitative layer: the things that require reading, interpreting, and summarizing unstructured text. Specifically, the app makes a small number of targeted API calls to analyze two types of documents:
Earnings call transcripts — the AI reads the most recent call and extracts management confidence, revenue guidance, margin trajectory, competitive positioning, and the strongest bull and bear arguments for the stock.
SEC filings — the AI compares the current year’s 10-K against the previous year’s and looks for changes in earnings quality, revenue quality, balance sheet health, and any new risks or red flags worth flagging.
Once all of that is ready — the quantitative data from code, the qualitative signals from AI — the scoring system combines it into a single recommendation.
Once we have data across all three pillars, we need a way to turn it into a single recommendation. Here’s how that works.
We score each stock across five dimensions:
Each dimension gets a score from −2 to +2 . Zero is neutral. Positive means the signal is favorable. Negative means it’s a warning sign.
We then combine those scores with weights — because not all dimensions matter equally for long-term investment decisions:
Overall Score =
30% × Quality
+ 25% × Growth
+ 20% × Valuation
+ 15% × Momentum
+ 10% × Risk
Quality and growth get the most weight because research consistently shows that businesses with strong fundamentals tend to outperform over the long run. Momentum is included because even a great business can be a bad investment if you buy it at the wrong time. Risk acts as a penalty — high debt or high volatility can drag down an otherwise good score.
Here’s what this looks like in practice. Imagine we’re evaluating a fictional software company called Meridian Corp :
Weighted total: +0.95 → Signal: 🔵 Weak Buy
The system says this company looks fundamentally sound and cheap, but the market hasn’t rewarded it yet — and growth is slowing. It’s worth watching, not rushing into.
The AI’s job is not to replace analysis — it’s to do the parts of analysis that are tedious for humans but easy for language models.
Reading a 60-page 10-K filing and extracting the ten most important facts. Summarizing an earnings call in plain English. Flagging when management’s language shifts from confident to defensive. Comparing this quarter’s results to the last four quarters and noting the trend.
The human still makes the final call. The AI just makes sure you’re starting from a complete picture — not guessing.
No matter how good a score looks, the system automatically surfaces these red flags:
Revenue or margins declining for multiple quarters in a row
High debt in a rising interest rate environment
Accounting irregularities (earnings restated, auditors changed)
Heavy insider selling without a disclosed reason
More than 50% of revenue from a single customer
A stock can look attractive on every metric and still be dangerous if one of these is true.
Here’s something I learned that no finance book will tell you upfront: even the best analysts are wrong a lot .
Markets are partly math and partly human psychology. A company can have perfect fundamentals and still fall 40% because sentiment shifts. A terrible business can triple because of a meme and a Reddit thread.
What a systematic approach like this gives you isn’t a crystal ball. It gives you a process — a way to make consistent, well-reasoned decisions based on evidence rather than emotion or hype. Over many decisions, that tends to work better than guessing.
That’s the goal of PaperProfit. Not to predict the future, but to help you think about stocks the way a professional would — clearly, consistently, and without getting swept up in the noise.
*Want to dig into the code behind this? See the technical documentation *
