---
source: "https://antoine.fi/mri-analysis-using-claude-code-opus"
hn_url: "https://news.ycombinator.com/item?id=48708941"
title: "I used Claude Code to get a second opinion on my MRI"
article_title: "Using Opus 4.8 to get a second opinion on an MRI and where it leaves me"
author: "engmarketer"
captured_at: "2026-06-28T17:26:23Z"
capture_tool: "hn-digest"
hn_id: 48708941
score: 39
comments: 56
posted_at: "2026-06-28T16:35:19Z"
tags:
  - hacker-news
  - translated
---

# I used Claude Code to get a second opinion on my MRI

- HN: [48708941](https://news.ycombinator.com/item?id=48708941)
- Source: [antoine.fi](https://antoine.fi/mri-analysis-using-claude-code-opus)
- Score: 39
- Comments: 56
- Posted: 2026-06-28T16:35:19Z

## Translation

タイトル: クロード コードを使用して MRI のセカンドオピニオンを取得しました
記事のタイトル: Opus 4.8 を使用して MRI のセカンドオピニオンを得る、そしてそれが私にもたらす影響
説明: クロード コードを使用して肩の怪我についてセカンドオピニオンを得ようとした私の経験。これは将来を垣間見るのに興味深いものですが、その結果にまだ完全な自信を持つのは難しいです。

記事本文:
アントワーヌのブログ
ホーム
トレーニング
について
Opus 4.8 を使用して MRI に関するセカンドオピニオンを取得し、それが私にもたらす影響
この記事は、Opus 4.8 を使用して MRI の結果を読み取り、診断に関するセカンドオピニオンのようなものを提供した私の経験についてです。もちろん、このテクノロジーがまだ存在していない可能性があることは承知しています。そのため、この記事を共有しています。もしかしたらそれは誰かを助けたり、少なくともちょっとした情報や娯楽を提供したりできるかもしれません。
免責事項：私はもちろん医師ではありません（これが実際に問題なのです！）ので、私の言うことはすべて割り引いて聞いてください。
いくつかのコンテキスト (スキップしても構いません)
数週間前から右肩に痛みを感じています。良くなってきたように見えましたが、整形外科医の意見を聞くことにしました。詳細は述べませんが、彼はクリニックで都合よく利用できるMRI検査を受けることを勧めました。私も同意し、主に肩甲下筋腱の「根尖付着部にグレードIII（幅>50％）の部分層断裂」があることを知りました。もちろん、これは私にとってほとんど意味がありませんが、彼らが提案した治療方針は広範でした。 MRIを撮ってから数分後にも始まりました。診療所から出てきたとき、私は彼らが銃を飛び越えたような気がしました。
ありがたいことに、出発する前に、MRIの結果のコピーと行ったすべての治療のリストを送ってもらい、合計3回繰り返すよう提案しました。
すべてを GPT 5.5 Pro に送信すると、すぐに次の 2 つの点にフラグが立てられました。
最近の臨床診療ガイドラインでは石灰化のない腱板腱障害に対して臨床医は衝撃波療法を使用したり推奨したりすべきではないとしているにもかかわらず、彼らは私の肩に衝撃波療法を実施しました。超音波検査では石灰化はないと言われました。
彼らは私にTraumeelを注射しました。

ドイツでは「治療適応のない」ホメオパシー医療。
それで私の自信は高まりませんでした。それで、MRIを分析することに興味を持ちました。
MRI の最初のレビューを行うために Opus をセットアップする
MRI パッケージは、拡張子のない数百のファイルを含む標準的な DICOM エクスポートで、合計約 266 MB でした。
分析のために、Claude Code 内で Opus 4.8 (xhigh) を使用して、コードを実行してパッケージをインストールできるようにすることにしました。作業が完了する前に、分析に必要なパッケージをインストールするように指示しました。クロード コードを使用することは、この問題に関して大量の作業を実行できるようにするために特に重要です。プログラマーにとっては明らかなことかもしれませんが、Claude Code と Claude.ai のチャットの違いは、たとえ 2 つが同じモデルを実行したとしても、非常に大きなものです。
それから始める時が来ました。 MRIのことは何も知らない私としては、クロードに綿密な計画を立ててから実行してもらうという設定にしました。私が出した唯一の指示は「右肩の痛みが2～3週間続く」というものでしたが、人間の医師が受けた指示よりも軽いことが後で分かりました。
約 1 時間後、次のようなレポートが返されました。
そのレポートの重大な問題は、医師が根尖挿入部にグレード III (50% 以上) の部分層断裂を認めたにもかかわらず、Opus 4.8 では腱は無傷であると報告したことです。
これにはかなり当惑しました。成績はもっと低いだろうと予想していましたが、その結果は極端でした。
判断するために、私はクロードに 2 つのレポートを比較してもらうことにしました。しかし今回は、もう少しコンテキストを与えました。人間のレポートを提供することに加えて、ChatGPT 5.5 Pro で行ったディスカッションも提供しました。そこで、私の診断が何であるかを理解する方法として試せる動きや位置を与えてもらいました。
計画文書から、Opus がとったアプローチは次のとおりです。
の

このアプローチは慎重かつ系統的であり、既存のコンテキストに偏らない新しい分析を取得する方法として複数のサブエージェントが使用されました。
再び、約 1 時間後に新しいレポートが届きました。
仲裁人の評決: 証拠は読者 A に有利です (中程度から高い信頼度)。軽度の挿入部腱炎。根尖挿入部を含め、個別の部分層または全層の断裂は確認されませんでした。
判決が互いにかけ離れているのは興味深いと思わずにはいられません。この報告書をさらに詳しく見てみると、オーパス社は、2 つの報告書の間には解決できないいくつかの論争があるが、今回の報告書では解決できると言うのをためらわなかったことがわかります。そして非常に決断的に。
信頼できる専門家の手に委ねられると、信じられないほど心が安らぎます。もう心配する必要はなく、プロセスをガイドしてもらうことができます。
AI は、その感情を不快な方法で完全に打ち砕くことができます。この AI 主導のセカンドオピニオンを取得した後、診断と治療計画は時期尚早で、事実が正当であると思われるよりも介入が多いように見えます...しかし、私も AI を完全に信頼できるかどうかはわかりません。そのため、私は別の医師に診てもらうか、今行っているリハビリで肩が良くなるか様子を見るかのどちらかという途方に暮れた状態に置かれています。
私の希望は、モデルの数世代後には、私たちが電子メールの校正に AI を信頼しているのと同じように、AI が MRI の検査を信頼できるようになることです。
これは記事の主旨ではないので、クリニックや医師の名前は出しません。 AI を使用してセカンドオピニオンを取得することについての私の技術的な好奇心を共有することです。私が間違っているかもしれないし、AI が間違っているかもしれない。医師の誤解もあったかもしれません。したがって、基本的に、これは医学的アドバイスとして受け取られるべきではありません:)
このブログの内容をお楽しみいただければ幸いです。何かご意見がございましたらお気軽にご連絡くださいw

ぜひご連絡ください。
Ruby on Rails 上に構築されたカスタム ブログ エンジンで作成されました :)

