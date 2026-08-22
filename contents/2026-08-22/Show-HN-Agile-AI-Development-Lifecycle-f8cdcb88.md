---
source: "https://valeriavg.dev/agile-ai-development-lifecycle"
hn_url: "https://news.ycombinator.com/item?id=49397671"
title: "Show HN: Agile AI Development Lifecycle"
article_title: "Agile AI Development Lifecycle - ValeriaVG"
image: "https://valeriavg.dev/og/agile-ai-development-lifecycle.png"
author: "valeriavg_dev"
captured_at: "2026-08-22T08:17:26Z"
capture_tool: "hn-digest"
hn_id: 49397671
score: 1
comments: 0
posted_at: "2026-08-22T08:08:59Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Agile AI Development Lifecycle

- HN: [49397671](https://news.ycombinator.com/item?id=49397671)
- Source: [valeriavg.dev](https://valeriavg.dev/agile-ai-development-lifecycle)
- Score: 1
- Comments: 0
- Posted: 2026-08-22T08:08:59Z

## Translation

タイトル: Show HN: アジャイル AI 開発ライフサイクル
記事のタイトル: アジャイル AI 開発ライフサイクル - ValeriaVG
説明: AI 以前のソフトウェア開発ライフサイクルを大規模な AI ネイティブ開発に変換するためのフレームワークと方法論。
HN テキスト: Schrödinger の AI-DLC フレームワーク

記事本文:
アジャイル AI 開発ライフサイクル - ValeriaVG ValeriaVG ブログ
アジャイル AI 開発ライフサイクル
2025 年末に、AWS は AI-DLC 手法をオープンソース化しました。一見すると、AI 以前のソフトウェア開発ライフサイクルから AI ネイティブの未来へと導く、採用しやすい (マークダウン ファイルをリポジトリにドロップするだけの) フレームワークのように見えます。その本質は、複合エンジニアリングや超大国と非常によく似ています。それらはすべて、人間によるレビューに依存するマークダウン散文のスキルの集合です。特に、どちらも AWS が AI-DLC をリリースする数か月前に一般に導入されましたが、可視性の点で AWS に匹敵するのは困難です。
AWS AI-DLC は、初期化、構想、開始、構築、運用の 5 つのフェーズに分類された、合計 30 以上の開発段階をユーザーがガイドできるように設計されています。すべての段階で、AI が生成したアーティファクトを徹底的に集団でレビューする必要があります。
このアプローチの長所は、導入が容易であることと、AWS によって設計されているという事実です。短所は、本質的に自動化されたウォーターフォール (5 つのフェーズのうち 4 つでドキュメントが生成される) であり、品質はレビュー担当者がどれだけ徹底的で知識があるかによって決まることです。
したがって、AI によって生成された大量のドキュメントを読むのが好きな超人が会社に配置されている場合を除き、特に少数のリポジトリを操作する 1 人のエンジニアのマシンの外で AI 手法を拡張する必要がある場合は、別の方法を探すかもし​​れません。
大規模な AI 導入者の第 2 陣営には、何か新しいものを発明したわけではないため、その方法論に単一の共通の名前がありません。彼らは既存のワークフローを進化させ、AI 機能と統合しました。
著名な出版物には、CNCF プラットフォーム ホワイト ペーパー、Spotify の Golden Paths、Netflix の「舗装道路」などがあります。
核となるアイデアは次のとおりです。

プロセスとツールを標準化して、チームが車輪を再発明することなくビジネス イニシアチブを実行しやすくします。これらの「舗装された道」の維持は、ストリームに合わせたチームと並んでチーム トポロジーによって普及したプラットフォーム エンジニアリング チームにかかっています。
すべてのプロセスが標準化されている場合、成熟したエンジニアリング組織から AI 対応のエンジニアリング組織に移行するには、既存のプロセスとプラットフォーム機能の上にエージェント ワークフローを作成する必要があります。
私は、最終的にこの方法論に適切な名前を付けることを提案します (そして、クラウドの巨人がそれを採用することを願っています)。
AI-DLC とアジャイル AI-DLC の主な違いは、開発段階では文書化ではなくプロトタイプとその進化が生成され、人間によるレビューではなく決定論的なチェックによって品質が保証されることです。
アジャイル AI-DLC は、「舗装されたパス」という概念に基づいて構築されています。これは、エンジニアリング チームだけでなく組織全体にわたる、さまざまな開始点から本番展開までの開発を導く、相互接続された標準化されたプロセスです。
舗装された道の例としては、次のようなものがあります。
デザイナーが Figma でプロトタイプを作成し、インタラクティブなプロトタイプのビルドをトリガーします
プロダクト マネージャーが Slack ボットにタグを付けて、デザイナーのプロトタイプをフルスタック機能で進化させます
エンジニアは、既存のシステム コンポーネントを利用して、アイデアから製品化まで社内プラットフォームを実現します。
決定論的チェックの例は次のとおりです。
独自のデザインシステムと抽象化
厳密なリンターコードチェックと型システム
スコアカードとそれを適用する CI/CD パイプライン
アジャイル AI-DLC は、その名前が示すように、アジャイル開発の基本ルールに従い、次のサイクルとして説明できます。
概念実証: チームが協力して迅速かつ探索的なプロトタイプを開発します。

問題空間を徹底的に調査し、解決策の実現可能性を検証します。 PoC の開発に使用される舗装された道は、問題の空間とチームの構成によって定義されます。この段階では、ドメインの専門家の関与が必要であり、期待される結果は、プロジェクトを続行するかどうかの決定です。
プロトタイプ : 検証された PoC がプロトタイプに進化します。舗装された道はチームの構成 (例: デザイン優先、API 優先など) によって異なります。ドメインの専門家は多少なりとも関与しており、この段階で期待される成果は、MVP バージョンの範囲と要件を定義することです。
実行可能な仕様 : 前の段階で収集された要件は、可能な限り実行可能な仕様 (単体、統合、エンドツーエンドまたはビジュアル テスト、契約、または厳密な機械強制ルール) に変換されます。自動的に検証できない要件は、チームメンバー、関係者、エージェントがアクセスできる方法で文書化および保存されます。ドメインの専門家がオンデマンドで関与し、この段階の結果として、一連のテスト自動化と実行可能な仕様が作成されます。
実装ループ : エージェントは、すべての要件が満たされ、すべてのチェックに合格するまで、一度に 1 つの要件をプロトタイプから実用的な実装に進化させるタスクを負います。
評価 : 実用的な実装は、必要に応じて手動チェックを含む、より広範な要件とチェックに照らして検証され、利害関係者にソリューションを実証します。ドメインの専門家がオンデマンドで関与し、この段階の結果として、次の反復で対処する必要がある一連の追加要件が得られます。
振り返り : チームは実装を振り返り、次のイテレーションに向けてパス、プロセス、要件を調整します。ドマイ

n 人の専門家がオンデマンドで関与し、この段階の結果は次のサイクルを開始する一連のコミットメントとなります。
ドキュメントではなく実行可能な仕様としての要件の概念は、AI 開発ライフサイクルに負荷がかかるため、より深く掘り下げる価値があります。この概念については、「Agentic XP: Moving Rigor Left in the Age of AI」で詳しく説明されています。
エクストリーム プログラミング (XP) は、ペア プログラミング、テスト駆動開発 (TDD)、小規模バッチ配信などの儀式でよく記憶されています。その核心となる洞察は、はるかに根本的なものでした。つまり、正しさは事後に検査されるのではなく、作成中に継続的に確立されるのです (Beck、2004)。 XP は、後から推測されるのではなく、開発中に正確さが明らかになるように整合性を維持することを目的としていました。
実際には、合意、作業方法、アーキテクチャを抽象化し、要件を自動化されたチェックとテストに変えることが必要になります。
予測可能なコードベースとプロセスは人間にとって採用しやすいものですが、私たちは AI が処理できるよりもはるかに高いレベルの曖昧さを許容することができ、また許容し続けてきました。したがって、同じことを行う正しい方法を 1 つだけ持つことは、AI によって生成されたソリューションの高品質を維持するために重要です。
この予測可能性は、静的コード分析や自動テスト、厳密なコンパイラー (Rust など) などの特殊なツールの使用によってもたらされますが、設計システム、API カタログ、ポリシーによって強制される仕様によってもたらされます。
アジャイル全体、特に XP は人間が間違いを犯すことを受け入れ、自動化の正確性を確保するために責任を転嫁します。作業が AI エージェントによって実行される場合、これはさらに重要になります。
「エージェント開発サバイバルガイド」の中で、私は「より多くのパワー - より多くのガードレール」の原則に従うことを主張しました。

これは組織全体での採用にも当てはまります。
ブレーンストーミングや使い捨てプロトタイプの作成など、重要ではないタスクについては、緩やかなルール (skills.md などの形式) を設定しても問題ありません。しかし、主要な業務や会社の評判にリスクをもたらすものについては、通路にコンクリートの壁を設け、人やエージェントが誤ってまたは意図的にそこから外れないようにすることが重要です。
そして、それが生産に関わるほぼすべてです。
アジャイル AI-DLC は、テスト自動化、継続的インテグレーションおよびデプロイ (CI/CD) パイプラインに依存し、自動化が不可能または実現可能でない場合にのみ文書化と人間によるレビューに頼ります。
アジャイル AI-DLC はどの成熟度レベルでも導入できますが、組織は DORA や SPACE などの指標で測定できるエンジニアリングの成熟度を高めることに取り組むことをお勧めします。
これらの基本的な指標に加えて、各「舗装されたパス」の導入を追跡し、その導入率、エラー率、再実行率を継続的に反復することをお勧めします。
パス採用率 : このパスを毎日、毎週、毎月使用しているユニーク ユーザーの数の尺度です。
パス エラー率 : 合計実行数に対する失敗したパス実行の割合です。
パス再実行率 : 一意のリクエスト数に対する同じリクエストによる実行の割合です。
そして最後に、すべてのパスで発生するコストをコスト予算とアラートの形で監視する必要があります。
アジャイル AI-DLC は、あらゆる規模と構造のチームに採用できますが、プラットフォーム チームとストリーム調整チームを組み合わせた場合に最も効果的に機能します。
プラットフォーム チーム : 舗装されたパスのメンテナンスとプラットフォーム機能の進化を担当します。パス、ドメイン、またはテクノロジーのクラスターを中心に編成できます。

。これらはプラットフォームの複雑さを吸収します。
ストリームに合わせたチーム : 一連の特定のビジネス目標や収益源など、結果ストリームを推進する責任を負います。彼らは複雑さを排除し、ビジネス上の問題を解決することに重点を置いています。
すべてのチームが目標とプラットフォームの両方に対するオーナーシップを感じることが重要です。プラットフォーム チームは、ストリームに沿ったチームと緊密に連携してニーズを予測し、必要な機能を積極的に構築する必要があります。また、ストリームに沿ったチームは、必要な、または将来必要になるプラットフォーム機能に対して提案や改善を行う必要があります。
組織に応じて、ドメインの専門家はプラットフォーム チームの一部となることも、独立したサテライト チームとして存在することもでき、専門知識とガイダンスを提供するために流れを調整したチームに引き込まれます。
アジャイル AI-DLC の全体的な成熟度は、エンジニアリングの成熟度と同様の方法で測定されます。つまり、企業が舗装されたパスをより積極的に使用し、成果が向上するほど、成熟度が高くなります。しかし、同社は最初の道が舗装された瞬間に採用される方法論を検討することができる。
アジャイル AI-DLC の核心は、アジャイルで柔軟な方法論であり、企業が既存の強みや働き方を強化し、反復的に改善できるようにします。
各プロジェクトに割り当てられた時間の少なくとも 20% を、日々の作業とプロセスの反映と改善に充てることをお勧めします。アジャイル AI-DLC やそのような改善を採用した結果として空いた時間は、可能な限り作業量を増やすのではなく、その反省と改善に充てるべきです。
アジャイル AI-DLC 手法に従うと、組織は AI で生成されたドキュメントをレビューするのではなく、成果、人間間のコラボレーション、戦略的思考に重点を置くことができます。
そして、これらの品質

将来がどうなるかに関係なく、これらは長期的な成功の前提条件です。
AI エージェントを使用したコーディングの時代を先取りしながら、アイデンティティの危機と開発者の劣悪なエクスペリエンスを乗り越えるための実用的なガイドです。
過去の AI 約束時代と、そもそもどのようにしてここに到達したかについてのエッセイ
個人の貢献者からリーダーに移行するのは簡単ではありません。それはなぜでしょうか?ある日、あなたはすべての作業を行う英雄ですが、次の日には他の人にそれを行うよう指導することが期待されます。マネージャーとして働き始めた頃に今の知識があればよかったのにと思います。おそらくこの記事は、私が陥ったいくつかの落とし穴を避けるのに役立つでしょう。

## Original Extract

A framework and methodology for transforming pre-AI software development lifecycle into AI-native development at scale.

Schrödinger's AI-DLC Framework

Agile AI Development Lifecycle - ValeriaVG ValeriaVG Blog
Agile AI Development Lifecycle
At the end of 2025, AWS open-sourced their AI-DLC methodology. At first it looks like an easy-to-adopt (just drop markdown files into your repository) framework that should take one from pre-AI software development lifecycle to the AI-native future. In its nature it is very similar to compound engineering and superpowers - all of them are a collection of skills in markdown prose that rely on human-in-the-loop review. Notably both of them were introduced to the public a few months before AWS released AI-DLC, but it's hard to compete with AWS in terms of visibility.
The AWS AI-DLC is designed to guide users through 30+ stages of development in total, grouped into 5 phases: initialization, ideation, inception, construction and operation. Every stage requires thorough mob-review of AI-generated artifacts.
The pros of this approach are ease of adoption and the fact that it is designed by AWS. The cons are that it is essentially an automated waterfall (4 out of 5 phases produce documentation) with quality being defined by how thorough and knowledgeable the reviewers are.
So unless your company is staffed by superhumans who enjoy reading tons of AI-generated documentation, you too might look for something else, especially if you need to scale the AI methodology outside of one engineer's machine working with a handful of repositories.
The second camp of AI adopters at scale doesn't have a single shared name for the methodology, because they didn't invent something new. They evolved their existing workflows and integrated them with AI capabilities.
Notable publications include CNCF Platforms White Paper , Spotify's Golden Paths and Netflix's "paved roads".
The core idea is to standardize their processes and tools to make it easier for the teams to execute on business initiatives without needing to reinvent the wheel. The maintenance of these "paved paths" would fall on the platform engineering teams, popularized by Team Topologies alongside stream-aligned teams.
When all the processes are standardized, transitioning from a mature engineering organisation to an AI-enabled engineering organisation is a matter of writing agentic workflows on top of the existing processes and platform capabilities.
I suggest we finally give this methodology a proper name (and hope that a cloud juggernaut would adopt it).
The key difference between AI-DLC and Agile AI-DLC is that the stages of development produce prototypes and their evolutions, not documentation, and the quality is assured by deterministic checks, rather than human review.
Agile AI-DLC is built on the concept of "paved paths": interconnected, standardized processes that guide the development from varying starting points to production deployment across the whole organisation, not just engineering teams.
Examples of paved paths can include:
A designer creates a prototype in Figma and triggers a build of an interactive prototype
A product manager tags a Slack bot to evolve the designer's prototype with full-stack capabilities
An engineer brings an internal platform from idea to production utilising existing system components
And examples of deterministic checks are:
Opinionated design system and abstractions
Strict linter code checks and type systems
Scorecards and CI/CD pipelines enforcing them
Agile AI-DLC, as the name suggests, follows the ground rules of Agile Development and can be described as the following cycles:
Proof of Concept : a quick, exploratory prototype is developed as the team collaboratively explores the problem space and validates the feasibility of the solution. The paved path used to develop the PoC is defined by the problem space and team's composition. This stage requires domain experts' involvement and the expected outcome is a decision on whether to proceed with the project or not.
Prototype : validated PoC is evolved into a prototype. The paved path depends on the team's composition (e.g. design-first, API-first, etc). Domain experts are involved to a lesser degree and the expected outcome of this stage is to define the scope and requirements for the MVP version.
Executable Specification : the requirements gathered during previous stages are translated into executable specifications: unit, integration, end-to-end or visual tests, contracts or strict machine-enforced rules as much as possible. Requirements that cannot be validated automatically are documented and stored in a manner that is accessible to team members, stakeholders and agents. Domain experts are involved on-demand and the outcome of this stage is a suite of test automation and executable specifications.
Implementation Loop : agents are tasked with evolving the prototype into a working implementation one requirement at a time until all the requirements are met and all the checks are passing.
Evaluation : the working implementation is validated against a broader set of requirements and checks, including manual checks, if needed, and demonstrating the solution to the stakeholders. Domain experts are involved on-demand and the outcome of this stage is a set of additional requirements that need to be addressed in the next iteration.
Retrospection : the team reflects on the implementation and makes adjustments to the paths, processes and requirements for the next iteration. Domain experts are involved on-demand and the outcome of this stage is a set of commitments that kicks off the next cycle.
The concept of requirements as executable specifications rather than documentation deserves a deeper dive as it is load-bearing for the AI development lifecycle. The concept is well-described in Agentic XP: Moving Rigour Left in the Age of AI :
Extreme Programming (XP) is often remembered for its rituals: pair programming, test-driven development (TDD) and small batch delivery. Its core insight was far more fundamental: correctness is established continuously during creation rather than inspected after the fact (Beck, 2004). XP was about maintaining alignment so that correctness emerged during development rather than being inferred later.
In practice that entails turning the agreements, ways of working and architecture into abstractions and the requirements into automated checks and tests.
Predictable codebases and processes are easier to adopt for a human, but we can and have been tolerating a far greater level of ambiguity than AI can handle. Therefore having one and only one correct way of doing the same thing is crucial for maintaining high quality of AI-generated solutions.
This predictability can come from the use of specialised tools, such as static code analysis and automated testing, strict compilers (such as the one Rust has, for example), but also from specifications enforced by design systems, API catalogs and policies.
Agile overall, and XP in particular, accepts that humans make mistakes and shifts the responsibility to ensure correctness on automation. Which becomes even more important if the work is executed by AI agents.
In the "Agentic Development Survival Guide" , I insisted on following the principle of "More Power - More Guardrails" which holds for the organisation-wide adoption as well.
It's okay to set loose rules (e.g. in the form of skills.md) for non-critical tasks, such as brainstorming or creating a throw-away prototype. But for anything that carries risks to main operations or company reputation, it is important to make sure that paths have concrete walls, preventing people and agents from accidentally or intentionally straying from it.
And that is almost everything that touches production.
Agile AI-DLC relies on test automation, continuous integration and deployment (CI/CD) pipelines and only resorts to documentation and human review when it isn't possible or feasible to automate it.
While Agile AI-DLC can be adopted at any level of maturity, it is recommended that the organisation commits to increasing their engineering maturity that can be measured by metrics such as DORA and SPACE .
In addition to these fundamental metrics, it is recommended to track adoption of each "paved path" and continuously iterate on their adoption, error and re-run rates, where:
Path Adoption Rate : is a measure of how many unique users are using this path daily, weekly and monthly.
Path Error Rate : is a percentage of failed path executions over the total number of executions.
Path Re-run Rate : is a percentage of executions with the same request over the number of unique requests.
And finally, every path should be monitored for costs it incurs in the form of cost budgets and alerts.
Agile AI-DLC can be adopted by teams of any size and structure, but it works best with the combination of Platform Teams and Stream-Aligned Teams .
Platform Teams : are responsible for the maintenance of the paved paths and evolution of platform capabilities. They can be organised around a cluster of paths, a domain or a technology. They absorb complexity of the platform.
Stream-Aligned Teams : are responsible for driving an outcome stream, such as a set of specific business goals or revenue streams. They focus on solving business problems, shedding the complexity.
It is important that all teams feel ownership over both goals and the platforms: platform teams should work closely with and anticipate the needs of stream-aligned teams to proactively build required capabilities and stream-aligned teams should make suggestions and even improvements to platform capabilities that they need or will need in the future.
Depending on the organisation, domain experts can be part of platform teams or exist as independent satellite teams, being pulled in by stream-aligned teams to provide expertise and guidance.
The overall maturity of Agile AI-DLC is measured in a similar way to engineering maturity: the more actively used paved paths the company has and the better the outcomes become - the higher the maturity. But the company can consider the methodology adopted the moment the first path is paved.
In its core, Agile AI-DLC is an agile and flexible methodology, allowing companies to enhance their existing strengths and ways of working and iteratively improve them.
It is recommended to dedicate at least 20% of time allocated to each project to reflect and improve daily work and processes. The time that would be freed as a result of adopting Agile AI-DLC and such improvements should be dedicated to their reflections and improvements rather than scaling the amount of work whenever possible.
Following the Agile AI-DLC methodology, you're ensuring that your organisation is focusing on the outcomes, collaboration between humans and strategic thinking, rather than reviewing AI-generated documentation.
And these qualities are prerequisites for long-term success, no matter what the future might hold.
A practical guide to navigating identity crisis and poor developer experience while trying to stay ahead of the curve of coding with AI agents.
An essay about the past AI-promises-era and how we got here in the first place
Transitioning from an individual contributor to a leader is not easy. Why would it be? One day you are the hero doing all the work, the next day you are expected to guide others to do it for you. I wish I had the knowledge I have now when I was starting out as a manager. Perhaps this article will help you avoid some pitfalls I dug myself into.
