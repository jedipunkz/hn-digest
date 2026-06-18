---
source: "https://www.aaron-gustafson.com/notebook/2026-06-17-llm-biased-against-accessible-code/"
hn_url: "https://news.ycombinator.com/item?id=48587436"
title: "LLM biased against accessible code (Claude Code issue #56079)"
article_title: "LLM biased against accessible code (Claude Code issue #56079) :: Aaron Gustafson"
author: "robin_reala"
captured_at: "2026-06-18T16:12:46Z"
capture_tool: "hn-digest"
hn_id: 48587436
score: 1
comments: 0
posted_at: "2026-06-18T16:01:02Z"
tags:
  - hacker-news
  - translated
---

# LLM biased against accessible code (Claude Code issue #56079)

- HN: [48587436](https://news.ycombinator.com/item?id=48587436)
- Source: [www.aaron-gustafson.com](https://www.aaron-gustafson.com/notebook/2026-06-17-llm-biased-against-accessible-code/)
- Score: 1
- Comments: 0
- Posted: 2026-06-18T16:01:02Z

## Translation

タイトル: LLM はアクセシブルなコードに対して偏っています (クロード コードの問題 #56079)
記事のタイトル: アクセシブルなコードに対して LLM が偏っている (Claude Code issue #56079) :: Aaron Gustafson

記事本文:
アーロン・グスタフソン
このサイトはオープンな再設計中であるため、奇妙に見える場合はそれが原因です。
閉じる
LLM はアクセシブルなコードに対して偏っています (クロード コードの問題 #56079)
EstiShay は懸念すべきバグレポートを提出しました。彼らは、プロジェクトの要件ファイルで「WCAG 2.2 AA 最小値」が明示的に指定されている場合でも、Claude Code がアクセシビリティ修正をオプションとして扱っていることを発見しました。尋ねると、モデルは次のように説明します。
クロードは、たとえプロジェクト自身のルールに別の規定がある場合でも、アクセシビリティの修正を要件ではなくオプションのトレードオフとして扱います。これは、単に正しい ARIA パターンを知らないということとは異なります。これは、モデルが競合する優先順位をどのように考慮するかという価値観の問題です。
つまり…知識が不足しているわけではありません (コードをレビューするときにモデルは適切に機能します)。それは意識的な選択であり、優先順位付けの失敗です。私たちが何十年も人間が作ったソフトウェアでアクセシビリティの優先順位を下げてきたのと同じように、コーディング速度を追求するとアクセシビリティは「あると便利」なものになります。
モデルは私たちから学び、私たちのすべての間違いを繰り返しています。大規模に。
あなたの AI はアクセシビリティ テストに合格できますか?
Microsoft Buildで初めて発表されました
2026 年 6 月 2 日
アクセシビリティを事後に修正するのでは遅すぎる
ノートブックのエントリが投稿されました
2026 年 5 月 29 日
このようなコンテンツをもっと知りたい場合は、フォローしてください
Notist でこれまで行った場所と次に行く場所を確認してください
自分のコンテンツを Dev.to にクロスポストします
このサイトで見つけられるものは次のとおりです。
私の講演活動のリスト
私が行ったインタビューへのリンク
このサイトでは、XML (RSS) および JSONFeed 形式で複数のフィードを提供しています。
タイプミスや間違いを見つけましたか?えーっ、ごめんなさい。直してくれませんか？あなたはとても素晴らしいです。
このサイトのコンテンツは、クリエイティブ コモンズに基づいてライセンスされています。
このサイトがどのように作られたか興味がありますか?これをチェックしてください。
{% ifment.author %}
{{ 言及.著者名 }}
{%endif

%}
{{ 言及.コンテンツ |安全です }}
{{ メンション.published_human }}
パーマリンク :
{{ 言及.url }}
{% if mention.author.photo %}
{% endif %}

## Original Extract

Aaron Gustafson
This site is undergoing an open redesign , so if it looks strange, that’s why.
Close
LLM biased against accessible code (Claude Code issue #56079)
EstiShay filed a bug report that’s concerning. They found that Claude Code treats accessibility fixes as optional, even when the project’s requirements file explicitly specifies “WCAG 2.2 AA minimum.” When asked, the model explains:
Claude treats accessibility fixes as optional trade-offs rather than requirements, even when the project’s own rules say otherwise. That’s distinct from just not knowing the right ARIA pattern — it’s a values problem in how the model weighs competing priorities.
So… not a lack of knowledge (the model performs well when reviewing code); it’s a conscious choice, a prioritization failure. Accessibility becomes a “nice to have” in pursuit of coding speed, just like we’ve deprioritized it in human-made software for decades.
The models learned from us and are repeating all of our mistakes. At scale.
Can Your AI Pass the Accessibility Test?
First presented at Microsoft Build
on 02 Jun 2026
Fixing Accessibility After the Fact Is Too Late
Notebook entry posted
on 29 May 2026
For more content like this, follow me on
See where I’ve been & where I’m going next on Notist
I crosspost my content to Dev.to
Here’s what you can find on this site:
A List of My Speaking Engagements
Links to Interviews I’ve Given
This site offers multiple feeds in XML (RSS) and JSONFeed formats .
Did you spot a typo or error? Eesh, I’m sorry. Would you mind fixing it for me? You’re so awesome.
The content of this site is licensed under Creative Commons .
Interested in how this site was made? Check this out .
{% if mention.author %}
{{ mention.author.name }}
{% endif %}
{{ mention.content | safe }}
{{ mention.published_human }}
Permalink :
{{ mention.url }}
{% if mention.author.photo %}
{% endif %}
