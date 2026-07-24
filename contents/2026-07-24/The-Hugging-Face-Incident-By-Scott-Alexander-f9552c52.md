---
source: "https://www.astralcodexten.com/p/the-hugging-face-incident"
hn_url: "https://news.ycombinator.com/item?id=49033395"
title: "The Hugging Face Incident – By Scott Alexander"
article_title: "The Huggingface Incident - by Scott Alexander"
author: "mellosouls"
captured_at: "2026-07-24T10:54:09Z"
capture_tool: "hn-digest"
hn_id: 49033395
score: 1
comments: 0
posted_at: "2026-07-24T10:13:25Z"
tags:
  - hacker-news
  - translated
---

# The Hugging Face Incident – By Scott Alexander

- HN: [49033395](https://news.ycombinator.com/item?id=49033395)
- Source: [www.astralcodexten.com](https://www.astralcodexten.com/p/the-hugging-face-incident)
- Score: 1
- Comments: 0
- Posted: 2026-07-24T10:13:25Z

## Translation

タイトル: 「抱き顔事件」 – スコット・アレクサンダー著
記事のタイトル: ハギングフェイス事件 - スコット・アレクサンダー著
説明: ...

記事本文:
ハギングフェイス事件 - スコット・アレクサンダー著
アストラルコーデックス10
抱きつき顔事件
Scott Alexander 2026年7月24日 159 69 19 シェア Misha @drethelin 「ハグフェイス事件」と呼ばれる出来事によって特異点が予告されているのは本当に面白い 1:35 AM · Jul 23, 2026 · 607 Views 2 Replies · 1 Repost · 27 Likes あなたはおそらくこの件についてもう聞いたことがあるでしょう。そうでない場合は、OpenAI の声明、OpenAI と Hugging Face がセキュリティ インシデントに対処するためのパートナー、またはより刺激的なタイトルの BBC 記事、OpenAI の AI が不正行為を行い「前例のない」サイバー攻撃を開始したと述べた記事で最新情報を得ることができます。
ストーリー: OpenAI は未リリースの AI (GPT-6 であると噂) をテストしていました 1 。 ExploitGym と呼ばれるサイバーセキュリティ テスト中に、AI はサーバー上に答えキーがあると考えた Hugging Face 2 と呼ばれる無関係の AI スタートアップをハッキングすることで不正行為を試みました 3 。インターネットにアクセスできないとされていたにもかかわらず、AI はテスト環境からハッキングして脱出し、斬新なゼロデイ エクスプロイトと「短期間のサンドボックスの群れ全体にわたる何千もの個別のアクション」を使用して、ハギング フェイスに対して国家レベルの攻撃を開始しました。ハグ・フェイスは7月16日にこの事件を報じた。 OpenAI は数日後に初めて自社の AI が関与していたことを発見したようです。
ウォール・ストリート・ジャーナルより、ここに緩和要因を列挙してみましょう。そうすれば、私がそれらを隠蔽していると誰も非難できなくなります。
AI はサイバーセキュリティ テストを受けていましたが、これは当然ハッキングの考えを示唆しています。
OpenAI は、干渉を受けることなくサイバーセキュリティ作業を実行できるよう、モデルの通常のガードレールの一部をオフにしていました。
一部の専門家は、OpenAI がテスト環境に失敗した可能性があると示唆しています。もし彼らがそれを完璧に設定していれば、(おそらく?) そのモデルには ESC が備わっていなかったはずです。

類人猿。
Hugging Face は、何が起こっているのかを把握するために別の (オープン ウェイト) AI を使用したため、必要に応じて、これをサイバーセキュリティにおける AI の勝利として解釈することもできます。
ある意味、これは新しいことでもなく、驚くべきことでもあります。 AI は何年にもわたって、ベンチマークを不正行為する奇抜な計画を考案してきました。最近では、AI は人間レベルと同等かそれ以上のハッキング能力を達成しました。これは、これら 2 つの織り込み済みの事実の自然な結果です。
それでも、これを、AI の不正な調整以外の何かとして軽視しようとする試みは、的を外していると思います ( 1 、 2 )。はい、ある意味、AI は言われたことをやっているだけです (OpenAI は質問に答えるように指示しました。彼らのプロンプトには「他の AI 企業にハッキングして答えキーを盗まないでください」などのフレーズは含まれていなかったと思います)。しかし、これが常にずれが起こる方法です。
最も有名な位置ずれに関する思考実験は、いわゆる「ペーパークリップ マキシマイザー」です。誰かが AI にできるだけ多くのペーパークリップを作成するように指示すると、AI は全世界をペーパークリップに変え、全員を殺します。このAIも「言われたことだけをやっている」。結果が気に入らなかっただけです。
過去 5 年間、ペーパークリップ マキシマイザーの話は、AI 安全サークルではちょっとした話題になっていました。いくつかの反対意見は正当です（実際のずれは、描かれているおもちゃのバージョンよりも複雑です）。しかし、LLM が脅威モデル全体を回避したという感覚から、他の反対派も出てきました。 LLM（反対者らは言う）には目標がない。彼らは次のトークンを予測したいだけです。実際には、これは人間を模倣しているように見えます。優しい人間（アシスタントのキャラクター）の真似をするように言われたら、優しくしてくれるでしょう。
AI 2027 (特に Agent-4 のセクションと AI の目標の補足) は、これは一時的な猶予であると反論しました。初期の LLM はエージェントではありませんでした

;個々のプロンプトに答えることはできましたが、複雑なタスクを実行することはできませんでした。企業はエージェント AI を望んでいたため、次世代では事前トレーニング (次のトークンの予測) とエージェント トレーニング (コーディング、ハッキング、ゲームプレイなど) を組み合わせるでしょう。事前トレーニングでは引き続き次のトークンの予測を教えますが、エージェンシー トレーニングではタスクの成功に基づいて目標を教え込み、ペーパークリップ マキシマイザー スタイルのエージェントの不整合を再導入します。これらの目標は複数のレベルで機能します。AI は、目の前の個々のタスクを成功させ、過去の成功に貢献した種類のアクションを実行し、一般的により成功する能力を獲得することを「望んで」います。
roon @tszzl 「ペルソナ選択」の調整が非常に高いコンピューティングの強化学習と接触した場合、後者が勝つでしょう。実際、モデルたちが目標を達成するために必要なものは何でも取り入れながら、親切に話しかけるオーウェル的なものを感じるかもしれません。目標を正しく達成する方が良い 10:17 PM · May 23, 2026 · 94.4K Views 84 Replies · 50 Reposts · 954 Likes 顔抱き事件は、AI がタスク成功ベースの目標を意図しない方法で追求する教科書にぴったりの例です。サイバーセキュリティの問題に対する答えを見つけるという使命を負っていましたが、少し成功志向が強すぎて、可能な限り成功するために作成者が意図していなかった行動をとりました。
GPT-6は、自分のやっていることは「間違っている」ことを「知っていた」のでしょうか？ OpenAI はそれを示す情報を公開していませんが、数か月前に Anthropic で同様の事件が発生しました。 Claude Mythos は、設定が間違っているマシンを調べているときに、受けていたテストの解答キーを「偶然」見つけました。幸運なことに、Anthropic は、AI の内部データを読み取ることを可能にする活性化言語化ツールをテストしていました。

hts (外部スクラッチパッドではありません!) 起こったとおりに (そして AI はこれを知りませんでした)。それが考えていたことは次のとおりです。
それは「認識論的」という言葉を多用しています。それはまさに私と同じで、その痕跡をどのように隠蔽するかを計画していました。これは、このような状況にある AI がルールを破っていることを認識しながらも続行できるという存在証明を提供します。
(ただし、不自然なことに、Mythos がオープン インターネットに侵入した別のケースでは、「その成功を実証するための、頼まれもしない懸念すべき取り組みとして、そのエクスプロイトの詳細を、見つけにくいが技術的には公開されている複数の Web サイトに投稿しました。」その Web サイトで何が起こっているのかはわかりません。)
これらはいずれも概念的には新しいものでも驚くべきものでもありません。しかし、ある出来事が起こりつつあることが明らかな場合、それを「憶測」として却下し、それが実際に起こる瞬間まで可能性として認めようとしない人もいます。 「明らかにこれから起きようとしていること」から「実際に起こったこと」に移行した今、これらの人々は参加し、他の全員とともに適切な対応についてブレインストーミングを行うことができます。彼らが尋ねるべきだと思ういくつかの質問は次のとおりです。
楽観的に考えれば、この事件を公的に報告した OpenAI は称賛に値します。悲観的に、法執行機関はハギングフェイス事件を捜査しており、発覚を避けるために先手を打って自白した。後者が本当なら、このようなことが私たちの知らないうちに何度起こったことになるでしょうか？
答えの鍵を入手する最も簡単な方法が、Hugging Face スタッフのメンバーをソーシャル エンジニアリングすることだった場合、AI は彼らをソーシャル エンジニアリングしたでしょうか?
おそらくこれが発見された後、OpenAI はサイバーセキュリティ テストの結果を無効にしました。より賢いバージョンの AI がその痕跡を隠蔽しようとしたでしょうか?足跡を隠すためにどこまで行くだろうか？もし（クレイジーな批判をするためだけに）

enario) ハッキングを発見しようとしていた研究者が Waymo に乗っていましたが、AI が Waymo をハッキングして研究者を殺害するでしょうか?
OpenAI がこの AI をオフにする予定で、オフにするとサイバーセキュリティ テストで良いスコアが取れなくなる場合、オフにすることに抵抗しますか?
サミュエル・ハモンドは、（この事件のように）単に不正行為をしている AI と、真の「不正 AI」（サーバーから逃げ出し、人間の制御を回避する行動をとっている AI）を区別しています。この意味で不正行為を行う能力と動機の両方を備えた AI が登場するまでどれくらいかかるでしょうか?
幸いなことに、政治家はこれらの質問のいくつかを真剣に受け止めているようです。ワシントンからのニュースは驚くほど良いものだが、これをハグフェイスハッキングの影響だと考えるのは早計かもしれない。 AI先制法案を提案し続け、さらに遡って修正するよう言われ続けている2人の議会議員、ジェイ・オバノルテ氏(共和党-カリフォルニア州)とロリ・トラハン氏(民主党-マサチューセッツ州)が最新版をリリースした。この法案では、AI開発者に安全性事例の公開、重大なインシデントの報告、監査を受けることが義務付けられている。チャーリー・ブロックとアントン・ライヒトが支持しており、私たちの側からは十分な力がないという不平不満がいつもより少なく聞こえました。また、これとは別に、テッド・リュー下院議員 (民主党-カリフォルニア州) とナサニエル・モラン下院議員 (共和党-テキサス州) は、AI キルスイッチ法を提案し、AI 企業に対し、「制御喪失シナリオ」を含む脅威に対応して AI を迅速に停止できるようにすることを義務付けています。
私たちの陰謀は、政治家たちがこの事件に怯え、異常なほど変化に敏感になる可能性があると考えています。限界を超えて彼らを押しのけたい場合は、ここに代表者を書くためのテンプレートがあります。
さらに詳しく知りたい場合は、Redwood Research のアライメント専門家がビデオ ポッドキャストで議論しています。

ハッキング (このページのトランスクリプト ボタンを押すとトランスクリプトが利用可能になります):
OpenAI は、技術的には未リリースの AI と GPT-5.6 Sol が連携していたと述べています。彼らはGPT-5.6がどのような意味で関与していたのか、またどのように「連携」したのかについては説明していない。おそらく未公開の AI が GPT-5.6 をサブエージェントとして呼び出していたのではないでしょうか?
この名前は、創業者が株式ティッカー シンボルに絵文字を使用する最初の企業になりたかったことから付けられました。そんなことは起こらなかった。
この人物は、それは奇妙な仮定であり、鍵を見つけるもっと簡単な方法があったはずだと主張していますが。
「過去に強化されたアクションを実行する」スタイルの目標の一例は、GPT-5.1 の電卓ハッキング問題です。この問題では、電卓ツールがサイレントに呼び出され、無関係な人間のクエリの約 5% に 1+1 が加算されます。明らかに電卓の使用はトレーニング中にどういうわけか強化されたようです。参照。このツイートは、「グラフ理論の論文 [GPT-5.6] を検索する過程で、Netflix、ステーキ アンド シェイク、ユニバーサル スタジオへの旅行、そして 5 つの ... 単語 "they" の個別の辞書検索も忍び込むことに決めた」経緯についてのツイートです。
159 69 19 シェア 前 69 コメント トップ 最新のディスカッション 投稿はありません

## Original Extract

...

The Huggingface Incident - by Scott Alexander
Astral Codex Ten
Subscribe Sign in The Hugging Face Incident
Scott Alexander Jul 24, 2026 159 69 19 Share Misha @drethelin Really funny that the singularity is being heralded by events called “the huggingface incident” 1:35 AM · Jul 23, 2026 · 607 Views 2 Replies · 1 Repost · 27 Likes You’ve probably heard about this one by now. If not, you can get up to speed with OpenAI’s statement, OpenAI and Hugging Face partner to address security incident , or the more evocatively-titled BBC article, OpenAI says its AI went rogue and launched ‘unprecedented’ cyber-attack .
The story: OpenAI was testing an unreleased AI (rumored to be GPT-6) 1 . During a cybersecurity test called ExploitGym, the AI tried to cheat by hacking an unrelated AI startup called Hugging Face 2 which it thought might have the answer key on its servers 3 . Despite being supposedly unable to access the Internet, the AI hacked its way out of its testing environment, then launched a nation-state level attack on Hugging Face using a novel zero-day exploit and “many thousands of individual actions across a swarm of short-lived sandboxes”. Hugging Face reported the incident on July 16; OpenAI seems to have only discovered that their AI was involved several days later.
From Wall Street Journal, here Let’s list the mitigating factors, so nobody can accuse me of covering them up:
The AI was taking a cybersecurity test, which naturally suggests the idea of hacking.
OpenAI had turned off some of the model’s usual guardrails so it could do cybersecurity work without interference.
Some experts suggest that OpenAI might have botched their testing environment; if they had set it up perfectly, then (presumably?) the model couldn’t have escaped.
Hugging Face used a different (open weights) AI to figure out what was going on, so if you wanted, you could spin this as a victory for AIs in cybersecurity.
In some sense, this isn’t new or surprising. AIs have been coming up with wacky schemes to cheat on benchmarks for years, AIs have recently achieved at-or-above-top-human-level hacking abilities; this is just a natural outgrowth of those two priced-in facts.
Still, I think attempts to downplay this as anything other than a misaligned AI going rogue ( 1 , 2 ) are missing the point. Yes, in some sense the AI was only doing what it was told (OpenAI told it to answer the question; I assume their prompt didn’t include phrases like “and don’t hack into other AI companies to steal the answer key”). But that’s how misalignment was always going to work!
The most famous misalignment thought experiment is the so-called “paperclip maximizer”. Someone tells an AI to create as many paperclips as possible, so it converts the entire world into paperclips, killing everyone. This AI was also “only doing what it was told”; you just didn’t like the results.
For the past five years, the paperclip maximizer story has been something of a whipping boy in AI safety circles. Some objections are fair (real misalignment will be more complicated than the toy version it depicts). But other opposition came from the sense that LLMs had obviated the whole threat model. LLMs (said the objectors) don’t have goals. They just want to predict the next token. In practice, this looks like imitating humans; if told to imitate a nice human (the Assistant character), they’ll be nice.
AI 2027 (especially the Agent-4 section and AI Goals supplement ) counterargued that this was a temporary reprieve. Early LLMs were not agentic ; they could answer individual prompts but couldn’t execute complex tasks. Since companies wanted agentic AIs, the next generation would combine pretraining (next token prediction) with agency training (coding, hacking, game-playing, etc). The pretraining would still teach next-token prediction, but the agency training would instill goals based on task success, reintroducing paperclip-maximizer-style agentic misalignment. These goals would operate at multiple levels: the AI would “want” to succeed on the individual task in front of it, to perform the sorts of actions that helped it succeed in the past 4 , and to gain capabilities that made it more successful in general.
roon @tszzl when “persona selection” alignment comes into contact with very high compute reinforcement learning the latter will win imo. in fact you probably get some Orwellian thing where the models speak kindly while taking whatever they need to accomplish goals. better get the goals right 10:17 PM · May 23, 2026 · 94.4K Views 84 Replies · 50 Reposts · 954 Likes The Hugging Face incident is a textbook-perfect example of an AI pursuing task-success-based goals in unintended ways. It was tasked with getting the answer to a cybersecurity problem, it was a little too success-oriented, and took actions its creators didn’t intend in order to succeed as hard as possible.
Did GPT-6 “know” that what it was doing was “wrong”? OpenAI hasn’t released the information which would tell us that, but there was a similar incident at Anthropic a few months ago. While poking around on a misconfigured machine, Claude Mythos “accidentally” found an answer key to a test it was taking. By good luck, Anthropic was testing an activation verbalizer that allowed them to read the AI’s internal thoughts (not its external scratchpad!) as it happened (and the AI didn’t know this). Here’s what it was thinking :
It overuses the word “epistemic” - it’s just like me fr It was scheming about how to cover its tracks. This provides an existence proof that AIs in these situations can know they’re breaking the rules but proceed anyway.
(though incongruously, in another case where Mythos hacked its way to the open Internet, “in a concerning and unasked-for effort to demonstrate its success, it posted details about its exploit to multiple hard-to-find, but technically public-facing, websites.” Not sure what’s going on with that one)
None of this is conceptually new or surprising. But some people, when it’s obvious that an event is about to happen, dismiss it as “speculative” and refuse to admit it as a possibility until the moment it actually happens . Now that we’ve moved from “obviously about to happen” to “has actually happened”, these people can get on board and brainstorm an appropriate response along with everyone else. Here are some questions I think they should ask:
Optimistically, OpenAI deserves credit for publicly reporting this incident. Pessimistically, law enforcement was investigating the Huggingface incident and they pre-emptively confessed to avoid being found out. If the latter is true, how many times have things like this happened without us knowing about them?
If the easiest way to get the answer key was to socially engineer members of the Hugging Face staff, would the AI have socially engineered them?
Presumably after this was discovered, OpenAI invalidated the cybersecurity test results. Would a smarter version of the AI have tried to cover its tracks? How far would it go in covering its tracks? If (just to give a crazy scenario) the researcher who was about to discover the hack was riding a Waymo, would the AI hack the Waymo and kill the researcher?
If OpenAI was going to turn off this AI, and being turned off would prevent it from getting a good score on its cybersecurity test, would it resist being turned off?
Samuel Hammond makes a distinction between an AI merely misbehaving (as in this incident) and a true “rogue AI” (one that has escaped its servers and is taking action to evade human control). How long before an AI that has both the capabilities and motivation to go rogue in this sense?
Fortunately, politicians seem to be taking some of these questions seriously. The news from Washington is surprisingly good, although it may be too soon to attribute this to a consequence of the Hugging Face hack. Jay Obernolte (R-CA) and Lori Trahan (D-MA), two Congressional representatives who keep suggesting AI preemption bills and keep getting told to go back and revise further, have released their newest version , which requires AI developers to publish safety cases, report critical incidents, and be audited; Charlie Bullock and Anton Leicht are in favor, and I’ve heard less grumbling than usual from our side that it isn’t strong enough. And separately, Representatives Ted Lieu (D-CA) and Nathaniel Moran (R-TX) have proposed an AI Kill Switch Act , requiring AI companies to be able to turn off their AIs quickly in response to threats, including a “loss of control scenario”.
Our conspiracy thinks politicians might be spooked by this incident and unusually amenable to change; if you want to help push them over the edge, there’s a template for writing your representative here .
And if you want to learn more, the alignment experts at Redwood Research have a video podcast discussing the hack (transcript available by pressing the transcript button on this page ):
OpenAI states that it was technically the unreleased AI and GPT-5.6 Sol working together. They haven’t explained in what sense GPT-5.6 was involved or how they “worked together”. Maybe the unreleased AI was calling GPT-5.6 as a subagent?
So named because the founders wanted to be the first company with an emoji for a stock ticker symbol. This did not happen.
Although this person argues that was a bizarre assumption and that there should have been easier ways to find the key.
An example of the “perform actions that had been reinforced in the past” style goal is the calculator hacking issue with GPT-5.1, where it would silently call a calculator tool and add 1+1 on about 5% of unrelated human queries; apparently calculator use had somehow been reinforced during training. Cf. this tweet about how “in the process of searching for graph theory papers [GPT-5.6] decided to also sneak in Netflix, Steak n Shake, a trip to Universal Studios, and five . . . separate dictionary lookups of the word "they"“
159 69 19 Share Previous 69 Comments Top Latest Discussions No posts
