import React, { Component } from "react";
import pl from '../../lib/topo-poland.json';
import * as topojson from "topojson-client";
import { Choropleth } from "../../lib/Choropleth";
import * as d3 from "d3";

export default class Topo extends Component {
  render() {
    var stats = [
      {
      id: 1,
      province: "Łódzkie",
      rate: 0.2
      },
      {
      id: 2,
      province: "Świetokrzyskie",
      rate: 2.9
      },
      {
      id: 3,
      province: "Wielkopolskie",
      rate: 3
      },
      {
      id: 4,
      province: "Kujawsko-Pomorskie",
      rate: 4.5
      },
      {
      id: 5,
      province: "Małopolskie",
      rate: 5
      },
      {
      id: 6,
      province: "Dolnośląskie",
      rate: 6
      },
      {
      id: 7,
      province: "Lubelskie",
      rate: 7
      },
      {
      id: 8,
      province: "Lubuskie",
      rate: 8
      },
      {
      id: 9,
      province: "Mazowieckie",
      rate: 9
      },
      {
      id: 10,
      province: "Opolskie",
      rate: 10
      },
      {
      id: 11,
      province: "Podlaskie",
      rate: 12
      },
      {
      id: 12,
      province: "Pomorskie",
      rate: 12
      },
      {
      id: 13,
      province: "Śląskie",
      rate: 13
      },
      {
      id: 14,
      province: "Podkarpackie",
      rate: 14
      },
      {
      id: 15,
      province: "Warminsko-Mazurskie",
      rate: 15
      },
      {
      id: 16,
      province: "Zachodniopomorskie",
      rate: 16
      }
  ];

  const voivodeships = topojson.feature(pl, pl.objects.POL_adm1);
  console.log(voivodeships);
  const projection = d3.geoIdentity().reflectY(true).fitSize([600,600], voivodeships)
  // const statemap = new Map(voivodeships.features.map(d => [d.id, d]));
  // const statemesh = topojson.mesh(pl, pl.objects.POL_adm1, (a, b) => a !== b);
  const chart = Choropleth(stats, {
      id: d => d.id,
      featureId: f => { console.log(f.properties.ID_1, f.properties.VARNAME_1); return f.properties.ID_1 },
      value: d => d.rate,
      scale: d3.scaleQuantize,
      domain: [1, 16],
      range: d3.schemeBlues[9],
      title: (f, d) => `${f.properties.VARNAME_1}`,
      features: voivodeships,
      width: 600,
      height: 600,
      projection
    });
    
    return (
      <div dangerouslySetInnerHTML={{__html: chart.outerHTML}} />
    )
  }
}
