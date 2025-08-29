# Shared Response Types

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared#Account">Account</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared#Attachment">Attachment</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared#BaseResponse">BaseResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared#Chat">Chat</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared#Message">Message</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared#Reaction">Reaction</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared#User">User</a>

# Accounts

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#AccountListResponse">AccountListResponse</a>

Methods:

- <code title="get /v0/get-accounts">client.Accounts.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#AccountService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#AccountListResponse">AccountListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# App

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#AppFocusResponse">AppFocusResponse</a>

Methods:

- <code title="post /v0/open-app">client.App.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#AppService.Focus">Focus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#AppFocusParams">AppFocusParams</a>) (<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#AppFocusResponse">AppFocusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Chats

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ChatGetResponse">ChatGetResponse</a>

Methods:

- <code title="post /v0/archive-chat">client.Chats.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ChatService.Archive">Archive</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ChatArchiveParams">ChatArchiveParams</a>) (<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared#BaseResponse">BaseResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v0/get-chat">client.Chats.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ChatService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ChatGetParams">ChatGetParams</a>) (<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ChatGetResponse">ChatGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v0/search-chats">client.Chats.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ChatService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ChatSearchParams">ChatSearchParams</a>) (<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/packages/pagination#CursorID">CursorID</a>[<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared#Chat">Chat</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Messages

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#MessageGetAttachmentResponse">MessageGetAttachmentResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#MessageSendResponse">MessageSendResponse</a>

Methods:

- <code title="post /v0/get-attachment">client.Messages.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#MessageService.GetAttachment">GetAttachment</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#MessageGetAttachmentParams">MessageGetAttachmentParams</a>) (<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#MessageGetAttachmentResponse">MessageGetAttachmentResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v0/search-messages">client.Messages.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#MessageService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#MessageSearchParams">MessageSearchParams</a>) (<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/packages/pagination#CursorID">CursorID</a>[<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared#Message">Message</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v0/send-message">client.Messages.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#MessageService.Send">Send</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#MessageSendParams">MessageSendParams</a>) (<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#MessageSendResponse">MessageSendResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Reminders

Methods:

- <code title="post /v0/clear-chat-reminder">client.Reminders.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ReminderService.Clear">Clear</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ReminderClearParams">ReminderClearParams</a>) (<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared#BaseResponse">BaseResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v0/set-chat-reminder">client.Reminders.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ReminderService.Set">Set</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ReminderSetParams">ReminderSetParams</a>) (<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared#BaseResponse">BaseResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# OAuth

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#UserInfo">UserInfo</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#OAuthAuthorizeCallbackResponse">OAuthAuthorizeCallbackResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#OAuthRegisterClientResponse">OAuthRegisterClientResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#OAuthTokenResponse">OAuthTokenResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#OAuthWellKnownAuthorizationServerResponse">OAuthWellKnownAuthorizationServerResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#OAuthWellKnownProtectedResourceResponse">OAuthWellKnownProtectedResourceResponse</a>

Methods:

- <code title="get /oauth/authorize">client.OAuth.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#OAuthService.Authorize">Authorize</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#OAuthAuthorizeParams">OAuthAuthorizeParams</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /oauth/authorize/callback">client.OAuth.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#OAuthService.AuthorizeCallback">AuthorizeCallback</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#OAuthAuthorizeCallbackParams">OAuthAuthorizeCallbackParams</a>) (<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#OAuthAuthorizeCallbackResponse">OAuthAuthorizeCallbackResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /oauth/userinfo">client.OAuth.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#OAuthService.GetUserInfo">GetUserInfo</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#UserInfo">UserInfo</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /oauth/register">client.OAuth.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#OAuthService.RegisterClient">RegisterClient</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#OAuthRegisterClientParams">OAuthRegisterClientParams</a>) (<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#OAuthRegisterClientResponse">OAuthRegisterClientResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /oauth/revoke">client.OAuth.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#OAuthService.RevokeToken">RevokeToken</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#OAuthRevokeTokenParams">OAuthRevokeTokenParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /oauth/token">client.OAuth.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#OAuthService.Token">Token</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#OAuthTokenParams">OAuthTokenParams</a>) (<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#OAuthTokenResponse">OAuthTokenResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /.well-known/oauth-authorization-server">client.OAuth.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#OAuthService.WellKnownAuthorizationServer">WellKnownAuthorizationServer</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#OAuthWellKnownAuthorizationServerResponse">OAuthWellKnownAuthorizationServerResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /.well-known/oauth-protected-resource">client.OAuth.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#OAuthService.WellKnownProtectedResource">WellKnownProtectedResource</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">githubcombeeperdesktopapigo</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#OAuthWellKnownProtectedResourceResponse">OAuthWellKnownProtectedResourceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
