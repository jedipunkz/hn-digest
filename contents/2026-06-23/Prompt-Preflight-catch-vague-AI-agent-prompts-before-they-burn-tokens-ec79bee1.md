---
source: "https://github.com/akg268/prompt-preflight/"
hn_url: "https://news.ycombinator.com/item?id=48639577"
title: "Prompt Preflight – catch vague AI-agent prompts before they burn tokens"
article_title: "GitHub - akg268/prompt-preflight: Local Codex plugin that catches vague AI-agent prompts before they burn tokens. · GitHub"
author: "akg268"
captured_at: "2026-06-23T03:03:06Z"
capture_tool: "hn-digest"
hn_id: 48639577
score: 2
comments: 0
posted_at: "2026-06-23T02:49:18Z"
tags:
  - hacker-news
  - translated
---

# Prompt Preflight – catch vague AI-agent prompts before they burn tokens

- HN: [48639577](https://news.ycombinator.com/item?id=48639577)
- Source: [github.com](https://github.com/akg268/prompt-preflight/)
- Score: 2
- Comments: 0
- Posted: 2026-06-23T02:49:18Z

## Translation

タイトル: プロンプト プリフライト – トークンを書き込む前に AI エージェントの曖昧なプロンプトを捕捉する
記事のタイトル: GitHub - akg268/prompt-preflight: トークンを書き込む前にあいまいな AI エージェント プロンプトをキャッチするローカル Codex プラグイン。 · GitHub
説明: トークンを書き込む前にあいまいな AI エージェント プロンプトを捕捉するローカル Codex プラグイン。 - akg268/プロンプトプリフライト

記事本文:
GitHub - akg268/prompt-preflight: トークンを書き込む前に曖昧な AI エージェント プロンプトをキャッチするローカル Codex プラグイン。 · GitHub
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
akg268
/
プロンプト-プリフライト
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビ

オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット .codex-plugin .codex-plugin .github/ workflows .github/ workflows docs docs フック フック スクリプト スクリプト スキル/ プロンプトプリフライト スキル/ プロンプトプリフライト src/ プロンプト_プリフライト src/ プロンプト_プリフライト テスト テスト .gitignore .gitignore .prompt-preflight.example.json .prompt-preflight.example.json ライセンス ライセンス README.md README.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
高価なモデルターンになる前に、指定不足のリクエストをキャッチします。
Prompt Preflight は、プロンプトが動作するのに十分な具体性があるかどうかをチェックするローカル Codex プラグインおよびスタンドアロン CLI です。曖昧さと間違いによるコストの両方が高い場合、リクエストは一時停止され、ユーザーに次の情報が表示されます。
ドメインを意識したより強力なプロンプトの例。
最も重要なギャップを埋めるための最大 3 つの質問。
このチェックでは決定的な Python ルールが使用されます。ネットワーク要求は行わず、モデルも呼び出しません。
プロンプト プリフライトは、モデルの作業が始まる前に漠然とした Codex プロンプトをキャッチします。
ユーザーが曖昧なリクエストを送信する
→ プロンプト プリフライトはローカルで実行されます
→ モデルターンを消費する前にコーデックスがブロックされる
→ ユーザーはより強力なプロンプト テンプレートと的を絞った質問を受け取ります
なぜこれが存在するのか
曖昧なプロンプトは、多くの場合、負荷の高いループを作成します。
漠然とした要望
→ モデルはプロジェクトのコンテキストを読み取ります
→ モデルが間違った解釈を生み出す
→ ユーザーが修正します
→ モデルは展開された会話を読み上げます
→ モデルが再び作業を行う
無駄なコストは最初の答えに限定されません。再試行には、以前のプロンプト、出力、修正、および追加のコンテキストも含まれます。
プロンプト プリフライトは、ループの前に説明を移動します。
漠然とした要望
→ ローカルプリフライトチェック
→ 対象を絞った明確化
→もう一つ強い要望
→私たち

豊富なモデル作品
トークンあたりの価格は減額されません。これにより、回避可能なモデル入力、不要な出力、繰り返しのツール作業、および修正回転が削減されます。
プリフライトがない場合、失敗した試行と再試行のおおよそのコストは次のとおりです。
失敗した入力 + 失敗した出力 + 修正コンテキスト + 置換入力 + 置換出力
プリフライトを使用すると、ローカル チェックでモデル トークンが消費されません。意図したパスは次のようになります。
明確なインプット + 有用なアウトプット
したがって、回避される可能性のあるトークンはおよそ次のようになります。
失敗した入力 + 失敗した出力 + 修正コンテキスト + 重複した作業
実際の節約量は、プロンプトの品質、モデルの動作、コンテキストのサイズ、タスクの複雑さによって異なります。 Prompt Preflight は現在、固定の節約率を主張していません。測定されたトークンテレメトリーは今後の課題です。
最大の利点は、リポジトリ全体の変更、移行、デプロイメント、アーキテクチャ作業、または反復的なイメージ生成など、誤った解釈によってコストがかかるタスクで期待されます。
車のイメージを作成する
画像生成が開始される前に、プロンプト プリフライトが応答します。
プロンプト:
「クルマのイメージを作る」
尋ねてみてください:
「車の[フォトリアリスティック/イラスト/3D]画像を作成します
[設定/背景]の[キーカラー、素材、特徴的なディテール]、
【カメラアングル・構図】から見た【照明・雰囲気】、
[アスペクト比]で。」
次のように答えて括弧内を埋めてください。
1. 車はどのように見えるべきですか?
2. どのようなビジュアルスタイルと雰囲気を望みますか?
3. どのような設定、構図、照明、アスペクト比を使用する必要がありますか?
これにより、任意の最初の画像に続いて数回の視覚的な修正が行われることがなくなります。
プロンプトが明確になった後の出力例:
ダッシュボードをより良くする
プロンプト プリフライトは次のことを提案します。
[目に見える結果] が得られるように、[特定のページ/コンポーネント] のダッシュボードを改善します。
[重要な動作または設計上の制約] を変更しないでください。

[テストまたは合格基準] で確認します。
モデルは、ファイルを読み取ったりコードを編集したりする前に、ターゲット、結果、境界、および完了の定義を受け取ります。
Codex モデルが UserPromptSubmit を通過する前に実行されます。
モデル、API キー、ネットワーク アクセス、外部サービスは使用しません。
フィードバックを選択する前に、ドメインごとにプロンプ​​トをルーティングします。
ソフトウェアと画像生成フィードバック プロファイルが含まれます。
「より具体的に」と言うだけでなく、カスタマイズされた書き換えを示します。
価値の高い質問を最大 3 つ尋ねます。
明確なプロンプトと会話によるフォローアップを通過させます。
1 回限りの [preflight:skip] バイパスをサポートします。
構成可能なブロック モードとナッジ モードをサポートします。
フック入力の形式が不正な場合はフェールオープンします。
評価とデバッグ用に構造化された JSON を提供します。
プロンプト プリフライトでは、次の 3 つのことが推定されます。
意図：どのような仕事を依頼されていますか？
あいまいさ: ドメイン固有のどの詳細が欠落していますか?
影響: 間違った解釈をすると、どれくらいの損害が発生するでしょうか?
プロンプトが実行可能で、曖昧さと影響の両方が設定されたしきい値を超えた場合にのみ中断されます。これにより、プラグインが簡単な質問、確認、またはすでに具体的な作業についてユーザーに質問するのを防ぎます。
現在のドメイン プロファイルには次のものが含まれます。
サポートされていないドメインでは、ソフトウェア固有の質問を受信するのではなく、保守的なフォールバックが使用されます。
Python 3.10 以降が必要です。
python3 scripts/prompt_preflight.py " 車の画像を作成する "
説明を必要とするプロンプトはステータス 2 で終了します。送信準備完了のプロンプトはステータス 0 で終了します。
python3 スクリプト/prompt_preflight.py \
「夜の濡れた東京の路上で、低いカメラアングル、映画のような照明、16:9 で赤い 1967 年型フォード マスタングのフォトリアリスティックな画像を作成します。」
完全な分析を検査します。
python3 scripts/prompt_preflight.py --json " プロジェクト全体を書き直す "
構造化出力には検出が含まれます

d 意図、曖昧さスコア、影響スコア、理由、質問、および提案されたプロンプト。
ベンチマークの曖昧なプロンプト検出
プロンプト プリフライトには、ソフトウェア作業、バグ修正、展開、移行、最適化、イメージ生成にわたる 100 個の意図的に曖昧なプロンプトの固定ベンチマークが含まれています。
python3 スクリプト/benchmark_vague_prompts.py
完全な結果を JSON として保存します。
python3 スクリプト/benchmark_vague_prompts.py \
--min-block-rate 0.90 \
--json-output ベンチマーク結果.json
ベンチマークは次のようにレポートします。
モデル作業前にブロックされた曖昧なプロンプトの数
確認する必要があるプロンプトを見逃した
曖昧さ、影響、明確化の平均スコア
検出された意図ごとにグループ化された結果
最初のベンチマークが教えてくれたこと
最初の 100 個のプロンプトの実行により、このプロジェクトが捕らえようとしている種類の回帰リスクが正確に明らかになりました。次のような短いアクション プロンプトでは、初期のスコアリングが甘すぎました。
APIを更新する
チェックアウトを修正する
分析を統合する
キャッシュの実装
これらのプロンプトは実行可能に見えますが、ターゲットの動作、制約、および受け入れ基準が省略されています。それらに基づいて行動すると、コストのかかるループが簡単に引き起こされる可能性があります。つまり、モデルが推測し、ユーザーがそれを修正し、モデルがコンテキスト内のより多くの会話履歴を使用して作業を繰り返します。
このベンチマークでは、ドメイン ルーティングの問題も明らかになりました。次のようなプロンプト:
家をレンダリングする
ソフトウェア プロジェクトのフィードバックではなく、イメージ生成のフィードバックを受け取る必要があります。アナライザーは一般的なビジュアル レンダリング プロンプトを画像生成リクエストとして処理するようになり、ユーザーはファイル、コンポーネント、プラットフォーム スタックではなく、スタイル、構成、ライティング、出力形式についての質問を受けることができます。
現在のデフォルトのしきい値を使用すると、ベンチマークは以下を検出します。
98 / 100 曖昧なプロンプト
10 / 10 のイメージ生成プロンプト
現在の 2 つのミスは次のとおりです。
不安定なテストを修正する
さらにテストを生成する
これらのミスは、キャリブレーションに役立つケースです。彼らはショー

なぜベンチマークは単なる虚栄的な指標ではないのか。ベンチマークは、望ましい動作が明らかな場合に、保守者に議論し、調整し、回帰テストに変換するための具体的なプロンプトを提供します。
これは回帰保護であり、トークンの節約を保証するものではありません。このベンチマークはモデル トークンを消費せず、曖昧でコストのかかるプロンプトを見逃してしまうような変更を検出するのに役立ちます。
このリポジトリには、 .github/workflows/benchmark.yml にある GitHub Actions ワークフローも含まれています。プッシュ、プル リクエスト、および手動ワークフロー ディスパッチで単体テストと 100 プロンプト ベンチマークを実行します。
python3 スクリプト/install_codex_plugin.py
インストーラーはプラグインを ~/plugins/prompt-preflight にコピーし、 ~/.agents/plugins/marketplace.json でパーソナル マーケットプレイスを作成または更新し、 codex plugin add prompt-preflight@personal を実行しようとします。
Codex CLI がシェル PATH 上にない場合でも、インストーラーはファイルとマーケットプレイスのセットアップを完了し、インストールを完了するために必要なコマンドと Codex アプリのリンクを出力します。
ファイルを書き込まずに変更をプレビューします。
python3 スクリプト/install_codex_plugin.py --dry-run
以下については、外部セットアップ ガイドを参照してください。
macOS、Linux、および Windows のインストール
個人のマーケットプレイスの構成
インストーラー オプションと手動フォールバック手順
インストール後、Codex を再起動し、新しいスレッドを開き、 /hooks を使用してフックを確認します。
Codex が実行されるプロジェクトに .prompt-preflight.json を作成します。
{
"有効" : true 、
"モード" : "ブロック" ,
「しきい値」 : 45 、
"最大質問数" : 3
}
block : モデル作業の前に曖昧な送信を停止します。
nudge : Codex に最初に明確にするよう指示しながらターンを許可します。
しきい値 : 割り込みの頻度を減らすには、値を上げます。
max_questions : 説明の質問を 1 ～ 5 に制限します。
Enabled : プロジェクトのプロンプト プリフライトを無効にします。
設定を変更せずに 1 つのリクエストをバイパスします。
車を作る

画像 [プリフライト:スキップ]
プライバシーとセキュリティ
プロンプト テキストはローカルで分析されます。プロンプト プリフライトでは次のことは行われません。
安価なモデルを呼び出して、高価なモデルを実行するかどうかを決定する
プロンプト分析中にファイルを変更する
他のローカル プラグインと同様に、フックを信頼する前に、 .codex-plugin/plugin.json 、hooks/hooks.json、および scripts/prompt_preflight_hook.py を確認してください。
ルールベースのインテント ルーティングでは、すべての語句を理解することはできません。
現在、ドメイン カバレッジは意図的に狭く、高精度になっています。
ユーザーがモデルに仮定を置くことを好む場合、明確化によって摩擦が生じる可能性があります。
トークンの節約はタスクに依存しており、まだ自動的に測定されていません。
中断する価値がない場合、プロンプトは [preflight:skip] を使用することがあります。
間違った分類は回帰テストにする必要があります。 --json を使用して疑わしいプロンプトを実行し、検出された意図、理由、および質問をキャプチャします。
python3 -m 単体テスト 検出 -s テスト -v
Codex フック コントラクトをスモークテストします。
python3 scripts/prompt_preflight_hook.py << ' EOF '
{"hook_event_name":"UserPromptSubmit","prompt":"車の画像を作成する"}
終了後
このプロジェクトには現在、曖昧かつ詳細なプロンプト、ドメイン ルーティング、バイパス動作、ナッジ モード、および不正な形式のフック入力に対する回帰が含まれています。
トークンと再試行の節約テレメトリー

[切り捨てられた]

## Original Extract

Local Codex plugin that catches vague AI-agent prompts before they burn tokens. - akg268/prompt-preflight

GitHub - akg268/prompt-preflight: Local Codex plugin that catches vague AI-agent prompts before they burn tokens. · GitHub
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
akg268
/
prompt-preflight
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit .codex-plugin .codex-plugin .github/ workflows .github/ workflows docs docs hooks hooks scripts scripts skills/ prompt-preflight skills/ prompt-preflight src/ prompt_preflight src/ prompt_preflight tests tests .gitignore .gitignore .prompt-preflight.example.json .prompt-preflight.example.json LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml View all files Repository files navigation
Catch underspecified requests before they become expensive model turns.
Prompt Preflight is a local Codex plugin and standalone CLI that checks whether a prompt is specific enough to act on. When ambiguity and the cost of being wrong are both high, it pauses the request and gives the user:
A domain-aware example of a stronger prompt.
Up to three questions that fill the most important gaps.
The check uses deterministic Python rules. It makes no network requests and calls no model.
Prompt Preflight catches a vague Codex prompt before model work begins:
User submits a vague request
→ Prompt Preflight runs locally
→ Codex gets blocked before spending a model turn
→ the user receives a stronger prompt template and targeted questions
Why this exists
A vague prompt often creates an expensive loop:
Vague request
→ model reads project context
→ model produces the wrong interpretation
→ user corrects it
→ model reads the expanded conversation
→ model does the work again
The wasted cost is not limited to the first answer. The retry also carries the earlier prompt, output, corrections, and additional context.
Prompt Preflight moves clarification before that loop:
Vague request
→ local preflight check
→ targeted clarification
→ one stronger request
→ useful model work
It does not reduce the price per token. It reduces avoidable model input, unwanted output, repeated tool work, and correction turns.
Without preflight, the approximate cost of a failed attempt and retry is:
failed input + failed output + correction context + replacement input + replacement output
With preflight, the local check consumes zero model tokens. The intended path becomes:
clarified input + useful output
The potential tokens avoided are therefore approximately:
failed input + failed output + correction context + duplicated work
Actual savings depend on prompt quality, model behavior, context size, and task complexity. Prompt Preflight does not currently claim a fixed savings percentage; measured token telemetry is future work.
The largest benefit is expected on tasks where a wrong interpretation is costly, such as repository-wide changes, migrations, deployments, architecture work, or iterative image generation.
Create a car image
Prompt Preflight responds before image generation begins:
Your prompt:
"Create a car image"
Try asking:
"Create a [photorealistic/illustrated/3D] image of a car with
[key colors, materials, and distinctive details], in [setting/background],
viewed from [camera angle/composition], with [lighting/mood],
in [aspect ratio]."
Fill in the brackets by answering:
1. What should the car look like?
2. What visual style and mood do you want?
3. What setting, composition, lighting, and aspect ratio should it use?
This prevents an arbitrary first image followed by several rounds of visual corrections.
Example output after the prompt is clarified:
Make the dashboard better
Prompt Preflight suggests:
Improve the dashboard in [specific page/component] so [observable outcome].
Keep [important behavior or design constraints] unchanged.
Verify with [tests or acceptance criteria].
The model receives a target, outcome, boundaries, and definition of done before it reads files or edits code.
Runs before a Codex model turn through UserPromptSubmit .
Uses no model, API key, network access, or external service.
Routes prompts by domain before selecting feedback.
Includes software and image-generation feedback profiles.
Shows a tailored rewrite instead of only saying “be more specific.”
Asks at most three high-value questions.
Lets clear prompts and conversational follow-ups pass through.
Supports a one-time [preflight:skip] bypass.
Supports configurable block and nudge modes.
Fails open if hook input is malformed.
Provides structured JSON for evaluation and debugging.
Prompt Preflight estimates three things:
Intent: What kind of work is being requested?
Ambiguity: Which domain-specific details are missing?
Impact: How expensive would a wrong interpretation be?
It interrupts only when the prompt is actionable and both ambiguity and impact cross the configured threshold. This prevents the plugin from interrogating users about simple questions, confirmations, or already-specific work.
Current domain profiles include:
Unsupported domains use a conservative fallback rather than receiving software-specific questions.
Requires Python 3.10 or later.
python3 scripts/prompt_preflight.py " Create a car image "
A prompt requiring clarification exits with status 2 . A prompt ready to send exits with status 0 :
python3 scripts/prompt_preflight.py \
" Create a photorealistic image of a red 1967 Ford Mustang on a wet Tokyo street at night, low camera angle, cinematic lighting, 16:9. "
Inspect the full analysis:
python3 scripts/prompt_preflight.py --json " Rewrite the whole project "
Structured output includes the detected intent , ambiguity score, impact score, reasons, questions, and suggested prompt.
Benchmark vague-prompt detection
Prompt Preflight includes a fixed benchmark of 100 intentionally vague prompts across software work, bug fixes, deployment, migration, optimization, and image generation.
python3 scripts/benchmark_vague_prompts.py
Save complete results as JSON:
python3 scripts/benchmark_vague_prompts.py \
--min-block-rate 0.90 \
--json-output benchmark-results.json
The benchmark reports:
Number of vague prompts blocked before model work
Missed prompts that should be reviewed
Average ambiguity, impact, and clarification scores
Results grouped by detected intent
What the first benchmark taught us
The first 100-prompt run exposed exactly the kind of regression risk this project is meant to catch. Early scoring was too lenient on short action prompts such as:
Update the API
Fix checkout
Integrate analytics
Implement caching
Those prompts look actionable, but they omit the target behavior, constraints, and acceptance criteria. Acting on them can easily trigger a costly loop: the model guesses, the user corrects it, and the model repeats the work with more conversation history in context.
The benchmark also exposed a domain-routing issue. A prompt like:
Render a house
should receive image-generation feedback, not software-project feedback. The analyzer now treats common visual render prompts as image-generation requests so the user gets questions about style, composition, lighting, and output format instead of files, components, or platform stack.
With the current default threshold, the benchmark catches:
98 / 100 vague prompts
10 / 10 image-generation prompts
The two current misses are:
Fix the flaky tests
Generate more tests
These misses are useful calibration cases. They show why the benchmark is not just a vanity metric: it gives maintainers concrete prompts to discuss, tune, and convert into regression tests when the desired behavior is clear.
This is a regression guard, not a token-savings guarantee. The benchmark consumes zero model tokens and helps catch changes that would let vague, costly prompts slip through.
The repository also includes a GitHub Actions workflow at .github/workflows/benchmark.yml . It runs the unit tests and the 100-prompt benchmark on pushes, pull requests, and manual workflow dispatch.
python3 scripts/install_codex_plugin.py
The installer copies the plugin to ~/plugins/prompt-preflight , creates or updates the personal marketplace at ~/.agents/plugins/marketplace.json , and attempts to run codex plugin add prompt-preflight@personal .
If the Codex CLI is not on your shell PATH , the installer still completes the file and marketplace setup, then prints the command and Codex app link needed to finish installation.
Preview the changes without writing files:
python3 scripts/install_codex_plugin.py --dry-run
See the external setup guide for:
macOS, Linux, and Windows installation
Personal marketplace configuration
Installer options and manual fallback steps
After installation, restart Codex, open a new thread, and review the hook with /hooks .
Create .prompt-preflight.json in the project where Codex runs:
{
"enabled" : true ,
"mode" : " block " ,
"threshold" : 45 ,
"max_questions" : 3
}
block : stop the vague submission before model work.
nudge : allow the turn while instructing Codex to clarify first.
threshold : raise it to interrupt less often.
max_questions : limit clarification questions from 1 to 5.
enabled : disable Prompt Preflight for a project.
Bypass one request without changing configuration:
Create a car image [preflight:skip]
Privacy and security
Prompt text is analyzed locally. Prompt Preflight does not:
Invoke a cheaper model to decide whether an expensive model should run
Modify files during prompt analysis
As with any local plugin, review .codex-plugin/plugin.json , hooks/hooks.json , and scripts/prompt_preflight_hook.py before trusting the hook.
Rule-based intent routing cannot understand every phrasing.
Domain coverage is intentionally narrow and high-precision today.
Clarification can add friction when the user prefers the model to make assumptions.
Token savings are task-dependent and are not yet measured automatically.
Prompts may use [preflight:skip] when interruption is not worthwhile.
Incorrect classifications should become regression tests. Run a questionable prompt with --json and capture its detected intent, reasons, and questions.
python3 -m unittest discover -s tests -v
Smoke-test the Codex hook contract:
python3 scripts/prompt_preflight_hook.py << ' EOF '
{"hook_event_name":"UserPromptSubmit","prompt":"Create a car image"}
EOF
The project currently has regression coverage for vague and detailed prompts, domain routing, bypass behavior, nudge mode, and malformed hook input.
Token and retry savings telemetr

[truncated]
