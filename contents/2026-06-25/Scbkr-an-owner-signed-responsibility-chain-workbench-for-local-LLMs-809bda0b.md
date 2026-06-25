---
source: "https://github.com/HIJO790401/scbkr-local-responsibility-model"
hn_url: "https://news.ycombinator.com/item?id=48667992"
title: "Scbkr – an owner-signed responsibility-chain workbench for local LLMs"
article_title: "GitHub - HIJO790401/scbkr-local-responsibility-model: Local AI responsibility-chain control layer with owner-signed Workbench, SCBKR S/C/B/K/R grammar, Data Center, four-store evidence reuse, FastAPI, React, and Tauri desktop runtime. · GitHub"
author: "look888"
captured_at: "2026-06-25T03:03:37Z"
capture_tool: "hn-digest"
hn_id: 48667992
score: 1
comments: 0
posted_at: "2026-06-25T02:12:51Z"
tags:
  - hacker-news
  - translated
---

# Scbkr – an owner-signed responsibility-chain workbench for local LLMs

- HN: [48667992](https://news.ycombinator.com/item?id=48667992)
- Source: [github.com](https://github.com/HIJO790401/scbkr-local-responsibility-model)
- Score: 1
- Comments: 0
- Posted: 2026-06-25T02:12:51Z

## Translation

タイトル: Scbkr – ローカル LLM 向けの所有者署名責任チェーン ワークベンチ
記事のタイトル: GitHub - HIJO790401/scbkr-local-responsibility-model: 所有者署名付きワークベンチ、SCBKR S/C/B/K/R 文法、データセンター、4 ストア証拠の再利用、FastAPI、React、および Tauri デスクトップ ランタイムを備えたローカル AI 責任チェーン制御レイヤー。 · GitHub
説明: 所有者署名付きワークベンチ、SCBKR S/C/B/K/R 文法、データセンター、4 ストア証拠の再利用、FastAPI、React、および Tauri デスクトップ ランタイムを備えたローカル AI 責任チェーン制御レイヤー。 - HIJO790401/scbkr-local-responsibility-model

記事本文:
GitHub - HIJO790401/scbkr-local-responsibility-model: 所有者署名付きワークベンチ、SCBKR S/C/B/K/R 文法、データセンター、4 ストア証拠の再利用、FastAPI、React、および Tauri デスクトップ ランタイムを備えたローカル AI 責任チェーン制御レイヤー。 · GitHub
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
{{ メッセージ

}}
ヒジョ790401
/
scbkr-ローカル責任モデル
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
HIJO790401/scbkr-local-responsibility-model
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
177 コミット 177 コミット .github/ workflows .github/ workflows apps apps config config core core data data docs docs docs schemas schemas scripts scripts testing testing .env.example .env.example README.md README.md package.json package.json pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイル ナビゲーション
SCBKR 地域責任連鎖モデル｜地域責任連鎖モデル
所有者署名付きのワークベンチ、データセンター、および 4 つの店舗の証拠を再利用するローカル AI 責任チェーン制御レイヤー。
8. オーナーサインゲート |ユーザー署名ゲート
10. 証拠関係ゲート｜証拠関係基準
SCBKR ローカル責任連鎖モデルは、ローカル AI 責任連鎖制御システムです。
その核心は、モデルが即座に応答できるようにすることではなく、モデルが生成される前に確認、署名、受け入れ、保存、再生できる一連の責任プロセスに入ることができるようにすることです。
チャットは自然言語ポータルです。
ワークベンチは責任連鎖の確認ベンチです。
S / C / B / K / R はタスク責任の構文です。
データセンターは再生可能なデータセンターです。
4 番目のデータベースは、インデックス付け可能なルール層です。
モデルは製図と編集のアシスタントです。
ユーザーが最終署名者です。
SCBKR を使用すると、ユーザーは自分のコンピュータで以下にアクセスできます。
モデルは参加できますが、モデルの権限を越えることはできません。
モデルはドラフトを生成できますが、ルールを有効にするにはユーザーが署名する必要があります。
モデルはデータを参照できますが、署名され、受け入れられ、取り消されていないデータのみを参照できます。
SCBKR は一般的なチャットボットではありません。
SCBKR は単なる RAG ではありません。
SCBKRは大きなモデル会社ではありません。
SCBKR は、モデルがユーザーの答えを直接決定できるツールではありません。
SCBKR はすべての履歴チャットを強化するわけではありません

文脈に沿って詰め込まれた長い会話型の製品。
SCBKR は、モデルがすべてを自動的に記憶できるようにするブラック ボックスの記憶システムではありません。
SCBKR は、モデルがユーザーに代わって自由に動作できるようにするエージェント シェルではありません。
SCBKR は、ローカル AI 責任チェーン制御層のセットです。
モデルをユーザー署名ルールと再生可能なデータセンターで機能させます。
モデルを支援させますが、その権限を超えないようにしてください。
ルールの蓄積を許可しますが、データベースに保存する前にユーザーが署名、受け入れ、再確認する必要があります。
ユーザー入力
→ システムが意図を判断する
→ 確認注文の作成
→ データセンター/Siku へのクエリ
→ 証拠関係ゲートは参照関係を決定します
→ モデル生成 タスクの理解
→ システム編 S / C / B / K / R
→ ワークベンチにドラフトが表示される
→ ユーザー変更・機種変更依頼
→ ユーザーの署名
→ モデル生成
→ ユーザーの受け入れ
→ 倉庫のご提案
→ ユーザーは書き込みを 2 回確認します
→ 4つのデータベースインデックス
→ データセンターでの再生
→ 後続のタスクは署名されたルールを参照します
SCBKR のポイントは、モデルにさらなる自由を与えることではなく、明確な責任の連鎖の中でモデルを機能させることです。
チャット
→ 作業台
→ S / C / B / K / R
→ オーナーのサイン
→ モデル生成
→ ユーザーレビュー
→ 2回目の保管確認
→ データセンター
→ 4店舗
→ 証拠の再利用
→ 新しいタスクコンテキスト
この閉ループの重要なポイントは次のとおりです。
後続のタスクは、所有者が署名し、レビューに合格し、有効期限が切れていない証拠のみを参照します。
C は通常のステップのリストではありません。
C は、タスクが実行できる理由、実行方法、エラーが発生して中断される場所を説明する必要があります。
モデル自体を検証してはなりません。
モデルは自己署名できない場合があります。
モデル単体での受付はできません。
モデルを単独で倉庫に保管することはできません。
モデルはデータセンター自体を変更または削除することはできません。
SCBKR は、採用できるかどうかを決定する前に、まず証拠関係を判断する必要があります。
ユーザーの署名がなければ、SCBKR は確立されません。
受入検査に合格しない場合はお預かりできません。
二次確認がなければ、物理的な書き込みはできません。
モデルは記述、理解、逆アセンブル、コンパイルすることしかできません。

下書き。
ルールが確立されているかどうかは、ユーザーの署名によって決まります。
一般的にチャットしたいだけの場合は、「チャット」に移動してください。
確認チケットを生成することが適切な場合は、提案カードが表示されます。
ユーザーが確認チケットの生成を明示的に要求した場合は、「ワークベンチ」を入力します。
チャットはルールを直接確立する責任を負いません。
ワークベンチは責任連鎖確認プラットフォームです。
8. オーナーサインゲート |ユーザー署名ゲート
オーナー署名ゲートは SCBKR の中核です。
モデルに署名することはできません。
モデルは、アシスタント、システム、モデル、ユーザーなどの偽の文字列を使用して署名を置き換えることはできません。
空の署名は確認できません。
草案は修正後に再署名する必要があります。
SCBKR ドラフトの内容が変更された後は、フロントエンド署名をクリアして、ユーザー署名待ちの状態に戻す必要があります。
確認済み = 偽
署名ステータス = 待機所有者署名
生成結果無効
review_result が無効です
storage_request / storage_plan / storage_result void
これは次のことを保証するためです。
ルールが確立されるたびに、古い署名やモデルの偽の署名を使用するのではなく、ユーザーの現在の署名からルールが作成されます。
SCBKR のデータ センターは表示ページではなく、モデルの将来の作業のためのルール ソース レイヤーです。
10. 証拠関係ゲート｜証拠関係基準
署名ステータス = 所有者署名済み
review_passed = true
ステータスを取り消したり、アーカイブしたり、置き換えたりしてはなりません
関係は direct_match / Same_domain / Same_logic / style_reference である必要があります
的を射るのに一般的な言葉だけに頼らないでください
SBKKR は以下を区別します。
direct_match は正式な基礎として使用できます
Same_domain は同じドメインの基礎として使用できます
like_logic は論理参照として使用できます
style_reference はスタイル参照として使用できます
like_grammar は文法の参照としてのみ使用でき、正式な基礎として使用することはできません。
candidate_only 候補者だが採用されなかった
無関係
競合競合、ユーザーの確認が必要です
引用例:
失効/アーカイブ/置き換えられたデータは使用できません。
like_grammar は、正式な基礎としてではなく、文法上の参照としてのみ使用できます。
署名済みの確認フォーム
承認された結果
すでに在庫があります

ルール
データセンターの録画を再生可能
4 つのデータ データベースにインデックスを付けることができます
次のタスクのためにチャット履歴全体を食べる必要はありません。必要なのは次のことだけです。
ツールは有効化されていないため、実行されたと主張することはできません。
このモデルはテストされていないため、使用可能であるとは言えません。
ユーザーが署名しない場合は確認できません。
責任の連鎖は確認されておらず、発生させてはなりません。
製品が受入検査に合格しなかった場合は、保管してはならない。
2 回目の確認がなければ、物理的な書き込みは許可されません。
失敗出力はメモリを汚染してはなりません。
非ネイティブ モデルの URL は、ネイティブ モデルのふりをしてはなりません。
現在、権限ロックには次のものが含まれます。
external_api = true
危険な操作を確認 = true
それ以外の場合、外部 API を呼び出すことはできません。
プロバイダーが LM Studio / Ollama / local モードとして表示される場合でも、それは外部モデル呼び出しとみなされ、external_api によって承認される必要があります。
API＆オーケストレーション層｜API＆オーケストレーション層
データセンターのクエリ/更新 API
永続化と監査層｜永続化と監査層
Four Stores / 証拠レイヤー｜Four Stores 証拠レイヤー
SCBKR ローカル責任チェーン モデル｜リリース候補は終了しています
SCBKR ローカル責任チェーンモデル｜リリース候補の調整
現在の技術段階:
P15-Pコア閉ループ完了
P15-Q リリース候補 バインディングはまだ完了していません
P15-P 完了した重要なポイント:
model_assisted_structed / scbkr_base_logic /draft_failed ドラフト ソース
LM Studio / Ollama / OpenAI互換のAPIアクセス
Windows デスクトップ プレビュー ビルドの基本
P15-Q はまだ完成していません。
これは製品コンセプトのギャップではなく、リリースおよび発売前のリリース候補の終了項目です。
1. Windows スモーク スクリプトは P15-P ゲートと一致します
古いスモーク スクリプトには、storage-confirm ペイロードに Second_confirm=true がありませんでした。
P15-P バックエンドの storage_confirm ゲートが厳格になっているため、古い煙は正しくブロックされます。
P15-Q は修理が必要です:
storage_confirmed = true
Second_confirm = true
確認済み_by = "ユーザー"
署名 = 明示的なテストのための所有者の署名
テストに合格するには、バックエンド ゲートを下げてはなりません。
所有者の署名が必要です
ストレージ所有者-si

特徴
署名: 「ユーザー」
アシスタント/モデル/システム偽署名
StorageRequest / storageconfirm は以下の前にチェックする必要があります。
ownerSignature.trim() を null にすることはできません
空の署名では、ユーザーに署名の入力を求める必要があり、虚偽の文字列を入力してはなりません。
草案は修正されました。責任の連鎖を確認する前に、ユーザー署名を再入力してください。
4. リリース候補リリースのセマンティクスの終了
現時点では、リポジトリにはまだプレビュー/スケルトン/MVP/未署名/非運用セマンティクスが含まれている可能性があります。
apps/desktop/src-tauri/tauri.conf.json
apps/desktop/src-tauri/Cargo.toml
apps/desktop/src-tauri/src/main.rs
0.15.0-rc.1
P15-Q リリース候補
5. ビルド/スモークスクリプトのリリース
scripts/build_desktop_release_windows.ps1
scripts/smoke_desktop_release_windows.ps1
健康
デスクトップのステータス
サンドボックスモデルのテスト
タスクを作成する
SCBKR ドラフトを作成する
所有者の署名で確認する
モデル生成を有効にする
生成する
レビューパス
ストレージリクエスト
ストレージ確認 (second_confirm=true)
データセンターのセクションを読む
完了
6. 実機受入
チャット
→ 確認フォームの生成
→ ワークベンチドラフト
→モデルの1つのレイヤーを変更します
→ ユーザーが再署名する
→責任連鎖の確認
→ モデル生成
→ ユーザーの受け入れ
→ 倉庫のご提案
→ ユーザーは書き込みを 2 回確認します
→ データセンター利用可能
→ 後続のタスクは所有者の署名付き証拠を参照します
16. 現在の走力
React + Vite + TypeScript Web UI
Windows デスクトップ/Tauri プレビュー ビルド
LM Studio / Ollama / OpenAI互換API設定
携帯電話/リモート デバイスを独自の SCBKR バックエンドに接続するための設計ベース
SCBKR の目標とする使用方法は、開発サーバーに依存して開始するだけではなく、ローカル デスクトップの入り口を提供し、ユーザーが自分のコンピューター上のローカル バックエンド、モデル、データ センターに接続できるようにすることです。
1. デスクトップアプリ/インストーラーモード
2. 開発者モード/開発者モード
1. デスクトップアプリ/インストーラーモード
LM Studio / Ollama / OpenAI互換APIを接続
アクションチャット / ワークベンチ / データセンター / モデル設定 / A

うだうだ
現在、デスクトップ パッケージには Windows/Tauri/sidecar ビルド ベースが含まれています。
P15-Pコア閉ループ完了
P15-Q リリース候補 バインディングはまだ完了していません
P15-Q はまだ完了する必要があります。
Windows スモーク スクリプトの配置 P15-P Second_confirm ゲート
プレビュー/スケルトンからリリース候補までのデスクトップ メタデータ
インストーラー / README / パッケージ バージョンがリリース候補セマンティクスに統合される
したがって、README では現在、このインストーラーが最終製品インストーラーであるとは宣言されていません。
公式リリース候補インストーラーは、P15-Q が閉じられた後にマークされます。
python -m pip install -e 。
フロントエンドの依存関係をインストールします。
npm --prefix apps/web install --package-lock=false
バックエンドを開始します。
python -m uvicorn apps.api.main:app --host 127.0.0.1 --port 8787
フロントエンドを開始します。
npm --prefix apps/web run dev
開く:
http://ローカルホスト:5500
開発者モードはエンド ユーザーの唯一のエントリ ポイントではありません。
これは、API、フロントエンド、モデル配線、SCBKR プロセス、およびデータセンターを検証するために使用されるローカル開発手法です。
http://ローカルホスト:5500
LM Studio の共通ネイティブ エンドポイント:
http://localhost:1234/v1
Ollama OpenAI 互換エンドポイント:
http://localhost:11434/v1
18. よく使用されるテスト
npm --prefix apps/web ビルドの実行
デスクトップ スケルトン/リリース チェック:
npm --prefix apps/desktop run check:skeleton
APIタイトルスモーク：
Python - << 'PY'
apps.api.main インポートアプリから
print(app.title)
ピィ
Windows デスクトップ ビルド/スモークは現在、P15-Q 向けに最終調整中です。
P15-Q は、P15-P の所有者の署名、レビュー ゲート、および 2 番目の確認ストレージ ゲートを揃えます。
powershell -ExecutionPolicy Bypass -ファイル scripts/build_desktop_release_windows.ps1
powershell -ExecutionPolicy Bypass -File scripts/smoke_desktop_release_windows.ps1
19. モデルアクセス
ローカル ループバック エンドポイントはネイティブ モデルとみなすことができます。
非ループバック エンドポイントは、外部モデル ワイヤとして扱う必要があります。
外部モデル接続は external_api/dangerous_operation_conf を経由する必要があります

RMEDガード。
20. 携帯電話およびリモートセルフアクセスの使用
ユーザーは自分のコンピュータ上でローカル バックエンドとローカル モデルを実行できます。
また、携帯電話、タブレット、その他のデバイスを介して SCBKR 作業環境に接続し直します。
電話機には LLM が組み込まれていません。
モバイル版は運用の入り口にすぎません。
実際のモデル、データ センター、4 つのライブラリ、台帳、ストレージは引き続き、ユーザーが指定したローカル バックエンド/デスクトップ サイドカー/API を担当します。
フロントエンド開発サーバーまたはデスクトップ ポータルで LAN アクセスが許可されており、携帯電話とコンピュータが同じ Wi-Fi 環境にある場合は、携帯電話を使用して以下を開くことができます。
http://{コンピュータ エリア ネットワーク IP}:5500
携帯電話を操作ポータルとして使用できます。
SCBKR の長期的な目標は、ユーザーが外部環境のローカル SCBKR / ローカル モデル作業環境に安全に接続できるようにすることです。
外にいる人たち
→ 携帯電話またはラップトップで SCBKR フロントエンドを開きます
→ ユーザーが設定した安全なチャネル経由
→ ホーム/スタジオコンピュータ用のSCBKRバックエンド
→ バックエンドに接続し、ローカルの LM Studio / Ollama / OpenAI 互換 API に接続します
→ チャット/ワークベンチ/署名/受諾/倉庫/データセンタークエリを完了
リモート セルフ アクセスは、次のようなユーザー設定を通じて実現できます。
プロバイダーが LM Studio / Ollama / local モードとして表示される場合でも、外部または非ループバック モデル ワイヤとして扱われ、external_api /angered_operation_confirmed ガードを通過する必要があります。
ユーザーが署名しない場合は確認できません。
責任の連鎖は確認されておらず、発生させてはなりません。
製品が受入検査に合格しなかった場合は、保管してはならない。
2 回目の確認がなければ、物理的な書き込みは許可されません。
モデルは自己署名できない場合があります。
モデル単体での受付はできません。
モデルを単独で倉庫に保管することはできません。
携帯電話、リモート ブラウザ、または外部デバイスは、操作の入り口にすぎません。
オーナーサインゲート
レビューゲート
保管確認ゲート
証拠関係ゲート
データセンター / Siku
レジャー/ハッシュ/リプレイ
共有コントロール。
GitHub リポジトリの右側にある About → Topics に記入することをお勧めします。
local-llm llm-agent ai-safety fastapi tauri デスク

op-app ollama lm-studio openai-compatibility local-first ai-workbench workflow-automation Agentic-workflow Audit-log rag
這いトピック對想：
SCBKR ローカル責任連鎖モデルは、ローカル AI 責任連鎖制御システムです。
一般的なチャットボットではありません。
これは単純な RAG ツールではありません。
モデル会社ではありません。
モデルがユーザーに代わって直接決定し、行動することはできません。
SCBKR は、モデルを生成する前に責任チェーン ワークフローに入れます。
自然言語エントリとしてのチャット
確認面としての作業台
責任連鎖文法としての S / C / B / K / R
再生可能なストレージ層としてのデータセンター
再利用可能なインデックス付きルール ストアとしての 4 つのストア
ルール終了条件としての所有者の署名
モデルは、支援、草案作成、記述、およびコンパイルを行うことができます。
モデルは、それ自体では確認、署名、レビュー、保存、更新、または削除を行うことができません。
ルールは所有者の署名後にのみ有効になります。
結果はレビューと 2 回目の確認後にのみ保存できます。
今後のタスクでは、所有者が署名し、レビューに合格し、取り消し/アーカイブ/置き換えられていない証拠のみを再利用できます。
P15-P炉心閉鎖完了。
P15-Q リリース候補の調整が保留中です。
二十三、產品原則
モデルは先の回答ではなく、先の交換です。
モデルは先生成ではなく先確認です。
モデルは変化する可能性がありますが、それ以上変形することはできません。
規則はモデル成立ではなく、使用者名成立です。
資料は自動記憶ではなく、庫内に保管されます。
引用は關鍵字命中、むしろ証拠関係です。
聊天は無限上下文ではなく、4庫インデックスの責任ある回収です。
遠端の操作はモデルを変えるのではなく、使用者が責任を持って行う環境です。
二十四、簽名
セマンティック ファイアウォールの創設者
Wen-Yao Hsu / ShenYao888π
所有者署名付きワークベンチ、SCBKR S/C/B/K/R を備えたローカル AI 責任チェーン制御層

文法、データセンター、4 ストア証拠の再利用、FastAPI、React、および Tauri デスクトップ ランタイム。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Local AI responsibility-chain control layer with owner-signed Workbench, SCBKR S/C/B/K/R grammar, Data Center, four-store evidence reuse, FastAPI, React, and Tauri desktop runtime. - HIJO790401/scbkr-local-responsibility-model

GitHub - HIJO790401/scbkr-local-responsibility-model: Local AI responsibility-chain control layer with owner-signed Workbench, SCBKR S/C/B/K/R grammar, Data Center, four-store evidence reuse, FastAPI, React, and Tauri desktop runtime. · GitHub
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
HIJO790401
/
scbkr-local-responsibility-model
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
HIJO790401/scbkr-local-responsibility-model
main Branches Tags Go to file Code Open more actions menu Folders and files
177 Commits 177 Commits .github/ workflows .github/ workflows apps apps config config core core data data docs docs schemas schemas scripts scripts tests tests .env.example .env.example README.md README.md package.json package.json pyproject.toml pyproject.toml View all files Repository files navigation
SCBKR 本地責任鏈模型｜Local Responsibility Chain Model
Local AI responsibility-chain control layer with owner-signed Workbench, Data Center, and four-store evidence reuse.
八、Owner Signature Gate｜使用者簽名閘門
十、Evidence Relation Gate｜證據關係判準
SCBKR 本地責任鏈模型 是一套本地 AI 責任鏈控制系統。
它的核心不是讓模型立刻回答，而是讓模型在生成之前，先進入可確認、可簽名、可驗收、可入庫、可回放的責任鏈流程。
Chat 是自然語言入口。
Workbench 是責任鏈確認台。
S / C / B / K / R 是任務責任語法。
Data Center 是可回放資料中心。
四庫是可索引規則層。
模型是草案與編譯助手。
使用者是最終簽名者。
SCBKR 允許使用者在自己的電腦上接入：
模型可以參與，但模型不能越權。
模型可以生成草案，但規則必須由使用者簽名才成立。
模型可以引用資料，但只能引用已簽名、已驗收、未撤銷的資料。
SCBKR 不是一般聊天機器人。
SCBKR 不是單純 RAG。
SCBKR 不是大模型公司。
SCBKR 不是讓模型替使用者直接決定答案的工具。
SCBKR 不是把所有歷史聊天硬塞進上下文的長對話產品。
SCBKR 不是讓模型自動記憶一切的黑箱記憶系統。
SCBKR 不是讓模型自由代理使用者行動的 agent 外殼。
SCBKR 是一套 Local AI Responsibility Chain Control Layer ：
讓模型在使用者簽名規則與可回放資料中心中工作。
讓模型能輔助，但不能越權。
讓規則能累積，但必須經過使用者簽名、驗收與二次確認入庫。
使用者輸入
→ 系統判斷 intent
→ 建立確認單
→ 查詢 Data Center / 四庫
→ Evidence Relation Gate 判斷引用關係
→ 模型產生 Task Understanding
→ 系統編譯 S / C / B / K / R
→ Workbench 顯示草案
→ 使用者修改 / 要求模型修改
→ 使用者簽名
→ 模型生成
→ 使用者驗收
→ 入庫建議
→ 使用者二次確認寫入
→ 四庫索引
→ Data Center 回放
→ 後續任務引用已簽名規則
SCBKR 的重點不是讓模型更自由，而是讓模型在明確的責任鏈中工作。
Chat
→ Workbench
→ S / C / B / K / R
→ Owner Signature
→ Model Generation
→ User Review
→ Second Confirm Storage
→ Data Center
→ Four Stores
→ Evidence Reuse
→ New Task Context
這個閉環的重點是：
後續任務只引用 owner-signed、review-passed、未失效的 evidence。
C 不是普通步驟清單。
C 要說明任務為什麼可以被執行、如何執行、錯在哪裡會中斷。
模型不得自行確認。
模型不得自行簽名。
模型不得自行驗收。
模型不得自行入庫。
模型不得自行修改或刪除 Data Center。
SCBKR 必須先判斷 evidence relation，再決定是否能採用。
沒有使用者簽名，SCBKR 不成立。
沒有驗收通過，不能入庫。
沒有二次確認，不能 physical write。
模型只能描述、理解、拆解、編譯草案。
規則是否成立，由使用者簽名決定。
如果只是一般聊天，走 Chat。
如果適合生成確認單，顯示建議卡。
如果使用者明確要求生成確認單，進入 Workbench。
Chat 不負責直接成立規則。
Workbench 才是責任鏈確認台。
八、Owner Signature Gate｜使用者簽名閘門
Owner Signature Gate 是 SCBKR 的核心。
模型不能簽名。
模型不能用 assistant 、 system 、 model 、 user 這類假字串代替簽名。
空簽名不得 confirmed。
修改草案後必須重新簽名。
任何 SCBKR 草案內容變更後，都必須清空前端簽名並回到等待使用者簽名狀態。
confirmed = false
signature_status = waiting_owner_signature
generation_result 作廢
review_result 作廢
storage_request / storage_plan / storage_result 作廢
這是為了確保：
每一次規則成立，都來自使用者當下明確簽名，而不是沿用舊簽名或模型假簽名。
SCBKR 的 Data Center 不是展示頁，而是模型未來工作的規則來源層。
十、Evidence Relation Gate｜證據關係判準
signature_status = owner_signed
review_passed = true
status 不得為 revoked / archived / superseded
relation 必須是 direct_match / same_domain / similar_logic / style_reference
不得只靠泛詞命中
SCBKR 會區分：
direct_match 可作為正式依據
same_domain 可作為同領域依據
similar_logic 可作為邏輯參考
style_reference 可作為風格參考
similar_grammar 只能參考語法，不得作為正式依據
candidate_only 候選但不採用
irrelevant 不相關
conflict 衝突，需使用者確認
引用範例：
revoked / archived / superseded 資料不可採用。
similar_grammar 只能作為語法參考，不得作為正式依據。
已簽名確認單
已驗收結果
已入庫規則
可回放 Data Center 紀錄
可索引四庫資料
下一次任務不需要吃完整聊天歷史，只需要：
工具未啟用，不得宣稱已執行。
模型未測通，不得宣稱可用。
使用者未簽名，不得 confirmed。
責任鏈未確認，不得生成。
驗收未通過，不得入庫。
未二次確認，不得 physical write。
失敗輸出不得污染記憶。
非本機模型網址不得假裝成本地模型。
目前權限鎖包含：
external_api = true
dangerous_operation_confirmed = true
否則不得呼叫外部 API。
即使 provider 顯示為 LM Studio / Ollama / local mode，也必須視為外部模型呼叫，必須經過 external_api 授權。
API & Orchestration Layer｜API 與協調層
Data Center Query / Update API
Persistence & Audit Layer｜持久化與審計層
Four Stores / Evidence Layer｜四庫證據層
SCBKR 本地責任鏈模型｜Release Candidate 收束中
SCBKR Local Responsibility Chain Model｜Release Candidate Alignment
目前技術階段：
P15-P 核心閉環已完成
P15-Q Release Candidate 收束尚未完成
P15-P 已完成重點：
model_assisted_structured / scbkr_base_logic / draft_failed 草案來源
LM Studio / Ollama / OpenAI-compatible API 接入
Windows desktop preview build 基礎
P15-Q 尚未完成。
這不是產品概念缺口，而是 Release Candidate 發行與上線前收束項。
1. Windows smoke script 對齊 P15-P gate
舊 smoke script 在 storage-confirm payload 缺少 second_confirm=true。
P15-P 後端 storage_confirm gate 已變嚴格，因此舊 smoke 會被正確擋下。
P15-Q 需要修：
storage_confirmed = true
second_confirm = true
confirmed_by = "user"
signature = 明確測試用 owner signature
不得為了測試通過而降低後端 gate。
owner-signature-required
storage-owner-signature
signature: "user"
assistant / model / system 假簽名
storageRequest / storageConfirm 前必須檢查：
ownerSignature.trim() 不可為空
空簽名必須提示使用者輸入簽名，不得補假字串。
草案已修改，請重新輸入使用者簽名後再確認責任鏈。
4. Release Candidate 發行語意收束
目前 repo 仍可能殘留 preview / skeleton / MVP / unsigned / not production 語意。
apps/desktop/src-tauri/tauri.conf.json
apps/desktop/src-tauri/Cargo.toml
apps/desktop/src-tauri/src/main.rs
0.15.0-rc.1
P15-Q Release Candidate
5. Release build / smoke script
scripts/build_desktop_release_windows.ps1
scripts/smoke_desktop_release_windows.ps1
health
desktop status
sandbox model test
create task
create SCBKR draft
confirm with owner signature
enable model_generate
generate
review pass
storage-request
storage-confirm with second_confirm=true
Data Center section read
complete
6. 實機驗收
Chat
→ 生成確認單
→ Workbench 草案
→ 模型修改一層
→ 使用者重新簽名
→ 確認責任鏈
→ 模型生成
→ 使用者驗收
→ 入庫建議
→ 使用者二次確認寫入
→ Data Center 可查
→ 後續任務引用 owner-signed evidence
十六、目前可跑能力
React + Vite + TypeScript Web UI
Windows Desktop / Tauri preview build
LM Studio / Ollama / OpenAI-compatible API 設定
手機 / 遠端裝置連回自有 SCBKR backend 的設計基礎
SCBKR 的目標使用方式不是只靠開發伺服器啟動，而是提供本地桌面入口，讓使用者在自己的電腦上連接本地後端、模型與資料中心。
1. Desktop App / Installer 模式
2. Developer Mode / 開發者模式
1. Desktop App / Installer 模式
連接 LM Studio / Ollama / OpenAI-compatible API
操作 Chat / Workbench / Data Center / Model Settings / Audit
目前桌面封裝已有 Windows / Tauri / sidecar build 基礎。
P15-P 核心閉環已完成
P15-Q Release Candidate 收束尚未完成
P15-Q 尚需完成：
Windows smoke script 對齊 P15-P second_confirm gate
desktop metadata 從 preview / skeleton 收束為 Release Candidate
installer / README / package version 統一到 Release Candidate 語意
因此目前 README 不將 installer 宣稱為最終 production installer。
正式 Release Candidate installer 會在 P15-Q 收束後標記。
python -m pip install -e .
安裝前端依賴：
npm --prefix apps/web install --package-lock=false
啟動後端：
python -m uvicorn apps.api.main:app --host 127.0.0.1 --port 8787
啟動前端：
npm --prefix apps/web run dev
打開：
http://localhost:5500
開發者模式不是最終使用者唯一入口。
它是用來驗證 API、前端、模型連線、SCBKR 流程與 Data Center 的本地開發方式。
http://localhost:5500
LM Studio 常見本機 endpoint：
http://localhost:1234/v1
Ollama OpenAI-compatible endpoint：
http://localhost:11434/v1
十八、常用測試
npm --prefix apps/web run build
Desktop skeleton / release check：
npm --prefix apps/desktop run check:skeleton
API title smoke：
python - << ' PY '
from apps.api.main import app
print(app.title)
PY
Windows desktop build / smoke 目前正在 P15-Q 收束中。
P15-Q 將對齊 P15-P 的 owner signature、review gate、second confirm storage gate。
powershell -ExecutionPolicy Bypass -File scripts/build_desktop_release_windows.ps1
powershell -ExecutionPolicy Bypass -File scripts/smoke_desktop_release_windows.ps1
十九、模型接入
本地 loopback endpoint 可視為本機模型。
非 loopback endpoint 必須視為外部模型連線。
外部模型連線必須經過 external_api / dangerous_operation_confirmed guard。
二十、手機與遠端自接入使用
使用者可以在自己的電腦上運行本地後端與本地模型，
並透過手機、平板或其他裝置連回自己的 SCBKR 工作環境。
手機不是內建 LLM。
手機端只是操作入口。
實際模型、資料中心、四庫、ledger、storage 仍然由使用者指定的本地後端 / desktop sidecar / API 負責。
若前端開發伺服器或桌面入口允許 LAN 存取，手機與電腦在同一個 Wi-Fi 下，可用手機打開：
http://{電腦區網IP}:5500
手機可以作為操作入口：
SCBKR 的長期目標是支援使用者在外部環境中，仍能安全連回自己的本地 SCBKR / 本地模型工作環境。
人在外面
→ 手機或筆電開啟 SCBKR 前端
→ 透過使用者自行設定的安全通道
→ 連回家中 / 工作室電腦的 SCBKR backend
→ backend 再連接本地 LM Studio / Ollama / OpenAI-compatible API
→ 完成 Chat / Workbench / 簽名 / 驗收 / 入庫 / Data Center 查詢
遠端自接入可以透過使用者自行設定的方式實現，例如：
即使 provider 顯示為 LM Studio / Ollama / local mode，也必須視為外部或非 loopback 模型連線，並經過 external_api / dangerous_operation_confirmed guard。
使用者未簽名，不得 confirmed。
責任鏈未確認，不得生成。
驗收未通過，不得入庫。
未二次確認，不得 physical write。
模型不得自行簽名。
模型不得自行驗收。
模型不得自行入庫。
手機端、遠端瀏覽器或外部裝置只是操作入口。
Owner Signature Gate
Review Gate
Storage Confirm Gate
Evidence Relation Gate
Data Center / 四庫
ledger / hash / replay
共同控制。
建議在 GitHub repo 右側 About → Topics 填入：
local-llm llm-agent ai-safety fastapi tauri desktop-app ollama lm-studio openai-compatible local-first ai-workbench workflow-automation agentic-workflow audit-log rag
這些 topic 對應：
SCBKR Local Responsibility Chain Model is a local AI responsibility-chain control system.
It is not a general chatbot.
It is not a simple RAG tool.
It is not a model company.
It does not allow the model to directly decide and act on behalf of the user.
SCBKR makes the model enter a responsibility-chain workflow before generation.
Chat as the natural-language entry
Workbench as the confirmation surface
S / C / B / K / R as the responsibility-chain grammar
Data Center as the replayable storage layer
Four stores as reusable indexed rule stores
Owner signature as the rule closure condition
The model can assist, draft, describe, and compile.
The model cannot confirm, sign, review, store, update, or delete by itself.
A rule only becomes valid after owner signature.
A result can only be stored after review and second confirmation.
Future tasks can only reuse evidence that is owner-signed, review-passed, and not revoked / archived / superseded.
P15-P core closure completed.
P15-Q Release Candidate alignment pending.
二十三、產品原則
模型不是先回答，而是先交代。
模型不是先生成，而是先確認。
模型可以參與，但不能越權。
規則不是模型成立，而是使用者簽名成立。
資料不是自動記憶，而是驗收後入庫。
引用不是關鍵字命中，而是 evidence relation。
聊天不是無限上下文，而是四庫索引與責任鏈回放。
遠端操作不是模型越權，而是使用者連回自己的責任鏈環境。
二十四、簽名
Founder of Semantic Firewall
Wen-Yao Hsu / ShenYao888π
Local AI responsibility-chain control layer with owner-signed Workbench, SCBKR S/C/B/K/R grammar, Data Center, four-store evidence reuse, FastAPI, React, and Tauri desktop runtime.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
