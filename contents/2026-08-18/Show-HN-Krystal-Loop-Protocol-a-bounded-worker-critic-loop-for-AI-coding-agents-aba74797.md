---
source: "https://github.com/KrystalUnity/krystal-loop-protocol"
hn_url: "https://news.ycombinator.com/item?id=49344975"
title: "Show HN: Krystal Loop Protocol–a bounded worker/critic loop for AI coding agents"
article_title: "GitHub - KrystalUnity/krystal-loop-protocol: A bounded build-check-critic-repair protocol for reliable multi-agent software work. · GitHub"
image: "https://opengraph.githubassets.com/72d0b621a0b7ea7bb881ac71084834633251412a435b7f8b42aa4c4a6022efca/KrystalUnity/krystal-loop-protocol"
author: "Eriksz"
captured_at: "2026-08-18T13:36:38Z"
capture_tool: "hn-digest"
hn_id: 49344975
score: 1
comments: 0
posted_at: "2026-08-18T12:58:38Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Krystal Loop Protocol–a bounded worker/critic loop for AI coding agents

- HN: [49344975](https://news.ycombinator.com/item?id=49344975)
- Source: [github.com](https://github.com/KrystalUnity/krystal-loop-protocol)
- Score: 1
- Comments: 0
- Posted: 2026-08-18T12:58:38Z

## Translation

タイトル: Show HN: Krystal Loop Protocol – AI コーディング エージェント用の制限されたワーカー/クリティカル ループ
記事のタイトル: GitHub - KrystalUnity/krystal-loop-protocol: 信頼性の高いマルチエージェント ソフトウェア作業のための制限されたビルド-チェック-クリティカル-修復プロトコル。 · GitHub
説明: 信頼性の高いマルチエージェント ソフトウェア作業のための、制限されたビルド-チェック-クリティカル-修復プロトコル。 - KrystalUnity/クリスタルループプロトコル

記事本文:
GitHub - KrystalUnity/krystal-loop-protocol: 信頼性の高いマルチエージェント ソフトウェア作業のための、制限付きの build-check-critic-repair プロトコル。 · GitHub
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
クリスタルユニティ
/
クリスタルループプロトコル
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
15 コミット 15 コミット フォルダーとファイル
.github/ ワークフロー .github/ ワークフロー アダプター アダプター

s 資産 アセットの例 例 .gitignore .gitignore AI_AGENT_INTEGRATION.md AI_AGENT_INTEGRATION.md LICENSE LICENSE LICENSE-APACHE-2.0 LICENSE-APACHE-2.0 LICENSE-CC-BY-4.0 LICENSE-CC-BY-4.0 PROTOCOL.md PROTOCOL.md README.md README.md RUNTIME_PROFILES.md RUNTIME_PROFILES.md SAFETY_AND_LIMITS.md SAFETY_AND_LIMITS.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
面倒な作業には高速 AI エージェントを使用します。主要なエージェントと本物の小切手を使用して、
プロジェクトは一貫性があり、機能しており、あなたの制御下にあります。
AI コーディング エージェントは、大量のコードを迅速に生成できます。より難しい問題が始まります
最初の印象的なデモの後: エージェントはコンテキストを失い、お互いのコンテキストを重ね合わせます。
変更を加え、自信に満ちた要約を信頼し、解決した問題を再び開き、静かに打ち切る
昨日機能した機能。
Krystal Loop Protocol (KLP) は、継続的に実行するための実用的な動作パターンです。
複数の AI エージェントにプロジェクトの制御を渡さずに構築します。
範囲を設定します。結果、許可されたファイル、禁止されたアクション、
チェック、予算、および停止条件。
分割してください。各従業員に、その結果に基づいて判断できる小さな成果を 1 つ与えます。
自分のもの。
構築してください。作業者は正確なリビジョンと事実に基づく引き継ぎを返します。
独自の合格判定ではありません。
確認してください。テスト、リンター、ビルド、またはその他の決定的チェックを実行する
他のモデルにどう思うかを尋ねる前に。
それを批判してください。新鮮な読み取り専用の批評家が、正確な改訂版をレビューします。
元のタスクと直接の証拠。
判断して修理します。それぞれの批評家の発見を確認または反論します。
直接証拠を確認し、宣言された制限内で確認されたブロッカーを修正します。
それを統合します。結合されたシステムを新しい成果物として扱い、確認します
またまた。
それを統治してください。ループはその委任内で自動的に完了します。あ
人は重要な変化とその結果としての行動のみを決定します。
一般的なエージェントの障害
KLPの対応

e
2 人のエージェントが同じ共有ファイルを編集します。
各ワーカーに明示的なファイルとアクションの境界を与えます。
作業者は自分の作業が完了したと言います。
事実の引き渡しと独立した受諾を区別する。
批評家が古いビルドをレビューします。
すべての判定を正確なアーティファクト リビジョンにバインドします。
批評家が自信を持って誤った問題を報告している。
すべての発見を直接証拠で確認または反論します。
エージェントは改善されずにループし続けます。
繰り返しの発見、使い果たされた予算、または停滞状態で停止します。
ユニットの変更は成功しますが、組み合わせると壊れます。
新しい成果物としての統合を確認します。
テストまたはエージェントのメッセージは、展開の許可として扱われます。
ライブアクションは人間による明示的な決定の背後に置いてください。
1 つのプロンプトから始める
マルチエージェントを構築する前に、これをリード エージェントに渡します。
Krystal Loop プロトコル コアの下で動作します。
ファイルを変更する前に、以下を含むバインドされたタスク コントラクトを作成します。
- その ID、リビジョン、親 ID、およびコーディネーターの委任。
- 正確な結果;
- 許可されたパスと保護されたパス。
- 禁止された行為と実際の副作用;
- 決定論的なチェックと必要な証拠。
- 時間、コスト、修理回数の制限。
- 人間による確認のために停止する必要がある状況。
タスクを独立して判断可能な作業単位に分割します。労働者は戻らなければならない
事実の引き渡しは正確な改訂に結び付けられており、自身の作業を証明してはなりません。
独立した読み取り専用の批判の前に、決定論的なチェックを実行します。確認して、
直接証拠を使用して反論するか、各批判者の発見を未解決のままにしておきます。治療する
統合された結果が新しいアーティファクトとして生成されます。通常の修理ラウンドは以下の条件で続行できます。
記録された代表団。重大な変化とその結果としての行動は次の段階で停止します。
権限の境界。
KLP には、特別なモデル、データベース、ベクトル ストア、またはメッセージ バスは必要ありません。
Git および Markdown ファイルから開始し、その後、耐久性のある調整を追加できます。
広報

オブジェクトにはそれが必要です。
含まれている例は、1 つの完全な境界パスを示しています。
DeepSeek Flash ワーカー
-> 正確なリビジョンと確定的なチェック
-> コントローラーによって編集され、封印されたレビュー パケット
-> 異系読み専評論家
-> 証拠に裏付けられた調査結果
-> 限定的な修復または統合
DeepSeek Codex ワーカーから始めて、
次に、を使用します
OpenAI 互換のクリティカル ハーネス。
ワーカー ランチャーは、分離された Codex プロファイルとレビューされた割り当てを使用します。
批評家は 1 つのリクエストを行い、その構造化された評決をローカルで検証し、
アクションを編集したり承認したりすることはできません。例は自動的に接続しません
労働者の出力から批評家の入力へ。コントローラーは検査して編集する必要があります
まずパケットを確認してください。
オフラインライフサイクルプルーフを実行する
失敗レビュー修復フィクスチャは、
プロバイダーのエグレスやモデルの支出なしで、制限されたライフサイクルを完了します。
出力 = " $( mktemp -d ) /klp-fixture "
python3 例/fail-review-repair/run_fixture.py --out-dir " $output "
python3 -m json.tool " $output /final-receipt.json "
一時的な Git プロジェクトを作成し、意図的に不完全なワーカーを記録します
リビジョン、失敗したチェックをキャプチャし、封印された批評家ハーネスを
ループバック偽プロバイダー、1 つの本当の発見を確認し、1 つの誤った発見を否定します。
承認された修理を 1 回適用し、統合されたアーティファクトを再チェックします。 2番目
テストでは、権限を超過するのではなく、ゼロラウンドの修理予算が停止されることが証明されています。
AI エージェントとオーケストレーターのための手順
OpenAI互換の批評家の例
オフラインの失敗レビュー修復フィクスチャ
KLP は自律的なソフトウェア ファクトリ、展開プラットフォーム、モデル ルーターではありません。
または、テストにより製品に欠陥がないことが証明されたと主張します。ポータブルな方法です
マルチエージェントの作業を拘束し、有用な証拠を保持し、不確実性を可視化します。
KLP は公共の建設者/批評家システムからインスピレーションを受けています

のような
ガントレットループ。と主張するものではありません
マルチエージェントコーディング、独立したレビュー、自動テストを発明します。その焦点
これはデモの後に起こることであり、実際のプロジェクトは理解できる状態でなければなりません
そして多くの変更に対処しています。
KLP Core v0.2 は暫定的な公開プロファイルです。このリポジトリには、
ポータブル ワーカー ランチャー、密封パケット批評家ハーネス、ローカルの偽プロバイダー
テスト、契約例、および再現可能なオフラインの失敗レビュー修復
固定具。このフィクスチャーはプロトコルの仕組みを証明します。モデルを認定するものではありません
品質、生産の安全性、または特定のエージェントのフレームワーク。
Krystal Loop プロトコルは分割ライセンスを使用します。
ソフトウェア、スクリプト、スキーマ、構成例、テストにはライセンスが必要です
Apache License 2.0 ( Apache-2.0 ) に基づく;
プロトコルとドキュメントのテキストは以下に基づいてライセンスされています。
クリエイティブ コモンズ表示 4.0 ( CC-BY-4.0 )。
ファイル レベルの境界については、ライセンス マップを参照してください。
信頼性の高いマルチエージェント ソフトウェア作業のための、制限されたビルド-チェック-クリティカル-修復プロトコル。
Readme ライセンスとその他 2 つのライセンスが見つかりました アクティビティ カスタム プロパティ 星
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A bounded build-check-critic-repair protocol for reliable multi-agent software work. - KrystalUnity/krystal-loop-protocol

GitHub - KrystalUnity/krystal-loop-protocol: A bounded build-check-critic-repair protocol for reliable multi-agent software work. · GitHub
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
KrystalUnity
/
krystal-loop-protocol
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
15 Commits 15 Commits Folders and files
.github/ workflows .github/ workflows adapters adapters assets assets examples examples .gitignore .gitignore AI_AGENT_INTEGRATION.md AI_AGENT_INTEGRATION.md LICENSE LICENSE LICENSE-APACHE-2.0 LICENSE-APACHE-2.0 LICENSE-CC-BY-4.0 LICENSE-CC-BY-4.0 PROTOCOL.md PROTOCOL.md README.md README.md RUNTIME_PROFILES.md RUNTIME_PROFILES.md SAFETY_AND_LIMITS.md SAFETY_AND_LIMITS.md View all files Repository files navigation
Use fast AI agents for the grunt work. Use a lead agent and real checks to keep
the project coherent, working, and under your control.
AI coding agents can produce a lot of code quickly. The harder problem begins
after the first impressive demo: agents lose context, overlap each other's
changes, trust confident summaries, reopen solved problems, and quietly break
features that worked yesterday.
Krystal Loop Protocol (KLP) is a practical operating pattern for continuing to
build with multiple AI agents without handing them control of the project.
Scope it. Write down the outcome, allowed files, forbidden actions,
checks, budget, and stop conditions.
Split it. Give each worker one small outcome that can be judged on its
own.
Build it. The worker returns the exact revision and a factual handover,
not its own pass verdict.
Check it. Run tests, linters, builds, or other deterministic checks
before asking another model what it thinks.
Critique it. A fresh, read-only critic reviews the exact revision against
the original task and direct evidence.
Adjudicate and repair it. Confirm or refute each critic finding with
direct evidence, then fix confirmed blockers within declared limits.
Integrate it. Treat the combined system as a new artifact and check it
again.
Govern it. The loop completes automatically inside its delegation. A
person decides only material changes and consequential actions.
Common agent failure
KLP response
Two agents edit the same shared file.
Give each worker an explicit file and action boundary.
A worker says its own work is complete.
Separate factual handover from independent acceptance.
A critic reviews an outdated build.
Bind every verdict to an exact artifact revision.
A critic confidently reports a false problem.
Confirm or refute every finding with direct evidence.
Agents keep looping without improvement.
Stop on repeated findings, exhausted budgets, or a plateau.
Unit changes pass but break when combined.
Review integration as a new artifact.
A test or agent message is treated as permission to deploy.
Keep live actions behind an explicit human decision.
Start With One Prompt
Give this to the lead agent before a multi-agent build:
Work under Krystal Loop Protocol Core.
Before changing files, write a bounded task contract containing:
- its ID, revision, parent identity, and coordinator delegation;
- the exact outcome;
- allowed and protected paths;
- forbidden actions and live side effects;
- deterministic checks and required evidence;
- time, cost, and repair-round limits;
- conditions that require stopping for human review.
Split the task into independently judgeable work units. Workers must return
factual handovers tied to exact revisions and must not certify their own work.
Run deterministic checks before independent, read-only criticism. Confirm,
refute, or leave each critic finding unresolved using direct evidence. Treat
the integrated result as a new artifact. Normal repair rounds may proceed under
the recorded delegation; material changes and consequential actions stop at
the authority boundary.
KLP does not require a special model, database, vector store, or message bus.
You can start with Git and Markdown files, then add durable coordination when
the project needs it.
The included examples demonstrate one complete bounded path:
DeepSeek Flash worker
-> exact revision and deterministic checks
-> controller-redacted, sealed review packet
-> different-family read-only critic
-> evidence-backed finding dispositions
-> bounded repair or integration
Start with the DeepSeek Codex worker ,
then use the
OpenAI-compatible critic harness .
The worker launcher uses an isolated Codex profile and a reviewed assignment.
The critic makes one request, validates its structured verdict locally, and
cannot edit or authorize any action. The examples do not automatically connect
worker output to critic input; the controller must inspect and redact the
review packet first.
Run the Offline Lifecycle Proof
The fail-review-repair fixture proves a
complete bounded lifecycle without provider egress or model spend:
output= " $( mktemp -d ) /klp-fixture "
python3 examples/fail-review-repair/run_fixture.py --out-dir " $output "
python3 -m json.tool " $output /final-receipt.json "
It creates a temporary Git project, records a deliberately incomplete worker
revision, captures a failed check, runs the sealed critic harness against a
loopback fake provider, confirms one real finding, refutes one false finding,
applies one authorized repair, and re-checks the integrated artifact. A second
test proves a zero-round repair budget stops instead of overrunning authority.
Instructions for AI agents and orchestrators
OpenAI-compatible critic example
Offline fail-review-repair fixture
KLP is not an autonomous software factory, deployment platform, model router,
or claim that tests prove a product has no defects. It is a portable way to
bound multi-agent work, retain useful evidence, and make uncertainty visible.
KLP is inspired by public builder/critic systems such as the
Gauntlet Loop . It does not claim to
invent multi-agent coding, independent review, or automated testing. Its focus
is what happens after the demo, when a real project must remain understandable
and working across many changes.
KLP Core v0.2 is a provisional public profile. This repository includes a
portable worker launcher, a sealed-packet critic harness, local fake-provider
tests, example contracts, and a reproducible offline fail-review-repair
fixture. The fixture proves protocol mechanics; it does not certify model
quality, production safety, or a particular agent framework.
Krystal Loop Protocol uses split licensing:
software, scripts, schemas, configuration examples, and tests are licensed
under Apache License 2.0 ( Apache-2.0 );
protocol and documentation text are licensed under
Creative Commons Attribution 4.0 ( CC-BY-4.0 ).
See the licensing map for the file-level boundary.
A bounded build-check-critic-repair protocol for reliable multi-agent software work.
Readme License and 2 other licenses found Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
