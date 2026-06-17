---
source: "https://arstechnica.com/ai/2026/06/ai-coding-agents-can-autonomously-direct-robot-training/"
hn_url: "https://news.ycombinator.com/item?id=48575722"
title: "AI coding agents taught robots how to install GPUs and cut zip-ties"
article_title: "AI coding agents taught robots how to install GPUs and cut zip ties - Ars Technica"
author: "pseudolus"
captured_at: "2026-06-17T21:43:29Z"
capture_tool: "hn-digest"
hn_id: 48575722
score: 2
comments: 0
posted_at: "2026-06-17T19:39:07Z"
tags:
  - hacker-news
  - translated
---

# AI coding agents taught robots how to install GPUs and cut zip-ties

- HN: [48575722](https://news.ycombinator.com/item?id=48575722)
- Source: [arstechnica.com](https://arstechnica.com/ai/2026/06/ai-coding-agents-can-autonomously-direct-robot-training/)
- Score: 2
- Comments: 0
- Posted: 2026-06-17T19:39:07Z

## Translation

タイトル: AI コーディング エージェントがロボットに GPU の取り付け方と結束バンドの切断方法を教えた
記事タイトル: AI コーディング エージェントがロボットに GPU の取り付け方と結束バンドの切断方法を教えた - Ars Technica
説明: Nvidia のロボット自己改善プログラムでは、AI コーディング エージェントのチームが参加しています。

記事本文:
コンテンツにスキップ
アルス テクニカ ホーム
セクション
フォーラム
購読する
検索
AI
ストーリーテキスト
サイズ
小
標準
大
幅
*
標準
ワイド
リンク
標準
オレンジ
※購読者限定
さらに詳しく
ストーリーにピンを留める
テーマ
ハイパーライト
検索
サインイン
サインインダイアログ...
サインイン
ロボットを訓練する方法
AI コーディング エージェントがロボットに GPU の取り付け方法と結束バンドの切断方法を教えました
Nvidia のロボット自己改善プログラムでは、AI コーディング エージェントのチームが協力しています。
23
AI コーディング エージェントのチームは、マザーボードへの GPU の挿入などのタスクを実行するようにロボットを訓練できます。
クレジット:
エヌビディア
AI コーディング エージェントのチームは、マザーボードへの GPU の挿入などのタスクを実行するようにロボットを訓練できます。
クレジット:
エヌビディア
テキスト
設定
ストーリーテキスト
サイズ
小
標準
大
幅
*
標準
ワイド
リンク
標準
オレンジ
※購読者限定
さらに詳しく
ナビゲーション用に最小化する
AI コーディング エージェントに、ロボット アーム、いくつかのコンピューティング リソース、およびロボットにさまざまなタスクを教えるための「豊富なトークン予算」が満載されたラボを提供するとどうなるでしょうか?エージェントは、ロボットに結束バンドをうまく切断したり、マザーボード上の薄いソケットに GPU を挿入したりする方法を教えるトレーニング計画を見つけ出すことができるようです。
AI が完全に自律的な方法で動作してロボットのトレーニングを自動化する方法を垣間見ることができるのは、新しいエージェント ハーネス フレームワークによって可能になりました。これは、AI モデルをラップしてさまざまなツールを使用できるようにすると同時に、メモリ、コンテキスト、制約、フィードバック ループなどの機能も提供するソフトウェアです。 ENPIRE と呼ばれるそのエージェント ハーネスは、Nvidia GEAR (Generalist Embodied Agent Research) ラボのロボット研究者と、ピッツバーグのカーネギー メロン大学およびカリフォルニア大学バークレー校の共同研究者によって開発されました。
「当社の NVIDIA GEAR ラボの一部は現在、一晩中たゆまぬ自己改善を行っています」と AI ディレクターの Jim Fan 氏は書いています。

NVIDIA で、LinkedIn の投稿で。 「私たちは午前中にレポートを読んだところです。」
ファン氏はまた、このような AI 主導のロボット トレーニングの目標について、NVIDIA の創設者兼 CEO のジェンセン フアン氏を引き合いに出し、「私たちは皆休暇を取っていても、ジェンセンは気付かないだろう」と冗談めかして説明しました。しかし、恩恵を受けるのは Nvidia ロボット研究者だけではありません。ファン氏は、チームはすべてをオープンソース化し、誰でも自宅で自分の「自走ロボット ラボ」を開催できるようにすると述べました。
ENPIRE ハーネスには 4 つのモジュールがあり、AI コーディング エージェントがタスクの自動リセットと検証を実行し、ロボットの動作をガイドするポリシーを洗練し、並行して動作する複数の物理ロボット全体でそのようなポリシーを評価し、ログの分析、研究論文の取り込み、トレーニング インフラストラクチャとアルゴリズム コードの改善によって障害に対処できるようにします。技術的な詳細については、2026 年 6 月 16 日にアップロードされた研究論文をご覧ください。
このハーネスは、OpenAI の Codex (GPT-5.5)、Anthropic の Claude Code (Opus 4.7)、Moonshot AI の Kim Code (Kimi K2.6) を含む 3 つの異なる AI コーディング エージェントでテストされました。コーディング エージェントのチームは、ロボット トレーニングに対するさまざまなアルゴリズム アプローチを独自に開発し、実世界の実験でテストし、自主的なテストの繰り返しサイクルにわたって全体の成功率を上げるのに役立つ変更はすべて保持しました。
AI主導のロボットトレーニングの成功と限界
ENPIRE を搭載した AI コーディング エージェントは、ロボットの自己改善のための戦略を開発しました。その戦略は、テーブル上の目標位置に合わせて T 字型のブロックを動かすようロボットに要求する標準的な「Push-T」タスクを含む、いくつかの操作タスクで 99% の成功率を達成しました。その他の作業には、ピンボックス内のピンの整理、結束バンドの結び方と切断、

グラフィックス カードを再度取り外して次の試用に向けてリセットする前に、GPU をマザーボードに取り付けます。
最も有望な結果は、ピンの挿入と編成タスクから得られたものと考えられます。そのロボット トレーニング シナリオでは、AI コーディング エージェントは、同じ人間の研究者の多くが開発した「最先端のヒューマンインザループ手法」よりも早く、ほぼ 100% の成功を達成しました。
このような実験では、最大 8 人の AI コーディング エージェントからなる大規模なチームが、4 人のエージェントからなる小規模なチームや単独で作業する 1 人のエージェントよりも迅速にロボット トレーニングで高い成功率を達成できることも示しました。たとえば、8 エージェント チームは 2 時間の調査時間で Push-T タスクで 99% の成功率を達成しました。これに対し、4 エージェント チームは 3 時間、単一エージェント チームは 5 時間近くかかりました。
しかし、人間の研究者たちは、AI コーディング エージェントを自律型ロボット トレーナーとして解放する際に、いくつかの重大な制限があることも発見しました。コーディングエージェントが「ログを読んだり、コードを書いたり、デバッグしたり、言語モデルのバックボーンを待ったり」するのに忙しい間、ロボットはアイドル状態で使用されないことがよくありました。また、コーディング エージェントのチームが大規模になると、互いのアイデアの要約に多くの時間が費やされ、実際にロボットを使用する時間が減り、コーディング エージェントは並行トレーニング セッションを開始するときに利用可能なコンピューティング リソースを最大限に活用できないことがありました。
より多くのエージェントとロボットが連携することによって実現される成功率の向上には、トークン消費量の増加という代償も伴いました。これは、Anthropic などの AI 開発者が、AI サービスの使用に伴うトークン関連コストの大幅な増加につながる価格変更を検討している現在、注目に値する考慮事項です。
AI ブームで潤った Nvidia は、複数のロボット工学の取り組みを通じて、物理 AI のビジョンを精力的に推進してきました。

同社は5月31日、汎用AI搭載ロボットを開発する研究機関向けに「リファレンス・ヒューマノイド・ロボット」を提供するため、中国の著名ロボット企業Unitreeと提携すると発表した。
6月初旬に韓国をめまぐるしく訪問した際、エヌビディアの創設者兼最高経営責任者（CEO）のジェンスン・ファン氏は、現代自動車の執行会長チョン・ウィソン氏とも会談し、AI搭載ロボットの量産規模の拡大について話し合った。現代自動車グループは米国のロボット企業ボストン・ダイナミクスを所有しており、同社はすでに4本足の「ロボット犬」スポットでよく知られており、同社の人型ロボット「アトラス」の商品化に取り組んでいる。
23件のコメント
コメント
フォーラムビュー
コメントを読み込んでいます...
前の話
次の話
よく読まれている
1.
Anthropic、Claude Agent SDK のトークンベースの課金を「一時停止」
2.
コモドールの最新ガジェットはソーシャルメディアとブラウザをブロックする折りたたみ式携帯電話です
3.
アマゾンが期待していた大型新型ロケットのうち、納入したのはヨーロッパだけ
4.
流出した財務文書によると、OpenAIは年間数十億ドルの損失を出している
5.
Windows および Linux ユーザー: セキュア ブート キーの更新期限が近づいています
カスタマイズ
Ars Technica は信号を分離してきました。
25年以上続く騒音。弊社独自の組み合わせにより、
技術的な知識と技術芸術への幅広い関心
Ars は、情報の海の中で信頼できる情報源です。後
すべてを知る必要はありません。重要なことだけを知っておく必要があります。

## Original Extract

Nvidia's self-improvement program for robots enlists teams of AI coding agents.

Skip to content
Ars Technica home
Sections
Forum
Subscribe
Search
AI
Story text
Size
Small
Standard
Large
Width
*
Standard
Wide
Links
Standard
Orange
* Subscribers only
Learn more
Pin to story
Theme
HyperLight
Search
Sign In
Sign in dialog...
Sign in
How to train your robot
AI coding agents taught robots how to install GPUs and cut zip ties
Nvidia’s self-improvement program for robots enlists teams of AI coding agents.
23
Teams of AI coding agents can train robots to do tasks such as inserting GPUs into motherboards.
Credit:
NVIDIA
Teams of AI coding agents can train robots to do tasks such as inserting GPUs into motherboards.
Credit:
NVIDIA
Text
settings
Story text
Size
Small
Standard
Large
Width
*
Standard
Wide
Links
Standard
Orange
* Subscribers only
Learn more
Minimize to nav
What happens when you give AI coding agents a lab full of robotic arms, some compute resources, and a “generous token budget” for teaching the robots various tasks? The agents can apparently figure out a training regimen that teaches the robots to successfully cut zip ties and even insert GPUs into thin sockets on motherboards.
That glimpse into how AI can act in a fully autonomous way to automate robot training was made possible by a new agent harness framework—software that wraps around AI models to enable their use of various tools while also providing capabilities such as memory, context, constraint, and feedback loops. That agentic harness, called ENPIRE , was developed by robotics researchers at the Nvidia GEAR (Generalist Embodied Agent Research) lab alongside collaborators from Carnegie Mellon University in Pittsburgh and the University of California, Berkeley.
“A part of our NVIDIA GEAR lab now self-improves tirelessly overnight,” wrote Jim Fan, director of AI at NVIDIA, in a LinkedIn post . “We just read the reports in the morning.”
Fan also jokingly described the goal of such AI-directed robot training, saying, “We all take a holiday and Jensen wouldn’t even notice,” in reference to Nvidia founder and CEO Jensen Huang. But it’s not only Nvidia robotics researchers who could benefit—Fan said the team would be open-sourcing everything so anyone can host their own “self-running robot lab at home.”
The ENPIRE harness has four modules that enable AI coding agents to perform automatic reset and verification on tasks, refine policies that guide robotic behavior, evaluate such policies across multiple physical robots working in parallel, and address failures by analyzing logs, ingesting research papers, and improving training infrastructure and algorithm code. More technical details are available in the research paper uploaded on June 16, 2026.
The harness was tested with three different AI coding agents, including OpenAI’s Codex with GPT-5.5, Anthropic’s Claude Code with Opus 4.7, and Moonshot AI’s Kimi Code with Kimi K2.6. Teams of the coding agents independently developed different algorithmic approaches to robot training, tested them in real-world experiments, and then retained whatever changes helped raise the overall success rate over repeated cycles of self-directed testing.
The success and limits of AI-directed robot training
Equipped with ENPIRE, the AI coding agents developed strategies for robotic self-improvement that achieved a 99 percent success rate across several manipulation tasks, including the standard “Push-T” task that challenges robots to move a T-shaped block to fit a target position on top of a table. Other tasks included organizing pins in a pin box, tying and cutting zip ties, and placing a GPU into a motherboard before unplugging the graphics card again to reset for the next trial.
The most promising result may have come from the pin insertion and organization task. In that robot-training scenario, AI coding agents achieved nearly 100 percent success faster than a “ frontier human-in-the-loop method” developed by many of the same human researchers.
Such experiments also showed how larger teams of up to eight AI coding agents could achieve high success rates in robot training more quickly than smaller four-agent teams or single agents working alone. For example, the eight-agent team achieved 99 percent success on the Push-T task in two hours of research time, compared to the four-agent team requiring three hours and the single-agent team requiring nearly five hours.
But the human researchers also discovered some crucial limitations when unleashing AI coding agents as autonomous robot trainers. The robots often sat idle and unused while the coding agents were busy “reading logs, writing code, debugging, or waiting for the language-model backbone.” Larger teams of coding agents also spent more time summarizing each other’s ideas and less time actually using the robots, and the coding agents sometimes failed to make full use of available compute resources when launching parallel training sessions.
The faster success rates enabled through more agents and robots working together also came at the cost of higher token consumption—a noteworthy consideration at a time when AI developers such as Anthropic are weighing pricing changes that would significantly increase the token-related costs of using AI services.
Flush with cash from the AI boom, Nvidia has been busily pushing its vision for physical AI through multiple robotics initiatives. On May 31, the company announced a partnership with the prominent Chinese robotics company Unitree to provide a “Reference Humanoid Robot” for research labs developing general-purpose AI-powered robots.
During a whirlwind tour of South Korea in early June, Nvidia founder and CEO Jensen Huang also met with Hyundai Motor Executive Chair Chung Euisun to discuss scaling up the mass manufacturing of AI-powered robots . Hyundai Motor Group owns the US robotics company Boston Dynamics, which is already well-known for its four-legged “robot dog” Spot and has been working to commercialize its Atlas humanoid robot.
23 Comments
Comments
Forum view
Loading comments...
Prev story
Next story
Most Read
1.
Anthropic "pauses" token-based billing for its Claude Agent SDK
2.
Commodore’s newest gadget is a flip phone that blocks social media and browsers
3.
Among the large new rockets Amazon was counting on, only Europe has delivered
4.
Leaked financial docs show OpenAI is losing billions of dollars a year
5.
Windows and Linux users: The deadline to update Secure Boot keys is near
Customize
Ars Technica has been separating the signal from
the noise for over 25 years. With our unique combination of
technical savvy and wide-ranging interest in the technological arts
and sciences, Ars is the trusted source in a sea of information. After
all, you don’t need to know everything, only what’s important.
