---
source: "https://www.latent.space/p/oai-v-langgraph"
hn_url: "https://news.ycombinator.com/item?id=48550132"
title: "In the Matter of OpenAI vs. LangGraph (2025)"
article_title: "In the Matter of OpenAI vs LangGraph - Latent.Space"
author: "mooreds"
captured_at: "2026-06-16T04:38:28Z"
capture_tool: "hn-digest"
hn_id: 48550132
score: 1
comments: 0
posted_at: "2026-06-16T03:17:18Z"
tags:
  - hacker-news
  - translated
---

# In the Matter of OpenAI vs. LangGraph (2025)

- HN: [48550132](https://news.ycombinator.com/item?id=48550132)
- Source: [www.latent.space](https://www.latent.space/p/oai-v-langgraph)
- Score: 1
- Comments: 0
- Posted: 2026-06-16T03:17:18Z

## Translation

タイトル: OpenAI 対 LangGraph の問題 (2025)
記事のタイトル: OpenAI と LangGraph の関係 - Latent.Space
説明: エージェント エンジニアリングにおける静かな戦争が騒々しくなります。

記事本文:
OpenAI と LangGraph の問題について - Latent.Space
OpenAI と LangGraph の問題について
エージェントエンジニアリングにおける静かな戦争が騒々しくなります。
75 7 6 共有 簡単なお知らせ: AI エンジニア CFP はまもなく終了します。 Computer Use、Voice、Reasoning などの「過小評価されているトラック」を検討し、CFP MCP 経由で応募してください (講演またはワークショップで対応します)。
今日の簡単な投稿に関連して、エージェントの信頼性トラックがあります。また、2025 年の AI エンジニアリングの現状に関する調査にも参加してください。
AI アテンション エコノミーのおかげで、永遠のパフォーマティブ オーガズムの涅槃の状態に存在するハイプボーイの聖職者たちが可能になりました。打ち上げのたびにすべてが変化するたびに精神は常に吹き飛ばされ、「もう終わった」と「もう戻った」というギガヘルツの振動で振動しています。 OpenAI の「Building Agents 実践ガイド」は、そのような最新の地球を揺るがすものです。
ただし、このガイドは Anthropic の同等のガイドほど評判は良くありません。
私たちと一緒に彼の何度も出演しているのを見ていると、ハリソン・チェイスはすぐに「怒る」人ではないので、このガイドを「見当違い」と呼び、一言ずつ分解することは、彼にとって言葉で戦うように見えるかもしれません 1 。
戦いの中心となるのは、ポッドで何度か議論した核心的な緊張です。「ビッグモデルがハンドルを握る」チームと「いや、コードを書かないといけない」チーム（かつてはチェーンと呼ばれていましたが、現在は「ワークフロー」という用語が勝ったようです）。
この議論に対する Harrison の完全な反論を読む必要がありますが、LangGraph 固有の部分を除くと、私にとって最も印象に残った議論は、ワークフロー内のすべての LLM 呼び出しをエージェントに置き換えても、エージェント システムを維持できるというものでした。
そして、理想的なエージェント フレームワークを使用すると、スペクトルの一方から開始してもう一方に移行し、コードを簡単に変更できるように最適化できます。
クラシックなT

すべてのフレームワーク設計者が通過する削減ライン これは、エージェントが多すぎることによる決定を覆したい場合があるため、必要であることがわかります。これは、講演者仲間の Augment Code が #1 の SWE-Bench エントリで見つけたものです。
はっきり言っておきますが、ビッグ モデルの人々がどこから来たのかを理解するのは簡単です。ビッグ ラボで十分に仕事をしている人なら、エンジニアが何百時間もかけて手作業で調整したワークフローが、次の大きなモデルのアップデートで一夜にして消え去るのを見たことがあるでしょう。これは、AI エンジニアが苦い教訓を何度も学ぶのと同じことです。これが、「苦い教訓を念頭に置いた AI エンジニアリング」がサミットで非常に大きな反響を呼んだトピックである理由です (現在、プラットフォーム全体で 124,000 回の視聴数)。
具体的には、今年の OpenAI と Gemini の Deep Research の両方の成功は、主に O3 を活用して研究の計画と実行を通じて推論し、その後、Bolt と Manus AI がワークフロー エンジニアリングをほとんど行わずにクロードと同じことを行ったことにより、ワークフローの「誘導バイアス」制約なしでモデルを単純に拡張する汎用エージェントの構築について、語るべきことがたくさんあることを証明したと思います。 O チームの研究者、Hyung Won Chung 氏は、構造を追加すると短期的には効果が得られますが、その構造はモデル (事前学習または推論計算) がスケールアップし続けるにつれてパフォーマンスが低下する傾向があると指摘しました。
この話から。私たちはこの洞察を文脈から少し外して使用しています。ヒョンウォンは内部モデルのアーキテクチャについて発言していましたが、それはモデルを中心に構築された外部システムにも当てはまると考えています。彼がその推定を支持するかどうか疑問に思う人もいるでしょう。あなたの目標が AGI を構築すること、水平プラットフォームを構築すること、特にモデル セレクターがあることさえ混乱する非技術的な消費者をターゲットにしたプラットフォームを構築することである場合、その立場をとるのは理解できます。

(データセット/人間によるフィードバック収集の目的で、奨励することさえあります)。
結局のところ、ハリソンが「実際には」戦う姿勢をとっていないと私が主張する理由は、彼がスペクトルが存在する余地を残しており、両方を行うためのオプションが必要なだけであるという、（ゲーム内で明らかな肌を持つ人にとっては）驚くほどバランスの取れた議論をしているからです。
これについては議論するのが難しいと思います。意味のある会話をするとしたら、実際には、このパレートフロンティアの現在の状態が今日実際にどこにあるのか（まだ凸状であるかどうかはわかりません）、そしてそれをどのように解消するかということが重要です。
本当のところは、ワークフローを作成する際に避けるべき悪いアイデアなど、間違いなく失敗するものがあるということです。また、その逆に、基礎となるモデルがアップグレードされても価値を維持するワークフロー システムも存在します。これは、昨年 1 月に AlphaCodium が最初にリリースされ、o1 リリース後の 11 月にその価値が「配布されなくなった」状態で見られたように、フロー エンジニアリングをカバーするポッドで説明したとおりです。
Harrison が記事の中で行ったもう 1 つの非常に優れた点は、関連するすべてのエージェント フレームワークの完全な比較表を今日公開したことです。ただし、もちろん彼ですら McCormick の罠から逃れることはできませんでした。全員のエージェント定義の記述的なメタフレームワークを、配布外の新しいエージェント定義と照合してテストすると便利です。
影響: インテント、メモリ、プランニング、認証、制御フロー、ツール。オーケストレーションのない一部のフレームワークを主観的に除外しました。オーケストレーションが優れたエージェント制御フローにとって非常に重要であると考えているためです。これは、目の肥えたエージェント エンジニアにとって、抽象化と機能の非常に公平な買い物リストだと思います。また、エージェント フレームワークが世界を約束しているにもかかわらず、いくつかのことを簡単に実行できない場合に、特定のギャップを感じる理由も明確に示しています。

イントロの話に戻りますが、心が常に混乱していると、それを取り戻すことはできません。人々の決断を助けることは、コミュニティにとって価値のあるサービスだと思います。
この種の議論から学ぶのが好きなら、私たちは昨年の NeurIPS でのディラン対フランクル対決の成功を倍増させ、また、関連する業界の議論の両側からの誠実な討論者である「グレート ディベート」と呼ばれるものへの投稿も受け付けています。誰もが勝ちますが、最もよく考えを変えることができる人が最も勝ちます。ペアで応募してください！
https://sessionize.com/ai-engineer-worlds-fair-2025
私が主張したいのは、実際にはそうではないということです。ハリソンは常に外交官です。
75 7 6 シェア 前へ 次へ この投稿に関するディスカッション コメント Restacks Dexter Awoyemi 2025 年 4 月 21 日 編集済み Latent.Space が「いいね！」しました。
LangChain 専門用語の「チェーン」という用語は、個々のノードとワークフロー全体の両方を指すようで、私には決して明確ではありませんでした。そのため、「ワークフロー」に置き換えられたことを嬉しく思います。 LangGraph はこれまでで最高のものです。
Anthropic の Building Effects Agents の記事では、エージェント ループと決定論的ワークフローの区別が他のどの考えよりも明確になっていますが、区別が過度に単純化されています。この空間に建築しているほとんどの人（私も含めて）は、それをより主体性のスペクトルとして見ているような印象を受けます。
私は、このエージェント フレームワークの内訳が進化するのを見てみたいと思っています。これは、最終的に非常に役立つリソースになる可能性があるものへの素晴らしいスタートです。
返信 シェア Sven Meyer 2025 年 4 月 21 日 編集「Harrison 氏は自身の記事で、関連するすべてのエージェント フレームワークの完全な比較表を公開しました」
ありがとうございます、とても興味深いです!!
返信 シェア Latent.Space らによる 2 件の返信 その他のコメント 5 件... トップ 最新のディスカッション 投稿はありません

## Original Extract

The silent war in Agent Engineering gets loud.

In the Matter of OpenAI vs LangGraph - Latent.Space
Subscribe Sign in In the Matter of OpenAI vs LangGraph
The silent war in Agent Engineering gets loud.
75 7 6 Share Quick reminder: AI Engineer CFPs close soon ! Take a look at “undervalued tracks” like Computer Use, Voice, and Reasoning , and apply via our CFP MCP (talks OR workshops, we’ll figure it out).
Relevant to today’s quick post we do have an Agent Reliability track. Also: take the 2025 State of AI Engineering Survey !
The AI attention economy has enabled a hypeboi priesthood who exist in a state of perpetual performative orgasmic nirvana, minds continually blown as every launch Changes Everything, vibing at gigahertz oscillations of “it’s so over” vs “so back”. OpenAI’s “ Practical Guide to Building Agents ” is the latest such earth shatterer:
This guide, however, has been less well received than Anthropic’s equivalent.
If you watch his multiple appearances with us, Harrison Chase is not someone who is quick to “anger”, so calling this guide “ misguided ” and doing a word by word teardown can seem like fighting words for him 1 .
At the heart of the battle is a core tension we’ve discussed several times on the pod - team “Big Model take the wheel” vs team “nooooo we need to write code” (what used to be called chains, now it seems the term “workflows” has won).
You should read Harrison’s full rebuttal for the argument, but minus the LangGraph specific parts, the argument that stood out best to me was that you can replace every LLM call in a workflow with an agent and still have an agentic system:
And the ideal agent framework lets you start from one side of the spectrum and move to the other, optimizing for making code easy to change :
the classic tradeoff line that every framework designer walks You’ll find this necessary because sometimes you DO want to reverse decisions from having too many agents - as fellow speaker Augment Code found in their #1 SWE-Bench entry :
To be clear it’s easy to understand where the Big Model folks are coming from: if you work with Big Lab enough, you’ve seen hundreds of engineer-hours of hand tuned workflows obliterated overnight with the next big model update — the AI Engineer equivalent of learning the Bitter Lesson again and again. This is why “ AI engineering with the Bitter Lesson in mind ” was such a resonant topic at the Summit (now at 124k views across platforms):
Specifically, I think the success of both OpenAI and Gemini’s Deep Research this year primarily leveraging O3 to reason through research planning and execution, and later Bolt and Manus AI doing the same with Claude, with very little workflow engineering, has demonstrated that there’s a lot to be said for building general purpose agents that simply augment models without the “inductive bias” constraints of workflows. O-team researcher Hyung Won Chung noted that adding more structure gets you wins in the short term, but that structure tends to lose in performance as the model (pretrain or inference compute) keeps scaling up.
from this talk . we are using this insight slightly out of context: Hyung Won was making statements about INTERNAL model architecture, but we think it also applies about EXTERNAL systems built around the model — one wonders if he’d also endorse that extrapolation. If your goal is to build AGI, to build a horizontal platform, particularly one targeting non-technical consumers who are confused by even having a model selector, then it’s an understandable position to take, and (even encourage, for the purposes of dataset/human feedback collection).
Ultimately the reason I argue Harrison isn’t -actually- taking a fighting stance is he leaves room for the spectrum to exist and makes a remarkably (for someone with obvious skin in the game) balanced argument that you’re going to just want options for doing both:
I find this hard to debate - if meaningful conversation is to be had, it really is more about where the current state of this Pareto frontier really is today (I’m not sure it is convex yet) and how to move it out.
What -is- true is that there is such a thing as bad ideas to avoid in creating workflows that will DEFINITELY get steamrolled, and also the converse - workflow systems that maintain value as their underlying models get upgrades - as we saw last year with AlphaCodium’s initial release in Jan and then its value persisting “out of distribution” in Nov after the release of o1 - as we discuss on our pod covering Flow Engineering .
The other pretty cool thing that Harrison did in his piece was publish a full comparison table of all relevant Agent Frameworks today , although of course even he couldn’t escape the McCormick trap . It’s useful to test our descriptive metaframework of everybody’s Agent definitions against a new out-of-distribution Agents definition:
IMPACT: Intent, Memory, Planning, Auth, Control flow, Tools. We have subjectively filtered out some frameworks without orchestration, because we view orchestration as very critical for good agent control flow I think this is a remarkably fair shopping list of abstractions and features for the discerning Agent Engineer — it also articulates why you feel certain gaps when an Agent Framework promises you the world and yet you can’t do some things easily.
To callback to our intro, if your mind is continually blown, it can never be made up. I think that helping people make up their minds is a valuable service to the community.
If you like learning from this kind of debate, we’re doubling down on the success of the Dylan v Frankle showdown from last year’s NeurIPS, and also accepting submissions for what we’re calling “The Great Debates” - good faith debaters from two sides of a relevant industry debate. Everybody wins, but the people who are best able to change minds win the most. Apply in pairs !
https://sessionize.com/ai-engineer-worlds-fair-2025
As I’ll argue: they’re actually not! Harrison is ever the diplomat.
75 7 6 Share Previous Next Discussion about this post Comments Restacks Dexter Awoyemi Apr 21, 2025 Edited Liked by Latent.Space Thanks for this!
The terminology of 'chain' in LangChain lingo has never been clear to me since it seems to refer to both individual nodes and the entire workflow, so I'm happy it has been superseded by 'workflow'. LangGraph is their best to date imo.
Anthropic's Building Effective Agents article makes the distinction between agentic loops and deterministic workflows clearer than any other thoughtpiece, even though it oversimplifies the distinction. I get the impression that most people building in this space (myself included) see it more as a spectrum of agency.
I'm keen to see this Agent Framework Breakdown evolve. It's a great start for what could ultimately become a very useful resource.
Reply Share Sven Meyer Apr 21, 2025 Edited "Harrison did in his piece was publish a full comparison table of all relevant Agent Frameworks"
Thanks for that , very interesting !!
Reply Share 2 replies by Latent.Space and others 5 more comments... Top Latest Discussions No posts
