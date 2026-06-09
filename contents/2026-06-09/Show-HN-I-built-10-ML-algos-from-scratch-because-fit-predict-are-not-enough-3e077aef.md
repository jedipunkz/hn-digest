---
source: "https://github.com/ml-from-scratch-book/code"
hn_url: "https://news.ycombinator.com/item?id=48461956"
title: "Show HN: I built 10 ML algos from scratch because fit() predict() are not enough"
article_title: "GitHub - ml-from-scratch-book/code: Companion code for Machine Learning From Scratch — 10 core ML algorithms built from scratch with NumPy, compared with Scikit-learn and PyTorch. · GitHub"
author: "akmoleksandr"
captured_at: "2026-06-09T16:05:03Z"
capture_tool: "hn-digest"
hn_id: 48461956
score: 4
comments: 1
posted_at: "2026-06-09T14:53:46Z"
tags:
  - hacker-news
  - translated
---

# Show HN: I built 10 ML algos from scratch because fit() predict() are not enough

- HN: [48461956](https://news.ycombinator.com/item?id=48461956)
- Source: [github.com](https://github.com/ml-from-scratch-book/code)
- Score: 4
- Comments: 1
- Posted: 2026-06-09T14:53:46Z

## Translation

タイトル: Show HN: fit()、predict() だけでは不十分なため、10 個の ML アルゴを最初から構築しました
記事のタイトル: GitHub - ml-from-scratch-book/code: ゼロからの機械学習のコンパニオン コード — NumPy を使用してゼロから構築された 10 個のコア ML アルゴリズムを、Scikit-learn および PyTorch と比較。 · GitHub
説明: ゼロからの機械学習のコンパニオン コード — NumPy を使用してゼロから構築された 10 個のコア ML アルゴリズムを、Scikit-learn や PyTorch と比較しました。 - スクラッチブック/コードからの ml

記事本文:
GitHub - ml-from-scratch-book/code: スクラッチからの機械学習のコンパニオン コード — NumPy を使用してゼロから構築された 10 個のコア ML アルゴリズムを、Scikit-learn や PyTorch と比較しました。 · GitHub
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
スクラッチブックからのml
/
コード
公共
通知

通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
10 コミット 10 コミット 01_linear_regression.ipynb 01_linear_regression.ipynb 02_logistic_regression.ipynb 02_logistic_regression.ipynb 03_正規化.ipynb 03_正規化.ipynb 04_k_nearest_neighbors.ipynb 04_k_nearest_neighbors.ipynb 05_naive_bayes.ipynb 05_naive_bayes.ipynb 06_決定_ツリー.ipynb 06_決定_ツリー.ipynb 07_ランダム_フォレスト.ipynb 07_ランダム_フォレスト.ipynb 08_勾配_ブースティング.ipynb 08_gradient_boosting.ipynb 09_xgboost.ipynb 09_xgboost.ipynb 10_neural_network.ipynb 10_neural_network.ipynb 11_model_optimization.ipynb 11_model_optimization.ipynb ライセンス ライセンス README.md README.md data_loader.py data_loader.py 要件.txt 要件.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Machine Learning From Scratch のコンパニオン コード リポジトリ。
Machine Learning From Scratch は、 fit() と detect() の背後にある「ブラック ボックス」を開きます。基本的な Python と高校数学を理解している人なら誰でも、5 段階のフレームワークを通じてコア ML アルゴリズムを完全に理解し、ゼロから構築できます。
直感 : アルゴリズムの背後にある平易な人間の思考を理解します。
形式化 : その直感をアクセス可能な数学的形式に変換します。
実装 : NumPy を使用して、数学を最初からクリーンなコードに組み込みます。
テスト : 実際のデータでコードを検証し、業界標準のライブラリと比較します。
ヒント : 実際的な長所、短所、および実際の使用法に関する洞察を学びます。
章
コンテンツ/コード
セットアップ
概要
ML の基礎
概要
データの概要
概要
ML に実際に必要な数学
概要
データの準備
概要
データローダー.py
線形回帰
01_linear_regression.ipynb
ロジスティック回帰
02_ロジスティック

_回帰.ipynb
正則化
03_正規化.ipynb
K最近傍法
04_k_nearest_neighbors.ipynb
ナイーブ・ベイズ
05_naive_bayes.ipynb
デシジョンツリー
06_意思決定ツリー.ipynb
ランダムフォレスト
07_ランダムフォレスト.ipynb
勾配ブースティング
08_gradient_boosting.ipynb
極端な勾配ブースティング (XGBoost)
09_xgboost.ipynb
ニューラルネットワーク
10_ニューラルネットワーク.ipynb
モデルを最大限に活用する
概要
11_モデル_最適化.ipynb
結論
概要
各章の概要
コードをローカルで実行するには、リポジトリのクローンを作成し、依存関係をインストールします。
git clone https://github.com/ml-from-scratch-book/code.git
CDコード
pip install -r 要件.txt
あるいは、ローカル設定を行わずに、すべてのノートブックを Google Colab で直接開くこともできます。colab.research.google.com にアクセスして、GitHub からインポートします。
機械学習の背後にある中心的な考え方と定義、つまりアルゴリズムとは何か、モデルとは何か、そしてそのすべてにおいてデータがどのような役割を果たすかを紹介します。関連性のあるアナロジーを使用して、ML 実践者が何を達成しようとしているのか、そしてどのようにそれに取り組むのかについての明確な頭の中でのイメージを構築します。
特徴、ターゲット、データ型、欠損値やクラスの不均衡などの一般的な問題など、データが実際にどのように見えるかを調査します。おもちゃのデータセットを超えて、ルールベースのアプローチよりも ML が実際に適切なツールであることを探ります。
ML に実際に必要な数学
ほとんどの ML アルゴリズムの基礎となる 3 つの数学的柱を効果的に詳しく説明します。
データの操作 (線形代数の基礎) : データを保存する方法としてのベクトル、行列、データを操作する方法としての行列の乗算およびその他の一般的な演算。
データを理解する (統計の基礎) : 分布、平均/分散、確率、サンプル サイズの重要性を実際の例を使って説明します。
データから学ぶ（最適化）：関数の概念からスタートし、

2 点間の限界までの距離、目的関数、勾配降下法、誤差を最小限に抑えることでモデルが実際に「学習」する方法。
トレーニング前に行われる実践的な作業 (欠損データの処理、カテゴリ変数のエンコード、特徴量のスケーリング、トレーニング/テストの分割、データ漏洩の回避) について説明します。この段階では、アルゴリズムで活用するデータの読み込み、前処理、分割を行うスクリプトを準備します。
この本の核心 — 10 個のコア ML アルゴリズムは、それぞれ 5 段階のフレームワークを使用してゼロから構築されています。
線形回帰 · ロジスティック回帰 · 正則化 · K 最近傍法 · ナイーブ ベイズ · デシジョン ツリー · ランダム フォレスト · 勾配ブースティング · XGBoost · ニューラル ネットワーク
ここで読者は、問題にアプローチする方法と、該当する場合は ML を使用して問題を解決する方法を学びます。言い換えれば、本書で得たすべての知識を、次の 3 つの段階で構成される ML プロジェクト開発の反復フレームワークにどのように組み込むかということです。
データ : 明確な問題ステートメント、データ品質、特徴エンジニアリング、データ漏洩防止の重要性の考えを強化します。
モデリング: 検証セット、k 分割相互検証、モデル選択、およびグリッド、ランダム、ベイジアン検索による Optuna によるハイパーパラメーター調整。
評価 : 回帰、バイナリおよびマルチクラス分類問題のメトリクスが、なぜ重要なのかの例とともに詳細に説明されています。
読者は一歩下がって、この本の後に自分がどうなったのか、そしてここから進むことができるさまざまな方向性について、より大きな全体像を見てみましょう。最も重要なことは、彼らがここで得た考え方と実践的なスキルは、実際にたどる道に関係なく大いに役立つということを理解して終了することです。
ゼロからの機械学習のコンパニオン コード — NumPy を使用してゼロから構築された 10 個のコア ML アルゴリズム (Scikit-learn および Scikit-learn と比較)

パイトーチ。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
1
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Companion code for Machine Learning From Scratch — 10 core ML algorithms built from scratch with NumPy, compared with Scikit-learn and PyTorch. - ml-from-scratch-book/code

GitHub - ml-from-scratch-book/code: Companion code for Machine Learning From Scratch — 10 core ML algorithms built from scratch with NumPy, compared with Scikit-learn and PyTorch. · GitHub
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
ml-from-scratch-book
/
code
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
10 Commits 10 Commits 01_linear_regression.ipynb 01_linear_regression.ipynb 02_logistic_regression.ipynb 02_logistic_regression.ipynb 03_regularization.ipynb 03_regularization.ipynb 04_k_nearest_neighbors.ipynb 04_k_nearest_neighbors.ipynb 05_naive_bayes.ipynb 05_naive_bayes.ipynb 06_decision_tree.ipynb 06_decision_tree.ipynb 07_random_forest.ipynb 07_random_forest.ipynb 08_gradient_boosting.ipynb 08_gradient_boosting.ipynb 09_xgboost.ipynb 09_xgboost.ipynb 10_neural_network.ipynb 10_neural_network.ipynb 11_model_optimization.ipynb 11_model_optimization.ipynb LICENSE LICENSE README.md README.md data_loader.py data_loader.py requirements.txt requirements.txt View all files Repository files navigation
Companion code repository for Machine Learning From Scratch .
Machine Learning From Scratch opens the "black box" behind fit() and predict() . Anyone with basic Python and high-school math can fully understand and build core ML algorithms from scratch via a 5-stage framework:
Intuition : Understand the plain-English human thinking behind the algorithm.
Formalization : Translate that intuition into accessible mathematical form.
Implementation : Put the math into clean, from-scratch code using NumPy.
Test : Validate your code on real data and compare it against industry-standard libraries.
Tips : Learn practical strengths, weaknesses, and real-world usage insights.
Chapter
Content/Code
Setup
Overview
Fundamentals of ML
Overview
Introduction to Data
Overview
The Math You Actually Need for ML
Overview
Data Preparation
Overview
data_loader.py
Linear Regression
01_linear_regression.ipynb
Logistic Regression
02_logistic_regression.ipynb
Regularization
03_regularization.ipynb
K-Nearest Neighbors
04_k_nearest_neighbors.ipynb
Naïve Bayes
05_naive_bayes.ipynb
Decision Tree
06_decision_tree.ipynb
Random Forest
07_random_forest.ipynb
Gradient Boosting
08_gradient_boosting.ipynb
Extreme Gradient Boosting (XGBoost)
09_xgboost.ipynb
Neural Network
10_neural_network.ipynb
Making the Best out of Models
Overview
11_model_optimization.ipynb
Conclusion
Overview
Chapter Overviews
To run the code locally, clone the repo and install the dependencies:
git clone https://github.com/ml-from-scratch-book/code.git
cd code
pip install -r requirements.txt
Alternatively, every notebook can be opened directly in Google Colab with no local setup — go to colab.research.google.com and import from GitHub.
Introduces the core ideas and definitions behind machine learning — what an algorithm is, what a model is, and what role data plays in all of it. Uses relatable analogies to build a clear mental picture of what ML practitioners aim to accomplish and how they go about it.
Explores what data actually looks like in practice: features, targets, data types, and common issues like missing values and class imbalance. Goes beyond toy datasets to explore when ML is actually the right tool over a rule-based approach.
The Math You Actually Need for ML
An efficient dive into three mathematical pillars underlying most ML algorithms:
Manipulating Data (Linear Algebra Essentials) : Vectors, matrices as ways to store data, matrix multiplication and other common operations as ways to manipulate it.
Understanding Data (Statistical Foundations) : Distributions, mean/variance, probability, significance of sample size explained via practical examples.
Learning from Data (Optimization) : Starting from the concept of a function and distance between two points to limits, objective functions, gradient descent, and how models actually "learn" by minimizing error.
Covers the practical work that happens before training: handling missing data, encoding categorical variables, feature scaling, train/test splitting, and avoiding data leakage. At this stage we prepare a script that loads, preprocesses and splits the data to be leveraged with algorithms.
The heart of the book — 10 core ML algorithms each built from scratch using the 5-stage framework:
Linear Regression · Logistic Regression · Regularization · K-Nearest Neighbors · Naïve Bayes · Decision Tree · Random Forest · Gradient Boosting · XGBoost · Neural Network
Here, the reader learns how to approach a problem and if applicable, solve it with ML. In other words, how to put all the knowledge gained in this book into an iterative framework of ML project development comprised of 3 stages:
Data : Reinforces the idea of the importance of a clear problem statement, data quality, feature engineering and preventing data leakage.
Modeling : Validation set, k-fold cross-validation, model selection and hyperparameter tuning with Optuna via Grid, Random and Bayesian search.
Evaluation : Metrics for regression, binary and multiclass classification problems covered in detail with examples of why they matter.
The reader takes a step back to look at the bigger picture of who they've become after this book and of different directions they can go from here. Most importantly they finish with the understanding that the mindset and practical skills they gained here goes a long way regardless of the exact path they follow.
Companion code for Machine Learning From Scratch — 10 core ML algorithms built from scratch with NumPy, compared with Scikit-learn and PyTorch.
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
1
fork
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
