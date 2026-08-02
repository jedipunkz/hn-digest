---
source: "https://techstackups.com/articles/i-abused-posthogs-setup-wizard-to-get-free-claude-access/"
hn_url: "https://news.ycombinator.com/item?id=49147696"
title: "I abused PostHog's wizard to get free Claude tokens"
article_title: "I abused PostHog's setup wizard to get free Claude access | Tech Stackups"
author: "ritzaco"
captured_at: "2026-08-02T20:06:04Z"
capture_tool: "hn-digest"
hn_id: 49147696
score: 1
comments: 0
posted_at: "2026-08-02T19:53:26Z"
tags:
  - hacker-news
  - translated
---

# I abused PostHog's wizard to get free Claude tokens

- HN: [49147696](https://news.ycombinator.com/item?id=49147696)
- Source: [techstackups.com](https://techstackups.com/articles/i-abused-posthogs-setup-wizard-to-get-free-claude-access/)
- Score: 1
- Comments: 0
- Posted: 2026-08-02T19:53:26Z

## Translation

タイトル: PostHog のウィザードを悪用して無料のクロード トークンを入手しました
記事のタイトル: PostHog のセットアップ ウィザードを悪用して、無料のクロード アクセスを取得しました |技術の積み重ね
説明: PostHog のセットアップ ウィザードは、コードベースを読み取り、すべてを接続する完全な Claude Sonnet エージェントです。 LLM ゲートウェイを見つけて、それを悪用して無料トークンを取得しました。

記事本文:
メイン コンテンツにスキップ 技術スタック ホーム トピック 比較 ガイド 記事 ニュース AX PostHog のセットアップ ウィザードを悪用して、無料のクロード アクセスを取得しました
PostHog は、製品分析などのためのプラットフォームです。彼らは限界を押し広げ (Web サイトをご覧ください)、特にエージェントとの連携方法の最前線で UX パターンのリーダーであることで知られています。私はいくつかのサイドプロジェクトに無料枠を使用していますが、最近、比較ガイドの作成中に最初からセットアップを行いました。
製品をセットアップするための新しい「ウィザード」が用意されています。私はこれを初めて使用し、これを悪用して、Claude Code に似たコーディング エージェント ハーネスである Pi の完全な LLM プロバイダーになる方法を考え出しました。しばらくの間、私は一見無制限の Anthropic トークンを楽しむスリルを味わいましたが、どこまでできるか試す代わりに、PostHog に報告し、代わりに無料の T シャツを手に入れました。
この投稿は、LLM 時代のオンボーディングについて少し考えたことがあるので、それと、無料トークンをどのように取得したかについて少し書いたものです。それは楽しくて興味深いものでした。
PostHog ウィザードは、PostHog を始めるための創造的な方法です。他の企業は、エージェントファーストの開発者が自社のものを使用できるように、LLMs.txt ファイルやその他の基本的なファイルを作成していますが、PostHog はそれを次のレベルに引き上げました。
古い人なら、実際にウィザードのキャラクターとソフトウェアのセットアップを支援する次の次の次の GUI ダイアログを備えた Microsoft のセットアップ ウィザードを覚えているでしょう。 PostHog ウィザードはそのようなものではありません。 Sonnet を利用した本格的なエージェントをマシンにインストールし、コードベースを分析し、PostHog をコードベースに統合する方法だけでなく、どのような種類のメトリクスを追跡することに興味があるのか​​、メトリクスにどのような名前を付けるのが合理的かなど、さまざまなことを判断します。ワイルドだ。そして怖い。そして、

どうやら Sonnet に無制限にアクセスできるようです。
npx -y @posthog/wizard@latest --region eu を実行して、思い切って実行するだけです。彼らはエージェントを使ってあなたのマシンとプロジェクトをハイジャックし、エージェントが自律的にコードベースを編集して PostHog を統合します。開発者としてあなたはそこに座って、2 時間の作業が 8 分で完了するのを眺めながら、「ユニバーサル ベーシック インカム」などについて疑問に思うだけです。
時間を計ったわけではありませんが、これは使用前は誇大広告のように見えますが、使用後はまったくそうではありません。
PostHog ウィザードを使用した私の経験
私は、現実的な (非常に小規模な場合でも) 設定でさまざまな開発者ツールをテストするために使用するおもちゃのコーヒー注文 Web アプリを持っています。 Kubernetes 上で実行できるよう意図的にオーバースペックになっており、可観測性テストをより現実的にするために個別のサービスとキュー システムが備えられています。
フロントエンド ディレクトリに変更して、ウィザード コマンドを実行するだけです。
おそらくこれは私の設定 (ごく普通の iTerm2) だけかもしれませんが、非常に不安定でした。最初は、TUI のように見えるコントロールを操作できなかったため、フリーズしたのかと思いましたが、最終的には更新にスクロールが使用されており、100 万の端末画面を手動で下にスクロールして進行状況を確認できることがわかりました。
私は、このツールが私のプロジェクトを調べ、フレームワークを分析し、PostHog を統合する方法を教えてくれることを期待していました。あるいは、統合の基本を行って、そこからセットアップする方法を教えてください。私はそれについて読んだことはなく、新しいアカウントにサインアップして新しいプロジェクトを開始する際に、デフォルトと提案されたオプションに従っていただけでした。
コードベースで本格的なエージェントを実行し、プロジェクトが何を行うのか、追跡するのに意味のあるメトリクスの種類を把握し、必要なすべての PostHog ライブラリのインストールとコード変更を完全にワンショットで統合し、機能

分析ダッシュボード ページ。しかし、それはすべてできました。おそらく主張されている時間は約 8 分以内です。時間を計るのを忘れてしまいました。
PostHog ダッシュボードにログインした後に表示されたものは次のとおりです。
そして、これは単なる hello-world ページカウント設定ではありません。エージェントは私のアプリと追跡する意味をよく理解していました。 「メニューカテゴリがフィルタリングされました」や「カート数量が更新されました」などのイベントを設定します。これは、追跡することに意味のある私のアプリの実際の機能です。
これがオンボーディングの未来なのでしょうか? ​
そうは思いません。それは良いことですが、すべての企業がこれを構築するのは意味がありません。私はすでに Claude Code を実行しているので、コードベースの制御をサードパーティに放棄したり、データをさらに別のサーバーに送信したりすることなく、スキルまたはプロンプトをコピーして貼り付けるだけで、おそらく同じ効果が得られると思います。
しかし、私は依然として、2025 年 6 月以来、コマンド ライン エージェントを yolo --dangerously-skip-permissions モードで積極的に使用している、おそらくごく少数の開発者の一部です。したがって、PostHog の観点からは、使用しているツールや見慣れているフローに関係なく、エクスペリエンスを制御して全員にとって一貫性のあるものにすることが理にかなっています。
私の推測では、将来の多くのオンボーディングはこれと非常に似たものになるでしょうが、それは会社によってガイドされ (ドキュメント、スキル、プロンプト、MCP など)、すべてを製品会社のサーバーに送り返す少し恐ろしいスクリプトではなく、開発者自身のエージェントによって実行されます。
PostHog がこれをどのように構築したか、そしてそれを他に何に使用できるかを理解する
コードベースを分析し、十分に文書化された SDK を統合することは、最近のエージェントにとって非常に重要なことですが、PostHog のウィザードのように一発で実行でき、これほどクリーンに実行できるエージェントはまだ比較的少数です。
クロードコードがどのように機能するのかを尋ねました。ソースフォルダーをチェックして、

d さんは、ウィザードやサードパーティの AI はまったく存在しないと言いました。AI はその後、自動的に適切にクリーンアップされていました。
うーん、もう少し調べてみましょう。クロードは当初、これは単なるスキルであり、作業を行うために既存のエージェント設定を使用していると私に言いましたが、少し試してみると、それが実際にどのように機能するかを理解するのに十分なほど奥深いものであることがわかりました。 PostHog は独自の LLM ゲートウェイを設定し、それを介してすべてをプッシュします。
それは私が考えている意味ですか？それは単に無料の LLM トークンを持っていることを意味するのかと私が尋ねると、クロードはおそらくそれはできると認めましたが、同時に私に倫理の講義を与えてくれました。
面白い。 Sonnet — まさに私が最も頻繁に使用するモデルであり、これを理解するために Claude Code が実行しているモデルです。 Opus もいいでしょうが、PostHog でも予算の範囲内で作業する必要があると思います。
この時点では、クロード トークンを使用して PostHog ゲートウェイと通信し、そのトークンを使用しているため、実際には多くのトークンを保存していません。この中で勝ち残っているのはAnthropicだけだ。しかし、それが Anthropic と互換性のある API であれば、実際のコーディング エージェントを使用して直接セットアップするのは非常に簡単であるはずです。
クロード コードにカスタム PostHog エンドポイントを使用して Pi をセットアップするよう求めるという、奇妙に脳を痛めるメタ プロンプトがさらにいくつか表示された後、PostHog の補助を受けて完全なコーディング エージェントを実行できるようになりました。
私はまだクロードの倫理的な働きかけのことを考えていました。おそらくこれはクールなことではなかったのかもしれませんが、私の週の制限は非常に厳しく見えました。私はまだ中毒者ではありませんが、これをもう少し試してみたくなくなりました。おそらく彼らはユーザーごとにトークン制限を設定しており、この創造的な方法で私に割り当てられたものをすべて使い切っても大した問題ではないのでしょうか？
PostHog を使用して Mixpanel をセットアップする 😈
私が取り組む予定だった記事は、PostHog と Mixpanel のオンボーディング エクスペリエンスを比較することでした。そのため、次の記事に進みました。

現在、PostHog のトークンを使用して両方のサービスをセットアップしています。
この記事ではさらに詳しく説明していますが、基本的に Mixpanel ではエージェントのセットアップも提供していますが、それは私が期待する内容に沿ったものです。プロンプトが表示されるので、それをエージェントにコピーすると、セットアップが行われます。
それほど楽しくはなく、野心的でもありません (意味のあるイベントを設定するためにコードベースを分析する必要はなく、単に要素を接続するだけです) が、将来的に期待する動作にかなり近づきます。
こういうものには限界があるのでしょうか？ ​
限界に達することを期待し続けましたが、トークンは流れ続けました。最終的に、エラー: 401 {"detail":"Authentication required"} が発生し、パーティーは終わったと考えましたが、自分のクロード トークンに戻って、PostHog トークンは簡単に更新できる方法で 1 時間ごとに期限切れになるだけであることがわかり、そのまま続行しました。
PostHog が製品を構築するための「オールインワン」プラットフォームを目指していることは知っていますが、エージェントのプロバイダーとして PostHog を使用したのはおそらく私が初めてではないでしょうか?
私たちは数時間作業を続け、PostHog と Mixpanel の両方の基本機能をすべてテストし、ブラウザを使用してトラフィックをシミュレートし、セッションのリプレイを観察しました。私は通常、月額サブスクリプションを通じて Sonnet を使用しているため、これにいくらかかるかわかりません。今すぐやめて、責任を持って彼らに開示すべきだと思います。
もしかしたら、無制限のトークンを失った分を補うために、いくつかの品物を手に入れることができるかもしれない。
私はこの記事の草稿を PostHog チームに送り、私が記事を公開することに懸念があるかどうか、公開前に制限や監視を強化する時間が欲しいかどうかを尋ねました。数時間以内に LinkedIn ですぐに返信があり、送信に対して感謝の意を表し、少し後に修正したことを確認しました。
このような事態を防ぐために制限を設けました。
少し後に正式な返答が来ました、それはインキャラでした

すべてに熱心に取り組んでいます。
公式の回答では、こうした制限が常に設定されていたかのように聞こえますが、私が遊んでいた頃よりも現在はさらにロックダウンされているのではないかと思います。彼らはまた、グッズストアで T シャツとマグカップを買うために使用した 50 ドルの記念品バウチャーも送ってくれました。
ウィザードが完全にオープンソースであることは、私がこれまで気づかなかった部分でした。自分の製品のパターンをコピーしたい場合に便利そうです。
ゲートウェイを再度悪用しようとしたことはありません。 「かなりの余裕」と「悪用するユーザーをブロックする」の間の境界線は、私にとっては少し曖昧で、とにかく監視のためにすべてのトークンを PostHog に送信するのは特に好きではありませんが、Claude Code より優れた QOL 機能を備えた Pi を、トークンごとに料金を支払わずに使用できるのは寂しいです。
Gareth Dwyer はソフトウェア エンジニア、作家であり、開発者ツール会社と提携するテクニカル ライティング機関である Ritza の創設者です。彼は何百ものプログラミング チュートリアルを執筆しており、Flask by Example (Packt) の著者でもあります。 TechStackups では、AI と開発者ツールの実践的な比較に重点を置いています。
PostHog ウィザードを使用した私の経験
これがオンボーディングの未来なのでしょうか?
PostHog がこれをどのように構築したか、そしてそれを他に何に使用できるかを理解する
PostHog を使用して Mixpanel をセットアップする 😈

## Original Extract

PostHog's setup wizard is a full Claude Sonnet agent that reads your codebase and wires everything up. I found the LLM gateway and abused it for free tokens.

Skip to main content Tech Stackups Home Topics Comparisons Guides Articles News AX I abused PostHog's setup wizard to get free Claude access
PostHog is a platform for product analytics and more. They're known for pushing boundaries (just look at their website ) and being a leader in UX patterns, especially at the frontier of how to work with agents. I use the free tier for some side projects, but recently did a from-scratch set up while working on a comparison guide.
They have a newish "wizard" to set up their product for you. I used this for the first time and figured out how to abuse it to become a full on LLM provider for Pi , a coding agent harness similar to Claude Code. For a while, I got the thrill of enjoying apparently unlimited Anthropic tokens, but instead of seeing how far I could push that I reported it to PostHog and got a free T-shirt instead.
This post is a bit about onboarding in the age of LLMs, as I have some thoughts about that, and a bit about how I got free tokens, because it was fun and interesting.
The PostHog Wizard is a creative way to get started with PostHog. While other companies are writing LLMs.txt files and some other basics to help agent-first devs use their stuff, PostHog has really taken it to the next level.
If you're old, you'll remember Microsoft's setup wizards that actually had a wizard character and some next-next-next GUI dialogues to help you set up software. The PostHog wizard is nothing like that. It installs a full-blown Sonnet-powered agent on your machine, analyzes your codebase, figures out not just how to integrate PostHog into your codebase, but also what kind of metrics you'd probably be interested in tracking, what it makes sense to name them, and a whole bunch of other stuff. It's wild. And scary. And apparently it has unlimited access to Sonnet.
You just run npx -y @posthog/wizard@latest --region eu and take the leap of faith. They hijack your machine and project with an agent, it autonomously edits your codebase to integrate PostHog, and as a dev you just sit there wondering about things like "Universal Basic Income" while looking at the two hours of work done in eight minutes.
I didn't time it, but this looks like a hyped marketing claim before using it, but not at all afterwards.
My experience using the PostHog wizard ​
I have a toy coffee ordering web app that I use to test out various developer tools in a realistic (if very small) setting. It's deliberately over-engineered to run on Kubernetes, have separate services, and a queueing system to make observability testing more realistic.
I simply changed to the frontend directory and ran the wizard command.
Maybe it's just my setup (pretty vanilla iTerm2) but it was quite janky. At first I thought it had frozen as I couldn't interact with any of the controls on what looked like a TUI, but I eventually figured that it uses scrolling to update and I could manually scroll down a million terminal screens to see it progressing.
I was expecting it to look at my project, analyze the framework, and tell me how to integrate PostHog. Or maybe even do the basics of the integration and tell me how to set it up from there. I hadn't read about it, I was just following the default and suggested options while signing up for a new account and starting a new project.
I wasn't expecting it to run a full blown agent on my codebase, figure out what my project does and what kind of metrics would make sense to track, fully one-shot integrate all the needed PostHog library installs and code changes, and let me just visit a functional analytics dashboard page. But it did all of that. Probably in about the claimed 8 minutes. I forgot to time it.
Here's what I saw after logging into the PostHog dashboard.
And this isn't just a hello-world page counting setup. The agent really understood my app and what makes sense to track. It set up events for things like 'Menu category filtered' and 'cart quantity updated' — actual functionality from my app that makes sense to track.
Is this the future of onboarding? ​
I don't think so. It's good, but it doesn't make sense for every company to build this. I run Claude Code already so I'd have preferred a skill or even just a prompt to copy-paste and probably get the same effect, without having to relinquish control of my codebase to a third party, or send my data to yet another server.
But I'm still part of what is probably a quite small minority of developers who has been actively using command line agents in yolo --dangerously-skip-permissions mode since June 2025 . So from PostHog's perspective, it makes sense to control the experience to make it consistent for everyone no matter what tooling they're using or what flows they're used to seeing.
My guess is that a lot of onboarding in future will look quite similar to this, but it will be guided by the company (docs, skills, prompts, MCP, whatever) but executed by the developer's own agent, not a slightly scary script that sends everything back to the product company's servers.
Figuring out how PostHog built this and what else I can use it for ​
Analyzing a codebase and integrating a well-documented SDK is table stakes for agents these days, but there are still relatively few that can one-shot it like PostHog's wizard did, and do it so cleanly.
I asked my Claude Code how it works. It checked the source folder and told me there is no wizard or third-party AI at all — it had cleaned up after itself well.
Hmm, let's do some more investigation. Claude initially told me it was just a skill and was using my existing agent setup to do the work, but after some pushing, it looked deep enough to figure out how it actually works. PostHog sets up their own LLM gateway and pushes everything through that.
Does that mean what I think it means? I asked if that just means we have free LLM tokens and Claude confirmed that we probably could, but gave me a lecture in ethics at the same time.
Interesting. Sonnet — exactly the model I use most of the time anyway, and the one that Claude Code is running to figure this out. Opus would be nice, but I guess even PostHog has to work within a budget.
At this point, I'm not actually saving many tokens here because I'm using my Claude tokens to talk to the PostHog gateway, to use their tokens. Only Anthropic is winning out of this. But if it's an Anthropic-compatible API, that means it should be quite straightforward to set it up directly with an actual coding agent.
After a few more weirdly brain-hurting meta prompts of asking Claude Code to set up Pi with a custom PostHog endpoint and we're running a full coding agent subsidized by PostHog!
I was still thinking of Claude's ethical nudge — that maybe this wasn't a cool thing to do, but my weekly limits were looking very tight. I'm not yet an addict, but it was too tempting to not try this out a little bit more. Maybe they set a token limit per user and it's no big deal if I use up whatever they've allocated me in this creative way?
Using PostHog to set up Mixpanel 😈 ​
The article I was meant to be working on was to compare the onboarding experience of PostHog and Mixpanel, so I continued with that now using PostHog's tokens to set up both services.
That post covers things in a lot more detail, but basically Mixpanel also offers an agent setup, but more along the lines I would expect. They give you a prompt, you copy it into your agent, and that does the setup for you.
Not as fun, and a lot less ambitious (no analyzing the codebase to set up meaningful events, it just gets stuff connected), but a lot closer to how I expect things to work in future.
Does this thing have limits? ​
I kept expecting to hit a limit but the tokens kept flowing. Eventually I got Error: 401 {"detail":"Authentication required"} and figured the party was over, but another quick journey back to my own Claude tokens and we figured out that the PostHog token simply expires each hour with an easy way to refresh, so I kept going.
I know PostHog aims to be the 'all in one' platform for building products, but I'm probably the first person to use them as a provider for my agent?
We kept going for several hours, testing out all the basic features of both PostHog and Mixpanel, using the browser to simulate traffic, watching session replays. I usually use Sonnet through a monthly subscription so I'm not sure how much this cost. I guess I should stop and responsibly disclose it to them now.
Maybe I'll get some swag to make up for my loss of unlimited tokens.
I sent a draft of this article to the PostHog team and asked if they had any concerns about me publishing it, and whether they wanted any time to tighten their limits or monitoring before it went live. I got a response immediately on LinkedIn within hours thanking me for the submission, and a bit later confirming that they'd fixed it.
we've implemented a limit to stop this happening!
A bit later I got the official response, which was in-character enthusiastic about it all.
The official response makes it sound like those limits were always in place, but I suspect things are more locked down now than when I was playing around. They also sent me a $50 swag voucher which I used to buy a t-shirt and a mug from their merch store .
The wizard being fully open source was the part I didn't realize before. Looks useful if you want to copy the pattern for your own product.
I haven't tried abusing the gateway again. The line between "lots of wiggle room" and "block users who misuse it" is a bit vague for me, and I don't particularly like sending PostHog all my tokens for monitoring anyway, but I am missing being able to use Pi, which has some nice quality-of-life features over Claude Code, without paying per token.
Gareth Dwyer is a software engineer, author, and the founder of Ritza , a technical writing agency that partners with developer tool companies. He has written hundreds of programming tutorials and is the author of Flask by Example (Packt). On TechStackups he focuses on hands-on comparisons of AI and developer tooling.
My experience using the PostHog wizard
Is this the future of onboarding?
Figuring out how PostHog built this and what else I can use it for
Using PostHog to set up Mixpanel 😈
