---
source: "https://www.anthropic.com/news/alberta-government-claude-cybersecurity"
hn_url: "https://news.ycombinator.com/item?id=48809198"
title: "Government of Alberta uses Claude to find and fix cybersecurity vulnerabilities"
article_title: "Government of Alberta uses Claude to find and fix cybersecurity vulnerabilities \\ Anthropic"
author: "surprisetalk"
captured_at: "2026-07-06T19:42:28Z"
capture_tool: "hn-digest"
hn_id: 48809198
score: 1
comments: 0
posted_at: "2026-07-06T19:15:37Z"
tags:
  - hacker-news
  - translated
---

# Government of Alberta uses Claude to find and fix cybersecurity vulnerabilities

- HN: [48809198](https://news.ycombinator.com/item?id=48809198)
- Source: [www.anthropic.com](https://www.anthropic.com/news/alberta-government-claude-cybersecurity)
- Score: 1
- Comments: 0
- Posted: 2026-07-06T19:15:37Z

## Translation

タイトル: アルバータ州政府がサイバーセキュリティの脆弱性を発見し修正するためにクロードを使用
記事のタイトル: アルバータ州政府がサイバーセキュリティの脆弱性を見つけて修正するためにクロードを使用 \ Anthropic
説明: アルバータ州政府は、システムをレビューし、脆弱性を見つけて修正するために、Opus モデルと Sonnet モデルの両方で Claude Code を使用してきました。

記事本文:
アルバータ州政府はサイバーセキュリティの脆弱性を発見し修正するためにクロードを使用 \ Anthropic メインコンテンツにスキップ フッターにスキップ 研究
ケーススタディ アルバータ州政府は、政府システム全体のサイバーセキュリティの脆弱性を発見して修正するためにクロードを使用しています
2025 年以来、アルバータ州政府は Opus モデルと Sonnet モデルの両方でクロード コードを使用してシステムをレビューし、脆弱性を見つけて修正しています。アルバータ州技術イノベーション省内のチームは、20 時間で 4 億 6,600 万行のコードをスキャンし、システム全体のセキュリティ ギャップを修正し、それらのシステムをより安全にするための新しいツールを構築しました。
私たちは、政府機関がクロードとクロード コードを使用して大規模なシステムをセキュリティで保護する方法の例として、彼らの経験の詳細を共有しています。政府はメリットを提供し、サービスを継続的に実行するためにこれらのシステムに依存しているため、これは重要な課題です。ただし、コードは多くの場合古く、安全性が低く、文書化が不完全です。アルバータ州はまた、他の政府が学ぶことができるようにその取り組みを文書化した技術白書集も出版している。ここで読むことができます。
アルバータ州のネイト・グルービッシュ技術革新大臣は、「アルバータ州の人々は、生活の中で最も機密性の高い情報に関して政府を信頼しており、それを保護するのが我々の責任である」と述べた。 「AI を使用してシステム全体の脆弱性を見つけて修正することで、従来のアプローチでは完了までに何年もかかっていた作業を数時間で完了しました。これが AI 時代の責任ある政府の姿であり、最善のものはまだ私たちの先にあります。」
アルバータ州の技術革新省は、社会サービスから公共安全、山火事対応まで、27 州すべての省庁のシステムを維持しています。これには、約 1,280 件のアプリケーションと 3 件のアプリケーションが含まれます。

,400のコードリポジトリ。そのほとんどは体系的なセキュリティレビューを受けておらず、安全でないコード、未対処のバグ、古いソフトウェアなどの蓄積された技術的負債は数十億ドルに達します。
同省のシステムには、納税記録、政府調達データ、社会福祉事件ファイルなどの機密性の高い情報が保管されています。そこで同省は 2025 年に、これらのシステムをより安全にし、長期にわたって維持しやすくするという使命を帯びた内部チームを設立し、クロードと協力してその実現に取り組みました。
すでに同省はクロードを次の目的で利用しています。
4 億 6,600 万行の政府コードを 20 時間で評価します。チームは、Claude Opus および Sonnet モデルを備えた Claude Code を使用して、管理しているコードベースで Claude を動作させました。約 50 人のエージェントが自律的かつ並行してシステムをスキャンし、セキュリティの脆弱性、基盤となるインフラストラクチャと展開プロセスの弱点、技術文書のギャップを見つけました。 Claude Code は 2 段階のルーチンを実行し、最初にルール エンジンで各リポジトリをスキャンして既知のパターンにフラグを立て、次にそれらのフラグを確認して、開発者が検証できるように、検出結果ごとに正確なファイルと行を引用しました。スキャンではアルバータ州が所有するすべてのリポジトリが対象となり、従来の自動スキャン ツールでは見逃されていた問題が特定されました。アルバータ州の実装には約 20 時間かかりました。研究チームは、そうでなければこの種の検討には約6年半かかった可能性があると推定している。
スキャンで見つかった脆弱性を修正します。スキャンによって脆弱性が特定された場合、Claude Code は多くの場合、修正を生成し、テストし、ビルドできます。パッチが安全であることを確認するために必要な自動テストがシステムにない場合、クロードは最初にテストを作成しました。コードが古すぎるか、既存の形式で効率的にパッチを適用するには複雑すぎる場合、クロードはそれを再構築しました。

より現代的で保守しやすい言語で。一部のシナリオでは、これらのシステムは、約 25 年前に Java で手作業でコーディングされ、最初の構築に 5 か月かかった補助金プログラム ポータルを含め、わずか 4 ～ 5 日で再構築できます。これらはすべて同省のエンジニアと協力して行われました。パッチは出荷前にチームによってレビューされ、承認されました。
継続的なセキュリティレビューを実行します。アルバータ州のサイバーセキュリティ チームは、開発プロセス全体を通じて実行される一連の専門のクロード レビュー エージェントも構築しました。 「レッド チーム」エージェントは、攻撃者と同じように外部からアプリケーションを調査し、脆弱性が悪用される方法をマッピングします。次に、「ブルー チーム」エージェントが国際セキュリティ標準に対するアプリケーションの防御を評価し、修正する正確なファイルを示す修正計画を作成します。追加のエージェントがコードの品質と一般の人々が目にする文章の明瞭さをチェックします。すべてのアプリケーションは、パスごとに約 95 のセキュリティ制御に対してチェックされます。これらのエージェントは Claude Agent SDK 上に構築されており、すべてのアプリケーションに対して堅牢な一連のチェックと分析を実行します。
アルバータ州は、独自のシステムのスキャン、保護、最新化に加えて、アルバータ AI アカデミーを通じて政府職員と一般の人々の両方に AI の使用方法を訓練しています。数千人の政府職員と 10,000 人を超える一般の人々がこのプラットフォームを使用して、プロンプトからエンタープライズ アプリケーションの配信まで、AI の効果的な使用の基礎を学びました。技術革新省はアカデミーを通じて、そのアプローチを単一のチームやプロジェクトを超えて、それを必要とするすべての省庁に拡大することを目指しています。
現在、クロードは同省の近代化に役立つコードの作成、レビュー、導入を支援しています。

ションの取り組み。次に、エンジニアと一緒にまったく新しいソフトウェアやツールを構築できる AI エージェントとの連携を拡大する予定です。
アルバータ州政府も近代化工事を継続する予定です。たとえば、ある省では 185 個のレガシー アプリケーションが運用環境で実行されていますが、これらのアプリケーションは保守に費用がかかり、更新も困難です。政府は、Claude Code を使用してこれらのシステムを分析し、その動作を理解し、最新のコーディング言語と規約に基づいて構築された 16 の再利用可能なアプリケーションに統合することを計画しています。目標は、複雑さを軽減し、メンテナンスコストを削減し、完了までに数年かかるであろう近代化作業をスピードアップすることです。
アルバータ州政府が対処しようと取り組んでいる技術的負債とセキュリティの脆弱性は、決して珍しいものではありません。これらは、世界中の多くの州、州、連邦機関のシステムに存在します。アルバータ州が発表した技術白書は、他の政府にも同様の問題に対処するための青写真を与えている。
ホワイトペーパーに加えて、アルバータ州は7月にエドモントンで業界デーを主催し、学んだことを共有する予定だ。そしてこの秋には、そのアプローチを州政府全体に拡大するプログラムを開始する予定だ。私たちはアルバータ州がこれらの取り組みを拡大する際に協力し続け、文書化されたアプローチが他の政府が独自のシステムを保護するのに役立つことを願っています。
Fable 5 のサイバー セーフガードとジェイルブレイク フレームワークの詳細
Sonnet 5 は、大規模なコーディング、エージェント、専門的な作業にわたって最先端のパフォーマンスを提供します。
Fable 5 は 7 月 1 日に全世界で復活します。また、Amazon、Microsoft、Google、その他の Glasswing パートナーと協力して、ジェイルブレイクの重大度をスコアリングするための業界全体のフレームワークも提案しています。
消費者の健康データのプライバシー ポリシー

## Original Extract

The Government of Alberta has been using Claude Code with both Opus and Sonnet models to review its systems, find vulnerabilities, and fix them.

Government of Alberta uses Claude to find and fix cybersecurity vulnerabilities \ Anthropic Skip to main content Skip to footer Research
Case Study Government of Alberta uses Claude to find and fix cybersecurity vulnerabilities across government systems
Since 2025, the Government of Alberta has been using Claude Code with both Opus and Sonnet models to review its systems, find vulnerabilities, and fix them. A team inside Alberta’s Ministry of Technology and Innovation scanned 466 million lines of code in 20 hours, remediated security gaps across its systems, and built new tools to make those systems safer.
We’re sharing details of their experience as an example of how government agencies can use Claude and Claude Code to secure their systems at a large scale. This is a critical challenge, as governments rely on these systems to deliver benefits and keep services running—yet the code is often old, insecure, and incompletely documented. Alberta has also published a collection of technical white papers documenting its efforts for other governments to learn from; you can read them here.
“Albertans trust their government with some of the most sensitive information in their lives, and it is our responsibility to protect it,” said Nate Glubish, Alberta’s Minister of Technology and Innovation. “By using AI to find and fix vulnerabilities across our systems, we accomplished in hours what would have taken a traditional approach years to complete. This is what responsible government looks like in the AI era, and the best is still ahead of us.”
Alberta’s Ministry of Technology and Innovation maintains the systems of all 27 provincial ministries, from social services to public safety to wildfire response. That includes roughly 1,280 applications and 3,400 code repositories. Most of it has never undergone a systematic security review, and the accumulated technical debt—insecure code, unaddressed bugs, outdated software—runs into the billions of dollars.
The Ministry’s systems hold highly sensitive information, including tax records, government procurement data, and social services case files. So in 2025, the Ministry set up an internal team with a mandate to make these systems more secure and easier to maintain over time, working with Claude to do so.
Already, the Ministry has used Claude to:
Assess 466 million lines of government code in 20 hours. The team put Claude to work on the codebases it maintains, using Claude Code with Claude Opus and Sonnet models. Around 50 agents worked autonomously and in parallel to scan the systems for security vulnerabilities, weaknesses in underlying infrastructure and deployment processes, and gaps in technical documentation. Claude Code ran a two-stage routine, first scanning each repository with a rules engine to flag known patterns, then reviewing those flags and citing the exact file and line for each finding so developers could verify them. The scan covered every repository Alberta owns and identified issues that traditional automated scanning tools had missed. It took about 20 hours for Alberta’s implementation; the team estimates that that kind of review could otherwise have taken around 6.5 years.
Fix the vulnerabilities the scan found. Where the scan identified a vulnerability, Claude Code could often generate a fix, test it, and build it. In cases where a system lacked the automated tests needed to confirm that a patch was safe, Claude wrote the tests first. Where the code was too outdated or too complex to patch efficiently in its existing form, Claude rebuilt it in a more modern and maintainable language. In some scenarios, these systems could be rebuilt in as little as four to five days, including a subsidy program portal that was originally hand-coded in Java roughly 25 years ago and took five months to build the first time. All of this was done in partnership with the Ministry’s engineers: before any patch shipped, it was reviewed and approved by the team.
Run continuous security review. Alberta’s cybersecurity team also built a set of specialized Claude review agents that run throughout the development process. A “red team” agent probes an application from the outside, the way an attacker might, and maps how a vulnerability might be exploited. A “blue team” agent then assesses the application’s defenses against an international security standard, and writes a remediation plan that points to the exact files to fix. Additional agents check code quality and the clarity of the writing the public sees. Every application is checked against roughly 95 security controls on each pass. These agents are built on top of the Claude Agent SDK and run a robust series of checks and analysis for every application.
In addition to scanning, securing, and modernizing its own systems, Alberta is training both government workers and the public in the use of AI through the Alberta AI Academy. Thousands of government employees and more than 10,000 members of the public have used the platform to learn the essentials of effective AI use, from prompting through enterprise application delivery. Through the Academy, the Ministry of Technology and Innovation aims to extend its approach beyond a single team or project to every ministry that needs it.
Today, Claude helps the Ministry write, review, and deploy code that aids in its modernization efforts. Next, it plans to expand that work with AI agents that can build entirely new software and tools alongside engineers.
The Government of Alberta also plans to continue its modernization work. One ministry, for example, has 185 legacy applications running in production, which are expensive to maintain and difficult to update. The Government plans to use Claude Code to analyze these systems, understand what they do, and consolidate them into 16 reusable applications built on modern coding languages and conventions. The goal is to reduce complexity, lower maintenance costs, and speed up modernization work that would otherwise take years to complete.
The technical debt and security vulnerabilities the Government of Alberta is working to address are hardly unique. They exist in the systems of many provinces, states, and federal agencies across the world. The technical white papers Alberta has released give other governments a blueprint for addressing these same issues.
In addition to the white papers, Alberta is hosting an industry day in Edmonton in July to share what it has learned. And this fall, it will begin a program to scale its approach across the provincial government. We’ll keep working with Alberta as it expands these efforts, and we hope the approach it has documented can help other governments secure their own systems.
More details on Fable 5’s cyber safeguards and our jailbreak framework
Sonnet 5 delivers frontier performance across coding, agents, and professional work at scale.
Fable 5 returns globally July 1. We're also proposing an industry-wide framework for scoring jailbreak severity, together with Amazon, Microsoft, Google, and other Glasswing partners.
Consumer health data privacy policy
