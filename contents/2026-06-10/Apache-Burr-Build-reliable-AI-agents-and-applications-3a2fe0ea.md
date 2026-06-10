---
source: "https://burr.apache.org/"
hn_url: "https://news.ycombinator.com/item?id=48477400"
title: "Apache Burr: Build reliable AI agents and applications"
article_title: "Apache Burr (Incubating) - Build Reliable AI Agents and Applications"
author: "anhldbk"
captured_at: "2026-06-10T16:21:05Z"
capture_tool: "hn-digest"
hn_id: 48477400
score: 47
comments: 17
posted_at: "2026-06-10T15:01:06Z"
tags:
  - hacker-news
  - translated
---

# Apache Burr: Build reliable AI agents and applications

- HN: [48477400](https://news.ycombinator.com/item?id=48477400)
- Source: [burr.apache.org](https://burr.apache.org/)
- Score: 47
- Comments: 17
- Posted: 2026-06-10T15:01:06Z

## Translation

タイトル: Apache Burr: 信頼性の高い AI エージェントとアプリケーションを構築する
記事のタイトル: Apache Burr (インキュベーション) - 信頼性の高い AI エージェントとアプリケーションの構築
説明: Apache Burr (Incubating) を使用すると、単純なチャットボットから複雑なマルチエージェント システムまで、意思決定を行うアプリケーションを簡単に開発できます。純粋な Python であり、魔法は必要ありません。

記事本文:
Apache Burr (インキュベート) - 信頼できる AI エージェントとアプリケーションの構築 Apache Burr インキュベート機能 統合 コミュニティ ダウンロード ドキュメント Discord Apache Incubating Project 信頼できる AI エージェントとアプリケーションの構築
Apache Burr (Incubating) を使用すると、単純なチャットボットから複雑なマルチエージェント システムに至るまで、意思決定を行うアプリケーションを簡単に開発できます。純粋な Python であり、魔法は必要ありません。
GitHub で表示 0 GitHub スター 0 k+ PyPI ダウンロード 0 + Discord メンバー シンプルで強力な Python API
クリーンで構成可能なインターフェイスを使用して、チャットボットからマルチエージェント システムまであらゆるものを構築します。
chatbot.py burr.core インポート アクション、State、ApplicationBuilder からチャットボット マルチエージェント ステート マシンをコピーします。
@action(読み取り=["メッセージ"]、書き込み=["メッセージ"])
def chat(状態: 状態, llm_client) -> 状態:
応答 = llm_client.chat(state["messages"])
状態を返す.update(
メッセージ=[*状態["メッセージ"]、応答]
)
アプリ = (
ApplicationBuilder()
.with_actions(チャット)
.with_transitions(("チャット", "チャット"))
.with_state(メッセージ=[])
.with_tracker("ローカル")
.build()
)
app.run(halt_after=["chat"], inputs={"llm_client": client}) Python 3.12 UTF-8 burr >= 0.30 AI アプリケーションの構築に必要なものすべて
Burr は、信頼性が高く、観察可能で、テスト可能な AI を活用したアプリケーションの構成要素を提供します。
アプリケーションをアクションと遷移のセットとして定義します。 DSL や YAML はなく、Python 関数とデコレータだけが使用されます。
Burr UI を使用すると、アプリケーションのすべてのステップをリアルタイムで監視、デバッグ、トレースできます。発生した状態の変化を確認します。
永続性と状態管理
状態をディスク、データベース、またはカスタム バックエンドに自動的に永続化します。中断したところからアプリケーションを再開します。
いずれかのステップで実行を一時停止し、人間の入力を待ちます。承認ワークフローや対話型エージェントに最適です。
アクションを並行して実行、ファンアウト / ファン i

n を使用して、複雑な DAG を構築します。モジュール設計用のサブアプリケーションを作成します。
過去の実行を再生し、個々のアクションを単体テストし、状態遷移を検証します。 AI システムに対する信頼を築きます。
Burr は、すでに使用しているツールやフレームワークと統合します。ロックインやラッパーはありません。
世界中のエンジニアから信頼されています
開発者やチームが Burr について述べていることをご覧ください。
「他のいくつかの難読化 LLM フレームワークを評価した結果、そのエレガントでありながら包括的な状態管理ソリューションが、AI の意思決定によって駆動されるロボットの展開に対する強力な答えであることが証明されました。」
「モジュール式 AI アプリケーションを構築したい場合、Burr を使用するのは簡単です。構築は非常に簡単で、デバッグが簡単になる UI が特に気に入っています。そして、チームを常に支援する用意ができているのが最高です。」
「今バーに出会って、すごいと思いました。皆さんがこれを構築するときにまさにこのニーズを予測していたようです。AI だからといって、奇妙な難解な概念は必要ありません。」
「 Burr の状態管理部分は、状態スナップショットの作成やビルドのデバッグ、リプレイ、さらにはそれを中心とした評価ケースの構築に非常に役立ちます。」
「私は過去数か月間 Burr を使用してきましたが、世の中の多くのエージェント LLM プラットフォーム (LangChain、CrewAi、AutoGen、Agency Swarm など) と比較して、Burr は複雑な動作を設計するためのより堅牢なフレームワークを提供します。」
「 LangChain から Burr への移行は状況を大きく変えるものでした。LangChain を操作するのに何日も何週間も費やしたのに比べ、私は Burr を使い始めるのにわずか数時間しかかかりませんでした。私はチームメイトに Burr を提案し、コードベース全体をそれにピボットしました。」
「 もちろん、[LangChain] を使用することはできますが、それが本当に本番環境に対応し、コードから本番環境までの時間が短縮されるかどうかはわかりませんが、私たちは 2 年間 LLM アプリを開発してきました。

答えはノーです。正直に言って、バーを見てください。後でよろしくお願いします。 」
「 他のいくつかの難読化 LLM フレームワークを評価した結果、そのエレガントでありながら包括的な状態管理ソリューションが、AI の意思決定によって駆動されるロボットの展開に対する強力な答えであることが判明しました。 」
「 モジュール式 AI アプリケーションを構築したい場合、Burr を使用するのは簡単です。構築はとても簡単で、デバッグが簡単になる UI が特に気に入っています。そして、いつでも助けてくれるチームが最高のチームです。 」
「ちょうど Burr に出会って、すごいと思いました。皆さんがこれを構築するときにまさにこのニーズを予測していたようです。 AIだからといって、変な難解な概念はありません。 」
「 Burr の状態管理部分は、状態スナップショットの作成やビルドのデバッグ、再生、さらにはそれを中心とした評価ケースの構築に非常に役立ちます。 」
「 私は過去数か月間 Burr を使用してきましたが、世の中の多くのエージェント LLM プラットフォーム (LangChain、CrewAi、AutoGen、Agency Swarm など) と比較して、Burr は複雑な動作を設計するためのより堅牢なフレームワークを提供します。 」
「 LangChain から Burr への移行はゲームチェンジャーでした。 LangChain を操作するのに何日も何週間も費やしたのに比べ、私は Burr を使い始めるのにわずか数時間しかかかりませんでした。私は Burr をチームメイトに提案し、コードベース全体をそれにピボットしました。 」
「 もちろん、[LangChain] を使用することはできますが、それが本当に本番環境に対応し、コードから本番環境までの時間が短縮されるかどうかは、私たちは 2 年間 LLM アプリを開発してきましたが、答えはノーです。正直に言って、バーを見てください。後でよろしくお願いします。 」
「 他のいくつかの難読化 LLM フレームワークを評価した結果、そのエレガントでありながら包括的な状態管理ソリューションが、AI の意思決定によって駆動されるロボットの展開に対する強力な答えであることが判明しました。 」
「バールを使用する」

モジュール式 AI アプリケーションを構築したい場合は、r を使用するのは簡単です。構築はとても簡単で、デバッグが簡単になる UI が特に気に入っています。そして、いつでも助けてくれるチームが最高のチームです。 」
「ちょうど Burr に出会って、すごいと思いました。皆さんがこれを構築するときにまさにこのニーズを予測していたようです。 AIだからといって、変な難解な概念はありません。 」
「 Burr の状態管理部分は、状態スナップショットの作成やビルドのデバッグ、再生、さらにはそれを中心とした評価ケースの構築に非常に役立ちます。 」
「 私は過去数か月間 Burr を使用してきましたが、世の中の多くのエージェント LLM プラットフォーム (LangChain、CrewAi、AutoGen、Agency Swarm など) と比較して、Burr は複雑な動作を設計するためのより堅牢なフレームワークを提供します。 」
「 LangChain から Burr への移行はゲームチェンジャーでした。 LangChain を操作するのに何日も何週間も費やしたのに比べ、私は Burr を使い始めるのにわずか数時間しかかかりませんでした。私は Burr をチームメイトに提案し、コードベース全体をそれにピボットしました。 」
「 もちろん、[LangChain] を使用することはできますが、それが本当に本番環境に対応し、コードから本番環境までの時間が短縮されるかどうかは、私たちは 2 年間 LLM アプリを開発してきましたが、答えはノーです。正直に言って、バーを見てください。後でよろしくお願いします。 」
「 他のいくつかの難読化 LLM フレームワークを評価した結果、そのエレガントでありながら包括的な状態管理ソリューションが、AI の意思決定によって駆動されるロボットの展開に対する強力な答えであることが判明しました。 」
「 モジュール式 AI アプリケーションを構築したい場合、Burr を使用するのは簡単です。構築はとても簡単で、デバッグが簡単になる UI が特に気に入っています。そして、いつでも助けてくれるチームが最高のチームです。 」
「ちょうどバーに出会ったんだけど、すごいなって思ったんだ、君みたいだね」

これを構築するときに、まさにこのニーズを予測しました。 AIだからといって、変な難解な概念はありません。 」
「 Burr の状態管理部分は、状態スナップショットの作成やビルドのデバッグ、再生、さらにはそれを中心とした評価ケースの構築に非常に役立ちます。 」
「 私は過去数か月間 Burr を使用してきましたが、世の中の多くのエージェント LLM プラットフォーム (LangChain、CrewAi、AutoGen、Agency Swarm など) と比較して、Burr は複雑な動作を設計するためのより堅牢なフレームワークを提供します。 」
「 LangChain から Burr への移行はゲームチェンジャーでした。 LangChain を操作するのに何日も何週間も費やしたのに比べ、私は Burr を使い始めるのにわずか数時間しかかかりませんでした。私は Burr をチームメイトに提案し、コードベース全体をそれにピボットしました。 」
「 もちろん、[LangChain] を使用することはできますが、それが本当に本番環境に対応し、コードから本番環境までの時間が短縮されるかどうかは、私たちは 2 年間 LLM アプリを開発してきましたが、答えはノーです。正直に言って、バーを見てください。後でよろしくお願いします。 」
助けを求め、プロジェクトを共有し、バーの将来に貢献してください。
メンテナーやコミュニティとチャットする
リポジトリにスターを付け、問題をファイルし、貢献します
フォローして最新情報やお知らせを入手してください
Apache Burr (Incubating) は、Apache Incubator のスポンサーにより、Apache Software Foundation (ASF) でインキュベーション中の取り組みです。インフラストラクチャ、コミュニケーション、および意思決定プロセスが他の成功した ASF プロジェクトと一致する形で安定していることがさらなるレビューで示されるまで、新しく受け入れられたすべてのプロジェクトはインキュベーションが必要です。インキュベーション ステータスはコードの完全性や安定性を必ずしも反映しているわけではありませんが、プロジェクトがまだ ASF によって完全に承認されていないことを示しています。
Apache Burr、Burr、Apache、Apache フェザー ロゴ、および Apache Burr プロジェクト

ct ロゴは、米国およびその他の国における Apache Software Foundation の登録商標または商標です。言及されているその他すべてのマークは、それぞれの所有者の商標または登録商標である可能性があります。

## Original Extract

Apache Burr (Incubating) makes it easy to develop applications that make decisions, from simple chatbots to complex multi-agent systems. Pure Python, no magic.

Apache Burr (Incubating) - Build Reliable AI Agents and Applications Apache Burr Incubating Features Integrations Community Download Docs Discord Apache Incubating Project Build reliable AI agents and applications
Apache Burr (Incubating) makes it easy to develop applications that make decisions, from simple chatbots to complex multi-agent systems. Pure Python, no magic.
View on GitHub 0 GitHub Stars 0 k+ PyPI Downloads 0 + Discord Members Simple, powerful Python API
Build anything from chatbots to multi-agent systems with a clean, composable interface.
chatbot.py Copy Chatbot Multi-Agent State Machine from burr.core import action, State, ApplicationBuilder
@action(reads=["messages"], writes=["messages"])
def chat(state: State, llm_client) -> State:
response = llm_client.chat(state["messages"])
return state.update(
messages=[*state["messages"], response]
)
app = (
ApplicationBuilder()
.with_actions(chat)
.with_transitions(("chat", "chat"))
.with_state(messages=[])
.with_tracker("local")
.build()
)
app.run(halt_after=["chat"], inputs={"llm_client": client}) Python 3.12 UTF-8 burr >= 0.30 Everything you need to build AI applications
Burr provides the building blocks for reliable, observable, and testable AI-powered applications.
Define your application as a set of actions and transitions. No DSL, no YAML — just Python functions and decorators.
The Burr UI lets you monitor, debug, and trace every step of your application in real time. See state changes as they happen.
Persistence & State Management
Automatically persist state to disk, databases, or custom backends. Resume applications from where they left off.
Pause execution and wait for human input at any step. Perfect for approval workflows and interactive agents.
Run actions in parallel, fan out / fan in, and build complex DAGs. Compose sub-applications for modular design.
Replay past runs, unit test individual actions, and validate state transitions. Build confidence in your AI systems.
Burr integrates with the tools and frameworks you already use. No lock-in, no wrappers.
Trusted by engineers worldwide
See what developers and teams are saying about Burr.
“ After evaluating several other obfuscating LLM frameworks, their elegant yet comprehensive state management solution proved to be the powerful answer to rolling out robots driven by AI decision making. ”
“ Using Burr is a no-brainer if you want to build a modular AI application. It is so easy to build with and I especially love their UI which makes debugging a piece of cake. And the always ready to help team is the cherry on top. ”
“ I just came across Burr and I'm like WOW, this seems like you guys predicted this exact need when building this. No weird esoteric concepts just because it's AI. ”
“ Burr's state management part is really helpful for creating state snapshots and build debugging, replaying and even building evaluation cases around that. ”
“ I have been using Burr over the past few months, and compared to many agentic LLM platforms out there (e.g. LangChain, CrewAi, AutoGen, Agency Swarm, etc), Burr provides a more robust framework for designing complex behaviors. ”
“ Moving from LangChain to Burr was a game-changer! It took me just a few hours to get started with Burr, compared to the days and weeks I spent trying to navigate LangChain. I pitched Burr to my teammates, and we pivoted our entire codebase to it. ”
“ Of course, you can use it [LangChain], but whether it's really production-ready and improves the time from code-to-prod, we've been doing LLM apps for two years, and the answer is no. Honestly, take a look at Burr. Thank me later. ”
“ After evaluating several other obfuscating LLM frameworks, their elegant yet comprehensive state management solution proved to be the powerful answer to rolling out robots driven by AI decision making. ”
“ Using Burr is a no-brainer if you want to build a modular AI application. It is so easy to build with and I especially love their UI which makes debugging a piece of cake. And the always ready to help team is the cherry on top. ”
“ I just came across Burr and I'm like WOW, this seems like you guys predicted this exact need when building this. No weird esoteric concepts just because it's AI. ”
“ Burr's state management part is really helpful for creating state snapshots and build debugging, replaying and even building evaluation cases around that. ”
“ I have been using Burr over the past few months, and compared to many agentic LLM platforms out there (e.g. LangChain, CrewAi, AutoGen, Agency Swarm, etc), Burr provides a more robust framework for designing complex behaviors. ”
“ Moving from LangChain to Burr was a game-changer! It took me just a few hours to get started with Burr, compared to the days and weeks I spent trying to navigate LangChain. I pitched Burr to my teammates, and we pivoted our entire codebase to it. ”
“ Of course, you can use it [LangChain], but whether it's really production-ready and improves the time from code-to-prod, we've been doing LLM apps for two years, and the answer is no. Honestly, take a look at Burr. Thank me later. ”
“ After evaluating several other obfuscating LLM frameworks, their elegant yet comprehensive state management solution proved to be the powerful answer to rolling out robots driven by AI decision making. ”
“ Using Burr is a no-brainer if you want to build a modular AI application. It is so easy to build with and I especially love their UI which makes debugging a piece of cake. And the always ready to help team is the cherry on top. ”
“ I just came across Burr and I'm like WOW, this seems like you guys predicted this exact need when building this. No weird esoteric concepts just because it's AI. ”
“ Burr's state management part is really helpful for creating state snapshots and build debugging, replaying and even building evaluation cases around that. ”
“ I have been using Burr over the past few months, and compared to many agentic LLM platforms out there (e.g. LangChain, CrewAi, AutoGen, Agency Swarm, etc), Burr provides a more robust framework for designing complex behaviors. ”
“ Moving from LangChain to Burr was a game-changer! It took me just a few hours to get started with Burr, compared to the days and weeks I spent trying to navigate LangChain. I pitched Burr to my teammates, and we pivoted our entire codebase to it. ”
“ Of course, you can use it [LangChain], but whether it's really production-ready and improves the time from code-to-prod, we've been doing LLM apps for two years, and the answer is no. Honestly, take a look at Burr. Thank me later. ”
“ After evaluating several other obfuscating LLM frameworks, their elegant yet comprehensive state management solution proved to be the powerful answer to rolling out robots driven by AI decision making. ”
“ Using Burr is a no-brainer if you want to build a modular AI application. It is so easy to build with and I especially love their UI which makes debugging a piece of cake. And the always ready to help team is the cherry on top. ”
“ I just came across Burr and I'm like WOW, this seems like you guys predicted this exact need when building this. No weird esoteric concepts just because it's AI. ”
“ Burr's state management part is really helpful for creating state snapshots and build debugging, replaying and even building evaluation cases around that. ”
“ I have been using Burr over the past few months, and compared to many agentic LLM platforms out there (e.g. LangChain, CrewAi, AutoGen, Agency Swarm, etc), Burr provides a more robust framework for designing complex behaviors. ”
“ Moving from LangChain to Burr was a game-changer! It took me just a few hours to get started with Burr, compared to the days and weeks I spent trying to navigate LangChain. I pitched Burr to my teammates, and we pivoted our entire codebase to it. ”
“ Of course, you can use it [LangChain], but whether it's really production-ready and improves the time from code-to-prod, we've been doing LLM apps for two years, and the answer is no. Honestly, take a look at Burr. Thank me later. ”
Get help, share your projects, and contribute to the future of Burr.
Chat with maintainers and the community
Star the repo, file issues, contribute
Follow for updates and announcements
Apache Burr (Incubating) is an effort undergoing incubation at The Apache Software Foundation (ASF), sponsored by the Apache Incubator . Incubation is required of all newly accepted projects until a further review indicates that the infrastructure, communications, and decision making process have stabilized in a manner consistent with other successful ASF projects. While incubation status is not necessarily a reflection of the completeness or stability of the code, it does indicate that the project has yet to be fully endorsed by the ASF.
Apache Burr, Burr, Apache, the Apache feather logo, and the Apache Burr project logo are either registered trademarks or trademarks of The Apache Software Foundation in the United States and other countries. All other marks mentioned may be trademarks or registered trademarks of their respective owners.
