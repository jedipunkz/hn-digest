---
source: "https://github.com/flexorch/flexorch-audit"
hn_url: "https://news.ycombinator.com/item?id=48584670"
title: "Show HN: Flexorch-audit – quality scoring and PII detection for LLM pipelines"
article_title: "GitHub - flexorch/flexorch-audit: Zero-dependency PII + quality + noise audit for LLM datasets (TR/EU/US) · GitHub"
author: "flexorch"
captured_at: "2026-06-18T13:17:06Z"
capture_tool: "hn-digest"
hn_id: 48584670
score: 1
comments: 0
posted_at: "2026-06-18T13:03:18Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Flexorch-audit – quality scoring and PII detection for LLM pipelines

- HN: [48584670](https://news.ycombinator.com/item?id=48584670)
- Source: [github.com](https://github.com/flexorch/flexorch-audit)
- Score: 1
- Comments: 0
- Posted: 2026-06-18T13:03:18Z

## Translation

タイトル: Show HN: Flexorch-audit – LLM パイプラインの品質スコアリングと PII 検出
記事のタイトル: GitHub - flexorch/flexorch-audit: LLM データセットの依存性ゼロの PII + 品質 + ノイズ監査 (TR/EU/US) · GitHub
説明: LLM データセットの依存性のない PII + 品質 + ノイズ監査 (TR/EU/US) - flexorch/flexorch-audit

記事本文:
GitHub - flexorch/flexorch-audit: LLM データセットの依存性ゼロの PII + 品質 + ノイズ監査 (TR/EU/US) · GitHub
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
フレクコーチ
/
フレクスコープ監査
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション操作

ション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
20 コミット 20 コミット .github/ workflows .github/ workflows アセット アセット デモ デモの例 例 src/ flexorch_audit src/ flexorch_audit テスト テスト .gitignore .gitignore CHANGELOG.md CHANGELOG.md COTRIBUTING.md CONTRIBUTING.md ライセンス ライセンスLIMITATIONS.md LIMITATIONS.md README.md README.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
依存関係のない PII 検出、品質グレーディング、LLM データセットのノイズ監査を 1 回の関数呼び出しで実行します。
ドキュメントを LLM パイプラインにフィードする前に、次の 3 つの質問に答える必要があります。
このテキストには個人データが含まれていますか? PII を言語モデルに送信すると、コンプライアンスのリスクが生じます。
テキストの品質は十分に高いですか?短いレコード、ノイズの多いレコード、または重複したレコードは、微調整や RAG の取得に悪影響を及ぼします。
騒音はどれくらいひどいですか?文字化けしたエンコーディングとシンボルの乱雑さにより、モデルの出力が静かに低下します。
これらの質問に答えるほとんどのツールには、重い NLP フレームワーク、モデルの重み付け、またはクラウド API が必要です。 flexorch-audit は、正規表現と Python の標準ライブラリのみを使用して、1 回の呼び出しで 3 つすべてに応答します。モデルの重み付け、ネットワーク呼び出し、外部パッケージはありません。
品質グレード — A/B/C/D 複合スコア: このテキストは一目で LLM 対応ですか?
ノイズ比 — ラインレベルのシンボルクラッター検出 (noise_ratio); 0.20 を超える値は、抽出アーティファクトの可能性があることを示します
PII 検出 — 8 か国 (TR/DE/FR/IT/NL/ES/UK/US) の 30 以上のタイプ + ユニバーサル タイプ。すべて正規表現ベースでチェックサム検証あり
バッチ監査 — Audit_batch() は、1 回の呼び出しでデータセット全体の重複率と PII カウントを集計します。
マスキング — 4 つの戦略: 編集、置換 (合成)、トークン、ハッシュ
実行時の依存関係なし — 純粋な Python stdlib、Python 3.10 以降
pip インストール flexorch-audit

クイックスタート
flexorch_audit インポート監査、マスクから
text = 開く ( "contract.txt" )。 read () # まず PDF/DOCX から抽出します
result = Audit ( text ) # デフォルトでは「und」 — すべての検出器がアクティブになります
# result = Audit(text, locale="tr") # TR のみの検出器に制限する
結果。品質_グレード # "B"
結果。品質スコア # 0.73 (0.0–1.0 複合)
結果。ノイズ比 # 0.04 (空白行/ゴミ行の割合; >0.20 = 低品質)
結果。 detect_ language # "und" (渡したロケール、呼び出し元が言語を制御)
結果。 pii_summary # [{"タイプ": "電子メール", "カウント": 2}, {"タイプ": "national_id_tr", "カウント": 1}]
# 完全な結果と生の指標 — dict アクセスも機能します。
result [ "pii" ] # [{"type": "email", "value": "ali@example.com", "start": 8, "end": 23}]
result [ "quality" ] # {"completeness": 1.0, "avg_length": 342, "duplicate_ratio": None}
result [ "ノイズ" ] # {"garbage_ratio": 0.0, "encoding_ok": True}
clean = マスク (テキスト、結果 [ "pii" ]、戦略 = "redact" )
# 「連絡先: [REDACTED_EMAIL]」
flexorch_audit から redact_for_llm をインポート
clean = redact_for_llm ( "TCKN: 12345678950、電子メール: ali@example.com" 、ロケール = "tr" )
# "TCKN: [REDACTED_NATIONAL_ID_TR]、メールアドレス: [REDACTED_EMAIL]"
# さまざまなマスキング戦略
redact_for_llm (テキスト、ロケール = "tr"、戦略 = "トークン") # <PII_NATIONAL_ID_TR_1>
redact_for_llm (テキスト、ロケール = "tr"、戦略 = "ハッシュ") # [3d4f9a1b2c8e7f0a]
redact_for_llm ( text , locale = "tr" , Strategy = "replace" ) # 静的な合成値
PII が見つからない → 元のテキストが変更されずに返されました。
flexorch_audit インポートestimate_tokensから
estimate_tokens (「機敏な茶色のキツネが怠惰な犬を飛び越えます。」) # → 16
推定トークン ( "" ) # → 0
ヒューリスティック: 単語 × 4/3 — tiktoken は必要ありません。英語およびほとんどのヨーロッパの言語の実際のトークナイザーの精度は 15% 以内です

数値。コンテキスト ウィンドウのサイジングとコスト予測の計画見積もりとして扱います。
from flexorch_audit import audit_batch
texts = [ データセット内のレコードのレコード [ "テキスト" ]
バッチ = Audit_batch ( texts ) # デフォルトでは locale="und"
バッチ [ "duplicate_ratio" ] # 0.12 — 完全に重複するレコードの割合
batch [ "avg_quality_score" ] # 0.78
バッチ [ "pii_summary" ] # [{"type": "email", "count": 47}, ...]
バッチ [ "results" ] # AuditResult のリスト、テキストごとに 1 つ
対象国
ロケール
検出器が作動しました
「および」 (デフォルト)
すべてのロケールを組み合わせた - ドキュメントの言語が不明な場合に使用します
「すべて」
「und」の別名
「とる」
TCKN · VKN · 電話番号tr · 名前 · IBAN_TR · 会社名_tr · MERSIS · 郵便番号_tr · 州_tr
「で」
Steueridentificationsnummer · Sozialversicherungsnummer
「フランス」
SIREN · SIRET · INSEE/NIR
「それ」
Codice Fiscale · Partita IVA
「nl」
BSN・KvK
「エス」
DNI/NIE・CIF
「イギリス」
NI番号・UTR
「私たち」
SSN・EIN・ITIN
「えー」
E.164 電話 · IBAN (EU+GB+CH+NO) · 会社名
ユニバーサル検出器 (ロケールに関係なく常にアクティブ): email · iban ·credit_card ·ip ·ip_v6
言語検出: flexorch-audit は依存関係がありません。言語検出ライブラリは含まれていません。
正しいロケールを自分で渡すか、「und」(デフォルト) を使用してすべての検出器をアクティブにします。
タイプ
説明
電子メール
RFC-5321電子メールアドレス
イバン
ISO 13616 IBAN — mod-97 validated; iban_tr または iban_intl が同じスパンで起動される場合は抑制されます
クレジットカード
16-digit groups, Luhn-validated
ip
IPv4アドレス
ip_v6
IPv6 — full, compressed :: , loopback forms
トルコ ( locale="tr" )
タイプ
説明
国民ID_tr
TCKN — 11-digit, modular arithmetic checksum
税ID_tr
VKN — 10-digit, Luhn-variant checksum
電話番号
Turkish mobile: +90 / 0 prefix + 10 digits
名前
ラベル接頭辞付きの名前: Adı: 、フルネーム: 、顧客名: 、など

c.
iban_tr
トルコの IBAN ( TR + 24 文字)、mod-97 検証済み
company_name_tr
TR 法的接尾辞を持つ会社: A.Ş. · Ltd.Şti. · Koll.Şti. · Koop. · T.A.Ş.
メルシス_いいえ
MERSIS — 16 桁の会社登録番号
postal_code_tr
トルコの郵便番号 (県番号 01–81)
province_tr
All 81 Turkish provinces
Germany ( locale="de" )
タイプ
説明
tax_id_de
Steueridentifikationsnummer — 11 桁、ISO 7064 MOD 11,2 チェックサム
social_id_de
Sozialversicherungsnummer — 地域 + 生年月日 + 文字 + シリアル
France ( locale="fr" )
タイプ
説明
siret_fr
SIRET — 14 桁、ラベルプレフィックスゲート付き
会社ID_fr
SIREN — 9 桁、ラベルプレフィックスゲート付き
social_id_fr
INSEE/NIR — 1 または 2 で始まる 15 桁
イタリア ( locale="it" )
タイプ
説明
National_id_it
Codice Fiscale — 16 文字の英数字、大文字の正規化
税金_id_it
Partita IVA — 11 桁、Agenzia delle Entrate チェックサム
オランダ ( locale="nl" )
タイプ
説明
国民ID_nl
BSN — 9 桁、11 チェック (加重和 mod 11)
company_id_nl
KvK — 8 桁、ラベルプレフィックスゲート
スペイン ( locale="es" )
タイプ
説明
国民ID
DNI (8 桁 + 文字、mod-23) および NIE (X/Y/Z プレフィックス、同じチェック)
税金ID
CIF — 文字プレフィックス + 7 桁 + 制御文字
イギリス ( locale="uk" )
タイプ
説明
social_id_uk
NI 番号 — 2 文字 + 6 桁 + A/B/C/D; HMRC 禁止プレフィックスは除外されます
Tax_id_uk
UTR — 10 桁、ラベルプレフィックスゲート付き
米国 ( locale="us" )
タイプ
説明
ssn
SSN — ###-##-#### 、無効なプレフィックス (000/666/9xx) は除外されます
Tax_id_us
EIN — XX-XXXXXXX 、IRS の無効なエリア プレフィックスは除外されます
National_id_us
ITIN — 9XX-7X/8X/9X-XXXX 中間グループが検証済み
EU / 国際 ( locale="eu" )
タイプ
説明
電話国際
E.164 国際電話 — 7 ～ 15 桁、TR (+90) を除く
iban_intl
IB

EU+GB+CH+NO の AN — ISO 13616 国+長さテーブル + mod-97
会社名_国際
国際的な接尾辞が付いている会社: GmbH · LLC · S.r.l.・B.V. ・SAS ・Inc. ・Ltd.等
ノイズ検出
noise_ratio は、空白行またはシンボル クラッタを含む行の割合を測定します。
result = 監査 ( "行をクリーン \n @@@ゴミ \n \n クリーン" )
結果。ノイズ比率 # 0.5 (4 行中 2 行にノイズがある)
行が空白 (ストリップの後) であるか、 @ # ! から 3 つ以上の連続文字が含まれている場合、その行は「ノイズが多い」となります。 ~ * = 。
clean = Mask ( text , result [ "pii" ], Strategy = "redact" ) # デフォルト
clean = マスク (テキスト、結果 [ "pii" ]、戦略 = "トークン" )
clean = マスク (テキスト、結果 [ "pii" ]、戦略 = "ハッシュ" )
clean = マスク (テキスト、結果 [ "pii" ]、戦略 = "replace" )
戦略
出力例
編集 (デフォルト)
[編集済み_メール]
交換する
user@example.com (静的合成)
トークン
<PII_EMAIL_1> (コールごとにタイプごとに一意)
ハッシュ
[3d4f9a1b2c8e7f0a] (SHA-256 の最初の 16 進数文字)
品質グレード
quality_grade (A ～ D) とquality_score (0.0 ～ 1.0) は複合信号です。
スコアの計算式: 完全性 × (0.4 × ノイズスコア + 0.4 × 長さスコア + 0.2)
length_score = min(char_count / 500, 1.0) · ノイズスコア = max(0, 1 − Garage_ratio × 10)
自動言語検出はありません。flexorch-audit には依存関係がありません。ロケールを明示的に渡すか、デフォルトの「und」を使用してすべての検出器をアクティブにします。 LIMITATIONS.md を参照してください。
自立型の名前検出 (ラベル接頭辞なし) には NLP/NER が必要ですが、含まれていません。
置換マスキングは静的な合成値を使用します。ロケールを意識したリアルな合成は実装されていません。
ライブラリはプレーン テキストを監査します。 PDF/DOCX の解析、電子請求書の抽出、パイプライン オーケストレーションは対象外です。
flexorch-audit は、プリロード フィルターとして LangChain または LlamaIndex パイプラインにスロットを挿入します。品質の監査、PII の検出、

必要に応じて、ドキュメントが LLM に到達する前にマスクします。
LangChain — example/langchain_loader.py
例から。 langchain_loader import AuditedLoader # プロジェクトにコピーします
ローダー = AuditedLoader (
テキスト = my_texts 、
locale = "tr" 、 # または "de"、 "fr"、 "us"、 "und" (すべて)
Mask_pii = True 、 # ロードする前に PII を編集します
min_grade = "B" , # 低品質のドキュメントをスキップする
）
docs = ローダー 。ロード()
# doc.metadata → {"quality_grade": "A", "quality_score": 0.91, "pii_summary": [...], ...}
LlamaIndex — サンプル/llamaindex_reader.py
例から。 llamaindex_reader import AuditedReader # プロジェクトにコピー
Reader = AuditedReader ( locale = "tr" 、mask_pii = True )
ドキュメント = リーダー 。 load_data ( my_texts 、 min_grade = "B" )
# doc.extra_info → {"quality_grade": "A", "quality_score": 0.91, "pii_summary": [...], ...}
どちらのローダーも薄いラッパー (約 60 行) であり、 langchain-core または llama-index-core を超える新しい依存関係はありません。それらをプロジェクトにコピーします。フレームワークのロックインはありません。
JavaScript / TypeScript でも利用可能
npm install @flexorch/audit
貢献する
LLM データセットの依存性のない PII + 品質 + ノイズ監査 (TR/EU/US)
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。嘆願

[切り捨てられた]

## Original Extract

Zero-dependency PII + quality + noise audit for LLM datasets (TR/EU/US) - flexorch/flexorch-audit

GitHub - flexorch/flexorch-audit: Zero-dependency PII + quality + noise audit for LLM datasets (TR/EU/US) · GitHub
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
flexorch
/
flexorch-audit
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
20 Commits 20 Commits .github/ workflows .github/ workflows assets assets demo demo examples examples src/ flexorch_audit src/ flexorch_audit tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE LIMITATIONS.md LIMITATIONS.md README.md README.md pyproject.toml pyproject.toml View all files Repository files navigation
Zero-dependency PII detection, quality grading, and noise audit for LLM datasets — in a single function call.
Before feeding documents into an LLM pipeline you need to answer three questions:
Does this text contain personal data? Sending PII to a language model is a compliance risk.
Is the text quality high enough? Short, noisy, or duplicate records hurt fine-tuning and RAG retrieval.
How bad is the noise? Garbled encodings and symbol clutter degrade model output silently.
Most tools that answer these questions require heavy NLP frameworks, model weights, or cloud APIs. flexorch-audit answers all three with one call — using only regex and Python's standard library. No model weights, no network calls, no external packages.
Quality grade — A/B/C/D composite score: is this text LLM-ready at a glance?
Noise ratio — line-level symbol clutter detection ( noise_ratio ); values above 0.20 indicate likely extraction artifacts
PII detection — 30+ types across 8 countries (TR/DE/FR/IT/NL/ES/UK/US) + universal types; all regex-based with checksum validation
Batch audit — audit_batch() aggregates duplicate ratio and PII counts across an entire dataset in one call
Masking — four strategies: redact, replace (synthetic), token, hash
Zero runtime dependencies — pure Python stdlib, Python 3.10+
pip install flexorch-audit
Quick start
from flexorch_audit import audit , mask
text = open ( "contract.txt" ). read () # extract from PDF/DOCX first
result = audit ( text ) # "und" by default — all detectors active
# result = audit(text, locale="tr") # restrict to TR-only detectors
result . quality_grade # "B"
result . quality_score # 0.73 (0.0–1.0 composite)
result . noise_ratio # 0.04 (fraction of blank/garbage lines; >0.20 = low quality)
result . detected_language # "und" (locale you passed in; caller controls language)
result . pii_summary # [{"type": "email", "count": 2}, {"type": "national_id_tr", "count": 1}]
# Full findings and raw metrics — dict access also works:
result [ "pii" ] # [{"type": "email", "value": "ali@example.com", "start": 8, "end": 23}]
result [ "quality" ] # {"completeness": 1.0, "avg_length": 342, "duplicate_ratio": None}
result [ "noise" ] # {"garbage_ratio": 0.0, "encoding_ok": True}
clean = mask ( text , result [ "pii" ], strategy = "redact" )
# "Contact: [REDACTED_EMAIL]"
from flexorch_audit import redact_for_llm
clean = redact_for_llm ( "TCKN: 12345678950, email: ali@example.com" , locale = "tr" )
# "TCKN: [REDACTED_NATIONAL_ID_TR], email: [REDACTED_EMAIL]"
# Different masking strategies
redact_for_llm ( text , locale = "tr" , strategy = "token" ) # <PII_NATIONAL_ID_TR_1>
redact_for_llm ( text , locale = "tr" , strategy = "hash" ) # [3d4f9a1b2c8e7f0a]
redact_for_llm ( text , locale = "tr" , strategy = "replace" ) # static synthetic value
No PII found → original text returned unchanged.
from flexorch_audit import estimate_tokens
estimate_tokens ( "The quick brown fox jumps over the lazy dog." ) # → 16
estimate_tokens ( "" ) # → 0
Heuristic: words × 4/3 — no tiktoken required. Accuracy within ~15% of the real tokenizer for English and most European languages; treat as a planning estimate for context window sizing and cost forecasting.
from flexorch_audit import audit_batch
texts = [ record [ "text" ] for record in dataset ]
batch = audit_batch ( texts ) # locale="und" by default
batch [ "duplicate_ratio" ] # 0.12 — fraction of exact-duplicate records
batch [ "avg_quality_score" ] # 0.78
batch [ "pii_summary" ] # [{"type": "email", "count": 47}, ...]
batch [ "results" ] # list of AuditResult, one per text
Country coverage
locale
Detectors activated
"und" (default)
All locales combined — use when document language is unknown
"all"
Alias for "und"
"tr"
TCKN · VKN · phone_tr · name · IBAN_TR · company_name_tr · MERSIS · postal_code_tr · province_tr
"de"
Steueridentifikationsnummer · Sozialversicherungsnummer
"fr"
SIREN · SIRET · INSEE/NIR
"it"
Codice Fiscale · Partita IVA
"nl"
BSN · KvK
"es"
DNI/NIE · CIF
"uk"
NI number · UTR
"us"
SSN · EIN · ITIN
"eu"
E.164 phone · IBAN (EU+GB+CH+NO) · company name
Universal detectors (always active regardless of locale): email · iban · credit_card · ip · ip_v6
Language detection: flexorch-audit is zero-dependency — no language detection library is included.
Pass the correct locale yourself, or use "und" (default) to activate all detectors.
Type
Description
email
RFC-5321 email address
iban
ISO 13616 IBAN — mod-97 validated; suppressed when iban_tr or iban_intl fires on same span
credit_card
16-digit groups, Luhn-validated
ip
IPv4 address
ip_v6
IPv6 — full, compressed :: , loopback forms
Turkey ( locale="tr" )
Type
Description
national_id_tr
TCKN — 11-digit, modular arithmetic checksum
tax_id_tr
VKN — 10-digit, Luhn-variant checksum
phone_tr
Turkish mobile: +90 / 0 prefix + 10 digits
name
Label-prefixed name: Adı: , Full Name: , Customer Name: , etc.
iban_tr
Turkish IBAN ( TR + 24 chars), mod-97 validated
company_name_tr
Company with TR legal suffix: A.Ş. · Ltd.Şti. · Koll.Şti. · Koop. · T.A.Ş.
mersis_no
MERSIS — 16-digit company registry number
postal_code_tr
Turkish postal code (province plate 01–81)
province_tr
All 81 Turkish provinces
Germany ( locale="de" )
Type
Description
tax_id_de
Steueridentifikationsnummer — 11 digits, ISO 7064 MOD 11,2 checksum
social_id_de
Sozialversicherungsnummer — area + DOB + letter + serial
France ( locale="fr" )
Type
Description
siret_fr
SIRET — 14 digits, label-prefix gated
company_id_fr
SIREN — 9 digits, label-prefix gated
social_id_fr
INSEE/NIR — 15 digits, starts with 1 or 2
Italy ( locale="it" )
Type
Description
national_id_it
Codice Fiscale — 16 chars alphanumeric, uppercase normalized
tax_id_it
Partita IVA — 11 digits, Agenzia delle Entrate checksum
Netherlands ( locale="nl" )
Type
Description
national_id_nl
BSN — 9 digits, 11-check (weighted sum mod 11)
company_id_nl
KvK — 8 digits, label-prefix gated
Spain ( locale="es" )
Type
Description
national_id_es
DNI (8 digits + letter, mod-23) and NIE (X/Y/Z prefix, same check)
tax_id_es
CIF — letter prefix + 7 digits + control character
United Kingdom ( locale="uk" )
Type
Description
social_id_uk
NI number — 2 letters + 6 digits + A/B/C/D; HMRC forbidden prefixes excluded
tax_id_uk
UTR — 10 digits, label-prefix gated
United States ( locale="us" )
Type
Description
ssn
SSN — ###-##-#### , invalid prefixes (000/666/9xx) excluded
tax_id_us
EIN — XX-XXXXXXX , IRS invalid area prefixes excluded
national_id_us
ITIN — 9XX-7X/8X/9X-XXXX middle group validated
EU / International ( locale="eu" )
Type
Description
phone_intl
E.164 international phone — 7–15 digits, TR (+90) excluded
iban_intl
IBAN for EU+GB+CH+NO — ISO 13616 country+length table + mod-97
company_name_intl
Company with international suffix: GmbH · LLC · S.r.l. · B.V. · SAS · Inc. · Ltd. etc.
Noise detection
noise_ratio measures the fraction of lines that are blank or contain symbol clutter:
result = audit ( "clean line \n @@@garbage \n \n clean" )
result . noise_ratio # 0.5 (2 noisy lines out of 4)
A line is "noisy" when it is blank (after strip) or contains 3+ consecutive characters from @ # ! ~ * = .
clean = mask ( text , result [ "pii" ], strategy = "redact" ) # default
clean = mask ( text , result [ "pii" ], strategy = "token" )
clean = mask ( text , result [ "pii" ], strategy = "hash" )
clean = mask ( text , result [ "pii" ], strategy = "replace" )
Strategy
Example output
redact (default)
[REDACTED_EMAIL]
replace
user@example.com (static synthetic)
token
<PII_EMAIL_1> (unique per type per call)
hash
[3d4f9a1b2c8e7f0a] (SHA-256 first 16 hex chars)
Quality grade
quality_grade (A–D) and quality_score (0.0–1.0) are composite signals:
Score formula: completeness × (0.4 × noise_score + 0.4 × length_score + 0.2)
length_score = min(char_count / 500, 1.0) · noise_score = max(0, 1 − garbage_ratio × 10)
No automatic language detection — flexorch-audit has zero dependencies. Pass locale explicitly, or use the default "und" to activate all detectors. See LIMITATIONS.md .
Free-standing name detection (without a label prefix) requires NLP/NER — not included.
replace masking uses static synthetic values; locale-aware realistic synthesis is not implemented.
The library audits plain text. PDF/DOCX parsing, e-invoice extraction, and pipeline orchestration are out of scope.
flexorch-audit slots into any LangChain or LlamaIndex pipeline as a pre-load filter — audit quality, detect PII, and optionally mask before your documents reach the LLM.
LangChain — examples/langchain_loader.py
from examples . langchain_loader import AuditedLoader # copy to your project
loader = AuditedLoader (
texts = my_texts ,
locale = "tr" , # or "de", "fr", "us", "und" (all)
mask_pii = True , # redact PII before loading
min_grade = "B" , # skip low-quality documents
)
docs = loader . load ()
# doc.metadata → {"quality_grade": "A", "quality_score": 0.91, "pii_summary": [...], ...}
LlamaIndex — examples/llamaindex_reader.py
from examples . llamaindex_reader import AuditedReader # copy to your project
reader = AuditedReader ( locale = "tr" , mask_pii = True )
docs = reader . load_data ( my_texts , min_grade = "B" )
# doc.extra_info → {"quality_grade": "A", "quality_score": 0.91, "pii_summary": [...], ...}
Both loaders are thin wrappers (~60 lines) with no new dependencies beyond langchain-core or llama-index-core . Copy them into your project — no framework lock-in.
Also available for JavaScript / TypeScript
npm install @flexorch/audit
Contributing
Zero-dependency PII + quality + noise audit for LLM datasets (TR/EU/US)
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Plea

[truncated]
