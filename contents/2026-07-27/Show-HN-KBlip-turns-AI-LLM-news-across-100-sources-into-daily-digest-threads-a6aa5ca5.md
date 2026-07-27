---
source: "https://kblip.com/releases"
hn_url: "https://news.ycombinator.com/item?id=49072174"
title: "Show HN: KBlip – turns AI/LLM news across 100 sources into daily digest threads"
article_title: "Releases · KBlip - The Signal Feed for the AI/LLM World"
author: "jonam21"
captured_at: "2026-07-27T17:31:28Z"
capture_tool: "hn-digest"
hn_id: 49072174
score: 2
comments: 2
posted_at: "2026-07-27T16:41:10Z"
tags:
  - hacker-news
  - translated
---

# Show HN: KBlip – turns AI/LLM news across 100 sources into daily digest threads

- HN: [49072174](https://news.ycombinator.com/item?id=49072174)
- Source: [kblip.com](https://kblip.com/releases)
- Score: 2
- Comments: 2
- Posted: 2026-07-27T16:41:10Z

## Translation

タイトル: Show HN: KBlip – 100 のソースにわたる AI/LLM ニュースを毎日のダイジェスト スレッドに変換します
記事のタイトル: リリース · KBlip - AI/LLM 世界のためのシグナル フィード
説明: AI/LLM の世界で新しいモデル、ツール、フレームワークが発表されます。
HN テキスト: こんにちは、HN、メーカーです。 KBlip は、Reddit、HN、arXiv、GitHub、YouTube、および最大 100 件の RSS/ブログ フィードを監視します。
AI 関連で、同じ記事の報道を 1 つの「スレッド」にまとめます (すべてのソースが次のようにリンクされています)。
証拠)、それを 5 つのフィードに分類します: リリース、ニュース、ソーシャル、開発 (リサーチ + ハードウェア)、
製品/アイデア、およびトレンド。 AI/LLM に関連した最新の出来事をキャッチアップするための私の日課は、私の時間のかなりの部分を占めており、異なる製品間を複数回切り替える必要があり、さらに悪いことに、それらの製品のいずれかのウサギの穴にはまってしまうことがあります。可視性を高め、すべてを 1 か所に集約して、情報をキャッチアップするための時間と労力を減らし、ノイズや冗長性もカットしたいと考えていました。強調すべき点がいくつかあります: セクションごとの毎日のダイジェスト (1 文の TL;DR)、
読むよりも聞きたい場合は、音声の「ブリーフミー」モードがあり、PWA としてインストールされます。 VPS で自己ホストされているソロ プロジェクト、LLM/埋め込み側の DeepSeek + Voyage。アカウントも広告も追跡もありません。

記事本文:
リリース · KBlip - AI/LLM 世界のためのシグナル フィード K Blip リリース ニュース 製品 ソーシャル トレンド開発 ⌘K · · リリース / 発売
新しいモデル、ツール、フレームワークの発売
概要 最近のヘッドライン 10 件の更新 拡大 📣 ニュース — Moonshot AI が Kimi K3 オープンウェイトをリリース: SimpleQA で 91.2% の 2.8T パラメーター MoE モデル、許容ライセンス [19] [1]
⚡ 能力 — AI コーディング エージェントは 750,000 LOC アプリを 3 日間でバグなし、31 回の検証パスでリファクタリングします [7]
🛠️ ツール — WISP ストリーミング エンジンは、3 層ストリーミングを介して消費者向けハードウェア上で Kim K3 などの 2T+ MoE モデルを実行します [9]
🛠️ ツール — Krasis ランタイムは、単一の RTX PRO 6000 Blackwell 96GB で第 4 四半期に Ornith-397B をストリーミングし、2,354 tok/s のプレフィルを達成しました [6]
🛠️ ツール — Open WebUI v0.11.0 には、チャット ビュー、管理パネル、軽量のタイポグラフィーを備えた完全なインターフェイスの再設計が含まれています [11]
🛠️ ツール — コミュニティは、Qwen モデル用のカスタム FlashAttendant を使用して SGLang を V100 GPU に移植します [10]
🛠️ ツール — マルチストリーム アテンションを備えた Minimax M3 サポートは、ローカルのロングコンテキスト処理のために llama.cpp に統合されました [22]
⚡ 機能 — POCKET-35B エージェント モデルは CPU で 59 t/s で実行され、GGUF 量子化を備えた電話機に適合します [23]
🛠️ ツール — World-Model-Optimizer は半分のコストでモデルを抽出してルーティングします、オープンソース [20]
🛠️ ツール — OpenClaude 改良されたコーディング エージェントが GitHub でリリースされ、クラウド API とローカル モデルをサポート [21]
Moonshot AI が寛容ライセンス付きの Kim K3 モデルをリリース
Moonshot AI は、「サービスとしてのモデル」を広く定義するライセンスに基づいて Kimi K3 言語モデルをリリースし、総収益がしきい値を超えない限り、サードパーティの推論サービスや微調整サービスを制限します。このライセンスは、オープン性と商業的保護のバランスをとることを目的としています。
Composer v3 は最初に 2.85,000 のダウンロードでからかわれた

時間
Composer v3 のティーザー イメージには、最初の 1 時間以内に 2,85,000 件のダウンロードがあり、初期の強い関心が示されています。この投稿は、AI ラボが多忙な一日を迎えることを示唆しており、複数のリリースや発表があることを示唆しています。
Vanara が CI 検証済みの Claude Code エージェント カタログの無料枠を開始
Vanara は、Claude Code エージェントのカタログであり、プッシュごとにパブリック CI で実行可能なスクリプトを介して検証され、要求どおりに動作することが保証されます。無料枠がオープンし、ユーザーは「npx vanara install security-pack」などのエージェント パックをインストールできるようになりました。
Anthropic の迅速なモデルリリースは競合他社を上回ります
アンスロピックは6月に『Claude Fable 5』をリリースしたが、輸出規制を理由に政府によって中止された。その後、OpenAI は 7 月に GPT-5.6 をリリースし、Fable に勝ると主張しました。 2 週間後、Anthropic は Opus 5 を Fable の半額で販売し、Fable と GPT-5.6 を抑えて独立系リーダーボードのトップになりました。これは Anthropic の 2 か月以内で 4 番目のモデルとなり、タイムリーなリリース戦略を示しています。
↑ 2026 年 7 月 27 日月曜日、午後 2 時 37 分更新 — Anthropic が Opus 5 をリリースし、OpenAI の GPT-5.6 主張を受けてリーダーボードのトップに立った。
カーソル研究部門が並列マルチエージェント群スキルをテスト
ユーザーは、エージェント群に関する Cursor の研究部門からの洞察を共有し、同様のイデオロギーを新しい並列マルチエージェント群スキルに適用します。このスキルは、複数のエージェントを同時に調整することで効率を向上させることを目的としています。
Krasis ランタイムは第 4 四半期にシングル RTX PRO 6000 Blackwell 96GB で Ornith-397B を実行
開発者は、限られた VRAM を介して大規模なモデルをストリーミングする MoE に重点を置いたランタイムである Krasis を構築しました。第 4 四半期には、単一の RTX PRO 6000 Blackwell 96GB 上で Ornith-1.0-397B を実行し、エキスパートを CPU RAM に保持し、VRAM 常駐を動的に管理することで、2,354 tok/s のプリフィルと ~20 ～ 24 tok/s のデコードを達成しました。
AI コーディング エージェントのリファクタリング 750k L

バグゼロで 3 日で OC アプリを完成
AI Sovereign Labs のケーススタディでは、AI コーディング エージェントが 3 日間で 750,000 行のアプリケーションを自律的にリファクタリングしたことが示されています。エージェントは 31 回の検証パスを実行し、201 個のエラーを修正し、バグ、リグレッション、技術的負債がゼロのコードを出荷しました。このプロジェクトは、最初から書き直したり、トレーニング データのクローンを作成したりするものではありませんでした。
XYZ AI Lab がオープンウェイト Deep Search エージェントである XYZ-Aquila-mini をリリース
XYZ AI Lab は、AI4AI パイプラインを介して Qwen3.6-35B-A3B から事後トレーニングされたオープンウェイト Deep Search エージェントである XYZ-Aquila-mini をリリースしました。このモデルは、人間が定義した制約と AI 主導の障害診断を備えた、限定探索検索タスク用に設計された思考モデルです。
WISP ストリーミング エンジンはコンシューマ ハードウェアで 2T+ MoE モデルを実行します
WISP は C/CUDA ストリーミング エンジンで、3 層ストリーミング (VRAM→RAM→NVMe SSD) を介して、消費者向けハードウェア上で Kim K3 (2.8T) や GLM-5.2 (744B) などの大規模 MoE モデルを実行できるようにします。自己組織化 LRU キャッシュ、MLA アテンション サポート、および 2.2 ～ 2.8 倍のスループットを実現する同一ファミリーの投機的デコードを備えています。 Mixtral-8x7B で動作することを確認しました。
コミュニティは、Qwen モデル用のカスタム FlashAttend を備えた SGLang を V100 GPU に移植します。
開発者は SGLang をフォークして NVIDIA V100 GPU のサポートを追加し、カスタム TeilLang FlashAttend を作成し、オープンソースの marlin-v100 と sm70 用の非ゲート flashinfer を統合しました。このフォークにより、4xV100 32GB NVLINK で Dflash および Laguna S2.1 サポートを備えた Qwen3.5/3.6 モデルを実行できるようになり、最大 4000 ～ 6000 のプレフィル トークン/秒と最大 100 トークンの生成を達成します。
インターフェースを完全に再設計した Open WebUI リリース v0.11.0
Open WebUI v0.11.0 には、チャット ビュー、管理パネル、より狭い会話列、より軽いタイポグラフィ、再配置された設定など、ユーザー インターフェイス全体のビジュアルな再設計が含まれています。アップデートはコンシの向上を目的としています

LLM のオープンソース Web UI のあらゆる側面にわたる堅実性と使いやすさ。
NVIDIA が、高性能オープン エンベディング ファミリである Nemotron 3 Embed をリリース
NVIDIA は、オープン埋め込みモデルの Nemotron 3 Embed ファミリをリリースし、検索ベンチマークで最高のパフォーマンスを実現したと主張しました。これらのモデルは、RAG とエージェントのワークフローの改善を目的として、NVIDIA NIM およびセマンティック検索デモを通じて利用できます。
Skill Router: コンテキスト ウィンドウを拡大することなく、大規模なエージェント スキル ライブラリを検索するためのローカルファースト ツール
Skill Router は、エージェント スキルのメタデータを SQLite にインデックス付けし、選択されるまで完全なコンテンツをコンテキストから切り離す、ローカルファーストのツールです。 CLI、対話型シェル、MCP stdio サーバー、およびオプションの HTTP API をサポートします。
Agent Context Lens v0.2.0 では、Codex の AGENTS.md チェーン インスペクションを追加
開発者は、Codex が階層化された AGENTS.md 命令をどのように解決するかを視覚化する、MIT ライセンスの無料 CLI ツールである Agent Context Lens v0.2.0 をリリースしました。これはユーザーがどの命令がアクティブ、シャドウされているか、またはチェーンの外側にあるのかを理解するのに役立ち、大規模なリポジトリのデバッグに役立ちます。
ユーザーは、モデルとシステムのプロンプトを高速に交換できる Windows 用の llama.cpp GUI を求めています
ユーザーは、llama.cpp は正常に動作しますが、モデルとシステム プロンプトをすばやく交換するための GUI が欠けていると報告しました。 Lm-studio はこれらの機能を提供しますが、同じ設定ではメモリの問題が原因で実行が大幅に遅くなります。 Llama-swap はモデルの交換を処理しますが、システム プロンプトは処理しません。
開発者が Ollama と統合したオープンソースのローカルファーストワークフロー自動化プラットフォームを発表
開発者は、Ollama をファーストクラスのプロバイダーとして扱うオープンソース ワークフロー自動化プラットフォームの v0.11.0 をリリースしました。これにより、ローカル LLM が確定的な自動化パイプラインを強化できるようになります。このプラットフォームは、エージェントのセマンティック メモリ、ドキュメント RAG、ビジュアル ワークフロー ビルダー、条件付き分岐、ヒューマンインをサポートしています。

-the-loop 承認ノード、およびマルチエージェント オーケストレーション。
ISONGraph シリアル化フォーマットにより、GraphRAG マルチホップ精度が 40% から 80% に向上
開発者は、ナレッジ グラフを LLM にフィードするための 10 個のグラフ シリアル化形式のベンチマークを行ったところ、形式の選択だけでマルチホップの精度が 40% から 80% に変動し、トークン コストが最大 70% 変動することがわかりました。 JSON はパフォーマンスが最も悪かったものの 1 つでした。著者は、LLM の理解に最適化されたプロパティ グラフ形式である ISONGraph を構築し、MIT ライセンスの下で 6 つの言語でリリースされました。
運用アナリストが、ChatGPT に数値ソースに DATA、DERIVED、または ESTIMATE のラベルを付けるよう強制するプロンプト トリックを共有します
運用アナリストは、ChatGPT にすべての数値、日付、または名前付き数値をソースにインラインでタグ付けさせるコピーアンドペーストのプロンプト行を投稿しました: 提供されたデータからの場合は [DATA]、計算された場合 (計算が示されている) の場合は [DERIVED]、一般知識からの場合は [ESTIMATE]。このトリックは、信憑性があるように見える捏造された数値をユーザーが信用するのを防ぐことを目的としています。
コーディング エージェント ハーネスを段階的に構築する
メンバー限定の記事では、コーディング エージェント ハーネスの構築について説明し、エージェント ループ、計画、サブエージェント、サンドボックス化、メモリ、およびチェックポイント設定をカバーしています。これは、カスタム エージェントの典型的な失敗と、同じタスクにおけるクロード コードの成功を対比しています。
AMD、新しい Venice チップでサーバー CPU の収益シェア 46% を達成、AI ハードウェアの焦点を CPU オーケストレーションに移す
AMD の新しい Venice CPU により、同社は x86 サーバーでの収益シェアを 2017 年のほぼゼロから 46% にまで押し上げました。この変化は、基本的なチャットボット クエリよりも最大 1,000 倍多くのトークンを消費する可能性があるエージェント AI ワークフローにおける CPU オーケストレーションの重要性の増大を浮き彫りにしています。
ASL V6: Python AI エージェント用のオープンソース AST レッドチーム エンジンがリリースされました
ASL V6、Python AI 時代のオープンソース抽象構文ツリー (AST) レッドチーム エンジン

nts が GitHub でリリースされました。このツールを使用すると、開発者はコードの AST を操作することで AI エージェントの脆弱性を自動的に調査できるため、展開前にセキュリティの弱点を特定できます。
Aegis Router のリリース: 従量課金制のエージェント IDE 用 API ルーター
Aegis Router は、コーディング エージェント向けに設計された API ルーターとしてリリースされ、従量課金制とキャッシュ対応ルーティングを特徴としています。 OpenAI 互換のインターフェイスを提供し、繰り返されるコンテキストの再処理を回避することで、長時間実行されるエージェント IDE セッションのコストを削減することを目的としています。
コミュニティは、Qwen 由来のコーディング モデルである Kat Coder 2.5 をテストし、完全な Star Fox 風のゲームを 1 つの HTML ファイルで生成します
Reddit ユーザーは、Qwen 3.6 35B A3B から派生したモデルである Kat Coder 2.5 をテストし、バニラ Three.js を使用してスター フォックスにインスピレーションを得た宇宙船ゲームを作成するよう促しました。このモデルは、船のコントロール、複数の敵タイプ、ボス、レベルの進行、HUD、および視覚効果を備えた完全にプレイ可能なゲームを含む 1 つの HTML ファイルを生成しました。ユーザーは Q4_K_M 量子化でモデルを実行し、テストでベース モデルを常に上回るパフォーマンスを示したことがわかりました。
OMP: デバッグとサブエージェントが組み込まれたオープンソースの 55k 行 Rust コーディング エージェント
OMP は、55,000 行の Rust で書かれたオープンソースのコーディング エージェントで、組み込みのデバッグ、LSP 統合、ブラウザ自動化、メモリ、サブエージェントを備えています。完全な自己完結型開発環境を提供することで、AI IDE を超えることを目指しています。
Kitewright が AI エージェント用の 7 MB ブラウザ自動化バイナリをリリース
Kitewright は、LLM エージェントに、Node.js や Playwright のオーバーヘッドなしでナビゲーション、スクリーンショット、抽出、PDF 生成のための実際のブラウザを提供する新しい 7 MB バイナリです。フットプリントを最小限に抑えるためにコンパイル言語で書かれています。
静的 MCP により、開発者は MCP ツールを静的として公開できます

サンドボックス実行の c ファイル
Static MCP と呼ばれる新しいオープンソース プロジェクトを使用すると、開発者はモデル コンテキスト プロトコル (MCP) ツールを静的ファイルとして公開し、サンドボックス環境で実行できます。このアプローチにより、サーバーを実行する必要がなくなるため、展開が簡素化され、セキュリティが強化されます。このプロジェクトは、2026 年 7 月 27 日に Hacker News で共有されました。
AntLing がトークン効率を高める 124B ハイブリッド リニア MoE モデルである Ling-3.0-flash をリリース
AntLing は、トークンあたりの有用な作業を最大化するように設計された 124B パラメーターのハイブリッド線形専門家混合モデルである Ling-3.0-flash を発表しました。このモデルは、レイテンシ、キャッシュ ミス、エネルギー コストが多くのターンにわたって蓄積される運用エージェント ループを対象としています。これは、現実世界のエージェント ワークロードに対するトークンの効率性と持続可能なインテリジェンスを重視します。
AgentVerity が LLM エージェントのルーティングと一貫性を評価するために起動
AgentVerity は、複数ステップのワークフローにわたるルーティングの正確性と一貫性に重点を置いた、LLM ベースのエージェント用の新しい評価フレームワークを導入します。このツールは、単一の「グリーン」品質スコアでは包括的なテスト範囲や再現性を証明できないことを強調しています。
Etsy の販売者は、AI リスティング ツールには大幅な編集が必要で、価格の提案は信頼できないことに気づいた
Etsy ショップのオーナーは、商品リスト、SEO、ダイナミックな AI ツールをテストしました

[切り捨てられた]

## Original Extract

New models, tools, and framework launches in the AI/LLM world.

Hi HN, maker here. KBlip watches Reddit, HN, arXiv, GitHub, YouTube, and ~100 RSS/blog feeds for anything
AI-related, clusters same-story coverage into a single "thread" (every source linked as
evidence), and sorts it into 5 feeds: Releases, News, Social, Developments (research + hardware),
Products/Ideas, and Trending. My daily routine to catch-up on the current happenings related to AI/LLM takes a big chunk of my time and I had to do multiple hops between different products and even worse sometimes I find myself lost in a rabbit hole in any of those products. Wanted to have better visibility and have everything aggregated in one place that takes less time and effort just to catch-up which also cuts off the noise/redundancy. A few things to highlight: a daily digest per section (one-sentence TL;DR),
an audio "brief me" mode if you'd rather listen than read, and it installs as a PWA. Solo project, self-hosted on a VPS, DeepSeek + Voyage for the LLM/embedding side. No accounts, no ads, no tracking.

Releases · KBlip - The Signal Feed for the AI/LLM World K Blip Releases News Products Social Trending Developments ⌘K ··· Releases / Launches
New models, tools, and framework launches
Overview Recent headlines 10 updates Expand 📣 News — Moonshot AI releases Kimi K3 open weights: 2.8T-parameter MoE model with 91.2% on SimpleQA, permissive license [19] [1]
⚡ Capability — AI coding agent refactors 750k LOC app in 3 days with zero bugs, 31 verification passes [7]
🛠️ Tool — WISP streaming engine runs 2T+ MoE models like Kimi K3 on consumer hardware via 3-tier streaming [9]
🛠️ Tool — Krasis runtime streams Ornith-397B at Q4 on single RTX PRO 6000 Blackwell 96GB, achieving 2,354 tok/s prefill [6]
🛠️ Tool — Open WebUI v0.11.0 ships full interface redesign with chat view, admin panel, lighter typography [11]
🛠️ Tool — Community ports SGLang to V100 GPUs with custom FlashAttention for Qwen models [10]
🛠️ Tool — Minimax M3 support with Multi-Stream Attention merged into llama.cpp for local long-context processing [22]
⚡ Capability — POCKET-35B agentic model runs on CPU at 59 t/s, fits on phones with GGUF quantization [23]
🛠️ Tool — World-Model-Optimizer distills and routes models for half the cost, open-source [20]
🛠️ Tool — OpenClaude improved coding agent released on GitHub , supports cloud APIs and local models [21]
Moonshot AI releases Kimi K3 model with permissive license
Moonshot AI has released the Kimi K3 language model under a license that defines 'Model as a Service' broadly, restricting third-party inference or fine-tuning services unless aggregate revenue exceeds a threshold. The license aims to balance openness with commercial protection.
Composer v3 teased with 2.85k downloads in first hour
A teaser image for Composer v3 shows 2.85k downloads within the first hour, indicating strong early interest. The post suggests AI labs are about to have a busy day, implying multiple releases or announcements.
Vanara launches free tier for CI-verified Claude Code agent catalog
Vanara is a catalog of Claude Code agents that are verified via runnable scripts in public CI on every push, ensuring they work as claimed. The free tier is now open, allowing users to install agent packs like `npx vanara install security-pack`.
Anthropic's rapid model releases outpace competitors
Anthropic released Claude Fable 5 in June, which was pulled by the government over export controls. OpenAI then released GPT-5.6 in July, claiming it beats Fable. Two weeks later, Anthropic dropped Opus 5 at half the price of Fable, topping independent leaderboards ahead of both Fable and GPT-5.6. This marks Anthropic's fourth model in under two months, showcasing a strategy of timely releases.
↑ Updated Mon, Jul 27, 2026, 02:37 PM — Anthropic releases Opus 5, topping leaderboards after OpenAI's GPT-5.6 claim.
Cursor research department tests parallel multi-agent swarm skill
A user shares insights from Cursor's research department on agent swarms, applying similar ideologies to a new parallel multi-agent swarm skill. The skill aims to improve efficiency by coordinating multiple agents concurrently.
Krasis runtime runs Ornith-397B at Q4 on single RTX PRO 6000 Blackwell 96GB
A developer built Krasis, an MoE-focused runtime that streams large models through limited VRAM. It runs Ornith-1.0-397B at Q4 on a single RTX PRO 6000 Blackwell 96GB, achieving 2,354 tok/s prefill and ~20–24 tok/s decode by keeping experts in CPU RAM and dynamically managing VRAM residency.
AI coding agent refactors 750k LOC app in 3 days with zero bugs
A case study from AI Sovereign Labs shows an AI coding agent autonomously refactoring a 750,000-line application in three days. The agent ran 31 verification passes, corrected 201 errors, and shipped code with zero bugs, regressions, or technical debt. The project was not a from-scratch rewrite or a clone of training data.
XYZ AI Lab releases XYZ-Aquila-mini, an open-weight Deep Search agent
XYZ AI Lab released XYZ-Aquila-mini, an open-weight Deep Search agent post-trained from Qwen3.6-35B-A3B via an AI4AI pipeline. The model is a thinking model designed for bounded-exploration search tasks, with human-defined constraints and AI-driven failure diagnosis.
WISP streaming engine runs 2T+ MoE models on consumer hardware
WISP is a C/CUDA streaming engine that enables running large MoE models like Kimi K3 (2.8T) and GLM-5.2 (744B) on consumer hardware via 3-tier streaming (VRAM→RAM→NVMe SSD). It features self-organizing LRU cache, MLA attention support, and same-family speculative decoding for 2.2-2.8x throughput. Verified working on Mixtral-8x7B.
Community ports SGLang to V100 GPUs with custom FlashAttention for Qwen models
A developer forked SGLang to add support for NVIDIA V100 GPUs, writing custom TeilLang FlashAttention and integrating open-source marlin-v100 and ungated flashinfer for sm70. The fork enables running Qwen3.5/3.6 models with Dflash and Laguna S2.1 support on 4xV100 32GB NVLINK, achieving ~4000-6000 prefill tokens/s and ~100 tokens/s generation.
Open WebUI releases v0.11.0 with full interface redesign
Open WebUI v0.11.0 ships a ground-up visual redesign of the entire user interface, including chat view, admin panel, narrower conversation column, lighter typography, and rearranged settings. The update aims to improve consistency and usability across all aspects of the open-source web UI for LLMs.
NVIDIA releases Nemotron 3 Embed, a high-performance open embedding family
NVIDIA released the Nemotron 3 Embed family of open embedding models, claiming top performance on retrieval benchmarks. The models are available via NVIDIA NIM and a semantic search demo, aiming to improve RAG and agent workflows.
Skill Router: local-first tool for searching large agent skill libraries without blowing up your context window
Skill Router is a local-first tool that indexes agent skill metadata into SQLite, keeping full content out of context until selected. It supports CLI, interactive shell, MCP stdio server, and an optional HTTP API.
Agent Context Lens v0.2.0 adds AGENTS.md chain inspection for Codex
A developer released Agent Context Lens v0.2.0, a free MIT-licensed CLI tool that visualizes how Codex resolves layered AGENTS.md instructions. It helps users understand which instructions are active, shadowed, or outside the chain, aiding debugging of large repositories.
User seeks llama.cpp GUI for Windows with fast model/system prompt swapping
A user reports that llama.cpp runs well but lacks a GUI for quickly swapping models and system prompts. Lm-studio offers these features but runs significantly slower with the same settings, likely due to memory issues. Llama-swap handles model swapping but not system prompts.
Developer launches open-source local-first workflow automation platform with Ollama integration
A developer released v0.11.0 of an open-source workflow automation platform that treats Ollama as a first-class provider, enabling local LLMs to power deterministic automation pipelines. The platform supports agent semantic memory, document RAG, visual workflow builder, conditional branching, human-in-the-loop approval nodes, and multi-agent orchestration.
ISONGraph serialization format improves GraphRAG multi-hop accuracy from 40% to 80%
A developer benchmarked 10 graph serialization formats for feeding knowledge graphs to LLMs, finding that format choice alone swings multi-hop accuracy from 40% to 80% and token cost by ~70%. JSON was among the worst performers. The author built ISONGraph, a property-graph format optimized for LLM comprehension, released under MIT license in 6 languages.
Ops analyst shares prompt trick to force ChatGPT to label number sources as DATA, DERIVED, or ESTIMATE
An operations analyst posted a copy-paste prompt line that forces ChatGPT to tag every number, date, or named figure inline with its source: [DATA] if from provided data, [DERIVED] if calculated (with calculation shown), or [ESTIMATE] if from general knowledge. The trick aims to prevent users from trusting fabricated figures that appear credible.
Building a coding agent harness step by step
A member-only article walks through constructing a coding agent harness, covering the agent loop, planning, subagents, sandboxing, memory, and checkpointing. It contrasts the typical failure of custom agents with the success of Claude Code on the same tasks.
AMD hits 46% server CPU revenue share with new Venice chips, shifting AI hardware focus to CPU orchestration
AMD's new Venice CPUs have propelled the company to a 46% revenue share in x86 servers, up from near zero in 2017. This shift highlights the growing importance of CPU orchestration for agentic AI workflows, which can consume up to 1,000x more tokens than basic chatbot queries.
ASL V6: Open-source AST red-teaming engine for Python AI agents released
ASL V6, an open-source abstract syntax tree (AST) red-teaming engine for Python AI agents, has been released on GitHub. The tool allows developers to automatically probe AI agents for vulnerabilities by manipulating their code's AST, helping identify security weaknesses before deployment.
Aegis Router launches: API router for agentic IDEs with measured-energy billing
Aegis Router launched as an API router designed for coding agents, featuring measured-energy billing and cache-aware routing. It offers an OpenAI-compatible interface and aims to reduce costs for long-running agentic IDE sessions by avoiding reprocessing of repeated context.
Community tests Kat Coder 2.5, a Qwen-derived coding model, generating a full Star Fox-like game in one HTML file
A Reddit user tested Kat Coder 2.5, a model derived from Qwen 3.6 35B A3B, by prompting it to create a Star Fox-inspired spaceship game with vanilla Three.js. The model produced a single HTML file containing a fully playable game with ship controls, multiple enemy types, bosses, level progression, HUD, and visual effects. The user ran the model at Q4_K_M quantization and noted it consistently outperformed its base model in their tests.
OMP: Open-Source 55k-Line Rust Coding Agent with Built-in Debugging and Subagents
OMP is an open-source coding agent written in 55,000 lines of Rust, featuring built-in debugging, LSP integration, browser automation, memory, and subagents. It aims to surpass AI IDEs by offering a complete, self-contained development environment.
Kitewright releases 7 MB browser automation binary for AI agents
Kitewright is a new 7 MB binary that gives LLM agents a real browser for navigation, screenshots, extraction, and PDF generation without the overhead of Node.js and Playwright. It is written in a compiled language for minimal footprint.
Static MCP lets developers publish MCP tools as static files with sandboxed execution
A new open-source project called Static MCP allows developers to publish Model Context Protocol (MCP) tools as static files, which are then executed in a sandboxed environment. This approach simplifies deployment and enhances security by eliminating the need for a running server. The project was shared on Hacker News on July 27, 2026.
AntLing releases Ling-3.0-flash, a 124B hybrid-linear MoE model for token efficiency
AntLing announced Ling-3.0-flash, a 124B-parameter hybrid-linear mixture-of-experts model designed to maximize useful work per token. The model targets production agent loops where latency, cache misses, and energy costs accumulate across many turns. It emphasizes token efficiency and sustainable intelligence for real-world agentic workloads.
AgentVerity launches to evaluate LLM agent routing and consistency
AgentVerity introduces a new evaluation framework for LLM-based agents, focusing on routing correctness and consistency across multi-step workflows. The tool highlights that a single 'green' quality score does not prove comprehensive test coverage or repeatability.
Etsy seller finds AI listing tools need heavy editing, pricing suggestions unreliable
An Etsy shop owner tested AI tools for product listings, SEO, and dynamic

[truncated]
