/**
 * Created by remy on 23/04/15.
 */

var articleID = 0;

function getBlogPosts(callback)
{
    var data;
    jQuery.ajax({
        url: "http://idp-api.herokuapp.com/blog",
        type: "GET",
        dataType: 'json',
        success: function(resultData) {
            callback(resultData);
            //data = resultData;
            //console.log(resultData);
        },
        error : function(jqXHR, textStatus, errorThrown) {
            callback(null);
            console.log(errorThrown);
        }
    })

    //return data;
}

function generatePosts()
{
    var data = getBlogPosts(function(data){
        for(var i = 0; i < data.length; i++) {
            generatePost(data[i]);
            articleID++;
        }
    });

}


function generatePost(blogPost)
{
    var article = $("#articleContainer").append("<article id='article'"+ articleID + "></article>");
   // var article = $("#article"+articleID);
    article.append("<h3><a href='#'>"+blogPost['titel']+"</a></h3>");
    article.append("<h6>Written by <a href='#'>"+blogPost['auteur']+"</a> on August 12, 2012.</h6>");
    article.append("<p>"+blogPost['text']+"</p>");
}
