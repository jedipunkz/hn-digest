---
source: "https://github.com/NetX-lab/Frontier"
hn_url: "https://news.ycombinator.com/item?id=48487049"
title: "Frontier: A Discrete-Event Simulator for Modern LLM Serving"
article_title: "GitHub - NetX-lab/Frontier: Frontier: A Discrete-Event Simulator for Modern LLM Serving · GitHub"
author: "matt_d"
captured_at: "2026-06-11T07:49:32Z"
capture_tool: "hn-digest"
hn_id: 48487049
score: 2
comments: 0
posted_at: "2026-06-11T06:44:01Z"
tags:
  - hacker-news
  - translated
---

# Frontier: A Discrete-Event Simulator for Modern LLM Serving

- HN: [48487049](https://news.ycombinator.com/item?id=48487049)
- Source: [github.com](https://github.com/NetX-lab/Frontier)
- Score: 2
- Comments: 0
- Posted: 2026-06-11T06:44:01Z

## Translation

タイトル: Frontier: 最新の LLM サービスのための離散イベント シミュレーター
記事のタイトル: GitHub - NetX-lab/Frontier: Frontier: 最新 LLM サービングのための離散イベント シミュレーター · GitHub
説明: Frontier: 最新の LLM サービスのための離散イベント シミュレーター - NetX-lab/Frontier

記事本文:
GitHub - NetX-lab/Frontier: Frontier: 最新の LLM サービスのための離散イベント シミュレーター · GitHub
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
ネットエックスラボ
/
フロンティア
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
主要支店

タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
6 コミット 6 コミット データ データ docs docs 例 例 figs figs フロンティア フロンティア ログ ログ 出力 出力テスト テスト .gitignore .gitignore .gitmodules .gitmodules AGENTS.md AGENTS.md ライセンス ライセンス README.md README.md 環境.yml 環境.yml 環境プロファイリング.yml environment_profiling.yml pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
最新の LLM サービスのための離散イベント シミュレーター
📍[2026/06] 初期バージョンがリリースされ、同じ場所に配置されたサービスと最新の最適化がサポートされました。細分化された配信のサポートは間もなく利用可能になります。乞うご期待！
Frontier は、最新の LLM サービス用の離散イベント シミュレーターです。これは、複雑な並列処理、ランタイム最適化、スパース モデル アーキテクチャ (MoE)、およびステートフル ワークロード (推論エージェント、RL ロールアウト) を組み合わせたサービス提供システム向けに構築されています。現時点では vLLM をシミュレートしていますが、他のサービス エンジンもすぐに含める予定です。
Frontier は、研究者やエンジニアが、GPU クラスターに繰り返しデプロイする時間と経済的コストを発生させることなく、サービス提供システムの設計とトレードオフをより深く理解できるように支援します。
コロケーション サービング : このブランチは、モノリシック コロケーション サービング クラスターをモデル化します。 PDD と AFD のサポートは、後の公開リリースで予定されています。
最新のランタイム最適化 : Frontier は、CUDA グラフ、投機的デコード / MTP、プレフィックス キャッシュ、量子化、チャンク プレフィル、階層型キャッシュなどの運用技術をスケジューラー、バッチ、エンジン ループの一部として取り込みます。これらの最適化によりバッチの形状、メモリの状態、リクエストごとの進行状況が変化するため、Frontier はそれらを単純な高速化要因ではなく実行時の動作としてモデル化します。
忠実度 : Frontier は、校正されたオペレーター、通信、転送、および KV キャッシュ メモリを組み合わせます。

モデルを使用して、シミュレーション結果を展開の決定に役立てます。これにより、ユーザーは SLA 制約の下で構成を比較し、現場外で大規模な GPU スケールの設計空間を探索し、粗い平均ケース モデルによって歪められる結論を避けることができます。
細分化されたアーキテクチャは意図的にこのリリースには含まれていませんが、間もなく利用可能になる予定です。
シミュレーション: Frontier は、プリコンパイルされたプロファイリング データベースを活用して、CPU のみのマシンでの E2E 実行を完全にサポートします。
プロファイリング: 新しいハードウェアまたはソフトウェア スタックでオペレーター レベルのパフォーマンス メトリクスを収集するために、Frontier プロファイリング モジュールを実行する場合のみ、少なくとも 1 つの GPU が必要です。
Frontier は、大規模な GPU クラスターで直接実行するとコストが高かったり、時間がかかったりするような仮定の研究用に設計されています。現在の草案では 4 つの使用例を示しています。このパブリック ブランチは現在、コロケーション パスを公開しています。細分化されたユースケースはロードマップの例です。
リリース パッケージをインストールし、リポジトリ ルートから追加機能をテストします。
python -m pip install -e ' .[テスト] '
PYTHONPATH= $PWD PYTHONDONTWRITEBYTECODE=1 pytest testing/unit/test_open_source_release_arch_guard.py -q -p no:cacheprovider
現在のリリース向けのコロケーション サンプルはシミュレーション モードごとに分割されており、デフォルトではワンクリックでスモークを実行するための式ベースの分析バックエンドが使用されます。
例/アーキテクチャ/コロケーション/run_all.sh
例/アーキテクチャ/コロケーション/オンライン/dense_model_basic_online.sh
例/architecture/co-location/online/moe_model_basic_online.sh
例/アーキテクチャ/コロケーション/オンライン/ Thinking_mode_basic_online.sh
例/
§── 建築/
│ └── コロケーション/
│ §── run_all.sh
│ §── オフライン/
│ └── オンライン/
§── 備品/
━── プロファイリング/
通信バックエンド
コロケーション サンプル スイートのデフォルトは、lightwe

ワンクリックでスモークを実行するための ight 式ベースのバックエンド:
--cc_backend_config_type 分析用
ASTRA-Sim にヒントを得たトポロジ モデルは、引き続き CLI の直接実験に利用できます。
--cc_backend_config_type astra_sim_analytical
collective_sim はオプションであり、 --cc_backend_config_typecollective_sim を明示的に選択した場合にのみ使用されます。
オプションのターゲットランタイムバックエンドを有効にするには:
git submodule update --init --recursivefrontier/cc_backend/backends/collective-sim
cd フロンティア/cc_backend/backends/collective-sim/sim
make -j " $( nproc ) "
ビルド後、 --cc_backend_config_typecollective_sim を選択する前に、frontier/cc_backend/backends/collective-sim/sim/datacenter/htsim_ndp が存在することを確認します。
ホスト C++ ランタイムが GLIBC または GLIBCXX の不一致を報告した場合は、ターゲット ランタイムでオプションのバックエンドを再構築します。
make -B -j " $( nproc ) "
環境とDocker
標準シミュレータ環境の場合は、environment.yml からインストールするか、編集可能な pip インストールを使用します。 vllm と flashinfer は CUDA と Python のバージョンに影響される可能性があるため、プロファイリング ラッパーは専用の環境を使用します。
conda env create -f 環境プロファイリング.yml
conda はフロンティアプロファイリングをアクティブ化します
バージョンの互換性を確認していない限り、既存の環境にプロファイリングの依存関係をやみくもにインストールしないでください。
実稼働 Docker ユーザーはパブリック イメージから開始できます。
docker pull fengyicheng/frontier-env
docker run --rm --gpus all --shm-size 16g \
--tmpfs /workspace/frontier/outputs \
--tmpfs /workspace/frontier/cache \
-v " $PWD " :/workspace/frontier \
-w /ワークスペース/フロンティア \
fengyicheng/フロンティア環境 bash
計画
私たちは、Frontier のコア コンポーネントをコミュニティに継続的にリリースしていきます。私たちは積極的に新機能の開発を行っており、以下をサポートする予定です。
サービング エンジンの統合 : SGLang および TensorRT-LLM のサポート

ラメワークス。
高度なキャッシュ: 階層キャッシュ機能の高度なランタイム バックエンドとして tair-kvcache をサポートします。
高度なモデルのサポート : DeepSeek-V4、Kimi 2.5 などの最先端のモデルのサポートを拡張しました。
シミュレーション アクセラレーション : シミュレーション アクセラレーション メカニズムをさらに導入します。
フロンティアは主にヴィドゥルの頂上に建てられています。次の優れたシステムは、ランタイム バックエンドとして参照または採用されています。コミュニティに貢献してくださった開発者の皆様に心より感謝いたします。
研究でこのリポジトリを使用する場合は、論文を引用してください (論文の公開後に引用の詳細が追加されます)。
@article { feng2026frontier 、
title = { フロンティア: 包括的で正確な LLM 推論シミュレーションに向けて } ,
著者 = { Feng、Yicheng と Tan、Xin と Deng、Yangtao と Jiang、Yimin と Zhu、Yibo と Xu、Hong } ,
ジャーナル = { arXiv プレプリント arXiv:2605.21312 } ,
年 = { 2026 }
}
お問い合わせ
ご質問がある場合は、Yicheng Feng ( yichengfeng@link.cuhk.edu.hk ) またはHong Xu ( hongxu@cuhk.edu.hk ) に電子メールを送信してください。
フィードバック、問題点、PR などは大歓迎です。ぜひご貢献をお願いいたします。
このプロジェクトは MIT ライセンスに基づいてライセンスされています。詳細については、LICENSE ファイルを参照してください。
Frontier: 最新の LLM サービスのための離散イベント シミュレータ
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
4
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Frontier: A Discrete-Event Simulator for Modern LLM Serving - NetX-lab/Frontier

GitHub - NetX-lab/Frontier: Frontier: A Discrete-Event Simulator for Modern LLM Serving · GitHub
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
NetX-lab
/
Frontier
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
6 Commits 6 Commits data data docs docs examples examples figs figs frontier frontier logs logs outputs outputs tests tests .gitignore .gitignore .gitmodules .gitmodules AGENTS.md AGENTS.md LICENSE LICENSE README.md README.md environment.yml environment.yml environment_profiling.yml environment_profiling.yml pyproject.toml pyproject.toml View all files Repository files navigation
A Discrete-Event Simulator for Modern LLM Serving
📍[2026/06] Initial version released, with support for co-located serving and modern optimizations. Support for disaggregated serving will be available soon. Stay tuned!
Frontier is a discrete-event simulator for modern LLM serving. It is built for serving systems that combine complex parallelisms, runtime optimizations, sparse model architectures (MoE), and stateful workloads (reasoning agents, RL rollouts). It simulates vLLM at the moment but we plan to include other serving engines soon.
Frontier helps researchers and engineers better understand serving system designs and tradeoffs without the time and financial costs of repeatedly deploying on GPU clusters.
Co-located Serving : This branch models a monolithic co-location serving cluster. PDD and AFD support is planned for a later public release.
Modern Runtime Optimizations : Frontier captures production techniques such as CUDA Graph, speculative decoding / MTP, prefix caching, quantization, chunked prefill, and hierarchical caching as part of the scheduler-batch-engine loop. These optimizations change batch shape, memory state, and per-request progress, so Frontier models them as runtime behavior rather than simple speedup factors.
Fidelity : Frontier combines calibrated operator, communication, transfer, and KV-cache memory models to make simulation results useful for deployment decisions. This helps users compare configurations under SLA constraints, explore large GPU-scale design spaces ex-situ, and avoid conclusions that would be distorted by coarse average-case models.
Disaggregated architectures are intentionally not included in this release but will be available soon.
Simulation: Frontier supports E2E execution entirely on CPU-only machines, leveraging pre-compiled profiling databases.
Profiling: At least 1 GPU is required exclusively for running the Frontier profiling module to collect operator-level performance metrics on new hardware or software stacks.
Frontier is designed for what-if studies that would be expensive or slow to run directly on large GPU clusters. The current paper draft demonstrates four use cases. This public branch currently exposes the co-location path; disaggregated use cases are roadmap examples.
Install the release package and test extras from the repository root:
python -m pip install -e ' .[test] '
PYTHONPATH= $PWD PYTHONDONTWRITEBYTECODE=1 pytest tests/unit/test_open_source_release_arch_guard.py -q -p no:cacheprovider
Current release-facing co-location examples are split by simulation mode and default to the formula-based analytical backend for one-click smoke runs:
examples/architecture/co-location/run_all.sh
examples/architecture/co-location/online/dense_model_basic_online.sh
examples/architecture/co-location/online/moe_model_basic_online.sh
examples/architecture/co-location/online/thinking_mode_basic_online.sh
examples/
├── architecture/
│ └── co-location/
│ ├── run_all.sh
│ ├── offline/
│ └── online/
├── fixtures/
└── profiling/
Communication Backend
The co-location example suite defaults to the lightweight formula-based backend for one-click smoke runs:
--cc_backend_config_type analytical
The ASTRA-Sim-inspired topology model remains available for direct CLI experiments:
--cc_backend_config_type astra_sim_analytical
collective_sim is optional and is used only when you explicitly select --cc_backend_config_type collective_sim .
To enable the optional target-runtime backend:
git submodule update --init --recursive frontier/cc_backend/backends/collective-sim
cd frontier/cc_backend/backends/collective-sim/sim
make -j " $( nproc ) "
After building, verify that frontier/cc_backend/backends/collective-sim/sim/datacenter/htsim_ndp exists before selecting --cc_backend_config_type collective_sim .
If your host C++ runtime reports a GLIBC or GLIBCXX mismatch, rebuild the optional backend in the target runtime:
make -B -j " $( nproc ) "
Environment and Docker
For the standard simulator environment, install from environment.yml or use editable pip installation. Profiling wrappers use a dedicated environment because vllm and flashinfer can be sensitive to CUDA and Python versions:
conda env create -f environment_profiling.yml
conda activate frontier-profiling
Do not blindly install profiling dependencies into an existing environment unless you have confirmed version compatibility.
Production Docker users can start from the public image:
docker pull fengyicheng/frontier-env
docker run --rm --gpus all --shm-size 16g \
--tmpfs /workspace/frontier/outputs \
--tmpfs /workspace/frontier/cache \
-v " $PWD " :/workspace/frontier \
-w /workspace/frontier \
fengyicheng/frontier-env bash
Plan
We will continuously release Frontier's core components to the community. We are actively developing new features and plan to support:
Serving Engines Integration : Support for SGLang and TensorRT-LLM frameworks.
Advanced Caching : Support for tair-kvcache as an advanced runtime backend for the Hierarchical Caching feature.
Advanced Model Support : Expanded support for state-of-the-art models such as DeepSeek-V4, Kimi 2.5, and more.
Simulation Acceleration : Introduce more simulation acceleration mechanisms.
Frontier is mainly built on top of Vidur. The following great systems have been referenced or adapted as runtime backends. We sincerely thank all the developers for their contributions to the community!
If you use this repo in your research, please cite our paper (citation details will be added after the paper is released):
@article { feng2026frontier ,
title = { Frontier: Towards Comprehensive and Accurate LLM Inference Simulation } ,
author = { Feng, Yicheng and Tan, Xin and Deng, Yangtao and Jiang, Yimin and Zhu, Yibo and Xu, Hong } ,
journal = { arXiv preprint arXiv:2605.21312 } ,
year = { 2026 }
}
Contact
Email Yicheng Feng ( yichengfeng@link.cuhk.edu.hk ) or Hong Xu ( hongxu@cuhk.edu.hk ) if you have any questions.
Feedback, issues, and PRs are highly welcome! We would love your contribution!
This project is licensed under the MIT license - see the LICENSE file for details.
Frontier: A Discrete-Event Simulator for Modern LLM Serving
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
4
forks
Report repository
Releases
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
