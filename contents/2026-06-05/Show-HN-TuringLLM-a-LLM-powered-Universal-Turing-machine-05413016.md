---
source: "https://github.com/gmlion/TuringLLM"
hn_url: "https://news.ycombinator.com/item?id=48416857"
title: "Show HN: TuringLLM – a LLM-powered Universal Turing machine"
article_title: "GitHub - gmlion/TuringLLM: An LLM-powered universal Turing machine · GitHub"
author: "gmlion"
captured_at: "2026-06-05T21:36:51Z"
capture_tool: "hn-digest"
hn_id: 48416857
score: 1
comments: 0
posted_at: "2026-06-05T19:09:43Z"
tags:
  - hacker-news
  - translated
---

# Show HN: TuringLLM – a LLM-powered Universal Turing machine

- HN: [48416857](https://news.ycombinator.com/item?id=48416857)
- Source: [github.com](https://github.com/gmlion/TuringLLM)
- Score: 1
- Comments: 0
- Posted: 2026-06-05T19:09:43Z

## Translation

タイトル: 表示 HN: TuringLLM – LLM を搭載したユニバーサル チューリング マシン
記事タイトル: GitHub - gmlion/TuringLLM: LLM を利用したユニバーサル チューリング マシン · GitHub
説明: LLM を搭載した万能チューリング マシン。 GitHub でアカウントを作成して、gmlion/TuringLLM の開発に貢献してください。
HN テキスト: こんにちは。
LLM によって完全に制御されるユニバーサル実行モデルでどこまでできるかを確認するために、TuringLLM を構築しました。基本的な考え方: チューリング マシンのステップ関数に LLM を使用します。状態と命令は MD ファイル (それぞれ STATE.md と INSTRUCTIONS.md) です。これらはチューリング マシンの「修正可能なテープ」を表します。各サイクルで、LLM は状態を読み取り、実行する対応する命令を見つけます。命令 (および条件) はフリーテキストとして記述され、実行中に状態とともに LLM 自体によって記述できます。各サイクルのスコープは制御されます。各 LLM 呼び出しは STATE と INSTRUCTIONS のみを受け取ります。それに加えて、コールスタックメカニズムにより、引数の受け渡しと戻り値を含むサブルーチンの階層呼び出しが提供されます。これを使用して、マルチエージェント パターンやメタ フレームワークを自由に実装できます。その「普遍性」をテストするために、Tree of Thoughts、LATS、Meta got、ADAS などの MAS 文献から 14 のパターンを実装しました。可能な限り共通の演算子を共有します。この部分はまだほとんど WIP であり、少し荒削りな部分があるかもしれません。そのため、お気軽に覗いて改善点を提案してください。ただし、実装は簡単で、すぐに新しいパターンが実行されるのが確認できます。必要なのは、INTERPRETER.md をビルドするよう求めるプロンプトだけです。サイクルとサブルーチンをグラフとして、または各サイクルをマシンの状態のログとしてレンダリングするビジュアライザーもあります。特にエージェントや MAS の研究に取り組んでいる方は、ご意見をお聞かせください。

記事本文:
GitHub - gmlion/TuringLLM: LLM を利用したユニバーサル チューリング マシン · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
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
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
グムリオン
/
チューリングLLM
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 Cod

e その他のアクション メニューを開く フォルダーとファイル
328 コミット 328 コミット ベンチマーク ベンチマーク docs docs インタプリタ インタプリタ スクリプト スクリプト src src .gitattributes .gitattributes .gitignore .gitignore CLAUDE.md CLAUDE.md INSTRUCTIONS.md INSTRUCTIONS.md ライセンス ライセンス通知 通知 README.md README.md VISUALIZER.md VISUALIZER.md new-instance.sh new-instance.sh package-lock.json package-lock.json package.json package.json sdlc-poc-artefact-spec.md sdlc-poc-artefact-spec.md sdlc-poc-interpreter-topology.md sdlc-poc-interpreter-topology.md setup-telegram.sh setup-telegram.sh tsconfig.json tsconfig.json Visualize.sh Visualize.sh Visualizer.html Visualizer.html すべてのファイルを表示 リポジトリ ファイルのナビゲーション
LLM を搭載した万能チューリング マシン。
LLM をチューリング マシンのステップ関数として扱います。すべて
それ以外の場合は終了します: 状態はディスク上に存在し、プログラムはマークダウンされ、実行されます。
再開可能で監視可能であり、「エージェント」は単にユーザーが作成したものです
ステートマシン。
サイクル ループは、サイクルごとに 1 回 LLM を呼び出します。 LLM はその状態を読み取ります
( MEMORY.md ) とプログラム ( INSTRUCTIONS.md ) の最初のものと一致します。
条件が適合する命令は（通常は書き換えることによって）動作します。
MEMORY.md ) が破棄されます。このサイクルは停止するまで繰り返されます。そこに
隠された会話状態ではありません。次のサイクルで確認されるのは、
前のものがディスクに残したものとまったく同じです。
TuringLLM は、何よりもまずシェル、つまり薄いユニバーサルです
一連のワンショット LLM 呼び出しをステートフルに変換するエグゼキューター
再開可能で監視可能なマシン。コールスタックプリミティブを同梱
(引数置換と戻り値のスプライシングを使用したプッシュ/ポップ フレーム)、
インスタンスごとに 2 つの git リポジトリ (サイクルごとに 1 つが自動コミット、1 つが
LLM 制御)、およびよく知られている MEMORY セクションの小さなセット
シェルがインターセプトします。 src/ 内のすべてとトップレベルのスクリプトは次のとおりです。

殻の部分。
LLM が各サイクルで何を行うかを決定するもの
シェルの外側、ユーザー作成のマークダウン内に存在します。
通訳/ 例のカタログを発送 —
マルチエージェントからのパターンの実行可能な実装を含む
システム文献 (自己洗練、反射、検証の連鎖、
計画と実行、思考の木、LATS、マルチエージェント討論、
MetaGPT、ChatDev、AFlow、ADAS) — ただし、シェル自体は知りません
彼らについて。これらがなくても TuringLLM を完全に正常に実行できます。
パターン カタログは、A/B を直接比較できるように設計されています。一部
デモではタスクを一定に保ち、戦略を変更します (ToT と LATS の比較)
同じ24ゲームのパズル。 MetaGPT と ChatDev の比較
wc-plus ビルドタスク);他の人は戦略を一定に保ち、変化します
タスク (3 つの計画分解デモはバイト数で実行されます)
3 つの異なるプログラムに関する戦略コードを使用して、公開された 3 つを表示します
パターンは 1 回の再帰に折りたたまれます)。カタログの配置方法を見る
A/B 比較用にアウト
完全な内訳については。 MAS 研究者ならカタログ
を再実装せずにパターンを比較するための基板です。
毎回周囲の足場。
シェルを構築してから、HITL を使用したディープリサーチのデモを実行します。
いくつかの明確な質問をする再帰的な研究者
事前に考え、考えている間に背景を収集し続け、
構造化レポートは、instances/<name>/workspace/report.md にあります。
npmインストール
npm ビルドを実行する
# インスタンスを作成します。 PROGRAM.md には、あらかじめ入力された状態で出荷されます。
# 研究目標が明確化されていないため、明確化が必要です。
./new-instance.sh interpreters/mas-papers/2-planning-decomposition/c-deep-research-hitl my-cdr-hitl
# オプションですが推奨: LLM の明確な質問をルーティングします。
# Telegram を通じてプッシュ通知として届くので、
# 答えます

