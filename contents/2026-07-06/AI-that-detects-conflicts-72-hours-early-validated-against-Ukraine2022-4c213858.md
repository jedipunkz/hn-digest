---
source: "https://github.com/eshanth23/sentinel"
hn_url: "https://news.ycombinator.com/item?id=48807525"
title: "AI that detects conflicts 72 hours early(validated against Ukraine2022)"
article_title: "GitHub - eshanth23/sentinel: SENTINEL - AI-powered global threat intelligence system for SCSP Hackathon 2026. SENTINEL is a real-time conflict prediction system that detects threats before they escalate by fusing 5 live intelligence signals into a single threat score (0-100). · GitHub"
author: "eshanth22"
captured_at: "2026-07-06T17:47:32Z"
capture_tool: "hn-digest"
hn_id: 48807525
score: 2
comments: 0
posted_at: "2026-07-06T17:07:27Z"
tags:
  - hacker-news
  - translated
---

# AI that detects conflicts 72 hours early(validated against Ukraine2022)

- HN: [48807525](https://news.ycombinator.com/item?id=48807525)
- Source: [github.com](https://github.com/eshanth23/sentinel)
- Score: 2
- Comments: 0
- Posted: 2026-07-06T17:07:27Z

## Translation

タイトル: 紛争を72時間前に検知するAI(ウクライナ2022に対して検証)
記事のタイトル: GitHub - eshanth23/sentinel: SENTINEL - SCSP ハッカソン 2026 向け AI を活用したグローバル脅威インテリジェンス システム。SENTINEL は、5 つのライブ インテリジェンス シグナルを 1 つの脅威スコア (0 ～ 100) に融合することで、脅威が拡大する前に脅威を検出するリアルタイム競合予測システムです。 · GitHub
説明: SENTINEL - SCSP ハッカソン 2026 向けの AI を活用したグローバル脅威インテリジェンス システム。SENTINEL は、5 つのライブ インテリジェンス シグナルを 1 つの脅威スコア (0 ～ 100) に融合することで、脅威が拡大する前に脅威を検出するリアルタイム競合予測システムです。 - GitHub - eshanth23/sentinel: SENTINEL - AI 搭載
[切り捨てられた]

記事本文:
GitHub - eshanth23/sentinel: SENTINEL - SCSP ハッカソン 2026 用の AI を活用したグローバル脅威インテリジェンス システム。SENTINEL は、5 つのライブ インテリジェンス シグナルを 1 つの脅威スコア (0 ～ 100) に融合することで、脅威が拡大する前に脅威を検出するリアルタイム競合予測システムです。 · GitHub
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
別のタブまたはウィンドウでアカウントを切り替えました。参照先へリロード

セッションをしてください。
アラートを閉じる
{{ メッセージ }}
エシャンス23
/
番兵
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
37 コミット 37 コミット バックエンド バックエンド フロントエンド フロントエンド .gitignore .gitignore DEMO_INSTRUCTIONS.txt DEMO_INSTRUCTIONS.txt README.md README.md 要件.txt 要件.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
SENTINEL — AI を活用したグローバル脅威インテリジェンス システム
ウクライナ侵攻を発生の72時間前に検知していたであろうリアルタイム紛争予測システム。
SCSP 国家安全保障ハッカソン 2026 用に構築 (ウォーゲーミング トラック)
デモ: https://github.com/user-attachments/assets/72deac0a-56df-4e25-acbd-76c621f63021
諜報機関は紛争が始まってから対応します。 2022 年のウクライナ侵攻は世界を驚かせました。信号が存在しなかったからではなく、信号をリアルタイムで融合させる人がいなかったからです。
意思決定者には、エスカレーションの前に脅威を検出する早期警告システムが必要です。
SENTINEL は、5 つの独立したデータ ソースを単一の 0 ～ 100 の脅威スコアに統合し、30 分ごとに更新します。
結果: 複数のシグナルが一致する場合 → エスカレーション前の早期警告
SENTINEL は 4 つの主要な競合に対してテストされました。
10 以上の監視地域を備えたインタラクティブな世界地図
色分けされた脅威レベル (NORMAL → ELEVATED → HIGH → CRITICAL)
30分ごとのリアルタイム更新
Conflict Brief Generator — 5 秒未満で CIA スタイルの諜報評価 (Groq + Llama 3)
AI Advisor チャットボット — 戦略的な質問をし、ライブの脅威コンテキストで回答を得る
主張検証エンジン — 5 つの情報源を相互参照することで軍事主張を事実確認します
3 つの確率的な将来経路: エスカレーション / 外交 / スタンドオフ

f
現在の信号に基づく尤度の割合
意思決定者に推奨されるアクション
過去の紛争が信号ごとにどのようにエスカレートしたかを視覚化する
4 つのシナリオ: 2022 年のウクライナ、1999 年のカルギル、1990 年の湾岸戦争、2024 年のイスラエルとイラン
SENTINEL がアラートを発したであろう正確な瞬間を観察してください
機首方位の視覚化による航空機追跡
65,000 のソースからのニュースの集約
各データソースの寄与を示す信号の内訳
観客
ユースケース
🏛️ 防衛と情報
危機対応チームへの早期警告
🤝 人道団体
移動イベントの 72 時間前にリソースを事前に配置
📰 ジャーナリストとOSINT
リアルタイムで申し立てを検証し、紛争が発展するにつれて紛争を追跡します
🎓 紛争研究者
予測モデルを検証し、エスカレーション パターンを調査する
🏢法人
サプライチェーン混乱予測、地政学的リスク評価
🛠️ 技術スタック
バックエンド: Python 3.11 • Flask • APScheduler
フロントエンド: React 18 • Leaflet.js • Recharts • Axios
AI: Groq API (Llama 3 - 70B パラメーター)
データ: GDELT • OpenSky • USGS • Yahoo Finance • ACLED
総コスト: $0.00 (すべて無料 API)
# リポジトリのクローンを作成する
git クローン https://github.com/eshanth23/sentinel.git
CDセンチネル
# バックエンドのセットアップ
CDバックエンド
Python -m venv venv
venv \S cripts \a をアクティブ化 # Windows |ソース venv/bin/activate (Mac/Linux)
pip install -r 要件.txt
# フロントエンドのセットアップ
cd ../フロントエンド
npmインストール
センチネルの実行
CD バックエンド
venv \S cripts \a アクティブ化
Python API.py
→ http://localhost:5000 で実行
CDフロントエンド
npmスタート
→ http://localhost:3000 が開きます
CDバックエンド
Pythonスケジューラ.py
→ ライブデータでキャッシュを更新します (2 ～ 3 分)
地域
スコア
レベル
トップシグナル
🔴イスラエル-ガザ-イラン
80/100
クリティカル
ニュース: 28/30、ACLED: 17/20
🔴ミャンマー
75/100
クリティカル
ACLED: 20/20、フライト: 20/20
🟠 ウクライナ-ロシア
72/100
高
ニュース

：27/30、ACLED：18/20
🟠 台湾-中国
68/100
高
フライト: 20/20、ニュース: 22/30
更新日: 2026 年 5 月
SENTINEL は 3 層設計に従っています。
📊 レイヤ 1: データ収集 (scheduler.py)
30 分ごとまたはオンデマンドで実行されます。 10 のグローバル リージョンに対して 5 つの API を呼び出し、脅威スコア (0 ～ 100) を計算し、結果を Threat_cache.json に保存します。
目的: 生のインテリジェンス信号を収集して処理する
🔌 レイヤ 2: API サーバー (api.py)
localhost:5000 上の Flask サーバー。キャッシュされたデータをロードし、それをフロントエンドに提供します。紛争概要、チャットボット、請求検証のための AI エンドポイントを提供します。
目的: REST API 経由でデータと AI 機能を公開する
🖥️ レイヤー 3: ダッシュボード (App.js)
localhost:3000 上のアプリケーションに反応します。インタラクティブなマップ、ビジュアライゼーション、AI ツール、および履歴リプレイを表示します。バックエンド API からすべてのデータを読み取ります。
目的: 脅威の監視と分析のためのユーザー インターフェイス
データ フロー: 外部 API → スケジューラ → キャッシュ → バックエンド → フロントエンド
✅ 回復力（キャッシュによりデモの失敗を防止）
✅ 費用対効果が高い (無料枠のまま)
✅ スケーラブル (レイヤーを個別に展開)
データフロー:
外部 API → スケジューラ → キャッシュ → バックエンド → フロントエンド
✅ 復元力 — キャッシュによりデモ中の API エラーを防止します
✅ 高速 — UI インタラクションの遅延ゼロ
✅ 費用対効果が高い — API 呼び出しを最小限に抑え、無料枠を維持します
✅ スケーラブル — 各レイヤーを個別に導入可能
API
適用範囲
更新頻度
コスト
グデルト
65,000 のニュースソース、100 以上の言語
リアルタイム
無料
オープンスカイ
世界中で 4,000 以上の ADS-B レシーバー
リアルタイム
無料
USGS
全球地震センサー (マグニチュード 2.5+)
リアルタイム
無料
ヤフーファイナンス
防衛請負業者トップ 5
毎日
無料
ACLED
50か国以上、200人以上の研究者
毎週
無料
グロク
ラマ 3 (70B パラメータ)
オンデマンド
無料 (14.4K/日)
🎓 構築ストーリー
タイムライン: 7 日間 (2026 年 4 月 17 日から 24 日まで)
過去の経験

ce: コーディングの知識がゼロ
学習: ビルド中にドキュメントから React + Flask を取得する
ハッカソン: SCSP 国家安全保障ハッカソン 2026 (ウォーゲーミング トラック)
🏆 SENTINEL の違い
既存のシステム
センチネル
❌ 単一信号 (ニュースまたは衛星)
✅ マルチシグナルフュージョン (5 ソース)
❌ 高価（数百万ドル）
✅ 永久無料 ($0.00)
❌ ブラックボックスAI
✅ 説明可能 (信号の内訳を示す)
❌ リアクティブ (イベント後)
✅ 予測 (72 時間ウィンドウ)
❌ 独自の
✅ オープンソース
🚀 将来のロードマップ
衛星画像の統合 (Sentinel Hub API)
しきい値を超えた場合の SMS/電子メール アラート
50 以上の監視対象地域に拡張
歴史データベース (20 年以上の紛争)
機械学習による重みの最適化
変化検出アルゴリズム (部隊増強の視覚化)
貢献は大歓迎です!助けが必要な領域:
学術的検証: より歴史的な紛争に対するテスト
データ サイエンス: 回帰分析によるスコアしきい値の最適化
UI/UX: ダッシュボードのデザインとユーザー エクスペリエンスの向上
DevOps: コンテナ化 (Docker)、CI/CD パイプライン
ドキュメント: チュートリアル、API ドキュメント、ビデオ ガイド
開始するには、問題または PR を開いてください。
いかなる目的でも自由に使用、変更、展開できます。
リポジトリ: github.com/eshanth23/sentinel
✅ 方法論に関する専門家のフィードバック
✅ 学術的検証パートナーシップ
✅ NGO/シンクタンクとの連携
✅ 現実世界での導入の機会
一人の生徒が作ったものです。無料で実行します。あらゆる人のために作られています。
🌍 「次の戦争を防ぐために建てられた。」
このプロジェクトに価値があると思われる場合は、他の人がこのプロジェクトを発見できるように、⭐リポジトリにスターを付けてください。
SENTINEL - SCSP ハッカソン 2026 向けの AI を活用したグローバル脅威インテリジェンス システム。SENTINEL は、5 つのライブ インテリジェンス シグナルを 1 つの脅威スコア (0 ～ 100) に融合することで、脅威が拡大する前に脅威を検出するリアルタイム競合予測システムです。

）。
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

SENTINEL - AI-powered global threat intelligence system for SCSP Hackathon 2026. SENTINEL is a real-time conflict prediction system that detects threats before they escalate by fusing 5 live intelligence signals into a single threat score (0-100). - GitHub - eshanth23/sentinel: SENTINEL - AI-powered
[truncated]

GitHub - eshanth23/sentinel: SENTINEL - AI-powered global threat intelligence system for SCSP Hackathon 2026. SENTINEL is a real-time conflict prediction system that detects threats before they escalate by fusing 5 live intelligence signals into a single threat score (0-100). · GitHub
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
eshanth23
/
sentinel
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
37 Commits 37 Commits backend backend frontend frontend .gitignore .gitignore DEMO_INSTRUCTIONS.txt DEMO_INSTRUCTIONS.txt README.md README.md requirements.txt requirements.txt View all files Repository files navigation
SENTINEL — AI-Powered Global Threat Intelligence System
Real-time conflict prediction system that would have detected the Ukraine invasion 72 hours before it happened.
Built for SCSP National Security Hackathon 2026 (Wargaming Track)
Demo: https://github.com/user-attachments/assets/72deac0a-56df-4e25-acbd-76c621f63021
Intelligence agencies react to conflicts after they start. The 2022 Ukraine invasion caught the world by surprise — not because the signals weren't there, but because no one was fusing them together in real-time.
Decision-makers need early warning systems that detect threats before escalation.
SENTINEL fuses 5 independent data sources into a single 0-100 threat score, updated every 30 minutes:
Result: When multiple signals align → Early warning before escalation
SENTINEL was tested against 4 major conflicts:
Interactive world map with 10+ monitored regions
Color-coded threat levels (NORMAL → ELEVATED → HIGH → CRITICAL)
Real-time updates every 30 minutes
Conflict Brief Generator — CIA-style intelligence assessments in <5 seconds (Groq + Llama 3)
AI Advisor Chatbot — Ask strategic questions, get answers with live threat context
Claim Verification Engine — Fact-check military claims by cross-referencing 5 sources
Three probabilistic future paths: Escalation / Diplomacy / Standoff
Percentage likelihoods based on current signals
Recommended actions for decision-makers
Visualize how past conflicts escalated signal-by-signal
4 scenarios: Ukraine 2022, Kargil 1999, Gulf War 1990, Israel-Iran 2024
Watch the exact moment SENTINEL would have fired alerts
Aircraft tracking with heading visualization
News aggregation from 65,000 sources
Signal breakdown showing each data source's contribution
Audience
Use Case
🏛️ Defense & Intelligence
Early warning for crisis response teams
🤝 Humanitarian Organizations
Pre-position resources 72 hours before displacement events
📰 Journalists & OSINT
Verify claims in real-time, track conflicts as they develop
🎓 Conflict Researchers
Validate prediction models, study escalation patterns
🏢 Corporations
Supply chain disruption forecasting, geopolitical risk assessment
🛠️ Tech Stack
Backend: Python 3.11 • Flask • APScheduler
Frontend: React 18 • Leaflet.js • Recharts • Axios
AI: Groq API (Llama 3 - 70B parameters)
Data: GDELT • OpenSky • USGS • Yahoo Finance • ACLED
Total Cost: $0.00 (all free APIs)
# Clone repository
git clone https://github.com/eshanth23/sentinel.git
cd sentinel
# Backend setup
cd backend
python -m venv venv
venv \S cripts \a ctivate # Windows | source venv/bin/activate (Mac/Linux)
pip install -r requirements.txt
# Frontend setup
cd ../frontend
npm install
Running SENTINEL
cd backend
venv \S cripts \a ctivate
python api.py
→ Runs on http://localhost:5000
cd frontend
npm start
→ Opens http://localhost:3000
cd backend
python scheduler.py
→ Updates cache with live data (2-3 min)
Region
Score
Level
Top Signals
🔴 Israel-Gaza-Iran
80/100
CRITICAL
News: 28/30, ACLED: 17/20
🔴 Myanmar
75/100
CRITICAL
ACLED: 20/20, Flights: 20/20
🟠 Ukraine-Russia
72/100
HIGH
News: 27/30, ACLED: 18/20
🟠 Taiwan-China
68/100
HIGH
Flights: 20/20, News: 22/30
Updated: May 2026
SENTINEL follows a three-layer design:
📊 Layer 1: Data Collection (scheduler.py)
Runs every 30 minutes or on-demand. Calls 5 APIs for 10 global regions, calculates threat scores (0-100), and saves results to threat_cache.json .
Purpose: Collect and process raw intelligence signals
🔌 Layer 2: API Server (api.py)
Flask server on localhost:5000 . Loads cached data and serves it to the frontend. Provides AI endpoints for conflict briefs, chatbot, and claim verification.
Purpose: Expose data and AI features via REST API
🖥️ Layer 3: Dashboard (App.js)
React application on localhost:3000 . Displays interactive map, visualizations, AI tools, and historical replays. Reads all data from the backend API.
Purpose: User interface for threat monitoring and analysis
Data Flow: External APIs → Scheduler → Cache → Backend → Frontend
✅ Resilient (cache prevents demo failures)
✅ Cost-effective (stays in free tiers)
✅ Scalable (deploy layers separately)
Data Flow:
External APIs → Scheduler → Cache → Backend → Frontend
✅ Resilient — Cache prevents API failures during demos
✅ Fast — Zero latency for UI interactions
✅ Cost-effective — Minimizes API calls, stays in free tiers
✅ Scalable — Each layer can be deployed independently
API
Coverage
Update Frequency
Cost
GDELT
65,000 news sources, 100+ languages
Real-time
Free
OpenSky
4,000+ ADS-B receivers worldwide
Real-time
Free
USGS
Global seismic sensors (mag 2.5+)
Real-time
Free
Yahoo Finance
Top 5 defense contractors
Daily
Free
ACLED
200+ researchers, 50+ countries
Weekly
Free
Groq
Llama 3 (70B params)
On-demand
Free (14.4K/day)
🎓 The Build Story
Timeline: 7 days (April 17-24, 2026)
Prior experience: Zero coding knowledge
Learning: React + Flask from documentation while building
Hackathon: SCSP National Security Hackathon 2026 (Wargaming Track)
🏆 What Makes SENTINEL Different
Existing Systems
SENTINEL
❌ Single-signal (news OR satellite)
✅ Multi-signal fusion (5 sources)
❌ Expensive ($millions)
✅ Free forever ($0.00)
❌ Black-box AI
✅ Explainable (shows signal breakdown)
❌ Reactive (post-event)
✅ Predictive (72-hour window)
❌ Proprietary
✅ Open source
🚀 Future Roadmap
Satellite imagery integration (Sentinel Hub API)
SMS/email alerts for threshold crossings
Expand to 50+ monitored regions
Historical database (20+ years of conflicts)
Machine learning optimization of weights
Change detection algorithms (troop buildup visualization)
Contributions welcome! Areas where help is needed:
Academic validation: Test against more historical conflicts
Data science: Optimize scoring thresholds via regression analysis
UI/UX: Improve dashboard design and user experience
DevOps: Containerization (Docker), CI/CD pipelines
Documentation: Tutorials, API docs, video guides
Open an issue or PR to get started.
Free to use, modify, and deploy for any purpose.
Repository: github.com/eshanth23/sentinel
✅ Expert feedback on methodology
✅ Academic validation partnerships
✅ NGO/Think Tank collaboration
✅ Real-world deployment opportunities
Built by one student. Runs for free. Built for everyone.
🌍 "Built to prevent the next war."
If you find this project valuable, please ⭐ star the repo to help others discover it.
SENTINEL - AI-powered global threat intelligence system for SCSP Hackathon 2026. SENTINEL is a real-time conflict prediction system that detects threats before they escalate by fusing 5 live intelligence signals into a single threat score (0-100).
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
