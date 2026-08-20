---
source: "https://techstrong.ai/articles/openai-unveils-zero-data-retention-for-frontier-models-previews-privacy-preserving-safety-system/"
hn_url: "https://news.ycombinator.com/item?id=49381251"
title: "OpenAI Unveils Zero Data Retention for Frontier Models"
article_title: "OpenAI Unveils Zero Data Retention for Frontier Models, Previews Privacy-Preserving Safety System - Techstrong.ai"
image: "https://techstrong.ai/wp-content/uploads/2026/08/openai_zdr_psp_770x330.jpg"
author: "CrankyBear"
captured_at: "2026-08-20T23:16:52Z"
capture_tool: "hn-digest"
hn_id: 49381251
score: 1
comments: 0
posted_at: "2026-08-20T22:41:59Z"
tags:
  - hacker-news
  - translated
---

# OpenAI Unveils Zero Data Retention for Frontier Models

- HN: [49381251](https://news.ycombinator.com/item?id=49381251)
- Source: [techstrong.ai](https://techstrong.ai/articles/openai-unveils-zero-data-retention-for-frontier-models-previews-privacy-preserving-safety-system/)
- Score: 1
- Comments: 0
- Posted: 2026-08-20T22:41:59Z

## Translation

タイトル: OpenAI、フロンティアモデルのデータ保持ゼロを発表
記事のタイトル: OpenAI がフロンティア モデルのゼロ データ保持を発表、プライバシー保護安全システムをプレビュー - Techstrong.ai
説明: OpenAI は、対象となる API 顧客にフロンティア モデルのデプロイメント向けにゼロ データ保持を提供し、新しい安全メカニズムを導入します。

記事本文:
OpenAI、フロンティアモデルのゼロデータ保持を発表、プライバシー保護安全システムをプレビュー - Techstrong.ai
コンテンツにスキップ
トグルナビゲーション 最新の記事
OpenAI、フロンティアモデルのゼロデータ保持を発表、プライバシー保護安全システムをプレビュー
OpenAI は、フロンティア モデルを使用して、対象となる API 顧客向けにゼロ データ保持アクセスを拡張しています。
ZDR は、対象となるプロンプトと出力が処理後に保持されるのを防ぎ、悪用監視ログに記録されないようにします。
OpenAI は、スタッフに基礎となる顧客コンテンツへのアクセスを許可せずに、有害な行動パターンを検出するプライベート セーフティ処理をプレビューしています。
OpenAI は、対象となる API 顧客にフロンティア モデル展開向けのゼロ データ保持を提供し、基礎となる顧客コンテンツを公開することなく誤用パターンを特定するように設計された新しい安全メカニズムを導入します。
OpenAI の新しいゼロ データ保持 (ZDR) は、人々の健康データや財務記録、機密の事業計画や独自の研究などの機密資料を扱う組織を対象としています。 ZDR を使用すると、OpenAI はリクエストの処理後にプロンプ​​トと出力を保持しません。このプログラムに基づき、OpenAI は、顧客のコンテンツを従業員がレビューのために利用することはできず、顧客が明示的にオプトインしない限り、エンタープライズ API データはモデルのトレーニングに使用されないと述べました。
ZDR は、要求可能なエンタープライズ/API の取り決めです。 Joe と Jane ChatGPT ユーザーは利用できません。
ZDR を導入するユーザーに対して、OpenAI は、顧客のコンテンツが顧客が管理するインフラストラクチャ上に留まるようにすることを約束します。同社は、コンテンツが顧客が管理するキーで暗号化された OpenAI インフラストラクチャ上に存在する代替案も開発中です。 OpenAIは、同社職員はこれらの鍵のコピーを所有しないだろうと述べた。
興奮しすぎる前に、

保証には多くの注意事項が伴います。たとえば、OpenAI は「特定の顧客に対して、モデルをゼロ データ保持の対象外にする権利」を留保します。さらに、データは悪用監視ログ用に保存される場合があり、一部の API 機能からのアプリケーション状態データはタスクまたはリクエストを満たすために保持されます。これには、ベクトル ストア データ、スレッド、またはその他のサーバー側オブジェクトが含まれます。
もちろん、ZDR は、OpenAI が関連資料を一切保持しないという絶対的な約束をするものではありません。 OpenAIは、明らかな児童性的虐待資料（CSAM）を報告する法的義務が引き続きあると述べた。潜在的な CSAM としてフラグが付けられた画像は、手動によるレビューとレポートのために引き続き保持されます。
では、ZDR はデータをどのように処理するのでしょうか? ZDR を有効にすると、OpenAI は対象となるプロンプト、応答、またはその他の顧客コンテンツを不正行為監視ログに保持しなくなります。 Responses API および Chat Completions API の場合、アプリケーションが store: true を送信した場合でも、store は強制的に false として扱われます。個別の「データでトレーニングしない」設定は提供されません。これは、API データがデフォルトでトレーニングから明確に除外されているためです。
OpenAI は、これらすべてがどのように機能するかを正確に説明したホワイトペーパーをリリースしていません。それでも、これが承認された API データ制御設定として機能することはわかっています。
さて、基礎となるインタラクションを保持または検査できない場合に、AI プロバイダーが危険な使用または不正使用をどのように監視できるかは、非常に良い問題です。 OpenAIは、既存のZDR互換の安全対策は一般的に各インタラクションを個別に評価すると述べた。それは良いことですが、ZDR は、エージェント ワークフローの一連のプロンプト、アカウント、またはステップ全体でのみ可視化されるリスクを見逃す可能性があります。実際に、これを言い換えてみましょう。 ZDR が何らかのリスクを見逃すかどうかは問題ではありません。 ZDR は高度なリスクを検出しません。
この問題は OpenAI だけではありません。すべての

他の AI 企業も、セキュリティ、プライバシー、悪用防止の間でバランスを取るという課題に直面しています。
また、「トレーニングに使用しない」と「保持率ゼロ」は異なる約束であることに留意することも重要です。 AI プロバイダーは、安全性レビュー、デバッグ、または法令順守のために、プロンプトと出力を一定期間保持しながら、API データでトレーニングしないことを約束する場合があります。さらに、ZDR は推論ペイロードの永続的なプロバイダー側​​ストレージに対処しますが、それでも一時的なメモリ内処理、自動不正行為スクリーニング、顧客構成の永続化、または法的に強制された保存が可能になります。
ZDRの導入に加えて、OpenAIは「Private Safety Processing (PSP)」のプレビューも導入しました。 PSP は、関連するインタラクション全体のパターンを検査し、疑わしい有害なアクティビティに関する狭い範囲のシグナルを返すように設計されています。 OpenAI のピンキーは、これらのアクションを生成したプロンプトや応答への OpenAI スタッフのアクセスを提供することなく、これを実現すると約束しています。
OpenAI と他のすべての AI ベンダーが ZDR と PSP を提供しているのは、企業や規制産業が AI フロンティア モデルの安全管理とデータ最小化の義務を調和させる必要があるためです。最近の強力な AI システムの導入では、安全性を監視するために機密コンテンツの保持を有効にすることがお客様に求められています。これは、コンプライアンス義務や内部セキュリティ ポリシーと衝突する可能性があるトレードオフです。
ZDR と PSP では、顧客データに問題が発生した場合、または AI が危険または違法な結果を生成した場合、AI ユーザーは責任を負わなければなりません。たとえば、誰かが病院から患者の医療記録を抽出する方法を見つけた場合です。
最後に、PSP は今のところプレビューにすぎません。実装方法、潜在的な問題をどのように検出するか、誤検知をどのように処理するか

答えのない質問はすべて。これらの質問に対する答えについては、2026 年 9 月に発行される技術ホワイト ペーパーをお待ちください。
LLM 用語のチートシート: AI 実務者向けの包括的なリファレンス
AI を活用した企業に対処するためにデータ保持戦略がどのように進化したか
米国司法長官、AIが子供に及ぼす悪影響を研究するよう議会に要請
自律型 AI エージェントを本番環境に安全に導入する方法
最新のソフトウェア開発と配信

## Original Extract

OpenAI will offer eligible API customers Zero Data Retention for frontier-model deployments and introduce a new safety mechanism.

OpenAI Unveils Zero Data Retention for Frontier Models, Previews Privacy-Preserving Safety System - Techstrong.ai
Skip to content
Toggle Navigation Latest Articles
OpenAI Unveils Zero Data Retention for Frontier Models, Previews Privacy-Preserving Safety System
OpenAI is expanding Zero Data Retention access for eligible API customers using frontier models.
ZDR prevents eligible prompts and outputs from being retained after processing and keeps them out of abuse-monitoring logs.
OpenAI is previewing Private Safety Processing to detect harmful behavior patterns without giving staff access to underlying customer content.
OpenAI will offer eligible API customers Zero Data Retention for frontier-model deployments and introduce a new safety mechanism designed to identify misuse patterns without exposing underlying customer content.
OpenAI’s new Zero Data Retention (ZDR) is aimed at organizations handling sensitive material, such as people’s health data and financial records, and confidential business plans and proprietary research. With ZDR, OpenAI will not retain prompts and outputs after a request is processed. Under the program, OpenAI said customer content will not be available to its personnel for review , and enterprise API data will not be used to train models unless a customer explicitly opts in.
ZDR is a requestable enterprise/API arrangement. It’s not available for Joe and Jane ChatGPT users.
For those who deploy ZDR, OpenAI promises that customer content stays on infrastructure controlled by the customer . The company is also developing an alternative in which content resides on OpenAI infrastructure encrypted with keys controlled by the customer; OpenAI said its personnel would not possess copies of those keys.
Before you get too excited, those guarantees come with numerous caveats. For example, OpenAI reserves “the right to make models ineligible for Zero Data Retention … for specific customers.” In addition, data may be stored for abuse-monitoring logs, and application state data from some API features will be kept to fulfill the task or request. This includes vector-store data, threads, or other server-side objects.
Of course, ZDR is not an absolute promise that OpenAI will never retain any related material. OpenAI said it remains legally obligated to report apparent child sexual abuse material (CSAM). Images flagged as potential CSAM will continue to be retained for manual review and reporting.
So, what does ZDR do with your data? With ZDR enabled, OpenAI won’t retain eligible prompts, responses, or other customer content in abuse-monitoring logs. For the Responses API and Chat Completions API, store is forcibly treated as false , even if an application sends store: true . It doesn’t provide a separate “do not train on my data” setting. That’s because API data is specifically excluded from training by default.
OpenAI hasn’t released a white paper on exactly how all this works. Still, we do know that it works as an approved API data-control setting.
Now, how an AI provider can monitor dangerous or abusive use when it cannot retain or inspect underlying interactions is a darn good question. OpenAI said existing ZDR-compatible safeguards generally assess each interaction independently. That’s nice, but ZDR may miss risks that become visible only across a sequence of prompts, accounts, or steps in an agentic workflow. Actually, let me rephrase this. It’s not a question of whether ZDR may miss some risks; ZDR won’t find sophisticated risks.
OpenAI isn’t alone in this problem. All the other AI companies face this balancing act between security, privacy, and preventing abuse.
It’s also important to keep in mind that “Not used for training” and “zero retention” are different promises. An AI provider may promise not to train on API data while still retaining prompts and outputs for a defined period for safety review, debugging, or legal compliance. Besides, while ZDR addresses durable provider-side storage of the inference payload, that still allows transient in-memory processing, automated abuse screening, customer-configured persistence, or legally compelled retention.
In addition to introducing ZDR, OpenAI also introduced a preview of “Private Safety Processing (PSP).” PSP is designed to examine patterns across related interactions and return narrowly scoped signals about suspected harmful activity. It will do this, OpenAI pinky promises, without providing OpenAI staff access to the prompts or responses that produced those actions.
OpenAI and all the other AI vendors are offering ZDR and PSP because enterprises and regulated industries need to reconcile AI frontier-model safety controls with data-minimization duties. Some recent deployments of powerful AI systems have required customers to enable retention of sensitive content for safety monitoring. This is a tradeoff that can clash with compliance obligations and internal security policies.
With ZDR and PSP, AI users will have to take responsibility for when something goes awry with customer data or the AI produces dangerous or illegal results. For example, if someone finds a way to extract a patient’s medical records from a hospital.
Finally, the PSP is only a preview for now. How it will be implemented, how it will detect potential problems and how it will handle false positives are all unanswered questions. For answers to those questions, stay tuned for a technical white paper that should be appearing in September 2026.
LLM Terminology Cheat Sheet: Comprehensive Reference for AI Practitioners
How Data Retention Strategies Have Evolved to Address the AI-Powered Enterprise
US Attorney General Urges Congress to Study AI’s Harmful Effects On Children
How to Safely Deploy Autonomous AI Agents in Production
Modern Software Development and Delivery
