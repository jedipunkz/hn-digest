---
source: "https://github.com/ariveinberg-max/aven-1"
hn_url: "https://news.ycombinator.com/item?id=49612189"
title: "I built a 20M-parameter LLM from scratch, including full RLHF (PPO)"
article_title: "GitHub - ariveinberg-max/aven-1: A from-scratch causal Transformer trained from random weights, no pretrained model or API. · GitHub"
image: "https://opengraph.githubassets.com/77c88c80a7f3efa60509405ff83c6ea6749d8ccc2913b36bfbf2f47895203d70/ariveinberg-max/aven-1"
author: "neurovance"
captured_at: "2026-09-08T16:06:49Z"
capture_tool: "hn-digest"
hn_id: 49612189
score: 1
comments: 0
posted_at: "2026-09-08T15:56:55Z"
tags:
  - hacker-news
---

# I built a 20M-parameter LLM from scratch, including full RLHF (PPO)

- HN: [49612189](https://news.ycombinator.com/item?id=49612189)
- Source: [github.com](https://github.com/ariveinberg-max/aven-1)
- Score: 1
- Comments: 0
- Posted: 2026-09-08T15:56:55Z

## Translation

Title: I built a 20M-parameter LLM from scratch, including full RLHF (PPO)
Article title: GitHub - ariveinberg-max/aven-1: A from-scratch causal Transformer trained from random weights, no pretrained model or API. · GitHub
Description: A from-scratch causal Transformer trained from random weights, no pretrained model or API. - ariveinberg-max/aven-1

Article text:
GitHub - ariveinberg-max/aven-1: A from-scratch causal Transformer trained from random weights, no pretrained model or API. · GitHub
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
ariveinberg-max
/
aven-1
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
10 Commits 10 Commits Folders and files
data data runs runs .gitignore .gitignore Aven-1-Colab.ipynb Aven-1-Colab.ipynb README.md README.md Start Brain.command Start Brain.command WRITEUP.md WRITEUP.md brain.py brain.py chat.html chat.html make_instructions.py make_instructions.py memory.py memory.py ppo.py ppo.py preferences.py preferences.py raft.py raft.py requirements.txt requirements.txt reward_model.py reward_model.py server.py server.py sources.py sources.py sweep.yaml sweep.yaml test_brain.py test_brain.py tokenizer.py tokenizer.py train.py train.py train_reward.py train_reward.py ui.html ui.html value_model.py value_model.py View all files Repository files navigation
A working first step toward your own AI: a 5,055,360-parameter byte-level Transformer, trained from random weights. No pretrained model, API key, or cloud inference. PyTorch provides tensor math, automatic differentiation, and AdamW; the architecture, training loop, checkpoint handling, and interface are in this project. This is not a general assistant, conscious brain, or a model with voice, vision, web access, or persistent conversational memory.
Double-click Start Brain.command in this folder. Keep the Terminal window open while using the app. Visit http://127.0.0.1:8765 if the browser does not open. Close the server with Control-C in its Terminal window; closing a browser tab does not stop training.
The first verified checkpoint is already available. Click Generate text to try it. Click Resume training to add the chosen number of training steps. Pause & save finishes the current step before saving. The app runs only on this computer.
If macOS does not launch the command file, run these commands in Terminal:
cd /Users/ariveinbers/Documents/Codex/my-ai-brain
.venv/bin/python server.py
What it learns
It predicts the next UTF-8 byte from up to 128 preceding bytes. Four causal Transformer blocks use five attention heads and a hidden width of 320. Batch size 8 and float32 are conservative starting defaults for an 8 GB Mac. GPU availability is checked at runtime; training selects MPS when available. Generation uses CPU to keep the interface simple and memory use modest. Close other heavy apps if your Mac experiences memory pressure.
The included data/demo.txt contains original, programmatically composed practice stories. Many share a template. They are a pipeline test, not a broad knowledge source. The last 10% of bytes are held out of gradient updates, but the similar templates make this an easy validation set. Falling loss on it does not establish general intelligence, useful conversation, or generalization to unrelated text. Output after a short run is expected to be rough.
The diagram shows actual final-position hidden-state magnitudes after generation, grouped into 32 nodes per block (10 channels per node). Brightness is normalized separately within each block. The connecting lines are a schematic, not a display of individual learned weights. It is not a live scan of thoughts.
Use UTF-8 plain text that you own, wrote, or have permission to train on. The starter accepts 4 KB–20 MB. More diverse, clean text is more useful than simply repeating a small document. It reads the entire corpus into memory, so the file-size limit is intentional.
To start a new run while preserving the current one, first pause training, wait for it to finish, and stop the server. Rename checkpoints to an unused name such as checkpoints-demo-backup . Then save your text as data/training.txt and reopen the app. The app will create new random weights for the next training session. The UI also accepts pasted text before a checkpoint exists. Keep backups private if your training text contains personal material.
To continue a CLI run, use the same corpus path every time:
.venv/bin/python train.py --data data/training.txt --steps 200 --resume
For the included demo, omit --data . A checkpoint records the corpus hash and rejects a different corpus on resume. It preserves model weights, optimizer state, step count, and CPU random-generator state. Checkpoints save every 20 steps, at normal completion, and on a handled pause. Force-quitting or losing power can lose steps since the previous save. Run one training process at a time; use the panel or CLI, not both simultaneously.
Training-data workspace vs. stored memory vs. learned weights
The panel now separates three things that are easy to conflate:
Learned weights ( checkpoints/latest.pt ): numbers adjusted by gradient descent. This is the only thing that makes the model's text generation change. You cannot edit these directly; you can only change them by training.
Training-data workspace ( data/sources/*.txt , managed from the "Training-data workspace" panel): the plain text a future run would learn from. Add, edit, or delete named sources here, then click Compile into training corpus to concatenate them into data/training.txt . Once a checkpoint exists, the corpus is locked (same rule as before) — compiling, saving, or deleting a source is refused until you preserve checkpoints/ and start a new run.
Memory ( memory.db , managed from the "Memory" panel): a plain SQLite table of notes you write, edit, and delete directly through the UI. It is not read by training or generation, and saving a note never changes a single model weight. It exists purely as external, inspectable storage — the opposite of the model's opaque learned parameters.
Deleting a memory record or a workspace source is permanent immediately; there is no undo.
Weights & Biases (cloud, free tier, project aven-1 ): pass --wandb to train.py for full run tracking at https://wandb.ai/ariveinberg-ari-research/aven-1 . This uploads run metrics and config — not your training text — to wandb.ai's servers:
Loss/perplexity : train loss, held-out loss, held-out perplexity, every 20 steps
Optimizer health : gradient norm (pre-clip) and total weight norm — the two numbers that would have caught an unstable learning rate immediately, instead of discovering it hours later from bad text output
Gradient/weight histograms per layer ( wandb.watch , every 100 steps) — see exactly which layer is diverging, not just an aggregate number
Sample generations table : a fixed set of probe prompts (greeting, arithmetic, antonym, calendar, identity), regenerated every 100 steps and logged as a table — scroll through training history and watch actual output quality change, instead of manually re-testing after the fact
Run config : architecture, dataset name/size/token count, tokenizer compression ratio, learning rate, batch size — everything needed to reproduce or compare a run
System metrics (CPU/memory) — captured automatically by W&B, no code needed
.venv/bin/pip install wandb # already in requirements.txt
.venv/bin/wandb login # opens a browser to create a free account / paste an API key
.venv/bin/python train.py --data data/training.txt --resume --steps 200 --wandb
The panel's Resume training button auto-adds --wandb once you've logged in (detected via ~/.netrc ), so runs launched from the browser log too.
Run ledger ( runs/ledger.jsonl ): one line per finished, paused, or errored run — architecture, dataset, final perplexity, timestamp. Plain JSON Lines, always written locally regardless of W&B, for a quick offline "what have I actually tried" scan.
checkpoints/status.json still drives the live panel in the browser; W&B and the ledger are for comparing runs after the fact.
Hyperparameter sweeps (Weights & Biases)
train.py --sweep runs an isolated, throwaway training trial under sweeps/ — it never touches checkpoints/ , reuses an existing tokenizer via --tokenizer-path (no re-training BPE per trial), and caches the encoded corpus so many short trials over the same data are fast. Combined with a W&B Sweep, this gives you the multi-run comparison views (parallel coordinates, scatter plots) on the Sweeps tab of your project:
.venv/bin/wandb sweep sweep.yaml # prints a sweep ID
.venv/bin/wandb agent < entity > /aven-1/ < sweep-id >
sweep.yaml searches learning rate, dropout, and batch size on a small, fast architecture (128-wide, 2 layers) over the instruction corpus, minimizing held-out perplexity. Stop the agent (Ctrl-C) whenever you have enough runs — each trial's leftover files live under sweeps/run-*/ and can be deleted freely; they're not your real model.
brain.py : embeddings, causal attention, Transformer blocks, and generation.
train.py : data split, optimization, validation, checkpointing, tokenizer training/reuse, W&B + ledger logging, sweep mode.
tokenizer.py : from-scratch byte-pair encoding, trained on your own corpus.
make_instructions.py : generates the original, programmatic instruction-tuning corpus ( data/instructions.txt ).
server.py and ui.html : local control panel, chat panel, and measured activity display.
memory.py : SQLite-backed notes storage ( memory.db ), separate from model weights.
sources.py : training-data workspace file management ( data/sources/ ) and corpus compilation.
sweep.yaml : Weights & Biases sweep config (learning rate / dropout / batch size search).
checkpoints/latest.pt : your trained weights, optimizer state, and tokenizer ( checkpoints/tokenizer.json ).
checkpoints/status.json : progress and loss history for the live panel.
runs/ledger.jsonl : one-line-per-run summary log.
sweeps/ : throwaway trial checkpoints and the encoded-corpus cache used by --sweep ; safe to delete.
work/train.log : messages from training launched in the panel.
test_brain.py : causal-mask, next-byte-target, learning, and generation tests.
Use an Apple Silicon Python 3.11 environment:
python3.11 -m venv .venv
.venv/bin/python -m pip install -r requirements.txt
.venv/bin/python -m unittest -v
The installed dependencies are pinned in requirements.txt . The first setup downloads software packages; running and training the model does not require internet access. No model weights or outside training datasets are downloaded.
GPU implementation reference: https://docs.pytorch.org/docs/stable/notes/mps.html
Improve the corpus and evaluate on a separate, genuinely different text collection. Then work on tokenization, longer context within the memory budget, and dialogue-specific training/evaluation. Vision and audio would need appropriate encoders, datasets, and objectives; actions require a separate tool-control system. Merely increasing the step count on these

[truncated]

## Original Extract

A from-scratch causal Transformer trained from random weights, no pretrained model or API. - ariveinberg-max/aven-1

GitHub - ariveinberg-max/aven-1: A from-scratch causal Transformer trained from random weights, no pretrained model or API. · GitHub
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
ariveinberg-max
/
aven-1
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
10 Commits 10 Commits Folders and files
data data runs runs .gitignore .gitignore Aven-1-Colab.ipynb Aven-1-Colab.ipynb README.md README.md Start Brain.command Start Brain.command WRITEUP.md WRITEUP.md brain.py brain.py chat.html chat.html make_instructions.py make_instructions.py memory.py memory.py ppo.py ppo.py preferences.py preferences.py raft.py raft.py requirements.txt requirements.txt reward_model.py reward_model.py server.py server.py sources.py sources.py sweep.yaml sweep.yaml test_brain.py test_brain.py tokenizer.py tokenizer.py train.py train.py train_reward.py train_reward.py ui.html ui.html value_model.py value_model.py View all files Repository files navigation
A working first step toward your own AI: a 5,055,360-parameter byte-level Transformer, trained from random weights. No pretrained model, API key, or cloud inference. PyTorch provides tensor math, automatic differentiation, and AdamW; the architecture, training loop, checkpoint handling, and interface are in this project. This is not a general assistant, conscious brain, or a model with voice, vision, web access, or persistent conversational memory.
Double-click Start Brain.command in this folder. Keep the Terminal window open while using the app. Visit http://127.0.0.1:8765 if the browser does not open. Close the server with Control-C in its Terminal window; closing a browser tab does not stop training.
The first verified checkpoint is already available. Click Generate text to try it. Click Resume training to add the chosen number of training steps. Pause & save finishes the current step before saving. The app runs only on this computer.
If macOS does not launch the command file, run these commands in Terminal:
cd /Users/ariveinbers/Documents/Codex/my-ai-brain
.venv/bin/python server.py
What it learns
It predicts the next UTF-8 byte from up to 128 preceding bytes. Four causal Transformer blocks use five attention heads and a hidden width of 320. Batch size 8 and float32 are conservative starting defaults for an 8 GB Mac. GPU availability is checked at runtime; training selects MPS when available. Generation uses CPU to keep the interface simple and memory use modest. Close other heavy apps if your Mac experiences memory pressure.
The included data/demo.txt contains original, programmatically composed practice stories. Many share a template. They are a pipeline test, not a broad knowledge source. The last 10% of bytes are held out of gradient updates, but the similar templates make this an easy validation set. Falling loss on it does not establish general intelligence, useful conversation, or generalization to unrelated text. Output after a short run is expected to be rough.
The diagram shows actual final-position hidden-state magnitudes after generation, grouped into 32 nodes per block (10 channels per node). Brightness is normalized separately within each block. The connecting lines are a schematic, not a display of individual learned weights. It is not a live scan of thoughts.
Use UTF-8 plain text that you own, wrote, or have permission to train on. The starter accepts 4 KB–20 MB. More diverse, clean text is more useful than simply repeating a small document. It reads the entire corpus into memory, so the file-size limit is intentional.
To start a new run while preserving the current one, first pause training, wait for it to finish, and stop the server. Rename checkpoints to an unused name such as checkpoints-demo-backup . Then save your text as data/training.txt and reopen the app. The app will create new random weights for the next training session. The UI also accepts pasted text before a checkpoint exists. Keep backups private if your training text contains personal material.
To continue a CLI run, use the same corpus path every time:
.venv/bin/python train.py --data data/training.txt --steps 200 --resume
For the included demo, omit --data . A checkpoint records the corpus hash and rejects a different corpus on resume. It preserves model weights, optimizer state, step count, and CPU random-generator state. Checkpoints save every 20 steps, at normal completion, and on a handled pause. Force-quitting or losing power can lose steps since the previous save. Run one training process at a time; use the panel or CLI, not both simultaneously.
Training-data workspace vs. stored memory vs. learned weights
The panel now separates three things that are easy to conflate:
Learned weights ( checkpoints/latest.pt ): numbers adjusted by gradient descent. This is the only thing that makes the model's text generation change. You cannot edit these directly; you can only change them by training.
Training-data workspace ( data/sources/*.txt , managed from the "Training-data workspace" panel): the plain text a future run would learn from. Add, edit, or delete named sources here, then click Compile into training corpus to concatenate them into data/training.txt . Once a checkpoint exists, the corpus is locked (same rule as before) — compiling, saving, or deleting a source is refused until you preserve checkpoints/ and start a new run.
Memory ( memory.db , managed from the "Memory" panel): a plain SQLite table of notes you write, edit, and delete directly through the UI. It is not read by training or generation, and saving a note never changes a single model weight. It exists purely as external, inspectable storage — the opposite of the model's opaque learned parameters.
Deleting a memory record or a workspace source is permanent immediately; there is no undo.
Weights & Biases (cloud, free tier, project aven-1 ): pass --wandb to train.py for full run tracking at https://wandb.ai/ariveinberg-ari-research/aven-1 . This uploads run metrics and config — not your training text — to wandb.ai's servers:
Loss/perplexity : train loss, held-out loss, held-out perplexity, every 20 steps
Optimizer health : gradient norm (pre-clip) and total weight norm — the two numbers that would have caught an unstable learning rate immediately, instead of discovering it hours later from bad text output
Gradient/weight histograms per layer ( wandb.watch , every 100 steps) — see exactly which layer is diverging, not just an aggregate number
Sample generations table : a fixed set of probe prompts (greeting, arithmetic, antonym, calendar, identity), regenerated every 100 steps and logged as a table — scroll through training history and watch actual output quality change, instead of manually re-testing after the fact
Run config : architecture, dataset name/size/token count, tokenizer compression ratio, learning rate, batch size — everything needed to reproduce or compare a run
System metrics (CPU/memory) — captured automatically by W&B, no code needed
.venv/bin/pip install wandb # already in requirements.txt
.venv/bin/wandb login # opens a browser to create a free account / paste an API key
.venv/bin/python train.py --data data/training.txt --resume --steps 200 --wandb
The panel's Resume training button auto-adds --wandb once you've logged in (detected via ~/.netrc ), so runs launched from the browser log too.
Run ledger ( runs/ledger.jsonl ): one line per finished, paused, or errored run — architecture, dataset, final perplexity, timestamp. Plain JSON Lines, always written locally regardless of W&B, for a quick offline "what have I actually tried" scan.
checkpoints/status.json still drives the live panel in the browser; W&B and the ledger are for comparing runs after the fact.
Hyperparameter sweeps (Weights & Biases)
train.py --sweep runs an isolated, throwaway training trial under sweeps/ — it never touches checkpoints/ , reuses an existing tokenizer via --tokenizer-path (no re-training BPE per trial), and caches the encoded corpus so many short trials over the same data are fast. Combined with a W&B Sweep, this gives you the multi-run comparison views (parallel coordinates, scatter plots) on the Sweeps tab of your project:
.venv/bin/wandb sweep sweep.yaml # prints a sweep ID
.venv/bin/wandb agent < entity > /aven-1/ < sweep-id >
sweep.yaml searches learning rate, dropout, and batch size on a small, fast architecture (128-wide, 2 layers) over the instruction corpus, minimizing held-out perplexity. Stop the agent (Ctrl-C) whenever you have enough runs — each trial's leftover files live under sweeps/run-*/ and can be deleted freely; they're not your real model.
brain.py : embeddings, causal attention, Transformer blocks, and generation.
train.py : data split, optimization, validation, checkpointing, tokenizer training/reuse, W&B + ledger logging, sweep mode.
tokenizer.py : from-scratch byte-pair encoding, trained on your own corpus.
make_instructions.py : generates the original, programmatic instruction-tuning corpus ( data/instructions.txt ).
server.py and ui.html : local control panel, chat panel, and measured activity display.
memory.py : SQLite-backed notes storage ( memory.db ), separate from model weights.
sources.py : training-data workspace file management ( data/sources/ ) and corpus compilation.
sweep.yaml : Weights & Biases sweep config (learning rate / dropout / batch size search).
checkpoints/latest.pt : your trained weights, optimizer state, and tokenizer ( checkpoints/tokenizer.json ).
checkpoints/status.json : progress and loss history for the live panel.
runs/ledger.jsonl : one-line-per-run summary log.
sweeps/ : throwaway trial checkpoints and the encoded-corpus cache used by --sweep ; safe to delete.
work/train.log : messages from training launched in the panel.
test_brain.py : causal-mask, next-byte-target, learning, and generation tests.
Use an Apple Silicon Python 3.11 environment:
python3.11 -m venv .venv
.venv/bin/python -m pip install -r requirements.txt
.venv/bin/python -m unittest -v
The installed dependencies are pinned in requirements.txt . The first setup downloads software packages; running and training the model does not require internet access. No model weights or outside training datasets are downloaded.
GPU implementation reference: https://docs.pytorch.org/docs/stable/notes/mps.html
Improve the corpus and evaluate on a separate, genuinely different text collection. Then work on tokenization, longer context within the memory budget, and dialogue-specific training/evaluation. Vision and audio would need appropriate encoders, datasets, and objectives; actions require a separate tool-control system. Merely increasing the step count on these

[truncated]
