---
source: "https://tenureai.dev/writing/every-ai-memory-benchmark-has-an-asterisk/"
hn_url: "https://news.ycombinator.com/item?id=48664538"
title: "Every AI Memory Benchmark Has an Asterisk"
article_title: "Every AI memory benchmark has an asterisk | Tenure"
author: "freewilly25"
captured_at: "2026-06-24T19:31:52Z"
capture_tool: "hn-digest"
hn_id: 48664538
score: 2
comments: 0
posted_at: "2026-06-24T19:22:31Z"
tags:
  - hacker-news
  - translated
---

# Every AI Memory Benchmark Has an Asterisk

- HN: [48664538](https://news.ycombinator.com/item?id=48664538)
- Source: [tenureai.dev](https://tenureai.dev/writing/every-ai-memory-benchmark-has-an-asterisk/)
- Score: 2
- Comments: 0
- Posted: 2026-06-24T19:22:31Z

## Translation

タイトル: すべての AI メモリ ベンチマークにはアスタリスクが付いています
記事のタイトル: すべての AI メモリ ベンチマークにはアスタリスクが付いています |在職期間
説明: Mem0

記事本文:
在職期間
プラットフォーム ▼ コア機能
すべての AI メモリ ベンチマークにはアスタリスクが付いています
Mem0 は LongMemEval の 93.4% を最先端のものとして公開しています。誰かがクリーンなハーネスを通して製品を実行し、73.8% を獲得しました。 Mem0 の CTO はスレッドに姿を現しますが、ギャップを否定しません。代わりに、フィールド内のすべての数字にはアスタリスクが付いていると彼は言います。彼の言うことは正しいし、その認め方は理解する価値がある。
Mem0 は LongMemEval で 93.4% を発表しました。クリーンなサードパーティ製ハーネスでは 73.8% が生成され、同じメモリ システムと同じデータで 19.6 ポイントの差が生じました。
このギャップは、ハードコードされたデータセット固有の等価性ルール、審査員が「はいに傾く」ように指示されたこと、出力をサンプリングする誰にも見えない隠された思考連鎖推論、LoCoMo 審査員の一方向のスコア引き上げメカニズムに起因します。
Mem0のCTOが返答した。彼はそのギャップを否定しなかった。同氏は、すべてのメモリベンダーは独自のハーネスを調整しており、唯一の本当の修正は全員が実行する共有ハーネスであり、コストと遅延が精度とともに報告されていると述べた。
彼は正しい。現状では、全員がアスタリスク付きの数字を報告しています。問題は、この分野が共有された多次元の評価フレームワークに移行するのか、それとも各ベンダーが独自の判断を調整し続けるのかということです。
Tenure の PrecisionMemBench は、まさにこの目的のために構築されています。多次元測定 (精度、ノイズ分離、レイテンシー、信念の変動性) と、最終的な雰囲気スコアの代わりに証拠の経路をチェックする決定論的評価です。
同じシステム、同じデータで 19.6 ポイントの差
Mem0 は、LongMemEval で 93.4% を最先端の総合スコアとして公表しました。とき
サードパーティは、クリーンな評価ハーネス (gpt-5 アンサー、
イエスに傾く指導のないバイナリジャッジ、5シード平均）、彼らが得ることができる最高のもの
73.8%でした。 S

アメメモリーシステム。同じベンチマークデータ。 19.6点差。
この種のギャップには説明が必要です。サードパーティが Mem0 の公開ベンチマークを調査
4月の発表直前に彼らが出荷したコミットを利用したところ、
いくつかのこと。
これはどれも隠されていません。 SOTAの11日前、4月3日のコミットメッセージ
アナウンスには次のように書かれています。「evals からのプロンプトの同期: コンテキストチェック、ルール 14 (矛盾)、
矛盾する数字、個人化スキャン、裁判官のバイアスチェック、思考の連鎖
<judge_ Thinking> タグ、5 ステップの最終チェック。」彼らのエンジニアは「BIAS」という言葉を入力しました。
チェックインジャッジ」と「5 ステップの最終チェック」を git に追加します。
CTO が現れて、静かな部分を大声で言います
Mem0 の CTO である Deshraj 氏がスレッドで返答しました。彼はそのギャップを否定しなかった。彼はそうしなかった
調査結果は間違っていたと主張する。代わりに、彼は別の主張をしました。これらの選択は次のとおりです。
ベンチマーク自体の欠陥への対応。ベンチマークには隠れた前提条件が含まれています
そのため、完璧に記憶を呼び起こしたとしても、質問を解決することはできません。推理の痕跡と
等価性ルールは、それらの欠陥を補おうとする試みでした。
しかしその後、彼はもっと興味深いことを言いました。彼はこう言いました。
「はい。これらのほとんどは、片側が一般的なハーネスで、もう片側が調整されたハーネスです。
トークンの予算が異なり、遅延は誰も言及せず、半分はエージェントで、すべて潰されています
すべてを隠す単一の精度スコアに変換します。唯一の本当の修正は共有することです
コストと遅延が精度とともに報告されるため、誰もが実行する問題を活用できます。
それまでは、世の中にあるすべての数字には、私たちの数字にもアスタリスクが付いています。」
これはスレッド全体の中で最も正直な発言です。それはギャップを守るものではありません。
それは、フィールド全体が、誰もが自分のチューニングをするゲームをしていることを認めることです
裁判官が判断した結果、得られた数値は比較できません。の

最も著名な企業の CTO
メモリ会社は公の場で、自社のベンチマーク数値には次のような特徴があると述べています。
アスタリスク。
同氏はさらにこう続けた。「正直に言って、これが私たちが最も気にかけている部分だ。私たちは前進しようとしているところだ」
エージェントの記憶をフィールドとして転送します。これは、できる限りの正確性を報告することを意味します。
コスト、トークン、そして既存のベンチマークがどこで不足しているのかを率直に把握すること。
彼らの周りで静かにゲームをしています。」
全員の番号にはアスタリスクが付いています
CTOの説明は正確です。 AI メモリ ベンチマークでは、現時点でのあなたはどうなっているのか
ベンダー間の比較が同じになることはほとんどありません。
各ベンダーは、独自の評価プロンプト、判断指示、および同等性ルールを調整します。同じメモリ システムを 2 つの異なるハーネスで実行すると、スコアが 20 ポイント異なる場合があります。
あるシステムはクエリごとに 500 個のトークンを取得し、別のシステムは 5,000 個のトークンを取得する可能性があります。コンテキストが増えると正確さが容易になりますが、コストについては誰も報告しません。
一部のシステムは、複数のステップにわたって複数のモデルを呼び出すエージェント ループを実行します。他のものは単一の取得パスを実行します。精度の数値だけでは、どれがどれであるかは分かりません。
精度、リコール、ノイズ分離、矛盾検出、信念の可変性、セッションターンのレイテンシはすべて単一の精度パーセンテージにまとめられており、実際のデプロイメントにとって重要なあらゆる側面が隠蔽されています。
その結果、あらゆるベンダーが高い数値を主張できる市場が誕生します。
ベンダー独自の評価設定では、技術的には真実です。しかし誰も答えません
開発者が実際に抱いている質問: これを統合すると、エージェントの記憶はどうなるでしょうか
実際にどのように動作しますか?また、それにはどのような費用がかかりますか?
精度とともにコストと遅延をレポートする共有ハーネス
CTO の処方箋はまさに正しいです。唯一の本当の解決策は、全員でハーネスを共有することです
コストとレイテンシを伴う実行

精度とともに報告されます。精度はひとつもありません
スコア。開発者が実際に何をトレードオフしているかを確認できる多次元レポート。
これは、Tenure の PrecisionMemBench が解決するために構築された問題です。崩れないよ
すべてを 1 つの数字にまとめます。次の 4 つの次元にわたってレポートします。
これらの次元は学術的なものではありません。これらは、本番環境の中断内容に直接対応します。
エージェントは正しい事実だけでなく 3 つの無関係な事実も取得し、状況を混乱させます。
モデル。メモリが蓄積されると、エージェントの速度が毎週遅くなります。できないエージェント
チームがどの API バージョンを使用しているかについての信念を更新します。単一の精度スコアが非表示になります
これらの障害モードのすべてを確認します。
アスタリスクは会話です
Mem0 スレッドを 1 つのベンダーのベンチマークに関する話として扱うのは簡単でしょう。
しかし、それは要点を外しています。 CTO の応答が本当のシグナルです。の
この分野で最も著名なメモリ会社は、記録上、自社のメモリ会社は次のように述べています。
番号にはアスタリスクが付いており、前に進む唯一の方法はハーネスを共有することです。
コストと遅延が精度とともに報告されます。
それはスキャンダルではありません。これはフィールドが実際にどこにあるかを示すものです。みんなでチューニング
ベンチマークには強制的な評価フレームワークが同梱されていないため、独自のハーネスが必要になります。
市場では単純な数値が評価されるため、誰もが単一の精度スコアを報告します。
この数字が比較できないことは誰もが知っていますが、誰も最初の数字になりたいとは思っていません。
ランディング ページに掲載するのが難しい多次元レポートを公開します。
アスタリスクはスキャンダルではありません。アスタリスクは、それ自体についての真実を伝えるカテゴリです。
AI のメモリはインフラになりつつあります。インフラストラクチャは、ベンダー固有の回答プロンプト、隠れた仮定、および 1 つの精度スコアにまとめられた判断の寛容さでは評価できません。
フィー

ld には共有ハーネスが必要です。最終的にはそうではありません。今。
PrecisionMemBench は、システムが何を取得したのか、何を見逃したのか、どのようなノイズを返したのか、すべてのメモリ プロバイダーが回答しなければならない層から始まります。
LLM-as-judge がエージェント評価のデフォルトになりました
最終回答判定では、代理人が主張内容を知ることを許可されていたかどうか、いつ知ることができたのか、欠席が確認されたかどうかを確認することはできません。
ほとんどの AI 評価が Linear 営業メールの失敗を見逃してしまう理由
LLM の裁判官であれば、そのメールを高く評価したでしょう。本当の失敗は、システムが事実を証明せずに送信を決定した以前に発生しました。
メモリの精度、ノイズ分離、セッションターンレイテンシ、および信念の可変性を測定するための取得ベンチマーク。

## Original Extract

Mem0

tenure
Platform ▼ Core Capabilities
Every AI memory benchmark has an asterisk
Mem0 publishes 93.4% on LongMemEval as state-of-the-art. Someone runs their product through a clean harness and gets 73.8%. The CTO of Mem0 shows up in the thread and doesn't deny the gap. Instead, he says every number in the field comes with an asterisk. He's right, and that admission is worth understanding.
Mem0 announced 93.4% on LongMemEval. A clean third-party harness produced 73.8%, a 19.6-point gap on the same memory system and the same data.
The gap traces to hardcoded dataset-specific equivalence rules, a judge instructed to "lean toward yes," hidden chain-of-thought reasoning invisible to anyone sampling outputs, and a one-directional score-lift mechanism in their LoCoMo judge.
The CTO of Mem0 responded. He didn't deny the gap. He said every memory vendor tunes their own harness, and the only real fix is a shared harness everyone runs against, with cost and latency reported alongside accuracy.
He's right. The status quo is everyone reporting numbers with an asterisk. The question is whether the field moves toward shared, multi-dimensional evaluation frameworks or keeps letting each vendor tune their own judge.
Tenure's PrecisionMemBench is built for exactly this: multi-dimensional measurement (precision, noise isolation, latency, belief mutability) and deterministic evaluation that checks evidence paths instead of vibes-scoring final answers.
A 19.6-point difference on the same system, the same data
Mem0 published 93.4% on LongMemEval as their state-of-the-art overall score. When a
third party ran their hosted product through a clean evaluation harness (gpt-5 answerer,
binary judge with no lean-toward-yes instruction, 5-seed mean), the best they could get
was 73.8%. Same memory system. Same benchmark data. A 19.6-point gap.
That kind of gap demands an explanation. The third party dug into Mem0's public benchmark
harness at the commit they shipped right before their April announcement, and found
several things.
None of this is hidden. The commit message from April 3rd, eleven days before the SOTA
announcement, reads: "Sync prompts from evals: CONTEXT CHECK, Rule 14 (contradictions),
conflicting numbers, personalization scan, BIAS CHECK in judge, chain-of-thought
<judge_thinking> tags, 5-step FINAL CHECK." Their engineer typed the words "BIAS
CHECK in judge" and "5-step FINAL CHECK" into git.
The CTO shows up and says the quiet part out loud
Deshraj, the CTO of Mem0, responded in the thread. He didn't deny the gap. He didn't
claim the findings were wrong. Instead, he made a different argument: these choices were
responses to flaws in the benchmarks themselves. The benchmarks contain hidden assumptions
that make questions unsolvable even with perfect memory retrieval. The reasoning traces and
equivalence rules were attempts to compensate for those flaws.
But then he said something more interesting. He said this:
"Yep. Most of these are a generic harness on one side and a tuned one on the other,
different token budgets, latency nobody mentions, half of them agentic, all squashed
into a single accuracy score that hides all of it. The only real fix is a shared
harness everyone runs against, with cost and latency reported alongside accuracy.
Until then every number out there, ours too, comes with an asterisk."
That is the most honest statement in the entire thread. It's not a defense of the gap.
It's an admission that the entire field is playing a game where everyone tunes their own
judge, and the resulting numbers are not comparable. The CTO of one of the most prominent
memory companies is saying, in public, that his own company's benchmark numbers come with
an asterisk.
He followed up: "Honestly this is the part we care about most. We're trying to move
agentic memory forward as a field, and that means reporting everything we can, accuracy,
cost, tokens, and being upfront about where existing benchmarks fall short instead of
quietly gaming around them."
Everyone's numbers have an asterisk
The CTO's description is accurate. In AI memory benchmarking right now, what you're
comparing between vendors is almost never the same thing:
Each vendor tunes their own evaluation prompts, judge instructions, and equivalence rules. The same memory system run through two different harnesses can produce scores 20 points apart.
One system might retrieve 500 tokens per query while another retrieves 5,000. More context makes accuracy easier, but nobody reports the cost.
Some systems run agentic loops that call multiple models across multiple steps. Others do a single retrieval pass. The accuracy numbers don't tell you which is which.
Precision, recall, noise isolation, contradiction detection, belief mutability, session-turn latency, all collapsed into a single accuracy percentage that hides every dimension that matters for real deployment.
The result is a market where every vendor can claim a high number, and every number
is technically true under the vendor's own evaluation setup. But none of them answer
the question a developer actually has: if I integrate this, what will my agent's memory
actually behave like, and what will it cost?
A shared harness that reports cost and latency alongside accuracy
The CTO's prescription is exactly right: the only real fix is a shared harness everyone
runs against, with cost and latency reported alongside accuracy. Not a single accuracy
score. A multi-dimensional report that lets developers see what they're actually trading off.
This is the problem Tenure's PrecisionMemBench was built to address. It doesn't collapse
everything into one number. It reports across four dimensions:
These dimensions are not academic. They correspond directly to what breaks in production:
the agent that retrieves the right fact but also three irrelevant ones and confuses the
model. The agent that gets slower every week as memory accumulates. The agent that can't
update its belief about which API version the team is on. A single accuracy score hides
every one of these failure modes.
The asterisk is the conversation
It would be easy to treat the Mem0 thread as a story about one vendor's benchmark
practices, but that misses the point. The CTO's response is the real signal. The
most prominent memory company in the space is saying, on the record, that their own
numbers come with an asterisk, and that the only way forward is shared harnesses with
cost and latency reported alongside accuracy.
That's not a scandal. It's a statement of where the field actually is. Everyone tunes
their own harness because the benchmarks don't ship with enforced evaluation frameworks.
Everyone reports a single accuracy score because the market rewards simple numbers.
Everyone knows the numbers aren't comparable, but nobody wants to be the first to
publish a multi-dimensional report that's harder to put on a landing page.
The asterisk is not the scandal. The asterisk is the category telling the truth about itself.
AI memory is becoming infrastructure. Infrastructure cannot be evaluated with vendor-specific answer prompts, hidden assumptions, and judge leniency collapsed into one accuracy score.
The field needs a shared harness. Not eventually. Now.
PrecisionMemBench starts at the layer every memory provider has to answer for: what did the system retrieve, what did it miss, and what noise did it return?
LLM-as-judge became the default for agent evaluation
Final-answer judging cannot see whether the agent was allowed to know what it claimed, when it could have known it, or whether absence was checked.
Why most AI evals would miss the Linear sales email failure
An LLM judge would have scored that email highly. The real failure happened earlier, when the system decided to send without proving the facts.
A retrieval benchmark for measuring memory precision, noise isolation, session-turn latency, and belief mutability.
