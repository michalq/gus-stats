import React, { Component } from 'react';
import Tree from '../D3Tree';
import { Container, Row, Col } from 'react-bootstrap';

export default class SubjectsTree extends Component {
  constructor(props) {
    super(props);

    this.state = {
      tree: null
    };
  }

  componentDidMount () {
    fetch('http://localhost:3030/subjects/tree').then(data => data.json())
      .then(data => this.setState({
        tree: data.data
      }));
  }
  render() {
    let subjects = {children: []};
    if (this.state.tree != null) {
      subjects = this.state.tree;
    }
    const branches = [];
    for (const root of subjects.children) {
        const chart = Tree(root, {
            label: d => d.name,
            title: (d, n) => `${d.id}`,
            link: (d, n) => `/subjects/${d.id}`,
            width: 1000,
        });
        branches.push(<div dangerouslySetInnerHTML={{__html: chart.outerHTML}} />);
    }
    return (
      <div>
        <div className="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
          <h1 className="h2">Subjects tree</h1>
        </div>
        {branches}
      </div>
    )
  }
}