# TODO

## API
[] Authentication (Just standard email auth for now but eventually SSO as well)
[] Design system for storing messages in the instance that application goes down and in-memory cache is cleared
  [] Create job for cleaning up after 30 days (make this known to customers based on application use case)
[] Figure out realtime notifications and how we post to the client

## Client
[] Retrieve messages from API
[] Dashboard with stored data
[] Display notification when received from server
[] Cache everything we receive from the server (max 30 days) for data visualization