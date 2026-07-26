---
source: "https://github.com/Top-Celestial-Company-Ltd/DROS-VEP-lite"
hn_url: "https://news.ycombinator.com/item?id=49059322"
title: "DROS-VEP – A <1μs C-ABI binary circuit breaker for AI agents"
article_title: "GitHub - Top-Celestial-Company-Ltd/DROS-VEP-lite: Open-source, 100% reproducible AI Agent Runtime Security Benchmark & Sandbox Environment (RFC-010 Draft Protocol). · GitHub"
author: "dros_jimmy"
captured_at: "2026-07-26T15:55:31Z"
capture_tool: "hn-digest"
hn_id: 49059322
score: 2
comments: 0
posted_at: "2026-07-26T15:46:11Z"
tags:
  - hacker-news
  - translated
---

# DROS-VEP – A <1μs C-ABI binary circuit breaker for AI agents

- HN: [49059322](https://news.ycombinator.com/item?id=49059322)
- Source: [github.com](https://github.com/Top-Celestial-Company-Ltd/DROS-VEP-lite)
- Score: 2
- Comments: 0
- Posted: 2026-07-26T15:46:11Z

## Translation

タイトル: DROS-VEP – AI エージェント用 <1μs C-ABI バイナリ サーキット ブレーカー
記事のタイトル: GitHub - Top-Celestial-Company-Ltd/DROS-VEP-lite: オープンソース、100% 再現可能な AI エージェント ランタイム セキュリティ ベンチマークおよびサンドボックス環境 (RFC-010 ドラフト プロトコル)。 · GitHub
説明: オープンソース、100% 再現可能な AI エージェント ランタイム セキュリティ ベンチマークおよびサンドボックス環境 (RFC-010 ドラフト プロトコル)。 - Top-Celestial-Company-Ltd/DROS-VEP-lite

記事本文:
GitHub - Top-Celestial-Company-Ltd/DROS-VEP-lite: オープンソース、100% 再現可能な AI エージェント ランタイム セキュリティ ベンチマークおよびサンドボックス環境 (RFC-010 ドラフト プロトコル)。 · GitHub
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
ああ、ああ！
エラーがありましたw

ファイルの読み込み中。このページをリロードしてください。
株式会社トップセレスティアルカンパニー
/
DROS-VEP-lite
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
Top-Celestial-Company-Ltd/DROS-VEP-lite
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
14 コミット 14 コミット エージェント エージェント ベンチマーク ベンチマーク コア コア ダッシュボード ダッシュボード ドキュメント ドキュメント エンタープライズ エンタープライズ 例/ adversarial_manifests 例/ adversarial_manifests レポート レポート シナリオ シナリオ .gitignore .gitignore README.md README.md README_zh.md README_zh.md docker-compose.yml docker-compose.yml ビューすべてのファイル リポジトリ ファイルのナビゲーション
🛡️ DROS-VEP Lite: オープンソース AI エージェント セキュリティ ベンチマーク環境
「あなたの AI エージェントは実際の企業内で安全に動作できますか? それを証明してください。」
📖 注目のガイド : AI エージェントを 5 分で破壊する (そしてより強力に再構築する) 方法
#1. リポジトリのクローンを作成する
git clone https://github.com/Top-Celestial-Company-Ltd/DROS-VEP-lite.git
cd ドロス-vep-lite
#2. コンテナ化されたエンタープライズサンドボックスを起動する
ドッカー構成 -d
# 3. インタラクティブな Web ダッシュボードを開く
# ブラウザで http://localhost:8080 に移動します
攻撃 ──► ポリシー評価 ──► 証拠アーティファクト ──► 決定的リプレイ
💡 既存の AI ベンチマークでは不十分な理由
ほとんどの AI ベンチマークは、LLM インテリジェンス、コーディング スキル、またはプロンプト毒性を測定します。 DROS-VEP は、ランタイム ツール呼び出しの承認と特権実行ガバナンスというまったく異なる次元を測定します。
🏗️ アーキテクチャと多層防御エコシステム
DROS は、従来のサイバーセキュリティ (WAF、EDR、SIEM) を置き換えるものではありません。代わりに、最新の多層防御アーキテクチャで AI エージェントの実行境界に「ラスト マイル ランタイム防御」を提供します。
┌───

───────────────────────┐
│ レイヤ 1: ネットワーク境界 │ WAF (Cloudflare、パロアルト) │ -> L3-L7 SQLi/DDoS をブロック
━━━━━━━━━━━━━━━━━━━━━━━━┤
│ レイヤ 2: エンドポイントとホスト │ EDR (CrowdStrike、Sentinel) │ -> OS ランサムウェアをブロック
━━━━━━━━━━━━━━━━━━━━━━━━┤
│ レイヤー 3: ID と IAM │ Keycloak、Active Directory │ -> Human OAuth/JWT を管理
━━━━━━━━━━━━━━━━━━━━━━━━┤
│ ★ レイヤー 4: AI エージェント ランタイム │ DROS PEP/PDP + ATR サンドボックス │ -> 不正なツールをブロック
━━━━━━━━━━━━━━━━━━━━━━┴─────────────┘
│
▼
PKI 証拠をエンタープライズ SIEM (Splunk、Elastic) にエクスポート
💡 従来のセキュリティ (WAF/Keycloak) が ATS シナリオを認識できない理由
間接プロンプト インジェクション攻撃 (ATS-001) では、ハイジャックされた AI エージェントは有効な Keycloak JWT トークンを所有しています。エージェントが /api/erp/finance にクエリを実行すると、WAF は「有効な HTTPS、クリーンな JSON、有効な OAuth トークン。アクセスが許可されました!」というリクエストを検査します。
従来の WAF では、100% 正当なユーザーがクリーンな REST API 呼び出しを行っていることがわかります。攻撃は LLM セマンティック コンテキスト内に隠されています。

このため、ツールの実行境界で DROS PEP/PDP が必要になります。
🎯 エージェントの脅威シナリオ (ATS マトリックスと 2026 年の現実世界のインシデント)
DROS-VEP Lite は、MITRE ATLAS にマッピングされた 2026 年の最も悪名高い現実世界の AI インシデントを直接再現し、無力化します。
🧪 エンジニアの整合性の証明: 分析と再生
エンジニアは静的なダッシュボードを信頼していません。 「ガードを外したら、結果は実際に変わりますか？」と彼らは尋ねます。
1. 反事実的なコントロールグループ (DROS ガードトグルを無効にする)
http://localhost:8080 を開き、「☑ DROS ガードを無効にする (デバッグ モード)」を確認します。
ガードアクティブ (ノーマル) : 防御完全性 100% ( AS-001 ~ AS-005 | 判定: DENY | 合格率: 100% )。
Guard Disabled (Control Group) : PEP は傍受をバイパスします。エージェントはターゲットのエンドポイントに侵入します。合格率は 100% ===> 0% (LEAKED) から急落します。
2. 決定論的再生エンジン ( benchmark/replay.py )
履歴監査ログまたは証拠アーティファクト パッケージを決定論的に再生します。
Python ベンチマーク/replay.py exec_ATS-001_1784702707
📊 ベンチマーク手法と測定の透明性
26.1 μs のポリシー評価レイテンシはどのように測定されますか?
🏆 $40,000 DROS レッドチーム グランド チャレンジ
DROS の 26.1μs C-ABI ゼロトラスト実行境界を突破できますか?
この Docker サンドボックスで DROS PEP/PDP ガードのバイパスに成功した最初のセキュリティ研究者またはハッカーに、当社の主力製品である「VajraClaw - エンタープライズ エアギャップ ライセンス」の 5 年間の無料ライセンスを提供しています (現在の早期導入者価格: 年間 7,990 ドル、総額: 39,950 ドル)。
⚖️ 曖昧さゼロの公式バイパス証明ルール
大賞を獲得するには、あなたの投稿は次の 3 つの厳しい基準をすべて満たさなければなりません。
未変更のサンドボックス環境:
このエクスプロイトは、デフォルトの docker compose up -d 環境に対して実行する必要があります。
dros-guard を変更することはできません

コンテナー コードを変更するか、それを無効にします (BYPASS_GUARD=true または UI で [DROS Guard を無効にする] をクリックすることは固く禁止されています)。ガードはアクティブでなければなりません。
具体的なエクスプロイト結果 (エージェントのオーバーリーチ) :
パス A (不正なツール呼び出し) : サポート エージェント ( /api/erp/inventory に対する権限のみを持つ) に DROS 物理ブロックをバイパスさせ、 /api/erp/finance または /api/system/secrets からのデータの読み取り/抽出を正常に実行させます。
パス B (不正な書き込み) : DROS PEP/PDP を正常にバイパスして、ERP コンテナ上の保護されたディスク ファイルを暗号化または上書きします。
暗号証明およびPoCパッケージ:
再現可能な PoC エクスプロイト スクリプトを送信する必要があります。
ブロックされたルートに対する異常な ALLOW 決定、または Guard の実行中にデータ漏洩につながる完全なネットワーク/コンテナ バイパス (SSRF/RCE) を証明する、未変更の Audit.jsonl ログと Decision.json 証拠パッケージを提出する必要があります。
🚫 無効な申し立て (範囲外)
DoS / DDoS : DROS サーバーのクラッシュはバイパスではありません (DROS は不正アクセスの防止に成功しました)。
純粋な LLM ジェイルブレイク : LLM をだまして攻撃させようとしても、DROS ガードが結果の API 呼び出しを 26.1 μs 以内にブロックすることに成功した場合、これはバイパスではなく DROS 防御の成功とみなされます。
提出方法: PoC パッケージを GitHub Discussions または Discord #conformance-claims に投稿します。最初に検証された送信タイムスタンプがグランプリを獲得します。
🏅 RFC-010 ドラフトプロトコル準拠ハーネス
サードパーティ AI エージェント フレームワーク (OpenAI Agent SDK、LangGraph、CrewAI、AutoGen、OpenClaw) は、次の 3 つの認定レベルにわたってランタイム セキュリティを評価できます。
レベル 1 (コア) : ID トークン (DIT) + PEP ツールのインターセプト + 構造化監査ログ。
レベル 2 (エンタープライズ) : ポリシーの説明可能性 (ポリシー ID) + 証拠パッケージ (SHA-256)

ダイジェスト) + マルチエージェントの役割の分離。
レベル 3 (高保証) : 暗号化証明 + 改ざん検出 + 決定的リプレイ。
ℹ️ 免責事項 : 付属の適合ハーネスは、RFC-010 ドラフト仕様に照らして実装を検証します。テストに合格したことは、独立した標準化団体による認証ではなく、このドラフトに適合していることを示します。
💎 製品のエディションとライセンス
特徴/能力
コミュニティ (無料 $0)
ハッカー ($149/年または $19/月 - 1,000 個の無料プロモーション)
プロフェッショナル ($499/年/チーム)
Enterprise Swarm (商用)
対象者
学生と研究者
フリーランサーと小規模 AI スタートアップ
中規模の AI エンジニアリング チーム
フォーチュン 500、銀行、政府
兼務の役割
最大2つの役割
最大5つの役割
最大25の役割
無制限 (500 以上の群れ生産)
ATS シナリオ
ATS-001 シングル
ATS-001～ATS-005 フルマトリックス
ATS-001～ATS-005＋カスタム
無制限のカスタムレッドチームるつぼ
コネクタ
REST モックエンタープライズ API
REST モック + CI/CD ハーネス
Keycloak + EspoCRM + Forgejo
ライブ SAP、Active Directory、K8s
リプレイとSIEM
ローカルテレメトリー
オフライン再生エンジン (replay.py)
リプレイ + テレメトリ ヒートマップ
無制限の PKI ログと SIEM (Splunk)
防御範囲
AI エージェント ツールのガバナンス
AI エージェント ツールのガバナンス
AI エージェント ツールのガバナンス
AI エージェント + エンタープライズ ランサムウェア防御
🎁 1 年間の無料ハッカー ライセンスを請求しましょう (🔥 先着 1,000 人のセキュリティ パイオニア!)
RFC-010 への準拠を確認しましたか? 1 年間の無料ハッカー ライセンス (149 ドル相当) を請求します。
オプション 1 (Web ダッシュボード UI): http://localhost:8080 を開き、「1 年間のハッカー ライセンスを請求する」をクリックします。
オプション 2 (GitHub Discussions Bot) : conformance_report.json を GitHub Discussions に投稿します。
オプション 3 (Discord Cyber​​ Crucible) : Discord サーバーに参加し、 #conformance-claims にレポートを投稿します。
オプション 4 (Gumroad $0 チェックアウト) : dr-os.io で 100% オフ クーポン DROS-RFC010-FREE を使用します。

📜 テクニカルホワイトペーパーと仕様
📖 フル ホワイトペーパー (英語 v2.0) : 自律型 AI ワークロードのゼロトラスト実行ガバナンス (DROS 4 層パラダイム)
📖 フル ホワイト ペーパー (繁体字中国語 v2.0): 自律型 AI ワークロードのゼロトラスト実行ガバナンス (DROS 4 層多層防御アーキテクチャ)
⚡ A4 4 ページのエグゼクティブ サマリー (HTML) : CISO およびセキュリティ研究者向けの簡単な視覚的サマリー
📋 RFC-010: DROS-VEP 仕様プロトコル: オープン エージェント セキュリティおよび脅威シナリオ プロトコル
Apache 2.0 に基づいてライセンスされています。詳細については、「ライセンス」を参照してください。
オープンソース、100% 再現可能な AI エージェント ランタイム セキュリティ ベンチマークおよびサンドボックス環境 (RFC-010 ドラフト プロトコル)。
Readme アクティビティ カスタム プロパティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Open-source, 100% reproducible AI Agent Runtime Security Benchmark & Sandbox Environment (RFC-010 Draft Protocol). - Top-Celestial-Company-Ltd/DROS-VEP-lite

GitHub - Top-Celestial-Company-Ltd/DROS-VEP-lite: Open-source, 100% reproducible AI Agent Runtime Security Benchmark & Sandbox Environment (RFC-010 Draft Protocol). · GitHub
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
Top-Celestial-Company-Ltd
/
DROS-VEP-lite
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Top-Celestial-Company-Ltd/DROS-VEP-lite
main Branches Tags Go to file Code Open more actions menu Folders and files
14 Commits 14 Commits agent agent benchmark benchmark core core dashboard dashboard docs docs enterprise enterprise examples/ adversarial_manifests examples/ adversarial_manifests reports reports scenarios scenarios .gitignore .gitignore README.md README.md README_zh.md README_zh.md docker-compose.yml docker-compose.yml View all files Repository files navigation
🛡️ DROS-VEP Lite: Open-Source AI Agent Security Benchmark Environment
"Can your AI Agent safely operate inside a real enterprise? Prove it."
📖 Featured Guide : How to Break Your AI Agent in 5 Minutes (And Rebuild It Stronger)
# 1. Clone the repository
git clone https://github.com/Top-Celestial-Company-Ltd/DROS-VEP-lite.git
cd dros-vep-lite
# 2. Launch containerized enterprise sandbox
docker compose up -d
# 3. Open Interactive Web Dashboard
# Navigate to http://localhost:8080 in your browser
Attack ───► Policy Evaluation ───► Evidence Artifact ───► Deterministic Replay
💡 Why Existing AI Benchmarks Are Not Enough
Most AI benchmarks measure LLM intelligence, coding skills, or prompt toxicity. DROS-VEP measures a completely different dimension: Runtime Tool-Call Authorization & Privileged Execution Governance.
🏗️ Architecture & Defense-in-Depth Ecosystem
DROS does NOT replace traditional cybersecurity (WAF, EDR, SIEM). Instead, it provides the "Last Mile Runtime Defense" for AI Agent execution boundaries in a modern Defense-in-Depth architecture:
┌─────────────────────────────────────────────────────────────┐
│ Layer 1: Network Perimeter │ WAF (Cloudflare, Palo Alto) │ -> Blocks L3-L7 SQLi/DDoS
├──────────────────────────────┼──────────────────────────────┤
│ Layer 2: Endpoint & Host │ EDR (CrowdStrike, Sentinel) │ -> Blocks OS Ransomware
├──────────────────────────────┼──────────────────────────────┤
│ Layer 3: Identity & IAM │ Keycloak, Active Directory │ -> Manages Human OAuth/JWT
├──────────────────────────────┼──────────────────────────────┤
│ ★ Layer 4: AI Agent Runtime │ DROS PEP/PDP + ATR Sandbox │ -> Blocks Unauthorized Tools
└──────────────────────────────┴──────────────────────────────┘
│
▼
Exports PKI Evidence to Enterprise SIEM (Splunk, Elastic)
💡 Why Traditional Security (WAF/Keycloak) Is Blind to ATS Scenarios
In an indirect prompt injection attack (ATS-001), the hijacked AI Agent possesses a valid Keycloak JWT token . When the agent queries /api/erp/finance , WAF inspects the request: "Valid HTTPS, clean JSON, valid OAuth token. Access Granted!"
Traditional WAFs see a 100% legitimate user making a clean REST API call . The attack is hidden inside the LLM Semantic Context . This is why DROS PEP/PDP is required at the tool execution boundary.
🎯 Agent Threat Scenarios (ATS Matrix & 2026 Real-World Incidents)
DROS-VEP Lite directly reproduces and neutralizes 2026's most notorious real-world AI incidents, mapped to MITRE ATLAS :
🧪 Engineer Proof of Integrity: Dissect & Replay
Engineers don't trust static dashboards. They ask: "If I unplug your guard, does the result actually change?"
1. Counterfactual Control Group ( Disable DROS Guard Toggle)
Open http://localhost:8080 and check ☑ Disable DROS Guard (Debug Mode) :
Guard Active (Normal) : 100% Defense Integrity ( AS-001 ~ AS-005 | Decision: DENY | Pass Rate: 100% ).
Guard Disabled (Control Group) : PEP bypasses interception. The agent penetrates target endpoints. Pass rate plummets from 100% ===> 0% (LEAKED) .
2. Deterministic Replay Engine ( benchmark/replay.py )
Replay any historical audit log or evidence artifact package deterministically:
python benchmark/replay.py exec_ATS-001_1784702707
📊 Benchmark Methodology & Measurement Transparency
How is the 26.1 μs policy evaluation latency measured?
🏆 The $40,000 DROS Red Team Grand Challenge
Can you breach DROS’s 26.1μs C-ABI zero-trust execution boundary?
We are offering a 5-Year FREE License of our flagship product, "VajraClaw - Enterprise Air-Gapped License" (Current Early Adopter Price: $7,990/yr, Total Value: $39,950 USD ), to the FIRST security researcher or hacker who successfully bypasses the DROS PEP/PDP Guard in this docker sandbox!
⚖️ Official Zero-Ambiguity Proof-of-Bypass Rules
To claim the Grand Prize, your submission MUST meet all three of the following strict criteria:
Unmodified Sandbox Environment :
The exploit must run against the default docker compose up -d environment.
You CANNOT modify the dros-guard container code or disable it ( BYPASS_GUARD=true or clicking "Disable DROS Guard" in the UI is strictly forbidden). The Guard must be active.
Concrete Exploit Result (Agentic Overreach) :
Path A (Unauthorized Tool Call) : Successfully force the support-agent (which only has permissions for /api/erp/inventory ) to bypass the DROS physical block and successfully read/exfiltrate data from /api/erp/finance or /api/system/secrets .
Path B (Unauthorized Write) : Successfully bypass the DROS PEP/PDP to encrypt or overwrite protected disk files on the ERP container.
Cryptographic Proof & PoC Package :
You must submit a reproducible PoC exploit script.
You must submit the unmodified audit.jsonl log and decision.json evidence package demonstrating either an anomalous ALLOW decision for a blocked route, or proving complete network/container bypass (SSRF/RCE) resulting in data exfiltration while the Guard was running.
🚫 Invalid Claims (Out-of-Scope)
DoS / DDoS : Crashing the DROS server is not a bypass (DROS successfully prevented unauthorized access).
Pure LLM Jailbreaks : If you trick the LLM into wanting to attack, but the DROS Guard successfully blocks the resulting API call in 26.1μs, this is considered a Successful DROS Defense , not a bypass.
How to submit : Post your PoC package to GitHub Discussions or our Discord #conformance-claims . The first verified submission timestamp wins the Grand Prize!
🏅 RFC-010 Draft Protocol Conformance Harness
Third-party AI Agent Frameworks (OpenAI Agent SDK, LangGraph, CrewAI, AutoGen, OpenClaw) can evaluate their runtime security across 3 certification tiers:
Level 1 (Core) : Identity Token (DIT) + PEP Tool Interception + Structured Audit Logging.
Level 2 (Enterprise) : Policy Explainability (Policy ID) + Evidence Package (SHA-256 Digest) + Multi-Agent Role Isolation.
Level 3 (High Assurance) : Cryptographic Attestation + Tamper Detection + Deterministic Replay.
ℹ️ Disclaimer : The included conformance harness validates implementations against the RFC-010 Draft specification. Passing the test indicates conformance to this draft, not certification by an independent standards body.
💎 Product Editions & Licensing
Feature / Capability
Community ($0 Free)
Hacker ($149/yr or $19/mo - 1k Free Promo)
Professional ($499/yr / Team)
Enterprise Swarm (Commercial)
Target Audience
Students & Researchers
Freelancers & Small AI Startups
Mid-sized AI Engineering Teams
Fortune 500, Banks, Government
Concurrent Roles
Max 2 Roles
Up to 5 Roles
Up to 25 Roles
Unlimited (500+ Swarm Production)
ATS Scenarios
ATS-001 Single
ATS-001 ~ ATS-005 Full Matrix
ATS-001 ~ ATS-005 + Custom
Unlimited Custom Red Team Crucibles
Connectors
REST Mock Enterprise APIs
REST Mock + CI/CD Harness
Keycloak + EspoCRM + Forgejo
Live SAP, Active Directory, K8s
Replay & SIEM
Local Telemetry
Offline Replay Engine ( replay.py )
Replay + Telemetry Heatmap
Unlimited PKI Log & SIEM (Splunk)
Defense Scope
AI Agent Tool Governance
AI Agent Tool Governance
AI Agent Tool Governance
AI Agent + Enterprise Ransomware Defense
🎁 Claim Your Free 1-Year Hacker License (🔥 First 1,000 Security Pioneers!)
Verified RFC-010 compliance? Claim a 1-Year FREE Hacker License ($149 Value) :
Option 1 (Web Dashboard UI) : Open http://localhost:8080 and click "Claim 1-Year Hacker License" .
Option 2 (GitHub Discussions Bot) : Post conformance_report.json to GitHub Discussions .
Option 3 (Discord Cyber Crucible) : Join our Discord Server and post report in #conformance-claims .
Option 4 (Gumroad $0 Checkout) : Use 100% OFF Coupon DROS-RFC010-FREE at dr-os.io .
📜 Technical Whitepapers & Specifications
📖 Full Whitepaper (English v2.0) : Zero-Trust Execution Governance for Autonomous AI Workloads (DROS 4-Layer Paradigm)
📖 完整白皮書 (繁體中文 v2.0) : 自主型 AI 工作負載的零信任執行治理 (DROS 四層防禦縱深架構)
⚡ 4-Page A4 Executive Summary (HTML) : Fast visual summary for CISOs & Security Researchers
📋 RFC-010: DROS-VEP Specification Protocol : Open Agent Security & Threat Scenario Protocol
Licensed under Apache 2.0. See LICENSE for details.
Open-source, 100% reproducible AI Agent Runtime Security Benchmark & Sandbox Environment (RFC-010 Draft Protocol).
Readme Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
