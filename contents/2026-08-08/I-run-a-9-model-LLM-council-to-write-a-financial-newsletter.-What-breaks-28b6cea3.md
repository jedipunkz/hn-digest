---
source: "https://marketdaily.ai/blog/eng-llm-council-judge-en-202608"
hn_url: "https://news.ycombinator.com/item?id=49221691"
title: "I run a 9-model LLM council to write a financial newsletter. What breaks"
article_title: "I run a 9-model LLM council with a judge to write a daily financial newsletter. Here's what actually breaks. | MarketDaily"
author: "delvinchang"
captured_at: "2026-08-08T13:38:13Z"
capture_tool: "hn-digest"
hn_id: 49221691
score: 1
comments: 0
posted_at: "2026-08-08T13:30:01Z"
tags:
  - hacker-news
  - translated
---

# I run a 9-model LLM council to write a financial newsletter. What breaks

- HN: [49221691](https://news.ycombinator.com/item?id=49221691)
- Source: [marketdaily.ai](https://marketdaily.ai/blog/eng-llm-council-judge-en-202608)
- Score: 1
- Comments: 0
- Posted: 2026-08-08T13:30:01Z

## Translation

タイトル: 私は金融ニュースレターを書くために 9 モデル LLM 評議会を運営しています。何が壊れるのか
記事のタイトル: 私は毎日の金融ニュースレターを書くために、裁判官と一緒に 9 モデルの LLM 評議会を運営しています。実際に壊れたものは次のとおりです。 |マーケットデイリー
説明: 私は、毎日の財務メールのダイジェストを配信する MarketDaily を運営しています。購読者は保有銘柄 (米国株と台湾株) を選択し、システムは 1 日に 2 回、それぞれの銘柄に対してパーソナライズされた HTML レポートを生成し、定時に送信します。最初のコミットは 2026-05-19 でした。それを書いている時点で

記事本文:
私は毎日の財務ニュースレターを執筆するために、裁判官とともに 9 モデルの LLM 評議会を運営しています。実際に壊れたものは次のとおりです。 |マーケットデイリー
マーケットデイリー ←
すべての記事
マーケットデイリー・エンジニアリング
私は毎日の財務ニュースレターを執筆するために、裁判官とともに 9 モデルの LLM 評議会を運営しています。実際に壊れたものは次のとおりです。
私は、毎日の金融メールのダイジェストを配信する MarketDaily を運営しています。購読者は保有銘柄 (米国株と台湾株) を選択し、システムは 1 日に 2 回、それぞれの銘柄に対してパーソナライズされた HTML レポートを生成し、定時に送信します。最初のコミットは 2026-05-19 でした。執筆時点では 75 日間、1,800 以上のコミットのプロダクション履歴があり、現在 21 人のサブスクライバーにサービスを提供しています。規模は小さいですが、耐障害性は基本的にゼロです。これは、市場が開く前に誰かの受信箱に金融コンテンツが届くことです。幻覚のような目標株価はバグレポートではなく、信頼の葬式だ。
この投稿は、それを存続可能にする信頼性レイヤー、つまりマルチモデル評議会、裁判官、および 31 チェックの決定論的監査ゲートについて説明します。失敗した3件を含む。
「モデル + フォールバック」の問題
単純なアーキテクチャは、メイン モデルを呼び出し、エラーが発生した場合は別のモデルにフォールバックするというものです。これにより、ほとんど問題にならない障害モードが処理されます。実際に害を及ぼす障害モードはサイレント劣化です。クォータのプレッシャーの下で、モデルはレポートのようなものを返します — 正しい形式、正しいセクション、例外は発生しません — しかし、内部の分析は混乱しています。エラーコードは発生しません。ユーザーはあなたが知る前にそれを知ります。
したがって、設計原則は次のようになりました。単一のモデルの出力を決して信頼しない、また、モデルの自己評価を決して信頼しないということです。 3 つの分離されたレイヤー:
市場データ ──► 事前構造（価格 vs MA20/MA50 — 方向性は
LLM ではなく、決定論的なコードによって決定されます)
│
▼
評議会: 9 つの構成された無料枠シート

7 つのプロバイダーにわたって
(Gemini x2、Groq、Ollama 経由のローカル GPU、Cloudflare
Workers AI x2、OpenRouter、Cerebras、OpenAI —
API キーのない座席は自動スキップされます。 ~5 独立した
声は通常の日に生きています）
│ 各議席: {無駄のない、信念、論文、key_risk}
│ 定足数: 2 議席以上、または議会の評決なし
▼
裁判官: 合意と反対意見を総合する
(フリーチェーン: Gemini lite → Groq → 両方が死亡した場合、
そのまま最高有罪判決を受けた席に着く）
│
▼
レポートの生成 (別個の LLM チェーン)
│
▼
監査: 最終的な HTML に対する 31 の決定的チェック
│
高失敗 §──► 60 秒待って、より強力なモデルで再試行してください
│ │ まだ高い
│ ▼
│ 決定論的フォールバック (LLM なし、価格なし)
│ ターゲット — 悪質なメールが建物から出ることはありません）
▼
毎日、例外なく、決まった時間に送信します
議会: 意見の相違は特徴であって騒音ではない
すべての株式は、実行ごとに 1 つの評議会ラウンドを取得します (ユーザー間でキャッシュされます)。各座席は同一の実際の市場データに加えて、厳しい構造的制約を取得します。方向はコードによってロックされます (価格対 MA20 対 MA50)。強気構造で「ショート」と表示された議席は、単純な if ステートメントによって中立に降格されます。このシステムの LLM には方向を転換する権限はありません。評議会は、構造的に許容される範囲内の理論、リスク対策、リーンなど、モデルが実際に役立つレイヤーでのみ機能します。
次に、反対意見を測定します。全議席が賛成の場合は 0、議会に長短の両方が含まれている場合は 2、そうでない場合は 1 となります。反対意見は最後のカードまで伝播します。反対意見の多い銘柄は下流で保守的な表現を強制され、示された自信は押し下げられます。信頼性は、過去の校正テーブルによって個別に 75% に制限されます。監査では、それを超える数値は意見ではなく、防御が失敗したものとして扱われます。
座席の選択は、ベンチマーク ch ではなく、クォータ エンジニアリングによって行われます。

アシング。完全な開示: 評議会は 42 日目 (2026 年 6 月 30 日) に着陸しました。それ以前は、システムはまさに単純なモデル + フォールバック アーキテクチャでした。この投稿が批判によって始めたものです。午前 5 時に命を奪われるのは、モデルの品質ではなく、相関関係にあるクォータの枯渇だからです。名簿を形成する実際の制約は次のとおりです。
Groq の無料利用枠は、1 分あたり (8,000 TPM)、モデルごとに 1 日あたり (200,000 TPD) に制限されています。評議会席では、レポート生成チェーンとは異なるモデルが使用されています。これは、1 つの共有バケットにより、評議会が毎日 200,000 トークンのうち 197,364 個を消費し、45 件のレポート呼び出しが呼び出しごとに 144 秒のバックアップ モデルに失敗したことを意味するためです。その夜のダイジェストは1時間35分遅れだった。
1 つのシート (私自身の GPU 上のローカル 14B モデル) が存在するのは、純粋に割り当てがゼロでネットワーク依存性がゼロであるためです。クラウド DNS が突然停止したり、すべてのベンダーの割り当てが一度になくなったりした場合に、唯一生き残るのがこれです。
あるモデルでは、プロンプトから価格が書き換えられ、385.25 が 385.00 になりました。加工精度。価格に関わるものは永久に禁止されていますが、議会の出力は数字のない JSON 意見であるため、議会では許可されています。爆発半径設計はモデルの信頼を打ち破ります。
シートにはサーキット ブレーカーが設定されています。クォータのデッド/キーの紛失/3 回連続の失敗により、そのラウンドのシートが無効になります。 HTTP 402 (ビリングウォール) は最初の攻撃でそれを強制終了します。支払いエラーの再試行は全くの無駄です。
裁判官: 合成、出口あり
裁判官はすべての議席の意見を JSON として認識し、最も騒々しい議席をオウム返しにするのではなく、同意と反対の意見を総合するように指示されます。裁判官自体にはフォールバックチェーンがあり、チェーン全体が停止した場合、システムは最高有罪判決を受けた議席の論文をそのまま採用します。回答した議席が 2 議席未満の場合、議会の評決はまったく行われず、株価は消滅します。

単純な単一モデルの道を進みます。カウンシル全体はフェイルセーフです。失敗しても部分的な結果が返され、送信がブロックされることはありません。 「送信を見逃さない」と「ガベージを送信しない」は 2 つの異なる保証であり、2 つの異なるメカニズムによって強制されます。
監査: 31 枚の小切手、それぞれ血で支払われました
Audit_digest() は、最終的な HTML に対して 31 の名前付き決定的チェックを実行します。どれもユーザーが実際に怒った（または怒ったであろう）様子をマッピングしています。緊張した規律（午前7時に「今日台湾株が上がった」と書くことはできません。市場は9時に開きます）、保有銘柄のカバレッジ（ユーザーが選択したすべての銘柄にアクションカードが必要です）、捏造検出（プレースホルダーXXXトークン、偽のURL、データフィードに存在しない収益予測）、プロンプト命令の漏洩（LLMが独自の命令をユーザー向けの出力にコピーします）。切り捨て検出、未定義の CSS クラス。
高失敗トリガー: 60 秒間待機し (無料利用枠は 1 分ごとに制限されています。5 秒後に再試行すると、同じ 429 ウィンドウに到達し、より弱いモデルに到達します)、より強力なモデルを前面に強制的に再生成して、それでも失敗する場合は、決定論的なフォールバック (価格レベルを意図的に含まないプレーン コードで組み立てられたレポート) を出荷します。
システムを構築した3つの出来事
429事件。ジェミニの 1 日の割り当てがなくなりました。バックオフ付き再試行ロジックでは、実行全体で呼び出しごとに最大 100 秒かかりました。ダイジェストは5時間遅れで配信されました。修正: 1 つのモデルで 2 つの連続した 429 = その日の割り当てが無効になり、実行のためにモデル全体を融合します。一時的なエラーに対しては正しい再試行ロジックは、クォータ エラーに対しては有害であるため、分類する必要があります。
CSS 事件は 2 つの行為で構成されます。 「既存の CSS クラスの再利用」を求められた LLM は、ニアミス名を自由形式で news-headline の代わりに news-title に変更します。スタイルシートルールなし、セクション全体

n はスタイルなしでレンダリングされます。行為 1: unknown_css_class HIGH チェックを追加しました。第 2 幕、1 週間後: 月曜日版のプロンプトには逐語的な骨格が欠けていたため、すべてのモデルがクラスを発明し、12 人のユーザー全員がチェックにヒットし、再試行でもヒットしました (同じプロンプト、同じ症状)。9 人のユーザーが劣化したフォールバック バージョンを取得しました。守備側が事件となった。本当の修正は、チェック前の決定論的な修復レイヤーでした。既知のニアミス名を実際のクラスにマッピングし、未知のものを削除します (CSS ルールが存在しないため、削除は視覚的に何も行われません)。教訓が 2 つあります。LLM に「既存のクラスを使用する」と指示するのは仕様ではなく、文字通りのスケルトンです。そして、HIGH チェックを追加する前に、「全員が一度に失敗したらどうなるか」を尋ねてください。
深層崩壊事件。ある朝、すべての無料割り当てが融合され、世代は最も弱いモデルに落ち、ストックごとの推論は 129 ～ 165 文字から 48 ～ 64 文字に崩壊しました。フォーマットは完璧でした。価格は存在していました。すべてのチェックに合格しました。それを人間が捕まえた。この修正は統計的なものでした。実際の生産履歴 (通常の日: 理由の長さの中央値 107 ～ 197 文字、悪い日: 中央値 48) に対して調整し、中央値が 80 を下回った実行にはシステム崩壊としてフラグを立てました。形式の確認は簡単です。深さをチェックするには、自分の調子の良かった日から測定したベースラインが必要です。
そのファミリーからのもう 1 つの水平防御: 1 回の実行で 3 人のユーザーに対して同じ HIGH チェックが発生した場合は、すぐに管理者に連絡します。ユーザーごとの再試行→フォールバック チェーンは、個々の障害を細かく処理し、全体的な障害をひどく処理します。これらは、一度に 1 ユーザーずつ、静かに全員を劣化させます。
ニュースレターは 75 日、評議会は 34 日、1 日 2 通送信されます。重要なのは、「9 つのモデルが 1 つのモデルよりも賢い」ということではありません。ポイントは、決定論的なコードによって方向が決定され、LLM が構造的に安全な領域内でのみ動作するアーキテクチャです。

さらに、すべての出力はそれを信頼しないコードによって検証され、すべてのインシデントは永続的な名前付きチェックになります。一か八かの LLM システムの信頼性はモデルの特性ではありません。それは建築資産です。
次の投稿: 31 個すべての監査チェックとその失敗モードの分類 - 独自のインシデントを引き起こしたチェックの完全な分析を含む。
LLM 評議会 + 裁判官:一個天天在生產環境寄信的多モデル仲裁系統(75 天實錄)
MarketDaily は、この正確なシステムで毎日のダイジェストを朝と夕方に送信します。現在は期間限定で無料であり、有料プランが復活した場合でも、早期購読者は永久に無料アクセスを維持できます。
訂正 (2026-08-02): この投稿では「31 チェック」と書かれていますが、実際の数は 30 です。1 つの「チェック」は docstring の使用例にのみ存在し、起動することはできません。また、独自の検証 grep がそれをカウントしました。監査システム自体がカウントに失敗することは、まさに次の記事で詳しく説明する種類の失敗です。
この記事は情報提供のみを目的としており、投資アドバイスではありません。投資にはリスクが伴います。自分の状況を評価してください。最終更新日: 2026-08-02

## Original Extract

I run MarketDaily, a daily financial email digest. Subscribers pick their holdings (US + Taiwan stocks), and twice a day the system generates a personalized HTML report for each of them and sends it at a fixed time. First commit was 2026-05-19; as of writing that

I run a 9-model LLM council with a judge to write a daily financial newsletter. Here's what actually breaks. | MarketDaily
MarketDaily ←
All articles
MARKETDAILY · ENGINEERING
I run a 9-model LLM council with a judge to write a daily financial newsletter. Here's what actually breaks.
I run MarketDaily, a daily financial email digest. Subscribers pick their holdings (US + Taiwan stocks), and twice a day the system generates a personalized HTML report for each of them and sends it at a fixed time. First commit was 2026-05-19; as of writing that's 75 days and 1,800+ commits of production history, currently serving 21 subscribers. Small scale, but the failure tolerance is basically zero: this is financial content landing in someone's inbox before market open. A hallucinated price target isn't a bug report, it's a trust funeral.
This post is about the reliability layer that makes it survivable: a multi-model council, a judge, and a 31-check deterministic audit gate. Including the three incidents where it failed.
The problem with "model + fallback"
The naive architecture is: call your main model, fall back to another on error. That handles the failure mode that almost never matters. The failure mode that actually hurts is silent degradation : under quota pressure the model returns something that looks like a report — right format, right sections, no exception raised — but the analysis inside is mush. No error code fires. Your user finds out before you do.
So the design principle became: never trust a single model's output, and never trust any model's self-assessment. Three separated layers:
market data ──► structure prior (price vs MA20/MA50 — direction is
decided by deterministic code, not by any LLM)
│
▼
COUNCIL: 9 configured free-tier seats across 7 providers
(Gemini x2, Groq, local GPU via Ollama, Cloudflare
Workers AI x2, OpenRouter, Cerebras, OpenAI —
seats with no API key auto-skip; ~5 independent
voices live on a typical day)
│ each seat: {lean, conviction, thesis, key_risk}
│ quorum: >= 2 seats or no council verdict
▼
JUDGE: synthesizes consensus AND disagreement
(free chain: Gemini lite → Groq → if both die,
take the highest-conviction seat verbatim)
│
▼
report generation (separate LLM chain)
│
▼
AUDIT: 31 deterministic checks on the final HTML
│
HIGH fail ├──► wait 60s, retry with a stronger model
│ │ still HIGH
│ ▼
│ deterministic fallback (no LLM, no price
│ targets — a bad email never leaves the building)
▼
send at the fixed time, every day, no exceptions
Council: disagreement is a feature, not noise
Every stock gets one council round per run (cached across users). Each seat gets identical real market data plus a hard structural constraint: direction is locked by code (price vs MA20 vs MA50). A seat that says "short" on a bullish structure gets demoted to neutral by a plain if statement. No LLM in this system has the authority to flip a direction. The council only operates on the layer where models are actually useful: thesis, counter-risk, and lean within the structurally allowed range.
Then we measure dissent: 0 if all seats agree, 2 if the council contains both long and short, 1 otherwise. Dissent propagates all the way to the final card — high-dissent stocks get forced-conservative wording downstream, and displayed confidence gets pushed down. Confidence is separately capped at 75% by a historical calibration table; the audit treats any number above that as a broken defense, not an opinion.
Seat selection is quota engineering, not benchmark chasing. Full disclosure: the council landed on day 42 (2026-06-30) — before that, the system was exactly the naive model+fallback architecture this post opened by criticizing. The seats deliberately span independent vendors' free tiers, because the thing that kills you at 5am is correlated quota exhaustion, not model quality. Some real constraints that shaped the roster:
Groq's free tier is limited per-minute (8,000 TPM) and per-model per-day (200,000 TPD). The council seat uses a different model than the report-generation chain, because one shared bucket meant the council ate 197,364 of the 200,000 daily tokens and 45 report calls fell through to a 144-second-per-call backup model. That evening's digest was 1h35 late.
One seat (a local 14B model on my own GPU) exists purely because it has zero quota and zero network dependency. It's the only survivor when cloud DNS blips or every vendor's quota dies at once.
One model rewrote a price from the prompt — 385.25 became 385.00. Fabricated precision. It's permanently banned from anything that touches prices, but it's allowed in the council, because council output is JSON opinions with no numbers in them. Blast-radius design beats model trust.
Seats have circuit breakers: quota-dead / missing key / 3 consecutive failures disables the seat for the round. HTTP 402 (billing wall) kills it on the first strike — retrying a payment error is pure waste.
Judge: synthesize, with an exit
The judge sees all seat opinions as JSON and is instructed to synthesize consensus and disagreement, not to parrot the loudest seat. The judge itself has a fallback chain, and if the whole chain dies, the system takes the highest-conviction seat's thesis verbatim. If fewer than 2 seats respond, there's no council verdict at all and the stock goes down the plain single-model path. The entire council is fail-safe: any failure returns partial results and never blocks the send. "Never miss a send" and "never send garbage" are two different guarantees, enforced by two different mechanisms.
Audit: 31 checks, each one paid for in blood
audit_digest() runs 31 named deterministic checks against the final HTML. Every single one maps to a real way a user was (or would have been) angry: tense discipline (at 7am you cannot write "Taiwan stocks rose today" — the market opens at 9), holdings coverage (every stock the user selected must have an action card), fabrication detection (placeholder XXX tokens, fake URLs, earnings estimates that don't exist in the data feed), prompt-instruction leakage (the LLM copying its own instructions into the user-facing output), truncation detection, undefined CSS classes.
HIGH failures trigger: wait a full 60 seconds (free tiers are per-minute limited; retrying after 5s just hits the same 429 window and lands on a weaker model), regenerate with a stronger model forced to the front, and if it still fails, ship the deterministic fallback — plain code-assembled report, deliberately containing no price levels.
Three incidents that built the system
The 429 incident. Gemini's daily quota ran out; the retry-with-backoff logic burned ~100 seconds per call across the whole run. The digest went out 5 hours late. Fix: two consecutive 429s on one model = quota is dead for the day, fuse the whole model for the run. Retry logic that's correct for transient errors is actively harmful for quota errors — you have to classify.
The CSS incident, in two acts. LLMs asked to "reuse the existing CSS classes" would freestyle near-miss names: news-title instead of news-headline . No stylesheet rule, whole section renders unstyled. Act one: we added an undefined_css_class HIGH check. Act two, one week later: the Monday-edition prompt lacked a verbatim skeleton, so every model invented classes, all 12 users hit the check, the retry hit it too (same prompt, same disease), and 9 users got the degraded fallback version. The defense became the incident. The real fix was a deterministic repair layer before the check: map known near-miss names back to real classes, strip unknown ones (no CSS rule exists, so stripping is a visual no-op). Two lessons: telling an LLM "use the existing classes" is not a spec — a verbatim skeleton is; and before adding any HIGH check, ask "what happens if everyone fails it at once."
The depth-collapse incident. All free quotas fused one morning, generation fell to the weakest models, and per-stock reasoning collapsed from 129–165 characters to 48–64. Format was perfect. Prices were present. Every check passed. A human caught it. The fix was a statistical one: calibrate against real production history (normal days: median reason length 107–197 chars; the bad day: median 48) and flag a run whose median drops below 80 as systemic collapse. Checking format is easy; checking depth requires a baseline measured from your own good days.
One more horizontal defense from that family: if the same HIGH check fires for 3 users in a single run, page the admin immediately. Per-user retry→fallback chains handle individual failures fine and systemic failures terribly — they degrade everyone quietly, one user at a time.
75 days of the newsletter, 34 days of the council, two sends a day. The point was never "9 models are smarter than one." The point is an architecture where direction is decided by deterministic code, LLMs operate only inside structurally safe boundaries, every output is verified by code that doesn't trust it, and every incident becomes a permanent named check. Reliability in high-stakes LLM systems is not a model property. It's an architecture property.
Next post: a taxonomy of all 31 audit checks and their failure modes — including the full anatomy of the check that caused its own incident.
LLM council + judge:一個天天在生產環境寄信的多模型仲裁系統(75 天實錄)
MarketDaily sends its daily digest with this exact system, morning and evening — currently free for a limited time, and early-bird subscribers keep free access permanently if paid plans ever return
Correction (2026-08-02): this post says "31 checks" — the real number is 30. One "check" exists only in a docstring usage example and can never fire, and our own verification grep counted it anyway. An audit system failing to count itself is exactly the kind of failure the next post dissects.
This article is for informational purposes only and is not investment advice. Investing involves risk; assess your own situation. Last updated: 2026-08-02
