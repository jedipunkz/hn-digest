---
source: "https://veritrooper.com/"
hn_url: "https://news.ycombinator.com/item?id=48416123"
title: "Show HN: Audit any AI/data pairing with Veritrooper"
article_title: "VERITROOPER — Deploy Accuracy Anywhere"
author: "brian8620"
captured_at: "2026-06-05T18:45:20Z"
capture_tool: "hn-digest"
hn_id: 48416123
score: 1
comments: 0
posted_at: "2026-06-05T18:08:09Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Audit any AI/data pairing with Veritrooper

- HN: [48416123](https://news.ycombinator.com/item?id=48416123)
- Source: [veritrooper.com](https://veritrooper.com/)
- Score: 1
- Comments: 0
- Posted: 2026-06-05T18:08:09Z

## Translation

タイトル: HN を表示: Veritrooper を使用して AI/データのペアを監査する
記事のタイトル: VERITROOPER — どこにでも精度を導入
説明: VERITROOPER は、お客様自身のデータに対して LLM を監査します。幻覚を検出し、障害を特定し、監査対応レポートと EU AI 法適合証拠記録を返します。

記事本文:
書き込まれたデータに対する LLM の精度が大幅に向上します。 VERITROOPER は、確信を持って間違った答えを見つけ出し、チームに正確な場所と理由を示し、修正を手渡します。争点となったすべての評決はベンダー間で検証されています。
誠実さ、誠実さ、透明性 — 設計によるもの。
結果は、その背後にあるプロセスと同じくらい価値があります。監査人、バイヤー、および自社のエンジニアにとって、それを生み出すすべてのステップは防御可能であるように構築されています。独立したベンダー間の検証、私たちに不利な丸め込みスコア、内蔵の幻覚トラップ、完全な改ざん明示サインオフで表示されるすべての失敗。ここでの誠実さは主張ではなく、メカニズムです。
これを一度も測定して証拠と呼んだわけではありません。 VERITROOPER は、米国の税法、OSHA 職場の安全規制、FDA の医薬品ラベル表示、および SEC 10-K 財務届出など、4 つの無関係な規制世界にわたってエンドツーエンドで実行され、それぞれ同じ 1,000 の質問監査、同じ 7 つのモデル (主力フロンティアまでのゲーム GPU の 7B)、同じクロスベンダー検証が行われています。 1 つのドメインが幸運である可能性があります。 4 人が同じように行動するのはパターンです。そして、数値はモデルのものです。VERITROOPER には独自のスコアはなく、モデルの床と天井を継承します。フロンティア モデルは 3 つすべてにおいてほぼ完璧に仕上がっています。モデルが弱いほど、監査の精度が回復します。 OSHA での Llama 3.1 70B の正直な 77.84 は、外れ値ではなく証拠です。実際の監査では、モデルが 1 を獲得したときに低い数値を返すことができなければなりません。
ベースライン = バニラ RAG 取得 (BM25 トップ-5) を使用したモデル。実際のデプロイメントで使用される方法です。監査済み = 同じモデルの同じ回答、正しいソース証拠に照らして再測定され、争点となったすべての評決が異なるベンダーのモデルによって検証されます。 Δ = モデルがテーブル上に残す精度 —

そのあらゆる点が、チームが解決できる特定の質問と根本原因にまで遡ります。すべての数値はタイムスタンプ付きのログから再現可能です。
再現性は、測定と推測の間の境界線です。そこで、同じ 238 の質問、同じモデル、2 台のサーバーを並べて、まったく同じ監査を 10 回連続で実行しました。ここにすべての実行があり、手つかずです。
コア チェックは、モデルの文言を明確な合格/不合格に変える決定論的な段階です。10 回すべてまったく同じ数値が返されました。監査済みの完全なスコアはその上に乗り、ほとんど動きません。最後のわずかな変化は正直です。当社の監査人自体が AI モデルであり、毎回完璧な AI は存在しません。それがこの製品が存在する理由のすべてです。このアーキテクチャは、独自の AI パーツさえも同じ正確な答えを何度も繰り返し保持します。
生データを入力し、詳細なアフターアクションを出力します。 VERITROOPER の安全パラシュートは、LLM がデータ上で失敗する場所を正確に特定します。わかりやすい言葉で言えば、エンジニアに渡すことができます。ステージをホバリングします。
1
2
3
4
5
6
7
ステージ1
データを選択してください。
税法、安全規制、ルールブック、財務書類など、書き留められたものはすべて VERITROOPER に指示します。 PDF、Word、HTML、CSV、JSON、プレーン テキスト、さらにはライブ データベース (SQLite、SQL ダンプ、DBF) を取り込み、自動的にチャンク化して解析します。 (テキストレイヤードキュメント - テキストを読み取りますが、スキャンされた画像は OCR されません。)
ローカルでもクラウドでも、7B からフロンティアまで、あらゆるベンダー。クロード、GPT、ジェミニ、ラマ、クウェン — 持っているものをすべて組み合わせてください。ベースラインでは、バニラ RAG 取得 (運用スタイル) でモデルが実行されるため、監査では、必要最低限​​の LLM ではなく、実際のデプロイメントが実際に何を行うかを測定します。次に、監査では、ベンダー間検証を使用して、正しいソース証拠に照らして同じモデルを再測定します。つまり、Δ は回復可能な精度ギャップであり、正確に失われた箇所です。あ

うぬぼれ、高揚をもたらすものではありません。
VERITROOPER は、計算、条件付き、精度、相互参照、例外、因果関係、および意図的に答えられない「トラップ」質問に及ぶ、グラウンド トゥルースをロックインして、データから質問セットを生成します。計算検証、証拠の根拠付け、および製造保護はフロントエンド モジュールによって処理されます。すべての質問は、LLM が確認する前に監査されます。
ターゲット LLM は、ベースラインのみ、およびベースラインと VERITROOPER の検証レイヤーの 2 つの方法で実行します。同じ質問で直接対決し、同じスコアを得ました。
100% 正しくない答えは医師に送られます。医師は診断専門家のチームであり、それぞれが異なる種類の失敗や質問の種類に合わせて調整されています。一般的な推論、拒否、曖昧さなど。
各専門家は、自分が処理した障害に関する構造化された調査結果 (評決、証拠、カテゴリー、および校正メモ) を返します。すべてのエントリはタイムスタンプ付きのログから再現可能です。
レポーターは医師の所見を取り入れ、何が失敗したか、なぜ失敗したか、どのカテゴリ、何を修正すべきかなど、平易な英語のアフターアクションを書きます。推奨事項は、モデル、データセット、証拠痕跡などの失敗データから得られます。一般的な AI のアドバイスではありません。エンジニアに渡すと、すぐに対応できる有用なデータが得られます。
VERITROOPER は、あらゆる文書に対して LLM を測定可能なほど正確にし、その方法を正確に証明します。税法、安全規制、医薬品のラベル表示、財務申告、ゲームのルールなど、書き留めたものすべてをデータセットとして渡すと、検証済みのグラウンドトゥルースを含む質問セットが生成され、2 つの方法でモデルが実行されます。1 回目は現実的な検索 (ベースライン) で、もう 1 回目は正しい情報源の証拠を前に置いた状態 (監査) です。それらの間のギャップは、モデルがテーブル上に残す精度です。障害はルーティングされる

専門の診断モジュールに送信され、争点となったすべての判定は、それを上書きできる独立したクロスベンダー検証者によって確認されます。テスト対象のモデルがそれ自身の答えについて最終的な決定権を有することはありません。実行は、モデルの何が間違っていたのか、その理由、そしてそれを回復するために正確に何を修正すべきかについての平易な英語のレポートで終わります。
規制された資料では、単純な検索ではできないことを行います。質問が相互参照されたセクションに依存している場合、パイプラインはその参照されたセクションのテキストを証拠に引き出します。ベースラインでは実行できないマルチホップ解決です。財務申告の場合、ラベル/値/期間テーブルを解析し、内蔵の財務計算機を使用してソース数値とすべての計算をチェックするため、捏造された数値や間違った年の値が検出されます。
LLM が自信を持って間違った答えを生成するとき、それを捕捉するために構築された障害モードは幻覚です。モデルは幻覚を起こしてもクラッシュしません。フラグも警告もエラーコードもありません。彼らは、真実ではないことについて確信を持っているように聞こえるだけです。これは、答えが実際に重要な設定において、裸の LLM の展開を妨げる原因で​​す。 VERITROOPER は、あらゆるドメイン、あらゆるモデルでそれを捕捉します。
逆に得られるものは、グラウンド トゥルースに対して調整された精度スコア、質問ごとに分類されたすべての障害のリスト (パターンとクラスターが特定された)、カテゴリごとの精度の内訳、障害回復率 (監査で回復したモデルのベースライン障害の割合)、および LLM が示した特定のギャップを埋めるための具体的なエンジニアリングの推奨事項です。すべての判定はタイムスタンプ付きのログから再現可能であり、ブラックボックスはありません。
ヨーロッパへの展開? 1 つのトグルで、同じ実行に EU AI 法適合証拠を追加します: 宣言された精度と堅牢性のテスト (第 15 条)、ドロップイン付録 IV 技術文書

メンテーション記録、反復的な精度ドリフトモニタリング（第 72 条）、日付が記載された署名済みの人的レビュー監査証跡（第 14 条）、およびカテゴリごとのギャップ診断（第 10 条）。 VERITROOPER は、適合ファイルが依存する証拠を作成します。プロバイダーの適合性評価に代わるものではなく、適合性を付与するものでもありません。
電話層 7B から主力フロンティアまで、7 つの異なる LLM にわたって 3 つの規制ドメイン データ セットをテストしました。クロスベンダー検証による正しい証拠に照らして測定すると、7B を実行する 2,000 ドルの RTX 4090 から主力のフロンティア API に至るまで、すべてのモデルが、IRS 税法に関する同じ 1,000 の質問で 95.9 ～ 100% の範囲内に収まりました。それがオチです。これはフロンティアだけの贅沢ではありません。ラップトップの 7B であっても、アースが解決されればフロンティア モデルのすぐ近くに来ます。そのため、監査では、データ上のギャップはほとんど回復可能であることが示され、ギャップを埋めるために修正すべき障害が正確に特定されます。
私だけ。ブライアン B. 平均的な男です。アメリカ陸空軍の退役軍人。
この 4 月以前には、VERITROOPER は名前としても、製品としても、コードとしても存在していませんでした。私は頭の中を占めるために D&D ゲームに取り組んでいました (障害のある獣医で、時間がありました)。私は Unity をゼロから独学で学び、LLM をダンジョン マスターとして接続して、ゲームがシーンを実行し、判決を下せるようにしました。厄介な問題が 1 つありました…
AIは私の嘘を見破ることができませんでした。何でも餌を与えることができ、自信を持って走りました。キャラクターが持っていないゴールドを費やしたり、存在しない呪文を唱えたり、キャラクターが実際には実行できないアクションを実行したりできました。私のくだらないことを呼んでくれるDM AIが欲しかった。そのようなものは存在しないことが判明しました。この問題を解決しようとしたのがVERITROOPERです。
結局のところ、DM が知らなかったために回答をでっち上げたのは、幻覚と呼ばれるより広範な問題でした。 DM、リー

今日の LLM ベースのナレーションは、事実に基づいて答えようとし、それができない場合は、正しいか間違っているか、または完全にでっち上げであるかに関係なく、自信を持って答えます。一言で言えばAI幻覚です。当時は全く分かりませんでした。
自分の目標を見つめていくうちに、それは可能だと思い始めました。このことについて仲の良い友人に話していると、彼は「もし RIFTS (別のゲーム システム) を実行できるようにしてくれたら本当に感心するよ」と率直に言いました。私は「できると思います」と言いました。あることが別のことにつながり、私が作成したこのシステムは、どの LLM を使用しても、どのようなデータセットでも機能するという認識が定着し始めました。
D&D SRD は、これを構築したときに実行した最初のデータ セットでした。見た目よりも難しく、アーキテクチャ全体をブートストラップしました。ゲーム メカニクスに関するルール法則化を確実に処理できる検証レイヤーは、税法、軍法、財務申告、空軍の訓練資料に一般化されます。それが、結果セクションのデータ セットで証明されています。
VERITROOPER は、D&D ゲームを作ろうとしている退屈な男が意図せず発見したものです。自宅の地下室で PC を持ち、問題を抱え、時間を持て余していた男が、開始から展開準備完了の開発までを短期間で行いました。
買収者、パイロットパートナー、技術レビュー担当者向け。
リクエストに応じてライブウォークスルーを行います。生のログ、完全なデータセット、および特許パッケージは、NDA に基づいて連絡すれば入手可能です。
技術証明パケットをダウンロード (PDF) →
誠実さ、誠実さ、透明性 — 設計によるもの
測定の信頼性が製品です。これらは数値を正確に保つための具体的なメカニズムであり、ボルトで取り付けられるものではなく、エンジンに組み込まれています。基本的なルール: 結果が真でないことが証明された場合、その結果は偽として扱われます。
明確な答えは、モデルの意見に関係なく、決定論的なグラウンドトゥルースチェックによって決定されます。争奪版

dict は独立した検証者に送られます。テスト対象のモデル自体がフロンティア モデルである場合、エンジンはその検証者をコード内の別のベンダーに強制します。
独立した検証者は、争点となったすべての評決をレビューし、それを無効にすることができます。
信頼度が低い場合、サード ベンダーが引き分けを破ります。ボーダーラインの判定は 1 人の採点者に依存しません。
評決の提案とその確認が同じベンダーによって行われることはありません。
数字はフロアであり、決して膨らむことはありません。
空または無回答は不合格として採点されます。(「答えられない」トラップ質問を含む) 黙って正解としてカウントされることはありません。
「ほぼ正しい」というのは依然として間違っています。ヘッジ、イエス/ノーの間違った極性、またはコミットの拒否はすべて失敗としてカウントされます。部分クレジットはありません。
質問自体の参照回答が不正な形式である場合、その回答はベースラインと監査済みの分母の両方から均等に削除され、カウントが出力されます。
すべてのレポートには、単なるわかりやすい数値ではなく、生の精度と検証者が調整した精度が並べて表示されます。
ベンダー間の完全な監査では、調整されたデルタは公表されている数値よりも広くなりました。報告された数値は上限ではなく、防御可能な下限です。
3 幻覚と自信過剰の検出
自信を持ってもっともらしく聞こえる失敗は、直接テストされます。
答えのない質問が意図的に混ぜ込まれています。サポート不可能なことを自信を持って答えるモデルは捕らえられ、カウントされます。

[切り捨てられた]

## Original Extract

VERITROOPER audits any LLM against your own data — detecting hallucinations, pinpointing failures, and returning audit-ready reports plus EU AI Act conformity-evidence records.

Make any LLM measurably more accurate on any written data. VERITROOPER catches the confident wrong answers, shows your team exactly where and why, and hands them the fixes — every contested verdict cross-vendor verified.
Integrity, Honesty & Transparency — by Design.
A result is only worth as much as the process behind it. Every step that produces one is built to be defensible — to your auditors, your buyers, and your own engineers: independent cross-vendor verification, scoring that rounds against us, built-in hallucination traps, every failure shown in full, tamper-evident sign-off. Integrity here isn’t a claim — it’s the mechanism.
We didn’t measure this once and call it proof. VERITROOPER has been run end-to-end across four unrelated regulated worlds — U.S. tax code, OSHA workplace-safety regulation, FDA drug labeling, and SEC 10-K financial filings — each the same 1,000-question audit, the same seven models (a 7B on a gaming GPU up to flagship frontier), the same cross-vendor verification. One domain could be luck; four behaving the same way is a pattern. And the numbers are your model’s — VERITROOPER carries no score of its own, it inherits the model’s floor and ceiling. Frontier models land near-perfect on all three; the weaker the model, the more accuracy the audit recovers. Llama 3.1 70B’s honest 77.84 on OSHA is the proof, not an outlier — a real audit has to be able to return a low number when the model earns one.
Baseline = the model with vanilla-RAG retrieval (BM25 top-5), the way real deployments serve. Audited = the same model’s same answers, re-measured against the correct source evidence, every contested verdict verified by a different vendor’s model. Δ = the accuracy your model is leaving on the table — every point of it traced to a specific question and root cause your team can fix. Every figure reproducible from timestamped logs.
Reproducibility is the line between a measurement and a guess. So we ran the exact same audit ten times in a row — same 238 questions, same model, two servers side by side. Here is every run, untouched.
Core check is the deterministic stage that turns the model’s wording into a clean pass/fail — it came back to the exact same number all ten times. The full audited score rides on top of it and barely moves. That last sliver of wiggle is honest: our auditors are themselves AI models, and no AI is perfect every single time — that’s the whole reason this product exists. The architecture holds even its own AI parts to the same accurate answer, again and again.
Raw data in, detailed after-action out. VERITROOPER's safety parachute pinpoints where your LLM fails on your data — in plain English you can hand to an engineer. Hover a stage.
1
2
3
4
5
6
7
Stage 1
Choose your data.
Point VERITROOPER at anything written down — tax code, safety regs, rulebooks, financial filings. It ingests PDF, Word, HTML, CSV, JSON, plain text, and even live databases (SQLite, SQL dumps, DBF), then chunks and parses automatically. (Text-layer documents — it reads the text, it doesn’t OCR scanned images.)
Local or cloud, 7B to frontier, any vendor. Claude, GPT, Gemini, Llama, Qwen — plug in what you've got. The baseline runs the model with vanilla-RAG retrieval — production-style, so the audit measures what your real deployment actually does, not a stripped-down LLM. The audit then re-measures the same model against the correct source evidence with cross-vendor verification — so the Δ is the recoverable accuracy gap, and exactly where it's lost. Audit, not serving uplift.
VERITROOPER generates the question set from your data, with ground truth locked in — spanning calculation, conditional, precision, cross-reference, exception, cause-effect, and deliberate unanswerable “trap” questions. Calc-verification, evidence grounding, and fabrication protection are handled by the front-end modules — every question is audited before the LLM ever sees it.
We run the target LLM two ways: baseline alone, and baseline plus VERITROOPER's verification layer. Head-to-head on identical questions, scored identically.
Any answer not 100% correct gets routed to the doctors — a team of diagnostic specialists, each tuned for a different type of failure or question type. General reasoning, refusals, ambiguity, and so on.
Each specialist returns structured findings on the failures they handled — verdict, evidence, category, and calibration notes. Every entry is reproducible from the timestamped log.
The Reporter takes the Doctor findings and writes a plain-English after-action — what failed, why, what category, what to fix. Recommendations come from your failure data: your model, your dataset, your evidence trail. Not generic AI advice. Hand it to an engineer — usable data they can act on immediately.
VERITROOPER makes any LLM measurably more accurate on any written material — and proves exactly how. You hand it a data set — tax code, safety regulations, drug labeling, financial filings, gaming rules, anything written down — and it generates a question set with verified ground truth, then runs the model two ways: once on realistic retrieval (the baseline) and once with the correct source evidence in front of it (the audit). The gap between them is the accuracy your model is leaving on the table. Failures are routed to specialist diagnostic modules, and every contested verdict is confirmed by an independent cross-vendor verifier that can override it — the model under test never gets the final say on its own answers. The run ends in a plain-English report on what the model got wrong, why, and exactly what to fix to recover it.
On regulated material it does what naive retrieval can’t: when a question depends on a cross-referenced section, the pipeline pulls that referenced section’s text into the evidence — multi-hop resolution the baseline can’t do. For financial filings it parses label/value/period tables and checks every calculation against the source numbers with a built-in financial calculator, so a fabricated figure or a wrong-year value gets caught.
The failure mode it's built to catch is hallucination — when an LLM confidently produces a wrong answer. Models don't crash when they hallucinate; there's no flag, no warning, no error code. They just sound certain about something that isn't true. That's what breaks naked LLM deployment in any setting where the answer actually matters. VERITROOPER catches it across any domain, with any model.
What you get out the other side: an adjusted accuracy score against ground truth, a per-question categorized list of every failure (with patterns and clusters identified), a per-category accuracy breakdown, a failure-recovery rate (the share of the model’s baseline failures the audit recovered), and concrete engineering recommendations that would close the specific gaps the LLM showed. Every verdict is reproducible from timestamped logs — no black box.
Deploying into Europe? One toggle adds the EU AI Act conformity evidence to the same run: declared accuracy and robustness testing (Article 15), a drop-in Annex IV technical-documentation record, recurring accuracy-drift monitoring (Article 72), a dated, signed human-review audit trail (Article 14), and a per-category gap diagnostic (Article 10). VERITROOPER produces the evidence a conformity file relies on — it does not replace the provider's conformity assessment or confer compliance.
Three regulated-domain data sets tested across seven different LLMs, phone-tier 7B to flagship frontier. Measured against the correct evidence with cross-vendor verification, every model — from a $2,000 RTX 4090 running a 7B up to flagship frontier APIs — lands in the 95.9–100% band on the same 1,000 IRS Tax Code questions. That's the punchline: this isn't a frontier-only luxury. Even a 7B on a laptop comes within shouting distance of a frontier model when grounding is solved — so the audit shows the gap on your data is mostly recoverable, and pinpoints exactly which failures to fix to close it.
Just me. Brian B. Average guy. Veteran of the US Army and Air Force.
Before this April, VERITROOPER didn't exist — not as a name, not as a product, not as code. I was working on a D&D game to occupy my mind (disabled vet, had time). I taught myself Unity from zero, wired up an LLM as the dungeon master so the game could run scenes and adjudicate rulings. I had one nagging problem…
The AI couldn't catch me lying. I could feed it anything and it would confidently run with it. I could spend gold the character didn't have, cast spells that didn't exist, and perform actions the character couldn't actually perform. I wanted a DM AI that called my crap. It turned out no such thing existed. Trying to solve this problem led to VERITROOPER.
As it turns out, the DM not knowing, and therefore making answers up, was a far broader problem called hallucination. The DM, like any LLM-based narration today, would try to answer factually, and if it couldn't it would answer confidently, whether right, wrong, or just completely made up. That's AI hallucination in a sentence. I had no idea at the time.
As I zeroed in on my goal, I began to think it was possible. I was talking about it to a good friend of mine, who offhandedly said "I'd be really impressed if you could make it run RIFTS" (another gaming system). I said "I think I can." One thing led to another, and the realization that this system I had created worked on any data set, using any LLM began to set in.
D&D SRD was the first data set I ever ran when I built this — harder than it looks, and it bootstrapped the whole architecture. Any verification layer that can reliably handle rules-lawyering on game mechanics generalizes to tax code, military law, financial filings, and Air Force training material — and that's what the data sets in the Results section prove.
VERITROOPER is the unintended discovery of a bored man trying to make a D&D game. Inception to deployment-ready development happened in a short amount of time from a guy in his basement with a PC, a problem, and time on his hands.
For acquirers, pilot partners, and technical reviewers.
Live walkthroughs by request. Raw logs, full dataset, and patent package available on contact under NDA.
Download Technical Proof Packet (PDF) →
Integrity, Honesty & Transparency — by Design
The credibility of the measurement is the product. These are the concrete mechanisms that keep the numbers honest — built into the engine, not bolted on. Foundational rule: if a result isn’t provably true, it’s treated as false.
Clear-cut answers are settled by deterministic ground-truth checks, independent of any model’s opinion. Contested verdicts go to an independent verifier — and when the model under test is itself a frontier model, the engine forces that verifier to a different vendor in code .
The independent verifier reviews every contested verdict and can override it.
On low confidence, a third vendor breaks the tie — no borderline call rests on a single grader.
Proposing a verdict and confirming it are never done by the same vendor.
The number is a floor, never inflated.
An empty or non-answer is scored as a failure — never silently counted correct (including on “unanswerable” trap questions).
“Almost right” is still wrong: hedging, wrong yes/no polarity, or refusing to commit all count as failures. No partial credit.
When a question’s own reference answer is malformed, it’s removed from both the baseline and audited denominators equally — and the count is printed.
Every report shows raw and verifier-adjusted accuracy side by side, never just the friendlier figure.
In full cross-vendor auditing, the adjusted deltas have run wider than the figures published — the reported number is the defensible floor, not the ceiling.
3 Hallucination & overconfidence detection
The confident, plausible-sounding failures are tested for directly.
Unanswerable questions are deliberately mixed in; a model that confidently answers the unsupportable is caught and counted a

[truncated]
