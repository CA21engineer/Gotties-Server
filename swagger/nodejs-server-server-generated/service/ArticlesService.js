'use strict';


/**
 * 記事一覧
 *
 * returns List
 **/
exports.articlesGET = function() {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = [ {
  "updated_at" : "2000-01-23T04:56:07.000+00:00",
  "before" : "",
  "user_id" : "user_id",
  "created_at" : "2000-01-23T04:56:07.000+00:00",
  "id" : "id",
  "after" : "",
  "title" : "title",
  "body" : "body",
  "category" : {
    "id" : "id",
    "title" : "title"
  }
}, {
  "updated_at" : "2000-01-23T04:56:07.000+00:00",
  "before" : "",
  "user_id" : "user_id",
  "created_at" : "2000-01-23T04:56:07.000+00:00",
  "id" : "id",
  "after" : "",
  "title" : "title",
  "body" : "body",
  "category" : {
    "id" : "id",
    "title" : "title"
  }
} ];
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}


/**
 * 記事詳細
 *
 * id Integer articles ID
 * returns Article
 **/
exports.articlesIdGET = function(id) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = {
  "updated_at" : "2000-01-23T04:56:07.000+00:00",
  "before" : "",
  "user_id" : "user_id",
  "created_at" : "2000-01-23T04:56:07.000+00:00",
  "id" : "id",
  "after" : "",
  "title" : "title",
  "body" : "body",
  "category" : {
    "id" : "id",
    "title" : "title"
  }
};
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}


/**
 * 新規記事作成
 *
 * body Body  (optional)
 * no response value expected for this operation
 **/
exports.articlesPOST = function(body) {
  return new Promise(function(resolve, reject) {
    resolve();
  });
}

