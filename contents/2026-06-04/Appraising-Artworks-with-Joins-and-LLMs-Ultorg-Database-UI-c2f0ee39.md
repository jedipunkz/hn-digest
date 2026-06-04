---
source: "https://www.ultorg.com/docs/edit/fill-with-ai/"
hn_url: "https://news.ycombinator.com/item?id=48397868"
title: "Appraising Artworks with Joins and LLMs (Ultorg Database UI)"
article_title: "Ultorg - Fill with AI"
author: "eirikbakke"
captured_at: "2026-06-04T13:13:49Z"
capture_tool: "hn-digest"
hn_id: 48397868
score: 1
comments: 1
posted_at: "2026-06-04T12:45:18Z"
tags:
  - hacker-news
  - translated
---

# Appraising Artworks with Joins and LLMs (Ultorg Database UI)

- HN: [48397868](https://news.ycombinator.com/item?id=48397868)
- Source: [www.ultorg.com](https://www.ultorg.com/docs/edit/fill-with-ai/)
- Score: 1
- Comments: 1
- Posted: 2026-06-04T12:45:18Z

## Translation

タイトル: 結合と LLM を使用したアートワークの評価 (Ultorg データベース UI)
記事のタイトル: Ultorg - AI で満たす
説明: Fill with AI は、人工知能を使用して、既存のデータと英語の指示に基づいて、Ultorg の観点で編集を段階的に行います。

記事本文:
データへの接続
データソースの基本
視点を構築する
基本的なインタラクション
データの提示
自動レイアウトエンジン
データの編集
データ編集の基本
保存と共有
パースペクティブの保存
その他のトピックス
インストールに関する注意事項
AI で埋めるアクションは、人工知能を使用して、Ultorg の観点でデータ編集を段階的に行います。パースペクティブに表示されるデータ (結合や 1 対多の関係など) は、英語で指定された指示とともに大規模言語モデル (LLM) に送信されます。
デモ: 結合と LLM を使用したアートワークの評価
このデモでは、OpenAI (gpt-4.1-mini) の LLM と Ultorg の Fill with AI 機能を使用して、ナショナル ギャラリー オブ アートのデータベース内の美術品の金額を推定します。
ビジュアル結合 を使用して、データベース内の他のテーブルから各アートワークのタイムライン エントリやその他の関連情報を取り込みます。この情報は、評価タスクで使用するために LLM に渡されます。結果は、追加のカスタム結合を通じて Excel ファイルに書き込まれます。
注: この機能には、Ultorg バージョン 2.2.0 以降 (2026-06-04) と Anthropic または OpenAI の API キーが必要です。
使用方法
Fill with AI を使用して自動データ編集を行うには:
現在のパースペクティブで、フィールドと結合を調整して、LLM にコンテキストの一部として表示するデータと、編集するフィールドを表示します。
上記のデモの例では、芸術作品のリストと、それぞれの書誌エントリを含むサブクエリ、および結果を書き込むテーブルを取得するための別の Appraisals サブクエリを示しました。
この例のように、結果を保存するために別のテーブルを使用する場合、AI による入力は既存の行のみを編集するため、処理する ID の行をテーブルに事前に入力する必要があります。私たちのデモでは、「評価」

ble は次のような Excel シートからのものでした。
(編集しているテーブルがクエリに行を提供しているテーブルと同じかどうかは無視してください。)
必要に応じて複数選択を使用して、編集するフィールドを選択します。
編集する列は、パースペクティブ内の任意のサブクエリに含めることができます。たとえば、上記のパースペクティブで「オブジェクト テキスト エントリ」の下のフィールドを選択することも同様にできます。
[編集] メニューの [AI で埋める] をクリックします。
この機能を初めて使用する場合は、AI プロバイダーを構成するように求められます。続行して、必要に応じて AI プロバイダーの API キーを指定します。
編集している各列にプロンプ​​トを表示します。以下の例を参照してください。たとえば、芸術作品を評価するには、次のようなものを使用します。
Appraisal USD : 「保険を目的とした、米ドルでの作品の評価」
評価メモ：「評価理由」
必要に応じて、その下の大きなテキスト フィールドに追加の指示を入力できます。
必要に応じて、使用する特定の AI モデルを [モデル] ドロップダウンから選択します。利用可能なモデルは AI プロバイダーによって異なり、速度、価格、機能が異なります。
Ultorg は、既知のデフォルト モデルが利用可能な場合はそれを選択し、互換性がないことがわかっているモデルを非表示にします。現在のデフォルトは、Anthropic の場合は claude-sonnet-4-6、OpenAI の場合は gpt-4.1-mini です。新しいモデルが頻繁にリリースされます。
[OK] を押します。 AI で生成された編集は、[AI で塗りつぶす] ダイアログが開いたままになっている間、徐々に表示されます。操作の実行中、パースペクティブはある程度の対話性を保持します。上下にスクロールしたり、フィールドを表示または非表示にしたりすることができます。
最後の編集が完了する前に [キャンセル] を押すと、それまでにステージングされた編集が残ります。必要に応じて、それらを元に戻すことができます (Ctrl+Z)。
全体的な操作が完了すると、LLM からの短い終了メッセージが表示されます。このメッセージにより、次のような問題が表面化する可能性があります。

プロンプトのあいまいさ、判断が難しいケース、または処理中に行われた仮定。
データが複数のバッチで処理された場合、最後に表示される概要は、各バッチの概要をまとめたものになります。
段階的な編集をデータ ソースに書き戻すには、[編集を適用] を押します。
提案された編集の前後の値を比較するには、[編集を破棄] を押してから、[元に戻す] と [やり直し] (Ctrl+Z/Y) を切り替えます。
AI による入力機能は、数値、日付、タイムスタンプなどの標準データ型の列で動作します。 Ultorg は、それぞれの場合に正しい出力形式に関する指示を LLM に渡します。このような指示を自分で書く必要はありません。
プロンプトの例
通常、各列を編集するには短いプロンプトで十分です。例:
「スウェーデン語に翻訳して」
「架空のサポート チケットの説明を 20 個生成します。」
(この種のプロンプトの場合、行を入力する前に、行がすでに存在しているか、Ctrl+Shift+Plus を使用して挿入されるようにステージングされている必要があることに注意してください。)
広告コピーのブレインストーミング:
「各カテゴリで、現在空白の行に追加の見出しをいくつか考えてください。」
上のスクリーンショットでは、 Fill with AI を呼び出す前に、 Insert Record ( ) アクションを使用して挿入をステージングしています。
隣接する「住所」列から通りの名前を抽出するには:
通り名 : 「住所から通り名を抽出します。」
営業見込み客とやり取りのデータベースに真偽の「エンゲージメント」フィールドを入力するには、次の手順を実行します。
エンゲージメント : 「このタイムライン エントリにその人側からエンゲージメントが含まれている場合は、このフィールドを true に設定します。たとえば、その人からの電子メール、またはその人との会議は「true」と見なされます。しかし、私が自分で送信する電子メールは、その人が応答したことを確認しない限り、「true」と見なされません。購読解除はエンゲージメントとしてカウントしません。エントリがそうでない場合は、

true に設定する場合は、値を false に設定します。」
セールスリードの資格:
リードスコア: 過去の履歴やその他の利用可能なデータに基づいて、各セールスリードに 1 ～ 10 のスコアを割り当てます。スコアが高いほど、現在の支援活動が優先されることを意味します。」
根拠 : 「与えられたリードスコアの根拠」。
ユーザーや作成者などのテーブルを軽く匿名化するには:
名 : 「性別が同じで音節数が同じで、最初の文字が同じである、ランダムな (ただし常に異なる) 名」
姓: 「最初の文字が同じで、音節数が同じで、ランダムな (ただし常に異なる) 姓」
LLM が複数の行にわたってコンテキストを維持する必要がある例:
「リストが進むにつれて、作品の良さについてますます懐疑的になる美術評論家による簡潔なメモ。リストの最後までに、査読者は品質の欠如に完全に憤慨しています。」
「1から20までの数字のリスト」
後者の種類のプロンプトの場合は、パースペクティブ内のデータの量がバッチ処理を引き起こすほど大きくないことを確認してください。これは、コンテキスト内の行を分離することになるためです。
Ultorg は、LLM に列編集タスク全体を理解させるために必要なプロンプトの残りの部分を提供します。
AI プロバイダーの構成
Ultorg デスクトップ アプリケーションは、中間サーバーを経由せずに、選択した AI プロバイダーに直接接続します。 [AI プロバイダーの構成] ダイアログでは、選択したプロバイダーを構成できます。
ドロップダウンからプロバイダーの種類を選択し、API キーを入力します。 「接続のテスト」を押して構成を検証します。
現在サポートされているプロバイダーは次のとおりです。
人間的。 API キーはここで作成できます。有料サービス。
オープンAI。 API キーはここで作成できます。有料サービス。
O と互換性のあるその他のプロバイダー

ペンAI API。
このようなプロバイダーの場合は、OpenAI を選択し、API ベース URL を適切な設定に変更します。
モデルはストリーミングおよび構造化された JSON 出力をサポートする必要があります。
API キーはオペレーティング システムのユーザー ディレクトリに保存されます。
セマンティクス
LLM に送信されるデータには次のものが含まれます。
自動編集タスクに関する一般的な手順。Ultorg にハードコードされています。
[AI を使用して入力] ダイアログで [OK] をクリックした時点でパースペクティブに表示されるデータ。既存の段階的なデータ編集が組み込まれています。バッチ処理が使用される場合、一度にルートレベルのレコードのサブセットのみが送信されます。
Ultorg で表示されるフィールド名。データ ソース レベルに存在する技術テーブルの列名は意図的に省略されています。
編集する各列について:
Ultorg に表示される列の名前。
列のデータ型と、そのデータ型の値を指定するための手順。
列が null 値をサポートするかどうか、および null 値を出力する方法。
列に指定する指示。
[AI で入力] ダイアログでオプションで指定する追加の指示。
LLM はすべての行を編集する義務はありません。ただし、編集できるのは、「OK」を押した時点でパースペクティブに表示されていた行のみです。下にスクロールしてさらに行をロードできます。いずれの場合も、選択した列のみが編集されます。
操作の実行中もパースペクティブでの作業を続けることができます。特に、新しいデータを編集すると操作が中断されます。
テキスト フィールドは、LLM から要求された出力形式の他のフィールドよりも前に自動的に順序付けされます。これにより、LLM が後で数値スコアなどを生成するときに、「根拠」タイプのフィールドの推論を最近のコンテキストに含めることができます。
バッチ処理
LLM は一度に処理できるデータ量に制限があります。

このコンテキストが絶対制限に近づきすぎない場合に最高のパフォーマンスが得られます。
AI による入力機能の場合、100 KB (約 30,000 個の「トークン」) を超えるパースペクティブ データは、自動的に小さなバッチに分割されます。ダイアログ ボックスに次のような通知が表示されます。
バッチ処理は透過的に行われるため、通常は心配する必要はありません。ただし、LLM はいつでもパースペクティブ内のすべての行を認識できるわけではないことを意味します。
一括編集
データの編集
編集オプション
contact@ultorg.com
Twitter / : @Ultorg
© 2026 Ultorg Inc.

## Original Extract

Fill with AI uses artificial intelligence to stage edits in an Ultorg perspective, based on existing data and your instructions in English.

Connecting to Data
Data Source Basics
Building Perspectives
Basic Interactions
Presenting Data
Auto-Layout Engine
Editing Data
Data Editing Basics
Saving and Sharing
Saving Perspectives
Other Topics
Installation Notes
The Fill with AI action uses artificial intelligence to stage data edits in an Ultorg perspective . The data showing in your perspective, including joins and one-to-many relationships , is sent to a Large Language Model (LLM) along with instructions that you provide in English.
Demo: Appraising Artworks with Joins and LLMs
In this demo, we estimate the dollar value of artworks in a database from the National Gallery of Art, using an LLM from OpenAI (gpt-4.1-mini) and the Fill with AI feature in Ultorg.
Using visual joins , we bring in timeline entries for each artwork, and other relevant information, from other tables in the database. This information gets passed to the LLM for use in the appraisal task. Results are written to an Excel file through an additional Custom Join .
Note: This feature is requires Ultorg version 2.2.0 or above (2026-06-04) and an API key from either Anthropic or OpenAI .
How to Use
To make automated data edits using Fill with AI :
In your current perspective, adjust fields and joins to show the data that you want the LLM to see as part of its context, as well as the field(s) to be edited.
In the example from our demo above, we showed a list of artworks and a subquery with bibliographical entries for each, as well as a separate Appraisals subquery to pull in the table that we want to write our results to:
If you use a separate table to store results, as in this example, it must be pre-populated with rows for the IDs you want to process, as Fill with AI will only edit existing rows. In our demo, the “Appraisals” table was from an Excel sheet like the following:
(Disregard if the table you are editing is the same one that is providing rows for the query.)
Select the field(s) you want to edit, using multiple selection if needed.
The columns to be edited can be in any subquery within the perspective. For example, we could equally well select a field under “Objects Text Entries” in the perspective above.
Click Fill with AI in the Edit menu.
If this is the first time you are using this feature, you will be asked to configure your AI provider. Proceed and provide an AI provider's API key as needed.
Provide a prompt for each column you are editing. See examples below. To appraise artworks, for instance, we might use:
Appraisal USD : “an appraisal of the artwork in US dollars, for insurance purposes”
Appraisal Notes : “rationale for the appraisal”
Additional instructions can be provided, if needed, in the larger text field underneath.
If desired, pick a specific AI model to use from the Model dropdown. Available models depend on your AI provider, and vary in speed, price, and capability.
Ultorg will pick a known default model if available, and hide models that are known to be incompatible. Current defaults are claude-sonnet-4-6 for Anthropic and gpt-4.1-mini for OpenAI. New models are frequently released.
Press OK . AI-generated edits will appear progressively while the Fill with AI dialog remains open. The perspective retains some interactivity while the operation runs; you can scroll up and down, hide or show fields, and such.
If you press Cancel before the last edit completes, the edits staged so far will remain. You can undo them if desired (Ctrl+Z).
When the overall operation completes, a short concluding message from the LLM will be displayed. This message may surface ambiguities in the prompt, cases that were hard to decide, or assumptions made during processing.
If data was processed in multiple batches , then the summary displayed at the end will be a summary of summaries from each batch.
To write staged edits back to the data source, press Apply Edits .
To compare values before and after the proposed edits, you can press Discard Edits and then toggle with Undo and Redo (Ctrl+Z/Y).
The Fill with AI feature works with columns of any standard data type , including numbers, dates, and timestamps. Ultorg will pass instructions to the LLM regarding the correct output format in each case. You do not have to write such instructions yourself.
Example Prompts
A short prompt is usually sufficient, for each column to be edited. Examples:
“translate to swedish”
“generate 20 descriptions for fictional support tickets.”
(For these kinds of prompts, note that rows must already exist, or be staged to be inserted via Ctrl+Shift+Plus, before they can be filled in.)
Ad copy brainstorm:
“come up with some extra headlines in each category, for the currently blank rows”
In the screenshot above, the Insert Record ( ) action was used to stage inserts before invoking Fill with AI .
To extract the street name from an adjacent “Address” column:
Street Name : “Extract the street name from the address.”
To populate a true/false “engagement” field in a database of sales leads and interactions:
Engagement : “Set this field to true if this timeline entry involved the person being engaged from their side. For example, an email from the person, or a meeting with the person, qualifies as 'true'. But an email that I am sending myself does not, unless I noted that the person responded. Do not count unsubscribes as engagements. For entries not set to true, set the value to false.”
Qualification of sales leads:
Lead score : “Assign a score from 1 to 10 to each sales lead based on past history and other available data. A higher score means prioritize for current outreach efforts.”
Rationale : “A rationale for the lead score given.”
To lightly anonymize a table of users, authors, or such:
First Name : “A random (but always different) first name of the same gender and similar number of syllables, with the first letter staying the same”
Last Name: “A random (but always different) last name with a similar number of syllables, with the first letter staying the same”
Examples where the LLM must maintain context across multiple rows:
“Terse notes from an art reviewer who gets increasingly skeptical of the artworks' merits as the list goes on. By the end of the list, the reviewer is completely exasperated by the lack of quality.”
“a list of numbers from 1 to 20”
For the latter kind of prompts, make sure the amount of data in the perspective is not so large as to cause batching , since this would separate the rows in the context.
Ultorg will supply the remaining parts of the prompt that are needed to make the LLM understand the overall column editing task.
Configuring an AI Provider
The Ultorg desktop application connects directly to an AI provider of your choice, rather than going through an intermediate server. The Configure AI Provider dialog lets you configure a provider of your choice:
Select the provider type from the dropdown, and enter an API key. Press Test Connection to validate the configuration.
The currently supported providers are:
Anthropic. An API key can be created here . Paid service.
OpenAI. An API key can be created here . Paid service.
Other providers compatible with the OpenAI API.
For such providers, select OpenAI and change API Base URL to the appropriate setting.
Models must support streaming and structured JSON output.
The API key will be stored in your operating system's user directory.
Semantics
The data sent to the LLM includes the following:
General instructions about the automated editing task, hardcoded in Ultorg.
The data visible in your perspective as of the time you click OK in the Fill with AI dialog, incorporating any existing staged data edits. If batching is used, only a subset of root-level records will be submitted at a time.
Field names as displayed in Ultorg. The technical table column names that exist at the data source level are intentionally omitted.
For each column to be edited:
The name of the column as displayed in Ultorg.
The data type of the column, and instructions for providing values of that data type.
Whether or not the column supports null values, and how to emit them.
The instructions you supply for the column.
Any additional instructions you optionally provide in the Fill with AI dialog.
The LLM is not obliged to edit every row. It can only, however, edit rows that were visible in the perspective as of the time you pressed OK . You can scroll down to load more rows. In any case, only the selected column(s) will be edited.
You can keep working with the perspective while the operation runs. New data edits, specifically, will interrupt the operation.
Text fields are automatically ordered before other fields in the output format that is requested from the LLM. This allows the reasoning in “rationale”-type fields to be included in the recent context when the LLM subsequently generates, for instance, numerical scores.
Batching
LLMs are limited in the amount of data they can process at any one time, and perform best when this context is not too close to the absolute limit.
For the Fill with AI feature, perspective data in excess of 100 KB (~30K “tokens”) will be automatically split into smaller batches. The dialog box will notify you accordingly:
Batching happens transparently, and you will not normally need to worry about it. However, it means that the LLM may not be aware of every row in your perspective at any given time.
Bulk Editing
Editing Data
Editing Options
contact@ultorg.com
Twitter / : @Ultorg
© 2026 Ultorg Inc.
