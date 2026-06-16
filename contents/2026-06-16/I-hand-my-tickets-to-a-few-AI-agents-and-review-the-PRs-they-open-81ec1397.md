---
source: "https://www.rafalmuszynski.pl/posts/2026/i-gave-myself-a-team-of-ai-engineers-that-open-their-own-prs"
hn_url: "https://news.ycombinator.com/item?id=48554403"
title: "I hand my tickets to a few AI agents and review the PRs they open"
article_title: "I gave myself a team of AI engineers that open their own PRs | Rafal Muszynski"
author: "takeit"
captured_at: "2026-06-16T13:57:12Z"
capture_tool: "hn-digest"
hn_id: 48554403
score: 1
comments: 0
posted_at: "2026-06-16T12:48:43Z"
tags:
  - hacker-news
  - translated
---

# I hand my tickets to a few AI agents and review the PRs they open

- HN: [48554403](https://news.ycombinator.com/item?id=48554403)
- Source: [www.rafalmuszynski.pl](https://www.rafalmuszynski.pl/posts/2026/i-gave-myself-a-team-of-ai-engineers-that-open-their-own-prs)
- Score: 1
- Comments: 0
- Posted: 2026-06-16T12:48:43Z

## Translation

タイトル: チケットを数人の AI エージェントに渡し、彼らが開いた PR を確認します
記事のタイトル: 独自の PR を開く AI エンジニアのチームを自分自身に与えました |ラファル・ムジンスキー
説明: チケットを取得し、一度に複数の独自のプル リクエストをオープンするエージェント。デモには午後かかった。残りは資格情報とアクセス許可の接続にかかりました。

記事本文:
私は、独自の PR を開く AI エンジニアのチームを自分自身に与えました |ラファル・ムジンスキー
コンテンツにスキップ
Rafal Muszynski li>a]:block [&>li>a]:px-4 [&>li>a]:py-3 [&>li>a]:text-center [&>li>a]:font-medium [&>li>a]:hover:text-accent sm:[&>li>a]:px-2 sm:[&>li>a]:py-1 hidden sm:mt-0 sm:ml-0 sm:flex sm:w-auto sm:gap-x-5 sm:gap-y-0">
投稿
svg]:stroke-accent" title="明暗を切り替える" aria-label="auto" aria-live="polite">
私は独自の PR を開く AI エンジニアのチームを自分自身に与えました
2026 年 6 月 7 日 • 8 分で読了 時間よりも多くのチケットを持っています。
私には本業があり、副業として自分の製品を開発し、そこでテストを行っています。この製品にはバックログがあり、そのほとんどは小規模なもので、ここは修正、あっちは微調整です。正確にはジュニアの仕事ではありません。私が別の開発者を持っていたら、私と同じようにやってくれる人に渡すようなものです。私はしません。それでチケットは座ります。
先週末、あることを試してみました。私はすでに実際のコーディングのほとんどを Claude Code に頼っていますが、それは依然として私がそこに座っていることを意味します。人にチケットを割り当てるのと同じようにチケットを割り当てて、プル リクエストとして戻ってきて確認できるとしたらどうなるでしょうか?ラップトップ上ではなく、私がラップトップの子守りをしている間でも、バックグラウンドで、何か他のことをしている間でもありません。
私が建てました。それは動作します。エージェントの名前は「Dude」です。 Linear で Dude にチケットを割り当てると、しばらくして、変更、説明、チケットへのリンクを含むプル リクエストが GitHub に表示されます。
これはそれを構築する物語であり、主に「AI がコードを書く」よりも困難だったすべての物語です。
男は Anthropic のマネージド エージェントで実行されていますが、メンタル モデルを理解するまでに 1 秒かかりました。私はエージェントを実行していません。 Anthropic はそれを実行します。次に何をするかを決定するループと、コードが実際に実行されるコンテナ (クローン作成、編集) を実行します。