携帯電話から。 Stdin も機能します (質問はインラインで出力されます)。
./setup-telegram.sh < BOT_TOKEN > instances/my-cdr-hitl # Human-in-the-loop を参照
# 実行
インスタンス/my-cdr-hitl/run.sh
2 番目のターミナルでビジュアライザを開いてコール スタックを監視します。
実行の展開に応じてイベント ストリームがライブで表示されます。
./visualize.sh my-cdr-hitl
戦略フレーム f000-plan-execute-clarify が表示されます。
いくつかの質問から ## Pendingquestions まで、シェルはそれらを次の方法で配信します。
チャンネルと LLM が循環し続け、背景を追加します
研究メモ — 待っている間。答えが到着したら（または
戦略はバックグラウンドが使い果たされたと判断します）それはタックル.mdをプッシュし、
再帰的ソルバーは最終レポートを作成します。
プロバイダーのデフォルトは、Haiku を使用した Claude Code です。に切り替えるには、
Anthropic API、OpenAI 互換エンドポイント、Ollama、またはローカル
GGUF モデル、instances/my-cdr-hitl/.env を編集 — を参照
プロバイダー。
┌───────────────────────────┐
│ インスタンス │
│ │
│ プログラム.md 説明書.md メモリ.md │
│ ┌─────────┐ ┌───────┐ ┌───────┐ │
│ │ # 目標 │ │ # 戦略 │ │ ## 状態 │ │
│ │ │ │ (アクティブなフレーム) │ │ 現在 │ │
│ │ (ユーザー- │◄───│ │ │ │
│ │ 執筆） │ │ # 副指導。 │───►│ ## 結果 │ │
│ │ │ │ (生成) │ │ ... │ │
│ ━━━━━━━┘ ━━━━━━┘ ━━━━━━┘ │
│ │
│ ワークスペース/履歴/ログ/ │
│ ┌───

─────┐ ┌───────┐ ┌───────┐ │
│ │ (プロジェクト │ │ 0001-a3f1b2c/ │ │ run-*.log │ │
│ │ アーティファクト、│ │ 0002-b4e2c3d/ │ │ (フル │ │
│ │ 独自の git │ │ ... │ │ 出力） │ │
│ │ リポ） │ │ │ │ │
│ ━━━━━━━┘ ━━━━━━┘ ━━━━━━┘ │
━━━━━━━━━━━━━━━━━━━━━━━━┘
│
┌───────────┼─────────┐
│ │ │
▼ ▼ ▼
┌─────┐ ┌─────┐ ┌────────┐
│ マシン │ │ シェル │ │ プロバイダー │
│ Git │ │ main.ts │ │ クロードコード、 │
│ (auto- │ │ (cycle │ │ api、openai、│
│ コミット │ │ ループ、 │ │ オラマ、 │
│ ごと │ │ 再試行、│ │ ローカル │
│ サイクル） │ │ 停止） │ │
━━━━━━━━━━━━━━━━━━━━━━━┘
3層設計
┌─────────────────────┐
│ PROGRAM.md (ユーザー層) │
│ │
│ 高いレベルの目標。ユーザーによって書かれました。 │
│ 機械による改造は一切行っておりません。 │
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

┤
│ INSTRUCTIONS.md (フレームごとのプログラム) │
│ │
│ # 戦略セクション: 不変メタプログラム │
│ PROGRAM.mdを解釈する。書き換えても存続します。 │
│ │
│ # サブ命令セクション: 可変動作 │
│エリア。各サイクルで生成および消費されます。 │
━━━━━━━━━━━━━━━━━━━━━━┤
│ シェル（ユニバーサルエグゼキューター） │
│ │
│ シングルサイクルループ。 LLM を呼び出し、再試行します │
│ 不完全なサイクル、自動コミット、スナップショット。 │
│ ハードコーディングされたフェーズはありません — INSTRUCTIONS.md が決定します。 │
━━━━━━━━━━━━━━━━━━━┘
PROGRAM.md と INSTRUCTIONS.md はどちらもシェルの外部で作成されます。
シェルの仕事は、(a) 適切なファイルを LLM にフィードすることです。
適切なタイミングで、(b) 既知のメモリーの小さなセットを解釈します。
通話後のセクション。それ以外はすべてユーザーコードです。
シェルは、小さな固定されたプリミティブのセットを提供します。上に構築されたものはすべて
それらの先頭にあるのは、interpreters/ の例を含む、ユーザーです。
コード。
アクティブなフレーム (呼び出しスタックの最上位) と cd を識別します。
そこに。
そのフレームの MEMORY.md と INSTRUCTIONS.md を読み取り、
プロンプトが表示され、LLM が呼び出されます。
LLM は、条件が一致する最初の ## 命令を見つけます。
MEMORY の ## 状態と一致し、その命令を実行します
アクション。通常は MEMORY.md を書き換えます。
シェルは新しい MEMORY を読み取り、既知のセクションを処理します。
インターセプト (以下を参照)、git へのコミット、履歴のスナップショット、
そしてループ。 ↻
サイクルをまたがる会話の継続性は暗示されません - LLM
前のサイクルで MEMORY.md に書き込まれた内容 (および
./scoped/ ユーザーコードをファイルします

残すことを選択しました）。
シェルは、フレームのインスタンスごとの呼び出しスタックを維持します。の
トップオブスタックフレームは「アクティブ」です: その INSTRUCTIONS.md と
MEMORY.md は次のサイクルを駆動します。 LLM はプッシュ/ポップの意図を通知します。
記憶を通して。最小限のプッシュは、
ロードするファイル:
## プッシュ
パス/to/some-file.md
プッシュ ターゲットは、 ## Push-Args 経由で名前付き引数を受け取ることもできます。
値は、キーと値のペアまたは YAML スタイルのブロック スカラー、および任意の値です。
プッシュされたファイルのテキスト内の {{key}} プレースホルダーは、前に置き換えられます。
新しいフレームが始まります:
## プッシュ
パス/to/some-file.md
## プッシュ引数
<キー>: <インライン値>
<キー>: |
<複数行ブロックのスカラー — 改行は保持される>
押してください。シェルが ## を認識すると、次の処理の前に MEMORY にプッシュします。
サイクルすると、現在のフレームがスタックに保存され、新しいフレームが作成されます。
指定されたファイルからロードされたアクティブなフレーム:
プッシュ前の発信者の記憶 プッシュ後のフレーム
─────────────────────
## 州 ## 州
空のドラフト
## プッシュ (INSTRUCTIONS.md ← <パス>,
すべての {{key}} が置き換えられた <path>
## 対応する値で引数をプッシュします)
キー: 値 (./scoped/ は空で始まります)
プッシュされたフレームが独自のフレームを実行している間、呼び出し元は一時停止されます。
ステートマシン。
ポップ。プッシュされたフレームの状態が Done に達すると、シェルは
それを破棄し、呼び出し元を再開します。 ## Return ブロック (場合)
存在すると、ポップされたフレームの MEMORY が呼び出し元の MEMORY にトップレベルとして接続されます。
セクション (キー foo は ## Foo になります)、呼び出し元の状態は次のとおりです。
<caller_state>_completed に名前が変更されました。
ポップ前にプッシュされたフレームのメモリ ポップ後の呼び出し元のメモリ
───────────────────

──────
## 州 ## 州
完了しました
## リターン ## キー
キー: 値 値
状態の _completed サフィックスにより、呼び出し元は
「プッシュしただけで結果が返された」と元のメッセージを区別する
プッシュを発行した場所の状態を示します。そうでない場合は、すぐにプッシュされます。
同じ命令を再起動して、もう一度押します。
深さ。プッシュ フレームは任意にネストできます。プッシュされたオペレータは
それ自体が他のものを押します。 2 つの量はランの深さを表します。
ドキュメントでは一貫してこれらを使用しています。
stack.length — スタック上のフレームのリテラル数。
ルート フレームは常に存在するため、常に ≥ 1 になります。
Depth — ゼロから数えられるネスト レベル。ルートフレーム
は深さ 0 で、プッシュするフレームは深さ 1 などとなります。
深さ == stack.length - 1 。
したがって、 d-cove ( cove.md → verify.md → Answer-independently.md )
深さ 2 に到達 — 3 フレーム。これまでにないフラットな通訳
1 人のオペレータをプッシュすると、深さ 1 (2 フレーム) に達します。シェルが停止する
state == 完了し、ルート フレームのみが残ったとき
( stack.length == 1 、つまり深さ 0)。
スタックは .call-stack.json に保存され、スナップショットが作成されます。
すべての履歴/エントリ。
例で使用されている規則 - プッシュ可能ファイルの呼び出し
「operators」を演算子/に入れると、純粋にユーザーになります
慣例。シェルは何も気にしません

[切り捨てられた]

## Original Extract

An LLM-powered universal Turing machine. Contribute to gmlion/TuringLLM development by creating an account on GitHub.

Hi,
I built TuringLLM to see how far we could go with a universal execution model totally controlled by an LLM. The basic idea: use the LLM for the step function of a Turing machine. State and instructions are MD files (STATE.md and INSTRUCTIONS.md, respectively). They represent the "modifiable tape" of the Turing machine. Each cycle, the LLM read the state and finds the corresponding instruction to execute. Instructions (and conditions) are written as free-text and, together with state, can be written by the LLM itself during execution. The scope of each cycle is controlled - each LLM invocation only receives STATE and INSTRUCTIONS. On top of it, a call-stack mechanism provides hierarchical invocation of subroutines with argument passing and return values. This can be used to implement freely multi-agents patterns and also meta-frameworks. To test its "universality", I implemented 14 patterns from the MAS literature, including Tree of Thoughts, LATS, Meta got, ADAS. They share common operatore when possible. is part is still very much a WIP and may be a bit rough around the edges, so feel free to peek and suggest improvements - but I can confirm it's easy to implement and see a new pattern run in no time; all it takes is some prompting to build the INTERPRETER.md. There's also a visualizer that renders the cycles and the subroutines as a graph or each cycle as a log of the machine's state. Let me know what you think, especially if you're working on agents or MAS research.

GitHub - gmlion/TuringLLM: An LLM-powered universal Turing machine · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
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
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
gmlion
/
TuringLLM
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
328 Commits 328 Commits benchmarks benchmarks docs docs interpreters interpreters scripts scripts src src .gitattributes .gitattributes .gitignore .gitignore CLAUDE.md CLAUDE.md INSTRUCTIONS.md INSTRUCTIONS.md LICENSE LICENSE NOTICE NOTICE README.md README.md VISUALIZER.md VISUALIZER.md new-instance.sh new-instance.sh package-lock.json package-lock.json package.json package.json sdlc-poc-artefact-spec.md sdlc-poc-artefact-spec.md sdlc-poc-interpreter-topology.md sdlc-poc-interpreter-topology.md setup-telegram.sh setup-telegram.sh tsconfig.json tsconfig.json visualize.sh visualize.sh visualizer.html visualizer.html View all files Repository files navigation
An LLM-powered universal Turing machine.
Treat an LLM as the step function of a Turing machine. Everything
else falls out: state lives on disk, the program is markdown, runs
are resumable and observable, and "agents" are just user-authored
state machines.
A cycle loop invokes the LLM once per cycle. The LLM reads its state
( MEMORY.md ) and program ( INSTRUCTIONS.md ), matches the first
instruction whose condition fits, acts (typically by rewriting
MEMORY.md ), and is destroyed. The cycle repeats until halt. There
is no hidden conversational state — what the next cycle sees is
exactly what the previous one left on disk.
TuringLLM is, before anything else, the shell : a thin universal
executor that turns a sequence of one-shot LLM calls into a stateful,
resumable, observable machine. It ships a call-stack primitive
(push/pop frames with arg substitution and return-value splicing),
two git repos per instance (one auto-committing every cycle, one
LLM-controlled), and a small set of well-known MEMORY sections the
shell intercepts. Everything in src/ and the top-level scripts is
part of the shell.
Anything that decides what the LLM is supposed to do on each cycle
lives outside the shell, in user-authored markdown.
interpreters/ ships a catalogue of examples —
including runnable implementations of patterns from the multi-agent
systems literature (Self-Refine, Reflexion, Chain-of-Verification,
Plan-and-Execute, Tree of Thoughts, LATS, Multi-Agent Debate,
MetaGPT, ChatDev, AFlow, ADAS) — but the shell itself doesn't know
about them. You can run TuringLLM perfectly well without them.
The pattern catalogue is designed for direct A/B comparison. Some
demos hold the task constant and vary the strategy (ToT vs LATS on
the same Game-of-24 puzzle; MetaGPT vs ChatDev on the same
wc-plus build task); others hold the strategy constant and vary
the task (the three planning-decomposition demos run byte-equal
strategy code on three different PROGRAMs to show three published
patterns collapse to one recursion). See How the catalogue is laid
out for A/B comparison
for the full breakdown. If you're a MAS researcher, the catalogue
is a substrate for comparing patterns without re-implementing the
surrounding scaffolding each time.
Build the shell, then run the deep-research-with-HITL demo — a
recursive researcher that asks you a few clarifying questions
up-front, keeps gathering background while you think, and produces a
structured report at instances/<name>/workspace/report.md .
npm install
npm run build
# Create the instance. PROGRAM.md ships pre-filled with an
# under-specified research goal that invites clarification.
./new-instance.sh interpreters/mas-papers/2-planning-decomposition/c-deep-research-hitl my-cdr-hitl
# Optional but recommended: route the LLM's clarifying questions
# through Telegram so they arrive as push notifications and you can
# answer from your phone. Stdin works too (questions print inline).
./setup-telegram.sh < BOT_TOKEN > instances/my-cdr-hitl # see Human-in-the-loop
# Run
instances/my-cdr-hitl/run.sh
In a second terminal, open the visualizer to watch the call stack
and event stream live as the run unfolds:
./visualize.sh my-cdr-hitl
You'll see the strategy frame f000-plan-execute-clarify write a
few questions to ## Pending Questions , the shell deliver them via
your channel, and the LLM keep cycling — appending background
research notes — while it waits. Once your answers arrive (or the
strategy decides background is exhausted) it pushes tackle.md and
the recursive solver writes the final report.
Provider defaults to Claude Code with Haiku. To swap to the
Anthropic API, an OpenAI-compatible endpoint, Ollama, or a local
GGUF model, edit instances/my-cdr-hitl/.env — see
Providers .
┌─────────────────────────────────────────────────────────────┐
│ INSTANCE │
│ │
│ PROGRAM.md INSTRUCTIONS.md MEMORY.md │
│ ┌─────────────┐ ┌──────────────────┐ ┌───────────┐ │
│ │ # Goal │ │ # Strategy │ │ ## State │ │
│ │ │ │ (active frame's) │ │ current │ │
│ │ (user- │◄───│ │ │ │ │
│ │ authored) │ │ # Sub-instruct. │───►│ ## Result │ │
│ │ │ │ (generated) │ │ ... │ │
│ └─────────────┘ └──────────────────┘ └───────────┘ │
│ │
│ workspace/ history/ logs/ │
│ ┌─────────────┐ ┌──────────────────┐ ┌───────────┐ │
│ │ (project │ │ 0001-a3f1b2c/ │ │ run-*.log │ │
│ │ artifacts, │ │ 0002-b4e2c3d/ │ │ (full │ │
│ │ own git │ │ ... │ │ output) │ │
│ │ repo) │ │ │ │ │ │
│ └─────────────┘ └──────────────────┘ └───────────┘ │
└─────────────────────────────────────────────────────────────┘
│
┌───────────────┼───────────────┐
│ │ │
▼ ▼ ▼
┌──────────┐ ┌──────────┐ ┌──────────────┐
│ Machine │ │ Shell │ │ Provider │
│ Git │ │ main.ts │ │ claude-code, │
│ (auto- │ │ (cycle │ │ api, openai, │
│ commit │ │ loop, │ │ ollama, │
│ per │ │ retry, │ │ local │
│ cycle) │ │ halt) │ │ │
└──────────┘ └──────────┘ └──────────────┘
Three-Layer Design
┌──────────────────────────────────────────────────┐
│ PROGRAM.md (user layer) │
│ │
│ High-level goals. Written by the user. │
│ Never modified by the machine. │
├──────────────────────────────────────────────────┤
│ INSTRUCTIONS.md (per-frame program) │
│ │
│ # Strategy section: immutable meta-program │
│ that interprets PROGRAM.md. Survives rewrites. │
│ │
│ # Sub-instructions section: mutable working │
│ area. Generated and consumed each cycle. │
├──────────────────────────────────────────────────┤
│ Shell (universal executor) │
│ │
│ Single cycle loop. Invokes the LLM, retries on │
│ incomplete cycles, auto-commits, snapshots. │
│ No hardcoded phases — INSTRUCTIONS.md decides. │
└──────────────────────────────────────────────────┘
PROGRAM.md and INSTRUCTIONS.md are both authored outside the shell.
The shell's job is to (a) feed the right files to the LLM at the
right time, and (b) interpret a small set of well-known MEMORY
sections after the call. Everything else is user code.
The shell offers a small, fixed set of primitives. Anything built on
top of them — including the examples in interpreters/ — is user
code.
Identifies the active frame (top of the call stack) and cd s
into it.
Reads that frame's MEMORY.md and INSTRUCTIONS.md , builds a
prompt, invokes the LLM.
The LLM finds the first ## Instruction whose Condition
matches MEMORY's ## State and runs that instruction's
Action , which typically rewrites MEMORY.md .
The shell reads the new MEMORY, processes any well-known sections
it intercepts (see below), commits to git, snapshots history,
and loops. ↻
No conversational continuity is implied across cycles — the LLM
sees exactly what the previous cycle wrote to MEMORY.md (and any
./scoped/ files the user code chose to leave behind).
The shell maintains a per-instance call stack of frames. The
top-of-stack frame is "active": its INSTRUCTIONS.md and
MEMORY.md drive the next cycle. The LLM signals push/pop intent
through MEMORY. The minimal push is a single section naming the
file to load:
## Push
path/to/some-file.md
Push targets can also receive named arguments via ## Push-Args .
Values are key/value pairs or YAML-style block scalars, and any
{{key}} placeholder in the pushed file's text is replaced before
the new frame starts:
## Push
path/to/some-file.md
## Push-Args
<key>: <inline value>
<key>: |
<multi-line block scalar — newlines preserved>
Push. When the shell sees ## Push in MEMORY, before the next
cycle it saves the current frame on the stack and creates a new
active frame loaded from the named file:
caller's memory before push pushed frame after push
─────────────────────────── ───────────────────────
## State ## State
drafted empty
## Push (INSTRUCTIONS.md ← <path>,
<path> with every {{key}} replaced
## Push-Args by the corresponding value)
key: value (./scoped/ starts empty)
The caller is paused while the pushed frame runs through its own
state machine.
Pop. When the pushed frame's state reaches done , the shell
destroys it and resumes the caller. The ## Return block, if
present, in the popped frame's MEMORY is spliced into the caller's MEMORY as top-level
sections (key foo becomes ## Foo ), and the caller's state is
renamed to <caller_state>_completed :
pushed frame's memory before pop caller's memory after pop
──────────────────────────────── ─────────────────────────
## State ## State
done drafted_completed
## Return ## Key
key: value value
The _completed suffix on the state is what lets the caller
distinguish "I just pushed and got a result back" from the original
state where it issued the push — it would otherwise immediately
re-fire the same instruction and push again.
Depth. Push frames can nest arbitrarily — a pushed operator may
itself push another. Two quantities describe how deep a run is, and
the docs use them consistently:
stack.length — the literal number of frames on the stack.
Always ≥ 1, because the root frame is always present.
depth — the nesting level, counted from zero. The root frame
is depth 0 , the frame it pushes is depth 1, and so on:
depth == stack.length - 1 .
So d-cove ( cove.md → verify.md → answer-independently.md )
reaches depth 2 — three frames; a flat interpreter that only ever
pushes one operator reaches depth 1 — two frames. The shell halts
when state == done and only the root frame remains
( stack.length == 1 , i.e. depth 0).
The stack is persisted to .call-stack.json and snapshotted into
every history/ entry.
The convention used by the examples — calling pushable files
"operators" and putting them in operators/ — is purely a user
convention. The shell doesn't care what the

[truncated]
