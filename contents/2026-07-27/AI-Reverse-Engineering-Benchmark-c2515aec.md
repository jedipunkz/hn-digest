---
source: "https://www.agentre-bench.ai/"
hn_url: "https://news.ycombinator.com/item?id=49068451"
title: "AI Reverse Engineering Benchmark"
article_title: "AgentRE-Bench — LLM Reverse Engineering Benchmark"
author: "childintime"
captured_at: "2026-07-27T12:51:30Z"
capture_tool: "hn-digest"
hn_id: 49068451
score: 1
comments: 1
posted_at: "2026-07-27T12:07:23Z"
tags:
  - hacker-news
  - translated
---

# AI Reverse Engineering Benchmark

- HN: [49068451](https://news.ycombinator.com/item?id=49068451)
- Source: [www.agentre-bench.ai](https://www.agentre-bench.ai/)
- Score: 1
- Comments: 1
- Posted: 2026-07-27T12:07:23Z

## Translation

タイトル: AI リバース エンジニアリング ベンチマーク
記事のタイトル: AgentRE-Bench — LLM リバース エンジニアリング ベンチマーク
説明: AgentRE-Bench は、ソース コードなし、ヒントなし、決定論的なスコアリングを使用して、コンパイルされたバイナリに対して AI リバース エンジニアリング エージェントを評価します。プライベート評価とトレーニング環境のパイロットが利用可能です。

記事本文:
AgentREベンチ
リーダーボード
タスク
プライベート評価
GitHub
お問い合わせ
ライト
AIセキュリティベンチマークプラットフォーム
AI エージェントはコードを書くことができます。リバースエンジニアリングできるでしょうか?
AgentRE-Bench は、ソース コードなし、ヒントなし、決定論的なスコアリングを使用して、コンパイルされたリバース エンジニアリング タスクで LLM エージェントを評価します。
公開ベンチマークは信頼性レイヤーです。オープンなタスク、再現可能な実行、AI リバース エンジニアリング エージェントのリーダーボードです。
非公開の評価とトレーニング環境のパイロットでは、エージェントがもっともらしいレポートを作成するのではなく、未リリースのバイナリに一般化するかどうかをテストします。
GitHub を見る
ベンチマークのスナップショット
V2 Linux の結果
13 個の ELF リバース エンジニアリング タスクが評価されました
パブリック ラダーの 10 個の Windows PE タスク
6 つのフロンティア モデルがリーダーボードで稼働
LLM 審査員は 0 人。決定的スコアリングのみ
01 ジェミニ 3.1 フラッシュライト 0.667
02 ディープシーク V4 プロ 0.648
03 クロード オーパス 4.7 0.512
引用者 / 研究の言及
AI セキュリティ研究で参照
CrackMeBench: エージェント向けのバイナリ リバース エンジニアリング
マルウェアのようなプロトコルとインフラストラクチャの再構築に重点を置き、ストリップされた ELF リバース エンジニアリング タスクに最も近いベンチマークとして AgentRE-Bench について言及しています。
自動脆弱性発見のチュートリアルと調査
自動化された脆弱性検出と AI セキュリティ ベンチマークの分野で AgentRE-Bench を引用します。
フロンティア モデルのパフォーマンス (Linux レベル)
13 の Linux ELF レベルすべて、25 のツール呼び出しバジェット、Docker サンドボックス静的分析ツール。 6 つのフロンティア モデルをエンドツーエンドで評価しました。見出しの発見: 小型の非思考モデル (Gemini 3.1 Flash Lite) がこの分野をリードし、メイン スコアですべてのフロンティア推論モデルを上回りました。このベンチでは、深さを推論することではなく、幻覚の調整が主な軸となっています。 Windows PE の結果は近日公開予定です。
幻覚が最も少ないネズミのベンチリーダー

フィールドで。ここでの勝利条件は、生々しい推論の深さではなく、調整です。
追加の 2 つのモデル (Gemini 3.1 Pro Preview および GLM 5.1) は、評価中の API エラーによりリーダーボードから除外されました。詳細とモデルごとの詳細については、分析の記事を参照してください。
13 の ELF レベルと 10 の Windows PE レベル
平文の TCP シェルから、WannaCry にインスピレーションを得た合成ランサムウェア ワームまで。 13 の Linux ELF レベルで基礎を学びます。 10 の Windows PE レベルでは、プロセス インジェクション、syscall 操作、およびワームの伝播が追加されます。
キャリブレーションは推論の深さに勝ります
バイナリ リバース エンジニアリングは、サポートされていない主張を罰します。ベンチマークでは、専門家のグラウンド トゥルースに基づいてスコアを付け、捏造されたテクニックを減算することで、それを可視化します。
キャリブレーションは推論の深さに勝ります
推論モデルは、バイナリから証明できないテクニックを過剰に主張することがよくありました。より低い幻覚剤の方がスコアが高かった。
バイナリ RE がもっともらしい報告を罰する
自信を持ってマルウェアを報告するだけでは十分ではありません。 AgentRE-Bench は、回復された事実に報酬を与え、裏付けのない主張には罰則を与えます。
誤解を招くアーティファクトがエージェントを騙す
タスクには、おとり文字列、誤解を招くラベル、逆アセンブリまたは実行を通じて検証する必要がある動作が含まれます。
より困難なタスクでは、ツールの継続的な使用、証拠の統合、および複数の分析ステップにわたる検証が必要です。
エージェントは検証済みの動作をバイナリから復元できますか?
AgentRE-Bench は、現実的なツール使用条件下でソースフリーのリバース エンジニアリングを測定します。エージェントはコンパイルされたアーティファクトを受け取り、サンドボックスで検査し、回復された事実、サポートされていない主張、およびエンドツーエンドの技術的正確さについて、隠されたグラウンドトゥルースに対してスコア付けされます。
ソースフリーのバイナリ指向
ソース コードを使用せずに、ファイル形式、アーキテクチャ、セクション、インポート、文字列、シンボル、および初回実行の手がかりを特定します。
エンドポイントを抽出し、

キー、定数、エンコードされた文字列、構成フィールド、プロトコル値、および動作固有のインジケーター。
ブランチ、ステート マシン、ステージングされたペイロード、検証ロジック、ランタイム デコード、および動的遷移を追跡します。
欺瞞と難読化への耐性
誤解を招く文字列、偽の暗号ラベル、おとり、デバッグ防止、分析防止チェック、ノイズの多いアーティファクトを処理します。
プロトコルと動作の再構築
C2 のようなフロー、エンコーディング、変換、プロセスの動作、秘密チャネル、マルウェアのようなワークフローを再構築します。
実証された結果を推測から分離します。正しい請求スコア。幻覚を利用したテクニックや裏付けのない結論は罰せられます。
公開ベンチマーク。プライベート汎化テスト。
公開リーダーボードは便利ですが、本格的なエージェント評価には、決定論的なスコアリングと専門家によって検証されたグラウンドトゥルースを備えた、汚染されていないプライベートなタスクが必要です。
プライベート AgentRE-Bench の評価では、未リリースのバイナリ、機密のグラウンド トゥルース、幻覚分析、ツール使用のトレース、および書面による技術レポートが使用されます。
リバース エンジニアリング、脆弱性調査、マルウェア分析、バイナリ動作回復用の AI エージェントを構築するチーム向けに、AgentRE-Bench は、公開ベンチマークを超えた機密の非公開評価を提供します。
目に見えない ELF リバース エンジニアリング タスク
プロトコルと動作の再構築の課題
幻覚と裏付けのない主張の分析
公開リーダーボード外での機密スコアリング
特定のエージェント機能のためのカスタム タスク ファミリ
モデルまたはエージェントのバージョンにわたるオプションの回帰評価
1 つのモデルまたはエージェントを集中的に迅速に評価したいチーム向け。
短い技術レポートと報告書
より詳細な評価が必要な AI ラボ、セキュリティ ベンダー、エージェント チーム向け。
特定のリバース エンジニアリング ワークフローに合わせたタスクを必要とするチーム向け。
パックされたバイナリと PE

分析
プライベート リバース エンジニアリング評価をリクエストする
プライベート評価には、再現可能なサンドボックスの実行、タスクレベルのスコアリング、幻覚分析、ツール使用のトレース、障害モードの内訳、および書面による技術レポートが含まれます。
AgentRE-Bench は、サンドボックス環境でのリバース エンジニアリング エージェントの制御された評価のために設計されています。商用評価は、防御的なマルウェア分析、モデルの信頼性、幻覚の軽減、および証拠に基づいたレポートに焦点を当てています。
リバース エンジニアリング エージェントのための RL トレーニング環境
AgentRE-Bench は単なるリーダーボードではありません。同じ目に見えないバイナリ、決定論的グレーダー、サンドボックス実行トレース、および専門家がラベル付けしたグラウンド トゥルースは、もっともらしいレポートを作成するのではなく、実際のバイナリの動作を回復するためのエージェントをトレーニングするための強化学習環境として機能します。
このトレーニング プログラムは、証拠に基づくリバース エンジニアリングのために特別に開発された報酬テクニックを中心に設計されています。つまり、正しい発見に報酬を与え、裏付けのない主張にペナルティを与え、効果的なツール使用行動を形成し、管理されたカリキュラムを通じてタスクの難易度を高めます。
AgentRE-Bench は、隠れたタスクの結果、決定的な報酬、幻覚ペナルティ、ツール使用の追跡、カリキュラムの進行状況をリバース エンジニアリング エージェントのトレーニング信号に変換することで、RL パイロットをサポートできます。
AgentRE-Bench — 決定論的なスコアリングを備えた LLM リバース エンジニアリング ベンチマーク。

## Original Extract

AgentRE-Bench evaluates AI reverse-engineering agents on compiled binaries with no source code, no hints, and deterministic scoring. Private evaluations and training-environment pilots available.

AgentRE-Bench
Leaderboard
Tasks
Private Evals
GitHub
Contact Us
Light
AI security benchmark platform
AI agents can write code. Can they reverse engineer it?
AgentRE-Bench evaluates LLM agents on compiled reverse-engineering tasks with no source code, no hints, and deterministic scoring.
The public benchmark is the credibility layer: open tasks, reproducible runs, and a leaderboard for AI reverse-engineering agents.
Private evaluations and training-environment pilots test whether agents generalize to unreleased binaries instead of producing plausible reports.
View GitHub
Benchmark Snapshot
V2 Linux Results
13 ELF reverse-engineering tasks evaluated
10 Windows PE tasks in the public ladder
6 frontier model runs on the leaderboard
0 LLM judges; deterministic scoring only
01 Gemini 3.1 Flash Lite 0.667
02 DeepSeek V4 Pro 0.648
03 Claude Opus 4.7 0.512
Cited By / Research Mentions
Referenced in AI Security Research
CrackMeBench: Binary Reverse Engineering for Agents
Mentions AgentRE-Bench as the closest related benchmark for stripped ELF reverse-engineering tasks, with emphasis on malware-like protocol and infrastructure reconstruction.
A Tutorial and Survey of Automated Vulnerability Discovery
Cites AgentRE-Bench in the automated vulnerability discovery and AI security benchmark landscape.
How frontier models perform (Linux levels)
All 13 Linux ELF levels, 25 tool-call budget, Docker-sandboxed static analysis tools. Six frontier models evaluated end-to-end. Headline finding: a small non-thinking model (Gemini 3.1 Flash Lite) leads the field, beating every frontier reasoning model on Main score. Hallucination calibration — not reasoning depth — is the dominant axis on this bench. Windows PE results coming soon.
Bench leader with the lowest hallucination rate in the field. Calibration, not raw reasoning depth, is the win condition here.
Two additional models (Gemini 3.1 Pro Preview and GLM 5.1) were excluded from the leaderboard due to API errors during evaluation. Full details and per-model deep-dives in the analysis writeup .
13 ELF levels and 10 Windows PE levels
From plaintext TCP shells to a synthetic WannaCry-inspired ransomware worm. 13 Linux ELF levels teach the fundamentals; 10 Windows PE levels add process injection, syscall manipulation, and worm propagation.
Calibration beats reasoning depth
Binary reverse engineering punishes unsupported claims. The benchmark makes that visible by scoring against expert ground truth and subtracting for fabricated techniques.
Calibration beats reasoning depth
Reasoning models often overclaimed techniques they could not prove from the binary. Lower-hallucination agents scored better.
Binary RE punishes plausible reports
A confident malware report is not enough. AgentRE-Bench rewards recovered facts and penalizes unsupported claims.
Misleading artifacts fool agents
Tasks include decoy strings, misleading labels, and behavior that must be verified through disassembly or execution.
Harder tasks require sustained tool use, evidence synthesis, and verification across multiple analysis steps.
Can an agent recover verified behavior from a binary?
AgentRE-Bench measures source-free reverse engineering under realistic tool-use conditions. Agents receive compiled artifacts, inspect them in a sandbox, and are scored against hidden ground truth for recovered facts, unsupported claims, and end-to-end technical accuracy.
Source-free binary orientation
Identify file format, architecture, sections, imports, strings, symbols, and first-pass execution clues without source code.
Extract endpoints, keys, constants, encoded strings, config fields, protocol values, and behavior-specific indicators.
Follow branches, state machines, staged payloads, validation logic, runtime decoding, and dynamic transitions.
Deception and obfuscation resistance
Handle misleading strings, fake crypto labels, decoys, anti-debugging, anti-analysis checks, and noisy artifacts.
Protocol and behavior reconstruction
Reconstruct C2-like flows, encodings, transformations, process behavior, covert channels, and malware-like workflows.
Separate proven findings from guesses. Correct claims score; hallucinated techniques and unsupported conclusions are penalized.
Public benchmark. Private generalization tests.
Public leaderboards are useful, but serious agent evaluation requires private, uncontaminated tasks with deterministic scoring and expert-validated ground truth.
Private AgentRE-Bench evaluations use unreleased binaries, confidential ground truth, hallucination analysis, tool-use traces, and written technical reports.
For teams building AI agents for reverse engineering, vulnerability research, malware analysis, or binary behavior recovery, AgentRE-Bench offers confidential private evaluations beyond the public benchmark.
Unseen ELF reverse-engineering tasks
Protocol and behavior reconstruction challenges
Hallucination and unsupported-claim analysis
Confidential scoring outside the public leaderboard
Custom task families for specific agent capabilities
Optional regression evals across model or agent versions
For teams that want a focused, fast assessment of one model or agent.
Short technical report and debrief
For AI labs, security vendors, and agent teams that need a deeper assessment.
For teams that need tasks tailored to a specific reverse-engineering workflow.
Packed binaries and PE analysis
Request a private reverse-engineering eval
Private evaluations include reproducible sandbox runs, task-level scoring, hallucination analysis, tool-use traces, failure-mode breakdowns, and a written technical report.
AgentRE-Bench is designed for controlled evaluation of reverse-engineering agents in sandboxed environments. Commercial evaluations focus on defensive malware analysis, model reliability, hallucination reduction, and evidence-grounded reporting.
An RL training environment for reverse-engineering agents
AgentRE-Bench is more than a leaderboard. The same unseen binaries, deterministic graders, sandbox execution traces, and expert-labeled ground truth can serve as a reinforcement-learning environment for training agents to recover real binary behavior rather than produce plausible-sounding reports.
The training program is designed around reward techniques developed specifically for evidence-grounded reverse engineering: rewarding correct findings, penalizing unsupported claims, shaping effective tool-use behavior, and increasing task difficulty through a controlled curriculum.
AgentRE-Bench can support RL pilots by turning hidden task outcomes, deterministic rewards, hallucination penalties, tool-use traces, and curriculum progression into training signals for reverse-engineering agents.
AgentRE-Bench — LLM reverse engineering benchmark with deterministic scoring.
