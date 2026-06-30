---
source: "https://okaneland.com/study/do-ai-agents-work-yet/"
hn_url: "https://news.ycombinator.com/item?id=48732019"
title: "AI agents finish a third of the job, and the math says why"
article_title: "Do AI agents actually work yet? What the benchmarks show · Okane Land"
author: "ermantrout"
captured_at: "2026-06-30T13:34:35Z"
capture_tool: "hn-digest"
hn_id: 48732019
score: 2
comments: 1
posted_at: "2026-06-30T12:50:39Z"
tags:
  - hacker-news
  - translated
---

# AI agents finish a third of the job, and the math says why

- HN: [48732019](https://news.ycombinator.com/item?id=48732019)
- Source: [okaneland.com](https://okaneland.com/study/do-ai-agents-work-yet/)
- Score: 2
- Comments: 1
- Posted: 2026-06-30T12:50:39Z

## Translation

タイトル: AI エージェントは仕事の 3 分の 1 を完了しており、その理由は計算でわかります
記事タイトル: AIエージェントは実際にまだ機能していますか?ベンチマークが示すもの · オカネランド
説明: AI エージェントはまだ機能していますか?ベンチマーク (TheAgentCompany、CRMArena-Pro、tau-bench、METR) と、ステップごとのスコアが自律的なマルチステップ作業に加算されない理由の背後にある複合信頼性の計算、およびエージェントが現在支払う場所。

記事本文:
金OKANE LAND The Studio The Primer The Palette The Proof The Study The Ledger Community ↗ dark OKANE Land › The Study The Study · Explainer
AI エージェントが仕事の 3 分の 1 を完了します。その理由は計算でわかります
編集者 · 2026 年 6 月 29 日 · 11 分で読みました · 調査済み
デモは本物で、リーダーボードは急速に上昇しています。ただし、実際の複数ステップの作業では、最も優れたエージェントが作業の約 3 分の 1 を完了します。その理由は算術演算によるもので、信頼性がさらに低下します。ここでは、ベンチマークが示す内容、その背後にある半減期の計算、および現在エージェントが実際に支払う金額を示します。
市場で最高の AI エージェントを現実的なソフトウェア会社内で解放し、チャット ツール、プロジェクト ボード、コード リポジトリ、および 175 の通常のオフィス タスクを与えると、それらの約 30% を自力で完了します。それは古いモデルではなく、タイプミスでもありません。これは Gemini 2.5 Pro であり、カーネギーメロン大学の TheAgentCompany ベンチマークでは、入手可能なエージェントの中で最も強力なエージェントの 1 つであり、完全自律性のスコアは 30.3% で、途中まで進んだことを部分的に認めた場合にのみ 39% まで上昇しています。残りは失敗したか、中途半端に終わった。
それをデモの横に置いてください。エージェントが起動されるたびに、クリーンな実行が表示されます。エージェントは旅行の予約、経費の申請、バグの修正をすべて単独で行い、デモは本物です。こちらも1個のサンプルサイズです。ビジネスは 1 回の成功でうまくいくのではなく、1,000 回目でうまくいきます。「ステージ上で 1 回機能する」と「無人で毎回機能する」の間の距離がすべての問題です。 「AI エージェントはまだ機能するか」に対する簡単な答えは次のとおりです。狭い、短い、監視付きタスクについては、まったくそのとおりです。また、無制限の複数ステップの作業については、人々は自動化することを約束し続けていますが、信頼性は高くなく、まだ自動化されていません。その理由は、来月パッチが適用されるバグではなく算術演算です。
エージェントはキーを獲得します

狭い線の内側にあると、線を広げた瞬間に失われます。お金を稼ぐために仕事をさせている場合:
範囲を狭くし、タスクを短くしてください。ジョブが長くなるほど信頼性は崖から落ちます。これに関する以下のデータは残酷です。エージェントは、ツール 1 つ、明確に定義されたジョブ 1 つ、いくつかの手順ですでに料金を支払います。
人間を検問所に入れてください。すべてのトークンを監視するのではなく、お金がかかる、または元に戻せないいくつかの動きを承認します。エージェントが草案を作成し、あなたがコミットします。
失敗を安価に、そして回復可能にします。エージェントに指示するのは、サンドボックス、ドラフト、提案書であり、本番データベースやライブ送信では決してありません。主要なエージェントの惨事は、いずれも、罰金からは程遠い、取り返しのつかない行為でした。
再試行とチェッカーを追加します。最初のパスの動作を検証する 2 番目のパスにより、失われた信頼性の多くが取り戻されます。裸の単発自律性は、最も弱い設定です。
リーダーボードは買わないでください。最も驚くべきスコアは、展開するエージェントからではなく、実際には購入できないオーダーメイドのマルチモデル リグから得られます。
これ以降はすべて証拠であり、その理由を示す数学です。
緩むと、ほとんどの仕事が失敗します
この 30% は異常値ではありません。 Salesforce は、エージェントが実際の販売およびサービス作業、つまりエージェントに販売する業務をどのように処理するかを確認するために、独自のベンチマーク CRMArena-Pro を構築しました。シングルステップタスクでは、最良のモデルは約 58% をクリアしました。タスクが通常のやりとりになり、エージェントが不足している詳細を尋ねる必要が生じた瞬間、その割合は約 35% に低下しました。同じエージェントは、特に注意するように指示されない限り、機密データの引き渡しを拒否することはほとんどなく、注意するように指示すると完了がさらに遅れました。これはベンダー自身の主張に対するベンダー自身の証拠であり、これを無視するのは困難です。
より深刻な問題は平均ではなく、平均値です。

一貫性。 Sierra のタウベンチは、エージェントを顧客サービス席に座らせ、複数ステップのチャットを通じて会社のポリシーに従うように依頼します。テストした最良のエージェントは平均的な小売タスクを約 61% の確率で正しく実行しており、これは使用可能であると思われます。次に、研究者らはさらに鋭い質問をしました。8 人の顧客に対して 1 回ずつ、同じタスクを 8 回連続で正しく実行できるでしょうか。毎回獲得したシェアは 25% を下回りました。 3 回のうち 2 回働くエージェントは、優れたデモであり、質の悪い従業員です。3 番目の結果を引き出す顧客は、返金、チャージバック、または苦情であるためです。
健全なタスクごとのスコアが依然として自律型ワーカーに加算されず、単なる掛け算にすぎない理由は次のとおりです。エージェントは各ステップで 95% の信頼性があり、これはほとんどの管理者よりも優れているとします。これらのステップを 20 個連鎖させて 1 つのタスクにまとめます。実際には、「旅行の予約と経費の申告」のような作業が含まれます。タスク全体が正しく完了する確率は、0.95 の 20 乗、つまり約 36% です。 10 段階ではまだ 60% にすぎません。長いタスクはサブタスクの連鎖であり、どれか 1 つが失敗すると全体が失敗するため、信頼性はさらに悪化し、下方にまで悪化します。
オックスフォード大学のトビー・オード氏は、2025 年の論文「AI エージェントの成功率には半減期はあるのか?」でこの名前を付けました。彼は、ベンチマーク データから作業を行って、エージェントの成功がタスクの長さに応じて指数関数的に減衰することを発見しました。これは、あたかも各エージェントが 1 分間の作業 (半減期) ごとに失敗する一定のリスクを抱えているのと同じくらい明確です。付箋にメモしておく価値がある結果の 1 つは、要求する信頼性が「9」増えるごとに、信頼できるタスクの長さをおよそ 10 で割ることです。 1 時間の仕事を半分の時間で終わらせるのに十分信頼できるエージェントは、8 ～ 9 回の作業が必要になると、その仕事のほんの一部しか信頼できなくなります。

10のt。デモは 50% のケースです。ビジネスは 99% のケースに基づいて運営されますが、それは別の、はるかに小規模な仕事です。
研究グループMETRは崖の上に数字を載せた。これは、エージェントの「50% の時間軸」、つまりエージェントが約半分の時間で一人で完了するのにかかる時間によってスコア化されたタスクの長さを測定します。引用される見出しは、この地平線は数か月ごとに 2 倍になっているというもので、実際は次のとおりです。フロンティアは、人間と同等の作業が数分から数時間にまで約 2 年間で上昇しました。ここで重要なのは下の行です。 80% の地平線は 5 回中 4 回信頼できる長さですが、50% の地平線よりも約 5 倍短くなります。 METR のデータによると、モデルは人間が 4 分以内にかかるほぼすべてのタスクを成功させましたが、人間が約 4 時間以上かかるタスクの成功率は 10% 未満でした。見出しは「長い仕事に半分の時間で取り組む」。 「信頼性が必要な場合には、ほとんどない」という結果も同じです。
また、エージェントがすでに人々とマッチングしていることを示すリーダーボードを見つけることもできます。現実世界のアシスタント タスクのベンチマークである GAIA では、人間のスコアは 92% でしたが、プラグインを使用した GPT-4 のリリース時は 15% でした。現在、上位のエントリも約 92% を獲得しています。細字部分を読んでください。これらの最高スコアは、複数のモデルに足場がボルトで固定され、テスト用に調整されたカスタム マルチモデル リグであり、独自の作成者によって提出され、独自に複製されたものではありません。外部の研究所が実際に再現した最良のエージェントは 75% 近くにあり、足場のない裸のモデルはそれをはるかに下回ります。あなたにとって重要な数値は、研究チームが手作りのアンサンブルを使用して隠されたテストセットでヒットできる数値ではありません。それは、専門家の世話をすることなく、購入して自分の仕事を指差し、信頼できるものであり、その数値は l よりもかなり低いです。

イーダーボード。
公平な反論は、これらすべての数字が急速に上昇しているということです。 METR の範囲は 4 ～ 7 か月ごとに 2 倍になっています。今日の 30% は 2 年後の 30% ではありません。したがって、正しい姿勢は「エージェントはおもちゃだ」でも「金曜日までにエージェントが会社を経営する」でもありません。それは、現在崖がある場所に建設し、崖が後退していくのを見守り続けることです。
2026 年半ばに実際にお金を稼いでいるエージェントは、無人でビジネスを運営しているわけではありません。彼らは、限定された、限定された、監督された作業を行っており、それをうまくやっています。つまり、最初のバージョンの草稿を作成し、受信トレイのトリアージを行い、データを取得し、定型文を作成し、重要な少数の決定とミスを見つけるための安価な方法を担当する人がいます。それが増強であり、それは報われます。罠はすべてのピッチの次の文であり、そこでは制限付きのヘルパーが静かに無人の従業員になります。その判決は36%が住んでいる場所です。
これを購入するのではなく構築する場合は、もう 1 つ注意してください。再試行とステップごとにメーター上のトークンが増加するため、最も頻繁に失敗する長くループする複数ステップの実行は、処理に最もコストがかかるものでもあります。信頼性を損なう仕事はマージンを損なう仕事と同じであり、これはエージェントにどこを指示すべきか、どこに指示すべきでないかについての強力なヒントとなります。エージェントがコーディングで特別に料金を支払う場合、速度には独自の請求が伴います。
AIエージェントはまだ機能していますか?リード付きです、はい、そしてリードは急速に長くなります。紐を外して、「エージェント」という言葉が約束するはずだったオープンエンドのマルチステップ作業で、彼らはその約3分の1を終了しました、そして計算によれば、デモがきれいに実行されたからといってそれがひっくり返ることはありません。半分の時間働くエージェントはあなたが監督するツールであり、あなたが信頼する従業員ではありません。失敗した半分を構築すれば、残りの半分は維持されます

給料で。
送信する価値のあるものがある場合は 1 通のメール
受信箱で領収書を受け取ります。
決まったスケジュールや詰め物はありません。何かをテストしたり、数値を実行したり、時間をかける価値のあるツールを見つけたりすると、メールが届きます。
無料。ダブルオプトイン、ワンクリックで購読解除。
本番環境でエージェントを実行しますか?フォーラムのメモを比較 ↗
情報源と調査方法
徐ら。 / カーネギー メロン (2025)、TheAgentCompany: 現実世界の必然的なタスクに関する LLM エージェントのベンチマーク (arXiv v3)。 arxiv.org/abs/2412.14161
Salesforce AI Research (2025)、CRMArena-Pro: 多様なビジネス シナリオおよびインタラクションにわたる LLM エージェントの総合的な評価。 arxiv.org/abs/2505.18878
Yao、Shinn、Razavi、Narasimhan / Sierra (2024)、tau-bench: 現実世界のドメインにおけるツール、エージェント、ユーザーの相互作用のベンチマーク。 arxiv.org/abs/2406.12045
METR (2025)、長時間のタスクを完了する AI 能力の測定。 metr.org/blog/2025-03-19-measuring-ai-ability-to-complete-long-tasks (arXiv:2503.14499)
METR (2026)、タイムホライズン、更新されたリーダーボードとメモ。 metr.org/time-horizons
Toby Ord / Oxford Martin AI Governance Initiative (2025)、AI エージェントの成功率には半減期はありますか? arxiv.org/abs/2505.05115
Mialon、Fourrier、Swift、Wolf、LeCun & Scialom (2023)、GAIA: 一般的な AI アシスタントのベンチマーク。 arxiv.org/abs/2311.12983
プリンストン (2026)、ホリスティック エージェント リーダーボード (HAL)、GAIA の結果。 hal.cs.princeton.edu/gaia
AI コーディング: より速い MVP、より遅いレビュー、そして誰も言及しないセキュリティ法案
LLM をラップすると、20 ドル請求すると、損失が発生します。モデル所有者が最初に遭遇する価格設定の罠
OKANE LAND · AIでお金を稼ぐ人のために · © 2026
購読する · ここから始める · メディア キット · プライバシー · クッキー
Google アナリティクスを使用して、どの駒が着地し、どのフロップが発生したかを確認します。これにより Cookie が設定されます (
プライバシーに関する注意事項に詳細が記載されています)。同意する

許可するか、拒否するとオフのままになります。

## Original Extract

Do AI agents work yet? The benchmarks (TheAgentCompany, CRMArena-Pro, tau-bench, METR) and the compounding-reliability math behind why high per-step scores do not add up to autonomous multi-step work, plus where agents pay now.

金 OKANE LAND The Studio The Primer The Palette The Proof The Study The Ledger Community ↗ dark Okane Land › The Study The Study · Explainer
AI agents finish a third of the job, and the math says why
the editors · Jun 29, 2026 · 11 min read · researched
The demos are real and the leaderboards are climbing fast. Turned loose on real multi-step work, though, the best agents finish about a third of it, and the reason is arithmetic: reliability compounds downward. Here is what the benchmarks show, the half-life math behind it, and where agents actually pay today.
Turn the best AI agent on the market loose inside a realistic software company, give it the chat tool, the project board, the code repo, and 175 ordinary office tasks, and it finishes about 30% of them on its own. That is not an old model and it is not a typo. It is Gemini 2.5 Pro, one of the strongest agents available, on Carnegie Mellon’s TheAgentCompany benchmark, scoring 30.3% fully autonomous and creeping to 39% only when you hand out partial credit for getting part-way. The rest it botched or left half-done.
Hold that next to the demos. Every agent launch shows a clean run: the agent books the trip, files the expense, fixes the bug, all by itself, and the demo is real. It is also a sample size of one. A business does not run on one good run, it runs on the thousandth, and the distance between “worked once on stage” and “works every time, unattended” is the entire question. The short answer to “do AI agents work yet” is this: on narrow, short, supervised tasks, genuinely yes, and on the open-ended multi-step work people keep promising they will automate, not reliably, not yet, and the reason is arithmetic rather than a bug that gets patched next month.
Agents earn their keep inside tight lines and lose it the moment you widen them. If you are putting one to work to make money:
Keep the scope narrow and the task short. Reliability falls off a cliff as a job gets longer, and the data below is brutal on this. One tool, one well-defined job, a handful of steps, is where agents already pay.
Put a human at the checkpoints. Not watching every token, but approving the few moves that spend money or cannot be undone. The agent drafts, you commit.
Make failure cheap and reversible. Point an agent at a sandbox, a draft, a proposal, never a production database or a live send. The headline agent disasters were each one irreversible action away from fine.
Add a retry and a checker. A second pass that verifies the first one’s work buys back a lot of the lost reliability. Bare single-shot autonomy is the weakest setup there is.
Do not buy the leaderboard. The most eye-popping scores come from bespoke multi-model rigs you cannot actually purchase, not from the agent you would deploy.
Everything after this is the evidence, and the math that says why.
Turned loose, they fail most of the job
That 30% is not an outlier. Salesforce built its own benchmark, CRMArena-Pro , to see how agents handle real sales and service work, the exact thing it sells agents to do. On single-step tasks the best model cleared about 58%. The moment the task became a normal back-and-forth where the agent had to ask for a missing detail, that fell to roughly 35%. The same agents almost never refused to hand over confidential data unless they were specifically told to be careful, and telling them to be careful dragged completion down further. This is the vendor’s own evidence against its own pitch, which makes it hard to wave away.
The deeper problem is not the average, it is the consistency. Sierra’s tau-bench sits an agent in a customer-service seat and asks it to follow company policy through a multi-step chat. The best agent it tested got the average retail task right about 61% of the time, which sounds usable. Then the researchers asked a sharper question: can it get the same task right eight times in a row, once for each of eight customers? The share it nailed every single time dropped below 25% . An agent that works two times in three is a fine demo and a poor employee, because the customer who draws the third outcome is a refund, a chargeback, or a complaint.
Here is why a healthy per-task score still does not add up to an autonomous worker, and it is just multiplication. Say an agent is 95% reliable on each individual step, which is better than most manage. Chain 20 of those steps into one task, the kind of thing “book my travel and file the expenses” actually involves, and the odds it gets the whole task right are 0.95 to the twentieth power, or about 36% . At ten steps it is still only 60%. Reliability compounds, and it compounds downward, because a long task is a chain of subtasks where failing any one fails the whole thing.
Oxford’s Toby Ord gave this its name in a 2025 paper, Is there a half-life for the success rates of AI agents? Working from the benchmark data, he found agent success decays exponentially with task length, as cleanly as if each agent ran a fixed risk of failing per minute of work, a half-life. One consequence is worth keeping on a sticky note: every extra “nine” of reliability you demand divides the length of task you can trust by roughly ten. An agent reliable enough to finish an hour-long job half the time can only be trusted on a job a fraction as long once you need it to work eight or nine times out of ten. The demo is the 50% case. A business runs on the 99% case, and that is a different, much smaller job.
The research group METR put numbers on the cliff. It measures an agent’s “50% time horizon,” the length of task, scored by how long it takes a person, that an agent finishes on its own about half the time. The headline that gets quoted is that this horizon is doubling every few months , and it is: the frontier has climbed from minutes to a few hours of human-equivalent work in about two years. The line underneath is the one that matters here. The 80% horizon, the length you can trust four times in five, runs about five times shorter than the 50% one. In METR’s data, models nailed almost every task that takes a human under four minutes and succeeded less than 10% of the time on tasks that take a human more than about four hours. “Half the time on a long task” is the headline. “Almost never, when it has to be reliable” is the same finding.
You can also find a leaderboard that says agents already match people. On GAIA, a benchmark of real-world assistant tasks, humans score 92% , versus 15% for GPT-4 with plugins when it launched, and the top entries now claim around 92% too. Read the fine print. Those top scores are custom multi-model rigs, scaffolding bolted onto several models and tuned for the test, submitted by their own authors and not independently reproduced. The best agent an outside lab has actually reproduced sits near 75% , and a bare model with no scaffolding is far below that. The number that matters for you is not the one a research team can hit on a hidden test set with a hand-built ensemble. It is the one you can buy, point at your work, and trust without a specialist babysitting it, and that number is a good deal lower than the leaderboard.
The fair counterpoint, and it is fair, is that all of these numbers are moving up quickly. METR’s horizon is doubling on the order of every four to seven months. The 30% of today is not the 30% of two years from now. So the right posture is neither “agents are a toy” nor “agents will run the company by Friday.” It is to build for where the cliff sits today, and to keep watching it recede.
The agents earning real money in mid-2026 are not running businesses unattended. They are doing narrow, bounded, supervised work, and doing it well: drafting the first version, triaging the inbox, pulling the data, writing the boilerplate, with a person owning the few decisions that count and a cheap way to catch the misses. That is augmentation, and it pays. The trap is the next sentence in every pitch, where the bounded helper quietly becomes an unattended employee. That sentence is where the 36% lives.
One more thing if you are building this rather than buying it: the long, looping, multi-step runs that fail most often are also the most expensive ones to serve , because every retry and every step is more tokens on the meter. The work that breaks your reliability is the same work that breaks your margin, which is a strong hint about where to point an agent and where not to. Where agents do pay in coding specifically, the speed comes with its own bill .
Do AI agents work yet? On a leash, yes, and the leash is getting longer fast. Off the leash, on the open-ended multi-step work the word “agent” was supposed to promise, they finish about a third of it, and the math says that will not flip just because the demo ran clean. An agent that works half the time is a tool you supervise, not an employee you trust. Build for the half that fails, and you keep the half that pays.
One email, when there's something worth sending
Get the receipts in your inbox.
No fixed schedule, no filler. You get an email when we've tested something, run the numbers, or found a tool worth your time.
Free. Double opt-in, unsubscribe in one click.
Running an agent in production? Compare notes in the forum ↗
Sources & how we researched this
Xu et al. / Carnegie Mellon (2025), TheAgentCompany: Benchmarking LLM Agents on Consequential Real World Tasks (arXiv v3). arxiv.org/abs/2412.14161
Salesforce AI Research (2025), CRMArena-Pro: Holistic Assessment of LLM Agents Across Diverse Business Scenarios and Interactions. arxiv.org/abs/2505.18878
Yao, Shinn, Razavi & Narasimhan / Sierra (2024), tau-bench: A Benchmark for Tool-Agent-User Interaction in Real-World Domains. arxiv.org/abs/2406.12045
METR (2025), Measuring AI Ability to Complete Long Tasks. metr.org/blog/2025-03-19-measuring-ai-ability-to-complete-long-tasks (arXiv:2503.14499)
METR (2026), Time horizon, updated leaderboard and notes. metr.org/time-horizons
Toby Ord / Oxford Martin AI Governance Initiative (2025), Is there a half-life for the success rates of AI agents? arxiv.org/abs/2505.05115
Mialon, Fourrier, Swift, Wolf, LeCun & Scialom (2023), GAIA: a Benchmark for General AI Assistants. arxiv.org/abs/2311.12983
Princeton (2026), Holistic Agent Leaderboard (HAL), GAIA results. hal.cs.princeton.edu/gaia
AI coding: faster MVP, slower review, and the security bill nobody mentions
Wrap an LLM, charge $20, lose money: the pricing trap the model owners hit first
OKANE LAND · FOR PEOPLE MAKING MONEY WITH AI · © 2026
subscribe · start here · media kit · privacy · cookies
We use Google Analytics to see which pieces land and which flop. That sets cookies (the
privacy note has the details). Accept to allow it, or decline and it stays off.
