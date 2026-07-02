---
source: "https://benemson.com/blog/agents/claudoro-pomodoro-timer-claude-code"
hn_url: "https://news.ycombinator.com/item?id=48760834"
title: "Claudoro: A Pomodoro Timer for Claude Code"
article_title: "Claudoro: a Pomodoro timer for Claude Code — BenEmson.com"
author: "iamflimflam1"
captured_at: "2026-07-02T13:31:27Z"
capture_tool: "hn-digest"
hn_id: 48760834
score: 1
comments: 0
posted_at: "2026-07-02T12:58:30Z"
tags:
  - hacker-news
  - translated
---

# Claudoro: A Pomodoro Timer for Claude Code

- HN: [48760834](https://news.ycombinator.com/item?id=48760834)
- Source: [benemson.com](https://benemson.com/blog/agents/claudoro-pomodoro-timer-claude-code)
- Score: 1
- Comments: 0
- Posted: 2026-07-02T12:58:30Z

## Translation

タイトル: Claudoro: クロード コードのポモドーロ タイマー
記事のタイトル: Claudoro: クロード コード用のポモドーロ タイマー — BenEmson.com
説明: Claudoro は、ステータス ラインに表示される、クロード コード用の無料のオープンソース ポモドーロ タイマーです。ライブ カウントダウン、信頼性の高いアラーム、ゼロ API トークンが表示されます。

記事本文:
Claudoro: クロード コード用のポモドーロ タイマー — BenEmson.com
メイン コンテンツにスキップ BenEmson.com ブログ
· Claudoro: クロード コード用のポモドーロ タイマー
Claudoro: クロード コード用のポモドーロ タイマー
私が一日中見ている場所にあるポモドーロ タイマー、クロード コードのステータス ライン。これを構築した理由と、それがどのように機能するかを説明します。
私は、自分の作業方法に正確に適合する小さなツールを構築するのが好きです。数年前、私はコマンド ライン用のポモドーロ タイマー pymodoro を作成しました。これは、私が長年頼りにしてきた幸せな小さな週末プロジェクトです。それは、私がすでに 1 日を過ごしていたターミナルに常駐するフォーカス クロックです。
古いツールを別の視点から見る
数週間前、私は重大な事故に遭い、脊椎を骨折し、しばらく仰向けになりました。私は問題から抜け出す方法を自分で構築する傾向があるので、ネガティブなことを抱えて座るのではなく、集中して何かを実現できるものを探しました。理想的には、ベッドから再構築できるほど十分に理解していることです。
ピモドーロが明らかな候補だった。しかし、今回は違う見方をしました。私は画面の 1 つのストリップ、つまりモデル、コンテキストのパーセンテージ、git ブランチ、その他のほとんどが空のスペースであるクロード コードのステータス行を 1 日に何時間も見つめています。もしタイマーがそこ、私の目がすでにある場所に住んでいたらどうなるでしょうか？覚えておくべきウィンドウも、Alt-Tab キーを押しながら移動するアプリもありません。ただ、私がいつも見ている 1 つの場所にフォーカスクロックがあるだけです。
🍅 22:47 ▕████████░░▏ ●●○○ Opus · 34% · main それは Claudoro です。ステータス ラインに表示されるクロード コード用のポモドーロ タイマーです。ライブのカチカチするカウントダウン、現在のブロックの進行状況バー、サイクル内のどこにいるのかを示すドットがすべて、すでに存在するステータス情報と並んで表示されます。何も実行していないときはセグメントが消えるため、開始と停止は

レイアウトを変更することはありません。
週末になるはずだった。足が生えてきました。リカバリー中は毎日それを使用し、実際に必要なパーツを追加し続けた結果、それが私が最も手に入れるツールになったのです。
クロードコードまたはターミナルから
私がずっと楽しんでいるのは、ワークフローの両側から答えが得られることです。クロード コード内では、 /pomo start 、 /pomo stop 、 /pomo stats を使用して実行します。モデルのラウンドトリップをまったく必要としない場合は、プロンプトから !pomo start 50 "architecture Spider" を使用して直接実行するか、任意の単純なターミナルから pomo を使用して、まったく同じタイマーを実行します。 1 つのタイマー、1 つの共有状態、私がたまたま手に持ったどの表面でも。この小さな柔軟性が、予想以上に重要であることがわかりました。
正しく仕上げることに気を配っていた部分
構築中に私にとって重要な点がいくつかありました。
既存のステータス ラインを踏みにじるのではなく、横に置く必要があるため、Claudoro はモデル、コンテキスト、git など、すでに表示されているものを使用して構成し、何かに触れる前に設定を静かにバックアップします。また、実行にコストはかかりません。タイマーは、クロード コードが線を描画するために呼び出す小さなローカル CLI にすぎないため、ループ内にモデルはなく、単一の API トークンも消費されません。レンダリング パスは意図的に低コストであり、ティックごとに 1 つの小さな状態ファイルを読み取ります。
アラームも信頼できるものでなければなりませんでした。ブロックが開始されると、Claudoro はスリープしてから鳴る小さな分離プロセスに引き渡します。そのため、ステータス ラインを非表示にしたり、すべてのセッションを閉じたりした場合でも、チャイムは鳴り続けます。一度に十数のセッションを実行すると、アラームが 1 つだけ発生します。
私が静かに満足しているビットは、いつも私をつまずかせていた 1 つの問題を解決します。私が豆を挽いてエスプレッソを淹れるためにうろうろしている間にブロックが長くなった場合、クラウドロは計画した時間と短い猶予時間のみをクレジットします。

、記録を放棄としてマークし、正直に言うと本当のスパンを保持します。私の統計は、午後デスクから離れても静かに膨らむことはありません。
そして、それはすべてマシン上に残ります。アカウント、ネットワーク、テレメトリはなく、状態ディレクトリの下にある追加専用のログだけです。これはノードであり、ランタイム依存関係はなく、MIT ライセンスが付与されており、セットアップは完全に元に戻すことができる単一の冪等コマンドです。
途中で、計画していなかったことがいくつか起こりました。 1 つ目は、統計ビュー、現在のストリーク、フォーカス ヒートマップ、およびトップ タグです。これは、自己完結型のオフライン ダッシュボードとしてもレンダリングされ、単一のローカル HTML ファイルにフォーカス履歴全体が表示されます。
2 つ目は、ポモドーロ テクニック自体への組み込みガイドで、単に時間を計るだけでなくこのメソッドをうまく使いたい場合に役立ちます。
npm install -g クラウドロ
pomo setup 次に、新しい Claude Code セッションを開き、 /pomo start を実行します。約 1 秒以内にステータス ラインにカウントダウンが表示されます。
もしあなたが試してみたら、どう感じたのか、何が正しいと感じたのか、何が壊れたのか、どうなったらよかったかなど、ぜひ聞きたいです。問題を開くか、単に返信してください。
そして、それがあなたの一日の中での役割を獲得した場合、あなたができる最も親切なことは、GitHub でそれにスターを付けることです。それは、そのアイデアが追求する価値があることを示す最も明白な兆候です。
Karpathy の LLM Wiki、6 週間
4 月 4 日、Andrej Karpathy は、個人の知識ベースに関する新しい考え方を盛り込んだ 1,000 語の要旨を発表しました。 6 週間後、私の保管庫は 2 時間で 6 つのプロジェクトを監査し、そのうちの 5 つで同じ欠陥を発見しました。 3 つのレイヤー、1 つの所有権ルール、および編集されたメモが大きくなる代わりに密度が高くなる理由。
GitHub スポンサーの Memecoin の同意ギャップの構造
誰かが私のプロジェクトを褒めてくれました。私は嬉しくなって、何も考えずに「はい」と答えたところ、その名前でミームコインがローンチされました。私のプロジェクトの信頼性は製品でした。人々はお金を失ったかもしれない。

ここで私が最初に尋ねるべきでした。
My Agent Memory Library はインディーズ記事の執筆に役立ちます
カスタム メモリ ライブラリは、エージェントが優れたインディーズ開発記事を書くのに役立ちますか?私たちはそれをライブでテストします: elfmem に基づいた調査、ボールトのリコール、ピア エージェントの通信、人間の編集者。出てきたのはこちらです。

## Original Extract

Claudoro is a free, open-source Pomodoro timer for Claude Code that lives in your status line: a live countdown, a reliable alarm, and zero API tokens.

Claudoro: a Pomodoro timer for Claude Code — BenEmson.com
Skip to main content BenEmson.com Blog
· Claudoro: a Pomodoro timer for Claude Code
Claudoro: a Pomodoro timer for Claude Code
A Pomodoro timer that lives where I already look all day: the Claude Code status line. Here is why I built it, and how it works.
I like building small tools that fit the exact shape of how I work. A few years ago I built a Pomodoro timer for the command line, pymodoro , a happy little weekend project that I ended up leaning on for ages. It did one thing well: a focus clock that lived in the terminal, where I already spent my day.
Looking at an old tool differently
A few weeks ago I had a serious accident and fractured a vertebra, which put me flat on my back for a while. I tend to build my way out of problems, so rather than sit with the negatives I went looking for something to focus on and make. Ideally something I already understood well enough to rebuild from bed.
Pymodoro was the obvious candidate. But this time I looked at it differently. I spend hours a day staring at a single strip of the screen, the Claude Code status line: model, context percentage, git branch, and otherwise mostly empty space. What if the timer lived there, right where my eyes already are? No window to remember, no app to alt-tab to, just a focus clock in the one place I am always looking.
🍅 22:47 ▕████████░░▏ ●●○○ Opus · 34% · main That is Claudoro: a Pomodoro timer for Claude Code that renders right in the status line. A live, ticking countdown, a progress bar for the current block, and dots for where you are in the cycle, all sitting alongside the status info that was already there. The segment disappears when nothing is running, so starting and stopping never nudges your layout.
It was meant to be a weekend. It grew legs. I used it every day through the recovery and kept adding the parts I actually wanted, until it had quietly become the tool I reach for most.
From Claude Code, or the terminal
The part I keep enjoying is that it answers from both sides of my workflow. Inside Claude Code I drive it with /pomo start , /pomo stop , /pomo stats . When I do not want a model round-trip at all, I run the very same timer straight from the prompt with !pomo start 50 "architecture spike" , or from any plain terminal with pomo . One timer, one shared state, whichever surface I happen to have my hands on. That small bit of flexibility turned out to matter more than I expected.
The parts I cared about getting right
A handful of things mattered to me while building it.
It had to sit alongside my existing status line rather than trample it, so Claudoro composes with whatever you already show, model, context and git, and quietly backs up your settings before it touches anything. It also had to cost nothing to run: the timer is just a small local CLI that Claude Code calls to paint the line, so there is no model in the loop and not a single API token spent. The render path is deliberately cheap, reading one small state file per tick.
The alarm had to be dependable too. When a block starts, Claudoro hands off to a small detached process that sleeps and then sounds, so the chime still lands even if I have hidden the status line or closed every session. Run a dozen sessions at once and exactly one alarm fires.
The bit I am quietly pleased with fixes the one thing that always used to trip me up. If a block runs long while I have wandered off to grind some beans and pull an espresso, Claudoro credits only the planned time plus a short grace, marks the record as abandoned, and keeps the true span for honesty. My stats never get quietly inflated by an afternoon away from the desk.
And it all stays on your machine: no accounts, no network, no telemetry, just an append-only log under your state directory. It is Node, zero runtime dependencies, MIT licensed, and the setup is a single idempotent command you can cleanly undo.
A couple of things crept in along the way that I had not planned. The first is a stats view, current streak, a focus heatmap and top tags, which also renders as a self-contained offline dashboard, your whole focus history in a single local HTML file.
The second is a built-in guide to the Pomodoro Technique itself, for when you want to use the method well rather than just run a clock.
npm install -g claudoro
pomo setup Then open a new Claude Code session and run /pomo start . The countdown appears in your status line within about a second.
If you give it a go, I would genuinely love to hear how you get on: what felt right, what broke, what you wish it did. Open an issue, or just reply.
And if it earns a place in your day, the kindest thing you can do is give it a star on GitHub. It is the clearest signal that the idea was worth chasing.
Karpathy's LLM Wiki, six weeks in
On 4 April, Andrej Karpathy published a thousand-word gist with a new way of thinking about personal knowledge bases. Six weeks in, my vault audited six projects in two hours and found the same flaw in five of them. Three layers, one ownership rule, and why compiled notes get denser instead of bigger.
Anatomy of a GitHub Sponsor Memecoin Consent Gap
Someone praised my project. I got flattered, said yes without thinking, and a memecoin launched in its name. My project's credibility was the product. People may have lost money. Here is what I should have asked first.
My Agent Memory Library Helps Write Indie Articles
Can a custom memory library help an agent write good indie dev articles? We test it live: elfmem-grounded research, vault recall, peer agent communication, and a human editor. Here is what came out.
