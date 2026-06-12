---
source: "https://mljar.com/blog/web-app-machine-learning/"
hn_url: "https://news.ycombinator.com/item?id=48502867"
title: "Build a Web App for Your Machine Learning Model"
article_title: "Build a Web App for your Machine Learning model"
author: "pplonski86"
captured_at: "2026-06-12T13:18:04Z"
capture_tool: "hn-digest"
hn_id: 48502867
score: 2
comments: 0
posted_at: "2026-06-12T11:43:34Z"
tags:
  - hacker-news
  - translated
---

# Build a Web App for Your Machine Learning Model

- HN: [48502867](https://news.ycombinator.com/item?id=48502867)
- Source: [mljar.com](https://mljar.com/blog/web-app-machine-learning/)
- Score: 2
- Comments: 0
- Posted: 2026-06-12T11:43:34Z

## Translation

タイトル: 機械学習モデル用の Web アプリを構築する
記事のタイトル: 機械学習モデル用の Web アプリを構築する
説明: MLJAR AutoML を使用して機械学習モデルをトレーニングし、Web アプリを自動的に生成し、単一予測およびバッチ予測のためにオンラインで公開する方法を学びます。

記事本文:
機械学習モデル用の Web アプリを構築する 製品 MLJAR Studio
Studio をダウンロード 6 11 2026 · Piotr Płoński 機械学習
機械学習モデル用の Web アプリを構築する
機械学習モデルの準備が整いました。次に何をすべきでしょうか?モデルを保存して Python で使用できます。 APIを作成できます。モデルを既存のアプリケーションに追加できます。機械学習モデルを使用して予測を行う方法は数多くあります。
このチュートリアルでは、初心者向けのオプションとして、機械学習モデル用の Web アプリを構築します。
一番いいところは？ HTML、CSS、または JavaScript を記述する必要はありません。パイソンのみ。
MLJAR AutoML を使用して機械学習モデルをトレーニングし、 app() を使用して Web アプリを自動的に生成し、MLJAR Platform を使用してアプリをオンラインで公開します。生成されたアプリケーションは、Python ノートブックから Web アプリを構築するためのフレームワークである Mercury を使用して作成されます。
最終結果では、2 つの Web アプリが作成されます。
ユーザーがフォームに入力すると 1 つの予測を取得する、単一予測用のアプリ。
バッチ予測用のアプリ。ユーザーは CSV ファイルをアップロードし、多くの行の予測を取得します。
完全なコードと結果は、私の GitHub リポジトリにあります。
https://github.com/pplonski/web-app-from-machine-learning-model
このチュートリアルでは、医療保険料を予測するアプリケーションを構築します。このモデルでは、以下のような患者の人口統計と健康関連の要因が使用されます。
目標値は、医療保険料を意味する 料金 です。数値を予測するため、これは回帰問題です。
MLJAR AutoML を使用してモデルをトレーニングします。 AutoML は、さまざまな機械学習アルゴリズムを自動的にチェックし、パフォーマンスを比較して、最適なモデルを選択します。トレーニング後、トレーニングされたモデルから Web アプリを生成します。

多くの人は Python コードを直接操作したくないため、これは便利です。彼らはシンプルな Web インターフェイスを好みます。たとえば、上司、顧客、医師、研究者、ビジネス チームメイトは、リンクを開いて値を入力し、予測を取得できます。
機械学習モデル用の Web アプリを作成する理由は何ですか?
機械学習モデルのトレーニングはプロジェクトの一部にすぎません。 2 番目の部分は、モデルを有用なものにすることです。モデルが Python スクリプトまたはノートブックに隠されている場合、それを使用できるのは技術ユーザーのみです。 Python を知らない人と共有するのは難しいです。
Web アプリはこの問題を解決します。ユーザーに次のことができる優れたインターフェイスを提供できます。
コードに触れずにモデルを使用します。
これは、機械学習モデルを共有するための非常に実用的な方法です。多くのプロジェクトでは、プロトタイプを表示したり、アイデアをテストしたり、ユーザーからフィードバックを収集したりするには、これで十分です。
pip install mljar 監視付きマーキュリー
mljar 監視パッケージは、機械学習モデルのトレーニングに使用されます。 mercury パッケージは、生成された Web アプリを実行するために使用されます。
リポジトリ内のrequirements.txtファイルには、次の2つのパッケージのみが含まれています。
mljar 監視付き
水銀
プロジェクトの構造
プロジェクト リポジトリには、次のファイルとディレクトリが含まれています。
機械学習モデルからのウェブアプリ/
§── AutoML_App/
§── AutoML_Results/
§── データ/
§── メディア/
§── 要件.txt
━── train_automl.py
最も重要なファイルは次のとおりです。
train_automl.py
モデルをトレーニングし、Web アプリを生成します。データ ディレクトリにはトレーニング データとテスト データが含まれています。 AutoML_Results ディレクトリには、MLJAR AutoML トレーニングの結果が含まれています。 AutoML_App ディレクトリには、生成された Web アプリが含まれています。
Web アプリケーションを手動で作成する必要はありません。 MLJAR AutoML が生成します。
データのロードから始めましょう。
パンダをPDとしてインポートする
df = PD

。 read_csv ( "データ/train.csv" )
データセットには患者情報と医療保険料が含まれています。対象の列はchargeです。これが予測したい値です。
データを入力特徴とターゲットに分割します。
X_train = df 。ドロップ ( columns = [ "charges" ] )
y_train = df [ "料金" ]
X_train 変数には、すべての入力列が含まれます。 y_train 変数にはターゲット列が含まれます。
MLJAR AutoML を使用して機械学習モデルをトレーニングする
監視付きインポート AutoML から
automl = AutoML (
results_path = "AutoML_Results" ,
)
automl 。フィット ( X_train , y_train )
以上です。 MLJAR AutoML は、これが回帰タスクであることを自動的に検出します。いくつかのモデルをトレーニングし、その結果を比較します。
この例では、AutoML は次のようなモデルをチェックします。
このプロジェクトからのトレーニング出力は、最適なモデルがアンサンブル モデルであることを示しています。
これらすべてのアルゴリズムを手動でテストする必要はありません。 AutoML はこれを自動的に実行し、完全なレポートを AutoML_Results ディレクトリに保存します。これは、大量のトレーニング コードを記述するのではなく、問題とデータに集中できるため、初心者にとって非常に役立ちます。
Webアプリを自動生成
ここからが最も興味深い部分です。
トレーニングされた AutoML モデルから Web アプリを生成します。
automl 。アプリ（
パス = "AutoML_App" ,
title = "保険料予測ツール" ,
)
このコマンドは、すぐに使用できる Web アプリを作成します。アプリは AutoML_App ディレクトリに生成されます。
フロントエンドのコードを記述する必要はありません。フォームを手動で作成する必要はありません。レイアウトを設計する必要はありません。 MLJAR AutoML は、トレーニングされたモデルと入力データに基づいてアプリを作成します。生成されたアプリは Mercury を使用して作成されます。 Mercury は Python ノートブックを Web アプリケーションとして提供できるため、最終的なアプリは依然として Python に基づいています。
automl.app() を実行すると、 AutoML_App ディレクトリに次のファイルが含まれます。

け:
AutoML_App/
§── app_support.py
§── automl.zip
§── config.toml
§── mljar_app.json
━──predict_batch.ipynb
§── 予測シングル.ipynb
§── 要件.txt
└── ランタイム.txt
最も重要なノートブックは次のとおりです。
予測シングル.ipynb
予測バッチ.ipynb
単一予測には、predict_single.ipynb ノートブックが使用されます。 detect_batch.ipynb ノートブックは、CSV ファイルからのバッチ予測に使用されます。 automl.zip ファイルには、予測を行うために必要なモデル ランタイム ファイルが含まれています。 requirements.txt ファイルには、アプリの実行に必要な Python パッケージが含まれています。
両方のアプリケーションでの Mercury ページ ビュー:
最初に生成されたアプリは単一予測用です。このアプリでは、ユーザーは次のような値をフォームに入力できます。
ユーザーがフォーム内の値を変更すると、予測が自動的に更新されます。これにより、アプリが非常にインタラクティブになります。入力値を変更すると、予測される医療保険料にどのような影響を与えるかをすぐに確認できます。
このタイプのアプリは、一度に 1 つのケースを確認したい場合に最適です。たとえば、技術者以外のユーザーは、Python コードを作成しなくても、ブラウザーでアプリを開いてフォームに入力し、予測を取得できます。
アプリには、予測に関する詳細情報も表示されます。各入力値について、それがトレーニング データの分布とどのように比較されるかを表示します。これは、モデルのトレーニングに使用されたデータと比較して、選択した値が典型的な値であるか異常であるかをユーザーが理解するのに役立ちます。
さらに、アプリには特徴の重要度プロットが表示されます。このプロットは、どの特徴がモデルの予測に最大の影響を与えたかを示します。これは、モデルが意思決定を行うために何を使用しているのかをユーザーがよりよく理解できるため便利です。
つまり、単一の予測アプリは単純な予測フォームだけではありません。また、追加のコンテキストも提供され、

モデルの動作を説明するのに役立ちます。
2 番目に生成されたアプリはバッチ予測用です。このアプリでは、ユーザーは多くの行を含む CSV ファイルをアップロードします。アプリはすべての行の予測を計算し、結果を含むファイルを返します。これは、ユーザーがすでにスプレッドシートまたは CSV ファイルにデータを持っている場合に非常に便利です。たとえば、次のようなファイルを準備できます。
年齢、性別、BMI、子供、喫煙者、地域
19 、女性、27.9 、0 、はい、南西
35 、男性 、30.5 、1 、いいえ、北西
52 、女性、31.2 、2 、いいえ、南東
アプリは行ごとに予測を追加します。ユーザーは、「予測のダウンロード」ボタンを使用して、計算された予測をダウンロードできます。
生成されたアプリは Mercury を使用してローカルで実行できます。生成されたアプリ ディレクトリに移動します。
cd AutoML_App/アプリ
依存関係をインストールします。
pip install -r 要件.txt
水星を開始します。
水銀
Mercury はローカルサーバーを起動します。ブラウザでアプリを開いて、コンピューターでテストできます。これは、アプリをオンラインで公開する前の良い手順です。フォームが機能するかどうか、予測が正しいかどうか、バッチ CSV アップロードが期待どおりに機能するかどうかを確認できます。
MLJAR AutoML にはヘルパー関数もあります。
automl 。ローカルアプリ ( )
Python コードから直接ローカル アプリを生成して起動できます。
ローカルで実行するのは良いことですが、ローカル アプリを他の人と共有するのは簡単ではありません。アプリをオンラインで公開する最も簡単な方法は、MLJAR プラットフォームを使用することです。 1 つの Python コマンドでアプリを公開できます。
automl 。パブリッシュ_アプリ ( )
リポジトリでは、次の行がコメント化されています。
# automl.publish_app()
アプリを公開するには、# を削除してスクリプトを実行します。パブリッシュからの期待される出力は次のとおりです。
アプリの公開を開始する
アプリワークスペースの作成
MLJAR プラットフォームへのサインイン
アプリURLの作成
作成したアプリの URL: https://model-forge-5fbd.ismvp.org
ファイルをアップロードしています：predict_single.ipynb
ファイルをアップロードしています：predict_batch.ipy

注意
ファイルのアップロード: app_support.py
ファイルのアップロード: mljar_app.json
ファイルのアップロード: config.toml
ファイルのアップロード:requirements.txt
ファイルのアップロード: runtime.txt
ファイルのアップロード: automl.zip
完了しました。アプリには https://model-forge-5fbd.ismvp.org からアクセスできます。
MLJAR AutoML は、生成されたアプリケーション ファイルを MLJAR プラットフォームにアップロードし、アプリの URL を作成します。 MLJAR プラットフォームを開いて Web アプリを検査できます。
アップロードされたすべてのファイルを一覧表示できます。
公開後、リンクを他の人と共有できます。 Python をインストールしなくても、ブラウザーでアプリを開いて機械学習モデルを使用できます。これは、クライアント、チームメイト、ビジネス ユーザー、研究者、学生、友人などとモデルを共有する場合に非常に便利です。
このチュートリアルで使用される完全なスクリプトは次のとおりです。
パンダをPDとしてインポートする
監視付きインポート AutoML から
# データの読み取り
df = pd 。 read_csv ( "データ/train.csv" )
X_train = df 。ドロップ ( columns = [ "charges" ] )
y_train = df [ "料金" ]
# AutoML をトレーニングする
automl = AutoML (
results_path = "AutoML_Results" ,
)
automl 。フィット ( X_train , y_train )
# アプリを作成する
automl 。アプリ（
パス = "AutoML_App" ,
title = "保険料予測ツール" ,
)
# ローカルでアプリを起動
# automl.local_app()
# platform.mljar.com にアプリをデプロイする
# automl.publish_app()
これは短いスクリプトですが、多くの機能を備えています。
機械学習モデルをトレーニングします。
必要に応じて、アプリをローカルで起動します。
必要に応じて、アプリをオンラインで公開します。
これが、私がこのワークフローが好きな理由です。プロジェクトをシンプルに保ちます。
生成されたアプリに関する重要な注意事項
Web アプリは自動的に生成されます。これは、アプリのコードを手動で記述する必要がないことを意味します。機械学習の部分に集中できます。もちろん、上級ユーザーは、生成されたファイルを開いてカスタマイズすることもできます。ただし、初心者にとっては、デフォルトで生成されたアプリがすでに非常に便利です。私

