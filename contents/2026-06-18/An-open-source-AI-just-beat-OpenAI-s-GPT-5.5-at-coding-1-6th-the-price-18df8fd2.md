---
source: "https://docs.z.ai/guides/llm/glm-5.2"
hn_url: "https://news.ycombinator.com/item?id=48581610"
title: "An open-source AI just beat OpenAI's GPT-5.5 at coding (1/6th the price)"
article_title: "GLM-5.2 - Overview - Z.AI DEVELOPER DOCUMENT"
author: "Raj_Sidwadkar"
captured_at: "2026-06-18T07:48:48Z"
capture_tool: "hn-digest"
hn_id: 48581610
score: 1
comments: 0
posted_at: "2026-06-18T06:40:39Z"
tags:
  - hacker-news
  - translated
---

# An open-source AI just beat OpenAI's GPT-5.5 at coding (1/6th the price)

- HN: [48581610](https://news.ycombinator.com/item?id=48581610)
- Source: [docs.z.ai](https://docs.z.ai/guides/llm/glm-5.2)
- Score: 1
- Comments: 0
- Posted: 2026-06-18T06:40:39Z

## Translation

タイトル: オープンソース AI がコーディングで OpenAI の GPT-5.5 を上回りました (価格は 1/6)
記事タイトル: GLM-5.2 - 概要 - Z.AI 開発者ドキュメント

記事本文:
GLM-5.2 - 概要 - Z.AI DEVELOPER DOCUMENT ドキュメント索引
/llms.txt で完全なドキュメントのインデックスを取得します。
さらに探索する前に、このファイルを使用して利用可能なすべてのページを検出します。
検索... ナビゲーション 言語モデル GLM-5.2 ガイド API リファレンス コーディング計画 リリースノート 規約とポリシー ヘルプセンター はじめる
GLM-5.2 1M コンテキストの紹介: 長期的なタスクを安定かつ実用的にする
ベンチマークと開発者の両方によって検証されたコーディング機能
ページをコピー ページをコピー 概要
プロジェクトレベルのコードベースの引き継ぎ: モデルにプロジェクト全体を一度に理解させる
現在のプロジェクトを読み、システム アーキテクチャ マップ、コア モジュールの責任、主要な API コントラクト、主要なデータ フロー、コア コール チェーン、潜在的な技術的負債、将来のリファクタリングで従う必要があるエンジニアリング上の制約を出力してください。
長期的なリファクタリング: 実際のエンジニアリング タスクをエンドツーエンドで実行させます
ビジネス ロジック、API シグネチャ、または実行時の動作を変更せずに、現在のモジュールの分離とリファクタリングを完了してください。まず、実行計画、影響範囲、リスク境界、検証方法を提供します。完了後、必要なテストを実行し、検証結果を出力します。
生産グレードの標準ストレステスト: 厳しいエンジニアリング上の制約を維持できるかどうかを確認する
現在のリポジトリのエンジニアリング標準に厳密に従ってください。新しい依存関係を導入したり、API コントラクトを変更したり、積極的に変更をコミットしたりしないでください。変更が完了したら、ビルド、lint、テストを実行し、検証結果と発見されたリスクを報告します。
モバイル オンデバイス デバッグ ループ: コードの実装からデバイスの検証まで
Kotlin にネイティブ Android クライアントを実装してください。

o 既存のサーバー側 API をサポートし、マルチセッションの会話、ストリーミング メッセージ、音声入力、通知、切断後の再接続をサポートします。完了したら、ADB を使用して実デバイスにインストールし、logcat とスクリーンショットを使用してデバッグします。
WeChat ミニ プログラムの開発: Web アプリから WeChat ミニ プログラムへの移行
現在の Web プロジェクトのすべての機能を WeChat ミニ プログラムに移行してください。 [native/Taro/uni-app] テクノロジー スタックを使用します。まず、ページ構造、コア ユーザー パス、バックエンド API コントラクト、およびパッケージ サイズ制限、ドメイン許可リスト、HTTPS 要件などのプラットフォーム制約を分析します。次に、ページ、コンポーネント、ページ ナビゲーション、およびデータ フローの実装を完了します。完了後、プロジェクトの実行方法、どの API が統合されているのか、どの機能がまだ未解決なのか、次に何を最適化できるのかを説明します。
ミニゲーム開発: ゲームプレイ ルールからプレイ可能なループまで
軽量のレベルベースのミニゲームを開発してください。まず、コアとなるゲームプレイ ループ、ステート マシン、レベル構造、スコアリング ルール、失敗と決済ロジックを設計し、次に開始、一時停止、再開、決済、再起動、ローカル保存などの基本機能を実装します。完了後、プロジェクトの構造、検証された機能、次のステップでの拡張の可能性について説明します。
研究の再現: 論文とデータから実行可能なエンジニアリング プロジェクトへ
この論文とデータセットに基づいて実験を再現してください。文書に明示的に記載されていない実装の詳細を記入します。 PyTorch を使用して、モデル アーキテクチャと損失関数を構築し、データ パイプラインとトレーニング/推論スクリプトを構築し、複数のファイル間で一貫性を保ってプロジェクトが正常に実行されることを確認します。実行時の問題を自律的に特定して修正し、書類が満たされていることを確認します

整列するまで項目ごとにリックし、再生パス、​​キーの変更、および残っているギャップを説明します。
コードからビデオへのループ: 自然言語のアイデアからデモ用ビデオまで
Remotionで新しいコンポジションを作成し、マップを追加してください。ロサンゼルスから開始し、ロサンゼルスに焦点を合わせたままカメラをズームアウトします。次に、ロサンゼルスからニューヨークまでのアニメーション ルートを描画し、カメラがそのルートをたどるようにします。旅にもう 1 つ立ち寄り先を追加します。今回はパリに行きます。
1M コンテキスト: 長期的なタスクを安定して実用化する
ベンチマークと開発者の両方によって検証されたコーディング機能
プロジェクトレベルのコンテキスト容量が強化され、コードベース全体を単一の推論ワークフロー内に配置できるようになります。
より安定した長期的なタスクの実行により、複雑なタスクを簡単に軌道から外れることなく継続的に進めることができます。
実稼働グレードのエンジニアリング標準への準拠がより確実になり、チーム開発ワークフローに厳しい制約を適用するのに役立ちます。
クライアント側およびモバイル エンジニアリング機能が強化され、アプリの生成を超えて完全なオンデバイス デバッグ ループがサポートされます。
API ドキュメント : API を呼び出す方法を学びます。
カール -X POST "https://api.z.ai/api/paas/v4/chat/completions" \
-H "コンテンツ タイプ: application/json" \
-H "認可: API キーのベアラー" \
-d '{
"モデル": "glm-5.2",
「メッセージ」: [
{
"役割": "システム",
"content": "あなたはシニア フルスタック ソフトウェア エンジニアで、フロントエンド開発、バックエンド アーキテクチャ設計、最新の Web テクノロジー スタックに精通しています。"
}、
{
"ロール": "ユーザー",
"content": "React + Node.js テクノロジー スタックを使用して、ホームページ、記事リスト ページ、記事詳細ページを含む個人のブログ Web サイトを設計して構築します。"
}
]、
「思考」: {
"タイプ": "有効"
}、
"reasoning_effort": "最大",
"max_tokens": 4096、
「テンペラ

チュア": 1.0
}'
ストリーミング呼び出しcurl -X POST "https://api.z.ai/api/paas/v4/chat/completions" \
-H "コンテンツ タイプ: application/json" \
-H "認可: API キーのベアラー" \
-d '{
"モデル": "glm-5.2",
「メッセージ」: [
{
"役割": "システム",
"content": "あなたはシニア フルスタック ソフトウェア エンジニアで、フロントエンド開発、バックエンド アーキテクチャ設計、最新の Web テクノロジー スタックに精通しています。"
}、
{
"ロール": "ユーザー",
"content": "React + Node.js テクノロジー スタックを使用して、ホームページ、記事リスト ページ、記事詳細ページを含む個人のブログ Web サイトを設計して構築します。"
}
]、
「思考」: {
"タイプ": "有効"
}、
"reasoning_effort": "最大",
「ストリーム」: true、
"max_tokens": 4096、
「温度」：1.0
}'
SDK をインストール # 最新バージョンをインストール
pip インストール zai-sdk
# またはバージョンを指定する
pip インストール zai-sdk== 0.2.3
インストールのインポートを確認する zai
print (zai.__version__ )
zai import ZaiClient からの基本的な呼び出し
client = ZaiClient( api_key = "your-api-key" ) # API キー
応答 = client.chat.completions.create(
モデル = "glm-5.2" 、
メッセージ = [
{
"ロール" : "システム" 、
"content" : "あなたはシニア フルスタック ソフトウェア エンジニアで、フロントエンド開発、バックエンド アーキテクチャ設計、最新の Web テクノロジー スタックに精通しています。" 、
}、
{
"ロール" : "ユーザー" 、
"content" : "ホームページ、記事リスト ページ、
[切り捨てられた]

## Original Extract

GLM-5.2 - Overview - Z.AI DEVELOPER DOCUMENT Documentation Index
Fetch the complete documentation index at: /llms.txt
Use this file to discover all available pages before exploring further.
Search... Navigation Language Models GLM-5.2 Guides API Reference Coding Plan Released Notes Terms and Policy Help Center Get Started
Introducing GLM-5.2 1M Context: Making Long-Horizon Tasks Stable and Practical
Coding Capabilities Validated by Both Benchmarks and Developers
Copy page Copy page ​ Overview
Project-Level Codebase Takeover: Let the Model Understand an Entire Project in One Go
Please read the current project and output a system architecture map, core module responsibilities, key API contracts, major data flows, core call chains, potential technical debt, and the engineering constraints that must be followed in future refactoring.
Long-Horizon Refactoring: Let It Run a Real Engineering Task End to End
Please complete the decoupling and refactoring of the current module without changing the business logic, API signatures, or runtime behavior. First provide the execution plan, impact scope, risk boundaries, and verification method. After completion, run the necessary tests and output the verification results.
Production-Grade Standards Stress Test: See Whether It Can Hold the Line on Hard Engineering Constraints
Please strictly follow the engineering standards of the current repository. Do not introduce new dependencies, do not modify API contracts, and do not commit changes proactively. After completing the modification, run the build, lint, and tests, then report the verification results and any uncovered risks.
Mobile On-Device Debugging Loop: From Code Implementation to Device Validation
Please implement a native Android client in Kotlin that connects to the existing server-side API and supports multi-session conversations, streaming messages, voice input, notifications, and reconnection after disconnection. After completion, install it on a real device using ADB, and debug it with logcat and screenshots.
WeChat Mini Program Development: Migrating from a Web App to a WeChat Mini Program
Please migrate all features of the current Web project into a WeChat Mini Program. Use the [native/Taro/uni-app] technology stack. First analyze the page structure, core user paths, backend API contracts, and platform constraints, including package size limits, domain allowlists, and HTTPS requirements. Then complete the implementation of pages, components, page navigation, and data flows. After completion, explain how to run the project, which APIs have been integrated, which features remain uncovered, and what can be optimized next.
Mini Game Development: From Gameplay Rules to a Playable Loop
Please develop a lightweight level-based mini game. First design the core gameplay loop, state machine, level structure, scoring rules, failure and settlement logic, then implement basic features including start, pause, resume, settlement, restart, and local save. After completion, explain the project structure, verified features, and possible next-step extensions.
Research Reproduction: From Paper and Data to a Runnable Engineering Project
Please reproduce the experiments based on this paper and dataset. Fill in implementation details not explicitly described in the paper. Use PyTorch to build the model architecture and loss functions, construct the data pipeline and training/inference scripts, and ensure the project runs successfully with consistency across multiple files. Autonomously identify and fix runtime issues, verify the paper’s metrics item by item until they are aligned, and explain the reproduction path, key changes, and any remaining gaps.
Code-to-Video Loop: From Natural-Language Ideas to a Demo-Ready Video
Please create a new composition in Remotion and add a map. Start from Los Angeles, zoom the camera out while keeping LA in focus. Then draw an animated route from Los Angeles to New York and have the camera follow the route. Add one more stop to the journey — this time, we are going to Paris.
1M Context: Making Long-Horizon Tasks Stable and Practical
Coding Capabilities Validated by Both Benchmarks and Developers
Stronger project-level context capacity, enabling an entire codebase to be placed within a single reasoning workflow;
More stable long-horizon task execution, allowing complex tasks to progress continuously without easily going off track;
More reliable adherence to production-grade engineering standards, helping enforce hard constraints in team development workflows;
Stronger client-side and mobile engineering capabilities, going beyond app generation to support a complete on-device debugging loop.
API Documentation : Learn how to call the API.
curl -X POST "https://api.z.ai/api/paas/v4/chat/completions" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer your-api-key" \
-d '{
"model": "glm-5.2",
"messages": [
{
"role": "system",
"content": "You are a senior full-stack software engineer, proficient in frontend development, backend architecture design, and modern web technology stacks."
},
{
"role": "user",
"content": "Design and build a personal blog website for me, including a homepage, article list page, and article detail page, using React + Node.js technology stack."
}
],
"thinking": {
"type": "enabled"
},
"reasoning_effort": "max",
"max_tokens": 4096,
"temperature": 1.0
}'
Streaming Call curl -X POST "https://api.z.ai/api/paas/v4/chat/completions" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer your-api-key" \
-d '{
"model": "glm-5.2",
"messages": [
{
"role": "system",
"content": "You are a senior full-stack software engineer, proficient in frontend development, backend architecture design, and modern web technology stacks."
},
{
"role": "user",
"content": "Design and build a personal blog website for me, including a homepage, article list page, and article detail page, using React + Node.js technology stack."
}
],
"thinking": {
"type": "enabled"
},
"reasoning_effort": "max",
"stream": true,
"max_tokens": 4096,
"temperature": 1.0
}'
Install SDK # Install latest version
pip install zai-sdk
# Or specify version
pip install zai-sdk== 0.2.3
Verify Installation import zai
print (zai. __version__ )
Basic Call from zai import ZaiClient
client = ZaiClient( api_key = "your-api-key" ) # Your API Key
response = client.chat.completions.create(
model = "glm-5.2" ,
messages = [
{
"role" : "system" ,
"content" : "You are a senior full-stack software engineer, proficient in frontend development, backend architecture design, and modern web technology stacks." ,
},
{
"role" : "user" ,
"content" : "Design and build a personal blog website for me, including a homepage, article list page, a
[truncated]
