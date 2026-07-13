---
source: "https://verial.xyz/systems/standards-that-are-code"
hn_url: "https://news.ycombinator.com/item?id=48898659"
title: "Standards that are code [MIT]: deterministic checks instead of LLM inference"
article_title: "Standards That Are Code | Verial"
author: "verial-lab"
captured_at: "2026-07-13T21:45:59Z"
capture_tool: "hn-digest"
hn_id: 48898659
score: 1
comments: 0
posted_at: "2026-07-13T20:50:38Z"
tags:
  - hacker-news
  - translated
---

# Standards that are code [MIT]: deterministic checks instead of LLM inference

- HN: [48898659](https://news.ycombinator.com/item?id=48898659)
- Source: [verial.xyz](https://verial.xyz/systems/standards-that-are-code)
- Score: 1
- Comments: 0
- Posted: 2026-07-13T20:50:38Z

## Translation

タイトル: コードである標準 [MIT]: LLM 推論の代わりに決定論的チェック
記事のタイトル: コードである標準 |ベリアル
説明: Agent Paradise Standards System (APSS): ベスト プラクティスを実行可能な自己検証標準としてエンコードするため、エージェントは確率的推論ではなく決定論的なレールで反復作業を実行します。

記事本文:
コードである標準 | Verial Verial エッセイ投稿システム ⌘K ← コードであるシステム標準
Agent Paradise Standards System (APSS): ベスト プラクティスを実行可能な自己検証標準としてエンコードすることで、エージェントは確率的推論ではなく決定論的なレールで反復作業を実行します。
滑り台とゴルトンボード
その成果はスピード、ワット数、自信です
目次 コーディング エージェントと仕事をした経験から、何度も説明し直さなければならないパターンがたくさんあります。たとえイライラしたり、重要な詳細が漏れてしまう可能性があるとしても、コンテキストがなければすべてのエージェント セッションは新たに開始されるため、これは当然のことです。これに関して私が聞いた経験則は、「3 回目のルール」です。エージェントに同じことを 3 回説明した場合、それを再利用可能なスキルに変換します。
これは良いルールであり、定期的に利用する再利用可能なワークフローを捉えます。しかし、エンジニアリングの観点から見ると、再利用性は半分にすぎません。拡張したい場合は、パフォーマンスも重要です。私たちは、お金を払っているハードウェアを最大限に活用したいと考えています。
多くのエージェントが同じパターンを並行して実行できるため、再利用可能なスキルが役に立ちます。しかし、それでも推論に基づいて実行され、推論にはコストがかかります。
では、反復可能なワークフローを決定論的なスクリプトに変えることができたらどうなるでしょうか?
滑り台とゴルトンボード #
今の雰囲気は「AIならできる、ただプロンプトすればいい」というものです。真に新しい作品にとって、それは正しい本能です。推論は、これまで人間が座って考える必要があったソフトウェアの部分に使用できる初めてのツールです。
しかし、推論には形があります。 Galton ボード (Plinko として知られるペグボード) を想像してください。コインを上部に落とすと、ピンの間をカチャカチャと音を立てて通過し、鐘のカーブのどこかに止まります。同じ入力、異なる出力。それは推論であり、設計上確率的です。そして、e

ベリーバウンスはちょっとした作業です。コインの跳ね返りが多ければ多いほど、決済までに時間がかかります。実際のマシンでは、コンピューティングにマッピングされ、コンピューティングには時間とワットがかかります。
スライドをイメージしてみましょう。毎回、上部から入って下部から出ます。1 つのパスです。それが決定論的なコードです。 Rust バイナリ、スクリプト。コンピューティングにおける最も古い保証であり、一度推論に夢中になると、まだ決定論が存在することを半分忘れていました。パスが 1 つあると作業量が減り、時間もワット数も減ります。
私たちがエージェントに一日中やってもらうことのほとんどは新しいものではありません。このファイルを検証します。その足場を生成します。設定を契約と照らし合わせて確認してください。私たちはこれらのコインをペグボードを通じて送信し続け、ワット、トークン、そして通常は決定的な質問に対する確率的な答えに対する信頼を支払います。
Agent Paradise Standards System は、決定論的な道に対する私たちの賭けです。略してAPSS。一行バージョン: コードである標準。
APSS の標準は、形状のあるアーティファクトを中心に構築されています。その形状をスキーマ、protobuf コントラクトとして固定すると、そのスキーマが唯一の信頼できる情報源になります。標準自体は Rust クレートであり、決定性と速度を重視して選択されています。
その 1 つの定義には、アーティファクトを生成するジェネレーターと、特定のアーティファクトが準拠していることを証明するバリデーターという 2 種類のコードがぶら下がっています。そして、すべてがアーティファクトにぶら下がっているため、サブスタンダードは、それを作成した標準に触れることなく、それを読み取って何か新しいもの (ビューまたはレポート) を生成する独自のジェネレーターを追加できます。
1 つの定義からの生成と検証は相互に強化します。標準は、同じ真実の情報源から、正しい形状で物事を生成し、物事がその中にあることを証明します。だからこそ私はそれを「賢い」と呼んでいます。標準はルールの説明ではなく、双方向で実行可能なルールです。
標準があれば

コードである場合、標準を記述するためのルールもコードである必要があります。したがって、APSS にはメタ標準があります。これは、他のすべての標準がどのように構成されているかを定義し、それ自体を検証する標準です。人間だけが強制する慣習は、継続的に注意を払わないと劣化する傾向があります。不適合はビルドの失敗であり、忘れられたガイドラインではないため、ツールが強制する規則は適用できません。
標準はソフトウェアと同様にバージョン管理されます。 V1 はすべて下位互換性を維持しており、その契約を破る必要がある場合は、エコシステムに対して V2 を使用する義務があります。標準には、サブ標準を含めることもできます。親の成果物に基づいて構築され、独自のタイムラインでバージョン管理された小さな標準です。
CodeCity は、それがどのように階層化されるかを示す良い例です。コード トポロジ標準は、コードベースをトポロジ アーティファクト、つまり言語に依存しない構造、複雑さ、結合のスナップショットに分析します。次に、標準以下のビジュアライゼーションがそのアーティファクトを消費してレンダリングします。以前に書いた 3D スカイライン ビューである CodeCity は、そのレンダリングの 1 つです。したがって、親はアーティファクトを生成し、サブスタンダードはそこからビューを生成します。
各規格には、単なるルールではなく、その根拠、理由も含まれています。これにより、標準が同意または異議を唱えることができる議論に変わり、システム全体が実験室に変わります。アイデアは実験として始まり、それが得られて初めて正式な標準に昇格します。
その成果はスピード、ワット数、自信です #
それを組み合わせると、エージェントの手続き記憶のようなものが得られます。通常であればセッションごとに再説明するプラクティスが実行可能になり、コードベース間で再利用可能になり、CI に直接接続されます。人間が標準を定義し、エージェントがそれを実行します。そして、緑色のチェックは、アーティファクトが決定論的に標準に準拠していることを意味しますが、その推論は決して約束するものではありません。ヴァ

Lidator 自体は依然として正しい必要がありますが、それは実行ごとではなく、1 回支払うコストです。
次にコストがかかり、ここで哲学ではなくなります。すべてのコンピューティングは、根本的には電力の問題です。エージェントが常に行うように求められることを考えてみましょう。それは、コードベースを理解することです。
APSS をリポジトリに指定すると、トポロジ全体、すべてのモジュール、その複雑さ、すべてがどのように結合するかを約 6 秒、およそ 100 ジュールで計算します。 1 モデルから同じマップを取得するには、コードをフィードする必要があり、このリポジトリだけで約 700,000 トークンになります。単一のロングコンテキスト読み取りでも、その約 1,000 倍にあたる 10 万ジュールのオーダーで実行され、モデルのマップは、アナライザーのマップが実行ごとに正確かつ同一である推測になります。
したがって、キロワット時で数万の決定論的なコードベース マップ、またはモデルから数十のマップが購入されます。これは保守的なケースです。実際の分析は 1 回の読み取りではなく、それらのエージェントによるループです。モデルの効率が上がるにつれて差は縮まりますが、ここでは明白な理由でコンピューティングが推論に勝ります。アナライザーは固定の少量のコンピューティングを実行し、モデルは大量のコンピューティングを実行します。
きちんとした数字よりも誠実さ。 2 つの側面は異なる方法で測定され、アナライザーのラップトップ CPU の数値とデータセンターのモデルの推定値が比較され、どちらもタスクに応じて変化します。これは正確な比率ではなく、桁として解釈してください。実際のエージェント分析では差が広がるだけであることに注意してください。
これは推論に反する議論ではありません。それをどこに使うかについての議論です。
システム全体が圧縮するルールは次のとおりです。タスクが反復的で、チェック可能な形状を持つ場合、それを推論せず、エンコードします。エンコードには、スキーマとその周りのコードなど、事前にコストがかかるため、繰り返しで利益が得られます。チェックが頻繁に実行されるほど、より多くの効果が得られます。

固定費は償却されません。
試してみてください。エージェントに CodeCity APSS Runbook を指定し、独自のコードベースを CodeCity としてレンダリングさせます。
これを採用してください。Cargo install apss を使用して CLI をインストールします。これはオープン ソースの MIT であり、アーティファクト シェイプは構築できるオープン コントラクトです。
貢献: 貢献ガイドはシンプルでエージェントフレンドリーです。新しい標準は実験として開始され、最初のコミットから足場が構築され、自己検証されます。
2026 年にローカルで測定。APSS コードベース (121 ファイル、2,145 関数、Rust リリース ビルド、Apple M1) での apss 実行コード トポロジ分析には約 6.4 秒かかり、推定 10 ～ 20 ワットで約 100 ジュールでした。壁掛け時計は計測されます。エネルギーは推定されるため、実験室の数値ではなく、桁の大きさとして扱います。 ↩
APSS コードベースは約 700,000 トークン (ソース文字を 4 で割ったもの) です。 1 回のパスで多くのトークンが実行されることを読み取ると、測定されたトークンごとの推論エネルギー (ML.ENERGY リーダーボード、トークンごとにおよそ 0.1 ～ 0.5 ジュール) と Google が開示したプロンプトごとのエネルギーから推定され、100,000 ジュールのオーダーで実行されます。現実的なエージェント分析では、このような呼び出しが多数発生し、コストが高くなります。これは意図的に控えめな推定値です。 ↩
システムに戻る Verial Verial — veritas (真実) と aerial (無限の空) から。私たちが真実を追求し、意図を持って構築する場合、宇宙マイクロ波背景放射が限界であるという信念。
© 2026 ベリアル。無断転載を禁じます。

## Original Extract

The Agent Paradise Standards System (APSS): encoding best practices as executable, self-validating standards, so agents run repetitive work on deterministic rails instead of stochastic inference.

Standards That Are Code | Verial Verial Essays Posts Systems ⌘K ← Systems Standards That Are Code
The Agent Paradise Standards System (APSS): encoding best practices as executable, self-validating standards, so agents run repetitive work on deterministic rails instead of stochastic inference.
The slide and the Galton board
The payoff is speed, watts, and confidence
Contents In my experience working with coding agents, there are many patterns I have to re-explain over and over. Even though it can be frustrating and important details may slip through, it's understandable because every agent session starts fresh unless it has context. A rule of thumb I've heard for this is the Third Time Rule : if you have explained the same thing to an agent three times, turn it into a reusable skill.
It is a good rule, and it captures the reusable workflows you lean on regularly. But from an engineering perspective, reusability is only half the story. If we want to scale, performance matters too. We want the most out of the hardware we are paying for.
A reusable skill helps, since many agents can run the same pattern in parallel. But it still runs on inference, and inference is expensive.
So what if we could turn a repeatable workflow into a deterministic script?
The slide and the Galton board #
The mood right now is "AI can do it, just prompt it." For genuinely new work, that is the right instinct. Inference is the first tool we have ever had for the parts of software that used to require a human to sit and think.
But inference has a shape. Picture a Galton board , the pegboard you may know as Plinko. Drop a coin at the top, it clatters through the pins, and it lands somewhere in a bell curve. Same input, different output. That is inference, stochastic by design. And every bounce is a little piece of work: the more the coin ricochets, the longer it takes to settle. In a real machine that maps to compute, and compute is time and watts.
Now picture a slide. In at the top, out at the bottom, one path, every time. That is deterministic code. A Rust binary, a script. The oldest guarantee in computing, and once we got excited about inference we half forgot we still had determinism. One path means less work: less time, fewer watts.
Most of what we ask agents to do all day is not new. Validate this file. Generate that scaffold. Check the config against the contract. We keep sending those coins through the pegboard, paying in watts, tokens and trust for a probabilistic answer to a question that usually has a deterministic one.
The Agent Paradise Standards System is our bet on the deterministic path. APSS, for short. The one-line version: standards that are code.
A standard in APSS is built around an artifact with a shape. You pin that shape down as a schema, a protobuf contract, and that schema becomes the single source of truth. The standard itself is a Rust crate, chosen for determinism and speed.
Two kinds of code hang off that one definition: a generator that produces the artifact, and a validator that proves a given artifact conforms. And because everything hangs off the artifact, a substandard can add its own generator that reads it and produces something new, a view or a report, without touching the standard that created it.
Generation and validation from one definition reinforce each other: the standard produces things in the correct shape and proves things are in it, from the same source of truth. That is why I call it smart. The standard is not a description of the rule, it is the rule, executable in both directions.
If standards are code, then the rules for writing a standard should be code too. So APSS has a meta-standard : the standard that defines how every other standard is structured, and validates them itself. A convention only humans enforce tends towards degradation if continuous attention is not paid. A convention the tooling enforces cannot, because nonconformance is a failing build, not a forgotten guideline.
Standards are versioned like software. All of V1 stays backwards compatible , and if you need to break that contract you owe the ecosystem a V2. A standard can also have substandards: smaller standards built on the parent's artifact, versioned on their own timeline.
CodeCity is a good example of how that layers. The Code Topology standard analyzes your codebase into a topology artifact, a language-agnostic snapshot of structure, complexity, and coupling. A visualization substandard then consumes that artifact and renders it. CodeCity , the 3D-skyline view I wrote about earlier, is one of those renderings. So the parent generates the artifact, and the substandard generates a view from it.
Each standard also carries its rationale, the why and not just the rule. That turns a standard into an argument you can agree with or challenge, and it turns the whole system into a lab: an idea starts as an experiment and gets promoted to an official standard only once it has earned it.
The payoff is speed, watts, and confidence #
Put it together and you get something like procedural memory for agents. The practices we would otherwise re-explain every session become executable, reusable across codebases, and wired straight into CI. Humans define the standard, agents execute it. And a green check means the artifact conforms to the standard, deterministically, which inference never promises. The validator itself still has to be correct, but that is a cost you pay once, not on every run.
Then there is the cost, and this is where it stops being philosophy. All computing is, at bottom, a power problem. Take something agents get asked to do all the time: understand a codebase.
Point APSS at a repository and it computes the whole topology, every module, its complexity, how it all couples, in about six seconds and roughly a hundred joules. 1 To get the same map from a model, you have to feed it the code, and this repository alone is around 700,000 tokens. Even a single long-context read runs on the order of a hundred thousand joules, 2 about a thousand times more, and the model's map is a guess where the analyzer's is exact and identical every run.
So a kilowatt-hour buys you tens of thousands of deterministic codebase maps, or a few dozen from a model. And that is the conservative case: a real analysis is not one read but an agentic loop of them. As models get more efficient the gap narrows, but computing beats inferring here for a plain reason, the analyzer runs a fixed, small amount of compute, and a model runs a great deal.
Honesty over a tidy number. The two sides are measured differently, a laptop CPU figure for the analyzer against a datacenter estimate for the model, and both move with the task. Read it as an order of magnitude, not a precise ratio, and note that a real agentic analysis only widens the gap.
None of this is an argument against inference. It is an argument about where to spend it.
The rule the whole system compresses down to: if a task is repetitive and has a checkable shape, do not infer it, encode it. Encoding costs something up front, the schema and the code around it, so the payoff is in repetition: the more often the check runs, the more that fixed cost amortizes toward nothing.
Try it: point an agent at the CodeCity APSS runbook and have it render your own codebase as a CodeCity.
Adopt it: install the CLI with cargo install apss . It is open source , MIT, and the artifact shapes are open contracts you can build against.
Contribute: the contribution guide is simple and agent-friendly. New standards start as experiments, scaffolded and self-validating from the first commit.
Measured locally, 2026. apss run code-topology analyze over the APSS codebase (121 files, 2,145 functions; Rust release build, Apple M1) took about 6.4 seconds, roughly 100 joules at an estimated 10 to 20 watts. Wall-clock is measured; energy is estimated, so treat it as an order of magnitude, not a lab figure. ↩
The APSS codebase is about 700,000 tokens (its source characters divided by four). Reading that many tokens in a single pass runs on the order of 100,000 joules, extrapolated from measured per-token inference energy (the ML.ENERGY leaderboard , roughly 0.1 to 0.5 joules per token) and Google's disclosed per-prompt energy . A realistic agentic analysis makes many such calls and costs more. This is an estimate, deliberately conservative. ↩
Back to Systems Verial Verial — from veritas (truth) and aerial (limitless sky). The belief that when we seek truth and build with intention, the cosmic microwave background is the limit.
© 2026 Verial. All rights reserved.
