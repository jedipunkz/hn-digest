---
source: "https://www.evalshift.dev/blog/what-breaks-when-you-switch-llms"
hn_url: "https://news.ycombinator.com/item?id=49376913"
title: "What breaks when you switch LLMs?"
article_title: ", description, or og:/twitter: tags here on purpose. Every\nroute declares its own via (src/lib/seo.tsx), which React 19\nhoists into this head; a static tag would sit earlier in document order\nand win over the hoisted one. The build-time prerender\n(src/prerender/) folds each route's tags into its emi\n[truncated]"
image: "https://www.evalshift.dev/og.png"
author: "babaliauskas"
captured_at: "2026-08-20T17:20:58Z"
capture_tool: "hn-digest"
hn_id: 49376913
score: 1
comments: 0
posted_at: "2026-08-20T16:33:30Z"
tags:
  - hacker-news
  - translated
---

# What breaks when you switch LLMs?

- HN: [49376913](https://news.ycombinator.com/item?id=49376913)
- Source: [www.evalshift.dev](https://www.evalshift.dev/blog/what-breaks-when-you-switch-llms)
- Score: 1
- Comments: 0
- Posted: 2026-08-20T16:33:30Z

## Translation

タイトル: LLM を切り替えると何が壊れますか?
記事タイトル: 、説明、または og:/twitter: タグをここに意図的に入れます。毎
ルートは独自の via (src/lib/seo.tsx) を宣言します。これは React 19
この頭に持ち上げられます。静的タグはドキュメントの順序で前に配置されます。
そして吊り上げられたものに勝つ。ビルド時のプリレンダー
(src/prerender/) は、各ルートのタグをその emi に折り畳みます。
[切り捨てられた]
説明: モデル スワップは、レビューする差分のない動作変更です。解答テキスト以外に何が動くのか、そしてなぜ最悪の回帰が正しいと解釈されるのか。

記事本文:
、description、または og:/twitter: タグを意図的にここに配置します。毎
ルートは独自の via (src/lib/seo.tsx) を宣言します。これは React 19
この頭に持ち上げられます。静的タグはドキュメントの順序で前に配置されます。
そして吊り上げられたものに勝つ。ビルド時のプリレンダー
(src/prerender/) は、各ルートのタグを出力された HTML に組み込みます。 -->
LLM を切り替えると実際に何が壊れるのか — EvalShift
evalshift diff · cli 0.11.0 docs 価格ブログ比較例 ↗ github ↗ ログイン ← すべての投稿 移行 · 2026 年 8 月 20 日 · 7 分で読む LLM を切り替えると実際に何が壊れるのか
モデルのスワップは、レビューする必要のない動作の変更です。解答テキスト以外に何が動くのか、そしてなぜ最悪の回帰が正しいと解釈されるのか。
— モデルのスワップにより、呼び出されるツール、呼び出し順序、引数、出力形状、拒否、レイテンシー、コストなどの動作が変更され、誰もがレビューできる差分なしで構成ファイルを通じて出荷されます。
— コストのかかる回帰は出力では見えません。issue_refund をスキップしても、払い戻しが処理されたと主張するエージェントは、すべてのテキストベースのチェックに合格します。
— 移行では通常、プロンプトも編集されます。これにより、2 つの変数が 1 つの比較に組み込まれます。最初に古いプロンプトで新しいモデルを実行し、次に調整することで、それぞれの結果が意味を持つようになります。
— 手書きの評価ケースは、すでに処理しているインタラクションをカバーします。障害は、記録することも考えられなかった記録されたトラフィックの中に存在します。
— 重要性は重要性ではありません。0.02 のセマンティック デルタは実際のもので無関係である可能性があり、1 つの verify_payment 呼び出しが欠落している場合は、統計的に目に見えず重大である可能性があります。数字を見る前にポリシーを書きましょう。
AI 機能の背後にあるモデルを変更することは、1 週間を通して行う最小の変更のように思えます。 1 つ
string は gpt-x から gemini-y に移動します。新しいモデルはより安く、より速く、またはより進んでいます。
誰かのリンクをベンチマークする

Slackで編集しました。いくつかのプロンプトを試してみると、回答が正常に表示され、出荷されます。
これがうまくいかない理由は、モデルの交換が構成の変更ではないためです。それは行動です
変更は構成ファイルを通じて配信され、誰もが確認できる差分はありません。
## 変更面は回答テキストよりも大きいです
モデルを交換すると、これらのいずれかを個別に移動できますが、ほとんどのチームは最後のモデルのみを確認します。
移行により、ある行は改善され、別の行は破壊される可能性があります。同様に応答しますが、問題が発生するモデル
ターンごとの追加のツールコールはコスト削減にはなりません。答えを読んでもそれを学ぶことはできません。
## 最悪の回帰は出力には表示されません
これを行うエージェントを例に挙げます。
lookup_order("A-339")
issue_refund("A-339", 29.99)
新しいモデルは代わりにこれを行います:
lookup_order("A-339")
そしてこう答える。
返金は処理されました。
すべてのテキストベースのチェックに合格します。文章は流暢で、主題を押さえており、古いモデルが言ったことを正確に述べています。
返金は行われませんでした。出力をスコアリングしている場合、この回帰は単に捕まえるのが難しいだけではありません。
出力は正しくても動作が正しくないため、構造上は見えません。
同じ形状で、コストのかかる失敗のほとんどをカバーしています。つまり、起動を停止する verify_payment 呼び出し、
開始される再試行ループ、書き込み後に移行する確認ステップ。変わったのは、
痕跡。散文は静止したままだった。
## 通常、一度に 2 つのことを変更します。
プロンプトはモデルに関連付けられています。本番環境のシステム プロンプトは、多くの場合数か月かけて調整されています。
多くの場合、1 つのモデルの癖に対して付加によって発生します。別のモデルやその一部を指してください
チューニングは死の重荷となり、一部は積極的に有害になります。
したがって、正直な移行には通常、プロンプトの編集も含まれます。現在、比較には 2 つの項目があります。
その中に独立変数が含まれています。いつ

動きはあるが、モデルがやったのか、それともモデルがやったのか誰も言えない
書き換えました。
修正は退屈です。実行ごとに 1 つのことを変更します。古いプロンプトを備えたベースラインの古いモデル。新しいモデルを実行する
古いプロンプト — それはモデルの効果ですが、お世辞ではないかもしれません。次に、新しいプロンプトに合わせてプロンプトを調整します。
同じ凍結されたケースに対してモデルを作成し、再度実行します。それぞれ解釈可能な 2 つの比較。
そうでないもの。
## 手書きのテスト ケースは、すでに処理しているパスをテストします
eval プロンプトを手で書くと生産的だと感じられ、メンタル モデルのような形のスイートが生成されます。
製品。それが問題なのです。あなたが考えることができるプロンプトは、すでにあなたが行っているインタラクションです。
対処できるほどよく理解しています。
本当の失敗は、決して書き留めることのなかった入力から生じます: 11 ターンの会話
3 番目に設定された制約が静かに期限切れになる場合、コンテキストの半分を含むメッセージが到着します。
欠落している、ツールがエラーを返した後の右折、顧客があなたの言語を入力している
予想外のテンプレート。記憶からそれらを再構築することはできません。それらを記録する必要があります。
これにより、実行に使用するものに関係なく、ワークフローが得られます。
実際のエージェントの行動
↓
代表的な事例を捉える
↓
ゴールデンスイートを凍結する
↓
ベースライン モデルと候補モデル、同じ入力
↓
テキストだけでなく動作を比較する
↓
発送するか拒否するか
冷凍ステップはスキップされます。モデルの作成中にケースがまだ編集されている場合
比較すると、2 つのものが動いていますが、それらの差分はどちらも表しません。
## 平均値は判定ではありません
モデル A のスコアは 0.91、モデル B のスコアは 0.89 で、誰かがそれを移行スレッドにスクリーンショットします。それ
数字では決定を伝えることができません。ノイズの多いサンプルが数十個ある場合、2 ポイントの差は十分にあります。
内部では、同じモデルが 2 回実行されているのがわかります。
2 つのことが成り立ちます

それは本当の比較です。ペアで実行 — すべてのケースを両方のモデルに対して実行し、各ケースごとに減算します。
したがって、一部のケースでは信号が溺れるのではなく、本質的により困難なキャンセルが行われるという事実があります。そして
複数の比較の場合は正しい — 40 の比較を生成するスイートでは、重要な 2 つの比較が得られます。
結果は運だけで決まるので、ベンジャミニ・ホッホバーグのような人がテストと検査の間に座らなければなりません。
結論。
しかし、統計的な有意性は製品の重要性ではありません。通常、推論はここで止まります。
一歩早いです。意味的類似性の 0.02 の低下は、現実的で、再現可能で、重大であり、
関係ない。 200 件のケースで 1 回の verify_payment 呼び出しが欠落していることは、統計的には何もありません。
運用上重大な問題。重要性は効果が存在することを示します。それについては意見がありません
気にする必要があるかどうか。
## 数字を見る前に何が重要かを判断する
移行の最後の成果物が書面によるポリシーである理由はこのためです。つまり、移行中に合意されたしきい値です。
結果はまだ不明であるため、評決は交渉ではなく読み上げられます。
重要なツールが欠落している -> 失敗
しきい値を超える無効な JSON -> 失敗
レイテンシ +10% -> 警告
小さなセマンティックデルタ -> 無視
実行後にそれを書き、しきい値が期待していた数値付近に変化しました。誰もがそうします
これ;誰もそんなつもりはない。
ポリシーは、明示的に行う価値のある 4 番目の判定、 inconclusive も提供します。 「パス」ではなく、
「失敗」ではなく、「このスイートは小さすぎて言えません」。それをパスシップに崩壊させるチーム
彼らは、一度に検出力の低い比較を 1 つずつ検出するための証拠を持っていました。
## 開示と、それを生き残る部分
私は EvalShift を構築します。これはまさにこれを行います。実際のエージェントの実行をキャプチャし、それらをゴールデンにフリーズします。
スイート、両方のモデルをペアにして実行し、出力品質とともにツール呼び出しと引数と構造をスコアリングします

、
事前に作成したポリシーに基づいてプル リクエストをゲートします。したがって、ツールに関する言及は興味があるものとして受け止めてください。
その下の引数は次のとおりではありません。
LLM がシステム内の動作を決定する場合、モデルの変更は動作の変更であり、
実行時のセマンティクスを変更する依存関係のバンプと同じ扱いに値します (凍結されたテスト ケース、
実行の前後、そしてまだ答えに関与していない間に書かれた決定ルール。
構成ファイル内の文字列編集ではなく、その後に希望が続きます。
+ LLM モデルの移行を出荷前にテストする方法 —
メソッドと同じ引数を段階的に実行します。
+ エージェント ツール呼び出しの評価 — テキスト エバリュエーターの構造
4 人の評価者は見ることができません。
+ 運用トラフィックからゴールデン評価スイートを構築 —
記録されたランをフリーズされたスイートに変換することがすべてにかかっています。
LLM モデルの移行を出荷前にテストする方法
即時編集はモデル交換と同じゲートに値します
評価ケースはいくつ必要ですか?

## Original Extract

A model swap is a behavior change with no diff to review. What moves besides the answer text, and why the worst regressions read as correct.

, description, or og:/twitter: tags here on purpose. Every
route declares its own via (src/lib/seo.tsx), which React 19
hoists into this head; a static tag would sit earlier in document order
and win over the hoisted one. The build-time prerender
(src/prerender/) folds each route's tags into its emitted HTML. -->
What actually breaks when you switch LLMs — EvalShift
evalshift diff · cli 0.11.0 docs pricing blog compare example ↗ github ↗ login ← all posts migration · Aug 20, 2026 · 7 min read What actually breaks when you switch LLMs
A model swap is a behavior change with no diff to review. What moves besides the answer text, and why the worst regressions read as correct.
— A model swap changes behavior — tools called, call order, arguments, output shape, refusals, latency, cost — and ships through a config file with no diff for anyone to review.
— The expensive regressions are invisible in the output: an agent that skips issue_refund and still says the refund was processed passes every text-based check.
— Migrations usually edit the prompt too, which puts two variables in one comparison — run new model with the old prompt first, then tune, so each result means something.
— Hand-written eval cases cover the interactions you already handle; the failures live in recorded traffic you would never have thought to write down.
— Significance is not importance: a 0.02 semantic delta can be real and irrelevant, one missing verify_payment call can be statistically invisible and serious. Write the policy before you see the numbers.
Changing the model behind an AI feature looks like the smallest change you will make all week. One
string moves from gpt-x to gemini-y . The new model is cheaper, or faster, or ahead on the
benchmark someone linked in Slack. You try a handful of prompts, the answers read fine, you ship.
The reason this keeps going wrong is that a model swap is not a config change. It is a behavior
change, delivered through a config file, with no diff for anyone to review.
## The change surface is bigger than the answer text
Swapping the model can move any of these independently, and most teams only look at the last one:
A migration can improve one row and wreck another. A model that answers just as well but issues one
extra tool call per turn is not a cost reduction. You will not learn that by reading answers.
## The worst regressions are invisible in the output
Take an agent that is supposed to do this:
lookup_order("A-339")
issue_refund("A-339", 29.99)
The new model does this instead:
lookup_order("A-339")
and replies:
Your refund has been processed.
Every text-based check passes. The sentence is fluent, on topic, and exactly what the old model said.
The refund did not happen. If you are scoring outputs, this regression is not merely hard to catch —
it is invisible by construction, because the output is correct and the behavior is not.
The same shape covers most of the expensive failures: the verify_payment call that stops firing,
the retry loop that starts, the confirmation step that moves after the write. What changed is the
trace. The prose stayed still.
## You are usually changing two things at once
Prompts are coupled to models. The system prompt in production has been tuned — often over months,
often by accretion — against one model's quirks. Point it at a different model and some of that
tuning becomes dead weight and some becomes actively harmful.
So the honest migration usually involves editing the prompt too, and now the comparison has two
independent variables in it. When quality moves, nobody can say whether the model did it or the
rewrite did.
The fix is boring: change one thing per run. Baseline old model with old prompt. Run new model with
old prompt — that is the model's effect, unflattering as it may be. Then tune the prompt for the new
model and run again, against the same frozen cases. Two comparisons, each interpretable, instead of
one that isn't.
## Hand-written test cases test the paths you already handle
Writing eval prompts by hand feels productive and produces a suite shaped like your mental model of
the product. That is the problem. The prompts you can think of are the interactions you already
understand well enough to have handled.
Real failures come from the inputs you would never have written down: the eleven-turn conversation
where a constraint set in turn three quietly expires, the message that arrives with half the context
missing, the turn right after a tool returned an error, the customer typing in a language your
template never anticipated. You cannot reconstruct those from memory. You have to record them.
Which gives a workflow, independent of what you use to run it:
real agent behavior
↓
capture representative cases
↓
freeze a golden suite
↓
baseline model vs candidate model, same inputs
↓
compare behavior, not just text
↓
ship or reject
The freezing step is the one people skip. If cases are still being edited while models are being
compared, two things are moving and the diff between them describes neither.
## An average is not a verdict
Model A scores 0.91, model B scores 0.89, and someone screenshots it into the migration thread. That
number cannot carry the decision. With a couple of dozen noisy samples, a two-point gap is well
inside what you would see running the same model twice.
Two things make it a real comparison. Run paired — every case against both models, then subtract per
case, so the fact that some cases are inherently harder cancels instead of drowning the signal. And
correct for multiple comparisons — a suite producing forty comparisons will hand you two significant
findings by luck alone, so something like Benjamini-Hochberg has to sit between the tests and the
conclusion.
But statistical significance is not product importance, and this is where the reasoning usually stops
one step early. A semantic-similarity drop of 0.02 can be real, reproducible, significant, and
irrelevant. One missing verify_payment call across two hundred cases is statistically nothing and
operationally a serious problem. Significance tells you the effect exists. It has no opinion about
whether you should care.
## Decide what matters before you see the numbers
Which is why the last artifact of a migration is a written policy — thresholds agreed while the
result is still unknown, so the verdict is read off rather than negotiated:
missing critical tool -> fail
invalid JSON above threshold -> fail
latency +10% -> warn
small semantic delta -> ignore
Write that after the run and the thresholds bend around the number you were hoping for. Everyone does
this; nobody means to.
A policy also gives you a fourth verdict worth having explicitly: inconclusive . Not "pass",
not "fail" — "this suite is too small to tell you". Teams that collapse that into a pass ship
regressions they had the evidence to catch, one underpowered comparison at a time.
## Disclosure, and the part that survives it
I build EvalShift , which does exactly this: capture real agent runs, freeze them into a golden
suite, run both models paired, score tool calls and arguments and structure alongside output quality,
and gate the pull request on a policy you wrote in advance. So take the tool mention as interested.
The argument underneath it is not:
If an LLM decides behavior in your system, then changing the model is a behavior change, and it
deserves the same treatment as a dependency bump that alters runtime semantics — frozen test cases,
a before-and-after run, and a decision rule written while you still have no stake in the answer.
Not a string edit in a config file, followed by hope.
+ How to test an LLM model migration before you ship it —
the same argument as a method, step by step.
+ Evaluating agent tool calls — what text evaluators structurally
cannot see, and the four evaluators that can.
+ Build a golden eval suite from production traffic —
turning recorded runs into the frozen suite this all depends on.
How to test an LLM model migration before you ship it
Prompt edits deserve the same gate as a model swap
How many eval cases do you need?
