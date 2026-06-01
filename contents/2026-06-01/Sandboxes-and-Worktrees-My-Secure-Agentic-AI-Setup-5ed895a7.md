---
source: "https://mikemcquaid.com/sandboxed-agent-worktrees-my-coding-and-ai-setup-in-2026/"
hn_url: "https://news.ycombinator.com/item?id=48347727"
title: "Sandboxes and Worktrees: My Secure Agentic AI Setup"
article_title: "Sandboxes and Worktrees: My secure Agentic AI Setup | Mike McQuaid"
author: "CoffeeOnWrite"
captured_at: "2026-06-01T00:26:56Z"
capture_tool: "hn-digest"
hn_id: 48347727
score: 3
comments: 0
posted_at: "2026-05-31T17:39:07Z"
tags:
  - hacker-news
  - translated
---

# Sandboxes and Worktrees: My Secure Agentic AI Setup

- HN: [48347727](https://news.ycombinator.com/item?id=48347727)
- Source: [mikemcquaid.com](https://mikemcquaid.com/sandboxed-agent-worktrees-my-coding-and-ai-setup-in-2026/)
- Score: 3
- Comments: 0
- Posted: 2026-05-31T17:39:07Z

## Translation

タイトル: サンドボックスとワークツリー: My Secure Agentic AI セットアップ
記事のタイトル: サンドボックスとワークツリー: 私の安全な Agentic AI セットアップ |マイク・マクエイド
説明: 一度に 1 つの AI の子守をやめます。サンドボックスにより安全に実行でき、Git ワークツリーにより並列実行が可能になります。より多くのトークンを使用すると、より多くの速度が得られます。

記事本文:
サンドボックスとワークツリー: 安全な Agentic AI セットアップ |マイク・マクエイド
マイク・マクエイド
CTPO および Homebrew プロジェクト リーダー
サンドボックスとワークツリー: 安全な Agentic AI セットアップ
私は 2021 年の初めに、GitHub で Copilot の内部アルファ版をテストするよう招待されて以来、AI ツールを使用しています (そこで 10 年間過ごしました)。
私は 2009 年から Homebrew を保守しています。
私は現在、「AI がコードの 90% を記述する」 (Dario Amodei の 2025 年初頭から 2025 年後半までの予測) を個人的に当てています。
何人かの人から私の現在の設定について詳しく尋ねられたので、それをここに示します。
エージェントはコード補完からコード生成へと段階的な変化をもたらし、多くの問題をワンショットで解決できるようになりました。
サンドボックス化により、エージェントが許可プロンプトの子守なしで暴走できるようになり、セキュリティと生産性が向上します。
Git ワークツリーは作業を並列化するため、より多くのトークン/支出が速度の向上に直接つながります。
AI をまだ「コード補完モード」で使用している場合、または GitHub Copilot のみを使用している場合は、チャンスを逃しています。
2025 年半ばから後半にかけて、Claude Code や OpenAI Codex などのエージェント ツールが十分に機能し、熱心に取り組んでいる場合でも編集よりもプロンプトの方が速くなりました。
Claude Code と OpenAI Codex に関する私の経験は主に次のとおりです。
OpenAI Codex (5.4 xhigh ): 時間がかかりますが、通常は最初からほとんどのことが正しく行われ、プロンプトや操作はあまり必要ありません
Claude Code (Opus 4.6 max ): デフォルトではかなり愚かですが、攻撃的なフック/ツール/プロンプトを使用してそれほど愚かにならないように調整できます。
(Opus 4.7 がリリースされたばかりです。もっと愚かかどうか見てみましょう…)
私が毎日選んでいるドライバーは OpenAI Codex ですが、トークンがすぐになくなってしまうので、どちらでも大丈夫であることを学びました。
(OpenAI: OSS メンテナーとして申請した無料トークンをください。ありがとう <3)。
多くの人が自分の C について考えることに多くの時間とエネルギーを費やしています。

LAUDE.md / AGENTS.md 。
私の経験では、パフォーマンスはモデルごと、バージョンごとに大きく異なるため、できるだけ最小限に抑える価値があります。
私のドットファイルには AGENTS-GLOBAL.md があり、Claude と Codex のすべてのプロジェクトにわたって必要な最低限の情報が含まれています。
エージェントとすぐに遭遇する主な問題は、人生の半分を「はい、これは安全な仕事をしてください」、「いいえ、この危険な仕事はやめてください」と言い続けなければならないことです。
これは退屈で遅いです。
私は怠け者なので、より良い自動化が必要でした。
さまざまなエージェントには、さまざまな「許可チェックのバイパス」、「サンドボックスなしで実行」、「YOLO」などのモードがあります。
これらのツールを使用して実際に生産性を高めたい場合は、基本的に 2 つのオプションがあります。
火遊びをするだけだと決めて権限を無効にし、何も問題が起こらないことを祈ります
サンドボックス環境で実行します。例: VM、別個のマシン、(システム) 権限、アクセス、トークンなどが制限されたサンドボックス。
私がオプション 2 を選択したのは、Homebrew メンテナーとして、自分のマシンで愚かで安全でないことをしないという責任を真剣に考えているからです。
残念ながら、私は Homebrew メンテナーでもあるため、ローカル macOS 開発に Docker を使用することにアレルギーがあります。
Sandvault は私が出会った中で最も優れた中間点です。
macOS サンドボックスを利用し、管理者以外の別のユーザー アカウントを作成して維持し、自由に実行できるようにします。
コードを漏洩することを除けば (OSS では心配していません)、エージェントが次のような可能性がある、私が懸念しているリスクベクトルの大部分を遮断できます。
たとえば行くバージョン管理にないファイルに対する rm -rf
私の例を使用してください機密性の高いリポジトリで処理を行うための GITHUB_TOKEN
機密ファイルを別の場所に持ち出す
インストールしたら ( brew install Sandvault )、sv codex または sv claude を実行して、選択したエージェントを起動できます。

e または sv シェルを使用してシェルを開始します。
ドットファイルを /Users/Shared/sv-${USER}/user に置くこともでき、それらは関連する Sandvault にコピーされます。 ( Sandvault-mike ) ユーザー。
これを行う方法の例が必要な場合は、私の dotfiles のスクリプト/同期を見てください。
また、Sandvault ユーザーに対して異なる色のプロンプトを使用して、自分がその中にいることを認識できるようにすることをお勧めします。
例については、私のドットファイル shprofile.sh と zprofile.sh を参照してください。
これはすべてうまく機能しますが、現在の $USER (例: mike ) と特権のない Sandvault ユーザー (例: Sandvault-mike ) の間でコードをより簡単に共有したい場合はどうすればよいでしょうか?
私は通常、自分のすべての OSS リポジトリを ~/OSS/* にクローンし、雇用主 (現在は Administrate ) のリポジトリを例えば OSS リポジトリにクローンします。 ~/管理。
代わりに、Sandvault で作成した /Users/Shared/sv-mike の下で /Users/Shared/sv-mike/repositories にクローンを作成し、/Users/Shared/sv-mike/worktrees を作成します (詳細は後ほど)。
これにより、両方のユーザーが読み書きできる安全な場所を確保できます。
グループの書き込み可能権限で git を混乱させないように、両方の ~/.gitconfig に追加する必要があることに注意してください。
【安全】
# Sandvault ディレクトリは別のユーザーが所有していることが想定されます。
ディレクトリ = / ユーザー / 共有 / sv - mike / リポジトリ /*
🌳 ワークツリー
Sandvault が Claude および Codex とうまく連携できるようになると、それらを独立して動作させるだけで生産性が大幅に向上しました。
ただし、現在使用しているトークンよりも多くのトークンを支払う場合、一度に 1 つのエージェントを実行することがボトルネックになります。
複数のリポジトリ間でマルチタスクを実行するのは簡単な次のステップですが、同じリポジトリでこれをもっと実行したいと思うようになりました。
私は最終的に homebrew と homebrew2 リポジトリを作成しました。これはひどいと感じましたが、まあまあ機能しました。
問題は、どのゲームで何をしたかを思い出すことでした。
Git 作業

ツリーを使用すると、同じリポジトリの複数のブランチを別々のディレクトリで同時にチェックアウトできます。
おそらくここではそれらが正しいソリューションであることはわかっていましたが (私は Git についての本を書いたので言い訳はできません)、それらを使って遊んでから何年も経っていました。
また、これらすべてを手動で構築する必要もありませんでした。
私はすでに Sandvault を中心にたくさんの git ヘルパー エイリアスを構築していました。
有能な CTO の友人が、ワークツリーを使用して多数のエージェントを実行する方法として Conductor を教えてくれました。
可能性はありましたが、私は Sandvault が大好きすぎて、うまくプレイさせる方法が見つかりませんでした。
代わりに、同様の機能を備えた Superset を見つけましたが、重要なことに、コマンドをオーバーライドできるようになりました。
「ワークツリーの場所」を /Users/Shared/sv-mike/worktrees に設定しました。
Sandvault コマンドを使用するように「エージェント」を設定しました (「プロンプトなし」オプションと「プロンプトあり」オプションも同様)。
これらはすべて、Sandvault が現在サポートしているものです (私は数日前に数行のコードで OpenCode を簡単に追加しました)。
次に、 /Users/Shared/sv-mike/worktrees に複製したプロジェクトに基づいて、多数の「プロジェクト」を追加しました。
それが役立つ場合には、基本的なものを実行してもらいます。 script/bootstrap を使用して、プロジェクトをより適切にセットアップします。
これがすべてセットアップされたら、プロジェクトごとにターミナルを簡単に起動し、そこで選択したツールやエージェントを実行できます。
ただし、実際に楽しくて強力になるのは、ワークツリーの作成です。
これにより、単一のプロンプトに基づいて、サンドボックス化された新しいワークツリーでエージェントを起動できます。
これがうまく機能すれば、限界は無限です。
私の個人的なワークフローは次のとおりです。
「Homebrew または作業コードの問題を読んで、TODO リストに追加してください」
「問題の説明をコピー＆ペーストし、それを解決する方法についてのブレインダンプをプロンプトに入れて、エージェントに作業してもらいます。」
時々私はただ発砲するだけです

異なるアプローチで複数のエージェントを同時に実行し、最も気に入らないエージェントを破棄します。
私はエージェントが制作した作品を他の人と共有する前に、必ずローカルでレビューします。
レビューには次の 1 つ以上が含まれます。
生成されたすべての出力の読み取り (例: 素晴らしい macOS Git GUI である Fork の使用)
生成された出力を手動で編集する (例: 現在選択しているエディターである Zed を使用する)
Zed を選んだ理由は、エディターに費やす時間が減り、Cursor の優れたオートコンプリートが必要なくなり、起動速度を重視するようになったからです。
生成された出力を強制的に編集するよう求めるプロンプト (例: 新しいプロンプトとしてローカル コード レビュー コメントを提供する)
検証テストを手動で実行する (例: ローカルで実行する、出力を確認する、AI エラー メッセージを表示する)
別の AI に AI の作業をレビューしてもらう (例: PR での Copilot コード レビュー)
CI エラーの提供 (GitHub Actions の出力など)
これらのうちどれだけ行うかは、コードへの精通度、その重要性、および他のガードレールにどれだけ自信があるかによって決まります。 CI、人によるレビュー。
スーパーセットの Zed ( zed .; exit ) と Fork ( fork .; exit ) の「プリセット」を使用して、ワンクリックでワークツリー内にプリセットを起動することでこれを簡単にします。
AI によって生成された他者の作業をレビューするとき、私は出力をレビューし、場合によっては手動でテストします。
でも、私はそっちのほうが怠け者なんです。
AI があなたのためにコードを書いているのであれば、レビュー担当者や同僚のためにテストの方法をさらに強化する必要があります。
現在、Homebrew では次のようなより高い障壁を設定しています。
PR で AI の開示を要求する
非メンテナが同時に開いている AI 生成 PR は 1 つ以上であることを要求しない
AI が PR を作成し、PR テンプレートを破棄するときにコメントなしで閉じます
大きすぎる PR を複数の小さい PR に分割する必要がある
低品質の AI PR の作成をやめようとしない人をブロックする
これはまだある

ただ、純粋に人間の仕事をレビューするのと比べると、疲れるしやる気も失せます。
その結果、人間ではなくエージェントと話しているように感じられる PR を単に終了してしまうことがあります。
ここまで読んで、「あなたは CTPO ではないのですか? コーディング以外のことを主にやるべきではないでしょうか?」と思った方は、次のようになります。
私はそうです、あなたがエンジニアなら読んでもあまり面白くありません
OK、もしあなたが主張するなら、もう一つ素晴らしいものをお見せしたいのです
私は、CTO エグゼクティブアシスタントに関する Obie Fernandez のツイートを読み、興味をそそられました。
私は 2 回目の CTPO の仕事に就いており、かなり整理されていると感じていますが、通常のシステムとメモリでは、仕事をうまく進めるために必要な膨大な量のコンテキストが不足しています。
私は人間のエグゼクティブアシスタントを雇ったことがありませんし、それを扱うにはコントロールフリークすぎるのではないかと強く疑っています。
スーパーセットには、 ctpo というプロジェクトがあります。
この中に、さまざまな会議のメモ、私の目標、会社の目標、個人の TODO などを入れます。
その後、会議の準備、短期/中期/長期の優先事項やリサーチを行うために質問することができます。
最初の「トレーニング」は、私が作成した社内のエンジニアリング文化に関する文書と、この Web サイトに書かれたコーパスに対して行われました。
このコーパスは、これも最新の AI ツールのおかげで、私が行ったさまざまな講演やポッドキャストのトランスクリプトを投稿とともに保存することで役に立ちました。
これにより、私の高レベルの価値観のほとんどをすでに認識しており、はるかに優れた記憶力 (これはほとんど正確です) を与えてくれるパーソナル アシスタントが私に与えられました。
最近 GitHub から連絡があり、Homebrew は、AI エージェント後のマージで減少ではなく増加が見られる少数のプロジェクトの 1 つであるようです。
これは偶然ではないと思います。この設定がその理由の大きな部分を占めています。
Claude/Codex、Sandvault、Superset、および私の CTPO アシスタントの間

これまでで最も生産的で、組織化されていると感じています。
サンドボックスとワークツリーを組み合わせることで、より多くの問題に対してより多くのトークンを投げるだけで、より多くのことを達成できるようになります。
職場では、自分自身の思い出や TODO リストの拡張版を手に入れているように感じます。これにより、手作業での管理がそれほど必要なくなります。
ここでも面白いことをやっているなら、連絡してください！
私は他の人から自分の改善点や間違っている点を学ぶことに興味があります。
これがお役に立てば幸いです。そして、来年の今頃、この記事が笑えるほど時代遅れになることを楽しみにしています。
皆さん、頑張ってください。今はワイルドなライドです 💜。
この投稿の下書きをレビューしてくれた Gus Fune と John Peebles に感謝します。
AIではなく私が書いたものです。私は AI の提案を求め、同意する場合にのみ適用します (同意しない場合も多い)。
私の子育て「スクリーンタイム」の哲学
新しいコンテンツに関する通知を受け取るために購読する
不定期にメールで最新情報をお知らせします。 RSS/Atom フィードは以下のリンクにあります。

## Original Extract

Stop babysitting one AI at a time. Sandboxing lets them run wild safely, Git worktrees let them run in parallel. Use more tokens, get more velocity.

Sandboxes and Worktrees: My secure Agentic AI Setup | Mike McQuaid
Mike McQuaid
CTPO and Homebrew Project Leader
Sandboxes and Worktrees: My secure Agentic AI Setup
I’ve been using AI tools since early 2021 when I was invited to test out the Copilot internal alpha at GitHub (where I spent 10 years).
I’ve maintained Homebrew since 2009.
I’ve now personally hit the “AI writes 90% of my code” ( Dario Amodei’s early 2025 prediction for late 2025).
I’ve been asked by a few folks to detail my current setup so: here it is.
Agents bring a step change from code completion to code generation and are good enough now to one-shot many problems
Sandboxing improves security and productivity by letting agents run wild without babysitting permission prompts
Git worktrees parallelise work so more tokens/spend directly translates to more velocity
If you’re still using AI in “code completion mode” or only GitHub Copilot, you’re missing out.
Mid-to-late 2025 was when agentic tools like Claude Code and OpenAI Codex got good enough that prompting became quicker than editing, even when you’re anal about it.
My experience of Claude Code and OpenAI Codex has mostly been:
OpenAI Codex (5.4 xhigh ): takes a while, generally does things mostly right first time, doesn’t need much prompting/steering
Claude Code (Opus 4.6 max ): fairly stupid by default, can be nudged with aggressive hooks/tools/prompts to being less so
(Opus 4.7 just got released: let’s see if it’s less stupid…)
My daily driver of choice is OpenAI Codex but I run out of tokens more quickly so have learned to be fine with either.
(OpenAI: please give me the free tokens I’ve applied for as an OSS maintainer, thanks <3).
A lot of people spend a lot of time and energy thinking about their CLAUDE.md / AGENTS.md .
My experience is their performance varies so much from model to model, version to version that it’s worth keeping them as minimal as possible.
I have an AGENTS-GLOBAL.md in my dotfiles that provides a decent minimum of what I care about across all projects in Claude and Codex.
The main problem you’ll bump into pretty quickly with agents is: you have to spend half your life going “Yes, do this safe thing”, “No, don’t do this dangerous thing”.
This is boring and slow.
I’m lazy so needed better automation.
The various agents have various “bypass permission checks”, “run without sandbox”, “YOLO”, etc. modes.
If you want to be actually productive with these tools you basically have two options:
decide you’re just going to play with fire and disable permissions and hope nothing goes wrong
run in a sandboxed environment e.g. a VM, separate machine, sandbox with reduced (system) permissions, access, tokens, etc.
I picked option 2 because I take my responsibility as a Homebrew maintainer seriously to not do stupid insecure shit on my machine.
Unfortunately, also because I’m a Homebrew maintainer, I’m allergic to using Docker for local macOS development.
sandvault is the nicest middle ground I’ve come across.
It makes use of macOS sandboxes and creates and maintains a separate non-admin user account where you can let it run wild.
Short of exfiltrating your code (which I’m not worried about with OSS), it closes the majority of risk vectors I care about where agents might:
go e.g. rm -rf on files not in version control
use my e.g. GITHUB_TOKEN to do things on sensitive repositories
exfiltrate sensitive files elsewhere
Once installed ( brew install sandvault ), you can run sv codex or sv claude to start your agent of choice or sv shell to start a shell.
You can also put your dotfiles under /Users/Shared/sv-${USER}/user and they will be copied to the relevant sandvault e.g. ( sandvault-mike ) user.
Take a look at my dotfiles’ script/sync if you want an example of how to do this .
I also recommend using a different-coloured prompt for your Sandvault user so you know when you’re inside it.
See my dotfiles shprofile.sh and zprofile.sh for examples.
This all works nicely but: what if you want to share code more easily between your current $USER (e.g. mike ) and the unprivileged sandvault (e.g. sandvault-mike ) user?
I would normally clone all my OSS repositories into ~/OSS/* and employer’s (currently Administrate ) into e.g. ~/Administrate .
Instead, I now clone under the Sandvault-created /Users/Shared/sv-mike into /Users/Shared/sv-mike/repositories and create /Users/Shared/sv-mike/worktrees (more on that later).
This lets me have somewhere safe that’s readable and writable to both users.
Note, you’ll need to add to your ~/.gitconfig for both to not freak out git with the group writable permissions:
[ safe ]
# It's expected that sandvault directories are owned by another user.
directory = / Users / Shared / sv - mike / repositories /*
🌳 Worktrees
Once I had Sandvault working nicely with Claude and Codex I could be a lot more productive with just letting them work independently.
However, running one agent at a time becomes the bottleneck when you’re paying for more tokens than you’re currently using.
Multitasking between multiple repositories was an easy next step but I found myself wanting to do this more on the same repository.
I ended up with a homebrew and homebrew2 repository which felt gross but kinda worked.
The problem was remembering what I did on which one.
Git worktrees let you have multiple branches of the same repository checked out simultaneously in separate directories.
I knew they were probably the right solution here (I wrote a book about Git so have no excuse) but it’d been ages since I’d played with them.
I also didn’t want to have to build all this manually.
I’d already built a bunch of git helper aliases around Sandvault.
A talented CTO friend pointed me at Conductor as a way of running a bunch of agents with worktrees.
It had potential but I love Sandvault too much and couldn’t figure out any way to make them play nicely.
Instead, I stumbled upon Superset which did similar but, importantly, allowed me to override commands.
I configured my “Worktree location” to /Users/Shared/sv-mike/worktrees .
I set my “Agents” to use their Sandvault commands (same for “No Prompt” and “With Prompt” options):
These are all those that Sandvault supports today (I easily added OpenCode a few days ago with a few lines of code).
I then added a bunch of “Projects” based on those I’d cloned into /Users/Shared/sv-mike/worktrees .
For those where it’s useful, I have them run a basic e.g. script/bootstrap so the project is better set up.
Once this is all set up, you can easily spin up terminals for each project and run whatever tool or agent you choose there.
Creating worktrees is where this actually gets fun and powerful though:
This lets you spin up an agent in a sandboxed new worktree based on a single prompt.
Once you have this working: the sky is the limit.
My personal workflow has gone from:
“reading Homebrew or work code problem, add to my TODO list”
“copy paste problem description plus braindump of how I think it should be solved into a prompt, let the agent work on it”
Sometimes I’ll just fire up multiple agents with different approaches to run at the same time and throw away those I like the least.
I always review any agent-produced work locally before I share it with others.
Review can involve one or more of:
reading all generated output (e.g. using Fork , a nice macOS Git GUI)
manually editing generated output (e.g. using Zed , my current editor of choice)
I picked Zed because I spend less time in my editor now, don’t need Cursor’s better autocomplete and care more about startup speed
prompting to force edits of generated output (e.g. providing local code review comments as new prompts)
manually performing verification tests (e.g. running things locally, reviewing the output, giving the AI error messages)
getting another AI to review the work of the AI (e.g. Copilot Code review in PRs)
providing CI failures (e.g. GitHub Actions output)
How many of these I do depends on my familiarity with the code, its criticality and how confident I am in other guardrails e.g. CI, human review.
I make this easier with Zed ( zed .; exit ) and Fork ( fork .; exit ) “Presets” in Superset to launch them in the worktree with one click.
When reviewing AI-generated work from others, I review the output and sometimes manually test.
I’m lazier there, though.
If AI is writing your code for you now, you need to step up and do more in the way of testing for your reviewers/coworkers.
We’ve set higher barriers here in Homebrew now including:
requiring AI disclosure on PRs
requiring non-maintainers do not have more than 1 AI generated PR open at once
closing without comment when AIs create the PR and discard our PR template
requiring PRs that are too large be split into multiple, smaller ones
blocking people who refuse to stop creating low-quality AI PRs
This is still sometimes tiring and demotivating though compared to reviewing the work purely of humans.
As a result, we will sometimes just close out PRs where it feels like we’re speaking to their agent and not the human.
If you’ve read all this and gone “aren’t you a CTPO? shouldn’t you be mostly doing non-coding things?”:
I am, it’s just not very interesting to read about if you’re an engineer
Ok, if you insist: I have another cool thing to show you
I read this tweet from Obie Fernandez about a CTO executive assistant and was intrigued.
I’m on my second CTPO job and feel fairly organised but that my usual systems plus my memory fail the huge amount of context I now need to do my job well.
I’ve never had a human executive assistant and strongly suspect I’d be too much of a control freak to handle one.
In Superset, I have a project called ctpo .
In this I put in various meeting notes, my goals, company goals, personal TODOs, etc.
I can then ask it questions to help me prepare for meetings, what my short/medium/long-term priorities are and do research.
The initial “training” was done on internal engineering culture documents I’ve written and the written corpus on this website.
This corpus was helped by, again thanks to modern AI tooling, storing transcripts from various talks and podcasts I’ve done along with my posts.
This has given me a personal assistant that already knows most of my high-level values and gives me a vastly better memory (which is mostly accurate).
GitHub reached out recently and apparently Homebrew is one of a smaller number of projects that’s seen an uptick rather than downtick in merges post-AI agents.
I don’t think that’s a coincidence: this setup is a big part of why.
Between Claude/Codex, Sandvault, Superset and my CTPO assistant I feel the most productive and organised I’ve ever been.
The combination of sandboxing and worktrees means I can just throw more tokens at more problems and get more done.
At work, I feel like I’ve got an extended version of my own memories and TODO lists that doesn’t require as much manual curation.
If you’re doing interesting things here too: get in touch!
I’m interested to learn from anyone else what I could be doing better or am doing wrong.
Hopefully this was useful and I am excited for it to be hilariously outdated this time next year.
Good luck out there everyone, it’s a wild ride just now 💜.
Thanks to Gus Fune and John Peebles for reviewing drafts of this post.
Written by me, not an AI. I solicit AI suggestions and apply only when I agree (often I don't).
My Parenting "Screen Time" Philosophy
Subscribe to be notified about new content
Occasional updates by email. RSS/Atom feeds are linked below.
