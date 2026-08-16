---
source: "https://conw.ai"
hn_url: "https://news.ycombinator.com/item?id=49323625"
title: "Show HN: Conw.ai – Independent local AI platform and developer API"
article_title: "conw.ai — The AI that learns with its users."
author: "thomasconway01"
captured_at: "2026-08-16T21:11:21Z"
capture_tool: "hn-digest"
hn_id: 49323625
score: 1
comments: 0
posted_at: "2026-08-16T21:00:47Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Conw.ai – Independent local AI platform and developer API

- HN: [49323625](https://news.ycombinator.com/item?id=49323625)
- Source: [conw.ai](https://conw.ai)
- Score: 1
- Comments: 0
- Posted: 2026-08-16T21:00:47Z

## Translation

タイトル: 表示 HN: Conw.ai – 独立したローカル AI プラットフォームと開発者 API
記事のタイトル: conw.ai — ユーザーとともに学習する AI。
説明: Conw は、独立してサービスを提供する自己学習 AI です。 Conway-Omega は確認された指導をすぐに記憶し、検証後にのみ重量の更新を促進します。

記事本文:
conw.ai — ユーザーとともに学習する AI。コンウ。 ai 学習方法 原則 計画 API FAQ 便利なチャットで学習します。何が引っかかったのか見てみましょう。
Conw は独立して提供されます。確認済みの指導はすぐに機能し、安全で役立つチャットはレビュー済みの学習候補になる可能性があります。保存された内容を確認して、どの返信を保持する価値があるかを知ることができます。
カード不要 30 秒のセットアップ 無料で使用可能 検証済みの学習 独立して提供されるローカル MLX 推論 便利なチャットから学習 リビング ルーム iMac 無料で使用 検証済みの学習 独立して提供されるローカル MLX 推論 有用なチャットから学ぶ リビング ルーム iMac 無料で使用 学習方法
あなたが教えます。フィルタリングします。それは何が引っかかっていたかを証明します。
1 つの悪いチャットでシステムを書き換えることなく、有益なシグナルから改善する保護された学習ループ。
返信を修正したり、評価したり、新しい単語を説明したりできます。確認された意味はすぐにあなたの個人的な記憶に入り、Conw は同じ会話でそれらを使用できます。
機密性が高く、安全でなく、フォーマットが不適切で、評価が低い返信は除外されます。レビュー済みの学習キューに入れることができるのは、短くて有用な例だけです。
バックグラウンド更新は、十分なクリーンなサンプルとアイドル状態のサーバーを待機します。ライブのものを置き換えるには、候補者がチェックに合格する必要があります。
iMacが静かになったらレビューします。
適格なサンプルは、保護されたキューで待機します。トレーニングは、十分な清潔な材料があり、提供がアイドル状態になった後にのみ開始されます。結果は候補であり、自動ライブ更新ではありません。
安全で有用な短い返信だけが生き残ります。低評価や汚染された返信は除外されます。
バックグラウンド作業は iMac がアイドルになるまで待機するため、ライブチャットが優先されます。
ライブで何かを変更する前に、候補はロードしてチェックに合格する必要があります。
まずはライブチャット、次に学習
私たちが真顔で言える3つのこと。
ほとんどの AI 製品は学習ループを隠しています。コ

nw は、レッスンが成功したか失敗したかを示し、重みの更新が拒否された場合でも有用なメモリを保持します。
使う人から学びます。
確認されたティーチングはすぐにメモリに書き込まれます。辞書で調べた意味は共有語彙になる可能性があります。未検証または安全でないレッスンは元のチャット内に残ります。今後体重を変更する場合は、最初に安全性チェックとクイズチェックに合格する必要があります。
確認された意味は記憶からすぐに機能します
辞書で検証された語彙はすべてのユーザーに役立ちます
安全でない、でっち上げられた、または毒を含んだレッスンは禁止されます
Conw は記念碑ではなくワークショップです。記憶、安全性、検証、ユーザーの所有権はそのままに、サービス提供モデルは変更できます。
私たちのサービングスタック。当社のハードウェア。交換可能なチェックポイント。
Conway-Retrain 12B は現在ライブであり、Gemma 4 ベース上で再トレーニングされ、外部の応答 API ではなく、独自のマシン上の MLX を通じて提供されます。以前の 188M Conway-Omega チェックポイントは製品内では非推奨になりましたが、Hugging Face でオープンソースとして公開されたままです。チェックポイントはインターフェイスの背後に留まるため、メモリと安全ルールは将来のモデル変更後も存続します。
推論は Conw によって提供され、別のモデル API には転送されません
チェックポイント固有のしきい値は、製品コードではなくキャリブレーション内に存在します。
学習ループは将来のチェックポイントの置き換えにも耐えられます。
最も環境に優しいトークンは、生成する必要がないものです。効率は、有用な答え、制限された出力、そして待っている人を決して飢えさせることのないトレーニングから始まります。
効率は発明するものではなく、測定できるものです。
Conway-Retrain 12B は現在、1 台の 16GB iMac からサービスを提供しています。短答トークンの上限により無駄が回避され、反復ガードにより出力の暴走が阻止され、チャット トラフィックがマシンを必要とするたびにバックグラウンド学習が生成されます。炭素数を適切に測定できるようになるまでは、炭素数を公表しません。
GPU の代わりに 16GB サービング マシン 1 台

クラスター
応答バジェットにより不必要なトークン生成が阻止される
推論が忙しい場合、バックグラウンド学習が延期される
独自のコードから Conway-Retrain を構築します。
api.conw.ai にある Conway-Retrain 12B および Conway-Omega 188M への従量課金制アクセス。チャット製品の背後にある同じモデルで、OpenAI SDK からアクセスできます。
OpenAI 互換エンドポイントのドロップイン - Base_url を交換して実行します。
いずれのモデルでも、入力および出力、1,000,000 トークンあたり £1.50。
名前付きキー、使用状況グラフ、最低 £5.00 のチャージ。
「ほとんどの AI マーケティングはフィクションです。私たちのものにはそんな余裕はありません。今日嘘をついたとしても、明日には気づくでしょう。」
公共の場で向上するモデルには、隠れる場所はありません。
何か見逃したことはありますか？ [email protected] — 人間が返信します。
Conw は本当に私のチャットから学習しますか?はい、ただし、すべてを盲目的にトレーニングするわけではありません。確認された教えはすぐに記憶に残ります。辞書で検証された意味は共有語彙になる可能性がありますが、センシティブなレッスン、安全でないレッスン、でっち上げられたレッスン、不適切な形式のレッスン、低評価のレッスンは除外されます。今後の体重更新についても、昇格前にチェックが必要です。
正直に言って、それはどのくらい賢いのでしょうか？これはコンパクトでローカルに提供されるモデルであり、フロンティア システムは広範さとトリビアの点でこれに勝ります。そのエッジはループです。確認された教育は個人の記憶からすぐに機能し、その後、検証された更新により、人々が実際に使用するモデルを改善できます。
本当に緑色なのでしょうか？ AI のエネルギーに関する主張のほとんどは証明することが難しいため、私たちはこの主張を控えめにしています。現在のサービスは 1 台の 16 GB iMac で実行され、不要な出力を制限し、推論にマシンが必要な場合にはバックグラウンド学習を一時停止します。適切な測定なしに炭素数を公表することはありません。
Conw は他の人のモデルのラッパーですか?いいえ。Conw は質問を外部の回答 API に転送しません。 Conway-Retrain 12B は、上に再トレーニングされた現在のモデルです。

Gemma 4 ベースを使用し、独自のマシンで MLX を通じてサービスを提供しました。以前の 188M Conway-Omega チェックポイントは製品内では非推奨になりましたが、オープンソースで Hugging Face 上に存続しています。
どうして無料なのでしょうか？コンパクトなローカル サービス設定によりコストが低く抑えられます。無料プランは、毎週のカジュアルな使用をカバーします。 Pro と Max は週ごとの手当を増額し、サービスと保護された学習ループの支払いを支援します。どの層でも同じ答えが得られます。広告もデータの販売もありません。
私のデータはどうなりますか?会話は保存されるので、適格な部分が保護された学習プロセスに参加できるようになります。私たちは広告を掲載したり、データを販売したりしません。アカウントとチャットの削除をご希望の場合は、メールでご連絡ください。削除させていただきます。
今日はそれに会いましょう。
役立つことを教えてください。
Conwは無料で使用できます。確認された指導はすぐに機能し、安全で有用な返信はレビューされた学習候補となる可能性があります。
カードはありません。スパムはありません。無料でご利用いただけます。コンウ。 ai 保護された学習ループを備えた小さな AI。独立して提供されます。
© 2026 conw.ai — まだ学生の独立系モデル。
役立つ教えが今保存されています。重みの更新はチェックを待ちます。

## Original Extract

Conw is an independently served, self-learning AI. Conway-Omega remembers confirmed teaching immediately and only promotes weight updates after verification.

conw.ai — The AI that learns with its users. conw . ai How it learns Principles Plans API FAQ Useful chats teach it. See what stuck.
Conw is independently served. Confirmed teaching works immediately, while safe and useful chats can become reviewed learning candidates — you can see what was saved and tell us which replies are worth keeping.
No card required 30-second setup Free to use Verified learning Independently served Local MLX inference Learns from useful chats Living-room iMac Free to use Verified learning Independently served Local MLX inference Learns from useful chats Living-room iMac Free to use How it learns
You teach. It filters. It proves what stuck.
A guarded learning loop that improves from useful signals without letting one bad chat rewrite the system.
Correct a reply, rate it, or explain a new word. Confirmed meanings enter your private memory immediately so Conw can use them in the same conversation.
Sensitive, unsafe, badly formatted, and down-rated replies stay out. Only short, useful examples can enter the reviewed learning queue.
Background updates wait for enough clean examples and an idle server. A candidate must pass checks before it can replace anything live.
When the iMac is quiet, it reviews.
Eligible examples wait in a guarded queue. Training starts only after there is enough clean material and serving is idle. The result is a candidate, not an automatic live update.
Only safe, useful, short replies survive. Down-rated and contaminated replies stay out.
Background work waits for the iMac to be idle so live chat keeps priority.
A candidate must load and pass its checks before anything live can change.
live chat first · learning second
Three things we can say with a straight face.
Most AI products hide the learning loop. Conw shows whether a lesson passed or failed, and keeps the useful memory even when the weight update is rejected.
It learns from the people who use it.
Confirmed teaching is written to memory immediately. Dictionary-checked meanings can become shared vocabulary; unverified or unsafe lessons stay inside the original chat. Any future weight change still has to pass safety and quiz checks first.
Confirmed meanings work immediately from memory
Dictionary-verified vocabulary can help every user
Unsafe, invented, or poisoned lessons stay out
Conw is a workshop, not a monument: the serving model can change while memory, safety, verification, and user ownership stay intact.
Our serving stack. Our hardware. A replaceable checkpoint.
Conway-Retrain 12B is now live, retrained on top of a Gemma 4 base and served through MLX on our own machine, not an external answer API. The earlier 188M Conway-Omega checkpoint is deprecated in-product, but stays published, open-source, on Hugging Face . The checkpoint stays behind an interface, so memory and safety rules survive future model changes.
Inference is served by Conw, not forwarded to another model API
Checkpoint-specific thresholds live in calibration, not product code
The learning loop survives future checkpoint replacements
The greenest token is the one we never need to generate. Efficiency starts with useful answers, bounded output, and training that never starves the person waiting.
Efficiency we can measure, not invent.
Conway-Retrain 12B currently serves from one 16GB iMac. Short-answer token ceilings avoid waste, repetition guards stop runaway output, and background learning yields whenever chat traffic needs the machine. We will not publish a carbon number until we can meter it properly.
One 16GB serving machine instead of a GPU cluster
Response budgets stop needless token generation
Background learning defers when inference is busy
Build on Conway-Retrain from your own code.
Pay-as-you-go access to Conway-Retrain 12B and Conway-Omega 188M, live at api.conw.ai — the same model behind the chat product, reachable from any OpenAI SDK.
Drop-in OpenAI-compatible endpoints — swap base_url and go.
£1.50 per 1,000,000 tokens, input and output, either model.
Named keys, usage graphs, and a £5.00 minimum top-up.
“Most AI marketing is fiction. Ours can't afford to be — you'd notice tomorrow if we lied today. ”
A model that improves in public has nowhere to hide.
Anything we missed? [email protected] — a human writes back.
Does Conw really learn from my chats? Yes, but not by blindly training on everything. Confirmed teaching enters memory immediately. Dictionary-verified meanings can become shared vocabulary, while sensitive, unsafe, invented, badly formatted, and down-rated lessons stay out. Any future weight update still needs checks before promotion.
How smart is it, honestly? It is a compact, locally served model, and frontier systems will still beat it on breadth and trivia. Its edge is the loop: confirmed teaching works immediately from private memory, then verified updates can improve the model where people actually use it.
Is it actually green? We keep this claim modest because most AI energy claims are difficult to prove. The current service runs on one 16GB iMac, caps needless output, and pauses background learning when inference needs the machine. We do not publish a carbon number without proper metering.
Is Conw a wrapper around someone else’s model? No. Conw does not forward your question to an external answer API. Conway-Retrain 12B is our current model, retrained on top of a Gemma 4 base and served ourselves through MLX on our own machine. The earlier 188M Conway-Omega checkpoint is deprecated in-product, but lives on, open-source, on Hugging Face .
How is it free? The compact local serving setup keeps costs low. The Free plan covers casual weekly use; Pro and Max raise your weekly allowance and help pay for serving and the guarded learning loop. Every tier gets the same answers. No ads, no selling your data.
What happens to my data? Conversations are stored so eligible parts can enter the guarded learning process. We don’t run ads and we don’t sell data. If you want your account and chats removed, email us and we’ll delete them.
Meet it today.
Teach it something useful.
Conw is free to use. Confirmed teaching works immediately, and safe, useful replies can become reviewed learning candidates.
No card. No spam. Free to use. conw . ai The little AI with a guarded learning loop. Independently served.
© 2026 conw.ai — an independent model, still in school.
Useful teaching is saved now. Weight updates wait for checks.
