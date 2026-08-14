---
source: "https://c2paviewer.com/articles/eu-ai-act-content-credentials"
hn_url: "https://news.ycombinator.com/item?id=49296515"
title: "EU AI Act and C2PA: What Article 50 Requires"
article_title: "EU AI Act and C2PA: What Article 50 Requires for AI Content | C2PA Viewer"
author: "Bluestein"
captured_at: "2026-08-14T09:58:18Z"
capture_tool: "hn-digest"
hn_id: 49296515
score: 2
comments: 0
posted_at: "2026-08-14T09:41:22Z"
tags:
  - hacker-news
  - translated
---

# EU AI Act and C2PA: What Article 50 Requires

- HN: [49296515](https://news.ycombinator.com/item?id=49296515)
- Source: [c2paviewer.com](https://c2paviewer.com/articles/eu-ai-act-content-credentials)
- Score: 2
- Comments: 0
- Posted: 2026-08-14T09:41:22Z

## Translation

タイトル: EU AI 法と C2PA: 第 50 条が要求するもの
記事のタイトル: EU AI 法と C2PA: AI コンテンツに第 50 条が要求するもの | C2PAビューア
説明: EU AI 法の第 50 条では、AI が生成したコンテンツを 2026 年 8 月 2 日までにマークすることが求められています。何が重要で、誰が遵守し、C2PA がどのようにそれを満たしているのかを学びましょう。

記事本文:
EU AI 法と C2PA: 第 50 条が AI コンテンツに要求するもの | C2PA Viewer メインコンテンツにスキップ C2PA Viewer サポートされているデバイスとサービス 記事 お問い合わせ 共有 フィードバック ホーム
規制 • 2026 年 4 月 26 日更新
EU AI 法と C2PA: 第 50 条が AI コンテンツに求めるもの
期限: EU AI 法の第 50 条は、2026 年 8 月 2 日に施行されます。マーキング義務の技術的詳細を解釈する欧州委員会の透明性に関する実務規範は、2026 年 4 月の時点でまだ最終決定段階にありました。この記事は、その日の時点で規制の本文と公開されたコード草案を反映しています。
EU AI 法、正式には規制 (EU) 2024/1689 は、人工知能を管理する欧州連合の法律です。第 50 条は、2026 年 8 月 2 日に発効する透明性義務を定めており、その義務の 1 つは AI によって生成されたコンテンツに固有のものであり、生成 AI システムからの出力は、人工的に生成されたものとして検出できる機械可読形式でマークされなければなりません。欧州委員会の透明性に関する実践規範草案で例として挙げられている技術メカニズムは、 C2PA Content Credentials です。第 50 条は、2026 年 8 月 2 日に発効します。違反に対する罰金は、1,500 万ユーロまたは世界の年間売上高の 3% のいずれか高い方に達します。
EU AI 法第 50 条が実際に要求していること
第 50 条は、特定の AI システムのプロバイダーと導入者に透明性の義務を課します。コンテンツの出所に関して最も重要な条項は第 50 条 (2) です。生成 AI システムのプロバイダーは、合成出力 (音声、画像、ビデオ、またはテキスト) が機械可読形式でマークされ、人工的に生成または操作されたものとして検出可能であることを保証しなければなりません。マーキングは、技術的に可能な限り、効果的で、相互運用可能で、堅牢で信頼性が高くなければなりません。

これら 4 つの形容詞は、欧州委員会が規制と並行して策定している透明性に関する実施規範の下で特定の意味を持っています。
効果的: マーキングは、検証当事者に対して AI によって生成されたコンテンツであることを実際に識別する必要があります。
相互運用性 : プロバイダー独自のツールだけでなく、準拠した検証者はマーキングを読み取ることができなければなりません。
堅牢 : マーキングはフォーマット変換や軽微な編集などの一般的な変換に耐える必要があります。
信頼性: マーキングは改ざんまたは偽造されたかどうかを検証者が検出できるように、改ざん明示的である必要があります。
第 50 条では、その他の透明性義務も規定しています。感情認識システムの導入者はユーザーに通知する必要があり、ディープフェイクを生成する導入者はそれを開示する必要があり、チャットボット スタイルの AI のプロバイダーは、システムと対話する人に対してシステムの人工的な性質を明確にしなければなりません。 50(2) のマーキング義務は、コンテンツ資格情報に直接技術的な影響を与えるものです。
EU AI 法は特に C2PA を義務付けていますか?
いいえ、EU AI 法はテクノロジー中立であり、施行条項に C2PA の名前は記載されていません。この規則は、マーキング標準を機能的な観点（機械可読、効果的、相互運用可能、堅牢、信頼性）で説明し、技術的な実装は標準化団体と欧州委員会の実施基準に委ねています。
実際には、C2PA コンテンツ認証情報は、透明性に関する委員会の実務規範草案で支持されているマーキング技術です。この規範では、Google の SynthID 透かしなどの補完的なシグナルと並んで、4 つの基準すべてを満たす技術ソリューションの例として C2PA が挙げられています。 C2PA は、Adobe、OpenAI、Google によってすでに導入されているオープン スタンダードであり、改ざんを検出可能にする暗号署名が付いています。
つまり、法律は C2PA を要求していません。

しかし、規制エコシステムはそれをしっかりと指摘しています。 C2PA を採用するプロバイダーは、準拠マーキングがどのようなものであるかについて利用可能な最も具体的な公式ガイダンスに準拠しています。
なぜただのウォーターマークではなく C2PA なのか
単純なピクセル透かし (目に見えるロゴまたは目に見えない知覚マーク) は、第 50 条の意味では相互運用できません。各プロバイダーは独自のウォーターマークを発明し、第三者が競合他社のマークを検証することはできず、暗号化された信頼チェーンもありません。 C2PA は、共有ファイル形式 (JUMBF)、共有クレーム スキーマ (JSON-LD)、および共有署名形式 (COSE) を定義することでこの問題を解決します。 1 人の検証者がすべてを読み取ることができます。技術的な詳細については、C2PA マニフェストの内容を参照してください。
第 50 条を遵守しなければならないのは誰ですか?
第 50 条は、生成 AI システムのプロバイダー (モデルを構築および提供する企業) と導入者 (AI を使用して人々に見せるコンテンツを生成する企業) という 2 つのカテゴリーの組織に適用されます。どちらにも義務がありますが、50(2) のマーキング義務はプロバイダーにあります。
地理的な範囲は広いです。同規則の第 2 条では、AI システムを EU 市場に投入する場合、その出力が EU 内で使用される場合、またはそのユーザーが EU 内に所在する場合、あらゆるプロバイダーが対象となります。 OpenAI などの米国拠点のプロバイダー、イスラエルの画像生成スタートアップ、または日本のカメラ メーカーは、ヨーロッパのユーザーがコンテンツに遭遇するとすぐに対象になります。
限られた例外があります。第 50 条第 2 項は、AI が入力内容を実質的に変更しない編集支援機能を実行する場合、または刑事犯罪の発見、予防、捜査、起訴を目的として使用が法律で許可されている場合には適用されません。標準的な生成ユースケース (画像を生成する AI ツール、テキストをビデオに変換するシステム、音声合成サービス) がこの義務に該当します。
いつ

■第50条は施行可能になるのか？
第 50 条は、EU AI 法が 2024 年 8 月 1 日に発効してから 2 年後の 2026 年 8 月 2 日に施行されます。同日が、一般的な透明性規則と汎用 AI モデルに関する義務にも適用されます。この規制は全体的に時間差が設けられています。
2025 年 2 月 2 日 : 禁止行為の禁止が発効
2025年8月2日：汎用AIモデル規定の新規モデル適用開始
2026 年 8 月 2 日 : 第 50 条を含む一般的な透明性義務が施行可能になります
2027 年 8 月 2 日 : 高リスク AI システム ルールが発効
C2PA コンテンツ認証情報が第 50 条をどのように満たすか
C2PA は、暗号化署名されたマニフェストを画像、ビデオ、オーディオ、ドキュメント ファイルに埋め込みます。マニフェストは、どの AI システムがコンテンツを生成または変更したか、コンテンツがいつ作成されたか、およびどの組織が申し立てに署名したかを宣言します。 C2PA 準拠の検証者 (C2PA ビューアを含む) は、このマニフェストを読み取り、公開された信頼リストと照合して署名を確認できます。
これを第 50 条の 4 つの基準にマッピングすると、次のようになります。
効果的: AI プロバイダーによって署名された「c2pa.created」アクションを含む有効な C2PA マニフェストは、コンテンツを AI 生成として明確にマークします。
相互運用性 : C2PA は、コンテンツの出所と信頼性のための連合によって維持されているオープン スタンダードです。準拠ツールは準拠マニフェストを読み取ります
堅牢 : マニフェストは、C2PA 対応ツールによって処理されるほとんどの形式変換と再エンコードに耐えます。単純な再保存によってストリッピングは依然として可能であるため、実施基準では C2PA と目に見えないウォーターマークを組み合わせることが奨励されています。
Reliable : マニフェストは、C2PA Trust List に対して検証可能な証明書チェーンを使用して COSE を使用して署名されます。改ざんにより署名が破られる
いくつかの大手 AI プロバイダーはすでに C2PA を製品に導入しています

ション。 Adobe Firefly、OpenAI DALL-E 3、OpenAI Sora、Google Imagen はすべて、出力にマニフェストを埋め込みます。 Midjourney は、C2PA をサポートしていない最も著名な生成 AI ツールであり、2026 年 8 月までにアプローチが変更されない限り、第 50 条に基づいて直接規制の対象となる可能性があります。各主要プラットフォームの現在のステータスについては、どの AI ツールが C2PA をサポートしているかを確認してください。
第 50 条の違反に対する罰則
第 50 条への違反は、EU AI 法の第 99 条に基づいて罰せられます。段階的なペナルティ構造は次のように機能します。
禁止されている AI 慣行の使用: 最大 3,500 万ユーロまたは世界の年間売上高の 7%
第 50 条に基づく透明性の欠如を含むその他の違反: 最大 1,500 万ユーロまたは世界の年間売上高の 3%
当局への誤った情報の提供: 最大 750 万ユーロまたは 1%
いずれの場合も、2 つの金額のうち高い方が適用されます。執行は各加盟国の国家監督当局を通じて行われ、EU AI 事務局が調整します。中小企業に対しては、規制は当局に比例して罰金を適用するよう指示しています。数十億ユーロの収益を持つ世界的な AI プロバイダーにとって、1,500 万ユーロの下限ではなく、売上高の 3% が拘束力のある数字です。
AI ツールの出力が準拠していることを確認する方法
現在、AI プロバイダーが第 50 条のマーキング要件を満たしているかどうかを確認する最も簡単な方法は、ツールからの実際の出力を検査することです。検証フローは次のとおりです。
サンプル出力を生成します。 AI ツールを使用して、画像、音声ファイル、またはビデオを作成します。一部のエディタではメタデータが削除されるため、別のエディタで再保存せずにダウンロードします。
ファイルを C2PA Viewer にドロップします。すべての処理はクライアント側で行われます。このツールは、マニフェストが存在するかどうか、マニフェストに署名した人、署名の内容を表示します。
マニフェストが存在し、

予想される組織によって署名されました。 Adobe は Firefly 出力に署名する必要があり、OpenAI は DALL-E 3 と Sora 出力に署名する必要があり、Google は Imagen 出力に署名する必要があります。
マニフェストに AI の生成が示されていることを確認します。 「claim_generator」フィールドには AI ツールの名前を指定し、マニフェストには人間ではなく AI システムに起因する「c2pa.created」アクションを含める必要があります。
証明書が C2PA 信頼リストに載っていることを確認します。認識されない証明書によって署名された有効なマニフェストは、相互運用性の基準を満たしません。
マニフェストが欠落しているか、不正な形式であるか、検証できない証明書によって署名されている場合、その出力は機械可読 AI マーキングに関する第 50 条の基準を満たしていません。 C2PA ファイルを検証する方法については、ステップバイステップのウォークスルーを参照してください。
実践的な第 50 条遵守チェックリスト
2026 年 8 月 2 日の期限に向けて準備を進めている AI プロバイダーおよび導入担当者向け:
対象範囲内のシステムのインベントリを作成します。 EU ユーザーに届くコンテンツを生成または大幅に変更する、提供または展開するすべての AI システムを特定します。
現在のマーキング動作を監査します。サンプル出力を生成し、検査します。どの形式がすでに C2PA マニフェストを保持しており、どの形式が C2PA マニフェストを保持していないのかに注意してください。
計画の実施。マーキングが欠落している場合、スコープは C2PA 署名を出力パイプラインに追加するように機能します。最新の画像およびビデオ ツールチェーンのほとんどには、C2PA SDK バインディングが含まれています。
署名証明書を登録します。 C2PA 信頼リストによって認識されている CA から証明書を取得するか、適切な信頼リスト管理者に登録された自己署名証明書を使用します。
モデルカードまたはシステムカードに記録します。第 50 条の義務では、効果的かつ信頼性の高いマーキングが必要です。マーキングがどのように機能するかを説明するドキュメントは、コンプライアンスと監査への対応をサポートします。
検証パスを提供します。ユーザーはあなたのマーキングを確認できる必要があります。次のような公開検証者へのリンク

C2PA Viewer または独自の組み込み検証ツールはこれを満たします。
実践規範を追跡します。欧州委員会は、2026 年までの透明性に関する実践規範の最終決定を進めています。最終ガイダンスでは、規則の本文を超える具体的な義務が追加される可能性があります。
EU AI 法では C2PA が必要ですか?
いいえ、EU AI 法では C2PA を明示的に指定していません。第 50 条 (2) では、AI が生成したコンテンツを、効果的で相互運用性があり、堅牢で信頼性の高い機械可読形式でマークすることを義務付けています。 C2PA コンテンツ認証情報は、4 つの基準をすべて満たす主要な技術メカニズムであり、欧州委員会の透明性に関する実務規範では、準拠するマーキングの例として C2PA がリストされています。
EU AI 法第 50 条はいつ発効しますか?
第 50 条は、EU AI 法の発効から 2 年後の 2026 年 8 月 2 日に施行されます。この日付は、一般的な透明性義務と汎用 AI モデルのルールに適用されます。
第 50 条に従わない場合、どのような罰則がありますか?
EU AI 法第 99 条に基づき、透明性違反には、最大 1,500 万ユーロまたは前会計年度の世界年間総売上高の 3% のいずれか高い方の行政罰金が科される可能性があります。全国

[切り捨てられた]

## Original Extract

Article 50 of the EU AI Act requires AI-generated content to be marked by August 2, 2026. Learn what counts, who complies, and how C2PA satisfies it.

EU AI Act and C2PA: What Article 50 Requires for AI Content | C2PA Viewer Skip to main content C2PA Viewer Supported Devices & Services Articles Contact Share Feedback Home
Regulation • Updated April 26, 2026
EU AI Act and C2PA: What Article 50 Requires for AI Content
Deadline: Article 50 of the EU AI Act becomes enforceable on August 2, 2026. The European Commission Code of Practice on Transparency, which interprets the technical detail of the marking obligation, was still being finalized as of April 2026. This article reflects the Regulation text and the published draft Code as of that date.
The EU AI Act, formally Regulation (EU) 2024/1689, is the European Union's law governing artificial intelligence. Article 50 sets transparency obligations that take effect on August 2, 2026, and one of those obligations is specific to AI-generated content: outputs from generative AI systems must be marked in a machine-readable format detectable as artificially generated. The technical mechanism the European Commission's draft Code of Practice on Transparency names by example is C2PA Content Credentials . Article 50 takes effect on August 2, 2026. Penalties for non-compliance reach 15 million EUR or 3 percent of worldwide annual turnover, whichever is higher.
What Article 50 of the EU AI Act actually requires
Article 50 imposes transparency obligations on providers and deployers of certain AI systems. The clause that matters most for content provenance is Article 50(2): providers of generative AI systems must ensure that synthetic outputs (audio, image, video, or text) are marked in a machine-readable format and detectable as artificially generated or manipulated. The marking must be effective, interoperable, robust, and reliable, as far as is technically feasible.
Those four adjectives carry specific meaning under the Code of Practice on Transparency that the European Commission has been developing alongside the Regulation:
Effective : the marking must actually identify the content as AI-generated to a verifying party
Interoperable : any compliant verifier must be able to read the marking, not only the provider's own tool
Robust : the marking should resist common transformations such as format conversion or minor edits
Reliable : the marking should be tamper-evident, so that a verifier can detect whether it has been altered or forged
Article 50 also covers other transparency duties: deployers of emotion recognition systems must inform users, deployers generating deepfakes must disclose them, and providers of chatbot-style AI must make the artificial nature of the system clear to the person interacting with it. The marking obligation in 50(2) is the one with direct technical consequences for content credentials.
Does the EU AI Act mandate C2PA specifically?
No. The EU AI Act is technology-neutral and does not name C2PA in its operative articles. The Regulation describes the marking standard in functional terms (machine-readable, effective, interoperable, robust, reliable) and leaves the technical implementation to standards bodies and to the Commission's Code of Practice.
In practice, C2PA Content Credentials are the marking technology favored by the Commission's draft Code of Practice on Transparency. The Code lists C2PA as an example of a technical solution that satisfies all four criteria, alongside complementary signals like Google's SynthID watermarking. C2PA is an open standard already deployed by Adobe, OpenAI, and Google, with cryptographic signatures that make tampering detectable.
In short: the law does not require C2PA, but the regulatory ecosystem points firmly at it. Providers who adopt C2PA match the most concrete official guidance available on what compliant marking looks like.
Why C2PA, not just any watermark
Plain pixel watermarks (visible logos or invisible perceptual marks) are not interoperable in the Article 50 sense. Each provider would invent its own watermark, no third party could verify a competitor's mark, and there is no cryptographic chain of trust. C2PA solves this by defining a shared file format (JUMBF), a shared claim schema (JSON-LD), and a shared signature format (COSE). One verifier can read all of them. See what is inside a C2PA manifest for the technical detail.
Who must comply with Article 50?
Article 50 reaches two categories of organizations: providers of generative AI systems (the companies that build and offer the model) and deployers (companies that use the AI to generate content shown to people). Both bear obligations, but the marking duty in 50(2) sits with providers.
Geographic scope is broad. Article 2 of the Regulation places any provider in scope if it puts an AI system on the EU market, if its outputs are used in the EU, or if its users are located in the EU. A US-based provider like OpenAI, an Israeli image-generation startup, or a Japanese camera manufacturer is in scope as soon as European users encounter their content.
There are limited exceptions. Article 50(2) does not apply when the AI performs an assistive editing function that does not substantially alter the input, or when the use is authorized by law for the purpose of detecting, preventing, investigating, or prosecuting criminal offenses. Standard generative use cases (an AI tool producing an image, a text-to-video system, a voice synthesis service) fall under the obligation.
When does Article 50 become enforceable?
Article 50 becomes enforceable on August 2, 2026, two years after the EU AI Act entered into force on August 1, 2024. The same date applies to the general transparency rules and to obligations on general-purpose AI models. The Regulation has a staggered timeline overall:
February 2, 2025 : Prohibitions on banned practices took effect
August 2, 2025 : General-purpose AI model provisions began applying for new models
August 2, 2026 : General transparency obligations including Article 50 become enforceable
August 2, 2027 : High-risk AI system rules take effect
How C2PA Content Credentials satisfy Article 50
C2PA embeds a cryptographically signed manifest into image, video, audio, and document files. The manifest declares which AI system generated or modified the content, when the content was produced, and which organization signed the claim. Any C2PA-compliant verifier (including C2PA Viewer) can read this manifest and confirm the signature against a published trust list.
Mapping this to the four Article 50 criteria:
Effective : a valid C2PA manifest containing a `c2pa.created` action signed by an AI provider unambiguously marks the content as AI-generated
Interoperable : C2PA is an open standard maintained by the Coalition for Content Provenance and Authenticity. Any compliant tool reads any compliant manifest
Robust : manifests survive most format conversions and re-encodings handled by C2PA-aware tools. Stripping is still possible through naive re-saves, which is why the Code of Practice encourages pairing C2PA with invisible watermarks
Reliable : the manifest is signed using COSE with a certificate chain verifiable against the C2PA Trust List. Tampering breaks the signature
Several major AI providers have already deployed C2PA in production. Adobe Firefly, OpenAI DALL-E 3, OpenAI Sora, and Google Imagen all embed manifests in their outputs. Midjourney is the most prominent generative AI tool that does not, and it faces direct regulatory exposure under Article 50 unless its approach changes before August 2026. See which AI tools support C2PA for the current status of each major platform.
Penalties for non-compliance with Article 50
Non-compliance with Article 50 is sanctioned under Article 99 of the EU AI Act. The graduated penalty structure works as follows:
Use of banned AI practices: up to 35 million EUR or 7 percent of global annual turnover
Other violations including transparency failures under Article 50: up to 15 million EUR or 3 percent of global annual turnover
Supplying incorrect information to authorities: up to 7.5 million EUR or 1 percent
In each case the higher of the two amounts applies. Enforcement runs through national supervisory authorities in each Member State, coordinated by the EU AI Office. For SMEs, the Regulation directs authorities to apply fines proportionately. For a global AI provider with multi-billion-euro revenue, 3 percent of turnover is the binding number, not the 15 million EUR floor.
How to verify your AI tool's outputs are compliant
The fastest way to check whether an AI provider satisfies Article 50's marking requirement today is to inspect a real output from the tool. The verification flow is:
Generate a sample output. Use the AI tool to produce an image, audio file, or video. Download it without re-saving through another editor, since some editors strip metadata.
Drop the file into C2PA Viewer . All processing happens client-side. The tool will display whether a manifest is present, who signed it, and what the signature claims.
Confirm a manifest is present and signed by the expected organization. Adobe should sign Firefly outputs, OpenAI should sign DALL-E 3 and Sora outputs, Google should sign Imagen outputs.
Confirm the manifest indicates AI generation. The `claim_generator` field should name the AI tool, and the manifest should include a `c2pa.created` action attributed to the AI system rather than to a human.
Confirm the certificate is on the C2PA Trust List. A valid manifest signed by an unrecognized certificate would not satisfy the interoperability criterion.
If a manifest is missing, malformed, or signed by an unverifiable certificate, the output does not meet Article 50's standard for machine-readable AI marking. See how to verify a C2PA file for a step-by-step walk-through.
Practical Article 50 compliance checklist
For AI providers and deployers preparing for the August 2, 2026 deadline:
Inventory in-scope systems. Identify every AI system you provide or deploy that generates or substantially modifies content reaching EU users.
Audit current marking behavior. Generate sample outputs and inspect them. Note which formats already carry C2PA manifests and which do not.
Plan implementation. Where marking is missing, scope work to add C2PA signing to the output pipeline. Most modern image and video toolchains have C2PA SDK bindings.
Register a signing certificate. Obtain a certificate from a CA recognized by the C2PA Trust List, or use a self-signed certificate registered with the appropriate trust list maintainer.
Document in your model card or system card. The Article 50 obligation requires effective and reliable marking. Documentation that explains how your marking works supports compliance and audit responses.
Provide a verification path. Users should be able to verify your marking. Linking to a public verifier such as C2PA Viewer or to your own embedded verifier satisfies this.
Track the Code of Practice. The European Commission is finalizing the Code of Practice on Transparency through 2026. Final guidance may add concrete obligations beyond the Regulation text.
Does the EU AI Act require C2PA?
No, the EU AI Act does not name C2PA explicitly. Article 50(2) requires AI-generated content to be marked in a machine-readable format that is effective, interoperable, robust, and reliable. C2PA Content Credentials are the leading technical mechanism that satisfies all four criteria, and the European Commission Code of Practice on Transparency lists C2PA as an example of compliant marking.
When does Article 50 of the EU AI Act take effect?
Article 50 becomes enforceable on August 2, 2026, two years after the EU AI Act entered into force. This date applies to general transparency obligations and to general-purpose AI model rules.
What are the penalties for not complying with Article 50?
Under Article 99 of the EU AI Act, transparency violations can be sanctioned with administrative fines up to 15 million EUR or 3 percent of total worldwide annual turnover for the preceding financial year, whichever is higher. National

[truncated]
