---
source: "https://github.com/humanlayer/12-factor-agents"
hn_url: "https://news.ycombinator.com/item?id=49323369"
title: "12-Factor Agents – Principles for building reliable LLM applications"
article_title: "GitHub - humanlayer/12-factor-agents: What are the principles we can use to build LLM-powered software that is actually good enough to put in the hands of production customers? · GitHub"
author: "rzk"
captured_at: "2026-08-16T21:11:56Z"
capture_tool: "hn-digest"
hn_id: 49323369
score: 2
comments: 0
posted_at: "2026-08-16T20:30:06Z"
tags:
  - hacker-news
  - translated
---

# 12-Factor Agents – Principles for building reliable LLM applications

- HN: [49323369](https://news.ycombinator.com/item?id=49323369)
- Source: [github.com](https://github.com/humanlayer/12-factor-agents)
- Score: 2
- Comments: 0
- Posted: 2026-08-16T20:30:06Z

## Translation

タイトル: 12-Factor Agent – 信頼性の高い LLM アプリケーションを構築するための原則
記事のタイトル: GitHub - humanlayer/12-factor-agents: 実際に運用顧客に提供するのに十分な LLM ベースのソフトウェアを構築するために使用できる原則は何ですか? · GitHub
説明: 実際に運用顧客に提供するのに十分な LLM ベースのソフトウェアを構築するために使用できる原則は何ですか? - humanlayer/12-factor-agents

記事本文:
GitHub - humanlayer/12-factor-agents: 実際に運用顧客に提供するのに十分な LLM ベースのソフトウェアを構築するために使用できる原則は何ですか? · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
人間層
/
12ファクターエージェント
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
273 コミット 273 コミット コンテンツ コンテンツ ドラフト dra

fts hack/ contributors_markdown hack/ contributors_markdown img img パッケージ パッケージ ワークショップ ワークショップ .gitignore .gitignore CLAUDE.md CLAUDE.md ライセンス ライセンス Makefile Makefile README.md README.md すべてのファイルを表示 リポジトリ ファイル ナビゲーション
12-Factor Agent - 信頼性の高い LLM アプリケーションを構築するための原則
12 Factor Apps の精神に基づいて。このプロジェクトのソースは https://github.com/humanlayer/12-factor-agents で公開されています。フィードバックや貢献を歓迎します。これを一緒に考えてみましょう！
AI エンジニア ワールド フェアに参加しませんでしたか?ここで講演をご覧ください
コンテキストエンジニアリングをお探しですか?因子 3 に直接ジャンプします
npx/uvx create-12-factor-agent に貢献したい - ディスカッション スレッドをチェックしてください
こんにちは、デックスです。私はしばらく AI エージェントをハッキングしてきました。
私は、プラグアンドプレイのクルー/ラングチェーンから、世界中の「ミニマリスト」スモラジェント、「プロダクショングレード」のラングラフ、グリップテープなどに至るまで、あらゆるエージェントフレームワークを試してきました。
私は YC 内外の多くの本当に強力な創業者と話をしましたが、彼らは皆 AI を使って本当に素晴らしいものを構築しています。彼らのほとんどは自分でスタックをローリングしています。実稼働の顧客対応エージェントにはフレームワークがあまり見当たりません。
「AI エージェント」と宣伝している製品のほとんどがそれほどエージェント的ではないことに私は驚きました。それらの多くはほとんどが決定論的なコードであり、適切なポイントに LLM ステップが散りばめられており、エクスペリエンスを真に魔法のものにしています。
エージェントは、少なくとも優秀なエージェントは、「これがプロンプト、これがツールの入ったバッグ、目標に到達するまでループする」というパターンには従いません。むしろ、それらはほとんど単なるソフトウェアで構成されています。
実際に運用顧客に提供するのに十分な LLM ベースのソフトウェアを構築するために使用できる原則は何でしょうか?
12-factor a へようこそ

男性諸君。デーリー以来、シカゴ市長が一貫して市の主要空港のあちこちに張り巡らされてきたように、皆さんがここに来てくれてうれしいです。
このガイドに関する初期のフィードバックについては、 @iantbutler01 、 @tnm 、 @hellovai 、 @stantonk 、 @balanceiskey 、 @AdjectiveAllison 、 @pfbyjy 、 @a-churchill 、および SF MLOps コミュニティに感謝します。
ショートバージョン: 12 の要素
たとえ LLM が指数関数的に強力になり続けたとしても、LLM を利用したソフトウェアの信頼性、拡張性、保守を容易にするコアとなるエンジニアリング技術が存在するでしょう。
私たちがここにたどり着いた経緯: ソフトウェアの簡単な歴史
要素 1: 自然言語からツール呼び出しまで
要素 3: コンテキスト ウィンドウを所有する
要素 4: ツールは単なる構造化された出力です
要素 5: 実行状態とビジネス状態を統合する
要素 6: シンプルな API による起動/一時停止/再開
要素 7: ツール呼び出しで人間と接触する
要素 8: 制御フローを所有する
要素 9: エラーをコンテキスト ウィンドウに圧縮する
要素 10: 小規模で集中的なエージェント
要素 11: どこからでもトリガーし、どこにいてもユーザーに対応
要素 12: エージェントをステートレス リデューサーにする
私のエージェントとしての歩みと、私たちがここにたどり着いた理由について詳しくは、「ソフトウェアの簡単な歴史」をご覧ください。簡単な概要は次のとおりです。
有向グラフ (DG) とその非循環の友人である DAG について詳しく説明します。まず指摘したいのは、...そうですね...ソフトウェアは有向グラフであるということです。私たちがプログラムをフローチャートとして表現していたのには理由があります。
約 20 年前、DAG オーケストレーターが普及し始めました。ここで話しているのは、 Airflow 、 Prefect などの古典的なもの、いくつかの前任者、そして ( dagster 、 inggest 、 Windmill ) などの新しいものです。これらは同じグラフ パターンに従っており、可観測性、モジュール性、再試行、管理などの利点が追加されています。
これを言ったのは私が最初ではありませんが、私が学び始めたときの最大の教訓

エージェントについては、DAG を破棄できるということでした。ソフトウェア エンジニアが各ステップやエッジ ケースをコーディングする代わりに、エージェントに目標と一連の移行を与えることができます。
そして、LLM にリアルタイムで意思決定を行わせ、経路を見つけ出します。
ここでの約束は、作成するソフトウェアの量を減らし、LLM にグラフの「エッジ」を与えて、LLM にノードを認識させるだけです。エラーから回復でき、記述するコードの量を減らすことができ、LLM が問題に対する新しい解決策を見つけられるかもしれません。
後で説明しますが、これは完全に機能しないことがわかります。
さらに一歩深く見てみましょう。エージェントを使用すると、次の 3 つのステップで構成されるループが得られます。
LLM はワークフローの次のステップを決定し、構造化された JSON を出力します (「ツール呼び出し」)。
決定論的なコードがツール呼び出しを実行する
結果はコンテキスト ウィンドウに追加されます
次のステップが「完了」したと判断されるまで繰り返します
初期イベント = { "メッセージ" : "..." }
context = [初期イベント]
True の場合:
next_step = llm を待ちます。決定_次のステップ (コンテキスト)
コンテキスト 。追加 ( next_step )
if ( next_step .tent == = "done" ):
next_step を返します。最終回答
result = await 実行ステップ ( next_step )
コンテキスト 。追加 (結果)
初期コンテキストは単なる開始イベント (ユーザー メッセージ、cron の起動、Webhook など) であり、llm に次のステップ (ツール) を選択するか、完了したかどうかを判断するよう依頼します。
027-エージェント-ループ-アニメーション.mp4
GIFバージョン
結局のところ、このアプローチは私たちが望んでいるほどうまく機能しません。
HumanLayer を構築するにあたり、私は既存の製品をよりエージェント的にしたいと考えている少なくとも 100 人の SaaS ビルダー (ほとんどが技術的な創設者) と話をしました。旅は通常次のようなものになります。
エージェントを構築することを決定する
プロダクトデザイン、UXマッピング、解決すべき問題は何か
すばやく移動したいので、$FRAMEWORK を取得してください

建物に着く
ほとんどの顧客向け機能には 80% では十分ではないことを認識してください
80% を超えるには、フレームワーク、プロンプト、フローなどのリバース エンジニアリングが必要であることに注意してください。
免責事項: これをどこで言うのが正確かはわかりませんが、ここは他のどのフレームワークよりも優れていると思われます。これは、決して意味するものではなく、世の中にある多くのフレームワーク、またはそれらに取り組んでいる非常に賢い人々のいずれかを掘り下げることを意味しました。これらは信じられないほどのことを可能にし、AI エコシステムを加速させてきました。
この投稿の成果の 1 つは、エージェント フレームワークの構築者が私や他の人の取り組みから学び、フレームワークをさらに改善できることを願っています。
特に、速く動きたいが深いコントロールが必要なビルダーに最適です。
免責事項 2 : MCP について話すつもりはありません。どこに当てはまるかがわかると思います。
免責事項 3 : 理由により、私は主に typescript を使用していますが、これらはすべて Python またはその他の好みの言語で動作します。
優れた LLM アプリケーションのための設計パターン
何百もの AI ライブラリを調べ、何十人もの創設者と協力した後、私の直感は次のとおりです。
エージェントを優れたものにする核となるものがいくつかあります
フレームワークに全力を尽くして、本質的にグリーンフィールドの書き換えに相当するものを構築することは逆効果になる可能性があります
エージェントを優れたものにするための核となる原則がいくつかあり、フレームワークを導入すれば、それらのほとんど/すべてを実現できます。
しかし、私が見た中でビルダーが高品質の AI ソフトウェアを顧客の手に渡す最も早い方法は、エージェント構築から小さなモジュール式のコンセプトを取り出し、それらを既存の製品に組み込むことです。
エージェントからのこれらのモジュール式の概念は、AI の背景がなくても、最も熟練したソフトウェア エンジニアによって定義および適用できます。
ビルダーが優れた AI ソフトウェアを顧客の手に渡すための最も早い方法は、小さなモジュール型のコンセプトを採用することです。

romエージェントを構築し、既存の製品に組み込む
私たちがここにたどり着いた経緯: ソフトウェアの簡単な歴史
要素 1: 自然言語からツール呼び出しまで
要素 3: コンテキスト ウィンドウを所有する
要素 4: ツールは単なる構造化された出力です
要素 5: 実行状態とビジネス状態を統合する
要素 6: シンプルな API による起動/一時停止/再開
要素 7: ツール呼び出しで人間と接触する
要素 8: 制御フローを所有する
要素 9: エラーをコンテキスト ウィンドウに圧縮する
要素 10: 小規模で集中的なエージェント
要素 11: どこからでもトリガーし、どこにいてもユーザーに対応
要素 12: エージェントをステートレス リデューサーにする
佳作・その他アドバイス
要素 13: 必要なすべてのコンテキストを事前に取得する
2025 年 3 月の Tool Use ポッドキャストのエピソードで、これについて多くのことを話しました。
このことについてはアウター ループで書いています
@hellovai と LLM パフォーマンスの最大化に関するウェビナーを行っています
この方法論を使用して OSS エージェントを got-agents/agents の下に構築します。
私たちは独自のアドバイスをすべて無視し、kubernetes で分散エージェントを実行するためのフレームワークを構築しました。
このガイドの他のリンク:
12 ファクター アプリ
効果的なエージェントの構築 (人為的)
ライブラリ パターン: フレームワークが悪である理由
関数呼び出し vs 構造化出力 vs JSON モード
OpenAI JSON と関数呼び出し
NotebookLM によるモデル機能の境界の検索
12-Factor エージェントに貢献してくださった皆様に感謝します。
すべてのコンテンツと画像は CC BY-SA 4.0 ライセンスに基づいてライセンスされています。
コードは Apache 2.0 ライセンスに基づいてライセンスされています
実際に運用顧客に提供するのに十分な LLM ベースのソフトウェアを構築するために使用できる原則は何でしょうか?
Readme ライセンス アクティビティ カスタム プロパティ スター
1.9k フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

What are the principles we can use to build LLM-powered software that is actually good enough to put in the hands of production customers? - humanlayer/12-factor-agents

GitHub - humanlayer/12-factor-agents: What are the principles we can use to build LLM-powered software that is actually good enough to put in the hands of production customers? · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
humanlayer
/
12-factor-agents
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
273 Commits 273 Commits content content drafts drafts hack/ contributors_markdown hack/ contributors_markdown img img packages packages workshops workshops .gitignore .gitignore CLAUDE.md CLAUDE.md LICENSE LICENSE Makefile Makefile README.md README.md View all files Repository files navigation
12-Factor Agents - Principles for building reliable LLM applications
In the spirit of 12 Factor Apps . The source for this project is public at https://github.com/humanlayer/12-factor-agents , and I welcome your feedback and contributions. Let's figure this out together!
Missed the AI Engineer World's Fair? Catch the talk here
Looking for Context Engineering? Jump straight to factor 3
Want to contribute to npx/uvx create-12-factor-agent - check out the discussion thread
Hi, I'm Dex. I've been hacking on AI agents for a while .
I've tried every agent framework out there , from the plug-and-play crew/langchains to the "minimalist" smolagents of the world to the "production grade" langraph, griptape, etc.
I've talked to a lot of really strong founders , in and out of YC, who are all building really impressive things with AI. Most of them are rolling the stack themselves. I don't see a lot of frameworks in production customer-facing agents.
I've been surprised to find that most of the products out there billing themselves as "AI Agents" are not all that agentic. A lot of them are mostly deterministic code, with LLM steps sprinkled in at just the right points to make the experience truly magical.
Agents, at least the good ones, don't follow the "here's your prompt, here's a bag of tools, loop until you hit the goal" pattern. Rather, they are comprised of mostly just software.
What are the principles we can use to build LLM-powered software that is actually good enough to put in the hands of production customers?
Welcome to 12-factor agents. As every Chicago mayor since Daley has consistently plastered all over the city's major airports, we're glad you're here.
Special thanks to @iantbutler01 , @tnm , @hellovai , @stantonk , @balanceiskey , @AdjectiveAllison , @pfbyjy , @a-churchill , and the SF MLOps community for early feedback on this guide.
The Short Version: The 12 Factors
Even if LLMs continue to get exponentially more powerful , there will be core engineering techniques that make LLM-powered software more reliable, more scalable, and easier to maintain.
How We Got Here: A Brief History of Software
Factor 1: Natural Language to Tool Calls
Factor 3: Own your context window
Factor 4: Tools are just structured outputs
Factor 5: Unify execution state and business state
Factor 6: Launch/Pause/Resume with simple APIs
Factor 7: Contact humans with tool calls
Factor 8: Own your control flow
Factor 9: Compact Errors into Context Window
Factor 10: Small, Focused Agents
Factor 11: Trigger from anywhere, meet users where they are
Factor 12: Make your agent a stateless reducer
For a deeper dive on my agent journey and what led us here, check out A Brief History of Software - a quick summary here:
We're gonna talk a lot about Directed Graphs (DGs) and their Acyclic friends, DAGs. I'll start by pointing out that...well...software is a directed graph. There's a reason we used to represent programs as flow charts.
Around 20 years ago, we started to see DAG orchestrators become popular. We're talking classics like Airflow , Prefect , some predecessors, and some newer ones like ( dagster , inggest , windmill ). These followed the same graph pattern, with the added benefit of observability, modularity, retries, administration, etc.
I'm not the first person to say this , but my biggest takeaway when I started learning about agents, was that you get to throw the DAG away. Instead of software engineers coding each step and edge case, you can give the agent a goal and a set of transitions:
And let the LLM make decisions in real time to figure out the path
The promise here is that you write less software, you just give the LLM the "edges" of the graph and let it figure out the nodes. You can recover from errors, you can write less code, and you may find that LLMs find novel solutions to problems.
As we'll see later, it turns out this doesn't quite work.
Let's dive one step deeper - with agents you've got this loop consisting of 3 steps:
LLM determines the next step in the workflow, outputting structured json ("tool calling")
Deterministic code executes the tool call
The result is appended to the context window
Repeat until the next step is determined to be "done"
initial_event = { "message" : "..." }
context = [ initial_event ]
while True :
next_step = await llm . determine_next_step ( context )
context . append ( next_step )
if ( next_step . intent == = "done" ):
return next_step . final_answer
result = await execute_step ( next_step )
context . append ( result )
Our initial context is just the starting event (maybe a user message, maybe a cron fired, maybe a webhook, etc), and we ask the llm to choose the next step (tool) or to determine that we're done.
027-agent-loop-animation.mp4
GIF Version
At the end of the day, this approach just doesn't work as well as we want it to.
In building HumanLayer, I've talked to at least 100 SaaS builders (mostly technical founders) looking to make their existing product more agentic. The journey usually goes something like:
Decide you want to build an agent
Product design, UX mapping, what problems to solve
Want to move fast, so grab $FRAMEWORK and get to building
Realize that 80% isn't good enough for most customer-facing features
Realize that getting past 80% requires reverse-engineering the framework, prompts, flow, etc.
DISCLAIMER : I'm not sure the exact right place to say this, but here seems as good as any: this in BY NO MEANS meant to be a dig on either the many frameworks out there, or the pretty dang smart people who work on them . They enable incredible things and have accelerated the AI ecosystem.
I hope that one outcome of this post is that agent framework builders can learn from the journeys of myself and others, and make frameworks even better.
Especially for builders who want to move fast but need deep control.
DISCLAIMER 2 : I'm not going to talk about MCP. I'm sure you can see where it fits in.
DISCLAIMER 3 : I'm using mostly typescript, for reasons but all this stuff works in python or any other language you prefer.
Design Patterns for great LLM applications
After digging through hundreds of AI libriaries and working with dozens of founders, my instinct is this:
There are some core things that make agents great
Going all in on a framework and building what is essentially a greenfield rewrite may be counter-productive
There are some core principles that make agents great, and you will get most/all of them if you pull in a framework
BUT, the fastest way I've seen for builders to get high-quality AI software in the hands of customers is to take small, modular concepts from agent building, and incorporate them into their existing product
These modular concepts from agents can be defined and applied by most skilled software engineers, even if they don't have an AI background
The fastest way I've seen for builders to get good AI software in the hands of customers is to take small, modular concepts from agent building, and incorporate them into their existing product
How We Got Here: A Brief History of Software
Factor 1: Natural Language to Tool Calls
Factor 3: Own your context window
Factor 4: Tools are just structured outputs
Factor 5: Unify execution state and business state
Factor 6: Launch/Pause/Resume with simple APIs
Factor 7: Contact humans with tool calls
Factor 8: Own your control flow
Factor 9: Compact Errors into Context Window
Factor 10: Small, Focused Agents
Factor 11: Trigger from anywhere, meet users where they are
Factor 12: Make your agent a stateless reducer
Honorable Mentions / other advice
Factor 13: Pre-fetch all the context you might need
I talked about a lot of this on an episode of the Tool Use podcast in March 2025
I write about some of this stuff at The Outer Loop
I do webinars about Maximizing LLM Performance with @hellovai
We build OSS agents with this methodology under got-agents/agents
We ignored all our own advice and built a framework for running distributed agents in kubernetes
Other links from this guide:
12 Factor Apps
Building Effective Agents (Anthropic)
Library patterns: Why frameworks are evil
Function Calling vs Structured Outputs vs JSON Mode
OpenAI JSON vs Function Calling
NotebookLM on Finding Model Capability Boundaries
Thanks to everyone who has contributed to 12-factor agents!
All content and images are licensed under a CC BY-SA 4.0 License
Code is licensed under the Apache 2.0 License
What are the principles we can use to build LLM-powered software that is actually good enough to put in the hands of production customers?
Readme License Activity Custom properties Stars
1.9k forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
