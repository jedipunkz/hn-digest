---
source: "https://github.com/navels/neal"
hn_url: "https://news.ycombinator.com/item?id=49168496"
title: "Show HN: Neal – Codex writes the code, Claude reviews it"
article_title: "GitHub - navels/neal: A plan-driven, multi-agent coding loop: separate planner, coder, and reviewer roles, each on the provider and model you choose, work together to implement your plan. · GitHub"
author: "navels"
captured_at: "2026-08-04T13:51:07Z"
capture_tool: "hn-digest"
hn_id: 49168496
score: 1
comments: 0
posted_at: "2026-08-04T13:14:27Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Neal – Codex writes the code, Claude reviews it

- HN: [49168496](https://news.ycombinator.com/item?id=49168496)
- Source: [github.com](https://github.com/navels/neal)
- Score: 1
- Comments: 0
- Posted: 2026-08-04T13:14:27Z

## Translation

タイトル: Show HN: Neal – Codex がコードを作成し、Claude がそれをレビューします
記事のタイトル: GitHub - navels/neal: プラン主導のマルチエージェント コーディング ループ: 選択したプロバイダーとモデルに応じて、別々のプランナー、コーダー、レビューアーの役割が連携してプランを実装します。 · GitHub
説明: プラン主導のマルチエージェント コーディング ループ: 選択したプロバイダーとモデルに応じて、別々のプランナー、コーダー、およびレビュー担当者の役割が連携してプランを実装します。 - へそ/ニール
HN テキスト: neal は、LLM コーディング エージェント (この場合は Codex GPT-5.4) を使用して大規模なフロントエンド コードベースの移行を自律的に実行しようとして私が作成した CLI です。このプロジェクト中に、neal の開発をガイドするいくつかのことが起こりました。 1. エージェントに「ブロックされない限り作業を続ける」ように指示しても、実際には長期間にわたって機能しません。 2. 大規模なプロジェクトは、より小さな作業単位に分割する必要があります。 3. コーディング エージェントは、コンテキストの破損を防ぐために、作業の各チャンクを新しいコンテキストで開始する必要があります。 4. コーディング エージェントは、途中で別のエージェントに敵対的レビューを実行させることで利益を得ることができます。私が最終的に完成したのは、プランナー、コーダー、およびレビュー担当者の役割を構成できるオーケストレーターです。最初は Codex と Claude の SDK を使用していましたが、後に OpenRouter と、モデルがオーケストレーションを処理できることを確認するための互換モードを追加しました。現在、互換性のある OpenRouter モデルは 44 種類あります。 - プランナー/レビュー担当者のループを通じて、提供された計画を実行します。これにより、作業が適切なサイズのチャンクに分割され、ニールが実行できるように計画文書がフォーマットされます。 - コーダー/読み取り専用レビューアーのループを通じて作業の各チャンクを実行します。両方のエージェントが満足したら、ニールは次のチャンクに進みます。すべてが完了したら、コーダー/レビュー担当者による最終パスが行われます。

r ループを使用して、実装が計画全体を満たしていることを確認します。 - 作業の各チャンクでコーダーのコンテキストをリセットして、コンテキストのドリフトを防ぎ、レビュー担当者のコンテキストが長時間実行されるのを防ぎます。移行は完了しました (549 件のコミット、3000 以上のファイルが追加/変更されました)。また、Codex が neal を使用して単独で解決できなかった 100 件のベンチマーク ケース (SWE-bench Pro から) を実行しました。その結果、Codex を両方の役割 (コーダーとレビュー担当者) で使用した場合に 8 件が解決され、レビュー担当者として Claude を使用した場合に 15 件が解決されたことが観察されました。私はもっ​​とエキサイティングなベンチマーク結果を期待していましたが、これは少なくともニールの価値を示し、特にレビューアーの役割に別のモデルを置きました。 neal は、npm install -g @navels/neal によってインストールされる npm パッケージです。
ニールセットアップ経由で設定されています
これは、Codex や Claude Code のサブスクリプション (CLI がローカルで認証されている限り) または OpenAI/Anthropic/OpenRouter API キー経由で動作します。 neal の機能とそれがどのようにして誕生したかについては、私のブログに詳しく書いています: https://navels.dev/blog/neal/ モデルの機能が急速に進歩しているため、このツールがそれほど長く役立つとは期待していませんが、プランナー/コーダーとして Claude Fable を、レビュアーとして Codex Sol を起用し、今でも毎日のドライバーとしてこのツールを使用しています。 neal を試してみた場合、特にうまく機能しないユースケースがある場合、または追加のモデルが利用可能であることが必要な場合は、フィードバックをお待ちしています。

記事本文:
GitHub - navels/neal: プラン主導のマルチエージェント コーディング ループ: 選択したプロバイダーとモデルに応じて、別々のプランナー、コーダー、およびレビュー担当者の役割が連携してプランを実装します。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
へそ
/

ニール
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
3 コミット 3 コミット .github/ workflows .github/ workflows docs docs サンプル サンプル スクリプト スクリプト src src test test .gitignore .gitignore .nvmrc .nvmrc AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md eslint.config.mjs eslint.config.mjs neal.yml neal.yml package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml renovate.json renovate.json tsconfig.base.json tsconfig.base.json tsconfig.json tsconfig.json tsconfig.test.json tsconfig.test.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ニール — アニールから: より良い状態に向けて、制御された調整を繰り返します。
安定した結果。
計画主導のマルチエージェントコーディングループ。選択したプロバイダーとモデルに応じて、個別の planner 、 coder 、 reviewer の役割が連携して計画を実装します。
neal は、リポジトリのローカル プランナー/コーダー/レビューアー ループです。プランを作成すると、Neal のプランナー/レビューアー ループによって、現在のリポジトリに基づいた実行形式、管理可能なスコープ、および高レベルの実装の詳細を備えた人間によるレビューが可能な実行プランに変換されます。次に、neal は、設定されたコーダーおよびレビューアーのロールを使用して各スコープを実行します。すべてのスコープが受け入れられた後、最終レビュー パスで、変更のセット全体が計画全体に対してチェックされます。実行アーティファクトは .neal/ に記録されるため、中断された作業を再開できます。
2 つの選択肢によりデザインが説明されます。
コーダーとレビューアーは異なるベンダーで実行されます。このレビューはまさに敵対的です。独自の盲点を持つ 2 番目のモデルが、1 つではなく最初のモデルをチェックします。

モデル自身の作品をグレーディングします。
各スコープは、新しいコンテキストからコーダーを開始します。長期計画は、コンテキストがいっぱいになるにつれて単一のエージェントのように変動することはありません。
ニールがあなたの役に立ったのであれば、MusiCares への寄付をご検討ください。それは
経済的、個人的な問題が発生したときにミュージシャンに頼れる場所を提供する非営利団体
あるいは医療危機。
単一のエージェントではなく役割。プランナー、コーダー、およびレビュー担当者は、独立した個別に構成可能な役割です。役割ごとにベンダーとモデルを混合します (コーデックス コード、クロード レビューなど)。
レビュアーは読み取り専用です。プロバイダーの登録時に宣言され、各アダプターで機械的に適用されます (Codex の OS サンドボックス、Claude の SDK ツール ホワイトリスト、OpenRouter モデルのニール所有のジェイル ツールセット)。 SECURITY.md および docs/providers.md を参照してください。
ニールコンパト。 3 つの役割すべてにわたって OpenAI 互換/OpenRouter モデルを認定し、日付付きの PASS/FAIL ホワイトリストを発行する組み込みハーネス。 docs/compatibility-models.md を参照してください。
衝突安全性と再開可能性。すべての実行の状態とアーティファクトは .neal/ の下に存在します。ニール再開は中断された実行を続行します。
制限された自律回復。 neal は、レビュー担当者/コーダーのデッドロックのクラスをそれ自体で解決し、本物のブロッカーをエスカレーションします。
スコープはそれ自体を分割できます。スコープが予想より大きいことが判明した場合、コーダーは不適切な実装を 1 つのコミットに強制するのではなく、スコープを派生サブプランに分割します。 docs/plan-format.md を参照してください。
各部分がどのように組み合わされるかについては、「アーキテクチャの概要」を参照してください。
内容: クイックスタート ·
ニールが存在する理由 ·
設置・
プロバイダーのセットアップ ·
コマンドツアー・
コマンド ·
終了コード ·
構成・
アーティファクト ·
平面形状・
安全上の注意事項
npm install -g @navels/neal
neal setup # コーダーとレビューアーの役割のプロバイダーを選択する
neal check # 設定とプロバイダーの準備を確認します
クロードまたはコーデックスのサブスクリプションをお持ちですか?円周率

から適格なモデル スラッグを確認します。
互換性ホワイトリストを作成し、OpenRouter を使用する
プロバイダー (「プロバイダーのセットアップ」を参照)。
実際の機能を平易な言語で説明し、ファイル (例: PLAN.md ) に書き込みます。
OAuth 2.0 を使用して、ログイン ページに「Google でサインイン」オプションを追加します。
既存の電子メール/パスワードによるログインと同じ方法で、結果のセッションを保存します。
次に、リポジトリのルートから neel を実行します。
neal run <PLAN.md へのパス>
注: 私は通常、計画ドキュメントをリポジトリの .gitignore-ed ディレクトリに保存しますが、ymmv に保存します。
ニールが実行する前に、洗練された計画を読みたいですか? 2 つのステップを実行します
代わりに個別に:
ニール・プラン < path-to-PLAN.md >
# 洗練された計画を確認/更新してから、次のようにします。
neal 実行 < PLAN.md へのパス >
ニールはなぜ存在するのか
neal は、大規模なフロントエンド アップグレード (Ember 3.28 から Ember 5) から生まれました。そこでは、エージェントが最初の指示から時間の経過とともにずれることになりました。そのため、作業を小さなチャンクに分割し、各チャンクの前にエージェント コンテキストをリセットしたいと思うようになりました。また、私自身をレビューする前に、クロードに Codex の作業をレビューしてもらい、両方のエージェントが結果に満足するまで所見と回答をコピー/ペーストしてもらうという手動ワークフローも組み込みたいと考えていました。ニールについては次のフローに落ち着きました。
どのような作業を行う必要があるか、そしてそれをどのように行うかを計画することから始めます。
ニールはこれをプランナー/レビューアー ループを通じて送信し、実行形式を与え、管理可能なスコープを定義し、実装アプローチを具体化します。
各スコープは、コーダー/レビューアーのループを通じて実行されます。コーダーは新しいコンテキストから開始し、レビューアーはスコープ間でそのコンテキストを維持します。
レビュー担当者が満足すると、ニールは前のスコープをコミットした状態で次のスコープに移動します。
すべてのスコープが完了したら、最後に変更セット全体がコーダー/レビューアーのループで実行されます。
スコープが大きすぎることが判明した場合、

コーダーはそれをサブプランに分割できます
コーダーが何かでブロックされている場合、レビューアー モデルに支援を求めることができます。
何らかの理由で neal が終了した場合、neal 再開は続行を試み、コーダーがブロックされていた場合は指示を求めます。
実行アーティファクトと状態は .neal/ に記録されます。
npm install -g @navels/neal
neal には Node.js >= 24.18.0 が必要で、構成されたプロバイダー CLI/SDK を駆動します
(OpenAI Codex、Anthropic Claude、または任意の OpenAI 互換/OpenRouter モデル)。
「プロバイダーのセットアップ」を参照してください。
corepack 有効化 && pnpm インストール && pnpm 開始 -- ヘルプ
コントリビューターのセットアップと開発については、CONTRIBUTING.md を参照してください。
リンク、検証コマンド、CI ゲート、正規ドキュメント参照。
neal は、プランナー、コーダー、レビューアーの役割に別個のプロバイダーを使用します。 neal セットアップは明示的なコーダーを書き込み、レビューアーのデフォルトは ~/.neal/config.yml になります。
一般的な構成では、コーディングに Codex を使用し、レビューに Claude を使用します。
エージェント：
コーダー:
プロバイダー：openai-codex
モデル：ヌル
査読者：
プロバイダー: anthropic-claude
モデル：ヌル
model: null により、プロバイダーはデフォルトのモデルを選択できます。プランナーも設定可能
ただし、それ以外の場合はコーダー構成から継承されます。
OpenRouter およびその他の OpenAI 互換 API
openai 互換を使用して、OpenRouter または別の API を通じてモデルを実行します。
OpenAI 互換のチャット補完を実装します。とは別のものです
openai-codex を使用し、Codex CLI またはログインを使用しません。
OpenRouter の場合は、API キーをエクスポートします。
エクスポート OPENROUTER_API_KEY=...
次に、エンドポイントを ~/.neal/config.yml に追加します。
プロバイダー:
openai_compatibility :
Base_url : https://openrouter.ai/api/v1
api_key_env : OPENROUTER_API_KEY
別の互換性のある API の場合は、URL とキー変数を変更します。参照
環境のみの構成の場合は docs/providers.md
ローカルエンドポイント。
ニールセットアップ
neal セットアップはローカル ランタイムと構成された OpenAI-compa を検出します

テーブル設定、
ただし、認証やプロンプトの送信は行いません。 openai 互換を選択してください
OpenRouter を通じて実行する各ロールを選択し、OpenRouter モデルを入力します。
deepseek/deepseek-v3.2 などのスラッグ。
すべてのロールに対して同じプロバイダーをスクリプト化するには:
neal setup --provider anthropic-claude --all-roles
既存の有効なプロバイダー設定は、確認しない限り上書きされません。
対話的に行うか、 --force を渡します。
次に、構成とライブプロバイダーのアクセスを確認します。
ニールチェック
neal check は、効果的なプランナー/コーダー/レビュー担当者の選択肢を出力し、プロンプトを表示します。
構成された各ロールに 1 つの小さなライブ リクエストを送信する前に。非インタラクティブな場合
シェルの場合、構成のみを検証し、ライブ検証がスキップされたことを報告します。
実際の作業で openai 互換モデルを使用する前に、neal compat を実行してください (「
docs/compat.md )。無料またはレートが大幅に制限されたモデル プール
通常、プログラマーとしては失敗し、レビュー担当者としては信頼できません。プロバイダーの詳細は次のとおりです
docs/providers.md 、および一般的なセットアップ/認証エラーは次のとおりです。
docs/troubleshooting.md 。
Writer コマンドには、事前に既存の HEAD コミットを含む Git リポジトリが必要です。
プロバイダーの実行。前にリポジトリの初期ベースラインを作成してコミットします。
ニールに仕事を計画、実行、実行、再開、または潰すよう依頼します。
を確認したい場合は、neal plan と neal execution を別々に使用してください。
実行前の正規化された計画ドキュメント:
ニール・プラン PLAN.md
ニールはPLAN.mdを実行します
ニールは PLAN.md --no-squash を実行します
1 つ以上の計画を調整して連続的に実行します。
ニールラン tmp/A.md tmp/B.md
neal run --no-squash tmp/A.md tmp/B.md
利用可能なオペレーターがいない状態でヘッドレス (CI、cron、またはベンチマーク ハーネス) を実行します。
演算子ブロックに答えます。
neal 実行 PLAN.md --unattended
neal run --unattended tmp/A.md tmp/B.md
プロバイダーを構成した後、neal を介して例を実行するには:
ニールセットアップ
ニールチェック
cd サンプル/問題トライアグ

e-js
ニールランPLAN.md
pnpmテスト
Examples/issue-triage-js/README.md を参照してください。
ガイド例と安全上の注意については、こちらをご覧ください。
ニール履歴書
ニール・レジューム --run < run-id >
ニール ステータスまたはニール履歴書で、選択した実行が待機中であることが示された場合にのみガイダンスを提供します。
neal raise --run < run-id > --message " 失敗したテストに限定された変更を保持し、必要な検証ゲートを再実行します。 "
ニールレビューは、すでにコミットされた作業のための別個の計画不要のワークフローです。の
コーダーが結果を提案し、レビュー担当者がそれを判断し、ニールが結果を拒否します。
レビュー中に行われたワークツリーの変更:
ニールレビュー --最後の 3 件
neal レビュー「認証/セッション処理に焦点を当てます。」 --since Origin/main
検査ステータス:
ニールステータス
ニールステータス --run < run-id >
neal status --json --run < run-id >
ニールステータス --すべて
ニールステータス --json --all
自動化契約
neal は、スクリプト、CI、およびベンチマーク ハーネスから駆動するための安定したマシン側の契約を維持します。これには、ライター終了コード、neal status --json 分類スキーマ、パッチ送信資格、ハーネス タイムアウトとトレース公開ガイダンスが含まれます。 docs/automation.md を参照してください。 neal-swebench は、こ​​の契約を通じてニールを SWE-bench Pro での役割ペアのベンチマークを行うように促します。
# セットアップ
ニールセットアップ
ニールチェック
ニール compat [--model < s

[切り捨てられた]

## Original Extract

A plan-driven, multi-agent coding loop: separate planner, coder, and reviewer roles, each on the provider and model you choose, work together to implement your plan. - navels/neal

neal is a CLI I wrote while trying to use an LLM coding agent (in this case, Codex GPT-5.4) to work autonomously on a migration of our large frontend codebase. A few things came up during that project that guided the development of neal: 1. Telling an agent to "keep working unless blocked" doesn't actually work over long stretches of time. 2. Large projects should be broken up into smaller chunks of work. 3. The coding agent should start each chunk of work with a fresh context to prevent context rot. 4. The coding agent benefits from having a different agent doing adversarial reviews along the way. What I ended up with is an orchestrator that - lets you configure planner, coder, and reviewer roles. I started with Codex and Claude using their SDKs but later added OpenRouter as well as a compatibility mode for verifying that a model can handle the orchestration. Currently there are 44 compatible OpenRouter models. - runs a plan that you provide through a planner / reviewer loop which splits the work into reasonable-sized chunks and formats the plan document so that it can be executed by neal. - runs each chunk of work through a coder / read-only reviewer loop. Once both agents are satisfied, neal moves on to the next chunk. Once everything is complete there is a final pass through the coder / reviewer loop to ensure the implementation satisfies the entire plan. - resets the coder's context with each chunk of work to prevent context drift, leaving the reviewer's context long-running. The migration landed (549 commits, over 3000 files added/modified). I also ran 100 benchmark cases (from SWE-bench Pro) that Codex failed to solve on its own through neal and observed 8 solved when using neal with Codex in both (coder and reviewer) roles and 15 solved when using Claude as the reviewer. I had hoped for a more exciting benchmark outcome but this at least demonstrated neal's value, especially putting a different model in the reviewer role. neal is an npm package installed via npm install -g @navels/neal
and configured via neal setup
It works with your Codex and/or Claude Code subscriptions (as long as those CLIs are authenticated locally) or via OpenAI/Anthropic/OpenRouter API keys. I have a more detailed writeup of neal's capabilities and how it came to be on my blog: https://navels.dev/blog/neal/ With how fast model capabilities are progressing I'm not expecting this tool to be very useful for too much longer, but I'm still using it as my daily driver, with Claude Fable as the planner/coder and Codex Sol as the reviewer. If you try neal I'd love to hear your feedback, especially if you have a use case that doesn't quite work or you want additional models available.

GitHub - navels/neal: A plan-driven, multi-agent coding loop: separate planner, coder, and reviewer roles, each on the provider and model you choose, work together to implement your plan. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
navels
/
neal
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits .github/ workflows .github/ workflows docs docs examples examples scripts scripts src src test test .gitignore .gitignore .nvmrc .nvmrc AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md eslint.config.mjs eslint.config.mjs neal.yml neal.yml package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml renovate.json renovate.json tsconfig.base.json tsconfig.base.json tsconfig.json tsconfig.json tsconfig.test.json tsconfig.test.json View all files Repository files navigation
neal — from anneal : repeated, controlled adjustment toward a more
stable result.
A plan-driven, multi-agent coding loop. Separate planner , coder , and reviewer roles, each on the provider and model you choose, work together to implement your plan.
neal is a local planner/coder/reviewer loop for your repo. You write a plan, and neal's planner/reviewer loop turns it into a human-reviewable execution plan with an execution shape, manageable scopes, and high-level implementation detail grounded in the current repository. neal then runs each scope with your configured coder and reviewer roles. After every scope is accepted, a final review pass checks the total set of changes against the entire plan. Run artifacts are recorded under .neal/ so interrupted work can resume.
Two choices explain the design:
The coder and reviewer run on different vendors. The review is genuinely adversarial: a second model with its own blind spots checks the first, instead of one model grading its own work.
Each scope starts the coder from a fresh context. A long plan doesn't drift the way a single agent does as its context fills up.
If neal has been useful to you, please consider donating to MusiCares. It's a
non-profit giving musicians a place to turn to in times of financial, personal,
or medical crisis.
Roles, not one monolithic agent. Planner, coder, and reviewer are independent, separately configurable roles. Mix vendors and models per role (e.g. Codex codes, Claude reviews).
The reviewer is read-only. Declared at provider registration and enforced mechanically in each adapter (OS sandbox for Codex, SDK tool allowlist for Claude, a neal-owned jailed toolset for OpenRouter models). See SECURITY.md and docs/providers.md .
neal compat . A built-in harness that qualifies any OpenAI-compatible / OpenRouter model across all three roles and emits a dated PASS/FAIL whitelist. See docs/compatible-models.md .
Crash-safe and resumable. Every run's state and artifacts live under .neal/ . neal resume continues an interrupted run.
Bounded autonomous recovery. neal resolves a class of reviewer/coder deadlocks itself, and escalates genuine blockers.
Scopes can split themselves. When a scope turns out bigger than expected, the coder splits it into a derived sub-plan instead of forcing a bad implementation into one commit. See docs/plan-format.md .
For how the pieces fit together, see the architecture overview .
Contents: Quickstart ·
Why neal exists ·
Installation ·
Provider setup ·
Command tour ·
Commands ·
Exit codes ·
Configuration ·
Artifacts ·
Plan shape ·
Safety notes
npm install -g @navels/neal
neal setup # pick providers for the coder and reviewer roles
neal check # verify config and provider readiness
No Claude or Codex subscription? Pick a qualified model slug from the
compatibility whitelist and use the OpenRouter
provider (see Provider setup ).
Describe a real feature in plain language and write it to a file (e.g., PLAN.md ):
Add a "Sign in with Google" option to the login page using OAuth 2.0.
Store the resulting session the same way the existing email/password login does.
Then run neal from your repository root:
neal run < path-to-PLAN.md >
Note: I usually keep plan docs in a .gitignore-ed directory in my repo but ymmv.
Want to read the refined plan before neal executes it? Run the two steps
separately instead:
neal plan < path-to-PLAN.md >
# review/update the refined plan, then:
neal execute < path-to-PLAN.md >
Why neal exists
neal grew out of a large frontend upgrade (Ember 3.28 to Ember 5) where the agent would drift over time from its initial instructions, which led to me wanting to break up the work into smaller chunks and reset the agent context before each chunk. I also wanted to incorporate my manual workflow of having Claude review Codex's work, copy/pasting findings and responses until both agents were satisfied with the result, before reviewing myself. I settled on this flow for neal:
start with a plan of what work needs to be done and how to do it
neal sends this through the planner/reviewer loop to give it an execution shape, define manageable scopes, and flesh out the implementation approach
each scope is run through the coder/reviewer loop with the coder starting with a fresh context and the reviewer keeping its context from scope to scope
when the reviewer is satisfied, neal moves on to the next scope with the previous scope committed
after all scopes are complete, the entire set of changes is run through the coder/reviewer loop a final time
if a scope is found to be too large, the coder can split it into a sub-plan
if the coder is blocked on something, it can consult the reviewer model for assistance
if neal exits for any reason, neal resume will attempt to continue, prompting for direction if the coder was blocked
run artifacts and state are recorded in .neal/
npm install -g @navels/neal
neal requires Node.js >= 24.18.0 and drives your configured provider CLIs/SDKs
(OpenAI Codex, Anthropic Claude, or any OpenAI-compatible / OpenRouter model).
See Provider setup .
corepack enable && pnpm install && pnpm start -- help
See CONTRIBUTING.md for contributor setup, the development
link, verification commands, CI gates, and canonical doc references.
neal uses separate providers for the planner, coder, and reviewer roles. neal setup writes explicit coder and reviewer defaults to ~/.neal/config.yml .
A typical config uses Codex for coding and Claude for review:
agent :
coder :
provider : openai-codex
model : null
reviewer :
provider : anthropic-claude
model : null
model: null lets the provider choose its default model. The planner can also be configured
but otherwise inherits from the coder config.
OpenRouter and other OpenAI-compatible APIs
Use openai-compatible to run models through OpenRouter or another API that
implements OpenAI-compatible Chat Completions. It's separate from
openai-codex and doesn't use the Codex CLI or login.
For OpenRouter, export your API key:
export OPENROUTER_API_KEY=...
Then add the endpoint to ~/.neal/config.yml :
providers :
openai_compatible :
base_url : https://openrouter.ai/api/v1
api_key_env : OPENROUTER_API_KEY
For another compatible API, change the URL and key variable. See
docs/providers.md for environment-only configuration and
local endpoints.
neal setup
neal setup detects local runtimes and configured OpenAI-compatible settings,
but it doesn't authenticate or send prompts. Choose openai-compatible for
each role you want to run through OpenRouter, then enter an OpenRouter model
slug such as deepseek/deepseek-v3.2 .
To script the same provider for every role:
neal setup --provider anthropic-claude --all-roles
Existing effective provider settings are not overwritten unless you confirm
interactively or pass --force .
Then verify the config and live provider access:
neal check
neal check prints the effective planner/coder/reviewer choices, then prompts
before sending one small live request to each configured role. In non-interactive
shells, it only validates config and reports that live verification was skipped.
Before using an openai-compatible model for real work, run neal compat (see
docs/compat.md ). Free or heavily rate-limited model pools
usually fail as coders and are unreliable reviewers. Provider details are in
docs/providers.md , and common setup/auth failures are in
docs/troubleshooting.md .
Writer commands require a Git repository with an existing HEAD commit before
provider execution. Create and commit the repository's initial baseline before
asking neal to plan, run, execute, resume, or squash work.
Use neal plan and neal execute separately when you want to review the
normalized plan document before execution:
neal plan PLAN.md
neal execute PLAN.md
neal execute PLAN.md --no-squash
Refine and execute one or more plans serially:
neal run tmp/A.md tmp/B.md
neal run --no-squash tmp/A.md tmp/B.md
Run headless (CI, cron, or a benchmark harness) with no operator available to
answer an operator block:
neal execute PLAN.md --unattended
neal run --unattended tmp/A.md tmp/B.md
To run an example through neal after configuring providers:
neal setup
neal check
cd examples/issue-triage-js
neal run PLAN.md
pnpm test
See examples/issue-triage-js/README.md
for the example guide and safety notes.
neal resume
neal resume --run < run-id >
Give guidance only when neal status or neal resume says the selected run is waiting for it:
neal resume --run < run-id > --message " Keep the change bounded to the failing test and rerun the required validation gate. "
neal review is a separate, plan-free workflow for already-committed work. The
coder proposes findings, the reviewer judges them, and neal rejects any
worktree changes made during the review:
neal review --last 3
neal review " Focus on auth/session handling. " --since origin/main
Inspect status:
neal status
neal status --run < run-id >
neal status --json --run < run-id >
neal status --all
neal status --json --all
Automation contract
neal keeps a stable machine-facing contract for driving it from scripts, CI, and benchmark harnesses: writer exit codes, the neal status --json classification schema, patch-submission eligibility, and harness timeout and trace-publishing guidance. See docs/automation.md . neal-swebench drives neal through this contract to benchmark role pairings on SWE-bench Pro.
# Setup
neal setup
neal check
neal compat [--model < s

[truncated]
