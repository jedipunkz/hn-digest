---
source: "https://www.csselectronics.com/pages/can-bus-reverse-engineering-ai-llm-claude"
hn_url: "https://news.ycombinator.com/item?id=48642917"
title: "Car Hacking with Claude Code"
article_title: "CAN Bus Reverse Engineering with AI [Claude Code]\n– CSS Electronics"
author: "martinfalch"
captured_at: "2026-06-23T11:02:05Z"
capture_tool: "hn-digest"
hn_id: 48642917
score: 1
comments: 0
posted_at: "2026-06-23T10:38:05Z"
tags:
  - hacker-news
  - translated
---

# Car Hacking with Claude Code

- HN: [48642917](https://news.ycombinator.com/item?id=48642917)
- Source: [www.csselectronics.com](https://www.csselectronics.com/pages/can-bus-reverse-engineering-ai-llm-claude)
- Score: 1
- Comments: 0
- Posted: 2026-06-23T10:38:05Z

## Translation

タイトル: クロードコードによるカーハッキング
記事のタイトル: AI を使用した CAN バス リバース エンジニアリング [クロード コード]
– CSSエレクトロニクス
説明: AI を使用して独自の CAN バス信号をリバース エンジニアリングします。 CANsub、python-can、および Claude Code スキルのデコード速度、RPM などを数分で確認できます。

記事本文:
AI を使用した CAN バス リバース エンジニアリング [Claude Code]
– CSSエレクトロニクス
コンテンツにスキップ
カートに追加したばかりです
AI を使用した CAN バス リバース エンジニアリング [Claude Code]
AI エージェントは CAN バス データをリバース エンジニアリングできますか?
この記事では、Martin Falch が当社の CANsub CAN バスを使用しています
Python-can を使用したインターフェース
リバースエンジニアリングするためのカスタムクロードコード「スキル」
複数の CAN バス信号。
CAN バス リバース エンジニアリングの基本については、CAN スニファーの記事を参照してください。
古典的な人体分析アプローチをカバーしています。この記事では、LLM で何ができるかを示します。
ネタバレ注意: 実際の実験から明らかなように、非常にうまく機能します。
ショーケース - このコンセプトは CAN スニッフィングの未来です。
しかし、私たちの言葉を鵜呑みにしないでください。このガイドは、今すぐ自分で試してみるのに役立ちます。
CANベースと外部リファレンスの比較
ショーケース×3
速度と回転数 (OBD2)
(2026 年 6 月更新)
CAN バス リバース エンジニアリングの実践的なショーケースを複数提供します。
あい！
課題: CAN バスのリバース エンジニアリングは難しい
独自の CAN 信号を手動でリバース エンジニアリングするのは、非常に難しいことで知られています。
一般的なバスでは数千台のバスが生産されます。
数十の CAN ID にわたって 1 秒あたりのフレーム数が増加し、関心のあるパラメータは生データのどこかに隠されています。
バイト。
それを抽出するのは、干し草の山から針を見つけるようなものです。
大量の CAN フレームをトロールして、どの CAN ID に信号が含まれている可能性があるかを特定します
ペイロード内のフィールドの正確な開始ビットとビット長を推測します。
バイト順序 (リトル エンディアンとビッグ エンディアン) および値が符号付きかどうかを判断します。
生の値を物理単位にマッピングするスケールとオフセットを解決します。
デコードされた結果が最終的に正しくなるまで、この試行錯誤のループを繰り返します。
これには時間がかかり、実際の専門知識が必要であり、信号ごとに数時間かかることもよくあります。
たとえば、分析のスクリプトを作成します。 Python は役立ちますが、需要もあります

DSの特化スキル。さらに、単一ではありません
あらゆる信号をデコードするスクリプトには、それぞれ独自の癖があります。
解決策: CANsub + クロード コード スキル
これらの課題に対処するために、AI/LLM ツールが非技術者ユーザーのリバース エンジニアリングに役立つかどうかを確認したいと思いました。
CAN信号。それをテストするために、クロード コードを構築しました。
スキルを習得し、それを使用して、 CANsub.2 と python-can を使用してさまざまな信号セットをリバース エンジニアリングしました。
注: ここでの目標は、この特定のスキルをプッシュすることではありません (ただし、使用することは歓迎です)。それは
概念的なものを示すために
ワークフロー。
次のセクションでは、その仕組み、開始方法、およびワークフローの実践的なショーケースについて説明します。
活動中。
以下に、最初のプロンプトからすぐに使用できる DBC ファイルまで、スキルが内部で何を行うのかを簡単に説明します。
Claude Code スキルは、LLM の命令、いくつかのコンテキスト資料、および
LLM が使用できる小さな Python スクリプトのコレクション。プロンプトに従って、AI エージェントは指示に従います
コンテキストとデータに基づいて、どのツールを使用するかを動的に決定します。
このスキルには、LLM の知識資料が含まれています: CANsub と python-can の仕組み、標準 OBD2
OBD 基準信号をデコードするための DBC - および CAN スニファーの記事
基本。
これにより、AI エージェントが賢明な決定を下し、幻覚を減らすことができます。
重要なのは、すべての CAN スニッフィング分析には、独自の生の CAN データを比較するための「参照」信号が必要であるということです。
反対。の
スキルは 2 つのアプローチをサポートしています。
独自の CAN データと CAN ベースの基準信号 (OBD2 など) を含む生の CSV ログ ファイルを提供します。
車両
速度)
上記が不可能な場合は、以下のように外部ソースを介してライブの基準信号を作成します。
信号が目に見える物理値（車のダッシュボード上の「X km/h」など）を持つ場合、

不可知的なツール、アプリなど -
を使用できます
基準信号としてビデオ録画を行います。あなたは、
CANsub を使用した独自の CAN トレース CSV。 iPhone で車のダッシュボードを撮影します。スキル
光学式文字認識 (OCR) を使用して、後続の分析のためにデジタル化された参照信号を抽出します。
準備ができた CAN リファレンスがない (および表示値がない) 信号の場合、スキルは小さなブラウザ アプリを起動できます。
「人間による入力」を可能にするため、
を反映します
現実世界の値、例:ドアの開閉やハンドルの位置など。
アプリは、分析の参照として、入力をタイムスタンプとともに記録し、CAN トレースに同期します。
トレースとリファレンスを使用して、LLM はさまざまな Python スクリプトを使用して段階的な分析を実行します。これ
ごとに並ぶことも含まれます
データ内の「候補フィールド」
それらを参照し、それらがどの程度緊密に連携するかによってランク付けします。次に、ズームインして正確な開始ビットを特定します。
長さ、バイトオーダー、符号。次に、スキルによってスケールとオフセットが決定されます。最後に、さまざまな健全性チェックと
修正が行われます。
分析が完了すると、AI エージェントはすぐに使用できる DBC ファイルを「アプリケーション/シグナル」に書き込みます。
サブフォルダー。それはまた「結果」を生み出します
プロット」は、レビュー用にリファレンス信号とデコード信号を示しています。さらにデコードすると、
信号を結合したアプリケーション DBC にマージするよう Claude に依頼できます。これを webCAN に直接ロードしてプロットできます。
デコードをライブで検証します。
私は数年前にオリジナルの CAN スニファーに関する記事を書き、最近 CANsub の発表に合わせて更新しました。
最初のドラフト以来、私は自然にさまざまな AI ツールを使い始めました。 ChatGPT + CAN に関する私の記事、
グラファナ
アシスタントなど。私は今、Claude Code + スキルを毎日使用していますが、これは素晴らしい「インフラストラクチャ」になると思いました。

'のための
スニッフィングができます。
私がスキル開発にどのように取り組んだかは次のとおりです。
私は意図的に Python スクリプト/ツールまたは SKILL.md を自分でレビュー/作成しませんでした
まず、関連するコンテキスト (CANsub の「Docs for AI」、python-can-cansub README など) を共有することから始めました。
次に、特定の CAN 信号のリバース エンジニアリングを手伝ってくれるようにクロードに依頼しました (データ/リファレンス入力を提供し、
答え）
私は結果を確認し、問題点を強調し、LLM に問題の解決を依頼しました。
このプロセスは、さまざまな課題を抱えて数回繰り返されました。
全体を通して、スキルは、私が使用した例への過度の適合を避けるために「サンプル外」でも評価されました。
スキル開発
このプロセスにはおそらく 5 ～ 10 時間かかりました。意図的にこれを磨きませんでした。
このスキルはさらに洗練される可能性がありますが (PR は大歓迎です)、この基本バージョンでも非常に優れたパフォーマンスを発揮します。
実際には、これには複数の課題が伴います。
コンテキスト: まず第一に、Claude や ChatGPT などの AI ツールは、「コンテキスト ウィンドウ」が最適に機能する場合に最適です。
無関係な情報で埋め尽くされていない。直接提供するテキストやデータが多ければ多いほど、彼らは「愚か」になります
が現れます。代わりに、関連するコンテキストを提供する必要があります。概要統計、具体的
これにより、主要な洞察に直接集中できるようになります。
ツール: ブラウザで Claude/ChatGPT を使用する場合、次のようなツールが自動的に使用される場合があります。
賢明な場合は Python - 例:非常に大きな CSV ファイルをアップロードする場合。そして、私の別のChatGPT + CANによると
記事に記載されているように、結果は非常に印象的なものになる可能性があります。ただし、新しい会話を開始すると、チャットボットは
基本的に、使用する各スクリプトを最初から再生成する必要があります。その場合、かなり一貫性のない、または遅い結果が得られます。
このためです。 Claude Code のスキルコンセプトは、ツール/ワークフロー/続きを「記憶」する方法です

今後はこうなります
繰り返し役立つため、簡単な反復的な改善も可能になります
ローカル アクセス: ブラウザベースの ChatGPT が提供されている分析ツールやスキルを使用した場合でも、
データを削除しても、たとえば CANsub を直接制御することはできません。さらに時間もかかるだろうし
ファイルをアップロード/ダウンロードする必要があるため、
繰り返し。 LLM を活用するローカル ワークスペースを持つことは大幅に効率的です
CANsub がこれに最適な理由
このワークフローはあらゆる Python-CAN インターフェイスに適応できますが、CANsub は特に適しています
AI を活用したリバース エンジニアリングに適しています。
LLM ワークフローでは、何が起こっているのかが不明確になることがあります。 LCD に即座に表示されます: 録音中か、録音内容は何ですか
バスの積載量は、
サイレント モードですが、CAN エラーなどでスタックしていますか?
webCAN により、
生/デコードされた CAN データをストリーミングして視覚化し、LLM で使用できるように旅行データを CSV にエクスポートします。それは
LLM ワークフローに最適なコンパニオン ツール
CANsub のサイレント (リッスンオンリー) モードにより、ライブバスを邪魔することなく安全に観察でき、エラーフレームも発生しません。
ロギングは重要です
分析中にバスの問題を発見するため
多くの車の OBD2 コネクタは独自の CAN を公開していません。 2 x CAN を使用すると、CAN1 で OBD2 をログに記録でき、CAN1 で独自の CAN をログに記録できます。
CAN2
非接触アダプター経由
1 つの CSV にエクスポートする場合
このワークフローをそのまま実行したい場合は、以下の項目を取得することをお勧めします。
1×CANsub.2 + 1
OBD2-DB9アダプター×
車が OBD2 コネクタ経由で独自の CAN を公開していない場合は、以下を追加します。
1×非接触
アダプター + 1 x レセプタクル DC
アダプター
次に、can-bus-reverse-engineering-skill に関する README に従います。
GitHub (10 分) 。
ヒント: サンプル CAN+OBD2 データを使用してスキルを試すことができます (README を参照)。
おそらく、ブラウザで ChatGPT を使用することに慣れているかもしれませんが、Claude のような AI コーディング ツールを試したことはありません。
コード。

もしそうなら、これは最初は気が遠くなるように思えるかもしれません。
ただし、README の手順に従えば、非常に簡単に始めることができることがわかります。それは本質的には
いくつかのものをインストールするだけです - そして
その後、Visual Studio Code を介して PC 上で直接、通常のプロンプト プロセスに戻ります。そして今
チャットボットはフォルダー、ファイル、Python、CANsub にアクセスできます。
そして、
この記事があなたの新人研修である場合は
クロード コード + スキルを日常業務で使い始めるためのポイントは何ですか?まあ、後でお礼を言ってもいいよ。
「AI CAN スニファー データ パック」を入手
この記事のショーケースを自分で再現してみませんか?
以下を含む「データ パック」をダウンロードします。
3 x 生の車用 CAN データと一致する OBD2 基準信号
Opel Astra CAN データ + ビジョン ワークフローのダッシュボード ビデオ
これを使用して 2 つのショーケースを複製し、新しい CAN 信号をリバース エンジニアリングできる可能性があります。
基準信号: CAN ベースと外部
始める前に、基準信号を提供する方法を決定します。
CAN ベースのリファレンス (ログ ファイル)
ターゲット信号に直接 CAN ベースのデコード可能なリファレンス/プロキシ (OBD2 速度/RPM または GPS-to-CAN など) がある場合
速度)、常に使用します
これ。あなたはログを記録します
独自の CAN と並行して CAN 信号を参照し、スキルはログから参照を直接抽出します。
これは
基準は機械的に正確であるため、最も正確なオプションです。 webCAN での CANsub の CSV エクスポートにより、
このログをキャプチャするのは簡単です。
外部参照（ストリーミング）
CAN ベースのリファレンスが存在しない場合は、たとえば次のようにしてリアルタイムで自分で生成します。の「リファレンスジェネレーター」アプリ
の
スキルを向上させるか、車のダッシュボードの速度などの物理的な値を撮影し、
スキル。あなた
ストリーミング中に現実世界の値を反映します。これは CAN ベースのリファレンスよりもノイズが多くなりますが、より不可知論的です
そして

他の方法では参照できない信号に対して機能します。
すでに実行できる独自の CAN 信号のリバース エンジニアリングに時間を費やすのは愚かに思えるかもしれません。
OBD2経由でリクエストとデコードを行います。ただし、多くの場合、独自の CAN 信号と同等の信号の方が正確です。
(例: 解像度
0.1 km/h 対 1 km/h）、より高い周波数（例：10 Hz 対、OBD2 の場合は 0.2 Hz）でも利用可能です。
周波数は
特に多くの OBD2 PID をリクエストする必要がある場合に関連します。たとえば、それらの間隔を空ける必要があるためです。それぞれ0.5秒。
これは、10 x OBD2 PID をリクエストするには、10 x 0.5s = 5 秒の「期間」が必要であることを意味します。このような場合にできることは、
OBD2 速度を 5 秒ごとに観察します。これは多くの使用例で問題になる可能性があります。
最後に、独自の CAN 信号は、たとえば、次の方法でサイレントに記録できます。 CANsub または CAN ロガーのような
CANedge。これ
できるという意味です
車両からの OBD2 データの要求を回避することで、取り付けが簡素化され、噴射のリスクが軽減され、バッテリーの消耗が軽減されます。
などなど。
自然と楽になりますよ
非常に多様でフルレンジの代表的なセットを提供する場合、関連する信号を分離します。
対応する生の独自データと比較するための参照データ。
私の経験では、CAN リファレンス信号がある場合、
いいね

[切り捨てられた]

## Original Extract

Reverse engineer proprietary CAN bus signals with AI. See how the CANsub, python-can and a Claude Code skill decode speed, RPM and more in minutes.

CAN Bus Reverse Engineering with AI [Claude Code]
– CSS Electronics
Skip to content
Just added to your cart
CAN Bus Reverse Engineering with AI [Claude Code]
Can an AI agent reverse engineer your CAN bus data?
In this article, Martin Falch uses our CANsub CAN bus
interface together with python-can
and a custom Claude Code 'skill' to reverse engineer
multiple CAN bus signals.
For basics on CAN bus reverse engineering see our CAN sniffer article ,
which covers the classic human-analysis approach. This article shows what is possible with an LLM.
Spoiler alert: It works extremely well , as evident from our practical
showcases - this concept is the future of CAN sniffing.
But don't take our word for it - this guide will help you try it yourself right now !
CAN-based vs. external reference
3 x showcases
Speed & RPM (OBD2)
(updated June 2026)
We'll provide multiple practical showcases of CAN bus reverse engineering with
AI!
The challenge: CAN bus reverse engineering is hard
Reverse engineering a proprietary CAN signal by hand is notoriously tough.
A typical bus produces thousands of
frames per second across dozens of CAN IDs, and the parameter you care about is hidden somewhere in the raw
bytes.
Extracting it is like finding a needle in a haystack :
Trawl huge volumes of CAN frames to spot which CAN ID might contain your signal
Guess the exact start bit and bit length of the field inside the payload
Figure out the byte order (little vs. big endian) and whether the value is signed
Solve for the scale and offset that map the raw value to a physical unit
Repeat this trial-and-error loop until the decoded result finally looks right
This is slow and demands real expertise - and often takes hours per signal.
Scripting the analysis in e.g. Python helps, but also demands specialized skills. Further, there is no single
script that decodes any signal, each has its own quirks.
The solution: CANsub + a Claude Code skill
To address these challenges, I wanted to see if AI/LLM tools could help non-technical users reverse engineer
CAN signals. To test it, I built a Claude Code
skill and used it to reverse engineer a diverse set of signals with the CANsub.2 and python-can .
Note: The goal here is not to push this specific skill (though you are welcome to use it). It is
to show a conceptual
workflow.
In the following sections I explain how it works, how you get started - and practical showcases of the workflow
in action.
Below I outline briefly what the skill does under the hood - from your first prompt to a ready-to-use DBC file.
The Claude Code skill is just a folder with instructions for the LLM, some context material and a
collection of small Python scripts the LLM can use. Following your prompt, the AI agent follows the instructions
step by step and decides on what tools to use dynamically based on the context and your data.
The skill includes knowledge material for the LLM: How the CANsub and python-can work, our standard OBD2
DBC for decoding OBD reference signals - and our CAN sniffer article for
basics.
This helps the AI agent make sensible decisions and hallucinate less.
Importantly, every CAN sniffing analysis needs a 'reference' signal to compare the raw proprietary CAN data
against. The
skill supports two approaches:
You provide a raw CSV log file with the proprietary CAN data and a CAN-based reference signal (e.g. OBD2
vehicle
speed)
If the above is not possible, you create a reference signal live yourself via an external source as per below
If your signal has a visible physical value - like 'X km/h' on your vehicle dashboard, diagnostic tool, app etc -
you can use a
video-recording as a reference signal. You record a
proprietary CAN trace CSV with the CANsub while e.g. filming the vehicle dashboard with your iPhone. The skill
uses Optical Character Recognition (OCR) to extract a digitized reference signal for the subsequent analysis.
For signals without a ready CAN reference (and with no display value), the skill can launch a small browser app
to enable 'human input' that
mirrors the
real-world value, e.g. opening/closing the door or the position of a steering wheel.
The app records your input with timestamps, synced to the CAN trace, as a reference to analyze against.
With a trace and a reference in hand, the LLM performs a step-by-step analysis using various Python scripts. This
includes lining up every
'candidate field' in the data against your
reference and ranking them by how closely they move together. It then zooms in to pin down the exact start bit,
length, byte order and sign. Next, the skill determines the scale and offset. Finally, various sanity checks and
corrections are performed.
With the analysis complete, the AI agent writes a ready-to-use DBC file in an 'application/signal'
subfolder. It also produces a 'result
plot' showing the reference vs. decoded signal for your review. As you decode more
signals, you can ask Claude to merge them into a combined application DBC, which you can load straight into webCAN to plot
and verify your decoding live.
I wrote our original CAN sniffer article some years ago and recently updated it with our launch of the CANsub.
Since the initial draft, I have naturally started using various AI tools - see e.g. my articles on ChatGPT + CAN ,
Grafana
Assistant and more. I now use Claude Code + skills daily and thought it would be a great 'infrastructure' for
CAN sniffing.
Below is how I approached the skill development:
I deliberately did not review/create the Python scripts/tools or SKILL.md myself
I began by sharing relevant context (our CANsub 'Docs for AI', our python-can-cansub README etc)
I then asked Claude to help reverse engineer specific CAN signals (I provided the data/reference inputs and knew
the answers)
I reviewed the results, highlighted problems and asked the LLM to help resolve them
This process was then repeated several times with a diverse set of challenges
Throughout, the skill was also evaluated 'out of sample' to avoid over-fitting to the examples I used for the
skill development
The process took me probably 5-10 hours - I intentionally did not polish this.
The skill could be refined much more (PRs are welcome), but even this basic version performs really well.
In practice this would involve multiple challenges:
Context: First of all, AI tools like Claude and ChatGPT work best when their 'context windows'
are not filled up with irrelevant information. The more text/data you provide them directly, the 'dumber' they
will appear. Instead, you need to provide them the relevant context - e.g. summary statistics, specific
scripted outputs etc. This lets them focus directly on the key insights
Tools: If you work with Claude/ChatGPT in your browser, it may automatically use tools such as
Python when it is sensible - e.g. if you upload a very large CSV file. And as per my separate ChatGPT + CAN
article, the results can be quite impressive. However, when you start up a new conversation, the chat bot will
essentially need to re-generate each script it uses from scratch - and you'll get rather inconsistent/slow results
because of this. Claude Code's skills concept is a way to 'memorize' tools/workflows/context that will be
repeatedly useful - and thus also enables easy iterative improvements to be made
Local access: Even if your browser-based ChatGPT used tools and skills to analyze provided
data, it would still be unable to directly control your CANsub, for example. It would also be more time-consuming
as you would need to upload/download files
repeatedly. Having a local workspace where you leverage the LLM is drastically more effective
Why the CANsub is ideal for this
This workflow can be adapted to any python-can interface, but the CANsub is especially well
suited for AI-driven reverse engineering:
In LLM workflows it can be unclear what is happening. The LCD shows it instantly: Are we recording, what is the
busload, are
we in silent mode, are we stuck with CAN errors etc
webCAN lets you
stream and visualize raw/decoded CAN data - and export trip data to CSV for use by your LLM. It's
the perfect companion tool for LLM workflows
The CANsub's silent (listen-only) mode lets you observe a live bus safely without disturbing it, and error frame
logging is critical
for spotting bus issues during analysis
Many car OBD2 connectors expose no proprietary CAN. With 2 x CAN you can log OBD2 on CAN1 and proprietary CAN on
CAN2
via a contactless adapter
for export to one CSV
If you wish to follow this workflow as-is I recommend getting below items:
1 x CANsub.2 + 1
x OBD2-DB9 adapter
If your car does not expose proprietary CAN via the OBD2 connector, add below:
1 x contactless
adapter + 1 x receptacle-DC
adapter
Next, follow the README on the can-bus-reverse-engineering-skills
GitHub (10 min) .
Tip: You can try the skill with our sample CAN+OBD2 data (see the README).
Maybe you're familiar with using ChatGPT in your browser - but you've never tried an AI coding tool like Claude
Code. If so, this may seem daunting at first.
However, if you follow the README steps you'll find that it's quite simple to get started. It's essentially
just installing a couple of things - and
then you are back to your normal prompting process, just directly on your PC via Visual Studio Code. And now
your chatbot can access folders, files, Python - and your CANsub.
And,
should it happen that this article is your onboarding
point for starting to use Claude Code + skills in your daily work? Well you can thank me later.
Get your 'AI CAN sniffer data pack'
Want to reproduce the showcases from this article yourself?
Download your 'data pack' including:
3 x raw car CAN data with matching OBD2 reference signals
Opel Astra CAN data + dashboard video for vision workflow
Use this to replicate two of our showcases - and potentially reverse engineer new CAN signals!
Reference signal: CAN-based vs. external
Before you start, determine how you will provide your reference signal:
CAN-based reference (log files)
If the target signal has a direct CAN based and decodable reference/proxy (such as OBD2 speed/RPM or GPS-to-CAN
speed), always use
this. You log the
reference CAN signal alongside the proprietary CAN and the skill extracts the reference straight from the log.
This is
the most precise option, since the reference is machine-accurate. The CANsub's CSV export in webCAN makes
capturing this log trivial.
External reference (streaming)
If no CAN based reference exists, generate one yourself in real time via e.g. the 'reference generator' app in
the
skill or by filming a physical value like your car's dashboard speed and using the vision-recognition tool of the
skill. You
mirror the real-world value as you stream. This is more noisy than a CAN based reference, but it is more agnostic
and works for signals you cannot reference any other way.
It may seem silly to spend time reverse engineering proprietary CAN signals that you are already able to
request and decode via OBD2. However, the proprietary CAN signal equivalent is often more precise
(e.g. a resolution
of 0.1 km/h vs. 1 km/h) and available at higher frequency (e.g. 10 Hz vs. e.g. 0.2 Hz for OBD2).
The frequency is
particular relevant if you need to request many OBD2 PIDs, as you need to space them out by e.g. 0.5 seconds each.
This means that requesting 10 x OBD2 PIDs requires a 'period' of 10 x 0.5s = 5 seconds. In such a case, you can only
observe OBD2 speed once every 5 seconds - which can be a problem in many use cases.
Finally, proprietary CAN signals can be silently recorded by e.g. the CANsub or a CAN logger like
the CANedge. This
means you can
avoid requesting OBD2 data from the vehicle, which simplifies installation, injection risk, battery drain mitigation
and more.
It is naturally easier
to isolate the relevant signal if you provide a highly diverse and full-range representative set of
reference data to compare against the corresponding raw proprietary data.
In my experience, if you have a CAN reference signal
lik

[truncated]
