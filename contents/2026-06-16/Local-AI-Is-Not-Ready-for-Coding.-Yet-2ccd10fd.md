---
source: "https://mmlac.com/blog/local-ai-not-ready-for-coding-yet/"
hn_url: "https://news.ycombinator.com/item?id=48556879"
title: "Local AI Is Not Ready for Coding. Yet?"
article_title: "Local AI Is Not Ready for Coding. Yet? - Methodos Mechanicus"
author: "speckx"
captured_at: "2026-06-16T16:37:41Z"
capture_tool: "hn-digest"
hn_id: 48556879
score: 2
comments: 0
posted_at: "2026-06-16T15:33:45Z"
tags:
  - hacker-news
  - translated
---

# Local AI Is Not Ready for Coding. Yet?

- HN: [48556879](https://news.ycombinator.com/item?id=48556879)
- Source: [mmlac.com](https://mmlac.com/blog/local-ai-not-ready-for-coding-yet/)
- Score: 2
- Comments: 0
- Posted: 2026-06-16T15:33:45Z

## Translation

タイトル: ローカル AI はコーディングの準備ができていません。まだ？
記事のタイトル: ローカル AI はコーディングの準備ができていません。まだ？ - メソドロス・メカニクス
説明: マルチエージェント ハーネス内で自律コーディング エージェントとしてローカル LLM (Devstral 24B、Qwen3.5 122B) を実行するフィールド レポート - 何が機能し、何が壊れたのか、そしてその背後にあるハードウェアの壁。

記事本文:
コンテンツにスキップ
-->
--> ブログ ハイライト プロジェクト 私について ブログ ハイライト ナビゲート
ローカル AI はコーディングの準備ができていません。まだ？
私は、最大容量の MacBook に適合する最も強力なモデルに、エージェント ハーネスの鍵を与えました。実際のツールを呼び出した後、事務手続きの中に埋もれてしまいました。
デヴストラル (24B): 熱心、有能、そして綴りができない
Qwen3.5 (122B): 本物のエージェント、そして本物の喪失
修正: システムをスタッドまで剥がす必要がありました
デヴストラル (24B): 熱心、有能、そして綴りができない
Qwen3.5 (122B): 本物のエージェント、そして本物の喪失
修正: システムをスタッドまで剥がす必要がありました
GasTown の投稿で、私は使い捨ての主張をしました。ワーカー エージェント (実際にコードを書くポールキャット) には最強のモデルは必要ありません。 Opus ではなく Sonnet で実行すると、API の請求書が自動車の支払いのように見えなくなります。
それをきっかけに、次のステップについて考えるようになりました。もし、polecat がタスクを選択して実装するスコープ指定されたワーカーにすぎない場合、API になぜお金を払う必要があるのでしょうか?自分のマシン上のモデルで実行してみませんか?トークンごとの請求やレート制限はなく、建物からデータが流出することもありません。 1 日に何十回も回転する労働者にとって、「無料でプライベート」というのは非常に大声で言います。
それでテストしてみました。正しくは、「ローカル モデルに fizzbuzz を書くよう依頼する」のではなく、ローカル モデルを実際の運用ハーネスにドロップして、クラウド エージェントとまったく同じようにジョブを実行するように指示します。
簡単に言えば、ラップトップに適合するモデルはそのジョブを実行できず、また、ジョブを実行できるモデルはあなたのラップトップに適合しないということです。長いバージョンの方が興味深いのは、ローカル モデルの能力が私が予想していたよりも高かったり、低かったりするためです (同じ文の場合もあります)。
私は公平な戦いを望んでいたので、クラウド エージェントが使用するものとまったく同じハーネスをローカル モデルに与えました。
マシン: 最大

完成した MacBook Pro — M5 Max、48 コア GPU、128 GB ユニファイド メモリ。これは、コンシューマ ハードウェアが保持できるモデルと同じくらいです。
提供: ホスト上の Ollama は、OpenAI 互換 API を介して公開され、オープンコード ハーネスによって駆動されるため、エージェントは実際のツール ( bash 、 read 、 write 、 edit 、 grep ) を利用できます。
2 人の出場者:
Devstral 24B — 専用のローカル コーディング モデル (第 8 四半期で最大 25 GB)。小型、高速、専用設計。
Qwen3.5 122B-a10b — 1,220 億パラメータの専門家混合 (約 10B アクティブ)、Q4 で約 81 GB のメモリ。このマシンに物理的にロードできる最も高性能なモデル。 「フロンティア モデルをローカルで実行する」と言われると、これがラップトップ上での意味のほぼ上限となります。
ハーネス: GasCity、以前に書いた GasTown セットアップの後継製品です。同じメンタル モデル — 計画する市長、実行するケナガイタチ、審査員、合併する製油所。ワークフローは、コンボイに束ねられたビーズ (タスク) として流れ、Polarcat はタスクを「割り当て」から「マージ」まで行うための式 (複数ステップのワークフロー) を実行します。
タスク自体は意図的に簡単なもので、「特定のテキストを 1 行含むファイルを作成する」というものでした。モデルがそれを実行できない場合は、他に何も問題はありません。それができれば、車輪がどこから外れるのかを知ることができます。
コントロール グループ ローカル モデルを判断する前に、Claude Sonnet ポールキャットを通じて同じタスクを実行しました。これはパイプライン全体を通過し、ファイルの作成、コミット、レビューアーへの引き渡し、 main へのマージ、ビードのクローズを行いました。ドラマはありません。それは重要です。それは、ハーネス、ディスパッチ、タスクがすべて健全であることを証明します。以下に続くものはすべてモデルの貢献であり、セットアップが壊れているわけではありません。
デヴストラル (24B): 熱心、有能、そして綴りができない #
エージェントのオンボーディング プロンプト全体 (約 25 KB の役割、ルール、プロトコル) を渡されました。Devstral は私が何かをしてくれました。

期待しないでください。それは自己紹介でした。
「私はソフトウェア エンジニアリングのタスクを支援するためにここにいます。まず、何がサポートが必要かをお知らせください。」
ツール呼び出しはゼロです。そこには、「あなたは自律的な労働者です。割り当てられたタスクを見つけて実行する方法は次のとおりです」という 15,000 個のトークンが読み取られ、正しい行動は礼儀正しく指示を待つことであると結論づけられました。フロンティア モデルは、同じプロンプトを開始号砲として扱います。デヴストラルはこれを会社のハンドブックとして扱いました。
しかし、ここに「ローカル AI の準備ができていない」ということをあまりにも愚かにしているひねりがあります。セレモニーを取り除き、直接明示的な指示を与えると、「このファイルを作成し、書き込みツールを使用して、 ls を実行してください」と、それがただ実行されました。ツールの呼び出し、ファイルの書き込み、検証。 4回に3回くらい。その能力は本物です。質問が具体的な場合にのみ表示されます。
そして、簡単なタスクで次のことを実行しました。
成果物の名前を変更したモデルには、 MINSTRAL.md という名前のファイルが必要でした。 Devstral は minstrel.md を作成し、私のスペルを黙って「修正」し、小文字にしました。そして、「Linux ではファイル名の大文字と小文字が区別されるため」、これで問題ないと完全な自信を持って説明しました。
その正当化は間違っているだけでなく、逆向きです。大文字と小文字が区別されるのはまさに、MINSTRAL.md と minstrel.md が 2 つの異なるファイルである理由です。モデルは、自分が犯していることに気づかなかったエラーを正当化するために、もっともらしく聞こえる文章を生成しました。これは、ローカル モデルの失敗モードを縮小したものです。1 つの要件を持つタスクに対して、自信があり、流暢で、微妙にオフになっています。
とにかく出力をチェックするグルー コードやスコープ編集の場合、24B コーディング モデルは本当に役立ちます。一行ずつ読む予定のないものについては、正しさのない自信は後で支払う税金です。
Qwen3.5 (122B): 本物のエージェント、そして本物のエージェント

紛失番号
122B モデルは別の動物であり、これが私が最も驚いた結果です。
Devstral をフリーズさせた同じ完全なオンボーディング プロンプトを渡された Qwen は、自らスタートを切りました。これは、work-discovery コマンドを独自に実行し、割り当てられたビーズを見つけて、それをアトミックに要求し、メールボックスをチェックして、ワークフロー定義を読み取りました。つまり、実際の 14 回の正しい形式のツール呼び出しが連続して行われました。 「工具が使えない」モデルではありません。それは、自分がシステム内の自律的なエージェントであることを理解し、システムの操作を開始しました。
その後、コードは 1 行も書かれませんでした。
それは残りのすべてのターンを洞窟探索に費やしました。護送船団を視察中。護送船団の子供たちを読んでいます。メタデータをダンプしています。 3 つの異なるビーズの依存関係を確認します。ワークフローを読み直します。 18 分間、非の打ちどころのない目的を持ったように見えるツール呼び出しが行われましたが、実際のタスク (ファイルに 1 行を書き込む) は決して行われませんでした。それは私たちの星系の地図作成の中で迷子になり、カバーすべき領域があることを忘れていました。
私はそれがコンテキスト ウィンドウの問題ではないことを確認しました。これは、切り捨てられることなく ~15,000 トークンのプロンプト全体を取り込み、探索することを選択しました。複雑さがモデルをオーバーフローさせることはありませんでした。それはそれを捕らえた。
本当の発見 私たちのハーネスは、モデルの知能にとってそれほど複雑ではありませんでした。クウェンは十分に知能が高いのです。モデルの実行機能には複雑すぎました。フロンティアモデルは、「私は小さなことをするためにここにいます」という言葉をアンカーとして持ち、儀式を進めます。 122B ローカル モデルはセレモニーをタスクにします。それがラインであり、私が見つけることを期待していたラインではありませんでした。
修正: システムをスタッドまで剥がす必要がありました #
複雑さが捕らえ物であるならば、明らかな実験はそれを取り除くことだった。そこで、ワーカーの「ローカル エディション」を構築しました。25 KB のオンボーディング プロンプトを約 40 行に削減し、

7 ステップのワークフロー (コンテキストのロード → レコード プール → ワークツリーのセットアップ → プリフライト → 実装 → セルフレビュー → 送信) を 2 つのステップに分割します。実行してから引き渡します。
Qwen は、スリムなセットアップでタスクに到達し、正確に正しい内容のファイルを作成しました。 「18 分、ファイルゼロ」から「完了」まで、単に指示を削除するだけで完了します。それは本当に希望が持てる結果です。
でも尻尾は長かったです。ファイルが存在する瞬間に最初に停止しました - コミットまたはハンドオフの前に勝利を宣言 - そのため、「完了」とは「ファイルが書き込まれた」ということではなく、コミットおよび送信されたことを意味するというプロンプトの叫び声を上げなければなりませんでした。次の実行では、実際に完全なシーケンスが実行されました…その後、私の側の設定バグ (テスト リポジトリにブランチ名の不一致があり、リモートがありませんでした) につまずきましたが、使い捨ての git リポジトリを即興で作成し、空白に押し込むことで対処しました。その一部はモデルではなく私のせいでした。しかし、より深い点は重要です。フロンティア エージェントは git エラーに気づき、回復したであろうということです。ローカル モデルは自信を持ってそれを突破しました。ファイル名の代わりに git init を使用しただけで、 minstrel.md と同じ障害モードでした。
教訓は「できない」ではありません。それは、エージェント プラットフォームについてあなたが気に入っているすべての抽象化レイヤーは、ローカル モデルが持っていない経営陣の予算で支払わなければならないレイヤーであるということです。プラットフォームをほとんど何もない状態に戻すことで、そこに到達することができます。これがプラットフォームの目的のほとんどです。
ここが「面白い」を「まだ」に変える部分です。
上記のものはすべて、5,000 ドルのラップトップで実行できる最も強力なモデルでしたが、それでもハーネスを自律的に駆動することはできませんでした。自然な反応は、「だから、より大きく、よりスマートなオープン モデルを使用してください」です。これは、DeepSeek-V3 (671B) のような真にフロンティア競争力のあるオープン ウェイトです。

r GLM-4.6 (~355B)。そしてそれができるのです！消費者向けハードウェアではそれはできません。
DeepSeek-V3 は約 400 GB を 4 ビットに量子化し、完全精度で最大 750 GB になります。適切に機能させるには、約 1 TB の GPU メモリが必要です。これは、実際には 8 倍の NVIDIA H200 ノード (1,152 GB の VRAM、快適に十分) を意味します。
購入するには: 8×H200 サーバー / DGX クラス ボックスは、CPU、ネットワーキング、NVMe、冷却を追加すると、最大 300,000 ～ 300,000 ～ 300,000 ～ 500,000+ で動作します。 GPU だけでもそれぞれ約 35 ～ 40,000 ドルかかります。
レンタル: 8 − GPU ノードの場合、1 時間あたり約 25 ～ 35 ∗ ∗ ∗ ∗ 8 − GPU ノードのオンデマンドの場合、1 時間あたり 25 ～ 35 をコールします**。または、常時接続のエージェント フリートが必要とする方法で暖かく保つ場合、8 − GP U ノード − デマンド − コール ∗ ∗ 月あたり 20,000 です。
つまり、自律コーディング ハーネスを実際に駆動できるモデルを実行するための請求額は、「毎月の高級車」と「住宅」の間くらいになります。それに反して、ホスト型フロンティア API (トークンごとに請求され、アイドルコストはゼロ、DRAM 不足を他の誰かが消費する) は、高価なオプションのようには見えなくなり、大人のオプションのように見え始めます。
Apple シリコンの脚注 はい、単一の 512 GB M3 Ultra Mac Studio に 671B モデルをロードできます。これは、約 10,000 ドルでそれを収納できる唯一の消費者向けボックスでした。キャッチは2つ。まず、世界的な DRAM 不足のさなか、Apple は 2026 年 3 月に 512 GB 層を廃止したため、現在は購入することさえできません。 2 番目に、次のことができた場合でも、1 秒あたり最大 6 トークンが生成され、単一の出力トークンを発行する前に 8,000 トークンのプロンプトを取り込むだけで最大 14 分間待機します。私たちのエージェントのオンボーディング プロンプトだけでも、最大 15,000 トークンでした。エージェント ループは、数十回の往復です。計算は終わりません。
これが論文全体を 1 つのフレームにまとめたものです。

机の上に収まるモデルではワークを運ぶことができず、ワークを運ぶことができるモデルにはデータセンターが必要です。中間点、つまり消費者のコストと消費者のレイテンシを伴うフロンティア機能は、現在存在しません。
これを実際に実行した後、私がたどり着いたのはここです。これは「ローカル AI はおもちゃである」ということではありません。それは私が見たものではないからです。
地元のモデルは、範囲を絞り、適切に構成された、単発の仕事を真に行うことができます。 Devstral はグルー コードを作成します。 Qwen は実際の正しいツール呼び出しを行い、自己起動します。あなたのユースケースが「とにかくレビューするプライベートなオートコンプリートと制限付き編集」である場合、ローカルですでに十分であり、プライバシーと限界費用ゼロの話は現実です。
ローカル モデルは、複雑なシステム内で自律エージェントとして機能する準備ができていません。彼らは、運用の儀式を進める際に目標にしっかりと固定し続けるための実行機能を欠いており、間違ったファイル名、間違った git リカバリ、もっともらしいが間違った正当化など、自信を持って失敗します。無人のループでは、自信を持って間違っているということは、高価な種類の間違いです。
ハーネスは途中でそれらに会う必要があります。単一の最大のロック解除は、より優れたモデルではなく、より単純なシステムでした。ローカル ワーカーが必要な場合は、意図的に愚かでフラットな、彼らにタスクを押し付けるワークフローを設計します。オーケストレーション層のエレガントなものはすべて、エグゼクティブ フュームで実行されるモデルにとっては摩擦です。
ハードウェア曲線が本当のゲートです。だからこそ「まだ」なのです。小型モデルは四半期ごとにスマートになり、ハーネスも簡素化できるようになり、いつか 100B クラスのモデルがあなたの膝の上に登場する日が来るでしょう。

[切り捨てられた]

## Original Extract

A field report on running local LLMs (Devstral 24B, Qwen3.5 122B) as autonomous coding agents inside a multi-agent harness — what worked, what broke, and the hardware wall behind it all.

Skip to content
-->
--> Blog Highlights Projects About Me Blog Highlights Navigate
Local AI Is Not Ready for Coding. Yet?
I gave the most powerful model that fits on a maxed-out MacBook the keys to our agent harness. It made real tool calls — then drowned in the paperwork.
Devstral (24B): Eager, Capable, and Can’t Spell
Qwen3.5 (122B): Genuinely Agentic — and Genuinely Lost
The Fix: We Had to Strip the System to the Studs
Devstral (24B): Eager, Capable, and Can’t Spell
Qwen3.5 (122B): Genuinely Agentic — and Genuinely Lost
The Fix: We Had to Strip the System to the Studs
In the GasTown post I made a throwaway claim: your worker agents — the polecats that actually write the code — don’t need your strongest model. Run them on Sonnet, not Opus, and your API bill stops looking like a car payment.
That got me thinking about the obvious next step. If a polecat is just a scoped worker that picks up a task and implements it, why pay for an API at all? Why not run it on a model on my own machine ? No per-token bill, no rate limits, no data leaving the building. For a fleet of workers that spin up dozens of times a day, “free and private” is a very loud pitch.
So I tested it. Properly — not “ask a local model to write fizzbuzz,” but drop a local model into our real production harness and tell it to go do a job , exactly like the cloud agents do.
The short version: the model that fits on your laptop can’t carry the job, and the model that can carry the job doesn’t fit on your laptop. The longer version is more interesting, because the local models were both more and less capable than I expected — sometimes in the same sentence.
I wanted a fair fight, so I gave the local models the exact same harness the cloud agents use.
Machine: a maxed-out MacBook Pro — M5 Max, 48-core GPU, 128 GB unified memory . This is as much model as consumer hardware will hold.
Serving: Ollama on the host, exposed over its OpenAI-compatible API, driven by the opencode harness so the agent gets real tools: bash , read , write , edit , grep .
The two contestants:
Devstral 24B — a dedicated local coding model (~25 GB at Q8). Small, fast, purpose-built.
Qwen3.5 122B-a10b — a 122-billion-parameter mixture-of-experts (~10B active), ~81 GB in memory at Q4. The most capable model I can physically load on this machine. When people say “run a frontier model locally,” this is roughly the ceiling of what that means on a laptop.
The harness: GasCity, the production successor to the GasTown setup I wrote about earlier. Same mental model — a Mayor that plans, polecats that implement, a Reviewer, a Refinery that merges. Work flows as beads (tasks) bundled into convoys , and a polecat runs a formula (a multi-step workflow) to take a task from “assigned” to “merged.”
The task itself was deliberately trivial: “Create a file with one line of specific text.” If a model can’t do that, nothing else matters. If it can, we learn where the wheels come off.
The control group Before judging the local models, I ran the identical task through a Claude Sonnet polecat. It sailed through the whole pipeline — created the file, committed it, handed off to the Reviewer, got merged to main , closed the bead. No drama. That matters: it proves the harness, the dispatch, and the task are all sound. Everything that follows is the model’s contribution, not a broken setup.
Devstral (24B): Eager, Capable, and Can’t Spell #
Handed our full agent onboarding prompt — about 25 KB of role, rules, and protocol — Devstral did something I didn’t expect. It introduced itself.
“I’m here to assist you with software engineering tasks. To get started, please let me know what you need help with.”
Zero tool calls. It read 15,000 tokens of “you are an autonomous worker, here is how to find and execute your assigned task,” and concluded that the correct move was to wait politely for instructions. A frontier model treats that same prompt as a starting gun; Devstral treated it as a company handbook.
But here’s the twist that makes “local AI isn’t ready” too glib: when I stripped the ceremony and gave it a direct, explicit instruction — “create this file, use your write tool, then run ls ” — it just did it. Tool call, file written, verified. About three times out of four. The capability is real; it only shows up when the ask is concrete.
And then, on the trivial task, it did this:
The model that renamed the deliverable I asked for a file named MINSTRAL.md . Devstral created minstrel.md — silently “corrected” my spelling and lowercased it — then explained, with total confidence, that this was fine “since filenames are case-sensitive on Linux.”
That justification is not just wrong, it’s backwards. Case-sensitivity is exactly why MINSTRAL.md and minstrel.md are two different files. The model produced a plausible-sounding sentence to rationalize an error it didn’t notice it was making. This is the local-model failure mode in miniature: confident, fluent, and subtly off — on a task with one requirement.
For glue code and scoped edits where you’re checking the output anyway, a 24B coding model is genuinely useful. For anything you’re not going to read line-by-line, that confidence-without-correctness is a tax you pay later.
Qwen3.5 (122B): Genuinely Agentic — and Genuinely Lost #
The 122B model is a different animal, and this is the result that surprised me most.
Handed the same full onboarding prompt that made Devstral freeze, Qwen self-started. It ran the work-discovery command on its own, found its assigned bead, claimed it atomically, checked its mailbox, and read the workflow definition — fourteen real, correctly-formed tool calls in sequence. This is not a model that “can’t use tools.” It understood it was an autonomous agent in a system and started operating the system.
Then it never wrote a single line of code.
It spent every remaining turn spelunking. Inspecting the convoy. Reading the convoy’s children. Dumping metadata. Checking dependencies on three different beads. Re-reading the workflow. Eighteen minutes of impeccable, purposeful-looking tool calls, and the actual task — write one line to a file — never happened. It got lost in the cartography of our system and forgot there was territory to cover.
I confirmed it wasn’t a context-window problem: it ingested the full ~15k-token prompt without truncation and chose to explore. The complexity didn’t overflow the model; it captured it.
The real finding Our harness wasn’t too complex for the model’s intelligence — Qwen is plenty intelligent. It was too complex for the model’s executive function. A frontier model holds “I am here to do ONE small thing” as an anchor while it navigates ceremony. The 122B local model let the ceremony become the task. That’s the line, and it’s not the line I expected to find.
The Fix: We Had to Strip the System to the Studs #
If complexity was the captor, the obvious experiment was to remove it. So I built a “local edition” of the worker: I cut the 25 KB onboarding prompt down to about 40 lines, and collapsed the seven-step workflow (load context → record pool → set up worktree → preflight → implement → self-review → submit) into two steps: do the thing , then hand it off.
Qwen, on the slim setup, reached the task and created the file with the exact right content. From “18 minutes, zero files” to “done” — purely by deleting instructions. That is a genuinely hopeful result.
But the tail was long. It first stopped the instant the file existed — declaring victory before committing or handing off — so I had to make the prompt scream that “done” means committed and submitted, not “file written.” On the next run it actually ran the full sequence… and then tripped over a configuration bug on my side (the test repo had a branch-name mismatch and no remote), which it handled by improvising a throwaway git repo and pushing into the void. Some of that was my fault, not the model’s. But the deeper point stands: a frontier agent would have noticed the git error and recovered. The local model barreled past it confidently — the same failure mode as minstrel.md , just with git init instead of a filename.
The lesson isn’t “it can’t be done.” It’s that every layer of abstraction you love about your agent platform is a layer the local model has to pay for in executive budget it doesn’t have. You can get there — by stripping the platform back down to almost nothing, which is most of what the platform was for.
Here’s the part that turns “interesting” into “not yet.”
Everything above was the most powerful model I can run on a $5,000 laptop , and it still couldn’t autonomously drive the harness. The natural response is “so use a bigger, smarter open model” — the genuinely frontier-competitive open weights like DeepSeek-V3 (671B) or GLM-4.6 (~355B). And you can! You just can’t do it on consumer hardware.
DeepSeek-V3 is roughly 400 GB quantized to 4-bit, ~750 GB at full precision. To serve it properly you want around 1 TB of GPU memory , which in practice means an 8× NVIDIA H200 node (1,152 GB of VRAM, comfortably enough).
To buy: an 8×H200 server / DGX-class box runs ~ 300 , 000 t o 300,000 to 300 , 000 t o 500,000+ once you add CPUs, networking, NVMe, and cooling. The GPUs alone are ~$35–40K each.
To rent: roughly 25 – 35 p e r h o u r ∗ ∗ f o r a n 8 − G P U n o d e o n − d e m a n d — c a l l i t ∗ ∗ 25–35 per hour** for an 8-GPU node on-demand — call it **~ 25–35 p er h o u r ∗ ∗ f or an 8 − GP U n o d eo n − d e man d — c a ll i t ∗ ∗ 20,000 a month if you keep it warm the way an always-on agent fleet needs.
So the bill to run a model that can actually drive an autonomous coding harness is somewhere between “a luxury car every month” and “a house.” Set against that, a hosted frontier API — billed per token, zero idle cost, someone else eating the DRAM shortage — stops looking like the expensive option and starts looking like the adult one.
The Apple-silicon footnote Yes, you can load a 671B model on a single 512 GB M3 Ultra Mac Studio — it was the one consumer box that could hold it, at around $10,000 . Two catches. First, Apple pulled the 512 GB tier in March 2026 amid the global DRAM squeeze, so you can’t even buy it now. Second, even when you could: ~6 tokens/second, and a ~14-minute wait just to ingest an 8,000-token prompt before it emits a single output token. Our agent’s onboarding prompt alone was ~15,000 tokens. An agentic loop is dozens of those round-trips. The math doesn’t close.
That’s the whole thesis in one frame: the model that fits on your desk can’t carry the work, and the model that can carry the work needs a data center. The middle ground — frontier capability at consumer cost and consumer latency — doesn’t exist today.
After actually running this, here’s where I land — and it’s not “local AI is a toy,” because that’s not what I saw.
Local models are genuinely capable at scoped, well-framed, single-shot work. Devstral writes glue code; Qwen makes real, correct tool calls and self-starts. If your use case is “private autocomplete and bounded edits I’m going to review anyway,” local is already good enough, and the privacy and zero-marginal-cost story is real.
Local models are not ready to be autonomous agents in a complex system. They lack the executive function to stay anchored to a goal while navigating operational ceremony, and they fail confidently — the wrong filename, the wrong git recovery, the plausible-but-false justification. In an unattended loop, confident-and-wrong is the expensive kind of wrong.
The harness has to meet them halfway. The single biggest unlock wasn’t a better model — it was a simpler system. If you want local workers, design a deliberately dumb, flat, push-the-task-at-them workflow. Everything elegant about your orchestration layer is friction to a model running on executive fumes.
The hardware curve is the real gate. This is why it’s “yet.” Small models are getting smarter every quarter, harnesses can be simplified, and one day the 100B-class model on your lap

[truncated]
