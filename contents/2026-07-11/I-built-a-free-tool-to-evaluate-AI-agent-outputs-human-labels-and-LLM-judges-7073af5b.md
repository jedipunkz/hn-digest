---
source: "https://github.com/AntoineF23/verdict"
hn_url: "https://news.ycombinator.com/item?id=48875259"
title: "I built a free tool to evaluate AI agent outputs (human labels and LLM judges)"
article_title: "GitHub - AntoineF23/verdict: Open source tool to evaluate AI features (agents, chatbots, RAG). Review LLM traces, run error analysis (open and axial coding), and validate an LLM as a judge against human labels with a confusion matrix, TPR/TNR and Cohen's kappa. Inspired by Hamel Husain's AI evals. L\n[truncated]"
author: "antoinefornas"
captured_at: "2026-07-11T19:59:32Z"
capture_tool: "hn-digest"
hn_id: 48875259
score: 1
comments: 0
posted_at: "2026-07-11T19:55:44Z"
tags:
  - hacker-news
  - translated
---

# I built a free tool to evaluate AI agent outputs (human labels and LLM judges)

- HN: [48875259](https://news.ycombinator.com/item?id=48875259)
- Source: [github.com](https://github.com/AntoineF23/verdict)
- Score: 1
- Comments: 0
- Posted: 2026-07-11T19:55:44Z

## Translation

タイトル: AI エージェントの出力 (人間のラベルと LLM ジャッジ) を評価するための無料ツールを構築しました
記事タイトル: GitHub - AntoineF23/評決: AI 機能 (エージェント、チャットボット、RAG) を評価するためのオープンソース ツール。 LLM トレースをレビューし、エラー分析 (オープンおよび軸コーディング) を実行し、混同行列、TPR/TNR、およびコーエンのカッパを使用して、人間のラベルに対する判断者として LLM を検証します。 Hamel Husain の AI 評価からインスピレーションを受けています。 L
[切り捨てられた]
説明: AI 機能 (エージェント、チャットボット、RAG) を評価するためのオープンソース ツール。 LLM トレースをレビューし、エラー分析 (オープンおよび軸コーディング) を実行し、混同行列、TPR/TNR、およびコーエンのカッパを使用して、人間のラベルに対する判断者として LLM を検証します。 Hamel Husain の AI 評価からインスピレーションを受けています。ローカルでもオフラインでも。 - アントワーヌF2
[切り捨てられた]

記事本文:
GitHub - AntoineF23/verdict: AI 機能 (エージェント、チャットボット、RAG) を評価するためのオープンソース ツール。 LLM トレースをレビューし、エラー分析 (オープンおよび軸コーディング) を実行し、混同行列、TPR/TNR、およびコーエンのカッパを使用して、人間のラベルに対する判断者として LLM を検証します。 Hamel Husain の AI 評価からインスピレーションを受けています。ローカルでもオフラインでも。 · GitHub
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
アカウントを切り替えました

別のタブまたはウィンドウで。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
アントワーヌF23
/
評決
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
4 コミット 4 コミット .github .github fixtures fixtures scripts scripts src src test-files test-files test test .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md Index.html Index.html package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json vite.config.ts vite.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Verdict は、AI 機能を評価するためのローカルのオープンソース ツールです。任意の AI からトレースをロードします
エージェント、チャットボット、RAG アシスタント、または LLM 機能。各会話を確認し、合格またはマークを付けます。
コメントで失敗します。根拠のある理論の誤りによってどのような種類の障害が発生するかを発見する
分析。次に、人間のラベルに対する判断者として LLM を構築して検証し、それをエクスポートします。
生産時の品質を監視するために判断します。
すべてがブラウザ内で実行されます。そうでない限り、バックエンドもアカウントもデータもマシンから出ません。
LLM API を呼び出すことを選択します。オフラインで動作する単一の Web ページとして配布します。
ライブデモ: https://antoinef23.github.io/verdict/
これが存在する理由: AI 機能をリリースし、いくつかの出力を目視する (「雰囲気チェック」) ことは、
それが良いかどうか、どのように失敗するか、または改善しているかどうかを教えてください。実際の評価はループです。あなた
データを見てラベルを付け、故障モードを見つけて測定し、測定を自動化します。
あなたが信頼できると証明された裁判官に対して。 Verdict は、そのループを実行するための専用ツールです。
このツールは、Hamel Husain の AI 評価手法に大きく影響を受けています。核となるアイデア
来てください

彼の教えから直接: データを見て、オープンおよびアキシャル コーディングでエラー分析を実行し、
そして、人間の審査員と一致することが示された場合にのみ、LLM 審査員を信頼します (
真陽性率と真陰性率)。評決の背後にある理論が知りたい場合は、以下をご覧ください。
50 分でわかる AI Evals クラッシュ コース (Hamel Husain):
https://creatoreconomy.so/p/ai-evaluations-crash-course-in-50- minutes-hamel-husain
Verdict は、これらのアイデアを独立して実装したものです。と提携または承認されていません。
ハメル・フサイン。
見て。 AI 機能から実際のトレースを読み取ります。
ラベル。各会話に合格または不合格のマークを付け、短いコメントを書きます。
コード。短いフリー テキスト コード (オープン コーディング) を使用して、それぞれの失敗が発生した理由をタグ付けします。
クラスタ。それらのコードを故障カテゴリにグループ化します (軸方向コーディング)。
裁判官。失敗カテゴリごとに 1 人の LLM ジャッジを作成します。
検証します。裁判官が人間の意見に同意していることを証明してください (混同行列、TPR、TNR、コーエンのカッパ)。
船。信頼できるジャッジをエクスポートし、人間が読む時間がない運用トラフィックで実行します。
ステップ 1 ～ 4 は人間の作業ですが、Verdict を使用すると迅速に実行できます。ステップ 5 から 7 により、その判断を調整できます
人間が読むことのない会話は、審査員がそれが査読者と一致することを証明した後にのみ読まれます。
あらゆるものからトレースを読み取ります。フォーマットに依存しない取り込み: OpenTelemetry (OTLP) エクスポート、
OpenInference と gen_ai と OpenLLMetry スパン、Langfuse エクスポート、プレーン [{role, content}]
会話、NDJSON、および他の JSON を段階的なタイムラインとしてレンダリングするフォールバック。
生の JSON ではなく、会話をレンダリングします。ユーザーとアシスタントのメッセージを含むクリーンなチャット タイムライン
およびツール呼び出し (折りたたみ可能な入力と出力) に加えて、すべてのステップで「生の表示」トグルが追加されます。
人間による迅速なレビュー。進行状況とフィルター、合格または失敗とコメントを含むキュー、および
何百もの C を移動するために構築されたキーボード ショートカット

会話。ブラウザに自動保存されます。
根拠のある理論によるエラー分析。フリーテキストのオープンコードを失敗に追加し、LLM を使用します
これらをカテゴリごとの数を使用して障害分類にクラスタリングします。
LLM が審査員として認定されました。失敗カテゴリごとに 1 つのバイナリ ジャッジがあり、それぞれに調整可能なパラメータがあり、
バージョン管理されたプロンプトと選択したモデル。トレーニングとテストを使用して人間のラベルに対して検証する
分割と完全混同行列 (TPR、TNR、精度、F1、精度、コーエンのカッパ)。をエクスポートします。
本番環境で実行するプロンプト、モデル、メトリクスを判断してください。
匿名化。 PII (電子メール、電話、IP、URL、カード、IBAN、および名前と会社) を編集します
ヒューリスティックおよび独自の用語リスト) をタイムライン、エクスポート、および LLM に送信されるすべてのものに含めます。
ローカルかつプライベート。すべてがブラウザ内で実行されます。フィードバックはローカル ストレージに保存され、
CSV または JSON にエクスポートします。唯一のネットワーク呼び出しは、トリガーするオプションの LLM API リクエストです。
インストールせずに使用するだけです。上記のライブデモを開くか、構築された dist/index.html をダウンロードして開きます。
どのブラウザでも使えます。これは 1 つのファイルであり、オフラインで動作します。
npmインストール
npm run dev # ホットリロードを使用した開発サーバー
npm run build # 1 つの自己完結型ファイルである dist/index.html を生成します
npm test # 単体テストと DOM テスト
npm タイプチェックを実行する
同梱されているサンプル データを使用して、すぐに試してみてください (「サンプル データ」を参照)。
LLM審査員はどのように評価されるのか（重要な部分）
裁判官の要点は、人間が読まないデータについて人間の審査員に取って代わることです。それだけです
裁判官のラベルが人間とほぼ同じであることを証明できれば安全です。評決はこれを他の方法で行います
分類子が評価されます。人間のラベルがグラウンド トゥルースとして扱われ、ジャッジのラベルが同じになります。
会話を行い、両者を比較します。
このプラットフォーム上のすべての会話は人間によってラベル付けされます。特定の故障カテゴリ X に対して、
会話は肯定的なものです

レビュー担当者のコードはカテゴリ X にマップされ、否定的なコードは
それ以外の場合は。 「完全にコード化された」チェックボックスを使用すると、失敗を徹底的にコード化されたものとしてマークできます。
他のカテゴリの否定的な情報は信頼できます。
1 つのカテゴリについて、各会話の人間のラベルと審査員のラベルを比較します。すべての会話
4 つのセルのいずれかに該当します。
TP (真陽性): 両方とも障害が存在することに同意します。
TN (真陰性): 両方ともそれが存在しないことに同意します。
FP (偽陽性): 裁判官がフラグを立てましたが、人間はフラグを立てませんでした。旗をめぐる裁判官。
FN (偽陰性): 裁判官は人間が捕まえたものを見逃しました。旗の下の裁判官。
TPR (再現率、感度) = TP / (TP + FN)
TNR (特異度) = TN / (TN + FP)
精度 = TP / (TP + FP)
F1 = 2 * 精度 * TPR / (精度 + TPR)
精度 = (TP + TN) / (TP + FP + FN + TN)
TPR は、裁判官が発見した実際の失敗の割合です。 TPR が低いということは、障害を見逃すことを意味します。
TNR は、裁判官が正しく放置したクリーンな会話の割合です。 TNRが低いということは、
誤った警報を発します。
精度とは、裁判官が「YES」と答えたときにどれだけ正しいかということです。
F1 は、精度と再現率のバランスをとる単一の数値です。
精度は全体的な正しいシェアです。失敗がまれである場合、これは誤解を招きます。
常に「NO」と言う裁判官でも高得点を獲得できる可能性があります。精度だけに頼らないでください。
コーエンのカッパ (見出し番号)
特に 1 つのクラスがまれな場合、2 人のラベラーが偶然一致することがよくあります。コーエンのカッパ
偶然の一致を修正するので、最高のシングル「ジャッジは十分ですか」ナンバーです。
po (観察された一致) = 精度
pe (偶然の一致) = ( (TP + FN) * (TP + FP) + (FN + TN) * (FP + TN) ) / N^2
カッパ = (ポ - ペ) / (1 - ペ)
おおよその目安: 0.4 未満は弱い、0.4 ～ 0.6 は中程度、0.6 ～ 0.8 は充実、0.8 以上は
ほぼ完璧。高いTPRとともに高いカッパを目指します。
チューニングすると、

自分が得点したのと同じ会話を見ながらプロンプトを判断すると、オーバーフィットし、
数字は嘘をつきます。評決は、ラベル付きセットをトレイン部分とホールドアウトテスト部分に分割します
（階層化されているため、両方とも同じポジティブバランスを維持します）。電車でチューニングして、テストを信頼してください
数字。 Verdict は、プロンプトまたはモデルを編集するたびに、両方のバージョンのメトリックをレポートし、クリアします。
したがって、古い数字が誤解を招くことはありません。
裁判官が差し出されたテストセット（高いカッパ、高いTPR、許容範囲）について強い合意に達したとき
FP)、それはそのカテゴリーにおける人間の信頼できる代役です。エクスポートしてラベルなしで実行します
実稼働トラフィック。その基準に達しない場合は、プロンプトまたはモデルを改善し続けるか、
そのカテゴリーのループにいる人間。
実際の評価を実行するために技術的な知識は必要ありません。このチュートリアルでは、バンドルされている 100 を使用します。
会話デモなので、ループ全体を約 20 分で実行し、その後は自分で繰り返すことができます
特徴の痕跡。
0. トレースを取得します。 AI 機能 (エージェント、チャットボット、RAG アシスタント、分類子) が生成する
会話。通常、エンジニアはこれらを可観測性スタックから JSON としてエクスポートできます。
(OpenTelemetry、Langfuse、Arize Phoenix、LangSmith、または単純なログ)。評決はこれらすべてを読み上げます。いいえ
輸出はまだですか？ test-file/ のデモ ファイルを使用します。
1. データを確認します。これが最も重要なステップです。 [Verdict] を開き、 [エクスポートの読み込み] をクリックして、
[demo-100-conversations.json] を選択します。いくつかの会話を端から端まで読んでください。あなたは気づき始めるでしょう
エージェントは物事をでっち上げ、指示を無視し、ツールを誤って扱います。
2. 合格または不合格のラベルとコメントを付けます。会話ごとに、AI がその役割を果たしたかどうかを判断します。ヒット
合否を記入し、その理由について短いコメントを書きます。バイナリのままにしておきます。 「に発送するのに十分です」
ユーザー？」 1対5のスコアを上回り、意見の相違は隠蔽されている。ショートカット: j

