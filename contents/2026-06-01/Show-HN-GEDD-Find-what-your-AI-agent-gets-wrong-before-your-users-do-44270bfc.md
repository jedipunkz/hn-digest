---
source: "https://github.com/aws-samples/sample-GEDD"
hn_url: "https://news.ycombinator.com/item?id=48349448"
title: "Show HN: GEDD – Find what your AI agent gets wrong (before your users do)"
article_title: "GitHub - aws-samples/sample-GEDD: Find what your AI agent gets wrong — before you have a rubric. Qualitative eval for PMs. · GitHub"
author: "balasvce19855"
captured_at: "2026-06-01T00:25:03Z"
capture_tool: "hn-digest"
hn_id: 48349448
score: 2
comments: 0
posted_at: "2026-05-31T20:31:24Z"
tags:
  - hacker-news
  - translated
---

# Show HN: GEDD – Find what your AI agent gets wrong (before your users do)

- HN: [48349448](https://news.ycombinator.com/item?id=48349448)
- Source: [github.com](https://github.com/aws-samples/sample-GEDD)
- Score: 2
- Comments: 0
- Posted: 2026-05-31T20:31:24Z

## Translation

タイトル: HN を表示: GEDD – AI エージェントが間違っている点を (ユーザーが間違っている前に) 見つけます
記事のタイトル: GitHub - aws-samples/sample-GEDD: ルーブリックを作成する前に、AI エージェントが間違っている点を見つけてください。 PM の定性的評価。 · GitHub
説明: ルーブリックを作成する前に、AI エージェントが間違っている点を見つけてください。 PM の定性的評価。 - aws-サンプル/サンプル-GEDD

記事本文:
GitHub - aws-samples/sample-GEDD: ルーブリックを作成する前に、AI エージェントが間違っている点を見つけてください。 PM の定性的評価。 · GitHub
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
aws-サンプル
/
サンプル-GEDD
公共
通知
にサインインする必要があります

通知設定を変更する
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
152 コミット 152 コミット .github .github grounded-evals grounded-evals CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE METHODLOGY.md METHODOLOGY.md README.md README.md SETUP.md SETUP.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
GEDD — AI エージェントが間違っている点を見つける - プロダクト マネージャーとドメイン エキスパートのためのクロード スキル
AI エージェントを出荷しました。次に、CEO、コンプライアンス、そしてそれを継承するチームに対して、それが機能することを証明する必要があります。エージェントは、ルーブリックが予期していなかった方法で失敗し、評価ツールは、何が失敗するかを確認する前に、何を測定すべきかを知ることを期待します。
GEDD はルーブリックを作成する前のツールです。ドメインの専門家が会話を行い、90 分後には実稼働評価パイプラインが完成します。
eval パイプラインは製品です。エージェントは、それが生み出すものにすぎません。
📖 なぜグラウンデッド・セオリーなのか?信頼できる AI エージェントのために — このリポジトリの背後にある長い形式の議論。
フローチャート TD
サブグラフ DE["🧑‍💼 DOMAIN EXPERT — /gedd in Claude Code"]
TB方向
S1["1️⃣ エージェントを定義"]
S2["2️⃣ システム プロンプト"]
S3["3️⃣ AgentCore にデプロイ"]
S4["4️⃣ ゴールデンクエリ"]
S5["5️⃣注釈を付けて判断する"]
S1 ==> S2 ==> S3 ==> S4 ==> S5
終わり
S5 ==>|"📄 session.json"|ハンドオフ:::ハンドオフ
ハンドオフ ==> S6
サブグラフ ML["🔧 ML ENGINEER — grounded-evals mlflow"]
TB方向
S6["6️⃣ SageMaker MLflow パイプライン"]
終わり
S3 -.->|"デプロイ"| AC["☁️ Bedrock AgentCore"]
S4 -.->|"呼び出し"| BR["🤖 クロード俳句 4.5"]
S6 -.->|"トラック"| SM["📊 SageMaker MLflow"]
S6 -.->|"ゲート"| CI["🚦 CI/CD パイプライン"]
classDef ハンドオフ fill:#fce4ec、ストローク:#c62828、ストローク幅:3px、ストローク-dasharray: 5 5
読み込み中
2

ペルソナ。 6つのステップ。 1 つのファイルがそれらを接続します。
テスト前にデプロイする理由は何ですか?エージェントに必要なのはシステム プロンプトのみです。ステップ 3 でデプロイすると、レイテンシー、IAM、コールド スタートを含むすべてのゴールデン クエリが実際のエンドポイントに対して実行されます。
パイプラインは線形ではなく、ループです。本番環境での失敗は、新しいテスト ケースにフィードバックされます。 eval スイートはエージェントとともに成長します。
フローチャート TD
サブグラフ EXPERT["🧑‍💼ドメインエキスパート"]
D["定義 + プロンプト + デプロイ"]
Q[「ゴールデン クエリ<br/><i>オープン コーディング手法</i>」]
A["注釈を付ける<br/><i>✓/⚠/✗ + エラー コード</i>"]
D --> Q --> A
終わり
サブグラフ ENGINEER["🔧 ML ENGINEER"]
J["ビルド ジャッジ<br/><i>ルーブリック + ウェイト + ハードフェイル</i>"]
K{"校正<br/>κ ≥ 0.80?"}
CI["CI/CD ゲート<br/><i>TSR ≥ 95%</i>"]
J --> K
K -->|"はい"| CI
K -->|"いいえ — 基準を修正します"| J
終わり
A -->|"session.json"| J
CI -->|"✅ 船"| PROD["🚀 プロダクション"]
PROD -.->|"🔄 新たな障害が発見されました"| Q
スタイル PROD 塗りつぶし:#c8e6c9、ストローク:#2e7d32
読み込み中
各ガイドはフライホイールのセクションにマップされます。
cd 接地評価
pip install -e 。
クロード
/ゲッド
90分 → ゴールデンデータセット + 審査員
pip インストール sagemaker-mlflow
grounded-evals mlflow \
--session session.json \
--tracking-uri $ARN \
--run-eval
デモを探索する
pip install -e " .[dev] "
グラウンデッド評価サービス
ローカルホストを開く:8080
17 のプリロードされたシナリオ
ドメインエキスパートが発見したもの
4 つのドメインにわたってテストしました。どのケースでも、専門家はエンジニアが見逃してしまうような失敗を発見しました。
これらは一般的な「幻覚」のラベルではありません。これらは、専門家自身の用語ではドメイン固有の故障モードであり、配置されたジャッジの基準になります。
フローチャート LR
CC["クロード コード<br/><i>/gedd スキル</i>"] --> SJ["session.json"]
SJ --> CLI["grounded-evals mlflow"]
CLI --> SM["SageMaker MLflow<br/><i>実験 + 審査員</i>"]
CLI --> BR["岩盤<br/><i>アゲン

tコア + クロード</i>"]
SM --> CICD["CI/CD<br/><i>回帰ゲート</i>"]
CICD --> BR
読み込み中
すべて AWS ネイティブです。認証用の IAM。アーティファクトの場合は S3。外部サービスはありません。
LLM 呼び出しは必要ありません。それぞれに、ゴールデン クエリ、注釈、エラー コード、生成された判定が事前にロードされています。
コマンド
何をするのか
チャット
会話型コーチング (ステップ 1 ～ 5)
評価
モデルに対してゴールデン クエリを実行する
注釈を付ける
応答にエラー コード ✓/⚠/✗ をマークします
裁判官
G-Eval ジャッジプロンプトを生成する
ミリフロー
SageMaker MLflow へのエクスポート (ステップ 6)
輸出
ゴールデン データセットを JSONL/CSV/JSON として書き込む
ステータス
セッションダッシュボード
分析する
エラーコードを評価次元にマッピングする
奉仕する
Web UIを起動する
骨折
ドメインをテスト カテゴリに分割する
チェック飽和
データセットのカバレッジを確認する
適用範囲
カテゴリ別の棒グラフの内訳
比較する
新しいプロンプトによって独自のカバレッジが追加されるかどうかを確認する
これが機能する理由
ほとんどの eval ツールは、何を測定する必要があるのかを尋ねます。 GEDD はこう尋ねます。実際に何が起こっているのでしょうか?
観察していないものを評価することはできません。事前に焼き付けられたルーブリックは、エージェント固有の失敗を見逃します。
基準は証拠によって重み付けされます。用量単位の混乱は、トーンスリップと同じ深刻さではありません。
あなたの評価はエージェントとともに進化します。フライホイールは新しい故障モードを自然に吸収します。
仕事には負荷がかかるものになります。判断基準はあなたの専門用語の中にあり、一般的な「有用性 1 ～ 5」ではありません。
GEDD がエージェントの間違いを見つけるのに役立つ場合、スターは他の人が間違いを見つけるのにも役立ちます。
ライセンス: MIT-0。 「ライセンス」を参照してください。セキュリティ: 「貢献」を参照してください。
ルーブリックを作成する前に、AI エージェントが間違っている点を見つけてください。 PM の定性的評価。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
ありました

ロード中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Find what your AI agent gets wrong — before you have a rubric. Qualitative eval for PMs. - aws-samples/sample-GEDD

GitHub - aws-samples/sample-GEDD: Find what your AI agent gets wrong — before you have a rubric. Qualitative eval for PMs. · GitHub
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
aws-samples
/
sample-GEDD
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
152 Commits 152 Commits .github .github grounded-evals grounded-evals CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE METHODOLOGY.md METHODOLOGY.md README.md README.md SETUP.md SETUP.md View all files Repository files navigation
GEDD — find what your AI agent gets wrong - A Claude Skill for Product Managers and Domain Experts
You shipped an AI agent. Now you need to prove it works — to your CEO, to compliance, to the team that inherits it. The agent fails in ways no rubric anticipated, and the eval tools expect you to know what to measure before you've seen what breaks.
GEDD is the tool for before you have a rubric. A domain expert has a conversation, and 90 minutes later you have a production eval pipeline.
The eval pipeline is the product. The agent is just the thing it produces.
📖 Why Grounded Theory? for reliable AI Agents — the long-form argument behind this repo.
flowchart TD
subgraph DE["🧑‍💼 DOMAIN EXPERT — /gedd in Claude Code"]
direction TB
S1["1️⃣ Define Agent"]
S2["2️⃣ System Prompt"]
S3["3️⃣ Deploy to AgentCore"]
S4["4️⃣ Golden Queries"]
S5["5️⃣ Annotate & Judge"]
S1 ==> S2 ==> S3 ==> S4 ==> S5
end
S5 ==>|"📄 session.json"| HANDOFF:::handoff
HANDOFF ==> S6
subgraph ML["🔧 ML ENGINEER — grounded-evals mlflow"]
direction TB
S6["6️⃣ SageMaker MLflow Pipeline"]
end
S3 -.->|"deploy"| AC["☁️ Bedrock AgentCore"]
S4 -.->|"invoke"| BR["🤖 Claude Haiku 4.5"]
S6 -.->|"track"| SM["📊 SageMaker MLflow"]
S6 -.->|"gate"| CI["🚦 CI/CD Pipeline"]
classDef handoff fill:#fce4ec,stroke:#c62828,stroke-width:3px,stroke-dasharray: 5 5
Loading
Two personas. Six steps. One file connects them.
Why deploy before testing? The agent only needs the system prompt. By deploying at Step 3, all golden queries run against the real endpoint — latency, IAM, cold starts included.
The pipeline isn't linear — it's a loop. Production failures feed back into new test cases. The eval suite grows with the agent.
flowchart TD
subgraph EXPERT["🧑‍💼 DOMAIN EXPERT"]
D["Define + Prompt + Deploy"]
Q["Golden Queries<br/><i>Open Coding methodology</i>"]
A["Annotate<br/><i>✓/⚠/✗ + error codes</i>"]
D --> Q --> A
end
subgraph ENGINEER["🔧 ML ENGINEER"]
J["Build Judge<br/><i>Rubric + weights + hard-fails</i>"]
K{"Calibrate<br/>κ ≥ 0.80?"}
CI["CI/CD Gate<br/><i>TSR ≥ 95%</i>"]
J --> K
K -->|"Yes"| CI
K -->|"No — fix criteria"| J
end
A -->|"session.json"| J
CI -->|"✅ Ship"| PROD["🚀 Production"]
PROD -.->|"🔄 New failure discovered"| Q
style PROD fill:#c8e6c9,stroke:#2e7d32
Loading
Each guide maps to a section of the flywheel:
cd grounded-evals
pip install -e .
claude
/gedd
90 min → golden dataset + judge
pip install sagemaker-mlflow
grounded-evals mlflow \
--session session.json \
--tracking-uri $ARN \
--run-eval
Explore Demos
pip install -e " .[dev] "
grounded-evals serve
Open localhost:8080
17 pre-loaded scenarios
What the Domain Expert Discovers
We tested across 4 domains. In every case, the expert caught failures an engineer would miss:
These aren't generic "hallucination" labels. They're domain-specific failure modes in the expert's own vocabulary — and they become the criteria in the deployed judge.
flowchart LR
CC["Claude Code<br/><i>/gedd skill</i>"] --> SJ["session.json"]
SJ --> CLI["grounded-evals mlflow"]
CLI --> SM["SageMaker MLflow<br/><i>Experiments + Judges</i>"]
CLI --> BR["Bedrock<br/><i>AgentCore + Claude</i>"]
SM --> CICD["CI/CD<br/><i>Regression gates</i>"]
CICD --> BR
Loading
All AWS-native. IAM for auth. S3 for artifacts. No external services.
No LLM calls needed. Each is pre-loaded with golden queries, annotations, error codes, and a generated judge.
Command
What it does
chat
Conversational coaching (Steps 1-5)
eval
Run golden queries against a model
annotate
Mark responses ✓/⚠/✗ with error codes
judge
Generate G-Eval judge prompt
mlflow
Export to SageMaker MLflow (Step 6)
export
Write golden dataset as JSONL/CSV/JSON
status
Session dashboard
analyze
Map error codes to eval dimensions
serve
Start the web UI
fracture
Fracture domain into test categories
check-saturation
Check dataset coverage
coverage
Bar-chart breakdown by category
compare
Check if a new prompt adds unique coverage
Why This Works
Most eval tools ask: what should we measure? GEDD asks: what is actually happening?
You can't evaluate what you haven't observed. Pre-baked rubrics miss your agent's unique failures.
Criteria are weighted by evidence. A dosage unit confusion isn't the same severity as a tone slip.
Your evaluation evolves with the agent. The flywheel absorbs new failure modes naturally.
Your work becomes load-bearing. The judge is in your domain vocabulary, not generic "helpfulness 1-5."
If GEDD helped you find what your agent gets wrong, a star helps others find it too.
License: MIT-0. See LICENSE . Security: see CONTRIBUTING .
Find what your AI agent gets wrong — before you have a rubric. Qualitative eval for PMs.
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
