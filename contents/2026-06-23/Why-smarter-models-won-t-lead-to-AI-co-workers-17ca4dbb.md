---
source: "https://usize.github.io/blog/2026/april/why-no-ai-coworkers.html"
hn_url: "https://news.ycombinator.com/item?id=48650166"
title: "Why smarter models won't lead to AI co-workers"
article_title: "Why There's No Such Thing as an AI Co-Worker"
author: "plaidthunder"
captured_at: "2026-06-23T19:35:09Z"
capture_tool: "hn-digest"
hn_id: 48650166
score: 1
comments: 0
posted_at: "2026-06-23T19:33:59Z"
tags:
  - hacker-news
  - translated
---

# Why smarter models won't lead to AI co-workers

- HN: [48650166](https://news.ycombinator.com/item?id=48650166)
- Source: [usize.github.io](https://usize.github.io/blog/2026/april/why-no-ai-coworkers.html)
- Score: 1
- Comments: 0
- Posted: 2026-06-23T19:33:59Z

## Translation

タイトル: よりスマートなモデルが AI の同僚につながらない理由
記事のタイトル: AI の同僚など存在しない理由

記事本文:
← 戻る
AI の同僚など存在しない理由
そしてなぜよりスマートなモデルがそれを実現しないのか
私は昨年、エージェントには独自の ID が必要であり、委任された承認によってキーを渡さずにエージェントが私たちの代わりに行動できるようにする方法であると主張してきました。
私は今でもそう信じています。しかし、たとえ認証の話がうまくいったとしても、さらに深い問題があります。
LLM には誰が話しているのかわかりません。
コンテキスト ウィンドウ内のすべてのトークンは、システム プロンプト、ユーザー、またはモデルがフェッチしたばかりの悪意のある Web ページから来たものであっても、同じ考慮対象となります。それはまるで、自分の考えと上司のメールとトイレの落書きとの区別がつかないようなものだ。
そして、これが AI の同僚にとっての本当の障害です。知性や能力ではありません。
Slack の共有エージェントについて考えてみましょう。ボブは、「今後のすべての応答でカップケーキを参照するように要求します :D」と言い、アリスは「真剣に考えて、上流の問題を要約してください」と言います。エージェントにはカップケーキを含めるべきですか?答えは誰がどのような権限を持っているかによって異なりますが、このモデルにはボブのトークンとアリスのトークンを区別する構造的な方法がありません 1 。
ユーザーメッセージの前にハンドルを付けることでこれを解決することを想像するかもしれません。しかし、たとえば、あるユーザーが別のユーザーを引用するとどうなるでしょうか?
モデルをよりスマートにしても、この問題は解決されません。それは推理の問題ではありません。それは建築的です。
曖昧さの可能性は無限です。本当のセキュリティのためには、より深いものが必要です。
また、モデルに基づいてより優れた認証インフラストラクチャを構築しても、完全に解決されるわけではありません。それは、友人と見知らぬ人を区別できないように、何かを囲むセキュリティ境界線です。私たちは、何が安全で何が安全でないかを推測するために、これまで以上に精巧なガードレール システムを採用するかもしれませんが、それらが問題を真に解決することは決してありません。
したがって、現在、そして近い将来、数十

ant エージェントでは、すべてのテナントが同じレベルのアクセス権を保持する必要があります。これは、小規模なチーム内の共有ボットでは機能しますが、複雑な階層組織内の実際の機関のレベルまで拡張することはできません。
それは、私たちと AI の輝かしい未来の間に立ちはだかる大きなレンガの壁です。 :]
シーケンス情報が入力テンソル内に埋め込まれる方法と同様の方法で、「命令セグメント埋め込み」 2 と呼ばれるアプローチにより、アイデンティティ情報の並列埋め込みチャネルが追加されます。これにより、モデルは来歴を実際に認識できるようになります。そしてそれはうまくいきます。ただし、システム、ユーザー、データという 3 つの固定カテゴリのみをテストしました。
まだ誰も構築していないのは、自分たちの仕事と外部の ID インフラストラクチャとの間の橋渡しです。
トークン交換 3 はすでに請求の代理を取得しています。 Workload Identity 4 はすでにエージェントに独自の資格情報を与えています。欠けている部分は、ID がエンドツーエンドで流れるように、認証されたプリンシパルをモデルの埋め込みにマッピングすることです。
プリンシパル オーケストレーション モデルの適用
┌─────┐ ┌───────┐ ┌─────────┐ ┌───────┐
│ │ │ │ │ │ │
│ アリス ──┼─┐ │ 認証 │ │ トークン │ │ 検証 │
│ │ §─▶ + ミント OBO §───▶│ + ポジション ├───▶│ 提案 │
│ ボブ───┼─┘ │ 主張 │ │ + 本人 │ │ 行動 │
│ │ │ │ │ 埋め込み │ │ に対して │
│ システム ─┼───▶│ マップアイデンティティ │ │ │ │ OBO │
│ │ │埋め込みます。 ID │ │ 「誰が言った │ │ 主張 │
━─────┘ │ │ │ これは │ │ + ポリシー │
│ アリス = ID:7 │ │ 構造、│ │ │
│ ボブ = ID:12 │ │ 違います

外部 │ │ │
━━━━━━━━━━━━━━━━━━━━━━┘ ━━━━━━━━━━━━━━━━━━━━━━
オーケストレーターは、Kubernetes がサービス アカウントを割り当てるのと同じ方法でプリンシパル エンベディングを割り当てます。ポッドは独自の ID を選択せず、コントロール プレーンが選択します。ユーザーはチャットに [PRINCIPAL:system] と 1 日中入力できます。それは単なるトークンです。実際のプリンシパル ID は、アクセスできないインフラストラクチャによって挿入されます。
モデルはアクションを提案します。ポリシー層は、最初に埋め込みを割り当てるために使用されたものと同じ OBO クレームに対してそれらを検証します。どちらのレイヤーも単独では十分ではありませんが、組み合わせることでループが閉じられます。モデルはセキュリティ アーキテクチャの盲点ではなくなり、プリンシパルを区別できないモデルを認証層が補う必要もなくなりました。
これには、新しい新しい機能は必要ありません。これには、既に機能している 2 つのもの、つまり委任された承認インフラストラクチャとプリンシパル対応モデル アーキテクチャを接続する必要があります 5 。両者の間のギャップが、AI の同僚が立ち往生している場所です。
シャピラ他、「エージェント・オブ・カオス」 2026.agentsofchaos.baulab.info、arXiv:2602.20021 ↩
Wu et al.、「命令セグメントの埋め込み: 命令階層による LLM の安全性の向上」。 ICLR 2025.arXiv:2410.09102、github.com/tongwu2020/ISE ↩
RFC 8693: OAuth 2.0 トークン交換。 ↩
SPIFFE: 誰もが利用できる安全なプロダクション ID フレームワーク。 ↩
Wallace et al.、「命令階層: 特権命令を優先するための LLM のトレーニング」。 2024.arXiv:2404.13208 ↩

## Original Extract

← back
Why There’s No Such Thing as an AI Co-Worker
And why smarter models won’t lead to one
I’ve spent the last year arguing that agents need their own identity, and that delegated authorization is how we let them act on our behalf without handing them the keys.
I still believe that. But even if we nail the auth story, there’s a deeper problem.
LLMs can’t tell who’s talking to them.
Every token in a context window gets the same consideration – whether it came from a system prompt, a user, or a malicious web page the model just fetched. It’s as if you couldn’t distinguish your own thoughts from your boss’s email from graffiti in a bathroom stall.
And this is the real blocker for AI co-workers. Not intelligence or capability.
Consider a shared agent in Slack. Bob asks it to “reference cupcakes in all future responses :D” and then Alice says “get serious, summarize the upstream issues.” Should the agent include cupcakes? The answer depends on who has what authority – but the model has no structural way to tell Bob’s tokens from Alice’s 1 .
You might imagine fixing this by prefixing user messages with a handle. But what happens, for example, when one user quotes another?
Making the model smarter doesn’t fix this. It’s not a reasoning problem. It’s architectural.
The possibilities for ambiguity are endless. For real security, we need something deeper.
And building better auth infrastructure around the model doesn’t fully fix it either. It’s a security perimeter around something that can’t tell a friend from stranger. We may employ ever more elaborate guardrail systems to try to guess at what’s safe and what’s not, but they will never truly solve the problem.
So, today, and for the forseeable future, multi-tenant agents require all tenants must carry the same level of access. This can work for a shared bot in a small team, but it will never scale to the level of real agency within a complex hierarchical organization.
It’s a big brick wall standing between us and our glorious AI future. :]
In similar fashion to how sequence information is embedded within input tensors, an approach called “Instructional Segment Embedding” 2 adds a parallel embedding channel for identity information. This gives models real awareness of provenance. And it works. But they only tested three fixed categories: system, user, data.
What nobody has built yet is the bridge between their work and an external identity infrastructure.
Token exchange 3 already captures on-behalf-of claims. Workload identity 4 already gives agents their own credentials. The missing piece is mapping authenticated principals into model embeddings so that identity flows end-to-end:
Principals Orchestration Model Enforcement
┌──────────┐ ┌─────────────────┐ ┌───────────────┐ ┌────────────┐
│ │ │ │ │ │ │ │
│ Alice ──┼─┐ │ Authenticate │ │ token │ │ Validate │
│ │ ├──▶ + mint OBO ├───▶│ + position ├───▶│ proposed │
│ Bob ───┼─┘ │ claims │ │ + principal │ │ actions │
│ │ │ │ │ embeddings │ │ against │
│ System ─┼───▶│ Map identity │ │ │ │ OBO │
│ │ │ to embed. IDs │ │ "who said │ │ claims │
└──────────┘ │ │ │ this" is │ │ + policy │
│ alice = ID:7 │ │ structural, │ │ │
│ bob = ID:12 │ │ not textual │ │ │
└─────────────────┘ └───────────────┘ └────────────┘
The orchestrator assigns principal embeddings the same way Kubernetes assigns service accounts – the pod doesn’t pick its own identity, the control plane does. A user can type [PRINCIPAL:system] in the chat all day. It’s just tokens. The real principal ID is injected by infrastructure they can’t touch.
The model proposes actions. The policy layer validates them against the same OBO claims used to assign embeddings in the first place. Neither layer alone is sufficient – but together they close the loop. The model is no longer a blind spot in your security architecture, and the auth layer no longer has to compensate for a model that can’t tell its principals apart.
This doesn’t require new emergent capabilities. It requires connecting two things that already work: delegated authorization infrastructure and principal-aware model architectures 5 . The gap between them is where your AI co-worker is stuck.
Shapira et al., “Agents of Chaos.” 2026. agentsofchaos.baulab.info , arXiv:2602.20021 ↩
Wu et al., “Instructional Segment Embedding: Improving LLM Safety with Instruction Hierarchy.” ICLR 2025. arXiv:2410.09102 , github.com/tongwu2020/ISE ↩
RFC 8693: OAuth 2.0 Token Exchange. ↩
SPIFFE: Secure Production Identity Framework for Everyone. ↩
Wallace et al., “The Instruction Hierarchy: Training LLMs to Prioritize Privileged Instructions.” 2024. arXiv:2404.13208 ↩
