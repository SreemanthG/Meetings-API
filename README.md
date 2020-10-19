<h1 align="center"> Meetings API </h1> <br>
<p align="center">
  <a href="https://gitpoint.co/">
    <img alt="GitPoint" title="GitPoint" src="/Assets/logo.svg" width=250>
  </a>
</p>
<br />



<p align="center">
  Schedule Meetings with API. Built with Golang
</p>

<p align="center">
  <a href="https://golang.org/" >
    <img alt="Golang" title="Golang" src="https://golang.org/lib/godoc/images/go-logo-blue.svg" height="50px">
  </a>
  &nbsp;&nbsp;&nbsp;
  <a href="https://www.mongodb.com/">
    <img alt="MongoDB" title="MongoDB" src="https://webassets.mongodb.com/_com_assets/cms/MongoDB_Logo_FullColorBlack_RGB-4td3yuxzjs.png" height="50px">
  </a>
</p>
<br />


<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Feedback](#feedback)
- [Contributors](#contributors)
- [Build Process](#build-process)
- [Backers](#backers-)
- [Sponsors](#sponsors-)
- [Acknowledgments](#acknowledgments)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Introduction
This is a basic version of managing API. This is my task done for appointy internship using golang and mongodb.


<p align="center">
  <img src = "http://i.imgur.com/HowF6aM.png" width=350>
</p>

## Features

A few of the things you can do with GitPoint:

* Schedules Meetings
* Lists all meetings within a time frame
* Create non overlapping meetings
* List all meetings of a participant
* Thread Safe

<p align="center">
  <img src = "http://i.imgur.com/IkSnFRL.png" width=700>
</p>
<br />


## Running Locally

Make sure you have [Go](https://golang.org/) and the [Mongodb](https://www.mongodb.com/) installed.

```sh
git clone https://github.com/SreemanthG/Meatings-API-Appointy.git
cd meetings-api-appointy
go get
go run main.go
```

## RESTful URLs

### Good URL examples
* Schedule a meeting:
    * post http://localhost/meetings
* Get a meeting using id:
    * GET http://localhost/meeting/5f8c88ed089352ae924e7c84
* List all meetings within a time frame:
    * GET http://localhost/meetings?start=1602255600&end=1602255600
* List all meetings of a participant:
    * GET http://localhost/meetings?participant=example@gmail.com

### Bad URL examples
* Schedule a meeting:
    * post http://localhost/meetings/213112
* Get a meeting using id:
    * GET http://localhost/meeting/5f8c88ed089352ae924e7c84/5f8c88ed089352ae924e7c84
* List all meetings within a time frame:
    * GET http://localhost/meetings?start=12-12-2020&end=12-12-2020
* List all meetings of a participant:
    * GET http://localhost/meetings?participant=examplename

## Request & Response Examples

### API Resources

  - [POST /meetings](#post-meetings)
  - [GET /meetings/[id]](#get-meetingsid)
  - [GET /meetings?start=start=[start time here]&end=[end time here]](#post-magazinesidarticles)
  - [GET /meetings?participant=[email]](#post-magazinesidarticles)
  
### POST /meetings

Example: http://localhost/meetings

Request body(JSON):
    
      {
      
        "title": "My new Meeting",
        "participants": [
            {"name":"sreemanth1","email":"sreemanth1@gmail.com","rsvp":"yes"},
            {"name":"sreemanth2","email":"sreemanth2@gmail.com","rsvp":"yes"}
        ],
        "start_Time": 1603188000,
        "end_Time": 1603202400
      
      }


Response body:
<br />
      

Success:

      {
      "InsertedID":"5f8d36b26199472986f1a690"
      }

Failure:

      One of the participants with email sreemanth1@gmail.com timings are clashing
     
### GET /meeting/[id]

Example: http://localhost/meeting/5f8d36b26199472986f1a690

Response body:
    
      {
        "_id": "5f8d36b26199472986f1a690",
        "title": "My new Meeting",
        "participants": [
            {
                "name": "sreemanth1",
                "email": "sreemanth1@gmail.com",
                "rsvp": "yes"
            },
            {
                "name": "sreemanth2",
                "email": "sreemanth2@gmail.com",
                "rsvp": "yes"
            }
        ],
        "start_Time": 1603188000,
        "end_Time": 1603202400,
        "creation_Timestamp": "2020-10-19T06:48:18.325Z"
      }

### GET /meetings?start=[start time here]&end=[end time here]

Example: http://localhost/meetings?start=1602255600&end=1603202400

Response body:
    
        [
          {
              "_id": "5f8c81b85fe813d139bcb1bf",
              "title": "My new test Meeting",
              "participants": [
                  {
                      "name": "sreemanth",
                      "email": "sreemanth@gmail.com",
                      "rsvp": "yes"
                  },
                  {
                      "name": "daksh",
                      "email": "daksh@gmail.com",
                      "rsvp": "no"
                  },
                  {
                      "name": "pranav",
                      "email": "pranav@gmail.com",
                      "rsvp": "maybe"
                  }
              ],
              "start_Time": 1602255600,
              "end_Time": 1602273600,
              "creation_Timestamp": "2020-10-18T17:56:08.676Z"
          },
          {
              "_id": "5f8c88ed089352ae924e7c84",
              "title": "My new test2 Meeting",
              "participants": [
                  {
                      "name": "daksh",
                      "email": "daksh@gmail.com",
                      "rsvp": "no"
                  },
                  {
                      "name": "pranav",
                      "email": "pranav@gmail.com",
                      "rsvp": "maybe"
                  }
              ],
              "start_Time": 1602255600,
              "end_Time": 1602273600,
              "creation_Timestamp": "2020-10-18T18:26:53.235Z"
          },
          {
              "_id": "5f8ca61c67bd98c1c301d19b",
              "title": "My new test2 Meeting",
              "participants": [
                  {
                      "name": "daksh",
                      "email": "daksh@gmail.com",
                      "rsvp": "yes"
                  },
                  {
                      "name": "pranav",
                      "email": "pranav@gmail.com",
                      "rsvp": "yes"
                  }
              ],
              "start_Time": 1602255600,
              "end_Time": 1602273600,
              "creation_Timestamp": "2020-10-18T20:31:23.395Z"
          },
          {
              "_id": "5f8d36b26199472986f1a690",
              "title": "My new Meeting",
              "participants": [
                  {
                      "name": "sreemanth1",
                      "email": "sreemanth1@gmail.com",
                      "rsvp": "yes"
                  },
                  {
                      "name": "sreemanth2",
                      "email": "sreemanth2@gmail.com",
                      "rsvp": "yes"
                  }
              ],
              "start_Time": 1603188000,
              "end_Time": 1603202400,
              "creation_Timestamp": "2020-10-19T06:48:18.325Z"
          }
        ]
     
### GET /meetings?participant=[email]
Example: http://localhost/meetings?participant=[email]

Response body:

        [
          {
              "_id": "5f8c81b85fe813d139bcb1bf",
              "title": "My new test Meeting",
              "participants": [
                  {
                      "name": "sreemanth",
                      "email": "sreemanth@gmail.com",
                      "rsvp": "yes"
                  },
                  {
                      "name": "daksh",
                      "email": "daksh@gmail.com",
                      "rsvp": "no"
                  },
                  {
                      "name": "pranav",
                      "email": "pranav@gmail.com",
                      "rsvp": "maybe"
                  }
              ],
              "start_Time": 1602255600,
              "end_Time": 1602273600,
              "creation_Timestamp": "2020-10-18T17:56:08.676Z"
          }
        ]    
        
     
## Feedback

Feel free to send us feedback on [Twitter](https://twitter.com/gitpointapp) or [file an issue](https://github.com/gitpoint/git-point/issues/new). Feature requests are always welcome. If you wish to contribute, please take a quick look at the [guidelines](./CONTRIBUTING.md)!

If there's anything you'd like to chat about, please feel free to join our [Gitter chat](https://gitter.im/git-point)!

## Contributors

This project follows the [all-contributors](https://github.com/kentcdodds/all-contributors) specification and is brought to you by these [awesome contributors](./CONTRIBUTORS.md).

## Build Process

- Follow the [React Native Guide](https://facebook.github.io/react-native/docs/getting-started.html) for getting started building a project with native code. **A Mac is required if you wish to develop for iOS.**
- Clone or download the repo
- `yarn` to install dependencies
- `yarn run link` to link react-native dependencies
- `yarn start:ios` to start the packager and run the app in the iOS simulator (`yarn start:ios:logger` will boot the application with [redux-logger](<https://github.com/evgenyrodionov/redux-logger>))
- `yarn start:android` to start the packager and run the app in the the Android device/emulator (`yarn start:android:logger` will boot the application with [redux-logger](https://github.com/evgenyrodionov/redux-logger))

Please take a look at the [contributing guidelines](./CONTRIBUTING.md) for a detailed process on how to build your application as well as troubleshooting information.

**Development Keys**: The `CLIENT_ID` and `CLIENT_SECRET` in `api/index.js` are for development purposes and do not represent the actual application keys. Feel free to use them or use a new set of keys by creating an [OAuth application](https://github.com/settings/applications/new) of your own. Set the "Authorization callback URL" to `gitpoint://welcome`.

## Backers [![Backers on Open Collective](https://opencollective.com/git-point/backers/badge.svg)](#backers)

Thank you to all our backers! 🙏 [[Become a backer](https://opencollective.com/git-point#backer)]

<a href="https://opencollective.com/git-point#backers" target="_blank"><img src="https://opencollective.com/git-point/backers.svg?width=890"></a>

## Sponsors [![Sponsors on Open Collective](https://opencollective.com/git-point/sponsors/badge.svg)](#sponsors)

Support this project by becoming a sponsor. Your logo will show up here with a link to your website. [[Become a sponsor](https://opencollective.com/git-point#sponsor)]

<a href="https://opencollective.com/git-point/sponsor/0/website" target="_blank"><img src="https://opencollective.com/git-point/sponsor/0/avatar.svg"></a>
<a href="https://opencollective.com/git-point/sponsor/1/website" target="_blank"><img src="https://opencollective.com/git-point/sponsor/1/avatar.svg"></a>
<a href="https://opencollective.com/git-point/sponsor/2/website" target="_blank"><img src="https://opencollective.com/git-point/sponsor/2/avatar.svg"></a>
<a href="https://opencollective.com/git-point/sponsor/3/website" target="_blank"><img src="https://opencollective.com/git-point/sponsor/3/avatar.svg"></a>
<a href="https://opencollective.com/git-point/sponsor/4/website" target="_blank"><img src="https://opencollective.com/git-point/sponsor/4/avatar.svg"></a>
<a href="https://opencollective.com/git-point/sponsor/5/website" target="_blank"><img src="https://opencollective.com/git-point/sponsor/5/avatar.svg"></a>
<a href="https://opencollective.com/git-point/sponsor/6/website" target="_blank"><img src="https://opencollective.com/git-point/sponsor/6/avatar.svg"></a>
<a href="https://opencollective.com/git-point/sponsor/7/website" target="_blank"><img src="https://opencollective.com/git-point/sponsor/7/avatar.svg"></a>
<a href="https://opencollective.com/git-point/sponsor/8/website" target="_blank"><img src="https://opencollective.com/git-point/sponsor/8/avatar.svg"></a>
<a href="https://opencollective.com/git-point/sponsor/9/website" target="_blank"><img src="https://opencollective.com/git-point/sponsor/9/avatar.svg"></a>

## Acknowledgments

Thanks to [JetBrains](https://www.jetbrains.com) for supporting us with a [free Open Source License](https://www.jetbrains.com/buy/opensource).
