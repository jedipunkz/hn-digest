---
source: "https://www.theregister.com/ai-ml/2026/05/31/netflix-wiz-creates-app-to-slash-ai-bills-then-open-sources-it/5248702"
hn_url: "https://news.ycombinator.com/item?id=48348679"
title: "Netflix Wiz creates app to slash AI bills, then open sources it"
article_title: "Netflix wiz creates app to slash AI bills, then open sources it"
author: "joebuckwilliams"
captured_at: "2026-06-01T00:25:59Z"
capture_tool: "hn-digest"
hn_id: 48348679
score: 11
comments: 2
posted_at: "2026-05-31T19:09:05Z"
tags:
  - hacker-news
  - translated
---

# Netflix Wiz creates app to slash AI bills, then open sources it

- HN: [48348679](https://news.ycombinator.com/item?id=48348679)
- Source: [www.theregister.com](https://www.theregister.com/ai-ml/2026/05/31/netflix-wiz-creates-app-to-slash-ai-bills-then-open-sources-it/5248702)
- Score: 11
- Comments: 2
- Posted: 2026-05-31T19:09:05Z

## Translation

タイトル: Netflix Wiz が AI 費用を削減するアプリを作成し、オープンソース化
記事のタイトル: Netflix のウィズが AI 費用を削減するアプリを作成し、それをオープンソース化
説明: プロジェクトのヘッドルームにより、多額の費用も節約できる

記事本文:
メインコンテンツへジャンプ
検索
トピックス
特別な機能
すべての特別な機能
Capgemini と AWS でそれを実現する
Nutanix: Kubernetes をスケールします。カオスではありません。
Netflix のウィズが AI 費用を削減するアプリを作成し、それをオープンソース化
Project Headroom は多額の費用も節約できる
Uber と Microsoft の COO が最近学んだように、企業のエンジニアに AI の積極的な使用を奨励すると、高額な使用料金が発生し、おそらく従業員の解雇による利益がすべて相殺される可能性があります。
Netflix の AI 請求額は、エージェントの指示を LLM に達する前にトークン単位で削減するソフトウェアを開発した同社のシニア エンジニア、テジャス チョプラ氏のおかげで、それほど驚くべきものではないかもしれない。
チョプラ氏は、トークンの 90% が、選択した巨大な思考マシンにとって冗長であると推定しています。
公式の Netflix プロジェクトではありませんが、Netflix のいくつかのチームがすでに Project Headroom を使用しており、多くの外部プロジェクトもそれに依存しています。
先週のオープンソースサミットでの講演で、チョプラ氏は、ヘッドルームがユーザーのために推定70万ドルを節約し、現在合計2,000億のトークンを他の場所で使えるようになったと述べた。
1 月に公開されたばかりのオープンソース アプリケーションとしては悪くありません。 Headroom は現在未加工の v0.22 ですが、GitHub 上で 2,000 個のスターを集めており、120 回以上フォークされています。
「当社のユーザーの多くは、何よりもトークンのコストに本当に苦しんでいる人たちです」とチョプラ氏はプレゼンテーションで述べた。
クロード・ソネットからの 287 ドルの請求書が、最初にチョプラ氏の注目をトークンエコノマイゼーションのアイデアにもたらしました。
請求書は典型的なホーム プロジェクトのものでした。少しのデバッグ、いくつかのリファクタリング、データベースにクエリを実行する MCP ツールです。当時、Claude Sonnet のトークンベースの価格設定は非常に寛大なものに見えました。100 万の入力トークンごとに 3 ドル、またはトークンを使用しない場合は 100 万ドルあたり 6 ドルでした。

コンテキスト ウィンドウのトークン制限が 200,000 であることを確認してください。それでも、その 287 ドルはすぐに増えました。
チョプラ氏はさらに詳しく調査した結果、このデータの多くが LLM にとって非常に冗長であることを発見しました。概して、彼自身の手作りの指示が原因ではありませんでした。むしろ、必要以上に冗長な JSON スキーマ、API 応答内のネストされたテンプレート、同一のデータベース列など、すべてのボイラープレートとマシンのメタデータがそれに伴って登場しました。
「これは散文ではありません。これは創造的な文章ではありません。これはテキストに見せかけた圧縮可能なデータです」とチョプラ氏はソフトウェアを紹介するブログ投稿で書いた。 2025 年、研究者グループは、ユーザー入力の読み取りが全トークン消費量の約 76% を占めていることを発見しました。
モデルプロバイダーには、トークンを保存するための独自のツールがあります。しかし、これまでのところ、これらのツールの設定はエンド ユーザーに対してやや偏ったものになっています。デフォルトでは、Claude のプレフィックス キャッシュ設定はわずか 5 分です。非アクティブ状態が 5 分間続くと、LLM がまったく同じデータを必要とする場合でも、コンテキスト ウィンドウ全体を更新する必要があります。 API ドキュメントには、1 時間の存続時間 (TTL) という別の設定が公開されています。しかし、落とし穴があります。 「読み取りコストを 90% 節約するには、書き込みコストの 2 倍を支払う必要があります」とチョプラ氏は聴衆に語った。スイートスポットを見つけるのはあなた次第です。
また、YCombinator が出資し、トークン圧縮をサービスとして提供する Token Company など、新しい商用トークン バーバーも多数登場しています。オープンソース側には、リポジトリへの呼び出しなどの詳細なコマンドの出力をトリミングする RTK (Rust Token Killer) があります。もう 1 つのオープンソース プロジェクトである LeanCTX は、RTK の亜種です。
これらのツールはすべて便利であるとチョプラ氏は認めましたが、操作を開発者のワークフローに限定するように Headroom を設計しました。そして、それはどれにもない何かを持っていました

アプリやサービスは、可逆圧縮を提供できます。
Headroom の仕事は、ユーザーのコンテキスト ウィンドウに入力されるすべてのソース素材 (会話履歴だけでなく、RAG が役立つと判断したログ、ツール出力、ファイル、ドキュメントの塊など) を、LLM に届く前に圧縮することです。
コンテキスト ウィンドウは、各ユーザー セッションに設定されたスペースです。最新のフロンティア モデルは、入力と出力の両方を保持する 200 万トークンに向けてコンテキスト ウィンドウを急速に拡大しています。
教皇レオが指摘するように、このような寛大さは複合的な祝福である。測定単位として、単一のトークンは人間の単語とほぼ同等です。従量課金制プランの場合、コンテキスト ウィンドウに情報を入力すればするほど、料金も高くなります。
Python および Node 上で実行される Headroom は、エンジニアのコンピュータ上でプロキシ (ポート 8787) として実行されます。ユーザーはコマンド ライン インターフェイスで LLM をラップし (つまり、「ヘッドルーム ラップ コーデックス」)、入力を解析します。
Headroom は多少のプログラミング コードと人間による指示を圧縮しますが、サーバー ログ (90% は廃棄可能)、MCP ツール出力 (70% が冗長な JSON)、データベース出力 (すべて 1 つのスキーマ)、およびファイル ツリー (多くの反復メタデータ) の分割に最適です。
Headroom の最初のステップは、CacheAligner と呼ばれるプロセスです。このプロセスは、すでに入力されている入力内で変更された情報のみを検索し、新しい情報のみを送信します。これにより、AI プロバイダーがユーザーのコンテキスト ウィンドウを保存するキャッシュである KV キャッシュ内のほとんど変更されていないテキストの本文全体を置き換える必要がなくなります。
「システム プロンプトに日付フィールドが含まれているか、セッションごとに変更される UUID が含まれている場合、実質的に毎回キャッシュ ミスが発生することになります」と同氏は聴衆に語った。 「そうなるとコストが跳ね上がります。」
それから、ルート

er プロセスはコンテンツのタイプを推測し、それを多数の圧縮プログラムの 1 つに送信します。抽象構文ツリー (AST) コンプレッサーは、プログラミング コードを圧縮します。 JSON コンプレッサーとドキュメント オブジェクト モデル (DOM) コンプレッサーは、それぞれ不要な JSON と Web ボイラープレートを切り取ります。
Headroom には、テキストまたは JSON 入力を調べ、統計分析に基づいてどのビットが実際に関連しているかを判断するいくつかの「スカッシャー」もあります。これらのツールは、モデルが元の非圧縮プロンプトにコールバックする必要がある頻度に基づいて、圧縮が過剰か過小かをフィードバック ループで学習します。
Compress Cache and Retrieve (CCR) と呼ばれる最後のプロセスは、LLM が元の圧縮されていないデータを参照する機能を提供します。データが圧縮されている場所にマーカーが付けられるため、LLM が元のコンテキストを取得したい場合は、ヘッドルーム MCP を呼び出して、ユーザーのマシンから必要な素材を取得できます。元のコンテキストは Redis または SQLite に保存されます。
チョプラ氏は、このソフトウェアスタックには、特にテストの精度に関して、まだやるべきことが残っていることを認めた。 CCR には元のプロンプトが保存されているため、これは簡単な作業です。財務データなど、他の特定の種類のデータ用にさらに多くのコンプレッサーを構築することもできます。
オーディオ、画像、ビデオにも取り組む必要があります (1 人のユーザーがビデオ解析のためにプロジェクトをすでにフォークしています)。 Chopra 氏によれば、関連プロジェクトは Headlight であり、間もなくオープンソースになる予定だという。ヘッドライトは各トークンの起源を追跡します。これは、マルチモデル作業の精度を確保するのに特に便利です。
研究によると、トークンに気を配ることはお金を節約するだけでなく、結果を向上させることができます。
エージェントは、モデルが使用できる量を超えるコンテキストを送信します。これにより、ユーザーの金庫が空になるだけでなく、LLM が実際に愚かになる可能性があります。
私たちの残りの部分と同様に、LLM も次のことを行います。

あまりにも多くの情報を提示されると混乱する。スタンフォード大学関係者のグループは、LLM がコンテキスト ウィンドウの最初と最後により注意を払い、中間ビットを無視する傾向があることを発見しました。
同様に、データ インテグレーターである Chroma の研究者らは、18 個の LLM にわたって、「入力長が長くなるにつれてパフォーマンスの信頼性がますます低くなっている」と推測しました。
彼らはこの現象を「文脈の腐敗」と呼んだ。
プロンプトをトリミングすると、遅延が改善される場合もあります。チョプラ氏はプレゼンテーションの中で、Headroom のユーザーの 1 人が音声起動アプリケーション用のソフトウェアをフォークした様子を中継しました。音声の場合、沈黙でもトークンを生成できます。ユーザーは、サービスが自然に聞こえるためにアプリから 200 ミリ秒以内に応答することを期待しているため、同社はヘッドルームを使用して、その遅延ウィンドウを可能な限り短縮しています。
また、ヘッドルームは、データセンターのエネルギー使用量によって世界が灼熱の地獄に陥るのではないかと心配している人たちに朗報を提供します。トークンが少ないということは、コンテキスト ウィンドウが小さくなることを意味します。つまり、少なくともジェボンのパラドックスが発生し、人々が猫のアニメーション映画をレンダリングするためにさらに電力を必要とする方法を見つけるまでは、エネルギーの使用量が少なくなります。 ®
AI + ML
Netflix のウィズが AI 費用を削減するアプリを作成し、それをオープンソース化
Project Headroom は多額の費用も節約できる
ウィキメディアの解雇後、ウィキペディア編集者がストライキとバナー破壊行為を計画
コミュニティから要求された多くの修正とモデレーションツールを担当するチームを解散させた後、財団が反乱を引き起こす
Postgres における AI とデータ主権: データセンターのエネルギー危機への答え
10億人のAIエージェントが電力網に乗り込む
国立宇宙センターでのロケット展示が NASA SLS の意図せぬ印象を引き出す
システム
EUのデジタル主権ブーブーは、このプロジェクトに起こった最高の出来事かもしれない
DIY か、死ぬか。ただ

CIAに買わせないでください
AWSはエンタープライズ需要がゼロにもかかわらず、イーロン・マスク氏のGrokをBedrockに押し込むと報じられている
フロンティアモデルのエナジードリンク
AI + ML
Googleは最近AIの機能化に本格的に取り組んでいる
セキュリティ
Anthropic、Mythos クラスのモデルを一般公開へ
セキュリティ
レドモンドが警察に通報、マイクロソフトに「屈辱を受けた」と不満を抱く0日ハンターが「骨が砕けるほどの衝撃」を誓う
AI + ML
Netflix のウィズが AI 費用を削減するアプリを作成し、それをオープンソース化
セキュリティ
軍隊の携帯電話が位置データを外国の敵に提供した
データ主権のトレードオフを克服する
避けられないトレードオフであるデータ主権はネットワークにとって実際に何を意味するのでしょうか?もっと詳しく知る。
プロンプトからエクスプロイトまで: LLM は API 攻撃をどのように変化させているか
最新のアプリケーションは API 主導で相互接続されており、多くの場合、過剰な権限が与えられているため、AI 支援攻撃の理想的な標的となっています。
未来の構築: Kubernetes 向けエンタープライズ データ サービスのロックを解除する
インフラストラクチャのサイロを排除し、標準化されたエンタープライズ グレードのクラウド ネイティブ プラットフォームを確立する方法を発見してください。
Microsoft 365 が見逃す高度な攻撃を Behavioral AI セキュリティで捕捉
Microsoft 365 は企業コミュニケーションのバックボーンであり、そのネイティブ セキュリティが既知のものや騒々しいものを除外します。
ランサムウェア侵害の混乱に足を踏み入れ、対応スキルをテストし、他の IT およびセキュリティの専門家とチームを組んでサイバー犯罪者を出し抜こう
仮想サイバー復旧シミュレーション
ランサムウェア攻撃の勢いは衰えていませんが、私たちも同様です。 Druva の人気イベント「Escape Ransomware」が完全にバーチャルになりました。
大規模なエージェント AI: パイロットから運用まで
AI の導入を大規模に推進することで、真の ROI を実現する方法を学びましょう。
今週のケトルでは、スチームデッキが炭鉱のカナリアであるかどうかについて考えます。

ハードウェア価格の将来と、Blue Origin の爆発が NASA の月ミッションに及ぼす影響
Project Headroom は多額の費用も節約できる
コミュニティから要求された多くの修正とモデレーションツールを担当するチームを解散させた後、財団が反乱を引き起こす
フロンティアモデルのエナジードリンク
今週のケトルでは、スチーム デッキがハードウェア価格の将来にとって炭鉱のカナリアであるかどうか、そして NASA の月ミッションに対する Blue Origin の爆発の影響について考えます。
Project Headroom は多額の費用も節約できる
コミュニティから要求された多くの修正とモデレーションツールを担当するチームを解散させた後、財団が反乱を引き起こす
フロンティアモデルのエナジードリンク
ロケットの爆発とハードウェアの価格の高騰により、ひどい新常態が生まれる
今週のケトルでは、スチーム デッキがハードウェア価格の将来にとって炭鉱のカナリアであるかどうか、そして NASA の月ミッションに対する Blue Origin の爆発の影響について考えます。
Netflix のウィズが AI 費用を削減するアプリを作成し、それをオープンソース化
Project Headroom は多額の費用も節約できる
ウィキメディアの解雇後、ウィキペディア編集者がストライキとバナー破壊行為を計画
コミュニティから要求された多くの修正とモデレーションツールを担当するチームを解散させた後、財団が反乱を引き起こす
国立宇宙センターでのロケット展示が終了

[切り捨てられた]

## Original Extract

Project Headroom could save you big money, too

Jump to main content
Search
TOPICS
Special Features
All Special Features
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
Netflix wiz creates app to slash AI bills, then open sources it
Project Headroom could save you big money, too
As the COOs from both Uber and Microsoft recently learned, encouraging company engineers to use AI aggressively can lead to hefty usage bills, perhaps even offsetting all the gains from laying off employees.
The AI bills at Netflix may not be so eye-popping thanks to company senior engineer Tejas Chopra, who has created software to prune agent instructions, as measured in tokens, before they hit the LLM.
Chopra has estimated that as much as 90% of tokens are redundant to the giant thinking machine of your choice.
Although not an official Netflix project, several teams there already use Project Headroom , and a number of external projects rely on it as well.
In a talk at the Open Source Summit last week, Chopra said that Headroom has saved an estimated $700,000 for its users, who collectively now have 200 billion tokens to spend elsewhere.
Not bad for an open source application that’s been out only since January. Headroom, currently at a still-raw v0.22, has gathered 2,000 stars on GitHub and has been forked over 120 times.
“A lot of our users are people who have been really burned by token costs, more than anything else,” Chopra said in his presentation.
A $287 bill from Claude Sonnet first brought Chopra’s attention to the idea of token economization.
The bill was typical home project stuff: a bit of debugging, some refactoring, MCP tools querying a database. At the time, Claude Sonnet’s token-based pricing seemed pretty generous: $3 for every million input tokens, or $6/million if you went over the 200,000 token limit for your context window. Still, that $287 added up quickly.
Upon deeper inspection, Chopra found a lot of this data was highly redundant to the LLM. By and large, his own hand-crafted instructions were not the culprit. Rather it was all the boilerplate and machine metadata that came along for the ride: Needlessly-verbose JSON schemas, nested templates within API responses, identical database columns.
“This isn’t prose. This isn’t creative writing. This is compressible data masquerading as text,” Chopra wrote in a blog post introducing his software. In 2025, a group of researchers found that reading user input accounted for about 76% of all token consumption.
The model providers have their own tools to save tokens. But to date, the settings on these tools are somewhat oblique to end users. By default, Claude has a prefix cache setting of just five minutes. After five minutes of inactivity, the entire context window needs to be refreshed, even if the LLM needs the exact same data. Another setting is exposed in the API documentation: a one-hour time to live (TTL). But there is a catch. "You pay two times the cost for your writes to get 90% savings for your reads," Chopra told the audience. It’s up to you to find the sweet spot.
There are also a number of new commercial token barbers popping up, such as YCombinator-funded Token Company, which offers token compression as a service. On the open source side there is RTK (Rust Token Killer), which trims to the output of verbose commands, such as calls to a repository. Another open source project, LeanCTX , is a variant of RTK.
All these tools are useful, Chopra admitted, but he designed Headroom to keep the operations confined to the developer’s workflow. And it had something none of the apps and services could offer: reversible compression.
Headroom’s job is to compress all the source material that is fed into the user’s context window – not only the conversation history, but also logs, tool outputs, files, chunks of documentation that the RAG found useful – before it arrives at the LLM.
The context window is the set space for each user session. The latest frontier models are rapidly expanding their context windows upwards towards two million tokens, which holds both input and output.
Such generosity is a mixed blessing, as Pope Leo might point out . As a unit of measurement, a single token is more or less equivalent to a human word. For pay-as-you go plans, the more you feed the context window, the more you’ll pay.
Running on Python and Node, Headroom runs as a proxy (port 8787) on the engineer’s computer. The user wraps their LLM at the command line interface (i.e. “headroom wrap codex”) and it then parses the input.
While Headroom does compress a bit of programming code and human instruction, it is best at chopping server logs (90% of which can be jettisoned), MCP tool outputs (70% redundant JSON), Database outputs (it’s all one schema), and file trees (much repeated metadata).
Headroom’s first step is a process called CacheAligner which looks only for information that has been changed within input that's already been entered, and ships only the new info, eliminating the need to replace an entire body of mostly unchanged text in KV Cache, the cache where the AI provider stores the user’s context window.
“If your system prompt contains a date field or contains some UUID that changes per session, you are effectively getting a cache miss every single time,” he told the audience. “That will blow up your costs.”
Then, a router process infers the type of content and sends it to one of a number of compressors. An Abstract Syntax Tree (AST) compressor squishes programming code. JSON and Document Object Model (DOM) compressors snip unneeded JSON and Web boilerplate, respectively.
Headroom also has some “squashers” that look at text or JSON input and decide which bits are actually relevant, based on statistical analysis. These tools learn in a feedback loop if they are over- or under-compressing, based on how often the model has to call back into the original uncompressed prompt.
The final process, called Compress Cache and Retrieve (CCR), offers that ability for the LLM to look at the original unsquashed data. It puts markers to where the data has been compressed, so if the LLM wishes to get the original context, it can call a Headroom MCP to retrieve the needed material from the user’s machine. The original context is stored on Redis or SQLite.
There is still work to be done to this software stack, Chopra admitted, particularly on testing accuracy. It should be an easy task because the CCR stores the original prompts. More compressors can also be built for other specific types of data, such as financial data.
Audio, image, and video will also have to be tackled (one user has already forked the project for video parsing). A related project, which Chopra says will be open source soon, is Headlight. Headlight will keep track of the origin of each token, which could be especially handy for ensuring the accuracy of multi-model work.
Minding your tokens does not only save money, it can improve results, research suggests.
Agents send more context than the model can possibly use, which, in addition to emptying the user’s coffers, can actually make the LLM dumber.
Like the rest of us, LLMs get confused when presented with too much information. A group of Stanford University boffins found that LLMs tend to pay more attention to the beginning and the end of the context window, and tend to disregard the middle bits.
Likewise, a set of researchers from data integrator Chroma deduced that , across 18 LLMs, “performance grows increasingly unreliable as input length grows.”
“Context rot,” they called this phenomenon.
Trimming prompts can also improve latency. In his presentation, Chopra relayed how one of Headroom’s users forked the software for a voice-activated application. With voice, even silence can generate tokens. The user expects a response from the app within 200 milliseconds for the service to sound natural, so the company is using Headroom to help shrink that latency window down as much as possible.
Headroom also offers some good news for those worrying about data centers heating the world into a fiery inferno with their energy usage. Fewer tokens means a smaller context window, which means less energy use – at least until Jevon's Paradox kicks in and people find even more power-hungry ways to render their animated cat movies. ®
AI + ML
Netflix wiz creates app to slash AI bills, then open sources it
Project Headroom could save you big money, too
Wikipedia editors plot strike and banner sabotage after Wikimedia layoffs
Foundation sparks revolt after disbanding team responsible for many community-requested fixes and moderation tools
AI and data sovereignty in Postgres: An answer to the datacenter energy crisis
A billion AI agents walk into a power grid
Rocket exhibit at National Space Centre pulls off unintentional NASA SLS impression
Systems
EU's digital sovereignty boo-boo may be the best thing to ever happen to the project
DIY or die. Just don't let the CIA buy it
AWS reportedly to tuck Elon Musk's Grok into Bedrock, despite zero enterprise demand
The energy drink of frontier models
AI + ML
Google has seriously leaned into AI enshittification lately
Security
Anthropic to release Mythos-class models to the public
Security
Disgruntled 0-day hunter 'humiliated' by Microsoft pledges 'bone shattering drop' as Redmond calls cops
AI + ML
Netflix wiz creates app to slash AI bills, then open sources it
Security
Troops’ phones gave away location data to foreign adversaries
Overcoming the trade-offs in data sovereignty
What does data sovereignty actually mean for your network, which trade-offs are unavoidable? Learn more.
From Prompt to Exploit: How LLMs Are Changing API Attacks
Modern applications are API-driven, interconnected, and often over-permissioned, making them an ideal target for AI-assisted attacks.
Architecting the Future: Unlocking Enterprise Data Services for Kubernetes
Join us to discover how to eliminate infrastructure silos and establish a standardized, enterprise-grade cloud-native platform.
Catch the Advanced Attacks Microsoft 365 Misses with Behavioral AI Security
Microsoft 365 is the backbone of enterprise communication, and its native security filters out the known and the noisy.
Step into the chaos of a live ransomware breach, test your response skills, and team up with other IT and security pros to outsmart cybercriminals
Virtual Cyber Recovery Simulation
Ransomware attacks aren’t slowing down, and neither are we. Druva’s hit event, Escape Ransomware, is now fully virtual.
Agentic AI at Scale: From Pilot to Production
Join us to learn how to unlock real ROI by driving adoption of AI at scale.
This week on The Kettle, we mull over whether the Steam Deck is a canary in the coal mine for the future of hardware prices, and the effect of Blue Origin's blowout on NASA's Moon missions
Project Headroom could save you big money, too
Foundation sparks revolt after disbanding team responsible for many community-requested fixes and moderation tools
The energy drink of frontier models
This week on The Kettle, we mull over whether the Steam Deck is a canary in the coal mine for the future of hardware prices, and the effect of Blue Origin's blowout on NASA's Moon missions
Project Headroom could save you big money, too
Foundation sparks revolt after disbanding team responsible for many community-requested fixes and moderation tools
The energy drink of frontier models
Exploding rockets and exploding hardware prices make for a lousy new normal
This week on The Kettle, we mull over whether the Steam Deck is a canary in the coal mine for the future of hardware prices, and the effect of Blue Origin's blowout on NASA's Moon missions
Netflix wiz creates app to slash AI bills, then open sources it
Project Headroom could save you big money, too
Wikipedia editors plot strike and banner sabotage after Wikimedia layoffs
Foundation sparks revolt after disbanding team responsible for many community-requested fixes and moderation tools
Rocket exhibit at National Space Centre pulls off

[truncated]
