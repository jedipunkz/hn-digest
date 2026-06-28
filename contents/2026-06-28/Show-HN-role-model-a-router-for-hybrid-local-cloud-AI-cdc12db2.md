---
source: "https://github.com/try-works/role-model"
hn_url: "https://news.ycombinator.com/item?id=48706181"
title: "Show HN: role-model, a router for hybrid local/cloud AI"
article_title: "GitHub - try-works/role-model: role-model is a protocol for assigning the right model for the right job. Use local and cloud AI together, or route between several cloud providers. · GitHub"
author: "try-working"
captured_at: "2026-06-28T11:33:46Z"
capture_tool: "hn-digest"
hn_id: 48706181
score: 1
comments: 1
posted_at: "2026-06-28T10:46:08Z"
tags:
  - hacker-news
  - translated
---

# Show HN: role-model, a router for hybrid local/cloud AI

- HN: [48706181](https://news.ycombinator.com/item?id=48706181)
- Source: [github.com](https://github.com/try-works/role-model)
- Score: 1
- Comments: 1
- Posted: 2026-06-28T10:46:08Z

## Translation

タイトル: Show HN: ロールモデル、ハイブリッド ローカル/クラウド AI 用ルーター
記事タイトル: GitHub - try-works/role-model: role-model は、適切なジョブに適切なモデルを割り当てるためのプロトコルです。ローカル AI とクラウド AI を一緒に使用するか、複数のクラウド プロバイダー間でルーティングします。 · GitHub
説明: role-model は、適切なジョブに適切なモデルを割り当てるためのプロトコルです。ローカル AI とクラウド AI を一緒に使用するか、複数のクラウド プロバイダー間でルーティングします。 - トライワークス/ロールモデル
HN テキスト: 皆さん、今日はロールモデルを発表します。ルーティング プロトコル、リファレンス ルーター ランタイム、およびより適切な情報に基づいたルーティングの決定を可能にする Pi の拡張機能です。 role-model は、選択されたルーティング戦略に基づいてリクエストをルーティングするコントローラー モデルへのフォールバックを備え、ほとんどが決定的です。このプロトコルは、ドメインとロールをモデルに割り当てることを中心に構造化されており、Pi のようなコンシューマー アプリケーションによって送信されるリクエストには、ルーティング メタデータを強化して精度を高めるためのタスク タイプが含まれています。組み込みのベンチマークを実行して、速度、品質、コスト全体のモデルのパフォーマンスと、実際のタスクで観察されたパフォーマンスを比較できます。 [0] でルーティングがどのように機能するかを示す図があります。ランタイムは、ローカル エンドポイント (LM Studio、llama.cpp など) に直接接続するか、ベンダーの llama-swap を介して複数のローカル モデル間のルーティングを行うローカル モデルをサポートします。昨日、ルーティングの基本について議論する別のモデル ルーターの投稿があったため、これを構築してテストすることで得た興味深い学習のいくつかについて説明することに焦点を当てます。 1. モデル ルーティングは基本的に将来を予測しようとしています。どのモデルがこのリクエストに対して (ユーザーが定義した基準に基づいて) 最適に実行するか? 2. リクエストをルーティングした後、それが正しい決定だったのか、あるいは他のモデルが実行したであろう判断だったかどうかを評価したいと考えます。

3. また、ルーターに (とりわけ) 困難さを評価させて単独で意思決定をさせるのは、理想からはほど遠いこともわかります。コンシューマ アプリケーションがルーターと連携して、リクエストに必要なものを定義することを希望します。 4. また、モデル間の区別がより明確になると、それがはるかに簡単になり、決定がはるかに正確になり、ルーティングの結果がより影響力を持つようになることもわかります。 ポイント 2 については、同じリクエストでプール内のモデルをベンチマークするためにローカルで実行できる評価を開始します。ここでの結果は、新しいリクエストをルーティングする際の入力としてポイント 1 に使用できます。ポイント 3 については、Pi 用の pi-role-model パッケージを構築しました。これにより、Pi エージェントは、難易度、優先ロール、または特定のモデル ID、必要な機能 (ツールの使用や画像入力など) などを含む role_model.intent メタデータを挿入できます。 Pi でこれをさらにカスタマイズし、メタデータを変更することで追加の方法でルーティングできるはずです。これが、私がロールモデルのルーティング プロトコルも構築した理由です。ポイント 4 については、モデル ルーティングが二次効果として実際に行うことは、特殊なモデルの市場を生み出すことです。モデルは、より小さい場合もそうでない場合もあり、より安価な場合もより高価な場合もあり、ローカルで実行可能である場合もあります。 2 つのフロンティア モデル (GPT 5.5 と Opus 4.8) の間でルーティングすることはほとんど意味がありません。品質、速度、コストの要素の 1 つが他の候補モデルの倍数であるモデル間をルーティングすることはより合理的であり、コード、散文、数学と科学、ビジュアルなどの特殊なドメイン モデルを用意することはさらに合理的です。この段階で、モデル ルーティングが本当に価値を持つようになります。ロールモデルにはリファレンス ランタイムがあり、それを継続的に構築しています (ルーティングを改善し、ユーザーがより詳細に制御できるようにするためにやるべきことはたくさんあります)

ロールモデルの最終的な目標は、コンシューマ アプリケーションで使用される推論リクエストの標準プロトコルが存在することであり、ルーター ミドルウェアであれ推論プロバイダであれ、プロバイダーがコスト、速度、品質の最適なバランスをとり、ユーザーの選択も尊重するモデルにルーティングできるようになり、一部のタスクではローカル モデルを使用し、他のタスクではクラウドを許可するようにユーザーがこれらの設定を制御できるようになります。リンク: [0] role-model - モデル ルーティング プロトコルのケース: https://try.works/role-model-the-case-for-a-model-routing-pr... [1] GitHub: https://github.com/try-works/role-model [2] ドキュメント: https://role-model.dev/

記事本文:
GitHub - try-works/role-model: role-model は、適切なジョブに適切なモデルを割り当てるためのプロトコルです。ローカル AI とクラウド AI を一緒に使用するか、複数のクラウド プロバイダー間でルーティングします。 · GitHub
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
トライワークス
/
ロールモデル
公共
通知
君はね

サインインして通知設定を変更できない
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
269 コミット 269 コミット .agent .agent .agents/ skill .agents/ skill .codex .codex .github .github .recursive .recursive .tmp .tmp apps/ docs-site apps/ docs-site cla cla docs docs 証拠/ スクリプト 証拠/ スクリプト パッケージ パッケージ プロトコル プロトコル ロールモデル-ルーターrole-model-router スキーマ/ role-model/ 分類スキーマ/ role-model/ 分類 スクリプト スクリプト署名 署名 testdata testdata .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLA.md CLA.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンスREADME.md README.md biome.json biome.json docker-compose.yml docker-compose.yml package.json package.json pi-role-model-extension-proposal-v3.md pi-role-model-extension-proposal-v3.md pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml Rust-toolchain.toml Rust-toolchain.toml skill-lock.json skill-lock.json tsconfig.base.json tsconfig.base.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
role-model は、機能を認識した AI ルーティング用のオープン プロトコルに加え、以下を実装するリファレンス ルーターです。
そのプロトコル。
これは、リクエストが何を必要とするか、エンドポイントが何を実行できるかを記述するための永続的な契約をルーターに与えます。
どのようなポリシーで許可されているか、および最終的なルーティング決定が行われた理由。
すべての AI ワークロードは、最終的に同じ質問に直面します。どのモデルがこのリクエストを処理すべきか?答えは
タスクのタイプ、必要な機能、コスト、レイテンシ、モデルがローカルで実行されているか、ローカルで実行されているかによって異なります。
雲。ロールモデルは、その決定を明示的、説明可能、移植可能にします。
大まかに言うと、ロールモデルは AI ルーティングをいくつかの安定した部分に分割します。
リクエストはタスクの種類を記述します。

d 機能、モダリティ、ツールのニーズ、および制約。
エンドポイント ID とプロファイルは、抽象的なモデル名ではなく、具体的なルーティング可能なエンドポイントを記述します。
ルーティング ポリシーは、ハード拒否、プリファレンス、予算、およびタイブレーク ルールを適用します。
可観測性アーティファクトは、決定、トレース、使用状況、および観察されたパフォーマンスを記録します。
リファレンス ルーターは、次の 3 つの展開形態にわたるハイブリッド ルーティングをサポートします。
ローカル / ローカル - 役割、タスク、機能に基づいたローカル モデル (ラマ スワップ ピアなど) 間のルート
ローカル/クラウド - コスト、遅延、タスクの難易度に基づいてローカル モデルとクラウド プロバイダー間のルートを決定します。
クラウド / クラウド - 機能と経済性に基づいたクラウド プロバイダー (OpenAI、DeepSeek、Moonshot など) 間のルート
これは、単一のランタイムで高速ローカル モデルからの迅速なチャット リクエストを処理し、複雑なコーディングをルーティングできることを意味します。
タスクを有能なクラウド モデルに割り当て、プライマリの機能が低下した場合は、より安価なクラウド エンドポイントにフォールバックします。
説明可能な 1 つのルーティング契約に基づいて。
エンド ユーザーの場合は、ソース ビルドよりもパッケージ化されたスタンドアロン ランタイムを好みます。
カール -fsSL https://raw.githubusercontent.com/try-works/role-model/main/scripts/install.sh |しー
インストーラーは最新の GitHub リリース アーカイブをダウンロードし、次の場所にインストールします。
~/.local/share/role-model-router/<version>/<target>/ で、role-model-router ランチャーを公開します。
~/.local/bin 。
irm https://raw.githubusercontent.com/try-works/role-model/main/scripts/install.ps1 |アイエックス
インストーラーは最新の GitHub リリース アーカイブをダウンロードし、次の場所にインストールします。
%LOCALAPPDATA%\Programs\RoleModelRouter\<version>\<target>\ を実行し、role-model-router.cmd を作成します。
ランチャー。
インストーラー スクリプトを使用したくない場合は、GitHub Releases から一致するアーカイブをダウンロードしてください。
pi-role-model パッケージは、Pi を外部で実行されているロールに接続します。

-モデルのランタイム。
まず、Role-Model ランタイムを起動してから、パブリック Pi パッケージをインストールします。
pi インストール npm:@try-works/pi-role-model
このリポジトリからローカル チェックアウト テストを行うには、パッケージを直接インストールします。
pi インストール ./packages/pi-role-model
Pi 内で次を実行します。
/role-model セットアップ
/ロールモデルの医師
/role-model エイリアス リスト
/role-model エイリアス選択
/role-model エイリアスは <エイリアス> を使用します
デフォルトでは、パッケージは http://127.0.0.1:3456 に接続し、Role-Model を
/api/role-model/downstream/openai を使用するロールモデルプロバイダー。 ROLE_MODEL_ENDPOINT を設定する
Pi を起動して別のローカル ランタイムを使用する前に。リモートエンドポイントには明示的な必要があります
信頼できるallowRemote動作、およびauthentication.requiredが失敗することを報告するランタイム
将来サポートされるトークン ソースが設定されていない限り、閉じられます。ローカル開発インストールの場合
および完全なコマンド リファレンスについては、を参照してください。
パッケージ/pi-role-model/README.md 。
Node.js 24 (node:sqlite および SEA のサポートに必要)
pnpm 10.x (corepack enable 経由)
Go 1.24+ (llama-swap ベンダー バイナリおよび Windows ランチャー用)
コアパックを有効にする
コアパック pnpm インストール
発煙試験
コアパック pnpm 実行煙
詳しいチュートリアルについては、 docs/public/quickstart.md を参照してください。
ブリッジと UI を開発モードで実行します (別のプロセス)。
# ターミナル 1: ブリッジサーバー
cd role-model-router/apps/runtime-host-bridge
corepack pnpm exec tsx scripts/start-for-qa.ts
# ターミナル 2: UI 開発サーバー
cd role-model-router/apps/runtime-ui
corepack pnpm exec 反応ルーター dev --port 5173 --host 127.0.0.1
次に、ブラウザで http://127.0.0.1:5173 を開きます。
実稼働ビルド (すべてのプラットフォーム)
UI を構築し、SEA ランタイムをパッケージ化します。
# UI静的ファイルをビルドする
corepack pnpm --filter @role-model-router/runtime-ui run build
# ブリッジを単一の実行可能ファイルとしてパッケージ化する
corepack pnpm run runtime:package-sea
出力: role-model-router/dist

/release/<プラットフォーム-アーキテクチャ>/role-model-runtime
専用のブラウザ ウィンドウを使用して完全な Windows パッケージを構築します。
#1. UIを構築する
corepack pnpm --filter @role-model-router/runtime-ui run build
# 2. パッケージブリッジ SEA ランタイム
corepack pnpm run runtime:package-sea
# 3. Go ランチャーを構築する
cd ロールモデルルーター/apps/launcher
go build -o ../../dist/release/win32-x64/role-model-launcher.exe main.go
# 4. UI ファイルをバンドルする
cp -r ../runtime-ui/build/client ../../dist/release/win32-x64/
次に、 dist/release/win32-x64/ にある role-model-launcher.exe をダブルクリックします。それは次のことを行います:
ポート 3456 でブリッジ サーバーを起動します。
Microsoft Edgeをアプリモードで開きます（専用ウィンドウ）
ブリッジから直接 UI を提供します (別個の開発サーバーなし)
これを読んでください
ご希望であれば
docs/public/README.md
ドキュメントハブ
docs/public/introduction.md
ロールモデルとは何か、そしてなぜそれが存在するのか
docs/public/quickstart.md
本当のエンドツーエンドのスモークラン
docs/public/concepts/how-role-model-works.md
システムの流れ
docs/public/concepts/protocol-overview.md
プロトコル表面
docs/public/concepts/routing-overview.md
ルーティングの決定はどのように行われるか
プロトコル/README.md
正規スキーマとフィクスチャ
ロールモデルルーター/README.md
リファレンスルーターパッケージとランタイムアプリ
docs/protocol/routing-policy.md
ルーティングポリシーリファレンス
docs/protocol/taxonomy-v1.md
分類 V1 のグループ、役割、タスク、および Pi 分類
docs/protocol/roles.md
ロールメタデータリファレンス
docs/プロトコル/tasks.md
タスクメタデータリファレンス
docs/operations/02-ci-and-release-flow.md
CI、リリースの自動化、およびワークフローの所有権
変更ログ.md
リリース履歴
謝辞
ロールモデルは、いくつかのオープンソース プロジェクトの成果に基づいて構築されています。
llama-swap - ローカル エンドポイントのプロセス監視、リクエスト転送、モデル スワップを処理するベンダーのローカル モデル ライフサイクル マネージャー
LiteLLM - 統合 LLM

プロバイダー カタログ、モデル メタデータ、および価格設定データがルーティング互換プロバイダー インベントリに情報を提供する API 抽象化
このリポジトリは、プロジェクト固有の BUSL-1.1 に基づいてライセンスされています。
追加使用許可。社内生産利用、評価、開発、
修正および非実稼働再配布はルートの下で許可されます
ライセンス。ホストまたは管理されたサードパーティ サービス、有料製品の埋め込み、および
サードパーティによる商用化には、別途商用ライセンスが必要です。
全条件については、「ライセンス」を参照してください。貢献には同意が必要です
結合する前に、コントリビューター ライセンス契約を確認してください。
個人の寄付のみ受け付けます。が所有する作品を投稿しないでください。
あなたが個人的に所有している場合を除き、雇用主、顧客、会社、またはその他の団体
このプロジェクトの規約に基づいて投稿する権利。
role-model は、適切なジョブに適切なモデルを割り当てるためのプロトコルです。ローカル AI とクラウド AI を一緒に使用するか、複数のクラウド プロバイダー間でルーティングします。
読み込み中にエラーが発生しました。このページをリロードしてください。
1
フォーク
レポートリポジトリ
リリース
2
v0.0.1-alpha.3
最新の
2026 年 6 月 27 日
+ 1 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

role-model is a protocol for assigning the right model for the right job. Use local and cloud AI together, or route between several cloud providers. - try-works/role-model

Hey everyone, I'm launching role-model today: a routing protocol, a reference router runtime, and an extension for Pi that allows for better informed routing decisions. role-model is mostly deterministic, with fallback to a controller model, that routes requests based on a chosen routing strategy. the protocol is structured around assigning domains and roles to models, where requests sent by consumer applications like Pi have task types to enrich routing metadata and thereby accuracy. you can to run the built-in benchmark to compare performance of models across speed, quality and cost, as well as observed performance on real tasks. I have a diagram on how routing works in [0]. The runtime supports local models, either directly to your local endpoint (LM Studio, llama.cpp etc), or routing between multiple local models via vendored llama-swap. Since there was another model router post yesterday where people discussed the basics of routing, I will focus on discussing some of the interesting learnings I've made building and testing this: 1. Model routing is essentially trying to predict the future: which model will perform optimally (based on criteria defined by the user) on this request? 2. After you have routed the request, you want to evaluate if it was the right decision or if some other model would have performed better 3. You also realize that having the router assess difficulty (among other things) to make decisions by itself is far from ideal - we'd prefer to have the consumer application work with the router to define what the request needs 4. You also realize that it becomes much easier, decisions become much accurate, and the outcomes of routing becomes more impactful when there is more of a distinction between models For point 2, I will be launching evals that you can run locally to benchmark models in your pool on the same requests. The outcomes here can then be used for point 1, as input when routing new requests. For point 3, I've built the pi-role-model package for Pi, which lets the Pi agent inject role_model.intent metadata including difficulty, preferred roles or even specfic model ids, required capabilities (say tool use or image input) and so on. You should be able to customized this further in Pi, and route in additional ways by changing metadata. This is why I've also built the role-model routing protocol. For point 4, what model routing really does as a second order effect is create a market for specialized models - models that may or may not be smaller, could be cheaper or more expensive, may be locally runnable. It makes little sense to route between two frontier models (GPT 5.5 and Opus 4.8); it makes more sense to route between models where one of the factors of quality, speed, cost is a multiple of the other candidate models, and it makes even more sense to have specialized domain models: code, prose, math and science, visuals and so on. It is at this stage model routing becomes really valuable. While role-model has a reference runtime that I'm continuously building out (there's lots to do to improve routing, as well as give users more granular control over routing decisions, and also ways to improve cross-model caching and also add techniques like FastContext), the ultimate goal of role-model is for there to be a standard protocol for inference requests that is used by consumer applications, so that the provider, be it a router middleware or an inference provider, will be able to route to a model that strikes the best balance between cost, speed and quality and also respects user choices, and even lets the user control these preferences to use local models for some tasks and allow cloud for others. Links: [0] role-model - the case for a model routing protocol: https://try.works/role-model-the-case-for-a-model-routing-pr... [1] GitHub: https://github.com/try-works/role-model [2] Docs: https://role-model.dev/

GitHub - try-works/role-model: role-model is a protocol for assigning the right model for the right job. Use local and cloud AI together, or route between several cloud providers. · GitHub
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
try-works
/
role-model
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
269 Commits 269 Commits .agent .agent .agents/ skills .agents/ skills .codex .codex .github .github .recursive .recursive .tmp .tmp apps/ docs-site apps/ docs-site cla cla docs docs evidence/ scripts evidence/ scripts packages packages protocol protocol role-model-router role-model-router schemas/ role-model/ taxonomy schemas/ role-model/ taxonomy scripts scripts signatures signatures testdata testdata .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLA.md CLA.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md biome.json biome.json docker-compose.yml docker-compose.yml package.json package.json pi-role-model-extension-proposal-v3.md pi-role-model-extension-proposal-v3.md pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml rust-toolchain.toml rust-toolchain.toml skills-lock.json skills-lock.json tsconfig.base.json tsconfig.base.json View all files Repository files navigation
role-model is an open protocol for capability-aware AI routing, plus a reference router that implements
that protocol.
It gives a router a durable contract for describing what a request needs , what an endpoint can do ,
what policy allows , and why a final routing decision was made .
Every AI workload eventually faces the same question: which model should handle this request? The answer
depends on task type, required capabilities, cost, latency, and whether the model is running locally or in
the cloud. role-model makes that decision explicit, explainable, and portable.
At a high level, role-model separates AI routing into a few stable pieces:
Requests describe task type, required capabilities, modalities, tool needs, and constraints.
Endpoint identities and profiles describe concrete routable endpoints rather than abstract model names.
Routing policy applies hard denies, preferences, budgets, and tie-break rules.
Observability artifacts record the decision, trace, usage, and observed performance.
The reference router supports hybrid routing across three deployment shapes:
Local / Local - route between local models (e.g., llama-swap peers) based on role, task, and capability
Local / Cloud - route between a local model and a cloud provider based on cost, latency, and task difficulty
Cloud / Cloud - route between cloud providers (e.g., OpenAI, DeepSeek, Moonshot) based on capability and economics
This means a single runtime can serve a quick chat request from a fast local model, route a complex coding
task to a capable cloud model, and fall back to a cheaper cloud endpoint when the primary is degraded, all
under one explainable routing contract.
For end users, prefer the packaged standalone runtime over a source build.
curl -fsSL https://raw.githubusercontent.com/try-works/role-model/main/scripts/install.sh | sh
The installer downloads the latest GitHub Release archive, installs it under
~/.local/share/role-model-router/<version>/<target>/ , and exposes a role-model-router launcher in
~/.local/bin .
irm https: // raw.githubusercontent.com / try - works / role - model / main / scripts / install.ps1 | iex
The installer downloads the latest GitHub Release archive, installs it under
%LOCALAPPDATA%\Programs\RoleModelRouter\<version>\<target>\ , and creates a role-model-router.cmd
launcher.
If you do not want to use installer scripts, download the matching archive from GitHub Releases.
The pi-role-model package connects Pi to an externally running Role-Model runtime.
Start the Role-Model runtime first, then install the public Pi package:
pi install npm:@try-works/pi-role-model
For local checkout testing from this repository, install the package directly:
pi install ./packages/pi-role-model
Inside Pi, run:
/role-model setup
/role-model doctor
/role-model alias list
/role-model alias choose
/role-model alias use <alias>
By default the package connects to http://127.0.0.1:3456 and registers Role-Model as the
role-model provider using /api/role-model/downstream/openai . Set ROLE_MODEL_ENDPOINT
before starting Pi to use a different local runtime. Remote endpoints require explicit
trusted allowRemote behavior, and runtimes that report authentication.required fail
closed unless a future supported token source is configured. For local development installs
and the full command reference, see
packages/pi-role-model/README.md .
Node.js 24 (required for node:sqlite and SEA support)
pnpm 10.x (via corepack enable )
Go 1.24+ (for llama-swap vendor binary and Windows launcher)
corepack enable
corepack pnpm install
Smoke test
corepack pnpm run smoke
For a fuller walkthrough, see docs/public/quickstart.md .
Run the bridge and UI in development mode (separate processes):
# Terminal 1: bridge server
cd role-model-router/apps/runtime-host-bridge
corepack pnpm exec tsx scripts/start-for-qa.ts
# Terminal 2: UI dev server
cd role-model-router/apps/runtime-ui
corepack pnpm exec react-router dev --port 5173 --host 127.0.0.1
Then open http://127.0.0.1:5173 in your browser.
Production build (all platforms)
Build the UI and package the SEA runtime:
# Build UI static files
corepack pnpm --filter @role-model-router/runtime-ui run build
# Package the bridge as a single executable
corepack pnpm run runtime:package-sea
Output: role-model-router/dist/release/<platform-arch>/role-model-runtime
Build a complete Windows package with dedicated browser window:
# 1. Build UI
corepack pnpm --filter @role-model-router/runtime-ui run build
# 2. Package bridge SEA runtime
corepack pnpm run runtime:package-sea
# 3. Build Go launcher
cd role-model-router/apps/launcher
go build -o ../../dist/release/win32-x64/role-model-launcher.exe main.go
# 4. Bundle UI files
cp -r ../runtime-ui/build/client ../../dist/release/win32-x64/
Then double-click role-model-launcher.exe in dist/release/win32-x64/ . It will:
Start the bridge server on port 3456
Open Microsoft Edge in app mode (dedicated window)
Serve the UI directly from the bridge (no separate dev server)
Read this
If you want
docs/public/README.md
the docs hub
docs/public/introduction.md
what role-model is and why it exists
docs/public/quickstart.md
a real end-to-end smoke run
docs/public/concepts/how-role-model-works.md
the system flow
docs/public/concepts/protocol-overview.md
the protocol surface
docs/public/concepts/routing-overview.md
how routing decisions happen
protocol/README.md
canonical schemas and fixtures
role-model-router/README.md
reference router packages and runtime apps
docs/protocol/routing-policy.md
routing policy reference
docs/protocol/taxonomy-v1.md
taxonomy V1 groups, roles, tasks, and Pi classification
docs/protocol/roles.md
role metadata reference
docs/protocol/tasks.md
task metadata reference
docs/operations/02-ci-and-release-flow.md
CI, release automation, and workflow ownership
CHANGELOG.md
release history
Acknowledgements
role-model builds on the work of several open-source projects:
llama-swap - the vendored local model lifecycle manager that handles process supervision, request forwarding, and model swapping for local endpoints
LiteLLM - the unified LLM API abstraction whose provider catalog, model metadata, and pricing data inform the routing-compatible provider inventory
This repository is licensed under BUSL-1.1 with a project-specific
Additional Use Grant. Internal production use, evaluation, development,
modification, and non-production redistribution are permitted under the root
license. Hosted or managed third-party services, paid product embedding, and
third-party commercialization require a separate commercial license.
See LICENSE for the full terms. Contributions require acceptance of
the Contributor License Agreement before they can be merged.
Only individual contributions are accepted. Please do not submit work owned by
an employer, client, company, or other entity unless you personally have the
right to contribute it under this project's terms.
role-model is a protocol for assigning the right model for the right job. Use local and cloud AI together, or route between several cloud providers.
There was an error while loading. Please reload this page .
1
fork
Report repository
Releases
2
v0.0.1-alpha.3
Latest
Jun 27, 2026
+ 1 release
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
