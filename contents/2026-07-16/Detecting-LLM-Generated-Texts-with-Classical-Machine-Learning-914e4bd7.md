---
source: "https://blog.lyc8503.net/en/post/llm-classifier/"
hn_url: "https://news.ycombinator.com/item?id=48936880"
title: "Detecting LLM-Generated Texts with \"Classical\" Machine Learning"
article_title: "Detecting LLM-Generated Web Fiction with \"Classical\" Machine Learning (AIGC Text Detection)"
author: "uneven9434"
captured_at: "2026-07-16T17:04:29Z"
capture_tool: "hn-digest"
hn_id: 48936880
score: 13
comments: 2
posted_at: "2026-07-16T16:41:37Z"
tags:
  - hacker-news
  - translated
---

# Detecting LLM-Generated Texts with "Classical" Machine Learning

- HN: [48936880](https://news.ycombinator.com/item?id=48936880)
- Source: [blog.lyc8503.net](https://blog.lyc8503.net/en/post/llm-classifier/)
- Score: 13
- Comments: 2
- Posted: 2026-07-16T16:41:37Z

## Translation

タイトル: 「古典的な」機械学習による LLM 生成テキストの検出
記事のタイトル: 「古典的な」機械学習による LLM 生成の Web フィクションの検出 (AIGC テキスト検出)
説明: .admonition { マージン: .75em 0;パディング: .6rem;オーバーフロー: 非表示;フォントサイズ: 12px; page-break-inside: 回避します。境界左: .3rem ソリッド #42b983;境界半径: .3rem;ボックスシャドウ: 0 0.1re

記事本文:
「古典的な」機械学習を使用した LLM 生成の Web フィクションの検出 (AIGC テキスト検出) 简体中文 / [英語] ホーム
2. 背景 (別名役に立たないとりとめのない話)
4. (ある程度) 成功した試み – scikit-learn SVM 4.1。データ生成
4.3. Web デモ用の JS 実装
4.5.攻撃と防御 – 検出の回避(?) 4.5.1.古典的な翻訳方法
「古典的な」機械学習による LLM 生成の Web フィクションの検出 (AIGC テキスト検出)
この記事は現在試験的に機械翻訳されているため、誤りが含まれている可能性があります。不明な点がある場合は、オリジナルの中国語版を参照してください。私は翻訳の改善に継続的に取り組んでいます。
2026 年初頭の時点で、主流の LLM で生成されたテキストは、従来の機械学習モデルを使用して人間が書いたコンテンツと効果的に区別できる強力な統計パターンを示しています。これが、いわゆる「AI 盗作チェッカー」の数が実際に内部で働いているのではないかと思います。
オンライン デモ: https://lyc8503.github.io/AITextDetector/
このデモで使用されるモデルは、汎用データでトレーニングされておらず、厳密な最適化や反復も受けていません。単一文の検出精度は、テスト セットで約 85% です。潜在的な制限を理解するために、使用する前にこの記事を読んでください。
コア コード (ドラフト) とトレーニング済みモデル ファイルは、GitHub: lyc8503/AITextDetector で入手できます。
背景 (別名役に立たないランブリング)
半年前、私がまだ学校で論文を書いていた頃、すでに AIGC (AI 生成コンテンツ) の論文チェックについての噂が広まっていました。 CNKI、Wanfang、および一部のサードパーティ AIGC 検出サービスなど、いくつかのプラットフォームをテストしたところ、手書きのテキストと LLM で生成されたテキストをかなりの精度で区別できることがわかりました。
それが火花を散らした

AIGC 検出が実際にどのように機能するか (そしてそれを回避する方法) についての私の好奇心。
しかし、当時私はラジオ、マインクラフト、東方に夢中で、あまりにも多くのことをやりくりしていたので、何度か失敗した後、そのアイデアは棚上げになりました。
結局、私は論文を偽造して読み進め、人生は前に進みました。しかし最近、Lofter を閲覧しているときに、AI が生成した低品質で性格から大きく外れているファンフィクションが大量にタグ付けされているのに遭遇しました。
どうすれば一目で AI だとわかるでしょうか?まあ、一部の人々 (またはギャル) は、投稿する前にマークダウン形式や AI によって生成されたセクション ヘッダーをクリーンアップすることを気にせず、記事の半分をペイウォールの内側に叩き込みます 😓
ただし、AI によって生成されたテキストのほとんどは、さまざまな文体やさまざまなプロンプトの中に埋もれており、すぐにはわかりません。何か違和感を感じたときには、もう手遅れです。一部のテキストは AI であることを証明するのがほぼ不可能で、私を不安にさせます。 AI が生成したうんこを数杯飲み込みすぎた後、ついに十分になりました。 Lofter の閲覧はここで終了します。VS Code を開いてください。
はい、こうして私は週末のプロジェクトのアイデアを復活させることになりました。それは、AI 生成のテキスト検出器を構築するというものでした…
AIGC 検出を検索すると、インターネットは現在、ほぼ完全に広告で汚染されています。すべての結果は、単なるエッセイ AI 書き換えサービスにすぎません。当時、私はノイズをかき分けて、 text perplexity と呼ばれるものを見つけました。
アイデアは単純です。既存の LLM を使用して、特定の文に出現する各単語の確率を推定します。ほぼすべての単語が LLM の予測で上位 (上位 N) にランク付けされている場合、その文は AI によって生成された可能性があります。逆に、予想外の単語が多ければ、それは人間が書いたものである可能性が高くなります。
期待できそうですね?時間をかけてこの方法を試してみましたが、結果は期待外れで、間違いが多かったです。

陽性と偽陰性があり、適切なしきい値を設定できませんでした。さらに、推論コストが高い、モデル間の一般化が不十分、大規模なモデルをローカルにデプロイするのが難しい、クローズドウェイトモデルの統合が難しいなど、実際的な問題もあります。全体として、このアプローチは洗練されておらず、信頼できるものでもありません。
失敗した試行 - 大量のソフト広告「チュートリアル」に騙されました
(ある程度) 成功した試み – scikit-learn SVM
オンライン リソースは役に立たなかったので、昔ながらの錬金術に戻りました。
Scikit-Learn、アクティベートしてください!その ロードマップ に従って、分類タスクの適切な開始点として線形 SVC と単純ベイズを直接選択できます。
(ささやき: これは私の直感とも一致しました。LLM には検出可能な単語選択パターンがあり、単純ベイズ分類器でもそれらを検出するはずです。信号がこれほど強いとは予想していませんでした。)
昔ながらの錬金術の伝統的な分類器にはラベル付きデータが必要です。そのため、トレーニングには人間が書いたテキストと確認済みの LLM 生成テキストが必要です。
私のアプローチ: 2010 年から 2022 年の間に発行された記事 (ChatGPT 以前) をフィルタリングして、特定の Ford 系および River 系プラットフォームから 2023 年に収集したデータを抽出しました。私は、エンゲージメントが非常に低い作品や非常に短い作品だけを除外し、人間が書いたサンプルとして 10,000 近くの数千文字のテキストをランダムに抽出しました。
次に、LLM を使用してこれらのテキストの章の要約を生成し、その要約を LLM にフィードバックして、記事全体を再生成させました。これにより、ジャンルが多様で、元の人間のコンテンツとほぼ一致する、LLM によって生成されたサンプルがほぼ同数得られました。
少なくとも理論的には。しかし、LLM API は高価なので、週末のプロジェクトに何千ドルも費やすつもりはありませんでした。そこで私は創造力を発揮し、複数の低コストまたは無料の API チャネルを利用してルールを回避しました。
Gemini : CLIProxyAPI を使用して連携

Antigravity/Gemini CLI クォータを API アクセスに変換します。AI Pro アカウントに最大 20 ドル支払うだけです
Qwen : qwen-code を使用すると、Qwen Plus API を無料でリバース エンジニアリングできます。
GLM-5 : 幸運なタイミング — OpenRouter が無料の GLM-5 パブリック ベータ版を提供していました (Pony Alpha)
Kimi、Deepseek、Doubao、GLM-4.7 : プロモーション コーディング プラン中にサインアップ - 初月 8.9 ドル、API アクセスのロックが解除
免責事項: これは推奨するものではありません。これらの行為はプラットフォームの利用規約に違反しており、禁止される可能性があります。しかし、プラットフォームはマーケティングの誇大広告で忙しすぎて気にすることができず、全額を支払うつもりはありませんでした。
プログラミングに重点を置いた LLM API の多くは、奇妙なことに呼び出しごとに料金を請求しますが、タスクを大量の入力にバッチ処理して、LLM に呼び出しごとにより多くのコンテンツを生成させることで、これを最適化することを悪用できます。それで…
私が定価で 2000 ドル相当の 3 億以上の Gemini トークンを使用したということはどういう意味ですか?!
最終的に、gemini-3-flash を使用してサマリーを生成し、7 つの異なるモデル ( gemini-3-pro 、 qwen-coder-plus 、 glm-5 、 glm-4.7 、 kimi-k2.5 、 doubao-seed-code 、 deepseek-v3.2 ) を使用して、LLM で生成された 7 セットのサンプルを生成しました。
データ生成の途中でしたが、待ちきれずにトレーニングを開始しました。
クロードに分類子コードを書くように依頼したところ、生のテキスト全体が単純にモデルにダンプされ、99.45% という疑わしい精度が達成されました… 待って、本当に?
クロードは駄目だよ。自分でやります。トレーニングのために、すべてのテキストを中国語の句読点を使用して文に分割し、中国語/英語以外の文字を削除してから、 scikit-learn の TF-IDF → LinearSVC を適用しました。いくつかのノイズを除去した後でも、文レベルの分類は依然として最大 85% の精度を達成しています。
このバグの多い最初のバージョンでも 88% の精度に達しました (後に 85% に最適化されました)
個々の文に含まれる情報は限られていますが、文ごとの精度が 85% であるということは、長い記事の場合は非常に信頼できることを意味します。

AIが生成したかどうかを判断する際に。このパフォーマンスは私の期待をはるかに上回りました。昔ながらの ML は今でも平手打ちを行っていますが、LLM に「このテキストは AI が生成したものですか?」と尋ねるだけの愚かなオンライン ツールよりもはるかに優れています。
すべてのデータを作成し終えた後、8 クラス モデル (人間 + 7 つの AI) をトレーニングしようとしましたが、LLM が類似しすぎているように見えます (おそらく相互に抽出されたものと思われます)。そのため、分類は面倒で、精度はわずか 50% 程度でした。
複数クラスの結果 - 明らかに分離できません。私のモデルは最悪かもしれないが、それは重要ではない
最終的に、7 つの個別のバイナリ分類器をトレーニングし、多数決を使用しました。つまり、2 つ以上のモデルが検出した文には AI としてフラグが付けられます。
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25
26
27
28
29
30
31
32
33
34
35
8536 個のサンプルをロードしました
トレイン チャプター サイズ: 6820、テスト チャプター サイズ: 1716
[ジェミニ] トレイン: 917,374 テスト: 228,051
[ジェミニ] フル TF-IDF + SVC ...
3,336,446 特徴量 -> acc=0.8809 f1=0.8082 [tn=143688 fp=10650 fn=16503 tp=57210]
[qwen] トレーニング: 1,315,338 テスト: 328,636
[qwen] フル TF-IDF + SVC ...
3,989,603 特徴 -> acc=0.8911 f1=0.8974 [tn=136293 fp=18045 fn=17739 tp=156559]
[ポニー] トレイン: 1,128,044 テスト: 278,663
[ポニー] フル TF-IDF + SVC ...
3,688,143 特徴 -> acc=0.8493 f1=0.8286 [tn=135085 fp=19253 fn=22755 tp=101570]
[kimi25] トレイン: 1,088,007 テスト: 269,567
[kimi25] フル TF-IDF + SVC ...
3,976,027 特徴 -> acc=0.8721 f1=0.8473 [tn=139390 fp=14948 fn=19534 tp=95695]
[glm47] トレーニング: 1,124,430 テスト: 279,109
[glm47] フル TF-IDF + SVC ...
3,980,772 特徴 -> acc=0.8436 f1=0.8222 [tn=134461 fp=19877 fn=23786 tp=100985]
[豆宝] 列車: 1,063,395 テスト: 264,121
[doubao] フル TF-IDF + SVC ...
4,243,728 フィーチャ -> acc=0.8940 f1=0.8700 [tn=142420 fp=11918 fn=16089 tp=93694]
[deepseekv32] トレイン: 1,176,294 テスト: 289,042
[deepseekv32] 完全な TF-IDF + SVC

...
4,361,691 フィーチャ -> acc=0.8529 f1=0.8403 [tn=134625 fp=19713 fn=22819 tp=111885]
====================================
概要
====================================
モデル s1 acc s1 f1
ジェミニ 0.8809 0.8082
クウェン 0.8911 0.8974
ポニー 0.8493 0.8286
きみ25 0.8721 0.8473
glm47 0.8436 0.8222
豆宝 0.8940 0.8700
ディープシークv32 0.8529 0.8403
すべてのモデルは 85% 以上の精度と 80% 以上の F1 を達成しており、かなり堅実です。また、AI が生成したテキストには複数のモデルによってフラグが付けられることが多いため、投票するのは完全に理にかなっていることにも気づきました。
MultinomialNB と SGDClassifier を試しましたが、精度がわずかに低下しました。 BERT はわずかな向上をもたらしましたが、必要な GPU 時間が多すぎるため、破棄されました。 AutoGluon をテストしたところ、どういうわけかバイナリ精度が 53% しか管理できませんでした。それらについては詳しく説明しません。
Web デモ用の JS 実装
この時点で、リポジトリを公開して終わりにすることもできました。しかし、毎回 Python を起動するのはあまりにも不便です。 Python API をホストすることもできましたが、それはサーバーのメンテナンスを意味し、私の厳格なサーバーレス哲学に違反します。
私の当初の計画: モデルを ONNX にエクスポートし、Wasm の ONNX Web ランタイム経由で推論を実行します。しかし、シリコン サーバントのクロードに協力を依頼したとき、明確に指定しなかったので、スクリプトから外れ、モデルをトリミングして JSON としてエクスポートしました。その後、ブラウザ推論用に TF-IDF + SVM を完全に JavaScript で実装しました。
うーん…実際には悪いアイデアではありません。 100 万文字のテキストでテストしました。私のマシンでは約 10 秒かかりましたが、許容範囲内です。一般的な数千文字の入力の場合、それは瞬時に行われます。
わかりました。これは単なるデモであり、JS アプローチの方が透明性が高いため、この少し愚かな実装はそのままにしておきます。 （私ではなくクロードを責めてください。）
精度に関しては、さまざまな機能制限をテストしました。最終的にパフォーマンスを優先し、500k の機能を維持しました。 JSON として保存されます。

は 107MB と膨らみます (ただし、サーバー側で gzip 圧縮すると、最大 38MB になります)。より小さいバージョン (50k ～ 80k) では精度が 3 ～ 4% 失われるだけでしたが、最終的な AI 検出率は大きく異なり、特に人間のテキストでは ±50% の相対差があり、誤検知につながりました。だから私は500kにこだわりました。
最終的な精度の低下: 以下に示すように、~1%:
1
2
3
4
5
6
7
8
9
10
11
12
13
14
======================================================
概要上位 500,000 C=1.0
======================================================
モデル s1 acc s2 acc Δacc s1 f1 s2 f1 Δf1
ジェミニ 0.8809 0.8721 -0.0088 0.8082 0.7986 -0.0096
クウェン 0.8911 0.8819 -0.0092 0.8974 0.8886 -0.0088
ポニー 0.8493 0.8383 -0.0109 0.8286 0.8173 -0.0113
kimi25 0.8721 0.8623 -0.0098 0.8473 0.8376 -0.0097
glm47 0.8436 0.8311 -0.0125 0.8222 0.8097 -0.0125
豆宝 0.8940 0.8869 -0.0071 0.8700 0.8624 -0.0076
ディープシークv32 0.8529 0.8419 -0.0110 0.8403 0.8291 -0.0112
平均 0.8592 0.8348
最小値: 0.8311
======================================================
テストパフォーマンス
以下のすべてのテストは、プルーニングされた Web バージョンを使用しており、完全な joblib モデルと同様に実行されるはずです。
現在のロジック: 7 つのバイナリ モデルすべてを使用して、入力テキストを文に分割し、クリーンアップして分類します。 2 つ以上のモデルが文にフラグを立てた場合、その文は AI の疑いとしてマークされ、強調表示されます。最終的な AI スコアは、フラグが立てられたキャラクターの割合です。分類:
はじめに

[切り捨てられた]

## Original Extract

.admonition { margin: .75em 0; padding: .6rem; overflow: hidden; font-size: 12px; page-break-inside: avoid; border-left: .3rem solid #42b983; border-radius: .3rem; box-shadow: 0 0.1re

Detecting LLM-Generated Web Fiction with "Classical" Machine Learning (AIGC Text Detection) 简体中文 / [English] Home
2. Background (aka Useless Rambling)
4. A (Somewhat) Successful Attempt – scikit-learn SVM 4.1. Data Generation
4.3. JS Implementation for Web Demo
4.5. Attack and Defense – Bypassing Detection(?) 4.5.1. Classic Translation Method
Detecting LLM-Generated Web Fiction with "Classical" Machine Learning (AIGC Text Detection)
This article is currently an experimental machine translation and may contain errors. If anything is unclear, please refer to the original Chinese version. I am continuously working to improve the translation.
As of early 2026, mainstream LLM-generated text exhibits strong statistical patterns that can be effectively distinguished from human-written content using traditional machine learning models. I suspect this is how many so-called “AI plagiarism checkers” actually work under the hood.
Online Demo: https://lyc8503.github.io/AITextDetector/
The model used in this demo is not trained on general-purpose data , nor has it undergone rigorous optimization or iteration. Its single-sentence detection accuracy is approximately 85% on the test set. Please read through this article before use to understand potential limitations.
The core code (drafts) and trained model files are available on GitHub: lyc8503/AITextDetector
Background (aka Useless Rambling)
Back when I was still writing my thesis at school half a year ago, rumors were already spreading about checking papers for AIGC (AI-generated content). I tested several platforms—CNKI, Wanfang, and some third-party AIGC detection services—and found they could indeed distinguish between my hand-written text and LLM-generated text with decent accuracy.
That sparked my curiosity about how AIGC detection actually works (and how to bypass it) .
But I was juggling too many things at the time—obsessed with radio , Minecraft, Touhou—and after a few failed attempts, I shelved the idea.
Eventually, I faked my way through the thesis, and life moved on. But recently, while browsing Lofter, I stumbled upon entire tags flooded with low-quality, wildly out-of-character AI-generated fanfics.
How can I tell they’re AI at a glance? Well, some folks (or gals) don’t even bother cleaning up Markdown formatting or AI-generated section headers before posting—and then they slap half the article behind a paywall 😓
Most AI-generated texts, however, are harder to spot—they’re buried among diverse writing styles, varied prompts, and not immediately obvious. By the time you realize something feels off, it’s too late. Some texts are borderline impossible to prove as AI, leaving me paranoid. After swallowing a few too many AI-generated turds, I’d finally had enough. Lofter browsing stops here—time to open VS Code!
Yes, that’s how I ended up reviving my weekend project idea: building an AI-generated text detector…
The internet is now almost entirely polluted with ads when searching for AIGC detection. Every result is just another essay-AI-rewriting service. Back then, I dug through the noise and found something called text perplexity .
The idea is simple: use an existing LLM to estimate the probability of each word appearing in a given sentence. If nearly every word ranks high in the LLM’s predictions (Top-N), the sentence is likely AI-generated. Conversely, if many words are unexpected, it’s more likely human-written.
Sounds promising, right? I spent some time trying this method, but results were disappointing—plenty of false positives and false negatives, and no reasonable threshold could be set. Plus, there are practical issues: high inference cost, poor cross-model generalization, difficulty deploying large models locally, and closed-weight models being hard to integrate. Overall, this approach isn’t elegant or reliable.
Failed attempt—got tricked by a bunch of soft-ad "tutorials"
A (Somewhat) Successful Attempt – scikit-learn SVM
Since online resources were useless, back to old-school alchemy.
Scikit-learn, activate! Following its Roadmap , we can directly pick Linear SVC and Naive Bayes as good starting points for our classification task.
(Whisper: this also matched my gut feeling—LLMs have detectable word-choice patterns; even a Naive Bayes classifier should pick them up. I just didn’t expect the signal to be this strong.)
Old-school alchemy traditional classifiers need labeled data—so we need human-written texts and confirmed LLM-generated texts for training.
My approach: I pulled data I’d scraped in 2023 from a certain Ford-like and River-like platform, filtering for articles published between 2010–2022 (pre-ChatGPT). I only filtered out extremely low-engagement or very short pieces, then randomly sampled nearly 10,000 multi-thousand-character texts as human-written samples.
Then, I used an LLM to generate chapter summaries of these texts, fed the summaries back into the LLM, and had it regenerate full articles. This gave me a roughly equal number of LLM-generated samples, diverse in genre and closely matching the original human content.
In theory, at least. But LLM APIs are expensive, and I wasn’t about to spend thousands on a weekend project. So I got creative—and skirted the rules leveraged multiple low-cost or free API channels:
Gemini : Used CLIProxyAPI to convert Antigravity/Gemini CLI quota into API access—just pay ~$20 for an AI Pro account
Qwen : qwen-code lets you reverse-engineer the Qwen Plus API—free
GLM-5 : Lucky timing—OpenRouter was offering free GLM-5 public beta (Pony Alpha)
Kimi, Deepseek, Doubao, GLM-4.7 : Signed up during a promotional coding plan—$8.9 first month, API access unlocked
Disclaimer: This is not a recommendation. These actions violate platform ToS and may get you banned. But the platforms are too busy with marketing hype to care, and I wasn’t about to pay full price.
Many programming-focused LLM APIs strangely charge per call, but we can abuse optimize this by batching tasks into massive inputs, forcing the LLM to generate more content per call. And so…
What do you mean I used over 300M Gemini tokens worth $2000 at full price?!
Ultimately, I used gemini-3-flash to generate summaries, and seven different models ( gemini-3-pro , qwen-coder-plus , glm-5 , glm-4.7 , kimi-k2.5 , doubao-seed-code , deepseek-v3.2 ) to generate seven sets of LLM-generated samples.
While I was halfway through data generation, I couldn’t wait and started training.
I asked Claude to write the classifier code, and it naively dumped the entire raw text into the model—achieving a suspicious 99.45% accuracy… Wait, really?
Claude’s useless. I’ll do it myself. For training, I split all texts into sentences using Chinese punctuation, cleaned non-Chinese/English characters, then applied scikit-learn’s TF-IDF → LinearSVC . After cleaning up some noise, sentence-level classification still hit ~85% accuracy!
Even this buggy first version hit 88% accuracy (later optimized to 85%)
Individual sentences carry limited info, but 85% accuracy per sentence means that for a longer article, we can be highly confident in judging whether it’s AI-generated. This performance far exceeded my expectations. Old-school ML still slaps—way better than those dumb online tools that just ask an LLM, “Hey, is this text AI-generated?”
After finishing all data, I tried training an 8-class model (human + 7 AIs), but the LLMs seem too similar—probably distilled from each other—so classification was messy, with only ~50% accuracy.
Multi-class results—apparently not separable. Maybe my model sucks, but whatever, not important
Eventually, I trained seven separate binary classifiers and used majority voting: a sentence is flagged as AI if ≥2 models detect it.
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25
26
27
28
29
30
31
32
33
34
35
loaded 8536 samples
train chapter size: 6820, test chapter size: 1716
[gemini] Train: 917,374 Test: 228,051
[gemini] full TF-IDF + SVC ...
3,336,446 features -> acc=0.8809 f1=0.8082 [tn=143688 fp=10650 fn=16503 tp=57210]
[qwen] Train: 1,315,338 Test: 328,636
[qwen] full TF-IDF + SVC ...
3,989,603 features -> acc=0.8911 f1=0.8974 [tn=136293 fp=18045 fn=17739 tp=156559]
[pony] Train: 1,128,044 Test: 278,663
[pony] full TF-IDF + SVC ...
3,688,143 features -> acc=0.8493 f1=0.8286 [tn=135085 fp=19253 fn=22755 tp=101570]
[kimi25] Train: 1,088,007 Test: 269,567
[kimi25] full TF-IDF + SVC ...
3,976,027 features -> acc=0.8721 f1=0.8473 [tn=139390 fp=14948 fn=19534 tp=95695]
[glm47] Train: 1,124,430 Test: 279,109
[glm47] full TF-IDF + SVC ...
3,980,772 features -> acc=0.8436 f1=0.8222 [tn=134461 fp=19877 fn=23786 tp=100985]
[doubao] Train: 1,063,395 Test: 264,121
[doubao] full TF-IDF + SVC ...
4,243,728 features -> acc=0.8940 f1=0.8700 [tn=142420 fp=11918 fn=16089 tp=93694]
[deepseekv32] Train: 1,176,294 Test: 289,042
[deepseekv32] full TF-IDF + SVC ...
4,361,691 features -> acc=0.8529 f1=0.8403 [tn=134625 fp=19713 fn=22819 tp=111885]
=====================================
SUMMARY
=====================================
model s1 acc s1 f1
gemini 0.8809 0.8082
qwen 0.8911 0.8974
pony 0.8493 0.8286
kimi25 0.8721 0.8473
glm47 0.8436 0.8222
doubao 0.8940 0.8700
deepseekv32 0.8529 0.8403
All models achieved over 85% accuracy and over 80% F1—pretty solid! I also noticed that AI-generated texts were often flagged by multiple models, so voting made perfect sense.
I tried MultinomialNB and SGDClassifier , but accuracy dropped slightly. BERT gave a minor boost but required too much GPU time—discarded. Even tested AutoGluon , which somehow managed only 53% binary accuracy. Won’t dive into those.
JS Implementation for Web Demo
At this point, I could’ve just published the repo and called it a day. But launching Python every time is way too inconvenient. I could’ve hosted a Python API, but that means server maintenance—violates my strict Serverless philosophy.
My original plan: export model to ONNX, run inference via ONNX Web Runtime in Wasm. But when I asked my silicon servant Claude to help, I didn’t specify clearly—and it went off-script, trimming and exporting the model as a JSON… then implemented TF-IDF + SVM entirely in JavaScript for browser inference.
Hmm… actually not a bad idea. I tested it on a 1-million-character text—it took about 10 seconds on my machine, acceptable. For typical few-thousand-character inputs, it’s instant.
Fine, since this is just a demo, and the JS approach is more transparent, I’ll keep this slightly silly implementation. (Blame Claude, not me.)
As for accuracy: I tested different feature limits. Ultimately prioritized performance and kept 500k features. Stored as JSON, it’s a bloated 107MB (though gzipped server-side, it’s ~38MB). Smaller versions (50k–80k) only lost 3–4% accuracy, but final AI detection rates varied significantly—especially on human texts, with ±50% relative differences, leading to false positives. So I stuck with 500k.
Final accuracy drop: ~1%, as shown below:
1
2
3
4
5
6
7
8
9
10
11
12
13
14
============================================================
SUMMARY top-500,000 C=1.0
============================================================
model s1 acc s2 acc Δacc s1 f1 s2 f1 Δf1
gemini 0.8809 0.8721 -0.0088 0.8082 0.7986 -0.0096
qwen 0.8911 0.8819 -0.0092 0.8974 0.8886 -0.0088
pony 0.8493 0.8383 -0.0109 0.8286 0.8173 -0.0113
kimi25 0.8721 0.8623 -0.0098 0.8473 0.8376 -0.0097
glm47 0.8436 0.8311 -0.0125 0.8222 0.8097 -0.0125
doubao 0.8940 0.8869 -0.0071 0.8700 0.8624 -0.0076
deepseekv32 0.8529 0.8419 -0.0110 0.8403 0.8291 -0.0112
AVG 0.8592 0.8348
MIN acc: 0.8311
============================================================
Testing Performance
All tests below use the pruned web version, which should perform similarly to the full joblib models.
Current logic: split input text into sentences, clean and classify using all 7 binary models. If ≥2 models flag a sentence, it’s marked as suspected AI and highlighted. Final AI score is the proportion of flagged characters. Classification:
Firs

[truncated]
