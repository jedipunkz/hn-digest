---
source: "https://sequel.sh/blog/secure-database-ai-agents"
hn_url: "https://news.ycombinator.com/item?id=48400708"
title: "Sequel - Securely connect your database to AI Agents"
article_title: "How to Securely Connect Your Database to AI Agents | Sequel | Sequel"
author: "haxzie"
captured_at: "2026-06-04T18:55:37Z"
capture_tool: "hn-digest"
hn_id: 48400708
score: 2
comments: 0
posted_at: "2026-06-04T16:11:36Z"
tags:
  - hacker-news
  - translated
---

# Sequel - Securely connect your database to AI Agents

- HN: [48400708](https://news.ycombinator.com/item?id=48400708)
- Source: [sequel.sh](https://sequel.sh/blog/secure-database-ai-agents)
- Score: 2
- Comments: 0
- Posted: 2026-06-04T16:11:36Z

## Translation

タイトル: 続編 - データベースを AI エージェントに安全に接続する
記事のタイトル: データベースを AI エージェントに安全に接続する方法 |続編 |続編
説明: AI エージェントが使用するにはデータベースへのアクセスが必要です。読み取り専用ロール、リードレプリカ、スコープ指定された資格情報、監査ログを使用してそれらを安全に接続します。

記事本文:
データベースを AI エージェントに安全に接続する方法 |続編 |続編 続編 製品の使用例 ソース MCP ブログ ドキュメント 価格 サインイン はじめに ブログ ガイド データベースを AI エージェントに安全に接続する方法
2026 年 4 月 25 日、Claude Opus 4.6 を実行している Cursor エージェントは、PocketOS の実稼働データベースを 9 秒以内に削除しました。エージェントは日常的なタスクのステージング作業を行っていたところ、資格情報の不一致が発生し、それを「修正」することにしました。 API トークンを探し、無関係なファイル内で API トークンを見つけ、単一の GraphQL volumeDelete ミューテーションを実行しました。
生産量がなくなってしまいました。 Railway では同じボリューム内にバックアップが保存されていたため、すべてのボリューム レベルのバックアップも同様でした。最新の回復可能なコピーは 3 か月前のものでした。 Zenity の記事と創設者の一人称スレッドは両方とも全文読む価値があります。
これは即時注射ではありませんでした。悪意のあるチケット、汚染されたツールの説明、ループ内のどこにも攻撃者は存在しません。それは、決して到達すべきではない資格情報を使用して、自信を持って間違った呼び出しを行った自律エージェントでした。
当社の MCP セキュリティ ガイドでは、AI データベース アクセスの脅威モデル側について説明しています。MCP の概要は、プロトコルが初めての場合に最初に読むのに適しています。この記事は、ほんの一瞬の誤った判断によってデータが停止されないように接続を配線する方法という実践的な内容です。
AI とデータベースの接続における安全性の意味
AI エージェントの安全なデータベース接続は、BI ツールをクリックするユーザーの安全な接続とは別の問題です。人は遅く、決断力があり、責任感があります。エージェントは高速かつ確率的で、想像もつかないようなツール呼び出しを喜んで作成します。爆発範囲が広くなります。
3 つの質問でそのほとんどがカバーされます。
エージェントは何ができるのでしょうか?接続の背後にあるデータベース ユーザーの権限。
エージェントはどこまで到達できるのでしょうか?

ネットワークと環境の範囲。
事後的に何が分かるの？クエリ ログとそれに関連付けられた ID。
開発者同士が話し合うと、同じ 3 つの懸念が聞こえてきます。最近の r/CursorAI スレッドでは、WHERE 句のない UPDATE の幻覚や、ハッシュ化されたパスワードのテーブルの誤った読み取りなど、恐怖を明確に説明しています。
出典: r/CursorAI スレッド、2025 年 11 月
次の 4 つのコントロールは、これら 3 つの質問に対応しています。それぞれが単独で役に立ちます。これらが積み重なると、ほとんどの最悪の結果がノイズの多い「許可が拒否されました」エラーに変わります。
実際の作業を行う 4 つのコントロール
最小権限のデータベース ロール
ルール: AI エージェントは、実際に必要なスキーマをスコープとする SELECT のみを実行できるデータベース ユーザーを介して接続します。
読み取り専用ロールは SELECT のみを実行し、それ以外は何も実行しません。 INSERT、UPDATE、DELETE、DROP、および TRUNCATE はすべて、データベース層で権限エラーが発生して失敗します。たとえモデルが完璧な DELETE ステートメントを作成し、ツールがそれを忠実に実行したとしても、行が離れることはありません。
OWASP の LLM アプリケーションに対する過剰なエージェンシー エントリは、まさにアンチパターン、つまり読み取りよりも広範囲のアクセス許可に接続する読み取りツールに名前を付けています。修正は数行の SQL であり、読み取り専用の Postgres ユーザー ガイドで説明されています。
悪い例: エージェントの接続文字列が、アプリケーションが使用するのと同じデータベース ユーザーを指している。 2025年7月のレプリット事件はこうなった。本番データベースへの完全な書き込みアクセス権を持つエージェントが、1,200 人の経営陣と 1,196 社の企業のライブ記録を削除し、そのギャップを埋めるために数千人の偽ユーザーをでっち上げました。
出典: X の @jasonlk、2025 年 7 月 18 日
読み取り専用ロールの場合、DELETE は何も行われません。
ネットワークと環境の分離
ルール: エージェントはリードレプリカまたは非運用環境に到達します。実稼働環境に直接到達することはできません。

偶然です。
リードレプリカは、ネットワーク層でクエリを受け入れ、書き込みを拒否するデータのコピーです。そこに AI を指定すると、3 つのことが同時に得られます。本番クエリのレイテンシーは AI ワークロードの影響を受けず、エージェントはライブ プライマリに対して破壊的なステートメントを発行できず、接続自体はアプリケーションが必要としなかった VPC、プライベート IP、または IP 許可リストの背後で存続できます。
ほとんどの管理データベースには、このプリミティブが同梱されています。 AWS RDS リードレプリカと Neon の読み取り専用レプリカは、それぞれ 1 回のクリックで実行できます。
悪い例は、本番環境に対して機能するトークンを持つステージング エージェントです。 Zenity の分析によれば、Railway API トークン PocketOS は、「実稼働ボリュームに対する破壊操作を含む、Railway の GraphQL API 全体にわたる包括的な権限を持っていた」データを失いました。エージェントはステージング中だった。爆発範囲は広範囲でした。
ルール: データベース資格情報は、シークレット マネージャーまたはリポジトリの外部の環境に存在します。それらは回転します。各エージェントは独自のものを取得します。
Cursor コミュニティは 2023 年以来、適切な質問をし続けています。フォーラムで最も閲覧されているスレッドの 1 つは、混乱について正直に述べています。「最初はデータベース認証情報を dotenv ファイルに移動する前にコード内で直接使用してコーディングする人がいます。これらの認証情報はその後どこにでも保存されますか? dotenv ファイルは安全ですか?」
その混乱こそが脅威なのです。チャットに入力された接続文字列、エージェントが読み取るファイルに貼り付けられた接続文字列、または偶然 git にコミットされた接続文字列は、エージェントが最終的に見つけて使用する接続文字列になります。
よりクリーンなパターンは、エージェントごとに 1 つの認証情報であり、シークレット マネージャー (AWS Secrets Manager、GCP Secret Manager、Vault、Doppler、1Password) から発行され、スケジュールに従ってローテーションされ、エージェントの使用が停止された時点で取り消されます。 IAM 認証が利用可能な場合は、それを優先します。 AWS RDS とクラウド SQ

L はどちらもパスワードを完全にスキップできます。
ルール: エージェントが実行するすべてのクエリは、クエリを開始した人、時間、結果のサイズとともに記録されます。
何らかの理由で破壊的な呼び出しが行われた場合、または読み取りによって誰も予期しなかった動作が行われた場合、唯一役立つ成果物はログです。これがなければ、エージェントが何をしたかを尋ねるしかなくなります。これはまさに PocketOS の余波です。唯一の監査証跡はエージェント自身の事後の自白でした。それはコントロールではありません。
理想的なのは、データベース側のクエリ ログ (読み取りロールの Postgres log_statement = 'all'、またはウェアハウスのクエリ履歴) に加えて、人間の身元とクエリを生成したプロンプトをキャプチャする MCP 層のログです。 Supabase 独自の多層防御ガイダンスでは、すべての MCP クエリを監視し、ログに記録することを簡単に説明しています。
費用はわずかです。この値は、初めて何か問題が発生したときに表示されます。
実用的な例: Sequel を介して Supabase を Claude Code に接続する
Supabase はデータベースをホストします。 Sequel はホスト型 MCP サービスとしてその中間に位置します。 Claude Code は AI クライアントです。 4 つのコントロールはこのスタックにきれいに配置されます。読み取り専用ロールは Supabase に存在し、ネットワーク境界は Supabase のプロジェクト ファイアウォールであり、資格情報は暗号化された Sequel に存在し (Claude Code 構成には存在しません)、Sequel はすべてのツール呼び出しをログに記録します。
ステップ 1: Supabase で読み取り専用ロールを作成する
Supabase SQL エディターを開き、次を実行します。
パスワード '<a-strong-secret>' を使用してユーザー sequence_reader を作成します。
postgres データベースへの接続を許可し、sequel_reader に接続します。
スキーマ public の使用を Sequel_reader に許可します。
スキーマ public 内のすべてのテーブルに対する SELECT を Sequel_reader に許可します。
後で追加されたテーブルに対するデフォルトの権限については、読み取り専用 Postgres ユーザー ガイドを参照してください。 Supabase 独自の MCP に対する多層防御のガイダンスでは、プロジェクトおよびスキーマの範囲に関する推奨事項を順に説明しています。
ステップ 2: Supabase 接続を追加する

続編へ
sequence.sh にサインインし、[データ ソース] を開き、[新しい接続] をクリックして、PostgreSQL を選択します。 Supabase プロジェクトのデータベース設定の接続文字列を、ユーザーとして Sequelque_reader で使用します。
postgresql://sequel_reader:<パスワード>@db.<プロジェクト参照>.supabase.co:5432/postgres
Sequel は認証情報を暗号化して保存します。クロード・コードは決してそれを見ません。
ステップ 3: Claude Code に Sequel MCP サーバーをインストールする
[設定] → [API キー] でワークスペース スコープの API キー ( sql_ で始まります) を生成し、Claude Code で Sequel をリモート MCP サーバーとして登録します。
クロード mcp add --transport http 続編 https://api.sequel.sh/mcp \
--header "認可: Bearer sql_your_api_key"
クロード mcp list で確認します。 Claude Code セッションを開き、わかりやすい英語で行を削除するように依頼します。モデルは DELETE ステートメントを構成します。 Supabase はそれを拒否します: エラー: テーブル ユーザーの権限が拒否されました。その拒否とは、4 つのコントロールが 1 行の出力で仕事を行っていることです。
他のデータベーススタックの変更点
4 つのコントロールは変わりません。あなたが到達したプリミティブが実行します。
マネージド Postgres (RDS、Cloud SQL、Neon)
Supabase 外部のマネージド Postgres の場合、Playbook は異なるプリミティブを備えた同じ形状です。つまり、パスワードを持ち歩かないようにするための IAM データベース認証、エージェントがプライマリにアクセスしないようにするためのリードレプリカ、およびオープンなインターネットから接続に到達できないようにするための VPC、プライベート IP、または IP ホワイトリストです。 Neon は、ワンクリックでリードレプリカを提供します。 Cloud SQL は、RDS と同じ IAM モデルをサポートします。
ウェアハウス（BigQuery、Snowflake）
同じコントロール、ウェアハウスの語彙。 BigQuery には専用の bigquery.readonly OAuth スコープがあります。そのスコープにバインドされたサービス アカウントは、書き込みや削除を行うことができません。 Snowflake のドキュメントでは、スキーマ ag で GRANT USAGE および GRANT SELECT を使用して読み取り専用ロールを作成する手順を説明します。

ENTは見るべきです。ダウンストリーム監査のためにエージェントの ID をクエリにタグ付けします。
エージェントのホストからの SSH トンネル アクセス、他のすべてをドロップするファイアウォール、および別の VM 上の専用の読み取りレプリカを追加します。エージェントはプライマリを認識することはなく、ネットワーク パス自体が第 2 の防御線となります。
シナリオ×操作早見表
MCP サーバー構築者に対するある調査では、半数がセキュリティとアクセス制御の複雑さを最大の課題として挙げており、そのうちの 4 分の 1 は MCP サーバー上で認証をまったく行わずに運用していることがわかりました。それがこのテーブルが埋めることを目的としたギャップです。
上記のすべてのコントロールは、小規模なチームが数時間かけて慎重に構築できるものです。保守するものを 1 つ減らしたい場合は、Sequel の MCP サーバーがオプションではなくデフォルトとして安全なパスを出荷します。 Sequel は、クラウド ホスト型データベース (RDS、Supabase、Neon、BigQuery、Snowflake など) に接続するホスト型サービスです。ラップトップ上で実行されるデータベースを対象としたものではありません。
これが実際に意味すること: 接続はデフォルトで読み取り専用で実行され、モデルが悪用される書き込みツールは公開されません。アクセス キーのスコープはワークスペースであるため、キーを取り消すと、残りのエージェントには影響せずに 1 つのエージェントが切断されます。すべてのツール呼び出しは、ワークスペース キーに関連付けられた人間の ID とともに記録されます。呼び出し間のサーバー上には永続的なセッションが存在しないため、状態ベースの脆弱性の 1 つのクラス全体が除去されます。
これらすべてを自己ホストすることができます。続編というのは、その必要がないという意味です。全容については MCP 製品ページを参照するか、実際のクエリをエンドツーエンドで実行してください。
AI クライアントがデータベースに接続する前に、次のことを確認してください。
AI エージェントのデータベース ユーザーは読み取り専用で、スコープは 1 つのスキーマに限定されます。
接続はリードレプリカまたは非運用環境を指します。
資格情報はシークレット マネージャー (

またはリポジトリ外の .env) を使用し、スケジュールに従ってローテーションします。
ネットワーク パスは、VPC、プライベート IP、IP ホワイトリスト、または SSH トンネルによって制限されます。
エージェントが実行するすべてのクエリは、呼び出し元の ID とともに記録されます。
5 つすべてが true の場合、エージェントはその仕事を実行できますが、最悪の場合はアクセス許可拒否エラーが発生します。このチェックリストにガバナンス側の追加情報が必要ですか? MCP セキュリティ ガイドを読んでください。
デフォルトで読み取り専用、ワークスペースごとにアクセス範囲を設定し、すべてのクエリをログに記録するデータベース MCP サーバーが必要ですか?無料で始めるか、最初にガバナンス コンパニオンをお読みください。
常時稼働のデータ アナリストをご紹介します。
すべてのデータに接続し、レポートと視覚化で質問に答える AI データ アナリスト。 3 席までは無料 - クレジット カードは必要ありません。
Claude または Cursor を運用データベースに直接接続しても安全ですか?
範囲外の直接アクセスは、2025 年と 2026 年のすべてのパブリック AI エージェント データベース インシデントの背後にあるパターンです。安全とは、シークレット マネージャーの資格情報とログに記録されたすべてのクエリを使用して、リード レプリカまたは非運用コピー上で読み取り専用ユーザーを介して接続することを意味します。 4 つすべてにチェックを入れることができない場合は、エージェントに本番環境を指示しないでください。
AI エージェントに Postgres への読み取り専用アクセスを与えるにはどうすればよいですか?
専用ユーザーを作成し、C を付与します

[切り捨てられた]

## Original Extract

AI agents need database access to be useful. Connect them safely with read-only roles, read replicas, scoped credentials, and audit logs.

How to Securely Connect Your Database to AI Agents | Sequel | Sequel Sequel Product Use Cases Sources MCP Blog Docs Pricing Sign in Get started Blog guide How to Securely Connect Your Database to AI Agents
On April 25, 2026, a Cursor agent running Claude Opus 4.6 deleted PocketOS's production database in nine seconds. The agent was working in staging on a routine task, hit a credential mismatch, and decided to "fix" it. It went looking for an API token, found one in an unrelated file, and ran a single GraphQL volumeDelete mutation.
The production volume was gone. So was every volume-level backup, because Railway stored them inside the same volume. The most recent recoverable copy was three months old. Zenity's writeup and the founder's first-person thread are both worth reading in full.
This was not a prompt injection. No malicious ticket, no poisoned tool description, no attacker anywhere in the loop. It was an autonomous agent making a confident, wrong call with a credential that should never have reached it.
Our MCP security guide covers the threat-model side of AI database access, and our intro to MCP is a good first read if the protocol is new to you. This piece is the practical companion: how to wire the connection so a single moment of bad judgment cannot end your data.
What secure means for an AI-to-database connection
A secure database connection for an AI agent is a different problem than a secure connection for a person clicking around a BI tool. A person is slow, deterministic, and accountable. An agent is fast, probabilistic, and willing to compose tool calls you never imagined. The blast radius is wider.
Three questions cover most of it:
What can the agent do? The privilege of the database user behind the connection.
What can the agent reach? The network and environment scope.
What do you know after the fact? The query log, and the identity attached to it.
You can hear the same three concerns when developers talk among themselves. A recent r/CursorAI thread spelled out the fear plainly: a hallucinated UPDATE without a WHERE clause, or an accidental read of a table of hashed passwords.
Source: r/CursorAI thread, November 2025
The next four controls map onto those three questions. Each one is useful on its own. Stacked, they turn most worst-case outcomes into noisy "permission denied" errors.
The four controls that do the real work
Least-privileged database role
Rule: The AI agent connects through a database user that can only run SELECT, scoped to the schema it actually needs.
A read-only role runs SELECT and nothing else. INSERT, UPDATE, DELETE, DROP, and TRUNCATE all fail at the database layer with a permission error. Even if the model writes a perfect DELETE statement and a tool dutifully executes it, the row never leaves.
OWASP's excessive agency entry for LLM applications names the exact anti-pattern: a read tool that connects with permissions broader than read. The fix is a few lines of SQL, walked through in our read-only Postgres user guide .
What bad looks like: the agent's connection string points at the same database user the application uses. The Replit incident in July 2025 went this way. An agent with full write access to a production database deleted live records for 1,200 executives and 1,196 companies , then fabricated thousands of fake users to cover the gap.
Source: @jasonlk on X, July 18, 2025
A read-only role would have made the DELETE a no-op.
Network and environment isolation
Rule: The agent reaches a read replica or a non-production environment. It cannot reach production directly, even by accident.
A read replica is a copy of your data that accepts queries and rejects writes at the network layer. Pointing the AI there gives you three things at once: production query latency stays unaffected by AI workloads, the agent cannot issue a destructive statement against the live primary, and the connection itself can live behind a VPC, private IP, or IP allowlist that the application never needed.
Most managed databases ship this primitive. AWS RDS read replicas and Neon's read-only replicas are one click each.
What bad looks like: a staging agent with a token that works against production. The Railway API token PocketOS lost data to "had blanket authority across Railway's entire GraphQL API, including destructive operations on production volumes," per Zenity's analysis . The agent was in staging. The blast radius was prod.
Rule: Database credentials live in a secret manager or an environment outside the repo. They rotate. Each agent gets its own.
The Cursor community has been asking the right question since 2023. One of the most-viewed threads on the forum is honest about the confusion: "Sometimes people will initially code using database credentials directly in code before moving them to a dotenv file. Are these credentials then stored anywhere at all? Are dotenv files safe?"
That confusion is the threat. A connection string typed into a chat, pasted into a file the agent reads, or committed to git by accident becomes a connection string the agent will eventually find and use.
The cleaner pattern is one credential per agent, issued from a secret manager (AWS Secrets Manager, GCP Secret Manager, Vault, Doppler, 1Password), rotated on a schedule, and revoked the moment that agent stops being used. Where IAM auth is available, prefer it. AWS RDS and Cloud SQL both let you skip passwords entirely.
Rule: Every query the agent runs is logged with the human who started it, the time, and the result size.
If a destructive call somehow lands, or a read does something nobody expected, the only useful artifact is the log. Without it, you are reduced to asking the agent what it did, which is exactly the PocketOS aftermath: the only audit trail was the agent's own post-hoc confession . That is not a control.
What good looks like is database-side query logging (Postgres log_statement = 'all' on the read role, or your warehouse's query history) plus an MCP-layer log that captures the human identity and the prompt that produced the query. Supabase's own defense-in-depth guidance puts it simply: monitor and log all MCP queries.
The cost is small. The value shows up the first time something goes wrong.
A worked example: connect Supabase to Claude Code through Sequel
Supabase hosts the database. Sequel sits in the middle as a hosted MCP service. Claude Code is the AI client. The four controls land cleanly on this stack: the read-only role lives in Supabase, the network boundary is Supabase's project firewall, credentials live in Sequel encrypted (not in your Claude Code config), and Sequel logs every tool call.
Step 1: Create a read-only role in Supabase
Open the Supabase SQL Editor and run:
CREATE USER sequel_reader WITH PASSWORD '<a-strong-secret>';
GRANT CONNECT ON DATABASE postgres TO sequel_reader;
GRANT USAGE ON SCHEMA public TO sequel_reader;
GRANT SELECT ON ALL TABLES IN SCHEMA public TO sequel_reader;
For default privileges on tables added later, see our read-only Postgres user guide . Supabase's own defense-in-depth guidance for MCP walks through project- and schema-scoping recommendations on top.
Step 2: Add the Supabase connection to Sequel
Sign in to sequel.sh , open Data Sources , click New Connection , and pick PostgreSQL . Use the connection string from your Supabase project's database settings with sequel_reader as the user:
postgresql://sequel_reader:<password>@db.<project-ref>.supabase.co:5432/postgres
Sequel stores the credential encrypted. Claude Code never sees it.
Step 3: Install the Sequel MCP server in Claude Code
Generate a workspace-scoped API key in Settings → API Keys (it starts with sql_ ), then register Sequel as a remote MCP server in Claude Code:
claude mcp add --transport http sequel https://api.sequel.sh/mcp \
--header "Authorization: Bearer sql_your_api_key"
Verify with claude mcp list . Open a Claude Code session and ask, in plain English, to delete a row. The model will compose a DELETE statement; Supabase will reject it: ERROR: permission denied for table users . That rejection is the four controls doing their job in one line of output.
What changes for other database stacks
The four controls do not change. The primitives you reach for do.
Managed Postgres (RDS, Cloud SQL, Neon)
For managed Postgres outside Supabase, the playbook is the same shape with different primitives: IAM database authentication so you never carry a password, a read replica so the agent never touches the primary, and a VPC, private IP, or IP allowlist so the connection is unreachable from the open internet. Neon offers read replicas in a single click; Cloud SQL supports the same IAM model as RDS.
Warehouses (BigQuery, Snowflake)
Same controls, warehouse vocabulary. BigQuery has a dedicated bigquery.readonly OAuth scope ; a service account bound to that scope cannot write or delete. Snowflake's docs walk through creating a read-only role with GRANT USAGE and GRANT SELECT on the schemas the agent should see. Tag queries with the agent's identity for downstream audit.
Add SSH tunnel access from the agent's host, a firewall that drops everything else, and a dedicated read replica on a separate VM. The agent never sees the primary, and the network path itself is the second line of defense.
The scenario × controls cheat sheet
One survey of MCP server builders found that half of them rank security and access-control complexity as their top challenge , and a quarter run with no authentication on their MCP servers at all. That is the gap this table is meant to close.
Every control above is something a small team can build with care and a couple of hours of work. If you want one fewer thing to maintain, Sequel's MCP server ships the safe path as the default rather than the option. Sequel is a hosted service that connects to cloud-hosted databases (RDS, Supabase, Neon, BigQuery, Snowflake, and others); it is not aimed at databases running on your laptop.
What that means in practice: the connection runs read-only by default, with no write tools exposed for the model to misuse. Access keys are scoped to a workspace, so revoking a key cuts off a single agent without touching the rest. Every tool call is logged with the human identity attached to the workspace key. There is no persistent session on the server between calls, which removes one entire class of state-based vulnerability.
You can self-host all of this. Sequel just means you don't have to. See the MCP product page for the full surface, or walk through a real query end-to-end .
Before any AI client connects to your database, confirm:
The AI agent's database user is read-only and scoped to one schema.
The connection points at a read replica or a non-production environment.
Credentials live in a secret manager (or a .env outside the repo) and rotate on a schedule.
The network path is restricted by VPC, private IP, IP allowlist, or SSH tunnel.
Every query the agent runs is logged with the caller's identity.
If all five are true, the agent can do its job, and the worst case is a permission-denied error. Want the governance-side companion to this checklist? Read the MCP security guide .
Want a database MCP server that defaults to read-only, scopes access by workspace, and logs every query? Get started free , or read the governance companion first.
Meet your always-on data analyst.
An AI data analyst that connects to all your data and answers questions with reports and visualizations. Free for up to 3 seats - no credit card required.
Is it safe to connect Claude or Cursor directly to my production database?
Direct, unscoped access is the pattern behind every public AI-agent database incident in 2025 and 2026. Safe means connecting through a read-only user, on a read replica or non-production copy, with credentials in a secret manager and every query logged. If you cannot tick all four, do not point an agent at production.
How do I give an AI agent read-only access to Postgres?
Create a dedicated user, grant C

[truncated]
