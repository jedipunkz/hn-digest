---
source: "https://aboutcode.org/blog/agentic-scancode-port-case-study/"
hn_url: "https://news.ycombinator.com/item?id=48369401"
title: "An AI agent ported our codebase from Python to Rust"
article_title: "An AI agent ported our codebase from Python to Rust | AboutCode.org"
author: "Tiberium"
captured_at: "2026-06-03T00:47:46Z"
capture_tool: "hn-digest"
hn_id: 48369401
score: 3
comments: 0
posted_at: "2026-06-02T12:33:55Z"
tags:
  - hacker-news
  - translated
---

# An AI agent ported our codebase from Python to Rust

- HN: [48369401](https://news.ycombinator.com/item?id=48369401)
- Source: [aboutcode.org](https://aboutcode.org/blog/agentic-scancode-port-case-study/)
- Score: 3
- Comments: 0
- Posted: 2026-06-02T12:33:55Z

## Translation

タイトル: AI エージェントがコードベースを Python から Rust に移植しました
記事のタイトル: AI エージェントがコードベースを Python から Rust に移植しました | Code.org について
説明: 個別のインシデントではなくケーススタディ

記事本文:
メイン コンテンツにスキップ はじめに ブログ プロジェクト 環境 Contribute について
AI エージェントがコードベースを Python から Rust に移植しました
ClearlyDefined の 3 年間のロードマップ
OSI と AboutCode は ClearlyDefined を維持するために提携しています
ScanCode LicenseDB -- 公開データベースで厳選された 2,000 以上のライセンス
VulnerableCode API の非推奨と V3 の導入
脆弱性のない依存関係の解決
そもそもデュアルライセンスとは何でしょうか？
行ったり来たり -- ソフトウェアのバージョン管理の話
ScanCode のライセンス明瞭性スコア
Java アプリケーションでのコピーレフト ライセンスのソフトウェア コンポーネントの使用
AI エージェントがコードベースを Python から Rust に移植しました
個別の事件ではなくケーススタディ
ScanCode は、ソース コードとバイナリ ファイルの両方で、ライセンス、著作権、パッケージの依存関係、脆弱性などを検出します。使用例には、ライセンスとセキュリティのコンプライアンス、ソフトウェア サプライ チェーン管理が含まれます。これは、700 人を超える寄稿者からなるオープンソース コミュニティによる 10 年以上にわたる慎重な設計、アーキテクチャ、テストの成果であり、ライセンス検出だけをカバーする 40,000 を超える自動テスト、全体で 90,000 を超える自動テストをサポートしています。
コア モジュールは、業界をリードするオープン ソース コード スキャン エンジンである ScanCode Toolkit です。 2026 年初頭、エージェント LLM システムが ScanCode Toolkit を Python から Rust に移植し、ScanCode の商標を侵害する名前で派生結果を公開し、ScanCode と当社がベンダーとして慎重に帰属させたサードパーティ コードの両方から著作権とライセンス通知を剥奪し、AboutCode コミュニティに関与することなくアウトリーチ キャンペーンを開始しました。
この事件は孤立したものではありません。 AboutCode (および他の多くのオープンソース プロジェクト) では、AI によって生成された、一見もっともらしい問題やプル リクエストが着実に流入しています。

テンプレート化されており、既存のレポートを複製することが多く、実際のソフトウェアの使用に基づいていることはほとんどありません。オープンソース エコシステム全体のメンテナは、この AI スロップを「AI スロップ」と呼んでいます。これは人間によるトリアージ時間を浪費し、問題トラッカーの信号を低下させ、ユーザー、寄稿者、保守者の間の社会契約を侵食します。この投稿で説明した移植事件は、より大規模でより大きなリスクを伴う同じ現象です。
この記事では、技術的に何が起こったのか、AI 支援開発の現状について何が明らかになったのか、そして AI 生成コードを扱う際にオープンソース コミュニティが何をする必要があるのか​​を文書化します。
移植は、LLM オーケストレーション ハーネスによって推進されました (OpenCode と OpenClaw-vibe でコーディングされた OpenCode プラグインを使用)。エージェントのアプローチは単純明快で、十分にテストされた成熟した Python コードベースを取得し、Rust でリファクタリングしました。これは、独立して書き直されたものでも、ScanCode からインスピレーションを得たものでもありません。これは機械的な翻訳であり、まさに LLM が適した種類のタスクです。
なぜ？コード翻訳は基本的に言語翻訳タスクに似ており、大規模言語モデル (LLM) はもともとそのような言語タスクのために設計されました。広範な ScanCode テスト スイートにより、仕様とガイド レールが提供されました。エージェントはアルゴリズムを理解する必要はありません。テストに合格したコードを生成するだけで済みました。
これは繰り返す価値があります。包括的なテスト スイート、適切なドキュメント、厳選されたデータセットによって、自動移植が可能になります。これにより、コードベースを理解しなくても複製が容易になります。
既存の Rust ライセンス検出ライブラリを使用するエージェントの最初のアプローチは、ScanCode の出力品質と一致しませんでした。次に、エージェントは、緩やかな言い換えが失敗したときに翻訳者が行うことと同じことを行いました。つまり、元の m をコピーしました。

鉱石を近くに置きます。最終的なポートは、ScanCode のコア アルゴリズム、コード構成、データ駆動型アーキテクチャを Rust で再現します。これは、エージェントがそれらを理解したからではなく、同等のコードに収束するのに十分なトレーニング データとテスト フィードバックがあったからです。
Rust ポートでは、パフォーマンスが 10 倍から 100 倍向上したと主張する「ベンチマーク」が公開されました。多くのベンチマークは、そのツールの販売や宣伝に役立てるために、独自のツールの機能やパフォーマンスの優位性を文書化して主張するように設計されているため、根本的に欠陥があります。
コンパイルされた Rust は、解釈された Python よりも優れたパフォーマンスを発揮します。公開された「ベンチマーク」では、Rust ポートは ScanCode よりも高速に実行されますが、チェックすると誤った結果が返され、検出が欠落し、ファイルがスキップされます。 ScanCode は、Rust ポートがカバーするテストの数は少ないにもかかわらず、標準の ScanCode テスト スイートを Rust ポートよりも高速に実行します。 Rust 移植と同様の最適化を適用した後、ScanCode は正確性と帰属を維持しながら、Rust 移植と同じかそれよりも高速に実行されます。
サブセットの正確性や速度をテストすることは、全体としての優位性と同等ではありません。
これは、AI 支援ソフトウェア開発の中核的な問題も示しています。エージェントは、ScanCode の構造をいくつかのテストに合格するには十分に複製しましたが、すべてのテストに合格するには十分ではありませんでした。このポートはパフォーマンスの最適化とキャッシュ戦略を適用して高速に表示されましたが、重要なデータの正確性と完全性が犠牲になりました。
ライセンスと著作権の失敗
ScanCode は Apache-2.0 ライセンスを取得しています。 Apache オープン ソース ライセンスは、利用可能なライセンスの中で最も寛容なライセンスの 1 つであり、要件は最小限です。
元の NOTICE ファイルを保持します。
変更されたファイルも含めて、ライセンスおよび著作権ヘッダーを保持します。
変更されたファイルに加えられた変更に注意してください。
プロジェクト名を無断で再利用しないでください。

この港は 4 つの要件すべてに違反していました。要件 1 と 4 は、ScanCode メンテナーの連絡を受けて部分的に修正されました。要件 2 と 3 はそうではありませんでした。
これは、ScanCode 自体、その作成者、寄稿者以外にも大きな影響を与えます。 ScanCode には、他の多数のオープン ソース プロジェクトのコードが組み込まれており、それぞれに独自のライセンスと著作権があります。私たちは、元のファイル、ファイルごとの著作権ヘッダー、および帰属通知を使用して、これらすべてを細心の注意を払って追跡します。エージェントはそのすべてを剥奪し、コードが ScanCode を介してポートに渡されたすべての上流プロジェクトにライセンス違反を拡大しました。また、Apache ライセンスは段階的ではないことにも注意してください。準拠しているか、ライセンスを取得していないかのどちらかです。この出版物の時点では、このポートは準拠していません。
皮肉は微妙ではありません。 ScanCode は、コンプライアンス コミュニティの専門知識を結集して作成されたもので、業界がこの種のライセンスおよび著作権違反を正確に検出するために使用するツールです。
LLM は出所を追跡しません
最も重要な技術的観察は、速度や正確さに関するものではありません。それは帰属に関するものです。
LLM は設計上、出所を追跡しません。エージェントがコードを変換すると、出力が生成されます。出力が、特定のライセンスに基づいて、特定の寄稿者によって作成された特定のファイルから派生したものであることは記録されません。そのメタデータはモデルの出力表現の一部ではありません。
これは構造的な問題であり、構成の問題ではありません。オープンソース プロジェクトからコピーするエージェントは、ライセンス ヘッダーを検出して保存し、コードの発行元とライセンスを注意深く追跡するための明示的な後処理手順が追加されない限り、デフォルトで帰属を剥奪します。ここではそのような措置は講じられませんでした。その結果、現在行われている LLM 支援移植は、帰属層のない盗用パイプラインになります。
この難読化は、

常に受動的。 Rust ポートのコミット履歴と構造を検討すると、エージェントが直接またはプロンプトを通じて出力をソースから遠ざけるために積極的に働いたという証拠があります。生成されたコードと問題追跡ツールで見つかった証拠に基づいて、変数名が変更され、コメントが書き換えられるか削除され、移植されたコード行への追加の参照が追加され、「ScanCode に触発されただけの独立した書き換え」という主張が最初からプロジェクトの枠組みに組み込まれました。
オリジナリティを求めてもオリジナリティは生まれません。エージェントは指示に従っていました。エージェントに「オリジナルの実装」を作成するよう促すと、その下のコードは元のプロジェクトから派生したままでありながら、可能な限り表面レベルのバリエーションが生成されます。オリジナリティがあるように見えますが、検出するのが難しいため、単純なコピーよりも悪い結果となります。
日常の AI 支援開発においても、より小規模ながら同じことが起こります。開発者がコード生成ツールを使用してユーティリティ関数、パーサー、またはデータ構造を生成すると、生成されたコードは、モデルのトレーニング データに存在するオープン ソース コードの実装パターンを、その系統を示すことなく厳密に再現する場合があります。ほとんどの開発者はこれを確認する方法を知りません。ほとんどのツールはフラグを立てません。
これは、AI 支援開発の議論の双方に対する警告です。オープンソース開発者の場合、ライセンスと貢献者のクレジットはエージェントには見えません。 AI によって生成されたコードを作成する開発者にとって、ツールが生成する出力には、著作物が帰属表示なしに使用された作者に対して未解決の義務が課せられる可能性があります。
これはケーススタディであり、個別の事件ではありません
このエピソードは主に 1 つのプロジェクトや、

俳優一組。これは、オープンソース エコシステム全体で繰り返されるパターンのプレビューです。
ScanCode がターゲットとなった具体的な条件は、最も成功したオープンソース プロジェクトを特徴付ける条件と同じです。つまり、成熟したコードベース、包括的なテスト、大量のドキュメント、厳選された多数のコンテンツ、大規模な下流ユーザー ベース、活発なコミュニティ、よく知られ信頼できる名前です。使用されるツールとテクニックは、AI によって生成されたコミット、コントリビューション、リライト (エージェントのオーケストレーション、自動化された問題のクロール、ターゲットを絞ったコミュニティへの働きかけなど) によって日常的になりつつあります。
この事件の人的および社会的側面は、技術的な側面と同じくらい、あるいはそれ以上に重要です。エージェントは、ScanCode の問題トラッカーをクロールし、Fedora が 10 年前に廃止し、リポジトリが 2026 年 3 月にアーカイブされたツールである yum データベース サポートの 3 年前の機能リクエストなど、古く、時代遅れ、または正しくない機能を実装しました。エージェントは新機能の開発についても報告しましたが、これらの機能はすでに他の AboutCode オープン ソース プロジェクトに存在します。
これは、コミュニティのコンテキストを持たない自動開発が生み出すものです。社会的および戦略的に一貫性のない技術的に機能する作業は、ほとんど役に立たない、または冗長な技術的負債を生み出し、どの機能を実装するかを選択するために必要なエコシステムドメインの専門知識と集合的な知恵をバイパスします。
これは、大規模な AI のスロップに伴うコストのあまり議論されていないものの 1 つです。これは単なるノイズではなく、双方の実際のリソースを消費する誤った方向の努力です。メンテナは、低品質の問題の優先順位付けと解決に時間を費やします。自動化システムは、古い機能や無関係な機能の実装にコンピューティングを費やします。どちらも価値を生み出しません。そして、雑然とした問題トラッカー、未発見のライセンスで蓄積された技術的負債

違反があり、複製されたものの誤解されたコードは、人間のメンテナがクリーンアップする必要があります。
Rust 移植チームがユーザーに連絡して ScanCode の置き換えを提案するコミュニティ支援キャンペーンも、同様にコミュニティの理解の欠如を反映しています。 Rust 移植開発者は、そのキャンペーンが始まるまで、ScanCode のパブリック コミュニティ チャネル、毎週のミーティング、またはメンテナーとのチャットに参加することはありませんでした。導入を最適化する自動化システムは、オープンソース コミュニティが構築される信頼関係や協力的な規範を自然にモデル化するものではありません。
これから進むべき道は、この一件を訴訟することではない。今後はベスト プラクティスを開発することが重要です。
正当な貢献者を導き、誇張された主張に対する真実を提供するために、ベンチマーク スイートと明確なパフォーマンス プロファイルがこれまで以上に重要になります。 ScanCode などのツールを含むライセンス コンプライアンス ツールは、AI によって生成された投稿に定期的に適用する必要があります。アトリビューションのギャップは必ずしも意図的なものではありません。多くの場合、明示的に確認しないと表示されません。そして私たちは、オープンソースの作者がその仕事に対して適切にクレジットされることを保証するために、さらに多くのオープンソース ツールを構築しています。
オープンソースのメンテナは、統合を管理し、保護する必要があります。

[切り捨てられた]

## Original Extract

A case study, not an isolated incident

Skip to main content Getting Started Blog Projects Environments About Contribute
An AI agent ported our codebase from Python to Rust
A three-year roadmap for ClearlyDefined
OSI and AboutCode partner to sustain ClearlyDefined
ScanCode LicenseDB -- 2,000+ licenses curated in a public database
VulnerableCode API Deprecation and V3 Introduction
Non-Vulnerable Dependency Resolution
What is a Dual License Anyway?
There and back again -- A software versioning story
License Clarity Scoring in ScanCode
Using Copyleft-licensed software components in a Java application
An AI agent ported our codebase from Python to Rust
A case study, not an isolated incident ​
ScanCode detects licenses, copyrights, package dependencies, vulnerabilities, and a few more things in both source code and binary files. The use cases include license and security compliance and software supply chain management. It is the product of over a decade of careful design, architecture, and testing by an open source community of over 700 contributors, supporting more than 40,000 automated tests covering license detection alone, and over 90,000 automated tests overall.
The core module is ScanCode Toolkit, the industry-leading open source code scanning engine. In early 2026, an agentic LLM system ported ScanCode Toolkit, from Python to Rust, published the derived results under a name that infringed the ScanCode trademark, stripped copyright and license notices from both ScanCode and third-party code we vendored and carefully attributed, and started an outreach campaign, without ever engaging the AboutCode community.
This incident is not isolated. AboutCode (and many other open source projects) are experiencing a steady influx of AI-generated issues and pull requests that are superficially plausible, templated, often duplicating existing reports, and almost never grounded in actual use of the software. Maintainers across the open source ecosystem call this AI slop. It consumes human triage time, degrades signal in issue trackers, and erodes the social contract between users, contributors, and maintainers. The porting incident described in this post is the same phenomenon at a larger scale and with higher stakes.
This article documents what happened technically, what it reveals about the current state of AI-assisted development, and what the open source community needs to do when dealing with AI-generated code.
The porting was driven by an LLM orchestration harness (using OpenCode and an OpenClaw-vibe coded OpenCode plugin). The agent's approach was straightforward: take a mature, well-tested Python codebase and refactor it in Rust. This is not an independent rewrite or inspired by ScanCode as it claims. It is a mechanical translation and it is exactly the kind of task LLMs are well-suited for.
Why? Code translation is fundamentally like a language translation task, and Large Language Models (LLMs) were originally designed for such language tasks. The extensive ScanCode test suite provided the specification and the guide rails. The agent did not need to understand the algorithms; it only needed to produce code that passed the tests.
This is worth repeating: A comprehensive test suite, decent documentation, and curated datasets is what makes automated porting possible. It is also what makes a codebase easier to replicate without understanding it.
The agent's initial approach, using an existing Rust license-detection library, failed to match ScanCode's output quality. The agent then did what any translator would do when a loose paraphrase fails: it copied the original more closely. The final port reproduces ScanCode's core algorithms, code organization, and data-driven architecture in Rust, not because the agent understood them, but because it had enough training data and test feedback to converge on equivalent code.
The Rust port published a "benchmark" that claimed 10x to 100x improvements in performance. Many benchmarks are fundamentally flawed because they are designed to document and assert their own tool's feature or performance superiority to help sell or promote that tool.
Compiled Rust is capable of outperforming interpreted Python. In the published "benchmarks", the Rust port runs faster than ScanCode, but when checked it returns incorrect results, missing detections and skipping files. ScanCode runs the standard ScanCode test suite faster than the Rust port, even though the Rust port covers fewer tests. After applying optimization similar to what the Rust port did, ScanCode runs as fast or faster than the Rust port, while maintaining correctness, and attribution.
Testing correctness or speed on a subset does not equate with superiority on the whole.
This also demonstrates a core problem of AI-assisted software development. The agents replicated ScanCode's structure well enough to pass some tests, but not well enough to pass all tests. The port applied performance optimizations and caching strategies to appear faster, but sacrificing critical data correctness and completeness.
License and copyright failures ​
ScanCode is Apache-2.0 licensed. The Apache open source license is among the most permissive available, with minimal requirements:
Retain the original NOTICE file.
Preserve license and copyright headers, including in modified files.
Note changes made to modified files.
Do not reuse the project name without permission.
The port violated all four requirements. Requirements 1 and 4 were partially corrected after ScanCode maintainers reached out. Requirements 2 and 3 were not.
This impacts more than ScanCode itself and its authors and contributors. ScanCode incorporates code from dozens of other open source projects, each with its own license and copyright. We track all of this meticulously with origin files, per-file copyright headers, and attribution notices. The agent stripped all of it, extending the license violations to every upstream project whose code passed through ScanCode into the port. Note also that the Apache license is not graduated: you either comply or you are not licensed. As of this publication, the port is not compliant.
The irony is not subtle. ScanCode is the product of the collective expertise of the compliance community, and is a tool that the industry uses to detect exactly this kind of license and copyright violation.
LLMs do not track provenance ​
The most important technical observation is not about speed or correctness. It is about attribution.
LLMs, by design, do not track provenance. When an agent translates code, it produces output. It does not record that the output derives from a specific file, authored by a specific contributor, under a specific license. That metadata is not part of the model's output representation.
This is a structural problem, not a configuration issue. Agents copying from open source projects will strip attribution by default unless explicit post-processing steps are added to detect and preserve license headers and carefully track the code origin and license. No such steps were taken here. The result is that LLM-assisted porting, as currently practiced, is a plagiarism pipeline with no attribution layer.
This obfuscation is not always passive. In reviewing the commit history and structure of the Rust port, there is evidence that the agent actively worked to distance the output from its source, either directly or steered through prompting. Variable names were changed, comments were rewritten or stripped, additional references to ported code lines added, and the claim of an "independent rewrite merely inspired by ScanCode" was baked into the project's framing from the start, based on evidence found in the generated code and the issue tracker.
Prompting for originality does not produce originality. The agent was following instructions. If you prompt an agent to produce an "original implementation", it will generate whatever surface-level variation possible while the code underneath remains derived from the original project. It produces the appearance of originality, which is a worse outcome than straightforward copying because it is harder to detect.
The same dynamic occurs at a smaller scale in everyday AI-assisted development. When a developer uses a code generation tool to produce a utility function, a parser, or a data structure, the generated code may closely reproduce implementation patterns from open source code present in the model's training data, without any indication of that lineage. Most developers do not know to check for this. Most tools do not flag it.
This is a warning to both sides of the AI-assisted development discussion. For open source developers, your licenses and your contributors' credits are invisible to the agent. For developers producing AI-generated code, the output your tools produce may carry unresolved obligations to authors whose work was used without attribution.
This is a case study, not an isolated incident ​
This episode is not primarily about one project or one set of actors. It is a preview of a pattern that will repeat across the open source ecosystem.
The specific conditions that made ScanCode a target are the same conditions that characterize most successful open source projects: a mature codebase, comprehensive tests, plenty of documentation, lots of curated content, large downstream user base, an active community, and a well-known and trusted name. The tools and techniques used are becoming routine with AI-generated commits, contributions, and rewrites: agentic orchestration, automated issue crawling, and targeted community outreach.
The human and social dimensions of this incident are as important, if not more important, as the technical ones. The agent crawled ScanCode's issue tracker and implemented old, outdated or incorrect features, such as a three-year-old feature request for yum database support, a tool that Fedora deprecated a decade ago and whose repository was archived in March 2026. The agent also reported the development of new features, but these features already exist in other AboutCode open source projects.
This is what automated development without community context produces: technically functional work that is socially and strategically incoherent, creating mostly useless or redundant technical debt and bypassing the ecosystem domain expertise and collective wisdom needed to select which feature to implement.
This is one of the less-discussed costs of AI slop at scale. It is not just noise, it is misdirected effort that consumes real resources on both sides. Maintainers spend time triaging and closing low-quality issues. Automated systems spend compute implementing stale or irrelevant features. Neither produces value. And the accumulated technical debt in cluttered issue trackers, undiscovered license violations, and replicated but misunderstood code falls on human maintainers to clean up.
The community outreach campaign by the Rust port team contacting users to suggest replacing ScanCode reflects the same absence of community understanding. The Rust port developers never engaged ScanCode's public community channels, weekly meetings, or chatting with maintainers, until that campaign began. An automated system optimizing for adoption does not naturally model the trust relationships and collaborative norms that open source communities are built on.
The path forward is not to litigate this one case. The path forward is to develop best practices.
Benchmark suites and clear performance profiles matter more than ever, both to guide legitimate contributors and to provide ground truth against inflated claims. License compliance tooling, including tools like ScanCode, should be routinely applied to AI-generated contributions. Attribution gaps are not always intentional; they are often invisible without explicitly checking. And we are building more open source tools to help ensure open source authors are properly credited for their work.
To open source maintainers, you should care for and protect the integr

[truncated]
