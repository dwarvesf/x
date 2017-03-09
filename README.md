# x
A mining storage for dwarvesf's extensions that being used in internal and external projects. 

# Services

Currently services provided:
- [x] Emailer 
    - Mandrill
    - Mailgun
    - Sendgrid
- [ ] Notifier
    - FCM 
    - GCM

# How to use

The mechanism of using this lib is very simple. It uses driver pattern (somehow like `database/sql` in Go source).

For example, to use Sendgrid as the mail provider:
- Setup the environment SENDGRID_API_ID and SENDGRID_API_KEY
- Register by blank import `import 	_ "github.com/dwarvesf/x/emailer/driver/sendgrid"`
- Open provider sendgrid
```
sendgrid, err := emailer.UseProvider("sendgrid")
if err != nil {
		
}
```
- Config the message:
```
var message = &emailer.Message{
		To:      "xxx@gmail.com",
		From:    "yyy@dwarvesf.com",
		Subject: "test",
		Text:    "test",
	}
```
- Send:  `sendgrid.Send(message)`

Every provider else works as the way like that.

# Contribution

* Fork it!
* Create your feature branch (for example FCM):

```
$ git checkout -b feature/fcm
```

* Write your function to implement the Notifier interface
* Commit your changes:

```
$ git commit -am "Add implementation FCM"
```

* Push to the branch:

```
$ git push origin feature/fcm
```

* Submit your pull request

# License

Copyright 2016 Dwarves Foundation

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
