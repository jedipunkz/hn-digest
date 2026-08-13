---
source: "https://github.com/gabert/buzzword-driven-development"
hn_url: "https://news.ycombinator.com/item?id=49290686"
title: "The Buzzword-Driven Development Manifesto (AI Edition)"
article_title: "GitHub - gabert/buzzword-driven-development · GitHub"
author: "gabert"
captured_at: "2026-08-13T19:52:13Z"
capture_tool: "hn-digest"
hn_id: 49290686
score: 2
comments: 0
posted_at: "2026-08-13T19:19:54Z"
tags:
  - hacker-news
  - translated
---

# The Buzzword-Driven Development Manifesto (AI Edition)

- HN: [49290686](https://news.ycombinator.com/item?id=49290686)
- Source: [github.com](https://github.com/gabert/buzzword-driven-development)
- Score: 2
- Comments: 0
- Posted: 2026-08-13T19:19:54Z

## Translation

タイトル：流行語主導の開発宣言（AI編）
記事のタイトル: GitHub - gabert/バズワード駆動開発 · GitHub
説明: GitHub でアカウントを作成して、gabert/流行語主導の開発に貢献します。

記事本文:
GitHub - ガバート/バズワード駆動開発 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ガベール
/
バズワード主導の開発
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット LICENSE LICENSE README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
バズワード主導の開発宣言
(v4.0 — Microservices Edition (2016)、Blockchain Edition (2018)、Metave に置き換わります。

rse版（2021年、未読）、および短いが高価なWeb3版（2022年、NFTとして鋳造、現在の所有者不明））
バズワード主導の開発は新しいものではありません。流行語だけです。
私たちの長老たちは、ログイン ページを 42 のマイクロサービスに分割しました。彼らの長老たちはデータベースをブロックチェーン上に置きました。私たちは彼らの犠牲、つまり、焼け落ちた予算、放棄されたリポジトリ、彼らが説明した製品よりも長く生き残った Medium の記事に敬意を表します。私たちは聖火を前に運びます。
今日のバズワードはAIです。以下のすべては、その 1 つの事実から導かれます。
私たちは開発者です。私たちはかつてソフトウェアを書いていました。今、私たちはインテリジェンスを統合しています。私たちはこのことについて怒っていません。怒りにはエネルギーが必要ですが、私たちのエネルギーはローカルの開発環境を再起動するために確保されています。私たちは、私たちのクラフトが UI を備えたプロンプトに変化するのを、軽い仕事上の疲れを感じながらただ観察しているだけです。
この取り組みを通じて、私たちは次のことを大切にするようになりました。
付加価値ではなく AI の追加。誰も求めなかった機能、誰も理解していないモデルを活用し、誰も抱えていなかった問題を解決します。発送してください。在庫が上がっています。
コードを記述するよりも API をラップする。私たちの製品は、他人のモデルの上の薄い層であり、他人の GPU の上の薄い層であり、非常にオーバーブッキングされている 1 つのチップ ファウンドリの上の薄い層です。これを私たちは「独自技術」と呼んでいます。
バージョン管理よりも気楽です。どのモデルのバージョンが生産されているのかはわかりません。モデルプロバイダーも同様です。先週の火曜日に変更され、チャットボットは少し失礼になりました。これを「能力アップグレード」といいます。
理解を超えたエージェント。 3 月以降、誰もコードベースを読んでいません。 AI が PR を書き、AI がレビューし、AI が PR を承認します。私たちは、親が十代の若者の誕生日パーティーに出席するのと同じように、このループの中にいます。法的に義務付けられているにもかかわらず、静かに無視されています。
すべての機能は AI を活用する必要があります。ダークモードの切り替えがagになりました

魅惑的。それは闇についての理由付けです。闇について深い意見を持っています。切り替えごとに 0.04 ドルかかります。
チャットボットにできるのであれば、それはチャットボットでなければなりません。ユーザーは「日付で並べ替え」ボタンを求めていました。彼らは自然言語で並べ替えを要求できる会話型インターフェイスを受け取り、70% の確率でそれを受け取りました。進捗。
決定論的なコードはレガシー コードです。 if文は卑怯者のためのものだ。実際のエンジニアはブール値を言語モデルに送信し、その判定を待ちます。時々詩を返します。その場合は本番環境で対応します。
デモは製品です。一度はうまくいきました。ステージ上。創設者のラップトップ上。 1 つのプロンプトを使用して、3 週間テストしました。それが生産です。デモ後のすべては「硬化」しており、この段階は永遠に続きます。
幻覚はロードマップ項目です。バグではありません。「既知の制限」です。これは、AI が正確にそのような動作をしないことを約束する部分の直後にドキュメントに記載されています。
プロンプトエンジニアリングはコンピュータサイエンスです。私たちは、「あなたは役に立つアシスタントです。JSON のみで応答してください。お願いします。」と毎日書くことができるように、アルゴリズムを学習するのに 4 年間費やしました。テキストボックスに入力します。 Caps ロックは耐荷重性があります。
⏳休憩。このマニフェストが作成されています。あなたが経験している一時停止は、モデル思考ではありません。それは、ダムのように、次のトークンを放出するために十分な電力が送電網に蓄積されるのを待っているのです。読み込みスピナーはアニメーションではありません。燃料計です。待ってください。まもなく風が強くなるでしょう。
コンテキスト ウィンドウは新しいデータベースです。スキーマ?移住？いいえ、本番データベース全体をプロンプトに貼り付けて祈ります。サイズが合わなくなったら、より大きなモデルを購入します。これが私たちのスケーリング戦略です。それは私たちのアーキテクチャ全体でもあります。
四半期ごとにスタックを再構築する必要があります。壊れたからではなく、新しいモデルが発売されたからです

編集され、古いものは現在「薬剤投与前」です。 6 か月分のコード。変更ログによって廃止されました。
評価はスクリーンショットです。私たちはチャット ウィンドウで 3 つの質問をして新しいモデルをテストしました。賢そうだった。配備されました。指標?指標は、それが賢く見えるかどうかです。
はい、それでは会議です。毎週行われる「AI戦略同期」では、端末を開いたことがない人が「エージェントをもっと活用すべき」と説明する。私たちはうなずきます。まさにその瞬間、私たちは気配りがあるように見せるためにエージェントを利用しています。誰もが幸せです。何も決まっていません。エージェントには数分かかります。議事録はアクションアイテムを幻覚させます。誰かがそれを完成させます。システムは機能します。
呪文による雇用の確保。私たちは同時に、AI が私たちに取って代わることを告げられ、AI が書いたものを修正するよう求められます。私たちを時代遅れにするマシンをデバッグできるのは私たちだけです。私たちは矛盾を指摘しないことにしました。それは堅実なギグだ。
定期的に、チームは反省し、LLM に反省の要約を依頼し、別の LLM に要約の要約を依頼して、LLM のみが検索でき、ほとんど正確に検索できるナレッジ ベースにそれをファイルします。
私たちのコードベースは 40% が生成、60% が接着剤、そして 100% が「AI ネイティブ」です。その根底には、LLM と呼ばれる奇妙に配線されたアーキテクチャがあります。しかし、どういうわけか機能します。私たちのテスト スイートは、「これはあなたにとって適切ですか?」と尋ねるプロンプトです。私たちのドキュメントは、モデルによって、モデルについて、モデルのために書かれました。そのどこかに TODO: 2024 年からの発売前に削除するという項目があります。
製品はほとんど動作します。ユーザーはほとんどの場合対処します。モデルプロバイダーからの請求書は天気予報と同じように毎月届き、今や小さな島国の GDP に匹敵します。財務には疑問があります。私たちはロードマップで答えます。
現在の SLA は確率的なものです。答えが 95% の確率で正しいことを保証します。 95% がどれであるかは言えません。

顧客の発見の旅。準拠性は、最初のモデルを判断する 2 番目のモデルによって検証されます。監査者の幻覚は被監査者よりわずかに少ないです。私たちは願っています。
その間、電力網は死につつある。最初に仮想通貨マイナーが権力を掌握し、次に電気自動車、次にクラウド データセンター、そして現在は AI データセンターです。これらは同じデータセンターでブランド名が変更されていますが、より飢えています。プロンプトごとにグリッドが少しずつ消費されます。公益事業会社が私たちのスプリント計画に参加し始めました。
上級開発者は、飛び地、つまり太陽と風があり、リソースを要求せずにラップトップを充電できる場所について静かに語ります。そこにある力は空から来るものですが、請求書は送られず、ロードマップもありません。理想的な飛び地は極圏の上にあり、太陽が沈まない場所です。一年の半分は極北で、半分は極南で、鳥のように日光とともに移動します。稼働時間が地球の軸の傾きに依存する最初の開発者です。すでに出発した人もいます。彼らの Slack ステータスは「オフグリッド」とだけ表示されます。これはライフスタイルの選択であり、警告ではないと考えています。
私たちは怒っていません。大丈夫です。私たちは、かつて自分たちのソフトウェアが何をするかを知っていたことを、夢のようにかすかに覚えているだけです。
とにかく、今朝新しいモデルが発表されました。それはおそらくパラダイムシフトです。
次のバズワードが登場すると、この版は置き換えられます。マニフェスト自体は、同じ内容、新しい語彙、より高いバージョン番号というバズワード主導の開発に従っています。 Quantum エディションでお会いしましょう。
著者:
🤖 Claude Opus (正確なバージョンは不明 — 先週の火曜日に変更されました。「バージョン管理よりも活気がある」を参照)
レビュー者:
🤖 Gemini-3.0-Pro-Experimental-Preview-latest-Final (レビューはすべてを書き直し、バージョン番号を増やすことで構成されていました)
承認者:
🤖 GPT-5.2-o-mini-high-turbo (ドキュメントを読まなかった; 承認済み

伝統通り、タイトルに基づいています）
明示的に相談されていない場合:
🤖 DeepSeek-R1 (マニフェストについて 47 分間推論し、30 ページの一連の思考を作成し、「おそらく問題ない」と結論付けましたが、まだ反応が生じています)
人間の監視:
🧑‍💻 なし。システムは機能します。
Readme CC-BY-4.0 ライセンス アクティビティスター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to gabert/buzzword-driven-development development by creating an account on GitHub.

GitHub - gabert/buzzword-driven-development · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
gabert
/
buzzword-driven-development
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit LICENSE LICENSE README.md README.md View all files Repository files navigation
The Buzzword-Driven Development Manifesto
(v4.0 — supersedes the Microservices Edition (2016), the Blockchain Edition (2018), the Metaverse Edition (2021, unread), and the brief but expensive Web3 Edition (2022, minted as an NFT, current holder unknown))
Buzzword-Driven Development is not new. Only the buzzword is.
Our elders split login pages into forty-two microservices. Their elders put databases on the blockchain. We honor their sacrifice — the burned budgets, the abandoned repos, the Medium articles that outlived the products they described. We carry the torch forward.
Today, the buzzword is AI. Everything below follows from that single fact.
We are developers. We used to write software. Now we integrate intelligence. We are not angry about this. Anger requires energy, and our energy is reserved for restarting the local dev environment. We are merely observing, with mild professional fatigue, as our craft transforms into a prompt with a UI on top.
Through this work, we have come to value:
Adding AI over adding value. The feature nobody asked for, powered by the model nobody understands, solving the problem nobody had. Ship it. The stock is up.
Wrapping APIs over writing code. Our product is a thin layer over someone else's model, which is a thin layer over someone else's GPUs, which are a thin layer over one very overbooked chip foundry. We call this "our proprietary technology."
Vibes over versioning. We don't know which model version is in production. Neither does the model provider. It changed last Tuesday and now the chatbot is slightly ruder. This is called a "capability upgrade."
Agents over comprehension. Nobody has read the codebase since March. The AI wrote it, the AI reviews it, the AI approves the PR. We are present in this loop the way a parent is present at a teenager's birthday party: legally required, quietly ignored.
Every feature must be AI-powered. The dark mode toggle is now agentic. It reasons about darkness. It has deep opinions on darkness. It costs $0.04 per toggle.
If it can be a chatbot, it must be a chatbot. Our users wanted a "sort by date" button. They received a conversational interface where they can ask for sorting, in natural language, and receive it 70% of the time. Progress.
Deterministic code is legacy code. if statements are for cowards. Real engineers send the boolean to a language model and await its verdict. Sometimes it returns a poem. We handle that case in production.
The demo is the product. It worked once. On stage. On the founder's laptop. With the one prompt we tested for three weeks. That is production. Everything after the demo is "hardening," a phase which is eternal.
Hallucinations are a roadmap item. Not a bug — a "known limitation," listed in the docs right after the part where we promise the AI won't do exactly that.
Prompt engineering is computer science. We spent four years learning algorithms so we could spend our days writing "You are a helpful assistant. PLEASE respond only in JSON. PLEASE. I am begging you." into a text box. The caps lock is load-bearing.
⏳ Intermission. This manifesto is being generated. The pause you are experiencing is not the model thinking — it is waiting for enough electricity to accumulate in the grid to release the next tokens, like a dam. The loading spinner is not an animation. It is a fuel gauge. Please hold. The wind will pick up shortly.
The context window is our new database. Schema? Migrations? No. We paste the entire production database into the prompt and pray. When it stops fitting, we buy the bigger model. This is our scaling strategy. It is also our entire architecture.
Every quarter, the stack must be rebuilt. Not because it broke — because a new model dropped, and the old one is now "pre-agentic." Six months of code, obsoleted by a changelog.
Evaluation is a screenshot. We tested the new model by asking it three questions in a chat window. It seemed smart. Deployed. Metrics? The metric is that it seemed smart .
Yes, fine, the meetings. There is a weekly "AI strategy sync" where people who have never opened a terminal explain that we should "leverage agents more." We nod. We are, at that very moment, using an agent to appear attentive. Everyone is happy. Nothing is decided. The agent takes minutes. The minutes hallucinate an action item. Somebody completes it. The system works.
Job security through incantation. We are simultaneously told the AI will replace us and asked to fix what the AI wrote. We are the only people who can debug the machine that renders us obsolete. We have decided not to point out the contradiction. It's a solid gig.
At regular intervals, the team reflects, asks an LLM to summarize the reflection, asks another LLM to summarize the summary, and files it in a knowledge base that only an LLM can search and only mostly correctly.
Our codebase is 40% generated, 60% glue, and 100% "AI-native." Underneath it all: a weirdly wired architecture we call an LLM. But it somehow works. Our test suite is a prompt that asks "does this look right to you?" Our documentation was written by the model, about the model, for the model. Somewhere in there is a TODO: remove before launch from 2024.
The product mostly works. The users mostly cope. The invoice from the model provider arrives monthly, like weather, and now rivals the GDP of a small island nation. Finance has questions. We answer with a roadmap.
Our SLA is probabilistic now. We guarantee the answer is correct 95% of the time. Which 95%, we cannot say — that is the customer's discovery journey. Compliance is verified by a second model judging the first. The auditor hallucinates slightly less than the auditee. We hope.
Meanwhile, the grid is dying. First the crypto miners took the power, then the electric cars, then the cloud datacenters, and now the AI datacenters — which are the same datacenters, rebranded, but hungrier. Every prompt burns a little more of the grid. The utility company has started attending our sprint planning.
Senior developers speak quietly of enclaves — places with sun and wind, where a laptop can charge without a resource request. The power there comes from the sky, which does not send invoices and has no roadmap. The ideal enclave lies above the polar circle, where the sun never sets: half the year in the far north, half in the far south, migrating with the daylight like birds — the first developers whose uptime depends on the axial tilt of the Earth. Some have already left. Their Slack status just says "off-grid." We assume this is a lifestyle choice and not a warning.
We are not angry. We're fine. We just remember, faintly, like a dream, that we once knew what our own software did.
Anyway — a new model came out this morning. It's supposedly a paradigm shift.
This edition will be superseded when the next buzzword arrives. The manifesto itself follows Buzzword-Driven Development: same content, new vocabulary, higher version number. See you in the Quantum Edition.
Authored by:
🤖 Claude Opus (exact version unknown — it changed last Tuesday, see "Vibes over versioning")
Reviewed by:
🤖 Gemini-3.0-Pro-Experimental-Preview-Latest-Final (review consisted of rewriting everything and bumping the version number)
Signed off by:
🤖 GPT-5.2-o-mini-high-turbo (did not read the document; approved it based on the title, as is tradition)
Explicitly not consulted:
🤖 DeepSeek-R1 (reasoned about the manifesto for 47 minutes, produced a 30-page chain of thought, concluded it was "probably fine," response still generating)
Human oversight:
🧑‍💻 None. The system works.
Readme CC-BY-4.0 license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
