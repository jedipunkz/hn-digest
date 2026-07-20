---
source: "https://firerun.io/blog/netlify-per-member-ai-spend-limits-2026/"
hn_url: "https://news.ycombinator.com/item?id=48982964"
title: "Netlify Adds a Kill Switch for Runaway AI Agent Spend"
article_title: "Netlify Adds a Kill Switch for Runaway AI Agent Spend — Firerun"
author: "gbourne"
captured_at: "2026-07-20T19:25:49Z"
capture_tool: "hn-digest"
hn_id: 48982964
score: 1
comments: 0
posted_at: "2026-07-20T18:34:28Z"
tags:
  - hacker-news
  - translated
---

# Netlify Adds a Kill Switch for Runaway AI Agent Spend

- HN: [48982964](https://news.ycombinator.com/item?id=48982964)
- Source: [firerun.io](https://firerun.io/blog/netlify-per-member-ai-spend-limits-2026/)
- Score: 1
- Comments: 0
- Posted: 2026-07-20T18:34:28Z

## Translation

タイトル: Netlify、AI エージェントの暴走に対応するキル スイッチを追加
記事のタイトル: Netlify、AI エージェントの暴走に対するキル スイッチを追加 — Firerun
説明: Netlify では、エージェントが 1 か月を費やすという苦情を受けて、チーム オーナーがチームごとだけでなくチーム メンバーごとに AI クレジットの支出を制限できるようになりました。

記事本文:
Netlify が AI エージェントの暴走に対応する Kill Switch を追加 — Firerun : Plausible のインストール チェックはヘッドのみをスキャンします。 --> AI エージェント用の機械可読サイト インデックス (llms.txt) コンテンツへスキップ 最終更新日 2026 年 7 月 20 日 FIRERUN 。ビルダー向けの BaaS およびデプロイ プラットフォームのニュースルーム
Netlify、AI エージェントの暴走に対応する Kill Switch を追加
Netlify では、エージェントが 1 回のプロンプトで 1 か月分の割り当てを使い切ってしまうという苦情を受けて、チーム オーナーが AI クレジットの支出をチームごとだけでなくチーム メンバーごとに制限できるようになりました。
Firerun 社説 · 2026 年 7 月 20 日
これが単なる機能ではなく修正である理由
Netlify では、プロ プランのチーム オーナーが、チーム全体に 1 つの制限を設定するのではなく、個々のチーム メンバーがエージェント ランナーを通じて消費できる AI クレジットの数を制限できるようになりました。この変更は、2026 年 7 月 15 日の時点で [全般] > [AI 有効化] の [チーム設定] にあり、メンバーごとのクレジット上限をオーバーライドで追加します。これには、メンバーの制限を 0 クレジットに設定して、エージェント ランナーのアクセスを完全に遮断することが含まれます ( Netlify 変更ログ )。
これまで、唯一の支出制御はチーム全体でした。チーム オーナーがエージェント ランナーと AI ゲートウェイを合わせて 1 つのクレジット上限を設定し、チーム全体がその上限に達すると、すべてのメンバーのエージェントの実行が停止されました (Netlify 変更ログ、2026 年 3 月)。これによって法案は守られたが、チームの割り当て分を誰かが使う前に一人が食べてしまうことを止めることはできなかった。
新しいメンバー クレジットの使用制限は、チームの上限の下に設定されます。チーム所有者はメンバーごとのデフォルトの制限を設定し、特定の人に対してそれを上書きします。つまり、パワー ユーザーには上限を高くし、必要なスペースが少ない人には上限を低くし、エージェントをまったく実行すべきではない人にはハード 0 を設定します。 Netlify は AI 推論使用状況ダッシュボードも作り直したので、所有者はどのメンバーが実際に支出を推進しているのかを確認できるようになりました (Netlify は

cs)。
これが単なる機能ではなく修正である理由
Netlify のクレジット システムは、AI モデルの使用量の 1 ドルを 180 クレジットに変換し、エージェント ランナーのコンピューティングは、エージェントが行うモデル呼び出しとは別に測定されます。モデルに依存した価格設定とは、1 つの高価なモデルに対する 1 つの高価なプロンプトが、固定価格のシートでは決して消費されない速度でクレジットを消費する可能性があることを意味します。今回のアップデートまで、チームはそれを 1 つのアカウントに抑える方法がありませんでした。 Netlify 独自のサポート フォーラムには、単一のエージェントの実行で共有残高が枯渇した後に本番環境のデプロイがブロックされたチームによる最近のスレッドが複数あります ( Netlify サポート フォーラム )。この苦情はフォーラム以外にも現れています。Trustpilot のレビュー担当者は、エージェントが「1 回のプロンプトで私の使用クレジットをすべて吸い上げ、サイトを 1 か月間無効にする前に変更さえも生成しなかった」と述べています (Trustpilot)。
それが今回のリリースで埋まるギャップです。メンバーごとの上限は、1 つの暴走プロンプトがチーム全体の予算ではなく、1 人の予算に合わせた壁にぶつかることを意味します。
新しいコントロールはオプトインおよび手動です。チーム オーナーはリスクに気づき、チーム設定を開いて個別の制限を設定する必要があります。メンバーごとのデフォルトの上限はなく、1 人が不均衡なシェアを使い果たしそうになったときに自動的に警告することもありません。 Netlify のフォーラムでは、このリリースでは触れられていない別の進行中のバグ パターンも示されています。「クレジットが不足しています」と報告しているチームは、ダッシュボードに利用可能な未使用のクレジットがまだ表示されているにもかかわらず、デプロイをブロックする「クレジット不足」バナーを報告しています。この不一致については、今月くらいにいくつかの投稿が投稿されました ( Netlify サポート フォーラム )。支出管理は、チェックしている残高が正確である場合にのみ役に立ちます。
私たちの見解: メンバーごとの制限は問題に対して適切な形状であり、エージェントのコストはアクターごとのリスクであり、コントロールはチーム レベルだけでなくそのレベルで存続する必要があります。ただし送料

チーム設定に組み込まれたオプトイン設定として、所有者には、おそらくまだ遭遇したことのない障害モードを予測する負担がかかります。より安全なデフォルトでは、所有者が明示的に上限を設定するまで、すべての新しいメンバーに自動的に上限が設定されます。これは、Netlify が他の課金に敏感な機能を制限する方法をすでに決定しているのと同じ本能です。
X LinkedIn HN Reddit リンクをコピー 読み続ける
Netlify 関数が型指定された構成の Magic ファイル名を廃止
Netlify は関数の設定を型指定されたコードに移動し、マジック ファイル名を置き換えたので、コーディング エージェントは設定を確実に検出して編集できます。
Genkit の Go SDK が Anthropic プラグインと Ollama プラグインで安定化
Google の Genkit Go SDK は 6 月 18 日に最初の安定リリースをリリースし、API の下位互換性を確保し、Claude、Vertex AI、および Ollama プラグインを出荷しました。
Genkit 1.39 はチャットベータ版を廃止し、プロダクションエージェントを導入
Genkit JS v1.39.0 では、ベータ版 Chat API が削除され、Firestore セッション永続性を備えた実稼働マルチターン エージェント フレームワークであるdefineAgent が同梱されています。
受信箱に新しい記事が届く
新しい記事が投稿された場合にのみメールでお知らせします。いつでも購読を解除してください。
概要 規約 プライバシー RSS llms.txt Cookie 設定 GitHub
© 2026 ファイアラン。独立した報道。私たちがカバーするプラットフォームとは無関係です。
分析には Cookie を使用します。あなたが受け入れるまで、彼らはオフのままです。プライバシー

## Original Extract

Netlify now lets Team Owners cap AI credit spend per team member, not just per team, after complaints of agents burning a month

Netlify Adds a Kill Switch for Runaway AI Agent Spend — Firerun : Plausible's installation check only scans the head. --> Machine-readable site index (llms.txt) for AI agents Skip to content Last Updated July 20, 2026 FIRERUN . The BaaS & deploy-platform newsroom for builders
Netlify Adds a Kill Switch for Runaway AI Agent Spend
Netlify now lets Team Owners cap AI credit spend per team member, not just per team, after complaints of agents burning a month's allotment in one prompt.
By Firerun Editorial · Jul 20, 2026
Why this is the fix, not just a feature
Netlify now lets Team Owners on Pro plans cap how many AI credits an individual team member can burn through Agent Runners, instead of only setting one limit for the whole team. The change, live in Team Settings under General > AI Enablement as of July 15, 2026, adds a per-member credit ceiling with overrides, including setting a member’s limit to 0 credits to shut off their Agent Runners access entirely ( Netlify changelog ).
Until now, the only spend control was team-wide: a Team Owner set one credit cap for Agent Runners and AI Gateway combined, and once the whole team hit it, every member’s agent runs stopped ( Netlify changelog, March 2026 ). That protected the bill but did nothing to stop one person from eating the team’s entire allotment before anyone else got to use it.
The new Member Credit Usage Limit sits underneath that team cap. A Team Owner sets a default per-member limit, then overrides it for specific people: a higher ceiling for a power user, a lower one for someone who needs less room, or a hard 0 for someone who shouldn’t be running agents at all. Netlify also reworked the AI inference usage dashboard so owners can see which members are actually driving the spend ( Netlify docs ).
Why this is the fix, not just a feature
Netlify’s credit system converts $1 of AI model usage into 180 credits, and Agent Runners compute is metered separately from the model calls an agent makes. Model-dependent pricing means one expensive prompt to one expensive model can consume credits at a rate a fixed-price seat never would, and until this update, a team had no way to contain that to one account. Netlify’s own support forum has multiple recent threads from teams whose production deploys got blocked after a single agent run drained a shared balance ( Netlify Support Forums ). The complaint shows up outside the forums too: a Trustpilot reviewer described an agent that “sucked all of my usage credits in one prompt… and didn’t even generate the changes before disabling my site for one month” ( Trustpilot ).
That’s the gap this release closes. A per-member cap means one runaway prompt now hits a wall sized for one person’s budget, not the whole team’s.
The new control is opt-in and manual. A Team Owner has to notice the risk, open Team Settings, and set individual limits, there’s no default per-member cap and no automatic alert when one person is about to blow through a disproportionate share. Netlify’s forums also show a separate, ongoing bug pattern this release doesn’t touch: teams reporting a “run out of credits” banner blocking deploys while their dashboard still shows unused credits available, a mismatch several posted about as recently as this month ( Netlify Support Forums ). Spend controls only help if the balance they’re checking against is accurate.
Our take: per-member limits are the right shape for the problem, agent cost is a per-actor risk and the control should live at that level, not just the team level. But shipping it as an opt-in setting buried in Team Settings puts the burden on the owner to anticipate a failure mode they’ve likely never hit yet. A safer default would cap every new member automatically until an owner explicitly raises it, the same instinct that already governs how Netlify gates other billing-sensitive features.
X LinkedIn HN Reddit Copy link Keep reading
Netlify Functions Retire Magic Filenames for Typed Config
Netlify moved Functions configuration into typed code, replacing magic filenames, so coding agents can discover and edit settings reliably.
Genkit's Go SDK Goes Stable with Anthropic and Ollama Plugins
Google's Genkit Go SDK hit its first stable release June 18, committing to backward API compatibility and shipping Claude, Vertex AI and Ollama plugins.
Genkit 1.39 Retires the Chat Beta, Lands Production Agents
Genkit JS v1.39.0 removes the beta Chat API and ships defineAgent, a production multi-turn agent framework with Firestore session persistence.
Get new articles in your inbox
We'll only email you when a new article drops. Unsubscribe anytime.
About Terms Privacy RSS llms.txt Cookie settings GitHub
© 2026 Firerun. Independent coverage. Not affiliated with the platforms we cover.
We use cookies for analytics. They stay off until you accept. Privacy
