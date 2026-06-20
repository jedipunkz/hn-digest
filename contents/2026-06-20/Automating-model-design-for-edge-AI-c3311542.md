---
source: "https://deepgate.ai/blog/neural-architecture-search"
hn_url: "https://news.ycombinator.com/item?id=48604630"
title: "Automating model design for edge AI"
article_title: "Automating model design for edge AI. | DeepGate"
author: "webstorms"
captured_at: "2026-06-20T00:56:43Z"
capture_tool: "hn-digest"
hn_id: 48604630
score: 1
comments: 0
posted_at: "2026-06-19T23:38:08Z"
tags:
  - hacker-news
  - translated
---

# Automating model design for edge AI

- HN: [48604630](https://news.ycombinator.com/item?id=48604630)
- Source: [deepgate.ai](https://deepgate.ai/blog/neural-architecture-search)
- Score: 1
- Comments: 0
- Posted: 2026-06-19T23:38:08Z

## Translation

タイトル: エッジ AI のモデル設計の自動化
記事のタイトル: エッジ AI のモデル設計の自動化。 |ディープゲート
説明: DeepGate のモデル検索およびコンパイラーは、同じマイクロコントローラー上のリファレンス モデルと比べて、最大 45 倍高速に実行し、最大 11 倍少ない RAM を使用するモデルを生成します。

記事本文:
エッジ AI のモデル設計を自動化します。 | DeepGate DeepGate ソリューション R&D チーム ブログ プラットフォーム お問い合わせ 2026 年 6 月 19 日
エッジ AI のモデル設計を自動化します。
DeepGate のモデル検索とコンパイラは、同じマイクロコントローラ上のリファレンス モデルと比べて最大 45 倍高速に実行し、最大 11 倍少ない RAM を使用するモデルを生成します。
マイクロコントローラーのモデルの構築は、依然として大部分が手作業のプロセスです。チームはモデルをゼロから設計するか、既存のアーキテクチャを適応させて、ターゲットのハードウェアに合わせて繰り返し変更します。リソースに制約のあるデバイスでは、実行するには大きすぎるか遅すぎるモデルと、デバイスには適合するがミスが多すぎて役に立たないモデルとの間のトレードオフに直面することがよくあります。
私たちは自動モデル設計システムの基礎を構築しました。ニューラル アーキテクチャ検索、DeepGate コンパイラ、開発プラットフォームを通じて得られるリアルハードウェア測定を組み合わせることで、ターゲットのマイクロコントローラに合わせたモデルを自動的に検索できます。音声内の話し言葉の検出から画像内の人物の存在の識別に至るまで、4 つの標準的な MLPerf Tiny ベンチマーク タスク全体にわたって、結果として得られたモデルは、参照モデルと比べて最大 45 倍高速に実行され、RAM の使用量は最大 11 倍少なくなりました。たとえば、Analog Devices MAX32655 で実行される MLPerf Tiny キーワード スポッティング ベンチマークでは、検索により、90% 以上の分類精度を維持しながら、推論レイテンシーが 104.3 ミリ秒から 2.3 ミリ秒に、RAM 使用量が 23.7 KB から 2.1 KB に減少しました。
このような利点により、機械学習モデルをより安価なハードウェアで実行し、バッテリー寿命を延長し、メモリを解放して他のタスクのためにコンピューティングできるようになります。効率性のフロンティアを推進することで、より高度な AI ワークロードをマイクロコントローラーの届く範囲内に移動し、ますます高機能なインテリジェンスを数十億人に提供します

デバイスの。
同じハードウェア上のリファレンス モデルを上回るパフォーマンス
私たちは、マイクロコントローラーでの機械学習の標準ベンチマーク スイートである MLPerf Tiny v1.4 での検索を評価しました。このベンチマークは、キーワード スポッティング、ビジュアル ウェイク ワード、CIFAR-10 画像分類、異常検出という 4 つの代表的なエッジ ワークロードをカバーしています。各タスクには、キーワード検出のトップ 1 精度 90% から異常検出の 0.85 AUC まで、事前に定義された品質目標があります。各ワークロードの目標は、参照モデルとの公正な比較を確保するために入力ディメンションを固定して、可能な限り最小かつ最速のモデルを作成しながら目標を達成することでした。
評価したボード全体で、当社の検索システムとコンパイラーは、最大 45 倍の高速推論と最大 11 倍の RAM 使用量の削減を実現しました。多くの場合、メモリはマイクロコントローラの主な制約となるため、これらのメモリ削減は特に重要です。場合によっては、ベンダー ツールチェーンのメモリ制限を超えたモデルでも、検索とコンパイル後に適合して正常に実行できることがありました。
以下の結果は、各ベンダーのツールチェーンでコンパイルされた MLPerf Tiny 参照モデルと、当社の検索システムによって自動的に検出され、DeepGate コンパイラーでデプロイされたアーキテクチャとを比較したもので、すべての結果は同じハードウェアで測定されました。ボードを切り替えたり、レイテンシーと RAM 使用量を切り替えたりして比較を調べてください。 RAM は、テンソル アリーナにピーク スタック サイズを加えたものとして測定されます。
DeepGate は最大 36.1 倍高速に動作します
実現方法: 2 つの補完的な検索方法
私たちは 2 つの検索システムを並行して実行し、特定のタスクに最も優れたパフォーマンスを発揮する方を使用しました。 MLPerf Tiny ワークロードでは、4 つの最終モデルのうち 3 つはニューラル アーキテクチャ検索 (NAS) システムから取得され、異常検出モデルはエージェント シーから取得されました。

チャーチ。
エージェント アーキテクチャ検索では、アーキテクチャまたはトレーニング レシピに対して一度に 1 つの変更を提案する LLM エージェントを使用します。結果のモデルをトレーニングし、実際のハードウェアでベンチマークを実行し、ターゲット メトリックが改善した場合にのみ変更を保持します。このアプローチには制限がなく、事前に定義された検索スペースの外側でアイデアを探索できますが、貪欲に機能し、一度に 1 つのモデルを改善します。
Supernet NAS は、ワンス・フォー・オールおよび MCUNet のアプローチを基にして拡張されており、参照モデルとの公正な比較のために入力解像度を固定しながら、int8 量子化対応トレーニングを使用してマイクロコントローラーの導入に適合しています。すべての候補アーキテクチャを個別にトレーニングするのではなく、単一のスーパーネットを、サイズ、速度、精度のトレードオフが異なる多くの異なるモデルに特化させることができます。
2 つのアプローチは、次のような補完的な長所を提供します。
どちらのアプローチも同じ社内インフラストラクチャ上で実行されます。各モデルは、DeepGate コンパイラーによって効率的な静的バイナリにコンパイルされ、複数のボードにわたって統合されたベンチマーク API を提供する開発プラットフォームを通じてターゲットのマイクロコントローラーにデプロイされます。結果として生じる遅延とメモリ使用量は、ターゲット ハードウェア上で直接測定されます。
私たちの長期的な目標は、タスクの定義からエッジ デバイスへの最適化されたモデルの展開に至るまで、高効率モデルの設計を自動化することです。これを達成するために、私たちは NAS とエージェントの検索方法を組み合わせて、両方のアプローチの長所を統合する単一の最適化ループを作成する方法を模索しています。
同時に、従来のニューラル ネットワーク レイヤーよりもメモリ使用量を減らし、高速に実行するように設計された新しい DeepGate レイヤーなど、検索システムで利用できるニューラル ネットワーク レイヤーのセットを拡張しています。これらのレイヤーを に組み込む

アーススペースは、リソースに制約のあるデバイスでさらに高い効率を実現し、これまでマイクロコントローラーの範囲を超えていると考えられていた AI ワークロードを可能にし、最終的には数十億台のデバイスにますます高性能なインテリジェンスをもたらします。
ご自身のモデルを縮小することに興味がある場合、または最適化されたビジョンおよびオーディオ モデルにアクセスすることに興味がある場合は、ぜひご連絡をお待ちしております。
Jacob 他、効率的な整数演算のみの推論のためのニューラル ネットワークの量子化とトレーニング、CVPR 2018。
Cai、Gan、Wang、Zhang、Han、ワンス・フォー・オール: 1 つのネットワークを訓練し、効率的な展開に特化する、ICLR 2020。
リン、チェン、リン、コーン、ガン、ハン、MCUNet: IoT デバイス上の小さな深層学習、NeurIPS 2020。

## Original Extract

DeepGate's model search and compiler produce models that run up to 45× faster and use up to 11× less RAM than reference models on the same microcontroller.

Automating model design for edge AI. | DeepGate DeepGate Solutions R&D Team Blog Platform Get in touch June 19, 2026
Automating model design for edge AI.
DeepGate's model search and compiler produce models that run up to 45× faster and use up to 11× less RAM than reference models on the same microcontroller.
Building models for microcontrollers is still largely a manual process. Teams either design models from scratch or adapt existing architectures, iteratively modifying them to fit the target hardware. On resource-constrained devices, they often face a trade-off between models that are too large or slow to run and models that fit on the device but make too many mistakes to be useful.
We’ve built the foundations of an automated model design system. By combining neural architecture search, the DeepGate compiler , and real-hardware measurements obtained through our development platform , we can automatically search for models tailored to a target microcontroller. Across the four standard MLPerf Tiny benchmark tasks, ranging from detecting spoken words in audio to identifying the presence of a person in an image, the resulting models ran up to 45× faster and used up to 11× less RAM than the reference models. For example, on the MLPerf Tiny keyword spotting benchmark running on the Analog Devices MAX32655, our search reduced inference latency from 104.3 ms to 2.3 ms and RAM usage from 23.7 KB to 2.1 KB , while maintaining over 90% classification accuracy.
Such gains can enable machine learning models to run on cheaper hardware, extend battery life, and free up memory and compute for other tasks. By pushing the efficiency frontier, we move more advanced AI workloads within reach of microcontrollers, bringing increasingly capable intelligence to billions of devices.
Outperforming the reference models on the same hardware
We evaluated our search on MLPerf Tiny v1.4 , the standard benchmark suite for machine learning on microcontrollers. The benchmark covers four representative edge workloads: keyword spotting, visual wake words, CIFAR-10 image classification, and anomaly detection. Each task has a predefined quality target, from 90% top-1 accuracy for keyword spotting to 0.85 AUC for anomaly detection. For each workload, the goal was to meet the target while producing the smallest and fastest model possible, with input dimensions kept fixed to ensure a fair comparison against the reference models.
Across the evaluated boards, our search system and compiler delivered up to 45× faster inference and up to 11× lower RAM usage. Because memory is often the primary constraint on microcontrollers, these memory reductions can be especially important: in some cases, models that exceeded memory limits under the vendor toolchain were able to fit and run successfully after search and compilation.
The results below compare the MLPerf Tiny reference model compiled with each vendor’s toolchain against architectures automatically discovered by our search system and deployed with the DeepGate compiler, with all results measured on the same hardware. Explore the comparisons by switching boards and toggling between latency and RAM usage; RAM is measured as the tensor arena plus peak stack size.
DeepGate runs up to 36.1× faster
How we did it: two complementary search methods
We ran two search systems side by side and used whichever performed best for a given task. On the MLPerf Tiny workloads, three of the four final models came from our neural architecture search (NAS) system, while the anomaly detection model came from our agentic search.
Agentic architecture search uses an LLM agent that proposes one change at a time – either to the architecture or the training recipe – trains the resulting model, benchmarks it on real hardware, and keeps the change only if the target metric improves. The approach is open-ended and can explore ideas outside any predefined search space, but it operates greedily, improving one model at a time.
Supernet NAS builds on and extends the Once-for-All and MCUNet approaches, adapted for microcontroller deployment using int8 quantization-aware training while keeping input resolution fixed for fair comparison against the reference models. Rather than training every candidate architecture independently, a single supernet can be specialised into many different models with different size, speed, and accuracy trade-offs.
The two approaches offer complementary strengths:
Both approaches run on the same in-house infrastructure. Each model is compiled into an efficient static binary by the DeepGate compiler and deployed to target microcontrollers through our development platform, which provides a unified benchmarking API across multiple boards. The resulting latency and memory usage are measured directly on the target hardware.
Our long-term goal is to automate the design of highly efficient models, from defining a task to deploying an optimised model on an edge device. To achieve this, we are exploring how to combine our NAS and agentic search methods into a single optimisation loop that unifies the strengths of both approaches.
At the same time, we’re expanding the set of neural network layers available to the search system, including novel DeepGate layers designed to use less memory and run faster than conventional neural network layers. Incorporating these layers into the search space will unlock even greater efficiency on resource-constrained devices, enabling AI workloads once thought beyond the reach of microcontrollers – and ultimately bringing increasingly capable intelligence to billions of devices.
If you’re interested in shrinking your own models – or accessing our optimised vision and audio models – we’d love to hear from you.
Jacob et al., Quantization and Training of Neural Networks for Efficient Integer-Arithmetic-Only Inference , CVPR 2018.
Cai, Gan, Wang, Zhang, Han, Once-for-All: Train One Network and Specialize it for Efficient Deployment , ICLR 2020.
Lin, Chen, Lin, Cohn, Gan, Han, MCUNet: Tiny Deep Learning on IoT Devices , NeurIPS 2020.
