---
source: "https://blogs.gnome.org/mcatanzaro/2026/07/20/some-changes-to-gnome-security-tracking/"
hn_url: "https://news.ycombinator.com/item?id=48981193"
title: "GNOME Changes Security Disclosure Policies Due to AI-Generated Reports"
article_title: "Some Changes to GNOME Security Tracking – Michael Catanzaro's Blog"
author: "Erenay09"
captured_at: "2026-07-20T17:24:08Z"
capture_tool: "hn-digest"
hn_id: 48981193
score: 2
comments: 0
posted_at: "2026-07-20T16:35:37Z"
tags:
  - hacker-news
  - translated
---

# GNOME Changes Security Disclosure Policies Due to AI-Generated Reports

- HN: [48981193](https://news.ycombinator.com/item?id=48981193)
- Source: [blogs.gnome.org](https://blogs.gnome.org/mcatanzaro/2026/07/20/some-changes-to-gnome-security-tracking/)
- Score: 2
- Comments: 0
- Posted: 2026-07-20T16:35:37Z

## Translation

タイトル: AI 生成レポートにより、GNOME がセキュリティ開示ポリシーを変更
記事のタイトル: GNOME セキュリティ追跡の一部の変更 – Michael Catanzaro のブログ
説明: AI によって生成されたセキュリティ脆弱性レポートの増加により、GNOME による脆弱性レポートの管理方法にいくつかの変更が必要な時期が来ています。これらのポリシー変更では、AI が生成したコンテンツを含むレポートと含まないレポートを意図的に区別しません。全員が同じルールに従う
[切り捨てられた]

記事本文:
コンテンツにスキップ
マイケル・カタンザーロのブログ
一般的な GLib プログラミング エラー
ビルド オプションのベスト プラクティス
クラッシュレポート用の高品質なバックトレースの作成
systemd-resolved、Split DNS、および VPN 構成について
GNOME セキュリティ追跡の一部の変更
AI によって生成されたセキュリティ脆弱性レポートの増加により、GNOME による脆弱性レポートの管理方法にいくつかの変更が必要な時期が来ています。
これらのポリシー変更では、AI が生成したコンテンツを含むレポートと含まないレポートを意図的に区別しません。すべての脆弱性レポートに対して同じルールに従うほうが、2 つの異なる方法を使用するより簡単です。記者が AI の使用を明らかにすることはめったにありません。問題レポートが AI によって生成されたものであるかどうかを推測する必要がないのは素晴らしいことです。通常は明らかですが、常にそうとは限りません。また、AI によって発見されない脆弱性報告はますます稀になってきています。 AI 以外のレポートは現在ではやや珍しいものになっているため、AI 以外のレポートを最適化することはあまり意味がありません。
従来、私は GNOME Security に報告されたすべてのセキュリティ問題に 90 日の開示期限を適用してきました。 90 日は業界標準のタイムラインですが、GNOME では特にうまく機能しません。実際には、ほぼすべての GNOME メンテナは次の 2 つの方法のいずれかで脆弱性レポートを処理します。
プロジェクトの管理者は、問題が報告されてから通常 1 ～ 3 週間以内に問題を迅速に修正します。
プロジェクトの管理者は問題をまったく修正しません。問題レポートは最終的に 90 日の開示期限に達し、その時点で機密保持の設定を解除しました。
90 日という期限は、問題が公になる前にプロジェクトの貢献者に問題を修正する時間を与えることを目的としていますが、実際にはメンテナはこの時間のほとんどを実際には活用していません。問題レポートを公開し、修正されたら CVE をリクエストします

または開示期限に達したときのいずれか早い方。 CVE が割り当てられると、定期的なプロジェクト管理者ではない貢献者がそれを修正しようとすることがあります。したがって、問題レポートを 90 日間秘密にしておくのは、役に立たない遅延が生じるだけです。
他のいくつかのプロジェクト、特に Linux カーネルは、AI によって発見される可能性のある脆弱性は攻撃者にすでに知られていると考えられることに基づいて、AI によって生成されたと思われる問題レポートに対して即時完全開示ポリシーを導入しています。しかし、このポリシーはかなり極端なようで、問題を早急に修正するようプレッシャーを感じているかもしれないメンテナーにとっては確かに不親切です。即時開示は GNOME ではうまく機能しません。
代わりに、2026 年 8 月 1 日以降に報告された問題については、30 日間の開示期限に切り替えます。これは良い妥協案のように思えます。 AI が生成する問題レポートの増加を考慮しない場合でも、期限を短くすることはおそらく GNOME にとってより効果的でしょう。
AI生成コンテンツを禁止するプロジェクトの手順
プロジェクトが AI 生成コンテンツを含む問題レポートを禁止している場合、GNOME Security に報告されたセキュリティ問題をプロジェクトの問題トラッカーに転送しません。脆弱性レポートの圧倒的多数に AI 生成コンテンツが含まれており、プロジェクトのポリシーに違反することになるためです。代わりに、GNOME セキュリティ問題トラッカーの問題レポートをすぐに閉じて、プロジェクトの管理者に ping を送信して、レポートの存在を知らせます。プロジェクトの問題トラッカーで脆弱性レポートを受け取りたい場合は、プロジェクトの AI ポリシーを変更して、脆弱性レポートの例外を作成してください。
残念ながら、GNOME メンテナは、この問題トラッカーと GitLab の機密問題にアクセスできません。

機密問題レポートに関して個々の開発者に CC を送信することは許可されていません。私はこれらの問題についてのみ即時開示を採用する予定でしたが、代わりに権限を拡張してすべての GNOME 開発者が問題トラッカーを表示できるようにする必要があるかもしれません。ご意見は大歓迎です。
私は 2020 年 11 月から GNOME セキュリティ問題の追跡を管理しています。(この作業をサポートしてくださった Red Hat に感謝します。) セキュリティ追跡は主に秘書の仕事です。私は問題が報告されたときと問題が解決されたときに追跡し、期限に達したら公開し、必要に応じて CVE を要求します。大した作業量ではありませんが、飽きてきたので、切り替える時期が来ました。 2026 年 11 月 1 日に、新しく報告されたセキュリティ問題の追跡を中止します。11 月中は、11 月 1 日より前に報告された問題の追跡のみに焦点を当てます。12 月 1 日までに、その一連の問題のすべての開示期限に達し、完了します。
現在、GNOME のセキュリティ問題を追跡している人は他にいません。あなたが経験豊富な GNOME コミュニティ メンバーで、この作業を引き継ぐことに興味がある場合は、お知らせください。開始をお手伝いします。 (セキュリティ追跡は初心者にとっては良い作業ではありません。)
これは、追跡インフラストラクチャを改善する機会となる可能性もあります。私は wiki ページを使用していますが、これはかなり原始的であり、かなりの手動メンテナンスが必要です。たとえば、問題レポートが閉じられると、ページの更新を忘れがちです。理想的には、Wiki を、問題の実際の状態に基づいて動的に更新する適切な Web アプリに置き換えます。
あなたのメールアドレスは公開されません。 * が付いているフィールドは必須です
Fedora ワークステーション、GNOME、Epiphany、および WebKitGTK 上

## Original Extract

Due to the increase in AI-generated security vulnerability reports, it is time for some changes in how GNOME manages vulnerability reports. These policy changes intentionally do not distinguish between reports that contain AI-generated content and those that do not. Following the same rules for all
[truncated]

Skip to content
Michael Catanzaro's Blog
Common GLib Programming Errors
Best Practices for Build Options
Creating Quality Backtraces for Crash Reports
Understanding systemd-resolved, Split DNS, and VPN Configuration
Some Changes to GNOME Security Tracking
Due to the increase in AI-generated security vulnerability reports, it is time for some changes in how GNOME manages vulnerability reports.
These policy changes intentionally do not distinguish between reports that contain AI-generated content and those that do not. Following the same rules for all vulnerability reports is simpler than having two different ways of doing things. Reporters rarely disclose AI use, and it’s nice to not have to guess whether the issue report is AI-generated or not; it’s normally obvious, but not always. Also, vulnerability reports that are not discovered by AI are becoming increasingly rare. Non-AI reports are now moderately unusual, so it really doesn’t make sense to optimize for them.
Traditionally, I have applied a 90 day disclosure deadline to all security issues reported to GNOME Security. 90 days is an industry standard timeline, but it doesn’t work particularly well for GNOME. In practice, almost all GNOME maintainers handle vulnerability reports in one of two ways:
The project maintainer fixes the issue quickly, typically within 1-3 weeks after it is reported.
The project maintainer does not fix the issue at all. The issue report eventually reaches the 90-day disclosure deadline, at which point I unset confidentiality.
The 90-day deadline is intended to allow project contributors time to fix the issue before it becomes public, but in practice, maintainers do not actually make use of most of this time. I disclose the issue report and request a CVE when it is fixed or when the disclosure deadline is reached, whichever comes first. Once a CVE is assigned, contributors who are not regular project maintainers will sometimes attempt to fix it. Accordingly, keeping the issue reports confidential for 90 days only introduces a delay that is not useful.
Some other projects, notably the Linux kernel, have implemented an immediate full disclosure policy for issue reports that seem to be AI-generated, on the basis that a vulnerability that can be discovered by AI is presumably already known to attackers. But this policy seems pretty extreme, and is certainly unkind to maintainers who might feel pressured to urgently fix the issue. Immediate disclosure would not work well for GNOME.
Instead, I will switch to a 30 day disclosure deadline for issues reported on August 1, 2026 or later. This seems like a good compromise. The shorter deadline would probably work better for GNOME even if not for the increase in AI-generated issue reports.
Procedure for Projects that Prohibit AI-Generated Content
If a project prohibits issue reports that contain AI-generated content, I will no longer forward security issues reported to GNOME Security to the project’s issue tracker, since the overwhelming majority of vulnerability reports contain AI-generated content and would violate the project’s policy. Instead, I will immediately close the issue report in the GNOME Security issue tracker, then ping the project maintainers to let them know about the existence of the report. If you prefer to receive vulnerability reports in your project’s issue tracker, then please change your project’s AI policy to make an exception for vulnerability reports .
Unfortunately, GNOME maintainers don’t have access to confidential issues in this issue tracker, and GitLab does not allow CCing individual developers on confidential issue reports. I had been planning to adopt immediate disclosure for these issues only, but perhaps we should instead expand the permissions to allow all GNOME developers to see the issue tracker. Opinions welcome.
I have been managing GNOME security issue tracking since November 2020. (Thank you to Red Hat for supporting this work.) Security tracking is largely a secretarial duty: I keep track of issues when they are reported and when they are closed, disclose them when the deadline is reached, and request CVEs when appropriate. It is not a huge amount of work, but I am getting tired of it, so it’s time for a change. I will discontinue tracking newly-reported security issues on November 1, 2026. During November, I will focus only on tracking issues reported prior to November 1. By December 1, all disclosure deadlines for that set of issues will have been reached, and I will be done.
Currently nobody else is tracking GNOME security issues. If you are an experienced GNOME community member and you are interested in taking over this work, let me know and I will help you get started. (Security tracking is not a good task for newcomers.)
This may also be an opportunity to improve our tracking infrastructure. I use a wiki page , but this is fairly primitive and requires considerable manual upkeep. It’s easy to forget to update the page when an issue report is closed, for example. Ideally, we would replace the wiki with a proper web app that dynamically updates based on the actual state of the issue.
Your email address will not be published. Required fields are marked *
On Fedora Workstation, GNOME, Epiphany, and WebKitGTK
