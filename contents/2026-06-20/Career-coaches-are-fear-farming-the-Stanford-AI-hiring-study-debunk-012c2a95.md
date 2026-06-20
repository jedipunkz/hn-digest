---
source: "https://placementist.com/insights/fear-farming-the-stanford-ai-hiring-study-debunk"
hn_url: "https://news.ycombinator.com/item?id=48607303"
title: "\"Career coaches\" are fear-farming the Stanford AI hiring study [debunk]"
article_title: "Fear-farming the Stanford AI hiring study [debunk] • placementist"
author: "nikkotyze"
captured_at: "2026-06-20T10:05:33Z"
capture_tool: "hn-digest"
hn_id: 48607303
score: 3
comments: 0
posted_at: "2026-06-20T08:06:17Z"
tags:
  - hacker-news
  - translated
---

# "Career coaches" are fear-farming the Stanford AI hiring study [debunk]

- HN: [48607303](https://news.ycombinator.com/item?id=48607303)
- Source: [placementist.com](https://placementist.com/insights/fear-farming-the-stanford-ai-hiring-study-debunk)
- Score: 3
- Comments: 0
- Posted: 2026-06-20T08:06:17Z

## Translation

タイトル: 「キャリアコーチ」がスタンフォード大学の AI 採用調査を恐怖に煽っている [誤りであることが判明]
記事のタイトル: スタンフォード AI 採用調査の恐怖を煽る [誤りを暴く] • プレースメント担当者
説明: AI がどこでもあなたを拒否するのではないかと心配していますか?誰か

記事本文:
スタンフォード大学の AI 採用調査を恐怖に煽る [誤りを暴く]
スタンフォード大学の新しい採用調査については 2 つのことが真実です。誰もが引用している調査結果は真実であり、それに関する解説の多くは抽象的なものにとどまっているようです。
このギャップは、1 つの不適切に構築されたツールに関する僅差の結果が、業界全体に対する評決となり、キャリアインフルエンサーが求職者のパニックを引き起こす新たな方法となった経緯を表しています。
バイラルなスタンフォード論文 (雇用におけるアルゴリズム的モノカルチャー) は本格的な研究であり、読む価値がありますが、「AI 採用」全体ではなく、パイメトリクスと呼ばれる 1 つの狭いゲームベースのツールについて研究したものです。パニックのほとんどは、1 つのベンダーを業界全体に一般化する人々によって引き起こされます。
保存する価値のある発見の 1 つは、全社的な公平性監査は、個々の職務が差別しているように見えても問題がないということです。一緒にプールされたパイメトリクスは、通常の悪影響のスクリーニングを通過しましたが、1,746 のポジションに分割され、最大 11% が黒人の応募者に不利に働きました。採用を企業レベルでのみ監査している場合は、変更する必要があります。
「どこでも拒否される」という恐ろしい話は、データにはほとんど存在しません。応募者の 84% は 1 つのポジションのみに応募しました。 10 個に適用されるのはわずか 0.02% です。著者らがこの悪夢を直接シミュレーションしたところ、すべてのモデルから拒否された人は一人もいませんでした。
ツール自体が本物です。 pymetrics は、会社の現在の従業員を「良い」従業員としてトレーニングし、ランダムなプロファイルを「悪い」従業員としてトレーニングするため、誰が仕事ができるかではなく、誰が既存のスタッフに似ているかを学習します。そして、この研究ではパフォーマンスを予測するという証拠はまったくありません。それは偏見があり、検証されていません。
著者らは、制限事項のセクションでこれについてほぼすべてを認めています。調査は注意深く行われ、ヘッジが行われた。注意しなかったのは、「AI がどこでもあなたを拒否している」という事態に至る連鎖で、すべてのステップで著者が意図的に付けた修飾子が落とされてしまいました。
調査中の論文

上: 雇用におけるアルゴリズムによるモノカルチャー (FAccT 2026)。
この研究論文は、「雇用におけるアルゴリズム的モノカルチャー」と呼ばれています。その中心的な考え方は「モノカルチャー」の概念です。すべての雇用主が同じ少数の AI ベンダーを通じて候補者を選考すると、1 つの偏ったモデルが 1 つの企業で害を及ぼすだけでなく、どこでも一度に締め出されてしまいます。当然の心配であり、結局のところ、当然の疑念でもあります。
それを研究するために、研究者らは単一ベンダーから4年間の実際の採用データを入手し、同じ人、同じ人種グループがフィルタリングされ続けているかどうかを確認した。
ここまでは順調ですね。次に、どのベンダーかを調べます。
すでに雇用されている人々に似ている、または雇っていない人々。それが科学全体です。
pymetrics という名前のベンダーを調べます。履歴書や応募者追跡システムではありません。
pymetrics はゲームをプレイするためのツールです。はい、ゲームです。リスクを取ること、処理速度、計画性などを測定すると思われる 12 ～ 16 のゲームをプレイします。これらのゲームのうち 12 件は、倉庫に応募する場合でも財務デスクに応募する場合でも同じです。
モデルはあなたのゲームプレイを採点し、「推奨」または「推奨しません」の 2 つの単語のいずれかを吐き出します (なぜ企業がその費用を支払うのでしょうか?)。それでおしまい。ああ、約 42% の確率で「推奨しない」と表示されます。
pymetricsプロセスの著者が論文で説明
各企業について、モデルはその会社の少なくとも 50 人の現従業員を「良い」例としてトレーニングし、ランダムな人々を「悪い」例としてトレーニングします。いいえ、「良い」グループはトップパフォーマーではありません。
それは現在その職に就いている人です。少なくともこの調査では、これらの現職企業が実際に優れた業績を上げていたことを示すものは何もない。
そして、「悪い」グループとは、役割を果たせなかった人、解雇された人、成績が落ちた人ではありません。それはランダムな見知らぬ人です。
そのため、モデルは、強い採用と弱い採用を分けるものなど、あなたが実際に望んでいることを学習することはありません。

あなたの組織のために。それはもっと愚かなことを学習します。あなたはすでに給与計算に載っている人々のように見えますか、それとも通りから来たランダムな人のように見えますか？
ボーナスとして、スコアは 330 日間キャッシュされます。年内に別の場所に応募すると、同じ保存されたゲームプレイで 2 回審査されることがよくあります。わーい！
しかしまあ、少なくともデータセットは本当に巨大です: 420 万件の申請、340 万人、そして 156 の雇用主。
pymetrics は以前、すべての申請者をプールすることで独自の公平性をチェックしていました。このようにプールされた結果は、通常の悪影響の審査を通過しました。黒人の応募者は52.5％、白人は58.3％で合格した。研究者らは、これが選択ツールを検証する間違った方法であることを正しく指摘しました。
米国雇用法 (タイトル VII) は、全社的な評価ではなく、各職務を個別に評価します。そこで彼らはデータを 1,746 個の位置に分割し、再度調べました。平均が隠していたバイアスが現れ、求人の約 11% が黒人応募者に不利に働き (統計補正後は 10.6%)、黒人応募者の約 4 分の 1 がそれ​​らの求人に採用されました。
これは貢献であり、良い貢献です。平均は差別を隠す可能性があります。企業レベルでのみ採用を監査する場合、これは間違いなく考慮すべき事項です。良い持ち帰り、保存する価値があります。
これは、これらのツールを購入する人にとって、報告する価値のあるデータガバナンスの問題を提起します。
標準的な企業コンプライアンスでは、雇用機会均等委員会 (EEOC) の公式データは厳密にファイアウォールで保護されています。これは候補者にとって完全にオプションであり、法的に自動評価エンジンや意思決定ツールにフィードすることができない分離されたリポジトリに置かれます。
では、156 社の雇用主のスコアにマッピングされた、評価ベンダーのデータセット内で人種データは何をしているのでしょうか?
研究論文の著者によって提供された
答えはムンです

デーン、それがポイントです。これは、pymetrics によって収集された自発的なセルフ ID です。申請者のわずか 40% のみがそれを提供しており、各雇用主はベンダーにそれを収集させるかどうかを選択します。すべて合法です。すべて標準です。そして、保護された EEO-1 申請とはまったく異なる流れです。
要約すると、候補者は、おそらく調達チームが検討したことのない設定によって管理され、あなたが管理していない画面上で、人口統計データを第三者に渡しているのです。
販売者は何か違法なことをしているのでしょうか？いいえ、会社レベルのコンプライアンス監査でこれが確認されることはありますか?また、いいえ。まあ - そうだね？
さて、紙の問題です。
問題 1: 彼らは 1 つのツールが悪いと証明したのに、すべてがそうであるかのように話しました。
この論文全体は「モノカルチャー」という言葉に基づいて構築されています。誰もが同じ支配的なシステムに集中するため、1 つの偏ったモデルが市場全体を汚染するという考えです。
それが現実的で恐ろしいように聞こえるように、冒頭では HireVue を引用しています。HireVue は、Fortune 100 の 60% 以上と、10 の最大の連邦政府機関のうち 8 つによって使用されています。恐ろしい数字ですよね？
ただし、彼らは HireVue を勉強していませんでした。彼らはパイメトリクスを研究しました。別の、はるかに小さな会社です。競争相手です。 HireVue は、構造化された面接と求人プレビューという異なる製品カテゴリであり、適切に設計されていれば科学的に有効な方法です。ただし、この文書では HireVue のバージョンをテストしていないため、それらを検証するつもりはありません。 pymetricsはゲームをします。
そして、フォーチュン 100 企業が実際に HireVue をどのように使用しているのかを説明する人は誰もいません。これは、構造化されたインタビュー、作品サンプル、リファレンスチェックなどを含むプロセスの一段階である可能性があります。
研究者自身の結果セクションでは、恐ろしいバージョンのモノカルチャー（同じモデルが複数の企業であなたを審査する）が「限られた」影響を伴う「まれな例」で起こることを認めている、「POに応募する応募者が非常に少ないため」

異なる雇用主の拠点でも、同じ基礎となる pymetrics モデルが提供されます。」
そこで彼らは代わりに、より弱い、より緩やかな定義を研究することに切り替えます。見出しのコンセプトはデータにはほとんど現れません。彼らは、HireVue のスケールをどこでも鳴らせるように主導しましたが、彼ら自身の数字によると、そうではありません。
問題 2: 彼ら自身のデータが最も恐ろしい主張を打ち消す
彼ら自身の数字を見てください。応募者の 84% が 1 つのポジションにのみ応募しました。 95% 以上が 1 つまたは 2 つに当てはまります。ちょうど 522 人 (0.02%) が 10 件に応募しました。 「どこでも拒否される」という悪夢を実現するには、同じベンダーを通じて広く申請する必要があります。ほとんど誰もしませんでした。
代表的な「10 件に応募、10 件から拒否」という統計は、見逃してしまうほど薄いデータの一部から抽出されたもので、効果を大きく見せるために、人々が実際よりもはるかに広範囲に応募する世界をシミュレートする必要がありました。
このメカニズムは存在しますが、実際に保持しているデータは薄いです (注意: このデータセットでは、pymetrics を介したアプリケーションのみが表示されます。誰かが 30 個のジョブに応募し、pymetrics に 1 回ヒットする可能性があります。そのため、「1 回適用された」というだけでも、同じベンダーのトラップが発生する可能性がいかに低いかを過小評価しています)。
「でも今は人々が大量に応募しているよ」とあなたは言うかもしれない。そうです。そしてそれは紙に当たります。彼らのデータは、AI アプリケーションが急増する前の 2022 年に終了します。さらに言えば、その急増はスムーズに実行されます。ワンクリックで適用、キーワード フィルター、履歴書が何百件も実行されます。
16 種類の行動ゲームをワンクリックで完了することはできません。彼ら自身のデータがそれを示しています。応募者の平均応募数は 1.24 件でした。パイメトリクスを大規模に処理する人は誰もいませんでした。フォーマットがそれを不可能にしているからです。大量応募の時代には、研究で測定されたものとは異なるシステムが溢れかえりました。
問題 3: 「証明」は実際には一致しない比較に基づいています
j ではなく共有アルゴリズムによる拒否クラスターを表示するには

通常の採用では、新聞には比較するものが必要です。単一の共有アルゴリズムが存在しない世界。
彼らは古い研究を使用しています: Kline et al. (2022 年)、米国の大手企業 108 社に 83,000 通の履歴書が送信されました。そのデータでは、拒否は独立しているように見え、別々の企業による別々の決定のように見えます。パイメトリクスでは、それらは一緒に固まります。したがって、アルゴリズムが原因であるに違いないと論文は述べています。
これは「通常の採用」のベンチマーク、つまり 108 社の人的スクリーニングです。実際に起こった点 (暗い) は、すべての雇用主が独自に決定した場合に予想される点 (明るい) とほぼ正確に一致しています。
素敵な話。落とし穴: 2 つのデータセットはほとんど似ておらず、論文もそれを認めています。
クラインの調査では、履歴書の約 23 ～ 25% に電話連絡がありました。パイメトリクスでは、約 50% が「合格」
Kline は米国のエントリーレベルの求人のみをカバーしています。 pymetrics は、エントリーレベルからディレクターまで、世界中のあらゆるものをカバーしています。データ内で最も一般的な都市はアメリカですらない (ロンドンです)
クラインは人間のコールバックを測定します。 pymetrics は 1 つのアルゴリズムのスコアを測定します。つまり、実際には「アルゴリズムがあるかないか」というわけではありません。これは、異なる国の、異なる職務レベルで、まったく異なる 2 つのプロセスです。
著者はさまざまな合格率に合わせて数学的に調整するため、結果を合格率のアーチファクトとして無視することはできません。しかし、調整はそこで止まり、地理、年功序列、または人間とアルゴリズムのギャップについては何もしません。誰がプールにいるか、誰が決定を下したかではなく、人々が通過する頻度を補正します。
そして彼ら自身の結論は、クラスタリングが「集中アルゴリズム評価の独特の特徴である可能性がある」というものです。 5月。これは正直な、ヘッジされたバージョンであり、合理的な仮説であり、証明された事実ではありません。新聞には「かもしれない」と書かれていたが、見出しには「そうなる」と書かれていた。
問題 4: 「どこでも拒否される」を実際にテストしたところ、問題は発生しませんでした

n
この論文の感情的な核心は、同じモデルがどこでもあなたを拒否するため、「アルゴリズム的にブラックボール」になり、雇用市場全体から凍結されるという考えです。
そこで彼らはそれをテストしました。彼らは 1,000 人の応募者を受け入れ、それぞれをすべての pymetrics モデル (495 つすべて) に対して実行して、誰がロックアウトされるかを確認しました。
その結果は、彼ら自身の言葉で言えば、「サンプリングされた応募者はすべてのパイメトリクス モデルによって拒否されない」ということです。一つもありません。最悪の状況にあった 1 人の人でも、52 件の仕事から推薦を受けました。
彼らが悪夢のようなシナリオを直接シミュレーションしたところ、どこからでも誰も締め出されることはありませんでした。誰もが、「雇用可能」よりも弱いどこかの少なくとも 1 つのモデルによって推薦されています (推薦は、まだ人間がレビューする山の中にあなたを置くだけです) が、それは「市場全体にブラックボール」バージョンを完全に下回っています。
次に、「25 件の求人に応募する必要がある」という一文が続きます。しかし、この数字はでっち上げられた世界から来ています。各申請者の実際の約 1.2 件の申請を 145 件の仮想的な申請に膨らませてから数えたのです。模型の上に模型を重ねたものです。
そして、どこでも拒否される人はいないデータセットで、「どこでも拒否される」ことを避けるために、25 の仕事に応募するように人々に伝えるという報道がアドバイスに変わるのは奇妙なことだ。
人々が応募しないわけではない

[切り捨てられた]

## Original Extract

Are you afraid AI is rejecting you everywhere? Someone

Fear-farming the Stanford AI hiring study [debunk]
Two things are true about the new Stanford hiring study: the finding everyone quotes is real, and much of the commentary around it seems to have stopped at the abstract.
That gap is how a narrow result about one badly built tool turned into a verdict on the whole industry and a new way for career influencers to farm job-seeker panic.
The viral Stanford paper ( Algorithmic Monocultures in Hiring ) is real research and worth reading, but it studied one narrow, game-based tool called pymetrics, not "AI hiring" as a whole. Most of the panic comes from people generalizing one vendor into an entire industry.
The one finding worth keeping: company-wide fairness audits can look clean while individual jobs discriminate. Pooled together, pymetrics passed the usual adverse-impact screen, but split into 1,746 positions, ~11% worked against Black applicants. If you audit hiring only at the company level, you should change it!
The scary "rejected everywhere" story barely exists in the data. 84% of applicants applied to just one position; only 0.02% applied to ten. When the authors directly simulated the nightmare, not one person was rejected by every model.
The tool itself is the real story. pymetrics trains on a company's current employees as "good" and random profiles as "bad", so it learns who resembles the existing staff, not who can do the job. And there's no evidence in the study that it predicts performance at all. It's biased and unvalidated.
The authors admit nearly all of this in the limitations section. The research was careful and hedged. What wasn't careful was the chain that turned it into "AI is rejecting you everywhere", every step dropped a qualifier the authors put there on purpose.
The paper in question: Algorithmic Monocultures in Hiring (FAccT 2026).
The research paper is called Algorithmic Monocultures in Hiring. Its core idea is the concept of "monoculture": if every employer runs candidates through the same handful of AI vendors, then one biased model doesn't just hurt you at one company, it locks you out everywhere at once. Reasonable worry, and as it turns out, reasonable doubts.
To study it, the researchers got four years of real hiring data from a single vendor and checked whether the same people and the same racial groups, kept getting filtered out.
So far, so good. Then you find out which vendor.
Resemble the people already hired, or don't. That's the entire science.
It looks at vendor named pymetrics. Not résumés and not an applicant tracking system.
pymetrics is a tool where you play games. Yes, games. You play 12 to 16 of them that supposedly measure things like risk-taking, processing speed, planning. 12 of those games are the same whether you're applying to a warehouse or a finance desk.
A model scores your gameplay and spits out one of two words: "recommend" or "do not recommend” (why companies pay for that?). That's it. Oh, and about 42% of the time, it says "do not recommend”.
pymetrics process authors describe in the paper
For each company, the model is trained on at least 50 of that company's current employees as the examples of "good" and random people as "bad". No, the "good" group isn't top performers.
It's whoever currently holds the job. At least in this study, nothing shows those incumbents were actually strong performers.
And the "bad" group isn't people who failed the role, or got fired, or underperformed. It's random strangers.
So the model never learns the thing you'd actually want like what separates a strong hire from a weak one for your organization. It learns something far dumber: do you look like the people already on the payroll or do you look like a random person off the street?
As a bonus, your score even gets cached for 330 days. Apply elsewhere within the year and you're often judged on the same saved gameplay, twice. Yay!
But hey, at least the dataset is genuinely big: 4.2 million applications, 3.4 million people and 156 employers.
pymetrics had previously checked its own fairness by pooling every applicant together. Pooled like that, it passed the usual adverse-impact screen. Black applicants passed at 52.5%, white at 58.3%. The researchers correctly pointed out that this is the wrong way to validate a selection tool.
US employment law (Title VII) evaluates each job separately , not the company-wide blend. So they split the data into 1,746 positions and looked again. Bias the average had hidden showed up: roughly 11% of jobs worked against Black applicants (10.6% after statistical correction), and about a quarter of Black applications landed in those jobs.
This is the contribution, and it's a good one. Averages can hide discrimination. If you only audit your hiring at the company level, this is definitely something to consider. Good takeaway, worth keeping.
it raises a data-governance question worth flagging for anyone buying these tools.
In standard corporate compliance, official Equal Employment Opportunity Commission (EEOC) data is strictly firewalled. It is completely optional for candidates and sits in an isolated repository that legally cannot feed into automated assessment engines or decision-making tools.
So what's the race data doing inside the assessment vendor's dataset, mapped to scores across 156 employers?
provided by the authors of the research paper
The answer is mundane, and that's the point. It's voluntary self-ID , collected by pymetrics. Only 40% of applicants gave it and each employer chooses whether to let the vendor collect it at all. All legal. All standard. And a completely different stream from the protected EEO-1 filing.
To sum it up, your candidates are handing demographic data to a third party, on a screen you don't control, governed by a setting your procurement team probably never reviewed.
Is the vendor doing anything illegal? No. Is your company-level compliance audit ever going to see this? Also no. Well - yay?
Now the problems with the paper.
Problem 1: They proved one tool is bad, then talked like all of them are
The whole paper is built on a word: MONOCULTURE. The idea that everyone gets funneled through the same dominant system, so one biased model poisons the whole market.
To make that sound real and scary, the opening cites HireVue: used by over 60% of the Fortune 100 and 8 of the 10 biggest federal agencies. Frightening number, right?
Except they didn't study HireVue. They studied pymetrics. A different, much smaller company. A competitor. HireVue is a different product category: structured interviews and job previews, methods that can be scientifically valid when properly designed. But this paper doesn't test HireVue's version of them, so I won't pretend it validates them either. pymetrics does games.
And nobody ever explains how those Fortune 100 companies actually use HireVue. It may be one stage in a process that also has structured interviews, work samples, reference checks.
Researcher's own results section admits the scary version of monoculture - the same model judging you at multiple companies - happens in, their words, "rare instances" with "limited" impact, "because very few applicants apply to positions at different employers served by the same underlying pymetrics model".
So they switch to studying a weaker, looser definition instead. The headline concept barely shows up in the data. They led with HireVue's scale to make it sound everywhere but their own numbers say it isn't.
Problem 2: Their own data kills the scariest claim
Look at their own numbers. 84% of applicants applied to exactly one position. Over 95% applied to one or two. Exactly 522 people - 0.02% - applied to ten. The "rejected everywhere" nightmare needs you to apply widely through the same vendor. Almost nobody did.
The flagship "apply to 10, rejected from 10" stat is pulled from a sliver of the data so thin you could miss it and to make the effect look bigger, they had to simulate a world where people applied far more broadly than they actually do.
The mechanism is there, but thin in the data they actually have (and note: this dataset only sees pymetrics-mediated applications. Someone might apply to thirty jobs and hit pymetrics once. So even "applied once" understates how rarely the same-vendor trap could spring).
"But people mass-apply now," you might say. True. And it cuts against the paper. Their data ends in 2022, before the AI-application surge. More to the point, that surge runs on frictionlessness: one-click apply, keyword filters and résumés fired off by the hundred.
You cannot one-click through 16 behavioral games. Their own data shows it. Applicants averaged 1.24 applications. Nobody was grinding through pymetrics at scale, because the format makes it impossible. The mass-apply era floods a different system than the one the study measured.
Problem 3: The "proof" leans on a comparison that doesn't really match
To show the rejections cluster because of the shared algorithm and not just normal hiring, the paper needs something to compare against. A world with no single shared algorithm.
They use an older study: Kline et al. (2022), 83,000 résumés sent to 108 big US companies. In that data, rejections look independent, like separate decisions by separate companies. In pymetrics, they clump together. So, the paper says, the algorithm must be the cause.
This is the "normal hiring" benchmark - human screening at 108 companies. The dots that actually happened (dark) sit almost exactly on the dots you'd expect if every employer decided on their own (light).
Nice story. The catch: the two datasets barely look alike, and the paper admits it.
In the Kline study, about 23–25% of résumés got a callback. In pymetrics, about 50% "pass"
Kline only covers entry-level US jobs. pymetrics covers everything from entry-level to director, all over the world. The most common city in the data isn't even American (it's London)
Kline measures human callbacks. pymetrics measures one algorithm's score. So it's not really "algorithm vs. no algorithm". It's two totally different processes, at different job levels, in different countries.
Authors math adjusts for the different pass rates, so you can't dismiss the finding as a pass-rate artifact. But the adjustment stops there, it does nothing about the geography, the seniority, or the human-vs-algorithm gap. It corrects for how often people pass, not who's in the pool or what's making the decision.
And their own conclusion is that the clustering "may be a distinctive feature of centralized algorithmic assessment". May. That's the honest, hedged version, a reasonable hypothesis, not a proven fact. The paper said "may", but the headlines said "does".
Problem 4: When they actually tested "rejected everywhere" it didn't happen
The paper's emotional core is the idea of being " algorithmically blackballed ", frozen out of the whole job market because the same model rejects you everywhere.
So they tested it. They took 1,000 applicants and ran each one against every pymetrics model - all 495 of them - to see who'd be locked out.
The result, in their own words: "no sampled applicant is rejected by every pymetrics model". Not one. The single worst-off person still got recommended by 52 jobs.
When they directly simulated the nightmare scenario, nobody was locked out everywhere. Everyone was recommended by at least one model somewhere which is weaker than "hireable" (a recommendation just puts you in a pile a human still reviews), but it flatly undercuts the "blackballed across the whole market" version.
Then comes the "you'd need to apply to 25 jobs" line. But that number comes from a made-up world: they inflated each applicant's real ~1.2 applications into 145 hypothetical ones, then counted. It's a model on top of a model.
And it's a strange thing for the coverage to turn into advice, telling people to apply to 25 jobs to avoid being "rejected everywhere" in a dataset where no one was rejected everywhere.
Not that people aren't applyi

[truncated]
