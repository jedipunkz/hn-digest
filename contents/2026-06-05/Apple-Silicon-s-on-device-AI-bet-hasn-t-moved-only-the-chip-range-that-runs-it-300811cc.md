---
source: "https://tbreak.com/apple-silicon-on-device-ai-doug-brooks-wwdc/"
hn_url: "https://news.ycombinator.com/item?id=48416850"
title: "Apple Silicon's on-device AI bet hasn't moved – only the chip range that runs it"
article_title: "Apple Silicon's on-device AI bet: same play, bigger range"
author: "Austin_Conlon"
captured_at: "2026-06-05T21:37:01Z"
capture_tool: "hn-digest"
hn_id: 48416850
score: 3
comments: 0
posted_at: "2026-06-05T19:08:56Z"
tags:
  - hacker-news
  - translated
---

# Apple Silicon's on-device AI bet hasn't moved – only the chip range that runs it

- HN: [48416850](https://news.ycombinator.com/item?id=48416850)
- Source: [tbreak.com](https://tbreak.com/apple-silicon-on-device-ai-doug-brooks-wwdc/)
- Score: 3
- Comments: 0
- Posted: 2026-06-05T19:08:56Z

## Translation

タイトル: Apple Silicon のオンデバイス AI への賭けは動かず、それを実行するチップ範囲のみが動いている
記事のタイトル: Apple Silicon のオンデバイス AI への賭け: 同じプレイ、より大きな範囲
説明: Apple のニューラル エンジン、ユニファイド メモリ、フュージョンに関する Apple Silicon のシニア PM、そしてオンデバイス AI への賭けが 9 年間動かない理由について語ります。

記事本文:
サインアップ
2026 年 6 月 5 日公開
2026 年 6 月 5 日更新
Apple Silicon のオンデバイス AI への賭けは動かず、それを実行するチップの範囲だけが動いている
WWDC の数日前、Apple のシリコン チームは次のように主張しました。オンデバイス AI への賭けは 9 年間変化しておらず、範囲全体で複雑化しているだけです。
WWDC がクパチーノで開幕する 3 日前、NVIDIA、クアルコム、インテルが台北のステージで順番に登壇し、クラウドとデバイス間のハイブリッド AI について、自分たちの言葉で主張した Computex 週間の終わりに、Apple のシリコン チームは静かな主張をしています。デバイス上で AI を実行するという彼らの賭けは、実際には動いていません。
Neural Engine は 2017 年から Apple チップに搭載されています。最初の A シリーズ パーツ以来、ユニファイド メモリは基本的な設計の選択肢となってきました。 9 年間で変わったのは哲学ではありません。それを実現するシリコンの製品群が、現在では 599 ドルの MacBook Neo から、2 つのダイが 1 つのパッケージに結合された M5 Max ワークステーション チップにまで広がっています。
本質的には、これが Computex の傍聴席でのブリーフィングにおける Doug Brooks の主張だった。 Brooks 氏は Apple Silicon のシニア プロダクト マネージャーであり、会話は記録上で製品に焦点を当てたもので、デモもマーケティング スライドもロードマップもありませんでした。チップだけです。
彼の冒頭の一文は、その後のすべての背骨となるため、引用する価値があります。
AI 用の優れたデバイスを構築したい場合は、優れたコンピューターを構築する必要があります。
これを回避策として読むこともできるし、Apple が AI の議論を自社がすでに製造しているハードウェアに戻そうとしていることを主張していると読むこともできるし、オンデバイス AI は電話や Mac にボルトオンで搭載されるわけではない、という実際の論文として読むこともできる。それはシリコンから設計されており、2017 年に A11 Bionic にニューラル エンジンが導入されて以来、Apple はまさにそれを行ってきました。ブルックス氏が言うように、博士のちらつきとともに

yness、「AI PC のトレンドが崩れるずっと前から」
同じ週に、他の 3 人のシリコン CEO が登壇しました。
Brooks 氏のケースは、Computex の残りの部分と比較する価値があります。この番組の基調講演枠には、NVIDIA の Jensen Huang、Qualcomm の Cristiano Amon、Intel の Lip-Bu Tan が参加し、それぞれが 2 時間を費やして、自分の言葉で、同じこと、つまり AI は完全にクラウド内で生きていくわけではないということについて議論しました。
Huang は、MediaTek で構築された新しい Windows PC スーパーチップである NVIDIA RTX Spark を発表しました。これは、NVLink-C2C を使用して Blackwell GPU ダイを Grace CPU ダイに接続し、両方にコヒーレント メモリを備えています。 「PCは再発明されつつある」と同氏は語った。同じ基調講演では、NVIDIA の Vera Rubin AI ファクトリー プラットフォームが本格的に稼働しました。賭けは明らかに両端にあります。クラウド規模のデータセンターと、連続体として機能するオンデバイスの PC シリコンです。
アモンのフレーミングはよりシャープでした。同氏は2026年を「エージェントの年」と宣言し、300ドルのエントリーレベルのWindows AI PCにはSnapdragon Cを、プレミアムPCにはSnapdragon X2 Eliteを導入し、クアルコムのクラウドシリコン市場への完全参入を示す最新のデータセンターAI推論チップラインであるDragonflyを発表した。同氏は、その基礎となる戦略を「コンピューティング連続体」、つまりデバイス、エッジ、クラウド全体に動的に割り当てられるワークロードと呼びました。 「エージェントはデバイスに関連付けられていません」と彼は言いました。 「ユーザーとともに動きます。」
Tan 氏の Intel 基調講演では、この事件の最も明白なバージョンが明らかになりました。 Perplexity CEO の Aravind Srinivas 氏のステージで、インテルは Core Ultra シリーズ 3 ラップトップでハイブリッド推論のライブ デモを実行しました。ローカル モデルは機密性の高いものにフラグを立て、それをデバイス上に保持し、機密性の低い素材のみをクラウドに送信しました。 Tan 氏はその後、「プライバシー、セキュリティ、コンプライアンス、コストがハイブリッド コンピューティングの必要性を高めている」とはっきりと述べました。
三社三雄

つまり、同じ引数の 3 つのフレーバーです。 AI がすべてクラウド上に存在するわけではありません。プライバシー、コスト、またはその両方のために、その一部はデバイス上に存在する必要があり、シリコンはそれを電話からワークステーションまでの範囲にわたって提供する必要があります。
これは基本的に、2017 年に A11 Bionic が Neural Engine を導入して以来、Apple が行ってきたことですが、細かい点では Apple が戦略を明らかにせずにそれを実行したということです。ブルックス氏は、「コンピューティング連続体」、「ハイブリッド コンピューティング」、または「AI ファクトリー」という言葉を使用しませんでした。彼の語彙は拡張性があり、バランスが取れており、統一されており、基礎的なものでした。これは Apple が A11 以来使用してきた言葉です。業界の残りの企業が言葉を再発明することで追いついたのか、それとも Apple の立場を全面的に採用することで追いついたのかは未解決の問題である。今週の Computex 週間でもはや問題となっているのは、オンデバイス AI が重要かどうかです。
基本は動いていない
Brooks 氏は、スケーラブルなアーキテクチャ、バランスの取れたアーキテクチャ、統合メモリ、ワットあたりのパフォーマンスへのこだわりという 4 つの基盤に繰り返し立ち戻りました。これらの言葉は、Apple Silicon の初期の頃から変わっていません。変化したのは、それぞれがどれだけ積極的にスケールされたかです。
特にニューラル エンジンは、電話機チップ上の単一ブロックから A シリーズと M シリーズで共有される機能に移行し、ブルックス氏の計算によると、数十億台の Apple デバイスが AI で高速化されています。フレームワークのストーリーは、2017 年の Core ML、そして今日の MLX オープンソース プロジェクト、Metal Performance Shaders、Foundation Models API、Apple Intelligence API と並行して進みました。これらはすべて、Neural Engine、CPU、または GPU の上に位置します。
これについて実際に何が難しいのか、何がチームを覚醒させているのかについて尋ねられたブルックス氏の答えは、テクノロジーというよりはむしろ規律の問題でした。
より多くのパフォーマンスとより多くのトランスをスローするのは本当に簡単です

制限することなく問題に対処します。電力効率を重視し続ける必要があります。私たちは、本来のパフォーマンスを維持するために、バッテリー寿命を大幅に犠牲にするつもりはありません。
これは、Apple のアプローチと、現在数キロワットの AI ワークステーションを歓迎している業界との間に静かに一線を画す返答だ。ブルックスが描いたように、Apple の限界はポケットの中にある携帯電話です。他のすべてはその制約から下方に設計されます。
統合メモリ: 効率性を重視して設計され、AI のために維持されます
Apple は、生成 AI が消費者の関心事になるずっと前にユニファイド メモリを選択しました。当初のケースは効率性、つまり CPU の大規模で低速なメモリ プールと GPU の小型で高速なメモリ プールの間での無駄なデータ シャッフルを排除することでした。 Brooks 氏が古いアーキテクチャについて説明したように、システムは「データのコピーを行ったり来たりするのに有意義な時間を費やしていました。これは実際にはユーザーの利益にはなりませんが、ある種の必要性がありました。」
ユニファイド メモリは、チップのすべての部分 (CPU、GPU、ニューラル エンジン) に対応できる単一のプールに置き換えました。この効率化がひそかに AI の利点となっていることに Apple がいつ気づいたかとの質問に対するブルックス氏の答えは、正確に言えば、その認識は事後だった、というものだった。 「統合メモリは現時点では非常に意味のあることだ」と彼は言う。 Apple がそれを再発明した後、それは当然の反応であり、彼は笑いました。
この単一のプールにより、各 Apple チップは、ワークロードに適したツールであるニューラル エンジン、GPU、CPU のいずれかを使用して、そのメモリ容量まで AI モデルを実行できます。コピーやドメインの譲渡はありません。 Brooks が言うように、ニューラル エンジンが結果を生成し、次に CPU がそれを必要とするとき、統合メモリがなければデータを移動することになります。それで、あなたはそうではありません。
Neural Accelerators: 最近のコンパウンダー
オンデバイス AI スタックの最新部分は Neural

アクセラレータ、そして彼らは、Apple のシリコン戦略が実際にどのように複合化するかについて最も興味深い話を伝えます。
ブルックスはその系譜を辿った。 Apple は、2019 年の A13 で初めて、専用の行列およびベクトル演算ユニットとしてニューラル アクセラレータを CPU に追加しました。彼の言葉を借りれば、彼らのスイートスポットは「音声認識や音声生成など、低遅延でバースト性の高い AI タスク」です。これらは、モデルを小さくする必要があり、答えが即座に得られる必要があるワークロードを処理します。
より大きな動きは昨年、Apple がすべての GPU コアにニューラル アクセラレータを搭載したときに起こりました。これは最初に iPhone の A19 Pro (iPhone 17 Pro で使用されているチップ) でヒットし、その後 M5 ファミリ全体に伝わり、この春には M5 Pro と M5 Max が搭載されました。その結果、Brooks 氏は次のように述べています。「GPU ドメインにおける AI パフォーマンスが大幅に向上しました。これは、LLM から画像生成まで、幅広い AI ワークロードにとって大きなメリットになりました。」
まとめると、この写真は、AI アクセラレーションが単一のブロック (ニューラル エンジン) ではなく、CPU のマトリックス ユニット、GPU のコアごとのアクセラレータ、およびニューラル エンジン自体の 3 つのドメインにまたがっているチップです。ポイントは選択です。デバイス上のさまざまなワークロードは、さまざまなシリコンに最適に配置されます。このチップは 3 つすべてを提供します。
実際にはどうなるか
Brooks 氏が要約で述べた事例は、すでに現場で使われているアプリを見るとより鮮明になります。 Appleは、ブルックス氏が説明したデバイス上のAIに大きく依存した製品を開発している4人の開発者からのコメントを共有した。
Good Snooze のオーナーである Jordi Bruin が提供する Whisper Transcription は、Neural Engine と GPU を使用してデバイス上で文字起こしワークフロー全体を実行します。最初のバージョンは 2023 年に出荷されましたが、そのときは M シリーズ MacBook で長文のオンデバイス文字起こしが実用化されたばかりでした。それ以来の進歩は素晴らしいものです

シリコンチームが指摘するのは、複合化です。
それ以来、モデルの改良、Core ML などの機械学習フレームワーク、M および A シリーズ チップのニューラル アクセラレータにより、文字起こしワークフローが大幅に高速化されました。ユーザーはリアルタイムの 300 倍以上の速度でファイルを文字に起こすことができるようになりました。つまり、1 時間のインタビューを iPhone または Mac 上で 15 秒以内にローカルで文字に起こすことができることになります。
Whisper Transcription はここ数年、私にとって頼りになるアプリであり、オーディオを扱う人にとっては、スピードと同じくらいプライバシーの問題も重要です。 Bruin氏の言葉を借りると、このアプリは「AppleシリコンとAppleフレームワークを利用して、デバイス上で文字起こしプロセス全体を実行するため、ユーザーは機密性の高い録音をリモートサーバーに送信せずにプライベートに保つことができる」という。 Apple Intelligence と Ollama や LM Studio などの他のローカル AI オプションを統合すると、下流機能 (チャット、概要、フォローアップ アクション) もローカルに維持できることになります。
CEO Swupnil Sahai が開発した SwingVision は、Apple がオンデバイス AI の「スポーツ トラッキング」ユースケースで言及したアプリです。 1 台のスマートフォンをテニスやピックルボールのプロ品質の統計および審判システムに変え、コートサイドのビデオからテレビ品質のハイライトを生成します。これは、携帯電話で実行する最も要求の厳しいワークロードの 1 つでもあります。1 秒あたり 60 フレームの 1080p 映像がリアルタイムで処理され、しばしば直射日光の当たるコート上で処理されます。
ボール追跡の精度を最大化するために、SwingVision の AI モデルは 1080p ビデオを 60 フレーム/秒で処理します。 Apple シリコンで動作する当社のモデルのおかげで、テニスやピックルボールのプレーヤーは、バッテリー寿命や過熱を心配することなく、世界中のどのコートでもリアルタイムのライン通話や音声フィードバックを体験できます。
最後の条項 — バッテリー寿命や過熱を心配する必要はありません

— これは、ブルックスから直接得られたものではない、iPhone 封筒の答えです。 SwingVisionがその存在証明です。
LM Studio は、最も広く使用されているローカル LLM デスクトップ アプリを実行し、最近その作成者である Adrien Grondin から Locally AI を取得し、同じモデル実行機能を iPhone と iPad にもたらします。彼らの経験は、Brooks が Core ML や Metal Performance Shaders と並んでリストしているオープンソース フレームワークである MLX を最も明確に支持しているものです。
MLX は信じられないほどパフォーマンスが高く、Apple シリコン向けに高度に最適化されており、ハードウェアとソフトウェアを最大限に活用する推論フレームワークを使用できると LM Studio の CEO、Yagil Burowski 氏は述べています。 iPhone と Mac でこのエクスペリエンスを実現することは「簡単」であり、まさに Brooks 氏が作成していた 1 回のビルドでどこでもスケールできるストーリーであると同氏は付け加えました。
Grondin のフレームはケースの iPad と iPhone 側にあります。 「LM Studio と Locally アプリは高度なモデルへのアクセスを民主化し、ユーザーが特定の LLM をデバイス上でローカルに簡単にダウンロードして実行できるようにします。Apple シリコンと MLX のおかげで、iPhone、iPad、Mac 全体で真にシームレスなオンデバイス AI エクスペリエンスを実現できました。」
Quoc Huy の iPad 初のメモ取りアプリである CollaNote は、同じ議論のマルチモーダル バージョンを作成します。ノート内の AI アシスタント、手書きにリンクされた講義の文字起こし

[切り捨てられた]

## Original Extract

Apple's senior PM for Apple Silicon on the Neural Engine, unified memory, Fusion, and why the on-device AI bet hasn't moved in nine years.

Sign Up
Published Jun 05, 2026
Updated Jun 05, 2026
Apple Silicon's bet on on-device AI hasn't moved — only the chip range that runs it
Days before WWDC, Apple's silicon team's argument: the on-device AI bet hasn't shifted in nine years — just compounded across the range.
Three days before WWDC opens in Cupertino, and at the end of a Computex week in which NVIDIA, Qualcomm and Intel all took turns on the Taipei stages arguing — in their own words — for hybrid AI between cloud and device, Apple's silicon team is making a quieter point: their bet on running it on the device hasn't really moved.
The Neural Engine has been in Apple chips since 2017. Unified memory has been a foundational design choice since the first A-series part. What's changed in nine years isn't the philosophy — it's the silicon range that now delivers it, which today stretches from a $599 MacBook Neo to an M5 Max workstation chip with two dies bonded into a single package.
That, in essence, was Doug Brooks's argument in a briefing on the sidelines of Computex. Brooks is the senior product manager for Apple Silicon, and the conversation was on-the-record and product-focused — no demo, no marketing slide deck, no roadmap. Just the chip.
His opening line is worth quoting, because it's the spine of everything that followed:
If you want to build a great device for AI, you need to build a great computer.
You can read that as evasion — Apple deflecting the AI conversation back to the hardware it already makes — or you can read it as the actual thesis: on-device AI doesn't get bolted on to a phone or a Mac; it gets designed in from the silicon up, and Apple has been doing exactly that since the A11 Bionic introduced the Neural Engine in 2017. As Brooks put it, with a flicker of dryness, "well before the AI PC trend broke."
The same week, three other Silicon CEOs took the stage
Brooks's case is worth setting against the rest of Computex . The show's keynote slots belonged to NVIDIA's Jensen Huang, Qualcomm's Cristiano Amon and Intel's Lip-Bu Tan, and each of them spent two hours arguing — in their own words — for the same thing: that AI is not going to live entirely in the cloud.
Huang announced the NVIDIA RTX Spark , a new Windows PC superchip built with MediaTek that uses NVLink-C2C to connect a Blackwell GPU die to a Grace CPU die with coherent memory across both. "The PC is being reinvented," he said. The same keynote put NVIDIA's Vera Rubin AI-factory platform into full production. The bet is plainly on both ends — cloud-scale data centres and on-device PC silicon, working as a continuum.
Amon's framing was sharper. He declared 2026 the "Year of the Agent," introduced Snapdragon C for $300 entry-level Windows AI PCs and Snapdragon X2 Elite for premium ones, and unveiled Dragonfly, a brand-new data-centre AI inference chip line that marks Qualcomm's full entry into the cloud silicon market. He called the underlying strategy the "Computing Continuum" — workloads dynamically allocated across devices, the edge, and the cloud. "The agent isn't tied to the device," he said. "It moves with the user."
Tan's Intel keynote landed the most explicit version of the case. On stage with Perplexity CEO Aravind Srinivas, Intel ran a live hybrid-inference demo on a Core Ultra Series 3 laptop: the local model flagged what was sensitive, kept that on the device, and sent only non-sensitive material to the cloud. Tan put it plainly afterwards — "privacy, security, compliance and cost are driving the need for hybrid compute."
Three companies, three stages, three flavours of the same argument. AI is not all going to live in the cloud. Some of it has to live on the device, for privacy or for cost or for both, and the silicon has to deliver that across a range from a phone to a workstation.
That is essentially the case Apple has been making since the A11 Bionic introduced the Neural Engine in 2017, with the small detail that Apple did it without naming a strategy. Brooks did not use the phrases "Computing Continuum" or "hybrid compute" or "AI factories." His vocabulary was scalable, balanced, unified, and foundational — words Apple has used since the A11. Whether the rest of the industry has caught up by reinventing language or by adopting Apple's positions wholesale is the open question. What is no longer the question this Computex week is whether on-device AI matters.
The fundamentals haven't moved
Brooks returned repeatedly to four foundations: a scalable architecture, a balanced architecture, unified memory, and an insistence on performance per watt. None of those words have changed since the early days of Apple Silicon. What has changed is how aggressively each has been scaled.
The Neural Engine, in particular, has gone from a single block on a phone chip to a feature shared across A-series and M-series, with billions of Apple devices now AI-accelerated by Brooks's count. The framework story rode along in parallel — Core ML in 2017, then today's MLX open-source project, Metal Performance Shaders, the Foundation Models API and the Apple Intelligence APIs. They all sit atop the Neural Engine, the CPU, or the GPU.
Asked about what's actually hard about this — what keeps the team awake — Brooks's answer was less about technology and more about discipline.
It's really easy to throw more performance and more transistors at problems without bounding it. You have to continue to keep that focus on power efficiency. We're not willing to sacrifice a huge amount of battery life for raw performance on its own.
It's a reply that quietly draws a line between Apple's approach and an industry currently celebrating multi-kilowatt AI workstations. As Brooks framed it, Apple's bound is the phone in your pocket. Everything else is engineered downwards from that constraint.
Unified memory: designed for efficiency, kept for AI
Apple chose unified memory long before generative AI was a consumer concern. The original case was efficiency — eliminating wasteful data shuffling between a CPU's large, slow memory pool and a GPU's smaller, faster one. As Brooks described the old architecture, the system spent meaningful time "copying data back and forth, which really doesn't benefit the user, but it was kind of a necessity."
Unified memory replaced that with a single pool every part of the chip can address: CPU, GPU, Neural Engine. When asked at what point Apple realised this efficiency play had quietly become an AI advantage, Brooks's answer was, accurately, that the realisation came after the fact. "Unified memory just makes so much sense right now," he said. After Apple re-invented it, it was the obvious response, and he laughed.
That single pool is what lets each Apple chip run an AI model up to its memory capacity, with whichever piece of silicon — Neural Engine, GPU, or CPU — is the right tool for the workload. No copies, no domain transfers. As Brooks put it, when the Neural Engine produces a result and the CPU needs it next, without unified memory, you're moving data around; with it, you're not.
Neural Accelerators: the recent compounder
The newest piece of the on-device AI stack is Neural Accelerators, and they tell the most interesting story about how Apple's silicon strategy actually compounds.
Brooks traced the lineage. Apple first added Neural Accelerators to the CPU in 2019 with the A13, as dedicated matrix and vector math units. Their sweet spot, in his words, is "low-latency, bursty AI tasks — things like speech recognition, speech generation." They handle the workloads where the model needs to be small, and the answer needs to be instant.
The bigger move came last year, when Apple put Neural Accelerators on every GPU core. That hit first on the A19 Pro in iPhone — the chip used in the iPhone 17 Pro — and then propagated through the M5 family, with M5 Pro and M5 Max picking them up this spring. The result, per Brooks: "a huge boost in AI performance in the GPU domain. That's been a huge benefit for a wide range of AI workloads, from LLMs to image generation."
Taken together, the picture is a chip in which AI acceleration is no longer a single block — the Neural Engine — but spread across three domains: the CPU's matrix units, the GPU's per-core accelerators, and the Neural Engine itself. The point is choice. Different on-device workloads land best on different silicon. The chip provides all three.
What that looks like in practice
The case Brooks made in the abstract becomes sharper when you look at the apps already in the field. Apple shared a comment from four developers whose products lean heavily on the on-device AI Brooks described.
Whisper Transcription , from Jordi Bruin, owner of Good Snooze, runs the entire transcription workflow on-device using the Neural Engine and GPU. The first version shipped in 2023, when M-series MacBooks had only just made long-form on-device transcription practical. The progress since then is the kind of compounding the silicon team would point to.
Model improvements, machine learning frameworks such as Core ML, and the Neural Accelerators in the M and A-series chips sped up the transcription workflow enormously since then. Users can now transcribe files at more than 300 times real-time speed, meaning an hour-long interview can be transcribed locally in under 15 seconds right on their iPhone or Mac.
Whisper Transcription has been my go-to app for a couple of years, and for anyone who works with audio, the privacy story matters as much as the speed. The app, in Bruin's words, "takes advantage of Apple silicon and Apple frameworks to run the whole transcription process on device, so users can keep sensitive recordings private instead of sending them to a remote server." Apple Intelligence integration, alongside other local AI options like Ollama and LM Studio, means downstream features — chat, summaries, follow-up actions — can also stay local.
SwingVision , from CEO Swupnil Sahai, is the app Apple mentioned in its "sports tracking" on-device AI use case. It turns a single smartphone into a pro-quality stats and officiating system for tennis and pickleball, generating TV-quality highlights from court-side video. It is also one of the more demanding workloads anyone runs on a phone: 1080p footage at 60 frames per second, processed in real time, often on a court in direct sun.
To maximise ball tracking accuracy, SwingVision's AI models process 1080p video at 60 frames per second. Thanks to our models running on Apple silicon, tennis and pickleball players can experience real-time line calling and audio feedback on any court in the world — all without worrying about battery life or overheating.
That last clause — without worrying about battery life or overheating — is the iPhone-envelope answer we did not get directly from Brooks. SwingVision is the existence proof.
LM Studio runs the most widely used local-LLM desktop app and recently acquired Locally AI from its creator, Adrien Grondin, bringing the same model-running capability to iPhone and iPad. Their experience is the cleanest endorsement of MLX, the open-source framework Brooks listed alongside Core ML and Metal Performance Shaders.
MLX is incredibly performant and highly optimised for Apple silicon, and we get to work with an inference framework that leverages the hardware and software best, said LM Studio CEO Yagil Burowski. Bringing the experience across iPhone and Mac, he added, was "straightforward" — exactly the build-once, scale-everywhere story Brooks was making.
Grondin's framing gets at the iPad-and-iPhone side of the case. "LM Studio and the Locally app democratise access to advanced models, enabling users to easily download and run specific LLMs locally on device. Apple silicon and MLX allowed us to enable a truly seamless on-device AI experience across iPhone, iPad and Mac."
CollaNote , Quoc Huy's iPad-first note-taking app, makes the multimodal version of the same argument. AI assistants inside notes, lecture transcription linked to handwrit

[truncated]
