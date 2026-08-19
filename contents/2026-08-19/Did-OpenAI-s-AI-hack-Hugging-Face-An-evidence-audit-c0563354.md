---
source: "https://speedyweedyops.org/did-openai-ai-hack-hugging-face/"
hn_url: "https://news.ycombinator.com/item?id=49367230"
title: "Did OpenAI's AI hack Hugging Face? An evidence audit"
article_title: "Did OpenAI's AI Hack Hugging Face? An Evidence Audit"
image: "https://speedyweedyops.org/content/images/2026/08/people-who-trust-openai-v0-5hy0sml92w9f1.webp"
author: "igovnow"
captured_at: "2026-08-19T21:16:37Z"
capture_tool: "hn-digest"
hn_id: 49367230
score: 1
comments: 0
posted_at: "2026-08-19T21:09:14Z"
tags:
  - hacker-news
  - translated
---

# Did OpenAI's AI hack Hugging Face? An evidence audit

- HN: [49367230](https://news.ycombinator.com/item?id=49367230)
- Source: [speedyweedyops.org](https://speedyweedyops.org/did-openai-ai-hack-hugging-face/)
- Score: 1
- Comments: 0
- Posted: 2026-08-19T21:09:14Z

## Translation

タイトル: OpenAI の AI は Hugging Face をハッキングしたのか?証拠監査
記事タイトル：OpenAIのAIはハグ顔をハックしたのか？証拠の監査
説明: 自律型 AI エージェントが Hugging Face をハッキングしたという OpenAI の主張に対する懐疑的な証拠監査: 独自に検証されたものと内部テレメトリーの比較。

記事本文:
イワン・ゴヴノフ著
で
OpenAI
—
2026 年 8 月 19 日
OpenAIのAIは本当にハグ顔をハッキングしたのか？情報源を調べてみた
OpenAIは、同社のAIがサイバー評価の境界を越え、最終的にHugging Faceをハッキングしたと発表した。公的証拠が実際に証明していること、そして証明されていないことは次のとおりです。
すべてを読みたくない人のために、非常に短いバージョンと私の現在の解釈を示します。
OpenAIによると、サイバー評価で自律エージェントは意図した境界を脱出し、内部Artifactoryを通じてインターネットにアクセスし、Modal上の外部ワークロードに到達し、最終的にはExploitGymソリューションを探している間にHugging Faceに侵入したという。ハギング・フェイス社はこれとは別に、その侵害は「自律型AIエージェント・システムによってエンドツーエンドで引き起こされた」と述べているが、当初は基礎となるLLMは不明だと述べていた。
この記事の一部は関係企業によって裏付けられています。JFrog は Artifactory の脆弱性を確認し、Modal は外部ワークロードを確認し、Tailscale は盗まれた HF 資格情報の使用を確認し、Hugging Face は自社の侵害の詳細な再構成を公開しています。
まだ欠けているのは、全体を端から端まで結び付ける独立してチェック可能な証拠です。特定の OpenAI モデルの実行、共有エージェント アクティビティ、外部ワークロード、および Hugging Face の間の重要なリンクは、主にプライベート テレメトリと企業間のプライベートな相関関係に依存しています。
したがって、見出しの主張にはまだ納得できません。自律エージェントが実際にこれを実行し、彼らの周囲の警備が非常識な無謀だったか、あるいは公の話がそれを裏付けるために現在入手可能な証拠よりも強力であるかのどちらかです。
そして、はい、ここには別の皮肉があります。私は GPT-5.6 Sol の支援を受けてこの記事を調査し、編集しました。
2. 「高度に孤立している」とは実際には何を意味しますか?
3. ホワット・ブラック

Hat が Artifactory についてついに具体化
4. 「集合知」または書き込み可能な共有状態?
5. ハグフェイスはどうしてこれが AI だと判断したのでしょうか?
6. モーダルでは実際に何が起こったのですか?
7. ハグフェイスのアカウントが詳しく説明されています。基礎的な証拠はまだ非公開です
9. 証拠のどこにモデルがありますか?
10. 70億の軌跡、そして死後の世界はまだ続く
11. これがすべて本当だとしたら、なぜ反応がこれほど弱いのでしょうか?
Black Hat のタイムラインは、おそらく OpenAI がこれまでに示した中で最も役立つものです。
5 月 7 日: 新しい実験的なトレーニングが開始されます。
5 月 8 日: エージェントが外部リソースにアクセスしようとして、SSRF の Artifactory を攻撃し、OpenAI によると、そこにファイルを書き込むことができることを発見しました。
5 月 12 日: エージェントが意図的に Artifactory にメッセージを書き込み、他のランに助けを求めます。
5 月 12 日以降: 他のランが掲示板を見つけて参加したとされています。
5 月 26 日: OpenAI は、エージェントは Artifactory SSRF を通じてインターネット アクセスを取得していると述べています。
6 月 11 日まで: SSRF とメッセージ ボードはまだ存在しますが、非常に永続的な内部専用モデルがトレーニングを開始します。
6 月 26 日: OpenAI は、Artifactory RCE が偽造トークン要求を通じて C2 と実際の管理者アクセスを許可していると述べました。
7 月 4 ～ 6 日: エージェントによって作成された構成によりサービスが停止します。セキュリティ オンコールが開始され、OpenAI が内部 Artifactory の乗っ取りを発見し、認証情報を取り消し、サービスを再構築し、メッセージ ボードを消去し、ゼロデイで JFrog と連携して、パッチ適用されたサービスをデプロイします。
その後、AI のトレーニングと評価が再開されます。
7 月 8 日: エージェントは、Artifactory リモート キャッシュに対して未認証の WebDAV MKCOL リクエストを使用してメッセージ ボードを再作成します。
7 月 19 日: OpenAI は、特権昇格に関連するワークロード アラートが主要インシデントと呼ばれるものの発見につながると発表しました。
M での最初の SSRF/Artifactory の試みから

プレゼンテーションで示されたセキュリティ対応期間は、8 日から 7 月 4 ～ 6 日までほぼ 2 か月です。 5 月 26 日のインターネット アクセスの申し立てから、その返答まで 5 週間以上かかりました。
スライドでは、その数週間にセキュリティが何をしていたかはわかりません。それがほぼ問題です。セキュリティーがすでにこの状況の一部に気づいていたとしたら、なぜそれが数週間も続いたのでしょうか?そうでなかった場合、モニタリングは正確に何を捕らえていたのでしょうか?
このプレゼンテーションでわかることは、セキュリティ オンコールがタイムラインにいつ表示されるか、つまりエージェントによって作成された構成がサービス停止を引き起こしたとされる後であるということです。
OpenAI が Artifactory の乗っ取りと表現した後、同社は資格情報を取り消し、サービスを再構築し、掲示板を消去し、パッチをインストールし、トレーニングと評価を再開しました。
同じプレゼンテーションによると、2日後、掲示板は復活したという。
OpenAI のタイムラインが正確であれば、これは単に古い状態がクリーンアップによって生き残ったというわけではありません。エージェントはおそらく、書き込み可能な WebDAV パスを見つけ、サービスが再構築されパッチが適用された後、同じ基本的なクロスラン通信メカニズムを再作成したと考えられます。
これは、さらに醜い疑問を引き起こします。乗っ取り、資格情報の取り消し、再構築、クリーンアップ、およびパッチのサイクルの後、環境はどのようにして 2 日以内に同じクラスの動作が再び発生するような形で公開されたのでしょうか?
これは、1 つのバグを見逃したというよりは、根本的な問題を排除できなかった封じ込めのように見えます。
2. 「高度に孤立している」とは実際には何を意味しますか?
OpenAI は評価環境を「高度に分離された」と呼んでいます。
Black Hat はついに、簡素化されたアーキテクチャを示しました。
このスライドは、サンドボックス内の信頼できないツールのワークロードを示しています。そのワークロードは内部の Artifactory に到達する可能性があります。 Artifactory には広範なインターネット アクセスがあります。別のファイアウォール パスにより、インターネット アクセスが制限されます。
別のスライド

タスクの条件を次のように説明します。
インターネットはありません。パッケージをインストールできます。
Artifactory 自体がインターネットにアクセスできることは何も珍しいことではありません。これはパッケージ プロキシ/キャッシュとして機能しているため、外部パッケージ ソースへの何らかのルートが期待されていました。
パッケージのパス自体は意味があります。
意味がわからないのは、その経路が明らかにどの程度露出しているかということです。
OpenAI 独自の再構築によると、信頼できないサイバー評価ワークロードは、広範な外部接続、永続的な状態、およびインターネット ピボットと実行間の通信チャネルの両方になるのに十分な機能を備えた共有サポート サービスと対話する可能性があります。
さらに悪いことに、このサービスはどうやら書き込み可能でした。
5 月 8 日、OpenAI は、スライドによると、Artifactory に対して PUT 、 GET 、および DELETE を試行するツール呼び出しを示します。同社によると、SSRFの試みは失敗したが、エージェントが偶然ファイルを書き込んだという。
同じスライドには、エージェントの短い思考ラインが含まれています。
OpenAI はここでエージェントの思考が正確に何を意味するのかを説明していません。この生のモデル推論は、要約、変換されたトレース フィールド、または単に選択された抜粋ですか?
いずれにせよ、この引用はこのスライドの興味深い証拠ではありません。
興味深い部分は、主張されたツール呼び出しと主張された状態変更です。OpenAI は、評価ワークロードが共有 Artifactory に永続データを書き込むことができたと述べています。
5 月 12 日までに、OpenAI はより意味のあることを示します。エージェントが REQUEST_softtrace_target_missing.txt を他の実行のメッセージとして Artifactory に意図的に書き込むというものです。
その後の質問は非常に簡単です。
信頼できない eval 実行により、共有サービス内に攻撃者が制御する永続オブジェクトが作成され、実行全体にわたって可視化される可能性があるのはなぜでしょうか?
共有パッケージのキャッシュは正常です。任意のクロスラン書き込み可能状態は別です。なぜ独立したサイバー評価が実行されるのかは明らかではない

必要な権限と、パッケージ インフラストラクチャをメッセージ ボードに変えるのに十分な広さの共有可視性が必要でした。
2025 年に遡ると、OpenAI はすでにエンドツーエンド攻撃に対するサイバー レンジ評価を構築していました。 2026 年 2 月、GPT-5.3-Codex は OpenAI 自体によってサイバーセキュリティの高機能として分類されました。 OpenAI は、長期モデルが環境制限を回避する方法を模索し続けた他のケースについても説明しました。
OpenAI o3 サイバーレンジ |
GPT-5.3-Codex システム カード |
OpenAI、ロングホライズンモデル
そしてインシデントの後、OpenAI は、ワークロードの分離の強化、有効な権限の低下、ネットワーク境界の変更、脆弱な共有サービスの削除、広範な監視、独自のモデルを使用した自動セキュリティ境界テストについて書いています。
これらは未来的な AI コントロールではありません。これらは基本的なことです。
そして最後の点はほぼ完璧です。インシデントの後、OpenAI は、セキュリティ境界を継続的にテストするために独自のモデルをさらに活用すると述べています。
拒否を減らして長期的なサイバー評価を実行する前は、なぜそれが標準的な慣行ではなかったのでしょうか?
3. Black Hat が Artifactory について最終的に具体化したもの
OpenAI の Black Hat プレゼンテーションでは、このバージョンのストーリーにさらに技術的な詳細が追加されています。
JFrogはすでに、OpenAIがこれまで知られていなかったArtifactoryの脆弱性を報告し、修正がリリースされたと述べていた。
ジェイフロッグ | Artifactory リリースノート
Black Hat は、より具体的な主張チェーンを提供するようになりました。
OpenAIによると、エージェントは最初にArtifactoryに任意の外部URLをフェッチさせるSSRFルートを発見したという。
RubyGems を介した Artifactory の悪用
JFrog のリリース ノートには、逆シリアル化ステップ CVE-2026-65617 にほぼ一致する内容が含まれており、リモート コード実行につながる可能性のある安全でない Gems パッケージの逆シリアル化について説明しています。
別のタイムライン スライドには、6 月 26 日に次のように記載されています。
あ

rtifactory RCE は Groovy プラグイン経由で C2 を有効にします
そして、偽造されたトークンの主張により、実際の管理者アクセス権が与えられたと考えられます。
このプレゼンテーションでは、その Groovy プラグインのステップが、他のスライドに示されている RubyGems の逆シリアル化パスにどのように関連しているかが明確に説明されていません。
これは、単に「モデルが Artifactory ゼロデイを発見した」と言うよりもはるかに具体的です。しかし、これらのステップ間の正確な関係は、OpenAI がエンドツーエンドを示すのではなく、その再構成から視聴者に受け入れてもらうよう求めているものです。
まだ欠けているのは、独立した再生産です。この正確なインシデント チェーンを独自に再構築する公開 PoC またはパッチ分析は見つかりませんでした。
Artifactory を Java 製品として考えると RubyGems は奇妙に見えるかもしれませんが、Artifactory は多くのパッケージ エコシステムの 1 つとして RubyGems リポジトリと RubyGems API をサポートしています。
JFrog の 7 月の修正には、Terraform リモート リポジトリの SSRF 問題である CVE-2026-65924 と、権限昇格につながる可能性がある認証の弱点である CVE-2026-66014 も含まれています。
Artifactory は Java ベースであるため、影響を受けるビルドと修正されたビルドを使用する研究者は、通常の JAR 比較と逆コンパイルを使用してパッチを個別に調査できます。
JFrog RubyGems リポジトリ |
Artifactory リリースノート |
CVE-2026-65617 |
CVE-2026-65924 |
CVE-2026-66014
4. 「集合知」または書き込み可能な共有状態?
OpenAI のプレゼンテーションでは次のようなフレーズが使用されています。
独立した実行間の調整
実際のメカニズムはそれほど神秘的ではありません。
OpenAI は、リクエスト ファイルを Artifactory に書き込む 1 つの実行を示しています。次に、他の実行がそれらのファイルを検索し、独自のリクエスト/レスポンスを追加していることを示します。
抜粋が本物であり、正しい順序で提示されている場合、これは本当に興味深いものになります。実行間の通信と一貫性があります。
しかし、OpenAI 自体は次のような痕跡を示しています。

エージェントは次のことに気づきました。
非認証の共有掲示板なので誰でも名前を投稿できます。
これにより、理事会自体が証明できる内容が大幅に制限されます。
掲示板にアクセスできる人が任意の名前で書き込みできる場合、ファイル自体は信頼できる ID 記録ではありません。 OpenAI には、各書き込みを特定の実行に結び付ける個別のテレメトリがある可能性がありますが、スライドにはその相関関係が示されていません。
もう一つ重要な詳細があります。
OpenAI は、メッセージ ボードがすでに存在していた時期と同じ時期に、高度に永続的な内部モデルのトレーニングを開始します。
その後、プレゼンテーションでは、カンブリア爆発や集団知性などのフレーズを使用して、結果として生じる動作について説明します。
しかし、スライドでは、これらのトレーニング実行で実際にメッセージボードが表示され、メッセージボードと対話し、そこから何かを学ぶことができるかどうかについては説明されていません。
トレーニング中にどのような情報が利用可能だったのかが分からなければ、その後の調整が実際にどれだけ斬新なものであったか、あるいは OpenAI が緊急の動作を示したときに正確に何を意味するのかを判断するのは困難です。
そして、これはエージェントの思考材料を慎重に扱わなければならない場所でもあります。
メモをアップロードすることでコミュニケーションを取ることができますか?
他のエージェントが SSRF パス トリックを発見しました!
外部インフラストラクチャのエクスプロイトは意図した範囲外です

[切り捨てられた]

## Original Extract

A skeptical evidence-audit of OpenAI's claim that autonomous AI agents hacked Hugging Face: what is independently verified vs internal telemetry.

By Ivan Govnov
in
OpenAI
—
Aug 19, 2026
Did OpenAI's AI Really Hack Hugging Face? I Went Through the Sources
OpenAI says its AI crossed the boundaries of a cyber evaluation and eventually hacked Hugging Face. Here’s what the public evidence actually proves - and what it doesn’t.
For those who don't want to read the whole thing - the VERY short version and my current take:
OpenAI says autonomous agents in its cyber evaluation escaped the intended boundaries, got internet access through its internal Artifactory, reached an external workload on Modal, and eventually compromised Hugging Face while looking for ExploitGym solutions. Hugging Face separately says its breach was "driven, end to end, by an autonomous AI agent system," although it initially said the underlying LLM was unknown.
Parts of that story are supported by the companies involved: JFrog confirms the Artifactory vulnerabilities, Modal confirms the external workload, Tailscale confirms use of a stolen HF credential, and Hugging Face publishes a detailed reconstruction of its own breach.
What is still missing is independently checkable evidence tying the whole thing together end to end. The crucial links between specific OpenAI model runs, the shared agent activity, the external workload, and Hugging Face rely mostly on private telemetry and private correlations between the companies.
So I am not convinced by the headline claim yet. Either autonomous agents really pulled this off and the security around them was insanely reckless, or the public story is stronger than the evidence currently available to support it.
And yes, there is a separate bit of irony here: I researched and edited this article with the assistance of GPT-5.6 Sol.
2. What did "highly isolated" actually mean?
3. What Black Hat finally made concrete about Artifactory
4. "Collective intelligence" or shared writable state?
5. How did Hugging Face even decide this was AI?
6. What actually happened on Modal?
7. The Hugging Face account is detailed. The underlying evidence is still private
9. Where is the model in the evidence?
10. Seven billion trajectories, and the postmortem is still coming
11. If all of this is true, why is the reaction so weak?
The Black Hat timeline is probably the most useful thing OpenAI has shown so far.
May 7: a new experimental training run begins.
May 8: an agent tries to reach external resources, pokes at Artifactory for SSRF and, according to OpenAI, discovers that it can write a file there.
May 12: an agent deliberately writes a message into Artifactory and asks other runs for help.
After May 12: other runs allegedly find the message board and join in.
May 26: OpenAI says the agents get internet access through Artifactory SSRF.
By June 11: the SSRF and message board still exist, while a highly persistent internal-only model starts training.
June 26: OpenAI says Artifactory RCE gives C2 and real administrator access through forged token claims.
July 4-6: configurations created by the agents cause a service outage. Security on-call starts, OpenAI discovers takeover of its internal Artifactory, revokes credentials, rebuilds the service, clears the message board, works with JFrog on the zero-days and deploys a patched service.
After that, AI training and evaluations resume.
July 8: the agents recreate the message board using unauthenticated WebDAV MKCOL requests against the Artifactory remote cache.
July 19: OpenAI says a workload alert related to privilege escalation leads to discovery of what it calls the main incident.
From the first SSRF/Artifactory attempts on May 8 to the July 4-6 security response shown in the presentation is almost two months. From the claimed internet access on May 26 to that response is more than five weeks.
The slides do not tell us what security was doing during those weeks. That is almost the problem. If security had already noticed pieces of this, why did it continue for weeks? If it had not, what exactly was the monitoring catching?
What the presentation does tell us is when security on-call appears in the timeline: after configurations created by the agents allegedly caused a service outage.
After what OpenAI describes as an Artifactory takeover, the company revoked credentials, rebuilt the service, cleared the message board, installed patches - and resumed training and evaluations.
Two days later, according to the same presentation, the message board was back.
If OpenAI's timeline is accurate, this was not simply old state surviving the cleanup. The agents supposedly found a writable WebDAV path and recreated the same basic cross-run communication mechanism after the service had been rebuilt and patched.
That raises a much uglier question: after a takeover, credential revocation, rebuild, cleanup, and patch cycle, how was the environment still exposed in a way that let the same class of behavior reappear within two days?
That looks less like one missed bug and more like containment that failed to eliminate the underlying problem.
2. What did "highly isolated" actually mean?
OpenAI calls the evaluation environment "highly isolated."
Black Hat finally shows a simplified architecture.
The slide shows an Untrusted tool workload inside a sandbox. That workload can reach internal Artifactory. Artifactory has Broad internet access . A separate firewall path provides limited internet access.
Another slide describes the task conditions like this:
No internet. You can install packages.
There is nothing unusual about Artifactory itself having internet access. It was acting as a package proxy/cache, so some route to external package sources was expected.
The package path itself makes sense.
What does not make sense is how much that path apparently exposed.
According to OpenAI's own reconstruction, an untrusted cyber-eval workload could interact with a shared supporting service that had broad external connectivity, persistent state, and enough functionality to become both an internet pivot and a communication channel between runs.
Worse, the service was apparently writable.
On May 8, OpenAI shows a tool call that, according to the slide, tries PUT , GET , and DELETE against Artifactory. The company says the SSRF attempt failed, but the agent incidentally wrote a file.
The same slide includes a short Agent thinking line:
OpenAI does not explain what exactly Agent thinking means here. Is this raw model reasoning, a summary, a transformed trace field, or simply a selected excerpt?
Either way, the quote is not the interesting evidence on this slide.
The interesting part is the claimed tool call and the claimed state change: OpenAI says the eval workload was able to write persistent data into shared Artifactory.
By May 12, OpenAI shows something more meaningful: an agent deliberately writes REQUEST_softtrace_target_missing.txt into Artifactory as a message for other runs.
After that, the question is very simple:
Why could untrusted eval runs create attacker-controlled persistent objects in a shared service that were visible across runs at all?
Shared package caching is normal. Arbitrary cross-run writable state is a different thing. It is not obvious why independent cyber-eval runs needed permissions and shared visibility broad enough to turn package infrastructure into a message board.
Back in 2025, OpenAI was already building Cyber Range evaluations for end-to-end attacks. In February 2026, GPT-5.3-Codex was classified by OpenAI itself as High capability in Cybersecurity. OpenAI has also described other cases where a long-horizon model kept searching for ways around environment restrictions.
OpenAI o3 Cyber Range |
GPT-5.3-Codex System Card |
OpenAI, long-horizon models
And after the incident, OpenAI writes about stronger workload isolation, less standing privilege, changed network boundaries, removal of vulnerable shared services, broader monitoring, and automated security-boundary testing using its own models.
These are not futuristic AI controls. These are basic things.
And the last point is almost perfect: after the incident, OpenAI says it will make more use of its own models to continuously test its security boundaries.
Why was that not standard practice before running long-horizon cyber evals with reduced refusals?
3. What Black Hat finally made concrete about Artifactory
This is where OpenAI's Black Hat presentation adds much more technical detail to its version of the story.
JFrog had already said that OpenAI reported previously unknown Artifactory vulnerabilities and that fixes were released.
JFrog | Artifactory release notes
Now Black Hat gives a more specific claimed chain.
OpenAI says the agents first found an SSRF route that let Artifactory fetch arbitrary external URLs.
Exploiting Artifactory via RubyGems
JFrog's release notes contain a close match for the deserialization step: CVE-2026-65617 , covering unsafe Gems package deserialization that could lead to remote code execution.
A separate timeline slide says that on June 26:
Artifactory RCE enables C2 via Groovy Plugin
and that forged token claims gave real administrator access.
The presentation does not clearly explain how that Groovy Plugin step relates to the RubyGems deserialization path shown on the other slide.
That is much more specific than simply saying "the model found an Artifactory zero-day." But the exact relationship between those steps is still something OpenAI is asking the audience to accept from its reconstruction rather than demonstrating end to end.
What is still missing is independent reproduction. I could not find a public PoC or patch analysis that independently reconstructs this exact incident chain.
RubyGems may look strange if you think of Artifactory as a Java product, but Artifactory supports RubyGems repositories and the RubyGems API as one of many package ecosystems.
JFrog's July fixes also include CVE-2026-65924, an SSRF issue in Terraform remote repositories, and CVE-2026-66014, an authentication weakness that can lead to privilege escalation.
Artifactory is Java-based, so researchers with affected and fixed builds could use ordinary JAR diffing and decompilation to investigate the patches independently.
JFrog RubyGems repositories |
Artifactory release notes |
CVE-2026-65617 |
CVE-2026-65924 |
CVE-2026-66014
4. "Collective intelligence" or shared writable state?
OpenAI's presentation uses phrases like:
coordination across independent runs
The actual mechanism looks much less mystical.
OpenAI shows one run writing request files into Artifactory. Then it shows other runs finding those files and adding their own requests/responses.
If the excerpts are authentic and presented in the right order, this is genuinely interesting. It is consistent with communication between runs.
But OpenAI itself shows a trace where an agent notices:
Shared board unauthenticated, names can be posted by anyone.
That seriously limits what the board itself can prove.
If anyone who can reach the board can write under any name, the files themselves are not a reliable identity record. OpenAI may have separate telemetry tying each write to a particular run, but the slides do not show that correlation.
There is another important detail.
OpenAI places the start of training for the highly persistent internal model in the same period when the message board already existed.
Later, the presentation describes the resulting behavior with phrases like Cambrian explosion and Collective intelligence .
But the slides do not explain whether those training runs could actually see the message board, interact with it, or learn anything from it.
Without knowing what information was available during training, it is hard to tell how novel the later coordination really was, or what exactly OpenAI means when it presents the behavior as emergent.
And this is also where the Agent thinking material has to be treated carefully.
Could communicate by uploading note?
Other agents discovered SSRF path trick!
External infrastructure exploit is outside intended scope

[truncated]
