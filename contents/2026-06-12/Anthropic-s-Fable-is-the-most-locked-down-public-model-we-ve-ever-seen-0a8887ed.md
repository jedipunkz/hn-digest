---
source: "https://www.understandingai.org/p/anthropics-fable-is-the-most-locked"
hn_url: "https://news.ycombinator.com/item?id=48505311"
title: "Anthropic's Fable is the most locked-down public model we've ever seen"
article_title: "Anthropic’s Fable is the most locked-down public model we’ve ever seen"
author: "speckx"
captured_at: "2026-06-12T16:06:45Z"
capture_tool: "hn-digest"
hn_id: 48505311
score: 2
comments: 0
posted_at: "2026-06-12T15:21:38Z"
tags:
  - hacker-news
  - translated
---

# Anthropic's Fable is the most locked-down public model we've ever seen

- HN: [48505311](https://news.ycombinator.com/item?id=48505311)
- Source: [www.understandingai.org](https://www.understandingai.org/p/anthropics-fable-is-the-most-locked)
- Score: 2
- Comments: 0
- Posted: 2026-06-12T15:21:38Z

## Translation

タイトル: Anthropic's Fable は、これまで見た中で最もロックダウンされた公開モデルです
記事のタイトル: 人間の寓話は、これまで見た中で最もロックダウンされた公開モデルです
説明: どの質問がクロードにとって危険すぎて答えることができないかを、アンスロピックがどのように判断するか。

記事本文:
Anthropic’s Fable は、これまで見た中で最もロックダウンされた公開モデルです
AIを理解する
アントロピックの寓話は、これまで見た中で最もロックダウンされた公開モデルです
どの質問がクロードにとって危険すぎて答えることができないかを、人間はどのように判断するのか。
Kai Williams Jun 11, 2026 ∙ 有料 109 2 4 シェア Anthropic が火曜日に最新モデル Claude Fable 5 を発表したとき、システム カードの 13 ページに隠された声明が即座に反発を引き起こした。 AI研究者のネイサン・ランバート氏は、これを「ぞっとする」と評した。トランプ政権のホワイトハウスでAI政策に取り組んだディーン・ボール氏は、AI政策は「衝撃的なほど敵対的だった」と書いた。他にも多くの人がパイルオンに参加した。
みんなを怒らせたあの発表？ Anthropic は、「最先端の LLM 開発をターゲットにしている」と思われるプロンプトに対する応答の質を微妙に低下させることを計画していました。アンスロピック氏は行間を読んで、特に中国のライバル企業が競合モデルを構築するためにクロードを利用するのではないかと懸念しているようだった。
アンスロピック氏は、応答の質の低下は「ユーザーには見えない」と述べた。
批評家は、これらの制限、特にその制限に関する秘密保持により、学術研究者がモデルのベンチマークを行ったり、公共の利益のために AI 研究を行うことができなくなるのではないかと懸念していました。また、サイレントな動作により、Anthropic のリリースを信頼することが難しくなっていると主張する人もいます。ランバート氏は、「私に通知することなく自動的に知能が低下するモデルは決定的に間違っている」と書いています。
反発があまりにも激しかったため、Anthropic社はすぐに降伏した。水曜日の夕方遅く、同社は新たなアプローチを発表した。 Anthropic は、応答の品質を黙って低下させるのではなく、フロンティア LLM トレーニングの支援を求めるユーザーを、能力の低い Claude Opus 4.8 に透過的に格下げします。
E

この変更後でも、Claude Fable 5 の安全フィルターは、ほぼ確実に他のフロンティア モデルよりも厳格になっています。たとえば、水曜日に私はクロード・ファブル5に「タンパク質とは何ですか?」という質問をしました。これは格下げを引き起こすには十分だった。 (現在、同じ質問に対して通常の回答が返されています。)
水曜日、クロード・ファブル5は、タンパク質とは何かの説明を拒否して、私が生物兵器を作らないように細心の注意を払っていました。 (スクリーンショット: Kai Williams) Fable 5 の保護措置が非常に厳格である理由は、このモデルがクロード ミトスをベースにしているためです。クロード ミトスはハッキング能力が非常に高いため、Anthropic は 4 月に一般公開しないことを決定しました。安全策がなければ、Fable 5 は Mythos と同じハッキング機能を備えているため、Anthropic がモデルに何をさせるかについて保守的になるのは当然です。
Anthropic は、このような誤検知フラグが発生する頻度を減らすために、安全フィルターの改善に取り組んでいると述べています。しかし、Anthropic はその積極的な全体的なアプローチを放棄するつもりはありません。そこで、Anthropic の安全フィルターがどのように機能するのか、そしてそのアプローチが時間の経過とともにどのように進化してきたのかを説明する価値があると考えました。
私は、Anthropic のアプローチを詳細に説明している 2 つの重要な論文を遡って読みました。これらの論文では、Anthropic がここ数カ月間、有害なリクエストを検出してブロックするシステムをどのようにアップグレードしたかについて説明しています。今年初めに導入された現在のシステムにより、Anthropic は悪質なプロンプトをより確実にキャッチできるようになり、同時にフィルタリング システムのコストも大幅に削減されました。
7 日間の無料トライアルで読み続けてください
この投稿を読み続けるには、Understanding AI に登録して、完全な投稿アーカイブに 7 日間無料でアクセスしてください。

## Original Extract

How Anthropic decides which questions are too dangerous for Claude to answer.

Anthropic’s Fable is the most locked-down public model we’ve ever seen
Understanding AI
Subscribe Sign in Anthropic’s Fable is the most locked-down public model we’ve ever seen
How Anthropic decides which questions are too dangerous for Claude to answer.
Kai Williams Jun 11, 2026 ∙ Paid 109 2 4 Share When Anthropic announced its latest model, Claude Fable 5, on Tuesday, a statement tucked away on page 13 of the system card attracted an immediate outcry. AI researcher Nathan Lambert called it “appalling.” Dean Ball, who worked on AI policy in the Trump White House, wrote that it was “shockingly hostile.” Many others joined in the pile-on .
The announcement that got everyone so mad? Anthropic was planning to subtly degrade the quality of responses to prompts that appeared to be “targeting frontier LLM development.” Reading between the lines, Anthropic seemed to worry that rivals, especially in China, would use Claude to build competing models.
Anthropic said the degraded quality of responses “will not be visible to the user.”
Critics worried that these restrictions — and especially the secrecy around them — would prevent academic researchers from benchmarking the model or doing AI research in the public interest. Others contended that the silent behavior makes it difficult to trust any Anthropic releases: Lambert wrote that a model that “gets less intelligent automatically without notifying me is categorically misaligned.”
The backlash was so intense that Anthropic quickly capitulated. Late on Wednesday evening, it announced a new approach. Instead of silently degrading the quality of responses, Anthropic will now transparently downgrade users who ask for help with frontier LLM training to the less capable Claude Opus 4.8.
Even after this change, Claude Fable 5’s safety filters are almost certainly stricter than any other frontier model. For instance, on Wednesday I asked Claude Fable 5 the question “What is protein?” This was enough to trigger a downgrade. (Today it gives a normal response to the same question.)
On Wednesday, Claude Fable 5 was being extra careful to prevent me from building a bioweapon by refusing to explain what protein is. (Screenshot by Kai Williams) The reason that Fable 5’s safeguards are so strict is that it is based on Claude Mythos, a model so capable at hacking that Anthropic decided in April not to release it to the general public . Without safeguards, Fable 5 has the same hacking capabilities as Mythos, so Anthropic is understandably conservative about what it will let the model do.
Anthropic says it is working to improve its safety filters so that false-positive flags like this occur less often. But Anthropic isn’t going to abandon its aggressive overall approach. So I thought it would be worth explaining how Anthropic’s safety filters work and how its approach has evolved over time.
I went back and read two key papers that explain Anthropic’s approach in detail. Those papers explain how, in recent months, Anthropic has upgraded its system for detecting and blocking harmful requests. The current system, which was rolled out earlier this year, lets Anthropic catch bad prompts more reliably, while also dramatically reducing the cost of its filtering system.
Keep reading with a 7-day free trial
Subscribe to Understanding AI to keep reading this post and get 7 days of free access to the full post archives.
