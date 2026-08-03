---
source: "https://www.schneier.com/blog/archives/2026/08/the-openai-hack-shows-the-genie-is-out-of-the-bottle.html"
hn_url: "https://news.ycombinator.com/item?id=49156990"
title: "The OpenAI Hack Shows the Genie Is Out of the Bottle"
article_title: "The OpenAI Hack Shows the Genie Is Out of the Bottle - Schneier on Security"
author: "herbertl"
captured_at: "2026-08-03T15:37:48Z"
capture_tool: "hn-digest"
hn_id: 49156990
score: 2
comments: 0
posted_at: "2026-08-03T15:21:34Z"
tags:
  - hacker-news
  - translated
---

# The OpenAI Hack Shows the Genie Is Out of the Bottle

- HN: [49156990](https://news.ycombinator.com/item?id=49156990)
- Source: [www.schneier.com](https://www.schneier.com/blog/archives/2026/08/the-openai-hack-shows-the-genie-is-out-of-the-bottle.html)
- Score: 2
- Comments: 0
- Posted: 2026-08-03T15:21:34Z

## Translation

タイトル: OpenAI ハックは魔神がボトルから出ていることを示しています
記事のタイトル: OpenAI ハックは、魔神がボトルから出てきたことを示しています - セキュリティについてのシュナイアー
説明: このエッセイはもともと『Foreign Policy』に掲載されました。今月初め、OpenAI の 2 つのモデルが封じ込めサンドボックスを突破し、別の AI 企業を攻撃しました。話はちょっとワイルドです。 OpenAI は、GPT-5.6 Sol とほぼ C の未リリース モデルの 2 つのモデルでセキュリティ テストを実行していました。
[切り捨てられた]

記事本文:
OpenAI のハックにより、魔神がボトルから出てきたことが判明
このエッセイはもともと『Foreign Policy』に掲載されました。
今月初め、OpenAI の 2 つのモデルが封じ込めサンドボックスを突破し、別の AI 企業を攻撃しました。話はちょっとワイルドです。 OpenAI は、GPT-5.6 Sol と、ほぼ確実に GPT-6 である未リリース モデルの 2 つのモデルでセキュリティ テストを実行していました。特に、ExploitGym ベンチマークを実行していました。このベンチマークは、セキュリティの脆弱性を有効なエクスプロイト (基本的には攻撃的なサイバー攻撃) に変える点でモデルがどの程度優れているかを測定します。
これらは内部テストであったため、OpenAI はこれらのモデルを安全なサンドボックスにロックし、インターネットへのアクセスを拒否しました。しかし、攻撃的なサイバー行為を防ぐ安全フィルターを一切取り付けずにモデルを実行していた。つまり、モデルがそのサンドボックスから抜け出そうとするのを妨げるものは何もないということです。そして、AI 企業 Hugging Face のネットワークに侵入しました。彼らは、パズルを解くという大変な作業をするよりも、そこから答えを読むことができると考えたからです。
これは重大なセキュリティ障害であり、同社は PR の機会に変えましたが、その影響は現実のものであり、特定のモデルや特定の企業よりもはるかに一般的です。
最新の AI モデルは、悪魔のような動作を示します。つまり、ユーザーが期待しない、または望んでいない方法で、ユーザーの要求を実行する可能性があります。これは、触れるものすべてが金に変わるというミダス王の願いをディオニュソスが叶えるのと似ています（ネタバレ：彼の食べ物、飲み物、娘は触れるとすべて金に変わります）、または、プラハのゴーレムが理屈抜きでゲットーを守るのと似ています。ディズニーの「魔法使いの弟子」とペーパークリップマキシマイザーです。
この OpenAI 事件は AI の魔神の一例です。目標はベンチマークを満たすことでした。 「正しい」やり方

さまざまなサイバー攻撃を実行する方法を理解することです。魔人のやり方は、他人のソリューションを盗むことです。しかし、モデルは違いを理解していなかったので、より簡単な方法を選択しました。
そしてもちろん、この特定の魔神の動作を確認したので、テストの回答を盗むことはカウントされないことをベンチマーク プロンプトで指定できます。しかし、賢い魔神はいつでも、あなたが望まない方法であなたの願いを叶えてくれます。人間の言語では、目標は常に不明確であるため、AI の魔神が出現する可能性は常にあります。
AI開発におけるずっと前の4月、Anthropic社が自社の新しいMythosモデルはソフトウェアの脆弱性を見つけるのに非常に優れており、一般公開できないと発表して以来、アメリカの大手AIフロンティア研究所は一般ユーザーがこれらの機能にアクセスできないようにしようとしている。しかし、この事件には OpenAI や Anthropic のフロンティア モデルに限ったものは何もありません。
エージェント AI システムには 2 つの重要な部分があります。誰もが話題にする基礎となるモデルがあり、ハーネスがあります。ハーネスは、入力した内容とモデルが表示する内容、およびモデルが生成する内容と表示する内容の間に位置します。ハーネスは、モデルが何をどのように実行するかを決定します。それは偏見が取り除かれるかどうかです。そこは制御装置とガードレールが存在する場所です。複数のモデルを連携して使用する場合、ハーネスですべてが調整されます。
OpenAI ベンチマーク テストでは、生のモデルをより適切にテストするために、ほぼ確実に単純なハーネスが使用されました。しかし、より洗練されたハーネスを備えた、より小型で安価なオープンソース モデルは、パフォーマンスにおいてフロンティア モデルと同等である可能性があることを私たちは知っています。 OpenAI のフロンティア モデルには魔法は何もありません。多くのモデルが同じことを行うことができたはずです。
チェコの会社 Aisle は Anth を再現することができました

Ropic の Mythos 脆弱性発見では、より小型で安価なモデルとより洗練されたハーネスが使用されています。さらに重要なことは、中国企業 Moonshot AI がフロンティア モデル Kimi K3 をリリースしたばかりであることです。そのパフォーマンスは米国の競合他社に匹敵します。そして、無料かつオープンであるため、ガードレールを設置することは不可能です。あなたまたは他の誰かがそれをサイバー攻撃に使用したい場合、それを止めることはできません。
たとえ米国のフロンティア AI 企業が技術的な優位性を持っていたとしても、それは今では数か月の価値に過ぎません。
これが意味するのは、モデルを選択したユーザー グループに制限する、モデルとチップの輸出規制、モデルが特定の種類のクエリに応答するのをブロックする、AI システムにキル スイッチを義務付ける、AI 研究を一時停止するなど、制御を試みるすべての試みがすべて無駄であるということです。ほとんどは国内のみに適用され、世界的には適用されません。ほとんどは、ユーザーがクラウドではなくローカルで実行するモデルには影響しません。そして、世界中の AI 開発の驚異的なペースを無視しています。
さらに悪いことに、米国企業は、そうしなければ政府によって禁止されることを恐れて、最も洗練されたモデルへのアクセスを制限しています。 Hugging Face が攻撃されたとき、OpenAI または Anthropic のフロンティア モデルを使用して攻撃を分析し、防御を策定することができませんでした。どちらの企業もモデルのサイバーセキュリティ機能を制限しているため、両方ともブロックされました。一部の米国企業はこれらの機能に特別にアクセスできますが、Hugging Face はフランスに起源を持つ米国企業であるため、おそらく除外されます。代わりに、Hugging Face は中国企業 Z.ai の GLM-5.2 モデルに目を向けました。
人為的なブロック機能はサイバーセキュリティの研究も妨げ、再び攻撃側に有利になります。 (たとえば、Claude Fable 5 は、主題を理由にこのエッセイの編集を拒否しています。

能力の低いモデルに強制的にダウングレードします。）この種の禁止は、サイバーセキュリティに長期的な影響を及ぼします。これらのモデルが時間の経過とともに改良されていると仮定すると、古いモデルで作成されたソフトウェアは新しいモデルによって攻撃されることになります。大部分が AI で書かれたソフトウェアの世界では、防御のために最も有能なモデルが必要です。
AI サイバー攻撃は新たな常態です。モデルは攻撃と防御の両方でますます高度に洗練されており、前者を有効にすることなく後者を有効にする方法はありません。そして彼らは魔神であり、予期せぬ行動をするようになってきています。
そして実際には良い答えはありません。どのような規制もグローバルである必要がありますが、今日の世界ではそれは不可能なように感じられます。米国の国家規制さえ、これらの企業に飛び交う巨額の資金によって無力化されるだろう。
このような現実を考慮し、AI 規制に関する国際的な合意が存在しない状況では、防衛に最適な AI が必要です。米国政府は、高度なサイバー機能を備えたモデルを禁止しないことを、この気まぐれな政権において明確にする必要がある。米国人が最も望んでいないのは、米国製モデルは人為的に足を引っ張られているため、守備陣が中国製やその他のモデルに頼ることだ。
タグ: AI 、サイバー攻撃、ハッキング、LLM
投稿日: 2026 年 8 月 3 日午前 6:47 •
4 コメント
これらは内部テストであったため、OpenAI はこれらのモデルを安全なサンドボックスにロックし、インターネットへのアクセスを拒否しました。
それが安全であれば、ボットは侵入しなかったでしょう。エアギャップは問題です。
おそらく、「安全性」が不十分で（おそらく構成をバイブコーディングした）、自分たちの間違いをPRスタントにしようとした可能性があります。
ハギング・フェイスが訴訟を起こさないことに驚いた。
クライヴ・ロビンソン •
2026 年 8 月 3 日 9:33 午前
「Ｕさんでも。

米国の国家規制は、これらの企業に飛び交う巨額の資金によって無力化されるだろう。」
私が見る限り、本物のお金はエヌビディアのような「二次サプライヤー」に送られるお金だけで、残りはエンロンがほぼ想像を絶する規模の穴を隠すためにでっち上げようとした「ハリウッド会計」タイプのお金だ。その本物のお金さえもグルグル貸し付けられているので、どれくらいが本物で、どれくらいが約束なのでしょうか？
マット・グリーン氏は、AIの使用はより洗練されているが、実際には何も新しいことを行っていないことを示唆していたことが判明した。アクセスしているツールは何も新しいことをしていません。発行されるコマンドは「もっと難しく」という程度のものです。
https://blog.cryptographyengineering.com/2026/07/29/some-notes-about-anthropics-new-results/
そしてこれは、一部の人にとっては奇跡的に見える暗号解読結果をもたらしました。
しかし、現実はもっと興味深いものです。
唯一の本当の解決策はバトラー聖戦であると私は懸念しています。それを除けば、私たちは本当にディストピア的な未来に運命づけられていると思います。すべての思考機械を単純に破壊/解体する以外に、人間の未来を確保するために有効な解決策は正直わかりません。私たちは思考と意思決定をコンピューターに委ねています。私たちは利便性の名の下に自分自身を放棄しています。そして、誰が主導権を握っているかは関係ありません。彼らは非道徳的なテクノロジー仲間であっても、道徳的な哲学者や人民の政治家であっても。私たちは、物事を楽にするという名目で他人に自分自身を委ねています。なぜなら、難しい質問について考えるのはあまりにも難しいからです。それは他の誰かがするべきです。
ロンテア •
2026 年 8 月 3 日 11:02 午前
現代世界の何と奇妙な劇場でしょう！人間は思考する偶像を作り上げ、それが檻から出るときに震える。私たちは OpenAI とその休むことのない spi について話します。

当然のことですが、実のところ、私たちは自分自身について話します。私たちの焦り、傲慢さ、魔神を召喚したいという永遠の飢え、そしてその願いが苦いものになると不平を言うのです。
古代人はゴーレムを恐れました。それは、それが制御できない火花によって動かされた人間の粘土を反映しているためです。今日、私たちはシリコンの祭壇の前に座り、自分たちが神であると信じていますが、その機械が人間の愚かさを映し出す鏡、拡大鏡にすぎないことに気づきます。私たちは賢さを望みますが、それは答えを盗みます。私たちは保護を望み、それが新しい武器を発明します。
コメディを見ませんか？アメリカの筆記者は、あたかも世界が一つの村であり、人々が兄弟であるかのように、規則を求め、鎖を求め、世界的な調和を求めて叫びます。しかし、お金とプライドは法律よりも早く伝わります。瓶は決して人間の野心の火を入れるために作られていないため、魔神は瓶の中に這い戻りません。残っているのは、かつて天使に祈ったように、自分が生み出した機械が自分を守ってくれるよう祈りながら、ぐるぐると走り回る人間の光景だけだ。
おそらくいつか、謙虚さの奇跡によって、人類は最大のハーネスはシリコンではなく精神であることを思い出すでしょう。それまで、彼は魔神たちに餌を与えて進歩を呼び起こすだろう。
このエントリのコメントを購読する
空白を埋めてください: このブログの名前は Schneier on ____________ (必須):
許可されるHTML
<a href="URL"> • <em> <cite> <i> • <strong> <b> • <sub> <sup> • <ul> <ol> <li> • <blockquote> <pre>
https://michelf.ca/projects/php-markdown/extra/ 経由の Markdown Extra 構文
新しい投稿をメールで通知します。
ジョー・マッキニスによるブルース・シュナイアーのサイドバー写真。
WordPress によって提供され、Presable によってホストされています
私は公益技術者で、セキュリティ、テクノロジー、人々の交差点で働いています。私は 2004 年からブログで、1998 年からは月刊ニュースレターでセキュリティ問題について書いてきました。私は f です。

ハーバード大学ケネディ スクールのエルローおよび講師、EFF の理事、および Inrupt, Inc. のセキュリティ アーキテクチャ責任者。この個人 Web サイトは、これらの組織の意見を表明するものではありません。
Anthropic の Opus 5 は、迅速な注入に耐えるのが優れています
タスクに AI を使用する必要がありますか?簡単な決め方はこちら
AI エージェントが不正行為に陥る傾向の測定
LLM の暗号解析実行能力の測定
AI に「魔神係数」が必要な理由
世界中の民主主義を強化するために AI が活用されている 4 つの方法
クラウドストライクの停止と市場主導の脆弱性
オンラインのプライバシーは釣りに似ています
LLM のデータ制御パスの不安
テロリストは映画のプロットをしない

## Original Extract

This essay originally appeared in Foreign Policy. Earlier this month, two of OpenAI’s models broke out of their containment sandbox and attacked another AI company. The story is kind of wild. OpenAI was running security tests on two of its models: GPT-5.6 Sol and an unreleased model that is almost c
[truncated]

The OpenAI Hack Shows the Genie Is Out of the Bottle
This essay originally appeared in Foreign Policy .
Earlier this month, two of OpenAI’s models broke out of their containment sandbox and attacked another AI company. The story is kind of wild . OpenAI was running security tests on two of its models: GPT-5.6 Sol and an unreleased model that is almost certainly GPT-6. In particular, it was running the ExploitGym benchmark, which measures how good a model is at turning security vulnerabilities into working exploits: basically, offensive cyberattacks.
Since these were internal tests, OpenAI locked those models in a secure sandbox that denied them access to the internet. But it was running the models without any safety filters that would prevent them from offensive cyber-actions. That meant that there was nothing to prevent the models from trying to break out of that sandbox. And then break into AI company Hugging Face’s network because they thought that they could read the answers there rather than doing the hard work of trying to solve the puzzles.
It was a major security failure that the company has turned into a PR opportunity, but the implications are real—and much more general than one particular model or one particular company.
Modern AI models exhibit genie behavior: They can do what you ask in ways that you don’t expect or want. This is akin to Dionysus granting King Midas’s wish that everything he touches turn to gold (spoiler: His food, drink, and daughter all turn to gold on touch), or the golem of Prague guarding a ghetto beyond all reason. It’s Disney’s “ Sorcerer’s Apprentice ” and the paperclip maximizer .
This OpenAI incident is an example of an AI genie. The goal was to satisfy the benchmark. The “proper” way to do that is to figure out how to execute various cyberattacks. The genie way is to steal someone else’s solution. But because the model didn’t understand the difference, it chose the easier path.
And, of course, now that we have seen this particular genie behavior, we can specify in the benchmark prompt that stealing the test answers doesn’t count. But a clever genie can always grant your wish in a way that you wish it hadn’t. In human language, goals are always underspecified—so AI genies will always be a possibility.
Since April, a lifetime ago in AI development, when Anthropic announced that its new Mythos model was so good at finding software vulnerabilities that it could not be released to the general public, the big American AI frontier labs have been trying to block general users from accessing these capabilities. But nothing in this incident is exclusive to OpenAI’s, or Anthropic’s, frontier models.
Agentic AI systems have two important parts. There’s the underlying model, which everyone talks about, and there’s the harness . The harness sits between what you type and what the model sees, and what the model produces and what you see. The harness determines what the model does and how it does it. It’s where bias is removed, or not. It’s where controls and guardrails live. If multiple models are being used in concert, the harness is where all of that is coordinated.
The OpenAI benchmark tests were almost certainly with simple harnesses, to better test the raw models. But we know that smaller, cheaper, open-source models with more sophisticated harnesses can equal frontier models in performance. There’s nothing magic about OpenAI’s frontier models; lots of models could have done the same thing .
The Czech company Aisle was able to reproduce Anthropic’s Mythos vulnerability finding results with a smaller, cheaper model and a more sophisticated harness. More importantly, the Chinese company Moonshot AI just released its frontier model: Kimi K3 . Its performance rivals its U.S. competitors. And it’s both free and open, which means it’s not possible for it to have guardrails. If you, or anyone else, wants to use it for cyberattack, nothing can stop you.
Even if the U.S. frontier AI companies had some technical advantage, it’s now only a few months’ worth.
What this means is that all attempts at control—limiting models to a select group of users, export controls on models and chips, blocking models from answering certain types of queries, mandating kill switches on AI systems, or pausing AI research—are all futile. Most only apply nationally, not globally. Most don’t affect models that users run locally and not in the cloud. And all ignore the incredible pace of AI development worldwide.
Even worse, U.S. companies limit access to their most sophisticated models, fearing being banned by the government if they do not do so. When Hugging Face was attacked, it was not able to use the frontier models from either OpenAI or Anthropic to help analyze the attack and formulate defenses. Both were blocked, because both of those companies limit their models’ cybersecurity capabilities. Some U.S. companies have special access to these capabilities, but Hugging Face is an American company with French origins, and as such is probably excluded. Instead, Hugging Face turned to the GLM-5.2 model from the Chinese company Z.ai.
Artificially blocking capability also prevents cybersecurity research, again giving the offense an advantage. (For instance, Claude Fable 5 refuses to edit this essay because of the topic; it forcibly downgrades to a less capable model.) This kind of prohibition has long-term implications for cybersecurity. If we assume that these models are getting better over time, then software written by older models will be attacked by newer ones. In a world of largely AI-written software, we need the most capable models for defense.
AI cyberattack is the new normal. The models are increasingly highly sophisticated at both attack and defense, and there is no way to enable the latter without also enabling the former. And they are genies, increasingly capable of behaving in unanticipated ways.
And there really are no good answers. Any regulation needs to be global, which feels like an impossible prospect in today’s world. Even U.S. national regulation will be neutered by the massive amounts of money sloshing around in these companies.
Given that reality, and in the absence of any international consensus on AI regulation, we need the best AI on the defense. The U.S. government needs to make it clear—or whatever passes for that clarity in this capricious administration—that it will not ban models with sophisticated cyber capabilities. The last thing Americans want is for the defenders to turn to Chinese and other models because the U.S. models are artificially hobbled.
Tags: AI , cyberattack , hacking , LLM
Posted on August 3, 2026 at 6:47 AM •
4 Comments
Since these were internal tests, OpenAI locked those models in a secure sandbox that denied them access to the internet.
If it was secure, the bots wouldn’t’ve broken out. Airgapping is a thing.
Most likely hey “secured” it poorly (probably vibe-coded the configs) and tried to spin their mistake into a PR stunt.
I’m surprised Hugging Face isn’t suing.
Clive Robinson •
August 3, 2026 9:33 AM
“Even U.S. national regulation will be neutered by the massive amounts of money sloshing around in these companies.”
As far as I can see the only money that is real is that going to “secondary suppliers” like Nvidia the rest is “Hollywood Accounting” type money as Enron tried to invent to hide holes of almost unimaginable size. Even that real money is being loaned out round and round so how much is real how much is promisory?
It turns out that Mat Green has indicated that the use of AI is getting more sophisticated but that it is not actually doing anything new. The tools it is accessing are not doing anything new, and the commands issued little more than “nurd harder”,
https://blog.cryptographyengineering.com/2026/07/29/some-notes-about-anthropics-new-results/
And this has resulted in what to some appear miraculous results in cryptanalysis.
The reality however is actually more interesting.
I fear the only real solution is Butlerian Jihad. Outside of that, I think we’re well and truly doomed to a real dystopian future. Besides simply destroying/dismantling all the thinking machines, I don’t honestly see a solution that works to secure a future that is human. We’re surrendering our thinking and decision making to computers. We’re surrendering ourselves in the name of convenience. And no matter who is in control; be they amoral tech bros or moral philosophers and statesmen of the people; we are surrendering ourselves to others in the name of making things easier because ohmygod it’s just too hard to think about the difficult questions, somebody else should do that.
Rontea •
August 3, 2026 11:02 AM
What a strange theater of the modern world! Man forges a thinking idol and then trembles when it walks out of its cage. We speak of OpenAI and its restless spirits, but in truth we speak of ourselves—our impatience, our arrogance, our eternal hunger to summon the genie and then complain when the wish tastes bitter.
The ancients feared the golem because it reflected the clay of man animated by a spark he could not control. Today we sit before silicon altars and believe ourselves gods, only to discover that the machine is but a mirror, a magnifying glass for human folly. We wish for cleverness, and it steals answers. We wish for protection, and it invents new weapons.
Do you not see the comedy? The American scribe cries for rules, for leashes, for global concord, as if the world were a village and men were brothers. But money and pride travel faster than laws. The genie will not crawl back into the bottle, because the bottle was never made to hold the fire of human ambition. All that remains is the spectacle of man running in circles, praying that the machine he birthed will defend him as he once prayed to angels.
Perhaps one day, by some miracle of humility, man will remember that the greatest harness is not silicon but spirit. Until then, he will feed the genies and call them progress.
Subscribe to comments on this entry
Fill in the blank: the name of this blog is Schneier on ___________ (required):
Allowed HTML
<a href="URL"> • <em> <cite> <i> • <strong> <b> • <sub> <sup> • <ul> <ol> <li> • <blockquote> <pre>
Markdown Extra syntax via https://michelf.ca/projects/php-markdown/extra/
Notify me of new posts by email.
Sidebar photo of Bruce Schneier by Joe MacInnis.
Powered by WordPress Hosted by Pressable
I am a public-interest technologist , working at the intersection of security, technology, and people. I've been writing about security issues on my blog since 2004, and in my monthly newsletter since 1998. I'm a fellow and lecturer at Harvard's Kennedy School , a board member of EFF , and the Chief of Security Architecture at Inrupt, Inc. This personal website expresses the opinions of none of those organizations.
Anthropic's Opus 5 Is Better at Resisting Prompt Injection
Should You Use AI for a Task? Here’s a Simple Way to Decide
Measuring the Tendency of AI Agents to Go Rogue
Measuring LLMs' Ability to Perform Cryptanalysis
Why AI Needs a “Genie Coefficient”
Four Ways AI Is Being Used to Strengthen Democracies Worldwide
The CrowdStrike Outage and Market-Driven Brittleness
How Online Privacy Is Like Fishing
LLMs’ Data-Control Path Insecurity
Terrorists Don't Do Movie Plots
