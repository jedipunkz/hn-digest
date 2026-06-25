---
source: "https://atkatana.com/blog/hyper-loops-draft.html"
hn_url: "https://news.ycombinator.com/item?id=48667742"
title: "The Carwash Problem: Why Your IT Organization Isn't Ready for AI-Generated Code"
article_title: "The Carwash Problem: Why Your IT Organization Isn't Ready for AI-Generated Code — Andrew Katana"
author: "atkatana"
captured_at: "2026-06-25T01:40:02Z"
capture_tool: "hn-digest"
hn_id: 48667742
score: 1
comments: 0
posted_at: "2026-06-25T01:33:13Z"
tags:
  - hacker-news
  - translated
---

# The Carwash Problem: Why Your IT Organization Isn't Ready for AI-Generated Code

- HN: [48667742](https://news.ycombinator.com/item?id=48667742)
- Source: [atkatana.com](https://atkatana.com/blog/hyper-loops-draft.html)
- Score: 1
- Comments: 0
- Posted: 2026-06-25T01:33:13Z

## Translation

タイトル: 洗車の問題: IT 組織が AI 生成コードの準備ができていない理由
記事のタイトル: 洗車の問題: IT 組織が AI 生成コードの準備ができていない理由 — Andrew Katana

記事本文:
洗車の問題: IT 組織が AI 生成コードに対応できない理由
私の AI デスクトップ エージェントが先週自殺しました。それは完全に私のせいでした。
セットアップは簡単でした。安定性の問題を解決するために、NVIDIA ドライバーをダウングレードするように依頼しました。それは続いた。新しいドライバーが出力され、古いドライバーが入力されます。再起動します。画面が真っ黒のままです。
問題は典型的なサプライチェーンの不一致でした。古いドライバー セットは、エージェント自体が動作するために必要な GPU アクセラレーションのデスクトップ環境をサポートしていませんでした。再起動時に、エージェントのプロセス マネージャーが依存していた表示スタックが機能しなくなったため、エージェントのプロセス マネージャーを初期化できませんでした。再起動前に最後に強制終了したプロセスは、再起動が必要なプロセスでした。
洗車の問題、つまりエージェントが自身の動作条件を破壊する方法でシステムを変更するという教科書的な例が必要な場合は、そこにあります。
これは悪いAIに関する話ではありません。これは、AI 主導の反復的な変化が、それを吸収するように設計されていないシステムに遭遇したときに何が起こるかについての物語です。
建築家のジョン・オースターハウトはかつて「少しの傾斜が複雑さを劇的に増大させる可能性がある」と書いています。 AI 支援開発は、その観察結果を運動法則に変えました。エージェントは 1 つの完璧なソリューションを作成するのではなく、反復します。彼らはループします。彼らは何かを試し、テストし、失敗し、調整し、再度試みます。ソリューションが実稼働品質に達するまでに、20、30、場合によっては 50 回の反復が必要です。
私はこのパターンをハイパー ループと呼んでいますが、もっと良い名前を付けてもいいと思っています。これは、ベースラインを設定し、解決策が見つかるまで繰り返し、その解決策をループが繰り返される次の組み立て段階に進めるプロセスです。各ステージはバンドル、テスト、および上方への昇格を行います。その結果、パイプラインは滝ではなく、むしろ粒子加速器のように見えます。

アーター。
重要な点は、このループには IT が必要ないということです。それは、企業資産やデータ上で、ユーザーが意識することなく実行されることが増えています。ラップトップと無料枠 AI アカウントを持つ平均的なナレッジ ワーカーは、すでにインフラストラクチャ上で次のブレークスルーを準備しています。
シングルスレッドのボトルネックは死んだ
エンタープライズ IT は過去 30 年間、生産手段の統合に費やしてきました。開発はITを介して行われます。導入は IT を介して行われます。変更管理は IT を通じて行われます。これはシングルスレッド モデルです。つまり、1 つのレビュー パイプライン、1 つの承認チェーン、1 つのコントロール セットです。
ソフトウェア成果物のソースは、IT を介したシングルスレッドではなくなりました。これはマルチスレッドで分散されており、ほとんど制御できないエンドポイント マシン上で実行されます。問題はこれをどうやって止めるかではありません。それは不可能ですし、試みるべきではありません。問題は、新しい現実に適合した受け入れプロセスをどのように構築するかです。
必要なのはテスト ハーネスです。これは、アーティファクトが公式開発チーム、請負業者、またはマイクロサービスの動作を促す方法を考え出したビジネス アナリストからのものかどうかに関係なく、本番環境と本番展開の間に位置する追加の監査ループです。
これは、エンタープライズ IT 用語で呼ばれる、ハーバード ビジネス レビューのオベリスク モデルです。コンサルティング業界は、AI によって、これまで若手スタッフのピラミッドが必要であったものを上級チームが提供できるようになることをすでに認識しています。同じ原則が内部にも当てはまります。つまり、組織はあらゆるレベルのソフトウェアを生産するようになりました。好むと好まざるにかかわらず、あなたもソフトウェア レビュアーになりました。
私がこの変化を乗り越えるのを見てきたすべての組織が、同じ 3 つの壁にぶつかりました。
1. セキュリティ テレメトリと爆発範囲の封じ込め
ほとんどの企業のデフォルトのセキュリティ体制は、既知の一連の開発を前提としています。

既知のパイプラインを通じてコードを生成するオペレーター。誰でもデプロイ可能な成果物を作成できるようになると、その前提は崩れます。
答えは、前線のゲートを厳しくすることではなく、端での封じ込めを強化することです。広角の可視性を実現するテレメトリ領域の監視ソリューションと、セーフティ ネットを提供するサーバーレス機能を探してください。目標は、冒険好きな従業員が何かを構築するのを妨げることではありません。これは、何かが壊れたときの爆発半径が街区ではなくインチ単位で測定されるようにするためです。
パイプラインのエントリ ポイントが 1 つしかなかったとき、システムのパッチ適用、依存関係の更新、およびソフトウェア サプライ チェーンのセキュリティはすでに困難でした。これに、パブリック レジストリ、ピン留めされていないバージョン、幻覚のパッケージ名から取得した AI 生成のアーティファクトをすべて掛け算します。
サプライ チェーンは、ますます実行可能な攻撃ベクトルとなっています。外部依存関係を引き出すハイパーループ内のすべてのループは、潜在的な妥協点となります。組織には、そうでないことが証明されるまで AI によって生成されたアーティファクトを高リスクとして扱う自動化された依存関係スキャン、アーティファクト レベルで強制される厳密な固定ポリシー、宣言されたものだけでなくどの依存関係が実際に使用されたかを証明できる実行時認証が必要です。
コスト面は変化しました。トークンの最大化（過剰な推論を実行して総当たりで解決すること）は市場によって修正されていますが、修正によって標準が生成されるわけではありません。これにより、ルーティング、優先順位付け、支出管理ソリューションが急増しています。
OpenRouter、Cloudflare AI Gateway、その他多数の企業が、AI によって生成された作業の課金プレーンとなるために競合しています。問題は、統一されたコストモデルがなければ、各チームの実験がクレジットカードを使ったシャドーITになってしまうことだ。答えは実験を禁止することではなく、実験をviにすることです。

予算はソースで決定されます。
ハイパーループ モデルは企業 IT にとって脅威ではありません。これは、IT が今後 10 年間に果たす最も重要な機能です。
単一の生産ソースとしての IT から、分散生産のための安全な動作環境のプロバイダーとしての IT への移行はすでに進行中です。この問題をうまく乗り越える組織は、企業がプロセス中に自らを破壊することなくソフトウェアを生産できるようにするテスト ハーネス (監査ループ、承認ゲート、ガードレール) を構築する組織になります。
代替案は、すべてのチームが独自のループ、独自の依存関係、独自のコスト モデル、および独自のセキュリティ体制を実行し、IT 部門がそのいずれかを知るのは何かが壊れた場合のみである未来です。
解決策を検討できるのは、何かが壊れていて注意が必要な場合だけです。それは変わらなければなりません。一貫した基準に照らしてレビューおよびチェックするための包括的な方法がなければ、これは困難な作業になるでしょう。
良いニュースです。私たちはアーキテクチャ パターンを持っています。オベリスク、ハイパーループ、テスト用ハーネス、これらは新しいアイデアではありません。これらは、ほとんどの企業が認めるよりも早く進行している問題に適用する必要があるだけです。
洗車の問題は現実です。唯一の問題は、AI エージェントが再起動して黒い画面になる前に、組織がこれに対処するシステムを構築しているかどうかです。
自宅のラボで構築され、地元のモデルを使用し、Andrew Katana が所有しています。

## Original Extract

The Carwash Problem: Why Your IT Organization Isn't Ready for AI-Generated Code
My AI desktop agent killed itself last week — and it was entirely my fault.
The setup was straightforward: I asked it to downgrade my NVIDIA drivers to fix a stability issue. It followed through. New drivers out, old drivers in. Reboot. Screen stays black.
The problem was a classic supply chain mismatch. The older driver set didn't support the GPU-accelerated desktop environment the agent itself needed to operate. On reboot, the agent's process manager couldn't initialize because the display stack it depended on was no longer functional. The last process it killed before the reboot was the one it needed to restart.
If you want a textbook example of the carwash problem — an agent modifying a system in a way that destroys its own operating conditions — there it is.
This is not a story about bad AI. It is a story about what happens when iterative, AI-driven change meets a system that was never designed to absorb it.
The architect John Ousterhout once wrote that "a little bit of slop can dramatically increase complexity." AI-assisted development has turned that observation into a law of motion. Agents don't write one perfect solution — they iterate. They loop. They try something, test it, fail, adjust, and try again. Twenty, thirty, sometimes fifty iterations before a solution reaches production quality.
I've been calling this pattern the Hyper-Loop — though I am open to a better name. It is a process that sets a baseline, iterates until a solution is found, then promotes that solution to the next assembly stage where the loop repeats. Each stage bundles, tests, and promotes upward. The result is a pipeline that looks less like a waterfall and more like a particle accelerator.
The critical detail: this loop does not require IT. It runs on your corporate assets, on your data, increasingly without your awareness. Your average knowledge worker with a laptop and a free-tier AI account is already cooking up their next breakthrough on your infrastructure.
The Single-Threaded Bottleneck Is Dead
Enterprise IT has spent the last thirty years consolidating the means of production. Development goes through IT. Deployments go through IT. Change management goes through IT. It is a single-threaded model — one review pipeline, one approval chain, one set of controls.
The source of software artifacts is no longer single-threaded through IT. It is multi-threaded, distributed, and running on endpoint machines you barely control. The question is not how to stop this — you cannot, and should not try. The question is how to build an acceptance process that matches the new reality.
What is needed is a testing harness — an additional audit loop that sits between production intent and production deployment, regardless of whether the artifact came from the official development team, a contractor, or a business analyst who figured out how to prompt their way to a working microservice.
This is the Obelisk model from the Harvard Business Review called out in Enterprise IT terms. The consulting industry has already recognized that AI enables senior teams to deliver what used to require pyramids of junior staff. The same principle applies internally: your organization now produces software from every level. You are now a software reviewer, like it or not.
Every organization I've watched navigate this shift hits the same three walls.
1. Security Telemetry and Blast Radius Containment
The default security posture for most enterprises assumes a known set of developers producing code through a known pipeline. When anyone can produce a deployable artifact, that assumption collapses.
The answer is not tighter gates at the front — it is better containment at the edge. Look for monitoring solutions in the telemetry space for wide-angle visibility, and serverless functions for providing a safety net. The goal is not to prevent adventurous employees from building things. It is to ensure that when something breaks, the blast radius is measured in inches, not city blocks.
System patching, dependency updates, and software supply chain security were already hard when the pipeline had one entry point. Now multiply that by every AI-generated artifact that pulls from public registries, unpinned versions, and hallucinated package names.
The supply chain is an increasingly viable attack vector. Every loop in the Hyper-Loop that pulls an external dependency is a potential compromise point. Organizations need automated dependency scanning that treats AI-generated artifacts as higher-risk until proven otherwise, strict pinning policies enforced at the artifact level, and runtime attestation that can prove what dependencies were actually used — not just what was declared.
The cost surface has shifted. Token maxing — running excessive inference to brute-force a solution — is being corrected by the market, but the correction is not producing a standard. It is producing a proliferation of routing, prioritization, and spend-management solutions.
OpenRouter, Cloudflare AI Gateway, and a dozen others are competing to be the billing plane for AI-generated work. The problem is that without a unified cost model, each team's experiments become shadow IT with a credit card. The answer is not to ban experimentation — it is to make experimentation visible and budgeted at the source.
The Hyper-Loop model is not a threat to Enterprise IT. It is the most important function IT will serve in the next decade.
The migration from IT as the single source of production to IT as the provider of a safe operating environment for distributed production is already underway. The organizations that navigate it successfully will be the ones that build the testing harness — the audit loop, the acceptance gate, the guardrails — that lets the business produce software without destroying itself in the process.
The alternative is a future where every team runs its own loop, its own dependencies, its own cost model, and its own security posture, and IT only finds out about any of it when something breaks.
I only get to look at the solution when something is broken and needs attention. That has to change. Without an overarching method to review and check against a consistent standard, this is going to be a rough ride.
The good news: we have the architectural patterns. The Obelisk, the Hyper-Loop, the testing harness — these are not new ideas. They just need to be applied to a problem that is moving faster than most enterprises are ready to admit.
The carwash problem is real. The only question is whether your organization builds the systems to handle it before your AI agent reboots into a black screen.
Built on a home lab, powered by local models, and owned by Andrew Katana .
