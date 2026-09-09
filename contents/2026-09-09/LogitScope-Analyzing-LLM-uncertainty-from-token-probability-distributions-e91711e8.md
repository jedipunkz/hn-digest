---
source: "https://github.com/ibm-granite/granite.debug-tools/tree/main/logitscope"
hn_url: "https://news.ycombinator.com/item?id=49628806"
title: "LogitScope: Analyzing LLM uncertainty from token probability distributions"
article_title: "granite.debug-tools/logitscope at main · ibm-granite/granite.debug-tools · GitHub"
image: "https://opengraph.githubassets.com/42ff2bf1d72f27baaec26ac82296cf5e40e4cb8787ae987d99bc20f6b2330943/ibm-granite/granite.debug-tools"
author: "mncharity"
captured_at: "2026-09-09T17:04:09Z"
capture_tool: "hn-digest"
hn_id: 49628806
score: 2
comments: 1
posted_at: "2026-09-09T16:08:18Z"
tags:
  - hacker-news
---

# LogitScope: Analyzing LLM uncertainty from token probability distributions

- HN: [49628806](https://news.ycombinator.com/item?id=49628806)
- Source: [github.com](https://github.com/ibm-granite/granite.debug-tools/tree/main/logitscope)
- Score: 2
- Comments: 1
- Posted: 2026-09-09T16:08:18Z

## Translation

Title: LogitScope: Analyzing LLM uncertainty from token probability distributions
Article title: granite.debug-tools/logitscope at main · ibm-granite/granite.debug-tools · GitHub
Description: Granite Debug Tools. Contribute to ibm-granite/granite.debug-tools development by creating an account on GitHub.

Article text:
granite.debug-tools/logitscope at main · ibm-granite/granite.debug-tools · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
ibm-granite
/
granite.debug-tools
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
More options Directory actions
History History main Breadcrumbs
Copy path Top Folders and files
.. examples examples logitscope logitscope .gitattributes .gitattributes .gitignore .gitignore README.md README.md pyproject.toml pyproject.toml View all files README.md
LogitScope is a Python framework for analyzing large language models through information metrics computed from token probability distributions. By examining how models distribute probability mass across vocabulary at each position, LogitScope provides objective, quantitative insights into model behavior, uncertainty, and decision-making without relying on semantic interpretation.
LogitScope transforms opaque LLM outputs into interpretable measurements by:
Quantifying uncertainty : Measures how confident or uncertain the model is at each token position
Detecting distribution patterns : Identifies when models are choosing between multiple options vs. making confident predictions
Revealing surprising outputs : Flags tokens that are statistically unexpected given the context
Tracking prediction quality : Monitors cumulative model performance through perplexity and other metrics
Exposing alternative paths : Shows top-k alternative tokens the model considered at each position
This enables debugging hallucinations, optimizing prompts, evaluating fine-tuning results, and understanding model behavior quantitatively.
📊 Information-theoretic metrics: Entropy, varentropy, skewentropy, surprisal, perplexity, and probability
🔌 Universal compatibility: Works with any HuggingFace causal language model
🚀 Efficient computation: Lazy evaluation with intelligent caching
💻 Flexible APIs: Python library for programmatic access or interactive web UI
⚡ Multi-platform: CPU, CUDA, and MPS (Apple Silicon) support
🔧 Extensible: Plugin architecture for custom metrics
# clone repo
git clone https://github.com/ibm-granite/granite.debug-tools
cd granite.debug-tools/logitscope
# create and activate uv venv
uv venv --python 3.11
source .venv/bin/activate
Basic Installation (Python API only)
For programmatic use of LogitScope without the web UI:
uv pip install -e .
Full Installation (with Web UI)
To use the interactive web interface:
uv pip install -e " .[ui] "
Development Installation
For contributing to LogitScope:
uv pip install -e " .[ui,dev] "
Requirements
LogitScope uses Python 3.11+ in addition to the following packages.
Core dependencies (always installed):
UI dependencies (optional, installed with [ui] ):
Development dependencies (optional, installed with [dev] ):
Use LogitScope programmatically in your code for analysis pipelines, experiments, or integration into existing workflows:
from transformers import AutoModelForCausalLM , AutoTokenizer
from logitscope import LogitScope
# Load model and tokenizer
model = AutoModelForCausalLM . from_pretrained ( 'HuggingFaceTB/SmolLM2-135M-Instruct' )
tokenizer = AutoTokenizer . from_pretrained ( 'HuggingFaceTB/SmolLM2-135M-Instruct' )
# Initialize LogitScope
scope = LogitScope ( tokenizer , model , device = 'cpu' )
# Analyze text
results = scope . measure ( 'The quick brown fox jumps over the lazy dog' )
# Access metrics as arrays
print ( f"Surprisal: { results . surprisal } " )
print ( f"Entropy: { results . entropy } " )
print ( f"Varentropy: { results . varentropy } " )
# Iterate over tokens with metrics
# Use repr() to show escape sequences (\n, \t, etc.) in CLI output
for token in results . iter_tokens ([ 'surprisal' , 'entropy' , 'perplexity' ]):
print (
f" { repr ( token [ 'token' ]):15s } | "
f"surprisal= { token [ 'surprisal' ]:.3f } | "
f"entropy= { token [ 'entropy' ]:.3f } | "
f"perplexity= { token [ 'perplexity' ]:.3f } "
)
# Inspect top-k alternatives at specific positions
# repr() makes special characters visible
top_tokens = results . top_k ( index = 5 , k = 10 )
for token , prob in top_tokens :
print ( f" { repr ( token ):15s } { prob :.4f } " )
View the examples/ directory for more examples on how to use the framework.
Note: The web UI requires the optional UI dependencies. Install with uv pip install -e ".[ui]" if not already installed.
Launch the web interface for visual exploration and real-time analysis:
python -m logitscope.ui --model HuggingFaceTB/SmolLM2-135M-Instruct --device cpu
Then open http://0.0.0.0:8000 in your browser.
Color-coded token visualization by metric
Live statistics with running averages
Interactive token explorer with top-k alternatives
Distribution plots and metric toggles
# Custom port
python -m logitscope.ui --model MODEL_NAME --port 8080
# GPU acceleration
python -m logitscope.ui --model MODEL_NAME --device cuda
# Apple Silicon
python -m logitscope.ui --model MODEL_NAME --device mps
How It Works
At each token position, language models produce a probability distribution over their entire vocabulary. LogitScope analyzes these distributions to extract meaningful signals about model behavior:
Tokenization : Input text is converted into tokens using the model's tokenizer
Model inference : The model generates logits (pre-softmax scores) for next-token predictions
Probability computation : Logits are transformed into probability distributions via softmax
Metric calculation : Information-theoretic properties are computed from the distributions
Analysis : Metrics reveal uncertainty patterns, distribution shapes, and prediction quality
Each metric captures a different aspect of the probability distribution:
Hallucination detection : Regions with high entropy and varentropy often indicate fabricated content. When a model generates facts with high confidence but also high uncertainty in surrounding tokens, it may be hallucinating.
Production monitoring : Track perplexity and average entropy over time to detect model degradation or unexpected behavior in deployment.
Output comparison : Quantitatively compare multiple model responses to the same prompt—lower perplexity and higher token probabilities indicate more "natural" outputs.
Fine-tuning evaluation : Measure before/after metrics on validation sets. Improved models show lower perplexity and higher probability on expected outputs without increased entropy on known-answer questions.
Prompt engineering : Observe how different prompt phrasings affect model confidence. Good prompts reduce entropy while maintaining high probability on correct tokens.
Ablation studies : When changing model architecture or training procedures, LogitScope provides objective metrics beyond just accuracy to quantify the impact.
Decision pattern analysis : Study how models make choices by examining entropy/varentropy patterns across different text types, domains, or generation strategies.
Uncertainty quantification : Measure whether model confidence scores correlate with actual correctness—essential for calibration studies.
Distribution characteristics : Investigate how probability distributions differ across languages, domains, or model sizes to understand architectural and training effects.
LogitScope(tokenizer, model, device='cpu', seed=None)
Main analysis interface that wraps any HuggingFace causal language model.
tokenizer (PreTrainedTokenizer) - HuggingFace tokenizer matching the model
model (PreTrainedModel) - HuggingFace causal language model
device (str) - Computation device: 'cpu' , 'cuda' , or 'mps'
seed (int, optional) - Random seed for reproducibility
measure(text: str) -> Results - Analyzes input text and returns a Results object with computed metrics
Container for analysis results with lazy-evaluated metrics. Metrics are computed on-demand when accessed and cached for efficiency.
Metric Properties (computed lazily):
surprisal - List of surprisal values for actual tokens chosen
entropy - List of entropy values, one per token position
varentropy - List of varentropy values (distribution variance)
skewentropy - List of skewentropy values (distribution asymmetry)
perplexity - List of perplexity values (cumulative quality)
iter_tokens(metrics: list[str]) -> Iterator[dict] - Iterates over tokens with their text, IDs, and requested metric values
top_k(index: int, k: int) -> list[tuple[str, float]] - Returns the top-k most probable tokens at a given position with their probabilities
input_ids - Token IDs as tensors
logits - Raw model logits (pre-softmax scores)
probs - Probability distributions over vocabulary
log_probs - Log probabilities (for numerical stability)
tokenizer - Tokenizer instance used for decoding
LogitScope automatically decodes tokens using tokenizer.decode() to produce human-readable text. This means:
Tokens from iter_tokens() are properly decoded with actual whitespace
Tokens from top_k() are decoded the same way
Works correctly with any tokenizer (GPT-2, LLaMA, etc.)
No manual string processing needed in your code or UI
# Tokens are already decoded!
for token in results . iter_tokens ():
print ( token [ 'token' ]) # Properly decoded by tokenizer
Development
LogitScope uses modern Python tooling for development:
# Install with UI and dev dependencies
uv pip install -e " .[ui,dev] "
# Format code
ruff format .
# Lint code
ruff check .
# Auto-fix issues
ruff check --fix .
Citation
@misc { ahmed2026logitscopeframeworkanalyzingllm ,
title = { LogitScope: A Framework for Analyzing LLM Uncertainty Through Information Metrics } ,
author = { Farhan Ahmed and Yuya Jeremy Ong and Chad DeLuca } ,
year = { 2026 } ,
eprint = { 2603.24929 } ,
archivePrefix = { arXiv } ,
primaryClass = { cs.AI } ,
url = { https://arxiv.org/abs/2603.24929 } ,
}
Footer
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Granite Debug Tools. Contribute to ibm-granite/granite.debug-tools development by creating an account on GitHub.

granite.debug-tools/logitscope at main · ibm-granite/granite.debug-tools · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
ibm-granite
/
granite.debug-tools
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
More options Directory actions
History History main Breadcrumbs
Copy path Top Folders and files
.. examples examples logitscope logitscope .gitattributes .gitattributes .gitignore .gitignore README.md README.md pyproject.toml pyproject.toml View all files README.md
LogitScope is a Python framework for analyzing large language models through information metrics computed from token probability distributions. By examining how models distribute probability mass across vocabulary at each position, LogitScope provides objective, quantitative insights into model behavior, uncertainty, and decision-making without relying on semantic interpretation.
LogitScope transforms opaque LLM outputs into interpretable measurements by:
Quantifying uncertainty : Measures how confident or uncertain the model is at each token position
Detecting distribution patterns : Identifies when models are choosing between multiple options vs. making confident predictions
Revealing surprising outputs : Flags tokens that are statistically unexpected given the context
Tracking prediction quality : Monitors cumulative model performance through perplexity and other metrics
Exposing alternative paths : Shows top-k alternative tokens the model considered at each position
This enables debugging hallucinations, optimizing prompts, evaluating fine-tuning results, and understanding model behavior quantitatively.
📊 Information-theoretic metrics: Entropy, varentropy, skewentropy, surprisal, perplexity, and probability
🔌 Universal compatibility: Works with any HuggingFace causal language model
🚀 Efficient computation: Lazy evaluation with intelligent caching
💻 Flexible APIs: Python library for programmatic access or interactive web UI
⚡ Multi-platform: CPU, CUDA, and MPS (Apple Silicon) support
🔧 Extensible: Plugin architecture for custom metrics
# clone repo
git clone https://github.com/ibm-granite/granite.debug-tools
cd granite.debug-tools/logitscope
# create and activate uv venv
uv venv --python 3.11
source .venv/bin/activate
Basic Installation (Python API only)
For programmatic use of LogitScope without the web UI:
uv pip install -e .
Full Installation (with Web UI)
To use the interactive web interface:
uv pip install -e " .[ui] "
Development Installation
For contributing to LogitScope:
uv pip install -e " .[ui,dev] "
Requirements
LogitScope uses Python 3.11+ in addition to the following packages.
Core dependencies (always installed):
UI dependencies (optional, installed with [ui] ):
Development dependencies (optional, installed with [dev] ):
Use LogitScope programmatically in your code for analysis pipelines, experiments, or integration into existing workflows:
from transformers import AutoModelForCausalLM , AutoTokenizer
from logitscope import LogitScope
# Load model and tokenizer
model = AutoModelForCausalLM . from_pretrained ( 'HuggingFaceTB/SmolLM2-135M-Instruct' )
tokenizer = AutoTokenizer . from_pretrained ( 'HuggingFaceTB/SmolLM2-135M-Instruct' )
# Initialize LogitScope
scope = LogitScope ( tokenizer , model , device = 'cpu' )
# Analyze text
results = scope . measure ( 'The quick brown fox jumps over the lazy dog' )
# Access metrics as arrays
print ( f"Surprisal: { results . surprisal } " )
print ( f"Entropy: { results . entropy } " )
print ( f"Varentropy: { results . varentropy } " )
# Iterate over tokens with metrics
# Use repr() to show escape sequences (\n, \t, etc.) in CLI output
for token in results . iter_tokens ([ 'surprisal' , 'entropy' , 'perplexity' ]):
print (
f" { repr ( token [ 'token' ]):15s } | "
f"surprisal= { token [ 'surprisal' ]:.3f } | "
f"entropy= { token [ 'entropy' ]:.3f } | "
f"perplexity= { token [ 'perplexity' ]:.3f } "
)
# Inspect top-k alternatives at specific positions
# repr() makes special characters visible
top_tokens = results . top_k ( index = 5 , k = 10 )
for token , prob in top_tokens :
print ( f" { repr ( token ):15s } { prob :.4f } " )
View the examples/ directory for more examples on how to use the framework.
Note: The web UI requires the optional UI dependencies. Install with uv pip install -e ".[ui]" if not already installed.
Launch the web interface for visual exploration and real-time analysis:
python -m logitscope.ui --model HuggingFaceTB/SmolLM2-135M-Instruct --device cpu
Then open http://0.0.0.0:8000 in your browser.
Color-coded token visualization by metric
Live statistics with running averages
Interactive token explorer with top-k alternatives
Distribution plots and metric toggles
# Custom port
python -m logitscope.ui --model MODEL_NAME --port 8080
# GPU acceleration
python -m logitscope.ui --model MODEL_NAME --device cuda
# Apple Silicon
python -m logitscope.ui --model MODEL_NAME --device mps
How It Works
At each token position, language models produce a probability distribution over their entire vocabulary. LogitScope analyzes these distributions to extract meaningful signals about model behavior:
Tokenization : Input text is converted into tokens using the model's tokenizer
Model inference : The model generates logits (pre-softmax scores) for next-token predictions
Probability computation : Logits are transformed into probability distributions via softmax
Metric calculation : Information-theoretic properties are computed from the distributions
Analysis : Metrics reveal uncertainty patterns, distribution shapes, and prediction quality
Each metric captures a different aspect of the probability distribution:
Hallucination detection : Regions with high entropy and varentropy often indicate fabricated content. When a model generates facts with high confidence but also high uncertainty in surrounding tokens, it may be hallucinating.
Production monitoring : Track perplexity and average entropy over time to detect model degradation or unexpected behavior in deployment.
Output comparison : Quantitatively compare multiple model responses to the same prompt—lower perplexity and higher token probabilities indicate more "natural" outputs.
Fine-tuning evaluation : Measure before/after metrics on validation sets. Improved models show lower perplexity and higher probability on expected outputs without increased entropy on known-answer questions.
Prompt engineering : Observe how different prompt phrasings affect model confidence. Good prompts reduce entropy while maintaining high probability on correct tokens.
Ablation studies : When changing model architecture or training procedures, LogitScope provides objective metrics beyond just accuracy to quantify the impact.
Decision pattern analysis : Study how models make choices by examining entropy/varentropy patterns across different text types, domains, or generation strategies.
Uncertainty quantification : Measure whether model confidence scores correlate with actual correctness—essential for calibration studies.
Distribution characteristics : Investigate how probability distributions differ across languages, domains, or model sizes to understand architectural and training effects.
LogitScope(tokenizer, model, device='cpu', seed=None)
Main analysis interface that wraps any HuggingFace causal language model.
tokenizer (PreTrainedTokenizer) - HuggingFace tokenizer matching the model
model (PreTrainedModel) - HuggingFace causal language model
device (str) - Computation device: 'cpu' , 'cuda' , or 'mps'
seed (int, optional) - Random seed for reproducibility
measure(text: str) -> Results - Analyzes input text and returns a Results object with computed metrics
Container for analysis results with lazy-evaluated metrics. Metrics are computed on-demand when accessed and cached for efficiency.
Metric Properties (computed lazily):
surprisal - List of surprisal values for actual tokens chosen
entropy - List of entropy values, one per token position
varentropy - List of varentropy values (distribution variance)
skewentropy - List of skewentropy values (distribution asymmetry)
perplexity - List of perplexity values (cumulative quality)
iter_tokens(metrics: list[str]) -> Iterator[dict] - Iterates over tokens with their text, IDs, and requested metric values
top_k(index: int, k: int) -> list[tuple[str, float]] - Returns the top-k most probable tokens at a given position with their probabilities
input_ids - Token IDs as tensors
logits - Raw model logits (pre-softmax scores)
probs - Probability distributions over vocabulary
log_probs - Log probabilities (for numerical stability)
tokenizer - Tokenizer instance used for decoding
LogitScope automatically decodes tokens using tokenizer.decode() to produce human-readable text. This means:
Tokens from iter_tokens() are properly decoded with actual whitespace
Tokens from top_k() are decoded the same way
Works correctly with any tokenizer (GPT-2, LLaMA, etc.)
No manual string processing needed in your code or UI
# Tokens are already decoded!
for token in results . iter_tokens ():
print ( token [ 'token' ]) # Properly decoded by tokenizer
Development
LogitScope uses modern Python tooling for development:
# Install with UI and dev dependencies
uv pip install -e " .[ui,dev] "
# Format code
ruff format .
# Lint code
ruff check .
# Auto-fix issues
ruff check --fix .
Citation
@misc { ahmed2026logitscopeframeworkanalyzingllm ,
title = { LogitScope: A Framework for Analyzing LLM Uncertainty Through Information Metrics } ,
author = { Farhan Ahmed and Yuya Jeremy Ong and Chad DeLuca } ,
year = { 2026 } ,
eprint = { 2603.24929 } ,
archivePrefix = { arXiv } ,
primaryClass = { cs.AI } ,
url = { https://arxiv.org/abs/2603.24929 } ,
}
Footer
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
