<h1 align="center"> Meetings API </h1> <br>
<p align="center">
  <a href="https://gitpoint.co/">
    <img alt="MeetApi" title="MeetApi" src="/Assets/logo.svg" width=250>
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
- [Running Locally](#running-locally)
- [RESTful URLs](#restful-urls)
- [Requests and Responses](#requests-and-responses)
- [Feedback](#feedback)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->
<br />



## Introduction
This is a basic version of API managing and scheduling meetings. This is my task done for appointy internship using golang and mongodb.



## Features

A few of the things you can do with MeetAPI:

* Schedules Meetings
* Lists all meetings within a time frame
* Create non overlapping meetings
* List all meetings of a participant
* Thread Safe


<br />


## Running Locally

Make sure you have [Go](https://golang.org/) and the [Mongodb](https://www.mongodb.com/) installed.

```sh
git clone https://github.com/SreemanthG/Meetings-API-Appointy.git
cd meetings-api-appointy
go get
go run main.go
```
<br />



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
<br />



## Requests and Responses

### API Resources

  - [POST /meetings](#post-meetings)
  - [GET /meetings/[id]](#get-meetings)
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
 <br />



    
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
<br />




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
<br />



    
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
<br />



        
## Database and Structure


<p align="center">
  <img src = "/Assets/database.JPG" width=700>
</p>


## Feedback

Feel free to send feedback on [Twitter](https://twitter.com/GSreemanth) or [file an issue](https://github.com/SreemanthG/meetings-api-appointy/issues/new). Feature requests are always welcome. You can contact me at sreemanth2001@gmail.com