ing、押し込み。チケットを渡して結果を待ちます。私がホストしているのは、チケットを割り当てるときに Webhook をキャッチし、実行を開始して邪魔をしない、薄い接着剤のサービスです。
その接着剤は、小さな Fly マシン上の約 100 行の Python です。これは、私が以前に配線したことのある小さな Webhook ボットのようなものです。彼らがレンタルしている最も安いボックスは、256 MB の RAM を搭載した共有 CPU-1x で、月額 1.94 ドルです。重労働はこちらではなく Anthropic 側で行われるため、それ以上にする必要はありません。 Dude にチケットを割り当てると Webhook が起動されます。接着剤はそれをキャッチし、チケットがどのリポジトリに接触しているかを判断し、それらのリポジトリに向けられたエージェント セッションを開始して、チケットを渡します。エージェントが終了すると、2 番目の Webhook が私に ping を送信します。開いたままにしておくストリームがなく、マシン上で何も実行されていません。
「どのリポジトリ」の部分は、ラベルを設定したときの単なるラベルです。私が Linear テンプレートを作成したので、Dude にファイルしたすべてのチケットが同じ形状になり、テンプレートには単一のリポジトリに固定するラベル ( backend 、 admin-ui 、 www 、 docs ) を付けることができます。それを設定すると、男はそこに留まります。それをオフのままにすると、Dude は変更に必要なリポジトリを見つけ出します。したがって、作業をディスパッチする全体の動作は次のとおりです。テンプレートを選択し、チケットを書き、場合によってはラベルを設定し、それを Dude に割り当てます。ラベルは機能する範囲を狭めます。 Dude に割り当てると、それが機能すると判断されます。
それから私は次の数時間を地味な 90% に費やしました。
これらのエージェントの 1 つのデモは 5 分程度です。チケットを読み取り、ファイルを編集し、PR を開きます。それを初めて見ると、終わったと思わせる部分です。それは行われていません。人間が関与せずに確実に作業を実行し、両者の間の距離が完全に配管になった時点で完了です。
コードを書くのが難しいのと同じように、どれも難しいことではありませんでした。認証、スコープ、構成など、遅い種類のものでした

n は 2 つのシステム間で正確に並ぶ必要がありました。それぞれの壁には本来の権利以上に時間がかかり、二度と再利用することのないものを私に教えてくれました。それが90％です。
最終的にはすべてを単独で実行しました。切符にちなんで名付けられた支店。概要とリンクバックを含むドラフト プル リクエスト。私はそれを見直しました。それは合理的な変更でした。
自分が割り当てたことを忘れていたチケットについて、自分が書いていない PR を見直すのには特別な感情があります。
計画は今でも私と常に関わっており、私が最も諦めたくない部分です。野心的なものについては、Claude Code を開き、満足するまで計画モードでアプローチを実行し、チケットを Linear に直接書き込みます。とにかく私が書いた計画が概要になります。それから私はそれを割り当てて立ち去ります。
こんなに頼りになるとは思っていませんでした。一度に複数発発射できるのです。割り当てられた各チケットは独自の独立した実行を開始するため、他のことをしている間に少数の作業が並行して実行されます。それが本当の利点であることがわかりました。1 人のエージェントの子守をやめたわけではなく、5 人のエージェントが同時に仕事をするようになったのです。小さくて独立したチケットの束なら、簡単に 10 倍になります。 「AI エンジニアがいる」という感覚はなくなり、何人かいるように感じ始めました。私の仕事はチケットを配って、戻ってきたものをレビューすることです。
私のラップトップではどれも実行できないため、ラップトップにいる必要はありません。携帯電話からいくつかのチケットを割り当てて保管し、後で PR を確認します。私が実際に毎日実行していることは、次のようなものです。人に任せて忘れることができる仕事のためのデュードと、自分自身で参加したいときのクロードコードセッションです。おい、座ることをクロード・コードに置き換えるのではなく、私がそこに座る必要のないチケットを受け取るだけだ。
しばらくの間、各チケットは 1 つのリポジトリを意味していました。変化が起こるまでは大丈夫だよ』

