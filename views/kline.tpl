<!DOCTYPE html>
<html>

<head>
    <meta charset="utf-8">
    <title>K线数据</title>
    <script src="https://go-echarts.github.io/go-echarts-assets/assets/echarts.min.js"></script>
    <link href="https://go-echarts.github.io/go-echarts-assets/assets/bulma.min.css" rel="stylesheet">
</head>

<body>
    <div class="select" style="margin-right:10px; margin-top:10px; position:fixed; right:10px;">
        <select onchange="window.location.href=this.options[this.selectedIndex].value"> -- + '/' + this.options[this.selectedIndex].text">
            <option value="Kline">Kline-K 线图</option>
            <<<range $index, $item := .Instruments>>>
                <option value="/kline/<<<$item>>>"> <<<$index>>> <<<$item>>> </option>
            <<<end>>>
        </select></div>


    <div class="container">
        <div class="item" id="sNHcHHgUchYa" style="width:900px;height:500px;"></div>
    </div>
    <script type="text/javascript">
        "use strict";
        var myChart_sNHcHHgUchYa = echarts.init(document.getElementById('sNHcHHgUchYa'), "white");
        var option_sNHcHHgUchYa = {
            title: { "text": <<<.title>>>, },
            tooltip: {},
            legend: {},
            dataZoom: [{ "type": "inside", "start": 50, "end": 100, "xAxisIndex": [0] }, { "type": "slider", "start": 50, "end": 100, "xAxisIndex": [0] }],
            xAxis: [{ "data": <<<.Dates>>>, "splitNumber": 20, "splitArea": { "show": false, }, "splitLine": { "show": false, } }],
            yAxis: [{ "axisLabel": { "show": true }, "scale": true, "splitArea": { "show": false, }, "splitLine": { "show": false, } }],
            series: [
                { "name": "kline", "type": "candlestick", "data": <<<.Bars>>>, "itemStyle": { "color": "#ec0000", "color0": "#00da3c", "borderColor": "#8A0000", "borderColor0": "#008F28" }, "label": { "show": false }, "emphasis": { "label": { "show": false }, }, "markLine": { "label": { "show": false } }, "markPoint": { "data": [{ "name": "highest value", "type": "max", "valueDim": "highest" }, { "name": "lowest value", "type": "min", "valueDim": "lowest" }], "label": { "show": true } }, },
            ],
            color: ["#c23531", "#2f4554", "#61a0a8", "#d48265", "#91c7ae", "#749f83", "#ca8622", "#bda29a", "#6e7074", "#546570", "#c4ccd3"],
        };
        myChart_sNHcHHgUchYa.setOption(option_sNHcHHgUchYa);
    </script>

    <br />
    <style>
        .container {
            display: flex;
            justify-content: center;
            align-items: center;
        }

        .item {
            margin: auto;
        }
    </style>
</body>

</html>