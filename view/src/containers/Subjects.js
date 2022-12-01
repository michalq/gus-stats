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

    return (
      <Container fluid className="main-content-container px-4">
        <Row className="page-header py-4">
          <h1 sm="4" className="text-sm-left">Subjects</h1>
        </Row>
        <Row>
          <Col md="9">
          <nav aria-label="breadcrumb">
            <ol className="breadcrumb">
              { ancestors }
              <li className="breadcrumb-item active" aria-current="page">{ this.state.subject }</li>
            </ol>
          </nav>
          <ol className="list-group list-group-flush">{ subjects }</ol>
          </Col>
        </Row>
      </Container>
    );
  }
}