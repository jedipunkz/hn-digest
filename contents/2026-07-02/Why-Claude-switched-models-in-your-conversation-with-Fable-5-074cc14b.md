---
source: "https://support.claude.com/en/articles/15363606-why-claude-switched-models-in-your-conversation-with-fable-5"
hn_url: "https://news.ycombinator.com/item?id=48759255"
title: "Why Claude switched models in your conversation with Fable 5"
article_title: "Why Claude switched models in your conversation with Fable 5 | Claude Help Center"
author: "adithyaharish"
captured_at: "2026-07-02T11:27:36Z"
capture_tool: "hn-digest"
hn_id: 48759255
score: 1
comments: 0
posted_at: "2026-07-02T10:35:08Z"
tags:
  - hacker-news
  - translated
---

# Why Claude switched models in your conversation with Fable 5

- HN: [48759255](https://news.ycombinator.com/item?id=48759255)
- Source: [support.claude.com](https://support.claude.com/en/articles/15363606-why-claude-switched-models-in-your-conversation-with-fable-5)
- Score: 1
- Comments: 0
- Posted: 2026-07-02T10:35:08Z

## Translation

タイトル: クロードがフェイブル 5 との会話でモデルを切り替えた理由
記事のタイトル: クロードがフェイブル 5 との会話でモデルを切り替えた理由 |クロード ヘルプセンター

記事本文:
クロードがフェイブル5との会話でモデルを切り替えた理由 |クロード ヘルプ センター メイン コンテンツにスキップ API ドキュメント リリース ノート サポートを受ける方法 English Français Deutsch Bahasa India Italiano 日本語 한국어 Português Pусский 简体中文 Español 繁體中文 English API ドキュメント リリース ノート サポートを受ける方法 English Français Deutsch Bahasa India Italiano 日本語 한국어 Português Pусский 简体中文 Español 繁體中文 English 記事の検索... すべてのコレクション
クロードがフェイブル5との会話でモデルを切り替えた理由
クロードがフェイブル5との会話でモデルを切り替えた理由
この記事では、リクエストがブロックされる理由、会話が別のクロード モデルに切り替わると何が起こるか、自動切り替えを管理する方法について説明します。
Claude Fable 5 の機能は、これまでに一般提供したすべてのモデルの機能をはるかに上回っています。これは、テストされた AI 機能のほぼすべてのベンチマークにおいて最先端であり、ソフトウェア エンジニアリング、ナレッジ ワーク、ビジョン、その他多くの分野で優れたパフォーマンスを示します。
これだけの性能を持つモデルをリリースするにはリスクが伴います。強力な保護策がなければ、サイバーセキュリティや生物学などの分野におけるクロード・フェイブル 5 の高度な機能がユーザーによって悪用され、壊滅的な被害をもたらす可能性のある大規模なサイバー攻撃や生物兵器が作成される可能性があります。これらの機能が、これまで Mythos クラスのモデル (Mythos Preview など) を厳選され精査された少数のパートナーにのみリリースしてきた理由です。
これらのリスクを認識し、一般ユーザーが Fable 5 の機能の大部分にアクセスできるようにするために、サービス利用規約と利用規約に沿って、特定の領域での応答をブロックする保護機能を備えたモデルを立ち上げました。私たちも反復してきました

Claude Fable 5 の最初の発売以来、安全対策を講じています。
Fable 5 のこれらのセーフガードによってブロックされたほとんどのユーザー クエリは、代わりに、次に高性能なモデルである Claude Opus 4.8 から応答を受け取る可能性があります (つまり、「フォールバック」)。私たちは、現在よりも誤検知を減らし、リスクの標的化に直接関係するモデルの使用を正確にブロックするために、これらの保護手段をより注意深く作成することに取り組んでいます。
Claude Fable 5 は、すべてのユーザー要求に対して自動安全性チェックを実行します。これらのチェックは、ユーザーが次の 4 つの領域でリクエストを送信したときに、Fable 5 から非 Mythos モデル (Opus 4.8 など) に視覚的にフォールバックすることを目的としています。
エクスプロイト、マルウェア、攻撃ツールの構築など、攻撃的なサイバーセキュリティ技術。 Claude Fable 5 は日常的なサイバーセキュリティタスクを支援しますが、ユーザーは高いフォールバック率を期待する必要があります。安全対策は、Mythos レベルの機能へのアクセスをブロックするように設計されています。
研究室の方法や分子機構など、生物学、化学、生命科学に関するクエリの大部分。短期的には、これは、良性の生物学の研究や、バイオテクノロジーのビジネス文書、医療画像と診断、医療の臨床および診断に関する質問、生物学の基本的な教育コンテンツなどの関連トピックを支援するモデルの能力に影響を与える可能性があります。
モデルの要約された考え方を抽出する試みを含む、Fable 5 に対する蒸留攻撃。
分散トレーニング インフラストラクチャ、ML アクセラレータ設計、特定の非標準チップ用のカーネル開発など、限られた最先端の LLM 開発タスク。
これらのブロックする安全策は意図的に広範囲に設定されており、ユーザー エクスペリエンスへの影響を軽減するために安全策を継続的に改善するよう取り組んでいます。リクエストがブロックされると、非 Mythos モデル (現在は Opus 4.8) にフォールバックする可能性があります。
チェックではモデルのすべても確認されます

最新のメッセージだけでなく、メモリ、コネクタからのコンテンツ、Web 検索結果、ファイルなども読み取るため、入力しなかったコンテンツによってブロックがトリガーされる可能性があります。
既定では、Claude、Claude Cowork、Claude Code、Claude Design、および Claude for Microsoft 365 で自動モデル切り替えが有効になっています。 モデルを自動的に切り替えると、Claude は同じ会話内の Claude Opus 4.8 でブロックされた Claude Fable 5 リクエストを再実行します。モデルが切り替わったことを説明する通知が表示され、応答には応答したモデルのラベルが付けられます。 Opus は、独自の強力な保護機能を備えた高機能なモデルであり、Fable 5 でブロックされている正当なリクエストのほとんどに対して、Opus は役立つ答えをくれるはずです。
切り替え後、モデル ピッカーは会話の残りの間 Opus に残ります。モデル ピッカーからいつでも Claude Fable 5 に戻すことができます。
注: 自動モデル切り替えが発生した後に Claude Fable 5 に戻す場合は、元のリクエストがまだ会話の一部であるため、同じ Fable 5 の安全対策によって会話が再びブロックされる可能性があることに注意してください。再試行する前に前のメッセージを編集すると、多くの場合役に立ちます。
Opus でもリクエストがブロックされた場合
オーパスには独自の安全システムがあります。 Opus でもリクエストがブロックされた場合は、メッセージを編集して再試行できます。特にサイバーに関しては、ユースケースに正当な防御目的があり、これらの安全対策の影響を受けている場合は、Opus のサイバー検証プログラム (CVP) を申請できます。リアルタイム サイバー セーフガードとサイバー検証プログラムの詳細については、こちらをご覧ください。
自動モデル切り替えを管理する
初めて Claude Fable 5 を選択すると、自動切り替えがデフォルトで有効になります。デフォルトでオンのままですが、いつでもオフにできます。
[設定] > [機能] (またはクロード コードの場合は [構成] > [モデルと出力]) に移動します。

メッセージにフラグが立てられたときにモデルを切り替えます。
自動モデル切り替えをオフにすると、ブロックされたリクエストはモデルを切り替える代わりに会話を一時停止します。その後、次のことが可能になります。
Edit your message and retry on Claude Fable 5
Send the same message to Opus manually
ブロックされたリクエストは、ブロックが発生したタイミングに応じて異なる請求が行われます。
入力時のブロック: Claude Fable 5 が出力を生成する前にリクエストがブロックされた場合、会話は直ちに Opus に切り替わります。 Opus の料金でのみ請求され、Opus の応答は使用制限または使用量にカウントされます。
ミッドストリームのブロック: リクエストがミッドストリームでブロックされた場合、ブロック前にストリーミングされた入力とトークンは Claude Fable 5 のレートで課金されます。残りの応答には Opus レートが適用されます。
ブロックされたリクエストがセキュリティや生物学のトピックに無関係であると思われる場合、またはこれらの分野での正当な作業がブロックされ続けている場合は、お知らせください。 Use "Send feedback" to report it.誤ってブロックされたリクエストの報告は、これらの安全策を絞り込み、改善するのに役立ちます。
今後、サイバー防衛と生物学の二重用途への割り当てを開放する方法を検討する予定です。当社の安全システムが成熟するにつれ、悪用に対する強力な保護を維持しながら、正当な生物学と防御的なサイバーセキュリティ作業をサポートすることを目指しています。
資格や申請方法など、プログラムの詳細については、決まり次第お知らせいたします。このヘルプセンターで最新情報を確認するか、ここで通知にサインアップしてください。
Where automatic model switching applies
自動モデル切り替えは、Claude Fable 5 を使用できる場所であればどこでも同じように機能します。
重要: Claude API を使用している場合、モデルの切り替えの動作は異なります。自動切り替えは自動ではないため、API 顧客は API で切り替えをオプトインして構成する必要があります。見る

詳細については開発者ドキュメントを参照してください。
クロード寓話 5 の詳細については、ブログをお読みください: クロード寓話 5 とクロード神話 5 。
Claude API のコンテキスト ウィンドウの大きさはどれくらいですか?クロード コードのモデル構成 対象モデル Mythos クラス モデルのデータ保持慣行 これはあなたの質問の答えになりましたか? 😞 😐 😃 目次 製品

## Original Extract

Why Claude switched models in your conversation with Fable 5 | Claude Help Center Skip to main content API Docs Release Notes How to Get Support English Français Deutsch Bahasa Indonesia Italiano 日本語 한국어 Português Pусский 简体中文 Español 繁體中文 English API Docs Release Notes How to Get Support English Français Deutsch Bahasa Indonesia Italiano 日本語 한국어 Português Pусский 简体中文 Español 繁體中文 English Search for articles... All Collections
Why Claude switched models in your conversation with Fable 5
Why Claude switched models in your conversation with Fable 5
This article explains why a request might be blocked, what happens when your conversation switches to a different Claude model, and how to manage automatic switching.
Claude Fable 5's capabilities far exceed those of every model we've previously made generally available. It is state-of-the-art on nearly all tested benchmarks of AI capability, showing exceptional performance in software engineering, knowledge work, vision, and many other areas.
Releasing a model this capable comes with risks. Without strong safeguards, Claude Fable 5's advanced capabilities in areas like cybersecurity and biology could be misused by users to create large-scale cyberattacks or bioweapons that could result in catastrophic damage. These capabilities are the reason we’ve previously only released Mythos-class models (like Mythos Preview) to a small number of selected and vetted partners.
Recognizing these risks, to allow general users to access the vast majority of Fable 5's capabilities, we've launched the model with safeguards that block its responses in some specific areas in line with our Terms of Service and Acceptable Use Policy . We’ve also been iterating on safeguards since our first launch of Claude Fable 5.
Most user queries blocked by these safeguards on Fable 5 may instead receive a response from our next-most-capable model, Claude Opus 4.8 (i.e., "fallback"). We're working on making these safeguards more discerning to precisely block uses of the model that directly relate to targeting risks, with fewer false positives than there are today.
Claude Fable 5 runs automated safety checks on every user request. These checks are intended to visibly fallback from Fable 5 to a non-Mythos model (e.g., Opus 4.8) when users submit requests in four areas:
Offensive cybersecurity techniques, such as building exploits, malware, or attack tooling. Claude Fable 5 can assist with routine cybersecurity tasks, but users should expect high fallback rates. The safeguards are designed to block access to Mythos-level capabilities.
Majority of biology, chemistry, and life sciences queries, such as lab methods or molecular mechanisms. In the near-term, this may impact the model’s ability to help with benign biology research and related topics, such as biotech business documentation, medical imaging and diagnostics, clinical and diagnostic healthcare questions, or basic educational content in biology.
Distillation attacks on Fable 5, including attempts to extract the model’s summarized thinking .
A narrow set of frontier LLM development tasks, such as distributed training infrastructure, ML accelerator design, and kernel development for certain non-standard chips.
These blocking safeguards are intentionally broad, and we work to continuously improve the safeguards to reduce their user-experience impact. When requests are blocked, they may fallback to a non-Mythos model, currently Opus 4.8.
The checks also review everything the model reads, not just your latest message—including memory, content from connectors, web search results, and files, so a block can be triggered by content you didn't type.
By default, automatic model switching is active in Claude, Claude Cowork, Claude Code, Claude Design, and Claude for Microsoft 365. When automatically switching models, Claude re-runs your blocked Claude Fable 5 request on Claude Opus 4.8 in the same conversation. You’ll see a notice explaining that the model switched, and the response will be labeled with the model that answered. Opus is a highly capable model with strong safeguards of its own, and for most otherwise legitimate requests blocked on Fable 5, Opus should give you a helpful answer.
After the switch, the model picker stays on Opus for the rest of the conversation. You can switch back to Claude Fable 5 anytime from the model picker.
Note: If you switch back to Claude Fable 5 after an automatic model switch occurs, note that the same Fable 5 safeguards may block the conversation again because the original request is still part of it. Editing your previous message before retrying often helps.
If the request is also blocked on Opus
Opus has its own safety systems. If your request is also blocked on Opus, you can edit your message and retry. For cyber specifically, if your use case has a legitimate defensive purpose and is being affected by these safeguards, you can apply for the Cyber Verification Program (CVP) for Opus. Learn more about real-time cyber safeguards and the Cyber Verification Program .
Manage automatic model switching
Automatic switching is enabled by default the first time you select Claude Fable 5. It stays on by default, and you can turn it off anytime:
Go to Settings > Capabilities (or Config > MODEL & OUTPUT in Claude Code).
Toggle Switch models when a message is flagged off.
With automatic model switching off, a blocked request pauses the conversation instead of switching models. You can then:
Edit your message and retry on Claude Fable 5
Send the same message to Opus manually
Blocked requests are billed differently depending on when the block happens:
Blocked on input: If a request is blocked before Claude Fable 5 produces any output, the conversation switches to Opus immediately. You're charged only at Opus rates, and the Opus response counts toward your usage limit or consumption.
Blocked midstream: If a request is blocked midstream, the input and the tokens streamed before the block are charged at Claude Fable 5 rates. The rest of the response is charged at Opus rates.
If your blocked request seems unrelated to security or biology topics, or if your legitimate work in these areas keeps getting blocked, let us know. Use "Send feedback" to report it. Reports of incorrectly blocked requests help us narrow and improve these safeguards.
Moving forward, we plan to consider ways to open up allocations for dual-use cyberdefense and biology research. As our safety systems mature, we aim to support legitimate biology and defensive cybersecurity work while keeping strong protections against misuse in place.
We'll share more details about the program, including eligibility and how to apply, as they become available. Watch this Help Center for updates, or sign up for notifications here .
Where automatic model switching applies
Automatic model switching works the same way everywhere you can use Claude Fable 5:
Important: If you're using the Claude API, model switching works differently. Automatic switching isn't automatic, and API customers must opt into and configure the switching in the API. See the developer documentation for details.
Read our blog to learn more about Claude Fable 5: Claude Fable 5 and Claude Mythos 5 .
How large is the Claude API’s context window? Claude Code model configuration Covered Models Data retention practices for Mythos-class models Did this answer your question? 😞 😐 😃 Table of contents Product
