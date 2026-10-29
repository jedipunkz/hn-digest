---
source: "https://oxide.computer/blog/self-hosted-ai-agents"
hn_url: "https://news.ycombinator.com/item?id=50290958"
title: "Running Self-Hosted AI Agents on Oxide"
article_title: "Running Self-Hosted AI Agents | Oxide Computer Company"
image: "https://oxide-computer-dntj3558l-oxidecomputer.vercel.app/img/meta/og-blog.png"
author: "chmaynard"
captured_at: "2026-10-29T16:40:07Z"
capture_tool: "hn-digest"
hn_id: 50290958
score: 3
comments: 0
posted_at: "2026-10-29T15:47:58Z"
tags:
  - hacker-news
---

# Running Self-Hosted AI Agents on Oxide

- HN: [50290958](https://news.ycombinator.com/item?id=50290958)
- Source: [oxide.computer](https://oxide.computer/blog/self-hosted-ai-agents)
- Score: 3
- Comments: 0
- Posted: 2026-10-29T15:47:58Z

## Translation

Title: Running Self-Hosted AI Agents on Oxide
Article title: Running Self-Hosted AI Agents | Oxide Computer Company
Description: The journey of running Claude Code and Amp self-hosted AI agents on the Oxide Cloud Computer as ephemeral virtual machines.

Article text:
Running Self-Hosted AI Agents | Oxide Computer Company Product
Customers are increasingly asking us the following question:
How do I run AI agents on Oxide?
The initial answer is simple. You run AI agents the same way you run other
workloads on Oxide—​in virtual machines. Customers already run long-lived
services in virtual machines, so running AI agents feels familiar.
However, AI agents aren’t exactly long-lived services—​they fall somewhere
between that and ephemeral CI/CD jobs, depending on how they are used. Run a
long-lived agent that provisions sub-agents and it looks like the former. Run an
agent to complete a bounded task using /goal and it looks like the latter.
Regardless of your use case, the workflow looks like this:
Run the AI agent in the virtual machine.
With that context, the more accurate answer is to run the AI agent’s
orchestrator as a virtual machine, configured to create ephemeral virtual
machines for each AI agent session. I’ll show how this works on our Oxide Cloud
Computer using both Claude Code self-hosted environments and Amp self-hosted
orbs.
Before describing how to run AI agents on an Oxide Cloud Computer, it’s
important to understand why AI agents need CPUs in the first place.
AI agents are just systems that perform inference and execute tool calls in
a loop. Put another way, AI becomes "agentic" when GPU-based inference needs
CPU-based tool execution.
The AI agent harness (e.g., Claude Code, Codex, Amp) is responsible for running
this loop. The harness runs as a process on a CPU, communicating with an LLM
over the network. When the LLM requests a tool call, the harness executes it on
the CPU it’s running on and sends the result back to the LLM.
Without a CPU, you can’t run the harness or execute tool calls.
Virtual machines, not containers
AI agents need strong isolation for both security and performance. You don’t
want an AI agent to delete important data or duel with other AI agents and
disrupt their execution.
Virtual machines provide stronger isolation than containers, making them
well-suited for AI agents. This isolation comes at a cost, summarized by this
table:
Using virtual machines instead of containers for workload orchestration raises
the following questions:
How do I provision the virtual machine running the AI agent process?
How do I clean up the virtual machine when the AI agent process finishes?
These questions are surprisingly difficult to work through, especially
when documentation leaves out the details. Look at this quote from the
Claude Code
self-hosted environments documentation .
Ephemeral, per-session containers: run each runner process in a fresh container
or VM that’s destroyed when the process exits
Not only does this push readers towards containers, and possibly worse isolation
as a result, but it also has a tone of "just do the right thing" without telling
readers how to do it. Excuse me, Claude, did you just tell me to go eff myself?
Needless to say, when our customers ask how to run AI agents on Oxide, they also
ask how to do it specifically with virtual machines, not containers.
Let’s see how to run AI agents on our Oxide Cloud
Computer using virtual machines, specifically using
Claude
Code self-hosted environments and
Amp self-hosted orbs . These were
chosen due to their popularity in the industry and to showcase how different AI
agents use certain Oxide features.
Claude Code self-hosted environments
Claude
Code self-hosted environments run Claude Code sessions on infrastructure you
control, in this case the Oxide Cloud Computer. The orchestrator dynamically
provisions ephemeral Oxide virtual machines to run Claude Code and deletes them
when finished.
The orchestrator virtual machine is the central component. It’s a long-running
virtual machine responsible for addressing the complexity mentioned in the
previous section. Recall these questions:
How do I provision the virtual machine running the AI agent process?
How do I clean up the virtual machine when the AI agent process finishes?
For provisioning, the orchestrator virtual machine runs a claude command
you’ve likely never used:
claude self-hosted-runner orchestrator \
--environment-secret-file /etc/claude-self-hosted/environment-secret \
--hooks-dir /etc/claude-self-hosted/hooks \
--hook-concurrency 1 \
--hook-timeout 60 \
--expected-spawn-seconds 300 This command authenticates to the self-hosted environment using the
--environment-secret-file option and waits for work orders, which are requests
to create an ephemeral runner for a user’s Claude Code session. When there are
no work orders, the number of ephemeral runners scales down to zero.
The --hooks-dir option is, to borrow an LLM-ism, "load-bearing." It points to
a directory containing hooks the orchestrator will call when work orders come
in. The only hook available is spawn-runner , which, when called, must create
the ephemeral virtual machine running the AI agent process. The AI agent process
is another claude command you’ve likely never used:
claude self-hosted-runner \
--capacity 1 \
--base-dir /var/lib/claude-self-hosted \
--use-anthropic-git-proxy \
--configure-git The --capacity 1 option ensures the process only runs a single
Claude Code session to support our ephemeral architecture, while the
--use-anthropic-git-proxy and --configure-git options ensure that your
connected Git repositories can be safely cloned within the virtual machine.
To make cleanup easier, the ephemeral virtual machine is configured to power
off once the AI agent process has finished. There are a few ways to do this, but
a popular way is to create a systemd service that runs at startup. The service
then runs the AI agent process and runs poweroff after it exits.
A stopped virtual machine is the signal used for cleanup. Unfortunately, the
orchestrator does not have a cleanup-runner hook, so the burden of cleanup is
on us. I chose to run a separate reaper process that periodically polls the
Oxide API for stopped virtual machines matching the desired name prefix. When
it finds a matching virtual machine, it deletes it and its boot disk, freeing up
resources for more Claude Code sessions.
If everything is wired up correctly, you’ll be able to select your self-hosted
environment when you create a Claude Code session on https://claude.ai . You
could also run claude with the --environment option, though you’ll need
to ask your Claude administrator for the self-hosted environment ID to pass as
its argument.
If you want to try this yourself, I’ve put together a showcase describing
how to use Claude Code self-hosted environments with Oxide. Find the code at
oxidecomputer/showcase .
Amp self-hosted orbs run Amp
threads on infrastructure you control, in this case the Oxide Cloud Computer.
The provider dynamically provisions ephemeral Oxide virtual machines to run Amp
orbs and deletes them when finished.
The Amp orb provider runs an amp command that authenticates to Amp and waits
for requests to spawn, pause, and delete Amp orbs:
AMP_ORB_PROVIDER_TOKEN = 'aop_live...' amp orb provider dispatch.sh The dispatch.sh script can contain any logic you wish, as long as it adheres
to the dispatch
contract defined by Amp, which is really just responding to the ORB_OPERATION
values spawn , pause , and delete . Here’s a pseudocode version of the
dispatch script I used:
remove_instance () {
oxide instance stop --project " $PROJECT " --instance " $orb_name "
oxide instance delete --project " $PROJECT " --instance " $orb_name "
}
case " $ORB_OPERATION " in
spawn )
remove_instance
if oxide disk view --project " $PROJECT " --disk " $orb_name " ; then
boot_disk = $( attach_disk " $orb_name " ) # resume from orb's disk
else
boot_disk = $( new_disk " $orb_name " " $ORB_IMAGE " ) # new orb from the image
fi
# User data starts `amp orb bootstrap` as `user` with this spawn's ORB_TOKEN.
oxide instance create --project " $PROJECT " \
--json-body <( instance_body " $boot_disk " " $ORB_TOKEN ")
;;
pause )
remove_instance # the disk is kept
;;
delete )
remove_instance
oxide disk delete --project " $PROJECT " --disk " $orb_name "
;;
esac Amp’s pause and delete operations free us from relying on a stopped virtual
machine as the signal for cleanup. Instead, we can delete the instance but
keep its boot disk for pause and delete the instance and its boot disk for
delete . This significantly simplifies the overall architecture and deployment
compared to Claude Code self-hosted environments.
If everything is wired up correctly, you’ll be able to select your self-hosted
orb provider when you create an Amp thread on https://ampcode.com . You could
also run amp with the --orb-provider option, using the provider name as
the argument.
If you want to try this yourself, I’ve put together a showcase
describing how to use Amp self-hosted orbs with Oxide. Find the code at
oxidecomputer/showcase .
Our Oxide Cloud Computer has the right primitives for today’s AI agents—​virtual
machines with strong isolation, disks with solid performance, control over VPC
networking, and an elastic API to dynamically provision everything. Even so,
we’re looking ahead to meet the needs of tomorrow’s AI agents.
Conversations with current and prospective customers about AI agents have
centered around the following themes:
Security is where most conversations start, especially as the number of
reported CVEs increases across the industry. We’ve already discussed the
strong isolation provided by virtual machines, but our security story
doesn’t stop there. Our Oxide Cloud Computer has silos for tenancy and
resource quota enforcement, projects for network isolation, and VPC firewall
rules to block undesirable network traffic. Additionally, our removal of
the BIOS, along with the virtual machine attestation work discussed in
RFD 605 , provides a secure
foundation for AI agent workloads.
If you run many AI agents, you may be wondering about scale. You can create as
many virtual machines as you can fit on your Oxide Cloud Computer, but we’re
working on multi-rack capabilities to expand a silo’s capacity beyond a single
rack. Once released, this will simplify the management of AI agents across
multiple racks.
Our API makes virtual machine provisioning trivial from an operational
standpoint, but it can take tens of seconds or even minutes after an API request
before a virtual machine is scheduled, booted, and running a workload. This
affects users in the form of cold-start latency, which a warm pool can help
mitigate, but not entirely eliminate. We accepted the latency in this blog post,
but our goal is to reduce it as much as possible.
Cleaning up virtual machines is also trivial using the Oxide API, but, depending
on the AI agent, there isn’t always a simple way to know when the workload
has finished. In this blog post, we handled this by treating a stopped virtual
machine as the signal that the workload has finished. This required us to wrap
our AI agent process in a systemd service that runs poweroff after the AI
agent process exits, along with a reaper process to poll the Oxide API for
virtual machines to clean up. Not the best user experience, especially if the AI
agent architecture doesn’t provide a delete operation, but we’re exploring how
to make this experience behave like a container that exits when the main process
is finished.
I’ve made it this far without talking about microVMs, but the microVM primitive
is exactly what some infrastructure providers are using to improve the user
experience of running AI agents without sacrificing the strong isolation they’ve
come to rely on from virtual machines.
The definition of microVMs can vary across the industry, but discussing them in
the context of AI agents has led to internal conversations about what a microVM
primitive could look like for Oxide. Our hypervisor team even spiked on work
to see how quickly we could launch a virtual machine and have it running its
workload. Turns out we’re not far off from popular microVM implementations!
We don’t have an

