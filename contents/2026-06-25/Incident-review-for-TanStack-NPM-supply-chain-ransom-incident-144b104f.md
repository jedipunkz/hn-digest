---
source: "https://grafana.com/blog/post-incident-review-for-tanstack-npm-supply-chain-ransom-incident/"
hn_url: "https://news.ycombinator.com/item?id=48675936"
title: "Incident review for TanStack NPM supply chain ransom incident"
article_title: "Post-incident review for TanStack npm supply chain ransom incident | Grafana Labs"
author: "dboreham"
captured_at: "2026-06-25T16:45:06Z"
capture_tool: "hn-digest"
hn_id: 48675936
score: 1
comments: 0
posted_at: "2026-06-25T16:37:33Z"
tags:
  - hacker-news
  - translated
---

# Incident review for TanStack NPM supply chain ransom incident

- HN: [48675936](https://news.ycombinator.com/item?id=48675936)
- Source: [grafana.com](https://grafana.com/blog/post-incident-review-for-tanstack-npm-supply-chain-ransom-incident/)
- Score: 1
- Comments: 0
- Posted: 2026-06-25T16:37:33Z

## Translation

タイトル: TanStack NPM サプライ チェーン身代金事件のインシデント レビュー
記事のタイトル: TanStack npm サプライ チェーン身代金事件の事後レビュー |グラファナ研究所
説明: 最近の TanStack サプライ チェーン身代金事件の内部調査を完了し、最初の調査結果を確認しました: この事件は厳密に Grafana Labs GitHub 環境に限定されていました。顧客の実稼働システムへの不正アクセスはなく、Grafana Cloud プラットフォームは
[切り捨てられた]

記事本文:
TanStack npm サプライ チェーン身代金事件の事後レビュー | Grafana Labs の新機能: Grafana 13 リリース、最新の AI、OSS プロジェクトの更新など、GrafanaCON2026 から
JP ダウンロード お問い合わせ Grafana Cloud
サイト検索 Ask AI ブログ TanStack npm サプライ チェーン身代金事件の事後レビュー: 顧客の実稼働システムへの不正アクセスはありません
5 月 27 日、私たちは最近の TanStack サプライ チェーン身代金事件の内部調査を完了し、最初の調査結果を確認しました。つまり、この事件は厳密に Grafana Labs の GitHub 環境に限定されていたということです。顧客の実稼働システムへの不正アクセスはなく、Grafana Cloud プラットフォームは影響を受けませんでした。
追加の独立した監査のために、サイバーセキュリティとインシデント対応のリーダーである Mandiant と協力しました。私たちは調査のために Grafana Labs のログ環境への API アクセスを提供し、調査のためにシステム全体でクエリを実行しました。マンディアント氏は、「公的機関やエンド ユーザーに提供された運用リポジトリ内でコードの改ざんやリポジトリ ポイズニングの証拠は存在しなかった」ことを確認しました。
インシデントを発見して以来、Grafana Labs のセキュリティ チームは、調査の完了とセキュリティ運用の強化という 2 つのワークストリームを並行して実行してきました。インシデント対応と修復の取り組みについてさらに詳しく共有するために、透明性の精神に基づいてこのブログを公開しています。
以前のアップデートを読む代わりに短いバージョンをお探しの場合は、TL;DR: TanStack サプライ チェーン攻撃は、5 月 11 日に Mini Shai-Hulud キャンペーンを通じて私たちを襲いました。当時、私たちはこのインシデントに関係するすべての認証情報を正常にローテーションしたと信じていました。一つ逃してしまいました。私はこの見落としを傲慢のせいにするつもりはありません。そのこと

当時私たちが持っていたのは、ローテーションが徹底的であると信じ込ませただけでした。私たちは間違っていました。
悪意のある攻撃者は、その見落とされた認証情報を利用して、リポジトリ コレクション全体のクローンを作成しました。その後、5 月 16 日に連絡を取り、コード漏洩を防ぐために身代金を要求しました。
Grafana Labs はオープンソース企業であるため、なぜこれが懸念されるのか疑問に思われるかもしれません。ソース コードのほとんどは公開されていますが、内部ツールや特定の Grafana Cloud 機能などについてはプライベート リポジトリを維持しています。これは重い決断でしたが、私たちは原則と FBI の文書化されたガイダンスに従い、支払いをしませんでした。
私たちは直ちに軽減策を開始し、顧客の実稼働システムへの不正アクセスがなく、Grafana Cloud プラットフォームが影響を受けていないことを確認しました。また、コードベースのダウンロード中に変更されていないことも確認しました。当社の顧客とオープンソース ユーザーは何もする必要はありません。
私たちは土曜日にこの事件について警告を受け、全社のチームが迅速かつ断固たる行動をとりました。 (あるいは、私の大好きなラッパー、ビッグ・ダディ・ケインの言葉を借りれば、グラファナ・ラボでは半歩も踏み出せません。)
これに応じて、Grafana Labs は 5 月 17 日にすべての GitHub アプリケーションを一時停止し、5 月 18 日に世界的なコード凍結を開始し、Vault、GitHub、Okta、Kubernetes、AWS、GCP、およびホストのログのクロスプラットフォーム監査を実施して、本番環境の顧客データが侵害されていないことを確認しました。
その後数週間、当社のエンジニアリング チームは以下を含む包括的な監査に貢献しました。
セキュリティに重点を置いた 1,500 件の PR レビューを完了
280 の GitHub アプリケーションを監査し、権限を剥奪し、いくつかを削除
1,200 のリポジトリをスキャンして改ざんの兆候がないか確認する
2,300 件の PR レビューを実行し、単一の重要な変更に対する不正な変更を探します

リポジトリ
インフラストラクチャ監査を終了し、レガシー システムを廃止する
広範囲にわたる新規アクセス監査の実施
それは大規模な仕事でしたが、各チームは自分たちの役割を果たすために並外れた方法でステップアップしました。エンジニアリング、セキュリティ、部門横断的なパートナーは精力的に対応し、私がここ Grafana Labs で常に大切にしてきた、コミュニティと顧客に対する協力と共通の取り組みを実証しました。
初期評価の結果、ダウンロードされたコンテンツには、ソース コードに加えて、一部の Grafana Labs チームが内部運用情報や当社のビジネスに関するその他の詳細を共同作業および保存するために使用する GitHub リポジトリが含まれていることがわかりました。これには、たとえば、専門的な環境で交換されるビジネス上の連絡先の名前や電子メール アドレス、および過去のマーケティング キャンペーンで使用された電子メール アドレスが含まれます。これは、運用システムや Grafana クラウド プラットフォームの使用を通じて取得または処理された情報ではありません。
あなたのドメインの電子メール アドレスが特定されたかどうかを知りたい場合は、Grafana Labs サポートにお問い合わせください。
5 月 11 日 19:21 - Shai Hulud 攻撃者によって自己ホスト型ランナー上で最初の悪意のあるコードが実行され、認証情報が漏洩しました。ローテーションされた認証情報。
07:21 5 月 14 日 - Grafana-delivery-bot を使用して脅威アクターによって行われた最初の悪意のあるコミットが、Shai Hulud 攻撃者から流出しました。
13:28 5 月 14 日 - リポジトリのデータ漏洩が開始されます。
5 月 15 日 20:57 - データ恐喝の脅威アクターが恐喝要求を公開します。
5 月 16 日 08:30 - Grafana Labs のセキュリティ チームが身代金の請求に気づき、確認を求め始めます。
17:39 5 月 16 日 - 侵害が確認されました。事件が宣言された。
5 月 16 日 19:33 - 既知のすべての影響を受ける認証情報と GitHub アプリケーションが一時停止/ローテーションされました。その他すべてのサスペンションと回転

GitHub アプリケーションとアクセス可能な認証情報が開始されます。
5 月 16 日 21:10 - すべての GitHub アプリケーションの一時停止が完了しました。
5 月 17 日 16:40 - 脅威アクターに関連付けられた GitHub アプリケーション アカウントによって行われたすべてのコード変更が特定され、元に戻されました。
5 月 17 日 16:52 - 根本原因、侵害の攻撃チェーンが特定されました。
5 月 17 日 17:21 - DockerHub の認証情報は侵害されていないと判断されました。
17:51 5 月 17 日 - すべての悪意のあるワークフローの実行が特定されました。影響を受けるシークレットの最終リストが編集およびローテーションされました。影響を受けるリポジトリからの他のすべての CI/共通シークレットのローテーションは続行されます。
5 月 17 日 23:23 - アクセス可能であることが確認された最後の資格情報がローテーションまたは一時停止されました。
03:08 5 月 18 日 - すべての重要でないコードと展開の変更の凍結を開始します。
5 月 25 日 08:00 - 全エンジニアリングのセキュリティ強化週間が始まります。
10:58 5 月 26 日 - コミット レビューが完了し、サービスの解凍が開始されます。リポジトリは、解凍される前に完全にレビューされ、短期的で詳細な範囲の認証情報に GitHub アプリケーション トークン ブローカーを使用するように移行されている必要があります。
10:54 5 月 27 日 - DockerHub にイメージを直接プッシュするリポジトリから、Google Cloud Artifact Registry へのプッシュへの移行が発生します。
5月27日 - 内部調査が完了。追加の攻撃活動や認証情報の漏洩は発見されませんでした。
6 月 2 日 08:00 - 全エンジニアリングのセキュリティ強化週間が終了します。
6 月 3 日 20:43 - データ損失に関するリポジトリのレビューが完了しました。
6月18日 - 厳重な調査が完了し、内部調査が裏付けられた。
調査は現在終了していますが、Grafana Labs のセキュリティ運用を改善する取り組みは継続されます。ドストエフスキーはかつて「理性が失敗すると悪魔が助けてくれる！」と述べました。私たちの哲学を強調するために「罪と罰」を引用しています。私たちは、実際に securi に大きな変化をもたらす変更を実装したかっただけです。

タイ。
私たちは過去 1 か月間、トークン ブローカー、きめ細かいアクセス制御、追加のアラート、静的分析などの影響の大きい制御の実行に費やしてきました。さらに、特定の GitHub アクションを廃止し、有効期間の短いトークンを使用したより厳密な範囲のアクションを使用するようになりました。
また、GitHub 組織を区画化し、すべてのアーカイブされたリポジトリをアクションが無効になった専用の組織に分離するプロセスも開始しました。
今後数週間以内に、当社の対応取り組みの概要と、インシデント後のレビューからセキュリティ体制をどのように改善したかに関する技術的な詳細を共有する予定です。
Grafana 13.1 リリース: コード更新としての可観測性、より多くのデータ ソースにわたる Grafana Assistant の拡張など
Grafana Cloud でのパフォーマンス テストと可観測性
Grafana Loki を使用してログ記録を開始する方法
Grafana Cloud を使用したアプリケーションとフロントエンドの可観測性の概要
Grafana k6 Cloud を使用したシフトレフトのパフォーマンス テスト
ダッシュボードのバージョン履歴を管理する
質問してください。面倒な作業は AI に任せましょう。
送信すると、当社のプライバシー ポリシーに同意したものとみなされます
インシデント対応と管理
インフラストラクチャとクラウドの可観測性
コスト管理と最適化
可観測性調査とレポート
Elasticsearch と Grafana クラウド
ドーナツは私たちの言葉を信じてください。今すぐ Grafana Cloud をお試しください。

## Original Extract

We completed our internal investigation of the recent TanStack supply chain ransom incident and confirmed our initial findings: The incident was strictly limited to the Grafana Labs GitHub environment. There was no unauthorized access to customer production systems, and the Grafana Cloud platform wa
[truncated]

Post-incident review for TanStack npm supply chain ransom incident | Grafana Labs What’s new: Grafana 13 release, the latest in AI, OSS project updates, and more from GrafanaCON2026
EN Downloads Contact Us Grafana Cloud
Site Search Ask AI Blog Post-incident review for TanStack npm supply chain ransom incident: No unauthorized access to customer production systems
On May 27, we completed our internal investigation of the recent TanStack supply chain ransom incident and confirmed our initial findings: The incident was strictly limited to Grafana Labs' GitHub environment. There was no unauthorized access to customer production systems, and the Grafana Cloud platform was not affected.
For an additional, independent audit, we engaged Mandiant, a leader in cybersecurity and incident response. We provided them with API access to Grafana Labs' log environment to conduct queries across our systems for their investigation, which started on June 1. Mandiant confirmed that there was “no evidence of code tampering or repository poisoning within public organizations or production repositories delivered to end users.”
Since we discovered the incident, the Grafana Labs security teams have been running two parallel workstreams: completing the investigation and hardening our security operations. We are publishing this blog in the spirit of transparency to share more details about our incident response and remediation efforts.
If you’re looking for the short version instead of reading our previous updates , here is the TL;DR: The TanStack supply chain attack hit us on May 11 via the Mini Shai-Hulud campaign. At the time, we believed we had successfully rotated every credential involved in this incident. We missed one. I won’t blame this oversight on hubris; the data we had at the time simply led us to believe our rotation was exhaustive. We were mistaken.
A bad actor utilized that overlooked credential to clone our entire repository collection. They then reached out on May 16, demanding a ransom to prevent a code leak.
Since Grafana Labs is an open source company, you might wonder why this is a concern. While most of our source code is public, we do maintain private repos for things like internal tools and specific Grafana Cloud features. It was a heavy decision, but we stuck to our principles and the FBI’s documented guidance : We did not pay.
We launched our mitigation efforts immediately, and we confirmed that there was no unauthorized access to customer production systems, and the Grafana Cloud platform was not affected. We also confirmed that while our codebase was downloaded, it was not altered. Our customers and open source users do not need to take any action.
We were alerted to the incident on a Saturday, and teams across the entire company took action quickly and decisively. (Or to borrow a phrase from one my favorite rappers Big Daddy Kane, ain't no half-stepping at Grafana Labs.)
In response, Grafana Labs suspended all GitHub applications on May 17, initiated a global code freeze on May 18, and conducted a cross-platform audit of Vault, GitHub, Okta, Kubernetes, AWS, GCP, and host logs to verify that no production customer data was compromised.
In the weeks following, our engineering teams contributed to a comprehensive audit that included but was not limited to:
Completing 1,500 security-focused PR reviews
Auditing 280 GitHub applications, stripping permissions and removing several
Scanning 1,200 repositories for any signs of tampering
Executing 2,300 PR reviews looking for unauthorized changes in a single critical repo
Finishing infrastructure audits and retiring legacy systems
Performing wide-ranging new access audits
It was a massive undertaking, but each team stepped up in an extraordinary way to do their part. Engineering, security, and cross-functional partners worked tirelessly to respond, demonstrating the collaboration and the shared commitment we have to our community and our customers that I have always valued here at Grafana Labs.
After the initial assessment, we found that in addition to source code, the downloaded content included GitHub repositories that some Grafana Labs teams use to collaborate on and store internal operational information and other details about our business. This includes, for example, business contact names and email addresses that would be exchanged in a professional setting and email addresses that were used in some past marketing campaigns. This was not information pulled from or processed through the use of production systems or the Grafana Cloud platform.
If you wish to know if email addresses with your domain were identified, please reach out to Grafana Labs support.
19:21 11 May - First malicious code executed on self-hosted runners by Shai Hulud threat actors, leaking credentials. Rotated credentials.
07:21 14 May - First malicious commit made by the threat actor using grafana-delivery-bot, leaked from Shai Hulud attackers.
13:28 14 May - Data exfiltration of repos begins.
20:57 15 May - Data extortion threat actor publishes their extortion demand.
08:30 16 May - Grafana Labs security team becomes aware of the claimed ransom and begins seeking confirmation.
17:39 16 May - Compromise confirmed; incident declared.
19:33 16 May - All known affected credentials and GitHub applications suspended/rotated. Suspension and rotation of all other GitHub applications and accessible credentials begins.
21:10 16 May - Suspension of all GitHub applications completed.
16:40 17 May - All code changes made by GitHub application accounts associated with the threat actor identified and reverted.
16:52 17 May - Root cause, attack chain of compromise identified.
17:21 17 May - DockerHub credentials determined not compromised.
17:51 17 May - All malicious workflow runs identified. Final list of affected secrets compiled and rotated. Rotation of all other ci/common secrets from affected repos continues.
23:23 17 May - Last of the potentially accessible credentials confirmed rotated or suspended.
03:08 18 May - Begin freeze of all non-critical code and deployment changes.
08:00 25 May - All-engineering security hardening week commences.
10:58 26 May - Commit review completed, service thawing begins. A repository needs to have been fully reviewed and transitioned to use a GitHub application token broker for short-term, finely-scoped credentials before being thawed.
10:54 27 May - Transition from repos directly pushing images to DockerHub to pushing to Google Cloud Artifact Registry occurs.
27 May - Internal investigation complete. No additional attack activity or compromised credentials were discovered.
08:00 2 June - All-engineering security hardening week concludes.
20:43 3 June - Review of repositories for data loss completed.
18 June - Mandiant investigation completed, corroborating internal investigation.
The investigation is now closed, but our work to improve security operations at Grafana Labs will continue. Dostoyevsky once noted that "when reason fails, the devil helps!" I’m quoting “Crime and Punishment” to underscore our philosophy: We only wanted to implement changes that actually moved the needle on security.
We’ve spent the past month executing high-impact controls, including a token broker, fine-grained access controls, additional alerting, and static analysis. In addition, we have moved off of certain GitHub Actions and now use more tightly scoped actions with short-lived tokens.
We have also started the process of compartmentalizing our GitHub organizations and isolating all archived repos into a dedicated organization with actions disabled.
We will share an overview of our response efforts and the technical details of how we improved our security posture from our post-incident review in the coming weeks.
Grafana 13.1 release: observability as code updates, extending Grafana Assistant across more data sources, and more
Performance testing and observability in Grafana Cloud
How to get started with logging using Grafana Loki
Intro to Application and Frontend Observability with Grafana Cloud
Shift left performance testing with Grafana k6 Cloud
Manage dashboard version history
Ask your questions. Let AI do the heavy lifting.
By submitting, you agree to our Privacy policy
Incident Response & Management
Infrastructure & Cloud Observability
Cost Management & Optimization
Observability Survey & Reports
Elasticsearch vs. Grafana Cloud
Donut take our word for it. Try Grafana Cloud today.
