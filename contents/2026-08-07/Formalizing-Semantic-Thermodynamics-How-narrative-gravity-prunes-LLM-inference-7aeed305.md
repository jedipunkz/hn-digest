---
source: "https://github.com/tauansloboda/semantic-thermodynamics"
hn_url: "https://news.ycombinator.com/item?id=49216790"
title: "Formalizing Semantic Thermodynamics: How narrative gravity prunes LLM inference"
article_title: "GitHub - tauansloboda/semantic-thermodynamics · GitHub"
author: "tauanvinicius"
captured_at: "2026-08-07T22:25:18Z"
capture_tool: "hn-digest"
hn_id: 49216790
score: 1
comments: 0
posted_at: "2026-08-07T22:05:19Z"
tags:
  - hacker-news
  - translated
---

# Formalizing Semantic Thermodynamics: How narrative gravity prunes LLM inference

- HN: [49216790](https://news.ycombinator.com/item?id=49216790)
- Source: [github.com](https://github.com/tauansloboda/semantic-thermodynamics)
- Score: 1
- Comments: 0
- Posted: 2026-08-07T22:05:19Z

## Translation

タイトル: 意味論的熱力学の定式化: 物語重力が LLM 推論をどのようにプルーニングするか
記事のタイトル: GitHub - tauansloboda/semantic-thermodynamics · GitHub
説明: GitHub でアカウントを作成して、tauansloboda/semantic-thermodynamics の開発に貢献します。

記事本文:
GitHub - タウアンスロボダ/セマンティック熱力学 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
タウアンスロボダ
/
意味論的熱力学
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
3 コミット 3 コミット DOCUMENTO FUNDACIONAL_ Termodinâmica Semântica v1.docx DOCUMENTO FUNDACIONAL_ Termodinâmica Semântica v1.docx README.md README.md Semantic_Thermodynamics_Whitepaper.pdf Semantic_Thermodynamics_Whitepaper.pdf

位相空間の崩壊_ 推論枝刈りにおける操作熱力学量としての物語.pdf 位相空間の崩壊_ 推論枝刈りにおける操作熱力学量としての物語.pdf Experimento_um.py Experimento_um.py Experimento_um_resultados.csv Experimento_um_resultados.csv Experimento_zero.py Experimento_zero.py Experimento_zero_resultados.csv Experimento_zero_resultados.csv すべてのファイルを表示 リポジトリ ファイルのナビゲーション
更新 (2026 年 8 月): 位相空間崩壊、構造摩擦、エントロピー比例則をカバーする完全な形式化ペーパーを追加しました。
意味論的熱力学: LLM エントロピーの最小化
このリポジトリには、セマンティック熱力学フレームワークの基礎的なホワイトペーパー、実行スクリプト、生の経験データが含まれています。
大規模言語モデルはベイジアン推論エンジンとして動作します。セマンティック エントロピーが高い環境では、無限の位相空間をマッピングするために余分な計算を消費します。ペルソナ、目的論的ベクトル、破壊的枝刈りの正確な公式である「ナラティブ グラビティ」を適用することで、モデルを強制的に決定論的な測地線にし、コストと待ち時間を大幅に削減できます。
実験結果 (実験ゼロ)
Formula v2.0 を使用して標準のデータ抽出プロンプトを再構築すると、次のことがわかりました。
完了トークンが 79.29% 削減されました。
システム遅延が 60.73% 短縮されました (3.4 秒から 1.3 秒)。
Semantic_Thermodynamics_Whitepaper.pdf : 完全な理論的枠組みとエントロピー比例の法則 ( $\Lambda$ )。
Experimento_zero.py : トークンとレイテンシの崩壊のベンチマークに使用される非同期 Python スクリプト。
Experimento_um.py : 「構造摩擦」と最適なセマンティック勾配をマッピングするスクリプト。
*.csv : OpenAI API からの生のテレメトリ データが実行されます。
API キーをエクスポートします:export OPENAI_API_KEY

==============================
python Experimento_zero.py を実行して、レイテンシの低下を確認します。
著者: タウアン ヴィニシウス グアヒバ スロボダ
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to tauansloboda/semantic-thermodynamics development by creating an account on GitHub.

GitHub - tauansloboda/semantic-thermodynamics · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
tauansloboda
/
semantic-thermodynamics
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits DOCUMENTO FUNDACIONAL_ Termodinâmica Semântica v1.docx DOCUMENTO FUNDACIONAL_ Termodinâmica Semântica v1.docx README.md README.md Semantic_Thermodynamics_Whitepaper.pdf Semantic_Thermodynamics_Whitepaper.pdf The Collapse of Phase Space_ Narrative as an Operational Thermodynamic Quantity in Inference Pruning.pdf The Collapse of Phase Space_ Narrative as an Operational Thermodynamic Quantity in Inference Pruning.pdf experimento_um.py experimento_um.py experimento_um_resultados.csv experimento_um_resultados.csv experimento_zero.py experimento_zero.py experimento_zero_resultados.csv experimento_zero_resultados.csv View all files Repository files navigation
UPDATE (August 2026): Added the full formalization paper covering Phase Space Collapse, Structural Friction, and the Entropic Proportionality Law.
Semantic Thermodynamics: LLM Entropy Minimization
This repository contains the foundational whitepaper, execution scripts, and raw empirical data for the Semantic Thermodynamics framework.
Large Language Models operate as Bayesian inference engines. In environments with high semantic entropy, they expend excess compute mapping infinite phase spaces. By applying "Narrative Gravity"—a precise formula of persona, teleological vectors, and destructive pruning—we can force the model into a deterministic geodesic, drastically reducing cost and latency.
Empirical Results (Experiment Zero)
By restructuring a standard data-extraction prompt using the Formula v2.0 , we observed:
79.29% reduction in completion tokens.
60.73% reduction in system latency (from 3.4s to 1.3s).
Semantic_Thermodynamics_Whitepaper.pdf : The complete theoretical framework and the Law of Entropic Proportionality ( $\Lambda$ ).
experimento_zero.py : The async Python script used to benchmark the token and latency collapse.
experimento_um.py : The script mapping the "Structural Friction" and the optimal semantic gradient.
*.csv : Raw telemetry data from the OpenAI API runs.
Export your API key: export OPENAI_API_KEY="sk-..."
Run python experimento_zero.py and watch the latency drop.
Author: Tauan Vinicius Guahyba Sloboda
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
