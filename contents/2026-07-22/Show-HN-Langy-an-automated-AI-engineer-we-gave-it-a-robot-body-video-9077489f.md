---
source: "https://langwatch.ai/blog/introducing-langy-your-automated-ai-engineer"
hn_url: "https://news.ycombinator.com/item?id=49007760"
title: "Show HN: Langy, an automated AI engineer (we gave it a robot body) [video]"
article_title: "Introducing Langy: Your Automated AI Engineer"
author: "jangletown"
captured_at: "2026-07-22T14:56:11Z"
capture_tool: "hn-digest"
hn_id: 49007760
score: 2
comments: 0
posted_at: "2026-07-22T14:52:14Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Langy, an automated AI engineer (we gave it a robot body) [video]

- HN: [49007760](https://news.ycombinator.com/item?id=49007760)
- Source: [langwatch.ai](https://langwatch.ai/blog/introducing-langy-your-automated-ai-engineer)
- Score: 2
- Comments: 0
- Posted: 2026-07-22T14:52:14Z

## Translation

タイトル: 番組 HN: 自動 AI エンジニア、ランジー (ロボットの体を与えました) [動画]
記事のタイトル: Langy の紹介: 自動化された AI エンジニア
説明: Langy は、LangWatch 内の自動化された AI エンジニアです。トレースを読み取り、シナリオのテストと評価を作成し、PR を開き、CI 内のすべての修正を証明します。
HN テキスト: ここの創設者。 Langy は、私たちのプラットフォームである LangWatch 内に住む AI エンジニアです。本番環境のトレースを読み取り、見つかった問題に対するシナリオ テストと評価を作成し、リポジトリでプル リクエストを開き、CI でシミュレーションを実行して修正を証明します。人間はまだ合体します。 Hugging Face の Reachy ロボットは、私たちが打ち上げを計画していた週に到着したため、打ち上げに使用したいと考えていました。そこで、実際にプラットフォーム上に存在する Langy に接続し、独自のカスタマー サポート音声エージェントをテストするよう依頼しました。ビデオでは、トレースなどを読み取るエージェント テストの書き込みを確認できます。 LangWatch はオープン ソース プラットフォーム (github.com/langwatch/langwatch) であり、それが駆動するシミュレーション テスト ライブラリである Scenario もオープン ソースです (その部分を自分で実行したい場合は、https://github.com/langwatch/scenario )。 Langy が内部でどのように動作するかについての完全な記事は、今週後半に公開される予定です。ロボットがどのように駆動するのか、実際と演出されたもの、レイテンシ、あるいはロボットの周りのハーネス全体が実際にどのように動作するのかについて詳しく知ることができてうれしく思います。

記事本文:
Langy の紹介: 自動化された AI エンジニア
プラットフォーム エージェント AI テスト エージェントに対して現実的なユーザー シナリオを実行して、運用前に問題を発見します。 LLM 評価 応答の品質と精度を測定して、運用に耐えられるエージェントを出荷します。 LLM 可観測性 エージェントのすべてのステップを追跡し、運用環境を完全に可視化してコストと遅延を監視します。 AI ガバナンス 新しい すべてのモデル、キー、ツールを管理します。予算、ルーティング ポリシー、および完全な監査証跡を備えた仮想キー プロンプト管理 完全な履歴と GitHub 同期を含む、コードとしてのバージョン、デプロイ、および A/B テスト プロンプト 音声 AI 顧客と会話する前に、音声 AI エージェントを大規模にテストおよびシミュレート LLM レッドチーム AI エージェントの安全性とセキュリティのギャップを明らかにする模擬攻撃 クロード コードの使用状況を追跡 クロード コード、コーデックス、およびすべてのコーディング エージェントの完全なトレース履歴とトークン支出 エンタープライズ カスタマー ブログ ドキュメント 料金 予約デモ サインイン サインアップ ← Langy を紹介するブログ: 自動化された AI エンジニア
Langy は、LangWatch 内の完全なコーディング エージェントです。実稼働トレースを読み取り、シナリオ テストと評価を作成し、コードベースでプル リクエストを開き、CI の変更を証明します。
本日、LangWatch 内で動作する AI エンジニアである Langy を立ち上げます。これは完全なコーディング エージェントです。本番環境のトレースを読み取り、見つかったものに対してシナリオ テストと評価を書き込み、GitHub に接続し、実際のコードベースでプル リクエストを開き、CI に書き込んだシミュレーションを実行して変更を証明します。評価結果がすでに保存されている場所にアクセスすると、緑色のチェックが付いた差分が返されます。
私たちが最初に行ったことは、Langy を MCP 経由でロボットに接続し、独自のカスタマー サポート音声エージェントをテストするように依頼したことです。シナリオが始まりました

シミュレーションそのもの。
AI は、これまで触れてこなかった作業にソフトウェアを押し込みました。エージェントは法律業務、医療業務、金融業務を行っており、エージェントを構築するエンジニアは弁護士、医師、銀行家ではありません。ソフトウェア チームにとってドメイン エキスパートの重要性はこれまで以上に高まっていますが、ほとんどのチームではまだ変更に貢献できません。その作業を行うのに十分な価値のあるエージェントは複雑で、その背後に評価パイプラインとシナリオ スイートがあり、それを改善するには、コードに触れ、それを適切に実行する AI エンジニアリングの専門知識が必要です。
私たちのプラットフォームも含め、この市場のすべてのプラットフォームがその協力を約束しました。実際には、ドメインの専門家は今でもスプレッドシートや Slack メッセージを通じて共同作業を行っており、エンジニアは他の人が試みようとするすべての仮説のボトルネックとなっていました。エンジニアもその仕事を望んでいません。他の人の即時調整を翻訳するために AI エンジニアになる人はいません。
ランジーが私たちの答えです。それを可能にするために変わったのは、コーディング エージェントが真に優れたものになったことと、LangWatch がエージェントのテスト、評価、改善がすでに行われるプラットフォームに成長したことです。エージェントが複雑になると、その作業はプラットフォームとコードベースという 2 つの場所に同時に存在することになり、Langy は両方の場所で活躍します。 PM またはドメインの専門家は、エージェントが何を別の方法で行うべきかを平易な言葉で説明します。 Langy はそれをシナリオのテストと評価に変え、それらを実行し、修正の草案を作成し、プル リクエストをオープンします。エンジニアはウィッシュを翻訳する代わりに差分を承認し、その PR 上で実行された CI には、Langy が作成したシミュレーションの合否が表示されるため、レビューは信頼性ではなく変更に関するものになります。
何千もの本番会話を読んで、エージェントが実際にどのように行動するかのパターンを見つけ、修正する価値のある問題に分類します。
評価を書く

ns およびシナリオ シミュレーションは、その動作を特定し、それを実行してエージェントがどこで問題を起こすかを確認します。
単一の不良トレースを失敗したテストとして再現し、保存しておきます。
GitHub に接続し、変更内容の説明とともにエージェントを修正するプル リクエストを開きます。
誰かがマージする前に、CI で独自のシミュレーションを実行して修正を証明します。
ループの残りの部分も処理します。つまり、2 つのモデルまたはプロンプトのバージョンを比較し、言語間でプロンプトを翻訳し、戻ってくる問題をキャッチするシグナルを設定します。
チームの全員がそこから何かを得る。 PM とドメインの専門家は、冗談のつもりで代理店を獲得します。彼らは問題を見つけ、プロンプトを改善し、ドメインが実際に必要とするシナリオを追加します。エンジニアは、Slack メッセージではなくレビュー可能なプル リクエストとしてドメインの専門知識を取得し、真に困難な問題に集中し続けることができます。そして、Langy はエンジニアのプラットフォーム上での作業も高速化します。つまり、1 つのトレースを見つけて、その背後にあるパターンを特定し、本番環境や先週の実験でエージェントが何を行ったかを説明するメトリクスや視覚化を、午後ではなく一目で確認できるようにします。
最後の部分は人間の能力を超えて拡大します。 Langy を数千の本番トレースに向けて、何が起こっているのかを尋ねます。Langy はトピックをクラスター化し、注目に値するパターンのシグナルを設定し、結果とそれぞれについて何をすべきかについての提案を返します。代わりに、エンジニアが手作業でいくつかのデータをサンプリングし、推定するという方法もありました。
今年は「AI エンジニア」という主張が目白押しですが、そのうちの 1 つに対しては懐疑的な反応を示すのが正しい反応です。この名前で出荷されるもののほとんどは、ダッシュボードの上部にあるチャット ボックスです。 Langy には手があり、実際のリポジトリに対するプル リクエストと、自身が作成したテストで実行されるグリーン CI の動作を示します。
今週はビデオやチュートリアルなどを公開します。

Langy の使用方法について説明します。乞うご期待。
私たちが気に入っている点は、すべての Langy の実行自体がエージェントの実行であり、他のものと同様に LangWatch 内でトレース、コストが計算され、評価されるため、Langy の動作を監視するために使用するのと同じツールを使用して監視します。
Langy は現在、LangWatch チームに展開されています。代理店での使用をご希望の場合は、デモを予約してください。セットアップも一緒に行います。
これを LangWatch を使用して本番環境に導入します。
エージェントを追跡し、評価を実行し、失敗を再現可能なテストに変換します。
エンジニアに相談する 無料で始める よくある質問
LangWatch AI を本番環境に出荷するチームのための LLM エンジニアリング プラットフォーム。

## Original Extract

Langy is an automated AI engineer inside LangWatch: it reads your traces, writes Scenario tests and evaluations, opens PRs, and proves every fix in your CI.

Founder here. Langy is an AI engineer that lives inside our platform, LangWatch. It reads your production traces, writes Scenario tests and evaluations for the problems it finds, opens a pull request on your repo, and proves the fix by running those simulations in CI. A human still merges. Reachy robot from Hugging Face arrived on the same week we were planing on launching it, so we wanted to use for the launch. So we it wired to our Langy which actually lives on the platform, and asked it to test our own customer-support voice agent. On the video you can see it writing the agent tests reading traces and everything. LangWatch is an open source platform (github.com/langwatch/langwatch) and Scenario, the simulation-testing library it drives, is open source too if you want to run that part yourself: https://github.com/langwatch/scenario . A full write-up on how Langy works under the hood is coming later this week. Happy to get into how it drives the robot, what is real vs staged, latency, or how the whole harness around it actually works.

Introducing Langy: Your Automated AI Engineer
Platform Agentic AI Testing Run realistic user scenarios against your agent to catch issues before production LLM Evaluation Measure response quality and accuracy so you ship agents that hold up in production LLM Observability Trace every agent step and monitor cost and latency with full production visibility AI Governance New Govern every model, key, and tool. Virtual keys with budgets, routing policies, and a full audit trail Prompt Management Version, deploy, and A/B test prompts as code, with full history and GitHub sync Voice AI Test and simulate your voice AI agents at scale, before they talk to customers LLM Red-teaming Simulated attacks that uncover safety and security gaps in AI agents Track your Claude Code Usage Full trace history and token spend for Claude Code, Codex, and every coding agent Enterprise Customers Blog Docs Pricing Book a demo Sign in Sign Up ← Blog Introducing Langy: Your Automated AI Engineer
Langy is a full coding agent inside LangWatch: it reads your production traces, writes Scenario tests and evaluations, opens the pull request on your codebase, and proves the change in your CI.
Today we are launching Langy, an AI engineer that works inside LangWatch. It is a full coding agent: it reads your production traces, writes Scenario tests and evaluations against what it finds, connects to your GitHub, opens a pull request on your actual codebase, and proves the change by running the simulations it wrote in your CI. You talk to it where your evaluation results already live, and what comes back is a diff with green checks attached.
The first thing we did with it: we plugged Langy into a robot over MCP and asked it to test our own customer-support voice agent. It kicked off the Scenario simulations itself.
AI has pushed software into work it never used to touch. Agents are doing legal work, health work, financial work, and the engineers building them are not lawyers, doctors, or bankers. The domain expert matters more to a software team than ever, and yet on most teams they still cannot contribute a change: any agent valuable enough to do that work is complex, with an evaluation pipeline and a scenario suite behind it, and improving it means touching code and having the AI engineering expertise to do it well.
Every platform in this market, ours included, promised that collaboration. In practice the domain experts still collaborate through spreadsheets and Slack messages, and the engineers became the bottleneck for every hypothesis anyone else wants to try. Engineers do not want that job either; nobody becomes an AI engineer to translate other people's prompt tweaks.
Langy is our answer. What changed to make it possible is that coding agents got genuinely good, while LangWatch grew into the platform where agent testing, evaluation, and improvement already happen. When your agents get complex, the work lives in two places at once, the platform and the codebase, and Langy is at home in both. A PM or domain expert describes what the agent should do differently, in plain language. Langy turns that into scenario tests and evaluations, runs them, drafts the fix, and opens the pull request. An engineer approves a diff instead of translating a wish, and the CI run on that PR shows the simulations Langy wrote, passing or failing, so the review is about the change rather than about trust.
Read through thousands of production conversations to find the patterns in how your agent actually behaves, and cluster them into the problems worth fixing.
Write the evaluations and Scenario simulations that pin that behavior down, then run them to see where the agent breaks.
Reproduce a single bad trace as a failing test you can keep.
Connect to your GitHub and open a pull request that fixes the agent, with the change explained.
Prove the fix by running its own simulations in your CI before anyone merges.
Handle the rest of the loop too: compare two model or prompt versions, translate prompts across languages, and set up the signals that catch a problem coming back.
Everyone on the team gets something out of that. PMs and domain experts get agency, pun intended: they find the problems, improve the prompts, add the scenarios their domain actually needs. Engineers get domain expertise landing as reviewable pull requests instead of Slack messages, and keep their focus for the genuinely hard problems. And Langy speeds engineers up on the platform too: finding that one trace, spotting the pattern behind it, pulling the metric or the visualization that explains what the agent has been doing in production or in last week's experiment, at a glance instead of an afternoon.
That last part scales further than a person can. Point Langy at a few thousand production traces and ask what is going on: it clusters the topics, sets up signals for the patterns worth watching, and comes back with findings and a proposal for what to do about each. The alternative was an engineer sampling a handful by hand and extrapolating.
"AI engineer" is a crowded claim this year, and skepticism is the right reaction to one more of them. Most of what ships under the name is a chat box on top of dashboards. Langy has hands, and it shows its work: a pull request on a real repo and a green CI run on tests it wrote itself.
This week we will be launching videos, tutorials, and more on how to use Langy. Stay tuned.
One detail we like: every Langy run is itself an agent run, traced, costed, and evaluated inside LangWatch like any other, so we watch Langy work with the same tools you will use to watch yours.
Langy is rolling out to LangWatch teams now. If you want it on your agent, book a demo and we will set it up with you.
Put this into production with LangWatch.
Trace your agents, run evaluations, and turn failures into repeatable tests.
Talk to an engineer Start free Frequently asked questions
LangWatch The LLM engineering platform for teams that ship AI to production.
