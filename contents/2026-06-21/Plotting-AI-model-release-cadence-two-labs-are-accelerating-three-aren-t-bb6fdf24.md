---
source: "https://swiftalerts.trade/the-cadence-trade-hn"
hn_url: "https://news.ycombinator.com/item?id=48615059"
title: "Plotting AI model release cadence: two labs are accelerating, three aren't"
article_title: "Plotting AI model release cadence: two labs are accelerating, three aren't"
author: "abipal15"
captured_at: "2026-06-21T04:35:57Z"
capture_tool: "hn-digest"
hn_id: 48615059
score: 1
comments: 0
posted_at: "2026-06-21T02:16:16Z"
tags:
  - hacker-news
  - translated
---

# Plotting AI model release cadence: two labs are accelerating, three aren't

- HN: [48615059](https://news.ycombinator.com/item?id=48615059)
- Source: [swiftalerts.trade](https://swiftalerts.trade/the-cadence-trade-hn)
- Score: 1
- Comments: 0
- Posted: 2026-06-21T02:16:16Z

## Translation

タイトル: AI モデルのリリース ペースのプロット: 2 つのラボが加速しているが、3 つは加速していない
説明: 明示的な方法論を使用した、ラボごとの累積フロンティア モデル リリースのグラフ。 Anthropic と OpenAI は加速する傾向を示しています。 Google、Meta、DeepSeek は同様です

記事本文:
分析 · 2026 年 6 月
AI モデルのリリース頻度のプロット: 2 つのラボが加速しているが、3 つは加速していない
方法論を使用したフロンティア モデルのリリース ペースのプロット
イーサン・モリック氏は今月、率直な観察を行った。AIの自己改善がたとえ弱いものであっても本物であるならば、それを導入している研究所は時間の経過とともにより早く出荷できるはずで、そうでない研究所は遅れをとるはずだ。同氏は、これはAnthropicとOpenAIではすでに見られたが、他の場所では見られないと主張した。実際にリリースデータがそれを裏付けているのか確認したかったのでプロットしてみました。
たとえ非常に限定的な方法であっても、AI の自己改善が可能であれば、AI 製品、ハーネス、モデルの両方を出荷するペースが上がるはずです。これは Anthropic と OpenAI で起こっているようですが、昨年追いつきつつあると思われるラボを含め、他のラボでは起こっていません。
イーサン・モリック、2026 年 6 月 19 日 [1]
この主張は反証可能ですが、AI の進歩の事例としては珍しいため、雰囲気ではなくデータに対してテストする価値があります。 2023 年第 1 四半期以降のラボごとのメジャー フロンティア モデル リリースの累計数は次のとおりです。
2023 年第 1 四半期から 2026 年第 2 四半期までの、ラボごとのメジャー フロンティア モデル リリースの累積数。傾きはリズムです。出典は注記[2]。
リリースとしてカウントされるもの: 個別に公開されているフロンティア モデル、フラッグシップ モデル、またはメジャー バージョン バンプ (GPT-4、GPT-4o、o1、o3、GPT-5、GPT-5.5、Claude 1 から Opus 4.8、Gemini 1 から 3.5、Llama 1 から 4、DeepSeek V2 から V4 および R1)。ポイント リリースとマイナー チェックポイントは、バージョン番号のインフレによる報酬を避けるために除外されます。分別のある人はこの線の引き方を変えるでしょうし、そうする場合は正確な数が少し変わります。
ケイデンスは証拠ではなく代理です。より急な傾きは、再帰的な自己改善だけでなく、より多くの資金、より良い管理、または単により頻繁に出荷するという決定とも一致します。このチャートは

モリック氏が説明したパターンがデータ内に存在します。因果関係を証明するものではありません。
累積数は誤解を招く可能性があります。累積線が上昇しているということは、リリースがまだ行われていることを意味するだけです。注目に値する信号は 2 次導関数であり、傾き自体が急になっているかどうかです。それが2番目のグラフです。
注目すべきはどのラインが曲がっているかです。 Anthropic と OpenAI は単に傾きが最も急であるだけではなく、右に向かって傾きが増加します。 Googleは2025年までほぼ横ばいだったが、2026年第2四半期に急成長した。Metaは2025年4月のLlama 4以降頭打ちとなり、それ以来フロンティアモデルを出荷していない。 DeepSeek は、加速することなく、四半期ごとに一定のペースで実行されます。
加速を分離するために、ここに年率換算のリリース率、つまり後続の 4 四半期のウィンドウを示します。このビューでは、平らな水平線は一定のリズムを意味します。上向きに曲がる線はケイデンスの加速を意味します。
年率換算のリリース率は、4 四半期のウィンドウを下回っています。フラットライン = 直線的なリズム。上向きに曲がる = ケイデンスが加速します。
2 つの研究室が曲がります。 3 つはそうではありません。 Anthropic は窓越しに年率をおよそ 3 倍に上昇させました。 OpenAI は 2 倍以上に増加しました。 Googleは2026年の追い上げまで横ばいを維持した。メタは衰退している。
これは支出と従業員数だけであり、ケイデンスギャップは拡大しないというデフレ的な見方があります。それが複合的であるという議論は、特定のループに基づいています。つまり、研究所は独自の製品を使用して後継製品を構築しています。 Anthropic エンジニアは、Claude コードを使用して、次の Claude のトレーニングと評価インフラストラクチャを作成します。 OpenAI は Codex on Codex を使用します。各リリースでは、次のリリースを生成するハーネスが改善されるため、次のリリースはより早く、より優れたものとして出荷されます。
これが何であり、何ではないかに注意してください。デプロイされたモデルはバージョン間でフリーズされるため、重み内でオンライン学習は発生しません。再帰は l で行われます

モデルではなく、組織のレベルです。これをオフライン RSI と呼びます。ループは、フォワード パス内ではなくリリース サイクル全体で閉じます。これは「自己改善する AI」よりもはるかに弱い主張ですが、実際にチャートはこれと一致しています。
他の 2 つは、再帰読み取りが予測する同じウィンドウに収まりました。まず、計算効率です。Tri Dao の FlashAttendant-4 は、2026 年 3 月に NVIDIA B200 で 71% の使用率に達しました [3]。同じグループの Mamba-3 は、トレーニング優先ではなく推論優先で明示的に設計されました [4]。サイクルあたりのトレーニングと推論が安くなるということは、四半期あたりのサイクルが増えることを意味します。第二に、人材の集中: 6 月 19 日の週に、Noam Shazeer (Transformer の共著者) が OpenAI に入社してアーキテクチャ研究を主導し、John Jumper (AlphaFold、2024 年ノーベル賞) が Google DeepMind を離れ Anthropic に移籍しました [5]。すでに最も早く出荷されているラボに人材が流れ込んでいます。
要点は反証可能な主張をテストすることであったため、正直な失敗モードは次のとおりです。
斜面の曲がりが止まります。次の 2 四半期で Anthropic と OpenAI の平坦化が見られる場合、その加速はループではなく 2026 年の成果であると考えられます。
遅れているラボは、真のオンライン学習を最初に提供します。継続的な加重内学習が Google、オープンソースの取り組み、またはコホート以外の誰かから提供される場合、オフライン RSI ケイデンスの利点は重要でなくなります。
リリース定義が機能します。ポイントリリースを数えると状況が変わります。私は意図的にそれらを除外しましたが、それは判断の判断であり、議論の対象となるのはデータ ファイルです。
メタのプラトーは戦略であり、失敗ではありません。メタは無差別級と異なるリリース哲学に移行しました。フラットなフロンティアリリースラインは、彼らが出荷したものを過小評価する可能性があります。
このグラフは AI が向上していることを証明するものではなく、再帰性を証明するものでもありません。それが示しているのは、より狭く、より防御可能であるということです。

まさに Mollick 氏の説明どおり、特定のラボではリリースのペースが加速していますが、3 つのラボではそうではありません。その差が偶然ではなくループであると考えれば、その差は広がり続けるはずです。資金や運のせいだと思ったら退行するはずです。いずれにせよ、今後 2 四半期のリリース日は完全なテストであり、予測は記録に残ります。
マーケット志向の人向け: ループが本物である場合、最もクリーンなリードスルーは、ほとんどがプライベートであるラボ自体ではなく、コンピューティング基板 (高速ラボは GPU と電力に集中的に資金を費やします) に対するものです。しかし、それは別の議論であり、推測の部分です。チャートは防御可能な部分です。
イーサン・モリック、X、2026 年 6 月 19 日: リズム観察。
リリース日は、AI リリース トラッカーや LLM 統計など、ラボの公開発表とリリース タイムラインから編集されています。基礎となる日付リストはリクエストに応じて入手可能です。
Tri Dao、FlashAttendant-4: アルゴリズムとカーネル パイプラインの共同設計、2026 年 3 月。1605 TFLOPs/s、B200 での使用率 71%。
Tri Dao、Mamba-3、パート 1、2026 年 3 月、トレーニング優先設計から推論優先設計への移行について。
Shazeer と Jumper の動きに関するレポート、2026 年 6 月 (研究室の発表。報道は 6 月 19 日の週に広く集約されました)。

## Original Extract

A chart of cumulative frontier model releases per lab, with explicit methodology. Anthropic and OpenAI show accelerating cadence; Google, Meta, and DeepSeek don

Analysis · June 2026
Plotting AI model release cadence: two labs are accelerating, three aren't
Plotting frontier model release cadence, with methodology
Ethan Mollick made an offhand observation this month: if AI self-improvement is real, even weakly, then the labs that have it should ship faster over time, and the ones that don't should fall behind. He claimed this was already visible at Anthropic and OpenAI but nowhere else. I wanted to check whether the release data actually supports that, so I plotted it.
If AI self-improvement, even in a very limited way, is possible, the cadence of shipping both AI products, harnesses, and models should go up. This appears to be happening at Anthropic and OpenAI, but not for any other labs, including those that seemed to be catching up last year.
Ethan Mollick, June 19, 2026 [1]
The claim is falsifiable, which is rare for AI-progress takes, so it's worth testing against data rather than vibes. Here's the cumulative count of major frontier model releases per lab since Q1 2023.
Cumulative count of major frontier model releases per lab, Q1 2023 to Q2 2026. Slope is cadence. Sources in notes [2].
What counts as a release: a distinct, publicly available frontier or flagship model or major version bump (GPT-4, GPT-4o, o1, o3, GPT-5, GPT-5.5; Claude 1 through Opus 4.8; Gemini 1 through 3.5; Llama 1 through 4; DeepSeek V2 through V4 and R1). Point releases and minor checkpoints are excluded to avoid rewarding version-number inflation. Reasonable people will draw this line differently, and the exact counts shift a little if you do.
Cadence is a proxy, not proof. A steeper slope is consistent with recursive self-improvement, but also with more funding, better management, or simply a decision to ship more often. This chart shows the pattern Mollick described exists in the data. It does not prove the causal mechanism.
Cumulative counts can mislead. A rising cumulative line only means releases are still happening. The signal worth caring about is the second derivative , whether the slope itself is steepening. That's the second chart.
The thing to look at is which lines are bending. Anthropic and OpenAI don't just have the steepest slopes, their slopes increase toward the right. Google sat nearly flat through 2025, then sprinted in Q2 2026. Meta plateaued after Llama 4 in April 2025 and hasn't shipped a frontier model since. DeepSeek runs a steady quarterly cadence without accelerating.
To isolate acceleration, here's the annualized release rate, a trailing four-quarter window. On this view a flat horizontal line means constant cadence; an upward-bending line means accelerating cadence.
Annualized release rate, trailing four-quarter window. Flat line = linear cadence. Upward bend = accelerating cadence.
Two labs bend up. Three don't. Anthropic roughly tripled its annualized rate over the window; OpenAI more than doubled. Google held flat until a 2026 catch-up; Meta is in decline.
There's a deflationary reading where this is just spending and headcount, and the cadence gap won't compound. The argument that it does compound rests on a specific loop: the labs use their own products to build their successors. Anthropic engineers use Claude Code to write training and eval infrastructure for the next Claude. OpenAI uses Codex on Codex. Each release improves the harness that produces the next release, so the next one ships sooner and better.
Note what this is and isn't. The deployed model is frozen between versions, so there's no online learning happening inside the weights. The recursion is at the level of the organization, not the model. Call it offline RSI: the loop closes across release cycles rather than within a forward pass. That's a much weaker claim than "self-improving AI," and it's the one the chart is actually consistent with.
Two other things landed in the same window that the recursion reading predicts. First, compute efficiency: Tri Dao's FlashAttention-4 hit 71% utilization on NVIDIA B200 in March 2026 [3], and Mamba-3, from the same group, was explicitly designed inference-first rather than training-first [4]. Cheaper training and inference per cycle means more cycles per quarter. Second, talent concentration: in the week of June 19, Noam Shazeer (Transformer co-author) joined OpenAI to lead architecture research, and John Jumper (AlphaFold, 2024 Nobel) left Google DeepMind for Anthropic [5]. Talent is flowing toward the labs already shipping fastest.
The honest failure modes, since the whole point was to test a falsifiable claim:
The slope stops bending. If the next two quarters show Anthropic and OpenAI flattening, the acceleration was a 2026 artifact, not a loop.
A lagging lab ships true online learning first. If continual in-weights learning arrives from Google, an open-source effort, or anyone outside the cohort, the offline-RSI cadence advantage stops mattering.
The release definition is doing the work. If you count point releases, the picture changes. I excluded them deliberately, but it's a judgment call, and the data file is the thing to argue with.
Meta's plateau is strategy, not failure. Meta shifted toward open-weight and a different release philosophy; a flat frontier-release line may understate what they shipped.
The chart doesn't prove AI is improving, and it doesn't prove recursion. What it shows is narrower and more defensible: two specific labs have a release cadence that is accelerating, and three don't, exactly as Mollick described. If you think that gap is a loop rather than a coincidence, it should keep widening. If you think it's funding or luck, it should regress. Either way, the next two quarters of release dates are a clean test, and the prediction is on the record.
For the markets-minded: the cleanest read-through, if the loop is real, is to the compute substrate (the fast labs spend concentrated dollars on GPUs and power) rather than to the labs themselves, which are mostly private. But that's a separate argument, and it's the speculative part. The chart is the defensible part.
Ethan Mollick, on X, June 19, 2026: the cadence observation .
Release dates compiled from public lab announcements and release timelines, including AI Release Tracker and LLM Stats . The underlying date list is available on request.
Tri Dao, FlashAttention-4: Algorithm and Kernel Pipelining Co-Design , March 2026. 1605 TFLOPs/s, 71% utilization on B200.
Tri Dao, Mamba-3, Part 1 , March 2026, on the shift from training-first to inference-first design.
Reporting on the Shazeer and Jumper moves, June 2026 (lab announcements; coverage aggregated widely the week of June 19).
