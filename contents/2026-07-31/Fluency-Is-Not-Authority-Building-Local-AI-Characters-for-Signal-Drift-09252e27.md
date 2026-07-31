---
source: "https://engineering.signaldrift.net"
hn_url: "https://news.ycombinator.com/item?id=49119156"
title: "Fluency Is Not Authority: Building Local AI Characters for Signal Drift"
article_title: "Fluency Is Not Authority — Building Local AI Characters"
author: "alcoveresearch"
captured_at: "2026-07-31T05:26:17Z"
capture_tool: "hn-digest"
hn_id: 49119156
score: 1
comments: 0
posted_at: "2026-07-31T04:54:39Z"
tags:
  - hacker-news
  - translated
---

# Fluency Is Not Authority: Building Local AI Characters for Signal Drift

- HN: [49119156](https://news.ycombinator.com/item?id=49119156)
- Source: [engineering.signaldrift.net](https://engineering.signaldrift.net)
- Score: 1
- Comments: 0
- Posted: 2026-07-31T04:54:39Z

## Translation

タイトル: 流暢さは権威ではありません: 信号ドリフト用のローカル AI キャラクターの構築
記事のタイトル: 流暢さは権威ではありません — ローカル AI キャラクターの構築
説明: Signal Drift がバンドルされたローカル言語モデルを作成されたゲーム状態内に保持する方法。

記事本文:
流暢さは権威ではない — ローカル AI キャラクターの構築
シグナルドリフト / エンジニアリングノート 001 ローカル推論・ナラティブシステム・ポストモーテム流暢性は
権威ではない
私は、940 MiB、15 億パラメータの言語モデルをクロスプラットフォームのインディー ゲーム内に出荷しました。対話を生成するのは簡単な部分でした。難しかったのは、流暢なキャラクターが、確実であるように聞こえるだけでゲームを書き換えることが決してできないようにすることでした。
Signal Ghost のすべての推論はプレーヤーのデバイス上で実行されます。正規のゲーム状態は作成されたままです。
物語ゲームを打破する最も簡単な方法の 1 つは、説得力のある文章を真実にしてしまうことです。
Signal Drift を構築している間、私が何度も戻ってきたシーンはシンプルでした。プレイヤーは架空のホストに侵入し、クライアントの説明と矛盾するログを見つけ、IRC でそのクライアントと対峙します。従来のダイアログ ツリーは、私が予想していた非難を処理できますが、プレイヤーが入力する可能性のあるすべての解釈を処理できるわけではありません。言語モデルはほぼすべてのことに応答できますが、制約のないモデルでは、自白、新たな陰謀者、または決して起こらなかった結果を発明する可能性があります。
プロットを書くのにモデルは必要ありませんでした。ゲームがすでに真実であると知っているプロットにキャラクターを住まわせるためにこれが必要でした。
コアルール モデルがキャラクターを演じます。作者のゲームが現実を決定します。
フリーテキストの対話が端末ゲームに適している理由
Signal Drift は、架空の端末、ログ、メール、IRC ルーム、証拠、および読み取り可能な Linux スタイルのコマンドを中心に構築された CRT ノワールのハッキング スリラーです。プレイヤーはホストをスキャンし、資格情報を回復し、SSH セッションを移動し、イベントを再構築し、見つけたものを公開、販売、変更、または消去するかどうかを決定します。
タイピングはすでにゲームの主要な動詞です。登場人物に「なぜこのタイムスタンプが監査ログと一致しないのですか?」と尋ねます。自然な感じがします

端末の横にあるダイアログホイールを操作することはできません。
目標は無限の対話ではありませんでした。それは、嘘に反論したり、発見されたファイルが何を意味するのかを尋ねたり、ライバルを挑発したり、裏切りを再訪したり、契約が終わった後に誰かと話したりするなど、作成されたビートの間の人間的な空間をカバーするためでした。
推論は外部 AI サービスなしでローカルで実行されます。
同じ機能が Windows、macOS、Linux にも搭載されています。
JavaScript は、状態と安全性に関して引き続き権威を持っています。
キャラクターは無限の歴史ではなく、設計された記憶を受け取ります。
ローカル推論により、API の依存関係と定期的な推論の請求がなくなりましたが、機能が無料になるわけではありません。これにより、コストがダウンロード サイズ、メモリ、起動時間、クロスプラットフォーム パッケージング、ハードウェアの差異、サポートに移されました。
Signal Ghost は、約 15 億 4,000 万のパラメーターを持つ Apache-2.0 モデルである Qwen2.5-1.5B-Instruct を使用します。 Signal Ghost v19 ダイアログ LoRA をトレーニングし、それをベース モデルに融合し、結果を GGUF に変換し、Q4_K_M に量子化しました。
出荷ファイルは 986,048,512 バイト、つまり約 940 MiB です。これは、現在のビルドで b9835 に固定されている llama.cpp を通じて実行されます。
より小規模なスモークテストモデルは、それが「単なるコンピュータープログラム」であることを発表したり、一般的なセキュリティの散文に陥ったり、劇的な瞬間が過ぎた後も話し続けたりするでしょう。コンパクトでも使えるわけではありませんでした。
1.5B Q4_K_M ビルドは、簡潔さを保ち、キャラクターの中に留まり、CPU セーフなパスでゲームの境界を尊重できる、私がテストした最小のバージョンでした。
3,057 の重複排除された監視付き微調整行
1,953 の選択されたプリファレンスと拒否されたプリファレンスのペア
40 の作成されたマルチターン会話 / 120 のプレフィックス
1,944 個のユニークなアシスタント出力
トレーニング シーケンスは、800 反復の MLX-LM LoRA パス、量子化モデルで観察された障害を使用した 160 反復の修復パス、480 反復でした。

リファレンスフリーの CPO スタイルの優先パスと、最終的な 80 回の反復による修復磨き。
コーパスは、プロットの状態ではなく、音声と動作を教えます。安定した制御タグは、キャラクター、会話モード、信頼関係層、意図、安全境界、および出力コントラクトを識別します。 LoRA は独自の声を持っています。実行時の取得により事実が提供されます。
JavaScript は、誰が話しているのか、会話がプライベートかパブリックか、ミッションがライブかどうか、プレイヤーが実際に観察した内容、どの関係層が適用されるか、そしてリクエストがモデルに到達する必要があるかどうかを選択します。
レンダラーは、目的の完了に関してモデルの出力を決して信頼しません。生成されたクレームでは、ドアのロックを解除したり、証拠を授与したり、等級を変更したり、関係を変更したり、買い手を発明したり、ミッションを終了したりすることはできません。
しっかりと根拠のある反応の場合、プロンプトには標準的な事実とアンカー用語の短いリストが含まれています。バリデーターは、サポートされていない数値、識別子、エンティティ、結果、グレード、プルーフ カウント、または熱変化を拒否します。また、応答は、この状態を即興で回避するのではなく、この状態に反応していることを示すために、提供されたアンカーを十分に保持する必要があります。
ドラフトが失敗した場合、Signal Ghost は、音声のずれ、繰り返し、グラウンディング、モード、またはその他の契約違反など、失敗を説明する 1 つの書き換えリクエストを受け取ります。それでも書き換えが失敗する場合、ゲームは作成された行を使用します。
そのフォールバックはエラー画面ではありません。それは物語システムの一部です。
記憶はデータ構造であり、トランスクリプトではありません
「登場人物たちがあなたのことを覚えている」というのは、簡単に言い過ぎになりがちです。 Signal Ghost はすべてを覚えているわけではありません。また、会話履歴全体をモデルに送信することもありません。
繰り返し登場する各キャラクターには、明示的に永続化されたレコードがあります。つまり、ローリングサマリー、最近の完全なターン、抽出された継続的な事実、合計および圧縮されたターン数、および最後のインタラクションのタイムスタンプです。

。
事実の抽出では、議論されたアーティファクト、信頼または脆弱性のプレッシャー、明示的に名前が付けられたキャラクター、「これを覚えておいてください」という合図、プレイヤーが選んだ名前、関係性の記述、および恐怖、怒り、孤独、疲労などの気分の合図など、ゲームで劇的に使用できるものを探します。
取得は意図的に小さく、理解しやすいものにしています。ベクトル検索ではなく、語彙と関連性に基づいてランク付けされます。通常の応答では、最も関連性の高い 3 つの事実と、最も関連性の高い 3 つのターンが返されます。明示的なリコール要求では、5 つのファクトと 4 つのターンに加えて、古い要約フラグメントを受け取ることができます。正確な事実が存在しない場合、プロンプトはキャラクターに記憶を捏造するのではなくギャップを認めるように指示します。
テストでは関連性が年表を上回ります。 1.5B ローカル モデルは、現在の質問に答える 3 つの交換が与えられた場合の方が、最近の任意の 8 つのターンとテキストのうねる壁が与えられた場合よりも優れたパフォーマンスを示しました。
パブリック IRC は、別の問題を追加します。プレイヤーは、他の誰かが話している間に 1 人のキャラクターについて言及することができます。適切な場合、Signal Ghost は、言及されたキャラクターの「経由で聞いた」記憶を保存します。これにより、すべての登場人物がすべてのプライベート メッセージに全知的にアクセスできるかのように装うことなく、部屋を越えた会話にある程度の継続性が与えられます。
思い出されたセリフがパフォーマンスを彩ります。証拠を書き換えることはありません。
ローカル モデルをエンジン サブシステムのように扱う
出荷用ランタイムでは、永続的な llama-server プロセスが優先されます。ゲームは、アプリケーションが読み込まれた直後にプリウォーミングを開始し、モデルを常駐させ、ランダムなループバック専用ポート 127.0.0.1 経由で完了リクエストを送信します。 Web UI は無効になっており、サーバーには推論スロットが 1 つあります。
優先サーバー
→ CPUセーフサーバー
→ ワンショットラマクリ
→ 作成したダイアログ 起動タイムアウトは 90 秒、生成リクエストはカッペです

d 45秒。障害が発生したサーバーは、継続的に再起動されるのではなく、短いバックオフに入ります。ゲームはシャットダウン中にアクティブなモデルプロセスを取得するため、Steam が「停止」のままになることはありません。
コンテキストは意図的に小さくなっています。Linux CPU では 1,024 トークン、他のデスクトップ ビルドでは 1,280 トークン、Linux および Steam Deck Vulkan パスでは 1,536 トークンです。通常、応答は 56 ～ 72 個の生成されたトークンを受け取りますが、プラットフォームに依存する上限は 72 ～ 96 です。 IRC のキャラクターには、研究論文のコンテキスト ウィンドウではなく、状態を認識した鋭い反応が必要です。
Metal を強制すると開発ホストでクラッシュするため、macOS は現在 CPU セーフな Accelerate/BLAS ルートを使用しています。 Linux と Steam Deck はバンドルされた Vulkan ランナーを優先し、CPU にフォールバックします。 Windows は現在、CPU セーフ パスを使用しています。
パフォーマンスにはばらつきがあるため、1 つの開発マシンをすべてのプレイヤーの結果として扱うべきではありません。これはローエンドのデータポイントです。16 GB RAM、8 つの CPU スレッド、1,280 トークンのコンテキストを搭載し、GPU レイヤーを備えていない 2018 Intel Core i7-8750H MacBook Pro です。
繰り返されるプレフィックスの数字は、llama.cpp のプロンプト キャッシュの恩恵を受けており、典型的なゲームプレイの遅延として解釈されるべきではありません。実際のプロンプトはさらに長く、ミッションの状態、記憶、プレイヤーの入力によって変化します。
あるベンチマーク出力には、「それがあなたが探しているものであれば、私がお手伝いします。」と書かれていました。それは安全で構造的には有効でしたが、当たり障りのないものでした。ルールベースの検証では、退屈な文をすべて認識するよりも、漏洩したプロンプトやでっち上げられた事実をより確実に拒否できます。
間違ったプロンプト形式は悪い動作のように見えました
Qwen2.5-Instruct はネイティブ ChatML ロール マーカーを想定しています。一般的な形式により拒否が発生しました。 ChatML と正しいストップ トークンにより、いくら散文を書いても解決できなかった問題が解決されました。
ブリッジは、ランナー バナー、エコーされた ChatML、コントロール トークン、スピーカー プレフィックス、コード フェンス、タイミングを取り除きます。

文章。登場人物の聖書のオウム返し、隠されたメトリクス、ハンドル リーク、トランスクリプト エコー、現実世界の侵入の詳細を拒否します。
キャラクターは 1 人の役に立つアシスタントに崩壊します
データには、拒否された十字文字の例が含まれています。サンプリングはキャラクターとモードによって異なります。証拠の質問はより厳密です。個人のダウンタイムはさらに異なる場合があります。
人間関係の対話が仕事に戻ってしまった
ダウンタイム、業務中、パブリックチャネル、ミッション、およびダイレクトメッセージの登録は個別に行われます。バリデーターは、感情的な回答が属する作業言語にフラグを立てます。
物語性のあるパッケージになった
応答が得られない場合、キャラクターに説得力はありません。プロセスの存続期間、プラットフォーム ライブラリ、コンテキスト バジェット、およびフォールバックはすべて、会話が占有されているか中断されているかを決定します。
スナップショット文を使用しない動作のテスト
正確な文字列のスナップショットは、サンプリングされたダイアログには適していません。代わりに契約をテストします。スピーカータグやマークダウンはなく、プロンプトの漏洩はなく、非表示のスコアや支払いはなく、現実世界の虐待の詳細はなく、キャラクター間の汚染はなく、個人的なダウンタイムでのミッション言語はなく、プライベートコンテキストが間違ったキャラクターに漏れることはありません。
裏切りや記憶の手がかりは、後のプロンプトに影響を与える可能性があります。
プロンプト抽出の試行では、作成されたワールド内境界を受け取ります。
ランナーが見つからなかったり失敗したりしても、セーブデータが破損したり、ミッションがブロックされたりすることはありません。
根拠のある反応は、裏付けのない実体、数値、結果をでっち上げることはできません。
その合計は、1,657 のゴールデン行、145 のキャストカバレッジ行、95 のメモリ行、120 のマルチターン プレフィックス、および 56 の修理ケースをカバーします。これらは、作成されたサンプルと契約の備品を検証します。これらは、すべての確率的出力が完璧であるという主張ではありません。
別のライブ モデル プローブは、複数のシードを使用して量子化された GGUF を実行します。区別が重要です。モデルは厳選されたセットを通過しても、3 番目のセットでは汎用になることができます。

ベンチマークが示したように、サンプル。
Signal Drift は 6 月 29 日にリリースされました。7 月 3 日、私は最初の完全な Signal Ghost リリースを出荷しました。これは、永続的なランナー、よりオープンな会話、より強力な声、関係のコンテキスト、パブリック IRC クロストーク、およびミッションを意識したフォールアウトです。 7 月 8 日には、より広範なオンボーディングとナラティブのアップデートを行いました。ガイド付きスターター侵入、より多くの音声と音楽、漏洩した録音、個人メール、より強力な継続性、そしてプレイヤーの足跡をより多く費やすフィナーレです。
2 番目の更新は重要な注意事項です。生成されたダイアログは、作成された構造を置き換えることはできません。キャラクターは見事に答えることができますが、プレイヤーが端末を理解できなかったり、証拠を見つけられなかったり、ある契約が次の契約にどのようにつながっているのかが見えなかったりすると、システムは意味のある反応をすることができません。
モデルは負のスペースを埋めます。建物を支える構造物ではありません。
他のゲームにも取り入れたい 5 つの教訓
01 モデルに狭いドラマチックな仕事を与えます。 「このキャラクターの反応を実行する」は、「ストーリーを実行する」よりも安全で便利です。
02 正規状態をモデルの外に保持します。流暢さは権威ではありません。
03 メモリを明示的に設計します。保持、取得、忘れ、正直なギャップはゲームデザインの決定事項です。
04 レイテンシーと障害を次のように扱います。

[切り捨てられた]

## Original Extract

How Signal Drift keeps a bundled local language model inside authored game state.

Fluency Is Not Authority — Building Local AI Characters
SIGNAL DRIFT / ENGINEERING NOTE 001 LOCAL INFERENCE · NARRATIVE SYSTEMS · POSTMORTEM Fluency Is
Not Authority
I shipped a 940 MiB, 1.5-billion-parameter language model inside a cross-platform indie game. Generating dialogue was the easy part. The difficult part was making sure a fluent character could never rewrite the game simply by sounding certain.
All inference for Signal Ghost runs on the player's device. Canonical game state remains authored.
One of the easiest ways to break a narrative game is to let a convincing sentence become true.
The scene I kept returning to while building Signal Drift was simple. The player breaks into a fictional host, finds a log that contradicts the client's briefing, and confronts that client in IRC. A conventional dialogue tree can handle the accusations I anticipated, but not every interpretation a player might type. A language model can respond to almost anything—but an unconstrained model may invent a confession, a new conspirator, or an outcome that never occurred.
I did not need a model to write the plot. I needed it to let a character inhabit a plot the game already knew was true.
CORE RULE The model performs the character. The authored game decides reality.
Why free-text dialogue belonged in a terminal game
Signal Drift is a CRT-noir hacking thriller built around fictional terminals, logs, mail, IRC rooms, evidence, and readable Linux-style commands. Players scan hosts, recover credentials, move through SSH sessions, reconstruct events, and decide whether to expose, sell, alter, or erase what they find.
Typing is already the game's primary verb. Asking a character “why does this timestamp disagree with the audit log?” feels natural in a way that selecting a dialogue wheel beside a terminal would not.
The goal was not infinite dialogue. It was to cover the human space between authored beats: challenging a lie, asking what a discovered file implies, provoking a rival, revisiting a betrayal, or talking to someone after the contract is over.
Inference runs locally, without an external AI service.
The same feature ships on Windows, macOS, and Linux.
JavaScript remains authoritative for state and safety.
Characters receive designed memory, not endless history.
Local inference eliminated an API dependency and recurring inference bill, but it did not make the feature free. It moved the cost into download size, memory, startup time, cross-platform packaging, hardware variance, and support.
Signal Ghost uses Qwen2.5-1.5B-Instruct , an Apache-2.0 model with roughly 1.54 billion parameters. I trained a Signal Ghost v19 dialogue LoRA, fused it into the base model, converted the result to GGUF, and quantized it to Q4_K_M.
The shipping file is 986,048,512 bytes—about 940 MiB. It runs through llama.cpp , pinned in the current build to b9835.
The smaller smoke-test model would announce that it was “just a computer program,” fall into generic security prose, or keep talking after the dramatic moment had passed. Compact was not the same as usable.
The 1.5B Q4_K_M build was the smallest version I tested that could remain terse, stay inside a character, and respect the game's boundaries on a CPU-safe path.
3,057 deduplicated supervised fine-tuning rows
1,953 chosen-versus-rejected preference pairs
40 authored multi-turn conversations / 120 prefixes
1,944 unique assistant outputs
The training sequence was an 800-iteration MLX-LM LoRA pass, a 160-iteration repair pass using failures observed in the quantized model, a 480-iteration reference-free CPO-style preference pass, and a final 80-iteration repair polish.
The corpus teaches voice and behavior, not plot state. Stable control tags identify character, conversational mode, rapport tier, intent, safety boundary, and output contract. The LoRA owns voice; runtime retrieval supplies facts.
JavaScript chooses who is speaking, whether the conversation is private or public, whether a mission is live, what the player has actually observed, which relationship tier applies, and whether the request should reach the model at all.
The renderer never trusts model output for objective completion. A generated claim cannot unlock a door, award evidence, change a grade, alter rapport, invent a buyer, or close a mission.
For tightly grounded reactions, the prompt contains a short list of canonical facts and anchor terms. The validator rejects unsupported numbers, identifiers, entities, outcomes, grades, proof counts, or heat changes. It also requires the reply to preserve enough supplied anchors to show that it is reacting to this state rather than improvising around it.
If a draft fails, Signal Ghost gets one rewrite request explaining the failure—voice drift, repetition, grounding, mode, or another contract violation. If the rewrite still fails, the game uses the authored line.
That fallback is not an error screen. It is part of the narrative system.
Memory is a data structure, not a transcript
“The characters remember you” can easily become an overclaim. Signal Ghost does not remember everything, and I do not send an entire conversation history back through the model.
Each recurring character has an explicit persisted record: a rolling summary, recent full turns, extracted standing facts, total and compacted turn counts, and a last-interaction timestamp.
Fact extraction looks for things the game can use dramatically: artifacts discussed, trust or vulnerability pressure, explicitly named characters, “remember this” cues, the player's chosen name, relationship statements, and mood cues such as fear, anger, loneliness, or exhaustion.
Retrieval is deliberately small and understandable. It is lexical and relevance-ranked rather than vector search. An ordinary reply receives the three most relevant facts and three most relevant turns. An explicit recall request can receive five facts and four turns, plus an older summary fragment. If the exact fact is absent, the prompt tells the character to admit the gap rather than fabricate memory.
Relevance beat chronology in testing. A 1.5B local model performed better when given the three exchanges that answered the current question than when given eight arbitrary recent turns and a rolling wall of text.
Public IRC adds another wrinkle: a player can mention one character while someone else is speaking. When appropriate, Signal Ghost stores an “overheard via” memory for the mentioned character. That gives cross-room conversation some continuity without pretending every character has omniscient access to every private message.
Remembered dialogue colors performance. It does not rewrite evidence.
Treating a local model like an engine subsystem
The shipping runtime prefers a persistent llama-server process. The game begins prewarming it shortly after the application loads, keeps the model resident, and sends completion requests over a random loopback-only 127.0.0.1 port. The web UI is disabled, and the server has one inference slot.
preferred server
→ CPU-safe server
→ one-shot llama-cli
→ authored dialogue The startup timeout is 90 seconds and a generation request is capped at 45 seconds. A failed server enters a short backoff rather than being relaunched continuously. The game reaps active model processes during shutdown so Steam does not remain stuck on “Stop.”
Context is intentionally small: 1,024 tokens on Linux CPU, 1,280 on other desktop builds, and 1,536 on the Linux and Steam Deck Vulkan path. Replies usually receive 56–72 generated tokens, with platform-dependent caps of 72–96. A character in IRC needs a sharp, state-aware reaction—not a research-paper context window.
macOS currently uses the CPU-safe Accelerate/BLAS route because forcing Metal crashed on the development host. Linux and Steam Deck prefer a bundled Vulkan runner and fall back to CPU. Windows currently uses the CPU-safe path.
Performance varies enough that one developer machine should not be treated as every player's result. This is a lower-end datapoint: a 2018 Intel Core i7-8750H MacBook Pro with 16 GB RAM, eight CPU threads, a 1,280-token context, and no GPU layers.
The repeated-prefix figures benefit from llama.cpp's prompt cache and should not be read as typical gameplay latency. Real prompts are longer and change with mission state, memory, and player input.
One benchmark output said, “I am here to help if that is what you are looking for.” It was safe and structurally valid, but bland. Rule-based validation can reject a leaked prompt or invented fact more reliably than it can recognize every boring sentence.
The wrong prompt format looked like bad behavior
Qwen2.5-Instruct expects native ChatML role markers. A generic format caused refusals. ChatML plus the correct stop token fixed a problem no amount of character prose would have solved.
The bridge strips runner banners, echoed ChatML, control tokens, speaker prefixes, code fences, and timing text. It rejects character-bible parroting, hidden metrics, handle leaks, transcript echoes, and real-world intrusion detail.
Characters collapse into one helpful assistant
The data includes rejected cross-character examples. Sampling differs by character and mode: evidence questions are tighter; personal downtime can vary more.
Relationship dialogue drifted back into work
Downtime, on-job, public-channel, mission, and direct-message registers are separate. Validators flag work language where an emotional answer belongs.
Packaging became narrative quality
A character is not convincing if the response never arrives. Process lifetime, platform libraries, context budgets, and fallbacks all shape whether a conversation feels occupied or broken.
Testing behavior without snapshotting sentences
Exact-string snapshots are a poor fit for sampled dialogue. I test contracts instead: no speaker tags or Markdown, no prompt leakage, no hidden score or payout, no real-world abuse detail, no cross-character contamination, no mission language in personal downtime, and no private context leaking into the wrong character.
A betrayal or memory cue can influence a later prompt.
A prompt-extraction attempt receives an authored, in-world boundary.
A missing or failed runner cannot corrupt a save or block a mission.
A grounded reaction cannot invent unsupported entities, numbers, or outcomes.
That total covers 1,657 golden rows, 145 cast-coverage rows, 95 memory rows, 120 multi-turn prefixes, and 56 repair cases. These validate authored examples and contract fixtures; they are not a claim that every stochastic output is perfect.
A separate live-model probe runs the quantized GGUF with multiple seeds. The distinction matters. A model can pass a curated set and still become generic on the third sample, as the benchmark demonstrated.
Signal Drift launched on June 29. On July 3, I shipped the first full Signal Ghost release: the persistent runner, more open conversation, stronger voices, relationship context, public IRC cross-talk, and mission-aware fallout. On July 8, I followed with a wider onboarding and narrative update—guided starter intrusion, more voice and music, leaked recordings, personal mail, stronger continuity, and a finale that spends more of the player's trail.
The second update is an important reminder: generated dialogue cannot replace authored structure. A character can answer beautifully, but if the player does not understand the terminal, cannot locate the evidence, or cannot see how one contract connects to the next, the system has nothing meaningful to react to.
The model fills negative space. It is not the structure holding the building up.
Five lessons I would carry into another game
01 Give the model a narrow dramatic job. “Perform this character's reaction” is safer and more useful than “run the story.”
02 Keep canonical state outside the model. Fluency is not authority.
03 Design memory explicitly. Retention, retrieval, forgetting, and honest gaps are game-design decisions.
04 Treat latency and failure as in

[truncated]
