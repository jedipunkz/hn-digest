---
source: "https://github.com/tetherto/qvac"
hn_url: "https://news.ycombinator.com/item?id=48687796"
title: "QVAC: Building local-first, peer-to-peer AI applications and systems"
article_title: "GitHub - tetherto/qvac: QVAC - Local AI SDK and libraries for building private, cross-platform, peer-to-peer AI applications. Run LLMs, speech-to-text, translation, and more locally on Linux, macOS, Windows, Android, and iOS. · GitHub"
author: "wslh"
captured_at: "2026-06-26T15:44:44Z"
capture_tool: "hn-digest"
hn_id: 48687796
score: 1
comments: 0
posted_at: "2026-06-26T15:26:38Z"
tags:
  - hacker-news
  - translated
---

# QVAC: Building local-first, peer-to-peer AI applications and systems

- HN: [48687796](https://news.ycombinator.com/item?id=48687796)
- Source: [github.com](https://github.com/tetherto/qvac)
- Score: 1
- Comments: 0
- Posted: 2026-06-26T15:26:38Z

## Translation

タイトル: QVAC: ローカルファーストのピアツーピア AI アプリケーションとシステムの構築
記事のタイトル: GitHub - tetherto/qvac: QVAC - プライベート、クロスプラットフォーム、ピアツーピア AI アプリケーションを構築するためのローカル AI SDK およびライブラリ。 LLM、音声テキスト変換、翻訳などを Linux、macOS、Windows、Android、iOS 上でローカルに実行します。 · GitHub
説明: QVAC - プライベート、クロスプラットフォーム、ピアツーピア AI アプリケーションを構築するためのローカル AI SDK およびライブラリ。 LLM、音声テキスト変換、翻訳などを Linux、macOS、Windows、Android、iOS 上でローカルに実行します。 - テザート/qvac

記事本文:
GitHub - tetherto/qvac: QVAC - プライベート、クロスプラットフォーム、ピアツーピア AI アプリケーションを構築するためのローカル AI SDK およびライブラリ。 LLM、音声テキスト変換、翻訳などを Linux、macOS、Windows、Android、iOS 上でローカルに実行します。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
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
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ

！
読み込み中にエラーが発生しました。このページをリロードしてください。
テザート
/
クヴァク
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1,778 コミット 1,778 コミット .claude .claude .cursor .cursor .github .github docs docs パッケージ パッケージ プラグイン/オープンコード プラグイン/オープンコード スクリプト スクリプト vcpkg-overlays vcpkg-overlays .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ウェブサイト •
ドキュメント •
サポート •
不和
QVAC は、ローカルファーストのピアツーピア AI アプリケーションとシステムを構築するためのオープンソースのクロスプラットフォーム エコシステムです。 QVAC を使用すると、LLM、音声、RAG などの AI タスクを Linux、macOS、Windows、Android、iOS 全体でローカルに実行したり、組み込みの P2P 機能を使用してピアに推論を委任したりできます。
ローカルファースト: AI モデルをロードし、独自のマシンで推論を実行します。サードパーティ API、SaaS、またはクラウドは関与しません。
P2P: BitTorrent、IPFS、ブロックチェーン ネットワークなど、AI を対象とした、止められないインターネット システムを構築します。
クロスプラットフォーム: ハードウェア、オペレーティング システム、JS ランタイム環境全体で一貫した開発者エクスペリエンス - コードを一度作成すれば、どこでも実行できます。
OpenAI 互換 API: より広範な AI エコシステムと統合します。
オープンソース: 完全に自由に使用および変更できます。上に構築し、貢献し、コミュニティの一員になりましょう。
QVAC は、JS SDK に統合された JavaScript ライブラリとツールで構成されています。 SDK は、QVAC を使用するための主要なエントリ ポイントです。これはタイプセーフであり、統一されたインターフェイスを通じてすべての QVAC 機能を公開します。 Node.js上で動作しますが、

裸のランタイム、および Expo 。
さらに、QVAC は、ツールを備えた CLI と、OpenAI 互換 API を公開する HTTP サーバーを提供します。 OpenAI API 形式を実装することで、QVAC はより広範な AI エコシステムと統合できます。
@qvac/sdk npm パッケージをプロジェクトにインストールします。次に、モデルをロードしてローカルで AI 推論を実行するか、組み込みの P2P 機能を使用して推論をピアに委任します。
サンプル ワークスペースを作成します。
mkdir qvac-examples
cd qvac-examples
npm init -y && npm pkg set type=module
SDKをインストールします。
npm インストール @qvac/sdk
クイックスタート スクリプトを作成します。
import {loadModel , LLAMA_3_2_1B_INST_Q4_0 , completed , unloadModel , } from "@qvac/sdk" ;
{を試してください
// モデルをメモリにロードします
const modelId = awaitloadModel ( {
モデルソース : LLAMA_3_2_1B_INST_Q4_0 、
モデルタイプ: "llm" 、
onProgress : (進行状況) => {
コンソール。ログ (進行状況) ;
} 、
} ) ;
// ロードされたモデルは複数回使用できます
const 履歴 = [
{
役割:「ユーザー」、
内容 : 「量子コンピューティングを一文で説明する」 、
} 、
] ;
const result = completed({modelId,history,stream:true});
for await (結果の const トークン .tokenStream) {
プロセスを標準出力 。書き込み (トークン) ;
}
// モデルをアンロードしてシステム リソースを解放する
await unloadModel ( {modelId } ) ;
}
キャッチ (エラー) {
コンソール。エラー ( "❌ エラー:" , エラー ) ;
プロセス 。出口 (1) ;
}
クイックスタート スクリプトを実行します。
ノードクイックスタート.js
機能
AI機能
完了: qvac-fabric-llm.cpp を介したテキスト生成とチャットのための LLM 推論。
テキスト埋め込み: qvac-fabric-llm.cpp による、セマンティック検索、クラスタリング、および取得のためのベクトル埋め込み生成。
翻訳: qvac-fabric-llm.cpp および Bergamot 経由のテキストからテキストへのニューラル機械翻訳 (NMT)。
文字起こし: qvac-ext-lib-whisper.cpp による音声からテキストへの自動音声認識 (ASR) または

NVIDIAインコ。
Text-to-Speech: ONNX ランタイムを介した Text-to-Speech (TTS) の音声合成。
OCR: ONNX ランタイムを介して画像からテキストを抽出する光学式文字認識 (OCR)。
画像生成: qvac-ext-stable-diffusion.cpp によるテキストから画像への生成。
微調整: LoRA を介して LLM をドメイン固有のタスクに適応させます。
マルチモーダル: 単一の会話コンテキスト内のテキスト、画像、その他のメディアに対する LLM 推論。
RAG: すぐに使える検索拡張生成ワークフロー。
委任された推論: ホールパンチ スタックを介してピアに推論を委任し、リソース共有を可能にします。
モデルのフェッチ: 分散モデル レジストリを介してピアから AI モデルをダウンロードします。
ブラインド リレー: リレー ノードを介してトラフィックをルーティングすることで、NAT/ファイアウォールを越えてピアを接続します。
プラグイン システム: 必要な AI 機能のみを組み込んで無駄のないアプリを構築し、カスタム機能をプラグインして SDK を拡張します。
ログ: 読み込み、推論、その他の操作中に何が起こっているかを可視化します。
ダウンロード ライフサイクル: モデルのダウンロードを一時停止および再開します。
シャード化されたモデル: 複数の部分にシャード化されたモデルをダウンロードします。
包括的な QVAC ドキュメントについては、 https://docs.qvac.tether.io を参照してください。
そこには、互換性マトリックス、環境/プラットフォームごとのインストール手順、各機能を使用するためのコード例のリファレンスなどが含まれています。
モノレポ構造の概要。 SDK、ライブラリ、ツールを含むすべての QVAC コンポーネントは /packages の下に存在します。すべてのコンポーネントが npm に公開されるわけではありません。
コア: エコシステム全体で共有される基本的な構成要素。
アドオン: 機能パッケージ — 各 QVAC 機能は 1 つ以上のアドオンによって実装されます。
SDK: 消費者向けの主要なエントリ ポイント。
ツール: エコシステムをサポートするユーザー向けのツールとサービス。
標準的な開発作業の場合

このモノリポジトリで使用される low については、 /docs/gitflow.md を参照してください。
各 QVAC コンポーネントの開発の詳細については、 /packages の下のそれぞれのサブディレクトリにあるドキュメントを参照してください。
QVAC アーキテクチャ全体については、 /docs/architecture を参照してください。
QVAC で何かを構築しましたか? README、Web サイト、またはアプリにバッジまたはバナーを追加します。これは、あなたのプロジェクトを強調し、他の人が QVAC を発見できるようにし、私たちのコミュニティを強化する簡単な方法です。
これらのバッジやバナーを使用することで、QVAC エコシステムの育成に貢献できます。
以下のバナーまたはバッジを選択してその Markdown スニペットをコピーするか、その画像 URL をコピーしてホストされている SVG アセットを直接使用します。
README ヘッダーに目立つように配置できる大判バッジ (240x60)。
[![QVAC で構築](https://raw.githubusercontent.com/tetherto/qvac/refs/heads/main/docs/branding/qvac-banner-dark-glow.svg)](https://github.com/tetherto/qvac)
バッジ
README 内の他のシールド/バッジと一緒に使用できるコンパクトなバッジ。
[![QVAC で構築](https://raw.githubusercontent.com/tetherto/qvac/refs/heads/main/docs/branding/qvac-badge-green-dark.svg)](https://github.com/tetherto/qvac)
について
QVAC - プライベート、クロスプラットフォーム、ピアツーピア AI アプリケーションを構築するためのローカル AI SDK およびライブラリ。 LLM、音声テキスト変換、翻訳などを Linux、macOS、Windows、Android、iOS 上でローカルに実行します。
Apache-2.0ライセンス
貢献する
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
72
フォーク
レポートリポジトリ
リリース
192
QVAC OpenCode プラグイン v0.1.0
最新の
2026 年 6 月 19 日
+ 191 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

QVAC - Local AI SDK and libraries for building private, cross-platform, peer-to-peer AI applications. Run LLMs, speech-to-text, translation, and more locally on Linux, macOS, Windows, Android, and iOS. - tetherto/qvac

GitHub - tetherto/qvac: QVAC - Local AI SDK and libraries for building private, cross-platform, peer-to-peer AI applications. Run LLMs, speech-to-text, translation, and more locally on Linux, macOS, Windows, Android, and iOS. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
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
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
tetherto
/
qvac
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1,778 Commits 1,778 Commits .claude .claude .cursor .cursor .github .github docs docs packages packages plugins/ opencode plugins/ opencode scripts scripts vcpkg-overlays vcpkg-overlays .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md View all files Repository files navigation
Website •
Docs •
Support •
Discord
QVAC is an open-source, cross-platform ecosystem for building local-first, peer-to-peer AI applications and systems. With QVAC, you can run AI tasks like LLMs, speech, RAG, and more locally across Linux, macOS, Windows, Android, and iOS — or delegate inference to peers using its built-in P2P capabilities.
Local-first: load AI models and perform inference on your own machine. No third-party APIs, SaaS, or cloud involved.
P2P: build unstoppable internet systems — like BitTorrent, IPFS, and blockchain networks, but for AI.
Cross-platform: consistent developer experience across hardware, operating systems, and JS runtime environments — write code once, run it everywhere.
OpenAI-compatible API: integrate with the broader AI ecosystem.
Open source: 100% free to use and modify — build on top, contribute back, be part of our community.
QVAC is composed of JavaScript libraries and tools that converge in the JS SDK. The SDK is the main entry point for using QVAC . It is type-safe and exposes all QVAC capabilities through a unified interface. It runs on Node.js, Bare runtime , and Expo .
Additionally, QVAC provides a CLI with tools and an HTTP server that exposes an OpenAI-compatible API . By implementing the OpenAI API format, QVAC can integrate with the broader AI ecosystem.
Install the @qvac/sdk npm package in your project. Then load models and run AI inference locally, or delegate inference to peers using the built-in P2P features.
Create the examples workspace:
mkdir qvac-examples
cd qvac-examples
npm init -y && npm pkg set type=module
Install the SDK:
npm install @qvac/sdk
Create the quickstart script:
import { loadModel , LLAMA_3_2_1B_INST_Q4_0 , completion , unloadModel , } from "@qvac/sdk" ;
try {
// Load a model into memory
const modelId = await loadModel ( {
modelSrc : LLAMA_3_2_1B_INST_Q4_0 ,
modelType : "llm" ,
onProgress : ( progress ) => {
console . log ( progress ) ;
} ,
} ) ;
// You can use the loaded model multiple times
const history = [
{
role : "user" ,
content : "Explain quantum computing in one sentence" ,
} ,
] ;
const result = completion ( { modelId , history , stream : true } ) ;
for await ( const token of result . tokenStream ) {
process . stdout . write ( token ) ;
}
// Unload model to free up system resources
await unloadModel ( { modelId } ) ;
}
catch ( error ) {
console . error ( "❌ Error:" , error ) ;
process . exit ( 1 ) ;
}
Run the quickstart script:
node quickstart.js
Functionalities
AI capabilities
Completion: LLM inference for text generation and chat via qvac-fabric-llm.cpp .
Text embeddings: vector embedding generation for semantic search, clustering, and retrieval, via qvac-fabric-llm.cpp .
Translation: text-to-text neural machine translation (NMT), via qvac-fabric-llm.cpp and Bergamot .
Transcription: automatic speech recognition (ASR) for speech-to-text via qvac-ext-lib-whisper.cpp or NVIDIA Parakeet .
Text-to-Speech: speech synthesis for text-to-speech (TTS) via ONNX Runtime .
OCR: optical character recognition (OCR) for extracting text from images via ONNX runtime.
Image generation: text-to-image generation via qvac-ext-stable-diffusion.cpp .
Fine-tuning: adapting LLMs to domain-specific tasks via LoRA.
Multimodal: LLM inference over text, images, and other media within a single conversation context.
RAG: out-of-the-box retrieval-augmented generation workflow.
Delegated inference: delegate inference to peers via the Holepunch stack , enabling resource sharing.
Fetch models: download AI models from peers via the distributed model registry.
Blind relays: connect peers across NATs/firewalls by routing traffic through relay nodes.
Plugin system : build lean apps by including only required AI capabilities, and extend the SDK by plugging in custom capabilities.
Logging: visibility into what's happening during loading, inference, and other operations.
Download Lifecycle: pause and resume model downloads.
Sharded models: download a model that is sharded into multiple parts.
For comprehensive QVAC documentation, see https://docs.qvac.tether.io .
There, you'll find the compatibility matrix, installation instructions per environment/platform , reference with code examples for using each functionality , and much more.
Monorepo structure overview. All QVAC components live under /packages , including the SDK, libraries, and tooling. Not every component is published to npm.
Core: foundational building blocks shared across the ecosystem.
Addon: capability packages — each QVAC capability is implemented by one or more addons.
SDK: primary entry point for consumers.
Tool: user-facing tools and services that support the ecosystem.
For the standard development workflow used in this monorepo, see /docs/gitflow.md .
For development specifics of each QVAC component, refer to the documentation in the respective subdirectory under /packages .
For the QVAC architecture as a whole, see /docs/architecture .
Built something with QVAC? Add a badge or banner to your README, website, or app. It is a simple way to highlight your project, help others discover QVAC, and strengthen our community.
By using these badges and banners, you help foster the QVAC ecosystem!
Choose a banner or badge below and copy its Markdown snippet, or copy its image URL and use the hosted SVG asset directly.
Large format badges (240x60) for prominent placement in your README header.
[![Built with QVAC](https://raw.githubusercontent.com/tetherto/qvac/refs/heads/main/docs/branding/qvac-banner-dark-glow.svg)](https://github.com/tetherto/qvac)
Badges
Compact badges for use alongside other shields/badges in your README.
[![Built with QVAC](https://raw.githubusercontent.com/tetherto/qvac/refs/heads/main/docs/branding/qvac-badge-green-dark.svg)](https://github.com/tetherto/qvac)
About
QVAC - Local AI SDK and libraries for building private, cross-platform, peer-to-peer AI applications. Run LLMs, speech-to-text, translation, and more locally on Linux, macOS, Windows, Android, and iOS.
Apache-2.0 license
Contributing
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
72
forks
Report repository
Releases
192
QVAC OpenCode Plugin v0.1.0
Latest
Jun 19, 2026
+ 191 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