t — バックエンドの微調整とそれを呼び出すフロントエンドは、たまたま 2 つの場所に存在する 1 つの作業です。それを2枚のチケットに分けて自分で並べていました。
これで、1 枚のチケットですべてをカバーできるようになります。 Dude は、変更がコードベースのどの部分に影響するかを調べ、それぞれを編集し、リポジトリごとに個別のプル リクエストを開きます。クロスリンクされ、マージする順序に関するメモが付けられます。一度に 1 つの割り当て、1 つの機能でリポジトリをレビューしました。
私が割り当てた小さくて明らかなチケット。野心的な人は最初に計画を立てます。いずれにしても、PR が見つかるまで私は蚊帳の外です。
インテリジェンスがボトルネックになることはありませんでした。
私がぶつかった壁はすべて統合でした。スコープが欠けているトークン、一致する必要がある 2 つの URL、人間を想定したデフォルト、間違った種類の資格情報に基づく権限。モデルがチケットを読み取ってコードを記述する部分は、そのまま機能しました。その周りに組み立てなければならなかったものはすべて揃っていませんでした。
それは新しい教訓ではありません。それはあらゆる統合が教えることです。しかし、「簡単な」部分がシステム作成ソフトウェアであり、「難しい」部分が OAuth である場合、状況は異なります。私たちはモデルが十分に優れているかどうかに多くの時間を費やします。このために、彼らはすでにそうなっていました。隙間は彼らの周りにある退屈な機械ばかりだった。
本業の 1 人チームにとって、魅力は、一度に複数の記事を書く代わりにレビューできることです。それは本当で、少額のチケットであればすでに元は取れます。
それが1か月分のチケットを持ちこたえるのか、それとも節約した時間を発送されたものを壊さないことに費やせるのか、それはこれから分かるだろう。

## Original Extract

Agents that pick up a ticket and open their own pull requests — several at once. The demo took an afternoon; wiring the credentials and permissions took the rest.

