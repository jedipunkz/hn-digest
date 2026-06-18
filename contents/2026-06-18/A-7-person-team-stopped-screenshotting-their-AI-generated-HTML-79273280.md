---
source: "https://display.dev/customers/indigo-engineering"
hn_url: "https://news.ycombinator.com/item?id=48582440"
title: "A 7-person team stopped screenshotting their AI-generated HTML"
article_title: "How Indigo Engineering gave their agents' output a home – Display.dev"
author: "ottilves"
captured_at: "2026-06-18T10:36:07Z"
capture_tool: "hn-digest"
hn_id: 48582440
score: 2
comments: 1
posted_at: "2026-06-18T08:25:44Z"
tags:
  - hacker-news
  - translated
---

# A 7-person team stopped screenshotting their AI-generated HTML

- HN: [48582440](https://news.ycombinator.com/item?id=48582440)
- Source: [display.dev](https://display.dev/customers/indigo-engineering)
- Score: 2
- Comments: 1
- Posted: 2026-06-18T08:25:44Z

## Translation

タイトル: 7 人のチームが AI 生成の HTML のスクリーンショットを撮るのをやめた
記事のタイトル: Indigo Engineering がエージェントの出力にホームを与えた方法 – Display.dev
説明: Indigo Engineering の 7 人チームは、AI が生成した HTML の Slack スクリーンショットを、レンダリングされ、バージョン管理され、コメント可能なページである display.dev に置き換えました。再設計の高速化、Google ドキュメントからの静かな終了、Anthropic エンタープライズ アップグレードの延期。

記事本文:
Indigo Engineering がエージェントの出力にホームを与えた方法 – Display.dev Display.dev / display █ 仕組み ユースケース 価格設定 ドキュメント ニュース サインイン 無料で始める 仕組み ユースケース 価格設定 ドキュメント ニュース テーマ サインイン 無料で始める Indigo Engineering がどのようにしてエージェントの出力にホームを与えたのか。
Indigo Engineering は、進歩的な主催者がコミュニティで力を築くのを支援する、簡単なソフトウェアとデータ製品を構築しています。 7 人のチームは米国中に散らばっており、機械学習による最適化を備えた有権者連絡用 CRM である Matchbook と、有権者ファイル強化ツールである Matter の 2 つの製品を出荷しています。
2026 年にソフトウェアを出荷するほとんどのチームと同様に、Indigo のエンジニアとデザイナーはコーディング エージェントに大きく依存しています。これらのエージェントは、デザイン出力、レポート、仕様書、内部ドキュメントなど、HTML とマークダウンの一定のストリームを生成します。作品はすぐに生まれます。チームメイトの前でレビューし、コメントし、修正してもらうという作業が問題の解決につながりました。
以前: Slack スレッド、HTML スクリーンショット、バージョン管理の悪夢
「display.dev を導入する前は、Slack でアーティファクト ファイルを共有し、そこにコメントを書き込んでいました。HTML アーティファクトの場合、これは大量の乱雑で混乱を招くスクリーンショットを意味していました。また、複数の会話が同じスレッドに混ざっていることを意味し、バージョン管理は悪夢で、変更やディスカッションを追跡するのは非常に面倒でした。」 — マックス・ウッド、Indigo Engineering CEO
チームはこれがどこにつながるかを理解していましたが、代替案も気に入りませんでした。
「最終的にはアーティファクトを追跡するために git を使い始めたかもしれませんが、それでもレンダリングは煩わしく、リビジョンには継続的な PR が必要になり、スレッドを追跡するのは困難になります。」
変更後: エージェント出力の共有ホーム
Indigo は表示するマークダウンと HTML アーティファクトを公開するようになりました

.dev に保存されており、そこで共同作業します。アーティファクトは実際のページとしてレンダリングされます。コメントはアーティファクトに保存されます。バージョンは自動的に追跡されます。 7 人のチーム全員がそれを使用します。
「私たちはこれをマークダウンと HTML アーティファクトの共有と共同作業に使用しています。この奇妙な新しいエージェント コーディング時代において、よりシームレスに共同作業するのに役立ちました。私たちの 7 人チーム全員がこれを使用しています。」 — マックス・ウッド、Indigo Engineering CEO
ワークフローがdisplay.devに移行すると、チームが予期していなかった3つの場所で成果が現れました。
新機能のイテレーションの高速化と大幅な再設計。 「display.dev のおかげで、私たちが泥沼にはまっていた新機能や大規模な再設計プロジェクトをより迅速に繰り返すことができました。」
Google ドキュメントが予期せず終了しました。 「これは、Google ドキュメントから移行するための予想外の優れた方法でもありました。なぜなら、当社のドキュメントの多くは現在マークダウンにネイティブであり、可視性とコラボレーションのためにそれらを使いやすい Google ドキュメント形式に変換することが非常に面倒になっていたからです。」
Anthropic エンタープライズ プランへの強制アップグレードの延期。 「また、これは、Claude Design の成果物で共同作業するために、おそらく強制されていたであろう Anthropic エンタープライズ プランへの切り替えを遅らせるのにも役立ちました。」
エージェントが作業を作成する場合、チームは出力をレビューして共同作業するための中立的な共有面を必要とします。そのための共有ワークスペースにより、どのツールが生成したシートでも全員にシートを購入するというプレッシャーがなくなりました。
display.dev がどのように改善できるかについて
完璧な新製品はありません。私たちはマックスに何がまだ足りないのか尋ねました。彼の答えは、ギャップについてだけでなく、私たちの働き方についても多くを語っています。
「display.dev は、私たちにとって完璧でエレガントなソリューションのようなものです。新しいソフトウェア製品には当然予想される問題もありましたが、問題が発生したときの display.dev チームの対応にはとても感謝しています」

フィードバックを共有しました。」 — マックス・ウッド、Indigo Engineering CEO
エージェントが制作を行い、人間がステアリングを握る場合でも、仕事にはやはり居場所が必要です。 display.dev は、エージェントが生成した HTML とマークダウンに実際のページとしてレンダリングする場所を提供し、チームのコメントを保持し、すべての変更をバージョン管理します。組織全体に対して単一の定額料金で行われます。 Indigo は、7 人全員が 1 つの面で共同作業しながら、再設計、内部ドキュメント、およびクロードの出力を実行します。エージェント時代のどのチームでも同じことができます。
AI が生成したアーティファクトのゲート付きパブリッシング。
© 2026 Display.dev 製品 仕組み ユースケース 価格 お客様 ブログ ニュース サポートについて サイトマップ Claude Code Cursor Codex Agent プラットフォーム オープンソース モデル すべての統合 → ソリューション HTML ファイルの共有 データ ダッシュボード HTML プレゼンテーション AI レポート すべてのソリューション → 代替案 vs Vercel vs GitHub Pages vs Cloudflare vs GitBook すべての代替案 → 法的プライバシー規約 セキュリティ Cookie ポリシー Cookie 設定 display.dev の Cookie
私たちは単一の分析ツール (PostHog) を使用して、訪問者がどのページを読んで、サイト内をどのように移動したかを確認します。広告Cookie、プロファイリング、データ販売はありません。フッターからいつでも気が変更できます。

## Original Extract

Indigo Engineering's 7-person team replaced Slack screenshots of their AI-generated HTML with display.dev: rendered, versioned, commentable pages. Faster redesigns, a quiet exit from Google Docs and a deferred Anthropic enterprise upgrade.

How Indigo Engineering gave their agents' output a home – Display.dev Display.dev / display █ How it works Use cases Pricing Docs News Sign in Get started free How it works Use cases Pricing Docs News Theme Sign in Get started free How Indigo Engineering gave their agents’ output a home.
Indigo Engineering builds straightforward software and data products that help progressive organizers build power in their communities. The team of seven is scattered across the US and ships two products: Matchbook, a voter-contact CRM with machine-learning optimization, and Matter, a voter-file enrichment tool.
Like most teams shipping software in 2026, Indigo’s engineers and designers lean heavily on coding agents. Those agents produce a constant stream of HTML and markdown: design output, reports, specs, internal docs. The work gets generated quickly. Getting it in front of teammates to review, comment on and revise was where things broke down.
Before: Slack threads, HTML screenshots and a versioning nightmare
“Before display.dev, we were sharing artifact files with each other in Slack and putting our comments there. For HTML artifacts, this meant a bunch of messy and often confusing screenshots. It also meant several conversations were mushed into the same thread, versioning was a nightmare, and tracking down changes or discussions was very annoying.” — Max Wood, CEO, Indigo Engineering
The team had seen where this leads and didn’t like the alternative either:
“We might have eventually started using git to track artifacts, but that would have still been annoying to render, we’d need constant PRs for revisions, and threads would be hard to follow.”
After: a shared home for agent output
Indigo now publishes its markdown and HTML artifacts to display.dev and collaborates on them there. The artifacts render as real pages. Comments live on the artifact. Versions are tracked automatically. The whole seven-person team uses it.
“We use it to share and collaborate on markdown and HTML artifacts, and it’s helped us work together more seamlessly in this weird new agentic coding era. Our whole 7-person team uses it.” — Max Wood, CEO, Indigo Engineering
Once the workflow moved to display.dev, the payoff showed up in three places the team didn’t expect to feel it.
Faster iteration on new features and a major redesign. “display.dev has helped us iterate much more quickly on new features and a big redesign project we’ve been mired in.”
An unexpected exit from Google Docs. “It’s also unexpectedly been a great way to move on from Google Docs since so many of our docs are now native to markdown and getting them into a friendly Google Doc format for visibility and collaboration had become a real chore.”
A deferred forced upgrade to Anthropic enterprise plans. “It’s also helped us delay a switch to Anthropic enterprise plans, which we probably would’ve been forced into so we could collaborate on Claude Design output.”
When agents produce the work, teams need a neutral, shared surface to review and collaborate on the output. A shared workspace for that removes the pressure to buy everyone a seat on whichever tool generated it.
On how display.dev could improve
No new product is perfect, and we asked Max what was still missing. His answer says as much about how we work as about the gaps:
“display.dev is sort of a perfect elegant solution for us. There have been issues that you’d reasonably expect from any new software product but we’ve been so grateful for the display.dev team’s responsiveness when we’ve shared feedback.” — Max Wood, CEO, Indigo Engineering
When agents do the producing and humans do the steering, the work still needs a home. display.dev gives agent-generated HTML and markdown a place to render as real pages, hold a team’s comments and version every change, at one flat price for the whole org. Indigo runs their redesign, their internal docs and their Claude output through it, with all seven people collaborating in one surface. Any team shipping in the agentic era can do the same.
Gated publishing for AI-generated artifacts.
© 2026 Display.dev Product How it works Use cases Pricing Customers Blog News About Support Sitemap For Claude Code Cursor Codex Agent platforms Open-source models All integrations → Solutions Share an HTML file Data dashboards HTML presentations AI reports All solutions → Alternatives vs Vercel vs GitHub Pages vs Cloudflare vs GitBook All alternatives → Legal Privacy Terms Security Cookie Policy Cookie settings Cookies on display.dev
We use a single analytics tool (PostHog) to see which pages visitors read and how they move through the site. No advertising cookies, no profiling, no data sold. You can change your mind any time from the footer.
