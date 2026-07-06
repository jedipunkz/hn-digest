---
source: "https://syfi.cs.washington.edu/blog/2026-06-29-ekka/"
hn_url: "https://news.ycombinator.com/item?id=48810123"
title: "Ekka: Automated Diagnosis of Silent Errors in LLM Inference"
article_title: "Ekka: Automated Diagnosis of Silent Errors in LLM Inference | SyFI Lab"
author: "matt_d"
captured_at: "2026-07-06T21:22:25Z"
capture_tool: "hn-digest"
hn_id: 48810123
score: 2
comments: 0
posted_at: "2026-07-06T20:31:59Z"
tags:
  - hacker-news
  - translated
---

# Ekka: Automated Diagnosis of Silent Errors in LLM Inference

- HN: [48810123](https://news.ycombinator.com/item?id=48810123)
- Source: [syfi.cs.washington.edu](https://syfi.cs.washington.edu/blog/2026-06-29-ekka/)
- Score: 2
- Comments: 0
- Posted: 2026-07-06T20:31:59Z

## Translation

タイトル: Ekka: LLM 推論におけるサイレント エラーの自動診断
記事のタイトル: Ekka: LLM 推論におけるサイレント エラーの自動診断 | SyFIラボ

記事本文:
未来のインテリジェンスのためのシステム
ホーム
人々
出版物
会談
ニュース
ブログ
Ekka: LLM 推論におけるサイレント エラーの自動診断
Yile Gu、Zhen Zhang、Shaowei Zhu、Xinwei Fu、Jun Wu、Yida Wang、Baris Kasikci
論文を読む (arXiv) · ICML 2026 ページ
TL;DR: LLM サービス フレームワークは迅速に出荷されますが、サイレント エラー (エラーは発生せずに出力品質を低下させるバグ) が発生する場合もあります。 Ekka は、診断を差分デバッグとして扱うことでそれらを自動的に診断します。バグのあるフレームワークの中間実行状態を調整して信頼できる参照と比較し、分岐しているコンポーネントを特定します。 vLLM と SGLang からの 17 の実際のバグに関して、Ekka は ~$30/件で 80% pass@1 診断精度を達成し、開発者によって確認された 4 つの新しいバグをすでに発見しています。
ICML 2026 に参加しますか? 2026 年 7 月 8 日水曜日、午後 5:00 ～ 6:45 KST のポスター セッション 5 — ホール A、#2405 にお越しください。 AI インフラの信頼性についてお話したいと思います。
vLLM や SGLang などの最新の LLM サービス フレームワークは、ページ アテンション、基数アテンション、カスタム CUDA カーネル、継続的なバッチ処理、および新しい最適化の安定した流れなど、AI インフラ エンジニアリングの注目すべき部分です。しかし、その速度には代償が伴います。スタックが複雑で高速に動くほど、サイレント エラーが発生しやすくなります。これは、フレームワークが実行を続け、完全に整形式の応答を返し、静かに間違った応答を生成するバグです。
クラッシュやアサーションの失敗とは異なり、サイレント エラーではスタック トレースやエラー メッセージは表示されません。一見もっともらしいが、微妙に、またはひどく間違っている出力だけが表示されます。これらのバグの診断は、ベンチマーク スコアの低下などの症状と、奥深くに埋め込まれた誤って設定されたアテンション ウィンドウである可能性がある根本原因との間に大きなセマンティック ギャップがあるため、非常に難しいことで知られています。

モデルスタック。現在、1 つを追跡するのは手作業であり、場合によっては手動で構成を切り替え、テンソルをダンプし、リファレンス実装と比較するという数か月にわたる作業が必要になります。
Ekka は、まさにこの問題に対処する自動診断システムです。
これは、vLLM からの実際のサイレント エラー ( #17689 ) です。 Gemma 3 を vLLM で提供すると、HellaSwag ベンチマークの精度は HuggingFace と比較して 30% 近く低下しましたが、実行時エラーや警告は発生しませんでした。両方のフレームワークに同じ GSM8K プロンプトが与えられました。
ロリーは、配達用にサブウェイを 2 個 7.50 ドル、チップス 2 袋を 1.50 ドル、クッキー 2 個を 1.00 ドルで注文します。配達料 20% とチップ 5.00 ドルがかかります。彼女の注文の値段はいくらですか?
そして、まったく異なる答えが得られました。
HuggingFace (参考) → 最終回答: $29 ✓
サブスク: 2 × 7.50 ドル = 15 ドル。チップ: 2 × $1.50 = $3。クッキー: 2 × $1.00 = $2。
小計: 15 + 3 + 2 = 20 。配送料: 20 の 20% = 4。合計: $29
vLLM (バグあり) → 最終回答: $100 ✗
サブスク: 2 × 7.50 ドル = 15 ドル。チップ: 2 × $1.50 = $3。クッキー: 2 × $1.00 = $2。
食品の合計コストは 15 + 15 + 15 = 45 です。配送料: 45 × 0.20 = 9。合計: $100
生成はエラーなく終了しますが、計算は静かに外れます。開発者がこれを根本原因まで追跡するのに 3 か月かかりました。6 層中 5 層ではなく、すべての層にスライディング ウィンドウ アテンションを誤って強制した設定ミスです。出力だけからは、計算内のどこでエラーが発生したかを示す信号はほとんどありません。
サイレントエラーは実際にはどのようなものですか?
設計を基礎付けるために、vLLM および SGLang GitHub の問題から収集された 90 件の実際のサイレント エラーの実証研究を実施しました。私たちは、解決済みの問題 70 件を根本原因と診断ワークフローを理解するために使用し、未解決の問題は Ekka を評価するために残しました。 E の形をした 3 つの所見

kaさんのデザイン。
図 1. vLLM および SGLang におけるサイレント エラーの症状の分布 (左) と根本原因の分布 (右)。
1. 症状が原因を示していることはほとんどありません。これまでのところ最も一般的な症状は、43.8% での精度の低下です。表面的には有効なテキストでも、ベンチマークではスコアが悪くなります。残りは、意味不明なループや繰り返しループなどの偽の出力と、一貫性のない出力です。単一の劣化した答えだけでは、計算がどこで間違ったのかについてはほとんど何もわかりません。
2. 根本原因はスタック全体に広がります。 30.6% のフレームワーク実装、25.5% のモデル実装、24.5% のカーネル バックエンドがすべて寄与しており、問題の約 19.4% は純粋な数値精度、つまり論理的欠陥がまったくない浮動小数点の不安定性です。重要なのは、すべてのサイレント エラーの約半分はモデルとカーネル実装のモデル スタック内で発生するため、有用な診断ツールはモデルをブラック ボックスとして扱うことができないことです。モデルを認識する必要があります。
3. 開発者はすでに手動で差分デバッグを行っています。最も一般的な診断アクションは構成の切り替えと最小限の再現で、それぞれ問題の 60% 以上で発生し、ケースの約半数で別のフレームワークとの比較が行われ、手動によるアクティベーションのダンプや検査と組み合わせられることがよくあります。これはまさに Ekka が自動化するワークフローです。
図 2. サイレント エラーを解決するときに開発者が実行する診断アクション。参照フレームワークとの比較やアクティベーションの検査は一般的ですが、手作業であり、手間がかかります。
問題点は、2 つのフレームワーク間でテンソルを手動で比較するのは面倒なことです。最適化されたサービング エンジンは、参照モデルとはまったく異なる内部モジュール構造、名前付け、メモリ レイアウト、および融合操作を使用します。テンソルを直接比較することは通常、手動で大幅に調整しない限り不可能です

努力。 Ekka の仕事は、その調整を自動的に行うことです。
モデル、サンプル プロンプト、2 つのフレームワークの構成、バグのあるターゲットと信頼できる参照のどちらを指定すると、Ekka は相違の原因である可能性が最も高いコンポーネントのランク付けされたレポートを出力します。まずコンテキストを収集します。両方のコードベースをコード インデックスに解析し、各モデルのアーキテクチャをモデル ツリーに圧縮し、PyTorch フォワード フックを使用してバグを再現しながらアクティベーションと呼び出しシーケンスを記録します。次に、エージェント診断を 3 つのステップで実行します。
図 3. Ekka のアーキテクチャ。一番上の行は静的および動的コンテキストを収集します。一番下の行では、ツールセットと知識ベースに裏付けられた 3 ステップのエージェント診断が実行され、ランク付けされた診断レポートが出力されます。
1. コンポーネント マッピングは、2 つのフレームワーク間で意味的に同等のコンポーネントを見つけます。これは名前の一致よりも困難です。vLLM は Q/K/V を 1 つの QKVProjection に融合します。ここで HuggingFace は 3 つのモジュールを保持し、一致する名前でさえ分岐する可能性があります (HuggingFace の RotaryEmbedding.forward は cos/sin を計算しますが、vLLM は回転を適用します)。圧縮されたモデル ツリーから作業して、エージェントはノードをマッピングし、名前があいまいな場合は get_class_defining ツールを使用して実装を検査します。マッピングは増分的に構築され、確認されたペアを削除するバリデータによってチェックされ、すべてのコンポーネントをマッピングまたは説明する必要があります。
2. Activation Alignment では、出力の形状、dtype、レイアウト、または順序が異なる可能性があるため、マッピングされたコンポーネントを実際に比較できるようにします。たとえば、HuggingFace の個別の Q/K/V を vLLM の融合出力と調整することは、3 つのテンソルを連結してバッチ次元をドロップすることを意味します。 Ekka はこれをコード生成としてフレーム化します。エージェントは小さな Python 後処理関数を記述します。

定型テンプレートを処理する固定テンプレートです。ヘルパー ツール ( get_tensor_sum 、 infer_tensor_match ) はテンソルの照合に役立ち、既知の正しいモデルからの検証済みのアライメントのナレッジ ベースはコンテキスト内の例を提供し、バリデーターは生成されたコードをチェックします。
3. エラー分析は、数値上のノイズから実際のバグを分離します。モデルは BF16 や FP8 などの低精度で実行されるため、浮動小数点誤差がレイヤー間で蓄積され、固定 L2 しきい値が機能しません。 Ekka は、参照自体の低精度と完全精度のギャップによってターゲットの偏差を正規化する堅牢な誤差率を使用します。
ここで、X T と X R はターゲットとリファレンスからの低精度出力、Y R はリファレンスのフル精度 FP32 出力です。正しいコンポーネントは 1 付近に留まります。バグのあるものはそのはるか上にスパイクします。エラーは下流に伝播するため、Ekka は呼び出しシーケンスを追跡し、変化点分析を適用して、エラー率が持続的に上昇している最も初期のコンポーネント (アライメント ノイズによる 1 回限りのスパイクではなく持続的なシフト) にフラグを立ててから、ランク付けされた疑わしいものを推論とともにレポートします。
私たちは、vLLM と SGLang からの 17 件の実際のサイレント エラーのベンチマークを構築しました (12 件はバグ調査から、5 件はその後報告されました)。実行可能な再現スクリプトを備えた Docker コンテナーとしてパッケージ化されました。バグは、Llama 4、Gemma 3、Qwen 3、ERNIE 4.5、MambaMixer 2 を含む幅広いモデルに及び、根本原因は、RMSNorm 次元の不一致や不適切なテンソル並列シャーディングから、FlashInfer カーネルのバグや BF16 蓄積エラーにまで及びます。各ケースには、グラウンドトゥルースのバグコンポーネントがラベル付けされています。
ここでは、元のバグ レポート、コードベース、モデル パス、再現スクリプトを与えられた 2 つの最先端のソフトウェア エンジニアリング エージェント、OpenCode と Mini-SWE-Agent を比較します。すべてのシステムが使用する

Claude Sonnet 4.5 をバックエンドとして使用し、pass@$k$ メトリクスでスコア付けされた、疑わしい上位 3 つのコンポーネントを出力します。
図 4. エンドツーエンドの診断精度。 Ekka は、汎用コーディング エージェントよりもはるかに優れた 80% 合格 @1 および 88% 合格 @5 に達します。
Ekka は 80% pass@1 および 88% pass@5 に達し、Mini-SWE-Agent より 28%、OpenCode より 34% 向上しています。要点: 汎用コーディング エージェントには、セマンティック ギャップを乗り越えるのに必要なドメイン固有の足場 (モデルを意識したマッピング、テンソル アラインメント、高精度のエラー分析) が欠けており、その足場こそが Ekka の精度を高めるものです。
アブレーションは各部分が重要であることを確認します。コンポーネント マッピングの場合、検証ループとコード検査ツールを追加すると、マッピング精度が ~60% から ~90% に向上し、エンドツーエンド パス@1 が検証なしの 0.47、ツールなしの 0.67 から 0.84 まで向上します。アクティベーションの調整では、コード検証ループとヘルパー ツールおよびナレッジ ベースにより、調整精度が ~7% から ~88% に向上し、ツールとナレッジ ベースが削除されると pass@1 は 0.39 に低下します。診断精度も、1.5 から 2.5 のエラーしきい値全体にわたって安定しています。BF16 と FP8 の両方で、正しいコンポーネントは 1.5 よりかなり下に集中していますが、グラウンド トゥルースのバグのあるコンポーネントの範囲は 3.1 から 1093 までです。
安いですよ。プロンプト キャッシュなしで Claude Sonnet 4.5 を使用することで、Ekka はケースごとに平均約 400 万の入力トークンと約 304,000 の出力トークンを生成し、最悪の場合でも診断あたり 30 ドル未満になります。これにより、サービス提供フレームワークの既存のテスト ワークフローに組み込むことが現実的になります。
本当のテストは、Ekka が厳選されたベンチマークを超えて一般化できるかどうかです。 1 か月にわたって、Ekka は vLLM と SGLang で 4 つの新しいサイレント エラーを診断しました。これらはすべて開発者によって確認されました。
これらは、SGLang-16132 を含む、テキスト、オーディオ、画像モデルにまたがります

これは、Ekka が当初設計した自己回帰テキスト パスのはるか外側にある拡散モデルのバグであり、これは差分デバッグ アプローチの一般性を物語っています。
LLM サービス フレームワークは高速に動作し続けるため、サイレント エラーは避けられない副作用となります。 Ekka は、開発者がすでに依存している、数時間から数か月かかる手動の差分デバッグ ワークフローが自動化できることを示しています。LLM エージェントに、フレームワーク間でコンポーネントをマッピングし、中間アクティベーションを調整し、精度と伝播を意識したエラー メトリックで相違を推論するための適切な足場を与えます。また、サイレント エラーを 80% pass@1 で診断し、1 件あたり約 30 ドル、4 件の新しいバグが確認され、現在も増加中です。
Ekka は現在、PyTorch モジュール レベルでローカライズされています。より細かい関数レベルまたはラインレベルの粒度への移行、JAX/Flax などの非 PyTorch スタックのサポート、より小さなバックボーンとより強力なキャッシュによるコストの削減は、自然な次のステップです。コンポーネントのマッピング、アクティベーションの調整、エラー分析のコア パイプラインは変わりません。

## Original Extract

Systems for Future Intelligence
Home
People
Publications
Talks
News
Blog
Ekka: Automated Diagnosis of Silent Errors in LLM Inference
Yile Gu, Zhen Zhang, Shaowei Zhu, Xinwei Fu, Jun Wu, Yida Wang, Baris Kasikci
Read the paper (arXiv) · ICML 2026 page
TL;DR: LLM serving frameworks ship fast, and sometimes ship silent errors — bugs that degrade output quality without raising any error. Ekka diagnoses them automatically by treating diagnosis as differential debugging: it aligns and compares the intermediate execution states of a buggy framework against a trusted reference, and pinpoints the component where they diverge. On 17 real-world bugs from vLLM and SGLang, Ekka hits 80% pass@1 diagnosis accuracy at ~$30/case , and has already found 4 new bugs confirmed by developers.
Attending ICML 2026? Come find us at Poster Session 5 — Hall A, #2405 , on Wed, Jul 8, 2026 · 5:00–6:45 PM KST . We’d love to chat about AI infra reliability!
Modern LLM serving frameworks like vLLM and SGLang are remarkable pieces of AI infra engineering: paged attention, radix attention, custom CUDA kernels, continuous batching, and a steady stream of new optimizations. But that velocity has a cost. The more complex and fast-moving the stack, the easier it is to introduce a silent error : a bug where the framework keeps running, returns a perfectly well-formed response, and quietly produces the wrong answer.
Unlike a crash or an assertion failure, a silent error gives you no stack trace and no error message: just output that looks plausible but is subtly, or badly, wrong. Diagnosing these bugs is notoriously hard because of the huge semantic gap between the symptom, say a dropped benchmark score, and the root cause, which might be a misconfigured attention window buried deep in the model stack. Today, tracking one down is a manual, sometimes months-long slog of toggling configs, dumping tensors, and comparing against a reference implementation by hand.
Ekka is an automated diagnosis system that takes on exactly this problem.
Here is a real silent error from vLLM ( #17689 ). Serving Gemma 3 with vLLM, accuracy on the HellaSwag benchmark dropped by nearly 30% compared to HuggingFace — with no runtime error or warning. Both frameworks were given the same GSM8K prompt:
Rory orders 2 subs for $7.50 each, 2 bags of chips for $1.50 each and 2 cookies for $1.00 each for delivery. There’s a 20% delivery fee and a $5.00 tip. What will her order cost?
and produced very different answers:
HuggingFace (reference) → Final Answer: $29 ✓
Subs: 2 × $7.50 = $15. Chips: 2 × $1.50 = $3. Cookies: 2 × $1.00 = $2.
Subtotal: 15 + 3 + 2 = 20 . Delivery fee: 20% of 20 = 4. Total: $29
vLLM (buggy) → Final Answer: $100 ✗
Subs: 2 × $7.50 = $15. Chips: 2 × $1.50 = $3. Cookies: 2 × $1.00 = $2.
The total cost of the food is 15 + 15 + 15 = 45. Delivery fee: 45 × 0.20 = 9. Total: $100
The generation finishes without error, but the math silently goes off the rails. It took developers three months to trace this to the root cause: a misconfiguration that incorrectly enforced sliding-window attention on all layers instead of 5 out of every 6. From the output alone, there is almost no signal pointing to where in the computation that error originated.
What do silent errors actually look like?
To ground the design, we conducted an empirical study of 90 real-world silent errors collected from vLLM and SGLang GitHub issues. We used the 70 closed issues to understand root causes and diagnosis workflows, and kept the open ones to evaluate Ekka. Three findings shaped Ekka’s design.
Figure 1. Symptom distribution (left) and root-cause distribution (right) of silent errors in vLLM and SGLang.
1. The symptom rarely points to the cause. The most common symptom by far is accuracy regression at 43.8%: superficially valid text that scores worse on benchmarks. The rest are bogus output such as gibberish and repetition loops, plus inconsistent output. A single degraded answer tells you almost nothing about where the computation went wrong.
2. Root causes are spread across the entire stack. Framework implementation at 30.6%, model implementation at 25.5%, and kernel backends at 24.5% all contribute, and about 19.4% of issues are pure numerical precision — floating-point instability with no logical defect at all. Crucially, roughly half of all silent errors originate inside the model stack of model and kernel implementation, so a useful diagnosis tool cannot treat the model as a black box. It has to be model-aware.
3. Developers already do differential debugging — by hand. The most common diagnosis actions are configuration toggling and minimal reproduction, each in over 60% of issues, and comparing against another framework appears in about half of cases, often paired with manually dumping and inspecting activations. This is exactly the workflow Ekka automates.
Figure 2. Diagnosis actions developers perform when resolving silent errors. Comparing against a reference framework and inspecting activations are common — but manual and laborious.
The catch: manually comparing tensors across two frameworks is painful. Optimized serving engines use completely different internal module structures, naming, memory layouts, and fused operations than a reference model. A direct tensor comparison is usually impossible without significant manual alignment effort. Ekka’s job is to make that alignment automatic.
Given the model, an example prompt, the two frameworks’ configs, and which is the buggy target vs. the trusted reference , Ekka outputs a ranked report of the components most likely responsible for the divergence. It first collects context — it parses both codebases into a code index , compresses each model’s architecture into a Model Tree , and uses PyTorch forward hooks to record activations and the call sequence while reproducing the bug — then runs an agentic diagnosis in three steps.
Figure 3. Ekka’s architecture. The top row collects static and dynamic context; the bottom row runs the three-step agentic diagnosis, backed by a toolset and a knowledge base, and emits a ranked diagnosis report.
1. Component Mapping finds semantically equivalent components across the two frameworks. This is harder than name matching: vLLM fuses Q/K/V into one QKVProjection where HuggingFace keeps three modules, and even matching names can diverge (HuggingFace’s RotaryEmbedding.forward computes cos/sin, while vLLM’s applies the rotation). Working from the compressed Model Tree, the agent maps nodes and uses a get_class_definition tool to inspect implementations when names are ambiguous. Mappings are built incrementally and checked by a validator that removes confirmed pairs and requires every component to be either mapped or explained.
2. Activation Alignment makes mapped components actually comparable, since their outputs can differ in shape, dtype, layout, or ordering. Aligning HuggingFace’s separate Q/K/V with vLLM’s fused output, for instance, means concatenating three tensors and dropping a batch dimension. Ekka frames this as code generation: the agent writes small Python postprocessing functions, with a fixed template handling the boilerplate. Helper tools ( get_tensor_sum , infer_tensor_match ) help match tensors, a knowledge base of validated alignments from known-correct models provides in-context examples, and a validator checks the generated code.
3. Error Analysis separates real bugs from numerical noise. Because models run in low precision such as BF16 or FP8, floating-point error accumulates across layers and a fixed L2 threshold doesn’t work. Ekka uses a robust error ratio that normalizes the target’s deviation by the reference’s own low-vs-full-precision gap:
where X T and X R are the low-precision outputs from the target and reference, and Y R is the reference’s full-precision FP32 output. A correct component stays near 1; a buggy one spikes well above it. Since errors propagate downstream, Ekka walks the call sequence and applies change-point analysis to flag the earliest component whose error ratio stays persistently elevated — a sustained shift, not a one-off spike from alignment noise — then reports the ranked suspects with reasoning.
We built a benchmark of 17 real-world silent errors from vLLM and SGLang — 12 from our bug study and 5 reported afterward — packaged as Docker containers with runnable reproduction scripts. The bugs span a wide range of models, including Llama 4, Gemma 3, Qwen 3, ERNIE 4.5, and MambaMixer 2, and root causes, from RMSNorm dimension mismatches and incorrect tensor-parallel sharding to FlashInfer kernel bugs and BF16 accumulation errors. Each case is labeled with the ground-truth buggy component.
We compare against two state-of-the-art software-engineering agents — OpenCode and Mini-SWE-Agent — each given the original bug report, the codebase, the model path, and a reproduction script. All systems use Claude Sonnet 4.5 as the backend and output their top-3 suspected components, scored with the pass@$k$ metric.
Figure 4. End-to-end diagnosis accuracy. Ekka reaches 80% pass@1 and 88% pass@5, well ahead of general-purpose coding agents.
Ekka reaches 80% pass@1 and 88% pass@5 , a 28% pass@1 improvement over Mini-SWE-Agent and 34% over OpenCode. The takeaway: general-purpose coding agents lack the domain-specific scaffolding — model-aware mapping, tensor alignment, precision-robust error analysis — needed to navigate the semantic gap, and that scaffolding is exactly what drives Ekka’s accuracy.
Ablations confirm each piece matters. For component mapping, adding the validation loop and the code-inspection tool lifts mapping accuracy from ~60% to ~90%, and raises end-to-end pass@1 from 0.47 without validation and 0.67 without the tool up to 0.84. For activation alignment, the code-validation loop plus helper tools and knowledge base push alignment accuracy from ~7% to ~88%, and pass@1 drops to 0.39 once tools and the knowledge base are removed. Diagnosis accuracy is also stable across error thresholds from 1.5 to 2.5: correct components cluster well below 1.5 across both BF16 and FP8, while ground-truth buggy components range from 3.1 all the way up to 1093.
It’s cheap. Using Claude Sonnet 4.5 without prompt caching, Ekka averages ~4.0M input and ~304K output tokens per case — under $30 per diagnosis even in the worst case. That makes it practical to fold into a serving framework’s existing testing workflow.
The real test is whether Ekka generalizes beyond a curated benchmark. Over the course of a month, Ekka diagnosed 4 new silent errors in vLLM and SGLang — all confirmed by the developers :
These span text, audio, and image models — including SGLang-16132, a diffusion-model bug well outside the autoregressive text path Ekka was originally designed around — which speaks to the generality of the differential-debugging approach.
LLM serving frameworks will keep moving fast, and silent errors are an inevitable side effect. Ekka shows that the manual, hours-to-months differential-debugging workflow developers already rely on can be automated : give an LLM agent the right scaffolding to map components across frameworks, align their intermediate activations, and reason about divergence with a precision- and propagation-aware error metric, and it diagnoses silent errors at 80% pass@1 and roughly $30 per case, with 4 confirmed new bugs and counting.
Ekka localizes at the PyTorch module level today; pushing to finer function- or line-level granularity, supporting non-PyTorch stacks such as JAX/Flax, and driving cost down with smaller backbones and stronger caching are the natural next steps. The core pipeline of component mapping, activation alignment, and error analysis stays the same.
