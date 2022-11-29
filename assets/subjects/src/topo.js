import * as d3 from "d3";
import pl from './topo-poland.json';
import * as topojson from "topojson-client";

const stats = [
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
d3.select("#app").append('div').html(chart.outerHTML);

// Copyright 2021 Observable, Inc.
// Released under the ISC license.
// https://observablehq.com/@d3/choropleth
function Choropleth(data, {
  id = d => d.id, // given d in data, returns the feature id
  value = () => undefined, // given d in data, returns the quantitative value
  title, // given a feature f and possibly a datum d, returns the hover text
  format, // optional format specifier for the title
  scale = d3.scaleSequential, // type of color scale
  domain, // [min, max] values; input of color scale
  range = d3.interpolateBlues, // output of color scale
  width = 640, // outer width, in pixels
  height, // outer height, in pixels
  projection, // a D3 projection; null for pre-projected geometry
  features, // a GeoJSON feature collection
  featureId = d => d.id, // given a feature, returns its id
  borders, // a GeoJSON object for stroking borders
  outline = projection && projection.rotate ? {type: "Sphere"} : null, // a GeoJSON object for the background
  unknown = "#ccc", // fill color for missing data
  fill = "white", // fill color for outline
  stroke = "white", // stroke color for borders
  strokeLinecap = "round", // stroke line cap for borders
  strokeLinejoin = "round", // stroke line join for borders
  strokeWidth, // stroke width for borders
  strokeOpacity, // stroke opacity for borders
} = {}) {
  // Compute values.
  const N = d3.map(data, id);
  const V = d3.map(data, value).map(d => d == null ? NaN : +d);
  const Im = new d3.InternMap(N.map((id, i) => [id, i]));
  const If = d3.map(features.features, featureId);

  // Compute default domains.
  if (domain === undefined) domain = d3.extent(V);

  // Construct scales.
  const color = scale(domain, range);
  if (color.unknown && unknown !== undefined) color.unknown(unknown);

  // Compute titles.
  if (title === undefined) {
    format = color.tickFormat(100, format);
    title = (f, i) => `${f.properties.name}\n${format(V[i])}`;
  } else if (title !== null) {
    const T = title;
    const O = d3.map(data, d => d);
    title = (f, i) => T(f, O[i]);
  }

  // Compute the default height. If an outline object is specified, scale the projection to fit
  // the width, and then compute the corresponding height.
  if (height === undefined) {
    if (outline === undefined) {
      height = 400;
    } else {
      const [[x0, y0], [x1, y1]] = d3.geoPath(projection.fitWidth(width, outline)).bounds(outline);
      const dy = Math.ceil(y1 - y0), l = Math.min(Math.ceil(x1 - x0), dy);
      projection.scale(projection.scale() * (l - 1) / l).precision(0.2);
      height = dy;
    }
  }

  // Construct a path generator.
  const path = d3.geoPath(projection);

  const svg = d3.create("svg")
      .attr("width", width)
      .attr("height", height)
      .attr("viewBox", [0, 0, width, height])
      .attr("style", "width: 100%; height: auto; height: intrinsic;");

  if (outline != null) svg.append("path")
      .attr("fill", fill)
      .attr("stroke", "currentColor")
      .attr("d", path(outline));

  svg.append("g")
    .selectAll("path")
    .data(features.features)
    .join("path")
      .attr("fill", (d, i) => color(V[Im.get(If[i])]))
      .attr("d", path)
    .append("title")
      .text((d, i) => title(d, Im.get(If[i])));

  if (borders != null) svg.append("path")
      .attr("pointer-events", "none")
      .attr("fill", "none")
      .attr("stroke", stroke)
      .attr("stroke-linecap", strokeLinecap)
      .attr("stroke-linejoin", strokeLinejoin)
      .attr("stroke-width", strokeWidth)
      .attr("stroke-opacity", strokeOpacity)
      .attr("d", path(borders));

  return Object.assign(svg.node(), {scales: {color}});
}