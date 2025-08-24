# Shared Response Types

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared#Attachment">Attachment</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared#BaseResponse">BaseResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared#Reaction">Reaction</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared#User">User</a>

# Accounts

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#Account">Account</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#AccountsResponse">AccountsResponse</a>

Methods:

- <code title="get /v0/get-accounts">client.Accounts.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#AccountService.GetAccounts">GetAccounts</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#AccountsResponse">AccountsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# App

Methods:

- <code title="post /v0/focus-app">client.App.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#AppService.FocusApp">FocusApp</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#AppFocusAppParams">AppFocusAppParams</a>) (<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared#BaseResponse">BaseResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Messages

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#Message">Message</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#SearchResponse">SearchResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#SendResponse">SendResponse</a>

Methods:

- <code title="post /v0/draft-message">client.Messages.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#MessageService.DraftMessage">DraftMessage</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#MessageDraftMessageParams">MessageDraftMessageParams</a>) (<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared#BaseResponse">BaseResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v0/search-messages">client.Messages.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#MessageService.SearchMessages">SearchMessages</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#MessageSearchMessagesParams">MessageSearchMessagesParams</a>) (<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/packages/pagination#CursorID">CursorID</a>[<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#Message">Message</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v0/send-message">client.Messages.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#MessageService.SendMessage">SendMessage</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#MessageSendMessageParams">MessageSendMessageParams</a>) (<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#SendResponse">SendResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Chats

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#Chat">Chat</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#FindChatsResponse">FindChatsResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#GetChatResponse">GetChatResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#LinkResponse">LinkResponse</a>

Methods:

- <code title="post /v0/archive-chat">client.Chats.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ChatService.ArchiveChat">ArchiveChat</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ChatArchiveChatParams">ChatArchiveChatParams</a>) (<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared#BaseResponse">BaseResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v0/find-chats">client.Chats.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ChatService.FindChats">FindChats</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ChatFindChatsParams">ChatFindChatsParams</a>) (<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/packages/pagination#CursorID">CursorID</a>[<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#Chat">Chat</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v0/get-chat">client.Chats.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ChatService.GetChat">GetChat</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ChatGetChatParams">ChatGetChatParams</a>) (<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#GetChatResponse">GetChatResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v0/get-link-to-chat">client.Chats.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ChatService.GetLinkToChat">GetLinkToChat</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ChatGetLinkToChatParams">ChatGetLinkToChatParams</a>) (<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#LinkResponse">LinkResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Reminders

Methods:

- <code title="post /v0/clear-chat-reminder">client.Reminders.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ReminderService.ClearChatReminder">ClearChatReminder</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ReminderClearChatReminderParams">ReminderClearChatReminderParams</a>) (<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared#BaseResponse">BaseResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v0/set-chat-reminder">client.Reminders.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ReminderService.SetChatReminder">SetChatReminder</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ReminderSetChatReminderParams">ReminderSetChatReminderParams</a>) (<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared#BaseResponse">BaseResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# OAuth
