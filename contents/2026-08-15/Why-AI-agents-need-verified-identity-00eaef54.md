---
source: "https://chrisbergeron.com/2026/08/11/why_ai_agents_need_verified_identity/"
hn_url: "https://news.ycombinator.com/item?id=49314608"
title: "Why AI agents need verified identity"
article_title: "Why AI agents need verified identity - Chris Bergerons Tech Blog"
author: "speckx"
captured_at: "2026-08-15T22:10:37Z"
capture_tool: "hn-digest"
hn_id: 49314608
score: 2
comments: 0
posted_at: "2026-08-15T21:44:48Z"
tags:
  - hacker-news
  - translated
---

# Why AI agents need verified identity

- HN: [49314608](https://news.ycombinator.com/item?id=49314608)
- Source: [chrisbergeron.com](https://chrisbergeron.com/2026/08/11/why_ai_agents_need_verified_identity/)
- Score: 2
- Comments: 0
- Posted: 2026-08-15T21:44:48Z

## Translation

タイトル: AI エージェントに検証済みの ID が必要な理由
記事のタイトル: AI エージェントに検証済みの ID が必要な理由 - Chris Bergerons Tech Blog
説明: ハウツー記事、プロジェクト、例を掲載したオンライン テクノロジー ジャーナル。

記事本文:
AI エージェントに検証済みの ID が必要な理由 - Chris Bergerons Tech Blog
ホーム アーカイブ カテゴリー タグ 概要 投稿日 2026 年 8 月 11 日 更新日 2026 年 8 月 11 日 AI エージェントに検証済みの ID が必要な理由
インターネット上で自分の名前を永久に所有するために 9 ドルを支払いました
私はインターネット上で自分の名前を所有するために 9 ドルを支払いました。レンタルしないでください。それを所有してください。暗号化して、移植可能に、永遠に。今日、私は同じことを可能にするオンランプを立ち上げます。
2 分間お付き合いください。なぜ、あなたがこれを読んでいるドメイン名よりも 9 ドルのハンドルの方が重要なのかを説明します。
20 年前、まだ「インフォテインメント」という言葉が存在する前に、私は自動車用のタッチスクリーン インダッシュ コンピューターを作りました。商品カテゴリーを作りました。これは当時としては斬新で物議を醸したものでしたが、あるいは単にスラッシュドットでの悪口だったのかもしれませんが、現在ではすべての新車に同梱されています。当時、車と最先端のテクノロジーを組み合わせたものではありませんでした。私がそうするまでは、非常に公に。
私は AI と Agentic Web に対しても同じアプローチをとっています。不足している部分を提供することでタイムラインを前進させます。未来を見る準備はできていますか?
パズルの欠けているピース
AI エージェントは現在、カレンダー作成、購入、調査、通信など、あらゆるワークフローに登場しています。同様に、アドレス指定可能である必要があります。インターネットがオンラインになったとき、次の 1 つの疑問が生まれ続けました。
Web アドレスとは何ですか?電話帳のどこにも見つかりません。
その Web アドレスが URL となり、それ以来、私たちがインターネットを検索し、読み、サーフィンする方法になっています。このブログには https://chrisbergeron.com があります。
Agentic Web はインターネットの次のバージョンであり、私たちはまさにそれが展開し始める瞬間に立っています。未来のアドレスは URL ではありません。それは Did:web です。そして昔と同じように、そこに行くには、新しい情報スーパーハイウェイへの入口ランプを利用する必要があります。そこで

username.md が登場します。9 ドルで、あなたが所有し、規制し、制御する正規のアドレスを取得できます。
分散型識別子 (DID) は、世界的に一意で可用性が高く、暗号的に検証可能なデジタル識別子です。これらは通常、個人、組織、データ モデル、または任意の抽象エンティティを指すことができる URI (Uniform Resource Identifier) として記述されます。
DID と、電子メール アドレスやユーザー アカウントなどの従来の識別子の違いは単純ですが奥が深いです。DID はどのサービス プロバイダーによっても所有されないのです。これはプラットフォーム間で機能し、ベンダー ロックインを防ぎます。
DID は W3C 標準です。プライバシーを保護し、同意とデータのポータビリティを可能にして、ユーザーが制御できるようにします。パスワードの代わりに、公開鍵と秘密鍵のペアを使用します。これは、あらゆる軸においてより強力なセキュリティ モデルです。つまり、DID は Web の分散公開キー インフラストラクチャ (DPKI) の基礎を形成します。
「分散化されているのに、なぜそれを販売するのですか？」
公正な質問です。 DID が分散化されている場合、DID を提供すると私が中央の権威者になるのではありませんか?
短期的にはその通りです、大声で言っても構いません。私は、今日では仕様書内に閉じ込められている機能への入り口を構築しました。すでに所有しているドメインで独自の DID ドキュメントをホストしたい場合は、ホストすることもできます。実際の例については、このブログの上部にあるメタデータを参照してください。それが私が始めた方法です。まず chrisbergeron.com でメタ タグを作成し、次に schema.org 準拠のプロファイルを作成しました。
私はその機能をすべての人に解放したいと考えており、username.md は単にそのための手段です。その過程で、私はエージェント Web 用の耐久性のある構成要素を出荷し続けるつもりです。
企業はこのようなツールを何年も使用してきました。普通の人は何も得られませんでした。今まで。
マーク・ザッカーバーグは最近、「未来はイブのためにある」というタイトルの記事を発表した。

リョネ。」もう一度読んでください。未来がすべての人のためのものであるなら、なぜそれを発表するのでしょうか？なぜなら、メタはあなたの入口になりたいからです。 Googleも同様です。マイクロソフトも同様です。彼らは、あなたが自分たちとだけ入る壁に囲まれた庭園内で AI エクスペリエンスを提供したいと考えています。
画面はもうご存知ですよね。 「Googleでサインインしてください。 Facebook でサインインしてください。」
それをクリックすると、サービス利用規約に同意したことになります。あなたは、AI との将来の関係全体を形作るためのライセンスを彼らに渡しました。それで彼らを信頼しますか？
エージェントの時代には、ログイン ボタンでは答えられない疑問が生じているためです。
AI エージェントは実際にあなたについて何を知っているのでしょうか?
あるエージェントが別のエージェントと話すとき、誰が何に署名したか?
何か問題が起こったとき、誰が責任を取るのでしょうか？
誰がエージェントを管理し、誰がコンプライアンスを証明するのでしょうか?
そもそも、エージェントが私たちに代わって行動することをどうやって信頼できるのでしょうか?
これらすべてに対する答えは、本人確認です。それが username.md によって提供されるものです。
2026 年における「本人確認済み」とは実際に何を意味するのか
それは次の 3 つのプロパティになります。
来歴 — メッセージはおそらく秘密鍵を保持するエンティティから送信されたものと考えられます。
アンカー — そのキーは安定した人間が読める名前 (ランダムな GUID ではなく chris.username.md ) にバインドされます。
移植性 — この名前は単一のプラットフォームに属しません。 DID を別のホストに移動しても、エージェントは引き続き同じ ID を解決できます。
ワイヤー上では次のようになります。 OIDC 検出応答:
1
$カール "https://login.username.md/.well-known/openid-configuration" 2>/dev/null | jq 。
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
{
"発行者" : "https://login.username.md" ,
"authorization_endpoint" : "https://login.username.md/oauth/v2/authorize" ,
"token_endpoint" : "https://login.username.md/oauth/v2/token" ,
"introspection_endpoint" : "https://login.username.md/oauth/v2/introspect"

、
"userinfo_endpoint" : "https://login.username.md/oidc/v1/userinfo" ,
"revocation_endpoint" : "https://login.username.md/oauth/v2/revoke" ,
"end_session_endpoint" : "https://login.username.md/oidc/v1/end_session" ,
"device_authorization_endpoint" : "https://login.username.md/oauth/v2/device_authorization" ,
"jwks_uri" : "https://login.username.md/oauth/v2/keys" ,
"スコープ_サポート" : [
"openid" 、
「プロフィール」 、
「メール」、
「電話」、
「住所」 、
「オフラインアクセス」
】
...切り詰められた
}
また、ハンドルが返すすべての応答には RFC 9421 署名が付けられます。
1
2
3
コンテンツダイジェスト: sha-256=:<base64(sha256(body))>:
署名入力: sig1=("content-digest" "@authority");created=...;keyid="platform-ed25519";alg="ed25519"
署名: sig1=:<base64(ed25519 署名)>:
ハンドルからのすべての応答は暗号署名されており、エージェントは約 4 行の Go で検証できます。詳細が必要な場合は、開発者ドキュメントがここにあります。
「Google でサインイン」がこれではない理由
Google と GitHub はアカウントの存在を証明できます。彼らはあなたがあなたの身元を所有していることを証明することはできません。彼らがあなたに渡す OAuth トークンは不透明です。誰でもそれを発行した人を確認できますが、あなた自身の署名キーを与えることはありません。 what@gmail.com からメッセージを送信しても、その背後にあなたが制御するキーはありません。これは Google のインフラストラクチャであり、キャプティブ ユーザーは善良なユーザーであるため、Google がキーを提供することはありません。彼らがあなたのアカウントを閉鎖した場合、あなたは単にオンライン上で存在しなくなるだけです。一部の報告によると、それは毎日何百人もの人々に起こっています。
OAuth は、人間がアプリにログインするために構築されました。エージェントには、OAuth に求められていなかったものが必要です。
ハンドルを購入すると、次の 3 つの特典が得られます。
あなたが公開する検証可能な身元。 chris.username.md は、署名された DID ドキュメント、OIDC 検出エンドポイント、ATProto ハンドル、およびプレーンテキストのエージェントに解決されます。

-読みやすいプロフィール。すべて9ドルです。
エージェントによるネイティブの発見可能性。あなたのハンドルにある署名済み SKILL.md は、あなたが提供する内容と連絡方法を他のエージェントに伝えます。彼らはそれを検証し、信頼することができます。
私よりも長生きする所有権。支払いは1回。サブスクリプション、ガス料金、財布は必要ありません。会社が消滅しても、ハンドルはパブリック DID レジストリ スナップショットに残ります。
なぜ今、なぜ私、なぜ username.md
私はヘルプ デスクから始まり、システム管理を経て、プリンシパル サイト信頼性エンジニアとして現在に至るまで、人生のほとんどをテクノロジーに費やしてきました。長年の読者は、DashPC と Dashboard Linux、そして今日のすべての車に搭載されているのと同じ種類のインフォテインメント システムを私がどのようにして公の場で構築したかを知っています。それはグリーンフィールドであり、初期のものであり、自動車とソフトウェアが出会うことで何が可能になるかを世界に示しました。
私は再び同じレンズを通して見ており、AI パズルの欠けているピースをあなたに渡します。私は20年間「早すぎる」と感じてきましたが、それと和解しました。 username.md が次の Cloudflare であるとは主張しません。でも、そうかもしれない。
新しい AI ゴールド ラッシュに今すぐ賭けることができます。ユーザー名を $9 で予約してください。それがすべてです。エージェント用に 1 つ購入してください。良い名前が無くなる前に、家族や友人のために購入してください。
9 ドルを支払えば、インターネット上で自分の名前を永久に所有できるようになります。暗号的に。持ち運び可能。
username.md は完全にオープンスタンダードに基づいて構築されており、それが重要な点です。 Did:web 、OIDC (RFC 6749 および OIDC Discovery)、RFC 9421 署名付き応答、および AT プロトコルなど、すべての仕様が公開およびリンクされています。このプロジェクトは、私の会社である The Holding Company と、私が 2002 年から運営している技術ブログ chrisbergeron.com によって支援されています。
私はインターネット上で自分の名前を永久に所有するために 9 ドルを支払いました。あなたもできるようになりました: username.md
AI エージェントに検証済み ID が必要な理由

ティティ
https://chrisbergeron.com/2026/08/11/why_ai_agents_need_verified_identity/
この記事が気に入りましたか?で私をサポートしてください
さらに素晴らしいコンテンツを購読してください
持株会社holdingco.com
AI エージェントに検証済みの ID が必要な理由
Raspberry Pi での Consul systemd のトラブルシューティング
Hashicorp Vault TOTP を使用した動的 SSH ポート
自分自身を所有する - あなたはあなたのブランドです
© 2026 Chris Bergeron Powered by Hexo & Icarus

## Original Extract

An online technology journal with how-to articles, projects and examples.

Why AI agents need verified identity - Chris Bergerons Tech Blog
Home Archives Categories Tags About Posted 08-11-2026 Updated 08-11-2026 Why AI agents need verified identity
I paid $9 to own my name on the internet forever
I paid $9 to own my name on the internet. Not rent it. Own it. Cryptographically, portably, forever. Today I’m launching the on-ramp that lets you do the same:
Stick with me for two minutes and I’ll show you why a $9 handle is about to matter more than the domain name you’re reading this on.
Twenty years ago I built a touchscreen in-dash computer for cars, before the word “infotainment” existed. I created a product category. It was novel and controversial for the time, or maybe that was just the curmudgeons on Slashdot, but today every new car ships with one. Back then, cars and bleeding-edge tech simply weren’t a thing people combined. Until I did, very publicly.
I’m taking that same approach to AI and the Agentic Web: advance the timeline by shipping the missing piece. Are you ready to see into the future?
The missing piece of the puzzle
AI agents are showing up in every workflow now: calendaring, purchasing, research, comms. As they do, they need to be addressable. When the internet came online, one question kept coming up:
What is a web address? I can’t find it anywhere in the phonebook.
That web address became the URL, and it’s how we find, read, and surf the internet ever since. This blog has one: https://chrisbergeron.com .
The Agentic Web is the next iteration of the internet, and we’re standing right at the moment it begins to unfold. The address of the future isn’t a URL. It’s a did:web . And much like the old days, getting there means taking an on-ramp to the new information superhighway. That’s where username.md comes in. For $9, you get a canonical address that you own, you regulate, and you control.
Decentralized identifiers (DIDs) are globally unique, highly available, cryptographically verifiable digital identifiers. They’re usually written as a Uniform Resource Identifier (URI) that can point to a person, an organization, a data model, or any abstract entity.
The difference between a DID and a traditional identifier like an email address or a user account is simple but profound: a DID isn’t owned by any service provider. It works across platforms, and it prevents vendor lock-in.
DIDs are a W3C standard. They preserve privacy, enable consent and data portability, and put the user in control. Instead of passwords, they use public/private key pairs, which is a stronger security model on every axis. In short, DIDs form the basis of a Decentralized Public Key Infrastructure (DPKI) for the web.
“Decentralized, so why are you selling it?”
Fair question. If a DID is decentralized, doesn’t offering one make me a central authority?
In the near term, yes, and I’m fine saying so out loud. I built an on-ramp to a capability that today is locked inside a specification document . If you’d rather host your own DID documents on a domain you already own, you can. Look at the metadata at the top of this blog for a working example. That’s how I started: meta tags on chrisbergeron.com first, then a schema.org-compliant profile.
I want to unlock that capability for everyone, and username.md is simply the means to that end. Along the way I plan to keep shipping durable building blocks for the agentic web.
The enterprise has had tools like this for years. Regular people got nothing. Until now.
Mark Zuckerberg recently published a piece titled “The Future is for Everyone.” Read that again. If the future is for everyone, why announce it? Because Meta wants to be your on-ramp. So does Google. So does Microsoft. They want to provide your AI experience inside a walled garden that you enter exclusively with them .
You already know the screen. “Sign in with Google. Sign in with Facebook.”
Click it and you’ve agreed to their terms of service. You’ve handed them the license to shape your entire future relationship with AI. Do you trust them with that?
Because the agentic era raises questions that a login button can’t answer:
What does an AI agent actually know about you?
When one agent talks to another, who signed what?
Who’s accountable when something goes wrong?
Who governs the agents, and who proves compliance?
How can we trust agents to act on our behalf at all?
The answer to every one of those is verified identity . That’s what username.md provides.
What “verified identity” actually means in 2026
It comes down to three properties:
Provenance — the message provably came from an entity that holds a private key.
Anchoring — that key is bound to a stable, human-readable name ( chris.username.md , not a random GUID).
Portability — the name belongs to no single platform. You can move the DID to a different host and agents still resolve the same identity.
Here’s what that looks like on the wire. An OIDC discovery response:
1
$ curl "https://login.username.md/.well-known/openid-configuration" 2>/dev/null | jq .
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
{
"issuer" : "https://login.username.md" ,
"authorization_endpoint" : "https://login.username.md/oauth/v2/authorize" ,
"token_endpoint" : "https://login.username.md/oauth/v2/token" ,
"introspection_endpoint" : "https://login.username.md/oauth/v2/introspect" ,
"userinfo_endpoint" : "https://login.username.md/oidc/v1/userinfo" ,
"revocation_endpoint" : "https://login.username.md/oauth/v2/revoke" ,
"end_session_endpoint" : "https://login.username.md/oidc/v1/end_session" ,
"device_authorization_endpoint" : "https://login.username.md/oauth/v2/device_authorization" ,
"jwks_uri" : "https://login.username.md/oauth/v2/keys" ,
"scopes_supported" : [
"openid" ,
"profile" ,
"email" ,
"phone" ,
"address" ,
"offline_access"
]
... truncated
}
And an RFC 9421 signature on every response your handle returns:
1
2
3
Content-Digest: sha-256=:<base64(sha256(body))>:
Signature-Input: sig1=("content-digest" "@authority");created=...;keyid="platform-ed25519";alg="ed25519"
Signature: sig1=:<base64(ed25519 signature)>:
Every response from your handle is cryptographically signed, and any agent can verify it in about four lines of Go . If you want the details, the developer docs are here .
Why “sign in with Google” is not this
Google and GitHub can prove your account exists. They cannot prove you own your identity . The OAuth token they hand you is opaque: anyone can check who issued it, but it never gives you a signing key of your own. Send a message from whatever@gmail.com and there is no key you control behind it. It’s Google’s infrastructure, and Google isn’t going to give you the keys, because a captive user is a good user. If they close your account, you simply cease to exist online. By some reports that happens to hundreds of people every day.
OAuth was built for humans logging into apps. Agents need something OAuth was never asked to be.
Buy a handle and you get three things:
A verifiable identity you publish. chris.username.md resolves to a signed DID document, an OIDC discovery endpoint, an ATProto handle, and a plain-text, agent-readable profile. All for $9.
Native discoverability by agents. The signed SKILL.md at your handle tells other agents what you offer and how to reach you. They can verify it and trust it.
Ownership that outlives me. One payment. No subscription, no gas fees, no wallet. If the company disappears, your handle survives in a public DID registry snapshot.
Why now, why me, why username.md
I’ve spent most of my life in technology, starting at a help desk, moving through systems administration, and landing where I am now as a Principal Site Reliability Engineer. Longtime readers know DashPC and Dashboard Linux, and how I built in public the same kind of infotainment system that ships in every car today. It was greenfield, it was early, and it showed the world what was possible where automobiles met software.
I’m looking through that same lens again, and I’m handing you the missing piece of the AI puzzle. I’ve spent 20 years being “too early,” and I’ve made peace with it. I’m not claiming username.md is the next Cloudflare. But it could be.
You can stake a claim in the new AI gold rush right now. Reserve your username for $9. That’s the whole ask. Buy one for your agent. Grab one for a family member or a friend before the good names are gone.
For $9 you can own your name on the internet, forever. Cryptographically. Portably.
username.md is built entirely on open standards, and that’s the whole point. Every spec is published and linked: did:web , OIDC (RFC 6749 and OIDC Discovery), RFC 9421 signed responses, and the AT Protocol. The project is backed by my company, The Holding Company , and by chrisbergeron.com , the tech blog I’ve run since 2002.
I paid $9 to own my name on the internet forever. Now you can too: username.md
Why AI agents need verified identity
https://chrisbergeron.com/2026/08/11/why_ai_agents_need_verified_identity/
Like this article? Support me with
Subscribe for more great content
The Holding Company holdingco.com
Why AI agents need verified identity
Troubleshooting Consul systemd on Raspberry Pi
Dynamic SSH ports with Hashicorp Vault TOTP
Own Yourself - You Are Your Brand
© 2026 Chris Bergeron Powered by Hexo & Icarus
