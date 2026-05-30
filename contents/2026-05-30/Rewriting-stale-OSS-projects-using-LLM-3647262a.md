---
source: "https://loopholelabs.io/blog/rewriting-oss-in-the-ai-era"
hn_url: "https://news.ycombinator.com/item?id=48334593"
title: "Rewriting stale OSS projects using LLM"
article_title: "The Math Changed: Rewriting Stale Open Source In The AI Era"
author: "axod"
captured_at: "2026-05-30T11:33:52Z"
capture_tool: "hn-digest"
hn_id: 48334593
score: 2
comments: 0
posted_at: "2026-05-30T10:09:18Z"
tags:
  - hacker-news
  - translated
---

# Rewriting stale OSS projects using LLM

- HN: [48334593](https://news.ycombinator.com/item?id=48334593)
- Source: [loopholelabs.io](https://loopholelabs.io/blog/rewriting-oss-in-the-ai-era)
- Score: 2
- Comments: 0
- Posted: 2026-05-30T10:09:18Z

## Translation

タイトル: LLM を使用した古い OSS プロジェクトの書き換え
記事のタイトル: 数学が変わった: AI 時代の陳腐化したオープンソースの書き換え
説明: 確立されたオープンソース プロジェクトの書き換えには何年もかかりました。 AI では、それはもう当てはまりません。ここでは、CRIU を自社用に Zig で書き直す理由、そもそもインフラストラクチャ プロジェクトが陳腐化する原因、および再構築と貢献の計算がもう一度検討する価値がある理由を説明します。

記事本文:
数学が変わりました: AI 時代における陳腐化したオープンソースの書き換え { ∞ } ループホール Labs アーキテクト
~/ ブログ / エンジニアリング / 数学が変わりました
数学が変わりました: AI 時代の陳腐化したオープンソースの書き換え
投稿者: jimmy moore 、創設エンジニア このページについて:
オープンソース プロジェクトが陳腐化する仕組み
最近まで、確立されたオープンソース プロジェクトを書き直すには何年もかかることがありました。 LLM ではそれが変わります。私たちは
CRIU を Zig で書き直すと、数年ではなく数か月で完了すると予想されます。
住宅の改修を行ったことがある人なら、あらゆる作業に適したツールがあることを知っています。何年も経てば小屋がいっぱいになってしまう
それぞれが 1 つの仕事をうまくこなす専門的なもの。ツールが必要になるのは数回だけかもしれませんが、
その瞬間が来ても、他には何もできません。
私の家 私の家を修理するための道具
═========================================
┌─────────────────┐
│ ハンマーソードリル │
│ 千枚通しペンチ │
│ レベルスクエアアックス │
╱╲ │ ジグソーサンダーオイル │
╱ ╲ │ バールカンナバイス │
┌────┐ │ テープレンチ ラチェット │
│ ▢▢ │ │ スパナ・釘打ち機・千枚通し │
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
ソフトウェアも同じ話です。構築する際には、より迅速に作業を進め、より適切に出荷できるようにするための専用ツールを使用します。
そして、どちらの種類のツールもメンテナンスが必要です。実際に切断する必要がある場合、錆びたノコギリは価値以上に苦痛になります。
古いソフトウェア ツールを使用すると、節約できた以上に時間がかかる可能性があります。ソフトウェアの難点は、
通常、鋭さを保っている人ではありません。他の誰かがそうしている、そしてあなたは彼らがまだそうであると信じています。
ツールが鈍くなってきたことに気づいたときの標準的なアドバイスは、次のとおりです。

良きオープンソース市民になりましょう: バグを報告し、
パッチを作成し、プルリクエストを送信します。私たちのほとんどはそれを行っています、そして相互主義がエコシステムを維持しているのです
走っている。しかし、そのモデルには歪みが見え始めています。メンテナは AI が生成したパッチに静かに溺れています。
一目でわかるバグレポートと、
燃やす
時間未満
レビュー。
貢献上流の経済学は、人間が人間がレビューできるように思慮深いパッチを作成したときに機能しました。
同等の速度。その方程式の一方がマシン速度で実行され、もう一方がそうでない場合、ループ全体が
壊れる。このモデルがまだ機能するかどうかを尋ねる価値はあります。
チェックポイント/復元はアーキテクトの仕事です。私たちは、お客様の実行中の Linux プロセスを受け入れます。
メモリ、スレッド、開いているファイルとソケット - それらをフリーズし、取得し、コンピュータ上で復活させます。
別のマシン。 Linux には、これを行うツールが 1 つあります: CRIU 。そうだった
10 年以上にわたって存在し、runc、podman、およびより広範な Kubernetes エコシステムに組み込まれています。
どのチームも最初に到達するのは明白なことです。
オープンソース プロジェクトが陳腐化する仕組み
私たちが遭遇し続けた問題、そして積み重ねてきた回避策は、実際にはそうではありませんでした。
CRIUの問題。それらは構造的なものであり、長きにわたるオープンソース プロジェクトが最終的に遭遇する類のものでした。
そして、端の周りにいくらパッチを当ててもそれらを修正することはできませんでした。
人々は情熱、ほとんど強迫観念を持っています。私たちは皆、人生の中で、一年間完全に夢中になったことがあるでしょう。
その間、そしてそれから遠ざかりました。何十年どころか、何年にもわたってその強度を維持するのは本当に困難です。
プロジェクトの開始時は、すべてが新しいことです。コミットするたびに針が動き、メンタルモデルが頭の中に収まり、
そして上達は早くて楽しいです。プロジェクトとして

変化に応じて、仕事は変わります。メンテナンスに多くの時間が費やされます。
回帰狩りを行い、既存のコードが腐らないようにします。リノベーションの例えに戻ります。私たちは皆、
プロジェクトは棚に上げられましたが、最後の 10% は面白くない部分だったため、90% で消滅しました。
コードを手作業で記述する場合、人が物事を進めるためのエネルギーは有限です。本来の原動力
最終的には燃え尽きて次へ進み、誰が残したとしてもプロジェクトは実行され続けます。
長命プロジェクトは、常に変化し続ける世界にも存在します。 CRIU が 2012 年に発足したとき、Go はちょうど
1.0に達しました。 Rust は最初の安定版リリースから 3 年かかりました。あと4人いたらジグは存在しなかっただろう。 Linux カーネル
io_uring がありませんでした。私たちが知っている eBPF は存在しませんでした。コンテナは基板ではなくニッチな好奇心だった
世界のコンピューティングの半分。ビルド システム、CI、パッケージ管理、テスト規約、ドキュメント ツール -
すべてが移動しており、多くの場合、複数回移動しています。
プロジェクトがその環境に合わせてリファクタリングを続けないと、プロジェクトの周囲の世界は徐々に適合しなくなります。の
伝えるのは微妙です。 CRIU のホームページは 2026 年になっても MediaWiki がインストールされています。
それ自体は重要ですが、間違った10年の壁紙のように見えます。プロジェクトが停止したことを静かにマークします
改修中。
┌───────────────────────────┐
│ ファイル編集 ヘルプ表示 │
━━━━━━━━━━━━━━━━━━━━━━━━━━━━┤
│ ◄

► ⟳ ⌂ http://geocities.com/~user/website-i-have-not-updated-lately │
━━━━━━━━━━━━━━━━━━━━━━━━━━━━┤
│ │
│ ★私のホームページへようこそ★ 訪問者: [ 0 0 0 0 4 2 7 ] │
│ ⚠ 建設中 ⚠ [<アニメーションワーカー.gif>] │
│ │
│ ウェブリング: << 前へ |ランダム |次へ >> 最終更新日: 1998 年 4 月 │
│ ゲストブックにサインしてください！ Netscape で見るのが最適 │
│ │
━━━━━━━━━━━━━━━━━━━━━━━━━━━━┘
そのドリフトの一部は修正できます。 CRIU は、runc、podman、および Kubernetes が出現した後、それらに組み込まれました。
既存のコアが適切なインターフェイスを公開している場合、新しいオーケストレーション レイヤーをその上に配置できます。
手術。他の種類のドリフトはまったく修正できません。確立された C コードベースを Rust または Zig に移行しない
いくつかのプルリクエストとともに。実装言語、メモリ モデル、ディスク上のフォーマット、およびコア
建築は選択された瞬間にすべてフリーズします。世界が前進する頃には、それに伴う移動コストも
やり直すコストに近づきます。
第三の力が働いており、それは構造的なものです。大企業はイノベーションを起こしません。中の人がいるからではありません
彼らにはアイデアが欠けていますが、会社はリスクを軽減し保護するために、多くの場合意図的に最適化されているためです。
既存の収入。イノベーションは本質的にリスクと破壊的であり、安定性を求めて調整されたシステムはそれを拒否します
免疫系が外来細胞を拒絶する仕組み。
成熟したプロジェクトが機能する

同じように。実際のユーザー、下流のエコシステム、そしてコア グループが存在すると、
長い間存在しているため、支配的な力は「何も壊さない」になります。査読者は慎重になります。範囲
引き締めます。野心的なパッチは議論の中で消えていきます。プロジェクトは静かに 1 つの位置に収束します: それを生かし続ける、維持する
安定していますので、最下層に触れないでください。大きなイノベーションは、関係なく誰でも押し進められるものではなくなる
彼らがそうしたかったかどうかについて。
企業の世界には、少なくとも逃し弁があります。大企業はスタートアップにリスクを負わせ、それまで待ってください
アイデアが実証され、結果が得られます。イノベーションは外部で起こり、既存企業がそれを購入する
既知の量。オープンソースには同等のものはありません。大規模で成熟したプロジェクトが小規模で革新的なプロジェクトを吸収することはほとんどありません。あ
代替実装を成功させるには、それ自体が独立している必要があります。構築されたものにはマージされません
交換する。成熟した OSS における有意義なイノベーションは、内部プロセスではありません。それは外部のものです。
パターンは探せばどこにでもあります。 Cassandra は MySQL コードベース内に構築されていません。
CockroachDB は、Postgres ワイヤ プロトコルを使用するにもかかわらず、Postgres 内に構築されていません。
Postgres と互換性のある SQL であり、単一ノード Postgres を超えたユーザーを対象としています。 Postgresならそうなるだろう
Postgres が内部的にその方向に進化できたとしたら、代わりにそれを構築する必要がありました。
外で。最新のデータベース環境 (Cassandra、Redis、ClickHouse、DuckDB、Tigerbeetle、および数十もの)
他の人）は、人々が何か違うものを求めていて、既存の企業がそれを吸収しようとしなかったときに起こったことです
それ。そのどれもが MySQL や Postgres に折り返されることはありませんでした。現職は誰が左に行っても走り続ける。
ソフトウェア開発には、既存の機能を実現するまでの緊張関係が常にありました。

g ツールと独自のバージョンをローリングします。
どちらの極端な場合にも、有名な故障モードがあります。一方では、些細なことのために何百もの依存関係を取り込みますが、
そして6か月後、そのうちの1つがハイジャックされているか非公開になっていることがわかりました。一方、使用を拒否する場合
あらゆるものに対応する既製のソリューションを使用し、世界中で使用されているパーサーとシリアル化ライブラリを書き直すのに何か月も費やします。
10年間持っています。
少なくとも最近まで、2 番目の極端を抑制してきたのは速度です。の独自バージョンをローリングする
確立されたツールは高価です。問題領域は大きく、エッジケースは敵対的で、プラットフォームのロングテール
癖は永遠に続く。ツールを再構築するために知っておくべきことをすべて知っていたとしても、数か月、あるいは数年かかります。
タイピングが邪魔になった。ほとんどの場合、自分でツールを構築するよりも、手持ちのツールを使い続けるほうが安上がりです。
欲しかった。ゼロから構築するのは、絶望的な人や資金に余裕のある人だけのものでした。
✻ クロード・コードへようこそ
cwd: ~/linux
╭───────────────────────────╮
│ > 死の痛みで Linux カーネルを Rust で書き換えます。 │
│ 10 分後に戻ってきますので、起動することを期待します。使うだけ│
│ 本当に必要な場合、または単に感じているだけの場合は危険です │
│冒険好き。 │
╰─────────────────────────╯
【10分後…】
╭───────────────────────────╮
│ > ごめんなさい、サビコンピを忘れてました

時間があります。また来ます│
│2日。 │
╰─────────────────────────╯
確立されたプロジェクトを（異なる言語で、異なる優先順位で）書き直すことへの障壁
特定の目的）は本質的になくなりました。
最も興味深いのは、AI がコードを書けることではありません。それは既存のプロジェクトが使えるようになるということです
参考。何十年にもわたって蓄積された意思決定、エッジケース、プラットフォームの癖が、かつてはあらゆることを妨げる壁となっていました。
書き直す。現在では、それらは自分のバージョンと照らし合わせて読み取ることができる仕様になっています。現職者は刑務所になるのをやめる
そしてドキュメント化が始まります。
代わりに構築するものは、考えられるあらゆる用途ではなく、特定の用途に合わせて最適化できます。既製
ツールはさまざまな問題をかなりうまく解決します。専用のツールが無料で問題をうまく解決します
誰にとってもすべてであるために、何年にもわたって蓄積された残骸のこと。たぶんあなたは別の形でそれを望んでいます
言語。もしかしたら、元のアーキテクチャでは禁止されている方法で、それをスタックの残りの部分と統合したいかもしれません。
もしかしたら、別のメモリ モデル、ディスク上のフォーマット、または同時実行性のストーリーが必要になるかもしれません。これらすべてが今や手の届くところにあります。
AIがやらないのはアーキテクチャや判断です。インフラ構築の難しい部分はまだ続く
人間から見たもの: 何を公開するか、何を隠すか、どこで抽象化が崩れるか、どのトレードオフが現実で、どのトレードオフが現実であるか
民間伝承。しかし、何か月にもわたるタイピング、プラットフォームの癖のロングテール、徹底的なテストマトリックス、
実際のアイデアを中心とした定型文: それが安っぽくなったものです。それを削除しても書き換えが簡単になるわけではありません。それ
書き換えが有限になります。
ライブ マイグレーションは私たちの核心です

建てる。アーキテクトの価値提案全体 (実行中のワークロードを複数の組織間で移動する)
クラウド、ノード間、ハードウェア間）は、チェックポイント/復元が高速で予測可能であり、私たちが修正できるかどうかに依存します。
壊れたとき。プロセスのすべてのステップを最適化し、ボトルネックを特定して排除する必要があります (ほぼ常に
ディスクまたはネットワークは当社の規模で）、バグが見つかったその日に修正版を出荷します。
CRIU は私たちをそこへほとんど連れて行ってくれます。明確にしておきますが、これは 10 年以上の実績を持つ本格的なソフトウェアです。
慎重な作業、深いカーネル統合、そして私たちが大いに尊敬する真のエンジニアリング努力。しかし、「ほとんどの
チェックポイント/復元が製品の基礎である場合、「その方法」だけでは十分ではありません。そして、その下流で働いています
自分のニーズとは関係のない未解決の問題が何百もある大規模なプロジェクト (関連するため無視できない問題)
製品に依存する同じコードパス

[切り捨てられた]

## Original Extract

Rewriting an established open source project used to take years. With AI, that's no longer true. Here's why we're rewriting CRIU in Zig for our own use, what makes infrastructure projects go stale in the first place, and why the rebuild-vs-contribute math might be worth a second look.

The Math Changed: Rewriting Stale Open Source In The AI Era { ∞ } Loophole Labs Architect
~/ BLOG / engineering / The Math Changed
The Math Changed: Rewriting Stale Open Source In The AI Era
By jimmy moore , Founding Engineer On this page:
How Open Source Projects Go Stale
Until recently, rewriting an established open source project could take years. With LLMs, that's changed. We're
rewriting CRIU in Zig, and expect it to be complete in months, not years.
Anyone who's done home renovations knows there's a tool for every job. Over the years you end up with a shed full
of specialized things that each do one job well. Tools you might only need a handful of times, but when the
moment comes, nothing else will do.
My house Tools to fix up my house
════════ ════════════════════════
┌───────────────────────────┐
│ hammer saw drill │
│ chisel awl pliers │
│ level square axe │
╱╲ │ jigsaw sander oil │
╱ ╲ │ crowbar plane vise │
┌────┐ │ tape wrench ratchet │
│ ▢▢ │ │ spanner nailer awl │
└────┘ └───────────────────────────┘
Software is the same story. As we build, we reach for the specialized tools that let us move faster and ship better.
And both kinds of tools need maintenance. A rusty saw is more pain than it's worth when you actually need to cut
wood, and a stale software tool can cost you more time than it ever saved. The catch with software is that you
usually aren't the one keeping it sharp. Someone else is, and you're trusting that they still are.
The standard advice, when you notice your tool getting dull, is to be a good open-source citizen: file the bug,
write the patch, send the pull request. Most of us have done it, and the reciprocity is what's kept the ecosystem
running. But that model is starting to show strain. Maintainers are quietly drowning in AI-generated patches and
bug reports that look right at a glance and
burn
hours under
review .
The contribute-upstream economics worked when humans wrote thoughtful patches for humans to review, at roughly
comparable speed. When one side of that equation runs at machine speed and the other doesn't, the whole loop
breaks down. It's worth asking whether the model still works at all.
Checkpoint/restore is what Architect does for a living. We take our customers' running Linux processes - their
memory, their threads, their open files and sockets - freeze them, pick them up, and bring them back to life on a
different machine. There is exactly one tool on Linux that does this: CRIU . It's been
around for more than a decade, it's wired into runc, podman, and the wider Kubernetes ecosystem, and it's the
obvious thing every team reaches for first.
How Open Source Projects Go Stale
The issues we kept running into - and the workarounds we kept piling on top of each other - weren't really
CRIU problems. They were structural ones, the kind every long-lived open source project eventually runs into,
and no amount of patching around the edges was going to fix them.
People have passions - almost obsessions. We've all had things in our lives we were utterly obsessed with for a
while, and then just grew away from. Maintaining that intensity over years, let alone decades, is genuinely hard.
At the start of a project, everything is new. Every commit moves the needle, the mental model fits in your head,
and progress is fast and fun. As the project matures, the work shifts. More of your time goes to maintenance,
regression hunts, and keeping the existing code from rotting under you. Back to the renovation analogy: we've all
got the project on the shelf that died at 90%, because the last 10% was the unfun part.
When code is written by hand, people only have finite energy to push things forward. The original driving forces
eventually burn out and move on, and the project keeps running on whoever's left.
A long-lived project also exists in a world that keeps changing around it. When CRIU started in 2012, Go had just
hit 1.0. Rust was three years from its first stable release. Zig wouldn't exist for another four. The Linux kernel
didn't have io_uring. eBPF as we know it didn't exist. Containers were a niche curiosity rather than the substrate
of half the world's compute. Build systems, CI, package management, testing conventions, documentation tooling -
all of it has moved, often more than once.
When a project doesn't keep refactoring toward its environment, the world around it slowly drifts out of fit. The
tells are subtle. CRIU's homepage is still a MediaWiki install in 2026. It's the kind of detail that doesn't
matter on its own, but reads like wallpaper from the wrong decade. It quietly marks when the project stopped
renovating.
┌──────────────────────────────────────────────────────────────────────────┐
│ File Edit View Help │
├──────────────────────────────────────────────────────────────────────────┤
│ ◄ ► ⟳ ⌂ http://geocities.com/~user/website-i-have-not-updated-lately │
├──────────────────────────────────────────────────────────────────────────┤
│ │
│ ★ WELCOME TO MY HOMEPAGE ★ visitors: [ 0 0 0 0 4 2 7 ] │
│ ⚠ UNDER CONSTRUCTION ⚠ [<animated worker.gif>] │
│ │
│ webring: << prev | random | next >> last updated: April 1998 │
│ sign my guestbook! best viewed in Netscape │
│ │
└──────────────────────────────────────────────────────────────────────────┘
Some of that drift can be patched over. CRIU did get bolted into runc, podman, and Kubernetes after they emerged.
If the existing core exposes a reasonable interface, new orchestration layers can sit on top of it without
surgery. Other kinds of drift can't be patched at all. You don't migrate an established C codebase to Rust or Zig
with a few pull requests. The implementation language, the memory model, the on-disk formats, and the core
architecture all freeze the moment they're chosen. By the time the world has moved on, the cost of moving with it
approaches the cost of starting over.
There's a third force at work, and it's structural. Big companies don't innovate. Not because the people inside
them lack ideas, but because the company has been optimized, often deliberately, to reduce risk and protect
existing revenue. Innovation is by definition risky and disruptive, and a system tuned for stability rejects it
the way an immune system rejects a foreign cell.
Mature projects work the same way. Once there are real users, a downstream ecosystem, and a core group that's
been around a long time, the dominant force becomes "don't break anything." Reviewers get cautious. Scope
tightens. Ambitious patches die in discussion. The project quietly converges on a position: keep it alive, keep
it stable, don't touch the bottom layer. Big innovation stops being something anyone can push through, regardless
of whether they wanted to.
In the corporate world there's at least an escape valve. Big companies let startups take on the risk, wait until
the idea proves out, and acquire the result. The innovation happens outside, and the incumbent buys it once it's
a known quantity. Open source has no equivalent. Large mature projects rarely absorb small innovative ones. A
successful alternative implementation has to stand on its own; it doesn't get merged into the thing it was built
to replace. Meaningful innovation in mature OSS isn't an internal process. It's an external one.
The pattern is everywhere once you look for it. Cassandra didn't get built inside the MySQL codebase.
CockroachDB didn't get built inside Postgres, even though it speaks the Postgres wire protocol, uses
Postgres-compatible SQL, and targets users who outgrew single-node Postgres. It's the thing Postgres would
have become if Postgres could have evolved in that direction internally, and instead it had to be built
outside. The modern database landscape (Cassandra, Redis, ClickHouse, DuckDB, Tigerbeetle, and dozens of
others) is what happened when people wanted something different and the incumbents weren't going to absorb
it. None of it got folded back into MySQL or Postgres. The incumbents keep running on whoever's left.
Software development has always had a tension between reaching for an existing tool and rolling your own version.
Both extremes have their famous failure modes. On one end you pull in hundreds of dependencies for trivial things,
and find out six months later that one of them has been hijacked or unpublished. On the other you refuse to use
ready-made solutions for anything, and spend months rewriting parsers and serialization libraries that the world
has had for a decade.
What's kept the second extreme in check, at least until recently, is speed. Rolling your own version of an
established tool is expensive. The problem space is large, the edge cases are hostile, the long tail of platform
quirks goes on forever. Even if you knew everything you needed to know to rebuild a tool, the months or years of
typing got in the way. Most of the time it was cheaper to live with the tool you had than to build the one you
wanted. Building from scratch was reserved for the desperate or the well-funded.
✻ Welcome to Claude Code
cwd: ~/linux
╭─────────────────────────────────────────────────────────────╮
│ > Rewrite the linux kernel in rust on pain of death. │
│ I'll be back in 10 mins and expect it to boot. Only use │
│ unsafe if you really have to, or if you're just feeling │
│ adventurous. │
╰─────────────────────────────────────────────────────────────╯
[10 mins later...]
╭─────────────────────────────────────────────────────────────╮
│ > Sorry I forgot about rust compile time. I'll be back in │
│ 2 days. │
╰─────────────────────────────────────────────────────────────╯
The barriers to rewriting an established project (in a different language, with different priorities, for a more
specific purpose) are essentially gone.
The most interesting thing isn't that AI can write code; it's that the existing project becomes a usable
reference. Decades of accumulated decisions, edge cases, and platform quirks used to be the wall blocking any
rewrite. Now they're a specification you can read against your own version. The incumbent stops being a prison
and starts being documentation.
What you build instead can be optimized for your particular usage rather than every possible usage. Off-the-shelf
tools solve a wide range of problems fairly well. A purpose-built tool can solve your problem really well, free
of the cruft that accumulates over years of being everything to everyone. Maybe you want it in a different
language. Maybe you want to integrate it with the rest of your stack in ways the original architecture forbids.
Maybe you want a different memory model, on-disk format, or concurrency story. All of these are now within reach.
What AI doesn't do is the architecture or the judgment. The hard parts of building infrastructure still come
from humans: what to expose, what to hide, where the abstractions break, which tradeoffs are real and which are
folklore. But the months of typing, the long tail of platform quirks, the exhaustive test matrices, the
boilerplate around the actual ideas: that's what's gotten cheap. Removing it doesn't make rewriting trivial. It
makes rewriting finite.
Live migration is core to what we build. Architect's whole value proposition (moving running workloads across
clouds, across nodes, across hardware) depends on checkpoint/restore being fast, predictable, and ours to fix
when it breaks. We need to optimize every step of the process, identify and eliminate bottlenecks (almost always
disk or network at our scale), and ship a fix the same day we find the bug.
CRIU gets us most of the way there. To be clear: it's a serious piece of software, with more than a decade of
careful work, deep kernel integration, and a real engineering effort we have a lot of respect for. But "most of
the way" isn't enough when checkpoint/restore is the foundation of your product. And working downstream of a
large project with hundreds of open issues unrelated to your needs (issues you can't ignore, because they touch
the same code paths your product depe

[truncated]
