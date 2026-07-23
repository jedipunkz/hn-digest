---
source: "https://juliahub.com/blog/frontier-models-physical-ai-evaluation"
hn_url: "https://news.ycombinator.com/item?id=49022265"
title: "Frontier model evaluation for Physical AI"
article_title: "GPT-5.6 vs Claude Fable 5 for Physical AI, which performs best? - Blog | JuliaHub"
author: "ViralBShah"
captured_at: "2026-07-23T15:03:27Z"
capture_tool: "hn-digest"
hn_id: 49022265
score: 1
comments: 0
posted_at: "2026-07-23T14:30:37Z"
tags:
  - hacker-news
  - translated
---

# Frontier model evaluation for Physical AI

- HN: [49022265](https://news.ycombinator.com/item?id=49022265)
- Source: [juliahub.com](https://juliahub.com/blog/frontier-models-physical-ai-evaluation)
- Score: 1
- Comments: 0
- Posted: 2026-07-23T14:30:37Z

## Translation

タイトル: 物理AIのフロンティアモデル評価
記事のタイトル: GPT-5.6 と物理 AI の Claude Fable 5、どちらが最高のパフォーマンスを発揮しますか? - ブログ |ジュリアハブ
説明: Dyad エージェントの最新のフロンティア モデルを 5 つのモデリングおよびシミュレーションの問題でテストし、精度、コスト、時間、作業スタイルを比較しました。

記事本文:
GPT-5.6 と物理 AI の Claude Fable 5、どちらが最もパフォーマンスが優れていますか?
GPT-5.6 と物理 AI の Claude Fable 5、どちらが最もパフォーマンスが優れていますか?
GPT-5.6 と物理 AI の Claude Fable 5、どちらが最もパフォーマンスが優れていますか?
私たちは、エージェントで最新のフロンティア モデル gpt-5.6-terra を、さまざまな難易度の 5 つのモデリングおよびシミュレーションの問題でテストしました。結果は次のとおりです: claude-fable-5 0.889 加重スコア
$9.60 / トライアル · 16.1 分 gpt-5.6-sol 0.814 加重スコア
$1.74 / トライアル · 13.4 分 gpt-5.6-terra 0.786 加重スコア
$1.25 / トライアル · 12.6 分 gpt-5.6-luna 0.727 加重スコア
$3.26 / トライアル · 25.0 分 コンテナーは、モデル、Dyad ハーネス、ピン留めされた構成、および密閉された問題を組み合わせます。エージェントは、グレーダーがモデルを封印されたグランド トゥルースと比較する前に、モデルを導出、コンパイル、シミュレーション、検証します。物理 AI 評価の仕組み
エージェントを解く ▸ 問題を読む ▸ 物理学を導き出す ▸ dyad/model.dyad を書く ▸ コンパイル ▸ シミュレート・検証 ▌ 採点する 採点者 — 両方をシミュレートし、時間を比較 → グラウンドトゥルースの提出スコア GRADE 1.0 0 0.96 1 · コンテナ MODEL gpt-5.6-terra HARNESS Dyad Agent CONFIG xhigh · 1M ctx · 128k 問題 · シールドされたワークスペース + グラウンド トゥルース ↓ 2 · エージェント ▸ 問題の読み取り ▸ 物理学の導出 ▸ dyad/model.dyad の書き込み ▸ コンパイル ▸ シミュレート · 検証 ↓ 3 · グレーダー 両方をシミュレートし、比較 ↓ 4 · 軌道の一致を評価 0.96 方法 · 制御された評価パイプラインの概略図。測定結果ではありません。
出典 · JuliaHub 内部密封評価 · 2026 年 7 月
物理 AI 評価の仕組み
エージェントを解く ▸ 問題を読む ▸ 物理学を導出する ▸ dyad/model.dyad を書く ▸ コンパイルする ▸ シミュレート・検証する ▌ 採点する 採点者 — 両方をシミュレートし、時間を比較してから、正解の提出スコア グレード 1.0 0 0.9

6 1 · コンテナ MODEL gpt-5.6-terra HARNESS Dyad Agent CONFIG xhigh · 1M ctx · 128k PROBLEM · SEALED ワークスペース + グラウンド トゥルース ↓ 2 · エージェント ▸ 問題の読み取り ▸ 物理学の導出 ▸ dyad/model.dyad の書き込み ▸ コンパイル ▸ シミュレート · 検証 ↓ 3 · グレーダー 両方をシミュレートし、比較 ↓ 4 · グレード軌道一致 0.96 方法 · 制御された評価パイプラインの概略図。測定結果ではありません。
出典 · JuliaHub 内部密封評価 · 2026 年 7 月
物理 AI 評価の仕組み
エージェントを解く ▸ 問題を読む ▸ 物理学を導き出す ▸ dyad/model.dyad を書く ▸ コンパイル ▸ シミュレート・検証 ▌ 採点する 採点者 — 両方をシミュレートし、時間を比較 → グラウンドトゥルースの提出スコア GRADE 1.0 0 0.96 1 · コンテナ MODEL gpt-5.6-terra HARNESS Dyad Agent CONFIG xhigh · 1M ctx · 128k 問題 · シールドされたワークスペース + グラウンド トゥルース ↓ 2 · エージェント ▸ 問題の読み取り ▸ 物理学の導出 ▸ dyad/model.dyad の書き込み ▸ コンパイル ▸ シミュレート · 検証 ↓ 3 · グレーダー 両方をシミュレートし、比較 ↓ 4 · 軌道の一致を評価 0.96 方法 · 制御された評価パイプラインの概略図。測定結果ではありません。
出典 · JuliaHub 内部密封評価 · 2026 年 7 月
物理 AI が生きるか死ぬかは、モデル化された物理が正しいかどうかによって決まります。航空機、分離カラム、または荷電粒子のモデルは、エンコードされた物理学が不可能であっても、クリーンにコンパイルして実行できます。エージェントはテストからのフィードバックによって方向転換し、テストは多くの場合同じエージェントによって作成されるため、エージェント AI はこの障害モードをさらに悪化させます。 Web 開発やコンパイラなどの分野では、正しい動作が含まれておりチェックできるため、ループは十分に機能します。エンジニアリングにおける本当の疑問は、「これは現実の世界と一致しているか?」ということですが、それを解決するのははるかに困難です。エージェントです。

作成したすべてのテストに合格する可能性がありますが、それらのテストは、モデルが実際に使用される領域では当てはまらない単純化に基づいています。
すでに存在する番号には信頼性の問題もあります。モデルプロバイダーが自己報告したベンチマークを額面通りに受け取る人はいませんが、それには十分な理由があります。スコアを公開しているプロバイダーは、スコアを調整するあらゆるインセンティブを持っている同じ当事者です。私たちのインセンティブは別の方向を向いています。 JuliaHub は、ベンダー全体に渡って複数のエージェント バックエンドを備えた Dyad を出荷しており、どのラボのモデルが提供するものであっても、ユーザーが可能な限り最高の Dyad エクスペリエンスを得ることができれば、当社は勝利を収めることができます。 1 つのモデルが物理モデリングに優れている場合は、それを見つけて推奨することが私たちの利益になります。そこで私たちは、実際はどれなのかを調査しました。
どのフロンティア モデルがその作業を最もよく処理するかを測定するために、他のすべてを静止させます。 Dyad AI エージェント ハーネス - モデリングとシミュレーションのワークフロー専用に構築された - は、問題、推論作業 (xhigh)、コンテキスト ウィンドウ (1M)、およびトークン バジェット (128k) とともに、すべてのトライアルにわたって固定されます。唯一の変数はモデルです。OpenAI の GPT 5.6 ファミリー ( terra 、 sol 、 luna ) と Anthropic の claude-fable-5 です。 4 つの主要な問題のそれぞれについてモデルごとに 3 つのトライアル、5 番目の問題についてそれぞれ 1 つの長期的なトライアル、合計 52 の採点された実行。
物理 AI 評価の仕組み。
私たちは内部評価のサブセット、つまりモデリングとシミュレーションの日々の作業から抽出された 5 つの封印された問題を調査の基礎としています。それらを難易度別に並べており、正しい解決策が連鎖しなければならない段階、小さなミスが致命的な場合の慎重な詳細、手の届く簡単なショートカット、仕様の解析からハーネスの修理に至るまでのモデリング自体に関するエンジニアリング作業の観点から定義されています。同じ特性に対してそれぞれを選択しました。それらは解決策を許可します。

物理的に間違っているにもかかわらず、グリーンにコンパイルして実行することはできません。したがって、グレーディングではコードは無視され、完全なシミュレートされた軌道が封印されたグラウンド トゥルースに対してスコア付けされます。 5 つの中で最も難しいのは NASA の HL-20 飛行体です。このビデオでは、この問題の簡単なバージョンが詳しく説明されています。研究では、より強力で密閉されたバージョンが実行されます。
私たちはすべてのモデルを完全なスイート (P1 ～ P4 でそれぞれ 3 回のトライアル、長い期間の P5 で 1 回) を通して実行し、すべてのトライアルを封印されたグラウンド トゥルースに照らして採点します。問題ごとのスコアは、難易度に応じた加重平均を通じて 1 つの数値にまとめられます。最も難しいものに重みが最も高くなります。コストと時間には重み付けはありません。トライアルあたりの平均コストとトライアルあたりの平均実時間はそれぞれタブを取得します。
コストと時間に対する加重スコア
$2 $4 $6 $8 $10 0.70 0.75 0.80 0.85 0.90 トライアルあたりのコスト (USD) 難易度加重スコア gpt-5.6-terra 0.786 · $1.25 gpt-5.6-sol 0.814 · $1.74 gpt-5.6-luna 0.727 · $3.26 claude-fable-5 0.889 · $9.60 gpt-5.6-terra 加重スコア 0.786 トライアルあたりのコスト $1.25 gpt-5.6-sol 加重スコア 0.814 トライアルあたりのコスト $1.74 gpt-5.6-luna 加重スコア 0.727 トライアルあたりのコスト $3.26 claude-fable-5 加重スコア 0.889 トライアルあたりのコスト $9.60 コストに対するスコアそして時間。 Fable は、トライアル 9.60 ドル、研究 125 ドルという価格でリードしています。これは GPT の 3 ～ 8 倍です。 Sol は価値の高い製品であり、トライアルあたり 1.74 ドルでメリットが 2 番目にあります。 Terraは最も安くて早いです。ルナはあらゆる軸に軌跡を描きます。
方法 · 難易度加重スコアでは、P1 ～ P4 では 3 回の試行、P5 では 1 回の試行を使用します。コストと実時間は、トライアルごとの加重されていない平均値です。
出典 · JuliaHub 内部密封評価 · 2026 年 7 月
コストと時間に対する加重スコア
$2 $4 $6 $8 $10 0.70 0.75 0.80 0.85 0.90 トライアルあたりのコスト (USD) 難易度加重スコア gpt-5.6-terra 0.786 · $1.25 gpt-5.6-sol 0.814 · $1.74 g

pt-5.6-luna 0.727 · $3.26 claude-fable-5 0.889 · $9.60 gpt-5.6-terra 加重スコア 0.786 トライアルあたりのコスト $1.25 gpt-5.6-sol 加重スコア 0.814 トライアルあたりのコスト $1.74 gpt-5.6-luna 加重スコア 0.727 トライアルあたりのコスト$3.26 claude-fable-5 加重スコア トライアルあたり 0.889 コスト $9.60 コストと時間に対するスコア。 Fable は、トライアル 9.60 ドル、研究 125 ドルという価格でリードしています。これは GPT の 3 ～ 8 倍です。 Sol は価値の高い製品であり、トライアルあたり 1.74 ドルでメリットが 2 番目にあります。 Terraは最も安くて早いです。ルナはあらゆる軸に軌跡を描きます。
方法 · 難易度加重スコアでは、P1 ～ P4 では 3 回の試行、P5 では 1 回の試行を使用します。コストと実時間は、トライアルごとの加重されていない平均値です。
出典 · JuliaHub 内部密封評価 · 2026 年 7 月
コストと時間に対する加重スコア
$2 $4 $6 $8 $10 0.70 0.75 0.80 0.85 0.90 トライアルあたりのコスト (USD) 難易度加重スコア gpt-5.6-terra 0.786 · $1.25 gpt-5.6-sol 0.814 · $1.74 gpt-5.6-luna 0.727 · $3.26 claude-fable-5 0.889 · $9.60 gpt-5.6-terra 加重スコア 0.786 トライアルあたりのコスト $1.25 gpt-5.6-sol 加重スコア 0.814 トライアルあたりのコスト $1.74 gpt-5.6-luna 加重スコア 0.727 トライアルあたりのコスト $3.26 claude-fable-5 加重スコア 0.889 トライアルあたりのコスト $9.60 コストに対するスコアそして時間。 Fable は、トライアル 9.60 ドル、研究 125 ドルという価格でリードしています。これは GPT の 3 ～ 8 倍です。 Sol は価値の高い製品であり、トライアルあたり 1.74 ドルでメリットが 2 番目にあります。 Terraは最も安くて早いです。ルナはあらゆる軸に軌跡を描きます。
方法 · 難易度加重スコアでは、P1 ～ P4 では 3 回の試行、P5 では 1 回の試行を使用します。コストと実時間は、トライアルごとの加重されていない平均値です。
出典 · JuliaHub 内部密封評価 · 2026 年 7 月
上記の単一の数字には構造があります。問題ごとに、同じ 3 つのレンズを試用して平均したもの:
問題ごとのモデルのスコア、コスト、時間
0.4 0.6 0.8 1.0 P1 P2 P3 P4 P5 スコア (a

vg トライアルより）簡単→難しい P1 簡単な寓話 1.000 ソル 1.000 テラ 1.000 ルナ 1.000 P2 寓話 1.000 ソル 1.000 テラ 1.000 ルナ 0.667 P3 寓話 1.000 ソル 0.667 テラ 1.000 ルナ1.000 P4 寓話 1.000 ソル 1.000 テラ 0.667 ルナ 1.000 P5 ベリーハード寓話 0.690 ソル 0.660 テラ 0.590 ルナ 0.383 集合体はどこから来たのか。スコア的には、フィールドはより簡単な問題を通じて天井を維持し、各 GPT は中央で 1 回滑り、すべてのラインが 5 位で下方にブレイクします。コストと時間の点では、P5 にお金と時間が費やされます。
方法 · P1 ～ P4 値は 3 回の試験の平均値です。 P5 は、モデルごとに 1 つの長期トライアルです。
出典 · JuliaHub 内部密封評価 · 2026 年 7 月
問題ごとのモデルのスコア、コスト、時間
0.4 0.6 0.8 1.0 P1 P2 P3 P4 P5 スコア (試行全体の平均) 簡単 → 難しい P1 簡単な寓話 1.000 ソル 1.000 テラ 1.000 ルナ 1.000 P2 寓話 1.000 ソル 1.000 テラ 1.000 ルナ 0.667 P3 寓話1.000 ソル 0.667 テラ 1.000 ルナ 1.000 P4 寓話 1.000 ソル 1.000 テラ 0.667 ルナ 1.000 P5 ベリーハード寓話 0.690 ソル 0.660 テラ 0.590 ルナ 0.383 骨材の由来。スコア的には、フィールドはより簡単な問題を通じて天井を維持し、各 GPT は中央で 1 回滑り、すべてのラインが 5 位で下方にブレイクします。コストと時間の点では、P5 にお金と時間が費やされます。
方法 · P1 ～ P4 値は 3 回の試験の平均値です。 P5 は、モデルごとに 1 つの長期トライアルです。
出典 · JuliaHub 内部密封評価 · 2026 年 7 月
問題ごとのモデルのスコア、コスト、時間
0.4 0.6 0.8 1.0 P1 P2 P3 P4 P5 スコア (試行全体の平均) 簡単 → 難しい P1 簡単な寓話 1.000 ソル 1.000 テラ 1.000 ルナ 1.000 P2 寓話 1.000 ソル 1.000 テラ 1.000 ルナ 0.667 P3 寓話1.000 ソル 0.667 テラ 1.000 ルナ 1.000 P4 寓話 1.000 ソル 1.000 テラ 0.667 ルナ 1.000 P5 ベリーハード寓話 0.690 ソル 0.660 テラ 0.590 ルナ 0.383

レゲートはから来ています。スコア的には、フィールドはより簡単な問題を通じて天井を維持し、各 GPT は中央で 1 回滑り、すべてのラインが 5 位で下方にブレイクします。コストと時間の点では、P5 にお金と時間が費やされます。
方法 · P1 ～ P4 値は 3 回の試験の平均値です。 P5 は、モデルごとに 1 つの長期トライアルです。
出典 · JuliaHub 内部密封評価 · 2026 年 7 月
各モデルには独自のアプローチがあります。
私たちはすべての裁判の記録を分析します。各ツール呼び出しは、問題の方向付け、物理学の導出とテスト、モデルの編集、コンパイル、またはコピーするものの検索など、その時点でエージェントが行っていたことによって分類され、試行ごとにカウントされ、モデルごとに平均されます。その結果、ワークスタイルの指紋が得られます。
モデルごとのワークスタイルの特徴
10 20 30 40 50 60 トライアルあたりのツールコール数 (平均) terra 50 % 28 % 12 % 10 % 0 % 41 コール sol 49 % 26 % 14 % 12 % 0 % 33 コール luna 56 % 27 % 9 % 7 % 1 % 59 コール寓話 38 % 39 % 12 % 11 % 0 % 28 呼び出し 方向付けと読み取り 導出とテスト モデルの編集 コンパイル ハント テラ 50 % 28 % 12 % 41 呼び出し sol 49 % 26 % 14 % 33 呼び出し ルナ 56 % 27 % 9 % 59 呼び出し 寓話 38 % 39 % 28 コール 10 20 30 40 50 60 トライアルごとのツールコール (平均) 同じ作業の 2 つのビュー。電話で見ると、ルナが最も多忙なエージェントであり、ファブルが最も無駄なエージェントです。時間の経過とともに、物理学の導出とテストがすべてのモデルのクロックを支配します。寓話は一度書きます。 Terra は編集に多くの時間を費やします。ルナは一人で狩りに時間を費やします。
メソッド・へ

[切り捨てられた]

## Original Extract

We tested the latest frontier models in the Dyad agent on five modeling and simulation problems, comparing accuracy, cost, time, and work style.

GPT-5.6 vs Claude Fable 5 for Physical AI, which performs best?
GPT-5.6 vs Claude Fable 5 for Physical AI, which performs best?
GPT-5.6 vs Claude Fable 5 for Physical AI, which performs best?
We tested the latest frontier models gpt-5.6-terra in our agent, on five modeling and simulation problems of varying difficulty. Here are the results: claude-fable-5 0.889 weighted score
$9.60 / trial · 16.1 min gpt-5.6-sol 0.814 weighted score
$1.74 / trial · 13.4 min gpt-5.6-terra 0.786 weighted score
$1.25 / trial · 12.6 min gpt-5.6-luna 0.727 weighted score
$3.26 / trial · 25.0 min A container combines a model, the Dyad harness, pinned configuration, and a sealed problem. The agent derives, compiles, simulates, and verifies a model before a grader compares it with sealed ground truth. How one physical-AI evaluation works
solve THE AGENT ▸ read problem ▸ derive physics ▸ write dyad/model.dyad ▸ compile ▸ simulate · verify ▌ grade it THE GRADER — SIMULATE BOTH, THEN COMPARE time → ground truth submitted score GRADE 1.0 0 0.96 1 · the container MODEL gpt-5.6-terra HARNESS Dyad Agent CONFIG xhigh · 1M ctx · 128k PROBLEM · SEALED workspace + ground truth ↓ 2 · the agent ▸ read problem ▸ derive physics ▸ write dyad/model.dyad ▸ compile ▸ simulate · verify ↓ 3 · the grader Simulate both, then compare ↓ 4 · grade trajectory match 0.96 Method · Schematic of the controlled evaluation pipeline; not a measured result.
Source · JuliaHub internal sealed evaluation · July 2026
How one physical-AI evaluation works
solve THE AGENT ▸ read problem ▸ derive physics ▸ write dyad/model.dyad ▸ compile ▸ simulate · verify ▌ grade it THE GRADER — SIMULATE BOTH, THEN COMPARE time → ground truth submitted score GRADE 1.0 0 0.96 1 · the container MODEL gpt-5.6-terra HARNESS Dyad Agent CONFIG xhigh · 1M ctx · 128k PROBLEM · SEALED workspace + ground truth ↓ 2 · the agent ▸ read problem ▸ derive physics ▸ write dyad/model.dyad ▸ compile ▸ simulate · verify ↓ 3 · the grader Simulate both, then compare ↓ 4 · grade trajectory match 0.96 Method · Schematic of the controlled evaluation pipeline; not a measured result.
Source · JuliaHub internal sealed evaluation · July 2026
How one physical-AI evaluation works
solve THE AGENT ▸ read problem ▸ derive physics ▸ write dyad/model.dyad ▸ compile ▸ simulate · verify ▌ grade it THE GRADER — SIMULATE BOTH, THEN COMPARE time → ground truth submitted score GRADE 1.0 0 0.96 1 · the container MODEL gpt-5.6-terra HARNESS Dyad Agent CONFIG xhigh · 1M ctx · 128k PROBLEM · SEALED workspace + ground truth ↓ 2 · the agent ▸ read problem ▸ derive physics ▸ write dyad/model.dyad ▸ compile ▸ simulate · verify ↓ 3 · the grader Simulate both, then compare ↓ 4 · grade trajectory match 0.96 Method · Schematic of the controlled evaluation pipeline; not a measured result.
Source · JuliaHub internal sealed evaluation · July 2026
Physical AI lives or dies on whether the modeled physics is correct. A model of an aircraft, a separation column or a charged particle can compile and run cleanly while the physics it encodes is impossible. Agentic AI makes this failure mode worse, because agents steer by feedback from tests, and the tests are often written by the same agent. In domains like web development or compilers that loop works well enough, since correct behavior is contained and checkable. In engineering, the real question is “does this match the real world?”, and that is much harder to close: the agent can pass every test it wrote while those tests rest on simplifications that don’t hold in the regime where the model will actually be used.
There is also a trust problem with the numbers that already exist. Nobody takes the model providers’ self-reported benchmarks at face value, and for good reason: the provider publishing the score is the same party with every incentive to tune for it. Our incentives point the other way. JuliaHub ships Dyad with multiple agent backends across vendors, and we win when our users get the best possible Dyad experience, whichever lab’s model delivers it. If one model is better at physical modeling, it is in our interest to find that out and recommend it. So we investigated: which one actually is?
To measure which frontier model handles that work best, we hold everything else still. The Dyad AI agent harness - purpose-built for modeling & simulation workflows - is pinned across every trial, along with the problems, the reasoning effort (xhigh), the context window (1M) and the token budget (128k). The only variable is the model: OpenAI’s GPT 5.6 family ( terra , sol , luna ) against Anthropic’s claude-fable-5 . Three trials per model on each of the four core problems, one long-horizon trial each on the fifth: 52 graded runs in all.
How Physical AI Evaluation Works.
We ground our exploration in a subset of our internal evals: five sealed problems drawn from the daily work of modeling & simulation. We order them by difficulty, defined in terms of the stages a correct solution must chain through, the careful details where a small slip is fatal, the easy shortcuts within reach, and the engineering work around the modeling itself - from parsing specs to repairing harnesses. We selected each for the same property: they admit solutions that compile and run green while being physically wrong. Grading therefore ignores the code and scores the full simulated trajectory against sealed ground truth. The hardest of the five is NASA’s HL-20 flight vehicle. An easier version of the problem is walked through in detail in this video ; the study runs a harder, sealed variant.
We run every model through the full suite - three trials on each of P1–P4, one on the long-horizon P5 - and grade every trial against sealed ground truth. Per-problem scores fold into a single number through a difficulty-weighted average, the hardest weighing the most. Cost and time carry no weighting: average cost per trial and average wall-clock per trial each get a tab:
Weighted score against cost and time
$2 $4 $6 $8 $10 0.70 0.75 0.80 0.85 0.90 cost per trial (USD) difficulty-weighted score gpt-5.6-terra 0.786 · $1.25 gpt-5.6-sol 0.814 · $1.74 gpt-5.6-luna 0.727 · $3.26 claude-fable-5 0.889 · $9.60 gpt-5.6-terra weighted score 0.786 cost per trial $1.25 gpt-5.6-sol weighted score 0.814 cost per trial $1.74 gpt-5.6-luna weighted score 0.727 cost per trial $3.26 claude-fable-5 weighted score 0.889 cost per trial $9.60 Score against cost and time. Fable leads at a price: $9.60 a trial, $125 for the study — three to eight times the GPTs. Sol is the value story, second on merit at $1.74 a trial. Terra is cheapest and fastest. Luna trails on every axis.
Method · Difficulty-weighted score uses three trials for P1–P4 and one trial for P5. Cost and wall-clock time are unweighted means per trial.
Source · JuliaHub internal sealed evaluation · July 2026
Weighted score against cost and time
$2 $4 $6 $8 $10 0.70 0.75 0.80 0.85 0.90 cost per trial (USD) difficulty-weighted score gpt-5.6-terra 0.786 · $1.25 gpt-5.6-sol 0.814 · $1.74 gpt-5.6-luna 0.727 · $3.26 claude-fable-5 0.889 · $9.60 gpt-5.6-terra weighted score 0.786 cost per trial $1.25 gpt-5.6-sol weighted score 0.814 cost per trial $1.74 gpt-5.6-luna weighted score 0.727 cost per trial $3.26 claude-fable-5 weighted score 0.889 cost per trial $9.60 Score against cost and time. Fable leads at a price: $9.60 a trial, $125 for the study — three to eight times the GPTs. Sol is the value story, second on merit at $1.74 a trial. Terra is cheapest and fastest. Luna trails on every axis.
Method · Difficulty-weighted score uses three trials for P1–P4 and one trial for P5. Cost and wall-clock time are unweighted means per trial.
Source · JuliaHub internal sealed evaluation · July 2026
Weighted score against cost and time
$2 $4 $6 $8 $10 0.70 0.75 0.80 0.85 0.90 cost per trial (USD) difficulty-weighted score gpt-5.6-terra 0.786 · $1.25 gpt-5.6-sol 0.814 · $1.74 gpt-5.6-luna 0.727 · $3.26 claude-fable-5 0.889 · $9.60 gpt-5.6-terra weighted score 0.786 cost per trial $1.25 gpt-5.6-sol weighted score 0.814 cost per trial $1.74 gpt-5.6-luna weighted score 0.727 cost per trial $3.26 claude-fable-5 weighted score 0.889 cost per trial $9.60 Score against cost and time. Fable leads at a price: $9.60 a trial, $125 for the study — three to eight times the GPTs. Sol is the value story, second on merit at $1.74 a trial. Terra is cheapest and fastest. Luna trails on every axis.
Method · Difficulty-weighted score uses three trials for P1–P4 and one trial for P5. Cost and wall-clock time are unweighted means per trial.
Source · JuliaHub internal sealed evaluation · July 2026
The single number above has anatomy. Per problem, averaged across trials, the same three lenses:
Per-problem model score, cost, and time
0.4 0.6 0.8 1.0 P1 P2 P3 P4 P5 score (avg over trials) easier → harder P1 easy fable 1.000 sol 1.000 terra 1.000 luna 1.000 P2 fable 1.000 sol 1.000 terra 1.000 luna 0.667 P3 fable 1.000 sol 0.667 terra 1.000 luna 1.000 P4 fable 1.000 sol 1.000 terra 0.667 luna 1.000 P5 very hard fable 0.690 sol 0.660 terra 0.590 luna 0.383 Where the aggregate comes from. On score, the field holds the ceiling through the easier problems, each GPT slips once in the middle, and every line breaks downward at P5. On cost and time, P5 is where the money and hours go.
Method · P1–P4 values are means across three trials; P5 is one long-horizon trial per model.
Source · JuliaHub internal sealed evaluation · July 2026
Per-problem model score, cost, and time
0.4 0.6 0.8 1.0 P1 P2 P3 P4 P5 score (avg over trials) easier → harder P1 easy fable 1.000 sol 1.000 terra 1.000 luna 1.000 P2 fable 1.000 sol 1.000 terra 1.000 luna 0.667 P3 fable 1.000 sol 0.667 terra 1.000 luna 1.000 P4 fable 1.000 sol 1.000 terra 0.667 luna 1.000 P5 very hard fable 0.690 sol 0.660 terra 0.590 luna 0.383 Where the aggregate comes from. On score, the field holds the ceiling through the easier problems, each GPT slips once in the middle, and every line breaks downward at P5. On cost and time, P5 is where the money and hours go.
Method · P1–P4 values are means across three trials; P5 is one long-horizon trial per model.
Source · JuliaHub internal sealed evaluation · July 2026
Per-problem model score, cost, and time
0.4 0.6 0.8 1.0 P1 P2 P3 P4 P5 score (avg over trials) easier → harder P1 easy fable 1.000 sol 1.000 terra 1.000 luna 1.000 P2 fable 1.000 sol 1.000 terra 1.000 luna 0.667 P3 fable 1.000 sol 0.667 terra 1.000 luna 1.000 P4 fable 1.000 sol 1.000 terra 0.667 luna 1.000 P5 very hard fable 0.690 sol 0.660 terra 0.590 luna 0.383 Where the aggregate comes from. On score, the field holds the ceiling through the easier problems, each GPT slips once in the middle, and every line breaks downward at P5. On cost and time, P5 is where the money and hours go.
Method · P1–P4 values are means across three trials; P5 is one long-horizon trial per model.
Source · JuliaHub internal sealed evaluation · July 2026
Each model has its own approach .
We analyze the transcripts behind every trial. Each tool call is bucketed by what the agent was doing at that moment - orienting in the problem, deriving and testing physics, editing the model, compiling, or hunting for something to copy - then counted per trial and averaged per model. The result is a work-style fingerprint:
Work-style fingerprints by model
10 20 30 40 50 60 tool calls per trial (mean) terra 50 % 28 % 12 % 10 % 0 % 41 calls sol 49 % 26 % 14 % 12 % 0 % 33 calls luna 56 % 27 % 9 % 7 % 1 % 59 calls fable 38 % 39 % 12 % 11 % 0 % 28 calls orient & read derive & test edit the model compile hunt terra 50 % 28 % 12 % 41 calls sol 49 % 26 % 14 % 33 calls luna 56 % 27 % 9 % 59 calls fable 38 % 39 % 28 calls 10 20 30 40 50 60 tool calls per trial (mean) Two views of the same work. By calls, Luna is the busiest agent and Fable the leanest. By time, deriving and testing physics dominates every model's clock. Fable writes once; Terra spends more time editing; Luna alone spends time hunting.
Method · To

[truncated]
