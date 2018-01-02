define({ "api": [
  {
    "type": "POST",
    "url": "/api/survey/device/add",
    "title": "Add survey to device",
    "name": "Add_survey_to_device_api",
    "group": "Survey",
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "survey_id",
            "description": "<p>Survey ID No</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "device_id",
            "description": "<p>Device ID No</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n    device_id : 'device_1',\n    survey_id : 'survey_1'\n}",
          "type": "Object"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    data : null,\n    status : 'success'\n}",
          "type": "type"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Not found survey Response:",
          "content": "{\n    \"error\":\"not found\",\n    \"status\":\"error\"\n }",
          "type": "Object"
        }
      ]
    },
    "filename": "app/apidoc/survey.js",
    "groupTitle": "Survey"
  },
  {
    "type": "GET",
    "url": "/api/survey/list",
    "title": "Get All Surveys List Api",
    "name": "No",
    "group": "Survey",
    "version": "1.0.0",
    "success": {
      "fields": {
        "200": [
          {
            "group": "200",
            "type": "Array",
            "optional": false,
            "field": "surveys",
            "description": "<p>Surveys list</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n \"data\": [\n     {\n     \"id\": \"srv_xlOv\",\n     \"dtime\": 0,\n     \"mtime\": 0,\n     \"ctime\": 1514863061,\n     \"name\": \"survey 1\",\n     \"questions\": [\n         {\n         \"id\": \"\",\n         \"content\": \"cau hoi thu 1\",\n         \"summary\": \"summary 1\",\n         \"type\": \"single\",\n         \"point\": 100,\n         \"manded\": true,\n         \"answers\": [\n             {\n             \"icon\": \"smiley1\",\n             \"content\": \"answer content 1\",\n             \"point\": 30,\n             \"link\": \"\"\n             }\n         ]\n         },\n         {\n         \"id\": \"\",\n         \"content\": \"cau hoi thu 2\",\n         \"summary\": \"summary 2\",\n         \"type\": \"multiple\",\n         \"point\": 100,\n         \"manded\": true,\n         \"answers\": [\n             {\n             \"icon\": \"\",\n             \"content\": \"answer content 1\",\n             \"point\": 30,\n             \"link\": \"\"\n             }\n         ]\n         },\n         {\n         \"id\": \"\",\n         \"content\": \"cau hoi tra loi\",\n         \"summary\": \"summary 2\",\n         \"type\": \"answer\",\n         \"point\": 100,\n         \"manded\": true,\n         \"answers\": []\n         }\n     ]\n     }\n ],\n \"status\": \"success\"\n }",
          "type": "Object"
        }
      ]
    },
    "filename": "app/apidoc/survey.js",
    "groupTitle": "Survey"
  },
  {
    "type": "POST",
    "url": "/api/survey/add",
    "title": "Add Survey Api",
    "name": "apiName",
    "group": "Survey",
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>name of survey</p>"
          },
          {
            "group": "Parameter",
            "type": "Array",
            "optional": false,
            "field": "questions",
            "description": "<p>list of questions</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "content",
            "description": "<p>content of questions</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "type",
            "description": "<p>type of questions:one of {single, multiple, single}</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "point",
            "description": "<p>point of questions</p>"
          },
          {
            "group": "Parameter",
            "type": "Boolean",
            "optional": false,
            "field": "manded",
            "description": "<p>required question</p>"
          },
          {
            "group": "Parameter",
            "type": "Array",
            "optional": false,
            "field": "answers",
            "description": "<p>list of answer (only for single and multiple question)</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "icon",
            "description": "<p>Icon of Single Question (only for single)</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n    \"name\":\"survey 1\",\n    \"questions\":[\n        {\n            \"content\":\"cau hoi thu 1\",\n            \"summary\":\"summary 1\",\n            \"type\":\"single\",\n            \"point\":100,\n            \"manded\":true,\n            \"answers\":[\n                {\n                    \"icon\":\"smiley1\",\n                    \"content\":\"answer content 1\",\n                    \"point\":30\n                }\n            ]\n        },\n        {\n            \"content\":\"cau hoi thu 2\",\n            \"summary\":\"summary 2\",\n            \"type\":\"multiple\",\n            \"point\":100,\n            \"manded\":true,\n            \"answers\":[\n                {\n                    \"content\":\"answer content 1\",\n                    \"point\":30\n                }\n            ]\n        },\n        {\n            \"content\":\"cau hoi tra loi\",\n            \"summary\":\"summary 2\",\n            \"type\":\"answer\",\n            \"point\":100,\n            \"manded\":true\n        }\t\n    ]\n}",
          "type": "Object"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"data\": {\n        \"id\": \"srv_xlOv\",\n        \"dtime\": 0,\n        \"mtime\": 0,\n        \"ctime\": 1514863061,\n        \"name\": \"survey 1\",\n        \"questions\": [\n            {\n                \"id\": \"\",\n                \"content\": \"cau hoi thu 1\",\n                \"summary\": \"summary 1\",\n                \"type\": \"single\",\n                \"point\": 100,\n                \"manded\": true,\n                \"answers\": [\n                    {\n                        \"icon\": \"smiley1\",\n                        \"content\": \"answer content 1\",\n                        \"point\": 30,\n                        \"link\": \"\"\n                    }\n                ]\n            },\n            {\n                \"id\": \"\",\n                \"content\": \"cau hoi thu 2\",\n                \"summary\": \"summary 2\",\n                \"type\": \"multiple\",\n                \"point\": 100,\n                \"manded\": true,\n                \"answers\": [\n                    {\n                        \"icon\": \"\",\n                        \"content\": \"answer content 1\",\n                        \"point\": 30,\n                        \"link\": \"\"\n                    }\n                ]\n            },\n            {\n                \"id\": \"\",\n                \"content\": \"cau hoi tra loi\",\n                \"summary\": \"summary 2\",\n                \"type\": \"answer\",\n                \"point\": 100,\n                \"manded\": true,\n                \"answers\": null\n            }\n        ]\n    },\n    \"status\": \"success\"\n}",
          "type": "Object"
        }
      ]
    },
    "filename": "app/apidoc/survey.js",
    "groupTitle": "Survey"
  }
] });
