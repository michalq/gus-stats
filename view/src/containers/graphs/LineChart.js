import React, { Component } from "react";
import { LineChart } from "../../lib/LineChart";

export default class LineChartComponent extends Component {
  render() {
    var data = [
      {
        date: '2010',
        value: 1,
        variableName: "test"
      },
      {
        date: '2011',
        value: 2,
        variableName: "test"
      },
      {
        date: '2012',
        value: 3,
        variableName: "test"
      },
      {
        date: '2010',
        value: 3,
        variableName: "test2"
      },
      {
        date: '2011',
        value: 5,
        variableName: "test2"
      },
      {
        date: '2012',
        value: 6,
        variableName: "test2"
      }
    ];
    var chart = LineChart(data, {
      x: d => d.date,
      y: d => d.value,
      z: d => d.variableName,
      yLabel: "↑ Smth",
      width: 500,
      height: 500,
      color: "steelblue",
      voronoi: false
    })

    return (
      <div dangerouslySetInnerHTML={{__html: chart.outerHTML}} />
    )
  }
}