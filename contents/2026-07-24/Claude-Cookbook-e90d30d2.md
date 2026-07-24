---
source: "https://platform.claude.com/cookbook/"
hn_url: "https://news.ycombinator.com/item?id=49031409"
title: "Claude Cookbook"
article_title: "Claude Cookbook"
author: "saikatsg"
captured_at: "2026-07-24T05:11:30Z"
capture_tool: "hn-digest"
hn_id: 49031409
score: 1
comments: 0
posted_at: "2026-07-24T05:09:15Z"
tags:
  - hacker-news
  - translated
---

# Claude Cookbook

- HN: [49031409](https://news.ycombinator.com/item?id=49031409)
- Source: [platform.claude.com](https://platform.claude.com/cookbook/)
- Score: 1
- Comments: 0
- Posted: 2026-07-24T05:09:15Z

## Translation

タイトル: クロード料理本
説明: Claude を使用して構築するための実践的なガイドとコード例。プロンプトテクニック、ツールの使用法、マルチモーダル機能などを学びます。

記事本文:
クロード料理本
料理本 クロードの料理本
クロードを効果的に使用するための実践的なガイドと例
プログラムによるツール呼び出し (PTC) コード実行環境でプログラムによってツールを呼び出すコードをクロードに作成させることで、待ち時間とトークンの消費を削減します。埋め込みによるツール検索 動的ツール検出のためのセマンティック埋め込みを使用して、Claude アプリケーションを数千のツールに拡張します。自動コンテキスト圧縮 会話履歴を自動的に圧縮することで、長時間実行されるエージェント ワークフローのコンテキスト制限を管理します。画像分析を改善するためのトリミング ツールをクロードに提供する クロードに、チャート、ドキュメント、図の詳細な分析のために画像領域にズームするためのトリミング ツールを提供します。フロントエンドの美学を求める 一般的な美学を避け、独特で洗練されたフロントエンド デザインをクロードに求めるためのガイド。クロード スキルの紹介 クロードの Excel、PowerPoint、PDF スキルを使用して、ドキュメントを作成し、データを分析し、ワークフローを自動化します。すべての料理本
すべてのカテゴリ タイトル カテゴリ 著者 日付 Messages API で Claude のエージェント検索ベンチマーク スコアを再現 2026 年 6 月 • Evals Tools プログラムによるツール呼び出し、サーバー側の圧縮、およびタスク バジェットを使用して、公開された DeepSearchQA および BrowseComp スコアを再現するメッセージ API ハーネスを構築します。
Fable 5 で安全性分類子ブロックを検出し、サーバー側または SDK ベースのクライアント側フォールバックを使用して Opus 4.8 にフォールバックします (ストリーミング動作や新しい請求の変更など)。
2 つの非同期マルチエージェント パターン (共有ハブを介したピア メッセージングを備えた固定 N エージェント チームと、動的に生成される非同期サブエージェント) は、必要最低限​​のメッセージングとライフサイクルの仕組みにまとめられています。
同じ内容で、ノートブック 00 から運用成熟度の 3 層 (Docker、Modal、Kubernetes) までリサーチ エージェントをデプロイします。

各層の ainer イメージと HTTP インターフェイス。
マルチエージェント コーディネーター構成による異種混合チーム — コーディネーターは、販売提案を組み立てるために、範囲指定されたツールセットを備えた 3 人の専門家 (Web 検索調査員、ファイル読み取り司書、ルールベースの価格設定担当者) を実行します。マルチエージェント フィールド、thread_created / thread_message_received イベント タイプ、およびロールごとのツールの範囲をカバーします。
Outcomes を使用して採点と改訂のループを構築します。ライターが引用研究要旨の下書きを作成し、ステートレスな採点者がすべての URL を取得してルーブリックと照合してすべての引用をチェックし、概要が合格するまでフィードバックによって改訂が行われます。 user.define_outcome、span.outcome_evaluation_* イベント、および採点者が実行できるルーブリックの書き方について説明します。
Claude 管理対象エージェントにメモリ ストアを与えて、複数のインタラクションにわたるユーザーの設定を学習して記憶できるようにします。
Claude Agent SDK を使用して脆弱性発見エージェントを構築します。このエージェントは、C ターゲットを脅威モデル化し、組み込みのファイル ツールでメモリ安全性のバグを追跡し、発見結果を優先順位付けして構造化されたレポートを作成します。
Claude をオンコール フローに接続します。アラートが発生すると、エージェントはログとランブックを読み取り、根本原因を特定し、修正 PR を開き、マージする前に承認を待ちます。
サンドボックス環境とファイル マウントを使用して、CSV をインタラクティブなグラフを備えた説明的な HTML レポートに変換するアナリストを構築します。
CSV でボットにメンションすると、スレッド内で分析レポートが取得され、同じセッションで複数ターンのフォローアップが行われます。
Claude 管理エージェント API のエントリポイント チュートリアル。エージェントに calc.py パッケージに埋め込まれた 3 つのバグを修正してもらい、エージェント/環境/セッションの作成、ファイルのマウント、ストリーミング イベント ループについて説明します。
管理対象エージェントのエンドツーエンドの運用ストーリー - ボールトにバックアップされた MCP 認証情報、session.status_idled Webhook パターン

または、存続期間の長い接続を使用しない人間参加型、およびリソースのライフサイクル CRUD 動詞。
サーバー側のプロンプトのバージョン管理 — v1 の作成、ラベル付きテスト セットに対する評価、v2 の出荷、回帰の検出、セッションをバージョン 1 に固定することによるロールバック。 Agents.update、sessions.create でのバージョンの固定、およびプロンプトがコードでない場合のレビュー ゲートの移動先について説明します。
複数の脅威情報ソースのクエリ、結果の相互参照、MITRE ATT&CK へのマッピング、SIEM と SOAR 統合のための構造化レポートの作成によって、IOC を自律的に調査するエージェントを構築します。
ディスク上の Agent SDK セッションをリスト、読み取り、名前変更、タグ付け、フォークして、トランスクリプト パーサーを作成せずに会話履歴サイドバーを構築します。
Claude を使用して、エンティティ抽出、関係マイニング、重複排除、およびマルチホップ グラフ クエリを実行して、非構造化テキストからナレッジ グラフを構築します。
長期にわたって実行されているエージェントのコンテキスト エンジニアリング戦略を比較し、それぞれがいつ適用されるか、コストがかかるか、構成方法を学びます。
OpenAI Agents SDK アプリを Claude Agent SDK に移植し、単一の経費承認エージェントの例を通じて各プリミティブ (ツール、ガードレール、セッション、ハンドオフ) をマッピングします。
自律的な診断、修復、事後文書化のための読み取り/書き込み MCP ツールを備えたインシデント対応エージェントを構築します。
バックグラウンド スレッドとプロンプト キャッシュを使用したインスタント セッション メモリ圧縮により、長時間にわたるクロードの会話を管理します。
コード実行環境でプログラム的にツールを呼び出すコードをクロードに作成させることで、レイテンシとトークンの消費を削減します。
動的ツール検出のためのセマンティック埋め込みを使用して、Claude アプリケーションを何千ものツールに拡張します。
会話履歴を自動的に圧縮することで、長時間実行されるエージェント ワークフローのコンテキスト制限を管理します。
Ele を使用して低遅延の音声アシスタントを構築する

音声合成およびテキスト読み上げと Claude を組み合わせた venLabs。
クロードに、チャート、文書、図表の詳細な分析のために画像領域をズームインするトリミング ツールを提供します。
一般的な美しさを避け、独特で洗練されたフロントエンド デザインをクロードに促すためのガイド。
Claude の Excel、PowerPoint、PDF スキルを使用して、財務ダッシュボードとポートフォリオ分析を構築します。
特殊な組織ワークフローでクロードを拡張するカスタム スキルを作成、展開、管理します。
Claude の Excel、PowerPoint、PDF スキルを使用して、ドキュメントを作成し、データを分析し、ワークフローを自動化します。
自律調査用の WebSearch を備えた Claude Code SDK を使用して調査エージェントを構築します。
サブエージェント、フック、出力スタイル、プラン モード機能を備えたマルチエージェント システムを構築します。
GitHub モニタリングと CI ワークフローのために MCP サーバー経由でエージェントを外部システムに接続します。
評価タスク ファイルとは独立して、ツール上でエージェント評価を並行して実行します。
Admin API を介してプログラムで Claude API の使用状況とコスト データにアクセスし、分析します。
Claude のメモリ ツールとコンテキスト編集を使用して、永続メモリを備えた AI エージェントを構築します。
ユーザーがクエリを作成している間に投機的にキャッシュをウォーミングすることで、最初のトークンまでの時間を短縮します。
バッチ ツールのメタパターン回避策を使用して、Claude 3.7 Sonnet で並列ツール呼び出しを有効にします。
クロードの拡張された思考を使用して、予算管理で透明性のある段階的な推論を実現します。
拡張された思考とツールを組み合わせて、複数ステップのワークフロー中に透明性のある推論を実現します。
3 つのシンプルなマルチ LLM ワークフロー パターンは、パフォーマンスの向上のためにコストまたは遅延をトレードします。
1 つの LLM を生成に使用し、もう 1 つを評価フィードバック ループに使用するワークフロー パターン。
セントラル LLM は、タスクをワーカー LLM に動的に委任し、それらを組み合わせた結果を合成します。
大量の Claude リクエストを 50% c で非同期に処理します。

バッチを使用した ost 削減。
RAG、思考連鎖、自己改善手法を使用して、自然言語クエリを SQL に変換します。
プロンプト キャッシュを使用して埋め込む前にチャンクにコンテキストを追加することで、RAG の精度が向上します。
カスタムタスク用に Amazon Bedrock で Claude 3 Haiku を微調整するためのステップバイステップガイド。
合成テスト ケースを生成して、クロード プロンプト テンプレートを効果的に評価および改善します。
プロンプトのコンテキストをキャッシュして再利用することで、コストを削減し、詳細な指示による迅速な応答を実現します。
評価と高度なテクニックを備えた法的文書の要約に関する包括的なガイド。
クロードとともに、概要インデックス作成および再ランキング手法を使用して、RAG システムを構築および最適化します。
クロードと一緒に、RAG と保険チケットの思考連鎖を使用して分類システムを構築します。
強制または自動選択のtool_choiceパラメータを使用して、クロードがツールを選択する方法を制御します。
Claude のビジョンと、栄養ラベルなどの画像から構造化データを抽出するツールを組み合わせます。
メッセージ継続による事前入力技術を使用して、max_tokens 制限を超える長い応答を生成します。
クロードのビジョン機能を使用して最適な画像処理パフォーマンスを実現するためのヒントとテクニック。
タイプセーフな Claude ツール使用対話のための Pydantic モデルを使用して、検証済みのツールを作成します。
Deepgram で音声を文字に起こし、準備として Claude を使用して面接の質問を生成します。
Wolfram Alpha LLM APIを計算クエリと答えのためのClaudeツールとして統合します。
クロードに算術演算と数学的問題解決のための計算ツールを提供します。
顧客検索と注文管理のためのツールを使用して、クロードと一緒に顧客サービス チャットボットを構築します。
Claude のツール使用機能を使用して、さまざまな入力から構造化された JSON データを抽出します。
白紙ページの問題を解決するためのタスクの開始プロンプトを生成するプロンプト エンジニアリング ツール

問題。
クロードが検証のための文書ベースの質問に答えるときに、詳細な出典の引用を提供できるようにします。
LangChain v1 の更新されたエージェント フレームワーク パターンを使用して、Claude 3 で RAG エージェントを構築します。
Claude 3 Haiku を使用して URL 抽出により Web ページのコンテンツを取得し、要約します。
抽出には Haiku サブエージェント、合成には Opus を使用して財務レポートを分析します。
画像の理解と推論には、LlamaIndex の Anthropic MultiModal LLM 抽象化を使用します。
技術ニュースを知識ベースとして使用し、Claude と MongoDB を使用してチャットボット RAG システムを構築します。
主要な指標に関するクロードのパフォーマンスを測定および改善するための堅牢な評価システムを構築します。
プロンプトでルールとカテゴリを定義して、カスタマイズ可能なコンテンツ モデレーション フィルターを構築します。
制約されたサンプリングを行わずに効果的なプロンプト手法を使用して、Claude から信頼性の高い JSON 出力を取得します。
Claude とデータベース スキーマ コンテキストを使用して、自然言語の質問から SQL クエリを生成します。
ビジョンベースのテキスト分析のために画像を Claude 3 API に渡すチュートリアル。
Claude 3 のビジョンを使用して、画像や PDF から非構造化テキストを抽出して構造化します。
Claude のビジョン分析機能を使用して、チャート、グラフ、プレゼンテーションから洞察を抽出します。
ReAct Agent パターンで DocumentAgent を使用して、大規模なドキュメント コレクション用の RAG を構築します。
ツールベースの推論とアクションのワークフローのために、LlamaIndex を使用して ReAct エージェントを作成します。
ドキュメントの検索と質問応答のために、LlamaIndex を使用して基本的な RAG パイプラインを構築します。
複数ドキュメントの検索に LlamaIndex RouterQueryEngine を使用して、クエリをさまざまなインデックスにルーティングします。
LlamaIndex エンジンを使用して、複雑なクエリを複数のドキュメントにわたるサブ質問に分解します。
クロードを松ぼっくりベクトル データベースに接続して、検索拡張生成とセマンティック検索を実現します。
Claude API を使用して PDF ドキュメントを処理および要約する

h テキストの抽出とエンコード。
研究ワークフローのために Claude 2 を使用した反復的な Wikipedia 検索を示す従来のノートブック。
料理本のアイデアはありますか?コミュニティへの貢献を歓迎します。

## Original Extract

Practical guides and code examples for building with Claude. Learn prompting techniques, tool use, multimodal capabilities, and more.

Claude Cookbook
Cookbook Claude Cookbook
Practical guides and examples for using Claude effectively
Programmatic tool calling (PTC) Reduce latency and token consumption by letting Claude write code that calls tools programmatically in the code execution environment. Tool search with embeddings Scale Claude applications to thousands of tools using semantic embeddings for dynamic tool discovery. Automatic context compaction Manage context limits in long-running agentic workflows by automatically compressing conversation history. Giving Claude a crop tool for better image analysis Give Claude a crop tool to zoom into image regions for detailed analysis of charts, documents, and diagrams. Prompting for frontend aesthetics Guide to prompting Claude for distinctive, polished frontend designs avoiding generic aesthetics. Introduction to Claude Skills Create documents, analyze data, automate workflows with Claude's Excel, PowerPoint, PDF skills. All Cookbooks
All Categories Title Categories Author(s) Date Reproduce Claude's agentic search benchmark scores in the Messages API Jun 2026 • Evals Tools Build a Messages API harness that reproduces published DeepSearchQA and BrowseComp scores, using programmatic tool calling, server-side compaction, and task budgets.
Detect safety classifier blocks on Fable 5 and fall back to Opus 4.8 with server-side or SDK-based client-side fallback, including streaming behavior and the new billing changes.
Two async multi-agent patterns — a fixed N-agent team with peer messaging through a shared hub, and dynamically spawned async subagents — reduced to their bare messaging and lifecycle mechanics.
Deploy the research agent from notebook 00 through three tiers of operational maturity (Docker, Modal, Kubernetes) with the same container image and HTTP interface at every tier.
Heterogeneous team via the multiagent coordinator config — a coordinator runs three specialists (web-search researcher, file-reading librarian, rules-based pricer) with scoped toolsets to assemble a sales proposal. Covers the multiagent field, the thread_created / thread_message_received event types, and per-role tool scoping.
Build a grade-and-revise loop with Outcomes: a writer drafts a cited research brief, a stateless grader fetches every URL and checks every quote against a rubric, and feedback drives revisions until the brief passes. Covers user.define_outcome, the span.outcome_evaluation_* events, and how to write a rubric the grader can act on.
Give your Claude Managed Agents a Memory store so they learn and remember your users' preferences across multiple interactions.
Build a vulnerability-discovery agent with the Claude Agent SDK that threat-models a C target, hunts memory-safety bugs with built-in file tools, and triages findings into a structured report.
Wire Claude into your on-call flow: when an alert fires, the agent reads logs and runbooks, pinpoints the root cause, opens a fix PR, and waits for your approval before merging.
Build an analyst that turns a CSV into a narrative HTML report with interactive charts, using a sandboxed environment and file mounting.
Mention the bot with a CSV to get an analysis report in-thread, with multi-turn follow-ups on the same session.
Entry-point tutorial for the Claude Managed Agents API. Walks through agent / environment / session creation, file mounts, and the streaming event loop by getting an agent to fix three planted bugs in a calc.py package.
End-to-end production story for Managed Agents — vault-backed MCP credentials, the session.status_idled webhook pattern for human-in-the-loop without long-lived connections, and the resource lifecycle CRUD verbs.
Server-side prompt versioning — create v1, evaluate against a labelled test set, ship v2, detect a regression, roll back by pinning sessions to version 1. Covers agents.update, version pinning on sessions.create, and where the review gate moves when prompts are not code.
Build an agent that autonomously investigates IOCs by querying multiple threat intel sources, cross-referencing findings, mapping to MITRE ATT&CK, and producing structured reports for SIEM and SOAR integration.
List, read, rename, tag, and fork Agent SDK sessions on disk to build a conversation history sidebar without writing a transcript parser.
Build knowledge graphs from unstructured text using Claude for entity extraction, relation mining, deduplication, and multi-hop graph querying.
Compare context engineering strategies for long-running agents and learn when each applies, what it costs, and how they compose.
Port an OpenAI Agents SDK app to the Claude Agent SDK, mapping each primitive (tools, guardrails, sessions, handoffs) through a single expense-approval agent example.
Build an incident response agent with read-write MCP tools for autonomous diagnosis, remediation, and post-mortem documentation.
Manage long-running Claude conversations with instant session memory compaction using background threading and prompt caching.
Reduce latency and token consumption by letting Claude write code that calls tools programmatically in the code execution environment.
Scale Claude applications to thousands of tools using semantic embeddings for dynamic tool discovery.
Manage context limits in long-running agentic workflows by automatically compressing conversation history.
Build a low-latency voice assistant using ElevenLabs for speech-to-text and text-to-speech combined with Claude.
Give Claude a crop tool to zoom into image regions for detailed analysis of charts, documents, and diagrams.
Guide to prompting Claude for distinctive, polished frontend designs avoiding generic aesthetics.
Build financial dashboards and portfolio analytics using Claude's Excel, PowerPoint, PDF skills.
Create, deploy, and manage custom skills extending Claude with specialized organizational workflows.
Create documents, analyze data, automate workflows with Claude's Excel, PowerPoint, PDF skills.
Build a research agent using Claude Code SDK with WebSearch for autonomous research.
Build multi-agent systems with subagents, hooks, output styles, and plan mode features.
Connect agents to external systems via MCP servers for GitHub monitoring and CI workflows.
Run parallel agent evaluations on tools independently from evaluation task files.
Programmatically access and analyze your Claude API usage and cost data via Admin API.
Build AI agents with persistent memory using Claude's memory tool and context editing.
Reduce time-to-first-token by warming cache speculatively while users formulate their queries.
Enable parallel tool calls on Claude 3.7 Sonnet using batch tool meta-pattern workaround.
Use Claude's extended thinking for transparent step-by-step reasoning with budget management.
Combine extended thinking with tools for transparent reasoning during multi-step workflows.
Three simple multi-LLM workflow patterns trading cost or latency for improved performance.
Workflow pattern using one LLM for generation and another for evaluation feedback loop.
Central LLM dynamically delegates tasks to worker LLMs and synthesizes their combined results.
Process large volumes of Claude requests asynchronously with 50% cost reduction using batches.
Convert natural language queries to SQL using RAG, chain-of-thought, and self-improvement techniques.
Improve RAG accuracy by adding context to chunks before embedding with prompt caching.
Step-by-step guide to finetuning Claude 3 Haiku on Amazon Bedrock for custom tasks.
Generate synthetic test cases to evaluate and improve your Claude prompt templates effectively.
Cache and reuse prompt context for cost savings and faster responses with detailed instructions.
Comprehensive guide to summarizing legal documents with evaluation and advanced techniques.
Build and optimize RAG systems with Claude using summary indexing and reranking techniques.
Build classification systems with Claude using RAG and chain-of-thought for insurance tickets.
Control how Claude selects tools using tool_choice parameter for forced or auto selection.
Combine Claude's vision with tools to extract structured data from images like nutrition labels.
Generate longer responses beyond max_tokens limit using prefill technique with message continuation.
Tips and techniques for optimal image processing performance with Claude's vision capabilities.
Create validated tools using Pydantic models for type-safe Claude tool use interactions.
Transcribe audio with Deepgram and generate interview questions using Claude for preparation.
Integrate Wolfram Alpha LLM API as Claude tool for computational queries and answers.
Provide Claude with calculator tool for arithmetic operations and mathematical problem solving.
Build customer service chatbot with Claude using tools for customer lookup and order management.
Extract structured JSON data from various inputs using Claude's tool use capabilities.
Prompt engineering tool that generates starting prompts for your tasks to solve blank-page problem.
Enable Claude to provide detailed source citations when answering document-based questions for verification.
Build RAG agents with Claude 3 using LangChain v1's updated agent framework patterns.
Fetch and summarize web page content using Claude 3 Haiku via URL extraction.
Analyze financial reports using Haiku sub-agents for extraction and Opus for synthesis.
Use LlamaIndex's Anthropic MultiModal LLM abstraction for image understanding and reasoning.
Build chatbot RAG system with Claude and MongoDB using tech news as knowledge base.
Build robust evaluation systems to measure and improve Claude's performance on key metrics.
Build customizable content moderation filters by defining rules and categories in prompts.
Get reliable JSON output from Claude using effective prompting techniques without constrained sampling.
Generate SQL queries from natural language questions using Claude with database schema context.
Tutorial on passing images to Claude 3 API for vision-based text analysis.
Extract and structure unstructured text from images and PDFs using Claude 3's vision.
Extract insights from charts, graphs, and presentations using Claude's vision analysis capabilities.
Build RAG for large document collections using DocumentAgents with ReAct Agent pattern.
Create ReAct agents with LlamaIndex for tool-based reasoning and action workflows.
Build basic RAG pipeline with LlamaIndex for document retrieval and question answering.
Route queries to different indices using LlamaIndex RouterQueryEngine for multi-document search.
Decompose complex queries into sub-questions across multiple documents using LlamaIndex engine.
Connect Claude with Pinecone vector database for retrieval-augmented generation and semantic search.
Process and summarize PDF documents using Claude API with text extraction and encoding.
Legacy notebook showing iterative Wikipedia searches with Claude 2 for research workflows.
Have an idea for a cookbook? We welcome community contributions.
