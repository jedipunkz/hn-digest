---
source: "https://www.korvo.xyz/arf"
hn_url: "https://news.ycombinator.com/item?id=49204008"
title: "Show HN: ARF – a record format for AI evaluation runs, with reproducible digests"
article_title: "ARF - a record format for evaluation runs | Korvo"
author: "akshay_bhardwaj"
captured_at: "2026-08-06T23:49:57Z"
capture_tool: "hn-digest"
hn_id: 49204008
score: 2
comments: 0
posted_at: "2026-08-06T23:25:29Z"
tags:
  - hacker-news
  - translated
---

# Show HN: ARF – a record format for AI evaluation runs, with reproducible digests

- HN: [49204008](https://news.ycombinator.com/item?id=49204008)
- Source: [www.korvo.xyz](https://www.korvo.xyz/arf)
- Score: 2
- Comments: 0
- Posted: 2026-08-06T23:25:29Z

## Translation

タイトル: Show HN: ARF – 再現可能なダイジェストを備えた AI 評価実行用のレコード形式
記事のタイトル: ARF - 評価実行のレコード形式 |コルボ
説明: Atlas Record Format (ARF) v0.1: 10 個の JSON レコード タイプと正規化ルール (RFC 8785 + SHA-256) により、2 つの実装が同じレコードの同じダイジェストに到達します。

記事本文:
ARF - 評価実行用のレコード形式 | Korvo Korvo のユースケース 統合 Medha Relay Atlas ARF アトラス レコード フォーマットのダウンロード · v0.1
ARF - 評価実行用のレコード形式
したがって、スコアは、それを生み出した質問、モデル、入力、判断にまで遡ることができます。レコードは同一のバイトに正規化されるため、2 つの実装は同じ SHA-256 ダイジェストに到達します。
Korvo Atlas はリファレンス実装です。
v0.1 は不安定であり、v1.0 より前の移行パスがないと変更される可能性があります。公開された形状は暫定的なものとして扱います。
形状に同意できない場合、またはそれを表現できない場合がありますか?問題をオープンします - v0.1 はそれを破るのに最適な時期です。
ほとんどの評価パイプラインはスコアの CSV で終了します。数週間後、行には 0.82 と表示されますが、どのプロンプトのバージョン、どのモデルのビルド、またはどの判断ステップでそれが生成されたのかは誰も言えません。
その代償は、数値を守らなければならない人にかかっています。エンジニアは回帰が生じた理由を尋ね、レビュー担当者はスコアが何をカバーしているかを尋ね、監査人は入力内容を確認するよう求めました。スコアはそのまま残ります。それらを意味するようになった文脈は意味を持ちません。
10 個の JSON レコード タイプと 1 つのハッシュ ルール。 ARF は新しい形式ではありません。これは Korvo Atlas がすでに実装している形式であり、同じバイトとダイジェストを製品の外部でチェックできるように書き留められています。
検証は完全にオフラインです。ネットワーク呼び出しはなく、スキーマ $id 値が逆参照されることはありません。これらの識別子 (名前空間パス korvo.xyz/arf/ns/v0.1 ) の名前タイプ - これらはフェッチ ターゲットではないため、まだ何も提供されていません。
すべての ARF レコードはプレーンな JSON オブジェクトです。これは、claim.ndjson からコピーされた実際の Claim です。
{
"id" : "6251dbd4-a1f2-4b8b-9cd0-594a46a2f5c3" ,
"statement" : "言語モデルを独自の合成出力でトレーニングすると、出力の品質が徐々に低下し、「モデル崩壊」と呼ばれる現象が発生します。」 、
「コン

信頼度" : "高い" 、
"sourceIds" : [ "bafb465b-47be-45fd-b94f-ea6599153e4f" ],
"questionId" : "7ebb02e3-c3db-4010-a03c-13d7e7ca4df8" ,
"ステータス" : "公開済み" ,
"タグ" : [ "モデル崩壊" , "合成データ" ],
"起源" : "人間" 、
"reviewStatus" : "未レビュー" ,
"createdBy" : "アリス・チェン" ,
"createdAt" : "2026-03-25T20:17:56.727Z" ,
"更新日時" : "2026-03-25T20:17:56.727Z"
アサーション自体 - レコードごとに 1 つのクレーム。段落に複数のクレームが含まれることはありません。 sourceIds このクレームが基づいている入力。少なくとも 1 つは必須であり、順序は重要です。順序を変更するとダイジェストが変更されます。 questionId このクレームが回答する質問。主張は疑問を向けるものであり、その逆ではありません。信頼度 著者がどの程度自信を持っているか - 低、中、高、検証済みのいずれか。これは明示された順位であり、計算されたスコアではありません。エンジンの判定は、代わりに評価版に反映されます。レコードは .ndjson として 1 行に 1 つ保存されます。未設定はキーが存在しないことを意味します。ARF は「値なし」を意味する null を書き込むことはありません。
レコードが相互に参照する方法
すべてのリンクは、ポインティング レコードのフィールドに保持される UUID です。 X --field--> Y を「X.field は Y の ID を保持します」と読みます。
クレーム --questionId----> 質問
--sourceIds-----> ソースを要求する
アーティファクト --claimIds-----> クレーム
アーティファクト --sourceIds-----> ソース
評価 --claimIds-----> クレーム
チャレンジ --claimId-------> クレーム
チャレンジ --challengerId--> バリデータ
承認 --claimId------> クレーム
承認 --validatorId---> バリデーター
リビジョン --artifactId----> アーティファクト
ChainRecord --artifactId----> アーティファクト
ArtifactBundle { アーティファクト、質問、クレーム[]、ソース[]、PublishedBy、PublishedAt }
このオブジェクト全体がハッシュ化される値です。 ArtifactBundle は公開時のエンベロープです。1 つのアーティファクト、その質問、および使用されたクレームとソースが 1 つの JSON オブジェクトに収集されます。

。そのオブジェクトがハッシュ化されるものです。
役割ごとにグループ化されています。必須フィールドは、各スキーマで必須とマークされているフィールドとまったく同じです。
質問例/question.ndjson 評価中の質問。
必須: ID、テキスト、ステータス、createdAt
ソース例/source.ndjson 引用できる入力。
必須: ID、タイプ、タイトル、createdAt
Artifact Examples/artifact.ndjson クレームとソースから組み立てられた文書による出力。
必須: ID、タイトル、タイプ、本文、claimIds、sourceIds、createdAt
何が主張されたか、そしてエンジンが何を結論付けたか。
Claim example/claim.ndjson 1 つのアサーションと、それが基づいているソース。
必須: ID、ステートメント、信頼度、sourceIds、createdAt
評価スキーマ/evaluation.json クレームに関するエンジンの結果。エンジンのペイロードはそのまま保存されます。
必須: id、claimIds、エンジン、protocolVersion、ペイロード、createdAt
誰がそれをチェックし、何を言ったか。
Validator の例/validator.ndjson クレームをレビューする人間またはエージェント。
必須: ID、名前、タイプ、createdAt
異議申し立ての例/challenge.ndjson クレームに対する記録された異議申し立て。
必須: id、claimId、challengerId、reason、createdAt
裏書きの例/endorsement.ndjson 重要な主張を支持する記録。
必須: id、claimId、validatorId、weight、createdAt
アーティファクトがどのように変更され、何が公開されたか。
リビジョン スキーマ/revision.json アーティファクトの特定時点のスナップショット。
必須: ID、artifactId、バージョン、スナップショット、createdAt
ChainRecord schemas/chain-record.json バンドル ハッシュを含むパブリケーション レコード。
必須: id、artifactId、contentHash、chain、txHash、publishedAt
サンプルレコードがまだ公開されていないため、 Evaluation 、 Revision 、および ChainRecord はスキーマにリンクされています。意図的な評価の場合:claimIds リンクはエンジン ペイロードから回復できず、捏造された例は

最も誤解されやすい分野を誤って伝えます。
正規化とは、JSON 値を 1 つの正確なバイト シーケンスに減らすことを意味し、正しい実装であれば同じバイトが生成されます。 4 段階、1 つの操作。
{
"id" : "602d54e3-0650-40c8-898c-e6d3c0e71d71" ,
"claimId" : "6251dbd4-a1f2-4b8b-9cd0-594a46a2f5c3" ,
"validatorId" : "d2f602ae-9be6-43c0-8c1e-622980547d49" ,
"comment" : "シュマイロフらの論文からの強力な証拠。モデルの崩壊は十分に文書化されています。" 、
「重み」：5、
"作成時刻" : "2026-03-25T20:17:56.727Z"
} 2. 正規形式 (RFC 8785) 293 バイト · 末尾の改行なし Copy {"claimId":"6251dbd4-a1f2-4b8b-9cd0-594a46a2f5c3","comment":"Shumailov らの論文からの強力な証拠。モデルの崩壊はよく文書化されています。","createdAt":"2026-03-25T20:17:56.727Z","id":"602d54e3-0650-40c8- 898c-e6d3c0e71d71","validatorId":"d2f602ae-9be6-43c0-8c1e-622980547d49","weight":5} 3. それらのバイトの SHA-256 をコピー 56c223f1dac62072d1999e00cd3ba854352682294425fa10ad102f1a0c73ea6e 4. エンコードされたハッシュ コピーsha256:56c223f1dac62072d1999e00cd3ba854352682294425fa10ad102f1a0c73ea6e 自分で確認してください
ステップ 2 を任意の SHA-256 に貼り付けると、ステップ 3 が表示されます。
printf '%s' '{"claimId":"6251dbd4-a1f2-4b8b-9cd0-594a46a2f5c3","comment":"シュマイロフらの論文からの強力な証拠。モデルの崩壊はよく文書化されています。","createdAt":"2026-03-25T20:17:56.727Z","id":"602d54e3-0650-40c8-8 98c-e6d3c0e71d71","validatorId":"d2f602ae-9be6-43c0-8c1e-622980547d49","weight":5}' | shasum -a 256 バイトを決定する 3 つのルール
オブジェクト キーは、UTF-16 コード単位であらゆる深さで並べ替えられます。キーを並べ替えてもダイジェストは変更されません。
配列の順序は保持され、重要です。 sourceId の並べ替えは別のレコードであり、別のダイジェストを生成する必要があります。
ダイジェストは、再解析または p ではなく、正規のバイトをカバーします。

retty で出力された文字列。末尾の改行はありません。
JSON 表現のない値 ( NaN 、 Infinity 、 unknown ) は、強制されるのではなく拒否されます。表現できない値のハッシュを発行すると、誰も再現できない証拠が生成されます。
同じ操作を ArtifactBundle に適用すると、その contentHash が得られます。 Bundle.ndjson は sha256:81ea9018…db24d673 を記録し、この方法で再計算するとその値が再現されます。
アルゴリズムは rfc8785-jcs-v1 としてバージョン管理され、バンドル内で保持されます。認識できないバージョン文字列に一致する実装は、デフォルトにフォールバックするのではなく、それを拒否する必要があります。別のアルゴリズムで暗黙的にハッシュすると、確実に間違った答えが生成されます。
オフライン。ネットワークも $id 逆参照もありません。
git クローン https://github.com/korvohq/atlas.git
CDアトラス
npmci
npm run arf:validate -- spec/v0.1/examples/*.ndjson 出力コピー PASS spec/v0.1/examples/artifact.ndjson:1 (アーティファクト)
PASS spec/v0.1/examples/bundle.ndjson:1 (バンドル)
PASS spec/v0.1/examples/challenge.ndjson:1 (チャレンジ)
PASS spec/v0.1/examples/claim.ndjson:1 (クレーム)
...
PASS spec/v0.1/examples/validator.ndjson:2 (バリデーター)
16 件が合格、0 件が失敗 レコードは、サーバーが使用するのと同じ JSON スキーマに対してチェックされます。タイプはレコードの形状から推測されるか、 --type=claim で強制されます。
権威のあるものはすべて spec/ 、ライセンスされた Apache-2.0 の下に存在します。 Korvo Atlas サーバーは引き続き AGPL-3.0 のみです。
spec/v0.1/README.md 各部分がどのように組み合わされるか。
canonicalization.md ハッシュ ルール、適合ベクトル、移植チェックリスト。
Disclosure.md 著者が記録を公開するつもりかどうか。
Bundle.schema.json ArtifactBundle の JSON スキーマ。
例/ 上記のコマンドで使用されるシード レコード。
例としては、実稼働データではなく、開発フィクスチャから読み取り専用でエクスポートされたシード レコードがあります。何もない

spec/v0.1/examples/ はライブ デプロイメントからのものです。
実行の公開アーカイブ: korvo.xyz/atlas 。このページ: korvo.xyz/arf 。名前空間 korvo.xyz/arf/ns/v0.1 は型に名前を付けますが、提供されません。

## Original Extract

Atlas Record Format (ARF) v0.1: ten JSON record types and a canonicalization rule (RFC 8785 + SHA-256) so two implementations reach the same digest for the same record.

ARF - a record format for evaluation runs | Korvo Korvo Use Cases Integrations Medha Relay Atlas ARF Download Atlas Record Format · v0.1
ARF - a record format for evaluation runs
So a score can be traced back to the question, model, inputs and judgement that produced it. Records canonicalize to identical bytes, so two implementations reach the same SHA-256 digest.
Korvo Atlas is the reference implementation.
v0.1 is unstable and may change without a migration path before v1.0 . Treat published shapes as provisional.
Disagree with a shape, or have a case it cannot express? Open an issue - v0.1 is the right time to break it.
Most evaluation pipelines end in a CSV of scores. Weeks later a row says 0.82 , and nobody can say which prompt version, which model build, or which judging step produced it.
The cost lands on whoever has to defend the number: the engineer asked why a regression appeared, the reviewer asked what the score covered, the auditor asked to see the inputs. The scores survive. The context that made them mean anything does not.
Ten JSON record types and one hashing rule. ARF is not a new format - it is the format Korvo Atlas already implements, written down so the same bytes and digests can be checked outside the product.
Validation is fully offline: no network calls, and schema $id values are never dereferenced. Those identifiers (namespace path korvo.xyz/arf/ns/v0.1 ) name types - they are not fetch targets, and nothing serves them yet.
Every ARF record is a plain JSON object. This is a real Claim , copied from claim.ndjson .
{
"id" : "6251dbd4-a1f2-4b8b-9cd0-594a46a2f5c3" ,
"statement" : "Training language models on their own synthetic outputs leads to progressive degradation of output quality, a phenomenon termed \"model collapse\"." ,
"confidence" : "high" ,
"sourceIds" : [ "bafb465b-47be-45fd-b94f-ea6599153e4f" ],
"questionId" : "7ebb02e3-c3db-4010-a03c-13d7e7ca4df8" ,
"status" : "published" ,
"tags" : [ "model collapse" , "synthetic data" ],
"origin" : "human" ,
"reviewStatus" : "unreviewed" ,
"createdBy" : "Alice Chen" ,
"createdAt" : "2026-03-25T20:17:56.727Z" ,
"updatedAt" : "2026-03-25T20:17:56.727Z"
} statement The assertion itself - one claim per record, never a paragraph containing several. sourceIds The inputs this claim rests on. At least one is required, and order is significant: reordering it changes the digest. questionId The question this claim answers. Claims point at questions, not the reverse. confidence How confident the author is - one of low , medium , high , verified . It is a stated position, not a computed score; engine verdicts live on Evaluation instead. Records are stored one per line as .ndjson . Unset means the key is absent - ARF never writes null to mean “no value”.
How records reference each other
Every link is a UUID held in a field on the pointing record. Read X --field--> Y as “X.field holds the id of Y”.
Claim --questionId----> Question
Claim --sourceIds-----> Source
Artifact --claimIds------> Claim
Artifact --sourceIds-----> Source
Evaluation --claimIds------> Claim
Challenge --claimId-------> Claim
Challenge --challengerId--> Validator
Endorsement --claimId-------> Claim
Endorsement --validatorId---> Validator
Revision --artifactId----> Artifact
ChainRecord --artifactId----> Artifact
ArtifactBundle { artifact, question, claims[], sources[], publishedBy, publishedAt }
this whole object is the value that gets hashed An ArtifactBundle is the publish-time envelope: one artifact, its question, and the claims and sources it used, collected into a single JSON object. That object is the thing that gets hashed.
Grouped by role. Required fields are exactly those marked required in each schema.
Question examples/question.ndjson The question under evaluation.
required: id, text, status, createdAt
Source examples/source.ndjson An input that can be cited.
required: id, type, title, createdAt
Artifact examples/artifact.ndjson A written output assembled from claims and sources.
required: id, title, type, body, claimIds, sourceIds, createdAt
What was asserted, and what an engine concluded.
Claim examples/claim.ndjson One assertion, plus the sources it rests on.
required: id, statement, confidence, sourceIds, createdAt
Evaluation schemas/evaluation.json An engine result about claims. The engine payload is stored verbatim.
required: id, claimIds, engine, protocolVersion, payload, createdAt
Who checked it, and what they said.
Validator examples/validator.ndjson A human or agent that reviews claims.
required: id, name, type, createdAt
Challenge examples/challenge.ndjson A recorded objection to a claim.
required: id, claimId, challengerId, reason, createdAt
Endorsement examples/endorsement.ndjson Recorded support for a claim, carrying a weight.
required: id, claimId, validatorId, weight, createdAt
How an artifact changed, and what was published.
Revision schemas/revision.json A point-in-time snapshot of an artifact.
required: id, artifactId, version, snapshot, createdAt
ChainRecord schemas/chain-record.json A publication record carrying the bundle hash.
required: id, artifactId, contentHash, chain, txHash, publishedAt
Evaluation , Revision and ChainRecord link to their schema because no example record is published yet. For Evaluation that is deliberate: the claimIds link cannot be recovered from an engine payload, and a fabricated example would misrepresent the field most likely to be misunderstood.
Canonicalization means reducing a JSON value to one exact byte sequence, so any correct implementation produces the same bytes. Four stages, one operation.
{
"id" : "602d54e3-0650-40c8-898c-e6d3c0e71d71" ,
"claimId" : "6251dbd4-a1f2-4b8b-9cd0-594a46a2f5c3" ,
"validatorId" : "d2f602ae-9be6-43c0-8c1e-622980547d49" ,
"comment" : "Strong evidence from the Shumailov et al. paper. Model collapse is well-documented." ,
"weight" : 5 ,
"createdAt" : "2026-03-25T20:17:56.727Z"
} 2. Canonical form (RFC 8785) 293 bytes · no trailing newline Copy {"claimId":"6251dbd4-a1f2-4b8b-9cd0-594a46a2f5c3","comment":"Strong evidence from the Shumailov et al. paper. Model collapse is well-documented.","createdAt":"2026-03-25T20:17:56.727Z","id":"602d54e3-0650-40c8-898c-e6d3c0e71d71","validatorId":"d2f602ae-9be6-43c0-8c1e-622980547d49","weight":5} 3. SHA-256 of those bytes Copy 56c223f1dac62072d1999e00cd3ba854352682294425fa10ad102f1a0c73ea6e 4. Encoded hash Copy sha256:56c223f1dac62072d1999e00cd3ba854352682294425fa10ad102f1a0c73ea6e Check it yourself
Paste step 2 into any SHA-256 and you get step 3.
printf '%s' '{"claimId":"6251dbd4-a1f2-4b8b-9cd0-594a46a2f5c3","comment":"Strong evidence from the Shumailov et al. paper. Model collapse is well-documented.","createdAt":"2026-03-25T20:17:56.727Z","id":"602d54e3-0650-40c8-898c-e6d3c0e71d71","validatorId":"d2f602ae-9be6-43c0-8c1e-622980547d49","weight":5}' | shasum -a 256 The three rules that decide the bytes
Object keys are sorted by UTF-16 code unit, at every depth. Reordering keys does not change the digest.
Array order is preserved and is significant. Reordering sourceIds is a different record, and must produce a different digest.
The digest covers the canonical bytes , not a re-parsed or pretty-printed string. There is no trailing newline.
Values with no JSON representation - NaN , Infinity , undefined - are rejected rather than coerced. Emitting a hash for an unrepresentable value would produce a proof nobody can reproduce.
The same operation applied to an ArtifactBundle gives its contentHash . bundle.ndjson records sha256:81ea9018…db24d673 , and recomputing it this way reproduces that value.
The algorithm is versioned as rfc8785-jcs-v1 , carried inside the bundle. An implementation that meets a version string it does not recognise must reject it rather than fall back to a default - silently hashing with a different algorithm produces a confidently wrong answer.
Offline. No network, no $id dereferencing.
git clone https://github.com/korvohq/atlas.git
cd atlas
npm ci
npm run arf:validate -- spec/v0.1/examples/*.ndjson Output Copy PASS spec/v0.1/examples/artifact.ndjson:1 (artifact)
PASS spec/v0.1/examples/bundle.ndjson:1 (bundle)
PASS spec/v0.1/examples/challenge.ndjson:1 (challenge)
PASS spec/v0.1/examples/claim.ndjson:1 (claim)
...
PASS spec/v0.1/examples/validator.ndjson:2 (validator)
16 passed, 0 failed Records are checked against the same JSON Schemas the server uses. The type is inferred from the record shape, or forced with --type=claim .
Everything authoritative lives under spec/ , licensed Apache-2.0 . The Korvo Atlas server remains AGPL-3.0-only.
spec/v0.1/README.md How the pieces fit together.
canonicalization.md Hashing rules, conformance vectors, porting checklist.
disclosure.md Whether an author intends a record to be published.
bundle.schema.json JSON Schema for the ArtifactBundle.
examples/ The seed records used by the command above.
Examples are seed records , exported read-only from development fixtures - not production data. Nothing in spec/v0.1/examples/ came from a live deployment.
Public archive of runs: korvo.xyz/atlas . This page: korvo.xyz/arf . The namespace korvo.xyz/arf/ns/v0.1 names types and is not served.
