---
source: "https://blog.roboflow.com/openai-gpt-5-6/"
hn_url: "https://news.ycombinator.com/item?id=49329575"
title: "GPT 5.6 Sol is the best \"vision\" model OpenAI ever released"
article_title: "GPT 5.6 Sol is the best \"vision\" model OpenAI ever released"
image: "https://storage.ghost.io/c/2c/8d/2c8d8c0d-1c15-4b6d-825e-02b78d61d40a/content/images/2026/07/img-blog-gpt-5-6-sol-best-vision-model-to-date.jpg"
author: "plurby"
captured_at: "2026-08-17T12:22:46Z"
capture_tool: "hn-digest"
hn_id: 49329575
score: 2
comments: 0
posted_at: "2026-08-17T12:09:42Z"
tags:
  - hacker-news
  - translated
---

# GPT 5.6 Sol is the best "vision" model OpenAI ever released

- HN: [49329575](https://news.ycombinator.com/item?id=49329575)
- Source: [blog.roboflow.com](https://blog.roboflow.com/openai-gpt-5-6/)
- Score: 2
- Comments: 0
- Posted: 2026-08-17T12:09:42Z

## Translation

タイトル: GPT 5.6 Sol は OpenAI がこれまでにリリースした最高の「ビジョン」モデルです
説明: GPT-5.6 Sol は、OpenAI のこれまでで最も強力なビジョン モデルです。私たちは Sol、Terra、Luna の検出、計数、OCR、抽出をテストし、その結果、速度、コストを主要な VLM と比較しました。

記事本文:
GPT 5.6 Sol は OpenAI がこれまでにリリースした最高の「ビジョン」モデルです
製品
プラットフォーム
デプロイ
デバイス上、エッジ上、VPC 内、または API 経由でモデルを実行します
ワークフロー
パイプラインとアプリケーションを構築するためのローコード インターフェイス
電車
ホストされたモデルのトレーニング インフラストラクチャと GPU アクセス
注釈を付ける
AI 支援のデータ注釈を使用して画像に迅速にラベルを付ける
宇宙
オープンソースのコンピューター ビジョン データセットと事前トレーニングされたモデル
業界別
あらゆる業界のソリューションを調べる
航空宇宙と防衛
サインイン
デモを予約する
始めましょう
検索
サインイン
デモを予約する
始めましょう
検索
ブログ
GPT 5.6 Sol は OpenAI がこれまでにリリースした最高の「ビジョン」モデルです
ピョートル・スカルスキー
発行済み
2026 年 7 月 16 日
•
6 分で読めます
先週、OpenAI は GPT-5.6 ラインナップを発表し、Sol、Terra、Luna モデルを導入しました。リリース ストリーム中、チームはコンピュータの使用に重点を置き、デスクトップ アプリケーションをナビゲートおよび操作できるモデルを示しました。 OpenAI は UI エージェントと詳細な 3D ビジュアライゼーションを強調しましたが、どちらもより強力な視覚的理解に依存しています。
彼らの視覚能力を測定するために、今後数週間以内にリリースする予定の VLM ベンチマークでモデルを実行しました。このベンチマークは、検出、計数、OCR、データ抽出などの一般的な視覚タスクをカバーします。この投稿では、それぞれの GPT-5.6 がどのように動作するかを詳しく見ていきます。
Sol は明らかに OpenAI がこれまでにリリースした最高のビジョン モデルです。この飛躍は特に物体検出と計数において顕著であり、GPT-5.5 は最も強力な VLM に大きく遅れをとっていた。 Terra と Luna は Sol ほど強力ではありませんが、どちらも GPT-5.5 よりも意味のある進歩を示しています。
Roboflow Playground で Sol、Terra、Luna をテストし、同じ視覚タスクでその結果を Claude Fable 5 や Gemini 3.5 Flash などのモデルと比較します。
検出は GPT-5.6 が示すところです

最も明確なジャンプ。 GPT-5.5 はベンチマークで 13.8 mAP@50 のスコアを記録しましたが、Sol は 46.2 に達しました。 Terra と Luna が 44.7 と 43.3 で僅差で続き、物体検出が大きな弱点から実用的な機能に移行しました。
文書レイアウトの検出は、GPT-5.6 の最も明確な長所の 1 つです。 Sol は、タイトル、段落、表、画像、署名をうまく処理しました。多くのドキュメント ワークフローは、OCR またはデータ抽出を開始する前に、ページの関連部分を見つけることから始まります。
GPT-5.6 は密度の高いシーンでも良好なパフォーマンスを発揮しました。錠剤と卵の例には、密集して詰め込まれた類似のオブジェクトが多数含まれており、これが VLM ベースの検出の一般的な弱点となります。従来の検出器とは異なり、VLM は各クラス ラベルと座標のセットをテキストとして生成します。オブジェクトの数が増えると、応答が長くなり、オブジェクトの欠落、重複、または座標エラーのリスクが増加します。それにもかかわらず、Sol は両方のシーンでほとんどのオブジェクトを検出しました。
最良の検出結果を得るには、GPT-5.6 モデルに絶対 XYXY 座標を画像ピクセルで返すように指示します。これは、0 ～ 1000 の範囲に正規化された YXYX 座標で最高のパフォーマンスを発揮した Gemini 3.5 フラッシュとは異なります。間違った座標形式を使用すると、ベンチマークで GPT-5.6 検出パフォーマンスが約 15 mAP ポイント低下しました。
いくつかのケースでは、GPT-5.6 Sol は画像の一見ランダムな部分にボックスを返しました。多くは、地上の真実と重複しないか、ほとんど重複しませんでした。ボックスは、目に見えるオブジェクトと一致するのではなく、直線の列や等間隔のグループなど、不自然なレイアウトを形成することがよくありました。
これらの例を OpenAI と共有しました。彼らのチームは、約 2,000 × 2,000 ピクセル以上の画像では、特に推論の労力が低い場合に、Sol の安定性が低下することを確認しました。推論の労力が増えると安定性が向上しますが、トークンの使用量、待ち時間、コストも増加します。サイズ変更 o

大きな画像を OpenAI API に送信する前にトリミングすることが、最も現実的な回避策です。
GPT-5.6 の全ラインナップにわたってカウンティングが向上しました。 Sol はベンチマークで 73.0% のスコアを獲得し、GPT-5.5 の 64.9% から増加しました。一方、Terra と Luna は 67.6% と 66.2% に達しました。ラインナップの中で最も安価なモデルである Luna は、依然として以前の OpenAI ベースラインを上回りました。
ベンチマークの一部として、オブジェクトを検出して合計を返す以上のことが必要なケースをテストしました。 Sol は、重なり合う金属ブラケットをカウントしましたが、これは従来の物体検出器と VLM の両方にとって困難なケースでした。ソル氏はまた、選択した得点ゾーン内でのみ弾痕を数え、どの物体を数えるべきか、ルールが適用される場所の両方を理解していることを示しました。
ブリスターパックははるかに難しいことが判明しました。別のプロンプトで、ソルに、空のスロットとパッケージ内にまだ密封されている錠剤を数えるように依頼しました。繰り返されるレイアウト、反射、および埋められたスロットと空のスロットの間の小さな視覚的な違いにより、両方のタスクが困難になりました。
異常なキャンディーの例では、別のタイプの失敗が明らかになりました。ソルは間違った数を数えましたが、モデルがキャンディーの数を数え間違えたのか、ターゲット カテゴリを誤解したのかは不明です。
OCR パフォーマンスは GPT-5.5 に近いままでした。 Sol は平均類似性スコア 90.7% を達成し、GPT-5.5 の 91.2% とわずか 0.5 ポイント差でしたが、Terra と Luna は 88.8% と 88.4% に達しました。テキスト抽出ではギャップがさらに大きく、Sol のスコアは 82.5% でしたが、GPT-5.5 のスコアは 87.6% でした。ルナとテラが81.4％、79.4％で続いた。
ベンチマークの一環として、完全な転写をターゲット抽出から分離しました。 OCR では、表示されているすべてのテキストを書き写すようにモデルに要求しますが、テキスト抽出では特定の情報を要求します。 Sol はどちらの設定でも手書きのメモで優れたパフォーマンスを発揮し、1 つのケースでは完全な文字起こしを作成し、要求されたデータを抽出しました。

別のe。
Sol は、複雑なビジュアル シーンに埋め込まれたテキストで優れたパフォーマンスを発揮しました。それは、汚れて摩耗したタイヤの曲面に沿って印刷されたタイヤサイズのシーケンスを読み取りました。別の例では、ホッケーの放送からライブスコアを抽出し、要求された形式で回答を返し、視覚的な読み取りと指示のフォローの両方をテストしました。
単純に見えるいくつかの抽出タスクは依然として失敗しました。ソルはブリスターパックに記載されている使用期限を読み取ることができませんでした。テキストは小さく、縦長で、コントラストが低く、反射の影響を受けており、これがエラーの原因である可能性があります。
GPT-5.6 ラインナップ全体でトークンの使用量が増えると、ビジョンが向上します。小規模なテストでは違いはそれほど重要ではありませんが、トークンの量が処理コストを直接増加させる大規模なテストでは、より重要になります。
Sol のベンチマークでは、画像あたり平均 10 秒近くかかりました。 Terra はそれを約 6 秒に短縮しましたが、Luna は 5 秒強で終了しました。 Luna は、ラインナップの中で最も強力なレイテンシーと品質のバランスを提供し、Gemini 3.5 Flash に近い速度を実現しながら、検出とカウントでは GPT-5.5 を上回ります。
私たちのベンチマークでは、Sol のコストは画像あたり約 2.5 セントで、Claude Fable 5 に次いで 2 番目に高価なモデルとなっています。 Terra は画像あたりの平均コストを約 1 セントに削減しましたが、Luna のコストは 0.5 セント未満でした。
Gemini 3.5 Flash は画像あたり 0.8 セントで、Sol よりもはるかに安価ですが、依然として検出とカウントのベンチマークをリードしています。これは、大規模な画像バッチ全体でコストが増加するデータ集約型のワークロードにとって強力なオプションになります。 Roboflow Playground を使用すると、Sol、Terra、Luna を、Claude Fable 5、Gemini 3.5 Flash、および他の VLM と同じタスクでテストできます。
GPT-5.6 では、OpenAI は以前よりも主要な VLM に大幅に近づきました。検出が弱点から使用可能な能力に移行し、インプをカウント

モデルファミリー全体を歩き回りました。
まだ明確な限界があります。 Gemini 3.5 フラッシュは、特にその価格において、ベンチマークにおける大量の検出とカウントには依然として優れた実用的な選択肢です。
GPT-5.6 は、OpenAI がビジョンをより真剣に捉えていることを示しています。 Sol には、特にコスト、遅延、検出が不安定なケースなど、依然として欠陥がありますが、進歩は無視できません。エージェント、画面の理解、ドキュメントのワークフロー、および視覚的な推論にとって、このリリースでは OpenAI が以前よりもはるかに強力なオプションになります。
研究でこの投稿を引用するには、次のエントリを使用します。
ピョートル・スカルスキー。 (2026 年 7 月 16 日)。
GPT 5.6 Sol は、OpenAI がこれまでにリリースした最高の「ビジョン」モデルです。ロボフローブログ: https://blog.roboflow.com/openai-gpt-5-6/
AI バスケットボール ショット評価ツールを構築する方法
AI データラベリングのトップモデル
AI ビデオ自動車損傷検査装置を構築する方法
クラウドに接続された AI 製品がオンプレミスで実行される仕組み
Qwen3.8-Max for Vision: ベンチマーク、強み、および実際のテスト
ドローン画像から小さな物体を検出する方法

## Original Extract

GPT-5.6 Sol is OpenAI’s strongest vision model yet. We tested Sol, Terra, and Luna across detection, counting, OCR, and extraction, then compared their results, speed, and cost with leading VLMs.

GPT 5.6 Sol is the best "vision" model OpenAI ever released
Products
Platform
Deploy
Run models on device, at the edge, in your VPC, or via API
Workflows
Low-code interface to build pipelines and applications
Train
Hosted model training infrastructure and GPU access
Annotate
Label images fast with AI-assisted data annotation
Universe
Open source computer vision datasets and pre-trained models
By Industry
Explore all industry solutions
Aerospace & Defense
Sign In
Book a demo
Get Started
Search
Sign in
Book a demo
Get Started
Search
Blog
GPT 5.6 Sol is the best "vision" model OpenAI ever released
Piotr Skalski
Published
Jul 16, 2026
•
6 min read
Last week, OpenAI announced the GPT-5.6 lineup, introducing the Sol, Terra, and Luna models. During the release stream , the team focused heavily on computer use , showing models capable of navigating and operating desktop applications. OpenAI highlighted UI agents and detailed 3D visualizations, but both depend on stronger visual understanding.
To measure their vision capabilities, we ran the models through our upcoming VLM benchmark, which we plan to release in the next few weeks. The benchmark covers common vision tasks, including detection, counting, OCR, and data extraction. In this post, we take a closer look at how GPT-5.6 performs across each of them.
Sol is clearly the best vision model OpenAI has released so far. The jump is especially visible in object detection and counting, where GPT-5.5 was far behind the strongest VLMs. Terra and Luna are not as strong as Sol, but both show meaningful progress over GPT-5.5.
Test Sol, Terra, and Luna in Roboflow Playground and compare their results with models such as Claude Fable 5 and Gemini 3.5 Flash across the same vision tasks.
Detection is where GPT-5.6 shows the clearest jump. GPT-5.5 scored 13.8 mAP@50 in our benchmark, while Sol reached 46.2. Terra and Luna followed closely at 44.7 and 43.3, moving object detection from a major weakness to a practical capability.
Document layout detection is one of the clearest strengths of GPT-5.6. Sol handled titles, paragraphs, tables, images, and signatures well. Many document workflows start with locating the relevant parts of a page before OCR or data extraction begins.
GPT-5.6 also performed well on dense scenes. The pills and eggs examples contain many similar objects packed closely together, a common weakness for VLM-based detection. Unlike traditional detectors, VLMs generate each class label and set of coordinates as text. As object count grows, the response becomes longer and the risk of missed objects, duplicates, or coordinate errors increases. Despite this, Sol detected most objects across both scenes.
For the best detection results, prompt GPT-5.6 models to return absolute XYXY coordinates in image pixels. This differs from Gemini 3.5 Flash, which performed best with YXYX coordinates normalized to a 0–1000 range. Using the wrong coordinate format reduced GPT-5.6 detection performance by around 15 mAP points in our benchmark.
In a few cases, GPT-5.6 Sol returned boxes in seemingly random parts of the image. Many had no overlap, or almost no overlap, with the ground truth. Instead of matching the visible objects, the boxes often formed unnatural layouts, such as straight rows or evenly spaced groups.
We shared those examples with OpenAI. Their team confirmed that Sol becomes less stable on images around 2,000 by 2,000 pixels or larger, especially at lower reasoning effort. Higher reasoning effort improves stability, but also increases token use, latency, and cost. Resizing or cropping large images before sending them to the OpenAI API is the most practical workaround.
Counting improved across the full GPT-5.6 lineup. Sol scored 73.0% in our benchmark, up from 64.9% for GPT-5.5, while Terra and Luna reached 67.6% and 66.2%. Luna, the cheapest model in the lineup, still outperformed the previous OpenAI baseline.
As part of the benchmark, we tested cases requiring more than spotting objects and returning a total. Sol counted heavily overlapping metal brackets, a difficult case for both traditional object detectors and VLMs. Sol also counted bullet holes only inside selected scoring zones, showing an understanding of both which objects to count and where the rule applied.
Blister packs proved much harder. In separate prompts, we asked Sol to count the empty slots and the pills still sealed inside the package. The repeated layout, reflections, and small visual differences between filled and empty slots made both tasks difficult.
The abnormal candy example exposed a different type of failure. Sol gave the wrong count, though it is unclear whether the model miscounted the candies or misunderstood the target category.
OCR performance stayed close to GPT-5.5. Sol achieved a 90.7% mean similarity score, only 0.5 points behind GPT-5.5 at 91.2%, while Terra and Luna reached 88.8% and 88.4%. The gap was larger in text extraction, where Sol scored 82.5% compared with 87.6% for GPT-5.5. Luna and Terra followed at 81.4% and 79.4%.
As part of the benchmark, we separated full transcription from targeted extraction. OCR asks the model to transcribe all visible text, while text extraction asks for a specific piece of information. Sol performed well on handwritten notes in both settings, producing a full transcription in one case and extracting a requested date in another.
Sol performed well on text embedded in complex visual scenes. It read a tire size sequence printed along the curved surface of a dirty, worn tire. In another example, it extracted the live score from a hockey broadcast and returned the answer in the requested format, testing both visual reading and instruction following.
Some simple-looking extraction tasks still failed. Sol could not read the expiration date printed on a blister pack. The text was small, vertical, low contrast, and affected by reflections, which may explain the error.
The vision gains come with higher token usage across the GPT-5.6 lineup. The difference matters less in small tests, but becomes more important at scale, where token volume directly increases processing costs.
Sol averaged close to 10 seconds per image in our benchmark. Terra reduced that to around 6 seconds, while Luna finished in slightly over 5 seconds. Luna offers the strongest latency-quality balance in the lineup, with speed close to Gemini 3.5 Flash while still outperforming GPT-5.5 on detection and counting.
In our benchmark, Sol cost roughly 2.5 cents per image, making it the second most expensive model after Claude Fable 5 . Terra reduced the average cost to about 1 cent per image, while Luna cost less than 0.5 cents.
At 0.8 cents per image, Gemini 3.5 Flash is much cheaper than Sol while still leading our detection and counting benchmarks. This makes it a strong option for data-intensive workloads where cost scales across large image batches. Roboflow Playground lets you test Sol, Terra, and Luna alongside Claude Fable 5, Gemini 3.5 Flash, and other VLMs on the same tasks.
With GPT-5.6, OpenAI is much closer to the leading VLMs than before. Detection moved from a weak point to a usable capability, and counting improved across the full model family.
There are still clear limits. Gemini 3.5 Flash remains a better practical choice for high-volume detection and counting in our benchmark, especially at its price.
GPT-5.6 shows OpenAI is now taking vision much more seriously. Sol still has flaws, especially around cost, latency, and some unstable detection cases, but the progress is hard to ignore. For agents, screen understanding, document workflows, and visual reasoning, this release makes OpenAI a much stronger option than before.
Use the following entry to cite this post in your research:
Piotr Skalski . (Jul 16, 2026).
GPT 5.6 Sol is the best "vision" model OpenAI ever released. Roboflow Blog: https://blog.roboflow.com/openai-gpt-5-6/
How to Build an AI Basketball Shot Evaluator
Top Models for AI Data Labeling
How to Build an AI Video Car Damage Inspector
How Cloud Connected AI Products Run On-Prem
Qwen3.8-Max for Vision: Benchmarks, Strengths, and Real-World Tests
How to Detect Small Objects in Drone Imagery
