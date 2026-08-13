---
source: "https://apify.com/filipmajchrzak/ai-claim-checkpoint"
hn_url: "https://news.ycombinator.com/item?id=49290923"
title: "I built a lint for AI-generated evidence"
article_title: "Ai Claim Readiness Chackpoint · Apify"
author: "fmajchrzak"
captured_at: "2026-08-13T19:51:48Z"
capture_tool: "hn-digest"
hn_id: 49290923
score: 1
comments: 0
posted_at: "2026-08-13T19:40:38Z"
tags:
  - hacker-news
  - translated
---

# I built a lint for AI-generated evidence

- HN: [49290923](https://news.ycombinator.com/item?id=49290923)
- Source: [apify.com](https://apify.com/filipmajchrzak/ai-claim-checkpoint)
- Score: 1
- Comments: 0
- Posted: 2026-08-13T19:40:38Z

## Translation

タイトル: AI が生成した証拠用の lint を構築しました
記事タイトル: Ai Claim Readiness Chackpoint · Apify
説明: AI ワークフローの主張に十分な拘束された証拠があるかどうかを、プロモーション、公開、または引き継ぎ前にチェックします。プロフェッショナルおよび B2B AI ワークフロー向けに設計

記事本文:
Ai Claim Readiness Chackpoint · Apify コンテンツにスキップ
AI エージェントとアプリ用のすぐに実行できるツール。どれか 1 つを選んで行ってください。
サーバーレス プログラムを構築して実行する
アプリやサービスと接続する
ブロックされずにこする
Webスクレイピングおよびクローリングライブラリ
MCP クライアントとシームレスに統合するためのアクターとツールを使用して Apify MCP サーバーを構成します。
Apify プラットフォームの完全なリファレンス
Python、JavaScript、TypeScript
初心者から上級者まで楽しめるコース
アクターを公開して報酬を受け取る
先月は140万ドルが支払われました。多くの開発者は 3,000 ドル以上の収入を得ています。
Apify に関するアドバイスと回答
Apify ストアに移動 Ai Claim Readiness チャックポイント
filipmajchrzak/ai-claim-checkpoint
無料で試す このアクターについて質問する AI ワークフローの主張に十分な拘束された証拠があるかどうかを、プロモーション、公開、または引き渡しの前にチェックします。プロフェッショナルおよび B2B AI ワークフロー向けに設計されています。
コミュニティ アクターの統計によって維持されます
README 入力出力 価格 API の問題 変更ログ AI Claim Readiness Checkpoint v0.5 (R7 有料ベータ版)
AI Claim Readiness Checkpoint は、ソフトで確定的なソフトウェア リリースの証拠チェックを行うための、有料の実験的な Apify Actor です。有料のパブリック ベータ版は、ビジネスまたは専門的な立場で活動する企業およびプロフェッショナル ユーザーのみを対象として提供されます。これは消費者向けの製品として意図されたものではありません。ただし、パブリック Apify アクターは、どの Apify ユーザーでも実行できます。そのため、これは購入者に表示される契約上の適格性であり、技術的な消費者アクセス制御ではありません。 v0.5/R7 ホスト型ラッパーは、提出されたテストとレビューのメタデータを 1 つの固定バージョン付きポリシー プリセットに対して評価し、再現可能な証拠バンドルを含む決定を生成します。その評価エンジンは、以下に示すバイト凍結された R6 ランタイムのままです。
アクターは、主張が正しいかどうかを判断するために AI モデルを使用しません。

それは本当です。
パブリック ベータ版は、 software_release_basic_v1 のみをサポートします。
失敗したテストがゼロのテスト証拠。
PASS ステータスの証拠を確認します。
どちらの観測も、指定された論理時間で 30 日以内のものです。
SHA-256 バインディングを使用した正規の JSON 証拠。
アダプターは壁時計を使用したり、欠落した証拠を発明したり、提出されたステータスを変更したり、失敗した証拠を修復したりすることはありません。
YYYY-MM-DDTHH:MM:SSZ 形式で明示的な UTC タイムスタンプを使用します。デフォルトの入力は合成であり、実際のリリース レコードではありません。
{ "inputMode" : "easy" 、 "preset" : "software_release_basic_v1" 、 "claimId" : "public-demo-synthetic-allow-001" 、 "subject" : "SYNTHETIC_PUBLIC_DEMO_NOT_A_REAL_RELEASE" 、 "logicalTime" : "2026-08-09T00:00:00Z" , "tests" : { "observedAt" : "2026-08-08T12:00:00Z" , "passed" : 12 , "failed" : 0 } , "review" : { "observedAt" : "2026-08-08T13:00:00Z" , "ステータス" : "PASS" , "レビューア" : "synthetic-public-demo" } }
証拠が利用できない場合にのみ、テストまたはレビューを null に設定します。
可能な事前設定の決定は次のとおりです。
許可 — 提出された証拠は固定ポリシーを機械的に満たしました。
ブロック — 提出された証拠が必要な管理と矛盾していました。
NEED_MORE_EVIDENCE — 必要な証拠がありませんでした。
Actor プロセスのステータスだけに依存しないでください。以下のすべてが当てはまる場合にのみ、ワークフローを続行します。
OUTPUT.status は VERIFIED_GATE_DECISION です。
VERIFICATION.sidecar_verified は true です。
OUTPUT.決定は VERIFICATION.決定と等しい。
OUTPUT.bundleSha256 および OUTPUT.bundleBytes は、 VERIFICATION.bundle_sha256 および VERIFICATION.bundle_bytes と同じです。
それ以外の場合は、停止して理由コードまたは障害レコードを調べてください。
個別にハッシュ固定された R6 検証ツールは、リプレイ バンドルとその決定をカバーします。 OUTPUT に追加されたパブリック ラッパー メタデータは検証されません。
単に verify_bundle.py を実行しないでください。

信頼できないリプレイ バンドルに存在します。まず、SHA-256 をストア リリースと同時に公開された v0.5/R7 トラスト アンカーと比較します。
5f9ad6f81ea07536582b69d2657962aafcf2c7613564680e7b546ff404350293
たとえば、ファイルを抽出した後、実行する前に次のようにします。
printf '%s %s\n' '5f9ad6f81ea07536582b69d2657962aafcf2c7613564680e7b546ff404350293' 'verify_bundle.py' | sha256sum -c - python verify_bundle.pyevidence_bundle.zip --sidecarevidence_bundle.zip.sha256.txt --trusted-verifier-sha256 5f9ad6f81ea07536582b69d2657962aafcf2c7613564680e7b546ff404350293
外部ハッシュ チェックが失敗した場合は停止します。ダイジェストは、バンドルを再配布した人の身元ではなく、この公開バージョンへのバイト バインディングを確立します。
ALLOW は、リリースの主張が真実であること、安全であること、準拠していること、安全であること、競争力があること、生産準備が整っていること、または販売準備ができていることを証明するものではありません。
有料ユニットは、1 つの完成した検証済みの機械評価とその再生可能な出力バンドルです。これは、ALLOW の結果、認定、法的アドバイス、セキュリティ レビュー、または好ましい結果に対する支払いではありません。
構成された Apify イベントごとの支払いイベントは、verified-decion です。検証済みの決定が ALLOW 、 BLOCK 、または NEED_MORE_EVIDENCE であっても、1 つのイベントの価格は同じです。すべての検証結果レコードがコミットされた後、アクターは RESERVED_NO_RETRY 状態の永続的な BILLING_ATTEMPT レコードを保存し、正確に検証された OUTPUT と一致するフェンスの両方を読み取ります。その場合にのみ、単一の課金データセット リクエストを行うことができます。この一致するフェンスまたはすでにコミットされた検証済みの出力のいずれかを参照する、その後のシリアル化された再起動または復活では、以前の結果が保持され、新しいデータセットやチャージの試行は行われません。不正な入力または拒否された入力、フェイルクローズされた結果、および未検証の結果は、このカスタム イベントを意図的に発行するものではありません。
BILLING_AT

TEMPT は、受領書ではなく、先行書き込み非再試行フェンスです。リクエストが Apify に到達したこと、データセット アイテムが保存されたこと、または請求が受け入れられたことを証明するものではありません。フェンスの永続性またはリードバックが不確実な場合、または最終的なデータセットとチャージのリクエストがあいまいなエラーを返した場合、アクターは検証された Key-Value ストアの結果を提供しますが、再試行は抑制されます。これにより、確認済みの有料イベントやデータセット行がなくても、購入者に確認済みの結果を与えることができます。プレーンな Key-Value ストア フェンスは、シリアル化されたプロセスの再エントリを保護します。これは、仮想的な重複コンテナに対するアトミックな比較作成ロックであるとは主張されておらず、ストレージを削除すると削除できる可能性があります。請求の確認手順については、PAID_BETA_TERMS.md および SUPPORT_POLICY.md を参照してください。
Apify Store の公開オファーは、検証済みの決定ごとに 0.49 米ドルで設定されています。 Verified-Decision は主要かつ唯一の有料イベントであり、実行あたりの最小最大コストは USD 0.49 で、個別の合成アクターの開始と自動データセット アイテムの有料イベントは存在せず、プラットフォーム使用量のパススルーは無効になっています。公開前の 0.5.1 スモーク実行では、chargeEventCounts.verified-decion = 1 が確認されました。購入者は、取引ごとに Apify によって表示される価格と請求条件に依存する必要があります。
すべての検証結果は、まず以下にリストされている 9 つの製品アーティファクトをコミットします。
Key-Value ストア。最終的にチャージされたデータセットに安全に到達する通常の実行
リクエストは 1 つのデータセット項目も書き込みます。運用中の BILLING_ATTEMPT レコード
条件付きです: 検証された OUTPUT の後にのみ書き込まれます。
フェンスの持続性が不確実な場合には存在しません。意図的に含まれていません
summary.artifactKeys は、請求管理レコードではなく、請求管理レコードであるためです。
製品のアーティファクト。
このアクタはソフト チェックポイントです。リリースアクションは実行されず、別のアクションが妨げられることもありません。

特権ワークフローがその結果をバイパスしないようにします。
パブリックベータ版で検証されたすべての出力には次のように記載されています。
supportedPreset=software_release_basic_v1 。
不正な入力、混合入力、不明な入力、大きすぎる入力、または一貫性のない入力は失敗して閉じられます。実行時のバイト変更、予期しない出力ファイル、ダイジェストの不一致、またはベリファイアの不一致もフェールクローズされます。
合成または非機密のソフトウェア リリース メタデータのみを送信してください。
資格情報、秘密、API トークン、秘密キー、個人データ、健康データ、規制されたデータ、クライアントの機密資料、輸出規制されている資料、または法的に特権のある資料を送信しないでください。リリースの件名とレビュー担当者には個人を特定しない識別子を使用します。
アクターは入力フィールドを処理し、決定、送信されたメタデータ、監査レコード、および再生バンドルを実行のデフォルトの Apify Key-Value ストアに書き込みます。最後の課金データセット リクエストに到達した実行では、1 つのデータセット アイテムも送信されます。アクター アプリケーションのログ メッセージには、固定決定または制限付きの失敗コードのみが含まれます。アプリケーションは、送信された識別子を意図的にログに記録しません。 Apify プラットフォームとシステムのロギングは、Actor コードの制御の範囲外です。
Actor コードは、サードパーティのサービスを呼び出したり、開発者が制御するストレージに実行データをコピーしたりしません。 Apify は実行入力、ログ、デフォルトの出力ストレージにのみ使用されます。ネットワークの分離は主張されていません。
アクターは、モデルのトレーニング、広告、または外部分析に入力または出力を使用しません。 Apify プラットフォームのストレージ、保持、削除、インフラストラクチャ、および譲渡に関する条件が適用されます。実行とそのストレージが不要になったら、Apify コンソールまたは API を使用して削除します。
発行者には、ユーザーのプライベート実行を検査するための日常的な権限がありません。サポートの場合は、パブリッシャーがレビューする前に、サニタイズされた情報のみを提供し、明示的にケース固有のアクセスを許可してください。

個人経営またはその保管場所。現在のパブリックベータ版のデータ処理およびサポート ルールについては、PRIVACY_NOTICE.md および SUPPORT_POLICY.md を参照してください。
商用モデルは、上記のカスタム検証済み決定イベントのみを使用し、イベントごとに 0.49 米ドルで支払う Apify です。 Actor は有料のパブリック ベータ版として公開されています。独立ユーザー U1 ～ U5 の検証と外部顧客のアカウンティングは公開後に監視されます。ストアのチェックアウト/実行画面は、有効な取引価格と Apify の料金を決定する権限を持ちます。
発行元は STUDIO MASAŻU FILIP MAJCHRZAK で、所有者は Filip Majchrzak です。ホストされたアクターの実装は独自のものであり、その完全なソースは公的に配布されていません。著作権 © 2026 フィリップ・マイシュルザク。公開により、ホストされたアクタを呼び出す許可が付与されますが、その独自のソースを取得または再配布することは許可されません。
意図的な例外の 1 つは、すべてのリプレイ バンドルに含まれるハッシュ固定された verify_bundle.py ソースです。その検証ソースは、BUNDLED_VERIFIER_LICENSE.md および実行の VERIFIER_LICENSE レコードの MIT ライセンスに基づいてライセンスされています。この例外は、残りのアクター実装のライセンスを付与しません。
ユーザーは送信されたデータに対する権利を保持し、生成された出力を合法的な目的でダウンロードして使用することができます。サードパーティのコンポーネントには、引き続きそれぞれのライセンスが適用されます。
ベータ版は、サービス レベル契約や可用性、応答時間、メンテナンス期間、特定の目的への適合性の保証はなく、現状のまま提供されます。
有料ベータ版の使用には、 PAID_BETA_TERMS.md 、 PRIVACY_NOTICE.md 、Apify の表示条件、および適用法に基づく放棄できない権利も適用されます。ストア価格またはプラットフォーム料金に関してこれらの条件が矛盾する場合、現在の Apify 購入/実行画面がそのプラットフォーム トランザクションを制御します。
「問題」タブを使用してください。

バグや使用方法に関する質問については、彼の Actor の Apify ページを参照してください。サポートはベスト エフォートであり、応答や解決時間の保証はありません。 SLAはありません。
問題の内容が公開される可能性があると想定します。サニタイズされた例、実行 ID、および出力ステータスのみを含めます。入力データ、シークレット、リプレイバンドル、または機密の詳細を決して投稿しないでください。セキュリティ、プライバシー、または請求イベントの問題が疑われる場合は、 fmajchrzak.ai@gmail.com に非公開で報告してください。エクスプロイトの詳細を公開問題に掲載しないでください。完全なサポートと料金レビューの境界は SUPPORT_POLICY.md にあります。
エンジン: OCOI 証拠ゲート ライト 0.1.1
パブリック ランタイム フリーズ: OCOI-EGL-PUBLIC-RUNTIME-FREEZE-20260811-R6
パブリックコア SHA-256: 09997e6c7de8e3045bbf80df5e4470107c2b823f3a200a8113fd3e119e24e84a
評価者 SHA-256: b1a1bd61268d37f7dc6daa37d69d4092dae53472ad3e75fe611742685749dd6c
信頼できるランチャー SHA-256: 41f30710ca9a1498b460acb7341cb386853bc2d59120da530e14e7c0b1aebfbf
ポータブル検証機 SHA-256: 5f9ad6f81ea07536582b69d2657962aafcf2c7613564680e7b546ff404350293
レコード SHA-256 の凍結: 577944cfa06593e1e6363fe85cd7fef959e8b4aaa3fd27c9d92592778284ec58
過去の R4 および R5 ランタイム ファイルは、証拠保存のためにバイト同一のままです。 v0.5/R7 有料ラッパーは凍結された R6 評価エンジンを使用します

[切り捨てられた]

## Original Extract

Checks whether an AI workflow claim has sufficient bound evidence before promotion, publication, or handoff. Designed for professional and B2B AI workfl...

Ai Claim Readiness Chackpoint · Apify Skip to content Get started
Ready-to-run tools for your AI agents and apps. Just pick one and go.
Build and run serverless programs
Connect with apps and services
Scrape without getting blocked
Web scraping and crawling library
Configure your Apify MCP server with Actors and tools for seamless integration with MCP clients.
Full reference for the Apify platform
Python, JavaScript, and TypeScript
Courses for beginners and experts
Publish your Actors and get paid
$1.4M paid out last month. Many developers earn over $3k.
Advice and answers about Apify
Go to Apify Store Ai Claim Readiness Chackpoint
filipmajchrzak/ai-claim-checkpoint
Try for free Ask questions about this Actor Checks whether an AI workflow claim has sufficient bound evidence before promotion, publication, or handoff. Designed for professional and B2B AI workflows.
Maintained by Community Actor stats
README Input Output Pricing API Issues Changelog AI Claim Readiness Checkpoint v0.5 (R7 paid beta)
AI Claim Readiness Checkpoint is a paid experimental Apify Actor for a soft, deterministic software-release evidence check. The paid public beta is intended and offered exclusively for businesses and professional users acting in a business or professional capacity; it is not intended as a consumer offering. A public Apify Actor can nevertheless be run by any Apify user, so this is buyer-visible contractual eligibility, not a technical consumer-access control. The v0.5/R7 hosted wrapper evaluates submitted test and review metadata against one fixed, versioned policy preset and produces a decision with a reproducible evidence bundle. Its evaluation engine remains the byte-frozen R6 runtime identified below.
The Actor does not use an AI model to judge whether a claim is true.
The public beta supports only software_release_basic_v1 .
test evidence with zero failed tests;
review evidence with PASS status;
both observations no more than 30 days old at the supplied logical time;
canonical JSON evidence with SHA-256 bindings.
The adapter does not use the wall clock, invent missing evidence, change submitted statuses, or repair failed evidence.
Use explicit UTC timestamps in YYYY-MM-DDTHH:MM:SSZ form. The default input is synthetic and is not a real release record.
{ "inputMode" : "easy" , "preset" : "software_release_basic_v1" , "claimId" : "public-demo-synthetic-allow-001" , "subject" : "SYNTHETIC_PUBLIC_DEMO_NOT_A_REAL_RELEASE" , "logicalTime" : "2026-08-09T00:00:00Z" , "tests" : { "observedAt" : "2026-08-08T12:00:00Z" , "passed" : 12 , "failed" : 0 } , "review" : { "observedAt" : "2026-08-08T13:00:00Z" , "status" : "PASS" , "reviewer" : "synthetic-public-demo" } }
Set tests or review to null only when that evidence is unavailable.
Possible preset decisions are:
ALLOW — the submitted evidence mechanically satisfied the fixed policy;
BLOCK — submitted evidence contradicted a required control;
NEED_MORE_EVIDENCE — required evidence was missing.
Do not rely on the Actor process status alone. Continue a workflow only when all of these are true:
OUTPUT.status is VERIFIED_GATE_DECISION ;
VERIFICATION.sidecar_verified is true ;
OUTPUT.decision equals VERIFICATION.decision ;
OUTPUT.bundleSha256 and OUTPUT.bundleBytes equal VERIFICATION.bundle_sha256 and VERIFICATION.bundle_bytes .
Otherwise stop and inspect the reason codes or failure record.
The separately hash-pinned R6 verifier covers the replay bundle and its decision. It does not verify the public wrapper metadata added to OUTPUT .
Never execute verify_bundle.py merely because it is present in an untrusted replay bundle. First compare its SHA-256 with the v0.5/R7 trust anchor published alongside the Store release:
5f9ad6f81ea07536582b69d2657962aafcf2c7613564680e7b546ff404350293
For example, after extracting the file but before executing it:
printf '%s %s\n' '5f9ad6f81ea07536582b69d2657962aafcf2c7613564680e7b546ff404350293' 'verify_bundle.py' | sha256sum -c - python verify_bundle.py evidence_bundle.zip --sidecar evidence_bundle.zip.sha256.txt --trusted-verifier-sha256 5f9ad6f81ea07536582b69d2657962aafcf2c7613564680e7b546ff404350293
Stop if the external hash check fails. The digest establishes byte binding to this published version, not the identity of whoever redistributed a bundle.
ALLOW does not prove that a release claim is true, safe, compliant, secure, competitive, production-ready, or sale-ready.
The paid unit is one completed, verified mechanical evaluation and its replayable output bundle. It is not a payment for an ALLOW result, certification, legal advice, security review, or a favorable outcome.
The configured Apify pay-per-event event is verified-decision . One event has the same price whether the verified decision is ALLOW , BLOCK , or NEED_MORE_EVIDENCE . After all verified result records have been committed, the Actor stores a durable BILLING_ATTEMPT record with state RESERVED_NO_RETRY , then reads back both the exact verified OUTPUT and the matching fence; only then may it make the single charged Dataset request. A later serialized restart or resurrection that sees either this matching fence or the already committed verified OUTPUT preserves the earlier result and makes no new Dataset or charge attempt. Malformed or rejected input, fail-closed outcomes, and unverified outcomes do not intentionally emit this custom event.
BILLING_ATTEMPT is a write-ahead no-retry fence, not a receipt: it does not prove that the request reached Apify, that a Dataset item was stored, or that a charge was accepted. If fence persistence or read-back is uncertain, or if the final Dataset-and-charge request returns an ambiguous error, the Actor delivers the verified Key-value store result but suppresses any retry. This can give the buyer a verified result without a confirmed paid event or Dataset row. The plain Key-value store fence protects serialized process re-entry; it is not claimed to be an atomic compare-and-create lock for hypothetical overlapping containers, and storage deletion can remove it. See PAID_BETA_TERMS.md and SUPPORT_POLICY.md for the billing-review procedure.
The public Apify Store offer is configured at USD 0.49 per verified-decision . verified-decision is the primary and only paid event, the minimum max cost per run is USD 0.49, separate synthetic Actor start and automatic Dataset-item paid events are absent, and platform-usage pass-through is disabled. A pre-publication 0.5.1 smoke run confirmed chargedEventCounts.verified-decision = 1 . Buyers should rely on the price and billing terms displayed by Apify for each transaction.
Every verified outcome first commits the nine product artifacts listed below to
the Key-value store. A normal run that safely reaches the final charged Dataset
request also writes one Dataset item. The operational BILLING_ATTEMPT record
is conditional: it is written only after the verified OUTPUT , and it can be
absent when fence persistence is uncertain. It is intentionally not included in
summary.artifactKeys because it is a billing-control record rather than a
product artifact.
This Actor is a soft checkpoint. It executes no release action and does not prevent a separately privileged workflow from bypassing its result.
Every public-beta verified output states:
supportedPreset=software_release_basic_v1 .
Malformed, mixed, unknown, oversized, or inconsistent input fails closed. Runtime-byte changes, unexpected output files, digest disagreement, or verifier disagreement also fail closed.
Submit only synthetic or non-sensitive software-release metadata.
Do not submit credentials, secrets, API tokens, private keys, personal data, health data, regulated data, confidential client material, export-controlled material, or legally privileged material. Use non-personal identifiers for the release subject and reviewer.
The Actor processes the input fields and writes the decision, submitted metadata, audit records, and replay bundle to the run's default Apify Key-value store. A run that reaches the final charged Dataset request also submits one Dataset item. Actor application log messages contain only a fixed decision or bounded failure code; the application does not intentionally log submitted identifiers. Apify platform and system logging is outside the Actor code's control.
The Actor code does not call third-party services or copy run data to developer-controlled storage. It uses Apify only for run input, logs, and default output storage. Network isolation is not claimed.
The Actor does not use inputs or outputs for model training, advertising, or external analytics. Apify platform storage, retention, deletion, infrastructure, and transfer terms apply. Delete runs and their storages through the Apify Console or API when they are no longer required.
The publisher does not have routine permission to inspect a user's private run. For support, provide only sanitized information and grant explicit, case-specific access before the publisher reviews a private run or its storage. See PRIVACY_NOTICE.md and SUPPORT_POLICY.md for the current public-beta data-handling and support rules.
The commercial model is Apify pay per event at USD 0.49 using only the custom verified-decision event described above. The Actor is published as a paid public beta. Independent-user U1–U5 validation and external-customer accounting are monitored post-publication. The Store checkout/run screen is authoritative for the effective transaction price and any Apify charges.
The publisher is STUDIO MASAŻU FILIP MAJCHRZAK , owned by Filip Majchrzak. The hosted Actor implementation is proprietary and its full source is not publicly distributed. Copyright © 2026 Filip Majchrzak. Public availability grants permission to invoke the hosted Actor, not to obtain or redistribute its proprietary source.
One deliberate exception is the hash-pinned verify_bundle.py source included in every replay bundle. That verifier source is licensed under the MIT License in BUNDLED_VERIFIER_LICENSE.md and the run's VERIFIER_LICENSE record. This exception does not license the rest of the Actor implementation.
Users retain their rights in submitted data and may download and use generated outputs for lawful purposes. Third-party components remain subject to their respective licenses.
The beta is provided as-is and as-available, without a service-level agreement or guaranteed availability, response time, maintenance period, or fitness for a particular purpose.
Use of the paid beta is also subject to PAID_BETA_TERMS.md , PRIVACY_NOTICE.md , Apify's displayed terms, and any non-waivable rights under applicable law. If these terms conflict about a Store price or platform charge, the current Apify purchase/run screen controls that platform transaction.
Use the Issues tab on this Actor's Apify page for bugs and usage questions. Support is best effort and has no guaranteed response or resolution time. There is no SLA.
Assume Issue content may be public. Include only sanitized examples, a run ID, and output status. Never post input data, secrets, replay bundles, or confidential details. Report suspected security, privacy, or billing-event issues privately to fmajchrzak.ai@gmail.com ; do not place exploit details in a public Issue. Full support and charge-review boundaries are in SUPPORT_POLICY.md .
Engine: OCOI Evidence Gate Lite 0.1.1
Public runtime freeze: OCOI-EGL-PUBLIC-RUNTIME-FREEZE-20260811-R6
Public core SHA-256: 09997e6c7de8e3045bbf80df5e4470107c2b823f3a200a8113fd3e119e24e84a
Evaluator SHA-256: b1a1bd61268d37f7dc6daa37d69d4092dae53472ad3e75fe611742685749dd6c
Trusted launcher SHA-256: 41f30710ca9a1498b460acb7341cb386853bc2d59120da530e14e7c0b1aebfbf
Portable verifier SHA-256: 5f9ad6f81ea07536582b69d2657962aafcf2c7613564680e7b546ff404350293
Freeze record SHA-256: 577944cfa06593e1e6363fe85cd7fef959e8b4aaa3fd27c9d92592778284ec58
The historical R4 and R5 runtime files remain byte-identical for evidence preservation; the v0.5/R7 paid wrapper uses the frozen R6 evaluation engine

[truncated]
