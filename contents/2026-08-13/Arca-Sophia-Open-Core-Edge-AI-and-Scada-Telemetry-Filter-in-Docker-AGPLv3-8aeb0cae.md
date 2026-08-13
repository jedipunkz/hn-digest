---
source: "https://github.com/Nodo9/arca-sophia-open-core"
hn_url: "https://news.ycombinator.com/item?id=49291382"
title: "Arca Sophia Open-Core–Edge AI and Scada Telemetry Filter in Docker (AGPLv3)"
article_title: "GitHub - Nodo9/arca-sophia-open-core: Conector industrial PLC y filtrado sintrópico para telemetría SCADA en Edge Computing (Open-Core AGPLv3 / SILIC-ETHIC). · GitHub"
author: "Nodo9"
captured_at: "2026-08-13T20:31:20Z"
capture_tool: "hn-digest"
hn_id: 49291382
score: 1
comments: 0
posted_at: "2026-08-13T20:26:12Z"
tags:
  - hacker-news
  - translated
---

# Arca Sophia Open-Core–Edge AI and Scada Telemetry Filter in Docker (AGPLv3)

- HN: [49291382](https://news.ycombinator.com/item?id=49291382)
- Source: [github.com](https://github.com/Nodo9/arca-sophia-open-core)
- Score: 1
- Comments: 0
- Posted: 2026-08-13T20:26:12Z

## Translation

タイトル: Arca Sophia Open-Core–Edge AI and Scada Telemetry Filter in Docker (AGPLv3)
記事のタイトル: GitHub - Nodo9/arca-sophia-open-core: エッジ コンピューティングにおける SCADA テレメトリ用の PLC 産業用コネクタとシントロピック フィルタリング (オープンコア AGPLv3 / SILIC-ETHIC)。 · GitHub
説明: エッジ コンピューティング (オープンコア AGPLv3 / SILIC-ETHIC) における SCADA テレメトリ用の産業用 PLC コネクタとシントロピック フィルタリング。 - Node9/アルカソフィア-オープンコア

記事本文:
GitHub - Nodo9/arca-sophia-open-core: コネクター産業用 PLC とエッジ コンピューティング (オープンコア AGPLv3 / SILIC-ETHIC) によるテレメトリ解析用の SCADA フィルター。 · GitHub
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
のど9
/
アルカソフィアオープンコア
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット Simulador_PLC Simulador_PLC .gitignore .gitignore ETHICS.md ETHICS.md LICENSE.txt LICENSE.txt README.md 読む

ME.md docker-compose.yml docker-compose.yml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Arca Sophia エコシステム — PLC 産業用コネクタおよび LCIM エンジン (オープンコア)
エッジ コンピューティングの分散推論を使用した、SCADA/PLC 環境向けのリアルタイムのシントロピック フィルタリングおよび異常軽減ソリューション。
📐 システムアーキテクチャ (オープンコア)
このプロジェクトは、産業間の相互運用性とコア ロジックの保護を保証するために、二重層モデルの下で動作します。
パブリック レイヤー (PLC_Simulator & Orchestration): 産業用テレメトリの取得とプロセス サイクルのシミュレーションのためのオープン インターフェイス (AGPLv3)。
コア層 (ics-core-sophia): INT4 での低エントロピーのセマンティック最適化を備えたローカル推論エンジン。
🚀 迅速な導入 (ローカル環境)
# パブリックリポジトリのクローンを作成する
git clone [https://github.com/nodo-raiz/bunker-codigo-public.git](https://github.com/nodo-raiz/bunker-codigo-public.git)
cd バンカーコード-パブリック
# コンテナインフラストラクチャを立ち上げる
docker-compose up -d --build
について
エッジ コンピューティング (オープンコア AGPLv3 / SILIC-ETHIC) における SCADA テレメトリ用の産業用 PLC コネクタとシントロピック フィルタリング。
github.com/Nodo9/arca-sophia-open-core リソース
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Conector industrial PLC y filtrado sintrópico para telemetría SCADA en Edge Computing (Open-Core AGPLv3 / SILIC-ETHIC). - Nodo9/arca-sophia-open-core

GitHub - Nodo9/arca-sophia-open-core: Conector industrial PLC y filtrado sintrópico para telemetría SCADA en Edge Computing (Open-Core AGPLv3 / SILIC-ETHIC). · GitHub
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
Nodo9
/
arca-sophia-open-core
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit Simulador_PLC Simulador_PLC .gitignore .gitignore ETHICS.md ETHICS.md LICENSE.txt LICENSE.txt README.md README.md docker-compose.yml docker-compose.yml View all files Repository files navigation
Ecosistema Arca Sophia — Conector Industrial PLC & Motor LCIM (Open-Core)
Solución de filtrado sintrópico y mitigación de anomalías en tiempo real para entornos SCADA/PLC mediante inferencia distribuida en Edge Computing.
📐 Arquitectura del Sistema (Open-Core)
El proyecto opera bajo un modelo de doble capa para garantizar la interoperabilidad industrial y la protección de la lógica central:
Capa Pública ( Simulador_PLC & Orquestación) : Interfaz Abierta (AGPLv3) para adquisición de telemetría industrial y simulación de ciclos de proceso.
Capa Núcleo ( ics-core-sophia ) : Motor de inferencia local con optimización semántica de baja entropía en INT4.
🚀 Despliegue Rápido (Entorno Local)
# Clonar repositorio público
git clone [https://github.com/nodo-raiz/bunker-codigo-public.git](https://github.com/nodo-raiz/bunker-codigo-public.git)
cd bunker-codigo-public
# Levantar infraestructura de contenedores
docker-compose up -d --build
About
Conector industrial PLC y filtrado sintrópico para telemetría SCADA en Edge Computing (Open-Core AGPLv3 / SILIC-ETHIC).
github.com/Nodo9/arca-sophia-open-core Resources
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
