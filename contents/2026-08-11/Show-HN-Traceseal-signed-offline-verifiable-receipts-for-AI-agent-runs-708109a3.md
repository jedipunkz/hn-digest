---
source: "https://traceseal.io/"
hn_url: "https://news.ycombinator.com/item?id=49254245"
title: "Show HN: Traceseal – signed, offline-verifiable receipts for AI agent runs"
article_title: "Traceseal — The Agent Accountability Platform"
author: "traceseal"
captured_at: "2026-08-11T06:48:00Z"
capture_tool: "hn-digest"
hn_id: 49254245
score: 1
comments: 0
posted_at: "2026-08-11T06:45:53Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Traceseal – signed, offline-verifiable receipts for AI agent runs

- HN: [49254245](https://news.ycombinator.com/item?id=49254245)
- Source: [traceseal.io](https://traceseal.io/)
- Score: 1
- Comments: 0
- Posted: 2026-08-11T06:45:53Z

## Translation

タイトル: HN を表示: Traceseal – AI エージェント実行の署名済みのオフライン検証可能な領収書
記事のタイトル: Traceseal — エージェント責任プラットフォーム
説明: AI エージェントの作業のための検証可能な領収書、グラウンドトゥルース検証、および再現可能なガバナンス。エージェントが行ったことの暗号化された証拠 - 誰でも検証可能。

記事本文:
EU AI 法第 50 条 — AI システムの透明性義務は 2026 年 8 月 2 日に発効します。領収書はコンプライアンスを証明する方法です。
トレースシール。イオ
プラットフォーム
AI エージェントが何をしたかを示す暗号化された証拠。
検証可能な領収書、真実の検証、およびエージェントの作業のための再現可能なガバナンス。すべての呼び出しの改ざんが明らかな監査証跡 - 誰でも検証可能で、オペレーターへのアクセスは必要ありません。
エージェントエコノミーのための信頼インフラストラクチャ。
エージェントは実際の仕事、つまりコードを書いたり、お金を動かしたり、制作に携​​わったりしています。 Traceseal はその下にあるアカウンタビリティ層です。実行された内容を証明し、グラウンド トゥルースと照合して検証し、ガバナンスの下で再現します。
すべてのエージェント呼び出しは、スキル ID、コンテンツアドレス指定されたマニフェスト、入出力ハッシュ、サンドボックス プロファイル、オペレーター署名などの署名付きの自己完結型レシートを発行します。
正規の JSON — 改ざんが封印を破る
認証は「実行された」だけではありません。コマンドはカーネル名前空間サンドボックス内で再実行され、禁止されたアクションは単に監視されるだけでなくサンドボックス自体によって強制されます。
bwrap 分離 - 読み取り専用ルート、ネットワークなし
領収書は透明性ログに連鎖します。第三者は 1 つのコマンドで履歴全体を再検証できます。オペレーターのマシン、キー、監査証跡にはアクセスできません。
終了コード: 0 有効 · 1 無効
3つの役割。検証可能な証拠が 1 つあります。
すべてのレシートには、何が実行されたか、誰がコードを承認したか、誰が実行を保証したかという 3 つのセクションがあります。第三者は 3 つすべてを検証できます。
公開者は、ed25519 キーを使用してスキル バンドルに署名します。署名は、透明性ログに記録される、コンテンツに対応したマニフェストを介してすべてのソース ファイルをカバーします。
対象: すべてのソース ファイルのマニフェスト
公開: 透明性ログエントリ
オペレーターは、カーネル名前空間サンドボックス内で署名付きスキルを実行します。実行時の記録

入力、出力、タイミング、サンドボックス構成をハッシュとして取得し、レコードに署名します。
レコード: 入力/出力コンテンツのハッシュ
署名済み: オペレーター ed25519 の署名
サードパーティは、traceseal-verify をインストールし、1 つのコマンドを実行します。検証者は、正規の JSON 上でオペレーターの署名をチェックします。オペレーター接続がありません。信頼性の仮定はありません。
終了コード: 0 有効・1 無効
EU AI 法では、エージェントの透明性が義務付けられています。
第 50 条の透明性義務は、2026 年 8 月 2 日に発効します。AI システムを導入する組織は、システムが何を行ったかを主張するのではなく、実証する必要があります。署名された実行受領書を使用すると、文書化作業が 1 つのコマンドの検証に変わります。第50条を読む→
オープン仕様。ベリファイアを開きます。公開ログ。
第三者がエージェントの責任を問うために必要なものはすべて、受領書仕様、検証者、実行時認証ツール、透明性ログ自体など、すべて公開されます。
独立した検証者。コマンド 1 つ、receipt.json 1 つ、ゼロトラスト前提。オープンレシート仕様で発送します。
実行時認証: サンドボックス内での実行を記録し、入力と出力をハッシュし、オペレーター キーでレシートに署名します。
ドロップイン LangChain 統合 — エージェント コードを変更せずに、チェーンおよびツール呼び出しの署名済みレシートを作成します。
トラスト ハーネス: エージェントの CLI を評価する再現可能な動作プローブ。署名付き、証人共同署名付きの Traceseal レシートを発行します。
透明性ログ: log.traceseal.io
30 秒以内に領収書を確認してください。
アカウントがありません。 APIキーがありません。領収書とオープンソース検証ツールだけです。

## Original Extract

Verifiable receipts, ground-truth verification and replayable governance for AI agent work. Cryptographic proof of what your agent did — verifiable by anyone.

EU AI Act, Article 50 — transparency obligations for AI systems take effect 2 August 2026 . Receipts are how you prove compliance.
traceseal . io
Platform
Cryptographic proof of what your AI agent did.
Verifiable receipts, ground-truth verification and replayable governance for agent work. A tamper-evident audit trail for every invocation — verifiable by anyone, no access to the operator required.
Trust infrastructure for the agent economy.
Agents are doing real work — writing code, moving money, touching production. Traceseal is the accountability layer underneath: prove what ran, verify it against ground truth, replay it under governance.
Every agent invocation emits a signed, self-contained receipt: skill identity, content-addressed manifest, input/output hashes, sandbox profile, operator signature.
canonical JSON — any tamper breaks the seal
Attestation goes beyond "it ran": commands re-execute inside a kernel-namespace sandbox, and forbidden actions are enforced by the sandbox itself, not merely observed.
bwrap isolation — read-only root, no network
Receipts chain into a transparency log. Any third party can re-verify the whole history with one command — no access to the operator's machine, keys or audit trail.
exit codes: 0 valid · 1 invalid
Three roles. One verifiable proof.
Every receipt has three sections: what ran, who authorised the code, and who vouches for the execution. Any third party can verify all three.
The publisher signs the skill bundle with an ed25519 key. The signature covers every source file via a content-addressed manifest, recorded in the transparency log.
Covers: manifest over all source files
Published: transparency log entry
The operator runs the signed skill inside a kernel-namespace sandbox. The runtime records inputs, outputs, timing and sandbox configuration as hashes, then signs the record.
Records: input/output content hashes
Signed: operator ed25519 signature
Any third party installs traceseal-verify and runs one command. The verifier checks the operator's signature over the canonical JSON. No operator connection. No trust assumptions.
Exit codes: 0 valid · 1 invalid
The EU AI Act makes agent transparency mandatory.
Article 50 transparency obligations take effect on 2 August 2026. Organisations deploying AI systems will need to demonstrate what their systems did — not assert it. Signed execution receipts turn that from a documentation exercise into a one-command verification. Read Article 50 →
Open spec. Open verifier. Public log.
Everything a third party needs to hold an agent to account is public — the receipt spec, the verifier, the runtime attestation tooling, and the transparency log itself.
The independent verifier. One command, one receipt.json, zero trust assumptions. Ships with the open receipt spec.
Runtime attestation: records execution inside the sandbox, hashes inputs and outputs, and signs the receipt with the operator key.
Drop-in LangChain integration — signed receipts for chain and tool invocations without changing your agent code.
The trust harness: reproducible behavioural probes that grade agent CLIs, emitting signed, witness-cosigned Traceseal receipts.
Transparency log: log.traceseal.io
Verify a receipt in 30 seconds.
No account. No API key. Just a receipt and the open-source verifier.
