---
source: ""
hn_url: "https://news.ycombinator.com/item?id=48751798"
title: "0/6 major aerospace documentation portals are AI Agent-ready"
article_title: ""
author: "priyanshu-j"
captured_at: "2026-07-01T19:32:24Z"
capture_tool: "hn-digest"
hn_id: 48751798
score: 2
comments: 0
posted_at: "2026-07-01T19:15:22Z"
tags:
  - hacker-news
  - translated
---

# 0/6 major aerospace documentation portals are AI Agent-ready

- HN: [48751798](https://news.ycombinator.com/item?id=48751798)
- Score: 2
- Comments: 0
- Posted: 2026-07-01T19:15:22Z

## Translation

タイトル: 0/6 の主要な航空宇宙ドキュメント ポータルが AI エージェント対応
HN テキスト: 私は、主要なドキュメント ポータルが AI エージェントをどの程度サポートしているかを評価するスコアリング エンジンである AeroScore の構築に時間を費やしました。つまり、この最初のバッチで私が評価した 6 つはどれも十分に準備されていません。方法論と基準は、AFDocs 仕様 (エージェント対応ドキュメントのオープンソース チェックリスト) と、航空宇宙分野では PDF が非常に重要視されているため、PDF を大量に使用するための 1 つの特定の拡張機能に基づいています。各サイトは、llms.txt の存在、robots.txt、および AFDocs リポジトリと aeroscore Web ページの両方で見つかるその他の多くの基準などの基準にわたって 0 ～ 100 のスコアが付けられます。主な調査結果:
-0/6 には llms.txt があります
-0/6 の URL バリアントをサポート
-最高スコアは ICAO の 69/100 でした。主要な航空宇宙文書ポータルがまだ AI エージェントに対応していないという事実は、この分野が完全に効率化されていないことを意味するため、重要です。 AI エージェントを使用すると、ディープ ドキュメンテーション ポータル全体のデータ分析を予想時間の数分の一で実行できるようになり、多くの分野に革命をもたらしました。航空宇宙分野でこれが欠けていると、生産性が失われます。リーダーボードは次の場所で公開されています: https://aeroscore.vercel.app AFDocs リポジトリは次のとおりです: https://github.com/agent-ecosystem/afdocs 私は航空宇宙工学の学部生で、自分の分野で AI がどのように実装され、どのように実装できるかに興味があります。これは aeroscore の v1 ですが、かなり荒削りなので、方法論に関するフィードバックをいただければ幸いです。将来のバージョンでは、改良されたルーブリックを使用してさらに多くのサイトを採点する予定です。

## Original Extract

I spent some time building AeroScore, a scoring engine that evaluates how well major documentation portals support AI agents. In short, none of the 6 that I evaluated in this first batch are sufficiently prepared. The methodology and criteria are based on the AFDocs Spec (an open-source checklist for agent-ready documentation) and one specific extension for PDF-heaviness since aerospace as a field places so much importance on PDFs. Each site gets scored 0-100 across criteria like llms.txt presence, robots.txt, and many others that can be found both on the AFDocs repo and at the aeroscore webpage. Key Findings:
-0/6 have an llms.txt
-0/6 support url variants
-the best score was ICAO with a 69/100 The fact that major aerospace documentation portals are not AI agent ready yet is important because it means the field is not being fully efficient. AI agents allow the analysis of data across deep documentation portals to occur in fractions of the expected time, and this has revolutionized many fields. The lack of this in aerospace means a loss of productivity. The leaderboard is live at: https://aeroscore.vercel.app The AFDocs repo is at: https://github.com/agent-ecosystem/afdocs I'm an incoming undergraduate aerospace engineering student interested in how AI is and can be implemented in my field. This is v1 of aeroscore, it's pretty rough and I'd appreciate feedback on the methodology greatly. In future versions, I'll be scoring more sites with a refurbished rubric.

