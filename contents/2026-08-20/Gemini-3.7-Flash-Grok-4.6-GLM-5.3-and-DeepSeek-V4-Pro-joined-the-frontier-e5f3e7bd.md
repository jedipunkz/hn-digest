---
source: "https://quesma.com/blog/baba-is-aug-2026/"
hn_url: "https://news.ycombinator.com/item?id=49377202"
title: "Gemini 3.7 Flash, Grok 4.6, GLM-5.3 and DeepSeek V4 Pro joined the frontier"
article_title: "Gemini 3.7 Flash, Grok 4.6, GLM-5.3 and DeepSeek V4 Pro joined the frontier - Quesma Blog"
image: "https://quesma.com/_astro/thumbnail.CU6fIcVH.png"
author: "stared"
captured_at: "2026-08-20T17:20:28Z"
capture_tool: "hn-digest"
hn_id: 49377202
score: 2
comments: 0
posted_at: "2026-08-20T16:57:19Z"
tags:
  - hacker-news
  - translated
---

# Gemini 3.7 Flash, Grok 4.6, GLM-5.3 and DeepSeek V4 Pro joined the frontier

- HN: [49377202](https://news.ycombinator.com/item?id=49377202)
- Source: [quesma.com](https://quesma.com/blog/baba-is-aug-2026/)
- Score: 2
- Comments: 0
- Posted: 2026-08-20T16:57:19Z

## Translation

タイトル: Gemini 3.7 Flash、Grok 4.6、GLM-5.3、DeepSeek V4 Pro がフロンティアに加わりました
記事タイトル: Gemini 3.7 Flash、Grok 4.6、GLM-5.3、DeepSeek V4 Pro がフロンティアに加わりました - Quesma Blog
説明: 2026 年 8 月の Baba Is Bench: Gemini 3.7 Flash、Grok 4.6、DeepSeek V4 Pro 0813 はそれぞれ、コストを 3 ～ 20 分の 1 に抑えながら、前世代を上回りました。無差別級の GLM-5.3 と Qwen3.8 の進歩は緩やかです。

記事本文:
Gemini 3.7 Flash、Grok 4.6、GLM-5.3、DeepSeek V4 Pro がフロンティアに加わりました - Quesma ブログ メイン コンテンツにスキップ 仕組み ベンチマーク ブログ 私たちについて トークン経済学 創設者と話す
仕組み ベンチマーク ブログ 私たちについて トークン経済学 創設者と話す Gemini 3.7 Flash、Grok 4.6、GLM-5.3、DeepSeek V4 Pro がフロンティアに加わりました
PNG をダウンロード この 8 月中旬は、LLM リリースに関しては 7 月と同じくらい暑いです: Grok 4.6 、 Gemini 3.7 Flash 、DeepSeek V4 Pro 0813、およびオープンウェイト モデル: Qwen3.8 Max 、 Qwen3.8 27B 、および GLM-5.3 。
インテリジェンス インデックスとインテリジェンス インデックス タスクあたりのコスト
人工分析から。次のようなベンチマークのスコアに基づいていることに注意してください。
Terminal-Bench v2.1、SciCode、人類最後の試験、GPQA ダイヤモンド - 必ずしも抽象的なような流動的なインテリジェンスではない
ARC-AGI-3 または Baba is You のパズル ゲーム。
これらのモデルは賢いので、パズル ゲームでモデルがどのように機能するかを確認するために Baba Is Bench を再実行することにしました。このゲームは人気があるにもかかわらず、ネタバレがないかチェックします。そして驚いたことに、SWE ベンチ Verified とは異なり、モデルが先の解決策を知っている兆候はありません。
Grok、Gemini Flash、DeepSeek では素晴らしい結果が得られます。バージョンのマイナーな変更に騙されないでください。バージョンの変更はすべて質的に異なる結果をもたらします。 GLM と Qwen の両方にとって、進歩はより緩やかでした。
前回と同様に、これを 2 ラウンドに分けて行います。まず、最初のステージである Intro をプレイするように求められ、次に少なくとも 7/8 レベルを通過した人には次のステージである Lake もプレイするように求められます。
分析したモデルを、以前の淡色表示のモデルと比較して表示します。
モデル 00 馬場さんは 01 どこに行きますか? 02 さて、これは何ですか？ 03 手の届かない 04 まだ手の届かない 05 火山 06 立ち入り禁止 07 芝生広場 pass@1 ↓ pass@3 ↓ ターン ↓ 出力トークン ↓ 総コスト ↑ Google Gemini 3.7 Flash 100% 100% 53 22k $

5.38 人間的クロード作品 5 100% 100% 7 23k $17.15 OpenAI GPT-5.5 100% 100% 8 15k $13.62 人間的クロード寓話 5 100% 100% 6 29k $40.98 OpenAI GPT-5.6 ソル 100% 100% 11 10k $11.91 Grok Grok 4.6 96% 100% 25 52k $14.01 Z.ai GLM-5.3 96% 100% 13 47k $6.20 キミ キミ K3 96% 100% 9 29k $12.56 人間クロード Opus 4.8 96% 100% 35 46,000 $63.14 Z.ai GLM-5.2 96% 100% 12 75,000 $6.20 Google Gemini 3.1 Pro 92% 100% 8 38,000 $12.41 Google Gemini 3.6 Flash 88% 100% 92 82,000 $124.31 OpenAI GPT-5.6 Terra 88% 100% 41 55k $34.28 DeepSeek DeepSeek V4 Pro 0813 75% 88% 18 130k $2.94 DeepSeek DeepSeek V4 フラッシュ 75% 100% 55 190k $2.82 Grok Grok 4.5 75% 88% 86 68k $50.66 DeepSeek DeepSeek V4 Flash 0731 75% 88% 22 156k $1.16 Anthropic Claude Sonnet 5 67% 88% 22 151k $42.21 OpenAI GPT-5.6 Luna 63% 88% 118 129k $59.32 Qwen Qwen3.8 Max 58% 75% 9 66,000 $10.75 Qwen Qwen3.6 27B 29% 38% 24 82,000 $5.59 DeepSeek DeepSeek V4 Pro 29% 38% 37 183,000 $11.49 Qwen Qwen3.7 最大 25% 38% 21 132,000 $16.98 Qwen Qwen3.8 27B 21% 38% 9 134k $9.07 パスタイム エージェントのステップ移動 解決済み タイムアウト 間違った答えは解決されていない
どちらの結果も上昇し、価格は大幅に下落したため、結果は驚くべきものとしか言いようがありません。
Gemini 3.7 Flash はすべてのトライアルを解決し、以前の Gemini 3.6 Flash よりも 20 倍以上安価になりました。
Grok 4.6 は 1 回の試行を除いてすべて解決しました (Grok 4.5 は 6 回の試行と 1 つのレベルを失敗しました) と同時に、3 倍以上安価でした。
DeepSeek V4 Pro 0813 は 3/8 しか解決しなかった以前のバージョンに対して 7/8 レベルを解決し、しかも 4 倍安価でした。
0% 25% 50% 75% 100% $1 $2 $5 $10 $20 $50 $100 完全なイントロ実行のコスト、対数スケール パス@1 Anthropic OpenAI Anthropic OpenAI Kim Anthropic Z.ai Google OpenAI DeepSeek Anthropic OpenAI Qwen Qwen Qwen Qwen Google Gemini 3.6 Flash Google Gemini 3.7 Flash Grok Grok 4.5 Grok Grok 4.6 ディープシーク ディープシーク

V4 プロ ディープシーク ディープシーク V4 プロ 0813 ディープシーク ディープシーク V4 フラッシュ 0731 Z.ai GLM-5.3
GLM および Qwen ファミリーでは、あまり成功しません。
GLM-5.3 では GLM-5.2 とまったく同じ結果が得られました。
Qwen3.8 Max は Qwen3.7 Max よりも優れていますが (はるかに多くの試行が解決されました)、すべてのレベルを完了するにはまだ遠いです。
Qwen3.8 27B の結果は Qwen3.6 27B よりも悪いですが、すべて実験ノイズの範囲内です。
それがうまくいかなかったことに驚きました。特に Qwen3.8 27B は人工分析で 52 のスコアを獲得しています。トランスクリプトで見たところによると、実際にボードを探索することなく、一度に解決しようとしました。これは、Qwen3.8 27B のデフォルトでは物事を非常に考えすぎるという Simon Willison の経験と一致します。 Baba Is You では明らかに考える必要がありますが、考えることとテストすることのバランスが重要です。多くのことを実行し、あまり考えていなかった以前の Grok および Gemini Flash モデルと比較対照してください。
$1 $2 $5 $10 $20 $50 $100 5 10 20 50 100 エージェント ターン (試行ごとの平均) * 1 レベルの未解決の総コスト ← より意図的であるほどトリガー ハッピーになる → ← 安価であるほど高価になる → Google Gemini 3.7 Flash Anthropic Claude Opus 5 OpenAI GPT-5.5 Anthropic Claude Fable 5 OpenAI GPT-5.6 Sol Grok Grok 4.6 Z.ai GLM-5.3 キミ キミ K3 Anthropic Claude Opus 4.8 Z.ai GLM-5.2 Google Gemini 3.1 Pro Google Gemini 3.6 Flash OpenAI GPT-5.6 Terra DeepSeek DeepSeek V4 Pro 0813* DeepSeek DeepSeek V4 Flash Grok Grok 4.5* DeepSeek DeepSeek V4 Flash 0731* Anthropicクロード ソネット 5* OpenAI GPT-5.6 ルナ*
エージェントのターン数と価格のグラフを見るのは興味深いです。 Grok と Gemini Flash は、トリガー ハッピー モデルからより穏やかなモデルに移行し、コストが大幅に削減されました。
8 つのイントロ レベルのうち少なくとも 7 に合格したモデルについては、次のステージである Lake で実行します。
コストがかかるため、レベルごとに 1 回のみ試行してください。

少なくとも、以前のモデルは高価でしたが、現在は大幅に改善されています。
クロード・ファブル 5 クロード・コード クロード・コード 01 アイシー・ウォーターズ
解決済み · 400万 · 277kトークン · $1.73 02ターン
解決済み · 1100万 · 110万トークン · $4.83 03 愛情
解決済み · 300万 · 229kトークン · $0.99 04 Pillar Yard
解決済み · 800万 · 777kトークン · $3.64 05 Brick Wall
解決済み · 2m · 206k トークン · $0.67 06 ロック
解決済み · 600万 · 795kトークン · $2.76 07 初心者の鍵屋
解決済み · 600万 · 695kトークン · $2.46 08ロックイン
解決済み · 500万 · 396kトークン · $1.70 09 変更なし
解決済み · 300万 · 213kトークン · $1.37 10 Two Doors
解決済み · 1600万 · 180万トークン · $5.96 11 Jelly Throne
解決済み · 1100万 · 57万トークン · $3.66 12 Crab Storage
解決済み · 1200万 · 150万トークン · $5.00 13 強盗
解決済み · 1000万 · 110万トークン · $4.52 b1 水没遺跡
解決済み · 1100万 · 676kトークン · $3.54 b2 Sunken Temple
停止 · 49m · 337k トークン · $10.41 14 2 時間 36 分 10.7M $53.25 Anthropic Claude Opus 5 Claude Code クロード コード 01 Icy Waters
解決済み · 300万 · 365kトークン · $0.60 02ターン
解決済み · 1600万 · 94万トークン · $2.33 03 愛情
解決済み · 700万 · 596kトークン · $0.83 04 Pillar Yard
解決済み · 1300万 · 948kトークン · $1.88 05 ブリックウォール
解決済み · 5m · 255k トークン · $0.38 06 ロック
解決済み · 1000万 · 783kトークン · $1.46 07 初心者の鍵屋
解決済み · 800万 · 845kトークン · $1.05 08ロックイン
解決済み · 1000万 · 988kトークン · $1.43 09 変更なし
解決済み · 600万 · 345kトークン · $0.52 10 Two Doors
解決済み · 4500万 · 280万トークン · $7.26 11 Jelly Throne
解決済み · 800万 · 725kトークン · $1.49 12 Crab Storage
解決済み · 1 時間 33 分 · 320 万トークン · $16.45 13 強盗
解決済み · 1200万 · 862kトークン · $1.81 b1 水没遺跡
解決済み · 1500万 · 999kトークン · $2.34 b2 Sunken Temple
停止 · 2 時間 35 分 · 170 万トークン · $22.12 14 6 時間 46 分 16.3M $61.96

Google Gemini 3.7 フラッシュ
ターミナル-2 01 氷水
解決済み · 700万 · 340万トークン · $0.34 02ターン
解決済み · 1200万 · 990万トークン · $0.75 03 愛情
解決済み · 300万 · 520kトークン · $0.09 04 Pillar Yard
解決済み · 400万 · 428kトークン · $0.08 05 ブリックウォール
解決済み · 1400万 · 950万トークン · $0.74 06 ロック
解決済み · 2,000万 · 2,370万トークン · $1.48 07 初心者の鍵屋
解決済み · 4m · 627k トークン · $0.12 08 ロックイン
解決 · 5m
[切り捨てられた]
まず、API 経由で実行されるモデルの所要時間を見てみましょう。
プロバイダーやモデルの帯域幅によってその時点で変化する可能性がありますが、実効モデル速度の大まかな推定値が得られます。日常のソフトウェア エンジニアリングにおいて、とにかくそれが私たちが気にかけていることです。
0m 3m 10m 30m 1h 2h 4h 6h 8h 10h 12h 14h 16h 18h 0 2 4 6 8 10 12 14 累積壁時計リニア / ログ レベル解決済み Google Grok DeepSeek DeepSeek Anthropic Kim Google OpenAI OpenAI Anthropic Anthropic Z.ai Z.ai Google Gemini 3.7 Flash Grok 4.6 DeepSeek V4 Pro 0813 DeepSeek V4 Flash 0731 Claude Opus 5 キミ K3 Gemini 3.6 Flash GPT-5.6 Sol GPT-5.5 Claude Fable 5 Claude Opus 4.8 GLM-5.2 GLM-5.3 Gemini 3.1 Pro
そしてすごいことに、Gemini 3.7 Flash はトークンごと (広く引用されていますが、ほとんど無関係です) だけでなく、解決された問題ごとにも高速です。ここでは、ほとんどのレベルで、Claude Fable 5 のみが高速です。
Grok 4.6 は GPT-5.6 Sol より 2 倍遅いですが、それでも 15 レベル中 13 を解決できます (GPT-5.6 Sol は 14 を解決します)。これはトップモデルのリーグへの昇格です。
DeepSeek V4 Pro 0813 はトップモデルの中で最も遅いかもしれませんが、最終的に同じトップリーグに到達した最初の無差別級モデルです。これは GLM-5.2 や Kim K3 の結果をはるかに上回っています。
GLM-5.3 は、2 つの線が重なる点までは GLM-5.2 とほぼ同じでした。ただし、もう 1 つのレベルが解決されました。
次にコストを見てみましょう。それは美しいです。
$0 $0.3 $1 $3 $10 $20

$30 $40 $60 $80 0 2 4 6 8 10 12 14 累積コスト線形・解決済みログレベル Google Grok DeepSeek DeepSeek Anthropic Kim Google OpenAI OpenAI Anthropic Anthropic Z.ai Z.ai Google Gemini 3.7 Flash Grok 4.6 DeepSeek V4 Pro 0813 DeepSeek V4 Flash 0731 Claude Opus 5 Kimi K3 ジェミニ 3.6 フラッシュ GPT-5.6 ソル GPT-5.5 クロード フェイブル 5 クロード オーパス 4.8 GLM-5.2 GLM-5.3 ジェミニ 3.1 プロ
Grok 4.6 は Fable 5 より安価ですが、GPT-5.6 Sol や Opus 5 よりは高価です。
Gemini 3.6 Flash は途方もなく高価でしたが、Gemini 3.7 Flash は最も安価なモデルの 1 つです。 12段階では、以前の最安モデルであるGPT-5.6 Solと比べて2倍安い。タスクが 13 個でも 2 倍安いです。 DeepSeek Pro と比較してのみ、わずかに高価です。
DeepSeek V4 Pro 0813 で何が起こるかはまったくの驚きです。オープンウェイトであるだけでなく、すべての中で最も安価です。
GLM-5.3 は以前のものより 2 倍安かった。
以前の実験では、さまざまなハーネスを使用しました。ここでは、Claude Code で実行した DeepSeek を除き、Harbor -default Terminus-2 を使用しました。 Grok Build と DeepSeek Harness もありますが、結果はわずかに悪かった (Grok) か、または大幅に劣っていました (DeepSeek)。とはいえ、これらは新しいツールであり、将来的には改良される可能性があります。ご興味がございましたら、さまざまなハーネスを使用したランニングに関するさらなる洞察を喜んで共有いたします。ハーネスはコストと結果の両方に影響します。
進歩は速く、より多くのプレーヤーがフロンティアゲームに参加し、オープンウェイトモデルははるかに低価格でフロンティアレベルの結果に到達しています。今後数か月以内にさらに多くのモデルが登場することを楽しみにしています。そして湖を解く最初のローカルモデルへ。
私の同僚のピョートル・グラボウスキー氏は次のように気づきました。
「第 2 層の研究室がフロンティアに追いついた」という表現は私の経験と一致します。ここ数日間、私は簡単/中程度の難易度の場合は Grok 4.6 または GLM-5.3 を使用していました。

ty のタスクを実行しましたが、Opus 5/GPT-5.6 Sol との間に大きな違いは見られませんでした (より難しいタスクの場合は、はい、まだギャップがあります)。
クロード コードの価格: 同じトークン、同じモデル、最大 40 倍の価格
Claude Code バイヤーズ ガイド、2026 年 8 月: Max、Team、Enterprise の費用、最良の取引を得る方法、Uber と Shopify が支出を制限する方法。
Kimi K3 はオープン、Opus 5 は良好、DeepSeek V4 Flash は安い: Baba Is You の LLM
パズル ゲーム Baba Is You に基づいた LLM エージェント ベンチマークである Baba Is Bench で、2026 年 7 月の新規リリース Kim K3、Claude Opus 5、Grok 4.5、Gemini 3.6 Flash、および DeepSeek V4 Flash 0731 を評価し、パス率、速度、コストを Claude Fable 5 および GPT-5.6 と比較します。
Baba は Fable 5 と GPT-5.6 Sol によって解決されますが、その代償はどのようなものでしょうか?
私たちはパズル ゲーム Baba Is You を Harbor フレームワークに移植し、Claude、GPT、Gemini、GLM、DeepSeek などの現在のモデルをベンチマークしました。人間の Twitch ストリーマーは、Claude Fable 5 よりも 4 倍高速です。
コーディング エージェントが実際に何を行うかを理解します。

## Original Extract

August 2026 on Baba Is Bench: Gemini 3.7 Flash, Grok 4.6 and DeepSeek V4 Pro 0813 each beat their predecessor while costing 3-20x less. For open-weight GLM-5.3 and Qwen3.8 progress is gradual.

Gemini 3.7 Flash, Grok 4.6, GLM-5.3 and DeepSeek V4 Pro joined the frontier - Quesma Blog Skip to main content How it works Benchmarks Blog About us Token economics Talk to the founder
How it works Benchmarks Blog About us Token economics Talk to the founder Gemini 3.7 Flash, Grok 4.6, GLM-5.3 and DeepSeek V4 Pro joined the frontier
Download PNG This mid August is as hot as July when it comes to LLM releases: Grok 4.6 , Gemini 3.7 Flash , DeepSeek V4 Pro 0813 as well as open-weight models: Qwen3.8 Max , Qwen3.8 27B , and GLM-5.3 .
Intelligence Index vs. Cost per Intelligence Index Task
from Artificial Analysis . Note that it is based on score of benchmarks like
Terminal-Bench v2.1, SciCode, Humanity’s Last Exam, GPQA Diamond - not necessarily fluid intelligence like in abstract
puzzle games of ARC-AGI-3 or Baba is You.
Since these models are smart, I decided to rerun the Baba Is Bench , to see how the models fare on a puzzle game. Even though the game is popular, we check for spoilers - and to our surprise, there are no signs of models knowing solutions ahead - unlike in SWE-bench Verified .
Results are stunning for Grok, Gemini Flash or DeepSeek. Don’t get deceived by minor version changes - they all give qualitatively different results! For both GLM and Qwen the progress was more gradual.
As previously, we do it in two rounds. First, asking to play the initial stage the Intro , then for those who pass at least 7/8 levels, also the next one the Lake .
We show the models we analyze, against dimmed previous ones.
Model 00 baba is you 01 where do i go? 02 now what is this? 03 out of reach 04 still out of reach 05 volcano 06 off limits 07 grass yard pass@1 ↓ pass@3 ↓ Turns ↓ Output tokens ↓ Total cost ↑ Google Gemini 3.7 Flash 100% 100% 53 22k $5.38 Anthropic Claude Opus 5 100% 100% 7 23k $17.15 OpenAI GPT-5.5 100% 100% 8 15k $13.62 Anthropic Claude Fable 5 100% 100% 6 29k $40.98 OpenAI GPT-5.6 Sol 100% 100% 11 10k $11.91 Grok Grok 4.6 96% 100% 25 52k $14.01 Z.ai GLM-5.3 96% 100% 13 47k $6.20 Kimi Kimi K3 96% 100% 9 29k $12.56 Anthropic Claude Opus 4.8 96% 100% 35 46k $63.14 Z.ai GLM-5.2 96% 100% 12 75k $6.20 Google Gemini 3.1 Pro 92% 100% 8 38k $12.41 Google Gemini 3.6 Flash 88% 100% 92 82k $124.31 OpenAI GPT-5.6 Terra 88% 100% 41 55k $34.28 DeepSeek DeepSeek V4 Pro 0813 75% 88% 18 130k $2.94 DeepSeek DeepSeek V4 Flash 75% 100% 55 190k $2.82 Grok Grok 4.5 75% 88% 86 68k $50.66 DeepSeek DeepSeek V4 Flash 0731 75% 88% 22 156k $1.16 Anthropic Claude Sonnet 5 67% 88% 22 151k $42.21 OpenAI GPT-5.6 Luna 63% 88% 118 129k $59.32 Qwen Qwen3.8 Max 58% 75% 9 66k $10.75 Qwen Qwen3.6 27B 29% 38% 24 82k $5.59 DeepSeek DeepSeek V4 Pro 29% 38% 37 183k $11.49 Qwen Qwen3.7 Max 25% 38% 21 132k $16.98 Qwen Qwen3.8 27B 21% 38% 9 134k $9.07 pass time agent steps moves solved timeout wrong answer not solved
The results are nothing short of stunning, as both results went up and prices went down significantly.
Gemini 3.7 Flash now solves every trial, and is over 20x cheaper than its predecessor, Gemini 3.6 Flash.
Grok 4.6 solved all but one attempt (vs Grok 4.5 that missed 6 attempts and one level), while being over 3x cheaper.
DeepSeek V4 Pro 0813 solved 7/8 levels vs its previous version that solved only 3/8 - and while being 4x cheaper.
0% 25% 50% 75% 100% $1 $2 $5 $10 $20 $50 $100 cost of the full intro run, log scale pass@1 Anthropic OpenAI Anthropic OpenAI Kimi Anthropic Z.ai Google OpenAI DeepSeek Anthropic OpenAI Qwen Qwen Qwen Qwen Google Gemini 3.6 Flash Google Gemini 3.7 Flash Grok Grok 4.5 Grok Grok 4.6 DeepSeek DeepSeek V4 Pro DeepSeek DeepSeek V4 Pro 0813 DeepSeek DeepSeek V4 Flash 0731 Z.ai GLM-5.3
With the GLM and Qwen family there is less success:
GLM-5.3 got exactly the same result as GLM-5.2.
Qwen3.8 Max is better than Qwen3.7 Max (way more attempts solved), but still far from completing all levels.
Qwen3.8 27B has worse results than Qwen3.6 27B - though, all within experimental noise.
I was surprised that it didn’t do better. Especially as Qwen3.8 27B scores 52 on Artificial Analysis . From what I’ve seen in the transcripts, it tried to solve in one go, without actually exploring the board. It matches Simon Willison’s experience that Qwen3.8 27B defaults to wildly overthinking things . While Baba Is You clearly needs thinking, it is a balance of thinking and testing. Compare and contrast with previous Grok and Gemini Flash models that were doing a lot and thinking too little.
$1 $2 $5 $10 $20 $50 $100 5 10 20 50 100 agent turns (mean per attempt) * one level unsolved total cost ← more deliberate more trigger-happy → ← cheaper more expensive → Google Gemini 3.7 Flash Anthropic Claude Opus 5 OpenAI GPT-5.5 Anthropic Claude Fable 5 OpenAI GPT-5.6 Sol Grok Grok 4.6 Z.ai GLM-5.3 Kimi Kimi K3 Anthropic Claude Opus 4.8 Z.ai GLM-5.2 Google Gemini 3.1 Pro Google Gemini 3.6 Flash OpenAI GPT-5.6 Terra DeepSeek DeepSeek V4 Pro 0813* DeepSeek DeepSeek V4 Flash Grok Grok 4.5* DeepSeek DeepSeek V4 Flash 0731* Anthropic Claude Sonnet 5* OpenAI GPT-5.6 Luna*
It is interesting to look at the chart of agent turns and price. Grok and Gemini Flash moved from trigger-happy models, to more moderate ones, resulting in a significantly reduced cost.
For models that passed at least 7 out of 8 intro levels, we run them on the next stage, the Lake .
Only one attempt per level, as it is costly. Or at least - it used to be costly with the previous models, as now it is getting significantly better.
Claude Fable 5 Claude Code Claude Code 01 Icy Waters
solved · 4m · 277k tokens · $1.73 02 Turns
solved · 11m · 1.1M tokens · $4.83 03 Affection
solved · 3m · 229k tokens · $0.99 04 Pillar Yard
solved · 8m · 777k tokens · $3.64 05 Brick Wall
solved · 2m · 206k tokens · $0.67 06 Lock
solved · 6m · 795k tokens · $2.76 07 Novice Locksmith
solved · 6m · 695k tokens · $2.46 08 Locked In
solved · 5m · 396k tokens · $1.70 09 Changeless
solved · 3m · 213k tokens · $1.37 10 Two Doors
solved · 16m · 1.8M tokens · $5.96 11 Jelly Throne
solved · 11m · 570k tokens · $3.66 12 Crab Storage
solved · 12m · 1.5M tokens · $5.00 13 Burglary
solved · 10m · 1.1M tokens · $4.52 b1 Submerged Ruins
solved · 11m · 676k tokens · $3.54 b2 Sunken Temple
stopped · 49m · 337k tokens · $10.41 14 2h 36m 10.7M $53.25 Anthropic Claude Opus 5 Claude Code Claude Code 01 Icy Waters
solved · 3m · 365k tokens · $0.60 02 Turns
solved · 16m · 940k tokens · $2.33 03 Affection
solved · 7m · 596k tokens · $0.83 04 Pillar Yard
solved · 13m · 948k tokens · $1.88 05 Brick Wall
solved · 5m · 255k tokens · $0.38 06 Lock
solved · 10m · 783k tokens · $1.46 07 Novice Locksmith
solved · 8m · 845k tokens · $1.05 08 Locked In
solved · 10m · 988k tokens · $1.43 09 Changeless
solved · 6m · 345k tokens · $0.52 10 Two Doors
solved · 45m · 2.8M tokens · $7.26 11 Jelly Throne
solved · 8m · 725k tokens · $1.49 12 Crab Storage
solved · 1h 33m · 3.2M tokens · $16.45 13 Burglary
solved · 12m · 862k tokens · $1.81 b1 Submerged Ruins
solved · 15m · 999k tokens · $2.34 b2 Sunken Temple
stopped · 2h 35m · 1.7M tokens · $22.12 14 6h 46m 16.3M $61.96 Google Gemini 3.7 Flash
Terminus-2 01 Icy Waters
solved · 7m · 3.4M tokens · $0.34 02 Turns
solved · 12m · 9.9M tokens · $0.75 03 Affection
solved · 3m · 520k tokens · $0.09 04 Pillar Yard
solved · 4m · 428k tokens · $0.08 05 Brick Wall
solved · 14m · 9.5M tokens · $0.74 06 Lock
solved · 20m · 23.7M tokens · $1.48 07 Novice Locksmith
solved · 4m · 627k tokens · $0.12 08 Locked In
solved · 5m
[truncated]
First, let’s look at the wall time, for models run via APIs.
While it might change with providers, or model bandwidth at a given time, it gives a ballpark estimate of effective model speed. In everyday software engineering it is what we care for, anyway.
0m 3m 10m 30m 1h 2h 4h 6h 8h 10h 12h 14h 16h 18h 0 2 4 6 8 10 12 14 cumulative wall clock linear · log levels solved Google Grok DeepSeek DeepSeek Anthropic Kimi Google OpenAI OpenAI Anthropic Anthropic Z.ai Z.ai Google Gemini 3.7 Flash Grok 4.6 DeepSeek V4 Pro 0813 DeepSeek V4 Flash 0731 Claude Opus 5 Kimi K3 Gemini 3.6 Flash GPT-5.6 Sol GPT-5.5 Claude Fable 5 Claude Opus 4.8 GLM-5.2 GLM-5.3 Gemini 3.1 Pro
And wow, Gemini 3.7 Flash is fast - not only per token (which is widely cited but largely irrelevant), but per problem solved. Here, for most levels, only Claude Fable 5 is faster.
Grok 4.6, while 2x slower than GPT-5.6 Sol, still manages to solve 13 of 15 levels (GPT-5.6 Sol solves 14). It’s a promotion to the league of top models.
DeepSeek V4 Pro 0813 might be the slowest of the top models - but it is the first open-weight model that finally reached the same top league. It is way beyond GLM-5.2 or Kimi K3 results.
GLM-5.3 was almost the same as GLM-5.2, to the point that two lines overlap; though, it solved one more level.
Now let’s look at costs. It’s beautiful.
$0 $0.3 $1 $3 $10 $20 $30 $40 $60 $80 0 2 4 6 8 10 12 14 cumulative cost linear · log levels solved Google Grok DeepSeek DeepSeek Anthropic Kimi Google OpenAI OpenAI Anthropic Anthropic Z.ai Z.ai Google Gemini 3.7 Flash Grok 4.6 DeepSeek V4 Pro 0813 DeepSeek V4 Flash 0731 Claude Opus 5 Kimi K3 Gemini 3.6 Flash GPT-5.6 Sol GPT-5.5 Claude Fable 5 Claude Opus 4.8 GLM-5.2 GLM-5.3 Gemini 3.1 Pro
Grok 4.6 is cheaper than Fable 5, but more expensive than GPT-5.6 Sol or Opus 5.
While Gemini 3.6 Flash was ridiculously expensive , Gemini 3.7 Flash is one of the cheapest models. At 12 levels, it is twice as cheap as GPT-5.6 Sol, the previous cheapest model. At 13 tasks it is still twice as cheap. Only vs DeepSeek Pro is it slightly more expensive.
What happens with DeepSeek V4 Pro 0813 is a pure wonder. Not only is it open-weight, but also the cheapest of all.
GLM-5.3 was 2x cheaper than its predecessor.
In previous experiments we used various harnesses. Here we used Harbor -default Terminus-2, except DeepSeek, which we ran on Claude Code. While there are Grok Build and DeepSeek Harness , results were slightly worse (Grok) or vastly inferior (DeepSeek). That said, these are newer tools, and likely to get improved in the future. If there is interest, we’re happy to share more insight on running with various harnesses — they influence both cost and results.
Progress is fast - with more players joining the frontier game, and open-weight models reaching frontier-level results at a much lower price. And I am looking forward to seeing more models in the coming months. And to the first local model solving the Lake.
As my colleague Piotr Grabowski noticed:
“Second-tier labs caught up with the frontier” matches my experience: last few days I’ve been using Grok 4.6 or GLM-5.3 for easy/medium difficulty tasks and couldn’t see major difference between them and Opus 5/GPT-5.6 Sol (for harder tasks yes, there’s still gap).
Claude Code pricing: same tokens, same model, up to 40x the price
Claude Code buyer’s guide, August 2026: what Max, Team, and Enterprise cost, how to get the best deal, and how Uber and Shopify cap the spend.
Kimi K3 is Open, Opus 5 is Good, DeepSeek V4 Flash is Cheap: LLMs on Baba Is You
We evaluate July 2026 fresh releases Kimi K3, Claude Opus 5, Grok 4.5, Gemini 3.6 Flash, and DeepSeek V4 Flash 0731 on Baba Is Bench, an LLM agent benchmark based on the puzzle game Baba Is You, comparing pass rate, speed, and cost with Claude Fable 5 and GPT-5.6.
Baba Is Solved by Fable 5 and GPT-5.6 Sol, but at what cost?
We ported the puzzle game Baba Is You to the Harbor framework, and benchmarked current models, including Claude, GPT, Gemini, GLM, and DeepSeek. A human Twitch streamer is 4x faster than Claude Fable 5.
Understand what your coding agents actually do.
