---
source: "https://ravegraph.beehiiv.com/p/swamp-is-interesting-because-it-doesn-t-trust-ai"
hn_url: "https://news.ycombinator.com/item?id=48505277"
title: "Swamp Is Interesting Because It Doesn't Trust AI"
article_title: "Swamp Is Interesting Because It Doesn't Trust AI"
author: "nickstinemates"
captured_at: "2026-06-12T16:06:48Z"
capture_tool: "hn-digest"
hn_id: 48505277
score: 1
comments: 0
posted_at: "2026-06-12T15:18:29Z"
tags:
  - hacker-news
  - translated
---

# Swamp Is Interesting Because It Doesn't Trust AI

- HN: [48505277](https://news.ycombinator.com/item?id=48505277)
- Source: [ravegraph.beehiiv.com](https://ravegraph.beehiiv.com/p/swamp-is-interesting-because-it-doesn-t-trust-ai)
- Score: 1
- Comments: 0
- Posted: 2026-06-12T15:18:29Z

## Translation

タイトル：SwampはAIを信用していないからこそ面白い

記事本文:
Swamp は AI を信頼していないため興味深い 検索 ログイン サインアップ ホーム
SwampはAIを信用していないからこそ面白い
SwampはAIを信用していないからこそ面白い
AI ツールの世界のほとんどは同じ方向を向いているようです。より多くの自律性、より多くのエージェント、より多くのブラックボックス。
デモは印象的ですが、実際は多くの場合、プロンプト、シェル スクリプト、および交差した指の集合です。非常に多くのことが蓄積され、「ハーネス」に束ねられました
だからこそ私はスワンプをとても面白いと思うのです。
それは最も強力なエージェント フレームワークだからでも、AGI を約束するからでもありません。
エンジニアを置き換えると主張しているからではありません。
実際には、ほぼ逆です。
誰も語らない信頼性の問題
私はここ数年、プラットフォーム エンジニアリングと SRE の分野で働いてきました。
最も信頼性の高いデプロイ パイプラインが最も賢いものであるとは限りません。これは、毎回同じチェックを同じ順序で実行するものです。これは、環境とビルド内のすべてが希望通り、期待通りであることを確認するためにチェックボックスをオンにする必要がある唯一の状況です。
クロードが Bash スクリプトを書くたびに...
私の最近のジョークの 1 つは次のとおりです。
クロードがクソみたいな bash スクリプトを書くたびに、私はこう思います。「これは Swamp にあるべきだ」と。
つまり、私が興味があるロジックは通常、実装ではなく、ワークフローであるということです。
何かが失敗したらどうなるか
それはワークフローの問題です。Swamp はワークフローを第一級の懸念事項として攻撃しているように感じています。
Swamp で最も興味深いのは AI ではありません
Swamp で最も興味深いのは、実際には AI ではないかもしれません。
それは、組織のワークフローを定義して実行する方法を提供する可能性があります。
大きな組織に行けば、同じプロセスを再発明している人々を見つけるでしょう。
チケット→デザイン→PR→レビュー

w → デプロイ
インシデント → 調査 → 検証 → 修正
変更申請→承認→リリース
セキュリティに関する調査結果 → 評価 → 修復
問題は、AI がその作業を実行できるかどうかではありません。
問題は、最終的に作品を説明できるかどうかです。
大規模な組織ではより多くのバッテリーが必要になる場合があります
Patrick Debois 氏は最近、大規模な組織ではおそらくより多くのバッテリーを搭載する必要があると良い点を指摘しました。
しかし、アダム・ジェイコブの反応も重要だったと思います。
最良のワークフローは、誰かが中央で予測するからといって現れるわけではありません。それらは、人々が現実の問題をエッジで解決し、パターンが徐々に共有されるために現れます。
そして、組織の AI ワークフローも同じように起こるのかもしれません。
プラットフォーム エンジニアが注意を払う必要がある理由
プラットフォーム エンジニアは舗装道路の建設に何年も費やしてきました。
Swamp はその考え方の自然な延長のように感じます。組織がインシデントを調査し、変更をレビューし、リリースを検証し、リスクを評価する方法を定義できたらどうなるでしょうか?
誰も読まないドキュメントとしてではありません。
Confluence ページのように誰も更新しません。
私の注意を引いたのは、Swamp がワークフローを実行できるということではありません。
それは、私たちに決定論を注入する場所を与えるということです。
私は必ずしも、より自律的なシステムを望んでいるわけではありません。
重要なチェックを確実に実行するシステムが必要です。
SRE とプラットフォーム エンジニアリングのバックグラウンドを持つ人にとって、それははるかに興味深い問題のように感じられます。
おそらく未来は自律エージェントではないでしょう。
最新のソフトウェア システムの信頼性エンジニアリング - クレーム、証拠、検証を使用して、システムが安全にリリースできることを証明します。

## Original Extract

Swamp Is Interesting Because It Doesn't Trust AI Search Login Sign Up Home
Swamp Is Interesting Because It Doesn't Trust AI
Swamp Is Interesting Because It Doesn't Trust AI
Most of the AI tooling world seems to be heading in the same direction. More autonomy, more agents, more black boxes.
The demos are impressive but the reality is often a collection of prompts, shell scripts and crossed fingers. So much has accreted and been bundled into a 'harness'
Which is why I find Swamp so interesting.
Not because it's the most powerful agent framework or because it promises AGI.
Not because it claims to replace engineers.
Actually, almost the opposite.
The Reliability Problem Nobody Talks About
I've spent the last few years working in platform engineering and SRE.
The most reliable deployment pipeline isn't the smartest one. It's the one that performs the same checks, in the same order, every single time. It's the one situation where you want a box ticking exercise to make sure that everythign in the environment and the build is as you want and expect it to be.
Every Time Claude Writes a Bash Script...
One of my running jokes recently has been:
Every time Claude writes a shitty bash script, I think: this should be in Swamp.
What I mean is that the logic I'm interested in isn't usually the implementation, it's the workflow.
what should happen if something fails
That's a workflow problem - Swamp feels like it's attacking workflows as a first-class concern.
The Most Interesting Thing About Swamp Isn't AI
The most interesting thing about Swamp might not actually be AI.
It might be that it gives us a way to define and execute organisational workflows.
Go into any large organisation and you'll find people reinventing the same processes:
Ticket → Design → PR → Review → Deploy
Incident → Investigation → Validation → Fix
Change Request → Approval → Release
Security Finding → Assessment → Remediation
The question isn't whether AI can perform the work.
The question is whether we can finally describe the work.
Bigger Organisations May Need More Batteries
Patrick Debois made a good point recently that larger organisations will probably need more batteries included.
But I also think Adam Jacob's response was important.
The best workflows don't emerge because somebody predicts them centrally. They emerge because people solve real problems at the edges and patterns gradually become shared.
And it might be how organisational AI workflows happen too.
Why Platform Engineers Should Care
Platform engineers have spent years building paved roads.
Swamp feels like a natural extension of that thinking - what if we could define how an organisation investigates incidents, reviews changes, validates releases or assesses risk?
Not as documentation nobody reads.
Not as Confluence pages nobody updates.
The thing that really caught my attention isn't that Swamp can run workflows.
It's that it gives us a place to inject determinism.
I don't necessarily want a more autonomous system.
I want a system that reliably performs the checks I care about.
For someone coming from an SRE and platform engineering background, that feels like a much more interesting problem.
Maybe the future isn't autonomous agents.
Reliability engineering for modern software systems — using claims, evidence, and validation to prove systems are safe to release.
