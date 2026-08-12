---
source: "https://stencil.so/blog/the-harness-problem"
hn_url: "https://news.ycombinator.com/item?id=49268543"
title: "We improved 15 LLMs at coding in one afternoon. Only the harness changed"
article_title: "We improved 15 LLMs at coding in one afternoon. Only the harness changed. — Stencil"
author: "latchkey"
captured_at: "2026-08-12T07:08:40Z"
capture_tool: "hn-digest"
hn_id: 49268543
score: 1
comments: 0
posted_at: "2026-08-12T06:28:12Z"
tags:
  - hacker-news
  - translated
---

# We improved 15 LLMs at coding in one afternoon. Only the harness changed

- HN: [49268543](https://news.ycombinator.com/item?id=49268543)
- Source: [stencil.so](https://stencil.so/blog/the-harness-problem)
- Score: 1
- Comments: 0
- Posted: 2026-08-12T06:28:12Z

## Translation

タイトル: 午後のコーディングで 15 個の LLM を改善しました。ハーネスのみ変更
記事のタイトル: 午後のコーディングで 15 個の LLM を改善しました。ハーネスのみ変更しました。 — ステンシル
説明: 16 のモデル、3 つの編集ツール、1 つの変数: ハッシュライン、コンテンツ ハッシュ ライン アンカーに基づいて構築された編集フォーマット、14/16 モデルのビート パッチ — そして最も弱いモデルが最も多くの利益を得ます。

記事本文:
ある午後のコーディングで 15 個の LLM を改善しました。ハーネスのみ変更しました。 — ステンシルステンシルブログ
研究
リンク プレビュー 午後 1 回のコーディングで 15 個の LLM を改善しました。ハーネスのみ変更しました。
実際には編集ツールが変更されただけです。それでおしまい。
編集形式ごとのモデルごとの合格率 · 3 実行 × 180 タスク、それぞれ新しいセッション ·
14/16 モデルではハッシュラインがパッチを上回ります。 v2 リビジョンは 12/16 でさらに改善されます —
最大のゲイン GPT-5.1 Codex Mini、60.0% → 77.5%
現在の会話は、GPT-5.3 と Opus のどちらのモデルがコーディングに最も優れているかについてほぼ完全に話題になっています。ジェミニ vs 今週落ちたもの。この枠組みは、実際にはボトルネックの 1 つであるハーネスがはるかにありふれたものであるにもかかわらず、モデルを重要な唯一の変数として扱うため、ますます誤解を招きます。
これは、ユーザーの第一印象 (制御不能なスクロールか、バターのようにスムーズか) を捉えるだけでなく、すべての入力トークンのソースであり、その出力とワークスペースに加えられたすべての変更の間のインターフェイスでもあります。
なぜわざわざ悩むのですか？ Opus は優れたモデルかもしれませんが、今日に至るまで Claude Code はサブエージェントの出力から生の JSONL を漏洩し、数十万のトークンを無駄にしています。オープン ハーネスでは、これを修正するだけで済みます。サブエージェントは構造化データを出力するようになりました。
ツールのスキーマ、エラー メッセージ、状態管理、「モデルが何を変更すべきかを認識している」から「問題が解決されている」までのすべて。実際に最も多くの失敗が起こるのはここです。
モデルに依存しないため、モデルはパラメーターにすぎないため、優れたテストの場となります。本当の変数はハーネスであり、想像を絶するほど制御できます。
とにかく、昨日変更した 1 つの変数についてです。
私たちが構築したものを説明する前に、最先端の技術を理解する価値があります。
Codex は apply_patch を使用します: 文字列を受け取ります

これは基本的に OpenAI 風味の差分である入力として、ハーネスは構造化スキーマに依存する代わりに、この BLOB が厳密なルール セットに従うことだけを期待します。 OpenAI の人々は間違いなく賢いため、トークン選択プロセスは、JSON スキーマや必要なツール呼び出しなどの他の制約がどのように機能するかのように、GPT の Codex バリアントの LLM ゲートウェイでこの構造に適合するようにほぼ確実に偏っています。
しかし、まったく気づかずにこれを他のモデルに与えるのでしょうか?パッチの失敗は屋根を通り抜けます。私たちのベンチマークにおける Grok 4 のパッチ失敗率は 50.7% で、GLM-4.7 のパッチ失敗率は 46.2% でした。これらは悪いモデルではありません。単に言語を話せないだけです。
クロード コード (および他のほとんどのコード) は str_replace を使用します。つまり、正確な古いテキストを検索し、新しいテキストに置き換えます。考えるのはとても簡単です。ただし、モデルは空白やインデントを含むすべての文字を完全に再現する必要があります。複数の試合?拒否されました。 「置換する文字列がファイル内に見つかりません」エラーは非常に一般的であるため、独自の GitHub 問題メガスレッド (+27 件の他の問題) が存在します。必ずしも最適ではありません。 Gemini は基本的に同じことを行いますが、それに加えていくつかのあいまいな空白のマッチングも行います。
カーソルは別個のニューラル ネットワークをトレーニングしました。つまり、ドラフト編集を取得してファイルに正しくマージすることが全体の仕事である微調整された 70B モデルです。ハーネスの問題は非常に難しいため、最も資金豊富な AI 企業の 1 つが別のモデルを投入することを決定しましたが、それでも彼らは自身のブログ投稿で「ファイル全体を完全に書き直す方が、400 行未満のファイルの補助的な diff よりも優れたパフォーマンスを発揮する」と述べています。
Aider 自身のベンチマークによると、フォーマットの選択だけで GPT-4 Turbo は 26% から 59% に変動しましたが、GPT-3.5 は有効な差分を確実に生成できなかったため、同じフォーマットでのスコアは 19% のみでした。形式はモデルと同じくらい重要です。
JetBrains conf の Diff-XYZ ベンチマーク

体系的に修正しました。モデルやユースケース全体で支配的な単一の編集形式はありません。 EDIT-Bench では、現実的な編集タスクで 60% 以上の合格@1 を達成したモデルは 1 つだけであることがわかりました。
ご覧のとおり、単純な「どうやって状況を変えるか」という問題に対する「最善の解決策」については、本当の意味でのコンセンサスはありません。私たちの見解では、これらのツールはいずれも、膨大な量のコンテキストを無駄にし、完全な再現に依存することなく、変更したい行に対する安定した検証可能な識別子をモデルに提供するものではありません。それらはすべて、すでに見たコンテンツを再現するモデルに依存しています。それができないとき、そしてたいていはできないとき、ユーザーはモデルのせいにする。
ここで我慢してください。モデルがファイルを読み取るとき、または何かを grep するときに、すべての行に 2 ～ 3 文字のコンテンツ ハッシュのタグが付けられて戻ってきた場合はどうなるでしょうか。
1:a3 |関数 hello() {
2:f1 | 「世界」を返します。
3:0e |} モデルは編集時にこれらのタグを参照します — 「行 2:f1 を置換、範囲 1:a3 から 3:0e を置換、3:0e の後に挿入」。前回の読み取り以降にファイルが変更されている場合、ハッシュは (楽観的には) 一致せず、何かが破損する前に編集が拒否されます。
疑似ランダムのタグを思い出せる場合は、自分が何を編集しているのかを知っている可能性があります。そうすれば、モデルは、その変更を表現するための信頼できる「アンカー」を示すために、古いコンテンツを再現したり、禁じられた空白を使用したりする必要がなくなります。
私たちの主な関心は実際のパフォーマンスであったため、フィクスチャは次のように生成されます。
React コードベースからランダムなファイルを取得します。
バグとして枠付けされた突然変異を、その逆が期待できる編集によって導入します (例: 演算子の交換、ブール反転、オフバイワンエラー、オプションのチェーンの削除、識別子の名前変更)。
問題の説明を平易な英語で生成します。
平均的なタスクの説明は次のようになります。
# `useCommitFilteringAndNavigation.js` のバグを修正
警備員

条項（早期返却）は削除されました。
問題は「useCommitFilteringAndNavigation」関数にあります。
欠落しているガード句 (早期リターンのある if ステートメント) を復元します。
当然のことながら、モデルは必ずしもまったく同じファイルであるとは限らない独自の解決策を思いつく可能性があるため、ここで 100% の成功率を期待することはできませんが、バグは十分に機械的なものであるため、ほとんどの場合、修正はミューテーションを元に戻すことで済みます。
タスクごとに 3 回実行、実行ごとに 180 タスク。毎回新しいエージェント セッション、4 つのツール (読み取り、編集、書き込み)。一時的なワークスペースを与え、プロンプトを渡し、エージェントが停止したら、フォーマット前後の元のファイルと比較します。
16 のモデル、3 つの編集ツール、結果は明白です。パッチはほぼすべてのモデルで最悪の形式で、ほとんどのモデルでハッシュライン一致またはビート置換が行われ、最も弱いモデルが最大の利益をもたらします。 Grok Code Fast 1 は 6.7% から 68.3% となり、10 倍の改善となりました。これは、パッチが壊滅的に失敗し、モデルの実際のコーディング能力が機械的な編集の失敗の陰にほぼ完全に隠れていたためです。 MiniMax は 2 倍以上になりました。 Grok 4 Fast の出力トークンは、再試行ループでのトークンの書き込みが停止されたため、61% 減少しました。
Gemini の成功率の +8% の向上は、ほとんどのモデルのアップグレードで実現されるよりも大きく、トレーニング コンピューティングのコストはゼロです。少し実験してみます (ベンチマークに約 300 ドルかかりました)。
多くの場合、モデルはタスクの理解が不安定ではありません。それは自分自身を表現するのが不安定です。あなたは着陸装置のせいでパイロットを責めています。
Anthropic は最近、非常に人気のあるオープンソース コーディング エージェントである OpenCode を、Claude Code サブスクリプションを通じて Claude にアクセスすることをブロックしました。
「OpenCode がプライベート API をリバースエンジニアリングした」という Anthropic の立場は、一見すると公平です。インフラストラクチャとルール。しかし、アクションサインが何であるかを見てください

同じく:
ハーネスを作らないでください。当社のものを使用してください。
それは単に人類学だけではありません。この記事を書いているときに、Google は私のアカウントを Gemini から完全に禁止しました。
レート制限はありません。警告されていません。無効 。ベンチマークを実行するためです。これは、Gemini 3 Flash が斬新なテクニックで 78.3% を達成したことを示したものと同じもので、最高の試みを 5.0 ポイント上回るものでした。何のためなのかさえわかりません。
それが逆になる理由は次のとおりです。異なる編集形式により、出力トークンが約 20% 削減されながら、独自のモデルが 5 ～ 14 ポイント改善されることを示しました。それは脅迫ではありません。無償の研究開発です。
競合他社のモデルに対してハーネスの最適化を行うベンダーはありません。 Anthropic は Grok に合わせて調整しません。 xAI は Gemini 用に調整しません。 OpenAI はクロードに合わせて調整しません。しかし、貢献者はさまざまなモデルを使用し、個人的に遭遇した障害を修正するため、オープンソースのハーネスはそれらすべてに合わせて調整されます。
モデルはお堀です。ハーネスは橋です。橋が燃えるということは、わざわざ渡ろうとする人が減るということだけです。ハーネスを解決済み、または重要ではないものとして扱うのは非常に短絡的です。
すべてのコード、ベンチマーク、および実行ごとのレポート: omp

## Original Extract

Sixteen models, three edit tools, one variable: hashline, an edit format built on content-hash line anchors, beats patch for 14/16 models — and the weakest models gain the most.

We improved 15 LLMs at coding in one afternoon. Only the harness changed. — Stencil Stencil Blog
RESEARCH
Link preview We improved 15 LLMs at coding in one afternoon. Only the harness changed.
In fact only the edit tool changed. That's it.
Pass rate per model per edit format · 3 runs × 180 tasks, fresh session each ·
hashline beats patch in 14/16 models; the v2 revision improves further in 12/16 —
largest gain GPT-5.1 Codex Mini, 60.0% → 77.5%
The conversation right now is almost entirely about which model is best at coding, GPT-5.3 or Opus. Gemini vs whatever dropped this week. This framing is increasingly misleading because it treats the model as the only variable that matters, when in reality one of the bottlenecks is something much more mundane: the harness.
Not only is it where you capture the first impression of the user (is it uncontrollably scrolling, or smooth as butter?), it is also the source of every input token, and the interface between their output and every change made to your workspace.
Why bother, you ask? Opus may be a great model, but Claude Code to this day leaks raw JSONL from sub-agent outputs, wasting hundreds of thousands of tokens. In an open harness, we get to just fix that: subagents output structured data now.
Tool schemas, error messages, state management, everything between "the model knows what to change" and "the issue is resolved." This is where most failures happen in practice.
Being model agnostic, it is a great testing ground, as the model is but a parameter. The real variable is the harness, which you have unimaginable control over.
Anyhow — about that one variable we changed yesterday.
Before we explain what we built, it's worth understanding the state of the art.
Codex uses apply_patch : It takes a string as input, which is essentially an OpenAI-flavored diff, and instead of relying on a structured schema, the harness just expects this blob to follow a strict set of rules. Since OpenAI folks are without a doubt smart, the token selection process is almost certainly biased to fit this structure at the LLM gateway for the Codex variants of GPT, similar to how other constraints like JSON schemas or required tool calls work.
But give this to any other model, completely unaware of it? Patch failures go through the roof. Grok 4's patch failure rate in our benchmark was 50.7% , GLM-4.7's was 46.2% . These aren't bad models — they just don't speak the language.
Claude Code (and most others) use str_replace : find the exact old text, swap in the new text. Very simple to think about. But the model must reproduce every character perfectly, including whitespace and indentation. Multiple matches? Rejected. The "String to replace not found in file" error is so common it has its own GitHub issues megathread (+27 other issues). Not exactly optimal. Gemini does essentially the same thing plus some fuzzy whitespace matching.
Cursor trained a separate neural network : a fine-tuned 70B model whose entire job is to take a draft edit and merge it into the file correctly. The harness problem is so hard that one of the most well-funded AI companies decided to throw another model at it, and even then they mention in their own blog post that "fully rewriting the full file outperforms aider-like diffs for files under 400 lines."
Aider's own benchmarks show that format choice alone swung GPT-4 Turbo from 26% to 59%, but GPT-3.5 scored only 19% with the same format because it couldn't reliably produce valid diffs. The format matters as much as the model.
The Diff-XYZ benchmark from JetBrains confirmed it systematically: no single edit format dominates across models and use cases. EDIT-Bench found that only one model achieves over 60% pass@1 on realistic editing tasks.
As you can see, there is no real consensus on the "best solution" to the simple "how do you change things" problem. Our take: none of these tools give the model a stable, verifiable identifier for the lines it wants to change without wasting tremendous amounts of context and depending on perfect recall. They all rely on the model reproducing content it already saw. When it can't — and it often can't — the user blames the model.
Now bear with us here. What if, when the model reads a file, or greps for something, every line comes back tagged with a 2-3 character content hash:
1:a3 |function hello() {
2:f1 | return "world";
3:0e |} When the model edits, it references those tags — "replace line 2:f1 , replace range 1:a3 through 3:0e , insert after 3:0e ." If the file changed since the last read, the hashes (optimistically) won't match and the edit is rejected before anything gets corrupted.
If they can recall a pseudo-random tag, chances are, they know what they're editing. The model then wouldn't need to reproduce old content, or god forbid whitespace, to demonstrate a trusted "anchor" to express its changes off of.
Since our primary concern was real-world performance, the fixtures are generated as follows:
Take a random file from the React codebase.
Introduce mutations, framed as bugs, via an edit whose inverse we can expect (e.g. operator swaps, boolean flips, off-by-one errors, optional chains removed, identifiers renamed).
Generate a description of the issue in plain English.
An average task description looks something like this:
# Fix the bug in `useCommitFilteringAndNavigation.js`
A guard clause (early return) was removed.
The issue is in the `useCommitFilteringAndNavigation` function.
Restore the missing guard clause (if statement with early return).
Naturally, we don't expect 100% success rate here, since the model can come up with a unique solution that isn't necessarily the exact same file, but the bugs are mechanical enough that most of the time, the fix is our mutation being reverted.
3 runs per task, 180 tasks per run. Fresh agent session each time, four tools (read, edit, write). We simply give it a temporary workspace, pass the prompt, and once the agent stops, we compare against the original file before and after formatting.
Sixteen models, three edit tools, and the outcome is unambiguous: patch is the worst format for nearly every model, hashline matches or beats replace for most, and the weakest models gain the most. Grok Code Fast 1 went from 6.7% to 68.3%, a tenfold improvement, because patch was failing so catastrophically that the model's actual coding ability was almost completely hidden behind mechanical edit failures. MiniMax more than doubled. Grok 4 Fast's output tokens dropped 61% because it stopped burning tokens on retry loops.
+8% improvement in the success rate of Gemini is bigger than most model upgrades deliver, and it cost zero training compute. Just a little experimenting (and ~$300 spent benchmarking).
Often the model isn't flaky at understanding the task. It's flaky at expressing itself. You're blaming the pilot for the landing gear.
Anthropic recently blocked OpenCode , a massively popular open-source coding agent, from accessing Claude through Claude Code subscriptions.
Anthropic's position "OpenCode reverse-engineered a private API" is fair on its face. Their infrastructure, their rules. But look at what the action signals:
Don't build harnesses. Use ours.
It's not just Anthropic either. While writing this article, Google banned my account from Gemini entirely:
Not rate-limited. Not warned. Disabled . For running a benchmark — the same one that showed Gemini 3 Flash hitting 78.3% with a novel technique that beats their best attempt at it by 5.0 pp. I don't even know what for.
Here is why that is backwards. We just showed that a different edit format improves their own models by 5 to 14 points while cutting output tokens by ~20%. That's not a threat. It's free R&D.
No vendor will do harness optimization for competitors' models. Anthropic won't tune for Grok. xAI won't tune for Gemini. OpenAI won't tune for Claude. But an open-source harness tunes for all of them, because contributors use different models and fix the failures they personally encounter.
The model is the moat. The harness is the bridge. Burning bridges just means fewer people bother to cross. Treating harnesses as solved, or even inconsequential, is very short-sighted.
All code, benchmarks, and per-run reports: omp
