---
source: "https://www.darktable.org/2026/06/meet-darktable-5.6-ai-tools/"
hn_url: "https://news.ycombinator.com/item?id=48655147"
title: "New AI tools in darktable 5.6"
article_title: "Meet the new AI tools in darktable 5.6 | darktable"
author: "mikae1"
captured_at: "2026-06-24T05:18:09Z"
capture_tool: "hn-digest"
hn_id: 48655147
score: 1
comments: 0
posted_at: "2026-06-24T04:32:20Z"
tags:
  - hacker-news
  - translated
---

# New AI tools in darktable 5.6

- HN: [48655147](https://news.ycombinator.com/item?id=48655147)
- Source: [www.darktable.org](https://www.darktable.org/2026/06/meet-darktable-5.6-ai-tools/)
- Score: 1
- Comments: 0
- Posted: 2026-06-24T04:32:20Z

## Translation

タイトル: darktable 5.6 の新しい AI ツール
記事のタイトル: darktable 5.6 の新しい AI ツールを紹介 |ダークテーブル
説明: darktable 5.6 には、最初の AI 機能が搭載されています。1 回のクリックであらゆる被写体の周囲のベクター マスクに変換する AI オブジェクト マスクと、ML ベースのノイズ除去とアップスケールのためのニューラル復元モジュールです。私たちは、オープン コンピューター ビジョン モデルを CPU 上でローカルにロードして実行する基盤を構築することに重点を置きました。
[切り捨てられた]

記事本文:
Darktable 5.6 の新しい AI ツールを紹介します
darktable 5.6 には、最初の AI 機能が搭載されています。1 回のクリックであらゆる被写体の周囲のベクター マスクに変換する AI オブジェクト マスクと、ML ベースのノイズ除去とアップスケールのためのニューラル復元モジュールです。私たちは、オープン コンピューター ビジョン モデルを CPU または GPU 上でローカルにロードして実行する基盤を構築することに重点を置きました。そのため、オープン モデルの状況が進化するにつれて、darktable は写真撮影に適した AI タスクをホストできるようになります。この投稿は、各ツールの機能、使用方法、および不十分な点を説明する実践ツアーです。
手順を進めるために、ウォークスルーで使用される 5 つのサンプル RAW を入手できます。
AI オブジェクト マスクのウォークスルーに使用される「Ebby_ND7_5751.NEF という名前の猫」(24 MB) をダウンロードします。
「エビーという名前の猫」© 2020 by Suki2019 、CC BY-SA 4.0 に基づいてライセンス供与されています。
「DSC01318.ARW」（20 MB）をダウンロード – マルチオブジェクト マスクの例に使用されます。
「ペンギンがいっぱい！」 © 2025 by Klogg、CC BY-SA 4.0 に基づいてライセンス許諾されています。
タイトル画像に使用される「_DSC0488.NEF」（15 MB）をダウンロードします。
「菊の花」© 2019 by Xavier Bartol (XavAL)、CC BY-SA 4.0 に基づいてライセンス許諾されています。
「07 Monkey ISO 6400.CR3」(43 MB) をダウンロードします – ニューラルノイズ除去のウォークスルーに使用されます。
「生のノイズ除去の相互比較」© 2025 by Dave22152、CC BY-SA 4.0 に基づいてライセンス供与されています。
「20250629_0012.NEF」（13 MB）をダウンロード – 高級ウォークスルーに使用されます。
「花の上のハナアブ」© 2025 by Fenny、ライセンスは CC BY-SA 4.0 に基づいています。
これらを共有してくれた Suki、Klogg、Xavier、Dave、Fenny に感謝します。
私たちが早い段階で設定したいくつかの原則は、その後のすべての決定を形作ることになりました。
オプション。 USE_AI を使用せずに構築すると、サブシステム全体が消えます。公式バイナリにはこれが含まれていますが、ディストリビューションはオプトアウトでき、あなたも同様にオプトアウトできます。既存のモジュールはどれもこれに依存しません。
地元。何もありません

あなたのマシンは大丈夫です。クラウド推論やテレメトリはなく、「データを使用してモデルを改善します」ということもありません。すべてのモデルは CPU 上で実行されます。GPU がある場合は GPU 上で実行されます。私自身写真家として、生のファイルをブラックボックス サービスに貼り付けることは決してありませんし、他の人にそれを頼むつもりもありませんでした。
厳選されたオープンモデル。私たちが提供するすべてのモデルは、ソース コード、トレーニング データ、ライセンスとともに別のリポジトリ (darktable-ai) に文書化されています。バイナリの重みは darktable 自体の中にバンドルされておらず、出所が明確でないモデルは含まれません。このバーは、当社のカタログが競合他社よりも小さいことを意味します。それが私たちが選んだ取引です。
拡張可能。カタログは出発点であり、柵ではありません。 darktable の各 AI タスク (オブジェクト マスキング、ノイズ除去、アップスケールなど) には、入力シェイプ、出力シェイプ、予期される前処理および後処理など、定義された ONNX インターフェイスがあります。そのインターフェースに一致するモデルはすべて手動でインストールでき、ファーストパーティのものと同様に使用できます。公開されたモデルを ONNX に変換する方法、または独自のモデルをトレーニングする方法を知っている場合は、出荷を待たずにそれを darktable に持ち込むことができます。カタログはすべての人のためのものです。さらに先に進みたければ、ドアは開いています。
デフォルトではオフです。有効にした機能で必要になるまで、何もダウンロードされません。
AI 機能は、ユーザーがオンにするまでオフのままです。ユーザーが要求しない限り、何もインストールまたはロードされません。 [設定] → [AI] を開き、[AI 機能を有効にする] にチェックを入れます。最も簡単に開始できるのは、ダウンロード/更新のデフォルト アクションです。ワンクリックで推奨セットがプルダウンされるため、すべての機能がすぐに動作します。インストールされると、モデルはディスク上に存在し、オンデマンドでロードされます。
GPU をお持ちの場合は、対応する実行プロバイダー (macOS の CoreML、NVIDIA の CUDA、AMD の ROCm、Windows の DirectML、Intel の OpenVINO) を接続します。セットアップはプラットフォームごとに異なります。 GPU アクセラレーション PA

マニュアルでは、それぞれについて説明します。 CPUも動作しますが、速度が遅くなります。
これは私のワークフローを最も変えた機能です。ダークテーブルで描画されたマスクは、ブラシ、パス、グラデーション、楕円、パラメトリック交差など、常に強力でしたが、たとえば、混雑した背景に対して動物をマスキングするには、常にブラシでの忍耐か、近似に落ち着くかのどちらかが必要でした。
AI オブジェクト マスクはそれをショートカットします。オブジェクトをクリックすると、そのオブジェクトの正確でハードエッジな選択が得られ、その場で通常のベクトル パスに変換されます。描画されたマスクと同じ種類で、移動できるエッジ ノード、調整できるフェザー、他のマスクと組み合わせることができる交差を備えた描画マスクが生成されます。
暗室で Ebby_ND7_5751.NEF という名前の Cat を開きます。マスキングをサポートするモジュールを選択します。デモではカラー バランス rgb を使用します。マスク マネージャーを開き、AI オブジェクト マスク ツールを選択します。
新しい画像を最初にクリックすると、エンコーダ パスがトリガーされます (ハードウェアに応じて 1 ～ 10 秒)。それが完了すると、この画像上に留まる限り、モデルはクリック プロンプトにインタラクティブに応答できるようになります。
Ebby の中心付近をクリックします。モデルは選択範囲を返し、キャンバス上に重ねて表示される閉じたパスにベクトル化されます。右クリックしてコミットします。
コミットする前に、2 つのコントロールによって結果が形成されます。
スムージングは、ベクトル化によって生成されるアンカー ポイントの数を制御します。デフォルトは 0 – アンカーの数が最大で、パスはシルエットにぴったりと密着します (編集が忙しくなります)。柔らかいエッジの被写体 (雲、花びら) の場合はこの値を上げると、アンカーの数が少なく、より滑らかな曲線が得られますが、きつい曲がり角では角を切ります。範囲は最大 1.3 です。
[境界を調整] は、エッジに焦点を当てた 2 番目のニューラル パスを実行するチェックボックスです。ほとんどの場合、目に見える画像の輪郭に対してパスをよりしっかりと引き込みます。

細い尻尾や突き出た耳など、最初のマスクで切り取られた細かい部分を表現するのに役立ちます。 (これは毛皮や髪を処理するものではありません。以下のフェザリング トリックはそれを行いますが、それは別の問題です。) クリックごとに 1 ～ 2 秒余分にコストがかかりますが、細かいシルエットのディテールがたくさんある被写体では、ほとんどの場合、それだけの価値があります。
最初の選択が完全に正しくない場合 (耳を外した、テーブルの一部をつかんだなど)、やり直しではなく追加のクリックで調整します。単純にクリックすると、別の肯定的なプロンプト (「これも含めてください」) が追加されます。 Shift キーを押しながらクリックすると、否定的なプロンプト (「これではありません」) が追加されます。エンコーダを再実行する必要はありません。各調整クリックはインタラクティブです。
見逃した部分を追加するためのあまり明白ではないヒント (たとえば、選択されていない耳) は、選択されていない領域の境界線のすぐ横をクリックしないでください。この種のエッジ隣接クリックはモデルを混乱させます (どちら側を指しているのかは曖昧です)。代わりに、欠落している部分に近いすでに選択されている領域をクリックします。モデルはそれを「この方向に選択範囲を拡張する」と読み取り、通常は耳をきれいに掴みます。
コミットされると、マスクは通常のパスになります。ノードを移動したり、パラメトリック マスク (輝度または色による制限に適しています) と交差させたり、標準のマスク コントロールでエッジをフェザリングしたりすることができます。
フェザリングによる毛皮のカバー範囲の向上
境界線の調整をオンにしても、エビーのコートのうっすらとしたエッジは、ハードエッジのベクター パスにきれいに変換されません。私が使うコツは、道端で戦わないことです。ブレンド面を柔らかくします。
マスクを適用した同じモジュールで、ブレンド オプションまで下にスクロールします。 [マスクの調整] で、フェザーの半径を約 5 ～ 10 ピクセルまで上げます。これにより、ハード マスクのエッジが文字通りのベクトル ラインではなく、画像のローカル コントラストに従います。毛皮で覆われた主題について

基礎となるパスにストランドごとの詳細がない場合でも、結果は自然なものとなります。
これは、ほぼすべての AI マスクで髪、毛皮、葉に使用するテクニックです。フェザー半径を小さくすると、手作業でエッジをブラシ修正するよりもはるかに小さな介入で済みます。
複数の被写体: ペンギン
ここで DSC01318.ARW を開きます。背景にペンギンのコロニーがあり、前景にペンギンのコロニーが他のペンギンの数メートル前に立っています。これは、クリック 1 つですべてが完了するという思い込みを打ち破る、一種の複数の被写体のシーンです。
AI オブジェクト マスク ツールを選択し、前景のペンギンをクリックします。モデルは、その 1 匹のペンギンだけをカバーするマスクを返します。ここで、肯定的なプロンプトとして背景のペンギンの 1 つをクリックして選択範囲を拡張してみます。前景のペンギンのマスクはそれを含むまで拡大しません。どれだけポジティブなクリックを追加しても、視覚的に分離された被写体が同じマスクに収まることはありません。各 AI オブジェクト マスクは、構造上 1 つの接続された領域です。
モデルの動作は被験者が接触するかどうかに依存するため、明確化します。物理的に接触しているペンギンの密集したクラスター内をクリックすると、モデルはペンギンを 1 つの接続された領域として扱うため、複数のペンギンを一度に掴む可能性があります。しかし、モデルに関する限り、前景のペンギンとコロニーは、それらの間に明確なギャップがあり、分離されています。
両方をマスクするには、右クリックして最初のマスクをコミットし、新しい AI オブジェクト マスクを開始して、2 番目のサブジェクトをクリックします。別の主題の数だけ繰り返します。
マスク マネージャーでは、デフォルトでマスクがユニオン モードでレイヤー化されます。すべて選択すると、モジュールが動作する結合領域が形成されます。
要点: AI は一度に 1 つの関連性のあるものを見つけます。ワンクリックで被写体にタッチできます。別々の被験者にはそれぞれマスクが 1 枚必要です。あなたはその人です

どのようなものを決定し、それらを組み合わせるオーケストレーター。
それは魔法ではありません。正直な注意点がいくつかあります:
その下はベクターマスクです。 AI ツールは領域を見つけて、それをベクトル パス (描画されたパス マスクで使用されるのと同じ種類) に変換します。強力ですが、ピクセルごとの詳細を表現することはできません。うっすらと生えた髪の毛、動物の毛皮、葉のエッジなどは、たとえモデルがどれほど優れていても、ピクセルを完璧に再現することはできません。上記の微調整テクニック (ブレンド時のマスク微調整の下で境界を微調整 + ぼかし半径を使用) は非常に役立ちますが、完全に解決するわけではありません。ベクトル マスクは、ストランドごとの精度にとって間違った抽象化です。
曖昧さは本物です。ギターを持っている人をクリックすると、クリックした場所とロードされているモデルに応じて、「ギターを持っている人」が 1 つのオブジェクトとして表示される場合もあれば、単に「ギター」が表示される場合もあります。 2 回目のクリックでほぼ毎回解決しますが、最初の試行で常に期待どおりになるとは限りません。
新しい画像の最初の選択が遅い。エンコーダーは画像ごとに 1 回 (ハードウェアに応じて 1 ～ 10 秒) 実行され、その後キャッシュされます。同じ画像に対するその後の選択と調整はインタラクティブに行われます。それに応じて計画を立ててください。ピンポンではなく、先に進む前に画像を完成させてください。
描画されたツールがまだ勝つ場合
ブラシとパスがどこにも進まないのには十分な理由があります。
オブジェクトではない選択範囲。空の隅、前景のくさび、「光があるところならどこでも」 - グラデーションとパスを使用すると、これがより速く行われます。
ピクセルパーフェクトな手動制御。マスクのエッジを特定のラインに配置する必要がある場合、ブラシが最適なツールです。
細かい詳細。セグメンテーション モデルは、マスクしたい場合がある種類の微細要素に基づいてトレーニングされていません。モデルと格闘するよりも、描いた方が早いです。
維持されている経験則: 何をマスキングしているのか説明できれば

単一の英語の名詞 (「猫」、「ボトル」、「自転車」) を使用する場合、AI オブジェクトのマスクが最速のパスです。できないなら描いてください。
2 番目の機能はニューラル復元です。これはライトテーブル (デフォルトでは右側のサイドバー) と暗室 (左側のサイドバー) の両方で利用できるユーティリティ モジュールです。ユーザーマニュアルには完全なリファレンスが記載されています。以下は実践的なツアーです。
オブジェクト マスクとは異なり、これはピクセル パイプラインの一部ではありません。入力画像を取得し、その上でニューラル ネットワークを実行し、結果を新しい画像としてライトテーブルに書き込みます。 raw denoise 、 denoise 、および upscale の 3 つのタスクがパネルを共有します。
これらのツールは、パイプライン モジュールとは異なる役割を果たします。速度が遅く、出力は常に新しい画像になるため、素早い色調調整には適していません。彼らが得意なことに関しては、素晴らしいです。
ウォークスルー: ISO 6400 モンキーのノイズ除去
暗室で07 Monkey ISO 6400.CR3を開きます。 100% にズームインすると、滑らかな表面 (猿の体、灰色の背景) と影の中にノイズが見えます。これは、カメラ センサーが高 ISO で生成するものです。ノイズは本物ですが、根底にあるディテールはまだ回復する必要があります。
ニューラル復元では、タスク = denoise およびモデル = denoise-nind (汎用 UNet) を選択します。強度スライダーを使用すると、エフェクトが強すぎる場合にエフェクトを元に戻すことができます。デフォルトはほとんどのファイルで機能します。結果が滑らかすぎる場合は、60 ～ 70% に下げます。 「プロセス」をクリックします。 GPU では数秒です。 CPU 例で

[切り捨てられた]

## Original Extract

darktable 5.6 ships its first AI features: an AI object mask that turns a single click into a vector mask around any subject, and a neural restore module for ML-based denoise and upscale. Our focus was on building the foundation that loads and runs open computer-vision models locally on your CPU or
[truncated]

Meet the new AI tools in darktable 5.6
darktable 5.6 ships its first AI features: an AI object mask that turns a single click into a vector mask around any subject, and a neural restore module for ML-based denoise and upscale. Our focus was on building the foundation that loads and runs open computer-vision models locally on your CPU or GPU – so darktable can host AI tasks suitable for photography as the open model landscape evolves. This post is the hands-on tour: what each tool does, how to use it, and where it falls short.
To follow along, you can grab the five sample raws used in the walkthroughs:
Download “a Cat named Ebby_ND7_5751.NEF” (24 MB) – used for the AI object mask walkthrough.
“A cat named Ebby” © 2020 by Suki2019 , licensed under CC BY-SA 4.0 .
Download “DSC01318.ARW” (20 MB) – used for the multi-object mask example.
“A whole lot of penguins!” © 2025 by Klogg , licensed under CC BY-SA 4.0 .
Download “_DSC0488.NEF” (15 MB) – used for the title image.
“Chrysanthemum Flowers” © 2019 by Xavier Bartol ( XavAL ), licensed under CC BY-SA 4.0 .
Download “07 Monkey ISO 6400.CR3” (43 MB) – used for the neural denoise walkthrough.
“A Raw Denoise Cross-comparison” © 2025 by Dave22152 , licensed under CC BY-SA 4.0 .
Download “20250629_0012.NEF” (13 MB) – used for the upscale walkthrough.
“Hoverfly on flower” © 2025 by Fenny , licensed under CC BY-SA 4.0 .
Thanks to Suki, Klogg, Xavier, Dave, and Fenny for sharing these.
A few principles we set early on, that ended up shaping every later decision:
Optional. Built without USE_AI , the whole subsystem disappears. The official binaries include it, but distros can opt out, and so can you – none of the existing modules depend on it.
Local. Nothing leaves your machine. No cloud inference, no telemetry, no “we’ll improve the model with your data”. Every model runs on your CPU, or your GPU if you have one. As a photographer myself, I’d never paste a raw file into a black-box service, and I wasn’t going to ask anyone else to.
Curated open models. Every model we offer is documented in a separate repository – darktable-ai – with its source code, training data and license. Binary weights are not bundled inside darktable itself, and models whose provenance isn’t clear don’t get included. That bar means our catalog is smaller than commercial competitors. That’s the trade we chose.
Extensible. The catalog is a starting point, not a fence. Each AI task in darktable (object masking, denoise, upscale, …) has a defined ONNX interface – input shapes, output shapes, expected pre- and post-processing. Any model matching that interface can be installed manually and used like a first-party one. If you know how to convert a published model to ONNX, or train your own, you can bring it to darktable without waiting for us to ship it. The catalog is for everyone; the door is open if you want to go further.
Off by default. Nothing is downloaded until a feature you enable needs it.
AI features stay off until you turn them on – nothing is installed or loaded unless you ask for it. Open preferences → AI and tick enable AI features . The easiest start is the download / update default action: one click pulls down the recommended set so every feature works immediately. Once installed, models live on disk and load on demand.
If you have a GPU, wire up a matching execution provider – CoreML on macOS, CUDA on NVIDIA, ROCm on AMD, DirectML on Windows, OpenVINO for Intel. Setup varies per platform; the GPU acceleration page in the manual walks through each one. CPU works too, just slower.
This is the feature that’s changed my workflow the most. Drawn masks in darktable have always been powerful – brush, path, gradient, ellipse, parametric intersection – but masking, say, an animal against a busy background has always meant either patience with the brush or settling for an approximation.
The AI object mask short-cuts that. A click on an object gives you an accurate, hard-edged selection of that object, converted on the fly to a regular vector path – the same kind drawn masks produce, with edge nodes you can move, feather you can adjust, and intersections you can combine with other masks.
Open a Cat named Ebby_ND7_5751.NEF in the darkroom. Pick any module that supports masking – I’ll use color balance rgb for the demo. Open its mask manager and pick the AI object mask tool.
The first click on a new image triggers the encoder pass – 1 to 10 seconds depending on hardware. Once it’s done, the model is ready to answer click prompts interactively for as long as you stay on this image.
Click somewhere near the centre of Ebby. The model returns a selection, vectorised into a closed path you can see overlaid on the canvas. Right-click to commit it.
Two controls shape the result before you commit:
Smoothing controls how many anchor points the vectorisation produces. The default is 0 – maximum anchors, path hugs the silhouette closely (and is busier to edit). Raise it for soft-edged subjects (clouds, flower petals) to get a smoother curve with fewer anchors, at the cost of cutting corners on tight bends. The range goes up to 1.3.
Refine boundary is a checkbox that runs a second neural pass focused on the edge. It pulls the path tighter against the visible image contours – mostly helps with small details the initial mask cut off, like a thin tail or a sticking-up ear. (It’s not what handles fur or hair – the feathering trick below does that, and it’s a different problem.) Costs an extra second or two per click, but on subjects with lots of fine silhouette detail it’s almost always worth it.
If the first selection isn’t quite right – say it missed an ear or grabbed a chunk of the table – refine with extra clicks rather than restarting. A plain click adds another positive prompt (“include this too”); shift-click adds a negative prompt (“not this”). The encoder doesn’t need to re-run; each refining click is interactive.
One non-obvious tip for adding a missed part – say an unselected ear: don’t click right next to the boundary on the unselected area. That kind of edge-adjacent click confuses the model (it’s ambiguous which side you mean). Click on an already-selected area close to the missed part instead. The model reads that as “extend the selection in this direction” and usually grabs the ear cleanly.
Once committed, the mask is a normal path. You can move its nodes, intersect it with parametric masks (good for restricting by luminance or colour), or feather its edges in the standard mask controls.
Better fur coverage with feathering
Even with refine boundary on, the wispy edge of Ebby’s coat doesn’t translate cleanly to a hard-edged vector path. The trick I use: don’t fight it on the path side; soften it on the blend side.
In the same module where you’ve applied the mask, scroll down to the blending options. Under mask refinement , bump the feather radius up to about 5–10 pixels. This makes the hard mask edge follow the image’s local contrast instead of the literal vector line. On a furry subject the result reads natural even though the underlying path has no per-strand detail.
This is the technique I use for hair, fur, and foliage in almost every AI mask. A small feather radius is a much smaller intervention than trying to brush-correct the edge by hand.
Multiple subjects: the penguins
Now open DSC01318.ARW . A colony of penguins in the background, with one in the foreground standing a few metres in front of the others – the kind of multi-subject scene that breaks the assumption that one click does it all.
Pick the AI object mask tool and click on the foreground penguin. The model returns a mask covering just that one penguin. Now try to extend the selection to one of the background penguins by clicking on it as a positive prompt – the foreground penguin’s mask doesn’t grow to include it. Visually-separated subjects can’t end up in the same mask, no matter how many positive clicks you add. Each AI object mask is one connected region by construction.
A clarification, since the model’s behaviour depends on whether subjects touch: clicking inside a tight cluster of penguins that are physically touching may grab several of them at once, because the model treats them as one connected region. But the foreground penguin and the colony, with a clear gap between them, are separate as far as the model is concerned.
To mask both, commit the first mask with right-click, start a new AI object mask, and click the second subject. Repeat for as many separate subjects as you have.
In the mask manager, the masks layer together by default in union mode – select them all and they form a combined region your module operates on.
The takeaway: the AI finds one connected thing at a time. Touching subjects can be one click. Separated subjects need one mask each. You’re the orchestrator who decides which things and combines them.
It’s not magic. A few honest caveats:
It’s a vector mask underneath. The AI tool finds a region for you and turns it into a vector path – the same kind a drawn path mask uses. Powerful, but it can’t represent per-pixel detail. Wispy hair, animal fur, foliage edges – these will never come out pixel-perfect, no matter how good the model. The refinement techniques shown above ( refine boundary + feather radius under mask refinement in blending) help a lot, but they don’t solve it completely. A vector mask is the wrong abstraction for strand-by-strand precision.
Ambiguity is real. Clicking on a person holding a guitar might give you “person holding guitar” as one object, or just “guitar”, depending on where you clicked and which model is loaded. A second click resolves it almost every time – but the first attempt isn’t always what you expect.
First selection on a new image is slow. The encoder runs once per image – 1 to 10 seconds depending on hardware – then caches. Subsequent selections and refinements on the same image are interactive. Plan accordingly: finish an image before moving on, rather than ping-ponging.
When the drawn tools still win
The brush and path aren’t going anywhere, and for good reason:
Selections that aren’t objects. A corner of the sky, a wedge of foreground, “wherever the light is” – gradients and paths do this faster.
Pixel-perfect manual control. When the mask edge has to land on a specific line, the brush is the right tool.
Tiny details. Segmentation models weren’t trained on the kinds of micro-elements you sometimes want to mask. Drawing is faster than fighting the model.
A rule of thumb that’s held up: if you can describe what you’re masking in a single English noun (“the cat”, “the bottle”, “the bike”), the AI object mask is the fastest path. If you can’t, draw it.
The second feature is neural restore – a utility module available in both lighttable (right sidebar by default) and darkroom (left sidebar). The user manual has the full reference ; below is a hands-on tour.
Unlike the object mask, this isn’t part of the pixel pipeline. It takes an input image, runs a neural network over it, and writes the result back to the lighttable as a new image. Three tasks share the panel: raw denoise , denoise , and upscale .
These tools sit in a different role from the pipeline modules. They’re slower, and the output is always a new image – so they’re a poor fit for a quick tonal tweak. For the things they’re good at, they’re remarkable.
Walkthrough: denoising the ISO 6400 monkey
Open 07 Monkey ISO 6400.CR3 in the darkroom. Zoom in to 100% – you’ll see the noise across the smooth surfaces (the monkey’s body, the grey background) and in the shadows. This is what a camera sensor produces at high ISO; the noise is real but the underlying detail is still there to recover.
In neural restore, pick task = denoise and model = denoise-nind (the general-purpose UNet). The strength slider lets you dial back the effect if it’s too aggressive – default works for most files; drop to 60–70% if the result is too smooth. Click process . On a GPU it’s a few seconds; on CPU ex

[truncated]
