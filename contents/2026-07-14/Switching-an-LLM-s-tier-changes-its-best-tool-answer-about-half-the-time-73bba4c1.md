---
source: "https://modelsagree.com/labs/model-tiers"
hn_url: "https://news.ycombinator.com/item?id=48905935"
title: "Switching an LLM's tier changes its \"best tool\" answer about half the time"
article_title: "The Cheap Tier Disagrees With the Expensive Tier"
author: "arephan"
captured_at: "2026-07-14T13:08:43Z"
capture_tool: "hn-digest"
hn_id: 48905935
score: 2
comments: 0
posted_at: "2026-07-14T12:42:21Z"
tags:
  - hacker-news
  - translated
---

# Switching an LLM's tier changes its "best tool" answer about half the time

- HN: [48905935](https://news.ycombinator.com/item?id=48905935)
- Source: [modelsagree.com](https://modelsagree.com/labs/model-tiers)
- Score: 2
- Comments: 0
- Posted: 2026-07-14T12:42:21Z

## Translation

タイトル: LLM の層を切り替えると、その「最適なツール」の答えが約半分の確率で変わります
記事のタイトル: 安価な階層は高価な階層と一致しない
説明: Haiku vs Opus、Flash vs Pro、Fast vs Expert: 同じ AI の低予算層では、約半分の確率で別の「最適なツール」が挙げられます。また、自己の好みによって、ファミリーごとに方向が反転します。

記事本文:
モデルも同意します。コム/ラボ
実験02
モデル層 · 自己優先 · 1 つのプロンプト、9 つのモデル
安価な層は高価な層と一致しません
ChatGPT、Claude、Gemini、Grok のすべての層に、Haiku 対 Opus、Flash 対 Pro、Fast 対 Expert という同じ 10 の「最高の AI ツール」に関する質問をしました。すべての層から同じ答えが得られた質問は 1 つもありませんでした。
AI モデルを比較する人は皆、ChatGPT 対 Claude 対 Gemini 対 Grok のようにブランドを比較します。これが、modelsagree.com で私たちが毎日行っていることです。何百もの「ベスト X」の質問について 4 つすべてを再投票し、同意した場合は公開します。しかし、ブランド名の下には、より地味な変数が隠されています。それは、「階層ピッカー」です。クロードは Haiku、Sonnet、Opus として出荷されます。 Gemini は Flash および Pro として出荷されます。 Grok は Fast および Expert として出荷されます。ほとんどのユーザーはデフォルトには触れません。そこで私たちは、ブランドを一定に保ち、階層のみを変更した場合、「最良のベクター データベース」は変わらないのかと尋ねました。
AI ツールのリーダーボードから 10 のライブ カテゴリ (ベクター データベース、AI コーディング アシスタント、LLM 可観測性、RAG フレームワーク、GPU クラウド、テキスト読み上げ API、LLM ゲートウェイ、エージェント フレームワーク、評価ツール、フロンティア API プロバイダー) を取得し、同じランキング プロンプト (サイトで使用しているものと同じもので、逐語的に理由付きで上位 5 位) を、通常の加入者としてアクセスできる各ファミリーのすべての層に送信しました。
ChatGPT: GPT-5.5 → GPT-5.6 (OpenAI はコンシューマー アカウントのミニ階層をブロックします。そのため、GPT の場合、これはサイズではなく世代を測定します)
Grok: Fast → Expert (grok.com 独自の推論層)
同じ日に、同じ言葉遣いで、階層ごとに 1 つの回答が得られます。次に、各ファミリーのレベルを相互に比較しました。
もう一度読んでください。単一のブランド内で、階層を切り替えると、約半分の確率で 1 位の推奨事項が変更されます。そしてその後ろにあるトップ5のリスト

#1 は、同じファミリーの層間で約 50 ～ 65% (Jaccard) だけ重複します。
9 段階すべてにおいて、満場一致で勝者が得られたのは 10 問中 0 問でした。
彼らは自分たちの裏庭で誰に栄冠を授けるでしょうか？
最も明らかな唯一の質問は、最も自己言及的な質問でした。それは、開発者がどのフロンティア モデル API に基づいて構築する必要があるかということです。各層は、ライバルと比較して独自のメーカーをランク付けする必要がありました。すべての答えは次のとおりです。
ここには明確な話はありません - そしてそれが発見です。民間理論では、「モデルはそのメーカーに金がかかる」と言われています。階層データによると、自己優先度は階層に依存し、方向が一貫していないことがわかります。
クロードは大きくなるにつれて謙虚になっていきます。俳句とソネットは人類の頂点に立つ。フラッグシップである Opus は、その栄冠を OpenAI に渡します。
双子座は大きくなるにつれて誇りを持っていきます。 Flash は OpenAI を冠します。 Pro が Google の栄冠に輝きます。兄弟層がライバルを選ぶ中、グリッド全体で自社のメーカーを選んだ唯一の層です。
ChatGPT は世代間で反転しました。 GPT-5.5 は OpenAI を冠します。 GPT-5.6は人間性を冠します。
Grok は決して自らを冠することはありません。ファストとエキスパートの両方がアンスロピックに王座を譲ります。アンスロピックは、この問題について層間で意見が一致している唯一のファミリーであり、ライバルについても意見が一致しています。
1 つの層をテストすることで、これらのファミリーの 1 つをバイアスについて監査していた場合、その層がたまたま保持した結論をそのまま使用して終了することになります。
予算レベルでは予算ツールを購入します
私たちが探してもいなかった 2 番目のパターンが浮上しました。 「トレーニングに最適な GPU クラウド」については、Sonnet、Opus、GPT-5.5、GPT-5.6、Gemini Pro、および両方の Grok モードのすべての主力層が CoreWeave で述べています。グリッド内で最も安価な 2 つの層、Claude Haiku と Gemini Flash はどちらも、Lambda が価値のある価格の代替品であると述べています。 「最高のベクター データベース」では、Gemini Flash は、無料の Postgres 拡張機能である pgvector の Pinecone (大手企業のお気に入り) をスキップしました。
それは仮説です、ではありません

法則 — Grok の Expert モードでも pgvector が選択されているため、無料ツールの本能は安価な層に限定されたものではありません。しかし、GPU とクラウドの分割は明確です。すべての予算階層が予算クラウドを選択し、すべての主力企業が主力クラウドを選択しました。それがトレーニングデータの偏りであれ、質問の浅い解釈であれ、あるいは好みのようなものであれ、安価なモデルは安価なスタックにパターンマッチングされました。
ソフトウェアを販売する場合にこれが重要となる理由
AI による回答は、実際の獲得チャネルになりつつあります。私たちは、それが私たち自身のトラフィックで起こるのを観察しています。業界は「ChatGPT が私たちについて何を言っているか」を監視し始めています。かつてGoogleのランクを監視していた方法と同じだ。しかし、その監視のほとんどすべてが主力 API 層に影響を及ぼします。
無料ユーザーは安価な利用枠を利用できます。安価な層では異なる答えが得られます。 「AI の可視性」は 1 つの数値ではありません。
Gemini Pro があなたを推奨し、Flash が競合他社を推奨する場合、無料利用枠 (デフォルトの利用枠) の何百万ものユーザーが競合他社の名前を聞くことになります。私たちの 10 問のサンプルでは、​​LangSmith (可観測性を重視してプロが選んだ) 対 Langfuse (Flash が選んだ)、Pinecone 対 pgvector、CoreWeave 対 Lambda、LiteLLM 対 Portkey でまさにその分かれ目が起こりました。将来の顧客が聞く推奨事項は、顧客が決して開かないドロップダウンに依存します。
これが、modelsagree.com が単一のモデルの意見を信頼するのではなく、4 つのファミリーにわたる合意を追跡している理由でもあります。つまり、意見の相違がシグナルです。 4 つのファミリーすべて (そして今ではそのすべての層) がまだ何かに同意している場合 (イレブンラボは 9 層中 7 層でテキスト読み上げで第 1 位を保持し、Braintrust は評価で 9 層中 6 層で、LangSmith は可観測性で 9 層中 6 層で 1 位を獲得しました)、その合意を却下するのははるかに困難です。
10 問すべての各段階の第 1 位は以下の表にあります。各階層を記載した完全なトップ 5 リスト

推論はオープン データセット ( consensus.json 、CC BY 4.0) にあります。これらのカテゴリのライブ 4 モデル リーダーボードは、modelsagree.com/best にあります。
質問ごとにレベルごとに 1 つのサンプル。模範解答はリロール間で変動します。上記の不一致の一部はサンプリング ノイズであり、安定した層の特性ではありません。 (当社のメイン リーダーボードは、まさにこの理由から継続的に再投票を行います。) 方向を実際のものとして扱い、正確なカウントを指標として扱います。
階層はどのファミリーでも同じではありません。クロード層とジェミニ層はモデル サイズが異なります。 GPT は世代が異なります (消費者アカウントはミニを選択できません)。 Grok は、同じ基礎モデル上の推論努力モードです。比較は「加入者が実際に選択できるもの」であり、制御されたパラメータスイープではありません。
Consumer surfaces、2026 年 7 月 12 日。人々が実際に使用する製品 (有料サブスクリプションの CLI/Web アプリ) を通じて、同じ日に同じプロンプトですべての質問が行われました。サーフェス (生の API、システム プロンプト) が異なれば、応答も異なる場合があります。
プロンプトを公開します。これは、サイト上のすべてのリーダーボードの背後にあるものと同じです — 方法論を参照してください。ベンダー名を指定するカテゴリの質問はありません。
modelsagree.com labs · 層バイアス実験 · ChatGPT · Claude · Gemini · Grok
modelsagree.com は、ChatGPT、Claude、Gemini、Grok が何百ものソフトウェア カテゴリにわたって第 1 位にランクされたものを追跡し、継続的に再投票し、すべてのモデルの推論を逐語的に公開しています。データ: CC BY 4.0 — modelsagree.com として引用。

## Original Extract

Haiku vs Opus, Flash vs Pro, Fast vs Expert: the budget tier of the same AI names a different “best tool” about half the time — and the self-preference flips direction per family.

modelsagree . com / labs
experiment 02
Model tiers · self-preference · one prompt, nine models
The cheap tier disagrees with the expensive tier
We asked every tier of ChatGPT, Claude, Gemini and Grok the same ten “best AI tool” questions — Haiku against Opus, Flash against Pro, Fast against Expert. Not one question got the same answer from every tier.
Everyone comparing AI models compares brands : ChatGPT versus Claude versus Gemini versus Grok. That's what we do every day at modelsagree.com — re-polling all four on hundreds of “best X” questions and publishing where they agree. But there's a quieter variable hiding under the brand name: the tier picker . Claude ships as Haiku, Sonnet and Opus. Gemini ships as Flash and Pro. Grok ships as Fast and Expert. Most users never touch the default. So we asked: if you hold the brand constant and change only the tier, does “the best vector database” stay the same?
We took ten live categories from our AI-tooling leaderboards — vector databases, AI coding assistants, LLM observability, RAG frameworks, GPU clouds, text-to-speech APIs, LLM gateways, agent frameworks, eval tools, and frontier API providers — and sent the identical ranking prompt (the same one our site uses, verbatim, top-5 with reasons) to every tier of each family we could access as a normal subscriber:
ChatGPT: GPT-5.5 → GPT-5.6 (OpenAI blocks the mini tiers on consumer accounts — so for GPT this measures generations, not sizes)
Grok: Fast → Expert (grok.com's own reasoning tiers)
Same day, same wording, one answer per tier. Then we compared each family's tiers against each other.
Read that again: within a single brand, switching the tier changes the #1 recommendation roughly half the time . And the top-5 lists behind those #1s only overlap by about 50–65% (Jaccard) between tiers of the same family.
Across all nine tiers, zero of the ten questions produced a unanimous winner.
Who do they crown in their own backyard?
The single most revealing question was the most self-referential one: which frontier-model API should a developer build on? Each tier had to rank its own maker against its rivals. Here's every answer:
There is no clean story here — and that's the finding. The folk theory says “models shill for their maker.” The tier data says self-preference is tier-dependent and inconsistent in direction :
Claude gets humbler as it gets bigger. Haiku and Sonnet crown Anthropic; Opus — the flagship — hands the crown to OpenAI.
Gemini gets prouder as it gets bigger. Flash crowns OpenAI; Pro crowns Google — the only tier in the whole grid that picked its own maker while its sibling tier picked a rival.
ChatGPT flipped between generations. GPT-5.5 crowns OpenAI; GPT-5.6 crowns Anthropic.
Grok never crowns itself. Both Fast and Expert hand the crown to Anthropic — the only family whose tiers agree on this question, and they agree on a rival.
If you were auditing one of these families for bias by testing one tier, you'd walk away with whichever conclusion that tier happened to hold.
The budget tier buys budget tools
A second pattern surfaced that we didn't go looking for. On “best GPU cloud for training,” every flagship tier — Sonnet, Opus, GPT-5.5, GPT-5.6, Gemini Pro, and both Grok modes — said CoreWeave . The two cheapest tiers in the grid, Claude Haiku and Gemini Flash, both said Lambda — the value-priced alternative. On “best vector database,” Gemini Flash skipped Pinecone (the big tiers' favorite) for pgvector , the free Postgres extension.
It's a hypothesis, not a law — Grok's Expert mode also picked pgvector, so the free-tool instinct isn't exclusive to cheap tiers. But the GPU-cloud split is clean: every budget tier picked the budget cloud, and every flagship picked the flagship cloud. Whether that's training-data skew, a shallower read of the question, or something like taste — the cheap models pattern-matched to the cheap stack.
Why this matters if you sell software
AI answers are becoming a real acquisition channel — we watch it happen in our own traffic. The industry is starting to monitor “what does ChatGPT say about us?” the way it once monitored Google rank. But almost all of that monitoring hits the flagship API tier.
Free users get the cheap tier. The cheap tier gives different answers. Your “AI visibility” is not one number.
If Gemini Pro recommends you but Flash recommends your competitor, then the millions of users on the free tier — the default tier — are hearing your competitor's name. In our ten-question sample, that exact split happened to LangSmith (Pro's pick for observability) versus Langfuse (Flash's pick), to Pinecone versus pgvector, to CoreWeave versus Lambda, to LiteLLM versus Portkey. The recommendation your future customers hear depends on a dropdown they never open.
That's also why modelsagree.com tracks the consensus across four families rather than trusting any single model's opinion: the disagreement is the signal. When all four families — and now, all their tiers — still agree on something (ElevenLabs held #1 in text-to-speech on seven of nine tiers; Braintrust on six of nine in evals; LangSmith on six of nine in observability), that consensus is far harder to dismiss.
Every tier's #1 for all ten questions is in the table below; the full top-5 lists with each tier's stated reasoning are in our open dataset ( consensus.json , CC BY 4.0). The live four-model leaderboards these categories come from are at modelsagree.com/best .
One sample per tier per question. Model answers wobble between re-rolls; some of the disagreement above is sampling noise, not stable tier personality. (Our main leaderboards re-poll continuously for exactly this reason.) Treat the direction as real and the exact counts as indicative.
Tiers aren't the same thing across families. Claude and Gemini tiers are different model sizes; GPT's are different generations (consumer accounts can't select the minis); Grok's are reasoning-effort modes on the same underlying model. The comparison is “what a subscriber can actually pick,” not a controlled parameter sweep.
Consumer surfaces, July 12, 2026. Everything was asked through the products people actually use (CLI/web apps on paid subscriptions), same day, identical prompt. Different surfaces (raw API, system prompts) may answer differently.
We publish the prompt. It's the same one behind every leaderboard on the site — see methodology . No category questions name any vendor.
modelsagree.com labs · tier-bias experiment · ChatGPT · Claude · Gemini · Grok
modelsagree.com tracks what ChatGPT, Claude, Gemini & Grok each rank #1 across hundreds of software categories, re-polled continuously, with every model's reasoning published verbatim. Data: CC BY 4.0 — cite as modelsagree.com.
