/**
 * 
 * @api {GET} /api/survey/list Get All Surveys List Api
 * @apiName No
 * @apiGroup Surveys
 * @apiVersion  1.0.0
 * 
 * 
 * @apiSuccess (200) {Array} surveys Surveys list
 * 
 * 
 * @apiSuccessExample {Object} Success-Response:
   {
       data : [],
       status:'success'
   }
 * 
 * 
 */

 /**
  * 
  * @api {POST} /api/survey/add Add Survey Api
  * @apiName apiName
  * @apiGroup Surveys
  * @apiVersion  1.0.0
  * 
  * 
  * @apiParam  {String} name name of survey
  * 
  * @apiParam  {Array} questions list of questions
  *
  * @apiParam  {String} content content of questions
  *
  * @apiParam  {String} type type of questions:one of {single, multiple, single}
  *
  * @apiParam  {Number} point point of questions
  *
  * @apiParam  {Boolean} manded required question
  *
  * @apiParam  {Array} answers list of answer (only for single and multiple question)
  *
  * @apiParam  {String} icon Icon of Single Question (only for single)
  *
  * @apiParam  {Content} content content of Single Question (only for single)
  *
  * @apiParam  {point} point point of answer
  *
  * @apiParamExample  {Object} Request-Example:
    {
        "name":"survey 1",
        "questions":[
            {
                "content":"cau hoi thu 1",
                "summary":"summary 1",
                "type":"single",
                "point":100,
                "manded":true,
                "answers":[
                    {
                        "icon":"smiley1",
                        "content":"answer content 1",
                        "point":30
                    }
                ]
            },
            {
                "content":"cau hoi thu 2",
                "summary":"summary 2",
                "type":"multiple",
                "point":100,
                "manded":true,
                "answers":[
                    {
                        "content":"answer content 1",
                        "point":30
                    }
                ]
            },
            {
                "content":"cau hoi tra loi",
                "summary":"summary 2",
                "type":"answer",
                "point":100,
                "manded":true
            }	
        ]
    }
  * 
  * 
  * @apiSuccessExample {Object} Success-Response:
    {
        "data": {
            "id": "srv_xlOv",
            "dtime": 0,
            "mtime": 0,
            "ctime": 1514863061,
            "name": "survey 1",
            "questions": [
                {
                    "id": "",
                    "content": "cau hoi thu 1",
                    "summary": "summary 1",
                    "type": "single",
                    "point": 100,
                    "manded": true,
                    "answers": [
                        {
                            "icon": "smiley1",
                            "content": "answer content 1",
                            "point": 30,
                            "link": ""
                        }
                    ]
                },
                {
                    "id": "",
                    "content": "cau hoi thu 2",
                    "summary": "summary 2",
                    "type": "multiple",
                    "point": 100,
                    "manded": true,
                    "answers": [
                        {
                            "icon": "",
                            "content": "answer content 1",
                            "point": 30,
                            "link": ""
                        }
                    ]
                },
                {
                    "id": "",
                    "content": "cau hoi tra loi",
                    "summary": "summary 2",
                    "type": "answer",
                    "point": 100,
                    "manded": true,
                    "answers": null
                }
            ]
        },
        "status": "success"
    }
  * 
  *@@apiErrorExample {type} Error-Response:
    {
        "error":"Answer at least 1",
        "status":"error"
    }
  * 
  */