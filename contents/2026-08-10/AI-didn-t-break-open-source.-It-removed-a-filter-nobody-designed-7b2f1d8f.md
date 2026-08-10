---
source: "https://erickxdev.vercel.app/en/blog/ai-didnt-break-open-source/"
hn_url: "https://news.ycombinator.com/item?id=49247364"
title: "AI didn't break open source. It removed a filter nobody designed"
article_title: "AI didn't break open source. It removed a filter nobody designed. — Erick Xavier"
author: "XaviDev"
captured_at: "2026-08-10T18:42:48Z"
capture_tool: "hn-digest"
hn_id: 49247364
score: 1
comments: 0
posted_at: "2026-08-10T18:04:24Z"
tags:
  - hacker-news
  - translated
---

# AI didn't break open source. It removed a filter nobody designed

- HN: [49247364](https://news.ycombinator.com/item?id=49247364)
- Source: [erickxdev.vercel.app](https://erickxdev.vercel.app/en/blog/ai-didnt-break-open-source/)
- Score: 1
- Comments: 0
- Posted: 2026-08-10T18:04:24Z

## Translation

タイトル: AI はオープンソースを破壊しませんでした。誰も設計していないフィルターを取り除いた
記事のタイトル: AI はオープンソースを破壊しませんでした。誰も設計していないフィルターが取り除かれました。 — エリック・ザビエル
説明: オープンソースは、誰も意図的に構築したわけではない品質フィルター、つまり貢献する前に理解する努力の上で数十年にわたって実行されてきました。それ

記事本文:
← ホーム
ブログ
☀️
2026-08-06 · AI、オープンソース、ソフトウェア エンジニアリング、生産性
AI がオープンソースを破壊したわけではありません。誰も設計していないフィルターが取り除かれました。
投稿者: erickdevz · フルスタック開発者、Viçosa、ブラジル
2026 年 1 月、curl プロジェクトは実施していたバグ報奨金制度を終了しました。
2019年以来。資金がなくなったからではなく、それより多くの支払いがあったからです。
確認された数十件の脆弱性に対して 10 万米ドル。シャットダウンしました
トリアージの経済性が崩壊したからである。同じ窓の周りで、
84 の Python プロジェクトを管理していた Jazzband 集団は、次のように発表しました。
それは終わりつつあった。 tldraw の創設者はリポジトリを次のように切り替えました。
すべての外部プルリクエストを自動クローズします。 Ghostty はゼロトレランスを採用しました
低労力の AI 提出に関するポリシー。
これらのプロジェクトはどれも、希望する人がいないためにトラブルに見舞われることはありませんでした。
貢献する。彼らは反対側の下に座屈した。
正しい診断は次のことを示しているため、その理由を理解することは価値があります。
業界のほとんどが構築している修正とは大きく異なります。
誰も知らなかったフィルターがそこにあった
ミッチェル・ハシモト氏、Ghostty の開発者 (以前は HashiCorp を共同設立)
it) は、プロジェクトの更新された AI でこれについて最も鋭い一文を書きました
ポリシー。彼のフレーミング:
エージェント プログラミングにより、自然な労力ベースのバックプレッシャーが排除されました
以前は労力の少ない貢献を制限していましたが、今ではあまりにも簡単になりました
最小限の労力で大量の悪質なコンテンツを作成します。
バックプレッシャー という言葉は無視してください。それが議論の全体だからです。
2023 年までに、
あなたがまだ知らなかったプロジェクト。リポジトリのクローンを作成します。構築してもらいます。読む
変更がどこに当てはまるかを見つけるのに十分なコードが含まれています。テストを実行します。ワークアウト
なぜテストが失敗したのか。
誰もそれを品質システムとして設計したわけではありません。しかし、それは一つとして機能した。の
努力は代償を伴うものだった

フロアを保証: 獲得したほぼ全員
PR が途中で何かを理解した限りでは。
通行料はなくなりました。 CI をコンパイルして渡すパッチの作成にはコストがかかります
数分経っても全く理解できませんでした。存在した上質なフロア
努力の副作用として、努力とともに消えていきます。
橋本氏はこれが反AIの立場ではないと強調している — Ghosttyは
AI の強力な支援を利用して書かれており、そのメンテナは毎日 AI を使用しています。の
line はツールに関するものではありません。外部委託する貢献者についてです。
モデルに考えさせ、確認請求書をモデルに渡します。
メンテナー。
システムを壊す非対称性
コードの生成が劇的に安くなりました。コードのレビューはこれ以上安くなりません
全然。
この非対称性は、前面ドアが開いているシステムにとって致命的です。の
カール番号
make it concrete: the share of vulnerability reports that turned out to
be real は何年にもわたって 15% を超えて推移していましたが、2025 年には 5% を下回りました。
AI が生成したレポートが殺到しました。2026 年の最初の 3 週間に、カール
received twenty submissions — seven of them inside a single sixteen-hour
ウィンドウ — そしてセキュリティ チームがそれぞれを読んだ後、再現を試みました
それを調べてコード パスを追跡しましたが、実際の脆弱性を説明するものはありませんでした。
ダニエル・ステンバーグ自身の投稿の要約
「千回のスロップによる死」 、
率直に言いました: 音量が上がるだけでなく、品質も下がります。
これまで以上に時間を費やして、これまで以上に成果を出さないようにしてください。
セキュリティ レポートはこれの最も残酷なバージョンです。
それらを無視してください。実際の脆弱性を無視すると、壊滅的なコストがかかります。それで
でっち上げられた報告書はすべて、公表する前に真剣に調査する必要がある
解雇された。生成コストはゼロに近づきます。の費用
責任を持って廃棄する割合は依然として高い。
正直な脚注:カール
プログラムを再開しました
1ヶ月くらい

その後、レポートの品質が回復しました。特に、何
戻ってきたのは、大したことはなく、大量、高品質でした。
人間が検証したレポート。これが重要な点です。問題は決して起こらなかったのです
あい。それは労力を省いた未レビューの出力でした。
AI の検出が間違った答えである理由
市場の本能は、検出器、つまり PR をスコアリングするツールを構築することでした。
それがモデルから得られた可能性はどのくらいですか。
これは 2 つの独立した理由で失敗し、どちらも致命的です。
1つ目は技術的なものです。 AI によって生成されたコンテンツの検出は信頼できません。
そしてモデルが改良されるにつれて信頼性は低くなります。 5%の検出器
偽陽性率、月に 100 件の PR がかかるプロジェクトを指す
は毎月、5人の人間の投稿者を詐欺で告発している。次のいずれかが必要です
プロジェクトの公開スレッドになると、それ以上の損失が発生します
得た。
2 つ目は概念的なものであり、より重要です。 AI が生成したコードはそうではありません。
問題。スロップについて最もうるさいメンテナーは、AI アシスタントを毎日使用しています。
日、そう言ってください。この物語全体の中で最も明確な例は、セキュリティです。
研究者はステンバーグ氏に AI 支援の研究結果の大規模なバッチを送信しました。
約 50 個の実際のバグを修正しました。同じツール、反対の結果 - なぜなら、
人間は出力を送信する前に理解して検証しました。
モデルで書かれた PR であり、
人間は正当な貢献です。誰かからの手書きの PR
AI が関与していなければ、彼らがやったことがジャンクであることを説明することはできません。
重要な変数は、コードの出所ではありません。それは、
人間は送信された内容を理解します。
プロジェクトがすでに手作業で行っていること
最も分かりやすい詳細は、これらのプロジェクトのポリシーが実際に何を要求しているのかということです。
matplotlib の貢献ガイド
実際の障害モードに直接名前を付けます。AI 出力の使用を警告します。
あなたを保証することなく

fully understand it, or without verifying it's the
correct approach — and says it will flag and reject low-value
それらの根拠に基づいて貢献します。 Ghostty の次の動きはこれ以上のものではありませんでした
検出器のいずれか。 it was a Vouch Request, where a first-time contributor
has to explain themselves in their own words — explicitly not written by
AI — PR を送信する前に。
これらはいずれも、実際には「AI を使用しない」というものではありません。の証明請求です
understanding — applied by hand, one PR at a time, spending exactly the
政策が保護する予定だった希少な資源。
そしてメンテナーも同じことを言っています。描画するとき
外部 PR を閉じた場合、
framed the move as temporary, pending better tooling.ここがその部分です
それは私を止めました: GitHub 自身の
質の低い貢献に関する公開討論、
a GitHub product manager floated, as one possible direction, defining a
set of rules or prompts and evaluating pull requests against them.の
プラットフォーム自体も同じアイデアに到達しました。書面による要求があります
まだ存在していない製品。
Restore the backpressure, don't police the origin
If the diagnosis is that the effort toll disappeared, the fix isn't to
誰がAIを使ったのか推測してください。 It's to rebuild the toll — and charge it in the right
通貨。
適切な通貨とは理解することです。そして理解度は次のような単位で測定できます。
非常に古い方法: 質問することです。
形はシンプルです。 When you open a PR, you get two or three specific
questions about that diff — questions you can only answer if you
understood the change, not if you skimmed the description.あなたは答えます。の
マージするとブロックが解除されます。
これには、検出器が決して検出しない特性があります。
それは誰も非難しません。 「これはAIが生成した」とは主張しません。あなたは尋ねます
someone to explain their own work — something any good-faith contributor
合理的で何かを見つける

いくつかのプロジェクトはすでに散文での依頼を行っています。
誤検知は安価です。質問が簡単すぎる場合は、
経験豊富な投稿者は 30 秒を失います。探知機としては最悪
事件は人間を詐欺師と呼んでいます。クイズの最悪のケースは軽いイライラです。
ツールに依存せず、将来性もあります。かどうかは関係ありません
コードはモデル、スタック オーバーフロー、またはその人自身のものから来ています。
頭。この疑問は、今から 3 モデル世代後も同じままです。
これを理解できる人間がいるでしょうか？
かつてあった場所に正確に料金所を再構築します。新しいものは追加されない
摩擦 — 常に存在し、これまでにのみ存在していた摩擦を元に戻します。
手作業で自由にできたので見えなくなりました。
クイズはAIが答えてくれます。可能ですが、差分を貼り付ける必要があります
質問をモデルに入力し、回答を読んで送信します。
これもまた通行料ですが、その通行料が重要なのです。目標は決して作ることではなかった
それは不可能です。それは無料ではないようにするためでした。
そして、無視すべきではないアクセシビリティの緊張もあります。
英語での質問は非ネイティブの投稿者にペナルティを課し、時間制限のあるテストを行う
神経発散者を罰する。これの本格的なバージョンは次のようになります。
メンテナによって構成可能で、時間制限がなく、確立されたものを免除することができます
貢献者。流暢な英語を書く人だけを通過させるフィルター
理解を測定しているのではなく、何か他のものを測定しているのです。
除外オープンソースを再現するものはすでにたくさんあります。 (気になる
これは個人的な意見です。私は英語のネイティブスピーカーではないので、英語を話すと失敗するでしょう。
私自身のアイデアの下手なバージョンです。)
公の場での会話では、2025 年と 2026 年が AI の侵入の瞬間であるとされていました
オープンソース。そのフレーミングは禁止、探知、コードをめぐる争いにつながる
出自 — そしてそれらはどれも問題を解決しません。
さらに役立つ読み物: オープンソースが実行された

何十年にもわたって高品質のフィルターを使用
誰も設計せず、ほとんど誰も気づかなかった - 理解する努力
貢献する前に。 AI がオープンソースを破壊したわけではありません。それは、
偶然のフィルターで、これから来るものを意図的に構築する必要がありました
無料で。
私は erickdevz 、フルスタック開発者です
ブラジルから。私はまさにこれを行うオープンソース ツールを構築しています。
理解ゲート
事前に投稿者に自分の PR についていくつかの質問をする
合併します。検出器ではなく、理解度チェックであり、設定可能で時間制限はありません。
まだ早いので、試してみるプロジェクトを探しています。

## Original Extract

Open source ran for decades on a quality filter no one built on purpose: the effort of understanding before contributing. That

← Home
Blog
☀️
2026-08-06 · ai, open-source, software-engineering, productivity
AI didn't break open source. It removed a filter nobody designed.
By erickdevz · full-stack developer, Viçosa, Brazil
In January 2026, the curl project shut down a bug bounty it had run
since 2019. Not because the money ran out — it had paid out more than
US$100,000 across dozens of confirmed vulnerabilities. It shut down
because the economics of triage had collapsed. Around the same window,
the Jazzband collective, which maintained 84 Python projects, announced
it was winding down. tldraw's founder flipped the repository to
auto-close every external pull request. Ghostty adopted a zero-tolerance
policy for low-effort AI submissions.
None of these projects hit trouble from a lack of people wanting to
contribute. They buckled under the opposite.
It's worth understanding why, because the correct diagnosis points to a
very different fix than the one most of the industry is building.
The filter nobody knew was there
Mitchell Hashimoto, who created Ghostty (and co-founded HashiCorp before
it), wrote the sharpest line about this in the project's updated AI
policy . His framing:
agentic programming eliminated the natural, effort-based backpressure
that used to limit low-effort contributions, and it's now far too easy
to produce large volumes of bad content with minimal effort.
Sit with the word backpressure , because it's the whole argument.
Think about what it took, up to 2023, to open a pull request against a
project you didn't already know. Clone the repo. Get it to build. Read
enough of the code to find where your change fit. Run the tests. Work out
why a test broke.
Nobody designed that as a quality system. But it functioned as one. The
effort was a toll, and the toll guaranteed a floor: almost anyone who got
as far as a PR had understood something along the way.
The toll is gone. Producing a patch that compiles and passes CI now costs
a few minutes and no understanding at all. The quality floor that existed
as a side effect of effort vanished along with the effort.
Hashimoto is emphatic that this is not an anti-AI position — Ghostty is
written with heavy AI assistance and its maintainers use AI daily. The
line isn't about the tool. It's about contributors who outsource the
thinking to a model and then hand the verification bill to the
maintainer.
The asymmetry that breaks the system
Generating code got dramatically cheaper. Reviewing code got no cheaper
at all.
That asymmetry is fatal to any system with an open front door. The
curl numbers
make it concrete: the share of vulnerability reports that turned out to
be real had run above 15% for years, then fell below 5% in 2025 as
AI-generated reports flooded in. In the first three weeks of 2026, curl
received twenty submissions — seven of them inside a single sixteen-hour
window — and, after the security team read each one, tried to reproduce
it, and traced the code paths, not one described a real vulnerability.
Daniel Stenberg's own summary, in his post
"death by a thousand slops" ,
was blunt: not only does the volume go up, the quality goes down, so you
spend more time than ever to get less out of it than ever.
Security reports are the cruelest version of this, because you can't just
ignore them. Dismissing a real vulnerability has a catastrophic cost. So
every invented report has to be investigated seriously before it can be
dismissed. The cost of generating approaches zero; the cost of
responsibly discarding stays high.
An honest footnote: curl
reopened the program
about a month later, when report quality climbed back up. Notably, what
came back wasn't the slop — it was high-volume, high-quality,
human-verified reports. Which is the whole point: the problem was never
AI. It was unreviewed output with the effort stripped out.
Why detecting AI is the wrong answer
The market's instinct was to build detectors: tools that score a PR by
how likely it is to have come from a model.
This fails for two independent reasons, and both are fatal.
The first is technical. Detecting AI-generated content isn't reliable,
and it gets less reliable as models improve. A detector with a 5%
false-positive rate, pointed at a project taking a hundred PRs a month,
accuses five human contributors of fraud every month. It takes one of
those becoming a public thread for the project to lose more than it
gained.
The second is conceptual, and it matters more. AI-generated code isn't
the problem. The maintainers loudest about slop use AI assistants every
day and say so. The clearest example in the whole saga: a security
researcher sent Stenberg a large batch of AI-assisted findings that led
to fixing around fifty real bugs. Same tool, opposite outcome — because a
human understood and verified the output before submitting it.
A model-written PR that's read, understood, tested, and defended by a
human is a legitimate contribution. A hand-typed PR from someone who
can't explain what they did is junk, without a line of AI involved.
The variable that matters isn't where the code came from. It's whether a
human understands what was submitted.
What projects are already doing by hand
The most telling detail is what these projects' policies actually ask for.
matplotlib's contributing guide
names the real failure mode directly: it warns against using AI output
without ensuring you fully understand it, or without verifying it's the
correct approach — and says it will flag and reject low-value
contributions on those grounds. Ghostty's next move wasn't a better
detector either; it was a Vouch Request, where a first-time contributor
has to explain themselves in their own words — explicitly not written by
AI — before they can submit a PR at all.
None of this is really "don't use AI." It's a request for proof of
understanding — applied by hand, one PR at a time, spending exactly the
scarce resource the policy was meant to protect.
And the maintainers say as much. When tldraw
closed external PRs , it
framed the move as temporary, pending better tooling. Here's the part
that stopped me: in GitHub's own
public discussion on low-quality contributions ,
a GitHub product manager floated, as one possible direction, defining a
set of rules or prompts and evaluating pull requests against them. The
platform itself reached for the same idea. There's written demand for a
product that doesn't quite exist yet.
Restore the backpressure, don't police the origin
If the diagnosis is that the effort toll disappeared, the fix isn't to
guess who used AI. It's to rebuild the toll — and charge it in the right
currency.
The right currency is understanding. And understanding is measurable in a
very old way: by asking questions.
The shape is simple. When you open a PR, you get two or three specific
questions about that diff — questions you can only answer if you
understood the change, not if you skimmed the description. You answer. The
merge unblocks.
This has properties a detector never will:
It accuses no one. You don't claim "this was AI-generated." You ask
someone to explain their own work — something any good-faith contributor
finds reasonable, and something several projects already ask for in prose.
False positives are cheap. If the questions are too easy for an
experienced contributor, they lose thirty seconds. A detector's worst
case is calling a human a fraud; a quiz's worst case is mild annoyance.
It's tool-agnostic and future-proof. It doesn't matter whether the
code came from a model, from Stack Overflow, or from the person's own
head. The question stays the same three model generations from now: is
there a human who understands this?
It rebuilds the toll exactly where it used to be. It doesn't add new
friction — it puts back the friction that was always there, and only ever
looked invisible because it came free with the manual work.
A quiz can be answered by AI. It can — but that requires pasting the diff
and the questions into a model, reading the answer, and submitting it.
That's a toll again, and the toll is the point. The goal was never to make
it impossible; it was to make it not free.
And there's an accessibility tension that shouldn't be waved away:
questions in English penalize non-native contributors, and timed tests
penalize neurodivergent people. Any serious version of this has to be
configurable by the maintainer, untimed, and able to exempt established
contributors. A filter that only passes people who write fluent English
isn't measuring understanding — it's measuring something else, and
reproducing an exclusion open source already has plenty of. (I care about
this one personally: I'm not a native English speaker, and I'd fail a
badly built version of my own idea.)
The public conversation framed 2025 and 2026 as the moment AI invaded
open source. That framing leads to bans, detectors, and fights over code
provenance — and none of them solve the problem.
The more useful reading: open source ran for decades on a quality filter
nobody designed and almost nobody noticed — the effort of understanding
before contributing. AI didn't break open source. It removed an
accidental filter, and forced us to build on purpose what used to come
for free.
I'm erickdevz , a full-stack developer
from Brazil. I'm building an open-source tool that does exactly this:
comprehension-gate
asks a contributor a few questions about their own PR before it can
merge. Not a detector — a comprehension check, configurable and untimed.
It's early and I'm looking for projects to try it on.
