---
source: "https://signatrust.net"
hn_url: "https://news.ycombinator.com/item?id=48481226"
title: "Do AI agents need cryptographically verifiable decision receipts?"
article_title: "Signatrust — Trust, verification & insurance layer for AI decisions"
author: "abokenan444"
captured_at: "2026-06-10T21:46:15Z"
capture_tool: "hn-digest"
hn_id: 48481226
score: 1
comments: 1
posted_at: "2026-06-10T19:14:24Z"
tags:
  - hacker-news
  - translated
---

# Do AI agents need cryptographically verifiable decision receipts?

- HN: [48481226](https://news.ycombinator.com/item?id=48481226)
- Source: [signatrust.net](https://signatrust.net)
- Score: 1
- Comments: 1
- Posted: 2026-06-10T19:14:24Z

## Translation

タイトル: AI エージェントには暗号的に検証可能な意思決定受領書が必要ですか?
記事のタイトル: Signatrust — AI 意思決定のための信頼、検証、保険レイヤー
説明: Signatrust は、自律システムの信頼層です。 OpenAPI により API が理解しやすくなりました。 OAuth により ID が委任可能になりました。 Signatrust は、AI による決定を証明可能にします。つまり、エージェントが行うすべてのアクションに対して、署名された検証可能な決定受領書が作成されます。

記事本文:
シグナトラスト
トラストチェーン
プラットフォーム
サイン
ゼロデータアクセス
パスポート
コンプライアンス
リスクと保険
受信確認
価格設定
ドキュメント
サインイン
始めましょう
AI 意思決定のための信頼、検証、保険レイヤー
すべての AI の決定に署名します。確認してください。保険をかけましょう。
Signatrust は、AI エージェントと自動化システムのための独立した信頼レイヤーです。
あらゆる決定を署名済みの領収書、監査された信頼スコア、保険対象のリスク プロファイルに変換します。
データが境界を離れることはありません。
署名・Ed25519 領収書
検証 · 独立した信頼スコア
AIリスクを保証・引き受け可能
代理人が行為を行い、誰もそれを証明できない場合
すでに起こった3つのこと。誰も領収書を持っていませんでした。
金融代理店は 02:14 に 487,000 ユーロの送金を承認しました。 6 か月後、監査人がどのポリシーが適用され、誰がその通話をレビューしたかを尋ねましたが、誰も何も証明できませんでした。
AIが住宅ローンを拒否した。規制当局は貸し手に根拠を提示するのに14日間の猶予を与えた。システムにはログがありましたが、どれも署名されておらず、改ざんが明らかでもなく、許容できるものでもありませんでした。
DevOps エージェントは、フリーズ ウィンドウ中の 03:40 に本番環境を変更しました。 ３チームがお互いを指し合った。 「エージェントがやりました」では、SLA が破棄された企業顧客を含め、誰も満足しませんでした。
体験談ではなく、匿名化された業界パターン。今日では、これらすべてに、署名されたチェーンリンクされた決定受領書があり、監査人、規制当局、または取引相手は、それを記録した当事者を信頼することなく、数秒で検証できます。
確認済み
STR-7F3A21C9D4
決定受領書 · Ed25519 · スペック v1.0
高リスク
エージェントファイナンスボット
モデル openai · gpt-4o · 2026.4
融資_拒否の決定
人間によるレビュー はい
入力 sha256:6718c8a836…abbcf768
出力 sha256:8a3466a5e7…8882c42a
ポリシー
eu-ai-act-高リスク内部クレジット-v3
Ed25519 署名
fj9FpBMPgPdEBlkZ44オサホ350+h1M

engjAANyrLoKXPv5VBI+uXN8UI+hD8MIWHbL8LZhYVK2veA+O3GDSCw==
監査証跡
決定を取得 2026-06-07T10:00:00Z
フィンガープリント付き SHA-256 の入出力をあなたの側で
正規化およびハッシュ化された領収書のシーケンス # 42
以前のレシートの改ざん防止にリンクされています
背骨が一本。 3層。たった 1 つの意思決定から保険が適用される AI リスクまで。
ほとんどのツールは、ロギング、監査、または保険などの 1 つのスライスを個別にカバーします。 Signatrust は単一の証拠ベースに基づいてチェーン全体を実行するため、その上のすべてのレイヤーは自己証明された請求ではなく、実際の領収書に基づいています。
エージェント、モデル、またはワークフローは、ローンの承認、支払いの実行、契約の締結、デプロイの出荷など、実際のアクションを実行します。
決定は、署名された連鎖決定受領書に封印されます: 誰が、いつ、どの権限の下で、どのモデルで行動したか — Ed25519 署名。
領収書は、監査人、AI 対 AI レビュー、知識ルール、または人間の専門家によって独立してチェック可能であり、システムへのアクセスは必要ありません。
検証された履歴は、エージェントの信頼スコア (0 ～ 100) (整合性、監視、ガバナンス、寿命) となり、レシートのみから再計算できます。
組織の代理店全体で集計: リスクレベル別のエクスポージャ、管理強度、保険等級、および相対引受指数。
保険会社は、今日の信用保険や自動車保険の価格設定と同じように、推測ではなく実際の証拠に基づいて AI 賠償責任保険の価格を設定します。
決定→受領→検証→信頼スコア→リスクスコア→保険。各レイヤーは個別にチェック可能です。自己証明されたものは何もありません。
実証済みのエンジンが 1 つあります。上位の製品ファミリー。
すべての Signatrust 製品は、同じ追加専用の Ed25519 署名付き台帳上の異なる表面です。新しい業種は信頼モデルをフォークすることなくプラグインできるため、信頼スコア、コンプライアンス、リスク プロファイルは人間、エージェント、オートマ全体にわたって一貫したままになります。

テッドシステム。
保管権なしでドキュメントに署名または承認します。ファイルはブラウザでハッシュされ、ハッシュのみが保存され、各署名は検証可能な領収書となります。
AI エージェントを登録し、そのモデルと権限を宣言し、API キーとポリシーを管理します。実行されるすべてのアクションは封印され、監査できます。
署名エンジン: 正規本文、Ed25519 署名、追加専用チェーン、およびベンダーが実装できるオープンレシート仕様。
独立した検証と専門家
システムにアクセスすることなく、誰でも領収書を再確認できます。人間の専門家のネットワークは、リスクの高い決定に副署名することができます。
エージェントごとの 0 ～ 100 の評判。レシートのみから再計算可能 - 誠実さ、危険な通話の監視、ガバナンス、寿命。
エクスポージャー、コントロールの強さ、保険等級、相対引受指数。保険会社ポータルが最上位にあります。
以下に基づく義務をサポートするために構築されています
エージェントが行動すると何が起こりますか?誰が何をしたか誰も証明できません。
データベースの行から、システムが何が起こったのかがわかります。記録が変更されなかったのか、誰の権限で実行されたのか、どのモデルで実行されたのかを証明するものではありません。今日の自律エージェントは毎分実際の行動を起こしており、その後の質問は監査人、規制当局、取引相手が実際に尋ねるものです。
間違ったサプライヤーに 100,000 ユーロを送金
誰が許可したのですか？どのモデルとバージョンですか?人間も巻き込まれていたのでしょうか？電話の時点で利用限度額は確認されましたか?
出願人は推論に対する法的権利を有します。どのポリシーが適用され、どの機能が重要で、どのモデル バージョンが答えを生み出したのでしょうか?
この取引相手に対して署名権限はありましたか?どの上限以内でしょうか？署名が完了する前に取引相手のリスクチェックが行われましたか?
実稼働インフラストラクチャを変更する
どのエージェント、どの権限、どの変更ウィンドウで

、オンコールの承認は誰にありますか?以前の状態は何だったのでしょうか。そして、誰もログを書き換えていないことを証明できますか?
独立した領収書がなければ、監査、規制当局、裁判所に耐えられる回答はありません。あなた自身のログは、あなた自身が書いたあなたに関する証拠です。 Signatrust は、あなたを信頼することなく、同じ事実に他の人が検証できる署名を提供します。
単なるログ行ではない暗号証拠
それぞれの決定は決定受領書に封印されます。これには、動作したエージェントとモデル、有効な権限とポリシー、入出力のフィンガープリント、タイムスタンプ、Ed25519 署名が含まれており、前の受領書に関連付けられています。
署名入り。すべてのレシートには、発行ノードからの Ed25519 署名が含まれています。
連鎖した。追加専用台帳のハッシュによる領収書リンク。 1 つを変更するとチェーンが切れます。
ポータブル。レシートは自己完結型の JSON であり、誰でも保存、共有、確認できます。
'signatrust' から { Signatrust } をインポートします。
const str = new Signatrust({ apiKey: process.env.SIGNATRUST_API_KEY });
// 入力と出力はローカルでハッシュされ、生データは決して残されません
const { レシート、share_url } = await str.sign({
モデル: { プロバイダー: 'openai' 、名前: 'gpt-4o' 、バージョン: '2026.4' }、
決定: {
タイプ: 'loan_rejection' 、
入力、出力、
リスクレベル: '高' 、
human_review: true 、
権限: [ 'credit.decide' ]、
ポリシー: [ 'eu-ai-act-high-risk' ]、
}、
});
コピー
独立した検証
システムにアクセスしなくても、誰でも領収書を確認できます
信頼性 (本文のハッシュ + 署名) は、レシートだけで証明できます。監査人、規制当局、取引相手、または領収書を持っているエンドユーザーは、オンラインまたはオフラインで数秒で確認できます。
ハッシュマッチ。レシートの本文は、それが主張する値に再ハッシュされます。
署名は有効です。 Ed25519 署名は、公開されたキーと照合してチェックされます。
チェーンはそのままです。領収書のリンク c

台帳の前任者とほぼ同じです。
本物の
ボディハッシュ一致パス
Ed25519 署名有効パス
チェーンリンケージそのままのパス
領収書STR-7F3A21C9D4は本物であり、改ざんされていません。
仕組み
3 行のコード。確認可能な領収書 1 枚。
SDK をエージェントにドロップします。すべての決定はローカルでフィンガープリントされ、封印され、署名され、連鎖されます。
SDK は、マシン上の入力と出力の sha256 を計算します。指紋だけが残ります。
ノードは正規のレシートを作成し、Ed25519、スタンピング モデル、アクセス許可、ポリシーを使用して署名します。
各領収書はハッシュによって前の領収書にリンクされます。改ざんが明らかな追加専用の台帳です。
第三者はあなたのデータにアクセスすることなく、独立して領収書を確認できます。
今すぐ本当の決断に署名してください
これにより、実行中の Signatrust ノードが呼び出されます。デモ エージェントが作成され、あなたの決定は実際の署名済みレシートに封印されて検証されます。また、それを改ざんして検証が失敗するのを確認することもできます。
すべてを確認してください。何も共有しません。
Signatrust は、Signatrust でもユーザーの決定を確認できないように設計されています。デフォルトでは、環境からは何も動作せず、SHA-256 フィンガープリントとメタデータのみが残され、コンテンツは残されません。私たちはデータを収集する会社というよりも、保管庫を構築する会社に近いです。
顧客、ケース、モデルの識別子はありません
生の意思決定コンテンツはなし - sha256 ハッシュのみ
何を共有するかを決めるのはあなたです
Signatrust はデフォルトでは運用データを収集しません。グローバルな信頼ネットワーク、共有リスクベンチマーク、そしてやがては保険モデルは、顧客が貢献を選択したもののみによって動かされ、常に匿名化され集約されます。
最大限の信頼。ベンチマークやトレーニングには貢献しません。銀行、政府、防衛機関のデフォルトの姿勢。
集計シグナルのみを共有 — 意思決定ボリューム、意思決定タイプ、エラー率、hu

男性レビュー率。プロンプト、出力、識別子はありません。
ベンチマーク レポート、比較洞察、価格設定の利点と引き換えに、より豊富な匿名化されたシグナルを自発的に提供します。
証拠に加えて評判とリーチ
オープン レシートの形式はエコシステム全体に広がり、すべてのレシートは、それを発行したエージェントに対して再計算可能な評判シグナルを提供します。
検証可能な履歴から得られる 0 ～ 100 の評判: 誠実さ、危険な通話の監視、ガバナンス、および寿命。
LangChain、CrewAI、AutoGen、または OpenAI Agents SDK でツール呼び出しをラップし、レシートを発行します。
オープンレシート形式とモデルコンテキストプロトコルツールにより、あらゆるエージェントがそのアクションを封印できます。
すべてのエージェントは、誰でも検証できるトラストパスポートを取得します
スコアが見出しです。パスポートがその証拠です。すべてのデータは、監査人が独立して再チェックできる領収書にリンクされています。取引相手や規制当局は、エージェントが自社のシステム内で動作する前にそれを要求でき、いつでも自分自身で再計算できます。
再計算可能。すべての数字は公開台帳上の領収書の関数であり、不透明なスコアリングはありません。
ポータブル。安定した URL にある署名付きスナップショット。ベンダーのアンケート、デューデリジェンス、RFP に埋め込みます。
プライバシーを尊重します。プロンプトや出力ではなく、カウントとレート。 K-匿名性は、公開前に管理されます。
スナップショット署名が有効です
サンプル データ — ライブ ページは実際のレシートから再計算します。
トラストチェーンのレイヤー2とレイヤー3
確認済みの領収書からコンプライアンスおよび AI 保険まで
レシート層が基礎となります。それに加えて、2 つの規制当局グレードの製品が直接提供されます。監査人が再チェックできるコンプライアンス レポートと、保険会社が引き受けることができるリスク プロファイルです。どちらも完全に実際の領収書に基づいており、決して自己証明されたものではありません。
EU AI法、GDPR、NIST AI RMFにマッピングされた規制当局対応レポート

ISO/IEC 42001 — 監査人が独立して再チェックできる検証可能な受領書に裏付けられたすべての管理。署名付きの改ざん防止コピーをワンクリックでエクスポートします。
EU AI 法第 2 条9、11–14 · GDPR 第 2 条。 5、25、30
人間による監視とデータ最小化の証拠
署名付きエクスポート、信頼しなくても検証可能
検証可能な履歴からの保険グレードのリスク プロファイル: リスク レベル別のエクスポージャー、運用管理の強度、保険性スコア、および相対引受インデックス。引受会社向けの比較シグナル — 派生したものであり、発明されたものではありません。
保険スコアと等級 (A ～ F)
相対引受指数 (ベースライン 1.00)
匿名化され、k-protected ネットワークベンチマーク
保険市場への 3 つのコンクリートの橋
リスク プロファイルはマーケティング上の主張ではありません。実際に AI リスクの価格を見積もる人々向けに、すでにクエリ可能でエクスポート可能であり、形作られています。
保険パートナー向けの読み取り専用ダッシュボード: ライブ保険等級、保険料指数、リスク レベル別のエクスポージャ、管理強度、およびサンプル内の領収書のワンクリック検証。
代理店ごとの保険グレード (A ～ F) とスコア
相対保険料指数とネットワーク ベースラインの比較
すべての品目に対する独立した検証
2 つの未認証エンドポイントがすでに運用されています。保険会社はプログラムでプロファイル データを取得し、その上に独自の価格設定モデルを再構築できます。
GET /api/v1/risk/{ag

[切り捨てられた]

## Original Extract

Signatrust is the trust layer for autonomous systems. OpenAPI made APIs understandable. OAuth made identity delegatable. Signatrust makes AI decisions provable — a signed, verifiable Decision Receipt for every action your agents take.

Signatrust
Trust chain
Platform
Sign
Zero data access
Passport
Compliance
Risk & insurance
Verify receipt
Pricing
Docs
Sign in
Get started
Trust, verification & insurance layer for AI decisions
Sign every AI decision. Verify it. Insure it.
Signatrust is the independent trust layer for AI agents and automated systems —
turning every decision into a signed receipt, an audited trust score, and an insurable risk profile.
Your data never leaves your perimeter.
Sign · Ed25519 receipts
Verify · independent trust score
Insure · underwriteable AI risk
When an agent acts and no one can prove it
Three things that already happened. None had a receipt.
A finance agent approved a €487,000 transfer at 02:14 . Six months later, when the auditor asked which policy applied and who reviewed the call, no one could prove a thing.
An AI denied a mortgage. The regulator gave the lender 14 days to produce the rationale. The system had logs — none of them signed, none of them tamper-evident, none of them admissible.
A DevOps agent modified production at 03:40 during a freeze window . Three teams pointed at each other. "The agent did it" satisfied no one — including the enterprise customer whose SLA was now broken.
Anonymised industry patterns, not testimonials. Today, every one of these would have a signed, chain-linked Decision Receipt any auditor, regulator or counterparty can verify in seconds — without trusting the party who logged it.
Verified
STR-7F3A21C9D4
Decision Receipt · Ed25519 · spec v1.0
High risk
Agent FinanceBot
Model openai · gpt-4o · 2026.4
Decision loan_rejection
Human review Yes
Input sha256:6718c8a836…abbcf768
Output sha256:8a3466a5e7…8882c42a
Policies
eu-ai-act-high-risk internal-credit-v3
Ed25519 signature
fj9FpBMPgPdEBlkZ44OsSaHo350+h1MengjAANyrLoKXPv5VBI+uXN8UI+hD8MIWHbL8LZhYVK2veA+O3GDSCw==
Audit trail
Decision captured 2026-06-07T10:00:00Z
Input & output fingerprinted SHA-256 · on your side
Receipt canonicalized & hashed ledger seq # 42
Linked to previous receipt tamper-evident
One spine. Three layers. From a single decision to insurable AI risk.
Most tools cover one slice — logging, or auditing, or insurance in isolation. Signatrust runs the whole chain on a single evidence base, so every layer above is derived from real receipts, not self-attested claims.
An agent, model or workflow takes a real action — approve a loan, execute a payment, sign a contract, ship a deploy.
The decision is sealed into a signed, chained Decision Receipt : who acted, when, under which permissions, with which model — Ed25519-signed.
The receipt is independently checkable by auditors, AI-vs-AI review, knowledge rules, or human experts — no access to your systems needed.
Verified history becomes an Agent Trust Score (0–100) — integrity, oversight, governance, longevity — recomputable from receipts alone.
Aggregated across an organization's agents: exposure by risk level, control strength, insurability grade and a relative underwriting index.
Insurers price AI liability cover on real evidence instead of guesses — the same way credit and motor insurance are priced today.
Decision → Receipt → Verification → Trust Score → Risk Score → Insurance. Each layer is independently checkable. Nothing is self-attested.
One proof engine. A family of products on top.
Every Signatrust product is a different surface on the same append-only, Ed25519-signed ledger. New verticals plug in without forking the trust model — so Trust Scores, compliance and risk profiles stay consistent across humans, agents and automated systems.
Sign or approve a document with zero custody — the file is hashed in your browser, only the hash is stored, each signature is a verifiable receipt.
Register an AI agent, declare its model and permissions, manage API keys and policies. Every action it takes can be sealed and audited.
The signing engine: canonical bodies, Ed25519 signatures, an append-only chain and an open receipt spec any vendor can implement.
Independent verification & experts
Anyone can re-check a receipt with no access to your systems. A network of human experts can countersign high-risk decisions.
A 0–100 reputation per agent, recomputable from receipts alone — integrity, oversight on risky calls, governance, longevity.
Exposure, control strength, insurability grade and a relative underwriting index. The insurer portal sits on top.
Built to support your obligations under
What happens when the agent acts — and no one can prove who did what?
A database row tells you what your system says happened. It doesn't prove the record wasn't changed, under whose authority it ran, or with which model. Today's autonomous agents take real actions every minute — and the questions that follow are the ones auditors, regulators and counterparties actually ask.
Transfers €100,000 to the wrong supplier
Who authorised it? Which model and version? Was a human in the loop? Was the spending limit checked at the moment of the call?
The applicant has a legal right to the reasoning. Which policy applied, which features mattered, and which model version produced the answer?
Did it have signing authority for this counterparty? Within what cap? Was the counterparty risk-checked before the signature went through?
Modifies production infrastructure
Which agent, with which permissions, in which change window, with whose on-call approval? What was the previous state — and can you prove no one rewrote the log?
Without an independent receipt, none of these have an answer that survives an audit, a regulator, or a court. Your own logs are evidence about you, written by you. Signatrust gives the same facts a signature anyone else can verify — without trusting you.
Cryptographic evidence, not just a log line
Each decision is sealed into a Decision Receipt: the agent and model that acted, the permissions and policies in force, fingerprints of the input and output, a timestamp, and an Ed25519 signature — chained to the previous receipt.
Signed. Every receipt carries an Ed25519 signature from the issuing node.
Chained. Receipts link by hash in an append-only ledger; altering one breaks the chain.
Portable. A receipt is self-contained JSON anyone can store, share and verify.
import { Signatrust } from 'signatrust' ;
const str = new Signatrust({ apiKey: process.env.SIGNATRUST_API_KEY });
// input & output are hashed locally — raw data never leaves
const { receipt, share_url } = await str.sign({
model: { provider: 'openai' , name: 'gpt-4o' , version: '2026.4' },
decision: {
type: 'loan_rejection' ,
input, output,
risk_level: 'high' ,
human_review: true ,
permissions: [ 'credit.decide' ],
policies: [ 'eu-ai-act-high-risk' ],
},
});
Copy
Independent verification
Anyone can check a receipt — without access to your systems
Authenticity (body hash + signature) is provable from the receipt alone. An auditor, regulator, counterparty, or the end user holding their receipt can confirm it in seconds, online or offline.
Hash match. The receipt body re-hashes to the value it claims.
Signature valid. The Ed25519 signature checks against the published key.
Chain intact. The receipt links cleanly to its predecessor in the ledger.
Authentic
Body hash matches pass
Ed25519 signature valid pass
Chain linkage intact pass
Receipt STR-7F3A21C9D4 is authentic and untampered.
How it works
Three lines of code. One verifiable receipt.
Drop the SDK into your agent. Every decision is fingerprinted locally, sealed, signed and chained.
The SDK computes sha256 of your input and output on your machine. Only fingerprints leave.
The node builds a canonical receipt and signs it with Ed25519, stamping model, permissions and policies.
Each receipt links to the previous one by hash — an append-only ledger where tampering is evident.
Any third party can verify the receipt independently, with no access to your data.
Sign a real decision right now
This calls the running Signatrust node. A demo agent is created for you, your decision is sealed into a real signed receipt and verified — and you can tamper with it to watch verification fail.
Verify everything. Share nothing.
Signatrust is engineered so that even Signatrust cannot see your decisions. By default nothing operational leaves your environment — only SHA-256 fingerprints and metadata, never content. We are closer to a company that builds a vault than one that collects data.
No customer, case or model identifiers
No raw decision content — only sha256 hashes
You decide what — if anything — is ever shared
Signatrust collects no operational data by default. A global trust network, shared risk benchmarks and, in time, insurance models are powered only by what customers choose to contribute — always anonymized and aggregated.
Maximum trust. No contribution to benchmarks or training. The default posture for banks, government and defense.
Share only aggregate signals — decision volume, decision types, error rates, human-review rates. No prompts, outputs or identifiers.
Voluntarily contribute richer anonymized signals in exchange for benchmarking reports, comparative insights and pricing benefits.
Reputation and reach, on top of the evidence
An open receipt format spreads across the ecosystem — and every receipt feeds a recomputable reputation signal for the agent that issued it.
A 0–100 reputation derived from verifiable history: integrity, oversight on risky calls, governance and longevity.
Wrap a tool call in LangChain, CrewAI, AutoGen or the OpenAI Agents SDK and emit a receipt.
An open receipt format and a Model Context Protocol tool so any agent can seal its actions.
Every agent earns a Trust Passport — verifiable by anyone
The score is the headline. The passport is the proof. Every datum links back to receipts an auditor can independently re-check. Counterparties and regulators can demand it before they let your agent act in their systems — and recompute it themselves at any time.
Recomputable. Every number is a function of receipts on the public ledger — no opaque scoring.
Portable. A signed snapshot at a stable URL. Embed it in vendor questionnaires, due diligence, RFPs.
Privacy-respecting. Counts and rates, not prompts or outputs. K-anonymity controls before any public disclosure.
Snapshot signature valid
Sample data — the live page recomputes from real receipts.
Layers 2 & 3 of the trust chain
From verified receipts to compliance & AI insurance
The receipt layer is the foundation. On top of it, two regulator-grade products fall out directly: compliance reports an auditor can re-check, and risk profiles an insurer can underwrite — both derived entirely from real receipts, never self-attested.
Regulator-ready reports mapped to the EU AI Act, GDPR, NIST AI RMF and ISO/IEC 42001 — every control backed by verifiable receipts an auditor can independently re-check. Export a signed, tamper-evident copy in one click.
EU AI Act Art. 9, 11–14 · GDPR Art. 5, 25, 30
Human-oversight & data-minimisation evidence
Signed export, verifiable without trusting us
An insurance-grade risk profile from verifiable history: exposure by risk level, the strength of operating controls, an insurability score and a relative underwriting index. A comparative signal for underwriters — derived, never invented.
Insurability score & grade (A–F)
Relative underwriting index (baseline 1.00)
Anonymized, k-protected network benchmarks
Three concrete bridges to the insurance market
The risk profile isn't a marketing claim — it's already queryable, exportable and shaped for the people who actually price AI risk.
A read-only dashboard for insurance partners: live insurability grade, premium index, exposure by risk level, control strength — and one-click verification of any receipt in the sample.
Per-agent insurability grade (A–F) & score
Relative premium index vs. network baseline
Independent verification on every line item
Two unauthenticated endpoints already in production. Underwriters can pull profile data programmatically and rebuild their own pricing models on top.
GET /api/v1/risk/{ag

[truncated]
