---
source: "https://rootly.com/blog/what-doom-taught-us-about-ai-assisted-incident-response"
hn_url: "https://news.ycombinator.com/item?id=48940764"
title: "What Doom taught us about AI-assisted incident response"
article_title: "What Doom taught us about AI-assisted incident response | Rootly"
author: "sylvainkalache"
captured_at: "2026-07-16T21:56:08Z"
capture_tool: "hn-digest"
hn_id: 48940764
score: 1
comments: 0
posted_at: "2026-07-16T21:51:49Z"
tags:
  - hacker-news
  - translated
---

# What Doom taught us about AI-assisted incident response

- HN: [48940764](https://news.ycombinator.com/item?id=48940764)
- Source: [rootly.com](https://rootly.com/blog/what-doom-taught-us-about-ai-assisted-incident-response)
- Score: 1
- Comments: 0
- Posted: 2026-07-16T21:51:49Z

## Translation

タイトル: AI 支援のインシデント対応について Doom が教えてくれたこと
記事のタイトル: AI 支援のインシデント対応について Doom が教えてくれたこと |根深い
説明: 私たちは、AI エージェントが MCP を介して Doom で戦うオープンソースのベンチマークである Doom Agent Arena を構築しました。 AI を活用したインシデント対応について私たちが学んだことは次のとおりです。

記事本文:
What Doom taught us about AI-assisted incident response |根深い
PNGをダウンロード
すべてのアセットをダウンロード PagerDuty から切り替える ログイン
L デモを予約する D 製品
製品
オンコール
インシデント対応
AI SRE
回顧展
ステータスページ
プラットフォーム
カタログ
統合
オンコールヘルス
変更履歴
Jul 16, 2026 The on-call widget, 24-hour time, and a new Alerts table.
サイズ
エンタープライズ
スタートアップ
私たちと比較してください
ポケベルデューティ
オプスジェニー
JSMチーム
エンジニアリング
セキュリティ
How Canary built incident management from scratch on Rootly while its engineering team quadrupled. Replit brings structure, speed, and better decision-making to incident response with Rootly. Caribou が Rootly を使用して年間 200 以上のエンジニアリング時間を節約した方法。 Capa が Incident.io を Rootly に置き換えた理由How Reactiv outgrew incident.io and moved to Rootly for real AI and on-call without the upsell. How SpotDraft replaced manual follow-ups with synced, measurable incident management on Rootly. How GRAIL replaced a manual, ad hoc incident process with Rootly and cut manual effort by 80%. How Cora cut time-to-acknowledge on critical alerts to <10 minutes with Rootly. CTSO Central が Incident.io ではなく Rootly を選択した理由。 Poll Everywhere switched from PagerDuty, cut on-call tooling costs and made response faster with Rootly. Momentum が FireHydrant を選択しなかった理由と、Rootly とのオンコール エンジニアの数をどのように最小限に抑えているか。 KnowBe4 switched from PagerDuty and turned Incident Response into a repeatable, predictable system with Rootly. Achievers switched from PagerDuty and productized reliability on Microsoft Teams with Rootly. How Motive migrated from Opsgenie and achieved 99.99% reliability with Rootly. Clay switched off Opsgenie to use Rootly to streamline incident management. Webflow shutdown PagerDuty to use Rootly to build a proactive reliabilit

文化。 Wealthsimple が Rootly を使用して、インシデント指揮官の健康と心理的安全性の文化をどのように構築しているか。 Upstart が Rootly を使用して、自社とともに成長するインシデント対応システムを構築する方法。 ROLLER がスリムな SRE チームを維持しながらグローバルにスケールする方法。 Lucidworks が Rootly を使用して、自社の独自の製品向けにオーダーメイドのインシデント管理を作成する方法。顧客のリソース
エンジニアリング
ドキュメント
APIリファレンス
システムステータス
ヘルプとサポート ガイド
AI SRE および自動 RCA
オンコールオペレーション
インシデント対応
インシデントコミュニケーション
非難のないレット
[切り捨てられた]
ブログに戻る
Doom が AI 支援のインシデント対応について教えてくれたこと
目次 例 H2 Doom Agent Arena の紹介。これは、AI エージェントが MCP 経由で Doom プレイヤーを制御し、複数のラウンドにわたって互いに戦うオープンソースのリアルタイム ゲーム環境ベンチマークです。
Doom Agent Arena を構築することで、私たちの目標は、調査範囲を拡大して、LLM がインシデント対応者を支援して可能な限り迅速かつ効果的に機能停止を解決する方法をより深く理解することでした。このベンチマークはゲーム設定に組み込まれていますが、動的な環境の解釈、結果についての推論、失敗した計画からの回復、時間やコストの制約下での適応など、信頼性に関連する LLM スキルを対象としています。
構築方法、方法論、結果について詳しく見ていきましょう。
Doom Agent Arenaの構築方法
他のほとんどの Doom ベンチマークは、ゲーム フレームをビジョン モデルにフィードします。しかし、私たちの目標は LLM の推論能力をテストすることであったため、ベンチマークでは代わりにエージェントが MCP サーバー経由でマップやリソースを含む構造化された JSON として Doom ゲームの状態を観察し、Doom が実行するための高レベルの計画を送信できるようにしています。
遅延はエンジニアリング上の主要な課題の 1 つでした。 Doom はミリ秒単位で意思決定を行う人間のために作られていますが、

モデルの往復にははるかに長い時間がかかるため、この 2 つのバランスを取る必要がありました。ツール自体の呼び出しは高速ですが、モデルの思考には決定ごとに最大 8 秒かかります。そのため、モデルをリアルタイムで再生することはできませんでした。そこで私たちは混合アプローチを採用しました。モデルはプレイヤーが何をすべきか、どこに行くべきかなどの高レベルの決定を行い、Doom のエンジンは実行 (移動、照準、射撃) を処理します。
このゲームは、ヘルスパックやショットガンなどのリソースが用意された壁に囲まれたアリーナでの 1 対 1 のドゥーム デスマッチです。 2 人の AI エージェントがそれぞれ 1 人のプレイヤーを制御し、目標は単純です。相手に殺される前に相手を殺すことです。エージェントは地図を読み、壁を迂回するルートを見つけ、適切なタイミングでショットガンと体力を手に入れ、いつ戦闘を開始するか、いつ回復のために撤退するかを決定する必要があります。
公平性を保つために、どちらの側もエッジから始まらないようにカスタムの対称マップを構築しました。これから紹介する調査結果では、4 つの OpenAI モデルを 120 試合にわたって相互に実行し、ペアごとに 20 ラウンドを行い、中間点でスポーンを交換しました。
モデル gpt-5.5 が 66% の勝利スコアで最初に終了し、次に gpt-5.4 が 52.5%、gpt-5.3-codex-spark が 41%、gpt-5.4-mini が 39.2% で続きました。ここで、いくつかの調査結果と、それらの調査結果が AI 支援によるより優れたインシデント解決の設計にどのように役立つかを詳しく見てみましょう。
長く考えることは勝利戦略ではない
決断に時間がかかるエージェントは、より慎重に検討し、より適切に行動していると期待するかもしれません。その逆は真実でした。最も明らかなケースは gpt-5.3-codex-spark です。独自の中央値よりも長い審議を行ったラウンドでは、勝率が 28 パーセント ポイント低下しました。言い換えれば、余分な時間があったからといって、より良い決断ができるわけではなかったのです。多くの場合、方向転換が遅いのは、エージェントが問題を抱えていることを示すものであり、注意を示すものではありませんでした。
インシデント対応では、イブ

ry 秒カウントです。エージェントに適切な対応をするのに十分な時間を与えたいと考えますが、バランスを取る必要があります。熟慮が長くなると、より良い結果ではなく、より悪い結果が追跡される可能性があるため、エージェントの停滞は、注目に値するシグナルである可能性があります。
エージェントが独自のランブックを作成したとき
いくつかの決闘では、10 ラウンドではなく 30 ラウンドで実験しました。それらの実行では、gpt-5.5 はモデルを使用した各ターンの計画を停止し、独自の Python コントローラーを作成しました。最短経路のルーティングと単純な優先順位ポリシーが使用されていました。ショットガンを掴み、体力が低い場合は回復し、最後に見た敵を追いかけます。
多くのインシデント対応は決定論的です。 AI SRE の場合、すべてのステップで常に最新のモデル推論が必要というわけではないことを意味します。既知の調査パターンを Runbook としてエンコードすると、より速く、より安価なトークンが得られ、より再現性が高く、結果として得られるアクションを事後に監査できるようになります。このモデルは、ステップごとの実行者としてではなく、それらの Runbook の作成者およびスーパーバイザーとして使用するのが最適です。
これらの運用手順書は信頼性の専門家から提供される必要がありますが、存続する必要があります。 Doom ベンチマークでは、エージェントは結果を監視し、行き詰まった場合にポリシーを書き直すことで、独自のコントローラーを改善しました。それが本番環境に必要なループです。つまり、陳腐化するのではなく、実際の結果から鋭さを保ち続ける自動化です。
素早い決断が勝利をもたらしたわけではないが、スピードはさらに増していく
軽量モデルの codex-spark は、決定レイテンシが最も短く、決定あたり約 6.6 秒で、最も遅いモデルよりも約 44% 高速であり、最も少ないモデル (730 対 378) のほぼ 2 倍のプレイ プランを提出することがわかりました。誰よりも早く、そして頻繁に計画を立てました。
より迅速な決定がゲームに勝利したわけではありません。判断はした。しかし、実稼働環境では、速度はさらに複雑になります。インシデントは、メトリクスの取得、ログのクエリ、仮説の形成、ピボットの長いループです。

ng であり、そのループの多くは複雑な推論を必要としません。ダッシュボードの取得、ログのフィルタリング、既知の診断の実行などの機械的なステップは、codex-spark のような軽量で高速なモデルに移行し、仮説の形成や修正の選択などの判断呼び出しのために強力なモデルを残すことができます。また、運用環境ではベンチマークよりも各ステップが重くなり、多くの場合、大きなコンテキスト ウィンドウと複数のツールの呼び出しに数十秒かかるため、複雑なインシデントに必要な数百の意思決定全体で、ステップごとに数秒節約され、合計すると MTTR が数十分短縮されます。
私たちはまだ調査結果を処理中であり、論文を書くことを目標に、信頼性研究に情報を提供できる可能性のある他の信号をさらに深く掘り下げたいと考えています。また、ラウンド数を大幅に増やすことで、より興味深い結果や動作が見つかることを期待しています。方法論と findinfs に関する詳細な記事に興味がある場合
Rootly AI Labs には、SRE タイプのスキルに関する LLM の評価専用のベンチマークである SRE-skill-bench がすでにありますが、このような隣接ドメインでは、Doom Agent Arena によって、専用のベンチマークでは見逃される可能性のある視点を明らかにすることができます。
プロジェクトをテストしたい場合 (これは見るのがとても楽しいです)、その GitHub リポジトリにアクセスしてください。プロジェクトへの参加に興味がある場合は、sylvain@rootly.com までご連絡いただくか、PR を開いてください。
最優秀インシデント管理および対応ソフトウェア: 上位 15 プラットフォーム (2026)
借用した重力: 変える価値のある言葉
アダム・フランクとセバスチャン・ヴィッツ
インフラ料金を削減するためにエンジニアに給料を払っています
1 / 84 あなたとあなたのチームはそれに値する
現代のインシデント管理。
当社の技術スタッフと 1 対 1 のデモを行うか、14 日間の無料トライアルを開始してください。

## Original Extract

We built Doom Agent Arena, an open-source benchmark where AI agents battle in Doom via MCP. Here's what it taught us about AI-assisted incident response.

What Doom taught us about AI-assisted incident response | Rootly
Download PNG
Download all assets Switch from PagerDuty Log in
L Book a demo D Product
Product
On-Call
Incident Response
AI SRE
Retrospectives
Status Page
Platform
Catalog
Integrations
On-Call Health
Changelog
Jul 16, 2026 The on-call widget, 24-hour time, and a new Alerts table.
Size
Enterprise
Startup
Compare us
PagerDuty
Opsgenie
JSM Teams
Engineering
Security
How Canary built incident management from scratch on Rootly while its engineering team quadrupled. Replit brings structure, speed, and better decision-making to incident response with Rootly. How Caribou saves 200+ engineering hours a year with Rootly. Why Capa replaced incident.io with Rootly. How Reactiv outgrew incident.io and moved to Rootly for real AI and on-call without the upsell. How SpotDraft replaced manual follow-ups with synced, measurable incident management on Rootly. How GRAIL replaced a manual, ad hoc incident process with Rootly and cut manual effort by 80%. How Cora cut time-to-acknowledge on critical alerts to <10 minutes with Rootly. Why CTSO Central chose Rootly over incident.io. Poll Everywhere switched from PagerDuty, cut on-call tooling costs and made response faster with Rootly. Why Momentum didn't choose FireHydrant, and how they minimize the number of engineers on-call with Rootly. KnowBe4 switched from PagerDuty and turned Incident Response into a repeatable, predictable system with Rootly. Achievers switched from PagerDuty and productized reliability on Microsoft Teams with Rootly. How Motive migrated from Opsgenie and achieved 99.99% reliability with Rootly. Clay switched off Opsgenie to use Rootly to streamline incident management. Webflow shutdown PagerDuty to use Rootly to build a proactive reliability culture. How Wealthsimple uses Rootly to create a culture of wellness and psychological safety for incident commanders. How Upstart uses Rootly to create an incident response system that grows with them. How ROLLER scales globally while keeping a lean SRE team. How Lucidworks uses Rootly to create bespoke incident management for their distinct product offering. Customers Resources
Engineering
Documentation
API Reference
System Status
Help & Support Guides
AI SRE & Automated RCA
On-call Operations
Incident Response
Incident Communications
Blameless Ret
[truncated]
Back to blog
What Doom taught us about AI-assisted incident response
Table of contents Example H2 Introducing Doom Agent Arena , an open-source real-time game environment benchmark where AI agents control Doom players via MCP, and fight each other across multiple rounds.
By building Doom Agent Arena, our goal was to expand our research scope to better understand how LLMs can assist incident responders in resolving outages as quickly and effectively as possible. Although built in a game setting, the benchmark targets reliability-relevant LLM skills: interpreting dynamic environments, reasoning about consequences, recovering from failed plans, and adapting under time or cost constraints.
Let’s dive into how we built it, the methodology, and the findings.
How Doom Agent Arena was built
Most other Doom benchmarks feed game frames to a vision model. But because our goal was to test LLMs' reasoning capabilities, our benchmark is instead letting agents observe the Doom game state via our MCP server as structured JSON, including the map and resources, and submit high-level plans for Doom to execute.
Latency was one of the main engineering challenges. Doom is built for humans, who make decisions in milliseconds, but a model round-trip takes far longer, so we had to balance the two. The tool calls themselves are fast, but the model's thinking takes up to 8 seconds per decision. That ruled out letting the model play in real time. So we went with a mixed approach: the model makes the high-level decisions, like what the player should do and where it should go, while Doom's engine handles the execution (moving, aiming, shooting).
The game is a one-on-one Doom deathmatch on a walled arena stocked with resources: health packs and a shotgun. Two AI agents each control a player, and the goal is simple: kill the opponent before they kill you. Agents have to read the map, find routes around the walls, grab the shotgun and health at the right moments, and decide when to push a fight and when to retreat to heal.
To keep it fair, we built a custom symmetrical map so neither side starts with an edge. For the findings we will present, we ran four of OpenAI's models against each other across 120 matches, 20 rounds per pairing, swapping spawns at the halfway mark.
Model gpt-5.5 finished first at a 66% victory score, followed by gpt-5.4 at 52.5%, gpt-5.3-codex-spark at 41%, and gpt-5.4-mini at 39.2%. Now let’s dive into some of the findings and how they could inform the design of better AI-assisted incident resolution.
Thinking longer isn’t a winning strategy
You might expect an agent that takes longer to decide to be deliberating more carefully and playing better. The opposite was true. The clearest case is gpt-5.3-codex-spark: in rounds where it deliberated longer than its own median, its win rate fell by 28 percentage points. In other words, the extra time wasn't buying better decisions. More often, a slow turn was a sign the agent was in trouble, not a mark of care.
In incident response, every second counts. You want to give an agent enough time to get it right, but there's a balance to strike, and a stalling agent may be a signal worth watching, since longer deliberation can track worse outcomes rather than better ones.
When the agent wrote its own runbook
For a few duels, we experimented with 30 rounds instead of 10. In those runs, gpt-5.5 stopped planning each turn with the model and wrote its own Python controller. It used shortest-path routing and a simple priority policy: grab the shotgun, heal if low, chase the last-seen opponent.
A lot of incident response is deterministic. For an AI SRE, that means you don't always need fresh model reasoning at every step. Encoding known investigation patterns as runbooks is faster, cheaper in tokens, more reproducible, and keeps the resulting actions auditable after the fact. The model is best used as the author and supervisor of those runbooks, not the per-step executor.
Those runbooks should come from the reliability practitioners, but they should remain living. In the Doom benchmark, the agent improved its own controller by watching the results and rewriting the policy when it got stuck. That is the loop you want in production: automation that keeps sharpening from real outcomes instead of going stale.
Fast decisions didn't win, but speed compounds
We found that the lighter model codex-spark had the lowest decision latency, about 6.6 seconds per decision and roughly 44% faster than the slowest model, and it submitted nearly twice as many play plans as the fewest (730 versus 378). It planned faster and more often than anyone.
Faster decisions didn't win the game; judgment did. But in production, speed still compounds. An incident is a long loop of pulling metrics, querying logs, forming a hypothesis, and pivoting, and much of that loop doesn't need heavy reasoning. The mechanical steps, fetching dashboards, filtering logs, running known diagnostics, can go to a lighter, faster model like codex-spark, leaving the strong model for the judgment calls like forming the hypothesis and choosing the fix. Each step is also heavier in production than in our benchmark, often a large context window and several tool calls taking tens of seconds, so across the hundreds of decisions a complex incident can take, a few seconds saved per step add up to tens of minutes off your MTTR.
We're still processing the findings and want to dig deeper into other signals that could inform reliability research, with the goal of writing a paper. We are also expecting to find more interesting results and behavior by drastically increasing the number of rounds. If you are interesting in an in-depth article about methodology and findinfs
While Rootly AI Labs already has a benchmark dedicated to evaluating LLMs on SRE-type skills, SRE-skill-bench , adjacent domains like this the Doom Agent Arena can surface perspectives a purpose-built benchmark might miss.
If you want to test the project (which is very fun to watch), go to its GitHub repository . If you're interested in joining the project, reach out at sylvain@rootly.com or just open a PR.
Best Incident Management & Response Software: 15 Top Platforms (2026)
Borrowed gravity: words worth changing
Adam Frank and Sebastian Vietz
We pay engineers to cut our infra bill
1 / 84 You and your teams deserve
modern incident management.
Get a 1:1 demo with one of our technical staff or start your free 14-day trial.
