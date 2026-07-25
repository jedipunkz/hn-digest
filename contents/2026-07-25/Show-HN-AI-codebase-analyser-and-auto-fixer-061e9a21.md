---
source: "https://www.getdebug.dev/blog/python-ai-app-prefilters"
hn_url: "https://news.ycombinator.com/item?id=49045992"
title: "Show HN: AI codebase analyser and auto-fixer"
article_title: "Python AI-app static analysis: what catches what — getdebug"
author: "nutifafa"
captured_at: "2026-07-25T09:28:03Z"
capture_tool: "hn-digest"
hn_id: 49045992
score: 1
comments: 0
posted_at: "2026-07-25T09:23:31Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AI codebase analyser and auto-fixer

- HN: [49045992](https://news.ycombinator.com/item?id=49045992)
- Source: [www.getdebug.dev](https://www.getdebug.dev/blog/python-ai-app-prefilters)
- Score: 1
- Comments: 0
- Posted: 2026-07-25T09:23:31Z

## Translation

タイトル: Show HN: AI コードベース アナライザーと自動修正ツール
記事のタイトル: Python AI アプリの静的分析: 何が何をキャッチするのか — getdebug
説明: Python AI アプリの正規表現プレフィルターを getdebug に追加し、Bandit、Semgrep、および (試行された) vulnhuntr に対してベンチマークを実行しました。ここでは、各ツールが実際に物事を捕捉する場所と捕捉しない場所を示します。
HN テキスト: 違いは、デバッグに対する「人間参加型」アプローチの考え方の変化です。私はセキュリティが非常に重要であるため、AI があらゆるアクションをエンドツーエンドで実行すべきではないと考えています。

記事本文:
Python AI アプリの静的分析: 何が何をキャッチするか — getdebug [ 200 OK ] [ ANALYZE ] [ .SARIF ] [ FIX-PR ] getdebug Docs Bench ↗ ブログの価格 サインインについて 始めましょう getdebug Docs Bench ブログの価格 テーマ アカウントについて
すべての投稿 2026 年 6 月 5 日 · Python · ai-app-security · ベンチマーク Python AI アプリの静的分析: 何が何をキャッチするのか
by fafa — 私たちがテストした 4 つのツールのそれぞれが実際に何を発見したかについて
getdebug CLI 0.4.0 には、0.3.0 でリリースされた JS/TS プレフィルターとともに Python AI アプリ正規表現プレフィルターが同梱されています。 5 つのカテゴリ: プロンプトインジェクション、安全でないロールマージ、pii-in-プロンプト、無制限ストリーム、安全でないツール出力。すべて決定論的。 LLM 呼び出しはありません。デフォルトの getdebug 分析。 Python リポジトリではミリ秒で実行され、コストはゼロで、常に実行されていたシークレットと依存関係 CVE パスに加えて、Python AI アプリのアンチパターンもカバーされるようになりました。
興味深い部分はパターンそのものではありません。パターンはよく知られた SDK イディオム (messages=[ { "role": "system"... } ] 、 stream=True 、 subprocess.run(tool_call.input.command) ) の単純な正規表現です。興味深いのは、開発中に他の 3 つのツールと比較して実行したベンチマークです。
4 つのツール、正直な地図
Bandit (PyCQA) は、Python-OSS の標準セキュリティ リンターです。手書きのルール。無料、高速、LLM なし。パイソンのみ。
Semgrep は、コミュニティ ルール パックを備えた多言語 SAST です。手書きのルール。無料、高速、LLM なし。 getdebug と同じ製品形状。
vulnhuntr (Protect AI、オープンソース) は、AI アプリの静的分析のカテゴリー リーダーとして認められています。 LLM 主導、Python のみ、エントリポイント検出ベース。
getdebug は、JS/TS + Python のパターンベースの正規表現プレフィルターに加えて、オプションの Ollama 経由のローカル LLM SAST パス (無料、デバイス上) と Claude 経由のホスト型 LLM SAST パス (有料) です。パイス

0.4.0 での追加は、特に正規表現レイヤーです。
テスト 1 — 脆弱性/安全なフィクスチャのペア (10 ファイル)
脆弱な 5 つと安全な 5 つの Python AI アプリ フィクスチャを、カテゴリごとに 1 ペアずつ作成しました。既存のJS/TSコーパスと同じ形状です。 4 つのツールはすべて同じセットで実行されました。
ツール TP FP FN プレシジョン リコール
取得デバッグ 5 1 0 83% 100%
盗賊 1 1 4 50% 20%
セムグレップ 1 1 4 50% 20%
vulnhuntr — — — (0 個のファイルが選択されています。以下を参照) Bandit と Semgrep はどちらも、汎用の subprocess.run(shell=True) ルールを介して unsafe-tool-output フィクスチャを起動します。これは、脆弱な亜種に対する真の陽性です。ただし、同じフィクスチャの安全なバリアント、つまりモデル出力がシェルに到達する前にホワイトリストを通じてマッピングされるバリアントでも起動されます。
# 安全なパターン — Bandit と Semgrep の両方がこれを FP としてフラグを立てます
ALLOWED = {"ホスト": "cat /etc/hosts", "稼働時間": "稼働時間"}
デフォルトハンドル(tool_call):
cmd = ALLOWED.get(tool_call.input.tag)
そうでない場合は cmd:「拒否されました」を返します
return subprocess.run(cmd,shell=True,capture_output=True).stdout どちらのツールも、cmd がモデルからではなく静的辞書から来たことを認識しません。彼らはshell=Trueを見て発火します。 getdebug の正規表現では、tool_call.input.X / block.input.X 参照をシンク引数に含める必要があります。これにより、汎用 SAST ツールではできない誤検知の範囲が狭まります。正直なアップデート: 0.5.5 で追加された scanPyShellToolArg 検出器は、この同じ安全なフィクスチャで CRITICAL を発火するようになりました。これは、tool_call.input.tag がshell=True シンクに到達していることを確認しますが、その値が最初に静的ホワイトリストを介してルーティングされたかどうかはわかりません。 0.5.5 でコーパスを再実行すると、このセットでは 5 TP / 1 FP / 0 FN (精度 83%) でデバッグが得られますが、元の 0.4.0 の数値が示したクリーン スイープではありません。このホワイトリストのケースの厳格化が追跡されます。
どちらのツールも、他の 4 つの行動カテゴリを完全に欠落しています — pii-in-prompt

、unsafe-role-merge、プロンプトインジェクション、unbounded-stream。それらは彼らのために設計されたものではありません。ルール パックには、 {"role": "system", "content": f"...{'$'}{name}..."} のパターンは含まれません。それが私たちの追加です。
テスト 2 — 現実世界の信号/ノイズ チェック
合成フィクスチャは動作を固定します。セキュリティ スキャナの実際のテストは、実際のコードで何を行うかです。 3 つの (動作する) ツールすべてを simonw/llm に対して実行しました。これは、LLM と通信するための Simon Willison のクリーンでよく管理された CLI、49 の Python ファイル (コミット 0d593ea; 生のツール出力は Bench/realworld/) です。
ツール 合計結果 シグナル
Bandit 1,261 1,228 は 'assert_used' (pytest);
AI アプリの適用範囲はゼロ
semgrep 3 3 つの汎用 SAST ヒット。
AI アプリの適用範囲はゼロ
getdebug 7 6 AI アプリの検出結果 (1 つのプロンプト インジェクション、
5 unbounded-stream) + 1 borderline role-merge この 49 ファイルのコードベースに関する Bandit の 1,261 件の検出結果は、ほぼ完全にノイズです。そのうち 1,228 件は、pytest アサーションに関するassert_used 警告です。これは長年にわたる Bandit の苦情です。デフォルト設定では、アサートが python -O で消えるため、すべてのアサートにセキュリティの匂いとしてフラグが立てられます。 -O 最適化をオプトアウトしている運用アプリ (ほぼ全員) の場合、これは純粋なノイズです。
Semgrep の 3 つの発見は現実的ですが一般的です。cli.py でフラグが付けられた exec() の使用法 (プラグイン システム用に意図的)、静的 HTML アセットの整合性属性の欠落、および非リテラル インポートです。これらはいずれも AI アプリに固有のものではありません。
getdebug の 7 つの検出結果のうち 6 つは AI アプリに分類されます。1 つはプロンプト インジェクション (ユーザーが指定した 2 つの文字列を連結する CLI のテンプレート機能)、および OpenAI プラグインの 5 つの無制限ストリーム ヒット (各ストリーム = True、スコープ内に with ブロックまたはタイムアウトなし)。脅威モデルに応じて、両方とも TP であると議論できます。これはシングルユーザー CLI であるため、2 つの CLI ユーザー入力間のプロンプト挿入は

uts は無理があり、ライブラリはストリーム管理をそのコンシューマに委任します。 7 番目の {"role": role} でのロールのマージは、境界線の誤検知です。ロールは、リクエスト入力ではなく、API 自体のストリーミング応答をエコーし​​ます。これは、優先順位を付けて排除する種類のものです。すべてに CRITICAL ではなく HIGH / MEDIUM のフラグが付けられ、トリアージのコンテキストが示されます。 CLI は、.getdebug-ignore ファイルを使用して検出結果を抑制します (インライン // getdebug:ignore 注釈はホスト スキャン機能です)。
vulnhuntr は、LLM 駆動の AI アプリ静的分析のカテゴリ リーダーであると明記されているため、現在のリリースを再確認しました。 2 つの調査結果があり、私たちは両方について公平であることを望んでいます。
走ります。 vulnhuntr 1.2.2 は Python 3.11 にインストールされ、正常に起動します。古い ModuleNotFoundError はなくなりました。他の LLM 駆動ツールと同様に、プロバイダー (API キー、または --llm ollama を介したローカル Ollama) が必要です。
しかし、ここでは何も見えません。そのファイル選択ヒューリスティックは「ネットワークに公開された」エントリ ポイントをターゲットにしており、simonw/llm は Web アプリではなく CLI であるため、vulnhuntr は分析するファイルを選択せず​​、このリポジトリには何も表示しません。独自の --dry-run で確認済み (ファイル 0、$0.00、キー不要)。
したがって、vulnhuntr は壊れていません。Web アプリのエントリ ポイントに範囲が限定されているため、このような CLI ライブラリでは何も操作できません。異なるツール、異なるターゲット: getdebug の AI アプリ プレフィルターは、Web アプリに関係なく、LLM に触れるあらゆる Python 上で実行されます。それが私たちが取り組んでいるギャップです。
LLM (チャット ラッパー、エージェント フレームワーク、ツール呼び出しバックエンド、バッチ サマライザー) を呼び出す Python アプリを出荷する場合は、次の 3 つをすべて実行する必要があります。
バンディット -r 。一般的な Python セキュリティ衛生用 (最初に .bandit 設定を介してassert_used をオフにします)。
semgrep --config auto 。言語を超えた SAST カバレッジについては。
npx @getdebug/cli@0.4.0 分析 。 AI アプリの行動パターンについては、他の 2 つは当てはまりません

ゲット。
それらは補完的なものです。それらはどれも他のものを包含しません。最初の 2 つは、一般的な SAST と Python の衛生状態を示しています。 getdebug は、「ユーザー オブジェクト全体を LLM プロンプトにシリアル化する」クラスのバグを捕捉します。これらのバグは、人間工学に苦労することなく、汎用 SAST で持続可能なルールを手動で記述することができません。
このページの getdebug.dev/bench にあるすべての番号を再現します。コーパス、方法論、ハーネスは GitHub の CodeSecBench で公開されています。
そしていつものように、再現できない結果が表示された場合、または getdebug がキャッチするはずのパターンがキャッチされなかった場合、コーパスはオープンです。PR は歓迎です。ハーネスは再現可能で、方法論は文書化されています。それがベンチマークを公の場で行うことの要点です。
npx @getdebug/cli@0.4.0 分析 。
または Homebrew 経由: brew install getdebug-ai/tap/getdebug
— Fafa Agbetsise / getdebug.dev 創設者
© 2026 getdebug · 出荷前にテスト

## Original Extract

We added Python AI-app regex prefilters to getdebug and benchmarked against Bandit, Semgrep, and (attempted) vulnhuntr. Here's where each tool actually catches things — and where they don't.

The difference is the mindset shift of "human-in-the-loop" approach to debugging. I believe security is so critical that AI should never execute every action end to end.

Python AI-app static analysis: what catches what — getdebug [ 200 OK ] [ ANALYZE ] [ .SARIF ] [ FIX-PR ] getdebug Docs Bench ↗ Blog Pricing About Sign in Get started getdebug Docs Bench Blog Pricing About Theme Account
All posts June 5, 2026 · python · ai-app-security · benchmark Python AI-app static analysis: what catches what
by fafa — on what each of the four tools we tested actually finds
getdebug CLI 0.4.0 ships Python AI-app regex prefilters alongside the JS/TS ones that landed in 0.3.0. Five categories: prompt-injection, unsafe-role-merge, pii-in-prompt, unbounded-stream, unsafe-tool-output. All deterministic. No LLM call. The default getdebug analyze . on a Python repo runs in milliseconds, costs zero, and now covers Python AI-app anti-patterns alongside the secrets + dependency-CVE passes it always ran.
The interesting part isn't the patterns themselves — they're straightforward regex on familiar SDK idioms ( messages=[ { "role": "system"... } ] , stream=True , subprocess.run(tool_call.input.command) ). The interesting part is the benchmark we ran while developing them, against the three other tools anyone would compare us to.
The four tools, the honest map
Bandit (PyCQA) is the Python-OSS standard security linter. Hand-written rules. Free, fast, no LLM. Python only.
Semgrep is multi-language SAST with community rule packs. Hand-written rules. Free, fast, no LLM. Same product shape as getdebug.
vulnhuntr (Protect AI, open source) is the stated category leader for AI-app static analysis. LLM-driven, Python-only, entry-point-detection based.
getdebug is what we ship: pattern-based regex prefilters in JS/TS + Python now, plus an optional local-LLM SAST pass via Ollama (free, on-device) and a hosted LLM SAST pass via Claude (paid). The Python additions in 0.4.0 are the regex layer specifically.
Test 1 — paired vulnerable/safe fixtures (10 files)
We wrote 5 vulnerable + 5 safe Python AI-app fixtures, one pair per category. Same shape as our existing JS/TS corpus. All four tools ran on the same set.
Tool TP FP FN Precision Recall
getdebug 5 1 0 83% 100%
bandit 1 1 4 50% 20%
semgrep 1 1 4 50% 20%
vulnhuntr — — — (0 files selected; see below) Bandit and Semgrep both fire on the unsafe-tool-output fixture via their generic subprocess.run(shell=True) rules — that's a true positive on the vulnerable variant. But they also fire on the safe variant of the same fixture, the one where the model output is mapped through an allowlist before reaching the shell:
# Safe pattern — Bandit + Semgrep both flag this as a FP
ALLOWED = {"hosts": "cat /etc/hosts", "uptime": "uptime"}
def handle(tool_call):
cmd = ALLOWED.get(tool_call.input.tag)
if not cmd: return "rejected"
return subprocess.run(cmd, shell=True, capture_output=True).stdout Neither tool knows that cmd came from a static dict, not from the model. They see shell=True and fire. getdebug's regex requires the tool_call.input.X / block.input.X reference to appear in the sink arg, which narrows the false-positive surface the generic SAST tools don't. Honesty update: the scanPyShellToolArg detector added in 0.5.5 now fires a CRITICAL on this same safe fixture — it sees tool_call.input.tag reaching a shell=True sink and can't tell the value was routed through the static allowlist first. Re-running the corpus on 0.5.5 scores getdebug at 5 TP / 1 FP / 0 FN (83% precision) on this set, not the clean sweep the original 0.4.0 numbers showed. Tightening this allowlist case is tracked.
Both tools miss the other four behavioural categories entirely — pii-in-prompt, unsafe-role-merge, prompt-injection, unbounded-stream. They aren't designed for them; the rule packs don't contain patterns for {"role": "system", "content": f"...{'$'}{name}..."} . That's our addition.
Test 2 — the real-world signal/noise check
Synthetic fixtures lock the behaviour in. The real test for a security scanner is what it does on actual code. We ran all three (working) tools against simonw/llm — Simon Willison's clean, well-maintained CLI for talking to LLMs, 49 Python files (commit 0d593ea; raw tool outputs in bench/realworld/).
Tool Total findings Signal
bandit 1,261 1,228 are 'assert_used' (pytest);
zero AI-app coverage
semgrep 3 3 generic-SAST hits;
zero AI-app coverage
getdebug 7 6 AI-app findings (1 prompt-injection,
5 unbounded-stream) + 1 borderline role-merge Bandit's 1,261 findings on this 49-file codebase is almost entirely noise : 1,228 of them are assert_used warnings on pytest assertions. This is a long-standing Bandit complaint — the default config flags every assert as a security smell because asserts disappear under python -O . For a production app where you opt-out of -O optimisation (which is almost everyone), this is pure noise.
Semgrep's 3 findings are real but generic: an exec() usage flagged in cli.py (intentional for the plugin system), a missing-integrity attribute on a static HTML asset, and a non-literal import. None of these are AI-app specific.
Six of getdebug's seven findings are AI-app categorized: one prompt-injection (the CLI's template feature concatenating two user-supplied strings) and five unbounded-stream hits in the OpenAI plugin (each stream=True with no with block or timeout in scope). Both arguable as TPs depending on threat model — this is a single-user CLI, so the prompt-injection between two CLI-user inputs is a stretch, and the library delegates stream-management to its consumer. The seventh, a role-merge on {"role": role} , is a borderline false positive — the role echoes the API's own streaming response, not request input — the kind of thing you'd triage away. Everything is flagged HIGH / MEDIUM , not CRITICAL , with context to triage. The CLI suppresses findings with a .getdebug-ignore file (inline // getdebug:ignore annotations are a hosted-scan feature).
vulnhuntr is the stated category leader for LLM-driven AI-app static analysis, so we re-checked the current release. Two findings, and we want to be fair about both:
It runs. vulnhuntr 1.2.2 installs and starts cleanly on Python 3.11 — the older ModuleNotFoundError is gone. Like any LLM-driven tool it needs a provider (an API key, or a local Ollama via --llm ollama ).
But it sees nothing here. Its file-selection heuristic targets “network-exposed” entry points, and simonw/llm is a CLI, not a web app — so vulnhuntr selects zero files to analyse and surfaces nothing on this repo. Confirmed with its own --dry-run (0 files, $0.00, no key needed).
So vulnhuntr isn't broken — it's scoped to web-app entry points, which makes it a no-op on a CLI library like this one. Different tool, different target: getdebug's AI-app prefilters run on any Python that touches an LLM, web app or not. That's the gap we're shipping into.
If you ship a Python app that calls an LLM — chat wrapper, agent framework, tool-calling backend, batch summariser — you should run all three :
bandit -r . for general Python security hygiene (turn off assert_used via .bandit config first).
semgrep --config auto . for cross-language SAST coverage.
npx @getdebug/cli@0.4.0 analyze . for the AI-app behavioural patterns the other two don't target.
They're complementary. None of them subsume the others. The first two catch general SAST and Python hygiene; getdebug catches the “serialised the whole user object into the LLM prompt” class of bugs that you can't hand-write a sustainable rule for in generic SAST without painful ergonomics.
Reproduce every number on this page at getdebug.dev/bench . Corpus, methodology, and harness are open — CodeSecBench on GitHub.
And as always: if you see a result you can't reproduce, or a pattern getdebug should catch but doesn't, the corpus is open — PRs welcome, harness is reproducible, methodology is documented. That's the whole point of doing the benchmark in public.
npx @getdebug/cli@0.4.0 analyze .
Or via Homebrew: brew install getdebug-ai/tap/getdebug
— Fafa Agbetsise / Founder, getdebug.dev
© 2026 getdebug · test before you ship
