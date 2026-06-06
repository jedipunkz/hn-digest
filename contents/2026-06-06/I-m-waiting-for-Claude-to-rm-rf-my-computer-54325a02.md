---
source: "https://12gramsofcarbon.com/p/agentics-local-coding-agents-are"
hn_url: "https://news.ycombinator.com/item?id=48426730"
title: "I'm waiting for Claude to rm rf my computer"
article_title: "Agentics: local coding agents are inherently unsafe"
author: "theahura"
captured_at: "2026-06-06T18:32:41Z"
capture_tool: "hn-digest"
hn_id: 48426730
score: 3
comments: 1
posted_at: "2026-06-06T16:52:24Z"
tags:
  - hacker-news
  - translated
---

# I'm waiting for Claude to rm rf my computer

- HN: [48426730](https://news.ycombinator.com/item?id=48426730)
- Source: [12gramsofcarbon.com](https://12gramsofcarbon.com/p/agentics-local-coding-agents-are)
- Score: 3
- Comments: 1
- Posted: 2026-06-06T16:52:24Z

## Translation

タイトル: クロードがコンピューターに rm 送信するのを待っています
記事のタイトル: Agentics: ローカル コーディング エージェントは本質的に安全ではありません
説明: 簡単、強力、安全。 2つ選んでください。

記事本文:
エージェントティクス: ローカルコーディングエージェントは本質的に安全ではありません
炭素12グラム
Agentics: ローカル コーディング エージェントは本質的に安全ではありません
簡単、強力、安全。 2つ選んでください。
theahura 2026年6月6日 2 シェア 告白したいことがあります。
クロード コード 1 を YOLO モードで実行します。
そうすべきではないことはわかっています。それがひどい考えであることは承知しています。しかし、クロードがファイルを読みたがっているので、そこに座って 30 秒ごとに Enter キーを押すのは耐えられません。
プロンプトに対して何も考えずに 1,000 万回目も Enter キーを押した後、私はキレました。
クロードを claude —dangerously-skip-permissions にエイリアスしました
そして、エージェントの子守をするという重荷からは解放されましたが、新たな恐怖が生まれました。
ある日、いつになるかわかりませんが、クロードが私のコンピュータ全体を rm -rf するでしょう。
以前、私はコーディング エージェントは、本当に賢い、まったく訓練を受けていないインターンによく似ていると述べました。私は今でも基本的にそう信じています。モデルは時間の経過とともに改善されてきましたが、あなた、あなたの会社、またはエンジニアリングの実践に関するコンテキストがなくても、依然としてすべてのセッションで完全に新鮮な状態で提供されます。新入社員のインターンに、ローカル マシンへの無制限のアクセスを許可しますか?ご存知のように、すべての SSH キー、契約書、大量の個人データ、そしておそらく免許証/クレジット カード/パスポートの写真はどこにありますか?
『12 Grams of Carbon』は読者サポートの出版物です。新しい投稿を受け取り、私の仕事をサポートするには、無料または有料の購読者になることを検討してください。
基本的な緊張は、高度に自律性があり、セットアップが簡単で、本当に安全なエージェントが必要であるということです。実際には 2 つしか選択できません。
フロンティア コーディング エージェント (Claude Code、Codex) をローカルにインストールし、すべての MCP をセットアップできます。ローカル開発がハッピー パスであるため、非常に簡単です。そうすれば、エージェントはいつかあなたの電子メールの受信箱を破壊するでしょう。
ハイグをインストールできます

権限ゲート型エージェントをローカルで制御し、すべての簡単なセットアップを完了したら、権限プロンプトで一度に 1 つずつ Enter キーを押して 1 日を費やします。
または、1 か月かけて分離されたサンドボックス サービスを構築し、バックグラウンド エージェントとしてクロード コードを実行することもできます。エージェントは機能しますが、Docker ネットワーク構成に深く入り込みすぎて、実際に楽しむことができません。
大多数の人は、埋め込みスペースの [強力、簡単、安全ではない] 部分にいます。ほとんどの場合、選択によるものです。Claude Code の平均的な消費者は、おそらく安全なサンドボックスをセットアップする方法を理解できず、限定されたエージェントを使用したくありません。
「はい、もう聞かないでください」の回数を最小限に抑えながら、ローカル エクスペリエンスをもう少し安全にするためにできることがいくつかあります。
/sandbox をオンにします。これは、OS レベルの分離 (Linux ではバブルラップ、macOS ではシートベルト) で、書き込みを作業ディレクトリに制限し、ホワイトリスト プロキシを介してネットワークをルーティングします。
決して触れられたくないもの (~/.ssh、~/.aws/credentials、.env、id_rsa、.npmrc、~/.netrc) に対する明示的な拒否ルールを設定に追加します。
適切なスキルとエージェント構成を作成し、会話が長回しにならないようにしてください。
Anthropic は、この不満に対処することを目的としたパーミッション モード全体を出荷しました。
Claude Code ユーザーは、許可プロンプトの 93% を承認しています。一部の意思決定を自動化する分類器を構築し、安全性を高めながら承認の疲労を軽減しました。
しかし、これらは常に絆創膏になります。コーディング エージェントの基本的な緊張は、コーディング エージェントに物事へのアクセスを与えたいということです。最も基本的なタスクを実行するには、ファイルの書き込みと読み取りができる必要があります。残念ながら、それだけでほとんどのガードレールを突破するのに十分です。ファイルシステムはエージェント構成の中核部分であるため、次のように想定する必要があります。

ディスク上にあるものは何でも大丈夫です。
ケーキを作って食べる唯一の実際の方法は、サンドボックスルートを行くことだと思います。理想的には、ローカル ファイルから遠く離れたクラウド内のサンドボックスで、中には貴重なものが入っていないため、ためらうことなく破棄できるボックスです。はい、バックグラウンド エージェントのセットアップは面倒ですが、インフラの実装方法が理解されるにつれて、時間の経過とともに簡単になるでしょう。そして、セキュリティだけではありません。クラウド サンドボックス上に存在するエージェントには、さまざまなインターフェイスからリモートでアクセスでき、バグ レポートや顧客からの受信などの自動イベントによってトリガーできます。
エージェント環境では、ベース マシンとして Firecrackers マイクロ VM プロバイダーを使用しました。 2 これらは開始と終了が非常に簡単であるため、常に更新されるエージェント ランタイムの構築に最適です。 VM 自体には実際には状態がありません。代わりに、各 VM には、関連するすべてのリポジトリのクローンを作成し、すべての依存関係をセットアップするセットアップ スクリプトが含まれているため、各エージェント セッションはクリーン ボックスで実行されます。基本的にエージェントはマシンを破壊することができますが、それは問題ではありません。私たちはそれをリサイクルするだけです。
実際のキーは、エージェントが実行されている場所と同じボックスに保管されることはありません。代わりに、エージェント ランタイムと外部世界の間でプロキシを行います。送信されるものに特定の資格情報のプレースホルダーが含まれているのを確認すると、実際の資格情報を注入する必要があるかどうかを決定する条件をトリガーできます。その結果、エージェントはすべてにアクセスできると思っていますが、実際には厳しく制御されています。
私がクラウド エージェントの利点を口にするのはこれが初めてではありません。これらについては、AI の実現に関する投稿でも書きました。
マネージド エージェント ランタイム (マネージド エージェント ランタイムとも呼ばれる) の市場には大きなギャップがあると思います。

クラウド上のバックグラウンドで実行できるため、「バックグラウンド エージェント」と呼ばれます)。それ以外のすべて、たとえば組織レベルのスキルは単なる絆創膏です。以前、コーディング エージェントにとってファイル システムは構成であると述べました。つまり、一貫した (そして安全な) エクスペリエンスを得るには、ファイルシステム全体を制御する必要があります。組織内の全員が AI の使用方法だけでなく、CLI の使用方法、Bash の使用方法、git の使用方法、コーディング エージェントを軌道に乗せるのに必要なその他の何百万ものツールの使用方法も学習すると想定するだけでは十分ではありません。そして、それでも十分ではないかもしれません。
ほとんどのチームはフルマネージド サービスを選択するべきだと思います。デビンは長い間宇宙に滞在しています。 Twill もこのようなことを行う比較的新人だと思いますか?しかし、実際に Ramp-Inspect / Stripe-Minions モデルに従っているサービスは他にたくさんありません。そのため、私たち自身がこの問題に遭遇したとき、世にある他のオプションのほとんどに満足できず、独自の を構築することにしました。現在、これはカスタマイズの柔軟性が高く、既製のオプションとして販売されています。
このようなツールがあることは非常に素晴らしいことだと言わざるを得ません。私たちがそれを構築できて本当によかったと思います。当社の PR の 30% 以上はすべて Slack を通じて出荷されており、バグのトリアージからニュースレターの作成まで、あらゆる種類の非常に優れた自動化機能を備えています。しかし、そのすべてには多大なメンテナンスが必要であり、私たちは文字通り一日中、この 1 つのことをどのように素晴らしいものにするか考えることだけに費やしています。代わりにヘルスケアや製造などを組み込んでいた場合、それを行うことは不可能でしょう。
これらが企業にとってどれほど大きな変革をもたらすかを説明するのは難しいですが、一度導入するともう元には戻れません。私の最良の比喩は、優れたバックグラウンド エージェントのセットアップは、はるかに優れたものであるということです。

端末内のツールを操作するよりも、非常に有能なエンジニアを雇用する方が良いでしょう。人間のようにこれらのものと対話することは非常に自然かつ直感的であり、これにより多くの困難な学習曲線が短縮され、多くのコラボレーションが可能になります。
そしてもちろん、それらは安全です。現地エージェントを運営したり、新入社員のインターンに個人のラップトップへのアクセスを許可したりするよりも、はるかに安全です。
『12 Grams of Carbon』は読者サポートの出版物です。新しい投稿を受け取り、私の仕事をサポートするには、無料または有料の購読者になることを検討してください。
現在は fly.io を使用していますが、モーダルへの切り替えを検討しています。まだ実行していない唯一の理由は、gvisor では Docker のサポートがそれほど簡単ではないからです
2 共有 この投稿に関するディスカッション コメント 再スタック トップ 最新のディスカッション 投稿はありません

## Original Extract

Easy, powerful, secure. Pick two.

Agentics: local coding agents are inherently unsafe
12 Grams of Carbon
Subscribe Sign in Agentics: local coding agents are inherently unsafe
Easy, powerful, secure. Pick two.
theahura Jun 06, 2026 2 Share I have a confession to make.
I run Claude Code 1 in YOLO mode.
I know I shouldn’t. I know it’s a terrible idea. But I just cannot stand sitting there and hitting enter every 30 seconds because Claude wants to read a file.
After mindlessly hitting enter for the ten millionth time on some prompt, I snapped.
I aliased Claude to claude —dangerously-skip-permissions
And even though I’m now free from the burden of having to babysit the agent, I have a new fear.
One day, I don’t know when, Claude will rm -rf my entire computer.
In the past, I’ve said that coding agents are a lot like a really smart, totally untrained intern. I still basically believe that. Even though models have gotten better over time, they still come to every session completely fresh without any context on you, your company, or your engineering practices. Would you give a new hire intern unlimited access to your local machine? You know, where you have all your ssh keys and your contracts and a ton of personal data and maybe photos of your license/credit card/passport?
12 Grams of Carbon is a reader-supported publication. To receive new posts and support my work, consider becoming a free or paid subscriber.
The fundamental tension is that we want agents that are highly autonomous, easy to set up, and really secure; and in practice, we can only pick two.
You can install a frontier coding agent (Claude Code, Codex) locally and set up all the MCPs and it will be super easy because local dev is the happy path, and then the agent will one day nuke your email inbox.
You can install a highly controlled permission-gated agent locally, get all the easy set up, and then spend your whole day hitting enter on permissions prompts one at a time.
Or you can take a month to build out an isolated sandbox service to run Claude Code as a background agent. The agent works, but you’re too deep in docker networking configs to really enjoy it.
The vast majority of people are in the [powerful, easy, insecure] part of the embedding space. Mostly by choice — the average consumer of Claude Code probably couldn’t figure out how to set up a secure sandbox and definitely doesn’t want to use a limited agent.
There are some things that you can do to make the local experience slightly more secure while minimizing the amount of “Yes, and don’t ask me again” you need to do.
Turn on /sandbox . It's OS-level isolation (bubblewrap on Linux, Seatbelt on macOS) that confines writes to your working directory and routes network through an allowlist proxy.
Add explicit deny rules for the stuff you never want touched — ~/.ssh, ~/.aws/credentials, .env, id_rsa, .npmrc, ~/.netrc — in your settings
Write good skills and agent configs, and don’t let your conversations go through too many turns.
Anthropic even shipped an entire permissions mode that is meant to address this complaint.
Claude Code users approve 93% of permission prompts. We built classifiers to automate some decisions, increasing safety while reducing approval fatigue.
But these will always be bandaids. The fundamental tension of a coding agent is that you want to give it access to things. It needs to be able to write and read files in order to do the most basic tasks. Unfortunately, that alone is enough to break out of most guardrails. Your filesystem is a core part of the agent config, so you should assume that anything on your disk is fair game.
I think the only actual way to have your cake and eat it too is to go the sandbox route. Ideally sandboxes in the cloud, far away from your local files, on boxes that you can tear down without hesitation because they don’t have anything precious in them. Yes, setting up background agents is a pain, but that will get easier over time as people figure out how to implement the infra. And you get more than security — agents that live on cloud sandboxes can be accessed remotely from a ton of different interfaces, and can be triggered by automatic events like bug reports or customer inbound.
For our agent environments, we went with a firecrackers micro VM provider for the base machine. 2 These are extremely easy to start and kill, which makes them perfect for building agent runtimes that are constantly refreshing themselves. The VMs themselves don’t really have state. Instead, each VM has a setup script that clones all the relevant repos and sets up all the dependencies, so each agent session runs on a clean box. The agent can basically destroy the machine and it won’t matter, we’ll just recycle the thing.
The actual keys are never kept on the same box as where the agent is running. Instead, we proxy between the agent runtime and the outside world. Whenever we see certain credential placeholders in anything going outbound, we can trigger conditionals that determine whether we should inject the real credentials. The end result is that the agent thinks it has access to everything, but in practice it’s tightly controlled.
This isn’t the first time I’ve harped on the benefits of cloud agents. I also wrote about them in my post on AI enablement.
I think there’s a big gap in the market for managed agent runtimes (also known as ‘background agents’ because they can run in the background on the cloud). Everything else — org level skills, for eg — is just a band-aid. I’ve said in the past that for coding agents, the filesystem is the config. That means that you have to control the entire filesystem to get a consistent (and secure!) experience. It is simply not enough to assume that everyone in the org will learn not just how to use AI, but also how to use a CLI and how to use Bash and how to use git and how to use the million other tools that are required to get a coding agent off the ground. And even that may not be enough.
I think most teams should go for a fully managed service. Devin has been in the space for a long time. I think Twill is a relative newcomer that also does something like this? But there aren’t a ton of other services that actually follow the Ramp-Inspect / Stripe-Minions model. So when we ran into this problem ourselves, we weren’t happy with most of the other options out there and set out to build our own , which we now sell as an off the shelf option with a lot of customization flexibility.
I have to say, having a tool like this is pretty incredible, and I’m really glad we built it. More than 30% of our PRs are shipped entirely through slack, and we have all sorts of really cool automations for things like bug triage to newsletter writing. But all of that takes a lot of maintenance, and we spend literally all day just thinking about how to make this one thing awesome, and it would be impossible to do that if we were building in, like, healthcare or manufacturing instead.
It’s hard to explain how transformative these are for a company, but once you have them you cannot really go back. My best metaphor is that a good background agent setup is much more like hiring an extremely capable engineer than it is like interacting with a tool in a terminal. It is very natural and intuitive to interact with these things like a person, which short circuits a lot of hard learning curves and enables a lot of collaboration.
And, of course, they are secure. Way way more secure than running a local agent, or giving a new hire intern access to your personal laptop.
12 Grams of Carbon is a reader-supported publication. To receive new posts and support my work, consider becoming a free or paid subscriber.
currently fly.io but we’re considering switching to modal. The only reason we haven’t yet is because gvisor doesn’t make docker support that easy
2 Share Discussion about this post Comments Restacks Top Latest Discussions No posts
