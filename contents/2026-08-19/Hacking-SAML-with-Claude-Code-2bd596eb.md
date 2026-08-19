---
source: "https://oblique.security/blog/hacking-saml/"
hn_url: "https://news.ycombinator.com/item?id=49368038"
title: "Hacking SAML with Claude Code"
article_title: "Hacking SAML with Claude Code | Oblique"
image: "https://oblique.security/blog/hacking-saml/image.png"
author: "ericchiang"
captured_at: "2026-08-19T23:14:36Z"
capture_tool: "hn-digest"
hn_id: 49368038
score: 1
comments: 0
posted_at: "2026-08-19T22:33:47Z"
tags:
  - hacker-news
  - translated
---

# Hacking SAML with Claude Code

- HN: [49368038](https://news.ycombinator.com/item?id=49368038)
- Source: [oblique.security](https://oblique.security/blog/hacking-saml/)
- Score: 1
- Comments: 0
- Posted: 2026-08-19T22:33:47Z

## Translation

タイトル: クロード コードを使用した SAML のハッキング
記事のタイトル: クロード コードを使用した SAML のハッキング |斜め
説明: SAML がいかに安全ではないかについて何年も文句を言い続けた後、私はクロード コードを使用して、見つけられるすべての SAML 実装をハッキングすることで、それを証明しようと決意しました。

記事本文:
クロードコードを使用した SAML のハッキング |斜め
ドキュメント
ほぼ 10 年前に、Dex のサポートを追加するという残念な決定を下して以来、私は SAML プロトコルを使用してきました。 SAML は、XML デジタル署名という脳を溶かすテクノロジーのため、非常に危険であるとよく言われます。私がこれまでやり取りしたすべての SAML ライブラリには完全な認証バイパスがあることが最終的に判明し、Dex での私自身の実装では複数の脆弱性が発生しました。
数か月前、私はこの経験を利用して、Sec Nerds SF カンファレンスで SAML 脆弱性の歴史について講演しました。これは私にとって少し暴言を吐く機会でもありましたが、バイパス技術やこの分野の最新の研究についてもブラッシュアップすることができました。
この頃、私は Niels Provos の『Finding Zero-Days with Any Model』を読んで、SAML 脆弱性に関する研究をハーネスに適用したらどうなるだろうかと考え始めました。私はアカウントのガードレールを削除するために Anthropic のサイバー検証プログラムに申請し、承認された後、私の調査とプロンプトに基づいてエクスプロイトを作成できる Opus モデルにアクセスできるようになりました。
SAML は悪いプロトコルだと主張したいのですが、それを証明できるでしょうか?
1 か月ほどの空き時間に、見つけられるすべての SAML 実装をハッキングしようとしました。 (残念なことに) プロトコルを構造的に破ることはできませんでしたが、多くのバグがあり、その多くは現在も生き続けています。
脆弱性を見つけてください、クロード
クロードにコーディングを依頼する場合とハッキングを依頼する場合の主な違いの 1 つは、規模です。大規模なコードベースであっても、個々のバグを追跡したり、機能を実装したりするために必要なコンテキストの量は比較的少ない場合があります。一方、ハッキングは徹底的に行うもので、最後までやり遂げます。

バグを見つけたり、追跡するスレッドがなくなったりします。
これを説明するために、基本的にすべての「ハッキング ハーネス」 (多数あります) は、いくつかのコア プリミティブを中心に結合されます。
複数のエージェント間で作業を分割する
エージェントが指示に使用できる中間結果用のストレージを提供する
実行する作業の優先順位付け、範囲指定、重複排除を行う
私も同じものを建てました。 (ソースコードは github.com/oblique-security/saml-research で確認してください)
ハーネスの組み立ては簡単でした。私がまだ驚いたのは、(Mythos、OpenAI/Hugging Face、モデルの脱出がこれほど報道されていたにもかかわらず)、Claude Opus にハッキングの方法を教える必要がまったくなかったことです。クロードに脆弱性のコーパスを与えると、多くの場合、他のライブラリでまさにその問題のレプリカを見つけようとすることになります。脅威モデルを提供し、Claude に独自の探索を行わせることで、より良い結果が得られました。これは、エージェント コーディングに関する私の経験に似ています。モデルが改良されるにつれて、提供される具体的な指示が減り、より一般的なガイダンスが提供されるようになります。
私のパイプラインは 2 つのコア フェーズに落ち着きました。基盤となるライブラリで奇妙な動作を見つけようとする「ガジェット」フェーズと、ガジェットを組み合わせてエンドツーエンドのエクスプロイトを作成することでガジェットを確認する「調査結果」フェーズです。これらは JSONL ファイルに書き込まれ、さまざまなステップで操作が実行されます。たとえば、提案された結果を証明するためにトークンを消費する前に、範囲外である提案された結果を拒否するステップです。
以下は、Claude が Node の xml-crypto の処理命令の処理で見つけたガジェットの例です。
{
"id" : "g-0005" ,
"title" : "C14N DIFFERENTIAL -- 処理命令はデータ テキストにフラット化されます: `<?php echo 1; ?>` は裸の文字 `echo 1;` に正規化されます" ,
「影響」 : [
「s-0108」 、
「s-0126」 、
「s-0210」
]、
「ステータス」：「確認済み」
どちら

発見フェーズは後に電子メールの切り捨てバイパスに変わりました。
<サンプル:応答>
<saml:アサーション>
<ds:署名><!--有効--></ds:署名>
<saml:件名>
- <saml:NameID>not-an-admin@example.com</saml:NameID>
+ <saml:NameID><?p not-an-?>admin@example.com</saml:NameID>
</saml:件名>
</saml:アサーション>
</samlp:Response> すべての SAML 実装を調査しています
Claude Max 20x プランを実行して、パイプラインで何が見つかるかを確認する時期が来ました。
2020 年以降、主要な SAML ライブラリでバイパスがおよそ四半期に 1 回発生しています。 GitHub Enterprise だけでも、2024 年から 2025 年にかけて 4 件ありました ( CVE-2024-4985 、 CVE-2024-6800 、 CVE-2024-9487 、 CVE-2025-23369 )。基本的にこれらはすべて、コンポーネントが微妙に異なる方法で XML を解釈し、署名されていないデータを検証済みであるかのように誤って処理することに起因しています。現在、多くのライブラリは、この種の攻撃に対する防御を試みる強化された API を採用しています。
インターネットを炎上させることはできませんでしたが、4 つの異なるプロジェクトで完全な認証バイパスを見つけることができました。
Authentik : NameID にコメントを挿入すると、別のユーザーのアカウントに切り捨てられ、そのユーザーとして認証される可能性があります ( CVE-2026-57580 )
PHP litesaml/lightsaml : 応答メッセージの署名のラッピング (CVE-2026-63182)
OneUptime : 応答メッセージの署名ラッピング ( OneUptime/oneuptime#2949 )
Java の saml-client : 応答メッセージの署名ラッピング ( justinbleach/saml-client#149 )
繰り返しになりますが、SAML 実装を手動で実行するつもりがある場合は、行わないでください (代わりに OpenID Connect を使用するだけでもよいでしょうか?)。
Authentik の脆弱性は、8 人の異なる独立した研究者が同時に報告したため、最も興味深いものでした。 AI を使ってハッキングしているのは私だけではないことは明らかです。
多くのトップレベルの「誰でもログイン」攻撃を観察した後、私は次のことを始めました。

プロトコルの他の部分では重要です。 SAML では、ユーザーの電子メールを含む「SAML 応答」が署名されますが、技術的にはすべてのメッセージに署名を含めることができます。最も一般的に実装されるのは、AuthnRequest、AttributeQuery、および LogoutRequest です。これらの他のメッセージは、SAML レスポンスほどセキュリティ上の注意が払われることはほとんどなく、これらのコンポーネントの 12 のプロジェクト (さらに増え続けています) で署名のバイパスが見つかり、その結果、情報漏洩や任意のログアウトの問題が発生しました。 TypeScript プロジェクト samlify で公的に報告されている例。
最後に、サービス妨害ベクトルです。 SAML では基本的に、インターネット経由で POST された任意の XML ドキュメントを処理する必要があり、ほぼすべての主要な SAML ライブラリは、認証されていないリクエストによるメモリ不足状態の影響を受けやすくなっています。 Go エコシステム用のパッチを受け入れることができましたが、すべての Python ライブラリと Node ライブラリには現在未修正のアクティブな DoS 問題があります。
Go の xmldsig では、チェックが欠けていたため、署名検証中に 2 次的なメモリ割り当てが発生しました。 JavaScript の xmldom も同様の割り当てを引き起こし (レポートは現在まだ非公開ですが)、すべてのノード ライブラリに影響を与えます。 Python パッケージは、libxmlsec1 に渡す一連の変換をフィルター処理しないため、XSLT テンプレートによってライブラリが任意の大きなドキュメントを生成する可能性があります。
<ds:変換アルゴリズム="http://www.w3.org/TR/1999/REC-xslt-19991116">
<xsl:stylesheet version="1.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform">
<xsl:template match="/"><out>
<xsl:call-template name="d"><xsl:with-param name="n" select="27"/></xsl:call-template>
</out></xsl:template>
<xsl:template name="d"><xsl:param name="n"/><xsl:choose>
<xsl:when test="$n > 0">
<xsl:call-template name="d"><xsl:with-param name="n" select="$n - 1"/></xsl:call-templa

て＞
<xsl:call-template name="d"><xsl:with-param name="n" select="$n - 1"/></xsl:call-template>
</xsl:when><xsl:otherwise><x/></xsl:otherwise>
</xsl:choose></xsl:template>
</xsl:スタイルシート>
</ds:Transform> 2026 年の脆弱性研究の状況は暗いです。私の報告の多くはメンテナに届きましたが、メンテナは応答しないか、応答したとしても明らかに圧倒されていました。オープンソースのメンテナがまだ報われない仕事ではなかったとしても、今では AI のおかげで、本物の脆弱性レポートといい加減な脆弱性レポートが殺到するようになりました。私が問題を報告したある企業は、生成されたレポートを除外するために画面録画を送信するよう要求しました (私はこれを拒否しましたが、脆弱性の影響のため、彼らはそれでも対応し、修正し、CVE を受け取りました)。
私が最も頻繁に交流したのは、OneUptime プロジェクトとのやり取りでした。彼らは…言ってしまえば… AI を受け入れていました。非公開のレポートに対して 1 か月間返答がなかった後、私は公開問題をオープンしましたが、その日のうちに非常に大規模な PR がマージし、人間の承認も得ずに問題を修正したと主張しました。その後、新しいバイパスを見つけ、新しい大きな固定が現れました。 3 番目のバイパスの後、大規模な PR は安全な場所に到達したように見えました。
複数の SAML 署名のバイパス (#2949)
署名付きエラー応答による SAML 署名バイパス (#2981)
XML 処理命令は SAML NameID を変更できる (#2988)
LLM は安全な SAML コードを書くのも苦手であることがわかりました。
最終的に、私が停止点を選んだのは、バグがなくなったからではなく、これらのレポートに対する欲求がそこになかったからです。
SAML は長い間存在してきたため、強風が吹けば多数の「誰でもログイン」の問題が解決される段階を超えていますが、例外は引き続き新しく作成された実装です。長年にわたるissの恩恵を受けていないプロジェクト

ues は同じバグを再実装し、同じ種類のバイパスを持ちます。そのため、もう一度強調します。独自の SAML 実装をロールしないでください。
このようなメタ分析がもっと普及するといいですね。セキュリティ エンジニアは、JWT が良いか悪いかなどについてよく議論します。 GitHub 全体をスキャンして具体的な脆弱性を見つけ出すことができるということは、CVE データや逸話とは異なり、モチベーションを高めるのに役立ちます。プロトコルに問題があると思われる場合でも、それを証明するツールが手に入ります。
SAML を製品に統合することを考えている場合は、必ず実装に対して LLM 評価を実行してください。一般的な脆弱性パターンを探したい場合は、お気軽に私までご連絡ください。
Oblique チームからの製品ニュースと最新情報
リンクをコピー
SOC 2 タイプ II
製品
ドキュメント
変更履歴
学ぶ
ブログ
アクセスポリシーレポート
ガイド
参考資料
会社名
について
キャリア
よくある質問
信頼
メディアキット
イベント
© 2026 Oblique Inc. 無断複写・転載を禁じます。

## Original Extract

After many years of complaining about how insecure SAML is, I decided to try to prove it by using Claude Code to hack every SAML implementation I could find.

Hacking SAML with Claude Code | Oblique
Docs
I’ve worked with the SAML protocol ever since I made the unfortunate decision to add support to Dex almost a decade ago. SAML is often described as actively dangerous due to the brain melting technology that is XML Digital Signatures. Every SAML library I’ve ever interacted with has eventually been found to have a full authentication bypass, and my own implementation in Dex resulted in multiple vulnerabilities .
A couple months ago, I used this experience to give a talk on the history of SAML vulnerabilities at the Sec Nerds SF conference . This was an opportunity for me to rant a little, but also brush up on bypass techniques and the latest research in the area.
Around this time, I was reading Niels Provos’s “Finding Zero-Days with Any Model” and started to wonder what would happen if I took my research on SAML vulnerabilities and gave it to a harness. I applied to Anthropic’s Cyber Verification Program to remove the guardrails on my account, and after being accepted, I had access to an Opus model that could write exploits fed by my research and prompting.
I like to claim that SAML is a bad protocol, but could I prove it?
During my spare time over the course of a month or so, I attempted to hack every SAML implementation I could find. While I wasn’t able to structurally break the protocol (much to my chagrin), there were plenty of bugs, many of which are still live today.
Find me vulnerabilities, Claude
One of the core differences between asking Claude to code and asking it to hack is scale. Even in a large codebase, the amount of context needed to track down an individual bug or implement a feature can be relatively small. On the other hand, hacking is about being exhaustive, and going until you find a bug or run out of threads to chase.
To account for this, essentially all “hacking harnesses” (of which there are many) coalesce around a few core primitives:
Break up work between multiple agents
Provide storage for intermediate results that agents can use for direction
Prioritize, scope, and dedupe the work to be done
I built the same. (Check out the source code at: github.com/oblique-security/saml-research )
Building the harness was easy. What I was still surprised by (even despite all the coverage of Mythos, OpenAI/Hugging Face, and model escapes) was that I didn’t really need to teach Claude Opus how to hack. Giving Claude a corpus of vulnerabilities often devolved into it trying to find replicas of those exact issues in other libraries. I found better results by providing a threat model and then letting Claude do its own exploration. This is similar to my experience with agentic coding: as the models get better, you provide fewer specific instructions, and more general guidance.
My pipeline settled on two core phases. A “gadget” phase attempting to find weird behavior in the underlying libraries, and a “findings” phase that combines gadgets then confirms them by writing an end-to-end exploit. These are written to JSONL files, which different steps perform operations on. For example, a step that rejects proposed findings that are out-of-scope before spending tokens to prove them.
Here’s an example of a gadget that Claude found in Node’s xml-crypto’s handling of processing instructions:
{
"id" : "g-0005" ,
"title" : "C14N DIFFERENTIAL -- processing instructions are FLATTENED to their data text: `<?php echo 1; ?>` canonicalizes to the bare characters `echo 1;`" ,
"impacts" : [
"s-0108" ,
"s-0126" ,
"s-0210"
],
"status" : "confirmed"
} Which the finding phase later turned into an email truncation bypass :
<samlp:Response>
<saml:Assertion>
<ds:Signature><!--valid--></ds:Signature>
<saml:Subject>
- <saml:NameID>not-an-admin@example.com</saml:NameID>
+ <saml:NameID><?p not-an-?>admin@example.com</saml:NameID>
</saml:Subject>
</saml:Assertion>
</samlp:Response> Investigating every SAML implementation
It was now time to take my Claude Max 20x plan and see what my pipeline could find.
Since 2020, there’s been a bypass in a major SAML library about once a quarter. GitHub Enterprise alone had four across 2024-2025 ( CVE-2024-4985 , CVE-2024-6800 , CVE-2024-9487 , CVE-2025-23369 ). Essentially all of these stem from components interpreting XML in subtly different ways, and mistakenly processing unsigned data as if it was verified. Today many libraries have adopted hardened APIs that attempt to defend against these kinds of attacks.
While I wasn’t able to set the Internet ablaze, I was able to find full authentication bypasses in four different projects:
Authentik : Injecting a comment in NameID can be used to truncate to another user’s account and authenticate as them ( CVE-2026-57580 )
PHP litesaml/lightsaml : Signature wrapping on Response message ( CVE-2026-63182 )
OneUptime : Signature wrapping on Response message ( OneUptime/oneuptime#2949 )
Java’s saml-client : Signature wrapping on Response message ( justinbleach/saml-client#149 )
Again, if you’re ever going to hand roll an SAML implementation, don’t (maybe even just use OpenID Connect instead?).
The Authentik vulnerability was the most interesting because eight different independent researchers reported it at the same time. Clearly, I’m not the only one hacking with AI.
After looking at many top level “login as anyone” attacks, I started poking around other parts of the protocol. In SAML, a “SAML Response” containing a user’s email is signed, but technically any message can have a signature, AuthnRequest, AttributeQuery, and LogoutRequest being the most commonly implemented. These other messages rarely get the kind of security attention the SAML Response does, and I found signature bypasses in twelve projects (and counting) for these components, resulting in information disclosure and arbitrary logouts issues. An example that’s been publicly reported in the TypeScript project samlify .
Finally, denial-of-service vectors. SAML fundamentally requires processing arbitrary XML documents POST’d over the internet, and almost every major SAML library was susceptible to out-of-memory conditions from unauthenticated requests. I was able to get patches accepted for the Go ecosystem, but all Python and Node libraries have active DoS issues that are currently unfixed.
Go’s xmldsig had a quadratic memory allocation during signature validation due to a missing check. JavaScript’s xmldom causes similar allocations (though the report is currently still private ) impacting all Node libraries. Python packages don’t filter the set of transforms they pass to libxmlsec1, so an XSLT template can cause the libraries to generate an arbitrarily large document:
<ds:Transform Algorithm="http://www.w3.org/TR/1999/REC-xslt-19991116">
<xsl:stylesheet version="1.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform">
<xsl:template match="/"><out>
<xsl:call-template name="d"><xsl:with-param name="n" select="27"/></xsl:call-template>
</out></xsl:template>
<xsl:template name="d"><xsl:param name="n"/><xsl:choose>
<xsl:when test="$n &gt; 0">
<xsl:call-template name="d"><xsl:with-param name="n" select="$n - 1"/></xsl:call-template>
<xsl:call-template name="d"><xsl:with-param name="n" select="$n - 1"/></xsl:call-template>
</xsl:when><xsl:otherwise><x/></xsl:otherwise>
</xsl:choose></xsl:template>
</xsl:stylesheet>
</ds:Transform> The state of vulnerability research in 2026 is bleak. Many of my reports landed on with maintainers who either didn’t respond, or if they did, were clearly overwhelmed. If being an open source maintainer wasn’t already a thankless job, they’re now getting flooded with both genuine and slop vulnerability reports , thanks to AI. One company I reported an issue to insisted I submit a screen recording to filter out generated reports (which I refused to do, but due to the impact of the vuln, they still responded, fixed it, and I got a CVE).
One of the most of-the-times interactions I had was with the OneUptime project, who… let’s say… have embraced AI. After no response to a private report for a month, I opened a public issue, only for a very large PR to merge later that day claiming to have fixed it with no human acknowledgement. I then found a new bypass, and a new large fixed appeared. After the third bypass, the large PRs seemed to arrive at something secure:
Multiple SAML signature bypasses (#2949)
SAML signature bypass via signed error response (#2981)
XML processing instruction can modify SAML NameID (#2988)
It turns out LLMs are also bad at writing secure SAML code.
Ultimately, I picked a stopping point not because I had run out of bugs, but because the appetite for these reports wasn’t there.
SAML has been around for long enough that we’ve grown past the point where a stiff breeze will shake out numerous “login as anyone” issues, though the exception continues to be newly written implementations. Projects without the benefit of many years of issues will reimplement the same bugs and have the same kinds of bypasses, so I would again stress: do not roll your own SAML implementation.
I hope this kind of meta analysis becomes more popular. Security Engineers will often debate things like whether or not JWT are good or bad. Having the ability to scan all of GitHub and come up with concrete vulnerabilities is motivating in a way that CVE data and anecdotes aren’t. If you think a protocol has issues, you now have the tools to prove it.
If you’re ever thinking of integrating SAML into a product, be sure to run an LLM evaluation against your implementation. And if you want some common vulnerability patterns to look for, feel free to reach out to me!
Product news and updates from the Oblique team
Copy link
SOC 2 Type II
Product
Documentation
Changelog
Learn
Blog
Access policies report
Guides
Reference
Company
About
Careers
FAQs
Trust
Media kit
Events
© 2026 Oblique Inc. All rights reserved.