[truncated]

## Original Extract

The journey of running Claude Code and Amp self-hosted AI agents on the Oxide Cloud Computer as ephemeral virtual machines.

Running Self-Hosted AI Agents | Oxide Computer Company Product
Customers are increasingly asking us the following question:
How do I run AI agents on Oxide?
The initial answer is simple. You run AI agents the same way you run other
workloads on Oxide—​in virtual machines. Customers already run long-lived
services in virtual machines, so running AI agents feels familiar.
However, AI agents aren’t exactly long-lived services—​they fall somewhere
between that and ephemeral CI/CD jobs, depending on how they are used. Run a
long-lived agent that provisions sub-agents and it looks like the former. Run an
agent to complete a bounded task using /goal and it looks like the latter.
Regardless of your use case, the workflow looks like this:
Run the AI agent in the virtual machine.
With that context, the more accurate answer is to run the AI agent’s
orchestrator as a virtual machine, configured to create ephemeral virtual
machines for each AI agent session. I’ll show how this works on our Oxide Cloud
Computer using both Claude Code self-hosted environments and Amp self-hosted
orbs.
Before describing how to run AI agents on an Oxide Cloud Computer, it’s
important to understand why AI agents need CPUs in the first place.
AI agents are just systems that perform inference and execute tool calls in
a loop. Put another way, AI becomes "agentic" when GPU-based inference needs
CPU-based tool execution.
The AI agent harness (e.g., Claude Code, Codex, Amp) is responsible for running
this loop. The harness runs as a process on a CPU, communicating with an LLM
over the network. When the LLM requests a tool call, the harness executes it on
the CPU it’s running on and sends the result back to the LLM.
Without a CPU, you can’t run the harness or execute tool calls.
Virtual machines, not containers
AI agents need strong isolation for both security and performance. You don’t
want an AI agent to delete important data or duel with other AI agents and
disrupt their execution.
Virtual machines provide stronger isolation than containers, making them
well-suited for AI agents. This isolation comes at a cost, summarized by this
table:
Using virtual machines instead of containers for workload orchestration raises
the following questions:
How do I provision the virtual machine running the AI agent process?
How do I clean up the virtual machine when the AI agent process finishes?
These questions are surprisingly difficult to work through, especially
when documentation leaves out the details. Look at this quote from the
Claude Code
self-hosted environments documentation .
Ephemeral, per-session containers: run each runner process in a fresh container
or VM that’s destroyed when the process exits
Not only does this push readers towards containers, and possibly worse isolation
as a result, but it also has a tone of "just do the right thing" without telling
readers how to do it. Excuse me, Claude, did you just tell me to go eff myself?
Needless to say, when our customers ask how to run AI agents on Oxide, they also
ask how to do it specifically with virtual machines, not containers.
Let’s see how to run AI agents on our Oxide Cloud
Computer using virtual machines, specifically using
Claude
Code self-hosted environments and
Amp self-hosted orbs . These were
chosen due to their popularity in the industry and to showcase how different AI
agents use certain Oxide features.
Claude Code self-hosted environments
Claude
Code self-hosted environments run Claude Code sessions on infrastructure you
control, in this case the Oxide Cloud Computer. The orchestrator dynamically
provisions ephemeral Oxide virtual machines to run Claude Code and deletes them
when finished.
The orchestrator virtual machine is the central component. It’s a long-running
virtual machine responsible for addressing the complexity mentioned in the
previous section. Recall these questions:
How do I provision the virtual machine running the AI agent process?
How do I clean up the virtual machine when the AI agent process finishes?
For provisioning, the orchestrator virtual machine runs a claude command
you’ve likely never used:
claude self-hosted-runner orchestrator \
--environment-secret-file /etc/claude-self-hosted/environment-secret \
--hooks-dir /etc/claude-self-hosted/hooks \
--hook-concurrency 1 \
--hook-timeout 60 \
--expected-spawn-seconds 300 This command authenticates to the self-hosted environment using the
--environment-secret-file option and waits for work orders, which are requests
to create an ephemeral runner for a user’s Claude Code session. When there are
no work orders, the number of ephemeral runners scales down to zero.
The --hooks-dir option is, to borrow an LLM-ism, "load-bearing." It points to
a directory containing hooks the orchestrator will call when work orders come
in. The only hook available is spawn-runner , which, when called, must create
the ephemeral virtual machine running the AI agent process. The AI agent process
is another claude command you’ve likely never used:
claude self-hosted-runner \
--capacity 1 \
--base-dir /var/lib/claude-self-hosted \
--use-anthropic-git-proxy \
--configure-git The --capacity 1 option ensures the process only runs a single
Claude Code session to support our ephemeral architecture, while the
--use-anthropic-git-proxy and --configure-git options ensure that your
connected Git repositories can be safely cloned within the virtual machine.
To make cleanup easier, the ephemeral virtual machine is configured to power
off once the AI agent process has finished. There are a few ways to do this, but
a popular way is to create a systemd service that runs at startup. The service
then runs the AI agent process and runs poweroff after it exits.
A stopped virtual machine is the signal used for cleanup. Unfortunately, the
orchestrator does not have a cleanup-runner hook, so the burden of cleanup is
on us. I chose to run a separate reaper process that periodically polls the
Oxide API for stopped virtual machines matching the desired name prefix. When
it finds a matching virtual machine, it deletes it and its boot disk, freeing up
resources for more Claude Code sessions.
If everything is wired up correctly, you’ll be able to select your self-hosted
environment when you create a Claude Code session on https://claude.ai . You
could also run claude with the --environment option, though you’ll need
to ask your Claude administrator for the self-hosted environment ID to pass as
its argument.
If you want to try this yourself, I’ve put together a showcase describing
how to use Claude Code self-hosted environments with Oxide. Find the code at
oxidecomputer/showcase .
Amp self-hosted orbs run Amp
threads on infrastructure you control, in this case the Oxide Cloud Computer.
The provider dynamically provisions ephemeral Oxide virtual machines to run Amp
orbs and deletes them when finished.
The Amp orb provider runs an amp command that authenticates to Amp and waits
for requests to spawn, pause, and delete Amp orbs:
AMP_ORB_PROVIDER_TOKEN = 'aop_live...' amp orb provider dispatch.sh The dispatch.sh script can contain any logic you wish, as long as it adheres
to the dispatch
contract defined by Amp, which is really just responding to the ORB_OPERATION
values spawn , pause , and delete . Here’s a pseudocode version of the
dispatch script I used:
remove_instance () {
oxide instance stop --project " $PROJECT " --instance " $orb_name "
oxide instance delete --project " $PROJECT " --instance " $orb_name "
}
case " $ORB_OPERATION " in
spawn )
remove_instance
if oxide disk view --project " $PROJECT " --disk " $orb_name " ; then
boot_disk = $( attach_disk " $orb_name " ) # resume from orb's disk
else
boot_disk = $( new_disk " $orb_name " " $ORB_IMAGE " ) # new orb from the image
fi
# User data starts `amp orb bootstrap` as `user` with this spawn's ORB_TOKEN.
oxide instance create --project " $PROJECT " \
--json-body <( instance_body " $boot_disk " " $ORB_TOKEN ")
;;
pause )
remove_instance # the disk is kept
;;
delete )
remove_instance
oxide disk delete --project " $PROJECT " --disk " $orb_name "
;;
esac Amp’s pause and delete operations free us from relying on a stopped virtual
machine as the signal for cleanup. Instead, we can delete the instance but
keep its boot disk for pause and delete the instance and its boot disk for
delete . This significantly simplifies the overall architecture and deployment
compared to Claude Code self-hosted environments.
If everything is wired up correctly, you’ll be able to select your self-hosted
orb provider when you create an Amp thread on https://ampcode.com . You could
also run amp with the --orb-provider option, using the provider name as
the argument.
If you want to try this yourself, I’ve put together a showcase
describing how to use Amp self-hosted orbs with Oxide. Find the code at
oxidecomputer/showcase .
Our Oxide Cloud Computer has the right primitives for today’s AI agents—​virtual
machines with strong isolation, disks with solid performance, control over VPC
networking, and an elastic API to dynamically provision everything. Even so,
we’re looking ahead to meet the needs of tomorrow’s AI agents.
Conversations with current and prospective customers about AI agents have
centered around the following themes:
Security is where most conversations start, especially as the number of
reported CVEs increases across the industry. We’ve already discussed the
strong isolation provided by virtual machines, but our security story
doesn’t stop there. Our Oxide Cloud Computer has silos for tenancy and
resource quota enforcement, projects for network isolation, and VPC firewall
rules to block undesirable network traffic. Additionally, our removal of
the BIOS, along with the virtual machine attestation work discussed in
RFD 605 , provides a secure
foundation for AI agent workloads.
If you run many AI agents, you may be wondering about scale. You can create as
many virtual machines as you can fit on your Oxide Cloud Computer, but we’re
working on multi-rack capabilities to expand a silo’s capacity beyond a single
rack. Once released, this will simplify the management of AI agents across
multiple racks.
Our API makes virtual machine provisioning trivial from an operational
standpoint, but it can take tens of seconds or even minutes after an API request
before a virtual machine is scheduled, booted, and running a workload. This
affects users in the form of cold-start latency, which a warm pool can help
mitigate, but not entirely eliminate. We accepted the latency in this blog post,
but our goal is to reduce it as much as possible.
Cleaning up virtual machines is also trivial using the Oxide API, but, depending
on the AI agent, there isn’t always a simple way to know when the workload
has finished. In this blog post, we handled this by treating a stopped virtual
machine as the signal that the workload has finished. This required us to wrap
our AI agent process in a systemd service that runs poweroff after the AI
agent process exits, along with a reaper process to poll the Oxide API for
virtual machines to clean up. Not the best user experience, especially if the AI
agent architecture doesn’t provide a delete operation, but we’re exploring how
to make this experience behave like a container that exits when the main process
is finished.
I’ve made it this far without talking about microVMs, but the microVM primitive
is exactly what some infrastructure providers are using to improve the user
experience of running AI agents without sacrificing the strong isolation they’ve
come to rely on from virtual machines.
The definition of microVMs can vary across the industry, but discussing them in
the context of AI agents has led to internal conversations about what a microVM
primitive could look like for Oxide. Our hypervisor team even spiked on work
to see how quickly we could launch a virtual machine and have it running its
workload. Turns out we’re not far off from popular microVM implementations!
We don’t have an

[truncated]