## Original Extract

My experiencing trying to get a second opinion on my shoulder injury using Claude Code. It's a fascinating glimpse into the future but it's hard to have full confidence in the results just yet.

Antoine's blog
Home
Workouts
About
Using Opus 4.8 to get a second opinion on an MRI and where it leaves me
This article is about my experience using Opus 4.8 to read the results of an MRI and give me a sort of second opinion on the diagnosis. Of course, I know the technology might not be there yet, which is why I'm sharing this article. Maybe it can help someone or at least provide a bit of information or entertainment.
Disclaimer: I'm of course not a doctor (this is actually the problem!) so please take everything I say with a grain of salt.
Some context (feel free to skip)
For a few weeks now, I've been experiencing some pain in my right shoulder. Even though it seemed to be getting better, I decided to get an opinion from an orthopedist. I won't go into the details, but he suggested I get an MRI, which the clinic conveniently had available. I agreed and mainly learned that I had a "Grade III (>50%-width) partial-thickness tear at the apical insertion" of my subscapularis tendon. This, of course, means little to me, but their suggested course of treatment was extensive; they even started a few minutes after I got the MRI. Coming out of the clinic, I had the feeling they had jumped the gun.
Thankfully, before I left, I asked them to send me a copy of the MRI results and a list of all the treatments they performed and suggested we repeat a total of 3 times.
I sent everything over to GPT 5.5 Pro, and right away it flagged two things:
They performed shockwave therapy on my shoulder even though a recent clinical practice guideline says clinicians should not use or recommend shockwave therapy for rotator-cuff tendinopathy without calcification; I was told during ultrasound that there was no calcification.
They injected me with Traumeel, which is registered in Germany as a homeopathic medicine "without a therapeutic indication".
That did not increase my confidence. So it made me curious to analyze the MRI.
Setting up Opus to do a first review of the MRI
The MRI package was a standard DICOM export containing a few hundred files without extensions, totaling around 266 MB.
For the analysis, I decided to use Opus 4.8 (xhigh) within Claude Code to give it the ability to run code and install packages. Before any work was done, I told it to install any packages that it might need for the analysis. Using Claude Code is especially important to enable it to perform significant amounts of work on this matter. It might seem obvious to coders, but the difference between Claude Code and Claude.ai's chat is enormous, even if those two run the same model.
It was then time to get started. Considering I know nothing about MRIs, I set things up to have Claude work hard on a detailed plan and then take action. The only instruction I gave was "right shoulder pain for 2–3 weeks," which I later realized was less than the human doctors received.
After around an hour, it came back with the report:
The critical problem with that report was that where the doctor saw a Grade III (greater-than-50%) partial-thickness tear at the apical insertion, Opus 4.8 reported an intact tendon!
This was quite disconcerting. I expected the grade to be lower, but that finding was extreme.
To adjudicate, I decided to have Claude do a comparison of the two reports. But this time, I gave it a bit more context; on top of giving the human report, I also provided it with a discussion I had with ChatGPT 5.5 Pro, where I had it give me movements and positions to try as a way of figuring out what my diagnosis was.
From the planning document, here is the approach Opus took:
The approach was careful and methodical, with multiple subagents used as a way of getting new analyses that weren't biased by the existing context.
Again, after around an hour, I got a new report:
Arbiter's verdict: Evidence favours Reader A (moderate-to-high confidence). Mild insertional tendinosis; NO discrete partial- or full-thickness tear identified, including at the apical insertion.
I can't help but find it fascinating that the verdicts are so far from each other. Looking further into the report, I can read that Opus wasn't afraid to say that there are some disputes between the two reports that it can't resolve, and yet this one it could; and very decisively.
There's something incredibly peaceful about being in the hands of an expert you trust. You don't have to worry anymore and can let them guide you through the process.
AI can absolutely shatter that feeling in an uncomfortable way: After having gotten this AI-driven second opinion, the diagnosis and treatment plan look premature and more intervention-heavy than the facts seemed to justify... but I don't know if I can fully trust AI either. So I'm left in a state of limbo where I either try my luck with another doctor or wait and see if my shoulder gets better with the rehab I'm doing.
My hope is that in a couple of model generations, we'll trust AI to review MRIs the way we trust it to proofread our emails.
I am not naming the clinic or doctor because this isn't the point of the article. It's about sharing my technical curiosity about using AI to get second opinions. I may be wrong, or the AI might be wrong. I could also have misunderstood the doctors. So basically, none of this should be taken as medical advice :)
I hope you enjoyed the content of this blog! Feel free to contact me if you have any thoughts, I would love to hear from you!
Made with a custom blog engine built on Ruby on Rails :)
