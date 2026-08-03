---
source: "https://github.com/Angeliasrl/norms-mcp-engine"
hn_url: "https://news.ycombinator.com/item?id=49154830"
title: "Show HN: A fail-closed evidence engine for norms used by AI agents"
article_title: "GitHub - Angeliasrl/norms-mcp-engine: Fail-closed evidentiary eligibility engine for AI-agent norms, with a self-linting claim map. · GitHub"
author: "babelprotocol"
captured_at: "2026-08-03T12:50:01Z"
capture_tool: "hn-digest"
hn_id: 49154830
score: 1
comments: 0
posted_at: "2026-08-03T12:26:21Z"
tags:
  - hacker-news
  - translated
---

# Show HN: A fail-closed evidence engine for norms used by AI agents

- HN: [49154830](https://news.ycombinator.com/item?id=49154830)
- Source: [github.com](https://github.com/Angeliasrl/norms-mcp-engine)
- Score: 1
- Comments: 0
- Posted: 2026-08-03T12:26:21Z

## Translation

タイトル: Show HN: AI エージェントが使用する規範のフェールクローズ型証拠エンジン
記事のタイトル: GitHub - Angeliasrl/norms-mcp-engine: 自己リンティング要求マップを備えた、AI エージェント規範のフェイルクローズされた証拠適格性エンジン。 · GitHub
説明: 自己リンティング要求マップを備えた、AI エージェント基準のフェールクローズ証拠適格性エンジン。 - Angeliasrl/norms-mcp-engine

記事本文:
GitHub - Angeliasrl/norms-mcp-engine: 自己リンティング要求マップを備えた、AI エージェント基準のフェールクローズ証拠適格性エンジン。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
アンジェリアスル
/
ノルム-mcp-エンジン
公共
通知
あなたはきっとsiでしょう

通知設定を変更するためにログインしました
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
7 コミット 7 コミット .github/ workflows .github/ workflows scripts scripts src src test test .gitattributes .gitattributes .gitignore .gitignore CLAIM_MAP.md CLAIM_MAP.md LICENSE LICENSE README.md README.mdclaims.mjsclaims.mjs package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
実験的なプレリリース。エンジンのみ。 MCP サーバーではありません。 I/O、ネットワークなし、
持続性または転送。
バージョン 0.1.0 · 請求ごとの証拠: CLAIM_MAP.md ·
独立して検証されていません。
純粋な関数としての許容性モデルとクレーム マップ リンター。 I/O なし、なし
ネットワーク、トランスポートなし。
これは MCP サーバーではありません。これは、MCP サーバーが使用するエンジンです。
ここにはエントリポイント、ツール ハンドラー、デプロイメントはありません。ノルム-mcp 、
サーバーは、公開されたアーティファクトとしてまだ存在していません。ネーミングの違い
意図的です: 存在しないトランスポートを要求することは、まさにそのようなものになります。
このライブラリは可視化するために存在するのでオーバークレームします。
npm テスト # 73 テスト + クレーム マップ同期チェック;ネットワークもアカウントも必要ありません
何をするのか
規範的コーパス内の定数ごとに、その証拠ステータスと
適格性の判定 — 消費者は、次のようなものに依存することを拒否できます。
今のところ何にも良くありません。このパッケージは計算します。何も公開しないのですが、
交通機関がないからです。
何も強制しません。適格性を計算してレポートします。決定
尋ねている人のものです。
この財産を表すイタリア語は「opponibile」です。つまり、主張できるものです。
誰かに対して、第三者の前で。 「強制可能」というのは間違った翻訳です —
執行とは、誰かがそれを適用することを意味します。最も近いのは許容性です:
請求は次の場合に成立します

挑戦した。
起源と検証は異なる質問に答え、定数は止まらない
批准されたときに宣言される。
SOURCE_DOCUMENT | 原点タイプOWNER_DECLARATION
検証状態 承認済み |未確認
未確認
承認済み
SOURCE_DOCUMENT
転写されており、決してチェックされていない
ソースと照らし合わせてチェックする
OWNER_DECLARATION
主張し、批准法を待っている
主張され、その後正式に承認された
右下のセルは、単軸スキームでは名前を付けることができないセルです。それは重要です
簿記を超えて：それは、後の読者が、簿記から派生した規範を区別できるようにするものです。
所有者がそう言ったから成立する規範からの文書。
バイナリ フィールドでは、checked および false と Never selected を区別できません。
通貨 現在 |古い |不明
権限ステータス 有効 |無効 |不明
有効期限 | 有効期限期限切れ |レビュー期限 |不明
UNKNOWN はクローズに失敗します - 否定的な結果が得られます - が、
ブロッキング。誰も調べなかったものは、不足していることが判明したものではありません。
'norms-mcp-engine/model' から {qualifiedAsGround } をインポートします。
適格なAsGround (エントリ) ;
// { 対象となる: false,
// ブロック: [], // 必要なものは何も見つかりませんでした
// 不明: ['authority_status'], // 誰もチェックしていません
// 未検査 : true }
資格と適用性
適格_as_ground = 検証状態 == 承認済み
AND 通貨 == 現在
AND 権限ステータス == 有効
かつ expiry_status == アクティブ
その答えは、まったくの根拠として使用される可能性があります。答えられないですよね
目の前の決定に耐える — 定数は批准できるし、最新のものでもよいし、
適切に承認され、有効期限が切れていないにもかかわらず、まだ何か他のものについてです。
'norms-mcp-engine/model' から { admissibleFor } をインポートします。
admissibleFor ( エントリ , { 件名 : 'android' } ) ;
エントリがスコープを宣言していない場合、結果は許容されます: false と
スコープ_既知: false

そして、適用するかどうかは呼び出し側の判断であることに注意してください。の
エンジンは一致を報告します。それは決して決定しません。
制限 (v0.1): スコープの一致は subject 間でのみ正確な値になります。
管轄区域と該当する業務。階層、ワイルドカード、または
否定。空のスコープまたは不正なスコープがスローされ、決して一致しません。
なぜなら、「指定されていない」という言葉は、暗黙のうちに「すべてに当てはまる」という意味であってはいけないからです。
批准にはその証拠がある
{
検証状態 : '承認済み' 、
批准: {
日付: '2026-07-31' 、
ドキュメント: 'POLICY.md' 、
sha256 : '<64 小文字の 16 進数>' ,
Section_id : 'sec-4-2' , // 安定した ID、解決済み
Section_label : '§4.2 Retention' // 人間が読める形式ですが、ずれの可能性があります
}
}
完全なブロックなしで RATIFIED をスローします。これは立ち止まる警備員です
「会話で確認」されて批准されることはありません。
改ざんではなく再検証
保存された sha256 は、検証が実行されたドキュメントのダイジェストです。
インデックスが移動しても批准が移動しない場合、証明は古くなります。
定数が false であるのとは異なります。
import { revalidate } から 'norms-mcp-engine/model' ;
再検証 (エントリ、コーパスインデックス) ;
// { 通貨: 'STALE'、理由: 'フィンガープリントがインデックスから逸脱しました。証明には再検証が必要です' }
インデックスにないドキュメントは STALE ではなく UNKNOWN を返します。
永続的であるべきルールに人為的な減衰条件を強制することは、
それ自体のエラー。
有効期限ポリシー: 条件付き |レビュー済み |永久的
PERMANENT は期限切れになりませんが、permanence.authority を記録しない限りスローします。
そして永続性と理由。永続性とは決定であり、決定が存在しないことではありません。の
check は validateEntry 内に存在するため、決して呼び出さないことでバイパスすることはできません
評価期限 。
条件が宣言されていない CONDITIONAL は、 ACTIVE ではなく UNKNOWN を生成します。
条件が存在しないことは、条件が存在しないという証拠にはなりません

解雇された。
各ドキュメントの ID バイト順:
レン(ID) || ID || len(コンテンツ) ||内容
2 つのプロパティ。どちらも実際のシステムに影響を及ぼします。
長さは UTF-8 バイト数です。 String.prototype.length は UTF-16 コードを返します
単位。 BMP の外にある 1 つの文字により、2 つの適合する実装が作成されます。
同意しません。プレフィックスは 8 バイトの符号なしビッグエンディアンです - 固定幅とエンディアン
は仕様の一部です。
連結ではなくフレーム化。 2 つの異なるコーパスを同じものに連結できる
バイトストリーム。このスイートは衝突を示し、それを除去するフレームを示しています。
OK 衝突: 単純な連結があいまいです
ドキュメントごとの正規形式: UTF-8、BOM なし、LF 行末、末尾の空白
ストリップされた、Unicode NFC。
散文で書かれたクレーム マップは、適用するために覚えておく必要がある慣例です。
これはテストになります。
'norms-mcp-engine/claimmap' から { lintClaimMap } をインポートします。
lintClaimMap (クレーム) ;
// { ok: false、結果: [{ id: '2'、ルール: 'R4'、メッセージ: '…' } ]、カウント: {…} }
このパッケージ自体のマップはデータであり、それ自体がリントします
claims.mjs は正規形式です。そこから CLAIM_MAP.md が生成されます。
scripts/build-claimmap.mjs 、コミットされたマークダウンが次の場合、npm テストは失敗します。
データからずれているか、マップが R1 ～ R10 を通過していない場合。
反転は意図的です。散文として残された地図は、誰かがそうしなければならない慣習である
忘れずに応募してください、そしてこのプロジェクト自体の歴史の中で、その慣例は失敗しました
3 つのリビジョンが実行中です。
状態: O (観察済み、第三者によって解決可能) · O-PENDING (報告済み、未解決)
まだ解決可能 — 証拠として認められない ) · D (派生) · A (オープン、
という名前の前提条件)。
MCP サーバーがありません。トランスポートもツール ハンドラーもエントリポイントもありません。
持続性はありません。純粋な関数。呼び出し元がストレージを所有しています。
applyRevalidation / applyExpiry ヘルパーはありません。再検証し、
評価期限の戻り値 v

命令;呼び出し側がそれらを適用します。
JSON スキーマはありません。シェイプは validateEntry およびスイートによって強制されます。
コードの独立した監査はありません。 7回のレビューラウンド。最後に読んだ2冊
ソース。人間の査読者は一人もいなかった。
独立した検証はありません。 1 人の作成者、1 回のテスト実行、外部なし
レプリケーション。
ここで行われたすべての申し立ての完全な証拠状況については、CLAIM_MAP.md を参照してください。
すべてのパブリック API は入力を検証し、ModelError または CorpusError をスローします。
安定したコードで。不正な形式の記録は決して評決として解釈されません。
否定的でも肯定的でもない — 不正な形式や不適格なものは、
異なる事実と、証拠の地位に関する図書館がそれらを混同してはなりません。
NormaliseDocuments は、コーパス、つまりダイジェストと
インデックスには同一のルールが適用されるため、どのドキュメントについて意見が相違することはありません
存在します。 buildIndex は null プロトタイプ オブジェクトを使用し、 revalidate は次のように検索します。
Object.hasOwn なので、 __proto__ などの ID は実際のレコードまたは
何もない。
Apache ライセンス 2.0 — Copyright 2026 Francesco Riva。詳細についてはライセンスを参照してください
テキスト。コミット 83438ef までのバージョンは MIT の下で公開されました。これとそれ以降
リビジョンは Apache-2.0 の下にあります。
CLAIM_MAP.md は同じライセンスの下にあります。ライセンスではなく実際的な要求
用語: いつ、何が請求されたかの記録であるため、修正されたコピーが必要です。
同じ名前で流通しない - その価値は変更されていないことにあります
事後。
自己リンティング請求マップを備えた、AI エージェント基準のフェールクローズ証拠適格性エンジン。
Readme Apache-2.0 ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Fail-closed evidentiary eligibility engine for AI-agent norms, with a self-linting claim map. - Angeliasrl/norms-mcp-engine

GitHub - Angeliasrl/norms-mcp-engine: Fail-closed evidentiary eligibility engine for AI-agent norms, with a self-linting claim map. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
Angeliasrl
/
norms-mcp-engine
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
7 Commits 7 Commits .github/ workflows .github/ workflows scripts scripts src src test test .gitattributes .gitattributes .gitignore .gitignore CLAIM_MAP.md CLAIM_MAP.md LICENSE LICENSE README.md README.md claims.mjs claims.mjs package.json package.json View all files Repository files navigation
Experimental prerelease. Engine only. Not an MCP server. No I/O, network,
persistence or transport.
Version 0.1.0 · claim-by-claim evidence: CLAIM_MAP.md ·
not independently verified.
The admissibility model and claim-map linter, as pure functions. No I/O, no
network, no transport.
This is not an MCP server. It is the engine an MCP server would consume.
There is no entrypoint, no tool handler and no deployment here. norms-mcp ,
the server, does not yet exist as a published artifact. The naming distinction
is deliberate: claiming a transport that is absent would be exactly the kind of
overclaim this library exists to make visible.
npm test # 73 tests + claim-map sync check; no network, no account required
What it does
Computes, for each constant in a normative corpus, its evidentiary status and
eligibility verdict — so a consumer can decline to rely on something that is
not currently good for anything. This package computes; it publishes nothing,
because it has no transport.
It enforces nothing. It computes eligibility and reports it. The decision
belongs to whoever is asking.
The Italian word for the property is opponibile : something you can assert
against someone, before a third party. Enforceable is the wrong translation —
enforcement means somebody applies it. The closest is admissibility : whether
a claim holds up when challenged.
Origin and verification answer different questions, and a constant does not stop
being declared when it becomes ratified.
origin.type SOURCE_DOCUMENT | OWNER_DECLARATION
verification_state RATIFIED | UNCONFIRMED
UNCONFIRMED
RATIFIED
SOURCE_DOCUMENT
transcribed, never checked
checked against its source
OWNER_DECLARATION
asserted, awaiting a ratification act
asserted, then formally approved
The bottom-right cell is the one a single-axis scheme cannot name. It matters
beyond bookkeeping: it is what lets a later reader distinguish norms derived from
documents from norms that stand because the owner said so.
A binary field cannot distinguish checked and false from never checked .
currency CURRENT | STALE | UNKNOWN
authority_status VALID | INVALID | UNKNOWN
expiry_status ACTIVE | EXPIRED | REVIEW_DUE | UNKNOWN
UNKNOWN fails closed — it yields a negative result — but never appears in
blocking . A thing nobody has examined is not a thing found wanting.
import { eligibleAsGround } from 'norms-mcp-engine/model' ;
eligibleAsGround ( entry ) ;
// { eligible: false,
// blocking: [], // nothing was found wanting
// unknown: ['authority_status'], // nobody has checked
// unexamined : true }
Eligibility, and applicability
eligible_as_ground = verification_state == RATIFIED
AND currency == CURRENT
AND authority_status == VALID
AND expiry_status == ACTIVE
That answers may this be used as a ground at all . It does not answer does it
bear on the decision in front of me — a constant can be ratified, current,
competently approved and unexpired, and still be about something else.
import { admissibleFor } from 'norms-mcp-engine/model' ;
admissibleFor ( entry , { subject : 'android' } ) ;
Where the entry declares no scope , the result is admissible: false with
scope_known: false and a note that applicability is the caller's judgement. The
engine reports the match; it never decides.
Limit (v0.1): scope matching is exact-value only across subject ,
jurisdiction and applicable_operations . No hierarchies, wildcards or
negation. An empty or malformed scope throws — it never becomes a match,
because "unspecified" must not silently mean "applies to everything".
Ratification carries its proof
{
verification_state : 'RATIFIED' ,
ratification : {
date : '2026-07-31' ,
document : 'POLICY.md' ,
sha256 : '<64 lowercase hex>' ,
section_id : 'sec-4-2' , // stable id, resolved against
section_label : '§4.2 Retention' // human-readable, may drift
}
}
RATIFIED without a complete block throws . This is the guard that stops
"verified in conversation" from ever becoming a ratification.
Revalidation, not falsification
The stored sha256 is the digest of the document the verification ran against.
If the index moves and the ratification does not, the proof is stale — which is
not the same as the constant being false.
import { revalidate } from 'norms-mcp-engine/model' ;
revalidate ( entry , corpusIndex ) ;
// { currency: 'STALE', reason: 'fingerprint diverged from index; proof requires revalidation' }
A document absent from the index yields UNKNOWN , not STALE .
Forcing an artificial decay condition onto a rule that ought to be permanent is
its own error.
expiry_policy: CONDITIONAL | REVIEWED | PERMANENT
PERMANENT never expires but throws unless it records permanence.authority
and permanence.reason . Permanence is a decision, not an absence of one. The
check lives in validateEntry , so it cannot be bypassed by never calling
evaluateExpiry .
CONDITIONAL without declared conditions yields UNKNOWN , not ACTIVE : the
absence of conditions is not evidence that none fired.
for each document, in id-byte order:
len(id) || id || len(content) || content
Two properties, both of which have bitten real systems:
Lengths are UTF-8 byte counts. String.prototype.length returns UTF-16 code
units. One character outside the BMP makes two conforming implementations
disagree. The prefix is 8-byte unsigned big-endian — fixed width and endianness
are part of the specification.
Framing, not concatenation. Two different corpora can concatenate to the same
byte stream. The suite demonstrates the collision and shows framing removing it:
ok COLLISION: plain concatenation is ambiguous
Per-document canonical form: UTF-8, no BOM, LF line endings, trailing whitespace
stripped, Unicode NFC.
A claim map written as prose is a convention someone has to remember to apply.
This makes it a test.
import { lintClaimMap } from 'norms-mcp-engine/claimmap' ;
lintClaimMap ( claims ) ;
// { ok: false, findings: [{ id: '2', rule: 'R4', message: '…' } ], counts: {…} }
This package's own map is data, and lints itself
claims.mjs is the canonical form. CLAIM_MAP.md is generated from it by
scripts/build-claimmap.mjs , and npm test fails if the committed Markdown has
drifted from the data or if the map does not pass R1–R10.
The inversion is deliberate. A map kept as prose is a convention someone has to
remember to apply, and in this project's own history that convention failed
three revisions running.
States: O (observed, resolvable by a third party) · O-PENDING (reported, not
yet resolvable — not admissible as evidence ) · D (derived) · A (open,
precondition named).
No MCP server. No transport, no tool handlers, no entrypoint.
No persistence. Pure functions; the caller owns storage.
No applyRevalidation / applyExpiry helpers. revalidate and
evaluateExpiry return verdicts; the caller applies them.
No JSON schema. Shapes are enforced by validateEntry and the suite.
No independent audit of the code. Seven review rounds; the last two read
the source. None was a human reviewer.
No independent verification. One author, one test run, no external
replication.
See CLAIM_MAP.md for the full evidentiary status of every claim made here.
Every public API validates its input and throws a ModelError or CorpusError
with a stable code . A malformed record is never interpreted as a verdict —
neither negative nor positive — because malformed and ineligible are
different facts, and a library about evidentiary status must not conflate them.
normaliseDocuments is the single admission gate for a corpus: the digest and
the index apply identical rules, so they cannot disagree about which documents
exist. buildIndex uses a null-prototype object and revalidate looks up with
Object.hasOwn , so an id such as __proto__ resolves to a real record or to
nothing.
Apache License 2.0 — Copyright 2026 Francesco Riva. See LICENSE for the full
text. Versions up to commit 83438ef were published under MIT; this and later
revisions are under Apache-2.0.
CLAIM_MAP.md is under the same licence. A practical request, not a licence
term: it is a record of what was claimed and when, so a modified copy should
not circulate under the same name — its value lies in not having been altered
after the fact.
Fail-closed evidentiary eligibility engine for AI-agent norms, with a self-linting claim map.
Readme Apache-2.0 license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
