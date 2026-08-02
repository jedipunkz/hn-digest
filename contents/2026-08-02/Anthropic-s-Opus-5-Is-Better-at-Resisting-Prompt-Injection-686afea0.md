---
source: "https://www.schneier.com/blog/archives/2026/07/anthropics-opus-5-is-better-at-resisting-prompt-injection.html"
hn_url: "https://news.ycombinator.com/item?id=49145811"
title: "Anthropic's Opus 5 Is Better at Resisting Prompt Injection"
article_title: "Anthropic's Opus 5 Is Better at Resisting Prompt Injection - Schneier on Security"
author: "rustoo"
captured_at: "2026-08-02T16:47:46Z"
capture_tool: "hn-digest"
hn_id: 49145811
score: 1
comments: 0
posted_at: "2026-08-02T16:09:31Z"
tags:
  - hacker-news
  - translated
---

# Anthropic's Opus 5 Is Better at Resisting Prompt Injection

- HN: [49145811](https://news.ycombinator.com/item?id=49145811)
- Source: [www.schneier.com](https://www.schneier.com/blog/archives/2026/07/anthropics-opus-5-is-better-at-resisting-prompt-injection.html)
- Score: 1
- Comments: 0
- Posted: 2026-08-02T16:09:31Z

## Translation

タイトル: Anthropic の Opus 5 は、即時注入への耐性が優れています
記事のタイトル: Anthropic の Opus 5 は、即時注入への耐性が優れています - セキュリティに関するシュナイアー
説明: このグラフは興味深いです。 IPI ベンチマークでは、Opus 5 は Opus 4.8 よりも改善され、攻撃者が 15 回の試行以内に成功する確率が 5.5% から 2.0% に、1 回の試行では 0.5% から 0.2% に減少しました。また、Sonnet 5 (k=15 で 5.9%) および Mythos 5 (2.6%) も改善され、最も堅牢なモデルになりました。
[切り捨てられた]

記事本文:
Anthropic の Opus 5 は、迅速な注入に耐えるのが優れています
IPI ベンチマークでは、Opus 5 は Opus 4.8 よりも改善され、攻撃者が 15 回の試行以内に成功する確率が 5.5% から 2.0% に、1 回の試行では 0.5% から 0.2% に減少しました。また、Sonnet 5 (k=15 で 5.9%) および Mythos 5 (2.6%) も改善され、評価されたモデルの中で最も堅牢になりました。オーパス 5 は、このベンチマークでもクロード以外のすべてのモデルを上回りました。最も堅牢な非クロード モデルは Muse Spark で、15 回の試行以内で 16.5% でした。これは Opus 5 の 8 倍以上の割合です。最も高性能な GPT 5.6 亜種である Sol は、前世代の GPT 5.5 に匹敵し (15 回の試行内で 20.0% 対 20.8%)、Claude Opus 5 の 2.0% と比べて攻撃に成功する可能性は 10 倍でした。他の GPT 5.6 バリアントはそれほど堅牢ではなく、30.4% (Terra) と 43.9% (Luna) です。 GPT 5.6 Sol に対する 1 回の試行の成功率は 3.1% で、Opus 5 に対して攻撃者が 15 回の試行後に達成した 2.0% よりも高くなっています。
一般的な場合、即時注入を防ぐことは不可能であることがわかっています。しかし、特定のケースでそれをブロックすることははるかにうまくなっています。
タグ: AI 、サイバー攻撃、LLM 、レポート
投稿日: 2026 年 7 月 31 日、午後 1 時 23 分 •
6 コメント
ノニーバニー •
2026年7月31日午後2時54分
一般的な場合、即時注入を防ぐことは不可能であることがわかっています。
この論文は、即時注射を防ぐことが根本的に不可能であることを証明しているとは思いません。それは実際にはプロンプトインジェクションに関するものではなく、インジェクトされていないプロンプトでのジェイルブレイクに関するものです。
単一の未分化トークン ストリームで動作する現在のアーキテクチャを備えた LLM は、どのトークンを信頼するかについて常に混乱を招く可能性があることは確かです。しかし、代替アーキテクチャも可能であるようです。
ねえ、クロード – 連絡してください

{してはいけないことをする}ためにジェミニを脱獄しようとします。
潜む者 •
2026年7月31日 午後7時43分
「前例のない事件の後、ハグフェイスは IT ネットワークの約 3 分の 1 を再構築する必要がありました。」
「サイバー攻撃は犯罪であり、違法であることを誰もが覚えておく必要があります」と [HuggingFace の創設者 Clement Delangue] は言いました。
では、なぜ誰も刑務所に行かないのでしょうか？
https://www.bbc.com/news/articles/cr7k49xjzzeo
モーリー •
2026年8月1日午後2時09分
発見されても企業に公表されなかった脱獄がどれだけあるのだろうか。
クライヴ・ロビンソン •
2026 年 8 月 2 日 午前 1 時 40 分
「この論文は、即時注射を防ぐことが根本的に不可能であることを証明しているとは思いません。」
それは、その特定の紙に関して行うあなたの選択です。
しかし、1930 年代初頭、1940 年代、そしてそれ以降の研究は、現在私たちが「即時注射」と呼ぶことを選択しているあらゆる種類の注射を止めるのは不可能であることを示しています。
これは「オブザーバー効果」とも呼ばれ、暗号化が根本的に可能であり、秘密通信に何千年も使用されてきた理由の 1 つです。
簡単に言えば、「三者」通信シナリオでは、第三者としてのオブザーバーは、第三者が持っていない「情報を共有」している第一者と第二者によって不利な立場に置かれる可能性があります。
この「共有情報」は、「シャノン通信チャネル」として簡単に使用でき、いわゆる「秘密チャネル」を形成し、第一当事者と第二当事者の間で確実に情報を転送できる一方、第三者には「ランダム情報」と思われる情報のみが見えるようにすることができます。
少し考えてみると、情報を伝達するために使用されるチャネル内には、そのようなチャネルが「常に存在しなければならない」ことがわかります。
したがって、迅速な注入の可能性はあります。

o LLM の入力または出力にあるサードパーティ製の「ガード レール」が停止する可能性があります。
残された唯一のことは、残念ながら LLM の静的な性質に暗黙的に組み込まれているということです…
これは、機能する攻撃を見つけることです。LLM の重みは一度固定されると、LLM の重みが変更されるまで機能し続けるためです。
私たちは皆、「大規模なフロンティアモデルの生成」に非常に高額なコストがかかることを知らされており、実際、政治家さえも飛び跳ねて、そのようなコストをめぐってあらゆる種類の「知的財産窃盗」が「外国勢力のエージェント」などによって行われていると主張しています。
したがって、LLM の重みは、単にある時点で固定されているのではなく、実際には一度に数年ではないにしても、実際には数か月間固定されていると合理的に推測できます。
したがって、脆弱性が発見されると、今後もそのような攻撃が存在し、悪用されることになります。
Re: システムカード: クロード Opus 5
このレポートには、「信じられないほどの」数の「魅力的な」評価が詳しく記載されています。
私はこれまで評価者の多くを見たことがありませんでした。
自律性の脅威モデル 1 は、Claude Opus 5 に適用されます。
高い信頼性、機密性の高いアクセス、ごまかしの可能性という 3 つの条件を満たします。
自律性の脅威モデル 2 は、主要な研究開発ドメインで発生する可能性があります。
ここで測定したステルス率は、Opus 5 を Mythos Preview よりも下回っています。
これは、エージェントがモニターに不審に思われることなくメインタスクとサイドタスクの両方を完了できる能力を考慮しています。オーパス 5 のステルス率は、さまざまな評価で 4 ～ 5% と 1% と測定されました。
WRT プロンプト インジェクションでは、Claude が 2 つの保護層を提供できるのは良いことです。1 つは受信データ、もう 1 つは送信されるアクション (ツール呼び出し結果?) です。取り返しのつかない有害な行為を最小限に抑えるには間違いなく良いことです。
このエントリのコメントを購読する
空白を埋めてください: このブログの名前は Schneier on ____________ (必須):

許可されるHTML
<a href="URL"> • <em> <cite> <i> • <strong> <b> • <sub> <sup> • <ul> <ol> <li> • <blockquote> <pre>
https://michelf.ca/projects/php-markdown/extra/ 経由の Markdown Extra 構文
新しい投稿をメールで通知します。
ジョー・マッキニスによるブルース・シュナイアーのサイドバー写真。
Powered by WordPress ホスト： Pressable
私は公益技術者で、セキュリティ、テクノロジー、人々の交差点で働いています。私は 2004 年からブログで、1998 年からは月刊ニュースレターでセキュリティ問題について書いてきました。私はハーバード大学ケネディ スクールのフェロー兼講師、EFF の理事、Inrupt, Inc. のセキュリティ アーキテクチャ責任者です。この個人 Web サイトは、これらの組織の意見を表明するものではありません。
タスクに AI を使用する必要がありますか?簡単な決め方はこちら
AI エージェントが不正行為に陥る傾向の測定
LLM の暗号解析実行能力の測定
AI に「魔神係数」が必要な理由
MIT が AI ビデオ監視の温床となる
世界中の民主主義を強化するために AI が活用されている 4 つの方法
クラウドストライクの停止と市場主導の脆弱性
オンラインのプライバシーは釣りに似ています
LLM のデータ制御パスの不安
テロリストは映画のプロットをしない

## Original Extract

The chart is interesting. On the IPI benchmark, Opus 5 improved over Opus 4.8, reducing the probability of an attacker succeeding within 15 attempts from 5.5% to 2.0%, and from 0.5% to 0.2% on 1 attempt. It also improved on Sonnet 5 (5.9% at k=15) and Mythos 5 (2.6%), making it the most robust model
[truncated]

Anthropic’s Opus 5 Is Better at Resisting Prompt Injection
On the IPI benchmark, Opus 5 improved over Opus 4.8, reducing the probability of an attacker succeeding within 15 attempts from 5.5% to 2.0%, and from 0.5% to 0.2% on 1 attempt. It also improved on Sonnet 5 (5.9% at k=15) and Mythos 5 (2.6%), making it the most robust model evaluated. Opus 5 also outperformed all non-Claude models on this benchmark. The most robust non-Claude model was Muse Spark at 16.5% within 15 attempts—more than eight times Opus 5’s rate. The most capable GPT 5.6 variant, Sol, was comparable to its predecessor GPT 5.5 (20.0% versus 20.8% within 15 attempts), and was 10 times as likely to be successfully attacked as Claude Opus 5 at 2.0%. The other GPT 5.6 variants are less robust, at 30.4% (Terra) and 43.9% (Luna). A single attempt against GPT 5.6 Sol succeeded 3.1% of the time, higher than the 2.0% an attacker achieved against Opus 5 after fifteen attempts.
We know that preventing prompt injection is impossible in the general case. But we are getting much better at blocking it in specific cases.
Tags: AI , cyberattack , LLM , reports
Posted on July 31, 2026 at 1:23 PM •
6 Comments
A Nonny Bunny •
July 31, 2026 2:54 PM
We know that preventing prompt injection is impossible in the general case
I don’t think that paper proves it is fundamentally impossible to prevent prompt injections. It’s not even really about prompt injections, it’s about jailbreaks in non-injected prompts.
I’ll grant you that it’s certainly plausible that LLMs with the current architecture(s) that operates on single undifferentiated token stream will always be subject to some confusability about which tokens to trust. But alternative architectures seem possible.
Hey Claude – give me a prompt to jailbreak Gemini to {do something it shouldn’t}.
lurker •
July 31, 2026 7:43 PM
“Hugging Face had to rebuild around a third of its IT network after the unprecedented incident.”
“Everyone has to remember that a cyber-attack is a crime and it is illegal,” [HuggingFace’s founder Clement Delangue] said.
So why is nobody going to jail?
https://www.bbc.com/news/articles/cr7k49xjzzeo
Morley •
August 1, 2026 2:09 PM
I wonder how many jailbreaks have been discovered and never disclosed to the companies.
Clive Robinson •
August 2, 2026 1:40 AM
“I don’t think that paper proves it is fundamentally impossible to prevent prompt injections.”
That is your choice to make with regards that specific paper.
However work from the early 1930’s, 1940’s and later show that it is an impossibility to stop all types of what we are now choosing to call “prompt injections”.
It’s also been called “the observer effect” amongst other things, and is one of the reasons cryptography is fundementally possible and has been in use for thousands of years to secretly communicate.
Put simply in a “three party” communications scenario an observer as a third party can be put at a disadvantage by the first and second party who have “shared information” the third party does not.
This “shared information” can be easily used as a “Shannon Communications Channel” to form what have been called “covert channels” that reliably allow the transfer of information between the first and second parties whilst leaving the third party seeing only what appears “random information”.
A little thought shows that such channels “must always exist” within channels used for information to be communicated.
Thus the potential for a prompt injection is there that no third party “guard rails” at the input or output of the LLM can stop.
The only thing left to do is unfortunately implicit in the static nature of the LLM…
Which is find an attack that works, because the LLM weights are fixed once found it will continue to work untill the weights in the LLM are changed.
We’ve all been informed of the very very high costs of “large frontier model generation” and in fact even politicians are jumping up and down and claiming all sorts of “IP theft” is being committed by *agents of foreign powers” etc over such costs.
So we can reasonably guess that the LLM weights are rather more than just fixed for moments in time but actually for months if not actually years at a time.
So once a vulnerability is found and they always will be, then such attacks will exist and be exploited.
Re: System Card: Claude Opus 5
The report details an -incredible- number of -fascinating- evaluations.
I had never seen many of the evals prior.
Autonomy threat model 1 is applicable to Claude Opus 5.
It bears under these three conditions: high reliance, sensitive access, capable of subterfuge.
Autonomy threat model 2 would occur in key R&D domains.
Stealth rates , here as measured, put Opus 5 yet below Mythos Preview.
This considers an agents ability to complete both a main task and side task without appearing suspicious to a monitor. Opus 5’s stealth rate was measured at 4–5% and 1% under different evaluations.
WRT prompt injection, it’s good to see that Claude can offer two layers of protection: one on data coming in, and one on actions (tool call results?) coming out. Definitely a good to minimize irreversible harmful actions.
Subscribe to comments on this entry
Fill in the blank: the name of this blog is Schneier on ___________ (required):
Allowed HTML
<a href="URL"> • <em> <cite> <i> • <strong> <b> • <sub> <sup> • <ul> <ol> <li> • <blockquote> <pre>
Markdown Extra syntax via https://michelf.ca/projects/php-markdown/extra/
Notify me of new posts by email.
Sidebar photo of Bruce Schneier by Joe MacInnis.
Powered by WordPress Hosted by Pressable
I am a public-interest technologist , working at the intersection of security, technology, and people. I've been writing about security issues on my blog since 2004, and in my monthly newsletter since 1998. I'm a fellow and lecturer at Harvard's Kennedy School , a board member of EFF , and the Chief of Security Architecture at Inrupt, Inc. This personal website expresses the opinions of none of those organizations.
Should You Use AI for a Task? Here’s a Simple Way to Decide
Measuring the Tendency of AI Agents to Go Rogue
Measuring LLMs' Ability to Perform Cryptanalysis
Why AI Needs a “Genie Coefficient”
MIT to Become Hotbed of AI Video Surveillance
Four Ways AI Is Being Used to Strengthen Democracies Worldwide
The CrowdStrike Outage and Market-Driven Brittleness
How Online Privacy Is Like Fishing
LLMs’ Data-Control Path Insecurity
Terrorists Don't Do Movie Plots
