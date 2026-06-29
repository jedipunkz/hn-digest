---
source: "https://syfi.cs.washington.edu/blog/2026-06-25-tracelab/"
hn_url: "https://news.ycombinator.com/item?id=48722460"
title: "TraceLab: Characterizing Coding Agent Workloads for LLM Serving"
article_title: "TraceLab: Characterizing Coding Agent Workloads for LLM Serving | SyFI Lab"
author: "matt_d"
captured_at: "2026-06-29T17:55:41Z"
capture_tool: "hn-digest"
hn_id: 48722460
score: 1
comments: 0
posted_at: "2026-06-29T17:44:00Z"
tags:
  - hacker-news
  - translated
---

# TraceLab: Characterizing Coding Agent Workloads for LLM Serving

- HN: [48722460](https://news.ycombinator.com/item?id=48722460)
- Source: [syfi.cs.washington.edu](https://syfi.cs.washington.edu/blog/2026-06-25-tracelab/)
- Score: 1
- Comments: 0
- Posted: 2026-06-29T17:44:00Z

## Translation

タイトル: TraceLab: LLM サービスのコーディング エージェント ワークロードの特性評価
記事のタイトル: TraceLab: LLM サービスのコーディング エージェント ワークロードの特徴付け | SyFIラボ

記事本文:
未来のインテリジェンスのためのシステム
ホーム
人々
出版物
会談
ニュース
ブログ
TraceLab: LLM サービスのコーディング エージェント ワークロードの特徴付け
カン・ジュー、マシュー・ジェイコブ、チェンシー・マー、イー・パン、ステファニー・ワン、アルビンド・クリシュナムルシー、バリス・カシクチ
主要な AI ラボが独自のコーディング エージェント (Anthropic の Claude Code、OpenAI の Codex、Google の Gemini CLI) を出荷しているため、これらのエージェントに効率的にサービスを提供することがシステムの問題としてますます重要になっています。ターミナル ベンチや SWE ベンチなどの既存のモデル品質ベンチマークは、サービス システムのパフォーマンスのモデル化にはあまり適していません。ツール呼び出しは比較的少なく、単一の独立したタスクに焦点を当てます。
このギャップを埋めるために、当社グループの日常的なコーディング エージェントの使用状況から収集された現実世界のデータセットである SyFI コーディング トレースをリリースし、コーディング エージェント用のサービス システムの設計をガイドします。また、トレース収集および分析パイプラインである TraceLab をオープンソース化することで、パーソナライズされたトレースを簡単に生成してデータセットに戻すことができるようになります。トレースとコードは https://github.com/uw-syfi/TraceLab で、ライブ デモは https://tracelab.cs.washington.edu/ で入手できます。
トレース収集とトレースファクト
私たちは、SyFI ラボで研究開発を行っている Claude Code と Codex の日常的な使用からこれらのトレースを収集しています。パイプラインは 2 つのステージで実行されます。まず、トレース コレクターが生の詳細なセッション ログを読み取り、キー フィールドを抽出し、周囲のロギング スキャフォールディングを破棄します。次に、トレース サニタイザーがツール呼び出しの入力や出力などの機密コンテンツを削除し、ユーザー名を削除し、セッション ID をマングリングすることで残りを匿名化します。結果として得られる匿名化されたトレースは、合計で約 4,300 セッションと 550 億トークンの規模で、サービング エンジンや分散サービング フリートにとっても十分な大きさです。

定常状態に達します。
私たちは主な観察を要約し、エージェント コーディングのユースケースをより適切にサポートするための将来のシステム研究への提案を提供します。
1 - 自律的な複数ステップの会話
コーディングエージェントは一発で答えてくれるものではありません。ユーザーのリクエストごとに、エージェントはモデル生成とツール呼び出しの自律ループを実行してから、最終的な回答を返します。私たちのトレースでは、各リクエストは平均 8.8 回の自主的なステップ (LLM ツール サイクル) を実行し、最終的な応答を返すまでに 10.8 回のツール呼び出しを発行します。その結果、すべての LLM ラウンドの 88% が人間ではなくツールの結果に応答します。
この自律性内で費やされる時間は、ループ全体に不均等に配分されます。 1 回の LLM 生成の平均時間は約 13 秒、各ツールの呼び出し時間は約 18 秒です。リクエストあたりのエンドツーエンドの応答時間はテールが大きく、中央値は約 38 秒ですが、平均は約 4 分で、p99 はほぼ 44 分です。ただし、最大のギャップはリクエスト間です。ユーザーは出力を読み、考え、次のリクエストを入力するのに時間を費やします。これには、中央値が約 1.4 分であるにもかかわらず、平均で 46.7 分かかります。
潜在的な研究の方向性。
合計数をほぼ一定に保ちながら、ラウンドごとに発行されるツール呼び出しの数を増やすことで、ステップ数が減り、ツールの並列実行の機会が増加します。
ユーザーが次のターンを開始しようとしているという早期の信号 (ユーザーがターミナル ウィンドウを再度アクティブにするか、入力を開始する) を使用して、ターン間の長いアイドル時間中に会話履歴をプリフェッチし、再入力します。
2- 長いコンテキスト、短い出力の生成
各ラウンドには非常に大きなプロンプトが含まれますが、生成される出力は比較的少量です。トレース全体で、モデルは 52.56 B のキャッシュされた入力トークンを読み取り、2.34 B の新しい入力トークンを事前に入力しますが、生成される出力トークンはわずか 1 億 8,690 万個であり、入力の数が出力 b を上回ります。

y 294× 。典型的なラウンドは 32k ～ 256k のトークン プレフィックスに配置され、わずか数百から数千のトークンを追加し、数百個をデコードします。ただし、新しい入力トークンの末尾は 128K を超える可能性があります。
速度に関しては、正規化されたデコード速度の中央値は全体で約 40.7 tok/s (Claude 46.8、Codex 33.9) であり、Codex の純粋なデコード速度は約 61.3 tok/s に達します。 Codex の各ステップの TTFT (約 3.1 秒) は、ラウンドの生成時間の約 25% です。
潜在的な研究の方向性。
短い増分プレフィルは、その独特のパフォーマンス特性により、長いプレフィルとは異なる方法で処理する必要があります。
エンドツーエンド時間を短縮するには、より適切なリクエスト ルーティングと自動スケーリング ポリシーを通じて、ツールの結果から次の世代に返すステップごとの TTFT を削減する必要があります。
3- 大量のツールを使用する、ロングテールの実行
自律的な LLM-ツール-LLM ループは、ほぼ完全に小規模なツール セット上で実行され、それらのツール呼び出しはシェルによって支配されます。 433,000 回のツール呼び出しのうち、76% はシェル/コマンドの実行 (ビルド、テスト、 git 、および同様のコマンドの実行) であり、ファイル編集 (11%) とファイルの読み取り/検索 (9%) が続きます。残りは、計画、サブエージェント、および Web ルックアップ コールで構成されます。 Claude は Codex (31 ツール) よりもはるかに幅広いツール語彙 (54 ツール) を利用していますが、どちらもそのボリュームの大部分が同じ少数のツール (主にシェル コマンド、ファイル編集、ファイルの読み取り/検索) に集中しています。
ツールのレイテンシは、ツールの種類全体と個々の種類の両方で大きくなり、Bash のようなツールでは 4 桁以上にも及びます。不均衡は顕著です。1 秒未満の呼び出しは全呼び出しの約 61% ですが、ツール時間の合計の約 1% にすぎません。一方、1 分を超える呼び出しの約 4% がその約 85% を消費します。
潜在的な研究の方向性。
KV キャッシュの再利用距離の推定に使用されるツール呼び出しレイテンシの予測は、次のことを考慮する必要があります。

ツールのカテゴリだけではなく、ツールの入力引数と実行履歴を考慮します。たとえば、同じ Bash 呼び出しは、1 秒未満の git チェックアウトから数分かかるコンパイルまで多岐にわたります。この方向での調査の一部は、コーディング エージェント ワークロードの再利用を意識した KV キャッシュ管理を研究する CacheWise に掲載されています。
4- プレフィックス キャッシュは効果的だが最適ではない
総合すると、プレフィックス キャッシュは優れているように見えます。入力トークンの 95.7% がキャッシュにヒットします。ただし、この平均値には、ステップを開始する 2 つの方法の間の明確な分裂が隠されています。ほとんどのステップはツールと結果の継続です。つまり、ツール呼び出しが終了すると、その結果がすぐに会話に追加され、モデルにフィードバックされます。これらのツール呼び出しのほとんどは待ち時間が短いため、最大 97.5% に達します。新しいユーザー メッセージから開始されるラウンドの処理ははるかに悪く、最大 84.4% に達するだけです。プロンプト キャッシュは非アクティブな状態が数分間続くと期限切れになり、ユーザーは日常的にリクエスト間よりも長く一時停止するため、次のリクエストが到着するまでにコンテキストが期限切れになり、最初から再入力する必要があります。
プレフィックス キャッシュ ミスの代償は多大です。処理された 23 億個の追加トークンのうち、真に新しいコンテキストを表すのは 19% のみです。残りの 81% は繰り返し作業であり、システムがすでに確認したトークンを再入力します。言い換えれば、すべての新しいトークンは平均して約 5 倍プレフィルされます。この冗長性を排除すると、サービスを提供するプロバイダーの事前入力コンピューティングが削減され、ユーザーの最初のトークンまでの時間が短縮されます。
潜在的な研究の方向性。
KV キャッシュ保持時間とストレージ オーバーヘッドの間のトレードオフの調査: 長いユーザー一時停止の間もエントリを維持すると、再プレフィル作業が節約されますが、メモリとストレージのコストが増大し、損益分岐点はモデル、ハードウェア、ワークロードによって異なる可能性があります。
私たちのリリースには、これを再現するために必要なものがすべてバンドルされています

分析し、独自の使用のために同様のトレースを構築するには、次のようにします。
トレースの例。生のクロード コードとコーデックス ログのサンプルと、正規化されサニタイズされたトレース形式を並べて表示します。
コレクションパイプライン。ローカルのクロード コードとコーデックスのログをクリーンな匿名化されたトレースに変換するトレース コレクターおよびサニタイザー。
データセットと分析スクリプト。完全な匿名化トレースと結果を再現するスクリプト。
リプレイクライアント。同様のリクエスト到着パターンを再構築して、vLLM や SGLang などのサービス エンジンを駆動するクライアント。
このトレースは、サービス提供システムにとって実際のコーディング エージェント トラフィックがどのようなものであるかを早期かつ具体的に示し、自己駆動ループ、ロングコンテキスト、ショート出力の性質、ロングテールのツール呼び出し、および適切ではあるが完璧ではないプレフィックス キャッシュ ヒット率を強調しています。私たちは、このトレースが偏っている可能性があることを認めています。パターンは私たちの特定のプロジェクトと習慣を反映しており、それがまさに私たちが完全なパイプラインをリリースする理由です。 Claude Code、Codex、またはその他のコーディング エージェントを使用している場合は、ぜひ https://tracelab.cs.washington.edu/ にアクセスし、自分のログで分析を実行し、よろしければサニタイズされたトレースを共有し、この最初のデータ ポイントをコミュニティが構築した共有リソースに変えるのにご協力ください。
TraceLab はオープンソースです。この成果を基にしている場合は、引用してください。

## Original Extract

Systems for Future Intelligence
Home
People
Publications
Talks
News
Blog
TraceLab: Characterizing Coding Agent Workloads for LLM Serving
Kan Zhu, Mathew Jacob, Chenxi Ma, Yi Pan, Stephanie Wang, Arvind Krishnamurthy, Baris Kasikci
As major AI labs ship their own coding agents—Claude Code from Anthropic, Codex from OpenAI, and Gemini CLI from Google—serving these agents efficiently is an increasingly important systems problem. Existing model-quality benchmarks, such as Terminal-Bench and SWE-bench, are poorly suited to modeling serving-system performance; they involve relatively few tool calls and focus on single, isolated tasks.
To close this gap, we are releasing the SyFI coding trace—a real-world dataset collected from our group’s everyday coding-agent usage—to guide the design of serving systems for coding agents. We are also open-sourcing the trace collection and analysis pipeline, TraceLab , making it easy to generate personalized traces and contribute them back to the dataset. The trace and code are available at https://github.com/uw-syfi/TraceLab , and a live demo at https://tracelab.cs.washington.edu/ .
Trace Collection and Trace Facts
We collect these traces from our own daily use of Claude Code and Codex doing research and development in the SyFI lab. The pipeline runs in two stages. First, a trace collector reads the raw, verbose session logs, extracts the key fields, and discards the surrounding logging scaffolding. Then, a trace sanitizer strips sensitive content—such as tool-call inputs and outputs—and anonymizes the rest by removing usernames and mangling session IDs. The resulting anonymized trace, at a scale of ~4,300 sessions and 55B tokens in total, is large enough for a serving engine or even a distributed serving fleet to reach steady state.
We summarize our major observations and provide suggestions for future systems research for better supporting agentic coding use cases.
1- Autonomous, Multi-Step Conversation
Coding agents do not answer in one shot. For each user request, the agent executes an autonomous loop of model generations and tool calls before returning a final answer. In our traces, each request takes an average of 8.8 self-directed steps, i.e., LLM-tools cycles, and issues 10.8 tool calls before giving the final answer. As a result, 88% of all LLM rounds respond to a tool result rather than a human.
The time spent within this autonomy is distributed unevenly across the loop. A single LLM generation averages ~13 s and each tool call ~18 s. End-to-end response time per request is heavy-tailed: the median is ~38 s, but the mean is ~4 min and the p99 is nearly 44 min. The largest gap, however, is between requests: the user spends time reading the output, thinking, and typing the next request, which takes an average of 46.7 min , despite a median of ~1.4 min.
Potential Research Directions.
Increasing the number of tool calls issued per round while holding the total count roughly fixed, thereby reducing the number of steps and increasing the opportunity for parallel tool execution
Using early signals that the user is about to start the next turn—the user re-activating the terminal window or beginning to type—to prefetch and re-prefill the conversation history during the long idle gaps between turns.
2- Long-Context, Short-Output Generation
Each round carries a very large prompt but generates comparatively little output. Across the trace, models read 52.56 B cached input tokens and prefill 2.34 B new ones , yet generate just 186.9 M output tokens —inputs outnumber outputs by 294× . A typical round sits on a 32k–256k-token prefix, appends only a few hundred to a few thousand tokens, and decodes a couple hundred out. However, the tail of new input tokens can go over 128K.
Speed-wise, the normalized decoding speed has a median of ~40.7 tok/s overall (Claude 46.8, Codex 33.9), and Codex’s pure-decoding rate reaches ~61.3 tok/s. Codex’s TTFT for each step (~3.1 s) is roughly 25% of a round’s generation time .
Potential Research Directions.
Short, incremental prefills need to be handled differently from those long prefills due to their distinct performance characteristics.
The per-step TTFT of returning from a tool result back into the next generation must be reduced through better request routing and auto-scaling policies for a shorter end-to-end time.
3- Tool-Heavy, Long-Tailed Execution
The autonomous LLM-tools-LLM loop runs almost entirely on a small set of tools, and those tool calls are dominated by the shell. Of 433 K tool calls , 76% are shell/command executions (running builds, tests, git , and similar commands), followed by file edits (11%) and file reads/searches (9%); planning, sub-agent, and web-lookup calls make up the rest. Claude draws on a far wider tool vocabulary (54 tools) than Codex (31), but both concentrate most of their volume in the same few—chiefly shell commands, file edits, and file reads/searches.
Tool latency is heavy-tailed both across tool types and within individual types , a tool like Bash spans four-plus orders of magnitude. The imbalance is stark: calls under 1 s are ~61% of all calls but only ~1% of total tool time, while the ~4% of calls lasting over a minute consume ~85% of it.
Potential Research Directions.
Tool-call latency prediction, used for estimating KV-cache reuse distance, should consider the tool’s input arguments and execution history rather than its category alone. For example, the same Bash call ranges from a sub-second git checkout to a multi-minute compile. Some of our explorations in this direction appear in CacheWise , which studies reuse-aware KV-cache management for coding-agent workloads.
4- Prefix Cache Effective But Not Optimal
In aggregate, prefix caching looks great: 95.7% of input tokens hit the cache. That average, however, hides a sharp split between the two ways a step can start. Most steps are tool-result continuations —a tool call finishes and its result is immediately appended to the conversation and fed back to the model. These tool calls mostly have short latency, so they hit ~97.5% . Rounds that start from a new user message fare much worse, hitting only ~84.4% : prompt caches expire after a few minutes of inactivity, and users routinely pause longer than that between requests, so by the time the next request arrives, the context has aged out and must be re-prefilled from scratch.
The cost of prefix cache miss is substantial. Of the 2.3 B append tokens processed, only 19% represent genuinely new context; the other 81% is repeated work , re-prefilling tokens the system has already seen. Put differently, every new token is prefilled roughly 5× on average . Eliminating this redundancy would cut prefill compute for serving providers and shorten time-to-first-token for users.
Potential Research Directions.
Investigation of the trade-off between KV-cache retention time and storage overhead: keeping entries alive across long user pauses saves re-prefill work but inflates memory and storage cost, and the break-even point likely varies with the model, hardware, and workload.
Our release bundles everything needed to reproduce this analysis and to build similar traces for your own usage:
Trace examples. Side-by-side samples of raw Claude Code and Codex logs alongside our normalized, sanitized trace format.
Collection pipeline. The trace collector and sanitizer that turn local Claude Code and Codex logs into clean, de-identified traces.
Dataset and analysis scripts. The full anonymized trace together with the scripts that reproduce the results.
Replay client. A client that reconstructs a similar request-arrival pattern to drive serving engines such as vLLM and SGLang.
This trace offers an early, concrete look at what real coding-agent traffic looks like to a serving system, highlighting the self-driven loop, the long-context, short-output nature, long-tailed tool calling, and a decent but not perfect prefix-cache hit rate. We admit this trace can be biased : the patterns reflect our particular projects and habits, which is exactly why we are releasing the full pipeline. If you use Claude Code, Codex, or any other coding agent, we’d love for you to visit https://tracelab.cs.washington.edu/ , run the analysis on your own logs, share a sanitized trace if you’re comfortable, and help turn this first data point into a shared, community-built resource.
TraceLab is open source. If you build on this work, please cite it:
