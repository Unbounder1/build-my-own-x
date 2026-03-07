Design a Slack-like chat system which supports sending messages to a group or to an individual user. Should support use cases to create a new group chat, add users to a group chat, send notifications to a group chat or an individual user. Add support for rich media (image etc) in a chat message. As a follow up, how to implement deletion of a message by user.
Design a webhook service which allows users to register their callback address and an eventId. Whenever eventId is triggered, the system should call the registered callback address with a specific payload. It is safe to assume that each eventId → callback address is unique and one eventId will trigger only one callback address. Assume a highly scalable webhook delivery system which can handle up to 1B events per day.
Design a multi-tenant CI/CD workflow system. Users trigger a workflow creation request via a git push. The workflow configuration is stored in a GitHub repository, and the repository ID can be obtained through an internal API service that analyzes the git push. Once the workflow is created, the job is scheduled (note that each workflow includes a sequence of jobs defined in the configuration file). Additionally, there is a front-end UI to display the progress of the workflow after it has been created.

```
Design a basic AI chat application like ChatGPT.

Requirements:

Users should be able to authenticate themselves.
Users should be able to send messages and receive back a response from an AI.
The AI response should stream in.
If the user refreshes the page, they get a fresh start. No persistence of the conversation is needed.
Consumer App (1 Million users)
Scoped to single browser ssession?
Time to first token - a few seconds
SEO requirments - basic
Text only
Some sort of moderation
Available moderation service
Auth Model - Single Tier
Chat history may require some soft storage
```