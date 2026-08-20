---
source: "https://meiert.com/blog/ai-and-tdd/"
hn_url: "https://news.ycombinator.com/item?id=49379618"
title: "AI-Driven Development Means Test-Driven Development"
article_title: "AI-Driven Development Means Test-Driven Development · Jens Oliver Meiert"
image: "https://d3rdtowr0c5lpf.cloudfront.net/media/social.png"
author: "speckx"
captured_at: "2026-08-20T20:17:35Z"
capture_tool: "hn-digest"
hn_id: 49379618
score: 1
comments: 0
posted_at: "2026-08-20T20:13:47Z"
tags:
  - hacker-news
  - translated
---

# AI-Driven Development Means Test-Driven Development

- HN: [49379618](https://news.ycombinator.com/item?id=49379618)
- Source: [meiert.com](https://meiert.com/blog/ai-and-tdd/)
- Score: 1
- Comments: 0
- Posted: 2026-08-20T20:13:47Z

## Translation

タイトル: AI 駆動開発とはテスト駆動開発を意味します
記事のタイトル: AI 駆動開発はテスト駆動開発を意味する · Jens Oliver Meiert
説明: TDD を採用するのに今ほど良い時期はないかもしれません。

記事本文:
Jens Oliver Meiert 書籍 アーカイブ 略歴 連絡先 検索 最新の著作物を使用してください: 最新の技術書、最新の非技術書、最新の最適化ツール、最新の防御ツール AI 駆動開発とはテスト駆動開発を意味します
2026 年 8 月 20 日に公開され、開発中、ai に提出されました。 (この投稿を Mastodon などで共有してください。) AI を使用して開発している場合は、テストから始める必要があります。あらゆる種類のモデルを使用した開発、保守、最適化から得た教訓が 1 つあるとすれば、これはこれです。これは、前のヒントで言い忘れたことです。テストから始めましょう。これを行う方法は、多くのモデルが事後テストに平手打ちするのを見てきたので、明示的にする必要があるかもしれませんが、まず作成または修正するものに対してテストを作成するよう依頼し、次に機能を開発するか修正を実装することです。洗い流して繰り返します。ただし、これは意図的に公開されており、文字よりも精神的に TDD (テスト駆動開発) を反映している可能性があります。テストから始めるとプロジェクトが保護されます。それぞれの機能が欠落しているか破損していることを確認します (そして、テストを確認して実行するか、少なくともサンプルを採取することをお勧めします)。 AI インスタンスが方向転換しないようにガードレールを設置します。AI インスタンスが行う作業はすべて、機能の開発または修正につながる必要があります。機能が実際に構築または修正されたことを確認します。それは自分自身を問題に集中させ続けます。現在、私は最初にフロントエンド開発者、2番目にWeb開発者、最後にプログラマーであるため、いくつかのプロセスをオフの立場から見ています。それでも、AI を使用することで、テスト駆動開発の価値が証明されました。 AI を扱っていて、プログラマーとしてはそれほど得意ではない場合は、テストから始めてみてください。テストには慣れる必要がありますが、「ただ出荷する」よりも簡単で安全です。私について
私は

イェンス (長名: イェンス・オリバー・マイヤート)、私はエンジニアリング リーダー、ゲリラ哲学者、インディーズ パブリッシャーです。私はさまざまな企業 (Google など) で技術リーダーおよびエンジニアリング マネージャーとして働いてきました。私は、Web およびツールの開発者 (コードの最適化、デジタル防御) を積極的に行っており、Web 標準 (HTML、CSS、WCAG など) の貢献者であり、本の著者 (O'Reilly、Frontend Dogma) でもあります。私は、Web 開発やエンジニアリング管理だけでなく、政治や哲学の分野でも、何かを試すのが大好きです。特に、私が信じている考えは、すべての人を大切にする場合にのみ健康になれるということです。ここmeiert.comでは、私の視点と経験のいくつかについて話します。 (慈善的に解釈しますが、批判的になり、自分の作品をより良くするためのフィードバックやアドバイスを共有してください。) 詳細

## Original Extract

There may have never been a better time to embrace TDD.

Jens Oliver Meiert Books Archive Biography Contact Search Go Use my latest work: latest tech book · latest non-tech book · latest optimization tool · latest defense tool AI-Driven Development Means Test-Driven Development
Published on Aug 20, 2026, filed under development , ai . ( Share this post , e.g., on Mastodon .) If you’re developing using AI, you need to start from tests. If there’s one lesson I learned from developing, maintaining, and optimizing using all sorts of models—something I missed calling out in previous tips !—, then it’s this. Start from tests. The way you do this—and you may need to be explicit as I’ve seen many models slap on tests after the fact—is by asking to first write tests for what you create or fix—and then develop the feature or implement the fix . Rinse and repeat—though this is intentionally open and may reflect TDD (test-driven development) more in spirit than in letter. Starting with tests protects your project: You confirm that the respective feature is missing or broken (and it’s recommended that you review and run the tests or, at least, take samples). You put a guardrail in place so that any AI instance doesn’t veer off—any work it does should lead to developing or fixing the feature. You confirm that the feature was indeed built or fixed. It keeps yourself centered on the problem. Now, I’m frontend developer first, web developer second, programmer last, and therefore look at some processes from an off-position. Still, working with AI has proven to me the value of test-driven development. If you’re working with AI, and aren’t a strong programmer, try starting with tests, too. While you will need to get used to testing, this is easier and safer than “just shipping.” About Me
I’m Jens (long: Jens Oliver Meiert), and I’m an engineering lead, guerrilla philosopher, and indie publisher. I’ve worked as a technical lead and engineering manager at various companies (e.g., Google); I’m an active web and tool developer (code optimization, digital defense), a contributor to web standards (like HTML, CSS, WCAG), and a book author (O’Reilly, Frontend Dogma). I love trying things—in web development and engineering management , but also in politics and philosophy , where I hold to one idea in particular: that we can only be well if we take good care of everyone. Here on meiert.com I talk about some of my perspectives and experiences. ( Interpret charitably but be critical—and share feedback and advice that makes my work better.) Learn More
