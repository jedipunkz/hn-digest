---
source: "https://zernie.com/blog/stop-writing-claude-md-rules/"
hn_url: "https://news.ycombinator.com/item?id=48830917"
title: "Stop Writing Claude.md Rules. Write Linting Rules Instead"
article_title: "Stop Writing CLAUDE.md Rules. Write Linting Rules Instead. — Ernie"
author: "speckx"
captured_at: "2026-07-08T12:14:55Z"
capture_tool: "hn-digest"
hn_id: 48830917
score: 1
comments: 0
posted_at: "2026-07-08T12:13:17Z"
tags:
  - hacker-news
  - translated
---

# Stop Writing Claude.md Rules. Write Linting Rules Instead

- HN: [48830917](https://news.ycombinator.com/item?id=48830917)
- Source: [zernie.com](https://zernie.com/blog/stop-writing-claude-md-rules/)
- Score: 1
- Comments: 0
- Posted: 2026-07-08T12:13:17Z

## Translation

タイトル: Claude.md ルールの作成を停止します。代わりにリンティング ルールを作成する
記事のタイトル: CLAUDE.md ルールの作成を停止します。代わりにリンティングルールを作成してください。 — アーニー
説明: CLAUDE.md ルールは、確率モデルが尊重するかもしれない願いです。 lint ルールは、過去にマージできない法律です。エージェントがレビューできる量を超えるコードを出荷した場合は、規約の文書化をやめて、規約の強制を開始してください。

記事本文:
CLAUDE.md ルールの作成を停止します。代わりにリンティングルールを作成してください。 — Ernie Ernie Home Skills Blog AI Playbook 話しましょう ← 記事に戻る EN / RU 2026 年 7 月 4 日 CLAUDE.md ルールの記述を停止します。代わりにリンティングルールを作成してください。
誰もが最初に持つ本能
エージェントがコードベースで愚かなことをします。 2 四半期前に移行した設計システムからインポートします。 console.log を運用パスにドロップします。これは、ループ内の try-catch 内にスイッチを含む 150 行の関数を作成します。
だからあなたは当然のことをします。 CLAUDE.md を開いて書き留めます。 「常に shadcn/ui を使用してください。 console.log は決して使用せず、ロガーを使用してください。関数は小さくしてください。」 You write it clearly. You write it in bold. Maybe you add a skill.
I did this for weeks. Everyone does. Vercel でさえ、スキル ライブラリを出荷していました。40 を超える React パフォーマンス ルールは美しく書かれ、エージェントが読み取ることができる SKILL.md ファイルとして構造化されています。 Genuinely good work.
そしてとにかくエージェントはまた愚かなことをします。
CLAUDE.md ルールは実際にはルールではありません。これは、正しいカードが引かれることを期待して、確率システムに入力する提案です。ほとんどの場合、そうなります。最初の試行で美しく慣用的なコードが得られると、自分が天才になったような気分になります。
しかし、この文では「ほとんどの場合」は多くの作業を行っています。エージェントが 1 週間でレビューするよりも多くのコードを午後に出荷している場合、「ほとんどの場合」は安全性ではありません。 It's a countdown.サイコロを十分に振ると、最終的にモデルはインデックスのないクエリをホット パスにドロップし、金曜日の午前 2 時にデータベースが溶ける時刻がわかります。
同じ PR、同じ規則、それを強制する 2 つの方法。散文では、ルールは CLAUDE.md に保存されており、エージェントはほとんどの場合それを尊重します。リントレーンでは CI チェックです。数十個出荷して、実際に最大 12% のずれが発生するかどうかを観察してください。

土地。
確率的に、それを尊重する可能性は最大 12%
まだ何も滑っていません。発送を続けてください。
違反する可能性は 0% — マージできません
0 が本番に到達しました。まだブロックするものはありません。
「ほとんどの場合」は安全性ではなく、カウントダウンです。法律はカウントダウンをすることはありません。
それが CLAUDE.md のことです。 It's firmly in pirate-code territory, more what you'd call guidelines than actual rules. It explains the why , which genuinely helps the agent get it right on the first try. What it can't do is make getting it wrong impossible . And once you're producing code faster than any human can read it, the gap between "usually right" and "can't be wrong" is the whole ballgame.
これはすでに解決しました。数十年前。
これは私を魅了します。私たちはこれを知っています。私たちはそれを何年も前から知っていました。
Unit tests exist because humans forget edge cases. Linters exist because humans can't agree on where to put a curly brace and shouldn't have to. CI exists because "works on my machine" stopped being funny after the third production outage. The entire history of software tooling is us admitting, over and over, that good intentions don't survive contact with a real codebase.そこで私たちはそれらに依存しないマシンを構築しました。
And then LLMs showed up and we collectively forgot all of it. 「本当に良い CLAUDE.md を書いてください。」 「スキルを追加するだけです。」 「もっとよく促してください。」 We took the one actor in the building that is definitionally probabilistic and decided the way to control it was… a nicely worded document.
CLAUDE.md explains the why and helps the agent get it right on the first try. lint ルールにより、間違いが起こらないようになっています。スキルを使うとスピードが上がります。リンターはあなたを正直に保ちます。 1 つしか持てない場合は、リンターを使用してください。
The reframe that changed how I work: stop trying to teach the agent your conventions.まきスタート

間違っているとマージに失敗します。可能なすべてのルールを散文から実行可能なチェックに移します。エージェントは、CI が単に拒否するような内容を覚えておく必要はありません。
レバレッジが最も高い単一のルール: 複雑さの上限
これを読んだ後に何かをするなら、これをしてください。
スタイルルールやインポート順序のことを言っているのではありません。つまり、コードの一部がどの程度大きく、どの程度複雑に絡み合っても許容されるかについて、厳しい制限が設けられているということです。 ESLint は何年もこれらを出荷していましたが、ほとんど誰もそれらを有効にしていないことが判明しました。
複雑さ — 循環的複雑さを制限する
max- Depth — ネストできる深さを制限します
max-lines-per-function — 強制的に分解します
max-params — インターフェイスを狭く保ちます
max-statements — 関数が一度に 12 のことを実行するのを停止します
SonarJS の認知的複雑さ — ネストされた条件文をより賢く扱う
これらがなぜ重要なのかを楽しい方法で学びました。これらがなければ、エージェントは、6 レベルのネスト、3 つの早期リターン、およびループ内の try-catch 内に埋め込まれたスイッチを備えた単一の 150 行関数を快く渡します。コンパイルされます。テストに合格します。それも機能します。その後、次のエージェントがそれに触って悪化させ、今ではエージェントが 2 人います。
{
「ルール」: {
"複雑さ" : [ "エラー" , 10 ],
"最大深度" : [ "エラー" , 3 ],
"関数ごとの最大行数" : [ "エラー" , 40 ],
"max-params" : [ "エラー" , 4 ],
"最大ステートメント" : [ "エラー" , 15 ]
}
コピー
これを --max-warnings=0 と組み合わせて、警告が失敗であることを確認し、何が起こるかを監視します。エージェントは 150 行のモンスターを渡すのをやめ、ヘルパーを抽出し、懸念を解消し始めます。厳しく始めて、積極的に痛む場所だけを緩めます。
ただし、このルールが何であるかについては正直に言ってください。これは console.log を禁止するのと同じ種類のものではないからです。これは法則です。シンボルがファイル内にあるかどうかのどちらかであり、規則に違反するコードのバージョンは存在しません。

dはまだ正常に読み取れます。複雑さの上限は、「これは読みやすいか」を表す数値です。そして、プロキシをハード ゲートにした瞬間、コードを生成するものはすべて、目標ではなくプロキシを最適化します。 max-lines-per-function: 40 を見つめるエージェントには、準拠する 2 つの方法があります。関数を正直に分解するか、関数を handleThingPart2() とカウンタをクリアするために名前を付けられたいくつかの使い捨てヘルパーに分割します。 2番目も喜んで実行します。実際のお気に入りの動きは脱出ハッチです。
// eslint-disable-next-line max-lines-per-function コピー
したがって、複雑さをゲートする場合は、そのドアも閉める必要があります。無効なコメント ( eslint-plugin-eslint-comments ) を禁止または追跡するか、エージェントが全体を迂回し、誰も信頼できない緑色のチェックマークを渡します。誰もが沈黙するために静かに学んだルールは、ルールがまったくないより悪いです。
避難ハッチを閉め、数字を教典ではなく出発点として扱い、キャップは依然として私が知る限り最もレバレッジの高いルールです。ただし、プロキシを不変式と間違えないでください。これは、何年も前、人間が 150 行の関数を書いていた時代に、私が最もオンにしておきたかった機能です。
ある時点から、何か問題が発生するたびに、「これは lint ルールでしょうか?」という疑問が頭の中でループし始めました。
そのスイッチが切り替わると、レビュアーではなくなります。あなたは、与えられた間違いが二度と起こらないようにする人になります。 「またこのスプリント」ではありません。これまで。
最初のものが最も難しいです。その後、彼らは雪だるま式に増えます。私のコードベースからの実際のもの: エージェントは、Datadog にルーティングするカスタム ロガーの代わりに、console.log を本番コードに追加し続けました。毎。シングル。時間。 CLAUDE.mdをどれだけ使っても修正できませんでした。 10 行の lint ルールは、forbid console.log 、suggest logger.error 、done を実行しました。それ以来、それについては考えていません。
そういう形なんです。レキュ

rring PR コメントはまだ書いていない lint ルールです。同じレビューメモを 3 回入力するたびに、それはレビューをやめて強制を開始するという合図です。
「カスタム ルールが 50 個あります」と聞くと、人々は反発します。私もそうでした。
はい、それらのルールの中には間違っているものもあります。誰かが悪い問題に遭遇し、イライラして、それを変更するために PR を開いた場合。そして、その PR の中で、チームがこれまで大声で話したことのない建築に関する会話をすることになります。実際にこのパターンを許可すべきでしょうか?なぜ禁止したのでしょうか?その PR は、それを引き起こしたルールよりも価値があります。このルールは、スクロールして離れていく Slack スレッド以外の場所で会話を強制的に発生させるものにすぎません。
失敗モードは PR を開かないときです。つまり、 eslint-disable がドロップされ、出荷され、ルールは沈黙のうちに終了します。そのため、注目すべき数値はルール数ではなく、無効化数です。抑圧が高まっている場合は、ルールが間違っているか、誰もルールをレビューしていないかのどちらかであり、どちらも調べる価値があります。チーム全体が回避するために学んだルールは強制ではありません。それは装飾です。
より厳格なルールの採用に対する昔からの反対意見は、常に「既存の違反が多すぎて修正できない」というものでした。エージェントの場合、その言い訳はほとんど通用しません。違反を指摘し、一括で優先順位を付けます。 AI と codemod により、移行の機械的な半分が 4 分の 1 から午後に変わります。高価なままになる半分は人間的なものです。ルールが正しいことを全員に同意させ、切り替えたときに開いていたすべてのブランチにマージ競合税を課す必要があります。それは常に実際のコストでした。編集を自動化することで、編集が多忙な作業の影に隠れてしまうことを防ぐことができます。
ESLint と言い続けているのは、それが私の世界だからです。この原則は ESLint に関するものではありません。 RuboCop、Ruff、clippy、golangci-lint。同じ考えで、毒を選んでください。オフ

-the-shelf リンターは、構文と一般的なバグ クラスを無料で処理します。何年も無視していたプラグイン (SonarJS、ユニコーン、完璧主義者) をオンにすると、多くの場合、午後になると間違いのカテゴリ全体が削除されます。
あなたが所有する部分はアーキテクチャであり、その多くは既製品でもあります。 「このレイヤーはそのレイヤーからインポートできません」は no-restricted-imports または eslint-plugin-boundaries です。 「すべての API ハンドラーは認証ラッパーを通過する」は、制限のない構文セレクターです。ルール コードを 1 行記述する前に、これらをオンにします。オーダーメイドのように見えるもののほとんどはそうではありません。
カスタム ルールは、かつては最も年上の 3 人の頭の中にのみ存在していた、真にあなただけの慣例のためのものです。そして、それをオーサリングすることは、思っているほど怖くありません。実際のルール (たとえば、インストルメント化されたクライアントを優先して生のフェッチを禁止する) は、基本的にセレクターとレポートです。
デフォルトのエクスポート {
作成 (コンテキスト) {
戻り値 {
"CallExpression[callee.name='fetch']" (ノード) {
コンテキスト 。レポート({
ノード、
メッセージ: 「生のフェッチではなく、apiClient を使用します。トレースと認証に接続されます。」 、
})
}、
}
}、
コピー
それがすべてです。一つ書けば恐怖は消え去る。その後、エージェントは雰囲気だけから吸収する方法がないすべての慣例について、いたるところでルールを目にするようになります。
糸くずは魔法ではないので、天井については正直に言いたいと思います。
lint ルールは、すでに見たものだけをキャッチできます。それはあなたが犯し、二度と犯さないと決めた間違いを暗号化します。あなたがまだ打っていないミスはまだ存在しており、それを監視するルールはありません。そして、lint は静的です。コードを読み取るだけで、実行することはできません。レート制限ミドルウェアがファイル内で定義されていることを喜んで確認しますが、それが実際にマウントされていないことはわかりません。コードはすぐそこにあります。ただ走っていないだけです。
それをキャッチするのはディです

重要な問題: 見たことのないもの、およびアプリが実際に起動したときにのみ表示されるもの。それはエージェントをより良く指導することではありません。それはエージェントが生成したものを検証することであり、独自のツールと独自の頭痛の種を備えた別の分野です (特に、スタックの一部の層は他の層よりも検証するのがはるかに安価です。私はそのループを強化することについて記事全体を書きました)。そして、その一部は完全にテストの上流にあり、最悪のバグをまったく表現できないように状態をモデリングします。どちらも独自の作品です。これではありません。
なぜなら、この記事は入力側、つまりエージェントに何をすべきかを伝える方法だけを扱っていたからです。短いバージョン: 書き留めたり期待したりするのはやめましょう。それを法律にしてください。
リンターは眠らず、疲れません。これは金曜日の午前 2 時に私が言える以上のことです。
CLAUDE.md を参照してください。 Пизи правила линтера。
Первая мысль, которая приходит в голову всем
Агент делает в твоём коде какую-нибудь глупость。 Импортирует из дизайн-системы, с которой ты слез два квартала назад. console.log が表示されます。 150 個のスイッチを試してキャッチしてください。
И ты делаегочевидное. CLAUDE.md を参照してください。 «Всегда используй shadcn/ui. console.log を参照してください。

[切り捨てられた]

## Original Extract

A CLAUDE.md rule is a wish a probabilistic model might honor. A lint rule is a law it can't merge past. When agents ship more code than you can review, stop documenting your conventions and start enforcing them.

Stop Writing CLAUDE.md Rules. Write Linting Rules Instead. — Ernie Ernie Home Skills Blog AI Playbook Let's talk ← Back to articles EN / RU July 4, 2026 Stop Writing CLAUDE.md Rules. Write Linting Rules Instead.
The instinct everyone has first
An agent does something dumb in your codebase. It imports from the design system you migrated off of two quarters ago. It drops a console.log into a production path. It writes a 150-line function with a switch inside a try-catch inside a loop.
So you do the obvious thing. You open CLAUDE.md and you write it down. "Always use shadcn/ui. Never use console.log , use the logger. Keep functions small." You write it clearly. You write it in bold. Maybe you add a skill.
I did this for weeks. Everyone does. Even Vercel shipped a skill library — 40+ React performance rules, beautifully written, structured as SKILL.md files for agents to read. Genuinely good work.
And then the agent does the dumb thing again anyway.
A CLAUDE.md rule isn't really a rule. It's a suggestion you feed into a probabilistic system, hoping it draws the right card. Most of the time it does. You get beautiful, idiomatic code on the first try and you feel like a genius.
But "most of the time" is doing a lot of work in that sentence. When an agent is shipping more code in an afternoon than you'd review in a week, "most of the time" isn't a safety property. It's a countdown. Roll the dice enough and eventually the model drops an unindexed query into a hot path, and you find out when the database melts at 2 AM on a Friday.
Same PRs, same convention, two ways to enforce it. In the prose lane the rule lives in CLAUDE.md — the agent honors it most of the time . In the lint lane it's a CI check. Ship a few dozen and watch where the ~12% that slip actually land.
~12% chance to honor it — probabilistically
Nothing slipped — yet. Keep shipping.
0% chance to violate — it cannot merge
0 reached prod. Nothing to block yet.
“Most of the time” is not a safety property — it’s a countdown. The law never counts down.
That's the thing about a CLAUDE.md. It's firmly in pirate-code territory, more what you'd call guidelines than actual rules. It explains the why , which genuinely helps the agent get it right on the first try. What it can't do is make getting it wrong impossible . And once you're producing code faster than any human can read it, the gap between "usually right" and "can't be wrong" is the whole ballgame.
We already solved this. Decades ago.
This one gets me. We know this. We've known it for years.
Unit tests exist because humans forget edge cases. Linters exist because humans can't agree on where to put a curly brace and shouldn't have to. CI exists because "works on my machine" stopped being funny after the third production outage. The entire history of software tooling is us admitting, over and over, that good intentions don't survive contact with a real codebase. So we built machines that don't rely on them.
And then LLMs showed up and we collectively forgot all of it. "Just write a really good CLAUDE.md." "Just add more skills." "Just prompt it better." We took the one actor in the building that is definitionally probabilistic and decided the way to control it was… a nicely worded document.
CLAUDE.md explains the why and helps the agent get it right on the first try. A lint rule makes sure it can't get it wrong. Skills speed you up. Linters keep you honest. If you can only have one — take the linter.
The reframe that changed how I work: stop trying to teach the agent your conventions. Start making the wrong thing fail to merge . Move every rule you can from prose into an executable check. The agent doesn't have to remember something the CI will simply refuse.
The single highest-leverage rule: complexity caps
If you do one thing after reading this, do this one.
I don't mean style rules or import ordering. I mean hard caps on how big and how tangled a piece of code is allowed to get. It turns out ESLint has shipped these for years and almost nobody turns them on:
complexity — caps cyclomatic complexity
max-depth — limits how deeply you can nest
max-lines-per-function — forces decomposition
max-params — keeps interfaces narrow
max-statements — stops a function from doing twelve things at once
SonarJS cognitive-complexity — smarter about nested conditionals
I learned why these matter the fun way. Without them, an agent will cheerfully hand you a single 150-line function with six levels of nesting, three early returns, and a switch buried inside a try-catch inside a loop. It compiles. It passes tests. It even works . Then the next agent touches it, makes it worse, and now you have two of them.
{
"rules" : {
"complexity" : [ "error" , 10 ],
"max-depth" : [ "error" , 3 ],
"max-lines-per-function" : [ "error" , 40 ],
"max-params" : [ "error" , 4 ],
"max-statements" : [ "error" , 15 ]
}
} Copy
Pair that with --max-warnings=0 so a warning is a failure, and watch what happens. The agent stops handing you 150-line monsters and starts extracting helpers and pulling concerns apart. Start strict, loosen only where it actively hurts.
But be honest about what this rule is , because it isn't the same kind of thing as forbidding console.log . That one is a law — the symbol is either in the file or it isn't, and there's no version of the code that breaks the rule and still reads fine. A complexity cap is a proxy : a number standing in for "is this readable." And the moment you make a proxy a hard gate, whatever's generating the code optimizes the proxy, not the goal. An agent staring down max-lines-per-function: 40 has two ways to comply — decompose the function honestly, or shatter it into handleThingPart2() and a couple of single-use helpers named to clear the counter. It will happily do the second. Its actual favorite move is the escape hatch:
// eslint-disable-next-line max-lines-per-function Copy
So if you gate on complexity you have to shut that door too: ban or track disable comments ( eslint-plugin-eslint-comments ), or the agent routes around the whole thing and hands you a green checkmark nobody should trust. A rule everyone quietly learned to silence is worse than no rule at all.
Shut the escape hatch, treat the numbers as a starting point rather than scripture, and caps are still the highest-leverage rule I know — just don't mistake the proxy for the invariant. It's the one I most wish I'd turned on years ago, back when it was humans writing the 150-line functions.
At some point a question started running on a loop in my head, every time something went wrong: can this be a lint rule?
Once that switch flips, you stop being a reviewer. You become the person who makes sure a given mistake can never happen again. Not "again this sprint." Ever.
The first one is the hardest. After that they snowball. A real one from my codebase: agents kept adding console.log to production code instead of the custom logger that routes to Datadog. Every. Single. Time. No amount of CLAUDE.md fixed it. A 10-line lint rule did: forbid console.log , suggest logger.error , done. I have not thought about it since.
That's the shape of it. A recurring PR comment is a lint rule you haven't written yet. Every time you find yourself typing the same review note for the third time, that's the signal: stop reviewing, start enforcing.
People hear "we have 50 custom rules" and recoil. I did too.
Yes, some of those rules will be wrong. Someone hits a bad one, gets annoyed, and opens a PR to change it. And right there, in that PR, you end up having the architectural conversation your team never quite had out loud. Should we actually allow this pattern? Why did we ban it? That PR is worth more than the rule that triggered it. The rule is just what forced the conversation to happen somewhere other than a Slack thread that scrolls away.
The failure mode is when they don't open the PR — they drop an eslint-disable , ship, and the rule dies in silence. That's why the number to watch is the disable count, not the rule count. If suppressions are climbing, either the rule is wrong or nobody's reviewing them, and both are worth finding out. A rule the whole team learned to route around isn't enforcement. It's decoration.
The old objection to adopting stricter rules was always "too many existing violations to fix." With an agent, that excuse is mostly dead. Point it at the violations and let it triage them in bulk; AI plus codemods turns the mechanical half of a migration from a quarter into an afternoon. The half that stays expensive is the human one — getting everyone to agree the rule is right, and eating the merge-conflict tax on every branch that was open when you flipped it. That was always the real cost. Automating the edits just stops it from hiding behind the busywork.
I keep saying ESLint because that's my world. The principle is not about ESLint. RuboCop, Ruff, clippy, golangci-lint. Pick your poison, same idea. Off-the-shelf linters handle syntax and the common bug classes for free; turning on plugins you'd been ignoring for years (SonarJS, unicorn, perfectionist) is often an afternoon that deletes entire categories of mistake.
The part that's yours is architecture, and even a lot of that ships off the shelf. "This layer can't import from that one" is no-restricted-imports or eslint-plugin-boundaries ; "every API handler goes through the auth wrapper" is a no-restricted-syntax selector. Turn those on before you write a single line of rule code. Most of what feels bespoke isn't.
Custom rules are for the conventions that are genuinely yours, the ones that used to live only in the heads of your three most senior people. And authoring one is less scary than it sounds. A real rule — say, banning raw fetch in favor of your instrumented client — is basically a selector and a report :
export default {
create ( context ) {
return {
"CallExpression[callee.name='fetch']" ( node ) {
context . report ({
node ,
message : "Use apiClient, not raw fetch — it wires in tracing and auth." ,
})
},
}
},
} Copy
That's the whole thing. Write one and the fear evaporates; after that you start seeing rules everywhere, for every convention an agent has no way to absorb from vibes alone.
I want to be honest about the ceiling, because linting isn't magic.
A lint rule can only catch what you've already seen. It encodes a mistake you've made and decided never to make again. The mistake you haven't hit yet is still out there, and no rule is watching for it. And linting is static: it reads the code, it can't run it. It'll happily confirm your rate-limiting middleware is defined in a file while having no idea it was never actually mounted. The code is right there. It's just not running.
Catching that is a different problem: the stuff you haven't seen, and the stuff that only surfaces when the app actually boots. That's not about instructing the agent better — it's about verifying what it produced, a different discipline with its own tools and its own headaches (not least that some layers of your stack are far cheaper to verify than others — I wrote a whole piece on tightening that loop ). And some of it is upstream of testing entirely: modeling state so the worst bugs can't be represented at all . Both are their own pieces; not this one.
Because this piece was only ever about the input side: how you tell the agent what to do. The short version: stop writing it down and hoping. Make it a law.
Linters don't sleep, and they don't get tired. That's more than I can say for myself at 2 AM on a Friday.
Хватит писать правила в CLAUDE.md. Пиши правила линтера.
Первая мысль, которая приходит в голову всем
Агент делает в твоём коде какую-нибудь глупость. Импортирует из дизайн-системы, с которой ты слез два квартала назад. Роняет console.log прямо в продовый путь. Пишет функцию на 150 строк со switch внутри try-catch внутри цикла.
И ты делаешь очевидное. Открываешь CLAUDE.md и записываешь это туда. «Всегда используй shadcn/ui. Никогда не используй console.log , т

[truncated]