単一予測とバッチ予測のための作業インターフェイスを提供します。多くのプロジェクトでは、モデルを他の人に見せてフィードバックを収集するにはこれで十分です。
このアプローチをいつ使用する必要がありますか?
このアプローチは、機械学習モデルを迅速に共有したい場合に適しています。
これは、モデルのユーザーが Python 開発者ではない場合に特に便利です。ノートブックや Python スクリプトを送信する代わりに、Web アプリへのリンクを送信できます。
このチュートリアルでは、機械学習モデルをトレーニングし、予測用の Web アプリを作成しました。モデルのトレーニングには MLJAR AutoML を使用しました。このモデルは、患者情報に基づいて医療保険料を予測します。
次に、automl.app() を使用して Web アプリを自動的に生成しました。
生成されたアプリには、2 つの便利な予測インターフェイスが含まれています。
1 つのサンプルに対して 1 つの予測アプリ。
CSV ファイルのバッチ予測アプリ。
また、Mercury を使用してアプリをローカルで実行する方法と、automl.publish_app() を使用してオンラインで公開する方法についても説明しました。
完全なコードは GitHub で入手できます。
https://github.com/pplonski/web-app-from-machine-learning-model
このワークフローはフロントエンド コードを記述する必要がないため、初心者に優しいです。 HTMLはありません。 CSSはありません。 JavaScriptは使用できません。パイソンのみ。楽しい予測をしてください！
コンピューター上の AI データ アナリスト
MLJAR Studio を使用してデータを探索し、洞察を見つけ、リポジトリを作成します