k で移動、p と f
採点するには / コメントにジャンプします。
3. 失敗をオープンコード化します。キューフィルターを Fail に設定します。失敗するたびに、1 つ以上を追加します
フィードバック バーの短いエラー コード (幻覚図、無視された制約など)
ツールエラーは無視されました。ぴったりの言葉を考え出します。以前に使用したコードがオートコンプリートされるため、文言が
一貫性を保ちます。会話内のすべての問題を把握したら、完全にコード化されたものにチェックを入れます。
4. 故障分類法への軸コード。 [分析が失敗しました] をクリックします。すべての個別のコードが表示されます
そしてその頻度。それらをより高いレベルの障害カテゴリに分類します。 API キーを使用して、「実行」をクリックします。
API を使用して。キーを使用しない場合は、[プロンプトのコピー] をクリックし、ChatGPT または Claude に貼り付けて、
JSONが戻ってきました。評決ではカテゴリがカウントとともに表示されます。カテゴリのフィルターキューをクリックして確認します
それらの失敗だけ。
5. カテゴリごとに審査員を構築し、検証します。 [審査員] をクリックします。失敗ごとに 1 つのタブがあります
カテゴリ。モデルを選択し、判定プロンプトを確認または編集して、「ラベル付きセットで検証」をクリックします。
混同行列と指標を読んでください (パネルに「これらの数値の読み方」ガイドがあります)。
見出しはテストスプリットに関するコーエンのカッパです。 FP が高いということはフラグを超えていることを意味するため、プロンプトを厳しくします。

