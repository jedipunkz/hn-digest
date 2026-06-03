---
source: "https://github.com/aerf-spec/aerf"
hn_url: "https://news.ycombinator.com/item?id=48369900"
title: "Show HN: AERF, signed receipts for AI agent actions"
article_title: "GitHub - aerf-spec/aerf: Agent Evidence Receipt Format (AERF) — an open specification for tamper-evident, independently verifiable records of AI agent actions. · GitHub"
author: "keertahacker"
captured_at: "2026-06-03T00:46:46Z"
capture_tool: "hn-digest"
hn_id: 48369900
score: 2
comments: 1
posted_at: "2026-06-02T13:18:19Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AERF, signed receipts for AI agent actions

- HN: [48369900](https://news.ycombinator.com/item?id=48369900)
- Source: [github.com](https://github.com/aerf-spec/aerf)
- Score: 2
- Comments: 1
- Posted: 2026-06-02T13:18:19Z

## Translation

タイトル: HN を表示: AERF、AI エージェントのアクションに対する署名済みの領収書
記事のタイトル: GitHub - aerf-spec/aerf: Agent Evidence Receipt Format (AERF) — AI エージェントのアクションの改ざん明示的で独立して検証可能な記録のためのオープン仕様。 · GitHub
説明: Agent Evidence Receipt Format (AERF) — AI エージェントのアクションの改ざん防止機能があり、独立して検証可能な記録のためのオープン仕様。 - aerf-spec/aerf

記事本文:
GitHub - aerf-spec/aerf: Agent Evidence Receipt Format (AERF) — AI エージェントのアクションの改ざん防止機能があり、独立して検証可能な記録のためのオープン仕様。 · GitHub
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
エアロスペック
/
アーフ
公共
通知
変更するにはサインインする必要があります

通知設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
7 コミット 7 コミット .github/ workflows .github/ workflows docs docs schemas schemas tools tools ベクトル ベクトル verifiers/ go verifiers/ go .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md DECISIONS.md DECISIONS.md ライセンス ライセンスLICENSE-spec LICENSE-spec Makefile Makefile README.md README.md RESOLUTION-PLAN.md RESOLUTION-PLAN.md SPEC.md SPEC.md THREAT-MODEL.md THREAT-MODEL.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AERF — エージェント証拠受領書フォーマット
v0.1.0-draft.1 — パブリック レビュー ドラフト、2026 年 5 月。まだ安定していません。
AERF は、AI エージェントの暗号化受信用のオープン ワイヤ形式です。
アクション。各領収書は、Ed25519 で署名された JSON ドキュメントであり、以下の内容が記録されます。
エージェントが何をしたのか、それを許可したポリシーは何なのか、それがいつ起こったのか、
そしてその行為の完全な証拠。領収書は独立しています
検証可能 — チェックするために AERF ソフトウェア、アカウント、サービスは必要ありません
1つ。このリポジトリの参照検証ツールは、のみを使用する単一の Go ファイルです。
標準ライブラリ。
AERF は AI ガバナンスの証拠とログ層に対応します
フレームワーク。改ざんの証拠と独立した検証可能性を提供します
コントロールの証拠
AIUC-1、
ヒパア、
SOC2、
ISO/IEC 42001、
EU AI法、
NIST AI RMF 、
SR 11-7 、および
SOX 404 。
AERF は完全なコンプライアンス プログラムに代わるものではありません。参照してください
コントロールごとのマッピングについては docs/COMPLIANCE.md
そして明らかなギャップ。
このリポジトリは仕様の本拠地であり、リファレンスです。
検証者。参考プロデューサーが住んでいる
エージェントミント-Python
( pip install Agentmint )。
AERF はエージェント AI 証拠にとって、コンテナ イメージにとっては署名と同じです
そして slsa-verifier は来歴を構築することです: 小さくて退屈で監査可能なもの
小さなボリのあるファイル形式

ng、監査可能な検証者。
{
"id" : " 7473e179-001c-4d3b-94bc-d0f53dd6eec6 " ,
"type" : " 公証証拠 " ,
"plan_id" : " bc023208-ea24-410a-a280-ff36820e18a6 " ,
"エージェント" : "クレームエージェント" ,
"アクション" : " submit:claim:CLM-9920 " ,
"in_policy" : true 、
"policy_reason" : " スコープが一致しました submit:claim:* " ,
"evidence_hash_sha512" : " b35d8ba27ad113c4...bb39e30c " ,
"証拠" : { "..." : " ... " },
"observed_at" : " 2026-05-06T16:22:33.490443+00:00 " ,
"policy_hash" : " 260eca8ac43ae65e...985d6bf1 " ,
"key_id" : " c348d3c785c92249 " ,
"plan_signature" : " 3e5b83e83b77bfa2...dea8ee01 " ,
"署名" : " 8bd989a95ab60863...04e97208 "
}
表示のために省略された証拠フィールド。参照してください
verifiers/go/example/receipt.json
ファイル全体の場合。証拠_ハッシュ_sha512 、ポリシー_ハッシュ 、
plan_signature 、および署名は切り詰められました。
署名は正規の JSON に対する Ed25519 署名です。
署名とタイムスタンプが削除されたレシートのエンコード。
領収書と発行者の公開鍵だけで完全な監査となる
アーティファクト。
git クローン https://github.com/aerf-spec/aerf.git
cd aerf/verifiers/go
verify.go を実行します。 example/receipt.json example/public_key.pem
# → 「OK 受信 7473e179...」と終了コード 0
# 改ざん検出を確認する
verify.go を実行します。 example/receipt-tampered.json example/public_key.pem
# → 「署名検証に失敗しました ...」と終了コード 1
改ざんされたファイルは元のファイルと 1 つのフィールドが異なります
( CLM-9920 → CLM-9921 )。署名はそうではありません。
。
§── README.md ここにあります。
§── SPEC.md AERF-EVIDENCE仕様（案）。
§── DECISIONS.md ロックされた + 保持された設計決定 C-1..C-20。
§── CHANGELOG.md
§── ライセンス Apache 2.0 — コード、スキーマ、サンプル。
§── LICENSE-spec CC BY 4.0 — 散文/仕様テキスト。
§── ドキュメント/
│ §─

─ COMPLIANCE.md コンプライアンス ナビゲーション ハブ。
│ └── フレームワーク/フレームワークごとの AERF マッピング ページ。
§── スキーマ/
│ └── aerf-v0.1.json JSON スキーマ (ドラフト 2020-12)
│ 証拠の領収書の形状。
└── 検証者/
└── 行きます/
§── verify.go ~200 行、stdlib のみ。
§── go.mod
§── README.md
━── 例/
§── 領収書.json
§── 領収書改ざん.json
§── public_key.pem
└──evidence-package.zip エージェントミントの証拠の完全な ZIP。
v0.1.0-draft.1 に含まれないもの
これらは意図的に延期されています。それらは後続のドラフトに反映されます。
テスト ベクター — 〜 8 個の適合ベクターのディレクトリ (genesis、
チェーン、タンパー、リプレイ、不正な形式など)。
Python および TypeScript 参照検証ツール。今日、Python は
参照プロデューサー。このリポジトリには Go ベリファイアのみが同梱されています。
規範的なディレクトリとしての準拠/
非規範的
ドキュメント/フレームワーク マッピング ページ
( AIUC-1 、
ヒパア、
SOC2、
ISO/IEC 42001、
EU AI法、
NIST AI RMF 、
SR11-7、
SOX 404)。
仕様ガバナンス下の規範準拠/ディレクトリ (ロックされている)
決定 C-20) は延期されたままである。
ガバナンス、貢献、およびセキュリティ ポリシーの文書。
CI ワークフローと検証ツールの事前構築済みリリース バイナリ。
AERF-AUTHZ プロファイル — 仕様ではこれが将来のものとして認められています
プロフィール（保留決定C-17）はあるが、それを特定していない。 v0.1 が出荷されます
EVIDENCE プロファイルのみ。
検証者のハッシュ チェーンと RFC 3161 タイムスタンプ チェックを参照します。
どちらも仕様で標準的に説明されています。 Go リファレンス
この草案の検証者は署名のみを強制します。参照
verifiers/go/README.md 。
これはパブリックレビュー草案です。ワイヤ形式は以前に変更される可能性があります
v0.1.0 安定版。ロックされた決定 (DECISIONS.md を参照)
v0.1.0 にバインドされています。保留された決定は、v0.1.0 が安定するまで公開されたままになります。
すべてのページヘッダーにタグを付けます

それに応じて仕様タイトルページも
ここで最終的なものとして引用されるものは何もありません。
プロデューサー: Agentmint-Python — pip install Agentmint
Verifier: このリポジトリ、verifiers/go/
このドラフトの時点での参照プロデューサー (agentmint 0.1.x)
2 つのロックされた決定に関する仕様から逸脱します (ジェネシス センチネル C-6
およびチェーンハッシュ入力 C-7)。 v0.1.0-draft.1 の例は、単一の
ギャップを回避するためのジェネシスレシート。ライブラリの修正は以下で追跡されます
v0.1.0-draft.2 の問題 #2。
保険会社はプランの領収書を機械読み取り可能として使用できます
引受受付および証拠受領書を不正開封防止請求として記録
証拠。付録 B を参照してください。
散文 (仕様、README、決定事項、変更履歴) — クリエイティブ・コモンズ
表示 4.0 インターナショナル (ライセンス仕様)。
コード、スキーマ、サンプル成果物 (検証者、スキーマ、
Recipe.json など) — Apache License 2.0 ( LICENSE )。
ファイルのライセンスが不明瞭な場合は、ディレクトリごとの README.md または
ファイル自体のヘッダーが制御します。
Aniketh Maddipati (agentmint.run)。
問題やPRを歓迎します。
Agent Evidence Receipt Format (AERF) — AI エージェントのアクションの改ざんが明らかで独立して検証可能な記録のためのオープン仕様。
Apache-2.0、不明なライセンスが見つかりました
ライセンスが見つかりました
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
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Agent Evidence Receipt Format (AERF) — an open specification for tamper-evident, independently verifiable records of AI agent actions. - aerf-spec/aerf

GitHub - aerf-spec/aerf: Agent Evidence Receipt Format (AERF) — an open specification for tamper-evident, independently verifiable records of AI agent actions. · GitHub
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
aerf-spec
/
aerf
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
7 Commits 7 Commits .github/ workflows .github/ workflows docs docs schemas schemas tools tools vectors vectors verifiers/ go verifiers/ go .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md DECISIONS.md DECISIONS.md LICENSE LICENSE LICENSE-spec LICENSE-spec Makefile Makefile README.md README.md RESOLUTION-PLAN.md RESOLUTION-PLAN.md SPEC.md SPEC.md THREAT-MODEL.md THREAT-MODEL.md View all files Repository files navigation
AERF — Agent Evidence Receipt Format
v0.1.0-draft.1 — Public Review Draft, May 2026. Not yet stable.
AERF is an open wire format for cryptographic receipts of AI-agent
actions . Each receipt is an Ed25519-signed JSON document that records
what an agent did, what policy permitted it, when it happened,
and the full evidence of the action. Receipts are independently
verifiable — no AERF software, account, or service is needed to check
one. The reference verifier in this repo is a single Go file using only
the standard library.
AERF addresses the evidence and logging layer of AI governance
frameworks. It supplies tamper-evidence and independent-verifiability
evidence for controls in
AIUC-1 ,
HIPAA ,
SOC 2 ,
ISO/IEC 42001 ,
the EU AI Act ,
NIST AI RMF ,
SR 11-7 , and
SOX 404 .
AERF does not replace full compliance programs; see
docs/COMPLIANCE.md for the per-control mapping
and explicit gaps.
This repository is the home of the specification and a reference
verifier. The reference producer lives in
agentmint-python
( pip install agentmint ).
AERF is to agentic AI evidence what cosign is to container images
and slsa-verifier is to build provenance: a small, boring, auditable
file format with a small, boring, auditable verifier.
{
"id" : " 7473e179-001c-4d3b-94bc-d0f53dd6eec6 " ,
"type" : " notarised_evidence " ,
"plan_id" : " bc023208-ea24-410a-a280-ff36820e18a6 " ,
"agent" : " claims-agent " ,
"action" : " submit:claim:CLM-9920 " ,
"in_policy" : true ,
"policy_reason" : " matched scope submit:claim:* " ,
"evidence_hash_sha512" : " b35d8ba27ad113c4...bb39e30c " ,
"evidence" : { "..." : " ... " },
"observed_at" : " 2026-05-06T16:22:33.490443+00:00 " ,
"policy_hash" : " 260eca8ac43ae65e...985d6bf1 " ,
"key_id" : " c348d3c785c92249 " ,
"plan_signature" : " 3e5b83e83b77bfa2...dea8ee01 " ,
"signature" : " 8bd989a95ab60863...04e97208 "
}
evidence field abbreviated for display; see
verifiers/go/example/receipt.json
for the full file. evidence_hash_sha512 , policy_hash ,
plan_signature , and signature truncated.
The signature is an Ed25519 signature over the canonical JSON
encoding of the receipt with signature and timestamp removed.
A receipt by itself plus an issuer's public key is a complete audit
artifact.
git clone https://github.com/aerf-spec/aerf.git
cd aerf/verifiers/go
go run verify.go example/receipt.json example/public_key.pem
# → "OK receipt 7473e179..." and exit code 0
# Confirm tamper detection
go run verify.go example/receipt-tampered.json example/public_key.pem
# → "FAIL signature verification FAILED ..." and exit code 1
The tampered file differs from the original by a single field
( CLM-9920 → CLM-9921 ). The signature does not.
.
├── README.md You are here.
├── SPEC.md The AERF-EVIDENCE specification (draft).
├── DECISIONS.md Locked + held design decisions C-1..C-20.
├── CHANGELOG.md
├── LICENSE Apache 2.0 — code, schemas, examples.
├── LICENSE-spec CC BY 4.0 — prose / specification text.
├── docs/
│ ├── COMPLIANCE.md Compliance navigation hub.
│ └── frameworks/ Per-framework AERF mapping pages.
├── schemas/
│ └── aerf-v0.1.json JSON Schema (Draft 2020-12) for the
│ EVIDENCE receipt shape.
└── verifiers/
└── go/
├── verify.go ~200 lines, stdlib only.
├── go.mod
├── README.md
└── example/
├── receipt.json
├── receipt-tampered.json
├── public_key.pem
└── evidence-package.zip Full agentmint evidence ZIP.
What's NOT in v0.1.0-draft.1
These are intentionally deferred. They will land in subsequent drafts:
Test vectors — a directory of ~8 conformance vectors (genesis,
chain, tamper, replay, malformed, etc.).
Python and TypeScript reference verifiers. Today, Python is the
reference producer ; only the Go verifier ships in this repo.
compliance/ as a normative directory — superseded by the
non-normative
docs/ framework mapping pages
( AIUC-1 ,
HIPAA ,
SOC 2 ,
ISO/IEC 42001 ,
EU AI Act ,
NIST AI RMF ,
SR 11-7 ,
SOX 404 ).
A normative compliance/ directory under spec governance (locked
decision C-20) remains deferred.
Governance, contributing, and security policy documents.
CI workflows and pre-built release binaries of the verifier.
AERF-AUTHZ profile — the spec acknowledges it as a future
profile (held decision C-17) but does not specify it. v0.1 ships
the EVIDENCE profile only.
Reference verifier hash-chain and RFC 3161 timestamp checks.
Both are described normatively in the spec; the Go reference
verifier in this draft enforces signatures only. See
verifiers/go/README.md .
This is a public review draft . The wire format may change before
v0.1.0 stable. Locked decisions (see DECISIONS.md )
are binding for v0.1.0; held decisions remain open until v0.1.0 stable.
We tag every page header and the spec title page accordingly so
nothing here gets cited as final.
Producer: agentmint-python — pip install agentmint
Verifier: This repo, verifiers/go/
The reference producer at the time of this draft ( agentmint 0.1.x)
diverges from the spec on two locked decisions (genesis sentinel C-6
and chain hash input C-7). The v0.1.0-draft.1 example is a single
genesis receipt to sidestep the gap; library fixes are tracked in
issue #2 for v0.1.0-draft.2.
Insurance carriers can use Plan receipts as machine-readable
underwriting intake and Evidence receipts as tamper-evident claims
evidence. See Appendix B.
Prose (SPEC, README, DECISIONS, CHANGELOG) — Creative Commons
Attribution 4.0 International ( LICENSE-spec ).
Code, schemas, example artifacts (verifiers, schemas,
receipt.json, etc.) — Apache License 2.0 ( LICENSE ).
If a file's license is unclear, the per-directory README.md or the
file's own header governs.
Aniketh Maddipati ( agentmint.run ).
Issues and PRs welcome.
Agent Evidence Receipt Format (AERF) — an open specification for tamper-evident, independently verifiable records of AI agent actions.
Apache-2.0, Unknown licenses found
Licenses found
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
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
