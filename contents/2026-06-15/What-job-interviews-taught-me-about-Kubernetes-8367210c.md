---
source: "https://notnotp.com/notes/what-job-interviews-taught-me-about-kubernetes/"
hn_url: "https://news.ycombinator.com/item?id=48546428"
title: "What job interviews taught me about Kubernetes"
article_title: "What job interviews taught me about Kubernetes"
author: "chmaynard"
captured_at: "2026-06-15T21:55:50Z"
capture_tool: "hn-digest"
hn_id: 48546428
score: 7
comments: 0
posted_at: "2026-06-15T20:12:37Z"
tags:
  - hacker-news
  - translated
---

# What job interviews taught me about Kubernetes

- HN: [48546428](https://news.ycombinator.com/item?id=48546428)
- Source: [notnotp.com](https://notnotp.com/notes/what-job-interviews-taught-me-about-kubernetes/)
- Score: 7
- Comments: 0
- Posted: 2026-06-15T20:12:37Z

## Translation

タイトル: 就職面接で Kubernetes について教えられたこと
説明: だから私は

記事本文:
就職面接で Kubernetes について教えられたこと
就職面接で Kubernetes について教えられたこと
それで最近は就職活動をしています。求人情報を読んだり、面接を受けたり、
十数社のエンジニアリングチームと話をしています。そして私は気づいた
最後にこれをやっていた 5 年前と比べて何か: 文字通り
今では全員が Kubernetes を使用しています。私が話をしたすべての企業。
前回の就職活動では全くそんなことはありませんでした。ありました
基本的には 3 つの陣営に分かれます: 稀な Kubernetes 採用者、systemd -on-VM/VPS/EC2 群衆、そしてサーバーレス派 (Lambda、
Cloud Run など）。
私が働いている場所には実際にビッグテック規模の企業があるので、これには驚きました。
したがって、K8s は私たちにとって明らかに理にかなっています。しかし、10 人のスタートアップでは、
2つのサービス？これらの場所はどれもマイクロサービスなどを行っていませんでした
ハイスケールに近い。そこで私はその理由を尋ねました。
ネタバレ: 彼らは技術的な側面にはあまり関心がありません
K8s。
技術面の面接は、特に理由を尋ねるのに最適な場所です。
CTO と直接話しているとき。それで私はそうしました。答えは次のとおりです。
基本的にはどこでも同じです。
1つ目は均一性です。すべてのサービスは同じ方法でデプロイされます。いいえ
決済サービスがベア VM 上で実行されていることを密かに知っている人
API が Docker Compose 上にある間、2019 年からの呪われた bash スクリプト
誰もそれに触れたことはありません。 1 つの方法ですべてをデプロイできます。
2 つ目は、共有された雇用可能な知識です。 K8s は基本的に言語です
今はフランカ。現在の仕事に就いた初日、私は次のようなリポジトリを作成しました。
Helm チャートと Kube 構成があり、全体像をしっかりと把握していました
1時間以内に建築。知識は YAML にあり、固定されているわけではありません。
誰かの頭。誰かを失った場合、代わりの人は 3 週間もかかりません
ドキュメントを調べて、何かがどのように実行されるかを理解しようとします。
今の会社ではオンカルで

l SRE は、次のような場合でもあらゆるサービスを稼働し続けることができます。
彼らはこれまで一度も触ったことがありません。彼らは Kubernetes と Kubernetes を知っています
パターンはどこのチームでも同じです。でそれをやってみてください
すべてのサービスが異なるように設定されている VM の束。 （注意：これだけです
もちろん、誰もエキゾチックなセットアップをしなかったとしても、これは当てはまります。)
3 番目はトレーサビリティ (コンプライアンスの有無にかかわらず) でした。私の現状では
会社、誰も kubectl apply -f 何かを直接適用することはできません
クラスターに。 Helm チャートを git にプッシュすると、トレースがあり、
MR 承認プロセスの後、FluxCD または ArgoCD が実際の展開を処理します。
影の中では何も起こりません。これはコンプライアンスと非常によく調和しています。
基本的には、これが ISO 認証を取得する方法です。そしてGitOpsペアなので
もちろん、Kubernetes を使用すると、これらすべてをほぼ無料で利用できます。
私が話を聞いたCTOたちは愚かな選択をしているわけではありません。彼らは現実を解決している
問題。
私は技術面だけに集中していましたが、久部は常に
私にとっては、技術的な問題に対する技術的な解決策です。でもたくさんあるみたいだね
の CTO は主に技術以外のメリットに関心を持っています。私よりも
と思った。彼らの技術的な問題にはそれが必要ありません。きっとそうしないでしょう
マニフェスト内でトポロジSpreadConstraintsが見つかりません。
気にする。 HPA、ポッド中断バジェット、ノード アフィニティ ルールはありません。ただ、
それ以外の場合に VM を持つノードの数と同じになります。しかし彼らは支払いを受け入れた
組織が複雑なソフトウェアを使用することの代償
利点。
正直、ほとんど大丈夫だと思います。しかし、それでもほとんどの企業はそう思います
それなしで始めるべきです。クラスターは、問題が発生した場合にデバッグするのが本当に困難です
うまくいかないとき、その段階では、製品ではなく製品にエネルギーを注ぐ必要があります。
インフラ。まだ次の大口顧客に売り込んでいるとき、
VPS とダーティ git pull の実行は完全に有効です
エメ

緊急修正。確かに最適ではありません。しかし、素早く、何が起こっているかを正確に知ることができます
起こっていること。なぜそうなるのかを理解するのに 2 時間も費やしたくないでしょう。
顧客の直前にポッドが CrashLoopBackOff でスタックする
電話する。
最近変化が起こった理由
なぜ変化が起こったのか、私はまだ完全には理解できません。 5年
以前は3つのキャンプはすべてうまくいっていました。現在、VM+ systemd 群衆
求人情報からは基本的に消え、サーバーレスはニッチなままであり、
K8sが勝ったばかりです。
私の最善の推測: マネージド K8 (EKS、GKE、AKS) は成熟し、才能のある
プールが反転しました: 十分な数の人々が、他のもののために雇用する必要があることを学びました
より難しい選択となった。そしてヘルムは「他人のチャートを使えばいい」としました。
本当のオプション。しかし、確信はありません。あなたがシフトのためにそこにいて、
もっと良い理論があれば、本当に知りたいです。
私の個人的な基準は、CTO が唯一のエンジニアではなくなった瞬間です
もう。 2 人目が現れるとすぐに、K8s は問題を解決します
現実になる。サーバーをセットアップしていない人がいますが、
導入する必要があります。 SSH キーではなく、適切なアクセス制御が必要な人
すべて。最終的には去って、知っていることをすべて持っていく人
彼らと一緒に。このとき、システムに知識を保持させたいのではなく、
人々。

## Original Extract

So I

What job interviews taught me about Kubernetes
What job interviews taught me about Kubernetes
So I've been job hunting lately. Reading job postings, doing interviews,
talking to engineering teams at like a dozen companies. And I noticed
something compared to five years ago when I was last doing this: literally
everyone is on Kubernetes now. Every single company I talked to.
Last time I was job hunting that wasn't the case at all. There were
basically three camps: the rare Kubernetes adopters, the systemd -on-VM/VPS/EC2 crowd, and the serverless people (Lambda,
Cloud Run, etc.).
That surprised me, because where I work we have actual Big Tech-scale
problems, so K8s makes obvious sense for us. But a 10-person startup with
two services? None of these places were doing microservices or anything
close to high scale. So I asked why.
Spoiler: they don't care much about the technical side of
K8s.
A technical interview is actually a great place to ask why, especially
when you're talking directly to the CTO. So I did. The answers were
basically the same everywhere.
First one was uniformity . Every service deploys the same way. No
one secretly knowing that the payments service runs on some bare VM with a
cursed bash script from 2019 while the API is on Docker Compose because
nobody ever touched it. One way to deploy, for everything.
Second was shared, hireable knowledge . K8s is basically a lingua
franca now. My first day at my current job, I pulled up the repo with the
Helm charts and Kube configs and had a solid picture of the whole
architecture within an hour. The knowledge is in the YAML, not stuck in
someone's head. Lose someone, their replacement isn't spending three weeks
digging through docs trying to figure out how anything runs.
At my current company, on-call SREs can keep any service up even if
they've never touched it before. They know Kubernetes, and Kubernetes
patterns are the same everywhere for all teams. Try doing that with a
bunch of VMs where every service is set up differently. (Caveat: this only
holds if nobody went exotic with the setup, of course.)
Third was traceability (with or without compliance). At my current
company, nobody can just kubectl apply -f something straight
to the cluster. You push a Helm chart to git, there's a trace, there's an
MR approval process, then FluxCD or ArgoCD handles the actual deployment.
Nothing happens in the shadow. That composes really well with compliance:
it's basically how we ace ISO certifications. And since GitOps pairs
naturally with Kubernetes, you get all of that almost for free.
The CTOs I talked to aren't making a dumb choice. They're solving real
problems.
I was focused on the technical side only, and Kube always has been a
technical solution to technical problems, for me. But it looks like a lot
of CTOs are interested primarily in the non-tech benefits. More than I
thought. Their technical problems just don't require it. I bet you won't
find any topologySpreadConstraints in their manifests, they don't
care. No HPA, no Pod Disruption Budgets, no node affinity rules. Just the
same number of nodes they'd have VMs otherwise. But they accepted to pay
the price of having a complex piece of software for the organizational
benefits.
Honestly, I think it's mostly fine. But I still think most companies
should start without it. Clusters are genuinely hard to debug when stuff
goes wrong, and at that stage you want your energy on the product, not the
infra. When you're still pitching to your next big customer, spinning up a
VPS and doing a dirty git pull is a totally valid
emergency fix. Suboptimal, sure. But fast, and you know exactly what's
happening. You really don't want to spend two hours figuring out why your
pod is stuck in CrashLoopBackOff right before a customer
call.
Why the shift happened recently
I still don't totally get why the shift happened when it did. Five years
ago all three camps were doing fine. Now the VM+ systemd crowd
has basically disappeared from job postings, serverless stayed niche, and
K8s just won.
My best guesses: managed K8s (EKS, GKE, AKS) got mature and the talent
pool flipped: enough people learned it that hiring for anything else
became the harder choice. And Helm made "just use someone else's chart" a
real option. But I'm not certain. If you were there for the shift and have
a better theory, I'd genuinely like to know.
My personal threshold would be the moment the CTO isn't the only engineer
anymore. As soon as a second person shows up, the problems K8s solves
become real. Now you've got someone who didn't set up the servers but
needs to deploy. Someone who needs proper access controls, not SSH keys to
everything. Someone who'll leave eventually and take everything they know
with them. That's when you want the system to hold the knowledge, not
people.
