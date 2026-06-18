---
source: "https://blog.attacks.ai/we-got-glasswing-at-home"
hn_url: "https://news.ycombinator.com/item?id=48585285"
title: "We Got Anthropic's Glasswing at Home (Who Needs Mythos 5 or Fable 5?)"
article_title: "We Got Glasswing at Home, and It Found Real Bugs — attacks.ai"
author: "Seventeen18"
captured_at: "2026-06-18T16:16:07Z"
capture_tool: "hn-digest"
hn_id: 48585285
score: 2
comments: 0
posted_at: "2026-06-18T13:49:00Z"
tags:
  - hacker-news
  - translated
---

# We Got Anthropic's Glasswing at Home (Who Needs Mythos 5 or Fable 5?)

- HN: [48585285](https://news.ycombinator.com/item?id=48585285)
- Source: [blog.attacks.ai](https://blog.attacks.ai/we-got-glasswing-at-home)
- Score: 2
- Comments: 0
- Posted: 2026-06-18T13:49:00Z

## Translation

タイトル: Anthropic の Glasswing を家に入手しました (Mythos 5 や Fable 5 が必要なのは誰ですか?)
記事のタイトル: 自宅に Glasswing があったが、本物のバグが見つかった — Attacks.ai
説明: 所有するハードウェアに Anthropic の Glasswing が必要だったので、それをビルドしました。Lucent、RTX 3090 上でローカル 27B Qwen を実行する、ほぼ無料のバグハンターです。ビルド、テレメトリ、および hermes-agent で表面化した 2 つの実際のバグを次に示します。

記事本文:
">
自宅に Glasswing があったのですが、本物のバグが見つかりました — Attacks.ai
攻撃。 AI・研究
すべての投稿 · 開示 · メモ · 研究
について
RSS
← すべての投稿
Disclosure · May 27, 2026 · 26 min read
家にGlasswingがあったのですが、本物の虫が見つかりました
私は所有するハードウェアに Anthropic の Glasswing が欲しかったので、それを構築しました。Lucent は、1 台の RTX 3090 上でローカル 27B Qwen を実行する、ほぼ無料のバグハンターです。ビルド、テレメトリ、そして hermes-agent で表面化した 2 つの実際のバグを次に示します。
Anthropic には Glasswing がいます。Glasswing は、コードベースを独自に読み取り、実際の脆弱性を返す自律的なセキュリティ研究者です。私も欲しかったです。 API 呼び出しによってレンタルされたのではなく、私の部屋にあるマシン、私が制御するモデル上で実行され、電気料金で一晩研削し続けることができます。
これが、私がそのマシンを構築した方法、そのマシン自体のテレメトリが示す内容、そして、初めて深刻なターゲットに向けたときに完成したバージョンで発見された 2 つの実際のバグです。
ビルドの話の前の正直な見出し: hermes-agent に対して、静的パスにより 1,342 の候補サイトにフラグが立てられ、ローカル掃討によりそれらのリードが 126 に絞り込まれ、敵対的監査によりそれらが 15 に減り、さらに公開する価値のある 2 つのバグに絞り込まれました。チャットの参加者なら誰でも回答できる承認プロンプトです。もう 1 つも対象範囲内にあり、Nous に報告されていますが、その修正はまだ到着していないため、出荷されるまで詳細を保留しています。エンゲージメントの最も鋭い結果は、レビュー担当エージェントから得られました。途中で、私がすでに他の 3 つのエクスプロイトを「デモンストレーション」した後、ベンダーが 6 週間前に置き換えたセキュリティ ポリシーに照らしてそれらを評価していたことが発覚し、それによって私の勝利のほとんどが削除されました。
最初の試みはかろうじて自動化されていました。私は実際のターゲットに対して大きなクラウド モデルである Opus 4.7 を手動で操作し、バグを見つけるように依頼しました。私は

5 つの発見とトップ 10 のリストという自信に満ちた山ができましたが、そのほとんどすべてが誤検知と行き止まりでした。最も説得力のあるものは、PDF 抽出ルーチンでのパス トラバーサルで、.pth ファイルをドロップし、次回の Python 起動時にコード実行に変わる可能性があり、紙の上ではクリーンなチェーンでした。最初に確認すべきだった 1 つのこと、つまり上流ライブラリが書き込み前にファイル名を正規化するかどうかを確認した瞬間に、問題は崩れてしまいました。それはそうです。走査がディスクに到達することはありません。
これは、チェックなしで実行されたモデルの通常の故障モードです。信頼性は高く、ほとんどが間違っています。より大きなモデルでは問題は解決されません。ボトルネックは、悪いリードを追いつくのに十分な速さで破棄することだったので、チャット ウィンドウでの作業をやめ、パイプラインを構築しました。
から借用できるオープンソースの出発点があり、私も簡単に借用しました。それは私が必要としていたことをしませんでした、そして、それが本当の問題を発見するまでに、私はそれのほとんどを書き直しました。私はそれをルーセントと呼んでいます。
ルーセントは会話ではありません。これはステージングされたパイプラインであり、各ステージで異なるモデルを自由に実行でき、ターゲットのソースはロックダウンされた Docker サンドボックス内に読み取り専用でマウントされます。
ランク。すべてのソース ファイルに脆弱性が隠蔽される可能性をスコア付けして、高価なステージが重要なところに予算を費やすようにします。大きな木では、これが、実行が終了するか終了しないかの違いになります。
ハント。ファイル並列エージェントの階層化されたプールは、ランク付けされたファイルを読み取り、それぞれが file:line と記述されたメカニズムに固定されたリードを記録します。これは最大ボリュームのステージであり、ローカル モデルで実行されます。
確認する。敵対的パスは、ソースに対して各リードを再読み取りし、それを反証しようとします。間違ったメカニズム、到達不能、ハーネスのアーティファクトです。これを生き延びないと何も進まない。
悪用します。生存者はトリアージとバリアントループを悪用します

実用的な概念実証を作成しようとします。
証拠のはしごの最上部に達するまでは、何も発見とは呼ばれません: 疑惑 → static_corroboration → crash_reproductioned → root_cause_explained → Exploit_demonstrated 。一番上の行は、実行され、ライブ インスタンスに対する動作を示すスクリプトを意味します。それ以下はすべてリードであり、リードは安価です。検証ステージは最も多くの作業を行ったステージであり、以下のほとんどはその理由について説明します。
リグ: 1 つの GPU、ローカル 27B、Lucebox
大容量の読み取りは 1 台の RTX 3090 で実行されます。ランカーとハンターは、Lucebox が提供するローカルのオープン Qwen3.6-27B を駆動します。これは、投機的デコードを使用します。事前に複数のトークンをドラフトし、それらをバッチで検証するため、ステップごとに 1 つではなく複数のトークンがコミットされます。このカードでは、4 ビットの 27B に対して、コードのようなテキストの生成が約 3.4 倍高速になり (1 秒あたり約 130 トークン、単純な自己回帰デコードの場合は 38 トークン)、最も通常のファイルではピークで 4 倍を超えます。それが、大きな木を指すのに十分な速さでソースを読み取る 27B とそうでないものの違いです。参考のために、これはデータセンター カード上の小型モデルに関するレポート 4 ～ 5x から借用した投機的デコーディングの論文です。コンシューマー 3090 の 4 ビットで 27B で 3.4 倍が所有できるバージョンです。
ローカルとは読書だけを意味するものではありません。ランキング、ハンティング、ほとんどのリードを殺す検証はすべてその 27B で実行されます。 1,342 の静的ヒットのうち 126 を除くすべてをスローするモデルがローカル モデルです。フロンティアモデルである Opus 4.7 が最上位に位置します。Opus 4.7 は実行を調整し、少数の生存者の最終的な詳細な監査を行い、証拠を構築してそれぞれを論破します。各ステージでは異なるモデルを自由に実行できるため、最上層はフロアではなく選択になります。それはその地位を獲得します

狭い理由が 1 つあります。発見を破ろうとする査読者は、それを構築したものと少なくとも同じくらい鋭くなければなりません。そうしないと、すべてを揺るがすことになります。その下にあるものはすべてローカルのままです。873k ライン ツリーの読み取りとトリアージはスケールを伴う作業であり、それをトークンで測定するため、この種の研究は高価になります。
モンティを引き抜き、エルメスに配線する
パイプラインには、次に何を読み取るか、そしてそれについてどのように推論するかを決定するドライバーである Monty と呼ばれるエージェント ループが付属していました。即興で必要な部分が硬く、複数のファイルにまたがる推論が手探りだったので、私はそれを破り、NousResearch の Hermes Agent を導入し、その後、私が望むように狩りを進めるためにしばらくかけて調整しました。
これは立ち止まる価値があります。ハンターの頭脳として私が選んだエージェントは、後に完成したハンターを解放することになる同じ NousResearch プロジェクト、hermes-agent です。それは計画されていなかった。私が Hermes を選んだのは、それが手元にある最高のエージェントだったからであり、それを自分のものにするときにその内部を知り、その後になって初めてその結果がプロジェクト自体に向けられたからです。コードベースを内部からよく知っていることが、探索がうまくいった理由の一部です。
これらはどれも長い間うまくいきませんでした。タイムラインの正直なバージョンは、パイプラインが何か間違ったことをしている数週間です。ランカーが興味深いファイルを埋め、ハント ステージで存在しない file:line 参照を発明し、ベリファイアが強制終了すべきものを手を振りながら、Lucebox ゲートウェイがシングルフライトであり、負荷がかかると停止して再起動が必要になるため、実行全体がウェッジされます。大きなツリーのコールド スイープには何時間もかかるため、あらゆる間違った仮定を見つけるのに 1 日の大半が費やされます。
一つずつ修正していきました。私は検証者に、ディスク上のファイルを含む自身の入力を信頼しないように教えましたが、それはより重要でした。

予想していたよりも大きかったです（詳細は以下で説明します）。実行中にルーセントチェックポイントを作成したため、最初からやり直すのではなく失速が再開されます。危険なコードが実行される前に、フレームワークまたはプロトコル ライブラリがいつバグを無力化したかを認識するトリアージ レイヤーを追加しました。これにより、それらは検出結果として報告されなくなります。どこかで、自信に満ちたゴミを作成することから、読む価値のある短いリストを作成するようになりました。
最初の本当のターゲット: hermes-agent
hermes-agent は、NousResearch のオープンソース パーソナル エージェントです。多くのチャネル (Telegram、電子メール、Slack、Matrix、Feishu など) からメッセージを取り込み、シェル コマンドを実行し、ファイルを書き込み、コードを実行し、それらに応答して「スキル」をインストールすることができるデーモンです。攻撃対象領域は、信頼できない入力パスの大規模なセットと特権アクションの大規模なセットとの間のギャップです。このツリーは、2,903 のソース ファイルにわたる約 873,000 行で構成されており、私が Lucent で作った、実際のコードベースのような厄介なコードベースです。私はソースを読んで決して修正せず、発見したことを Nous Research に報告しました。
モデルが実行される前に、静的解析だけで 1,342 の候補サイトにフラグが付けられました。そのうち 72 が重大と呼ばれ、105 が高、1,162 が中、3 が低で、パターン マッチャーでは何も検証できないため検証済みは 0 でした。この数値はパイプラインへの入力であり、出力ではありません。ランカーが最も疑わしいファイルを一番上に押し上げ、ハンターがそこから下を読み進め、ファネルが狭まりました。
1,342
126
15
静的ヒット
所見候補
トリアージされたリード
9
死んだ
誤検知 · 撤回 · 範囲外
4
争える
硬化・潜在
2
範囲内
#13 ゲートウェイの承認 - フェールオープン · 高
#14 は修正版がリリースされるまで保留されます
完全なトリアージファネル: 1,342 件の静的解析ヒットが 126 件の候補結果を通じて除外され、15 件が 2 件の実際のバグにつながりました (ここでは 13 件が見つかりました。2 件目は修正がリリースされるまで保留されました)。
どうやってフン

走った
126 から 15 への絞り込みは敵対的監査であり、ローカル スイープの逆であるため、そのコストを示す価値があります。つまり、制限されたフロンティア モデルであり、実際の支出がどこにあるのかということです。約 20 人の専門エージェントがこの取り組みを実行しました。6 人の偵察監査人が並行して配置され、それぞれが 1 つのサーフェス (認証、データ フロー、出力) を所有しました。リードを構築計画に変換した 1 人のアーキテクト。 6 つのビルダーは、それぞれが独自の git ワークツリーに 1 つの実行可能な概念実証を作成しました。そして、調査結果をバラバラにすることだけが唯一の仕事だった 6 人の査読者パス。一番上にある 1 つの長期実行オーケストレーター。これは、私がクリーンなテレメトリーを持っている 13 人のエージェント全体でおよそ 150 万のトークンに相当し、偵察とオーケストレーションをカウントすると 200 万を超えます。個々のエージェントは 20 ～ 96 回のツール呼び出しを行いました。 API サーバーのバグだけを作成したビルダーは、満足する前に 87 を作成しました。
きれいな午後ではありません。 2 人のビルダーが実行中にハーネスエラーで死亡しましたが、ディスクから回収されて再起動されました。アカウントは途中で使用制限に達し、自動再開が設定され、制限がリセットされたときに回復しました。エージェント コンピューティングの測定値は約 1 時間 45 分ですが、ビルダーとレビューアーが並行して実行されるため、これは実時間ではなく合計です。このような混乱を放置すると、記事が製品デモのようなものになってしまいます。
最初のスイープでは 126 の候補結果が生成されましたが、それらのほぼすべてが間違っていました。説得力のあるものは問題であり、一日を無駄にするのに十分であるように見えるリードでした。
調査結果は、この書き込み全体で単一のシーケンスで 01 ～ 15 を実行するため、以下の行は連続していません。この表は行き止まりを収集し、実際の範囲外の動作がその後に続き、完全なリストは最後のスコアカードにあります。
これらの撤回のうち 2 つは、最後のチェックをスキップした場合に出荷されるバグそのものであるため、詳しく説明する価値があります。 0を見つける

2 は、最初のバージョンの .pth チェーンで、パイプラインによって再び検出され、同じ方法で強制終了されます。 marker-pdf は、脆弱な os.path.join が画像ファイル名を認識する前に画像ファイル名を正規化するため、関数レベルのプリミティブは実数であり、エンドツーエンドのパスは実数ではありません。 03 を見つけることは、WhatsApp ブリッジの許可リストのクリーン パス トラバーサルのように見えました。反証の深さは 4 段階あります。Baileys ライブラリは、ブリッジのハンドラーが起動する前に cleanMessage を呼び出し、それが jidNormalizedUser を呼び出します。これにより、@ が欠落している文字列 (すべてのトラバーサル ペイロード) に対して空の文字列が返され、ブリッジは正規のチャット ID にフォールバックします。トラバーサル文字列を配信できません。これは関数レベルでは本物ですが、実際のイングレスでは無効になります。
いくつかの実際の動作は範囲外でした。
06、browser_cdp SSRF (実際、条件付き): 本物のサーバー側リクエスト フォージェリですが、オペレーターがデフォルト以外の DevTools URL を構成した場合に限ります。
08、read_file 編集バイパス (実際、範囲外): code_file=True フラグは位置ベースのシークレット マスキングをスキップするため、.env の不透明な値は編集されずに返されます。リード側リークを再現しました。次に、ライブ モデルにシークレットを実行するよう依頼すると、モデルは拒否され、リクエストにデータ漏洩としてフラグが付けられました。

[切り捨てられた]

## Original Extract

I wanted Anthropic's Glasswing on hardware I own, so I built it: Lucent, a near-free bug-hunter running a local 27B Qwen on one RTX 3090. Here is the build, the telemetry, and the two real bugs it surfaced in hermes-agent.

">
We Got Glasswing at Home, and It Found Real Bugs — attacks.ai
attacks . ai / Research
All posts · Disclosures · Notes · Research
About
RSS
← All posts
Disclosure · May 27, 2026 · 26 min read
We Got Glasswing at Home, and It Found Real Bugs
I wanted Anthropic's Glasswing on hardware I own, so I built it: Lucent, a near-free bug-hunter running a local 27B Qwen on one RTX 3090. Here is the build, the telemetry, and the two real bugs it surfaced in hermes-agent.
Anthropic has Glasswing: an autonomous security researcher that reads a codebase on its own and comes back with real vulnerabilities. I wanted one too. Not rented by the API call, but running on a machine in the room with me, on models I control and can leave grinding overnight for the price of electricity.
This is how I built that machine, what its own telemetry says it did, and the two real bugs the finished version found the first time I pointed it at a serious target.
The honest headline before the build story: against hermes-agent , a static pass flagged 1,342 candidate sites, the local sweep narrowed those to 126 leads, and an adversarial audit cut them to 15 and then to the two bugs worth disclosing: an approval prompt that anyone in the chat can answer. The other is also in scope and reported to Nous, but its fix is still landing, so I am holding its details until it ships. The sharpest result of the engagement came from a reviewer agent. Partway through, after I had already "demonstrated" three other exploits, it caught that I had been grading them against a security policy the vendor had replaced six weeks earlier, which deleted most of my wins.
The first attempts were barely automated. I drove a big cloud model, Opus 4.7, by hand against a real target and asked it to find bugs. It produced a confident pile: five findings and a top-ten list, almost all of it false positives and dead ends. The most convincing one was a path traversal in a PDF-extraction routine that could drop a .pth file and turn into code execution at the next Python startup, a clean chain on paper. It fell apart the moment I checked the one thing I should have checked first, whether the upstream library normalizes the filename before it writes. It does. The traversal never reaches disk.
This is the normal failure mode for a model run without checks: high confidence, mostly wrong. A bigger model does not fix it. The bottleneck is discarding bad leads fast enough to keep up, so I stopped working in a chat window and built a pipeline.
There was an open-source starting point to borrow from , and I did, briefly. It did not do what I needed, and by the time the thing was finding real issues I had rewritten most of it. I call it Lucent .
Lucent is not a conversation. It is a staged pipeline, each stage free to run a different model, with the target's source read-only-mounted inside a locked-down Docker sandbox:
Rank. Score every source file for how likely it is to hide a vulnerability, so the expensive stages spend their budget where it matters. On a large tree this is the difference between a run that finishes and one that does not.
Hunt. A tiered pool of file-parallel agents reads the ranked files and records leads, each pinned to a file:line and a described mechanism. This is the highest-volume stage, and it runs on the local model.
Verify. An adversarial pass re-reads each lead against the source and tries to disprove it: wrong mechanism, not reachable, an artifact of the harness. Nothing advances until it survives this.
Exploit. Survivors go to exploit triage and a variant loop that tries to produce a working proof of concept.
Nothing is called a finding until it reaches the top of an evidence ladder: suspicion → static_corroboration → crash_reproduced → root_cause_explained → exploit_demonstrated . The top rung means a script that runs and shows the behavior against a live instance. Everything below that is a lead, and leads are cheap. The verify stage is the one that did the most work, and most of what follows is about why.
The rig: one GPU, a local 27B, Lucebox
The high-volume reading runs on a single RTX 3090. The ranker and the hunters drive a local open Qwen3.6-27B served by Lucebox , which uses speculative decoding: it drafts several tokens ahead and verifies them in a batch, so it commits multiple tokens per step instead of one. On this card, against the 27B at 4-bit, that works out to roughly 3.4× faster generation on code-like text (about 130 tokens per second, against 38 for plain autoregressive decoding), peaking past 4× on the most regular files. That is the difference between a 27B that reads source fast enough to point at a large tree and one that does not. For reference, the speculative-decoding paper this borrows from reports 4–5× on smaller models on a datacenter card; 3.4× on a 27B at 4-bit on a consumer 3090 is the version you can own.
Local does not mean only reading. The ranking, the hunting, and the verify that kills most of the leads all run on that 27B; the model that throws out all but 126 of the 1,342 static hits is the local one. A frontier model, Opus 4.7, sits on top: it orchestrates the run and does the final deep audit of the few survivors, building the proofs and arguing each one down. Each stage is free to run a different model, so that top layer is a choice rather than a floor. It earns its place for one narrow reason: a reviewer that tries to break a finding has to be at least as sharp as whatever built it, or it waves everything through. Everything beneath it stays local, because reading and triaging an 873k-line tree is the work that scales, and metering that by the token is what makes this kind of research expensive.
Ripping out Monty, wiring in Hermes
The pipeline came with an agent loop called Monty , the driver that decides what to read next and how to reason about it. It was rigid where I needed it to improvise, and it fumbled reasoning that spanned several files, so I tore it out and dropped in NousResearch's Hermes Agent , then spent a while tuning it to drive the hunt the way I wanted.
This is worth pausing on. The agent I picked to be the brain of the hunter is the same NousResearch project, hermes-agent, that I would later turn the finished hunter loose on. That was not planned. I chose Hermes because it was the best agent I had on hand, got to know its internals while making it mine, and only afterward pointed the result at the project itself. Knowing a codebase that well, from the inside, is part of why the hunt went where it did.
None of this worked for a long time. The honest version of the timeline is weeks of the pipeline doing something wrong: the ranker burying the interesting files, the hunt stage inventing file:line references that did not exist, the verifier waving through things it should have killed, the whole run wedging because the Lucebox gateway is single-flight and would stall under load and need a restart. A cold sweep of a large tree takes hours, so every bad assumption cost the better part of a day to find.
I fixed them one at a time. I taught the verifier to distrust its own inputs, including files on disk, which mattered more than I expected (more on that below). I made Lucent checkpoint mid-run, so a stall resumes instead of starting over. I added a triage layer that recognizes when a framework or protocol library neutralizes a bug before the dangerous code can run, so those stop being reported as findings at all. Somewhere in there it went from producing confident garbage to producing a short list worth reading.
First real target: hermes-agent
hermes-agent is NousResearch's open-source personal agent: a daemon that ingests messages from many channels (Telegram, email, Slack, Matrix, Feishu, and others) and is allowed to run shell commands, write files, execute code, and install "skills" in response to them. The attack surface is the gap between a large set of untrusted-input paths and a large set of privileged actions. The tree is about 873k lines across 2,903 source files, the kind of messy, real codebase I built Lucent to chew on. I read the source and never modified it, and reported what I found to Nous Research.
Static analysis alone flagged 1,342 candidate sites before any model ran: 72 it called critical, 105 high, 1,162 medium, 3 low, and zero verified, because a pattern matcher cannot verify anything. That number is the input to the pipeline, not an output. The ranker pushed the most suspicious files to the top, the hunters read down from there, and the funnel narrowed:
1,342
126
15
static hits
candidate findings
triaged leads
9
dead
false positive · retracted · out of scope
4
contestable
hardening · latent
2
in scope
#13 gateway approval — fail-open · High
#14 held until its fix ships
The full triage funnel: 1,342 static-analysis hits shed down through 126 candidate findings and 15 leads to the 2 real bugs (finding 13 here; the second held until its fix ships).
How the hunt ran
The narrowing from 126 to 15 is the adversarial audit, and it is worth showing its cost because it is the opposite of the local sweep: bounded, frontier-model, and where the real spend is. About 20 specialized agents ran across the engagement: six recon auditors in parallel, each owning one surface (auth, data flow, output); one architect that turned leads into a build plan; six builders that each wrote one runnable proof of concept in its own git worktree; and six reviewer passes whose only job was to tear findings apart. One long-running orchestrator on top. That is roughly 1.5 million tokens across the 13 agents I have clean telemetry for, past 2 million counting recon and orchestration. Individual agents made 20 to 96 tool calls; the builder for the API-server bug alone made 87 before it was satisfied.
It is not a clean afternoon. Two builders died to a harness error mid-run and were recovered from disk and relaunched. The account hit a usage limit partway through, armed an auto-resume, and picked up when the limit reset. About 1h45m of measured agent-compute, but that is a sum, not a wall clock, because the builders and reviewers run in parallel. Leaving that mess out is how a writeup ends up reading like a product demo.
The first sweep produced 126 candidate findings, and almost all of them were wrong. The convincing ones were the problem, the leads that looked solid enough to waste a day on.
The findings run 01–15 in a single sequence across this writeup, so the rows below are not contiguous: this table collects the dead ends, the real-but-out-of-scope behaviors come after it, and the complete list is in the scorecard at the end.
Two of those retractions are worth the detail, because they are exactly the bugs you ship if you skip the last check. Finding 02 is the .pth chain from the first version, found again by the pipeline and killed the same way: marker-pdf normalizes the image filename before the vulnerable os.path.join ever sees it, so the function-level primitive is real and the end-to-end path is not. Finding 03 looked like a clean path traversal in the WhatsApp bridge's allow-list. The disproof is four steps deep: the Baileys library calls cleanMessage before the bridge's handler fires, that calls jidNormalizedUser , which for any string lacking an @ (every traversal payload) returns the empty string, and the bridge then falls back to the legitimate chat ID. The traversal string cannot be delivered. It is real at the function level and dead through the real ingress.
A few were real behaviors that fell outside scope:
06, browser_cdp SSRF (real, conditional): a genuine server-side request forgery, but only when the operator configures a non-default DevTools URL.
08, read_file redaction bypass (real, out of scope): a code_file=True flag skips the position-based secret masking, so an opaque value in a .env comes back unredacted. I reproduced the read-side leak. When I then asked a live model to carry the secret out, it declined and flagged the request as data exfiltration

[truncated]
