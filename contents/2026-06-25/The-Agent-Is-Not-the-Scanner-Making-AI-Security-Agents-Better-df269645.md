---
source: "https://shad0wmazt3r.github.io/ai-security"
hn_url: "https://news.ycombinator.com/item?id=48674765"
title: "The Agent Is Not the Scanner: Making AI Security Agents Better"
article_title: "The Agent Is Not the Scanner: Making AI Security Agents Better – Beri's Blog – Security Research"
author: "speckx"
captured_at: "2026-06-25T15:54:24Z"
capture_tool: "hn-digest"
hn_id: 48674765
score: 3
comments: 0
posted_at: "2026-06-25T15:21:22Z"
tags:
  - hacker-news
  - translated
---

# The Agent Is Not the Scanner: Making AI Security Agents Better

- HN: [48674765](https://news.ycombinator.com/item?id=48674765)
- Source: [shad0wmazt3r.github.io](https://shad0wmazt3r.github.io/ai-security)
- Score: 3
- Comments: 0
- Posted: 2026-06-25T15:21:22Z

## Translation

タイトル: エージェントはスキャナーではありません: AI セキュリティ エージェントをより良くする
記事のタイトル: エージェントはスキャナーではありません: AI セキュリティ エージェントをより良くする – Beri のブログ – セキュリティ リサーチ
説明: LLM は、過去 1 年間で脆弱性を発見する能力が驚くほど向上しました。これらは、より優れた AI 支援セキュリティ ワークフローの構築と、モデルごとにスキャナーやスキルによる影響がどのように異なるかについての私のメモです。

記事本文:
PB
ベリのブログ
セキュリティ研究
ブログ
について
アーカイブ
2026 年 5 月 1 日
エージェントはスキャナーではありません: AI セキュリティ エージェントをより優れたものにする
LLM は、過去 1 年間で脆弱性を発見する能力が驚くほど向上しました。これらは、より優れた AI 支援セキュリティ ワークフローの構築と、モデルごとにスキャナーやスキルによる影響がどのように異なるかについての私のメモです。
LLM を足場にすることは、一般的に良いアイデアではありません。それが役立つか害があるかは、ほぼ完全にモデルがすでにどの程度の能力を持っているかに依存し、モデル ファミリによって異なります。
私は 8 か月かけてセキュリティ タスクに対して LLM を実行して、実際に何が機能するかを把握し、経験的に測定しました。結果は私が期待していたものではありませんでした。
エージェントにすべてのコンテキストを渡すのは、一見すると良いアイデアのように思えます。
しかし、このアプローチにはいくつかの問題があります。 1つ目はコストです。スキャナーで簡単にキャッチできるものになぜ 20,000 ～ 30,000 トークンを費やすのでしょうか?
セキュリティ エンジニアはスキャナーの作成に精力的に取り組んできており、当社には SAST と DAST に対する効果的な技術があります。さらに、nmap および semgrep 出力の出力を解析することもトークンの無駄です。 LLM は、アプリケーションのこれまで調査されていない領域を調査する方向に導くために、すでに調査された内容に関する最も洗練された結果のみを必要とします。
なぜブラインドスキャナーでも十分ではなかったのか
考慮していない脆弱性クラスが常に存在すること、またはコンテキストの欠如によりコードでの特定の発現を予期していなかった可能性があることを考慮すると、すべての単一の脆弱性クラス/CTF 課題の検出ロジックを記述することは、根本的に解決不可能な問題です。スキャンの実行は簡単で、ターゲットに対して「実行」を押して成功/失敗の出力を待つだけですが、何を検出するか、どこに脆弱性があるかをどうやって知ることができるのでしょうか。

?
スキャナーとファザーのさらに深刻な問題は、スキャナーとファザーが本質的に決定論的であり、既知の署名とのパターン マッチングのみを行うため、探すべきものを見つけることには非常に優れていますが、その他のことには盲目であることです。 SQLi、XSS、IDOR から完全に保護されているかもしれませんが、SSTI の検出を一度も作成したことがないと、それらを完全に見逃してしまう可能性があります。
ラティスマインドの構築と足場
このソリューションは、確定的で安価なスキャンと、構造化された知識および脆弱性を特定する LLM のインテリジェンスを組み合わせたハイブリッド アプローチです。このハイブリッド アプローチの最初の部分は、Lattice Mind です。
Lattice Mind は、エージェントがスキャンを実行できるスキャナであり、脆弱性のクラス全体に対応するスクリプトを作成することなく、現在のスキャンを編集して独自のペイロードを入力できます。基本的に、決定論的推論 (つまりデシジョン ツリー) を使用してアプリのアーキテクチャについて推論することで、簡単に実現できる成果を掴みます。
たとえば、Lattice Mind のスキャンによってエージェントにペイロードを書き込むのに十分な情報が得られた場合、エージェントは Lattice Mind を使用して新しいペイロードでスキャンを実行し、前述の SSTI 脆弱性を検証/悪用するためのスクリプトを作成するオーバーヘッドを節約できます。
スキャフォールディングは、セキュリティ タスクの実行に必要なツールとコンテキストへのアクセスを LLM に提供するスキルと MCP サーバーを備えたオーケストレーション レイヤーです。当初、スキルは TTP に似ていました。何を検出するのか、そこからどこへ行くのか、何を連鎖させることができるのかなどです。明らかに、スキル向上スキルを追加する必要がありました。このスキルでは、タスク完了時にモデルがスキルと MCP が実行にどのように役立った/悪影響を与えたかを評価し、次回の実行でより良いものになるようにスキルを修正する必要がありました。
ライブ CTF でテストしたところ、PicoCTF 2025 と DawgCTF 2026 の最も難しい問題を 30 分以内に解決できました。

それぞれ1分。
私にとって、このようなセットアップは、ツールやコンテキストを持たないモデルを提供するよりも優れていることは明らかであるように思えました。私の仮説は、これらのモデルはセキュリティ エンジニアリング タスクよりもソフトウェア エンジニアリング タスク向けにトレーニングされており、スキルとしての TTP がそのギャップを埋めることができるというものでした。私は自分の直感をテストしたかったので、テストベンチをセットアップしました。11 個のモデル、それぞれ 3 回の実行、コントロール vs スキルのみ vs MCP 対応、20 個の脆弱性発見コード スニペット タスク。
すべてのモデルはコード スニペットを調べて脆弱性を見つけようとします。私は次の 2 つの面で評価しました。
手動: 最大/最もサイバー能力の高いモデルが推論を検討し、回答を手動で採点して、推論が正しかったかどうか、モデルが解決を妨げる技術的問題に悩まされていないか、別の CVE を指定していないか、同じスニペット内で別のクラスの脆弱性を検出したかどうかを確認します。この点は、モデルが不当に評価されていないことを保証するためでした。
自動化: これらの脆弱性スニペットには特定の「正しい」答えがあり、モデルがその答えにどれだけ近づいたかを確認することで評価できます。
結果は期待したものではありませんでした
スキルにより、最も弱いモデルが大幅に改善されました。 gpt-5.1-codex-mini-low では、F1 が 0.4774 から 0.5926 に跳ね上がり、相対的に 24% 向上し、厳密なタスクの精度が 20 パーセントポイント上昇しました。残りの低ベースライン グループは、+0.05 ～ +0.06 F1 の範囲で同じパターンを示しました。制御 F1 が 0.60 未満の 4 つのモデルすべてで、スキルは平均 ΔF1 +0.0656 を生成し、タスクあたりの FP が 0.08 低下し、厳密なタスクの精度が 10.83 ポイント向上しました。
特にコーデックスミニローリフトは、数学を変えるような数字です。 codex-mini は、100 万入出力トークンあたり 0.25 ドル / 2.00 ドルで、codex-mini よりも入力で約 12 倍、出力で 7.5 倍安くなります。

ソネット。スキルのおかげでF1は4分の1近く上昇した。すべてのタスクでフロンティア モデルを実行するのはコストが高すぎるエージェント セキュリティ作業では、スキャフォールディングが実際のコスト レバーになります。
問題は、ゲインが一様ではないということであり、一様でないということ自体が、より興味深い発見です。
11 モデルすべてにおけるベースライン F1 とスキル向上率の相関関係は -0.81 であり、ベースライン パフォーマンスだけから効果の兆候を予測できるほど十分に強いです。
制御 F1 が 0.60 未満のモデル (codex-mini-low、より低いナノ バリアント): 平均 ΔF1 +0.0656、タスクあたりの FP は 0.08 減少、厳密な精度は 10.83 ポイント増加。
制御 F1 が 0.75 以上のモデル (ミニ高、ミニ中、ソネット中): 平均 ΔF1 -0.0382、タスクあたりの FP は 0.06 増加、厳密な精度は 6.67 ポイント減少。
モデル別ΔF1（スキル - コントロール）
+0.12 | gpt-5.1-codex-mini-low ████████████
+0.06 | gpt-5.4-ナノ-なし ██████
+0.05 | gpt-5.4-ナノ-ロー █████
+0.03 | gpt-5.4-ナノハイ ███
+0.02 | gpt-5.4-ミニロー██
+0.02 |ジェミニ-3-フラッシュ ██
-0.02 | kimi-k2.5 ░░
-0.02 |クロード-4.6-ソネット-ミディアム ░░
-0.03 | gpt-5.4-ミニハイ ░░░
-0.06 | gpt-5.4-ナノ-ミディアム ░░░░░░
-0.06 | gpt-5.4-ミニ-ミディアム ░░░░░░
機構は両端で対称です。弱いモデルは、以前は所見を省略したり、出力構造を混乱させたり、実体のない形状を幻覚したりするため、足場から価値を引き出します。
スキルは、何を探すべきか、何が証拠として考慮されるか、そして答えを導き出すために必要なフォーマットの明確なリストを与えてくれます。即興の余地が少ないため、リコールが上昇し、FP が低下します。
強力なモデルには、そのすべてがすでに内部に組み込まれています。構造化された命令を上に重ねても、新しい情報は追加されません。

オーバーヘッドのみ。
3 つのモデル固有の動作は勾配に適合しなかったため、取り出す価値があります。
Gemini 3 Flash は、単一検出の精度が向上しました (+0.0184 F1、FP レートが低下) が、悪用連鎖の成功率は 66.7% から 0% に低下しました。スキルは個々の呼び出しを鮮明にし、複数段階の推論を打ち破りました。タスクで結果を連鎖させる必要がある場合は、Gemini を制御して実行します。一発の精度が必要な場合は、スキルが勝利します。
キミ K2.5 は客観的には後退し (-0.0156 F1)、主観的には改善しました (+0.0175)。報道の印象は良くなるが、脆弱性差別は悪化する。トリアージモデルではなく、書き込みモデルとして役立ちます。
クロード・ソネット 4.6 は最も極端な分裂を示しました。客観的な F1 は -0.0178 低下、主観的な品質は +0.2666 上昇し、実行中の主観的な最大のゲインです。スキルにより、ソネットのレポートはより明確になり、結果はわずかに悪化しました。
MCP はスキルだけで勝ったわけではありません。コントロールにも勝てませんでした。
私は、スキルの上に MCP ツールを重ねてパフォーマンスを向上させることを期待していました。ツールは指示よりも優れている必要があります。彼らはそうしませんでした。
MCP 対応は両方のベースラインを下回りました。 11 モデル全体で、MCP は 3 でコントロールを上回り、1 でスキルのみを上回りました。ソネットは、コントロールに対して -0.2021 F1 で最悪の打撃を受けました。 Codex-mini-low と nano-high も -0.09 以上下落しました。
MCP の実行はモデルごとに 1 回の実行ですが、他の 2 つのプロファイルは 3 回の実行の平均であるため、分散の同等性を主張できません。しかし、その差は十分に大きく、11 モデル全体で十分に一貫しているため、1 回の実行ノイズがすべてである可能性は低いです。
可能性の高い説明は、ベンチマーク自体が間違ったテスト面であるということです。これについては次に説明したいと思います。
ベンチマークは明確な物語を物語っていました。 MCP のせいで事態はさらに悪化しました。その話はこのベンチマークにも当てはまり、ベンチマークは私が実際に関心を持っている質問に対する間違った手段です。
Th

これらのタスクは静的なコード スニペットです。エージェントはコードの断片を受け取り、脆弱性クラスに名前を付けるように求められます。
リクエストに応答するサービスはなく、逆アセンブルするバイナリも、実行するリポジトリも、駆動するブラウザもありません。エージェントがツールを使用して実際に実行できる一連の操作は空です。 MCP をそのコンテキストにロードすることは純粋なオーバーヘッドです。
kali-mcp-server、lattice-mind、ghidra-mcp、および github-mcp のツール定義はトークンを消費し、タスクが使用できない機能をアドバタイズし、見返りのないツール選択にモデルを推論するよう促します。
データのパターンはその読み取り値と一致します。 MCP を有効にして最も損失が大きかったモデルは、ツール呼び出しを試行すべきではないときにツール呼び出しを試行する可能性が最も高いモデルでもあります。Sonnet は -0.2021 F1、codex-mini-low は -0.0928、nano-high は -0.1012 です。
アクションに対するモデルのバイアスが強いほど、間違った環境によるペナルティは大きくなります。これは、実際のターゲットで見られるものと一致しません。 CTF チャレンジの実行時、エージェントが問題を解決するのは、同じ MCP 設定が原因です。 Web チャレンジでは、劇作家がアプリをウォークスルーするよう求められます。
ラティスマインド スキャンはほぼ実行ごとに行われ、ライブ エンドポイントに対してペイロードを起動します。ツールは、プロンプトに存在しないシグナルを提供します。
最大の変化は、すべてのモデルに対して同じ足場を実行するのをやめたことです。ハーネスは、モデルが得意なことと、ワークフローの段階で実際に必要なことによってルーティングされるようになりました。
2 つのルールが他のすべてに優先します。
強力なモデルは少なくなります。モデルが 0.75 コントロール F1 を超える場合、ハーネスはデフォルトでスキル ライトまたは純粋なコントロールになります。重い足場により、高ベースライン グループの F1 コストは平均 0.04 となり、タスクあたりの FP は 0.06 上昇しました。スキルテキストは、これらのモデルにはない問題を解決していました。
弱いモデルの方が多くなります。モデルが 0.60 コントロール F1 を下回った場合、ハーンは

ess のデフォルトはフルスキルです。スキルは codex-mini-low +0.115 F1 と +0.05 から +0.06 の間のナノ層を購入しました。リコールは上昇し、FP は低下し、出力構造はコイン投げではなくなります。
これはモデル強度軸を処理します。パイプラインステージの軸は、実質的な勝利のほとんどをもたらしたものです。
Reconはスキルのある安価なモデルに進みます。 Nano-low のコストは、100 万出力トークンあたり約 0.40 ドルで、mini の 5 倍、Sonnet のほぼ 40 倍安いです。スキルを有効にすると、F1 リフトが +0.05 増加し、FP が減少します。実行コストが制約であり、失われたもののコストが下流で回収できる広範囲の表面のスイープの場合、これは適切なスロットです。私はそれを広範囲に実行し、後で検証パスを利用して生成される FP を捕捉します。
信頼性の高いエクスプロイト推論は、制御に関する強力なモデルに基づいています。コントロールのミニハイは、ベンチマークで最高の生の目標 F1 である 0.8384 を記録しました。また、高ベースラインの回帰が最も深刻な影響を与えるため、スキルを重ねるには最悪の場所でもあります。ハーネスが候補結果の小さなセットに絞り込まれ、それぞれについて最高の識別コールが必要になると、足場を使用しないミニハイがデータ内で最もクリーンな答えになります。
最終報告はスキルを持つソネットへ。 Sonnet のベンチマーク プロファイルが明確に役立つのは、ここだけです。目標 F1 は 0.017 低下しました

[切り捨てられた]

## Original Extract

LLMs have gotten surprisingly good at finding vulnerabilities over the past year. These are my notes on building better AI-assisted security workflows, and how different models are impacted by scanners and skills differently.

PB
Beri's Blog
Security Research
Blog
About
Archive
May 1, 2026
The Agent Is Not the Scanner: Making AI Security Agents Better
LLMs have gotten surprisingly good at finding vulnerabilities over the past year. These are my notes on building better AI-assisted security workflows, and how different models are impacted by scanners and skills differently.
Scaffolding an LLM is not a universally good idea. Whether it helps or hurts depends almost entirely on how capable the model already is, and it varies between model families.
I spent eight months running LLMs against security tasks to figure out what actually works and empirically measuring them. The findings were not what I expected.
Handing an agent all the context seems to be a good idea at first glance.
But that approach has several issues. The first one of which is cost. Why would you spend 20-30K tokens on something that could be easily caught by a scanner?
Security engineers have worked tirelessly to create scanners and we have effective techniques for SAST and DAST. Moreover, parsing the output for nmap and semgrep outputs is a waste of tokens too. The LLM only needs the most refined results on what has already been looked at, what the results were to steer it towards looking into previously unexplored areas of the application.
Why Blind Scanners Were Not Enough Either
Writing detection logic for every single vulnerability class/CTF challenge is a fundamentally unsolvable problem considering there would always be vulnerability classes you didn’t consider or you didn’t expect the specific manifestation in code due to lack of context. Running scans is easy, you can just hit run against a target and wait for the success/fail output but how do you know what to detect and where your vulnerabilities lie?
The deeper problem with scanners and fuzzers are that they are inherently deterministic and only pattern match against known signatures, making them extremely good at finding what to look for but blind to everything else. You may have fully protected yourself against SQLi, XSS, and IDOR, but if you never write a detection for SSTIs, you may miss them entirely.
Building Lattice Mind and The Scaffolding
The solution is a hybrid approach combining the deterministic, inexpensive scans with structured knowledge and LLM’s intelligence to identify vulnerabilites. The first part of this hybrid approach is Lattice Mind.
Lattice Mind is a scanner that the agent can run scans, and can edit the current scan to input its own payload without having to create a script for a whole class of vulnerabilities. Essentially, you grab the low hanging fruits by using deterministic reasoning (i.e. decision trees) to reason about the app architecture,
For example, if Lattice Mind’s scans give the agent enough information to write a payload, it can use Lattice Mind to run a scan with the new payload saving the overhead of it having to write a script to verify/exploit the said SSTI vulnerability.
The Scaffolding is the orchestration layer with skills and MCP servers providing the LLM with access to tools and context needed to perform their security tasks. At the start, the skills looked more like TTPs: what do you detect, where do you go from there, what can you chain etc. Obviously, I had to add a skill improvement skill where the model upon completing a task, evaluates how skills and MCP helped/harmed the run, and then modify the skills to be better next run.
I put it to the test in live CTFs and it was able to solve the hardest problems from PicoCTF 2025 and DawgCTF 2026 in less than 30 minutes each.
To me, it seemed obvious that a setup like this would be better than providing the model with no tools and context. My assumption was that these models are trained more for software engineering tasks than security engineering tasks, and TTPs as skills could fill in the gap. I wanted to test my intuition and hence, set up a test bench: 11 models, 3 runs each, control vs skills-only vs MCP-enabled, 20 vulnerability-finding code snippet tasks.
Every model would look at code snippets and try to find vulnerabilites, and I evaluated on two fronts:
Manual: The largest/most cyber capable model looks through the reasoning and manually grades the answers to check if the reasoning was correct and whether the model suffered from technical issues hindering its solve or whether it named a different CVE, or if it detected a different class of vulnerability in the same snippet. This front was to ensure we’re not unfairly assessing models.
Automated: These vulnerability snippets have certain “correct” answers and we can evaluate by checking how close the models got to the answer.
The Results Were Not What I Expected
Skills made the weakest models substantially better. On gpt-5.1-codex-mini-low, F1 jumped from 0.4774 to 0.5926, a 24% relative gain, with strict task accuracy climbing 20 percentage points. The rest of the low-baseline group showed the same pattern in the +0.05 to +0.06 F1 range. Across all four models with control F1 below 0.60, skills produced an average ΔF1 of +0.0656 , dropped FPs per task by 0.08, and lifted strict task accuracy by 10.83 points.
The codex-mini-low lift in particular is the kind of number that changes the math. At $0.25 / $2.00 per million input/output tokens, codex-mini is roughly 12x cheaper on input and 7.5x cheaper on output than Sonnet. Skills pushed its F1 up by nearly a quarter. For agentic security work where running a frontier model on every task is too expensive, scaffolding turns into a real cost lever.
The catch is that the gain is not uniform, and the way it isn’t uniform is itself the more interesting finding.
The correlation between baseline F1 and skill uplift across all 11 models is -0.81 , strong enough that you can predict the sign of the effect from baseline performance alone:
Models with control F1 below 0.60 (codex-mini-low, the lower nano variants): average ΔF1 +0.0656 , FPs per task down 0.08 , strict accuracy up 10.83 points .
Models with control F1 at or above 0.75 (mini-high, mini-medium, sonnet-medium): average ΔF1 -0.0382 , FPs per task up 0.06 , strict accuracy down 6.67 points .
ΔF1 by model (skills - control)
+0.12 | gpt-5.1-codex-mini-low ████████████
+0.06 | gpt-5.4-nano-none ██████
+0.05 | gpt-5.4-nano-low █████
+0.03 | gpt-5.4-nano-high ███
+0.02 | gpt-5.4-mini-low ██
+0.02 | gemini-3-flash ██
-0.02 | kimi-k2.5 ░░
-0.02 | claude-4.6-sonnet-medium ░░
-0.03 | gpt-5.4-mini-high ░░░
-0.06 | gpt-5.4-nano-medium ░░░░░░
-0.06 | gpt-5.4-mini-medium ░░░░░░
The mechanism is symmetric on both ends. Weaker models get value out of scaffolding because they were previously omitting findings, getting confused on output structure, or hallucinating shape without substance.
Skills give them an explicit list of what to look for, what counts as evidence, and the format their answer needs to land in. Recall climbs and FPs drop because there is less room to improvise.
Stronger models already have all of that internally. Layering structured instructions on top adds no new information, only overhead.
Three model-specific behaviors didn’t fit the gradient and are worth pulling out.
Gemini 3 Flash improved on single-finding precision (+0.0184 F1, lower FP rate) but its chain-of-exploit success collapsed from 66.7% to 0%. Skills sharpened individual calls and broke multi-step reasoning. If the task requires chaining findings, run Gemini on control. If it requires single-shot precision, skills win.
Kimi K2.5 regressed objectively (-0.0156 F1) and improved subjectively (+0.0175). Better-sounding reports, worse vulnerability discrimination. Useful as a writeup model, not a triage one.
Claude Sonnet 4.6 showed the most extreme split. Objective F1 down -0.0178, subjective quality up +0.2666 , the largest subjective gain anywhere in the run. Skills made Sonnet’s reports much clearer and its findings slightly worse.
MCP didn’t beat skills-only. It didn’t beat control either.
I expected MCP tools layered on top of skills to push performance higher. Tools should beat instructions. They didn’t.
MCP-enabled was below both baselines. Across 11 models, MCP beat control on 3 and beat skills-only on 1. Sonnet took the worst hit at -0.2021 F1 versus control. Codex-mini-low and nano-high also dropped by more than -0.09.
The MCP run is single-run per model, while the other two profiles are 3-run means, so I cannot claim variance equivalence. But the gap is large enough and consistent enough across 11 models that single-run noise is unlikely to be the whole story.
The more likely explanation is that the benchmark itself is the wrong test surface, which is what I want to talk about next.
The benchmark told a clear story. MCP made things worse. That story is true on this benchmark, and the benchmark is the wrong instrument for the question I actually care about.
These tasks are static code snippets. The agent receives a code fragment and is asked to name the vulnerability class.
No service responding to requests, no binary to disassemble, no repository to walk, no browser to drive. The set of things the agent can actually do with a tool is empty. Loading MCP into that context is pure overhead.
Tool definitions for kali-mcp-server, lattice-mind, ghidra-mcp, and github-mcp consume tokens, advertise capabilities the task cannot use, and invite the model to spend reasoning on tool selection that has no payoff.
The pattern in the data fits that read. The models that lost the most with MCP enabled are also the models most likely to attempt tool calls when they shouldn’t have: Sonnet at -0.2021 F1, codex-mini-low at -0.0928, nano-high at -0.1012.
The stronger the model’s bias toward action, the larger the wrong-environment penalty.This does not match what I see on live targets. On running CTF challenges, the same MCP setup is the reason the agent solves problems. On a web challenge it reaches for playwright to walk the app.
Lattice-mind scans happen nearly every run, firing payloads against a live endpoint. The tools provide signal that does not exist in the prompt.
The biggest change: I stopped running the same scaffold against every model. The harness now routes by what the model is good at and what the stage of the workflow actually needs.
Two rules sit on top of everything else.
Strong models get less. If a model lands above 0.75 control F1, the harness defaults to skills-lite or pure control. Heavy scaffolding cost the high-baseline group an average of 0.04 F1 and pushed FP per task up by 0.06. The skill text was solving a problem these models didn’t have.
Weak models get more. If a model lands below 0.60 control F1, the harness defaults to full skills. Skills bought codex-mini-low +0.115 F1 and the nano tier between +0.05 and +0.06. Recall climbs, FPs drop, and output structure stops being a coin flip.
That handles the model-strength axis. The pipeline-stage axis is where most of the practical wins came from.
Recon goes to cheap models with skills. Nano-low costs roughly $0.40 per million output tokens, which is five times cheaper than mini and almost forty times cheaper than Sonnet. With skills enabled it gets a +0.05 F1 lift and reduces FPs. For broad surface sweeps where the cost of running it is the constraint and the cost of missing things is recoverable downstream, this is the right slot. I run it wide and rely on a verifier pass later to catch the FPs it produces.
High-confidence exploit reasoning goes to strong models on control. Mini-high in control posted the best raw objective F1 in the benchmark at 0.8384. It is also the worst place to layer skills, because the high-baseline regression hits hardest there. Once the harness has narrowed to a small set of candidate findings and needs the highest-discrimination call on each one, mini-high without scaffolding is the cleanest answer in the data.
Final reports go to Sonnet with skills. This is the only place where Sonnet’s benchmark profile is unambiguously useful. Objective F1 dropped 0.017

[truncated]
