---
source: "https://cookbooks.datacebo.com/cookbook/real_vs_synthetic_vs_fake_data/"
hn_url: "https://news.ycombinator.com/item?id=50382872"
title: "Benefits of Performance Testing with AI-Generated Synthetic Data"
article_title: "Real Vs Synthetic Vs Fake Data | DataCebo Cookbooks"
image: "https://cookbooks.datacebo.com/og-image.png"
author: "skadamat"
captured_at: "2026-11-05T19:36:05Z"
capture_tool: "hn-digest"
hn_id: 50382872
score: 2
comments: 0
posted_at: "2026-11-05T18:53:06Z"
tags:
  - hacker-news
---

# Benefits of Performance Testing with AI-Generated Synthetic Data

- HN: [50382872](https://news.ycombinator.com/item?id=50382872)
- Source: [cookbooks.datacebo.com](https://cookbooks.datacebo.com/cookbook/real_vs_synthetic_vs_fake_data/)
- Score: 2
- Comments: 0
- Posted: 2026-11-05T18:53:06Z

## Translation

Title: Benefits of Performance Testing with AI-Generated Synthetic Data
Article title: Real Vs Synthetic Vs Fake Data | DataCebo Cookbooks
Description: Real Vs Synthetic Vs Fake Data — a step-by-step DataCebo cookbook for generating synthetic data with the Synthetic Data Vault (SDV).

Article text:
Real Vs Synthetic Vs Fake Data | DataCebo Cookbooks cookbooks
About API Docs Source Back May 6, 2026
Benefits of Performance Testing with AI-Generated Synthetic Data ¶
Software systems need close-to-real-world data to test against. Teams typically choose between three approaches:
In this tutorial, we'll generate fake data with the DayZSynthesizer and synthetic data with the GaussianCopulaSynthesizer , then compare both against the real data — first statistically, then through a simulated fraud detection scenario.
When to use this cookbook: You're evaluating whether to use fake or synthetic data in place of production data — for example, when populating a dev/test environment or stress-testing a downstream service like a fraud detection pipeline.
Prerequisite: DayZSynthesizer is part of SDV Enterprise. The GaussianCopulaSynthesizer used for the synthetic-data sections is available in SDV Community.
import warnings
warnings . filterwarnings ( 'ignore' )
Loading and Exploring the Data ¶
We'll work with the credit_card_transactions dataset. Here's an overview of its columns:
User : user ID linking transactions together
Card : card issuer, encoded as an integer
Year : year the transaction occurred
Month : month the transaction occurred
Day : day the transaction occurred
Time : time the transaction occurred
Use Chip : if an EMV chip card was used, or if the transaction was online
Merchant Name : name of the merchant
Merchant City : registered city for the merchant
Merchant State : registered state for the merchant
Zip : registered zip code for the merchant
MCC : merchant category code, a 4-digit number identifying the type of transaction
Errors? : any errors experienced as part of the transaction
Is Fraud? : whether the transaction was ultimately fraudulent
from sdv.datasets.demo import download_demo
datasets , full_metadata = download_demo (
modality = 'multi_table' ,
dataset_name = 'credit_card_transactions'
)
full = datasets [ 'credit_card_transactions-ibm_v2' ]
full [ "Amount" ] = full [ "Amount" ] . str . replace ( "$" , "" ) . astype ( float )
full [ "Zip" ] = full [ "Zip" ] . astype ( str ) . str . replace ( ".0" , "" , regex = False )
full [ "Time" ] = full [ "Time" ] . str . split ( ":" ) . str [ 0 ] . astype ( int )
real_data = full . sample ( frac = 0.1 , random_state = 0 ) . reset_index ( drop = True )
del datasets , full
metadata = full_metadata . tables [ 'credit_card_transactions-ibm_v2' ]
for col , sdtype in [
( "User" , "categorical" ),
( "Card" , "categorical" ),
( "Amount" , "numerical" ),
( "Merchant City" , "categorical" ),
( "Merchant State" , "categorical" ),
( "Time" , "numerical" ),
]:
metadata . update_column ( column_name = col , sdtype = sdtype )
Let's preview the data:
real_data . head ()
User
Card
Year
Month
Day
Time
Amount
Use Chip
Merchant Name
Merchant City
Merchant State
Zip
MCC
Errors?
Is Fraud?
0
1264
0
2015
10
20
19
59.63
Swipe Transaction
2301046522991844568
Westlake
OH
44145
5813
NaN
No
1
1693
3
2011
8
19
17
41.63
Swipe Transaction
6661973303171003879
North Hollywood
CA
91605
5211
NaN
No
2
763
2
2009
2
10
17
16.84
Swipe Transaction
-8288218236188507426
Pomeroy
OH
45769
5812
NaN
No
3
1133
1
2006
3
16
13
13.32
Swipe Transaction
-3150043395187512028
Edinburgh
IN
46124
5921
NaN
No
4
550
0
2020
2
25
3
8.90
Chip Transaction
-337314228625997565
Parkton
NC
28371
5921
NaN
No
Now let's explore the dataset to understand its key properties — how many transactions and users it contains, where the values fall, and how the categorical columns relate. Each of the following questions is designed to surface something important about the data that the synthesizers will need to reproduce.
How large is the dataset and how many unique users does it have?
print ( f 'Sample rows: { len ( real_data ) : , } ' )
print ( f 'Columns: { real_data . shape [ 1 ] } ' )
print ( f 'Unique users: { real_data [ "User" ] . nunique () : , } ' )
Sample rows: 2,438,690
Columns: 15
Unique users: 2,000
The dataset is large enough to exercise statistical modeling — over a million transactions across a couple thousand unique users. That gives the synthesizer plenty of signal to learn distributions and correlations from, and gives us enough rows to see meaningful differences between real, fake, and synthetic data downstream.
Which columns have missing values?
# Check for missing values
missing = real_data . isnull () . sum ()
missing [ missing > 0 ]
Merchant State 272248
Errors? 2399915
dtype: int64
Only Merchant State and Errors? have missing values. We'll configure our fake data generator to match these proportions later.
What does the distribution of transaction amounts look like?
The vast majority of transactions are small — clustered near \$0 — with a long right tail extending past \$1,500. This heavy skew is typical of real-world spending behavior, where everyday purchases dominate and large transactions are rare.
Breaking it down into ranges reveals more detail:
Negative amounts represent refunds or chargebacks — they peak around -\$75. The \$0–\$250 range shows a steep decline from very small purchases, with periodic spikes at round-dollar amounts. Transactions above \$250 are sparse, forming a long tail out to \$1,500+.
What years are represented in this dataset?
Transaction volume grows steadily from the 2000s through 2020, reflecting increasing credit card adoption over time. The drop visible in 2020 reflects that the dataset only covers the first few months of that year (as the dataset was created at the beginning of 2020), not a real decrease in transaction volume.
How balanced is the fraud label?
The fraud label is highly imbalanced — only a tiny fraction of transactions are flagged as fraud. This is realistic: real-world fraud rates are typically below 1% of transactions. The imbalance matters for any downstream modeling task: a naive model that always predicts "not fraud" would already be ~99% accurate, so we'll need to be careful about how we evaluate performance later.
How do the key categorical columns relate to each other?
The chart below shows the flow between payment method, transaction errors, and fraud status.
The Sankey chart shows how transactions flow across payment method, error status, and fraud label. The thick band on the right indicates that most transactions have no recorded errors — the Errors? column is blank (i.e. null) for the vast majority of rows. The smaller bands break out the rare error types: bad PIN, insufficient balance, technical glitches, etc. Fraud is rare across all payment methods and is concentrated among transactions with no recorded error.
Generating Fake Data with DayZSynthesizer ¶
The DayZSynthesizer generates fake data by randomly sampling values from user-defined ranges and categories. Unlike a statistical model, it does not learn any patterns or correlations from the real data.
To make the fake data as realistic as possible, we configure it step by step with information from the real dataset.
First, we instantiate a DayZSynthesizer from the metadata. The synthesizer starts in a default state where it knows the column types but nothing about the actual values, ranges, or null rates in the real data — those are the things we configure in the next steps.
from sdv.single_table import DayZSynthesizer
fake_synthesizer = DayZSynthesizer ( metadata )
Next, we tell the synthesizer the missing-value proportions for each column that has nulls. By default, DayZSynthesizer assumes every column is fully populated, which would produce fake data without any nulls. Real credit card data has nulls in Merchant State and Errors? , so we explicitly set those proportions to match — otherwise the fake data would look unrealistically complete.
# Set missing value proportions to match real data
for col in [ 'Merchant State' , 'Errors?' ]:
fake_synthesizer . set_missing_values (
column_name = col ,
missing_values_proportion = round (
real_data [ col ] . isnull () . sum () / len ( real_data ), 10
)
)
Next, we tell the synthesizer the valid categories for each categorical column. Without this, the synthesizer doesn't know what values are legal for Merchant Name , Use Chip , etc., and would either fail to generate them or produce placeholders. We pass in the unique values from the real data so the fake data uses the same vocabulary.
cat_columns = [
'User' , 'Card' , 'Use Chip' , 'Merchant Name' ,
'Merchant City' , 'Merchant State' , 'Zip' , 'MCC' ,
'Errors?' , 'Is Fraud?'
]
for col in cat_columns :
fake_synthesizer . set_category_values (
column_name = col ,
category_values = real_data [ col ] . dropna () . unique () . tolist ()
)
Finally, we set the numerical and datetime bounds for each non-categorical column. Without these, the synthesizer has no idea what range to sample from for Amount , Year , Month , or Day . We use the real data's min/max for each — this gives the fake data realistic ranges, even though the values within those ranges are sampled uniformly at random.
for col in [ 'Amount' , 'Year' , 'Month' , 'Day' , 'Time' ]:
cast = float if col == 'Amount' else int
fake_synthesizer . set_numerical_bounds (
column_name = col ,
min_value = cast ( real_data [ col ] . min ()),
max_value = cast ( real_data [ col ] . max ())
)
if col != 'Amount' :
fake_synthesizer . set_rounding_scheme (
column_name = col ,
num_decimal_digits = 0
)
Now let's generate fake data with the same number of rows as our real dataset.
num_sample_rows = len ( real_data )
fake_data = fake_synthesizer . sample ( num_sample_rows )
fake_data [[ 'Year' , 'Month' , 'Day' , 'Time' ]] = fake_data [[ 'Year' , 'Month' , 'Day' , 'Time' ]] . astype ( int )
fake_data . head ()
User
Card
Year
Month
Day
Time
Amount
Use Chip
Merchant Name
Merchant City
Merchant State
Zip
MCC
Errors?
Is Fraud?
0
1451
1
1995
2
3
15
2718.94
Chip Transaction
3533508019517891227
Bradley Beach
NaN
1840
3405
NaN
No
1
429
2
2002
11
12
11
691.59
Online Transaction
6655367386832070574
Airville
Algeria
80016
7276
NaN
No
2
843
5
2002
10
28
2
4568.73
Online Transaction
5979176342209983088
Ore City
Uganda
97470
5311
NaN
Yes
3
1402
4
2015
5
20
12
2811.24
Swipe Transaction
-5140159998192253928
Arverne
AL
22935
5651
NaN
Yes
4
644
4
1992
7
23
8
2437.87
Online Transaction
686332575408057196
Bridgeport
NJ
7823
5499
NaN
Yes
Why does the fake data look uniform within each range?
By design. DayZSynthesizer is built for cases where you have no real data to learn from — only schema and ranges — so it samples uniformly within the bounds you provide. Its only configuration knobs are bounds ( set_numerical_bounds , set_datetime_bounds ), categories ( set_category_values ), null proportions ( set_missing_values ), and rounding ( set_rounding_scheme ); there is no distribution-shaping parameter. To capture real-world shapes (skews, modes, correlations), you need a synthesizer that learns from real data — which is exactly what GaussianCopulaSynthesizer does next.
Generating Synthetic Data with GaussianCopulaSynthesizer ¶
The GaussianCopulaSynthesizer learns the statistical distributions and correlations in the real data, then generates new rows that preserve those patterns. This is the key difference from fake data — synthetic data is modeled , not randomized.
Notice that we don't need any of the manual configuration we did for DayZ — no category lists, no numerical bounds, no datetime ranges. The synthesizer learns all of these properties directly from the data.
from sdv.single_table import GaussianCopulaSynthesizer
synthetic_synthesizer = GaussianCopulaSynthesizer ( metadata )
synthetic_synthesizer . fit ( real_data )
synthetic_data = synthetic_synthesizer . sample ( num_sample_rows )
synthetic_data . head ()
User
Card
Year
Month
Day
Time
Amount
Use Chip
Merchant Name
Merchant City
Merchant State
Zip
MCC
Errors?
Is Fraud?
0
1594
1
2012
9
5
14
35.35
Swipe Transaction
-4500542936415012428
Breaux Bridge
IL
2145
5499
NaN
No
1
1798
0
2012
7
2
12
134.92
Swipe Transaction
-4334232547381218591
Pensacola
OH
97504
5499
NaN
No
2
332
2
2011
2
13
7
36

[truncated]

## Original Extract

Real Vs Synthetic Vs Fake Data — a step-by-step DataCebo cookbook for generating synthetic data with the Synthetic Data Vault (SDV).

Real Vs Synthetic Vs Fake Data | DataCebo Cookbooks cookbooks
About API Docs Source Back May 6, 2026
Benefits of Performance Testing with AI-Generated Synthetic Data ¶
Software systems need close-to-real-world data to test against. Teams typically choose between three approaches:
In this tutorial, we'll generate fake data with the DayZSynthesizer and synthetic data with the GaussianCopulaSynthesizer , then compare both against the real data — first statistically, then through a simulated fraud detection scenario.
When to use this cookbook: You're evaluating whether to use fake or synthetic data in place of production data — for example, when populating a dev/test environment or stress-testing a downstream service like a fraud detection pipeline.
Prerequisite: DayZSynthesizer is part of SDV Enterprise. The GaussianCopulaSynthesizer used for the synthetic-data sections is available in SDV Community.
import warnings
warnings . filterwarnings ( 'ignore' )
Loading and Exploring the Data ¶
We'll work with the credit_card_transactions dataset. Here's an overview of its columns:
User : user ID linking transactions together
Card : card issuer, encoded as an integer
Year : year the transaction occurred
Month : month the transaction occurred
Day : day the transaction occurred
Time : time the transaction occurred
Use Chip : if an EMV chip card was used, or if the transaction was online
Merchant Name : name of the merchant
Merchant City : registered city for the merchant
Merchant State : registered state for the merchant
Zip : registered zip code for the merchant
MCC : merchant category code, a 4-digit number identifying the type of transaction
Errors? : any errors experienced as part of the transaction
Is Fraud? : whether the transaction was ultimately fraudulent
from sdv.datasets.demo import download_demo
datasets , full_metadata = download_demo (
modality = 'multi_table' ,
dataset_name = 'credit_card_transactions'
)
full = datasets [ 'credit_card_transactions-ibm_v2' ]
full [ "Amount" ] = full [ "Amount" ] . str . replace ( "$" , "" ) . astype ( float )
full [ "Zip" ] = full [ "Zip" ] . astype ( str ) . str . replace ( ".0" , "" , regex = False )
full [ "Time" ] = full [ "Time" ] . str . split ( ":" ) . str [ 0 ] . astype ( int )
real_data = full . sample ( frac = 0.1 , random_state = 0 ) . reset_index ( drop = True )
del datasets , full
metadata = full_metadata . tables [ 'credit_card_transactions-ibm_v2' ]
for col , sdtype in [
( "User" , "categorical" ),
( "Card" , "categorical" ),
( "Amount" , "numerical" ),
( "Merchant City" , "categorical" ),
( "Merchant State" , "categorical" ),
( "Time" , "numerical" ),
]:
metadata . update_column ( column_name = col , sdtype = sdtype )
Let's preview the data:
real_data . head ()
User
Card
Year
Month
Day
Time
Amount
Use Chip
Merchant Name
Merchant City
Merchant State
Zip
MCC
Errors?
Is Fraud?
0
1264
0
2015
10
20
19
59.63
Swipe Transaction
2301046522991844568
Westlake
OH
44145
5813
NaN
No
1
1693
3
2011
8
19
17
41.63
Swipe Transaction
6661973303171003879
North Hollywood
CA
91605
5211
NaN
No
2
763
2
2009
2
10
17
16.84
Swipe Transaction
-8288218236188507426
Pomeroy
OH
45769
5812
NaN
No
3
1133
1
2006
3
16
13
13.32
Swipe Transaction
-3150043395187512028
Edinburgh
IN
46124
5921
NaN
No
4
550
0
2020
2
25
3
8.90
Chip Transaction
-337314228625997565
Parkton
NC
28371
5921
NaN
No
Now let's explore the dataset to understand its key properties — how many transactions and users it contains, where the values fall, and how the categorical columns relate. Each of the following questions is designed to surface something important about the data that the synthesizers will need to reproduce.
How large is the dataset and how many unique users does it have?
print ( f 'Sample rows: { len ( real_data ) : , } ' )
print ( f 'Columns: { real_data . shape [ 1 ] } ' )
print ( f 'Unique users: { real_data [ "User" ] . nunique () : , } ' )
Sample rows: 2,438,690
Columns: 15
Unique users: 2,000
The dataset is large enough to exercise statistical modeling — over a million transactions across a couple thousand unique users. That gives the synthesizer plenty of signal to learn distributions and correlations from, and gives us enough rows to see meaningful differences between real, fake, and synthetic data downstream.
Which columns have missing values?
# Check for missing values
missing = real_data . isnull () . sum ()
missing [ missing > 0 ]
Merchant State 272248
Errors? 2399915
dtype: int64
Only Merchant State and Errors? have missing values. We'll configure our fake data generator to match these proportions later.
What does the distribution of transaction amounts look like?
The vast majority of transactions are small — clustered near \$0 — with a long right tail extending past \$1,500. This heavy skew is typical of real-world spending behavior, where everyday purchases dominate and large transactions are rare.
Breaking it down into ranges reveals more detail:
Negative amounts represent refunds or chargebacks — they peak around -\$75. The \$0–\$250 range shows a steep decline from very small purchases, with periodic spikes at round-dollar amounts. Transactions above \$250 are sparse, forming a long tail out to \$1,500+.
What years are represented in this dataset?
Transaction volume grows steadily from the 2000s through 2020, reflecting increasing credit card adoption over time. The drop visible in 2020 reflects that the dataset only covers the first few months of that year (as the dataset was created at the beginning of 2020), not a real decrease in transaction volume.
How balanced is the fraud label?
The fraud label is highly imbalanced — only a tiny fraction of transactions are flagged as fraud. This is realistic: real-world fraud rates are typically below 1% of transactions. The imbalance matters for any downstream modeling task: a naive model that always predicts "not fraud" would already be ~99% accurate, so we'll need to be careful about how we evaluate performance later.
How do the key categorical columns relate to each other?
The chart below shows the flow between payment method, transaction errors, and fraud status.
The Sankey chart shows how transactions flow across payment method, error status, and fraud label. The thick band on the right indicates that most transactions have no recorded errors — the Errors? column is blank (i.e. null) for the vast majority of rows. The smaller bands break out the rare error types: bad PIN, insufficient balance, technical glitches, etc. Fraud is rare across all payment methods and is concentrated among transactions with no recorded error.
Generating Fake Data with DayZSynthesizer ¶
The DayZSynthesizer generates fake data by randomly sampling values from user-defined ranges and categories. Unlike a statistical model, it does not learn any patterns or correlations from the real data.
To make the fake data as realistic as possible, we configure it step by step with information from the real dataset.
First, we instantiate a DayZSynthesizer from the metadata. The synthesizer starts in a default state where it knows the column types but nothing about the actual values, ranges, or null rates in the real data — those are the things we configure in the next steps.
from sdv.single_table import DayZSynthesizer
fake_synthesizer = DayZSynthesizer ( metadata )
Next, we tell the synthesizer the missing-value proportions for each column that has nulls. By default, DayZSynthesizer assumes every column is fully populated, which would produce fake data without any nulls. Real credit card data has nulls in Merchant State and Errors? , so we explicitly set those proportions to match — otherwise the fake data would look unrealistically complete.
# Set missing value proportions to match real data
for col in [ 'Merchant State' , 'Errors?' ]:
fake_synthesizer . set_missing_values (
column_name = col ,
missing_values_proportion = round (
real_data [ col ] . isnull () . sum () / len ( real_data ), 10
)
)
Next, we tell the synthesizer the valid categories for each categorical column. Without this, the synthesizer doesn't know what values are legal for Merchant Name , Use Chip , etc., and would either fail to generate them or produce placeholders. We pass in the unique values from the real data so the fake data uses the same vocabulary.
cat_columns = [
'User' , 'Card' , 'Use Chip' , 'Merchant Name' ,
'Merchant City' , 'Merchant State' , 'Zip' , 'MCC' ,
'Errors?' , 'Is Fraud?'
]
for col in cat_columns :
fake_synthesizer . set_category_values (
column_name = col ,
category_values = real_data [ col ] . dropna () . unique () . tolist ()
)
Finally, we set the numerical and datetime bounds for each non-categorical column. Without these, the synthesizer has no idea what range to sample from for Amount , Year , Month , or Day . We use the real data's min/max for each — this gives the fake data realistic ranges, even though the values within those ranges are sampled uniformly at random.
for col in [ 'Amount' , 'Year' , 'Month' , 'Day' , 'Time' ]:
cast = float if col == 'Amount' else int
fake_synthesizer . set_numerical_bounds (
column_name = col ,
min_value = cast ( real_data [ col ] . min ()),
max_value = cast ( real_data [ col ] . max ())
)
if col != 'Amount' :
fake_synthesizer . set_rounding_scheme (
column_name = col ,
num_decimal_digits = 0
)
Now let's generate fake data with the same number of rows as our real dataset.
num_sample_rows = len ( real_data )
fake_data = fake_synthesizer . sample ( num_sample_rows )
fake_data [[ 'Year' , 'Month' , 'Day' , 'Time' ]] = fake_data [[ 'Year' , 'Month' , 'Day' , 'Time' ]] . astype ( int )
fake_data . head ()
User
Card
Year
Month
Day
Time
Amount
Use Chip
Merchant Name
Merchant City
Merchant State
Zip
MCC
Errors?
Is Fraud?
0
1451
1
1995
2
3
15
2718.94
Chip Transaction
3533508019517891227
Bradley Beach
NaN
1840
3405
NaN
No
1
429
2
2002
11
12
11
691.59
Online Transaction
6655367386832070574
Airville
Algeria
80016
7276
NaN
No
2
843
5
2002
10
28
2
4568.73
Online Transaction
5979176342209983088
Ore City
Uganda
97470
5311
NaN
Yes
3
1402
4
2015
5
20
12
2811.24
Swipe Transaction
-5140159998192253928
Arverne
AL
22935
5651
NaN
Yes
4
644
4
1992
7
23
8
2437.87
Online Transaction
686332575408057196
Bridgeport
NJ
7823
5499
NaN
Yes
Why does the fake data look uniform within each range?
By design. DayZSynthesizer is built for cases where you have no real data to learn from — only schema and ranges — so it samples uniformly within the bounds you provide. Its only configuration knobs are bounds ( set_numerical_bounds , set_datetime_bounds ), categories ( set_category_values ), null proportions ( set_missing_values ), and rounding ( set_rounding_scheme ); there is no distribution-shaping parameter. To capture real-world shapes (skews, modes, correlations), you need a synthesizer that learns from real data — which is exactly what GaussianCopulaSynthesizer does next.
Generating Synthetic Data with GaussianCopulaSynthesizer ¶
The GaussianCopulaSynthesizer learns the statistical distributions and correlations in the real data, then generates new rows that preserve those patterns. This is the key difference from fake data — synthetic data is modeled , not randomized.
Notice that we don't need any of the manual configuration we did for DayZ — no category lists, no numerical bounds, no datetime ranges. The synthesizer learns all of these properties directly from the data.
from sdv.single_table import GaussianCopulaSynthesizer
synthetic_synthesizer = GaussianCopulaSynthesizer ( metadata )
synthetic_synthesizer . fit ( real_data )
synthetic_data = synthetic_synthesizer . sample ( num_sample_rows )
synthetic_data . head ()
User
Card
Year
Month
Day
Time
Amount
Use Chip
Merchant Name
Merchant City
Merchant State
Zip
MCC
Errors?
Is Fraud?
0
1594
1
2012
9
5
14
35.35
Swipe Transaction
-4500542936415012428
Breaux Bridge
IL
2145
5499
NaN
No
1
1798
0
2012
7
2
12
134.92
Swipe Transaction
-4334232547381218591
Pensacola
OH
97504
5499
NaN
No
2
332
2
2011
2
13
7
36

[truncated]
