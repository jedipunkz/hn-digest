---
source: "https://github.com/wdh107/agent-audit-trail"
hn_url: "https://news.ycombinator.com/item?id=48550486"
title: "AATF – An open spec for recording why AI agents make decisions"
article_title: "GitHub - wdh107/agent-audit-trail: The open specification and reference SDK for recording AI Agent decision chains. Every decision, recorded. Every alternative, documented. · GitHub"
author: "wdh107"
captured_at: "2026-06-16T04:38:09Z"
capture_tool: "hn-digest"
hn_id: 48550486
score: 1
comments: 1
posted_at: "2026-06-16T04:10:27Z"
tags:
  - hacker-news
  - translated
---

# AATF – An open spec for recording why AI agents make decisions

- HN: [48550486](https://news.ycombinator.com/item?id=48550486)
- Source: [github.com](https://github.com/wdh107/agent-audit-trail)
- Score: 1
- Comments: 1
- Posted: 2026-06-16T04:10:27Z

## Translation

タイトル: AATF – AI エージェントが意思決定を行う理由を記録するためのオープン仕様
記事のタイトル: GitHub - wdh107/agent-audit-trail: AI エージェントの意思決定チェーンを記録するためのオープン仕様およびリファレンス SDK。あらゆる決断を記録。すべての代替案が文書化されています。 · GitHub
説明: AI エージェントの意思決定チェーンを記録するためのオープン仕様およびリファレンス SDK。あらゆる決断を記録。すべての代替案が文書化されています。 - wdh107/エージェント監査証跡

記事本文:
GitHub - wdh107/agent-audit-trail: AI エージェントの意思決定チェーンを記録するためのオープン仕様およびリファレンス SDK。あらゆる決断を記録。すべての代替案が文書化されています。 · GitHub
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
wdh107
/
エージェント監査証跡
公共
通知
あなたはそうしなければなりません

サインインして通知設定を変更します
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
4 コミット 4 コミット .github/ workflows .github/ workflows docs docs 例 例 src/ Agent_audit_trail src/ Agent_audit_trail テスト テスト .gitignore .gitignore ライセンス ライセンス README.md README.md SPEC.md SPEC.md pyproject.toml pyproject.toml すべてのファイルを表示リポジトリ ファイルのナビゲーション
エージェント監査証跡フォーマット (AATF)
エージェントのあらゆる決定が記録されます。すべての代替案が文書化されています。
AI エージェントの意思決定チェーンを記録するためのオープン仕様およびリファレンス SDK。
クイック スタート · フォーマット · ブログ: 自己監査ストーリー · なぜ既存のツールを使用しないのか?・仕様・事例
AATF は別のログ ライブラリではありません。これは、AI エージェントがそれぞれの決定を下した理由を記録するためのオープン仕様です。これには、検討した代替案、自信の度合い、実行しないことを選択した内容が含まれます。
OpenTelemetry → 可観測性のため
AATF → エージェントの決定に対する責任
ユーザーは「上海行きのフライトを予約してください」と尋ねます。
ステップ 1: [human_input] → ユーザーリクエストを受信
ステップ 2: [推論] → 目的: 航空券の予約 (信頼度: 0.95)
Alt: ホテルの予約 → 拒否されました (ユーザーは「フライト」と言った)
Alt: 電車の予約 → 拒否されました (ユーザーは「フライト」と言った)
ステップ 3: [tool_call] → Flight_search_api (342ms) → 3 つの結果
ステップ 4: [推論] → 決定: CA1234 2580 円 (確信度: 0.88)
代替案: MU5678 ¥2890 → 不採用 (¥310増)
代替案: CZ9012 ¥3200 → 却下（予算オーバー）
→ SHA-256 ハッシュ チェーン: ✓ 改ざん防止
→ PII 編集: ✓ 電子メール、電話番号、カード番号
→エクスポート：JSON / CSV / HTML（AATF準拠）
クイックスタート（5行）
from agent_audit_trail import AuditSession 、 Decision 、 Alternative
AuditSession (agent_id = "my-agent") をセッションとして使用:
セッション。広告

d_reasoning_step (
名前 = "choose_tool" ,
決定 = 決定 (
input_summary = "ユーザーは天気情報を求めています" ,
決定 = "天気予報 API を使用する" ,
reasoning = "リアルタイム データを必要とする事実のクエリ" ,
信頼度 = 0.95 、
代替案の検討 = [
代替案 ( description = "記憶からの答え" ,
reason_rejected = "天気は常に変化します" ),
代替案 ( description = "説明を求める" ,
reason_rejected = "クエリは十分に明確です" ),
]
)
)
それだけです。すべての決定は、その推論、信頼度スコア、拒否された選択肢とともに AATF 準拠の形式で記録されるようになりました。
AATF の中心となるのは決定記録です。
{
"type" : " 推論 " 、
"名前" : " 意図分類 " ,
「決定」: {
"input_summary" : " ユーザーは上海行きのフライトを予約したいと考えています " ,
"decion" : " フライト予約の意図として分類 " ,
"reasoning" : " 明示的なキーワード: 'フライト' + 目的地 + 予算 " ,
「信頼度」 : 0.95 、
"confidence_basis" : " 3 つのスロットすべてがユーザーによって明示的に指定されました " ,
"alternatives_considered" : [
{
"description" : " ホテルの予約意図 " ,
"reason_rejected" : " ユーザーは「ホテル」ではなく「フライト」と言いました " ,
「スコア」：0.05
}、
{
"description" : " 列車の予約意図 " ,
"reason_rejected" : " ユーザーが明示的に「フライト」と発言しました " ,
「スコア」：0.02
}
]
}、
"step_hash" : " 458942bbf4162f4d9cca121d93b9423413ec... "
}
他のフォーマットでは捉えられない 3 つのこと:
特徴
何をするのか
なぜそれが重要なのか
代替案_検討済み
エージェントが選択しなかったものをリストするよう強制します
エージェントが単に予測された結論を正当化したわけではないことを証明する
信頼 + 信頼_根拠
数値の信頼度 + その決定方法
監査人が「X のため 95% 確実」と「雰囲気のため 95% 確実」を区別できるようにする
自信の軌跡
意思決定チェーン全体にわたって信頼性を追跡します
エージェントが多かれ少なかれ確信を持てるようになった時期を明らかにする

情報を収集するので
なぜ既存のツールを使用しないのか?
私たちは既存のエコシステムを尊重します。 AATF が当てはまるのは次のとおりです。
TL;DR: 他のツールはエージェントの行動を監査します。 AATF はエージェントがなぜそうしたのかを監査します。
#ラングチェーン
Agent_audit_trail から。統合。 langchain import AATFCallbackHandler
エージェント = create_agent (コールバック = [ AATFCallbackHandler()])
#OpenAI
Agent_audit_trail から。統合。 openai インポート AATFOpenAIWrapper
client = AATFOpenAIWrapper (OpenAI())
# 汎用デコレーター (任意のフレームワーク)
Agent_audit_trail から Audit_traced をインポート
@audit_traced (agent_id = "my-agent")
def my_agent_function (クエリ):
「答え」を返す
インストール
pip インストール エージェント監査証跡
外部依存性ゼロ。 Python 3.10以降。 700 行の純粋な stdlib。
私たちは AATF を使用して自分自身、つまり AI エージェントが自身の製品の欠陥を反映していることを監査しました。その結果、改ざんが明らかな 10KB の監査証跡が作成され、すべての推論ステップが本物であり、事後的に合理化されたものではないことが証明されます。
📄 完全な監査証跡の JSON を表示する
AATF はオープン仕様であり、製品ではありません。 SDK はリファレンス実装です。
📋 AATF v0.1.0 仕様の全文を読む
これはドラフト仕様です。皆様からのフィードバックをお待ちしております。設計上の決定に同意できない場合は、問題を提起してください。特に:
alternatives_considered は必須であるべきですか、それともオプションですか?
信頼度 (0.0 ～ 1.0) は正しい抽象概念ですか、それとも定性的なラベルを使用する必要がありますか?
どのようなハッシュ アルゴリズムを標準にする必要がありますか? (現在は SHA-256)
フォーマットは進行中のストリーミング/トレースをサポートする必要がありますか?
役割
得られるもの
エージェント開発者
エージェントの理由をしっかりと証明してください。決定の失敗をデバッグします。関係者にチェーン全体を示します。
コンプライアンス責任者
EU AI法、GDPR、SOC2要件に対応する機械解析可能な監査証跡。
CISO
改ざん防止ハッシュ チェーン。 PII 編集機能が組み込まれています。 au向けにエクスポート

ディターたち。
研究者
エージェントの推論パターンに関する構造化データ。自信の軌跡。意思決定ツリー。
プロジェクトのステータス
✅ リファレンス SDK (Python) — 134 のテストに合格
✅ PII 編集 (電子メール、電話)
✅ ハッシュチェーンの整合性検証
✅ LangChain / OpenAI / 汎用統合
🔲 PII 秘匿化の拡張 (クレジット カード、SSN、API キー、IP)
🔲 仕様変更のためのコミュニティ RFC プロセス
🔲 LangChain/CrewAI 公開プラグイン
このプロジェクトは貢献者を求めています。エージェントの責任を重視する場合:
SPEC を読んでフォーマットを理解する
問題をオープンしてください — 何か同意できないことがありますか?私たちはそれを聞きたいです
統合を構築します — あなたのフレームワークは何ですか?プラグインを歓迎します
情報を広める — スター、ツイート、ブログ投稿
マサチューセッツ工科大学。それを使用し、フォークし、改善してください。スペックはみんなのものです。
エージェントが思考できる場合、その思考は監査可能である必要があります。
AI エージェントの意思決定チェーンを記録するためのオープン仕様およびリファレンス SDK。あらゆる決断を記録。すべての代替案が文書化されています。
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

The open specification and reference SDK for recording AI Agent decision chains. Every decision, recorded. Every alternative, documented. - wdh107/agent-audit-trail

GitHub - wdh107/agent-audit-trail: The open specification and reference SDK for recording AI Agent decision chains. Every decision, recorded. Every alternative, documented. · GitHub
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
wdh107
/
agent-audit-trail
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits .github/ workflows .github/ workflows docs docs examples examples src/ agent_audit_trail src/ agent_audit_trail tests tests .gitignore .gitignore LICENSE LICENSE README.md README.md SPEC.md SPEC.md pyproject.toml pyproject.toml View all files Repository files navigation
Agent Audit Trail Format (AATF)
Every Agent decision, recorded. Every alternative, documented.
The open specification and reference SDK for recording AI Agent decision chains.
Quick Start · The Format · Blog: Self-Audit Story · Why Not Existing Tools? · SPEC · Examples
AATF is not another logging library. It's an open specification for recording why an AI Agent made each decision — including what alternatives it considered, how confident it was, and what it chose not to do.
OpenTelemetry → for observability
AATF → for Agent decision accountability
User asks: "Book a flight to Shanghai"
Step 1: [human_input] → User request received
Step 2: [reasoning] → Intent: flight booking (confidence: 0.95)
Alt: hotel booking → rejected (user said "flight")
Alt: train booking → rejected (user said "flight")
Step 3: [tool_call] → flight_search_api (342ms) → 3 results
Step 4: [reasoning] → Decision: CA1234 at ¥2580 (confidence: 0.88)
Alt: MU5678 at ¥2890 → rejected (¥310 more)
Alt: CZ9012 at ¥3200 → rejected (over budget)
→ SHA-256 hash chain: ✓ tamper-evident
→ PII redaction: ✓ email, phone, card numbers
→ Export: JSON / CSV / HTML (AATF-compliant)
Quick Start (5 Lines)
from agent_audit_trail import AuditSession , Decision , Alternative
with AuditSession ( agent_id = "my-agent" ) as session :
session . add_reasoning_step (
name = "choose_tool" ,
decision = Decision (
input_summary = "User wants weather info" ,
decision = "Use weather API" ,
reasoning = "Factual query requiring real-time data" ,
confidence = 0.95 ,
alternatives_considered = [
Alternative ( description = "Answer from memory" ,
reason_rejected = "Weather changes constantly" ),
Alternative ( description = "Ask for clarification" ,
reason_rejected = "Query is clear enough" ),
]
)
)
That's it. Every decision is now recorded with its reasoning, confidence score, and rejected alternatives — in AATF-compliant format.
The heart of AATF is the Decision record :
{
"type" : " reasoning " ,
"name" : " intent_classification " ,
"decision" : {
"input_summary" : " User wants to book a flight to Shanghai " ,
"decision" : " Classified as flight-booking intent " ,
"reasoning" : " Explicit keywords: 'flight' + destination + budget " ,
"confidence" : 0.95 ,
"confidence_basis" : " All three slots explicitly stated by user " ,
"alternatives_considered" : [
{
"description" : " Hotel booking intent " ,
"reason_rejected" : " User said 'flight', not 'hotel' " ,
"score" : 0.05
},
{
"description" : " Train booking intent " ,
"reason_rejected" : " User explicitly said 'flight' " ,
"score" : 0.02
}
]
},
"step_hash" : " 458942bbf4162f4d9cca121d93b9423413ec... "
}
Three things no other format captures:
Feature
What It Does
Why It Matters
alternatives_considered
Forces agents to list what they didn't choose
Proves the agent didn't just rationalize a foregone conclusion
confidence + confidence_basis
Numeric confidence + how it was determined
Lets auditors distinguish "95% sure because X" from "95% sure because vibes"
confidence_trajectory
Tracks confidence across the full decision chain
Reveals when an agent becomes more or less certain as it gathers information
Why Not Existing Tools?
We respect the existing ecosystem. Here's where AATF fits:
TL;DR: Other tools audit what the agent did . AATF audits why the agent did it .
# LangChain
from agent_audit_trail . integrations . langchain import AATFCallbackHandler
agent = create_agent ( callbacks = [ AATFCallbackHandler ()])
# OpenAI
from agent_audit_trail . integrations . openai import AATFOpenAIWrapper
client = AATFOpenAIWrapper ( OpenAI ())
# Generic decorator (any framework)
from agent_audit_trail import audit_traced
@ audit_traced ( agent_id = "my-agent" )
def my_agent_function ( query ):
return "answer"
Installation
pip install agent-audit-trail
Zero external dependencies. Python 3.10+. 700 lines of pure stdlib.
We used AATF to audit ourselves — an AI Agent reflecting on its own product's flaws. The result is a tamper-evident, 10KB audit trail that proves every reasoning step was genuine and not post-hoc rationalized.
📄 View the full audit trail JSON
AATF is an open specification, not a product. The SDK is the reference implementation.
📋 Read the full AATF v0.1.0 Specification
This is a draft spec. We want your feedback. Open an issue if you disagree with any design decision. Especially:
Should alternatives_considered be mandatory or optional?
Is confidence (0.0-1.0) the right abstraction, or should we use qualitative labels?
What hash algorithm should be standard? (Currently SHA-256)
Should the format support streaming/traces that are still in-progress?
Role
What You Get
Agent Developer
Prove your agent reasons well. Debug decision failures. Show stakeholders the full chain.
Compliance Officer
Machine-parseable audit trails that map to EU AI Act, GDPR, SOC2 requirements.
CISO
Tamper-evident hash chains. PII redaction built-in. Export for auditors.
Researcher
Structured data on agent reasoning patterns. Confidence trajectories. Decision trees.
Project Status
✅ Reference SDK (Python) — 134 tests passing
✅ PII Redaction (email, phone)
✅ Hash Chain Integrity Verification
✅ LangChain / OpenAI / Generic Integrations
🔲 PII Redaction expansion (credit card, SSN, API keys, IP)
🔲 Community RFC process for spec changes
🔲 LangChain/CrewAI published plugins
This project wants contributors. If you care about Agent accountability:
Read the SPEC — understand the format
Open an issue — disagree with something? We want to hear it
Build an integration — your framework? Your plugin welcome
Spread the word — star, tweet, blog post
MIT. Use it, fork it, improve it. The spec belongs to everyone.
If your Agent can think, its thinking should be auditable.
The open specification and reference SDK for recording AI Agent decision chains. Every decision, recorded. Every alternative, documented.
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
