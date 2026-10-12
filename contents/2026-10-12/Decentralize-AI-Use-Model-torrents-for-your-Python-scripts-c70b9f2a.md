---
source: "https://twitter.com/thepirateface/status/2109691589954310152"
hn_url: "https://news.ycombinator.com/item?id=50061664"
title: "Decentralize AI: Use Model torrents for your Python scripts"
article_title: "Pirate Face on X: \"https://t.co/KMt2I9h5KY\" / X"
image: "https://pbs.twimg.com/media/HUchCRibIAALunT?format=webp&name=large"
author: "llmbababoom"
captured_at: "2026-10-12T20:08:26Z"
capture_tool: "hn-digest"
hn_id: 50061664
score: 3
comments: 0
posted_at: "2026-10-12T19:29:04Z"
tags:
  - hacker-news
---

# Decentralize AI: Use Model torrents for your Python scripts

- HN: [50061664](https://news.ycombinator.com/item?id=50061664)
- Source: [twitter.com](https://twitter.com/thepirateface/status/2109691589954310152)
- Score: 3
- Comments: 0
- Posted: 2026-10-12T19:29:04Z

## Translation

Title: Decentralize AI: Use Model torrents for your Python scripts
Article title: Pirate Face on X: "https://t.co/KMt2I9h5KY" / X
Description: Introducing Pirate Face Hub: Download Hugging Face Models Directly From the Swarm

Article text:
Pirate Face on X: "https://t.co/KMt2I9h5KY" / X Article
Introducing Pirate Face Hub: Download Hugging Face Models Directly From the Swarm
Download AI models from your terminal or existing Python workflows through the Pirate Face swarm. No code changes.
Pirate Face now lists 5,149 models as torrents, most of them added by the community. Until today, using one meant leaving your normal workflow: open the model page, copy the magnet, run a torrent client, then point your code at the files.
When Pirate Face reached the front page of Hacker News , commenters agreed that BitTorrent is the right format for model weights. The friction was everywhere else: tooling, and trusting files from strangers. Pirate Face Hub removes both. Your existing code stays exactly as it is.
Pirate Face Hub includes a pf download command. With uv installed, run:
uvx -p 3.13 --from pirateface pf download org/model Replace org/model with the Hugging Face model ID you want to download. The command downloads the model into your Hugging Face cache, using the swarm when peers are available and falling back to Hugging Face when needed. No magnets to copy or separate torrent client to open.
The command selects Python 3.13 because the package currently supports Python 3.12 and 3.13. The package README documents the CLI and Python requirements.
Pirate Face Hub works with transformers, diffusers, sentence-transformers, and anything that downloads through huggingface_hub.
Run any script with the package added:
uv run --with pirateface python your_script.py Or install it once and forget about it:
pip install pirateface That is the whole setup. There is no config, no new API, and no environment variable to set. Code that already loads models keeps working, and the files now come from the swarm:
from transformers import AutoModelForCausalLM
model = AutoModelForCausalLM. from_pretrained ( "org/model" ) # unchanged That includes hf_hub_download and snapshot_download. The Hugging Face hf download command also works when Pirate Face Hub is installed in the same Python environment.
Hugging Face still tells your code which files a model has and what each one should look like. Pirate Face Hub changes only where the bytes come from.
For each file, the package asks Pirate Face for the model's torrent and downloads the file from peers in the swarm.
If Pirate Face has no torrent for the model, or no peer answers within 30 seconds, that file comes from Hugging Face as usual.
Downloading over BitTorrent works like any torrent client: other peers can see your IP address. The seeding guide covers what that means in practice.
Every missing model joins the queue
When you download a model Pirate Face does not have yet, the package sends one request in the background asking Pirate Face to rescue it. It sends only the model name. Pirate Face's own workers download the model from Hugging Face, check every file, build the torrent and list it. You can watch the queue on the Rescue board .
To turn requests off, set PF_NO_RESCUE=1.
Contributing has never been easier
Every download through Pirate Face Hub uses the public swarm, and every model it does not find yet becomes a listed torrent that anyone can seed. Add one package to the scripts you already run, and they become part of the infrastructure behind Making Models Permanent .
Pirate Face Hub is open source under the MIT license. The code is on GitHub and the package is on PyPI . It needs Python 3.12 or 3.13.
See what’s happening and join the conversation
Continue with phone Continue with Apple Continue with Google or Log in with username or email Relevant people
© 2026 X Corp. Pirate Face @thepirateface Introducing Pirate Face Hub: Download Hugging Face Models Directly From the Swarm
Download AI models from your terminal or existing Python workflows through the Pirate Face swarm. No code changes.
Pirate Face now lists 5,149 models as torrents, most of them added by the community. Until today, using one meant leaving your normal workflow: open the model page, copy the magnet, run a torrent client, then point your code at the files.
When Pirate Face reached the front page of Hacker News , commenters agreed that BitTorrent is the right format for model weights. The friction was everywhere else: tooling, and trusting files from strangers. Pirate Face Hub removes both. Your existing code stays exactly as it is.
Pirate Face Hub includes a pf download command. With uv installed, run:
uvx -p 3.13 --from pirateface pf download org/model Replace org/model with the Hugging Face model ID you want to download. The command downloads the model into your Hugging Face cache, using the swarm when peers are available and falling back to Hugging Face when needed. No magnets to copy or separate torrent client to open.
The command selects Python 3.13 because the package currently supports Python 3.12 and 3.13. The package README documents the CLI and Python requirements.
Pirate Face Hub works with transformers, diffusers, sentence-transformers, and anything that downloads through huggingface_hub.
Run any script with the package added:
uv run --with pirateface python your_script.py Or install it once and forget about it:
pip install pirateface That is the whole setup. There is no config, no new API, and no environment variable to set. Code that already loads models keeps working, and the files now come from the swarm:
from transformers import AutoModelForCausalLM
model = AutoModelForCausalLM. from_pretrained ( "org/model" ) # unchanged That includes hf_hub_download and snapshot_download. The Hugging Face hf download command also works when Pirate Face Hub is installed in the same Python environment.
Hugging Face still tells your code which files a model has and what each one should look like. Pirate Face Hub changes only where the bytes come from.
For each file, the package asks Pirate Face for the model's torrent and downloads the file from peers in the swarm.
If Pirate Face has no torrent for the model, or no peer answers within 30 seconds, that file comes from Hugging Face as usual.
Downloading over BitTorrent works like any torrent client: other peers can see your IP address. The seeding guide covers what that means in practice.
Every missing model joins the queue
When you download a model Pirate Face does not have yet, the package sends one request in the background asking Pirate Face to rescue it. It sends only the model name. Pirate Face's own workers download the model from Hugging Face, check every file, build the torrent and list it. You can watch the queue on the Rescue board .
To turn requests off, set PF_NO_RESCUE=1.
Contributing has never been easier
Every download through Pirate Face Hub uses the public swarm, and every model it does not find yet becomes a listed torrent that anyone can seed. Add one package to the scripts you already run, and they become part of the infrastructure behind Making Models Permanent .
Pirate Face Hub is open source under the MIT license. The code is on GitHub and the package is on PyPI . It needs Python 3.12 or 3.13.
Dodge This Security @shotgunner101 2h Can you guys offer a daily/weekly archive of all of the torrent files/magnet links compressed? That way we can archive all of them in case your site ever gets taken down due to future legislation? 2 1 2 46

## Original Extract

Introducing Pirate Face Hub: Download Hugging Face Models Directly From the Swarm

Pirate Face on X: "https://t.co/KMt2I9h5KY" / X Article
Introducing Pirate Face Hub: Download Hugging Face Models Directly From the Swarm
Download AI models from your terminal or existing Python workflows through the Pirate Face swarm. No code changes.
Pirate Face now lists 5,149 models as torrents, most of them added by the community. Until today, using one meant leaving your normal workflow: open the model page, copy the magnet, run a torrent client, then point your code at the files.
When Pirate Face reached the front page of Hacker News , commenters agreed that BitTorrent is the right format for model weights. The friction was everywhere else: tooling, and trusting files from strangers. Pirate Face Hub removes both. Your existing code stays exactly as it is.
Pirate Face Hub includes a pf download command. With uv installed, run:
uvx -p 3.13 --from pirateface pf download org/model Replace org/model with the Hugging Face model ID you want to download. The command downloads the model into your Hugging Face cache, using the swarm when peers are available and falling back to Hugging Face when needed. No magnets to copy or separate torrent client to open.
The command selects Python 3.13 because the package currently supports Python 3.12 and 3.13. The package README documents the CLI and Python requirements.
Pirate Face Hub works with transformers, diffusers, sentence-transformers, and anything that downloads through huggingface_hub.
Run any script with the package added:
uv run --with pirateface python your_script.py Or install it once and forget about it:
pip install pirateface That is the whole setup. There is no config, no new API, and no environment variable to set. Code that already loads models keeps working, and the files now come from the swarm:
from transformers import AutoModelForCausalLM
model = AutoModelForCausalLM. from_pretrained ( "org/model" ) # unchanged That includes hf_hub_download and snapshot_download. The Hugging Face hf download command also works when Pirate Face Hub is installed in the same Python environment.
Hugging Face still tells your code which files a model has and what each one should look like. Pirate Face Hub changes only where the bytes come from.
For each file, the package asks Pirate Face for the model's torrent and downloads the file from peers in the swarm.
If Pirate Face has no torrent for the model, or no peer answers within 30 seconds, that file comes from Hugging Face as usual.
Downloading over BitTorrent works like any torrent client: other peers can see your IP address. The seeding guide covers what that means in practice.
Every missing model joins the queue
When you download a model Pirate Face does not have yet, the package sends one request in the background asking Pirate Face to rescue it. It sends only the model name. Pirate Face's own workers download the model from Hugging Face, check every file, build the torrent and list it. You can watch the queue on the Rescue board .
To turn requests off, set PF_NO_RESCUE=1.
Contributing has never been easier
Every download through Pirate Face Hub uses the public swarm, and every model it does not find yet becomes a listed torrent that anyone can seed. Add one package to the scripts you already run, and they become part of the infrastructure behind Making Models Permanent .
Pirate Face Hub is open source under the MIT license. The code is on GitHub and the package is on PyPI . It needs Python 3.12 or 3.13.
See what’s happening and join the conversation
Continue with phone Continue with Apple Continue with Google or Log in with username or email Relevant people
© 2026 X Corp. Pirate Face @thepirateface Introducing Pirate Face Hub: Download Hugging Face Models Directly From the Swarm
Download AI models from your terminal or existing Python workflows through the Pirate Face swarm. No code changes.
Pirate Face now lists 5,149 models as torrents, most of them added by the community. Until today, using one meant leaving your normal workflow: open the model page, copy the magnet, run a torrent client, then point your code at the files.
When Pirate Face reached the front page of Hacker News , commenters agreed that BitTorrent is the right format for model weights. The friction was everywhere else: tooling, and trusting files from strangers. Pirate Face Hub removes both. Your existing code stays exactly as it is.
Pirate Face Hub includes a pf download command. With uv installed, run:
uvx -p 3.13 --from pirateface pf download org/model Replace org/model with the Hugging Face model ID you want to download. The command downloads the model into your Hugging Face cache, using the swarm when peers are available and falling back to Hugging Face when needed. No magnets to copy or separate torrent client to open.
The command selects Python 3.13 because the package currently supports Python 3.12 and 3.13. The package README documents the CLI and Python requirements.
Pirate Face Hub works with transformers, diffusers, sentence-transformers, and anything that downloads through huggingface_hub.
Run any script with the package added:
uv run --with pirateface python your_script.py Or install it once and forget about it:
pip install pirateface That is the whole setup. There is no config, no new API, and no environment variable to set. Code that already loads models keeps working, and the files now come from the swarm:
from transformers import AutoModelForCausalLM
model = AutoModelForCausalLM. from_pretrained ( "org/model" ) # unchanged That includes hf_hub_download and snapshot_download. The Hugging Face hf download command also works when Pirate Face Hub is installed in the same Python environment.
Hugging Face still tells your code which files a model has and what each one should look like. Pirate Face Hub changes only where the bytes come from.
For each file, the package asks Pirate Face for the model's torrent and downloads the file from peers in the swarm.
If Pirate Face has no torrent for the model, or no peer answers within 30 seconds, that file comes from Hugging Face as usual.
Downloading over BitTorrent works like any torrent client: other peers can see your IP address. The seeding guide covers what that means in practice.
Every missing model joins the queue
When you download a model Pirate Face does not have yet, the package sends one request in the background asking Pirate Face to rescue it. It sends only the model name. Pirate Face's own workers download the model from Hugging Face, check every file, build the torrent and list it. You can watch the queue on the Rescue board .
To turn requests off, set PF_NO_RESCUE=1.
Contributing has never been easier
Every download through Pirate Face Hub uses the public swarm, and every model it does not find yet becomes a listed torrent that anyone can seed. Add one package to the scripts you already run, and they become part of the infrastructure behind Making Models Permanent .
Pirate Face Hub is open source under the MIT license. The code is on GitHub and the package is on PyPI . It needs Python 3.12 or 3.13.
Dodge This Security @shotgunner101 2h Can you guys offer a daily/weekly archive of all of the torrent files/magnet links compressed? That way we can archive all of them in case your site ever gets taken down due to future legislation? 2 1 2 46