[切り捨てられた]

## Original Extract

Learn how to train a Machine Learning model with MLJAR AutoML, automatically generate a Web App, and publish it online for single and batch predictions.

Build a Web App for your Machine Learning model Products MLJAR Studio
Download Studio Jun 11 2026 · Piotr Płoński machine-learning
Build a Web App for your Machine Learning model
Your Machine Learning model is ready. What should you do next? You can save the model and use it in Python. You can create an API. You can add the model to an existing application. There are many ways to make predictions with a Machine Learning model .
In this tutorial, I will show you a beginner-friendly option: we will build a Web App for a Machine Learning model.
The best part? You don't need to write HTML, CSS, or JavaScript. Only Python.
We will train a Machine Learning model with MLJAR AutoML , automatically generate a Web App with app() , and then publish the app online with MLJAR Platform . The generated application is created with Mercury , a framework for building web apps from Python notebooks.
In the final result, we will have two web apps:
An app for single prediction, where the user fills the form and gets one prediction.
An app for batch prediction, where the user uploads a CSV file and gets predictions for many rows.
You can find the full code and results in my GitHub repository:
https://github.com/pplonski/web-app-from-machine-learning-model
In this tutorial, we will build an application for predicting medical insurance charges. The model will use patient demographic and health related factors, such as:
The target value is charges , which means medical insurance charges. This is a regression problem because we predict a number.
We will use MLJAR AutoML to train the model. AutoML will automatically check different Machine Learning algorithms, compare their performance, and select the best model. After training, we will generate a Web App from the trained model.
This is useful because many people don't want to work directly with Python code. They prefer a simple web interface. For example, your boss, client, doctor, researcher, or business teammate can open a link, enter values, and get the prediction.
Why create a Web App for a Machine Learning model?
Training a Machine Learning model is only one part of the project. The second part is making the model useful. If your model is hidden in a Python script or notebook, only technical users can use it. It is hard to share it with people who don't know Python.
A Web App solves this problem. You can give users a nice interface where they can:
use the model without touching code.
This is a very practical way to share Machine Learning models. In many projects, this can be enough to show a prototype, test an idea, or collect feedback from users.
pip install mljar-supervised mercury
The mljar-supervised package is used to train the Machine Learning model. The mercury package is used to run the generated Web App.
The requirements.txt file in the repository has only these two packages:
mljar-supervised
mercury
Project structure
The project repository contains the following files and directories:
web-app-from-machine-learning-model/
├── AutoML_App/
├── AutoML_Results/
├── data/
├── media/
├── requirements.txt
└── train_automl.py
The most important file is:
train_automl.py
It trains the model and generates the Web App. The data directory contains the training and test data. The AutoML_Results directory contains the results from MLJAR AutoML training. The AutoML_App directory contains the generated Web App.
You don't need to manually write the web application. MLJAR AutoML generates it for you.
Let's start with loading the data.
import pandas as pd
df = pd . read_csv ( "data/train.csv" )
The dataset contains patient information and medical insurance charges. The target column is charges . This is the value that we want to predict.
We split the data into input features and target:
X_train = df . drop ( columns = [ "charges" ] )
y_train = df [ "charges" ]
The X_train variable contains all input columns. The y_train variable contains the target column.
Train Machine Learning model with MLJAR AutoML
from supervised import AutoML
automl = AutoML (
results_path = "AutoML_Results" ,
)
automl . fit ( X_train , y_train )
That's all. MLJAR AutoML will automatically detect that this is a regression task. It will train several models and compare their results.
In this example, AutoML checks models like:
The training output from this project shows that the best model is an Ensemble model.
You don't need to manually test all these algorithms. AutoML does it for you and saves the full report in the AutoML_Results directory. This is very helpful for beginners because you can focus on the problem and the data, not on writing a lot of training code.
Generate Web App automatically
Now comes the most interesting part.
We will generate a Web App from the trained AutoML model.
automl . app (
path = "AutoML_App" ,
title = "Insurance Charges Predictor" ,
)
This command creates a ready-to-use Web App. The app is generated in the AutoML_App directory.
You don't need to write frontend code. You don't need to create forms manually. You don't need to design the layout. MLJAR AutoML creates the app based on your trained model and input data. The generated app is created with Mercury. Mercury can serve Python notebooks as web applications, so the final app is still based on Python.
After running automl.app() , the AutoML_App directory contains files like:
AutoML_App/
├── app_support.py
├── automl.zip
├── config.toml
├── mljar_app.json
├── predict_batch.ipynb
├── predict_single.ipynb
├── requirements.txt
└── runtime.txt
The most important notebooks are:
predict_single.ipynb
predict_batch.ipynb
The predict_single.ipynb notebook is used for single prediction. The predict_batch.ipynb notebook is used for batch prediction from a CSV file. The automl.zip file contains the model runtime files needed to make predictions. The requirements.txt file contains Python packages needed to run the app.
Mercury page view with both applications:
The first generated app is for single prediction. In this app, the user can enter values in a form, for example:
After the user changes any value in the form, the prediction is updated automatically. This makes the app very interactive. You can change the input values and immediately see how they affect the predicted medical insurance charges.
This type of app is great when someone wants to check one case at a time. For example, a non-technical user can open the app in the browser, fill in the form, and get the prediction without writing any Python code.
The app also shows more information about the prediction. For each input value, we display how it compares with the training data distribution. This helps the user understand whether the selected value is typical or unusual compared to the data used to train the model.
Additionally, the app displays a feature importance plot. This plot shows which features had the biggest impact on the model predictions. It is useful because the user can better understand what the model is using to make its decision.
So the single prediction app is not only a simple prediction form. It also gives extra context and helps explain the model behavior.
The second generated app is for batch prediction. In this app, the user uploads a CSV file with many rows. The app computes predictions for all rows and returns a file with results. This is very useful when users already have data in a spreadsheet or CSV file. For example, they can prepare a file like this:
age , sex , bmi , children , smoker , region
19 , female , 27.9 , 0 , yes , southwest
35 , male , 30.5 , 1 , no , northwest
52 , female , 31.2 , 2 , no , southeast
The app will add predictions for each row. User can download computed predictions with Download predictions button.
You can run the generated app locally with Mercury. Go to the generated app directory:
cd AutoML_App/app
Install dependencies:
pip install -r requirements.txt
Start Mercury:
mercury
Mercury will start a local server. You can open the app in your browser and test it on your computer. This is a good step before publishing the app online. You can check if the form works, if the predictions are correct, and if the batch CSV upload works as expected.
There is also a helper function in MLJAR AutoML:
automl . local_app ( )
It can generate and start the local app directly from Python code.
Running locally is nice, but sharing a local app with other people is not easy. The easiest way to publish the app online is to use MLJAR Platform. You can publish the app with one Python command:
automl . publish_app ( )
In the repository, this line is commented:
# automl.publish_app()
To publish the app, remove the # and run the script. The expected output from publish is:
Start app publish
Creating app workspace
Signing in to MLJAR platform
Creating app URL
Created app URL: https://model-forge-5fbd.ismvp.org
Uploading file: predict_single.ipynb
Uploading file: predict_batch.ipynb
Uploading file: app_support.py
Uploading file: mljar_app.json
Uploading file: config.toml
Uploading file: requirements.txt
Uploading file: runtime.txt
Uploading file: automl.zip
Finished. You can access your app at: https://model-forge-5fbd.ismvp.org
MLJAR AutoML will upload the generated application files to MLJAR Platform and create a URL for your app. You can open MLJAR Platform and inspect your web app:
You can list all uploaded files:
After publishing, you can share the link with other people. They can open the app in the browser and use your Machine Learning model without installing Python. This is very convenient when you want to share your model with: clients, teammates, business users, researchers, students, friends.
Here is the full script used in this tutorial:
import pandas as pd
from supervised import AutoML
# Read data
df = pd . read_csv ( "data/train.csv" )
X_train = df . drop ( columns = [ "charges" ] )
y_train = df [ "charges" ]
# Train AutoML
automl = AutoML (
results_path = "AutoML_Results" ,
)
automl . fit ( X_train , y_train )
# Create App
automl . app (
path = "AutoML_App" ,
title = "Insurance Charges Predictor" ,
)
# Start App locally
# automl.local_app()
# Deploy App on platform.mljar.com
# automl.publish_app()
It is a short script, but it does a lot:
Trains Machine Learning models.
Optionally starts the app locally.
Optionally publishes the app online.
This is why I like this workflow. It keeps the project simple.
Important note about generated apps
The Web App is generated automatically. This means that you don't need to manually write the app code. You can focus on the Machine Learning part. Of course, advanced users can later open the generated files and customize them. But for beginners, the default generated app is already very useful. It gives you a working interface for single and batch predictions. For many projects, this is enough to show your model to other people and collect feedback.
When should you use this approach?
This approach is good when you want to quickly share a Machine Learning model.
It is especially useful when the users of the model are not Python developers. Instead of sending them a notebook or Python script, you can send them a link to a Web App.
In this tutorial, we trained a Machine Learning model and created a Web App for predictions. We used MLJAR AutoML to train the model. The model predicts medical insurance charges based on patient information.
Then we used automl.app() to automatically generate a Web App.
The generated app contains two useful prediction interfaces:
Single prediction app for one sample.
Batch prediction app for CSV files.
We also saw how to run the app locally with Mercury and how to publish it online with: automl.publish_app()
The full code is available on GitHub:
https://github.com/pplonski/web-app-from-machine-learning-model
This workflow is beginner-friendly because you don't need to write frontend code. No HTML. No CSS. No JavaScript. Only Python. Happy predicting!
AI Data Analyst on Your Computer
Use MLJAR Studio to explore data, find insights, and create repo

[truncated]
