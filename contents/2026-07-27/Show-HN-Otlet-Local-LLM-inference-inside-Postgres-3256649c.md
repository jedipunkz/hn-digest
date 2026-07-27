---
source: "https://github.com/joshmeek/otlet"
hn_url: "https://news.ycombinator.com/item?id=49070611"
title: "Show HN: Otlet – Local LLM inference \"inside\" Postgres"
article_title: "GitHub - joshmeek/otlet: Local LLM inference next to your data · GitHub"
author: "josmek"
captured_at: "2026-07-27T15:38:02Z"
capture_tool: "hn-digest"
hn_id: 49070611
score: 1
comments: 1
posted_at: "2026-07-27T14:56:34Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Otlet – Local LLM inference "inside" Postgres

- HN: [49070611](https://news.ycombinator.com/item?id=49070611)
- Source: [github.com](https://github.com/joshmeek/otlet)
- Score: 1
- Comments: 1
- Posted: 2026-07-27T14:56:34Z

## Translation

タイトル: HN を表示: Otlet – Postgres の「内部」ローカル LLM 推論
記事のタイトル: GitHub - joshmeek/otlet: データの隣のローカル LLM 推論 · GitHub
説明: データの隣にあるローカル LLM 推論。 GitHub でアカウントを作成して、joshmeek/otlet の開発に貢献してください。
HN テキスト: 私は、Postgres の「内部」でローカル LLM 推論を実行するための Postgres 拡張機能である Otlet に取り組んできました (推論はバックグラウンド ワーカーで実行され、Postgres はジョブ、入力、出力、受信を保存します)。これは、エンティティ解決の問題 (テーブル内のどの行が同じエンティティを参照しているかを判断する) に取り組んでいるときに開始しました。データはすでに Postgres にありましたが、モデルを使用すると、行を別のプロセスに取り込み、結果を書き戻す必要がありました。データが存在する場所のすぐ隣にあるデータベースに、そのワークフローをもっと保持できるかどうかを確認したかったのです。現在、Otlet では、SQL からモデル作業を開始し、構造化された出力を要求し、各実行の記録を保持し、ソース データに触れる前に提案された書き込みをレビューすることができます。リポジトリ内の例では、最初のパスには小規模なモデルを使用し、不確実なケースにはより強力なモデルを使用しています。これは私がハッキングしてきたものであり、まだ運用環境での使用を推奨するものではありません。試してみたい人には、Docker のデモがあります。私がこの記事を投稿しているのは、まだ細かく切り刻まれているのを見るのが荒いうちに、アプローチを共有したかったからです。

記事本文:
GitHub - joshmeek/otlet: データの隣のローカル LLM 推論 · GitHub
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
ジョシュミーク
/
オレット
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メインBr

anches タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
207 コミット 207 コミット ベンチマーク ベンチマーク クレート クレート docker/ postgres docker/ postgres docs docs scripts scripts .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md すべてのファイルを表示 リポジトリ ファイル ナビゲーション
Otlet は Postgres 内でローカル LLM 推論を実行する Postgres 拡張機能で、行の隣で読み取って動作します。
私はエンティティ解決の問題のためにそれを構築し始めました。新しいデータが Postgres に到着しましたが、厳しい判断によりデータベースを離れる必要がありました。行を別のシステムにコピーせずに、データがすでに存在する場所でローカル モデルでその作業を実行できるようにしたいと考えていました。
機密行をモデルプロバイダーに送信する必要がないため、推論をデータのそばに置いておく方が安全かつ効率的です。 Postgres の権限、トランザクション、来歴、監査状態はパス内に留まり、アプリケーションは個別のデータ エクスポートと結果取り込みのパイプラインを回避します。その結果として得られる決定とフィードバックは、教育機関が所有するシステム内で複雑化し続けます。
Otlet は、pgrx 拡張機能と常駐 Postgres バックグラウンド ワーカーを使用します。 SQL からモデル作業を要求し、行に対して作業をキューに入れ、ソース変更後に派生状態を更新し、Postgres を離れることなく結果を検査できます。
qwen35_4b が登録されており、customer_notes に行が含まれていると仮定します。 SELECT 結果を otlet.ask に渡します。
選択出力
オレットから。尋ねてください（
' qwen35_4b ' 、
' これらの顧客メモを一文に要約してください。 '、
( SELECT jsonb_agg(to_jsonb(n))
FROM customer_notes n WHERE customer = ' Riverline Labs ' )、
' {"タイプ":"オブジェクト","必須":["概要"],"追加プロパティ":false,"プロパティ":{"概要":{"タイプ":"文字列"}}} '
);
代表的な出力:
出力
--------------------------------------------------

--------------------------------------------------------------------------------------
{"summary": "Riverline Labs は、金曜日までに CSV エクスポート、Otlet のソース テーブル変更の説明、および調達概要を要求しています。"}
(1行)
Otlet は、選択された行の横でモデルを実行し、JSON を検証し、otlet の下にレシートを記録しました。
エンティティ解決では、最初のパスでは安価なローカル モデルが使用され、最初の回答が決定コントラクトを満たさない場合はより強力なローカル モデルが使用されます。ローカル ランタイムを起動し、両方の GGUF ファイルを登録します。
./scripts/otlet-setup.sh
セットアップを再実行すると、Otlet 拡張機能の状態を再構築する際に、PostgreSQL ボリュームとモデル アーティファクトが再利用されます。ユーザーテーブルはそのまま残ります。コンテナーまたはイメージが変更されると、セットアップは拡張機能を再インストールする前に永続的なプリロード状態をクリアします。
名前の選択
オレットから。 register_model ( ' qwen3_1_7b ' 、 ' /var/lib/postgresql/otlet-models/Qwen3-1.7B-Q8_0.gguf ' )
すべてを結合する
名前の選択
オレットから。 register_model ( ' qwen35_4b ' , ' /var/lib/postgresql/otlet-models/Qwen3.5-4B-Q4_K_M.gguf ' );
名前
------------
qwen3_1_7b
qwen35_4b
(2列)
Otlet タスクは、 subject_id および行形式の input を返す SQL クエリを読み取ります。エンティティ解決のチュートリアルでは、2 つのアプリケーション テーブルから public.otlet_demo_vendor_pair_input を構築します。短縮されたタスク呼び出しには、SQL API、出力コントラクト、トレース設定、入力整形、および決定プリセットが含まれます。
SELECT 名、モデル名
オレットから。作成_タスク (
task_name => 'entity_resolution_demo' 、
input_query => $$
SELECT subject_id、入力
パブリックから。 otlet_demo_vendor_pair_input
subject_id で注文
$$、
命令 => ' 各ベンダーのペアが同じエンティティであるかどうかを決定します。一致、信頼度、短い理由、および 1 つの一致する型指定されたアクションを返します。 '、
出力スキーマ => ' {
"タイプ":"オブジェクト",
"必須":["

一致"、"信頼性"、"理由"],
"追加プロパティ":false,
"プロパティ":{
"一致":{"列挙型":["同じエンティティ","異なるエンティティ","不明瞭"]},
"信頼":{"列挙型":["低"、"中"、"高"]},
"理由":{"タイプ":"文字列","最大長":240}
}
} '、
モデル名 => ' qwen3_1_7b ' 、
runtime_options => ' {"max_tokens":256,"reasoning":"off","inference_cache":true,"世代トレース":true,"世代トレース_最大トークン":16,"世代トレース_トップ_k":3} ' ,
input_shaping => ' {"evidence_fields":["candidate_evidence"],"action_id_fields":{"left_id":"left_id","right_id":"right_id"}} ' ,
Decision_contract => ' {"preset":"entity_resolution_evidence_v1"} '
);
SELECT タスク名、安価なモデル名、強力なモデル名
オレットから。 set_model_selection_policy (
「entity_resolution_demo」、「qwen3_1_7b」、「qwen35_4b」
);
SELECT オレット 。 run_task ( 'entity_resolution_demo ') AS queued_jobs;
名前 |モデル名
------------------------+---------------
エンティティ_解像度_デモ | qwen3_1_7b
(1行)
タスク名 |安いモデル名 |強力なモデル名
------------------------+-----------------+-------------------
エンティティ_解像度_デモ | qwen3_1_7b | qwen35_4b
(1行)
キューに入れられたジョブ
-------------
4
(1行)
次のコマンドを使用して、ソース行と完全な命令を含む完全にチェックされたバージョンを実行します。
./scripts/otlet-demo.sh
代表的な受け入れられた出力と入力されたアクション。クエリでは、理由が 48 文字に切り取られます。
r を選択します。 subject_id 、 r 。出力 - >> ' 一致 ' AS 一致、
left( r . 出力 - >> ' 理由 ' , 48 ) AS 理由、
。アクションタイプ
オレットから。 rを実行します
オレットに参加します。アクション a ON a 。 job_id = r 。ジョブID
ここで r 。 task_name = 'entity_resolution_demo'
r で注文します。件名_id ;
件名ID |一致 |理由 |アクションタイプ
------------------------+-----------------+------------------------------------------+---------------
売ります

or-1001:ベンダー-313 |異なるエンティティ |競合する安定した識別子が見つかりました。 |新しいエンティティ
ベンダー-1001:ベンダー-314 |異なるエンティティ | 4 つの競合する安定した識別子が見つかりました。新しいエンティティ
ベンダー-1001:ベンダー-42 |同じエンティティ |同じ送金口座と納税者番号が一致する |マージ候補
ベンダー-1001:ベンダー-77 |異なるエンティティ |競合する安定した識別子が見つかりました。 |新しいエンティティ
(4列)
一致するペアの受信トレースを検査します。安価な出力は JSON スキーマには合格しましたが、より厳格な決定コントラクトには合格しませんでした。オレットはペアを強力なモデルにエスカレーションしました。
SELECT 選択役割、選択ステータス、モデル名、
schema_validation_status、prompt_tokens、generated_tokens、
runtime_fingerprint_hash はフィンガープリントとして NULL ではありません
オレットから。推論受信トレースステータス
WHERE task_name = 'entity_resolution_demo'
AND subject_id = ' ベンダー-1001:ベンダー-42 '
ORDER BY 試行インデックス;
選択_役割 |選択ステータス |モデル名 |スキーマ検証ステータス |プロンプトトークン |生成されたトークン |指紋をとられた
----------+-----------------+----------------------+------------------------------------------+--------------+---------------+--------------
安い |拒否されました | qwen3_1_7b |合格しました | 720 | 27 | t
強い |受け入れられました | qwen35_4b |合格しました | 738 | 126 | t
(2列)
Otlet は両方の試行を記録し、受け入れられた出力から merge_candidate を作成します。このアクションにはオペレーターの承認が必要です。ソースベンダーの行は変更されません。
完全なデモでは、行とペアのウォッチ、候補ドリフト、CustomScan の鮮度、ポータブル ウォッチの定義、および制限された update_row をチェックします。レシートの編集、ロール付与、キャンセル、モデルロードアドミッション、メモリプレッシャー、キャッシュ境界、プロンプトおよびランタイムフィンガープリント、不変条件、Docker クラッシュログをカバーします。
データの隣のローカル LLM 推論
Readme MIT ライセンス アクティビティ スター
0f

orks レポート リポジトリ 使用者
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Local LLM inference next to your data. Contribute to joshmeek/otlet development by creating an account on GitHub.

I’ve been working on Otlet, a Postgres extension for running local LLM inference “inside” Postgres (inference runs in a background worker, while Postgres stores the jobs, inputs, outputs, and receipts) I started it while working on an entity resolution problem (determining which rows in a table refer to the same entity). The data was already in Postgres, but using a model meant pulling rows into another process and then writing the results back. I wanted to see whether I could keep more of that workflow in the database right next to where the data lives Right now Otlet lets you start model work from SQL, require structured output, keep a record of each run, review proposed writes before they touch source data, etc. The example in the repo uses a small model for the first pass and a stronger model for uncertain cases This is just something I've been hacking on, not something I would recommend attempting to use in production yet. There’s a Docker demo if anyone wants to try it. I’m mostly posting it because I wanted to share the approach while it’s still rough to see it get ripped to shreds

GitHub - joshmeek/otlet: Local LLM inference next to your data · GitHub
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
joshmeek
/
otlet
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
207 Commits 207 Commits benchmarks benchmarks crates crates docker/ postgres docker/ postgres docs docs scripts scripts .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md View all files Repository files navigation
Otlet is a Postgres extension that runs local LLM inference inside Postgres , next to the rows it reads and acts on
I started building it for an entity-resolution problem. New data arrived in Postgres, but the hard judgment would have to leave the database. I wanted a local model to do that work where the data already was, without copying rows into another system
Keeping inference beside the data is safer and more efficient because sensitive rows do not need to travel to a model provider. Postgres permissions, transactions, provenance, and audit state stay in the path, and the application avoids a separate data-export and result-ingest pipeline. The resulting decisions and feedback keep compounding in a system the institution owns
Otlet uses a pgrx extension and a resident Postgres background worker. You can ask for model work from SQL, queue work over rows, refresh derived state after source changes, and inspect the result without leaving Postgres
Assume qwen35_4b is registered and customer_notes contains the rows. Pass the SELECT result to otlet.ask :
SELECT output
FROM otlet . ask (
' qwen35_4b ' ,
' Summarize these customer notes in one sentence. ' ,
( SELECT jsonb_agg(to_jsonb(n))
FROM customer_notes n WHERE customer = ' Riverline Labs ' ),
' {"type":"object","required":["summary"],"additionalProperties":false,"properties":{"summary":{"type":"string"}}} '
);
Representative output:
output
----------------------------------------------------------------------------------------------------------------------------------------
{"summary": "Riverline Labs requests CSV export, clarification on Otlet's source table changes, and a procurement summary by Friday."}
(1 row)
Otlet ran the model beside the selected rows, validated the JSON, and recorded a receipt under otlet
Entity resolution uses a cheap local model for the first pass and a stronger local model when the first answer does not meet the decision contract. Start the local runtime, then register both GGUF files:
./scripts/otlet-setup.sh
Rerunning setup reuses the PostgreSQL volume and model artifacts while rebuilding Otlet extension state; user tables stay in place. When the container or image changes, setup clears persisted preload state before reinstalling the extension
SELECT name
FROM otlet . register_model ( ' qwen3_1_7b ' , ' /var/lib/postgresql/otlet-models/Qwen3-1.7B-Q8_0.gguf ' )
UNION ALL
SELECT name
FROM otlet . register_model ( ' qwen35_4b ' , ' /var/lib/postgresql/otlet-models/Qwen3.5-4B-Q4_K_M.gguf ' );
name
------------
qwen3_1_7b
qwen35_4b
(2 rows)
An Otlet task reads any SQL query that returns subject_id and row-shaped input . The entity-resolution walkthrough builds public.otlet_demo_vendor_pair_input from two application tables. The shortened task call includes the SQL API, output contract, trace settings, input shaping, and decision preset:
SELECT name, model_name
FROM otlet . create_task (
task_name => ' entity_resolution_demo ' ,
input_query => $$
SELECT subject_id, input
FROM public . otlet_demo_vendor_pair_input
ORDER BY subject_id
$$,
instruction => ' Decide whether each vendor pair is the same entity. Return match, confidence, a short reason, and one matching typed action. ' ,
output_schema => ' {
"type":"object",
"required":["match","confidence","reason"],
"additionalProperties":false,
"properties":{
"match":{"enum":["same_entity","different_entity","unclear"]},
"confidence":{"enum":["low","medium","high"]},
"reason":{"type":"string","maxLength":240}
}
} ' ,
model_name => ' qwen3_1_7b ' ,
runtime_options => ' {"max_tokens":256,"reasoning":"off","inference_cache":true,"generation_trace":true,"generation_trace_max_tokens":16,"generation_trace_top_k":3} ' ,
input_shaping => ' {"evidence_fields":["candidate_evidence"],"action_id_fields":{"left_id":"left_id","right_id":"right_id"}} ' ,
decision_contract => ' {"preset":"entity_resolution_evidence_v1"} '
);
SELECT task_name, cheap_model_name, strong_model_name
FROM otlet . set_model_selection_policy (
' entity_resolution_demo ' , ' qwen3_1_7b ' , ' qwen35_4b '
);
SELECT otlet . run_task ( ' entity_resolution_demo ' ) AS queued_jobs;
name | model_name
------------------------+------------
entity_resolution_demo | qwen3_1_7b
(1 row)
task_name | cheap_model_name | strong_model_name
------------------------+------------------+-------------------
entity_resolution_demo | qwen3_1_7b | qwen35_4b
(1 row)
queued_jobs
-------------
4
(1 row)
Run the complete checked version, including the source rows and full instruction, with:
./scripts/otlet-demo.sh
Representative accepted outputs and typed actions. The query clips reasons to 48 characters:
SELECT r . subject_id , r . output - >> ' match ' AS match,
left( r . output - >> ' reason ' , 48 ) AS reason,
a . action_type
FROM otlet . runs r
JOIN otlet . actions a ON a . job_id = r . job_id
WHERE r . task_name = ' entity_resolution_demo '
ORDER BY r . subject_id ;
subject_id | match | reason | action_type
------------------------+------------------+------------------------------------------+-----------------
vendor-1001:vendor-313 | different_entity | Conflicting stable identifiers found. | new_entity
vendor-1001:vendor-314 | different_entity | 4 conflicting stable identifiers found | new_entity
vendor-1001:vendor-42 | same_entity | Same remittance account and tax ID match | merge_candidate
vendor-1001:vendor-77 | different_entity | Conflicting stable identifiers found. | new_entity
(4 rows)
Inspect the receipt trace for the matching pair. The cheap output passed JSON Schema and failed the stricter decision contract. Otlet escalated the pair to the strong model:
SELECT selection_role, selection_status, model_name,
schema_validation_status, prompt_tokens, generated_tokens,
runtime_fingerprint_hash IS NOT NULL AS fingerprinted
FROM otlet . inference_receipt_trace_status
WHERE task_name = ' entity_resolution_demo '
AND subject_id = ' vendor-1001:vendor-42 '
ORDER BY attempt_index;
selection_role | selection_status | model_name | schema_validation_status | prompt_tokens | generated_tokens | fingerprinted
----------------+------------------+------------+--------------------------+---------------+------------------+---------------
cheap | rejected | qwen3_1_7b | passed | 720 | 27 | t
strong | accepted | qwen35_4b | passed | 738 | 126 | t
(2 rows)
Otlet records both attempts and creates merge_candidate from the accepted output. The action requires operator approval. The source vendor rows remain unchanged
The full demo checks row and pair watches, candidate drift, CustomScan freshness, portable watch definitions, and bounded update_row . It covers receipt redaction, role grants, cancellation, model-load admission, memory pressure, cache bounds, prompt and runtime fingerprints, invariants, and Docker crash logs
Local LLM inference next to your data
Readme MIT license Activity Stars
0 forks Report repository Used by
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
