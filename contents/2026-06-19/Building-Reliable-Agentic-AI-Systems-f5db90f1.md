---
source: "https://martinfowler.com/articles/reliable-llm-bayer.html"
hn_url: "https://news.ycombinator.com/item?id=48595716"
title: "Building Reliable Agentic AI Systems"
article_title: "Building Reliable Agentic AI Systems"
author: "Garbage"
captured_at: "2026-06-19T07:59:31Z"
capture_tool: "hn-digest"
hn_id: 48595716
score: 1
comments: 0
posted_at: "2026-06-19T07:12:53Z"
tags:
  - hacker-news
  - translated
---

# Building Reliable Agentic AI Systems

- HN: [48595716](https://news.ycombinator.com/item?id=48595716)
- Source: [martinfowler.com](https://martinfowler.com/articles/reliable-llm-bayer.html)
- Score: 1
- Comments: 0
- Posted: 2026-06-19T07:12:53Z

## Translation

タイトル: 信頼性の高いエージェント AI システムの構築

記事本文:
課題: 前臨床データの迷路をナビゲートする
ソリューション: PRINCE - 進化したプラットフォーム
システム アーキテクチャ: 信頼性の高いエージェント RAG システムのエンジニアリング
Agentic RAG システム
ユーザーの意図を明確にする
思考と計画: プロセスの反映
リフレクション エージェント: データの検証と十分性
Writer エージェント: 回答の合成とフォーマット
本番LLMシステムにおける信頼の構築
透明性と説明可能性
回復力のためのエンジニアリング: エラー処理と回復
データ品質の強化: 固有表現の認識と注釈
旅は続く: 反復開発
信頼性の高いエージェント AI システムの構築
本番環境に対応したエージェント AI システムの構築におけるケーススタディ
このペーパーでは、クラウドホスト型プラットフォームである前臨床情報センター (PRINCE) について説明します。
医薬品業界の課題に対処するために、バイエル AG が Thoughtworks と共同で開発した
開発。 PRINCE はエージェント検索拡張生成を活用します
Text-to-SQL により、数十年にわたる安全性研究レポートが統合されます。 PRINCEの進化を解説します
キーワードベースの検索から、複雑な質問に答えることができるインテリジェントなリサーチアシスタントまで
質問や規制文書の作成など。私たちはエンジニアリング上の重要な決定を次のように反映しています。
コンテキスト エンジニアリングのレンズ — 情報がどのように形成され、専門化されたデータ間でルーティングされるか
エージェントとハーネス エンジニアリング - オーケストレーション、リカバリ、可観測性がどのように構築されたか
制御と信頼性を維持するためにモデルの周囲に配置されます。システムは信頼を優先します。
透明性、説明可能性、人間参加型の統合。 PRINCE が AI のデモンストレーションを行う
医薬品の変革の可能性を高め、データへのアクセス性を大幅に向上させ、
ガバナンスとコンプライアンスを確保しながら研究を効率化します。
Sarang Kulkarni は Thoughtwor の主席コンサルタントです

交差点で働いているksさん
ソフトウェアエンジニアリング、データプラットフォーム、応用AI。彼は建築に重点を置いています
実稼働グレードの GenAI システム、特に検索拡張生成 (RAG) と
マルチエージェント ワークフローをサポートし、チームがこれらのシステムを初期のアイデアから現実世界に移行するのを支援します
使用します。サラン氏は、Thoughtworks のグローバル AI サービス開発チームにも貢献し、教育を行っています。
オライリー
本番環境に対応した RAG アプリケーションの構築に関するコース。
課題: 前臨床データの迷路をナビゲートする
ソリューション: PRINCE - 進化したプラットフォーム
システム アーキテクチャ: 信頼性の高いエージェント RAG システムのエンジニアリング
Agentic RAG システム
ユーザーの意図を明確にする
思考と計画: プロセスの反映
リフレクション エージェント: データの検証と十分性
Writer エージェント: 回答の合成とフォーマット
本番LLMシステムにおける信頼の構築
透明性と説明可能性
回復力のためのエンジニアリング: エラー処理と回復
データ品質の強化: 固有表現の認識と注釈
旅は続く: 反復開発
前臨床創薬は本質的に複雑でデータ集約的です。
研究者は、効率的にアクセスし、
この重要な段階で生成された膨大な量の情報を分析します。
従来のキーワードベースの検索方法は、多くの場合、厳密なブール値に依存します。
論理的であり、微妙で複雑な問題に直面すると、しばしば不十分になります。
前臨床研究の質問の性質。
大規模言語モデル (LLM) の出現により、変革の機会がもたらされました。によって
LLM の生成能力と情報検索システムの精度を組み合わせた検索拡張生成 (RAG) が、有望な技術として浮上しています。
このアプローチは、前臨床データ アクセスに革命をもたらす可能性を秘めています。
研究者からposへ

複雑な質問を自然言語で送信し、正確でコンテキストに富んだ情報を受け取ります
独自のデータに基づいた回答を提供します。
バイエルはこの可能性を早期に認識し、これらがどのようにして実現されるかを探ることに尽力しました。
これらの技術は、前臨床研究における長年の課題に対処できる可能性があります。
この投稿では、バイエルが生成 AI に初期投資した経緯について説明します。
その結果、Agentic RAG 上に構築されたエージェント AI システムである PRINCE が誕生しました。このケーススタディ
技術的なアーキテクチャ、エンジニアリング上の決定、教訓を探ります
困難な迷路から前臨床データ検索を変革する中で学んだこと
直感的な会話体験を実現します。
PRINCE の背後にあるエンジニアリング上の決定の多くは、コンテキストのレンズを通して理解できるようになりました。
エンジニアリングとハーネス エンジニアリングですが、システムが最初に設計されたときはこれらの用語を使用していませんでした。コンテキストエンジニアリングが各モデルにどのような情報を形づくるか
受信したもの、受信しなかったもの、および特殊なステップ間でコンテキストがどのように移動したか
研究、考察、執筆。ハーネス工学が周囲の足場を形作った
モデル: オーケストレーション、ツール境界、状態永続性、再試行、フォールバック、検証、
リフレクション ループ、可観測性、人間によるレビュー。
この投稿では技術的なアーキテクチャとエンジニアリングの課題に焦点を当てていますが、私たちの論文では
Frontiers in Artificial Intelligence に掲載された記事では、
製品の進化とビジネスへの影響について詳しく説明します。
課題: 前臨床データの迷路をナビゲートする
多くの大規模な企業と同様、バイエルの前臨床研究の状況
製薬組織の特徴は、多様性と広範さです。
データの配列。これには、さまざまな研究からの高度に構造化されたデータセットと、膨大なデータセットが含まれます。
非構造化の量
研究レポートなどのテキスト文書に埋め込まれた情報、
出版物、および

規制当局への提出。研究者は頻繁に
これにアクセスして分析する際に大きなハードルに遭遇しました
情報を効果的に:
データサイロ: 情報は断片化され、多数のデータに分散されていました。
異種のシステムとリポジトリがあるため、
特定の化合物に関連する前臨床データの包括的かつ全体的なビュー
または勉強します。
限定的な検索機能: 従来のキーワードベースの検索エンジン
前臨床用語の複雑さと多様性に苦労し、
研究上の質問、多くの場合、無関係、不完全、または圧倒的な結果が得られます
結果。
時間のかかる手動分析: 特定の洞察の抽出または編集
複数の文書にまたがる情報にはかなりの手作業が必要でした。
研究者の貴重な時間を中核的な科学活動から逸らすことになります。
これらの固有の課題は、さらなる改善の必要性を明らかにしました。
前臨床データに対する効率的かつインテリジェントな統合アプローチ
検索と分析。
ソリューション: PRINCE - 進化したプラットフォーム
これらの課題に対処するために、バイエルは前臨床試験を開発しました。
インフォメーション センター (PRINCE) プラットフォーム。 PRINCEは統一されたものとして考案されました
前臨床データへのゲートウェイ、最初は統合に重点を置く
以前は構造化された研究メタデータをサイロ化し、それらを「検索可能な」方法で公開しました。
この初期フェーズでは、ユーザーが高度なフィルターを適用して取得できるようになりました。
主に構造化された研究メタデータからの情報。
ただし、バイエルの貴重な前臨床試験の重要な部分は、
知識は、長年にわたって蓄積された非構造化 PDF 調査レポートに存在します。
数十年。長年にわたる多数のシステム移行により、構造化された
これらのレポートに関連付けられたメタデータは、不完全、欠落、または
間違った注釈も含まれています。重要なのは、権威ある「金」
s

標準」情報が承認された PDF 内に一貫して存在していました
研究報告書。
Generative AI、特に RAG の出現は、
この豊富な非構造化データを解放します。 RAGを統合することで
PRINCE はフィルターベースからパラダイムをシフトし始めました。
「検索」ツールを自然言語の「質問」システムに組み込むことで、研究者は次のことが可能になります。
これらの研究レポートの内容を直接問い合わせてください。
この進化は、PRINCE の 3 つの異なる進化を反映しています。
フェーズ:
検索: 最初のフェーズでは、統合ゲートウェイの作成に焦点を当てました。
数千件の非臨床研究レポート、複数の社内データサイロを統合
さまざまな前臨床ドメインを
検索可能な形式。主に構造化メタデータを活用します。
Ask: このフェーズでは、AI を活用した質問応答システムを導入しました。
検索拡張生成 (RAG)。これにより、研究者は直接洞察を得ることが可能になりました
ポーズによる履歴レポートのスキャン PDF などの非構造化データから
自然言語での質問。
Do: 現在のフェーズでは、PRINCE を次のことができるアクティブな研究助手として位置付けています。
複雑なタスクの実行。これは、マルチエージェント システムの統合によって実現されます。
プラットフォームが複雑なクエリを処理し、ワークフローを調整し、サポートできるようにします。
規制文書の草案作成などの活動。
検索から依頼へのこの意図的な進化は、戦略的アプローチを表しています。
さらなる効率性と革新性を求める業界のニーズに応える
前臨床開発。研究者にますます強力な情報を提供することで、
PRINCE は、前臨床データにアクセスし、分析し、それに基づいて行動するためのツールを提供することを目指しています。
データに基づいた意思決定の迅速化、不必要な実験の必要性の削減、
そして最終的には、より安全でより効果的な製品の開発を加速します。
治療法。
システムアーチ

構造: 信頼性の高い Agentic RAG システムの設計
システムは、堅牢なバックエンドを搭載した対話型の会話型 UI として機能します。
インフラストラクチャ。そのアーキテクチャは、複雑なクエリを処理し、
正確でコンテキスト豊富な回答は、LangGraph を使用して調整され、
FastAPI アプリケーション。
図 1 は、システム コンテキスト (UI、バックエンド、データ) を示しています。
ストア、LLM フォールバック、オブザーバビリティ - 図 2
システムがその専門エージェントをどのように調整するかに焦点を当てます。
図 1: システムのコンテキストとサポート
プラットフォーム。
ユーザーリクエスト: ユーザーがリクエストを送信すると、プロセスが開始されます。
React で構築された会話型 UI。
オーケストレーション: ユーザーのリクエストは、LangGraph ベースのオーケストレーション レイヤーにルーティングされます。
バックエンド。このワークフロー エンジンは、進行する複数段階のプロセスを調整します。
を通して
ユーザーの意図の明確化、思考と計画、調査の実施（RAG と
テキストから SQL)、
データの完了を検証し、最後に Writer エージェントを通じて応答を生成します。
の
ワークフローには、データの完全性を保証するための意図的な一時停止ポイントとフィードバック ループが含まれています
前に
進んでいます。 (このエージェント ワークフローの詳細については、専用のセクションで説明します)
後で。)
データの取得と状態管理: Researcher エージェントは包括的なデータと対話します。
そして
分散データエコシステム:
すべての研究レポートのベクトル表現は OpenSearch に保存され、
情報検索の中核となる知識ベース。
さまざまな ETL と調和から得られる、厳選された構造化データ
プロセスには Athena 経由でアクセスします。
エージェントの実行状態は注意深く追跡されます。各論理の後で
ステップ (LangGraph ノードの実行)、対応する状態は永続化されます。
LangGraph チェックポイントを使用する PostgreSQL。
より幅広い用途

