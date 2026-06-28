---
source: "https://github.com/edonadei/caliper"
hn_url: "https://news.ycombinator.com/item?id=48709606"
title: "Show HN: Caliper – pass@k reliability testing for Claude Code and Codex skills"
article_title: "GitHub - edonadei/caliper: Local-first eval harness for Claude Code, Codex and Pi skills · GitHub"
author: "edonadei"
captured_at: "2026-06-28T18:23:04Z"
capture_tool: "hn-digest"
hn_id: 48709606
score: 1
comments: 0
posted_at: "2026-06-28T17:42:41Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Caliper – pass@k reliability testing for Claude Code and Codex skills

- HN: [48709606](https://news.ycombinator.com/item?id=48709606)
- Source: [github.com](https://github.com/edonadei/caliper)
- Score: 1
- Comments: 0
- Posted: 2026-06-28T17:42:41Z

## Translation

タイトル: HN を表示: Caliper – クロード コードおよびコーデックス スキルの pass@k 信頼性テスト
記事タイトル: GitHub - edonadei/caliper: クロード コード、コーデックス、Pi スキル用のローカルファースト評価ハーネス · GitHub
説明: クロード コード、コーデックス、Pi スキル用のローカル ファースト評価ハーネス - edonadei/caliper
HN テキスト: クロード コードとコーデックスのスキルをテストするのは難しいです。難しいというのは、標準的な方法がないということです。何かについてスキルを一度評価すると、それが機能するように見えます。あなたはそれを公開します。その後、新しいスーパー モデル (GLM 5.2 など?) がリリースされると、一部が静かに壊れますが、ユーザーが苦情を言うまではわかりません。私も同じ問題に直面したので、それを防ぐために軽量なものを構築しようとしました。キャリパー。これは、隔離された環境でスキルを k 回実行し、pass@k スコア (この k 回で何回成功したか) を与える、ローカルで軽量なハーネスです。非決定論的なテクノロジーであるため、「一度は機能した」とだけ言うことはできません。 k回でどれくらい経過したかを答える必要があります。成功は YAML 仕様で定義します。スキーマを保持し、人間が読み取れるようにするために YAML を選択しました。 LLM ジャッジ、Python アサーション、またはその両方を使用します。 JSON 抽出を使用した簡単な評価例を次に示します。これを YAML ファイルに記述します。
- 名前: アクションアイテムをクリーンな JSON として抽出します。
プロンプト: "/tmp/transcript.txt を読み取り、
アクション アイテムを /tmp/actions.json に追加します。」
期待: "すべての項目に次の値が含まれる有効な JSON 配列
所有者、タスク、期限。値下げフェンスはありません。」
アサート: |
jsonをインポートする
items = json.load(open("/tmp/actions.json"))
Assert isinstance(項目, リスト)
assert all({"所有者","タスク","期限"} <= i.keys()
項目内の i の場合)
次に、CLI を使用して、これを実行します。 Caliper run extract-actions.eval.yaml --k 5 --baseline --baseline フラグの優れている点は、次のとおりです。

スキルなしですべてが再実行されるため、スキルが作業を行っているか、ベース エージェントがとにかく通過するかどうかを確認できます: ID Task k(5) pass@k
task-1 アクションアイテムを JSON 5/5 として抽出します 100% PASS
スキル100%で
スキルなし 60%
デルタ +40%
ほとんどのモデルは、ほとんどの場合、JSON を正しく取得する方法を知っています (JSON 抽出はすでに 2 年前に解決されています)。しかし、それだけです、「ほとんどの場合」がバグです。このデルタは、そのスキルが実際にどのように役立ったかを示しています。 (0% の場合もあれば、-100% の場合もあります!) また、お気に入りのハーネスを使用してすぐに開始できる 2 つのスキルも作成しました。 Claude Code、Codex または Pi: - Evaluate-skill: ワークフローを離れることなく eval を実行および管理します。 - Grille-skill: SKILL.md を読み取り、「良い」とは何かについてインタビューし、3 つのタスク仕様 (ハッピー パス、エッジ ケース、敵対的) を作成し、実行します。 コマンドを使用してスキルをインストールできます: npx skill@latest add edonadai/caliper 現在のところ、claude-code、codex、 pi、claude-api、openai-api。エージェントとジャッジを別々のバックエンドとして実行できるため、一方でスキルを実行し、もう一方でジャッジを行うことができます。 GitHub: https://github.com/edonadei/caliper
PyPI: https://pypi.org/project/caliper-eval/ もちろん、これは最初のステップです。自動評価者レイヤーは大幅に改善できると思います。評価仕様の作成と反復の手間が増え、より多くのハーネスがサポートされるようになります。このレイヤーを自己改善のより大きなシステムに組み込んでみてはいかがでしょうか?エージェント評価も構築しているのであれば、それをどのように扱っているかを知りたいと思っています。

記事本文:
GitHub - edonadei/caliper: クロード コード、コーデックス、Pi スキル用のローカルファースト評価ハーネス · GitHub
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
エドナデイ
/
キャリパー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
主要支店

タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
86 コミット 86 コミット .github/ ワークフロー .github/ ワークフロー キャリパー キャリパーの例/ スターターパックの例/ スターターパックのスキル テスト テスト .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md README.md README.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルナビゲーション
Caliper — エージェントのスキルの信頼性テスト
スキルを k 回実行し、追跡して比較できる pass@k スコアを取得し、スキルがベース エージェントに勝っていることを証明します。すでに使用しているエージェント (Claude Code、Codex、または Pi) と連携して動作します。
npx スキル@最新の追加 edonadei/caliper
スキルを使用した場合と使用しない場合で各タスクを実行すると、Caliper は違いを表示します。
ID タスク k (3) pass@k
task-1 従来のコミットメッセージを書き込みます msg 3/3 100% PASS
task-2 有効な構成ファイルを生成します 2/3 96% 合格
スキル 98% で ###################-
スキルなし 55% ###########--------
デルタ +43% アップ
エージェントのスキルをテストするのは困難です。今日、このプロンプトでマシン上で機能するスキルは、モデルの更新または 1 行のプロンプト編集後に明日には機能しなくなる可能性があります。 Caliper は信頼性を測定可能にします。成功とはどのようなものかを定義し、スキルを繰り返し実行し、長期的に追跡できる pass@k スコアを取得します。
Caliper を使用して次のような質問に答えます。
私の即時編集は実際にスキルを向上させましたか?
スキルは機能していますか? それともスキルなしでもベース エージェントは機能しますか?
先週通過したワークフローをまだ通過しますか?
Claude Code、Codex、または Pi のどのエージェントがこのスキルをより確実に実行しますか?
パス A — Agentic (エージェントに運転してもらいます)
npx スキル@最新の追加 edonadei/caliper
2.仕様を対話的に生成する
エージェント (Claude Code または Codex) で:
/グリルスキル ./私のスキル/SKILL.md
Grille-skill は SKILL.md を読み取り、インタビューし、3 つのタスク .eval.yaml (ハッピー パス、エッジ ケース、

敵対的）。
/evaluate-skill run my-skill.eval.yaml --k 3 --baseline
過去の実行を参照:
/評価-スキルリスト
/評価-スキルレポート私のスキル
パス B — CLI (自分で実行)
pipx install Caliper-eval # Python 3.10+が必要
2.仕様を書く
# 私のスキル.eval.yaml
スキル：
パス: ./SKILL.md
バックエンド: クロードコード
裁判官：
バックエンド: クロードコード
タスク:
# 自動評価者 — LLM 裁判官が記録を読んで決定します
- name : 従来のコミットメッセージを書き込みます
プロンプト: " ステージングされた git diff をコミット メッセージとして要約します。"
期待してください: >
応答は従来のコミット メッセージです。つまり、簡潔な件名です。
72 文字未満の行と、その理由を説明する本文が続きます。
何が変わったかだけではなく、変化が生じたのです。
# スクリプトの実行 — 決定論的な Python アサーション
- name : 有効な構成ファイルを生成します。
クリーンアップ: rm -f /tmp/app.config.json
プロンプト: " /tmp/app.config.json に 'ポート' 8080 で構成を生成します。 "
アサート : |
jsonをインポートする
pathlibインポートパスから
data = json.loads(Path("/tmp/app.config.json").read_text())
データをアサート["ポート"] == 8080
Expect: 審査員 LLM によって採点されます。 assert: Python としてローカルで実行されます。いずれかまたは両方を使用してください。
キャリパー実行 my-skill.eval.yaml --k 3 --baseline
4. 出力を読み取る
CALIPER - 私のスキル - k=3 - クロードコード
ID タスク k (3) pass@k
task-1 従来のコミットメッセージを書き込みます msg 3/3 100% PASS
task-2 有効な構成ファイルを生成します 2/3 96% 合格
スキル 98% で ###################-
スキルなし 55% ###########--------
デルタ +43% アップ
結果は .caliper/results/my-skill/2026-06-19T14-23-01Z.json に保存されました
仕様に何を入れるべきかわからないですか?
Eval Starter Pack には 4 つのコピー＆ペーストが含まれています
テンプレート。それぞれが実際のエージェントの失敗 (誤った成功、ツールの誤用、
暴走ループ、プロンプト回帰)。すべてのテンプレートは、
バンドルされたサンプルを編集して自分のスキルをポイントします

2つか3つある
コメント行。
.eval.yaml 仕様
│
▼
Harness ──── エージェント (クロード コード / コーデックス / パイ) に対してスキルを実行します
│
▼
裁判官 ──── LLM 自動評価者および/または決定論的な Python アサーション
│
▼
pass@k スコア + 保存されたトランスクリプト
各試行は、セッション履歴のない、隔離された一時的なホームで実行されます。結果は JSON として保存され、後で検査して比較することができます。
リポジトリには 2 つのエージェント スキルが同梱されています。両方を次のようにインストールします。
npx スキル@最新の追加 edonadei/caliper
Evaluate-skill — eval の実行と管理
通常のワークフロー内から eval を作成、検証、実行、要約できます。別のターミナルは必要ありません。 Caliper が存在しない場合、スキルは Caliper を自動的にインストールします。
または、すでに Caliper がインストールされており、スキルを手動で接続したい場合は、次のようにします。
キャリパーのインストール-スキル クロード-コード
キャリパー取り付けスキルコーデックス
ファイルを書き込まずにプレビューする:
キャリパーのインストールスキル クロードコード --dry-run
次に、それをクロード コードで使用します。
/evaluate-skill run my-skill.eval.yaml --k 3
/evaluate-skill my-skill.eval.yaml を検証する
またはコーデックスでは次のようになります。
Evaluate-skill スキルを使用して、k=3 で my-skill.eval.yaml を実行し、結果を要約します。
グリルスキル — 対話的に eval を作成する
まだ eval を持っていませんか?グリルスキルを使用して、それらを作成することができます。 SKILL.md を読み取り、適切な行動がどのようなものであるかについてインタビューし、3 つのタスク仕様 (ハッピー パス、エッジ ケース、敵対的) を生成します。次に、eval とループを実行します。k=1 で検証、k=3 で測定し、コミットする前にベースラインを作成します。
/グリルスキル ./私のスキル/SKILL.md
すでにスキルのディレクトリに存在する場合、パスは必要ありません。
/グリルスキル
スキルの隣に .eval.yaml がすでに存在する場合、grill-skill は最初から開始するのではなく、既存のタスクを読み取り、ギャップについてインタビューします。
期間
それは何ですか
スペック
スキルを説明する .eval.yaml ファイル

、ジャッジ、および実行するタスク
バックエンド
スキルを実行するエージェント ( claude-code 、 codex 、 pi 、 claude-api 、 openai-api )
裁判官
合格/不合格を決定するもの — トランスクリプトを読み取る LLM (expect: )、Python アサーション (assert: )、またはその両方
pass@k
信頼性スコア: k 回実行し、スキルが成功する頻度を測定します。
ベースライン
スキルを使用せずに同じタスクを再実行して、スキルが作業を行っていることを証明する
試みる
単一タスクの 1 つの分離された実行 - 新しい一時ホーム、セッション履歴なし
バックエンドの選択
バックエンド
必要なもの
こんな方に最適
クロードコード
Claude Code CLI がインストールされ認証されている
Claude Code のスラッシュ コマンド スキルのテスト
コーデックス
Codex CLI がインストールされています ( npm install -g @openai/codex )
コーデックスのスキルをテストする
円周率
pi CLI がインストールされ ( npm install -g @earendil-works/pi-coding-agent )、認証されている
pi スキルのテスト (agentskills.io)
クロード・アピ
ANTHROPIC_API_KEY 環境変数
API ベースのエージェント、CLI は不要
オープンナイAPI
OPENAI_API_KEY 環境変数
OpenAI API エージェント
エージェント バックエンドとジャッジ バックエンドは独立しています。クロード ジャッジまたはその他の組み合わせで Codex スキルをテストできます。
クロード CLI をインストールして認証します。バックエンド: claude-code は既存のクロード コード認証を使用します。追加の構成は必要ありません。
エクスポート ANTHROPIC_API_KEY=...
コーデックスのセットアップ
npm install -g @openai/codex
コーデックスログイン
バックエンド: codex は codex exec を呼び出します。 OpenAI API にはフォールバックしません。 Codex デスクトップ アプリがインストールされている場合、Caliper は PATH 上の codex よりもアプリにバンドルされたバイナリを優先します。 CODEX_CLI_PATH を設定して、特定のバイナリを強制します。
エクスポート OPENAI_API_KEY=...
パイのセットアップ
npm install -g @earendil-works/pi-coding-agent
pi # 次に認証します (例: サブスクリプションプロバイダーの /login、またはプロバイダー API キーを設定します)
バックエンド: pi は pi --print --mode json を実行し、pi の --skill フラグ (エージェントスク

ils.io 標準)。 ~/.pi/agent の認証と設定を再利用します。仕様のモデル: 設定されている pi の設定済みのデフォルトをオーバーライドします。 PI_CLI_PATH を設定して、特定のバイナリを強制します。注: pi の組み込みのデフォルトプロバイダーは google であるため、model: のない仕様は、認証されているプロバイダーを解決するために pi 構成に依存します。
キャリパー更新-cli --check
推奨されるワークフロー
気になる 1 つの動作の仕様を作成します。
仕様を反復しながら --k 1 を指定して実行します。
LLM 審査員が間違っていると推測する可能性がある事実 (ファイル、JSON、コマンド出力) については、assert: を追加します。
タスクが安定したら、--k 3 以上に移行します。
--baseline を追加して、スキルが違いを生み出していることを証明します。
開発者が同じ評価を実行できるように、スキルとともに仕様をコミットします。
/evaluate-skill run my-skill.eval.yaml --k 3 --baseline --verbose
仕様フォーマット
スキル：
path : ./SKILL.md # スキルファイルへのパス (ベースラインのみの実行の場合はオプション)
バックエンド: クロードコード # クロードコード |コーデックス |円周率 |クロード・アピ |オープンナイAPI
model : <model-name> # オプションのモデル オーバーライド
裁判官：
バックエンド: クロードコード # クロードコード |コーデックス |円周率 |クロード・アピ |オープンナイAPI
model : <model-name> # オプションのモデル オーバーライド
サンドボックス:
追加パス :
- 各試行内で ./bin # が PATH の先頭に追加されます
禁止ファイル:
- " .* \\ .eval \\ .yaml$ " # エージェントが仕様を読み取れないようにします
- " ./.caliper/.* " # エージェントが保存された結果を読み取れないようにします
タスク:
- name : 短いタスク名
setup : <shell command> # オプション、各試行の前に実行されます
cleanup : <shell command> # オプション、試行のたびに常に実行されます
プロンプト: <エージェントに送信されるプロンプト>
Expect : <自然言語の成功条件>
アサート : |
# オプションのインライン Python アサーション
True をアサートする
- name : 外部アサーション スクリプトを含むタスク
プロンプト:「レポートを生成する」
アサート: ./assertions/check_report.py
各タスクには次のルールが必要です

Expect または Assert のいずれかとして。タスク ID は、 task-001 、 task-002 などとして自動的に割り当てられます。
裁判官のバックエンドは、試行記録全体を読み取り、期待条件が満たされたかどうかを判断します。バックエンドがツール呼び出しのトレース (クロード コード、コーデックス) をキャプチャすると、それらのトレースも含まれます。裁判官は、最終的なテキストのみに依存することなく、「エージェントがツール X を使用した」などのことを検証できます。
裁判官：
バックエンド: クロードコード
決定論的アサーション (assert:)
Python アサーションはローカルで実行されます。 LLM 裁判官が推測する可能性のある事実については、以下を使用してください。
ファイルが存在する / 正確なファイルの内容
タスク:
- name : 出力ファイルを書き込みます
クリーンアップ: rm -f /tmp/out.txt
プロンプト: " hello world を /tmp/out.txt に書き込みます "
アサート : |
pathlibインポートパスから
パス = パス("/tmp/out.txt")
assert path.exists(), "出力ファイルは作成されませんでした"
アサート path.read_text().strip() == "hello world"
Expect と Assert の両方が存在する場合は、両方ともパスする必要があります。
コマンド
説明
キャリパーラン＜仕様＞
評価スペックを実行する
新しいキャリパー[名前]
インタラクティブなウィザードで新しい仕様を作成する
キャリパー検証 <仕様>
仕様ファイルを検証する
キャリパー一覧【スペック】
仕様と保存された実行のリストを表示します
キャリパーレポート <仕様または結果>
保存された結果を再レンダリングする
キャリパーのインストールスキル <バックエンド>
バンドルされた評価スキルを Claude Code または Codex にインストールします

[切り捨てられた]

## Original Extract

Local-first eval harness for Claude Code, Codex and Pi skills - edonadei/caliper

Skills for Claude Code and Codex are hard to test. What I mean by hard is that there's no standard way to do it. You evaluate the skill once on something, it looks like it works. You publish it. Then the new super model releases (GLM 5.2 anyone?), it will quietly break for some part, and you won't find out until your users complain. I also faced the same problem, so I tried to build something lightweight to stop doing that. Caliper. It's a local and lightweight harness that runs a skill k times in isolated environments and gives you a pass@k score (How much times it succeeded in these k times). As a non-deterministic technology, you can't just say "it worked once". You need to answer how much it passed in k times. You define success in a YAML spec. I picked YAML to keep a schema and make it still readable for a human. You either use a LLM judge, a Python assertion, or both: Here's an simple evaluation example with a JSON extraction, so you write this in a YAML file: tasks:
- name: Extracts action items as clean JSON
prompt: "Read /tmp/transcript.txt and write the
action items to /tmp/actions.json."
expect: "A valid JSON array where every item has
owner, task, due. No markdown fences."
assert: |
import json
items = json.load(open("/tmp/actions.json"))
assert isinstance(items, list)
assert all({"owner","task","due"} <= i.keys()
for i in items)
Then with the CLI, you'll run it: caliper run extract-actions.eval.yaml --k 5 --baseline What's cool about the --baseline flag is that it will re-runs everything without the skill, so you can see whether the skill is doing the work or the base agent was going to pass anyway: ID Task k(5) pass@k
task-1 Extracts action items as JSON 5/5 100% PASS
With skill 100%
No skill 60%
Delta +40%
Most models know how to get the JSON right most of the time (JSON extraction was solved by 2 years old already). But that's it, "most of the time" is the bug. That delta shows how the skill actually helped. (It's sometimes 0%, sometimes -100%!) I also created two skills you can get started right away with your favorite harness, e.g. Claude Code, Codex or Pi: - evaluate-skill: run and manage evals without leaving your workflow - grill-skill: reads your SKILL.md, interviews you about what "good" looks like, writes a 3-task spec (happy path, edge case, adversarial), and runs it You can install the skill with the command: npx skills@latest add edonadei/caliper I for now support claude-code, codex, pi, claude-api, openai-api. You can run the agent and the judge as separate backends, so you can run a skill on one and judge with another. GitHub: https://github.com/edonadei/caliper
PyPI: https://pypi.org/project/caliper-eval/ Of course, it's a first step. I think the autorater layer can be vastly improved, more handholding to create and iterate on evaluation specs, supporting more harness, why not including this layer into a self-improvement bigger system? If you're also building agentic evaluations, I'm genuinely interested to hear how you are handling that.

GitHub - edonadei/caliper: Local-first eval harness for Claude Code, Codex and Pi skills · GitHub
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
edonadei
/
caliper
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
86 Commits 86 Commits .github/ workflows .github/ workflows caliper caliper examples/ starter-pack examples/ starter-pack skills skills tests tests .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md README.md README.md pyproject.toml pyproject.toml View all files Repository files navigation
Caliper — Reliability testing for agent skills
Run your skill k times, get a pass@k score you can track and compare, and prove the skill beats the base agent. Works with the agent you already use — Claude Code, Codex, or Pi.
npx skills@latest add edonadei/caliper
Run each task with and without the skill, and Caliper shows you the difference:
ID Task k (3) pass@k
task-1 Writes a conventional commit msg 3/3 100% PASS
task-2 Generates a valid config file 2/3 96% PASS
With skill 98% ###################-
No skill 55% ###########---------
Delta +43% up
Agent skills are hard to test. A skill that works on your machine, on this prompt, today, might fail tomorrow after a model update or a one-line prompt edit. Caliper makes reliability measurable: define what success looks like, run the skill repeatedly, and get a pass@k score you can track over time.
Use Caliper to answer questions like:
Did my prompt edit actually improve the skill?
Is the skill doing the work, or would the base agent pass without it?
Does it still pass the workflows it passed last week?
Which agent — Claude Code, Codex, or Pi — runs this skill more reliably?
Path A — Agentic (let your agent drive)
npx skills@latest add edonadei/caliper
2. Generate a spec interactively
In your agent (Claude Code or Codex):
/grill-skill ./my-skill/SKILL.md
grill-skill reads your SKILL.md , interviews you, and writes a 3-task .eval.yaml (happy path, edge case, adversarial).
/evaluate-skill run my-skill.eval.yaml --k 3 --baseline
Browse past runs:
/evaluate-skill list
/evaluate-skill report my-skill
Path B — CLI (run it yourself)
pipx install caliper-eval # requires Python 3.10+
2. Write a spec
# my-skill.eval.yaml
skill :
path : ./SKILL.md
backend : claude-code
judge :
backend : claude-code
tasks :
# Autorater — the LLM judge reads the transcript and decides
- name : Writes a conventional commit message
prompt : " Summarize the staged git diff as a commit message. "
expect : >
The response is a conventional-commit message: a concise subject
line under 72 characters, followed by a body explaining why the
change was made, not just what changed.
# Script execution — a deterministic Python assertion
- name : Generates a valid config file
cleanup : rm -f /tmp/app.config.json
prompt : " Generate a config at /tmp/app.config.json with a 'port' of 8080. "
assert : |
import json
from pathlib import Path
data = json.loads(Path("/tmp/app.config.json").read_text())
assert data["port"] == 8080
expect: is graded by the judge LLM; assert: runs locally as Python. Use either or both.
caliper run my-skill.eval.yaml --k 3 --baseline
4. Read the output
CALIPER - my-skill - k=3 - claude-code
ID Task k (3) pass@k
task-1 Writes a conventional commit msg 3/3 100% PASS
task-2 Generates a valid config file 2/3 96% PASS
With skill 98% ###################-
No skill 55% ###########---------
Delta +43% up
Results saved to .caliper/results/my-skill/2026-06-19T14-23-01Z.json
Not sure what to put in a spec?
The Eval Starter Pack has four copy-paste
templates, each catching a real agent failure (false success, tool misuse,
runaway loops, prompt regressions). Every template runs green as-is against a
bundled example, then points at your own skill by editing two or three
commented lines.
.eval.yaml spec
│
▼
Harness ──── runs your skill against the agent (Claude Code / Codex / Pi)
│
▼
Judge ──── LLM autorater and/or deterministic Python assertions
│
▼
pass@k score + saved transcript
Each attempt runs in an isolated temporary home with no session history. Results are saved as JSON you can inspect and diff later.
The repo ships two agent skills. Install both with:
npx skills@latest add edonadei/caliper
evaluate-skill — run and manage evals
Create, validate, run, and summarize evals from inside your normal workflow — no separate terminal needed. The skill installs Caliper automatically if it's missing.
Or, if you already have Caliper installed and want to wire up the skill manually:
caliper install-skill claude-code
caliper install-skill codex
Preview without writing files:
caliper install-skill claude-code --dry-run
Then use it in Claude Code:
/evaluate-skill run my-skill.eval.yaml --k 3
/evaluate-skill validate my-skill.eval.yaml
Or in Codex:
Use the evaluate-skill skill to run my-skill.eval.yaml with k=3 and summarize the result.
grill-skill — create evals interactively
Don't have evals yet? grill-skill guides you through creating them. It reads your SKILL.md , interviews you about what good behavior looks like, and generates a 3-task spec (happy path, edge case, adversarial). Then it runs the eval and loops — k=1 to validate, k=3 to measure, baseline before you commit.
/grill-skill ./my-skill/SKILL.md
No path needed if you're already in the skill's directory:
/grill-skill
If an .eval.yaml already exists next to your skill, grill-skill reads the existing tasks and interviews you about gaps instead of starting from scratch.
Term
What it is
Spec
A .eval.yaml file that describes the skill, judge, and tasks to run
Backend
The agent that executes the skill ( claude-code , codex , pi , claude-api , openai-api )
Judge
What decides pass/fail — an LLM reading the transcript ( expect: ), Python assertions ( assert: ), or both
pass@k
Reliability score: run k times, measure how often the skill succeeds
Baseline
Re-run the same tasks without the skill to prove the skill is doing the work
Attempt
One isolated run of a single task — fresh temporary home, no session history
Choosing a backend
Backend
Requires
Best for
claude-code
Claude Code CLI installed and authenticated
Testing Claude Code slash-command skills
codex
Codex CLI installed ( npm install -g @openai/codex )
Testing Codex skills
pi
pi CLI installed ( npm install -g @earendil-works/pi-coding-agent ) and authenticated
Testing pi skills (agentskills.io)
claude-api
ANTHROPIC_API_KEY env var
API-backed agents, no CLI needed
openai-api
OPENAI_API_KEY env var
OpenAI API agents
The agent backend and judge backend are independent — you can test a Codex skill with a Claude judge, or any other combination.
Install and authenticate the claude CLI. backend: claude-code uses your existing Claude Code auth — no extra configuration needed.
export ANTHROPIC_API_KEY=...
Codex setup
npm install -g @openai/codex
codex login
backend: codex calls codex exec . It does not fall back to the OpenAI API. If the Codex desktop app is installed, Caliper prefers the app-bundled binary over codex on PATH . Set CODEX_CLI_PATH to force a specific binary.
export OPENAI_API_KEY=...
pi setup
npm install -g @earendil-works/pi-coding-agent
pi # then authenticate (e.g. /login for a subscription provider, or set the provider API key)
backend: pi runs pi --print --mode json and loads the skill natively via pi's --skill flag (the agentskills.io standard). It reuses your ~/.pi/agent auth and settings — the spec's model: overrides pi's configured default when set. Set PI_CLI_PATH to force a specific binary. Note: pi's built-in default provider is google , so a spec with no model: relies on your pi config to resolve a provider you are authenticated for.
caliper update-cli --check
Recommended workflow
Create a spec for one behavior you care about.
Run with --k 1 while iterating on the spec.
Add assert: for facts an LLM judge might guess wrong (files, JSON, command output).
Move to --k 3 or higher once the task is stable.
Add --baseline to prove the skill is making a difference.
Commit the spec alongside the skill so contributors can run the same eval.
/evaluate-skill run my-skill.eval.yaml --k 3 --baseline --verbose
Spec format
skill :
path : ./SKILL.md # path to the skill file (optional for baseline-only runs)
backend : claude-code # claude-code | codex | pi | claude-api | openai-api
model : <model-name> # optional model override
judge :
backend : claude-code # claude-code | codex | pi | claude-api | openai-api
model : <model-name> # optional model override
sandbox :
extra_path :
- ./bin # prepended to PATH inside each attempt
forbidden_files :
- " .* \\ .eval \\ .yaml$ " # prevents agent from reading the spec
- " ./.caliper/.* " # prevents agent from reading saved results
tasks :
- name : Short task name
setup : <shell command> # optional, runs before each attempt
cleanup : <shell command> # optional, always runs after each attempt
prompt : <prompt sent to the agent>
expect : <natural-language success condition>
assert : |
# optional inline Python assertion
assert True
- name : Task with external assertion script
prompt : " Generate a report "
assert : ./assertions/check_report.py
Each task needs at least one of expect or assert . Task IDs are assigned automatically as task-001 , task-002 , and so on.
The judge backend reads the full attempt transcript and decides whether the expect condition was met. When the backend captures tool-call traces (Claude Code, Codex), those traces are included — the judge can verify things like "the agent used tool X" without relying on the final text alone.
judge :
backend : claude-code
Deterministic assertions ( assert: )
Python assertions run locally. Use these for facts the LLM judge might guess:
file exists / exact file contents
tasks :
- name : Writes an output file
cleanup : rm -f /tmp/out.txt
prompt : " Write hello world to /tmp/out.txt "
assert : |
from pathlib import Path
path = Path("/tmp/out.txt")
assert path.exists(), "Output file was not created"
assert path.read_text().strip() == "hello world"
When both expect and assert are present, both must pass.
Command
Description
caliper run <spec>
Run an evaluation spec
caliper new [name]
Create a new spec with the interactive wizard
caliper validate <spec>
Validate a spec file
caliper list [spec]
List specs and saved runs
caliper report <spec-or-result>
Re-render saved results
caliper install-skill <backend>
Install the bundled evaluate-skill into Claude Code or Codex

[truncated]
