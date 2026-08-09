---
source: "https://github.com/Rehanrana11/evidence-verify"
hn_url: "https://news.ycombinator.com/item?id=49230728"
title: "Show HN: Would your AI audit logs survive an audit challenge?"
article_title: "GitHub - Rehanrana11/evidence-verify: Verify AI audit logs: chain integrity, tamper detection, and the independence check most vendors fail. EU AI Act Article 12. · GitHub"
author: "rmasoodx22"
captured_at: "2026-08-09T12:29:37Z"
capture_tool: "hn-digest"
hn_id: 49230728
score: 1
comments: 0
posted_at: "2026-08-09T12:17:10Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Would your AI audit logs survive an audit challenge?

- HN: [49230728](https://news.ycombinator.com/item?id=49230728)
- Source: [github.com](https://github.com/Rehanrana11/evidence-verify)
- Score: 1
- Comments: 0
- Posted: 2026-08-09T12:17:10Z

## Translation

タイトル: Show HN: あなたの AI 監査ログは監査の課題に耐えられますか?
記事のタイトル: GitHub - Rehanrana11/evidence-verify: AI 監査ログを検証する: チェーンの整合性、改ざん検出、およびほとんどのベンダーが失敗する独立性チェック。 EU AI 法第 12 条 · GitHub
説明: AI 監査ログを検証します: チェーンの整合性、改ざん検出、およびほとんどのベンダーが失敗する独立性チェック。 EU AI 法第 12 条 - Rehanrana11/evidence-verify

記事本文:
GitHub - Rehanrana11/evidence-verify: AI 監査ログを検証します: チェーンの整合性、改ざん検出、およびほとんどのベンダーが失敗する独立性チェック。 EU AI 法第 12 条 · GitHub
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
リハンラナ11
/
証拠検証
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
5 コミット 5 コミット ライセンス ライセンス README.md README.md Battery-forge-meta-prompt.md Battery-forge-meta-prompt.md Battery.p

y Battery.pydemo_anchor.jsondemo_anchor.jsonevidence_verify.pyevidence_verify.pymake_samples.pymake_samples.pysample_anchored.jsonlsample_anchored.jsonlsample_selfattested.jsonlsample_selfattested.jsonlsample_tampered.jsonlsample_tampered.jsonl すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI 監査ログでは解決できない質問に答えるコマンドが 1 つあります。
このログがそれを作成したシステムによって変更されていないことを誰が保証しますか?
$ 証拠検証受領書.jsonl
チェックされたエントリ: 5
チェーン判定: CHAIN VALID
独立性 : 自己証明 — 独立性の挑戦を生き延びることはできません。
このチェーン内のすべてのハッシュは、同じものによって生成および保存されています。
それが説明するシステム。侵害されたプロデューサーは、
チェーン全体。外部アンカーが見つかりません。
なぜこれが存在するのか
EU AI 法の第 12 条では、高リスク AI の自動イベントログが義務付けられています
システム（附属書 III の義務は、2026 年 7 月のオムニバス改正に基づいて 2027 年 12 月 2 日から適用されるようになりました。条項に基づく保持期間は 6 か月以上）
19/26(6))。ほとんどのベンダーはハッシュチェーンされたログで答えます。
ハッシュチェーンはシーケンスを証明します。それは独立性を証明するものではありません。もし
ログを書き込むシステムが危険にさらされている、または単純に動機付けられている可能性があります。
履歴を書き換えてチェーン全体に再署名します。ログは日記ではなく、
宣誓供述書。執行者が公証人になることはできません。
証拠検証では 3 つのレベルがチェックされます。
CHAIN — すべてのエントリは前のエントリにリンクされていますか?
タンパー — すべてのエントリのコンテンツは依然としてそのハッシュと一致していますか?
INDEPENDENCE — チェーンは外部のアンカーを参照しますか?
生産システム？そうでない場合は、はっきりとそう述べています。
ログを含む、以下の受信形式に従うあらゆるログで機能します。
他のゲートウェイやプロキシから。現在のベンダーのエクスポートをポイントして、
監査人が何をみるか見てみましょう。
ビルド -o 証拠に進みます。
./証拠検証サンプル

_selfattested.jsonl
Python リファレンス実装 (同一のロジック、クロステストに使用):
python3evidence_verify.pysample_selfattested.jsonl
終了コード
0 = チェーンは有効でアンカーされています (--anchor が指定されている場合はラウンドトリップが検証されます)。 1 = 自己認証、不一致、または壊れているか改ざんされています。
2 = 使用法またはファイルのエラー。 CI は 0 でゲートできます。
領収書のフォーマット (v0.1 — 仕様 RFC が予定されています)
1 行に 1 つの JSON オブジェクト (JSONL):
{
"シーケンス" : 1 、
"ts" : " 2026-08-09T10:01:00Z " ,
"イベント" : { "モデル" : " ... " 、 "request_sha256" : " ... " 、 "response_sha256" : " ... " 、 "決定" : " 許可 " },
"prev_hash" : " <前のエントリの sha256; ジェネシスの 64 個のゼロ> " ,
"entry_hash" : " <sha256 of seq|ts|canonical(event)|prev_hash> " ,
"アンカー" : { "タイプ" : " 外部 " 、 "プロバイダー" : " ... " 、 "ref" : " ... " }
}
canonical(event) = ソートされたキーを含む、空白なしの JSON (RFC 8785 スタイル)。
Go と Python の実装はバイト互換です。
チェーンを固定します (出力を生成システムの外側に保存します - 別個の
リポジトリ、Gist、S3、タイムスタンプ サービス):
pythonevidence_verify.pyアンカー領収書.jsonl >アンカー.json
アンカー (ローカル ファイルまたは https URL) に対して検証します。
pythonevidence_verify.py verify reports.jsonl --anchor アンカー.json
pythonevidence_verify.py verify reports.jsonl --anchor https://example.com/anchor.json
ラウンドトリップ検証は、ハッシュ チェーンだけではできないものをキャッチします: 完全な履歴
書き換え (再署名されたチェーン) と切り捨て。独立性の保証は、
アンカーが存在する場所の独立性とまったく同じくらい強力です。
Sample_selfattested.jsonl — 有効なチェーン、アンカーなし (ほとんどのベンダーが出荷するもの)
Sample_anchored.jsonl — 外部アンカー参照を含む有効なチェーン
sample_tampered.jsonl — 1 つのエントリがサイレント編集されます。検証者はそれをキャッチします
python3 make_samples.py で再生成します。
Ed25519 サイン

性質の検証 (非対称、サードパーティによる検証可能)
eIDAS 互換のアンカーの適格タイムスタンプ検証
アダプター: Kong プラグインのエクスポート、LangChain コールバック、LiteLLM、OpenAI
コンプライアンス ログのエクスポート (デフォルトでは最大 30 日間保存されます)
6 か月の保持ギャップ）
Apache-2.0。検証は永久に無料です。それがポイントです。
AI 監査ログを検証します: チェーンの整合性、改ざん検出、およびほとんどのベンダーが失敗する独立性チェック。 EU AI 法第 12 条。
Readme Apache-2.0 ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Verify AI audit logs: chain integrity, tamper detection, and the independence check most vendors fail. EU AI Act Article 12. - Rehanrana11/evidence-verify

GitHub - Rehanrana11/evidence-verify: Verify AI audit logs: chain integrity, tamper detection, and the independence check most vendors fail. EU AI Act Article 12. · GitHub
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
Rehanrana11
/
evidence-verify
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
5 Commits 5 Commits LICENSE LICENSE README.md README.md battery-forge-meta-prompt.md battery-forge-meta-prompt.md battery.py battery.py demo_anchor.json demo_anchor.json evidence_verify.py evidence_verify.py make_samples.py make_samples.py sample_anchored.jsonl sample_anchored.jsonl sample_selfattested.jsonl sample_selfattested.jsonl sample_tampered.jsonl sample_tampered.jsonl View all files Repository files navigation
One command that answers the question your AI audit logs can't:
who guarantees this log wasn't altered by the system that produced it?
$ evidence verify receipts.jsonl
entries checked : 5
chain verdict : CHAIN VALID
independence : SELF-ATTESTED — would not survive independence challenge.
Every hash in this chain was produced and stored by the same
system it describes; a compromised producer can re-sign the
entire chain. No external anchor found.
Why this exists
The EU AI Act's Article 12 requires automatic event logging for high-risk AI
systems (Annex III obligations now apply from December 2, 2027 under the July 2026 Omnibus amendment; retention ≥6 months under Articles
19/26(6)). Most vendors answer with hash-chained logs.
A hash chain proves sequence . It does not prove independence . If the
system that writes the log is compromised — or simply motivated — it can
rewrite history and re-sign the entire chain. The log is a diary, not an
affidavit. The enforcer cannot be the notary.
evidence verify checks three levels:
CHAIN — is every entry linked to the previous one?
TAMPER — does every entry's content still match its hash?
INDEPENDENCE — does the chain reference an anchor outside the
producing system? If not, it says so, plainly.
It works on any log that follows the receipt format below — including logs
from other gateways and proxies. Point it at your current vendor's export and
see what an auditor would see.
go build -o evidence .
./evidence verify sample_selfattested.jsonl
Python reference implementation (identical logic, used for cross-testing):
python3 evidence_verify.py sample_selfattested.jsonl
Exit codes
0 = chain valid and anchored (round-trip verified when --anchor is given). 1 = self-attested, mismatch, or broken/tampered.
2 = usage or file error. CI can gate on 0 .
Receipt format (v0.1 — spec RFC coming)
One JSON object per line (JSONL):
{
"seq" : 1 ,
"ts" : " 2026-08-09T10:01:00Z " ,
"event" : { "model" : " ... " , "request_sha256" : " ... " , "response_sha256" : " ... " , "decision" : " allow " },
"prev_hash" : " <sha256 of previous entry; 64 zeros for genesis> " ,
"entry_hash" : " <sha256 of seq|ts|canonical(event)|prev_hash> " ,
"anchor" : { "type" : " external " , "provider" : " ... " , "ref" : " ... " }
}
canonical(event) = JSON with sorted keys, no whitespace (RFC 8785-style).
The Go and Python implementations are byte-compatible.
Anchor a chain (store the output OUTSIDE the producing system — a separate
repo, a gist, S3, a timestamping service):
python evidence_verify.py anchor receipts.jsonl > anchor.json
Verify against the anchor (local file or https URL):
python evidence_verify.py verify receipts.jsonl --anchor anchor.json
python evidence_verify.py verify receipts.jsonl --anchor https://example.com/anchor.json
Round-trip verification catches what hash chains alone cannot: full history
rewrites (re-signed chains) and truncation. The independence guarantee is
exactly as strong as the independence of where the anchor lives.
sample_selfattested.jsonl — valid chain, no anchor (what most vendors ship)
sample_anchored.jsonl — valid chain with external anchor references
sample_tampered.jsonl — one entry silently edited; the verifier catches it
Regenerate with python3 make_samples.py .
Ed25519 signature verification (asymmetric, third-party verifiable)
eIDAS-compatible qualified timestamp validation for anchors
Adapters: Kong plugin export, LangChain callback, LiteLLM, OpenAI
compliance-log export (which retains ~30 days by default — mind the
6-month retention gap)
Apache-2.0. Verification stays free forever. That's the point.
Verify AI audit logs: chain integrity, tamper detection, and the independence check most vendors fail. EU AI Act Article 12.
Readme Apache-2.0 license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
