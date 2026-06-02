---
source: "https://zozo123.github.io/sandboxes-why-how-when/"
hn_url: "https://news.ycombinator.com/item?id=48356254"
title: "The Sandbox Shift – sandboxes are the new containers, for AI-written code"
article_title: "The Sandbox Shift — a field manual for running untrusted code"
author: "zozo123-IB"
captured_at: "2026-06-02T00:40:52Z"
capture_tool: "hn-digest"
hn_id: 48356254
score: 4
comments: 1
posted_at: "2026-06-01T12:59:22Z"
tags:
  - hacker-news
  - translated
---

# The Sandbox Shift – sandboxes are the new containers, for AI-written code

- HN: [48356254](https://news.ycombinator.com/item?id=48356254)
- Source: [zozo123.github.io](https://zozo123.github.io/sandboxes-why-how-when/)
- Score: 4
- Comments: 1
- Posted: 2026-06-01T12:59:22Z

## Translation

タイトル: サンドボックス シフト – サンドボックスは AI で書かれたコードの新しいコンテナーです
記事のタイトル: サンドボックス シフト — 信頼できないコードを実行するためのフィールド マニュアル
説明: Docker によりコードが移植可能になりました。サンドボックスを使用すると、AI が作成したコードを安全に実行できるようになり、エージェントにコンピューター全体を提供して、エンドツーエンドで開発、テスト、構築、出荷できるようになります。インタラクティブな決定機能を備えた、SWE、MLE、AI 研究者向けの、なぜ、いつ、どのように行うかを簡潔にまとめた現場マニュアル。

記事本文:
コンテンツにスキップ
封じ込めフィールドマニュアル · No.01 2026
サンドボックス
シフト。
Docker はコードを移植可能にしました。サンドボックスにより安全に実行できます —
作者はもう人間じゃないんだから。これはモデルであり、誰が読むよりも速くコードを記述します。
世代が安くなりました。安全な実行が新たなボトルネックとなっています。
サンドボックス · ユニット エージェントが内部で動作する使い捨てコンピュータ。ファイル、ネットワーク、機密情報へのアクセスが制御され、編集、実行、テスト、構築、出荷が可能で、その後消滅します。モデルを暴走させても十分安全です。ジョブを完了できる程度に完了していること。
コンテナ時代 · 2013 信頼できる人間がコードを書く
目標: 移植性と再現性。コードはレビューされ、保証されます。環境は長生きするペットです。
サンドボックス時代・現在 モデルがコードを書く
目標: 隔離、インスタント、使い捨て。コードは安全であることが証明されるまで有罪です。環境は、千単位で、はかない家畜です。
著者は信頼されていません。モデルで記述されたコードには、挿入されたペイロード、幻覚的な rm -rf / 、またはマルウェアに解決されるタイプミスされた依存関係が含まれる可能性があります。すべての実行を読むことはできません。
爆発範囲は制限されています。このボックスは、ファイル システム、ネットワーク出力、シークレット、および (分離が弱い場合は) ホスト カーネル自体など、コードが触れる可能性のあるものを決定します。
大規模に再現可能。実行ごとに同じクリーンな状態、何千もの同時実行、スポーンとキルのコストが低い。これがなければ、eval の報酬信号は単なるノイズになります。
封じ込めはまだ半分にすぎません。サンドボックスは、エージェントが内部で作業するコンピューター全体でもあり、開発、テスト、構築、展開を行うための 1 つの場所であり、人間が座っていなくてもタスクがエンドツーエンドで実行されます。
編集 → 実行 → テスト → ビルド → デプロイ → 観察 ↺ エージェントのループ - 自律的に何千回も繰り返されます
Docker はアーティファクトをパッケージ化します (ビルド、出荷、実行)。サンドボックスはエージェントに作業全体を任せます

フローとキー。コードを実行するだけのボックスを与えると、計算機が得られます。編集、構築、出荷できるものを与えれば、開発者が得られます。
人々が実際にたどり着く 4 つの仕事:
dev ハーネスのワークショップ コーディング エージェント (Claude Code とその友人たち) に、ラップトップやメイン ブランチに触れることなく、編集、インストール、実行、中断などの開発を行うためのボックスを提供します。
test · eval 動作することを証明する 単一の行を信頼する前に、あなたまたは AI が書いたばかりのコードを白紙の状態で実行し、評価します。
Ship AI が作成したコードをデプロイする モデルが作成したコードを、包含されたランタイム内の実稼働環境で実行します。爆発範囲は設計によって制限されます。
並列 10 個を一度に 1 つの環境を 10 通りにフォークし、エージェントに 10 個の機能とバグを同時に追跡させます。合格したものはそのままにしておきます。残りはビンに入れてください。
そして、同じプリミティブを役割ごとに選択します。
エージェントは PR を開いてシェル コマンドを実行するようになりました。人間が完全にレビューしていないコードを出荷します。
モデルで生成されたコードと SQL は、実際のデータセットとパイプラインに対して実行されます。
すべてのロールアウトとすべての評価では、信頼できないコードが一度に数千個実行され、報酬は再現可能でなければなりません。
境界決定
盗むものが何もない信頼できないコードは、外部、つまり公開された一時的なボックス、爆発半径が最も低い場所に属します。
データを必要とする信頼できないコードは VPC 内に属し、ハード的に隔離されています。これで、実際のネットワーク到達範囲で何かをフェンシングすることになります。
隔離はしご - 下部は弱くて速く、上部は強くて重いです。脅威モデルを保持する最も低い段を選択します。
サブプロセス + 制限 同じカーネル、同じユーザー。タイムアウトと ulimit 。信頼できるコードのみ。
名前空間 · cgroups · seccomp nsjail、bubblewrap。半信頼性コード用の安価なカーネルレベルのフェンシング。
コンテナrunc/containerd。共有カーネル — 1 つのカーネル CVE はエスケープです。
gVisor ユーザー空間カーネルが syscal をインターセプトする

ls。コンテナのエルゴノミクス、攻撃対象領域の縮小。
microVM · Firecracker 〜100 ミリ秒で起動する実際の VM 境界。ホストあたり数千回。エージェントのスイートスポット。
フル VM · エアギャップ 最大限の分離、最大限のコスト。本当に敵対的な人のために。
分離強度 ↔ ブート遅延 ↔ 密度 / コスト
microVM は、VM グレードの分離には時間がかかり、コストがかかるという古いルールを打ち破りました。これが、実行ごとの使い捨てを経済的に現実的なものにしています。
本当の教訓 Docker は分離性で LXC に勝てたわけではなく、開発者のエクスペリエンスで勝ったのです。サンドボックスも同様に勝ち得ます。つまり、最も厚い壁ではなく、最高の API と最速のブートです。
4つの質問。ライブ判定、推奨される段 (はしごを強調表示します)、および配置。ブラウザからは何も残りません。
…VM?孤立した観点から言えば、そうです。斬新な点は、約 100 ミリ秒で起動し、ホストごとに数千をパックすることです。その経済性により、ロールアウトごとの使い捨てが可能になります。
…コンテナ？コンテナはホスト カーネルを共有します。 1 つのカーネル CVE はエスケープです。信頼できるコードには問題ありませんが、レビューされていないモデル出力には問題がありません。このギャップこそが、gVisor と microVM が存在する理由です。
…誇大宣伝？私たちは何十年にもわたって信頼できないコードを実行してきました。真実。変わったのは作成者と量です。コードは現在、機械で書かれ、レビューされることなく、人間が精査できるよりも早く生成されます。絶縁はエッジケースからデフォルトの基板に移行します。
広大で動きの速い市場 - 高速で使い捨ての API 主導の分離という同じ形状に収束しています。
デイトナ · E2B · Blacksmith · Tensorlake · exe.dev · Modal · islo.dev · …
どれを選択しても、このページのコンセプトは変わりません。 clubbox.sh が存在するので、盲目的に選択する必要はありません。すべてのプロバイダーで同じタスクを実行し、サンドボックスを交換できる商品に変えます。
完全開示 islo.dev はこれらのプロバイダーの 1 つであり、私たちのものです。あるのはそれだけです

町にいるクマが 1 匹います。 🐻
はるかに大きな規模で、再びコンテナ革命が起こります。 2026 年にサンドボックスなしでソフトウェアを出荷する場合、2016 年には Docker を使用しません。
サンドボックスシフト。静的ページ — ビルド、トラッキング、Cookie はありません。
信頼できないコードを意図したとおりに実行します。 · ソース · インテルのストーリーを反映したデザイン
開示。ヨッシ・エリアス著、
islo.dev (そう、クマ 🐻) と協力し、積極的に貢献している人
clubbox.sh : 上記のすべてのプロバイダーで実行されます。
なぜ、いつ、どのようにベンダー中立なのか。風景の名前には、私たちの名前も含まれます。

## Original Extract

Docker made code portable. Sandboxes make AI-written code safe to run — and give an agent a whole computer to dev, test, build and ship in, end to end. A concise field manual on why, when & how, for SWEs, MLEs and AI researchers, with an interactive decider.

Skip to content
CONTAINMENT FIELD MANUAL · No.01 2026
The Sandbox
Shift .
Docker made code portable . Sandboxes make it safe to run —
because the author isn't human anymore. It's a model, writing code faster than anyone can read it.
Generation got cheap. Safe execution is the new bottleneck.
sandbox · the unit A disposable computer an agent works inside — it can edit, run, test, build & ship there, with controlled reach into files, network & secrets, then vanish. Safe enough to let a model run wild; complete enough to let it finish the job.
Container era · 2013 Trusted human writes the code
Goal: portability & reproducibility. Code is reviewed and vouched for. Environments are long-lived pets.
Sandbox era · now A model writes the code
Goal: isolation, instant, disposable. Code is guilty until proven safe. Environments are ephemeral cattle, by the thousand.
The author is untrusted. Model-written code can carry an injected payload, a hallucinated rm -rf / , or a typo'd dependency that resolves to malware. You can't read every run.
Blast radius is bounded. The box decides what the code may touch — filesystem, network egress, secrets, and (if isolation is weak) the host kernel itself.
Reproducible, at scale. Identical clean state every run, thousands in parallel, cheap to spawn and kill. Without it, an eval's reward signal is just noise.
Containment is only half the story. A sandbox is also the whole computer an agent works inside — one place to dev, test, build and deploy, where it takes a task end to end with no human in the seat.
edit → run → test → build → deploy → observe ↺ the agent's loop — repeated autonomously, thousands of times over
Docker packaged the artifact : build, ship, run. The sandbox hands the agent the entire workflow — and the keys. Give it a box that can only run code and you get a calculator; give it one that can edit, build and ship, and you get a developer.
Four jobs people actually reach for them:
dev The harness's workshop Give the coding agent — Claude Code & friends — a box to develop in: edit, install, run, break things, without touching your laptop or your main branch.
test · eval Prove it works Run code you or your AI just wrote on a clean slate and grade it, before you trust a single line.
deploy Ship AI-written code Run model-authored code in production inside a contained runtime, where its blast radius is bounded by design.
parallel Ten at once Fork one environment ten ways and let agents chase ten features and bugs at the same time. Keep what passes; bin the rest.
And the same primitive, by role — pick yours:
Agents now open PRs and run shell commands. You ship code no human fully reviewed.
Model-generated code and SQL run against real datasets and pipelines.
Every rollout and every eval runs untrusted code — thousands at once — and the reward must be reproducible.
The boundary decision
Untrusted code with nothing to steal belongs outside — a public, ephemeral box, lowest blast radius.
Untrusted code that needs your data belongs inside the VPC , isolated hard — now you're fencing something with real network reach.
An isolation ladder — weak & fast at the bottom, strong & heavy at the top. Pick the lowest rung that holds your threat model.
subprocess + limits Same kernel, same user. A timeout and a ulimit . Trusted code only.
namespaces · cgroups · seccomp nsjail, bubblewrap. Cheap kernel-level fencing for semi-trusted code.
container runc / containerd. Shared kernel — one kernel CVE is an escape.
gVisor A user-space kernel intercepts syscalls. Container ergonomics, smaller attack surface.
microVM · Firecracker A real VM boundary that boots in ~100 ms, thousands per host. The agent sweet spot.
full VM · air-gapped Maximum isolation, maximum cost. For the genuinely hostile.
isolation strength ↔ boot latency ↔ density / cost
microVMs broke the old rule that VM-grade isolation must be slow and expensive — which is what makes per-run disposability economically real.
the real lesson Docker didn't beat LXC on isolation — it won on developer experience. Sandboxes get won the same way: the best API and the fastest boot, not the thickest wall.
Four questions. Live verdict, recommended rung (it highlights the ladder), and placement. Nothing leaves your browser.
…a VM? In isolation terms, yes. The novelty is booting it in ~100 ms and packing thousands per host — that economics is what makes per-rollout disposability possible at all.
…containers? Containers share the host kernel; one kernel CVE is an escape. Fine for code you trust, thin for unreviewed model output. That gap is why gVisor and microVMs exist.
…hype? we've run untrusted code for decades. True. What changed is authorship and volume: code is now machine-written, unreviewed, and generated faster than humans can vet. Isolation moves from edge case to default substrate.
A vast, fast-moving market — converging on the same shape: fast, disposable, API-driven isolation.
Daytona · E2B · Blacksmith · Tensorlake · exe.dev · Modal · islo.dev · …
Pick any of them and the concepts on this page don't change. crabbox.sh exists so you don't have to choose blindly — it runs the same task across every provider, turning the sandbox into a commodity you can swap.
full disclosure islo.dev is one of these providers — and it's ours. There's only one bear in town. 🐻
It's the container revolution again — at a far larger scale. Shipping software without sandboxes in 2026 is not using Docker in 2016.
The Sandbox Shift. Static page — no build, no tracking, no cookies.
Run untrusted code like you mean it. · source · design nods to intel-story
Disclosure. Written by Yossi Eliaz ,
who works with islo.dev (yes — the bear 🐻) and actively contributes to
crabbox.sh , which runs across every provider above.
The why / when / how is vendor-neutral; the landscape names names — ours included.
