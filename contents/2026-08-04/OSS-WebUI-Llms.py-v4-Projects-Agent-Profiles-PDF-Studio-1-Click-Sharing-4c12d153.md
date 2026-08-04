---
source: "https://llmspy.org/docs/latest"
hn_url: "https://news.ycombinator.com/item?id=49167161"
title: "OSS WebUI Llms.py v4: Projects, Agent Profiles, PDF Studio, 1-Click Sharing"
article_title: "llms.py v4 is here! | llmspy.org"
author: "mythz"
captured_at: "2026-08-04T11:56:32Z"
capture_tool: "hn-digest"
hn_id: 49167161
score: 3
comments: 0
posted_at: "2026-08-04T11:34:59Z"
tags:
  - hacker-news
  - translated
---

# OSS WebUI Llms.py v4: Projects, Agent Profiles, PDF Studio, 1-Click Sharing

- HN: [49167161](https://news.ycombinator.com/item?id=49167161)
- Source: [llmspy.org](https://llmspy.org/docs/latest)
- Score: 3
- Comments: 0
- Posted: 2026-08-04T11:34:59Z

## Translation

タイトル: OSS WebUI Llms.py v4: プロジェクト、エージェント プロファイル、PDF Studio、1-Click 共有
記事タイトル: llms.py v4 が登場! | llmspy.org
説明: llms.py v4 の発表 - 新機能と更新が満載された最新のメジャー リリース

記事本文:
llms.py v4 が登場しました! | llmspy.org llms.py llms.py 検索 ⌘ K ドキュメント llms.py v4 が登場しました。 v3 リリース ノート はじめに インストール クイック スタート機能 CLI Web UI テーマ チャット UI エージェント プロファイル プロジェクト パブリッシング Analytics コア ツール サーバー ツール PDF Studio 電卓 UI コードの実行 UI KaTeX による数学のレンダリング モデル セレクター システム プロンプト ライブラリ プロバイダー 音声入力エージェント ブラウザ拡張機能 拡張機能の概要 組み込み拡張機能 UI 拡張機能 サーバー拡張機能ツールのサポート スキルの管理 コンピューターの使用 Gemini ファイル検索ストア モデル コンテキスト プロトコルマルチモーダル メディアの生成 認証資格情報 GitHub 認証構成 構成 アバター デプロイメント デプロイメント Docker デプロイメント カスタム llms .py ビルド llms.py v4 が登場しました。 llms.py v4 が登場しました!
llms.py v4 の発表 - 新機能とアップデートを満載した最新のメジャー リリース
これまで最大のリリースである llms.py v4 を発表できることを嬉しく思います。このマイルストーン バージョンでは、
私たちが構築してきたすべてのものを一緒に: チャット スレッド、プロジェクト、画像を公開して共有します。
新しい ai.llmspy.org ショーケースへの音声、設定可能なエージェント プロファイル
特殊なワークフロー、プロバイダーがホストするサーバー ツール、および新しいオーディオの巨大な波に対して、
音声、画像生成モデル。今すぐアップグレードしてすべてのロックを解除してください。
llms.py をインストールする最も簡単な方法は、 pip を使用することです。
pip インストール llms-py
最新バージョンに更新して最新の機能を入手してください。
pip install llms-py --upgrade
アップグレード後は、外部拡張機能もアップグレードすることをお勧めします。
llms --すべて更新
構成を最新の llms.json および Providers-extra.json にリセットするには、次のコマンドを実行します。
新しい PDF Studio 拡張機能は、コードファーストのドキュメント デザインを llms.py にもたらします。コードを記述する方法で PDF をデザインします。左側にプレーンテキストの Typst テンプレート、右側にレンダリングされたページが表示されます。

そうです、データを入力または編集するとリアルタイムで更新されます。
PDF Studio は、システムの PATH で Typst CLI コンパイラが利用可能になると自動的に有効になる組み込み拡張機能です。
# 貨物 (錆び)
カーゴインストール --locked typst-cli
# macOS (自作)
醸造インストールタイプスト
リアルタイムのライブ プレビューとスキーマ駆動型のフォーム UI
請求書、証明書、明細書などのドキュメントは、コンテンツ (.json データ) からレイアウト (.typ テンプレート) を分離します。コード ビューまたはインタラクティブな自動生成フォーム ビューで生の JSON を使用してドキュメント データを編集し、キーストロークごとにコンパイルされた PDF を即座に再レンダリングします。
AI を活用した編集とビジョン入力
[AI で編集] パネルを使用して、ドキュメントのレイアウトとスタイルの変更を平易な英語で説明します。スクリーンショットまたはサンプル PDF を添付すると、ビジョン モデルがデータ構築から設計を自動的に切り離して、再利用可能な Typst テンプレートと JSON データ モデルを作成し、自動修復コンパイラ エラー パスを備えます。
豊富なビジュアルピッカーにより、マークアップを覚えなくてもスタイリングが直感的に行えます。
タイポグラフィーとフォント ピッカー: システム フォントとカスタム フォントのプレビュー、サイズ変更、行の高さ、太さの調整。
ページ設定 : インタラクティブな用紙サイズ (A3 ～ A6、レター、リーガル、スライド)、余白、および向き。
共有ライブラリ ( lib.typ ) : lib.preview.typ ライブ デザイン システム検証でセントラル ハウス スタイルを維持します。
型付きアプリケーション コードの生成
ワンクリックでドキュメントの JSON スキーマを、厳密に型指定された C#、Python、TypeScript、または JavaScript クラスに変換します。バックエンド アプリケーションは、テンプレートの予想される構造と一致する JSON を確実に生成します。
完全なチュートリアルについては、完全な PDF Studio ドキュメントを確認してください。
新しいパブリッシング拡張機能を使用すると、AI の作成物を他のユーザーと共有できます。
世界 - ワークスペースから直接。チャット スレッド用の公開の読み取り専用リンクを生成します (静的)。

プロジェクトと個々のメディア ファイルはすべて ai.llmspy.org/m でホストされており、
すぐにご利用いただけます。
ツールバーの「共有」アイコンをクリックして、公開パネルを開きます。パブリッシャーアカウントの接続
完全に無料で匿名です。電子メールや個人情報は必要ありません。また、いつでも切断できます。
あなたが公開するものはすべて、新しい ai.llmspy.org Web サイト (一般公開) でホストされます。
コミュニティで共有されている AI 生成の最高のコンテンツを閲覧して発見できるショーケースです。
Images 、 Audio 、および Projects ギャラリーを切り替えて、作成されたものを調べます。
コード ブロック、テーブル、リッチ マークダウン形式を含む会話全体を、
永続的なスタイル付きの Web ページ。 8 つの美しいテーマから選択して、共有スレッドに
Nord のクールな北極のトーンから Matrix のネオングリーンの輝きまで、独特の外観。
プロジェクトを共有する (ゲームなど)
ゲーム、Web アプリ、または HTML/JS/CSS のフォルダーなどの静的プロジェクトを、ライブ パブリック URL にデプロイします。
シングルクリック。拡張機能はプロジェクトのビルド フォルダーを自動検出し、パッケージ化して、
共有可能なリンク。
画像と音声を共有する
AI によって生成された画像をメディア ギャラリー ライトボックスから直接共有し、生成されたオーディオ クリップを共有します
オーディオ プレーヤーから - どちらもメタデータとともにアップロードされ、独自の公開ページで提供されます。
完全なチュートリアルについては、公開ドキュメントを参照してください。
ストリーミングはすべてのプロバイダーにわたって実装されるようになったため、完全な完了が到着するのを待つのではなく、チャット UI で応答がリアルタイムでトークンごとにレンダリングされ、以前は非ストリーミング応答のみをサポートしていたプロバイダーを含むすべてのモデルに即時にフィードバックが提供されます。
ストリーミングは、Anthropic、OpenAI、Google、OpenRouter、Groq、Mistral、Cerebras、xAI、Fireworks、Ollama、LM Studio、その他すべてのプロバイダーで同じように機能します。

PenAI 互換プロバイダー - そのため、どのモデルやプロバイダーを選択しても、一貫したリアルタイム出力が得られます。
新しいプロファイル マネージャーは、エージェント プロファイルの表示、上書き、作成のための完全な UI を提供します。手動でファイルを編集する必要はありません。 [エージェント プロファイル セレクター] ドロップダウンの下部にある [プロファイルの管理] オプションからアクセスします。
組み込みプロファイル ( Chat 、 Coder 、 Planner ) は読み取り専用として表示され、サーバー定義のデフォルトが表示されます。 [設定] タブから組み込みプロファイルのデフォルトのモデルとテーマをオーバーライドできます。オーバーライドはユーザー設定として保存され、元のプロファイル ファイルを変更することなくすぐに有効になります。
+ ボタンを使用してまったく新しいカスタム プロファイルを作成します。それぞれに独自の名前、アバター、デフォルト モデル、テーマ、ツール/スキル制限、システム プロンプト ファイルが含まれます。カスタム プロファイルは、 SYSTEM.md 、 SYSTEM.template 、およびすべてのテンプレート変数ファイルのインライン編集に加えて、UI から直接プロンプト ファイルを追加および削除する機能をサポートします。
完全なチュートリアルについては、Profile Manager のドキュメントを参照してください。
OpenRouter を介したオーディオおよび音声モデル
llms.py は、 OpenRouter を通じて利用可能な音声および音声生成モデルをサポートするようになり、チャットおよび画像モデルとともに、音楽生成、音声合成、およびテキスト読み上げ機能をモデル セレクターにもたらします。
モデル セレクターのオーディオ (🔊) 出力フィルターと音声 (👤) 出力フィルターを使用して、利用可能なすべてのモデルを参照します。
オーディオモデル
モデル ID コンテキスト 価格 Lyria 3 クリップ プレビュー google/lyria-3-clip-preview 1.0M 無料 Lyria 3 Pro プレビュー google/lyria-3-pro-preview 1.0M 無料 GPT オーディオ openai/gpt-audio 128K $2.50 / $10.00 GPT Audio Mini openai/gpt-audio-mini 128K $0.60 / $2.40
Google の Lyria 3 は、DeepMind の最新の音楽生成モデルです - 現在、Clip プレビューと Pro プレビューの両方が無料で利用できます

OpenRouter上で。
モデル ID 価格 MAI-Voice-2 microsoft/mai-voice-2 無料 / $22.00 Grok Voice TTS 1.0 x-ai/grok-voice-tts-1.0 無料 / $15.00 Gemini 3.1 Flash TTS google/gemini-3.1-flash-tts-preview $1.00 / $22.00 Voxtral Mini TTSミストラライ/voxtral-mini-tts-2603 $0.05 / $0.20
最新のオーディオおよび音声モデルを選択するには、プロバイダーの構成をリセットします。
llms --リセットプロバイダー
画像生成の改善
OpenRouter プロバイダーをリファクタリングして、Chat Completion API の代わりに OpenRouter イメージ モデルのイメージ生成 API を使用します。これは将来性があり、イメージ生成のパフォーマンスと信頼性が向上します。
最新のシュート イメージ モデルのサポートを追加します。
フラックス、ドリームシェイパー、ジャガーノート、Ilustmix
新しいサーバー ツール機能は、OpenRouter サーバー ツールと Anthropic のサーバー ツールの両方を含む、プロバイダーがホストするツールの組み込みサポートを追加します。これにより、クライアント側の設定を行わずに、Web 検索、Web フェッチ、コード実行などのプロバイダーがホストする機能をチャット UI から直接有効にすることができます。
各ツールは標準の JSON スキーマで定義されており、フロントエンドはこれを使用して構成 UI とモデルに提示されるツール定義を動的に生成します。
ツールバーから [ツール] パネルを開き、[サーバー ツール] タブに切り替えて、選択したモデルのプロバイダーで使用できるツールを確認します。
詳細については、サーバー ツールのドキュメントを参照してください。
新しいエージェント プロファイル機能を使用すると、それぞれに独自のシステム プロンプト、デフォルト モデル、UI テーマ、制限されたツール/スキル セット、カスタム アバター、およびフッター アクション ボタンを備えた特殊な AI エージェントを構成できます。
チャット ヘッダーのエージェント プロファイル セレクター ドロップダウンからエージェントを切り替えます。組み込みプロファイルには、Chat (汎用)、Planner (複数ステップの推論とタスク分解)、Coder (ディスクへのコードの読み取りおよび書き込み) が含まれます。
強力な 2 フェーズのワークフロー: を使用します。

プランナーは目標を構造化されたアーキテクチャ計画に分割し、フッターのアクション ボタン ( PLAN.md に保存 、 プランの実行 、 PLAN.md の実行 ) を介してコーダーに引き渡し、段階的に実装します。
実際のワークフローは次のとおりです。プランナーには「SF GALAGA ゲームを構築する」というプロンプトが与えられ、完全なアーキテクチャ計画を作成して PLAN.md に保存し、コーダーが完全なゲームを実装しました。
カスタムパーソナルアシスタント
アシスタントのサンプル プロファイルは、ID、ユーザー コンテキスト、値、ツールの個別のマークダウン ファイルを構成する SYSTEM.template を使用して、カスタム アバターを使用して完全にパーソナライズされた AI コンパニオンを構築する方法を示しています。
カスタム アバター、onlyTools/onlySkills の制限、条件付きフッター アクション、複合システム プロンプト テンプレートを含む完全な設定リファレンスについては、エージェント プロファイルのドキュメントを参照してください。
新しいプロジェクト機能は、安全なワークスペース サンドボックス システムを提供します。各プロジェクトは、1 つ以上の許可されたディレクトリを定義します。すべての AI エージェント ファイルシステム ツール (読み取り、書き込み、編集、検索、リスト) は、それらのパスにのみ制限されます。
プロジェクトは、カスタム絶対パスとともにパス エイリアス (サーバーの作業ディレクトリの場合は $WORKSPACE、システムの一時ディレクトリの場合は $TEMP) をサポートします。存在しないディレクトリは保存時に自動的に作成されます。
左上のヘッダーにある [ワークスペースとプロジェクト] ドロップダウンからプロジェクトを切り替えるか、デフォルトのワークスペースにリセットしてサンドボックスの制限を削除します。プロジェクトはユーザーごとに ~/.llms/user/{username}/projects/projects.json に保存されます。
詳細については、プロジェクトのドキュメントを参照してください。
最新モデルをチェックしてください。 Grok 4.3、GPT-5.5、GPT Image 2、Opus 4.7、Sonnet 4.6、Gemini 3.1、Gemma 4、DeepSeek v4、Kimi K2.6、GLM 5.1、MiMo V2.5 など...
画像生成モデルを再構築する
の新しい Recraft 画像生成モデルが利用可能になりました

オープンルーターから:
Recraft V4 Pro (画像あたり 0.25 ドル)
llms.json の Anthropic プロバイダーの npm 構成を変更して、 @ai-sdk/anthropic の代わりに @ai-sdk/anthropic-cli を使用することで、既存のクロード コード サブスクリプションを利用できるようになりました。例:
{
"人類" : {
"有効" : true 、
"npm" : "@ai-sdk/anthropic-cli" ,
...
}
}
このプロバイダーはすべてのリクエストをクロード バイナリにルーティングするため、その機能と統合は制限されています。ツール呼び出しやスキルはサポートされていません。それ以外の場合は、システム プロンプト、チャット履歴、画像やドキュメントの添付ファイルなど、期待される他の機能もサポートされます。
同時に、クロード バイナリのいくつかの賢さの恩恵を受けることができ、長い会話を処理するためのシステム プロンプトと最適化が組み込まれているため、一部のユースケースでは良いオプションとなる可能性があります。
Fireworks の大規模言語モデルのサポート
新しいプロバイダーとして、市場をリードする 200 tok/s で GLM 5 、 Kim K2.5 、 MiniMax M2.5 、 DeepSeek V3.2 などの主要なオープンソース モデルをホストする高速推論プラットフォームである Fireworks AI のサポートが追加されました。
すべてのテキスト モデルは推論とツールの使用をサポートしているため、Fireworks は速度が重要なエージェント ワークフローにとって優れた選択肢となります。こちらは Fireworks 作成による Kim K2.5 です

[切り捨てられた]

## Original Extract

Announcing llms.py v4 - the latest major release packed with new features and updates

llms.py v4 is here! | llmspy.org llms.py llms.py Search ⌘ K Documentation llms.py v4 is here! v3 Release Notes Getting Started Installation Quick Start Features CLI Web UI Themes Chat UI Agent Profiles Projects Publishing Analytics Core Tools Server Tools PDF Studio Calculator UI Run Code UI Rendering Math with KaTeX Model Selector System Prompts Library Providers Voice Input Agent Browser Extensions Extensions Overview Built-in Extensions UI Extensions Server Extensions Tool Support Manage Skills Computer Use Gemini File Search Stores Model Context Protocol Multimodal Media Generation Authentication Credentials GitHub Auth Configuration Configuration Avatars Deployment Deployment Docker Deployment Custom llms .py Build llms.py v4 is here! llms.py v4 is here!
Announcing llms.py v4 - the latest major release packed with new features and updates
We're thrilled to announce llms.py v4 - our biggest release yet! This milestone version brings
together everything we've been building: publish and share your chat threads, projects, images and
audio to the new ai.llmspy.org showcase, configurable Agent Profiles
for specialized workflows, provider-hosted Server Tools , and a huge wave of new Audio ,
Speech , and Image generation models. Upgrade now to unlock it all.
Easiest way to install llms.py is using pip :
pip install llms-py
Get the latest features by updating to the latest version :
pip install llms-py --upgrade
After upgrading, it's recommended to also upgrade any external extensions:
llms --update all
To reset your configuration to the latest llms.json and providers-extra.json run:
The new PDF Studio extension brings code-first document design to llms.py. Design PDFs the way you write code: a plain-text Typst template on the left, the rendered page on the right, updating in real-time as you type or edit data.
The PDF Studio is a built-in extension that is automatically enabled whenever the Typst CLI compiler is available on your system PATH :
# Cargo (Rust)
cargo install --locked typst-cli
# macOS (Homebrew)
brew install typst
Real-Time Live Preview & Schema-Driven Form UI
Documents like invoices, certificates, and statements separate layout ( .typ template) from content ( .json data). Edit document data using raw JSON in Code View or an interactive, auto-generated Form View -re-rendering the compiled PDF instantly on every keystroke.
AI-Powered Editing & Vision Input
Describe document layout and styling changes in plain English using the Edit with AI panel. Attach a screenshot or sample PDF to let vision models automatically decouple design from data-building a reusable Typst template and JSON data model for you, complete with auto-healing compiler error passes.
Rich visual pickers make styling intuitive without memorizing markup:
Typography & Font Picker : System & custom font previewing, sizing, line-height, and weight adjustments.
Page Setup : Interactive paper sizing (A3–A6, Letter, Legal, Slides), margins, and orientation.
Shared Library ( lib.typ ) : Maintain a central house style with lib.preview.typ live design system verification.
Typed Application Code Generation
Turn your document's JSON schema into strongly typed C#, Python, TypeScript, or JavaScript classes with a single click-ensuring your backend application produces JSON matching your template's expected structure.
Check out the full PDF Studio Documentation for the complete walkthrough.
The new Publishing extension lets you share your AI creations with the
world - straight from your workspace. Generate public, read-only links for chat threads, static
projects, and individual media files, all hosted on ai.llmspy.org/m and
available instantly.
Click the Share icon in the toolbar to open the publishing panel. Connecting a publisher account
is completely free and anonymous - no email or personal details required - and you can disconnect at any time.
Everything you publish is hosted on the new ai.llmspy.org website - a public
showcase where you can browse and discover the best AI-generated content shared by the community.
Switch between the Images , Audio , and Projects galleries to explore what's been created:
Publish an entire conversation - including code blocks, tables, and rich markdown formatting - as a
permanent, styled web page. Choose from 8 beautiful themes to give your shared threads a
distinctive look, from the cool arctic tones of Nord to the neon-green glow of Matrix .
Share Projects (e.g. Games)
Deploy a static project - a game, web app, or any folder of HTML/JS/CSS - to a live public URL with a
single click. The extension auto-detects your project's build folder, packages it, and returns a
shareable link.
Share Images & Audio
Share AI-generated images directly from the media gallery lightbox and generated audio clips
from the audio player - both uploaded with their metadata and served on their own public page.
See the Publishing docs for the full walkthrough.
Streaming is now implemented across all providers , so responses render token-by-token in real time in the Chat UI instead of waiting for the full completion to arrive - giving immediate feedback for every model, including providers that previously only supported non-streamed responses.
Streaming works the same way across every provider - Anthropic, OpenAI, Google, OpenRouter, Groq, Mistral, Cerebras, xAI, Fireworks, Ollama, LM Studio and all other OpenAI-compatible providers - so you get consistent real-time output no matter which model or provider you choose.
The new Profile Manager provides a full UI for viewing, overriding, and creating Agent Profiles - no manual file editing required. Access it from the Manage Profiles option at the bottom of the Agent Profile Selector dropdown.
Built-in profiles ( Chat , Coder , Planner ) are displayed as Read-only with their server-defined defaults visible. You can override the default model and theme for any built-in profile through the Settings tab - overrides are stored as user preferences and take effect immediately without modifying the original profile files.
Create entirely new custom profiles with the + button - each with its own name, avatar, default model, theme, tool/skill restrictions, and system prompt files. Custom profiles support inline editing of SYSTEM.md , SYSTEM.template , and all template variable files, plus the ability to add and delete prompt files directly from the UI.
See the Profile Manager docs for the full walkthrough.
Audio & Speech Models via OpenRouter
llms.py now supports Audio and Speech generation models available through OpenRouter , bringing music generation, audio synthesis, and text-to-speech capabilities to the model selector alongside chat and image models.
Use the Audio (🔊) and Speech (👤) output filters in the model selector to browse all available models.
Audio Models
Model ID Context Price Lyria 3 Clip Preview google/lyria-3-clip-preview 1.0M Free Lyria 3 Pro Preview google/lyria-3-pro-preview 1.0M Free GPT Audio openai/gpt-audio 128K $2.50 / $10.00 GPT Audio Mini openai/gpt-audio-mini 128K $0.60 / $2.40
Google's Lyria 3 is DeepMind's latest music generation model - both the Clip and Pro previews are available currently free on OpenRouter.
Model ID Price MAI-Voice-2 microsoft/mai-voice-2 Free / $22.00 Grok Voice TTS 1.0 x-ai/grok-voice-tts-1.0 Free / $15.00 Gemini 3.1 Flash TTS google/gemini-3.1-flash-tts-preview $1.00 / $22.00 Voxtral Mini TTS mistralai/voxtral-mini-tts-2603 $0.05 / $0.20
To pick up the latest Audio and Speech models, reset your providers configuration:
llms --reset providers
Image Generation Improvements
Refactor OpenRouter provider to use Image Generation API for OpenRouter image models instead of the Chat Completion API, which is future proof and improves image generation performance and reliability.
Add support for latest Chutes Image Models:
Flux, Dreamshaper, Juggernaut & Ilustmix
The new Server Tools feature adds built-in support for provider-hosted tools, including both OpenRouter Server Tools and Anthropic's Server Tools . This lets you enable provider-hosted capabilities like web search , web fetch , and code execution directly from the chat UI without any client-side setup.
Each tool is defined in standard JSON Schema which the frontend uses to dynamically generate the configuration UI and the tool definition presented to the model.
Open the Tools panel from the toolbar, switch to the SERVER TOOLS tab to see the tools available for the selected model's provider:
See the Server Tools docs for more details.
The new Agent Profiles feature lets you configure specialized AI agents, each with its own system prompt, default model, UI theme, restricted tool/skill set, custom avatar, and footer action buttons.
Switch between agents from the Agent Profile Selector dropdown in the chat header. Built-in profiles include Chat (general-purpose), Planner (multi-step reasoning and task decomposition), and Coder (reads and writes code to disk).
A powerful two-phase workflow: use the Planner to break a goal into a structured architecture plan, then hand off to the Coder via footer action buttons ( Save to PLAN.md , Execute Plan , Execute PLAN.md ) to implement it step by step.
Here's the workflow in action - the Planner was given the prompt "Build a Sci-Fi GALAGA game" , produced a full architecture plan saved to PLAN.md , and the Coder implemented the complete game:
Custom Personal Assistant
The assistant example profile shows how to build a fully personalized AI companion using a SYSTEM.template that composes separate markdown files for identity, user context, values, and tools - with a custom avatar:
See the Agent Profiles docs for the full configuration reference including custom avatars, onlyTools / onlySkills restrictions, conditional footer actions, and composite system prompt templates.
The new Projects feature provides a secure workspace sandboxing system. Each project defines one or more allowed directories - all AI agent filesystem tools (read, write, edit, search, list) are restricted exclusively to those paths.
Projects support path aliases ( $WORKSPACE for the server's working directory, $TEMP for the system temp dir) alongside any custom absolute paths. Non-existent directories are automatically created on save.
Switch between projects from the Workspaces & Projects dropdown in the top-left header, or reset to Default Workspace to remove sandbox restrictions. Projects are persisted per-user at ~/.llms/user/{username}/projects/projects.json .
See the Projects docs for more details.
Checkout the latest models inc. Grok 4.3, GPT-5.5, GPT Image 2, Opus 4.7, Sonnet 4.6, Gemini 3.1, Gemma 4, DeepSeek v4, Kimi K2.6, GLM 5.1, MiMo V2.5 and more...
Recraft Image Generation Models
New Recraft image generation models from are now available from Open Router:
Recraft V4 Pro ($0.25 per image)
You can now make use of your existing Claude Code Subscription by changing the Anthropic provider npm configuration in your llms.json to use @ai-sdk/anthropic-cli instead of @ai-sdk/anthropic , e.g:
{
"anthropic" : {
"enabled" : true ,
"npm" : "@ai-sdk/anthropic-cli" ,
...
}
}
This provider routes all requests to the claude binary so it's functionality and integration is limited, e.g. it doesn't support tool calling or skills. Otherwise it supports other features you'd expect like System Prompts, Chat History, Image and Document attachments, etc.
At the same time it's able to benefit from some smarts in the claude binary it's built-in system prompts and optimizations for handling long conversations, so it can be a good option for some use cases.
Support for Fireworks Large Language Models
Added support for Fireworks AI as a new provider, a fast inference platform hosting the leading open-source models including GLM 5 , Kimi K2.5 , MiniMax M2.5 and DeepSeek V3.2 at a market-leading 200 tok/s .
All text models support reasoning and tool use , making Fireworks an excellent choice for agentic workflows where speed matters. Here's Kimi K2.5 via Fireworks creat

[truncated]
