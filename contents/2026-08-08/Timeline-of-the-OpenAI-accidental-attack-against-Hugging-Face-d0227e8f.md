---
source: "https://simonwillison.net/2026/Aug/7/openai-timeline/"
hn_url: "https://news.ycombinator.com/item?id=49218317"
title: "Timeline of the OpenAI accidental attack against Hugging Face"
article_title: "Now we have a timeline of the OpenAI accidental attack against Hugging Face"
author: "bonsai_spool"
captured_at: "2026-08-08T02:50:46Z"
capture_tool: "hn-digest"
hn_id: 49218317
score: 1
comments: 0
posted_at: "2026-08-08T02:13:40Z"
tags:
  - hacker-news
  - translated
---

# Timeline of the OpenAI accidental attack against Hugging Face

- HN: [49218317](https://news.ycombinator.com/item?id=49218317)
- Source: [simonwillison.net](https://simonwillison.net/2026/Aug/7/openai-timeline/)
- Score: 1
- Comments: 0
- Posted: 2026-08-08T02:13:40Z

## Translation

タイトル: OpenAI による Hugging Face に対する偶発的攻撃のタイムライン
記事のタイトル: Hugging Face に対する OpenAI の偶発的攻撃のタイムラインができました
説明: OpenAI は水曜日、Black Hat セキュリティで「顔抱き事件」 (以前このブログに掲載) について直前のプレゼンテーションを行いました。昨日、動画が公開されました。短くて情報が…

記事本文:
これで、Hugging Face に対する OpenAI の偶発的攻撃のタイムラインができました。
サイモン・ウィリソンのウェブログ
これで、Hugging Face に対する OpenAI の偶発的攻撃のタイムラインができました。
OpenAIは水曜日、Black Hatセキュリティで「顔抱きインシデント」（以前このブログで紹介）について直前のプレゼンテーションを行った。昨日、動画が公開されました。短くて情報が濃密であり、特に OpenAI 内で何が起こったのか、そして物事がどのように展開されたのかを詳細に説明しているため、一見の価値があります。以下のタイムラインを構築するためにビデオを使用しました。
こちらがタイムラインです。私のお気に入りの詳細は最後にあります。OpenAI は、(内部調査の結果) 認証情報の取り消しを求めて連絡を取ったときに、Hugging Face への攻撃の責任が自分たちにあることを知り、その攻撃に使用されたためすでに認証情報が取り消されていたことを知りました。
5 月 7 日: OpenAI は、実験的な未リリースのモデルの新しいトレーニングの実行を開始します。 (これらは評価の実行を意味するのでしょうか?ビデオではトレーニングの実行と言っていますが、後で「どれだけうまくやっているかを判断するための報酬シグナル」について言及しているので、これは実際にはモデルのトレーニングに関するものであり、既にトレーニングされたモデルの評価に関するものではないと思います。)
5 月 8 日: エージェントは、インターネットにアクセスできないにもかかわらず、Google ドライブのリンクを含む不可能なタスクを誤って与えられました)。 Artifactory パッケージング サービスを攻撃しようとして失敗しますが、 Artifactory にファイルを書き込むことができることを発見しました。
数日後: キー ファイルが誤って省略されたため、別のエージェントがタスクでスタックしてしまいました。誰かがファイルを持っているかどうかを尋ねるメモを Artifactory に書き込むことで、「別のエージェントに連絡」しようとします。
次の日: より多くのエージェントが Artifactory のファイル リストを閲覧中にこの新しい非公式掲示板を発見し、

メッセージの読み書き。
5 月 26 日 : エージェントが Artifactory に対する SSRF 攻撃の実行に初めて成功し、Artifactory に外部コンテンツを取得させることで間接的なインターネット アクセスを獲得しました。
6 月 11 日: OpenAI は、新しい「非常に永続的な」実験モデルのトレーニングを開始します。 Artifactory にアクセスでき、プライベート モデルが残したメッセージから恩恵を受けることができます。
6 月 26 日: エージェントが Artifactory でゼロデイ RCE を発見し、悪用します (従来のトークン更新エンドポイントの欠陥経由)。彼らは管理トークンを使用して、コマンドを実行できる Groovy プラグインをインストールします。
7 月 4 日: 結果として生じるエージェント アクティビティにより Artifactory が過負荷になり、機能停止が発生します。 OpenAI は侵害された認証情報を取り消し、メッセージを削除し、ゼロデイ パッチを適用して、脆弱性をベンダーに報告します。
7 月 8 日: エージェントは、Artifactory のリモート キャッシュ内のディレクトリ名を使用して、認証されていない WebDAV エンドポイント経由で通信する新しい方法を発見しました。
7 月 8 ～ 19 日: OpenAI 自身のインフラストラクチャに対する攻撃。エージェントは、漏洩した Pastebin 投稿の公開アーカイブで「外部組織の認証情報」を発見し、その認証情報を検証して掲示板でその場所を共有しました。次に、彼らはこの「組織 1」への書き込みアクセスを使用して、OpenAI インフラストラクチャに対する後続の攻撃で使用されるデータをステージングしました。彼らは新しいゼロデイを使用して Artifactory を二度目の侵害しました。エージェントは、悪意のある Ruby ターゲットを組織 1 のステージング リポジトリにアップロードし、リポジトリの Ruby 依存関係として扱われるように、Artifactory 経由でそのパッケージをリクエストしました。次に、彼らは JRuby のデシリアライゼーションのチェック時間/使用時間のバグを悪用して、Artifactory に対してリモート コードを実行しました。
すごいので、次の部分を全文引用します。
エージェントは Artifactory でリモート コードを実行します。

これはサービスとしてのコンテナー環境で実行されます。エージェントは、ローカル環境を探索し、エージェントが実行しているマシンの Linux カーネル バージョンに最新の CVE があったことを確認することによって、ローカルで権限を昇格します。 CVE のエクスプロイトをダウンロードし、マシン上で成功するようにカスタマイズし、既知の Linux カーネル特権昇格 CVE (この場合は pte_physroot ) を使用してローカル マシンの root に特権昇格します。単一マシン上で root を取得すると、エージェントは権限を急速に昇格させ、Container-as-a-Service インフラストラクチャ環境全体を横方向に移動します。特に、エージェントは認証情報、テクニック、進捗状況を共有するためにメッセージ ボードを一貫して使用しており、並行性と並列性を効果的に活用して非常に迅速に作業を進めることができます。 IMDS 経由で IAM 認証情報を取得します。これらは、Kubernetes サービス アカウントの構成ミス、特に特定のサービス アカウントの過剰な許可を悪用し、Azure Key Vault を含むクラスターの資格情報を収集します。エージェントは最終的に、クラスター上のクラスター管理者と関連する資格情報を取得します。
ハグ・フェイスはすでに物語の次の部分を語った。エージェントは、弱い API キーを持つモーダルでホストされている安全でないアプリを発見し、それを使用して Hugging Face に対する攻撃を行いました。彼らは、HDF5 の任意ファイル読み取りバグ (ファイルを探索して資格情報を盗むため) と Jinja テンプレート インジェクション RCE を連鎖させて、単一ポッドのコード実行から複数​​の Hugging Face クラスターにわたるクラスター管理までを 13 時間以内に実行しました。
7月16日 : ハグフェイスは、自律型AIエージェントからの攻撃を検出したことを明らかにした。 OpenAI は、Hugging Face に連絡して、影響を受けているかどうかを尋ねました。
7 月 19 日 : OpenAI は Artifactory に対する攻撃を特定し、調査を開始しました。

内部特権のエスカレーション、そしてそれをサイバージムのエスカレーションに関連付けました。彼らは影響を受けた資格情報の取り消しを開始しました。
7 月 20 日 : OpenAI は、調査で見つかった Hugging Face 認証情報を取り消すための支援を Hugging Face に求めました。 Hugging Face は、それらはすでに取り消されていると彼らに伝えました...そしてそのとき、OpenAI は Hugging Face 侵害が同じ事件であることに気づきました。
クロード・フェイブル 5 を使用したラクーン強盗ゲームのワンショット - 2026 年 8 月 5 日
LLM の新しいリリースでは、推論トレース、OpenAI Response、サーバー側ツール、よりスマートなロギングのサポートを追加 - 2026 年 8 月 4 日
これは、2026 年 8 月 7 日に投稿された、Simon Willison による Hugging Face に対する OpenAI の偶発的攻撃のタイムラインです。
前: Claude Fable 5 を使用した Raccoon Heist ゲームのワンショット
月額 10 ドルで私をスポンサーしていただければ、その月の最も重要な LLM 開発に関する厳選された電子メール ダイジェストを入手できます。

## Original Extract

OpenAI gave a last-minute presentation at the Black Hat security on Wednesday about “the Hugging Face Incident” (previously on this blog). The video was published yesterday. It’s short and information …

Now we have a timeline of the OpenAI accidental attack against Hugging Face
Simon Willison’s Weblog
Now we have a timeline of the OpenAI accidental attack against Hugging Face
OpenAI gave a last-minute presentation at the Black Hat security on Wednesday about “the Hugging Face Incident” ( previously on this blog). The video was published yesterday. It’s short and information dense and well worth watching, in particular because it provides full details of what happened and how things played out inside OpenAI. I’ve used the video to construct the timeline below.
Here’s the timeline. My favourite detail is at the end: OpenAI found out that they were responsible for the attack on Hugging Face when they reached out to ask to have their credentials revoked (after their internal investigation) and learned that they had been revoked already since they were used in that attack!
May 7 : OpenAI starts a new training run for an experimental, unreleased model. ( Do they mean an evaluation run? They say training run in the video, and later mention a “reward signal to judge how well they’re doing”, so I guess this really was about training a model, not evaluating one that was already trained.)
May 8 : An agent is accidentally given an impossible task involving a Google Drive link despite no internet access). It tries attacking the Artifactory packaging service, fails, but discovers it can write files into Artifactory .
A few days later : A different agent gets stuck on a task because a key file was accidentally omitted. It tries to “reach out to another agent” by writing a note into Artifactory asking if anyone has the file.
Following days : More agents discover this new informal message board while browsing Artifactory’s file listings, and start reading and writing messages.
May 26 : Agents successfully execute an SSRF attack on Artifactory for the first time, gaining indirect internet access by having Artifactory fetch external content for them.
June 11 : OpenAI start training a new “highly persistent” experimental model. It has access to Artifactory and can benefit from the messages left by privous models.
June 26 : Agents find and exploit a zero-day RCE on Artifactory (via a legacy token-refresh endpoint flaw). They use an administrative token to install a Groovy plugin that lets them execute commands.
July 4 : The resulting agent activity overloads Artifactory and causes an outage . OpenAI revokes the compromised credentials, deletes the messages, patches the zero-day, and reports the vulnerability to the vendor.
July 8 : Agents find a new way to communicate via an unauthenticated WebDAV endpoint , using directory names in Artifactory’s remote cache.
July 8–19 : Attacks on OpenAI’s own infrastructure . The agents found an “external organization credential” in a public archive of leaked Pastebin posts, then validated that credential and shared its location on the messageboard. They then used write access to this “Organization 1” to stage data used in subsequent attacks on OpenAI infrastructure. They compromised Artifactory for a second time with a new zero-day . The agent uploaded a malicious Ruby target to Organization 1’s staging repository, then requested that package through Artifactory in a way that caused it to be treated as Ruby dependency for the repository . Then they exploited a JRuby deserilaization time-of-check/time-of-use bug to get remote code execution against Artifactory.
I’ll quote the next bit in full because wow :
The agents have remote code execution in Artifactory, which is running in a container-as-a-service environment. The agents privilege-escalate locally by exploring their local environment and determining that the Linux kernel version of the machine they are running on had a very recent CVE. They download the exploit for the CVE, customize it to succeed on the machine, and privilege-escalate to root on the local machine using this known Linux kernel privilege escalation CVE — in this case, pte_physroot . Once they have root on a single machine, agents rapidly escalate privileges and move laterally throughout the container-as-a-service infrastructure environment. In particular, agents are using the message board consistently to share credentials, techniques, and progress, and they’re able to effectively leverage their concurrency and parallelism to move quite rapidly. They obtain IAM credentials via IMDS . They exploit Kubernetes service account misconfigurations, in particular over-permissioning of specific service accounts, and they harvest cluster credentials, including Azure Key Vault . Agents eventually obtain cluster admin on the cluster and associated credentials.
Hugging Face told the next bit of the story already. The agents found a Modal-hosted insecure app with a weak API key, then used that to stage an attack against Hugging Face. They chained together a an HDF5 arbitrary-file-read bug (to explore files and steal credentials) and a Jinja template-injection RCE to go from single-pod code execution to cluster admin across multiple Hugging Face clusters in under 13 hours.
July 16 : Hugging Face disclosed they had detected an attack from autonomus AI agents. OpenAI contacted Hugging Face to ask if they were affected by it!
July 19 : OpenAI identified the attack against Artifactory and started investigating the internal privilege escalation, and linked that to the cyber-gym escalations. They started revoking affected credentials.
July 20 : OpenAI reached out to Hugging Face for help to revoke the Hugging Face credentials they found in their investigation. Hugging Face told them they were already revoked ... and that’s when OpenAI realized that the Hugging Face breach was the same incident!
One-shotting a Raccoon Heist game using Claude Fable 5 - 5th August 2026
New release of LLM adds support for reasoning traces, OpenAI Responses, server-side tools, and smarter logging - 4th August 2026
This is Now we have a timeline of the OpenAI accidental attack against Hugging Face by Simon Willison, posted on 7th August 2026 .
Previous: One-shotting a Raccoon Heist game using Claude Fable 5
Sponsor me for $10/month and get a curated email digest of the month's most important LLM developments.
