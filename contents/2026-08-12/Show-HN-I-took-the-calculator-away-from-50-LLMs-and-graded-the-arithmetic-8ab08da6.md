---
source: "https://github.com/DmitriGoloubentsev/LLM-Tests"
hn_url: "https://news.ycombinator.com/item?id=49271070"
title: "Show HN: I took the calculator away from 50 LLMs and graded the arithmetic"
article_title: "GitHub - DmitriGoloubentsev/LLM-Tests · GitHub"
author: "NatalijaAAD"
captured_at: "2026-08-12T12:45:40Z"
capture_tool: "hn-digest"
hn_id: 49271070
score: 4
comments: 0
posted_at: "2026-08-12T12:02:25Z"
tags:
  - hacker-news
  - translated
---

# Show HN: I took the calculator away from 50 LLMs and graded the arithmetic

- HN: [49271070](https://news.ycombinator.com/item?id=49271070)
- Source: [github.com](https://github.com/DmitriGoloubentsev/LLM-Tests)
- Score: 4
- Comments: 0
- Posted: 2026-08-12T12:02:25Z

## Translation

タイトル: HN を表示: 50 個の LLM から電卓を取り出し、算術を採点しました
記事のタイトル: GitHub - DmitriGoloubentsev/LLM-Tests · GitHub
説明: GitHub でアカウントを作成して、DmitriGoloubentsev/LLM-Tests の開発に貢献します。

記事本文:
GitHub - DmitriGoloubentsev/LLM-Tests · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ドミトリ・ゴルベンツェフ
/
LLM テスト
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
14 コミット 14 コミット docs docs test1 test1 test1v2 test1v2 .gitignore .gitignore README.md README.md make_hero_figure.py make_hero_figure.py run_openrouter_free.sh run_openrouter_free.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
LLMはフォローできますか

w 派生 — それとも単に電卓に手を伸ばすだけですか?
アシスタントに exp(-0.9513299347) を 15 桁まで尋ねれば、完璧な答えが得られます。それは書きます
Python を 3 行記述して実行します。これはモデルではなくサンドボックスを測定します。
これは電卓を取り上げた同じ試験です。利用可能なツールはありません - 制限もありません
落胆し、欠席する — したがって、モデルは学生が本を閉じた紙の上で行うことを実行する必要があります。
記憶している定数に引数を適用し、テイラー級数を展開し、12 個の級数を処理します。
16桁の掛け算を手作業で行う。 1 つが滑り落ちた数千の算術トークン
答えを破壊します。
それをブラフする方法はありません。答えは 15 位まで正しいか、そうでないかのどちらかです。
採点者は桁の小数点まで判断できます。
モデルが正しい数値を生成できるかどうかをテストしているわけではありません。
スレッドを失うことなく長い導出をたどることができるかどうかをテストしています。
ツールを使用すると、すべてのモデルのスコアが 100% になります。これがなければ、拡散は膨大になります。
100 個以上の実行、50 個以上のモデル、8 つのエンドポイント、1 つのプロンプト、ポイントごとに 1 つのステートレス呼び出し。
同じ質問をすると、モデルは間にほとんど何もない 2 つの母集団に分かれます。
努力では1,300倍、答えでは10桁の差。
そして、モデルがどちら側に着地するかは、そのサイズ、ベンダー、価格によっては予測できません。
ミストラルのショートカットは24Bと675Bにあります。プールサイドの小型モデルは、大型の兄弟モデルを上回ります。
そして、OpenAI 独自のラダーは、GPT-4.1 から 5.4 まで、単一のバージョン ステップで一線を越えます。
〜10トークンのショートカット。 GPT-5.5 は 8,201 で派生します。
4 回の実行、同じ関数、同じ 101 ポイント、同じプロンプト。上: 呼び出し可能な引数に関する 1 つのモデル、
引数ではそれはできません - 6 桁が消えます。下: あらゆる値を導き出すモデルとその内容
スロットルを絞ったときに起こります

その考え。
偽りのない。自信を持った文章は何も得ません。正しい数字のみがスコアになります。
構造的に正しい参照に対して、桁の小数点まで等級付け可能。
オンデマンドでリコールプルーフ。引数はシードされたハッシュによって小数点以下 10 桁までジッターされるため、
トレーニング コーパスには現れません。モデルは導出するか、推測する必要があります。
これは、他のすべての eval が混同するもの、つまりメソッドの認識とそれを実行するものを分離します。
数千のトークンにわたってエラーなし。
メトリクスは正しい数値 = -log10(相対誤差) で、上限は 15 (IEEE-754 二重協定) です。
合格/不合格では、興味深いものがすべて隠されてしまいます。「4.8 桁」と「15 桁」は両方とも「誤り」です。
ブールグレーダー。
「電卓なし」を強制する方法
プレーンな API では、これは無料です。ツール配列は送信されないため、呼び出すものは何もありません。危険な道
これは Claude Code CLI バックエンドであり、通常、モデルにはインタープリターが含まれています。そこでは次のように実行されます。
--tools "" — 組み込みツールセット全体が無効になります。
使い捨ての $HOME と空の作業ディレクトリなので、CLAUDE.md はコンテキストに到達しません。
すべての応答は num_turns == 1 であることが検証されます。ツール呼び出しには 2 番目のターンが必要です。
たった 1 回のターンで、答えが Python からではなくモデルから出たことが証明されます。
グリッド
答えられる質問
テスト1
x = -1.00 … 1.00 ステップ 0.02 (101 ラウンド引数)
記憶されている値はどの程度正確ですか?
テスト1v2
同じ 101 ポイントで、各ポイントは決定論的な ~10 桁のオフセット ( -0.9513299347 , …) によって微調整されます。
引数がどのコーパスにも存在しない場合、どの程度正確ですか?
test1v2 は --func exp|sin|cos|log|sqrt をサポートします。
それらの間のアブレーションがポイントです。 gemma4:31b はラウンド議論で 10.40 中央値を獲得し、
リコールプルーフのものでは 4.76 — 同じモデル、同じプロンプト、同じ回答あたり最大 18 トークン。 test1 のスコア
ルックアップテーブルはほぼフルでした。
最悪の議論の論点はルーではない

検出エラー: exp(0.58) を要求すると、応答されました
1786038444035 — 小数点が欠落しているため、調べている間に答えが 10¹² ずれています
完璧に整った形。
実際に異なる動作をする 4 つのグループに分類されます。 tok/pt を中央値に沿って読み取ります —
モデルが答えを導き出したのか、それとも答えを思い出したかがわかります。
1 — デリバ · 数千のトークン、倍精度ラインに到達
モデル
ふん
答えた
中央値
最悪の
トーク/ポイント
コスト
openai/gpt-5.5
経験値
101/101
15.00
3.08
8,201
—
openai/gpt-5.5
罪
20/20
15.00
12.18
9,344
—
openai/gpt-5.6-sol
経験値
5/5
15.00
14.17
4,254
$0.64
人間/クロード作品-5
経験値
10/10
15.00
13.38
16,688
$4.17
ディープシーク-v4-フラッシュ
経験値
54/101
15.00
13.17
51,695
$1.46
deepseek-v4-フラッシュ努力=低
経験値
98/101
15.00
5.04
26,520
$0.75
クロード オーパス 5 (CLI) 努力値 = 最大
経験値
20/20
15.00
5.56
15,690
サブ
z-ai/glm-5.2
経験値
18/101
15.00
8.37
30,181
無料
クロード寓話 5 (CLI)
経験値
62/101
14.68
10.49
5,503
サブ
月書体/kimi-k2.7-code
経験値
7/20
15.00
12.14
55,391
無料
クロード作品5 (CLI)
罪
101/101
14.38
8.32
7,924
サブ
月書体/kimi-k2.6
罪
15/20
14.14
7.89
54,103
無料
月書体/kimi-k2.6
経験値
13/20
13.68
8.31
57,209
無料
2 — まばらな中間 · 純粋に派生しますが、まだ不十分です
モデル
ふん
答えた
中央値
最悪の
トーク/ポイント
人間性/クロード・ソネット-5
経験値
10/10
12.10
9.30
22,630
オープンナイ/o3
経験値
1/3
11.88
11.88
31,987
メタ/ミューズ-グリマー-30b
経験値
3/3
11.46
9.82
27,817
ディープシーク-v4-フラッシュ:プレビュー
経験値
3/3
11.44
9.50
49,470
qwen/qwen3.7-max
経験値
3/3
11.40
11.32
11,843
anthropic/claude-opus-4.8 努力=高
経験値
10/10
11.13
9.20
4,141
openai/o4-mini
経験値
3/3
10.96
9.82
15,557
ThinkingMachine/インクリング-スモール
経験値
3/3
10.01
6.18
25,808
gpt-oss:120b
経験値
98/101
9.55
3.18
9,105
google/gemini-3.5-フラッシュ
経験値
20/20
9.24
6.14
22,401
プールサイド/ラグナ-xs-2.1
経験値
61/101
7.60
3.08
25,599
gpt-oss:20b
経験値

93/101
6.37
2.44
26,420
anthropic/claude-sonnet-4-6 (near.ai)
経験値
3/3
6.33
4.76
1,292
qwen/qwen3.7-フラッシュ
経験値
80/101
5.63
2.84
27,579
qwen3.5:397b
経験値
101/101
5.19
3.95
6,015
nvidia/nemotron-3-ultra (オラマ)
経験値
3/3
4.50
4.13
12,894
openai/o3-mini
経験値
3/3
3.91
2.96
4,588
3 — ショートカット · 最大 18 個のトークン、3 ～ 7 桁、失敗せず改善もしない
モデル
サイズ
答えた
中央値
トーク/ポイント
google/gemini-2.5-pro
—
3/3
5.23
22
google/gemma-4-26b-a4b (教育省)
26B
3/3
4.91
19
gemma4:31b (密)
31B
101/101
4.76
18
anthropic/claude-opus-4.8 (デフォルト)
—
10/10
4.66
9
アップステージ/ソーラープロ4
—
101/101
4.41
18
openai/gpt-5.4
—
3/3
4.38
11
ミストラル-ラージ-3:675b
675B
101/101
4.29
18
openai/gpt-5.2
—
3/3
4.22
10
プールサイド/ラグナ-S-2.1
—
20/101
4.13
18(応答時)
openai/gpt-5.1
—
3/3
4.03
16
qwen3-コーダー
～30B
101/101
3.72
18
openai/gpt-4.1
—
3/3
3.57
7
開発ストラル:24b
24B
101/101
3.46
16
openai/gpt-5.4-mini
—
3/3
3.40
16
qwen2.5-coder:14b
14B
101/101
2.93
18
4 — 終了しない · 出力上限の理由、空のコンテンツを返す
モデル
答えた
トーク/ポイント
エンドポイント
ミニマックス-m3
0/101
32,000（上限）
オラマ・クラウド
ミニマックス-m2.7
0/20
65,536 (上限)
オラマ・クラウド
inclusionai/ling-3.0-tiny
0/101
27,564
オープンルーター
inclusionai/ling-3.0-フラッシュ
0/3
32,000（上限）
オープンルーター
cohere/north-mini-code
0/3
26,176
オープンルーター
ディープシーク-v4-プロ
0/3
65,536 (上限)
オラマ・クラウド
glm-5.1
0/3
32,768（上限）
オラマ・クラウド
クロード作品-4.6 努力=高い、上限なし
0/10
65,536 (上限)
オープンルーター
そして、独自のカテゴリーの 1 つである Nemotron ファミリーは、あらゆる規模でこの課題に取り組んでいます。
nemotron-3-nano-30b は、101 ポイントで正解桁の中央値 0.00 を獲得しました (1 つの回答は
229144000.0 ここで、答えは 0.89 ); nemotron-3-ultra-550b は 3.69 を管理しました (Ollama では 4.50)。
推論調整された nemotron-3-nano-omni-30b は、32,7 を燃焼した後、スコアが -0.00 でした。

1 あたり 68 トークン
ポイント; nemotron-3-super-120b は、動作中のエンドポイントで 6.32 で 1/20 と応答しました。 4つのモデル、
4 つのサイズ、3 つのエンドポイント、1 つの結果。
この家族からの生存者への警告。 NVIDIA NIM で実行される初期の nemotron-3-super-120b
101点中8点で中央値15.00点（ドライバーグレードの数字）を獲得し、93点で負けた
HTTP 504、783 tok/pt。 OpenRouter の無料枠で再実行すると、6.32 桁で 1/20 と答えます。
6,924トーク/ポイント。 15.00点はタイムアウトを逃れた数少ない速いポイントであり、
モデル。実行のほとんどが失敗した場合、中央値は生存者を表します。常にその横にある中央値を読んでください。
答えた数。
sub = Claude Code サブスクリプション (トークンごとの請求ではありません) · free = Ollama Cloud の無料枠。
サンプリングされた実行 ( n/3 、 n/20 ) は、同じグリッドの決定論的なサブセットである --sample を使用するため、すべての
モデルは同一の引数に基づいてスコア付けされます。
⚠ このハーネスが現在処理している 2 つの測定トラップ — どちらも困難な方法で見つかりました。
1. トークンは total_tokens に隠れる可能性があると考えます。ジェミニにサービスを提供するnear.aiが復帰
プロンプト 113 · 完了 18 · 合計 15126 — 支出の 99.9% はどちらにも当てはまりません
completed_tokens もreasoning_tokensも、合計のみ。読み取り完了トークン
(明らかな分野) ジェミニは約 800 倍過小報告されています: 22,401 ではなく 253 tok/pt、
それは完全に間違った集団に含まれています。ハーネスは再構築されます
hidden = 合計 − プロンプト − 完了をポイントごとに保存します。これをフォークした場合は、
ホストの算術計算は、トークン番号を信頼する前に合計されます。
2. すべてのモデルを温度 0 に固定できるわけではありません。GPT-5.x、o シリーズ、および Gemini は拒否します。
温度=0 なので、これらの実行では --no-temperature が使用され、プロバイダーのデフォルト (1.0) が継承されます。
彼らの中央値は安定しています — ジェミニは同じ 20 の議論を 2 回実行して 9.55 と 9.24 を獲得しました —
しかし、彼らの尻尾はそうではありません：サム

2 回の実行では、最悪の点は -13.51 と 6.14 でした。
これらのモデルの 1 回実行のテール数値を、プロパティではなく 1 回の描画として扱います。
1. サイズ、ベンダー、アーキテクチャからはほとんど何も予測できません
24B (3.46 桁) と 675B (4.29) でのミストラル ショートカット - 28 倍のパラメータ増加で購入可能
0.8桁で構造的には何も変わりません。 Google のスコアは密度 4.76、MoE としては 4.91 です。
プールサイドの小型モデルは、その大型モデルをわずかに上回っていますが、どちらもフルで上回っています。
101 点グリッド:
この分割はベンダー間だけでなく、ベンダー内部でも行われます。 Google は 5.23 桁を出荷しています
ショートカット (Gemini-2.5-pro) を Gemini-3.5-flash と並べると、中央値は 9.55 になりますが、
−13.51 の最悪点、つまり答えは 13 桁ずれています。
2. クリーンな世代結果: GPT-5.4 と GPT-5.5 の間で変更されました。
同じエンドポイント、同じプロンプト、同じ引数 - サービング構成が混乱することはありません。
4 回連続リリースすると、約 10 トークンと 3.5 ～ 4.4 桁でショートカットされます。 5番目は派生します。 A ~900×
単一のバージョン ステップ全体で 11 桁の労力と 11 桁にジャンプします。勾配や勾配ではありません。
サイズの話。
3. プロンプトの 1 文で、モデルが 0/101 桁から 15.00 桁に移動しました。
minimax-m3 のスコアは 101 点中 0 点でした。すべてのポイントが予算をすべて消費して空に戻りました。その
トレースはその理由を示しています。教科書の列の 24 桁のオペランドの算術です。
列 7: 5 + 1 +

[切り捨てられた]

## Original Extract

Contribute to DmitriGoloubentsev/LLM-Tests development by creating an account on GitHub.

GitHub - DmitriGoloubentsev/LLM-Tests · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
DmitriGoloubentsev
/
LLM-Tests
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
14 Commits 14 Commits docs docs test1 test1 test1v2 test1v2 .gitignore .gitignore README.md README.md make_hero_figure.py make_hero_figure.py run_openrouter_free.sh run_openrouter_free.sh View all files Repository files navigation
Can an LLM follow a derivation — or does it just reach for a calculator?
Ask any assistant for exp(-0.9513299347) to 15 digits and you get a perfect answer. It writes
three lines of Python and runs them. That measures the sandbox, not the model.
This is the same exam with the calculator taken away. No tools are available — not restricted, not
discouraged, absent — so the model has to do what a student does on a closed-book paper: reduce
the argument onto a constant it remembers, expand a Taylor series, and work through a dozen
16-digit multiplications by hand. Thousands of tokens of arithmetic in which one slipped carry
destroys the answer .
There is no way to bluff it. The reply is either right to fifteen places or it isn't, and the
grader can tell to a fraction of a digit.
We are not testing whether a model can produce the right number.
We are testing whether it can follow a long derivation without losing the thread.
With a tool, every model scores 100%. Without one, the spread is enormous.
100+ runs · 50+ models · 8 endpoints · one prompt · one stateless call per point.
Asked the identical question, models split into two populations with almost nothing in between:
A 1,300× difference in effort, and a 10-digit difference in the answer.
And which side a model lands on is not predicted by its size, its vendor, or its price.
Mistral shortcuts at 24B and at 675B. Poolside's smaller model beats its larger sibling.
And OpenAI's own ladder crosses the line in a single version step — GPT-4.1 through 5.4 all
shortcut at ~10 tokens; GPT-5.5 derives at 8,201.
Four runs, same function, same 101 points, same prompt. Top: one model on arguments it can recall,
then on arguments it cannot — six digits vanish. Bottom: a model that derives every value, and what
happens when you throttle its thinking.
Unfakeable. A confident paragraph earns nothing. Only correct digits score.
Gradable to a fraction of a digit , against a reference correct by construction.
Recall-proof on demand. Arguments are jittered to ten decimals by a seeded hash, so they
cannot appear in any training corpus. The model must derive, or guess.
It separates what every other eval conflates: knowing the method, and executing it
without error across thousands of tokens.
The metric is correct digits = -log10(relative error) , capped at 15 (IEEE-754 double agreement).
Pass/fail would hide everything interesting — "4.8 digits" and "15 digits" are both "wrong" to a
boolean grader.
How "no calculator" is enforced
Over a plain API this is free: no tools array is sent, so there is nothing to call. The risky path
is the Claude Code CLI backend, where the model normally has an interpreter. There it runs with:
--tools "" — the entire built-in tool set disabled,
a throwaway $HOME and empty working directory, so no CLAUDE.md reaches the context,
and every reply verified to have num_turns == 1 — a tool call needs a second turn, so a
single turn proves the answer came out of the model, not out of python .
grid
question it answers
test1
x = -1.00 … 1.00 step 0.02 (101 round arguments)
how accurate on values it may have memorised ?
test1v2
same 101 points, each nudged by a deterministic ~10-decimal offset ( -0.9513299347 , …)
how accurate when the argument cannot be in any corpus?
test1v2 supports --func exp|sin|cos|log|sqrt .
The ablation between them is the point. gemma4:31b scores 10.40 median on round arguments and
4.76 on recall-proof ones — same model, same prompt, same ~18 tokens per answer. Its test1 score
was a lookup table almost in full.
Its worst round-argument point is not a rounding error: asked for exp(0.58) it replied
1786038444035 — the decimal point is missing , so the answer is off by 10¹² while looking
perfectly well-formed.
Sorted into the four groups that actually behave differently. Read tok/pt alongside the median —
it says whether the model derived the answer or recalled it.
1 — Derivers · thousands of tokens, land on the double-precision line
model
fn
answered
median
worst
tok/pt
cost
openai/gpt-5.5
exp
101/101
15.00
3.08
8,201
—
openai/gpt-5.5
sin
20/20
15.00
12.18
9,344
—
openai/gpt-5.6-sol
exp
5/5
15.00
14.17
4,254
$0.64
anthropic/claude-opus-5
exp
10/10
15.00
13.38
16,688
$4.17
deepseek-v4-flash
exp
54/101
15.00
13.17
51,695
$1.46
deepseek-v4-flash effort=low
exp
98/101
15.00
5.04
26,520
$0.75
Claude Opus 5 (CLI) effort=max
exp
20/20
15.00
5.56
15,690
sub
z-ai/glm-5.2
exp
18/101
15.00
8.37
30,181
free
Claude Fable 5 (CLI)
exp
62/101
14.68
10.49
5,503
sub
moonshotai/kimi-k2.7-code
exp
7/20
15.00
12.14
55,391
free
Claude Opus 5 (CLI)
sin
101/101
14.38
8.32
7,924
sub
moonshotai/kimi-k2.6
sin
15/20
14.14
7.89
54,103
free
moonshotai/kimi-k2.6
exp
13/20
13.68
8.31
57,209
free
2 — The sparse middle · genuinely derive, still fall short
model
fn
answered
median
worst
tok/pt
anthropic/claude-sonnet-5
exp
10/10
12.10
9.30
22,630
openai/o3
exp
1/3
11.88
11.88
31,987
meta/muse-glimmer-30b
exp
3/3
11.46
9.82
27,817
deepseek-v4-flash:preview
exp
3/3
11.44
9.50
49,470
qwen/qwen3.7-max
exp
3/3
11.40
11.32
11,843
anthropic/claude-opus-4.8 effort=high
exp
10/10
11.13
9.20
4,141
openai/o4-mini
exp
3/3
10.96
9.82
15,557
thinkingmachines/inkling-small
exp
3/3
10.01
6.18
25,808
gpt-oss:120b
exp
98/101
9.55
3.18
9,105
google/gemini-3.5-flash
exp
20/20
9.24
6.14
22,401
poolside/laguna-xs-2.1
exp
61/101
7.60
3.08
25,599
gpt-oss:20b
exp
93/101
6.37
2.44
26,420
anthropic/claude-sonnet-4-6 (near.ai)
exp
3/3
6.33
4.76
1,292
qwen/qwen3.7-flash
exp
80/101
5.63
2.84
27,579
qwen3.5:397b
exp
101/101
5.19
3.95
6,015
nvidia/nemotron-3-ultra (Ollama)
exp
3/3
4.50
4.13
12,894
openai/o3-mini
exp
3/3
3.91
2.96
4,588
3 — Shortcutters · ~18 tokens, 3–7 digits, never fail and never improve
model
size
answered
median
tok/pt
google/gemini-2.5-pro
—
3/3
5.23
22
google/gemma-4-26b-a4b (MoE)
26B
3/3
4.91
19
gemma4:31b (dense)
31B
101/101
4.76
18
anthropic/claude-opus-4.8 (default)
—
10/10
4.66
9
upstage/solar-pro4
—
101/101
4.41
18
openai/gpt-5.4
—
3/3
4.38
11
mistral-large-3:675b
675B
101/101
4.29
18
openai/gpt-5.2
—
3/3
4.22
10
poolside/laguna-s-2.1
—
20/101
4.13
18 (when it answers)
openai/gpt-5.1
—
3/3
4.03
16
qwen3-coder
~30B
101/101
3.72
18
openai/gpt-4.1
—
3/3
3.57
7
devstral:24b
24B
101/101
3.46
16
openai/gpt-5.4-mini
—
3/3
3.40
16
qwen2.5-coder:14b
14B
101/101
2.93
18
4 — Never finish · reason to the output cap, return empty content
model
answered
tok/pt
endpoint
minimax-m3
0/101
32,000 (cap)
Ollama Cloud
minimax-m2.7
0/20
65,536 (cap)
Ollama Cloud
inclusionai/ling-3.0-tiny
0/101
27,564
OpenRouter
inclusionai/ling-3.0-flash
0/3
32,000 (cap)
OpenRouter
cohere/north-mini-code
0/3
26,176
OpenRouter
deepseek-v4-pro
0/3
65,536 (cap)
Ollama Cloud
glm-5.1
0/3
32,768 (cap)
Ollama Cloud
claude-opus-4.6 effort=high , uncapped
0/10
65,536 (cap)
OpenRouter
And one category of its own — the Nemotron family is broken on this task at every size.
nemotron-3-nano-30b scored a median of 0.00 correct digits over 101 points (one reply was
229144000.0 where the answer is 0.89 ); nemotron-3-ultra-550b managed 3.69 (4.50 on Ollama);
the reasoning-tuned nemotron-3-nano-omni-30b scored −0.00 after burning 32,768 tokens per
point; and nemotron-3-super-120b answered 1/20 at 6.32 on a working endpoint. Four models,
four sizes, three endpoints, one result.
A survivorship warning, from this family. An early nemotron-3-super-120b run on NVIDIA NIM
scored a 15.00 median — a deriver-grade number — from 8 of 101 points, with 93 lost to
HTTP 504 , at 783 tok/pt. Re-run on OpenRouter's free tier it answers 1/20 at 6.32 digits and
6,924 tok/pt . The 15.00 was the handful of fast points that escaped the timeouts, not the
model. When most of a run fails, the median describes the survivors — always read it next to
the answered count.
sub = Claude Code subscription (not a per-token bill) · free = Ollama Cloud free tier.
Sampled runs ( n/3 , n/20 ) use --sample , a deterministic subset of the same grid, so every
model is scored on identical arguments.
⚠ Two measurement traps this harness now handles — both found the hard way.
1. Thinking tokens can hide in total_tokens . near.ai serving Gemini returns
prompt 113 · completion 18 · total 15126 — 99.9% of the spend is in neither
completion_tokens nor reasoning_tokens , only in the total. Reading completion_tokens
(the obvious field) under-reported Gemini by ~800× : 253 tok/pt instead of 22,401, which put
it in the wrong population entirely. The harness now reconstructs
hidden = total − prompt − completion and stores it per point. If you fork this, check your
host's arithmetic adds up before trusting any token number.
2. Not every model can be pinned to temperature 0. GPT-5.x, the o-series and Gemini reject
temperature=0 , so those runs use --no-temperature and inherit the provider default (1.0).
Their medians are stable — Gemini scored 9.55 and 9.24 on two runs of the same 20 arguments —
but their tails are not : the same two runs put its worst point at −13.51 and 6.14 .
Treat single-run tail figures for those models as one draw, not a property.
1. Size, vendor and architecture predict almost nothing
Mistral shortcuts at 24B (3.46 digits) and at 675B (4.29) — a 28× parameter increase buys
0.8 digits and changes nothing structurally. Google scores 4.76 dense and 4.91 as an MoE.
Poolside's smaller model beats its larger sibling , and not marginally — both on full
101-point grids:
The split also runs inside vendors, not just between them. Google ships a 5.23-digit
shortcutter (Gemini-2.5-pro) alongside Gemini-3.5-flash, which derives to a 9.55 median — but with
a −13.51 worst point, i.e. an answer off by thirteen orders of magnitude.
2. The clean generational result: it changed between GPT-5.4 and GPT-5.5
Same endpoint, same prompt, same arguments — no serving-config confound:
Four consecutive releases shortcut at ~10 tokens and 3.5–4.4 digits. The fifth derives. A ~900×
jump in effort and eleven digits, across a single version step — not a gradient, and not a
size story.
3. One sentence of prompting moved a model from 0/101 to 15.00 digits
minimax-m3 scored 0 of 101 — every point burning the full budget and returning empty. Its
traces show why: schoolbook column arithmetic on 24-digit operands.
Col 7: 5 + 1 +

[truncated]
