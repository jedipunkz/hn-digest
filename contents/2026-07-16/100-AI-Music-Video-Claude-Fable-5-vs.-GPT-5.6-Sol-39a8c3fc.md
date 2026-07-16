---
source: "https://www.tryai.dev/blog/ai-music-video-arena-claude-vs-gpt-5.6"
hn_url: "https://news.ycombinator.com/item?id=48939524"
title: "$100 AI Music Video: Claude Fable 5 vs. GPT-5.6 Sol"
article_title: "$100 AI Music Video: Claude Fable 5 vs. GPT-5.6 Sol · TryAI"
author: "hershyb_"
captured_at: "2026-07-16T20:53:35Z"
capture_tool: "hn-digest"
hn_id: 48939524
score: 5
comments: 4
posted_at: "2026-07-16T20:03:23Z"
tags:
  - hacker-news
  - translated
---

# $100 AI Music Video: Claude Fable 5 vs. GPT-5.6 Sol

- HN: [48939524](https://news.ycombinator.com/item?id=48939524)
- Source: [www.tryai.dev](https://www.tryai.dev/blog/ai-music-video-arena-claude-vs-gpt-5.6)
- Score: 5
- Comments: 4
- Posted: 2026-07-16T20:03:23Z

## Translation

タイトル: $100 AI ミュージック ビデオ: Claude Fable 5 vs. GPT-5.6 Sol
記事タイトル: $100 AI ミュージック ビデオ: Claude Fable 5 vs. GPT-5.6 Sol · TryAI
説明: Claude Fable 5 と GPT-5.6 Sol に同じ曲、予算、Web 検索、ローカル ffmpeg を与え、それぞれが自律的にミュージック ビデオを監督できるようにしました。

記事本文:
$100 AI ミュージック ビデオ: Claude Fable 5 vs. GPT-5.6 Sol · TryAI TryAI モデル比較 ギャラリー ブログ サインイン 無料で試してみる すべての投稿ビデオ比較エージェント $100 AI ミュージック ビデオ: Claude Fable 5 vs. GPT-5.6 Sol
Claude Fable 5 と GPT-5.6 Sol に同じ曲、予算、Web 検索、ローカル ffmpeg を与え、それぞれが自律的にミュージック ビデオを監督できるようにしました。
8 min read 私たちは 1 つの仕事で小さなエージェント ハーネスを構築しました。モデルに曲、ハードドルの予算、ツールのセットを渡し、その後邪魔をせずに完全なミュージック ビデオを自力で制作させます。モデルは、どのビデオ モデルが存在するかを調査し、クリップを生成し、独自のフッテージを視聴し、ffmpeg で編集し、最終カットを組み立てます。
前回のビルドオフの読者の何人かが、実際にモデル間でツールの使用がどのように異なるのかを確認したいと述べたので、フロンティアレベルのモデルに、何を調査するか、何を生成するか、どのように編集するかを各モデルが独自に決定する、無制限で長期的なタスクを与えました。すべてのツール呼び出しをログに記録するので、それぞれの呼び出しがどのように機能したかを正確に確認できます (以下の完全なトランスクリプト)。
Claude Fable 5 と GPT-5.6 Sol の 2 つのモデルをそれぞれ 2 つの予算 (25 ドルと 100 ドル) で合計 4 回実行しました。すべての実行には、同じ曲 (ブルーノ・マーズとマーク・ロンソンの「アップタウン・ファンク」)、短いテキストの説明、およびタイムスタンプ付きの歌詞トランスクリプトが提供されます。
各モデルは、6 つのツールを使用して自律的なツール呼び出しループを実行しました。
plan : 考えるためのツール (コストもアクションも不要)。
web_search : 生成モデルとその API を調査し、ミュージック ビデオに関する情報を取得します (必要な場合)。
get_budget : 残りの予算を確認します。
generate_image とgenerate_video : 予算を消費する唯一のツール。モデルは、任意の FAL モデルまたはレプリケート モデルを選択し、独自のパラメーターを渡すことができます。
run_command : ffmpeg/ffprobe が利用可能なローカル シェル。オーディオの分析、c の切り取り、連結に使用されます。

唇を重ねて、最後のビデオを多重化します。
予算がゼロになると、有料生成は拒否されますが、モデルは編集を続けることができます。すべてのモデル メッセージ、ツール呼び出し、チャージ、およびエラーがログに記録されました。ハーネス全体は github.com/hershalb/music-video-arena でオープンソースであるため、自分で実行できます。
以下の各クリップは、モデルの最終的な自己組み立て出力 .mp4 で、オリジナルの曲が多重化されたフルレングスです。
4 つのランはすべて自然に終了し (ステップや時間制限に達するものはありませんでした)、4 つすべてがオリジナルの曲が多重化された有効なフルレングスのビデオを作成しました。
「生成支出」は従量制の FAL コストであり、予算の上限となります。どちらのモデルも 25 ドルでほぼ使い切ってしまいました。 100 ドルとして、36.57 ドル（ソル）と 48.60 ドル（フェイブル）を費やしたため、より多くの予算がより多くの映像につながりました。これには、以下に追加するモデル自体の実行コストは含まれません。
独自のツールを選択する必要があり、モデルは多様化しました。 4 回の実行のうち 3 回は純粋なテキストからビデオへの変換でした。 25 ドルの GPT-5.6 Sol のみが、画像からビデオへのパイプラインを使用していました (最初に静止画を生成し、次にそれらをアニメーション化します)。 GPT-5.6 Sol は 100 ドルで、1 回の実行で 3 つの異なるビデオ モデルを混合しました。
価格は、特に明記されていない限り、出力ビデオの 1 秒あたりの FAL のリスト料金です。 Hailuo 2.3 Standard はビデオごとに価格設定されており (6 秒のクリップあたり約 0.28 ドル)、Seedance 1.0 Pro はトークン価格です (5 秒の 1080p クリップあたり約 0.62 ドル、上記は実効 1 秒あたりのレートとして示されています)。実行ごとに生成される個別のクリップは 46 ～ 80 の範囲でした。
各実行でツール呼び出しがどのように費やされたか (失敗した生成呼び出しを含む試行回数がカウントされます)。
クロード・ファブル 5 · $100 GPT-5.6 ソル · $100
各実行の完全なトランスクリプト、すべての計画、ツール呼び出し、コマンドは次のとおりです: Fable 5 · $25 、 Sol · $25 、 Sol · $100 、 Fable 5 · $100 。
「失敗した呼び出し」とは、エラーを返した生成リクエストです (ほとんどの場合、システムへの一時的なネットワーク障害が発生します)。

プロバイダ）。これらは請求されませんでしたが、モデルは再試行するためにステップを費やしました。
予算は世代 (FAL) の支出のみを測定します。 Claude Fable 5 の LLM トークン コスト ($10 / 1M 入出力あたり $50) と GPT-5.6 Sol ($5 / $30) を加算すると、各実行の合計コストが得られます。
Claude Fable 5 の場合、トークンだけで 16.99 ドルから 25.05 ドルが実行され、これは各実行総額の約 30 ～ 40% でした。 GPT-5.6 Sol のトークンコストは、同様のトークン量にもかかわらず、3 ～ 4 ドル近くにとどまりました。
4 つの実行すべてで同じ入力: 曲、短いテキストの説明、タイムスタンプ付きの歌詞トランスクリプト。各モデルは FAL 上で独自の世代モデルを選択し、独自の ffmpeg 編集を行いました。
実時間には、モデル自体の再試行とプロバイダー キューでの待機時間が含まれます。
生成費用は、モデルごとの価格表からのベストエフォート型の見積もりです。
アリーナはオープンソースです: github.com/hershalb/music-video-arena 。自分の曲と予算に合わせて、比較したいモデルを交換して、それらが何を構築するかを確認してください。問題や PR は大歓迎です。セットアップに関するフィードバックもお待ちしています。
ミュージック ビデオはどれも素晴らしいものではありませんでしたが、モデルたちがどのようにしてそこに到達するのかを見るのは非常に興味深く、フロンティア レベルのモデルには明らかにギャップがまだ存在することがわかりました。いくつかの点に注意してください:
キャラクターとストーリーの一貫性は、4 人全員にとって苦労しました。登場人物がショット間を行き来し、どのビデオも最初から最後まで一貫したストーリーラインを保持していません。
モデルたちは歌詞を文字通りに受け止めています。 「ドラゴンを引退させたいな、おい」と言うと、画面上に実際のドラゴンが表示されます。数ショットまでは面白いですが、しばらくすると少し奇妙になりました。
テンポ合わせが弱い。カットはビートに沿って着地しますが (すべて ffmpeg ビート検出を実行しました)、クリップ内のモーション、ダンス、カメラの移動が曲のテンポと一致することはほとんどないため、少しずれていると感じることがよくあります。行の例「gotta k

は自分自身です、私はとてもきれいです」は、主人公がキスする動作が遅すぎることを示しています。
GPT-5.6 Sol (25 ドル) は最も独創的なエディタでした。これは、テキストとアニメーション静止画をビデオ効果でオーバーレイするもので、他の実行では試みられなかった手法です。残りのほとんどは、生成されたクリップをつなぎ合わせただけです。 GPT 5.6 Sol $100 も、Fable のように 1 つのビデオ モデルに固執するのではなく、複数のビデオ モデルを試しました。
実際に編集を繰り返した人は誰もいませんでした。クリップが存在すると、モデルは連結されて多重化されますが、再カットしたりエフェクトを追加したりするために戻ってくることはほとんどなく、クリップが適切であることを確認するために自分のクリップを真剣に調査する人は誰もいませんでした。 GPT-5.6 Sol の 100 ドル版では、いくつかの本当に低品質の AI クリップが同梱されていましたが、Claude Fable 5 はたまたま、より一貫性のある出力を持つモデルを選択しました。これの一部はおそらくモデルの制限ですが、自己レビューの欠如は注目に値します。
どちらのモデルもReplicateには触れていません。 FAL キーとレプリケート キーの両方が使用可能でしたが、4 回の実行すべてで FAL のみが使用されました。
クロード・フェイブル5はより高価な選択でした。 GPT-5.6 Sol よりも速く終了するにもかかわらず、実行あたりのコストは高くなります (そして全体で最も高いのは $73.65)。主観的には、Fable $100 のビデオの方がわずかに好みでしたが、どれも私たちを驚かせるものではありませんでした。
100ドルは予算が多すぎたかもしれません。どちらのモデルも上限近くで過ごすことを望まず、歩数も控えめに保ちました。その余裕があれば、たとえば、一貫したキャラクター イメージを事前に生成し、そこからアニメーション化することもできましたが、どちらもそうすることを選択しませんでした。
モデルがより賢くなるにつれて、より主観的/様式的なタスクを改善できるかどうかがわかりますが、現時点ではまだ改善の余地がたくさんあります。
ここで説明したすべてのモデルは、TryAI で 1 つのアカウントで利用でき、サブスクリプションなしで従量課金制で利用できます。
古い投稿 GPT-5.6、Grok 4.5、Claude、および Muse Spark が同じ 4 つのアプリを構築 TryAI すべてのフロンティア AI モデル

1か所。チャット、画像、ビデオ、音楽 - 使用した分だけお支払いください。
© 2026 TryAI.無断転載を禁じます。
すべての AI モデルを試してみたい人向けに構築されています。

## Original Extract

We gave Claude Fable 5 and GPT-5.6 Sol the same song, a budget, web search, and local ffmpeg, then let each autonomously direct a music video.

$100 AI Music Video: Claude Fable 5 vs. GPT-5.6 Sol · TryAI TryAI Models Compare Gallery Blog Sign in Try free All posts video comparison agents $100 AI Music Video: Claude Fable 5 vs. GPT-5.6 Sol
We gave Claude Fable 5 and GPT-5.6 Sol the same song, a budget, web search, and local ffmpeg, then let each autonomously direct a music video.
8 min read We built a small agentic harness with one job: hand a model a song, a hard dollar budget, and a set of tools, then get out of the way and let it produce a full music video on its own. The model researches which video models exist, generates clips, watches its own footage, edits with ffmpeg, and assembles a final cut.
A few readers of our last build-off said they wanted to see how tool use actually varies between models, so we gave frontier-level models an open-ended, long-horizon task where each model decides on its own what to research, what to generate, and how to edit. We log every tool call, so you can see exactly how each one worked (full transcripts below).
We ran two models, Claude Fable 5 and GPT-5.6 Sol , each at two budgets ($25 and $100), for four runs total. Every run got the same song (Bruno Mars and Mark Ronson's "Uptown Funk"), a short text description, and a time-stamped lyric transcript.
Each model ran an autonomous tool-calling loop with six tools:
plan : a tool for thinking (no cost, no action).
web_search : to research generation models and their APIs and fetch information about music videos (if needed).
get_budget : to check the remaining budget.
generate_image and generate_video : the only tools that spend budget. The model can pick any FAL or Replicate model and pass its own parameters.
run_command : a local shell with ffmpeg/ffprobe available, used to analyze audio, cut and concatenate clips, and mux the final video.
Once the budget hits zero, paid generation is refused, but the model can keep editing. Every model message, tool call, charge, and error was logged. The whole harness is open source at github.com/hershalb/music-video-arena , so you can run it yourself.
Each clip below is the model's final, self-assembled output.mp4, full length with the original song muxed in.
All four runs finished on their own (none hit a step or time limit) and all four produced a valid, full-length video with the original song muxed in.
"Generation spend" is the metered FAL cost, which is what the budget caps. At $25 both models nearly exhausted it. At $100 they spent $36.57 (Sol) and $48.60 (Fable), so more budget did translate into more footage. It does not include the cost of running the model itself, which we add below.
Left to choose their own tools, the models diverged. Three of the four runs went pure text-to-video. Only GPT-5.6 Sol at $25 used an image-to-video pipeline (generating stills first, then animating them). GPT-5.6 Sol at $100 mixed three different video models in a single run.
Prices are FAL's listed rates, shown per second of output video unless noted. Hailuo 2.3 Standard is priced per video (about $0.28 per 6s clip), and Seedance 1.0 Pro is token-priced (~$0.62 per 5s 1080p clip, shown above as its effective per-second rate). Distinct clips generated per run ranged from 46 to 80.
How each run spent its tool calls (this counts attempts, including failed generation calls).
Claude Fable 5 · $100 GPT-5.6 Sol · $100
Each run's full transcript, every plan, tool call, and command, is here: Fable 5 · $25 , Sol · $25 , Sol · $100 , Fable 5 · $100 .
"Failed calls" are generation requests that returned an error (mostly transient network failures to the provider). They were not charged, but the model spent steps retrying them.
The budget only meters generation (FAL) spend. Adding the LLM token cost for Claude Fable 5 ($10 / $50 per 1M input/output) and GPT-5.6 Sol ($5 / $30), gives the total cost of each run.
For Claude Fable 5, the tokens alone ran $16.99 to $25.05, about 30-40% of each run's total. GPT-5.6 Sol's token cost stayed near $3-4 despite similar token volume.
Same inputs for all four runs: song, a short text description, and a time-stamped lyric transcript. Each model chose its own generation models on FAL and did its own ffmpeg editing.
Wall-clock time includes the model's own retries and any waiting on provider queues.
Generation spend is a best-effort estimate from a per-model price table.
The arena is open source: github.com/hershalb/music-video-arena . Point it at your own song and budget, swap in whichever models you want to pit against each other, and see what they build. Issues and PRs welcome, we would love feedback on the setup.
None of the music videos were great, but watching how the models got there was pretty interesting and does show where gaps still clearly exist for frontier-level models. A few things notes:
Character and story consistency was a struggle for all four. Recurring characters drift between shots, and none of the videos hold a coherent storyline from start to finish.
The models take lyrics very literally. "Make a dragon wanna retire, man" gets you an actual dragon on screen. It's interesting for a few shots, but got a little weird after a while.
Tempo matching is weak. The cuts land on the beat (they all ran the ffmpeg beat detection), but the motion inside the clips, dancing, camera moves, rarely matches the song's tempo, so it often feels a little off. An example line "gotta kiss myself I'm so pretty", shows the main character making a kissing motion way too slowly.
GPT-5.6 Sol at $25 was the most inventive editor. It overlaid text and animated still images with video effects, techniques none of the other runs tried. The rest mostly just stitched generated clips together. GPT 5.6 Sol $100 also tried multiple video models instead of just sticking with one like Fable did.
Nobody really iterated on the edit. Once clips existed, the models concatenated and muxed, but rarely went back to re-cut or add effects, and none seriously probed their own clips to confirm they were any good. GPT-5.6 Sol's $100 run shipped some genuinely low-quality AI clips, while Claude Fable 5 happened to pick a model with more coherent output. Some of this is probably a model limitation, but the lack of self-review is notable.
Neither model touched Replicate. Both FAL and Replicate keys were available, but all four runs used FAL exclusively.
Claude Fable 5 was the pricier pick. It cost more per run (and the most overall, at $73.65) despite finishing faster than GPT-5.6 Sol. Subjectively, we slightly preferred the Fable $100 video, though none blew us away.
$100 was probably too much budget. Neither model wanted to spend near the cap, and both kept their step counts modest. With that headroom they could have, for example, generated consistent character images up front and animated from those, but neither chose to.
We'll see if models can improve on more subjective/stylistic tasks as they continue to get smarter, but for now there's still a lot of room for improvement.
Every model mentioned here is available on TryAI with one account, pay-as-you-go, no subscription.
Older post GPT-5.6, Grok 4.5, Claude, and Muse Spark build the same 4 apps TryAI Every frontier AI model in one place. Chat, image, video, and music — pay only for what you use.
© 2026 TryAI. All rights reserved.
Built for people who want to try every AI model.
