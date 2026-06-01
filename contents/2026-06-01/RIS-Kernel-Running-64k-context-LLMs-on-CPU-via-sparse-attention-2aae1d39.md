---
source: "https://github.com/santosardr/riskernel"
hn_url: "https://news.ycombinator.com/item?id=48348450"
title: "RIS-Kernel: Running 64k context LLMs on CPU via sparse attention"
article_title: "GitHub - santosardr/riskernel: RIS-Kernel: A Model-Agnostic Architecture for Long-Context LLM Inference via Sparse Attention · GitHub"
author: "santosardr"
captured_at: "2026-06-01T00:26:24Z"
capture_tool: "hn-digest"
hn_id: 48348450
score: 2
comments: 0
posted_at: "2026-05-31T18:44:54Z"
tags:
  - hacker-news
  - translated
---

# RIS-Kernel: Running 64k context LLMs on CPU via sparse attention

- HN: [48348450](https://news.ycombinator.com/item?id=48348450)
- Source: [github.com](https://github.com/santosardr/riskernel)
- Score: 2
- Comments: 0
- Posted: 2026-05-31T18:44:54Z

## Translation

タイトル: RIS-Kernel: スパース アテンションを介して CPU 上で 64k コンテキスト LLM を実行する
記事のタイトル: GitHub - santosardr/riskernel: RIS-Kernel: スパース アテンションによるロングコンテキスト LLM 推論のためのモデルに依存しないアーキテクチャ · GitHub
説明: RIS-Kernel: スパース アテンションによるロングコンテキスト LLM 推論のためのモデルに依存しないアーキテクチャ - santosardr/riskernel

記事本文:
GitHub - santosardr/riskernel: RIS-Kernel: スパース アテンションによるロングコンテキスト LLM 推論のためのモデルに依存しないアーキテクチャ · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Spark インテリジェントなアプリを構築してデプロイする
GitHub モデルのプロンプトの管理と比較
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
サントサルドル
/
危険人物
公共
通知
ch にサインインする必要があります

アンジュ通知設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット コード コード データ データ ライセンス ライセンス README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
RIS-Kernel: スパース アテンションによるロングコンテキスト LLM 推論のためのモデルに依存しないアーキテクチャ
このリポジトリには、高速化されていない汎用の CPU ハードウェア上で大規模なコンテキスト ウィンドウ (64k+ トークン) を実行するシステム レベルのスパース アテンション推論エンジンである RIS-Kernel の公式実装が含まれています。
大規模な言語モデルでの完全なセルフアテンションは $O(N^2)$ にまで拡張され、長いコンテキストのドキュメント分析は 65,536 トークンに制限され、高価な GPU クラスターが必要になります。 Reduced Interaction Sampling (RIS) 推論エンジンは、モデルに依存しないアーキテクチャとしてこの制約に対処します。 RIS は、重みを変更せずに、コモディティ メモリの制限内に収まる疎な確率幾何学を使用して、セルフ アテンションの複雑さを $O(N \log N)$ に軽減します。 2 つのレジームにわたって Qwen2-1.5B-Instruct の RIS を検証します。 32,768 個のトークン (ネイティブの高密度注意が上限となる) での制御された評価では、密度 1% およびアンサンブル シード 70 個の RIS-Stochastic は 75.00% の精度を達成し、ネイティブの高密度ベースライン (71.88%) を上回りましたが、密度 5% およびシード 10 個の RIS-Stochastic はそれに一致しました (71.88%)。これは、まばらな注意が正則化機能として機能することを示しています。複数のシードにわたる低密度 (1%) ではシーケンス レベルのノイズが除去されますが、高密度 (5%) では気が散るノイズが再導入されます。最も厳しい予算の下で、RIS-Structural はわずか 10 シードで 1% の密度で 68.75% の精度に達し、ゼロコンテキストの下限 (59.38%) と比較してコンテキスト ギャップの 75% を回復します。 65,536 トークンでは、集中的な注意がメモリ不足障害を引き起こすため、RIS は yie

lds の検索は、ゼロコンテキストの下限 (51.56%) よりも最大 14.06 パーセントポイント向上しました。すべての評価は、高速化されていない汎用の CPU サーバー (16 ～ 128 GB の RAM) で実行され、ロングコンテキスト LLM 推論が GPU アクセラレーションなしの標準的な学術用ハードウェアで実行可能であることを示しています。
RIS-Kernel は、実行時に注意喚起を傍受するモデルに依存しないレイヤーとして機能します。 Reduced Interaction Sampling (RIS) を実装することで、標準の Transformer の $O(N^2)$ メモリと計算ボトルネックをバイパスします。
Qwen2-1.5B を概念実証 (PoC) として利用します。 RIS がコンパクトなモデルで検索を安定させ、誘導できることを実証することは、このアーキテクチャが厳しいパラメータ制約下でもコンテキストの一貫性を維持し、より大規模なアーキテクチャに自然に拡張できることを証明します。
⚠️ ハードウェアの免責事項とパフォーマンス
この実装は CPU のみの実行用に最適化されており、汎用学術マシン (標準ワークステーションや部門サーバーなど) でのロングコンテキストの実験が可能です。
RAM 要件: 安定した 65,536 トークン推論セッションには、~100GB+ RAM が必要です。
CPU パフォーマンス:
プレフィル: 65,000 トークンの場合は約 50 分 (1 回限りのコスト、その後キャッシュ)。
生成: トークンあたり最大 5 秒。
GPU 注: CUDA サポートは実験的です。 GPU で実行すると、プリフィル/生成時間が大幅に短縮されますが、高い VRAM が必要になります。
🛠️ フォルダー構造とコンポーネント
リポジトリは、ローカルで実行することも、再現可能な Code Ocean カプセルとして実行することもできるように構造化されています。
code/ : すべての実行スクリプト、エントリ ポイント、および視覚化モジュール。
code/scripts/ris_attention.py: Reduced Interaction Sampling スパース ジオメトリのコア実装。
code/scripts/inference_ris_v3.py: デュアルハッシュ キャッシュと PFUS を利用した高性能 CPU バウンド推論エンジン。
code/scripts/benchmark/ : 実行スクリプト

コンテキスト ウィンドウと密度全体でスイープを実行します。
code/article/fig/ : プロットを生成するための視覚化スクリプト。
data/ : コンテキストドキュメント ( genppi.txt 、 aom.txt など) のマウントされたローカル ディレクトリ。
results/ : ベンチマークと生成された数値が出力されるディレクトリ。
python3 -m venv venv
ソース venv/bin/activate
pip install torch --index-url https://download.pytorch.org/whl/cpu
pip install -r コード/スクリプト/要件-cpu.txt
2. コンテキストを準備する
コンテキスト記事はすでに data/ フォルダーの下にプリロードされています。独自の PDF を使用したい場合は、原稿リポジトリの extract_pdf.py ユーティリティを使用して、PDF をクリーンなテキスト ブロックに前処理できます。
Python を使用して推論エンジンを起動します。
PYTHONPATH=コード/スクリプト Python コード/スクリプト/inference_ris_v3.py \
--model_class qwen2 \
--ウィンドウ 65536 \
--context_files データ/genppi.txt \
--密度 0.05 \
--n_seeds 1
主な議論:
--window : トークン単位のコンテキスト ウィンドウ サイズ。
--density : アクティブな注意の密度割合 (例: 1% の場合は 0.01、5% の場合は 0.05)。
--n_seeds : アンサンブルする確率的に投影されたマスクの数。
--save_graph : アテンション トポロジを .dot ファイルにエクスポートします。
--save_graph フラグを使用して、スパース アテンション トポロジをエクスポートできます。結果の .dot ファイルを Graphviz または Gephi で開き、アテンション検索マップを検査します。
このコードは、科学的な透明性と再現性を実現するために、MIT ライセンスに基づいて利用できます。この著作物を使用する場合は、プレプリントを引用してください。
@misc { santos2026riskernel 、
著者 = { サントス、アンダーソン R. } 、
title = { RIS-Kernel: スパース アテンションによるロングコンテキスト LLM 推論のためのモデルに依存しないアーキテクチャ } ,
年 = { 2026 } 、
パブリッシャー = { Zenodo } 、
ドイ = { 10.5281/zenodo.20476759 } 、
URL = { https://doi.org/10.5281/zenodo.20476759 }
}
について
RIS-Kernel: モデルに依存しない Ar

スパース アテンションによるロングコンテキスト LLM 推論のアーキテクチャ
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

RIS-Kernel: A Model-Agnostic Architecture for Long-Context LLM Inference via Sparse Attention - santosardr/riskernel

GitHub - santosardr/riskernel: RIS-Kernel: A Model-Agnostic Architecture for Long-Context LLM Inference via Sparse Attention · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Spark Build and deploy intelligent apps
GitHub Models Manage and compare prompts
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
santosardr
/
riskernel
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits code code data data LICENSE LICENSE README.md README.md View all files Repository files navigation
RIS-Kernel: A Model-Agnostic Architecture for Long-Context LLM Inference via Sparse Attention
This repository contains the official implementation of RIS-Kernel , a systems-level sparse attention inference engine that runs massive context windows (64k+ tokens) on commodity, unaccelerated CPU hardware.
Full self-attention in large language models scales as $O(N^2)$ , limiting long-context document analysis to 65,536 tokens and requiring costly GPU clusters. The Reduced Interaction Sampling (RIS) inference engine addresses this constraint as a model-agnostic architecture. Without modifying weights, RIS reduces self-attention complexity to $O(N \log N)$ using sparse stochastic geometry that fits within commodity memory limits. We validate RIS on Qwen2-1.5B-Instruct across two regimes. In controlled evaluations at 32,768 tokens (where native dense attention serves as the upper bound), RIS-Stochastic at 1% density and 70 ensemble seeds achieves 75.00% accuracy, outperforming the native dense baseline (71.88%), while RIS-Stochastic at 5% density and 10 seeds matches it (71.88%). This demonstrates that sparse attention acts as a regularizer: low density (1%) over multiple seeds filters out sequence-level noise, whereas higher density (5%) reintroduces distractor noise. Under the tightest budget, RIS-Structural reaches 68.75% accuracy at 1% density with just 10 seeds, recovering 75% of the contextual gap relative to the zero-context floor (59.38%). At 65,536 tokens, where dense attention triggers out-of-memory faults, RIS yields retrieval gains of up to 14.06 percentage points over the zero-context floor (51.56%). All evaluations run on commodity, unaccelerated CPU servers (16–128 GB of RAM), demonstrating that long-context LLM inference is feasible on standard academic hardware without GPU acceleration.
RIS-Kernel acts as a model-agnostic layer that intercepts attention calls at runtime. By implementing Reduced Interaction Sampling (RIS), it bypasses the $O(N^2)$ memory and compute bottleneck of standard Transformers.
We utilize Qwen2-1.5B as a Proof of Concept (PoC). Demonstrating that RIS can stabilize and guide retrieval in a compact model proves that the architecture maintains contextual coherence even under severe parameter constraints, scaling naturally to larger architectures.
⚠️ Hardware Disclaimer & Performance
This implementation is optimized for CPU-only execution to enable long-context experiments on commodity academic machines (like standard workstations or departmental servers).
RAM Requirements : ~100GB+ RAM is required for stable 65,536 token inference sessions.
CPU Performance :
Prefill : ~50 minutes for 65k tokens (one-time cost, cached thereafter).
Generation : ~5 seconds per token.
GPU Note : CUDA support is experimental. Running on GPU will drastically reduce prefill/generation times but requires high VRAM.
🛠️ Folder Structure & Components
The repository is structured to run both locally and as a reproducible Code Ocean capsule:
code/ : All execution scripts, entry points, and visualization modules.
code/scripts/ris_attention.py: Core implementation of the Reduced Interaction Sampling sparse geometry.
code/scripts/inference_ris_v3.py: High-performance CPU-bound inference engine utilizing dual-hash caching and PFUS.
code/scripts/benchmark/ : Execution scripts for running sweeps across context windows and densities.
code/article/fig/ : Visualization scripts for generating plots.
data/ : Mounted/local directory for context documents ( genppi.txt , aom.txt , etc.).
results/ : Directory where benchmarks and generated figures are outputted.
python3 -m venv venv
source venv/bin/activate
pip install torch --index-url https://download.pytorch.org/whl/cpu
pip install -r code/scripts/requirements-cpu.txt
2. Prepare Context
The context articles are already pre-loaded under the data/ folder. If you wish to use your own PDFs, you can use the extract_pdf.py utility from the manuscript repository to preprocess them into clean text blocks.
Launch the inference engine using python:
PYTHONPATH=code/scripts python code/scripts/inference_ris_v3.py \
--model_class qwen2 \
--window 65536 \
--context_files data/genppi.txt \
--density 0.05 \
--n_seeds 1
Key Arguments:
--window : Context window size in tokens.
--density : Active attention density fraction (e.g., 0.01 for 1%, 0.05 for 5%).
--n_seeds : Number of stochastically projected masks to ensemble.
--save_graph : Exports the attention topology to a .dot file.
You can export the sparse attention topology with the --save_graph flag. Open the resulting .dot file in Graphviz or Gephi to inspect the attention retrieval maps.
The code is available for scientific transparency and reproducibility under the MIT License. If you use this work, please cite the preprint:
@misc { santos2026riskernel ,
author = { Santos, Anderson R. } ,
title = { RIS-Kernel: A Model-Agnostic Architecture for Long-Context LLM Inference via Sparse Attention } ,
year = { 2026 } ,
publisher = { Zenodo } ,
doi = { 10.5281/zenodo.20476759 } ,
url = { https://doi.org/10.5281/zenodo.20476759 }
}
About
RIS-Kernel: A Model-Agnostic Architecture for Long-Context LLM Inference via Sparse Attention
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