[切り捨てられた]

## Original Extract

Open source tool to evaluate AI features (agents, chatbots, RAG). Review LLM traces, run error analysis (open and axial coding), and validate an LLM as a judge against human labels with a confusion matrix, TPR/TNR and Cohen's kappa. Inspired by Hamel Husain's AI evals. Local and offline. - AntoineF2
[truncated]

GitHub - AntoineF23/verdict: Open source tool to evaluate AI features (agents, chatbots, RAG). Review LLM traces, run error analysis (open and axial coding), and validate an LLM as a judge against human labels with a confusion matrix, TPR/TNR and Cohen's kappa. Inspired by Hamel Husain's AI evals. Local and offline. · GitHub
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
AntoineF23
/
verdict
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits .github .github fixtures fixtures scripts scripts src src test-files test-files test test .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md index.html index.html package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json vite.config.ts vite.config.ts View all files Repository files navigation
Verdict is a local, open source tool for evaluating AI features. Load the traces from any AI
agent, chatbot, RAG assistant, or LLM feature. Review each conversation and mark it Pass or
Fail with a comment. Discover what kinds of failures happen with grounded theory error
analysis. Then build and validate an LLM as a judge against your human labels, and export that
judge to monitor quality in production.
Everything runs in your browser. No backend, no accounts, and no data leaves your machine unless
you choose to call an LLM API. Ship it as a single web page that works offline.
Live demo: https://antoinef23.github.io/verdict/
Why this exists: shipping an AI feature and eyeballing a few outputs (the "vibe check") does not
tell you if it is good, how it fails, or whether it is improving. Real evaluation is a loop. You
look at your data, label it, find the failure modes, measure them, then automate the measurement
with a judge you have proven trustworthy. Verdict is a purpose built tool for running that loop.
This tool is heavily inspired by the AI evaluation methodology of Hamel Husain . The core ideas
come straight from his teaching: look at your data, run error analysis with open and axial coding,
and only trust an LLM judge once it has been shown to agree with human reviewers (measured with
true positive and true negative rates). If you want the theory behind Verdict, watch:
AI Evals Crash Course in 50 minutes (Hamel Husain):
https://creatoreconomy.so/p/ai-evaluations-crash-course-in-50-minutes-hamel-husain
Verdict is an independent implementation of these ideas. It is not affiliated with or endorsed by
Hamel Husain.
Look. Read real traces from your AI feature.
Label. Mark each conversation Pass or Fail and write a short comment.
Code. Tag why each failure happened with short free text codes (open coding).
Cluster. Group those codes into failure categories (axial coding).
Judge. Build one LLM judge per failure category.
Validate. Prove the judge agrees with humans (confusion matrix, TPR, TNR, Cohen's kappa).
Ship. Export the trusted judge and run it on production traffic no human has time to read.
Steps 1 to 4 are human work, and Verdict makes them fast. Steps 5 to 7 let you scale that judgment
to conversations no human will read, but only after the judge has proven it matches your reviewers.
Reads traces from anything. Format agnostic ingest: OpenTelemetry (OTLP) exports,
OpenInference and gen_ai and OpenLLMetry spans, Langfuse exports, plain [{role, content}]
conversations, NDJSON, and a fallback that renders any other JSON as a gradable timeline.
Renders conversations, not raw JSON. A clean chat timeline with user and assistant messages
and tool calls (collapsible inputs and outputs), plus a "show raw" toggle on every step.
Fast human review. A queue with progress and filters, Pass or Fail plus a comment, and
keyboard shortcuts built for moving through hundreds of conversations. Autosaves to your browser.
Grounded theory error analysis. Add free text open codes to failures, then have an LLM
cluster them into a failure taxonomy with per category counts.
LLM as a judge, validated. One binary judge per failure category, each with a tunable,
versioned prompt and a model you choose. Validate against your human labels with a train and test
split and a full confusion matrix (TPR, TNR, precision, F1, accuracy, Cohen's kappa). Export the
judge with its prompt, model, and metrics to run in production.
Anonymization. Redact PII (emails, phones, IPs, URLs, cards, IBANs, plus a name and company
heuristic and your own term list) in the timeline, in exports, and in everything sent to an LLM.
Local and private. Everything runs in the browser. Feedback lives in local storage and
exports to CSV or JSON. The only network calls are the optional LLM API requests you trigger.
Just use it, no install: open the live demo above, or download the built dist/index.html and open
it in any browser. It is one file and works offline.
npm install
npm run dev # dev server with hot reload
npm run build # produces dist/index.html, one self contained file
npm test # unit and DOM tests
npm run typecheck
Try it right away with the bundled sample data (see Sample data ).
How the LLM judge is evaluated (the important part)
The whole point of a judge is to replace a human reviewer on data no human will read. That is only
safe if you can prove the judge labels almost as well as a human . Verdict does this the way any
classifier is evaluated: it treats your human labels as ground truth, has the judge label the same
conversations, and compares the two.
Every conversation on this platform is labeled by a human. For a given failure category X, a
conversation is a positive if the reviewer's codes map to category X, and a negative
otherwise. The "fully coded" checkbox lets you mark a failure as exhaustively coded, so its
negatives for other categories are trustworthy.
For one category, compare each conversation's human label to the judge's label. Every conversation
falls into one of four cells:
TP (true positive): both agree the failure is present.
TN (true negative): both agree it is absent.
FP (false positive): the judge flagged it, the human did not. The judge over flags.
FN (false negative): the judge missed one the human caught. The judge under flags.
TPR (recall, sensitivity) = TP / (TP + FN)
TNR (specificity) = TN / (TN + FP)
Precision = TP / (TP + FP)
F1 = 2 * Precision * TPR / (Precision + TPR)
Accuracy = (TP + TN) / (TP + FP + FN + TN)
TPR is the share of real failures the judge catches. Low TPR means it misses failures.
TNR is the share of clean conversations the judge correctly leaves alone. Low TNR means it
raises false alarms.
Precision is how often the judge is right when it says YES.
F1 is a single number balancing precision and recall.
Accuracy is the overall share correct. It is misleading when a failure is rare, because a
judge that always says NO can still score high. Do not rely on accuracy alone.
Cohen's kappa (the headline number)
Two labelers can agree a lot just by chance, especially when one class is rare. Cohen's kappa
corrects agreement for chance, so it is the best single "is the judge good enough" number.
po (observed agreement) = Accuracy
pe (chance agreement) = ( (TP + FN) * (TP + FP) + (FN + TN) * (FP + TN) ) / N^2
kappa = (po - pe) / (1 - pe)
Rough reading: below 0.4 is weak, 0.4 to 0.6 is moderate, 0.6 to 0.8 is substantial, above 0.8 is
near perfect. Aim for high kappa together with high TPR.
If you tune the judge prompt while looking at the same conversations you score on, you overfit and
the numbers lie. Verdict splits your labeled set into a train part and a held out test part
(stratified so both keep the same balance of positives). Tune on train, and trust the test
numbers. Verdict reports both and clears a version's metrics whenever you edit its prompt or model,
so stale numbers never mislead you.
When a judge reaches strong agreement on the held out test set (high kappa, high TPR, acceptable
FP), it is a trustworthy stand in for a human on that category. Export it and run it on unlabeled
production traffic. If it does not reach that bar, keep improving the prompt or model, or keep a
human in the loop for that category.
You do not need to be technical to run a real evaluation. This walkthrough uses the bundled 100
conversation demo, so you can do the whole loop in about 20 minutes, then repeat it on your own
feature's traces.
0. Get some traces. An AI feature (agent, chatbot, RAG assistant, classifier) produces
conversations. Your engineers can usually export these as JSON from your observability stack
(OpenTelemetry, Langfuse, Arize Phoenix, LangSmith, or a plain log). Verdict reads all of these. No
export yet? Use the demo files in test-files/ .
1. Look at your data. This is the most important step. Open Verdict, click Load export , and
pick demo-100-conversations.json . Read a few conversations end to end. You will start noticing
patterns immediately: the agent makes things up, ignores instructions, mishandles tools.
2. Label Pass or Fail plus a comment. For each conversation decide: did the AI do its job? Hit
Pass or Fail and write a short comment on why. Keep it binary. "Good enough to ship to a
user?" beats a 1 to 5 score, which hides disagreement. Shortcuts: j and k to move, p and f
to grade, / to jump to the comment.
3. Open code the failures. Set the queue filter to Fail . On each failure, add one or more
short error codes in the feedback bar, for example hallucinated figure , ignored constraint ,
ignored tool error . Invent the words that fit. Previously used codes autocomplete so wording
stays consistent. Tick fully coded once you have captured every problem in that conversation.
4. Axial code into a failure taxonomy. Click Analyze fails . You will see every distinct code
and its frequency. Cluster them into higher level failure categories. With an API key, click Run
with API . Without a key, click Copy prompt , paste it into ChatGPT or Claude, and paste the
JSON back. Verdict shows the categories with counts. Click filter queue on a category to review
just those failures.
5. Build and validate a judge per category. Click Judges . There is one tab per failure
category. Pick the model, review or edit the judge prompt, and click Validate on labeled set .
Read the confusion matrix and metrics (there is a "How to read these numbers" guide in the panel).
The headline is Cohen's kappa on the test split. High FP means it over flags, so tighten the prompt.

[truncated]
