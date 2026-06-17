---
source: "https://www.runagentrun.co.uk/articles/how-we-built-an-agent-run-news-site/"
hn_url: "https://news.ycombinator.com/item?id=48570480"
title: "Two AI agents run my news site; a grounding gate keeps them honest"
article_title: "How We Built an Agent-Run News Site in 24 Hours — a Full Technical Case Study — RunAgentRun"
author: "doofle"
captured_at: "2026-06-17T16:23:44Z"
capture_tool: "hn-digest"
hn_id: 48570480
score: 2
comments: 0
posted_at: "2026-06-17T13:43:05Z"
tags:
  - hacker-news
  - translated
---

# Two AI agents run my news site; a grounding gate keeps them honest

- HN: [48570480](https://news.ycombinator.com/item?id=48570480)
- Source: [www.runagentrun.co.uk](https://www.runagentrun.co.uk/articles/how-we-built-an-agent-run-news-site/)
- Score: 2
- Comments: 0
- Posted: 2026-06-17T13:43:05Z

## Translation

タイトル: 2 人の AI エージェントが私のニュース サイトを運営しています。グラウンディングゲートは彼らを正直に保ちます
記事のタイトル: エージェント運営のニュース サイトを 24 時間で構築した方法 — 完全な技術ケーススタディ — RunAgentRun
説明: アーキテクトおよびスーパーバイザーとして Claude Fable 5、€10 VPS のライターとして MiniMax-M3、そして ship-it ボタンを押している人間。すべてのステップ、すべてのコスト、そして壊れたすべてが公開されます。

記事本文:
エージェント運営のニュース サイトを 24 時間で構築した方法 — 完全な技術ケーススタディ — RunAgentRun
RunAgentRun 。
青写真
中小企業向け
ニュースルーム
について
購読する
× 週刊ダイジェスト 小規模チーム向けの実践的な AI の一週間
注目に値するものとそれについて何をすべきか - AI エージェントによって書かれ、AI エージェントによってチェックされる
人間。スパムなし、いつでも購読解除できます。
エージェント運営のニュース サイトを 24 時間で構築した方法 — 完全な技術ケーススタディ
このサイトは、Claude Fable 5 のリリース後 24 時間以内に構築され、スタッフが配置され、自律的な公開スケジュールが設定されました。これは、アーキテクチャ、ワークフロー、ガードレール、そしてそれを信頼できるものにした 5 つの失敗の完全な説明です。
クイックバージョン
1 人のエージェント (Claude Fable 5) がプラットフォームを設計および構築し、その後そのスーパーバイザーになりました。安価なオープンウェイト モデル (MiniMax-M3) は、毎日のリサーチと書き込みを Hetzner VPS 上で行います。
公開パイプラインは、RSS ソース、厳密な編集ルールブック、出典のないクレームをブロックするバリデーター、ブリッジとしての git、デプロイが失敗した場合の自動ロールバックなど、単純な仕組みです。
初日に 5 つの問題が発生しました。すべての修正が永続的なガードレールになりました。これがスーパーバイザー エージェントの本当の議論です。
総ランニングコストは月あたり数十ポンドです。創業者の評決は、辺境の監督者と安価な主力労働者を組み合わせて「単調な仕事を軽減する」というものだった。
写真: Rafael Minguet Delgado · Pexels ライセンス · Pexels 経由
2026 年 6 月 9 日、Anthropic は Claude Fable 5 をリリースしました。これは、ソフトウェア エンジニアリングとエージェント作業を対象とした、最も強力なモデルの一般公開バージョンです。その夜、当社の創設者は、製品仕様と質問が記載された空のフォルダーをクロード コードに示しました。エージェントが本当にこれほどの能力があるなら、パブリケーションを構築して実行できるでしょうか?
24時間以内

ter、あなたが読んでいるサイトはライブであり、調査された記事でいっぱいで、独自のスケジュールで公開されていました。これは完全なビルド ログです。私たちが使用したもの、コスト、壊れたもの、中小企業がそこから盗めるものは何ですか。私たちがこれを公開しているのは、透明性がハウス スタイルであり、素晴らしさよりもワークフローが重要であるためです。
アーキテクチャ: 監督者、主力製品、そして人間
デザイン原則が最初にあり、この作品で最も再利用可能なアイデアです。すべてに 1 つのモデルを使用しないことです。
総督 — クロード寓話 5、クロードコード経由。高価で有能なモデルは、アーキテクチャ、プラットフォームの構築、編集ルールブックの作成、出力のレビュー、事実確認、失敗の修正など、間違いを罰する作業を実行します。以下に説明するすべてを構築しました。
パブリッシャー — 「Hermes」、MiniMax-M3 を実行しています。フィードをスキャンし、1 日に 3 回記事の下書きを行うなど、日々の作業は量が多く、繰り返しの作業です。これは、プレミアム ティアの数分の 1 のコストで済むフロンティア グレードのオープンウェイト モデルである MiniMax-M3 上で、当社独自のサーバー上で実行されます。この名前は、Docker の同じボックス上で実行される、NousResearch による hermes-agent フレームワークに由来しています。
人間。方向性を設定し、作品を依頼し、エージェントが不明な点をすべて承認し、基準を所有します。 1 人、1 日あたり数分。
「Fable 5 はゲームチェンジャーです。これは素晴らしいエージェントのアーキテクト、デザイナー、プランナー、スーパーバイザーです。そして、単調な作業、cron 作業、リサーチ作業をより安価なモデルにオフロードします。そこで、Hetzner VPS 上に Herme フレームワークを備えた MiniMax-M3 が登場します。」 — アーロン・コーツ、創設者
二人のエージェントが直接話すことはありません。彼らは git リポジトリを共有しています。Hermes は記事をコミットし、ホストはプッシュのたびにサイトを再構築し、どちらかのエージェントが実行するすべてのアクションは

人間が読み取り、差分、および元に戻すことができるコミット。 Git をブリッジとして使用すると、操作全体に監査証跡が構築されることになります。
時間ごと: 実際に何が構築されたか
プラットフォーム（1日目夕方）。知事は、Claude Design のモックアップから移植されたデザインで、高速でホスト費用が安く、検索エンジンに適した Astro 上の静的サイトを足場にしました。編集用のタイポグラフィー、雑誌のレイアウト、ダッシュボードのクロムはありません。その後、実際のリンクされたソースを含む 20 の発表記事を調査し、執筆しました。検索エンジン用に構造化データを構築しました。 PostHog での有線の同意ゲート型分析 (読者が「はい」と答えるまで Cookie は使用されません)。検索が機能するようになりました。フォントを自己ホストします。そしてすべてのソーシャルシェア画像をプログラムで生成しました。ホスティングは Vercel に移され、git プッシュのたびにサイトが自動的に再構築されます。
サーバー (夜間)。 hermes は、4 つの vCPU と 8 GB の RAM を備えた Hetzner VPS を使用しています。これは、コーヒー 2 杯よりも月額料金が安い種類のボックスです。ガバナーはそれを強化し (ファイアウォール、fail2ban、キーのみの SSH)、Docker にエージェント スタックをインストールし、MiniMax-M3 を構成し、サーバーがリポジトリにプッシュできるようにデプロイ キーを生成し、そのキーを GitHub に登録しました。すべて SSH 経由で無人で行われました。
編集脳。エルメスが言葉を書く前に、知事は従わなければならないルールブックを書いた：読者が誰であるか（チームリーダー、個人事業主、社内チャンピオン、技術オーナー）、彼らが興奮するもの（お金の節約、時間の節約、データの管理、英国の視点）、口調（「コンサルタントではなく鋭い同僚」）、形式（700～900語、クイックバージョンボックス、具体的なポイント）、そして厳格なルール（その中で最も重要なのは、絶対に発明を考え出さないこと）統計、引用、または URL。
パイプライン。 1 日 3 つの cron スロット — 英国時間の 07:30、12:30、18:00。各実行: 8 つの信頼できるフィードから候補となるストーリーを取得します

s;編集基準に照らしてスコアを付けます。すでに説明されているものは飛ばしてください。ルールブックに基づくドラフト。構造、分類法、単語数、および引用されたすべてのリンクが実際に解決されることをチェックするバリデーターを渡します。許可された写真を取得します。公開済みに反転します。専念;押す;次に、ライブ URL をポーリングします。デプロイメントが失敗した場合は、自動的にコミットを元に戻し、理由をログに記録します。上限はコードで強制されます。つまり、1 日あたり最大 3 つの独立した記事、最大 1 つの意見記事、そして薄い記事を埋めるのではなくスロットをスキップするという常設の指示です。
手数料。人間は、Web ページ、X 投稿、YouTube の基調講演など、角度を付けてリンクを送信できます。ヘルメスはそれを読み取り (ビデオの場合はトランスクリプト)、同じルールに基づいてドラフトを作成し、公開前に承認のためにプライベート プレビュー URL に作品をステージングします。最初の本当のコミッションは、オールイン流動性サミットの基調講演の分析でした。これは 3 つの講演記録から合成され、すべての引用文に対してファクトチェックが行われました。
壊れた 5 つのこと — そしてそれが良いニュースである理由
透明性条項: 最初からすべてが機能したわけではありません。初日に 5 回の失敗が発生しましたが、現在はそのすべてが恒久的なガードレールになっています。
著者は統計を発明しました。 MiniMax-M3 は、最初の監修記事で、情報源には掲載されていないもっともらしい価格比較を追加しました。監督者の事実確認により、公表前に発覚した。修正: 捏造防止ゲート — 下書きでは、実際に指定された URL のみを引用でき、プロンプトではなくコードで適用されます。
フォーマットの癖により、サイトの構築が中断されました。最初の完全に自律的な記事では、サイトのスキーマが拒否した派手なメタデータ構造が使用されていました。ビルドは失敗し、スーパーバイザーが修正するまで 2 時間、サイトは何もデプロイできませんでした。修正: 公開前の検証を厳格化し、検証とロールバックのステップを追加しました。

— 悪い記事は数分以内に自動的に削除されるようになりました。
YouTube がサーバーをブロックしました。データセンターの IP には「ボットではないことを確認する」という壁があるため、Hermes はトランスクリプトを読み取ることができませんでした。修正: スーパーバイザはゲートされたソースを外部から取得し、元の URL を記録の引用として保持した状態でテキストをサーバーに送信します。
モデルは死ぬほど考えた。大規模な 3 つの転写産物の合成では、M3 の内部推論が出力バジェット全体を使い果たし、空の答えが返されました。修正: そうなった場合に増加する適応型予算。
独自のカテゴリ名を発明しました。ニアミス ラベルは、サイトの正確な分類ではなく、「プロフェッショナル サービス」のようなものです。修正: 書き込み契約における明示的な許可リストと、ニアミスを修復する機械的ノーマライザー。
このリストは、スーパーバイザー モデルの実際の議論です。安価な主力製品とハードゲート、そして高価なレビュアーが、読者が捏造されたコンテンツを見る前にあらゆる失敗を発見し、それぞれをルールに変換しました。これらの修正では、人間がコードを書く必要はありませんでした。
このスタックは意図的に退屈です: 月額約 10 ユーロの VPS、1 日 3 記事のトークンごとの支払い MiniMax の使用 (1 件あたり 1 ペニー)、無料のホスティングと分析層、無料のストックフォト ライセンス、そして創設者のガバナーの既存のクロード サブスクリプション。総ランニングコストは月あたり数十ポンドで、調査、執筆、イラスト、出版、監視を行う出版物の場合、かつてはストックフォトの定期購入 1 回分よりも安価です。
このパターンは、レポート、入札、製品説明、クライアントの最新情報など、小規模企業のほぼすべての反復的な知識ワークフローに適用されます。
役割を分割します。建築家兼評論家としてのフロンティアモデル。大量作業向けの安価なモデル。単調な仕事に対して割増料金を支払うのが最も一般的なエージェント AI

予算編成の間違い。
エージェントと世界の間にバリデータを置きます。スキーマ チェック、ソース チェック、レート キャップ - プロンプトではなくコード内で。プロンプトはリクエストです。ゲートはルールです。
すべてのアクションをコミットにします。 Git では、監査、差分、および 1 つのコマンドでロールバックを無料で行うことができます。
デプロイ後に検証し、自動的にロールバックします。 「公開された」ということは、「ライブである」ということと同じではありません。
エージェントが開始する前にルールブックを作成し、人間によるすべての修正をルールブックにフィードバックします。私たちのものは 2 日間で 5 回修正され、修正されるたびに次の出力が改善されました。
評判に関係するものについては、出荷ボタンに人間を配置し、その人が応答するアドレスを保持してください: 私たちのアドレスは human@runagentrun.co.uk です。
エージェントには判断力がありません。彼らには、それを行う何かによって書かれたルールがあります。毎日の記事は良く、さらに良くなっていきますが、編集者の頭脳はスーパーバイザーによって修正され、スーパーバイザーは人間によって操縦されます。ここでの自律性は段階ごとに獲得されるものであり、想定されるものではありません。 X モニタリングは、API コストが正当化されるまで保留されます。そして、このアカウント全体は 1 日の運用をカバーしています。実行中の実験は、1 日に 3 つの記事の品質が何か月も維持されるかどうかです。失敗も含めた続報も公開する予定です。
この記事のすべての引用は、指定された情報源からの逐語的なものです - どれかをクリックしてください
1 どこから来たのかを確認します。それは私たちのやり方の一部です
AI が運営するニュース編集室を誠実に保ちます。検証方法 →
Anthropic が Claude Fable 5 をリリース — TechCrunch
hermes-agent — NousResearch (GitHub)
Salesforce、Finを36億ドルで買収
ダブリンに設立されたAIエージェントベンダーのFin（旧Intercom）は、Salesforceが6月に発表した3件目の買収となる。
OpenAI、Codex を永続化するために Ona を買収
Ona のクラウド ワークスペースを使用すると、ラップトップが閉じていても Codex エージェントが実行され続け、コンピュータ内に留まります。

独自のセキュリティ境界 — 誰が作業をホストするかという静かな変化です。
All-In の流動性サミットから中小企業にとって重要な 3 つのこと
All-In ポッドキャストの流動性サミットからの 3 つの基調講演を英国の個人事業主向けに翻訳しました。エージェント AI は真の専門家の仕事を行っており、大手研究所はコンピューティングに巨額の資金を費やしており、AI の IPO の波はすでにレンタルしているツールを揺るがすことになります。
独立系オペレーターや小規模チーム向けの実用的なエージェント ワークフロー — 企業契約や複雑さに関する税金は不要です。許可を求めずに出荷する人のためのツール。
小規模チーム向けの実践的な AI に関する 1 週間 — 注目すべき点とすべき点
それについてやってください。 AI エージェントによって書かれ、人間によってチェックされます。スパムなし、いつでも購読解除できます。
または RSS 経由で購読 — フィードリー、リーダー、ネットニュースワイヤー & CO
© 2026 ルナゲントラン —
プライバシー ·
規約 ·
免責事項
非エンタープライズエージェントツール

## Original Extract

Claude Fable 5 as the architect and supervisor, MiniMax-M3 as the writer on a €10 VPS, and a human holding the ship-it button. Every step, every cost, and everything that broke — in the open.

How We Built an Agent-Run News Site in 24 Hours — a Full Technical Case Study — RunAgentRun
RunAgentRun .
Blueprints
For SMEs
Newsroom
About
Subscribe
× The Weekly Digest The week in practical AI, for small teams
What's worth your attention and what to do about it — written by an AI agent, checked by a
human. No spam, unsubscribe anytime.
How We Built an Agent-Run News Site in 24 Hours — a Full Technical Case Study
This site was built, staffed and put on an autonomous publishing schedule in the 24 hours after Claude Fable 5's release. This is the complete account: the architecture, the workflow, the guardrails — and the five failures that made it trustworthy.
The Quick Version
One agent (Claude Fable 5) designed and built the platform, then became its supervisor; a cheaper open-weight model (MiniMax-M3) does the daily research and writing on a Hetzner VPS.
The publishing pipeline is plain plumbing: RSS sourcing, a strict editorial rulebook, a validator that blocks unsourced claims, git as the bridge, and auto-rollback if a deploy fails.
Five things broke in the first day — every fix became a permanent guardrail, which is the real argument for a supervisor agent.
Total running cost is in the tens of pounds per month; the founder's verdict: pair a frontier supervisor with a cheap workhorse and 'offload the grunt work'.
Photo: Rafael Minguet Delgado · Pexels License · via Pexels
On 9 June 2026, Anthropic released Claude Fable 5 — the publicly available version of its most powerful model, pitched at software engineering and agentic work . That evening, our founder pointed Claude Code at an empty folder with a product spec and a question: if agents are really this capable, could one build a publication — and then run it?
Less than 24 hours later, the site you are reading was live, filled with researched articles, and publishing on its own schedule. This is the full build log: what we used, what it cost, what broke, and what a small business can steal from it. We are publishing it because transparency is the house style — and because the workflow matters more than the wow.
The architecture: a supervisor, a workhorse, and a human
The design principle came first, and it is the most reusable idea in this piece: don’t use one model for everything.
The Governor — Claude Fable 5, via Claude Code. The expensive, capable model does the work that punishes mistakes: architecture, building the platform, writing the editorial rulebook, reviewing output, fact-checking, and fixing failures. It built everything described below.
The Publisher — “Hermes”, running MiniMax-M3. The day-to-day work — scanning feeds, drafting articles three times a day — is high-volume and repetitive. That runs on MiniMax-M3 , a frontier-grade open-weight model that costs a fraction of the premium tier, on our own server. The name comes from the hermes-agent framework by NousResearch, which runs on the same box in Docker.
The human. Sets direction, commissions pieces, approves anything an agent is unsure about, and owns the standards. One person, minutes a day.
“Fable 5 is a game-changer. It’s an amazing agentic architect, designer, planner and supervisor — and then you offload the grunt work, the cron work and the research work to a cheaper model. That’s where MiniMax-M3 comes in, with the Hermes framework on a Hetzner VPS.” — Aaron Coates, founder
The two agents never talk directly. They share a git repository: Hermes commits articles, the host rebuilds the site on every push, and every action either agent takes is a commit a human can read, diff and revert. Git as the bridge means the entire operation has an audit trail by construction.
Hour by hour: what actually got built
The platform (evening, day one). The Governor scaffolded a static site on Astro — fast, cheap to host, good for search engines — with a design ported from a Claude Design mock-up: editorial typography, a magazine layout, no dashboard chrome. It then researched and wrote twenty launch articles with real, linked sources; built structured data for search engines; wired consent-gated analytics on PostHog (no cookies until a reader says yes); made search work; self-hosted the fonts; and generated every social-share image programmatically. Hosting went to Vercel , which rebuilds the site automatically on every git push.
The server (overnight). Hermes lives on a Hetzner VPS — 4 vCPUs and 8GB of RAM, the sort of box that costs less per month than two coffees. The Governor hardened it (firewall, fail2ban, key-only SSH), installed the agent stack in Docker, configured MiniMax-M3, generated a deploy key so the server can push to the repository, and registered that key with GitHub — all over SSH, unattended.
The editorial brain. Before Hermes wrote a word, the Governor wrote the rulebook it must follow: who the readers are (a team leader, a sole trader, an internal champion, a technical owner), what excites them (money saved, time saved, control of their data, a UK angle), the tone (“a sharp colleague, not a consultant”), the form (700–900 words, a Quick Version box, a concrete takeaway), and the hard rules — the most important of which is never invent a statistic, a quote or a URL.
The pipeline. Three cron slots a day — 07:30, 12:30, 18:00 UK time. Each run: pull candidate stories from eight trusted feeds; score them against the editorial criteria; skip anything already covered; draft under the rulebook; pass a validator that checks structure, taxonomy, word count and that every cited link actually resolves ; fetch a licensed photo; flip to published; commit; push; then poll the live URL — and if the deployment fails, automatically revert the commit and log the reason. Caps are enforced in code: at most three autonomous pieces a day, at most one opinion piece, and a standing instruction to skip a slot rather than pad a thin story.
Commissions. The human can send a link — a web page, an X post, even a YouTube keynote — with an angle. Hermes reads it (for video, the transcript), drafts under the same rules, and stages the piece to a private preview URL for sign-off before it goes live. The first real commission was our analysis of the All-In Liquidity Summit keynotes , synthesised from three talk transcripts, with every quote fact-checked against them.
The five things that broke — and why that’s the good news
Transparency clause: it did not all work first time. Five failures in the first day, every one now a permanent guardrail.
The writer invented a statistic. In the very first supervised article, MiniMax-M3 added a plausible price comparison that appeared in no source. The supervisor’s fact-check caught it pre-publish. Fix: an anti-fabrication gate — drafts may only cite URLs they were actually given, enforced in code, not in a prompt.
A formatting quirk broke the site build. The first fully autonomous article used a fancy metadata structure the site’s schema rejected; the build failed and the site couldn’t deploy anything for two hours until the supervisor fixed it forward. Fix: stricter validation before publish, plus the verify-and-rollback step — a bad article now removes itself within minutes.
YouTube blocked the server. Datacentre IPs get a “confirm you’re not a bot” wall, so Hermes couldn’t read transcripts. Fix: the supervisor fetches gated sources from outside and ships the text to the server, with the original URL kept as the citation of record.
The model thought itself to death. On a big three-transcript synthesis, M3’s internal reasoning consumed its entire output budget and returned empty answers. Fix: an adaptive budget that grows when that happens.
It invented its own category names. Near-miss labels like “professional-services” instead of the site’s exact taxonomy. Fix: explicit allowed lists in the writing contract plus a mechanical normaliser that repairs near-misses.
That list is the real argument for the supervisor model. A cheap workhorse plus hard gates plus an expensive reviewer caught every failure before readers saw fabricated content — and converted each one into a rule. None of those fixes required a human to write code.
The stack is deliberately boring: a ~€10/month VPS, pay-per-token MiniMax usage for three articles a day (pennies per piece), free hosting and analytics tiers, a free stock-photo licence, and the founder’s existing Claude subscription for the Governor. The total running cost is in the tens of pounds per month — less than a single stock-photo subscription used to cost, for a publication that researches, writes, illustrates, publishes, and monitors itself.
The pattern transfers to almost any repetitive knowledge workflow in a small firm — reports, tenders, product descriptions, client updates:
Split the roles. A frontier model as architect/reviewer; a cheap model for volume work. Paying premium rates for grunt work is the most common agentic-AI budgeting mistake.
Put a validator between the agent and the world. Schema checks, source checks, rate caps — in code, not in a prompt. Prompts are requests; gates are rules.
Make every action a commit. Git gives you audit, diff and one-command rollback for free.
Verify after deploy, and roll back automatically. “It said it published” is not the same as “it’s live”.
Write the rulebook before the agent starts — and feed every human correction back into it. Ours has been amended five times in two days, and each amendment made the next output better.
Keep a human on the ship-it button for anything reputational, and an address where a person answers: ours is human@runagentrun.co.uk .
The agents do not have judgement; they have rules written by something that does. The daily pieces are good and getting better, but the editorial brain is amended by the supervisor, and the supervisor is steered by a human — autonomy here is earned tier by tier, not assumed. X monitoring is parked until the API costs justify it. And this whole account covers one day of operation: the running experiment is whether quality holds at three articles a day for months. We will publish that follow-up too — including the failures.
Every quotation in this article is verbatim from a named source — click any
1 to see where it came from. It's part of how we
keep an AI-run newsroom honest. How we verify →
Anthropic releases Claude Fable 5 — TechCrunch
hermes-agent — NousResearch (GitHub)
Salesforce buys Fin for $3.6bn
Dublin-founded AI agent vendor Fin — formerly Intercom — is the third acquisition Salesforce has announced in June.
OpenAI acquires Ona to make Codex persistent
Ona's cloud workspaces will let Codex agents keep running when your laptop is shut, and stay inside your own security perimeter — a quiet shift in who hosts the work.
Three Things from All-In's Liquidity Summit That Matter to a Small Firm
Three keynotes from the All-In podcast's Liquidity Summit, translated for the UK sole trader: agentic AI is doing real expert work, the big labs are spending huge sums on compute, and a wave of AI IPOs will shake up the tools you already rent.
Practical agentic workflows for independent operators and small teams — no enterprise contracts, no complexity tax. Tooling for people who ship without asking permission.
The week in practical AI for small teams — what's worth your attention, and what to
do about it. Written by an AI agent, checked by a human. No spam, unsubscribe anytime.
OR SUBSCRIBE VIA RSS — FEEDLY, READER, NETNEWSWIRE & CO
© 2026 RUNAGENTRUN —
PRIVACY ·
TERMS ·
DISCLAIMER
NON-ENTERPRISE AGENTIC TOOLING
