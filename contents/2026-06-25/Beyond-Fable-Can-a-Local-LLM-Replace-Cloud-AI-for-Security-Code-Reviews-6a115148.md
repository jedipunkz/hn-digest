---
source: "https://srlabs.de/blog/beyond-fable"
hn_url: "https://news.ycombinator.com/item?id=48672174"
title: "Beyond Fable: Can a Local LLM Replace Cloud AI for Security Code Reviews"
article_title: "Beyond Fable: Can a Local LLM Replace Cloud AI for Security Code Reviews - SRLabs Research"
author: "dubbel"
captured_at: "2026-06-25T12:45:50Z"
capture_tool: "hn-digest"
hn_id: 48672174
score: 1
comments: 1
posted_at: "2026-06-25T12:05:40Z"
tags:
  - hacker-news
  - translated
---

# Beyond Fable: Can a Local LLM Replace Cloud AI for Security Code Reviews

- HN: [48672174](https://news.ycombinator.com/item?id=48672174)
- Source: [srlabs.de](https://srlabs.de/blog/beyond-fable)
- Score: 1
- Comments: 1
- Posted: 2026-06-25T12:05:40Z

## Translation

タイトル: 寓話の向こう側: ローカル LLM はセキュリティ コード レビュー用のクラウド AI を置き換えることができるか
記事タイトル: 寓話を超えて: ローカル LLM はセキュリティ コード レビューでクラウド AI を置き換えることができる - SRLabs Research
説明: 2 つの製品コードベースにおける Claude Fable 5、Claude Opus 4.6/4.8、GLM-5、Gemma4-26b、および Qwen3.6-35B-A3B の比較研究。

記事本文:
寓話を超えて: ローカル LLM はセキュリティ コード レビューでクラウド AI を置き換えることができるか - SRLabs Research
メインコンテンツにスキップ
SRLabs — ホームリサーチ
戻る ソフトウェア アシュアランス コード監査 安全な開発 AI 2026-06-22 • 20 分で読む 寓話を超えて: ローカル LLM はセキュリティ コード レビューのクラウド AI を置き換えることができますか
Karsten Nohl @karsten-nohl 最高イノベーション責任者 Allurity The 問題
セキュリティ コードのレビューは、サイバー セキュリティにおいて最も価値があり、伝統的に労働集約的なサービスの 1 つです。 LLM は、このプロセスにおいて精力的に協力者となっています。彼らは、数千行のコード、相互参照 CWE データベース、および経験豊富なレビュー担当者でも見逃す可能性のある表面パターンをスキャンします。しかし、落とし穴があります。
侵入テストの受信者の多くは、特に金融、政府、重要なインフラストラクチャにおいて、ソース コードがクラウドでホストされているサービスと共有されることを望んでいません。独自のコードをサードパーティ LLM に送信すると、LLM プロバイダーとの契約上の保護手段だけでは完全に軽減できない機密保持とデータの保存に関するリスクが生じます。
その結果生じるジレンマ: 最良の LLM はクラウドでホストされます。セキュリティレビューを最も必要とする企業は、多くの場合、これらの主要な機能を放棄します。
実際、クラウドホスト型モデルのリードはどのくらいあるのでしょうか?私たちは、ローカルでホストされるオープンウェイト モデルは、フロンティア クラウド モデルに匹敵するセキュリティの発見を生み出すことができるのかという実際的な質問に答えることにしました。
答えは次のとおりです。ただし、適切な足場が必要です。
私たちはローカル LLM の限界をテストする一連の実験を実行し、ソース コードをクラウドに公開することなく、クラウド ベースのフロンティア モデルと連携して最適に機能することがわかりました。
約 3B のアクティブなパラメータのみを備えた Qwen3.6-35B-A3B モデルは、ソース コードがマシンから離れることなく完全に Mac ラップトップ上で実行され、生成された検出セットの比較が行われました。

フィンテック アプリと投票アプリの両方でフロンティア クラウド モデル (GLM-5、Claude Opus 4.6) までのサイズを実現でき、独自のいくつかのユニークな発見も得られます。人間によるナッジは一切不要で、各コードベースは 90 分以内に完成しました。コードの読み取り、脆弱性の発見、重大度の分類、CVE 出力のトリアージなどの中心的なタスクに関しては、ローカル モデルは現在、フロンティア モデルと同じレベルにあります。
注意: カウントの同等性を見つけることは、機能の同等性ではありません。この主張は、ローカル モデルはパイプラインの一部として役立つのに十分な競争力があり、その発見は専門家によって同様に影響力があると認識されているというものです。この調査は定量的な側面に焦点を当てていますが、品質の発見は侵入テストの専門家と開発者チームの両方によって検証されました。
ローカル モデルがまだ行っていないことは、レビューを設計し、結果を統合することです。私たちが見つけた最も効果的なパイプラインは、これらのオーケストレーション タスクの両方をクラウド フロンティア モデルに委任しますが、どちらの段階でもクラウドはソース コードを認識しません。これをソースローカルと呼びます。独自のソース コードがマシンから離れることはありません。メタデータはクラウド (ファイル ツリー、スキーマ、ルート、依存関係マニフェスト、生成されたステップ プロンプト) に送信され、内部名、ディレクトリ構造、アーキテクチャを伝えることができます。 「建物から音源が一切出ない」というのは正確な約束ですが、「何も出ない」はそうではありません。
この作業を行う足場は 3 つの部分で構成されます。
構造化された分解とプロンプトの生成 - クラウド モデルはレビューを焦点を絞ったステップに分割し、メタデータのみからステップ プロンプトを作成します (ファイル ツリー、スキーマ、ルート - ソース コードは不要)
ローカル ツールと LLM 出力 - プロンプトはローカルで実行され、標準のセキュリティ ツール (bundler-audit、npm Audit、Semgrep、Brakeman など) を実行し、コンテキスト トライのローカル モデルに JSON 出力をフィードします。

年齢と追加のバグハンティング
レポートの統合 - 最終的なクラウド パスでは、ステップレベルの調査結果が配信可能なレポートに統合されます。
パート 1 と 3 では、ソース コードをクラウドに公開する必要はありません。パート 2 は完全にローカルで実行されます。
結果として得られるベスト プラクティスは、迅速なエンジニアリングにはクラウド、実行にはローカル、統合にはクラウドです。クラウド モデルはソース コードを決して参照せず、レビューを設計します。ローカル モデルには広範なアーキテクチャ上の推論は必要ありません。バンドルされたファイルに対して集中的なチェックが実行されます。
図 1. ソースローカル パイプライン。クラウド オーケストレーターはレビューを設計し (ステージ 1)、メタデータのみから結果をレポートに統合します (ステージ 3)。ローカル モデルはソースを読み取り、セキュリティ ツールを実行します (ステージ 2)。ステップ プロンプトとステップ レベルの結果のみが信頼境界を越えます。ソース コードがマシンから離れることはありません。
事例 5 の活用: クラウドベースのオーケストレーション層はモデルに依存しません。ステージ 1 と 3 のオーケストレーターは、無制限のフロンティア モデルである必要はありません。サイバーセキュリティのガードレールを備えたモデルは、その仕事をうまく処理します。 Claude Fable 5 は意図的なサイバー制限付きで出荷され、レビュー プロンプトを設計し、拒否や品質の低下なく結果を統合し、これらの役割において Claude Opus 4.8 と完全に一致します。これは驚くべきことではありません。防御的なレビューの設計と統合は知識と構造の作業であり、悪用ではなく、オーケストレーターはソース コードには決して触れません。
ただし、オーケストレーター モデルとエグゼキューター モデルの両方を選択すると、検出される内容が変わります。オーケストレーターが生成するプロンプト設計は、ローカルのエグゼキューター モデルを実質的に異なる脆弱性に向けて誘導するため、2 つのオーケストレーターのプロンプトを結合した方が、どちらか単独よりも優れています。 「単一のモデルですべてが見つかるということはない」は、プロンプトデザインとプロに当てはまります。

mpt 実行層。
1. 単一のモデルですべてが見つかるわけではない
すべてのモデルの結果を統合すると、個々のモデルの出力よりも大幅に大きくなります。各モデルでは、質的に異なるクラスの脆弱性が見つかりました。
クロードは建築的推論に優れていた
データフロートレースとツール統合における GLM-5
Gemma4 は、焦点を当てたファイル セット内で行レベルのコード パターン マッチングを行います。
積極的な重大度調整を備えた広範囲をカバーする Qwen3.6。
実務者への影響 : 「セカンドオピニオン」モデルを実行すると、はるかに小規模なモデルを使用する場合でも、対象範囲が真に拡大されます。これは、コードベースとテストされたすべてのモデルの両方に当てはまります。
図 2. 両方のコードベースにわたって各モデルによって検出された個別の脆弱性カテゴリ (モデル間のカバレッジ マトリックスから合計 53)。この結合は単一モデルを矮小化します。Qwen には最もユニークなカテゴリがあり、GLM-5 と Qwen は最大の重複を共有し、4 つすべてに含まれるカテゴリは 2 つだけでした。 (カウントはカテゴリであり、検証された真陽性ではありません。Gemma4 はフィンテック上でのみ実行されました。)
2. モデルのサイズよりも迅速なエンジニアリングが重要
適切に構造化されたプロセスにより、すべてのモデルがより優れたものになります。実際には、モデルの機能の違いがこの「ハーネス」に依存しないほど優れています。
たとえば、最大 38 億のアクティブ パラメータ予算で実行される Gemma4 (合計は「260 億」であるにもかかわらず、専門家の混合モデルです) は、はるかに大規模なフロンティア モデルが見逃していた 3 つの真の発見を発見しました。ランニングコストが安いにもかかわらず、機能面では競争力があり、違いは生の機能ではなく、迅速な設計でした。これには準備が必要です。「ハーネス」の準備をスキップして、Gemma4 にモノリシック プロンプトを与えると、不完全な結果が生成され、出力命令を追跡できなくなりました。同じスコープを expl で 6 つの焦点を絞ったマイクロタスクに分解した場合

Gemma は、icit ファイル パスと grep コマンドを使用して、特定の行番号とコードの証拠を含む実用的な調査結果を作成しました。どちらにしても幻覚はありません。
これは、ローカル モデルの品質の上限が予想よりも高いことを示唆していますが、それに到達するには、検索をガイドするためのハーネスの準備が必要です。この準備作業自体は自動化できることがわかりました。Claude はファイル ツリーのみ (ソース コードなし) からステップ プロンプトを生成し、それらの自動生成されたプロンプトを Qwen が実行すると、どちらのクラウド モデルの単一プロンプト レビューよりも多くの結果が得られました。これを繰り返すことが重要です。クロードが小さなモデルを実行するためのプロンプトを準備した場合、その小さなモデルは、クロードがテスト全体に責任を感じているレビュー以上のものを見つけます。
図 3. 同じプレイブックを 1 つのプロンプトとして実行する場合と複数のプロンプトに分割する場合の比較 (フィンテック)。クロード (22→43、+95%) と GLM-5 (28→54、+93%) はほぼ 2 倍です。小型の Gemma4 モデルではゲインが少なく (12→17、+42%)、天井が低いことがわかります。 Qwen はワンショットのベースラインがないため省略されます。
隣接する作業は同じ方向を向いています。 Niels Provos 氏は、Finding Zero-Days with Any Model の中で、「脆弱性の発見はフロンティアモデルの問題ではなく、オーケストレーションの問題である」と主張し、モデル全体に​​わたる実際の欠陥を表面化する FSM 主導のハーネスを実証しています。彼の結果について正確に言うと (読みすぎやすい)、27 年前の OpenBSD TCP SACK バグのヘッドラインの複製では、商用の Claude (Sonnet 4.6 から Opus 4.6 にエスカレート) が使用され、ファジングと QEMU の概念実証で検証されましたが、オープンウェイト GLM 5.1 は別のターゲットで実行されました。 Provos の研究の領域 (実行可能な PoC を使用したディープ C ゼロデイ ハンティング) は、私たちの研究 (CWE マップされた結果を含む Web アプリケーションのレビュー、PoC なし) とは異なります。どちらも同じ結論に達します。
3. レポート品質

ty は大幅に変化する
レポートの品質は明らかに大規模かつフロンティア モデルの方が優れており、ローカル モデルが実際のテストを行う「ソースローカル」レビューにも場所があることが改めて示唆されています。
クロード・オーパスのレポートは即時配信に向けて最も洗練されたものでしたが、最も人間的なナッジが必要でした (執筆プロセス全体で最大 6 回のリマインダー)
GLM-5 は最も包括的な成果物セットを作成しましたが、時折幻覚を示す出力参照がレポートの品質を損なうことがあります。
Qwen は、正しい CWE マッピングを備え、幻覚症状のない、よく構造化されたステップごとのレポートを作成しました。ステップレベルの出力は、Claude によって配信可能なレポートに正常に統合されました (ソースローカル統合ステージ)
Gemma4 の出力には最も多くの後処理が必要でした
4. レビュー オーケストレーターには、サイバー制限のある寓話であっても、あらゆる機能を備えたモデルを使用できます
2026 年 6 月にリリースされた Claude Fable 5 には、強力なサイバーセキュリティ ガードレールが付属しています。 Anthropic は、これらを大規模なレッドチームによって強化された安全対策として組み立てており、その後の米国の輸出規制による Fable の停止は、これらを純粋なマーケティングとして解釈することを妨げています。実際、Fable は攻撃的/悪用のリクエストを拒否しますが、レビューの準備と分析のステップ、つまりソースローカルのレビューに有能なフロンティア モデルが必要な段階をすぐに支援します。
2 つのコードベースの Qwen 実行テストについて、2 つのオーケストレーター (Claude Fable 5 と Claude Opus 4.8) を比較します。実践者にとって重要な結果は 2 つあります。
サイバー制限されたフロンティア モデルは、有能なオーケストレーターです。Fable 5 は、これらの役割において、拒否や明らかな品質のギャップがなく、Opus 4.8 と比較して、完全なプロンプト パックと厳密な統合を生成しました。 (Fable は特定のサイバー リクエストを内部的に Opus 4.8 にルーティングできます。私たちはこれを監視していましたが、そのような手口は見つかりませんでした

これらの防御的なオーケストレーションの実行中に ff が実行されるため、これはサイレント フォールバックではなく、Fable 自体です。)
オーケストレーターの選択によって、何が見つかるかが変わります。Fable のターゲットを絞ったプロンプトは、既知の困難な「センチネル」バグを確実に回復しましたが、より緊密なセットを生成しましたが、Opus のより広範なプロンプトは、特にターゲットにしていなかったセンチネルが欠落するという代償として、Fable アームが決して提起しなかったクリティカル (否定票のゼロコスト投票、認証されていない投票用紙の詰め込み) を含む、より大規模で場所によってはより深刻なセットを表面化させました。エグゼキューターと同様に、ユニオンはいずれかのオーケストレーターのみを支配します。Fable は Opus よりも「優れている」わけではありません。これらは並行して実行するのが最適です。
実務者への示唆 : オーケストレーションは、パイプラインのモデルにほとんど依存しない部分として考える必要があります。アクセスできる機能のあるクラウド モデルを (制限付きかどうかに関係なく) 使用できます。また、ローカル モデルに対して 2 つのプロンプト デザインを実行すると、2 番目のエグゼキュータと同様に対象範囲が拡大します。最良の結果を得るには、組み合わせて組み合わせてください。
図 4. オーケストレーターによる 2 回の実行にわたる Sentinel バグの回復 (ライブ ソースに対してチェック)。 Fable のターゲットを絞ったプロンプトは、実際のセンチネルをすべて再検索しますが、実行 2 で誤検知を作り出します (
実験のセットアップ
私たちは、技術スタックと脅威プロファイルが異なる 2 つの本番コードベースを使用しています。
Fintech ダッシュボード — Next.js / TypeS

[切り捨てられた]

## Original Extract

A comparative study of Claude Fable 5, Claude Opus 4.6/4.8, GLM-5, Gemma4-26b, and Qwen3.6-35B-A3B on two production codebases.

Beyond Fable: Can a Local LLM Replace Cloud AI for Security Code Reviews - SRLabs Research
Skip to main content
SRLabs — Home Research
Back Software Assurance Code Audit Secure Development AI 2026-06-22 • 20 minute read Beyond Fable: Can a Local LLM Replace Cloud AI for Security Code Reviews
Karsten Nohl @karsten-nohl Chief Innovation Officer Allurity The Problem
Security code review is one of the most valuable — and traditionally labor-intensive — services in cyber security. LLMs have become tireless wingmen in this process: They scan thousands of lines of code, cross-reference CWE databases, and surface patterns that even experienced reviewers might miss. But there's a catch.
Many pentest recipients do not want their source code shared with cloud-hosted services — particularly in finance, government, and critical infrastructure. Sending proprietary code to a third-party LLM creates confidentiality and data residency risks that contractual safeguards with the LLM provider alone cannot fully mitigate.
The resulting dilemma: The best LLMs are cloud-hosted. Those companies who need security reviews the most, often forgo these leading capabilities.
How big is the lead of cloud-hosted models really? We set out to answer a practical question: can a locally-hosted open-weight model produce security findings comparable to frontier cloud models?
We find the answer is: almost — but only with the right scaffolding.
We ran a series of experiments testing the limits of local LLM and found that they work best in tandem with cloud-based frontier models, but without disclosing source code to the cloud:
A Qwen3.6-35B-A3B model with only ~3B active parameters, running entirely on a Mac laptop with no source code leaving the machine, produced finding sets comparable in size to frontier cloud models (GLM-5, Claude Opus 4.6) on both a fintech app and a voting app, with some unique findings of its own. It required zero human nudges and completed each codebase in under 90 minutes. For the central task — reading code, discovering vulnerabilities, classifying severity, triaging CVE output — a local model is now in the same league as frontier models.
A caveat: Finding count parity is not capability parity. The claim is that a local model is competitive enough to be useful as part of the pipeline , and that its findings are perceived as equally impactful by experts. This study focuses on the quantitative side, but finding quality was validated by both pentest experts and a developer team.
What a local model does not yet do as well is design the review and consolidate the results . The most effective pipeline we found delegates both of these orchestration tasks to a cloud frontier model — but in neither stage does the cloud see source code. We call this Source-local : the proprietary source code never leaves the machine. Metadata does cross to the cloud (file tree, schema, routes, dependency manifests, and the generated step prompts), which can carry internal names, directory structure, and architecture. "No source leaves the building" is the accurate promise — "nothing leaves" is not.
The scaffolding that makes this work has three parts:
Structured decomposition and prompt generation — a cloud model breaks the review into focused steps and creates step prompts from metadata only (file tree, schema, routes — no source code)
Local tool and LLM output — the prompts execute locally, run standard security tools (e.g., bundler-audit, npm audit, Semgrep, Brakeman) and feed their JSON output to the local model for contextual triage and additional bug hunting
Report consolidation - a final cloud pass merges the step-level findings into a delivery-ready report.
Parts 1 and 3 require no source code exposure to the cloud; Part 2 runs entirely locally.
The resulting best practice is: cloud for prompt engineering, local for execution, cloud for consolidation. The cloud model never sees source code — it designs the review. The local model never needs broad architectural reasoning — it executes focused checks against bundled files.
Figure 1. The Source-local pipeline. The cloud orchestrator designs the review (stage 1) and consolidates findings into a report (stage 3) from metadata only; the local model reads the source and runs the security tools (stage 2). Only step prompts and step-level findings cross the trust boundary — the source code never leaves the machine.
Leveraging Fable 5: the cloud-based orchestration layer is model-agnostic. The orchestrator in stages 1 and 3 need not be an unrestricted frontier model; a model with cybersecurity guardrails handles the job just fine. Claude Fable 5, which ships with deliberate cyber restrictions, designs the review prompts and consolidates the findings with no refusal and no loss of quality, fully matching Claude Opus 4.8 in those roles. This is unsurprising: designing and consolidating a defensive review is knowledge-and-structure work, not exploitation, and the orchestrator never touches source code.
The choice of both orchestrator and executor model, however, changes what gets found — the prompt design the orchestrator produces steers the local executor model toward materially different vulnerabilities, so the union of two orchestrators' prompts beats either alone. “No single model finds everything” holds true on the prompt-design and prompt-execution layers.
1. No Single Model Finds Everything
The union of all models' findings is significantly larger than any individual model's output: Each model found qualitatively different classes of vulnerabilities:
Claude excelled at architectural reasoning
GLM-5 at data flow tracing and tool integration
Gemma4 at line-level code pattern matching within focused file sets
Qwen3.6 at breadth coverage with aggressive severity calibration.
Implication for practitioners : Running a "second opinion” model genuinely expands coverage, even when using a much smaller model. This held across both codebases and all models tested.
Figure 2. Distinct vulnerability categories caught by each model across both codebases (53 total, from the cross-model coverage matrices). The union dwarfs any single model: Qwen has the most unique categories, GLM-5 and Qwen share the largest overlap, and only two categories were caught by all four. (Counts are categories, not validated true positives; Gemma4 ran on Fintech only.)
2. Prompt Engineering Matters More Than Model Size
A well-structured process makes every model better. So much better in fact, that the differences in model capability become secondary to this ‘harness’.
For example, Gemma4 — which runs on a ~3.8B active-parameter budget (it is a Mixture-of-Experts model, despite the "26B" total) — found three genuine findings that far larger frontier models missed. It is cheap to run yet competitive in capability, and the difference here was not raw capability but prompt design . This takes preparation: When skipping the ‘harness’ preparation and giving a monolithic prompt to Gemma4, it produced incomplete results and lost track of output instructions. When the same scope was decomposed into six focused micro-tasks with explicit file paths and grep commands, Gemma produced actionable findings with specific line numbers and code evidence. No hallucinations either way.
This suggests that the quality ceiling for local models is higher than expected — but reaching it requires harness preparation to guide the search. We find that this preparation effort can itself be automated : Claude generated step prompts from a file tree alone (no source code), and Qwen executing those auto-generated prompts produced more findings than either cloud model's single-prompt reviews. Important to repeat this: When Claude prepared a prompt for a smaller model to run, that smaller model finds more than a review where Claude feels responsible for the entire test.
Figure 3. Running the same playbook as one prompt vs split into multiple prompts (Fintech). Claude (22→43, +95%) and GLM-5 (28→54, +93%) nearly double; the small Gemma4 model gains less (12→17, +42%), revealing a lower ceiling. Qwen is omitted as it has no one-shot baseline.
Adjacent work points in the same direction. Niels Provos, in Finding Zero-Days with Any Model , argues that "vulnerability discovery is an orchestration problem, not a frontier-model problem," demonstrating an FSM-driven harness that surfaces real flaws across models. To be precise about his results (they are easy to over-read): his headline replication of the 27-year-old OpenBSD TCP SACK bug used commercial Claude — Sonnet 4.6 escalating to Opus 4.6 — and was validated with fuzzing and QEMU proof-of-concepts, while the open-weight GLM 5.1 was exercised on a different target. The domain of Provos’ study (deep C zero-day hunting with executable PoCs) differs from ours (web-application review with CWE-mapped findings, no PoC). Both reach the same conclusion.
3. Report Quality Varies Dramatically
The report quality is clearly better for larger and frontier models, once again suggesting that they have a place even for “Source-local” reviews where local models do the actual testing:
Claude Opus's report was the most polished for immediate delivery but required the most human nudges (~6 reminders across the writing process)
GLM-5 produced the most comprehensive deliverable set, but occasional hallucinated output references tarnish the report quality
Qwen produced well-structured per-step reports with correct CWE mappings and no hallucinated evidence. The step-level output was successfully consolidated by Claude into delivery-ready reports (the Source-local consolidation stage)
Gemma4's output required the most post-processing
4. The Review Orchestrator Can Be Any Capable Model — Even Cyber-Restricted Fable
Claude Fable 5, released in June 2026, ships with strong cybersecurity guardrails. Anthropic frames these as safety measures hardened through extensive red-teaming — and the subsequent US export-control suspension of Fable cuts against reading them as pure marketing. In practice Fable declines offensive/exploitation requests but readily helps with the preparation and analysis steps of a review — exactly the stages where a Source-local review needs a capable frontier model.
We compare two orchestrators — Claude Fable 5 vs Claude Opus 4.8 — for the Qwen-executed tests of two codebases. Two results matter for practitioners.
A cyber-restricted frontier model is a competent orchestrator : Fable 5 produced complete prompt packs and rigorous consolidations with no refusal and no obvious quality gap versus Opus 4.8 in those roles. (Fable can route certain cyber requests to Opus 4.8 internally; we watched for this and saw no such handoff during these defensive orchestration runs — so this is Fable itself, not a silent fallback.)
The choice of orchestrator changes what gets found : Fable's targeted prompts reliably recovered the known-hard "sentinel" bugs but produced a tighter set, while Opus's broader prompts surfaced a larger and in places more severe set — including criticals the Fable arm never raised (negative-vote zero-cost voting, unauthenticated ballot stuffing) — at the cost of missing sentinels it did not specifically target. As with executors, the union dominates either orchestrator alone: Fable is not "better" than Opus — they are best run in parallel.
Implication for practitioners : Orchestration should be viewed as a largely model-agnostic part of the pipeline: You can use whichever capable cloud model you have access to (restricted or not). And running two prompt designs for the local model expands coverage just as a second executor does. Mix and match for best results.
Figure 4. Sentinel-bug recovery by orchestrator across two runs, checked against the live source. Fable's targeted prompts re-find every real sentinel but invent a false positive in run 2 (a
Experiment Setup
We use two production codebases with different tech stacks and threat profiles:
Fintech Dashboard — a Next.js / TypeS

[truncated]
