---
source: "https://krishna-modi12.github.io/frontend-design-pro/"
hn_url: "https://news.ycombinator.com/item?id=49396937"
title: "Show HN: Front end skill pack for AI agents, with machine-enforced quality gates"
article_title: "frontend-design-pro — the design rules an AI agent actually follows"
image: ""
author: "KrishnaModi12"
captured_at: "2026-08-22T06:21:49Z"
capture_tool: "hn-digest"
hn_id: 49396937
score: 1
comments: 0
posted_at: "2026-08-22T05:52:19Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Front end skill pack for AI agents, with machine-enforced quality gates

- HN: [49396937](https://news.ycombinator.com/item?id=49396937)
- Source: [krishna-modi12.github.io](https://krishna-modi12.github.io/frontend-design-pro/)
- Score: 1
- Comments: 0
- Posted: 2026-08-22T05:52:19Z

## Translation

タイトル: Show HN: 機械強制品質ゲートを備えた AI エージェント用のフロントエンド スキル パック
記事のタイトル:frontend-design-pro — AI エージェントが実際に従う設計ルール
説明: AI コーディング エージェントがフロントエンド コードを作成するときにロードするスキル パック。リクエストごとに 19 のスキルの 1 つをルーティングし、出力を 60 の機械チェックされた制約に保持し、11 のゲートすべてを通過した場合にのみ出荷されます。リクエストをルーティングし、ブラウザでチェックを実行します。

記事本文:
フロントエンドデザインプロ — AI エージェントが実際に従う設計ルール コンテンツへスキップ フロントエンドデザインプロ 数字 仕組み 壁 インストール v 14.11.1 GitHub AI フロントエンド スキル パック
エージェントが何を求めるかを変更します。
レジストリ経由。機械による強制。デフォルトでアンチスロップ — リクエストは 19 のスキルのうちの 1 つを正確にロードし、リクエストが書き込むものはすべて出荷前に 60 回のチェックに保留されます。
ルーターの背後には、参照深度の 371,905 個のトークンが存在します。これは、すべてのスキルが指すことができるマテリアルであり、リクエストがそこにルーティングされるまでロードされることはありません。パックはエージェント全体に渡されるわけではありません。ルーターはリクエストをトリガー キーワードと照合し、19 のスキルのうち 1 つとそれが宣言するコア ファイルを開き、残りはすべてディスク上に残ります。
より良いものを構築するためのより迅速な方法。今日試してみてください。
このパックが代わりに求めるもの
ラインを出荷する前に、ルート、予算、確認を行います。
まずはスペック。常に制約します。一度生成します。
ルーターはプロンプトを読み取り、それをすべてのスキルのトリガー キーワードと照合し、最も具体的に一致した 1 つと、宣言したコア ファイルをロードします。一致しません。推測ではなく明確な質問をします。以下で試してみてください。
書き込まれる内容は、機械チェックされた 60 の制約 (AST 17 + 正規表現 43) に保持されます。前半は TypeScript コンパイラー API を介して実行され、2 番目はパターンとして実行されます。小切手が実際に発動することを証明するための 10 の意図的な反例が存在します。
1 つのスキルとその宣言された依存関係。他には何も開きません。一般的なリクエストでは、利用可能な全深度に対して数千のトークンがロードされ、出荷されたサンプルはすべて tsc --strict でコンパイルされます。
以下にリクエストを入力してください。これにより、ページ用に作成された類似レジストリではなく、実際のレジストリが実行されます。
ヒーロー、価格設定、お客様の声、弁当グリッド、ロゴウォール、比較表、FAQ、CTA、フッター — その他

PTY の状態とオンボーディング
スキル スキル/ランディング ページ/SKILL.md
コア core/accessibility-baseline.md
コア core/validate-checklist.md
5,999 トークン — 利用可能な深度の 1.6 %
60 個のチェックで、エージェントが求め続けているデフォルトをキャッチします。
テキストとして照合するのではなく、TypeScript コンパイラ API を詳しく調べました。正規表現は、読み取っているツリーを理解していないと確認できない 8 つのルールです。
ソース上で正規表現として実行します。チェックがより速くなり、以下のこのページ独自のチェッカーの半分はブラウザで実行できます。
Acme、John Doe、user123、$99.99、「ワークフローのレベルを上げる」 — 他に何も指定されていない場合にエージェントが到達するデフォルト。
入り口での跳ね返り、弾力性、背中の緩和はありません。移行なし: すべて。モーションの縮小は機能的であり、文字列の言及ではありません。
Inter、Roboto、Arial、Poppins、DM Sans、Space Grotesk は表示面として禁止されています。フォールバック スタックでは問題ありませんが、ID を持ち込むことはありません。
100vh はモバイル ツールバーを無視します。 min-h-[100dvh] が唯一受け入れられるフォームであり、出荷されるすべてのサンプルでパターンによってチェックされます。
コンポーネントを貼り付けます。以下の両方のスニペットは、出荷前に scripts/test_constraints.py に対して両方向でクロスチェックされます。
人為的なマウント時の読み込み遅延はありません。実際の非同期からスケルトンを駆動します。
唯一のフォントとして禁止されている表示面はありません。フォールバックとしてのみ許可されます。
本文にグラデーション テキストはありません。計算された色は透明なので、コントラストを測定できません。
コンポーネント コードに任意の 16 進色は使用できません - OKLCH トークンのみ。
CSS トークン定義内に 16 進値がありません。
min-h-screen なし — 100vh はモバイル ツールバーを無視します。 min-h-[100dvh] を使用します。
プレースホルダー名や株価はありません。
プレースホルダー コメントは不要です。出力は完全である必要があります。
プレースホルダのブランド名は使用しません。セクターに合ったブランド名を作成してください。
進行状況のコピーでは、3 つのドットではなく省略記号文字が使用されます。

。
すべて遷移なし — アニメーション化するプロパティに名前を付けます。
アウトラインなしは、目に見えるフォーカス リングと組み合わせる必要があります。
43 の正規表現チェックのうち 15 がここで実行されます。残りはコンパイラまたはファイルシステムを必要とします。スイートの残りの半分は TypeScript AST を実行し、8 つのページスコープのルールは意図的に除外されています。これは、これらのルールが画面に関しては正しく、テキスト ボックスに貼り付けられたフラグメントに関しては間違っているためです。それらの ID は近似ではなく、ここでは省略されています。
パックの残りの部分が何を生み出すか
デフォルトのブランチを追跡します。ゲート付きのバージョンスタンプ付きアーカイブの場合: gh release download --repo Krishna-Modi12/frontend-design-pro --pattern '*.skill'
© 2026 フロントエンドデザインプロ · MIT ライセンス取得

## Original Extract

A skill pack an AI coding agent loads while it writes frontend code. It routes one of 19 skills per request, holds the output to 60 machine-checked constraints, and ships only when all 11 gates pass. Route a request and run the checks in your browser.

frontend-design-pro — the design rules an AI agent actually follows Skip to content frontend-design-pro Numbers How it works The wall Install v 14.11.1 GitHub AI frontend skill pack
Change what your agent reaches for.
Registry-routed. Machine-enforced. Anti-slop by default — a request loads exactly one of 19 skills, and everything it writes is held to 60 checks before it ships.
Behind the router sits 371,905 tokens of reference depth — material every skill can point into, none of it loaded until a request routes there. The pack is not handed to an agent whole: a router matches your request against trigger keywords, opens exactly one of 19 skills plus the core files it declares, and everything else stays on disk.
A faster way to build better. Try it today.
what this pack asks for instead
A route, a budget, a check — before a line ships.
Spec first. Constrain always. Generate once.
The router reads your prompt, matches it against every skill's trigger keywords, and loads exactly one — the most specific match — plus the core files it declares. No match, and it asks a clarifying question instead of guessing. Try it below.
What gets written is held to 60 machine-checked constraints (17 AST + 43 regex) — the first half walked through the TypeScript compiler API, the second run as patterns. Ten deliberate anti-examples exist to prove the checks actually fire.
One skill plus its declared dependencies — nothing else opens. A typical request loads a few thousand tokens against the full depth available, and every shipped example compiles under tsc --strict.
Type a request below. This runs the real registry — not a lookalike written for the page.
Heroes, pricing, testimonials, bento grids, logo walls, comparison tables, FAQ, CTAs, footers — plus empty states and onboarding
skill skills/landing-pages/SKILL.md
core core/accessibility-baseline.md
core core/validate-checklist.md
5,999 tokens — 1.6 % of the depth available
60 checks catch the defaults agents keep reaching for.
Walked through the TypeScript compiler API rather than matched as text — the eight rules a regex cannot see without understanding the tree it's reading.
Run as regular expressions over source. Faster to check, and the half this page's own checker below can run in your browser.
Acme, John Doe, user123, $99.99, 'Elevate your workflow' — the defaults an agent reaches for when nothing else is specified.
No bounce, elastic or back easing on an entrance. No transition: all. Reduced motion is functional, not a string mention.
Inter, Roboto, Arial, Poppins, DM Sans, Space Grotesk are banned as a display face — fine in a fallback stack, never carrying identity.
100vh ignores the mobile toolbar. min-h-[100dvh] is the only accepted form, checked by pattern on every shipped example.
Paste a component. Both snippets below are cross-checked against scripts/test_constraints.py in both directions before this ships.
No artificial mount-time loading delay — drive skeletons from real async.
No banned display face as the sole font — allowed only as a fallback.
No gradient text on body copy — the computed colour is transparent, so contrast cannot be measured.
No arbitrary hex colour in component code — OKLCH tokens only.
No hex value inside a CSS token definition.
No min-h-screen — 100vh ignores the mobile toolbar; use min-h-[100dvh].
No placeholder names or stock values.
No placeholder comments — the output must be complete.
No placeholder brand names — invent one that fits the sector.
Progress copy takes the ellipsis character, not three dots.
No transition-all — name the properties that animate.
outline-none must be paired with a visible focus ring.
15 of the 43 regex checks run here. The rest need a compiler or a filesystem: the other half of the suite walks a TypeScript AST, and eight page-scoped rules are deliberately excluded because they are right about a screen and wrong about a fragment pasted into a text box. Their IDs are absent here rather than approximated.
what the rest of the pack produces
Tracks the default branch. For the gated, version-stamped archive instead: gh release download --repo Krishna-Modi12/frontend-design-pro --pattern '*.skill'
© 2026 frontend-design-pro · MIT licensed
