---
source: "https://techcrunch.com/2026/06/08/microsofts-open-source-tools-were-hacked-to-steal-passwords-of-ai-developers/"
hn_url: "https://news.ycombinator.com/item?id=48457830"
title: "Microsoft's open source tools were hacked to steal passwords of AI developers"
article_title: "Microsoft's open source tools were hacked to steal passwords of AI developers | TechCrunch"
author: "raffael_de"
captured_at: "2026-06-09T10:19:30Z"
capture_tool: "hn-digest"
hn_id: 48457830
score: 155
comments: 58
posted_at: "2026-06-09T07:33:16Z"
tags:
  - hacker-news
  - translated
---

# Microsoft's open source tools were hacked to steal passwords of AI developers

- HN: [48457830](https://news.ycombinator.com/item?id=48457830)
- Source: [techcrunch.com](https://techcrunch.com/2026/06/08/microsofts-open-source-tools-were-hacked-to-steal-passwords-of-ai-developers/)
- Score: 155
- Comments: 58
- Posted: 2026-06-09T07:33:16Z

## Translation

タイトル: Microsoft のオープンソース ツールがハッキングされ、AI 開発者のパスワードが盗まれた
記事のタイトル: Microsoft のオープンソース ツールがハッキングされ、AI 開発者のパスワードが盗まれた |テッククランチ
説明: Microsoft は、ハッキングの報告を受けて、Azure および AI コーディング ツール用の数十の GitHub コード リポジトリを閉鎖しました。

記事本文:
TechCrunch デスクトップのロゴ
TechCrunch モバイルのロゴ
最新の
Microsoftのオープンソースツールがハッキングされ、AI開発者のパスワードが盗まれた
Microsoftは、ハッカーが明らかにどのようにしてプロジェクトに侵入し、コードにパスワードを盗むマルウェアを注入したのかを調査するため、GitHubでホストされている数十のオープンソースプロジェクトへのアクセスを遮断した。
影響を受けるプロジェクトの多くは、Microsoft のクラウド サービス Azure や、Claude Code、Gemini のコマンド ライン インターフェイス、VS Code など、開発者が AI 開発アプリでコーディングするために使用するその他のツールに関連しています。
このハッキングを最初に報告したセキュリティ会社 Cloudsmith とコミュニティ主導のマルウェア分析サイト OpenSourceMalware によると、ハッカーが AI コーディング アプリで侵害されたツールを開いたときに、このマルウェアによってユーザーのパスワードやその他の機密認証情報が盗まれる可能性がありました。
影響を受けるツールをダウンロードした人が何人いるかはすぐにはわかりません。
404 Media が最初に報じたように、Microsoft はリポジトリを削除したことを認めました。
Microsoftの広報担当者Ben HopeはTechCrunchに対し、同社は「潜在的な悪意のあるコンテンツを調査するため、一部のリポジトリを一時的に削除した」と語った。
「これらのリポジトリの一部はレビュー後に復元されましたが、その他のリポジトリは作業が継続している間オフラインのままになる可能性があります。」
「調査の一環として、影響を受けるリポジトリからコンテンツをプルダウンした可能性がある少数の顧客に通知しました。引き続き調査を続け、顧客の対応が必要なことがさらに特定された場合は、確立したサポート チャネルを通じて直接連絡します」とホープ氏は付け加えました。
TechCrunchの質問に対し、Microsoftは影響を受ける顧客の具体的な数をすぐには明らかにしなかった。
Microsoft に属する少なくとも 70 のプロジェクトが「無効化」されました。

Microsoft が所有するコード ホスティング サイトである GitHub 上のプロジェクトのページにアクセスしようとしたときに読み込まれるメッセージによるものです。 「GitHub の利用規約に違反したため、このリポジトリへのアクセスは GitHub スタッフによって無効にされました。」
これは、コンピューターにコードがインストールされている多数のユーザーにマルウェアを植え付ける目的で、ハッカーが広く人気のあるオープンソース プロジェクトに侵入したここ数カ月の最新の例です。これらのハッキングは、多数のソフトウェア製品や特定の種類のユーザーによって頻繁に使用されるコードを標的とするため、「サプライ チェーン」攻撃として知られています。ユーザーはクラウド システムや大量の顧客データにアクセスできる場合があるため、ハッキングには有利である可能性があります。
オープンソース プロジェクトの単独開発者がハッカーの標的になることは珍しいことではありませんが、場合によっては開発者の信頼を得るための長期にわたる取り組みの一環として行われますが、この種の攻撃から防御するリソースを備えている Microsoft のような大手テクノロジー企業が侵害されることはまれです。
Ars Technica によると、これはマイクロソフトが過去数週間でハッカーに同社のオープンソース プロジェクトを侵害することを可能にした既知の侵害としては 2 件目です。 5月中旬、セキュリティ研究者らは、MicrosoftのオープンソースプロジェクトであるDurable Task（開発者によるアプリ構築を支援するツール）がハッキングされたと発表した。 OpenSourceMalware は、Microsoft の最新のインシデントは Durable Task プロジェクトの「再侵害」であると述べ、Microsoft が最初の試みやまったく新しい明確な侵害でハッカーを根絶できなかった可能性を示唆しています。
Microsoft からのコメントを追加して更新しました。
記事内のリンクを通じて購入すると、少額の手数料が発生する場合があります。これは編集上の独立性に影響しません。
ザック・ウィテカー
セキュリティエディター
ザック・ウィテカーが警備員

TechCrunchの編集者です。彼はまた、今週のセキュリティに関する週刊サイバーセキュリティ ニュースレターの執筆者でもあります。
Signal の zackwhittaker.1337 で暗号化されたメッセージを介して彼に連絡できます。電子メールで彼に連絡したり、支援活動を確認するには zack.whittaker@techcrunch.com に連絡したりすることもできます。
6月18日
ロサンゼルス
マッハ インダストリーズ、ファウンダーズ ファンド、シンケイ システムズのリーダーから、規模を拡大し成功するために何が必要かについて内部を見てみましょう。炉端での率直なチャットや影響力の高いネットワーキングを通じて、貴重な洞察と新しいつながりを得ることができます。
WWDC 2026: Siri AI、iOS 27、Apple Intelligence などに関するすべての発表
これはトークンポカリプスの夜明けでしょうか?
創設者らがVCの恐ろしい話を共有、中には名前を挙げている人もいる
GoogleはSpaceXにコンピューティング費用として月額9億2000万ドルを支払う
ミラ・ムラティは慎重にスポットライトに戻ります
IPOを前に、Anthropicのダニエラ・アモデイ氏はAIのリターンに対する疑念を一蹴
Microsoft、OpenClaw からインスピレーションを得たパーソナル アシスタント「Scout」を発売

## Original Extract

Microsoft shut down dozens of GitHub code repositories for Azure and AI coding tools after a reported hack.

TechCrunch Desktop Logo
TechCrunch Mobile Logo
Latest
Microsoft’s open source tools were hacked to steal passwords of AI developers
Microsoft has cut off access to dozens of its open source projects hosted on GitHub as it investigates how hackers apparently breached the projects and injected password-stealing malware into the code.
Many of the affected projects relate to Microsoft’s cloud service Azure and other tools used by developers to code with AI development apps, such as Claude Code, Gemini’s command line interface, and VS Code.
According to security firm Cloudsmith and community-driven malware analysis site OpenSourceMalware , which were some of the first to flag the hack, the malware allowed the hackers to steal the users’ passwords and other sensitive credentials when they opened the compromised tools in their AI coding apps.
It’s not immediately known how many people have downloaded the affected tools.
Microsoft confirmed it pulled the repos, as first reported by 404 Media .
Microsoft spokesperson Ben Hope told TechCrunch that the company has “temporarily removed some repositories as we investigated potential malicious content.”
“Some of these repos have been restored after review, while others may remain offline while work continues.”
“As part of our investigation, we notified a small number of customers who may have pulled down content from the affected repositories. We will continue to investigate, and if anything further is identified that requires customer action, we will reach out directly through our established support channels,” added Hope.
Microsoft did not immediately provide the specific number of customers affected, when asked by TechCrunch.
At least 70 projects belonging to Microsoft have been “disabled,” per a message loading when trying to access the projects’ pages on GitHub, a code-hosting site that Microsoft owns. “Access to this repository has been disabled by GitHub Staff due to a violation of GitHub’s terms of service.”
This is the latest example in recent months of hackers breaching widely popular open source projects with the aim of planting malware on a large number of users who have the code installed on their computers. These hacks are known as “supply chain” attacks as they target code that is often used in a large number of software products, or by a specific kind of user, which may be advantageous to hack as they sometimes have access to cloud systems and large amounts of customers’ data.
While it’s not uncommon for sole developers of open source projects to be targeted by hackers — in some cases as part of long-running efforts to gain the trust of the developer — it is rare for large tech giants like Microsoft, which have the resources to defend against these kinds of attacks, to get breached.
This is Microsoft’s second known breach over the past few weeks that has allowed hackers to compromise its open source projects, per Ars Technica . In mid-May, security researchers said that Microsoft’s open source project Durable Task, a tool that helps developers build apps, was hacked. OpenSourceMalware said that Microsoft’s latest incident is a “re-compromise” of the Durable Task project, suggesting that Microsoft may not have eradicated the hackers on its first attempt or an entirely new, distinct breach.
Updated with comment from Microsoft.
When you purchase through links in our articles, we may earn a small commission . This doesn’t affect our editorial independence.
Zack Whittaker
Security Editor
Zack Whittaker is the security editor at TechCrunch. He also authors the weekly cybersecurity newsletter, this week in security .
He can be reached via encrypted message at zackwhittaker.1337 on Signal. You can also contact him by email, or to verify outreach, at zack.whittaker@techcrunch.com .
June 18
Los Angeles
Get an inside look at what it takes to scale and succeed from leaders at Mach Industries, Founders Fund, and Shinkei Systems. Through candid fireside chats and high-impact networking, you’ll walk away with valuable insights and new connections.
WWDC 2026: Everything announced on Siri AI, iOS 27, Apple Intelligence and more
Is this the dawn of the Tokenpocalypse?
Founders share VC horror stories, and some are naming names
Google will pay SpaceX $920M per month for compute
Mira Murati steps back into the spotlight, carefully
Ahead of its IPO, Anthropic’s Daniela Amodei shrugs off doubts about AI’s returns
Microsoft launches Scout, an OpenClaw-inspired personal assistant