I gave myself a team of AI engineers that open their own PRs | Rafal Muszynski
Skip to content
Rafal Muszynski li>a]:block [&>li>a]:px-4 [&>li>a]:py-3 [&>li>a]:text-center [&>li>a]:font-medium [&>li>a]:hover:text-accent sm:[&>li>a]:px-2 sm:[&>li>a]:py-1 hidden sm:mt-0 sm:ml-0 sm:flex sm:w-auto sm:gap-x-5 sm:gap-y-0">
Posts
svg]:stroke-accent" title="Toggles light & dark" aria-label="auto" aria-live="polite">
I gave myself a team of AI engineers that open their own PRs
7 Jun, 2026 • 8 min read I have more tickets than time.
I have a day job and a product of my own on the side, where I test things. The product has a backlog, and most of it is small — a fix here, a tweak there. Not junior work, exactly. The kind of thing I’d hand to another developer if I had one — someone who’d do it the way I would. I don’t. So the tickets sit.
Last weekend I tried something. I already lean on Claude Code for most of my actual coding, but that still means I’m the one sitting there. What if I could assign a ticket the way I’d assign it to a person, and have it come back as a pull request I review? Not on my laptop, not while I babysit it — in the background, while I do something else.
I built it. It works. The agent has a name: Dude. You assign a ticket to Dude in Linear, and a while later a pull request shows up on GitHub with the change, a description, and a link back to the ticket.
This is the story of building it, which is mostly the story of everything that was harder than “the AI writes code.”
Dude runs on Anthropic’s Managed Agents , and the mental model took me a second to get straight: I don’t run the agent. Anthropic runs it — the loop that decides what to do next, and the container where the code actually executes: the cloning, the editing, the pushing. I hand off a ticket and wait for the result. What I host is a thin glue service that catches a webhook when I assign a ticket, kicks off a run, and gets out of the way.
That glue is around a hundred lines of Python on a tiny Fly machine — the kind of small webhook bot I’ve wired up before . The cheapest box they rent, a shared-cpu-1x with 256MB of RAM, $1.94 a month — it doesn’t need to be more than that, since the heavy lifting happens on Anthropic’s side, not here. Assigning a ticket to Dude fires a webhook; the glue catches it, works out which repo or repos the ticket touches, starts an agent session pointed at them, and hands it the ticket. When the agent finishes, a second webhook pings me. No stream to hold open, nothing running on my machine.
The “which repo” part is just a label — when I set one. I made a Linear template so every ticket I file for Dude comes out the same shape, and the template can carry a label — backend , admin-ui , www , docs — that pins it to a single repo. Set it and Dude stays there; leave it off and Dude works out which repos the change needs. So the whole act of dispatching work is: pick the template, write the ticket, maybe set a label, assign it to Dude. The label narrows where it works; assigning to Dude decides that it works.
Then I spent the next couple of hours on the unglamorous 90%.
The demo of one of these agents is a five-minute thing. It reads a ticket, edits files, opens a PR. Watching that the first time is the part that makes you think it’s done. It is not done. Done is when it reliably does the job with no human in the loop, and the distance between those two is entirely plumbing.
None of it was hard in the way writing code is hard. It was the slow kind — auth, scopes, configuration that had to line up exactly across two systems. Each wall cost more time than it had any right to and taught me nothing I’ll ever reuse. That’s the 90%.
Eventually it did the whole thing on its own. A branch named after the ticket. A draft pull request with a summary and a link back. I reviewed it; it was a reasonable change.
There’s a specific feeling to reviewing a PR you didn’t write, for a ticket you forgot you’d assigned.
The planning still happens with me in the loop, which is the part I’m least willing to give up. For anything ambitious I open Claude Code, work the approach in plan mode until I’m happy with it, and have it write the ticket straight into Linear. The plan I’d have written anyway becomes the brief. Then I assign it and walk away.
What I didn’t expect to lean on as much: I can fire several at once. Each assigned ticket spins up its own isolated run, so a handful work in parallel while I do something else. That turned out to be the real gain — not that I’ve stopped babysitting one agent, but that five are working at the same time. On a stack of small, independent tickets that’s an easy 10x. It stopped feeling like “I have an AI engineer” and started feeling like I have a few, and my job is handing out tickets and reviewing what comes back.
And because none of it runs on my laptop, I don’t have to be at one. I’ll assign a few tickets from my phone, put it away, and check the PRs later. What I actually run day to day is a mix: Dude for the work I can hand off and forget, and Claude Code sessions when I want to be in it myself. Dude doesn’t replace sitting down with Claude Code — it just takes the tickets that don’t need me sitting there.
For a while each ticket meant one repo. That’s fine until the change isn’t — a backend tweak and the frontend that calls it are one piece of work that happens to live in two places. I was splitting those into two tickets and lining them up myself.
Now one ticket can cover all of it. Dude works out which parts of the codebase the change touches, makes the edits in each, and opens a separate pull request per repo — cross-linked, with a note on the order to merge them. One assignment, one feature, reviewed a repo at a time.
The small, obvious tickets I just assign. The ambitious ones get the plan first. Either way I’m out of the loop until there’s a PR to look at.
The intelligence was never the bottleneck.
Every wall I hit was integration: a token missing a scope, two URLs that had to match, a default that assumed a human, a permission that lived on the wrong kind of credential. The part where a model reads a ticket and writes the code — that just worked. Everything I had to assemble around it did not.
That’s not a new lesson. It’s the one every integration teaches. But it lands differently when the “easy” part is a system writing software for you and the “hard” part is OAuth. We spend a lot of breath on whether the models are good enough. For this, they already were. The gap was all the boring machinery around them.
For a one-person team with a day job, the appeal is the trade: review instead of write, several at a time. That’s real, and on the small tickets it already pays for itself.
Whether it holds up across a month of tickets, or whether I spend the saved time un-breaking what it ships, I’ll find out.