カチオンレベルの状態は次のように管理されます。
DynamoDB 。
このシステムは、OpenAI、Anthropic、
Google とオープンソース プロバイダー。これらのプラットフォームは、統合されたプラットフォームを通じてすべてのモデルを公開します。
OpenAI 互換エンドポイントにより、モデルの交換や最適なツールの選択が容易になります。
それぞれのタスク。また、コントロール プレーンも管理し、レート制限やその他の安全策を適用します。
虐待を防ぐため。
復元力とエラー処理: 堅牢性は重要な設計原則であり、
複数のフォールバック メカニズムを導入:
特定の LLM が失敗した場合、システムは自動的に再試行します。
代替モデルまたはプラットフォームにフォールバックする前に、リクエストを数回繰り返します。
サービスの継続性を確保します。
一時的な障害から迅速に回復するには、再試行が必要です。
個々の LLM 呼び出しレベルと論理ノード レベルの両方で実装されます (つまり、
エージェントの計画の全ステップ)。
また、エージェントにはエラーのコンテキストが提供されるため、別のグラフを作成できます。
対応としての行動の軌道または代替計画。
可観測性と評価: システム全体のパフォーマンスと評価が監視されます。
信頼性:
一般的なシステムの健全性とメトリクスは、 Cloudwatch を使用して追跡されます。
Langfuse は主要な可観測性ツールとして機能し、

