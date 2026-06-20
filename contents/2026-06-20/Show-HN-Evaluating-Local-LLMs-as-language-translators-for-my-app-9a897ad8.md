---
source: "https://lector.dev/eval/"
hn_url: "https://news.ycombinator.com/item?id=48604349"
title: "Show HN: Evaluating Local LLMs as language translators for my app"
article_title: "Which local LLM translates best? A reproducible eval"
author: "3stacks"
captured_at: "2026-06-20T00:57:24Z"
capture_tool: "hn-digest"
hn_id: 48604349
score: 4
comments: 2
posted_at: "2026-06-19T23:06:45Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Evaluating Local LLMs as language translators for my app

- HN: [48604349](https://news.ycombinator.com/item?id=48604349)
- Source: [lector.dev](https://lector.dev/eval/)
- Score: 4
- Comments: 2
- Posted: 2026-06-19T23:06:45Z

## Translation

タイトル: HN を表示: アプリの言語翻訳者としてローカル LLM を評価する
記事のタイトル: どのローカル LLM が最適に翻訳されますか?再現可能な評価
HN テキスト: この種の評価を実行するのはこれが初めての試みなので、方法論に関するフィードバックが欲しいです。ネイティブ スピーカーから新しい翻訳を取得しない限り、ソースがモデルの入力に含まれていないことを保証することはできませんが、トップ モデルを使用した経験から、それらは非常に正確であると感じます。比較的小規模な言語からのやや不明瞭なテキストに遭遇した場合でも、適切な慣用的な意味に関しては、翻訳は通常 Google 翻訳よりも優れています。

記事本文:
どのローカル LLM が最もよく翻訳されますか?再現可能な評価
2026-06-19 · ソース → 英語 · アフリカーンス語 / ドイツ語 / スペイン語
ローカル LLM の翻訳能力はどのくらいですか?実際にクラウドが必要ですか?
24 のオンデバイス、セルフホスト、クラウド モデルの再現可能なベンチマークを英語に翻訳し、
リソースの少ないケース (アフリカーンス語) が前面と中央にあります。見出し: アフリカーンス語→英語でローカル 18 GB
モデルはフロンティア クラウドと統計的に結びつきます。同じ盲目タトエバ文、同じプロンプト、貪欲
デコード、COMET (意味) と chrF++ (表面) によるマルチリファレンスのスコア付け。
Lector の翻訳モデルを選択するために構築されました。
オープンで再現可能: ハーネス、すべてのモデルの生の出力、およびシードされたテスト セット:
github.com/heuwels/llm-lang-eval
1 つのズームされた COMET 軸上の 3 つの言語にわたる 24 のモデルすべて: 各行がモデル
(アフリカーンス語でランク付け)、言語ごとに 1 つのドット、そのコネクターが言語を越えて広がります。ズームにより、
微妙な違いが判読可能。 ~1.5 COMET 未満のギャップはサンプリング ノイズです (「有意性」を参照)。
88 90 92 94 96 gpt-5 gpt-5 · アフリカーンス語: 95.3 gpt-5 · ドイツ語: 93.1 gpt-5 · スペイン語: 93.9 gemini-2.5-pro gemini-2.5-pro · アフリカーンス語: 95.3 gemini-2.5-pro · ドイツ語: 93.5 gemini-2.5-pro · スペイン語: 94.4 クロード・オプス-4.8 クロード・オプス-4.8 · アフリカーンス語: 95.1 クロード・オプス-4.8 · ドイツ語: 93.3 クロード・オプス-4.8 · スペイン語: 94.5 クロード・ソネット-4.6 クロード・ソネット-4.6 · アフリカーンス語: 95.1 クロード・ソネット-4.6 · ドイツ語: 93.1 claude-sonnet-4.6 · スペイン語: 94.6 gemma-4-12b-qat gemma-4-12b-qat · アフリカーンス語: 95.0 gemma-4-12b-qat · ドイツ語: 93.1 gemma-4-12b-qat · スペイン語: 93.6 gpt-4o-mini gpt-4o-mini · アフリカーンス語: 95.0 gpt-4o-mini · ドイツ語: 92.9 gpt-4o-mini · スペイン語: 94.3 gpt-4o gpt-4o · アフリカーンス語: 94.9 gpt-4o · ドイツ語: 93.4 gpt-4o · スペイン語: 94.2 misstral-large misstral-larg

e · アフリカーンス語: 94.8 ミストラル-ラージ · ドイツ語: 93.2 ミストラル-ラージ · スペイン語: 94.0 ラマ-3.3-70b ラマ-3.3-70b · 英語: 94.7 ラマ-3.3-70b · ドイツ語: 92.9 ラマ-3.3-70b · スペイン語: gemini-2.5-flash gemini-2.5-flash · 英語: 94.7 gemini-2.5-flash · ドイツ語: 92.4 gemini-2.5-flash · スペイン語: 92.4 deepseek-v3.2 deepseek-v3.2 · 英語: 94.7 deepseek-v3.2 · ドイツ語: deepseek-v3.2 · 英語: 94.0 gemma-2-27b gemma-2-27b・英語：94.5 gemma-2-27b ・ドイツ語：91.2 gemma-2-27b ・英語：92.2 gemma-4-12b gemma-4-12b ・英語：94.5 gemma-4-12b ・ドイツ語：93.0 gemma-4-12b ・英語：93.3 gemma-3-12b ・英語：94.5 gemma-3-12b · ドイツ語: 92.2 gemma-3-12b · 英語: 92.6 gemma-3-27b · 英語: 92.6 gemma-3-27b · 英語: 92.9 gemma-3-27b · 英語: 94.3 gemma-3-27b · 英語: ministral-3-14b · ドイツ語: 91.8 ministral-3-14b ·英語: 93.5 gemma-4-e4b · 英語: 94.2 gemma-4-e4b · ドイツ語: 92.2 gemma-4-e4b · 英語: 94.1 qwen3.5-9b · ドイツ語: 92.2 qwen3.5-9b · スペイン語: 93.0 misstral-small-3.2-cloud misstral-small-3.2-
[切り捨てられた]
chrF++ (サーフェス) 上の COMET (つまり、×100)、言語ごと。 chrF++
報酬文字が参照と重複するため、有効な言い換えがドッキングされます
(「景色は素晴らしいです」 vs 「景色は息を呑むほどです」);彗星
意味をスコアし、クレジットします。緑 = 先行バンド (~1.5 COMET 以内、a
統計的には同点であり、勝者は 1 名ではありません）。行はチャートの順序 (英語の COMET によるランク付け) を共有しているため、
ほぼ同点のスコアは、スコアが均等であっても、離れた場所に位置する可能性があります。言語ごとに n=200。
費用とできること

味方が走る
クラウド モデルがそのボードの最上位を占めていますが、これはローカル モデルの研究であり、ほとんどの
クラウド フィールドはボックス上ではまったく実行できません。フロンティア API (GPT-5、Claude、Gemini) はクローズド ウェイトです。
したがって、それらを自己ホストするという選択肢は決してありませんでした。 OpenRouter を通じてアクセスできるオープン モデルのほとんどは、
ハードウェアとしては大きい: Llama 3.3 は 70B、Mistral Large はさらに大きく、さらには 24B と 27B のオープン モデル
(Mistral Small、Gemma 2 および 3 は 27B)、OS と KV がインストールされると、18 GB Mac の天井または天井を超えて設置されます。
キャッシュがそのシェアを奪います。本当に適しているのは、オンデバイス層とセルフホストボックス層です。
独自のフィールド。
88 90 92 94 96 gemma-4-12b-qat gemma-4-12b-qat · アフリカーンス語: 95.0 gemma-4-12b-qat · ドイツ語: 93.1 gemma-4-12b-qat · スペイン語: 93.6 gemma-4-12b gemma-4-12b · アフリカーンス語: 94.5 gemma-4-12b · ドイツ語: 93.0 gemma-4-12b · スペイン語: 93.3 gemma-3-12b gemma-3-12b · アフリカーンス語: 94.5 gemma-3-12b · ドイツ語: 92.2 gemma-3-12b · スペイン語: 92.6 ministral-3-14b ministral-3-14b · アフリカーンス語: 94.3 ministral-3-14b · ドイツ語: 91.8 ministral-3-14b · スペイン語: 93.5 gemma-4-e4b gemma-4-e4b · アフリカーンス語: 94.2 gemma-4-e4b · ドイツ語: 92.2 gemma-4-e4b · スペイン語: 93.1 qwen3.5-9b qwen3.5-9b ·アフリカーンス語: 94.1 qwen3.5-9b · ドイツ語: 92.2 qwen3.5-9b · スペイン語: 93.0 gemma-3n-e4b gemma-3n-e4b · アフリカーンス語: 93.7 gemma-3n-e4b · ドイツ語: 92.3 gemma-3n-e4b · スペイン語: 92.9 ollama-llama3.1-8b ollama-llama3.1-8b · アフリカーンス語: 93.5 ollama-llama3.1-8b · ドイツ語: 91.6 ollama-llama3.1-8b · スペイン語: 92.6 apfel-foundation apfel-foundation · アフリカーンス語: 92.1 apfel-foundation · ドイツ語: 89.6 apfel-foundation · スペイン語: 90.9 qwen2.5-coder-14b qwen2.5-coder-14b · アフリカーンス語: 91.8 qwen2.5-coder-14b · ドイツ語: 91.4 qwen2.5-coder-14b · スペイン語: 93.0 アフリカーンス語 ドイツ語 スペイン語 · COMET、

軸88から96にズーム。各名前の前のドット = 層
実際に適合する最も強力なモデルは、7.5 GB の gemma-4-12b-qat です。
フロンティアバンドの一番上に座っているのと同じ人です。 apfel-foundation は Apple の組み込み Foundation モデルであり、macOS に同梱され、Apfel ハーネスを通じて実行されます。回答内容については立派なスコアを獲得しているが、アフリカーンス語の文の 26% (200 文中 52 文) で拒否またはエラーがあり、スペイン語では 8% であった。Apple はサポート言語にアフリカーンス語をリストしていないため、最下位付近に位置している。
コストはほとんどかかりません。翻訳は非常に小さく、約 80 トークンなので、クラウド全体が
スイープ (3 つの言語にわたる 24 のモデル、およびホールドアウトとクローズ プローブ) は、OpenRouter で 13.62 ドルになりました。
フロンティアモデルであっても、翻訳あたり 1 セント未満です。トークンごとの価格は依然として 1 つまたは 2 つの注文にまたがります
Gemini Flash や GPT-4o-mini のような安価な層に対するフロンティア API の規模は大きいですが、この点では
どちらに行っても、絶対的な請求額は少額です。
決断を左右するのは、あなたがすでに所有しているものです。予備の Mac は埋没コストなので、ローカル モデル
電気を超えて検索するたびに無料です。既存のクロードプランも、その範囲内でマージンが無料です。
これが、私が最初に Anthropic OAuth ルートに到達する理由です。 OpenRouter は 3 つのうちの唯一の 1 つです
これにより実際のトークンごとの請求が追加され、自己ホストできない特定のモデルが必要な場合にその役割を果たします。
の計画はありません。したがって、正直な答えは通常、手元にあるものを使用することです。予備のボックスでローカルのコンピュータを実行します。
このモデルでは、既存のプランがすでにルックアップをカバーしており、どちらも使用しない場合、安価なクラウド層は 1 件当たり数ペニーです。
千。 12B 以来、自宅で走ることができ、すでにアフリカーンス語のフロンティアを結び付けており、1 人あたりのフロンティア料金を支払っています。
この特定のジョブではトークンはほとんど購入しません。
汚染チェック:

目に見えないデータで生き残れるのか？
注意点については注意してください。正確なトレーニング終了日はすべてのモデルについて公開されているわけではありません。
最近追加された文は文体や難易度が微妙に異なる場合があるため、小さな Δ は「意味」ではなく「耐えられる」と読んでください。
汚染を正確に測定します。
オウム返しのプローブ: 記憶、直接測定
回復 = 空白になった単語の完全一致。大きなプラスのギャップ = 暗記。ゼロに近い
= 本物のコンテキスト予測。セルあたり n≈150 なので、~±10 以内のギャップはノイズになります。
数字だけが多くを物語っています。実際の翻訳は次のとおりです。
ほとんどが同意しません。緑色は、その文の文ごとの chrF++ が最も高いことを示します。
アフリカーンス語: モデルが分割される場所
モデルが明確な陣営に分かれ、複数のモデルが 1 つの文言に同意し、複数のモデルが別の文言に同意する文 (1 回限りの文言は最後まで折りたたまれます)。キャンプごとに × ティアドットをカウントします。緑 = 最も近い基準キャンプから約 6 chrF 以内。
ドイツ語: モデルが分割される場所
モデルが明確な陣営に分かれ、複数のモデルが 1 つの文言に同意し、複数のモデルが別の文言に同意する文 (1 回限りの文言は最後まで折りたたまれます)。キャンプごとに × ティアドットをカウントします。緑 = 最も近い基準キャンプから約 6 chrF 以内。
スペイン語: モデルが分割される場所
モデルが明確な陣営に分かれ、複数のモデルが 1 つの文言に同意し、複数のモデルが別の文言に同意する文 (1 回限りの文言は最後まで折りたたまれます)。キャンプごとに × ティアドットをカウントします。緑 = 最も近い基準キャンプから約 6 chrF 以内。
タスク。ブラインドソース→英語。モデルはソース文のみを認識します。参考資料は採点のために差し出されます。
データ。 Tatoeba 文ペア (CC-BY 2.0 FR)、シードされたランダム サンプル、利用可能な場合は複数の参照、長さでフィルター処理。
プロンプト。 1 つの固定ユーザー メッセージ。モデル間で同一です (モデルごとの調整はありません)。思考連鎖が無効になっています (reasoning_effort: none )。翻訳は必要ありません

e.
構造化された出力。すべてのモデルは、 json_schema response_format を介して {"translation": "…"} を出力するように制約されます。これがイコライザーです。そうでない場合、小規模なモデルはプレーン テキストで「大声で考え」、答えをプリアンブルに埋め込んでしまいます。制約付きデコードではそれが不可能になり、すべてのモデルに同一の制約が与えられ、Lector 自体のプロンプトが反映されます。
デコード中。温度 = 0 (貪欲)、18 GB ホスト上に一度に 1 つのモデルが常駐します (JIT ロード/エビクト)。
メトリクス。 chrF++ および sacreBLEU 経由の BLEU (署名は記録されています)。 COMETの予定です。
重要性: ランクではなくバンドを読み取ります。言語ごとに n=200 (アフリカーンス語は主に単一参照)、したがって、システムごとの COMET 95% 信頼区間は、いずれにしてもおよそ 1 ～ 2 ポイントです。 ~1.5 COMET 未満の差はサンプリング ノイズです。緑色の先頭バンドは統計的な同点であり、その中の並べ替え順序は意味がありません。セグメントごとのブートストラップ CI は今後の課題です。
これはプロキシです。これは、レクターの実際の単語/フレーズの辞書検索タスクではなく、一般的な文の MT を測定します。モデル選択の強力なシグナルであり、「レクターの出力の採点」ではありません。
汚染、それが大きな問題です。 Tatoeba はこれらのモデルの事前トレーニングに含まれているため、高いスコアは言語についての推論ではなくペアの暗記を反映している可能性があり、スコアだけで 2 つを区別することはできません。上記の汚染チェックセクションでは、カットオフ後のホールドアウトでこれを制限しています。絶対スコアを疑念を持って扱い、前後の差分と見出しの数値に対する相対的なギャップを重視します。
英語化は簡単な方向であり、ここでのアフリカーンス語は主に単一参照です。それに応じてお読みください。
llm-lang-eval によって生成 · ハーネスと生の生成はオープンで再現可能 · Luke Boyle によって構築

## Original Extract

This is my first attempt at running an eval of this nature so would love some methodology feedback. I can't guarantee the sources weren't already in the model's inputs without getting novel translations from native speakers, but from my experience using the top models, they feel very accurate. Even encountering somewhat obscure texts from a relatively small language the translations generally beat Google Translate for proper idiomatic meaning.

Which local LLM translates best? A reproducible eval
2026-06-19 · source → English · Afrikaans / German / Spanish
How good are local LLMs at translation, and do you actually need the cloud?
A reproducible benchmark of 24 on-device, self-hosted, and cloud models translating into English, with the
low-resource case (Afrikaans) front and centre. The headline: on Afrikaans→English a local 18 GB
model lands in a statistical tie with frontier cloud. Same blinded Tatoeba sentences, same prompt, greedy
decoding, scored multi-reference with COMET (meaning) and chrF++ (surface).
Built to pick a translation model for Lector .
Open and reproducible: harness, every model's raw outputs, and the seeded test sets:
github.com/heuwels/llm-lang-eval
All 24 models across three languages on one zoomed COMET axis: each row a model
(ranked by Afrikaans), one dot per language, the connector its cross-language spread. The zoom makes
the tight differences legible; gaps under ~1.5 COMET are sampling noise (see Significance ).
88 90 92 94 96 gpt-5 gpt-5 · Afrikaans: 95.3 gpt-5 · German: 93.1 gpt-5 · Spanish: 93.9 gemini-2.5-pro gemini-2.5-pro · Afrikaans: 95.3 gemini-2.5-pro · German: 93.5 gemini-2.5-pro · Spanish: 94.4 claude-opus-4.8 claude-opus-4.8 · Afrikaans: 95.1 claude-opus-4.8 · German: 93.3 claude-opus-4.8 · Spanish: 94.5 claude-sonnet-4.6 claude-sonnet-4.6 · Afrikaans: 95.1 claude-sonnet-4.6 · German: 93.1 claude-sonnet-4.6 · Spanish: 94.6 gemma-4-12b-qat gemma-4-12b-qat · Afrikaans: 95.0 gemma-4-12b-qat · German: 93.1 gemma-4-12b-qat · Spanish: 93.6 gpt-4o-mini gpt-4o-mini · Afrikaans: 95.0 gpt-4o-mini · German: 92.9 gpt-4o-mini · Spanish: 94.3 gpt-4o gpt-4o · Afrikaans: 94.9 gpt-4o · German: 93.4 gpt-4o · Spanish: 94.2 mistral-large mistral-large · Afrikaans: 94.8 mistral-large · German: 93.2 mistral-large · Spanish: 94.0 llama-3.3-70b llama-3.3-70b · Afrikaans: 94.7 llama-3.3-70b · German: 92.9 llama-3.3-70b · Spanish: 93.8 gemini-2.5-flash gemini-2.5-flash · Afrikaans: 94.7 gemini-2.5-flash · German: 92.4 gemini-2.5-flash · Spanish: 92.4 deepseek-v3.2 deepseek-v3.2 · Afrikaans: 94.7 deepseek-v3.2 · German: 92.6 deepseek-v3.2 · Spanish: 94.0 gemma-2-27b gemma-2-27b · Afrikaans: 94.5 gemma-2-27b · German: 91.2 gemma-2-27b · Spanish: 92.2 gemma-4-12b gemma-4-12b · Afrikaans: 94.5 gemma-4-12b · German: 93.0 gemma-4-12b · Spanish: 93.3 gemma-3-12b gemma-3-12b · Afrikaans: 94.5 gemma-3-12b · German: 92.2 gemma-3-12b · Spanish: 92.6 gemma-3-27b gemma-3-27b · Afrikaans: 94.4 gemma-3-27b · German: 92.6 gemma-3-27b · Spanish: 92.9 claude-haiku-4.5 claude-haiku-4.5 · Afrikaans: 94.3 claude-haiku-4.5 · German: 92.9 claude-haiku-4.5 · Spanish: 94.4 ministral-3-14b ministral-3-14b · Afrikaans: 94.3 ministral-3-14b · German: 91.8 ministral-3-14b · Spanish: 93.5 gemma-4-e4b gemma-4-e4b · Afrikaans: 94.2 gemma-4-e4b · German: 92.2 gemma-4-e4b · Spanish: 93.1 qwen3.5-9b qwen3.5-9b · Afrikaans: 94.1 qwen3.5-9b · German: 92.2 qwen3.5-9b · Spanish: 93.0 mistral-small-3.2-cloud mistral-small-3.2-
[truncated]
COMET (meaning, ×100) over chrF++ (surface), per language. chrF++
rewards character overlap with the reference, so it docks valid paraphrases
( "scenery is magnificent" vs "landscape is breathtaking" ); COMET
scores meaning and credits them. Green = leading band (within ~1.5 COMET, a
statistical tie, not a single winner). Rows share the chart's order (ranked by Afrikaans COMET), so
near-ties can sit a place apart despite equal rounded scores. n=200 per language.
Cost, and what you can actually run
Cloud models fill the top of that board, but this is a study of local models, and most of
the cloud field can't run on the box at all. The frontier APIs (GPT-5, Claude, Gemini) are closed weights,
so self-hosting them was never an option. The open models I could reach through OpenRouter are mostly too
big for the hardware: Llama 3.3 is 70B, Mistral Large is larger again, and even the 24B and 27B open models
(Mistral Small, Gemma 2 and 3 at 27B) sit at or past the ceiling of an 18 GB Mac once the OS and the KV
cache take their share. What genuinely fits is the on-device and self-hosted-box tiers, so here is that
field on its own.
88 90 92 94 96 gemma-4-12b-qat gemma-4-12b-qat · Afrikaans: 95.0 gemma-4-12b-qat · German: 93.1 gemma-4-12b-qat · Spanish: 93.6 gemma-4-12b gemma-4-12b · Afrikaans: 94.5 gemma-4-12b · German: 93.0 gemma-4-12b · Spanish: 93.3 gemma-3-12b gemma-3-12b · Afrikaans: 94.5 gemma-3-12b · German: 92.2 gemma-3-12b · Spanish: 92.6 ministral-3-14b ministral-3-14b · Afrikaans: 94.3 ministral-3-14b · German: 91.8 ministral-3-14b · Spanish: 93.5 gemma-4-e4b gemma-4-e4b · Afrikaans: 94.2 gemma-4-e4b · German: 92.2 gemma-4-e4b · Spanish: 93.1 qwen3.5-9b qwen3.5-9b · Afrikaans: 94.1 qwen3.5-9b · German: 92.2 qwen3.5-9b · Spanish: 93.0 gemma-3n-e4b gemma-3n-e4b · Afrikaans: 93.7 gemma-3n-e4b · German: 92.3 gemma-3n-e4b · Spanish: 92.9 ollama-llama3.1-8b ollama-llama3.1-8b · Afrikaans: 93.5 ollama-llama3.1-8b · German: 91.6 ollama-llama3.1-8b · Spanish: 92.6 apfel-foundation apfel-foundation · Afrikaans: 92.1 apfel-foundation · German: 89.6 apfel-foundation · Spanish: 90.9 qwen2.5-coder-14b qwen2.5-coder-14b · Afrikaans: 91.8 qwen2.5-coder-14b · German: 91.4 qwen2.5-coder-14b · Spanish: 93.0 Afrikaans German Spanish · COMET, zoomed axis 88 to 96; dot before each name = tier
The strongest model that actually fits, gemma-4-12b-qat at 7.5 GB, is the
same one sitting in the frontier band up top. apfel-foundation is Apple's built-in Foundation model, the one that ships with macOS, run through the Apfel harness. It scores respectably on what it answers, but it refused or errored on 26% of the Afrikaans sentences (52 of 200), against 8% in Spanish, and Apple doesn't list Afrikaans among its supported languages, which is why it sits near the bottom.
Cost barely enters into it. A translation is tiny, roughly 80 tokens, so the entire cloud
sweep (24 models across three languages, plus the holdout and the cloze probe) came to $13.62 on OpenRouter,
well under a cent per translation even on the frontier models. Per-token pricing still spans an order or two
of magnitude, the frontier APIs against the cheap tiers like Gemini Flash or GPT-4o-mini, but at this token
count the absolute bill is small whichever way you go.
What moves the decision is what you already own. A spare Mac is a sunk cost, so a local model
is free per lookup beyond the electricity. An existing Claude plan is free at the margin too, within its
limits, which is why I reach for the Anthropic OAuth route first. OpenRouter is the only one of the three
that adds a real per-token bill, and it earns its place when you need a specific model you can't self-host or
don't have a plan for. So the honest answer is usually to use what you've got: a spare box runs a local
model, an existing plan already covers the lookups, and with neither, the cheap cloud tier is pennies per
thousand. Since a 12B you can run at home already ties the frontier on Afrikaans, paying frontier rates per
token buys very little for this particular job.
Contamination check: does it survive on unseen data?
Caveat on the caveat: exact training-cutoff dates aren't published for every model, and
recently-added sentences may differ subtly in style or difficulty, so read a small Δ as "holds up", not as a
precise measurement of contamination.
Parroting probe: memorisation, measured directly
Recovery = exact match of the blanked word. A large positive gap = memorisation; near-zero
= genuine context prediction. n≈150 per cell, so gaps within ~±10 are noise.
The numbers only say so much. Here are the actual translations where models
disagree most. Green marks the highest per-sentence chrF++ for that sentence.
Afrikaans: where the models split
Sentences where models split into clear camps , several agreeing on one wording, several on another (one-off wordings collapsed to a tail). Count × tier-dots per camp; green = within ~6 chrF of the closest-to-reference camp.
German: where the models split
Sentences where models split into clear camps , several agreeing on one wording, several on another (one-off wordings collapsed to a tail). Count × tier-dots per camp; green = within ~6 chrF of the closest-to-reference camp.
Spanish: where the models split
Sentences where models split into clear camps , several agreeing on one wording, several on another (one-off wordings collapsed to a tail). Count × tier-dots per camp; green = within ~6 chrF of the closest-to-reference camp.
Task. Blinded source → English. The model sees only the source sentence; references are held out for scoring.
Data. Tatoeba sentence pairs (CC-BY 2.0 FR), seeded random sample, multi-reference where available, length-filtered.
Prompt. One fixed user message, identical across models (no per-model tuning). Chain-of-thought disabled ( reasoning_effort: none ); translation needs none.
Structured output. Every model is constrained to emit {"translation": "…"} via a json_schema response_format . This is the equaliser: small models otherwise "think out loud" in plain text and bury the answer in preamble. Constrained decoding makes that impossible, gives every model the identical constraint, and mirrors how Lector itself prompts.
Decoding. temperature = 0 (greedy), one model resident at a time on an 18 GB host (JIT load/evict).
Metrics. chrF++ and BLEU via sacreBLEU (signatures recorded). COMET planned.
Significance: read bands, not ranks. n=200 per language (Afrikaans largely single-reference), so per-system COMET 95% confidence intervals are roughly 1 to 2 points either way. Differences below ~1.5 COMET are sampling noise: the green leading band is a statistical tie and the sort order within it is not meaningful. Per-segment bootstrap CIs are future work.
This is a proxy. It measures general sentence MT, not Lector's actual word/phrase dictionary-lookup task, a strong signal for model choice, not "Lector's output graded."
Contamination, the big one. Tatoeba is in these models' pretraining, so a high score can reflect memorising the pair rather than reasoning about the language, and the score alone can't separate the two. The contamination-check section above bounds this with a post-cutoff holdout; treat absolute scores with suspicion and weight the pre-vs-post deltas and relative gaps over the headline numbers.
Into-English is the easy direction , and Afrikaans here is largely single-reference. Read accordingly.
Generated by llm-lang-eval · harness & raw generations are open and reproducible · built by Luke Boyle
