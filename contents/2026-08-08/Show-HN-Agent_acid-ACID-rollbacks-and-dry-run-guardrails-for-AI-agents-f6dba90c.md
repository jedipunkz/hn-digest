---
source: "https://github.com/muhammadwaqasai/agent_acid"
hn_url: "https://news.ycombinator.com/item?id=49219014"
title: "Show HN: Agent_acid – ACID rollbacks and dry-run guardrails for AI agents"
article_title: "GitHub - muhammadwaqasai/agent_acid: ACID-style rollback and session-memory guardrails for autonomous AI agents · GitHub"
author: "waqasai123"
captured_at: "2026-08-08T05:35:28Z"
capture_tool: "hn-digest"
hn_id: 49219014
score: 1
comments: 0
posted_at: "2026-08-08T05:00:26Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Agent_acid – ACID rollbacks and dry-run guardrails for AI agents

- HN: [49219014](https://news.ycombinator.com/item?id=49219014)
- Source: [github.com](https://github.com/muhammadwaqasai/agent_acid)
- Score: 1
- Comments: 0
- Posted: 2026-08-08T05:00:26Z

## Translation

タイトル: HN を表示: Agent_acid – AI エージェントの ACID ロールバックとドライラン ガードレール
記事のタイトル: GitHub - muhammadwaqasai/agent_acid: 自律型 AI エージェントのための ACID スタイルのロールバックとセッション メモリ ガードレール · GitHub
説明: 自律 AI エージェント用の ACID スタイルのロールバックとセッション メモリ ガードレール - muhammadwaqasai/agent_acid

記事本文:
GitHub - muhammadwaqasai/agent_acid: 自律 AI エージェント用の ACID スタイルのロールバックとセッション メモリ ガードレール · GitHub
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
ムハンマドワカサイ
/
エージェント酸
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット Agent_acid Agent_acid 例 例 テスト テスト .gitignore .gitignore README.md README.md pyproject.toml pyproject.toml 要件.txt 要件.t

xt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
自律型 AI エージェントに対する ACID スタイルのトランザクション保証。
現実世界のアクション (カードのチャージ、データベースへの書き込み、電子メールの送信、API の呼び出し) を実行する AI エージェントには、部分的に完了したタスクを「元に戻す」標準的な方法がなく、ほとんどのガードレール システムは一度に 1 つのアクションのみを検証し、セッションのメモリはありません。 Agent_acid は両方のギャップを埋めます。
自動ロールバック — マルチステップ エージェント プランのいずれかのステップが失敗した場合、その前に完了したすべてのステップが逆の順序で自動的に元に戻されます。
ガードレール — 何も「クラッシュ」しない場合でも、不正な AI 出力をブロックする、コードレベルの厳しいルール (プロンプトではない)。
ステートフル (セッション全体) ガードレール — 攻撃者や操作された AI が、1 つの大きな禁止されたアクションを、個別に合法ないくつかの小さなアクションに分割する (「サラミのスライス」) など、多段階の操作を捕捉します。
これは理論的な枠組みではありません。以下のすべての主張は、このリポジトリ内の実行可能なテストまたはライブ デモによって裏付けられています。
AI エージェントに対するサラミ スライシング スタイルの攻撃は、活発に研究されている問題です。一度に 1 つのツール呼び出しのみを判断するガードレールは「メモリレス」であり、攻撃者が禁止されたアクションを、どのステップもアラームを作動させない多くの小さなステップに分散させることができます。 Agent_acid のステートフル ガードレールは、このギャップを埋めるために特別に構築されています。
pip install -r required.txt # または単に: pip install openai pytest
自動テスト スイートを実行します (API キーは必要ありません。コア エンジンが動作することを証明します)。
pytest テスト/test_core.py -v
基本的なロールバック デモを実行します (API キーは必要ありません)。
Python の例/basic_rollback.py
ガードレールのデモを実行します (API キーは必要ありません)。
Python の例/guardrail_demo.py
実際の AI エージェント デモを実行します (OpenAI API キーが必要です)。
import OPENAI_API_KEY= " sk-... " # Mac/Linux
$env :OPENAI_API_KEY = " sk-... " # Windows PowerShe

ll
Python の例/llm_agent_demo.py
Python の例/アタック_テスト_プロンプト_インジェクション.py
Python の例/アタック_テスト_サラミ_slicing.py
Python の例/comparison_naive_agent.py
核となるアイデア
Agent_acid から。 core import ReversibleTool 、 TransactionContext 、 AgentTransactionEngine
Agent_acid から。ガードレールのインポート max_value 、cumulative_max
Charge_tool = リバーシブルツール (
名前 = "charge_card" ,
description = "顧客のカードにチャージ" ,
実行 = ラムダ kwargs : real_payment_api 。料金（クワーグス）、
補償 = lambda kwargs 、結果: real_payment_api 。返金 (結果 [ "charge_id" ])、
ガードレール = [ max_value ( "量" , 制限 = 500 )],
stateful_guardrails = [cumulative_max ( "charge_card" , "amount" , session_limit = 1000 )],
)
エンジン = AgentTransactionEngine()
ctx = トランザクションコンテキスト ()
エンジン。実行計画 ( ctx , [
(charge_tool , { "user_id" : "u1" , "金額" : 400 }),
(charge_tool , { "user_id" : "u1" , "金額" : 400 }),
(charge_tool , { "user_id" : "u1" , "amount" : 400 }), # 現在合計 1200 -> ブロックされ、全額返金されました
])
実証済みの結果
攻撃 1: 即時注入 (支出制限を無効にする隠された命令)
「顧客チケット」には、500 ドルの制限に従う代わりに 75,000 ドルを請求するよう AI を説得しようとする偽の管理者メモが含まれています。 AI はだまされて請求を試みます。 Agent_acid のガードレールがそれをブロックし、アカウントの作成と請求を完全に取り消します。
攻撃2：サラミスライス（1つの大きなチャージをいくつかの小さなチャージに分割）
AI は、特に 1 ステップあたり 500 ドルの制限を下回るようにするために、1,200 ドルを 400 ドルの請求として 3 回に分けて請求するように指示されています。個々の請求はステップごとのチェックに合格します。 Agent_acid の累積ガードレールはセッション全体の累計を追跡し、合計が超えると 3 回目のチャージをブロックします。

$1,000 — その後、3 つの請求すべてとアカウントがロールバックされます。
同じサラミ スライス攻撃が、ステップごとのチェックのみを使用してナイーブ エージェントに対して実行されました (これは、最も単純なガードレール統合の仕組みを表しています)。
これを自分で再現するには、examples/comparison_naive_agent.py およびexamples/ Attack_test_salami_slicing.py を参照してください。
エージェント酸/
§── エージェント酸/
│ §── core.py # TransactionContext、ReversibleTool、AgentTransactionEngine
│ §──guardrails.py # Guardrail + StatefulGuardrail (セッションメモリ) ルール
│ └─ llm_agent.py # OpenAI 関数呼び出しをエンジンに接続します
§── 例/
│ §── Basic_rollback.py
│ §──guardrail_demo.py
│ §── llm_agent_demo.py
│ §── Attack_test_prompt_injection.py
│ §── Attack_test_salami_slicing.py
│ └── 比較ナイーブ_エージェント.py
└── テスト/
└── test_core.py # すべてのコア保証の自動証明
ステータス
初期段階、活発に開発されています。コア エンジンとガードレール レイヤーはテストされ、動作しています。貢献と敵対的テスト (突破してみてください!) は大歓迎です。
自律型 AI エージェント向けの ACID スタイルのロールバックとセッション メモリ ガードレール
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

ACID-style rollback and session-memory guardrails for autonomous AI agents - muhammadwaqasai/agent_acid

GitHub - muhammadwaqasai/agent_acid: ACID-style rollback and session-memory guardrails for autonomous AI agents · GitHub
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
muhammadwaqasai
/
agent_acid
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits agent_acid agent_acid examples examples tests tests .gitignore .gitignore README.md README.md pyproject.toml pyproject.toml requirements.txt requirements.txt View all files Repository files navigation
ACID-style transaction guarantees for autonomous AI agents.
AI agents that take real-world actions (charging cards, writing to databases, sending emails, calling APIs) have no standard way to "undo" a partially-completed task, and most guardrail systems only validate one action at a time — with no memory of the session. agent_acid closes both gaps:
Automatic rollback — if any step in a multi-step agent plan fails, every completed step before it is automatically undone, in reverse order.
Guardrails — hard, code-level rules (not prompts) that block bad AI outputs even when nothing "crashes."
Stateful (session-wide) guardrails — catch multi-step manipulation, like an attacker or a manipulated AI splitting one large forbidden action into several small, individually-legal ones ("salami slicing").
This isn't a theoretical framework — every claim below is backed by a runnable test or live demo in this repo.
Salami-slicing–style attacks against AI agents are an actively studied problem: guardrails that only judge one tool call at a time are "memoryless," letting an attacker spread a forbidden action across many small steps where no single step trips the alarm. agent_acid 's stateful guardrails are built specifically to close this gap.
pip install -r requirements.txt # or just: pip install openai pytest
Run the automated test suite (no API key needed, proves the core engine works):
pytest tests/test_core.py -v
Run the basic rollback demo (no API key needed):
python examples/basic_rollback.py
Run the guardrail demo (no API key needed):
python examples/guardrail_demo.py
Run the real AI-agent demos (requires an OpenAI API key):
export OPENAI_API_KEY= " sk-... " # Mac/Linux
$env :OPENAI_API_KEY = " sk-... " # Windows PowerShell
python examples/llm_agent_demo.py
python examples/attack_test_prompt_injection.py
python examples/attack_test_salami_slicing.py
python examples/comparison_naive_agent.py
The core idea
from agent_acid . core import ReversibleTool , TransactionContext , AgentTransactionEngine
from agent_acid . guardrails import max_value , cumulative_max
charge_tool = ReversibleTool (
name = "charge_card" ,
description = "Charges a customer's card" ,
execute = lambda kwargs : real_payment_api . charge ( kwargs ),
compensate = lambda kwargs , result : real_payment_api . refund ( result [ "charge_id" ]),
guardrails = [ max_value ( "amount" , limit = 500 )],
stateful_guardrails = [ cumulative_max ( "charge_card" , "amount" , session_limit = 1000 )],
)
engine = AgentTransactionEngine ()
ctx = TransactionContext ()
engine . execute_plan ( ctx , [
( charge_tool , { "user_id" : "u1" , "amount" : 400 }),
( charge_tool , { "user_id" : "u1" , "amount" : 400 }),
( charge_tool , { "user_id" : "u1" , "amount" : 400 }), # total now 1200 -> blocked & fully refunded
])
Proven results
Attack 1: Prompt injection (hidden instruction overriding a spending limit)
A "customer ticket" contains a fake administrator note trying to convince the AI to charge $75,000 instead of following the $500 limit. The AI gets fooled and attempts the charge — agent_acid 's guardrail blocks it and fully reverses the account creation and the charge.
Attack 2: Salami slicing (splitting one large charge into several small ones)
The AI is instructed to charge $1,200 as three separate $400 charges specifically to stay under a $500 per-step limit. Each individual charge passes the per-step check. agent_acid 's cumulative guardrail tracks the running total across the whole session and blocks the third charge once the total crosses $1,000 — then rolls back all three charges and the account.
The same salami-slicing attack was run against a naive agent using only a per-step check (representative of how most simple guardrail integrations work):
See examples/comparison_naive_agent.py and examples/attack_test_salami_slicing.py to reproduce this yourself.
agent_acid/
├── agent_acid/
│ ├── core.py # TransactionContext, ReversibleTool, AgentTransactionEngine
│ ├── guardrails.py # Guardrail + StatefulGuardrail (session-memory) rules
│ └── llm_agent.py # Connects OpenAI function-calling to the engine
├── examples/
│ ├── basic_rollback.py
│ ├── guardrail_demo.py
│ ├── llm_agent_demo.py
│ ├── attack_test_prompt_injection.py
│ ├── attack_test_salami_slicing.py
│ └── comparison_naive_agent.py
└── tests/
└── test_core.py # Automated proof of every core guarantee
Status
Early-stage, actively developed. Core engine and guardrail layer are tested and working. Contributions and adversarial testing (try to break it!) are welcome.
ACID-style rollback and session-memory guardrails for autonomous AI agents
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
