---
source: "https://www.alyph.ai/"
hn_url: "https://news.ycombinator.com/item?id=49215505"
title: "Show HN: Alyph, a manual transmission for LLM context"
article_title: "Alyph | The Canvas for AI Conversations"
author: "rrr_oh_man"
captured_at: "2026-08-07T20:30:08Z"
capture_tool: "hn-digest"
hn_id: 49215505
score: 1
comments: 0
posted_at: "2026-08-07T19:58:51Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Alyph, a manual transmission for LLM context

- HN: [49215505](https://news.ycombinator.com/item?id=49215505)
- Source: [www.alyph.ai](https://www.alyph.ai/)
- Score: 1
- Comments: 0
- Posted: 2026-08-07T19:58:51Z

## Translation

タイトル: Show HN: Alyph、LLM コンテキストの手動送信
記事タイトル: アリフ | AI 会話のためのキャンバス
説明: Alyph は AI 会話のためのキャンバスです。うさぎの穴が始まる前にチャットを分岐させ、答えが得られたら煩わしいやり取りを取り除き、300 を超えるモデルを並べて比較し、すべてのプロジェクトを独自のコンテキストに保ちます。個人の創業者、小規模チーム、独立系コンサルタント向けに構築されています。
HN text: LLM を使い始めて以来、過去に戻って履歴を変更し、別のブランチに進みたいという気持ちが常にありました。私にとってのひらめきの瞬間は、Google AI Studio で AI の応答を実際に編集できるようになったときでした。私は、はい、すべての思考の流れと同じスレッドの変形を分岐させるには、多元宇宙タイプの状況が必要であることに気づきました。同じ思いでGeminiとChatGPTを並行して試してみたいと思いました。 # アイデア Alyph の元のコンセプトは 2024 年に遡ります。初期のコンセプトは、より SF 寄りのものがいくつかありました (コードベースを見ると、Stargate、Hermes などの名前が表示されます)。ホロデッキのように、トピックやコンテンツの「惑星」を訪れる 3D にすることも考えました。しかし、今の実行ははるかに現実的です。 # 今日の Alyph Alyph は、LLM に並行して質問できる 2D キャンバスです。ウサギの穴をたどり、最終的に答え (コードであれテキストであれ) を取得したら、その答えを取り出して、その間の不要なやり取りをすべて削除できます。このアプリには、同じボード上の他のプレイヤーと一緒にプレイできるマルチプレイヤー モードもあります。 (以下のデモボードを参照) # 背景 標準チャット UI の中心的な問題は、私が考えるところ、コンテキスト ポイズニングです。 LLM が実際にどのように動作するかを調べてみると、メモリがゼロであることがわかります。これらは、携帯電話上の非常に派手なオートコンプリートで、全体を（再）読み取ります。

送信するすべてのメッセージでチャットを最初から完了できます。これは、すべての間違った方向転換、侮辱 (おっと)、失敗したデバッグ ループ、および悪いうさぎの穴が、後続のすべてのメッセージとともに送信されることを意味します。あなたの完全な履歴は、未来を予測するために使用されます。単にクリーンな FaqSection.tsx が必要な場合には、これは役に立ちません。より強力なモデルでは、プロンプトを表示する際の言葉の選択は、それに付随するコンテキストほど重要ではないことに気づきました。少なくとも私のユースケースには当てはまりません。たとえば、LLM に非常に高品質のしきい値 (完全に型チェックされた、明確なフロントエンド コンポーネント、主流言語の厳密なスタイル ガイド) を持つコードベースを与えると、本当に高品質なものが非常に簡単に生成される傾向があります。さらに、モデルに多くのコンテキストを与えすぎると、回答の品質が低下する傾向があります。たとえば、Claude は 300,000 ～ 500,000 トークン以下では素晴らしい性能を発揮しますが、それを超えると非常に早く劣化する可能性があります (一方、少なくとも私の経験では、Gemini は大規模なコンテキストで少しパフォーマンスが向上する傾向があります)。それで、あなたは何をしますか？現時点での主流の答えは、アルゴリズムにコンテキストを選択させることです。しかし、私の強い仮定は、モデルはコンテキスト ウィンドウがますます大きくなり続けるということです。コンシューマ モデルでは 32K から始まりましたが、最近では Grok は 2M になっています。 1億個のコンテキストを持つモデルがあれば世界はどう変わるでしょうか? 100B？だからこそ、私は（異端ですが！）ハーネス、コンテキスト圧縮アルゴリズム、エージェントなどが現時点ではストップギャップだと考えています。私の場合、それらは実際にはうまく機能しません（時々重要なことを忘れたり、脱線してスロを生み出すため）、肉の代理人として「受け入れ」または「拒否」を押す私がいると、本当にブラックボックスのように感じます。遠方の管理人です。私は遠いマネージャーにはなりたくない。雑草の中にいたい。それで私が構築したもの

これは、時代錯誤の手動ツールだと言う人もいるかもしれません。いわばマニュアルトランスミッション。ベースラインを 100% コントロールしたいと考えています。モデルに表示させたいものを正確に入力し、実験し、その中にフォルダーをダンプするだけです (ちなみに、Alyph は `.gitignore` ファイルを尊重します。これは素晴らしいことです)。モデルが最終的にはさらに大きなコンテキスト ウィンドウを許可すると仮定すると、そのベースラインを手動で制御できることが重要だと思います。もちろん、Alyph にはまだ荒削りな部分もありますが、それは愛情のこもった作業です。実際に他の人にも効果があるかどうかを確認するために、ぜひ共有したいと思います。それが嫌なら言ってください。好きなら、私にも教えてください。一緒にチームを組んでくれる方も探していますので、ぜひ連絡ください。ライブ: https://alyph.ai 、無料で試せます。 YC のデモボードは次のとおりです: https://my.alyph.app/board/d806842f-6d52-45a4-8bb6-da46e95e5..

[切り捨てられた]

記事本文:
アリフ | AI 会話のキャンバス Alyph 無料でサインアップ 機能 学ぶ 価格の比較 なぜ AI 会話のキャンバス。
大きなアイデアを持ち、多くのことに挑戦する人々のために作られました。
デモを試す 無料で $5 のウェルカム クレジットで開始でき、サブスクリプションは必要ありません
Acme リテイナーを取るべきですか、それともスロットを開いたままにすべきですか?
GPT-5.6 ソル 持ちますが、リテイナーが常にクリープするため、時間には制限があります。
クロード・ソネット 5 あなたがそれを依頼する特定のクライアントの名前を出せる場合にのみ、それを受け入れてください。 ▍
アプリについて質問しましたが、14 メッセージ後、まだローカル マシンをデバッグ中です。
AI はあなたのチャットを覚えていません。それを再読するのです。
通常のチャットでは、送信するすべてのメッセージに、寄り道、エラー、修正を含む会話全体が含まれます。修正は役に立ちますが、残りの部分は無料で使用できるようになり、次の答えが損なわれます。
1 メッセージが増加 · ~0.2k トークン メッセージ 9 9 メッセージが増加 · ~2.1k トークン メッセージ 16 16 メッセージが増加 · ~4.8k トークン キャンバス上の同じプロジェクト。
インストールプロセスには、14 件の戦闘メッセージと 1 件の修正メッセージが必要でした。プルーニングによりこれらの中間メッセージが削除されるため、失敗した試行は最終的な回答によって上書きされ、スレッドは完全にクリーンなままになります。
Pythonでターミナルアプリを作りたいと思っています。どこから始めればよいでしょうか?
GPT-5.6 Sol Use Textual — ウィジェット、レイアウト、イベントを 1 回のインストールで: pip install textual
インストール時のエラー: 致命的なエラー: 'Python.h' ファイルが見つかりません
GPT-5.6 Sol dev ヘッダーがありません。 xcode-select --install を実行して、再試行してください。
新しいもの: ld: -lportaudio のライブラリが見つかりません
GPT-5.6 Sol brew install portaudio pkg-config を実行し、再インストールします。それでクリアです。
進捗 — インポートした瞬間にセグメンテーション違反が発生します。
GPT-5.6 ソル ホイール破損。ソースから再構築: pip install --no-binary :all: textual
修正されました。 OK — 実際のアプリに戻ります。
GPT-5.6 ソルを修正しました。 N

メインループは App、compose()、on_mount です。こちらがきれいな骨格です。
行ったり来たりの剪定 答えが重要なときは、全員に尋ねます。
まったく同じ質問とコンテキストを ChatGPT、Claude、Gemini、および他の 300 以上のモデルに同時に送信すると、回答を並べて読んで判断できます。
クライアントは 12 か月間の維持料金の 20% オフを希望しています。受け取ったほうがいいでしょうか？
GPT-5.6 Sol 双方に90日間の離脱条項を付けて対応します。
クロード・ソネット 5 割引は顧客に何を期待するかを教えるため、10% でカウンターします。
Gemini 3.5 Flash 年間が現金で前払いされている場合にのみ割引を受け入れます。
決断によっては別の人が必要になることもあります。
共同創設者、アシスタント、またはクライアントとキャンバスを共有すると、彼らは AI とまったく同じスペースで執筆および編集できるようになります。
2 オンラインアクメ提案。金曜日までに範囲と価格が必要です。
クロード・ソネット 5 3 層のドラフト。監査を最上位層にバンドルしました。
監査を主導し、価格は二番目に決めます。 ▍ サラが編集中です
Y You S Sarah 忙しい日々のために作られました。
9 時にクライアントの仕事、正午に管理タスク、3 時にコーディング、9 時に犬。すべてが独自のスレッド内に存在するため、コンテキストの切り替えにはまったくコストがかかりません。
09:00 · クライアントワーク Smith & Co: 価格設定ページ。
クロード・ソネット 5 年次 最初に切り替えます。 2 つのバリエーションが用意されています。
GPT-5.6 Sol フォローアップ ドラフト。しっかりと、温かく、そして厳かに1ページ。
クロード ソネット 5 既知の問題。 esbuild を 0.21 に固定します。
hospital-letter.pdf 残りのツールボックス。
1 つのウォレットに 300 を超えるモデル。
OpenAI、Anthropic、Google、DeepSeek、その他数百もの。プロバイダーの価格にインフラストラクチャのマージンを加えたもの。
契約書、スプレッドシート、またはプロジェクト フォルダー全体をドロップインします。モデルのコンテキスト ウィンドウが唯一の制限です。
ヒーロー画像を要求したコピーのすぐ隣にヒーロー画像を生成します。ビデオ機能も開発中です。
すべてのキャンバスも読み取ります

標準チャットとして使用できるため、いつでもビューを切り替えることができます。
ユーザーに代わって考えることを約束するツールは、粗末なものを生み出します。 Alyph はあなたの代わりに仕事をするわけではありません。質問、ファイル、モデル、その他の情報を含む作業内容が保存されます。呼び出しを行うと、出力は独自のものになります。
私は一人でスタートアップを経営していますが、これが私がそれを維持する方法です。
私は朝はクライアントの仕事をし、夜はコーディングをし、合間には獣医師の請求を受けています。私は AI に毎月約 2,000 ドルを費やしていますが、そのすべてが Alyph を介して直接支払われます。
サブスクリプションやスロットリングはありません。実際に使用した分に対して料金を支払い、制限を自分で設定します。
プロバイダー価格にインフラストラクチャ マージンを加えた価格で 300 を超えるモデル。完全に厳しい支出制限のある従量課金制。 5GBのストレージが無料で付属。データが公開モデルのトレーニングに使用されることはありません。 5 ドルの無料クレジットから無料で始められます。キャンバスを開く 購読はありません。キャンセルを忘れることはありません。

## Original Extract

Alyph is the canvas for AI conversations. Branch a chat before a rabbit hole begins, prune the noisy back-and-forth once you get an answer, compare over 300 models side by side, and keep every project in its own context. Built for solo founders, small teams, and independent consultants.

Since I started using LLMs, I've always had this feeling of wanting to go back, change the history, and then go down another branch. The lightbulb moment for me was when Google AI Studio started to let you actually edit the AI's responses. I realized, YES, I need a multiverse type of situation to branch all my thought streams and variants of the same thread. I wanted to try Gemini and ChatGPT in parallel on the same thought. # Idea The original concept for Alyph goes back to 2024. I had some early concepts that were a lot more sci-fi (if you look at the codebase, you’ll see names like Stargate, Hermes, etc.). I even thought about making it 3D, like a holo deck, where you visit "planets" of topics and content. But the execution now is much more, eh, pragmatic. # Alyph today Alyph is a 2D canvas where you can ask LLMs stuff in parallel. When you go down a rabbit hole and finally get an answer (be it code or text), you can pull that answer up and delete all the unnecessary back-and-forth in between. The app also has a multiplayer mode where you can do that together with others on the same board. (see demo board below) # Background The core problem with standard chat UIs, as I see it, is context poisoning. If you look into how LLMs actually work, you realize they have zero memory. They are a very fancy autocomplete on your phone that (re)reads the entire complete chat from scratch with every single message you send. This means: all your wrong turns, the insults (oops), the failed debugging loops, and the bad rabbit holes get sent with every subsequent message. Your complete history is used to predict the future. That doesn’t help when you just want a clean FaqSection.tsx. I noticed that with the more powerful models, the choice of words when prompting isn't as important as the context that goes with it. At least it isn't for my use cases. For example, if you give the LLM a codebase with a very high-quality threshold (fully type-checked, clear front end components, strict style guide in a mainstream language) they tend to produce really high-quality stuff very easily. Furthermore, if you give models too much context, the answers tend to degrade. Claude, for example, is amazing under 300k-500k tokens, but above that, it can degrade very quickly (whereas Gemini tends to perform a bit better with massive context, at least in my experience). So what do you do? The mainstream answer right now is to let an algorithm try to pick and choose the context for you. But: My strong assumption is that models will keep getting bigger and bigger context windows — we started at what, 32K for consumer models, and now Grok is at 2M these days. How would the world change if we had models with 100M context? 100B? That's why I think (heresy!) harnesses, context compacting algorithms, agents etc. are a stop gap right now. They don't really work well in my cases (as they sometimes forget important things and go off on a tangent to produce slop), and it really feels like a black box with me as the meat proxy pressing "accept" or "reject". I'm a distant manager. I don't want to be a distant manager. I want to be in the weeds. So what I built is, what some might say, is an anachronistic manual tool. A manual transmission, so to speak. I want to have 100% control over my baseline. I want to put in exactly what I want the model to see, experiment, and just dump a folder into it (Alyph btw respects your `.gitignore` file, which is neat). With the assumption that models will eventually allow even bigger context windows, having manual control over that baseline is important, I think. Alyph still has some rough edges, of course, but it's a labor of love. I would love to share it to see if it actually works for others. If you hate it, tell me. If you love it, tell me too. I'm also looking for somebody to team up with, so hit me up. Live at: https://alyph.ai , it's free to try. Here's a demo board for YC: https://my.alyph.app/board/d806842f-6d52-45a4-8bb6-da46e95e5..

[truncated]

Alyph | The Canvas for AI Conversations Alyph Sign up free Features Learn Compare Pricing Why The canvas for AI conversations.
Built for people with big ideas and many hats.
Try the demo Free to start with a $5 welcome credit and no subscription
Should I take the Acme retainer or keep the slot open?
GPT-5.6 Sol Take it but cap the hours because retainers always creep.
Claude Sonnet 5 Take it only if you can name the specific client you would drop for it. ▍
You asked about your app but fourteen messages later you are still debugging your local machine.
AI doesn't remember your chat. It rereads it.
In normal chats every message you send carries the whole conversation including the detour, the errors, and the fix. The fix helps but the rest rides along rent free and poisons the next answer.
1 message go up · ~0.2k tokens Message 9 9 messages go up · ~2.1k tokens Message 16 16 messages go up · ~4.8k tokens The same project on a canvas.
The installation process took fourteen messages of fighting and one message of fix. Pruning removes those middle messages so the final answer overwrites the failed attempts and your thread continues completely clean.
I want to build a terminal app in Python. Where do I start?
GPT-5.6 Sol Use Textual — widgets, layout, and events in one install: pip install textual
Error on install: fatal error: 'Python.h' file not found
GPT-5.6 Sol You're missing the dev headers. Run xcode-select --install and try again.
New one: ld: library not found for -lportaudio
GPT-5.6 Sol brew install portaudio pkg-config, then reinstall. That clears it.
Progress — now it segfaults the moment I import it.
GPT-5.6 Sol Broken wheel. Rebuild from source: pip install --no-binary :all: textual
Fixed. OK — back to the actual app.
GPT-5.6 Sol Fixed. Now the main loop: App, compose(), and on_mount. Here is the clean skeleton.
Back-and-forth pruned When the answer matters you ask everyone.
Send the exact same question and context to ChatGPT, Claude, Gemini, and any of the other 300+ models simultaneously so you can read the answers side by side and decide.
A client wants 20% off for a 12-month retainer. Should I take it?
GPT-5.6 Sol Take it with a 90-day out clause for both sides.
Claude Sonnet 5 Counter at 10% because discounts teach clients what to expect.
Gemini 3.5 Flash Only accept the discount if the year is prepaid in cash.
Some decisions need another person.
Share a canvas with your co-founder, your assistant, or your client so they can write and edit in the exact same space as the AI.
2 online Acme proposal. I need the scope and price by Friday.
Claude Sonnet 5 Three tiers drafted. I bundled the audit into the top tier.
Lead with the audit and price it second. ▍ Sarah is editing
Y You S Sarah Built for busy days.
Client work at nine, administrative tasks at noon, coding at three, and the dog at nine. Everything lives in its own thread so context switching costs absolutely nothing.
09:00 · client work Smith & Co: the pricing page.
Claude Sonnet 5 Annual toggle first. Two variants ready.
GPT-5.6 Sol Follow-up draft. Firm, warm, and strictly one page.
Claude Sonnet 5 Known issue. Pin esbuild to 0.21.
hospital-letter.pdf The rest of the toolbox.
Over 300 models in one wallet.
OpenAI, Anthropic, Google, DeepSeek, and hundreds more. Provider prices plus an infrastructure margin.
Drop in a contract, a spreadsheet, or an entire project folder. The model's context window is your only limit.
Generate your hero image right next to the copy that asked for it. Video capabilities are on the way.
Every canvas also reads as a standard chat so you can switch views whenever you like.
Tools that promise to think for you produce slop. Alyph does not do your work for you. It holds your work including the questions, the files, the models, and the mess. You make the calls so the output stays uniquely yours.
I run a one-person startup and this is how I keep it together.
I handle client work in the morning, coding at night, and a vet claim in between. I spend about $2,000 a month on AI and all of it goes directly through Alyph.
No subscription and no throttling. You pay for what you actually use and you set the limits yourself.
Over 300 models at provider prices plus an infrastructure margin. Pay-as-you-go with completely hard spending limits. 5 GB of storage included for free. Your data is never used to train public models. Free to start with $5 of complimentary credit. Open the canvas No subscription. You will never have to remember to cancel.
