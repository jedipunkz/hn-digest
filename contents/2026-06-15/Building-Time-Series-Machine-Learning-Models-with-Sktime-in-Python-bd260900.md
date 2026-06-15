---
source: "https://www.kdnuggets.com/building-time-series-machine-learning-models-with-sktime-in-python"
hn_url: "https://news.ycombinator.com/item?id=48541527"
title: "Building Time-Series Machine Learning Models with Sktime in Python"
article_title: "Building Time-Series Machine Learning Models with sktime in Python - KDnuggets"
author: "eigenBasis"
captured_at: "2026-06-15T14:13:44Z"
capture_tool: "hn-digest"
hn_id: 48541527
score: 1
comments: 0
posted_at: "2026-06-15T14:08:33Z"
tags:
  - hacker-news
  - translated
---

# Building Time-Series Machine Learning Models with Sktime in Python

- HN: [48541527](https://news.ycombinator.com/item?id=48541527)
- Source: [www.kdnuggets.com](https://www.kdnuggets.com/building-time-series-machine-learning-models-with-sktime-in-python)
- Score: 1
- Comments: 0
- Posted: 2026-06-15T14:08:33Z

## Translation

タイトル: Python の SKtime を使用した時系列機械学習モデルの構築
記事のタイトル: Python の sktime を使用した時系列機械学習モデルの構築 - KDnuggets
説明: この記事では、sktime を使用して Python で時系列機械学習モデルを構築し、ワークフローを予測するためのコア データ構造を調査します。

記事本文:
Python の sktime を使用した時系列機械学習モデルの構築
この記事では、sktime を使用して Python で時系列機械学習モデルを構築し、ワークフローを予測するためのコア データ構造を探索します。
センサーの読み取り値、サーバーのメトリクス、または時間の経過とともに到着するデータを扱う場合は、標準の scikit-learn パイプラインが完全に適合しないことはすでにご存知でしょう。時系列データには、季節性、トレンド、時間的順序、将来の値が過去の値に依存するという事実など、表形式モデルでは無視される構造があります。
sktime は、このために特別に構築された Python ライブラリです。これは、scikit-learn スタイルの API (適合、予測、変換) を提供しますが、時系列向けにゼロから設計されています。時系列の予測、分類、回帰、クラスタリングをすべて一貫したインターフェイスで実行できます。
この記事では、産業用 HVAC センサーからの温度測定値の予測という問題例に取り組みます。 sktime が時系列データを処理する方法、前処理パイプラインを構築する方法、予測者を適合させる方法、およびそれらを評価する方法を学びます。
コードは GitHub で入手できます。
Python 3.10 以降と、パンダに関する基本的な知識が必要です。必要なものをすべてインストールします。
pip install sktime pmdarima statsmodels
すべてのオプションの依存関係を一度に取得したい場合は、 pip install sktime[all_extras] でカバーします。
sktime が解決している問題を理解するのに役立ちます。 scikit-learn では、データは 2D テーブルです。行はサンプル、列は特徴です。時系列データは、各「行」が実際には時間の経過に伴う一連の値であり、それらの値の順序が重要であるため、この仮定を破ります。
使用する主なデータ コンテナは次のとおりです。
時間インデックス自体については、sktime はいくつかの時間インデックスをサポートしています: DatetimeIndex 、 PeriodIndex 、 Int64In

dex 、および pandas オブジェクトの RangeIndex 。インデックスは単調である必要があります。 DatetimeIndex を使用している場合は、 freq 属性を設定する必要があります。
現実的なデータセットを作成してみましょう。工場にある HVAC センサーが温度を 1 時間ごとに記録していると想像してください。測定値には毎日の季節的パターン (勤務時間中に高くなる)、夏によるわずかな上昇傾向、および多少のノイズが見られます。
numpyをnpとしてインポート
パンダをPDとしてインポートする
np.random.seed(42)
# 2026 年 1 月 1 日から始まる 90 日間の毎時測定値
n_時間 = 90 * 24
タイムスタンプ = pd.date_range(start="2026-01-01", period=n_hours, freq="h")
# 傾向: 90 日間で徐々に 5 度上昇
トレンド = np.linspace(0, 5, n_hours)
# 毎日の季節性: 気温は午後 2 時に最高になり、午前 4 時に下がります。
1 日の時間 = np.arange(n 時間) % 24
daily_cycle = 4 * np.sin(2 * np.pi * (1 日の時間 - 4) / 24)
# ノイズ
ノイズ = np.random.normal(0, 0.8, n_hours)
# 基礎温度20℃前後
温度 = 20 + トレンド + 日次サイクル + ノイズ
# いくつかの欠損値 (センサーのドロップアウト) を導入します。
ドロップアウトインデックス = [300, 301, 302, 1440, 1441]
温度[ドロップアウトインデックス] = np.nan
y = pd.Series(温度, インデックス=タイムスタンプ, name="temp_celsius")
y.index.freq = pd.tseries.frequency.to_offset("h")
print(y.head())
print(f"\n形状: {y.shape}")
print(f"欠落値: {y.isna().sum()}")
print(f"インデックスタイプ: {type(y.index)}")
2026-01-01 00:00:00 16.933270
2026-01-01 01:00:00 17.063277
2026-01-01 02:00:00 18.522783
2026-01-01 03:00:00 20.190095
2026-01-01 04:00:00 19.821941
周波数: h、名前: temp_celsius、dtype: float64
形状: (2160,)
欠損値: 5
インデックスの種類:
# トレーニングとテスト用の時系列データの分割
時系列データの分割は表形式データとは異なります。行をシャッフルすることはできません。常に時系列に分割する必要があります。つまり、以前のデータでトレーニングし、後のデータでテストします。
sktime は温度を提供します

この目的のために、oral_train_test_split:
sktime.splitからインポートtemporal_train_test_split
# 過去 7 日間 (168 時間) をテストセットとして保持します
y_train, y_test = Temporal_train_test_split(y, test_size=168)
print(f"トレイン: {y_train.index[0]} → {y_train.index[-1]}")
print(f"テスト: {y_test.index[0]} → {y_test.index[-1]}")
print(f"トレーニング サイズ: {len(y_train)}、テスト サイズ: {len(y_test)}")
電車: 2026-01-01 00:00:00 → 2026-03-24 23:00:00
テスト: 2026-03-25 00:00:00 → 2026-03-31 23:00:00
トレインサイズ: 1992、テストサイズ: 168
この関数により、分割がクリーンかつ時系列に行われることが保証され、将来からトレーニング セットにデータが漏洩することはありません。
# 予測の範囲を定義する
モデルをフィッティングする前に、どのタイムステップを予測したいかを sktime に伝える必要があります。これが ForecastingHorizo​​n です。
sktime.forecasting.base から ForecastingHorizon をインポート
# 168 歩先を予測 (7 日間の毎時データ)
# is_relative=False は、絶対タイムスタンプを使用していることを意味します
fh = ForecastingHorizon(y_test.index, is_relative=False)
print(f"水平線の長さ: {len(fh)}")
print(f"最初の予測点: {fh[0]}")
print(f"最後の予測点: {fh[-1]}")
水平線の長さ: 168
最初の予測ポイント: 2026-03-25 00:00:00
最終予測ポイント: 2026-03-31 23:00:00
fh = [1, 2, 3, ..., 168] のような相対的な範囲を使用することもできます。これは、「1 ステップ先、2 ステップ先、...」を意味します。予測が必要な実際のタイムスタンプがある場合、絶対的なホライズンがより明確になります。
# 前処理および予測パイプラインの構築
実際のセンサー データには欠損値、季節的パターン、傾向が含まれており、予測前または予測中にこれらすべてを処理する必要があります。 sktime の TransformedTargetForecaster を使用すると、予測器による変換を単一の推定器に連鎖させることができます。変換は、フィッティングの前にターゲット系列 y に適用されます。

g、予測中に途中で自動的に逆転します。
sktime.forecasting.exp_smoothing からのインポート ExponentialSmoothing
sktime.forecasting.compose からインポート TransformedTargetForecaster
sktime.transformations.series.impute から import Imputer
sktime.transformations.series.detrend からインポート Deseasonalizer、Detrender
パイプライン = TransformedTargetForecaster(
ステップ=[
# ステップ 1: 線形補間を使用して欠落しているセンサー読み取り値を埋める
("インピューター", Imputer(method="linear")),
# ステップ 2: 予測者が定常系列を認識できるように、線形トレンドを削除します。
("detrender"、Detrender())、
# ステップ 3: 日次の季節性を削除します (24 時間サイクルの 1 時間ごとのデータの場合は sp=24)
("deseasonalizer", Deseasonalizer(model="additive", sp=24)),
# ステップ 4: クリーン化された定常残差を予測する
("予測者", ExponentialSmoothing(トレンド=なし、季節=なし)),
】
)
パイプライン.fit(y_train, fh=fh)
y_pred = パイプライン.predict()
print(y_pred.head())
2026-03-25 00:00:00 21.210066
2026-03-25 01:00:00 21.788986
2026-03-25 02:00:00 22.615184
2026-03-25 03:00:00 23.688449
2026-03-25 04:00:00 24.621127
周波数: h、名前: temp_celsius、dtype: float64
Imputer(method="linear") は、周囲の読み取り値の間を線形補間することで欠損値を埋めます。これはセンサー データに適しています。
Detrender() は、線形トレンドをトレーニング シリーズに当てはめて減算します。予測時にはトレンドを追加します。
Deseasonalizer(sp=24) は残差から 24 時間サイクルを削除します。 spは季節期間の略です。
最後に、ExponentialSmoothing は、トレンドを除去し、季節性を除去した残差を予測します。
detect() が呼び出されると、すべての逆変換が逆の順序で自動的に適用され、元の温度スケールで予測が戻されます。
sktime は標準の評価指標と統合されています。予測の場合、平均絶対値

一般的な選択としては、誤差 (MAE) と平均絶対パーセント誤差 (MAPE) があります。
sktime.performance_metrics.forecasting import から (
平均絶対誤差、
平均絶対パーセント誤差、
)
mae = 平均絶対誤差(y_test, y_pred)
Mape = means_absolute_percentage_error(y_test, y_pred)
print(f"MAE: {mae:.3f} °C")
print(f"MAPE: {mape*100:.2f}%")
# 別の Forecaster に切り替える
sktime インターフェイスの最大の利点の 1 つは、基礎となるアルゴリズムの交換に必要な変更が 1 行だけであることです。指数平滑法の代わりに ARIMA モデルを試して比較してみましょう。
sktime.forecasting.arima から ARIMA をインポート
Pipeline_arima = TransformedTargetForecaster(
ステップ=[
("インピューター", Imputer(method="linear")),
("detrender"、Detrender())、
("deseasonalizer", Deseasonalizer(model="additive", sp=24)),
# クリーン化された残差に対する ARIMA(1,1,1)
("予報者", ARIMA(order=(1, 1, 1), stop_warnings=True)),
】
)
Pipeline_arima.fit(y_train, fh=fh)
y_pred_arima = Pipeline_arima.predict()
mae_arima = 平均絶対誤差(y_test, y_pred_arima)
mape_arima = means_absolute_percentage_error(y_test, y_pred_arima)
print(f"有馬前: {mae_arima:.3f} ℃")
print(f"有馬マップ: {mape_arima*100:.2f}%")
有馬前: 0.586 °C
有馬楓：2.41%
重要な点は、前処理ステップ (代入、トレンド除去、非季節化) が同一のままであることです。最後の予測者を変更しただけで、他のすべてはそれを中心にきれいに構成されました。
# 時間を超えた相互検証
単一のテスト期間を延長することは誤解を招く可能性があります。 sktime は、時間的な順序を尊重するスプリッターを介して時系列の相互検証を提供します。
SlidingWindowSplitter はローリング ウィンドウを使用します。トレーニング ウィンドウは時間とともに前方にスライドし、常に同じ長さを保ちます。 ExpandingWindowSplitter は、トレーニング セットを累積的に拡大します。

これは、利用可能なすべての履歴を使用したい場合に適しています。
sktime.split からインポート ExpandingWindowSplitter
sktime.forecasting.model_evaluation からインポート評価
# ウィンドウの拡張: 1800 時間の列車セットから開始し、168 時間のウィンドウで評価します
cv = ExpandingWindowSplitter(
初期ウィンドウ=1800、
fh=リスト(範囲(1, 169)),
ステップ長=168、
)
結果 = 評価(
予測者=パイプライン、
y=y、
CV=CV、
スコアリング = 平均絶対誤差、
return_data=False、
)
print(results[["test__DynamicForecastingErrorMetric", "fit_time"]].round(3))
print(f"\n平均 CV MAE: {results['test__DynamicForecastingErrorMetric'].mean():.3f} °C")
test__DynamicForecastingErrorMetric fit_time
0 0.627 0.274
1 0.585 0.100
平均CV MAE: 0.606 °C
Evaluate は、フォールドごとのメトリクスとタイミングを含む DataFrame を返します。相互検証 MAE は、データ内のさまざまな時間枠にわたってモデルが一貫して一般化していることを確認します。
この記事では sktime の中核となる予測ワークフローについて説明しましたが、ライブラリは基本的な予測タスクをはるかに超えています。
また、時系列分類、不確実性推定による確率的予測、複数の関連する時系列にわたる共有モデルのトレーニング、逐次予測のための従来の機械学習アルゴリズムの適応、モデル選択とワークフローの調整の自動化もサポートしています。
sktime の最大の強みの 1 つは、一貫した API と広範な Python 機械学習エコシステムとの統合であり、初心者と経験豊富な実践者の両方にとって実験が容易になります。 sktime のドキュメントとサンプル ノートブックは特によく書かれており、予測や時間データの問題を定期的に扱う場合はブックマークする価値があります。
Bala Priya C はインド出身の開発者兼テクニカル ライターです。彼女は数学とプログラムの交差点で働くのが好きです

ng、データサイエンス、コンテンツ作成。彼女の興味と専門分野には、DevOps、データ サイエンス、自然言語処理が含まれます。彼女は読書、執筆、コーディング、コーヒーが好きです。現在、彼女はチュートリアル、ハウツー ガイド、意見記事などを作成して、学習し、開発者コミュニティと知識を共有することに取り組んでいます。 Bala は、魅力的なリソースの概要やコーディング チュートリアルも作成しています。
Rust での高性能機械学習モデルの構築
実際に役立つ機械学習モデルを構築するためのヒント
予測モデルの構築: Python でのロジスティック回帰
大規模な言語モデルを使用したアプリを作成するためのデータ パイプラインの構築
雇用を獲得する機械学習ポートフォリオを構築するための究極のガイド
Django を使用した機械学習アプリケーションの構築
おすすめの無料コース トップ 5
無料の電子ブック「KDnuggets Artificial Intelligence Pocket Dictionary」と、データ サイエンス、機械学習、AI、分析に関する主要なニュースレターを直接受信箱に入手してください。
購読すると、KDnuggets プライバシー ポリシーに同意したことになります
Python の sktime を使用した時系列機械学習モデルの構築 3 データ クリーニングと準備のための Pandas のトリック クロード コードとローカル モデルのペアリング 数値 Perfor のための 3 つの NumPy のトリック

[切り捨てられた]

## Original Extract

In this article, we’ll build time-series machine learning models in Python using sktime and explore its core data structures for forecasting workflows.

Building Time-Series Machine Learning Models with sktime in Python
In this article, we’ll build time-series machine learning models in Python using sktime and explore its core data structures for forecasting workflows.
If you work with sensor readings, server metrics, or any data that arrives over time, you already know that standard scikit-learn pipelines don't quite fit. Time series data has structure that tabular models ignore: seasonality, trend, temporal ordering, and the fact that future values depend on past ones.
sktime is a Python library built specifically for this. It gives you a scikit-learn-style API — fit, predict, transform — but designed from the ground up for time series. You can do forecasting, classification, regression, and clustering on time series, all with a consistent interface.
In this article, you'll work through an example problem: forecasting temperature readings from an industrial HVAC sensor. You'll learn how sktime handles time series data, how to build preprocessing pipelines, how to fit forecasters, and how to evaluate them.
You can get the code on GitHub .
You'll need Python 3.10 or higher and a basic familiarity with pandas. Install everything you need with:
pip install sktime pmdarima statsmodels
If you'd rather have all optional dependencies in one shot, pip install sktime[all_extras] covers them.
It helps to understand the problem sktime is solving. In scikit-learn, your data is a 2D table — rows are samples, columns are features. Time series data breaks this assumption because each "row" is actually a sequence of values over time, and the order of those values matters.
The main data containers you'll use are:
For the time index itself, sktime supports several time indexes: DatetimeIndex , PeriodIndex , Int64Index , and RangeIndex on your pandas objects. The index must be monotonic. If you're using DatetimeIndex , the freq attribute should be set.
Let's create a realistic dataset. Imagine an HVAC sensor in a factory that records temperature every hour. The readings have a daily seasonal pattern (higher during working hours), a slight upward trend due to summer, and some noise.
import numpy as np
import pandas as pd
np.random.seed(42)
# 90 days of hourly readings starting Jan 1, 2026
n_hours = 90 * 24
timestamps = pd.date_range(start="2026-01-01", periods=n_hours, freq="h")
# Trend: gradual 5-degree rise over 90 days
trend = np.linspace(0, 5, n_hours)
# Daily seasonality: temperature peaks at 2pm, dips at 4am
hour_of_day = np.arange(n_hours) % 24
daily_cycle = 4 * np.sin(2 * np.pi * (hour_of_day - 4) / 24)
# Noise
noise = np.random.normal(0, 0.8, n_hours)
# Base temperature around 20°C
temperature = 20 + trend + daily_cycle + noise
# Introduce a few missing values (sensor dropout)
dropout_indices = [300, 301, 302, 1440, 1441]
temperature[dropout_indices] = np.nan
y = pd.Series(temperature, index=timestamps, name="temp_celsius")
y.index.freq = pd.tseries.frequencies.to_offset("h")
print(y.head())
print(f"\nShape: {y.shape}")
print(f"Missing values: {y.isna().sum()}")
print(f"Index type: {type(y.index)}")
2026-01-01 00:00:00 16.933270
2026-01-01 01:00:00 17.063277
2026-01-01 02:00:00 18.522783
2026-01-01 03:00:00 20.190095
2026-01-01 04:00:00 19.821941
Freq: h, Name: temp_celsius, dtype: float64
Shape: (2160,)
Missing values: 5
Index type:
# Splitting Time Series Data for Training and Testing
Splitting time series data is different from tabular data — you can't shuffle rows. You must always split chronologically: train on earlier data, test on later data.
sktime provides temporal_train_test_split for this purpose:
from sktime.split import temporal_train_test_split
# Hold out the last 7 days (168 hours) as the test set
y_train, y_test = temporal_train_test_split(y, test_size=168)
print(f"Train: {y_train.index[0]} → {y_train.index[-1]}")
print(f"Test: {y_test.index[0]} → {y_test.index[-1]}")
print(f"Train size: {len(y_train)}, Test size: {len(y_test)}")
Train: 2026-01-01 00:00:00 → 2026-03-24 23:00:00
Test: 2026-03-25 00:00:00 → 2026-03-31 23:00:00
Train size: 1992, Test size: 168
The function ensures the split is clean and chronological — no data leakage from the future into the training set.
# Defining the Forecasting Horizon
Before fitting any model, you need to tell sktime which time steps you want to predict. This is the ForecastingHorizon .
from sktime.forecasting.base import ForecastingHorizon
# Predict 168 steps ahead (7 days of hourly data)
# is_relative=False means we're using absolute timestamps
fh = ForecastingHorizon(y_test.index, is_relative=False)
print(f"Horizon length: {len(fh)}")
print(f"First forecast point: {fh[0]}")
print(f"Last forecast point: {fh[-1]}")
Horizon length: 168
First forecast point: 2026-03-25 00:00:00
Last forecast point: 2026-03-31 23:00:00
You can also use relative horizons like fh = [1, 2, 3, ..., 168] , which means "1 step ahead, 2 steps ahead, ...". Absolute horizons are cleaner when you have actual timestamps you want predictions for.
# Building a Preprocessing and Forecasting Pipeline
Real sensor data has missing values, seasonal patterns, and trend — you need to handle all of these before or during forecasting. sktime's TransformedTargetForecaster lets you chain transformations with a forecaster into a single estimator. The transformations are applied to the target series y before fitting, and automatically reversed on the way out during prediction.
from sktime.forecasting.exp_smoothing import ExponentialSmoothing
from sktime.forecasting.compose import TransformedTargetForecaster
from sktime.transformations.series.impute import Imputer
from sktime.transformations.series.detrend import Deseasonalizer, Detrender
pipeline = TransformedTargetForecaster(
steps=[
# Step 1: Fill missing sensor readings using linear interpolation
("imputer", Imputer(method="linear")),
# Step 2: Remove the linear trend so the forecaster sees a stationary series
("detrender", Detrender()),
# Step 3: Remove the daily seasonality (sp=24 for hourly data with 24-hour cycles)
("deseasonalizer", Deseasonalizer(model="additive", sp=24)),
# Step 4: Forecast the cleaned, stationary residuals
("forecaster", ExponentialSmoothing(trend=None, seasonal=None)),
]
)
pipeline.fit(y_train, fh=fh)
y_pred = pipeline.predict()
print(y_pred.head())
2026-03-25 00:00:00 21.210066
2026-03-25 01:00:00 21.788986
2026-03-25 02:00:00 22.615184
2026-03-25 03:00:00 23.688449
2026-03-25 04:00:00 24.621127
Freq: h, Name: temp_celsius, dtype: float64
Imputer(method="linear") fills missing values by linearly interpolating between the surrounding readings, which works well for sensor data.
Detrender() fits a linear trend to the training series and subtracts it; on prediction it adds the trend back.
Deseasonalizer(sp=24) removes the 24-hour cycle from the residuals; sp stands for seasonal period.
Finally, ExponentialSmoothing forecasts the detrended, deseasonalized residuals.
When predict() is called, all inverse transformations are applied in reverse order automatically, and you get back predictions in the original temperature scale.
sktime integrates with standard evaluation metrics . For forecasting, mean absolute error (MAE) and mean absolute percentage error (MAPE) are common choices.
from sktime.performance_metrics.forecasting import (
mean_absolute_error,
mean_absolute_percentage_error,
)
mae = mean_absolute_error(y_test, y_pred)
mape = mean_absolute_percentage_error(y_test, y_pred)
print(f"MAE: {mae:.3f} °C")
print(f"MAPE: {mape*100:.2f}%")
# Swapping in a Different Forecaster
One of the biggest advantages of the sktime interface is that swapping the underlying algorithm requires changing just one line. Let's try an ARIMA model in place of exponential smoothing and compare.
from sktime.forecasting.arima import ARIMA
pipeline_arima = TransformedTargetForecaster(
steps=[
("imputer", Imputer(method="linear")),
("detrender", Detrender()),
("deseasonalizer", Deseasonalizer(model="additive", sp=24)),
# ARIMA(1,1,1) on the cleaned residuals
("forecaster", ARIMA(order=(1, 1, 1), suppress_warnings=True)),
]
)
pipeline_arima.fit(y_train, fh=fh)
y_pred_arima = pipeline_arima.predict()
mae_arima = mean_absolute_error(y_test, y_pred_arima)
mape_arima = mean_absolute_percentage_error(y_test, y_pred_arima)
print(f"ARIMA MAE: {mae_arima:.3f} °C")
print(f"ARIMA MAPE: {mape_arima*100:.2f}%")
ARIMA MAE: 0.586 °C
ARIMA MAPE: 2.41%
The key point is that the preprocessing steps — imputation, detrending, deseasonalization — stayed identical. You only changed the final forecaster, and everything else composed cleanly around it.
# Cross-Validating Across Time
Holding out a single test window can be misleading. sktime provides time series cross-validation through splitters that respect temporal ordering.
SlidingWindowSplitter uses a rolling window: the training window slides forward in time, always staying the same length. ExpandingWindowSplitter grows the training set cumulatively as you move forward, which is more appropriate when you want to use all available history.
from sktime.split import ExpandingWindowSplitter
from sktime.forecasting.model_evaluation import evaluate
# Expanding window: start with 1800-hour train set, evaluate on 168-hour windows
cv = ExpandingWindowSplitter(
initial_window=1800,
fh=list(range(1, 169)),
step_length=168,
)
results = evaluate(
forecaster=pipeline,
y=y,
cv=cv,
scoring=mean_absolute_error,
return_data=False,
)
print(results[["test__DynamicForecastingErrorMetric", "fit_time"]].round(3))
print(f"\nMean CV MAE: {results['test__DynamicForecastingErrorMetric'].mean():.3f} °C")
test__DynamicForecastingErrorMetric fit_time
0 0.627 0.274
1 0.585 0.100
Mean CV MAE: 0.606 °C
evaluate returns a DataFrame with per-fold metrics and timing. The cross-validation MAE confirms that the model generalizes consistently across different time windows in the data.
This article covered the core forecasting workflow in sktime, but the library extends far beyond basic prediction tasks.
It also supports time-series classification , probabilistic forecasting with uncertainty estimates, training shared models across multiple related time series, adapting traditional machine learning algorithms for sequential forecasting, and automating model selection and tuning workflows .
One of sktime's biggest strengths is its consistent API and integration with the broader Python machine learning ecosystem, making experimentation easier for both beginners and experienced practitioners. The sktime docs and example notebooks are especially well-written and are worth bookmarking if you regularly work with forecasting or temporal data problems.
Bala Priya C is a developer and technical writer from India. She likes working at the intersection of math, programming, data science, and content creation. Her areas of interest and expertise include DevOps, data science, and natural language processing. She enjoys reading, writing, coding, and coffee! Currently, she's working on learning and sharing her knowledge with the developer community by authoring tutorials, how-to guides, opinion pieces, and more. Bala also creates engaging resource overviews and coding tutorials.
Building High-Performance Machine Learning Models in Rust
Tips for Building Machine Learning Models That Are Actually Useful
Building Predictive Models: Logistic Regression in Python
Building Data Pipelines to Create Apps with Large Language Models
The Ultimate Guide to Building a Machine Learning Portfolio That Lands Jobs
Building Machine Learning Application with Django
Our Top 5 Free Course Recommendations
Get the FREE ebook 'KDnuggets Artificial Intelligence Pocket Dictionary' along with the leading newsletter on Data Science, Machine Learning, AI & Analytics straight to your inbox.
By subscribing you accept KDnuggets Privacy Policy
Building Time-Series Machine Learning Models with sktime in Python 3 Pandas Tricks for Data Cleaning & Preparation Pairing Claude Code with Local Models 3 NumPy Tricks for Numerical Perfor

[truncated]
