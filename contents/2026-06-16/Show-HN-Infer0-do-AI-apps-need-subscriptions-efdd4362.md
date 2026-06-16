---
source: "https://infer0.com/"
hn_url: "https://news.ycombinator.com/item?id=48556523"
title: "Show HN: Infer0 – do AI apps need subscriptions?"
article_title: "Home | infer0"
author: "sumolessons"
captured_at: "2026-06-16T16:38:02Z"
capture_tool: "hn-digest"
hn_id: 48556523
score: 5
comments: 0
posted_at: "2026-06-16T15:12:22Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Infer0 – do AI apps need subscriptions?

- HN: [48556523](https://news.ycombinator.com/item?id=48556523)
- Source: [infer0.com](https://infer0.com/)
- Score: 5
- Comments: 0
- Posted: 2026-06-16T15:12:22Z

## Translation

タイトル: Show HN: Infer0 – AI アプリにはサブスクリプションが必要ですか?
記事のタイトル: ホーム |推測0
説明: あなた自身の知性をもたらしてください。 infer0 を使用すると、エンドユーザーが AI プロバイダー キーを接続してアプリを認証できるようになり、開発者は一度統合すれば任意のプロバイダーにルーティングできます。
HN text: AI サイドプロジェクトに関して私を悩ませていることの 1 つは、推論コストです。従来のソフトウェアでは、発売が成功すれば通常、より高い利益が得られます。しかし、AI 製品の場合、成功すると予想外に多額の費用がかかる可能性があります。このため、私はより安価で機能の低いモデルを選択するようになり、特定のアイデアを検討することさえ躊躇するようになりました。すべてのサイドプロジェクトがさらに月額 20 ドルのサブスクリプションになることは望ましくありませんが、推論コストの補助を積極的に行う VC 支援の企業と競争することもできません。そこで私は次のアイデアを思いつきました。ユーザーが自分の推論に対して単純に料金を支払ったらどうなるでしょうか?これはローカルに設定された API キーを通じて一部のアプリですでに発生していますが、モデルを拡張できるでしょうか?ユーザーが独自の AI アカウントを持っている場合、開発者は変動する推論コストを負担することなく、AI を活用した製品を構築できます。サブスクリプションビジネスにするべきではない AI アプリケーションがどれだけあるでしょうか?課題は、開発者がユーザーの API キーを扱いたくないこと、ユーザーが試行するすべてのアプリに API キーを配布したくないこと、そして推論コストを通過させるためだけに支払い方法を収集するという煩わしさを誰も望んでいないことです。これが私の最新のサイド プロジェクト、infer0.com の裏話です。これは AI 推論の SSO に似ています。ユーザーは AI プロバイダーに一度接続すると、アプリは認証トークンを使用して infer0 を通じて推論を要求します。開発者は API キーを管理したり、モデルのコストを自分で支払ったりする必要はありませんが、ユーザーは複数のアプリケーションにわたって同じ AI アカウントを使用できます。これはひどいアイデアかもしれません。なぜなら、誰も信用しないでしょうし、信用する人がいると私は確信しているからです。

ユーザー API 資格情報の処理に関するリスクについて、私は十分に理解していませんでした。しかし、それを構築する必要があると感じました。それでは、大まかな最初のパスを示します。

記事本文:
0 " />
0 " />
0を推測します
クイックスタート
価格設定
よくある質問
ドキュメント
サインイン
サインイン
あなた自身の知性をもたらしてください。
infer0 を使用すると、エンドユーザーが AI プロバイダー キーを接続して認証できるようになります。
それらを使用するアプリ。開発者はAPI形式を使用して一度統合する
それは彼らのスタックに適合します。 infer0 はそれぞれのプロバイダーにルートします
ユーザーはコード行を変更せずに選択できます。
開発者設定で OAuth アプリを作成し、クライアント ID とシークレットを取得します。 infer0 OAuth フローをアプリに追加します。サポートされている API 形式を使用します。 infer0 はプロバイダー間を自動的に変換します。
各ユーザーは、OpenAI、Anthropic、または Google API キーを接続します。キーは AES-256-GCM で暗号化され、安全に保管されます。ユーザーは、プロバイダーごとに 1 日の支出制限を設定して、コストを管理できます。
OAuth フロー中に、各ユーザーはアプリが使用できる AI プロバイダーを選択します。後で [承認] ページから、割り当てられたプロバイダーを変更したり、承認ごとの 1 日の使用量制限を設定したりできます。
アプリは、SDK と一致するエンドポイントにユーザーのアクセス トークンを渡します。 infer0 は AI プロバイダーを検索し、支出制限をチェックし、リクエストを転送し、レスポンスを SDK が期待する形式に変換します。
コードベースに既にある OpenAI または Anthropic SDK を使用します。 infer0 は、チャット完了、メッセージ、および応答の形式をサポートします。プロバイダー固有のルーティング ロジックは必要ありません。
プロバイダー API キーは AES-256-GCM で暗号化され、アプリに公開されることはありません。保管したり、扱ったり、見ることさえしません。
各ユーザーは自分の AI プロバイダーの料金を支払います。アプリが請求、クレジット、API コストを処理することはありません。料金は自分のホスティング料金のみを支払います。
OpenAI、Anthropic、または Google API キーを接続します。認可されたアプリごとに使用するプロバイダーを切り替えます。予算内に収まるように、プロバイダーごとに 1 日の支出制限を設定します。
承認された各アプリには、独自のプロバイダーの割り当てと 1 日の制限が与えられます。

2番目の制限。これらは、承認ページからいつでも変更できます。ワンクリックでアプリを一時停止または取り消します。
いつでもプロバイダーを追加、削除、または切り替えることができます。承認ごとおよびプロバイダーごとに 1 日の使用量制限を設定します。ダッシュボードからアプリのアクセスを即座に取り消します。
infer0 は開発初期段階にあります。それを使って何かを構築している場合、
それについて教えてください
あなたのプロジェクトをここで紹介します。
infer0 は、プロンプトや完了の内容をログに記録したり保存したりしません。レート制限のメタデータは保持されますが、メッセージ本文がディスクに書き込まれることはありません。
プロバイダー API キーは、保存前に AES-256-GCM で暗号化されます。トークンはハッシュ化されます。暗号化キーはデータベースとは別に管理されます。
暗号化されたデータは、infer0 スタッフが読み取ることはできません。暗号化キーはデータベースとは別に管理され、アプリケーションに公開されることはありません。
ユーザーは、「認証」ページから任意のアプリのアクセスを取り消すことができます。取り消すと、関連付けられているすべてのトークンが即座に無効になり、それ以降のリクエストがブロックされます。
アカウント プロファイル、暗号化されたプロバイダー構成、および OAuth 認証レコードは保持されます。ユーザーはいつでもプロバイダーを削除し、承認を取り消すことができます。
AI プロバイダーのユーザー API キーは有効なままであり、影響を受けません。サービスが再開されるまで、アプリの infer0 リクエストは失敗します。アプリ内でこれを適切に処理してください。

## Original Extract

Bring your own intelligence. infer0 lets end users connect their AI provider keys and authorize apps, while developers integrate once and route to any provider.

One thing that’s been bothering me about AI side projects is inference costs. With traditional software, a successful launch usually means higher profits. But with AI products, success can mean unexpectedly large bills. This has pushed me toward cheaper, less capable models and made me hesitate to even explore certain ideas. I don’t want every side project to become another $20/month subscription, but I also can’t compete with VC-backed companies willing to subsidize inference costs. Then I had this idea: what if users simply paid for their own inference? This already happens in some apps through locally configured API keys, but could the model be extended? If users bring their own AI account, developers can build AI-powered products without taking on variable inference costs. How many AI applications shouldn’t be subscription businesses at all? The challenge is that developers don’t want to handle user API keys, users don’t want to hand them out to every app they try, and nobody wants the friction of collecting payment methods just to pass through inference costs. That’s the backstory to my latest side project, infer0.com. It's a bit like SSO for AI inference. Users connect their AI providers once, and apps use auth tokens to request inference through infer0. Developers don’t manage API keys or pay model costs themselves, while users can bring the same AI accounts across multiple applications. This may be a terrible idea, both because nobody will trust it and because I’m sure there are risks around handling user API credentials that I haven’t fully appreciated. But I felt the need to build it. So here’s a rough first pass.

0 " />
0 " />
infer 0
Quickstart
Pricing
FAQ
Docs
Sign in
Sign in
Bring your own intelligence.
infer0 lets end users connect their AI provider keys and authorize
apps to use them. Developers integrate once using the API format
that fits their stack. infer0 routes to whichever provider each
user chooses without changing a line of code.
Create an OAuth app in Developer Settings to get a client ID and secret. Add the infer0 OAuth flow to your app. Use any supported API format. infer0 translates between providers automatically.
Each user connects their OpenAI, Anthropic, or Google API key. Keys are encrypted with AES-256-GCM and stored securely. Users can set a per-provider daily spend limit to control costs.
During the OAuth flow, each user chooses which AI Provider the app may use. They can change the assigned provider or set a per-authorization daily spend limit later from their Authorizations page.
Your app passes the user's access token to the endpoint that matches your SDK. infer0 looks up their AI Provider, checks spend limits, forwards the request, and translates the response back into the format your SDK expects.
Use the OpenAI or Anthropic SDK your codebase already has. infer0 supports Chat Completions, Messages, and Responses formats. No provider-specific routing logic needed.
Provider API keys are encrypted with AES-256-GCM and never exposed to your app. You don't store, handle, or even see them.
Each user pays their own AI provider bills. Your app never handles billing, credits, or API costs. You only pay for your own hosting.
Connect your OpenAI, Anthropic, or Google API key. Switch which provider each authorized app uses. Set daily spend limits per provider to stay within budget.
Each authorized app gets its own provider assignment and daily spend limit. Change these anytime from your Authorizations page. Pause or revoke an app in one click.
Add, remove, or switch providers anytime. Set daily spend limits per authorization and per provider. Revoke any app's access instantly from your dashboard.
infer0 is in early development. If you are building something with it,
tell us about it
and we will feature your project here.
infer0 does not log or store prompt or completion content. Metadata for rate limiting is retained, but message bodies are never written to disk.
Provider API keys are encrypted with AES-256-GCM before storage. Tokens are hashed. Encryption keys are managed separately from the database.
Encrypted data cannot be read by infer0 staff. Encryption keys are managed separately from the database and never exposed to the application.
Users can revoke any app's access from their Authorizations page. Revoking immediately invalidates all associated tokens and blocks further requests.
Account profiles, encrypted provider configs, and OAuth authorization records are retained. Users can delete providers and revoke authorizations at any time.
User API keys with their AI provider remain valid and unaffected. App requests to infer0 will fail until service resumes. Handle this gracefully in your app.
