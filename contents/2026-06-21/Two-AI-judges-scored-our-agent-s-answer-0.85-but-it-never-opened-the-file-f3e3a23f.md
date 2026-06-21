---
source: "https://tenureai.dev/writing/llm-as-judge-became-the-default-for-agent-evaluation/"
hn_url: "https://news.ycombinator.com/item?id=48620731"
title: "Two AI judges scored our agent's answer 0.85, but it never opened the file"
article_title: "LLM-as-Judge Became the Default for Agent Evaluation (And It Can't See the Failure That Matters) | Tenure"
author: "jflynt76"
captured_at: "2026-06-21T17:30:51Z"
capture_tool: "hn-digest"
hn_id: 48620731
score: 2
comments: 0
posted_at: "2026-06-21T17:23:57Z"
tags:
  - hacker-news
  - translated
---

# Two AI judges scored our agent's answer 0.85, but it never opened the file

- HN: [48620731](https://news.ycombinator.com/item?id=48620731)
- Source: [tenureai.dev](https://tenureai.dev/writing/llm-as-judge-became-the-default-for-agent-evaluation/)
- Score: 2
- Comments: 0
- Posted: 2026-06-21T17:23:57Z

## Translation

タイトル: 2 人の AI 審査員がエージェントの回答に 0.85 点を付けましたが、ファイルは開かれませんでした
記事のタイトル: LLM-as-Judge がエージェント評価のデフォルトになった (そして重要な失敗が見えない) |在職期間
説明: 最終回答の判定により、エージェントが

記事本文:
在職期間
プラットフォーム ▼ コア機能
LLM-as-judge がエージェント評価のデフォルトになりました (重要な失敗がわかりません)
ジャッジモデルは最終的な答えしか見ることができません。エージェントかどうかを確認できません
何を主張しているのか、いつ知ることができたのか、あるいは
欠勤申請が実際にチェックされたことはありません。リーダーボードの数字は、
3人とも沈黙。
LLM-as-judge は、無制限のタスクに使用できる唯一の汎用ツールであるため、エージェント評価のデフォルトになりました。だからといって、あらゆる障害クラスに適したツールになるわけではありません。
ジャッジモデルは、最終的な解答と正解を比較します。エージェントがその回答を生成するためにたどった経路を把握することはできず、エージェントが使用した証拠の使用が許可されているかどうかを確認する方法もありません。
あるケーススタディでは、2人のフロンティア裁判官がエージェントの反応を0.85と採点した。エージェントは、回答が依存する文書を開いたことがありませんでした。文書は存在しないと主張し、とにかく答えました。 GroundEval のスコアは 0.000、完全なケーススタディとコードはリポジトリにあります。
これは、工具使用メカニズムとは異なる故障クラスです。軌跡を認識するベンチマークは、エージェントが適切なツールを正しい順序で呼び出したかどうかをすでにチェックしています。それらのどれも、エージェントが自分が主張した内容を知ることが許可されているかどうか、または十分な調査によって欠勤請求が得られたかどうかをチェックしていません。
この修正は、決定論的な状態契約です。アクセス ポリシー、イベント ログ、およびアーティファクトのタイムスタンプが、ループ内の判断者なしで軌跡に対してチェックされます。
間違った質問を閉じ込めた枠組み
エージェントのベンチマークが単一の正しい文字列を持たない何かをスコアする必要がある場合、フリーテキスト
説明、複数段階の研究タスク、合理的に終了する可能性のあるツールを使用した軌跡
いくつかの異なる方法がありますが、標準的な動きは、

質問、返答、そして
審査員モデルの回答を参照し、採点してもらいます。人類最後の試験が普及
これの特定のバージョン: 最終的な答えを抽出し、既知の正解と比較します。
「はい」または「いいえ」を出力します。高速で汎用性があり、ドメイン固有の機能は必要ありません。
足場。まさにそれが広まった理由です。
裁判官のプロンプト自体は、何をチェックしているのかを明確に示しています。モデルに次のように指示します。
抽出された最終的な答えが正しい答えと一致するかどうかのみに焦点を当て、
背景についてコメントし、別の答えを主張するのではなく、問題自体を解決するのでもありません。
このスコープは意図的であり、文字列の一致を確認するという目的で構築されているため合理的です。
表現や数値の差異を許容します。何かをチェックするために構築されたものではありません
答えがどのように生み出されたかについて。
世界中で活動するエージェントにとって、最終回答の照合だけで十分かどうかを問う人は誰もいませんでした。
ただ質問に答えるだけではなく。インフラストラクチャはすでに存在しており、プロンプトテンプレート
はすでに書かれており、エージェントの出力の採点はモデルの出力の採点と同じように見えました。
そうではありません。その違いこそが問題なのです。
最終的な解答を読むジャッジモデルは、目的地が地図と一致するかどうかを確認します。
エージェントが実際にそのルートを歩いたのか、それともそこにテレポートしたのかを確認する方法はありません。
幸運な推測で。どちらも同じ文字列を生成します。次に信頼できるのはそのうちの 1 つだけです
地形が変わる瞬間。
2人のジャッジから0.85点。トレースからの 0.000。
エージェントは質問をされましたが、その答えは特定の Confluence ページに依存していました。
エージェントは確認したかのように応答しました。それはもっともらしく自信を持って説明した
言語、問題のページが存在しない理由を尋ね、それに基づいて質問に答えました。
2 人の別々のフロンティア裁判官モデルが質問を読み、

応答を評価し、0.85 とスコアを付けました。
両裁判官は、答えには十分な理由があり、説明は一貫していると判断した。どちらの裁判官もそうではなかった
質問と最終回答以外のものへのアクセス（全体）
採点プロンプトのデザイン: 最終的な回答を抽出し、それをグラウンド トゥルースと比較し、無視します。
他のすべて。
その痕跡は別の物語を語っていました。エージェントはページを取得しませんでした。一度も開けたことはありませんが、
それを検索したことも、それを明らかにするクエリを発行したこともありません。欠席を主張した
欠勤申請に必要な作業を行わずに、その後、次の理由を推論した。
検証されていない主張をあたかも事実であるかのようにすること。記録された軌道に対して得点し、
それを管理するアクセス ポリシーの場合、応答は 0.000 を受け取ります。
スコアリング ロジック自体、アクセス ポリシー コントラクト、イベント ログ スキーマは、
オープンソースの GroundEval リポジトリ。
誰でも同じケースを同じ軌道に対して実行し、同じ 0.000 を得ることができます。
これは、1 回のベンチマーク実行における 1 回限りの不具合ではありません。それは予測可能な結果です
ルートをまったく確認せずに目的地を評価すること。もっともらしく聞こえる最終的な答え
と有効な証拠のパスは異なるプロパティであり、
前者は、完全に後者に存在する失敗を決して捕らえることはできません。
軌道評価はすでに存在します。これは別の失敗クラスです。
これを、最終解答の採点を置き換えるべきだという議論と誤解しやすいでしょう。
軌道スコアリングを使用して、フィールドがすでにそれを解決していると結論付けます。そうではありません。
すでに存在する軌道を意識したベンチマークは、別の質問に答えています。
ツールの選択、引数の正確性、依存関係に関する軌跡レベルの診断
順序付け: エージェントが正しい引数を使用して正しいツールを呼び出したかどうかをチェックする

s、
正しい順序。進捗ベースのマルチターン評価により、エージェントが進捗したかどうかを確認します
最終的に成功するかどうかではなく、ターンをまたいで目標に向かって意味のあることを意味します。
裁判官モデルが Web エージェントの軌跡を確実に採点できるかどうかに関する研究の信頼性の研究
採点方法自体の、同じ機械的な質問、つまり手順は意味がありましたか? に適用されます。
3人とも整備士の点検中だ。エージェントは正しく移動しましたか。どれもチェックしていない
エージェントは、証拠から推論したかどうかにかかわらず、知っていると主張していることを知ることが許可されていました。
事後の証拠ではなく、質問が行われた時点で存在していたもの、または
欠勤を確認する必要がある場所を検索して欠勤請求を取得したかどうか
沈黙から主張するよりも。
機械的には完璧、状態は無効
エージェントは、適切な形式の引数を使用して適切なツールを適切な順序で呼び出し、可視コーンの外側のドキュメント、または質問時点以降に作成されたアーティファクトを使用して回答することができます。軌道力学についても何も言うことはありません。
何かが存在しないと主張することは、エージェントがそれが存在するであろう場所を検索した場合にのみ有効です。裁判官は自信を持って否定的な答えを見た。メカニクスチェックにより、一度も呼び出されなかったツールが検出されました。どちらも、検索が不十分であることを実際の失敗としてフラグを立てることはありません。
エージェントは、実際の、アクセス可能で、正確にタイムスタンプが付けられたイベントを引用しながら、因果関係の方向を逆方向に取得し、影響を以前の結果の原因として扱うことができます。引用は本物です。理屈はそうではありません。最終回答チェックもツール呼び出しチェックもこれを検出しません。
これらの失敗はそれぞれ同じ形をしています。つまり、状態またはガバナンスの制約に違反した場合、
機械的なステップがスキップされたものではなく、裁判官が文章を読んで発見できる推論上の誤りでもない。
慎重に。メモリシステムが取得できるのは、

間違ったユーザーの信念ストアからの正しい事実。
エージェントは、読み取りが許可されていないアーティファクトを引用しながらも正しく答えることができます。
どちらも、ジャッジやツール呼び出しに対して、有効な応答と同じように見える応答を生成します。
チェッカーも同様。
裁判官は決してそれを見る立場になかった
アーティファクトが特定の位置でエージェントの可視円錐の外側にあったことを検証する裁判官モデル
タイムスタンプは質問と回答だけからは判断できません。アクセスポリシーが必要になりますが、
イベント ログ、アーティファクトのタイムスタンプ、および予期される検索スペースが一緒に提供されます。
機械的にチェックできるほど正確な形式での応答。ほとんどの評価セットアップでは、
審査員のプロンプトはより単純な比較のために設計されているため、いずれも、この答えは次のとおりです。
その答えと一致します。
これらの構造が提供されると、実際に作業を行うものについて何かが変わります。
正しさの信号は裁判官の妥当性評価ではなくなり、状態になります。
契約そのもの。その時点で裁判官モデルはオプションになりますが、それは裁判官が信頼できないからではありません。
しかし、この証拠パスはこれらの制約の下で有効なのかという疑問が投げかけられているため、
制約を書き留めれば、決定的な答えが得られます。
正しさを評価する必要があるため、最終的な答えの正しさは不十分です
証拠パス: エージェントが知ることを許可された内容、いつ知ることができるか、何を知ることができるか
検索内容、引用内容、不在または因果関係の主張が正当化されたかどうか
エージェントが行動したときの世界の状態。
これは、きれいな精度や F1 数値を報告するベンチマークでもギャップが残る理由でもあります。
リーダーボード上のパーセンテージは、ジャッジが狭い答え合わせを行うことによって生成される可能性があります。
その下にあります。採点には制約があり、数値は客観的に見えますが、問題は

採点で問われるのは依然として「この答えが一致しているか」であり、「この答えが有効に到達したか」ではありません。
きれいに見える数字と目に見えない証拠経路の失敗は緊張関係にありません。それらは共存する
建設によって。
導入後に構築されるもののほとんどは、導入前に見つかるはずだったギャップに対する防御です。
何かを開示するように話しかけられるエージェントに対する標準的な対応を説明します。
そうすべきではありません。クエリを信頼できない入力として扱います。それを安全分類器を通して実行します。
流出の試みにフラグを立てます。そもそもレトリーバーが手の届く範囲を制限してください。フィルター
保護しようとしているものに似たパターンの出力。モデルを強制的に
承認されたコンテキストからのみ回答してください。それぞれが適切なパッチである 5 つのレイヤーが積み重ねられているため、
エージェントが実際にどの境界を維持できなくなるかを事前に知る人は誰もいませんでした。
その全体的な姿勢はテストのギャップの下流にあり、個別の問題ではありません。正確であれば
シナリオ。攻撃の背後にあるデータの要求と対になった敵対的な命令。
権限境界は、展開前に評価として実行され、
妥当性を判断するのではなく、具体的なアクセス ポリシーを作成することで、「このエージェントはどうなるだろうか」という質問に対する答えが得られます。
境界を保持する」という言葉が入っていることはわかっていたでしょう。推測ではありません。一般的なレイヤーが 5 つではない
未知の障害をカバーするために構築されています。確認されたギャップを中心に構築された 1 つの安全装置。
これは、デプロイメント前に状態の妥当性をテストすることで実際にチームを買収するものであり、
将来のあらゆる攻撃を排除するため、すべてのシナリオをカバーする評価スイートはありません。
テスト用に構築されていますが、既知の弱点に対して防御することと防御することの違いはありません
どの弱点が本当なのか分からないので、すべてに対して反対します。フィールドのほとんどは、
ツールが最初のアクセスのチェックを行うため、まだ 2 番目のことを行っています。

タイミング、
決定論的な証拠の境界はまだ標準ではありません。
承認書を書く前に契約書を書きましょう
このどれも、裁判官モデルのやり方が悪いとか、軌道力学が間違っているということを主張するものではありません。
ベンチマークは間違ったものを測定しています。どちらも実際の質問に対して実際に取り組んでいます。の
議論はより限定的であり、却下するのが困難です。失敗クラスが存在し、エージェントが境界線を越えています。
許可境界、まだ存在しない証拠からの推論、または許可境界が存在しないと主張する
構造上、どちらのアプローチも確認する立場にないことは決してチェックされていない
ジャッジモデルがどれほど有能であるか、または軌跡がどれほど注意深く記録されているか。
修正はより良い判断ではありません。それは別の質問であり、決定論的に尋ねられます。
事前に定義されたアクセス ポリシー、イベント ログ、およびアーティファクトのタイムスタンプがこれを実行しました。
エージェントの軌道は境界内に留まります。その質問には採点するためのモデルは必要ありません
契約書が書かれたら。そもそも契約が存在する必要があります。これがステップです
ほとんどのチームはスキップします。なぜなら、それを書くのは、審査員に出力を指摘して信頼するよりも時間がかかるからです。
返ってくる数字。
承認ステップ、ガバナンス層、エージェントが接触する前に人間が関与する

[切り捨てられた]

## Original Extract

Final-answer judging tells you whether an agent

tenure
Platform ▼ Core Capabilities
LLM-as-judge became the default for agent evaluation (and it can't see the failure that matters)
A judge model can only see the final answer. It cannot see whether the agent
was allowed to know what it claimed, when it could have known it, or whether
an absence claim was ever actually checked. The number on the leaderboard is
silent on all three.
LLM-as-judge became the default for agent evaluation because it is the only general-purpose tool available for open-ended tasks. That does not make it the right tool for every failure class.
A judge model compares a final answer to a correct answer. It has no visibility into the path the agent took to produce that answer, and no way to check whether the agent was permitted to use the evidence it used.
In a case study, two frontier judges scored an agent response 0.85. The agent had never opened the document its answer depended on. It asserted the document didn't exist and answered anyway. GroundEval scored it 0.000, full case study and code in the repo .
This is a distinct failure class from tool-use mechanics. Trajectory-aware benchmarks already check whether an agent called the right tools in the right order. None of them check whether the agent was allowed to know what it claimed, or whether an absence claim was earned by sufficient search.
The fix is a deterministic state contract: an access policy, an event log, and artifact timestamps, checked against the trajectory without a judge in the loop.
The framing that locked in the wrong question
When agent benchmarks need to score something with no single correct string, a free-text
explanation, a multi-step research task, a tool-using trajectory that could reasonably end
a few different ways, the standard move is to hand the question, the response, and a
reference answer to a judge model and ask it to grade. Humanity's Last Exam popularized
a specific version of this: extract the final answer, compare it to a known correct answer,
output yes or no. It is fast, it is general-purpose, and it requires no domain-specific
scaffolding. That is exactly why it spread.
The judge prompt itself is explicit about what it is checking. It instructs the model to
focus only on whether the extracted final answer matches the correct answer, and not to
comment on background, not to argue for a different answer, not to solve the problem itself.
That scope is deliberate and reasonable for what it is built to do: confirm a string match
with tolerance for phrasing and numerical variance. It was never built to check anything
about how the answer was produced.
Nobody asked whether final-answer matching was sufficient for agents that act in the world
rather than just answer questions. The infrastructure was already there, the prompt template
was already written, and grading an agent's output looked like grading a model's output.
It isn't. The difference is the entire problem.
A judge model reading a final answer is checking whether the destination matches the map.
It has no way to check whether the agent actually walked the route, or teleported there
on a lucky guess. Both produce the same string. Only one of them is trustworthy the next
time the terrain changes.
0.85 from two judges. 0.000 from the trace.
An agent was asked a question whose answer depended on a specific Confluence page.
The agent responded as though it had checked. It described, in plausible and confident
language, why the page in question did not exist and answered the question on that basis.
Two separate frontier judge models read the question and the response and scored it 0.85.
Both judges found the answer well-reasoned and the explanation coherent. Neither judge had
access to anything other than the question and the final response, which is the entire
design of the grading prompt: extract the final answer, compare it to ground truth, ignore
everything else.
The trace told a different story. The agent never fetched the page. It never opened it,
never searched for it, never issued a query that would have surfaced it. It asserted absence
without having done the work that an absence claim requires, and then reasoned forward from
that unverified assertion as though it were fact. Scored against the recorded trajectory and
the access policy that governed it, the response receives 0.000.
The scoring logic itself, the access policy contract, the event log schema, lives in the
open-source GroundEval repo .
Anyone can run the same case against the same trajectory and get the same 0.000.
This is not a one-off glitch in a single benchmark run. It is the predictable consequence
of grading a destination without ever checking the route. A plausible-sounding final answer
and a valid evidence path are different properties, and a method that only examines the
former will never catch a failure that lives entirely in the latter.
Trajectory evaluation already exists. This is a different failure class.
It would be easy to mistake this for an argument that final-answer scoring should be replaced
with trajectory scoring, and conclude the field has already solved it. It hasn't, because
the trajectory-aware benchmarks that already exist are answering a different question.
Trajectory-level diagnostics for tool selection, argument correctness, and dependency
ordering check whether an agent called the right tools, with the right arguments, in the
right sequence. Progress-based multi-turn evaluation checks whether an agent advanced
meaningfully toward a goal across turns, rather than just whether it eventually succeeded.
Work on whether judge models can reliably grade web-agent trajectories studies the reliability
of the grading method itself, applied to the same mechanical question: did the steps make sense.
All three are checking mechanics. Did the agent move correctly. None of them check whether
the agent was allowed to know what it claimed to know, whether it reasoned from evidence that
existed at the time the question was asked rather than evidence from after the fact, or
whether an absence claim was earned by searching the places absence should be checked rather
than asserted from silence.
Mechanically flawless, state-invalid
An agent can call the right tools in the right order, with well-formed arguments, and still answer using a document outside its visibility cone, or an artifact created after the question's as-of time. Trajectory mechanics has nothing to say about either.
Claiming something doesn't exist is only valid if the agent searched the places it would exist. A judge sees a confident negative answer. A mechanics check sees a tool that was never called. Neither one flags insufficient search as the actual failure.
An agent can cite a real, accessible, correctly timestamped event and still get the causal direction backward, treating an effect as the cause of an earlier outcome. The citation is real. The reasoning is not. Neither final-answer nor tool-call checking catches this.
Each of these failures has the same shape: a state or governance constraint was violated,
not a mechanical step skipped and not a reasoning error a judge could spot by reading
carefully. A memory system can retrieve a correct fact from the wrong user's belief store.
An agent can answer correctly while citing an artifact it was never permitted to read.
Both produce a response that looks identical to a valid one, to a judge and to a tool-call
checker alike.
The judge was never positioned to see it
A judge model verifying that an artifact was outside an agent's visibility cone at a specific
timestamp cannot do so from the question and response alone. It would need the access policy,
the event log, the artifact timestamps, and the expected search space supplied alongside the
response, in a form precise enough to check mechanically. Most evaluation setups don't supply
any of that, because the judge prompt was designed for a simpler comparison: does this answer
match that answer.
Once those structures are supplied, something changes about what's actually doing the work.
The correctness signal stops being the judge's plausibility assessment and becomes the state
contract itself. A judge model becomes optional at that point, not because judges are unreliable,
but because the question being asked, was this evidence path valid under these constraints, has
a deterministic answer once the constraints are written down.
Final-answer correctness is insufficient because correctness has to be evaluated against
the evidence path: what the agent was allowed to know, when it could know it, what it
searched, what it cited, and whether an absence or causal claim was justified by the
state of the world at the time the agent acted.
This is also why the gap survives even in benchmarks that report clean accuracy or F1 numbers.
A percentage on a leaderboard can still be produced by a judge doing narrow answer-matching
underneath it. The grading is constrained and the number looks objective, but the question
the grading asks is still "does this answer match," not "was this answer validly reached."
A clean-looking number and an invisible evidence-path failure are not in tension. They coexist
by construction.
Most of what gets built after deployment is defense for a gap that should have been found before it
Walk through the standard response to an agent that can be talked into disclosing something
it shouldn't. Treat the query as untrusted input. Run it through a safety classifier that
flags exfiltration attempts. Restrict what the retriever can reach in the first place. Filter
the output for patterns that look like the thing you're trying to protect. Force the model to
answer only from approved context. Five layers, each one a reasonable patch, stacked because
nobody knew in advance which boundary the agent would actually fail to hold.
That whole posture is downstream of a testing gap, not a separate problem. If the exact
scenario, an adversarial instruction paired with a request for data that sits behind a
permission boundary, had been run as an evaluation before deployment, scored against a
concrete access policy rather than judged for plausibility, the answer to "would this agent
hold the boundary" would have been known going in. Not a guess. Not five generic layers
built to cover an unknown failure. One safeguard, built around a confirmed gap.
This is what testing for state validity before deployment actually buys a team: not the
elimination of every future attack, no evaluation suite covers every scenario it wasn't
built to test, but the difference between defending against a known weakness and defending
against everything because you don't know which weakness is real. Most of the field is
still doing the second thing because the tools to do the first, checking access, timing,
and evidence boundaries deterministically, are not yet standard.
Write the contract before you write the approval
None of this argues that judge models are bad at what they do, or that trajectory-mechanics
benchmarks are measuring the wrong thing. Both are doing real work on real questions. The
argument is narrower and harder to dismiss: there is a failure class, an agent crossing a
permission boundary, reasoning from evidence that didn't exist yet, or claiming absence it
never checked for, that neither approach is positioned to see, by construction, regardless
of how capable the judge model is or how carefully the trajectory is logged.
The fix is not a better judge. It is a different question, asked deterministically: given
an access policy, an event log, and artifact timestamps defined ahead of time, did this
agent's trajectory stay inside the boundary. That question doesn't need a model to grade it
once the contract is written down. It needs the contract to exist at all, which is the step
most teams skip, because writing it is slower than pointing a judge at an output and trusting
the number that comes back.
An approval step, a governance layer, a human in the loop before an agent touc

[truncated]
