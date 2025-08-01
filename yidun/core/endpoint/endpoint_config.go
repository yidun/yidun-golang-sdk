package endpoint

var (
	EndpointJSON = `
	{
		"antispam": {
		  "cn-hangzhou": [
			"localhost:8111"
		  ]
		},
		"crawler": {
		  "cn-hangzhou": [
			"as.dun.163.com",
			"as.dun.163yun.com"
		  ]
		},
		"file": {
		  "cn-hangzhou": [
			"as.dun.163.com",
			"as.dun.163yun.com"
		  ]
		},
		"media": {
		  "cn-hangzhou": [
			"as.dun.163.com",
			"as.dun.163yun.com"
		  ]
		},
		"mediaSubmit": {
		  "cn-hangzhou": [
			"as.dun.163.com",
			"as.dun.163yun.com"
		  ],
		  "sg-singapore": [
			"as-media-sg.dun.163yun.com",
			"as-media-sg.dun.163.com"
		  ]
		},
		"report": {
		  "cn-hangzhou": [
			"as.dun.163.com",
			"as.dun.163yun.com"
		  ]
		},
		"digitalBook": {
		  "cn-hangzhou": [
			"as.dun.163.com",
			"as.dun.163yun.com"
		  ]
		},
		"audioCheck": {
		  "cn-hangzhou": [
			"as.dun.163.com",
			"as.dun.163yun.com"
		  ],
		  "sg-singapore": [
			"as-audio-sg.dun.163yun.com",
			"as-audio-sg.dun.163.com"
		  ],
			"de-frankfurt": [
			  "as-audio-frankfurt.dun.163.com",
			  "as-audio-frankfurt.dun.163yun.com"
			],
		  "us-virginia": [
			  "as-audio-virginia.dun.163.com",
			  "as-audio-virginia.dun.163yun.com"
			]
		},
		"audioCommon": {
		  "cn-hangzhou": [
			"as.dun.163.com",
			"as.dun.163yun.com"
		  ],
			"sg-singapore": [
			  "as-audio-sg.dun.163.com",
			  "as-audio-sg.dun.163yun.com"
			],
			"de-frankfurt": [
			  "as-audio-frankfurt.dun.163.com",
			  "as-audio-frankfurt.dun.163yun.com"
			],
			"us-virginia": [
			  "as-audio-virginia.dun.163.com",
			  "as-audio-virginia.dun.163yun.com"
			]
		},
		"videoCheck": {
		  "cn-hangzhou": [
			"as.dun.163.com",
			"as.dun.163yun.com"
		  ],
		  "sg-singapore": [
			"as-video-sg.dun.163yun.com",
			"as-video-sg.dun.163.com"
		  ]
		},
		"videoCommon": {
		  "cn-hangzhou": [
			"as.dun.163.com",
			"as.dun.163yun.com"
		  ]
		},
		"videoSolutionCheck": {
		  "cn-hangzhou": [
			"as.dun.163.com",
			"as.dun.163yun.com"
		  ],
		  "sg-singapore": [
			"as-videosolution-sg.dun.163yun.com",
			"as-videosolution-sg.dun.163.com"
		  ],
		"de-frankfurt": [
		  "as-videosolution-frankfurt.dun.163.com",
		  "as-videosolution-frankfurt.dun.163yun.com"
		],
		 "us-virginia": [
			  "as-videosolution-virginia.dun.163.com",
			  "as-videosolution-virginia.dun.163yun.com"
			]
		},
		"videoSolutionCommon": {
		  "cn-hangzhou": [
			"as.dun.163.com",
			"as.dun.163yun.com"
		  ],
			"sg-singapore": [
			  "as-videosolution-sg.dun.163.com",
			  "as-videosolution-sg.dun.163yun.com"
			],
			"de-frankfurt": [
			  "as-videosolution-frankfurt.dun.163.com",
			  "as-videosolution-frankfurt.dun.163yun.com"
			],
			"us-virginia": [
			  "as-videosolution-virginia.dun.163.com",
			  "as-videosolution-virginia.dun.163yun.com"
			]
		},
		"liveAudio": {
		  "cn-hangzhou": [
			"as.dun.163.com",
			"as.dun.163yun.com"
		  ],
		"sg-singapore": [
			  "as-liveaudio-sg.dun.163.com",
			  "as-liveaudio-sg.dun.163yun.com"
			],
		 "us-virginia": [
			  "as-liveaudio-virginia.dun.163.com",
			  "as-liveaudio-virginia.dun.163yun.com"
			]
		},
		"liveVideo": {
		  "cn-hangzhou": [
			"as.dun.163.com",
			"as.dun.163yun.com"
		  ]
		},
		"liveVideoSolutionCheck": {
		  "cn-hangzhou": [
			"as.dun.163.com",
			"as.dun.163yun.com"
		  ],
		  "sg-singapore": [
			"as-livesolution-sg.dun.163yun.com",
			"as-livesolution-sg.dun.163.com"
		  ],
			"us-virginia": [
				  "as-livesolution-virginia.dun.163.com",
				  "as-livesolution-virginia.dun.163yun.com"
				]
		},
		"liveVideoSolutionCommon": {
		  "cn-hangzhou": [
			"as.dun.163.com",
			"as.dun.163yun.com"
		  ],
			"sg-singapore": [
			  "as-livesolution-sg.dun.163.com",
			  "as-livesolution-sg.dun.163yun.com"
			],
			"us-virginia": [
			  "as-livesolution-virginia.dun.163.com",
			  "as-livesolution-virginia.dun.163yun.com"
			]
		},
		"keyword": {
		  "cn-hangzhou": [
			"as.dun.163.com",
			"as.dun.163yun.com"
		  ]
		},
		"anticheat": {
		  "cn-hangzhou": [
			"ac.dun.163.com",
			"ac.dun.163yun.com"
		  ]
		},
		"sms": {
		  "cn-hangzhou": [
			"sms.dun.163.com",
			"sms.dun.163yun.com"
		  ]
		},
		"captcha": {
		  "cn-hangzhou": [
			"c.dun.163.com",
			"c.dun.163yun.com"
		  ]
		},
		"mobileverify": {
		  "cn-hangzhou": [
			"ye.dun.163.com",
			"ye.dun.163yun.com"
		  ]
		},
		"auth": {
		  "cn-hangzhou": [
			"verify.dun.163.com",
			"verify.dun.163yun.com"
		  ]
		},
		"imageCheck": {
		  "cn-hangzhou": [
			"as.dun.163.com",
			"as.dun.163yun.com"
		  ],
		  "cn-beijing": [
            "as-image-bj.dun.163.com",
            "as-image-bj.dun.163yun.com"
          ],
		  "sg-singapore": [
			"as-image-sg.dun.163yun.com",
			"as-image-sg.dun.163.com"
		  ],
          "us-virginia": [
            "as-image-virginia.dun.163.com",
            "as-image-virginia.dun.163yun.com"
          ],
          "us-silicon-valley": [
             "as-image-sv.dun.163.com",
             "as-image-sv.dun.163yun.com"
          ]
		},
		"imageCommon": {
		  "cn-hangzhou": [
			"as.dun.163.com",
			"as.dun.163yun.com"
		  ],
		  "cn-beijing": [
             "as-image-bj.dun.163.com",
             "as-image-bj.dun.163yun.com"
          ],
          "sg-singapore": [
            "as-image-sg.dun.163.com",
            "as-image-sg.dun.163yun.com"
          ],
          "us-virginia": [
            "as-image-virginia.dun.163.com",
            "as-image-virginia.dun.163yun.com"
          ],
          "us-silicon-valley": [
             "as-image-sv.dun.163.com",
             "as-image-sv.dun.163yun.com"
          ]
		},
		"text-check": {
		  "cn-hangzhou": [
			"as.dun.163.com",
			"as.dun.163yun.com"
		  ],
		  "cn-beijing": [
			"as-text-bj.dun.163.com",
			"as-text-bj.dun.163yun.com"
		  ],
		  "cn-shanghai": [
			"as-text-sh.dun.163yun.com"
		  ],
		  "cn-guangzhou": [
			"as-text-gz.dun.163yun.com"
		  ],
		  "sg-singapore": [
			"as-text-sg.dun.163yun.com",
			"as-text-sg.dun.163.com"
		  ],
		  "de-frankfurt": [
			"as-text-frankfurt.dun.163.com"
		  ],
		  "us-virginia": [
			"as-text-virginia.dun.163.com"
		  ],
		  "jp-tokyo": [
			"as-text-tokyo.dun.163.com",
			"as-text-tokyo.dun.163yun.com"
		  ]
		},
		"text-api": {
		  "cn-hangzhou": [
			"as.dun.163.com",
			"as.dun.163yun.com"
		  ]
		},
		"aigc-stream-api": {
		  "cn-hangzhou": [
			"as.dun.163.com",
			"as.dun.163yun.com"
		  ]
		},
		"aigc-stream-check-api": {
		  "cn-hangzhou": [
			"as.dun.163.com",
			"as.dun.163yun.com"
		  ]
		},
		"openapi": {
			"cn-hangzhou": [
			  "openapi.dun.163.com"
			]
		},
		"pretreatment": {
		  "cn-hangzhou": [
			"as.dun.163.com",
			"as.dun.163yun.com"
		  ]
		},
		"list": {
		  "cn-hangzhou": [
			"as.dun.163.com",
			"as.dun.163yun.com"
		  ]
		},
		"fingerprint": {
		  "cn-hangzhou": [
			"fp-query.dun.163.com",
			"fp-query.dun.163yun.com"
		  ]
		},
		"irisk": {
		  "cn-hangzhou": [
			"ir-open.dun.163.com"
		  ],
		  "sg-singapore": [
			"ir-open.guardease.com"
		  ],
		  "de-frankfurt": [
			"ir-open.guardease.com"
		  ],
		  "us-virginia": [
			"ir-open.guardease.com"
		  ],
		  "jp-tokyo": [
			"ir-open.guardease.com"
		  ]
		},
		"profile": {
		  "cn-hangzhou": [
			"rp.dun.163.com"
		  ]
		},
        "aigcStream": {
          "cn-hangzhou": [
            "as.dun.163.com"
          ]
        },
        "aigcStreamPush": {
          "cn-hangzhou": [
            "as.dun.163.com"
          ]
        }
	  }
	  `
)