[切り捨てられた]

## Original Extract

The Challenge: Navigating the Preclinical Data Maze
The Solution: PRINCE - An Evolutionary Platform
System Architecture: Engineering a Reliable Agentic RAG System
The Agentic RAG System
Clarify User Intent
Think & Plan: Process Reflection
The Reflection Agent: Data Validation and Sufficiency
The Writer Agent: Answer Synthesis and Formatting
Building Trust in a Production LLM System
Transparency and Explainability
Engineering for Resilience: Error Handling and Recovery
Enhancing Data Quality: Named Entity Recognition and Annotation
The Journey Continues: Iterative Development
Building Reliable Agentic AI Systems
A Case Study in building production-ready agentic AI systems
This paper presents the Preclinical Information Center (PRINCE), a cloud-hosted platform
developed by Bayer AG with Thoughtworks to address pharmaceutical industry challenges in drug
development. PRINCE leverages Agentic Retrieval-Augmented Generation
and Text-to-SQL to integrate decades of safety study reports. We describe PRINCE's evolution
from keyword-based search to an intelligent research assistant capable of answering complex
questions and drafting regulatory documents. We reflect on key engineering decisions through
the lens of context engineering—how information was shaped and routed between specialized
agents—and harness engineering—how orchestration, recovery, and observability were built
around the models to maintain control and reliability. The system prioritizes trust through
transparency, explainability, and human-in-the-loop integration. PRINCE demonstrates AI's
transformative potential in pharmaceuticals, significantly improving data accessibility and
research efficiency while ensuring governance and compliance.
Sarang Kulkarni is a Principal Consultant at Thoughtworks, working at the intersection of
software engineering, data platforms, and applied AI. He focuses on building
production-grade GenAI systems, particularly Retrieval-Augmented Generation (RAG) and
multi-agent workflows, and helps teams take these systems from early ideas to real-world
use. Sarang also contributes to Thoughtworks’ Global AI Service Development team and teaches
an O’Reilly
course on building production-ready RAG applications.
The Challenge: Navigating the Preclinical Data Maze
The Solution: PRINCE - An Evolutionary Platform
System Architecture: Engineering a Reliable Agentic RAG System
The Agentic RAG System
Clarify User Intent
Think & Plan: Process Reflection
The Reflection Agent: Data Validation and Sufficiency
The Writer Agent: Answer Synthesis and Formatting
Building Trust in a Production LLM System
Transparency and Explainability
Engineering for Resilience: Error Handling and Recovery
Enhancing Data Quality: Named Entity Recognition and Annotation
The Journey Continues: Iterative Development
Preclinical drug discovery is inherently complex and data-intensive.
Researchers face the significant challenge of efficiently accessing and
analyzing vast volumes of information generated during this critical phase.
Traditional keyword-based search methods, often reliant on rigid Boolean
logic, frequently fall short when confronted with the nuanced and intricate
nature of preclinical research questions.
The advent of Large Language Models (LLMs) has presented a transformative opportunity. By
combining the generative power of LLMs with the precision of information retrieval systems, Retrieval-Augmented Generation (RAG) has emerged as a promising technique.
This approach holds the potential to revolutionize preclinical data access, enabling
researchers to pose complex questions in natural language and receive accurate, context-rich
answers grounded in proprietary data.
Recognizing this potential early, Bayer committed to exploring how these
technologies could address longstanding challenges in preclinical research.
In this post, we share that journey—how Bayer's early investment in generative AI
has resulted in PRINCE, an agentic AI system built on Agentic RAG. This case study
explores the technical architecture, engineering decisions, and lessons
learned in transforming preclinical data retrieval from a challenging maze
into an intuitive conversational experience.
Many of the engineering decisions behind PRINCE can now be understood through the lens of context
engineering and harness engineering, although when the system was first designed we did not use these terms. Context engineering shaped what information each model
received, what it did not receive, and how context moved between specialized steps such as
research, reflection, and writing. Harness engineering shaped the scaffolding around the
models: orchestration, tool boundaries, state persistence, retries, fallbacks, validation,
reflection loops, observability, and human review.
While this post focuses on the technical architecture and engineering challenges, our paper
published in Frontiers in Artificial Intelligence covers the
product evolution and business impact in more detail.
The Challenge: Navigating the Preclinical Data Maze
The preclinical research landscape at Bayer, like many large
pharmaceutical organizations, is characterized by a diverse and extensive
array of data. This includes highly structured datasets from various studies, alongside vast
amounts of unstructured
information embedded within text documents such as study reports,
publications, and regulatory submissions. Researchers frequently
encountered significant hurdles in accessing and analyzing this
information effectively:
Data Silos: information was fragmented and scattered across numerous
disparate systems and repositories, making it exceedingly difficult to gain a
comprehensive, holistic view of preclinical data related to a specific compound
or study.
Limited Search Capabilities: traditional keyword-based search engines
struggled with the complexity and variability of preclinical terminology and
research questions, often yielding irrelevant, incomplete, or overwhelming
results.
Time-Consuming Manual Analysis: extracting specific insights or compiling
information across multiple documents required considerable manual effort,
diverting valuable researcher time away from core scientific activities.
These inherent challenges highlighted a clear need for a more
efficient, intelligent, and integrated approach to preclinical data
retrieval and analysis.
The Solution: PRINCE - An Evolutionary Platform
To address these challenges, Bayer developed the Preclinical
Information Center (PRINCE) platform. PRINCE was conceived as a unified
gateway to preclinical data, initially focusing on consolidating
previously siloed structured study metadata and exposing them in a “Searchable” manner.
This initial phase allowed users to apply advanced filters and retrieve
information primarily from structured study metadata.
However, a significant portion of Bayer's valuable preclinical
knowledge resides within unstructured PDF study reports accumulated over
decades. Due to numerous system migrations over the years, the structured
metadata associated with these reports could be incomplete, missing, or
even contain incorrect annotations. Crucially, the authoritative “gold
standard” information was consistently present within the approved PDF
study reports.
The emergence of Generative AI, particularly RAG, provided the key to
unlocking this wealth of unstructured data. By integrating RAG
capabilities, PRINCE began to shift the paradigm from a filter-based
'search' tool to a natural language 'ask' system, enabling researchers to
query the content of these study reports directly.
This evolution reflects PRINCE's progression through three distinct
phases:
Search: the initial phase focused on creating a unified gateway to
thousands of nonclinical study reports, consolidating multiple in-house data silos from
various preclinical domains into a
searchable format, primarily leveraging structured metadata.
Ask: this phase introduced an AI-powered question-answering system utilizing
Retrieval Augmented Generation (RAG). This enabled researchers to derive insights directly
from unstructured data, including scanned PDFs from historical reports, by posing
questions in natural language.
Do: the current phase positions PRINCE as an active research assistant capable of
executing complex tasks. This is achieved through the integration of multi-agent systems,
allowing the platform to handle intricate queries, orchestrate workflows, and support
activities like drafting regulatory documents.
This deliberate evolution from Search to Ask to Do represents a strategic
response to the industry's need for greater efficiency and innovation in
preclinical development. By providing researchers with increasingly powerful
tools to access, analyze, and act upon preclinical data, PRINCE aims to enable
faster data-driven decision-making, reduce the need for unnecessary experiments,
and ultimately accelerate the development of safer, more effective
therapies.
System Architecture: Engineering a Reliable Agentic RAG System
The system functions as an interactive conversational UI, powered by a robust backend
infrastructure. Its architecture, designed for handling complex queries and delivering
accurate, context-rich answers, is orchestrated using LangGraph and served via a
FastAPI application.
Figure 1 provides the system context—UI, backend, data
stores, LLM fallbacks, and observability—while Figure 2
zooms into how the system coordinates its specialized agents.
Figure 1: System context and supporting
platforms.
User Request: the process begins when a user submits a request through the
Conversational UI which is built with React.
Orchestration: the user's request is routed to a LangGraph-based orchestration layer in
the backend. This workflow engine coordinates a multi-stage process that progresses
through
clarifying user intent, thinking and planning, conducting research (using RAG and
Text-to-SQL),
validating data completion, and finally generating a response through the Writer agent.
The
workflow includes deliberate pause points and feedback loops to ensure data completeness
before
proceeding. (We explore the details of this agentic workflow in a dedicated section
later.)
Data Retrieval and State Management: the Researcher agents interact with a comprehensive
and
distributed data ecosystem:
Vector representations of all study reports are stored in OpenSearch , forming
the core knowledge base for information retrieval.
Curated structured data , resulting from various ETL and harmonization
processes, is accessed via Athena .
The state of the agent's execution is meticulously tracked. After each logical
step (a LangGraph node execution), the corresponding state is persisted in
PostgreSQL using a LangGraph checkpointer .
Broader application-level state is managed in
DynamoDB .
The system leverages internal GenAI platforms that host models from OpenAI, Anthropic,
Google, and open-source providers. These platforms expose all models via a unified
OpenAI-compatible endpoint, making it easy to swap models and choose the best tool for
each task. They also manage the control plane, enforcing rate limits and other safeguards
to prevent abuse.
Resilience and Error Handling: robustness is a critical design principle, with
multiple fallback mechanisms in place:
If a specific LLM fails, the system automatically retries
the request several times before falling back to an alternative model or platform to
ensure service continuity.
To recover quickly from transient failures, retries are
implemented at both the individual LLM call level and the logical node level (i.e., an
entire step in the agent's plan).
Also, agents are provided the context of the errors so that they can chart a different
trajectory or alternative plan of action as a response.
Observability and Evaluation: the entire system is monitored for performance and
reliability:
General system health and metrics are tracked using Cloudwatch .
Langfuse serves as the primary observability tool, providing d

[truncated]
