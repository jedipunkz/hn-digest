---
source: "https://github.com/bjoern-janson/resolution-horizon"
hn_url: "https://news.ycombinator.com/item?id=48969902"
title: "Resolution Horizon – Finding the mathematical limit where AI overfits to noise"
article_title: "GitHub - bjoern-janson/resolution-horizon: A framework for measuring the limits of invariant discovery under bounded observation. · GitHub"
author: "bjoern_janson"
captured_at: "2026-07-19T17:50:00Z"
capture_tool: "hn-digest"
hn_id: 48969902
score: 1
comments: 0
posted_at: "2026-07-19T17:15:40Z"
tags:
  - hacker-news
  - translated
---

# Resolution Horizon – Finding the mathematical limit where AI overfits to noise

- HN: [48969902](https://news.ycombinator.com/item?id=48969902)
- Source: [github.com](https://github.com/bjoern-janson/resolution-horizon)
- Score: 1
- Comments: 0
- Posted: 2026-07-19T17:15:40Z

## Translation

タイトル: Resolution Horizon – AI がノイズにオーバーフィットする数学的限界を見つける
記事のタイトル: GitHub - bjoern-janson/resolution-horizon: 限定された観察の下で不変発見の限界を測定するためのフレームワーク。 · GitHub
説明: 制限された観察の下で不変発見の限界を測定するためのフレームワーク。 - ビョルン・ジャンソン/解像度ホライズン

記事本文:
GitHub - bjoern-janson/resolution-horizon: 限定された観察の下で不変発見の限界を測定するためのフレームワーク。 · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
ビョルン・ヤンソン
/
解像度-ホライズン
公共
通知
通知を変更するにはサインインする必要があります

設定
追加のナビゲーション オプション
コード
ビョルン・ジャンソン/解像度ホライズン
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
17 コミット 17 コミット docs docs .gitignore .gitignore ライセンス ライセンス README.md README.md demo.ipynbデモ.ipynb Experiment.py Experiment.py Instrument.py Instrument.py test_suite.py test_suite.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
制限された観察の下での不変発見の限界を測定する。
Resolution Horizo​​n は、動的システムの有限でノイズの多い観測から数学的構造をどの程度回復できるかを研究するための実験的研究フレームワークです。
このフレームワークは、発見可能な構造をシステムの固有の特性として扱うのではなく、構造の回復が力学、観察者、利用可能な計算リソースによって共同で決定されることを提案しています。
中心的な仮説は、すべての観測プロセスには測定可能な分解能の範囲 ( $k^*$ ) があり、それを超えると追加の分析では真の不変量が明らかにならず、代わりに観測の不確実性がモデル化され始める構造的複雑さの有限レベルがあります。
このプロジェクトは、微分幾何学、力学システム、統計的推定、情報理論を組み合わせて、この境界を研究するための統一された経験的フレームワークを作成します。
動的システムの有限観測を考慮すると、構造深度を増加させると、最初は不変量の発見が改善されますが、臨界深度を超えると、推定の不確実性が支配的になり、実験的に測定可能な分解能の範囲が生成されます。
主な経験的量は情報効率交換率 $\eta(k)$ であり、圧縮値と構造的曖昧さのバランスを取ります。
$$\eta(k) = \frac{\Delta L(k)}{C_G(k)} = \frac{L(\mathcal{D}_N \mid \text{Baseline}) - L\left(\mathcal{D}_

N \ \Big\vert\ \frac{\text{Lie}_k(\hat{F})}{\hat{G}^{(k)}}\right)}{L(\theta_G) + \log N\left(\delta, \ \hat{G}^{(k)}(\epsilon_k) \cap B_R, \ d_{C^{k+1}}\right)}$$
最適な発見範囲は次のように定義されます。
$$k^* = \arg\max_k \eta(k)$$
地平線は、データと環境の制約によって形成される創発的な応答曲面です。
$$k^* = f\left(N, \sigma, E_{\mathrm{cov}}, \lambda_{\min}(\Gamma)\right)$$
これらの方向性の予測は、プロジェクトの主要な反証可能な主張を形成します。
3 つの体制段階のフットプリント
双子の情報曲線間の相互作用により、代数深度 $k$ にわたる独特の結合された幾何学的シグネチャが生成されます。
レジーム I (ディスカバリー): $\Delta L(k) \uparrow$ および $C_G(k) \downarrow$ 。構造の圧縮は効率的に上昇し、座標の曖昧さは縮小します。
レジーム II (Horizo​​n): $\eta(k)$ がグローバル最大値に達します。これは最適な構造解像度境界 ( $k = k^*$ ) です。
レジーム III (誤解決): $\Delta L(k) \downarrow$ および $C_G(k) \uparrow$ 。この機器は等価関係自体をオーバーフィットし、高次微分ノイズを偽の背景対称として解析します。
§──instrument.py # 定量測定コアとメトリクス計算機
§── Experiment.py # 正則化バイパスプロトコルと因果関係介入
§── test_suite.py # 自動トポロジ検証ハーネス
└── README.md # プロジェクトの仕様と展開ガイド
哲学
Resolution Horizon は、科学的発見そのものを測定可能な計算プロセスとして扱います。尋ねる代わりに:
どのような数学的構造が存在するのでしょうか?
フレームワークは次のことを尋ねます。
有限データ、有限精度、有限計算を使用する有界オブザーバーによって回復可能な数学的構造は何ですか?
ライセンス
MITライセンスに基づいてリリースされています。
について
lを測定するためのフレームワーク

限定された観察の下での不変発見の模倣。
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

A framework for measuring the limits of invariant discovery under bounded observation. - bjoern-janson/resolution-horizon

GitHub - bjoern-janson/resolution-horizon: A framework for measuring the limits of invariant discovery under bounded observation. · GitHub
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
bjoern-janson
/
resolution-horizon
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
bjoern-janson/resolution-horizon
main Branches Tags Go to file Code Open more actions menu Folders and files
17 Commits 17 Commits docs docs .gitignore .gitignore LICENSE LICENSE README.md README.md demo.ipynb demo.ipynb experiment.py experiment.py instrument.py instrument.py test_suite.py test_suite.py View all files Repository files navigation
Measuring the limits of invariant discovery under bounded observation.
Resolution Horizon is an experimental research framework for studying how much mathematical structure can be recovered from finite, noisy observations of dynamical systems.
Rather than treating discoverable structure as an intrinsic property of a system, the framework proposes that structural recovery is jointly determined by the dynamics, the observer, and the available computational resources.
The central hypothesis is that every observation process possesses a measurable resolution horizon ( $k^*$ ): a finite level of structural complexity beyond which additional analysis no longer reveals genuine invariants and instead begins to model observational uncertainty.
The project combines differential geometry, dynamical systems, statistical estimation, and information theory into a unified empirical framework for studying this boundary.
Given finite observations of a dynamical system, increasing structural depth initially improves invariant discovery, but beyond a critical depth, estimation uncertainty dominates, producing an experimentally measurable resolution horizon.
The primary empirical quantity is the information efficiency exchange rate $\eta(k)$ , balancing compression value against structural ambiguity:
$$\eta(k) = \frac{\Delta L(k)}{C_G(k)} = \frac{L(\mathcal{D}_N \mid \text{Baseline}) - L\left(\mathcal{D}_N \ \Big\vert\ \frac{\text{Lie}_k(\hat{F})}{\hat{G}^{(k)}}\right)}{L(\theta_G) + \log N\left(\delta, \ \hat{G}^{(k)}(\epsilon_k) \cap B_R, \ d_{C^{k+1}}\right)}$$
The optimal discovery horizon is defined as:
$$k^* = \arg\max_k \eta(k)$$
The horizon is an emergent response surface shaped by data and environmental constraints:
$$k^* = f\left(N, \sigma, E_{\mathrm{cov}}, \lambda_{\min}(\Gamma)\right)$$
These directional predictions form the primary falsifiable claims of the project.
The Three-Regime Phase Footprint
The interaction between the twin informational curves produces a distinctive coupled geometric signature across algebraic depth $k$ :
Regime I (Discovery): $\Delta L(k) \uparrow$ and $C_G(k) \downarrow$ . Structural compression rises efficiently while coordinate ambiguity contracts.
Regime II (Horizon): $\eta(k)$ achieves its global maximum. This is the optimal structural resolution boundary ( $k = k^*$ ).
Regime III (Misresolution): $\Delta L(k) \downarrow$ and $C_G(k) \uparrow$ . The instrument overfits the equivalence relation itself, parsing high-order derivative noise as false background symmetry.
├── instrument.py # Quantitative measurement core & metric calculators
├── experiment.py # Regularization Bypass protocol & causal interventions
├── test_suite.py # Automated topological verification harness
└── README.md # Project specification and deployment guide
Philosophy
Resolution Horizon treats scientific discovery itself as a measurable computational process. Instead of asking:
What mathematical structure exists?
The framework asks:
What mathematical structure is recoverable by a bounded observer with finite data, finite precision, and finite computation?
License
Released under the MIT License.
About
A framework for measuring the limits of invariant discovery under bounded observation.
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
