---
source: "https://github.com/prathikmurugan/jsonbit"
hn_url: "https://news.ycombinator.com/item?id=50582997"
title: "Jsonbit – Reduce token usage when AI reads logs"
article_title: "GitHub - prathikmurugan/jsonbit · GitHub"
image: "https://opengraph.githubassets.com/f8ca48587b066bed1d99d381d2169bc6784cfa1ea40565dee3923f7a9f293118/prathikmurugan/jsonbit"
author: "ComputedCamp995"
captured_at: "2026-11-20T22:30:35Z"
capture_tool: "hn-digest"
hn_id: 50582997
score: 1
comments: 0
posted_at: "2026-11-20T22:08:26Z"
tags:
  - hacker-news
---

# Jsonbit – Reduce token usage when AI reads logs

- HN: [50582997](https://news.ycombinator.com/item?id=50582997)
- Source: [github.com](https://github.com/prathikmurugan/jsonbit)
- Score: 1
- Comments: 0
- Posted: 2026-11-20T22:08:26Z

## Translation

Title: Jsonbit – Reduce token usage when AI reads logs
Article title: GitHub - prathikmurugan/jsonbit · GitHub
Description: Contribute to prathikmurugan/jsonbit development by creating an account on GitHub.

Article text:
GitHub - prathikmurugan/jsonbit · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects, and automated work
Build & Ship Actions CI/CD and automation
Code Quality Quality analysis, fixes, and enforcement
Codespaces Cloud development environments
Secure Code Security Vulnerability detection and remediation
Secret Protection Secret scanning and leak prevention
Developers Learn Documentation
Changelog Changelog Your source for the latest features, improvements, and fixes across GitHub products, all in one place.
Businesses By Stage Enterprises
Add-ons GitHub Advanced Security Enterprise-grade code security
Copilot for business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Resources Learn & Connect Documentation
GitHub Universe GitHub Universe The event for developers, by developers — talks, workshops, and what’s next for GitHub.
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
prathikmurugan
jsonbit Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
13 Commits 13 Commits Folders and files
.github .github benchmarks benchmarks bin bin python/ jsonbit python/ jsonbit rust rust scripts scripts src src test test .gitignore .gitignore .mcp.json .mcp.json AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md opencode.json opencode.json package-lock.json package-lock.json package.json package.json pyproject.toml pyproject.toml tsconfig.json tsconfig.json View all files Repository files navigation
JsonBit compacts nested JSON before it reaches your LLM. Send a compact digest in one prompt, or let your agent query the original data locally via SQL/MCP and read only the results it needs.
pip install jsonbit
jsonbit digest path/to/your-file.json
Pass the output to your LLM in place of the raw JSON. (Rust cargo , Node.js npm / npx , and standalone binaries are listed under Installation & Setup .)
Example output from jsonbit digest :
# orders[4](orderId,customer_name,customer_tier,status,total)
@stats: total[sum=619.00,avg=154.75,min=95.50,max=201.00]
ORD-101,Alice,pro,shipped,171.75
ORD-102,Bob,starter,pending,201
ORD-103,Cara,pro,shipped,95.5
ORD-104,Dan,pro,shipped,150.75
# orders_items[6] -> orders (sku,qty,price)
@stats: qty[sum=15,avg=2.5,min=1(n=1),max=4] price[sum=246.25,avg=41.04,min=10.50,max=74.50]
0:USB-C,2,10.5|MOUSE,3,50.25
1:MOUSE,4,50.25
2:HUB-4,1,74.5|USB-C,2,10.5
3:MOUSE,3,50.25
Workflow
Command / MCP Tool
Best For
Fidelity & Guarantees
1-Turn Compaction
jsonbit digest / olap_digest
Single-prompt LLM calls & holistic summaries where the model needs every row.
Compact LLM text view containing all rows, escaped delimiters, hoisted child tables, and precomputed stats (rounded averages are approximate) .
Agentic Local Querying
jsonbit card + sql / olap_load + olap_sql
Coding agents (OpenCode, Claude Code, Cursor) querying medium/large .json or .jsonl files.
Emits a compact schema and statistics summary (independent of row count); runs exact read-only SQL & BitVec64 filters locally in microseconds.
Binary Storage
jsonbit pack / jsonbit rehydrate
Archiving or caching normalized .jsonbit snapshots on disk ( -80.6% vs. minified JSON on W1 – W10 ).
Lossless round-trip verified across all 33 benchmark/adversarial datasets and unit tests (preserving nulls, colliding keys, and u64 > i64::MAX ).
Table of Contents
Agent Integrations (OpenCode, Claude Code, Codex, Cursor)
Install the pre-built native jsonbit binary using your preferred package manager (no Rust compiler required for Python or Node.js):
Python ( pip / uv ) — Includes native CLI + Python SDK
pip install jsonbit
Rust ( cargo )
cargo install jsonbit
Node.js ( npm / npx ) — Includes native CLI + TypeScript SDK
npm install -g @jsonbit/core
# or run directly without global install:
npx @jsonbit/core --help
Pre-Built Binaries & Building from Source
Standalone binaries for Linux ( x64 , arm64 ), macOS ( arm64 , x64 ), and Windows ( x64 ) are attached to every GitHub Release .
To build from source (requires Rust 1.75+ ):
git clone https://github.com/prathikmurugan/jsonbit.git
cd jsonbit
cargo install --path .
Verify installation:
jsonbit --version
Agent Integrations
jsonbit includes a built-in stdio JSON-RPC 2.0 MCP server ( jsonbit mcp ) with zero external runtime dependencies. It exposes five tools: olap_load , olap_digest , olap_ask , olap_sql , and olap_rehydrate .
Add the following to opencode.json in your project root (or globally in ~/.config/opencode/opencode.json ):
{
"$schema" : " https://opencode.ai/config.json " ,
"mcp" : {
"jsonbit" : {
"type" : " local " ,
"command" : [ " jsonbit " , " mcp " ],
"enabled" : true
}
}
}
Claude Code ( .mcp.json or CLI)
Register via the Claude Code CLI:
claude mcp add jsonbit -- jsonbit mcp
Or include in .mcp.json :
{
"mcpServers" : {
"jsonbit" : {
"command" : " jsonbit " ,
"args" : [ " mcp " ]
}
}
}
Cursor & Windsurf
Add the .mcp.json configuration above to .cursor/mcp.json or ~/.codeium/windsurf/mcp_config.json .
OpenAI Codex CLI ( AGENTS.md )
Codex and terminal agents automatically read AGENTS.md in the repository root to invoke jsonbit via CLI or MCP when inspecting .json , .jsonl , or .jsonbit files.
Command
Arguments
Description
jsonbit card
<file.json|.jsonl|.jsonbit>
Prints the compact 3NF schema and statistics summary (independent of row count) plus storage compression stats.
jsonbit digest
<file.json|.jsonl|.jsonbit>
Prints the 1-Turn Factorized 3NF Holistic Digest (all rows, hoisted child tables, and precomputed summary stats).
jsonbit ask
<file> "<question>"
Evaluates a question against the bitmap index and prints the verified context slice.
jsonbit answer
<file> "<question>"
Evaluates a question and prints only the scalar answer value.
jsonbit sql
<file> "<SELECT ...>"
Executes a read-only SQL query against the 3NF columnar store ( WHERE , IN , BETWEEN , LIKE , GROUP BY , ORDER BY , LIMIT , OFFSET , CSR JOIN ).
jsonbit pack
<input.json> <output.jsonbit>
Serializes the bit-packed 3NF tables into a binary .jsonbit snapshot.
jsonbit rehydrate
<file.json|.jsonl|.jsonbit>
Reconstructs and prints the original nested JSON document.
jsonbit mcp
(none)
Starts the concurrent stdio JSON-RPC 2.0 MCP server ( olap_load , olap_digest , olap_ask , olap_sql , olap_rehydrate ).
Programmatic SDK Usage (Python, TypeScript, Rust)
Python ( pip install jsonbit )
from jsonbit import JsonBit
with JsonBit ( "benchmarks/data/github-repos.json" ) as olap :
card = olap . card ()
ans = olap . answer ( "How many repositories have more than 150000 stars and more than 30000 forks?" )
rows = olap . sql ( "SELECT defaultBranch, COUNT(*) AS cnt FROM records GROUP BY defaultBranch ORDER BY cnt DESC" )
olap . pack ( "repos.jsonbit" )
TypeScript / Node.js ( @jsonbit/core )
import { JsonBit , rehydrateFrom3NF } from '@jsonbit/core'
const engine = JsonBit . from ( jsonData )
const card = engine . getDataCard ( ) . text
const rows = engine . sql ( `SELECT orderId, customer.name, total FROM orders WHERE total > 500` )
const original = rehydrateFrom3NF ( engine )
engine . close ( )
Rust Library ( cargo add jsonbit )
use jsonbit :: { execute_sql , JsonBitEngine } ;
let engine = JsonBitEngine :: fr
[truncated]
flowchart TD
Input["Nested JSON / JSONL Input"] --> Normalizer["3NF Normalizer & CSR Indexer (rust/normalizer.rs)"]
Normalizer --> ColStore["Bit-Packed Columnar Store (rust/column.rs)"]
ColStore --> Bitmap["BitVec64 Bitmap Engine (rust/bitmap.rs)"]
Bitmap --> Card["OLAP Data Card (rust/engine.rs)"]
Bitmap --> SqlExec["SQL & Predicate Executor (rust/query_exec.rs)"]
ColStore --> Snapshot[".jsonbit Binary Snapshot (rust/snapshot.rs)"]
ColStore --> Rehydrate["Lossless 3NF Rehydrator (rust/normalizer.rs)"]
Loading
Storage Layout ( rust/column.rs & rust/normalizer.rs )
ForIntColumn (Frame-of-Reference Integers): Computes [min_val, max_val] per column and packs unsigned deltas val - min_val into ceil(log2(max_val - min_val + 1)) bits per row inside a Vec<u64> buffer. Column headers store sum , min , max , and non_null_count for O(1) unfiltered aggregations.
DictStringColumn (Dictionary + Inverted Bitmaps): Interns unique strings into a string arena, bit-packs dictionary IDs into ceil(log2(|D|)) bits per row, and builds a 64-bit word-aligned BitVec64 inverted index per category ( |D| <= 64 ) for O(1) equality lookups and POPCNT frequency histograms.
BoolBitmapColumn (1-Bit Booleans): Packs 64 boolean values into each u64 word alongside a 1-bit validity mask.
CSR 1-to-Many Join Index: Child arrays (e.g., orders[].items[] ) are normalized into child tables with a contiguous csr_offsets: Vec<u32> array mapping each parent row i to its child slice csr_offsets[i]..csr_offsets[i + 1] .
All benchmarks execute directly against the compiled Pure-Rust jsonbit binary ( target/release/jsonbit ).
1. Live gemini-3.8-flash Evaluation Across 10 Real-World Workloads (27 Questions)
Evaluates Raw JSON , Minified JSON , YAML , TOON ( @toon-format/toon ), and jsonbit across 10 production workloads ( W1 – W10 ), including flat tables, 1-to-Many e-commerce orders, sparse .jsonl OpenTelemetry logs with free-text messages, singleton Kubernetes configs, IAM primitive arrays, wide 30-column tables, and multi-byte UTF-8 financial ledgers. Across the 27 questions, single-prompt jsonbit digest uses 57.4% fewer estimated tokens than Minified JSON ( 101,302 vs. 237,891 ) with 26/27 vs. 13/27 accuracy:
Agentic local-query comparison (separate workflow — schema summary + Rust SQL execution, not a full-row text encoding):
Methodology note: Token counts above are gpt-tokenizer ( o200k_base ) estimates for reproducible cross-format comparison; the jsonbit SQL path total includes the schema summary prompt plus generated SQL query tokens.
For qualitative tasks where the LLM must read all records in a single turn and write a multi-fact executive summary (naming all departments, products, microservice failure messages, and international counterparties across W3 , W4 , W5 , and W10 ):
2. Token Scaling Curve ( 4 $\rightarrow$ 25,000 Nested Orders)
While linear text encodings ( Raw JSON , Minified JSON , YAML , TOON ) scale at $O(N)$ with row count and exceed LLM context windows at scale, jsonbit maintains an $O(1)$ prompt footprint ( ~67–473 tokens):
Note: For 5,000 - and 25,000 -row payloads (>2 MB), YAML and TOON token counts are ratio-extrapolated from the 500 -row exact encoding; jsonbit card tokens and Rust ingest times are measured directly on the full dataset.
3. Memory Footprint & Query Speed ( Pure-Rust jsonbit vs. jq CLI)
Workload breakdown: On tiny payloads ( W1 ), singleton configs ( W2 ), financial webhooks with heavy quoting ( W7 , +5.8% ), and wide flat 30-column tables ( W9 , -2.6% ), 1-turn digest is close to TOON . The largest reductions come from nested 1-to-Many hierarchies ( W4 , -55.2% ), sparse .jsonl logs ( W5 , -69.7% ), and primitive arrays ( W8 , -73.8% ), where 3NF hoisting and prefix/dictionary factoring eliminate repeated structure.
Storage Accounting ( W1 – W10 ):
In-memory bit-packed column store: 280.6 KB (Minified JSON) $\rightarrow$ 30.6 KB ( -89.1% )
Serialized on-disk .jsonbit snapshots ( jsonbit pack ): 280.6 KB (Minified JSON) / 427.7 KB (raw files) $\rightarrow$ 54.3 KB ( -80.6% vs. Minified JSON, -87.3% vs. raw files)
Query Latency (Warm In-Memory Rust Engine vs. Cold jq CLI Subprocess):
jq CLI ( stedolan/jq , cold subprocess including

[truncated]

## Original Extract

Contribute to prathikmurugan/jsonbit development by creating an account on GitHub.

GitHub - prathikmurugan/jsonbit · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects, and automated work
Build & Ship Actions CI/CD and automation
Code Quality Quality analysis, fixes, and enforcement
Codespaces Cloud development environments
Secure Code Security Vulnerability detection and remediation
Secret Protection Secret scanning and leak prevention
Developers Learn Documentation
Changelog Changelog Your source for the latest features, improvements, and fixes across GitHub products, all in one place.
Businesses By Stage Enterprises
Add-ons GitHub Advanced Security Enterprise-grade code security
Copilot for business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Resources Learn & Connect Documentation
GitHub Universe GitHub Universe The event for developers, by developers — talks, workshops, and what’s next for GitHub.
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
prathikmurugan
jsonbit Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
13 Commits 13 Commits Folders and files
.github .github benchmarks benchmarks bin bin python/ jsonbit python/ jsonbit rust rust scripts scripts src src test test .gitignore .gitignore .mcp.json .mcp.json AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md opencode.json opencode.json package-lock.json package-lock.json package.json package.json pyproject.toml pyproject.toml tsconfig.json tsconfig.json View all files Repository files navigation
JsonBit compacts nested JSON before it reaches your LLM. Send a compact digest in one prompt, or let your agent query the original data locally via SQL/MCP and read only the results it needs.
pip install jsonbit
jsonbit digest path/to/your-file.json
Pass the output to your LLM in place of the raw JSON. (Rust cargo , Node.js npm / npx , and standalone binaries are listed under Installation & Setup .)
Example output from jsonbit digest :
# orders[4](orderId,customer_name,customer_tier,status,total)
@stats: total[sum=619.00,avg=154.75,min=95.50,max=201.00]
ORD-101,Alice,pro,shipped,171.75
ORD-102,Bob,starter,pending,201
ORD-103,Cara,pro,shipped,95.5
ORD-104,Dan,pro,shipped,150.75
# orders_items[6] -> orders (sku,qty,price)
@stats: qty[sum=15,avg=2.5,min=1(n=1),max=4] price[sum=246.25,avg=41.04,min=10.50,max=74.50]
0:USB-C,2,10.5|MOUSE,3,50.25
1:MOUSE,4,50.25
2:HUB-4,1,74.5|USB-C,2,10.5
3:MOUSE,3,50.25
Workflow
Command / MCP Tool
Best For
Fidelity & Guarantees
1-Turn Compaction
jsonbit digest / olap_digest
Single-prompt LLM calls & holistic summaries where the model needs every row.
Compact LLM text view containing all rows, escaped delimiters, hoisted child tables, and precomputed stats (rounded averages are approximate) .
Agentic Local Querying
jsonbit card + sql / olap_load + olap_sql
Coding agents (OpenCode, Claude Code, Cursor) querying medium/large .json or .jsonl files.
Emits a compact schema and statistics summary (independent of row count); runs exact read-only SQL & BitVec64 filters locally in microseconds.
Binary Storage
jsonbit pack / jsonbit rehydrate
Archiving or caching normalized .jsonbit snapshots on disk ( -80.6% vs. minified JSON on W1 – W10 ).
Lossless round-trip verified across all 33 benchmark/adversarial datasets and unit tests (preserving nulls, colliding keys, and u64 > i64::MAX ).
Table of Contents
Agent Integrations (OpenCode, Claude Code, Codex, Cursor)
Install the pre-built native jsonbit binary using your preferred package manager (no Rust compiler required for Python or Node.js):
Python ( pip / uv ) — Includes native CLI + Python SDK
pip install jsonbit
Rust ( cargo )
cargo install jsonbit
Node.js ( npm / npx ) — Includes native CLI + TypeScript SDK
npm install -g @jsonbit/core
# or run directly without global install:
npx @jsonbit/core --help
Pre-Built Binaries & Building from Source
Standalone binaries for Linux ( x64 , arm64 ), macOS ( arm64 , x64 ), and Windows ( x64 ) are attached to every GitHub Release .
To build from source (requires Rust 1.75+ ):
git clone https://github.com/prathikmurugan/jsonbit.git
cd jsonbit
cargo install --path .
Verify installation:
jsonbit --version
Agent Integrations
jsonbit includes a built-in stdio JSON-RPC 2.0 MCP server ( jsonbit mcp ) with zero external runtime dependencies. It exposes five tools: olap_load , olap_digest , olap_ask , olap_sql , and olap_rehydrate .
Add the following to opencode.json in your project root (or globally in ~/.config/opencode/opencode.json ):
{
"$schema" : " https://opencode.ai/config.json " ,
"mcp" : {
"jsonbit" : {
"type" : " local " ,
"command" : [ " jsonbit " , " mcp " ],
"enabled" : true
}
}
}
Claude Code ( .mcp.json or CLI)
Register via the Claude Code CLI:
claude mcp add jsonbit -- jsonbit mcp
Or include in .mcp.json :
{
"mcpServers" : {
"jsonbit" : {
"command" : " jsonbit " ,
"args" : [ " mcp " ]
}
}
}
Cursor & Windsurf
Add the .mcp.json configuration above to .cursor/mcp.json or ~/.codeium/windsurf/mcp_config.json .
OpenAI Codex CLI ( AGENTS.md )
Codex and terminal agents automatically read AGENTS.md in the repository root to invoke jsonbit via CLI or MCP when inspecting .json , .jsonl , or .jsonbit files.
Command
Arguments
Description
jsonbit card
<file.json|.jsonl|.jsonbit>
Prints the compact 3NF schema and statistics summary (independent of row count) plus storage compression stats.
jsonbit digest
<file.json|.jsonl|.jsonbit>
Prints the 1-Turn Factorized 3NF Holistic Digest (all rows, hoisted child tables, and precomputed summary stats).
jsonbit ask
<file> "<question>"
Evaluates a question against the bitmap index and prints the verified context slice.
jsonbit answer
<file> "<question>"
Evaluates a question and prints only the scalar answer value.
jsonbit sql
<file> "<SELECT ...>"
Executes a read-only SQL query against the 3NF columnar store ( WHERE , IN , BETWEEN , LIKE , GROUP BY , ORDER BY , LIMIT , OFFSET , CSR JOIN ).
jsonbit pack
<input.json> <output.jsonbit>
Serializes the bit-packed 3NF tables into a binary .jsonbit snapshot.
jsonbit rehydrate
<file.json|.jsonl|.jsonbit>
Reconstructs and prints the original nested JSON document.
jsonbit mcp
(none)
Starts the concurrent stdio JSON-RPC 2.0 MCP server ( olap_load , olap_digest , olap_ask , olap_sql , olap_rehydrate ).
Programmatic SDK Usage (Python, TypeScript, Rust)
Python ( pip install jsonbit )
from jsonbit import JsonBit
with JsonBit ( "benchmarks/data/github-repos.json" ) as olap :
card = olap . card ()
ans = olap . answer ( "How many repositories have more than 150000 stars and more than 30000 forks?" )
rows = olap . sql ( "SELECT defaultBranch, COUNT(*) AS cnt FROM records GROUP BY defaultBranch ORDER BY cnt DESC" )
olap . pack ( "repos.jsonbit" )
TypeScript / Node.js ( @jsonbit/core )
import { JsonBit , rehydrateFrom3NF } from '@jsonbit/core'
const engine = JsonBit . from ( jsonData )
const card = engine . getDataCard ( ) . text
const rows = engine . sql ( `SELECT orderId, customer.name, total FROM orders WHERE total > 500` )
const original = rehydrateFrom3NF ( engine )
engine . close ( )
Rust Library ( cargo add jsonbit )
use jsonbit :: { execute_sql , JsonBitEngine } ;
let engine = JsonBitEngine :: fr
[truncated]
flowchart TD
Input["Nested JSON / JSONL Input"] --> Normalizer["3NF Normalizer & CSR Indexer (rust/normalizer.rs)"]
Normalizer --> ColStore["Bit-Packed Columnar Store (rust/column.rs)"]
ColStore --> Bitmap["BitVec64 Bitmap Engine (rust/bitmap.rs)"]
Bitmap --> Card["OLAP Data Card (rust/engine.rs)"]
Bitmap --> SqlExec["SQL & Predicate Executor (rust/query_exec.rs)"]
ColStore --> Snapshot[".jsonbit Binary Snapshot (rust/snapshot.rs)"]
ColStore --> Rehydrate["Lossless 3NF Rehydrator (rust/normalizer.rs)"]
Loading
Storage Layout ( rust/column.rs & rust/normalizer.rs )
ForIntColumn (Frame-of-Reference Integers): Computes [min_val, max_val] per column and packs unsigned deltas val - min_val into ceil(log2(max_val - min_val + 1)) bits per row inside a Vec<u64> buffer. Column headers store sum , min , max , and non_null_count for O(1) unfiltered aggregations.
DictStringColumn (Dictionary + Inverted Bitmaps): Interns unique strings into a string arena, bit-packs dictionary IDs into ceil(log2(|D|)) bits per row, and builds a 64-bit word-aligned BitVec64 inverted index per category ( |D| <= 64 ) for O(1) equality lookups and POPCNT frequency histograms.
BoolBitmapColumn (1-Bit Booleans): Packs 64 boolean values into each u64 word alongside a 1-bit validity mask.
CSR 1-to-Many Join Index: Child arrays (e.g., orders[].items[] ) are normalized into child tables with a contiguous csr_offsets: Vec<u32> array mapping each parent row i to its child slice csr_offsets[i]..csr_offsets[i + 1] .
All benchmarks execute directly against the compiled Pure-Rust jsonbit binary ( target/release/jsonbit ).
1. Live gemini-3.8-flash Evaluation Across 10 Real-World Workloads (27 Questions)
Evaluates Raw JSON , Minified JSON , YAML , TOON ( @toon-format/toon ), and jsonbit across 10 production workloads ( W1 – W10 ), including flat tables, 1-to-Many e-commerce orders, sparse .jsonl OpenTelemetry logs with free-text messages, singleton Kubernetes configs, IAM primitive arrays, wide 30-column tables, and multi-byte UTF-8 financial ledgers. Across the 27 questions, single-prompt jsonbit digest uses 57.4% fewer estimated tokens than Minified JSON ( 101,302 vs. 237,891 ) with 26/27 vs. 13/27 accuracy:
Agentic local-query comparison (separate workflow — schema summary + Rust SQL execution, not a full-row text encoding):
Methodology note: Token counts above are gpt-tokenizer ( o200k_base ) estimates for reproducible cross-format comparison; the jsonbit SQL path total includes the schema summary prompt plus generated SQL query tokens.
For qualitative tasks where the LLM must read all records in a single turn and write a multi-fact executive summary (naming all departments, products, microservice failure messages, and international counterparties across W3 , W4 , W5 , and W10 ):
2. Token Scaling Curve ( 4 $\rightarrow$ 25,000 Nested Orders)
While linear text encodings ( Raw JSON , Minified JSON , YAML , TOON ) scale at $O(N)$ with row count and exceed LLM context windows at scale, jsonbit maintains an $O(1)$ prompt footprint ( ~67–473 tokens):
Note: For 5,000 - and 25,000 -row payloads (>2 MB), YAML and TOON token counts are ratio-extrapolated from the 500 -row exact encoding; jsonbit card tokens and Rust ingest times are measured directly on the full dataset.
3. Memory Footprint & Query Speed ( Pure-Rust jsonbit vs. jq CLI)
Workload breakdown: On tiny payloads ( W1 ), singleton configs ( W2 ), financial webhooks with heavy quoting ( W7 , +5.8% ), and wide flat 30-column tables ( W9 , -2.6% ), 1-turn digest is close to TOON . The largest reductions come from nested 1-to-Many hierarchies ( W4 , -55.2% ), sparse .jsonl logs ( W5 , -69.7% ), and primitive arrays ( W8 , -73.8% ), where 3NF hoisting and prefix/dictionary factoring eliminate repeated structure.
Storage Accounting ( W1 – W10 ):
In-memory bit-packed column store: 280.6 KB (Minified JSON) $\rightarrow$ 30.6 KB ( -89.1% )
Serialized on-disk .jsonbit snapshots ( jsonbit pack ): 280.6 KB (Minified JSON) / 427.7 KB (raw files) $\rightarrow$ 54.3 KB ( -80.6% vs. Minified JSON, -87.3% vs. raw files)
Query Latency (Warm In-Memory Rust Engine vs. Cold jq CLI Subprocess):
jq CLI ( stedolan/jq , cold subprocess including

[truncated]
