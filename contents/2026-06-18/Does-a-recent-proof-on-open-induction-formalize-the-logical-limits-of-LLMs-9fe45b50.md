---
source: "https://arxiv.org/abs/2606.17275"
hn_url: "https://news.ycombinator.com/item?id=48588207"
title: "Does a recent proof on open induction formalize the logical limits of LLMs?"
article_title: "[2606.17275] Syntactic Systems Cannot See Semantic Invariants"
author: "yaccb3"
captured_at: "2026-06-18T18:56:14Z"
capture_tool: "hn-digest"
hn_id: 48588207
score: 1
comments: 0
posted_at: "2026-06-18T16:54:56Z"
tags:
  - hacker-news
  - translated
---

# Does a recent proof on open induction formalize the logical limits of LLMs?

- HN: [48588207](https://news.ycombinator.com/item?id=48588207)
- Source: [arxiv.org](https://arxiv.org/abs/2606.17275)
- Score: 1
- Comments: 0
- Posted: 2026-06-18T16:54:56Z

## Translation

タイトル: 開放帰納法に関する最近の証明は、LLM の論理的限界を形式化していますか?
記事のタイトル: [2606.17275] 構文システムは意味上の不変条件を認識できない
説明: arXiv 論文 2606.17275 の要約ページ: 構文システムは意味的不変条件を認識できない

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2606.17275
ヘルプ |高度な検索
コンピューター サイエンス > コンピューター サイエンスの論理
[2026 年 6 月 15 日に提出]
タイトル: 構文システムは意味上の不変条件を認識できない
要約: 私たちは小さな未解決の質問から始めます。Hetzl と Vierling は、帰納法の 2 つの理論、開帰納法と節集合サイクルは比較できないものであるかどうかを尋ねました。彼らは一方の方向性を証明し、もう一方の方向性を残しました。ここで締めくくりますが、証明は恥ずかしいほど短いものです。加算のルールは最初の引数が $0$ または後続引数である場合にのみ実行でき、Skolem 定数はそのどちらでもないため、項 $a{+}b$ と $b{+}a$ には決して触れることができず、これらに決して触れることができないマシンはそれらが等しいことを決して証明できません。 2 つの理論を分けるのは 2 つの定数の順序であり、その順序は記号に関するものではなく、数値に関する事実です。この証明から、そのような引数の形状に名前を付ける小さな一般原則である構文不変原則を抽出します。次に、$\mathsf{P}$ と $\mathsf{NP}$ を解決するための既知の障壁において、この同じ形状が非公式にどのように現れるかについて、いくつかの推測的なコメントで終わります。各障壁は、その障壁内の技術が到達できない記述のレベルを示していると思われます。この類推は現実のものであるため、私たちはこれを定理ではなく提案として取り上げますが、それを擁護できる段階を超えてそれを押し進めることはしません。その過程で、$\SAT$ の高速アルゴリズムが存在する場合、常に書き留めることができるマシンとして展示可能であるかどうか、あるいはそれが可能であるかどうかという、たとえ話が示唆しているものの決着がつかない未解決の疑問を提起します。

場合によっては、数値の関数としてのみ見つかることがあります。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI (登録保留中)
関連するDOI:
https://doi.org/10.5281/zenodo.20618697
もっと学ぶために集中する
関連リソースにリンクする DOI
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。
arXiv に連絡する arXiv に連絡するにはここをクリックしてください
お問い合わせ
arXiv メールを購読する 購読するにはここをクリックしてください
購読する

## Original Extract

Abstract page for arXiv paper 2606.17275: Syntactic Systems Cannot See Semantic Invariants

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2606.17275
Help | Advanced Search
Computer Science > Logic in Computer Science
[Submitted on 15 Jun 2026]
Title: Syntactic Systems Cannot See Semantic Invariants
Abstract: We start from a small open question, where Hetzl and Vierling asked whether two theories of induction, open induction and clause set cycles, are incomparable. They proved one direction and left the other open. Here we close it, and the proof is almost embarrassingly short, because the rules for addition can only fire when the first argument is $0$ or a successor, a Skolem constant is neither, so the terms $a{+}b$ and $b{+}a$ can never be touched, and a machine that can never touch them can never prove they are equal. The thing that separates the two theories is the order of two constants, and that order is a fact about numbers, not about symbols. We extract from this proof a small general principle, the Syntactic Invariance Principle, that names the shape of such arguments. We then close with a few speculative remarks on how this same shape appears, informally, in the known barriers to settling $\mathsf{P}$ versus $\mathsf{NP}$, where each barrier seems to point to a level of description that the techniques in the barrier cannot reach. We raise this as a suggestion rather than a theorem, since the analogy is real but we do not push it past the point where we can defend it. Along the way we raise an open question that the analogy suggests but does not settle, on whether a fast algorithm for $\SAT$, were it to exist, would always be exhibitable as a machine you can write down or whether it could be found, in some cases, only as a function on the numbers.
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Related DOI :
https://doi.org/10.5281/zenodo.20618697
Focus to learn more
DOI(s) linking to related resources
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
contact arXiv Click here to contact arXiv
Contact
subscribe to arXiv mailings Click here to subscribe
Subscribe
