---
source: "https://github.com/healthchainai/HealthChain"
hn_url: "https://news.ycombinator.com/item?id=48764545"
title: "Show HN: HealthChain – Python SDK to connect AI models to live EHR systems"
article_title: "GitHub - healthchainai/HealthChain: Python SDK for healthcare AI: connect models to live EHR systems, skip the integration tax 💫 🏥 · GitHub"
author: "adjks"
captured_at: "2026-07-02T17:52:22Z"
capture_tool: "hn-digest"
hn_id: 48764545
score: 1
comments: 0
posted_at: "2026-07-02T17:18:58Z"
tags:
  - hacker-news
  - translated
---

# Show HN: HealthChain – Python SDK to connect AI models to live EHR systems

- HN: [48764545](https://news.ycombinator.com/item?id=48764545)
- Source: [github.com](https://github.com/healthchainai/HealthChain)
- Score: 1
- Comments: 0
- Posted: 2026-07-02T17:18:58Z

## Translation

タイトル: Show HN: HealthChain – AI モデルをライブ EHR システムに接続するための Python SDK
記事のタイトル: GitHub - healthchainai/HealthChain: ヘルスケア AI 用の Python SDK: モデルをライブ EHR システムに接続し、統合税をスキップします 💫 🏥 · GitHub
説明: ヘルスケア AI 用の Python SDK: モデルをライブ EHR システムに接続し、統合税をスキップします 💫 🏥 - healthchainai/HealthChain

記事本文:
GitHub - healthchainai/HealthChain: ヘルスケア AI 用の Python SDK: モデルをライブ EHR システムに接続し、統合税をスキップします 💫 🏥 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
ヘルスチェナイ
/
ヘルスチェーン

公共
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
165 コミット 165 コミット .github .github クックブック クックブック dev-templates dev-templates docs docs healthchain healthchain ノートブック ノートブック リソース リソース スクリプト スクリプト テスト テスト .dockerignore .dockerignore .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml CITATION.cff CITATION.cff CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md COTRIBUTING.md COOKBOOK_CONTRIBUTORS.md COOKBOOK_CONTRIBUTORS.md Dockerfile Dockerfile LICENSE LICENSE MAINTAINERS.md MAINTAINERS.md README.md README.md SECURITY.md SECURITY.md mkdocs.yml mkdocs.yml pyproject.toml pyproject.toml 要件-doc.txt 要件-doc.txt uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ヘルスケア AI 用のオープンソース Python SDK
モデルは簡単な部分ですが、それが本番環境に到達するのを妨げる統合は難しい部分です。 HealthChain はこれを処理し、タイプセーフな FHIR リソース、リアルタイム EHR 接続、本番環境に対応した展開など、信頼できるツールをモデルとエージェントに提供します。したがって、構築したものはデモの外でも保持されます。
pip インストールヘルスチェーン
クイックスタート
# FHIR ゲートウェイ プロジェクトのスキャフォールディング
healthchain 新しい my-app -t fhir-gateway
cd 私のアプリ
# ローカルで実行
ヘルスチェーンサーブ
app.py を編集してモデルを追加し、healthchain.yaml を編集してコンプライアンス、セキュリティ、デプロイ設定を構成します。
すべてのコマンドについては、CLI リファレンスを参照してください。
AI 開発者や研究者がヘルスケア AI を出荷するための最も迅速な方法。必要なものすべてがすぐに利用でき、ユーザーに合わせて拡張できるように構築されています。
あらゆる深刻なヘルスケア A

私は、同じ統合インフラストラクチャをゼロから構築するプロジェクトです。ロジスティック回帰、70B パラメータ モデル、エージェント ワークフローのいずれを導入する場合でも、トレーニング済みモデルと実際の臨床システムとの間にある壁は同じです。つまり、実際の FHIR API、マルチサイト導入、監査可能なガバナンスです。既製のソリューションは存在せず、AI と医療プロトコルの両方を理解するエンジニアは希少で、維持するのが困難です。
HealthChain はその複雑さを処理するため、実際に重要なこと、つまりモデルと患者に集中できます。
リアルタイム向けに最適化 - 古いデータのエクスポートではなく、ライブ FHIR API と統合ポイントに接続します。
自動検証 - タイプセーフな FHIR モデルで医療データの破損を防止
ネイティブ LLM、エージェント、ML サポート - 任意のモデルまたはエージェント (LLM、LangChain、scikit-learn) を接続し、結果を FHIR として出力します
既存のスタックと連携 - FastAPI、LangChain、HuggingFace、spaCy と統合
本番環境に対応した基盤 - Docker 化された導入、構成可能なセキュリティ、NHS および HIPAA 環境向けに構築されたアーキテクチャ
TLDR AI ニュースレターで紹介されました (90 万人以上の開発者)
Epic とのオープンソース統合に関して Medplum によって紹介されました
NHS Python オープンソース カンファレンスで発表 (トークを見る)
NHS AI 導入経験に基づいて構築 – 起源のストーリーを読む
🤝 パートナーシップとプロダクションでの使用
あなたの製品または組織のために HealthChain を検討していますか?統合、パイロット、コラボレーションについて話し合うために連絡するか、Discord に参加してコミュニティとつながりましょう。
ヘルスチェーンから。ゲートウェイインポート HealthChainAPI、FHIRGateway
ヘルスチェーンから。もふもふr4b インポート患者
# ヘルスケアアプリケーションを作成する
app = HealthChainAPI (title = "マルチ EHR 患者データ")
# 複数の FHIR ソースに接続する
fhir = FHIRGゲートウェイ()
もふもふadd_source ( "エピック" , "fhir://fhi

r.epic.com/r4?client_id=epic_client_id" )
もふもふadd_source ( "cerner" , "fhir://fhir.cerner.com/r4?client_id=cerner_client_id" )
@ ふぃる 。集合体 (患者)
def enrich_patient_data (id : str 、source : str ) -> 患者 :
"""接続されている EHR から患者データを取得し、AI の機能強化を追加します"""
バンドル = fhir 。検索(
患者さん、
{ "_id" : id },
ソース、
add_provenance = True 、
来歴_tag = "ai強化" ,
)
返品バンドル
アプリ。 register_gateway ( fhir )
# 入手可能場所: GET /fhir/transform/Patient/123?source=epic
# 入手可能場所: GET /fhir/transform/Patient/123?source=cerner
if __name__ == "__main__" :
アプリ。実行 (ポート = 8888)
パイプラインの構築 [ ドキュメント ]
ヘルスチェーンから。パイプラインのインポート パイプライン
ヘルスチェーンから。パイプライン。コンポーネント。統合インポート SpacyNLP
ヘルスチェーンから。 ioインポートドキュメント
# 医療 NLP パイプラインを作成する
nlp_pipeline = パイプライン [ドキュメント]()
nlp_パイプライン 。 add_node ( SpacyNLP . from_model_id ( "en_core_web_sm" ))
nlp = nlp_pipeline 。ビルド()
doc = Document (「患者は高血圧と糖尿病を患っています。」)
結果 = nlp (ドキュメント)
spacy_doc = 結果 。 nlp 。 get_spacy_doc ()
print ( f"Entities: { [( ent . text , ent . label_ ) for ent in spacy_doc . ents ] } " )
print ( f"FHIR 条件: { result . fhir . issue_list } " ) # FHIR バンドルに自動変換
🛣️ 私たちが目指しているもの
🔒 本番環境のセキュリティとコンプライアンス — API 認証、監査ログ、NHS/HIPAA 導入向けの構成可能なセキュリティ
📋 構成としてのガバナンス — healthchain.yaml のファーストクラスの展開アーティファクトとしての臨床安全性、データ アクセス契約、コンプライアンス標準
🔌 より深い EHR 接続 - より多くの FHIR ソース、ライブ データ パターン、パイロット展開による実際の統合例
📊 可観測性 — 導入電話番号

医療システムのメトリと監査証跡
🤖 臨床 AI エージェント用のツールキット - エージェントが臨床データを確実に操作できるようにする、型付き FHIR ツール、検証ループ、評価
HealthChain は、次世代のヘルスケア開発者によって、次世代のヘルスケア開発者のために構築されています。モデルを遡及データからライブ システムに移行する研究者や、何かを出荷する前に FHIR の学習に何か月も費やしたくない AI 開発者です。最も優れた貢献は、実際の問題に直面し、それについて具体的な意見を持っている人から来ます。
医療データや研究データを扱いますか?クックブックを寄稿する — ユースケースを持参してください。私が個人的にサポートします。
ガイドラインについては CONTRIBUTING.md をお読みください
技術的な質問とアイデア → GitHub ディスカッション
パイロットとパートナーシップ → 電子メール
このプロジェクトは、HL7 とボストン小児病院によって開発された fhir.resources と CDS Hooks 標準に基づいて構築されています。
© 2024–2026 ドット実装 AI。 HealthChain は、 dotimplement ai によって管理されているオープンソース プロジェクトです。
ヘルスケア AI 用の Python SDK: モデルをライブ EHR システムに接続し、統合税をスキップします 💫 🏥
healthchainai.github.io/HealthChain/
トピックス
Apache-2.0ライセンス
行動規範
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
39
フォーク
レポートリポジトリ
リリース
26
ヘルスチェーン v0.15.0
最新の
2026 年 6 月 23 日
+ 25 リリース
このプロジェクトのスポンサーになる
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Python SDK for healthcare AI: connect models to live EHR systems, skip the integration tax 💫 🏥 - healthchainai/HealthChain

GitHub - healthchainai/HealthChain: Python SDK for healthcare AI: connect models to live EHR systems, skip the integration tax 💫 🏥 · GitHub
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
healthchainai
/
HealthChain
Public
Uh oh!
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
165 Commits 165 Commits .github .github cookbook cookbook dev-templates dev-templates docs docs healthchain healthchain notebooks notebooks resources resources scripts scripts tests tests .dockerignore .dockerignore .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml CITATION.cff CITATION.cff CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md COOKBOOK_CONTRIBUTORS.md COOKBOOK_CONTRIBUTORS.md Dockerfile Dockerfile LICENSE LICENSE MAINTAINERS.md MAINTAINERS.md README.md README.md SECURITY.md SECURITY.md mkdocs.yml mkdocs.yml pyproject.toml pyproject.toml requirements-doc.txt requirements-doc.txt uv.lock uv.lock View all files Repository files navigation
Open-Source Python SDK for Healthcare AI
The model is the easy part — the integration that blocks it from ever reaching production is the hard part. HealthChain handles it, giving your models and agents tools they can trust: type-safe FHIR resources , real-time EHR connectivity , and production-ready deployment . So what you build holds up outside the demo.
pip install healthchain
Quick Start
# Scaffold a FHIR Gateway project
healthchain new my-app -t fhir-gateway
cd my-app
# Run locally
healthchain serve
Edit app.py to add your model, and healthchain.yaml to configure compliance, security, and deployment settings.
See the CLI reference for all commands.
The quickest way for AI developers and researchers to ship healthcare AI — everything you need out of the box, built to scale with you.
Every serious healthcare AI project builds the same integration infrastructure from scratch. Whether you're deploying a logistic regression, a 70B-parameter model, or an agentic workflow, the wall between a trained model and a live clinical system is the same: real FHIR APIs, multi-site deployments, auditable governance. No off-the-shelf solution exists, and engineers who understand both AI and healthcare protocols are scarce and hard to retain.
HealthChain handles that complexity so you can focus on what actually matters: the model and the patient.
Optimized for real-time - Connect to live FHIR APIs and integration points instead of stale data exports
Automatic validation - Type-safe FHIR models prevent broken healthcare data
Native LLM, agent & ML support - Wire up any model or agent — LLMs, LangChain, scikit-learn — and output results as FHIR
Works with your existing stack - Integrates with FastAPI, LangChain, HuggingFace, and spaCy
Production-ready foundations - Dockerized deployment, configurable security, and an architecture built for NHS and HIPAA environments
Featured in TLDR AI Newsletter (900K+ developers)
Featured by Medplum for open source integration with Epic
Presented at NHS Python Open Source Conference ( watch talk )
Built from NHS AI deployment experience – read the origin story
🤝 Partnerships & Production Use
Exploring HealthChain for your product or organization? Get in touch to discuss integrations, pilots, or collaborations, or join our Discord to connect with the community.
from healthchain . gateway import HealthChainAPI , FHIRGateway
from healthchain . fhir . r4b import Patient
# Create healthcare application
app = HealthChainAPI ( title = "Multi-EHR Patient Data" )
# Connect to multiple FHIR sources
fhir = FHIRGateway ()
fhir . add_source ( "epic" , "fhir://fhir.epic.com/r4?client_id=epic_client_id" )
fhir . add_source ( "cerner" , "fhir://fhir.cerner.com/r4?client_id=cerner_client_id" )
@ fhir . aggregate ( Patient )
def enrich_patient_data ( id : str , source : str ) -> Patient :
"""Get patient data from any connected EHR and add AI enhancements"""
bundle = fhir . search (
Patient ,
{ "_id" : id },
source ,
add_provenance = True ,
provenance_tag = "ai-enhanced" ,
)
return bundle
app . register_gateway ( fhir )
# Available at: GET /fhir/transform/Patient/123?source=epic
# Available at: GET /fhir/transform/Patient/123?source=cerner
if __name__ == "__main__" :
app . run ( port = 8888 )
Building a Pipeline [ Docs ]
from healthchain . pipeline import Pipeline
from healthchain . pipeline . components . integrations import SpacyNLP
from healthchain . io import Document
# Create medical NLP pipeline
nlp_pipeline = Pipeline [ Document ]()
nlp_pipeline . add_node ( SpacyNLP . from_model_id ( "en_core_web_sm" ))
nlp = nlp_pipeline . build ()
doc = Document ( "Patient presents with hypertension and diabetes." )
result = nlp ( doc )
spacy_doc = result . nlp . get_spacy_doc ()
print ( f"Entities: { [( ent . text , ent . label_ ) for ent in spacy_doc . ents ] } " )
print ( f"FHIR conditions: { result . fhir . problem_list } " ) # Auto-converted to FHIR Bundle
🛣️ What we're building towards
🔒 Production security and compliance — API authentication, audit logging, and configurable security for NHS/HIPAA deployments
📋 Governance as config — clinical safety, data access agreements, and compliance standards as a first-class deployment artifact in healthchain.yaml
🔌 Deeper EHR connectivity — more FHIR sources, live data patterns, and real-world integration examples from pilot deployments
📊 Observability — deployment telemetry and audit trails for healthcare systems
🤖 A toolkit for clinical AI agents — typed FHIR tools, validation loops, and evals so agents can work with clinical data reliably
HealthChain is built by and for the next generation of healthcare developers — researchers moving models from retrospective data into live systems, AI developers who don't want to spend months learning FHIR before they can ship anything. The best contributions come from people who have hit a real problem and have something specific to say about it.
Working with healthcare or research data? Contribute a cookbook — bring your use case, I'll personally support you through it
Read CONTRIBUTING.md for guidelines
Technical questions and ideas → GitHub Discussions
Pilots and partnerships → email
This project builds on fhir.resources and CDS Hooks standards developed by HL7 and Boston Children's Hospital .
© 2024–2026 dotimplement ai. HealthChain is an open source project maintained by dotimplement ai .
Python SDK for healthcare AI: connect models to live EHR systems, skip the integration tax 💫 🏥
healthchainai.github.io/HealthChain/
Topics
Apache-2.0 license
Code of conduct
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
39
forks
Report repository
Releases
26
HealthChain v0.15.0
Latest
Jun 23, 2026
+ 25 releases
Sponsor this project
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
