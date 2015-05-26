/**
 * Created by remy on 22/04/15.
 */

var chart;
var chartData;
var pieData;


function getBattery(callback) {
    var data;
    jQuery.ajax({
        url: "http:idp-api.herokuapp.com/spin/archive/batterij",
        type: "GET",
        dataType: 'json',
        success: function (resultData) {
            callback(resultData);
            //data = resultData;
            console.log(resultData);
        }, error: function (jqXHR, textStatus, errorThrown) {
            callback(null);
            console.log(errorThrown);
        }
    });
}


function getModes(callback) {
    var data;
    jQuery.ajax({
        url: "http:idp-api.herokuapp.com/spin/archive/mode",
        type: "GET",
        dataType: 'json',
        success: function (resultData) {
            callback(resultData);
            //data = resultData;
            //console.log(resultData);
        }, error: function (jqXHR, textStatus, errorThrown) {
            callback(null);
            console.log(errorThrown);
        }
    });
}


function generateData() {
    var data = getBattery(function (data) {
        chartData = new Array(data.length);
        for (var i = 0; i < data.length; i++) {
            chartData[i] = parseInt(data[i]["batterij"]);
        }
        console.log(chartData);
        createChart();
    });

}






function generatePieData() {
    var data = getModes(function (data) {
        pieData = [0,0,0]
        for (var i = 0; i < data.length; i++) {
            if(data[i]["mode"] == "dansen")
            {
                pieData[0]++;
                console.log(pieData[0]);
            }
            else if(data[i]["mode"] == "lopen")
            {
                pieData[1]++;
            }
            else
            {
                pieData[2]++;
            }
        }
        console.log(pieData);
        createPie(pieData);
    });
}



function createPie(datas)
{
    console.log(datas);
    var ctx = document.getElementById('myChart').getContext('2d');
    var data = [
        {
            value: datas[0],
            color:"#F7464A",
            highlight: "#FF5A5E",
            label: "dansen"
        },
        {
            value: datas[1],
            color: "#46BFBD",
            highlight: "#5AD3D1",
            label: "lopen"
        },
        {
            value: datas[2],
            color: "#FDB45C",
            highlight: "#FFC870",
            label: "autonoom"
        }
    ]
  
    var options = {
        //Boolean - Whether we should show a stroke on each segment
        segmentShowStroke : true,

        //String - The colour of each segment stroke
        segmentStrokeColor : "#fff",

        //Number - The width of each segment stroke
        segmentStrokeWidth : 2,

        //Number - The percentage of the chart that we cut out of the middle
        percentageInnerCutout : 0, // This is 0 for Pie charts

        //Number - Amount of animation steps
        animationSteps : 100,

        //String - Animation easing effect
        animationEasing : "easeOutBounce",

        //Boolean - Whether we animate the rotation of the Doughnut
        animateRotate : true,

        //Boolean - Whether we animate scaling the Doughnut from the centre
        animateScale : false

    }
    // For a pie chart
    var myPieChart = new Chart(ctx).Pie(data,options);
}

function createChart()
{
    var options = {scaleShowLabels : false};

    var buyerData = {
        labels : ["January","February","March","April","May","June"],
        datasets : [
            {
                fillColor : "rgba(172,194,132,0.4)",
                strokeColor : "#ACC26D",
                pointColor : "#fff",
                pointStrokeColor : "#9DB86D",
                data : chartData
            }
        ]
    };
    var buyers = document.getElementById('buyers').getContext('2d');
     chart = new Chart(buyers).Line(buyerData, options);
}

function createrad()
{
    var ctx = document.getElementById('ding').getContext('2d');
    var options = {
        //Boolean - Whether to show lines for each scale point
        scaleShowLine : true,

        //Boolean - Whether we show the angle lines out of the radar
        angleShowLineOut : true,

        //Boolean - Whether to show labels on the scale
        scaleShowLabels : false,

        // Boolean - Whether the scale should begin at zero
        scaleBeginAtZero : true,

        //String - Colour of the angle line
        angleLineColor : "rgba(0,0,0,.1)",

        //Number - Pixel width of the angle line
        angleLineWidth : 1,

        //String - Point label font declaration
        pointLabelFontFamily : "'Arial'",

        //String - Point label font weight
        pointLabelFontStyle : "normal",

        //Number - Point label font size in pixels
        pointLabelFontSize : 10,

        //String - Point label font colour
        pointLabelFontColor : "#666",

        //Boolean - Whether to show a dot for each point
        pointDot : true,

        //Number - Radius of each point dot in pixels
        pointDotRadius : 3,

        //Number - Pixel width of point dot stroke
        pointDotStrokeWidth : 1,

        //Number - amount extra to add to the radius to cater for hit detection outside the drawn point
        pointHitDetectionRadius : 20,

        //Boolean - Whether to show a stroke for datasets
        datasetStroke : true,

        //Number - Pixel width of dataset stroke
        datasetStrokeWidth : 2,

        //Boolean - Whether to fill the dataset with a colour
        datasetFill : true,

        //String - A legend template
        legendTemplate : "<ul class=\"<%=name.toLowerCase()%>-legend\"><% for (var i=0; i<datasets.length; i++){%><li><span style=\"background-color:<%=datasets[i].strokeColor%>\"></span><%if(datasets[i].label){%><%=datasets[i].label%><%}%></li><%}%></ul>"

    }
    var data = {
        labels: ["Eating", "Drinking", "Sleeping", "Designing", "Coding", "Cycling", "Running"],
        datasets: [
            {
                label: "My First dataset",
                fillColor: "rgba(220,220,220,0.2)",
                strokeColor: "rgba(220,220,220,1)",
                pointColor: "rgba(220,220,220,1)",
                pointStrokeColor: "#fff",
                pointHighlightFill: "#fff",
                pointHighlightStroke: "rgba(220,220,220,1)",
                data: [65, 59, 90, 81, 56, 55, 40]
            },
            {
                label: "My Second dataset",
                fillColor: "rgba(151,187,205,0.2)",
                strokeColor: "rgba(151,187,205,1)",
                pointColor: "rgba(151,187,205,1)",
                pointStrokeColor: "#fff",
                pointHighlightFill: "#fff",
                pointHighlightStroke: "rgba(151,187,205,1)",
                data: [28, 48, 40, 19, 96, 27, 100]
            }
        ]
    };
    var myRadarChart = new Chart(ctx).Radar(data, options);
}