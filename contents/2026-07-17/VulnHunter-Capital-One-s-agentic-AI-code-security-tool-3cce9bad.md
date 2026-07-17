---
source: "https://www.capitalone.com/tech/open-source/announcing-vulnhunter/"
hn_url: "https://news.ycombinator.com/item?id=48946692"
title: "VulnHunter: Capital One's agentic AI code security tool"
article_title: "VulnHunter: an open-source, agentic AI code security tool | Capital One Tech"
author: "medina"
captured_at: "2026-07-17T13:05:31Z"
capture_tool: "hn-digest"
hn_id: 48946692
score: 3
comments: 1
posted_at: "2026-07-17T12:42:12Z"
tags:
  - hacker-news
  - translated
---

# VulnHunter: Capital One's agentic AI code security tool

- HN: [48946692](https://news.ycombinator.com/item?id=48946692)
- Source: [www.capitalone.com](https://www.capitalone.com/tech/open-source/announcing-vulnhunter/)
- Score: 3
- Comments: 1
- Posted: 2026-07-17T12:42:12Z

## Translation

タイトル: VulnHunter: Capital One のエージェント AI コード セキュリティ ツール
記事のタイトル: VulnHunter: オープンソースのエージェント AI コード セキュリティ ツール |キャピタルワンテック
説明: Capital One が開発したエージェント推論ワークフローを備えたオープンソース AI コード セキュリティ ツールである VulnHunter のリリースを発表します。

記事本文:
Capital One のオープンソースのエージェント AI コード セキュリティ ツール。
ソフトウェア セキュリティのルールは、ほとんどの防御者が追いつくことができないほどの速さで変化しています。
高度な AI モデルにより、悪意のある者がソフトウェアの脆弱性を発見して悪用する障壁が大幅に低くなりました。かつてはかなりのスキルと時間を必要としていたものを、現在では自動化、高速化、拡張できるようになりました。世界は、高度に洗練された次世代 AI 攻撃機能が手頃な価格で事実上すべての敵にアクセスできるようになるまでの時間がますます短くなっていることに直面しています。業界全体で、組織はこのパラダイムシフトへの準備を急いでいます。
ネットワークのセグメンテーション、ID 制御、監視などの従来の環境保護は依然として不可欠ですが、もはやそれだけでは十分ではありません。この新たな現実における最終的な防御には、アプローチの変更が必要です。組織は、攻撃者が高度なモデルを導入して脆弱性を発見し悪用する前に、コード内の脆弱性を検討して検出し、修正する必要があります。
Capital One では、AI を利用した脅威に対する正しい対応は待つことではなく、最先端の AI 主導の防御を構築し、それをあらゆる場所の防御者の手に委ねることであると判断しました。
そのため、私たちは本日、プロアクティブで攻撃者の観点からの分析をソース コードに直接適用するように設計された高度なエージェント AI セキュリティ ツールである VulnHunter のオープンソース リリースを発表します。
Capital One の内部で開発された VulnHunter は、従来の受動的脆弱性スキャナーではありません。これは、潜在的に悪用可能な欠陥を特定し、予想される攻撃経路をマッピングし、高度にターゲットを絞ったコード修復を提案するエージェント推論ワークフローを備えた防御ツールの変化を表しています。
開発者のエクスペリエンスのために構築
VulnHu のユーティリティを完全にロック解除するには

使いやすさが重要であることはわかっていました。従来のセキュリティ ツールの永続的な課題は、開発者の実際の日常のワークフローがあまり考慮されず、主に厳密なサイバーセキュリティの実践を強制することを目的として構築されていることが多いことです。
VulnHunter を構築する際には、開発者第一の考え方を導入しました。私たちは、セキュリティ ツールがエンタープライズ規模で成功する唯一の方法は、開発者が実際に使用したいものであるかどうかであることを知っていました。
私たちは、重要な瞬間における開発者のエクスペリエンスを非常に効率的にすることに重点を置きました。 VulnHunter は、開発プロセス全体を通じて従来の摩擦点を意図的に取り除き、最小限に抑えることで、開発者の負担を誤警報のトリアージから遠ざけます。代わりに、ワークフローは証拠に裏付けられた即時のコード修復に重点を置いています。
内部: VulnHunter のユニークな機能
VulnHunter は、投機的なアラートを最小限に抑え、実用的な修復を最大限に高めるために設計されたいくつかの主要な技術革新を導入しています。
自身の結論に異議を唱えるように設計された改ざんエンジン: 私たちの目標は、開発者に届く前に誤検知を最小限に抑えることです。何らかの発見を明らかにした後、VulnHunter は、自身の主張を反証するために特別に設計された構造化された推論ワークフローを実行します。この改ざんエンジンは、当てはまらない仮定、悪用経路の論理的なギャップ、攻撃の成功を妨げる条件を積極的に検索します。サポートされていない仮定に依存する結果は直ちに破棄されるように設計されています。
その結果、開発者の注意を引いたものは、すでに社内の厳しい課題を乗り越えたものになります。フラグが立てられたすべての脆弱性は、ツールが除外しようとして失敗した脆弱性です。
攻撃者優先のフォワード分析: 従来のツールでは、多くの場合「シンクファースト」分析を活用し、潜在的な攻撃の可能性を検討しています。

危険なコード パターンを分離して、仮想の攻撃者を逆方向に検索します。このアプローチでは、エンジニアリング チームに誤検知が大量に発生する可能性があります。
VulnHunter は、このモデルを反転して、悪役の正確な行動をシミュレートします。 API、ネットワーク メッセージ、ファイル アップロードなど、攻撃者がアクセスできる可能性のあるエントリ ポイントから始まり、アプリケーション ロジック、データ変換、内部セキュリティ チェックポイントを通じて理由が進みます。 VulnHunter は、攻撃者が実際にシステムとどのように対話するかをモデル化することで、攻撃者が本当に突破できるかどうかを評価します。
証拠に裏付けられた修復モデリング: 欠陥が改ざんエンジンを通過しても、VulnHunter は単に警報を鳴らして開発者に推測を任せるだけではありません。問題を見つけることから、それを解決するために取り組むことに変わります。
VulnHunter はコードベース全体で裏付けとなる証拠を収集し、生き残っているエクスプロイト パス全体を計画します。これは、欠陥の明確な説明を提供し、攻撃者が取得する特定の機能またはアクセスを詳細に説明し、エンジニアリング レビューのために焦点を絞ったターゲットを絞ったコード変更を生成するように設計されています。
VulnHunter をコミュニティにリリースする前に、私たちは独自のコードで VulnHunter を実行しました。私たちは、数十のビジネス分野にまたがる数千のリポジトリにわたる脆弱性を迅速かつ効率的に特定し、修復することができました。これまで私たちのチームは多大な時間を要し、手動でトリアージを行っていましたが、今では検証済みで実用的な結果が迅速かつ効果的に得られます。
集団的自衛権への取り組み
現代のソフトウェア サプライ チェーンは深く相互接続されています。広く使用されているオープンソース コンポーネントの 1 つの脆弱性が、数千の企業に同時に波及する可能性があります。単一の組織だけではこの課題を解決できないため、私たちは VulnHunter をオープンソース化しています。この現実に対処するための防御ツールは次のとおりです。

保護するコードベースと同様に、広く配布、テスト、改善する必要があります。
Capital One のオープンなコラボレーションへの取り組みに基づいて、VulnHunter のリリースにより、より広範なテクノロジーおよびセキュリティ コミュニティがワークフローを検査し、その前提に疑問を呈し、この新しい防御アプローチの改善に貢献できるようになります。
GitHub で VulnHunter を使ってみる
VulnHunter は現在利用可能です。これを実行するには、Claude Opus 4.8 にアクセスし、動作する Claude Code 環境にアクセスする必要があります。リポジトリには、クイックスタート ガイド、アーキテクチャ ドキュメント、VulnHunter がコード パスをトレースして修復を生成する方法を示す注釈付きのサンプル ワークフローが含まれています。既知の制限事項と現在の開発ロードマップはリポジトリに文書化されています。
バグの報告、推​​論ワークフローへの変更の提案、モデル サポートの拡張など、貢献したい場合は、CONTRIBUTING.md に問題とプル リクエストを送信するためのプロセスの概要が記載されています。
リポジトリ: github.com/capitalone/vulnhunter
モデルの最適化: Claude Opus 4.8 モデル
初期実装: クロード コード スキル
VulnHunter は Claude Opus 4.8 と Claude Code を念頭に置いて作成および活用されていますが、そのフレームワークとスキルはコーディング ハーネスや基盤モデル全体で活用できる可能性があります。
脅威の状況は待っていません。 VulnHunter は、攻撃者が脆弱性に到達する前に、防御者が脆弱性を発見して修正できる、より厳密で証拠に基づいた方法を提供するために構築されました。安全なソフトウェアは、開発者、企業、そして私たち全員が構築するシステムに依存する人々に利益をもたらす共有基盤であるため、私たちはこれをリリースしています。コミュニティが次にこれを使って何を構築するのかを楽しみにしています。
Capital One Tech について詳しく見る
GitHub で VulnHunter を使ってみましょう。
資本の仕組みを探る

1 つは、複雑な課題を大規模に解決する AI の構築です。
キャリアの機会を見つけてください。
当社は、大規模なリアルタイム データ、AI、機械学習、クラウドの力を利用して、業界の困難な問題を解決します。
ゼロトラスト革命: 従来のネットワーク セキュリティが失敗する理由
デジタル脅威の状況の変化により、ネットワーク セキュリティの根本的な変化が求められています。
コンテキスト エンジニアリング: オープンソースのコンテキスト仕様の紹介
第 1 回 Capital One AI シンポジウムからの洞察
Capital One Web サイトから離れようとしています
開示声明: © 2026 Capital One.意見は著者個人のものであり、必ずしも Capital One の意見ではありません。特に明記されていない限り、Capital One は言及されている第三者とは提携も承認もしておらず、リンクされている第三者サイトのコンテンツやプライバシー ポリシーについては責任を負いません。使用または表示される商標およびその他の知的財産は、それぞれの所有者の財産です。

## Original Extract

Announcing the release of VulnHunter, an open-source AI code security tool with an agentic reasoning workflow developed by Capital One.

Capital One’s open-source, agentic AI code security tool.
The rules of software security are changing faster than most defenders can keep pace.
Advanced AI models have dramatically lowered the barrier for bad actors to discover and exploit vulnerabilities in software. What once required significant skill and time can now be automated, accelerated, and scaled. The world faces an increasingly short window of time before highly sophisticated, next-generation AI attack capabilities become affordable and accessible to virtually every adversary. Across the industry, organizations are racing to prepare for this paradigm shift.
Traditional environmental protections like network segmentation, identity controls, and monitoring remain essential, but are no longer sufficient on their own. The ultimate defense in this new reality requires a shift in approach: organizations need to consider and detect the vulnerabilities in their code and fix them before adversaries can deploy advanced models to discover and exploit them.
At Capital One, we decided that the right response to AI-enabled threats wasn't to wait, but to build cutting-edge AI-driven defenses and put them in the hands of defenders everywhere.
That’s why we are announcing today the open-source release of VulnHunter , an advanced agentic AI security tool designed to apply proactive, attacker-perspective analysis directly to the source code.
Developed internally at Capital One, VulnHunter is not a traditional, passive vulnerability scanner. It represents a shift in defensive tooling with an agentic reasoning workflow to identify potentially exploitable defects, map prospective attack paths, and propose highly targeted code remediations.
Built for the developer experience
To fully unlock the utility of VulnHunter, we knew ease of use mattered. A persistent challenge with traditional security tools is that they are often built primarily to enforce rigid cybersecurity practices, without much consideration for a developer’s actual day-to-day workflow.
We brought a developer-first mindset when building VulnHunter. We knew the only way a security tool can be successful at enterprise scale is if it is something developers actually want to use.
We focused on making the developer experience highly efficient in the moments that matter. By intentionally rounding out and minimizing traditional points of friction throughout the development process, VulnHunter shifts the developer's burden away from triaging false alarms. Instead, the workflow is focused on immediate, evidence-backed code repair.
Under the hood: the unique capabilities of VulnHunter
VulnHunter introduces several key technical innovations designed to minimize speculative alerts and maximize actionable repair:
Falsification engine designed to challenge its own conclusions: Our goal is to minimize false positives before they ever reach a developer. After surfacing any finding, VulnHunter runs a structured reasoning workflow specifically designed to disprove its own argument. This falsification engine actively searches for assumptions that don't hold, logical gaps in the exploit path, and conditions that would prevent the attack from succeeding. It is designed to immediately discard findings that rely on unsupported assumptions.
The result: what reaches a developer's attention has already survived a rigorous internal challenge. Every flagged vulnerability is one the tool has tried and failed to rule out.
Attacker-first forward analysis: Conventional tools often leverage “sink-first” analysis, looking at potentially dangerous code patterns in isolation to search backward for a hypothetical attacker. This approach can flood engineering teams with false positives.
VulnHunter flips this model to simulate a bad actor’s exact journey. It begins at potential attacker-accessible entry points — such as APIs, network messages, or file uploads — and reasons forward through application logic, data transformations, and internal security checkpoints. By modeling how an attacker actually interacts with a system, VulnHunter evaluates whether an attacker can truly break through.
Evidence-backed remediation modeling: When a defect survives the falsification engine, VulnHunter doesn't just sound the alarm and leave the guesswork to developers. It shifts from finding the problem to working to solve it.
VulnHunter gathers supporting evidence across the codebase to map out the entire surviving exploit path. It is designed to provide a clear explanation of the defect, detail the specific capabilities or access an attacker would gain, and generate focused, targeted code changes for engineering review.
Before releasing VulnHunter to the community, we ran it on our own code. We were able to identify and remediate vulnerabilities across thousands of repositories, spanning tens of business areas, with speed and efficiency. What took our teams significant time and manual triage before now produces verified, actionable findings quickly and effectively.
A commitment to collective defense
Modern software supply chains are deeply interconnected. A single vulnerability in a widely-used open-source component can ripple across thousands of enterprises simultaneously. We’re open-sourcing VulnHunter because no single organization can solve this challenge alone. The defensive tools to address this reality need to be just as widely distributed, tested, and improved as the codebases they protect.
Building on Capital One’s commitment to open collaboration, the release of VulnHunter enables the broader tech and security community to inspect the workflow, challenge its assumptions, and contribute improvements to this new defensive approach.
Get started with VulnHunter on GitHub
VulnHunter is available now. To run it, you will need access to Claude Opus 4.8 and access to a working Claude Code environment. The repository includes a Quickstart guide, architecture documentation, and annotated example workflows showing how VulnHunter traces code paths and generates remediations. Known limitations and the active development roadmap are documented in the repository.
If you want to contribute — whether that's reporting a bug, proposing a change to the reasoning workflow, or expanding model support — the CONTRIBUTING.md outlines the process for submitting issues and pull requests.
Repository: github.com/capitalone/vulnhunter
Model Optimization: Claude Opus 4.8 model
Initial Implementation: Claude Code skill
While VulnHunter was authored and leveraged with Claude Opus 4.8 and Claude Code in mind, the framework and skills have potential to be leveraged across coding harnesses and foundation models.
The threat landscape isn't waiting. We built VulnHunter to give defenders a more rigorous, evidence-driven way to find and fix vulnerabilities before attackers can reach them. We're releasing it because secure software is a shared foundation that benefits developers, enterprises, and the people who depend on the systems we all build. We look forward to seeing what the community builds with this next.
Learn more about Capital One Tech
Get started with VulnHunter on GitHub .
Explore how Capital one is building AI to solve complex challenges at scale .
Discover career opportunities .
We use real-time data at scale, AI and machine learning and the power of the cloud to solve challenging industry problems.
Zero trust revolution: Why legacy network security fails
The shifting landscape of digital threats demands a fundamental change in network security.
Context engineering: Introducing open-source Context Specs
Insights from the inaugural Capital One AI Symposium
You are now leaving the Capital One website
DISCLOSURE STATEMENT: © 2026 Capital One. Opinions are those of the individual author and are not necessarily those of Capital One. Unless noted otherwise, Capital One is not affiliated with, nor endorsed by, any third parties mentioned and is not responsible for the content or privacy policies of any linked third-party sites. Any trademarks and other intellectual property used or displayed are property of their respective owners.
