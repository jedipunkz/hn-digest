---
source: "https://blog.coinkite.com/adding-to-public-record/"
hn_url: "https://news.ycombinator.com/item?id=49181980"
title: "AI review missed Coldcard flaw"
article_title: "Adding to the Public Record on Our Ongoing Investigation | COINKITE Blog"
author: "hughw"
captured_at: "2026-08-05T12:48:29Z"
capture_tool: "hn-digest"
hn_id: 49181980
score: 1
comments: 0
posted_at: "2026-08-05T12:35:43Z"
tags:
  - hacker-news
  - translated
---

# AI review missed Coldcard flaw

- HN: [49181980](https://news.ycombinator.com/item?id=49181980)
- Source: [blog.coinkite.com](https://blog.coinkite.com/adding-to-public-record/)
- Score: 1
- Comments: 0
- Posted: 2026-08-05T12:35:43Z

## Translation

タイトル: AI レビューで Coldcard の欠陥が見逃された
記事のタイトル: 進行中の調査に関する公的記録への追加 |コインカイトのブログ
説明: 影響を受けるユーザー向けの新しい情報、公開 COLDCARD セキュリティ レコード、およびセキュリティ クリティカルなコードをレビューするチーム向けのレッスン。

記事本文:
ブログ
キャリア
お問い合わせ
RSS
メールニュースレター
ストア
×
ホーム
ブログ
キャリア
お問い合わせ
RSS
メールニュースレター
ストア
← 投稿に戻る
現在進行中の調査に関する公的記録への追加
私たちは現時点で本当の怒りがあることを理解しています。ユーザーは実際に被害を受けています
損失が発生しており、影響を受けた人々にとっては、公式声明だけでは十分ではありません。
私たちは引き続き影響を受ける顧客を直接サポートし、他の顧客にも連絡するよう促します。
影響を受ける可能性があることを認識している他のユーザーに送信します。緊急: もし
あなたのシードは、影響を受けるファームウェアを使用して生成されました。少なくとも 50
独立したプライベートなサイコロロールであり、あなたの資金は強力な組織によって保護されていません。
固有の BIP-39 パスフレーズ、それらの資金を新しいウォレットに移動します
今。
また、この脆弱性はあらゆる企業の建物に対する警告であると考えています。
私たちだけではなく、ビットコインのハードウェアとソフトウェア、そして私たちはこれを今公開しています。
他の企業が確認するのに時間がかかるため、詳細はまだ新しいですが
さらなる損失の可能性を防ぐための独自のコード。
その取り組みをサポートするために、新しいリソースを公開しました。
Coinkite.com/historyal-disclosures 、
すべての既知の公安調査を記録するページ。
開示、専門的レビュー、内部調査結果、およびセキュリティ勧告
COLDCARD デバイスに影響します。
私たちがこれを公開しているのは、研究者、ジャーナリスト、セキュリティ チームが取り組んでいるからです。
彼ら自身のレビューには、決定的な記録が必要です。この情報は随時更新していきます
捜査が続くにつれて。
独立した研究者が公にしているように
確認された、このファームウェア
バグは、2 つの無関係なサブモジュール間の境界に生息しているようです。
暗号化ロジックやビットコイン固有のロジックではなく、親コード内にあります。
ほとんどの社内および第三者によるレビューの対象となります。
フラグチェックが正しいように見えたため、バグは気付かれずに放置されました

d、そして
その潜在的な影響はリリースごとに増大しました。
私たちは、このバグがどのように発生するのかをより広範なエコシステムが理解することが重要であると考えています。
同様の結果を回避できるように、それが発生した理由と検出を回避した理由を説明します。
AI 支援レビューで何が把握でき、何が把握できなかったか
私たち自身のコードレビューにおける AI の使用について疑問があることは承知しています。カバーします
これは完全に事後調査ですが、現在活発な調査が行われていることから、
すぐに言えることは次のとおりです。
私たちは、重要なコードベースに対して AI 支援によるレビューを実行しました。
悪用の数週間前。この脆弱性は発見されませんでした。以来、
この事件を受けて、私たちは、キミを含むフロンティアモデルに対してコードをテストしました。
K3、クロード・ファブル、コーデックス 5.6。誰も捕まえられなかった。
これが、私たち、そして AI ツールに依存している他の人にとって、次のことを明確にする理由です。
彼らが現在捕らえているものと、捕獲できないかもしれないもの。
他のチームに求めていること
私たちは、学んだことを独自のレビュープロセスとして共有することに引き続き取り組んでいます
続けます。チームがセキュリティ クリティカルなコードの AI レビューに依存している場合、
特にビルドとサブモジュールの境界に対してテストすることをお勧めします。
これらの新しく強力な AI モデルの結果、多くのビットコインが
オープンソース コードに依存するプロジェクトを含むプロジェクトでは、早急な対応が必要です。
レビュー。
私たちはできる限りのことに取り組み、直接の影響を受ける人々をサポートしています。
そして私たちが知っているすべてを公開します。
@COLDCARDwallet
@SATSCARD
@TAPSIGNER
@theBLOCKCLOCK
@arcasafes
ビデオチュートリアル
@コインカイト
Coinkite製
ストアを見る
製品

## Original Extract

New information for affected users, the public COLDCARD security record, and lessons for teams reviewing security-critical code.

Blog
Careers
Contact
RSS
Email Newsletter
Store
×
Home
Blog
Careers
Contact
RSS
Email Newsletter
Store
← Back to posts
Adding to the Public Record on Our Ongoing Investigation
We understand there is real anger at this moment. Users have suffered real
losses, and for those impacted, no public statement is enough.
We continue to support affected customers directly and urge others to reach
out to any other users they are aware of who may be affected. Urgently: if
your seed was generated with our affected firmware, without at least 50
independent, private dice rolls, and your funds aren’t protected by a strong,
unique BIP-39 passphrase, move those funds to a new wallet
now .
We also believe this vulnerability is a warning for every company building
Bitcoin hardware and software, not only us—and we’re publishing this now,
while the details are still fresh, because other companies need time to check
their own code to prevent potential further loss.
To support that effort, we’ve published a new resource:
coinkite.com/historical-disclosures ,
a page that records all known public security research, coordinated
disclosures, professional reviews, internal findings, and security advisories
affecting COLDCARD devices.
We’re publishing it because researchers, journalists, and security teams doing
their own review need a definitive record to work from. We’ll keep this updated
as the investigation continues.
As independent researchers have publicly
corroborated , this firmware
bug appears to have lived at a boundary between two unrelated submodules, not
in the parent code, and not in the cryptographic or Bitcoin-specific logic that
are the subject of most internal and third-party reviews.
Because the flag check looked correct, the bug silently went unnoticed, and
its potential impact grew with every release.
We believe it’s important for the broader ecosystem to understand how this bug
arose, and why it evaded detection, so they can avoid similar consequences.
What AI-Assisted Review Did, and Didn’t, Catch
We know there are questions about our own use of AI in code review. We’ll cover
this fully in our post-mortem, but given the active investigation right now,
here’s what we can say immediately.
We’ve run AI-assisted review against our critical codebases, including in the
weeks before the exploit. It did not catch this vulnerability. Since the
incident, we’ve also tested our code against frontier models, including Kimi
K3, Claude Fable, and Codex 5.6. None of them caught it.
It’s a reason for us, and anyone else relying on AI tools, to be specific about
what they currently catch and what they might not.
What We’re Asking Other Teams to Do
We remain committed to sharing what we learn as our own review process
continues. If your team relies on AI review of security-critical code, we
recommend you test it specifically against build and submodule boundaries.
As a result of these new and powerful AI models, we believe many Bitcoin
projects, including those that rely on open-source code, require immediate
review.
We’re addressing what we can, we’re supporting the people directly affected,
and publishing everything we know.
@COLDCARDwallet
@SATSCARD
@TAPSIGNER
@theBLOCKCLOCK
@arcasafes
Video Tutorials
@Coinkite
Made with by Coinkite
View store
Products
