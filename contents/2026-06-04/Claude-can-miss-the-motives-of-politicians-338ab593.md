---
source: "https://futuresearch.ai/llms-miss-motives-politicians/"
hn_url: "https://news.ycombinator.com/item?id=48399278"
title: "Claude can miss the motives of politicians"
article_title: "Claude can miss the motives of politicians"
author: "ddp26"
captured_at: "2026-06-04T16:13:12Z"
capture_tool: "hn-digest"
hn_id: 48399278
score: 7
comments: 0
posted_at: "2026-06-04T14:33:23Z"
tags:
  - hacker-news
  - translated
---

# Claude can miss the motives of politicians

- HN: [48399278](https://news.ycombinator.com/item?id=48399278)
- Source: [futuresearch.ai](https://futuresearch.ai/llms-miss-motives-politicians/)
- Score: 7
- Comments: 0
- Posted: 2026-06-04T14:33:23Z

## Translation

タイトル: クロードは政治家の動機を見逃すことがある
説明: フロンティア AI エージェントは事実を正確にカタログ化しますが、「なぜ今?」と尋ねることはできません。政治的および制度的行動を予測するとき。 BTF-2 ベンチマークの 3 つのケーススタディは、開催国の期限、制度的デフォルト、信用獲得の機会など、主体の構造的インセンティブがいかに失われているかを示しています。
[切り捨てられた]

記事本文:
クロードは政治家の動機を見逃す可能性があります今後の調査 ☰ 解決策
クロードは政治家の動機を見逃すことがある
クロード 4.6 オーパスは人々のインセンティブを微妙に理解できないことがある
LLM チャットボットに、人々と協力したり交渉したりする方法についてアドバイスを求めるときは、説明した状況を考慮して、全員の動機やインセンティブが何であると考えるかを明示的にリストするよう依頼するとよいでしょう。そうすれば、重要な動機が欠けているのを見つけた場合、それを伝えることができ、それが考慮されることがわかります。
そうしないと、LLM またはエージェントが、人間には見分けるのが難しくない動機を見逃してしまう可能性があります。これは私たちの予測評価からの例です。これは、主人公が直面した主な動機に関してクロード Opus 4.6 から完全に外れています。 (この件については、ランエージェントでも 2 回調べました。) 質問は、「ブラジルのカマラ本会議は、2025 年 10 月 15 日から 12 月 31 日までに循環経済法案を承認するでしょうか?」というものでした。 10 月 15 日の観点からクロード氏が迅速に調査したところ、この法案は 1 年以上繰り返し予定されながらも採決されなかったことが判明した。 BTF-2 ベンチマークにおけるクロード オーパス 4.6 の予測エージェントは、12 月 31 日までにそれが起こる可能性を 30% と示しました。
実際、この法案は2週間後の10月29日に可決された。その理由は、ブラジルが11月10日に国連気候サミットであるCOP30を主催していたためだった。ルーラ政権は、目玉となる環境法案を提出することなく、自らの目玉サミットに到達することはできなかった。しかし、「COP30」という単語は、Opuss の 17 件の検索クエリのどれにも、読まれたページにも、中間の推論ブロックにも現れませんでした。
代理人は「この法案の状況はどうなっていますか?」と尋ねました。繰り返し。 「これが突然政府の利益になるような何が起こっているのか？」と尋ねようとは考えなかった。 (A

より優れた予測者がこれを考えたので、これが後から考えずに予想できたことであることがわかります。) 検索および読み取りを行った内容のサブセットを見て、それがどれほど徹底的であったかを確認できますが、依然として重要な動機が欠けています。
これは、Opus 4.6 エージェントの最悪のスコア予測 130 件を調査した論文で見つかった 2 つの主要な戦略的推論の失敗のうちの 1 つです。これに付随する失敗は、レトリックをコミットメントとして読み取ることであり、政治家の発言を文字通りに受け取りすぎることです。これはその逆で、彼らが言っていないことを見逃しています。政治家が直面する重要なインセンティブは、はっきりと語られないこともありますが、人間はかなり確実に推測することができます。
Opus 4.6 のこの失敗モードの例をさらに 2 つ挙げます。
「2025 年度の安全な道路をすべての人に提供する補助金は年末までに発表されますか?」オーパスは、トランプ時代の運輸省が安全補助金をバイデン時代の約10％の速度で処理し、25％の速度で安全補助金を与え、処理の遅さは無能の証拠として扱っていたという擁護報道を見た。本当の動機は逆でした。トランプ寄りのDOTには、年内を延期し、ニュース価値のあるマスを蓄積し、プログラムのブランドを変更し、年末にトランプブランドの発表を1つだけ取り下げる十分な理由があった。 USDOTは12月23日、「トランプ運輸長官、大きくて美しいインフラの構築に10億ドルを投資」という見出しの下、521のコミュニティに9億8,200万ドルを寄付すると発表した。
「司法省は12月31日までにカリフォルニア州のSB 627に対して施行前異議申し立てを提出するだろうか？」カリフォルニア州では、連邦職員がカリフォルニア州での作戦中にマスクを着用することを軽犯罪とする法律が可決され、2026年1月1日から施行された。オーパスは25％を与え、連邦政府が単純に「無視して擁護」する可能性を重視し、職員にマスク着用を継続し法廷で免責を主張するよう指示した。しかし、「無視して防御する」ことは、すべての個人を無視します。

連邦職員は1月1日から軽罪訴追の対象となる。州の刑事法が連邦職員を直接標的とする場合。 1 月 1 日の発効日は単なる目安であり、待つ理由ではありません。
私たちは予測の際にこれを検出し、エージェントは一般に政治問題に関して非常に敏感ですが、この失敗を頻繁に検出したため、他の人が何を達成しようとしているのか、そしてその理由をモデル化する必要があるアドバイスを求めるときに、Opus 4.6 や他のフロンティアエージェントを信頼することに慎重になりました。
これを回避するための 1 つの提案は、AI に人々が達成しようとしていると考えられることをリストするよう依頼することです。そうでないものを見つけた場合は、それを伝えることができるため、それを考慮に入れることができます。
こちらもご覧ください: AI は人々の言葉を信じます (レトリックをコミットメントとして読むことに関する関連投稿)、オーパスはより良い研究を行い、ジェミニはより良い判断力を持ち、エージェントを 2 回実行する、エージェントは時々大惨事になる、Hugging Face に関する BTF-2 データセット、および BTF-2 データセットのリリース
一般的な問い合わせ? hello@futuresearch.ai までご連絡ください。

## Original Extract

Frontier AI agents catalog facts accurately but fail to ask 'why now?' when forecasting political and institutional actions. Three case studies from the BTF-2 benchmark show how missing an actor's structural incentive, like a host-country deadline, an institutional default, or a credit-taking opport
[truncated]

Claude can miss the motives of politicians futu re search ☰ Solutions
Claude can miss the motives of politicians
Claude 4.6 Opus can subtly fail to understand people's incentives
When asking an LLM chatbot for advice in how to collaborate or negoitate with people, it might be a good idea to ask it to explicitly list what it thinks everyone's motivations and incentives are, given the situation you describe to it. That way, if you see a key motivation it misses, you can tell it that, and you'll know it will take that into account.
If you don't, there is a chance the LLM or agent will miss motivations that are not hard for a human to spot. Here's an example from our forecasting evals, a total miss from Claude Opus 4.6 on primary motivation the main actor faced. (I looked into this case also in run agents twice .) The question was: "Will Brazil's Câmara Plenary approve the circular economy bill between October 15 and December 31, 2025?" Claude's quick research, from the perspective of Oct 15, showed the bill had been repeatedly scheduled and not voted on for over a year. A Claude Opus 4.6 forecasting agent in the BTF-2 benchmark gave a 30% chance it would happen by Dec 31.
In fact, the bill passed on October 29, two two weeks later. The reason was that Brazil was hosting COP30, the UN climate summit, on November 10. The Lula government could not arrive at its own showcase summit without shipping its flagship environmental legislation. But the word "COP30" did not appear in any of Opuss's seventeen search queries, in any page it read, or in any intermediate reasoning block.
The agent asked "what is the status of this bill?" repeatedly. It didn't think to ask: "What is going on that might make this very much in the govnerment's interests all of a sudden?" (A better forecaster did think of this, so we know this was something that could be anticipated without hindsight.) You can see a subset of what it did search and read to see how thorough it was, while still missing the critical motivation:
This is one of two dominant strategic-reasoning failures we found in our paper auditing 130 of the Opus 4.6 agent's worst scoring forecasts. The companion failure, reading rhetoric as commitment , is about taking what politicians say too literally. This one is the opposite: missing what they don't say. The key incentive a politician faces is sometimes not stated out loud, but humans can infer them pretty reliably.
Two more examples of this failure mode from Opus 4.6:
"Will FY2025 Safe Streets for All grants be announced by year-end?" Opus saw advocacy reports that the Trump-era DOT was processing safety grants at roughly 10% of Biden-era speed and gave 25%, treating slow processing as evidence of inability. The actual motive was the opposite. A Trump-aligned DOT had every reason to delay through the year, accumulate newsworthy mass, rebrand the program, and drop a single Trump-branded announcement at year-end. USDOT announced $982 million to 521 communities on December 23 under the headline "Trump's Transportation Secretary Invests $1 Billion into Building Big, Beautiful Infrastructure."
"Will DOJ file a pre-enforcement challenge to California's SB 627 by December 31?" California had passed a law making it a misdemeanor for federal officers to wear masks during California operations, effective January 1, 2026. Opus gave 25%, weighted heavily by the possibility that the federal government would simply "ignore and defend": instruct agents to keep wearing masks and assert immunity in court. But "ignore and defend" leaves every individual federal agent exposed to misdemeanor prosecution starting January 1. When a state criminal statute directly targets federal officers. The January 1 effective date was a clock, not a reason to wait.
While we detect this in forecasting, and the agents are in general quite perceptive on political matter, we detected this failure often enough that it makes me wary about trusting Opus 4.6, or other frontier agents, when I'm asking it for advice that requires modeling what other people are trying to accomplish and why.
One suggestion to avoid this is: ask the AI to list what it thinks people are trying to accomplish. If you spot something it doesn't, you can tell it, so it can take it into account.
Also see: AI takes people at their word (the companion post on reading rhetoric as commitment), Opus does better research, Gemini has better judgment , run agents twice , agents sometimes catastrophize , the BTF-2 dataset on Hugging Face , and the BTF-2 dataset release .
General inquiry? You can reach us at hello@futuresearch.ai.
