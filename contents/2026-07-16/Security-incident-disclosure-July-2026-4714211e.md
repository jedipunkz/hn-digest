---
source: "https://huggingface.co/blog/security-incident-july-2026"
hn_url: "https://news.ycombinator.com/item?id=48935615"
title: "Security incident disclosure – July 2026"
article_title: "Security incident disclosure — July 2026"
author: "sbulaev"
captured_at: "2026-07-16T15:26:58Z"
capture_tool: "hn-digest"
hn_id: 48935615
score: 2
comments: 0
posted_at: "2026-07-16T15:07:06Z"
tags:
  - hacker-news
  - translated
---

# Security incident disclosure – July 2026

- HN: [48935615](https://news.ycombinator.com/item?id=48935615)
- Source: [huggingface.co](https://huggingface.co/blog/security-incident-july-2026)
- Score: 2
- Comments: 0
- Posted: 2026-07-16T15:07:06Z

## Translation

タイトル: セキュリティ インシデントの開示 – 2026 年 7 月
記事のタイトル: セキュリティ インシデントの開示 — 2026 年 7 月
説明: 私たちは、オープンソースとオープン サイエンスを通じて人工知能を進歩させ、民主化する旅の途中にあります。

記事本文:
セキュリティインシデントの開示 — 2026 年 7 月
ハグ顔モデル
記事に戻る a]:hidden">
セキュリティインシデントの開示 — 2026 年 7 月
AI による侵入の分析 非対称問題
当社は、限られた内部データセットと当社のサービスで使用されるいくつかの認証情報への不正アクセスを特定しました。パートナーまたは顧客データが影響を受けたかどうかの評価はまだ完了しており、必要に応じて影響を受ける当事者に直接連絡します。公開されているユーザー向けモデル、データセット、またはスペースが改ざんされた証拠は見つかっておらず、ソフトウェア サプライ チェーン (コンテナー イメージと公開されたパッケージ) はクリーンであることが確認されました。
侵入は、AI プラットフォームが独自に公開されている場所、つまりデータ処理パイプラインから始まりました。悪意のあるデータセットは、データセット処理で 2 つのコード実行パス (リモート コード データセット ローダーとデータセット構成内のテンプレート インジェクション) を悪用し、処理ワーカーでコードを実行しました。そこから、攻撃者はノードレベルのアクセスにエスカレートし、クラウドとクラスターの認証情報を収集し、週末にかけて複数の内部クラスターに横方向に移動しました。
このキャンペーンは、自律エージェント フレームワーク (エージェントのセキュリティ研究ハーネスに基づいて構築されているようですが、使用されている LLM はまだ不明) によって実行され、公共サービス上で実行される自己移行型のコマンド アンド コントロールを使用して、短期間のサンドボックスの群れ全体で何千もの個別のアクションを実行しました。これは、業界が予測してきた「エージェント攻撃者」のシナリオと一致します。
ルートの脆弱性を修正しました。初期アクセスに使用されるデータセットのコード実行パスが閉じられます。
影響を受けたクラスター全体で攻撃者の足場を根絶し、侵害されたノードを再構築しました。
影響を受けた資格情報とトークンを取り消してローテーションし、より広範な予防的ローテーションを開始しました。

エクレット。
追加のガードレールとより厳格なアドミッション制御をクラスターに導入しました。
検出とアラートが改善され、曜日を問わず、重大度の高い信号が数分でレスポンダに通知されるようになりました。
私たちは外部のサイバーセキュリティ法医学専門家と協力して問題を調査し、セキュリティ ポリシーと手順を見直しています。最後に、私たちはこの事件を法執行機関にも報告しました。
予防策として、アクセス トークンをローテーションし、アカウントの最近のアクティビティを確認することをお勧めします。影響を受けていると思われる場合、またはセキュリティ上の懸念を報告したい場合は、 security@huggingface.co までご連絡ください。
24時間体制で対応してくれたHugging Faceの各チームに感謝するとともに、これにより混乱が生じたことをお詫び申し上げます。セキュリティに終わりはありません。私たちは基準を引き上げ続けます。
AI による侵入を分析する
この攻撃は当初、AI 支援検出によって表面化されました。当社の異常検出パイプラインは、セキュリティ テレメトリに対する LLM ベースのトリアージを使用して、実際の信号を日常のノイズから分離します。侵害の兆候を示したのは、これらの信号の相関関係でした。
数万の自動アクションの群れが何をしたかを理解するために、17,000 を超える記録されたイベントで構成される完全な攻撃者のアクション ログに対して LLM 主導の分析エージェントを実行しました。これにより、タイムラインを再構築し、侵害の痕跡を抽出し、接触した認証情報をマッピングし、おとり活動からの真の影響を分離することができました。このアプローチのおかげで、通常は数日かかる作業を数時間で実行し、敵のスピードに匹敵することができました。
この分析に使用できるモデルの選択には、予想外の制約がありました。これについては以下で説明します。
ログ分析を開始したとき、私たちは最初に商用 API の背後にあるフロンティア モデルを使用しました。これは問題ありませんでした

rk: 分析には大量の実際の攻撃コマンド、エクスプロイト ペイロード、C2 アーティファクトの送信が必要ですが、これらのリクエストはプロバイダーの安全ガードレールによってブロックされ、インシデント対応者と攻撃者を区別できませんでした。代わりに、独自のインフラストラクチャ上のオープンウェイト モデルである GLM 5.2 でフォレンジック分析を実行しました。これには 2 番目の利点がありました。攻撃者のデータも、攻撃者が参照した資格情報も環境から流出しませんでした。
この経験は、計画する価値のあるギャップを示しています。ジェイルブレイクされたホスト型モデルか、制限のない無差別級モデルか、どのモデルが攻撃者のエージェントに動力を与えていたのかはわかりません。いずれにせよ、攻撃者は使用ポリシーに拘束されませんでしたが、私たち自身のフォレンジック作業は、最初に試したホストされたモデルのガードレールによってブロックされました。防御側への実践的な教訓: ガードレールのロックアウトを回避し、攻撃者のデータと資格情報が環境から流出するのを防ぐために、独自のインフラストラクチャで実行できる有能なモデルを精査し、インシデントが発生する前に準備しておきます。これはホスト型モデルの安全対策に反対するものではなく、このフィードバックを関係するプロバイダーと共有しています。
自律型の AI 主導の攻撃ツールは、もはや理論上のものではありません。これにより、広範で忍耐強い複数段階のキャンペーンを実行するコストが削減され、マシンの速度で動作します。現在、オンライン プラットフォームを防御するということは、データとモデルの表面を第一級の攻撃対象領域として扱い、防御に AI を使用して歩調を合わせることを意味します。私たちはそこに投資し続け、学んだことを共有し続けます。
画像、音声、ビデオをアップロードするには、テキスト入力をドラッグするか、貼り付けるか、ここをクリックします。ここをタップまたは貼り付けて画像をアップロード コメント · コメントするにはサインアップまたはログインしてください

## Original Extract

We’re on a journey to advance and democratize artificial intelligence through open source and open science.

Security incident disclosure — July 2026
Hugging Face Models
Back to Articles a]:hidden">
Security incident disclosure — July 2026
Analyzing an AI-driven intrusion The asymmetry problem
We identified unauthorized access to a limited set of internal datasets and to several credentials used by our services. We are still completing our assessment of whether any partner or customer data was affected, and we will contact any affected parties directly as required. We have found no evidence of tampering with public, user-facing models, datasets, or Spaces, and our software supply chain (container images and published packages) was verified clean.
The intrusion started where AI platforms are uniquely exposed: the data-processing pipeline. A malicious dataset abused two code-execution paths in our dataset processing (a remote-code dataset loader and a template-injection in a dataset configuration) to run code on a processing worker. From there, the actor escalated to node-level access, harvested cloud and cluster credentials, and moved laterally into several internal clusters over a weekend.
The campaign was run by an autonomous agent framework (appearing to be built on an agentic security-research harness - used LLM still not known) executing many thousands of individual actions across a swarm of short-lived sandboxes, with self-migrating command-and-control staged on public services. This matches the "agentic attacker" scenario the industry has been forecasting.
Fixed the root vulnerability: the dataset code-execution paths used for initial access are closed.
Eradicated the attacker's foothold across the affected clusters and rebuilt the compromised nodes.
Revoked and rotated the affected credentials and tokens, and began a broader precautionary rotation of secrets.
Deployed additional guardrails and stricter admission controls on our clusters.
Improved our detection and alerting so a high-severity signal pages a responder in minutes, any day of the week.
We are working with outside cybersecurity forensic specialists to investigate the issue and review our security policies and procedures. Finally, we have also reported this incident to law enforcement agencies.
As a precaution, we recommend rotating any access tokens and reviewing recent activity on your account. If you believe you are affected, or want to report a security concern, contact us at security@huggingface.co .
We are grateful to the teams across Hugging Face who responded around the clock, and we are sorry for any disruption this caused. Security is never finished; we will keep raising the bar.
Analyzing an AI-driven intrusion
The attack was initially surfaced through AI-assisted detection. Our anomaly-detection pipeline uses LLM-based triage over security telemetry to separate real signals from the daily noise, and it was the correlation of those signals that flagged the compromise.
To understand what a swarm of tens of thousands of automated actions did, we ran LLM-driven analysis agents over the full attacker action log, comprised of more than 17,000 recorded events. This allowed us to reconstruct the timeline, extract indicators of compromise, map the credentials touched, and separate genuine impact from decoy activity. Thanks to this approach, we were able to do in hours what would usually take days, and match the adversary's speed.
The choice of models we could use for this analysis was constrained in a way we did not anticipate; we describe this below.
When we started the log analysis, we first used frontier models behind commercial APIs. This did not work: the analysis requires submitting large volumes of real attack commands, exploit payloads, and C2 artifacts, and these requests were blocked by the providers' safety guardrails, which cannot distinguish an incident responder from an attacker. We ran the forensic analysis instead on GLM 5.2, an open-weight model, on our own infrastructure. This had a second benefit: no attacker data, and none of the credentials it referenced, left our environment.
This experience points to a gap worth planning for. We do not know which model powered the attacker's agents, whether a jailbroken hosted model or an unrestricted open-weight one; either way, the attacker was bound by no usage policy, while our own forensic work was blocked by the guardrails of the hosted models we first tried. The practical lesson for defenders: have a capable model you can run on your own infrastructure vetted and ready before an incident, both to avoid guardrail lockout and to keep attacker data and credentials from leaving your environment. This is not an argument against safety measures on hosted models, and we are sharing this feedback with the providers concerned.
Autonomous, AI-driven offensive tooling is no longer theoretical. It lowers the cost of running a broad, patient, multi-stage campaign, and it operates at machine speed. Defending an online platform now means treating the data and model surface as a first-class attack surface, and using AI on defense to keep pace. We will keep investing there, and keep sharing what we learn.
Upload images, audio, and videos by dragging in the text input, pasting, or clicking here . Tap or paste here to upload images Comment · Sign up or log in to comment
