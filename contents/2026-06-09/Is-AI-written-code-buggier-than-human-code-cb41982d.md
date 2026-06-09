---
source: "https://www.repowise.dev/blog/engineering/is-ai-written-code-buggier-than-human-code"
hn_url: "https://news.ycombinator.com/item?id=48459397"
title: "Is AI-written code buggier than human code?"
article_title: "Is AI-Written Code Buggier Than Human Code?"
author: "raghavchamadiya"
captured_at: "2026-06-09T13:09:02Z"
capture_tool: "hn-digest"
hn_id: 48459397
score: 2
comments: 0
posted_at: "2026-06-09T11:00:09Z"
tags:
  - hacker-news
  - translated
---

# Is AI-written code buggier than human code?

- HN: [48459397](https://news.ycombinator.com/item?id=48459397)
- Source: [www.repowise.dev](https://www.repowise.dev/blog/engineering/is-ai-written-code-buggier-than-human-code)
- Score: 2
- Comments: 0
- Posted: 2026-06-09T11:00:09Z

## Translation

タイトル: AI が書いたコードは人間のコードよりもバグが多いですか?
記事のタイトル: AI が書いたコードは人間のコードよりもバグがあるのか​​?
説明: AI エージェントのコードが人間のコードよりも多くのバグを引き起こすかどうかをテストするために、28 のリポジトリにわたる 112,382 のコミットを git 非難しました。サイズを調整した後は、そうではなくなり、ラインが長持ちします。

記事本文:
AIが書いたコードは人間のコードよりもバグが多いのでしょうか? repowise Docs 価格 ブログスターを調べる — pip install repowise サインイン ブログ / エンジニアリング AI が書いたコードは人間のコードよりもバグが多いですか?私たちはそれを解明するために 112,000 件のコミットを非難しました
10 人のエンジニアに、AI が書いたコードが手で書いたコードよりもバグが多いかどうかを尋ねると、自信に満ちた 10 件の答えが得られます。それぞれの方向で約半数、数字で裏付けられた答えは 0 つです。これは、誰もが事前知識を持っていても、誰もデータを持っていない質問の 1 つです。テイクは大音量で、測定値はほとんどなく、存在する数少ない測定値は、マージ後のコードが実際に何をするかというよりも、AI コードについて人々がどのように感じているかを調査する傾向があります。
そこで私たちは、欠陥予測フィールドが他のすべてのものを測定するのと同じ方法で、実際のプロジェクトの git 履歴を遡って、すべてのバグ修正をバグを導入したコミットのせいにすることによって測定しました。実際のプル リクエストをマージするコーディング エージェントの最初の 1 年にわたって、28 のパブリック リポジトリと 112,382 件のコミットにわたってこれを実行し、データについて簡単な質問をしました。エージェントがコミットを書いたとき、同じコードベースに人間のコミットよりも、エージェントがバグを植え付ける可能性が高かったでしょうか?
短いバージョン: いいえ。どちらかといえばその逆。そして、実際に着地したエージェントが書いたセリフは、人間が書いたセリフよりも長く残り続けました。この投稿の残りの部分は、私が実際にそう信じてしまうような警告を含む長いバージョンです。
この質問の最初の問題は、「AI が書いた」ということは一つのことではないということです。慎重なリファクタリングを通じてクロード コードを駆動するシニア エンジニアと、スケジュールに従って PR を開く無人ボットはどちらも「AI コード」であり、これらを一緒くたにしてしまうと、存在するシグナルがすべて隠蔽されてしまいます。そこで、何かを測定する前に、eig を読み取るコミットごとの来歴検出器を構築しました。

さまざまなシグナル: ボット アカウント ID、サービス電子メール アドレス、コミット メッセージ フッター、共同作成者トレーラー、エージェント ブランチ プレフィックスや PR 本文マーカーなどのマージされた PR 証拠。次に、124 件のコミットを 6 人の独立したレビュー担当者に渡してブラインド検証したところ、精度は 96.2% で、8 つの検出チャネルのうち 6 つが完璧でした。 1 つの実際の障害モードは、人間がエージェントの PR 内でフォローアップ コミットをプッシュすることで、その信頼性がダウングレードされるため、フィルタリングできるようになります。
来歴を把握した上で、エージェントのコミットを 3 つの層に分割し、重要なことに、それらを決してプールしませんでした。
T1、ボットエージェント。ほぼ自律的で、直接のループに人間は関与しません。 Copilot のコーディング エージェント、Cursor のクラウド エージェントである Devin のことを考えてください。
T2、人間主導のエージェント。 Claude Code、Codex、およびその仲間たち。開発者が開発を進めながらレビューを行っています。これは、実際のエージェント コミットの圧倒的多数です。
T3、AI 支援。共著者の予告編だけで、他には何もなく、最も軽いタッチです。
階層の動作が異なるため、この階層が重要であることがわかります。このストーリーを正直に説明すると、階層を分離する必要があります。
「このコミットによりバグが発生した」ことを測定する方法
ここでの標準ツールは、Śliwerski、Zimmermann、Zeller にちなんで名付けられた SZZ です。この考え方は機械的です。バグを修正するコミットを見つけて、変更した行を gitBLame してその行に最後に触れたコミットを見つけ、それより前のコミットをバグを導入したコミットと呼びます。これをリポジトリ全体にわたって実行すると、バグを誘発するコミットとバグを誘発しないコミットのラベル付きデータセットが得られます。
各リポジトリ内で SZZ を実行し、同じコードベースでエージェントのコミットと人間のコミットを比較しました。これにより、一部のプロジェクトが他のプロジェクトよりもバグが多いという明らかな混乱を制御できました。次に、追加された行、削除された行、およびファイルのタッチを使用してロジスティック モデルを当てはめます。

コントロールとしてのヘッドに加えて、リポジトリの修正効果が加えられているため、すべての結果は「変更の規模がすでに説明している以上のバグのリスク」であることがわかります。このサイズ制御はオプションではありません。コミットによってバグが発生するかどうかを判断する唯一の最も強力な予測因子は、コミットの大きさです。コントロールをスキップすると、ほとんどの場合、エージェントが人間よりも大きいか小さい diff を書き込むかを測定することになります。
もう 1 つ規律があり、それが最終的に信頼できる結果と喜ばしい結果の違いとなりました。ナイーブな SZZ には、エージェントに有利なバイアスが組み込まれています。これにより、修正コ​​ミットはバグ誘発要素としてカウントされなくなり、エージェントは不釣り合いな量の修正を行うため、単純な方法でそれらを静かに保護します。これを把握するために、すべての結果に対して必須の感度チェックとして、より厳密なバリアント (B-SZZ) を実行しました。結果がフレンドリーなバリアントでのみ表示され、厳密なバリアントでは消えてしまう場合、それは本物ではありません。その境界線がどこに該当するかを以下で正確に説明します。
見出し: エージェントのコミットはバグを誘発しない
これは中心的な結果であり、同じリポジトリ内の人間のコミットと比較して、作成者層ごとにバグを導入するコミットの調整された確率です。 1.0 未満は、人間のベースラインよりもバグが少ないことを意味します。
著者層ごとにバグが発生する確率を調整し、変更のサイズとチャーンを制御
どの層も人間のライン以下に着地します。人間主導のエージェント (T2) のオッズ比は 0.57、95% 信頼区間は 0.42 ～ 0.76 であるため、区間全体は 1.0 未満になります。ボット エージェント (T1) は 0.75 [0.43, 0.95] です。最もライトタッチのアシスト層 (T3) は 0.96 [0.69, 1.08] で、これは人間と区別がつきません。いずれにせよ、大部分が人間によって作成されたコミットから期待されるものとまったく同じです。
今では正直さが伝わります。厳密な B-SZZ バリアントでは、T2 効果の減衰

null に向かって進みます。1.0 にギリギリ近づく [0.68, 1.01] の間隔で 0.79 に移動します。したがって、私はエージェントの「バグが 43% 少ない」という強い主張を支持しません。私が支持するのは、私たちがテストしたすべてのバリアントとすべての層に当てはまるため、より弱くてより永続的な主張です。同じリポジトリ内でエージェントのコミットが人間のコミットよりも多くのバグを引き起こすという証拠はなく、ポイント推定値は一貫して保護的であり、悪化することはありません。インターネットの声の大きい半分が確信していること、つまり AI が追加のバグを大量に送り込んでいることは、非難グラフには表示されません。
バグ誘発は 1 つのレンズです。もう 1 つは耐久性です。エージェントがコード行を書いたとき、そのコードはコードベースに残るのでしょうか、それとも 1 か月後に破棄されて書き直されるのでしょうか? 「エージェント コードはどうせすべて破棄される」というのは、バグ番号を受け入れる人であってもよくある意見なので、私たちはそれも測定しました。
リポジトリ全体でファイルをサンプリングし、現在の HEAD の 6 か月以上前に追加された行を取得し、まだ存在する部分を計算し、各リポジトリ内でペアにして、同じプロジェクトと同じ時代のエージェントと人間の行を比較しました。
HEAD までのライン生存、人間対人間主導のエージェント、リポジトリ内でペアリング
測定するのに十分な人間主導のエージェント ボリュームを備えたすべてのリポジトリでは、エージェント ラインは人間のラインよりも高い割合 (合計 17.9 パーセント ポイント) で生存しており、信頼区間はゼロから十分に離れていました。斜面はすべて同じように傾いています。調査の中での 1 つの例外は、1 つのリポジトリ内の 1 つのボット エージェント (T1) のデプロイメントです。このボット エージェントは、他のファイルの約 3 倍の頻度で自身のファイルを再修正し、当然のことながら、そのラインの大量生産が速くなります。導入パターンは、やはりエージェントのブランド名を上回っています。
雰囲気に反するこの結果がふさわしい

メカニズムは 3 つあると思います。
1つ目はレビューによる選定です。これらはマージされたコミットであり、プロジェクトが強制するレビューの制限をすでにクリアしていることを意味します。エージェントの PR が信頼できる人間のチームメイトよりも厳しく精査されること、また保護効果の一部が実際にはモデルの生の出力ではなくレビュー担当者の厳密性を測定していることは十分に考えられます。これは私が最も重要視している注意事項なので、もう一度説明します。
2 つ目は、エージェントが不釣り合いに保守要員であることです。エージェントのコミットが実際に行うことを特徴付けると、まったく新しい機能ではなく修正に大きく偏っています。修正形式のコミットは、小規模でローカライズされ、すでに存在して機能するコードに触れることであり、単にグリーンフィールド機能のコミットよりもリスクが低く、エージェントのボリュームの多くがまさにその形式です。
3 つ目は、おそらく過小評価されている退屈なものです。人間のコミットの多くもあまり優れていません。人間のベースラインは最高の状態の上級エンジニアではありません。それは、実際のリポジトリには必ず含まれる、慌ただしい金曜日、馴染みのないサブシステム、そして「ただ出荷するだけ」という現実的な組み合わせです。単に一貫性を保っているだけのエージェントであれば、そのハードルをクリアすることができます。
誰かがこれを読みすぎる前に私が提起したい警告
このような研究は、主張を拒否している事柄と同じくらい信頼できるものであるため、私は見出しよりもこれらのことを大きく取り上げたいと思っています。
これらはレビューされ、マージされたコミットです。この結果は、レビューされていないエージェントの出力がブランチに直接ダンプされたことについては何も述べておらず、保護効果の一部はおそらくモデルではなくレビューの厳格さによるものと考えられます。
検出されたエージェント コードのみが表示されます。エージェントの出力を貼り付け、マーカーを付けずに自分の名前でコミットする開発者は私たちには見えず、「人間」としてカウントされます。したがって、ここでのすべてのエージェント番号は上限ではなく下限です。
エージェント負荷の高いリポジトリ

若いね。エージェント時代は約 12 か月に及びます。これが、すべての比較がリポジトリ内または一致するコントロールに対して行われ、生のプールされたレートでは決してない理由であり、ウィンドウが 2 倍になる 6 か月後に再実行する必要がある理由です。
コーパスはランダムなサンプルではありません。それは帰属を維持するプロジェクトと AI に隣接したツールに偏っています。私たちは、カウンターウェイトとして退屈な非 AI プロジェクト (17 年前の Java フレームワーク、CMS、可観測性、およびチャット プラットフォーム) を意図的にオーバーサンプリングしましたが、人口全体の割合を主張しているわけではありません。
SZZはSZZです。責任ベースの誘導は、複雑なコミットや大規模なリファクタリングによるメソッドの既知の制限を継承しており、これを防ぐことはできますが、完全に排除することはできません。それが、厳密な B-SZZ バリアントがあらゆる番号に対応する理由のすべてです。
欠陥を予測する群衆にとってのボーナス発見
この作品からは、それ自体言及する価値のあることが 1 つ抜け落ちていました。それは、理解する前に私たちに刺さったからです。エージェントがほとんどのコミットを書き込むリポジトリでは、誰もがバグ、キーワード、ファイルレベルで「このファイルは修正によって修正された」ヒューリスティックにラベルを付ける安っぽい方法が崩壊します。修正 Firehose は非常に一定であるため、最も極端なリポジトリでは、ファイルの 86% が「バグがある」とカウントされ、予測子の精度がコイントスに近づくほど低下します。しばらくの間、これは「エージェント リポジトリの欠陥予測」のように見えました。
まったく、そうではありません。壊れるのはラベル付けのショートカットであり、基礎となる構造ではありません。コミットレベルの SZZ ラベルに移行し、明らかな自己修正チャーンをスパム解除すると、予測がすぐに戻ってきました。通常のシグナル (以前の修正、チャーン、所有権) はエー​​ジェント コードでは正常に機能します。これはモデリングの問題を装った測定の問題であり、Repowise のコード健全性スコア全体がその種のラベル付けを正しく行うことに基づいて構築されているため、この区別は私たちが非常に重視している点です。より深いバージョンが必要な場合は、

これらの履歴ベースの信号をどのように正直に保つかについて、プロセス メトリクスが構造メトリクスを上回る理由、およびヘルス スコアが実際にバグを予測するかどうかについての関連投稿は、両方とも漏洩トラップについて詳しく説明しています。
では、ロボットを信じるべきでしょうか?
もしあなたが判決を求めているのなら、AI エージェントがコードベースにバグを大量に送り込んでいるという懸念は、28 件の実際のプロジェクトの責任履歴によって裏付けられていません。 112,382 件のコミット全体で、エージェントのコミットは同じリポジトリ内の人間のコミットよりも欠陥を引き起こす可能性が低く、点推定値は保護的な傾向にあり、エージェントが作成した行は人間のコミットよりも長生きしました。この効果は本物ですが控えめであり、ほとんど完全に人間主導の層、つまり人がまだ差分をレビューしている層に存在します。最後の部分は脚注ではありません。データが伝えるストーリーは、「ロボットの方が優れている」ということではなく、「有能な人間とエージェントが、人間がコードを読み続けている状態で、人間が単独で作業する場合と少なくとも同等に耐えるコミットを生成する」ということです。正直に言うと、これは私たちのほとんどがすでに行っているワークフローです。
完全なレポート、リポジトリごとの数値、厳密バリアント感度テーブル、およびすべての分析スクリプト (独自のコーパスで再実行できる) は、Repowise ベンチマークで公開されています。

[切り捨てられた]

## Original Extract

We git-blamed 112,382 commits across 28 repos to test whether AI-agent code introduces more bugs than human code. After controlling for size, it doesn't, and its lines last longer.

Is AI-Written Code Buggier Than Human Code? repowise Docs Pricing Explore Blog Star — pip install repowise Sign in Blog / engineering Is AI-written code buggier than human code? We blamed 112,000 commits to find out
Ask ten engineers whether AI-written code is buggier than the code they write by hand and you'll get ten confident answers, roughly half in each direction, and zero of them backed by a number. It's one of those questions where everyone has a prior and nobody has the data. The takes are loud, the measurements are scarce, and the few measurements that exist tend to be a survey of how people feel about AI code rather than what the code actually does once it's merged.
So we measured it the way the defect-prediction field measures everything else: by going back through the git history of real projects and blaming every bug fix to the commit that introduced the bug. We did this across 28 public repositories and 112,382 commits spanning the first year of coding agents merging real pull requests, and then we asked a simple question of the data. When an agent wrote the commit, was it more likely to be the one that planted a bug than a human commit in the same codebase?
The short version: no. If anything, the opposite. And the agent-written lines that did land stuck around longer than the human-written ones. The longer version, with the caveats that make me actually believe it, is the rest of this post.
The first problem with this question is that "AI wrote it" isn't one thing. A senior engineer driving Claude Code through a careful refactor and an unattended bot opening PRs on a schedule are both "AI code," and lumping them together would hide whatever signal exists. So before measuring anything we built a per-commit provenance detector that reads eight different signals: bot account identities, service email addresses, commit-message footers, co-author trailers, and merged-PR evidence like agent branch prefixes and PR-body markers. We then blind-validated it, handing 124 commits to six independent reviewers, and it came back at 96.2% precision, with six of the eight detection channels perfect. The one real failure mode, a human pushing a follow-up commit inside an agent's PR, gets its confidence downgraded so we can filter it.
With provenance in hand we split agent commits into three tiers and, importantly, never pooled them:
T1, bot agents. Near-autonomous, no human in the immediate loop. Think Devin, the Copilot coding agent, Cursor's cloud agents.
T2, human-driven agents. Claude Code, Codex, and friends, where a developer is steering and reviewing as they go. This is the overwhelming majority of agent commits in the wild.
T3, AI-assisted. A co-author trailer and not much else, the lightest touch.
That tiering turns out to matter, because the tiers behave differently, and any honest version of this story has to keep them apart.
How you measure "this commit caused a bug"
The standard tool here is SZZ, named after Śliwerski, Zimmermann and Zeller. The idea is mechanical: find the commits that fix bugs, then git blame the lines they changed to find the commit that last touched those lines, and call that earlier commit the one that introduced the bug. Do this across a whole repo and you get a labeled dataset of bug-inducing and not-bug-inducing commits.
We ran SZZ within each repository, comparing agent commits to human commits in the same codebase, which controls for the obvious confound that some projects are just buggier than others. Then we fit a logistic model with the lines added, lines deleted, and files touched as controls, plus a repo fixed effect, so that every result reads as "bug risk beyond what the size of the change already explains." That size control is not optional. The single strongest predictor of whether a commit introduces a bug is how big the commit is, and if you skip the control you mostly end up measuring whether agents write bigger or smaller diffs than humans.
There's one more piece of discipline that ended up being the difference between a believable result and a flattering one. Naive SZZ has a built-in bias in favor of agents here. It excludes fix commits from being counted as bug-inducers, and agents do a disproportionate amount of fixing, so the naive method quietly shields them. To catch that, we ran a stricter variant (B-SZZ) as a mandatory sensitivity check on every single result. If a finding only shows up under the friendly variant and evaporates under the strict one, it isn't real. I'll tell you below exactly where that line falls.
The headline: agent commits are not more bug-inducing
Here is the core result, the adjusted odds of a commit introducing a bug, by authorship tier, relative to human commits in the same repo. Below 1.0 means fewer bugs than the human baseline.
Adjusted odds of introducing a bug by authorship tier, controlled for change size and churn
Every tier lands at or below the human line. Human-driven agents (T2) come in at an odds ratio of 0.57, with a 95% confidence interval of 0.42 to 0.76, so the whole interval sits below 1.0. Bot agents (T1) are at 0.75 [0.43, 0.95]. The lightest-touch assisted tier (T3) is at 0.96 [0.69, 1.08], which is just indistinguishable from human, exactly what you'd expect from commits that are mostly human-authored anyway.
Now the honesty pass. Under the strict B-SZZ variant, the T2 effect attenuates toward the null: it moves to 0.79 with an interval of [0.68, 1.01] that just barely touches 1.0. So I would not stand behind the strong claim that agents are "43% less bug-prone." What I will stand behind, because it holds across every variant and every tier we tested, is the weaker and more durable claim: there is no evidence that agent commits introduce more bugs than human commits in the same repo, and the point estimates are consistently protective, never worse. The thing the loud half of the internet is sure about, that AI is shipping a wave of extra bugs, is not visible in the blame graph.
Bug induction is one lens. Another is durability: when an agent writes a line of code, does it survive in the codebase, or does it get ripped out and rewritten a month later? "Agent code all gets thrown away anyway" is a common deflection even from people who accept the bug numbers, so we measured that too.
We sampled files across the repos, took the lines added six or more months before the current HEAD, and computed what fraction were still present, paired within each repo so we're comparing agent and human lines from the same project and the same era.
Line survival to HEAD, human versus human-driven agent, paired within repo
In every repo with enough human-driven-agent volume to measure, the agent lines survived at a higher rate than the human lines, by 17.9 percentage points in aggregate, with a confidence interval well clear of zero. The slopes all tilt the same way. The one exception anywhere in the study was a single bot-agent (T1) deployment in one repo that re-fixes its own files about three times as often as the others, and unsurprisingly its lines churn out faster. Deployment pattern, again, beats the agent's brand name.
A result this counter to the vibe deserves a mechanism, and I think there are three honest ones.
The first is selection by review. These are merged commits, which means they already cleared whatever review bar the project enforces. It is entirely plausible that agent PRs get scrutinized harder than a trusted human teammate's, and that some of the protective effect is really measuring reviewer stringency rather than the model's raw output. I'll come back to this, because it's the caveat I take most seriously.
The second is that agents are disproportionately the maintenance crew. When we characterized what agent commits actually do , they skew heavily toward fixes rather than net-new features. Fix-shaped commits, small, localized, touching code that already exists and works, are simply lower-risk than greenfield feature commits, and a lot of agent volume is exactly that shape.
The third is the boring one that's probably underrated: a lot of human commits are also not very good. The human baseline is not a senior engineer at their best; it's the realistic mix of rushed Fridays, unfamiliar subsystems, and "just ship it" that any real repo contains. An agent that is merely consistent can clear that bar.
The caveats I'd raise before anyone over-reads this
I want these louder than the headline, because a study like this is only as trustworthy as the things it refuses to claim.
These are reviewed, merged commits. The result says nothing about unreviewed agent output dumped straight to a branch, and some of the protective effect is likely review stringency, not the model.
We can only see detected agent code. A developer who pastes an agent's output and commits it under their own name with no marker is invisible to us and counts as "human." Every agent number here is therefore a floor, not a ceiling.
Agent-heavy repos are young. The agent era is about twelve months deep. That's why every comparison is within-repo or against matched controls, never a raw pooled rate, and why we owe a re-run in six months when the windows double.
The corpus is not a random sample. It skews toward projects that keep attribution and toward AI-adjacent tooling. We deliberately oversampled boring non-AI projects (a 17-year-old Java framework, a CMS, observability and chat platforms) as counterweight, but we're not claiming population-wide rates.
SZZ is SZZ. Blame-based induction inherits the method's known limits with tangled commits and large refactors, which we guard against but cannot fully eliminate. That's the whole reason the strict B-SZZ variant rides along with every number.
A bonus finding for the defect-prediction crowd
One thing fell out of this work that's worth its own mention, because it bit us before we understood it. In repos where agents write most of the commits, the cheap way everyone labels bugs, keyword and file-level "this file was touched by a fix" heuristics, falls apart. The fix firehose is so constant that in the most extreme repo, 86% of files counted as "buggy," which drags any predictor's accuracy down toward a coin flip. For a while that looked like "agent repos break defect prediction."
It isn't, quite. What breaks is the labeling shortcut , not the underlying structure. When we moved to commit-level SZZ labels and de-spammed the obvious self-fix churn, prediction came right back; the usual signals (prior fixes, churn, ownership) work fine on agent code. It's a measurement problem masquerading as a modeling problem, which is a distinction we care about a lot, since the whole Repowise code-health score is built on getting that kind of labeling right. If you want the deeper version of how we keep these history-based signals honest, the companion posts on why process metrics beat structural metrics and whether our health score actually predicts bugs both go into the leakage traps in detail.
So, should you trust the robot?
If you came for a verdict: the fear that AI agents are flooding codebases with bugs is not supported by the blame history of 28 real projects. Across 112,382 commits, agent commits were no more likely to introduce a defect than human commits in the same repo, the point estimates leaned protective, and agent-written lines outlived human ones. The effect is real but modest, and it lives almost entirely in the human-driven tier, the one where a person is still reviewing the diff. That last part is not a footnote. The story the data tells isn't "the robots are better," it's "a competent human plus an agent, with the human still reading the code, produces commits that hold up at least as well as that human working alone." Which, honestly, is the workflow most of us are already in.
The full report, the per-repo numbers, the strict-variant sensitivity tables, and every analysis script (so you can re-run it on your own corpus) are open in the Repowise benchmark re

[truncated]
