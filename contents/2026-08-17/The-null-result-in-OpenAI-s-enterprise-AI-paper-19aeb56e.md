---
source: "https://theworkingmodel.co/analysis/null-result-openai-enterprise-paper/"
hn_url: "https://news.ycombinator.com/item?id=49327276"
title: "The null result in OpenAI's enterprise AI paper"
article_title: "The null result in OpenAI's enterprise AI paper — The Working Model"
author: "aicoding"
captured_at: "2026-08-17T07:43:52Z"
capture_tool: "hn-digest"
hn_id: 49327276
score: 1
comments: 0
posted_at: "2026-08-17T06:51:36Z"
tags:
  - hacker-news
  - translated
---

# The null result in OpenAI's enterprise AI paper

- HN: [49327276](https://news.ycombinator.com/item?id=49327276)
- Source: [theworkingmodel.co](https://theworkingmodel.co/analysis/null-result-openai-enterprise-paper/)
- Score: 1
- Comments: 0
- Posted: 2026-08-17T06:51:36Z

## Translation

タイトル: OpenAI のエンタープライズ AI 論文のヌル結果
記事のタイトル: OpenAI のエンタープライズ AI 論文の null 結果 — The Working Model
説明: OpenAI

記事本文:
実用モデル
ホーム
分析
ハイプオメーター
OpenAI のエンタープライズ AI 論文のヌル結果
OpenAI は今週、Columbia と Wharton と共同で、2026 年 3 月までの ChatGPT Enterprise テレメトリーに基づいて構築された作業報告書を発表しました。誰もが引用する調査結果は要約にあります。時間をかける価値があるのは、論文で報告され、その後過去に進む表 2 の係数です。
良い紙ですよ。同社は自社の顧客に関するデータを公開しているベンダーでもあり、給与明細には 2 人の共著者が含まれているため、そのカテゴリの他の記事を読むのと同じように読んでください。
見出しの調査結果はご想像の通りです。使用量は 7 倍に増加しました。採用者の偏りは大きい。タスクの使用は集中的ではなく広範囲に行われます。どれももっともらしいが、特に驚くべきものはない。
私が何度も戻ってくるのは別の場所です。
すでに導入している企業では、従業員数が多いほど従業員 1 人あたりの使用量が少なくなることが予測されます。 1 人あたりのメッセージ数、1 人あたりの週次アクティブ ユーザー数、1 人あたりのトークン数が減少しており、これらはすべて統計的に有意です。これは、スケーリング効果として合理的に表されます。大企業は希薄化する。
ただし、そのテーブルには 4 つの列があり、4 番目は週ごとのアクティブ ユーザーごとのメッセージです。
括弧内は企業ごとにまとめられた標準誤差。 ∗∗∗ は 1% での有意性を示します。 3 つの有意な陰性と 1 つのヌル。 null は何をすべきかを指示するものです。
何もない。フラット。標準誤差は推定値の 4 倍のサイズです。従業員数 200,000 人の企業の熱心な AI ユーザーは、従業員数 500 人の企業の熱心な AI ユーザーと同じくらい熱心です。
これに頼る前に、正直な注意点が 1 つあります。 4 番目の回帰の R² は 0.100 ですが、他の 3 つの回帰では 0.36 ～ 0.48 です。これはノイズの多い仕様であるため、ヌルは、厳密に適合したモデルのヌルよりも弱い証拠となります。しかし、点推定は本質的にゼロであり、単なるインプではありません

読み上げても、サインは何も反転しません。
つまり、大企業の使用状況の差は、ユーザーがツールを使用した後の使い方とは何の関係もありません。それは完全に何人がそれを使用しているかの問題です。
問題がエンゲージメントにある場合は、トレーニングを購入します。問題が侵入である場合、そのどれも制約には影響しません。
2 つの読み方はまったく異なる場所にあなたを導くので、これは小さな違いではないと思います。問題がエンゲージメントである場合は、より良いプロンプト、より良いトレーニング、より良い社内伝道、より良いユースケース ライブラリを購入することになります。それが今のイネーブルメント業界全体です。問題が侵入である場合は、2 回目にログインしなかったのは誰で、その理由を確認する必要があります。
最初のことを測定し、それが 2 番目のことを説明すると仮定するのは簡単です。この論文は、そうではないことを示すまともな証拠です。
データからは分からないのは、なぜ非返品者が返品しないのかということであり、それが null の結果によって緊急に生じる問題です。このデータセットには、誰が座席を提供され、誰がそれを断ったのかを示す分母はなく、開催されなかった 2 番目のセッションの記録もありません。そこで起こっていることは何であれ、イネーブルメント機構が機能する前に起こっています。これは、オンボーディング後の段階をほぼ完全に中心に組織してきた業界にとって、かなり厄介な発見です。
2 番目に注目すべき点は、そもそも採用を予測するものです。
研究者らは、これらの企業が何かを購入するかなり前の2021会計年度に測定された、従業員1人当たりの研究開発株、従業員1人当たりの資本化されたソフトウェア、従業員1人当たりの販管費の3種類の蓄積された無形資本をテストした。ここでの販管費は、組織能力と管理能力の粗雑な代用手段です。それは最も強力で最も堅牢なものになりました。物的資本の集約度は別の方向に作用した

スケールを制御すると、導入に悪影響を及ぼします。
したがって、最初に移行した企業は、最高のエンジニアリングや最も深いソフトウェア資産を持っている企業ではありませんでした。彼らは、仕事の進め方を再設計するオーバーヘッド機能の構築にすでに何年も費やしてきた人たちでした。地味な層。景気後退の際に最初に切り捨てられ、誰も決算発表を行わない層。
私はこれに不快感を感じますが、おそらく正しいと思います。これは、汎用テクノロジーについて私たちが知っているすべてのことと一致しており、特定の業界で技術的に最強の組織が最も早く適応できる組織ではないことが多いというパターンとも一致しています。吸収する能力は構築する能力とは別の資産であり、私たちはそれを測定したり資金を提供したりすることがあまり得意ではありません。
それはまた、かなり厳しい短期的な動向を示唆している。これを吸収するのに最適な立場にある企業が最大かつ最も価値のある企業であり、モデル自体が誰にとっても同様の価格で広く入手可能である場合、一般入手可能性は何も平準化されません。それはギャップを広げます。同紙はそのことを簡潔に述べた後、そのまま放置している。
3 番目の部分は、内部的に見たいものを変更するものです。
導入から 6 か月後、ほとんどのガバナンスの設計方法と比べて、年功序列による強度の勾配が間違った方向に進んでいます。初期のキャリアの従業員や研修生は、同じ企業の平均的なアクティブ ユーザーよりも 1 週間におよそ 8 ～ 9 多いメッセージを送信します。経営陣、創設者、パートナーからの派遣は減ります。アナリストやマーケティング、コミュニケーションのスタッフも熱心に取り組んでいます。
そして、タスクの組み合わせも同じ線に沿って分割されます。サンプル全体を通して、主なアクティビティはドキュメントと技術文書であり、全アクティブ ユーザーの半数以上が技術的なデジタル作業とメッセージの草稿です。 P

生産。幹部は、トピックの概要、事実と数字、法律と規制に関する質問、財務と税金の質問など、別の場所に集まります。消費と合成。
どれもそれ自体はスキャンダルではありません。しかし、それをまとめると、運用状況は、最も経験の浅い人々がモデルの支援を受けて大量の文書を作成し、最も上級の人々が要約を読んでおり、彼らの間にあるレビュー層が、ジュニアの成果物がより遅く、明らかに労力がかかる世界向けに設計されているということです。
タスクのリスクについても関連する調査結果があります。法律、規制、財務、税務のタスクは、ユーザーリーチの観点からは広範囲に及んでいますが、メッセージシェアの観点からは小規模です。たまに触る人も多いです。ボリュームで監視している場合 (ボリュームは管理コンソールで提供されるものであるため、ほとんどの組織ではそうなっています)、これらのカテゴリのサンプリングは最も少なく、不良な出力によるコストが最も高くなるカテゴリです。これは簡単に修正できる測定誤差であり、ほぼ普遍的なものだと思います。
ここでは通常よりも重要なため、いくつかの注意点があります。
上記のものはすべて相関関係があります。この点については同紙も注意しているし、私も注意するつもりだ。 API、内部ツール、他のベンダー、個人アカウントではなく、ChatGPT Enterprise 内でのみ使用量を測定します。非導入企業として分類された企業は、このデータセットでは認識できないような重い AI プログラムを実行している可能性があります。これにより、論文内の導入企業と非導入企業の比較がすべて弱められます。役職はその時点のものであり、不完全です。この財務分析は米国の公開企業のみを対象としているため、中堅企業や非公開企業については何もわかりません。そして重要なことに、これらはどれも生産量を測定するものではありません。品質でも、収益でも、従業員数でもありません。使用量は生産性ではありませんし、著者も生産性であるとは決して主張しません。
私が考えたフレーミング

最も強くプッシュバックするのは暗黙的なものであり、トークンが多いほど良いということです。論文の強度測定のほとんどは体積測定です。企業が従業員あたりのメッセージで下位 4 分の 1 に属し、上位の企業よりも多くの価値を引き出している可能性は十分にありますが、このデータセットにはそれらを区別するものは何もありません。
この論文が証明していること、そして私が支持していると思うことは、その結論にあります。採用は展開の始まりにすぎません。企業はこれを使用するかどうかをまだ決定していません。彼らはそれがどこに属するのかを検討していますが、そのプロセスは時間がかかり、組織的であり、外部からはほとんど見えません。
今後 2 年間で広がるギャップは、モデルへのアクセスによってもたらされるものではありません。誰もがそれを持っています。それは契約締結後 18 か月間に何が起こるかによって決まるが、その期間を十分に測定して契約を獲得できるかどうかを知る人はほとんどいない。
Chatterji、A.、Holtz、D.、Rakholia、N.、Tambe、P.、Weeratunga、G. (2026)。
組織による AI の使用方法: ChatGPT の証拠。ワーキングペーパー、2026 年 8 月 11 日。
PDF
ワーキング モデルは、調査、提出書類、リリース ノートを読み取って、実際に何が変更されたかを通知します。無料。

## Original Extract

OpenAI

The Working Model
Home
Analysis
Hype-o-meter
The null result in OpenAI's enterprise AI paper
OpenAI published a working paper this week with Columbia and Wharton, built on ChatGPT Enterprise telemetry through March 2026. The findings everyone will quote are in the abstract. The one worth your time is a coefficient in Table 2 that the paper reports and then moves past.
It's a good paper. It's also a vendor publishing data about its own customers, with two co-authors on its payroll, so read it the way you'd read anything else in that category.
The headline findings are what you'd expect. Usage grew sevenfold. Adopters skew large. Task use is broad rather than concentrated. All plausible, none of it especially surprising.
What I keep coming back to is elsewhere.
Among firms that have already adopted, larger headcount predicts lower usage per employee. Fewer messages per head, fewer weekly active users per head, fewer tokens per head, all statistically significant. This is presented, reasonably, as a scaling effect. Big companies dilute.
But there are four columns in that table, and the fourth is messages per weekly active user.
Standard errors clustered by firm in parentheses; ∗∗∗ denotes significance at 1%. Three significant negatives and one null. The null is the one that tells you what to do.
Nothing. Flat. A standard error four times the size of the estimate. A 200,000-person company's engaged AI users are as engaged as a 500-person company's engaged AI users.
One honest caveat before I lean on this. That fourth regression has an R² of 0.100, against 0.36 to 0.48 for the other three. It's a noisier specification, so the null is weaker evidence than a null in a tightly fitted model would be. But the point estimate is essentially zero, not merely imprecise, and the sign flips nothing.
Which means the large-enterprise usage gap has nothing to do with how people use the tool once they're using it. It's entirely a question of how many people are using it at all.
If your problem is engagement, you buy training. If your problem is penetration, none of that touches the constraint.
I don't think this is a small distinction, because the two readings send you to completely different places. If your problem is engagement, you're buying better prompts, better training, better internal evangelism, better use case libraries. That's the entire enablement industry right now. If your problem is penetration, you should be looking at who never logged in a second time and why.
It's easy to measure the first thing and assume it explains the second. This paper is decent evidence that it doesn't.
What the data can't tell you is why the non-returners don't return, and that's the question the null result makes urgent. The dataset has no denominator for who was offered a seat and declined it, no record of the second session that never happened. Whatever is happening there is happening before any of the enablement machinery gets a chance to work, which is a fairly awkward finding for an industry that has organised itself almost entirely around the post-onboarding phase.
The second thing worth sitting with is what predicts adoption in the first place.
The researchers tested three kinds of accumulated intangible capital, all measured in fiscal 2021, well before any of these firms bought anything: R&D stock per employee, capitalised software per employee, and SG&A stock per employee. SG&A is the crude proxy here for organisational and managerial capability. It came out strongest and most robust. Physical capital intensity ran the other direction, negatively associated with adoption once you control for scale.
So the firms that moved first weren't the ones with the best engineering or the deepest software estate. They were the ones that had already spent years building the overhead function that redesigns how work happens. The unglamorous layer. The layer that gets cut first in a downturn and that no one puts in an earnings call.
I find this uncomfortable and probably correct. It's consistent with everything we know about general purpose technologies, and it's consistent with the pattern where the technically strongest organisation in a given industry is frequently not the one that adapts fastest. Capability to absorb is a separate asset from capability to build, and we're not very good at measuring or funding it.
It also suggests a fairly grim near-term dynamic. If the firms best positioned to absorb this are also the largest and most valuable ones, and the models themselves are broadly available at similar prices to everyone, then general availability doesn't level anything. It widens the gap. The paper says as much, briefly, and then leaves it alone.
The third piece is the one that changes what I'd want to look at internally.
Six months after adoption, the intensity gradient by seniority runs the wrong way relative to how most governance is designed. Early-career workers and trainees send roughly eight to nine more messages per week than the average active user at the same firm. Executives, founders and partners send fewer. Analysts and marketing and communications staff also run hot.
And the task mix splits along the same line. Across the whole sample, the dominant activities are documentation and technical writing — more than half of all active users — technical digital work, and drafting messages. Production. Executives cluster somewhere else: topic overviews, facts and figures, legal and regulatory questions, financial and tax queries. Consumption and synthesis.
None of that is scandalous on its own. But put it together and the operating picture is that your least experienced people are producing a high volume of written material with model assistance, your most senior people are reading summaries, and the review layer between them was designed for a world where the junior output was slower and more obviously effortful.
There's a related finding on task risk. Legal, regulatory, financial and tax tasks show up as widespread by user reach but small by message share. A lot of people touch them occasionally. If you're monitoring by volume, which most organisations are because volume is what the admin console gives you, you will sample those categories least and they are the ones where a bad output is most expensive. That's a straightforwardly fixable measurement error and I'd guess it's near-universal.
Some caveats, because they matter more than usual here.
Everything above is correlational. The paper is careful about this and I'll be careful too. It measures usage only inside ChatGPT Enterprise: not the API, not internal tooling, not other vendors, not personal accounts. Firms classified as non-adopters may be running heavy AI programmes that this dataset simply cannot see, which attenuates every adopter-versus-non-adopter comparison in the paper. Job titles are point-in-time and incomplete. The financial analysis covers US public companies only, so it tells you nothing about mid-market or private firms. And critically, none of this measures output. Not quality, not revenue, not headcount. Usage is not productivity and the authors never claim it is.
The framing I'd push back on hardest is the implicit one, which is that more tokens is better. Most of the paper's intensity measures are volume measures. It's entirely possible for a firm to be in the bottom quartile on messages per employee and extracting more value than someone in the top, and nothing in this dataset would distinguish them.
What the paper does establish, and what I think holds up, is in its conclusion: adoption is only the beginning of deployment. Firms aren't deciding whether to use this anymore. They're working out where it belongs, and that process is slow, organisational, and mostly invisible from outside.
The gap that opens over the next two years won't come from model access. Everyone has that. It'll come from what happens in the eighteen months after the contract is signed, and almost nobody is measuring that period well enough to know whether they're winning it.
Chatterji, A., Holtz, D., Rakholia, N., Tambe, P., & Weeratunga, G. (2026).
How Organizations Use AI: Evidence from ChatGPT. Working paper, 11 August 2026.
PDF
The Working Model reads the research, the filings and the release notes so you don't have to, and tells you what actually changed. Free.
