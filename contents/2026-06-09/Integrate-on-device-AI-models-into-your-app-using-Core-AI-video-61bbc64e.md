---
source: "https://developer.apple.com/videos/play/wwdc2026/326/"
hn_url: "https://news.ycombinator.com/item?id=48459516"
title: "Integrate on-device AI models into your app using Core AI [video]"
article_title: "Integrate on-device AI models into your app using Core AI - WWDC26 - Videos - Apple Developer"
author: "sgt"
captured_at: "2026-06-09T13:08:33Z"
capture_tool: "hn-digest"
hn_id: 48459516
score: 1
comments: 0
posted_at: "2026-06-09T11:14:08Z"
tags:
  - hacker-news
  - translated
---

# Integrate on-device AI models into your app using Core AI [video]

- HN: [48459516](https://news.ycombinator.com/item?id=48459516)
- Source: [developer.apple.com](https://developer.apple.com/videos/play/wwdc2026/326/)
- Score: 1
- Comments: 0
- Posted: 2026-06-09T11:14:08Z

## Translation

タイトル: Core AI を使用してオンデバイス AI モデルをアプリに統合する [ビデオ]
記事のタイトル: Core AI を使用してオンデバイス AI モデルをアプリに統合する - WWDC26 - ビデオ - Apple Developer
説明: Qwen、Mistral、SAM3 などを含む、Apple シリコン用に最適化された人気のオープンソース モデルの厳選されたコレクションをご覧ください...

記事本文:
Core AI を使用してオンデバイス AI モデルをアプリに統合する - WWDC26 - ビデオ - Apple Developer
英語で見る
Core AI を使用してオンデバイス AI モデルをアプリに統合する
新しい Core AI フレームワークを使用して Apple シリコン向けに最適化された、Qwen、Mistral、SAM3 などの人気のオープンソース モデルの厳選されたコレクションをご覧ください。 Mac でモデルをダウンロード、実行、ベンチマークし、わずか数行のコードでアプリにモデルを統合する方法を学びます。モデルのコンパイルとオンデバイスの特殊化のための新しいワークフローを検討して、初回のモデルの読み込みを高速化します。 Xcode のコア AI ツールを使用してランタイム パフォーマンスをプロファイリングし、最適化する方法をご覧ください。
1:16 - アプリのコンセプト: カメラベースの語彙学習
7:40 - Core AI モデル リポジトリを使用してモデルを取得する
10:55 - Swift 統合コードの作成
13:05 - モデルの特殊化のレイテンシの診断
17:00 - Ahead-of-time (AOT) コンパイル
コア AI モデルを事前にコンパイルする
MLX を使用した分散推論とトレーニングを探索する
MLX を使用して Swift の数値計算を探索する
MLX を使用して Mac 上でローカル エージェント AI を実行する
コードをコピーする
11:01 - SAM3 画像セグメンテーションをロードして実行します
CoreAIImageSegmenter をインポート
// ロード
letsegmenter = ImageSegmenter を待ってみる (resourcesAt: sam3ModelURL)
// 使用する
let response = try awaitsegmenter.segment(画像: inputImage, プロンプト: "flower" )
マスク = 応答.セグメント.ファーストにしますか? .マスク
コードをコピーする
11:28 - 言語モデルをロードしてセッションを作成する
FoundationModels をインポートする
CoreAILanguageModels をインポート
// モデルインスタンスを作成する
let model = try await CoreAILanguageModel (resourcesAt: qwen3ModelURL)
// モデルを使用してセッションを作成します
let session = LanguageModelSession (モデル: モデル)
// レスポンスを生成する
let response = try await session.respond(to: "..." )
コードをコピーする
12:29 - @Gene を使用して構造化出力を生成する

有能な
FoundationModels をインポートする
CoreAILanguageModels をインポート
@Generable
struct VocabCard {
let chineseWord: 文字列
let english意味: 文字列
let exampleSentence: 文字列
}
let model = try await CoreAILanguageModel (resourcesAt: modelURL)
let session = LanguageModelSession (モデル: モデル)
let response = try await session.respond(
to: "花の単語カードを作成する" ,
生成: VocabCard 。自分自身
)
let カード: VocabCard = response.content
コードをコピーする
17:22 - 事前にコア AI モデルをコンパイルする
$ xcrun coreai-build apply MyModel.aimodel --platform iOS
Core AI の概要 — サーバー不要、トークンあたりのコストなし、クラウド遅延なしで、高度なオンデバイス AI 機能をアプリに導入できる新しいテクノロジー セットです。
1:16 - アプリのコンセプト: カメラベースの語彙学習
デモ アプリの紹介 — 学生が現実世界のオブジェクトにカメラを向けて、翻訳、例文、セグメント化された画像を含む単語カードを生成し、すべてデバイス上で実行できる iOS 言語学習アプリです。
アプリの AI 要件 (コンテンツ、言語、デバイスの制約) を定義し、適切なモデルを選択する方法: テキスト指示による画像セグメンテーションには SAM3、語彙カード生成には Qwen 0.6B (119 言語推論モデル)。
7:40 - Core AI モデル リポジトリを使用してモデルを取得する
coreai-models GitHub リポジトリを使用して、既製のエクスポート レシピを持つ人気モデルを見つける方法 - カタログを参照し、SAM3 および Qwen のエクスポート スクリプトを実行し、最適化された .aimodel ファイルを取得します。
Xcode で .aimodel ファイル (サイズ、プラットフォーム ターゲット、関数シグネチャ、テンソル形状) を検査し、coreai-models Swift パッケージを追加し、アプリの依存関係として CoreAILM および CoreAISegmentation ライブラリを選択する方法。
10:55 - Swift 統合コードの作成
ボットを使用するための Swift コードの書き方

h モデル - SAM3 をロードしてテキスト プロンプト セグメンテーションを実行し、単一の CoreAILanguageModel 行で Qwen をロードし、型付き語彙カード フィールドの構造化された @Generable 出力を備えた Foundation Models の使い慣れた LanguageModelSession API を使用します。
13:05 - モデルの特殊化のレイテンシの診断
新しい Core AI Instruments テンプレートを使用して、初回実行の遅延がモデルの特殊化 (特定のデバイス用の Core AI モデルをコンパイルするプロセス) によって引き起こされていることを特定し、それをいつどのように適切に処理するかを理解します。
意図的なデプロイメント戦略を設計する方法: 初回実行エクスペリエンスを使用して機能を導入し、すべてのユーザーの更新サイズが肥大化するのを避けるためにモデルをアプリ バンドルから除外し、ユーザーがオプトインした場合にのみバックグラウンド アセットを介してオンデマンド モデルのダウンロードをトリガーします。
17:00 - Ahead-of-time (AOT) コンパイル
coreai-build コマンドを使用して、開発マシン上で事前にコンパイルを実行する方法。これにより、デバイス アーキテクチャ固有のコンパイル済みモデル アセットが生成され、初回実行時のデバイス上での特殊化時間が大幅に短縮されます。
完全な iOS エクスペリエンスのライブ デモ: AOT コンパイルによる高速モデルの準備、実際のオブジェクト (岩、木、ひまわり) の SAM3 セグメント化、およびキャッシュされたモデルからのその後のシームレスな推論を伴う Qwen による中国語語彙カードの生成。
同じ Swift コードが変更を加えずに macOS でどのように実行されるか — 写真のフォルダーに対するバッチ処理の追加、より高品質な推論とピンイン生成のための Qwen3 8B へのステップアップ、カリキュラム生成のためのより長いコンテキストの使用、ロードトリップの写真を完全な授業計画に処理するライブ macOS デモ。
概要: Core AI は、プライベートなマルチプラットフォームのオンデバイス AI エクスペリエンスを構築するために必要なものをすべて提供します。サーバーもトークンごとのコストもかかりません。

、クラウドの遅延はありません。
Core AI を使用してオンデバイス AI モデルをアプリに統合する
Apple デベロッパー エンタープライズ プログラム
App Store 小規模ビジネス プログラム
セキュリティ研究装置プログラム

## Original Extract

Discover a curated collection of popular open-source models — including Qwen, Mistral, SAM3, and more — optimized for Apple silicon using...

Integrate on-device AI models into your app using Core AI - WWDC26 - Videos - Apple Developer
View in English
Integrate on-device AI models into your app using Core AI
Discover a curated collection of popular open-source models — including Qwen, Mistral, SAM3, and more — optimized for Apple silicon using the new Core AI Framework. Learn how to download, run, and benchmark models on your Mac, and integrate them into your app with just a few lines of code. Explore a new workflow for model compilation and on-device specialization to speed up first-time model load. Find out how to profile and optimize runtime performance with Core AI tools in Xcode.
1:16 - App concept: camera-based vocab learning
7:40 - Getting models with the Core AI models repository
10:55 - Writing the Swift integration code
13:05 - Diagnosing model specialization latency
17:00 - Ahead-of-time (AOT) compilation
Compiling Core AI models ahead of time
Explore distributed inference and training with MLX
Explore numerical computing in Swift with MLX
Run local agentic AI on the Mac using MLX
Copy Code
11:01 - Load and run SAM3 image segmentation
import CoreAIImageSegmenter
// Load
let segmenter = try await ImageSegmenter (resourcesAt: sam3ModelURL)
// Use
let response = try await segmenter.segment(image: inputImage, prompt: "flower" )
let mask = response.segments.first ? .mask
Copy Code
11:28 - Load a language model and create a session
import FoundationModels
import CoreAILanguageModels
// Create model instance
let model = try await CoreAILanguageModel (resourcesAt: qwen3ModelURL)
// Create session using the model
let session = LanguageModelSession (model: model)
// Generate response
let response = try await session.respond(to: "..." )
Copy Code
12:29 - Generate structured output with @Generable
import FoundationModels
import CoreAILanguageModels
@Generable
struct VocabCard {
let chineseWord: String
let englishMeaning: String
let exampleSentence: String
}
let model = try await CoreAILanguageModel (resourcesAt: modelURL)
let session = LanguageModelSession (model: model)
let response = try await session.respond(
to: "Create a vocab card for flower" ,
generating: VocabCard . self
)
let card: VocabCard = response.content
Copy Code
17:22 - Compile a Core AI model ahead of time
$ xcrun coreai-build compile MyModel.aimodel --platform iOS
Overview of Core AI — a new set of technologies that lets you bring advanced on-device AI capabilities to your apps with no server, no cost per token, and no cloud latency.
1:16 - App concept: camera-based vocab learning
Introduction to the demo app — an iOS language-learning app where students point their camera at real-world objects to generate vocab cards with translations, example sentences, and segmented images, all running on-device.
How to define your app's AI requirements — content, language, and device constraints — and select the right models: SAM3 for text-prompted image segmentation and Qwen 0.6B (a 119-language reasoning model) for vocab card generation.
7:40 - Getting models with the Core AI models repository
How to use the coreai-models GitHub repository to find popular models with ready-made export recipes — browsing the catalog, running the export script for SAM3 and Qwen, and getting optimized .aimodel files.
How to inspect .aimodel files in Xcode (size, platform targets, function signatures, tensor shapes), add the coreai-models Swift package, and select the CoreAILM and CoreAISegmentation libraries as app dependencies.
10:55 - Writing the Swift integration code
How to write the Swift code to use both models — loading SAM3 and running text-prompted segmentation, loading Qwen with a single CoreAILanguageModel line, and using the familiar LanguageModelSession API from Foundation Models with structured @Generable output for typed vocab card fields.
13:05 - Diagnosing model specialization latency
Using the new Core AI Instruments template to identify that first-run latency is caused by model specialization — the process that compiles a Core AI model for the specific device — and understanding when and how to handle it gracefully.
How to design a deliberate deployment strategy: using a first-run experience to introduce the feature, keeping models out of the app bundle to avoid bloating update size for all users, and triggering on-demand model download via Background Assets only when the user opts in.
17:00 - Ahead-of-time (AOT) compilation
How to use the coreai-build command to perform compilation ahead-of-time on your development machine — generating device-architecture-specific compiled model assets that dramatically reduce on-device specialization time during the first-run experience.
Live demo of the complete iOS experience: fast model preparation with AOT compilation, SAM3 segmenting real objects (rocks, wood, sunflower), and Qwen generating Mandarin vocab cards — with seamless subsequent inferences from the cached model.
How the same Swift code runs on macOS with no changes — adding batch processing for folders of photos, stepping up to Qwen3 8B for higher-quality reasoning and pinyin generation, using longer context for curriculum generation, and a live macOS demo processing road trip photos into a full lesson plan.
Summary: Core AI gives you everything you need to build private, multi-platform on-device AI experiences — no server, no cost per token, no cloud latency.
Integrate on-device AI models into your app using Core AI
Apple Developer Enterprise Program
App Store Small Business Program
Security Research Device Program
