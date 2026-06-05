---
source: "https://research.ibm.com/blog/granite-libraries-project-switch"
hn_url: "https://news.ycombinator.com/item?id=48417645"
title: "Granite Switch: Building AI more like software"
article_title: "Introducing Granite Libraries and Project Granite Switch - IBM Research"
author: "yangikan"
captured_at: "2026-06-05T21:35:35Z"
capture_tool: "hn-digest"
hn_id: 48417645
score: 1
comments: 0
posted_at: "2026-06-05T20:19:44Z"
tags:
  - hacker-news
  - translated
---

# Granite Switch: Building AI more like software

- HN: [48417645](https://news.ycombinator.com/item?id=48417645)
- Source: [research.ibm.com](https://research.ibm.com/blog/granite-libraries-project-switch)
- Score: 1
- Comments: 0
- Posted: 2026-06-05T20:19:44Z

## Translation

タイトル: Granite Switch: ソフトウェアに近い AI の構築
記事のタイトル: Granite ライブラリとプロジェクト Granite Switch の紹介 - IBM Research
説明: IBM は、LLM にソフトウェア エンジニアリングの厳密さとモジュール性をもたらします。

記事本文:
2026 年 6 月 4 日リリース 6 分で読む ソフトウェアに近い AI の構築
ソフトウェア エンジニアリングの厳密さとモジュール性を LLM にもたらす 2 つの新しいツールである Granite Libraries と Project Granite Switch を紹介します。
Granite Switch : アダプター機能をデプロイ可能なモデルに構成する
Granite ライブラリ: RAG、安全性、およびコア機能のアダプター機能を参照してダウンロードします
Mellea : アダプター関数を型指定された Python 関数として呼び出す
Granite 4.1 モデル : すべてを強化する基本モデル
最先端の AI が、あなたが求めていたものに似た画像を生成するようになったのは、それほど昔のことではありません。この数年間で、私たちはカンブリア紀の爆発的な機能を目の当たりにしてきました。AI は、完璧なテキストを生成したり、ミッション クリティカルなエンタープライズ ワークフローを実行したり、エージェントを調整してアプリ全体を自律的に実行したりできるようになりました。
しかし、これらすべての進歩にもかかわらず、AI モデルを構築して操作する方法は、他のソフトウェアとはまったく異なります。企業ユーザーはモデルができるだけ正確で効率的であることを望んでいますが、多くの場合、そこに到達することはほとんどの開発者にとって手の届かないところにあります。その理由の 1 つは、これまで AI モデルを従来のソフトウェアに使用するのと同じ種類のプラグアンドプレイのビルディング ブロックに分割することが非常に困難だったことです。
最新のソフトウェア アプリケーションは、1 つの巨大な粘土ブロックではなく、多くのレゴ ブロックで構成されるオブジェクトのように、連携して動作する小さな自己完結型の部分で構成されています。何かが壊れた場合、エンジニアは、コードベースの残りの部分には手を触れずに、問題のモジュールを見つけて修正し、テストし、再デプロイできます。機能はさまざまなインターフェイスの背後で分離されており、アプリの一部はさまざまなチームによって開発され、仕様に照らしてテストされ、必要に応じて置き換えられます。

でした。アプリケーションは 1 つの目的を果たすことができますが、その内部は未分化の塊ではありません。
今日の LLM は現代の驚異であり、収益報告書を解読するのと同じくらい簡単に、各国の首都に関する質問に答えることができます。ただし、すべての機能はパラメーターの重みのコレクション全体に分散されます。特定の状況に対するモデルの反応方法を変更するには、全体を再トレーニングするか、非常に詳細かつ正確なプロンプトを作成する必要があります。これらはいずれも迅速な解決策ではなく、ソフトウェア システムを改善するのと同じように、複数のチームが協力して AI モデルを改善することもできません。
IBM Research は、ジェネレーティブ コンピューティングと呼ばれるアプローチを通じて、LLM にソフトウェア エンジニアリングの厳密さとモジュール性をもたらすテクノロジーに取り組んできました。
IBM Researchの言語およびマルチモーダル・モデル担当ディレクター、ルイス・ラストラス氏は、「モデルはデータを含む単なるコードであり、コードよりもデータの方がはるかに多いだけです」と述べています。 「私たちは LLM 用のソフトウェアの教訓を学んでいません。部分を個別に構築できるのです。」
IBM は、ジェネレーティブ コンピューティングのビジョンに私たちを近づける一連の連携ツールを発表します。ソフトウェア モジュール性のアイデアは、特殊なタスクに合わせて AI モデルをカスタマイズできるアダプターのコレクションである Granite Libraries のきっかけとなりました。これにより、モデル全体を再トレーニングすることなく、モデルが対象のタスクを迅速に実行できるようになります。中心となるアイデアは、ソフトウェア ライブラリの関数のように、入力と出力が定義された「アダプター関数」です。
このコンテキストにおけるアダプター関数は、従来のモデルとは異なる種類の出力を生成するようにトレーニングされた小規模なモデル アダプターです。これらのアダプター関数は、自由形式のテキストを生成する代わりに、文書の関連性のスコアリングや q の書き換えなど、特定のタスクを実行します。

幻覚を検出したり、安全性を判断したりするのに役立ちます。
チームはまた、Granite ライブラリにある特殊なコンポーネントを動的に管理できるようにする既存のモデル アーキテクチャ用のツールキットである Project Granite Switch も導入しています。最近リリースされた Granite 4.1 のモデルと、IBM のジェネレーティブ コンピューティング用オープンソース ライブラリである Mellea と組み合わせることで、開発者は予測不可能なテキスト生成を信頼性の高い決定論的なプログラミング関数に変えるツールを手に入れることができます。
Granite Libraries は、ソフトウェアを非常に強力にしたのと同じ種類のカスタマイズを AI モデルにもたらすように設計されています。
IBM は、一般的な企業ワークフローをサポートするように設計された 3 つのライブラリをリリースしました。 RAG ライブラリには、クエリの書き換え、応答性の評価、幻覚の検出、引用の生成など、主要な検索拡張生成タスク用のアダプターが含まれています。コア ライブラリは、要件チェック、確実性スコアリング、コンテキスト アトリビューションなどの基本的な機能を提供します。このリリースの締めくくりとして、Guardian Library により、別個のガードレール モデルを使用せずに、モデルがインラインの安全性、事実性、ポリシー チェックを直接実行できるようになります。これらの Granite ライブラリは、すべての Granite 4.1 モデルで利用できます。
これらのライブラリはモジュール式で独立してトレーニングされているため、企業は必要に応じてそれらを採用し、今日のソフトウェアの依存関係の管理方法とよく似た機能を段階的に追加できます。
各アダプター機能は、1 つのタスクのエキスパートになるようにトレーニングされています。たとえば、要件チェッカーはモデル応答と一連の制約を受け取り、制約が満たされているかどうかを返します。 Granite 4.1 3B がこれを実行するよう明示的に要求されると、命令フォローイングの一般的なベンチマークである IFEval で 51% のバランスの取れた精度を達成します。同じメートルなら

odel には新しい Granite Library 要件チェック アダプター機能が搭載されており、その精度は 84% に跳ね上がります。
アダプターを使用すると、小規模モデルは、注意深いプロンプトのみによる基本モデルよりも、特定のタスクにおいて比較的優れた性能を発揮します。そして、Mellea は、これらのアダプター関数がソフトウェアのように動作できるようにするものです。Mellea は、特定のアダプターをアクティブにするために必要なタグを自動的に挿入し、フォーマット ルールをリアルタイムで厳密に適用し、すべてを標準の Python 関数としてパッケージ化します。これにより、生の AI テキストの予測不可能な性質からメイン アプリケーションが隔離されます。
Granite ライブラリを使用すると、Granite 基本モデルは、ソフトウェア インターフェイスを通じて明確に定義された機能を実行するようにトレーニングされた、タスク固有のエキスパート (低ランク アダプター (LoRA) またはアクティブ化された LoRA (aLoRA)) を呼び出すことができます。これにより、より小さなモデルでも、より低い推論コストで、大規模なジェネラリスト モデルと同等の狭いタスクを実行できるようになります。
ライブラリ アダプターがアクティブな場合、モデルはそのタスクにおいてひたすら優秀になることができます。基本モデルは変更されませんが、aLoRA の切り替えにほとんどコストをかけずに、その動作を必要に応じて正確に規定できるようになりました。
以下では、Granite ライブラリを使用するとこのプロセスがどれほど高速になるかを示します。このインタラクティブは、同じマルチステップ RAG パイプライン上で 2 つの Granite Switch チェックポイント (1 つは aLoRA を使用し、もう 1 つは標準 LoRA を使用) のベンチマークをアニメーションでリプレイしたものです。レースを加速または減速するには、ボックスの左上にあるさまざまな速度をクリックします。
20 分程度の余裕があれば、Colab でこのレースを再現することもできます。 Colab は通常 1 つの GPU しか提供しないため、2 つのサーバーは順番に実行されますが、生成されたリプレイにより、あたかも同時にレースを行ったかのようにテレメトリがつなぎ合わされます。
アーキテクチャ層: Pr

ジェクトグラナイトスイッチ
Project Granite Switch は、GitHub で利用できる新しい実験的なツールキットで、コンパイラがソース コードやソフトウェア ライブラリからバイナリを生成する方法と同様に、数分で新しいモデルを作成するために使用できます。
Granite Switch により、基本の Granite モデルとそのア​​ダプター機能が、推論時に効率的にアクティブ化される単一のモデルとして機能できます。これは、既存のコア Granite モデルに新しい「スイッチング」レイヤーを追加し、次にアダプターの重みをベース モデルに接着し、フォーマット タグを追加し、新しいチャット テンプレートを準備することによって実現されます。 Granite Switch は、さまざまなタスクごとにまったく新しい AI モデルを起動する必要がなく、必要なときに正確に適切なアダプターを動的にオン/オフします。基本モデルは Granite Switch 内で引き続きアクセスできます。つまり、基礎となるモデルを一切変更せずに新しい機能を利用できます。
この独立したスイッチング層により、LoRA と aLoRA の両方を、大規模展開用のオープンソース推論エンジンである vLLM 内で実行できるようになります。現実の世界では、通常、単一のビジネス タスクには、安全性チェックの実行、データの検索、回答の検証などの一連のアクションが必要です。また、異なるアダプターを切り替えると、AI は短期記憶をクリアし、各ステップですべてを最初から再計算する必要があり、処理速度が低下します。アクティブ化された LoRA を使用することにより、Granite Switch モデルは再読み取りのために一時停止することなくメモリを 1 つのステップから次のステップに転送できるため、複数ステップのワークフローが劇的に高速化されます。
追加のトランスフォーマー層を基本モデルに挿入し、そのアテンション メカニズムを利用してアクティブなアダプターの状態に関連する値を読み取り、保存することにより、エキスパートを切り替える時期が来たときに特別な制御トークンがモデルに信号を送ります。それは

オーケンは、鉄道操車場でどの列車がどこに行くかを指示する指令員のように機能し、スイッチング層がレール自体として機能します。
Granite 4.1: IBM のこれまでで最も高性能なモデル
Granite ライブラリと Project Granite Switch の可能性は、その上で実行できる有能なモデルがなければ失われます。 Granite 4.1 により、IBM は最近、これまでで最もパフォーマンスの高いモデル スイートをリリースしました。
Granite 4.1 ファミリは、その重量クラスを上回る性能を発揮するように設計されており、8B モデルは以前の Granite 32B 専門家混合 (MoE) モデルのパフォーマンスと同等またはそれを上回り、30B モデルはエンタープライズ タスクにおいて Llama 3.3 70B などの大幅に大型のモデルと競合します。他のタスクにすぐに適応できる小型でパフォーマンスの高いモデルは、狭いタスクには対応できない可能性がある大規模で汎用的なモデルよりも、はるかにコストが低くなります。
比較的少量の高品質データでトレーニングすることにより、これらのモデルは、多くのフロンティア推論モデルよりも低いレイテンシーと運用コストを維持しながら、ツールの呼び出しと命令のフォローにおいて競争力のあるスコアを達成しました。
このリリースは、業界をリードする表とグラフの抽出のための Granite Vision 4.1 や新しい音声モデルとガードレール モデルを含む、より広範なエコシステムの一部です。これらはすべて約 15 兆のトークンでトレーニングされており、オープンソースの Apache 2.0 ライセンスの下でリリースされており、世界的な展開に向けて 12 の主要言語をサポートしています。
IBM は、AI モデルをソフトウェアと同じように構成可能にして、企業ユーザーにより多くの価値を提供するという広範な目標の一環として、Granite ライブラリを導入しました。機能をモジュール式コンポーネントに分割することで、開発者は、適応しやすく、運用コストが低く、本番環境での予測が容易な AI システムを構築できます。
モジュール化によって生成 AI の導入におけるすべての課題が解決されるわけではない

大規模ではありますが、より持続可能でエンタープライズ対応のシステムへの実用的な道を提供します。
Granite Switch : アダプター機能をデプロイ可能なモデルに構成する
Granite ライブラリ: RAG、安全性、およびコア機能のアダプター機能を参照してダウンロードします
Mellea : アダプター関数を型指定された Python 関数として呼び出す
Granite 4.1 モデル : すべてを強化する基本モデル
ニュースレターを購読する ホーム
量子の未来がNY Tech Weekで注目を集める
ニュース ピーター・ヘス 2026年6月4日 AI
著名な数学者サブハッシュ・コート氏がIBM Researchに入社
ニュース マイク・マーフィー 2026 年 6 月 1 日 数理科学
IBM、MITとの新たな研究の道筋を描く
Q&A キム・マルティノー 2026 年 5 月 11 日 発見の加速
高速レースとコンピューティングの最前線が出会う場所
研究キム・マルティノーとデイブ・モッシャー 2026 年 4 月 30 日 発見の加速

## Original Extract

IBM is bringing the rigor and modularity of software engineering to LLMs.

04 Jun 2026 Release 6 minute read Building AI more like software
Introducing Granite Libraries and Project Granite Switch, two new tools that bring the rigor and modularity of software engineering to LLMs.
Granite Switch : Compose adapter functions into a deployable model
Granite Libraries : Browse and download adapter functions for RAG, safety, and core capabilities
Mellea : Call adapter functions as typed Python functions
Granite 4.1 models : The base models that power it all
It wasn’t that long ago that the state of the art in AI was generating you an image that sort of looked like the thing you’d asked for. In the interceding years, we’ve seen a Cambrian explosion of capabilities, with AI that can now generate flawless text, run mission-critical enterprise workflows, or orchestrate agents to run entire apps autonomously.
But even with all these advances, the ways we build and interact with AI models is quite different than any other piece of software. Enterprise users want their models to be as accurate and efficient as possible, but getting there is often out of reach for most developers. One reason is that up until now, it’s been incredibly difficult to break AI models down into the same kinds of plug-and-play building blocks we use for traditional software.
A modern software application is made up of smaller, self-contained pieces that work together, rather like an object composed of many LEGO bricks, instead of one giant block of clay. When something breaks, an engineer can find the offending module, fix it, test it, and redeploy — without touching the rest of the codebase. Capabilities are separated behind various interfaces, and part of the app can be developed by different teams, tested against their specifications, and replaced as needed. The application may serve one purpose, but its internals are not an undifferentiated mass.
Today's LLMs are modern marvels, able to answer questions about the capitals of countries as easily as they can decipher an earnings report. But every capability is diffused across the entire collection of parameter weights. To change the way the model reacts to a given situation, you either need to retrain the whole thing, or write extremely detailed and accurate prompts. None of these are quick solutions, and none of them allow multiple teams to work together to improve AI models just like they improve software systems.
IBM Research has been working on technologies that bring the rigor and modularity of software engineering to LLMs through an approach called generative computing .
“Models are just code with data, just a lot more data than code,” said Luis Lastras, IBM Research’s director of language and multimodal models. “We haven’t learned the lessons of software for LLMs — we can build pieces separately.”
IBM is launching a set of coordinated tools that bring us closer to the vision of generative computing. The idea of software modularity was the spark for Granite Libraries , a collection of adapters that can customize AI models for specialized tasks. It enables a model to quickly execute targeted tasks without having to retrain the entire model. The core idea is the “adapter function,” which has a defined input and output, like a function in a software library.
An adapter function in this context is a small model adapter that’s trained to generate a different type of output than a traditional model. Instead of producing open-ended text, these adapter functions carry out a specific task, whether that’s scoring a document for relevance, rewriting a query, detecting a hallucination, or making a safety decision.
The team is also introducing Project Granite Switch , a toolkit for existing model architectures that enables them to dynamically manage the specialized components found in Granite Libraries. Coupled with the recently released models in Granite 4.1 , and Mellea , IBM’s open-source library for generative computing, developers now have a tool that turns unpredictable text generation into a reliable, deterministic programming function.
Granite Libraries is designed to bring the same kind of customization to AI models that has made software so powerful.
IBM has released three libraries designed to support common enterprise workflows. The RAG Library includes adapters for key retrieval-augmented generation tasks such as query rewriting, answerability assessment, hallucination detection, and citation generation. The Core Library provides foundational capabilities, including requirement checks, certainty scoring, and contextual attribution. Rounding out the release, the Guardian Library enables models to perform in-line safety, factuality and policy checks directly, without a separate guardrail model. These Granite Libraries are available for all Granite 4.1 models.
Because these libraries are modular and independently trained, enterprises can adopt them as needed and add more capabilities incrementally, a bit like how software dependencies are managed today.
Each adapter function is trained to be an expert at one task. The requirement checker, for example, takes a model response and a set of constraints, and returns whether the constraints are satisfied. When Granite 4.1 3B is explicitly prompted to do this, it achieves 51% balanced accuracy on IFEval , a popular benchmark for instruction following. If the same model is equipped with the new Granite Library requirement-check adapter function, its accuracy jumps to 84%.
The adapter makes a small model comparatively better at a specific task than the base model could be through careful prompting alone. And Mellea is what allows these adapter functions to act like software: It automatically inserts the tags needed to activate specific adapters, strictly enforces formatting rules in real time, and packages everything as a standard Python function. This insulates the main application from the unpredictable nature of raw AI text.
With Granite Libraries, a Granite base model can call task-specific experts — low-rank adapters (LoRAs), or activated-LoRAs (aLoRAs) — that have been trained to perform a well-defined function through a software interface. This gives a smaller model the ability to perform narrow tasks on par with large generalist models — with much lower inference costs.
When a library adapter is active, the model can become single-mindedly excellent at that task. While the base model stays unchanged, its behavior can now be precisely prescribed as needed, with nearly no cost to switch aLoRAs in and out.
Below, you can see how much faster this process can be with Granite Libraries. This interactive is an animated replay of benchmarking two Granite Switch checkpoints — one using aLoRA, and one using standard LoRA — on the same multi-step RAG pipeline. To speed up or slow down the race, click on the different speeds at the top left of the box:
If you have about 20 minutes to spare, you can even recreate this race in Colab . The two servers run sequentially, as Colab usually only provides one GPU, but the replay generated stitches together their telemetry as if they had raced simultaneously.
The architecture layer: Project Granite Switch
Project Granite Switch is a new experimental toolkit available on GitHub that can be used to compose new models in minutes, similar to the way compilers can produce binaries from source code and software libraries.
Granite Switch allows a base Granite model and its adapter functions to act as a single model that activates them efficiently at inference time. It accomplishes this by adding a new “switching” layer to an existing core Granite model, and then gluing the adapter's weights to the base model, adding formatting tags, and preparing a new chat template. Instead of needing to spin up a brand-new AI model for every different task, Granite Switch dynamically flips the right adapter on and off exactly when it’s needed. The base model is still accessible within Granite Switch, meaning the new capabilities are available without changing the underlying model in any way.
This independent switching layer allows both LoRAs and aLoRAs to run within vLLM , the open-source inferencing engine for large-scale deployments. In the real world, a single business task usually requires a chain of actions, like running safety checks, finding data, and verifying answers. And switching between different adapters forces the AI to clear its short-term memory and recalculate everything from scratch at each step, which slows things down. By using activated LoRA, a Granite Switch model can carry its memory forward from one step to the next without pausing to re-read it, dramatically speeding up multi-step workflows.
By inserting an extra transformer layer into the base model and commandeering its attention mechanism to read and save values related to the active adapter’s state, a special control token signals to the model when it’s time to switch experts. That token acts like a dispatcher at a railyard instructing which trains to go where, with the switching layer acting as the rails themselves.
Granite 4.1: IBM’s most capable models to date
The potential of Granite Libraries and Project Granite Switch would be lost without capable models to execute on top of. With Granite 4.1 , IBM recently released its most performant suite of models yet.
The Granite 4.1 family is designed to punch above its weight class, with the 8B model matching or exceeding the performance of the previous Granite 32B mixture-of-experts (MoE) model and the 30B model competing with significantly larger models like Llama 3.3 70B on enterprise tasks. Small, performant models that can be quickly adapted to other tasks are far less costly to serve than a massive, generalist model that might be incompetent at narrow tasks.
By training on a relatively small amount of high-quality data, these models achieved competitive scores in tool calling and instruction following while maintaining lower latency and operational costs than many frontier reasoning models.
The release is part of a broader ecosystem that includes Granite Vision 4.1 for industry-leading table and chart extraction, as well as new speech and guardrail models. All were trained on approximately 15 trillion tokens and are released under the open-source Apache 2.0 license, supporting 12 major languages for global deployment.
IBM has introduced Granite Libraries as part of a broader goal to make AI models as composable as software to deliver more value to enterprise users. By separating capabilities into modular components, developers can build AI systems that are easier to adapt, cheaper to operate, and more predictable in production.
Modularity won’t solve every challenge of deploying generative AI at scale, but it offers a practical path toward more sustainable, enterprise-ready systems.
Granite Switch : Compose adapter functions into a deployable model
Granite Libraries : Browse and download adapter functions for RAG, safety, and core capabilities
Mellea : Call adapter functions as typed Python functions
Granite 4.1 models : The base models that power it all
Subscribe to our newsletter Home
The future of quantum takes center stage at NY Tech Week
News Peter Hess 04 Jun 2026 AI
Renowned mathematician Subhash Khot joins IBM Research
News Mike Murphy 01 Jun 2026 Mathematical Sciences
IBM charts a new research path with MIT
Q & A Kim Martineau 11 May 2026 Accelerated Discovery
Where the frontiers of high-speed racing and computing meet
Research Kim Martineau and Dave Mosher 30 Apr 2026 Accelerated Discovery
