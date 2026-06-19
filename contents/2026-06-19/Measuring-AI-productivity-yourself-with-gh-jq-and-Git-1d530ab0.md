---
source: "https://getunblocked.com/blog/measuring-ai-productivity/"
hn_url: "https://news.ycombinator.com/item?id=48599704"
title: "Measuring AI productivity yourself with gh, jq, and Git"
article_title: "Measuring AI productivity: Why your numbers are flat and your rework is up"
author: "dennispi"
captured_at: "2026-06-19T16:03:30Z"
capture_tool: "hn-digest"
hn_id: 48599704
score: 1
comments: 0
posted_at: "2026-06-19T15:26:30Z"
tags:
  - hacker-news
  - translated
---

# Measuring AI productivity yourself with gh, jq, and Git

- HN: [48599704](https://news.ycombinator.com/item?id=48599704)
- Source: [getunblocked.com](https://getunblocked.com/blog/measuring-ai-productivity/)
- Score: 1
- Comments: 0
- Posted: 2026-06-19T15:26:30Z

## Translation

タイトル: gh、jq、Git を使用して AI の生産性を自分で測定する
記事のタイトル: AI の生産性の測定: 数値が横ばいで手戻りが増える理由
説明: エンジニアリング組織全体で AI の使用率は最大 65% 増加していますが、PR スループットは 7.76% しか増加していません。ここ

記事本文:
AI の生産性の測定: 数値が横ばいで手戻りが増加している理由 製品 成果 顧客 セキュリティ ブログ ドキュメント 価格 ログイン 始める デモの予約 製品 コンテキスト エンジン コーディング エージェントと MCP AI コード レビュー 開発者 Q&A 結果 マージ可能なコード 保存トークン 手戻りの削減 顧客 セキュリティ ブログ ドキュメント 価格 ログイン 始める デモを予約 すべての記事 AI の生産性の測定: 数値が横ばいで手戻りが増加している理由
エンジニアリング組織全体で AI の使用率は最大 65% 増加しましたが、PR スループットは 7.76% しか増加しませんでした。ここでは、生産性ダッシュボードがモーションと配信を区別できない理由と、gh、jq、git だけを使用して来週ベースラインとして設定できる 4 つの指標を説明します。
最近ソーシャル メディアを開くという不運に見舞われた人は、おそらくあちこちでソーシャル メディアを目にしたことがあるでしょう。トケノミクスが到来し、リーパーが登場しました。
トークンノミクスとは、私たちの多くが AI ツールを使い始め、山ほどのコードを生成し、リファクタリングを出荷し、エージェントが夜間作業している間に就寝することで生産性を感じている一方で、マージ可能な出力と運用環境へのソフトウェアの出荷が追いついていないという概念を指します。それが説明しているのはトークンの利回りです。この世代のすべてがビジネス、顧客、そしてあなた自身の正気にとってどのような結果をもたらすのでしょうか? Linux Foundation でさえ、Tokenomics Foundation と今後のカンファレンス Tokenomicon を発表しました。これは Necronomicon への素晴らしい感謝です。彼らは AI 予算に関するカンファレンスを死者の書にちなんで名付けましたが、これはまさに正しいと思います。
このようなジョークが存在するのは、請求書が本物であり、痛ましいからです。 AI は現在、エンジニアリング予算の中で最も急速に成長している項目の 1 つであり、Linux Foundation 自身の言葉を借りれば、「その支出を測定し管理する規律が追いついていない」とのことです。
DX は LA から 400 の組織調査を実施しました

2024 年から 2026 年初頭までを調査したところ、プル リクエストのスループットは 7.76% しか増加していないのに対し、AI の使用率は最大 65% 増加したことがわかりました。 METR は 5 月に 349 人の技術労働者を対象に調査を行い、その中には 87 人のソフトウェア エンジニアが含まれており、自己申告の中央値は価値のある仕事が 2 倍、作業のスピードが 3 倍高かったとのことです。しかし、METR 自身のスタッフは、調査のサブグループの中で最も低い推定値を示しました。彼らは、チームが認識ギャップの研究について読んで観察バイアスがあったためだと考えています。測定データを最もよく知っている人が自分の感覚を信じていないということは、何かを物語っています。
2025 年の別の METR 調査 では、16 人の経験豊富なオープンソース開発者が、それぞれのリポジトリで 246 の実際のタスクに取り組みました。各タスクは、AI の使用を許可または禁止するためにランダムに割り当てられました。開始前、彼らは AI によって 24% 高速化されると予測していました。終了後、彼らはこれにより 20% 速くなったと信じていました。
19% 長くかかりました。知覚の誤りは、実際の経験と接触しても生き残った。調査ではあなたを救うことはできません。
METR は、2026 年 2 月に、57 人の開発者と 800 を超えるタスクを対象とした別の研究でこの研究を追跡調査しました。彼らは、信頼区間がゼロを超えており、METR自体は「速度向上の何らかの証拠」しか主張していないものの、速度の低下が縮小し、場合によっては逆転することさえ発見した。この期間中に 3 つの重要なことが変わりました。モデルは改良され、Opus 4.5 は研究の途中で出荷されました。開発者はこれらのツールをさらに練習しました。そして、コーホート自体も変化し、新しく採用された 47 人の開発者が、より多様なリポジトリのセットに取り組むようになりました。これは研究を台無しにするものではなく、発見です。このフィールドのすべての数値はスナップショットです。特定のリポジトリを持つ特定の人々が、特定のモデルを使用して結果を達成します。借用した結論が会社に移されることはありません。それで唯一の

価値のある生産性の数値は、自分で測定するものです。
ダッシュボードに生産性が高いと表示されるのはなぜですか? #
制約のない AI はより多くのコードと大きな差分を生成するため、コストがレビューに向かうことは誰もが知っています。 Faros は、AI を高度に導入しているチームではレビュー時間が 91% 増加していると測定しました。つまり、PR が開いたときにカウントするスループット グラフは役に立ちません。 tokenmaxxing リーダーボードと同様、人々が AI を使い始める動機になるかもしれませんが、有用な結果は追跡されていません。懸念されるのは、このレビュー キューが十分にバックアップされる場合です。チームは当然のことを行います。彼らはレビューをやめます。別の Faros レポートによると、31.3% 多い PR がまったくレビューされずにマージされています。
CircleCI は別の興味深い事実を発見しました。チームの中央値は、メインブランチのスループットが 7% 低下した一方で、フィーチャーブランチのスループットが 15% 増加しました。機能ブランチのスループットが最大 50% 増加した上位 10% のチームでさえ、メインブランチのアクティビティはほぼ横ばいでした。そして、メインブランチのマージ成功率は現在 70.8% にとどまっており、ここ 5 年間で最低となっています。動きは活発ですが、配信はフラットです。そしてその間にあるすべてのトークンが燃えています。
「Jカーブだ」と彼らは叫ぶだろう。 「ディップの中で測定しているんですね！」 J カーブから抜け出すことがどのようなものかを定義しなければ、これは反証不可能なままになります。実際の J カーブの主張には、日付が記載された終了基準が必要です。そして、曲線がいつ変わるのか、そしてそれを証明する数字が何なのかを誰も書き留めないのであれば、それは有用なモデルではありません。それはトークンを燃やすグラフを使った言い訳です。
before がなければデルタを証明することはできません。そして、ほとんどのチームは、これらの新しいツールを展開した日に「以前」を削除しました。 「ベースラインは何でしたか?」という質問に最初に答えることが重要です。
各エンジニアを他のチームではなく、AI 以前の自分の歴史と比較する同じエンジニアのベースライン設定は便利です。

技術。こうすることで、在職期間、チーム変更、季節性、さらには自己選択などの混乱をなくすことができます (早期採用者はおそらく最も強力なエンジニアでした)。これが機能することの証拠は、同じ企業で AI ユーザーの PR スループットが自社のベースラインに対して前年比 30% 増加し、非 AI ユーザーの PR スループットが 5% 増加したという DX の金融サービスの事例から得られています。
来週のベースラインを取得するには、次のことを行います。各エンジニアについて、サイクル タイム、レビュー時間、マージ後 30 日以内の再作業、および欠陥回避率の 4 つの項目を追跡します。次に、AI 支援による作業を残りの作業に対してセグメント化します。必ずチームレベルで集計してください。個人のパフォーマンスには絶対に使用しないでください。エンジニアがこれが個人の生産性を追跡するためであると考えた瞬間、それは成績表のように見え、データは消滅します。
これらにはベンダーは必要ありません。コードが GitHub または GitLab 上に存在する場合、データはすでに存在します。以下の例のスタック全体では、GitHub CLI ( gh )、 jq 、 git 、および cron ジョブを使用します。
サイクル タイムは、プル リクエストを開いてからマージするまでの時間です。平均値ではなく中央値を使用します。そうしないと、1 つの古い PR によって分析が台無しになってしまいます。
bash gh pr list --statemerged --limit 100 をコピーします \
--json createdAt,mergedAt \
| jq '[.[] | (.mergedAt|fromdateiso8601) - (.createdAt|fromdateiso8601)]
|並べ替える
| {merged_prs: 長さ,
median_cycle_hours: ((.[length/2|floor]) / 3600 | Floor)}' この数値を混乱させる 2 つの要因: ドラフト PR とボット
これを行う必要があるのは、チームがドラフト PR を活用している場合のみです。上記のクロックは、PR の作成時に開始するように設計されています。チームが CI を実行したり、AI エージェントにリモートでレビューさせたり、進捗状況を共有したりするために早めにドラフトを開いた場合、作業がまだ半分しか終わっていない間にタイマーが実行されるため、サイクル タイムは実際より悪くなります。 GitHub は、ドラフトが準備完了に切り替わる瞬間を記録します。

r レビューして、代わりにそこから測定できるようにします。
bash gh apigraphql -f owner="$(gh repo view --json owner -q .owner.login)" \ をコピーします。
-f repo="$(gh repo view --json name -q .name)" -f query='
query($owner: 文字列!, $repo: 文字列!) {
リポジトリ(所有者: $owner、名前: $repo) {
pullRequests(状態: MERGED、最後: 100) {
ノード {
作成日
マージ済み
timelineItems(itemTypes: READY_FOR_REVIEW_EVENT、最初: 1) {
ノード { ... on ReadyForReviewEvent { createdAt } }
}
}
}
}
}' \
| jq '[.data.repository.pullRequests.nodes[]
| ((.timelineItems.nodes[0].createdAt // .createdAt) | fromdateiso8601) as $start
| (.mergedAt | fromdateiso8601) - $start]
|並べ替える
| {merged_prs: 長さ,
median_cycle_hours_from_ready: ((.[length/2|floor]) / 3600 | Floor)}' ボット # あなたのチームはおそらく少なくとも dependabot を実行しているでしょう。そして、その PR は数分でマージされるか、場合によっては数か月間無視されたままになります。いずれにせよ、ここではロボットを測定したくない場合は、 --search "-author:app/dependabot" を使用してロボットを除外します。
1 つのクエリで 2 つの数値が得られます。コードが人間の目を待つ時間と、まったく目を通さずにマージされるコードの量。ファロスからの 31.3% を覚えていますか?自分で測定する方法は次のとおりです。
bash gh pr list --statemerged --limit 100 をコピーします \
--json createdAt,reviews \
| jq '{
合計: 長さ、
merged_with_zero_reviews: [.[] | select((.reviews|length)==0)] |長さ、
最初のレビューまでの時間の中央値:
([.[] | select((.reviews|length) > 0)
| ((.reviews |map(.submittedAt) |sort |first |fromdateiso8601)
- (.createdAt|fromdateiso8601))]
|並べ替え | (.[長さ/2|床]) / 3600 |床）
}' 注意点が 1 つあります。最初のレビューは必ずしも徹底的なレビューを意味するわけではありません。ゴム印による承認は、すべての行を読んだ人と同様にカウントされます。この分析では、実際のリビジョンを評価するのではなく、傾向を監視しようとしているため、これは問題ありません。

うわー。ただし、2 つの数字を一緒に見る必要があります。レビューの到着が速くなっている一方で、レビューがまったくない状態でマージされる PR が増えている場合は、速度が上がっていない可能性があります。単にレビューを停止しただけです。
リワークはさまざまな方法で測定できます。ここでは、Git History、Python、および cron のスプラッシュのみを必要とする高速なアプローチを紹介します。目標は、「今月削除したすべてのコードのうち、真新しいコードはどれくらいですか?」という質問に答えることです。古いコードが削除されるのは正常であり、祝われることです。 3 週間前のコードが削除されるということは、コードを書くのにお金を払い、レビューするのにお金を払い、そしてそれを置き換えるのに 3 回目のお金を払ったことを意味します。
測定の仕組み: 過去 30 日間のすべてのコミットを追跡します。コミットが削除した行ごとに、その行が何年前のものかを gitblame に尋ねます。削除されてから 30 日未満の場合は、やり直しとしてカウントします。若い欠失をすべての欠失で割ったものが割合です。これはダッシュボードで課金されるメトリクスであり、内部的には同じ 3 つの git コマンドのバージョンです。
このスクリプトを rework_rate.py としてリポジトリのルートに保存し、Python 3 で実行します。 警告の 1 つ: 削除された行のブロックごとに gitblame が 1 回実行されるため、大規模でビジーなリポジトリでは速度が遅くなります。そのため、月次 cron に入れてください。
注: デフォルトは 30 日です。チームのベロシティが高い場合は、14 日バージョンを実行してください。
rework_rate.py #!/usr/bin/env python3 をコピーします
"""再作業率: 削除されてから AGE 日以内に削除された行の割合。
使用法: python3 rework_rate.py [window_days] [age_days] (リポジトリのルートで実行)"""
インポート re、サブプロセス、sys
WINDOW = sys.argv[1] if len(sys.argv) > 1 else "30"
AGE = int(sys.argv[2]) if len(sys.argv) > 2 else 30
def git(*args):
return subprocess.run(["git", *args], capture_output=True, text=True).stdout
若い = 合計 = 0
git("log", f"--since={WIN) のエントリの場合

DOW} 日前", "--no-merges",
"--pretty=%H %ct").splitlines():
sha、ts = エントリ.split()
削除済み_at = int(ts)
for f in git("diff-tree", "--no-commit-id", "--name-only", "-r", sha).splitlines():
diff = git("diff", "-U0", f"{sha}^", sha, "--", f)
for m in re.finditer(r"^@@ -(\d+)(?:,(\d+))?", diff, re.M):
開始、カウント = int(m.group(1))、int(m.group(2) または 1)
カウント == 0 の場合:
続ける
blame = git("blame", "--line-porcelain", f"-L{start},+{count}",
f"{しゃ}^"、"--"、f)
re.finditer(r"^author-time (\d+)",blame,re.M)のaの場合：
合計 += 1
削除済みの場合 - int(a.group(1)) < AGE * 86400:
若い += 1
合計の場合:
print(f"削除された行: {合計} | 若い (<{AGE}d): {若い} | "
f"やり直し率: {若い/合計:.1%}")
それ以外の場合:
print("no deletions in window") モノリポジトリでこの投稿を書いているときにこれを実行しました。その後、836 件のコミットと 215,653 行の削除: 27%。先月削除した行の 4 行に 1 行は 30 日以内のものでした。 14 日のしきい値では 20.8% になります。これは、ほとんどのやり直しは AI ピル化されると迅速に行われるためです。
業界のベンチマークは存在せず、必要もないので、探しに行かないでください。すべてのリポジトリのチャーンは異なり、スクリプトはファイル名の変更に従わず、1 つの大きなフォーマット パスがスパイクとして読み取られます。それはどれも重要ではありません。

[切り捨てられた]

## Original Extract

Across engineering organizations AI usage is up ~65% but PR throughput rose only 7.76%. Here

Measuring AI productivity: Why your numbers are flat and your rework is up Products Outcomes Customers Security Blog Docs Pricing Log In Get Started Book a Demo Products Context Engine Coding Agents and MCP AI Code Review Developer Q&A Outcomes Mergeable Code Save Tokens Reduce Rework Customers Security Blog Docs Pricing Log In Get Started Book a Demo All articles Measuring AI productivity: Why your numbers are flat and your rework is up
Across engineering organizations AI usage is up ~65% but PR throughput rose only 7.76%. Here's why your productivity dashboard can't tell motion from delivery, and four metrics you can baseline next week with nothing but gh, jq, and git.
If you've had the misfortune of opening social media recently, you've probably seen it everywhere. Tokenomics has come and the Reapers are here.
Tokenomics refers to the concept that while many of us have started to use AI tools and are feeling productive by generating mountains of code, shipping refactors, and going to bed while agents work overnight, the mergeable output and shipped software to production has not kept pace. What it's describing is token yield: what is the outcome of all of this generation for the business, for your customers, and for your own sanity? Even the Linux Foundation has announced its Tokenomics Foundation as well as its upcoming conference Tokenomicon , a brilliant hat tip to the Necronomicon . They named the conference for AI budgets after the Book of the Dead, which feels about right.
These jokes exist because the bill is real and it hurts. AI is now one of the fastest-growing line items on an engineering budget and in the Linux Foundation's own words, "the discipline to measure and govern that spend has not kept pace."
DX ran a 400 organization study from late 2024 through early 2026 and found that pull request throughput was only up 7.76% while AI usage rose ~65%. METR surveyed 349 technical workers in May , 87 software engineers among them, and the median self-report was two times more valuable work and three times faster. However, METR's own staff gave the lowest estimates of any subgroup in their survey. They believe it's because their team has read about the perception gap research and had an observation bias. When the people who are most familiar with the measurement data don't trust their own feelings, that tells you something.
In another 2025 METR study , 16 experienced open source developers worked on 246 real tasks on their own repos. Each task was randomly assigned to either allow or forbid AI use. Before starting, they predicted that AI would make them 24% faster. After finishing, they believed it had made them 20% faster.
They took 19% longer. The perception error survived contact with the real experience. Surveys can't save you.
METR followed up that study with another in February 2026 with 57 developers and over 800 tasks. They discovered that the slowdown shrank or in some cases even reversed, though the confidence intervals cross zero and METR itself claims only "some evidence for speedup." Three key things had changed in this time period. The models got better, Opus 4.5 had shipped mid-study. Developers got more practice with these tools. And the cohort itself changed, with 47 newly recruited developers working on a more diverse set of repos. This isn't a knock on the research, it's the finding. Every number in this field is a snapshot. Specific set of people with a specific set of repos using specific models to accomplish their outcomes. Borrowed conclusions don't transfer into your company. So the only productivity number worth anything is the one you measure yourself.
Why does my dashboard say we're productive? #
We all know that unconstrained AI produces more code and bigger diffs, so the cost moves towards review. Faros measured that review time was up 91% in teams with high AI adoption. That means your throughput chart counting when the PR opens isn't helpful. Like a tokenmaxxing leaderboard, while it may motivate people to start using AI it isn't tracking to useful outcomes. What's concerning is when this review queue backs up far enough. Teams do the obvious thing. They stop reviewing. A separate Faros report found that 31.3% more PRs are merged with no review at all .
CircleCI found another interesting fact . The median team pushed 15% more throughput to feature branches while their main-branch throughput dropped 7%. Even the top 10% of teams, with feature branch throughput up ~50%, saw main-branch activity stay roughly flat. And main branch merge success now sits at 70.8%, the lowest in over 5 years. Motion is up, yet delivery is flat. And all the tokens in between are on fire.
"It's the J-curve," they'll shout. "You're measuring in the dip!" This will remain unfalsifiable if we never define what climbing out of the J-curve looks like. A real J-curve claim requires a dated exit criterion. And if no one will write down when the curve turns and what number proves it, it's not a useful model. It's an excuse with a graph that's leading you to burn tokens.
You can't prove a delta without a before. And most teams deleted their "before" the day they rolled out these new tools. It's critical that you first answer the question, "What was our baseline?"
Same-engineer baselining where you compare each engineer to their own pre-AI history, not to other teams, is a useful technique. This way it can kill confounds like tenure, team changes, seasonality, or even self-selection (early adopters were likely your strongest engineers). The proof that this works is from DX's financial services case where AI users had an increase of 30% PR throughput year over year against their own baseline and non-AI users had a 5% increase at the same company.
To get a baseline next week, here's what you do. For each engineer, track four items: Cycle time, review time, rework within 30 days of merge, and defect escape rate. Then segment AI-assisted work against the rest. Make sure you aggregate at the team level. Never use this for individual performance. The moment an engineer thinks that this is to track their personal productivity, it looks like a report card and your data will die.
None of this needs a vendor. If your code lives on GitHub or GitLab, the data already exists. The whole stack in the examples below uses GitHub CLI ( gh ), jq , git , and a cron job.
Cycle time is the time from opening a pull request to merging it. Use the median, not the mean, otherwise one stale PR will ruin your analysis.
Copy bash gh pr list --state merged --limit 100 \
--json createdAt,mergedAt \
| jq '[.[] | (.mergedAt|fromdateiso8601) - (.createdAt|fromdateiso8601)]
| sort
| {merged_prs: length,
median_cycle_hours: ((.[length/2|floor]) / 3600 | floor)}' Two things will mess with this number: Draft PRs and Bots
You only need to do this if your team leverages draft PRs. The clock above is designed to start when the PR is created. If your team opens drafts early to run CI, have AI agents review in remote, or share progress, the timer will run while that work is still half done and your cycle time will look worse than it is. GitHub records the moment a draft flips to ready for review so you can measure from that instead:
Copy bash gh api graphql -f owner="$(gh repo view --json owner -q .owner.login)" \
-f repo="$(gh repo view --json name -q .name)" -f query='
query($owner: String!, $repo: String!) {
repository(owner: $owner, name: $repo) {
pullRequests(states: MERGED, last: 100) {
nodes {
createdAt
mergedAt
timelineItems(itemTypes: READY_FOR_REVIEW_EVENT, first: 1) {
nodes { ... on ReadyForReviewEvent { createdAt } }
}
}
}
}
}' \
| jq '[.data.repository.pullRequests.nodes[]
| ((.timelineItems.nodes[0].createdAt // .createdAt) | fromdateiso8601) as $start
| (.mergedAt | fromdateiso8601) - $start]
| sort
| {merged_prs: length,
median_cycle_hours_from_ready: ((.[length/2|floor]) / 3600 | floor)}' Bots # Your team is likely running at least Dependabot. And its PRs either merge themselves in minutes or sometimes sit ignored for months. Either way, you don't want to measure a robot here, exclude them with: --search "-author:app/dependabot" .
One query that gets you two numbers. How long code waits for human eyes and how much of it merges with no eyes at all. Remember that 31.3% from Faros? This is how you measure your own:
Copy bash gh pr list --state merged --limit 100 \
--json createdAt,reviews \
| jq '{
total: length,
merged_with_zero_reviews: [.[] | select((.reviews|length)==0)] | length,
median_hours_to_first_review:
([.[] | select((.reviews|length) > 0)
| ((.reviews | map(.submittedAt) | sort | first | fromdateiso8601)
- (.createdAt|fromdateiso8601))]
| sort | (.[length/2|floor]) / 3600 | floor)
}' One caveat: first review doesn’t always mean a thorough review. A rubber-stamp approval counts the same as someone who read every single line. In this analysis, that's fine because you're trying to watch for a trend, not grading the actual reviews. But you should watch the two numbers together. If reviews are arriving faster while more PRs merge with no review at all, you may not be speeding up—you’ve just stopped reviewing.
Rework can be measured in many different ways, here is an approach that's fast and requires nothing but git history, python, and a splash of cron. The goal is to answer the question “of all the code that we deleted this month, how much was brand new?” Old code getting deleted is normal and celebrated. Three week old code getting deleted means you paid to write it, paid to review it, and then paid a third time to replace it.
How the measurement works: Walk every commit from the last 30 days. For each line a commit deleted, ask git blame how old that line was. If it was less than 30 days old when it was deleted, count it as rework. Young deletions divided by all deletions, that's the rate. This is the metric that dashboards charge for, and under the hood, it's some version of these same three git commands.
Save this script as rework_rate.py at the root of your repo and run it with Python 3. One word of warning: it’s slow on big and busy repos because it runs git blame once for every block of deleted lines, so put it in a monthly cron.
Note : 30 days is the default. If your team is high velocity, run a 14 day version.
Copy rework_rate.py #!/usr/bin/env python3
"""Rework rate: share of deleted lines under AGE days old when deleted.
Usage: python3 rework_rate.py [window_days] [age_days] (run at repo root)"""
import re, subprocess, sys
WINDOW = sys.argv[1] if len(sys.argv) > 1 else "30"
AGE = int(sys.argv[2]) if len(sys.argv) > 2 else 30
def git(*args):
return subprocess.run(["git", *args], capture_output=True, text=True).stdout
young = total = 0
for entry in git("log", f"--since={WINDOW} days ago", "--no-merges",
"--pretty=%H %ct").splitlines():
sha, ts = entry.split()
deleted_at = int(ts)
for f in git("diff-tree", "--no-commit-id", "--name-only", "-r", sha).splitlines():
diff = git("diff", "-U0", f"{sha}^", sha, "--", f)
for m in re.finditer(r"^@@ -(\d+)(?:,(\d+))?", diff, re.M):
start, count = int(m.group(1)), int(m.group(2) or 1)
if count == 0:
continue
blame = git("blame", "--line-porcelain", f"-L{start},+{count}",
f"{sha}^", "--", f)
for a in re.finditer(r"^author-time (\d+)", blame, re.M):
total += 1
if deleted_at - int(a.group(1)) < AGE * 86400:
young += 1
if total:
print(f"deleted lines: {total} | young (<{AGE}d): {young} | "
f"rework rate: {young/total:.1%}")
else:
print("no deletions in window") I ran this while writing this post on our monorepo. 836 commits and 215,653 deleted lines later: 27%. One in four lines we deleted last month was under 30 days old. At a 14-day threshold it’s 20.8%, because most rework happens fast once AI-Pilled.
Don't go looking for an industry benchmark because there isn't one and you don't need one. All repos churn differently, the script doesn't follow file renames and one big formatting pass will read as a spike. None of that matters because the

[truncated]
