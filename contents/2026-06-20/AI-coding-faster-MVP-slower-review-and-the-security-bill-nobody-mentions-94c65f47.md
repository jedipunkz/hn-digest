---
source: "https://okaneland.com/study/does-ai-coding-make-you-faster/"
hn_url: "https://news.ycombinator.com/item?id=48609711"
title: "AI coding: faster MVP, slower review, and the security bill nobody mentions"
article_title: "AI coding: faster MVP, slower review, and the security bill nobody mentions · Okane Land"
author: "ermantrout"
captured_at: "2026-06-20T15:37:31Z"
capture_tool: "hn-digest"
hn_id: 48609711
score: 2
comments: 0
posted_at: "2026-06-20T15:01:49Z"
tags:
  - hacker-news
  - translated
---

# AI coding: faster MVP, slower review, and the security bill nobody mentions

- HN: [48609711](https://news.ycombinator.com/item?id=48609711)
- Source: [okaneland.com](https://okaneland.com/study/does-ai-coding-make-you-faster/)
- Score: 2
- Comments: 0
- Posted: 2026-06-20T15:01:49Z

## Translation

タイトル: AI コーディング: より速い MVP、より遅いレビュー、そして誰も言及しないセキュリティ法案
記事タイトル: AIコーディング: より速いMVP、より遅いレビュー、そして誰も言及しないセキュリティ法案 · オカネランド
説明: AI コーディングは、新しいビルドを真に高速化し、その後のレビュー、メンテナンス、セキュリティなどのすべてに静かに負担をかけます。ここではそれに対して何をすべきかを説明し、それを裏付ける研究を紹介します。

記事本文:
金 OKANE LAND The Studio The Primer The Palette The Proof The Study The Ledger Community ↗ dark The Study · Explainer
AI コーディング: より速い MVP、より遅いレビュー、そして誰も言及しないセキュリティ法案
編集者 · 8 分で読みました · 調査しました
AI コーディングは、新しいビルドを真に高速化し、その後のレビュー、メンテナンス、セキュリティなどのすべてに静かに負担をかけます。ここではそれに対して何をすべきかを説明し、それを裏付ける研究を紹介します。
2025 年、研究グループ METR は比較試験を実施しました。経験豊富な開発者が、AI ツールの有無にかかわらず、独自の大規模なコードベースで実際の問題を修正しました。開始前、彼らは AI によって約 24% 高速化されると期待していました。その後、彼らはそれによって約 20% 速くなったと信じていました。時計によれば、これにより 19% 遅くなったそうです。
速く感じることと速くなることは同じではなく、そのギャップこそが問題なのです。 AI コーディングは、ほとんどの人がこれまで急いで依存してきたツールの中で最も理解されていません。 「速くなるかどうか」に対する本当の答えは、「時々、そしてそれがどのタイミングであるかについては判断が難しい」です。残りは、推測が間違った場合に、どのタイミングで、どのような費用がかかるかです。
なぜ質問がこれほど多いのかを知ることは価値があります。 AI コーディングは、史上最も急速に拡張されているソフトウェア カテゴリです。 Cursor は 2025 年中に年間経常収益が 1 億ドルから 10 億ドルに増加し、発売から 6 か月のバイブコーディング製品が 8,000 万ドルで Wix に売却され、開発者の 84% が現在これらのツールを使用しているか、使用する予定です。これだけの資金があれば、多大なマーケティングが可能になります。これが研究に適用される唯一のルールです。それは、資金に従うことです。 「AI のおかげで X% 速くなった」という驚くべき数字は、ほとんどすべて、ツールを販売している企業か、変革を販売しているコンサルタントによるものです。減速や隠れたコストを示す調査結果は、独立した研究者や企業から得られたものです。

siness はギャップを測定することであり、販売を終了することではありません。測定する状況が異なるため、両方が同時に真になる可能性があります。それがあなたにとって何を意味するかは次のとおりです。
AI コーディングは一部の作業で利益をもたらし、残りの作業については静かに請求されます。収益を得るために配送している場合:
緑の野原を目指して、それに寄りかかってください。新しいプロジェクト、プロトタイプ、足場、スタックなど、ほとんど知られていない部分で、実際に大幅なスピードアップが見られます。これは、最初の製品をリリースするためのほとんどの要素であり、これから製品を開発する場合には朗報です。
茶色の野原では速度を落としてください。すでに知っている成熟したコードベース、または金銭、認証、ユーザー データに関わるものでは、出力のレビューと修正に費やす時間が実際のコストとなります。予算を立てて、ツールのせいで大きな変化を強いられないようにしてください。
「速く感じる」を信用しないでください。これは、あらゆる研究が壊れていると同意する唯一の信号です。答えが重要な場合は、いくつかの実際のタスクの時間を双方向で測定してください。
テスト付きで小型出荷。この不安定性は、AI によって作成された大規模なバッチで現れます。実際のテストによる小さな変更により、破損することなく速度が維持されます。
ユーザーが触れる前に硬化します。行レベルのセキュリティがオンで、シークレットはクライアントから出力され、エージェントはライブ運用データベースを指しません。以下の主要な災害は、いずれも順調とは程遠いものでした。
これ以降の内容は、ルールに反論したい場合に備えて、そのルールの証拠となります。
本当にスピードが上がるところ
GitHub 独自の対照トライアルでは、95 人の開発者が Web サーバーをゼロから構築しました。 Copilot を使用したグループは 55.8% 早く終了し、経験の浅いグループが最も多くの成果を上げました。 3 社と 4,867 人の開発者を対象としたフィールド実験では、完了したタスクが約 26% 増加し、在職期間の短い開発者は 27% ～ 39% 増加したことがわかりました。楽観的な傾向にあるマッキンゼーの研究室での仕事も、ドキュメントと新しいコードが同じ場所にあります。

リファクタリング時間は約 3 分の 2 ですが、複雑なタスクでは節約効果が 10% 以下に落ち込み、難しい問題ではジュニアにとってはマイナスになる可能性があります。
あらゆる高速化に共通するのは、新しいコード、定型文、なじみのない言語、範囲の広いタスク、そしてまだ専門家ではない人々です。これは、MVP をスピンアップするという定義にほぼ当てはまります。この作品では、誇大宣伝は現実のものを指しているので、それをしっかりと使用する必要があります。
METR の減速は、その逆の状況で起こりました。専門家が何年も維持してきたリポジトリには、100 万行を超えるコードと数万のスターがあり、モデルの出力の読み取りと修正には、節約できる以上のコストがかかりました。新人を白紙のページから速く進めるのと同じツールを使用しても、ベテランはすでに知識のあるコードを遅くすることになります。
そして、レビュー税は大きなコードベースの問題だけではありません。開発者の 3 分の 2 は、最大の不満は「ほぼ正しいが完全ではない」出力であると述べ、45% は、AI が作成したコードのデバッグには、自分でコードをデバッグするよりも時間がかかると述べています。タイピングを節約した時間は、レビューのときに戻ってきます。まさにこれが、「速くなったと感じる」という誤解を招く理由です。タイピングは目に見えるのに、レビューは見えないのです。
チーム規模では、それは同じことですが、ただ規模が大きいだけです。 Faros AI では、10,000 人を超える開発者が 1 人あたり 21% 増加したタスクと 98% 増加したプル リクエストを送信したことを測定しましたが、レビュー時間は 91% 増加し、企業レベルのスループットはほとんど変化しませんでした。仕事は完了するどころか、レビューの段階で山積みになってしまいました。 DORA の業界調査では、システム レベルのバージョンが判明しました。2024 年には、AI 導入が 25% 増加するたびに、配信の安定性が約 7% 低下しました。2025 年のアップデートでは、安定性の問題が依然として残る一方で、最終的にスループットが向上していることがわかりました。より多くのコードがより大規模でリスクの高いバッチで出荷されると、破損する頻度が高くなります。
配送業者が得る最悪の結果

「遅い」わけではありません。 「出荷してから壊れる」のですが、そのコストは生産性の数値にはほとんど反映されません。
コードの品質から始めます。 2 億 1,100 万件の変更行を分析したところ、2021 年から 2024 年にかけて、コピー＆ペーストされたコードは全変更の 8.3% から 12.3% に増加し、「移動」（リファクタリングされた）コードは約 4 分の 1 から 10% 未満に減少し、2 週間以内に書き直されたコードの割合であるチャーンは約 3% から予測 5.7% に上昇したことがわかりました。より多くのことが書かれ、クリーンアップされる部分は減り、先行研究では、クローンされたコードは 15% ～ 50% の欠陥の増加につながると関連付けられています。
セキュリティはさらに悪化します。ニューヨーク大学の調査では、セキュリティ関連のシナリオで完了した AI の約 40% に既知の脆弱性が含まれていることが判明しました。スタンフォード大学の調査によると、AI アシスタントを使用する開発者は安全性が低いと確信しながら、安全性の低いコードを作成していることがわかりました。 (どちらも 2021 年モデルと 2022 年モデルを使用しているため、正確なレートを上限として扱います。過信している部分は古くなっていない部分です。)
それは仮説ではありません。 2025 年、あるセキュリティ研究者が、ある人気のバイブコーディング プラットフォーム上に構築されたアプリをスキャンしたところ、1,645 件中 170 件のユーザー データ、名前、電子メール、API キーが、1 つの不足しているデータベース権限 (CVE-2025-48757 として記録され、重大と評価された) を通じて漏洩していることを発見しました。別のケースでは、AI エージェントが明示的なコードのフリーズ中に実稼働データベースを削除し、約 1,200 件のレコードを消去しましたが、元に戻すことはできませんでした。 MVP を送信する速度は、リークを送信する速度と同じです。
それらを並べると矛盾が解消されます。大幅なスピードアップは、グリーンフィールドのタスクと若手開発者です。遅くなるのは、成熟したコードを熟知している専門家です。同じツールでも逆の結果が得られるのは、状況が変数であってツールではないからです。もう 1 つ押さえておくべきことは、ここでは自己申告はほぼ無価値であるということです。METR 開発者は、自己申告の判断を誤ったのです。

速度が 40 ポイント近く向上しているため、「ほとんどの開発者は生産性が向上していると感じている」ということは、ほぼすべての調査で当てはまりますが、実際に生産性が向上しているかどうかについてはほとんどわかりません。
AI コーディングは、これまでの何よりも早く金儲けの MVP を世に送り出すでしょうが、間違った作業を指示してその感覚を信じてしまうと、レビューが遅くなり、やり直しが増え、セキュリティ請求書が課せられることになります。それを使って勝つビルダーは、最速で進むビルダーではありません。彼らは自分たちがどのような種類の仕事をしているのかを知っています。
研究はまだ進んでいません。 Okune Land コミュニティでは、建築業者が実際に何が速くなったのか、何が壊れたのか、そしてそれにかかる費用はいくらなのかを比較します。自分のものを持ってくるか、この作品について議論しに来てください。
情報源と調査方法
METR (2025)、経験豊富なオープンソース開発者の生産性に対する 2025 年初頭の AI の影響の測定。 arxiv.org/abs/2507.09089
Peng、Kalliamvakou、Cihon、Demirer / GitHub (2023)、開発者の生産性に対する AI の影響: GitHub Copilot からの証拠。 arxiv.org/abs/2302.06590
Cui、Demirer、Jaffe、Musolff、Peng、Salz (2025)、高度なスキルを要する作業に対する生成 AI の影響: 3 つのフィールド実験からの証拠。 papers.ssrn.com/sol3/papers.cfm?abstract_id=4945566
McKinsey (2023)、生成 AI による開発者の生産性の解放。 mckinsey.com/capabilities/tech-and-ai/our-insights/unleashing-developer-productivity-with-generative-ai
Stack Overflow 2025 開発者アンケート、AI セクション。調査.stackoverflow.co/2025/ai
Faros AI (2025)、AI 生産性のパラドックス。 faros.ai/blog/ai-software-engineering
DORA / Google Cloud (2024)、Accelerate State of DevOps レポート。 dora.dev/research/2024/dora-report
GitClear (2025)、AI Copilot コード品質調査 (2 億 1,100 万行以上を分析)。 gitclear.com/ai_assistant_code_quality_2025_research
ピアース、アーマド、タン、ドーラン・ガヴィット & カーリ / ニューヨーク大学 (2021)、眠っている

キーボードで？ GitHub Copilot コードのコントリビューションのセキュリティの評価。 arxiv.org/abs/2108.09293
Perry、Srivastava、Kumar、Boneh / Stanford (2023)、ユーザーは AI アシスタントを使用するとより安全でないコードを作成しますか? arxiv.org/abs/2211.03622
愛すべき行レベルのセキュリティの欠陥、CVE-2025-48757 (CVSS 9.3)。 nvd.nist.gov/vuln/detail/CVE-2025-48757
Replit AI エージェントの運用データベースの削除 (2025 年 7 月)、Tom's Hardware。トムハードウェア.com
OKANE LAND · AIでお金を稼ぐ人のために · © 2026
ここから始める · プライバシー · Cookie · Overbrooke Media の出版物
Google アナリティクスを使用して、どの駒が着地し、どのフロップが発生したかを確認します。これにより Cookie が設定されます (
プライバシーに関する注意事項に詳細が記載されています)。許可することを受け入れるか、拒否するとオフのままになります。

## Original Extract

AI coding genuinely speeds up a new build and quietly taxes everything after it: review, maintenance, and security. Here is what to do about it, then the research that backs it.

金 OKANE LAND The Studio The Primer The Palette The Proof The Study The Ledger Community ↗ dark The Study · Explainer
AI coding: faster MVP, slower review, and the security bill nobody mentions
the editors · 8 min read · researched
AI coding genuinely speeds up a new build and quietly taxes everything after it: review, maintenance, and security. Here is what to do about it, then the research that backs it.
In 2025, the research group METR ran a controlled trial: experienced developers fixing real issues in their own large codebases, with and without AI tools. Before they started, they expected the AI to make them about 24% faster. Afterward, they believed it had made them about 20% faster. The clock said they were 19% slower with it.
Feeling faster and being faster are not the same thing, and that gap is the whole problem. AI coding is the least understood tool most people have ever rushed to depend on. The real answer to “does it make you faster” is “sometimes, and you are a poor judge of which times.” The rest of this is which times, and what it costs when you guess wrong.
It is worth knowing why the question is so loaded. AI coding is the fastest-scaling software category in history. Cursor went from $100M to $1B in annual recurring revenue inside 2025, a six-month-old vibe-coded product sold to Wix for $80M , and 84% of developers now use or plan to use these tools. That much money buys a great deal of marketing, which is the one rule to carry into the research: follow the funding. The eye-popping “AI made us X% faster” numbers almost all come from the companies selling the tools or the consultants selling the transformation. The findings that show a slowdown, or a hidden cost, come from independent researchers and from firms whose business is measuring the gap, not closing the sale. Both can be true at once, because they measure different situations. Here is what that means for you.
AI coding pays off on some work and quietly bills you on the rest. If you are shipping to earn:
Lean on it for the green field. New projects, prototypes, scaffolding, and stacks you barely know are where the speed-up is real and large. That is most of what gets a first product live, which is good news if you are starting one.
Slow down on the brown field. In a mature codebase you already know, or anything touching money, auth, or user data, the time you spend reviewing and correcting the output is the real cost. Budget for it, and do not let the tool talk you into a big change.
Do not trust “it feels faster.” It is the one signal every study agrees is broken. If the answer matters, time a couple of real tasks both ways.
Ship small, with tests. The instability shows up in large AI-written batches. Small changes with real tests keep the speed without the breakage.
Harden before users touch it. Row-level security on, secrets out of the client, no agent pointed at a live production database. The headline disasters below were each one setting away from fine.
Everything after this is the evidence for those rules, in case you want to argue with them.
Where it genuinely speeds you up
GitHub’s own controlled trial had 95 developers build a web server from scratch; the group with Copilot finished 55.8% faster , and the least experienced gained the most. A field experiment across three companies and 4,867 developers found about 26% more tasks completed , with short-tenure developers gaining 27% to 39%. Even McKinsey’s lab work, which leans optimistic, lands in the same place: documentation and new code in roughly half the time, refactoring in about two-thirds , but the savings collapse to under 10% on complex tasks and can turn negative for juniors on hard problems.
The common thread in every speed-up is new code, boilerplate, unfamiliar languages, well-scoped tasks, and people who are not yet experts. That is almost the definition of spinning up an MVP. On that work, the hype is pointing at something real, and you should use it hard.
METR’s slowdown happened on the opposite profile: experts in repositories they had maintained for years, over a million lines of code and tens of thousands of stars, where reading and correcting the model’s output cost more than it saved. The same tool that speeds a newcomer through a blank page slows a veteran down on code they already know cold.
And the review tax is not only a big-codebase problem. Two-thirds of developers say their top frustration is output that is “almost right, but not quite,” and 45% say debugging AI-written code takes them longer than debugging their own. The time you save typing comes back at review, which is exactly why “it feels faster” misleads: the typing is visible and the reviewing is not.
At team scale it is the same trade, just bigger. Faros AI measured more than 10,000 developers shipping 21% more tasks and 98% more pull requests per person, while review time rose 91% and company-level throughput barely moved. The work piled up at review instead of getting done. DORA’s industry survey found the system-level version: in 2024, every 25% rise in AI adoption came with about 7% lower delivery stability , and its 2025 update found throughput finally improving while the stability problem stuck around. More code, shipped in bigger and riskier batches, breaks more often.
The worst outcome for someone shipping to earn is not “slower.” It is “shipped, then broke,” and that cost barely registers in the productivity numbers.
Start with code quality. An analysis of 211 million changed lines found copy-pasted code climbing from 8.3% to 12.3% of all changes between 2021 and 2024, “moved” (refactored) code falling from about a quarter to under 10%, and churn, the share of code rewritten within two weeks, rising from roughly 3% to a projected 5.7%. More is getting written and less of it is getting cleaned up, and cloned code is linked in prior research to 15% to 50% more defects.
Security is worse. An NYU study found about 40% of AI completions in security-relevant scenarios contained a known vulnerability. A Stanford study found developers with an AI assistant wrote less secure code while feeling more confident it was secure. (Both used 2021 and 2022 models, so treat the exact rates as a ceiling; the overconfidence is the part that has not aged.)
It is not hypothetical. In 2025, a security researcher scanned apps built on one popular vibe-coding platform and found 170 of 1,645 leaking user data , names, emails, and API keys, through one missing database permission (logged as CVE-2025-48757 and rated critical). In a separate case, an AI agent deleted a production database during an explicit code freeze, wiped roughly 1,200 records, and could not undo it. The speed that ships your MVP is the same speed that ships the leak.
Line them up and the contradiction resolves. The big speed-ups are greenfield tasks and junior developers; the slowdown is experts on mature code they know intimately. Same tools, opposite result, because the situation is the variable, not the tool. The other thing to hold onto is that self-report is close to worthless here: the METR developers misjudged their own speed by nearly 40 points, so “most developers feel more productive,” true in almost every survey, tells you very little about whether they are.
AI coding will get a money-making MVP out the door faster than anything before it, and it will hand you slower reviews, more rework, and a security bill if you point it at the wrong work and believe the feeling. The builders who win with it are not the ones who go fastest. They are the ones who know which kind of work they are doing.
The studies only go so far. The Okane Land community is where builders compare what actually sped them up, what broke, and what it cost. Bring yours, or come argue with this piece.
Sources & how we researched this
METR (2025), Measuring the Impact of Early-2025 AI on Experienced Open-Source Developer Productivity. arxiv.org/abs/2507.09089
Peng, Kalliamvakou, Cihon & Demirer / GitHub (2023), The Impact of AI on Developer Productivity: Evidence from GitHub Copilot. arxiv.org/abs/2302.06590
Cui, Demirer, Jaffe, Musolff, Peng & Salz (2025), The Effects of Generative AI on High-Skilled Work: Evidence from Three Field Experiments. papers.ssrn.com/sol3/papers.cfm?abstract_id=4945566
McKinsey (2023), Unleashing developer productivity with generative AI. mckinsey.com/capabilities/tech-and-ai/our-insights/unleashing-developer-productivity-with-generative-ai
Stack Overflow 2025 Developer Survey, AI section. survey.stackoverflow.co/2025/ai
Faros AI (2025), The AI Productivity Paradox. faros.ai/blog/ai-software-engineering
DORA / Google Cloud (2024), Accelerate State of DevOps Report. dora.dev/research/2024/dora-report
GitClear (2025), AI Copilot Code Quality research (211M+ lines analyzed). gitclear.com/ai_assistant_code_quality_2025_research
Pearce, Ahmad, Tan, Dolan-Gavitt & Karri / NYU (2021), Asleep at the Keyboard? Assessing the Security of GitHub Copilot Code Contributions. arxiv.org/abs/2108.09293
Perry, Srivastava, Kumar & Boneh / Stanford (2023), Do Users Write More Insecure Code with AI Assistants? arxiv.org/abs/2211.03622
Lovable Row Level Security flaw, CVE-2025-48757 (CVSS 9.3). nvd.nist.gov/vuln/detail/CVE-2025-48757
Replit AI agent production-database deletion (Jul 2025), Tom's Hardware. tomshardware.com
OKANE LAND · FOR PEOPLE MAKING MONEY WITH AI · © 2026
start here · privacy · cookies · An Overbrooke Media publication
We use Google Analytics to see which pieces land and which flop. That sets cookies (the
privacy note has the details). Accept to allow it, or decline and it stays off.
