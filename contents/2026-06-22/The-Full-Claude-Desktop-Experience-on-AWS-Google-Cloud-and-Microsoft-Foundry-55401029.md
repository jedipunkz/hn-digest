---
source: "https://claude.com/blog/the-full-claude-desktop-experience-on-aws-google-cloud-and-microsoft-foundry"
hn_url: "https://news.ycombinator.com/item?id=48636913"
title: "The Full Claude Desktop Experience on AWS, Google Cloud, and Microsoft Foundry"
article_title: "Claude Desktop on AWS, Google Cloud, and Microsoft Foundry | Claude"
author: "hackerBanana"
captured_at: "2026-06-22T22:33:17Z"
capture_tool: "hn-digest"
hn_id: 48636913
score: 2
comments: 1
posted_at: "2026-06-22T22:06:30Z"
tags:
  - hacker-news
  - translated
---

# The Full Claude Desktop Experience on AWS, Google Cloud, and Microsoft Foundry

- HN: [48636913](https://news.ycombinator.com/item?id=48636913)
- Source: [claude.com](https://claude.com/blog/the-full-claude-desktop-experience-on-aws-google-cloud-and-microsoft-foundry)
- Score: 2
- Comments: 1
- Posted: 2026-06-22T22:06:30Z

## Translation

タイトル: AWS、Google Cloud、Microsoft Foundry での完全なクロード デスクトップ エクスペリエンス
記事のタイトル: AWS、Google Cloud、Microsoft Foundry 上の Claude デスクトップ |クロード
説明: AWS、Google Cloud、Microsoft Foundry 上の推論を使用して、完全な Claude デスクトップ エクスペリエンス (チャット、Claude Cowork、Claude Code) をデプロイします。本日よりベータ版でご利用いただけます。

記事本文:
AWS、Google Cloud、Microsoft Foundry 上の Claude デスクトップ |クロード
-->
クロード製品のご紹介 クロード
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
AWS、Google Cloud、Microsoft Foundry での完全な Claude デスクトップ エクスペリエンス
AWS、Google Cloud、Microsoft Foundry での完全な Claude デスクトップ エクスペリエンス
共有 リンクをコピー https://claude.com/blog/the-full-claude-desktop- experience-on-aws-google-cloud-and-microsoft-foundry
AWS、Google Cloud、Microsoft Foundry を通じて Claude Desktop を使用している組織は、チャット、Claude Cowork、Claude Code などの完全なデスクトップ エクスペリエンスをすべて 1 つのアプリで利用できるようになりました。
IT チームは、複数の製品にわたって推論を独自の環境内に保持し、ユーザーごとの SSO、MDM ポリシー テンプレート、オフライン インストーラー オプション、デバイス上で完全に実行できる M365 コネクタを使用して組織全体に Claude Desktop を展開できるようになりました。
推論は、構成したリージョンのクラウド上で実行され、会話履歴はローカルに保存されます。データ コネクタが到達するエンドポイントと、Anthropic が受信する集約テレメトリを制御します。
組織全体に 1 つの表面
今日まで、AWS、Google Cloud、Microsoft Foundry を通じて Claude Desktop を使用している顧客は、Claude Cowork と Claude Code にのみアクセスできました。現在、1 つの展開ですべての役割がカバーされ、各サーフェスには独自のポリシー キーがあるため、誰がいつ何を取得するかを決定できます。
チャットですぐに答えが得られます。

問題について考えています。クロードは、従業員に任せたい作業を担当します。クロードは、承認されたソースを調査し、デバイス上にすでにあるファイルを操作して、成果物を構築し、完了すると結果が明らかになります。 Claude Code は、ターミナル内で作業せずにエージェント コーディングを希望するエンジニア向けです。
Claude Desktop を組織全体に展開するということは、既存のシステム内で作業することを意味します。
他の仕事用アプリと同様にサインインします。従業員は、IAM Identity Center、Workforce Identity Federation、Microsoft Entra ID、または Okta などの OIDC プロバイダーなど、他のすべてのものに使用するものと同じ職場アカウントを使用します。共有キーをローテーションする必要がなく、エンドユーザーのマシンにクラウド認証情報も必要ありません。
すでに管理している他のアプリと同様に展開します。セットアップ UI からポリシー テンプレートをエクスポートし、Intune、GPO、または Jamf を通じてプッシュします。オフライン インストーラーはエアギャップ環境をカバーします。
誰かが見る前に、それが機能することがわかります。すべてのコネクタをテストし、プロバイダーがサービスを提供するクロード モデルを確認し、接続を検証します。これらすべてをロールアウト前に行います。モデル ガードは、設定が正しく構成されていない場合でも、GovCloud を含めてクロード上でルーティングを継続します。
小規模から始めて、採用の増加に応じて拡大します。チャット、Claude Cowork、および Claude Code にはそれぞれ独自のポリシー キーがあるため、技術者以外のチームにチャットと Claude Cowork、Claude Code のエンジニアリングを提供し、チームが各サーフェスを採用するにつれてアクセスを拡大できます。強制拒否ルールはすべてのタブに適用されます。
クロードを作品のある場所に連れて行きましょう。 Microsoft 365 コネクタにより、Claude は独自の Entra アプリを通じてメールやドキュメントにアクセスできるようになり、テナントの許可リストと GCC High/DoD エンドポイントのベータ サポートが提供されます。最も厳格な常駐要件の場合は、ローカル コネクタを使用すると、デバイスと Microsoft の間で接続が維持されます。
「私たちは既存のクラウド環境を通じてクロード デスクトップを迅速に展開しました。

オンメント — 個別のベンダー契約はありません。当社独自の LLM ゲートウェイにより、大規模なインフラストラクチャの構築を行わずに、1 つのチームが世界中の数百のユーザーに導入できるようになりました。」 - サラン・オー、分析/AI チームリーダー、ハンファソリューションズ
管理者向けの展開ガイドでは、SSO、ポリシー テンプレート、ロールアウト前の検証について説明します。または、アカウント チームにご連絡ください。ロールアウトの計画をお手伝いします。
前へ 前へ 0 / 5 次へ 次へ
クロードとともに構築するチーム向けの製品ニュースとベスト プラクティスをさらに詳しくご覧ください。
1M コンテキストが Opus 4.6 および Sonnet 4.6 で一般利用可能になりました
製品発表 1M コンテキストが Opus 4.6 および Sonnet 4.6 で一般利用可能になりました 1M コンテキストが Opus 4.6 と Sonnet 4.6 で一般利用可能になりました 1M コンテキストが Opus 4.6 と Sonnet 4.6 で一般利用可能になりました 1M コンテキストが Opus 4.6 と Sonnet 4.6 で一般利用可能になりました 2026 年 2 月 5 日 Claude Opus 4.6 で金融を前進
Enterprise AI Claude Opus 4.6 による金融の進歩 Claude Opus 4.6 による金融の進歩 Claude Opus 4.6 による金融の進歩 Claude Opus 4.6 による金融の進歩 2025 年 7 月 24 日 人間チームがクロード コードを使用する方法
Enterprise AI 人間チームがクロード コードを使用する方法 人間チームがクロード コードを使用する方法 人間チームがクロード コードを使用する方法 人間チームがクロード コードを使用する方法 2025 年 10 月 30 日 Brex がクロード コードを使用してコードの品質と生産性を向上させる方法
Enterprise AI Brex が Claude Code でコードの品質と生産性を向上させる方法 Brex が Claude Code でコードの品質と生産性を向上する方法 Brex が Claude Code でコードの品質と生産性を向上する方法 Brex が Claude Code でコードの品質と生産性を向上する方法 Claude を使用して組織の運営方法を変革する
製品のアップデート、ハウツー、コミュニティのスポットライトなど。毎月あなたの受信箱に配信されます。
購読する 購読する あなたの情報を提供してください

毎月の開発者ニュースレターを受け取りたい場合は、電子メール アドレスを入力してください。いつでも購読を解除できます。
ホームページ ホームページ 次へ 次へ ありがとうございます!あなたの提出物は受理されました！おっと！フォームの送信中に問題が発生しました。書き込みボタン テキスト ボタン テキスト 学習ボタン テキスト ボタン テキスト コード ボタン テキスト ボタン テキスト 書き込み 聴衆向けにユニークな声を開発するのを手伝ってください こんにちはクロード!聴衆に合わせたユニークな声を開発するのを手伝ってくれませんか?私からさらに詳しい情報が必要な場合は、すぐに 1 ～ 2 つの重要な質問をしてください。より良い仕事をするために役立つドキュメントをアップロードした方がよいと思われる場合は、お知らせください。このタスクをより適切に達成するのに役立つ場合は、Google ドライブ、ウェブ検索など、アクセスできるツールを使用できます。分析ツールは使用しないでください。返答はフレンドリーで簡潔、会話的なものにしてください。
できるだけ早くタスクを実行してください。意味があればアーティファクトでも構いません。アーティファクトを使用する場合は、どの種類のアーティファクト (インタラクティブ、ビジュアル、チェックリストなど) がこの特定のタスクに最も役立つかを検討してください。ご協力いただきありがとうございます！
文章のスタイルを改善する こんにちはクロード!私の書き方を改善してもらえませんか？私からさらに詳しい情報が必要な場合は、すぐに 1 ～ 2 つの重要な質問をしてください。より良い仕事をするために役立つドキュメントをアップロードした方がよいと思われる場合は、お知らせください。このタスクをより適切に達成するのに役立つ場合は、Google ドライブ、ウェブ検索など、アクセスできるツールを使用できます。分析ツールは使用しないでください。返答はフレンドリーで簡潔、会話的なものにしてください。
できるだけ早くタスクを実行してください。意味があればアーティファクトは素晴らしいものになります。アーティファクトを使用する場合は、どの種類のアーティファクト (インタラクティブ、ビジュアル、チェックリストなど) がこの特定のタスクに最も役立つかを検討してください。ご協力いただきありがとうございます！

創造的なアイデアをブレインストーミングします、こんにちはクロード!創造的なアイデアをブレインストーミングしていただけますか?私からさらに詳しい情報が必要な場合は、すぐに 1 ～ 2 つの重要な質問をしてください。より良い仕事をするために役立つドキュメントをアップロードした方がよいと思われる場合は、お知らせください。このタスクをより適切に達成するのに役立つ場合は、Google ドライブ、ウェブ検索など、アクセスできるツールを使用できます。分析ツールは使用しないでください。返答はフレンドリーで簡潔、会話的なものにしてください。
できるだけ早くタスクを実行してください。意味があればアーティファクトは素晴らしいものになります。アーティファクトを使用する場合は、どの種類のアーティファクト (インタラクティブ、ビジュアル、チェックリストなど) がこの特定のタスクに最も役立つかを検討してください。ご協力いただきありがとうございます！
Learn 複雑なトピックを簡単に説明する こんにちはクロード!複雑なトピックを簡単に説明してもらえますか?私からさらに詳しい情報が必要な場合は、すぐに 1 ～ 2 つの重要な質問をしてください。より良い仕事をするために役立つドキュメントをアップロードした方がよいと思われる場合は、お知らせください。このタスクをより適切に達成するのに役立つ場合は、Google ドライブ、ウェブ検索など、アクセスできるツールを使用できます。分析ツールは使用しないでください。返答はフレンドリーで簡潔、会話的なものにしてください。
できるだけ早くタスクを実行してください。意味があればアーティファクトは素晴らしいものになります。アーティファクトを使用する場合は、どの種類のアーティファクト (インタラクティブ、ビジュアル、チェックリストなど) がこの特定のタスクに最も役立つかを検討してください。ご協力いただきありがとうございます！
これらのアイデアを理解するのを手伝ってください、こんにちはクロード!これらの考えを理解するのを手伝ってもらえますか?私からさらに詳しい情報が必要な場合は、すぐに 1 ～ 2 つの重要な質問をしてください。より良い仕事をするために役立つドキュメントをアップロードした方がよいと思われる場合は、お知らせください。必要に応じて、Google ドライブ、ウェブ検索など、アクセス権のあるツールを使用できます。

このタスクをより適切に達成できるようサポートしてください。分析ツールは使用しないでください。返答はフレンドリーで簡潔、会話的なものにしてください。
できるだけ早くタスクを実行してください。意味があればアーティファクトは素晴らしいものになります。アーティファクトを使用する場合は、どの種類のアーティファクト (インタラクティブ、ビジュアル、チェックリストなど) がこの特定のタスクに最も役立つかを検討してください。ご協力いただきありがとうございます！
試験または面接の準備をする こんにちはクロード!試験や面接の準備をしてもらえますか？私からさらに詳しい情報が必要な場合は、すぐに 1 ～ 2 つの重要な質問をしてください。より良い仕事をするために役立つドキュメントをアップロードした方がよいと思われる場合は、お知らせください。このタスクをより適切に達成するのに役立つ場合は、Google ドライブ、ウェブ検索など、アクセスできるツールを使用できます。分析ツールは使用しないでください。返答はフレンドリーで簡潔、会話的なものにしてください。
できるだけ早くタスクを実行してください。意味があればアーティファクトは素晴らしいものになります。アーティファクトを使用する場合は、どの種類のアーティファクト (インタラクティブ、ビジュアル、チェックリストなど) がこの特定のタスクに最も役立つかを検討してください。ご協力いただきありがとうございます！
コード プログラミングの概念を説明する こんにちはクロード!プログラミングの概念について説明していただけますか?私からさらに詳しい情報が必要な場合は、すぐに 1 ～ 2 つの重要な質問をしてください。より良い仕事をするために役立つドキュメントをアップロードした方がよいと思われる場合は、お知らせください。このタスクをより適切に達成するのに役立つ場合は、Google ドライブ、ウェブ検索など、アクセスできるツールを使用できます。分析ツールは使用しないでください。返答はフレンドリーで簡潔、会話的なものにしてください。
できるだけ早くタスクを実行してください。意味があればアーティファクトは素晴らしいものになります。アーティファクトを使用する場合は、どの種類のアーティファクト (インタラクティブ、ビジュアル、チェックリストなど) がこの仕様に最も役立つかを検討します。

重要なタスク。ご協力いただきありがとうございます！
私のコードを見てヒントを教えてください、こんにちはクロード!私のコードを見て、ヒントを教えていただけますか?私からさらに詳しい情報が必要な場合は、すぐに 1 ～ 2 つの重要な質問をしてください。より良い仕事をするために役立つドキュメントをアップロードした方がよいと思われる場合は、お知らせください。このタスクをより適切に達成するのに役立つ場合は、Google ドライブ、ウェブ検索など、アクセスできるツールを使用できます。分析ツールは使用しないでください。返答はフレンドリーで簡潔、会話的なものにしてください。
できるだけ早くタスクを実行してください。意味があればアーティファクトは素晴らしいものになります。アーティファクトを使用する場合は、どの種類のアーティファクト (インタラクティブ、ビジュアル、チェックリストなど) がこの特定のタスクに最も役立つかを検討してください。ご協力いただきありがとうございます！
バイブコードを使ってこんにちはクロード！一緒にバイブコードを作ってくれませんか？私からさらに詳しい情報が必要な場合は、すぐに 1 ～ 2 つの重要な質問をしてください。より良い仕事をするために役立つドキュメントをアップロードした方がよいと思われる場合は、お知らせください。このタスクをより適切に達成するのに役立つ場合は、Google ドライブ、ウェブ検索など、アクセスできるツールを使用できます。分析ツールは使用しないでください。返答はフレンドリーで簡潔、会話的なものにしてください。
できるだけ早くタスクを実行してください。アーティファクトがあれば素晴らしいです。

[切り捨てられた]

## Original Extract

Deploy the full Claude desktop experience - chat, Claude Cowork, and Claude Code - using inference on AWS, Google Cloud and Microsoft Foundry. Available today in beta.

Claude Desktop on AWS, Google Cloud, and Microsoft Foundry | Claude
-->
Meet Claude Products Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
The full Claude Desktop experience on AWS, Google Cloud, and Microsoft Foundry
The full Claude Desktop experience on AWS, Google Cloud, and Microsoft Foundry
Share Copy link https://claude.com/blog/the-full-claude-desktop-experience-on-aws-google-cloud-and-microsoft-foundry
Organizations that use Claude Desktop through AWS, Google Cloud, and Microsoft Foundry now get the full Desktop experience — chat, Claude Cowork, and Claude Code, all in one app.
Now IT teams can keep inference inside their own environment across products, and deploy Claude Desktop organization-wide with per-user SSO, MDM policy templates, an offline installer option, and an M365 connector that can run entirely on the device.
Inference runs on your cloud in the regions you configure and conversation history is stored locally. You control the endpoints data connectors reach and the aggregated telemetry Anthropic receives.
One surface for the entire organization
Until today, customers using Claude Desktop through AWS, Google Cloud, and Microsoft Foundry only had access to Claude Cowork and Claude Code. Now, one deployment covers every role, and each surface has its own policy key, so you decide who gets what, and when.
Chat for quick answers and thinking through a problem. Claude Cowork for the work your people would rather hand off: Claude researches across approved sources, works with the files already on the device and builds the deliverable, surfacing results when it’s done. Claude Code for engineers who want agentic coding without living in a terminal.
Deploying Claude Desktop organization-wide means working within the systems you already have.
Sign in like any work app. Employees use the same work account they use for everything else: IAM Identity Center, Workforce Identity Federation, Microsoft Entra ID, or any OIDC provider like Okta. No shared keys to rotate, no cloud credentials on end-user machines.
Deploy like any app you already manage. Export policy templates from the setup UI and push them through Intune, GPO, or Jamf. An offline installer covers air-gapped environments.
Know it works before anyone sees it. Test every connector, confirm which Claude models your provider serves, and verify the connection, all before rollout. A model guard keeps routing on Claude, including in GovCloud, even if a setting is misconfigured.
Start small, expand as adoption grows. Chat, Claude Cowork, and Claude Code each have their own policy key, so you can give non-technical teams chat and Claude Cowork, engineering Claude Code, and then broaden access as teams adopt each surface. Your hard-deny rules apply across every tab.
Bring Claude to where the work lives. A Microsoft 365 connector gives Claude access to mail and documents through your own Entra app, with tenant allowlisting and beta support for GCC High/DoD endpoints. For the strictest residency requirements, use our local connector, and the connection stays between the device and Microsoft.
"We rolled out Claude Desktop fast through our existing cloud environment — no separate vendor contract. Our own LLM Gateway let one team deploy it to hundreds of users worldwide, with no heavy infrastructure build-out." - Sarang Oh, Analytics/AI Team Leader, Hanwha Solutions
For admins, the deployment guide walks through SSO, policy templates, and pre-rollout validation. Or contact your account team and we'll help you plan the rollout.
Prev Prev 0 / 5 Next Next eBook
Explore more product news and best practices for teams building with Claude.
1M context is now generally available for Opus 4.6 and Sonnet 4.6
Product announcements 1M context is now generally available for Opus 4.6 and Sonnet 4.6 1M context is now generally available for Opus 4.6 and Sonnet 4.6 1M context is now generally available for Opus 4.6 and Sonnet 4.6 1M context is now generally available for Opus 4.6 and Sonnet 4.6 Feb 5, 2026 Advancing finance with Claude Opus 4.6
Enterprise AI Advancing finance with Claude Opus 4.6 Advancing finance with Claude Opus 4.6 Advancing finance with Claude Opus 4.6 Advancing finance with Claude Opus 4.6 Jul 24, 2025 How Anthropic teams use Claude Code
Enterprise AI How Anthropic teams use Claude Code How Anthropic teams use Claude Code How Anthropic teams use Claude Code How Anthropic teams use Claude Code Oct 30, 2025 How Brex improves code quality and productivity with Claude Code
Enterprise AI How Brex improves code quality and productivity with Claude Code How Brex improves code quality and productivity with Claude Code How Brex improves code quality and productivity with Claude Code How Brex improves code quality and productivity with Claude Code Transform how your organization operates with Claude
Product updates, how-tos, community spotlights, and more. Delivered monthly to your inbox.
Subscribe Subscribe Please provide your email address if you'd like to receive our monthly developer newsletter. You can unsubscribe at any time.
Homepage Homepage Next Next Thank you! Your submission has been received! Oops! Something went wrong while submitting the form. Write Button Text Button Text Learn Button Text Button Text Code Button Text Button Text Write Help me develop a unique voice for an audience Hi Claude! Could you help me develop a unique voice for an audience? If you need more information from me, ask me 1-2 key questions right away. If you think I should upload any documents that would help you do a better job, let me know. You can use the tools you have access to— like Google Drive, web search, etc.—if they’ll help you better accomplish this task. Do not use analysis tool. Please keep your responses friendly, brief and conversational.
Please execute the task as soon as you can—an artifact would be great if it makes sense. If using an artifact, consider what kind of artifact (interactive, visual, checklist, etc.) might be most helpful for this specific task. Thanks for your help!
Improve my writing style Hi Claude! Could you improve my writing style? If you need more information from me, ask me 1-2 key questions right away. If you think I should upload any documents that would help you do a better job, let me know. You can use the tools you have access to— like Google Drive, web search, etc.—if they’ll help you better accomplish this task. Do not use analysis tool. Please keep your responses friendly, brief and conversational.
Please execute the task as soon as you can—an artifact would be great if it makes sense. If using an artifact, consider what kind of artifact (interactive, visual, checklist, etc.) might be most helpful for this specific task. Thanks for your help!
Brainstorm creative ideas Hi Claude! Could you brainstorm creative ideas? If you need more information from me, ask me 1-2 key questions right away. If you think I should upload any documents that would help you do a better job, let me know. You can use the tools you have access to— like Google Drive, web search, etc.—if they’ll help you better accomplish this task. Do not use analysis tool. Please keep your responses friendly, brief and conversational.
Please execute the task as soon as you can—an artifact would be great if it makes sense. If using an artifact, consider what kind of artifact (interactive, visual, checklist, etc.) might be most helpful for this specific task. Thanks for your help!
Learn Explain a complex topic simply Hi Claude! Could you explain a complex topic simply? If you need more information from me, ask me 1-2 key questions right away. If you think I should upload any documents that would help you do a better job, let me know. You can use the tools you have access to— like Google Drive, web search, etc.—if they’ll help you better accomplish this task. Do not use analysis tool. Please keep your responses friendly, brief and conversational.
Please execute the task as soon as you can—an artifact would be great if it makes sense. If using an artifact, consider what kind of artifact (interactive, visual, checklist, etc.) might be most helpful for this specific task. Thanks for your help!
Help me make sense of these ideas Hi Claude! Could you help me make sense of these ideas? If you need more information from me, ask me 1-2 key questions right away. If you think I should upload any documents that would help you do a better job, let me know. You can use the tools you have access to— like Google Drive, web search, etc.—if they’ll help you better accomplish this task. Do not use analysis tool. Please keep your responses friendly, brief and conversational.
Please execute the task as soon as you can—an artifact would be great if it makes sense. If using an artifact, consider what kind of artifact (interactive, visual, checklist, etc.) might be most helpful for this specific task. Thanks for your help!
Prepare for an exam or interview Hi Claude! Could you prepare for an exam or interview? If you need more information from me, ask me 1-2 key questions right away. If you think I should upload any documents that would help you do a better job, let me know. You can use the tools you have access to— like Google Drive, web search, etc.—if they’ll help you better accomplish this task. Do not use analysis tool. Please keep your responses friendly, brief and conversational.
Please execute the task as soon as you can—an artifact would be great if it makes sense. If using an artifact, consider what kind of artifact (interactive, visual, checklist, etc.) might be most helpful for this specific task. Thanks for your help!
Code Explain a programming concept Hi Claude! Could you explain a programming concept? If you need more information from me, ask me 1-2 key questions right away. If you think I should upload any documents that would help you do a better job, let me know. You can use the tools you have access to— like Google Drive, web search, etc.—if they’ll help you better accomplish this task. Do not use analysis tool. Please keep your responses friendly, brief and conversational.
Please execute the task as soon as you can—an artifact would be great if it makes sense. If using an artifact, consider what kind of artifact (interactive, visual, checklist, etc.) might be most helpful for this specific task. Thanks for your help!
Look over my code and give me tips Hi Claude! Could you look over my code and give me tips? If you need more information from me, ask me 1-2 key questions right away. If you think I should upload any documents that would help you do a better job, let me know. You can use the tools you have access to— like Google Drive, web search, etc.—if they’ll help you better accomplish this task. Do not use analysis tool. Please keep your responses friendly, brief and conversational.
Please execute the task as soon as you can—an artifact would be great if it makes sense. If using an artifact, consider what kind of artifact (interactive, visual, checklist, etc.) might be most helpful for this specific task. Thanks for your help!
Vibe code with me Hi Claude! Could you vibe code with me? If you need more information from me, ask me 1-2 key questions right away. If you think I should upload any documents that would help you do a better job, let me know. You can use the tools you have access to— like Google Drive, web search, etc.—if they’ll help you better accomplish this task. Do not use analysis tool. Please keep your responses friendly, brief and conversational.
Please execute the task as soon as you can—an artifact would be great if it

[truncated]
