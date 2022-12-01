import React, { Component } from 'react';
import Col from 'react-bootstrap/Col';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';

export default class Subjects extends Component {
  constructor(props) {
    super(props);

    this.state = {
      subjects: [],
      ancestors: [],
      subject: "Root"
    };
  }
  
  componentDidMount () {
    let url = 'http://localhost:3030/subjects';
    if (this.props.subjectId) {
      url = url + '/' + this.props.subjectId;
    }
    fetch(url).then(data => data.json())
      .then(data => this.setState({
        subjects: data.data.children,
        ancestors: data.data.ancestors,
        variables: !!data.data.links.$variables,
        subject: data.data.name
      }));
  }

  render() {
    const subjects = this.state.subjects.map((sbj, i) => (
      <a href={'/subjects/' + sbj.id} key={i} className="list-group-item list-group-item-action d-flex justify-content-between align-items-start">
        <div className="ms-2 me-auto">
          { sbj.name }
        </div>
        <span className="badge bg-primary rounded-pill">{ sbj.children_qty }</span>
      </a>
    ));
    const ancestors = this.state.ancestors.reverse().map((anc, i) => (
      <li className="breadcrumb-item"><a href={"/subjects/" + anc.id}>{ anc.name }</a></li>
    ));

    let content = <div>
      <h3>Children</h3>
        <ol className="list-group list-group-flush">{ subjects }</ol>
    </div>
    if (this.state.variables) {
      content = <div>
        <h3>Variables</h3>
      </div>
    }
    return (
      <div>
        <div className="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
          <h1 className="h2">Dashboard</h1>
        </div>
        <div>
          <nav aria-label="breadcrumb">
            <ol className="breadcrumb">
              { ancestors }
              <li className="breadcrumb-item active" aria-current="page">{ this.state.subject }</li>
            </ol>
          </nav>
          { content }
        </div>
      </div>
    );
  }
}